package agent

import (
	"context"
	"fmt"

	"github.com/code-certification/certify/internal/domain"
)

// Strategy determines how many pipeline stages to execute.
type Strategy int

const (
	// StrategyQuick runs prescreen only — cheap gate.
	StrategyQuick Strategy = iota
	// StrategyStandard runs prescreen → review → scoring (3 stages).
	StrategyStandard
	// StrategyFull runs all 5 stages including decision and remediation.
	StrategyFull
)

// PipelineConfig configures the review pipeline.
type PipelineConfig struct {
	Strategy Strategy
	Models   domain.ModelAssignments
}

// Pipeline orchestrates review stages in sequence.
type Pipeline struct {
	stages []Stage
}

// NewPipeline builds a pipeline based on strategy and model assignments.
func NewPipeline(provider Provider, cfg PipelineConfig) *Pipeline {
	var stages []Stage

	if cfg.Models.Prescreen != "" {
		stages = append(stages, NewPrescreenStage(provider, cfg.Models.Prescreen))
	}

	if cfg.Strategy >= StrategyStandard {
		if cfg.Models.Review != "" {
			stages = append(stages, NewReviewStage(provider, cfg.Models.Review))
		}
		if cfg.Models.Scoring != "" {
			stages = append(stages, NewScoringStage(provider, cfg.Models.Scoring))
		}
	}

	// Full strategy adds decision + remediation (not implemented as stages yet)

	return &Pipeline{stages: stages}
}

// Run executes the pipeline stages in order.
// Each stage can short-circuit (stop the pipeline).
// Errors cause graceful degradation, not failure.
func (p *Pipeline) Run(ctx context.Context, input StageInput) (ReviewResult, error) {
	var accumulated StageResult
	reviewed := false

	for _, stage := range p.stages {
		result, cont, err := stage.Execute(ctx, input)
		if err != nil {
			// Graceful degradation: stop pipeline, return what we have
			return p.toResult(accumulated, reviewed), nil
		}

		// Merge stage result into accumulated
		accumulated.TokensUsed += result.TokensUsed
		if result.ReviewOutput != "" {
			accumulated.ReviewOutput = result.ReviewOutput
			input.ReviewOutput = result.ReviewOutput // feed forward
			reviewed = true
		}
		if result.Scores != nil {
			accumulated.Scores = result.Scores
		}
		if result.Confidence > accumulated.Confidence {
			accumulated.Confidence = result.Confidence
		}
		if result.Status != "" {
			accumulated.Status = result.Status
		}
		if result.Actions != nil {
			accumulated.Actions = result.Actions
		}

		if !cont {
			break
		}
	}

	return p.toResult(accumulated, reviewed), nil
}

func (p *Pipeline) toResult(sr StageResult, reviewed bool) ReviewResult {
	return ReviewResult{
		Reviewed:     reviewed,
		ReviewOutput: sr.ReviewOutput,
		Scores:       sr.Scores,
		Status:       sr.Status,
		Actions:      sr.Actions,
		Remediation:  sr.Remediation,
		Confidence:   sr.Confidence,
		TokensUsed:   sr.TokensUsed,
	}
}

// --- Coordinator: batch orchestration ---

// CoordinatorConfig configures the review coordinator.
type CoordinatorConfig struct {
	Models      domain.ModelAssignments
	Strategy    Strategy
	TokenBudget int // max tokens to spend across all reviews (0=unlimited)
}

// Coordinator manages file-level dedup, token budgets, and pipeline execution.
type Coordinator struct {
	provider    Provider
	config      CoordinatorConfig
	reviewed    map[string]ReviewResult // file path → cached result
	tokensSpent int
}

// NewCoordinator creates a review coordinator.
func NewCoordinator(provider Provider, cfg CoordinatorConfig) *Coordinator {
	return &Coordinator{
		provider: provider,
		config:   cfg,
		reviewed: make(map[string]ReviewResult),
	}
}

// ReviewUnit returns the review result for a unit, deduplicating by file.
func (c *Coordinator) ReviewUnit(ctx context.Context, unit domain.Unit, source string, ev []domain.Evidence) ReviewResult {
	filePath := unit.ID.Path()

	// Check file-level cache
	if cached, ok := c.reviewed[filePath]; ok {
		return cached
	}

	// Check token budget
	if c.config.TokenBudget > 0 && c.tokensSpent >= c.config.TokenBudget {
		result := ReviewResult{Reviewed: false}
		c.reviewed[filePath] = result
		return result
	}

	// Build evidence summary
	var summaries []string
	for _, e := range ev {
		summaries = append(summaries, e.Summary)
	}
	evSummary := ""
	if len(summaries) > 0 {
		evSummary = fmt.Sprintf("%v", summaries)
	}

	// Build and run pipeline
	pipeline := NewPipeline(c.provider, PipelineConfig{
		Strategy: c.config.Strategy,
		Models:   c.config.Models,
	})

	result, err := pipeline.Run(ctx, StageInput{
		Unit:            unit,
		SourceCode:      source,
		Evidence:        ev,
		EvidenceSummary: evSummary,
	})
	if err != nil {
		result = ReviewResult{Reviewed: false}
	}

	c.tokensSpent += result.TokensUsed
	c.reviewed[filePath] = result
	return result
}

// Stats returns coordinator statistics.
func (c *Coordinator) Stats() (filesReviewed, totalFiles, tokensSpent int) {
	reviewed := 0
	for _, r := range c.reviewed {
		if r.Reviewed {
			reviewed++
		}
	}
	return reviewed, len(c.reviewed), c.tokensSpent
}

package agent

import "github.com/code-certification/certify/internal/domain"

// Router maps task types to specific models with fallback.
type Router struct {
	assignments domain.ModelAssignments
}

// NewRouter creates a model router.
func NewRouter(assignments domain.ModelAssignments) *Router {
	return &Router{assignments: assignments}
}

// ModelFor returns the model ID for the given task type.
// Falls back to the fallback model if no specific assignment exists.
func (r *Router) ModelFor(task TaskType) string {
	model := r.directAssignment(task)
	if model != "" {
		return model
	}
	return r.assignments.Fallback
}

func (r *Router) directAssignment(task TaskType) string {
	switch task {
	case TaskPrescreen:
		return r.assignments.Prescreen
	case TaskReview:
		return r.assignments.Review
	case TaskScoring:
		return r.assignments.Scoring
	case TaskDecision:
		return r.assignments.Decision
	case TaskRemediation:
		return r.assignments.Remediation
	default:
		return ""
	}
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iksnae/code-certification/internal/discovery"
	gh "github.com/iksnae/code-certification/internal/github"
	"github.com/spf13/cobra"
)

var (
	initPath string
	initPR   bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap certification in the current repository",
	Long:  "Creates .certification/ directory with config, policies, records, overrides, and GitHub workflows.",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := initPath
		if root == "" {
			root, _ = os.Getwd()
		}
		certDir := filepath.Join(root, ".certification")

		// Create directory structure
		dirs := []string{
			certDir,
			filepath.Join(certDir, "policies"),
			filepath.Join(certDir, "records"),
			filepath.Join(certDir, "overrides"),
		}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0o755); err != nil {
				return fmt.Errorf("creating %s: %w", d, err)
			}
		}

		// Detect languages
		fmt.Println("  Detecting languages...")
		langs := discovery.DetectLanguages(root)
		if len(langs) > 0 {
			for _, l := range langs {
				cfg := ""
				if l.HasConfig {
					cfg = " (config found)"
				}
				fmt.Printf("  • %s: %d files%s\n", l.Name, l.FileCount, cfg)
			}
		} else {
			fmt.Println("  • No recognized languages detected")
		}

		// Write config
		configPath := filepath.Join(certDir, "config.yml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			cfg := generateConfig(langs)
			if err := os.WriteFile(configPath, []byte(cfg), 0o644); err != nil {
				return fmt.Errorf("writing config: %w", err)
			}
			fmt.Println("  ✓ Created .certification/config.yml")
		} else {
			fmt.Println("  • .certification/config.yml already exists")
		}

		// Write global policy
		policyPath := filepath.Join(certDir, "policies", "global.yml")
		if _, err := os.Stat(policyPath); os.IsNotExist(err) {
			if err := os.WriteFile(policyPath, []byte(globalPolicy), 0o644); err != nil {
				return fmt.Errorf("writing policy: %w", err)
			}
			fmt.Println("  ✓ Created .certification/policies/global.yml")
		}

		// Generate language-specific policies
		for _, l := range langs {
			policy := languagePolicy(l.Name)
			if policy == "" {
				continue
			}
			name := l.Name + ".yml"
			path := filepath.Join(certDir, "policies", name)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				if err := os.WriteFile(path, []byte(policy), 0o644); err != nil {
					return fmt.Errorf("writing %s policy: %w", l.Name, err)
				}
				fmt.Printf("  ✓ Created .certification/policies/%s\n", name)
			}
		}

		// Generate GitHub workflow files
		workflowDir := filepath.Join(root, ".github", "workflows")
		if err := os.MkdirAll(workflowDir, 0o755); err == nil {
			workflows := map[string]string{
				"certification-pr.yml":      gh.GeneratePRWorkflow(),
				"certification-nightly.yml": gh.GenerateNightlyWorkflow(),
				"certification-weekly.yml":  gh.GenerateWeeklyWorkflow(),
			}
			for name, content := range workflows {
				path := filepath.Join(workflowDir, name)
				if _, err := os.Stat(path); os.IsNotExist(err) {
					if err := os.WriteFile(path, []byte(content), 0o644); err == nil {
						fmt.Printf("  ✓ Created .github/workflows/%s\n", name)
					}
				}
			}
		}

		// Summary of what was created
		summary := fmt.Sprintf("## Certification Bootstrap Summary\n\n"+
			"**Languages detected**: %d\n", len(langs))
		for _, l := range langs {
			summary += fmt.Sprintf("- %s (%d files)\n", l.Name, l.FileCount)
		}
		summary += "\n**Generated files**:\n" +
			"- `.certification/config.yml` — Mode: advisory\n" +
			"- `.certification/policies/` — Starter policy packs\n" +
			"- `.github/workflows/` — PR, nightly, weekly workflows\n\n" +
			"**Next steps**:\n" +
			"1. Review and customize `.certification/config.yml`\n" +
			"2. Run `certify scan` to discover code units\n" +
			"3. Run `certify certify` to evaluate\n" +
			"4. Run `certify report` to see results\n"

		if initPR {
			// Create a branch and PR
			branchName := "certify/bootstrap"
			fmt.Printf("\n  Creating PR on branch %s...\n", branchName)

			cmds := [][]string{
				{"git", "checkout", "-b", branchName},
				{"git", "add", ".certification/", ".github/workflows/"},
				{"git", "commit", "-m", "chore: bootstrap code certification"},
				{"git", "push", "-u", "origin", branchName},
			}
			for _, c := range cmds {
				cmd := exec.Command(c[0], c[1:]...)
				cmd.Dir = root
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("git command failed: %w", err)
				}
			}

			prCmd := exec.Command("gh", "pr", "create",
				"--title", "chore: bootstrap code certification",
				"--body", summary,
				"--base", "main",
			)
			prCmd.Dir = root
			prCmd.Stdout = os.Stdout
			prCmd.Stderr = os.Stderr
			if err := prCmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "warning: gh pr create failed: %v\n", err)
				fmt.Println("  Files committed to branch. Create PR manually.")
			}
		} else {
			fmt.Print(summary)
		}

		fmt.Println("\n✓ Certification initialized. Run 'certify scan' to discover code units.")
		return nil
	},
}

func init() {
	initCmd.Flags().StringVar(&initPath, "path", "", "Path to repository (default: current directory)")
	initCmd.Flags().BoolVar(&initPR, "pr", false, "Create initialization as a pull request")
}

func generateConfig(langs []discovery.LanguageInfo) string {
	cfg := `# Certify — Configuration
mode: advisory

scope:
  include: []
  exclude:
    - "vendor/**"
    - "node_modules/**"
    - "dist/**"
    - "build/**"
    - "testdata/**"
    - "**/*_test.go"
    - "**/*.test.ts"
    - "**/*.spec.ts"

agent:
  # Auto-detection: when OPENROUTER_API_KEY (or CERTIFY_API_KEY) is set in
  # your environment or as a GitHub secret, Certify automatically enables
  # conservative AI-assisted review (prescreen-only, free-tier models,
  # 10k token budget). No configuration needed — just add the secret.
  #
  # To use the full agent pipeline, set enabled: true and configure models:
  # enabled: true
  # provider:
  #   type: openrouter
  #   base_url: https://openrouter.ai/api/v1
  #   api_key_env: OPENROUTER_API_KEY
  # models:
  #   prescreen: qwen/qwen3-coder:free
  #   review: qwen/qwen3-coder:free
  #   scoring: qwen/qwen3-coder:free
  #
  # To explicitly disable AI even when a key is present, uncomment:
  # enabled: false

expiry:
  default_window_days: 90
  min_window_days: 7
  max_window_days: 365

issues:
  enabled: false
`
	return cfg
}

const globalPolicy = `name: global
version: "1.0.0"

rules:
  - id: lint-clean
    dimension: correctness
    description: "Code must pass linter with zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: test-pass
    dimension: testability
    description: "All tests must pass"
    severity: error
    metric: test_failures
    threshold: 0
`

func languagePolicy(lang string) string {
	switch lang {
	case "go":
		return `name: go-standard
version: "1.0.0"
language: go

rules:
  - id: go-vet-clean
    dimension: correctness
    description: "go vet must report zero issues"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: go-test-pass
    dimension: testability
    description: "All Go tests must pass"
    severity: error
    metric: test_failures
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
`
	case "typescript", "javascript":
		return `name: ts-standard
version: "1.0.0"
language: ts

rules:
  - id: eslint-clean
    dimension: correctness
    description: "ESLint must report zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
`
	case "python":
		return `name: python-standard
version: "1.0.0"
language: py

rules:
  - id: lint-clean
    dimension: correctness
    description: "Linter must report zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
`
	default:
		return ""
	}
}

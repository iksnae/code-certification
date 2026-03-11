package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iksnae/code-certification/internal/discovery"
	gh "github.com/iksnae/code-certification/internal/github"
	"github.com/iksnae/code-certification/internal/workspace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap certification in the current repository",
	Long:  "Creates .certification/ directory with config, policies, records, overrides, and GitHub workflows.",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := flagString(cmd, "path")
		if root == "" {
			var err error
			root, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}
		}

		wsMode := flagBool(cmd, "workspace")
		if wsMode {
			return runWorkspaceInit(root)
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

		initPR := flagBool(cmd, "pr")
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

func bindInitFlags() {
	initCmd.Flags().String("path", "", "Path to repository (default: current directory)")
	initCmd.Flags().Bool("pr", false, "Create initialization as a pull request")
}

func runWorkspaceInit(root string) error {
	fmt.Println("🔍 Workspace mode: discovering git submodules...")

	subs, err := workspace.DiscoverSubmodules(root)
	if err != nil {
		return fmt.Errorf("discovering submodules: %w", err)
	}

	if len(subs) == 0 {
		return fmt.Errorf("no git submodules found in %s", root)
	}

	fmt.Printf("  Found %d submodule(s):\n\n", len(subs))

	var unconfigured []workspace.Submodule
	for _, s := range subs {
		status := "✓ certify configured"
		if !s.HasConfig {
			status = "✗ no certify setup"
			unconfigured = append(unconfigured, s)
		}
		fmt.Printf("  %-40s %s\n", s.Path, status)
	}
	fmt.Println()

	if len(unconfigured) == 0 {
		fmt.Println("✓ All submodules have certify configured.")
		return nil
	}

	fmt.Printf("  Initializing %d unconfigured submodule(s)...\n\n", len(unconfigured))
	for _, s := range unconfigured {
		fmt.Printf("  → %s\n", s.Path)
		subRoot := filepath.Join(root, s.Path)
		if err := runSubcommand("init", "--path", subRoot); err != nil {
			fmt.Fprintf(os.Stderr, "    warning: init failed for %s: %v\n", s.Path, err)
		}
	}

	// Create workspace-level .certification directory
	certDir := filepath.Join(root, ".certification")
	os.MkdirAll(certDir, 0o755)

	fmt.Println("\n✓ Workspace initialized. Run 'certify scan --workspace' next.")
	return nil
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
		return goStandardPolicy
	case "typescript", "javascript":
		return tsStandardPolicy
	case "python":
		return pythonStandardPolicy
	case "rust":
		return rustStandardPolicy
	case "swift":
		return swiftStandardPolicy
	default:
		return ""
	}
}

const goStandardPolicy = `name: go-standard
version: "1.3.0"
language: go

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

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
    exclude_patterns: ["*_test.go"]

  - id: max-complexity
    dimension: maintainability
    description: "Cyclomatic complexity should be manageable"
    severity: warning
    metric: complexity
    threshold: 20

  - id: max-params
    dimension: maintainability
    description: "Functions should have at most 5 parameters"
    severity: warning
    metric: param_count
    threshold: 5

  - id: max-nesting
    dimension: readability
    description: "Function body nesting should not exceed 4 levels"
    severity: warning
    metric: max_nesting_depth
    threshold: 4

  - id: no-ignored-errors
    dimension: correctness
    description: "Error returns should not be silently discarded"
    severity: warning
    metric: errors_ignored
    threshold: 0

  - id: max-func-lines
    dimension: readability
    description: "Functions should be concise (under 100 lines)"
    severity: warning
    metric: func_lines
    threshold: 100

  - id: no-defer-in-loop
    dimension: correctness
    description: "defer inside loops causes resource leaks"
    severity: error
    metric: defer_in_loop
    threshold: 0

  - id: context-first-param
    dimension: correctness
    description: "context.Context should be the first parameter"
    severity: warning
    metric: context_not_first
    threshold: 0

  - id: no-init-func
    dimension: maintainability
    description: "Avoid init() functions — prefer explicit initialization"
    severity: warning
    metric: has_init_func
    threshold: 0

  - id: limit-global-mutable
    dimension: security
    description: "Minimize package-level mutable variables"
    severity: warning
    metric: global_mutable_count
    threshold: 2

  - id: max-methods
    dimension: maintainability
    description: "Types should not have more than 15 methods (god object)"
    severity: warning
    metric: method_count
    threshold: 15

  - id: max-cognitive-complexity
    dimension: readability
    description: "Cognitive complexity should not exceed 25"
    severity: warning
    metric: cognitive_complexity
    threshold: 25

  - id: no-unsafe-imports
    dimension: security
    description: "Avoid unsafe, os/exec, and similar dangerous imports"
    severity: warning
    metric: unsafe_import_count
    threshold: 0

  - id: no-hardcoded-secrets
    dimension: security
    description: "No hardcoded passwords, API keys, or secrets in source"
    severity: critical
    metric: hardcoded_secrets
    threshold: 0

  - id: wrap-errors
    dimension: operational_quality
    description: "Errors should be wrapped with context using fmt.Errorf(%%w)"
    severity: warning
    metric: errors_not_wrapped
    threshold: 0

  - id: no-panic-in-library
    dimension: correctness
    description: "Library code should return errors, not panic"
    severity: warning
    metric: panic_calls
    threshold: 0

  - id: no-os-exit
    dimension: testability
    description: "Avoid os.Exit() — makes code untestable"
    severity: warning
    metric: os_exit_calls
    threshold: 0
`

const tsStandardPolicy = `name: ts-standard
version: "1.2.0"
language: ts

rules:
  - id: lint-clean
    dimension: correctness
    description: "Linter must report zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: test-pass
    dimension: testability
    description: "All tests must pass"
    severity: error
    metric: test_failures
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
    exclude_patterns: ["*.test.ts", "*.spec.ts", "*.test.tsx", "*.spec.tsx"]

  - id: max-complexity
    dimension: maintainability
    description: "Cyclomatic complexity should be manageable"
    severity: warning
    metric: complexity
    threshold: 20

  - id: max-params
    dimension: maintainability
    description: "Functions should have at most 5 parameters"
    severity: warning
    metric: param_count
    threshold: 5

  - id: max-nesting
    dimension: readability
    description: "Function body nesting should not exceed 4 levels"
    severity: warning
    metric: max_nesting_depth
    threshold: 4

  - id: max-func-lines
    dimension: readability
    description: "Functions should be concise (under 100 lines)"
    severity: warning
    metric: func_lines
    threshold: 100

  - id: max-cognitive-complexity
    dimension: readability
    description: "Cognitive complexity should not exceed 25"
    severity: warning
    metric: cognitive_complexity
    threshold: 25

  - id: no-unsafe-imports
    dimension: security
    description: "Avoid child_process, eval, vm, and similar dangerous imports"
    severity: warning
    metric: unsafe_import_count
    threshold: 0

  - id: no-hardcoded-secrets
    dimension: security
    description: "No hardcoded passwords, API keys, or secrets in source"
    severity: critical
    metric: hardcoded_secrets
    threshold: 0

  - id: no-empty-catch
    dimension: correctness
    description: "Catch blocks should not be empty"
    severity: warning
    metric: empty_catch_blocks
    threshold: 0
`

const pythonStandardPolicy = `name: python-standard
version: "1.2.0"
language: py

rules:
  - id: lint-clean
    dimension: correctness
    description: "Linter must report zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: test-pass
    dimension: testability
    description: "All tests must pass"
    severity: error
    metric: test_failures
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
    exclude_patterns: ["test_*.py", "*_test.py"]

  - id: max-complexity
    dimension: maintainability
    description: "Cyclomatic complexity should be manageable"
    severity: warning
    metric: complexity
    threshold: 20

  - id: max-params
    dimension: maintainability
    description: "Functions should have at most 5 parameters (excluding self/cls)"
    severity: warning
    metric: param_count
    threshold: 5

  - id: max-nesting
    dimension: readability
    description: "Function body nesting should not exceed 4 levels"
    severity: warning
    metric: max_nesting_depth
    threshold: 4

  - id: max-func-lines
    dimension: readability
    description: "Functions should be concise (under 100 lines)"
    severity: warning
    metric: func_lines
    threshold: 100

  - id: max-cognitive-complexity
    dimension: readability
    description: "Cognitive complexity should not exceed 25"
    severity: warning
    metric: cognitive_complexity
    threshold: 25

  - id: no-unsafe-imports
    dimension: security
    description: "Avoid subprocess, pickle, ctypes, and similar dangerous imports"
    severity: warning
    metric: unsafe_import_count
    threshold: 0

  - id: no-hardcoded-secrets
    dimension: security
    description: "No hardcoded passwords, API keys, or secrets in source"
    severity: critical
    metric: hardcoded_secrets
    threshold: 0

  - id: no-empty-except
    dimension: correctness
    description: "Bare except:pass blocks silently swallow errors"
    severity: warning
    metric: empty_catch_blocks
    threshold: 0
`

const rustStandardPolicy = `name: rust-standard
version: "1.2.0"
language: rs

rules:
  - id: lint-clean
    dimension: correctness
    description: "Clippy must report zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: test-pass
    dimension: testability
    description: "All tests must pass"
    severity: error
    metric: test_failures
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0

  - id: max-complexity
    dimension: maintainability
    description: "Cyclomatic complexity should be manageable"
    severity: warning
    metric: complexity
    threshold: 20

  - id: max-params
    dimension: maintainability
    description: "Functions should have at most 5 parameters (excluding self)"
    severity: warning
    metric: param_count
    threshold: 5

  - id: max-nesting
    dimension: readability
    description: "Function body nesting should not exceed 4 levels"
    severity: warning
    metric: max_nesting_depth
    threshold: 4

  - id: max-func-lines
    dimension: readability
    description: "Functions should be concise (under 100 lines)"
    severity: warning
    metric: func_lines
    threshold: 100

  - id: max-cognitive-complexity
    dimension: readability
    description: "Cognitive complexity should not exceed 25"
    severity: warning
    metric: cognitive_complexity
    threshold: 25

  - id: no-unsafe-imports
    dimension: security
    description: "Avoid std::process, libc, and similar dangerous imports"
    severity: warning
    metric: unsafe_import_count
    threshold: 0

  - id: no-hardcoded-secrets
    dimension: security
    description: "No hardcoded passwords, API keys, or secrets in source"
    severity: critical
    metric: hardcoded_secrets
    threshold: 0

  - id: no-panic-unwrap
    dimension: correctness
    description: "Avoid panic!() and .unwrap() — use Result/Option properly"
    severity: warning
    metric: panic_calls
    threshold: 0
`

const swiftStandardPolicy = `name: swift-standard
version: "1.0.0"
language: swift

rules:
  - id: lint-clean
    dimension: correctness
    description: "SwiftLint must report zero errors"
    severity: error
    metric: lint_errors
    threshold: 0

  - id: test-pass
    dimension: testability
    description: "All tests must pass"
    severity: error
    metric: test_failures
    threshold: 0

  - id: no-todos
    dimension: readability
    description: "No TODO/FIXME comments in certified code"
    severity: warning
    metric: todo_count
    threshold: 0
    exclude_patterns: ["*Tests.swift", "*Test.swift"]

  - id: max-complexity
    dimension: maintainability
    description: "Cyclomatic complexity should be manageable"
    severity: warning
    metric: complexity
    threshold: 20
`

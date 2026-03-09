package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap certification in the current repository",
	Long:  "Creates .certification/ directory with default config, policies, and record storage.",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, _ := os.Getwd()
		certDir := filepath.Join(root, ".certification")

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

		// Write default config
		configPath := filepath.Join(certDir, "config.yml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			if err := os.WriteFile(configPath, []byte(defaultConfig), 0o644); err != nil {
				return fmt.Errorf("writing config: %w", err)
			}
			fmt.Println("  ✓ Created .certification/config.yml")
		} else {
			fmt.Println("  • .certification/config.yml already exists")
		}

		// Write default global policy
		policyPath := filepath.Join(certDir, "policies", "global.yml")
		if _, err := os.Stat(policyPath); os.IsNotExist(err) {
			if err := os.WriteFile(policyPath, []byte(defaultGlobalPolicy), 0o644); err != nil {
				return fmt.Errorf("writing policy: %w", err)
			}
			fmt.Println("  ✓ Created .certification/policies/global.yml")
		}

		fmt.Println("\n✓ Certification initialized. Run 'certify scan' to discover code units.")
		return nil
	},
}

const defaultConfig = `# Code Certification System Configuration
mode: advisory

scope:
  include: []
  exclude:
    - "vendor/**"
    - "node_modules/**"
    - "testdata/**"
    - "**/*_test.go"
    - "**/*.test.ts"

agent:
  enabled: false

expiry:
  default_window_days: 90
  min_window_days: 7
  max_window_days: 365
`

const defaultGlobalPolicy = `name: global
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

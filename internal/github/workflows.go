// Package github handles GitHub integration (workflows, PR annotations, issues).
package github

// GeneratePRWorkflow returns the YAML for a PR certification workflow.
func GeneratePRWorkflow() string {
	return `name: Certification PR Review
on:
  pull_request:
    branches: [main]

permissions:
  contents: read
  pull-requests: write
  issues: write

jobs:
  certify:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install certify
        run: go install github.com/iksnae/code-certification/cmd/certify@latest

      - name: Scan code units
        run: certify scan

      - name: Run certification
        run: certify certify
        env:
          OPENROUTER_API_KEY: ${{ secrets.OPENROUTER_API_KEY }}

      - name: Generate report
        run: certify report --format json > certification-report.json

      - name: Comment on PR
        if: always()
        run: certify review
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          OPENROUTER_API_KEY: ${{ secrets.OPENROUTER_API_KEY }}
`
}

// GenerateNightlyWorkflow returns the YAML for a nightly certification sweep.
func GenerateNightlyWorkflow() string {
	return `name: Certification Nightly Sweep
on:
  schedule:
    - cron: '0 2 * * *'  # 2 AM UTC daily
  workflow_dispatch:

permissions:
  contents: write
  issues: write

jobs:
  sweep:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install certify
        run: go install github.com/iksnae/code-certification/cmd/certify@latest

      - name: Scan code units
        run: certify scan

      - name: Run full certification
        run: certify certify
        env:
          OPENROUTER_API_KEY: ${{ secrets.OPENROUTER_API_KEY }}

      - name: Generate report
        run: certify report --format text

      - name: Commit updated records
        run: |
          git config user.name "certify[bot]"
          git config user.email "certify@users.noreply.github.com"
          git add .certification/
          git diff --staged --quiet || git commit -m "chore: nightly certification sweep"
          git push
`
}

// GenerateWeeklyWorkflow returns the YAML for a weekly certification report.
func GenerateWeeklyWorkflow() string {
	return `name: Certification Weekly Report
on:
  schedule:
    - cron: '0 8 * * 1'  # 8 AM UTC every Monday
  workflow_dispatch:

permissions:
  contents: read
  issues: write

jobs:
  report:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install certify
        run: go install github.com/iksnae/code-certification/cmd/certify@latest

      - name: Generate weekly report
        run: certify report --format text

      - name: Generate JSON report
        run: certify report --format json > weekly-report.json
`
}

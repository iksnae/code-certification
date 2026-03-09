# Changelog

## 0.1.1 — 2026-03-09

### Fixed
- **Scanner excludes non-code files**: Images, CSS, JSON, HTML, build artifacts, lock files, fonts, and binaries are no longer discovered or certified. Only certifiable source code extensions (35 types across 20+ languages) are scanned.
- **Dashboard offline fallback**: Dashboard now builds report from on-disk records when the CLI binary isn't in the extension's PATH.

### Added
- **OpenAI provider support**: `OPENAI_API_KEY` auto-detected. Default models: gpt-4o-mini, gpt-4o, gpt-4.1-mini, gpt-4.1-nano, o3-mini.

## 0.1.0 — 2026-03-09

### Initial Release

- **Dashboard**: Interactive report card WebView — overall grade, grade distribution, quality dimension bars, language breakdown, filterable unit table
- **Tree View**: Explorer sidebar showing certified units organized by directory
- **CodeLens**: Inline grade annotations on Go and TypeScript functions
- **Status Bar**: Persistent grade badge — click to open dashboard
- **Diagnostics**: Warnings for Grade D/F units, info markers for expiring-soon certifications
- **Provider Configuration**: Visual setup for any OpenAI-compatible provider — 11 presets + custom URL
- **Model Browser**: Dynamic model discovery from any provider via `certify models`
- **Commands**: 8 command palette actions — scan, certify, report, dashboard, configure, browse models, install CLI, open unit

#!/usr/bin/env bash
# Mutation battery for the unassessed-unit denominator invariants.
#
# Each mutant reintroduces one specific defect this branch fixed. A mutant that
# SURVIVES means the corresponding test does not actually constrain the code it
# claims to.
#
# Two rules this harness exists to enforce, because the earlier throwaway
# versions of it broke both:
#
#   1. BASELINE GREEN GUARD. The unmutated suite is run first, and nothing else
#      runs unless it is green. Without it, a battery reports kills on a tree
#      that was already failing, and every "kill" is meaningless.
#
#   2. A COMPILE FAILURE IS NOT A KILL. A mutant that does not build is a broken
#      mutant, reported as BUILD-FAIL and counted against the run. Scoring it as
#      a kill is the worst failure mode a battery has: it reports a mutant dead
#      while it is in fact untested, which inflates the score in exactly the
#      direction that flatters the author.
#
# Usage: scripts/mutation-test.sh [-v]
# Exit:  0 all mutants killed · 1 a mutant survived or failed to build
#        2 baseline was not green (nothing was measured)
set -uo pipefail

REPO_ROOT="$(git -C "$(dirname "${BASH_SOURCE[0]}")" rev-parse --show-toplevel)"
cd "$REPO_ROOT"

VERBOSE=0
[[ "${1:-}" == "-v" ]] && VERBOSE=1

# Mutant table: NAME | FILE | ORIGINAL | MUTATED | TEST_PACKAGE
# ORIGINAL must occur exactly once in FILE, or the mutant is reported invalid.
MUTANTS=(
"run-grade-second-aggregation|cmd/certify/certify_cmd.go|card := report.GenerateCard(records, \"\", p.commit, time.Now())|card := report.Card{OverallScore: 0.4938271604938272, OverallGrade: \"F\"}; _ = records|./cmd/certify/"
"store-skip-legacy-normalization|internal/record/store.go|	if rec.Unsupported {\n		rec = rec.WithUnassessedVerdict()\n	}|	if false {\n		rec = rec.WithUnassessedVerdict()\n	}|./internal/record/ ./internal/report/"
"unassessed-verdict-keeps-grade|internal/domain/record.go|		Grade:      GradeNA,|		Grade:      GradeF,|./internal/record/ ./internal/engine/"
"card-score-denominator|internal/report/card.go|		c.OverallScore = totalScore / float64(c.AnalyzableUnits)|		c.OverallScore = totalScore / float64(c.TotalUnits)|./internal/report/ ./cmd/certify/"
"card-passrate-denominator|internal/report/card.go|		c.PassRate = float64(c.Passing) / float64(c.AnalyzableUnits)|		c.PassRate = float64(c.Passing) / float64(c.TotalUnits)|./internal/report/"
"badge-drops-population|internal/report/badge.go|		c.OverallGrade, c.PassRate*100, FormatUnitPopulation(c.TotalUnits, c.UnsupportedCount))|		c.OverallGrade, c.PassRate*100, fmt.Sprintf(\"%d units\", c.TotalUnits))|./internal/report/"
"unit-population-drops-qualifier|internal/report/card.go|	return fmt.Sprintf(\"%d of %d units analyzable\", total-unsupported, total)|	return fmt.Sprintf(\"%d units\", total)|./internal/report/"
"agent-avgscore-denominator|internal/agent/architect_snapshot.go|		snap.Metrics.AvgScore = totalScore / float64(analyzable)|		snap.Metrics.AvgScore = totalScore / float64(len(records))|./internal/agent/"
"agent-package-counts-unassessed|internal/agent/architect_snapshot.go|		if r.Unsupported {\n			a.unsupported++\n		} else {\n			a.scores = append(a.scores, r.Score)\n		}|		a.scores = append(a.scores, r.Score)|./internal/agent/"
"hotspot-risk-over-all-units|internal/agent/architect.go|		analyzable := h.Units - h.Unsupported|		analyzable := h.Units|./internal/agent/"
"workspace-weight-over-all-units|internal/agent/workspace_snapshot.go|		units := subSnap.Metrics.TotalUnits - subSnap.Metrics.UnitsUnsupported|		units := subSnap.Metrics.TotalUnits|./internal/agent/"
"recurrent-areas-count-unassessed|internal/report/detailed.go|		if r.Unsupported {\n			continue\n		}\n		dir := filepath.Dir(r.UnitPath)|		dir := filepath.Dir(r.UnitPath)|./internal/report/"
"report-tree-stats-denominator|internal/report/report_tree.go|		s.avgScore = totalScore / float64(s.analyzable())|		s.avgScore = totalScore / float64(len(units))|./internal/report/"
"workspace-aggregate-weight|internal/workspace/aggregate.go|		totalWeightedScore += s.Score * float64(s.Analyzable())|		totalWeightedScore += s.Score * float64(s.Units)|./internal/workspace/"
"workspace-aggregate-divisor|internal/workspace/aggregate.go|		wc.OverallScore = totalWeightedScore / float64(wc.AnalyzableUnits)|		wc.OverallScore = totalWeightedScore / float64(wc.AnalyzableUnits+wc.TotalUnsupported)|./internal/workspace/"
"language-detail-denominator|internal/report/full.go|			avg = sum / float64(len(a.scores))|			avg = sum / float64(a.units)|./internal/report/"
)

pass=0
survived=0
buildfail=0
invalid=0
declare -a SURVIVED_LIST=() BUILDFAIL_LIST=() INVALID_LIST=()

restore_all() { git checkout -- . 2>/dev/null || true; }
trap restore_all EXIT

if ! git diff --quiet || ! git diff --cached --quiet; then
	echo "REFUSING: working tree is dirty. Commit or stash first — this harness"
	echo "restores files with 'git checkout --' and would discard your changes."
	exit 2
fi

echo "=============================================="
echo " BASELINE GREEN GUARD"
echo "=============================================="
if ! go build ./... >/dev/null 2>&1; then
	echo "BASELINE NOT GREEN: tree does not build. Nothing measured."
	exit 2
fi
baseline_log="$(mktemp)"
if ! go test ./... >"$baseline_log" 2>&1; then
	echo "BASELINE NOT GREEN: unmutated suite fails. Nothing measured."
	echo "--- baseline failures ---"
	grep -E '^(---|FAIL)' "$baseline_log" | head -30
	rm -f "$baseline_log"
	exit 2
fi
rm -f "$baseline_log"
echo "baseline: unmutated suite GREEN — mutants are now meaningful"
echo

for entry in "${MUTANTS[@]}"; do
	[[ -z "$entry" ]] && continue
	IFS='|' read -r name file original mutated pkgs <<<"$entry"

	# Interpret \n in the table as real newlines.
	original="$(printf '%b' "$original")"
	mutated="$(printf '%b' "$mutated")"

	if [[ ! -f "$file" ]]; then
		echo "?? $name — INVALID (no such file: $file)"
		invalid=$((invalid + 1))
		INVALID_LIST+=("$name")
		continue
	fi

	occurrences="$(python3 - "$file" "$original" <<'PY'
import sys
src = open(sys.argv[1]).read()
print(src.count(sys.argv[2]))
PY
)"
	if [[ "$occurrences" != "1" ]]; then
		echo "?? $name — INVALID (original text occurs $occurrences times, need exactly 1)"
		invalid=$((invalid + 1))
		INVALID_LIST+=("$name")
		continue
	fi

	python3 - "$file" "$original" "$mutated" <<'PY'
import sys
path, old, new = sys.argv[1], sys.argv[2], sys.argv[3]
src = open(path).read()
open(path, "w").write(src.replace(old, new, 1))
PY

	# A mutant must compile before its test result means anything.
	if ! go build ./... >/dev/null 2>&1; then
		echo "!! $name — BUILD-FAIL (mutant did not compile; NOT a kill)"
		buildfail=$((buildfail + 1))
		BUILDFAIL_LIST+=("$name")
		git checkout -- "$file"
		continue
	fi

	out="$(go test $pkgs 2>&1)"
	rc=$?
	git checkout -- "$file"

	if [[ $rc -ne 0 ]]; then
		echo "ok $name — KILLED"
		[[ $VERBOSE -eq 1 ]] && echo "$out" | grep -E '^    .*_test\.go:' | head -3
		pass=$((pass + 1))
	else
		echo "XX $name — SURVIVED (no test constrains this line)"
		survived=$((survived + 1))
		SURVIVED_LIST+=("$name")
	fi
done

total=$((pass + survived + buildfail + invalid))
echo
echo "=============================================="
echo " killed=$pass survived=$survived build-fail=$buildfail invalid=$invalid  (of $total)"
echo "=============================================="
((survived)) && printf ' SURVIVED: %s\n' "${SURVIVED_LIST[@]}"
((buildfail)) && printf ' BUILD-FAIL: %s\n' "${BUILDFAIL_LIST[@]}"
((invalid)) && printf ' INVALID: %s\n' "${INVALID_LIST[@]}"

if ((survived || buildfail || invalid)); then
	exit 1
fi
exit 0

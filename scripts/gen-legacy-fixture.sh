#!/usr/bin/env bash
# Regenerate testdata/legacy-store from the binary that actually wrote that shape.
#
# The fixture under testdata/legacy-store/records/ is the on-disk output of
# certify as it existed at BASE_REF — before the `unsupported` field existed.
# It is captured, never authored. A hand-written "legacy" fixture agrees with
# whatever the author already believes, which is how a fixture written as
# {"status":"exempt","grade":"N/A"} let the grade-distribution defect through:
# the real pre-fix binary wrote {"status":"decertified","grade":"F"}, and only
# that shape reaches the code path where a stored grade contradicts a
# backfilled Unsupported flag.
#
# Run from anywhere. Rewrites testdata/legacy-store/records/ in place.
set -euo pipefail

BASE_REF="${BASE_REF:-6c9110de3}"
REPO_ROOT="$(git -C "$(dirname "${BASH_SOURCE[0]}")" rev-parse --show-toplevel)"
DEST="$REPO_ROOT/testdata/legacy-store/records"

WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

echo "==> building certify at $BASE_REF"
git -C "$REPO_ROOT" worktree add --detach -q "$WORK/base" "$BASE_REF"
(cd "$WORK/base" && go build -o "$WORK/certify-base" ./cmd/certify)

echo "==> building the mixed corpus (5 Go units, 4 Swift units)"
CORPUS="$WORK/corpus"
mkdir -p "$CORPUS/src"
cat >"$CORPUS/go.mod" <<'EOF'
module example.com/corpus

go 1.24
EOF
for i in 1 2 3 4 5; do
	cat >"$CORPUS/src/mod$i.go" <<EOF
package src

// Add$i returns the sum of a and b.
func Add$i(a, b int) int {
	if a > b {
		return a + b
	}
	return b + a
}
EOF
done
for i in 1 2 3 4; do
	cat >"$CORPUS/src/View$i.swift" <<EOF
import Foundation

struct View$i {
    let title: String
    func render() -> String {
        return "View$i: \\(title)"
    }
}
EOF
done

git -C "$CORPUS" init -q .
git -C "$CORPUS" add -A
git -C "$CORPUS" -c user.email=fixture@example.com -c user.name=fixture commit -qm init

echo "==> running the $BASE_REF binary against it"
(cd "$CORPUS" && "$WORK/certify-base" scan >/dev/null && "$WORK/certify-base" certify --skip-agent >/dev/null)

echo "==> capturing records into $DEST"
rm -rf "$DEST"
mkdir -p "$DEST"
cp "$CORPUS"/.certification/records/*.json "$DEST/"

git -C "$REPO_ROOT" worktree remove --force "$WORK/base"

echo "==> captured $(find "$DEST" -name '*.json' | wc -l | tr -d ' ') records"
echo "    verdict fields as written by $BASE_REF:"
for f in "$DEST"/*.json; do
	python3 - "$f" <<'PY'
import json, sys
d = json.load(open(sys.argv[1]))
print("      %-28s status=%-12s grade=%-3s score=%s unsupported=%r"
      % (d["unit_id"], d["status"], d["grade"], d["score"],
         d.get("unsupported", "<absent>")))
PY
done

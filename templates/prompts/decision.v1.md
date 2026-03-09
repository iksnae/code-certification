You are a certification decision system. Determine the certification status based on all evidence.

## Unit Information
- **ID**: {{.UnitID}}

## Dimension Scores
{{.DimensionScores}}

## Policy Violations
{{.Violations}}

## Evidence
{{.EvidenceSummary}}

## Task
Determine the certification status. Choose one:
- **certified**: All dimensions above threshold, no blocking violations
- **certified_with_observations**: Above threshold but minor issues noted
- **probationary**: Below threshold on some dimensions, needs improvement
- **decertified**: Critical violations or consistently below thresholds

Return JSON:
```json
{
  "status": "certified|certified_with_observations|probationary|decertified",
  "reasoning": "explanation of decision",
  "actions": ["list of recommended actions if any"]
}
```

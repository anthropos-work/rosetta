**Type:** tik (#3, under TOK-01)

# M46 · iter-04 — gen-batch preview / dry-run mode

Implemented deliverable #3: the offline batch-review surface.

## What landed (rext `stack-seeding/cmd/gen-batch/main.go`)
- New `--preview` flag — IMPLIES `--dry-run` (no LLM call, no key, no seeding). Renders, per expanded
  member: its routing context (`story=N batch=B role seniority bias` + a CACHED/to-generate status), the
  full rendered MOTHER prompt, the already-cached generated JSON (pretty-printed) or `<not yet generated>`,
  and a per-member estimated cost. Ends with a TOTAL estimated-cost line + a stdout summary.
- New `--preview-out <path>` — writes the full render to a file (the stdout summary still prints). Requires
  `--preview` (a guarded error otherwise).
- `renderPreview()` reuses the existing `estimateUsage` + `aiwrap.CostOf` (the model→price table) for the
  estimate. VALUES-BLIND: it never builds an LLM client, never reads/echoes a key — a pure offline review
  of what a real run WOULD send + spend.

## Tests (fixtures-first, no key/cost) — `cmd/gen-batch/main_test.go`
4 new tests (a panicking fixture proves the LLM is never called), green + the full stack-seeding suite green:
- `PreviewRendersPromptsAndCost` — renders per-member prompts (org/role/reserved-name) + `<not yet
  generated>` + the estimated-cost line; no LLM call.
- `PreviewShowsCachedJSONAndZeroCost` — a pre-cached batch shows the cached JSON + `$0.0000`.
- `PreviewOutWritesFile` — `--preview-out` writes the render to a file; stdout announces the write.
- `PreviewOutRequiresPreview` — `--preview-out` without `--preview` is a guarded error.

## Real offline smoke (no key)
`gen-batch --seed presets/gen-batch-20.seed.yaml --preview` renders all 20 members with their full prompts +
a total estimated cost of **$0.0062** (gpt-4o-mini, matching the M45 doc's ~$0.005/N=20). No LLM call.

## Re-measure
Gate primary metric (the M42 semantic sweep on a generated org) NOT measured this iter — it moves at the
real-run gate-proving tik (#5). Progress here is the author-review deliverable landed toward the gate.

## Close — 2026-06-26

**Outcome:** the gen-batch preview / dry-run mode (deliverable #3) landed + unit-proven (4 new tests + full
suite green + a real offline smoke). An author can now review a batch's prompts + cached JSON + estimated
cost before committing a real run, with no key/cost.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (org-scale gate proven on the real-run sweep, tik #5; this iter lands a deliverable the
gate depends on)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** (none beyond TOK-01)
**Side-deliverables:** none.
**Routes carried forward:** deliverable #4 (`--gen-batches` fence + throughput/429 verification) + #5 (the
real-run gate-proving sweep) remain on TOK-01's plan.
**Lessons:** Making `--preview` IMPLY `--dry-run` (rather than a third independent mode) keeps the no-key /
no-LLM / no-seed invariant automatic — the preview can never accidentally spend. Reusing the live
`estimateUsage` + `CostOf` for the estimate means the preview's cost number matches the real run's ceiling
pre-check exactly (the author sees the same dollars the `--max-cost` guard will).

---
iter: 07
milestone: M21
iteration_type: tik
status: closed-fixed
created: 2026-06-12
---

# M21 iter-07 — tik (sixth tik under TOK-01): wire structure capture + AUTO-PROVISION on replay

Under TOK-01 + M21-D7 (option A). Continue `STRUCT-M21-codeify`: make `stacksnap` ITSELF capture the schema structure
and APPLY it before the row replay. This automates stages 3→4 (the schema half); the serve rows (registration +
permissions) are iter-08 → gate met. Executing-at-bring-up is M22.

## Active strategy reference
TOK-01 + M21-D7 (option A) + the iter-06 capture core (`directus.CaptureStructure`, `manifest.Structure`).

## Re-survey (Phase 1 Step 0)
furthest-passing-stage = 6 (demonstrated); gate met-by-tooling pending the apply wiring. The iter-06 capture core is
committed (rosetta-extensions `2c42ed5`); this iter wires it through capture + replay.

## Cluster / target identified
The apply side has two halves: (1) **capture** must store the structure artifact in the snapshot; (2) **replay** must
APPLY it to a bootstrapped-gap stack before the row replay, redefining the exit-5 "cache miss" into "provision the
schema from the captured structure, then load". Today replay exits 5 on a bootstrapped-gap directus stack (M21-D3) —
this iter makes it auto-provision instead.

## Hypothesis / deliverables
1. **Capture wiring:** `capture.Surface.CapturesStructure` (directus opts in) + an optional `StructureCapturer`
   interface (type-asserted, so existing capture fakes are untouched). `capture.Run` calls it, stores the
   `_structure.sql` payload, sets `manifest.Structure`. `captureAdapter` implements it via `directus.CaptureStructure`.
2. **Apply wiring:** `replayCmd` — on cache MISS, locate the surface's structure-bearing cached snapshot, apply its
   structure SQL to the offset target (multi-statement Exec), re-probe the digest, re-resolve. Hit → replay rows
   (exit 0). Still miss → exit 5 (genuine divergence). Empty schema (never bootstrapped) still exit 4.
3. **Recipe truth:** `provision.go` ProvisionPlan step-2 (content-schema) updated — no longer "NOT YET AUTOMATED";
   `stacksnap replay` now auto-provisions the structure.
4. **Tests:** unit (capture stores structure; replay auto-provision flow against fakes) + a **live two-harness
   integration test** (source harness with the 26-table structure → `stacksnap capture` → fresh bootstrapped target →
   `stacksnap replay` auto-provisions → digest converges to `6cd35278…` → exit 0, real rows loaded).

## Expected lift
No furthest-stage ordinal move (stays 6-demonstrated) — this is a code-ification build iter graded on deliverables
landing + the integration test going green. It converts stages 3–4 from hand-applied to **stacksnap-applied** (the
gate's "stacksnap applies the captured structure" clause, half-met; serve rows in iter-08 complete it).

## Phase plan (multi-step planned shape — the apply-side wiring)
1. Capture wiring (capture.go + adapter + directus.Surface) + unit test.
2. Apply wiring (replayCmd auto-provision + the conn exec-script + redefined exit semantics) + unit test.
3. provision.go recipe-truth update + its tests.
4. `go build/vet/test ./...` green.
5. Live two-harness integration test (capture → bootstrap → replay auto-provision → exit 0 + rows).
6. Adversarial verification workflow on the diff + the integration evidence (ultracode).

## Escalation conditions
- If the multi-statement structure Exec can't run over pgx (protocol limitation), switch to a statement-splitter or a
  dedicated ExecScript; record + proceed (not a blocker — a known pgx simple-protocol path exists).
- If auto-provision risks applying to a non-gap (diverged) stack, gate it on cache-miss + a structure artifact +
  fail-closed on apply error (exit 5, no half-provision).

## Acceptable close-no-lift outcomes
- n/a — a build iter; closed-fixed when the wiring lands + unit + integration tests green.

## Test discipline
`go build ./... && go vet ./... && go test ./...` + the live integration test (throwaway harnesses, torn down at close).

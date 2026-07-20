# iter-03 — progress

**Type:** tik (under TOK-01, Track A) — resolves USER-BLOCKER-M235-01.

## Execution
- **scrub** (`scrub.go`): `NameTokens` + `AddNameReplacements` (token-split names) + word-boundary-aware
  matching + `SurvivingToken` leak probe. +5 unit tests (owner-first-name-only, word-boundary no-corruption,
  SurvivingToken, first-writer-wins, NameTokens). Existing scrub regression tests still green.
- **capture** (`cmd/content-capture/main.go`): source the session owner's identity from `public.users`
  (`sessions.owner_id` → firstname/lastname + email local-part) → `<<ACTOR_0>>`; token-split actors + owner;
  **G-post** capture-time fail-closed post-condition (refuse to write a fixture where a sourced name survives).
- **offline gate** (`content_stories_test.go`): `TestEmbeddedContent_NoStructuralPII` now also asserts the set
  carries `<<ACTOR_0>>` (the zero-placeholders regression tripwire — RED on buggy fixtures, GREEN after fix).
- **re-capture**: all 9 fixtures re-written through the fixed path (prod read-only, marco_read, counts-only).
- **docs**: session-clone-spec.md §3+§6, safety.md §3.8, KB-1 content-sessions.yaml header.
- rext commit `25a5459`, tag `playbill-m235-scrub-fix`.

## Verify (counts / exit-codes only — no content read into context)
- `<<ACTOR_0>>` placeholders: **9/9** files, **545** total (was **0**); per-file 44–107 (correlates with
  transcript length → no over-redaction explosion).
- 8 flagged first names: **0** matches across all 9 fixtures (was 8/9 leaking); 0 files contain any.
- Full unit gate GREEN; module-wide `go vet` + `go test ./...` clean.

## Close — 2026-07-19

**Outcome:** USER-BLOCKER-M235-01 resolved — scrub now sources the session owner's real identity + token-splits
names + fails closed on a surviving sourced name; 9 fixtures re-captured provably clean (0 leaked names, 545
placeholders present). Milestone live gate not moved (this is Track-A readiness/hardening).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (Track-A readiness prerequisite; the live (session×action) landing gate is Track-B → later tiks / M236)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** milestone-root decisions.md "RESOLUTION of USER-BLOCKER-M235-01".
**Side-deliverables:** word-boundary-aware scrub matching (a strict correctness improvement over plain-substring; benefits org + name scrubbing) — landed in the same scrub change, unit-proven.
**Routes carried forward:** the fixture-matrix closure (4 new captures: +1 asmt-voice-pass, +1 asmt-doc-pass, hiring not-passed, interview not-passed) → next tik under TOK-01 Track A; Playthrough + coverage descriptors; non-sim product sections; M230 clone re-anchor.
**Lessons:** the leaked identity lived one join away from where the scrub looked (owner `public.users`, not `jobsimulation.actors`). The durable gate is the *capture-time* fail-closed post-condition (knows the names in-process), NOT an offline test (can't know arbitrary names) — the offline gate can only prove the machinery fired (placeholder presence).

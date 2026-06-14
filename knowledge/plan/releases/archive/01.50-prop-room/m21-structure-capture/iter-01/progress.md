# M21 iter-01 — progress

**Type:** tok (bootstrap) — the milestone's unconditional iter-01, per build-mstone-iters Phase 0 rule 1.
**Protocol shape:** the protocol doc (`alignment_testing.md`) is a fidelity-measurement discipline, not a per-iter
build protocol; this bootstrap tok authors the concrete per-iter shape (staged-pipeline build toward the binary gate).

## Work done
1. **BYS context loaded** — milestone overview/progress/decisions, protocol doc, no prior iters (first invocation).
2. **Infrastructure probed (live):** Docker 28.5.1 running; `directus/directus:11.6.1` image cached locally; the
   directus snapshot cache present + complete (`.agentspace/snapshots/directus/6cd35278…/`, 9 tables, ~25 MB COPY
   payloads); `stack-snapshot` builds clean (`go build ./...` exit 0) on go1.25.3.
3. **Phase 0b KB-fidelity gate** — discharged YELLOW from fresh in-area research (recorded in `spec-notes.md`
   § Pre-flight audits — iter-01). The contract M21 builds against (`snapshot-spec.md` store-fork) is faithful to
   live code; the known stale areas are routed to M24 and are not M21's contract.
4. **Static baseline established** — the 6-stage pipeline + furthest-passing-stage = 2 of 6 (recorded in
   `spec-notes.md` § Baseline gate-distance). Confirmed the crux code points live: `provision.go:102-108` placeholder,
   exit codes 4/5 at `cmd/stacksnap/main.go:63,72`, `directus_files` absent from `directus.Surface()`/the cache.
5. **TOK-01 authored** — the staged-pipeline strategy (see milestone-root `decisions.md` → TOK-01, and `iter-01/overview.md`).

## Close — 2026-06-11

**Outcome:** Authored TOK-01 (staged-pipeline build toward the binary serve-anonymously gate; structure artifact
keyed by source digest, applied before row replay). Established the 6-stage metric + the static baseline (stage 2/6).
No gate progress (tok — strategy work, by definition).
**Type:** tok (bootstrap)
**Status:** closed-fixed (the bootstrap tok's planned deliverable — the initial strategy + baseline — landed in full)
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this is a BOOTSTRAP tok — does not exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (toks don't count toward the cap) — (6) protocol-stop: n — Outcome: continue (bootstrap tok continues into iter-02 tik within the same call)
**Decisions:** none intra-iter (the strategy of record is the milestone-root TOK-01 entry).
**Side-deliverables:** none.
**Routes carried forward:** the structure-source question (option a/b/c) → resolved in iter-02 by evidence; the
`directus_files` capture wiring → a stage-3 sub-task under TOK-01.
**Lessons:** the declared `iteration_protocol_ref` (alignment_testing.md) is a *measurement* discipline, not a
per-iter build protocol — its snapshot-fidelity dimension is the right faithfulness frame for the *served* content,
but M21's per-iter loop is a staged-pipeline build with a binary gate. Future structure/capture milestones with
observable (non-score) gates should expect the bootstrap tok to author the iter shape rather than inherit it.

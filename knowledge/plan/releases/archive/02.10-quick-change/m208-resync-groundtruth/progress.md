# M208 — Progress

Section checklist (closure = all boxes land).

- [x] `make pull` stack-dev + stack-demo platform to origin/main (skiller gone) + app to v1.334 + siblings; capture before/after refs
- [x] Remove vestigial `stack-dev/skiller/` + `stack-demo/skiller/` clones
- [x] Rebuild images + re-migrate against merged `public` schema — cold `make up` rc=0; clean-slate `make reset-db`+`make migrate` creates the full `public` taxonomy from scratch (no skiller schema on a clean DB), given the `extensions` prerequisite (Finding 1)
- [x] Confirm 4-subgraph compose / no skiller container / `SKILLER_RPC_ADDR=http://backend:8083` — all confirmed; router serves the absorbed taxonomy subgraph (existing-volume run); backend needs `INVITATION_HMAC_SECRET` (Finding 2, routed)
- [x] Pin the merge fact-sheet (moved tables, `org_id IS NULL` predicate, measured 42,790 count, RPC, 4-subgraph list, `extensions` prerequisite) — `corpus/services/backend.md` § + banner; `corpus/services/skiller.md` stub banner
- [x] Opportunistic M25-D9 — **surfaced on clean-slate** (not Fate-1); routed **Fate-3 to M211** (extensions-bootstrap + PG-readiness) + M209 Risk-2 cross-ref

## M208: Final Review

Close review (2026-07-08). The committed diff is **100% documentation** (fact-sheet + milestone records +
two sibling-overview routing edits + state.md; zero code/test/lint surface — confirmed by
`git diff --name-only e319d2f..HEAD`). HARDEN was correctly N/A (no code to deepen tests for; the release's
testable surface is M209 rext).

### Scope
- [x] All 6 section boxes checked; every `overview.md` In-scope item delivered; the opportunistic M25-D9 item
  legitimately re-fated (Fate-3 → M211 — did not fall out as a trivial Fate-1, which the overview permitted).

### Code Quality
- [x] No code touched → no lint/type/consistency surface. Nothing to fix.

### Documentation
- [x] Fact-sheet accurate + internally self-consistent; both anchor cross-refs resolve; 0 orphan TODO/FIXME;
  no new developer-facing unit → per-unit handbook contract N/A. KB-1/2/3 (pre-merge prose) correctly OUT of
  M208's charter → M210 (Fate 2), not a Phase-7 gap.

### Tests & Benchmarks
- [x] No test/benchmark surface in this diff (docs-only). rext test suites are M209's scope. Nothing to run.

### Decision Triage
- [x] The merged-shape facts are already blended into knowledge — the `backend.md` fact-sheet **IS** the
  knowledge blend (with the `extensions` clean-bring-up prerequisite captured in-line). Remaining decisions
  (KB-1/2/3, Finding 1/2, environment-park, live-de-risk, fact-sheet-placement) are milestone-specific
  archive or already-blended — no further blending owed.

**Consolidated fix-queue: 0 findings.** Deferral re-audit (Phase 1b) GREEN. Clean close.

## M208: Completeness Ledger (Phase 9, section variant)

- **Done (Fate 1) — all 5 `overview.md` In-scope items:**
  1. Re-sync both stacks — `platform` `5e1ae6b→0808b92`, `app` `a848cccb→c3c45e01` v1.334.1 (86-commit merge pull) + sibling set; before/after refs captured (spec-notes table).
  2. Remove vestigial `stack-dev/skiller/` + `stack-demo/skiller/` clones — done; no residual compose wiring.
  3. Rebuild + re-migrate against merged `public` — live de-risk GREEN (cold `make up` rc=0; clean-slate `reset-db`+`migrate` builds full public taxonomy, no skiller schema on a clean DB).
  4. Confirm 4-subgraph compose / no skiller container / `SKILLER_RPC_ADDR=http://backend:8083` — all confirmed.
  5. Pin the merge fact-sheet — `corpus/services/backend.md §` + banner, `corpus/services/skiller.md` stub.
- **Confirmed-covered (Fate 2):** KB-1 (backend.md prose) → M210 · KB-2 (skiller.md body) → M210 · KB-3 (5→4 subgraphs) → M210 · Finding 2 (`INVITATION_HMAC_SECRET` dev `.env` gap) → M211 / `/stack-secrets`.
- **Annotated (Fate 3):** Finding 1 (extensions-bootstrap + PG-readiness, M25-D9 class) → M211 (`overview.md` edited) + M209 Risk-2 cross-ref.
- **Dropped:** none.
- **Release-scope-breaking deferral (escape hatch):** **none.**

Zero escape-hatch entries → no user sign-off required → proceed to merge.

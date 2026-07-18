---
title: "Deferral Audit — M225 dress-the-set (close)"
date: 2026-07-17
scope: milestone
invoked-by: close-milestone
---

## Verdict
**YELLOW** — single, consciously-tracked carries with accepted destinations and standing user sign-off. No
repeat-deferral of *promised* milestone work without a decision; no chronic/silent scope erosion. Both carries
are routed to an explicit re-fate at the v2.4 **release** close (the "last-chance / extra-scrutiny" gate).

## Summary
- Total deferrals in scope: **2** (both INHERITED carries; **M225 introduced ZERO new deferrals**)
- Single deferrals: 2
- Repeat deferrals (of promised feature work): 0
- Chronic patterns flagged: 0

## Deferral Inventory

### M225's own ledger — CLEAN
- Every Scope.In item **landed Fate-1** (S1 bring-up guard + S2 3-seat coverage gate MET + S3 GREEN recruiter
  playthrough + S4 docs). `overview.md` Out: (the live cross-machine proof) is **Fate-2, already owned by M226**,
  not a deferral.
- S3's candidate playthrough is a **conscious pillar split** (recruiter playthrough = the gate; candidate covered
  on the presence side by S2's candidate coverage manifests), recorded in **D3** — a scope decision, **not** a
  deferral.
- KB-1 / D1 / D2 / D3 carry no "defer/postpone/later" — all are Fate-1 landings or investigations.

### DEF-CARRY-A — 8 pre-existing demo-stack test failures (inherited)
```yaml
- id: DEF-CARRY-A
  item: "8 pre-existing demo-stack test failures (6× test_cockpit.py [4 removed-academy-CTA + 2 v2.3.1 overlay-JS] + test_purge + test_reap)"
  origin_milestone: pre-v2.4 (v2.3.1/v2.3.2 cockpit hotfixes); first carried explicitly at M224 close
  first_deferred_on: 2026-07-16 (M224 D6)
  last_seen_in: M224 decisions.md:210 (D6) + state.md standing backlog + rext demo-stack tests (HEAD-identical)
  destination: "standing test-debt backlog → a future demo-stack test-debt harden pass"
  reason_recorded: "HEAD-identical; in files M224/M225 never touched; predates v2.4; outside the hiring domain — fixing = scope-bleed"
  partial_attempted: no
```

### DEF-CARRY-B — the M204 assign-WRITE declared TODO (inherited)
```yaml
- id: DEF-CARRY-B
  item: "assignment-monitoring.assign-and-track.UC1 — the assign-WRITE half (two-backend org-admin WRITE flow)"
  origin_milestone: M204 (v2.0 opening night)
  first_deferred_on: 2026-07-02 (M204, declared in-manifest TODO)
  last_seen_in: corpus/ops/demo/playthroughs.md ("15 live Playthroughs, 1 TODO") + rext playthroughs manifest (reports `unimplemented`)
  destination: "v2.4 release close (its declared-TODO fate) — a DECLARED, manifest-tracked build-reference gap, not silent erosion"
  reason_recorded: "declared build-reference gap tracked in the manifest; out of M204's declared 3 manager journeys; reports `unimplemented`"
  partial_attempted: no
```

## Repeat-Deferral Patterns
**None of the blocking kind.** Neither carry is a repeat-deferral of *promised, desired milestone work* pushed
forward without a decision:
- **DEF-CARRY-A** has been carried M224→M225 (2 milestones), which trips the aging trigger — but it is an
  **inherited-failure carry** in untouched, unrelated files (the `CHRONIC_DEFER of a wanted feature` pattern the
  gate blocks on does not apply). It has a conscious decision (M224 D6), a destination (standing test-debt backlog),
  and standing user sign-off (M224 close).
- **DEF-CARRY-B** is a **declared in-manifest TODO** (reports `unimplemented`) carried since v2.0 — explicit,
  tracked, and surfaced in every count. Its designated re-fate point is the **release** close.

## Fate-1 Investigation

### DEF-CARRY-A
- **Fate-1 (land now, complete) feasible:** no.
- **Why:** the failures live in the rext **demo-stack** section (`test_cockpit.py` / purge / reap), a domain M225
  never touched (M225 = coverage manifest + playthrough + set-dress-guard docs). Fixing them inside an M225 close
  is scope-bleed into an unrelated module; a proper fix is a dedicated demo-stack test-debt harden pass.
- **Fate:** KEEP-DEFERRED (carry) — driver-authorized this pass; re-fate explicitly at v2.4 release close.

### DEF-CARRY-B
- **Fate-1 (land now, complete) feasible:** no (and out of M225's scope entirely).
- **Why:** the assign-WRITE half is a two-backend org-admin WRITE flow — a declared build-reference gap owned by
  the Playthroughs manifest, not part of M225's recruiter-vantage deliverable. M225 *added* a live playthrough
  (hiring), moving the live count 14→15 while leaving this one declared TODO untouched.
- **Fate:** KEEP-DEFERRED (declared TODO) — its explicit fate belongs to the v2.4 release close.

## Recommendations
- **DEF-CARRY-A → KEEP-DEFERRED-WITH-SIGNOFF (carry).** Fresh decision dated **2026-07-17**: still a valid carry —
  nothing M225 learned changes the calculus (M225 touched no demo-stack test file); destination unchanged (standing
  test-debt backlog → a future demo-stack test-debt harden pass). **Re-fate explicitly at v2.4 release close.**
- **DEF-CARRY-B → KEEP-DEFERRED (declared TODO).** Fresh decision dated **2026-07-17**: unchanged; its declared-TODO
  fate is a v2.4 **release-close** decision (close-release Phase 1b — the last-chance gate), per the M204 manifest
  declaration.

## Applied Changes
- This report authored (fresh dated re-fate for both carries).
- M225 `decisions.md` gains a short **D-AUDIT** note referencing this report + the two carries' fates.
- No plan mutations, no code fixes, no roadmap edits — both carries already have durable homes (M224 D6 + state.md
  standing backlog for A; the Playthroughs manifest + playthroughs.md for B). The v2.4 close-release Phase 1b audit
  inherits both with `release` scope and the "extra scrutiny" mandate.

## Blocking Items (require user decision)
**None.** No repeat-deferral of promised milestone work; both inherited carries have conscious decisions,
destinations, and standing sign-off, re-fated fresh this pass and routed to the v2.4 release close. Verdict YELLOW →
`SEVERITY=warning`; the milestone close proceeds.

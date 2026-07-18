---
milestone: M228
slug: second-night
version: v2.4 "casting call"
milestone_shape: iterative
status: archived
created: 2026-07-17
last_updated: 2026-07-18
depends_on: M222, M223, M224, M225, M226, M227
exit_gate: "On billion.taildc510.ts.net, a default /demo-up N (no flags) yields, reproducibly on a cold reset-to-seed, with the M227 corrections rendering: (1) the hiring org present, is_hiring=true, exactly 5 managers + 45 candidates; (2) the recruiter hero lands on the comparison surface and sees ≥ the retuned floor (≥6, a small margin below the seeded ~8/position) comparable non-junk rows per EACH of the 5 positions, each candidate on exactly 1 sim; (3) the 2 candidate heroes render usable assessed profiles with external emails + gender-matched avatars visible; (4) the org reads as hiring, hiring-only content (no training/assessment sims in its surfaces); (5) p95 click→ACCESS < 5 s for the recruiter vantage; (6) coexists with the 3 workforce orgs; (7) 0 platform-repo edits."
iteration_protocol_ref: corpus/ops/verification.md (+ corpus/ops/demo/coverage-protocol.md + corpus/ops/demo/latency-budget.md — the M226 remote-origin cold-reset-to-seed acceptance run, re-run on the M227-corrected data)
delivers: the corrected live-proof record; any live finding folds into the relevant corpus doc
---

# M228 — Second night

## Goal
The **corrected** (M227) hiring demo re-proven **live on billion**, over the tailnet, **default `/demo-up`, no flags**
— the M226 7-condition gate (retuned to the believable per-position count) holds on a cold reset-to-seed, with the
four M227 believability corrections **rendering live**.

## Exit gate (measurable — the 7-condition live billion proof, RETUNED for M227)
On `billion.taildc510.ts.net`, a **default** `/demo-up N` (no flags) yields, **reproducibly on a cold reset-to-seed**,
with the M227 corrections rendering:
1. the hiring org present, **`is_hiring=true`**, **exactly 5 managers + 45 candidates**;
2. the recruiter hero lands on the comparison surface and sees **≥ the retuned floor (`≥6`, a small margin below the
   seeded ~8/position) comparable non-junk rows per EACH of the 5 positions**, **each candidate on exactly 1 sim**;
3. the 2 candidate heroes render **usable assessed profiles** — **external emails** (gmail/outlook/…) + **gender-matched
   avatars** visible;
4. the org **reads as hiring**, **hiring-only content** (no training/assessment sims in its surfaces);
5. **p95 click→ACCESS < 5 s** for the recruiter vantage;
6. **coexists with the 3 workforce orgs** on the cockpit;
7. **0 platform-repo edits.**

## Why iterative
The M226 "prove-on-billion" pattern (itself the M215 "prove-on-odyssey" lineage): the last breakages surface only on a
live cross-machine cold run over the tailnet. The corrected data (1-sim/candidate, external emails, hiring-only content,
matched avatars) is re-proven live — the path is discovered, not enumerated.

## The M227 corrections re-proven (the delta vs M226)
The billion infra + the 4 hiring demo-patches + the tailscale-serve hiring front are ALL proven (M226). **Only the DATA
changed** (M227, rext tag `casting-call-m227-sections`):
- **(a) 1 sim/candidate** → each of the 5 positions has ~8 candidates (was ~43); gate floor retuned `≥40 → ≥6`.
- **(b) hiring-only content** — the hiring org surfaces ONLY `SIMULATION_TYPE_HIRING` sims.
- **(c) external candidate emails** — candidates show gmail/outlook/proton/icloud, NOT `@meridian-talent.com`.
- **(d) gender-matched avatars** — every user's face matches their name's inferred gender.

## Iteration protocol
`corpus/ops/verification.md` + `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/latency-budget.md` — the M226
remote-origin cold-reset-to-seed acceptance run, re-run on the M227-corrected data. Same billion-safety rules (one
driver, no detached scripts on-host, assert from a tailnet peer, never kill a mid-build).

## Scope

### In
- Consume rext tag `casting-call-m227-sections` on billion; a default `/demo-up 1` (no flags) cold reset-to-seed; drive
  the retuned 7-condition gate + the 4 M227-correction render checks to green, reproducibly.
- **Latency is GATED here** (condition 5) for the recruiter vantage.

### Out
- Nothing net-new; this is the corrected-demo acceptance closer. Any un-patchable surface **escalates**.

## Depends on
**M227** (the corrected seed/content, rext tag `casting-call-m227-sections`, 63c3e8d). The `billion` demo from M226 was
LEFT UP (at `casting-call-m226-c2-race-fix`) as the prior live-proof artifact.

## KB dependencies
- `corpus/ops/demo/tailscale-serve.md` · `corpus/ops/demo/latency-budget.md` · `corpus/ops/verification.md` ·
  `corpus/ops/safety.md` · `corpus/services/hiring.md` (the M227-corrected read-model).

## Delivers → knowledge/corpus
The corrected live-proof record; any latency/render finding folds into the relevant corpus doc.

## Demo-patch?
Whatever M224/M226 pinned, **re-proven live at final code** (the M221 `REPROVE-…-at-final-code` discipline). M227 added
NO new demo-patch (pure seed/content tooling). **A platform-repo edit is never in bounds.**

## Risks carried
- **R1 (re-surface)** — a render gate that passed deterministically at M227 re-appears on the cold remote run (the M215/M221
  lesson: last breakages are cross-machine). The M227 render checks (external emails / avatars / hiring-only) were NEVER
  live-rendered (M227's local re-prove was env-blocked) — this is the sharpest new risk.
- **R4 (re-surface)** — the 45×5 whole-org-hydration latency; with 1-sim/candidate the compare table is now ~8/position,
  smaller, so the R4 warm-up transient should be milder, but re-measure (never assume).
- **Billion memory floor** (7.3 GiB + 15 GiB swap) under the 2-app hiring demo — watched from tik-01 (M226 held).

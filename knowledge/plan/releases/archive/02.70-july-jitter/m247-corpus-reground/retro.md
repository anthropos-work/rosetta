# M247 — Retro (corpus re-ground)

## Summary
The corpus-reground fan-out lane of v2.7 "july jitter". Re-grounded the corpus to the CONSOLIDATED platform
M246 proved live: `skillpath.md` → a merged-into-app **redirect** (mirrors `skiller.md`) + moved to the README
**archived/merged** table; the **4→3 subgraph** reclassification across ~30 echo files (each mention re-read in
context, not blind-sed); **4 net-new app-domain fact sheets** (`coursebuilder` / `ai-labs` / `askengine` /
`academy-backend`, `TEMPLATE.md`-shaped); `ai-readiness.md` refreshed for the aireadiness-package refactor
(platform-facts only); roadrunner resolved to the **negative** (stays ORPHANED, not archived). Section close,
all 8 sections Fate 1. **DOC-ONLY: zero rext, zero platform-repo edits, no code-of-record tag.**

## Incidents This Cycle
None. No source stack (documentation repository), so 0 flakes / 0 test regressions by construction. Fidelity
spot-check GREEN (0 corrections vs `stack-demo/app` @ `v1.351.1`); 0 broken links across 30 touched docs. No
bug surfaced during build, harden, or close.

## What Went Well
- **Fidelity-over-coverage as the doc analog.** With no test/coverage stack, the meaningful robustness signal
  was verifying every authored fact against current `stack-demo/app` source (ent schemas, migrations, model
  ids, env keys, scoring bands) — GREEN with 0 corrections. The load-bearing `ai-labs` "shared-purse UNBUILT"
  caveat was confirmed by the *absence* of a `checkout.session.completed` handler, not assumed.
- **Context-aware reclassification, not sed.** The ~37 skillpath mentions were re-read in place; the sweep
  deliberately LEFT the non-service uses (studio-desk's own `/api/skillpath` builder route, CMS skill-path
  *content*, seeder skill-path *data*, DB-dump role-name quirks) — grep-verified 0 residual "4 subgraphs"
  (bar the intentional historical note) and 0 stale live-skillpath-as-service claims.

## What Didn't (go as smoothly)
- **A deferral routing went stale mid-flight.** M247 D0 routed the inert rext-hygiene items (dormant skillpath
  INJECTED key + `test_injection.py`/`exposure_claim_guard.py` fixtures + audit prose) to "M251/rext-hygiene",
  but **M251 closed first — without owning them** (they were never in its `In:` scope). Caught at the Phase-1b
  deferral audit (the destination-milestone-closed aging trigger). Re-fated to a **documented-inert standing
  note** (0 functional impact; all four files are rext, 0 tracked in the rosetta repo). Lesson: a Fate-3 to a
  sibling that *doesn't formally own the item* is fragile when closes serialize — an inert item with no true
  owner is better recorded as an accepted standing note than routed to a milestone that won't pick it up.
- **An inherited hand-off arrived without a landing.** M251 punted an *optional* `verification.md` index
  anchor to M247 for lane-collision avoidance; M247 declined it (out of the consolidation charter, a
  test-health concern) → re-fated Fate-2 to the release-close pass. Non-blind, optional — no gap, but a
  reminder that a cross-lane punt needs the receiving milestone to actually own it in scope, or it strands.
- **Serialized-close bookkeeping contention.** `state.md` (frontmatter) and the roadmap v2.7 count row are
  touched by every close; M251 already advanced them on release. To keep the Phase-11 merge clean *by
  construction* (no conflict to force-resolve), M247's branch-side Phase-10 edits were kept off those shared
  lines (the build-time `state.md` tweak was reverted to base so it auto-merges to M251's version), and the
  `state.md` M247-close update + the "2 of 9 → 3 of 9" count bump landed as a **post-merge commit on release**.

## Carried Forward
- **ai-readiness demo-seeder fidelity** (31-skill demo set + track-keyed named sims + evaluated-skills
  set-dress) + the **D-07 demopatch re-pin** + the 4 demo-section compute line-anchors → **M250** (Fate-2;
  M250 also `Delivers → ai-readiness.md`).
- **ops/demo spec-doc reconcile** (content-stories-spec/routes, demopatch-spec, cockpit-spec, latency-budget,
  secrets-spec, studio parts of frontend-tier/studio-desk) → the **code milestones M248/M249/M250/M252/M253**
  + the **release-close consistency pass** (Fate-2).
- **Optional `verification.md` demo-stack-suite + run-unit-roster index anchor** (inherited from M251) →
  **release-close consistency pass** (Fate-2; optional + non-blind).
- **Inert rext-hygiene items** (dormant skillpath key + injection test fixtures + audit prose) → **documented
  standing note**, swept opportunistically by whichever rext milestone next edits those files (M249/M252) —
  no sibling `overview.md` edit made.

## Metrics Delta
(from `metrics.json`)
- Files: **36 changed** (30 corpus `.md` + `.gitignore` + 4 milestone plan artifacts); **4 new fact sheets**;
  CLAUDE.md service-doc count **23 → 27**; +1229/-427 lines.
- Fidelity: **GREEN, 0 corrections**. Broken links: **0** / 30 docs. Residual "4 subgraphs": **0** (bar 1
  intentional historical note). Stale live-skillpath claims: **0**.
- Deferral audit: **YELLOW**, 0 blockers (1 aged-out re-fated + 3 Fate-2). Platform-repo edits: **0**.
  Code-of-record: **N/A (doc-only)**.

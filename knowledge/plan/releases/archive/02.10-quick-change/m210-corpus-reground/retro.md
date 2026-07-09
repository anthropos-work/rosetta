# M210 Retro — Corpus + skills re-ground

**Closed:** 2026-07-08 · `closed-complete` · `section` · complexity medium.
**Outcome:** made the rosetta corpus **internally consistent** with the merged platform + M209's landed `public.*`
rext code. Adopted the colleague's correct architecture/subgraph/service half, flipped the 6 rext-facing tooling-doc
bodies (+ directus-local) `skiller.*→public.*`, reconciled db-access ↔ tooling, swept the 4 skill files + `CLAUDE.md`
to the merged 4-subgraph/no-skiller compose. **RESOLVES the KB-1/2/3 body-flip deferrals** M208+M209 routed here.
Zero platform-repo edits.

## Summary
The committed diff is **100% documentation** — 50 `.md` files (45 corpus/skills doc bodies + 5 milestone/plan
records), **0** `.go`/`.ts`/`.sh`. The core, load-bearing outcome (grep-verified corpus-wide at close): **0 stale
`skiller.<table>` tooling-query references**, 0 leftover interim disclosure notes, **4 subgraphs** consistent
everywhere, 0 broken relative `.md` links, db-access ↔ tooling reconciled on `public.*`. The corpus now aligns with
M209's landed code. HARDEN was correctly N/A (no code/test surface to deepen).

## Metrics delta (from `metrics.json`)
- **Tests:** N/A — docs-only milestone, 0 code/test surface. Inherited v2.0 baseline unchanged (rext Go 1763 funcs,
  10 live Playthroughs). **Flake:** 0 (no tests).
- **Core outcome:** 0 stale `skiller.<table>` refs · 0 interim notes · 4 subgraphs · 0 broken links · db-access ↔
  tooling reconciled (`public.skills` 42,763). Verdict GREEN.
- **Close findings:** 1 (nice-to-have, no-change-needed — the app==backend subgraph dual-naming is the corpus's
  established convention, not an M210 defect). **0 must-fix / 0 should-fix / 0 doc gaps / 0 test gaps / 0 blends.**
- **Deferral audit:** GREEN (11 in scope → 7 resolved-at-destination by M210, 4 still-open confirm-only; 0 repeat,
  0 chronic, 0 aged-out, 0 escape-hatch; 0 M210-originated).

## Incidents this cycle
- **No P0/P1/P2. No regressions. No flakes.** No fixes were owed at close (the review found 0 must-fix); the sole
  finding was a no-change-needed observation.

## What went well
- **Don't-fabricate discipline held.** The design note asserted a "43/44 → 44/44" member-count edit in
  `profile-completeness-spec.md`; an exhaustive search (corpus + `.claude/` + rext authoring copy + git history)
  proved **no such literal ever existed** and the file's schema refs were already `public.*`. The milestone made the
  one genuine merge-sweep fix (a node-id prose line) and **did not invent a phantom count** — the correct outcome
  under the three-fate rule + the audit don't-fabricate principle.
- **Selective per-section adoption beat cherry-pick.** The colleague's monolithic docs commit spanned all 6 sections
  and would have conflicted with M208's already-landed `backend.md`/`skiller.md` fact-sheet; per-section adoption
  kept the section discipline and preserved M208's authoritative fact-sheet (no duplicate merge section).
- **Superseded a now-stale interim note instead of adopting it.** The colleague's `stack-snapshot/SKILL.md` hunk
  warned the taxonomy surface "still targets skiller … exit 4" — written pre-M209. Since M209 already re-pointed the
  surface, M210 wrote an accurate M209-done note rather than importing a stale caveat.
- **Clean lockstep with M209.** The tooling-doc bodies now state `public.*` in step with M209's landed code; the
  design's "keep M210 adjacent-and-sequential-to-M209" call paid off — 0 self-contradicting-corpus risk.

## What didn't (go as smoothly)
- **The design note's "43/44" count was inaccurate** — it cost an exhaustive verification pass to confirm the
  literal never existed (rather than a quick edit). Correctly resolved by evidence, not fabrication, but a reminder
  that end-state design notes can carry phantom specifics.
- **Residual app==backend naming duality** persists in the corpus (some subgraph rows say `backend`, others `app`).
  It's the corpus's long-standing convention (repo=`app`, container/RPC=`backend`) and out of M210's charter, so it
  was left as-is — but it's a latent cosmetic-consistency item a future corpus-wide style pass could harmonize.

## Carried forward
- **DEF-M208-01 / M25-D9** (extensions-schema bootstrap + PG-readiness) → **M211** (Fate-3; `overview.md` pinned).
- **DEF-M208-02** (`INVITATION_HMAC_SECRET` dev `.env` gap) → **M211 / `/stack-secrets`** (Fate-2).
- **DEF-M209-04** (recapture the `public.*` taxonomy — a data op needing a COPY-byte source) → **M211** (Fate-3;
  exit-gate item).
- **TEST-1** (rext `stack-seeding/README` test-count drift, pre-existing since M41) → **close-release rext roll /
  M211 rext re-tag** (Fate-2; rext frozen at `2f06e78` this close).
- **Push-gated KEEP** (origin push of `main` + tags + rext consumption re-pin) — the user's manual gate; unchanged.

All four still-open items are single, freshly-fated, confirm-only routings to a downstream milestone (M211) or
close-release. None escape-hatch; none repeat/aged. **M211 is the FINAL v2.1 milestone** — after it,
`/developer-kit:close-release` rolls the `v2.1` rext tag and merges → `main`.

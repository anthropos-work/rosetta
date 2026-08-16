**Type:** tik (no TOK chain — Before-You-Start case (b); see `overview.md` § Type selection)
**Protocol:** [`corpus/ops/verification.md`](../../../../../../corpus/ops/verification.md)

# M265 iter-01 — progress

## What the iter found

**The headline is not any single clause.** It is that v2.9 shipped five milestones' worth of taxonomy
realignment, every one green, and the demo's simulation library still rendered **zero cards**. Nothing
in the release measured that, because nothing measured a *rendered surface*.

### The defect the gate existed to catch

Taxonomy and Directus content are two **separate** snapshot surfaces. Replaying the taxonomy swaps
`public.skills` wholesale (43,584 → 3,562); the content replays unchanged, still pinning skills by
node-id inside JSON. Measured: **302 distinct ids referenced, 187 retired.** The resolver's
`skills[].name` is non-null, so **one** dead id nulls the **entire** list:

```
ERROR graphql resolver error error="input:38:7: publicJobSimulations[5].skills[1].name
                                    ent: skill not found"
```

Observed: an empty library, sim detail pages returning empty HTTP 200 — with `/api/health` 200, 10/10
containers up and `public.skills = 3562` all green. **A row count cannot catch a hollow row, and a
liveness probe cannot catch an empty list.**

Repaired by `rosetta-extensions/stack-snapshot/realign`, which runs on every replay and rewrites dead
ids through `public.skill_redirects`. Final live figures: **26 json columns scanned, 515 dangling refs
→ 0, 515 repaired**; `skill not found` in the backend log **258 → 0**.

### Four self-inflicted defects, each the shape of the bug being fixed

1. **A hand-maintained column list, wrong within the hour.** It named four columns, repaired all four,
   verified clean, exited 0 — and the next page load still failed, because ids also live in
   `sequences.validation_evaluation_criteria` and `skill_paths.chapter_list`, nested a level down.
   Replaced with catalog discovery + exact-token substitution, so depth is irrelevant.
2. **`EXCEPTION WHEN others THEN dead := NULL`** turned a five-args-to-four-placeholders `format()`
   bug into "this column holds nothing dead". The remap reported success and changed nothing.
3. **An anti-vacuity rule with too broad a trigger** recorded a *successful* 55,116-row replay as a
   FAILED surface on cold stacks, where the content schema legitimately does not exist yet.
4. **A scan pattern narrower than its subject.** `[A-Z]{4,8}` matched 3,543 of the canon's 3,562
   skills; the 19 it missed carry a digit in the stem (`K-DIGB2B-2416`). It did not under-report — it
   reported **clean** while the app still failed. A verifier narrower than what it verifies gives a
   false answer, which is strictly worse than not looking.

### Seed + harness defects surfaced by running the suite

- `playthroughs/seed/pt-world.seed.yaml` pinned **three retired role NAMES**. M262 fixed this class in
  `stack-seeding/presets/` — the directory the bug was found in — and pt-world lives in another
  section, so it was never scanned. Fenced by `seed_role_guard.py`, repo-wide.
- `seed_role_guard` was then **not registered in the guard family** — written, committed, and wired to
  run exactly never.
- `seed-facts.ts` still mirrored the same three retired names (the **third** place they lived).
- The AI-readiness "completed" hero mapped her skills **before the cycle opened** (`now − 7…67 days`
  against a cycle opening at `now − 45`), so she rendered NOT STARTED with three completed steps in
  the DB. Clamped, with the `45` now a shared constant.
- `pt-onboarding-org-prepared` asserted a skill (`Data Analysis`) the consolidation **retired
  outright**; re-pinned to `Descriptive Analytics`, derived from the role's actual canon profile —
  deliberately *not* the redirect target, because a redirect answers "what replaced this SKILL", not
  "what does this ROLE offer".

### Three dev-path defects that hid each other

`/dev-up` on the new canon had never been run. An additional `dev-N` could not be migrated at all: the
positional `N` was **silently ignored** (so it targeted the main dev stack), the atlas DSN dialled
`::1`, and a first-cut fix assumed dev-N has its own clone set — it does not, unlike demo. Each
failure presented three layers from its cause: no migration → backend `Exited(1)` on an enforcer
panic → no taxonomy → "the catalog is empty".

## Retraction made inside this iter

`pt-assignment-assign` was characterised as **not taxonomy-caused** and routed forward, on the strength
of two true measurements (seeded assignments carry `plan_id`; the nav has moved to the v2 surface)
composed into a story about a superseded surface. **Retracted.** The real cause was one hop further
out — the assign picker reads `publicSkillPaths`, still nulled by the single node-id the scan's own
pattern could not see. Widen the pattern, repair the 515th ref, and it passes.

**Two true facts and a plausible join are not a diagnosis.** The check that would have caught it is the
one that did: `skill not found` was not zero when the conclusion was written.

## Close — 2026-08-16

**Outcome:** All five gate clauses measured MET on cold bring-ups. The release's central defect — a
demo whose simulation library rendered zero cards while every probe passed — was found, repaired and
re-proven live.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-1

### Gate clause measurements

| # | Clause | Measured |
|---|--------|----------|
| 1 | cold `/demo-up` green on the new canon | **MET** — 0 failed surfaces, no `set-dress INCOMPLETE`; `✓ taxonomy replayed: public.skills = 3562` · `✓ demo-patches: all applied` · `✓ verify live: all probes passed` · `✓ container liveness: 10/10`. Realignment correct in BOTH orderings in one run: `SKIPPED — content schema not provisioned yet` (taxonomy surface) then `26 json column(s) scanned; 515 dangling ref(s) before, 0 after; 515 repaired` (directus surface) |
| 2 | `/dev-up` green on the new canon | **MET** — `dev-3` carries the canon: **3,562** skills / **706** roles / **12,835** redirects, 141 tables migrated, backend up, `taxonomy rows=3562 ok` |
| 3 | full Playthrough suite incl. the taxonomy Playthrough | **MET** — **222 passed / 0 failed**, cold reset-to-seed at rext `v2.9.17-rext` |
| 4 | seed closure + per-hero richness floor | **MET** — `all 247 seeded verified-skill node-id(s) resolve; 590 of 591 membership(s) populated` |
| 5 | `/taxonomy` navigable live | **MET** — walked category → skill (`/taxonomy/skill/ai-adoption-change-management`); `Taxonomy` present in the primary nav |

**Decisions:** D-M265-1 … D-M265-6 (milestone-root `decisions.md`)

**Side-deliverables:** none — every fix listed above was in service of a gate clause, so none are
side-discovery. Recorded here explicitly because an empty field and an unconsidered field look the
same.

**Routes carried forward:**
- The **claim-census ratchet** is broken across six corpus files (`dependency_map` 29→36,
  `external_services` 102→120, `backend` 30→32, `messenger` 13→16, `storage` 18→24, and the net-new
  `taxonomy-canon.md` at 37). Five predate M265. → **Fate 3 → `/developer-kit:close-release`**:
  re-baselining a ratchet is a target change, not a stale-reference repair, and five of the six grew
  in earlier milestones, so it is release-level by construction.
- Four pre-existing guard-family REDs (`clone_drift`, `decommissioned_instruction`, `demo_knob`,
  `unreadable_repo_claim`) flag files M265 did not touch. → **Fate 3 → `/developer-kit:close-release`**.
- Three archived-milestone scratchpads (`work-m257`, `work-m257x`, `work-m258`) are sweep candidates
  under the wrapper's archived sweep; `work-m257x` alone holds hundreds of evidence artifacts, so the
  destructive sweep was **not** performed. → **Fate 3 → `/developer-kit:close-release`** (its scratchpad
  sweep is the right owner, with the user's eyes on it).

**Lessons:**
- **A verifier narrower than its subject reports clean.** That is not under-reporting; it converts an
  open question into a false answer. Measure the verifier against the population it verifies —
  `[A-Z]{4,8}` vs 3,562 canon ids would have shown 3,543 in one query.
- **A hand-maintained list of the places a value can hide is wrong the moment the data changes shape.**
  Discovery beats enumeration wherever the subject is data-shaped.
- **A false alarm on the most common path is worse than no alarm** — it teaches the reader to discount
  the real one. Anti-vacuity rules need triggers as narrow as their intent.
- **A probe's denominator must share the scope of what it qualifies**, or it is two questions wearing
  one answer.
- **Two true facts and a plausible join are not a diagnosis** (see the retraction above).

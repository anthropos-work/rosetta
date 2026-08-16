# M265 — Progress

## Running ledger

### iter-01 — prove it live, and find the defect six green gates could not see

**The gate.** On a cold bring-up, ALL of: (1) `/demo-up` green end-to-end on the new canon;
(2) `/dev-up` green on the new canon; (3) the full Playthrough suite passing, INCLUDING the net-new
taxonomy Playthrough; (4) seed closure green WITH the per-hero richness floor satisfied;
(5) `/taxonomy` navigable live.

#### What this iter found

The headline is not any single clause. It is that **v2.9 shipped five milestones' worth of
taxonomy realignment, every one of them green, and the demo's simulation library still rendered
zero cards.** Nothing in the release measured that, because nothing in the release measured a
*rendered surface*. M265 exists to be the thing that does, and it earned its place on the first
run.

**The defect (`realign`).** Taxonomy and Directus content are two SEPARATE snapshot surfaces.
Replaying the taxonomy swaps `public.skills` wholesale (43,584 → 3,562); the content replays
unchanged, still pinning skills by node-id inside JSON. Measured: **302 distinct ids referenced,
187 retired.** The resolver's `skills[].name` is non-null, so ONE dead id nulls the ENTIRE list:

```
ERROR graphql resolver error error="input:38:7: publicJobSimulations[5].skills[1].name
                                    ent: skill not found"
```

Observed: an empty library, empty sim detail pages returning HTTP 200 — with `/api/health` 200,
10/10 containers up, and `public.skills = 3562` all green. **A row count cannot catch a hollow
row, and a liveness probe cannot catch an empty list.**

Fixed by `stack-snapshot/realign`, which runs on every replay and rewrites dead ids through
`public.skill_redirects`. Live: **26 json columns scanned, 257 dangling before, 0 after, 257
repaired**; `skill not found` in the backend log went **258 → 0**; both simulation Playthroughs
went red → green.

**Three self-inflicted defects while building it, each the shape of the bug it was fixing:**

1. **The hand-maintained column list was wrong on arrival.** It named four columns, repaired all
   four, verified clean, exited 0 — and the next page load still failed, because ids also live in
   `sequences.validation_evaluation_criteria` and `skill_paths.chapter_list`, nested a level down.
   Replaced with catalog discovery + exact-token substitution, so depth is irrelevant.
2. **`EXCEPTION WHEN others THEN dead := NULL`** turned a five-args-to-four-placeholders `format()`
   bug into "this column holds nothing dead". The remap reported success and changed nothing; only
   the independent post-scan caught it. A swallowed error that yields an empty work-list is
   indistinguishable from clean input.
3. **Driving the loop from the redirect tables did not finish** — 24,017 redirects × 26 columns is
   ~600k UPDATEs for the 257 that apply. Killed after two minutes. Now driven by the dead ids each
   column actually contains: work proportional to the damage, not to the taxonomy's history.

**Two more seed defects, same family as M262's, found by running the suite:**

- `playthroughs/seed/pt-world.seed.yaml` pinned **three retired role NAMES**. M262 fixed this class
  in `stack-seeding/presets/` — the directory the bug was FOUND in — and pt-world lives in another
  section, so it was never scanned. The suite's reset-to-seed leg failed outright on it. Fenced by
  `seed_role_guard.py`, which walks **every** seed yaml in the repo.
- `seed_role_guard` itself was then **not registered in the guard family** — written, committed, and
  wired to run exactly never. A fence nobody runs is not a fence.

**Not a product regression, worth recording:** `pt-assignments-nav-v2` (the user-reported "assign
content opens the legacy page") was red for two stale-locator reasons — the entry was RENAMED to
"Assign Program", and the nav-group expansion was guarded by `count()`, which does not auto-wait, so
on a slow render the expand was skipped IN SILENCE and the failure surfaced later as "entry not
found". The assertion under test is unchanged and passes: the nav lands on the v2 list.

#### Clause status

| # | Clause | State |
|---|--------|-------|
| 1 | `/demo-up` green cold on the new canon | ✅ green (`public.skills = 3562`, all demo-patches applied, 10/10 containers, all probes) |
| 2 | `/dev-up` green on the new canon | ⬜ not yet measured |
| 3 | full Playthrough suite, incl. the taxonomy Playthrough | 🟡 in progress — the content defect above is fixed; residual reds under triage |
| 4 | seed closure + per-hero richness floor | ✅ `all 247 seeded verified-skill node-id(s) resolve; 590 of 591 membership(s) populated` |
| 5 | `/taxonomy` navigable live | ✅ walked category → skill (`/taxonomy/skill/ai-adoption-change-management`); `Taxonomy` is in the primary nav |

#### Carried out of this iter, not silently

- **The claim-census ratchet is broken across six corpus files** (`dependency_map` 29→36,
  `external_services` 102→120, `backend` 30→32, `messenger` 13→16, `storage` 18→24, and
  `taxonomy-canon.md` new at 37). Five predate M265 — they grew during v2.9's corpus work and
  nobody re-baselined. Re-baselining a ratchet is a target change, not a stale-reference repair, so
  it is **routed to `/developer-kit:close-release`** rather than done here.
- Four other guard-family REDs (`clone_drift`, `decommissioned_instruction`, `demo_knob`,
  `unreadable_repo_claim`) are pre-existing and flag files M265 did not touch.

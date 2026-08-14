# M259 — Decisions

## `D-M259-1` — VERDICT **GO**

The consolidation is legible, the bundle is checked in and measurable without touching production, the
redirect map is clean, and **every risk to this project is in our own tooling** rather than the platform's.
Downstream milestones proceed as designed, with two resizings (`D-M259-2`, `D-M259-3`).

## `D-M259-2` — M262 cannot be a remap; it needs a resolve-or-drop path

**~67 % of retired skills have no successor** (26,518 of 39,353). "Remap seeded refs through the redirect
map" covers a third of the exposure. M262's `In:` list stands but its shape changes: prefer the redirect,
fall back to re-resolving by name against the canon, and drop only when neither lands — never fabricate.

**Not a Fate-3 edit:** M262 already lists "remap seeded refs THROUGH the redirect map" and the per-hero
richness floor; this refines HOW, and the milestone has not started. Recorded here, to be read at its start.

## `D-M259-3` — five tables join the capture surface (M260/M261)

`skill_redirect`, `job_role_redirect`, `category_translation`, `specialization_translation`,
`taxonomy_canon_state`. This is **net-new scope** neither milestone named at design time, because the tables
were not known then. It is small and squarely inside both milestones' stated goals ("re-ground every
taxonomy-size assumption", "a stack replays the new canon"), so it lands there rather than becoming its own
milestone.

## `D-M259-4` — the "60K skills" figure is now REFUTED, not merely unverified

`shared_libraries.md` graded "60K skills" **UNVERIFIED** on the correct reasoning that a public-only capture
cannot see private rows. The platform's own pre-consolidation total is **43,584**, so 60K is now positively
contradicted. The **≥42,790 public floor is vindicated**: the private remainder is 794 skills / 41 roles,
exactly the shape the floor language predicted. **M264 lands the corpus edit** — this milestone only
establishes the fact.

## `D-M259-5` — `skills-and-job-roles` is out of scope, permanently

The 12,201-skill ESCO repo (last commit 2024-04-08) is a **different lineage**, not an earlier version of the
same data. ESCO survives as *provenance* on the canon (`esco_uri`/`isco` on 454 of 706 roles, `onet_code` on
495), not as its source. Its "dormant ≥18 months" grading in `org-repos.md` **stands unchanged** — the
revamp did not wake it, and no milestone in v2.9 should pull or diff it.

## `D-M259-6` — the fresh pull invalidated 19 corpus anchors; they were REPAIRED, not deferred

Pulling both clone sets (M259's own first deliverable) moved `app` +175 and `next-web-app` +86, and
`repair_postcondition` immediately went RED on 19 line-anchors across 12 files. The first instinct was to route
them Fate-3 to M264 and record a baseline with a written reason.

**The ratchet refused** — *"refusing to raise the baseline — these sites are new and must be repaired, not
accepted"* — and it was right. So they were repaired here, against the freshly-pulled sources.

Two things worth carrying forward:

1. **They were not a re-numbering.** `app/main.go:314`'s cited sentinel-fold comment no longer exists in the
   file, and the `MESSENGER_ENABLED` / `CUSTOMERIO_SYNC_ENABLED` gate NAMES moved out of `main.go` into a
   **net-new `app/env_guards.go`** (`:61`, `:62`), while `main.go` still reads them at `:292`. An integer bump
   would have produced confident, wrong citations.
2. **A blanket string replacement broke two sibling anchors.** `platform-migration-status.md:155` carries a
   CHAINED citation — `app/main.go:1437`, then bare `:1458`, `:1460`, `:1485`, … — and swapping the leading
   file to `env_guards.go` silently re-pointed the bare continuations at a 208-line file. The guard caught it
   (`anchor-out-of-range`). **In a chained citation, never change the anchor that others hang off; make the
   continuations explicit instead.** Both were made explicit as part of the repair.

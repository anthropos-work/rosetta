# M40 Spec Notes

Authoritative design: [`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review,
2026-06-24; root-cause workflow w7t4wq2z4). The serve-grant is a **replay-side** synthesis — no seeding, no
platform-repo edits.

## Pre-flight audits — M40 (all sections)
KB-fidelity: **GREEN** (report: `kb-fidelity-audit.md`). Topic→doc→code triples:
- serve rows → `corpus/ops/snapshot-spec.md` (l.380-430) → `rext stack-snapshot/directus/structure.go`
- anonymous read → `corpus/ops/safety.md` §2.9 → cms Directus client (`DIRECTUS_TOKEN` blank, verified live)
- firewall admissibility → `corpus/ops/safety.md` / `firewall.go` `AssertStructuralMetadata` (relations/fields carry NO tenant-scope cols)

## ROOT CAUSE (verified live on demo-3, supersedes the original a/b/c framing)
The per-stack Directus has **`directus_relations = 0`, `directus_fields = 0`** — only collections + the 5 prod-public
permissions are synthesized. Consequences:
- **(a) `versions`** — needs ONLY a public-read `directus_permissions` row (system collection, no relation). Live: `GET /versions` anon → 403.
- **(b)/(c) the nested O2M/M2M aliases** (`simulations.sequences`, `simulations.library_categories`,
  `library_categories.macro_category`, `sequences.roles`, `sequences.task_checks.sub_checks`, …) are **UNKNOWN to
  Directus** without their `directus_relations` + `directus_fields` rows → `fields=sequences` returns
  `"… don't have permission … or it does not exist"`; `fields=sequences.scenario_intro` is silently stripped.
- **(c) does NOT need a platform nil-guard.** `sequences|read` is ALREADY granted on demo-3's public policy and
  `/items/sequences` serves the full `scenario_intro` directly; only the *nested* read fails, and `directus_relations=0`
  is the complete explanation. BOTH the library half AND the activity-feed half ship in tooling — no platform sign-off
  needed, no escalation. (The library/activity fork in overview.md "Open questions" is resolved: not a fork.)

## The fix (M40-D3): synthesize the relational web dynamically
In `structure.go`, capture (from the sanctioned `--dsn`, firewall-gated, version-robust — NO hardcoded list):
1. **`directus_relations`** rows whose `many_collection` AND `one_collection` both lie within the served-collection
   closure (the content collections + the library-category target/junction collections + macro-categories).
2. **`directus_fields`** rows for the served collections (so the alias fields + their interfaces register).
3. **The missing target/junction collections** as served collections: `library_categories`, `library_macro_categories`,
   `simulations_library_categories`, `skill_paths_library_categories` (+ register + public-read grant).
4. **`versions`** (directus_versions) public-read `directus_permissions` row.

Both `directus_relations`/`directus_fields` are admissible under `AssertStructuralMetadata` (zero tenant-scope columns,
verified). The junction/target collections get DDL via the existing dynamic structure capture (they are user collections).

## rext `stack-snapshot/directus/structure.go` — SYNTHESIZE public-read `directus_permissions` (on `PublicPolicyID`)
TODO: extend the structure replay to ADD public-read rows for the collections cms's anonymous path needs but prod's
public policy does not grant. NOTE: the existing `servePermissionsRowsSQL` only COPIES rows already on prod's public
policy → these three classes must be synthesized explicitly.

## (a) `directus_versions` (SYSTEM collection) — the dominant blocker
TODO: synthesize the public-read row that unblocks cms `skillpath.go:64` `GetSkillPath` →
`GetLatestOrCreateVersion` → `version.go:40` `GET /versions` (anon 403, treated as fatal) → unblocks the entire
skill-paths library + every sim/path detail page.

## (b) library-category collections — the sims-list 403
TODO: synthesize public-read rows for the library-category collections `ListPublicJobSimulations` (cms
`jobsimulation.go:305`) expands → fixes 403 → empty relation → `ToDomain` panic `"index out of range [0]"`.

## (c) `simulations.sequences` O2M nested read — the activity-feed strip
TODO: INVESTIGATE FIRST — under the public policy the O2M nested read is STRIPPED even with `sequences|read`
granted, so cms `GetJobSimulation` gets `s.Sequences==[]` → panics at `jobsimulation.go:1097` (`s.Sequences[0]`) →
the activity feed's per-row simulation federation returns null → feed empties. Determine whether the O2M is
grantable under the Directus public policy without a platform nil-guard (READ-ONLY — cannot edit cms). If not: ship
the library half independently; escalate the activity-feed half for platform sign-off.

## Activity DATA (already correct — NOT in scope)
TODO (verify only, do not seed): 21 completed sessions already in `jobsimulation.sessions` /
`public.local_jobsimulation_sessions`. This milestone is purely a serve-grant gap.

## Regression test — three surfaces serve >0 on a fresh demo
TODO: re-replay the snapshot into demo-3 and assert `/library/ai-simulations`, `/library/skill-paths`, and
`/profile/activities` each serve >0.

## KPI "AI simulations completed" = 0 — separate, re-verify after
TODO: re-verify after the feed fix; its source `public.local_jobsimulation_sessions = 21` has no CMS dependency, so
it is likely a separate frontend/auth-context issue — flag separately if it persists.

## Delivers → `corpus/ops/snapshot-spec.md`
TODO: author the public-policy serve-grant extension (the synthesized public-read rows on the `PublicPolicyID`, what
each unblocks, and the O2M-strip caveat).

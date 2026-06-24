# M40 Spec Notes

Authoritative design: [`.agentspace/profile_gaps.md`](../../../../.agentspace/profile_gaps.md) (live-demo review,
2026-06-24; root-cause workflow w7t4wq2z4). The serve-grant is a **replay-side** synthesis — no seeding, no
platform-repo edits.

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

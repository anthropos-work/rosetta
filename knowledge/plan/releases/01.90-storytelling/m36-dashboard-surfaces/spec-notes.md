# M36 — Spec notes

Authoritative design: [`.agentspace/seeding_gaps.md`](../../../../.agentspace/seeding_gaps.md) §3b (the
dashboard surfaces + the aggregates needing distributions), §3c (the claimed-vs-verified aha), §6a #5.

## The dashboard spine (REST `/api/workforce/*` → `app/internal/workforce/*.go`)
_(verified-side: `user_skill_evidences`; mapped-side: `membership_skills` — mapped must outnumber verified per
skill for a believable funnel. Members need `job_role_id`/name + `joined_at` + tags.)_

## Aggregates needing believable distributions
_(self-eval accuracy = `user_level` vs `anthropos_level`; verification funnel; role-readiness [needs
`job_role_skills`]; AI-readiness [AI-named skills]; succession [validation + interview rows]; growth arc
[early-low/late-high]. The two employee heroes = the dashboard's standout high/low rows.)_

## Scope hard line
_(seed the spine for the seeded story; do NOT chase every widget — the milestone's biggest growth risk.)_

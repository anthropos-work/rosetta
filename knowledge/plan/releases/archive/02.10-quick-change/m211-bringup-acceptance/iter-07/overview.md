---
iter: 07
milestone: M211
iteration_type: tik
iter_shape: bring-up-fix
status: planned
created: 2026-07-08
---

# iter-07 — tik: land the dev cold-path corrections (11-container count + reset-db extensions re-bootstrap)

**Active strategy reference:** TOK-01 "Warm-first cache-migrate, then cold-prove both stacks" — move (2)
"close the dev-side M25-D9 gap" + move (4)'s pre-cold-run hardening. Every tik routes a fix to its surface
(here: corpus) and re-measures.

**Step 0 — Re-survey.** TOK-01's next-tik direction (from iter-06 close) is: NEXT SESSION run (e) coverage +
Playthroughs + the full COLD `/dev-up`/`/demo-up`. Before spending a 20-45-min cold `/dev-up`, the cheap
pre-cold hardening is to fix the two KNOWN cold-`/dev-up` blockers surfaced but not yet closed:
1. iter-04 routed forward: the dev-up SKILL still says "**12 containers**" (health check "confirm 12 healthy
   containers") — post-merge the skiller container is gone; warm `docker ps` shows **11**. A cold `/dev-up`
   would fail its documented health assertion on a phantom 12th container.
2. M208 finding (overview.md Pre-surfaced requirement): `make reset-db` re-runs migrations against a wiped DB
   with **no `extensions` schema** → app/cms `vector(1536)` + gin-trigram migrations fail cold. The
   first-build flow already creates extensions before `make migrate`, but the reset-db path (auto-migrate)
   isn't documented to re-bootstrap extensions+schemas+policy first.

Both are cold-`/dev-up` correctness fixes; landing them BEFORE the cold run removes two known failures.

**Cluster / target identified:** the dev bring-up docs (`.claude/skills/dev-up/SKILL.md` +
`corpus/ops/setup_guide.md`). Evidence: iter-04's routed-forward observation + M208's pre-surfaced finding
(both named in overview.md). Corpus-only surface (no platform edit; no rext change needed).

**Hypothesis:** correcting the container count to 11 and documenting the reset-db extensions re-bootstrap
removes the two known cold-`/dev-up` blockers, so the eventual cold `/dev-up` (tik-11) passes its health +
migrate steps without a mid-run surprise.

**Expected lift:** metric (fully-met sub-conditions) stays 5/6 — this is prerequisite hardening, not a
sub-condition flip. The lift is de-risking the cold `/dev-up` proof (two known failures removed).

**Phase plan:** triage (done in Phase 0 survey) → route fix to corpus surface → re-measure (confirm no stale
"12 container" / undocumented-reset-db-extensions residue remains in the dev bring-up path).

**Escalation conditions:** none expected (corpus doc fix). If the container count turns out NOT to be 11
against the graphql profile definition, re-check before editing.

**Acceptable close-no-lift outcomes:** n/a — the target is concrete doc corrections; either they land
(closed-fixed) or the survey falsifies the need (closed-no-lift with the finding recorded).

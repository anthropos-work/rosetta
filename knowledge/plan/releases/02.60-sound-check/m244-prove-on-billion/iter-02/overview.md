---
iter: 02
milestone: M244
iteration_type: tik
status: closed-fixed
created: 2026-07-22
---

# iter-02 — Foundation tik (under TOK-01)

**Type:** tik. **Active strategy:** TOK-01 (staged cold billion bring-up → gate-parts a–h one-cluster-per-tik).

## Step 0 re-survey — TOK-01 target SUBSTITUTED (not a re-scope)
TOK-01 named tik-1 "from-scratch billion setup + scp secrets/cache." Re-survey found a **47 h-old M238-era demo-1 already running** under `devops@/home/devops/panorama/` (mis-pinned rext.tag=m237 / checkout=m238; autoverify 47 h stale) — unusable for M244, but the workspace **already has** secrets (60K) + snapshot cache (1.5G) + the rext clone. Substitution under the same TOK-01: **adopt the devops workspace → re-pin m243 → teardown stale (capture serve-reap) → cold bring-up.** No from-scratch setup, no 1.4G scp.

## Cluster / target
Establish the green cold demo on billion — the enabling precondition for gates (b)(c)(h) which all gate on a fresh green `autoverify.json` — and discharge the two checks that run before/around the bring-up: gate (a) ORG-CLEAN (read-only, first) and gate (e) DEF-M226-01 serve-reap (tested at teardown).

## Hypothesis / expected lift
Re-pin + cold reset-to-seed at m243 yields a fresh green autoverify + all peer origins serving; +2 gate parts (a, e) discharged; the green stack unblocks b/c/d/f/g/h.

## Phase plan
Pre-flight rung zero (pin) → gate (a) go-test cleanliness → purge teardown (gate e capture) → cold `--public-host` bring-up (login shell, supervised) → verify fresh green autoverify + origins + content.

## Escalation conditions
Bring-up BRINGUP_EXIT nonzero / ENOSPC (DEF-M239-01 class) → investigate; platform-source wall → ESCALATE.

## Acceptable close-no-lift
n/a — this tik lands the foundation.

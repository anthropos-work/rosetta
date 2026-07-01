# M53 — v1.10b "fit-up" cold-rebuild acceptance record

**Date:** 2026-07-01 · **Stack:** demo-1 (offset 10000) · **rext pin:** `v1.10.1` (authoring HEAD `e91f004`)
**Verdict:** **NOT GREEN — 5/6 acceptance criteria + academy F6 PASS from cold; AB4 (manager half) is a
release-BLOCKER routed back to M51.**

The single live demo the M47–M52 fixes were iterated against was **destroyed** (`/demo-down 1 --purge`, all
17 containers + network removed, ALL demo-1 images purged — M49 #6 verified) and **cold-rebuilt** from the
`v1.10.1` release tag by a single `/demo-up 1` (no manual steps). The acceptance bar was then asserted.

## The rext release tag
- **`v1.10.1`** (annotated) rolled on the authoring copy at `e91f004`, rolling up `fit-up-m47..m52` (46 rext
  commits) + the M53 academy F6 commit. `.agentspace/rext.tag` bumped to `v1.10.1`; the `stack-demo/rosetta-extensions`
  consumption clone re-pinned to `v1.10.1` via a clean fetch + checkout. Canonical pin recorded in
  `corpus/ops/rosetta_demo.md`.

## The acceptance bar

| # | Criterion (owner) | Result | Evidence |
|---|---|---|---|
| AB1 | all backends healthy — no silent `app Exited` (M47/M49) | ✅ PASS | 17 demo-1 containers **Up, 0 Exited**; autoverify: `/api/health` 200 on :18082, `sentinel.casbin_rules=1150`, all liveness+readiness probes passed |
| AB2 | cold-start snapshot filled with NO prompt (M47) | ✅ PASS | Set-dress replayed from the filled 1.4 GB cache with **no prompt**: taxonomy 330261 rows / directus 11986 / sim-embeddings 1490. `/demo-up` is replay-only (KB-1) — the cache was filled by M47's turnkey capture |
| AB3 | set-dress + seed (3 orgs incl. AI-readiness) + verify + cockpit — no #7 abort | ✅ PASS | `/demo-up` EXIT 0. Seed: **org rows=3** (Cervato, Solvantis, Northwind), 9-identity roster, ai-readiness-config 14 + funnel 1263, users 2700, profiles 8743, jobsim 1078, certs 374 …; cockpit serving :17700 |
| AB4 | **both-vantage M42 coverage GREEN on the existing orgs (M50)** | ❌ **BLOCKER** | **Employee (maya @ Cervato): GATE MET** (reachable=59, failing=0, escapes=0, persona=0, notReached=0, frontier=EXHAUSTED). **Manager (dan-manager @ Cervato, M50's gate): GATE NOT MET** — failingSections=**2**, both `/enterprise/workforce/ai-readiness` (empty for Cervato). See below + `decisions.md` AB4-REGRESSION |
| AB5 | AI-readiness dashboard criteria hold on the 3rd org (M51) | ✅ PASS | Manager `dana-manager` @ **Northwind Aviation**: GATE MET (reachable=70, failing=0, escapes=0). Dashboard renders **50/100 org readiness, 199 members, 173/199 (87%) functional+, the 3-step funnel + By-team grid** — fast (541 meaningful chars, no 180s timeout; the M51 loadMembers-bound patch applied). 199 matches the shipped funnel (KB-2) |
| AB6 | cockpit [Download manifest] = complete inlined `seed-generation-manifest.yaml` (M52) | ✅ PASS | `GET :17700/seed-generation-manifest.yaml` → HTTP 200, 7593 B, `application/x-yaml`, attachment. Complete: `population` (**all 3 orgs** + heroes), `generation` (prompt_template + batches + `max_cost_usd: 0.3`), `snapshot_sources`, `excludes` |
| F6 | academy: content + menu-link + non-anonymous session (M53) | ✅ PASS | (i) content real (copilot/claude-code/ai-engineering chapters render); (ii) **9 cockpit [Academy] links** → `http://localhost:13077/`, each `data-academy-persona="member"`; (iii) academy launched with **both** e2e_persona gates (`BENCHMARK_VISUAL_BYPASS=1` + `NEXT_PUBLIC_E2E_AUTH=1` in the running process env) → the cockpit link's `e2e_persona=member` cookie drives a signed-in member. Cosmo AI chat absent by design (no keys) — no `/api/ai/chat` assertion |

## The AB4 blocker (routed to M51 — NOT fixed in M53)

**What failed:** the **M50 canonical M42 manager gate** — `dan-manager` @ Cervato Systems (`run-coverage.sh 1
manager` default) — is RED from cold: `failingSections=2`, both on `/enterprise/workforce/ai-readiness`
(`ai-readiness-org-score` + `ai-readiness-funnel`, kind=`empty`). The page renders HTTP 200 with 0 ejects but
shows **"No AI readiness data yet for this org"** (org header = Cervato Systems) — because the 199 AI-readiness
snapshots are seeded ONLY for **Northwind Aviation** (a `closed` cycle — the M51 showcase-org design).

**Root cause (M51 iter-05 D3):** M51 added `/enterprise/workforce/ai-readiness` to `MANAGER_MANIFEST.seedPaths`
(`stack-verify/e2e/lib/coverage-manifest.ts:520`) **UNCONDITIONALLY**, so EVERY manager sweep asserts the funnel
renders — but the data is org-specific. M51's gate ran ONLY `dana-manager` @ Northwind (passes), so it never
re-ran the M50 `dan-manager` @ Cervato sweep and never saw the regression its seedPath introduced. M50 (closed
BEFORE M51) had a GREEN manager gate because the seedPath didn't exist yet.

**Why M53 caught it:** the fix-on-live serialization across M47..M52 never re-ran the M50 Cervato manager sweep
after M51's manifest change; M53's from-cold both-vantage assertion is the first joint re-measurement — exactly
the "cold-rebuild surfacing a late regression" M53 exists to catch. The dana-manager @ Northwind gate (AB5) is
GREEN, proving the *data + dashboard* are correct — only the *manifest conditioning* is wrong.

**The fix is M51's (org-condition the ai-readiness seedPath + descriptors — assert only for the showcase-org
manager vantage).** Per the acceptance-not-fix rule, M53 does NOT repair it. Since M51 is archived, this
escalates to the orchestrator/user to route (re-open M51 or a tracked follow-up) before v1.10b closes.

## Disposition
- **v1.10b is NOT ready to close** until AB4's manager half is GREEN again (with AB5 still GREEN).
- Everything else — the whole bring-up chain, all 3 orgs, the AI-readiness dashboard, the manifest download,
  the academy F6, the M49 #6 teardown — **works from cold at `v1.10.1`.**

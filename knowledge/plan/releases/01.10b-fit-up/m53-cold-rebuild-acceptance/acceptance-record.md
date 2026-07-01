# M53 — v1.10b "fit-up" cold-rebuild acceptance record

**Date:** 2026-07-01 · **Stack:** demo-1 (offset 10000) · **rext pin:** `v1.10.1` (authoring HEAD `117fe41`)
**Verdict:** **GREEN — 6/6 acceptance criteria + academy F6 PASS from cold.** AB4 was RED on first assertion
(an M51-owned gate regression); it was **fixed at the acceptance gate** (a conscious, recorded exception to
M53's no-fix-code rule — same class as the academy F6 exception — because M51 is archived) and **re-verified
GREEN on both manager vantages**.

The single live demo the M47–M52 fixes were iterated against was **destroyed** (`/demo-down 1 --purge`, all
17 containers + network removed, ALL demo-1 images purged — M49 #6 verified) and **cold-rebuilt** from the
`v1.10.1` release tag by a single `/demo-up 1` (no manual steps). The acceptance bar was then asserted.

## The rext release tag
- **`v1.10.1`** (annotated) — originally rolled on the authoring copy at `e91f004` (rolling up `fit-up-m47..m52`,
  46 rext commits + the M53 academy F6 commit); **re-rolled at the acceptance gate to `117fe41`** to include the
  AB4 org-conditional-manager-manifest fix (a local, unpushed tag re-roll — `git tag -d` + re-create annotated;
  NOT a force-push). `.agentspace/rext.tag` stays `v1.10.1`; the `stack-demo/rosetta-extensions` consumption
  clone re-pinned to the re-rolled `v1.10.1` via a clean fetch + checkout. Canonical pin recorded in
  `corpus/ops/rosetta_demo.md`.

## The acceptance bar

| # | Criterion (owner) | Result | Evidence |
|---|---|---|---|
| AB1 | all backends healthy — no silent `app Exited` (M47/M49) | ✅ PASS | 17 demo-1 containers **Up, 0 Exited**; autoverify: `/api/health` 200 on :18082, `sentinel.casbin_rules=1150`, all liveness+readiness probes passed |
| AB2 | cold-start snapshot filled with NO prompt (M47) | ✅ PASS | Set-dress replayed from the filled 1.4 GB cache with **no prompt**: taxonomy 330261 rows / directus 11986 / sim-embeddings 1490. `/demo-up` is replay-only (KB-1) — the cache was filled by M47's turnkey capture |
| AB3 | set-dress + seed (3 orgs incl. AI-readiness) + verify + cockpit — no #7 abort | ✅ PASS | `/demo-up` EXIT 0. Seed: **org rows=3** (Cervato, Solvantis, Northwind), 9-identity roster, ai-readiness-config 14 + funnel 1263, users 2700, profiles 8743, jobsim 1078, certs 374 …; cockpit serving :17700 |
| AB4 | **both-vantage M42 coverage GREEN on the existing orgs (M50)** | ✅ PASS (fixed at gate) | **Employee (maya @ Cervato): GATE MET** (reachable=59, failing=0, escapes=0, persona=0, notReached=0, frontier=EXHAUSTED). **Manager (dan-manager @ Cervato, M50's gate): GATE MET after the fix** — reachable=**69**, failingSections=**0** (was 2), escapes=0, persona=0, notReached=0, frontier=EXHAUSTED; the ai-readiness page is no longer primed/asserted for the base org. See "The AB4 fix" below + `decisions.md` AB4-FIX |
| AB5 | AI-readiness dashboard criteria hold on the 3rd org (M51) | ✅ PASS (re-verified post-fix) | Manager `dana-manager` @ **Northwind Aviation**: GATE MET (reachable=70, failing=0, escapes=0). **Both ai-readiness sections still PASS** (`ai-readiness-org-score` + `ai-readiness-funnel`, 541 meaningful chars) — the AB4 fix scopes the page to the showcase org, so Northwind still primes + asserts it. Dashboard renders **50/100 org readiness, 199 members, the 3-step funnel** — fast (no 180s timeout; the M51 loadMembers-bound patch applied). 199 matches the shipped funnel (KB-2) |
| AB6 | cockpit [Download manifest] = complete inlined `seed-generation-manifest.yaml` (M52) | ✅ PASS | `GET :17700/seed-generation-manifest.yaml` → HTTP 200, 7593 B, `application/x-yaml`, attachment. Complete: `population` (**all 3 orgs** + heroes), `generation` (prompt_template + batches + `max_cost_usd: 0.3`), `snapshot_sources`, `excludes` |
| F6 | academy: content + menu-link + non-anonymous session (M53) | ✅ PASS | (i) content real (copilot/claude-code/ai-engineering chapters render); (ii) **9 cockpit [Academy] links** → `http://localhost:13077/`, each `data-academy-persona="member"`; (iii) academy launched with **both** e2e_persona gates (`BENCHMARK_VISUAL_BYPASS=1` + `NEXT_PUBLIC_E2E_AUTH=1` in the running process env) → the cockpit link's `e2e_persona=member` cookie drives a signed-in member. Cosmo AI chat absent by design (no keys) — no `/api/ai/chat` assertion |

## The AB4 fix (an M51-owned gate regression, fixed at the M53 acceptance gate)

**What failed on first assertion:** the **M50 canonical M42 manager gate** — `dan-manager` @ Cervato Systems
(`run-coverage.sh 1 manager` default) — was RED from cold: `failingSections=2`, both on
`/enterprise/workforce/ai-readiness` (`ai-readiness-org-score` + `ai-readiness-funnel`, kind=`empty`). The page
rendered HTTP 200 with 0 ejects but showed **"No AI readiness data yet for this org"** (org header = Cervato
Systems) — because the 199 AI-readiness snapshots are seeded ONLY for **Northwind Aviation** (a `closed` cycle —
the M51 showcase-org design). A gate-correctness bug (0 escapes, persona green), NOT a content gap.

**Root cause (M51 iter-05 D3):** M51 added `/enterprise/workforce/ai-readiness` to `MANAGER_MANIFEST.seedPaths`
(`stack-verify/e2e/lib/coverage-manifest.ts`) **UNCONDITIONALLY**, so EVERY manager sweep primed + asserted the
funnel — but the data is org-specific. M51's gate ran ONLY `dana-manager` @ Northwind (passes), so it never
re-ran the M50 `dan-manager` @ Cervato sweep and never saw the regression its seedPath introduced. M50 (closed
BEFORE M51) had a GREEN manager gate because the seedPath didn't exist yet. M53's from-cold both-vantage
assertion is the first joint re-measurement — exactly the "cold-rebuild surfacing a late regression" M53 exists
to catch.

**The fix (rext test/gate artifact — zero platform edits; commit `117fe41`):** make the manager manifest
**org-conditional**. `manifestFor(vantage, expectedOrg)` now returns the full showcase manifest (`MANAGER_MANIFEST`,
which primes + asserts `/enterprise/workforce/ai-readiness`) ONLY when the org is the AI-readiness showcase org
(`AI_READINESS_SHOWCASE_ORG = 'Northwind Aviation'`, case-insensitive substring); for any other manager org it
returns a new `MANAGER_MANIFEST_BASE` that omits both the seedPath and the descriptor. `coverage.spec.ts` threads
`COVERAGE_EXPECTED_ORG` into `manifestFor`. +3 unit tests (showcase includes it, base/empty omits it, no
collateral drop); 27/27 manifest unit tests pass. The employee manifest path is unchanged.

**Why this is a sanctioned M53 exception:** per the acceptance-not-fix rule a failed assertion routes to its
owning milestone. M51 (archived) owns this regression. With M51 closed, the user **APPROVED fixing it at the
acceptance gate** — a conscious, recorded exception (the same class as the academy F6 exception, D1/D2) rather
than re-opening an archived milestone. Recorded in `decisions.md` AB4-FIX.

**Re-verification (both manager vantages, from the same cold demo-1):**
- `dan-manager` @ Cervato (M50 base-Workforce gate): **GATE MET** — reachable=**69/150**, failingSections=**0**
  (was 2), escapes=0, persona=0, notReached=0, frontier=EXHAUSTED. Confirmed ai-readiness is no longer in the
  reached page set (base manifest omits it). Persona all-ok; studio-desk cross-port follow OK.
- `dana-manager` @ Northwind (M51 showcase gate, AB5): **GATE MET** — reachable=**70/150**, failingSections=0,
  escapes=0, persona=0; ai-readiness seedPath IS crawled (position #3) and **both ai-readiness sections PASS**
  (541 meaningful chars) — the showcase gate is intact.

## Disposition
- **v1.10b is GREEN from cold at the re-rolled `v1.10.1` (`117fe41`)** — 6/6 acceptance criteria + academy F6.
- The whole bring-up chain, all 3 orgs, the AI-readiness dashboard, the manifest download, the academy F6, the
  M49 #6 teardown, and **both-vantage M42 coverage** now pass from cold.

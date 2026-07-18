**Type:** tik (under TOK-01 `reprove-hiring-on-billion`). Apply Finding-1's serve fix + measure the recruiter
vantage (C2/C3/C5) from this Mac. Protocol: verification.md + coverage-protocol.md + latency-budget.md.

# iter-03 — work log

1. **Consumed** `casting-call-m226-serve-hiring` (rext ee1bdf2) on billion (fetch + checkout + rext.tag); the
   running demo-1 stayed up (tooling clone only).
2. **Surgical `tailscale serve` re-apply** with the fixed `gen_tailscale_serve.py` (regenerate plan + apply, no
   rebuild) — `tailscale serve status` now fronts :13001.
3. **Confirmed reachable from this Mac:** peer `curl https://billion:13001` → exit 0 / **307** (was exit 35).
4. **Measured C2/C3/C5 from this Mac** (the peer vantage) — see scoreboard.

## Measurements (from this Mac → billion, HTTPS)

- **C2 — recruiter comparison ≥40 rows/each of 5 positions: ✅ GREEN.** `run-hiring-render.sh 1 rae-recruiter
  --hiring` COVERAGE_RENDER_GATE=1 PASSED. Recruiter lands on `/enterprise/activity-dashboard`; per-sim **network
  total = 44 for EACH of the 5 shared sims** (min 44 ≥ floor 40), page-1 renders 20, scores 61-100 (14-18 distinct
  per sim — non-degenerate/rankable), **junk=0, prod-ejects=0, 0 errors, any403=false**. Real candidates
  (Mei Costa, Owen Dubois @meridian-talent.com).
- **C3 — 2 candidate heroes render usable assessed profiles: ✅ GREEN.** `m224-candidate-heroes.spec.ts` (made
  remote-capable via `CANDIDATE_HOST`), 3/3 passed: Cara (assessed) `/home` shows a **Completed** hiring assignment
  + skills; Cody (assigned-only) `/home` shows an **Assigned/pending** assignment; Rae (recruiter, no-regression)
  reaches `/enterprise` with the full hiring nav re-skin. No ejects to :13000/login; real content.
- **C5 — recruiter p95 click→ACCESS < 5 s: ✅ GREEN.** `run-latency.sh 1 recruiter` (the new vantage), 5 cold
  runs, gated on a fresh-green autoverify.json copy — **reached ACCESS 5/5, p50 0.44 s, p95 1.50 s < 5.0 s.** The
  3rd measured access path (after v2.3's employee 1.46 s / manager 1.40 s). ERR_ABORTED lines = benign aborted
  RSC-prefetches, not login-path failures.
- **C4 — reads as hiring:** now **fully confirmed** — Rae's nav shows the hiring re-skin ("Results | AI
  Simulations | AI Interviews | Candidates Feedback"), on top of iter-02's is_hiring=true (DB) + cockpit label.

## Scoreboard after iter-03

| # | condition | verdict |
|---|-----------|---------|
| 1 | is_hiring + exactly 5 mgr / 45 cand | ⚠ 3 admin + 47 candidate (Finding-2, open) |
| 2 | recruiter comparison ≥40 rows / each 5 positions | ✅ GREEN (44 × 5, junk=0) |
| 3 | 2 candidate heroes usable assessed profiles | ✅ GREEN (Cara Completed, Cody Assigned) |
| 4 | reads as hiring | ✅ GREEN (is_hiring + cockpit label + nav re-skin) |
| 5 | recruiter p95 click→ACCESS < 5 s | ✅ GREEN (p95 1.50 s, 5/5 ACCESS) |
| 6 | coexists with 3 workforce orgs on cockpit | ✅ GREEN |
| 7 | 0 platform-repo edits | ✅ GREEN |

**6/7 GREEN. Only C1 (Finding-2) remains.**

## Close — 2026-07-17

**Outcome:** Finding-1's serve fix applied + WORKS (peer-reachable :13001); the 3 recruiter-vantage conditions
measured GREEN from this Mac — C2 (44 comparable rows × 5 sims, junk=0), C3 (both candidate profiles usable), C5
(recruiter p95 1.50 s < 5 s, the 3rd measured vantage). Scoreboard 3/7 → **6/7 GREEN**; only C1 (count) open.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (6/7 GREEN; C1 count discrepancy remains — Finding-2 routed)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — Outcome: continue (iter-04)
**Decisions:** iter-03 D1 (surgical serve re-apply, no rebuild); D2 (candidate-heroes remote parameterization).
**Side-deliverables:** rext `casting-call-m226-c3-remote` (9396adc) — the C3 harness remote-capability (`CANDIDATE_HOST`), proven live.
**Routes carried forward:**
- iter-04: resolve **Finding-2** (C1: seed 3 admin + 47 candidate vs gate's exactly 5+45). Seeder fix (place the 2 candidate heroes at candidate-index slots / compensate the admin count) so the population is 5 admin + 45 candidate — the gate is the contract. Handler `PROVE-M226-iter04-count-5-45`.
- iter-04/05: a reproducible **DEFAULT cold re-bring-up** at the fixed rext tag (so the *default* `up-injected.sh 1` itself fronts :13001 + the counts hold, no hand-holding) — the gate's reproducibility clause + fold the serve/recruiter fix into the default path.
**Lessons:** Finding-1's fix was clean + cheap (surgical serve re-apply, no rebuild) — the demo-patches + images were untouched, only the tailscale-serve front changed. The recruiter vantage p95 (1.50 s) matches the employee/manager profile — the hiring app shares the authenticated-shell fast path. The candidate-heroes localhost-hardcoding was the SAME remote-capability gap as Finding-1 — M226 is systematically surfacing the hiring tooling's localhost-only assumptions (proven only on M225 localhost).

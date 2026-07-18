**Type:** tik (under TOK-01 `reprove-hiring-on-billion`). The first live cold bring-up on billion + the first
7-condition measurement from this Mac. Protocol: verification.md + coverage-protocol.md + latency-budget.md.

# iter-02 — work log

1. **Teardown** the stale v2.3.2 `demo-1` (`rosetta-demo down 1 --purge`): 16 containers removed, data purged,
   images reclaimed, hostlock released. VERIFIED from this Mac: base ports refuse (curl exit 7); on-host ss = 0
   listeners (cockpit + academy reaped); `tailscale serve` = "No serve config". Clean M221 F5/F5b/F12, 0 orphans.
2. **Substrate cutover:** confirmed `sections`↔`harden` is test-only (3 test files, 0 runtime); cut billion's
   rext → `casting-call-m225-harden` (be431c3) + rext.tag SoT; pinned all 12 platform clones to the local M225
   `clones.lock.json` SHAs (next-web-app 012a58578→**23bdbb5db** = the demo-patch `pre_sha256` pin; app, cms cut
   back too). billion substrate == the local GREEN proof.
3. **Default cold `up-injected.sh 1` (NO FLAGS)** — remote-reach auto-discovered billion (all 6 rungs). Ran
   SYNCHRONOUSLY via a tethered background ssh (remote foreground; never detached-on-host; monitored by
   file-write/image liveness). **~10.5 min, rc=0.** (One false-start: pre-flight "Go NOT on PATH" — Go IS
   installed at /usr/local/go/bin, a non-interactive-ssh PATH gap; clean fail-loud, re-launched with the PATH
   export. Not a gate finding.)
4. **autoverify ALL GREEN:** demo-patches all applied (0 refused/skipped — the substrate pin held), taxonomy
   42790 skills, isolation clean (55 writes / 74869 rows on demo-1 only, 0 shared/external), **hiring org
   set-dressed: 5 shared positions + 294 candidate HIRING sessions**. autoverify.json green ts=07:03:38Z.
5. **First 7-condition measurement FROM THIS MAC** — see the scoreboard below.

## The 7-condition scoreboard (iter-02, first measurement)

| # | condition | verdict | evidence |
|---|-----------|---------|----------|
| 1 | hiring org + is_hiring + **exactly 5 mgr / 45 cand** | ⚠ DISCREPANCY | Meridian Talent, is_hiring=true, 50 total, but **3 admin + 47 candidate** (Finding-2) |
| 2 | recruiter comparison **≥40 rows/each 5 positions** | ⛔ BLOCKED | data GREEN (294 sess / 5 sims ≈ 59 avg > 40) but RENDER unreachable — 13001 not tailscale-served (Finding-1) |
| 3 | 2 candidate heroes render usable profiles | ⛔ BLOCKED | candidates land on 13001 (same serve gap, Finding-1) |
| 4 | org reads as hiring | ✅ GREEN | is_hiring=true (DB) + cockpit "Hiring/Meridian" label + roster org_is_hiring (full nav render → iter-03) |
| 5 | **recruiter p95 click→ACCESS < 5 s** | ⛔ BLOCKED | recruiter lands on 13001 (serve gap); harness recruiter vantage PREPARED |
| 6 | coexists with 3 workforce orgs on cockpit | ✅ GREEN | cockpit shows 4 orgs: 2 workforce (Maya/Tom; Dan/Sara/Nick) + Northwind (Aria/Ben/Dana) + Meridian-hiring (Rae/Cara/Cody) |
| 7 | 0 platform-repo edits | ✅ GREEN | next-web/app/platform/jobsim 0-dirty; cms `?? studio/` = disclosed M221 D-05h; all demo-patches self-reverted |

**3 GREEN (4,6,7) · 1 discrepancy (1) · 3 blocked-by-serve-gap (2,3,5).** Two findings, both attributed.

## Findings (attributed to surface, per TOK-01 step 5)

**Finding-1 (R1-class, live-only) — the hiring app :3001 is not tailscale-served.** `gen_tailscale_serve.py`'s
`UI_BROWSER_FACING` never included the M224 hiring 2nd app, so `:13001` has no HTTPS front over the tailnet →
the recruiter + candidates (who land on the hiring app) are unreachable from a peer. Reachable on localhost (the
M225 proof), dead cross-machine — the M215/M221 lesson. **Fix PREPARED + committed + tagged**
(`casting-call-m226-serve-hiring`, rext ee1bdf2): add `("hiring", 3001)` to `UI_BROWSER_FACING` + the recruiter
latency vantage. Sanity-checked (serve_commands fronts 13001; 107 serve tests pass). APPLY + PROVE in iter-03.
Empirically confirmed live: `tailscale serve status` lacks 13001; peer `curl https://billion:13001` → exit 35.

**Finding-2 — the hiring org seeds 3 admin + 47 candidate, not the gate's exactly 5 + 45.** The seed preset
intends "5 admin + 45 candidate (first 10% of 50 = 5 admin)". But the 3 hiring-story heroes occupy population
slots 0,1,2 (the admin-index range); Rae(0) stays admin, but Cara(1) + Cody(2) are candidate heroes whose
role-override forces `candidate` → 2 would-be-admin slots become candidate → **3 admin + 47 candidate**,
deterministically (the local M225 proof had this too — the "exactly 5+45" was never literally produced). Routed
to iter-03 for a decision: seeder fix to hit exactly 5+45 (the gate is the contract) vs a gate-wording correction
(3 recruiters + 47 candidates is functionally equivalent). Leans seeder-fix. Handler `PROVE-M226-iter03-count-5-45`.

## Close — 2026-07-17

**Outcome:** First casting-call demo stood up GREEN on billion (cold reset-to-seed, default no-flags, ~10.5 min);
first 7-condition measurement complete — 3 GREEN (4,6,7), 1 discrepancy (1), 3 blocked-by-Finding-1 (2,3,5); two
findings precisely attributed, Finding-1 fix committed+tagged for iter-03.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (3/7 GREEN; 3 blocked by Finding-1; 1 discrepancy Finding-2 — fixes prepared/routed)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — Outcome: continue (iter-03)
**Decisions:** iter-02 D1 (pin to lock SHAs), D2 (rext→harden), D3 (teardown --purge); D4 (Finding-1 fix); D5 (Finding-2 attribution).
**Side-deliverables:** rext `casting-call-m226-serve-hiring` (ee1bdf2) — the Finding-1 serve+recruiter fix, prepared+tagged, to be applied+proven in iter-03.
**Routes carried forward:**
- iter-03: consume `casting-call-m226-serve-hiring` on billion + surgical `tailscale serve` re-apply (add 13001) → measure C2 (recruiter render ≥40×5), C3 (candidate profiles), C5 (recruiter p95) from this Mac. Handler `PROVE-M226-iter03-recruiter-vantage`.
- iter-03/04: resolve Finding-2 (3+47 vs 5+45) — seeder fix vs gate-wording. Handler `PROVE-M226-iter03-count-5-45`.
- later: a reproducible DEFAULT cold re-bring-up at the fixed tag (so the default bring-up itself fronts 13001 + the counts hold) — the gate's reproducibility clause.
**Lessons:** The M215/M221 "last breakage is cross-machine" held precisely — the one real bug (serve gap) was invisible on localhost. The substrate pin to the local lock SHAs made every demo-patch apply clean (0 drift). billion's 7.3 GiB held (load ~1.8, no OOM) for the 2-app demo. attribute-before-fix worked: both findings were characterized to their exact surface before any change.

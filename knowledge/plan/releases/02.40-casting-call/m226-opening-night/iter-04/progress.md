**Type:** tik (under TOK-01). Fix Finding-2 (C1 count → 5+45) + a reproducible DEFAULT cold re-bring-up proving
all 7 conditions. Protocol: verification.md + coverage-protocol.md + latency-budget.md.

# iter-04 — work log

1. **Root cause (confirmed in users.go):** `roleForHero` vantage-fidelity (hiring end-user hero → candidate;
   funnel requires it) + heroes ride the first slots (M38-D7) + `roleForIndex` first-5-slots-admin → Cara(2) +
   Cody(3) override 2 admin-band slots → 3 admin + 47 candidate. The preset's own "→5+45" comment never held.
2. **Fix (one-line preset):** `role_mix.admin 0.1 → 0.14` → admin band int(50×0.14)=7 → 5 real admins + 45
   candidate. Updated `presets_test.go` (asserts 0.14 + 7-slot band). Seeding Go tests green. Committed+tagged
   `casting-call-m226-count-5-45` (5d0297e — bundles serve-hiring + c3-remote + count).
3. **DEFAULT cold re-bring-up #1 FAILED** (rc=1): `bind 0.0.0.0:13001 — address already in use`. Root cause =
   **Finding-3**: an orphaned tailscale-serve front on :13001 from iter-03's surgical re-apply (added :13001 to
   the live serve but not the on-disk reset plan) → the teardown ran the stale iter-02 reset plan (no 13001) →
   didn't clear it. Recovered: `tailscale serve --https=13001 off` → teardown --purge → clean slate.
4. **DEFAULT cold re-bring-up #2 GREEN** (rc=0, ~10 min): 13001 bound cleanly, autoverify all-green, hiring org
   5 positions + 284 sessions. **The default path fronted :13001 (Finding-1 fix proven in the default path).**
5. **Re-measured ALL 7 from this Mac.**

## The 7-condition scoreboard (iter-04b — default cold cycle at the fixed tag)

| # | condition | verdict | evidence |
|---|-----------|---------|----------|
| 1 | is_hiring + **exactly 5 mgr / 45 cand** | ✅ GREEN | Meridian Talent, is_hiring=true, **5 admin + 45 candidate** (Finding-2 fixed) |
| 2 | recruiter comparison **≥40 rows/each 5 positions** | ✅ GREEN | 5 shared positions × **42 comparable candidates** each (DB) + rendered 42×5, junk=0, prod-ejects=0 |
| 3 | 2 candidate heroes usable assessed profiles | ✅ GREEN | Cara (Completed) / Cody (Assigned) render on /home |
| 4 | reads as hiring | ✅ GREEN | is_hiring=true + Rae nav "Results \| AI Simulations \| AI Interviews \| Candidates Feedback" |
| 5 | **recruiter p95 click→ACCESS < 5 s** | ✅ GREEN | ACCESS 5/5, p50 0.66 s, **p95 1.09 s** (green gate fresh) |
| 6 | coexists with 3 workforce orgs | ✅ GREEN | cockpit 4 orgs (2 workforce + Northwind + Meridian-hiring) |
| 7 | 0 platform-repo edits | ✅ GREEN | clones clean; cms `?? studio/` disclosed |

**7/7 GREEN on a default cold reset-to-seed cycle.** Finding-1 proven IN THE DEFAULT PATH (:13001 fronted, no
hand-holding).

## Findings

**Finding-3 (artifact, self-resolved).** iter-03's SURGICAL serve re-apply added :13001 to the live serve but not
the on-disk reset plan → the iter-04 teardown didn't clear it → the re-bring-up's :13001 compose bind conflicted
(M215 F12). NOT the default flow (a full default bring-up writes a reset plan including :13001, so teardown clears
it — proven: re-bring-up #2 came up clean after the orphan was cleared). Minor hardening candidate: the pre-bind
reap could also clear stale serve fronts on the demo's offset ports. Routed to harden/iter-05.

**Finding-4 (C2 harness insights-capture race).** On a fresh stack, `run-hiring-render`'s capture of the
`insightsByJobSimulations` network response is timing-fragile — 3 of 5 runs returned `list=0` (query fired but the
harness missed the window), 2 returned `list=5` + PASSED (42×5). The M215/M221 "the harness is the fragile thing"
lesson. NOT a product issue (the data is present, C3-Rae renders the Results scoreboard). Fix (rext harness
robustness) → iter-05, so the 2nd cold cycle's C2 is reliably automated-green. Handler `HARDEN-M226-iter05-c2-race`.

## Close — 2026-07-17

**Outcome:** Finding-2 fixed (C1 = exactly 5+45); a default cold re-bring-up at the fixed tag proved 7/7 GREEN
(Finding-1's serve fix in the default path, the 5+45 counts, all recruiter/candidate conditions). Surfaced +
attributed Finding-3 (orphan, self-resolved) + Finding-4 (C2 harness race, routed).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (7/7 on ONE cold cycle — the gate's "reproducibly" clause wants a 2nd clean cycle; and C2's harness needs the Finding-4 fix for a reliable automated pass) — PROVISIONALLY MET, per the M221 precedent
**Phase 5 grading:** (1) gate-met: n (provisional — 1 cycle, C2 flaky) — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — Outcome: continue (iter-05)
**Decisions:** iter-04 D1 (preset bump 0.14), D2 (test update), D3 (full default re-bring-up); Finding-3, Finding-4.
**Side-deliverables:** rext `casting-call-m226-count-5-45` (5d0297e) — the Finding-2 count fix + updated preset test.
**Routes carried forward:**
- iter-05: **fix Finding-4** (C2 harness insights-capture robustness — poll/extend the wait) + a **2nd clean default cold cycle** → 7/7 reliably automated-green → the gate's "reproducibly" clause MET. Handler `HARDEN-M226-iter05-c2-race`.
- harden: Finding-3 (pre-bind reap clears stale serve fronts on offset ports) — a nice-to-have.
**Lessons:** The billion proof surfaced FOUR findings invisible on localhost — the serve gap (F1), the count displacement (F2), the surgical-re-apply orphan (F3), the harness capture race (F4) — three of them harness/tooling, exactly the M215/M221 pattern. The count fix (0.14) held; C2's margin (44→42) stayed above 40. Provisional 7/7 on one cold cycle; reproducibility needs the 2nd.

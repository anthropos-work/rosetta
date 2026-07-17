# M226 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | Phase 0b KB-fidelity GREEN; billion recon (stale v2.3 demo up, prereqs green, C-6 mem risk, rext cutover needed); authored TOK-01 `reprove-hiring-on-billion` | 0/7 (no lift — tok) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | teardown stale v2.3 + substrate cutover (rext→m225-harden, 12 clones pinned to lock SHAs) + default cold `up-injected.sh 1` GREEN (~10.5 min, rc=0) + FIRST 7-condition measurement from this Mac | **3/7 GREEN (C4,6,7); C1 discrepancy 3+47 vs 5+45; C2/3/5 blocked by the :13001 serve gap** | closed-fixed — see iter-02/progress.md |
| iter-03 | tik | applied Finding-1 fix (consume `casting-call-m226-serve-hiring` + surgical serve re-apply → :13001 peer-reachable); measured C2/C3/C5 from this Mac | **6/7 GREEN — C2 (44×5, junk=0), C3 (Cara/Cody usable), C5 (recruiter p95 1.50 s <5 s); C4 fully confirmed. Only C1 (count) open** | closed-fixed — see iter-03/progress.md |
| iter-04 | tik | Finding-2 fix (`role_mix.admin 0.1→0.14`, tag `casting-call-m226-count-5-45`) + full DEFAULT cold re-bring-up at the fixed tag + re-measured all 7 | **7/7 GREEN on one default cold cycle (PROVISIONAL — C1 now 5+45; Finding-1 proven in default path). Findings 3 (orphan, self-resolved) + 4 (C2 harness race) surfaced** | closed-fixed — see iter-04/progress.md |
| iter-05 | tik | Finding-4 fix (C2 harness insights-capture poll, tag `casting-call-m226-c2-race-fix`) + 2nd clean DEFAULT cold cycle + re-measured all 7 (C2 reliable) + corpus fold-ins | **7/7 GREEN — 2nd cold cycle. GATE MET (2 clean cold cycles, both 7/7, 0 platform edits)** | closed-fixed (GATE MET) — see iter-05/progress.md |

## GATE MET — 2026-07-17

**The M226 7-condition exit gate is MET.** On `billion.taildc510.ts.net`, a **default `/demo-up 1` (no flags)**
yielded, **reproducibly across 2 clean cold reset-to-seed cycles** (iter-04b + iter-05), measured from this Mac
(the tailnet peer): (1) hiring org `is_hiring=true`, **exactly 5 admin + 45 candidate**; (2) recruiter comparison
**42 comparable candidates × each of the 5 positions** (≥40), junk=0, prod-ejects=0; (3) both candidate heroes
(Cara Completed / Cody Assigned) render usable profiles; (4) reads as hiring (nav "Results"); (5) recruiter **p95
click→ACCESS 1.09 s / 2.36 s < 5 s** (the 3rd measured vantage); (6) coexists with the 3 workforce orgs (12 heroes
/ 4 orgs on the cockpit); (7) **0 platform-repo edits**. rext code-of-record `casting-call-m226-c2-race-fix`
(4bd68ff). Four cross-machine findings surfaced + fixed (all tooling/harness/seed, 0 platform edits): F1 serve
gap, F2 count displacement, F3 surgical-orphan (self-resolved), F4 harness race.

**Next:** `/developer-kit:harden-mstone-iters` (final pass) → `/developer-kit:close-milestone`. Optional harden:
Finding-3 (pre-bind reap clears stale serve fronts on offset ports — a nice-to-have, not gate-blocking).

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

## Next iter

**iter-05 (tik, under TOK-01):** confirm **reproducibility** (the gate's "reproducibly" clause) + fix Finding-4.
Handler `HARDEN-M226-iter05-c2-race`. Steps: (1) fix the C2 harness insights-capture race (rext `stack-verify`:
poll/extend the `insightsByJobSimulations` capture wait so a cold-stack query is reliably caught); (2) a **2nd
clean default cold cycle** (`down 1 --purge` → `up-injected.sh 1` NO FLAGS) → re-measure all 7 → **7/7 reliably
automated-green** → the gate is firmly MET (2 cold cycles, per the M221 one-cycle-is-provisional precedent).
Optional harden: Finding-3 (pre-bind reap clears stale serve fronts on offset ports).

# M226 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | Phase 0b KB-fidelity GREEN; billion recon (stale v2.3 demo up, prereqs green, C-6 mem risk, rext cutover needed); authored TOK-01 `reprove-hiring-on-billion` | 0/7 (no lift — tok) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | teardown stale v2.3 + substrate cutover (rext→m225-harden, 12 clones pinned to lock SHAs) + default cold `up-injected.sh 1` GREEN (~10.5 min, rc=0) + FIRST 7-condition measurement from this Mac | **3/7 GREEN (C4,6,7); C1 discrepancy 3+47 vs 5+45; C2/3/5 blocked by the :13001 serve gap** | closed-fixed — see iter-02/progress.md |
| iter-03 | tik | applied Finding-1 fix (consume `casting-call-m226-serve-hiring` + surgical serve re-apply → :13001 peer-reachable); measured C2/C3/C5 from this Mac | **6/7 GREEN — C2 (44×5, junk=0), C3 (Cara/Cody usable), C5 (recruiter p95 1.50 s <5 s); C4 fully confirmed. Only C1 (count) open** | closed-fixed — see iter-03/progress.md |

## Next iter

**iter-04 (tik, under TOK-01):** resolve **Finding-2** (C1: the hiring org seeds 3 admin + 47 candidate, not the
gate's exactly 5+45). Handler `PROVE-M226-iter04-count-5-45`. The 2 candidate heroes (Cara, Cody) occupy admin-index
population slots (0,1,2) and override to `candidate`, dropping admins 5→3. Seeder fix (rext `stack-seeding`): place
the candidate heroes at candidate-index slots (or compensate the admin count) so the population is exactly 5 admin
+ 45 candidate — the gate is the contract. Then a **reproducible DEFAULT cold re-bring-up** at the fixed rext tag
(so the *default* `up-injected.sh 1` itself fronts :13001 + the counts hold, no hand-holding) — the gate's
reproducibility clause. Re-measure all 7 from this Mac to confirm 7/7 GREEN reproducibly.

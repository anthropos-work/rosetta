# M228 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | Phase 0b KB-fidelity GREEN; billion recon (M226 demo UP @ c2-race-fix, devops-operated, 20 images cached, 7 serve fronts); authored TOK-01 `reprove-corrected-hiring-on-billion` | 0/7 (no lift — tok) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | teardown M226 demo + rext cutover → `casting-call-m227-sections` + default cold `up-injected.sh 1` GREEN (UP_RC=0) + FIRST measurement from this Mac | **5/7 GREEN (C1,3,5,6,7 + fix#2); C2/C4 blocked by F2 (training leak) + F3 (2-sim); F1 seed FK error** | closed-fixed — see iter-02/progress.md |

## Next-iter queue
- iter-03 (tik): root-cause + FIX **F1** (succession/interview_extraction_results FK), **F2** (hiring-only training leak — find the exact 2nd-session seeder + add the fix#1 guard), **F3** (1-sim/candidate leak). Then default cold re-bring-up + warm re-measure C2 (5/5) + C4 (hiring-only). Handler `PROVE-M228-iter03-fix-believability-leaks`.

## Findings ledger (open)
- **F1** succession/interview_extraction_results FK seed error (AI-readiness Org C) — seed reports "failed" (non-fatal). → iter-03.
- **F2** fix#1 hiring-only GAP — 2 TRAINING sims leak (7 candidate sessions). → iter-03.
- **F3** fix#3 1-sim/candidate GAP — 17/45 candidates on 2 sessions. → iter-03.
- **F4** RESOLVED (reads-as-hiring GREEN; render-harness nav check was a cold-timing false-read).
- **F5** cold-tailnet warm-before-gate (M226 F5 recurrence) — measurement methodology; warm passes. → iter-03 warm re-measure.

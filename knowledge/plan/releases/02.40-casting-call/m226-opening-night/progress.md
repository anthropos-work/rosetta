# M226 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | Phase 0b KB-fidelity GREEN; billion recon (stale v2.3 demo up, prereqs green, C-6 mem risk, rext cutover needed); authored TOK-01 `reprove-hiring-on-billion` | 0/7 (no lift — tok) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | teardown stale v2.3 + substrate cutover (rext→m225-harden, 12 clones pinned to lock SHAs) + default cold `up-injected.sh 1` GREEN (~10.5 min, rc=0) + FIRST 7-condition measurement from this Mac | **3/7 GREEN (C4,6,7); C1 discrepancy 3+47 vs 5+45; C2/3/5 blocked by the :13001 serve gap** | closed-fixed — see iter-02/progress.md |

## Next iter

**iter-03 (tik, under TOK-01):** apply Finding-1's fix + measure the recruiter vantage. Handler
`PROVE-M226-iter03-recruiter-vantage`. Steps: consume `casting-call-m226-serve-hiring` (rext ee1bdf2) on billion +
surgical `tailscale serve` re-apply (front :13001) → confirm https://billion:13001 reachable from this Mac →
measure **C2** (recruiter comparison ≥40 rows / each of 5 sims, `run-hiring-render.sh 1 rae-recruiter --hiring`
COVERAGE_RENDER_GATE=1), **C3** (Cara/Cody candidate profiles), **C5** (recruiter p95 click→ACCESS < 5 s,
`run-latency.sh 1 recruiter`, gated on fresh-green autoverify.json) from this Mac. Also decide **Finding-2**
(3+47 vs exactly 5+45 — seeder fix vs gate-wording; handler `PROVE-M226-iter03-count-5-45`). Later: a reproducible
DEFAULT cold re-bring-up at the fixed tag (the gate's reproducibility clause + the serve fix in the default path).

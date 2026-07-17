# M226 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | Phase 0b KB-fidelity GREEN; billion recon (stale v2.3 demo up, prereqs green, C-6 mem risk, rext cutover needed); authored TOK-01 `reprove-hiring-on-billion` | 0/7 (no lift — tok) | closed-fixed — see iter-01/progress.md |

## Next iter

**iter-02 (tik, under TOK-01):** the substrate cutover + first default `/demo-up 1` on billion + first 7-condition
measurement from this Mac. Handler `PROVE-M226-iter02-first-cold-bringup`. Steps: cold-teardown the stale v2.3.2
`demo-1` (serve reset + academy respawner reap, M221 F5/F5b/F12; verify base ports freed + no survivor from this
Mac) → cut billion's rext (+ platform) over to `casting-call-m225-harden` (confirm `sections`↔`harden` test-only)
→ run a default cold `up-injected.sh 1` synchronously (NO FLAGS; ~15–25 min 2-app rebuild; never detach) → measure
the 7-condition gate FROM THIS MAC. Attribute every failing condition (R1 render re-surface / R4 45×5 hydration
latency / Clerkenstein-seed wiring / OOM) to its surface before any fix.

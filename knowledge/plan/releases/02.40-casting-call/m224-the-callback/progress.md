# M224 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | KB-fidelity GREEN (hiring.md FAPI-pointer fix inline); authored TOK-01 (recruiter-render-first) | baseline UNMEASURED (presumed 0 rows) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | Clerkenstein org `publicMetadata.isHiring` wired end-to-end (seeder roster → FAPI); `/align-run` GREEN 100/100 ×2; rext tag `casting-call-m224-iter02` | UNCHANGED (fix-half/scaffold — no render yet) | closed-fixed — see iter-02/progress.md |
| iter-03 | tik | recruiter cockpit seat (Rae Ramirez, manager→admin, slot-1, funnel-skipped) + `curatedTalent` skill family + manifest regenerated; rext tag `casting-call-m224-iter03` | UNCHANGED (scaffold — no render yet) | closed-fixed — see iter-03/progress.md |

## Next iter

**iter-04 (tik, under TOK-01) — THE FIRST GATE READING:** bring up a **LOCAL demo** consuming rext
`casting-call-m224-iter03` (offset ports; the M223 funnel seeds 45 candidates × 5 sims + the recruiter seat +
isHiring re-skin). Author a recruiter **render-probe** under `stack-verify/e2e/` (log in as `rae-recruiter` via
the cockpit handshake → reach each of the 5 `[simId]` comparison scoreboards → count comparable candidate rows +
capture the score distribution + run closure/eject checks). Take the **baseline** `min(rows-per-sim)` and
**attribute** the gap (Clerkenstein-identity vs render-gate vs seed) per hiring.md's read-path — **no fix before
the attribution** (the M219 trap). NB: 9.7 GiB Docker VM (< the 12 GB UI-tier prereq) — watch for a next-web OOM;
`demo-1-*` images are cached. Post-gate: the 2 candidate `/profile` heroes (funnel/candidate-role hero-awareness).

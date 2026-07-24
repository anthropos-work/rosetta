**Type:** tik — TOK-01 cluster 1 (gate a green). The AI-readiness demopatch re-point + green re-verify.

# M254 · iter-03 — progress

## Close — 2026-07-24

**Outcome:** re-authored the drifted `app-aireadiness-snapshot-loadmembers` demopatch for the consolidated
app (path/anchor/replacement/shas), shipped to origin (tag `july-jitter-m254-aireadiness-repoint`), re-pinned
billion, cold reset-to-seed → **autoverify green:true, 0 warnings**. Gate (a) MET.
**Type:** tik
**Status:** closed-fixed
**Gate:** MET (a — consolidated stack builds + comes up GREEN, fresh green autoverify, cold reset-to-seed,
0 platform edits)
**Phase 5 grading:** (1) gate-met: n (gate (a) is one of 8 parts; overall gate not yet met — b–h pending) —
(2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 2 of 5) —
(6) protocol-stop: n — Outcome: continue (iter-04 = read-only sweeps b/c/d/g)
**Decisions:** D1 (demopatch re-point authored against app@3df8536, manifest-only fix).
**Side-deliverables:** none.
**Routes carried forward:** gates b–h (iter-04 read-only sweeps → iter-05 latency → iter-06 mutating tail).
**Lessons:** (1) The self-healing apply_patch gate keys on the ANCHOR, not the path — but a re-pointed `path`
is required when the FILE moves (self-heal can't find an anchor in a file that doesn't exist). The
consolidation moved the whole read-path package, so a path re-point + anchor re-author was needed, not just a
sha re-pin. (2) The robust re-verify pattern (detached `setsid nohup` on billion + local short-poll for a
`.done` sentinel) completed cleanly in ~12 min and survives ssh disconnects — the standard for every
subsequent long billion op.

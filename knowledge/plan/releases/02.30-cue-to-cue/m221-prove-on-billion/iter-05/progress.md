# iter-05 (tik) — Phase C cycle 2: fix the dominant blocker + re-prove

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST").

## Close — 2026-07-15

**Outcome:** Landed F1 (the store-root shadowing fix, 3 places), F5 (native-academy supervised-respawn reap) and
F5b (gate-8 generated-file dirt) — all RED-fenced, shipped as rext `cue-to-cue-m221-r3`. **But the live re-prove
FALSIFIED F1 on the box:** a DEFAULT cold `up-injected.sh 1` on billion came up `green:false, warnings:3`,
**`public.skills=0`** (unchanged from the iter-04 baseline), directus exited(1) on a postgres `get_namespace_oid`
schema error. The shipped automated store resolution did NOT reproduce iter-04's manual `STACKSNAP_STORE` pin
(which loaded 42,790 rows) against billion's real directory topology. Gate distance did NOT move — still ~3/8.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET — re-prove came up non-green (skills=0, directus schema failure); the F1 fix is insufficient on
the live box. F5/F5b landed and RED-fenced but are hygiene fixes that do not move the 8-gate metric directly.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (tik #4 of 5) — (6) protocol-stop: n — Outcome: continue → iter-06 re-diagnoses F1 on the box.
**Decisions:** D-M221-05a..05f (iter-05/decisions.md). 05f records the orchestrator-graded falsification (the
iter-05 agent ran out of budget after landing the fixes but before grading the re-prove).
**Side-deliverables:** F5 (supervised-respawn reap) + F5b (gate-8 dirt) landed clean and RED-fenced — real fixes,
but not gate-movers this iter.
**Routes carried forward (Fate-3 → iter-06):**
- **F1 re-diagnosis (dominant):** WHY does the store resolution still land on an empty root on billion despite
  fixes A (resolver-prefers-snapshots) + B (`dev-setdress.sh --store`)? The manual pin worked; the shipped
  resolution did not. Is `up-injected.sh`'s demo set-dress path even invoking the fixed `dev-setdress.sh` seam,
  or a different entry? Is the `--store` root computed wrong on billion's layout?
- **The directus schema failure:** is it a consequence of the empty taxonomy (no schema → `get_namespace_oid`
  fails), or an independent ordering bug? (F2 rides this.)
- **F4** (academy client-side render), **F10** (field-exercise freshness-abort + `assert_ports_free`),
  **BURNIN-M221-dev-public-host**, **F-M220-4** — not reached; later tik.
**Lessons:** "de-risked to certainty" in a baseline diagnosis (a manual pin that worked) is NOT the same as "the
shipped automated fix reproduces it on the live box." The manual `STACKSNAP_STORE=... ` pin loading 42,790 rows
proved the *store exists and is loadable*; it did not prove the *resolution logic finds it* under the real demo
bring-up. The live re-prove is what caught the gap — exactly why Phase C exists.

## Ledger
- iter-05 (tik): landed F1/F5/F5b (rext r3, RED-fenced) but the live re-prove FALSIFIED F1 — skills still 0 on
  billion, gate unchanged ~3/8; re-diagnosis routed to iter-06 — see iter-05/progress.md

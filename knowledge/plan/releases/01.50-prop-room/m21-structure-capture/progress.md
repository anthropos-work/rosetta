# M21 — Progress

**Status:** in-progress (iter-01 closed). **Shape:** iterative (exit gate in `overview.md`).
**Build with:** `/developer-kit:build-mstone-iters`.
**Active strategy:** TOK-01 (staged-pipeline build toward the binary serve-anonymously gate — see `decisions.md`).
**Furthest pipeline stage passing:** 2 of 6 (static baseline; live baseline in iter-02).

## Running ledger
_Appended after each iter (tik = a standard iter toward the gate; tok = a strategy/retro iter)._

- iter-01 (tok/bootstrap): authored TOK-01 (staged-pipeline strategy) + the 6-stage metric + static baseline
  (stage 2/6); Phase 0b KB-fidelity YELLOW; infra confirmed runnable (Docker + cached directus image + complete
  row cache) — see iter-01/progress.md

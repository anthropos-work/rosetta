# M20 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN. The closing milestone of v1.3b → then `/developer-kit:close-release`._

## Deliverables
- [ ] **Set-dress chaining** — `up-injected.sh` runs `stacksnap replay` → `stackseed` after migrate, reusing the M13 `dev-setdress` pass; default-on + non-fatal; `--no-setdress` escape.
- [ ] **Atomicity contract** — snapshot+seed both-or-neither (no half-set-dressed 403 state); retry-safe via the M17 guards.
- [ ] **Cold-start capture** — the documented DSN-export / restore-a-`pg_dump`-then-`--dsn` workflow; the MCP-adapter spike resolved (build or document-only).
- [ ] **demo auto-set-dress preset** chosen + wired (lean: a demo preset over `dev-min`).
- [ ] **`corpus/ops/snapshot-cold-start.md`** authored; demo recipes + `demo-up`/`demo-down` skills updated.

## Verification
- [ ] A `demo-N` with a warm cache comes up auto-set-dressed (real catalog + a seeded org → login + authorized routes 200).
- [ ] Non-fatal proven: a cold cache (no snapshot) warns + still seeds; no abort.
- [ ] Prod-safety held — capture stays read-only/bounded/firewalled/confirmed; the safety.md drift guards still pass.
- [ ] Go/py/shellcheck clean; flake 0.

## Notes
_(build notes appended here)_

# M20 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN. The closing milestone of v1.3b → then `/developer-kit:close-release`._

## Deliverables
- [x] **Set-dress chaining** — `up-injected.sh` runs the M13 `dev-setdress` pass (now stack-type-aware via `--stack-type demo`) after migrate, before the M18 verify; default-on + non-fatal; `DEMO_NO_SETDRESS=1` escape. ONE engine, two lifecycles. (#M20-D1)
- [x] **Atomicity contract** — seed ALWAYS runs after the cache-first replay (the seed is the floor); a catalog-only stack 403s, a replay miss degrades to structural-only but still logs in 200; retry-safe via the M17 guards. (#M20-D3)
- [x] **Cold-start capture** — net-new `corpus/ops/snapshot-cold-start.md`: the DSN-export / restore-a-`pg_dump`-then-`--dsn` workflow. MCP-adapter spike RESOLVED → **document-only** (the MCP returns JSON rows, not COPY bytes; an adapter would re-serialize COPY format for zero gain). (#M20-D4)
- [x] **demo auto-set-dress preset** — `small-200` (vs dev's `dev-min`), wired as the demo default in the stack-type-aware engine; an explicit `--seed` overrides. (#M20-D2)
- [x] **`corpus/ops/snapshot-cold-start.md`** authored; `safety.md` §2.7 + demo recipes + `demo-up` skill updated; cross-linked from snapshot-spec/db-access/recipe/CLAUDE.md. (`demo-down` needs no change — teardown is unaffected by auto-set-dress.)

## Verification
- [x] A `demo-N` chain wired so a warm cache comes up auto-set-dressed (real catalog + a `small-200` seeded org); the seed-after-replay floor → login 200, not 403. (Engine path proven by the DevSetdress suite against the real script; live demo bring-up = the operator/close-time smoke.)
- [x] Non-fatal proven: a cold cache (replay rc=1) warns + still seeds; behavioral test `test_chain_is_behaviourally_non_fatal_under_set_e` + `test_cache_miss_is_non_fatal_seed_still_runs` (dev+demo).
- [x] Prod-safety held — capture is NEVER on a bring-up (replay-only, pinned by `test_capture_is_never_invoked_on_a_bring_up` + `test_*_replays_both_surfaces_then_seeds`'s `assertNotIn capture`); per-stack Directus firewall-checked; n=0 guard intact; the M15 safety.md drift guards (read+write) stay GREEN after the §2.7 edit.
- [x] Go/py/shellcheck clean; flake 0. (Python 343 pass / 11 live-deselected; Go snapshot+seeding suites green; shellcheck clean on both scripts.)

## Notes
- The "code" is in `rosetta-extensions` (extensions commit `e4d2f9b`, tagged `dress-rehearsal-m20`); the rosetta side is documentation (the cold-start runbook + safety §2.7 + skill/recipe updates). Both committed.
- Tests added: +7 dev-setdress (stack-type/atomicity) + +9 demo chain (7 static fence + 2 behavioral) = +16.

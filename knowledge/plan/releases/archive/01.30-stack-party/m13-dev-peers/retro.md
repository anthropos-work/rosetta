# M13 — Retro

**Milestone:** Dev stacks as full-fidelity peers — local Directus + auto-snapshot + light seed · **Shape:** section · **Closed:** 2026-06-07
**Branch:** `m13/dev-peers` → merged `--no-ff` into `release/01.30-stack-party` · **Tag:** `stack-party-m13` @ `cca4464`

## Summary

The **meatiest milestone of v1.3** and the one that completes the dev/demo convergence thesis: M12 made dev a
peer of demo for **N-allocation**; M13 makes it a peer for **DATA**. A `dev-stack up` bring-up now gets the demo
treatment by default — a new **set-dressing pass** (`dev-stack/dev-setdress.sh`, default-on + non-fatal) stands
up a per-stack Directus, replays the real public taxonomy + content cache-first, and applies the new `dev-min`
light seed so a fresh dev stack is a believable, never-empty world. The load-bearing design move (M13-D2): rather
than re-encode M10's Directus recipe in shell, a small Go runner (`stack-snapshot/cmd/provision-plan`) makes the
**library-only** M10 `ProvisionPlan`/`EnvContract`/`Validate` contract **executable**, so the recipe AND the
prod-Directus firewall have one source of truth shared by both stack kinds. Built in 3 sections (dev-min preset →
the set-dressing pass → corpus), hardened in 1 pass, closed clean. Prod-safety held unchanged throughout: capture
is never run on dev (replay only), media stays refs-only (blob bytes = v1.4), and the n=0-dev guard is now
doubled (the set-dress pass + `stackseed --reset`).

## Incidents This Cycle

**No defects shipped; 3 robustness bugs caught + fixed in harden** (all P2-ish, all in M13's own new code, each
mutation-pinned by a regression test that fails on the pre-fix code):

1. **Whitespace-env firewall slipped a non-prod-safe pass.** `provision-plan --check-env` guarded only the
   exact-`""` case, so a whitespace-only `--base-addr`/`--dsn` slipped past into `Validate` (which rejects an
   empty DSN but not a blank BaseAddr) and printed "prod-safe" for a meaningless address. Fixed to trim +
   fail-closed (the firewall's correct posture).
2. **The set-dress re-run hint always said `--no-snapshot`.** `${no_snapshot:+…}` always expands (the var is
   `0`/`1`, both non-empty), so a manual re-run hint wrongly told a snapshot-ON dev to skip the snapshot. Fixed
   to echo only the flags the bring-up actually used.
3. **A `set -u` trailing-flag leak.** A trailing `--dsn`/`--seed` with no value leaked a raw `$2: unbound
   variable` instead of a clean usage message. Added a `needval` guard.

One *test-quality* note (not a defect): the §2 PR review found `build_cli` resolved the cmd package one dir too
high (the stub tests short-circuit the real build path, so it hid until a real-build test was added). Fixed to
build `./cmd/<x>` from the module root + a go-guarded real-build regression test (M13-D4).

No flakes (5/5 at close — Go shuffled + dev-stack sequential; 0 during harden). No regressions in the reused M10
surface (its `directus` package stays 100%, byte-for-byte). No P1s.

## What Went Well

- **Reuse-the-contract, don't re-encode-it (M13-D2).** Making the M10 library contract executable via a tiny
  runner gave dev + demo one source of truth for the per-stack Directus recipe + the firewall. This is the
  cleanest possible "dev gets the demo treatment" — no forked shell recipe to drift.
- **The replay machinery needed zero change.** `stacksnap replay` already accepted `dev-N` (`pg.ParseStackN`)
  and already derived its write-target from the stack offset (`pg.DSNForOffset`), so "dev as a replay target"
  was pure bring-up wiring — the KB-fidelity audit (GREEN) confirmed this before any code was written.
- **All 3 open questions resolved to decisions during build** (Q1 heaviness → cache-first + non-fatal; Q2
  default-on → `--no-snapshot`/`--no-setdress`; Q3 dev-min size → 10-user role-mix floor). No thrash.
- **Prod-safety stayed invariant.** Every guard from v1.1/v1.2 held: read-only capture (never on dev), refs-only
  media, `AssertPublicOnly`, the 3-layer isolation guard, and now a doubled n=0-dev guard.

## What Didn't (smaller stuff)

- **A per-unit handbook-index miss escaped build + harden** — the new `cmd/provision-plan` command shipped
  without a row in the stack-snapshot README Packages table; caught at close (Phase 3) and fixed. Same class of
  miss M10 had (the missing taxonomy/directus README rows). The dev-stack README + both corpus docs were
  complete; only the *origin module's* index lagged. Carrying-forward note below.
- **The harden bugs were all in the new shell + the new runner's edge paths** — the firewall green-light and the
  re-run hint were both "looks right, expands wrong" bash/Go-zero-value traps. Adversarial-style edge tests at
  build (not just harden) would have caught them one phase earlier.

## Carried Forward

- **None blocking.** No new deferrals (M13 added zero — all scope landed Fate 1).
- **DEF-M10-01** (S3 media blob bytes + cloud SnapshotStore backend) → **v1.4**, inherited from v1.2, signed by
  the user at v1.3 design (2026-06-07), parked in `roadmap-vision.md § v1.4 seeds`. Unchanged this milestone;
  M13 reused the refs-only media path verbatim. Re-confirmed in this close's deferral audit (GREEN).
- **Process nudge for M14/M15** (not a tracked item): when a milestone adds a new top-level command/unit, add its
  component-index row in the SAME commit as the code — the per-unit handbook-index check keeps catching this at
  close. Cheap to front-load.

## Metrics Delta (from `metrics.json`)

- **Go test funcs:** 708 → **720** (+12: stack-snapshot 212→223 the provision-plan runner; stack-seeding 232→233 the dev-min preset pin). Both modules `-race -count=1` green; gofmt + `go vet` clean.
- **Python (dev-stack pytest):** 33 → **38** (+5: the set-dress wiring + dev-setdress edge/regression tests). All other Python suites green (stack-core+demo-stack 67, injection+verify 69).
- **Coverage (M13-touched):** `cmd/provision-plan` 93% (90.7%→93% at harden; residual = the `os.Exit` shim + an invariant-pinned defensive branch); `directus` + `blueprint` 100%.
- **Flake:** **0** (5/5 — Go shuffled + dev-stack sequential). shellcheck + py_compile CLEAN.
- **Review:** 6 findings — 1 doc must-fix (the README index row), 5 GREEN/no-action (4 adversarial scenarios + 3 decision-triage blends, ref-tagged). Deferral re-audit GREEN (0 new). 0 escape-hatch.

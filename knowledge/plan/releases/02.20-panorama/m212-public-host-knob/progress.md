# M212 — progress

Section checklist (closure = all boxes land + a dry `up-injected.sh` run with `STACK_PUBLIC_HOST` set bakes the
MagicDNS host into every browser-facing value; unset ⇒ byte-identical to today).

- [x] `HOST` var + `STACK_PUBLIC_HOST` default in `up-injected.sh` (+ `FAPI_HOST` dotted default, `BIND_HOST` gate — D-IMPL-1/2)
- [x] `/demo-up --public-host` flag → `STACK_PUBLIC_HOST` plumbed to scripts (up-injected.sh flag + env-var; SKILL doc)
- [x] next-web build-args + `.env.local` overlays substituted
- [x] studio-desk build-args substituted
- [x] `inject.py --fapi-host "$FAPI_HOST:..."` (pk mint — inject.py already host-parametric; caller + MagicDNS round-trip test)
- [x] `gen_injected_override.py` host-param plumbing (emission → M214; wired-but-unemitted seam — D-IMPL-3)
- [x] cockpit `--host 0.0.0.0` (opt-in) + host into `--app-base`/`--fapi-host`/`--academy-base` (cockpit.py needed no change)
- [x] `ant-academy.sh` host sub + gated `-H 0.0.0.0` bind
- [x] `demo_web` Directus rewrite substituted
- [x] `want_ep` cache-validators invalidate on HOST change
- [x] `stack_registry.py` additive `external_host` (+ `set_host` reconcile-upsert, CLI, `rosetta-demo status` render — D-IMPL-4)
- [x] unset-knob regression check (byte-identical) + tests

## Closure (2026-07-11)
**DONE.** All 12 sections landed. Code in rext @ tag **`panorama-m212`** (sha `d4f6da6`), 3 commits (stack-core →
stack-injection → demo-stack, built bottom-up). Zero platform-repo edits. Tests: stack-core **95**, stack-injection
**123** (8 skipped), demo-stack **340** (+ the live-docker `test_migrate_race_live` not run) — all green; both scripts
shellcheck-clean. Closure contract met: a dry `up-injected.sh` bakes the MagicDNS host into every browser-facing value
when `STACK_PUBLIC_HOST` is set, and is byte-identical to today when unset (pinned by `TestHostKnobClosure` +
`test_host_knob_derivation_*`). KB-fidelity Phase 0b = YELLOW (KB-1 homed to M214). Decisions D-IMPL-1..4 recorded.

## M212: Hardening

Passes: 2 (stopped — remaining coverage misses are module-entry guards / argparse-unreachable only; full
dimension scan clean; no flakes). Code lives in **rext** (`.agentspace/rosetta-extensions/`, branch `main`);
tag `panorama-m212` re-pointed `d4f6da6` → post-harden HEAD.

**Scope manifest (milestone-touched, `d0cdfbb..HEAD` in rext):**
- `stack-core/stack_registry.py` (allocate/set_ports `external_host` + `set_host` reconcile-upsert) — tests `stack-core/tests/test_stack_registry.py`
- `stack-injection/inject.py` (`--fapi-host` help / pk-mint) — tests `stack-injection/tests/test_injection.py`
- `stack-injection/gen_injected_override.py` (the `host` seam through `build_lines`/`frontend_lines` + `--public-host`) — same test file
- `demo-stack/up-injected.sh` (HOST/FAPI_HOST/BIND_HOST derivation, `--public-host` flag, want_ep validator, demo_web rewrite, set-host call) — tests `demo-stack/tests/test_frontend_build.py`
- `demo-stack/ant-academy.sh` (`$HOST` STUDIO_URL + gated `-H 0.0.0.0`) — tests `demo-stack/tests/test_ant_academy.py`
- `demo-stack/rosetta-demo` (status render) — covered via registry/status tests

### Pass 1 — 2026-07-11
**Coverage delta (milestone-touched Python):**
- `inject.py`: 33% → 98% statements (only the `__main__` guard line remains)
- `gen_injected_override.py` / `stack_registry.py`: already 99% — held.

**Tests added:**
- `test_injection.py` (`TestInjectMainArtifacts`): 7 — drove `inject.main()` end-to-end (0%-covered by build): four-artifact write for the byte-identical default (127.0.0.1) + the M212 MagicDNS host, the round-trip self-check, custom bapi-ip/webhook-secret, append-not-truncate on a pre-seeded `.env`, the `$`-sentinel error path, the D-IMPL-1 boundary (inject.py mints ANY non-`$` host — dotted enforcement is `@clerk/backend`'s downstream), and a base64 padding-residue grid.
- `test_frontend_build.py` (in `TestFrontendBuildBehaviour`): 3 — `want_ep` cache-validator **HOST-invalidation** (the overview's "top-risk item"): next-web rebuild-on-host-change + reuse-on-host-match, studio-desk rebuild-on-host-change. The build phase only varied OFFSET; a future edit could hardcode `localhost` back into `want_ep` (offset tests still pass) and silently reuse a stale localhost image on a `--public-host` stack — now pinned.

**Bugs fixed inline:** none — the M212 production code was solid; the gaps were untested-surface, not defects.
**Flakes stabilized:** none observed.
**Knowledge backfill:** no KB-worthy findings — the discoveries (inject-mints-any-host boundary; want_ep-embeds-$HOST) are the D-IMPL-1/§Scope facts already recorded in `decisions.md`/`overview.md`; the clerkenstein dotted-host doc gap stays homed to M214 (KB-1). Nothing net-new to blend.

### Pass 2 — 2026-07-11
**Coverage delta:** `stack_registry.py` held at 99% (the additions exercise already-covered branches for behavioral edges, not new lines). Shell flag-parse is behavioral (no line coverage).

**Tests added:**
- `test_frontend_build.py` (in `TestFrontendBuildBehaviour`): 4 — EXERCISE the `--public-host` operator-surface **flag** (build phase tested only the env-var path): flag derives HOST/FAPI_HOST/BIND_HOST identically and wins over an empty env; bare `N` is a byte-identical no-op; `--public-host` with no value fails loudly (the `${2:?}` guard); an unknown argument is rejected (the `*)` arm). Extended the lib-only harness with a `source_args` seam.
- `test_stack_registry.py` (`TestExternalHost`): 2 — `set_host` additive-preservation (stamps `external_host` without clobbering existing ports/created) + idempotent re-stamp (a re-run at a new MagicDNS name updates in place, no duplicate row, no stale value).

**Bugs fixed inline:** none.
**Flakes stabilized:** none.
**Knowledge backfill:** no KB-worthy findings (operator-surface flag behavior + registry additivity are already in `overview.md` §Scope / `decisions.md` D-IMPL-4).

### Stop condition
Stopped after Pass 2: the remaining line-coverage misses are all module-entry `__main__` guards + one argparse-unreachable `return 1` (subparsers are `required=True`), which carry no behavioral value; the Step 2b dimension scan surfaced nothing new; no flaky tests across the verification runs. Post-harden rext totals: stack-core **97**, stack-injection **130** (8 skipped), demo-stack **350** — all green; both scripts shellcheck-clean.

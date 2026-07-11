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

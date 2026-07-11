# M212 — retro (the single host knob)

## Summary
The v2.2 "panorama" release-opening milestone. Introduced ONE opt-in browser-facing host knob — `STACK_PUBLIC_HOST`
(default `localhost`, so an unset stack is byte-identical to today) — surfaced as the `/demo-up --public-host` flag
and threaded through every rext emitter that bakes a browser-facing `localhost`/`127.0.0.1`: next-web + studio-desk
build-args and `.env.local` overlays, the pk/FAPI host, the cockpit + ant-academy bases + their opt-in `0.0.0.0`
bind, the `demo_web` Directus content-URL rewrite, the `want_ep` cache-validators, and an additive `external_host`
on the stack registry. All 12 sections landed; zero platform-repo edits; zero Go touched. Closed clean on the first
close pass — merged `--no-ff` into `release/02.20-panorama`.

## Incidents This Cycle
None. No P1/P2, no flakes (gate 5/5 across all 3 touched suites), no regressions. Build + 2 harden passes + 4 close
adversarial scenarios surfaced zero defects in the production code.

## What Went Well
- **The byte-identity contract paid for itself.** Designing the knob as "unset ⇒ exactly today" (D-DESIGN-1) made
  the whole review cheap — every emitter change is provably a no-op when the flag is absent, pinned by
  `TestHostKnobClosure` + `test_host_knob_derivation_*`, so the risk surface collapsed to "the opt-in path only".
- **The two design-time forks were the right calls.** The separate dotted `FAPI_HOST` default (`127.0.0.1`, D-IMPL-1)
  kept the pk valid AND byte-identical; gating the `0.0.0.0` bind on the knob (D-IMPL-2) kept external listeners
  opt-in. Both were caught at build, not close.
- **The top-risk item was pre-empted.** The `want_ep` cache-validator HOST-invalidation (a stale localhost-baked
  image silently reused on a `--public-host` stack) was flagged in the overview and pinned by harden Pass 1 — the
  close adversarial review found it already handled.
- **Clean seam discipline for M214.** The `gen_injected_override` host param is wired end-to-end but deliberately
  un-emitted (D-IMPL-3), pinned byte-identical — M214 flips only the terminal emission, no rework.

## What Didn't
- **A pre-existing handbook count-drift surfaced at close.** `demo-stack/README.md:66` quotes `test_tooling.py
  (50 tests)`; actual 111 — drift accumulated across prior releases, not an M212 regression. Because the rext
  code-of-record is FROZEN at tag `panorama-m212` (and this rosetta-only close must not advance/re-tag rext), it
  could not be fixed in-place. Routed rather than papered over.
- **Minor tooling friction (non-blocking):** the box's default `python3` (3.14) has a broken `pyexpat`; pytest +
  the JUnit tally had to run under python 3.12. No milestone impact; noted for the next close.

## Carried Forward
- **D-CLOSE-1** — `demo-stack/README.md` `test_tooling` count 50→111 → **v2.2 `/developer-kit:close-release`** (which
  legitimately re-tags/advances rext + runs release-level doc hygiene). Fate 2, in-release.
- **KB-1 / D-IMPL-1** — the clerkenstein.md dotted-FAPI-host constraint → **M214** (`tailscale-serve.md` +
  `clerkenstein.md` update). Fate 2, confirmed-owned.
- **D-IMPL-3** — the CORS + Clerk sign-in/web-app URL emission at the public host → **M214**. Fate 2, confirmed-owned.

## Metrics Delta
- **Tests:** 577 collected across the 3 M212-touched rext suites (stack-core 97 · stack-injection 122+8skip ·
  demo-stack 350) — 569 pass / 8 skip / 0 fail (JUnit-authoritative). +16 net-new from the 2 harden passes.
- **Coverage:** inject.py 33%→98%; gen_injected_override.py + stack_registry.py held at 99%.
- **Flake:** 0 (5/5 gate). **Go:** 0 touched. **Shellcheck:** clean (both scripts). **Platform-repo edits:** 0.
- Full machine-readable record: `metrics.json`.

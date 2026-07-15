---
title: "KB Fidelity Audit — M221 prove-on-billion"
date: 2026-07-15
scope: milestone:M221
invoked-by: build-mstone-iters (Phase 0b, pre-bootstrap-tok gate)
---

## Verdict
**YELLOW** — no blind areas, no surviving stale load-bearing claim. One genuine stale-exposure
finding in a named KB-dependency doc (`tailscale-serve.md`) was **FIXED INLINE** (Phase 6); one
incidental completeness gap noted. All of the parent's hard spot-checks PASS. **Does not block the
bootstrap tok.**

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Remote-access runbook (`--public-host`, F1–F12) | `corpus/ops/demo/tailscale-serve.md` | `demo-stack/{up-injected.sh,ant-academy.sh,cockpit.py,reap.sh,tailscale_autohost.py}`, `stack-injection/{gen_injected_override.py,exposure_claim_guard.py}` | PAIRED |
| Bring-up auto-verify (iteration protocol) | `corpus/ops/verification.md` | `stack-verify/live/autoverify.sh` | PAIRED |
| Believability / coverage gate | `corpus/ops/demo/coverage-protocol.md` | `stack-verify/e2e/run-coverage.sh` | PAIRED |
| Functional-flow Playthroughs | `corpus/ops/demo/playthroughs.md` | `playthroughs/e2e/` | PAIRED |
| Click→ACCESS latency budget (<5 s gate) | `corpus/ops/demo/latency-budget.md` | `stack-verify/e2e/run-latency.sh` | PAIRED |
| Exposure contract (default-on remote flip) | `corpus/ops/safety.md` Part 3 | `stack-injection/exposure_claim_guard.py`, `demo-stack/up-injected.sh` | PAIRED |

All six named KB dependencies exist and pair with live code. No BLIND-AREA.

## Fidelity Findings

### F1 — tailscale-serve.md denies the academy's live LAN exposure (STALE → FIXED INLINE) — **the parent's flagged item**
- **Source:** `corpus/ops/demo/tailscale-serve.md` lines 462 (surface table), 464 (bind table), ~626 (safety framing).
- **Expected (doc):** the presenter cockpit **and** ant-academy bind **loopback by default** — "It is `""` (loopback) by default and `0.0.0.0` under a public host."
- **Actual (code):** `ant-academy.sh:296` — `bind_args=(); [ -n "${STACK_PUBLIC_HOST:-}" ] && bind_args=(-H 0.0.0.0)`. With no public host, **no `-H` is passed and `next dev`'s own default is `0.0.0.0`** → the academy binds **`*:13077`** on every demo. Only the **cockpit** is genuinely loopback by default (`cockpit.py:572` — `--host` default `127.0.0.1`, and `up-injected.sh:1713` omits `--host` when `BIND_HOST=""`).
- **Corroboration:** `safety.md` Part 3 (lines 432–437, M220) states it **correctly** — table row `ant-academy (3077+off) | *:13077 | YES — HTTP 200 ❌`, and "the *'gated on the knob'* framing is **only true of the cockpit**." The two docs contradicted each other; the milestone's own `overview.md` (`FIX-M221-academy-loopback-bind`) also has the correct contract.
- **Why it slipped:** M220's correction fixed the **container** exposure claim (§3.1 retraction box) but left the **host-native academy** claim denying the exposure. The fence meant to catch this (`exposure_claim_guard.py`, "fails if any doc denies it") is itself **container-only / blind to host-native listeners** (lines 141–143 check `directus_lines`/`frontend_lines`/`build_lines` only) — the exact F-M220-5 gap. So the doc-lie survived S0/S1 exactly as F-M220-5 predicts.
- **Verdict:** STALE → **FIXED INLINE** (Phase 6). Lines 462, 464, and the safety-framing bullet now split the cockpit (`127.0.0.1` loopback default) from ant-academy (`*:13077`, `next dev` default `0.0.0.0`, world-published on every demo), matching `safety.md` Part 3, and point at `FIX-M221-academy-loopback-bind`.
- **Note on the overview's carry:** the parenthetical "**The docs were corrected in M220 (Fate 1)**" was **imprecise** (only partially corrected) until this fix; now accurate. The load-bearing content of F-M220-5 (academy binds 0.0.0.0, code fix + guard extension pending) was already correct.

### F2 — `aiReadinessStep1Score` double-round fix (c6648d1) — ALIGNED
- **Source:** overview `REPROVE-M221-battery-at-final-code`.
- **Actual:** `stack-seeding/seeders/ai_readiness_funnel.go:383–394` rounds **once** — `v := int(heldWeight/totalWeight*float64(aiReadinessStep1Max) + 0.5)`. The double-round intermediate survives only in the explanatory comment. Divergence values verified against the platform's single-round: held 2.5→**12**, 4.0→**18**, 5.5→**25** (the seeder now agrees with the platform where it previously reported 11/19/26).
- **Verdict:** ALIGNED. The fix is in the pinned `cue-to-cue-m220-r6`.

### F3 — `NEXT_PUBLIC_BACKEND_API_URL` blackholed twin (F-7) is genuinely DORMANT — ALIGNED
- **Source:** overview `PROBE-M218-backend-api-url-twin`.
- **Actual:** `up-injected.sh:693` bakes `NEXT_PUBLIC_BACKEND_API_URL=$SCHEME://$HOST:$((8082+OFFSET))` → `https://billion…:18082` under a public host (confirmed). All **17 readers** in `next-web-app` are client-side: the two app-router `page.tsx` files that would default to server components (`invite/[token]`, `reminders/[token]/unsubscribe`) both carry `'use client'`; the rest are hooks (`use*`), UI components, and browser HTTP helpers. No route handler / middleware / RSC reader. M218 D10's "every current reader is client-side" holds → the gun is dormant, as the carry claims.
- **Verdict:** ALIGNED. (The carry's DoD — fence a future server-side reader — remains genuine M221 work.)

### F4 — M217 pre-bind-reap fix landed but unproven-in-field — ALIGNED
- **Actual:** `up-injected.sh:800` now `. "$HERE/reap.sh"` (the missing source that made `reap_port` dead code exit 127). `assert_ports_free` defined at `reap.sh:250`. Both match the carry ("fixed and unit-proven; not field-proven").
- **Verdict:** ALIGNED.

### F5 — `REPROVE-battery-at-final-code` commit ordering — ALIGNED
- **Actual:** `c6648d1` (double-round) and `b5bf65b` (coverage stale-report) are both **NOT** ancestors of `cue-to-cue-m219-r8` (the grading tag) → they landed after it, exactly as the carry states; both **are** in the pinned `cue-to-cue-m220-r6` (the code M221 runs). The seed-path change (`c6648d1`) genuinely restarts the battery count per M218 D13.
- **Verdict:** ALIGNED.

### F6 — `FIX-M221-devstack-test-spin` DISCHARGED — VERIFIED (not on faith)
- **Actual:** ran the full suite — `pytest .agentspace/rosetta-extensions/dev-stack/tests/` → **116 passed, 4 skipped in 129.68s**. No spin (completed well under the 280 s timeout). Matches the overview's "116 passed · 4 skipped · 127 s." Both root-cause fixes present: `DEV_SETDRESS_USE_STUB_BINS=1` honored (`dev-setdress.sh:116`) and set by the harnesses (`test_dev_stack.py:263, 621`, the M220-close F-M220-1b fix at :605); the secret-preflight wiring (D29) is present and non-fatal.
- **Verdict:** correctly DISCHARGED.

## Completeness Gaps

1. **(incidental)** `playthroughs.md` does not state "**perf is a NON-GOAL**" explicitly — the overview's "Also lands" justification for a separate `stack-verify` latency surface paraphrases principle **P2** ("functional truth, not pixel truth", line 152). The decision is sound and concretely implemented (`stack-verify/e2e/run-latency.sh` exists, M218), so this is a wording gap, not a stale claim. Optional 1-line note; not applied (pre-iter, low value, avoids over-editing a stable doc).

## Applied Fixes
- `corpus/ops/demo/tailscale-serve.md` — three inline corrections (lines 462, 464, safety-framing bullet) so the doc stops denying the academy's live LAN exposure and matches `safety.md` Part 3 + `FIX-M221-academy-loopback-bind`. (No frontmatter to bump — the doc uses a heading-style intro, not YAML.)

## Open Items (require user decision)
_None._ The one stale finding was small and unambiguous and is fixed inline. The exposure-guard host-native blindness that let it slip is already tracked in-scope as `FIX-M221-academy-loopback-bind` ("Extend `exposure_claim_guard` to the host-native listeners"), not a new item.

## Gate Result
**YELLOW: proceed.** No blind area; no surviving stale load-bearing claim (the one found is fixed);
every inherited-carry spot-check the bootstrap tok relies on (double-round fix, academy bind,
backend-api-url dormancy, reap fix, battery commit-ordering) is ALIGNED with code; the DISCHARGED
`devstack-test-spin` is verified by an actual green run, not taken on faith. The bootstrap tok may
author its strategy against these docs. Eyes-open note for the tok: the `exposure_claim_guard` fence
remains blind to host-native listeners — that blindness is `FIX-M221-academy-loopback-bind`'s own DoD.

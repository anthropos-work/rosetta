---
title: "KB Fidelity Audit — M237 clean stage"
date: 2026-07-21
scope: milestone:M237
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Clone-freshness mechanism | `corpus/ops/rosetta_demo.md` §"Clone freshness — skip-if-present never updates (`F-M236-CLOSE-1`)" | `demo-stack/ensure-clones.sh` (phase c `make init`, phase e provenance), `platform/Makefile` (`init`/`pull`/`status`) | PAIRED |
| R1 pristine sweep (all manifests) | `corpus/ops/demo/demopatch-spec.md` §2.1 + G5 row | `demo-stack/ensure-clones.sh` `PATCH_MANIFESTS` (phase f), `demo-stack/patches/` (14 dirs) | PAIRED |
| Fresh-clone billion re-triage methodology | `corpus/ops/verification.md` §"PRE-FLIGHT RUNG ZERO" + §"Drive every remote bring-up through a LOGIN shell"; `corpus/ops/demo/tailscale-serve.md`; billion reach (MEMORY) | `demo-stack/up-injected.sh`, tailscale-serve tooling | PAIRED |

## Fidelity Findings
1. **Clone-freshness doc vs code — ALIGNED.** `rosetta_demo.md` §Clone-freshness states truthfully: `make init` is a pure existence check (matches `platform/Makefile:19` skip-if-present); `clones.lock.json` records `{ref,sha}` as "provenance of what was built, NOT a freshness signal" and `ref` degrades to `"HEAD"` on detached (matches `ensure-clones.sh` phase e); "There is no automatic freshness rung today ... `F-M236-CLOSE-1` is the finding that owns closing that gap. The gap to close is **visibility**, not automatic mutation." The doc scopes exactly M237 Section 1. Verdict: ALIGNED (describes current pre-fix code; is the Delivers target to flip in Phase 5).
2. **R1 sweep doc vs code — ALIGNED.** `demopatch-spec.md` §2.1 states R1 iterates a hardcoded 3-entry `PATCH_MANIFESTS` (quoted block matches `ensure-clones.sh:215-219` byte-for-byte); "There are **14 manifests under `demo-stack/patches/`**" (confirmed: 14 `*.yaml` dirs); "The structural fix ... is `F-M236-CLOSE-2` and is **not** in v2.5." The G5 row (line 52) also says "covers 3 of the 14 manifests." Verdict: ALIGNED (describes current bug; is the Delivers target to flip in Phase 5 — BOTH the §2.1 body and the G5 row must be updated).
3. **Billion login-shell landmine — ALIGNED.** `verification.md` §"Drive every remote bring-up through a LOGIN shell" documents `ssh <host> 'bash -lc "…"'` and the `/usr/local/go/bin` profile-only PATH trap that mimics a missing prereq. Matches the orchestrator's methodology note. Verdict: ALIGNED.

## Completeness Gaps
- None load-bearing. The new `DEMO_ADVANCE_CLONES` env knob (Section 1) is NEW work this milestone creates, not a pre-existing blind area; it will be documented in `corpus/ops/demo/demo-up-defaults.md` (a `demo_knob_guard.py` requirement) + `rosetta_demo.md` inline with the implementation.

## Applied Fixes
- None needed. All three topics PAIRED + ALIGNED. spec-notes.md triples recorded.

## Open Items (require user decision)
- None.

## Gate Result
GREEN: proceed to build-milestone Phase 1. The two Delivers docs (`rosetta_demo.md`, `demopatch-spec.md`) accurately describe the current pre-fix state and are the exact targets to update in Phase 5 once the code lands.

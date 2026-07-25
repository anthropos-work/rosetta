---
title: "KB Fidelity Audit — M251 test-health"
date: 2026-07-23
scope: milestone:M251
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| run-unit unit-spec roster / `UnitSpecsAreExecuted` guard | `corpus/ops/demo/coverage-protocol.md` (unit-spec sweep, `content-denominator.json expected_pairs=49` @ :917) | `rext stack-verify/e2e/run-unit.sh`, `rext stack-verify/tests/test_e2e_collection_integrity.py` | PAIRED |
| cockpit overlay in-flight window (M218 removal) | `corpus/ops/demo/cockpit-spec.md` :271–277 (documents the 30s→removed change, "Fixed cue-to-cue-v2.3.1"); `corpus/ops/demo/latency-budget.md` | `rext demo-stack/cockpit.py::_OVERLAY_JS`, `rext demo-stack/tests/test_cockpit.py::TestOverlayJs` | PAIRED |
| cockpit per-hero [Academy] link (removed 2026-07-15) | `corpus/ops/demo/cockpit-spec.md` (no render claim); `corpus/services/ant-academy.md` | `rext demo-stack/cockpit.py:858-862` (removal comment), `tests/test_cockpit.py::TestAcademyLink/TestAcademyCatalogEntryEdges/TestServedPanelWithAcademy` | PAIRED |
| public-host per-port serve/reset fronting (M226/M220 hiring 13001) | `corpus/ops/demo/tailscale-serve.md`, `corpus/ops/safety.md` §3 | `rext stack-injection/gen_tailscale_serve.py` (UI_BROWSER_FACING hiring 3001 @ :44), `tests/test_host_prereqs_m215.py::TestF12ServeResetGenerator` | PAIRED |
| demo-stack python suite index anchor | `corpus/ops/verification.md` (anchor does NOT yet exist) | `rext demo-stack/tests/*.py` | DOC-ONLY (optional deliverable — DEFERRED to M247, see decisions.md) |

## Fidelity Findings
No STALE knowledge-doc claims found. The milestone's premise is that the *tests* (rext code) are stale
against deliberately-changed *behaviour*; the knowledge docs already describe the current truthful behaviour:

1. **Overlay 30s window** — cockpit-spec.md:271–277 explicitly narrates the 30s in-flight window as
   **removed** ("any return to the cockpit clears the flag"). ALIGNED with `_OVERLAY_JS` (no `30000`,
   `getItem` read-path gone, `removeItem` clear + `pageshow` reset present). The *test* lags, not the doc.
2. **Per-hero academy link** — no corpus doc claims the cockpit renders a per-hero `class="btn academy"`
   link (grep: 0 hits). cockpit.py:858-862 documents the removal in-code. ALIGNED.
3. **Hiring port 13001 fronting** — gen_tailscale_serve.py:44 + tailscale-serve.md document hiring joining
   the UI-tier serve/reset fronting (M226). ALIGNED. The *test*'s `_UI_PORTS` set lags (missing 3001).

## Completeness Gaps
None load-bearing for M251. The optional `corpus/ops/verification.md` demo-stack-suite index anchor
(overview.md "Delivers →") is intentionally NOT authored in this lane — it is DEFERRED to M247 to avoid a
cross-lane collision on the corpus doc M247 owns (recorded in decisions.md). Not a blind area: the code the
anchor would index already exists and is exercised.

## Applied Fixes
None needed — no stale knowledge-doc claims to correct. (The stale assertions live in rext *test code*,
which is M251's implementation subject, not a knowledge doc.)

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to build-milestone Phase 1. Every topic PAIRED (or an intentionally-deferred optional
doc-only deliverable), every knowledge-doc claim ALIGNED with current code, no blind areas, no critical
undocumented behaviours in the code M251 modifies.

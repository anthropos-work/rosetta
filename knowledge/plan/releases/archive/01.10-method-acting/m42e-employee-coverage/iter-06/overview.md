---
iter: 06
milestone: M42e
iteration_type: tik
iter_shape: tooling
status: closed-fixed
created: 2026-06-25
---

# iter-06 — tooling-iter: the `/` sentinel false-positive (innerText, not textContent)

**Type:** tik / **iter_shape: tooling** (coverage-protocol.md "Iter type selection -> Tooling-iter"). Fixes the
last harness false-positive in the residual: the root `/` is flagged `error` because the assertion read the
whole `<body>` (incl. an inlined i18n JSON blob whose translation VALUES include "Something went wrong"). Ships
the assertion fix + re-sweeps. This is tik #5 of the session (the 5-tik cap fires after it closes).

## Active strategy reference
**TOK-01: sweep-then-route-by-leverage.** The `/` false-positive is a deterministic 1-failure clear via a
1-line assertion change (`textContent` -> `innerText`) -- the highest-confidence residual move; the 2 empty
skill-paths are a deeper content gap (next session).

## Re-survey (Phase 1 Step 0)
iter-05 re-sweep is the current state `(failing=3, escapes=2)`. The `/` sentinel false-positive is still
present + confirmed (iter-03 D3 / iter-05 D4a: `hasMain=false`, the sentinel matched the `errorSettingUpGPTRealtime`
i18n string in the serialized body). The assertion fix (a WIP `innerText` change in `crawl.ts`) is uncommitted
in the rext tree -- adopt + commit it as this iter's deliverable.

## Cluster / target identified
The per-page non-emptiness assertion in `crawl.ts::assertPage`. The fix: read **`innerText()`** (visible text)
instead of **`textContent()`** (which serializes hidden + inlined content, incl. Next.js's inlined i18n
translation table). The i18n table carries values like "Something went wrong" that false-match the error
sentinel on a fully-rendered page; `innerText` structurally excludes it.

## Hypothesis
Switching the text read to `innerText` makes the `/` root read its REAL visible text (not the 189 KB inlined
JSON), so the error-sentinel no longer false-matches -> `/` stops being flagged `error`. `failing` 3 -> 2 (the
2 empty skill-paths remain). escapes unchanged (2).

## Expected lift
`failing` 3 -> 2. Risk: `innerText` could change other pages' textLen (it returns only visible text, which can
be SHORTER than textContent) -- the re-sweep verifies no previously-ok page drops below the textFloor.

## Phase plan (multi-step tooling-iter -- planned 2 lines)
- **Line 1 (ship):** adopt the `crawl.ts` `textContent` -> `innerText` change; compile-validate.
- **Line 2 (use):** Phase D re-sweep vs live demo-3; record `(failing, escapes)` + confirm `/` cleared + no new
  empties from the visible-text floor.
- **Phase E -- close + cap:** grade; this is tik #5 -> the 5-tik cap fires; emit the session report.
- Fold the innerText-vs-textContent rule into coverage-protocol.md (protocol-evolution; supersedes the prior
  "prefer <main>" note's caveat).

## Escalation conditions
- `innerText` drops a legitimate page below the textFloor (a new false-empty) -> tighten the floor or escalate
  to a Tier-2 selector for that page; route forward. Not a blocker.

## Acceptable close-no-lift outcomes
- If `innerText` clears `/` but introduces a new false-empty elsewhere (net failing unchanged), close on the
  landed assertion capability + route the new false-empty; the assertion still improved (fewer false-positives).

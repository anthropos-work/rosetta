**Type:** tik (run 6, tik 1). Active strategy: TOK-02.

# iter-13 — progress

## Probe (TOK-02 mandate: confirm the slow-fetch cause FIRST)
Static read of `next-web-app/packages/ui/src/AISimulation/AISimulationResultContainer.tsx` (stack-dev checkout): the two extraction fetches (`useGetSimulationExtractionSchema` / `useGetInterviewExtractionReport`, gated by `isExtractionEnabled`, lines 508-517) feed ONLY the `interviewExtractionPlan` / `interviewExtractionData` PROPS passed to `<AISimulationResult>` — they are **NOT** in the `BaseLoading` gate (lines 551-565: `sessionLoading || userDataLoading || resultLoading || membershipLoading || interactionsLoading || callInteractionsLoading || jobSimDataLoading`). So scoping the container FETCH to manager cannot change what gates the player ack. Strong static evidence the fetch-cause is wrong.

Live network+DOM probe on billion (both interview player seats, ephemeral spec, deleted after use) CONFIRMED it:
- `intv-voice-pass/player`: ack first seen **5319 ms**; body plateaus at **128** chars (nav shell) from ~2.3 s to ~4.3 s, then ack paints at ~5.3 s (body→334, main→205).
- `intv-voice-fail/player`: ack **5419 ms**; body plateaus at **126** until ~4.4 s, ack at ~5.4 s (body→332).
- The extraction legs (`interviewExtractionUserReport`, `simulationExtractionSchema`) DID fire and returned **200** — irrelevant to the ack timing. (iter-11's "~25 s" was a cold server; warm it is ~5.4 s.)

**Root cause:** the nav shell alone (126-128 chars) sits ABOVE `MIN_SETTLED_CHARS` (120), so `settle()`'s length-stability early-exit declares the chrome plateau "settled" (~2.3 s) BEFORE the ack paints (~5.4 s); `assert()` then reads no ack. The warm retry re-incurs the identical plateau, so it never helps. → **TOK-02's FALLBACK is correct**; the preferred demopatch path is falsified.

## Fix (fallback: shape-aware settle — harness-side, NOT a demopatch → NO re-bake)
`stack-verify/e2e/lib/content-result-page.ts`:
- Hoisted the EN/IT ack regex to a module const `INTERVIEW_ACK_RE` (shared by `settle()` + `assert()`).
- `settle(shape?)` now consults a `contentReady(shape)` guard: for `player-interview` it early-exits only once the acknowledgement is actually present in `<body>`; for every other shape (and shapeless callers) `contentReady` is always true → behaviour unchanged.
- `assert()` player-interview uses the hoisted const.
`stack-verify/e2e/tests/content-stories.spec.ts`: both `settle()` calls now pass `shape`.

Because the content-stories harness runs LOCALLY (driving billion's UNMODIFIED app over the tailnet), the fix needs **no billion re-pin and no image re-bake** — a large deviation from TOK-02's cost estimate, justified by the probe.

## Tests
`content-result-page.unit.spec.ts`: +4 virtual-clock staged-page specs (new `FakeStagedPage` renders in stages):
1. ← REGRESSION: does NOT early-exit on the 128-char chrome plateau; waits for the late ack + the pair then LANDS.
2. already-acked → settles promptly (fast path).
3. ack-never-paints → polls to deadline, then asserted false (no hang, no false-land).
4. blast-radius: `player-scored` settles on stable length, ack-blind (guard is player-interview-only).
**Mutation-verified teeth:** removing `&& (await this.contentReady(shape))` fails 2 of the 4. 48/48 unit green; `tsc --noEmit` clean.

## Live re-sweep (billion, foreground 6.0m)
`run-content-stories.sh 1 --host billion.taildc510.ts.net` → **LANDED 47 / 47** (simulation 44/44, skill-path-legacy 2/2, skill-path-new 1/1); runner exit 0. Both interview player pairs: `ok=True shape=player-interview mainLen=205 reason="interview acknowledgement rendered"`.

## rext shipped
Commit `2bb0473` on `main` (pushed). Consumption tag `sound-check-m244-content-sweep-robustness` MOVED to it (annotated, peels `^{}` → 2bb0473 on origin — rung-zero verified). Billion's build-time pin is NOT touched (harness is not a build input).

## Close — 2026-07-22

**Outcome:** Gate (b) content-stories **45/47 → 47/47** live on billion; primary metric **3/8 → 4/8**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (overall a-h; gate part (b) now MET — 4/8)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (tik 1/5) — (6) protocol-stop: n — Outcome: **continue**
**Decisions:** D1, D2, D3 (iter-13/decisions.md).
**Side-deliverables:** none.
**Routes carried forward:** gates (c)/(f)/(h)/(d) + DEF-M239-01 stay per TOK-02's run-6 sequence (next: iter-14).
**Lessons:** a tok-named "preferred fix" is a hypothesis, not a mandate — the mandated probe turned a 20-50 min re-bake into a ~10-line harness fix with no image work. When a sweep reads a "false empty," suspect the HARNESS's settle/readiness model before the app: a page that renders in stages defeats a length-stability settle whose floor sits below the chrome.

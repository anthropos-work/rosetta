**Type:** tik (tooling-iter) — under TOK-01. Ships tooling (the content-stories sweep fixes) AND uses it (runs gate b) in one iter. Protocol: coverage-protocol.md content-stories sweep.

# iter-03 — gate (b) content-stories sweep

## Step 0 re-survey
billion demo-1 GREEN at m243 (iter-02), fresh green autoverify, cockpit serves a 49-pair content-manifest. Target valid.

## What happened (the arc)
1. **bash-3.2 PARSE WALL (fixed).** `run-content-stories.sh` would not run at all on this macOS peer: the `EXPECTED_PAIRS` pin used a heredoc INSIDE a command substitution (`$(python3 - <<'PY' … PY)`), which macOS's stock **bash 3.2** cannot parse (`unexpected EOF`). The sweep is documented to run "from any tailnet peer" — and a macOS peer is bash 3.2. **Fix:** hoisted the heredoc out of `$()`. Verified `bash -n` under 3.2. (rext `a77e89c`.) The other 5 runners already parse under 3.2 — this was the only one.
2. **Gate (b) COLD run → 37/49 FAIL.** 12 non-landing, all `mainLen:0` (empty `<main>`), 0 dropped.
3. **Root-cause: TAILNET COLD-RENDER FLAKINESS, not a data/platform bug.** The failing SET *changed run-to-run* (run1 12 fails, a 20s-settle run 8 fails, different members); every failing session had **full result content in the DB** (e.g. asmt-code-pass IT: 1408 chars of genuine Italian `explanation_summary`), and warm re-visits landed them. The spec's flat **8s settle budget** (`content-result-page.ts`) is too short for a COLD content-story result page over the tailnet (auth SSR + hydration + result GraphQL + fan-out render), so it reads `<main>` before it paints → false empty.
4. **Two robustness fixes (shipped).** (a) `CONTENT_SETTLE_MS` env-overridable, runner sets a higher base for `--host`; (b) **per-pair warm retry** (`CONTENT_PAIR_RETRIES`, default 2 for `--host`) — a reload re-visits a failed pair HOT so the settle clears. A flat 20s settle made the sweep ~28 min for the same result the retry gets in ~6 min, so the committed default is **12s base + 2 retries** (fast base, retry catches stragglers). (rext `01206e7`; tag `sound-check-m244-content-sweep-robustness` → 01206e7 on origin.)
5. **RESULT: reliably 46/49** (37 cold → 46, flakiness eliminated). 3 residuals persist across every config incl. 2 retries.

## The 3 deterministic residuals (NOT flaky — routed forward)
All 3 have **full result content in the DB** but render empty/incomplete — a **voice/interview RESULT-RENDER issue**, not seed or sweep:
- **intv-voice-pass** (SIMULATION_TYPE_INTERVIEW/it, val_attempt=1 explen=1236, interview_extraction=1) → *"no interview acknowledgement text"* — the **player-interview** shape. **This is gate (g) territory** (the interview plan-section-id alignment assertion M244 owns). The natural next step (gate g) likely lands this pair.
- **hire-voice-fail** (SIMULATION_TYPE_HIRING/it voice, explen=1042) → empty render.
- **asmt-voice-pass-en** (SIMULATION_TYPE_ASSESSMENT/en voice, explen=1180) → empty render.
The 2 voice cells render empty despite full content — entangled with the **M240 "voice presence-only"** decision (the recording isn't ported; whether these should render the feedback text or be excluded from the 49 denominator is a fix-shape/design question). NB: most voice cells DID land (hire-voice-pass, intv-voice-fail, asmt-voice-pass-2, etc.), so it is not a blanket voice exclusion.

## Close — 2026-07-22

**Outcome:** gate (b) tooling made runnable + robust on a macOS tailnet peer (3 rext fixes shipped+tagged+pushed); the sweep's tailnet cold-render flakiness FIXED (37→**46/49** reliable). Gate (b) NOT green: 3 deterministic residuals (voice/interview result-render, full DB content) route to gate (g) + a voice-presence-only fix-shape decision.
**Type:** tik (tooling-iter)
**Status:** closed-fixed-partial
**Gate:** NOT MET (gate b 46/49; 2/8 gate parts overall)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **y** (the gate-b residual needs a fix-shape decision — voice/interview render: presence-only-exclude vs demopatch vs gate-g fix — entangled with gate (g) + the M240 voice-presence-only design; surfaced with billion left green for continuation) — (5) cap-reached: n — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D1 (tooling-iter reclassification), D2 (retry over ever-higher settle) — iter-03/decisions.md
**Side-deliverables:** the bash-3.2 parse fix is arguably the load-bearing one — without it NO macOS operator can run the content-stories sweep at all (independent of the flakiness).
**Routes carried forward:** (b-residual) gate (g) interview render likely lands intv-voice-pass; the 2 voice cells (hire-voice-fail, asmt-voice-pass-en) need a voice-render/presence-only decision. Then re-prove gate (b) → target the corrected landable count.
**Lessons:** (1) run the harness runners on the DUMBEST target shell they claim to support — a macOS peer is bash 3.2; a heredoc-in-`$()` is invisible until then. (2) A browser sweep over the tailnet is latency-flaky COLD; the fix is a per-pair warm retry, not an ever-higher flat settle (which just makes it slow). (3) Distinguish flaky (failing SET changes run-to-run) from deterministic (same SET) by re-running — do NOT fix a "render bug" that is really a too-short wait.
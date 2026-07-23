# iter-08 — decisions (local)

- **D1 — the root cause is the DEMOPATCH render-gate SCOPE, not the seed/language.** The two interview clones
  had identical data (val_attempts=1, extraction=1) and differed only in `evaluation_status` (passed vs failed).
  The passed clone rendered EMPTY (mainLen 0), the failed clone the ACK (205). The discriminator is a render-gate
  branch: the M232 `next-web-interview-flag-result` demopatch widened the interview render gate for a demo for
  BOTH scopes, so a PASSED interview (extraction report available) took the player-scope `InterviewReport` branch
  (empty for the player) while a FAILED one fell through to the `UserHiddenResult` ack. Not a language/regex bug.

- **D2 — fix at the demopatch, scoped to the manager vantage; leave the container fetch open.** An interview has
  no player-facing scored report (the sweep's own `player-interview` shape says so). So the demo-widen must open
  the render gate ONLY for the MANAGER scope: `(isManagerScope && !(POSTHOG absent))`. The player then always
  renders the ack (pass AND fail) → the sweep passes; the manager report (gate g) is untouched. Only the RESULT
  (render) demopatch needs scoping — the CONTAINER (fetch) demopatch stays open (the player fetches the data but
  the render gate blocks it; a harmless wasted fetch, one fewer sha to recompute + one fewer test to touch).

- **D3 — verify the sha-computation METHOD before trusting the recomputed post_sha256.** The demopatch G7 apply
  post-condition re-checks post_sha256, so a wrong value would REFUSE the patch (fail-loud, but it would block the
  manager report / gate g on the next bake). I verified my apply-replay (`file.replace(anchor, replacement)`)
  reproduces the CURRENT post_sha `f4eaea3a…` EXACTLY before trusting the new value `99903aec…`. Pristine file
  matches `pre_sha256`; `isManagerScope` (= isManagerView) confirmed in scope at the anchor.

- **live proof coupled to gate h.** A demopatch bakes into the next-web IMAGE, so the live effect (interview
  player renders the ack) lands at the gate-(h) cold reset-to-seed (re-pin billion to `8756ec0` + re-bake).
  Evidence now: root cause proven, fix sha-verified, demopatch + tooling + build tests green.

- **0 platform edits.** The demopatch IS the sanctioned zero-platform-edit mechanism (demopatch-spec.md).

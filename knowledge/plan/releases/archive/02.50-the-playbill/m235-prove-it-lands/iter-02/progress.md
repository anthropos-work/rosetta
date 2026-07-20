**Type:** tik under TOK-01 (Track A) — protocol `playthroughs.md` + `coverage-protocol.md`.

# iter-02 progress — IN-FLIGHT, BLOCKED (user-blocker; no close section)

## Phase 1 Step 0 — re-survey (mandatory) → surfaced a safety blocker

Planned: source + capture + scrub 4 new simulation sessions. The re-survey of the existing fixture substrate
(the exact file this tik would extend) found the anonymization scrub is systematically NOT removing personal
names. Evidence + root cause: see `overview.md` and `decisions.md`. Diagnostic (read-only, no capture run):

- `grep '<<[A-Z_0-9]+>>'` over all 9 `fixture/content/*.json` → **0 matches** (no placeholders anywhere).
- per-fixture given-name heuristic → **8/9** ship a real first name in the copied LLM feedback.
- root cause: `content-capture/main.go:94-116` sources names only from `actors.username`/`alias` (empty here);
  the leaked name comes from the session owner's `public.users` identity, never added to the scrub `repl` map.

## No code landed
This tik ran only its mandatory re-survey; it did NOT run `content-capture`, edit the fixture, or touch the
scrub. There is nothing to commit for the planned target — the blocker surfaced before any implementation.

## Exit
`EXIT_REASON: user-blocker` (Phase 5 §4). Recorded in the milestone-root `decisions.md` (USER-BLOCKER-M235-01);
surfaced to the user. The iter stays **in-flight** (no `## Close` section, committed under a non-closing
`blocker(M235/02):` prefix so it is not treated as a closed iter) pending the user's scrub decision.

_(No `## Close` section — this iter has not closed. Per Phase 4 Step 0, a user-blocker is not a close status.)_

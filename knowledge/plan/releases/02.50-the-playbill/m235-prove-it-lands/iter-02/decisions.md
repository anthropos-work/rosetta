# iter-02 — decisions (local)

## INCOMPLETE-EXIT-2026-07-19 — user-blocker at re-survey
- **What got done:** the mandatory re-survey (read-only) + a precise diagnostic + root-cause of a safety
  finding on the fixture substrate. TOK-01 remains the active strategy; the planned target is untouched.
- **What's left:** the fixture matrix closure (4 new captures) — BLOCKED on the user's scrub decision.
- **What blocked progress:** the anonymization scrub is systematically ineffective (0 placeholders; 8/9
  fixtures leak a real first name). See `overview.md` root cause. Expanding the real-PII footprint by capturing
  4 more sessions is a decision the data-controller must make first.
- **Resume path (next session):** per the user's ruling — (a) harden scrub + re-capture, (b) accept-as-is +
  proceed to capture, or (c) narrow. OR proceed with the non-capture Track-A targets (descriptors / non-sim
  player-path builders / M230 clone re-anchor), which are unblocked.

## The finding detail (for the user's decision)
- **Fact:** 0 `<<ACTOR_i>>`/`<<ORG>>` placeholders across all 9 `contentsession/fixture/content/*.json`; a real
  customer first name survives in 8/9 (Filippo/Raffaele/Madelynn/Simone/Cristian/Marco/Henry/Tram).
- **Root cause:** `content-capture/main.go:94-116` builds the scrub `repl` map only from
  `jobsimulation.actors.username`/`alias` (empty for these sessions). The name in the LLM feedback comes from
  the session owner's `public.users` identity, never sourced into `repl`.
- **Posture note:** `session-clone-spec.md` §6 documents + the data-controller accepted a "best-effort scrub /
  residual re-identification risk, VPN/tailnet-scoped." The material new fact is the scrub removed ZERO names
  (systematic, not occasional) — worth a data-controller re-affirm before extending.
- **Suggested scrub fix (if the user chooses to harden):** source the candidate's real name (the session
  owner's `public.users.first_name`/`name`) into the capture's `repl` map, AND split every actor/owner name
  into whitespace tokens, scrubbing each token ≥3 chars (so a first-name-only reference like "Raffaele" is
  caught). Keep the deterministic + structural-PII gate; add a name-leak regression test. Then re-capture all 9.

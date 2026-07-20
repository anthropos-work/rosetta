---
iteration: iter-03
iteration_type: tik
status: closed-fixed
created: 2026-07-19
---

# iter-03 — harden the anonymization scrub + re-capture the 9 fixtures (resolve USER-BLOCKER-M235-01)

**Active strategy reference:** TOK-01 (bootstrap, two-track) — Track A "Readiness (buildable + unit-provable)".
This tik is a **prerequisite hardening** of the fixture substrate before Track A step 1 extends it (the
4 new captures). It resolves the user-blocker the run-1 iter-02 surfaced.

**User ruling (2026-07-19):** "Fix scrub + re-capture." (USER-BLOCKER-M235-01, milestone-root decisions.md.)

## Cluster / target identified
iter-02 re-survey (read-only) proved the scrub removed ZERO names: 0 `<<ACTOR>>`/`<<ORG>>` placeholders in
any of the 9 `contentsession/fixture/content/*.json`; 8/9 leak a real first name in the copied LLM feedback.
Root cause (`cmd/content-capture/main.go`): the scrub `repl` map is built ONLY from
`jobsimulation.actors.username`/`.alias` (empty for these sessions). The leaked first name comes from the
session OWNER's `public.users` identity (`jobsimulation.sessions.owner_id` → `public.users.firstname/lastname`),
which is never sourced into `repl`.

## Hypothesis
Sourcing the session owner's real identity (firstname/lastname + email local-part) from `public.users` into
the scrub `repl` map — mapped to the PLAYER placeholder `<<ACTOR_0>>` — plus whitespace-token-splitting every
actor/owner name (each token ≥3 chars) so a bare first-name mention is caught, will scrub the leaked names.
Re-capturing the 9 fixtures through the fixed path will write placeholders where names were and leave no real
name behind.

## Phase plan
- **Scrub package:** add `ExpandNameTokens`/`AddNameReplacements` (full name + ≥3-char whitespace tokens → placeholder);
  unit tests proving the owner-first-name-only path scrubs a bare first name.
- **Capture (`cmd/content-capture`):** source `public.users` for `sessions.owner_id` (firstname/lastname/email
  local-part) → `<<ACTOR_0>>`; token-split actors + owner; add a **capture-time post-condition** (G-post): after
  scrubbing a session, ASSERT none of the sourced name tokens survive (case-insensitive) in any output string
  leaf — FAIL the capture if one does (so a leak can never be written). Names are known in-process, never persisted.
- **Offline cleanliness gate:** extend `TestEmbeddedContent_NoStructuralPII` to also assert the fixture SET
  carries `<<ACTOR_0>>` + `<<ORG>>` placeholders (the "0 placeholders anywhere" bug's regression tripwire) —
  arbitrary first names can't be detected offline, so the capture-time post-condition is the name-leak gate.
- **Re-capture:** run `content-capture` against prod READ-ONLY (marco_read via ~/.pgpass, over Tailscale) —
  counts-only stdout, raw content never printed. Re-writes all 9 `fixture/content/*.json`.
- **Verify:** grep (exit-code / counts only, no content printed) that the 8 flagged names are ABSENT and
  placeholders are PRESENT across all 9 fixtures; run the full unit gate.
- **Docs:** fix KB-1 (the `content-sessions.yaml` stale synthesize-first header) + correct the "what gets
  scrubbed" description in `session-clone-spec.md` §6 + `safety.md` §3.8 (owner-identity name now scrubbed;
  keep the honest best-effort / residual-risk / VPN-scoped posture).
- **Record + re-tag:** milestone `decisions.md` (resolution of USER-BLOCKER-M235-01); rext tag `playbill-m235-scrub-fix`.

## Expected lift
Gate metric (live (session×action) landings) does not move this tik — this is Track-A readiness/hardening.
Provable outcome: 9/9 fixtures clean (0 leaked names, placeholders present); unit gate green; capture-time
name-leak post-condition in place.

## Escalation conditions
- Prod unreachable (Tailscale/pgpass) → cannot re-capture → user-blocker (re-capture is the user's ruling).
- A sourced name token STILL survives after the fix → root-cause deeper leak surface, do not ship a partial scrub.

## Acceptable close-no-lift outcomes
N/A for the leak-fix goal — this tik must land the scrub fix + clean fixtures (Fate-1). If prod is unreachable
the iter does not close (user-blocker).

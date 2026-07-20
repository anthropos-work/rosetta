---
iter: iter-02
milestone: M235
iteration_type: tik
status: in-flight-blocked
created: 2026-07-19
active_strategy: TOK-01
---

# M235 · iter-02 — tik (fixture matrix closure) — BLOCKED at re-survey (user-blocker)

**Type:** tik under **TOK-01** (Track A). **Status: in-flight, BLOCKED — no `## Close` section** (the
mandatory Phase 1 Step 0 re-survey surfaced a safety user-blocker before any capture ran; no code landed).

## Planned target (TOK-01 next-tik direction)
The **fixture matrix closure** — source + `content-capture` + scrub the 4 missing simulation sessions (a 2nd
assessment voice-PASS, an assessment doc-PASS, a hiring not-passed, an interview not-passed) to satisfy the
assessment PASSED set (2 voice / 1 code / 1 document) + each sim type in passed AND not-passed; fix KB-1; unit-
prove the fan-out + fixture-cleanliness gate + datadna closure.

## Re-survey (Phase 1 Step 0) — what it found instead
Before capturing NEW real prod sessions, the re-survey inspected the existing fixture substrate M235 extends
and found the anonymization scrub is **not removing personal names** — a systematic PII exposure, not the
"occasional residual" the spec describes:

- **0 `<<ACTOR_i>>`/`<<ORG>>` placeholder tokens exist in ANY of the 9 checked-in fixtures** — the seeder's
  `fillPlaceholders` (content_stories_write.go) has nothing to fill.
- **8 of 9 fixtures ship a real customer FIRST NAME** in the copied LLM feedback: Filippo, Raffaele (24×),
  Madelynn, Simone, Cristian, Marco, Henry, Tram — each as "{Name} ha/showed/demonstrated…".

### Root cause (code-cited)
`cmd/content-capture/main.go:94-116` builds the scrub `repl` map ONLY from `jobsimulation.actors.username` /
`.alias` (`coalesce(...,'')`, guarded by `if username != ""`). For these sessions those columns are empty →
`repl` carries no names → `scrub.Scrub(text, repl)` replaces nothing. But the candidate's real first name
appears throughout the LLM feedback because it originates from the **user identity** (the session owner's
`public.users` name), which the capture never sources into `repl`. So the scrub has no knowledge of the name
that is actually in the text.

## Why this is a user-blocker (Phase 5 §4), not a route-forward
M235's central Track-A work is to capture 4 MORE real prod sessions into this exact fixture — expanding the
exposure by ~44%. The decision — **(a)** harden the scrub (source the owner's real name + strip its
first/last components) + re-capture the existing 9, **(b)** re-affirm the accepted-residual + VPN/tailnet-scope
posture as-is (`safety.md` §3.8) and proceed, or **(c)** narrow the sourcing — is a **data-controller / safety
decision that changes what code lands** in the capture tik (and possibly reworks the closed M232 deliverable +
its shipped fixtures). Per Phase 5 §4 that is a user-blocker; expanding the real-PII footprint before the user
weighs in on the anonymization control's effectiveness is the wrong default. No fake proof, no platform edit.

## What is NOT blocked (available for the next session under the user's decision)
The other Track-A targets don't touch the scrub or add PII and remain buildable: the Playthrough + coverage
descriptors for the existing sessions, the non-simulation product player-path builders (content-stories-spec §6),
and the M230 clone re-anchor. They can proceed once the user rules on the scrub decision (or in parallel if the
user greenlights them independently).

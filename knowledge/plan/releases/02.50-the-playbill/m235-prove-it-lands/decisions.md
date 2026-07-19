# M235 — Decisions

_(decisions recorded as they arise during build)_

## TOK-01: two-track — build-and-unit-prove readiness HERE, route the live browser proof — 2026-07-19

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Split M235's gate into the two things it actually is, and drive them as separate tracks:

- **Track A — Readiness (buildable + unit-provable in THIS environment; the bulk of the tiks, all Fate-1).**
  Build everything that makes the live proof possible, and prove it with the existing unit/validator gates:
  1. **Fixture matrix closure** — source + `content-capture` + scrub the 4 missing simulation sessions (a 2nd
     assessment voice-PASS, an assessment doc-PASS, a hiring not-passed, an interview not-passed) so the
     assessment PASSED set = 2 voice / 1 code / 1 document and every sim type is present in passed AND
     not-passed. Fix KB-1 in the same edit. Prove via `TestContentStorySeeder_*` + `TestEmbeddedContent_NoStructuralPII` + `datadna measure-closure`.
  2. **Non-simulation product sections** — add skill-path (legacy runtime rows + `local_skill_path_session`
     mirror), academy (`academy_chapter_progress`, depends on M230 catalog), and ai-labs (presence-only)
     fixture sections + the non-simulation player-path builders (`content-stories-spec.md` §6 homes these to
     M235). Re-project `content-manifest.json`; keep the honesty gate (`CanonicalFileMatchesProjection`) + the
     fail-closed resolver green.
  3. **Playthrough + coverage descriptors** — declare a Playthrough use case per (session × action) in the
     playthroughs manifest (manifest-first, P5) + a coverage descriptor asserting non-zero rendered VALUES
     (the M219/M222 mirror-table trap → `textMatch`-on-values, so a blank clone reads RED). Prove via
     `ptvalidate` (both-way integrity + precondition-coverage + datadna closure).
  4. **M230 carry-forward** — re-anchor the drifted local `next-web` clone (the 2 demopatch manifests) as a
     cold-`/demo-up` prerequisite; the `ANT_ACADEMY` coverage descriptor consuming tag
     `playbill-m230-academy-fs-published`; the `getPublicCatalogView` 2nd-manifest anonymous-routes follow-on.

- **Track B — the FORMAL live gate (needs a running stack).** The cold reset-to-seed + Playwright run proving
  every CTA lands NON-EMPTY for both vantages, 0 ejects. Attempt a local cold `/demo-up` once Track A makes it
  ready; if the local box genuinely can't (M230 precedent: drifted clone + heavy bring-up + prior docker
  trouble), document the constraint honestly (a `closed-no-lift` iter with documented falsification) and route
  the definitive live proof to **M236 (on-`billion`)** — exactly as the M230 carry-forward projected ("M235
  primarily; M236 as the live confirmation"). Exit `EXIT_REASON: user-blocker` ONLY if a genuine environment
  DECISION is needed that changes what code lands. Never fake a live proof; never platform-edit to force a render.

**Rationale:** The iterative protocol (playthroughs + coverage) is a measure→triage→fix→re-measure loop against
a LIVE sweep — but the sweep needs a stack, and none is up. Rather than stall the whole milestone on a heavy,
M230-blocked local bring-up, Track A lands + unit-proves 100% of what's provable without a browser (the fixture
substrate, the manifest projection, the descriptors, the clone re-anchor), so the live proof — whether it runs
locally or on billion at M236 — is a measurement, not a build. Sequencing the fixture substrate first is the
dependency spine (manifest projection, descriptors, and live landing all read it).

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).
**Distance-to-gate context:** gate metric = (session × action) landing NON-EMPTY on a cold reset-to-seed for
both vantages, 0 ejects. Current live-proven = 0 (no stack). Fixture readiness ≈ 5/9 requirements met
(assessment+training both-states ✓; missing +1 asmt-voice-pass, +1 asmt-doc-pass, hiring not-passed, interview
not-passed; 3 product sections absent). Track A drives readiness → 100%; Track B is the live measurement.
**Next-tik direction:** iter-02 = the fixture matrix closure (the 4 missing simulation sessions + KB-1 fix +
seeder/fixture-cleanliness/closure unit proofs).

## USER-BLOCKER-M235-01: the anonymization scrub removes ZERO names — decide before extending the fixture — 2026-07-19

**Surfaced by:** iter-02 (tik) Phase 1 Step 0 re-survey, before any new capture ran. **EXIT_REASON: user-blocker.**

**The finding (rigorously verified, read-only):** the content-story anonymization scrub is systematically NOT
removing personal names from the shipped M232 fixtures:
- **0** `<<ACTOR_i>>`/`<<ORG>>` placeholder tokens exist in ANY of the 9 `contentsession/fixture/content/*.json`
  (so the seeder's `fillPlaceholders` fills nothing).
- **8 of 9** fixtures ship a real customer FIRST NAME in the copied LLM feedback (Filippo, Raffaele ×24,
  Madelynn, Simone, Cristian, Marco, Henry, Tram — each "{Name} ha/showed/demonstrated…").

**Root cause (code-cited):** `cmd/content-capture/main.go:94-116` builds the scrub replacement map ONLY from
`jobsimulation.actors.username`/`.alias` (both empty — `coalesce(...,'')` — for these sessions). The candidate's
real first name appears throughout the LLM feedback because it comes from the session owner's **`public.users`
identity**, which the capture never sources into the replacement map. The scrub therefore has no knowledge of
the name that is actually in the text.

**Posture:** `session-clone-spec.md` §6 + `safety.md` §3.8 document a data-controller-ACCEPTED "best-effort
scrub / residual re-identification risk, VPN/tailnet-scoped." The **material new fact** is that the scrub
removed **zero** names (systematic, every session), not the "occasional residual" the acceptance was premised on.

**Why it blocks M235:** the milestone's central Track-A task is to capture **4 more** real prod sessions into
this exact fixture (+~44% footprint). Whether to **(a)** harden the scrub (source the owner's real name + strip
first/last name tokens) + re-capture the existing 9, **(b)** re-affirm the accepted-residual + VPN-scope posture
and proceed to capture, or **(c)** narrow the sourcing — is a data-controller decision that changes what code
lands (and may rework the closed M232 deliverable + its shipped fixtures). Expanding the real-PII footprint
before the user rules is the wrong default. No fake proof; no platform edit.

**Recommendation:** harden the scrub before extending — source `public.users.first_name`/`name` for the session
owner into the `repl` map AND token-split every actor/owner name (scrub each token ≥3 chars), add a name-leak
regression test, re-capture all 9 fixtures, THEN proceed to the 4 new captures. (The `git`-diffable fixtures make
a re-capture auditable.) The user may instead re-affirm accept-as-is given the VPN-scope control.

**Not blocked (available next session regardless of the ruling):** the Playthrough + coverage descriptors for
the existing sessions, the non-simulation product player-path builders (`content-stories-spec.md` §6), and the
M230 clone re-anchor — none touch the scrub or add PII.

## RESOLUTION of USER-BLOCKER-M235-01 — user ruled "Fix scrub + re-capture" (2026-07-19, iter-03)

**User ruling (2026-07-19):** "Fix scrub + re-capture." Landed in iter-03 (tik under TOK-01, Track A) as a
prerequisite hardening of the fixture substrate before extending it.

**What changed (rext `stack-seeding`, tag `playbill-m235-scrub-fix`):**
- **`cmd/content-capture/main.go`** now sources the **session owner's real identity** from `public.users`
  (`sessions.owner_id` → `firstname`/`lastname` + email local-part) into the scrub `repl` map, mapped to the
  **player placeholder `<<ACTOR_0>>`**. This is the name that was leaking (threaded through the LLM feedback);
  it was never in `jobsimulation.actors` (whose `username`/`alias` are empty for these sessions), which is why
  the original scrub removed zero names.
- **`scrub` package**: `NameTokens` (full name + each ≥3-char whitespace token) + `AddNameReplacements`
  (first-writer-wins) so a **bare first-name** mention is caught; matching is now **word-boundary-aware**
  (a short token never corrupts a common word — "Ann" ≠ "announce"). `SurvivingToken` is the new leak probe.
- **Capture-time fail-closed post-condition (G-post)**: after scrubbing a session the capture asserts **no
  sourced name token survives** any free-text leaf and **refuses to write the fixture** if one does (prints
  only the field name + token length — never a value). This is the name-leak gate the offline test cannot be
  (it can't know arbitrary names); the names are known in-process and never persisted.
- **Offline cleanliness gate** (`TestEmbeddedContent_NoStructuralPII`) extended: still fails on structural PII,
  now **also asserts the set carries the `<<ACTOR_0>>` placeholder** — the "sourced zero names → zero
  placeholders" regression tripwire (it went RED on the buggy fixtures, GREEN after re-capture).
- **Re-captured all 9 fixtures** through the fixed path (read-only prod, marco_read, counts-only).

**Verified clean (counts/exit-codes only, no content read into context):**
- `<<ACTOR_0>>` placeholders present in **9/9** fixtures (**545** total, was **0**); per-file 44–107,
  correlating with transcript length → no over-redaction explosion.
- The **8 flagged first names** (Filippo/Raffaele/Madelynn/Simone/Cristian/Marco/Henry/Tram): **0** matches
  across all 9 fixtures (was 8/9 leaking).
- Full unit gate GREEN (contentsession + scrub + seeders content-story); module-wide `go vet`/`go test` clean;
  existing scrub regression tests still pass (word-boundary change is backward-compatible).

**Posture (unchanged, now honestly described):** still **best-effort, NOT provably clean** — a name never
*sourced* (a third party mentioned in passing) can still survive; residual re-identification risk remains
**ACCEPTED by the data-controller, VPN/tailnet-scoped**. What changed is the *what-gets-scrubbed* description:
the owner-identity name is now removed, and a *sourced* name can no longer be silently written. Docs corrected:
`session-clone-spec.md` §3+§6, `safety.md` §3.8, and the KB-1 stale `content-sessions.yaml` header.

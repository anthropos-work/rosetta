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

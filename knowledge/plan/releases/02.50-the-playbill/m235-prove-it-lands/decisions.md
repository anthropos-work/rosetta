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

## USER-BLOCKER-M235-02: Track-A step-3 "coverage descriptor" mechanism doesn't exist as TOK-01 assumed — a scope/sequencing decision — 2026-07-19

**Surfaced by:** iter-05 planning (Phase 0d/Phase 1 pre-flight against the coverage harness), before any iter-05
code landed. **EXIT_REASON: user-blocker.** No half-iter (iter-05 was never opened as a dir; this is a
milestone-level finding, like USER-BLOCKER-M235-01).

**The finding (concrete, code-cited):** TOK-01 Track A step 3 planned "a coverage descriptor asserting non-zero
rendered VALUES (the M219/M222 mirror-table trap → textMatch-on-values)" for the content-stories result pages.
The existing coverage harness **structurally cannot express that**:
- `stack-verify/e2e/lib/coverage-manifest.ts` `pageDescriptorFor` matches a page descriptor by **EXACT
  normalized path** (`p.path === norm`, `coverage-manifest.ts:989-991`). The content-stories player result page
  is `/sim/<slug>/result/<sessionId>` — a **dynamic** path (per-seeded-session uuid); there is no
  dynamic-segment normalization, so no static descriptor path can match it.
- `lib/crawl.ts` reaches pages by **crawling in-app nav from the hero's landing** (BFS over links). The
  content-stories result pages are reached **only via the cockpit "login as content-player-N" CTA**, never
  linked from the hero crawl — so the sweep never lands on them.

So the "coverage descriptor" mechanism as planned isn't a manifest add — it needs **new harness plumbing**: a
content-stories spec that logs in as each `content-player-<idx>` seat (the Playthroughs-style cockpit
seat-switch, `playthroughs/e2e/lib/hero-login.ts`) + resolves each session's **exact** result URL from the
seeded `content-manifest.json` + reuses the shared `AISimulationResultContainer` page-object. That plumbing is
authored **and calibrated against a LIVE seeded render** (selectors, the mirror-table manager scoreboard, the
per-session score/feedback fence) — exactly the Track-B work TOK-01 routed to **M236 (on-billion)**. Authoring
it blind here would ship an **incorrect** (not merely uncalibrated) descriptor into a load-bearing harness.

**Why it's a decision (changes what code lands):** the remaining TOK-01 Track A steps split into:
- **Step 3 (descriptors) + Step 4 (M230 carry-forward: next-web clone re-anchor + ANT_ACADEMY coverage
  descriptor)** — decisively **Track-B / M236-coupled** (need a live seeded render + the new seat-login sweep
  plumbing to author + validate correctly). Recommend **Fate-2/3 → M236**, and M236's plan must now include the
  **new content-stories seat-login coverage/Playthrough plumbing** (this finding), not "add a manifest section".
- **Step 2 (non-simulation product sections: skill-path / academy / ai-labs)** — the one remaining
  *offline-buildable* Track-A surface, but it is a **large new-seeder chunk** (a skill-path/academy fixture
  model + prod sourcing + a runtime-fan-out seeder + the `local_skill_path_session` mirror — comparable in size
  to the whole M232 `ContentStorySeeder`), not a single tik. Opening it in the remaining run budget risks a
  half-landed iter (which the run's dirty-tree ban forbids).

**The decision for the user/orchestrator:** run-2's deliverable is the **scrub fix (USER-BLOCKER-M235-01
resolved, provably clean) + the SIMULATION content-product substrate closure** (13 sessions; assessment PASSED
= 2 voice / 1 code / 1 document; every type passed AND not-passed) — the gate's core substrate, unit-proven.
Should the next work be (a) a **dedicated fresh session** to build the Step-2 non-sim product seeders, and/or
(b) **fold Steps 3-4 into M236** as the live-authored content-stories coverage/Playthrough plumbing (with the
new-plumbing requirement above)? Both change what lands next; neither is cleanly completable in remaining
run-2 budget without a half-iter. No blind descriptors; no platform edit.

**Not blocked / already delivered this run:** the scrub fix + the 4-cell fixture-matrix closure + both
regenerated honesty manifests are landed, committed, tagged, unit-proven, and provably leak-clean.

## RESOLUTION of USER-BLOCKER-M235-02 — user ruled "Build non-sim seeders, then close" (2026-07-20, run 3)

**User ruling (2026-07-20):** "Build non-sim seeders, then close." Landed across run-3 iters 05–08 under
TOK-01 Track A step 2, all Fate-1 (offline-buildable + unit-provable). The LIVE proof legitimately routes to
M236 (Fate-3, user-authorized).

**What landed (rext `stack-seeding`, tags `playbill-m235-nonsim-{skillpath,ailabs,academy}`):**
- **iter-05 — skill-path-legacy (real progress).** A separate code-owned non-sim exhibit registry
  (`seeders/content_nonsim.go`) + `ContentStoryNonSimSeeder` writing `skillpath.skill_path_sessions` (real
  progress) + the `public.local_skill_path_sessions` MIRROR (non-blank manager scoreboard, the M219/M222
  trap), owned by `content-player-<idx>` seats, pinned to REAL public `skill_path_id`s sourced OFFLINE from
  the captured snapshot; the `/skill-path/<id>` + mirror-manager routes project; version "2" collision-safe.
- **iter-06 — ai-labs (presence-only, M231 §5).** The presence projection arm (label + seat, NO CTA) + the
  `public.lab_sessions` status/spend seeder arm (12-char hex id) surfaced on `/labs` + `/enterprise/labs`.
- **iter-07 — academy (skill-path-new).** app_base=academy, a REAL public `/library/<slug>` course CTA
  (direct academy-origin + e2e_persona seam), no manager view. The `academy_chapter_progress` write is the
  live `app/cmd/academy-seed` platform binary (M236), not an in-process rext write.
- A believable `Label` field (real course/lab titles) renders in the cockpit; the honesty gate regenerated to
  18 sessions (all 4 products resolve); all Go + 47 Python render tests GREEN, 6 pre-existing failures
  unchanged. Deterministic + idempotent + audited + fail-closed + n=0-guarded, single-sourced with the
  projection (the owner↔seat invariant is unit-proven).

**What is UNIT-PROVEN here vs ROUTED to M236 (Fate-3, user-authorized 2026-07-20).** Unit-proven = the
seeders write structurally-correct rows + the manifest sections RESOLVE + the cockpit renders them. The LIVE
(session × action)-lands proof needs a running stack → **M236** (its `In:` list is EDITED to own it). The
Fate-3 handoff M236 now carries: (1) the NEW content-stories seat-login coverage/Playthrough plumbing
(USER-BLOCKER-M235-02's core finding — the exact-path/hero-crawl harness can't reach the dynamic-URL,
cockpit-seat-reached result pages; it must be authored + CALIBRATED against a live render); (2) the per-section
live-calibration checklists (skill-path version-match/status/mirror — iter-05 decisions.md; ai-labs
lab_sessions DDL — iter-06; academy progress-write/route/catalog — iter-07); (3) the M230 carry-forward live
items (ANT_ACADEMY coverage descriptor + next-web clone re-anchor + getPublicCatalogView 2nd manifest).

**Milestone disposition.** After iter-08 the OFFLINE-buildable scope is EXHAUSTED (all 3 non-sim sections
built + unit-proven; the live proof legitimately routes to M236). The milestone PRAGMATIC-CLOSES per the
user's "then close" mandate — the actual `/developer-kit:close-milestone` merge is a separate step the
orchestrator/user drives. No live gate was faked; no platform-repo edit. EXIT_REASON for the run:
protocol-stop (offline clusters exhausted; live proof routed to M236).

# v2.6 "sound check" — design spec (design-roadmap, 2026-07-20)

Reliability / field-hardening release. Trigger: live demo defects — "still not all gets built and
provisioned as expected." Tooling + docs only, **zero platform-repo edits** (dead surfaces → sha-pinned
demo-patch or escalate). House shape: **barrier → parallel fixes → prove-on-billion** (v2.3/v1.10b lineage).

Branch `release/02.60-sound-check`; tag will be `v2.6`. Continues the `M2xx` scheme at M237.

## Phase 0 verdicts
- **0a deferral audit: proceed.** All v2.5 deferrals re-fated 2026-07-20 with named M237/M238 destinations; nothing blocking. (Report: `releases/archive/02.50-the-playbill/release-deferrals.md`.)
- **0b KB blind areas:** clone-freshness anchored (`rosetta_demo.md` §Clone freshness, v2.5); secrets DNA (`secrets-spec.md`) exists but needs a Bedrock-cred extension (M239 Delivers); media-porting (Chime/S3) has NO doc anchor → M240 Delivers a media-substrate spec + a `safety.md` amendment for raw media.

## Research provenance
`.agentspace/scratch/tasks/` — three agents (provisioning reliability / content-stories fidelity / carries+history+cockpit+playthrough). Headline: only defect #1 (menu) is clone-staleness; #2–#5 each reproduce on origin/main. Almost every fix needs a live billion bring-up + prod read → the barrier + the iterative closer bracket the release.

## User decisions (2026-07-20)
1. **Talk-to-data → FULL.** Wire real AWS Bedrock creds via the `/stack-secrets` provisioning mechanism; reference `../hyper-studio/.env.example` (key set `AWS_ACCESS_KEY_ID`/`AWS_SECRET_ACCESS_KEY`/`AWS_SESSION_TOKEN`/`AWS_REGION` + `CLAUDE_CODE_USE_BEDROCK`). Extend the secret-coverage DNA for the `app` service. (M239)
2. **Media → PORT IT.** Capture + re-host the Chime/S3 voice recording + document blobs so the manager can hear the call / see the document. **Expands the customer-PII surface to raw audio + full documents** → M240 carries a HARD internal gate: fresh data-controller sign-off + `safety.md` amendment for raw media + a voice/document anonymization decision (you cannot token-scrub a voice). Likely pulls in DEF-M10-01 (S3 read access). (M240)
3. **Language → EN-only fallback per tuple.** M241 opens with a read-only pool-count query (IT sessions per tuple); toggle where IT exists, EN-only where absent. No blocking. (M241)

## Milestones (8), execution order

### M237 — "clean stage" (SECTION, HARD go/no-go barrier)
**Goal:** the demo builds from CURRENT platform source, and the ambiguous UI defects are re-triaged on a correct build — so every downstream fix is scoped against reality, not stale code.
**In:**
- Fix clone-freshness in `rext demo-stack/ensure-clones.sh`: a **fetch-verified** freshness assertion (never suppressed-stderr — the billion `root` host-key failure that gave 12-vs-202) + an opt-in advance-to-`origin/main`-or-pinned-tag path + a real pin model so "pinned" vs "stale-by-neglect" is distinguishable (today both read `ref:"main"`/`"HEAD"`).
- Fix **F-M236-CLOSE-2**: the R1 pristine sweep enumerates all 14 patch manifests, not the hard-coded 3.
- Bring up a **fresh-clone demo on billion**; produce a **confirmed-defect ledger**: verify #1 menu now hierarchical for managers; RE-TRIAGE #2 academy-language + #4 library-empty on the fresh build (which survive?).
**Delivers →** `corpus/ops/rosetta_demo.md` (clone-freshness mechanism), `corpus/ops/demo/demopatch-spec.md` (R1 all-manifests).
**Why barrier:** any UI-defect triage on a stale-clone demo is untrustworthy (the M217 "clean stage" logic).
**Depends on:** none (opens the release). **Parallel:** none (gates everything).

### M238 — "ant-academy reliability" (SECTION)
**Goal:** a hero can follow a course and actually consume a chapter; language switch works.
**In:** fix #3 (Start→404) — the demo academy chapter-body path is unwired (bodies are backend-authoritative, no FS fallback; the catalog demopatch covers only the catalog). Wire a chapter-body demo path (a chapter-body FS-fallback demopatch analogous to `academy-fs-published-fallback`, OR wire the academy backend for the demo). Fix #2 (language error — re-triaged in M237; likely the same backend-null path). Extend the **academy presence/coverage sweep**.
**Depends on:** M237. **Parallel with:** M239, M240, M243.

### M239 — "enterprise surfaces" (SECTION)
**Goal:** talk-to-data works live; the library grid loads first-time; the hierarchical manager menu is confirmed.
**In:** talk-to-data **(a)** flag enablement (`NEXT_PUBLIC_DEMO_FLAGS_ALL` or a flag-gate demopatch, M219/M232 pattern) + **(b) real AWS Bedrock creds** provisioned via `/stack-secrets` + the secret-coverage DNA extension for `app` (hyper-studio template) + mounted/env-wired into the app compose service. Fix #4 (library empty-first-load: the client-fetch race + the open non-offset `:5050` `apps/web` endpoint carry). Confirm #1 hierarchical menu renders for managers (presence sweep).
**Delivers →** `corpus/ops/secrets-spec.md` (Bedrock cred class for `app`).
**Open:** the demo secrets-posture for AWS creds (safety.md note — same class as AI-provider keys, now present-not-absent for `app`).
**Depends on:** M237. **Parallel with:** M238, M240, M243.

### M240 — "content-stories fidelity" (SECTION, with a HARD media-safety gate)
**Goal:** the cockpit's claim matches the session — right type, playable call, visible document — at a believable pass rate.
**In:**
- Defect 1 (selection): tighten `rext stack-seeding sourcing.go` to constrain the public sim's type to the cell type (exclude the interview sim from non-interview cells); re-pin `content-sessions.yaml`.
- Defect 3 (document): write the dropped `input_data` at seed time (`content_stories_write.go` / a content-specific criterion column set); + **port the document blob** if the body is a `storage_upload` (per user decision 2).
- Defect 2 (voice): **port the Chime/S3 recording** — capture the recording reference + re-host the audio in the demo storage tier + flip `chime_status` to available (per user decision 2).
- Pass-rate (#4-feature): add a score-band to `SelectionSpec` (`AND s.score BETWEEN 70 AND 95`), flip the tiebreak to `score ASC` (prefer lower), 100% only as fallback; re-capture.
- **HARD internal gate (before any customer media lands in a demo):** fresh data-controller sign-off + `safety.md` amendment covering raw audio + full documents + a voice/document anonymization decision. Likely consumes **DEF-M10-01** (S3 read access).
**Delivers →** a media-substrate spec (`corpus/ops/demo/` — new), `corpus/ops/safety.md` §3.8 amendment (raw media).
**Depends on:** M237. **Parallel with:** M238, M239, M243. **Note:** re-capture needs prod read (`~/.pgpass`).

### M241 — "content-stories language" (SECTION, opens with a pool-count go/no-go)
**Goal:** each session is consumed in its intended language, with an EN/IT cockpit toggle.
**In:** read-only prod pool-count query (IT sessions per requirement tuple — the interview-scarcity go/no-go); add `s.language` to `sourcing.go` SELECT + optional filter; add a `language` field to the fixture + `content_manifest.go` projection (re-touch the `CanonicalFileMatchesProjection` honesty gate); use `cs.Language` instead of the hard-coded `sessLanguageEnglish`; source EN+IT pairs per tuple where IT exists; **EN-only fallback per tuple** where absent (toggle hidden/disabled there); cockpit toggle swaps the login-and-land target. Extend the **content-stories sweep** for language (assert structure/presence, never the translated value — P2 forbids copy assertions).
**Depends on:** M240 (shares `stack-seeding` + the re-capture). **Parallel:** none (serial after M240).

### M242 — "cockpit UX" (SECTION)
**Goal:** the Content-stories tab reads clearly and the heroes are legible by role.
**In:** (1) row layout — regroup by requirement tuple `(sim_type, modality)` → `target | passed login options | not-passed login options` on one row (render-only; fields exist); (2) tab selector — move into the white header, right, vertically centered (restructure `cockpit.py` header to flex; preserve the byte-identical-when-no-content-manifest invariant); (3) hero icon bg by user-type (manager=orange / employee=indigo, reuse badge palette; derive a candidate color = `is_hiring && vantage != manager`). Extend the **cockpit specs**.
**Depends on:** M240 + M241 (the row regroup wants the pass/fail variants + the language axis). **Parallel:** none.

### M243 — "assign-WRITE Playthrough" (SECTION) [realizes reserved M238]
**Goal:** the one net-new hero journey — a manager assigns content with a deadline and it lands.
**In:** `playthroughs/manifest/assignment-monitoring.yaml` UC1 (`assign-and-track.UC1`, currently `TODO`); a new `/enterprise/assignments` page object; possibly a `pt-world` precondition (assignable content + target member) in lockstep with `seed-worlds.yaml`; the spec `e2e/tests/assignment-assign.spec.ts` tagged `@pt:...UC1`. Takes the corpus 15→16 live Playthroughs, 0 TODO. Needs a live browser drive against a running demo.
**Depends on:** M237 (fresh demo). **Parallel with:** M238/M239/M240.

### M244 — "prove on billion" (ITERATIVE, the closer) [realizes reserved M237]
**Goal:** re-prove the whole feature — v2.5's headline AND every v2.6 fix — live on billion, cold reset-to-seed.
**Exit gate:** on a cold reset-to-seed on billion: (a) ORG-CLEAN reports 0 surviving source-org tokens (or each dispositioned) — RUN FIRST, read-only, before the bring-up; (b) content-stories `run-content-stories.sh` green at the shipped harness with the CQ-1 grader fix + CQ-2 runner wiring + externally-sourced `EXPECTED_PAIRS` (discharges CLOSE-D3); (c) the 39 live-browser specs execute green (T-3); (d) the anonymous academy `/library`+`/free` twin renders real cards (S-1); (e) DEF-M226-01 — the serve-reap self-resolution claim is **actively tested** or DROPPED; (f) the 3 v2.3 drift-carries burned-in live (BURNIN-M221 / F-M220-4 / PROBE-M218-c3); (g) the interview plan-section-id alignment assertion added + green (S-8/S-9); (h) every v2.6 fix (academy course-start, talk-to-data live answer, library, content fidelity incl. media, language toggle, cockpit UX) proven live; p95 click→ACCESS < 5 s hero vantages. 0 platform edits.
**Iteration protocol:** `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md` + `coverage-protocol.md` + `playthroughs.md`.
**Why iterative:** live-proof is measurement-driven; iters until the gate (M221/M236 lineage).
**Re-scope trigger:** 5 consecutive toks without a viable strategy → user-strategic-replan.
**Depends on:** M238, M239, M240, M241, M242, M243 (all fixes). **Parallel:** none (terminal).

## Execution graph
```
M237 clean stage (barrier)
  ├─▶ M238 academy ─────────────┐
  ├─▶ M239 enterprise ──────────┤
  ├─▶ M240 content-fidelity ─┐  │
  │      └─▶ M241 language ─┐ │  │
  │            └─▶ M242 cockpit-UX
  ├─▶ M243 assign-WRITE ────────┤
  └────────────────────────────▶ M244 prove-on-billion (closer)
```

## Reservation remap
Reserved **M237 "re-prove-on-billion"** → realized as **M244**. Reserved **M238 "assign-WRITE"** → realized as **M243**. (Renumbered to read in execution order; the archived `release-deferrals.md` reservations explicitly permit this.)

## Risks / open decisions
- **R1 (blocks-quality): raw-media PII.** Porting real customer voice + documents is a larger data-controller call than v2.5's scrubbed text; M240's internal gate (sign-off + safety amendment + anonymization decision) must clear before any customer audio lands in a demo. Voice can't be token-scrubbed.
- **R2 (blocks-scope): language scarcity.** IT interview sessions may not exist; M241's pool query decides per-tuple coverage; EN-only fallback where absent.
- **R3 (degrades-quality): AWS Bedrock creds in the demo.** A new present-not-absent secret class for `app`; secrets-posture note in safety.md.
- **R4 (dependency): prod read + live billion.** M240/M241 re-capture + M244 re-prove need `~/.pgpass` prod read + a reachable billion. Both confirmed available at v2.5 close.
- **R5 (process): v2.5 not pushed to origin.** The v2.6 branch cut from local `main`; `main` + tag `v2.5` still local-only. Flag to user.

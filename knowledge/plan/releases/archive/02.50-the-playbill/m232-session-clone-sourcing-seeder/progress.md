# M232 — Progress

## Sections
(derived from overview.md + roadmap.md §M232 + content-stories-routes.md, at build time)

- [x] **S1 — Prod-session sourcing + source-pinned content-session fixture.** ✅ rext `6ee2d48`
  The public-anchoring (inner-join `directus.simulations` on `private=false AND tenant_id IS NULL
  AND status='published'`) + non-manager-played + per-(type × modality × passed/not-passed) selection
  query (a reusable Go builder + documented), authoring-time capture (via the read-only postgres MCP)
  of each pinned session's NON-PII structural projection (score, skill node-ids+scores, criterion
  shape, actor/interaction counts, modality), and a checked-in, anonymized, source-pinned
  `content-sessions` fixture + the blueprint `ContentSession`/`ContentProduct` model + loader.
  Assessment set = 2 voice + 1 code + 1 document; + training passed/not-passed; + hiring; + interview.

- [x] **S2 — Anonymization layer + ContentStorySeeder core fan-out.** ✅ rext `3a2b61e`
  The anonymize-by-construction transform (structured-id re-key/re-tenant, token regen, timestamp
  shift, minted anonymized player seat → owner-is-player-vantage, free-text SYNTHESIZED never copied)
  + the seeder reconstructing per pinned session: `jobsimulation.sessions` + the
  `public.local_jobsimulation_sessions` MIRROR + `validation_attempt_results` + `_skill_results` +
  `_criterion_results` + `validation_check_results`, passed + not-passed, all G14-valid enums,
  idempotent, PerStackIsolated + audited, registered in the seeder set. Closure stays green
  (resolve-or-drop skill node-ids, never fabricate).

- [x] **S3 — Modality substrate: transcript + CODE + DOCUMENT + INTERVIEW reports.** ✅ rext `bea0324`
  `actors` (PII names → anon identity) + `interactions` transcript (action_type real, action_payload
  synthesized); CODE modality (CodeSubmission / code criterion input_format); DOCUMENT modality
  (CollaborativeAsset / storage_upload); INTERVIEW full `interview_extraction_results.user_report` +
  `manager_report` JSON (believable report, not the minimal `{}` the SuccessionSeeder writes).

- [x] **S4 — Interview render flags + manifest source-pin + docs.** ✅ rext `4e4add6`
  Enable `flag_interview_{player,manager}_report` in the demo (a sha-pinned `demopatch` on the demo's
  own ephemeral clone, or a demo PostHog bootstrap — D3); fold the pinned prod sources into the
  `seed-generation-manifest.yaml` projection (source-pin contract, deterministic reseed); AMEND
  `corpus/ops/safety.md` Part 3 to the honest posture (anonymized-real, VPN/tailnet-scoped,
  source-pinned bounded exception); author `corpus/ops/demo/session-clone-spec.md` (the deliverable).

## M232: Final Review — Completeness Ledger (section)

**Done (Fate 1):** S1 sourcing + source-pinned fixture · S2 ContentStorySeeder core fan-out (session + mirror co-write
+ attempt/skill/criterion/check, G14 enums, re-tenant, non-manager) · S3 modality substrate (transcript + code +
document + interview report) · S4 interview flag-gate demopatches + manifest source-pin projection + `safety.md` §3.8
amendment + `session-clone-spec.md`. **+ the REWORK (user decision 2026-07-19): synthesize → COPY-THE-REAL-CONTENT +
best-effort scrub** — new `scrub/` pkg, `cmd/content-capture` (prod read-only, streams prod→scrub→fixture, raw never
in agent context, ran once → 9 real scrubbed sessions), `contentsession/` model+embed, replay seeder; docs flipped to
the honest copy+scrub / residual-risk-accepted / VPN-scoped posture. Guardrail tests: CopiesRealContent (byte-copied),
PlaceholdersFilled (no source name/org survives), NoStructuralPII (re-scans the shipped fixture), scrub (no
email/phone/url survives). rext tag `playbill-m232-sections-copyreal`. All ✅.

**Confirmed-covered (Fate 2):** INTERVIEW report exact plan-section render fidelity → **M235** (prove-it-lands — its
coverage gate owns "every session×action lands on a non-empty result page"; M232 seeds the REAL copied report).

**Annotated (Fate 3):** dedicated minted per-session player seats → **M234** (D2, handoff; reuse-a-member baseline
renders). The **14 pre-existing demo-stack test failures** (SSR + studio/pubweb `urls.ts` pins drifted vs the newer
`stack-demo` clone; + cockpit/host/purge) → **v2.5 release close** (D8) — subsystems M232 never touched; M232's own
files match `stack-demo` byte-for-byte + pass. **⚠ REPEAT pattern:** this same drifted-clone test-failure class
recurred at v2.4 close-release — worth a dedicated re-anchor bugfix at the v2.5 release close.

**Dropped:** none. **Escape-hatch:** none.

**Verdict:** M232 delivers the copy-real session-clone seeder the user asked for. The residual re-identification risk
is documented + **data-controller-accepted** (`safety.md` §3.8), best-effort-scrubbed, VPN/tailnet-scoped. Deferral
audit YELLOW (the recurring 14-test carry flagged for release-close re-anchor; 0 escape-hatch). KB-fidelity GREEN.

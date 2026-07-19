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

- [ ] **S3 — Modality substrate: transcript + CODE + DOCUMENT + INTERVIEW reports.**
  `actors` (PII names → anon identity) + `interactions` transcript (action_type real, action_payload
  synthesized); CODE modality (CodeSubmission / code criterion input_format); DOCUMENT modality
  (CollaborativeAsset / storage_upload); INTERVIEW full `interview_extraction_results.user_report` +
  `manager_report` JSON (believable report, not the minimal `{}` the SuccessionSeeder writes).

- [ ] **S4 — Interview render flags + manifest source-pin + docs.**
  Enable `flag_interview_{player,manager}_report` in the demo (a sha-pinned `demopatch` on the demo's
  own ephemeral clone, or a demo PostHog bootstrap — D3); fold the pinned prod sources into the
  `seed-generation-manifest.yaml` projection (source-pin contract, deterministic reseed); AMEND
  `corpus/ops/safety.md` Part 3 to the honest posture (anonymized-real, VPN/tailnet-scoped,
  source-pinned bounded exception); author `corpus/ops/demo/session-clone-spec.md` (the deliverable).

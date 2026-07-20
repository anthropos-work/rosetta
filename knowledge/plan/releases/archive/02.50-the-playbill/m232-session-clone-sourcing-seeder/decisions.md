# M232 — Decisions

## D1 (S1) — Anonymize BY CONSTRUCTION (the core safety decision)
The seeder does NOT copy customer free-text — it copies only the NON-PII structural skeleton (source
session-id pin, real public sim_id, sim_type, modality, score, pass/fail, duration, actor/interaction
counts) and SYNTHESIZES every free-text facet (LLM feedback, transcript, submission, actor names,
interview report). Per `content-stories-routes.md` §3.3, M232 owns the synthesize-vs-redact choice and
the brief leans synthesize. Consequence: the checked-in fixture is **provably PII-free**, the strongest
posture, and the `safety.md` Part 3 amendment (S4) records a NARROW, honest exception — real session
STRUCTURES sourced (scores/shape), synthesized content, source-pinned, VPN/tailnet-scoped.

## D2 (S1) — Owner = an existing player-vantage population member (REUSE, not mint)
Open question ("reuse hero seats or mint per-session anonymized player seats — brief leans mint, each
must map to a real seeded `public.users` row") RESOLVED toward **reuse a distinct non-hero player member**
of the target org. Rationale: reuse trivially satisfies "maps to a real seeded users row" + "never a
manager seat" (owner-is-player BY CONSTRUCTION), needs zero extra user-minting (no re-implementing the
UsersSeeder identity/casbin path), and a real member owning a real content-story session is MORE
believable than a thin single-session minted seat. Non-hero slots keep hero dashboards clean.
- Surfaced for M234: if the cockpit's "become this session's player" UX needs a dedicated seat per
  session, mint then — but the reuse baseline already renders. (Handoff noted; not a deferral of M232 scope.)

## D3 (S1) — Target org = the first NON-HIRING EffectiveStory
Content-story sessions land in the first non-hiring story org (has player members). They appear in that
org's activity (believable — a few extra sessions of extreme score don't skew a 220-member org). A
dedicated isolated content-story org is an M234 open question (cleaner cockpit grouping), not needed for
M232 to render.

## D4 (S1) — Fixture is a go:embed'd curated artifact (not per-stack config)
The exhibit set is a FIXED, source-pinned, code-owned artifact (like the AI-readiness prompts), so it is
go:embed'd in the `contentsession` package rather than declared per-stack in the blueprint. The seeder
loads `Embedded()`; M233 projects the pins from the same embed into `seed-generation-manifest.yaml`
(single source, honesty-gateable).

## D5 (S1) — Sourcing is authoring-time only; the seeder never touches prod
`sourcing.go` BUILDS the selection SQL (the reproducible record of the public-anchoring + non-manager +
per-cell selection, D6) but never RUNS it. Sourcing is an authoring-time activity against the read-only
postgres MCP (`db-access.md`); the seeder is fully offline (reads only the embed + the stack's own
replayed taxonomy). This keeps a demo box prod-access-free and the read boundary honored.

## D6 (S3) — Transcript uses only DB-valid action_types; interview report plan-shaped
- The `interactions.action_type` DB enum admits ONLY `email` + `call` (a COPY bypasses Ent, so an
  invalid value would insert-but-be-invisible — the G14 class). VOICE→`call`, everything else→`email`.
  Document uploads ride as `email` interactions with `attachment_refs` (NOT `storage_upload`, which is a
  proto-only constant the DB rejects). Confirmed against the live value distribution (email 309744 / call
  7382).
- CODE grades via `collaborative_asset` (the editor-content diff) — there is NO `code` input_format enum
  value; the `code_submissions` Judge0 row is a separate substrate. So a code clone writes BOTH a
  code_submission (the run) and a collaborative_asset (the graded editor content).
- **INTERVIEW report is PLAN-DRIVEN → Fate-2, covered by M235.** The next-web `AISimulationResultContainer`
  iterates a SEPARATE CMS `ExtractionPlan` and looks up `results[sectionId]`; FULL render fidelity needs the
  seeded section-ids to match the replayed plan for the one public interview sim. M232 seeds a
  structurally-valid, believable `{"results":{...}}` envelope (score_grade section + narrative + manager
  attention_points) — insertable + renders the header quality badge. Exact plan-section-id alignment is a
  render-coverage refinement, and **M235 "prove-it-lands" already owns it**: its exit gate is "every
  (session × action) lands on a NON-EMPTY result page for both vantages", and its iteration protocol
  "triages each blank landing to its true read-model; fix in seeder/manifest or route to a demo-patch". So
  the interview page's plan-section match is exactly what M235's iteration surfaces + fixes — Fate-2 (no
  plan edit needed, no M232 scope deferred). (The two interview PostHog flags are enabled in S4.)

## D7 (S4) — Interview flags via a two-part sha-pinned demopatch (not a PostHog bootstrap)
The interview flags are enabled by TWO sha-pinned demopatches (the M219 `next-web-aireadiness-flag-gate`
twin), NOT a PostHog bootstrap. Rationale (agent-verified): a demo bakes no PostHog key, so `posthog.init()`
never runs and a `bootstrap.featureFlags` block is unreachable; baking a key would query a real/absent PostHog
(still undefined/false) AND re-open the third-party egress `next-web-no-thirdparty` closed (safety §3.6). The
flags are compiled-in client source → the demopatch ladder's sanctioned case. TWO patches because the two
flags feed ONE combined gate per file, in TWO shared `packages/ui` files: `next-web-interview-flag-container`
(the FETCH gate, `AISimulationResultContainer.tsx`) + `next-web-interview-flag-result` (the RENDER gate,
`AISimulationResult.tsx`). Single-line anchors (no chaining — each own file). Wired into `up-injected.sh`'s
BOTH frontend builds (packages/ui is shared → apps/web + apps/hiring) + the patchset fingerprint + LIFO revert.
Verified: load + apply + idempotent-reapply + revert-clean against a git working tree; the pin test
(`test_interview_flag_patch_m232.py`, 11 cases) fences shape + wiring + a live-anchor drift pin.

## D8 (S4) — demo-stack suite: 4 self-caused (FIXED) + 14 pre-existing (NOT M232 scope)
`python3 -m unittest discover -s demo-stack/tests` = 683 tests, 18 failures. Triaged:
- **4 SELF-CAUSED → FIXED.** `test_frontend_build` tag-guard tests: my adding the 2 interview manifests to
  `next_web_patchset_fp` correctly changed the `demo.patchset` fingerprint, so the guard rebuilt an image the
  test had labeled with the OLD fingerprint. The tag-guard behaviour is CORRECT (a new patch must invalidate a
  cached image — the §5-bis "applied != shipped" fence); the FIX was to sync the test harness's hardcoded
  `_PATCHSET_MANIFESTS` list (add the 2 interview ids). `test_frontend_build` now GREEN (89/89).
- **14 PRE-EXISTING, out of M232 scope** (subsystems M232 never touches): SSR + studio/pubweb urls.ts patches'
  pinned `pre_sha256` drifted vs the current `stack-demo/next-web-app` clone (it's at a NEWER next-web tag; e.g.
  studio pin `0d4c37902…` vs stack-demo `d92fa701…`; the tests say "re-anchor needed #8") = 6
  (`test_ssr_origin_chain` 4 + `test_demopatch` 2); plus `test_cockpit` academy-link/overlay 6,
  `test_host_prereqs_m215` 1, `test_purge` 1 — all unrelated pre-existing failures on this box.
- **My work is CLEAN.** Interview manifests match stack-demo BYTE-FOR-BYTE (verified), so they PASS the
  live-clone checks; my new pin test (11) + aireadiness (9) + frontend_build (89) all green; 0 interview-related
  failures. The 14 pre-existing failures (demo-stack hygiene: the urls.ts/SSR re-anchor + the cockpit/host/purge
  reds) are a Fate-3 candidate for the v2.5 release close, not M232 work. (Self-healing anchor, demopatch-spec
  §6: the drifted patches still APPLY at demo-up; only the whole-file sha baseline drifted.)

## D9 (REWORK 2026-07-19) — COPY the real content + best-effort scrub (SUPERSEDES D1's anonymize-by-construction)
User decision, explicit: the anonymize-by-construction / synthesize-free-text design (D1) was NOT what was
asked. The interesting part of a played session IS its free-text (real conversation, real feedback, real
submission), which synthesis fabricated. **REWORKED the CONTENT layer to COPY the real free-text + scrub PII
where detectable.** The INFRASTRUCTURE is unchanged (fan-out, mirror co-write, G14 enums, source-pin,
public-anchored D6, non-manager, re-tenant, the two interview flag-gate demopatches, the manifest projection).
- **New `scrub` package** (tested): redact emails/phones/urls + replace known names/org → placeholders.
- **New `cmd/content-capture`**: reads prod **read-only** (marco_read via ~/.pgpass over Tailscale, `SET`
  read-only, SELECTs only), COPIES each pinned session's real fan-out content, SCRUBS it, writes the checked-in
  `contentsession/fixture/content/<key>.json`. Raw content streams prod→scrub→fixture — never enters an agent's
  context (prints counts only). Ran once → 9 real scrubbed blobs.
- **New `contentsession/content.go`**: the captured-content model + go:embed loader.
- **Reworked seeder**: `content_stories_{write,modality}.go` now REPLAY the copied content (fill
  `<<ACTOR_i>>`/`<<ORG>>` placeholders with the demo persona/org), not synthesize. Copies the REAL skill node-ids
  the candidate was assessed on. `resolveTaxonomyRefs` draw removed (node-ids copied, not drawn).
- **HONEST posture** (docs updated): safety.md §3.8 + session-clone-spec.md flipped from "provably PII-free by
  construction" to "real content COPIED, best-effort scrubbed, NOT guaranteed clean — residual re-identification
  risk ACCEPTED by the data-controller (2026-07-19), VPN/tailnet-scoped as the control." manifest
  anonymization_transform, README, CLAUDE.md all reworded.
- **Tests**: scrub unit (no email/phone/url/name/org survives); `TestEmbeddedContent_NoStructuralPII` (fixture
  cleanliness re-scan); `TestContentStorySeeder_CopiesRealContent` (seeded free-text == captured, filled,
  byte-for-byte → proves COPIED not synthesized); `PlaceholdersFilled`; kept fan-out/enum/mirror/determinism/
  public-anchored/non-manager. D1's synth-specific tests (NoTaxonomyDropsSkills, CriterionInputFormat) removed.
- Zero platform-repo edits still holds (interview flags stay demopatches; the prod read is authoring-time read-only).

# M222 — Progress

_Section checklist, derived from `overview.md` § Scope.In. To be worked by `/developer-kit:build-milestone`._

## Sections

- [x] **S0 — KB-fidelity gate** (pre-milestone) — the blind area IS the gate; resolved by authoring
      `corpus/services/hiring.md` (S1) from the S3 render-probe's live-code trace.
- [x] **S1 — Author `corpus/services/hiring.md`** (the BLIND AREA — the hiring model, before any code) — registered
      from `next-web-app.md`.
- [x] **S2 — The `is_hiring` gate thread** (blueprint `StoryOrg.IsHiring` field + `narrative: hiring` +
      `ResolvedStory.IsHiringOrg()` + `org.go` thread + the deterministic hiring `OrgID` reservation + unit tests)
- [x] **S3 — Render-probe BA-1/BA-2/BA-3/BA-6** (throwaway hand-seed vs the live dockerized `apps/web` on `billion`)
      — **GO** (see `decisions.md` D1–D4 + `spec-notes.md`)
- [x] **S4 — Record the GO/NO-GO + the exact seeder-output contract** — **GO**; the score is the MIRROR table
      `public.local_jobsimulation_sessions`, NOT `jobsimulation.sessions` (D2). NOT `apps/hiring`-only (D1).
- [x] **S5 — Registrations** in `seeding-spec.md`/`stories-spec.md`/`README.md` (the 4th story + the `is_hiring` gate)

## M222: Final Review

_Close-milestone review (2026-07-15). The milestone shipped clean — the only fixes were one doc-fidelity tighten
and the required Phase-2c adversarial record; all gates ran and passed._

### Scope
- [x] S0–S5 all checked; no dropped/forgotten items. The single Fate-3 (job_position replay → M223) is recorded in
      D4 + M223 `overview.md` Scope.In #4. Deferral audit **GREEN** (`audit-deferrals/deferral-audit-2026-07-15-m222-close.md`).

### Code Quality
- [x] [verified] `go vet` clean · `gofmt` clean · `-race` clean (seeders/blueprint/manifest). Gate code is
      consistent with the `aiReadinessNarrative` sibling pattern; no dead code, no boundary issues. The `org.go`
      thread is a single load-bearing value (`st.IsHiringOrg()`, was hardcoded `false`).

### Adversarial (Phase 2c)
- [x] 3 scenarios recorded in `decisions.md` § Adversarial review (manifest omitempty byte-identity; the dual-signal
      OR; exact-match narrative). The highest-value one (omitempty preset drift) is already test-covered by
      `manifest/hiring_test.go::TestBuildPopulation_ProjectsIsHiring` (asserts exactly one `is_hiring:` key).

### Documentation
- [x] [fixed] `hiring.md` `useGetClerkOrganization` code quote: aligned the optional-chaining to the real source
      (`organization?.publicMetadata?.…`) — every other file:line claim spot-checked GREEN against the READ-ONLY
      platform clones (resolver_queries.go:1088/1089/1134, intelligence.go:1692/1728-1751/1801,
      local_jobsimulation_session.go:52, useNavbarSections.tsx:301, template.tsx:90, FreeTrialContainer.tsx:29).
- [x] `hiring.md` discoverable (linked from `next-web-app.md`, `seeding-spec.md`, `stories-spec.md`); follows TEMPLATE.

### Tests & Benchmarks
- [x] Full stack-seeding suite: **965 pass / 0 fail / 0 skip** (13 packages, `-count=1`). Flake gate **5/5** green
      (random `-shuffle`). Gate test `TestOrgSeeder_IsHiringGate` present + GREEN + **RED-provable** (reverting
      `org.go` to hardcoded `false` fails it: Recruitco/Talentworks want true, got false).
- [x] Handbook test-count reconciled: README "832 test funcs / 83 files / 13 packages" = 830 `func Test` + 2
      `func Fuzz` — accurate, no drift.

### Decision Triage
- [x] D1 (GO), D2 (mirror-table contract), D3 (dual-write), D4 (job_position drop) — the knowledge-worthy content
      was blended into `corpus/services/hiring.md` + `stories-spec.md` + `seeding-spec.md` at build time (verified
      accurate); D1/D4 carry explicit `M222 D#` back-references. Full records stay in `decisions.md` (archive).

## M222: Completeness Ledger (section variant)

_Every scope item placed into exactly one three-fate category. Cross-checked against overview.md In/Out,
progress.md checkboxes, spec-notes.md promises, decisions.md, and code TODOs (none)._

### Done (Fate 1) — landed complete in M222
- **S1 · `corpus/services/hiring.md`** — the BLIND-AREA hiring READ-model (dual-write, the mirror-table
  read-path, the seeder-output write-set, the isEnterprise divergence). file:line-verified against the clones.
- **S2 · the `is_hiring` gate** — `blueprint.StoryOrg.IsHiring` + `HiringNarrative="hiring"` +
  `ResolvedStory.IsHiringOrg()` + `OrgSeeder` threads `st.IsHiringOrg()` (was hardcoded `false`) + reserved
  `HiringOrgID()` + `manifest.Org.IsHiring` (omitempty) — with RED-provable tests.
- **S3/S4 · render-probe + GO/NO-GO** — BA-1/BA-2/BA-3/BA-6 answered on the live `billion` substrate; **GO**
  (D1), the mirror-table seeder-output contract (D2), the dual-write (D3). BA-3 escalation trigger refuted.
- **S5 · registrations** — `next-web-app.md` + `seeding-spec.md` + `stories-spec.md` + rext README.
- **S0 · KB-fidelity gate** — resolved by authoring hiring.md from the S3 live trace.

### Confirmed-covered (Fate 2) — owned by a named downstream milestone of this release (no plan edit)
- The full 50-person seed + the candidate-assessment funnel → **M223** (overview.md `In`).
- The cockpit heroes (1 manager + 2 candidates) + Clerkenstein `publicMetadata.isHiring` wiring → **M224**.
- Any latency work → **M226** (reported not gated until then).

### Annotated (Fate 3) — attached to a milestone at close (plan edited)
- **The `directus.job_position` snapshot replay is DROPPED from M223 Scope.In** (D4): the captured snapshot has
  0 `job_position` rows and the scoreboard never reads the entity — the 5 "positions" ARE 5 real captured
  `SIMULATION_TYPE_HIRING` sims. Applied to `m223-casting-the-ensemble/overview.md` (Scope.In #4 struck through +
  frontmatter `delivers:` updated) at build time; audited GREEN at close.

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch)
- **None.** No item leaves v2.4; no cross-release punt; no user sign-off required.

**Verdict:** all In-scope items delivered as Fate 1; Out-scope items are Fate-2 (sequential release graph); one
applied Fate-3 (job_position → M223). **Zero escape-hatch deferrals → proceed to merge, no sign-off needed.**

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

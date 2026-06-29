---
milestone: M50
slug: content-seeding-fill
version: v1.10b "fit-up"
milestone_shape: iterative
status: planned
created: 2026-06-29
last_updated: 2026-06-29
complexity: large
exit_gate: "On a COLD reset-to-seed demo, both Maya (employee) and Dan (manager) render fully-populated across every annotation-listed surface, proven by a re-run of the M42 semantic coverage gate (employee + manager), 0 prod-eject escapes."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
delivers: corpus/ops/demo/{profile-completeness-spec,stories-spec}.md updates (new seeders + backfills)
issues: Demo Content + Heroes-as-Maya + Heroes-as-manager (annotation.md)
---

# M50 — Content & seeding fill

## Goal
On a cold reset-to-seed demo (built on M47+M49's clean bring-up), the **existing** heroes — Maya (employee) and
Dan (manager) — read as fully-populated people across every surface the field review found empty.

## Exit gate
On a **COLD** reset-to-seed demo, both vantages render populated across: `/profile/activities` (activity + xp +
skill-paths completed) and `/home` in-progress; `/library/skill-paths`; the Workforce **Growth / Verification /
Talent** tabs (incl. **languages** + **certifications**); `/enterprise/assignments`; `/enterprise/members`
(**location / join-date / last-activity**); and the **hero academy link + a non-anonymous academy session** —
**proven by a re-run of the M42 semantic coverage gate** (employee + manager), **0 prod-eject escapes**.

## Why iterative (not section)
The real root-causes are **hypotheses until observed on the clean bring-up**: some gaps need a **new** seeder
(`MemberLanguagesSeeder` does not exist), some need **backfills** (member location/join-date/last-activity), and
some may simply **vanish** once set-dress actually runs (several "empty" surfaces are plausibly downstream of the
demo-up #7 abort, which skipped set-dress entirely). The fix list emerges from the first observation pass — a fixed
`In:` checklist now would be speculative.

## Iteration protocol
`corpus/ops/demo/coverage-protocol.md` — the M42 Playwright **semantic believability** gate (real seeded content +
per-section cardinality + persona self-consistency + 0 prod-eject escapes). Each tik: observe → fix a cluster →
re-measure both vantages.

## Candidate fix surface (emergent — NOT fixed)
- **Hero academy link + session** (annotation: no menu link; anonymous session if accessed directly) — pairs with
  M49 #5 (clone) + the academy-content decision below.
- **`/profile/activities` + `/home`** — `HeroActivitySeeder` / `ActivitySeeder` output resolving against the
  replayed content (the M10 content-ref linkage).
- **`/library/skill-paths`** — confirm the recaptured Directus content replays (likely M47-fixed).
- **Workforce Growth/Verification** — target-roles + population-evidence aggregates resolving against the replayed
  taxonomy.
- **Workforce Talent** — a **new `MemberLanguagesSeeder`**; certifications coverage (the existing
  `CertificatesSeeder` writing enough, role-coherent rows).
- **`/enterprise/assignments`** — assignment resource-refs resolving (M10 linkage).
- **`/enterprise/members`** — backfill `location` / backdated `created_at` join-date / `last_activity` roll-up.

## Open questions (resolve during build)
- **ant-academy course content** — shared-Directus replay vs a dedicated academy snapshot surface (so the catalog
  isn't `0 chapters / 0 skill-paths`). Decide here; the AI-keys decision (M49) gates the academy AI chat.
- Which "empty" surfaces were demo-up-#7 artifacts vs genuine seeder gaps — settled by the first observation pass.

## Depends on
**M49** (the clean bring-up to diagnose against) + **M48** (current code). **Parallel with:** none (monopolizes the
single demo stack — the 1-demo constraint).

## Re-scope trigger
A surface that cannot be populated without a **platform edit** → **escalate** (`unimplementable-without-platform-edit`),
never edit the platform.

## KB dependencies (read as contract)
- `corpus/ops/demo/coverage-protocol.md` (the gate), `corpus/ops/demo/stories-spec.md` (the 7-table chain),
  `corpus/ops/demo/profile-completeness-spec.md` (the density rubric), `corpus/ops/seeding-spec.md`.

## Delivers
- **→ rosetta-extensions:** the new/fixed seeders (e.g. `MemberLanguagesSeeder`, member-field backfills), tagged
  `fit-up-m50`.
- **→ rosetta:** `profile-completeness-spec.md` + `stories-spec.md` updates (the new seeders + backfills + the
  academy-content decision).

# iter-27 — decisions

## D1 — REVISES iter-26 D1: the 3 ai-readiness sub-renders are HARNESS locator mismatches, NOT seed gaps
iter-26 characterized all 3 as seed/wiring DATA gaps (route → 3 rext stack-seeding fixes). This iter's
re-survey against billion's LIVE DB + the v1.341.0 read paths + the actual next-web UI **falsifies that on all
three**: the data is fully seeded and rendering correctly; the harness locators assert text/structure the
shipped v1.341.0 UI does not produce. Each was confirmed empirically:

- **(a) manager-dashboard.UC1 byTeam.** Org C (Vertex Logistics, `3a7a47f1-…`) HAS **13 tags + 42
  membership_tags**; a reproduction of `aggregateByTeam`'s exact input (started-AND-tagged members) returns
  **12 teams** (Product 7, Design 6, Engineering 6, …). The read path is sound: `queryMemberTags`
  (`members.go:485`) joins `membership_tags ⋈ tags` — exactly what `TagsSeeder` writes (all 40 memberships
  matched, 0 id-mismatch); `loadMembers→hydrateMembers` populates `mem.Tags`; both `buildLiveResponse` and
  `buildResponseFromSnapshots` feed tagged members to `aggregateByTeam`. The UI (`SnapshotTab.tsx::ByTagTable`)
  renders the section **unconditionally**, titled **"AI Readiness by Team"** (`i18n aiReadiness.byTag.title`;
  "byTag" is only the internal key — `matrix.toggle.byTag` also renders "By Team"). The harness locator asserted
  `/AI Readiness by Tag/i` — "Tag" is not a substring of "Team", so it missed a correctly-rendered section.

- **(b) manager-dashboard.UC2 interview panel = 24 chars.** The seeded
  `jobsimulation.interview_aggregated_reports` row for Org C is **rich** (8396 chars, session_count 31,
  narrative.patterns=4 / unexpected=4 / insights=4 / catalog_kpis=4 / catalog_charts=2). `queryInterviewAggregatedReports`
  (`how_we_measure_v2.go:596`) is keyed by **organization_id ONLY** — its own v1.341.0 comment says it is
  "NOT gated through ai_readiness_sims, whose configured interview sim_ref can differ" — which **refutes
  iter-26's sim_ref-alignment hypothesis** outright. The 4 findings blocks RENDER (the spec's own assertions at
  lines 82–87 PASS; Playwright fails at the first failure and failed at line 97, so it got past them). The
  failing `>900`-char check measured `interviewBreakdown().innerText()`, which scopes to the **24-char heading
  `<span>` "AI Interview — breakdown"**; the findings render as SIBLINGS of that heading inside the
  `openStep===3 && interview` card `<div>` (`HowWeMeasureTab.tsx:1579`; the `INTERVIEW_GROUPS.map` findings at
  :1734 are inside the card), so measuring the heading saw nothing beneath it.

- **(c) member-funnel.UC2 deadline.** The active cycle end_date (2026-08-22) is present + resolved. `DueDate.tsx`
  renders **"Due Aug 22 · N days left"** — the date via `toLocaleDateString(locale,{month:'short',day:'numeric'})`,
  **NO YEAR**. The M219-hardened `dueDate()` locator required a **4-digit year** in every date alternative
  (`\d{4}` in all four shapes), so it never matched the real render — even though its own inline example
  ("Due by Aug 12") was already year-less.

**Common class:** these are three facets of ONE finding — the Playthrough locators were authored against an
ASSUMED UI that differs from the correct, fully-seeded v1.341.0 render. 0 seed changes, 0 platform edits.

## D2 — fix: three harness-locator corrections in the Playthroughs e2e lib (+ one spec line)
`playthroughs/e2e/lib/ai-readiness-page.ts`:
- (a) `byTeam()`: `/AI Readiness by Tag/i` → `/AI Readiness by Team/i` (+ rationale comment).
- (b) added `interviewBreakdownPanel()` scoping to the breakdown CARD (the deepest div carrying both the
  "AI Interview — breakdown" heading AND a findings label, `.last()`); `aireadiness-manager-howwemeasure.spec.ts`
  line 92's `>900` measurement points at it instead of the heading. STRONGER: it now measures the real panel
  body, not a 24-char heading that could never clear the floor.
- (c) `dueDate()`: added two year-less date alternatives (`\w{3,9}\s+\d{1,2}` / `\d{1,2}\s+\w{3,9}`) — KEEPING
  the `due|deadline` anchor, so it does NOT regress to the bare-`by` toothlessness M219 fixed.

Harness runs from the LOCAL rext authoring copy against billion → **no billion re-pin, no re-seed**. tsc clean.
Validation: the 4 ai-readiness Playthroughs GREEN on billion (10 passed 2.3m); then the full 16-suite.

## D3 — Playthrough locators are billion-verifiable-only: assert against the ACTUAL demo render, never an assumed one
Lesson for the protocol: a Playwright locator that has NEVER been run against the real target can encode a
plausible-but-wrong string/structure and stay latent until the "prove on billion" run. All three here were of
that shape (Tag/Team, heading-vs-card scope, year-bearing-vs-year-less date). When a sub-render "fails" on a
demo whose DATA is provably present, root-cause the LOCATOR against the live UI+i18n BEFORE concluding a seed
gap — the empirical order is: DB has data? → read path produces it? → UI renders it (component + i18n)? →
locator matches the render? The failure was at the last rung for all three.

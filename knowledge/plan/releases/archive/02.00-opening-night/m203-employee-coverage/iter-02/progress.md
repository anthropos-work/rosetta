**Type:** tik (under TOK-01) — protocol: `corpus/ops/demo/playthroughs.md` §"The iteration protocol".

# iter-02 progress

- **Re-survey (Phase 1 Step 0):** confirmed pt-world is seeded on demo-1 (orgs Meridian Labs + Halcyon Retail);
  pt-employee (Pat Ellis) has **24 verified / 36 total** user_skills (≫ Spotlight's ≥2 threshold).
- **Recon (be-the-human probe):** /profile → h1 "Pat Ellis", tabs Career Profile / Skills / Activities. The
  Skills tab (/profile/skills) renders the verified-skill surface: "Verified Skills 8", "Role Skills 10",
  "All Skills 20", "Skill Gaps (5)", "Growth", 46 SVG charts, "Re-rate skills" + per-skill "Edit" controls.
- **Declared** profile.verified.UC1 + profile.growth.UC1 in `playthroughs/manifest/profile.yaml` (unique ids,
  both-way integrity, precondition `verified-skill` resolves against seed-worlds.yaml). ptvalidate: VALID
  (3 UCs, 3 live Playthroughs, 0 TODO).
- **Extended** `ProfilePage` with Skills-tab semantic accessors (skillsTab/openSkillsTab, verifiedSkillsStat,
  roleSkillsStat, skillGapsStat, skillCharts, growthStat, rerate/edit) — found-only antd landmarks (§5.2),
  no data-testid, no platform edit (P3).
- **Built** `profile-verified.spec.ts` (@pt:pt-profile-verified) + `profile-growth.spec.ts` (@pt:pt-profile-growth).
- **Ran** against demo-1: both PASS (16.5s / 16.8s). Full suite (3 Playthroughs + 12 unit) 15/15 pass — no regression.
- **Reconciled** (ptreport --gate no-regressions): **3/3 passing (100.0%), failing=0** — gate GREEN.
- Go tests + gofmt + go vet clean.

## Close — 2026-07-02

**Outcome:** +2 employee use cases green (profile.verified.UC1, profile.growth.UC1); employee coverage
3/3 declared passing; no-regressions gate GREEN on demo-1.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the milestone gate needs the full employee set — Skill Paths + AI Simulations still to land,
plus the 5-run reset-to-seed determinism proof; this iter advanced coverage, gate not yet reached).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (product-id reuse), D2 (self-eval routed forward), D3 (chart-presence assertion under P2) — see decisions.md
**Side-deliverables:** none
**Routes carried forward:**
  - iter-03 → `profile.self-evaluation.UC1` (the WRITE flow — the "Re-rate skills" / per-skill "Edit" self-rate
    path timed out under probe; needs its own iter to anchor the control + assert at the persist boundary, and
    the FIRST mutating flow → must prove reset-to-seed. Fate-3, handler PT-M203-iter03-self-eval.)
  - iter-04+ → Skill Paths (skill-paths.legacy.UC1), then AI Simulations (chat/interview/code) per TOK-01.
**Lessons:** the real antd /profile has near-zero a11y — section labels ("Verified Skills") render as
STAT-CARD text, not h1-h4 headings; the sanctioned anchor is scoped-within-<main> getByText on the visible
stat label. Charts assert as `main().locator('svg').first()` visible (presence, not geometry — P2). Pure-read
journeys need NO reset (no mutation) — the fastest greens, exactly as TOK-01 predicted.

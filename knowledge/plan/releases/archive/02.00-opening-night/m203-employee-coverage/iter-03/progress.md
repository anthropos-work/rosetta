**Type:** tik (under TOK-01) — protocol: `corpus/ops/demo/playthroughs.md` §"The iteration protocol".

# iter-03 progress

- **Re-survey + target substitution (Phase 1 Step 0):** probed the self-rate flow — the "Re-rate skills" /
  per-skill "Edit" controls' click is intercepted (the rate modal "Rate your skills as a DevOps Engineer" is
  in the DOM but stays hidden on click; a genuine drivability quirk). Re-reading the M203 gate showed the
  gate's Profile journey = "verified-skill chart + claimed-vs-verified gap + **work/education timeline**" —
  self-eval is a non-gate M201-corpus extra, while the timeline is the third GATE piece and uncovered.
  Substituted target → `profile.timeline.UC1` (still TOK-01, a deterministic read); re-routed self-eval.
- **Recon:** the Career Profile tab (default /profile view) renders "Work" (3 entries: DevOps Engineer @
  Meridian Labs "Jan 2024 - Present", SRE @ Drift Analytics "Jul 2021 - Jan 2024") + "Education" (University
  of Edinburgh "Feb 2015 - Feb 2019") — the M41 ProfileSeeder history.
- **Declared** profile.timeline.UC1; ptvalidate VALID (4 UCs, 4 live Playthroughs, 0 TODO).
- **Extended** ProfilePage with the Career-tab accessors (careerTab/openCareerTab, workSection, educationSection,
  timelineDateEntry — a tolerant "Mon YYYY - (Present|Mon YYYY)" range match scoped in <main>, found-only §5.2).
- **Built + ran** profile-timeline.spec.ts (@pt:pt-profile-timeline): PASS (16.3s).
- **Reconciled** (ptreport --gate no-regressions): **4/4 passing (100.0%), failing=0** — GREEN.
- Go tests + gofmt + vet clean.

## Close — 2026-07-02

**Outcome:** +1 employee use case green (profile.timeline.UC1); the full **Profile** gate journey now covered
(identity + verified-skill chart + claimed-vs-verified gap + work/education timeline). Employee coverage 4/4
declared passing; no-regressions GREEN on demo-1.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (the milestone gate still needs Skill Paths + AI Simulations journeys + the 5-run
reset-to-seed determinism proof; the Profile journey is complete but the full employee set is not yet green).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (target substitution: timeline over self-eval, gate-relevance), D2 (self-eval drivability quirk routed forward) — see decisions.md
**Side-deliverables:** none
**Routes carried forward:**
  - iter-04 → **Skill Paths** (skill-paths.legacy.UC1) per TOK-01 step 2 — the next gate journey. Needs a
    legacy path w/ an assessment sim in the replayed Directus catalog (M201 seed finding). Handler PT-M203-iter04-skillpaths.
  - later → `profile.self-evaluation.UC1` (non-gate M201 extra) — needs the rate-modal click-intercept solved
    (a real user-gesture / actionability fix, or assert at the interactive-state boundary). Handler PT-M203-selfeval.
  - later → AI Simulations (chat/interview/code) per TOK-01 step 3.
**Lessons:** when a routed-forward target proves finicky, re-check it against the GATE before sinking iter
budget — a gate-relevant sibling (the timeline) was the higher-value, lower-risk target hiding in plain sight.
The Career Profile tab is the DEFAULT /profile view (no tab click strictly needed, but openCareerTab is
idempotent-safe). Timeline entries assert via a tolerant date-range regex (presence of a real entry, P2).

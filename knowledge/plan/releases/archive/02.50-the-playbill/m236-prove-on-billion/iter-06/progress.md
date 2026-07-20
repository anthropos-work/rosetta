# iter-06 — tik: the residual 4

**Type:** tik

## Step 1–2 — the interview manager pairs were a FALSE FAIL

Probed `/enterprise/activity-dashboard/interviews/<simId>/<membershipId>`. It renders **correctly**:

```
AI Interviews / AI Readiness Interview / Nadia Ferrari
Answer Depth: In-depth · Adequate · Brief     Sort by: Start Date
Status      Start Date      End Date        AVG Time   Answer Depth
Completed   Mar 23, 2026    Mar 23, 2026    7min       -            View Report
```

A real player name, a real attempt row, a real action. **593 chars of genuine content.**

It is simply a **different surface** from the sim scoreboard: a breadcrumb over an attempts table with
*View Report*, and **no `<player>'s Results for <sim>` header** — which is exactly what the
`manager-dashboard` shape asserts on. So the shape graded a working page as broken.

Added `manager-interview` as the **fifth** calibrated shape, selected **by route** like the others.

This is the third shape-calibration finding of the milestone, and the pattern is now unmistakable: **every
time a "defect" turned out to be a shape mis-match, the page was fine and the assertion was wrong; every
time the assertion was right, the failure was real and had one clean cause.** Route-derived shape selection
is what keeps those two cases distinguishable.

## Step 3 — re-measure

**27 / 31 → 29 / 31.**

```
content-stories: LANDED 29 / 31
  simulation:        26/26   ← COMPLETE (13 player + 13 manager)
  skill-path-legacy:  3/4
  skill-path-new:     0/1

  x1  page.goto timeout (>180 s)  — one skill-path manager route
  x1  route rendered a not-found  — academy
```

**The simulation arm is complete**: every one of the 26 (session × action) pairs for the largest product
lands on real, non-empty content, live on `billion`, for both vantages.

## The residual 2 — characterized, not guessed

**(a) The skill-path manager route HANGS.** `/enterprise/activity-dashboard/skill-paths/df9d2142…/50e1bb5e…`
exceeds a **180 s** navigation timeout. Notably:
- It is **independent of the membership-id fix** (it failed identically before and after, and the id in the
  URL is now the corrected membership id).
- Its **sibling** skill-path manager route (`sp-genai-in-progress`) **passes**, so the surface works.
- The difference between them is size: the hanging one is the **completed 13-chapter** path; the passing one
  is a 3-chapter path at 45%.

That shape — one heavy instance hanging while a light sibling passes — is the arithmetic signature
`latency-budget.md` teaches to read as a **per-item fan-out**, not a broken route. Not fixed here: it is a
distinct investigation and the iter's declared scope was the residual triage.

**(b) The academy pair** is unchanged and needs the catalog + progress fill (`app/cmd/academy-seed` wired
into the cold bring-up) — known since iter-03, never in doubt.

## Close — 2026-07-20

**Outcome:** +2 pairs, and the **simulation arm is 26/26 complete**. Both remaining pairs are characterized
with named handlers; neither is mysterious.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y (5 tiks this session)** — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D1 (the interview manager view is a distinct surface — a false fail), D2 (the skill-path hang is a heavy-instance fan-out signature, characterized not fixed)
**Side-deliverables:** none.
**Routes carried forward:**
- **The skill-path manager hang** (1 pair) → **iter-07**, handler `SKILLPATH-M236-iter07-manager-hang`. Start from the heavy-vs-light sibling contrast and the `latency-budget.md` fan-out signature.
- **Academy** (1 pair) → handler `ACADEMY-M236-iterTBD-catalog-fill`: wire `app/cmd/academy-seed --user-id <the academy content-player owner> --fixture …` into the cold bring-up, then re-point the CTA from the anonymous `/library/<slug>` preview to the authed chapter route.
- **The p95 click→ACCESS measurement (HERO vantages only, per B2)** — a gate component not yet taken this milestone. → handler `LATENCY-M236-iterTBD-hero-p95`.
- **A final cold reset-to-seed reproduction** — the gate requires reproducibility, and the current stack has had its manifest re-exported and cockpit restarted in-place. → handler `REPRO-M236-iterTBD-cold-cycle`.
**Lessons:**
- **Three of this milestone's "defects" were wrong assertions, and one was a real bug with one clean cause.** The ratio is the lesson: when a gate is new, disbelieve the gate first. Probe the page before triaging the product.
- **Route-derived classification kept being the fix.** Every shape added was selected by URL, and every mis-grade came from inferring shape by content. In a system where surfaces differ but share a design language, the URL is the only cheap, stable discriminator.

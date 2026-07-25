---
iteration_type: tik
iter: 10
milestone: M254
status: closed-fixed
---

# M254 · iter-10 — overview

**Type:** tik · **Active strategy:** TOK-01 (cluster-per-tik live re-prove), **cluster 4** — the
mutating / seed-destroying serial tail: gate **(e)** studio builder Playthrough + the rest of gate **(h)**
(the live Playthroughs green).

## Step 0 — Re-survey (mandatory)
TOK-01's next-tik direction (cluster 4) named "the studio builder Playthrough + the Playthroughs green."
Re-survey confirms this target is current and un-absorbed:
- billion demo UP + peer-reachable (web 307 / hiring 307 / cockpit 200), 16 demo-1 containers (devops driver),
  rext clone pinned `july-jitter-m254-academy-nonode-hostrobust` (dfdd9bc) = authoring copy, `rext.tag` matches.
- iter-09 landed the pt-world reset + PROVED the browser suite runs (first 6 green incl a manager Playthrough)
  but did NOT complete the full 117-test suite (context saturated at cap). So (e)+(h)-Playthroughs are
  proven-to-run, not proven-green. Target still meaningful.

## Cluster / target identified
Run the **full Playthrough suite to completion** against the live billion demo and get a clean four-state
`ptreport` map. iter-09 ran a partial suite (6/117), which MUTATED the pt-world → for a **deterministic
full-suite verdict** re-reset pt-world on billion first (`--reset-only`, on-host via the devops driver), THEN
browse from this peer (`PT_HOST=billion.taildc510.ts.net PT_APP_SCHEME=https`). Assert:
- **(e)** the studio builder Playthroughs (`pt-studio-advanced-generate` + `pt-studio-guided-generate`) green.
- **(h)** the 16 live Playthroughs green (6 employee + 4 manager + 4 AI-readiness + hiring-recruiter +
  assignment-assign) + the locator/manifest specs; `ptreport --gate no-regressions` passes.

## Hypothesis
The demo is fresh (iter-08 cold reset-to-seed, autoverify green) + pt-world-seeded; iter-09 proved cohesion.
A clean re-reset + full run should land all `live` Playthroughs green (0 regressions).

## Expected lift
Gate parts **(e)** and the Playthrough half of **(h)** move to MET → milestone at effective gate-met
(6 hard-MET + e/h via Playthroughs + the 3 recorded coordinator-approved dispositions).

## Phase plan (protocol: playthroughs.md + verification.md)
1. Re-reset pt-world on billion (`--reset-only`, devops driver, on-host docker+DB) — foreground, ~1-2 min.
2. Launch the browser suite DETACHED from this peer (sentinel + log); FOREGROUND long-timeout poll loop
   (never background+yield); heartbeat the journal each cycle.
3. Reconcile `ptreport` four-state map; assert (e)+(h) green.
4. Any REAL failure → fix (rext / demopatch; rung-zero: tag + `git push --tags` origin, re-pin) or disposition.
5. Record the 3 coordinator-approved dispositions ((f)-FCP-p95, (c)-academy-durability, (g)-testhealth) into
   milestone `progress.md` + `decisions.md` + `carry-forward.md`.

## Escalation conditions
- A real Playthrough failure with no zero-platform-edit fix → user-blocker.
- The suite cannot complete (harness/plumbing) → route + re-scope.
- All 16 live green → state milestone at effective gate-met (gate-met exit).

## Acceptable close-no-lift outcomes
N/A expected — the suite is proven-to-run; this iter drives it to completion.

# M256 · iter-02 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 1.
Iter shape per `corpus/ops/demo/playthroughs.md` § The iteration protocol (steps 4 + 6).

## Phase 0b / 0d

0b **SKIPPED** — the milestone's standing pre-flight verdict (YELLOW, iter-01) is inherited; this tik
redirects into no new subsystem. 0d **SKIPPED** — a measurement-only iter wires no artifacts through a
generate/validate pipeline (the skill's explicit skip case).

## Phase 2 — the runs

Three consecutive `run-playthroughs.sh 2 --reset` invocations from the authoring copy, with the demo's own
`bin/` prepended to `PATH` (the M204 iter-05 gate-run prereq — the runner shells out to the *pinned*
`stackseed`, which is not on the login PATH). `PT_HOST=localhost`, `PT_APP_SCHEME=http`, `workers: 1`,
`retries: 0`. **No harness code was modified** — a baseline measured on changed code is not a baseline.

**132 tests per run — 18 browser Playthroughs + 114 unit specs. `132 passed` on all three runs; 0 flake,
0 red.** `ptreport` reconciled 18/18 `[PASS]`. So D-v28-3's batch-gate rule has nothing to escalate: the
consolidated red set is **empty**.

Hero login worked on every run, which independently retires iter-01 D5's autoverify alarm: the fake-FAPI
serves the browser fine.

## Phase 3 — the baseline (the denominator)

### Per-Playthrough, per run (seconds)

| Playthrough | run 1 (cold) | run 2 | run 3 | **median** |
|---|---:|---:|---:|---:|
| pt-activity-drilldown | 7.97 | 6.32 | 4.09 | **6.32** |
| pt-aireadiness-manager-howwemeasure | 4.72 | 4.69 | 4.40 | **4.69** |
| pt-assignment-assign | 8.76 | 4.03 | 4.39 | **4.39** |
| pt-profile-growth | 4.68 | 3.92 | 4.35 | **4.35** |
| pt-profile-verified | 4.70 | 3.92 | 4.05 | **4.05** |
| pt-aisim-chat-launch | 4.42 | 3.98 | 3.30 | **3.98** |
| pt-profile-timeline | 4.10 | 3.63 | 3.91 | **3.91** |
| pt-profile-identity | 3.63 | 2.95 | 3.50 | **3.50** |
| pt-skillpath-legacy | 3.77 | 3.02 | 3.16 | **3.16** |
| pt-workforce-roster | 3.35 | 2.79 | 2.98 | **2.98** |
| pt-workforce-succession | 2.90 | 2.68 | 2.71 | **2.71** |
| pt-aireadiness-member-done | 2.34 | 2.38 | 1.89 | **2.34** |
| pt-hiring-recruiter-compare | 3.20 | 2.10 | 1.90 | **2.10** |
| pt-aireadiness-manager-dashboard | **17.19** | 1.86 | 1.62 | **1.86** |
| pt-workforce-funnel | 2.50 | 1.75 | 1.85 | **1.85** |
| pt-aireadiness-member-progress | 1.84 | 1.77 | 1.78 | **1.78** |
| _pt-studio-guided-generate_ (studio lane) | 1.89 | 1.84 | 1.84 | **1.84** |
| _pt-studio-advanced-generate_ (studio lane) | 1.52 | 1.19 | 1.26 | **1.26** |

### The three figures D-v28-9 requires kept apart

| Figure | Value | Status |
|---|---:|---|
| **Median per-Playthrough, 16 non-studio** | **3.326 s** | **the GATED metric** — clause 1 target **≤ 2.628 s** (0.79×) |
| Median per-Playthrough, all 18 | 3.067 s | cross-check only (the studio pair drags it down) |
| Suite wall-clock (132 tests) | 85.4 / 56.6 / **54.4** s → median **56.6 s** | **REPORTED, not gated** (D-v28-12) |
| Studio lane | 1.26 s / 1.84 s | excluded from the median — **and see D6: it is not LLM-bound** |
| Mean per non-studio | 3.373 s | — |
| Slowest non-studio | 6.32 s (`pt-activity-drilldown`) | — |

**Pinned statistic (must be recomputed identically at the end or the ratio is meaningless):** the baseline is
the **median across the 16 non-studio Playthroughs of each Playthrough's median across 3 consecutive
`--reset` runs**, where run 1 is the first run after bring-up (**cold, and deliberately INCLUDED**). The
cross-check — median of the three per-run medians — is **3.228 s**, 3 % below the headline; either is
defensible, so the headline is pinned to remove the choice.

### ⚠️ Environment (the `latency-budget.md` rule — this number is not comparable to billion's)

Host `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM **9.70 GiB** (vs the documented 12 GB UI-tier
floor). Stack `demo-2`, offset 20000, **localhost / http**, brought up `--no-public-host` from a warm image
cache. rext authoring copy `main` @ `6ca8764`; the demo consumes pin `cockpit-deeplinks-v1`, and no
`playthroughs/e2e` runtime file differs between them.

**Per D-v28-12, none of these numbers may be quoted as comparable to billion's 228 s.** The absolute
billion re-measure is routed to M258.

## Phase 3 addendum — cold-start dominates run 1, which is why n=3 was mandatory

`pt-aireadiness-manager-dashboard` cost **17.19 s** on run 1 and **1.86 s / 1.62 s** after — a **9×** swing.
`pt-assignment-assign` 8.76 → 4.03, `pt-activity-drilldown` 7.97 → 4.09. Run 1's suite wall-clock is
**85.4 s** against 56.6 / 54.4 s. This is next-web first-render / route-compile warm-up, not test noise.

**Consequence for the gate:** the per-run medians are **3.935 / 2.989 / 3.228 s** — an **18 %** spread
between the coldest and warmest run, against a **21 %** target. A single-run measurement could pass or fail
clause 1 on warm-up alone. That is why the pinned statistic is a **per-test median across runs** and why the
protocol (n=3, cold run included) must be **identical** at both ends of the milestone. Recorded as D7.

## Close — 2026-07-28

**Outcome:** the denominator exists: **median per non-studio Playthrough = 3.326 s** (n=3, cold run
included, local `demo-2`), so clause 1's target is **≤ 2.628 s**; suite wall-clock median **56.6 s**
(reported). 18/18 green, 0 flake, empty red set. Two findings changed the milestone's aim: the studio lane
is **not** LLM-bound because `pt-studio-advanced-generate` is a **false green** (D6), and the median's real
driver is the **login handshake paid by every test**, not the `networkidle` inheritance TOK-01 named (D8).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (clause 1 needs ≤ 2.628 s; nothing has been optimised yet — this iter deliberately changed nothing. Clauses 2 and 3 unstarted.)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (first tik; it delivered its planned scope) — (4) user-blocker: n (see below) — (5) cap-reached: n (1 of 5 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D6 (`pt-studio-advanced-generate` is a false green; the "irreducible LLM lane" has no referent on this host), D7 (the measurement protocol is pinned because run-to-run spread ≈ the gate's own target), D8 (the median's driver is the per-test login, so iter-03 re-targets to `storageState` reuse).
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M256-studio-false-green` → **Fate 3, the clause-2 tik of this milestone.** Make
  `advancedDesignerRendered()` assert a post-generation landmark only, and prove it RED without a
  generation (D6). This is a clause-2 deliverable (a negative control), not a separate errand.
- `DOC-M256-llm-lane-premise` → **Fate 3, the same tik.** `corpus/ops/demo/playthroughs.md` and the M256
  overview both describe the advanced builder as running "to its completion boundary"; both need the
  correction once the spec is fixed, so the doc is rewritten **once**, against the fixed behaviour.
**Lessons:**
1. **A duration that is too GOOD is evidence, not luck.** A test documented as a 2–3 minute live-LLM
   round-trip finishing in 1.5 s was the only visible symptom of a false green. The three D-v28-9 figures
   were kept apart to *budget* the LLM lane; keeping them apart is what exposed that the lane was empty.
2. **Measure the cold run, then never compare across warmth.** Run 1 alone would have overstated the
   baseline by 18 % — conveniently, in the direction that makes a later "improvement" easy to claim.
3. **A localhost baseline understates a `networkidle` defect and overstates its fix.** `networkidle`
   resolves quickly against fast, sparse localhost requests (`page-object.ts` says so, and M244 proved the
   deadlock only over the tailnet). So the L1 lever is a **correctness** fix worth landing on its own terms,
   but its local *timing* payoff will be small — the honest reason to re-target iter-03 to the login.

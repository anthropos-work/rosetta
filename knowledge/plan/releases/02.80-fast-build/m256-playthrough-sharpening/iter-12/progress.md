# M256 · iter-12 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 4 · **Handler:** `NEGCTL-M256-cross-vantage`.

## Phase A — probe the contrast vantages

Four candidate contrast vantages driven live on `demo-2` before any control was written.

| contrast vantage | result | verdict |
|---|---|---|
| `pt-manager` (Org A, `narrative: ai-transformation`) on `/home` | bodyLen **3040**, 32 links, 21 buttons — **ALIVE**; `doneHeroTitle` **0**, `progressFunnelTitle` **0** | **USE** |
| `pt-manager` on `/ai-readiness` | bodyLen **2069**, 18 links, 14 buttons — **ALIVE**; `dashboardHeading` **0**, `methodHeading` **0** | **USE** |
| `pt-manager` on `/profile` | her own name renders; `Pat Ellis` **0** | **USE** |
| `pt-manager` on the **hiring** app | **ejected to `https://app.anthropos.work/login?redirect_url=…`**, bodyLen **162** | **REJECT** |
| `pt-manager`'s own profile *structural* stats | `verifiedSkillsStat` **1**, `skillCharts` **10**, `workSection` **1** | **REJECT** |

Two rejections are findings, not gaps:

1. **The hiring contrast EJECTS THE BROWSER TO PRODUCTION.** The absence would have been "true" while the
   browser was no longer in the demo at all — the iter-07 dead-page class in a new costume, and an
   out-of-demo escape besides. Recorded in the control file so nobody re-tries it.
2. **A STRUCTURAL final has no contrast vantage in this world.** The M44 profile-completeness seeder gives
   *every* member a career and skills, so there is no hero for whom `verifiedSkillsStat` / `skillCharts` /
   `workSection` is legitimately absent. This generalises: `mappedSkillsStat`, `coverageGap`,
   `memberRows().first()` and the rest of the workforce finals render for **any** populated org. **The
   cross-vantage mechanism can only discriminate an outcome that is ORG- or HERO-specific.** Pretending
   otherwise would have produced 9 controls that pass for any org — the vacuity iter-07 refuted.

## Phase B — what landed

**`negative-controls.spec.ts`** — a **third file class**, and its architecture is the load-bearing decision.
A control must NOT live inside the Playthrough it covers: a second login roughly doubles that Playthrough's
duration, and clause 1 gates the **median per Playthrough**, so paying 16 in-test controls would break the
speed clause to satisfy the honesty clause. The file therefore declares **no `@pt:` id** — it is not a
Playthrough, `ptreport` does not reconcile it, it never enters the median — and it batches by vantage so N
absences cost ONE login. (The milestone's own plan anticipated this: *"they are excluded from the timed p50
either way"*.)

**Every control asserts LIVENESS first** — a populated body and rendered nav, **polled** — then the absence.
That ordering is the iter-07 refutation encoded as a rule: a dead page satisfies every absence assertion,
so an absence is only evidence once the app is proven up.

**The coverage link is machine-checked, not claimed.** The control file declares which Playthroughs it
covers; the fence unions those links with the `@pt-negative-control:` lines on the mutating specs, and
**fails closed** on a link naming an id no Playthrough declares (phantom coverage) or on a token that does
not look like an id (a typo, or the tag written in prose).

**Two self-inflicted defects were caught by the guards as they were built** — both worth recording because
both are repeats of earlier iters' lessons:
- The fence's own explanatory prose contained the new tag **literally**, so the fence classified *itself* as
  a control spec and parsed its own sentences as an id list. That is iter-07's phantom-`@pt:`-in-a-comment
  defect one grammar later. Fixed structurally (`*.unit.spec.ts` can never be a control) **and** by
  convention (the tag is spelled apart in prose), plus the strict-token check that makes a silent skip
  impossible.
- The detection regex was the `/g` one, whose `lastIndex` persists across `.test()` calls — it would have
  matched every *other* file and silently halved the control set. Split into a non-global twin.

**A bounded interaction (side-deliverable).** `assignments-page.ts`'s exhausted-retries path ended in an
**unbounded** `waitFor`, so when all three attempts failed the run did not report *"submit never enabled"* —
it hung for the full **240 s** test budget and died naming that line instead of the real problem. Observed
twice, both times on the **first run of a batch** (iter-11 Phase C run 1, iter-12 Phase C run 1), green on
every warm run. The comment above it already stated the right intent; the missing `timeout` is exactly what
prevented it. Now bounded, so the spec's own assertion reports the failure.

**One control earned its keep on its first run.** It found that
`pt-aireadiness-manager-howwemeasure`'s step-name sub-assertions **match on a non-readiness org** — because
`/ai-readiness` without the feature renders a live **upsell** panel that names the very steps. Those
sub-assertions can therefore be satisfied by the not-enabled state. The control still covers that
Playthrough through `methodHeading()` (genuinely discriminating); the weak sub-assertions are routed
(`PT-M256-readiness-step-asserts`). *A negative control does not only confirm absence — it identifies
assertions that prove less than they appear to.*

## Phase C — re-measure, and the measurement itself became the finding

**Coverage (the planned metric): negative controls 8 → 13 of 24**, fence-computed
(8 self-declared + 5 via the control spec), the 11 uncovered named on every run. Suite **150 passed**,
`ptreport` **24/31 passing, 0 failing, 0 unimplementable** on runs 2–4.

**Flake: 1 in the batch** — `pt-assignment-assign` on run 1 (the 240 s unbounded-`waitFor` hang, since
bounded). Reported, not hidden.

### ⚠️ Clause 1 is NOT decidable at n=3 on this host, and the earlier "MET" readings were sampling noise

Six full-suite runs in this session, **same host, same code for the original 16 specs**:

| run | per-run median, all 22 non-studio | per-run median, ORIGINAL 16 only |
|---|---:|---:|
| iter-11 run 1 | 0.5701× | **0.5281×** |
| iter-11 run 2 | 0.8330× | 0.7259× |
| iter-11 run 3 | 0.7928× | 0.6866× |
| iter-12 run 2 | 0.9172× | 0.8103× |
| iter-12 run 3 | 1.1121× | **1.0762×** |
| iter-12 run 4 | 0.5967× | 0.5290× |
| **median (n=6)** | **0.8129×** | **0.7063×** |
| **spread** | **1.95×** | **2.04×** |

**The statistic varies by a factor of two on code that did not change.** The original-16 subset is the
control — no iter after 03 touched those specs — and it reads 0.528× on one run and 1.076× on another. There
is **no trend**: the most recent run (0.529×) matches the oldest (0.528×), with the extreme in between. So
this is not thermal degradation, it is **variance the pinned statistic does not absorb**: a median of 3
consecutive runs can land anywhere from ~0.53× to ~1.08× depending on which three runs it catches.

**Consequences, stated plainly:**

1. **Clause 1 cannot be declared MET on current evidence.** At n=6 the gated full-set figure is **0.8129×**
   against a `≤ 0.79×` gate — *outside* it. The flattering denominator (the original 16, 0.706×) is inside
   it, and choosing that one would be exactly the dishonesty this milestone has spent eleven iters refusing.
2. **The per-iter history needs re-reading.** iter-04's 0.5434×, iter-06's 0.6245×, iter-08's 0.5950×,
   iter-09's 0.5652×, iter-11's 0.6863× were each presented as a re-verification, and iters 06 and 08
   explicitly attributed their movement to variance. That attribution was right; what nobody tested was
   whether the *gate verdict itself* survives the variance. **It does not.**
3. **This is a release-decision, not an iter decision.** D-v28-12 re-cut clause 1 as a ratio against a
   same-stack baseline measured hours earlier; the ratio is only meaningful if the host is stable between
   the two measurements, and it is not. The remedies are all roadmap-level: raise n and report the spread;
   make the measurement **paired** (re-measure the baseline in the same batch); normalise within-run against
   an invariant leg; or move the measurement to a stable host (which is where D-v28-12 came from). **Escalated.**

*Also worth noting for whoever picks this up:* `pt-assignment-assign` is the single largest contributor to
the median's instability (4.60 s at iter-11, 6.44 s at iter-12 run 4) and it is the same Playthrough whose
retry ladder has now needed bounding twice. A speed lever aimed at it would improve both the median and its
variance.

## Close — 2026-07-28

**Outcome:** **negative controls 8 → 13 of 24**, via a third file class (cross-vantage controls that live
*outside* the Playthroughs so they cannot inflate the gated median) with a **machine-checked, fail-closed
coverage link**. Two candidate vantages were **rejected on measurement** — the hiring one ejects the browser
to production, and structural finals have no contrast vantage at all, which bounds the mechanism's
applicability honestly. And Phase C produced a finding larger than the iter's own scope: **clause 1's
verdict is not decidable at n=3 on this host** — the statistic spans **2.04×** on unchanged code, so the
"MET" readings from iter-03 onward were sampling noise and the n=6 figure (**0.8129×**) is *outside* the
gate.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 1 **NOT DECIDABLE at n=3** (n=6 full-set **0.8129×** vs `≤ 0.79×`; spread 2.04× on
unchanged code) — **escalated as a measurement-validity question against D-v28-12**; clause 2 mutating
**6/5 MET**, `blocked` **1/1 MET**, negative controls **13 of 24**; clause 3 verdict half **COMPLETE**,
landed half short (org-admin 2/4, onboarding 1/5); **D-v28-5 still unfixed**, still blocked behind
`FIX-M256-cockpit-manifest-drift`.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik landed its planned scope) — (3) re-scope: n (0 of 31 curated UCs `unimplementable`; the coverage clause is progressing) — (4) **user-blocker: y — a GATE CLAUSE has become undecidable on the available host, and every remedy (raise n · pair the baseline · normalise within-run · move to a stable host) is a change to D-v28-12, a release-level decision. Continuing to add coverage under a clause whose verdict is a coin flip would accumulate work against an unfalsifiable target** — (5) cap-reached: n (2nd tik) — (6) protocol-stop: n — Outcome: exit-4
**Decisions:** D53 (a control must live OUTSIDE the Playthrough it covers, because clause 1 gates the median and an in-test control would break the speed clause to satisfy the honesty clause), D54 (a control asserts LIVENESS before absence — polled, not read once), D55 (the coverage link is machine-checked and fail-closed in both directions), D56 (the cross-vantage mechanism discriminates only ORG- or HERO-specific outcomes; a structural final has no contrast vantage and must be sharpened instead), D57 (the hiring contrast vantage ejects to PRODUCTION and must never be used as a control), D58 (the fence harvested its own prose — the iter-07 phantom-comment class, one grammar later; fixed structurally AND by convention), D59 (a `/g` regex reused with `.test()` is stateful and would have halved the control set), D60 (**clause 1's verdict is not decidable at n=3 — the statistic spans 2.04× on unchanged code; the earlier MET readings were sampling noise, and the remedy is a release decision**).
**Side-deliverables:**
- `assignments-page.ts`'s last unbounded `waitFor` **bounded** — it converted a legible "submit never
  enabled" failure into an opaque 240 s hang, twice.
**Routes carried forward:**
- `NEGCTL-M256-cross-vantage` → **13 of 24; the mechanism is now bounded, which changes the shape of what
  remains.** The 11 uncovered split into **two classes with different costs**: (a) **9 structural** finals
  (`pt-workforce-*` ×4, `pt-activity-drilldown`, `pt-profile-{verified,growth,timeline}`,
  `pt-hiring-recruiter-compare`) which have **no** contrast vantage and must instead have their finals
  **sharpened to name real seeded data** — O(tests), and it makes those Playthroughs stronger regardless of
  the control; (b) **2 studio** which are blocked behind `FIX-M256-studio-false-green` (asserting a control
  on a known false green would certify it).
- `PT-M256-readiness-step-asserts` → **NEW.** `pt-aireadiness-manager-howwemeasure`'s `MANAGER_STEP_NAMES`
  asserts match page-wide and are satisfied by the **not-enabled upsell panel** on `/ai-readiness`. Re-scope
  them inside the method panel. Found by the control on its first run.
- `FENCE-M256-bounded-interaction` → **sharpened + now evidence-backed.** Four unbounded `waitFor` calls
  remain (`assignments-page.ts:62,115`, `activity-dashboard-page.ts:71`, `org-admin-page.ts:51`). None is
  inside a retry loop, so none is proven harmful — but the class has now cost two 240 s hangs. The fence is
  a source scan; **the harden pass is its natural home** (it fires next).
- `MEASURE-M256-clause1-sampling` → **NEW and ESCALATED (D60).** Clause 1's pinned statistic (median per
  Playthrough over 3 consecutive runs, against a baseline measured in a different batch) is not robust to
  this host's variance. Options, all release-level: raise n and publish the spread; make the measurement
  paired; normalise within-run against an invariant leg; or move it to a stable host. **`pt-assignment-assign`
  is the largest single contributor** and a speed lever aimed at it would cut both the median and its variance.
- Everything else from iter-11's list stands unchanged.
**Lessons:**
1. **A mechanism's LIMIT is as much a deliverable as its reach.** The cross-vantage control was routed as if
   it applied to all 16 remaining Playthroughs. Measurement says it applies to the org/hero-specific ones and
   cannot apply to structural ones at all. Knowing that converts a vague "16 to go" into two named classes
   with different costs — and it stops nine vacuous controls from being written.
2. **Guard your new grammar against your own documentation.** Two iters running, a tag introduced with a
   fence has been harvested from the prose explaining it. If a tool scans source for a marker, the tool's own
   file is source.
3. **Before trusting a ratio, measure the denominator's variance.** Eleven iters reported clause 1 as MET, and
   each report was honest about its own number. Nobody asked how much the number moves when nothing changes.
   It moves by 2×, which means the gate had not actually been passed — it had been sampled favourably. **A
   relative gate needs its noise floor published next to it, or it is not falsifiable.**

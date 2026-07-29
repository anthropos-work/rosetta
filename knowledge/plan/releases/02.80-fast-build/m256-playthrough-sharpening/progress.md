# M256 — progress

## Running ledger

- iter-01 (tok/bootstrap): TOK-01 authored — the headline parallelism lever proven **off the critical path**
  by D-v28-12's own re-cut (clause 1 is a per-test median, which worker count cannot move); replaced by the
  residual-`networkidle` per-test lever (12 login sites + 8 unfenced harness violations); cluster order
  inverted so **org-admin goes first** (it discharges clause 2's mutating floor *and* half of clause 3); a
  live local `demo-2` stood up as the measurement surface; iter-01's own D4 **refuted in-iter** by the
  Phase-0b audit (`actor.entitlement` is declared-only). — see `iter-01/progress.md`
- iter-02 (tik): **the baseline exists** — median per non-studio Playthrough **3.326 s** (n=3, cold run
  included, local `demo-2`, 18/18 green, 0 flake), so clause 1's target is **≤ 2.628 s**; suite wall-clock
  median **56.6 s** (reported, not gated). Two findings re-aimed the milestone: `pt-studio-advanced-generate`
  is a **FALSE GREEN** (it asserts the route's own header, and the LLM call was still in flight 19 s after
  the suite ended) so the gate's "irreducible LLM lane" has no referent; and the median's driver is the
  **per-test login handshake**, not the `networkidle` inheritance — iter-03 re-targets to seat-grouped
  `storageState` reuse. — see `iter-02/progress.md`
- iter-03 (tik): **clause 1's speed half MET** — median per non-studio Playthrough **3.326 s -> 2.014 s =
  0.6055x** (gate <= 0.79x), 0 flake over 3 runs, by banning `networkidle` at **20 harness sites** (12 login
  call sites inheriting the default + 2 per-surface `goto` overrides + 6 unbounded settles) and widening its
  fence from one route to the whole harness (`networkidle-fence.unit.spec.ts`, mutation-verified). Phase A's
  leg probe (2854 ms vs 423 ms) **falsified iter-02 D8 before any code was written**, rescuing the lever D8
  had written off and de-scoping `storageState` reuse (~200 ms, plus a false-green hazard). A latent hydration
  flake the removed settle had been masking was surfaced and fixed semantically. rext tag
  `fast-build-m256-networkidle-fence`, **on origin**. — see `iter-03/progress.md`
- iter-04 (tik, `closed-fixed-partial`): **the org-admin product exists** — the 9th product and one of the
  M201 corpus's un-homed clusters for five releases, from **0 of 4** covered to **4 declared / 2 landed
  GREEN** (`pt-orgadmin-tag-create`, `pt-orgadmin-setting-toggle` — both mutate **and** read back through a
  full reload). The milestone's open question is **answered: all four surfaces have real read-back surfaces**,
  none `unimplementable`. The other 2 are declared **`TODO`** with written diagnoses + named handlers and
  their specs parked in `e2e/drafts/` (**zero standing red**, D-v28-3). Clause 1 **re-verified on the grown
  denominator: 0.5434x**, with an original-16 honesty cross-check at 0.5347x. Clause 2's mutating count
  **1 -> 3** of 5. rext tag `fast-build-m256-orgadmin`, **on origin**. — see `iter-04/progress.md`
- iter-05 (tik, `closed-no-lift`): both parked journeys' hypotheses **tested and REFUTED**, with better
  evidence than a pass would have given. The create-role dialog has a **hidden mandatory step** — "Suggest
  skills" transforms it and the primary button becomes **`Generate`** (a live-LLM leg, so the journey belongs
  with the studio lane, not the median) — which confirms the enabled-`Save`-with-an-EMPTY-alert as a **real
  product defect**. The assign-tags modal is unreachable through **four** measured routes, including a
  refutation of **iter-04's own `Escape` fix** (it closes the MODAL, not the dropdown) and the finding that
  **`check({force:true})` flips the DOM but not antd's React state** — `isChecked()` true, submit disabled,
  nothing assigned. *`force: true` can manufacture a control the application does not know about.* Suite
  unchanged and green (136 passed, 20/22, 0 failing). rext tag `fast-build-m256-orgadmin-diagnoses`, **on
  origin**. — see `iter-05/progress.md`

- iter-06 (tik, `closed-fixed-partial`): **clause 2's mutating floor MET — 3 → 5, and machine-counted for the
  first time.** Both of the iter's own hypotheses were **refuted by live measurement** before a spec line was
  written: the skill-path CTA does *not* flip on `Start` (the row lands `progress=0, started_at=NULL`, so a
  started-but-unadvanced path reads "Start" **forever** — a real product wart, reported not worked around), and
  the sim launch writes **nothing** (0 `jobsimulation.sessions` rows), on a surface that could never have gone
  green anyway (`/sim/<slug>/session-list` gates on a Clerkenstein-absent `externalId` → a **false RED** trap).
  The two writes came instead from **writes the suite already made**: `pt-skillpath-legacy` extended by one click
  (closing a spec-vs-manifest fidelity gap — the manifest already promised "progress tracks") and the net-new
  self-cleaning `pt-skillpath-bookmark`. **H3 accepted:** a mutating Playthrough's **pre-state read IS its
  negative control** when the final assertion is a strict inequality/negation — and the retro-audit found **all
  three pre-existing mutating Playthroughs already had it, unnamed**. TOK-01 move 2's undelivered half landed as
  the **`@pt-mutation` registry + fence** (`MUTATES=5 READ-ONLY=14 UNKNOWN=2`, computed; mutation-verified). Side:
  a real **245 s suite flake** diagnosed to an **unreachable retry loop** (unbounded first `click` inherits the
  test budget) and fixed. Clause 1 re-verified on the grown denominator: **0.6245×**. `141 passed` ×3, 0 flake.
  rext tag `fast-build-m256-clause2-writes`, **on origin**. — see `iter-06/progress.md`

- iter-07 (tik, `closed-no-lift`): **the milestone's biggest RISK is retired, and the routed clause-2 mechanism
  is refuted.** (a) **Re-scope trigger NOT tripped:** the Phase-0b audit's F5 (*"no pre-onboarding state exists
  and none can be declared"*) **conflated org membership with onboarding completion** — onboarding lives in
  `public.user_params.onboarding` (`jsonb`, no onboarding table), is **NULL for all 191 seeded users** (the
  pre-onboarding state is the **DEFAULT**), and `/onboarding` **drives** for both seats. Onboarding is
  **UNBUILT, not impossible**; clause 3 keeps its full scope. (b) **GraphQL outcome-ablation REFUTED** — it
  yields a **dead page, not an empty surface** (`bodyLen 2147 → 24`, 0 nav, 0 buttons), so the control cannot
  discriminate; the replacement (**cross-vantage discrimination**) is identified with its O(tests) cost stated,
  not hand-waved. (c) **iter-02's studio false-green DIAGNOSIS overturned:** the blamed `Simulation Advanced
  Builder` string **never renders** (5 min poll); the matcher fires on **empty section scaffolding** at +2.1 s —
  so the previously-routed fix was a **no-op** that would have shipped as a fix. (d) This iter's own re-survey
  caught an **iter-06 defect**: a *comment* in the new fence file minted a **phantom Playthrough id**, invisible
  because `run-playthroughs.sh` reconciles with `ptreport` (which does not scan `@pt:` tags) — fixed + fenced
  across every file in `tests/`. rext tag `fast-build-m256-negctl-falsified`, **on origin**. — see
  `iter-07/progress.md`

- iter-08 (tik, `closed-fixed`): **the ONBOARDING product exists** — the last whole surface in the M201 curated
  corpus that no e2e suite had ever touched (*the first thing every real user does*), un-homed for five releases,
  and the one the milestone's own audit had priced as impossible. **1 live Playthrough**
  (`pt-onboarding-complete`, mutating **#6**) + **all 5 curated onboarding UCs declared with written verdicts** —
  every verdict **harness/seed work, 0 `unimplementable`**, which is the concrete reason the re-scope trigger
  never fired. The proof shape is the cheapest in the suite: **`/onboarding` is its own read-back** (it SERVES
  the flow while incomplete and REDIRECTS to `/home` once complete), so one route gives both the pre-state
  absence (the negative control) and the persisted post-state. Seat choice was load-bearing — `pt-free` is
  driven by **0** other use cases, so an irreversible write cannot perturb another Playthrough. Clause 1
  improved to **0.5950×**, and the original-16 cross-check (1.772 s = 0.5328×) **retires iter-06's suspected
  regression as run-to-run variance**. `145 passed` ×3, 0 flake; `ptreport` 22/29 passing, 0 failing.
  `@pt-mutation` registry: **MUTATES=6 READ-ONLY=14 UNKNOWN=2**. rext tag `fast-build-m256-onboarding`, **on
  origin**. — see `iter-08/progress.md`

- iter-09 (tik, `closed-fixed`): **clause 3's VERDICT half is COMPLETE** — [`coverage-verdicts.md`](coverage-verdicts.md)
  gives all **28** M201 curated use cases a written verdict, and **0 are `unimplementable-without-platform-edit`**:
  every uncovered one is blocked on harness work, a fixture, a credential, or a deliberate tier reservation —
  never on the platform refusing to be driven. **The re-scope trigger cannot fire on current evidence.** Pricing
  them properly paid off immediately: `workforce-intelligence.organization-feedback` had been un-homed for five
  releases, and the one question a verdict forces (*is its data already seeded?*) collapsed it — `feedback
  rows=70` on every reset — so it cost **a page object and a spec** and was **landed in the same iter**
  (`pt-workforce-org-feedback`, asserting **both** sentiment polarities because a one-sided recap is the empty
  state the final rules out). Two structural findings: the **résumé fixture is a SHARED blocker** for
  `profile-skills.import` + onboarding's import UC (pay it once — it would be the suite's first file-upload
  Playthrough), and **`profile-skills.self-evaluation`'s M206 reservation is WEAK** (a persist-then-observe write
  needing no mirror engine → recommended for re-homing to M257). Clause 1 **0.5652×**; `146 passed` ×3, 0 flake;
  `ptreport` 23/30 passing, 0 failing. rext tag `fast-build-m256-verdicts`, **on origin**. — see
  `iter-09/progress.md`

- iter-10 (tik, `closed-no-lift`): **D-v28-5 could not be measured, and finding out why exposed a worse,
  previously unrecorded defect.** Driving the presenter's real clicks on the cockpit never switched hero at all —
  because **`cockpit-manifest.json` and `fake-fapi-roster.json` have drifted completely apart**:
  `run-playthroughs.sh --reset` re-exports the **roster** (30 `pt-*` keys — which is why 23 Playthroughs are
  green) but **never the cockpit manifest** (still `maya-thriving`, `tom-struggling`, … baked at bring-up). So the
  cockpit renders **35 `[Log in as]` buttons naming heroes that no longer exist**, and the handshake's
  *deliberate* best-effort tolerance for an unknown `__clerk_identity` (`server.go:347-349` — right for a
  malformed deep-link) turns that into a **successful-looking WRONG login**: the presenter gets whoever was last
  active, with no error in the UI and none in the log. That is very likely the substrate of D-v28-5's
  "two-or-more clicks" symptom — **recorded as a hypothesis, not a finding**, because it cannot be isolated until
  the cockpit can select a hero that exists. **No code shipped, deliberately** (iter-07 D31's lesson: an
  unverified lifecycle fix that closes a handler and changes nothing is the worst outcome available). Suite
  untouched and green. — see `iter-10/progress.md`

- iter-11 (tik, `closed-fixed`): **clause 2's `blocked` sub-clause is DISCHARGED — 0 → 1 — and the Playthrough
  that discharged it paid for itself on its first run.** `blocked` had been 0 for 23 Playthroughs not because
  nobody wrote the test but because **the seeded world contained no refusal**: the g3 `FEATURE_JOB_SIMULATIONS`
  grant was written for *every* membership (20/20 · 40/40 · 40/40 · 40/40, measured). So the deliverable is a
  **seed** change — `StoryOrg.sim_feature_disabled` → the single recognition point `SimFeatureEnabled()` → the
  `UsersSeeder` guard, an **opt-OUT** because the granted default is load-bearing for every existing world — plus
  `pt-aisim-org-feature-blocked`, which asserts the refusal from **four** directions (the deny dialog PRESENT, the
  **org NAMED** — *"contact your administrator at Halcyon Retail"*, so it proves WHICH tenant was refused — the
  launch confirmation ABSENT, and the URL still on the detail route; *a dead page satisfies exactly one*). Phase A
  probed first and was right to: two in-repo comments described the deny surface incompatibly ("deny modal" vs
  "empty `<main>`"), and the "empty main" reading would have shipped a spec that passes for any broken page.
  **`NEGCTL-M256-cross-vantage` is now LIVE on its first pair** — the same locator asserted ABSENT by
  `pt-aisim-chat-launch` (Org A, granted) and PRESENT here (Org B, withheld), so each is the other's control:
  negative controls **6 → 8 of 24**, and the count is now **computed by the fence** (uncovered ids named, floor
  pinned) instead of narrated. **The RED was worth more than the green:** the new Playthrough failed its first
  live run because `stackseed --reset` deleted only `g2` grants — **`g3` had accumulated for four releases (731
  rows for 140 memberships, 540 ORPHANED)** and, ids being deterministic, a stale row **silently re-granted** the
  feature, so the withheld org came up 20/20 and the world was never in its declared state. Fixed as a class
  (`resetCasbinPTypes = {g2,g3}`, exact-set pinned, never a TRUNCATE). *An additive leftover in a reset path is
  invisible for exactly as long as every test wants the thing PRESENT.* Also caught in-iter: the seed's new field
  made the stack's **pinned** `stackseed` hard-fail the reset *after* truncating — so the tag was pushed and
  `stack-demo/rosetta-extensions` re-pinned + rebuilt, and run 4 is that binary. **148 passed ×4** consecutive
  reset-to-seed runs, 0 flake; `ptreport` **24/31 passing, 0 failing, 0 unimplementable**. Clause 1 re-verified
  **0.6863×** — the drift from iter-09's 0.5652× is the MACHINE, proven by the untouched-original-16 cross-check
  moving with it (0.5284× → 0.6055×). rext tag `fast-build-m256-blocked-outcome`, **on origin**. — see
  `iter-11/progress.md`

- iter-12 (tik, `closed-fixed`): **negative controls 8 → 13 of 24 — and the mechanism's LIMIT was measured,
  which matters more than the six it covered.** A cross-vantage control now lives in its own file class
  (`negative-controls.spec.ts`, **no `@pt:` id** — not a Playthrough, not reconciled, **never in the gated
  median**, batched by vantage so N absences cost ONE login), asserts **liveness before absence** (polled — a
  bare `.count()` after a `domcontentloaded` nav reads the pre-hydration DOM and calls a working app dead),
  and its coverage is a **machine-checked fail-closed LINK** (a control naming an id no Playthrough declares,
  or a token that is not id-shaped, fails the fence). **The limit:** a contrast vantage discriminates only an
  **org- or hero-specific** outcome; a **structural** final (a stat label, a chart, a table's first row)
  renders for any populated org — measured, Org A's manager reads `verifiedSkillsStat` 1 / `skillCharts` 10 /
  `workSection` 1 — so **9 of the remaining 11 cannot have one** and must instead have their finals sharpened
  to name real seeded data. Writing them anyway would have re-introduced **the exact vacuity iter-07
  refuted**, through the mechanism adopted to replace it. Two vantages rejected on measurement: the hiring one
  **ejects the browser to PRODUCTION** (`app.anthropos.work/login`, bodyLen 162), and a Playthrough on a known
  false green must not be given a control (it would certify it). **A control earned its keep on its first
  run:** `pt-aireadiness-manager-howwemeasure`'s step-name asserts **match on a non-readiness org**, because
  `/ai-readiness` without the feature renders a live *upsell* panel naming the steps. Two self-inflicted
  defects caught while building the guards, both repeats: the fence **harvested its own prose** (iter-07's
  phantom-comment class, one grammar later) and its detector reused a `/g` regex whose `lastIndex` would have
  halved the control set. Side: `assignments-page.ts`'s last **unbounded `waitFor`** bounded — it turned a
  legible failure into an opaque 240 s hang, twice.
  **AND THE HEADLINE, which is an escalation: clause 1 is NOT DECIDABLE at n=3 on this host.** Six full-suite
  runs, same box, **the original 16 specs unchanged since iter-03**: the control subset spans **0.5281× →
  1.0762× (2.04×)** with **no trend** (newest 0.529×, oldest 0.528×, the extreme in between). The gated figure
  at **n=6 is 0.8129× — OUTSIDE the `≤ 0.79×` gate**; the flattering denominator (original-16, 0.7063×) is
  inside it, and picking that one would be the dishonesty this milestone has refused for eleven iters. So the
  MET readings from iter-03 onward were **favourable samples, not verdicts** — *a relative gate needs its
  noise floor published next to it or it is not falsifiable.* Every remedy (raise n + publish the spread ·
  pair the baseline · normalise within-run · a stable host) changes **D-v28-12**, so it is escalated, not
  actioned. `150 passed`, `ptreport` **24/31 passing, 0 failing, 0 unimplementable**; 1 flake in the batch
  (the now-bounded 240 s hang). rext tag `fast-build-m256-negctl-crossvantage`, **on origin**. — see
  `iter-12/progress.md`

- iter-13 (tik, `closed-fixed`): **negative controls 13 → 16 of 24 — by refuting the premise that the nine
  structural finals could not have one.** iter-12's measurement was right (a stat LABEL / a chart count / a
  "Work" section renders for every member, so no vantage can falsify it, and no suppression switch can exist);
  the *reading* of it was wrong. Those finals were structural **because they were written structurally**. Phase
  A measured three independent hero-specific facts on the same two surfaces, each reading **0** for the same
  contrast seat that could not falsify the structural version: the rendered `Verified Skills` stat equals the
  seeded `skills.verified` **exactly** (8 / 3 / 2 across the three seats) and `All Skills` equals seeded
  `verified + mapped` **exactly** (8+12=20, 2+8=10 — two heroes, independently); the `"<role> at <org>"`
  context line; the `My Closest Roles` recommendation, which renders only once the matcher has enough VERIFIED
  evidence (present for the 8-verified hero, absent for the 3- and 2-verified seats). So the three finals were
  re-aimed — the old structural ones **retained as intermediates** — and the numbers are **machine-linked to
  the seed** by a net-new `lib/seed-facts.ts` + a **fail-closed** `seed-facts-fence.unit.spec.ts` (its FIRST
  test asserts the parse is not vacuous, because a reconciliation over an empty regex parse passes every
  comparison silently — the milestone's signature defect). **10 mutants RED**, including each Playthrough
  driven on the contrast seat and each control assertion individually. **The iter shipped a false RED and
  caught it by watching:** `\b` in a `hasText` regex is unreliable because `textContent` concatenates sibling
  nodes with no separator (`…Meridian LabsFeb 2024 - Present…`), so both new conjunction locators read 0 on a
  page that plainly rendered the thing — the same constant is safe under `getByText` and broken under
  `hasText`, now documented at both constants and pinned by a test that holds the bug.
  **AND the `pt-assignment-assign` flake is ROOT-CAUSED AND FIXED, so clause 1's flake half is met by a fix
  rather than a favourable batch.** The pre-fix batch was `165 passed` ×2 then **1 failure on run 3**. Three
  hypotheses were refuted by measurement — iter-11's bloated policy (`g3 = 171` for 191 memberships, **0
  orphans**, exactly as designed), the antd `maskClosable` re-click (it *throws* on the mask; the modal
  survives), and `press('Enter')` with the dropdown closed (`aria-expanded` stays true) — plus the modal
  **surviving 151 s unattended**. The trace named the real one: the modal is **ROW-SCOPED** (*"Assign Skill
  Path to `<member>`"*, rendered by the row's action cell), it opened 2.2 s after the first row painted while
  the table was still settling, the settling re-render detached the Select's input and **took the modal with
  it**, and the remaining time decomposes exactly as the ladder's own bounds — 3 × 15 s + 20 s + 15 s = **84 s,
  the reported duration**. *Every bound in that ladder was correct; bounding makes a stuck attempt yield, it
  does not make a **dead subject** detectable.* Fixed in three parts (a `dialogIsOpen()` re-open guard at the
  top of every attempt; `waitForMembersTableSettled()` so it does not race; and `openBuilderAndPickSkillPath()`,
  because once a modal can be re-opened the target must be read from the builder that **accepted** the pick or
  the read-back can assert the wrong row). Recovery proven **deterministically** — the exact failing state
  reached with a real user action (the modal's own Cancel; `Escape` is disabled, measured), then the ladder
  re-opened the builder and the pick took. **Post-fix gate: `166 passed` ×3 consecutive cold reset-to-seed
  runs, rc 0, `ptreport` 24 passing / 0 failing / 7 TODO / 0 unimplementable, 0 flake**; the assign Playthrough
  6.9–11.5 s where it had spiked to 84 s. — see `iter-13/progress.md`

- iter-14 (tik, `closed-fixed`): **negative controls 16 → 20 of 24 — the same move as iter-13, one level up,
  and the level is what mattered.** iter-13's contrast for a *person* final is another person; reused here it
  would have been wrong, because the four Workforce-Intelligence finals are read BY `pt-manager` and are about
  **her org's aggregates** — she cannot falsify her own dashboard. **The contrast vantage follows the SUBJECT
  of the final** (D68), so an org-aggregate final needs a manager of a **second seeded TENANT**
  (`pt-ai-manager`, Org C), and Phase A qualified her on measurement rather than convenience: her org renders
  all four surfaces *completely* (40 members, 109 mapped, 88 verified, 182 sims, 20 roster rows, a full
  succession projection, 20 feedback rows), so the ONLY thing absent is Org A's data. The four finals —
  a Role column · two stat LABELS + an `<svg>` · `/ready/i` + `/at.?risk/i` · a recap label + both polarity
  words, each with `rows > 0` — were **measurably true of that other tenant too**, i.e. all four would have
  passed if the manager had been served a different customer's workforce. They now name the seed's own org
  facts: the **org email domain** (`@<org.slug>.com`, the **pagination-proof** anchor — a roster shows 20 of
  40, so any one row is a bet on sort order, while the domain is on every org-member row; bounded honestly at
  15 of 20, the rest being `Candidate`-role members on external addresses), `Overall Members` == the seeded
  `org.size`, and the seeded hero with her seeded role. **11 mutants RED**, in two deliberate groups (D73):
  each control's absence re-aimed at Org C's own data (the control can fire) **and each sharpened Playthrough
  driven on Org C** (the final discriminates) — the second group being the one that matters, since the first
  alone can certify healthy controls over still-vacuous Playthroughs. The seed link was mutated three ways
  (size, slug, and a parse-emptying rename → the fail-closed non-vacuous test fires). **Two self-inflicted
  errors caught by measurement, both instructive:** `readOrgStat` was written as a copy of
  `ProfilePage.readSkillStat` and returned `null` against a dashboard that plainly renders `40` — the two
  cards have **opposite shapes** (`"Verified Skills\n8"` vs a bare label `<span>` whose parent reads
  `"40Overall Members40 active"`, value first and no whitespace, because `textContent` concatenates) — and
  **the Phase A probe's own settle predicate certified a hydrating grid as populated** (`rows > 5` is
  satisfied by 20 content-free skeleton rows), which nearly banked a permanent absence that does not exist;
  *a settle predicate the empty state satisfies is not a settle predicate* (D71), the same could-not-fail
  defect committed in the instrument. What survived that retraction is real and routed:
  `pt-activity-drilldown`'s `rows > 0` **is** satisfied by the skeleton, and its control is available from the
  drill-down, which names members with their roles. `170 passed` ×3 consecutive cold reset-to-seed runs, rc 0,
  **0 flake**, `ptreport` 24 passing / 0 failing / 7 TODO / 0 unimplementable; the four sharpened Playthroughs
  run **1.4–1.9 s**, all below the suite median, so a tenancy proof is no more expensive than a structural
  one. — see `iter-14/progress.md`

## Baseline — MEASURED (iter-02, 2026-07-28)

| Figure | Value |
|---|---:|
| **Median per-Playthrough, 16 non-studio — the GATED metric** | **3.326 s** |
| **Clause 1 target (0.79x)** | **<= 2.628 s** |
| Median per-Playthrough, all 18 (cross-check) | 3.067 s |
| Suite wall-clock, 132 tests (REPORTED, not gated) | median **56.6 s** (85.4 cold / 56.6 / 54.4) |
| Studio lane, excluded (and NOT LLM-bound — iter-02 D6) | 1.26 s / 1.84 s |

**Pinned statistic (D7)** — recompute identically or the ratio is meaningless: the median across the 16
non-studio Playthroughs of each Playthrough's median across **3 consecutive `--reset` runs**, run 1 being
the first (cold) run after bring-up and **included**.

**Post-iter-14 — THE CURRENT FIGURE.** n=3, same host, cold reset-to-seed each:

| statistic | run 1 | run 2 | run 3 | median (n=3) | range |
|---|---:|---:|---:|---:|---:|
| all 22 non-studio (**REPORTED, per D-v28-13**) | 0.6013× | 0.6164× | 0.6464× | **0.6314×** | 1.075× |
| ORIGINAL 16 only (**the control subset**) | — | — | — | 0.5562× | — |

`170 passed` ×3, **0 flake**, rc 0 ×3. Suite wall-clock 1.2 / 1.2 / 1.3 m. The four sharpened Playthroughs
read **1.4–1.9 s** each — all *below* the suite median, so sharpening a READ costs nothing measurable (the
opposite of iter-06's finding that proving a WRITE costs more than proving a render).

**Still more evidence for the D-v28-13 recut.** The ORIGINAL-16 control subset — code no iter has touched
since iter-03 — has now been observed at **0.5281× · 1.0762× · 0.7517× · 0.9321× · 0.5562×** across five
batches on one host. A gate at 0.79× sits inside its own noise floor. **iter-14 landed no speed mechanism, so
clause 1's leg half has nothing new to measure**; its flake half is **MET** (0 flake ×3).

`ptreport`: **24/31 passing, 7 `[TODO]`, 0 failing, 0 unimplementable.** `@pt-mutation` registry, computed:
**MUTATES=6 READ-ONLY=16 UNKNOWN=2**. `@pt-negative-control` registry, computed: **20 of 24** (8
self-declared + 12 via the control spec); named uncovered: `pt-activity-drilldown`,
`pt-hiring-recruiter-compare`, `pt-studio-advanced-generate`, `pt-studio-guided-generate`. **`blocked`
outcomes: 1.**

**Post-iter-13 (superseded as the current figure).** Post-fix batch, n=3, same host, cold reset-to-seed each:

| statistic | run 4 | run 5 | run 6 | median (n=3) | range |
|---|---:|---:|---:|---:|---:|
| all 22 non-studio (**REPORTED, per D-v28-13**) | 1.2327× | 1.1124× | 0.8419× | **1.0523×** | 1.46× |
| ORIGINAL 16 only (**the control subset**) | — | — | — | 0.9321× | — |

`166 passed` ×3, **0 flake**, rc 0 ×3. Suite wall-clock 2.1 / 1.9 / 1.5 m. The pre-fix batch (runs 1–3, which
carried the assign flake) read median **0.9170×**, range 0.8118×–1.1425×, original-16 **0.7517×**.

**This is more evidence for the D-v28-13 recut, not against it.** The ORIGINAL-16 control subset — code no
iter has touched since iter-03 — has now been observed at **0.5281× · 1.0762× · 0.7517× · 0.9321×** across
batches on one host. A gate at 0.79× sits inside its own noise floor, exactly as D-v28-13 recorded. **iter-13
landed no speed mechanism, so clause 1's leg half has nothing new to measure**; its flake half is **MET** —
and met because the flake was diagnosed and fixed (iter-13 D65), not re-rolled.

`ptreport`: **24/31 passing, 7 `[TODO]`, 0 failing, 0 unimplementable.** `@pt-mutation` registry, computed:
**MUTATES=6 READ-ONLY=16 UNKNOWN=2**. `@pt-negative-control` registry, computed: **16 of 24** (8 self-declared
+ 8 via the control spec). **`blocked` outcomes: 1.**

**Post-iter-12 (superseded as the current figure; the escalation it raised still stands) — a NON-VERDICT (see iter-12 D60).** Six full-suite runs across
iters 11–12 on the same host, with the ORIGINAL 16 specs unchanged since iter-03:

| statistic | min | max | spread | median (n=6) | gate |
|---|---:|---:|---:|---:|---:|
| all 22 non-studio (**the GATED figure**) | 0.5701× | 1.1121× | 1.95× | **0.8129×** | ≤ 0.79× → **OUTSIDE** |
| ORIGINAL 16 only (**the control subset**) | 0.5281× | 1.0762× | **2.04×** | 0.7063× | — |

**Clause 1 cannot be declared MET.** The control subset is code no iter touched and it varies by 2.04× with
no trend, so the pinned statistic (median of 3 consecutive runs, against a baseline measured in a *different*
batch) does not absorb this host's variance — a batch of 3 lands anywhere from ~0.53× to ~1.08×. The readings
below are retained as the per-batch record, but each must now be read as **one sample of a distribution**, not
as a verdict. **Escalated as `MEASURE-M256-clause1-sampling`** (a D-v28-12 decision). Nothing here retracts
iter-03's speed work, which was measured **directly at the leg** (2854 ms → 423 ms for the same navigation),
not as a suite ratio.
`ptreport`: **24/31 passing, 7 `[TODO]`, 0 failing, 0 unimplementable.** `@pt-mutation` registry, computed:
**MUTATES=6 READ-ONLY=16 UNKNOWN=2**. `@pt-negative-control` registry, computed: **13 of 24** (8 self-declared
+ 5 via the control spec). **`blocked` outcomes: 1.**

**Post-iter-11 (one batch of 3; 22 non-studio):** median per
non-studio Playthrough **2.282 s = 0.6863×** of baseline — gate `<= 0.79×` **MET**; **honesty cross-check** over
the ORIGINAL 16 only: **2.014 s = 0.6055×**. **0 flake** over **4** consecutive reset-to-seed runs (`148 passed`
×4 — the 4th on the stack's own re-pinned `stackseed`). Suite wall-clock 56.1 / 71.9 / 74.2 s (reported, not
gated). `ptreport`: **24/31 passing, 7 `[TODO]`, 0 failing, 0 unimplementable.** `@pt-mutation` registry,
computed: **MUTATES=6 READ-ONLY=16 UNKNOWN=2**. `@pt-negative-control` registry, computed (net-new): **8 of 24**.
**`blocked` outcomes: 1** (was 0). *On the drift from iter-09's 0.5652×:* the cross-check over the 16 specs this
iter did not touch moved by the same factor (0.5284× → 0.6055×), and a subset with zero code change cannot
regress from code — so it is the environment (four back-to-back suite runs on a 9.70 GiB Docker VM against the
documented 12 GB floor), reported as variance with the control shown, exactly as iter-08 did in the opposite
direction. Clause 1 holds with margin either way.

**Post-iter-09 (superseded; 21 non-studio):** median per
non-studio Playthrough **1.880 s = 0.5652×** of baseline — gate `<= 0.79×` **MET**; **honesty cross-check** over
the ORIGINAL 16 only: **1.757 s = 0.5284×**. **0 flake** over 3 consecutive cold reset-to-seed runs (`146
passed` ×3). `ptreport`: **23/30 passing, 7 `[TODO]`, 0 failing, 0 unimplementable.** `@pt-mutation` registry,
computed: **MUTATES=6 READ-ONLY=15 UNKNOWN=2**.

**Post-iter-08 (superseded; 20 non-studio):** median per
non-studio Playthrough **1.979 s = 0.5950×** of baseline — gate `<= 0.79×` **MET**; **honesty cross-check** over
the ORIGINAL 16 only: **1.772 s = 0.5328×**. Suite wall-clock median **55.3 s** (reported, not gated). **0 flake**
over 3 consecutive cold reset-to-seed runs (`145 passed` ×3). `ptreport`: **22/29 passing, 7 `[TODO]`, 0
failing.** `@pt-mutation` registry, computed: **MUTATES=6 READ-ONLY=14 UNKNOWN=2**. *This retires iter-06's
suspected regression:* its original-16 cross-check read 2.006 s and the drift from iter-04's 0.5347× was
attributed to laptop variance rather than asserted as a regression — the same cross-check now reads 0.5328×,
within noise of iter-04, so variance it was.

**Post-iter-06 (superseded; on the 19-non-studio denominator):** median per
non-studio Playthrough **2.077 s = 0.6245×** of baseline — gate `<= 0.79×` **MET**; **honesty cross-check** over
the ORIGINAL 16 only: **2.006 s = 0.6033×**. Suite wall-clock median **53.5 s** (reported, not gated). **0 flake**
over 3 consecutive cold reset-to-seed runs (`141 passed` ×3). `ptreport`: **21/23 passing, 2 `[TODO]`, 0 failing.**
*Slower than iter-04's 0.5434×, expectedly and in two parts:* by design, `pt-skillpath-legacy` grew 3.16 → 4.14 s
and the new `pt-skillpath-bookmark` sits above the median — **proving a write costs more than proving a render**,
which is the trade clause 2 asks for; and the original-16 cross-check drifted 1.778 → 2.006 s with no code change
to those specs (run-to-run variance on a 9.70 GiB Docker VM against the documented 12 GB floor).

**Post-iter-04 (re-measured on the GROWN denominator, 18 non-studio):** median per non-studio Playthrough
**1.808 s = 0.5434x** of baseline — gate `<= 0.79x` **MET**; **honesty cross-check** over the ORIGINAL 16
only: **1.778 s = 0.5347x**, so the gain is in the existing tests, not an artifact of adding fast ones (the
two new Playthroughs are *slower* than the median). Suite wall-clock median **44.5 s** (reported, not gated).
**0 flake** over 3 consecutive runs (`136 passed` x3). `ptreport`: **20/22 passing, 2 `[TODO]`, 0 failing.**
(Post-iter-03 was 2.014 s = 0.6055x on the 16-test denominator.)

**Environment:** `Kirality-Mac-Pro-6.local`, darwin 25.1.0, Docker VM **9.70 GiB** (vs the 12 GB floor);
`demo-2` offset 20000, **localhost/http**, `--no-public-host`. **Per D-v28-12 no number here may be quoted
as comparable to billion's 228 s** — the absolute billion re-measure is routed to M258.

## Next-iter routing

Fate-3 items land here.

| Handler | What | Target |
|---|---|---|
| `FIX-M256-autoverify-fapi-libressl` | `autoverify.sh` check (d) probes the fake-FAPI with LibreSSL `curl`, which cannot handshake the mkcert leaf on macOS → warns *"NOBODY CAN LOG IN"* on a working stack (iter-01 D5). Give it a probe independent of the host TLS stack. | a later tik of M256 |
| `DOC-M256-ptworld-reset-comment` | `playthroughs/seed/pt-world.seed.yaml`'s header claims the showcase world is "not touched by pt-world's reset". `doReset` takes **no org filter** — it is (audit F6). | a later tik of M256 |
| `PERF-M256-parallel-lane` | The cookie/`__client`-scoped Clerkenstein registry **or** one fake-FAPI per worker. Both priced in iter-01 D1. A **wall-clock** lever, not a median one — no M256 gate clause needs it. | a future release milestone |
| `PT-M256-orgadmin-role-create` | **SHARPENED (iter-05 D18).** "Suggest skills" transforms the dialog and the primary button becomes **`Generate`** — "Core skills" is a MANDATORY step, so drive that path and assert at its **completion boundary**, budgeting the journey with the **studio lane, not the median** (it is live-LLM). Separately: **report the product defect** — `Save` is enabled with the form incomplete and fails silently with an EMPTY `alert` region. | next iter |
| `PT-M256-orgadmin-member-tag` | **SHARPENED (iter-05 D19).** Unreachable through 4 measured routes: dropdown intercepts pointer events over its own modal (invisible to actionability checks); `Escape` closes the MODAL not the dropdown; an in-modal outside-click leaves 9 menuitems open; `check({force:true})` flips the DOM but not antd's React state (submit stays disabled, tally stays 0). Untried: the `<label for=…>` element; a route that never opens a dropdown. If both fail, declare **`unimplementable-without-platform-edit`** with the four-attempt evidence. | next iter |
| ~~`PT-M256-clause2-fifth-write`~~ | **DONE (iter-06 D24).** Answered differently than D20 framed it — not from a new surface but from **writes the suite already made**: `pt-skillpath-legacy` extended by one click + the net-new `pt-skillpath-bookmark`. Mutating **5**, machine-counted by the `@pt-mutation` fence. | closed iter-06 |
| ~~`NEGCTL-M256-ablation-harness`~~ | **REFUTED (iter-07 D29).** GraphQL outcome-ablation yields a **dead page, not an empty surface** (`bodyLen 2147 → 24`, 0 nav, 0 buttons) — the control would pass for any Playthrough including one asserting pure chrome, so it cannot discriminate. A gentler ablation needs per-operation shapes: **O(queries), not O(surfaces)**. Do not re-try verbatim. | closed iter-07 |
| `NEGCTL-M256-cross-vantage` | **20 of 24 (iter-14) — 4 remain, and BOTH of the reachable two are now priced.** iter-14 took the four workforce finals by the same move as iter-13's profile three, with one generalisation that is the reusable part: **the contrast vantage follows the SUBJECT of the final** (D68) — a person-final needs another person, an ORG-aggregate final needs a second seeded **TENANT** (`pt-ai-manager`, Org C, measured to render all four surfaces completely). **`pt-activity-drilldown`** is next and both halves are measured: (a) its `contentRows().first()` visible + `count() > 0` pair **is satisfied by 20 content-free skeleton rows** (measured live; not a false green, because the drill step needs a row `<a>`, but a weak assertion) → assert rows that carry TEXT; (b) its control comes from the **drill-down**, which names members WITH their roles — Org A's first content drills to `Pat Ellis / DevOps Engineer`, the SAME content id on Org C drills to `Mei Costa` / `Theo Lindqvist`; the first row is deterministically a hero-session one (the grid sorts by most-recent and hero sessions are seeded at reset). **`pt-hiring-recruiter-compare`** still needs a **same-vantage** control (its contrast vantage ejects the browser to production, iter-12); Org D's authored facts give a seed-pinned cardinality final (`size: 40`, `role_mix 0.1/0.9` → 5 shared positions / ~36 candidates), but where the ABSENCE comes from is the open question — and a written verdict is the honest fallback, never a manufactured control. The 2 STUDIO remain blocked behind `FIX-M256-studio-false-green`. | **iter-15+** |
| ~~`NEGCTL-M256-cross-vantage` (iter-13 framing)~~ | **(Superseded by the row above.)** **16 of 24 (iter-13) — and the "9 STRUCTURAL" class below was NOT a wall.** iter-12's measurement was right and its conclusion was too strong: those finals were structural *because they were written structurally*, and the mechanism's limit was the ASSERTION, not the vantage. Three are now covered (`pt-profile-{verified,growth,timeline}`) by re-aiming each final at the hero's own seeded data — magnitudes machine-linked to the seed by `lib/seed-facts.ts` + a fail-closed fence — after which the SAME contrast seat falsifies all three. **The remaining 6 of that class are reachable by the same move:** `pt-workforce-{roster,funnel,succession,org-feedback}`, `pt-activity-drilldown`, `pt-hiring-recruiter-compare`. ⚠ The hiring one needs care — iter-12 measured that its contrast vantage **ejects the browser to PRODUCTION**, so its control must come from a sharpened final on the *same* vantage, never a contrast org. The 2 STUDIO remain blocked behind `FIX-M256-studio-false-green`. | **iter-14+** |
| ~~`NEGCTL-M256-cross-vantage` (iter-12 framing)~~ | **(Superseded by the row above.)** (a) **9 STRUCTURAL** finals — `pt-workforce-{roster,funnel,succession,org-feedback}`, `pt-activity-drilldown`, `pt-profile-{verified,growth,timeline}`, `pt-hiring-recruiter-compare` — have **NO contrast vantage** (a stat label / chart / first table row renders for any populated org; measured: Org A's manager reads `verifiedSkillsStat` 1, `skillCharts` 10, `workSection` 1). Their route is to **sharpen the final to name real seeded data** — O(tests), and it strengthens the Playthrough regardless. Writing contrast controls for them would re-introduce iter-07's refuted vacuity. (b) **2 STUDIO** are blocked behind `FIX-M256-studio-false-green` — a control on a known false green would certify it. | **iter-13+** |
| `PT-M256-readiness-step-asserts` | **STILL OPEN after iter-14 — unstarted, and it should ride with the iter-15 sharpening batch** (same defect shape: an assertion satisfied by the wrong state). **NEW (iter-12), found by a control on its first run.** `pt-aireadiness-manager-howwemeasure`'s `MANAGER_STEP_NAMES` assertions match **page-wide** and are satisfied by the **not-enabled upsell panel** on `/ai-readiness` (which names the very steps). Re-scope them inside the method panel. The Playthrough is still covered via `methodHeading()`, which does discriminate. | a later tik of M256 |
| `MEASURE-M256-clause1-sampling` | **NEW (iter-12 D60) — ESCALATED, and it gates the milestone's own verdict.** Clause 1's pinned statistic is not robust to this host's variance: over 6 runs the CONTROL subset (unchanged since iter-03) spans **0.5281×–1.0762× (2.04×)** with no trend, and the gated figure at n=6 is **0.8129×, outside the `≤ 0.79×` gate**. The MET readings from iter-03 on were favourable samples. Remedies, all **D-v28-12** decisions: raise n + publish the spread · make the measurement **paired** (baseline in the same batch) · normalise within-run against an invariant leg · move it to a stable host. `pt-assignment-assign` is the largest single contributor to both the median and its variance — a lever aimed at it would cut both. | **user / roadmap call** |
| ~~`NEGCTL-M256-cross-vantage` (iter-11 framing)~~ | **MECHANISM PROVEN LIVE (iter-11).** Negative controls **8 of 24**, and the count is now **computed by the mutation fence**, which names the 16 uncovered ids on every run. The reference implementation is the `pt-aisim-chat-launch` ∥ `pt-aisim-org-feature-blocked` pair: one locator, two orgs, opposite verdicts, both live. **What it teaches about cost:** that pair was cheap because the two vantages differ by **seeded state** (a withheld g3 grant), and a symmetric pair contributes **two** controls. A pair that differs only by test code is the O(tests) case below. So triage the remaining 16 by *"is there a hero/org for whom this outcome legitimately does not exist in the seed?"* first — the seed-state cases are the cheap tail. | **iter-12+** |
| ~~`NEGCTL-M256-cross-vantage` (original framing)~~ | **(iter-07 D30, superseded by the row above.)** Clause 2's negative controls stood at **5 of 21**; the 5 mutating ones get it free from their pre-state read (D22), and the **16 non-writing** ones need a different mechanism. Run each Playthrough's **own final locator against a contrast vantage** — a hero/org for whom the outcome legitimately does not exist. Real absence, app stays alive, and it proves **WHICH** data not merely **THAT** data (the M219 lesson, per-Playthrough). Cost stated honestly: **O(tests), not O(surfaces)** — budget it across more than one tik. | **iter-08+** |
| `FIX-M256-studio-false-green` | **RE-AIMED (iter-07 D31) — the old diagnosis was FALSE.** iter-02 blamed the route's `Simulation Advanced Builder` header; a 5-minute poll of the real journey shows that string **NEVER renders**. The matcher actually fires on the designer's **empty section scaffolding** ("Scenario Characters" / "Mission Tasks" headings) at **+2.1 s**, before the LLM draft populates it. **Deleting the header alternative is a NO-OP — do not ship it as a fix.** Assert a **POPULATED** section instead (a character card / a non-zero `designer.actors.counter.label` count). Evidence is attached to the locator in `studio-builder-page.ts`. Until it lands, both studio Playthroughs stay `@pt-mutation: UNKNOWN`, never `MUTATES`. | a later tik of M256 |
| `DOC-M256-llm-lane-premise` | `playthroughs.md` § the `studio` product + the M256 overview + D-v28-9 all describe the advanced builder as reaching a generation completion boundary. Correct **once**, against the fixed behaviour. **Still not dischargeable (iter-07):** a section *heading*'s presence does not answer *"did the generation complete on this host?"*. Measuring section **CONTENT** answers the fix and the doc premise together — keep them one piece of work. | the same tik as the fix |
| ~~`ONBOARD-M256-build`~~ | **DONE (iter-08).** The onboarding product exists: `pt-onboarding-complete` live (mutating #6) + **all 5 curated onboarding UCs declared with written verdicts**, every verdict harness/seed work and **0 `unimplementable`**. | closed iter-08 |
| `ONBOARD-M256-import-path` | The **4 remaining** onboarding UCs, each with its specific missing piece already written into `manifest/onboarding.yaml`: a **résumé fixture** (spec §5.4's `fixtures/` dir is still EMPTY — this would be the suite's FIRST file-upload Playthrough) + an async LLM import; the **org-prepared trigger condition**, not yet identified (iter-08 measured the *import form* for a hero WITH a populated profile); an **org-less actor** (F5's one kernel of truth — needs a member-less user + a roster seat); an **Org C stage-0** seat; a **day-0 hiring-org** seat (the only onboarding UC whose final spans two apps). | a later tik of M256 |
| ~~`VERDICT-M256-remaining-uncovered`~~ | **DONE (iter-09).** [`coverage-verdicts.md`](coverage-verdicts.md): all **28** curated UCs verdicted, **0 `unimplementable`**. Clause 3's verdict half is COMPLETE. | closed iter-09 |
| `PT-M256-resume-fixture-pair` | `profile-skills.import.UC1` + `onboarding.enterprise-workforce-standard.UC1` share **one** blocker: a checked-in **résumé fixture**. `playthroughs/fixtures/` has been reserved and EMPTY since spec §5.4 — **no shipped Playthrough has ever exercised a file upload** — so the first pays the fixture *and* the real-file-chooser pattern. Land them **together** so that cost is paid once. | a later tik of M256 |
| `PT-M257-self-evaluation` | **Re-home recommendation (iter-09 D39).** `profile-skills.self-evaluation.UC1`'s M206 reservation is WEAK: its curated final is persist-then-observe (`user_skill_evidences.user_level`), needing no LLM, no integration, no fixture — exactly the MUTATING shape clause 2 hunted for four iters. Re-homing a reservation is a **roadmap decision**, so it is recorded as a recommendation, not actioned. | M257 (user/roadmap call) |
| `PT-M257-talk-to-data` | `talk-to-data.query.UC1` — real + wired (`app/internal/askengine`), but needs the `ask_*` tables migrated on the demo **and live Bedrock credentials**. An unavailable credential is not something an iter can fix; it also belongs in the separately-budgeted integration lane, not the timed median. | M257+ |
| `FIX-M256-cockpit-manifest-drift` | **NEW (iter-10 D41), and it BLOCKS D-v28-5.** `run-playthroughs.sh --reset` re-exports `fake-fapi-roster.json` (M211 iter-16) but **not** `cockpit-manifest.json`, so on any Playthrough-reset demo the cockpit lists heroes that no longer exist and every seat selection **silently** falls back to the last-active seat (`clerk-frontend/server.go:347-349`). 23 Playthroughs stay green while the human-facing cockpit is entirely stale. Fix shape: re-export the cockpit manifest alongside the roster so the two artifacts move together — **verify on a live bring-up**, and do not regress the roster refresh. Separately consider making the unknown-key fallback **loud** on a demo. | **next iter** |
| `D-v28-5-cockpit-logout` | **BLOCKED on the above (iter-10 D42).** A gate clause in its own right, still unfixed after 10 iters — but no longer merely unstarted: it is **not measurable** until the cockpit can select a current hero. The double-click symptom is plausibly a *consequence* of the drift (a presenter who does not get the hero they clicked clicks again) — plausible, **unmeasured**. Re-measure on a cockpit whose manifest matches its roster BEFORE designing a fix. By the user's explicit call it gets **no Playthrough**. | after the drift fix |
| ~~`BLOCKED-M256-refusal-surface`~~ | **DONE (iter-11) — clause 2's `blocked` sub-clause MET, 0 → 1.** Answered exactly as the routing predicted the *surface* but not the *mechanism*: the deny modal was right, and the way to reach it was a **seed** change, not a test one — `sim_feature_disabled: true` on `pt-world` Org B withholds the g3 grant so Sentinel's own enforcer refuses (`pt-aisim-org-feature-blocked`). It also exposed that `--reset` had been leaking g3 grants for four releases. Original note kept for the record: | closed iter-11 |
| ~~(original)~~ | Clause 2's `>= 1 blocked` outcome, then **0**. `actor.entitlement` is declared-only (iter-01 D4), so it needs a REAL refusal. Strongest candidate, and the locator already exists: `SimulationPage.orgMemberCannotStartModal()` — which `pt-aisim-chat-launch` currently asserts **ABSENT**. Seed a member whose org lacks the `FEATURE_JOB_SIMULATIONS` g3 grant and the deny modal becomes the outcome (M203 iter-05 documented the mechanism from the other direction). | a later tik of M256 |
| ~~`ONBOARD-M256-assessment`~~ | **DONE (iter-07 D28) — trigger NOT tripped.** The audit's F5 conflated org membership with onboarding completion. Onboarding is **UNBUILT, not impossible**; clause 3 keeps its full scope. Build routed as `ONBOARD-M256-build`. | closed iter-07 |
| ~~`FENCE-M256-bounded-interaction`~~ | **DONE (harden pass 1, commit `cfaa1a9`).** `playthroughs/e2e/tests/bounded-interaction-fence.unit.spec.ts` + **6 sites bounded**. **iter-12's framing was wrong in both clauses, and measurement says so:** it counted *"four unbounded `waitFor` calls … none inside a retry loop, so none proven harmful"* — but D25's root cause was a **`click()`**, not a `waitFor`, and **two of those four ARE the guard of a retry loop**, which is the same unreachability through a different door. Playwright's action default is `0` (no timeout), so both loops written *after* D25 (`openAssignBuilderForFirstAssignable`, iter-03; `clickUntilDialog`, iter-04) reproduced D25's shape exactly — unbounded click inside, unbounded wait guarding, only `dialog().waitFor` bounded — each declaring a 30 s budget enforceable from neither position. **A fence scoped to the spelling of the bug you already found is the mistake iter-03 corrected for `networkidle`.** Clicks moved INSIDE the `try` (D25's yield-to-next remedy; a bounded click outside it aborts the loop on first detach — same outcome, third road). Exception boundary **enumerated**: 28 straight-line sites deliberately out of scope, with the reason. RED-proven twice — D25's historical snippet (both shapes, not one) and a live revert of the org-admin click. | closed harden-1 |
| ~~`FLAKE-M256-assign-under-bloated-policy`~~ | **CLOSED (iter-13 D65) — and the hypothesis was WRONG, which is why measuring it first was right.** It recurred on iter-13's Phase-D run 3. The policy was measured **clean**: `g2 = 191`, `g3 = 171` for 191 memberships (Org B's 20 correctly withheld), **0 orphans**. Two more plausible mechanisms were also refuted by probe (the antd `maskClosable` re-click *throws* on the mask and the modal survives; `press('Enter')` with the dropdown closed leaves `aria-expanded` true and `dialogCount` 1), plus the modal **surviving 151 s unattended**. **The trace named the real cause:** the assign modal is **ROW-SCOPED** (*"Assign Skill Path to `<member>`"*, rendered by the member row's action cell), so a members-table re-render **unmounts** it — and it opened 2.2 s after the first row painted, while the table was still settling. The remaining time decomposes exactly as the retry ladder's own bounds: 3 × 15 s + 20 s + 15 s = **84 s, the reported duration**. *Every bound was correct; bounding makes a stuck attempt yield, it does not make a **dead subject** detectable.* Fixed three ways (a `dialogIsOpen()` re-open guard per attempt · `waitForMembersTableSettled()` · `openBuilderAndPickSkillPath()` so the member named is the one the ACCEPTED builder targets), recovery proven deterministically, regression-pinned in the bounded-interaction fence, and the 3× gate re-run **0 flake**. | closed iter-13 |
| `DOC-M256-claudemd-pt-count` | **NEW (iter-11), housekeeping.** `CLAUDE.md` still reads "18 live Playthroughs"; it points at `playthroughs.md` as authoritative, which now reads **24**. Reconcile ONCE at milestone close rather than on every increment (the count has moved five times inside this milestone). | milestone close |
| `FIX-M257-content-stories-pair-count` | `run-content-stories.sh` re-implements `buildPairs()` inline, omits `manager_presence_only`, computes 47 against the pinned 45 and `sys.exit(2)`s — the content-stories sweep refuses to start (audit Gap 7). | M257 / M258 (they compose the sweep) |

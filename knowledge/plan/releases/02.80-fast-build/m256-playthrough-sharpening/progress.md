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

**Post-iter-08 (re-measured on the GROWN denominator, 20 non-studio) — THE CURRENT FIGURE:** median per
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
| `NEGCTL-M256-cross-vantage` | **Gate-critical, replaces the above (iter-07 D30).** Clause 2's negative controls stand at **5 of 21**; the 5 mutating ones get it free from their pre-state read (D22), and the **16 non-writing** ones need a different mechanism. Run each Playthrough's **own final locator against a contrast vantage** — a hero/org for whom the outcome legitimately does not exist. Real absence, app stays alive, and it proves **WHICH** data not merely **THAT** data (the M219 lesson, per-Playthrough). Cost stated honestly: **O(tests), not O(surfaces)** — budget it across more than one tik. | **iter-08+** |
| `FIX-M256-studio-false-green` | **RE-AIMED (iter-07 D31) — the old diagnosis was FALSE.** iter-02 blamed the route's `Simulation Advanced Builder` header; a 5-minute poll of the real journey shows that string **NEVER renders**. The matcher actually fires on the designer's **empty section scaffolding** ("Scenario Characters" / "Mission Tasks" headings) at **+2.1 s**, before the LLM draft populates it. **Deleting the header alternative is a NO-OP — do not ship it as a fix.** Assert a **POPULATED** section instead (a character card / a non-zero `designer.actors.counter.label` count). Evidence is attached to the locator in `studio-builder-page.ts`. Until it lands, both studio Playthroughs stay `@pt-mutation: UNKNOWN`, never `MUTATES`. | a later tik of M256 |
| `DOC-M256-llm-lane-premise` | `playthroughs.md` § the `studio` product + the M256 overview + D-v28-9 all describe the advanced builder as reaching a generation completion boundary. Correct **once**, against the fixed behaviour. **Still not dischargeable (iter-07):** a section *heading*'s presence does not answer *"did the generation complete on this host?"*. Measuring section **CONTENT** answers the fix and the doc premise together — keep them one piece of work. | the same tik as the fix |
| ~~`ONBOARD-M256-build`~~ | **DONE (iter-08).** The onboarding product exists: `pt-onboarding-complete` live (mutating #6) + **all 5 curated onboarding UCs declared with written verdicts**, every verdict harness/seed work and **0 `unimplementable`**. | closed iter-08 |
| `ONBOARD-M256-import-path` | The **4 remaining** onboarding UCs, each with its specific missing piece already written into `manifest/onboarding.yaml`: a **résumé fixture** (spec §5.4's `fixtures/` dir is still EMPTY — this would be the suite's FIRST file-upload Playthrough) + an async LLM import; the **org-prepared trigger condition**, not yet identified (iter-08 measured the *import form* for a hero WITH a populated profile); an **org-less actor** (F5's one kernel of truth — needs a member-less user + a roster seat); an **Org C stage-0** seat; a **day-0 hiring-org** seat (the only onboarding UC whose final spans two apps). | a later tik of M256 |
| `VERDICT-M256-remaining-uncovered` | **Clause 3's other half.** Written verdicts still owed for `workforce.organization-feedback`, `profile-skills.import`, `talk-to-data.query`, plus the 5-release-old **M206/M207** reservations. `manifest/onboarding.yaml`'s TODO block is the template — a verdict names the specific missing piece, which is what proves it is not `unimplementable`. | next iter |
| `D-v28-5-cockpit-logout` | **A gate clause in its own right, still UNSTARTED across 8 iters.** The cockpit logout / Back-to-Cockpit double-click defect. Same seat-switch machinery every Playthrough drives (`hero-login.ts` / the M37 handshake); by the user's explicit call it gets **no Playthrough**. | next iter |
| `BLOCKED-M256-refusal-surface` | Clause 2's `>= 1 blocked` outcome, still **0**. `actor.entitlement` is declared-only (iter-01 D4), so it needs a REAL refusal. Strongest candidate, and the locator already exists: `SimulationPage.orgMemberCannotStartModal()` — which `pt-aisim-chat-launch` currently asserts **ABSENT**. Seed a member whose org lacks the `FEATURE_JOB_SIMULATIONS` g3 grant and the deny modal becomes the outcome (M203 iter-05 documented the mechanism from the other direction). | a later tik of M256 |
| ~~`ONBOARD-M256-assessment`~~ | **DONE (iter-07 D28) — trigger NOT tripped.** The audit's F5 conflated org membership with onboarding completion. Onboarding is **UNBUILT, not impossible**; clause 3 keeps its full scope. Build routed as `ONBOARD-M256-build`. | closed iter-07 |
| `FENCE-M256-bounded-interaction` | Generalise iter-06 D25: a source-scan fence asserting no unbounded `click`/`press` sits inside a retry loop in the harness. The defect class is real (a 245 s in-suite timeout that passed in 6.0 s alone) and the fix was per-site; the fence is what stops the next one. | a later tik of M256 |
| `FIX-M257-content-stories-pair-count` | `run-content-stories.sh` re-implements `buildPairs()` inline, omits `manager_presence_only`, computes 47 against the pinned 45 and `sys.exit(2)`s — the content-stories sweep refuses to start (audit Gap 7). | M257 / M258 (they compose the sweep) |

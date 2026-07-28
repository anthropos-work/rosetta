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
| `PT-M256-clause2-fifth-write` | **Gate-critical (iter-05 D20).** Clause 2 needs >= 5 mutating and stands at 3, with both org-admin candidates blocked on platform behaviour. Choose the 5th/6th deliberately: the `Remove Tags` bulk action (may avoid the modal entirely), the profile self-evaluation write (also needs a clause-3 verdict), or an onboarding completion. | next iter |
| `FIX-M256-studio-false-green` | `advancedDesignerRendered()` matches the route's own `Simulation Advanced Builder` header, so `pt-studio-advanced-generate` passes ~1.3 s before the generation completes (iter-02 D6). Assert a post-draft-only landmark and prove it RED with no generation. **This IS a clause-2 negative control**, not a side errand. | the clause-2 tik of M256 |
| `DOC-M256-llm-lane-premise` | `playthroughs.md` § the `studio` product + the M256 overview + D-v28-9 all describe the advanced builder as reaching a generation completion boundary. Correct **once**, against the fixed behaviour. | the same tik |
| `FIX-M257-content-stories-pair-count` | `run-content-stories.sh` re-implements `buildPairs()` inline, omits `manager_presence_only`, computes 47 against the pinned 45 and `sys.exit(2)`s — the content-stories sweep refuses to start (audit Gap 7). | M257 / M258 (they compose the sweep) |

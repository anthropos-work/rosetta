# M256 · iter-03 — decisions

## D9 — Phase A FALSIFIED iter-02's D8: `networkidle` *is* the lever; `storageState` reuse is de-scoped

iter-02 D8 inferred, from `pt-aireadiness-manager-dashboard` costing 1.86 s warm, that `networkidle` was
cheap on localhost and that the login handshake was the median's driver. **Phase A measured the legs
directly and the inference was wrong in one direction and right in another:**

| Leg | Measured |
|---|---:|
| `selectSeat` `POST /v1/demo/select` | **44 ms** / 9 ms — negligible |
| `goto /profile`, `waitUntil: 'networkidle'` | **2854 ms** |
| `goto /profile`, `waitUntil: 'domcontentloaded'` | **423 ms** (+ ~940 ms for the content assertion) |
| restored `storageState`, `domcontentloaded`, no handshake | 243 ms (+ 913 ms) |

**Where D8 was wrong:** `networkidle` costs **~1.5 s per test** on `/profile`, locally. D8 generalised from
a *single route* (`/ai-readiness`) that happens to settle fast; the cost is **per-route**, not per-doctrine —
and the `/profile` family sits in the middle of the distribution, which is exactly where a median moves.

**Where D8 was right:** the handshake is not worth attacking. With `domcontentloaded` the whole
handshake+SSR is **423 ms**, and `storageState` reuse saves only **~200 ms** beyond it
(1156 ms vs 1363 ms end-to-end).

**De-scoped, with a measurement rather than a deferral.** ~200 ms/test does not justify the machinery, and
the machinery carries a **false-green hazard**: Clerkenstein holds one global seat and `handleMe` resolves it
with **no cookie input** (iter-01 D1), so a restored state whose seat has moved on renders *a* hero — the
wrong one — and a render-presence assertion still passes. That is iter-02 D6's failure mode, deliberately
manufactured for 200 ms. Not built, and not carried forward as debt: the answer is *no*.

**The general lesson (why Phase A existed):** a doctrine can be right about a mechanism and wrong about a
magnitude. D8 read one route's timing as the doctrine's cost. The 30-second leg probe cost less than the
machinery it prevented — and it also rescued the lever it had wrongly written off.

## D10 — The redundant `networkidle` settle was accidentally LOAD-BEARING; the fix stays semantic

Removing the redundant `waitForLoadState('networkidle')` from `assignments-page.openSkillPathsTab` turned
`pt-assignment-assign` red once (**240 s timeout** at `dialog().waitFor`, then green on the next two runs).

**Root cause — a hydration race, not a missing settle.** Playwright's `click()` auto-waits for
*actionability* (visible / stable / enabled / hit-target), but actionability is a DOM property, and the
"Assign Skill Path" handler is attached by **React hydration**, which can land after it. A click in that
window is delivered and does nothing: the modal never opens, and the test burns its whole budget waiting for
a dialog no click ever requested. The `networkidle` settle had happened to outlast hydration, so it **masked
a real flake for two releases**.

**The fix is a bounded click-RETRY** (re-issue until the dialog appears, 30 s ceiling, a loud error naming
both candidate causes) plus a **semantic table-row hydration gate** replacing the network heuristic.
Restoring the settle was rejected on principle: it would re-hide a genuine flake behind a timing accident,
and **P6 says a flaky Playthrough is a defect in the Playthrough**. A test suite whose green depends on an
unrelated network heuristic outlasting an unrelated render is not measuring what it claims to.

**Recorded because it will recur:** every one of the 6 stripped settles sat behind an existing semantic wait,
so each *could* have been masking a race. Only this one was. Any future settle removal should expect to
surface one and should fix it semantically.

## D11 — Clause 1's speed half is MET on the pre-coverage suite, and must be RE-MEASURED after coverage

| | Baseline (iter-02) | After (iter-03) | Ratio |
|---|---:|---:|---:|
| **Median per non-studio Playthrough** | **3.326 s** | **2.014 s** | **0.6055×** |
| Median per-Playthrough, all 18 | 3.067 s | 1.954 s | 0.637× |
| Suite wall-clock (REPORTED) | 56.6 s | 52.0 s | — |
| Flake across 3 consecutive runs | 0 | **0** (134 passed ×3) | — |

Gate target is **≤ 0.79×**, so the speed half clears it with margin (0.61× vs 0.79×) — **but the gate says
all clauses are measured on the POST-coverage suite**, and clause 3 grows the denominator 18 → ~27. The
ratio therefore stays **provisional** until the final re-measure, under D7's pinned protocol, and the new
Playthroughs must be authored with `waitUntil: 'domcontentloaded'` from the first line — which the widened
fence now **enforces**, so a new spec cannot silently re-inherit the default.

**Attribution, so the number is not a mystery.** The gain is concentrated exactly where the defect was:
`profile-verified` 4.05 → 1.90 (0.47×), `profile-timeline` 3.91 → 1.89 (0.48×), `activity-drilldown`
6.32 → 3.35 (0.53×, the stripped settle), `profile-growth` 4.35 → 2.59, `profile-identity` 3.50 → 2.07,
`workforce-succession` 2.71 → 1.60. The AI-readiness surfaces moved ~0 (1.86 → 1.96) — they were already
fast. Nothing regressed beyond noise; the largest rise is `pt-studio-advanced-generate` 1.26 → 1.44, itself
a false green (D6) whose timing is not meaningful until it is fixed.

## D12 — The fence is renamed, because its name had become a description of its blind spot

`home-login-networkidle.unit.spec.ts` → **`networkidle-fence.unit.spec.ts`**. The old name was accurate and
that was the problem: it fenced the four `/home`-landing logins, i.e. the bug already found, and 20 sites
lived outside it. The invariant is now the doctrine — *no harness navigation may gate on `networkidle`* —
with a fail-closed floor of **18 login sites** (one per Playthrough, growing with coverage) and a
whole-harness source scan that strips comments so the doctrine can still be *documented* without tripping
the ban. Both violation shapes are mutation-verified RED.

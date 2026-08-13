# M256 · iter-03 — progress

**Type:** tik · **Active strategy:** `TOK-01` move 2 (lever substituted per iter-02 D8, then D8 itself
falsified by this iter's Phase A — see D9).

## Phase A — the diagnostic probe (the honesty gate on the claim)

Timed the login's legs directly against live `demo-2` before writing any production change:

```
selectSeat POST /v1/demo/select                 44 ms   (9 ms on repeat)
goto /profile  waitUntil: 'networkidle'       2854 ms
goto /profile  waitUntil: 'domcontentloaded'   423 ms  + 940 ms to assert content
restored storageState, domcontentloaded        243 ms  + 913 ms
```

**Outcome: the iter's own hypothesis was falsified and the previously-rejected lever was rescued.**
`networkidle` costs ~1.5 s **per test** on `/profile` (iter-02 D8 had generalised from `/ai-readiness`, a
route that settles fast — the cost is per-ROUTE); `storageState` reuse saves only ~200 ms beyond the
`domcontentloaded` fix and carries a false-green hazard. Recorded as **D9**; reuse **de-scoped with a
measurement**, not deferred.

## Phase B/C — what landed (all in `rosetta-extensions`, zero platform edits)

`playthroughs` section, commit `a3fe75a`, tag **`fast-build-m256-networkidle-fence`** — **pushed to origin**
(rung zero verified via `git ls-remote --tags origin`).

| Class | Sites | Change |
|---|---:|---|
| (a) login call sites inheriting the `'networkidle'` default | **12** | pin `waitUntil: 'domcontentloaded'` |
| (b) per-surface `goto` overrides pinning `networkidle` | **2** | `skill-path-page.ts`, `simulation-page.ts` → `domcontentloaded`, each with the doctrine + the measured 2854/423 ms in the comment |
| (c) unbounded `waitForLoadState('networkidle').catch(() => {})` settles | **6** | **deleted** — each sat behind an existing semantic wait (`waitForURL` / `waitFor({state:'visible'})`) and in front of an auto-retrying assertion |
| **Total** | **20** | **zero `networkidle` remains in harness runtime code** |

**The fence, widened + renamed** (D12): `home-login-networkidle.unit.spec.ts` →
`networkidle-fence.unit.spec.ts`. Two invariants now, both fail-closed:
1. **EVERY** `loginAsHero` pins `domcontentloaded` and never `networkidle` — floor **18** (one per
   Playthrough; it grows with coverage and cannot shrink silently). Previously: only `/home`-landing, floor 4.
2. **NO** harness runtime file contains a `networkidle` `goto` or settle — a whole-`lib/`+`tests/` source scan
   with comments stripped, so the doctrine can still be *described* without tripping its own ban. Floor: ≥ 20
   files scanned (a scanner that reads nothing proves nothing).

**Mutation-verified both ways:** re-pinning `simulation-page.ts` to `networkidle` → RED; dropping the
`waitUntil` from `profile-identity.spec.ts` → RED with the exact offender message; both restored → green.

**A latent flake surfaced and was fixed semantically** (D10). Stripping the settle in `openSkillPathsTab`
turned `pt-assignment-assign` red once — a **240 s** timeout at `dialog().waitFor`. Cause: a React
**hydration race** (the affordance is actionable before its handler attaches, so the click is delivered and
does nothing). The settle had been **accidentally load-bearing**, masking it. Fixed with a bounded click-retry
+ a semantic table-row gate — **not** by restoring the settle, which would re-hide a real flake behind a
timing accident (P6).

## Phase D — re-measure (D7's pinned protocol: n=3, `--reset` each, cold run included)

| Figure | Baseline | After | Ratio |
|---|---:|---:|---:|
| **Median per non-studio Playthrough — the GATED metric** | **3.326 s** | **2.014 s** | **0.6055×** |
| Median per-Playthrough, all 18 | 3.067 s | 1.954 s | 0.637× |
| Suite wall-clock (REPORTED, not gated) | 85.4 / 56.6 / 54.4 → **56.6 s** | 39.5 / 52.0 / 52.0 → **52.0 s** | — |
| Flake across 3 consecutive runs | 0 | **0** — `134 passed` ×3 | — |

**Gate target ≤ 0.79× → clause 1's speed half is MET at 0.61×**, with margin. Same environment as the
baseline (`Kirality-Mac-Pro-6.local`, Docker VM 9.70 GiB, `demo-2` offset 20000, localhost/http). **Not
comparable to billion's 228 s** (D-v28-12).

Per-run post-fix non-studio medians: **1.805 / 2.101 / 2.633 s** — the cold-run ordering *inverted* (the cold
run is now the fastest), which is consistent with cold-start cost having been dominated by the very
`networkidle` waits that are gone.

## Close — 2026-07-28

**Outcome:** median per non-studio Playthrough **3.326 s → 2.014 s = 0.6055×** against a ≤ 0.79× gate,
0 flake over 3 runs, from banning `networkidle` at **20 harness sites** and widening its fence from one route
to the whole harness. Phase A falsified the iter's own targeting hypothesis before any code was written, which
both rescued the lever iter-02 had written off and prevented ~200 ms of risky `storageState` machinery. A
latent hydration flake the old settle had been masking was surfaced and fixed semantically.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — clause 1's **speed half is met** (0.61× ≤ 0.79×, 0 flake) but is **provisional** until re-measured on the post-coverage suite; clause 1's flake half is met; **clause 2** (negative controls, ≥5 mutating, ≥1 `blocked`) and **clause 3** (onboarding ×5 + org-admin ×4 + written verdicts) are unstarted, as is D-v28-5.
**Phase 5 grading:** (1) gate-met: n (clauses 2 and 3 unstarted) — (2) triggered-tok: n — (3) re-scope: n (the tik delivered above its expected lift) — (4) user-blocker: n (the one red was inside this iter's planned scope and was fixed in-iter, Fate 1 — not an unrelated-suite regression) — (5) cap-reached: n (2 of 5 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D9 (Phase A falsified D8: `networkidle` is the lever, `storageState` reuse de-scoped on measurement), D10 (the stripped settle was masking a real hydration flake; fixed semantically, not restored), D11 (clause 1 speed half MET at 0.61×, provisional pending the post-coverage re-measure), D12 (the fence renamed because its name described its blind spot).
**Side-deliverables:** none — the flake fix (D10) is inside the planned scope (it was caused by, and is part of, the `networkidle` removal).
**Routes carried forward:** none new. The four items in `../progress.md` § Next-iter routing still stand, unchanged.
**Lessons:**
1. **A doctrine can be right about the mechanism and wrong about the magnitude.** `networkidle` was already
   banned in doctrine and already fenced — for one class. What was missing was a *number*. The 30-second leg
   probe (2854 ms vs 423 ms) is what turned a style rule into a 39 % median cut, and it also reversed a
   wrong write-off from the previous iter. Measure the leg before targeting the lever.
2. **Fence the invariant, not the incident.** The old fence's name — `home-login-networkidle` — was an
   accurate description of its blind spot. 20 violations lived one route away from a guard that passed. When
   writing a guard, ask what the *rule* is, not what the *bug* was.
3. **When you remove a redundant wait, expect to find what it was hiding.** Six settles were removed; one was
   masking a genuine hydration race that had been latent for two releases. Restoring it would have been the
   fast fix and the wrong one.

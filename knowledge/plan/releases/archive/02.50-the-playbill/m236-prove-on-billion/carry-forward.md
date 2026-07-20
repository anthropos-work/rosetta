---
title: "M236 Carry-Forward — Routes from prove-on-billion"
date: 2026-07-20
status: archived
close_status: closed
gate_target: "Both tabs work live on billion — all 29 landable (session x action) pairs render real, non-empty content for player + manager vantages, the academy grid renders real cards (Thread A) — reproducibly on a cold reset-to-seed, p95 click->ACCESS < 5 s for the HERO vantages only, 0 platform edits."
gate_target_denominator: 29
gate_target_denominator_history: "AUTHORED 31, CORRECTED to 29 at iter-07 on product-source evidence (the skill-path manager surface renders 'Coming soon'); see overview.md 'Gate denominator' + iter-07/decisions.md D2. The struck-through prose trail is the deliberate audit record."
gate_achieved: "gate-met — 29/29 cold on billion (see overview.md + decisions.md CLOSE-D3 for the harness-changed-after-measurement residual)"
routes_to: "v2.5 release close (/developer-kit:close-release) — M236 is the FINAL v2.5 milestone; there is no further milestone to route into"
---

## TL;DR

M236 met its gate. This carry-forward exists for the items the close **could not** dispose of, plus **two
findings discovered during the close continuation** that are larger than the milestone that found them.

The headline item is the **standing demo-stack test failures** — the repeat-deferral that held the merge. It
has now been **re-baselined under an explicit user decision** and is handed to the release close with a real
characterisation instead of a stale label. **Its fate is the release close's to take, and it is deliberately
not taken here.**

## Route 1 — the standing demo-stack test failures (the repeat-deferral)

- **Was:** *"14 pre-existing demo-stack test failures, 6 of them `pre_sha256` pin drift"*, carried across
  **10 milestones and 2 releases**, whose declared destination (the v2.4 release close) had **already fired
  without landing it**.
- **User decision (2026-07-20, in `decisions.md` → CLOSE-D2 → RESOLUTION):** **re-baseline now; decide the
  fate at release close.** Explicitly not fixed here, not dropped, and not silently rolled forward again.
- **Now:** **→ `rebaseline-standing-failures.md`** (this directory). Read that document, not the old label.

**What the release close needs to know before it decides:**

> **AUTHORITATIVE COUNT (re-measured at the v2.5 release close, 2026-07-20): 8 on macOS · 7 expected on
> Linux.** The count is **host-dependent** — always state the host OS with it. The 6 clone-dependent
> failures of the carried 14 **did not reproduce** at the close, which independently confirms the
> `stack-demo` clone set is pristine. `14` is the DIRTY-clone reading and is superseded everywhere;
> `19` (M236's own full-suite sweep) folded in 5 stack-core doc-truth-guard failures that were FIXED at
> the M236 close and are therefore not part of the standing set. See the release `metrics.json`
> → `standing_failures`.

| | |
|---|---|
| Count reproduced before any change (dirty clone) | **14** — the carried count was accurate *for that clone state* |
| Count on a clean, stable-`main` clone set | **8 (macOS) · 7 expected (Linux)** ← **authoritative** |
| Re-measured at the v2.5 release close | **8 observed, macOS** — clone set confirmed clean |
| Real product/tooling defects among them | **0** |
| `pre_sha256` pin drift among them | **0 — the carried diagnosis is refuted** |

- **6 of the 14 were never test failures** — they were a **dirty clone** (leftover applied demo patches)
  reporting itself as one. They pass once the clone is pristine.
- **The remedy the old label implied was the dangerous action.** Re-anchoring the "drifted" pins would have
  re-pinned five manifests to *patched* content, permanently disarming the drift detector for those files.
  Every pin is in fact correct at **both** the stale ref and `main`.
- **The remaining 8 are all test-side debt** — 7 stale assertions against deliberately-changed behaviour, 1
  environment-conditional test that fails instead of skipping on macOS.
- **The count is host-dependent:** 8 on macOS, expected 7 on Linux. Any future carry **must** state the host
  OS or it will drift again for exactly the reason it drifted the first time.
- **Recommended fate: Fate 1 (LAND).** All 8 are now cheap test-side edits with no product risk and no live
  stack required. Offered as input; the decision is the close's.

## Route 2 — F-M236-CLOSE-1: `/demo-up` rebuilds images from clones it never updates

**Recommend treating this as higher priority than Route 1, and separately from it.**

`ensure-clones.sh` populates a demo's clone set via `make init`, which is **skip-if-present**. There is no
fetch, no pull and no checkout for the platform repos anywhere in the bring-up. Once a clone exists it is
never advanced. Every `/demo-up` — including a full cold teardown-and-rebuild — compiles **fresh images from
stale source**, and says so in its own log: *"app already exists, skipping."*

Measured 2026-07-20, identically on **both** boxes: `app` **249** commits behind `main`, `next-web-app`
**202**, `messenger` 28, `ant-academy` 60, `graphql-wundergraph` 6, `cms` 4.

- This is the **upstream generator** of the pin-drift-looking class in Route 1. It is not part of the 14 and
  should not inherit their fate.
- `clones.lock.json` **cannot detect it**: it records `ref` as `"HEAD"` for every detached clone, so it is
  structurally unable to distinguish *deliberately pinned* from *stale by neglect*.
- It surfaced as a **user-reported symptom** — an old next-web-app left menu — which is exactly what it
  produces. Root-caused in full in `decisions.md`; the fix is a ref policy for the demo clone set.
- **Also carries a methodology warning worth propagating:** the first measurement of this drift was wrong
  (12 instead of 202) because a `git fetch` failed silently behind `2>/dev/null`. Never measure drift through
  a suppressed-stderr fetch.

## Route 3 — F-M236-CLOSE-2: the "pristine" safety net sweeps 3 manifests of ~15

`ensure-clones.sh`'s **R1** rung reverts stale patches left by a crashed build — the guarantee that a demo
clone is pristine before a build. Its hard-coded manifest list has **3** entries; `patches/` carries **15**.
Any patch outside those three, left applied by an interrupted build, survives **every** subsequent bring-up.

Both boxes were carrying leftovers, in disjoint sets: 5 in local `next-web-app`, 2 in `billion`'s
`ant-academy`. **Fix this before landing Route 1**, or the clone-dependent failures return the next time a
build is interrupted and the next reader re-derives the whole analysis.

## Route 4 — items the close routed Fate-3 to the release close (unchanged)

Carried forward from `decisions.md` → CLOSE-D1; listed here so the close has one place to look.

| item | note |
|---|---|
| `ACADEMY-M236-iter08-public-catalog-twin` — anon `/library` + `/free` render 0 cards | needs a 2nd demopatch manifest + next-web rebuild + live re-prove |
| `apps/web` client GraphQL endpoint on non-offset `:5050` | batched with the above — same rebuild, one re-prove |
| `DEF-M235-03` — M204 assign-WRITE declared TODO | inherited, correctly routed past M236 |
| CLOSE-D3 — the harness changed after the gate was measured | a cheap live re-run of `run-content-stories.sh` discharges it; not gate-invalidating |

## Provenance

- `decisions.md` → CLOSE-D1 / CLOSE-D2 / **RESOLUTION** / F-M236-CLOSE-1 / F-M236-CLOSE-2
- `rebaseline-standing-failures.md` — the full re-baseline, with resolved refs for reproducibility
- Fresh cold bring-up at stable `main` on `billion` (2026-07-20) — see `overview.md`

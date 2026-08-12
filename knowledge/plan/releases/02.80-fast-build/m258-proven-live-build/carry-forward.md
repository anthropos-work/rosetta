---
milestone: M258
title: "proven-live build (the closer)"
release: v2.8 "fast build"
close_status: closed-incomplete
close_authority: "USER RULING — clauses 1, 2, 4, 5 proven; clause 3 NOT MET and never to be recorded as met"
gate_target: "composed p50 <= 480 s over 3 consecutive cold reset-to-seed cycles, zero standing red, 0 platform edits, presenter-usable world"
gate_status: "4 of 5 clauses proven; the TIMING clause was never measured clean"
inherits_to: "/developer-kit:close-release (v2.8) — the release's final milestone, so there is no later in-release destination"
escape_hatch_deferrals: 0
last_updated: 2026-08-12
---

# M258 — carry-forward

## TL;DR

**M258 shipped the thing it promised and did not measure the number it promised.** A demo stack now comes
up **and proves itself**: `demo-4` is live, built by the fixed tooling from the newest platform mains, and
it returned `red_set: []` in its own bring-up. What was never taken is a **clean composed p50 over 3 cold
cycles** — the box is a permanently contended workstation, the user concluded that contention is not
removable, and ruled the goal achieved on the other four clauses plus a ~402 s projection.

**Read the numbers with their status attached, or do not quote them:**

| number | what it is |
|---|---|
| **840.01 s** | **instrument-rejected.** 3/3 reps `headroom=FAIL`; `buildbench` calls them "not usable measurements" |
| **401.60 s** | a **PROJECTION**, composed from separately-measured halves. Never measured as one cycle |
| **~290 s** | **deliberately not banked** — warm-cache, missing the export/unpack leg that is 46.2 % of a cold cycle, on the quietest box of the milestone |
| **179.37 s** | `batch_gate`'s own p50 — **inside M256's 200 s budget, while contended.** The batch half is not what is slow |

**There is no clean clause-3 measurement anywhere in this milestone.** Anyone reporting one is quoting a
projection or a warm cache.

## Root-cause clusters

### Cluster 1 — The timing clause has no clean measurement, and the blocker is the host

- **Affected:** gate clause 3.
- **Root cause:** the gate host is a **permanently contended workstation** shared with another user
  project's parallel campaign. `load1` minimum over ~30 minutes of polling was **11.93**, trending to
  **62.88**, against a headroom limit of 10. External, measured, and not ours to stop.
- **What was tried, and it worked as far as it could.** iter-08 ended the wait by **automating the
  trigger** rather than waiting for quiet: `autoarm-campaign.sh` samples at 15 s and fires on three
  consecutive `load1 ≤ 5.0`. It fired **91 s** after arming, into a **75 s** dip — a window a
  ~2-minute hand poll had correctly failed to find. The campaign then ran 3/3 cold with `red_count 0`
  and 3/3 `headroom=FAIL`, because the contending campaign resumed **60 s after launch**.
- **Scope to resolve:** one clean 3-rep campaign on a host that can hold `load1 < 10` for ~30 minutes.
  Not engineering — scheduling, or a different host. The auto-arm stays armed.
- **Fate:** LAND-NEXT → release close. **Not an escape hatch** — nothing is punted past v2.8 unsigned.
- ⚠️ **`billion` is DEMO-ONLY and `odysseus` is RETIRED** (`D-v28-15`). `hostprofiles/` holds only
  `billion.json` and a retired laptop's, so **`build-budget.md`'s gate clause 1 is not gradeable today**.

### Cluster 2 — Instrument items the gate reads its own number from

- **Affected:** `PROFILE-M257-provisional-fields` (inherited M255 → M257 → M258, **AGED_OUT**) ·
  `FIX-M257-campaign-kill-orphans-bringup` · `FIX-M257-sampler-disk-units-vm` ·
  `MEASURE-M257-macmini-true-idle` · `FIX-M257-frontend-floor-is-billion-shaped` ·
  `FIX-M257-image-listing-conflates-empty-and-unreadable` · `SPLIT-M258-iter09-copy-vs-reindex` ·
  `SETTLE-M258-iter13-studio-desk-cold-time` (`D75`) · `FIX-M257-census-interpreter-namespace-import`.
- **Root cause:** the harness that grades the gate has never itself been graded on this host — and the
  host it *was* built for no longer exists in the project.
- **Two that must not be quoted until they land:** `SPLIT-…copy-vs-reindex` shipped its instrument but
  **no run has produced its line**, so **do not assert whether the taxonomy replay is COPY- or
  REINDEX-bound**; and `D75`'s studio-desk **time** axis is **UNMEASURED with its original estimate
  refuted** (it applied `billion`'s x86_64/containerd s/GB to an arm64/overlayfs host — the exact
  cross-host error `build-budget.md`'s opening rule exists to prevent). **Only the 350 MB space win is
  measured**, and the corpus now says so.
- **Fate:** LAND-NEXT → release close.

### Cluster 3 — Chronics that crossed ≥ 2 milestones

- **`FIX-M257-content-stories-pair-count` — LANDED at this close, and its description REFUTED.** Carried
  from the M256 audit through M257 as *"the sweep refuses to start"*; measured against the canonical
  preset it yields **45 = the pin**, with **no row carrying both flags**. `buildPairs()` sets
  `has_manager_view: false` on a manager-presence-only row, so the missing clause never changed a count.
  **The sweep was never blocked.** Three iterations verified the defect by *reading* the two
  implementations and inherited the panic with it — **a code read is not a measurement.** The clause
  landed anyway, because the exclusion depended on a second field happening to correlate.
- **`F2` — `ptvalidate` invoked nowhere outside its own tests. LAND-NEXT, and the reason is not time.**
  Wiring it as a *binding* pre-flight changes the gate behaviour of the runner **every bring-up now
  depends on**; validating that means driving a full suite, which resets `demo-4` — the user's only stack
  and this milestone's binding end state. Advisory-first is not offered: that is a partial landing.
- **`RATCHET-M257-literal-ceilings-breached` — pre-existing breach of 8, and it did not grow here.**
  Measured pristine **248** / HEAD **249** against a ceiling of **240**; the harden added exactly one
  literal, in the instrument's own sanctioned `dated` class, and this close added **none**. **No ceiling
  was raised at any point.** Either attribute and raise with a reason, or pay the debt down.

### Cluster 4 — Known pre-existing failure sets (measured, not inherited on faith)

- **~46 rext-internal census failures** and **9 `demo-stack` live-clone failures**.
- **Re-attributed at this close** against a pristine `git archive HEAD` extract **re-sited at the same
  depth as the real clone** — the first siting SKIPPED 5 tests rather than running them, and said so,
  which is the same trap the harden pass hit and named. Pristine reports **12**, current **11**, and the
  diff by name is **empty in the introduced direction**: **nothing was introduced.**
- **One failure WAS introduced by this close and fixed inline:** a prose comment it added contained a
  phrase `test_frontend_build.py` counts, taking a fence 3 → 4. **A comment can break a fence that counts
  a phrase** — the sibling of this milestone's own *a fence can be satisfied by its own comment*.
- **Do not upgrade this to "swept clean".** The full `stack-core` sweep **did not complete** on this host,
  consistent with its documented behaviour across both milestones; scoped runs plus pristine-extract
  attribution were used instead.
- **Fate:** LAND-NEXT → release close.

### Cluster 5 — Structural, lower severity

`FIX-M258-iter03-guard-scans-its-own-scratch` (+ its `test_fence_provenance` sibling) and
`ROUTE-M258-iter09-literal-ratchets-scan-the-demo-clone` share one root: guards that census
`stacks/demo-N/clones/**` — the *platform's* source inside a demo's ephemeral clone. **Consequence: any
ratchet figure measured on a box that has run a demo, without excluding `stacks/`, is not a measurement of
this repo.** · `ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone` is structural
(`ant-academy` runs natively, so G5 has no ephemeral clone to discard) and **not repaired: reverting
tracked files is a forbidden op, and that clone is what the user's stack serves from.** ·
`FIX-M258-iter14-purge-leaves-276MB` priced and deliberately not taken, validated to within 1 MB
(predicted ≈276, measured **277**) · `TARGET-M258-iter13-browser-only-deps` (the 838 MB production-dependency
tail) · the iter-17 REPORT/ROUTE set · `FIX-M258-iter15-hiring-under-set-dressed` **does not reproduce**
(50 rows vs 38), re-scoped to WATCH · **`FIX-M257-dockerignore-env-pattern-unpaired`** — ⚠️ **the tidy
one-line fix BAKES THE REAL CLERK KEY**; needs a re-include and a real build to validate.

### Not ours to fix

**`make bootstrap-dev` is BROKEN in the platform** (iter-18 `D92`) — reported, **0 platform edits**. Owned
by [`platform-defect-register.md`](../../platform-defect-register.md), which is that class's existing home.

## Projected post-resolution state

One clean 3-rep campaign on an uncontended host is the whole of clause 3. The arithmetic that exists
already fits: **iter-05's gateable single-box bring-up 247.79 s + the batch half ~129–179 s ≈ 377–427 s**,
against a 480 s ceiling — and `LEVER-M257-L5-setdress` is **still unspent**, now with a named target (the
taxonomy replay, **~88 %** of `set_dress`) rather than an opaque span. Nothing in the composition needs
engineering; it needs a quiet box.

## Cross-references

- [`progress.md`](progress.md) § Gate Outcome Ledger — the close record
- [`deferrals-audit.md`](deferrals-audit.md) — every item, its fate, and the blocking-grading representation
- [`metrics.json`](metrics.json) — the machine-readable numbers, each with its status attached
- [`hardening-ledger.md`](hardening-ledger.md) — 5 passes, STABILIZED
- [`../m257x-platform-realignment/carry-forward.md`](../m257x-platform-realignment/carry-forward.md) —
  the inherited clusters, of which cluster 1 (the production-bucket pointer, contained in CODE and on no
  running stack) remains the standing safety item

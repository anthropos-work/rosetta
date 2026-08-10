---
iter: 258
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-10T13:50:17Z
closed: 2026-08-10T14:08:12Z
---

# iter-258 — prove the advance BUILDS: a cold `demo-2` on the advanced refs

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them*. This tik operates under the user's binding closing
condition `D-M257x-256-1`, which `TOK-08`'s own framing accommodates: a bring-up is the **census** of
whether the platform assembles, where every prior clause-1 reading has been a **sample** taken on a stack
built at a ref nobody re-verified.

## Step 0 — Re-survey before targeting (mandatory), and it moved three things

`TOK-08`'s next-tik direction is stale by construction — the user's 2026-08-10 closing condition
(`D-M257x-256-1`) supersedes it, and iter-257 named
`ROUTE-M257x-256-the-advance-is-unproven` *"the milestone's critical path"*. The re-survey confirms the
route is still open and still meaningful, and corrected **three** briefed facts before any work:

1. **The briefed *"demo-1 is the only compose project that has ever existed on this box, so slot 2 is
   free"* is right in its conclusion and wrong in its premise.** `demo-stack/stacks/` holds **`demo-2`
   and `demo-4`** skeleton dirs (Jul 31, 3 files each: `.env.demo-N`, `docker-compose.demo.yml`, and for
   demo-2 an empty `data/`), and `demo-stack/stacks/registry.json` records **slot 2** — while omitting
   the live `demo-1`. That file is the **pre-M12 demo-only legacy registry**, kept for provenance;
   `stack_registry.py:51` puts the canonical unified allocator at
   `stack-core/.stacks/registry.json`, which holds **`demo-1` only** (`adopted: true, status: up`).
   Slot 2 **is** free in the file that decides allocation. Recorded because a future reader will open
   the wrong registry exactly as this iter first did.
2. **The clone advance is REAL and on disk.** `stack-demo/{app,next-web-app,ant-academy}` are at
   `3eaadae68` / `19423a1fb` / `249430c3`, matching `clones.pin.json` byte-for-byte at all six repos.
3. **`demo-1`'s green verdict provably does NOT cover the advance** — the fact that makes this iter
   necessary rather than redundant. `demo-1/autoverify.json` is
   `{"green":true,"warnings":0,"ts":"2026-08-06T10:21:56Z"}`, but its **build-scratch clone**
   `demo-1/clones/app` is at **`ad9f3c498`** — the pre-advance ref, which `CLAUDE.md` records as **28
   commits behind** `3eaadae6`. The stack that is green is not the stack the corpus now describes.

## Cluster / target identified

`ROUTE-M257x-256-the-advance-is-unproven`. iter-256 fast-forwarded three clones and closed with the
explicit disclaimer *"**Not claimed:** that anything builds."* Nothing since has built them. Under
`D-M257x-256-1` the milestone cannot close until they do — **demo first, then dev**.

Target: a **cold `demo-2`** — a slot with no containers that have ever existed — on the advanced refs.
Non-destructive by construction: `demo-1` is never touched.

## Hypothesis

The advanced refs assemble into a working demo stack. The live risk is the **`app` advance's new
migration** plus a terraform commit reading *"the backend migration pipeline has been a silent no-op
since the atlas 0.7.0 bump"* — so a migration surprise is treated as a live hypothesis, not a remote one.

## Expected lift

Boolean, not timing: `up-injected.sh 2` exits 0 **and** `stacks/demo-2/autoverify.json` reads
`green:true`. That is clause 1's unit of evidence at the advanced refs — the first such reading that
exists.

## THE PIN DECISION — taken deliberately, recorded, not defaulted

Pre-flight rung zero **passes**: `.agentspace/rext.tag` = `fast-build-m257x-iter-101` and that tag **is on
origin** (`0011c10aba0ff0950341cb410265ee59d070afe3`; origin carries 470 tags). The M236
unreachable-pin failure is not present. But the pin is **157 iters stale**, and the rext authoring copy is
**5 commits ahead of origin/main**, unpushed and untagged.

**DECISION (`D-M257x-258-1`): accept the stale pin, deliberately, on experimental-control grounds.**

Three reasons, in descending weight:

1. **It isolates the variable.** The green `demo-1` (2026-08-06T10:21:56Z, 0 warnings) was produced by
   **this exact tooling version** at app `ad9f3c498`. Holding the tooling constant makes the platform-ref
   advance the **single changed variable**. Bumping the pin would change two at once and make a failure
   unattributable — which is the whole reason this route is open.
2. **Bumping it would not change what gets built.** iter-257 measured that `DEMO_ADVANCE_CLONES` defaults
   to **`0`** and nothing outside `ensure-clones.sh` sets it, so a default bring-up applies **no pin at
   all** — it builds the clones as checked out. The newer `clones.pin.json` that iter-256 advanced is
   therefore not consumed by this bring-up either way.
3. **Publishing a tag is a side-effect, not a control.** Tagging the authoring copy mid-experiment would
   put 157 iters of untested-in-a-bring-up tooling under the run.

**The gap is stated rather than closed:** iter-256's `clones.pin.json` advance and iter-257's
`clone_pin_guard` arm D are **absent** from the tag this demo clones. If the bring-up fails inside rext
tooling rather than inside platform code, the stale pin becomes a candidate cause and must be re-tested
at a fresh tag before the failure is attributed to the platform. Routed as
`ROUTE-M257x-258-the-pin-is-157-iters-stale` regardless of outcome.

## Pre-registrations — sealed in this iter's FIRST commit, before the bring-up runs

Per `TOK-08`'s standing discipline. Three of the five predict **against** this run's hope.

| | claim | prediction |
|---|---|---|
| PR-1 | `up-injected.sh 2` reaches `autoverify green:true` on the **first** attempt | **REFUTED** — clauses 1–2 are already `MET WITH DISCLOSURE` that a freshly built stack failed its first full run in **2 of 2** attempts, and this run adds an unproven 28-commit advance on top |
| PR-2 | the `app` advance's new migration applies cleanly | **HOLDS** — the terraform commit names a **prod CI pipeline** no-op, not a broken migration file; the local `make migrate`/atlas path was never the thing that no-op'd |
| PR-3 | the iter-101 tooling builds **no** decommissioned service | **HOLDS** — `INJECT_CANDIDATES="app cms jobsimulation"` at `:216-217` is filtered at `:1682-1703` by a derivation from *the platform compose's own build set*, so cms/jobsimulation drop out at platform `0c91421` |
| PR-4 | `demo-1` is bit-untouched by a `demo-2` bring-up — 11 containers, same container IDs, same uptime lineage | **HOLDS** — compose is `-p demo-2` scoped |
| PR-5 | the absent host profile produces an **advisory** warning, not a block | **HOLDS** — `build-budget.md` `D-M255-1`: the identical assert hard-fails in `buildbench` (a gate) and is advisory in the bring-up (an operator). **Verified, not assumed** — if it blocks, that is itself the finding |

## Phase plan

- **Phase A** — seal this overview (first commit). Re-verify the three environmental facts the run
  depends on and that `demo-1` is healthy *before* starting.
- **Phase B** — cold `up-injected.sh 2`, logged to a file, run to completion. Heartbeat it.
- **Phase C** — read the verdict: exit code + `autoverify.json`. If RED, root-cause to a named layer
  (platform code / migration / rext tooling / host) rather than to a symptom.
- **Phase D** — re-measure: `demo-1` untouched check, PR adjudication, close.

## Escalation conditions

- A failure that is **attributable to the stale pin** → re-test at a fresh tag before blaming the
  platform; do not publish a platform verdict from a tooling failure.
- A failure requiring a **platform-repo edit** to clear → `user-blocker`. This milestone is 0 platform
  edits by construction.
- Anything that would require touching `demo-1` → hard stop, `user-blocker`. Non-negotiable.

## Acceptable close-no-lift outcomes

A bring-up that **fails with a root cause named and cited** is a complete iter under this protocol: the
route says *the advance is unproven*, and proving it **broken** discharges that route exactly as proving
it green does. What would NOT satisfy it is a bring-up that fails and is reported as "slow" or "flaky"
without a named layer.

## Timing discipline (binding for this iter)

The host is **permanently contended** by third-party work the user cannot stop (an `anima8` audio pytest
battery was resident at open; load avg 2.79 / 4.84 / 7.79 at 13:44Z against 10.27 / 12.32 / 13.61 at the
orchestrator's pre-flight 17 minutes earlier). **A boolean survives contention; a timing does not.**
Every duration in this iter is labelled **CONTENDED** and is **not** a host baseline. `billion`'s
666.29 s p50 and its phase attribution are **not portable here** — this box is `overlayfs`, which pays no
image-unpack leg at all, and that leg is 46.2 % of `billion`'s budget.

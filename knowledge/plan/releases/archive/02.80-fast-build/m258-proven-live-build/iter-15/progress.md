# M258 iter-15 — progress

**Type:** tik · **Active strategy:** `TOK-02`, `TIK-C` — `END-M258-one-stack`, the milestone's binding
end state and simultaneously its largest Class A reclaim.

Measured 2026-08-12 from 10:40Z on `macmini` (Apple M4 Pro, arm64, Docker Desktop VM, overlayfs),
`load1` **19.7–26.3**. Every figure carries that environment.

## Phase A — clearing the way (two things that could have wrecked the transition)

**The clause-3 waiter was still armed** (`autoarm-campaign.sh`, pid 7619, sampling every 15 s, set to
fire a 3-rep cold campaign **that tears `demo-1` down**). Firing during `END-M258-one-stack` would have
destroyed the stack being built and could have left the user with **no** stack at the exact moment this
milestone promises him one. Stopped deliberately, reason written to `.autoarm-outcome` (`D72`). For the
record it never came close: `load1` minimum since arming **14.21** against a threshold of **5.0**.
**Clause 3 remains NOT MET and is never to be recorded as met.**

**Stack ownership was re-resolved from docker** rather than inherited: `iter-07 D23` had `demo-1` on the
authoring clone, but both demos now bind out of the **consumption** clone (`D73`). The authoring clone
still holds a stale 263 MB `stacks/demo-1/` — the two-clone confusion that produced iter-06's `D19`.

## Phase B — the pin guard fired, on my own error

```
✗ FATAL: rext pin mismatch.
    the consumption clone is at : fast-build-m258-iter-14
    .agentspace/rext.tag pins   : fast-build-m258-iter-09
```

I had checked the clone out at the new tag and not updated the declaration. **Both remedies the guard
offers would have been wrong** — reverting the clone reverts the feature under test (the M236 shape),
and `DEMO_ALLOW_UNPINNED_REXT=1` disables the check that just worked. The right fix is the third:
update the declaration, because `rext.tag` is an **intent**, not a lock (`D71`).

Feature-presence then verified rather than the tag alone (M236 rung zero): `studio-desk.Dockerfile`
present, `down -v` present, all three UI builds on `-f "$dockerfile"`, and **no `Dockerfile.dev` target
left anywhere**.

## Phase C — newest platform repos

| repo | before | after |
|---|---|---|
| platform | `0c91421` | **`766df6c`** |
| app | `3eaadae68` | **`c52dbc51e`** (+76) |
| next-web-app | `19423a1fb` | **`3379072e9`** (+59) |
| ant-academy | `249430c3` | **`7ae25e95`** |
| sentinel · studio-desk | already current | — |

Fast-forward only; no clone carried local commits.

## Phase D — the build (peak state, and the release's image work in one table)

Peak, with three stacks up **and** `demo-3` building — the worst case the end state exists to undo:

```
Images 25.71 GB · Build Cache 27.77 GB (18.96 reclaimable) · Volumes 9
Docker.raw allocated 53.84 GB   ·   host stack dirs 4.3 GB   ·   free 170 GiB
```

The UI tier, same box, same day — **L1 and `TIK-A` together**:

| image | demo-2 (pre-L1) | demo-3 (new) | |
|---|---|---|---|
| next-web | 4.04 GB | **417 MB** | 9.7× |
| hiring | 3.94 GB | **380 MB** | 10.4× |
| studio-desk | 1.70 GB | **1.35 GB** | 1.26× |
| **UI tier total** | **9.68 GB** | **2.15 GB** | **4.5×** |

The `TIK-A` lever shipped in a real bring-up — the log names it
(*"rext studio-desk.Dockerfile, multi-stage prune-and-copy"*) and the **anti-silence `vite` assert ran
and passed**. ⚠️ Its export leg here was 1.5 s / unpack 0.3 s because the layers were already cached
from iter-13's probe build: **not a cold-export datum**, and not quoted as one.

## Phase E — verification, and the distinction that decided the milestone

`autoverify` came back **green: false, warnings: 1** — and the one warning is a *seeded-cardinality*
shortfall, not a broken stack: **the hiring org has 38 candidate sessions against a threshold of ≥40**
(positions 5/5 OK). Everything else passed: taxonomy **42,790**, all liveness+readiness probes,
demo-patches all applied (none refused), *"frontend builds: ok (the running images are this run's)"*,
cockpit answering, Clerkenstein answering, academy rendering real cards, **10/10 containers**.

The **batch gate** then ran the full suite and returned **`verdict: red`, `red_count: 15`**
(`batch_seconds` 629, `restore_seconds` 29, 15 passing / 15 failing / 1 unimplemented of 31).

⚠️ **Read what that red set is a verdict ON, because it decided whether the user keeps a usable stack.**
The batch is `reset-to-seed → pt-world → suite → restore`. The Playthroughs run against **`pt-world`, the
DECOUPLED TEST SEED — not the presenter demo world**; the restore leg then puts the demo world back. So
the red set grades *test* data, and the stack a presenter would drive is what the restore produced. It
was measured directly rather than inferred:

```
restored presenter world:  orgs=4  users=591  memberships=591  skills=42790
cockpit seats:             12 across 4 stories   (deep links 13)
   AI Transformation & Reskilling 3 · SDR Onboarding & Ramp 3
   AI Readiness Diagnostic 3 · Candidate Hiring & Comparison 3
live surfaces:  cockpit :37700 -> 200 · web :33000 -> 307 · studio :39000 -> 302
                backend :38082/api/health -> 200
```

That is the documented healthy shape exactly (iter-06/07: *"4 story orgs / 591 users / 12 cockpit
seats"*), and it satisfies the gate's clause 5.

**The red set is escalated, not swept** (`D-v28-3` requires it). Its 15 entries are **not one thing**:
**4 are plain timeouts** (`org-admin` ×3 at 60 s, `pt-activity-drilldown` at 20 s) taken at
**`load1` 26–33 against 10 cores with `retries: 0`** — the exact condition `iter-07 D28` says
manufactures false reds. The other **11 are DATA-SHAPE assertions** (*"the roster rows belong to THIS
org … @pt-meridian-labs.com"*, *"the org's shared positions are exactly 5 … fewer means a starved
SIMULATION_TYPE_HIRING pool / cold snapshot cache"*, *"at least one member with no skill-path
assignment"*), and they agree with autoverify's own under-set-dress warning. **Neither cause was
confirmed here** — no `SQLSTATE 42P01` / `does not exist` appears anywhere in the bring-up log, so the
"newest platform moved a table the seeder writes" hypothesis is **unproven, not established**, and must
not be reported as diagnosed. Routed with both candidate causes named.

## Phase F — the teardown, in the mandatory order

Order was enforced **in code**, not by discipline: `teardown-others.sh` refuses unless `demo-3` is up,
and re-checks it between every step. Within the sequence the order is *least-valuable-to-the-user first*,
so the stack he validates on is the last to go:

| # | stack | why | after |
|---|---|---|---|
| 1 | **demo-1** (mine) | zero cost to him; frees resources for the running batch | 3 stacks |
| 2 | **dev** (`anthropos`) | | 2 stacks |
| 3 | **demo-2** (the USER'S) | **last**, on pre-L1 images, replaced by demo-3 | **1 stack** |

A real finding fell out of it: **`compose down` FAILED on both demos** — *"service sentinel has neither
an image nor a build context specified: invalid compose project"* — and the **label sweep is what
actually removed all 11 containers each time**. The M257x iter-55 defensive path is not belt-and-braces;
it is the thing that works when an older stack's generated compose no longer parses.

**The reclaim, measured `system df` before/after — never from the SIZE column (`D53`):**

| axis | peak (3 stacks + demo-3 building) | final (1 stack) | Δ |
|---|---|---|---|
| **Docker.raw allocated (real SSD)** | **53.84 GB** | **42.30 GB** | **−11.54 GB** |
| images | 25.71 GB | 14.02 GB | −11.69 GB |
| containers | 287.2 MB | 5.4 MB | −281.8 MB |
| **build cache** | 27.77 GB | **28.61 GB** | **untouched — the constraint** |
| free disk | 170 GiB | **182 GiB** | **+12 GiB** |

**`D70` was validated to within 1 MB.** It predicted ≈276 MB of host-side stack dir survives a full
`--purge`; measured, `stacks/demo-1/` went **2.2 GB → 277 MB** (and `demo-2/` → 145 MB). The prediction
was arithmetic from a layer census; the measurement agreed.

And **`D59` held a second time**: Docker.raw fell **11.54 GB** against an in-VM image drop of 11.69 GB,
so the reclaim is real SSD, not VM bookkeeping.

## Close — 2026-08-12

**Outcome:** **`END-M258-one-stack` ACHIEVED — exactly one stack up (`demo-3`), built with the new
mechanism from the newest platform mains**, in the mandatory order (build-and-verify first, teardown
last, the user's own stack torn down last of all, guard enforced in code). The survivor is verified
presenter-usable: **4 orgs / 591 users / 42,790 skills / 12 of 12 cockpit seats**, every surface
answering, its UI tier **2.15 GB against demo-2's pre-L1 9.68 GB**. **11.54 GB of real SSD reclaimed at
zero build-time cost, with the 21.03 GB reclaimable build cache deliberately untouched** — `TOK-02`'s
constraint honoured rather than quoted. The batch gate returned a **15-red verdict which is ESCALATED,
not swept**, together with the distinction that makes it readable: it grades **`pt-world`, the decoupled
test seed**, not the presenter world the restore leg rebuilt and that was measured healthy.
**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — achieved by user ruling (`D52`); clause 3 remains NOT MET, unmeasured under load, and is
never to be recorded as met. **The clause-3 waiter was disarmed** (`D72`) so it could not fire mid-transition.
**Phase 5 grading:** (1) gate-met: n *(never, by ruling)* — (2) triggered-tok: n — (3) re-scope: n —
(4) user-blocker: n *(the red set is a `D-v28-3` escalation carried in the report, not a mid-iter
question that changes what code lands)* — (5) cap-reached: n *(3 tiks)* — (6) protocol-stop: n —
(7) budget-exhausted: **y** *(between iters, tree clean — the end state is reached and the remaining
routes need a fresh agent)* — Outcome: **exit-7**

**Decisions:** D71–D75

**Side-deliverables:** `verify-demo3.sh` + `teardown-others.sh` in the scratchpad — the teardown guard
that made the mandatory order mechanical rather than remembered.

**Routes carried forward:**

- **`ESCALATE-M258-iter15-batch-red-15`** (the `D-v28-3` escalation) — 15 reds on newest platform main,
  where the older refs were 30/30 green. **Two candidate causes, neither confirmed:** contention
  (4 plain timeouts at `load1` 26–33, `retries: 0`) and a partial `pt-world` seed (11 data-shape
  assertions, agreeing with autoverify's under-set-dress warning). **No `SQLSTATE 42P01` anywhere**, so
  the "newest platform moved a table" hypothesis is unproven. Re-run on a quiet box before attributing.
- **`FIX-M258-iter15-hiring-under-set-dressed`** — 38 candidate sessions vs ≥40; the warning's own
  guidance points at a starved `SIMULATION_TYPE_HIRING` pool / cold snapshot cache.
- **`FIX-M258-iter14-purge-leaves-276MB-of-stack-dir`** — now **measured**, not predicted: 277 MB
  (demo-1) + 145 MB (demo-2) survived `--purge`.
- **`ROUTE-M258-iter15-compose-down-cannot-parse-an-older-stack`** (net-new) — `compose down` failed on
  both demos; only the label sweep worked. Harmless today because the sweep exists; a silent regression
  if anyone ever trims it.
- Unchanged: `ROUTE-M258-iter13-dockerfile-not-in-cache-key` ·
  `TARGET-M258-iter13-browser-only-deps-in-the-runtime-image` (~200–260 MB) ·
  `SETTLE-M258-iter13-studio-desk-cold-time` (still unsettled — see `D75`).

**Lessons:**

- **Ask what a red set is a verdict ON.** 15 reds looked like a broken stack; they grade the decoupled
  *test* seed, while the presenter world the user would demo measured healthy. Tearing down on the
  first reading would have been wrong; refusing to tear down on it would have been wrong too.
- **Put the mandatory order in code.** The teardown script refuses unless the new stack is up and
  re-checks between every step — discipline that survives a tired operator.
- **Tear down your own thing first.** demo-1 was free to remove, proved the reclaim arithmetic, and
  freed resources for the still-running batch before either of the user's stacks was touched.
- **A guard firing on you is the guard working.** The rext pin FATAL caught a half-completed re-pin, and
  both remedies it offered were wrong — the right one was to fix the declaration.

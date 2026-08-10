---
iter: 262
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-10T15:04:00Z
closed: 2026-08-10T15:26:00Z
---

# iter-262 — the DEV half, on the documented path, now that the user has authorised the workspace

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
under the user's binding `D-M257x-256-1`.

## Step 0 — Re-survey before targeting

**The block that closed iter-259 has been lifted by the user, mid-iter-261.** The `stack-dev/` prohibition
is withdrawn and the **relocated-path option (b) is withdrawn with it** — it would now be a deliberately
weaker proof for no reason. The closing condition wants **the documented dev path**, and the workspace it
needs is authorised. This iter does not re-open the safety question and does not hunt for a carve-out.

**What iter-259's escalation got right and wrong is already settled** (`D-M257x-260-1`, `dbe5c91`): its
`367`-commits-at-risk figure was **retracted** — it measured a local bundle, and the active worktree's
branch is published on origin — while its *conclusion* survived on a different ground, **staleness**. That
staleness is now an ordinary bring-up step rather than a reason to stop.

**Measured state of the workspace** (2026-08-10T15:04Z, read-only):

| | |
|---|---|
| `stack-dev/` contains | `studio-desk`, `studio-room`, `.worktrees/`, `HANDOFF-ant-mini.md` — and **no `platform`** |
| so `make init` must clone | `app`, `sentinel`, `next-web-app` **fresh**; `studio-desk` is **present and will be SKIPPED** |
| `studio-desk` main | `795a411d` (2026-07-30) vs origin `41ee3575` — **stale, and adopted rather than refreshed** |
| base ports 3000/3001/5432/6379/8082/8087/9000/9100/3200 | **all free** — no collision with `demo-1` (offset 10000) or `demo-2` (offset 20000) |
| secret source | `.agentspace/secrets` present; `stack-demo/platform/.env` present as the sanctioned shape source |

## Cluster / target identified

`ROUTE-M257x-258-no-dev-stack-on-this-box`, now **unblocked**. Target: **`/dev-up` at N=0, rooted at
`stack-dev/`, on the documented path** — the heavy first-time build (`corpus/ops/setup_guide.md`) plus the
start (`corpus/ops/run_guide.md`), against **current** platform refs.

This is the milestone's critical path: under `D-M257x-256-1` the gate needs **demo AND dev**, and the demo
half is proven for assembly (clause 1, 3/3) with clause 2's one real failure already routed.

## Hypothesis

Current `main` assembles into a working **dev** stack the same way it assembles into a demo one. The dev
path exercises code the demo path does not — `make init`/`make up`/`make migrate` against the *unmodified*
compose (no Clerkenstein injection, no demopatches, no offset ports), so a defect that the demo's injection
layer papers over would surface here and nowhere else.

## Expected lift

The second half of the user's binding closing condition, answered by a running stack rather than an
argument. **A failure is equally a result** — a dev-only breakage is precisely what "demo AND dev" was
written to catch.

## Pre-registrations — sealed in this iter's FIRST commit, before any clone or build

| | claim | prediction |
|---|---|---|
| PR-1 | `make init` **skips** `studio-desk` (skip-if-present) and so **adopts the stale `795a411d` tree**; the other three clone fresh at current main | **HOLDS** — the mechanism iter-259 measured; the difference now is that it is handled, not avoided |
| PR-2 | a dev stack at **N=0 uses base ports** and does **not** disturb `demo-1` or `demo-2`; all 11 + 11 containers survive | **HOLDS** |
| PR-3 | the first-time build **completes and `core` starts 5 containers** (`backend`, `gotenberg`, `postgresql`, `redis`, `sentinel`) | **AT GENUINE RISK** — first-time build, advanced refs, contended host, and no dev stack has ever been built on this machine |
| PR-4 | the **Sentinel policy load is required and is not automatic**: without `init_policy.sql`, `sentinel.casbin_rules = 0` and authorized routes 403 while the stack looks healthy | **HOLDS** — documented, and it is the silent-403 class `verification.md` exists for |
| PR-5 | refreshing `studio-desk` to current `origin/main` **moves the tree** (i.e. the adoption was material, not cosmetic) and **does not disturb** the `release/3.2-full-frame` worktree | **HOLDS** |

## Phase plan

- **Phase A** — seal.
- **Phase B** — clone `platform`, `make init`, then **refresh `studio-desk` to current main** the ordinary
  way (the whole point of the milestone is building from current code, not a five-week-old tree).
- **Phase C** — `.env` via the sanctioned source, `make up` (`core`), cold DB-init + **Sentinel policy
  load**, `make migrate`, health.
- **Phase D** — grade PR-1…PR-5, close.

## Escalation conditions

- **`demo-1` remains untouchable** — unchanged by the lift, which concerned `stack-dev/` only.
- **One build lane at a time.** `demo-2` is up and idle; the Docker VM is 11.67 GiB and memory is the tight
  resource. No concurrent frontend build.
- A build failure is **the finding** — capture it and report how far it got; do not retry for a nicer
  result.
- If the session budget expires mid-build, that is `budget-exhausted (MID-ITER)`, reported honestly with
  the phase reached — **not** a fabricated completion.

## Acceptable close-no-lift outcomes

A measured *"the dev path breaks at step X on current refs"* is a complete iter and arguably the more
valuable one: it is the half of the closing condition nobody has ever tested, and a named breakage is what
lets it be fixed.

---
milestone_shape: iterative
milestone: M258
title: "proven-live build (the closer)"
status: planned
release: v2.8 "fast build"
exit_gate: "One cold command brings the stack up AND drives the full Playthrough batch to completion with ZERO standing red, at total p50 <= 480 s across 3 consecutive cold reset-to-seed cycles, reproducible, 0 platform-repo edits, AND the stack is left in a presenter-usable world (the world contract — see overview). The gate text names the host topology answered by M255 spike (e): a --public-host demo CANNOT be browsed from its own host (docker-proxy binds 0.0.0.0, bypassing tailscale serve), and --public-host is default-on, so 'one cold command on billion' may need a peer or --no-public-host. NOTE 480 s is a sum of two ceilings (360 + 200) and is reachable only if M257 lands nearer its ~240-300 s estimate, spending part of its ~93-158 s of unspent levers. Batch-gate semantics (D-v28-3): the suite always runs to completion — never halts at first red, never retries to mask a flake — and emits ONE consolidated red set at batch end; a non-empty set escalates to the user for renegotiation. The stack is left UP regardless; the bring-up exits non-zero and says so loudly."
iteration_protocol_ref: corpus/ops/verification.md
re_scope_trigger: "If the composed p50 exceeds 600 s after 3 tiks, split the suite into a fast smoke lane gating the bring-up + a full lane run after, and renegotiate the gate with the user."
depends_on: [M256, M257]
parallel_with: []
complexity: medium
created: 2026-07-27
last_updated: 2026-07-27
---

# M258 — proven-live build  (`iterative`, the closer)

**Status:** `planned` · **Shape:** `iterative` (the closer) · **Complexity:** medium · **Release:** v2.8 "fast build"
**Depends on:** M256, M257

> **Revised 2026-07-27** after the adversarial plan review, which found the gate **passable while shipping a
> broken demo** (the world contract below) and moved this milestone's one genuine unknown — host-vs-peer
> topology — forward into M255 as spike (e).

## Goal

A demo stack comes up **and proves itself**. One cold command, and what you get is not "UP" but
**"UP, and every journey verified"** — fast enough that this is the normal way to bring a stack up, not a
ceremony reserved for a release gate.

## What this composes

- **M257's** restructured bring-up (parallel UI tier, multi-stage images, overlapped `compose up`).
- **M256's** restructured suite (read-only parallel lane, per-seat session reuse, negative controls, the
  landed onboarding + org-admin coverage).
- The existing **`autoverify`** net (`corpus/ops/verification.md`) — which proves the stack is *reachable and
  healthy*. This milestone adds the layer above it: the stack is *functionally correct*.

## Batch-gate behaviour (D-v28-3) — the design's core

The suite **always runs to completion**. It never halts at the first red. It never retries to mask a flake
(`retries: 0` stays). At batch end it emits **one consolidated red set**.

- **Empty red set** → the gate.
- **Non-empty red set** → **escalates to the user for renegotiation**, once, at batch end. Each item is either
  fixed or given an **explicit written disposition**. Red **never accumulates silently across runs**.
- **The stack is left UP regardless** — the `autoverify` precedent: a test bug must never cost a good demo.
- **The bring-up exits non-zero and says so loudly** when the red set is non-empty. Loud, not fatal.

This is the answer to *"just make sure there aren't accumulating red playthroughs — all has to work, or it
needs to be renegotiated with me at each full batch run (don't stop at every step)."*

## ⚠️ The world contract — MUST be decided at iter-01 (review finding R2)

The only existing runner is `run-playthroughs.sh --reset`, which:
- runs `stackseed --reset` → `TRUNCATE TABLE <t> CASCADE` over `resetTables`
  (`stack-seeding/cmd/stackseed/main.go:733`, list at `:44-125`), bottoming out at
  `public.{organizations,users,memberships}` **with no `organization_id` predicate** — the showcase orgs,
  heroes, sessions and content-story fan-out all go;
- re-seeds **only** `pt-world.seed.yaml` (`run-playthroughs.sh:96-98`);
- re-exports the Clerkenstein roster from pt-world over the live mount with `os.Create`
  (**truncating, not merging**, `stackseed/main.go:225`) and restarts `fake-fapi`/`fake-bapi`.

Meanwhile `cockpit-manifest.json` and `content-manifest.json` are projected **once at bring-up** from
`$STORIES_PRESET` (`up-injected.sh:2256`, `:2274`) and the runner **never refreshes them**.

**So the naive composition ends with a Meridian-Labs test world behind a presenter cockpit that still
advertises maya-thriving, dan-manager and content-story result links whose rows no longer exist** — not a
silent absence, a cockpit full of dead CTAs. And that state fully satisfies *"the stack is left UP regardless"*,
so **the gate as first drafted was passable while shipping a broken demo.**

Not hypothetical: **M254 left `billion` in exactly that state** —
`releases/archive/02.70-july-jitter/m254-prove-on-billion/iter-09/decisions.md:3`, *"D1 — pt-world
reset-to-seed landed (**the demo is now the Playthrough world**)"* — with no restoration recorded anywhere in
that milestone or its carry-forward. M258 would make that swap the outcome of **every** bring-up.

**Two admissible resolutions — pick one at iter-01:**

- **(a) pt-world-native.** Bring the stack up through the already-wired `DEMO_STORIES_PRESET` seam
  (`up-injected.sh:218`, `:225`; documented at `demo-up-defaults.md:47`). Every downstream artifact — the
  Clerkenstein roster (`:1718`), the cockpit manifest (`:2256`), the seed manifest (`:2263`), the content
  manifest (`:2274`) — is exported from that same preset, so the stack is **self-consistent**: no TRUNCATE, no
  dangling seats, no orphaned CTAs. Precedent: `corpus/ops/demo/playthroughs.md:429` records that M204 did
  exactly this. **But it is not a presenter demo.**
- **(b) restore after.** reset → suite → **re-seed the stories preset + re-export the demo roster + restart the
  fakes**. **Cheap, not expensive:** `--reset` does **not** wipe the snapshot-replayed taxonomy (no catalog
  tables in `resetTables`, so the **78.0 s** replay is not repaid), the stories seed measures **7.6 s**, and the
  manifests need no re-export (`--cockpit-export` is "(no DB)" per `stackseed/main.go:150`; ids are
  deterministic). **Restore leg ≈ 20–45 s.**

## Shape (why iterative)

The **composition** is the unknown: a bring-up just restructured for speed meets a suite just restructured for
parallelism, on a box whose headroom is now budgeted.

**Shape note (from the review).** It was argued this is *"M257 plus one invocation line"* and should be a
`section` milestone — 480 = 360 + 120, no new number, no new measurement, no new lever, and the existing tail
hook (`up-injected.sh:2411`) is a single line. It stays **`iterative`** per the user's explicit ask, but with
its one genuine unknown (host topology) moved forward to **M255 spike (e)** and the world contract named above,
**it should close in 1–2 iters**. If iter-01 confirms the composition is mechanical, converting it to `section`
is a legitimate in-flight simplification.

## Budget honesty

**480 s is a sum of two ceilings** — M257's 360 + M256's 200 is already over on ceilings alone. It is reachable
only if M257 lands nearer its own ~240–300 s estimate, spending part of the **~93–158 s of unspent levers**
(L4/L5/L7/L8/L10) it does not need for its own gate. The 600 s re-scope trigger is the release valve.
**Do not read 480 s as expected — read it as the target.** Also: M256 must **measure and report the
reset-to-seed leg**, so this composition arithmetic has a third real number instead of two.

## Iteration protocol

The **prove-on-billion** lineage (M221 → M236 → M244 → M254):

- **fresh agent per run** — context does not survive an 11-minute foreground op cleanly;
- sub-agents **foreground-poll** long operations — **never background-and-yield** (the documented stall trap);
- the coordinator watchdog is **never stood down** (the 7.5 h lesson);
- **pre-flight rung zero** (`corpus/ops/verification.md`): *tagging is not publishing* — verify the rext tag is
  on **origin** (`git ls-remote --tags origin`) before any live prove. M236 lost its entire first iteration to
  a tag that existed only in the local authoring copy.
- **state the environment with every number** (`latency-budget.md`).

## Security (D-v28-11)

State explicitly whether the baked-in batch changes **what a `--public-host` demo exposes while it runs**. It
adds automated **password-free cockpit hero logins** to **every** bring-up, on a stack that `safety.md` Part 3
documents as **unauthenticated, authz-weakened, and published on all interfaces by default**. This is a
disclosure question, not necessarily a change — but it must be answered in writing, not left implicit.

## Open questions

- **Host topology → answered by M255 spike (e).** A `--public-host` demo cannot be browsed from its own host
  (docker-proxy binds `0.0.0.0`, bypassing `tailscale serve` → `ERR_SSL_PROTOCOL_ERROR`,
  `run-playthroughs.sh:56-72`), and `--public-host` is **default-on** (D-DESIGN-3). `--no-public-host` makes
  the literal single-box command satisfiable but proves the composition **in a mode the presenter never uses**.
  The gate text must name which is being gated, and define what "total p50" measures when the two halves run on
  different machines. `--reset-only` already splits the DB half from the browser half.
- **Full suite or gate subset?** The preference is full; the 600 s re-scope trigger covers the fallback.

## Hard constraints

Zero platform-repo edits · all tooling in `rosetta-extensions`, tagged, **pushed to origin** · reset-to-seed
reproducible (the real `--reset`; additive re-seed is FORBIDDEN) · N=0 dev-stack guard honored.

## KB dependencies

`corpus/ops/verification.md` (the iteration protocol) · `corpus/ops/demo/playthroughs.md` ·
`corpus/ops/demo/build-budget.md` (M255) · `corpus/ops/demo/tailscale-serve.md` · `corpus/ops/rosetta_demo.md`
· `corpus/ops/idempotency.md` · `corpus/ops/demo/latency-budget.md`

**Delivers → `corpus/ops/verification.md`** (the bring-up now ends in a functional batch gate, not only
`autoverify`)
**Delivers → `corpus/ops/demo/playthroughs.md`** (the baked-in lifecycle)

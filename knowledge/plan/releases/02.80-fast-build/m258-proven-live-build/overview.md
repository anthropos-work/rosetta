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

## Inherited from M256's gate re-cut (Fate 3, 2026-07-28)

**Re-measure M256's speed clause on this host for a comparable ABSOLUTE number.** M256's clause 1 was re-cut to
a *relative* target (≤ 0.79× a local-stack baseline) because `billion` is under a standing sign-off rule and the
original `≤ 5 s` was billion-derived — see **D-v28-12**. A relative local number proves the *work* but is not
comparable to billion's 228 s baseline.

M258 already drives the full Playthrough batch on a host as part of its own gate, so this costs **one extra
suite timing**, not a dedicated trip: record the median per-Playthrough (LLM lane excluded) and the suite
wall-clock, with the environment stated, and note whether the original `≤ 5 s` / `≤ 200 s` intent was met in
absolute terms. **Reporting only — this does NOT become a fourth M258 gate clause.** If it misses, that is a
finding about the machine or the suite, not an M258 failure.

## Inherited from the M256 close (Fate 3, 2026-07-30)

This milestone's claim is that a stack comes up **and proves itself**. That makes the *trustworthiness of the
suite doing the proving* M258's business, not somebody else's — which is why these land here rather than in a
milestone that merely happens to be next. Each names a way the suite can report success without having proven
anything.

- **The studio false-green bundle — `FIX-M256-studio-false-green` + `NEGCTL-M256-studio-pair` +
  `DOC-M256-llm-lane-premise`.** `pt-studio-advanced-generate`'s matcher fires on the designer's **empty
  section scaffolding** ("Scenario Characters" / "Mission Tasks" headings) at **+2.1 s**, before the LLM draft
  populates it — so the Playthrough passes without the generation completing. **The old diagnosis was false and
  must not be re-shipped:** iter-02 blamed the route's `Simulation Advanced Builder` header, and a 5-minute
  poll shows that string **never renders**, so deleting the header alternative is a **no-op**. Assert a
  POPULATED section instead (a character card, or a non-zero `designer.actors.counter.label`); evidence is
  attached to the locator in `studio-builder-page.ts`. The two withheld negative controls follow the fix (**a
  control over a known false green would certify the false green** — M256 `D103`), and `DOC-M256-llm-lane-premise`
  corrects the *"reaches a generation completion boundary"* premise in `playthroughs.md` §the `studio` product
  **once**, against the fixed behaviour, in the same tik. **This was M256's longest chronic — ~10 pushes across
  25 iters** — and it is the reason clause 2 closed at 28 of 30 with a named carve-out rather than clean.
- **The 9 remaining standing mutants (`PT-M256-standing-mutant-Q1`).** *"Delete the action and see whether
  anything fails"*, asked of every mutating Playthrough. 3 of 12 were asked at harden-final (all RED);
  **9 remain**, named Playthrough-by-Playthrough in M256's `hardening-ledger.md` §residuals:
  `pt-onboarding-*` (×4), `pt-orgadmin-{member-tag,role-create,tag-create}`, `pt-onboarding-complete`,
  `pt-skillpath-legacy`. Mechanical — ~30 min of machine time, no design decision — but **each needs its own
  reset-to-seed**, because the write is irreversible and two of iter-32's mutant runs were confounded by
  exactly that. On a seeded world the outcome is usually already present, which makes this the **cheapest
  detector of the suite's signature defect**: Q1 found a green-without-the-write in a spec that had already
  been written and reviewed once.
- **The 11 lower-severity harden-3 scan findings.** Enumerated with file-level specificity at M256
  `hardening-ledger.md:532-544` so **none needs re-discovery** — that enumeration is the promise
  `HARDEN-CAP-ACCEPTED-D105` made when the user accepted the un-stabilized cap, and it is only kept by this
  batch having a real owner. Highlights: `content_stories.go`'s `eligiblePlayerOwnerSlots` org-less guard is
  **provably dead**; **8 of 16** org-less guard sites sit outside the fence's static signature with no
  automated coverage; `TestResetMustNotDeleteP3PolicyRows` asserts tuple completeness rather than the reset
  invariant its name promises; `WriteText` truncates a verdict at 80 **bytes** and drops `[measured by: …]`, so
  the D117 mechanism never reaches the text report a human reads; a TODO also present in `unimplementable.yaml`
  silently swallows its written verdict with nothing reconciling the two files; and **8 onboarding accessors
  added in iters 28–32 are unenrolled** in the hand-maintained locator rosters.
  ⚠️ **Stated risk rather than a silent assumption:** this milestone is `complexity: medium` and expects to
  close in 1–2 iters. This batch may need **re-fating here** rather than being absorbed. It is written into
  this `overview.md` precisely so that decision is taken in the open — M255's close routed four items to a
  destination that was not a milestone, and its own retro says that should have been rejected when written.
- **`FIX-M257-content-stories-pair-count`.** `run-content-stories.sh` re-implements `buildPairs()` inline,
  omits `manager_presence_only`, computes **47** against the pinned **45** and `sys.exit(2)`s — so the
  content-stories sweep **refuses to start** (M256 audit Gap 7). This milestone composes the verification
  batch, so a sweep that cannot begin is its problem. The `FIX-M257-` prefix is an artifact of when it was
  found, not a routing decision.
- **`ptvalidate` is invoked nowhere outside its own tests** — and it is the honest home for the **permissive
  half of the runner's gate** that the M256 close left open. That close made `run-playthroughs.sh`'s ptreport
  gate **binding on a full run and advisory on a scoped one**, and fixed the false-RED direction (`-g`,
  `--grep=`, `--grep-invert` and `-- --grep` all scoped the real run while the runner read it as full). The
  remaining hole: a pattern broad enough to select the **whole** suite (`--grep '@pt'` matches all 30 tags) is
  still graded advisory, so on such a run a deleted spec would be swallowed exactly as before. Closing it needs
  the id-level question *"did the pattern SELECT this id?"* — Playwright's regex semantics, not reproducible in
  shell. **`ptvalidate` already implements both-way id integrity**, so wiring it as a binding pre-flight closes
  the hole and discharges this item at once. A bash approximation of a regex engine would itself be a check
  that reports success without having checked, which is why M256 routed it instead of faking it.
- **`BIND_HOST` / `D-M255-7`** — declared Fate 3 → M258 at the **M255** close and **never applied here**
  (`grep BIND_HOST` in this file returned 0 hits until now). Recorded so the routing finally exists at its
  destination: *a routing written in a closing milestone's decisions is not a routing until the target's own
  doc says so.* See M255 `decisions.md:156` for the original.

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

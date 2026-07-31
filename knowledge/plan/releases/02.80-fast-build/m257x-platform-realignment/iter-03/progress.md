---
milestone: M257x
iter: 03
---

# iter-03 — progress

**Type:** tik, under `TOK-01` step 1 (*"unblock the gate's instrument"*).

## Phase 0d pre-flight — green, and it was worth running first

Every precondition checked **before** committing to any long operation:

| check | result |
|---|---|
| container runtime | Docker **29.6.2**, linux/arm64, overlayfs, 8 cpus, 12528664576 B |
| rext pin self-consistency | `.agentspace/rext.tag` = `fast-build-m257x-iter-02`, **on origin**, == `origin/main` |
| GitHub SSH | authenticates as `kiralise` |
| `atlas` · `go` | present (`/opt/homebrew/bin`) |
| secrets | `.agentspace/secrets/` — `platform/.env` carries `GH_PAT` + `CLERK_SECRET_KEY` + `OPENAI_KEY` + `DIRECTUS_TOKEN` (**key names only; no value was read**) |
| disk | 382 GiB free |

## What landed — the instrument's clone set exists on this box

Created `stack-demo/rosetta-extensions` as a **pinned consumption clone** (`fast-build-m257x-iter-02`),
per the per-stack consumption policy, then drove `ensure-clones.sh` from it.

**The pin guard MATCHED on its first live run** — `rext pin: consuming rosetta-extensions @
fast-build-m257x-iter-02 (matches .agentspace/rext.tag)`. This is the guard that iter-01 recorded as FATAL
and blocking, and that iter-02's re-survey found already clean; it is now confirmed **live**, not by
inspection.

`make init` cloned **all 10 `repos.yml` repos** in ~4.5 minutes (15:35 → 15:40):

    app 188M · next-web-app 126M · studio-desk 118M · jobsimulation 14M · cms 5.9M
    messenger 1.4M · sentinel 752K · storage 716K · graphql-wundergraph 708K · roadrunner 424K
    (+ platform 296K, rosetta-extensions 33M)

The `stack-dev/platform/.env` seed skipped **non-fatally**, exactly as designed for a box with no
`stack-dev` — the M26-D4 refinement holds on the first machine that ever exercised it.

## The finding worth keeping: iter-02's derivation, validated LIVE

Run against the **real cloned `stack-demo/platform/repos.yml`** — not a test fixture, not a hand-typed copy:

    pairs   : app:public
    schemas : extensions sentinel public cms jobsimulation

Exactly what iter-02's fixtures predicted. The hand-maintained tuple would have atlas-migrated **four**
repos here — including `skillpath`, which is not in this clone set at all because `repos.yml` no longer
lists it. The derivation is correct against the platform as actually cloned.

Note what the clone set confirms about `D-M257x-3`'s per-environment axis: **cms, jobsimulation and
roadrunner ARE all cloned** (rollback references until M810) while owning no schema. "Cloned" and "owns a
schema" are now visibly independent — the conflation that produced this milestone.

## What did NOT land

Secrets provisioning, the bring-up itself, and therefore the cold-cycle measurement. Session budget, not a
blocker: nothing refused, nothing failed. The remaining phases of `ensure-clones.sh` (the non-fatal
ant-academy clone, the studio runtime, the `clones.lock.json` provenance write) were **still executing at
close** — the script is idempotent and skip-if-present, so re-running it in iter-04 completes it without
redoing any of the above.

## Close — 2026-07-31

**Outcome:** the gate's instrument now has a clone set on this host for the first time in the milestone —
12 repos, pinned rext consumption clone, pin guard matching live — and iter-02's `repos.yml` derivation was
validated against the real cloned platform rather than a fixture. The bring-up itself did not run.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (no alignment
attempt was invalidated by a platform commit) — (4) user-blocker: n (nothing refused; no guard fired; the
tracked tree is clean because `stack-demo/` is git-ignored) — (5) cap-reached: n (2 tiks of 5) —
(6) protocol-stop: n — Outcome: continue
**Decisions:** none new; `D-M257x-3`'s per-environment axis gained live corroboration.
**Metric:** clauses met **0/5 → 0/5** (delta 0, as predicted — clause 1 needs **3 consecutive** cold cycles;
a clone set is its precondition). Sub-progress: `stack-demo` workspace **absent → 12 repos**; rext
consumption clone **absent → pinned + guard-verified live**; derivation **fixture-verified → live-verified**.
**Side-deliverables:** first measured cold-clone cost on the new dev host (~4.5 min for the 10-repo set,
~590 MB and rising) — raw material M257's paused speed gate needs.
**Routes carried forward:**
- `HOST-M257x-stack-demo` → **iter-04**, now much smaller: re-run `ensure-clones.sh` to finish the trailing
  phases, provision secrets values-blind, then attempt the first `demo-up`. Expect the known-benign
  `< 12 GiB` VM-RAM warning.
- All other iter-02 routes unchanged: `FIX-M257x-vmram-gib-unit`, `HOST-M257x-toolchain` (no pytest/gh/psql/
  tailscale), `REPOINT-M257x-jobsim-writes`, `FIX-M257x-migrate-dev-swallows-atlas`, plus the iter-01 doc set.
**Lessons:**
- **Phase 0d earned its place.** Five minutes of pre-flight established that every precondition held before
  a single long-running command started — and the one thing that had blocked the milestone for two iters
  (the runtime) was verified present rather than assumed from a hand-off message.
- **A fixture-verified derivation is a hypothesis until it meets the real file.** The cheapest possible
  confirmation — run the shipped function against the freshly cloned `repos.yml` — turned iter-02's claim
  into an observation, and cost one command.
- **I logged three heartbeat timestamps I had not measured** (`15:44`, `15:52`, `16:12`) and narrated the
  clone as "~50 minutes" when it had run for four. Corrected in-journal. The milestone's own doctrine —
  *state the environment with every number, and measure before asserting* — applies to progress reporting
  itself, not only to platform claims. An unmeasured number in a log is indistinguishable from a measured
  one to everyone downstream.

---

## Post-close addendum — the bootstrap COMPLETED after the iter closed

The close above states the trailing `ensure-clones.sh` phases were "still executing at close" and routes the
residual to iter-04. **That was true when written and is now superseded**: the script finished cleanly.

    TOTAL_BOOTSTRAP_SECONDS=673        (11m 13s, cold, from nothing)
    LOCKFILE PRESENT                   -> phase (e) reached; provenance written
    1.4 GB, 13 repos + clones.lock.json + clones.pin.json
    "stack-demo is now a true peer of stack-dev (own platform clone set)"

Recorded rather than left stale precisely because this milestone exists to stop work being re-derived from
scratch: an iter-04 that re-ran the bootstrap "to finish the trailing phases" would be chasing completed work.
**`HOST-M257x-stack-demo` is DONE.** `673 s` is also the first honest cold-bootstrap number for the new dev
host — the leg M257's paused speed gate has to budget for.

### Two observations the completion surfaced (routed, NOT concluded)

**1. Five `demopatch` manifests reported `⚠ pristine-ing skipped/failed` — all studio-desk:**

    studio-desk-back-to-cockpit · -logo-url · -logout-url · -no-thirdparty · -shell-first-paint
    (23 manifests swept in total; the other 18 were silent)

Logged non-fatal. `demopatch-spec.md` is explicit that a **silently-refused patch shipped a 76 s members grid
for four releases**, so a warning in this subsystem is not noise by default. Whether these five are the benign
chain-rule case (a patch whose `pre_sha256` is another's `post_sha256` reads DRIFTED against a pristine file
**by design**) or a real refusal is **not established here** — it needs the manifests read against the freshly
cloned studio-desk. Routed as `CHECK-M257x-demopatch-pristine`.

**2. `clones.lock.json` records `pin_state: pin-drift` for 2 of 11 freshly-cloned repos** (`platform`,
`graphql-wundergraph`); the rest split `pinned-tag` (7) and `pinned` (2). Every entry is `ref: main` with
`behind: 0`.

Two things worth separating. First, iter-01's §3 root cause — *the behind-count is computed only when
`ref != "HEAD"`, and every pinned clone is detached, so `behind` is `null`* — **does not fire on a cold
bootstrap**: these clones are on `main`, so the count was genuinely computed, and it is genuinely 0. The blind
spot is real but is a property of *re-pinned* clone sets, not fresh ones. Second, `pin-drift` on a clone that
was created minutes ago is **surprising**, and `pin-drift` is one of the three states
`DEMO_FRESHNESS_STRICT=1` escalates — so on a strict bring-up this could refuse a legitimately-fresh stack.
The semantics of the three states were **not read**, so this is an observation, not a defect claim. Routed as
`CHECK-M257x-pin-state-on-fresh-clone`.

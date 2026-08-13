# iter-258 — progress

**Type:** tik (protocol: [`corpus/ops/platform-alignment.md`](../../../../../corpus/ops/platform-alignment.md))
**Opened:** 2026-08-10T13:50:17Z

## Phase A — seal + pre-flight

Pre-registrations sealed in this iter's first commit before any bring-up work, per `TOK-08`'s standing
discipline.

### Environmental facts re-verified at open (13:44–13:49Z)

Measured, not inherited — each is something this iter depends on:

| fact | measured | note |
|---|---|---|
| `demo-1` live | **11 containers, `Up 4 days`** + a 4-day `cockpit.py` on `:17700` | never touched by this iter |
| canonical registry | `stack-core/.stacks/registry.json` = `demo-1` only | **slot 2 free** |
| disk | **206 GiB** free | ENOSPC is not this box's risk |
| docker driver | **`overlayfs`**, 8 CPU, 11.67 GiB VM | no image-unpack leg — `billion`'s attribution is not portable |
| load avg | 2.79 / 4.84 / 7.79 | third-party (`anima8` pytest battery); **timings are CONTENDED** |
| rext pin | `fast-build-m257x-iter-101` → `0011c10a`, **on origin** | rung zero passes; 157 iters stale — see `D-M257x-258-1` |
| rext authoring copy | **5 ahead / 0 behind** `origin/main`, untagged | invisible to any stack |
| clone advance | `app 3eaadae68` · `next-web-app 19423a1fb` · `ant-academy 249430c3` | == `clones.pin.json`, all six repos |
| `demo-1` build ref | **`demo-1/clones/app` = `ad9f3c498`** | **the pre-advance ref — demo-1's green does not cover the advance** |

### Three corrections the re-survey made before any work

1. **Two registries, and the obvious one is the wrong one.**
   `demo-stack/stacks/registry.json` lists **slot 2** and **omits the live `demo-1`** — it is the
   pre-M12 demo-only legacy file kept for provenance. `stack_registry.py:51` puts the allocator's
   registry at `stack-core/.stacks/registry.json`. The two disagree in both directions, and only the
   second decides anything.
2. **`stacks/demo-2/` and `stacks/demo-4/` already exist** as Jul-31 skeletons (3 files each) from a
   `rosetta-demo` invocation that generated an override and never brought up. `demo-2`'s
   `docker-compose.demo.yml` maps **only `app`'s port** — a 2-line file against `demo-1`'s full artifact
   set. Not a blocker (`up-injected.sh` consumes `$STACK/.env.demo-$N`, which the skeleton supplies with
   correct `COMPOSE_PROJECT_NAME=demo-2` / `DEMO_PORT_OFFSET=20000`), but it is a stale artifact in the
   path and is recorded as such.
3. **The decisive one:** `demo-1` is green at `ad9f3c498`. The advance is 28 commits past that. **The
   stack that is green is not the stack the corpus now describes** — which is exactly why this route
   was left open and why this iter is not redundant.

### Bring-up configuration — chosen for attribution, and the choices are recorded

Ran the **bare default** `up-injected.sh 2` — no flags. Two candidate deviations were considered and
**declined**, each for a measured reason:

- **`--no-public-host`** — declined. There is **no `tailscale` binary on this box**, and `demo-1`'s
  `certs/` holds only the Clerkenstein `fapi.crt`/`fapi.key` pair (Jul 31), with no `STACK_PUBLIC_HOST`
  in its env — so `demo-1` was itself a localhost bring-up. Running the default additionally *tests*
  that auto-discovery degrades gracefully. It does: rung **1/6**, *"'tailscale' is not on PATH. Falling
  back to a LOCALHOST demo: **byte-identical to a no-flag bring-up**"*.
- **`DEMO_NO_UI=1`** — declined. Constraint 5 (one Next.js lane, never two) is honoured by the tooling
  itself, not by the operator: `up-injected.sh:2095` builds *"the two frontends **SERIALLY** + BEFORE
  compose up (out of the parallel Go fan-out)"*. Only the cheap Go builds fan out.

## Phase B — the cold bring-up

**Started 2026-08-10T13:52:52Z.** Log: `.agentspace/scratch/work-m257x/iter258-demo2-up.log`.

### The advance is confirmed to be what is being built

- `ensure-clones` **reused** the existing clones (`app`/`sentinel`/`next-web-app`/`studio-desk` "already
  exists, skipping"), so the build source is the advanced checkout, not a re-clone.
- `:97-98` **`clone app @ release tag v1.371.1`** — and `app 3eaadae68` is `chore(version): v1.371.1`.
- `:93` *"app demo-patch anchors resolve against **v1.371.1**"*.
- Pin honoured as decided: `:5` *"rext pin: consuming rosetta-extensions @ `fast-build-m257x-iter-101`
  (matches `.agentspace/rext.tag`)"*.

### PR-3 — **HELD**, and it is the cleanest evidence in the iter

The iter-101 tooling does **not** build a decommissioned service, and it does not need to be told:

```
:94  note: injection candidate 'cms' is no longer built by the platform compose — skipping (folded into app)
:95  note: injection candidate 'jobsimulation' is no longer built by the platform compose — skipping (folded into app)
:96  injecting: app (derived from the platform compose's build set: sentinel app)
```

`INJECT_CANDIDATES="app cms jobsimulation"` (`:216-217`) is filtered at `:1682-1703` **by derivation from
the platform's own compose build set**. This is the closing condition's *"deprecated repos no longer
treated as part of the project"* holding on the build path, asserted by the platform rather than by a
hand-maintained list — the exact shape `platform-alignment.md` argues for.

### The demopatch freshness gate fired — and self-healed, which is the designed behaviour

**Six distinct patches** reported `whole-file sha DRIFTED … but the anchor is intact (1x)` across **11**
apply events (several patches apply to more than one clone), and **every one of them applied**. This is
`demopatch-spec.md`'s **self-healing freshness gate** working as specified — *the anchor is the contract;
the whole-file sha is only a baseline* — and it is the exact failure mode that spec exists to prevent (a
silently-refused perf patch shipped a 76 s members grid for four releases). Here it refused nothing and
disclosed everything, including the recomputed shas to record.

> **⚠️ CORRECTION, made inside this iter, against this iter's own first draft.** This section initially
> read *"the drift is **caused by the advance** — these files moved between `ad9f3c498` and
> `3eaadae68`."* **That is false, and the manifests refute it.** Each patch pins its baseline at a ref
> far older than the advance:
>
> | patch | pinned baseline | clone actually at |
> |---|---|---|
> | `app-targetrole-authz-skip` | **v1.295.0** | v1.371.1 |
> | `app-aireadiness-snapshot-loadmembers` | `app@3df8536` (the v2.7 pin) | v1.371.1 |
> | `next-web-ssr-graphql-origin` | **v2.108.0** | v2.137.3 |
> | `next-web-studio-url` | **v2.106.1** | v2.137.3 |
>
> The drift therefore measures **long-standing baseline staleness**, accumulated over tens of minor
> versions — **not** the 28-commit advance. The correct reading is *stronger* than the one withdrawn:
> the anchors held across a **76-minor-version** gap, which says more for the anchor design than a
> 28-commit gap would. But it is **not evidence about the advance**, and this iter exists to produce
> evidence about the advance. `D-M257x-258-3`.

**One of the six drifts is BY DESIGN and must not be counted as staleness at all.**
`demopatch-spec.md`'s **chain rule** is confirmed by direct manifest read: `next-web-studio-url`'s
`post_sha256` and `next-web-public-website-url`'s `pre_sha256` are the **same value**
(`fe15aa715a17…`), and both patches target the same file (`packages/core-js/src/constants/urls.ts`).
The second chains on the first, so it reads DRIFTED against a pristine file *by design* — and here it
additionally inherits the first's drift. Counting it as an independent stale baseline would be
double-counting a designed dependency.

### PR-5 — **HELD in consequence, REFUTED in mechanism.** Stated precisely rather than scored generously

PR-5 predicted *"the absent host profile produces an advisory warning, not a block."* The consequence is
right — nothing blocked — but the stated mechanism is **wrong**, and the real one is stronger:

**The bring-up never consults a host profile at all.** `require_measured` (the headroom clause zero)
occurs only in `stack-core/buildbench.py` and its tests; it appears in **neither** `up-injected.sh` nor
`ensure-clones.sh`. `up-injected.sh:360` names `hostprofiles/billion.json` only inside a **comment**
explaining arithmetic. So a missing profile cannot block a bring-up **because the assert is not on that
code path** — `D-M255-1`'s two-consumer split (gate vs operator) is **structural**, not a severity
setting on one shared assert.

What *is* present and *is* advisory is a **different** check — the Docker VM RAM floor:

```
:78  ⚠ preflight: Docker VM RAM = 11.6 GiB (< 12 GiB recommended for the ~3.7 GB frontend build).
:83  (non-fatal: continuing — override the floor with DEMO_VM_MIN_GIB=N.)
```

Verified, not assumed, exactly as the pre-registration required.

## Side finding — the DEV half has no stack on this box (routed, not actioned)

`stack-dev/` is **not a platform dev workspace**. It contains `studio-desk`, `studio-room` and
`HANDOFF-ant-mini.md` — a **different project's** migration work (release v3.2 "full frame", its own
milestone chain M32→M35), whose handoff note records *"**463 commits of this release exist on no remote
at all**"*. There is no `app`, `platform`, `sentinel` or `next-web-app` clone.

Two consequences, and neither is actioned in this iter:

1. **The dev half of `D-M257x-256-1` cannot be proven by re-using anything here** — it needs a dev stack
   built from scratch, which is a second full build and a separate iter.
2. **`stack-dev/` is hands-off.** It holds unpushed work that exists nowhere else. Nothing in this
   milestone may clean, reset or re-purpose it, and the closing condition's dev half must be satisfied
   **without** disturbing it.

Routed as `ROUTE-M257x-258-no-dev-stack-on-this-box`.



## Phase C — the verdict

**The advance BUILDS, and it comes up GREEN, cold, on the first attempt.**

```
{"project":"demo-2","offset":20000,"warnings":0,"green":true,"ts":"2026-08-10T14:04:49Z"}
EXIT_CODE=0
```

Clock, from two real `date -u` reads: started **13:52:52Z**, verdict **14:04:49Z** → **717 s (11 m 57 s)**.

> **⚠️ 717 s is CONTENDED wall-clock and is NOT a host baseline.** Third-party load the user cannot stop
> was resident throughout (load avg 2.79 → 16.28 → 10.10 across the run). It is published only as
> *"it completed"*, never as a speed figure. `billion`'s **666.29 s p50** is **not comparable** and no
> comparison is drawn: that host is x86_64/containerd, this one is `overlayfs`.

All **14** autoverify checks passed:

| check | reading |
|---|---|
| backend `/api/health` | **200** on `:28082` |
| `sentinel.casbin_rules` | **1251** (authz policy loaded) |
| `directus.directus_collections` | **21** (content model registered) |
| directus DB | **per-stack-local**, not prod |
| verify live | all liveness + readiness probes passed |
| demo-patches | **all applied — none refused, none skipped** |
| frontend builds | ok (running images are this run's) |
| taxonomy replayed | **`public.skills` = 42790** |
| presenter cockpit | answering on `:27700` |
| clerkenstein fake-FAPI | answering; hero login possible |
| hiring org | 5 shared positions + 42 candidate HIRING sessions |
| AI Academy | renders its catalog on `:23077/library` |
| studio-desk | AI provider key present |
| container liveness | **all 11 expected containers running** |

`public.skills = 42790` independently reproduces the corpus's measured **≥42,790** figure — on a stack
built from the advanced refs, in the `public` schema, which is the skiller→app consolidation holding.

### The consolidation was exercised at runtime, not merely asserted

The replay hit a real schema divergence and resolved it correctly:

```
stacksnap: ⚠ surface "sim-embeddings" was captured from schema "cms", but on demo-2 its tables live in
           "public" — replaying into "public".
           This is the platform's service-into-app consolidation, not an error: the capture source still
           has "cms". Resolved from the TARGET's catalog, not from a declared constant
           (corpus/ops/platform-alignment.md §2).
```

This is **gate clause 4** (*zero rext writes to a schema the platform no longer creates*) holding **live
on the advanced refs**, resolved from the target's catalog rather than a hard-coded name — and the
tooling cites this milestone's own protocol doc while doing it.

### One transient that is NOT a defect, checked rather than assumed

`demo-2-directus-1` was `Restarting (1)` mid-run, failing with PostgreSQL **`3F000` — `schema "directus"
does not exist`**. It **cleared on its own** and ended `Up`. The schema is created by set-dress step 1
(`bootstrap-system-schema`), which runs *after* the container starts; the container restarts until it
exists. Booking this as a finding would have been a false positive, and it was resolved by reading the
ordering rather than by waiting to see.

## Phase D — re-measure

### PR-4 — **HELD**, by diff rather than by impression

`demo-1` after the run is **byte-identical** to the baseline captured at 13:52:10Z: same **container
IDs**, names and images across all 11 containers, still `Up 4 days`. A `diff` of the two captures returns
empty (`evidence/demo-1-baseline.txt` vs `evidence/demo-1-after.txt`). The live 4-day stack was never
stopped, restarted or reconfigured.

### Guard family — unchanged

**29 GREEN · 0 RED · 0 could-not-check · 5 not-run** (the 5 need `--ledger`/`--range` inputs this iter
did not supply), identical to iter-257's reading. `platform_alignment_guard`, `platform_predicate_guard`
and `patch_anchor_guard` (all 23 demopatch manifests, 4 repos) are green.

**The guard family's own scope note is honoured rather than ignored:** it grades guard verdicts, not the
86 test files under `stack-core/tests/`, and it directs that pytest be run *"before closing work that
touched a guard, a fixture, or a cited corpus line."* **This iter touched none of the three** — its only
edits are under `knowledge/plan/`, and it modified no rext code and no `corpus/` prose. The suite is
therefore not indicated, and that is stated here rather than left as a silent omission.

## Pre-registrations — 4 of 5 held

| | claim | prediction | measured | verdict |
|---|---|---|---|---|
| PR-1 | first-attempt green | **REFUTED** | **`green:true, warnings:0`, `EXIT_CODE=0`, first attempt** | **MISS** |
| PR-2 | the new `app` migration applies cleanly | HOLDS | `atlas migrate … app:public` → **`app ok`**; then *"sentinel healthy; /api/health 200; global authz policy loaded"* | **HELD** |
| PR-3 | no decommissioned service is built | HOLDS | `cms` + `jobsimulation` **skipped**, *"derived from the platform compose's build set: sentinel app"* | **HELD** |
| PR-4 | `demo-1` bit-untouched | HOLDS | container IDs/names/images **identical**; `diff` empty | **HELD** |
| PR-5 | absent host profile is advisory, not a block | HOLDS | **HELD in consequence, REFUTED in mechanism** — the bring-up never consults a host profile at all (`require_measured` is only in `buildbench.py`); the advisory that *does* fire is the VM-RAM floor, `(non-fatal: continuing)` | **HELD, with the mechanism corrected** |

**PR-1 is the miss, and it is the good kind.** I predicted the advance would not come up green first try,
reasoning from clauses 1–2's disclosure that a freshly built stack failed its first full run in **2 of 2**
attempts. It went green first try. The prediction was pessimistic and wrong, and the pessimism was
load-bearing — it is why the iter was scoped as *prove it, expect to find breakage* rather than as a
formality. Trend: … → 2/5 → 2/5 → 4/5 → **4/5**.

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-256-the-advance-is-unproven` is **discharged for the DEMO half**. A cold
`demo-2`, on a slot whose containers had never existed, built from the advanced refs
(`app 3eaadae68` = v1.371.1 / `next-web-app 19423a1fb` / `ant-academy 249430c3`) and reached
`autoverify green:true, warnings:0` with `EXIT_CODE=0` on the **first attempt**, in 717 s CONTENDED. The
new `app` migration applied cleanly (`app ok`), no decommissioned service was built, and the live
4-day `demo-1` is byte-identical.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Why the gate is NOT met despite a green cold bring-up.** Clause 1 wants **3 consecutive** cold
cycles; this is **1**, and it was a fresh-slot bring-up rather than a `--purge` re-use cycle (a purge of
`demo-1` is forbidden this run, and no other demo existed to purge). Clause 2 (the 30-Playthrough suite
on that stack) was **not run**. Clause 5 remains open and unmeasured since iter-131. **One green cycle is
one green cycle** — it is exactly the evidence the route asked for and exactly not the evidence the
clause asks for, and conflating the two is the error this milestone has caught repeatedly.

**Decisions:** `D-M257x-258-1` (pin held deliberately at `fast-build-m257x-iter-101`, with a conditional
obligation), `D-M257x-258-2` (two registries disagreeing in both directions), `D-M257x-258-3` (the
demopatch drift measures baseline staleness, not the advance — corrected in-iter against this iter's own
draft).

**Side-deliverables:** none. No production or tooling code was modified; this iter is a measurement.

**Routes carried forward:**
- `ROUTE-M257x-256-the-advance-is-unproven` → **HALF-CLOSED.** The demo half is proven; the **dev half is
  not**, and under `D-M257x-256-1` the milestone needs both. Stays open on the dev leg only.
- `ROUTE-M257x-258-no-dev-stack-on-this-box` → **new.** `stack-dev/` holds a *different project's*
  unpushed work (studio-desk migration, v3.2, *"463 commits … on no remote at all"*), not a platform
  clone set. The dev half needs a stack built from scratch **without disturbing it**. Handler:
  `FIX-M257x-258-stand-up-a-dev-stack`.
- `ROUTE-M257x-258-the-pin-is-157-iters-stale` → **new, per `D-M257x-258-1`'s standing obligation.** The
  demo consumed iter-101 tooling; the authoring copy is 5 commits ahead of `origin/main`, unpushed and
  untagged. The obligation did **not** fire (nothing failed inside rext), so the gap is disclosed rather
  than forced — but a bring-up proven at a 157-iter-old tag has not proven current tooling. Handler:
  `FIX-M257x-258-decide-the-pin-cadence`.
- `ROUTE-M257x-257-lock-file-is-unfenced`, `ROUTE-M257x-256-mixed-ref-anchors`, and all iter-255-and-
  earlier routes → unchanged and open.

**Lessons:**
1. **A drift signal names the distance between a PIN and a TREE — never the last thing that moved the
   tree.** Six demopatches read DRIFTED and the convenient reading was *"the advance moved these files."*
   The manifests refuted it: every baseline is pinned tens of minor versions back. The corrected reading
   was the stronger one (anchors held across a ~76-minor gap), which is worth remembering the next time a
   correction looks like a downgrade.
2. **Hold the tooling constant when the question is about the code.** Keeping the rext pin stale was
   uncomfortable and correct: `demo-1` went green under that exact tooling, so the advance was the single
   changed variable and *"it built"* means something. Bumping the pin would have bought currency at the
   price of attribution.
3. **Check the ordering before booking the crash.** A crash-looping Directus with a hard PostgreSQL error
   code looked like the finding and was the documented boot sequence. The cost of checking was one log
   read; the cost of not checking would have been a false defect in a milestone whose currency is
   whether its findings are real.

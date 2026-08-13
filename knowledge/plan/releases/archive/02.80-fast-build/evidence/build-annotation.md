# Demo bring-up — build-process annotation & timings

**Purpose.** A measured, phase-by-phase anatomy of what a `/demo-down --purge` + `/demo-up` cycle actually
does and where the wall-clock goes — captured as raw material for a future *build-speedup* roadmap.
Everything below is **MEASURED** on one instrumented run unless explicitly marked `ESTIMATE`.

**Status:** first capture. One cold-images run, `n=1` — treat per-phase numbers as indicative, not p50.

---

## 1. The run

| | |
|---|---|
| **Date** | 2026-07-27 07:03:47Z → 07:15:00Z |
| **Host** | `billion.taildc510.ts.net` (Tailscale VM), Linux 6.8.0-134 x86_64 |
| **Stack** | `demo-1`, offset 10000, `--public-host billion.taildc510.ts.net` |
| **rext pin** | `july-jitter-v27-close-followups` @ `a5b1288` (on origin — rung zero verified) |
| **Platform clones** | `DEMO_ADVANCE_CLONES=pinned` → `clones.pin.json` (app `v1.351.1`, cms `v0.255.1`, jobsimulation `v0.253.0`) |
| **Command** | `rosetta-demo down 1 --purge` then `up-injected.sh 1 --public-host …` |
| **Outcome** | `up_rc=0`; `autoverify.json` → `{"warnings":0,"green":true}`; `demopatch.log` **empty** (all 23 patches applied, none refused/skipped) |
| **Artifacts** | `billion:/home/devops/panorama/cleanbuild-20260727T070347Z.log` (+ `.samples.tsv`), driver `reset-clean-build.sh` |

> **State the environment with every number** (`latency-budget.md` rule). These are **tailnet-VM** numbers,
> not laptop numbers.

### Host capacity — read this before trusting any comparison

| Resource | Value | Note |
|---|---|---|
| vCPU | **8** | |
| RAM | **7.3 GiB** | **below** the tooling's own 12 GiB floor — `up-injected.sh:258` warns every run |
| Swap | 15 GiB | |
| Disk `/` | 193 G, **40 G free** at start (80 % used) | floor is 20 GiB (`up-injected.sh:298`) |
| Docker build cache | **102.7 GB across 614 entries** (88.08 GB reclaimable) at start | |

### Cache state — what "clean" meant here

`--purge` removes the **stack's images** and its data dir. It does **not** touch the BuildKit **layer cache**,
which persisted. So this run is **"images purged, layer cache warm"** — the normal, sanctioned path.
A genuinely cold run (`docker builder prune -af` as well) would be **slower**; that variant is **not measured**.

Observed per-image cache hits: `next-web` 2 CACHED steps, `hiring` 2, `studio-desk` 5 (all build steps),
`app`/`cms`/`jobsimulation` 5 each.

---

## 2. Headline

| | |
|---|---|
| **Total cycle** | **672.4 s — 11 m 12 s** |
| Teardown (`down 1 --purge`) | 19.7 s (2.9 %) |
| Bring-up (`up-injected.sh`) | **650.7 s (96.8 %)** |
| **UI-tier image builds (3 images)** | **446.4 s — 66.4 % of the whole cycle** |
| **Docker image export/unpack alone** | **288.4 s — 42.9 % of the whole cycle** |
| Everything else (backends, compose, set-dress, verify) | 204.3 s (30.4 %) |

**One sentence:** two Next.js images that together weigh **9.4 GB** dominate the cycle, and **most of their
cost is not compilation — it is writing the image to disk.**

---

## 3. Top-level phase table

| # | Phase | Start (s) | End (s) | Duration | % | Cacheable / skippable? |
|---|---|---|---|---|---|---|
| P0 | pre-state capture | 0.00 | 0.75 | 0.8 s | 0.1 % | instrumentation only — not part of a normal run |
| P1 | **teardown** `rosetta-demo down 1 --purge` | 0.75 | 20.47 | **19.7 s** | 2.9 % | no (purge is the point) |
| P2 | post-teardown capture | 20.47 | 20.96 | 0.5 s | 0.1 % | instrumentation only |
| P3 | `tailscale serve reset` | 20.96 | 20.96 | 0.0 s | 0.0 % | — |
| P4 | **bring-up** `up-injected.sh` | 20.96 | 671.66 | **650.7 s** | 96.8 % | see §4 |
| P5 | post-state capture | 671.66 | 672.38 | 0.7 s | 0.1 % | instrumentation only |

---

## 4. Bring-up sub-phases (P4 = 650.7 s)

| # | Sub-phase | Start | End | Dur | % of cycle | Serial? | Cacheable / skippable |
|---|---|---|---|---|---|---|---|
| 4.1 | host pre-flight (Go/atlas, VM RAM, disk, ssh-agent) | 20.96 | 37.31 | 16.4 s | 2.4 % | serial | `DEMO_NO_HOST_PREFLIGHT=1` |
| 4.2 | secrets provision + clone ensure/advance | 37.31 | 37.79 | 0.5 s | 0.1 % | serial | idempotent, already fast |
| 4.3 | build-scratch clones + disarmed-colony inject + app demopatches (×3) | 37.79 | 40.58 | 2.8 s | 0.4 % | serial | no |
| 4.4 | **backend images ×5 — built in PARALLEL** (app, cms, jobsimulation, fake-fapi, fake-bapi) | 40.58 | 66.77 | **26.2 s** | 3.9 % | **parallel** ✅ | layer cache |
| 4.5 | `stackseed` build + roster export + Clerkenstein wiring | 66.77 | 67.90 | 1.1 s | 0.2 % | serial | — |
| 4.6 | **`next-web` image** (9 demopatches + build) | 67.90 | 295.23 | **227.3 s** | **33.8 %** | **serial** ❌ | patch-set fingerprint reuse; `DEMO_NO_UI=1` |
| 4.7 | **`studio-desk` image** (5 demopatches + build) | 295.23 | 306.42 | 11.2 s | 1.7 % | **serial** ❌ | fully cached this run |
| 4.8 | **`hiring` image** (7 demopatches + build) | 306.42 | 514.35 | **207.9 s** | **30.9 %** | **serial** ❌ | same |
| 4.9 | `compose up` — **also builds 5 more images inline** (sentinel, storage, roadrunner, graphql, postgresql) | 514.35 | 551.09 | 36.7 s | 5.5 % | mixed | layer cache |
| 4.10 | registry record + `tailscale serve` + egress check | 551.09 | 570.28 | 19.2 s | 2.9 % | serial | — |
| 4.11 | **set-dress** (Directus provision + snapshot replay + seed) | 570.28 | 668.80 | **98.5 s** | **14.6 %** | serial | `DEMO_NO_SETDRESS`-class knobs |
| 4.12 | post-seed fixups (Casbin reload, FK indexes, Directus drift, URL rewrite) | 668.80 | 669.07 | 0.3 s | 0.0 % | serial | no |
| 4.13 | cockpit + ant-academy start | 669.07 | 669.56 | 0.5 s | 0.1 % | serial | — |
| 4.14 | autoverify | 669.56 | 671.66 | 2.1 s | 0.3 % | serial | non-fatal by design |

> **The UI tier is built serially on purpose** — `up-injected.sh` logs
> *"building the UI tier serially (before compose up) — next-web, then studio-desk, then hiring"*.
> On a 7.3 GiB box that is a defensible RAM guard, but it is the single largest structural cost. See §7-L1.

### 4.11 set-dress detail (98.5 s)

| Step | Dur | Volume |
|---|---|---|
| per-stack Directus provision start | 1.5 s | |
| **Directus `bootstrap` + `taxonomy` replay** | **78.0 s** | 10 tables, **330,261 rows**, + reindex of 2 pgvector indexes |
| `directus` content replay | 0.8 s | 14 tables, 11,986 rows |
| Directus container restart + `/server/health` | 10.3 s | |
| `sim-embeddings` replay | 0.2 s | 4 tables, 1,490 rows |
| **stories seed** | **7.6 s** | taxonomy 42,790 + content-stories 2,037 + nonsim 6 rows |

---

## 5. Per-image build anatomy

### 5.1 `next-web` — 227.3 s wall, image **4.77 GB**

Build context: `stack-demo/next-web-app`, `Dockerfile.dev`, **single stage** on `node:24-alpine`.

| Step | Op | Time |
|---|---|---|
| #2 | load metadata `node:24-alpine` | 0.7 s |
| #5 | load build context | 0.2 s |
| #6/#7 | `corepack enable` / `WORKDIR` | **CACHED** |
| #8 | `COPY . .` | 1.1 s |
| #9 | `RUN rm -rf node_modules apps/*/node_modules packages/*/node_modules configs/*/node_modules` | 0.3 s |
| #10 | `RUN --mount=type=cache,…/pnpm/store pnpm install --frozen-lockfile` | 29.1 s |
| #11 | `RUN pnpm turbo build --filter=@anthropos/web-app --concurrency=1` | 53.2 s |
| **#12** | **exporting layers + unpacking to `demo-1-next-web:latest`** | **141.9 s** (of which *unpacking* 37.2 s) |

**Three pathologies visible in that table:**

1. **`#12` export = 141.9 s = 62 % of the build.** Not compilation — serialising a 4.77 GB image.
2. **`COPY . .` precedes dependency install** (#8 → #9 → #10). Any source byte — and the 9 demopatches
   guarantee changed bytes — invalidates `COPY`, then `rm -rf node_modules`, then `pnpm install`. The
   `--mount=type=cache` pnpm store softens the download but the install still re-runs (29.1 s).
3. **`--concurrency=1`** serialises turbo on an 8-core box.

### 5.2 `hiring` — 207.9 s wall, image **4.67 GB**

Same shape (rext `hiring.Dockerfile` filtering `@anthropos/hiring-app` out of the same next-web clone):
export **136.7 s**, turbo build 42.6 s, `pnpm install` 25.5 s.

### 5.3 `studio-desk` — 11.2 s wall, image **1.71 GB**

Build context `stack-demo/studio-desk`, `Dockerfile.dev`, **single stage** on `node:24-alpine`.
**All five build steps CACHED** this run (the patch set was byte-identical to the previous build), so only
`#11` export ran: **9.8 s**.

Layer ordering here is *correct* (`COPY package*.json` → `npm ci` → `COPY . .`), and it uses
`--mount=type=cache,target=/root/.npm`. But it is still **single-stage**: `npm ci` installs **dev**
dependencies and nothing is pruned, so the runtime image ships the whole toolchain —
**32,568 JS/CSS files, 266 MB, inside a 1.71 GB image**, to serve a Vite production bundle.

> Not the build bottleneck today (11 s), but it is 1.71 GB of export cost every time the patch set changes,
> and it is the same single-stage pattern as the two expensive images.

### 5.4 Image inventory after the run

| Image | Size |
|---|---|
| `demo-1-next-web` | **4.77 GB** |
| `demo-1-hiring` | **4.67 GB** |
| `demo-1-studio-desk` | **1.71 GB** |
| `demo-1-postgresql` | 475 MB |
| `demo-1-cms` | 468 MB |
| `demo-1-app` | 448 MB |
| `demo-1-jobsimulation` | 430 MB |
| `demo-1-graphql` | 155 MB |
| `demo-1-roadrunner` / `sentinel` / `storage` | 71 / 56 / 51 MB |
| `demo-1-fake-fapi` / `fake-bapi` | 28 / 27 MB |

**The three Node images = 11.15 GB of ~12.6 GB total.** The seven Go images together are 1.7 GB.

---

## 6. Resource profile — what is *not* the bottleneck

Sampled every 10 s for the whole run (67 samples, `cleanbuild-*.samples.tsv`):

| Metric | Value |
|---|---|
| **peak load1** | **4.90** (of 8 cores) |
| avg load1 | 2.26 |
| peak memory used | 5,405 MB (of 7.3 GiB) |
| peak swap used | 2,452 MB |
| min disk available | 34 GB |
| top-CPU process during export windows | `unpigz` (~40 %) |

**The box was never CPU-saturated** — average load 2.26 on 8 cores, peak 4.90. Memory peaked at ~74 % with
modest swap. So despite the 7.3 GiB-vs-12 GiB warning, **RAM was not the limiter on this run**; the cycle is
dominated by **serialised I/O** (image export/unpack, compression) and **deliberate serialisation** of the UI
tier. Raising RAM alone would buy little; **parallelising and shrinking the images** is where the time is.

---

## 7. Where the time actually goes — ranked levers

Ranked by measured seconds recoverable. `ESTIMATE` marks savings not yet demonstrated.

| # | Lever | Evidence | Est. saving | Shape |
|---|---|---|---|---|
| **L1** | **Multi-stage the two Next.js images** — ship `.next/standalone` + static instead of the full build tree with dev deps. 4.77 GB → a few hundred MB collapses the 141.9 s + 136.7 s export steps **and** the 85.7 s unpack leg (L9). | §5.1 #12, §5.2, §5.4 | **~200–250 s** `ESTIMATE` | Dockerfile change = **platform-repo edit** → must be a `demopatch`, or an rext-owned Dockerfile like the existing `hiring.Dockerfile` |
| **L2** | **Build `next-web` and `hiring` in parallel** instead of serially — *and reorder the UI tier to overlap `compose up`*. See §8.1: the serialization is **unconditional**, and the UI tier runs *before* compose, so postgres boot + 4 atlas migrations + replay + seed idle for ~7.5 min. | 4.6/4.8, §6, §8.1 | **~200 s** `ESTIMATE` | rext `up-injected.sh` + a RAM gate |
| **L3** | **Move `COPY . .` after dependency install** in the next-web Dockerfile so the 29.1 s + 25.5 s `pnpm install` layers survive a source-only change (which every demopatch is). See §8.2 — the install layer has **never once been reused**. | §5.1 #8→#10, §8.2 | **~55 s** + an ~8 GB/cycle disk leak | platform-repo edit → demopatch / rext Dockerfile |
| **L4** | **Drop `--concurrency=1`** from `pnpm turbo build` on an 8-core host. | §5.1 #11 | ~20–35 s `ESTIMATE` | build-arg or demopatch |
| **L5** | **Speed the taxonomy replay** (78.0 s for 330 k rows + 2 pgvector reindexes). Candidates: build the vector indexes **after** COPY in one pass, `UNLOGGED`-then-`SET LOGGED`, or ship a pre-built PG data dir. The 3 replay surfaces are independent and run serially. | §4.11 | ~30–50 s `ESTIMATE` | rext `stack-snapshot` |
| **L6** | **Prune the BuildKit cache on a schedule.** It grew **102.7 GB → 105.4 GB** in one cycle and reclaimable went **88.08 → 99.66 GB**; free disk 40 G → 38 G. Not a *time* win, but the documented ENOSPC failure mode surfaces as a cryptic *"redis exited (1)"* (M239-F1). With §8.2's ~8 GB/cycle leak the real runway is **~4–5 cycles**, not 15–20. | §1, §8.2, §9 | 0 s (risk, not time) | ops / rext teardown |
| **L7** | **Multi-stage `studio-desk`** — `npm ci --omit=dev` in a runtime stage; 1.71 GB → ~200 MB. Small today (11 s) but it is 9.8 s of export on every patch-set change, and shrinks the runtime footprint. | §5.3 | ~8 s `ESTIMATE` | platform-repo edit → demopatch / rext Dockerfile |
| **L8** | **Cache the Directus bootstrap + restart** (15.6 s of container boot latency, not data work — the most compressible slice of set-dress). | §4.11 | ~15 s `ESTIMATE` | rext |
| **L9** | **The containerd-snapshotter unpack leg** — 38.3 + 37.6 + 9.8 = **85.7 s** paid *unconditionally*, even on a 100 %-cached build (studio-desk exported 0 new layer bytes and still paid 9.8 s, because `--purge` had removed the image). Not a build flag; the fix is **image shrink = L1**. Disabling the snapshotter is a host `daemon.json` change and a bad trade (it would orphan all 26 existing images). | §5.1/§5.2/§5.3 | folded into L1 | host config — *not* rext, *not* demopatch |
| **L10** | **Small serial fat**: ~12 serial `git fetch`es + 23 serial `demopatch revert` shells in `ensure-clones.sh`; Go tooling compiled 4–5×/bring-up (`stacksecrets` into a throwaway `mktemp -d`, `stackseed` **twice**); the 4 independent `atlas migrate apply` targets run serially; phase 39 re-emits the *entire* tailscale serve plan just to add one port. | §8.4 | ~20–50 s `ESTIMATE` | rext |

**L1 + L2 + L3 together plausibly take the cycle from ~11 m to ~4–5 m** `ESTIMATE` — because they attack the
same 446 s block from three angles (smaller images to export, exported concurrently, with dependency layers
that actually survive).

---

## 8. Deeper findings (from a 21-agent adversarial sweep, 12 confirmed / 2 refuted)

These were verified independently of the timed run and go beyond what the wall-clock alone shows.

### 8.1 The UI-tier serialization is unconditional — and its stated reason no longer applies

`build_frontends()` (`up-injected.sh:1240-1263`) has **exactly one** conditional: `[ "$NO_UI" = 1 ]`.
There is no RAM check and no parallel path. `preflight_vm_ram()` (`:246-271`) declares its `bytes`/`gib`
**`local`**, assigns no global and returns no verdict — it only logs *"(non-fatal: continuing)"*. So the
7.3 GiB warning is **cosmetic**; nothing branches on it.

`git log -S "SERIALLY"` dates the serialization to `2b7b32f` (M19, 2026-06-09) — **a month before**
`--public-host` (M212, 2026-07-11) made billion a demo host. The stated rationale (`:444-445`) is *"so the
build RAM spike never overlaps the Go builds"* — but the Go builds finish at T+66.8 s and the UI tier starts
at T+67.9 s, so they never overlap anyway.

**Two separable wins:** (a) the UI tier runs at `:1877` and `compose up` at `:1924`, so postgres boot,
4 atlas migrations, snapshot replay and the seed sit idle for ~7.5 min — **reordering is rext-only and
conflicts with nothing** (runtime RAM is ~0.66 GiB); (b) studio-desk sets **no** `NODE_OPTIONS` (Vite/esbuild)
and costs 9.8–24.6 s — overlapping *it* is nearly free. Only next-web and hiring carry
`--max-old-space-size=4096`, and that is a V8 *ceiling*, not a reservation.

### 8.2 The `pnpm install` layer has never once been reused — and leaks ~8 GB per cycle

`docker buildx du --verbose` holds **16 distinct** `mount / from exec /bin/sh -c pnpm install --frozen-lockfile`
entries at **4.029 GB each = 64.46 GB** — **61 % of the entire build cache** — every one with
`Usage count: 1`, paired with 16 distinct `COPY . .` entries. Two per bring-up (next-web + hiring).

The layer ordering is the *amplifier*; the *trigger* is that the context genuinely drifts every build.
It is **not** the minted Clerk pk or the `.env` overlays — `demo-stack/frontend/next-web.dockerignore`
(copied in transiently at `up-injected.sh:620`/`:1108`) already excludes `.env*`, `.git`, `node_modules`,
`.next`, `.turbo`, `dist`, `*.log`. But **`next-web-app` itself has no `.dockerignore`**, so the 82 MB `.git`
of a fresh ephemeral clone is in context.

**Verified safe to fix:** all 8 next-web/hiring demopatch manifests target `.ts`/`.tsx` sources — **none**
touches a manifest or lockfile — so a manifests-first `COPY` would give a genuinely stable install layer.
It must copy the root `package.json` + `pnpm-lock.yaml` + `pnpm-workspace.yaml` **plus all 16 workspace
`package.json`s** (the install log reports *"Scope: all 16 workspace projects"*; a naive `COPY package*.json ./`
would break `--frozen-lockfile`).

### 8.3 The same monorepo is built twice

Both `next-web` and `hiring` build from `ctx=$DEMO/next-web-app` (`up-injected.sh:490` and `:1008`) with
instruction-identical Dockerfiles down to the install line, differing only in the turbo `--filter`, port,
CMD and demopatch set. Verified at diff_id level: the bottom 6 layers (~214 MB) **are** shared and log
CACHED in both; the 3.37 GB install layer is **not** (`126579da…` vs `dab7fc6b…`).
**Recoverable ~145–160 s** — not the full 207.9 s, since hiring's 42.6 s turbo build is genuine work.

### 8.4 `--purge` defeats every image cache, including 5 hidden ones

`rosetta-demo:336-341` does `docker images | grep -E "^demo-$n-" | xargs docker rmi -f`, removing the three
UI images **and** 5 images that `docker compose up` then rebuilds inline (postgresql, graphql, sentinel,
storage, roadrunner) with **no per-service log file to attribute them to** — 15.3 s warm here, 120–300 s
`ESTIMATE` if truly cold. So the three cache-reuse checks (`:562`, `:849`, `:1077`) can **never** hit on a
purge cycle.
**Operational lever:** plain `rosetta-demo down N` (no `--purge`) keeps the images and makes a re-up cost
seconds — the fast-cycle option whenever a wiped DB is not required.

### 8.5 Corpus claims this run contradicts

| Claim | Where | Reality |
|---|---|---|
| studio slowness *"was pure memory starvation, **not a slow build**"* | `frontend-tier.md:248-250` | Refuted — export/unpack is 288.4 s and the box never exceeded load 4.90/8 |
| *"the ~3.7 GB build cache"* | `frontend-tier.md:271` | **~28× off** — the cache is 105.4 GB. `DEMO_DISK_MIN_GIB=20` is sized against the wrong number |
| *"~3 min per frontend"* | `frontend-tier.md:231` + 3 mirrors | Right for the two Next apps, **~7× wrong** for studio-desk; and `frontend-tier.md` mentions "hiring" **zero times in 623 lines**, so the total undercounts by a whole 208 s image |
| *"~3.7 GB first build"* | `up-injected.sh:794` | Measured **4.77 GB** / 4.67 GB |

> Retracting the "~3 min per frontend" number requires the v2.7 **C1 mirrored-count rule**
> (`demopatch-spec.md:222-226`): all four docs move in one commit, and `demo-up-defaults.md` is
> machine-fenced both ways by `stack-core/demo_knob_guard.py`.

### 8.6 Hazard spotted in passing (not a speed issue)

`$STACK/certs` **survives `--purge`** and the whole mint block is guarded on `[ ! -f $CERTS/fapi.crt ]`
(`up-injected.sh:1859`), so billion's `tailscale cert` minted **2026-07-11** has never been re-minted.
A 90-day cert can silently expire — around **2026-10-09**.

> **Caution for the roadmap:** L1/L3/L4/L7 all touch Dockerfiles in the canonical platform repos. Rosetta's
> hard rule is **zero platform-repo edits**, so each must land either as a sha-pinned `demopatch`
> (`demopatch-spec.md`) or as an rext-owned Dockerfile in the pattern `hiring.Dockerfile` already
> establishes. That precedent matters: **an rext-owned Dockerfile for next-web/studio-desk is already a
> sanctioned shape.**

---

## 9. Storage growth per cycle

| | Before | After | Δ |
|---|---|---|---|
| Images | 31 / 23.91 GB | 31 / 23.91 GB | 0 |
| Build cache | 614 / **102.7 GB** (88.08 GB reclaimable) | 606 / **105.4 GB** (**99.66 GB** reclaimable) | **+2.7 GB size, +11.6 GB reclaimable** |
| Disk free `/` | 40 G (80 %) | 38 G (81 %) | **−2 G** |

Every cycle leaves ~2 G behind and orphans ~11.6 G more cache.

---

## 10. Reproduce it

```bash
ssh root@billion
sudo -u devops setsid nohup bash /home/devops/panorama/reset-clean-build.sh \
  > /home/devops/panorama/cleanbuild.launch.log 2>&1 &

# then poll (the run is ~11 min; do NOT background-and-yield — foreground-poll)
RUN=$(cat /home/devops/panorama/.cleanbuild.current)
grep '@@PHASE@@' /home/devops/panorama/$RUN.log
```

The driver emits a `[elapsed|UTC]` prefix on every line, `@@PHASE@@` markers, and a 10 s resource sampler
to `$RUN.samples.tsv`. Per-image BuildKit output lands in
`stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/build-<service>.log`.

---

## 11. Not measured / open questions

- **Truly cold build** (`docker builder prune -af` first) — the numbers here keep the layer cache. Needed to
  size the *worst* case and the real value of L3.
- **n=1.** No p50/p95. Re-run 3× before a roadmap commits to a target.
- **Laptop vs tailnet VM.** Only the VM is measured.
- **Does `next-web`'s `Dockerfile.dev` have a production sibling?** If a multi-stage Dockerfile already
  exists upstream, L1 is a *selection* change, not a *new Dockerfile* — much cheaper. **Check first.**
- **Is the 4.9 peak load1 a plateau or a ceiling?** If BuildKit is I/O-throttled, L2's parallel win could be
  smaller than the naive 200 s.
- The corpus asserts `~3 min / ~3.7 GB first build` for next-web/hiring (`up-injected.sh:794`, and
  `frontend-tier.md`). **Measured: 227 s / 4.77 GB and 208 s / 4.67 GB** — the time is roughly right, the
  **size claim is ~1 GB stale**.

---

## Appendix — related finding (a runtime issue, not a build issue) — **FIXED**

The same instrumented run surfaced a **runtime** defect in studio-desk: Clerkenstein's fake **FAPI** never
registered `GET /v1/me/organization_memberships` (only the **BAPI**
`GET /v1/users/{userID}/organization_memberships` existed, `clerkenstein/clerk-backend/server.go:47`).
clerk-js 404'd and burned a 3-attempt retry ladder (**1606 ms + 2265 ms ≈ 4.05 s**) on **every** studio load,
gating `body.page-loaded` and therefore all page content. Browser FCP measured **6936 ms**.

Fixed on `fix/studio` (`rosetta-extensions@bc65850`) and verified live on billion: `canAccess` **4049 ms →
38 ms**, FCP **6936 ms → 2152 ms**, 404s **3 → 0**. Zero platform-repo edits.

**Full write-up: [`studio-slowness-investigation.md`](./studio-slowness-investigation.md)** — including why
M253's `< 1 s` first-paint gate stayed green through all of it, and the remaining byte-side lane (S1–S7).

> **Relevant to a build roadmap:** studio-desk is the **fastest** of the three UI images (9.8–24.6 s) and its
> Dockerfile layer ordering is the **correct** one. Do not conflate "studio is slow" (runtime, now fixed)
> with "the studio build is slow" (it is not).

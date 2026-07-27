# The bring-up build budget

_What "fast" means for a `/demo-down --purge` + `/demo-up` cycle, how the wall-clock is attributed, the
measured baseline, the headroom contract, and the harness that grades it. Authored by **M255** (v2.8 "fast
build") — before it, the corpus had **one** bring-up measurement, at **n = 1**, produced by a one-off shell
script that existed on a single box and was never in version control._

> **Why this doc exists.** [`latency-budget.md`](latency-budget.md) exists because the corpus asserted for four
> releases that a cockpit login took *"~2–5 s, which we can't shorten"* — it was 39 s and it was shortenable.
> This is the same shape, one layer down. The corpus asserted *"~3 min per frontend"*, *"~3.7 GB first build"*,
> and *"the ~3.7 GB build cache"*. Measured: **227 s / 4.77 GB**, **208 s / 4.67 GB**, and a build cache of
> **105.4 GB**. The time was roughly right; the sizes were stale by a factor, and the free-disk floor was
> sized against the wrong one of them.

---

## The rule that comes first

> **State the environment with every number.** Inherited verbatim from
> [`latency-budget.md`](latency-budget.md), and M255 earned it twice over:
>
> - the **same Dockerfile and the same context** produce a **4.84 GB** image on `billion` (x86_64) and a
>   **2.88 GB** image on the M1 Pro laptop (arm64) — a 40 % difference from `node_modules` alone;
> - `billion` uses the **containerd image store** and therefore pays an `unpacking to …` leg on every image;
>   the laptop uses **overlay2** and pays **none**. A measured 37.6 s unpack on one host is 0 s on the other.
>
> A bare "the frontend image is ~4.7 GB" is not a fact about the platform. It is a fact about one host.

---

## The definition

> **READY** := `up-injected.sh` has exited `0` **and** the stack's own `autoverify.json` reads
> `{"green": true, "warnings": 0}`.

Not "containers started". Not "the build finished". The same discipline
[`verification.md`](../verification.md) already applies to a bring-up: **UP means verified-working**. A cycle
that ends fast and red has not ended.

**The measured quantity is the FULL CYCLE** — `rosetta-demo down N --purge` *plus* `up-injected.sh N` — because
that is the unit an operator actually experiences, and because `--purge` is what makes the run comparable
(see *the cache variant*, below).

---

## The variant: "cold images, warm layer cache"

`--purge` removes **the stack's images** and its data dir. It does **not** touch the BuildKit **layer cache**.
Every gated number in v2.8 is measured in that state, and the reason is not convenience:

- it is **the sanctioned path** — what a real teardown-and-rebuild does;
- a genuinely cold run (`docker builder prune -af` first) is a *different, slower* thing, and **v2.8 cut it
  from the gate** (D-v28-8). It doubled a campaign that already runs against a ~4–5 cycle disk runway, to test
  a hypothesis the warm run already answers: three (size, unpack) points — **8.03 / 8.05 / 5.73 s per GB** —
  plus studio-desk paying **9.8 s of unpack having exported zero new layer bytes**, which is precisely the
  size-driven, cache-independent signature that mattered.

> **Do not prune before a baseline.** Reclaim **between** campaigns, never before one. Pruning first silently
> converts rep 1 into a truly-cold run and breaks comparability with every earlier number.

One **truly-cold** data point does exist, informationally: the laptop measurement below ran against an empty
BuildKit cache. It is reported, not gated.

---

## The per-phase attribution model

A cycle decomposes into two top-level phases and eleven bring-up sub-phases. `buildbench` derives the
sub-phases from `up-injected.sh`'s **own** progress lines against a checked-in anchor list, and the table
**must sum back to the bring-up phase** — a sub-phase table that does not add up means something is
unattributed and every attribution in it is suspect.

| # | phase | what it covers |
|---|---|---|
| P1 | **teardown** | `rosetta-demo down N --purge` |
| P4 | **bring-up** | everything below |
| 4.1 | `host_preflight` | Go/atlas/VM-RAM/disk/ssh-agent checks |
| 4.2 | `secrets_provision` | secret-coverage pre-flight + clone ensure/advance |
| 4.3 | `clones_and_inject` | build-scratch clones, disarmed-colony inject, the `app` demo-patches |
| 4.4 | `backend_builds` | 5 Go images, **in parallel** |
| 4.5 | `seed_tooling` | `stackseed` build, roster export, Clerkenstein wiring |
| 4.6 | `ui_next_web` | the next-web image — **serial** |
| 4.7 | `ui_studio_desk` | the studio-desk image — **serial** |
| 4.8 | `ui_hiring` | the hiring image — **serial** |
| 4.9 | `compose_up` | plus 5 more images built inline (postgresql, graphql, sentinel, storage, roadrunner) |
| 4.10 | `serve_and_egress` | registry record, `tailscale serve`, egress check |
| 4.11 | `set_dress` | Directus provision, snapshot replay, seed |
| 4.12 | `autoverify` | the bring-up-tail verify |

**Fail-closed.** A required anchor that does not match is reported as **missing** and flips the ledger's
`phases_complete` to false. Anchors that legitimately do not fire — the UI tier under `DEMO_NO_UI=1`, the
serve block without a public host — are reported **separately** as *not applicable*, and the ledger's env
snapshot is what makes that distinction derivable after the fact.

### Inside an image build

BuildKit's own `#N DONE Xs` lines are authoritative, and the export step is split into its two halves:

| leg | what it is | why it is called out |
|---|---|---|
| `exporting layers` | serialising + compressing the image | size-proportional, ~34 MB/s measured on billion |
| `unpacking to …` | the **containerd snapshotter** materialising it | paid **unconditionally**, even on a 100 %-cached build — studio-desk paid 9.8 s having exported nothing |

---

## The baseline

**`billion.taildc510.ts.net`** — 8 vCPU · 7.3 GiB RAM · 15 GiB swap · x86_64 · Linux 6.8.0-134 · Docker 29.6.2
with the **containerd image store** · demo-1 at offset 10000 with `--public-host`.

First, the **instrumented `n=1` anatomy** — the run every lever in this release was ranked from. **It is not
the gated number**; the `n=3` campaign below is, and it lands within 0.9 % of this.

| | |
|---|---|
| **Total cycle** | **672.4 s — 11 m 12 s** *(n = 1, 2026-07-27 — superseded as the baseline by the 666.29 s p50 below)* |
| teardown | 19.7 s (2.9 %) |
| bring-up | 650.7 s (96.8 %) |
| **UI-tier image builds (3)** | **446.4 s — 66.4 % of the cycle** |
| **image export + unpack alone** | **288.4 s — 42.9 % of the cycle** |
| peak load1 | **4.90 of 8 cores** · avg 2.26 |
| peak memory | 5,405 MB of 7.3 GiB · peak swap 2,452 MB |

**One sentence: two Next.js images that together weigh 9.4 GB dominate the cycle, and most of their cost is
not compilation — it is writing the image to disk.**

Full anatomy, including the ranked lever table: `releases/02.80-fast-build/evidence/build-annotation.md`
in the plan tree.

### The n ≥ 3 campaign — **the gated baseline**

**`buildbench run 1 --reps 3 --profile billion --public-host billion.taildc510.ts.net --label m255-baseline`**,
2026-07-27, same host as above. Artefacts: `billion:/home/devops/panorama/m255/campaign/`. **All three reps
`rc=0`, `autoverify green: true, warnings: 0`, `phases_complete: true`, zero missing anchors, headroom assert
OK.**

> **The baseline is `n=3 p50 = 666.29 s (11 m 06 s)`.** It supersedes the annotation's **672.4 s** — which
> was `n=1` and, it turns out, an unusually representative one: the two differ by **0.9 %**. **Every v2.8
> reduction target is measured against 666.29 s.**

> 📌 **Provenance — read before quoting or re-deriving these numbers.** The campaign ran **09:59–11:37Z on
> 2026-07-27** (reps completed 11:11 / 11:26 / 11:37; `campaign.json` written 11:37), on an otherwise-idle
> `billion`. Possible **third-party activity on the host was reported from ~13:11Z** — **94 minutes after the
> last rep finished**, so there is no overlap and none of these figures is contaminated (user-confirmed
> 2026-07-27). Two corroborations, both independent of that question: three totals across two separate
> sessions cluster within 2 % (**658 / 666 / 672 s**), and rep-02's 881 s outlier is *causally* attributed —
> 206 s of its 215 s excess sits in two sub-phases, matched to a reclaim that evicted 7 cache records /
> 356.8 MB. Host contention smears across sub-phases; it does not concentrate like that.
>
> **Re-confirm on the first post-freeze campaign** (`billion` is under a user freeze until ~2026-07-29, see
> [`state.md`](../../../knowledge/plan/state.md) § Process flags): the **n=3 p50**, spike (a)'s
> **146.8 s → 2.9 s** export cut, and spike (d)'s **disk-bound** attribution — the three timing-derived claims.
> The **barrier verdict needs no re-confirmation**: `4.84 GB → 379 MB` is bytes on disk, not a stopwatch, and
> a 50× export reduction cannot be produced by host contention.

| | p50 | min | max |
|---|---|---|---|
| **total cycle** | **666.29 s** | 658.15 s | 881.01 s |
| teardown (P1) | 28.08 s | 24.99 s | 33.17 s |
| bring-up (P4) | 633.15 s | 633.11 s | 852.92 s |

Per sub-phase, p50 (rep-01 / rep-02 / rep-03):

| sub-phase | p50 | share | reps |
|---|---|---|---|
| `ui_next_web` | **216.47 s** | 32.5 % | 214.06 / 249.89 / 216.47 |
| `ui_hiring` | **207.26 s** | 31.1 % | 207.26 / 204.65 / 209.12 |
| `set_dress` | 97.07 s | 14.6 % | 97.07 / 97.37 / 96.46 |
| `compose_up` | 36.26 s | 5.4 % | 36.26 / 44.40 / 35.78 |
| `backend_builds` | 25.42 s | 3.8 % | 25.42 / 30.32 / 24.59 |
| `serve_and_egress` | 18.73 s | 2.8 % | 18.85 / 18.73 / 18.09 |
| `host_preflight` | 14.63 s | 2.2 % | 14.63 / 15.52 / 14.47 |
| `ui_studio_desk` | 12.41 s | 1.9 % | 12.41 / **185.27** / 12.13 |
| `clones_and_inject` | 2.42 s | 0.4 % | 2.37 / 2.48 / 2.42 |
| `autoverify` | 2.12 s | 0.3 % | 2.20 / 2.12 / 2.09 |
| `seed_tooling` | 1.01 s | 0.2 % | 1.70 / 1.01 / 0.74 |
| `secrets_provision` | 0.88 s | 0.1 % | 0.88 / 1.16 / 0.79 |

**The n=1 headlines all survive n=3, and two get sharper:**

- **UI-tier image builds** = **436.1 s — 65.5 %** of the cycle (n=1 said 446.4 s / 66.4 %).
- **Export + unpack alone** = **307.5 s — 46.2 %** at p50, summed across all nine built images
  (305.9 / 359.6 / 307.5 per rep; n=1 said 288.4 s / 42.9 %). Nearly half the cycle is still *writing images
  to disk*, and n=3 puts that share **higher**, not lower.
- **Peak `load1` 4.06 / 4.56 / 4.22** — every rep well under the clause-1 limit of `cores − 2 = 6`, and all
  three **below** the n=1 figure of 4.90. Spike (d)'s "not a CPU plateau" reading holds at n=3.

**The headroom model was validated, not just applied.** The assert predicted
`heap_commitment = 1 lane × 3900 MiB + 1500 MiB idle = 5400 MiB`. The three reps peaked at **5446 / 5579 /
5398 MB** — within **~3 %**. `lane_heap_measured_peak_mib = 3900` now rests on a campaign, not on one run.

> **One honest caveat on spike (d).** 1 of 221 samples reached **100 % disk `%util`** (rep-03). Average util
> stayed 20.8–23.9 %, so the "not an I/O ceiling" conclusion stands — but *"zero samples ≥ 90 %"* is an **n=1**
> claim and must not be restated at n=3.

#### rep-02 is the finding, not the outlier to discard

rep-02 cost **881.01 s — 32 % above p50** — and 206 s of that 215 s excess sits in two sub-phases:
`ui_studio_desk` went **12.41 s → 185.27 s** (+173 s) and `ui_next_web` **216.5 s → 249.9 s** (+33 s).
Studio-desk's cache chain broke
at step 2 of 6 (`WORKDIR /app` uncached, 1 of 11 steps cached against 5 in the neighbouring reps), so it paid a
full **`npm ci` of 136.5 s** that the other two reps got for free.

**The cause is the reclaim step, and it refuted this doc's own first draft of the reasoning for it.** That
draft justified `--filter until=24h` on the grounds that *"the base-image and CACHED-step records are touched
by every rep, so their last-used clock keeps resetting and they survive."* Measured:

| | records | build cache pruned |
|---|---|---|
| reclaim after **rep-01** | 546 → **539** (−7) | **356.8 MB** |
| reclaim after **rep-02** | 581 → **581** (−0) | **0 B** |

rep-01 *did* serve studio-desk's chain from cache — and the first reclaim evicted it anyway. **Being served
from cache does not reliably refresh a BuildKit record's last-used clock.** Seven records and 356.8 MB bought
+173 s of wall clock.

Three consequences, all of which are why the campaign protocol reads the way it does:

1. **Report `p50`, never the mean.** The mean of these three reps is 735.2 s — **10 % high**, describing a
   cache eviction rather than a bring-up.
2. **`n ≥ 3` is a floor, not a nicety.** At `n=2` this campaign would have reported **~773 s** and every v2.8
   lever would have been priced against a number that does not exist.
3. **The eviction is one-off, not per-rep.** It clears the *pre-campaign-age* records; from rep 3 on the
   campaign is in steady state (rep-01 658–666 s ≈ rep-03). **Budget a warm-up rep** — or read `min` alongside
   `p50`, which here agree to within 1.2 %.

### The informational laptop run

**M1 Pro · macOS 25.1 · Docker Desktop 28.5.1 · 10 VM CPUs · 9,937 MiB VM RAM · 58 GiB VM disk · arm64 ·
overlay2 · BuildKit cache EMPTY (truly cold).** Not gated — every v2.8 gate is measured on `billion`.

> **What was measured, and what was not — because the assert said no.** The intent was a full `n=1` cycle
> here. It did not run: `buildbench assert-headroom --profile laptop` **FAILED clause 1** —
> `peak load1 10.69 exceeded cores-2 (8)` — because the workstation was doing other work at the time. Disk
> and memory were fine (25.3 GiB free against a 19 GiB requirement).
>
> **That is the assert working, on its first live outing, and the result is reported rather than overridden.**
> A cycle measured on a host at load 10.7/10 would have produced a number whose slowness says nothing about
> the bring-up — precisely the class of number D-M255-1 exists to refuse. The lesson generalises: **a
> developer workstation is not a bench.** It has other jobs. `billion` is the gate host because it does not.
>
> What *is* measured below is one real `hiring.Dockerfile` build taken on a quiet box — the dominant leg, and
> enough to fix the laptop profile's per-lane memory peak and to expose three host-shape differences that
> matter more than a wall-clock total would have. The one field this leaves derived rather than measured is
> the profile's `projected_image_gib`, and it is labelled as such in `laptop.json`.

One real `hiring.Dockerfile` build, 2026-07-27:

| | laptop (arm64, truly cold) | billion (x86_64, warm layer cache) |
|---|---|---|
| `pnpm install --frozen-lockfile` | **74.4 s** | 25.5 s |
| `pnpm turbo build --filter=hiring-app` | **140.4 s** | 42.6 s |
| export | **74.9 s** (layers only) | 136.7 s (99.0 layers + 37.6 unpack) |
| image | **2.88 GB** | 4.67 GB |
| total | **296 s** | 208 s |
| VM memory: idle → peak | 615 → 4,838 MiB | 1,500 → 5,405 MiB |

**Three things a laptop comparison teaches, and none of them is "the laptop is faster":**

1. **Compilation is 3.3× SLOWER here** (140.4 s vs 42.6 s) despite ten cores and more RAM. A laptop is not a
   faster `billion`; it is a differently-shaped host. This is why every v2.8 gate is measured on `billion`.
2. **There is no unpack leg** — overlay2, not containerd. Lever L9 is a `billion`-only phenomenon.
3. **The image is 40 % smaller** for arch reasons alone.

---

## The headroom contract

Three clauses, evaluated against a **measured, checked-in host profile**
(`rosetta-extensions/stack-core/hostprofiles/{billion,laptop}.json`). Any failure fails the whole assert.

| # | clause | fails when |
|---|---|---|
| 1 | **CPU** | peak `load1` > `cores − 2` |
| 2 | **memory** | `lanes × measured-per-lane-peak + idle` > 80 % of the memory budget |
| 3 | **disk** | free < `disk_floor_gib` + `projected_image_gib` |

**Clause 2 uses the MEASURED per-lane peak, not the V8 ceiling — deliberately.** The effective ceiling is
**8192 MiB** (`apps/web/package.json:98` and `apps/hiring/package.json:92` re-assign
`NODE_OPTIONS=--max_old_space_size=8192` **inline** for the `next build` child, overriding the Dockerfile's
`4096` — so `ENV NODE_OPTIONS` is **not** a usable seam for lowering it). But a ceiling is not a reservation:
against a 7,500 MiB budget, a ceiling-based test would "prove" that `billion` cannot run the single lane it
demonstrably runs every day. The profiles therefore record both, and the assert uses the measured one —
**3,900 MiB** on billion, **4,223 MiB** on the laptop.

**And the measured one held up.** Clause 2's arithmetic predicts a whole-host commitment of **5,400 MiB** for
one lane on billion; the n=3 campaign peaked at **5,446 / 5,579 / 5,398 MB**. A headroom model that is right
to ~3 % across three independent cycles is a model, not a guess — which is what makes the *refusals* it issues
(the laptop, below) worth honouring rather than overriding.

**The derived consequence M257 has to price in:**

> `max_parallel_ui_lanes = floor((0.8 × budget − idle) / measured-lane-peak)` = **1 on billion**, **1 on the
> laptop.** Neither host fits two concurrent Next.js build lanes in RAM.

**On macOS the budget is the Docker VM allocation — never host totals.** A 16 GiB laptop has a 9,937 MiB
engine. Reading host `free`/`df` there over-states RAM by ~60 % and disk without bound; that is precisely how
the M239-F1 ENOSPC walked past a GREEN pre-flight.

### Two contracts, deliberately different — D-M255-1

`up-injected.sh`'s pre-flights are **advisory by design**: *"never block a genuinely good bring-up on a soft
heuristic"* (`:279`, `:319`). M255 does **not** retract that, and the retraction it *does* make is narrower
and different: `preflight_vm_ram()` declares its variables `local`, assigns no global and returns no verdict,
so **nothing branches on it** — it is advisory in the sense of *inert*, which is not the same as *tolerant*.

`buildbench` is not an operator. It is a **measuring instrument feeding a release gate**, and a gate number
measured on a host without headroom is not a number. So the identical assert **hard-fails** there.

| consumer | contract | on failure |
|---|---|---|
| `up-injected.sh` pre-flight | operator-facing | **warns, continues** |
| `buildbench` pre-rep + post-rep | gate-facing | **aborts the rep, exits 1** |

---

## The campaign protocol

**The binding constraint is the TRANSIENT, not the net.** A steady rep nets only **~1.7–2.1 GiB** of resident
disk (measured across the campaign: 38.85 → 36.23 GiB free over three reps and two reclaims), but *mid-cycle*
it swings roughly **18 GiB** — `--purge` frees the old images, the rebuild re-writes them, and the export leg
stages layers before unpacking them. Clause 3 is sized against that swing, which is why 36–42 GiB of free disk
is comfortable and 25 GiB is the floor.

1. **Declare the starting state.** Every ledger records `docker system df` before and after, plus the Docker
   VM's free disk and the host's.
2. **Assert before the rep, hard.** `free ≥ disk_floor_gib + projected_image_gib` — **25 GiB** on billion
   (7 GiB reserve + the **measured 18 GiB peak consumption** of one cold-images cycle).
3. **Reclaim between reps, precisely:**
   ```bash
   docker builder prune -f --filter until=24h
   ```
   `until=<duration>` prunes records **not used within** that window, where `prune -af` would make the next rep
   **truly cold** and silently break comparability. That much is the reason for the filter and it holds.

   > **⚠️ But `until=24h` is NOT a guarantee that rep-touched records survive — measured, and it cost 173 s.**
   > The reasoning this step used to carry — *"CACHED-step records are touched by every rep, so their
   > last-used clock keeps resetting and they survive"* — is **false as stated**. In the n=3 campaign, rep-01
   > served studio-desk's whole chain from cache, and the reclaim immediately after it **evicted that chain
   > anyway** (7 records, 356.8 MB), so rep-02 paid a full 136.5 s `npm ci`. Being *served* from cache does not
   > reliably refresh a BuildKit record's last-used clock.
   >
   > It is a **one-off**, not a per-rep tax — it clears the pre-campaign-age records, and the next reclaim
   > pruned **0 B / 0 records**. Plan for it: **budget a warm-up rep, report `p50`, and read `min` beside it.**

4. **Expect the reclaim's value to come from the IMAGE prune, not the cache prune.** After rep-01: **5.321 GB**
   of dangling images against **356.8 MB** of build cache. After rep-02: **0 B** and **0 B**. Build-cache growth
   per steady rep is **+1.7 to +2.2 GB** (the rebuild rep: +4.4 GB) — an order of magnitude less than the
   ~11.6 GiB/rep this doc previously claimed.
5. **Never prune before the baseline.** (Restated because it is the mistake that costs a whole campaign.)

### ⚠️ A mid-campaign ENOSPC does not look like ENOSPC

It presents as the cryptic **`redis exited (1)`** (M239-F1) — *not* as a build error, *not* as "no space left
on device" in the failing phase. Under a speed campaign that is actively dangerous, because it reads as
**"the lever I just added broke the stack"**. Before debugging any bring-up failure that appears during a
build-speed change, run `df -h /` and `docker system df`.

`DEMO_DISK_MIN_GIB` was **20**, reasoned from *"the ~3.7 GB frontend build"* and *"the ~3.7 GB build cache"*.
Both numbers are stale by roughly an order of magnitude. It is now **25**, derived from the measurement above,
and the profile carries the same arithmetic so the operator warning and the gate cannot drift apart.

---

## The parallelism rule (union-apply)

`next-web` and `hiring` build from **the same clone** — `ctx=$DEMO/next-web-app` (`up-injected.sh:490`,
`:1008`) — with different demo-patch sets. Building them in parallel naively races the shared working tree
against three of `demopatch`'s own guards: **G2** (drift-refuse), **G4** (idempotent), **G5** (self-revert).
A G2 refusal is **non-fatal and silent**, so the failure mode is *an image that ships unpatched while the
bring-up grades green*.

> **THE RULE: apply the UNION once · build both images in parallel from the single clone · revert once, LIFO.**

It is safe because of what the 11 distinct manifests actually are — **5 shared**, **5 under disjoint `apps/*`
trees** (`apps/web` ×3 vs `apps/hiring` ×2), **1 shared-package outlier** — and it is *kept* safe by a fence,
`stack-core/union_apply_guard.py`, which derives both lists from `up-injected.sh` and refuses:

- a shared member that is not the same manifest;
- a single-image member under another app's tree (**CROSS-APP**);
- a single-image member under a shared tree with no written waiver (**UNWAIVED SHARED TREE**);
- a waiver left behind after its manifest became shared (**STALE WAIVER**);
- the `urls.ts` chain declared out of order or half-applied (its `pre_sha256` **is** the other's `post_sha256`).

Union-apply also **removes one apply/revert cycle of that chained pair** — exactly where G2 drift-refusals
historically bite.

### The outlier, and a correction to the decision that waived it

`next-web-ssr-graphql-origin` → `packages/graphql/src/server/server.graphql.ts` is applied by the next-web
build only, and D-v28-7 waived it as *"inert for the hiring image — behaviour-identical when
`WUNDERGRAPH_SSR_ENDPOINT` is unset"*.

**The premise is false.** `stack-injection/gen_injected_override.py:367` sets
`WUNDERGRAPH_SSR_ENDPOINT=http://graphql:8080/graphql` **on the hiring container too**, and `apps/hiring` does
import the patched module (`apps/hiring/src/app/api/bunny/recording/[sessionId]/route.ts` →
`createServerGraphQLClient`). Under union-apply, hiring's behaviour **changes**.

It changes for the **better**: that route currently resolves its SSR origin from the build-inlined *public*
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` and therefore carries the same blackhole M218 measured at 37.5 s for
next-web. But *"improves a latent defect"* is not *"inert"*, and the difference is operational: an inert patch
needs no proof, a behaviour change needs one. **M257 must re-verify the hiring recruiter Playthrough after
flipping union-apply on.**

### The image-isolation invariant

Stacks share the BuildKit **layer cache**. Stacks never share **images** — every image is tagged `demo-N-*`
and a bring-up refuses to reuse one whose baked offset endpoint, minted publishable key, or demo-patch-set
fingerprint does not match. **Cache layers, never images.** M257 asserts it.

---

## The harness

`rosetta-extensions/stack-core/buildbench.py`, at a tag pushed to origin.

```bash
buildbench.py run 1 --reps 3 --profile billion --public-host <magicdns> --label baseline
buildbench.py report .buildbench/baseline-<ts>
buildbench.py assert-headroom --profile laptop --lanes 2      # exit 1 = this plan oversubscribes the host
buildbench.py env-snapshot                                    # every knob, value, and where it is read
buildbench.py parse --log <an older cycle log>                # back-fill a past run into the same schema
```

**Every ledger entry is self-describing** — the exact `argv` of both legs, the rext HEAD and tag (and whether
the clone was dirty), and **every `DEMO_*` knob with its effective value and whether that value came from the
environment or the script's own default**.

That last part is not bookkeeping. Neither `autoverify.json` (which emits only
`project/offset/warnings/green/ts` — `stack-verify/live/autoverify.sh:381-385`) nor the phase log records
**which services were in scope**, so **a `DEMO_NO_UI=1` cycle and a full-UI cycle leave indistinguishable
artefacts while differing by ~66 % of the wall clock**. A number nobody can attribute to a configuration is
not evidence.

**Exit codes** follow the `stack-core` guard convention: `0` ok · `1` a rep failed or the assert fired ·
`2` the harness could not run. An empty campaign directory reports as a **FINDING**, never as a pass.

A campaign writes `campaign.json` plus a `rep-NN/` per rep holding `ledger.json` (the full phase/build/sample
ledger), `samples.tsv`, and `cycle.log`. **The M255 baseline lives at
`billion:/home/devops/panorama/m255/campaign/`** and is what every v2.8 comparison re-derives from — re-read it
with `buildbench report`, don't re-type numbers out of this doc.

### Rung zero applies here too

A remote stack consumes `rosetta-extensions` **at a tag fetched from origin** (the M217 FATAL pin guard). A
`buildbench` that exists only in the authoring copy is **unreachable** to `billion`, and the failure looks like
a missing feature rather than a missing tag. `git push --tags` is part of shipping the harness. See
[`../verification.md`](../verification.md) pre-flight rung zero.

---

## See also

- [`../../../knowledge/plan/releases/02.80-fast-build/evidence/build-annotation.md`](../../../knowledge/plan/releases/02.80-fast-build/evidence/build-annotation.md)
  — the instrumented n=1 anatomy + the ranked lever table L1–L10 this release is designed from
- [`latency-budget.md`](latency-budget.md) — the *other* budget: click→ACCESS, and the doc this one is modelled on
- [`frontend-tier.md`](frontend-tier.md) — what the UI tier is and why it exists
- [`demopatch-spec.md`](demopatch-spec.md) — the 7 guards union-apply has to stay inside
- [`demo-up-defaults.md`](demo-up-defaults.md) — every knob the env snapshot captures, with its real default
- [`../safety.md`](../safety.md) — §3.5.4, the cert-renewal path and its silent-fallback failure mode
- [`../verification.md`](../verification.md) — what "verified-working" means, and rung zero

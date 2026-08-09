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

A cycle decomposes into two top-level phases and **twelve** bring-up sub-phases. `buildbench` derives the
sub-phases from `up-injected.sh`'s **own** progress lines against a checked-in anchor list, and the table
**must sum back to the bring-up phase** — a sub-phase table that does not add up means something is
unattributed and every attribution in it is suspect.

> **The count is TWELVE, not eleven** — corrected M257x. `BRINGUP_ANCHORS`
> (`rosetta-extensions/stack-core/buildbench.py:115-131`, entries at `:118-130`, @ rext `415240f`) holds **12**:
> `host_preflight · secrets_provision · clones_and_inject · backend_builds · seed_tooling · ui_next_web ·
> ui_studio_desk · ui_hiring · compose_up · serve_and_egress · set_dress · autoverify`. The **4.1–4.12 table
> immediately below already listed all twelve**, and so does the p50 table in "The baseline" — only this
> sentence was short, which is the shape a sum-back invariant is supposed to catch and a prose count is not.
> Derived, not counted by eye:
>
> ```sh
> python3 -c "import re;s=open('stack-core/buildbench.py').read();\
> b=re.search(r'BRINGUP_ANCHORS.*?= \[(.*?)\n\]',s,re.S).group(1);\
> n=re.findall(r'\(\"([a-z_]+)\",',b);print(len(n),n)"     # -> 12
> ```
>
> **Widened once** (Rule 57): `grep -rn "BRINGUP_ANCHORS" .` across the whole rext tree returns exactly two
> hits — the definition at `buildbench.py:115` and its single consumer at `:250`. There is no second anchor
> list, so the number does not move. Three of the twelve are conditional (`ui_*` on the UI tier,
> `serve_and_egress` on `--public-host`) and are reported as *not applicable* rather than missing when their
> feature is off — a conditional phase is still a declared phase.

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

> 🛑 **SUPERSEDED — `D-v28-15` (2026-07-31) retired `odysseus` too, and NOTHING below has been re-pointed.
> Measured M257x iter-225.** `D-v28-15` supersedes `D-v28-14` in the same day: `billion` is demo-deployment
> only, **`odysseus` is retired**, and dev/test is **LOCAL to the new Mac**. Read every `odysseus` sentence
> below as *"the host this doc was written for, which no longer exists for this purpose"* — this doc names
> `billion` 44 times, `odysseus` 26 and `laptop` 22, and **names the sanctioned dev host zero times.**
>
> **The operational consequence, measured on the sanctioned host (`Mac16,11`, arm64, 12 cores, 24 GiB,
> Docker Desktop `overlay2`, VM 8 CPU / 11.67 GiB / 125.4 GiB disk with 64.97 GiB free):**
>
> 1. **There is no host profile for it.** `rosetta-extensions/stack-core/hostprofiles/` holds exactly two —
>    `billion.json` (8-core x86_64 Linux, containerd) and `laptop.json` (10-core / 16 GiB M1 Pro, 9,937 MiB
>    VM, 58 GiB VM disk). Neither describes this machine. **Both name `odysseus` as the gate host**, citing
>    the superseded `D-v28-14`, and **no `odysseus.json` was ever written** — so the host the tooling names
>    as its gate has no profile either.
> 2. **Nothing would notice if you used another host's profile.** `buildbench --profile <name>` is
>    operator-supplied; there is no host auto-detection. A missing *named* profile is correctly exit-2
>    (`load_host_profile`: *"a missing profile is exit-2 territory, never a pass"*), but a **present and
>    inapplicable** one is graded silently. `host_facts()` is recorded into the run JSON and **never
>    compared to the profile.** The harness already refuses an `autoverify.json` verdict *"that does not
>    describe the run under test"* — the same discipline is simply not applied to the host profile.
> 3. **`require_measured` does not cover this.** Clause zero fails a `None` *measurement input* (a dead
>    sampler, an unanswered disk probe). A profile describing the wrong machine supplies numbers, not
>    `None`, so clause zero passes.
> 4. **The VM sits below the build floor.** `up-injected.sh`'s `preflight_vm_ram` floor is **12 GiB
>    binary**; this VM is **11.67 GiB**, so the warning fires on every bring-up here. Non-fatal by design
>    (`DEMO_VM_MIN_GIB` overrides), and it is the unit trap M257x iter-12 already documented: Docker
>    Desktop's slider is decimal GB, so setting exactly "12 GB" yields ~11.2 GiB and never clears it.
> 5. **Disk is NOT the constraint here.** 64.97 GiB free inside the VM against `disk_floor_gib + `
>    `projected_image_gib` = **25 GiB**.
> 6. **CPU saturation probably is.** `laptop.json`'s own history is the precedent: a full cycle was
>    *attempted and REFUSED* by clause 1, `peak load1 10.69 > cores-2`, because *"a developer workstation is
>    not a bench; it has other jobs."* The sanctioned dev host is a developer workstation.
>
> **So gate clause 1 is not gradeable on the sanctioned host today**, and the blocker is not the bring-up —
> it is that **no measured profile exists for the only machine the release is allowed to test on.** What
> unblocks it is a `buildbench` measurement run on a quiet box producing a checked-in `Mac16,11` profile.
> Do not grade a cycle here against `laptop.json` or `billion.json`: a wall-clock figure never transfers
> between hosts, which is this section's own standing rule.

> ⚠️ **Host change — `D-v28-14` (2026-07-31): the gate host is `odysseus`, not `billion`.** Everything in this
> section records **where the M255 baseline was measured**, and stays true as such. It is **not** a statement
> about where v2.8's gates are graded: `billion` is the demo machine now (deploy-only, never dev/test), and
> every `buildbench` campaign from here on runs on **`odysseus.taildc510.ts.net`** — 8 cores · 7,780 MB RAM ·
> **zero swap** · 189 G free of 193 G · x86_64 · Linux 6.8.0-117 · Docker 29.6.2 · the **containerd image
> store** (`Storage Driver: overlayfs` with `io.containerd.snapshotter.v1`, the same class as `billion`, so the
> unpack leg is paid there too — *not* the laptop's classic overlay2 graphdriver).
>
> **`billion`'s figures do not transfer.** This doc's own first rule is why: the same Dockerfile and context
> yield **4.84 GB on `billion`** and **2.88 GB on the arm64 laptop** (§*The rule that comes first*).
> **`odysseus`'s own baseline is UNMEASURED at time of writing** — M257 owes it (`n ≥ 3`, plus a checked-in
> `odysseus.json` host profile), and until it exists there is **no odysseus number to quote**, projected or
> otherwise. A reduction target is measured against **the baseline of the host it runs on**.

**`billion.taildc510.ts.net`** — 8 vCPU · 7.3 GiB RAM · **16 GiB** swap · x86_64 · Linux 6.8.0-134 · Docker 29.6.2
with the **containerd image store** · demo-1 at offset 10000 with `--public-host`.

First, the **instrumented `n=1` anatomy** — the run every lever in this release was ranked from. **It is not
the gated number**; the `n=3` campaign below is, and it lands within 0.9 % of this.

| | |
|---|---|
| **Total cycle** on `billion` | **672.4 s — 11 m 12 s** *(n = 1, 2026-07-27 — superseded as the baseline by the 666.29 s p50 below)* |
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

### The n ≥ 3 campaign on `billion` — **that host's gated baseline**

**`buildbench run 1 --reps 3 --profile billion --public-host billion.taildc510.ts.net --label m255-baseline`**,
2026-07-27, same host as above. Artefacts: `billion:/home/devops/panorama/m255/campaign/`. **All three reps
`rc=0`, `autoverify green: true, warnings: 0`, `phases_complete: true`, zero missing anchors, headroom assert
OK.**

> **`billion`'s baseline is `n=3 p50 = 666.29 s (11 m 06 s)`.** It supersedes the same host's annotation figure
> of **672.4 s on `billion`** — which was `n=1` and, it turns out, an unusually representative one: the two
> differ by **0.9 %**. **A v2.8 reduction target is measured against the baseline of the host it runs on** — on
> `billion`, against its 666.29 s; on `odysseus`, against `odysseus`'s own baseline, which per `D-v28-14` is
> **not yet measured**.

> 📌 **Provenance — read before quoting or re-deriving these numbers.** The campaign ran **09:59–11:37Z on
> 2026-07-27** (reps completed 11:11 / 11:26 / 11:37; `campaign.json` written 11:37), on an otherwise-idle
> `billion`. Possible **third-party activity on the host was reported from ~13:11Z** — **94 minutes after the
> last rep finished**, so there is no overlap and none of these figures is contaminated (user-confirmed
> 2026-07-27). Two corroborations, both independent of that question: three totals across two separate
> sessions cluster within 2 % (**658 / 666 / 672 s**), and rep-02's 881 s outlier is *causally* attributed —
> 206 s of its 215 s excess sits in two sub-phases, matched to a reclaim that evicted 7 cache records /
> 356.8 MB. Host contention smears across sub-phases; it does not concentrate like that.
>
> **Re-confirm on `odysseus`'s baseline campaign.** This note used to say *"re-confirm on the first post-freeze
> campaign"* — that freeze has expired, and `D-v28-14` has since made `billion` demo-only, so there is no
> "first post-freeze campaign on `billion`" to wait for. The three timing-derived claims re-home to the
> `odysseus` baseline campaign M257 owes, and they **ride along free**: that campaign already runs `n ≥ 3` cold
> cycles with per-phase attribution, which is the whole instrument each of them needs. The three: the **n=3 p50
> of 666.29 s on `billion`**, spike (a)'s **146.8 s → 2.9 s** export cut, and spike (d)'s **disk-bound**
> attribution. **What a different host can re-confirm is the SHAPE, not the seconds** — an `odysseus` figure
> corroborates *that export dominates* or *that the run is not I/O-ceilinged*; it can never restate `billion`'s
> numbers.
>
> The **barrier verdict needs no re-confirmation**: `4.84 GB → 379 MB` is bytes on disk, not a stopwatch, and
> a 50× export reduction cannot be produced by host contention.

| | p50 | min | max |
|---|---|---|---|
| **total cycle** on `billion` | **666.29 s** | 658.15 s | 881.01 s |
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
overlay2 · BuildKit cache EMPTY (truly cold).** Not gated — a v2.8 gate is measured on the release's dedicated
bench host (`billion` for the M255 baseline; **`odysseus`** since `D-v28-14`), never on a workstation.

> **What was measured, and what was not — because the assert said no.** The intent was a full `n=1` cycle
> here. It did not run: `buildbench assert-headroom --profile laptop` **FAILED clause 1** —
> `peak load1 10.69 exceeded cores-2 (8)` — because the workstation was doing other work at the time. Disk
> and memory were fine (25.3 GiB free against a 19 GiB requirement).
>
> **That is the assert working, on its first live outing, and the result is reported rather than overridden.**
> A cycle measured on a host at load 10.7/10 would have produced a number whose slowness says nothing about
> the bring-up — precisely the class of number D-M255-1 exists to refuse. The lesson generalises: **a
> developer workstation is not a bench.** It has other jobs. A **dedicated** bench host is used because it does
> not — `billion` for the M255 baseline, and **`odysseus`** for every v2.8 gate since `D-v28-14`.
>
> What *is* measured below is one real `hiring.Dockerfile` build taken on a quiet box — the dominant leg, and
> enough to fix the laptop profile's per-lane memory peak and to expose three host-shape differences that
> matter more than a wall-clock total would have. The one field this leaves derived rather than measured is
> the profile's `projected_image_gib` — and **the only thing labelling it is prose**, inside `laptop.json`'s
> `notes` blob, **which no code reads**. The loader validates it *identically* to the four genuinely-measured
> numeric fields (`buildbench.py:409-412`), so nothing mechanical stops a derived number being quoted as one.
> **There is no `provisional_fields` list and no provenance object** — M257 owes the real mechanism (it is a
> carried M255 item); until then, treat "measured or derived?" as a question the profile answers only to a
> human who opens it.

One real `hiring.Dockerfile` build, 2026-07-27:

| | laptop (arm64, truly cold) | billion (x86_64, warm layer cache) |
|---|---|---|
| `pnpm install --frozen-lockfile` | **74.4 s** | 25.5 s |
| `pnpm turbo build --filter=hiring-app` | **140.4 s** | 42.6 s |
| export | **74.9 s** (layers only) | 136.7 s (99.0 layers + 37.6 unpack) |
| image | **2.88 GB** | 4.84 GB (see below) |
| total | **296 s** | 208 s |
| VM memory: idle → peak | 615 → 4,838 MiB | 1,500 → 5,405 MiB |

> **On the 4.67-vs-4.84 GB conflict — quote 4.84 GB.** `billion`'s hiring image was measured twice on
> 2026-07-27: **4.67 GB** in the `n=1` annotation (the `demo-1-hiring` tag) and **4.84 GB** in spike (a)'s
> baseline build; `up-injected.sh:300` hedges *"4.67-4.84 GB"*. **4.84 GB is the profile's value** — it is the
> divisor behind `laptop.json`'s `projected_image_gib: 12` (billion's measured 18 GiB cycle peak scaled by
> 2.88/4.84), and it is what §*The rule that comes first* and the spike (a) barrier verdict already use. An
> `odysseus.json` `projected_image_gib` must be derived from a **measured odysseus** image, not from either.

**Three things a laptop comparison teaches, and none of them is "the laptop is faster":**

1. **Compilation is 3.3× SLOWER here** (140.4 s vs 42.6 s) despite ten cores and more RAM. A laptop is not a
   faster `billion`; it is a differently-shaped host. This is why a v2.8 gate is measured on a **dedicated**
   bench host — `billion` for the M255 baseline, `odysseus` since `D-v28-14` — and never on a workstation.
2. **There is no unpack leg** — overlay2, not containerd. Lever L9 is a **containerd-image-store** phenomenon,
   not a `billion` one: it is paid on `billion` **and on `odysseus`** (both use the containerd image store) and
   on neither the laptop nor any other classic-overlay2 host.
3. **The image is 40 % smaller** for arch reasons alone.

---

## The headroom contract

**Four clauses** — three thresholds plus a **clause zero** — evaluated against a **measured, checked-in host
profile** (`rosetta-extensions/stack-core/hostprofiles/{billion,laptop}.json`; M257 adds `odysseus.json`). Any
failure fails the whole assert.

| # | clause | fails when |
|---|---|---|
| **0** | **measured-ness** (`require_measured`) | a **required input is `None`** — the instrument produced nothing, so there is no verdict to give |
| 1 | **CPU** | peak `load1` > `cores − 2` |
| 2 | **memory** | `lanes × measured-per-lane-peak + idle` > 80 % of the memory budget |
| 3 | **disk** | free < `disk_floor_gib` + `projected_image_gib` |

> **Clause zero is the one a brand-new host hits FIRST, and it is unpredictable from a three-clause reading.**
> `buildbench.py:451` names it in its own comment (implemented `:461-470`) and it fires in **all three**
> production call sites — `pre_rep_assert` requires `disk_avail_gib` (`:724`), the post-rep gate requires
> `peak_load1` **and** `disk_avail_gib` (`:1175`), the `assert-headroom` CLI requires `disk_avail_gib` (`:1470`).
> It exists because the first version *skipped* any clause whose input was `None`, so a rep whose sampler died
> handed in `peak_load1=None` and the gate **returned `ok=True` on a host it had never measured** — the same
> "an empty result read as a pass" class the rest of `stack-core` refuses. The failures it emits are named
> `unmeasured_peak_load1` / `unmeasured_disk_avail_gib`. **On a fresh host with an unpopulated sampler or a
> disk probe that did not answer, `unmeasured_disk_avail_gib` is the first thing you will see** — that is the
> assert working, not a broken profile.

> **Clause 1 says *peak*, and under the standalone CLI it is not one.** `buildbench run` genuinely takes the
> peak (`max` over the sampler's rows, `buildbench.py:1173`). `buildbench assert-headroom` has no campaign to
> sample, so it reads a **single instantaneous** `os.getloadavg()[0]` (`:1466`) — and the failure message still
> says *"peak load1 …"*. Read a CLI verdict as *"load right now"*: a quiet moment on a busy box passes, and a
> transient spike fails. It is the `run` path's peak that gates a number.

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

> `max_parallel_ui_lanes = max(1, min(ui_image_count, floor((0.8 × budget − idle) / measured-lane-peak)))`
> = **1 on billion**, **1 on the laptop.** Neither host fits two concurrent Next.js build lanes in RAM.

**Both guards in that formula are load-bearing, and an earlier draft of this doc omitted them.** The code
(`buildbench.py:430-436`) clamps at `ui_image_count` (3) and **floors at 1**; the bare `floor(…)` returns **0**
on an undersized host where the code returns **1**. `billion` and the laptop agree with the bare version by
coincidence (1.153 → 1, 1.736 → 1) — **a third host need not**, which matters now that `odysseus` is the gate
host and its profile is unwritten.

> **The floor of 1 is NOT a claim that one lane FITS** (`buildbench.py:425-428`, verbatim intent). A host too
> small for a single lane still has to build one, so the *plan* is 1 either way — on such a host it is
> `headroom_assert(lanes=1)` **clause 2** that fails, and that is where the fact gets reported. Do not read a
> return of 1 as "headroom verified"; read it as "the plan is one lane". The two are checked in different
> places on purpose. This is load-bearing for **L2**, whose entire premise is `max_parallel_ui_lanes = 1`.

**On macOS the budget is the Docker VM allocation — never host totals.** A 16 GiB laptop has a 9,937 MiB
engine. Reading host `free`/`df` there over-states RAM by ~60 % and disk without bound; that is precisely how
the M239-F1 ENOSPC walked past a GREEN pre-flight.

### Two contracts, deliberately different — D-M255-1

`up-injected.sh`'s pre-flights are **advisory by design**: *"never block a genuinely good bring-up on a soft
heuristic"* (`:280`, `:320`). M255 does **not** retract that, and the retraction it *does* make is narrower
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

> ### 🔴 Step zero on a host with an EMPTY BuildKit cache: run ONE warm-up cycle and DISCARD it
>
> The gated variant is *cold images, **warm** layer cache* (§*The variant*), and v2.8 deliberately **cut** the
> truly-cold run from the gate (`D-v28-8`) because it is *a different, slower thing*. **A brand-new host has no
> layer cache at all** — `odysseus` was probed on 2026-07-31 with **0 images, 0 containers, 0 build cache** — so
> **rep 1 there measures the truly-cold variant the gate excludes.** That is a different *measurement class*,
> not an unlucky rep, and a `p50` computed over it **restates the excluded variant as the gated one**.
>
> **THE RULE: on a host whose BuildKit cache is empty, at least one full warm-up cycle MUST run and be
> DISCARDED before the `n ≥ 3` campaign begins**, so that every counted rep starts from a populated cache. The
> warm-up is a cycle, not a rep: it is not written into the campaign dir and never enters the `p50`.
>
> **This is not the same warm-up as step 3's** — cross-reference, do not conflate. Step 3's exists for the
> **one-off reclaim eviction** (7 records / 356.8 MB → **+173 s**) and protects *comparability between reps of
> one campaign*. This one is about **measuring the right variant at all**. On a fresh host both apply and this
> one is the stronger: step 3's warm-up saves a `p50` from an outlier; this one stops the campaign from grading
> a variant the gate does not accept.

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

> **Invoke it through `python3` — there is no wrapper.** Despite its shebang, `buildbench.py` is committed
> **mode `100644`** (not executable), and the rext clone has no `bin/`, no `pyproject.toml`, and no
> console-script entry point. `./buildbench.py` and a bare `buildbench.py` both fail; only the explicit
> interpreter form runs. (Elsewhere in the corpus `rext <section>/<script>` is shorthand for *"that path inside
> the rext clone"*, not a command on `PATH`.)

```bash
cd <the rext clone>
# --profile names a checked-in hostprofiles/<name>.json; an unknown name exits 2.
# ⚠️ M257x iter-226: `odysseus.json` will NEVER land — `D-v28-15` retired that host on 2026-07-31.
# Name the profile of the host you are ACTUALLY on. `hostprofiles/` currently holds only
# `billion.json` (8-core x86_64, containerd) and `laptop.json` (10-core/16 GiB M1 Pro) — so on the
# sanctioned dev host there is NO applicable profile yet, and nothing in buildbench compares the
# profile you name to the machine you are on (M257x iter-225). Measure one first; do not borrow.
python3 stack-core/buildbench.py run 1 --reps 3 --profile <your measured host> \
        --public-host <magicdns> --label baseline
python3 stack-core/buildbench.py report stack-core/.buildbench/baseline-<ts>
python3 stack-core/buildbench.py assert-headroom --profile laptop --lanes 2   # exit 1 = plan oversubscribes
python3 stack-core/buildbench.py env-snapshot                                 # every knob + where it is read
python3 stack-core/buildbench.py parse --log <an older cycle log>             # back-fill into the same schema
```

**The full flag surface** (verified against the argparse constructed at `buildbench.py:1464`). The ones above are the
common path; these are the rest, and two of them decide whether a campaign is comparable at all:

| verb | flag | default | what it does |
|---|---|---|---|
| `run` | `--out <dir>` | `<rext>/stack-core/.buildbench/<label\|campaign>-<utc-ts>` | where the campaign dir is written |
| `run` | `--no-reclaim` | off | **skip the between-reps reclaim.** Changes the variant — reps then accumulate cache |
| `run` | `--reclaim-until <dur>` | `24h` | the `docker builder prune --filter until=` window (§*The campaign protocol* step 3) |
| `run` | `--dry-run` | off | plan the campaign without running a cycle |
| `parse` | `--samples <file>` | — | fold a `samples.tsv` into the parsed ledger |
| `parse` | `--build-logs <dir>` | — | fold per-image build logs in |
| `parse` | `--n <int>` | `1` | the stack `N` the log came from (anchor matching is `demo-N`-aware) |
| `parse` | `--no-public-host` | off | the run had no `--public-host` → its serve/registry anchors are *not applicable*, not missing |
| `parse` | `--no-ui` | off | the run had `DEMO_NO_UI=1` → its UI-tier anchors are *not applicable*, not missing |
| `report` / `assert-headroom` / `env-snapshot` | `--json` | off | machine-readable output instead of the printed summary |

**Two env knobs, and the first is a trap now that the gate host has changed:**

| knob | read at | default | note |
|---|---|---|---|
| `BUILDBENCH_PROFILE` | `:1205`, `:1217` | — | supplies the default for `--profile` on **both** `run` and `assert-headroom`. **With it unset, `--profile` defaults to `billion`** — so an `odysseus` campaign must pass `--profile odysseus` explicitly or export `BUILDBENCH_PROFILE=odysseus`, or it will grade itself against the wrong host's profile |
| `BUILDBENCH_LANES` | `:873` | `1` | the lane count `run` asserts headroom for. **Env-only — there is no `--lanes` flag on `run`** (`--lanes` exists on `assert-headroom` alone) |

*(`--public-host` also falls back to `STACK_PUBLIC_HOST` (`:1405`) — the same stack-wide knob the rest of the
demo family reads.)*

> **`--reps 1` exits 0 and is NOT a gateable number.** The harness separates **integrity** (`ok`) from
> **quotability** (`gateable`): a sound sub-`n=3` campaign carries `not_gateable_reason` — *"n=1 is below the
> n>=3 floor — sound, but not a gate number"* — printed by `report`. **The exit code does not distinguish
> them**: a `--reps 1` smoke run reports `ok: True` and exits `0`. That is exactly how a non-gateable figure
> gets quoted as a gate — at `n=2` this very campaign would have reported **~773 s on `billion`** against its
> real **666.29 s on `billion`** p50. **Read `gateable`, not the exit code, before quoting any number.**

> **`gateable` has a SECOND clause — the host identity — and the rule above is why it had to.** M257x
> iter-229 taught `buildbench` to refuse a `--profile` that does not describe the machine under it, and wired
> the refusal into `run_campaign`: a mismatch exits with the reserved `EXIT_HOST_IDENTITY`, and the escape
> hatch `BUILDBENCH_ALLOW_HOST_MISMATCH=1` lets the campaign *measure* while keeping its exit code non-zero.
> **That put the entire disclosure in the exit code — the one channel the paragraph above tells you not to
> read.** `campaign.json` is the file this release quotes from, `print_report`'s stdout is what gets pasted
> into a ledger, and `buildbench.py report <dir>` re-aggregates the same rep ledgers in a *different process*
> that never sees the run's exit code at all. All three read clean. Fixed in the M257x harden pass:
> `build_report` now rolls every rep's `host_identity` up **worst-first** into `report["host_identity"]`,
> `print_report` states the verdict on **every** run (including `MATCH` — a disclosure that appears only when
> it is bad teaches nothing about the runs where it is silent), and:
>
> | rolled-up verdict | meaning | `ok` | `gateable` |
> |---|---|---|---|
> | `match` | every rep demonstrated its profile describes this host | ✅ | ✅ |
> | `mismatch` | **any** rep ran against a profile that does not | ❌ **RED** | ❌ |
> | `unmeasured` | the probe could not read the engine | ✅ | ❌ |
> | `absent` | the ledger predates the check (any pre-iter-229 campaign directory) | ✅ | ❌ |
>
> **`absent` is a verdict, not a default.** Treating a missing field as `match` would have awarded a
> gate-quality green to every campaign directory produced before the check shipped — including the ones the
> baseline numbers in this document came from. Nothing is known to be *wrong* with those runs, which is why
> `absent` is not RED; what they cannot do is *demonstrate* the profile described the host, which is exactly
> what `gateable` asserts. `mismatch` is RED and not merely un-gateable because a p50 aggregated across two
> machines is not a measurement of either. Fenced by
> `stack-core/tests/test_buildbench.py::TestIdentityReachesTheQUOTABLEArtifact` (8 tests), whose control arm
> asserts a matched campaign is **still** gateable — a clause that refuses everything grades nothing.

> **`parse --json` is a DEAD flag.** It is declared (`:1432`) and then never consulted: the `parse` branch ends
> in an unconditional `print(json.dumps(out, indent=2))` (`:1511`), so `parse` **always** emits JSON with or
> without it. Harmless, but do not read its presence as implying a non-JSON `parse` mode exists.

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
`billion:/home/devops/panorama/m255/campaign/`** and is what a comparison **on `billion`** re-derives from; per
`D-v28-14`, a comparison on another host re-derives from **that host's own** campaign dir (`odysseus`'s is owed
by M257). Either way: re-read the ledger with `buildbench report`, don't re-type numbers out of this doc.

### Rung zero applies here too

A remote stack consumes `rosetta-extensions` **at a tag fetched from origin** (the M217 FATAL pin guard). A
`buildbench` that exists only in the authoring copy is **unreachable** to the remote bench host — `odysseus`
since `D-v28-14` — and the failure looks like a missing feature rather than a missing tag. `git push --tags` is
part of shipping the harness. See
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

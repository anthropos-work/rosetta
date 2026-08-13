---
title: "KB Fidelity Audit — M257 first-light build"
date: 2026-07-31
scope: milestone:M257
invoked-by: user
---

## Verdict

**RED** — blocked.

Three blocker-severity classes: (1) the milestone's own declared `iteration_protocol_ref`
(`build-budget.md`) asserts in **five places** that the v2.8 gate is measured on `billion` and that
`666.29 s` is what every reduction target is measured against — both superseded by **D-v28-14**
(2026-07-31), and both are exactly what an M257 developer would read as truth when pricing levers;
(2) **`FIX-M256-autoverify-fapi-libressl` is a true blind area** — zero corpus coverage of any kind,
and `verification.md:204` asserts the *opposite* signal; (3) **`corpus/ops/verification.md`'s
authoritative check-set enumeration stopped at M225** and is three checks short of the gate M257
reads (`autoverify green:true / 0 warnings`).

Mitigating (and the reason this is RED-fixable-in-hours, not RED-structural): the fresh-Linux-VM
provisioning knowledge M257's first tik needs **does exist and is host-generic**
(`tailscale-serve.md:119-131` + `setup_guide.md:110-140`), and `odysseus.json` is a **pure drop-in**
— no code change required.

---

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Build budget / READY / attribution | `corpus/ops/demo/build-budget.md` | `stack-core/buildbench.py` | PAIRED (stale) |
| Host profiles + headroom assert | `build-budget.md:267-316` | `stack-core/hostprofiles/*.json`, `buildbench.py:378-495` | PAIRED (stale) |
| `buildbench` CLI surface | `build-budget.md:413-441` | `buildbench.py:1187-1310` | PAIRED (incomplete) |
| **Authoring a NEW host profile** | — | `buildbench.py:389-402` (runtime-only) | **BLIND-AREA** |
| **`odysseus` the host** | — | — | **BLIND-AREA** |
| Fresh-Linux-VM provisioning | `tailscale-serve.md:113-176`, `setup_guide.md:110-140` | `up-injected.sh` pre-flights | PAIRED — **OK** |
| UI-tier build claims (§8.5 retraction) | `frontend-tier.md` | `up-injected.sh:241-816` | PAIRED (anchors drifted) |
| `autoverify` check set (the gate) | `corpus/ops/verification.md` | `stack-verify/live/autoverify.sh` | PAIRED (stale, under-enumerated) |
| `FIX-M256-demo2-service-self-termination` | `verification.md:603-631` | (unimplemented) | DOC-ONLY (one WRONG claim) |
| `FIX-M256-autoverify-fapi-libressl` | — | `autoverify.sh:259-269` | **BLIND-AREA** |
| **Baseline mirror fence** | — | `stack-core/tests/test_baseline_mirror_fence.py` | **CODE-ONLY (mechanical blocker)** |
| demopatch G1-G7 | `demopatch-spec.md` | `demopatch` | PAIRED — no host-bound claims |
| Knobs/defaults | `demo-up-defaults.md` | parsers | PAIRED — OK |
| idempotency / safety / snapshot-spec | those docs | — | PAIRED — no host-bound claims |

---

## Findings — grouped by audit area

### 1. Host-bound claims now stale or mis-scoped

Distinguishing *"where M255's baseline WAS measured"* (still true) from *"where v2.8 gates WILL be
measured"* (now false):

**NOW FALSE — must be retracted or re-scoped:**

- `[STALE]` `corpus/ops/demo/build-budget.md:139` — *"**Every v2.8 reduction target is measured
  against 666.29 s.**"* → D-v28-14 + `overview.md:7`: billion's 666.29 s **does not transfer**;
  M257 must re-measure on odysseus (n≥3) before any lever is priced. **The single most dangerous
  line in the audit** — it is inside the milestone's own `iteration_protocol_ref`.
- `[STALE]` `build-budget.md:230` — *"Not gated — every v2.8 gate is measured on `billion`."*
- `[STALE]` `build-budget.md:240` — *"`billion` is the gate host because it does not [have other
  jobs]."* The *reasoning* ("a developer workstation is not a bench") is still true and transfers to
  odysseus; the *conclusion* naming billion does not.
- `[STALE]` `build-budget.md:261` — *"This is why every v2.8 gate is measured on `billion`."*
- `[STALE]` `.agentspace/rosetta-extensions/stack-core/hostprofiles/billion.json:4` — `"role": "the
  tailnet demo VM — the host EVERY v2.8 exit gate is measured on"`. The **machine-readable** profile
  asserts the superseded fact; `billion.json:13` repeats it (*"THE gated baseline every v2.8
  reduction target is a percentage of"*).
- `[STALE]` `build-budget.md:150-152` — *"**Re-confirm on the first post-freeze campaign** (`billion`
  is under a user freeze until ~2026-07-29)"*. The freeze expired **and** billion became permanently
  demo-only. **Three timing-derived claims are flagged for re-confirmation on a host that no longer
  exists for that purpose** — an orphaned instruction, not just a date.
- `[STALE]` `build-budget.md:446` (Rung zero) — *"unreachable to `billion`"*. Rule holds; host wrong.

**HISTORICAL AND CORRECTLY SCOPED — leave alone:**

- `[OK]` `build-budget.md:108` — the billion host spec for the M255 baseline.
- `[OK]` `build-budget.md:132-133`, `:439-440` — *"The M255 baseline lives at
  `billion:/home/devops/panorama/m255/campaign/`"*. True. **But `:440-441`'s continuation —
  *"and is what every v2.8 comparison re-derives from"* — is now MIS-SCOPED**: M257 re-derives from
  an odysseus baseline that does not yet exist.
- `[OK]` `build-budget.md:101`, `:141-148`, `:22-24` — provenance and the *"state the environment
  with every number"* rule. The rule is what makes D-v28-14 binding.

**THE HIGH-VALUE UNKNOWN (doc is right; the answer is unknown):**

- `[OK→ACTION]` `build-budget.md:262` — *"**There is no unpack leg** — overlay2, not containerd.
  Lever L9 is a `billion`-only phenomenon."* The doc correctly states that lever value is
  **storage-driver-dependent**. **odysseus's storage driver is unrecorded anywhere.** L1's
  ~200-250 s estimate *folds in L9's 85.7 s unpack leg* (`overview.md:68,78`). If odysseus runs
  overlay2, **L1 loses ~86 s of its estimate and the "thin margin" warning at `overview.md:81-89`
  becomes a miss.** This must be the first thing tik-01 records.

### 2. `odysseus` knowledge — BLIND AREA, with a real mitigation

- `[MISSING]` `corpus/**` — **zero occurrences of `odysseus`.** Only `knowledge/plan/state.md`,
  `knowledge/plan/roadmap.md:301-315`, and the M257 `overview.md`.
- `[OK]` `corpus/ops/demo/tailscale-serve.md:119-131` — **the prereq list the roadmap points at
  (`roadmap.md:321`) IS really there and IS specific enough to follow.** A 6-row table with exact
  install commands: (1) Docker+Compose, (2) **Go 1.25.x** with the literal
  `curl … go1.25.12.linux-amd64.tar.gz | sudo tar -C /usr/local -xz`, (3) **atlas CLI** with its
  literal install line, (4) tailscale operator, (5) ssh-agent (auto-started), (6) snapshot cache.
  Mirrored canonically at `setup_guide.md:110-140`. It is host-generic (framed as *"a fresh Linux
  VM"*), so it transfers to odysseus unchanged. **This is NOT the RED.**
- `[OK]` `tailscale-serve.md:133-152` — the **F2b login-shell trap** (*a false "Go NOT on PATH" from
  `ssh host 'cmd'`*) is documented with the disproof one-liner
  `ssh <host> 'bash -lc "go version"'`. Directly relevant: M257 provisions a **no-Go** host over ssh
  and will hit the same symptom for a genuine reason — it must not conflate the two.
- `[MISSING]` **No corpus doc explains how to AUTHOR a host profile.** `build-budget.md` has no such
  section (headings verified at `:17-452`); it documents profile *consumption* only. The 7 required
  keys, the `schema: "hostprofile/v1"` pin, and which numbers are measured-vs-derived exist **only**
  as runtime validation at `buildbench.py:387-402`. M257's first deliverable is `odysseus.json`.

### 3. Host-profile docs vs code

- `[OK]` `build-budget.md:270` — both `billion.json` and `laptop.json` exist at the cited path.
- `[OK]` `build-budget.md:284` — `lane_heap_measured_peak_mib` **3900** / **4223** matches
  `billion.json:25` / `laptop.json:20` exactly.
- `[OK]` disk floor — `disk_floor_gib` is **7** in both (`billion.json:27`, `laptop.json:23`); the
  **25** is `7 + projected_image_gib 18`, and `build-budget.md:327-328` states it correctly.
  Triple-fenced (`up-injected.sh:316`, `test_buildbench.py:264-268`,
  `demo-stack/tests/test_frontend_build.py:641-646`).
- `[STALE]` `build-budget.md:269` + table `:272-276` — *"**Three clauses**"* → the code has **FOUR**.
  `buildbench.py:445` names it in its own comment: *"`require_measured` is clause **zero**"*
  (implemented `:451-460`, emitting `unmeasured_peak_load1` / `unmeasured_disk_avail_gib`). It fires
  in all three production call sites (`:714`, `:971`, `:1262`) and is mutation-tested. **A reader of
  the doc cannot predict a `unmeasured_disk_avail_gib` failure** — which is exactly the failure a
  brand-new host with an unpopulated sampler will produce first.
- `[WRONG]` `build-budget.md:293-294` — *"`max_parallel_ui_lanes = floor((0.8 × budget − idle) /
  measured-lane-peak)`"* → real formula at `buildbench.py:420-426` is
  `max(1, min(ui_image_count, floor((0.8·budget − idle)/lane)))`. The doc omits **both** the clamp at
  `ui_image_count` (3, pinned by `test_buildbench.py:270-274`) and the **floor at 1**. Billion and
  laptop agree by coincidence (1.153→1, 1.736→1); **a third host need not.** The doc's version
  returns **0** on an undersized host where the code returns **1**, and
  `buildbench.py:415-418`'s emphatic *"**The floor of 1 is not a claim that one lane FITS** … do not
  read a return of 1 as 'headroom verified'"* is **absent from the doc**. Load-bearing: L2's entire
  premise (`overview.md:69`) is *"the headroom assert derives `max_parallel_ui_lanes = 1`"*.
  Mirrored-wrong in `billion.json:34` and `laptop.json:29` too — a three-way omission.
- `[STALE]` `build-budget.md:274` — clause 1 says *peak* `load1`. True under `run`
  (`buildbench.py:969`, `max(load)` at `:334`); under the standalone `assert-headroom` CLI it is
  `os.getloadavg()[0]` at `buildbench.py:1258` — a **single instantaneous sample** that the failure
  message still labels "peak". The doc's own example at `:233` describes a point sample.
- `[WRONG]` `build-budget.md:244-245` — *"the profile's `projected_image_gib` … is **labelled as
  such** in `laptop.json`"* → there is **no `provisional_fields` list, no provenance object, no
  machine-readable marker**. Repo-wide case-insensitive grep for `provisional` returns exactly one
  hit: free text inside `laptop.json:30`'s `notes` blob — and `notes` is **never read by any code**.
  The loader validates `projected_image_gib` **identically** to the five measured fields
  (`buildbench.py:399-402`). Confirms the M255-inherited item verbatim.
- `[WRONG]` `build-budget.md:254` says the billion hiring image is **4.67 GB**; `laptop.json:30,32`
  say **4.84 GB** — and 4.84 is the divisor behind `projected_image_gib: 12`. `up-injected.sh:279`
  hedges *"4.67-4.84 GB"*. **M257 must derive odysseus's `projected_image_gib` from one of these.**
- `[OK — the good news]` **`odysseus.json` is a pure drop-in.** No hardcoded map, no `choices=[…]`,
  no enum. `load_host_profile` (`buildbench.py:378-403`) is a bare `d / f"{name}.json"` join
  (`:381`); the glob at `:385` only builds the error hint and will *add* `odysseus` to it
  automatically. `--profile` is a `default=`, not `choices=` (`:1196`, `:1208`). Unknown name →
  exit 2 (`:1235-1237`, `:1253-1255`).
  **Three caveats:**
  - `[MISSING]` A dropped-in profile gets **zero CI coverage**: `test_buildbench.py:354` and `:425`
    hardcode `for name in ("billion", "laptop")` instead of globbing `HOSTPROFILE_DIR`. A malformed
    `odysseus.json` ships green and fails only at first invocation. Two-line fix.
  - `[MISSING]` `name` is **not** in the 7-key required list (`buildbench.py:389-390`) but
    `headroom_assert` does `profile["name"]` at `:492` → `KeyError` on a profile that omits it.
  - `[STALE]` If odysseus should become the default, `"billion"` is the hardcoded argparse default
    twice (`:1196`, `:1208`); otherwise export `BUILDBENCH_PROFILE=odysseus`.
- `[MISSING] **MECHANICAL BLOCKER**` `stack-core/tests/test_baseline_mirror_fence.py:33` pins
  `_PROFILE = billion.json` as the **single source** for the gated baseline, and `:98-112` fences
  **8 prose sites** (including `build-budget.md`, `CLAUDE.md`, `roadmap.md`, `state.md`, and
  **M257's own `overview.md`**) against it: any line matching `baseline|p50|total cycle` that
  carries a number in the **640-700 s band** without also carrying `666.29` is reported as drift.
  **When M257 writes odysseus's measured baseline into prose, this fence goes red or forces the
  prose to keep quoting billion's number.** The fence is **undocumented in the corpus** (0 hits for
  `mirror fence` / `test_baseline_mirror` across `corpus/`).

### 4. `buildbench` CLI surface

- `[OK]` All 5 documented subcommands (`run`, `report`, `assert-headroom`, `env-snapshot`, `parse`)
  exist with the documented flag spellings. **No false promises.** `build-budget.md:420`'s
  `--lanes 2` → exit 1 example is arithmetically verified.
- `[WRONG] **immediate, first-tik**` `build-budget.md:418-422` (mirrored `stack-core/README.md:32-36`)
  — **the invocation examples are not literally runnable.** `buildbench.py` is committed **`100644`**
  (`git ls-files -s` → `100644 42dbd585…`) despite carrying a shebang; there is **no `rext` wrapper,
  no `bin/`, no `pyproject.toml`, no console_script**, and nothing in the repo executes it. Only
  `python3 stack-core/buildbench.py …` works. No test covers the shell path (every test does
  `import buildbench as bb` / `bb.main([...])`). **M257 drives this harness over ssh on a fresh
  host — this is the first thing that breaks.**
- `[STALE]` `build-budget.md:419` — `buildbench.py report .buildbench/baseline-<ts>`. Actual default
  out-dir is `<REXT>/stack-core/.buildbench/<label>-<ts>` (`buildbench.py:1229-1230`); the relative
  path resolves only from inside `stack-core/`. Also `stack-core/.gitignore` does **not** ignore
  `.buildbench/` — campaign output is not git-ignored.
- `[MISSING]` **13 parser flags + 2 env knobs have no doc row anywhere.** `run`: `--out` (`:1198`),
  `--no-reclaim` (`:1199`), `--reclaim-until` (`:1200` — the doc discusses `until=24h` at `:330-334`
  but never names the flag), `--dry-run` (`:1201`). `parse`: `--samples`, `--build-logs`, `--n`,
  `--no-public-host`, `--no-ui` (`:1217-1221`). `--json` on `report`/`assert-headroom`/`env-snapshot`.
  Env: **`BUILDBENCH_LANES`** (`:864` — the *only* way to change lane count for `run`) and
  **`BUILDBENCH_PROFILE`** (`:1196`, `:1208`) — grep for `BUILDBENCH` across all of `corpus/`
  returns **zero hits**. The `parse` omissions are sharpest: with default `--n 1` and no
  `--no-ui`/`--no-public-host`, a legitimate log is reported as **broken**.
- `[MISSING]` `buildbench.py:1117-1120` computes `gateable = ok and len(ledgers) >= 3` plus
  `not_gateable_reason`. **Neither appears in `build-budget.md`** — a `--reps 1` run exits 0 with
  `ok: True` and is silently not quotable.
- `[WRONG]` `buildbench.py:1224` — `parse --json` is a **dead flag**: the handler (`:1282-1304`)
  never reads `a.json` and unconditionally prints JSON, while every sibling branches on it.
- `[STALE]` `buildbench.py:22` + `stack-core/README.md:22` still assert *"orphans ~11.6 GiB of
  BuildKit cache"* per rep — **refuted by `build-budget.md:348-349`** (*"+1.7 to +2.2 GB"*, calling
  11.6 GiB *"an order of magnitude"* off). M257 sizes its campaign disk runway from this.

### 5. The two inherited M256 fixes

**(a) `FIX-M256-demo2-service-self-termination` — documented, unimplemented, one WRONG claim.**

- `[OK]` `corpus/ops/verification.md:603-631` — a dedicated section carrying every symptom
  (`Exited (0)` `:605`, *"DB too many ping failures"* `:607-608`, *"14 of 16"* `:609`, the 20-empty-row
  render `:612-613`, the not-ENOSPC disclaimer `:617-619`, `docker start` recovery `:620`) and the
  routing line **`FIX-M256-demo2-service-self-termination → M257/M258`** at **`:627`** — the only
  occurrence of that identifier in the corpus. Restated at `playthroughs.md:1216-1222`.
- `[OK]` Code confirms the gap: exhaustive grep of `autoverify.sh` finds **no** `docker ps`, no
  `.State.Running`, no container count. `docker inspect` at `:105`/`:363` is a *presence* gate only —
  it succeeds on a stopped container. The identifier appears nowhere in rext: unimplemented.
- `[WRONG]` `verification.md:623-626` — *"`autoverify` probes HTTP endpoints and DB state … so
  nothing in the cheap-win set fires."* The clause is literally correct but the sentence reads as
  *"autoverify cannot see this"*, **and that is false for autoverify as a whole**:
  `stack-verify/lib/services.sh:43-44` carries `jobsimulation` and `cms` rows, both in the demo
  `--services` scope (`up-injected.sh:2494`); an `Exited (0)` container un-publishes its port →
  `code=000` → `status=down` (`services.sh:130-133`) → `verify.sh` rc≠0 → `autoverify.sh:159` warns →
  `green:false`. **A re-run would have gone red.** The stack stayed green because of the
  **stale-verdict class** (the JSON was written once at bring-up tail and never refreshed — the F-6/
  F-10 hazard the same doc names at `:252-260`, `:293-303`), **not** a missing check. M257 could
  build the wrong fix from this paragraph.
- `[MISSING]` The **genuine** coverage hole is narrower and undocumented: `fake-fapi` and `fake-bapi`
  are compose services (`gen_injected_override.py:8,560`) with **no row** in `services.sh`'s
  13-entry `SERVICES` table (`:35-64`). That is the real "16 containers vs 13 probed" gap; `fake-bapi`
  is covered by nothing at all.

**(b) `FIX-M256-autoverify-fapi-libressl` — TRUE BLIND AREA.**

- `[MISSING]` **Zero corpus coverage.** Exhaustive grep across `corpus/`: `LibreSSL` **0 hits**,
  `HTTP 000` **0 hits**, `handshake failure` **0 hits**, the identifier itself **0 hits** (in
  `corpus/` *and* in rext). No routing line, no symptom record, no breadcrumb.
- `[WRONG]` `verification.md:204` — the only corpus text about check (d) — states the **opposite**
  signal (the check exists because *"verify stayed green"*), with no note that on macOS it produces
  the **reverse** false warning. An operator seeing `⚠ NOBODY CAN LOG IN` on a healthy demo has no
  documented path to the right conclusion.
- `[OK]` Code confirms: `autoverify.sh:259-269` uses bare host `curl -fsSk --max-time 3` with **no**
  `docker exec` fallback, no TLS-version pin, no openssl leg. `-k` suppresses **trust** rejection
  only — it does nothing for a handshake the stack cannot negotiate. Both legs collapse to one
  boolean → `warn` → `green:false` on a working stack. (Inverse, also undocumented: because `-k`
  skips trust, this check **cannot** catch the M213/tailscale cert-chain problem `safety.md:782`
  describes.)

**And the structural finding under both:**

- `[STALE]` `corpus/ops/verification.md` is the authoritative check-set doc **and its enumeration
  stopped at M225.** `autoverify.sh` runs **13 checks**; `verification.md` documents 8. Missing
  entirely: **(a2) frontend-build `buildfail.log`** (`autoverify.sh:214-221`), **(f) academy
  `/library/` catalog** (`:312-347`, documented only as an aside in `frontend-tier.md:160`), **(g)
  studio-desk AI provider key** (`:349-381` — `verification.md` never mentions studio-desk at all).
  Its heading at `:194` still reads *"**The four** cheap-wins"* against an eight-check block; `:96`
  says *"**two** decisive checks"* against up to four; `:201` (*"a demo-patch was REFUSED"*) is
  contradicted by its own `-s`-trap section at `:500-514` (the check is three-state: absent → warn,
  non-empty → warn, empty → pass).
- `[MISSING]` The **demo-only gating** of checks (c)-(g) (`autoverify.sh:243-244`, `:382`) and the
  `DEMO_NO_COCKPIT` / `DEMO_STORIES=0` / `DEMO_NO_UI` / `DEMO_NO_ACADEMY_FILL` / port>65535 skips are
  undocumented. `verification.md:199-204` lists cockpit + fake-FAPI with no demo-only qualifier.
- `[MISSING]` `dev-stack/dev-stack:298` calls autoverify **without `STACK_DIR`** → a dev bring-up
  writes **no `autoverify.json`** and silently skips checks (a)/(a2). Undocumented.
- `[MISSING]` §2 (`verify live`) contributes **exactly one** warning regardless of how many probes
  failed (`autoverify.sh:159`). 12 dead services and 1 dead service both yield `warnings: 1`.
- `[OK]` **`green`/`warnings` semantics are correctly documented.** `autoverify.sh:398` sets
  `green = (warnings == 0)` literally; `build-budget.md:33-34`'s READY definition is correct (though
  `warnings: 0` is redundant with `green: true` — **M257's gate is one condition, not two**). The
  real risk is the **denominator**, not the semantics.
- `[STALE]` `build-budget.md:430` cites `autoverify.sh:381-385` for the JSON emit; actual **`:396-400`**
  (`:381-385` lands mid-check-(g)). Same range drifted in code at `buildbench.py:17` (`382-386`) and
  `:802` (`376-381`) — all three wrong, by different amounts.
- `[STALE]` `build-budget.md:302` cites advisory pre-flights at `up-injected.sh:279` / `:319`; actual
  non-fatal continues are at **`:268`** (RAM) and **`:341`** (disk), prose at `:276`.

### 6. `frontend-tier.md` §8.5 retraction anchors — **all four drifted**

`overview.md:104-105` enumerates `frontend-tier.md` **×4 sites — `:231`, `:249`, `:262`, `:271`**.
Verified against the file (676 lines) at HEAD:

- `[WRONG]` `:231` → *"Browser-trusted FAPI cert (M31; M213 remote path)"* — not a retracted claim.
- `[WRONG]` `:249` → *"offset ports with its own minted pk"* — not a claim.
- `[WRONG]` `:262` → *"Non-fatal — actually true now (v1.10b M49 #7)"* — not a claim.
- `[WRONG]` `:271` → **a blank line.**

**Root cause identified:** those are the **pre-M255 line numbers**. `git show 8fded8e^` (the parent
of M255's own doc commit *"…and four corpus claims the barrier proved wrong"*) has the four claims at
exactly `:231`, `:249`, `:262`, `:271`. **M255's commit shifted them +24 and the overview was never
re-anchored.** The live sites at HEAD:

| claim | overview says | **actual at HEAD** |
|---|---|---|
| *"one ~3-minute, ~3.7 GB cached build per frontend"* | `:231` | **`frontend-tier.md:255`** |
| *"a ~3.7 GB next-web compile"* | `:249` | **`:273`** |
| *"pure memory starvation, not a slow build"* | *(named in prose, **no line cited**)* | **`:274`** |
| *"the ~3.7 GB next-web build spike"* (16 GB host) | `:262` | **`:286`** |
| *"the ~3.7 GB build cache"* | `:271` | **ALREADY RETRACTED by M255** at `:299-306` |

So the enumerated set is wrong in three ways: **the line numbers are all stale**, **one of the four
(`:271`) is already fixed** and would be a no-op, and **one live claim (`:274`, memory starvation) is
named in the overview's prose but has no line in the enumeration.** The correct live set is
**`:255`, `:273`, `:274`, `:286`** — still ×4, but a different four.

- `[OK]` `corpus/ops/demo/README.md:139` — *"the honest 'one ~3-min cached build per new demo-N'
  residual"*. **Anchor correct.**
- `[OK]` `CLAUDE.md:318` — the frontend-tier index entry carrying the same residual. **Anchor
  correct.**
- `[STALE — not enumerated]` `up-injected.sh:816` **and** `:1251` both log
  *"~3 min / ~3.7 GB first build"*. `overview.md:112` cites only **`up-injected.sh:794`** — which at
  HEAD is a `demopatch` interview-flag branch, **not** a 3.7 GB claim. **Two live code sites, one
  wrong citation.** Also `up-injected.sh:241`, `:265`, `:273`, `:298` carry `~3.7 GB` in comments and
  the operator warning.
- `[OK]` `demo-up-defaults.md:151` correctly carries the *already-retracted* framing (*"the old value
  was reasoned from a '~3.7 GB frontend build / ~3.7 GB build cache' that is stale by an order of
  magnitude"*) — consistent with `overview.md:108-109` excluding it from the set.
- `[OK]` `frontend-tier.md` mentions "hiring" **4 times** (not zero). `overview.md:111` says *"mentions
  'hiring' zero times in 623 lines"* — the file is now **676 lines** and the hits are in the M255
  disk-headroom section (`:294`) and elsewhere. The *substantive* point (hiring's build is
  undocumented as a UI-tier frontend) still stands, but the "zero times / 623 lines" figure is stale.

---

## Blind areas (these drive RED)

1. **`odysseus` has no corpus presence at all** — 0 hits across `corpus/`. Everything M257 measures,
   asserts and reports is on a host the knowledge base does not know exists.
2. **No doc explains how to AUTHOR a host profile.** `build-budget.md` documents consumption only.
   The 7 required keys / `schema` pin / measured-vs-derived split live solely as runtime validation
   in `buildbench.py:387-402`. M257's first deliverable is `odysseus.json`.
3. **`FIX-M256-autoverify-fapi-libressl` is undocumented in every form**, and `verification.md:204`
   asserts the opposite signal. M257 must fix a defect with no written contract.
4. **The baseline mirror fence is undocumented** (`test_baseline_mirror_fence.py`, 0 corpus hits)
   while hard-pinning `billion.json` as the single source across 8 prose sites — including M257's own
   `overview.md`.
5. **odysseus's storage driver is unrecorded**, and `build-budget.md:262` says lever L9 (~86 s, folded
   into L1) is containerd-only. L1's price is unknown until this is measured.
6. **`verification.md`'s check-set enumeration is 3 checks short** of the gate M257 reads.

---

## What the bootstrap strategy must account for

- **Re-scope every "billion is the gate host" claim before iter-02 prices a lever.** Six sites:
  `build-budget.md:139`, `:230`, `:240`, `:261`, `:440-441`, `:446`, plus `billion.json:4` and `:13`.
  Split each into *historical* (keep) vs *forward-looking* (retract). `:150-152`'s re-confirm
  instruction is orphaned and needs re-homing to odysseus or explicit closure.
- **Sequence `odysseus.json` against the mirror fence, not after it.** The profile is a pure drop-in
  (no code change), but `test_baseline_mirror_fence.py:33` hardcodes `billion.json` as *the* gated
  baseline and fences 8 prose sites in the 640-700 s band. Decide up front whether the fence becomes
  gate-host-parameterised or billion's `gated_baseline` is retired. Also add `name` to the required-key
  list (`buildbench.py:389`) and glob `HOSTPROFILE_DIR` in `test_buildbench.py:354`/`:425`, or
  odysseus.json ships with zero CI validation.
- **Fix the harness invocation before the first campaign, not during it.** `buildbench.py` is
  committed `100644` with no wrapper and no PATH entry — the doc's five examples fail on a fresh
  remote host. Cheapest path: `chmod +x` + commit the mode, or document `python3 stack-core/buildbench.py`.
  Export `BUILDBENCH_PROFILE=odysseus` (the argparse default is `"billion"` at `:1196`/`:1208`).
- **Record odysseus's storage driver in tik-01 and re-price L1 against it.** overlay2 → L9's ~86 s
  does not exist → L1 falls toward its ~140 s floor → `overview.md:81-89`'s "no cushion" warning
  becomes a structural miss, which is a **re-scope signal at 420 s**, not a grind.
- **Treat `verification.md` as under-describing the gate, not describing it.** The gate is
  `green == (warnings == 0)` over **13** checks of which 5 are demo-only and several silently skip;
  the doc lists 8 and says "four". Before trusting a green, know which checks were actually in scope.
  And `verification.md:623-626` is wrong about *why* the M256 stack stayed green — it was the
  **stale-verdict class**, not a blind check set; the container-liveness fix should be scoped to the
  real hole (`fake-fapi`/`fake-bapi` have no `services.sh` row).
- **Re-anchor the §8.5 retraction set before writing the grep gate.** The live claims are
  `frontend-tier.md:255`, `:273`, `:274`, `:286` (not `:231/:249/:262/:271`); `:271`'s claim was
  already retracted by M255 at `:299-306`; the code sites are `up-injected.sh:816` **and** `:1251`
  (not `:794`); `README.md:139` and `CLAUDE.md:318` are correct. Since the gate is a **grep** for
  strings, the drift is recoverable — but the enumeration is what a developer will work from.
- **Correct the harness's own stale numbers while you are in there:** `buildbench.py:22` /
  `stack-core/README.md:22` still say *"~11.6 GiB orphaned per rep"* against the doc's measured
  *"+1.7 to +2.2 GB"* — M257 sizes its disk runway from that figure. And the 4.67-vs-4.84 GB conflict
  (`build-budget.md:254` vs `laptop.json:30,32`) is the divisor for odysseus's `projected_image_gib`.

---

## Applied Fixes

None. This audit was run report-only; no corpus or rext file was edited.

## Gate Result

**RED — blocked.** Resolve blind areas 1-4 and the six host-scope stale claims before entering the
iter loop. Items 5-6 and the CLI/anchor findings are recordable as `KB-N` items and can be absorbed
into the milestone's own `Delivers →` targets (`frontend-tier.md`, `build-budget.md`), which already
cover the retraction and the per-host numbers.

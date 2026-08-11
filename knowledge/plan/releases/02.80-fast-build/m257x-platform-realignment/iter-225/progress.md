**Type:** tik — under `TOK-08`, on the working-stack half of the user redirect.

## Phase 1 — sealed

Predictions `P-225-1..4` sealed before either host profile was opened. Host measured first: `Mac16,11`,
`arm64`, **12 CPU**, **24 GiB**, 208 GiB free on `/`; Docker Desktop **`overlay2`**, VM **8 CPU /
11.67 GiB**.

## Phase 2 — the two profiles

`rosetta-extensions/stack-core/hostprofiles/` holds exactly two, both `measured: 2026-07-27`:

| | `billion.json` | `laptop.json` | **this host** |
|---|---|---|---|
| kind | native-linux | docker-desktop-vm | docker-desktop-vm |
| arch | x86_64 | arm64 (M1 Pro) | **arm64 (`Mac16,11`)** |
| cores | 8 | 10 | **12** |
| mem budget | 7,500 MiB (host) | 9,937 MiB (VM) | **11,948 MiB (VM)** |
| disk budget | 192 GiB | 58 GiB | **125.4 GiB (VM)** |
| storage driver | containerd | overlay2 | overlay2 |

**`P-225-1` HOLDS. Neither profile describes the sanctioned dev host** — `laptop.json` is a 10-core /
16 GiB M1 Pro with a 9,937 MiB VM and a 58 GiB virtual disk; this is a 12-core / 24 GiB `Mac16,11` with an
11,948 MiB VM and a 125.4 GiB virtual disk.

**And both profiles name a host that no longer exists for the purpose.** `billion.json`'s `role` reads
*"NO LONGER THE GATE HOST: **`D-v28-14`** (2026-07-31) moved all dev/test work to `odysseus`… the exit
gates are measured on odysseus"*; `laptop.json`'s notes say the same twice. **`D-v28-15`, the same day,
supersedes `D-v28-14`: `odysseus` is retired and dev/test is LOCAL to the new Mac.** There is **no
`odysseus.json`** — so the gate host the tooling names has no profile either, and the one it should name
has none.

## Phase 3 — `P-225-2` is REFUTED, and what is actually there is worse in one direction

`require_measured` is **not** what a host without a profile hits. Read at
`stack-core/buildbench.py:438-469`, clause zero fails a **`None` measurement input** — a sampler that
died, a disk probe that did not answer — *"a headroom assert that passes because it measured nothing is
worse than no assert."* A profile describing the wrong machine supplies numbers, not `None`, so clause
zero passes cleanly. The prediction named the right contract for the wrong failure.

What is actually there:

- **`--profile` is operator-supplied. There is no host auto-detection at all** (`run_campaign` takes
  `profile_name`; `load_host_profile` resolves it by filename).
- A **missing named** profile is handled exactly right — `FileNotFoundError`, exit-2, with the docstring
  *"a missing profile is exit-2 territory, never a pass"* and the available names listed.
- A **present but inapplicable** profile is **graded silently.** `host_facts()` (hostname, kernel, cores,
  docker version) is collected and written into the run's JSON at `:905` — and **never compared to the
  profile.** So `--profile laptop` on this Mac yields a verdict computed from another machine's
  `cores`, `mem_budget_mib`, `lane_heap_measured_peak_mib`, `disk_floor_gib` and a
  `projected_image_gib` that `laptop.json` itself flags **PROVISIONAL — the only non-measured field**.

The symmetry is the finding: **the harness already refuses an `autoverify.json` verdict *"that does not
describe the run under test"*** (`:809`). The identical discipline is simply not applied to the host
profile — the object that decides whether the run should have been attempted at all.

## Phase 4 — the rest of the sizing

- **`P-225-3` HOLDS.** `build-budget.md` names `billion` **44** times, `odysseus` **26**, `laptop` **22**,
  `M1 Pro` **2** — and the sanctioned dev host **zero**. Its host set is disjoint from `D-v28-15`'s.
- **`P-225-4` HOLDS.** `up-injected.sh`'s `preflight_vm_ram` floor is **12 GiB binary**
  (`DEMO_VM_MIN_GIB:-12`, compared against `docker info MemTotal` in bytes); this VM is **11.67 GiB**, so
  the warning fires on **every** bring-up here — 2.75 % under, well inside the predicted 10 % boundary.
  Non-fatal by design. It is also the unit trap **M257x iter-12 already documented in that file**: Docker
  Desktop's slider is decimal GB, so following the "raise it to 12 GB" instruction literally yields
  ~11.2 GiB and never clears the warning.
- **Disk is NOT the constraint.** The VM has **64.97 GiB free** of 125.4 GiB, against
  `disk_floor_gib + projected_image_gib` = **25 GiB**. The clause that *"actually bites"* on billion is
  slack here.
- **CPU saturation probably is, and there is a precedent in the profile itself.** `laptop.json`'s
  `why_no_full_cycle_number`: a full cycle was **attempted and REFUSED** by clause 1 — `peak load1 10.69
  exceeded cores-2 (8)` — with disk and memory fine, because *"a developer workstation is not a bench; it
  has other jobs."* `D-v28-15` makes a developer workstation the sanctioned dev host.

## Close — 2026-08-09

**Outcome:** the sizing question has a decidable answer — **gate clause 1 is not gradeable on the
sanctioned dev host today**, and the blocker is not the bring-up: **no measured profile exists for the
only machine the release is allowed to test on**, both checked-in profiles describe other hardware and
name a retired gate host, and nothing in `buildbench` would notice if you borrowed one.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-225-1` (a sizing answer is a deliverable; a fabricated cycle is not),
`D-M257x-225-2` (the refuted prediction is published with its mechanism, not quietly restated).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Predictions, graded:**

| id | prediction | result |
|---|---|---|
| `P-225-1` | `laptop.json` does not describe this host; neither profile does | **HELD** |
| `P-225-2` | `buildbench` refuses to grade here, citing `require_measured` | **REFUTED** — clause zero fails `None` *inputs*, not an inapplicable profile; nothing refuses, and nothing compares profile to host |
| `P-225-3` | `build-budget.md`'s host set is disjoint from the sanctioned host | **HELD** — billion 44 / odysseus 26 / laptop 22 / sanctioned host 0 |
| `P-225-4` | the 12 GB VM prereq is within 10 % of this VM | **HELD** — floor is 12 **GiB**, VM is 11.67 GiB, −2.75 %, warning fires |

**Repair landed:** `corpus/ops/demo/build-budget.md` — a superseding banner above the now-stale
`D-v28-14` host-change block, carrying all six measured consequences above. Checked before writing:
**0 corpus-scoped citers of `build-budget.md:NN`** (all 55 hits are frozen `knowledge/plan/` iter records,
which must keep their historical values), so the ~40-line insertion re-points nothing.

**Suite state at close** — `guard_family` with `--platform stack-demo/platform`: **24 GREEN · 0 RED · 5
not-run**, the 5 being commit/ledger-scoped members with no input supplied — **not** a whole-family green,
and the runner says so (`EXIT 2`). `anchor_offset_guard` over the probe range: OK, 0 graded of 59 seen.
No pytest section run — this iter changed no rext code.

**Routes carried forward:**
- `ROUTE-M257x-225-no-profile-for-sanctioned-host` → **this is the actual unblocker for gate clause 1.**
  A `buildbench` measurement run on a **quiet** box producing a checked-in `Mac16,11` profile. Note the
  precedent: the last workstation attempt was refused by clause 1 at load 10.69, so "quiet" is a real
  precondition, not a formality.
- `ROUTE-M257x-225-profile-vs-host-identity-check` → `buildbench` should refuse a profile that does not
  describe the host it is running on, mirroring the `autoverify` verdict guard at `:809`. `host_facts()`
  is already collected; nothing consumes it for this.
- `ROUTE-M257x-225-hostprofile-role-strings-name-a-retired-gate-host` → `billion.json` `role` and
  `laptop.json` notes cite the superseded `D-v28-14` and name `odysseus` as the gate host. A rext edit
  (+ push), deliberately not opened here as a third line.
- `ROUTE-M257x-224-drift-guard-blind-to-stale-clone`, `ROUTE-M257x-222-pin-advance-needs-a-reproof`,
  `ROUTE-M257x-223-classify-the-ten-drifted-baselines` → all still open, unchanged.

**Lessons:**
1. **Name the right contract, or the prediction grades the wrong thing.** `P-225-2` guessed
   `require_measured` because the corpus describes it as *"the clause a fresh host hits FIRST"* — true of
   a host whose **sampler** is unpopulated, false of a host whose **profile** is absent. Two different
   fresh-host failures; the doc's phrasing does not distinguish them, and this iter did not either until
   it read the code.
2. **A guard that refuses a mismatched VERDICT and accepts a mismatched PROFILE has the discipline and is
   pointing it one object too late.** The profile decides whether the run was worth attempting; the
   verdict only reports what happened.
3. **A superseding decision does not propagate itself.** `D-v28-15` superseded `D-v28-14` on the same day
   and left 26 `odysseus` sentences, two profile `role` strings and a whole host-change banner behind it.
   The corpus rule *"a retraction that reaches the prose and not the code has not landed"* has a twin: a
   supersession that reaches the ledger and not the prose has not landed either.
4. **Sizing is a deliverable.** The cheapest disqualifying question — *"is there a profile for this
   host?"* — answered clause 1's attemptability in under an hour without a single container being built.

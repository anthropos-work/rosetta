**Type:** tik · **Active strategy:** `TOK-02` step 2 — *fix the instrument's units before trusting its refusal*

## Line 1 — the defect, measured on the machine it is about

Clause 1 of the gate's HEADROOM assert reads *"peak load1 ≤ cores − 2"*. Three probes, all on this host,
2026-08-11:

| quantity | value | what it describes |
|---|---|---|
| `os.cpu_count()` (the process that reads `os.getloadavg()`) | **12** | the macOS host |
| `sysctl hw.logicalcpu` | **12** | the macOS host — a second, independent probe |
| `docker info` NCPU | **8** | the Docker Desktop **VM** |
| `macmini.json` `cores` | **8** | the VM allocation, per the profile's own `budget_source` |

`Sampler.run` and the `assert-headroom` CLI both call `os.getloadavg()` — **host** load across **12**
cores. Clause 1 graded it against **8**. Limit **6**, where the correct limit is **10**.

**Two independent probes for the host count, deliberately.** iter-04's lesson was that two agreeing *weak*
signals are not one strong signal; `os.cpu_count()` and `sysctl hw.logicalcpu` are not independent in
mechanism, so the honest statement is that the *engine/host split* is what carries the finding — and that
split is confirmed by a probe of a different kind (`docker info` vs the kernel).

**The instrument already knew.** `engine_facts()`'s docstring states the exact numbers — *"`os.cpu_count()`
is 12 and `docker info` NCPU is 8"* — and `profile_describes_host()` grades cores *"against the quantity the
profile's own `kind` declares it to be."* Clause 1 was the third site and it used the wrong one. This is a
slip, not a misunderstanding, which is why the fix is a shared helper rather than a local patch: there are
now **three** consumers and one definition.

## Line 2 — the fix, and the direction it fails

`load1_core_basis(profile, *, observed_host_cores=None)` — the core count of the machine the sample came
**from**. Precedence, strongest evidence first:

1. **`observed_host_cores`** — passed by both live paths from the *same process* that read the loadavg, so
   numerator and denominator provably describe one machine and no declaration can drift out from under it.
2. **`profile["host_logical_cores"]`** — the declared fallback, for grading a load1 recorded elsewhere.
3. **`profile["cores"]` for `native-linux` only** — there the engine shares the host kernel, so the two are
   the same number by construction and `profile_describes_host` already asserts it.
4. Otherwise **`None` → the clause FAILS** with `load1_core_basis`.

**Rung 4 is the whole point.** Falling back to `profile["cores"]` *is* the defect, so the ungradeable case
must be loud. `kind` also became a loader-required key: without it, `cores` could be a host total or a VM
allocation and nothing can tell which.

Supporting edits: `host_logical_cores` declared on both `docker-desktop-vm` profiles, each with its own
provenance — **12, measured** on `macmini`; **10, DECLARED not re-measured** on `laptop`, taken from that
profile's own `budget_source` prose because the machine is retired and cannot be probed again.

## Line 3 — proving it, including proving it can fail

**Live, before and after, on the real profile.** At 20:25:37Z, host load1 **23.48** (top process `Python`
at 792 % CPU — the user's own work; this box is the *permanently contended* one `TOK-02` step 3 is written
for):

```
headroom [macmini] FAIL — lanes=1 max_parallel_ui_lanes=2 free=49.2 GiB load1=23.478 vs 12.0 cores (profile cores=8.0)
  ✗ peak_load1: peak load1 23.48 exceeded cores-2 (10) on a 12-core host … NB the 12 is the core count of
    the machine the load1 sample came from, which on a docker-desktop-vm profile is NOT profile['cores'] (8)
```

That is a **correct refusal under corrected units**, and it is worth reading twice: the fix did **not**
make this host pass. It made the verdict mean something. At the load this box ran during iter-04 (3.84–7.31)
the corrected clause passes and the old one would have refused; at 23.48 both refuse. The fix moves the
threshold to the right machine — it does not disarm the clause, and the CLI now prints both numbers so a
reader can see which machine each belongs to.

**Mutation control — the new tests are RED without the fix.** `load1_core_basis` was monkeypatched back to
`float(profile["cores"])` in memory and the three sensitive tests re-run:

```
MUTATION CONTROL: RED as required   (3 failed)
  test_THE_REGRESSION_…                 -> "peak load1 7.31 exceeded cores-2 (6) on a 8-core host"
  test_a_vm_profile_with_no_host_core_count_is_UNGRADEABLE  -> "8.0 is not None"
  test_a_vm_profile_grades_against_the_HOST_core_count…     -> "8.0 != 12.0"
```

The first failure message is the defect quoted verbatim by the test that pins it: **limit 6 on a host whose
load was 7.31**.

**Eleven new tests**, and the four that matter most:

| test | what it stops |
|---|---|
| `…_is_UNGRADEABLE` | the negative control — basis absent ⇒ the arm goes **RED**, never falls back. A capability probe that fails open disarms the check it guards |
| `test_THE_REGRESSION_…` | the bug itself, on the shipped profile at iter-04's real observed peak (7.31), with a **precondition assert** that the peak still trips the *old* limit — so the test cannot silently stop exercising the defect |
| `test_a_genuinely_saturated_gate_host_still_FAILS` | that the fix did not disarm clause 1 |
| `test_the_live_paths_actually_PASS_the_observed_host_core_count` | **written-but-never-invoked** — asserts the CLI's emitted verdict carries the observed basis *and* that the source has **exactly 2** live call sites passing it |

Also: `test_the_laptop_refusal_on_record_still_refuses` (a historical refusal must not be quietly re-graded)
and a discovered-not-listed sweep asserting every shipped profile can grade its own clause 1.

`python3.12 -m pytest tests/test_buildbench.py -q` → **120 passed** (was 109).

## Line 4 — blast radius, checked rather than assumed

Making `kind` loader-required and adding a field to two profiles could break other consumers. Swept:

- `demo-stack/tests/test_frontend_build.py:687` reads a hostprofile **as JSON**, not through the loader —
  unaffected by a new required key.
- `test_baseline_mirror_fence.py` constructs its own profile bodies and reads them with its own reader —
  **28 passed**, unchanged.
- `stack-core/README.md`'s `hostprofiles/*.json` row listed **two** profiles and no `kind`/`host_logical_cores`
  contract. Updated: three profiles, which one is the gate host, which is retired, and why the two new
  requirements exist. *(A row that enumerates the shipped set is a claim, and it had gone stale the moment
  `macmini.json` landed at iter-04.)*
- Guards that read `README.md`: `test_corpus_index_guard` + `test_fence_registry_population_m257x` +
  `test_claim_twin_guard` → **63 passed**.

## Close — 2026-08-11

**Outcome:** Clause 1 of the gate's HEADROOM assert was grading **the macOS host's `load1` (12 cores)
against the Docker VM's core allocation (8)** — a limit of 6 where the correct limit is 10, failing
**closed**. Fixed at the definition (`load1_core_basis`), consumed by clause 1, observed from the same
process that reads the loadavg at both live call sites, and **proven able to fail**: the ungradeable case
is a new RED clause, and restoring the old basis in memory turns three tests red. `kind` is now
loader-required; both `docker-desktop-vm` profiles declare `host_logical_cores` with provenance.
**Metric delta: none, zero by design** — no lever touched, no cycle run. This is `TOK-02` step 2, whose
whole purpose is that the *next* iter's refusals can be believed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(this is a tik; the streak reset at iter-05)* — (3) re-scope: n *(no p50 exists yet; the trigger reads one)* — (4) user-blocker: n — (5) cap-reached: n *(tik 1 of 5)* — (6) protocol-stop: n — (7) budget-exhausted: n — **Outcome: continue**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D4)
**Side-deliverables:** `stack-core/README.md`'s `hostprofiles/*.json` row — it enumerated **two** profiles
and had been stale since `macmini.json` landed at iter-04. In the fix's blast radius (it now also carries
the `kind` / `host_logical_cores` contract), so it landed here rather than being routed.
**Routes carried forward** (Fate 3, named handlers, → **iter-07** unless stated):
- `BASELINE-M257-macmini-n3` → **iter-07**, now unblocked on both counts: the gate names this host
  (iter-05) and the instrument grades the right machine (this iter). Pre-flight facts gathered here — VM
  disk **51.6 GiB** free against a 22 GiB floor+projected, `demo-1` is a **free, registered, container-less
  slot**, and a demo-up demonstrably works on this box (`demo-2` is live with 11 containers).
- `INVESTIGATE-M257-load1-48` → **narrowed again, and worth re-aiming rather than closing.** This iter
  measured host `load1` **23.48** on a 12-core box — 2× `cores-2` — from the user's own concurrent build.
  odysseus's 48.7 is no longer reproducible, but "a load1 far above core count" is now observed on a
  *second* host, which weakens the odysseus-specific reading and strengthens the standing
  uninterruptible-sleep hypothesis at `buildbench.py:349-350`.
- `MEASURE-M257-macmini-true-idle`, `PROFILE-M257-provisional-fields`, and iter-03/04's remaining routes
  carry unchanged.
**Lessons:**
- **A codebase that states a distinction in one function and forgets it in another has a units bug, not a
  knowledge gap.** `engine_facts()`'s docstring contains the exact two numbers this fix is about. The cure
  is not more documentation — it is **one definition with three consumers**, which is what
  `load1_core_basis` now is.
- **A check that is right by coincidence is indistinguishable from a check that is right.** `laptop.json`
  carries a real, correct clause-1 refusal — on the one host class where the VM allocation and the host
  core count happen to be equal. That refusal was the strongest available evidence the clause was sound,
  and it was evidence of nothing.
- **Grade "does it fail closed or open" before deciding how urgent a units bug is, but fix it either way.**
  This one only ever produced *refusals*, so nothing was silently certified — yet it would have made every
  refusal in the coming campaign unreadable, which is a different kind of expensive.

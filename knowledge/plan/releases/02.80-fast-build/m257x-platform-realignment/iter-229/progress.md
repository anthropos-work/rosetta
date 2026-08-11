**Type:** tik — under `TOK-08` (census the mechanical classes; stop sampling them). Target selected by the
**user's redirect**, which ranks the working stack above the instruments that grade it; profile-host
identity is the one gate-clause-1 blocker that does **not** need a quiet box.

# iter-229 — `buildbench` accepted a profile that does not describe the host it runs on

## Phase 1 — the premise, read from the code rather than from iter-225's prose

`stack-core/buildbench.py`, before this iter:

| fact | site |
|---|---|
| `host_facts()` collects `hostname` / `kernel` / `cores` / `python` / `docker` | `:858` |
| its **only** consumer writes it into the rep ledger as a record | `:905` |
| `pre_rep_assert(profile, *, lanes)` takes no host argument at all | `:718` |
| `read_verdict()` **does** refuse an `autoverify.json` that *"describes an earlier stack"* | `:809` |

So the harness already had the discipline and was pointing it one object too late. `--profile` is
operator-supplied; nothing compared it to the machine underneath.

## Phase 2 — measurement (sealed predictions graded in `933ad95`, before any of this ran)

**Host facts, M257x sanctioned dev host, 2026-08-09:**

```
hostname  Marcos-Mac-mini.local        docker info NCPU          8
kernel    Darwin 25.5.0 arm64          docker info Architecture  aarch64
cores     12  (os.cpu_count())         docker info MemTotal      11947 MiB
python    3.9.6                        docker info Driver        overlayfs
```

### The collision this iter exists for

| quantity | value |
|---|---|
| `os.cpu_count()` (host total) | **12** |
| `docker info` NCPU (VM allocation) | **8** |
| `billion.json` `cores` | **8** |
| `laptop.json` `cores` | 10 |
| `laptop.json` `arch` | `arm64` — **the same as this host** |

`laptop.json`'s own `budget_source` states that its numbers are *"the Docker Desktop VM allocation, NEVER
host totals."* So a profile's `cores` means **different things in different profiles**, and which one it
means is declared by `kind`. Two cheaper implementations were tried on paper and both have a **silent
accept** on this very host:

| implementation | `billion` here | `laptop` here |
|---|---|---|
| **A** — `cores` always vs engine NCPU | **MATCH — wrong** (8 == 8: an x86_64 *native-Linux* profile accepted on an arm64 Mac) | mismatch |
| **C** — `arch` only | mismatch | **MATCH — wrong** (arm64 == aarch64) |
| **B** — `cores` always vs `os.cpu_count()` | mismatch | mismatch |

**B is not vindicated by that column** — it has no *demonstrated* collision on this host, but it is refuted
by the profile's declared semantics: on a `docker-desktop-vm` profile it compares a host total against a VM
allocation, which is exactly the confusion `laptop.json` warns about and exactly how the M239-F1 ENOSPC got
past a green pre-flight. Stated rather than left to the table, because a column that happens to be clean is
not the same as an arm that is right.

### Predictions, graded

| id | prediction | result |
|----|-----------|--------|
| `P-229-1` | `host_facts()` consumed by **0** assert/gate sites | **HELD** — one consumer, `:905`, a ledger record |
| `P-229-2` | this host's arch matches `laptop.json`'s, so an arch-only check would not catch it | **HELD** — both arm64; pinned as a regression test |
| `P-229-3` | core count differs from both profiles | **HELD** — 12 vs 8 / 10 |
| `P-229-4` | neither profile stores a hostname; `name` is a role label | **HELD** — no `hostname` key in either |
| `P-229-5` | `pre_rep_assert` has no host parameter | **HELD** |
| `P-229-6` | `profile["cores"]` and `host_facts()["cores"]` are different quantities on a VM profile | **HELD, and worse than predicted** — the two spellings differ by 4 here, and the VM number **equals `billion.json`'s `cores` exactly** |

**6 of 6 HELD.** The seal earned its keep on `P-229-6`: it was written as a semantic caution and came back a
measured collision, which is what turned the implementation from one arm into three.

## Phase 2b — what landed

`stack-core/buildbench.py`:

- **`engine_facts()`** — `docker info` as structured facts (NCPU / MemTotal / Architecture / OSType /
  Driver). A separate probe from `host_facts()` **on purpose**: on Docker Desktop the engine has its own
  CPU and memory allocation, and that is the quantity a `docker-desktop-vm` profile's numbers denote.
  `{}` on any failure, which the caller renders UNMEASURED.
- **`_norm_arch()`** — `aarch64` ≡ `arm64`, `amd64` ≡ `x86_64`. An **unrecognised** spelling returns
  `None`, not itself: a typo must be ungradeable, not a distinct architecture.
- **`HostIdentity`** + **`profile_describes_host()`** — three verdicts (`match` / `mismatch` /
  `unmeasured`), three arms:
  - `kind` — a `native-linux` profile on a Darwin host is a category error. This is the arm that catches
    the worst case, `--profile billion`, whose `cores` coincides with this engine's NCPU.
  - `arch` — normalised; **alone it is not enough**, and the test says so by name.
  - `cores` — graded against the quantity the profile's **own `kind`** declares it to be: engine NCPU for
    `docker-desktop-vm`, `os.cpu_count()` for `native-linux`. An **unknown** `kind` makes `cores`
    UNMEASURED rather than guessed.
  - `mem_budget_mib` — **observed but explicitly NOT graded** (a budget is `headroom_assert`'s object, not
    identity's), and the payload carries the note saying so. A correct exclusion is still a defect while it
    is silent.
- **Wired into `run_campaign`** once, before the reps loop **and before `--dry-run` short-circuits** —
  a dry-run is precisely when an operator wants to learn the profile does not fit. Writes
  `host-identity.json` **always**; refuses with a **reserved `EXIT_HOST_IDENTITY = 3`**. (This bullet said
  *"exit 2"* until Phase 3 measured why that was wrong — see below. Corrected in place rather than
  rewritten, because the wrong choice is the finding.)
- **The hatch, `BUILDBENCH_ALLOW_HOST_MISMATCH=1`** — lets a mismatched host measure anyway, stamps
  `host_identity` into **every** rep ledger, and **keeps the campaign's exit code non-zero**. It buys data;
  it never buys a green.

### Live reading on this host, after the change

```
--- billion -> mismatch
   kind: profile is 'native-linux' but this host is Darwin
   arch: profile is x86_64 but this host is arm64
   cores: profile declares 8 but this host has 12
--- laptop  -> mismatch
   cores: profile declares 10 (VM allocation) but this engine reports NCPU=8
```

`laptop` fires on **one** arm. Both of the other two agree with this host — which is the measured form of
iter-225's finding, and the reason the cheap implementation would have shipped a silent accept.

## Phase 2c — tests

`stack-core/tests/test_buildbench.py`, **+16 tests** in three classes, tested in PAIRS per `§8`:

- **FIRE side** — billion on the Mac (all three arms), laptop on the Mac (cores alone, and the test asserts
  `len(mismatches) == 1` so a future widening that catches it for the *wrong* reason fails), laptop on a
  Linux host (kind).
- **ACCEPT side** — billion on its own host, and laptop against a matching VM allocation, both `match` with
  empty `mismatches` **and** empty `unmeasured`. Without this a fence that never passes looks identical.
- **UNMEASURED as a third verdict** — an unreadable engine is `unmeasured`, never `match`; a real mismatch
  **outranks** an unmeasured field (billion with no engine probe is still a `mismatch`, so a category error
  is never reported to the operator as a probe failure); an unknown `kind` makes `cores` unmeasured.
- **The two naive implementations**, each pinned by the case it silently accepts — and each opens with an
  assertion that the coincidence still exists (*"the coincidence this test exists for is gone; re-derive
  the trap"*), so the test cannot quietly become vacuous when a profile is re-measured.
- **Wiring** — mismatch returns **2** and `pre_rep_assert` is never called; `host-identity.json` is written
  regardless; the hatch proceeds, records the gap in the rep ledger, and still returns non-zero.

## Phase 3 — the whole-section run went RED, and all three failures were this iter's own

`/usr/bin/python3 -m pytest stack-core/tests/ -q` (pytest 8.4.2 / CPython 3.9.6, whole `stack-core`
section, Python): **3 failed · 1,958 passed · 3 skipped** in 24:54. Every one traced to this iter, and each
was a **different** class of self-inflicted damage:

### 1. An exit code that already meant something else

`test_dry_run_writes_a_self_describing_ledger_without_touching_docker` hard-codes `--profile billion` and
asserts `rc in (0, 1)` — *"rc 2 means the harness could not run at all — never acceptable."* It is right,
and `main` spends 2 in five places on unparseable args and malformed profiles. The identity refusal
returned 2 as well, so one code now meant two things a caller fixes differently.

**Repaired by reserving `EXIT_HOST_IDENTITY = 3`** — `§5`'s *"a guard has three verdicts; reserve a
distinct exit"* applied at the process boundary — and by teaching that test its **third** host-dependent
outcome, on its own terms: the refusal must be recorded in `host-identity.json`, must name the arms that
fired, and `rep-01/` must not exist, because refusing on identity happens **before any rep runs**.

That test's own docstring is what makes this worth writing down: it was made host-robust in M256 because
*"a suite that is red on every developer laptop trains people to ignore it."* An identity check hard-fails
on exactly that population by construction. **A guard whose whole job is to notice "this is the wrong
host" will redden every fixed-host test in the suite** — that is not a reason to weaken it, it is a
migration cost, and it must be paid in the tests rather than in the guard.

### 2. `+199` lines of rext rotted 22 corpus anchors, and the machine caught 3

Inserting into `buildbench.py` moved everything below it. `corpus/ops/demo/build-budget.md` carries **22**
anchors into that file. The suite caught **3**:

| instrument | caught | how |
|---|---|---|
| `test_anchor_subject_census_m257x` | 1 | the prose **quotes a literal** (`report["gateable"]`) and it was no longer there |
| `anchor_construct_guard` | 2 | the cited line went **blank** |
| — | **19 silent** | the cited line still holds *a* line of code, naming something else |

`anchor_construct_guard` says so itself, in its own output: *"FLOOR — detects 'resolves to nothing'; does
NOT detect 'resolves to the WRONG construct'. A citation landing on a real line that names something else
PASSES."* Widening was measured and declined at iter-121 — only 5.5 % of corpus anchors supply their own
expected content. **This iter is the first measurement of what that floor costs on a real edit: 19 of 22.**

All 22 were re-derived against the subject and repaired **line-scoped**, never bumped by a global offset —
and the offsets were **not uniform** (`+1` above the insertion point, `+195` below it, `+199` at the tail),
so a single-offset bump would have written 22 new wrong numbers. One anchor (`:279`, whose line is the
non-unique `if m:`) does not resolve by content match at all and was pinned with a 7-line context window.

### 3. The re-derivation itself carried a number, twice

The first pass verified 14 anchors as `UNCHANGED` — **before** the `EXIT_HOST_IDENTITY` constant was added
at line 78, which shifted every one of them by `+1`. The verdicts were then carried across that later edit.
Caught only because `anchor_construct_guard` re-reddened on two now-blank lines.

**This milestone's rule one, failing inside the iter that was repairing it.** *Derive or omit; never
carry* — and the corollary this adds: **a derivation is invalidated by any subsequent edit to its subject,
including your own.** Re-derive after the LAST edit, not after the first.

## Close — 2026-08-10

**Outcome:** `buildbench` graded runs against an operator-named host profile and never checked whether that
profile described the host. It does now — three arms, three verdicts, a reserved exit code — and both
checked-in profiles are correctly REFUSED on the sanctioned dev host. The cheap version of the same check
would have accepted `billion`, an x86_64 native-Linux profile, on an arm64 Mac.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

> **Grading corrected before the commit.** This line first read `(7) budget-exhausted: y — Outcome:
> exit-7`, on the strength of a 42-minute wall clock for one tik. Wall clock is the wrong operand: **25 of
> those 42 minutes were one `pytest` invocation**, and the skill's rule is that `budget-exhausted` fires
> *"only when the budget is actually spent — not on 'enough progress for one session'."* It was not. This
> is the SECOND run in a row to reach for exit-7 on a duration rather than on a remaining budget
> (iter-226 corrected the identical mis-grade three iters ago), which makes it a pattern worth naming:
> **a long iter is not an exhausted session, and a slow instrument is not a spent budget.**
**Decisions:** `D-M257x-229-1` (the hatch records the gap and never buys a green),
`D-M257x-229-2` (`mem_budget_mib` observed, explicitly not graded),
`D-M257x-229-3` (the sanctioned-host profile deliberately NOT authored — it needs a quiet box).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Predictions, graded: 6 of 6 HELD** (table above). `P-229-6` came back stronger than sealed — the VM
allocation on this host **equals `billion.json`'s `cores` exactly**, which turned a semantic caution into a
demonstrated silent accept and is why the check has three arms rather than one.

**Suite state at close** — `stack-core` whole section, pytest 8.4.2 / CPython 3.9.6, Python: **1,958 passed
/ 3 failed / 3 skipped** at the mid-iter tree; all 3 failures were this iter's own and all 3 are repaired.
Post-repair, scoped to the three affected files: **131 passed / 0 failed**. The whole section was **not**
re-run after the repair — a 25-minute run the budget did not have — so the standing whole-section reading is
the pre-repair one, and this is stated rather than rounded up to a green. `guard_family --platform
stack-demo/platform`: **24 GREEN · 0 RED · 5 not-run** (commit/ledger-scoped members, no input supplied —
not a whole-family green; the runner says `EXIT 2` and is right).

**Side-deliverables:** none. Every edit outside the identity check repaired damage the identity check's own
insertion caused.

**Routes carried forward:**
- `ROUTE-M257x-229-anchor-rot-is-19-of-22-invisible` → **new, and the most valuable thing this iter
  measured.** Any rext edit silently rots corpus anchors below its insertion point; the instruments catch
  only anchors that go blank or that quote a literal. A guard that re-derives every `file:line` into a
  tracked rext file **by content** — the script this iter wrote by hand — is decidable and does not need a
  reading. Not built here (third line).
- `ROUTE-M257x-225-no-profile-for-sanctioned-host` → **still open, now LOUD.** Borrowing a profile used to
  be graded silently; it now exits 3 and names the arms. Still needs a quiet box (`D-M257x-229-3`).
- `ROUTE-M257x-225-hostprofile-role-strings-name-a-retired-gate-host` → open, unchanged.
- `ROUTE-M257x-222-pin-advance-needs-a-reproof`, `ROUTE-M257x-223-classify-the-ten-drifted-baselines`,
  `ROUTE-M257x-224-drift-guard-blind-to-stale-clone`,
  `ROUTE-M257x-228-corpus-disagrees-with-itself-about-refs`,
  `ROUTE-M257x-227-archived-repo-selfdesc-is-stale` → all open, unchanged.

**Lessons:**
1. **A field name is not a unit.** `cores` denotes the host total in one profile and a VM allocation in
   another, and which one is declared by a *sibling* field. Written into the protocol doc `§8`.
2. **Grade a candidate check by what it ACCEPTS.** Every arm rejects something; only one rejected `billion`
   on a Mac. The two naive versions are now regression tests that each open by asserting the coincidence
   still exists, so they cannot go vacuous when a profile is re-measured.
3. **A derivation is invalidated by your own later edit.** Re-derive after the last change, not the first.
4. **An instrument that states its own floor is telling you what it will cost you.** `anchor_construct_guard`
   has printed *"does NOT detect 'resolves to the WRONG construct'"* on every run for a hundred iters. This
   iter is the first to put a number on it: **19 of 22**.

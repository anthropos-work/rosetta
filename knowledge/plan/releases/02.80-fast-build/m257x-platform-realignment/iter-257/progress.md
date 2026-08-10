**Type:** tik — under `TOK-08`, on `ROUTE-M257x-256-workspace-pin-is-not-the-canonical-pin`.

_Opened 2026-08-10 11:09 CEST. PR-1…PR-5 sealed in this iter's first commit (`95dd2e2`), before any
measurement._

## Phase A — the census of readers

Every rext mention of `clones.pin`, by file: `demo-stack/tests/test_tooling.py` (17),
`demo-stack/ensure-clones.sh` (13), `stack-core/clone_pin_guard.py` (4),
`stack-core/tests/test_clone_pin_guard.py` (3), `stack-core/README.md` (1),
`demo-stack/tests/test_aireadiness_snapshot_loadmembers_m254.py` (1).

Partitioned by **which file they mean**:

| reader | file it means | what it does with it |
|---|---|---|
| `ensure-clones.sh:204` | canonical → workspace | seeds the copy, **copy-if-absent** |
| `ensure-clones.sh:234-252` | **workspace copy** | `DEMO_ADVANCE_CLONES=pinned` checks each clone out at it |
| `ensure-clones.sh:377-384`, `:428` | **workspace copy** | the freshness / PIN-DRIFT report |
| `clone_pin_guard:151` | canonical (derived from its own location) | arms A/B/C |

**Two production sites read the copy, both inside one file; zero other files read it, and — before this
iter — zero fences did.**

## Phase B — the decision, and the arm

`ensure-clones.sh:373` names the workspace file an *optional operator pin declaration*, so overwriting
it would destroy a deliberate state the tooling itself calls `pinned`. The copy semantics therefore
**stay**, and the asymmetry is encoded in a new **arm D** on `clone_pin_guard` (`D-M257x-257-1`):
a phantom key is a FINDING, a value difference is a DISCLOSURE, no `stack-*/` is NOT-RUN-with-a-reason.
Wired into `guard_family` via the family's own `--repo-root` (never a new flag — the never-two-clone-sets
rule the four platform-facing guards already follow).

**Watched going RED, then GREEN:**

```
5 finding(s) — [D phantom] the WORKSPACE pin stack-demo/clones.pin.json names 'cms' … 'jobsimulation'
… 'messenger' … 'roadrunner' … 'storage'
arm D read 1 workspace copy/copies; 3 value drift(s) DISCLOSED, not failed:
    ant-academy = 22df69dd8 vs canonical 249430c39
    app         = ad9f3c498 vs canonical 3eaadae68
    next-web-app= 8297c684c vs canonical 19423a1fb
```

## Phase C — the repair, and what the drifts meant

This box's copy was reconciled to the canonical: **5 phantom keys removed, 3 values adopted**, 11 → 6
repos. The three value drifts were **not** an operator declaration — they were the pre-iter-256 seed,
which means `DEMO_ADVANCE_CLONES=pinned` on this box would have checked the clones **back out at the
pre-advance shas while logging `pinned`**. A barrier that disagrees with the barrier it was copied from
is worse than no barrier.

`ensure-clones.sh` also gained a **disclosure** at the seed site: when the copy exists and differs from
the canonical, it logs the divergence, names the one-line fix, and names the fence that grades it. It
changes no behaviour — copy-if-absent is deliberate; the silence was not.

## Phase D — what the iter found about its own predecessor

`D-M257x-257-2`: iter-256's *"the pin changes what a cold bring-up on any box builds"* is **retracted at
all three publishing sites**. `DEMO_ADVANCE_CLONES` defaults to `0` and nothing outside
`ensure-clones.sh` sets it, so a default bring-up applies **no pin**; a fresh box takes each repo's
default-branch tip from `git clone` + `make init`, and an existing workspace builds whatever its clones
are checked out at. The **actions** iter-256 took were right for reasons it did not state.

And `D-M257x-257-3`: inserting the disclosure **moved cited lines in `ensure-clones.sh`**, turning three
guards RED from one cause (`anchor_construct_guard`, `demo_knob_guard`, `repair_postcondition`).
Re-pointed `demo-up-defaults.md`'s two anchors `:212 → :220` and `:467 → :475`. §7 rule 4, applied to
our own repo, with the fences supplying the citer list.

## Close — 2026-08-10

**Outcome:** The fence and the mechanism were reading different files, and only one was watched. Arm D
now grades every `stack-*/clones.pin.json` against the canonical — phantom keys fail, value drift is
disclosed, an absent workspace is NOT-RUN. This box's copy is reconciled (11 → 6 repos), so
`DEMO_ADVANCE_CLONES=pinned` no longer silently undoes iter-256's advance. And iter-256's stated reason
for that advance is retracted in place: a default bring-up applies no pin at all.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: y — Outcome: exit-7

**Pre-registrations — 4 of 5 held.** (Trend: 5/5 → 4/4 → 2/5 → 3/5 → 2/5 → 2/5 → **4/5**.)

| | claim | prediction | measured | verdict |
|---|---|---|---|---|
| PR-1 | `clone_pin_guard` has an arm reading the WORKSPACE copy | REFUTED — zero | **zero**; `:151` derives the canonical, the CLI positional is caller-supplied and nothing routes it to a workspace | **HELD** |
| PR-2 | exactly one rext call site reads the copy | REFUTED — two or more | **two** (`:234` advance, `:377` freshness) — but **both in one file**, and zero other files | **HELD**, with the qualifier stated: two *sites*, one *file* |
| PR-3 | `pinned` is the default, so the phantom checkout is ACTIVE | REFUTED — default `0` | **`0`**, and nothing outside `ensure-clones.sh` sets it | **HELD** — and it is why four releases did not trip over it |
| PR-4 | the pin is the ONLY canonical rext artifact seeded copy-if-absent | REFUTED — at least one more | **it is the only one.** The other copy-if-absent seed (`:114`, `platform/.env`) comes from `stack-dev`, not from a canonical rext artifact | **MISS** |
| PR-5 | `clones.lock.json` agrees with the clones as they are | REFUTED — stale | **stale**: `app` and `next-web-app` disagree with reality and it still records `next-web-app behind: 4` | **HELD** |

**Suite state at close** — Python, `stack-core`, `/usr/bin/python3 -m pytest` (CPython 3.9.6):
`test_clone_pin_guard.py` **29 passed** (23 → 29, +6 for arm D);
`test_frozen_expectation_census_m257x.py` + `test_clone_pin_guard.py` **128 passed**. Guard family
(`--platform`, repo root) **29 GREEN / 0 RED / 5 not-run** after the re-point, with `clone_pin_guard`
reporting *"arm D read 1 workspace copy/copies"*.

**Whole section: 1 failed / 2,164 passed / 3 skipped in 2,936.79 s (48:57)** — and the one failure is
worth more than the run. `test_test_collection_fence.py` caught that this iter had appended
`ArmD_TheWorkspaceCopies` **after** `test_clone_pin_guard.py`'s `if __name__ == "__main__"` guard, so
`python3 test_clone_pin_guard.py` would not collect the six new tests **and would still print `OK`** —
*a test that stops running looks exactly like a test that passed* (iter-254's lesson, arriving from a
new direction). Guard moved to EOF; verified both ways: `pytest` **58 passed** across that file and the
fence that caught it, and direct execution now reports `OK` over the full class list. **No second
whole-section run was taken for a two-line block move within one test file** — the fence that found it
and the file it found it in were both re-run instead.

**Side-deliverables:**
- Three literal ratchets re-pinned with recorded reasons: `DOCSTRING_LITERAL_CEILING` 237 → **238**,
  `TEST_MODULE_LITERAL_CEILING` 636 → **637**, `COMMENT_LITERAL_CEILING` 220 → **222** — the last one
  because **a `#:` re-pin reason is itself a `#` comment, so writing down why a ceiling moved moves this
  ceiling.**
- `clone_pin_guard.py::workspace_pins` graded in `derivation_registry.DECISIONS` (`DECLINE:tree-scan`).
- `D-M257x-257-4`: the noun census read a line-number citation followed by a verb as a measurement of
  that verb. Reworded rather than widening the census — and the first attempt to explain the defect
  **reproduced it inside the explanation**.

**Routes carried forward:**
- `ROUTE-M257x-257-lock-file-is-unfenced` → **new.** `stack-demo/clones.lock.json` is the runtime record
  beside the pin and **nothing grades it**; measured stale in 2 of 5 entries, still reporting
  `next-web-app behind: 4` when the real answer is 0. It also carries only 5 repos (no `ant-academy`).
  Handler: `FIX-M257x-257-fence-the-clone-lock`.
- `ROUTE-M257x-256-mixed-ref-anchors` · `ROUTE-M257x-256-the-advance-is-unproven` → open, and the second
  is now the milestone's critical path under the user's closing condition.
- `ROUTE-M257x-256-workspace-pin-is-not-the-canonical-pin` → **CLOSED** by arm D + the reconciliation.
- All iter-255 and earlier routes → unchanged and open.

**Lessons:**
1. **Fence the file the MECHANISM reads, not the file the doc calls canonical.** Arm D is four lines of
   idea and it had been missing since iter-222 because the guard's own docstring named the canonical
   file so confidently that nobody asked which file `pinned` actually opens.
2. **A right action with a wrong reason is worse than it looks.** iter-256 advanced the pin *and* the
   checkouts, which was correct; its stated reason would have justified advancing only the pin, which
   would have changed nothing a bring-up consumes.
3. **Editing rext moves cited lines exactly like advancing a clone does.** The class §7 rule 4 was
   written for is not about platform repos; it is about any file the corpus cites by line.

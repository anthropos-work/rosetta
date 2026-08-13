# iter-157 — progress

**Type:** tik · **Strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)

## Phase A — the partition census

`SURVEY-M257x-iter150-partition-completeness-elsewhere` asked where else a mechanism declares a partition
and never derives that the parts **cover** the domain. The largest declared partition in the tree is
`repair_postcondition`'s fence registry: every fence declares `FENCE_KIND ∈ {postcondition, standalone}`,
and the module states *"Anything else, or nothing at all, is refused."*

**The reading, and it is a mismatch of exactly two:**

| reading | number |
|---|---|
| modules in `stack-core/` declaring a module-level `FENCE_KIND` | **25** (5 `postcondition` · 20 `standalone`) |
| modules the registry enumerated | **23** |
| modules declaring a kind that nothing read | **2** — `guard_family`, `predicate_enumerator` |

The registry's membership was `GUARD_DIR.glob("*_guard.py")`. Neither of those two filenames ends in
`_guard`, so each fell through a partition with **no third bucket** — classified by neither branch, and in
the function's output **indistinguishable from a module that is not a fence at all**. iter-150's shape
exactly: *treated as the safe case BY OMISSION.*

The family's own report carried the arithmetic in plain sight for as long as both existed —
`5 participating … 18 standalone` against 25 declarations — and nothing subtracted.

## Phase B — two false claims, both about this same line

⚠️ **`D-M257x-157-1` — the code counted FILENAMES while every claim about it was written in terms of the
DECLARATION.** Two sites, both wrong, neither found by review:

- `repair_postcondition.discover_fences.__doc__`: *"Derived from the filesystem on every run, so a fence
  added tomorrow enrols itself or makes this fail loudly naming its own filename. There is deliberately no
  list to maintain."* — true of the claim, false of the code. A fence added tomorrow **not named
  `*_guard.py`** neither enrolled nor failed.
- `guard_family.py:67`: *"`FENCE_KIND` — read **STATICALLY** by `repair_postcondition.py`."* — it was
  not. The declaration sat there, correct in content, read by nothing.

This is the class iter-150 named at a constant (*a comment claiming a derivation nobody performs*) showing
up at a **registry**, and it is the second consecutive iter in which the false statement is the tooling's
own prose about itself rather than a claim about the platform. **A glob is not a derivation** — it is a
hand-written predicate with a wildcard in it, and it fails the same way a hand-written list does, minus the
visibility.

## Phase C — the repair

Membership now follows the **declaration**, not the filename: the loop walks `glob("*.py")` and enrols any
module declaring a legal kind. **The `*_guard.py` REQUIREMENT is unchanged** — a guard-named file that
declares nothing, or declares something illegal, is still refused; that arm already worked and the widening
must not weaken it. One branch may pass in silence — a non-guard-named module declaring nothing — because
most modules in the directory are not fences and must not be forced to say so; it is asserted explicitly so
that widening it later is a deliberate act rather than a drift.

Reading after: **5 participating · 20 standalone = 25**, equal to the declaration count, and the ratchet
verdict is unchanged (`0 site(s) reported`, `OK`). The two newly-enrolled modules both declare
`standalone`, so neither is asked for `postcondition_sites()` and the collect path is untouched — which is
why this widening could land inside the iter rather than routing.

## Phase D — the fence

`stack-core/tests/test_fence_registry_completeness_m257x.py`, **12 tests**, every class above the
`__main__` guard with collection parity asserted in-file.

Arms, **both directions**, because only one was ever checked: every declaring module is enrolled · the
registry invents no member that declares nothing · the two modules that exposed this are enrolled **by
enrolment, not by filename** (so renaming either *into* the glob keeps the arm passing for the right
reason — §5 rule 71) · `guard_family`'s claim about being read statically is asserted as the **property**
it claims rather than as its sentence · and four preservation arms over a temp registry (guard-named +
nothing → refused; guard-named + illegal → refused; **non**-guard-named + illegal → refused, the widening's
own new edge, which the skip branch would otherwise swallow; non-guard-named + nothing → silently skipped).

Controls: **anti-vacuity on the declaration set** (≥20 declaring modules, and both buckets non-empty — a
regex that stopped matching would make both completeness arms compare two empty sets) and a **mutation
control** that an empty directory raises `CouldNotRun` rather than returning two empty lists.

**RED-PROOF against the real pre-fix code taken from `HEAD`** (loaded as a module and pointed at the live
guard dir): `PRE-FIX enrolled: 23 · declared: 25 · missing: ['guard_family', 'predicate_enumerator']`.

⚠️ **Two defects in this iter's own first draft of the fence, both mine, both fixed:** it called a function
by a name I guessed (`fence_registry`; the real one is `discover_fences`) — 11 of 12 tests failed on an
`AttributeError`, which is the cheap failure; and `assertRegex` applies `re.search` **without** `re.M`, so
`^FENCE_KIND` meant *start of file* and the arm read as *"guard_family declares nothing"*. **The second is
the interesting one: it is a false NEGATIVE produced by an anchor that meant something different from what
it looked like** — iter-152's `search()`-vs-`^` defect in mirror image, inside the fence written for the
partition class, one iter after the iter that found the original.

## Gates

- `test_fence_registry_completeness_m257x` — **12/12**
- blast radius (`repair_postcondition` ×2, `guard_family` ×2, `fence_provenance`) — see the close section
- live `repair_postcondition` — `5 participating; 20 standalone; 0 site(s) reported` → **OK**
- **Not re-run, and saying so** (§5 rule 60): the full `stack-core` section, `stack-verify`, `dev-stack`,
  `demo-stack`, `stack-injection`. This iter touched one module and added one test file, both in
  `stack-core`; the scoped run covers `repair_postcondition`'s own tests, its consumer `guard_family`, and
  the provenance module they share.

**No `N` reading taken, so no `N` movement is claimed.** Gate **4 of 5**, unchanged.

## Close — 2026-08-08

**Outcome:** the fence registry selected members by **filename glob** while both claims about it said
**declaration** — **25 modules declare a `FENCE_KIND`, 23 were enumerated**, and the two that were not
(`guard_family`, `predicate_enumerator`) fell through a partition with no third bucket, reporting the gap
as a pass. Membership now follows the declaration (**5 + 20 = 25**), the ratchet verdict is unchanged, and
a 12-test fence covers both directions. RED-proofed against `HEAD`'s code: **23 enrolled of 25 declared**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (§9); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
n — Outcome: continue
**Decisions:** `D-M257x-157-1` … `D-M257x-157-6` (see [`decisions.md`](decisions.md))
**Side-deliverables:** `test_repair_postcondition.py::test_01` **re-pointed, not deleted** — it asserted
`on_disk == registry`, encoding the registry's then-current *spelling*; the registry is now correctly a
superset and equality went RED indistinguishably from a regression. Now the subset direction (its stated
property) plus an anti-vacuity guard; the other direction lives in the new fence where it can be stated in
terms of declarations. **The FIFTH instance of `§5` rule 71 in four iters.**
**Routes carried forward:**
- `SURVEY-M257x-iter150-partition-completeness-elsewhere` — **CLOSED by this iter** on the tree's largest
  declared partition. Other partitions exist (`platform_alignment_guard.ALLOWED_STATES`,
  `claim_census_guard.BARE_FILE_ALLOW`) and are allow-lists rather than partitions; recorded, not swept.
- `SURVEY-M257x-iter156-other-reporting-layers` — **still open, and narrowed**: `autoverify.sh:204` was
  graded and is NOT the defect (`D-M257x-157-5`). The remaining runners are ungraded.
- Unchanged and still queued: `FIX-M257x-iter156-cannot-run-sniff-reads-merged-stream` ·
  `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` (**now with a fifth instance, and the
  five-in-four-iters rate is itself the argument for sizing it a sharper predicate**) ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`
**Lessons:** `§5` gains **rule 73** — *a glob is not a derivation*, plus the partition-with-no-third-bucket
shape, the widen-only-when-the-consumer-is-unchanged rule, the requirement-survives-the-widening rule, and
**anchors are load-bearing in both directions and the API decides which** (`assertRegex` applies
`re.search` **without** `re.M`, which produced a false NEGATIVE inside this iter's own fence).

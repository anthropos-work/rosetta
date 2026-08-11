**Type:** tik — under [`TOK-08`](../decisions.md) *(census the mechanical classes; stop sampling them)*.

# iter-159 — the spelling-pin class is TWO classes, and only one of them has a haystack

## Phase A — the labeled set

`SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` has been the largest standing route by
instance count since iter-154, and it has never been started: iter-155 **sized and rejected** the sweep
with numbers (**2,854** string-literal assertions / **766** expression-shaped across **109 files**) and
recorded the next step as *"a sharper predicate, not the sweep."* Iters 156, 157 and 158 each added an
instance without building one. **Seven confirmed instances in five iters.**

The asset nobody had used: **each of those seven has a commit that repaired it.** That is a labeled set
— ground truth this milestone produced as a by-product — and it converts *"is this predicate sharp?"*
from an argument into a measurement. Assembled as seven `(file, pre-repair commit, post-repair commit)`
triples in `stack-core/labeled_spelling_pins.py`.

## Phase B — the predicate, and what the proof did to it

**The hypothesis: the naive predicate looked at the NEEDLE; the property lives in the HAYSTACK.** A
needle cannot tell you whether a string is a domain value or a copied fragment of source. Three clauses,
each independently RED-proofed:

1. **git-tracked** — a test reading a file its subject *produced* (`open(self.log).read()`, captured
   stdout) is asserting **behaviour**, which is what a fence should do. That file is not in git.
2. **source extension** — a checked-in `.json` golden or `.md` corpus doc is a different class.
3. **not parsed** — a derivation through `ast.parse`/`ast.unparse`/`json.loads` has left raw text for a
   parsed construct. **Without this clause the instrument would flag `§8`'s own prescribed repair** — and
   specifically iter-158's fix, the single most damaging thing a census of this class could do.

**The first proof was misleading, and fixing the measurement is the finding.** Graded at *file* level it
read `DISCRIMINATION 1/4` — "the predicate measures the file, not the pin". That grading was wrong: the
repairs removed **one** pinned assertion from files carrying 9, 51 and 56 legitimate source-text
assertions, so a file-level count barely moves. Re-graded at the right grain — *was the assertion the
repair DELETED among the lines the predicate flagged?* — the same predicate scores **4 of 4**.

| measure | result |
|---|---|
| **RECALL** (fired on the pre-repair form) | **4 / 6** = 67 %, against a pre-registered 50 % floor |
| **DISCRIMINATION** (flagged the exact deleted assertion) | **4 / 4** — L410, L252, L763+766, L360 |
| **BLIND SPOT** (declared in advance, confirmed) | **1 / 1** |

### The finding: this is not one class, it is two, split by WHERE the spelling lives

The two recall misses are not a gap a better haystack predicate closes. Both were investigated at source
and **neither reads a source file at all**:

- **(a) IN THE HAYSTACK** — the test searches checked-in code as raw text. **4 of 7 instances.** This
  instrument censuses them, exactly (4/4).
- **(b) IN THE VALUE** — the haystack is *legitimate behaviour*, but a value the test **supplies or
  expects** is a hand-written literal duplicating something the subject derives. **3 of 7 instances**:
  iter-155's scope fence (asserts over `out.stdout`, but feeds the literal fixture
  `"postgresql redis sentinel backend gotenberg"`), iter-157's `assertEqual(on_disk, registry)`, and
  iter-158's `stderr="Traceback (most recent call)\n"`.

**Sub-signature (b) is precisely what `§5` rule 71's prescribed structural repair addresses** — *derive
the expectation from the same source the code derives from* — and it is **unfenced**. Rule 71 named the
repair for the half no instrument was looking at.

The recall number is reported **as measured, on the pre-registered denominator**. Re-labelling the two
misses "blind" after seeing them would lift it to 4/4 and would be the over-claim `D-M257x-158-3`
refuses.

## Phase C — the census

```
spelling-pin-census: 961 unexempted candidate(s) (0 declared exempt) over 6004 assertions
in 106 test files (34 of which read a git-tracked source file as text)
```

**2,854 → 961**, and — per iter-114's rule — the denominator is now stated: 961 of **6,004** assertions.
Concentration, free from the same run and the routing signal for the sweep:

| section | candidates | | top file | |
|---|---|---|---|---|
| demo-stack | **780 (81 %)** | | `test_cockpit.py` | 235 |
| stack-verify | 86 | | `test_tooling.py` | 171 |
| dev-stack | 56 | | `test_frontend_build.py` | 119 |
| stack-core | 28 | | `test_host_prereqs_m215.py` | 72 |
| stack-injection | 11 | | `test_ant_academy.py` | 63 |

**961 is bounded and attributed, but it is not yet sweepable, and this iter does not claim it is.** Most
are the residual rule 71 explicitly sanctions (*"keep only the structural half it can honestly assert —
this script calls X"*). Grading them is the sweep, pre-declared out of scope in this iter's `overview.md`
and routed forward.

## Phase D — the instrument is fenced, and it caught a live RED-proof of iter-157

**A `FENCE_KIND` declaration was written, and the tree refused it.** `repair_postcondition.py` answered
`COULD NOT RUN — guard(s) declare no legal FENCE_KIND ... spelling_pin_census.py (declares 'census';
legal: ('postcondition', 'standalone'))`. That is **iter-157's registry failing closed on an unknown
kind, in production, unplanned** — a live RED-proof of the previous iter's work.

The declaration was then removed **on its merits, not by widening `LEGAL_KINDS`**: a family member is a
fence **at zero and kept there**, and this census stands at 961. Enrolling it would make the family
permanently RED and train readers to skip it — iter-155's *"a warning nobody can act on"* argument. The
promotion path is written into the module docstring in order.

**Fence: `tests/test_spelling_pin_census_m257x.py`, 19 tests, all green.** Every clause tested in **both**
directions (`§8` guards-in-pairs), plus an anti-vacuity control written against the subject (an empty
scan must exit `2 = CANNOT RUN`, never a silent zero — `§9`) and a mutation control on the labeled-set
proof (a stale path would silently shrink the denominator, so it is asserted readable).

**Mutation-controlled, measured, not asserted:**

| mutation | fence response |
|---|---|
| clauses 1+2 dropped (admit every file read) | **4 failures / 19** |
| clause 3 dropped (stop recognising a parser launder) | **1 failure / 19** |

## Gates

- `tests/test_spelling_pin_census_m257x.py` — **19 passed, 0 failed**; both mutations RED.
- `repair_postcondition.py` — **OK**, registry restored: **5 participating + 20 standalone = 25**, matching
  iter-157's corrected count.
- `test_fence_registry_completeness_m257x` + `test_repair_postcondition` + `test_fence_provenance` —
  **77 passed, 0 failed**.

**NOT re-run, named in full (`§5` rule 60, and its rule-71 corollary that an omission from this list reads
as coverage):** the stack-core suite in full (~20–35 min; this iter **added three new files and modified
none**, `git status` shows three `??` and nothing else), and the suites of **demo-stack, dev-stack,
stack-verify, stack-injection, stack-seeding, stack-snapshot, stack-secrets, alignment, playthroughs and
clerkenstein** — the iter modified no file in any of them.

## Close — 2026-08-08

**Outcome:** the largest standing route moves from **sized-and-rejected** to **enumerated with a proven
instrument** — 2,854 → **961 over a stated 6,004 denominator** — and the proof against the milestone's own
seven repaired instances says the class is **two classes**: 4 of 7 live in the haystack (censused, 4/4
discrimination) and 3 of 7 live in a hand-written **value**, structurally invisible to any haystack
predicate and **unfenced**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**no `N` reading taken, so the metric is
UNMEASURED not unmoved (`§9`); a successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule**) —
(3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted:
n — Outcome: continue
**Decisions:** `D-M257x-159-1` … `D-M257x-159-4` (see [`decisions.md`](decisions.md))
**Side-deliverables:** a live, unplanned RED-proof of iter-157's declaration-driven registry — it refused
an unknown `FENCE_KIND` and named the offending module.
**Routes carried forward:**
- `SURVEY-M257x-iter154-other-fences-may-be-pinned-to-spellings` — **HALF CLOSED, and the halves are now
  named.** Sub-signature (a) has a proven instrument and a bounded population; the survey continues as
  the two items below.
- `SWEEP-M257x-iter159-grade-the-961-haystack-candidates` — **NEW.** Grade and declare-exempt the
  enumerated population, run it to zero, then promote the census into the fence family via the documented
  path. 81 % sits in demo-stack and 235 in one file, so the sweep is attributable rather than diffuse.
- `FIX-M257x-iter159-value-side-subsignature-is-unfenced` — **NEW, and the sharper of the two.** 3 of 7
  confirmed instances pin a hand-written value, not a haystack; rule 71's own prescribed repair targets
  exactly this half and nothing enumerates it.
- Unchanged and still queued: `SURVEY-M257x-iter158-noise-classifier-is-narrow-by-choice` ·
  `SURVEY-M257x-iter156-other-reporting-layers` · `SURVEY-M257x-iter152-half-up-services-are-ungradeable` ·
  `FIX-M257x-iter155-add-injected-rows-vs-the-pinned-liveness-only-decision` ·
  `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` · `-iter143-wrong-head-is-unfenced` ·
  `-iter143-scope-derivation-by-grep` · `-iter143-appending-to-the-protocol-doc-rots-the-ledger` ·
  `-iter144-correction-vs-retraction-unfenced` · `SURVEY-M257x-iter144-orphan-arm-is-the-residual` ·
  `FIX-M257x-iter142-path-arm-window` · `-iter142-value-change-articles` · `-iter142-tier-b-underflag` ·
  `FIX-M257x-iter135-adjudicated-live-defects` · `-iter140-receipts-not-checkable-here` ·
  `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter131-predicate-sets-not-enumerated`
**Lessons:** **grade a census at the grain of its claim.** The same predicate read `1/4` and `4/4` on the
same data; the difference was whether the measurement asked *"did the file's count fall"* or *"was the
repaired assertion among the flagged lines."* A coarse measurement had refuted a correct instrument, and
it would have been recorded as a refutation. And: **when a labeled set says a predicate misses, check
whether the misses are the same class before improving the predicate** — here they were a second
sub-signature, so "improve recall" would have been work against a target that does not exist.

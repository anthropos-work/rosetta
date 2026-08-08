**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

# iter-169 — close the hand-listed-stage class, and fence its population

## Phase 0d — pre-flight, and it is the whole story of this iter

The plan was: migrate the one remaining hand-listed battery, then fence the class. The pre-flight — run the
battery's own baseline **before** touching anything — came back **RED**:

```
FAIL: test_00_the_unmutated_baseline_is_GREEN
AssertionError: stack-core/tests/test_buildbench.py is RED before any mutation:
  ['TestSampler::test_the_stop_event_does_not_shadow_a_Thread_attribute']
```

A **BASELINE RED with no attributable test** is the published signature of
`FIX-M257x-iter111-staged-battery-dependency-is-underived` — and the battery reporting it was that route's
**last open member**. Every prior occurrence had been closed by appending one filename to a stage list. The
cheap move was available, obvious, and **wrong**.

One command separated the two: the same test fails **in place**, outside any staged tree. Staging was never
implicated. The cause is that the assertion reads

```python
self.assertTrue(callable(getattr(s, "_stop", None)))
```

and **CPython 3.14 removed `Thread._stop`** (measured: `hasattr(threading.Thread(), "_stop")` is `False`;
`3.14.6`). The Sampler is correct and never touched `_stop`. `D-M257x-169-1`, booked as `§5` **rule 74** —
*a false RED wearing an open route's signature is a decoy; grade the red before you act on the route.*

## Phase A — the class, censused by property

`§5` rule 73 forbids selecting a population with a glob, so the battery population is derived from what a
battery **is**: a test module binding a module-level MUTANT registry.

| population (7) | stages a file set? | derives it? |
|---|---|---|
| `test_m257x_claim_twin_mutation_battery` | yes | yes (iter-168) |
| `test_m257x_repair_reach_mutation_battery` | yes | yes (iter-168) |
| `test_m257x_repair_postcondition_mutation_battery` | yes | yes (iter-168) |
| `test_m257x_mechanical_fences_mutation_battery` | yes | yes (iter-168) |
| `test_repair_leak_guard_mutation_battery` | yes | yes (iter-166/168) |
| `test_m255_mutation_battery` | yes | **hand-listed → this iter** |
| `test_m220_mutation_battery` | **no** | **exempt, and it now proves it** |

iter-168 recorded the population as **six**; the number was never derived. It is **seven**, and the seventh
is a genuine exemption — m220 mutates one subject into a gitignored sibling *beside* the real tree, so the
staged tree **is** the real tree and there is no dependency set to derive. The exemption is re-established on
every run rather than remembered (`test_03_every_exemption_proves_it_stages_no_file_set`), because an
exemption that is merely asserted is a hand-list with better manners.

## Phase B — the sixth occurrence was already live

Deriving m255's stage set returned one `.py` its hand-list did not carry: **`stack-core/fence_provenance.py`**
— the same module iter-111 added by hand to a different battery. It had produced no symptom because
`demo_knob_guard` imports it inside `main()`, which a suite run never reaches. **A latent registry defect is
not an averted one.** The class was never five past occurrences; it was five past occurrences **and one
standing**.

The widening is a *resolution* change, not a bigger haystack: imports resolve against the **importing file's
own directory** first (what the interpreter does), then the two root-relative conventions. That is what
reaches `stack-injection/platform_topology.py` from `stack-injection/gen_injected_override.py` with the repo
as root — the very import harden pass 1 lost a battery to.

**The overview's escalation condition was cleared by measurement, not argument.** All five already-migrated
batteries derive **identical** sets before and after:

| battery | old | new | identical |
|---|---|---|---|
| claim_twin | 6 | 6 | ✅ |
| repair_reach | 4 | 4 | ✅ |
| repair_postcondition | 6 | 6 | ✅ |
| repair_leak | 5 | 5 | ✅ |
| mechanical_fences | 10 | 10 | ✅ |

m255 goes 12 hand-listed → **13 derived** (8 by import + 5 `extra` shell/data files imports cannot reveal).

## Phase C — the fence, and the narrowing it forced

`stack-core/tests/test_battery_stage.py` — **13 tests**, all green. It fails when a battery stages a file set
without deriving it, when any battery carries a literal path registry at all, or when an exemption starts
copying files. Controls, per the standing `TOK-08` requirement: a synthetic hand-listed battery that **must**
be caught by all three classifiers, a plain test module that **must not** be, a cross-section sibling import,
a stdlib-shadow refusal, and the deliberate over-approximation asserted so nobody "fixes" it.

**The fence went RED on its first run — on a true positive of its shape and a false positive of its intent.**
`MD, AN, VA = "a.py", "b.py", "c.py"` in the mechanical-fences battery parses as an `ast.Tuple` but binds
three scalars; no collection exists at runtime and nothing can fall out of date. The classifier was narrowed
to exclude tuple-unpacking — and because iter-158 caught a narrowing that graded 14 of 14 broken checks
green, the narrowing ships with a control that grades **the same three literals in both syntactic forms** and
requires opposite verdicts (`D-M257x-169-5`).

## Phase D — verification, and what it did NOT cover

| run | result |
|---|---|
| `test_m255_mutation_battery` (full: baseline + every mutant + signature-distinctness) | **3/3 OK**, 11.2 s — was RED at iter open |
| `test_battery_stage` (new) | **13/13 OK** |
| `test_buildbench` + `test_union_apply_guard` + `test_demo_knob_guard` in place | **132/132 OK** |
| `test_m257x_repair_reach_mutation_battery` + `test_repair_leak_guard_mutation_battery` | **11/11 OK**, 109 s |
| `test_m220_mutation_battery` (the exemption) | **12/12 OK**, 107 s |
| `test_test_collection_fence` | **16/16 OK** |
| corpus guards vs the edited protocol doc: markdown-structure · corpus-citation · anchor-construct · derived-value · claim-twin | **5 × rc=0** |

**Not covered, stated rather than implied (`§5` rule 60):** the full `stack-core` suite was not run (~20–35 min,
and rule 51's timing leg is unusable on this host); the `claim_twin`, `repair_postcondition` and
`mechanical_fences` batteries were not re-executed — their derived sets were proven **identical** instead,
which is evidence about the derivation and not about those batteries' mutants; and the other four
`rosetta-extensions` sections were not run at all (`§5` rule 68 — "the whole suite" names its denominator).

## Close — 2026-08-08

**Outcome:** the five-occurrence hand-listed-stage class is **closed at its population, not at its last
member**. The sixth occurrence was found **already live** (`fence_provenance.py`, symptomless because the
import sits in `main()`); the population was censused by property at **7**, not the asserted 6, with the
seventh exempt *and proving it*; and the battery's own BASELINE RED turned out to be a **decoy** — a rotted
`Thread._stop` assertion under CPython 3.14, wearing the open route's exact signature.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this iter closed `closed-fixed`; no no-prog
streak, and no `N`/`P` reading was taken so the metric is UNMEASURED, not unmoved — `§9`) — (3) re-scope: n
— (4) user-blocker: n — (5) cap-reached: n (1 tik this run) — (6) protocol-stop: n — (7) budget-exhausted: n
— Outcome: **continue**
**Decisions:** `D-M257x-169-1` … `D-M257x-169-5` (see [`decisions.md`](decisions.md))

**Why `closed-fixed` and not `closed-fixed-partial`:** the planned scope was the census, the migration and
the population fence, and all three landed. The `Thread._stop` repair was **not** a side discovery — the
battery's baseline had to be green for the migration to be verifiable at all, so it sat on the iter's
critical path.

**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` — **CLOSED**, at the population. Open since
  iter-111, five occurrences, six members.
- `FIX-M257x-iter168-m255-battery-stages-across-sections` — **CLOSED.**
- `FIX-M257x-iter166-stage-derivation-covers-code-not-data` — **still open, and now narrower.** m255's five
  shell/data dependencies ride in `extra`. The residual is unchanged in kind: imports cannot reveal a DATA
  dependency, and `extra` is a hand-list by construction. It is disclosed in `battery_stage.py` and is the
  one hand-list the new fence deliberately permits.
- `SURVEY-M257x-iter169-rotted-assertions-beyond-Thread-_stop` — **NEW.** One assertion pinned to a CPython
  private name was RED at HEAD and unnoticed because it was masked by a battery whose red read as a
  different class. Nothing enumerates assertions pinned to interpreter internals. iter-168's own rule
  applies: *measure the hazard, or "the same problem exists elsewhere" is only a mood.*
- `SURVEY-M257x-iter168-ratchet-input-vs-assertion-scope` — unchanged, still open.
- `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` — **advanced, fourth consecutive iter**
  (`battery_stage` now has a suite and a population fence of its own, alongside `waiver_ledger` and
  `frozen_capture`).
- The standing queue, unchanged.

**Lessons:** **a route predicts a cause; it does not certify one.** The strongest evidence available at this
iter's open — matching failure signature, matching subject, an open route naming that exact file — pointed at
a conclusion that was false, and the class's own repair history made the wrong move the cheapest one. One
command outside the staged tree cost nothing and settled it. **Where a known class and a fresh symptom
coincide, the coincidence is the thing to test first.**

And the smaller one, which is really iter-168's lesson turned on itself: iter-168 reported the population as
six without deriving it, one iter after earning *measure the hazard*. Deriving it took a minute and returned
**seven**. **A count you did not compute is a count you inherited.**

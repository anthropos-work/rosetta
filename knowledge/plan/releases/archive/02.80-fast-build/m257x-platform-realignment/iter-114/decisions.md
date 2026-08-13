# iter-114 — decisions

## D-M257x-114-1 — the denominator is declared, and only a corpus-derived one may carry a ratio

**The measurement that forced it** is `TOK-07` rule 4's, restated because it is easy to file as
bookkeeping and it is not: iter-108 was graded **46/46 = 100 %** by this fence, the arithmetic was
correct, and the same propositions stood false one file away. The denominator was a `raw/` directory —
what *one reading detected* — at a per-pass recall this milestone has measured at **33–83 %**.

**Decided.** `repair_reach_guard` takes exactly one of `--enumeration` (a `predicate_enumerator.py` site
set: corpus-derived, per predicate, enumerated **before** any repair) or `--ledger` (a read's `raw/`).
The report **names which** on its own line, above the number it qualifies. Only
`corpus-derived-per-predicate` may print `reach t/N = P%`.

**Rejected — printing the percentage with a warning next to it.** The percentage is the thing that gets
quoted in a close; the warning is not. `TOK-07` rule 4's words are *"unable to print a reach percentage at
all"*, and a caveat is not an inability. The same asymmetry decided the `--json` shape: **the `reach_pct`
key is OMITTED, not set to null**, so a consumer reading `doc["reach_pct"]` raises rather than formatting
a figure the run was never entitled to. A null with a note beside it still formats.

**Also refused (exit 2), each because it makes the ratio meaningless in a different way:** both inputs at
once (not a state one report can name); neither (not a measurement); an empty enumeration; a **malformed
site**, which is refused rather than dropped — a silently shrinking denominator is the one direction a
reach number flatters itself in, and this fence's own tests already named that hazard.

## D-M257x-114-2 — an UNSETTLED enumeration is not a denominator either

iter-113's ceiling makes a new failure available: an enumeration whose headroom nobody adjudicated is a
**candidate list**, not a population. Grading a repair against one would report reach over a set still
being decided — the same defect one step earlier in the pipeline.

**Decided.** `read_enumeration` refuses (exit 2) an enumeration that reports `seed_recall_failures` (its
own forms could not re-find the sites they came from) or `unsettled_headroom` (its ceiling is open). Both
are keys the enumerator already emits, so this is a **fail-closed read of an existing contract**, not a
new one to keep in sync.

## The controls, and why each exists rather than being obvious

- **Positive control (`test_02_an_enumeration_denominator_DOES_print_a_percentage`).** "It refuses to
  print a percentage" is trivially satisfied by a tool that never prints one. Only the pair measures
  anything. This is §5 rule 2's shape applied to a refusal.
- **Anti-vacuity control (`test_02_the_shipped_iter113_enumeration_is_an_acceptable_denominator`).** It
  loads **the artifact iter-113 actually checked in** and asserts 24 predicates / 71 sites. Without it the
  whole class could be exercising a JSON shape nothing in the repository produces — the failure
  §8's iter-94 rule names: write the anti-vacuity control against the guard's **subject**, not its inputs.
- **The unchanged known-answer fixture.** iter-81's repair vs iter-76's ledger still classifies
  **109 / 35 / 3 / 4 / 1 = 152**, exit 1, with `graphql-wundergraph.md:13` named from both readings. An
  extension that quietly softened the fence would show up here first, and it did not.

## The baseline this iter measured on the way past

`--enumeration iter-113/enumeration.json --range 461b547` reports **`reach 0/71 = 0.0%`**. iter-113
touched `knowledge/` and the protocol doc, never the corpus sites, so 0 is correct — and step 2 now starts
from a **measured** zero rather than from an assumption, which is the difference between a repair that can
prove its reach and one that reports it.

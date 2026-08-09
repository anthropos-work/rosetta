# iter-218 — pre-registered numeric claims, SEALED BEFORE ANY REPAIR

Measured at rosetta `1b5fc29` / rosetta-extensions `e1dadde`. Every reading holds each fence's **own**
scope machinery fixed (`_CLONE_CONTEXT`, `fence_provenance.corpus_sources`) so that **the number-matcher
is the only variable** — iter-209's discipline.

## V1 — the family: SIX number+word constructs, and the split is the finding

| module | construct | reads `**N** noun`? |
|---|---|---|
| `derived_count_guard` | `_LABELLED` | **yes** (`\*{0,2}`) |
| `derived_count_guard` | `_ARROW` | **yes** (`\*{0,2}` both sides) |
| `derived_count_guard` | `_BARE_INT` | **yes** (`^\**…\**$`) |
| `derived_count_guard` | `_N_OF_M` | partially — reads `**28 of 29**`, **blind to `28** of **29`** |
| `guard_family` | `_STATED_COUNT` | **no** |
| `platform_predicate_guard` | `_REPO_COUNT` | **no** |
| `suite_census` | `_DECLARED_GO_COUNT_RE` | **no** |

**Every emphasis-aware construct is in ONE module.** The three that are not are in three different
modules, which is why no reader ever saw them as one class.

## V2 — G2's live scope and its blind spot

Over `fence_provenance.corpus_sources` = **114 documents**, inside G2's own `_CLONE_CONTEXT` window:
**21 live matches · 2 emphasis-blind.**

## V3 — the two, adjudicated BEFORE the repair

1. `corpus/ops/update_guide.md:96` — *"The `make pull` command updates the **4** repos defined in
   `repos.yml`"*. **In scope, and it is exactly the claim G2 exists to grade.** It is TRUE today
   (`repos.yml` lists four), so the defect is a **LATENT false-GREEN, not a live one** — stated in
   iter-213's terms, and stated before the repair rather than after.
2. `corpus/architecture/org-repos.md:71` — *"sources **9** service repos' `//terraform` modules"*.
   **OUT of scope**, and it would be a **FALSE RED**: `9 ≠ 4`. It is not a claim about the clone set —
   the possessive `repos'` heads a longer noun phrase. **The un-emphasised spelling of this sentence
   would ALSO be a false RED today**, so the blind spot is currently hiding a false RED, not only a
   false GREEN.

## V4 — the scope narrowing that does NOT work, refuted before it was tried

Requiring `_CLONE_CONTEXT` on the SAME LINE instead of in a ±4-line window **does not separate them**
(line 71 itself contains *"does not clone it"*) and costs **21 → 14** live matches. Refuted.

## V5 — the rule that does, and its price

Excluding a **possessive** `repos'` / `repos’`:

* costs **0** live matches — 21 before, 21 after;
* leaves the widened matcher adding **exactly 1** claim, `update_guide.md:96`;
* and that claim grades **GREEN** (4 = 4), so **zero false REDs** — iter-209's precondition met, which
  is the whole condition on landing this at all.

## Pre-registered stop condition

**G2's live finding count must not move.** If widening `_REPO_COUNT` turns the corpus RED anywhere, the
widening is REFUSED and routed — the iter closes on the measurement, not on a repair.

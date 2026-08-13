# iter-112 — decisions

## `D-M257x-112-1` — the judgement/mechanism boundary, drawn where it can actually be held

An enumerator that "mechanically finds every instance of a predicate" is not possible as stated, and
pretending otherwise would reproduce this milestone's signature defect one layer up — a check reporting
a state it did not measure. A predicate is a *proposition*; a corpus scan matches *strings*. Something
has to turn one into the other, and that something is judgement.

**So the boundary is drawn explicitly and the judgement half is fenced rather than hidden:**

- **choosing a predicate's search FORM is judgement** — derived from the seed's own text where the
  proposition is carried by a literal, authored where it is not;
- **enumerating that form over the corpus is mechanical and complete** — every `.md` in scope, matched
  against both raw lines and `claim_ledger.normalize_document`'s re-flowed text, because corpus prose
  wraps and five hand sweeps in this milestone missed claims for exactly that reason;
- **seed recall is FAIL-CLOSED** — a form that cannot find the site it was written for is a RED, not a
  warning. This is the control that fires when the enumeration, rather than the corpus, is what is
  broken;
- **an underivable predicate is exit 2, never 0 sites** — *"0 sites"* is the most convincing possible
  way for a check that skipped to read like a check that passed (§5 rule 8).

`FENCE-M257x-iter112` = `stack-core/predicate_enumerator.py`, 18 tests, controls per `TOK-06`'s binding
clause (kept at `TOK-07`) and all of them shown firing against mutated copies.

## `D-M257x-112-2` — the fence caught its OWN derivation on the first real run, twice

Recorded because a fence's first live run is the only time its controls are graded against something
nobody tuned them for.

**Run 1 — 1 refusal, 5 seed-recall REDs.** The derivation padded each seed by one neighbouring line,
reasoning that prose wraps. It does — but the reading books the **range the proposition occupies**, and
padding pulled tokens off *adjacent* propositions: `P04` (a claim about speech models) derived
`studio/tools/pdf2md.py`; `P07` (a claim about a PostHog flag) derived a LiveKit call-site citation.
Both then failed seed recall, and `P17` was refused outright because its literal sat two lines below a
seed the ledger had collapsed from the booked range `37-41` to `:37`.

**Two fixes, and one of them is a rule.** Range seeds are first-class (`path:lo-hi`, recalled if **any**
line in the range is found) — the reading books ranges for a reason and collapsing them is a narrowing.
And the derivation reads the booked range **and not one line more**: *if a range is too narrow that is
the READING's statement to correct, not this fence's to guess around.* Both pinned by mutation controls
that re-inject the padding and show the neighbour's token coming back.

**Run 2 — 36.07×, which is a number about English.** With the padding gone but no cap, an uncapped
derivation reported 1046 sites from 29 seeds. A seed line carries a dozen incidental tokens — `10.0.0`,
`$GH_PAT`, a stray `CURRENT` — and enumerating those measures the corpus's **vocabulary**, not the
predicate. Fixed with specificity ranking (`path:line` citation > slashed path > other literal) and a
cap of 4 forms per seed, both with their reasons recorded in the module.

**Neither run produced a wrong number that shipped.** Both were caught by the fence's own refusals
before a report existed. That is the property the controls were written for, and this is the first
evidence it holds.

## `D-M257x-112-3` — the measured reach: a large share of this residual is PROSE, not citations

The finding that decides how step 2 has to be done, and it was not predictable from the outside.

`derive_forms` reaches a predicate whose proposition is carried by a **literal** — a `file.go:19`
citation, a symbol, a module path, a version. Run against iter-109's 24 predicates it refused two
outright and mis-derived five more, and opening the seed lines shows why:

| predicate | seed text | why derivation cannot reach it |
|---|---|---|
| `P04` | *"GPT-4o Mini TTS, TTS v2 HD, TTS v2"* | the proposition IS the prose; `TTS v2 HD` is not an identifier |
| `P22` | *"Public subnets: Application Load Balancer (ALB), Cosmo Router"* | ditto — `Cosmo Router` is two English words |
| `P06` | *"the split is on the endpoint only, not on the agent name"* | a negation over prose; there is no literal at all |

> **This is the same boundary `claim_twin_guard` already draws** — it matches **quoted verbatim forms**
> from an authored ledger, precisely because a token-derivation cannot reach a prose claim. What is new
> is that it is now **measured for this residual** rather than assumed: 22 of 24 predicates needed an
> authored form, and `P16`/`P24` are left DERIVED deliberately as the control that the derived path
> still runs.

Every authored form is a **substring of its own seed line**, so seed recall is satisfied by
construction and the fence's control still has teeth against the *twins* (it constrains where a form
may come from, not where it may match).

## `D-M257x-112-4` — the first per-predicate multiplier ever reported here, and it is NOT yet trustworthy

**29 detected seeds → 211 enumerated sites → 7.28×**, seed recall 100 %, corpus `2a273ad`+ scope
`corpus/services` + `corpus/architecture` (40 files). Artifacts: `enumeration.txt`, `enumeration.json`.

Read it honestly, because the headline number is the least useful thing on the sheet:

- **12 of 24 predicates read `NO-EXPANSION` (×1.0).** By `TOK-07`'s own guard-rail — *a multiplier near
  1.0 indicts the ENUMERATION, not the predicate* — that is **12 verdicts against this ledger's forms**,
  not 12 unique predicates. An authored single literal is easy to make too narrow.
- **4 read implausibly broad** — `P16` ×48, `P18` ×29, `P22` ×37, `P24` ×19. Those forms are matching
  vocabulary. `Cosmo Router` at 37 sites is not 37 publications of *"the router is in the public
  subnets"*; it is 37 mentions of a deleted component.
- **What IS credible sits in the middle, and it is real**: `P10` ×10 (`studio/gen.py`), `P21` ×6,
  `P09` ×6 (`AIReadinessClient.tsx`), `P15` ×4, `P06` ×3, `P03`/`P07`/`P20` ×2. **Every one of those is
  a site iter-109 did not book** — which is the twin population `D-M257x-109-4` predicted and nothing
  had ever enumerated.

**So the instrument lands and the measurement does not.** The iter closes
**`closed-fixed-partial`** on exactly that line: the enumerator, its controls and the ledger are
delivered; a **trustworthy** enumeration of all 24 is not, and the untrustworthy part is named per
predicate rather than averaged into a headline.

**Routed:** `FIX-M257x-iter112-forms-need-a-second-pass` — per-predicate form review, with the two
failure shapes now named (too-narrow → `NO-EXPANSION`; too-broad → vocabulary), fenced by seed recall
and read against the sites themselves. **Step 2 (the repair) must not run against this ledger as it
stands** — repairing 211 sites of which some are vocabulary would be worse than repairing 46.

## `D-M257x-112-5` — the new fence is NOT added to `guard_family`'s census, deliberately

`census()` globs `*_guard.py` plus an explicit extras list; `predicate_enumerator.py` matches neither,
so the family neither runs it nor complains. That is correct and chosen: the tool is **ledger-scoped**
— it cannot run without a `--ledger` that exists only per-iter — and adding it would put a permanent
`NOT-RUN` row in a family summary that already carries six.

**But it IS covered where coverage matters:** it is added by name to
`test_fence_provenance.py::_guards_declaring_json`, so iter-111's machine-mode contract (parseable
`--json`, tree inside the document, flag actually read) binds it from its first commit. A fence exempt
from the family must not thereby be exempt from the property every fence has to satisfy.

# iter-180 — decisions

## `D-M257x-180-1` — one rationale covering two entries is TWO claims wearing one, and the counts hid it

`derivation_registry.py` declined `guard_family::union` and `guard_family::declaring_modules` with a
single shared sentence: *"each returns the same population `census` does (modulo `CENSUS_EXCLUSIONS`)."*
Re-derived at HEAD `4b60aa2`, independently of iter-177's reading:

| derivation | size | measured relation |
|---|---|---|
| `guard_family::union` | **27** | `census ∪ CENSUS_EXCLUSIONS` — **exactly**. The sentence is TRUE. |
| `guard_family::census` | 26 | — |
| `guard_family::declaring_modules` | **26** | `census Δ {guard_family, repair_postcondition}` — **the sentence is FALSE in both directions**, and adding the exclusions back yields `union` minus `repair_postcondition`, still one member away. |
| `repair_postcondition::discover_fences` | **26** | **identical** to `declaring_modules` — and REGISTERED. |

`CENSUS_EXCLUSIONS = {guard_family}`. **`len(census) == len(declaring_modules)`**, so every count-based
comparison of the two reads green — the same cardinality coincidence iter-177's mutation control
characterises from the other side, now shown to have concealed a *second* defect on the other side of the
repo.

**Decision: state the relation PER ENTRY, never one sentence for two.** The same argument as
`D-M257x-121-2` (enumerated tables over adjectives), reached from the other direction: a rationale that
covers two sites is not economical, it is unfalsifiable at the site where it is wrong.

## `D-M257x-180-2` — the repair is a GRAMMAR, not a rewrite; and the fence was RED-proven before it

Rewriting the sentence fixes today's reading and leaves the class exactly as rottable as it was. So the
class was censused first — **2 of 76** registry entries name a sibling derivation in backticks, derived
rather than listed, and they are precisely the two under review — and then given a machine-gradeable
form:

> `RELATION: <module::attr> == <module::attr> [| <module::attr>]`

resolved live on every run by
`tests/test_frozen_expectation_census_m257x.py::ARationaleThatAssertsASetRelationIsGRADED`, which asserts
**both directions**: every sibling-naming rationale must carry a clause, and every clause must hold.

**The fence was run against the defect before the defect was repaired** — both new arms RED, naming the
two entries and the missing relation. A fence first seen green over a repaired tree is a fence nobody has
watched fail (§9 iter-149).

## `D-M257x-180-3` — the pre-registered escalation did NOT fire, and that is reported rather than assumed

iter-180's `overview.md` sealed the condition before any code: *"if the resolver needs a per-site lookup
table, stop and keep the prose repair only — a fence that is itself a registry is the tax iter-178
declined to pay."* Measured: one generic resolver (`module::attr`; call it if callable, then flatten to a
set of `str`) covers **all five operands** across both clauses, including
`repair_postcondition::discover_fences`, which returns a **pair** of lists. No lookup table, so no new
registry, so no new tax.

The flatten is load-bearing and has its own control: a resolver taking only the first element of that
pair would compare a 6-member set against a 26-member one and report a **real-looking disagreement**
rather than raising — the quiet direction, which is the one that needs the control.

## `D-M257x-180-4` — the open survey is not settled here; it is put on a measurement

`SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` asks why one derivation of this
population is REGISTERED while another is DECLINED. iter-180 measures the fact the survey rests on —
`declaring_modules` and `discover_fences` return **exactly the same set** — and stops there.

**Decision: do not resolve it in this iter.** iter-175's reason still holds and is unchanged by the
measurement: deciding it changes what the frozen-expectation census treats as a *candidate*, which is a
scope change to a running instrument, not a rationale fix. What changes is that the survey no longer rests
on a comment's word. The comment that asserted it has been replaced by a clause the suite grades.

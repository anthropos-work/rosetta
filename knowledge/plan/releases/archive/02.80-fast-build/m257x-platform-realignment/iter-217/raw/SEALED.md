# iter-217 — pre-registered numeric claims, SEALED BEFORE ANY REPAIR

Measured at rosetta `0c9cd79` / rosetta-extensions `3965790`, by
[`probe-variants.py`](probe-variants.py), which imports the shipped
`derivation_registry` and holds **every** piece of its machinery fixed
(`_MEASURED_NOUNS`, `_measurement_units`, `_classify_measurement`, `_unit_line`,
`_CENSUS_SKIP`) so that **the separator is the only variable** — iter-209's
discipline, and the reason iter-209's own hand-written slugger was 16× wrong.

## V1 — the baseline, all three ratchets EXACT

`python3 stack-core/derivation_registry.py --ceilings` at `3965790`:

| ratchet | live | ceiling | mark |
|---|---|---|---|
| `DOCSTRING_LITERAL_CEILING` | 195 | 195 | exact +0 |
| `COMMENT_LITERAL_CEILING` | 159 | 159 | exact +0 |
| `TEST_MODULE_LITERAL_CEILING` | 492 | 492 | exact +0 |

Exit 0. **Any movement measured below is this iter's, not inherited slack.**

## V2 — the CLOSE leg (the routed defect)

Admitting a markdown emphasis run **closing on the operand** — `**292**`, `**292 of 704**` —
yields **39 NEW matches** repo-wide that the live regex sees not at all:
**18 in non-test modules · 21 in test modules · 24 of the 39 classify `standing`.**

## V3 — the MIRROR leg is EMPTY, and that is a measurement

Admitting the emphasis run **opening on the noun** instead — `292 **modules**` — yields
**0** new matches on this tree. The idiom is one-directional here. The leg is shipped anyway,
so the symmetry is **by construction and not by coincidence** (`§5` — *two spellings of one rule
drift*), and it therefore ships with a **staged** control, because a live population of 0 cannot
prove the branch fires (`§9`).

## V4 — BACKTICKS ARE REFUSED, and the refusal is priced

Admitting a backtick to the same position adds **19 further matches** (58 total instead of 39).
Inspected, they are **code-span tails, not measurements** — `` `http://sentinel:8087` matches ``,
`` `anchor_offset_guard.py:321` citations ``, `` `> 0` assertion ``, `` `up-injected.sh:2550` passes ``.
A backtick delimits a **code identifier**; `**` delimits **emphasis on a figure**. The two are
different constructs and only one of them is this census's subject. The refusal is **re-derived on
every run**, never asserted as a constant (`§5` — *a refusal gate is a derivation too*).

## V5 — one of the 39 is a FALSE POSITIVE, named and sized before it is counted

`corpus_citation_guard.py:2` — *"of 387 lines carrying one, only **4** name exactly one corpus doc"* —
where `name` is a **verb**. `names?` is a legitimate `_MEASURED_NOUNS` member, so the ambiguity is in
the vocabulary and not in the separator; the live regex has always had it for the un-emphasised
spelling. Disclosed at **1 of 39** rather than absorbed.

## V6 — the falsifiable prediction

The three ratchets rise by a **partition of the 39 that must reconcile**:

> `Δdocstring + Δcomment + Δtestmod = 39`, with the test-module leg taking all **21** test-module
> hits and the other **18** splitting between docstrings and comments.

**The only admissible reason for a shortfall is row de-duplication** (the censuses key rows by
`file:line::text`, so two identical texts on one line collapse to one row) — and if it fires it must be
**itemised**, never netted out. A partition that does not reconcile and is not itemised is a defect of
this iter, not an artefact.

---
iter: 115
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-07
---

# iter-115 — the repair, over the enumerated set: 24 predicates, 71 sites, every pair-half

**Active strategy reference:** [`TOK-07`](../decisions.md#tok-07-enumerate-the-predicate-not-the-anchor--2026-08-06)
— **step 2, second half.** iter-114 landed the first half (the guard's denominator); this iter is the
work the denominator exists to grade. `TOK-07` rules 2 and 3, quoted because they are the whole shape
of this iter:

> **2.** For every predicate, enumerate its instances mechanically across the whole corpus BEFORE
> repairing any of them. **The class is the unit of work, not the citation.**
>
> **3.** Never fix one side of a pair. A predicate is **closed when every enumerated instance is
> closed, or it is not closed.** A repair that closes one instance and leaves its twin **manufactures a
> self-contradiction**, which is strictly worse than the defect it started from.

**Step 0 — re-survey (done, and it is measured, not asserted).**

- `git diff --stat 461b547 HEAD -- corpus/` is **empty**: the corpus has not moved since the
  enumeration was taken, so `iter-113/enumeration.json`'s line numbers are live coordinates.
- The pre-repair baseline is re-run and reproduces: `--enumeration iter-113/enumeration.json
  --range 461b547` → **`reach 0/71 = 0.0 %`**, `denominator: corpus-derived-per-predicate`.
- All **14** platform clones re-verified at this iter's open against iter-109's ground-truth table:
  `platform 0c91421d · app ad9f3c49 · next-web-app 8297c68 · sentinel f2c4619 · studio-desk 41ee357 ·
  ant-academy 22df69d · cms ca50c81 · jobsimulation 462343b · messenger fa47850 · storage 4ce8ece ·
  roadrunner 87d8d44 · graphql-wundergraph 60c229f · app/studio + cms/studio aeec036a`. **No fetch**
  (§5 rule 41a), read-only.

The target is untouched and still meaningful.

## Cluster / target identified

The whole enumerated set. Not a cluster within it — `TOK-07` rule 3 makes a partial predicate a
**worse** state than an unrepaired one, and the enumeration exists precisely so the repair is not
scoped by what a reading happened to see.

**71 sites · 24 predicates · 22 files.** The distribution is not uniform and the tail is where the
prior passes died: **P21 alone is 22 sites over 9 files**, P10 is 11 over 5, P09 is 6 in one file.

## Hypothesis

Repairing every enumerated instance of each predicate — with the **correction re-derived from platform
source at the pinned trees**, never copied from the ledger (§5 rule 22's own worked lesson,
`platform-alignment.md:1442`) — closes the predicates rather than sampling them, and the reach guard
grades it against the corpus-derived denominator rather than a prior reading's detections.

**Three findings are already queued against this repair and each is a trap with a name:**

1. **P08's pin is off by two.** `ai-readiness.md:52` pins `:496`; the M51 block opens at `:498`.
   **Re-derive it** — a pin copied out of a ledger is the third generation of the same defect, and this
   passage is published as a worked example of a repaired anchor.
2. **P13's superlative has a counter-example inside the corpus** — `external_services.md:495` records
   the router as also once built from a git+url context. The repair must not restate a superlative the
   corpus itself refutes one file away.
3. **P24 is one survivor against ten witnesses.** `sentinel.md:5` asserts a compose block that ten
   other corpus sites record as deleted. The repair is a correction, not a new claim.

**Two enumerated PAIRS exist so that neither can be repaired one-sided**, and both are promotions the
enumerator made that the seed set did not contain:

- **P10 `cms.md:171`** — a within-file twin of the `bash -c` claim at `:287`.
- **P12 `ai_architecture.md:212`** — the same fact at a **different pin** (`:1594-1597` vs
  `:1594-1600`). A one-sided repair here manufactures exactly the `external_services.md:554`/`:565`
  self-contradiction iter-108 created.

## Expected lift

**`P` does not move — no reading is taken.** `TOK-07` reads **last**; §9's refinement applies and its
mandated words are used at close: **no `N` movement is claimed**; `P` is **UNMEASURED, not unmoved**.

What must be true at close:

- `repair_reach_guard --enumeration iter-113/enumeration.json --range <base>..HEAD` reports
  **71/71**, `denominator: corpus-derived-per-predicate`, exit 0;
- every predicate is closed at **every** enumerated instance — the per-predicate roll-up is stated
  site by site, so a partial predicate cannot hide inside an aggregate;
- each correction cites platform source at a stated ref, and the ref is stated **in the sentence that
  depends on it** (§5 rule 22, the three-times-rotted anchor);
- the suite the touched fences own is green, with **the invocation stated** beside every count.

## Phase plan

Multi-step by design and declared here so the scope-creep tripwire counts against the planned shape:
one step per predicate group, **committed incrementally** — a 71-site repair abandoned mid-flight
would strand edits across ~20 files, which is the exact risk that deferred this work twice.

## Escalation conditions

- **User-blocker** only per the protocol's list (test gates RED, unrelated-suite regression).
- **Route forward** if a correction cannot be derived from the pinned trees — record the residual and
  say so; never invent a pin.
- If a correction turns out to be **wrong at the ref** (the iter-22 class), the line that misled is
  itself a finding: hunt it, do not silently substitute.

## Acceptable close-no-lift outcomes

- An enumerated site turns out **not** to publish its predicate on re-reading. That is a finding
  against `FIX-M257x-iter113-adjudication-is-judgement`, recorded with the reason, and the site is
  still touched so the denominator is honoured rather than renegotiated.

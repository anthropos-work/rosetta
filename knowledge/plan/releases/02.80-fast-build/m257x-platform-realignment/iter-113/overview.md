---
iter: 113
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: archived
opened: 2026-08-07
---

# iter-113 — the second pass on the FORMS: a subject ceiling, so "rare" has to be earned

**Active strategy reference:** [`TOK-07: enumerate the predicate, not the anchor`](../decisions.md#tok-07-enumerate-the-predicate-not-the-anchor--2026-08-06)
— **step 1, second pass.** `TOK-07` rule 2's guard-rail is the whole subject of this iter:

> **a multiplier that comes back near 1.0× is evidence the ENUMERATION is not working — not that the
> predicate is rare.**

**Step 0 — re-survey (mandatory, done before targeting).** The enumerator reproduces byte-identically at
rext `436d926` against the iter-112 ledger: **24 predicates · 29 seeds → 211 sites · 7.28× · seed recall
100 % · 12 `NO-EXPANSION`**, in 0.18 s. The routed blocker
`FIX-M257x-iter112-forms-need-a-second-pass` is untouched and still the thing standing in front of
`TOK-07` step 2. No substitution.

**Phase 0d pre-flight (RUN, not skipped).** This iter wires a new ledger through the enumerator pipeline,
so the pipeline was dry-run first against the *existing* ledger: **PASS**, exit 0, output identical to
`iter-112/enumeration.txt` modulo the provenance stamp (which now reads clean rather than `DIRTY`,
because iter-112's configuration is committed). No pre-existing tooling debt blocks this iter.

## Cluster / target identified

iter-112 shipped the mechanism and **refused to bank its own number**, which was right. The number is
untrustworthy in **two opposite directions**, and both are properties of the FORMS:

| shape | predicates | what the form is actually doing |
|---|---|---|
| **too narrow** | 12 read `NO-EXPANSION` (P01 P02 P04 P05 P08 P11 P12 P13 P14 P17 P19 P23) | the form is a verbatim literal lifted off the seed — `markdownManager.go:19`, `"when the two disagree, the name wins"`, `"so external users cannot enter"`. A string that occurs once by construction re-finds its own seed and nothing else. |
| **too broad** | 4 read as vocabulary (P16 ×48, P18 ×29, P22 ×37, P24 ×19) | the form names the predicate's **subject**, not the predicate. `Cosmo Router` ×37 is 37 mentions of a deleted component, most of which assert nothing about the VPC's public subnets. |

**Both are the same defect**: the ledger has **one form tier** and is being asked to do two jobs — *identify
the proposition* and *bound where the proposition could live*. So the iter's deliverable is the **second
tier**, not 24 hand-widened strings.

## Hypothesis

Give every predicate a **SUBJECT form set** (broad: every site that so much as mentions the predicate's
topic) alongside its **PREDICATE form set** (specific: the sites that publish the proposition). Then the
ceiling is computable and the question `TOK-07` forbids answering by assertion becomes mechanical:

- **`headroom` = subject-sites − predicate-sites**, listed by `file:line`.
- **`headroom == 0` → SATURATED.** No site mentions this subject that the predicate form misses. A
  `NO-EXPANSION` predicate that is also SATURATED has a **small class, PROVEN** — there is nothing wider
  to find, because the corpus contains nothing else on the subject.
- **`headroom > 0` → the form's verdict is UNSETTLED**, and the fence NAMES the candidate sites. Each is
  then either a real instance (→ widen the predicate form; the site joins the enumerated set) or a
  subject-mention that does not publish the proposition (→ an explicit `excluded` row **with a reason**).
  **Unadjudicated headroom is RED**, not a pass.

That is the answer to *"how did you distinguish a narrow form from a small class"* — the distinction is
**measured against a ceiling and listed**, never asserted. And the same tier fixes the vocabulary
direction for free: `Cosmo Router` is not a bad form, it is a form **on the wrong tier**.

## Expected lift

The primary metric (`P`) does **not** move this iter — no reading is taken, per `TOK-07`'s
enumerate → repair → **read last** ordering. §9's iter-type refinement applies: this iter's `P` is
**UNMEASURED**, not unmoved, and it must not be reported as a flat reading.

What must move is the **denominator's trustworthiness**, and it is graded, not narrated:

1. every predicate carries both tiers, or the run is **exit 2 UNMEASURED** (a predicate whose ceiling is
   unknown cannot have its `NO-EXPANSION` discriminated, and a check that skipped must not read like one
   that passed — §5 rule 8);
2. **zero unadjudicated headroom** at close;
3. the per-predicate multiplier is re-reported **per predicate**, with every surviving `NO-EXPANSION`
   **named** and carrying its ceiling evidence.

## Phase plan

Planned as a **two-step tooling shape** (declared, so the scope-creep tripwire counts against *this*
shape and not against a single-target tik):

- **Step 1 — ship the mechanism.** Subject tier in `predicate_enumerator.py`: ceiling, headroom,
  adjudication, the refusals. **Binding: a mutation control AND an anti-vacuity control that can actually
  fire.** The anti-vacuity control here has one specific job — catch a subject tier that was **copied from
  the predicate tier rather than authored**, which would make every predicate read SATURATED and every
  small class look proven. That is this fence's vacuity mode and it must be shown firing.
- **Step 2 — use it.** Author the two-tier ledger over the 24 predicates, run it, adjudicate the headroom
  to zero, and report the multiplier per predicate.

## Escalation conditions

- **Route forward** if headroom adjudication exceeds what the iter can close honestly — report the
  measured headroom, close partial, and name the residual per predicate. Do **not** bulk-exclude to reach
  green; a bulk exclusion is tuning-to-green wearing an adjudication's clothes.
- **User-blocker** only for the protocol's own list (test gates RED, unrelated-suite regression).
- **Not** a user-blocker: a predicate that stays `NO-EXPANSION` with a proven ceiling. That is a finding.

## Acceptable close-no-lift outcomes

- The second pass runs and **the 12 stay 12** with SATURATED ceilings — i.e. the forms were fine and the
  residual really is that thin. That is a falsification of this iter's hypothesis and a real answer to the
  user's question; it closes the blocker either way, because the discriminator is what was missing.
- The mechanism refuses its own first ledger (iter-112's did, twice). A refusal caught before a number
  ships is the control working.

## What this iter does NOT do

`TOK-07` step 3 — **the read is not taken here.** Zero platform-repo edits; `stack-demo/**` untouched; no
clone fetched (§5 rule 41a); `rosetta-extensions` stays on `main`; clause 5 is not re-cut, narrowed or
argued.

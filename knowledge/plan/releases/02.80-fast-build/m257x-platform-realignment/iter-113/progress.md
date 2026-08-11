**Type:** tik · `iter_shape: tooling` (two planned steps: ship the mechanism, then use it — declared in
`overview.md` so the scope-creep tripwire counts against *that* shape, per
[`build-mstone-iters` Phase 2](../../../../../.claude/skills) and §9's iter-shape refinement)

**Active strategy:** `TOK-07` step 1, second pass — the routed blocker
`FIX-M257x-iter112-forms-need-a-second-pass`.

---

# iter-113 — the ceiling: `NO-EXPANSION` stops being an assertion

## The question this iter had to answer

iter-112 shipped the enumerator, ran it, and **refused to bank its own number**. Its report said 12 of 24
predicates read `NO-EXPANSION`, and `TOK-07` rule 2 scores that against the FORMS:

> **a multiplier that comes back near 1.0× is evidence the ENUMERATION is not working — not that the
> predicate is rare.**

But the rule cuts both ways and iter-112 could not tell which side it was on, because **"the form is too
narrow" and "the class really is that small" produce the identical number.** Nothing in the report
separated them. That is why step 2 was blocked: repairing 211 sites of which some were vocabulary would
have been worse than repairing 46.

## Step 1 — the mechanism

**Two form tiers per predicate**, and the ceiling falls out of the difference:

| tier | what it finds |
|---|---|
| `forms` | sites that **publish** the proposition |
| `subject_forms` | sites that so much as **mention** the topic |

- `ceiling` = every site mentioning the subject — an upper bound on where the proposition *could* live.
- `headroom` = subject sites − predicate sites, **named by `file:line`, never merely counted**.
- `headroom == 0` → **SATURATED**: nothing wider exists to find. A `NO-EXPANSION` predicate that is also
  saturated has a **small class, proven**.
- `headroom > 0` → **UNSETTLED**, and each candidate is either a missed instance (widen the form) or a
  subject-mention that publishes nothing (an `excluded` row **carrying a reason**). **Unadjudicated
  headroom is RED.**

The same tier fixes the opposite failure for free. iter-112's four vocabulary forms — `Cosmo Router` ×37
for a claim about VPC public subnets — were never bad forms; they were forms **on the wrong tier**.

### The vacuity mode, stated before the controls were written

Copy `forms` into `subject_forms` and every predicate reads SATURATED, every small class reads PROVEN,
and the fence certifies a ledger it never settled. Four controls, each shown firing:

1. **Lexical refusal (exit 2)** — a subject form equal to or containing one of its own predicate forms is
   *narrower*; the tiers are inverted and the ledger is refused at load.
2. **Coverage invariant (RED)** — every document the predicate tier finds the proposition in must be
   reached by the subject tier. **At FILE granularity deliberately**: a line-level test goes RED on
   ordinary wrapped prose, which is how `anchor_offset_guard` false-REDed an ordinary append (§8).
3. **Aggregate anti-vacuity (RED)** — if *no* predicate anywhere in the ledger has a subject tier that
   reaches even one site its predicate tier misses, the tier was copied rather than authored. Aggregate on
   purpose: one saturated predicate is a finding, a wholly saturated ledger is a tell.
4. **Stale-exclusion (RED)** — an `excluded` row that excludes nothing is an adjudication that outlived
   the site it was written for.

Plus the paragraph rule. Corpus prose wraps at ~110 columns, so a subject token routinely lands a line
above the sentence publishing the proposition — **one publication seen twice**. Subject hits sharing a
paragraph with a predicate hit are not headroom. The paragraph is the unit because §8's iter-93 rule is
that *a reader lands on a paragraph and not on a file* — and the twin this whole tok exists to catch, the
`ai`-fold pair at `external_services.md:554` / `:565`, is **eleven lines and several paragraphs apart**,
so it survives the suppression untouched. Both halves are pinned in one test, and the mutation control
widens the paragraph to the whole document and watches the distant twin vanish.

**Tests: 30 (was 18), 30 passed. 64 passed with `test_fence_provenance.py` alongside** (`/usr/bin/python3
-m pytest stack-core/tests/test_predicate_enumerator.py stack-core/tests/test_fence_provenance.py -q`,
82.83 s — the invocation is stated with the count, per iter-111).

## Step 2 — the measurement, and it moved DOWN

**29 seeds → 71 sites, 2.45×, seed recall 100 %, 0 open headroom, exit 0.**

iter-112 reported **7.28×**. **The honest direction was down**, and the reason is the whole finding:

| | iter-112 | iter-113 |
|---|---|---|
| enumerated sites | 211 | **71** |
| multiplier | 7.28× | **2.45×** |
| `NO-EXPANSION` | 12 | 16 |
| …of which **settled against a measured ceiling** | **0** | **16** |
| ceiling examined | *(not computed)* | **368 sites** |
| candidates adjudicated with a reason | *(n/a)* | **254** |

**162 of iter-112's 211 sites were four vocabulary forms** (P16 48, P18 58, P22 37, P24 19). Moved to the
subject tier where they belonged, those four predicates enumerate **5 sites between them**. The 7.28× was
mostly a count of how often the corpus says "Cosmo Router".

**And real expansion appeared exactly where it had been hidden** — the twins `D-M257x-109-4` predicted:

| | iter-112 | iter-113 | what the widened form reached |
|---|---|---|---|
| **P21** | 6 | **22** | the `*_RPC_ADDR` re-point, published at 22 sites |
| **P10** | 10 | **11** | `cms.md:171`'s `exec gen.py` — a **within-file twin** of `:287` |
| **P15** | 4 | **5** | |
| **P02** | 1 | **2** | |
| **P12** | 1 | **2** | `ai_architecture.md:212` pins the same nil-default at **`:1594-1600`** while the seed pins **`:1594-1597`** — repairing one would have manufactured exactly the self-contradiction `TOK-07` rule 3 forbids |

### How a narrow form was distinguished from a small class

**Mechanically, and then by named judgement — and the two halves are reported separately because only one
of them is mechanical.**

- **Mechanical half (complete, checkable):** the ceiling. For each predicate the subject tier enumerates
  every site in `corpus/services` + `corpus/architecture` that mentions the topic — **368 sites**. Any
  site the predicate form misses is surfaced as a named candidate. A predicate cannot be called rare while
  a single candidate is unexamined; the run exits 1.
- **Judgement half (named, not mechanical):** **254 candidates were read and excluded with a reason** —
  each is a site publishing a *different proposition about the same subject*, most often the **corrected**
  one. This is judgement. The fence guarantees the candidate set is complete and that every candidate
  carries a reason; **it does not guarantee the reasons are right**, and this close does not claim it does.

**So the honest verdict per predicate is one of two, and they are not equivalent:**

- **`SMALL-CLASS-PROVEN` — 1 predicate (P20).** Zero headroom: no site in the corpus mentions the subject
  that the predicate form does not already reach. Nothing was judged.
- **`SMALL-CLASS-ADJUDICATED` — 15 predicates.** Headroom existed and every candidate was excluded with a
  reason. These rest on 254 readings of mine.

The distinction is the reason the two verdicts are separate strings in the tool rather than one.

### What the ceiling found on the way past — three items step 2 must not repair blind

1. **`P08`'s pin is off by two.** The predicate says the M51 iter-08/09 block sits at
   `ai-readiness.md:496`. The ceiling surfaced `:498`, which is where the block actually opens. **Step 2
   must RE-DERIVE the anchor, never copy it** (§5 rule 22).
2. **`P13`'s quantifier has a counter-example in the corpus.**
   `external_services.md:495` records that the **router** was also once built from a `git+url` context —
   evidence against *"customerio-sync was the ONLY compose service ever built that way."* Not a
   publication of the predicate, so not an instance; material to what the repair should say.
3. **`P24` is a lone survivor against ten witnesses.** The corpus states the messenger compose block was
   **deleted** at ten sites. `sentinel.md:5`'s surviving presupposition that it still exists is the
   outlier — which is what a settled ceiling is supposed to reveal.

### Two REDs this iter made and fixed inside itself

Recorded as its own, per iter-111's standard, and neither was a hidden failure it exposed:

- **The coverage control fired on real data.** Tightening `P13` and `P24`'s subject forms left both tiers
  blind to the very document their proposition lives in. The repair was the subject form; the invariant
  was not touched.
- **The stale-exclusion control fired on real data.** Promoting `cms.md:171` and `ai_architecture.md:212`
  into the enumerated set left their exclusions behind — three rows excluding nothing. Removed.

Both were caught by controls written before the ledger existed, which is the only reason they are
footnotes rather than a wrong number.

## What this iter did NOT do

**No reading was taken.** `TOK-07`'s order is enumerate → repair → **read last**, and step 2 (repair the
whole predicates against this denominator) has not started. Per §9's iter-type refinement, this iter's `P`
is **UNMEASURED, not unmoved** — it must not be counted toward the 3-no-prog tok trigger. Gate unchanged
at **4 of 5**.

Zero platform-repo edits · `stack-demo/**` untouched · no clone fetched (§5 rule 41a) ·
`rosetta-extensions` on `main`, no tag cut · clause 5 not re-cut, narrowed or argued.

---

## Close — 2026-08-07

**Outcome:** the blocker is closed. `NO-EXPANSION` is no longer an assertion: every one of the 16 flat
predicates is settled against a **measured subject ceiling of 368 sites**, with **254 candidates
adjudicated by named reason** and **2 promoted into the enumerated set** — one of them a within-file twin
and one a same-fact-different-pin twin whose one-sided repair would have manufactured a
self-contradiction. The multiplier moved **7.28× → 2.45×, downward and correctly**: 162 of iter-112's 211
sites were four vocabulary forms sitting on the wrong tier.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** D-M257x-113-1 … D-M257x-113-5
**Side-deliverables:** none
**Routes carried forward:**
- `TOK-07` step 2 (repair whole predicates against the **enumerated** set, `repair_reach_guard` graded on
  the corpus-derived denominator with its provenance printed) → **next iter — now UNBLOCKED**
- `TOK-07` step 3 (the read) → last, unchanged
- `FIX-M257x-iter113-adjudication-is-judgement` → the 15 `SMALL-CLASS-ADJUDICATED` verdicts rest on 254
  readings by one agent; a second pair of eyes on the exclusion reasons is worth a harden pass, and the
  ledger is checked in so it can be audited without re-deriving
- `FIX-M257x-iter111-staged-battery-dependency-is-underived` → open
- `FIX-M257x-iter111-buildbench-parse-json-is-a-noop-flag` → open
- `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` → open, de-ranked
- `DEF-M257x-iter101-briefing-rext-tree` → open
**Lessons:**
- **A verdict of "rare" has to be earned against a ceiling, or it is a sample calling itself a census.**
  The cheap version of this fence would have reported 12 rare predicates and been believed.
- **Report the direction that indicts the previous number.** 7.28× → 2.45× reads like a regression and is
  the opposite: the earlier figure was counting vocabulary.
- **Separate the mechanical half of a verdict from the judged half in the OUTPUT, not in the prose.**
  `SMALL-CLASS-PROVEN` (1) and `SMALL-CLASS-ADJUDICATED` (15) are different strings because they carry
  different warranties, and collapsing them would have let 254 judgement calls wear a measurement's voice.

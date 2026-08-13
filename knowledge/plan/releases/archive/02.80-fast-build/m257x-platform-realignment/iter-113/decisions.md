# iter-113 — decisions

## D-M257x-113-1 — a `NO-EXPANSION` verdict is UNMEASURED without a ceiling, so the ceiling is mandatory

`TOK-07` rule 2 says a multiplier near 1.0× is evidence the enumeration is broken. iter-112 obeyed it and
could go no further, because **"the form is too narrow" and "the class is small" produce the identical
number**. The tool had no way to tell them apart and correctly refused to guess.

**Decided:** every predicate carries a second, broader form tier (`subject_forms`); a predicate with none
is **exit 2 UNMEASURED**, never a verdict. Without a ceiling the `NO-EXPANSION` label cannot be
discriminated, and printing it anyway is a check that skipped wearing the voice of one that passed
(§5 rule 8).

**Rejected:** defaulting `subject_forms` to `forms` when absent. It is the single cheapest way to make
every predicate read SATURATED, and a default that fabricates a green is worse than a refusal.

## D-M257x-113-2 — the paragraph is the unit of publication, and the twin it must not swallow is pinned

The corpus wraps at ~110 columns. A subject token routinely lands one line above the sentence publishing
the proposition, so a naive `subject − predicate` difference reports **one publication as two** and buries
the real twins in wrap noise. Measured on the fixture world: 3 of 4 raw headroom sites were this.

**Decided:** subject hits sharing a **paragraph** with a predicate hit are not headroom. The paragraph is
the unit because §8's iter-93 rule is that *a reader lands on a paragraph and not on a file*.

**The risk this takes, stated and tested:** a suppression that swallows too much would hide the very twins
this tok exists to catch. The `ai`-fold pair — `external_services.md:554` / `:565` — sits **eleven lines
and several paragraphs apart**, so it survives. Both halves are asserted in one test
(`test_a_wrapped_subject_token_is_NOT_headroom_but_a_distant_twin_IS`), and a mutation control widens the
paragraph to the whole document and watches the distant twin disappear.

**Related:** the coverage invariant is at **FILE** granularity for the same reason — a line-level subset
test would go RED on ordinary wrapped prose, which is exactly how `anchor_offset_guard` false-REDed an
ordinary append (§8).

## D-M257x-113-3 — the multiplier moved 7.28× → 2.45×, and DOWN is the honest direction

**Measured.** iter-112: 29 seeds → 211 sites, 7.28×. iter-113: 29 seeds → **71** sites, **2.45×**.

**162 of iter-112's 211 sites came from four vocabulary forms** — `P16` 48, `P18` 58, `P22` 37, `P24` 19.
Those forms named the predicate's SUBJECT, not the predicate. Moved to the tier they belonged on, the four
enumerate **5 sites between them**. The 7.28× was substantially a count of how often the corpus writes
"Cosmo Router" about a component that was deleted.

**And the expansion that was real appeared where iter-112 could not see it** — the twin population
`D-M257x-109-4` predicted: `P21` 6 → **22**, `P10` 10 → **11**, `P15` 4 → 5, `P02` 1 → 2, `P12` 1 → 2.

**The number to carry forward is not the multiplier.** It is the **denominator: 71 sites**, each of which
either publishes the proposition or was examined and excluded by name. `repair_reach_guard` is graded
against that in step 2 (`TOK-07` rule 4).

## D-M257x-113-4 — the verdict is split in the OUTPUT, because only part of it is mechanical

16 predicates read `NO-EXPANSION` after the second pass. All 16 are settled — and they are **not settled
the same way**, so the tool emits two different strings rather than one:

| verdict | n | warranty |
|---|---|---|
| `SMALL-CLASS-PROVEN` | **1** (P20) | headroom is **zero**. No site mentioning the subject is unreached. Nothing was judged. |
| `SMALL-CLASS-ADJUDICATED` | **15** | headroom existed; every candidate carries an exclusion **with a reason**. This rests on judgement. |

**368 candidate sites were enumerated; 254 were read and excluded; 2 were promoted into the enumerated
set.** The fence guarantees the candidate set is complete and that no candidate is unexamined. **It does
not guarantee the 254 reasons are right, and this iter does not claim it does** — which is why the two
verdicts are separate tokens and why `FIX-M257x-iter113-adjudication-is-judgement` is routed forward
rather than closed. Collapsing them into one string would have let 254 judgement calls wear a
measurement's voice — the `fence_provenance` defect (`D-M257x-111-1`) in a new costume.

## D-M257x-113-5 — two REDs, both this iter's own, both caught by controls written before the ledger existed

Recorded as its own per iter-111's standard rather than dressed up as hidden failures it exposed:

1. **Coverage RED.** Tightening `P13` and `P24`'s subject forms to proposition terms left both tiers blind
   to the document their own proposition lives in (`customerio-sync.md`, `sentinel.md`). **The repair was
   the subject form; the invariant was not touched** — the direction that would have hollowed the fence
   out is relaxing the invariant, and §8's iter-98 rule already names it.
2. **Stale-exclusion RED.** Promoting `cms.md:171` and `ai_architecture.md:212` into the enumerated set
   left three exclusion rows excluding nothing. Removed.

Neither reached a published number, because both controls existed before the real ledger was authored.
That ordering is the point: iter-112's controls caught its derivation twice for the same reason.

## Three findings the ceiling surfaced that step 2 must not repair blind

- **`P08`'s pin is off by two.** The predicate pins the M51 iter-08/09 block at `ai-readiness.md:496`; the
  block opens at **`:498`**. Step 2 **re-derives** the anchor (§5 rule 22), never copies it.
- **`P13`'s "only" has a counter-example inside the corpus.** `external_services.md:495` records the
  **router** as having been built from a `git+url` context too. Not a publication of the predicate, so not
  an instance — but material to what the repair may assert.
- **`P24` is one survivor against ten witnesses.** The corpus states the messenger compose block was
  deleted at ten sites; `sentinel.md:5`'s presupposition that it still exists is the outlier.

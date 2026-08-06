# iter-110 — decisions

## `D-M257x-110-1` — the tok is DELIBERATE, and the streak is graded rather than assumed

Phase 0 rule 2 was checked before a word of strategy was written, because this milestone has twice had a
tok's trigger asserted rather than derived.

| tik | reading taken? | `N` |
|---|---|---|
| iter-107 (`TOK-06` step 2) | **no** | UNMEASURED |
| iter-108 (`TOK-06` step 3, the repair) | **no** — and it says so in §9's mandated words | UNMEASURED |
| iter-109 | **yes** | 33 → 36 (`P` 22 → 24) |

**One of the last three tiks measured.** §9's iter-type refinement — *a delta requires two measurements;
an iter that took no reading has an UNMEASURED metric, not an unmoved one* — means the streak's
precondition is unestablished. The trigger did not fire and could not have.

`TOK-07` is therefore **deliberate**: author-initiated, and **non-terminating** for `TOK-06`'s stated
reason — it sequences work iter-109 already routed and opens no new territory, so there is nothing
unreviewed for the next tik to commit to. Recorded here so that the *next* reader does not have to
re-derive whether an exit was owed.

## `D-M257x-110-2` — the multiplier series was RE-VERIFIED from the ledger, not quoted forward

The strategy's whole argument is four numbers, and this milestone's most-repeated defect is a number that
travelled without being re-derived (§5 rules 11/12; harden pass 25 found one inside the *rationale of a fix
for that class*). So each was opened at its source before it was used:

| pass | figure | verified at |
|---|---|---|
| iter-96 | 13 anchors → **51 sites** / 23 files; **38** sites an anchor-wise repair would have left | `iter-96/progress.md` lines 1, 11, 13, 14 |
| iter-98 | 20 anchors / 21 predicates → **37 sites** / 22 files | milestone `progress.md:1373-1374` |
| iter-102 | 52 anchors → **98 sites** found → 94 repaired | milestone `progress.md:1500-1504` |
| iter-108 | 31 primary anchors **derived from `iter-103/raw/`**; reach 46/47 raw, **46/46** over the upheld union | milestone `progress.md:1817-1828` |

**One correction to the framing this iter was handed.** The natural way to say it is *"iter-108's
multiplier was 1.0×"*. That is not what the ledger shows: **iter-108 reports no site-expansion figure at
all.** There is no multiplier to be low, because there was no expansion step to produce one. The stronger
and more accurate statement — and the one `TOK-07` uses — is that **the expansion step is absent from the
close, not present-and-poor.**

## `D-M257x-110-3` — a refuted premise does not condemn the instruments built under it

`TOK-06`'s premise is refuted (`D-M257x-109-3`): the residual is a standing pool sampled slowly, not a flow
being replenished, so fencing the inflows could not drain it. The reflex on a refuted strategy is to revert
its work.

**Measured, that reflex would have been wrong.** `TOK-06`'s induction leg took repair-induced anchors from
**21 % → 5.6 %** of the residual (band #10: 2 of 36, the lowest in the series, against ~2 per cycle for six
prior cycles at a far smaller repair size). `anchor_offset_guard` and `fence_provenance` are in production
and earning; `wrong-tree` went **4 → 1 → 1 → 0**.

**So steps 0–2 are RE-RANKED, not reverted** — and the one leg whose *ranking* rested entirely on the
refuted premise, the drift fence (`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`), is **de-ranked and
kept open** rather than cancelled: drift that is standing rather than arriving still wants a fence, just
not first.

**Generalised into the protocol** (§9), because the shape will recur in any milestone that revises a
strategy on measurement: *grade a refuted strategy leg-by-leg against what each leg measured, not
wholesale against its premise.*

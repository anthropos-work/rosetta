# iter-215 — decisions

## `D-M257x-215-1` — the miss set is PARTITIONED and ITEMISED; nothing is glossed

`unrecognised_blocker_tables` declared **91** misses and printed 91 identically-shaped lines. It now
labels each with **why** it missed, using `is_ledger_table`'s own vocabulary, and the report groups
them. `unrecognised_blocker_tables` becomes a **projection** of the new `blocker_table_misses` — one
walk, two views — so the labelled and unlabelled populations cannot drift (iter-210's rule).

The reconciliation is **fail-closed**: if the buckets do not sum to the declared total the guard prints
`MISS PARTITION DOES NOT RECONCILE` and exits **2**. An itemisation that is a subset of the population
is the failure mode this class keeps producing.

## `D-M257x-215-2` — ⚠️ this iter's own first cut smuggled an unmeasured verdict, and a 170-iter-old fixture caught it

The first grouping named the largest bucket **`no claim column`** and printed it as a bare COUNT with
the gloss *"summary tables; deriving 0 claims from them is correct, not a miss."*

`test_19`'s `_UNREADABLE_LEDGER` fixture — checked in around iter-44 — is headed
`| # | Anchor | Claim | What is true |`, is unmistakably a real ledger, and lands **in that bucket**,
because `_CLAIM_COL` (`false claim|what is wrong|the claim|issue|^text$`) cannot read the bare word
`Claim`. The test failed on the missing word `UNREADABLE`, and the reason it failed was the gloss.

Re-measured on the live milestone afterwards: **2 of the 57** carry a column a human would call a claim
column — `iter-116/raw/r29-B.md:100` (`claim | stated | measured`, 3 rows) and **`iter-48/raw/D.md:72`
(`Claim | Verified against`, 44 rows)**. So **U2 is FALSIFIED as written**, and the real near-miss set
is **36**, not 34.

Every bucket is now itemised and named for **what matched** — `neither column recognised`,
`claim column recognised only`, `anchor column recognised only`, `both recognised, same column` — with
no bucket carrying a semantic verdict the mechanism cannot support. A mechanical label that asserts a
meaning is exactly this milestone's defect class, committed while repairing it.

## `D-M257x-215-3` — the vocabulary widening is ROUTED, not landed

Three unreadable spellings are now measured: **`corpus claim`** (`iter-82/raw/r15-B.md:109`, 3 rows,
also carrying `corpus anchor`), bare **`Claim`** (44 rows at `iter-48/raw/D.md:72`), and **`claim`**
(`iter-116/raw/r29-B.md:100`). Widening `_CLAIM_COL` to read them changes the derived denominator and
therefore `claim_twin_guard`'s RED surface — iter-209's precondition (**zero false REDs measured
first**) applies and has not been paid. Routed with its exhibits.

Also routed, and **bounded**: `_ANCHOR_COL` spells `file` unbounded and matches *"pro**file**s"*.
Measured on the accept side — **0 of 68** accepted ledger tables have an anchor column that matched
`file` only as a substring — so it inflates a label and decides nothing.

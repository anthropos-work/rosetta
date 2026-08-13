# M258 iter-05 — decisions

## D13 — 247.79 s is quoted as the bring-up half; 249.13 s is NOT averaged into it, and 344.82 s is not dropped

Three reps, one of which passes headroom. The temptations were to (a) quote the **p50 249.13 s** because
buildbench printed it, or (b) average reps 2 and 3 because they agree within 1.34 s. Both are refused.

- **The p50 is over three reps, two of which the instrument itself calls "not usable measurements."** A
  p50 that includes unusable samples is not made usable by being a p50. `rep_is_ok` is the contract, and
  it excludes reps 1 and 2.
- **Rep 2 missed by 0.62** (`peak_load1` 10.62 vs a floor of 10). That is agonisingly close and it is
  still a miss. **A threshold honoured only when it is comfortable is not a threshold** — this release
  re-cut `D-v28-12` rather than grade inside a noise floor, and the same discipline applies to a margin
  this small. Rep 2 is reported as **corroboration** (it tells us rep 3 is not a lucky sample) and
  excluded from the quoted figure.
- **Rep 1 is reported, not dropped.** 344.82 s at `peak_load1` 21.77 is the spread `C2` demands be
  published. buildbench's own note asks whether eviction explains it; it does not — **load** does — and
  saying so is better than deleting the outlier.

So the honest statement is: **the single-box bring-up half is 247.79 s (n=1 gateable), corroborated at
249.13 s by a rep that missed headroom by 0.62, with a contended outlier at 344.82 s.** Not a p50 over
3 cold cycles, and not offered as one.

## D14 — the studio-desk route is closed as a FINDING, not as a lever, and the distinction is load-bearing

`CHECK-M258-iter02-studio-desk-is-the-untouched-leg` was routed as a *candidate lever if the composed
budget needs room* — 115.35 s, the largest UI leg, the one L1 never touched.

Measured across three reps it is **7.12 s** [7.11–8.05]. The 115.35 s was a **cold** build; iter-02's
bring-up followed a re-pin, so nothing was cached. Optimising it would have bought **≈0 s** on any cycle
that is not the first after a cache wipe.

Two consequences worth writing down:

1. **The route is closed with its question answered**, and it simultaneously **explains the 108.32 s
   delta** the milestone was told to explain before trusting either bring-up figure. One measurement,
   two open items discharged.
2. **The reserve the milestone was counting on is still `LEVER-M257-L5-setdress`.** `set_dress` measured
   **81.23 s** here against M257's 82.04 s — unmoved, and still the largest single phase. If the composed
   gate later needs room, that is where it is, exactly as M257's close said.

**The general form, because this milestone has now hit it twice:** *a phase's rank in a single-sample
table is not its rank in the budget.* iter-02 ranked studio-desk first on n=1; it is not even in the top
three warm.

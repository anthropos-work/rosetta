# iter-140 — decisions

## `D-M257x-140-1` — a receipt whose number was written from the CONCLUSION stops demonstrating anything

Both failures this iter share one authoring shape, and it is not carelessness about the *claim*:

| receipt | claimed | measured | conclusion |
|---|---|---|---|
| `sentinel.md:5` | *"returns **one unrelated hit**"* | **0 lines, exit 1** | **strengthened** — the absence is total |
| `latency-budget.md:365` | *"returns **one** non-test occurrence"* | **22 lines / 4 files; 3 non-test; 1 code** | **intact** — one *code* occurrence, two comments |

You know the answer; you write the command that demonstrates it; you fill in the count from **what you
know** rather than from **what it printed**.

> **Rule.** Write the number from the output. **The defect is not the number — it is that the receipt no
> longer demonstrates anything.** A reader who runs `latency-budget.md`'s grep sees **22** where the page
> says **one**, and the rational response is to distrust the paragraph, *including the parts that are
> exactly right*. A receipt is the strongest form of citation precisely because it is checkable; a
> receipt that fails its own check is weaker than no receipt at all.

Positive controls were run in the identical invocation form for both (`colony` at `fa47850d` returns hits in
messenger's own `cmd/` package, three files), so *"the command form is broken"* is excluded
and the absence is real.

## `D-M257x-140-2` — a class is censusable iff its subject carries its own HEAD

iter-138 tried to census bare `:NN` pins and iter-139 retracted the result at **0/12**. iter-140 censused
published receipts and got a usable answer at **7/9 reproduce**. Same strategy (`TOK-08`), same author,
one iter apart, opposite outcomes.

> **Rule.** **The variable was never the strategy — it was the subject's decidability.** A bare `:NN`
> pin has no resolvable head (*which file does a continuation pin continue?*). A **receipt names its own
> command, pathspec and ref**, so it is self-contained and re-runnable by construction.
>
> Before censusing a class, ask **"does an instance of this class carry enough to resolve itself?"** If
> not, the census will measure your resolver's guesses, not the corpus — which is exactly the 0/12.

This retires the temptation to read iter-139 as *"censusing does not work here."* It works; it needs a
subject that can be resolved.

## `D-M257x-140-3` — the not-checkable bucket is published, not absorbed

**9 of 22** receipts were runnable on this box at the ref each names. The other **13** are reported as
**not-checkable**, in their own count, and are **neither passes nor failures**.

Reporting *"7 of 9 reproduce"* without the 22 would be precisely the over-claim iter-139 retracted one
iter ago — a rate over the population the instrument could reach, presented as a rate over the class.
Routed as `FIX-M257x-iter140-receipts-not-checkable-here`.

## Upheld claims counted as results

**Seven receipts reproduced exactly**, including three with structure and not just a count —
`security_compliance.md:235` (8 hits, and both named files correct), `dependency_map.md:59` (6 lines, and
all three named files correct), `studio-room.md:367` (22 hits across 3 files). Recorded because a census
that only reports its failures is not a census.

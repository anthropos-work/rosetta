# iter-61 — decisions

## `D-M257x-61-1` — iter-60's GREEN over-reported, and the briefed "17 files" was right all along

iter-60 fenced the profile predicate in two forms (command, table-first-cell), repaired those, and
read **GREEN**. Re-surveyed at the same platform ref, the class's **larger half** was still standing
in a third form the guard could not see — the **noun phrase**, *"the default `graphql` profile"* —
at **34 raw sites across 17 files**.

**That is exactly the "17 files / 30 occurrences" this milestone was handed**, and `D-M257x-60-7`
recorded it as an *undercount* on the strength of a broader grep. It was not an undercount. It was a
**different, well-defined construct** — and it is the one that carries the class. `D-M257x-60-7`'s
third row is **withdrawn**; the other two corrections in it (the `main.go` line numbers) stand.

The general form is the one worth keeping: **a fence whose reach is narrower than its class
over-reports its own GREEN**, and the over-report is invisible precisely because the fence is the
thing you would use to check. The countermeasure is not a wider regex — it is enumerating the
*forms* a predicate can be written in before believing a GREEN.

## `D-M257x-61-2` — two new constructs, still constructs

* **Noun phrase**: a backticked token adjacent to the literal word `profile` (or `profile (\`tok\`)`).
* **Table row**: `| \`NAME_RPC_ADDR\` | \`value\` | …` — the same binding as `NAME=VALUE` with no `=`.
  `messenger.md` carried two stale values in this shape for the whole time the fence read GREEN.

Neither is a substring widening. "GraphQL" the API — named constantly here — still cannot match,
because the token must be inside a backtick span *and* adjacent to the word `profile`.

## `D-M257x-61-3` — prose needs two discriminators a command line does not

35 raw hits → 22 real, and **13 were the guard's own** — removed by rules derived from the corpus's
own writing, not by exceptions:

* **Negation.** *"and no `cms` profile"* asserts the token's **absence**. Anchored to the END of the
  preceding text so only an **adjacent** negation counts; a "no" earlier in the sentence about
  something else cannot launder a real claim (tested both ways). **Note the recursion**: this is the
  form iter-60's own corrections are written in — `D-M257x-60-5` made the corpus say *"there is no
  cms profile"*, so an undiscriminating widening reads **this milestone's repairs** as fresh defects.
* **Ref-pin**, the exemption G2/G4/G5 already used and G1 did not. `external_services.md:425`
  describes what `b56d731` did. `_REF_PINNED` also gains the **bare backticked sha**, the corpus's
  usual way of opening a sentence about a commit, which carries most of the historical narrative.

## `D-M257x-61-4` — the residual is routed WHOLE, not repaired in part

22 sites / 12 files, enumerated in `evidence/residual.md` with the command to regenerate it.
**Not repaired here, and deliberately not repaired in part.** §5 rule 19's scope-edge corollary is
explicit: a claim leaks to the edge of the previous repair's scope and **pools there**, so a subset
repair leaves a half-consistent corpus that costs the next auditor its budget in adjudication.
`FIX-M257x-iter61-profile-prose-class` → iter-62, as one unit.

The instrument landing without its repair is the honest split: the fence is now **RED and correct**,
where before it was **GREEN and narrow**. A RED fence with an enumerated residual is a strictly
better state than a GREEN one that cannot see the class.

## `D-M257x-61-5` — the routed citation target was substituted, and re-measured before routing on

`TOK-05` named iter-61 = the 21 `main.go:N` citations. Step 0 found both candidates real and chose
the fence gap, because the gap was created **by this milestone one iteration earlier** and makes a
committed GREEN misleading. The citation class is re-routed to iter-62 with a **refreshed**
measurement rather than iter-58's: of 16 distinct `app/main.go:N` citations at app `v1.366.0`,
**5 still land on the construct they claim** (`:446`, `:524`, `:604`, `:816`, `:992` — all verified
by reading the line) and the rest have moved (`:971-973` is now a comment about collapses, `:1178`
is `defer cancelServerContext()`, `:1196` is the **skiller** handler, not jobsim).

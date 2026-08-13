# iter-245 — decisions

## `D-M257x-245-1` — the platform slice gets NO existence fence, and the measurement is the reason

iter-244 fenced the rext slice because a rext section is an **unambiguous prefix**. A repo name is not.
`app/` is a directory inside studio-desk; `jobsimulation/` is a directory inside `app/internal/`. **4 of
4 first-pass findings were head collisions**, and the fourth defeated a citing-document-based
disambiguator because the *sentence* named the repo while the *document* did not.

An existence fence here would false-RED, which `§8` rule 6 says gets a fence disabled. The route closes at
**zero for the repo-rooted clone-present slice**, and stays open for the repo-relative (203) and
clone-absent (131) classes — neither decidable from this clone set.

## `D-M257x-245-2` — grade the decidable half of a range; keep the route open on the other, in the same line

`ROUTE-M257x-h59` was routed whole because *which line of a range carries the claim* has a 490-anchor blast
radius. **Whether the file has that many lines never did.** The arm grades resolution + bounds and the
disclosure states both halves in one sentence, so a green on the graded half can never read as coverage of
the ungraded one.

## `D-M257x-245-3` — the bounds arm under-flags by design

An anchor is out of bounds only if it is out of bounds at **every ref its block names, and at the
worktree**. Measured: 2 of the arm's first 5 findings were ref mis-attribution, one of them
`app/CLAUDE.md:289-294` — an anchor that is exactly correct at HEAD and fired because a *different* claim
in the same block named an older commit. Under-flagging is the correct direction; a bounds arm that REDs a
correct anchor gets disabled on first contact.

## `D-M257x-245-4` — the inherited RED is repaired, and its cause is reported as a pattern

`test_anchor_subject_census_m257x` was RED on the tree that opened this run, from iter-240's own edit to
`setup_guide.md`. With iter-244's finding, that is **two consecutive iters leaving a different `stack-core`
test RED**. The common cause is not carelessness: both iters' closes quote **`guard_family`**, which runs
the guards and **not** the test suite behind them. Repaired and routed as
`ROUTE-M257x-245-guard-family-green-is-not-suite-green`.

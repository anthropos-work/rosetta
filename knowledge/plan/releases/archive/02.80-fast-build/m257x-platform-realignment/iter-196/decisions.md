# iter-196 — decisions

## `D-M257x-196-1` — the half-true sentence, caught one iter after it cost nine

Reading Go made TypeScript the last unread language, and the obvious next sentence is *"Playwright
needs a live stack, therefore unreadable."* **It is half-true**, and the false half is worth **424
tests**: `playwright test --list` enumerates with no stack, no browser and no server.

This is precisely iter-195's defect one language over, and it was avoidable only because iter-195 had
just written the rule down. The generalisation worth keeping: **an unreadable-because clause names a
capability, and capabilities are per-verb.** Playwright cannot *run* without a stack and can *list*
without one; the exclusion was written against the strongest verb and inherited by every weaker one.

**First enumeration:**

| section | tests | files | notes |
|---|---:|---:|---|
| `playthroughs` | 215 | 45 | deps **borrowed** — ships no `node_modules` |
| `stack-verify` | 209 | 30 | own install |
| **TOTAL** | **424** | **75** | |

The 75 files corroborate iter-186/187's file counts (45 + 30) exactly, independently derived.

## `D-M257x-196-2` — the true half must survive the finding, so the VOCABULARY is fenced

These are e2e tests against a live demo stack. **A listed test PARSES and REGISTERS; it has not
passed.** The risk created by this iter is that `424` gets quoted next to Go's `2,714` as though both
were verdicts — and the two are different kinds of statement.

So the fence is on the words, not only the number: `ts_census` returns `tests` / `files` / `listed` /
`borrowed_deps` and **an arm asserts there is no `pass` or `fail` key**. If the vocabularies ever
converge, a reader can quote one as the other and be right about the words, which is how a population
becomes a green.

The printed block says it in the tool as well: *"This tool reports a POPULATION here and a VERDICT for
Go, and the two must never be quoted in the same breath."*

## `D-M257x-196-3` — a cold checkout cannot enumerate `playthroughs` unaided, and that is disclosed

`playthroughs/e2e` ships no `node_modules`; the enumeration only worked by borrowing `stack-verify`'s
install and setting `NODE_PATH`. The six **Go** sections needed nothing at all.

That asymmetry is a fact about this repo, not an implementation detail: a census that quietly borrows
dependencies hides a real prerequisite from whoever reads its number. `borrowed_deps` is returned,
printed, and fenced — and the fence is written so that `playthroughs` shipping its own install later
**fails loudly as good news** rather than passing silently.

## `D-M257x-196-4` — the same NOT-LISTED rule as iter-195's unrunnable

No `Total: N tests in M files` summary line means the listing did not happen, and it is recorded as
`listed: 0` rather than as `tests: 0`. Proved with a synthetic section carrying a deliberately broken
`playwright.config.ts`. Third instance of one rule in three iters — iter-191's false CANNOT-RUN,
iter-195's silent-zero build failure, and this — which is why it is now written as a general rule in
the protocol rather than as a third worked example.

## `D-M257x-196-5` — the population is a FLOOR, not a pin

`424` and `75` are asserted with `assertGreaterEqual`. A pin would go RED on every legitimate new spec
— the failure mode `test_the_real_preset_ships_four_orgs` demonstrated for three milestones, training
readers to ignore a working guard. A floor catches the direction that matters: silent shrinkage.

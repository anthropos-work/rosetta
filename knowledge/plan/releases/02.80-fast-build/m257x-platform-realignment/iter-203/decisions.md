# iter-203 — decisions

## `D-M257x-203-1` — the docstring census REPORTS three classes and asserts on none of them

`docstring_measurement_literals(root)` enumerates every measurement-shaped number in a string literal
that is **not** inside a `print(...)`, across rext non-test modules, and classifies each:

| class | example | count |
|---|---|---|
| `dated` | *"at iter-122 an early draft reported 256 findings, all false"* | **53** |
| `doc-relative` | *"the correction sits 65 lines away"* | **7** |
| `standing` | *"205 of 695 line-pinned citations are bare basenames"* | **35** |
| | **total** | **95 across 62 sites** |

**Only `standing` can rot, and it is still not asserted on.** `dated` and `doc-relative` are
mechanically decidable; *"this present-tense sentence asserts a current property of the tree"* is not —
an undated sentence about a past draft reads exactly like a standing one. Failing on `standing` would
grade prose, which is the over-reach `D-M257x-202-4` already priced twice inside one iter (434 pairs,
then 21 re-labels).

So what is fenced is the **size**: a ratchet at 95 that the population may fall below and may not rise
above without a recorded reason. The class is now bounded and visible instead of unmeasured. Both
mechanical classes are additionally required to be **non-empty**, because a classification with no
instances cannot be wrong — which is not the same as being right — and the `dated` split carries a
**mutation control** that strips the marker from one sentence and requires the classification to move.

## `D-M257x-203-2` — the census's own denominator came from a throwaway probe, and that is the class

This iter's `overview.md` and the first draft of the census's block comment both said **85 across 55
sites**. The census itself returns **95 across 62**. The 85 came from a scratch probe written minutes
earlier that did not apply the ordinal rule and skipped a different set of helper modules.

**Quoting a scratch measurement beside a fence is the exact defect the fence enumerates**, committed in
the act of describing it. Corrected in place, at both sites, with the provenance stated — a number in a
census's own comment has to come from that census.

## `D-M257x-203-3` — `205 of 695` was stale in both operands, and it named no unit

`claim_census_guard.basename_index`'s docstring carried *"205 of 695 line-pinned citations"* as a
standing property of the corpus. Derived on this tree:

| grain | bare | line-pinned |
|---|---|---|
| pair (claiming unit × citation) | **410** | **979** |
| **distinct citation text** — what the sentence means | **292** | **704** |

**Both operands had moved, and the numerator moved mostly because of iter-202.** Stopping the extension
alternation truncating `.json`/`.jsx`/`.tsx`/`.graphqls` recovered 35 citations that were invisible and
re-spelled 63 more, and those are overwhelmingly **bare filenames** — `package.json`,
`AIReadinessClient.tsx`, `useNavbarSections.tsx`. So the first member of this class is a figure this
milestone invalidated one iter earlier and did not notice.

**The missing unit is half the defect.** *"205 of 695"* is checkable against neither grain because it
names neither. Repaired to state both, and fenced by `TheBasenameShareIsDERIVED`, which recomputes
both operands from the live census and asserts the docstring carries them — plus an **anti-vacuity**
arm requiring the two grains to actually differ (if they ever coincide, the unit label is untested and
the fence proves nothing) and an arm requiring the words *distinct* and *pair* to be present at all.

This is iter-199's *make the printed count derive itself*, transposed: a docstring cannot self-format,
so the derivation lives in a test that recomputes and asserts the text. It can still go stale — but no
longer silently.

## `D-M257x-203-4` — the registry's completeness fence went RED on its own author for the third time

Adding `docstring_measurement_literals` turned `test_every_executable_derivation_is_classified` RED
within a minute of the function existing, exactly as it did for `printed_arithmetic_totals` (iter-193)
and `printed_measurement_literals` (iter-199). Classified at source as
`DECLINE:verdict` — the return value is findings, so a test literal equal to it would be a baseline with
its own regeneration discipline, not a frozen copy of a corpus constant.

Worth recording as a **positive** result rather than a chore: a registry that catches its own author
three times, on three different days, is a registry that has not silently fallen behind the tree — which
is the entire claim `unclassified()` was built to make.

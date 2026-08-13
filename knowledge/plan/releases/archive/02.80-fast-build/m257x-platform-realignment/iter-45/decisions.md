# iter-45 — decisions

## `D-M257x-45-1` — three fences, not one general linter

iter-42 classified iter-41's 18 blockers by the **cheapest instrument that could have caught each**, and
the three residual classes want three different instruments, not one clever one:

| fence | asserts | reaches |
|---|---|---|
| `markdown_structure_guard` | a document's STRUCTURE is undamaged | `#6` |
| `anchor_construct_guard` | a resolvable `file:line` lands on a CONSTRUCT | `#13`, `#16` |
| `derived_value_guard` | a stated scalar equals the source it names | `#10`, `#11` |

Merging them would have produced a single "corpus lint" whose findings could not be reasoned about
separately, and whose false-positive budget would be shared across three unrelated rule families. Kept
apart, each one's reach is a separate, falsifiable number.

## `D-M257x-45-2` — every rule measured tree-wide BEFORE adoption, and the rejected drafts kept

§5 rule 2, applied to lint rules. Each guard's docstring records the false-positive rate of the drafts
that did not ship, next to the rule that replaced them:

- **M1, first draft: 7 hits, 6 false** (86%) — an INDENTED blockquote inside a list item followed by the
  next sibling item, which is correct markdown. Narrowed to **column-0 on both sides** → 1 hit, blocker #6.
- **M2, three drafts** — `in in-app`, then `default-ON on demo`. A hyphenated compound is one word to a
  reader and two to `\b`, so **both** sides of the pair needed a hyphen guard.
- **The anchor self-reference rule, first draft: 134 findings, essentially all of them ports** (`:5050`,
  `:3000`). Restricting bare `:NNN` to the *immediately followed by above/below/earlier* construct took it
  to 2, both real.

## `D-M257x-45-3` — blocker #17 is DROPPED from this instrument, not tuned into it

`platform-migration-status.md:60` attaches `app/main.go:604` to a sentence naming four domains; `:604`
wires jobsimulation only. An "enumeration-attached anchor" rule fired on **6 of 7** tree-wide candidates
and still missed #17, because the other three domains' wiring calls sit within a few lines of `:604`.

Narrowing the window until #17 fires and its neighbours do not is **Trap A** — tuning a fence to its own
answer key. Catching #17 requires deciding what a sentence *claims*, which is the line this whole fence
family does not cross. **Routed to step 4's hand repair**, named, rather than manufactured into a fence.

## `D-M257x-45-4` — the five blockers are asserted against a CAPTURED fixture, not the live corpus

The first draft of the behaviour suite asserted all five **live**, and passed. Every one of those
assertions would have **failed at iter-46**, whose entire job is repairing those five sites — and the
obvious repair would have been to edit the fence's test to match. That is how a fence stops asserting
anything: not by deletion, but by maintenance.

So `tests/fixtures/mechanical/{red,green}/` holds two **line-faithful repo roots** captured from the
tree iter-41 measured. Two of the five defects are *relationships between line numbers* (`:788` citing
`:447` in the same file; `:110` citing line 815 of another), which the claim-twin fixture's ±2-line
neighbourhood shape would destroy — so this fixture snapshots whole files, with the one large vendored
platform source under **line-preserving elision** rather than compaction.

The live tree is still asserted, on the three properties that **survive the repair**: the resolver still
resolves (>50 anchors), unmeasured docs are still named, the ratchet still grades the tree GREEN.

Generalized into `platform-alignment.md` §8 as **rule 7**.

## `D-M257x-45-5` — the green twin must still be REACHED and MEASURED, not merely silent

A fence that stopped resolving a site prints exactly the same clean run as one that passed it (§5 rule 8,
arriving through the fixture instead of through the code). So `test_10b` asserts the repaired twin still
resolves **2** anchors, and `test_17` asserts it is still **measured** — silence has to be earned by the
repair, not by loss of reach.

For the same reason the green transforms are **declared and mechanical**, never hand-edited. The anchor
transform re-points to the *nearest preceding content line* — deliberately weaker than "the correct
construct", because `anchor_construct_guard` cannot assert that (see `D-M257x-45-3`), and a green twin
that proved more than its fence claims would be a fixture tuned past its instrument.

## `D-M257x-45-6` — the mutation battery found three real holes on its first run

20 mutants: 1 declared-GREEN no-op that survives, 19 kills, ≥5 inversions, and **one mutant per reporting
path in all three modules** — the direct debt of harden passes 7–9, where two of the previous fence's
honesty mechanisms turned out never to have been implemented and one deleted clean with 15/15 green.

Three mutants came back **GREEN**, each naming a hole in a behaviour suite that had looked complete:

| mutant | the hole |
|---|---|
| `md-indent-tolerance-removed` | M1 is "column-0 on **both** sides" and only one side was tested |
| `md-empty-scan-reads-as-clean` | a `corpus/` with zero scannable files reported a clean pass |
| `value-unmeasured-counts-as-measured` | `measured` counted a doc no scalar was ever read from |

All three closed with named tests. **The battery is not a formality that ratifies a suite; on this
iteration it was the only thing that measured it.**

## `D-M257x-45-7` — a defect in iter-44's own ratchet: `--accept` rewrote records it did not move

Registering the three new fences replaced `claim_twin_guard`'s baseline `reason` — iter-44's registration
sentence — with a sentence about three fences that postdate it. `registered_at` already preferred the
prior value; `reason` did not, and **that asymmetry is what made the loss silent in review**.

Fixed: `--reason` is written only onto fences whose site set this acceptance actually moved. iter-44's
reason restored verbatim, two regression tests added (**and the negative half too** — a fence that IS
lowered must still take the new reason, or the ratchet's progress would be filed under the sentence that
explained the previous state), plus a mutant in iter-44's own battery.

**A record that silently rewrites itself is this milestone's own defect class, found this time in the
instrument built to prevent it.**

## `D-M257x-45-8` — a captured fixture is not this repository's source

The fixture's vendored `assignments.go` made `test_write_target_schema_fence` report **`stack-core`** as
an unclassified Go-bearing rext section — a section that ships no Go, over a file belonging to the
platform. The next step would have been scoring another team's captured source for rext's schema writes.

Fixed in the fence, not the fixture: all three of its walks now prune `fixtures` **directly under
`tests`**. Narrow by construction, and the regression test asserts **both** directions — a section that
genuinely keeps Go under a directory called `fixtures` elsewhere is still scanned.

The alternative — classifying `stack-core` in `SECTION_COVERAGE` to make the fence pass — is the exact
move that fence's own failure message forbids: *"Do NOT default to 'n/a' to make this pass: that is the
whole defect, written down."*

## `D-M257x-45-9` — the rext pin is NOT moved

`.agentspace/rext.tag` and the `stack-demo` clone stay at `fast-build-m257x-iter-37`. Everything this
iteration ships is **offline guard and test code on no runtime path** — three `*_guard.py`, their tests,
their fixture, and a `--reason` bookkeeping fix in a commit-time hook. Nothing a bring-up executes
changed, so re-pinning would alter what the next gate measurement runs against for no benefit. TOK-02's
conjunction-risk note (*clauses 1/2/4 should be undisturbed — should be, and that must be verified rather
than assumed when the pin moves*) is therefore **not yet triggered**; it fires at the first iteration that
changes runtime source.

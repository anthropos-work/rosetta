# iter-75 — decisions

## `D-M257x-75-1` — the "92 unrepaired citations" are 0 defects, 77 unreachable and 26 undecidable

Adjudicated against `git ls-files` over every clone — **7,265 tracked basenames across 13 clones**,
the repository's own answer rather than a directory walk:

| fate | sites | basenames |
|---|---|---|
| UNIQUE — exactly one tracked path | **77** | 26 |
| MULTI — more than one | **26** | 19 |
| **ABSENT — no such file anywhere** | **0** | 0 |

**Not one citation in the class names a file that does not exist.** The class is the ninth reach
limit of this milestone, not a repair backlog: iter-73 taught a bare `<name>.<ext>:N` to **reach**
the resolver and did not teach the resolver to **find** it, because every route `resolve()` owns is
positional and an ops doc citing `` `up-injected.sh:1487` `` supplies no position at all.

**Fourth consecutive routed count to collapse on adjudication** — 64 → 5, 23 → 1, 21 → 0, and now
92 → 0 defects. The orchestrator's sentence is now carried by four measurements: *every routed count
is a hypothesis.*

## `D-M257x-75-2` — uniqueness is the safety argument, and 26 sites staying unresolved is the rule working

The new route fires only when `git ls-files` names **exactly one** path for the basename. `main.go`
is **57** tracked files, `main.tf` 10, `mixin.go` 3; `studioManager.go` is 2 —
`app/internal/cms/studio/` and `cms/internal/studio/`, the merged copy and the standalone husk,
which is precisely the pair a directory guess would get wrong and precisely the fold this milestone
exists to document.

Same argument `AMBIGUOUS_BASENAMES` already makes for `.md`, and the same one this guard's docstring
prices at *"134 findings, essentially all of them ports."* **26 sites stay unresolvable and are
named in the reach line.** An unresolvable citation that cannot be resolved without guessing is
coverage, not debt.

Two further restrictions, each with its own mutant:

- **Bare citations only.** A path-qualified citation has already said where the file lives;
  resolving it by basename would override the document with a guess about its directory — a
  different act from filling a silence.
- **`git ls-files`, never a walk.** A build artifact or scratch file with a unique basename must not
  become a resolvable target, and must not make a genuinely unique basename read ambiguous.

## `D-M257x-75-3` — the two `rosetta-extensions` clones are ONE witness, and the bug was mine twice in one iteration

`rosetta-extensions` is cloned twice under this tree: the per-stack consumption copy at
`stack-demo/rosetta-extensions` (**pinned at a tag**) and the authoring copy at
`.agentspace/rosetta-extensions`. Both directories carry the same **name**.

The first adjudication script keyed its universe by `<clone-name>/<relpath>` and collapsed them
(right); the dry-run script keyed by absolute path and split them (wrong). **The same
misunderstanding produced two different wrong answers in one iteration** — the adjudication said
`up-injected.sh` was in *"2 places"* printing the same path twice, and the dry run silently found
only 31 of the 77.

Resolved by deciding **which clone is the witness**, not by de-duplicating harder: the authoring
copy is the current one, the per-stack clone is pinned, and `resolve()`'s **pre-existing** rext
fallback already prefers the authoring copy — so the rule follows the guard instead of inventing a
preference. Pinned as a test (`test_the_two_rosetta_extensions_clones_do_not_split_a_basename`) and
as mutant M4.

**The rule worth carrying:** when a repository appears twice in a tree, a basename index must say
*which copy is the witness* before it says whether a name is unique. Two clones of one repo are one
witness. §5 rule 32's lesson one level in — *two instruments disagreeing is a finding*, and here
both instruments were mine.

## `D-M257x-75-4` — a route that fires silently is a reach claim nobody can audit

`resolve()` returns a bare `Path`, so the rule that found a file cannot be read from its return
value. The new route increments `RESOLVE_ROUTES` and the guard prints `resolved via
bare-unique-basename x77` beside its existing reach lines.

This is deliberately **not** a widening of `run()`'s return tuple: that tuple has grown to seven
positional values and already broke four callers once (`RF-M257x-iter71-run-returns-a-tuple`).
A module-level counter cleared at `run()` entry is the smaller edit and the one that does not
re-open a known defect.

## `D-M257x-75-5` — a 0-finding dry run is only evidence if the pipeline can report a finding

The dry run returned **77 newly resolvable, 0 findings**, which is the shape of a broken pipeline
as much as of a clean corpus — and iter-73's comparable widening had turned the corpus RED with 6,
so 0 was the *surprising* answer. The same code path was therefore fed three known-bad inputs
against a real 2,693-line target: `:99999` → `anchor-out-of-range`, `:155` → `anchor-on-blank-line`,
`:1`/`:2693` → clean.

§5 rule 2 applied to a *derivation* rather than to a search. **A measurement that returns the
convenient answer earns a control before it is believed**, and the cost here was one command.

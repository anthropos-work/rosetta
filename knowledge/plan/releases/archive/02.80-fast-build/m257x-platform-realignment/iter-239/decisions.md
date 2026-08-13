# iter-239 — decisions

## `D-M257x-239-1` — the skill's own declaration is the contract; a runnable example is not

**Both sides were candidates.** Three of the four runnable `/stack-snapshot` examples in the live corpus
spell it verb-first (`replay 1`), so the *majority* spelling in the corpus is the one the skill does not
declare. A majority is not a contract, and two independent statements inside the skill settle the
direction:

* `argument-hint: [dev-N|demo-N] [replay|capture|status] …` — target first;
* the skill's own flow line, `.claude/skills/stack-snapshot/SKILL.md:32`, `/stack-snapshot N replay` —
  target first.

The skill agrees with itself twice; the corpus disagreed three times. **Repaired the corpus, in the
direction of the declaration.** The same reasoning settles the bare-`N` half: `stack-seed/SKILL.md` writes
`--stack demo-N` / `--stack dev-N` at **7 of 7** of its own command examples and has no bare form anywhere.

Recorded rather than assumed, because the opposite call was defensible on volume alone — and volume was
the wrong test.

## `D-M257x-239-2` — arm D grades a DERIVED set of skills, never a named one

The first hand-written pass of this iter's instrument hard-coded five "generic stack-ops skills"
(`stack-list`, `stack-secrets`, `stack-seed`, `stack-snapshot`, `stack-update`) from the CLAUDE.md
convergence note. **Derived from the argument-hints, the set is three**: `stack-list` declares `(no args)`
and `stack-update` declares `[dev-N]` — a *dev-only* slot, so `demo-N` is not a legal target there and the
grading would have been wrong in both directions on two of five members.

The guard therefore derives its subject set from the presence of a literal `dev-N|demo-N` slot in the
hint. `§5` — **a fence pinned to a SPELLING is not pinned to a PROPERTY**; a hard-coded list also goes
stale the day a sixth stack-ops skill lands. `test_arm_D_target_slot_set_is_DERIVED_not_hard_coded` adds a
new skill at test time and requires the guard to grade it without being edited.

## `D-M257x-239-3` — retired skill names get NO arm of their own

The six v1.3-M14 hard-renames were censused: **96 sites across 6 names, 0 of them runnable in the live
corpus.** All 96 are historical planning records (`knowledge/plan/**`, `CHANGELOG.md`) or a SKILL.md prose
line naming its own predecessor (*"Formerly `/demo-seed`"*), which is exactly what those documents are for
— correcting them would falsify the record.

No arm was added, for two reasons. **Arm A already catches it**: a retired name has no `SKILL.md`, so the
moment one becomes runnable it is an unresolvable invocation. And a dedicated arm would have to *spell* the
retired names to fence them, which is the `retracted_pin_guard` class one document over — a retraction that
reproduces the thing it retracts re-arms the finding it closes.

## `D-M257x-239-4` — the two undeclared presets are classified, not repaired

`stack-seeding/presets/` holds **8** seed files; the `/stack-seed` hint enumerates **6**. The 6 it names all
resolve (0 false promises). The 2 it omits — `gen-batch-20`, `gen-batch-org-fill` — are the AI-generation
batch descriptors of `ai-generation-spec.md`, driven through `--seed <path>`, not `--preset`. **An
under-declaration is not a false promise**, and nothing presents them as `--preset` values. Classified here,
repaired nowhere — the same disposition iter-237 gave its 28 orphan env names.

## `D-M257x-239-5` — bare-`N` is graded as CORRECTNESS, not as SAFETY

Traced before booking severity, because "ambiguous stack target" invites a safety framing this evidence does
not support:

* `blueprint.ParseStackN("1")` **succeeds** — it takes everything after the first `-`, and a name with no
  `-` parses whole. The bare form is accepted silently; it does not fail loudly and it does not misparse.
* `isolation.TargetFor` sets `IsProd` **only** for the literal `"production"` / `"prod"`. An unqualified
  target is non-prod, so the never-write-prod firewall holds unchanged.

So the finding is that the reader must supply a prefix the tool's every other example carries — real, and
worth a fence — but the isolation contract is untouched. Stated in the guard's own docstring so the fence
cannot later be quoted as evidence of a safety gap it never measured.

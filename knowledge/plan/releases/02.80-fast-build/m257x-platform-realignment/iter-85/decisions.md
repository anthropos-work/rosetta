# iter-85 — decisions

## D-M257x-85-1 — the scope was declared NARROW on purpose, and that is the iter-83 lesson applied to itself

40 findings were available; this iter repaired **9** (Q2's 7 + 2 leak sites) plus the rext defect, and
**declared the rest out of scope at open**.

iter-83 measured what happens when a repair reports a completeness it did not measure: **74.1 %**, eleven
predicates called discharged, four of them wrong. The defence is not "try harder" — it is **declare a
scope you can complete and grade against it**. `closed-fixed` grades planned scope (Phase 4 Step 0), and
the residual is routed with a ledger, not left as an impression.

Q1 is deliberately *not* here despite being the cheapest class (13 anchors, ~4 re-derivations): mixing a
bulk re-anchor into the pass that repairs seven judgement-heavy deleted-fact claims would blur the
grading, and the grading is the deliverable.

## D-M257x-85-2 — the `dev-stack` default profile is DERIVED, never a literal — and the negative path was tested in the real call form

`dev-stack:186`/`:414` held `profile="graphql"`. Substituting `"core"` would have been wrong for the
reason `D-M257x-59-2` and TOK-04 P4 already give: **derive, else fence, else declare.** A literal is
correct until the next rename, and the platform renamed this one **four days ago**.

Both entry points now default empty and resolve via `platform_topology.default_profile()`, which reads
`backend`'s own `profiles:` list and **raises rather than guessing** — the same primitive
`gen_injected_override.py` has used on the demo path since iter-55. The dev path kept the literal for four
more releases, which is the measured cost of fixing one caller instead of the shared derivation.

**The first cut of this fix was WRONG, and the suite caught it — recorded rather than smoothed over.**
I made the derivation **fatal** on any failure. That broke **13 tests** in `test_dev_public_host.py`,
whose fixtures use *synthetic* platform dirs with no real compose file, and it took the M220 mutation
battery down with them — not by killing a mutant, but by breaking its **baseline**
(*"the UNMUTATED subject fails its own suite … no RED below means anything"*). A stricter contract than
the codebase's own is still a regression.

**The correct shape already existed one file over**: `gen_injected_override.profile_for()` catches
`TopologyError` and returns `FALLBACK_PROFILE`, with a fence
(`test_profile_is_derived_on_the_production_path`) proving the constant can never reach a real bring-up.
This now **mirrors it exactly and imports the SAME constant** — a second `"core"` literal in the shell
would have been the very defect being removed, one file over.

**Three things proven rather than asserted:** `derived => core` against the real clone; `core` via the
fallback against a synthetic dir **without dying**; and a **die** that still fires when `python3` itself
cannot run — because an *empty* `--profile` selects only the floor, which is the same failure this change
exists to remove. 129/129 green across `test_dev_public_host` + `test_dev_stack`.

**And my own comment block contradicted the code for one edit cycle** — it still said the derivation
*"fails loud instead of silently bringing up nothing"* after I had added the fallback branch. Deleted.
That is `CHECK-M257x-iter77-narration-vs-documentation` committed by the author of the fix for it, inside
the same function, within minutes. Third occurrence of that class in this run.

The `/dev-up` skill docs were realigned in the same commit. Fixing a tool and leaving its documentation
describing the old default induces the doc↔tool mismatch the pass exists to remove.

## D-M257x-85-3 — the leak fence caught MY repair, and the disposition is a WAIVER, not a paraphrase

`repair_leak_guard` — the guard iter-81 did not run — went **RED** on this iter's own working tree,
naming `corpus/ops/platform-alignment.md:1060`: §5 **rule 40**, which I authored at iter-83, quotes the
false `graphql`-profile sentence verbatim as its worked example.

**Waived, with a written reason, in `repair_leak_waivers.json`.** The guard's own contract is that a
waiver can only make the fence quieter and is therefore **reported on every run**, and requires both a
`path` and a `form_contains` — no wildcards, no path-only exemptions. Chosen over paraphrasing because a
rule that will not name the sentence it is about teaches nothing, and the quote sits inside an explicit
*"…while the sentence '…' stood untouched"* construction no reader can misread.

**Worth recording for what it says about the fence family.** The reach fence I built at iter-83 graded
this repair 11/11 and was *silent* about this — correctly: reach is *"did you open what you were given"*,
and I did. The leak fence asks a different question and caught what reach could not. **Three fences, three
questions, and this iter needed two of them.** The one that fired is the one iter-81 skipped.

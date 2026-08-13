# iter-91 — decisions

## `D-M257x-91-1` — the user's question answered: a QUALIFIED NO, with something better in its place

Asked: should `guard_family.py` **refuse to run** against a clone whose remote-tracking refs are stale,
rather than answering from what it can see?

**Not as put, and the reason is a measurement rather than a preference.**

1. **"Stale" is not locally decidable.** A clone fetched sixty seconds ago can already be behind; one
   fetched last week can be current. Deciding it needs the network — and a fence that cannot run offline
   is a fence that stops being run (the same pressure that made `_git_out` memoised: an uncached run took
   `anchor_construct_guard` from ~1 s to 10.9 s, and *a guard slow enough to skip is a guard that stops
   being run*). So the network check exists and is **opt-in**: `--verify-remote`.
2. **"Cannot see the objects it needs" IS locally decidable**, and it is the condition that actually bit.
   That is fixed **at the point of use**, which is strictly better than a family-level heuristic: only the
   guard knows which refs it needs. This runner places sixteen heterogeneous guards, and encoding their ref
   requirements here would rebuild exactly the hand-maintained tuple §2 exists to refuse.
3. **Refusal is kept for the one case that is both locally decidable and unambiguous:** a platform-facing
   run against a clone with **no `origin/main` at all` → `exit 2` UNMEASURED, *before any guard runs*.

**And the gap nobody had named was smaller and worse.** The runner printed `platform=<dir>` and never a
**sha**. Every `13 GREEN` transcript in this milestone therefore names a *directory*, not a commit — which
is precisely how a green reading gets quoted forward into a run brief with no way to re-check it. The
reference line now prints the corpus sha, the platform sha, its `origin/main`, whether they agree, and the
fetch age, on every run.

## `D-M257x-91-2` — `unresolvable` and the silent worktree fallback are now UNMEASURED, not GREEN

`platform_alignment_guard` had exactly one positive control: `subject_checked == 0`. So **total** resolver
failure tripped, and **partial** blindness was folded into a GREEN verdict.

Two conditions now return **exit 2**:

- **the silent worktree fallback** — `cited_text`'s `auto` ladder ends by reading the CHECKOUT when no ref
  resolves, returning provenance `worktree(no-ref)`. The string has always been there; **nothing read it.**
- **`unresolvable > 0`** — each is a claim the run did not check.

Exit 2 rather than 0 or 1 because *"I could not measure this"* is neither a pass nor a failure, and
`guard_family.py` already renders exit 2 as CANNOT-RUN and refuses to call the family green. The escape
hatch is `ALIGNMENT_ALLOW_UNMEASURED=1`, which **records** the gap rather than hiding it — the same
contract as `--allow-not-run`.

Measured before and after, on the shipped map:

| reference | before | after |
|---|---|---|
| `auto`, refs present | GREEN (90 resolved, 0 unresolvable) | **GREEN — unchanged** |
| `worktree` (the stale-clone fallback) | RED, 8 findings, **4 unresolvable ungraded** | **UNMEASURED (exit 2)** |

The second row is the point: it *was* returning a verdict — and a RED one, which looks like diligence —
while 4 citations had not been checked at all.

## `D-M257x-91-3` — the 7-guard conjunction-pair enumeration (Decision 1 item 3)

The generalisation iter-89 recorded and iter-90 proved: **a specification with seven guards needs at least
one test per PAIR that can interact, not one per guard.** Enumerated. Of the 21 unordered pairs, most are
independent axes (G1 path × G4 idempotency have no shared state or ordering); **12 can interact**, and each
is graded below against what the suite actually asserts.

| pair | interaction | before iter-90 | now |
|---|---|---|---|
| G2 × G5 | self-healed apply ⇒ sha is neither pin ⇒ revert refuses | **UNCOVERED — the defect** | ✅ covered |
| G4 × G5 | can a patch still come off after an idempotent re-apply? | **UNCOVERED** | ✅ covered |
| G1 × G5 | does the path firewall run on the verb that also does `git checkout`? | **UNCOVERED** | ✅ covered |
| G6 × G5 | is demo-only scope enforced on revert? | **UNCOVERED** | ✅ covered |
| G5 × G5′ | chain: two patches on one file, LIFO unwind (cross-**invocation**) | **UNCOVERED** | ✅ covered |
| G2 × G7 | self-heal recomputes post ⇒ G7 must check the MARKER, not the recorded sha | covered | ✅ |
| G3 × G5 | a staged write must self-revert | covered | ✅ |
| G1 × G2 | no byte written before the path assert | covered | ✅ |
| G1 × G6 | both halves of the firewall, one call | covered | ✅ |
| G2 × G4 | "already patched" is G2's coherence probe, not a sha match | covered | ✅ |
| G3 × G7 | post-condition, then the unstaged assert — ordering | covered | ✅ |
| **G5 × G7** | a FAILED post-condition restores pristine — and must also drop the journal | — | ⚠️ **routed** |

**Five interacting pairs were uncovered; four now have tests.** The two landed this iter (G1×G5, G6×G5) are
the safety-critical ones and they were entirely untested: **every** G1/G6 escape test in the file drives
`apply`, while `revert` is the verb that writes *and* runs `git checkout -- <path>`. `cmd_revert` did call
`assert_demo_clone_path` — but nothing said so, and nothing would have noticed its removal. Mutation-proven:
deleting that call turns 3 subtests RED.

Routed: **`CHECK-M257x-iter91-g5xg7-journal-on-postcondition-failure`** — apply's post-condition-failure
branch restores the pristine bytes and drops the journal, but the drop is not asserted, because triggering a
genuine short write needs fault injection the harness does not have. Named rather than quietly skipped.

## `D-M257x-91-4` — two guards report GREEN against a tree they never read (side finding)

Surfaced by iter-91's own mutation run, not sought. With the freshness refusal mutated out, the family ran
against an empty temp dir and reported:

```
guard-family: 2 GREEN · 2 RED · 9 could-not-check · 3 not-run
```

Nine guards correctly said *COULD NOT RUN — no corpus/ under <tmp>*. But **`story_org_count_guard` and
`union_apply_guard` returned GREEN** — against a directory containing no corpus at all. They resolve their
inputs relative to the rext checkout and **ignore `--repo-root` entirely**.

For `union_apply_guard` that is arguably correct: its subject IS the rext manifest set. For
`story_org_count_guard` it is not — its own summary says *"and every doc agrees"*, and those docs live in
the **corpus**, i.e. in the tree it just ignored.

Not chased — it is a third line of investigation and this iter declared three steps. Routed:
**`CHECK-M257x-iter91-guard-repo-root-scoping`**. The reason it matters is the same reason this whole iter
exists: **a guard that answers about a tree it did not read is a cannot-tell wearing a verdict**, and one
of the two is in the family whose green this milestone quotes.

## `D-M257x-91-5` — `CHECK-M257x-iter90-realmanifest-baseline` re-derived and left OPEN, deliberately

Re-examined with the wider evidence from iter-90 (three manifests, two vehicles). It is **not** a re-pin
task, and re-pinning now would destroy the signal a future iter needs:

Under M217 **the anchor is the contract; the whole-file sha is only a baseline.** A test asserting that a
shipped `pre_sha256` still matches a live, persistently-updated clone therefore asserts the property the
design deliberately stopped requiring — it is the *same* rot M217 removed from apply and iter-90 removed
from revert, still resident in the test layer. Re-pinning would make it green today and stale again on the
next `make pull`, which is the definition of the rot.

Left open, with the shape of the answer recorded: the assertion should be re-scoped to *the anchor resolves
exactly once in the live clone* — a property that survives upstream drift, which is the whole point. Not
landed here because it touches three tests across two vehicles and belongs with the adjudication, not with
a fence iter.

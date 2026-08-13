**Type:** tik (under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07))

# iter-245 — the platform slice measured ZERO, and the zero paid for something better

## Phase B — the reading, and the instrument that produced four artifacts before it produced a fact

Of the **335** corpus citations whose first path segment names a clone present in `stack-demo/`
(**105 distinct paths**), graded at the clone's worktree **and** at `origin/main`:

**4 distinct paths appeared to be missing. All four were the instrument.** Each is a **head collision** —
a first segment that names a real repo while meaning something else:

| apparent finding | what it actually is |
|---|---|
| `app/services/config.ts`, `app/core/main.ts` | **studio-desk repo-relative.** studio-desk has its own `app/` directory; both files exist under `studio-desk/app/` |
| `app/services/userService.ts` | the same, cited from `clerkenstein.md` — the **sentence** names studio-desk (*"studio-desk's `STUDIO_ACCESS_ROLES`"*) while the citing **document** does not, so a citing-doc-based disambiguator misses it |
| `jobsimulation/ai/ai.go` | a **mid-sentence abbreviation** of `app/internal/jobsimulation/ai/ai.go`; the same sentence spells the full path two clauses earlier |

**So the platform slice measures ZERO wrong paths, and `P-244-1`'s falsification branch fires as
written.** The result is worth having *because* it is a zero: `anchor_construct_guard`'s 599-strong
out-of-reach bucket is **genuinely out of reach, not a hiding place for defects** — at least on the
repo-rooted, clone-present slice.

**And it is why the platform slice must NOT get iter-244's fence.** The rext sections are unambiguous
prefixes; repo names are not. `app/` is a directory inside studio-desk. `jobsimulation/` is a directory
inside `app/internal/`. **4 of 4 first-pass findings were head collisions** — an existence fence over this
population would be a false-RED generator, and that is measured rather than assumed.

## Phase C/D — what the zero paid for: the range population, graded on the half that IS decidable

The census turned up a better target on the way past. **117 of the 335 were RANGE citations
(`path:NN-MM`), and 117/117 resolved** — while `anchor_construct_guard` was reporting **490 range
citations, 24.8 % of its qualified population, in NEITHER of its two counts** (harden pass 59's
`ROUTE-M257x-h59-range-anchors-are-ungraded`).

The route is open because *which line of a range carries the claim* is a design decision with a 490-anchor
blast radius. **But that is not the only question a range answers.** Two properties need no such
decision:

> **does the path resolve**, and **does the range lie inside the file**.

`NN-MM` past EOF is wrong under **every** reading of which line matters. So iter-245 grades that half and
leaves the route open on the other, saying so in the disclosure line.

**Result — 490 range citations, from "in neither count" to:** `349 resolved and BOUNDS-checked, 141
refused`, and the refusals are head-keyed exactly like the single-line bucket (`infrastructure` ×28,
`terraform` ×16, `services.tf` ×9).

### It found 3 real defects on its first run — one class

All three in `corpus/services/storage.md`: **a bare basename inside `corpus/services/<svc>.md` that the
sentence says means a DIFFERENT repo.**

| site | cited | the sentence says | resolved to | repair |
|---|---|---|---|---|
| `storage.md:74`, `:192` | `README.md:81-87` | *"**platform**'s own README"* | `storage/README.md` (33 lines) | → `platform/README.md:81-87` |
| `storage.md:77` | `main.go:518-523` | *"**app**'s two boot guards"* | `storage/main.go` (18 lines) | → `app/main.go:518-523` |

Both targets were verified to carry the claim: `platform/README.md:81-87` is the AWS-credentials block
verbatim, and `app/main.go:518-523` is the empty-bucket `log.Fatalf`. **The repair is the citation, not
the prose.**

### Two of the first five findings were the arm's own ref handling, and both are now controls

* `repos.yml:29-31` — whose own sentence says *"**at that ref**"*. Fixed by extending the arm the
  **alternate-ref ladder** the single-line arm already uses (`block_ref_candidates`).
* `CLAUDE.md:299` → `app/CLAUDE.md:289-294` — which is **exactly** the *"don't delete the `ai` repo"*
  block at HEAD (356 lines), and fired only because a *different* claim in the same block names
  `b948604f`, where the file is 221 lines. `block_ref` attached a ref the anchor never claimed. Fixed by
  making the **worktree the last candidate**: under-flagging is the correct direction for a fence, and a
  bounds arm that REDs a correct anchor over a mis-attributed ref gets disabled on first contact.

## Phase E — pre-registrations, graded after the last edit

| id | prediction | outcome |
|---|---|---|
| **P-245-1** | ≤ 40 distinct paths missing at both refs | **HELD — and the answer is 0.** The falsification branch fires as written |
| **P-245-2** | ≥ 1 finding is a RANGE citation | **HELD, by the other arm** — 0 existence findings, but the bounds arm this iter built found **3**, all ranges |
| **P-245-3** | `infrastructure` reports could-not-check, never green | **HELD** — absent from the clone set; **45** single-line + **28** range citations sit in the refusal buckets, named |
| **P-245-4** | `app` supplies the plurality of findings | **REFUTED** — `app` supplied 3 of the 4 apparent findings and **all three were artifacts**; the real defects were 3/3 in `storage.md` |
| **P-245-5** | ≥ 1 finding in a frozen-legacy repo | **HELD** — all 3 are `storage`, one of the six `ROUTE-M257x-241` named as graded by nothing |

## Close — 2026-08-10

**Outcome:** the platform slice of iter-244's class measures **ZERO** with a proven instrument (4 of 4
first-pass findings were head collisions), and the measured zero is what justifies **not** extending the
fence there. The iter's shipped deliverable is instead the **range-bounds arm** on
`anchor_construct_guard`: **490 range citations — 24.8 % of the qualified population and in neither of the
guard's counts since harden pass 59 — are now `349 resolved and BOUNDS-checked, 141 refused`**, gating at
zero after **3 real repairs** in `storage.md`.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-245-1` (the platform slice gets NO existence fence — head ambiguity, measured 4/4) ·
`D-M257x-245-2` (grade the decidable half of a range; leave `h59` open on the undecidable half and say so
in the same line) · `D-M257x-245-3` (the bounds arm under-flags by design: worktree is the last ref
candidate) · `D-M257x-245-4` (the inherited RED is repaired, and its cause reported as a pattern).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6), scoped
`-k "anchor or guard_family or fence_registry or repair_postcondition or rext_path"`: **356 passed / 0
failed**, including **6 net-new** bounds-arm cases. Guard family (`--platform`, from repo root): **27 GREEN
/ 0 RED / 0 could-not-check / 5 not-run**.

**Side-deliverables:**
- **A second inherited RED, from the same pair of iters as iter-244's.**
  `test_anchor_subject_census_m257x` was RED on the tree that opened this run: `platform-alignment.md:1054`
  cited `setup_guide.md:504` for the literal `migrations: true`, which **iter-240's own edit moved to
  :514**. Verified pre-existing at `a9b0fef` by reading the blob there. Re-pointed. **The pattern is now
  two-for-two: iters 239 and 240 each left a `stack-core` test RED, and both were found only because
  iters 244/245 ran pytest** — those iters' closes quote the guard-family runner, which does not run the
  test suite. Routed as `ROUTE-M257x-245-guard-family-green-is-not-suite-green`.

**Routes carried forward:**
- `ROUTE-M257x-245-guard-family-green-is-not-suite-green` → **new**, two-for-two.
- `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` → open (iter-244).
- `ROUTE-M257x-h59-range-anchors-are-ungraded` → **HALF-CLOSED.** The bounds/resolution half is graded and
  gating; the which-line-carries-the-claim half remains open and is now stated in the same disclosure line.
- `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` → **measured, and closed for the repo-rooted
  clone-present slice at ZERO.** Still open for the **repo-relative** (203) and **clone-absent** (131,
  `infrastructure` ×45) classes, neither of which is decidable from this clone set.
- `ROUTE-M257x-241-wider-citation-surface-is-ungraded` → **partly served** — the 3 repairs are all in
  `storage`, one of its six frozen-legacy repos.
- `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` · `ROUTE-M257x-238-claude-md-fences-are-unmaintained` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-critical-env-list-is-unfenced` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **A measured ZERO is a licence not to build.** The instinct under a census strategy is to fence every
   class. Here the zero, plus the 4-of-4 head-collision reading, is precisely the evidence that a fence
   would be a false-RED generator — and *that* is the deliverable.
2. **An undecidable question usually has a decidable component. Split before you defer.** `h59` was routed
   whole because "which line carries the claim" is hard. "Does the file have that many lines" was never
   hard, and it was sitting inside the same route for a full harden cycle.
3. **A head that names a repo does not mean the path is rooted at that repo** — the platform-slice twin of
   iter-244's `alignment/` case, and the reason these two slices need different instruments.
4. **`guard_family` green is not suite green**, and two consecutive iters proved it. A runner that reports
   27 GREEN says nothing about the 1,985-test suite behind those guards.

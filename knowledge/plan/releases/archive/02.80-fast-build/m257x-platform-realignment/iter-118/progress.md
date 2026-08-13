**Type:** tik · `iter_shape: census` · **`TOK-08` class 2 — platform-source citation resolution.**

# iter-118 — the fence was green over a subject it only two-thirds reached

## The numbers, in `TOK-08`'s shape

| | value |
|---|---|
| **enumerated population** | **1,070 platform-source citation candidates** across 112 files |
| **already false at census** | **0** — the fence was and is green over everything it resolves |
| **reach** | **675 / 1,070 = 63.1 %** |
| **denominator provenance** | **`citation-candidates-minus-non-citations`**, printed in the report and in `--json` |
| **excluded, named** | **119 non-citations** (URL/scheme authorities) |

**Invocation:** `python3 stack-core/anchor_construct_guard.py` with `ROSETTA_ROOT` set, rext `4304930`.

**The finding is not a defect count; it is a denominator.** This fence already reported *"675 resolved,
514 unresolvable"* and already named its unresolvable heads — the discipline was in place. What nobody
had done was **look inside the head**. 119 of the 514 are `http://sentinel:8087`-shaped: matched because
`sentinel` is a path-ish head and `8087` is a line number. **They were never citations**, so the fence
was grading its own coverage against a target that included things it could not possibly resolve.

That is iter-114's rule — *a reach metric is settled by its DENOMINATOR's provenance* — **arriving one
layer over, inside the guard that reports reach.** It printed no ratio at all before this iter, which is
how the inflation survived: there was no number for anyone to distrust.

## What is now measured rather than hidden

The 395 genuine unresolvables, named:

| head | n | why it cannot be resolved |
|---|---|---|
| `(bare)` | 276 | bare `` `:NNN` `` — a port, or a continuation of a platform file named earlier. **The same shape class 1 measured as undecidable** |
| `main.go` | 27 | every Go repo has one; the basename is not an identifier |
| `studioManager.go` · `wiring.go` · `mixin.go` · tail | ~92 | single-file basenames needing a repo-disambiguation rule |

**Routed, not attempted** — the repo-disambiguation rule is a real lever and it is a different iter's
work. The tripwire held: planned scope landed, the rest routed.

## The sweep, against iter-117's sealed definition

| class | population | findings | reach | controls that can fire |
|---|---|---|---|---|
| 1 — intra-corpus citation | 1,520 | **0** | **100 %** | mutation ×7 · anti-vacuity ×3 · regression ×6 |
| 2 — platform-source citation | 1,070 | **0** | **63.1 %** | mutation ×3 · anti-vacuity ×3 |

**The full mechanical sweep `TOK-08` pre-registered is COMPLETE.** The class list was fixed in iter-117,
could only grow, and did not. **`TOK-08`'s trigger is therefore armed: the next iter is the grading
reading**, and it grades `P >= 19` (enumeration-first REFUTED, hand back to the user) against
`P <= 18` (working, say so with the number), on the baseline `P = 37` at corpus `f581de09`.

## Close — 2026-08-07

**Outcome:** Class 2 censused. **1,070 platform-source citation candidates**, **0 false**, reach
**675/1,070 = 63.1 %** over a denominator that now names its provenance — after **119 non-citations**
(URL/scheme authorities, `http://sentinel:8087`-shaped) were excluded, counted and named. The fence had
been reporting coverage against a target inflated by things that were never citations, and printed no
ratio at all, which is why it survived. Anti-vacuity is the load-bearing control when a fix REMOVES from
a denominator, and it is written that way. **The full mechanical sweep is complete as pre-registered
(class 1 at 100 % reach, class 2 at 63.1 %, both at 0 findings), so `TOK-08`'s grading reading is now
the next iter.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (`TOK-08`'s branch is
**not gradeable until the reading**; grading it on the sweep alone would be the flattering reading) —
(4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n —
(7) budget-exhausted: **y — session budget spent BETWEEN ITERS, tree clean, both repos pushed** —
Outcome: exit-7
**Decisions:** `D-M257x-118-1` (a URL is not a citation — the reach denominator was inflated) ·
`D-M257x-118-2` (anti-vacuity is load-bearing when the fix removes from a denominator) ·
`D-M257x-118-3` (the sweep is complete as pre-registered; the residual is named, not closed)
**Side-deliverables:** `test_iter45_mechanical_fences.py::test_21` hard-coded a four-fence baseline set
and broke whenever a fence joined the ratchet. Now DERIVES the expected set from disk — iter-44's own
lesson (a fence registry is derived, never maintained as a list) applied to the test guarding it — with
an anti-vacuity floor. Does not change the close status.
**Routes carried forward:**
- **iter-119 = `TOK-08`'s GRADING READING.** The sweep is complete; the branch is `P >= 19` refuted /
  `P <= 18` working, against `P = 37` at `f581de09`. **Do not soften either side.**
- **`FIX-M257x-iter118-bare-basename-needs-repo-disambiguation`** → net-new, open. ~92 single-file
  basenames + `main.go` ×27 are unresolvable only because the basename is not an identifier. A
  service-doc → repo mapping would reach most of them and lift class 2's 63.1 %.
- `FIX-M257x-iter116-intra-corpus-miscitation-is-the-largest-class` → resolution half closed at
  iter-117; **construct half open**, now with a number (4 of 387 bare-pin lines are machine-resolvable)
- `FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block` · `-induction-fences-do-not-scale` ·
  `FIX-M257x-iter113-adjudication-is-judgement` · `FIX-M257x-iter115-anchor-guard-resolves-fixture-paths-live` ·
  `FIX-M257x-iter111-staged-battery-dependency-is-underived` · `-buildbench-parse-json-is-a-noop-flag` ·
  `FIX-M257x-iter108-stackcore-suite-hangs` · `FIX-M257x-iter107-drift-fence-satisfiable-by-prose` (de-ranked) ·
  `DEF-M257x-iter101-briefing-rext-tree` → all open
**Lessons:**
1. **A green fence with a named unresolvable head still has to have the head READ.** The discipline that
   makes coverage visible does not make anyone look. 119 non-citations sat in that bucket across
   multiple iters, in a fence explicitly built so its reach could not shrink in silence.
2. **A fence that prints no ratio cannot have its denominator distrusted.** The inflation survived
   because there was no number to argue with. Printing the reach is what made it falsifiable.
3. **When a fix REMOVES from a denominator, the anti-vacuity control is the load-bearing one.** The
   mutation half only proves the exclusion fires; the other half is all that stops a predicate that
   excludes everything from reporting 100 % reach over nothing.

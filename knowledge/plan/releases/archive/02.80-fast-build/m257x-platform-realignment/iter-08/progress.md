---
milestone: M257x
iter: 08
---

# iter-08 — progress

**Type:** tik (under `TOK-01`, step 3 — *land the fences, each watched going RED, before trusting any green*)

## The re-survey refuted the iter it inherited from

iter-07 routed `FENCE-M257x-write-fence-scans-one-section-of-nine` forward on two claims. Step 0
re-measured both **before any work**, and both failed.

| iter-07 claimed | measurement |
|---|---|
| the scope limit is undocumented — *"nothing says so"* | **FALSE.** A **ten-line justification comment sits directly above `SCORED_SECTIONS`**, with a correct Trap-A rationale for excluding `stack-snapshot`. iter-07 read line 92 and the module docstring, and not the ten lines between them. |
| the fix is to widen the tuple | **A no-op.** The fence's *own scanner*, run over every Go-bearing section: `stack-seeding` **92** constructs, `stack-snapshot` / `stack-secrets` / `clerkenstein` / `alignment` / `playthroughs` **0 each**. Widening would score nothing and report GREEN. |
| (assumed) the fence can pass vacuously | Already closed — `test_the_seeders_actually_write_something` asserts `found > 40`. iter-06 had thought of it. |

**This is the milestone's dominant defect committed by the milestone itself** — a state reported without
being measured — and §5's closing rule already names it: *verify a claim before escalating it, including a
claim made by an audit.* Recorded prominently in `iter-07/progress.md` too, because an inherited false
finding is the thing this milestone exists to stop.

## What actually survived, and it is worth the iter

**1. The exemption's REASONS had gone stale — and iter-07 is what made them stale.** The comment argued
`stack-snapshot` was safely unscored because its one stale target *"already fails LOUD at replay time …
rc=4"*, tracked by `REPOINT-M257x-cms-similarity-writes`. **iter-07 closed that route and deleted that
signal**: the replay now resolves and succeeds. A justification whose evidence has been removed still reads
as live — which is precisely how it misled the reader who quoted the constant above it.

**2. Nothing mapped a section to which of §8's three layers covers it.** *"Is `stack-snapshot` fenced?"* has
a correct answer — **the live layer**, because after `D-M257x-8` its write target is resolved at run time and
there is genuinely nothing static to see — and it was written down nowhere. Gate clause 4 rested on a reader
guessing it right. iter-07 guessed wrong.

**3. The fence's own SCOPE was a hand-maintained list of the system's parts** — the same shape as the migrate
tuple iter-02 deleted, one level up, and in the worst possible place: **a fence only asserts about what it
already scans, so an unclassified section is invisible by construction.** It cannot go RED, because nothing
looks at it.

## What was built

`SECTION_COVERAGE` — every Go-bearing rext section, its `(layer, reason)`, derived and fenced:

    stack-seeding    static  seeders COPY into a stack's Postgres with literal schema names
    stack-snapshot   live    replay resolves its target schema at RUN TIME (D-M257x-8)
    stack-secrets    n/a     writes .env files, never Postgres
    clerkenstein     n/a     an in-memory Clerk mock; no Postgres write path
    alignment        n/a     scores mirror-vs-source genes; no Postgres write path
    playthroughs     n/a     drives a browser against a running stack; no direct Postgres write

`SCORED_SECTIONS` is now **derived from that map**, so the two cannot drift; widening scope means
classifying a section, not editing a tuple. And the map is derived against the repo: a section that gains
Go code and no classification goes **RED naming itself**.

**The subtlest of the five new tests is `test_the_static_layer_actually_scores_its_sections`.** A section
classified `static` that yields **zero** scoreable constructs is **mis-classified, not covered** — and it
reports GREEN, which is strictly worse than leaving it out, because it *looks* fenced. That is exactly the
trap iter-07's own pre-compute predicted and this iter measured; it is now unwritable rather than merely
warned about.

## Fences — 5 new, all mutation-verified RED

| # | mutation | fence that fired |
|---|---|---|
| M1 | a new Go-bearing section appears, unclassified | `test_every_go_section_declares_its_coverage` |
| M2 | `stack-snapshot` mis-classified `static` (the *looks-covered* trap) | `test_the_static_layer_actually_scores_its_sections` |
| M3 | a classified section removed from the map | `test_every_go_section_declares_its_coverage` |
| M4 | `SCORED_SECTIONS` hand-edited independently of the map | `test_scored_sections_is_derived_from_the_map` |
| M5 | a coverage reason degraded to a stub | `test_every_entry_declares_a_known_layer_and_a_reason` |

Each mutant was **collected before being run** — the Python analogue of iter-07's new `§8 rule 5` (a mutant
that does not compile is not a RED fence). All five produced a real, named failure.

## Gate clause 4

**Now claimable, and on better evidence than iter-07 would have claimed it on.** The condition (zero rext
writes to a schema the platform no longer creates) holds; `REXT_TRANSITIONAL_SCHEMAS` is empty; and the
*asserting* side is now a written, machine-checked statement of **which layer covers which section** rather
than a reader's guess. Left for the milestone close to call formally — this iter's job was to make the claim
honest, not to award it.

## Routes carried forward

| item | why | target |
|---|---|---|
| `FIX-M257x-academy-not-serving` | Now the **only genuine ✗** in autoverify — the other two are `CHECK-M257x-bringup-evidence-logs-absent` (evidence-absence, not defects). It is the last thing between here and a green cold cycle, and **clause 1 needs three of them**. | next tik |
| `CHECK-M257x-bringup-evidence-logs-absent` | Unchanged from iter-07; bears on how clause 1's "0 warnings" is read. | next cold cycle |

## Close — 2026-07-31

**Outcome:** the inherited finding was **refuted by measurement** (the boundary was documented; widening
would have scored 0), and the narrower real gap was fixed: the fence's own scope is now **derived** from the
repo with a declared layer+reason per section, so an unclassified section goes RED naming itself, and a
section that merely *looks* covered cannot.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-12` (a fence's scope is derived and fenced like any other list of the system's parts)
**Side-deliverables:** none.
**Routes carried forward:** `FIX-M257x-academy-not-serving` (now clause-1's last blocker) · `CHECK-M257x-bringup-evidence-logs-absent`
**Lessons:**
- **Read the lines around the line you are quoting.** iter-07 quoted `SCORED_SECTIONS = ("stack-seeding",)`
  and reported its rationale absent; the rationale was the ten lines immediately above it. Grepping to a
  line number and reading only that line is a search that returns a true substring and a false conclusion —
  the same family as §5's other false-absence traps, and not yet on that list. Added as §5 rule 10.
- **A fence only asserts about what it already scans**, so its scope is the one list that can never announce
  its own staleness. Derive it. (`platform-alignment.md` §8, new subsection.)
- **"I scanned it" and "I found nothing to check in it" are different findings**, and only one is coverage.
- **When you close a route, grep for the comments that cited it.** iter-07's fix silently invalidated the
  safety argument in a comment 200 lines away, and that stale argument is what the next reader trusted.

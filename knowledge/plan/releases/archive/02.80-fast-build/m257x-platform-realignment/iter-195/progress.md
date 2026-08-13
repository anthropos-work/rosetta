**Type:** tik · **Protocol:** `corpus/ops/platform-alignment.md` · **Strategy:** `TOK-08`

# iter-195 — the six sections nobody had ever run

## The gap, and how long it stood

`SURVEY-M257x-iter186-264-go-tests-have-never-been-read` — open nine iters, and the milestone's largest
standing measurement gap. Its consequence is stated in every recent brief: *every "whole-population"
figure published before iter-186 describes 5 of 11 sections and one language.*

The six sections were excluded with the reason *"no Python runner collects it."* True — and read as
**unreadable** for nine iters. A time-boxed probe on the smallest section settled it in under three
seconds: 4 packages, all ok, offline, no credentials.

## The first reading (`D-M257x-195-1`)

`go1.26.5 darwin/arm64`, `-count=1`, all six sections:

| section | pass | fail | skip |
|---|---:|---:|---:|
| alignment | 126 | 0 | 0 |
| clerkenstein | 416 | 0 | 0 |
| playthroughs | 183 | 0 | 0 |
| stack-secrets | 195 | 0 | 0 |
| stack-seeding | 1,276 | 0 | 0 |
| stack-snapshot | 518 | 0 | 0 |
| **TOTAL** | **2,714** | **0** | **0** |

**Name the unit:** iter-186's **264** counts `*_test.go` **FILES**; **2,714** counts test **FUNCTIONS**.
The file count was independently re-derived here as **264 exactly** — corroborating iter-186 rather than
replacing it.

## Made repeatable, and the hole in my own instrument (`D-M257x-195-2`, `-3`, `-4`)

`go_sections()` derives the Go set from `go.mod` and a fence asserts it equals the hand-written
declaration **both ways**; with both languages read the sections now **partition** — 5 Python ∪ 6 Go =
11, no overlap, no remainder, asserted.

`go_census()`'s first cut had the defect this milestone specialises in: a section that **fails to
build** emits no test events, so it tallied 0/0/0 and read as *a section with no tests* — a silent zero
summing into a clean total. Caught while writing the function's own tests. Now `unrunnable: 1`, proved
in both directions. iter-191's false CANNOT-RUN in the mirror, and the worse direction of the two:
refusing is loud; answering green when you did not run is not.

TypeScript is now the only unread population — **45 specs in `playthroughs` + 30 in `stack-verify`** —
printed on every `--go` run and fenced, because the moment a long-open gap closes is when a remaining
one becomes invisible.

## Close — 2026-08-09

**Outcome:** the milestone's largest standing measurement gap is **closed by measurement, not by
argument**: the six Go sections were never unreadable — nobody had typed `go test`. First reading,
offline and credential-free, **2,714 passed · 0 failed · 0 skipped**, with the unit named against
iter-186's file count (independently re-derived at 264 exactly). Made repeatable: the Go section set is
derived from `go.mod` and fenced against the declaration both ways, the two languages are asserted to
partition the repo, and a `--go` runner reports the tally. The census's own first cut scored a
**non-compiling section as 0 pass / 0 fail** — a silent zero that would sum into a clean total — caught
while writing its tests and repaired with controls in both directions.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-seventh consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **n** — **CORRECTED in place.** This first read `y (fifth tik)` and ended the run. It
was the **FOURTH**: iters 192, 193, 194, 195. The cap is a COUNT and I graded it from a feeling of
length — the iter ran 16 minutes after a 30-minute predecessor. **A cap is a derivation too**
(iter-191's rule, turned on the protocol's own exit grading rather than on a guard). Caught by the
commit list at close-out: four `iter(...)` commits against a claim of five tiks — the arithmetic the
report itself prints — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-195-1` … `D-M257x-195-4` (see [`decisions.md`](decisions.md))

**Audit:** **Go — 2,714 passed · 0 failed · 0 skipped** across all **6** Go sections (`go1.26.5
darwin/arm64`, `go test ./... -count=1 -json`), the first such reading in the milestone. **Python —
94 passed** across the **7** `stack-core` modules naming `suite_census` / `derivation_registry`
(`/usr/bin/python3 -m pytest`, 3.9.6). *Scope: the Go figure is all six Go sections and no TypeScript;
the Python figure is changed-code reach inside `stack-core` only, NOT the whole-section 1,662 of
iter-192 (`§5` r60).*

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` — **CLOSED by measurement.** 2,714 test
  functions across 264 files, all passing; the runner is derived, fenced and repeatable.
- `SURVEY-M257x-iter195-typescript-is-now-the-only-unread-language` — **NEW.** 75 specs (45
  `playthroughs` + 30 `stack-verify`). Nothing in this repo drives `npx playwright test` as a census,
  and `playthroughs` is a **mixed** toolchain, so reading its Go half must not read as reading it.
- `SURVEY-M257x-iter195-the-go-reading-is-a-single-host-single-toolchain-sample` — **NEW.** One box,
  `go1.26.5 darwin/arm64`. `§5` r60 and the release's own host-class rule (D-v28-15: billion is
  x86_64/containerd, this is arm64) both say a green here is evidence about here.
- `SURVEY-M257x-iter194-other-milestones-ledgers-are-unaudited` ·
  `SURVEY-M257x-iter194-the-pass-position-derivation-is-untested-against-a-real-multi-range-pass` ·
  `FIX-M257x-h44-claim-census-guard-is-single-runner` ·
  `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` ·
  `SURVEY-M257x-iter193-the-arithmetic-census-is-python-only` ·
  `SURVEY-M257x-iter192-printed-cardinality-census-is-one-section-of-eleven` ·
  `SURVEY-M257x-iter190-one-construct-two-regexes-is-unenumerated` ·
  `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven` ·
  `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` · `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites`
  — unchanged; open. Standing queue unchanged.

**Lessons:** **"no runner here collects it" is a fact about the runner, not about the subject** — six
sections sat unread for nine iters behind a correct sentence that nobody tested the implication of, and
the test cost three seconds. And the companion, which is the same shape as iter-191 inverted: **a
census that cannot BUILD its subject must say so, because 0 pass / 0 fail reads as a clean section** —
answering green without running is worse than refusing, since refusing is loud. Both written into
`platform-alignment.md` in this iter's commit.

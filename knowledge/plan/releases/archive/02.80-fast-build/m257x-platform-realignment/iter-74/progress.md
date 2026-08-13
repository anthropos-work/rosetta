# iter-74 — progress

**Type:** tik, under `TOK-05`. Shape: single planned target
(`CHECK-M257x-iter73-ambiguous-grew`), Phases A–E as declared in `overview.md`.

---

## Phase A — attribution: is 12 → 39 reach, or rot?

**H1 holds, and it is exact.** Splitting the ambiguous class by which regex alternative matched the
citation — the only thing iter-73's widening changed — partitions it perfectly:

| ref source | `bare-code` | `md` | `path` | total |
|---|---|---|---|---|
| **ambiguous** | **27** | 0 | **12** | **39** |
| block-pinned | 14 | 0 | 31 | 45 |
| default | 6 | 3 | 54 | 63 |
| no-clone | 0 | 16 | 9 | 25 |

The `path` ambiguous count is **12 — unchanged from iter-71's measurement, to the citation.** All
27 of the growth is the `bare-code` alternative, which **could not be counted at all** before
iter-73 because a bare `<name>.<ext>:N` never reached `resolve()`. **The corpus did not get worse;
the fence started seeing a class it had been blind to.** That is the orchestrator's framing,
confirmed mechanically rather than accepted.

`CHECK-M257x-iter73-ambiguous-grew` asked which of two things the growth was. It is the first.

## Phase A′ — and inside the class, a real fence defect

Attribution answers *where the 27 came from*. It does not say the 39 are all legitimate. Reading
the class, **21 of the 39 sit in one document** — `platform-migration-status.md` — whose citation
region is a **10-row markdown table**. That is not a coincidence about that document; it is a
property of `_block_of`:

> `_block_of` walked to the nearest **blank line**. A markdown table has **no blank lines between
> its rows.** So for every citation inside a table, the "block" was **the entire table**, and every
> sha named in ANY row chose the ref for the citations in EVERY row.

This contradicts a rule **this milestone already derived and already implements next door** — §5
rule 33: *"a pin's scope is the claim's own block — a markdown **CELL** in a table, a wrapped
sentence in prose"*, implemented in `platform_predicate_guard._pin_window` since iter-63. **Two
guards, two definitions of "block", and only one of them matched the rule the corpus records.**

The instance that settles it, measured before any edit — `external_services.md`'s provider table:

| row | content |
|---|---|
| 539 | `| **AWS Bedrock (EU)** | … `app/internal/askengine/bedrock.go:25` … |` |
| 540 | `| **Mistral (EU)** | … `app/internal/cms/studio/markdownManager.go:19` … |` |
| **542** | `| **Anthropic Direct** | … measured at app `b948604` … |` |

Rows 539 and 540 carry **no pin**. Both were being read at `b948604` — a ref named by **a fourth
provider's row**. Rule 33's *"the pin crosses a ROW boundary"* mechanism, in the guard that
**chooses refs** rather than the one that grants exemptions, and the consequence is strictly worse:
an exemption merely silences a check, while a mis-chosen ref makes the guard **adjudicate a
citation against a file the document never named**.

## Phase B — the fix, dry-run before it was landed

iter-73's lesson 1 applied first (*twelve findings means land-and-repair; two hundred means
measure-and-route*). Simulating the cell-scoped window with the guard **untouched**:

| ref source | before | after (predicted) |
|---|---|---|
| ambiguous | 39 | **24** |
| block-pinned | 45 | 48 |
| default | 63 | 75 |
| no-clone | 30 | 30 |
| **findings** | **0** | **0** |

18 citations change ref source; 5 change the **file actually read**. Zero verdict changes — so the
whole class was *land-and-repair with nothing to repair*, and that was known before the edit.

Landed as `_block_of(lines, i, col)`: **a CELL in a table, the blank-line block in prose**, with a
table row also a boundary for the prose walk. The prose half is untouched — narrowing it would
re-introduce the one-line window that has been a bug three times in this milestone (iter-63,
iter-68, and iter-71's *surviving* mutant). `col` is the citation's own offset; absent, a table row
falls back to the whole ROW, the conservative direction (a wider window can only make a citation
`ambiguous`, never make it read at a ref nobody wrote down).

**`col=None` is passed deliberately for self-references**, and the comment says why: those offsets
are matched against a *substituted* copy of the line (markdown links stripped), so `m.start()`
indexes a string that is not `line` and would cut the wrong cell.

Live, the prediction reproduced exactly: **ambiguous 39 → 24, block-pinned 45 → 48, default 63 →
75, 0 findings, guard OK.**

## Phase B′ — the fix's own reach limit, found by its own residual

Reading the 24 residual, three of them diverged in a way the class was not supposed to contain —
and two were in a **table that the row test did not recognise**:

```
> | side | measured at platform `0dab54d` / app `9d00a313` v1.367.0 |
> |---|---|
> | **consumer** | … `app/main.go:471` … `app/internal/storage/service.go:22` … |
```

`storage.md`'s v9.0 fold block is a **blockquoted** table. `_TABLE_ROW` (copied verbatim from the
sibling guard) anchors at `|`, so `> |` fails the row test, drops into the prose branch, and takes
**the whole quoted table** as its window — *the exact defect being fixed, surviving inside the first
draft of the fix for it.*

Not a one-document curiosity, and the count is why the prefix was admitted rather than the
document: **75 blockquoted table rows across 14 files** (`rg -c '^\s*>\s*\|.+\|\s*$'`, against 2197
plain rows in 86 files). `_TABLE_ROW` now admits an optional `>` prefix.

Final live reading: **ambiguous 39 → 20 · block-pinned 45 → 45 · default 63 → 82 · no-clone 30 ·
177 resolved · 0 findings.**

## Phase C — adjudicating the residual 20

An ambiguity is only worth anything if it can change an answer. Measured per citation, classifying
at **every** live sha its own window names as well as at the default:

| | count |
|---|---|
| ambiguous citations whose verdict **AGREES** across every candidate ref | **19** |
| ambiguous citations whose verdict **DIVERGES** | **1** |

The 19 are inert: contrast blocks (`deleted by X + Y, merged Z, now W`) where the cited line is a
construct at all of them. The single divergent one is `platform-migration-status.md:71` →
`app/internal/jobsimwiring/wiring.go:123`, whose cell names **three** live shas — and reading the
cell settles it without any rule at all: it says *"`app` @ `9d00a313` v1.367.0 — re-resolved M257x
iter-68; the first four stood at … at `b948604` v1.366.0 and at … `5ba17044` v1.363.2, which is
what one working day costs a line-number citation."* The cell is **contrasting** the three
deliberately, and it names the one it is asserting. The default ladder reads it at `origin/main` =
`9d00a313`, which is the ref the sentence names.

**No corpus repair.** Picking a sha out of a contrast block by rule would be §4 Trap A — a rule
fitted to a sentence — and `block_ref`'s docstring already commits to falling back and **counting**
instead. The residual is the designed fallback, now measured to be inert in 19 of 20 cases and
adjudicated by reading in the 20th.

## Phase D — gates

| gate | result |
|---|---|
| `anchor_construct_guard` | **OK** — 177 resolved / 112 files; `default x82, block-pinned x45, no-clone x30, ambiguous x20` |
| `platform_alignment_guard` | **OK** — F 74 citations (20 subject-checked · 53 range-only · 1 outside · 0 unresolvable), unchanged |
| `platform_predicate_guard` | **OK** — the corpus and the platform's configuration agree |
| `markdown_structure_guard` | **OK** — no structural damage |
| `corpus_index_guard` | **OK** — 84 docs / 6 index-bearing dirs |
| `CITE_REF=worktree` | **still discriminates** — `override x175`, **7** findings. The escape hatch survived a change to the window it overrides |
| `tests/test_iter45_mechanical_fences.py` | **68** (was 62); new class `InATableTheBlockIsTheCell` 6/6 |
| mutation battery | **5 mutants caught · no-op control SURVIVED** |
| `stack-core` suite | **775 tests, 1F in 693.9 s** — `test_claim_twin_guard_iter48_answer_key::test_02_the_green_twin_of_every_site_stays_SILENT`, the perishable iter-48 fixture. **Baseline 769/1F matched by IDENTITY**, +6 = exactly this iter's new tests |
| `stack-injection` · `dev-stack` · `demo-stack` | untouched sections; iter-71/73's runs stand (332 OK · 151 OK **solo** · 1048/7F by identity) |

### The mutation battery

Every mutant is an **inversion of one clause of the landed rule**, and the control changes
**docstring prose only** — the only kind that cannot alter behaviour by construction, which is
iter-73's lesson 2 (*if the control fails, the battery has not run*). The guard is restored
**byte-identical** in a `finally`, asserted.

| mutant | caught |
|---|---|
| M1 table branch removed (the pre-iter-74 blank-line window) | **3F** |
| M2 `col` ignored inside a table (whole ROW, never the cell) | **1F** |
| M3 `>?` dropped — a blockquoted row stops being a row | **1F** |
| M4 a table row stops being a boundary for the prose walk | **1F** |
| M5 `run()` stops passing the citation's column | **1F** |
| **no-op control** (docstring prose) | **SURVIVED — OK** |

Every fixture is a shape the corpus actually writes (a provider table with the pin in another
provider's row; a one-row/two-clause evidence cell; a blockquoted `> |` fold table; a paragraph
butted against a table). Three iterations of this milestone have lost a mutant to a fixture that
agreed with the implementation instead of with the corpus; these were derived from the four
documents named above.

## Phase D′ — P3 at close, and the adjudication ref moved *during* the iteration

| pin | at open | at close |
|---|---|---|
| platform clone | `0dab54d` | `0dab54d` — **level with origin/main** |
| `app` clone | `b948604` v1.366.0 | unchanged |
| `app` **origin/main** | `9d00a313` v1.367.0 | **`7177374` = `v1.367.0-4-g717737471`** |
| rext pin (`.agentspace/rext.tag`) | `fast-build-m257x-iter-67` | unchanged — this iter touched offline guard code only |

**The adjudication ref moved without any clone advancing.** The `auto` ladder prefers `origin/main`,
so a `git fetch` run as a routine P3 check silently re-pointed **every default-adjudicated citation**
(82 of them) at a newer file — §5 rule 26's *"an input that can change without appearing in a diff is
not a controlled input"*, in its purest form. Both ref-aware guards were re-run at the new ref:
**`anchor_construct_guard` OK (177 resolved, `ambiguous 20`, unchanged)** and
**`platform_predicate_guard` OK, consumer side measured `@ origin/main@7177374`.**

Then the citation delta the advance actually cost (`D-M257x-59-3` / §7 rule 4's second half), over
every corpus citation landing in the `app` clone, `9d00a313` → `7177374`:

| | |
|---|---|
| **HELD** (cited line byte-identical) | **48** |
| **MOVED** | **1** |
| **DEAD** | **0** |

And the one "MOVED" **is not a defect** — it is `clerk-integration.md:103`'s `` `app/go.mod:31` @
`5ba17044` ``, a claim that names its own ref. At `5ba17044` line 31 is
`github.com/clerk/clerk-sdk-go/v2 v2.7.0`, exactly as written; it is `go-humanize` at `9d00a313` and
`mimetype` at `7177374`. The guard reads it at its pin (`block-pinned`) and is right to.
**The delta script — written today — forced both refs and manufactured a move that isn't one**, which
is iter-69's lesson reappearing inside a measurement built the same day: *a pin is a DATE; repairing
a dated claim onto a floating ref moves a CORRECT claim onto something that will move again.* Net
real breakage across a 4-commit app advance: **zero**.

## Phase A″ — a hand-off number that did not survive re-derivation (recorded, not repaired here)

The orchestrator handed forward two numbers. **39 reproduces exactly.** The other one splits:

- **"92 bare citations still unresolvable" — the COUNT reproduces** (91 distinct bare-code
  citations across 103 sites).
- **The HEAD LIST does not.** iter-73 wrote *"`gen.py` x10, `intelligence.go` x8, `main.go` x7"*;
  measured now the heads are **`up-injected.sh` x32, `intelligence.go` x5, `20260722104506.sql` x5,
  `main.go` x2 — and `gen.py` x0.** `gen.py` is cited **eleven times** in the corpus and **not once**
  in `file:N` form: every one is a RANGE (`` `gen.py:484-492` ``), which the guard's regex cannot
  match because it requires a closing backtick immediately after the digits.

So the count and the heads came from **two different instruments**, and only the count was measured
with the regex it was attributed to. Routed to iter-75 with the derivation, so the class that gets
repaired is the one that was measured. **Also derived and worth carrying:** of the 239 sites the
guard reports "unresolvable", **88 are URLs** (`http://backend:8083` and the `ENV=http://host:port`
form) — not citations at all. The unresolvable denominator has never been net of them.

### An instrumentation defect in the gate run itself

The suite's first run was invoked as `python3 -m unittest discover -s tests -q 2>&1 | tail -14`, and
**the tail window did not contain the verdict.** unittest writes `Ran N tests` / `OK|FAILED` to
stderr as it finishes, while the guard subprocesses it spawns write to a **block-buffered** pipe
that flushes *after* — so the last 14 lines were guard chatter and the two lines the gate exists to
read had been pushed out of the window. Worse, the pipeline's exit code is `tail`'s, so it reported
**0 over a suite that exited 1**.

Same family as §5 rule 1 (*never let a search's stderr go unread*), one step along: **a truncating
filter can discard the verdict while the command still looks like it succeeded.** Re-run with the
full output captured to a file — which is the only reason the 775/1F below is a measurement rather
than an impression.

## Close — 2026-08-04

**Outcome:** `CHECK-M257x-iter73-ambiguous-grew` is settled, and it is settled by **partition rather
than by judgement**: the `ambiguous` bucket's 12 → 39 is **100% the newly-reachable `bare-code`
alternative** (27), while the pre-existing `path` partition is **still 12, to the citation** — the
fence started seeing a class it had been blind to, and the corpus did not move by one site. Reading
the class anyway then found a real defect *inside* it: **`_block_of` walked to the nearest blank
line, and a markdown table has none between its rows**, so a citation inside a table took the WHOLE
TABLE as its window and every sha named in any row chose the ref for the citations in every row —
contradicting **§5 rule 33**, which this milestone derived in iter-63 and has implemented in the
sibling guard ever since. Measured instance: `external_services.md`'s AWS Bedrock and Mistral rows
were being read at a ref named by the **Anthropic Direct** row. The window is now the CELL in a
table and the blank-line block in prose — and **the first draft of that fix contained the defect it
removes**, because the row test (copied verbatim from the sibling) anchors at `|` and `storage.md`'s
fold block is a **blockquoted** table, one of **75 such rows across 14 files**. Final: **ambiguous
39 → 20, block-pinned 45, default 82, 177 resolved, 0 findings** — nothing in the corpus was wrong;
**19 citations stopped being adjudicated against a file their own document never named.** The
residual 20 was then adjudicated rather than assumed: classified at **every** ref its own window
names, **19 agree and 1 diverges**, and the 1 is a cell that names in words the ref it asserts.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (1 tik of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-74-1` (the growth is reach, proved by splitting the class along the
dimension reach moved), `D-M257x-74-2` (in a table the block is the CELL — the two guards now agree),
`D-M257x-74-3` (a blockquoted table row is still a table row; the first draft of the fix contained
the defect it removes), `D-M257x-74-4` (the residual 20 adjudicated by reading — 19 inert, 1
self-naming — and NOT closed by a fitted rule), `D-M257x-74-5` (a hand-off whose count was measured
and whose head list was not), `D-M257x-74-6` (a routine `git fetch` moved the adjudication ref
mid-iteration; 48 held / 1 correctly-pinned / 0 dead).
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter73-unresolvable-92` — **re-derived and re-headed** (91 distinct bare-code
  citations / 103 sites; heads `up-injected.sh` x32, `intelligence.go` x5, `20260722104506.sql` x5,
  `user_resource.go` x4, `ensure-clones.sh` x4). Carries `D-M257x-74-5`'s correction: the routed
  head list came from a different instrument than the routed count, and **88 of the 239
  "unresolvable" sites are URLs, not citations**. iter-75's target.
- **Closed here:** `CHECK-M257x-iter73-ambiguous-grew` (and with it the narrower
  `CHECK-M257x-iter71-ambiguous-blocks` it superseded).
- Unchanged: `FENCE-M257x-iter70-line-or-port` · `RF-M257x-iter71-run-returns-a-tuple` ·
  `CHECK-M257x-iter70-studio-room-lines` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**)
  · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.

**Lessons:**

1. **When a count grows in the same pass that REACH grows, split it by the reach dimension before
   touching anything.** It is one derivation and it decides between two opposite responses. Rule 16
   says an unread metric is indistinguishable from an unmoved one; this is its mirror — **an unread
   metric that becomes read looks exactly like a metric that got worse.** Landed as §5 rule 35.
2. **When a rule is already derived, check every implementation of it, not the one you are
   editing.** Rule 33 had ruled on the table window in iter-63 and one of the two guards that needed
   it never got it. Nobody notices, because each guard is internally consistent and both are green;
   the wrong one is wrong *silently*.
3. **An ambiguity is only worth repairing if it can change an answer.** 19 of the 20 residual
   citations classify identically at every ref their own window names. Counting a fallback is
   coverage; assuming it is a defect is not.
4. **The first draft of a fix is the most likely place to find the defect it removes.** Twice in
   this milestone now (iter-68's default argument, this iter's `|`-anchored row test). Re-read the
   fix as if it were the code under audit, before the battery.
5. **A pin is still a date, even inside a measurement script you wrote today.** The app-delta script
   forced two refs and manufactured a "MOVED" out of a claim that is true at the ref it names.
6. **A `| tail -N` on a suite can hide the verdict, and the pipeline's exit code then belongs to
   `tail`.** Capture a gate's full output to a file. `-q` plus block-buffered subprocess stdout is
   enough to push `Ran N tests` / `FAILED` out of any fixed window.

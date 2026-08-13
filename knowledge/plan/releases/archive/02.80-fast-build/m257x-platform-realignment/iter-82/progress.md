# iter-82 — progress

**Type:** tik, under `TOK-05`. Deliverable: **the re-read**. No repair landed, by instruction and by
design — a repair validated by the pass that found the work is the defect this milestone exists to end.

---

## THE RESULT

| | pre-repair (#13 / #14, iter-76) | **post-repair (#15 / #16, this iter)** |
|---|---|---|
| reading 1 | **77** | **29** |
| reading 2 | **75** | **30** |
| union (distinct anchors) | 152 booked → 140 upheld | **41** |

**`N₁₅ = 29`, `N₁₆ = 30`.** Per-seat, re-derived from the reports on disk (not from any seat's
own summary line):

| seat | #15 | #16 |
|---|---|---|
| A | 3 | 4 |
| B | 3 | 6 |
| C | 4 | 5 |
| D | 4 | 3 |
| E | 6 | 6 |
| F | 3 | 2 |
| G (diff) | 6 | 4 |
| **total** | **29** | **30** |

Arithmetic control: 59 booked findings parsed across the 14 reports; 29 + 30 = 59. Every report's
counted `### B<n>` headings match; every report carries its per-file `wc -l` positive controls.

**Clause 5 is graded ONLY by a reading that returns zero. It did not. The gate stays 4 of 5.**

The 140 → 41 movement is real and large, and it is **not** a claim that 99 defects were fixed: the
two numbers are a *booked union* and an *unadjudicated union* respectively. The pre-repair 152 lost
12 to adjudication (92.1 % upheld). This 41 has had **no adjudication at all** — that is the next
iter's job, and iter-80 made *"adjudicate before repairing"* binding precisely so this number is not
mistaken for a work list.

## The instrument — frozen, and PROVEN so

Not asserted this time. Verified on every knob, and one knob was verified by re-execution:

- **briefing byte-identical** — sha256 `3858ec53…`; `git log` on the path returns **one** commit
  (`012edd2`, iter-76). Untouched since it was taken.
- **file set identical** — the same 40 files (`corpus/architecture/*.md` + `corpus/services/*.md`).
- **seat count identical** — 7 per reading (A–F full-read + G adversarial diff), 2 readings.
- **ground truth identical** — all **12** clone shas re-derived at open and matched the briefing's
  table exactly. No clone had moved since iter-76.
- **partition method re-executed and reproduced iter-76's hand EXACTLY** — run at `012edd2` it
  returned 40 files / 9,544 lines with seat totals 1791 / 1621 / 1537 / 1495 / 1542 / 1558 and
  per-seat file lists identical to the recorded partition. That is the positive control for the
  method itself, and it is the thing §5 rule 25 says must be storable rather than describable.

**What moved, and why it had to:** iter-81's repair changed line counts (+410 / −251 over 33 files),
so the corpus is now 40 files / **9,712** lines and the fixed method **deals a different hand**. Same
consequence iter-76 recorded at #11/#12; recorded, not engineered away. Freezing the *output* of the
method instead of the *method* is the drift, not the cure for it.

## The pre-registered predictions, graded unsoftened

| # | prediction | outcome |
|---|---|---|
| 1 | neither reading returns zero | **HOLDS** — 29 and 30 |
| 2 | each of `N₁₅`, `N₁₆` in **[10, 45]** | **HOLDS** — 29 and 30 |
| 3 | union > max(N₁₅, N₁₆) | **HOLDS** — 41 distinct vs max 30 |
| 4 | seat G books ≥ 1 (the repair introduces defects) | **HOLDS** — G booked 6 / 4 |
| 5 | the 11 repaired predicates contribute **zero** | **FALSIFIED** — see below. This is the iter's headline, as pre-registered |
| 6 | the held storage carve-out is re-booked | **FALSIFIED** — see below |

### Prediction 5 falsified: the iter-81 repair is INCOMPLETE

Pre-registered wording: *"If any seat books a finding inside a predicate iter-81 claimed to repair,
**the repair was incomplete** and that is this iter's headline, not a footnote."*

`corpus/services/graphql-wundergraph.md:13`, booked by **both** readings (seat B in each):

> "The `graphql` *profile name* survives in compose and is now simply the default profile — it no
> longer names a router service."

Measured at `platform 0dab54d`, the ground-truth ref: the token `graphql` appears in **no**
`profiles:` key. The eight that exist are `core`, `backend`, `all`, `storage-legacy`,
`customerio-sync`, `messenger`, `studio-desk`, `frontend`; the Makefile reads `PROFILE ?= core`. The
name does **not** survive in compose, and the default profile is **not** `graphql`. Both halves of
the sentence are false.

That site falls squarely inside **P4 — *"`graphql` is a live profile / the default"*, ~10 sites**,
one of the 11 predicates iter-81 reports as repaired. **One reading finding it would be a miss; both
readings finding it independently is a hole in the repair**, and it is the exact predicate whose
danger the root `CLAUDE.md` documents at length (asking for a retired profile token *exits 0* and
starts three containers, so the stack looks alive and the application is absent).

**Adjudication of the remaining 40 is still deferred** — this one is graded here only because it
decides a pre-registered prediction and cost one command.

**Prediction 6 falsified, and this is load-bearing for the report:** `storage.md` `:55` / `:154` /
`:181` — the `/tmp`-sandbox-vs-production-bucket contradiction held by instruction
(`DEF-M257x-iter80-storage-prod-bucket`, escalated, awaiting the user) — was booked by **neither
reading**. The only `storage.md` anchor in the union is `:8`. **The held carve-out accounts for zero
of the 41.** It cannot be offered as an explanation for any part of a non-zero N.

**Recall is still the dominant term.** The two readings share only **15** of 41 distinct anchors —
overlap ≈ 51 % / 55 %, unchanged from the <60 % prior that has held across every paired measurement
in this milestone. This was the stated reason prediction 1 was written against the gate's interest,
and it is why a zero was never the honest expectation: **repairing the union of two readings cannot
repair what neither reading saw**, and a fresh pair draws a fresh sample from the same remainder.

## Composition of the 41

Per corpus file (all 14 seats, 59 booked):

| n | file |
|---|---|
| 8 | `ai-readiness.md` |
| 6 | `platform-migration-status.md` |
| 5 | `external_services.md` |
| 4 | `backend.md` · `dependency_map.md` |
| 3 | `academy-backend.md` · `alignment_testing.md` · `architecture_overview.md` |
| 2 | `graphql-wundergraph.md` · `cms.md` · `storage.md` · `hiring.md` · `shared_libraries.md` · `skillpath.md` |
| 1 | `gotenberg.md` · `ant-academy.md` · `jobsimulation.md` · `services/README.md` · `CLAUDE.md` · `roadrunner.md` · `service_taxonomy.md` · `frontend_architecture.md` |

**The dominant class is the stale line anchor**, and it is overwhelmingly **cross-repo** — corpus
citations into `rosetta-extensions` and into `app` whose target has moved by a few lines. The
clearest cluster: `ai-readiness.md`'s rext anchors, booked independently by seats B and G in **both**
readings, consistently drifted by **+4**. This is the class `CHECK-M257x-iter77-cross-repo-pin`
flagged as unmeasured — **it is now measured, and it is the largest single contributor.**

Anchors booked by **both** readings (the highest-confidence set, 15):
`alignment_testing.md:360` · `architecture_overview.md:74` · `dependency_map.md:19` ·
`dependency_map.md:58` · `external_services.md:208` · `platform-migration-status.md:71` ·
`shared_libraries.md:181` · `academy-backend.md:15` · `ai-readiness.md:213` · `ai-readiness.md:219` ·
`backend.md:43` · `graphql-wundergraph.md:13` · `skillpath.md:34` · `storage.md:8`

## One finding adjudicated in-run, because it was cheap and it was mine to check

Three seats (C#16, D#16, and C#15 as an out-of-scope note) booked or flagged that
`service_desired_count` reads **`= 1`** at `:19` in both `storage` and `messenger`, contradicting the
map's `= 0` at `storage/terraform/main.tf:38` and `messenger/terraform/main.tf:29`.

**It is a FALSE POSITIVE, verified directly rather than by taking a seat's word:**

```
storage   63bffc8:terraform/main.tf:38   service_desired_count = 0
messenger a0ec933:terraform/main.tf:29   service_desired_count = 0
```

Both cited refs resolve in the clones; both claims are exact at the ref the claim itself names, with
the ordering rationale in-comment. The clones sit at older shas (`4ce8ece5` / `fa47850d`) where the
value genuinely is `1` at `:19`. This is **`CHECK-M257x-iter76-seat-ref-discipline`** recurring —
seats grading a ref-pinned claim against the checkout despite a briefing section telling them not to.
**Third occurrence. A rule stated in a briefing is still not a rule enforced by an instrument.**

## Escalation condition — NOT triggered

Written at open: *"if the union exceeds 60, the loop is not converging and the next move is a TOK."*
Union is **41**. The condition does not fire; the repair-then-re-read loop is converging (140 → 41).

## A bookkeeping gap found while closing: iter-81 left NO iteration record

`iter-81/` contains an empty `raw/` and nothing else — **no `overview.md`, no `progress.md`, no
`decisions.md`** — and the milestone `progress.md` has **zero** occurrences of "iter-81". A 33-file,
+410/−251 repair that the gate turns on is recorded in the plan **nowhere**.

**The 11 predicates exist, as a written list, only inside the iter-81 commit message** (a 58-line
message; the list is lines 19–29). That has two consequences worth stating plainly:

1. **Adjudicating "did the repair miss any of the 11" requires the 11** — and they are recoverable
   only from `git log`, not from any plan artifact.
2. It is the same class as `CHECK-M257x-iter82-commit-message-narration` in the other direction: the
   commit message is doing documentation's job, and this run has already measured that *that
   particular message* contains a loose anchor (`:17-19`). **A commit message is testimony, not
   evidence** — and it is now also the sole record of a gate-critical work item.

Not repaired here (writing iter-81's record retrospectively would be authoring history, not
recovering it). Routed as **`FIX-M257x-iter82-iter81-has-no-record`**.

## Routed forward — NOT actioned here

- **`FIX-M257x-iter82-reread-union`** — adjudicate the 41 before any repair. Binding per iter-80.
- **`CHECK-M257x-iter77-cross-repo-pin`** — no longer unmeasured: it is the largest class in the 41.
- **`CHECK-M257x-iter76-seat-ref-discipline`** — third occurrence, now with a worked example.
- **`DEF-M257x-iter80-storage-prod-bucket`** — still held, still escalated, **and provably not part
  of the 41**.

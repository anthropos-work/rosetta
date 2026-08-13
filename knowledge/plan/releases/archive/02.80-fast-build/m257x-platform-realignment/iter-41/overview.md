---
iter: 41
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-02
---

# iter-41 — `MEASURE-M257x-iter41-clause5-sixth-pass`

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`).

## Step 0 — re-survey

Clause 5 is the only open gate clause (1, 2, 3, 4 all hold). Five passes have run — iters 21/33, 34, 38, 39
— returning **25, 13, 11, 17, 37** blockers. iter-39 established that **this series measures the
instruments, not the corpus**: five different auditor counts and briefings (5 → 5 → 6 → 7), each better
informed than the last. Nothing in it licenses a claim of convergence *or* of growth.

## The design property that makes this pass different — and it is the whole point

**Verified at open:** `git diff <iter-39 close>..HEAD -- corpus/services/ corpus/architecture/` is **EMPTY**.
iter-40's claim-scoped repair touched `corpus/ops/**`, `.claude/skills/**`, `corpus/README.md` and
`CLAUDE.md` — **not one of the 40 in-scope files.**

So pass six reads a corpus **byte-identical** to the one pass five finished repairing, with an instrument
**held fixed**. That has never been true before: every prior pass differed from its predecessor in *both*
the corpus (a repair landed between them) and the instrument (more auditors, better briefing). The two
variables were perfectly confounded for five consecutive measurements.

**This makes the result interpretable in a way no previous pass was:**

| pass six returns | what it means |
|---|---|
| **~0** | The corpus is clean and pass five's repairs held. Clause 5 closes. |
| **large** | On a byte-identical corpus with a fixed instrument, a large count means **the measurement is not repeatable** — the residual is a function of *who reads*, not of what is written. |

The second outcome is not "the corpus still has defects." It is *"clause 5 asks for a property that this
instrument cannot measure twice."* That is a materially different — and more useful — finding, and it is the
one that would justify escalating rather than opening a seventh pass.

## Instrument — HELD FIXED at iter-39's, deliberately

Run 21's hand-off advised re-partitioning again. **Overruled, with the reason stated**: iter-39's central
finding is that the headline series measured instrument changes. A sixth number produced by a seventh
instrument would extend that series and settle nothing. **Comparability is this pass's deliverable**, so
every knob is held:

- **Seven auditors**: six full-read partitions (A–F) + one adversarial diff-reader (G). Same as iter-39.
- **Same briefing**: every §5 rule, each file's own repair history, the same blocker/minor grading rule, the
  same ground-truth sources.
- **Same partition METHOD**: files sorted by line count descending, snake-dealt A→F then F→A.
- **All 40 files read in full, top-to-bottom.** No narrowing. Both narrowing strategies are disproven in
  this milestone (by search term, iter-21; by file set, iter-38).

> **Note on rule 18(b) — the tension is real and is resolved by the method, not by ignoring it.** Rule 18(b)
> says partition a confirming pass differently. Holding the *method* fixed does not hold the *partition*
> fixed: iter-39's repairs moved 20 files' line counts by +739/−250, so the same size-sort deals a
> **different hand**. Checked, and the properties iter-39 relied on all survive: `ai_architecture` (A) and
> `security_compliance` (E) split; `backend` (B) / `architecture_overview` (C) / `service_taxonomy` (F)
> split three ways; `studio-desk` (F) / `studio-room` (D) split; `academy-backend` (A) / `ant-academy` (E)
> split; the merged/archived family lands in **five** different hands. The method delivers both properties
> at once — which is why it is the right thing to hold fixed.

### Partition (computed, 40 files / 9,163 lines)

| auditor | lines | files |
|---|---|---|
| **A** | 1716 | external_services · ai_architecture · graphql-wundergraph · academy-backend · coursebuilder · skiller · TEMPLATE |
| **B** | 1541 | ai-readiness · backend · cms · ai-labs · clerk-integration · customerio-sync · architecture/README |
| **C** | 1505 | alignment_testing · architecture_overview · chronos · roadrunner · messenger · services/README · db-backup |
| **D** | 1472 | studio-room · clerkenstein · shared_libraries · sentinel · next-web-app · gotenberg · intelligence |
| **E** | 1443 | ant-academy · hiring · security_compliance · storage · askengine · dependency_map |
| **F** | 1486 | studio-desk · service_taxonomy · jobsimulation · platform-migration-status · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of `d5bd838..b925199 -- corpus/` — iter-39's own 20-file sweep, +739/−250 |

## PRE-REGISTERED PREDICTION — written before any auditor reports

**Count: 8–20 blockers.**

Derivation, and its weakness stated up front. The one stable quantity in the series is *"blockers found in
the immediately prior pass's repaired text"*: **9** (iter-34), **11** (iter-38), **~25** (iter-39). It has
risen every time, and iter-39 repaired the most files yet (20). Arguing *down* from ~25 rests entirely on
iter-39 having run the most thorough adversarial half of the series (8 self-inflicted found and fixed,
against 6 and 2 before it) — which is a real difference but an unquantified one. **The lower bound is soft
and the whole interval may be refuted, as the last two were.**

**Named specific predictions** (so the pass can miss them):

1. At least one blocker in text iter-39 wrote to *explain* a correction — the retraction-as-new-claim shape
   that has appeared in every adversarial pass.
2. **Fewer than 5 blockers in the files iter-39 never edited** (20 of 40). If this fails, the repaired-text
   density model is refuted outright rather than merely non-predictive.

**Consequent prediction: clause 5 does NOT close.** Registering that is what makes a zero result meaningful
if it happens, and what stops a non-zero result being re-framed as a disappointment.

## Phase plan

- **Phase A — read.** Seven auditors in parallel. Each verifies every falsifiable claim against the platform
  clones at origin `2adcf71`, the live `demo-1` Postgres, or `docker-compose.yml`/`repos.yml`. **Blocker** =
  a claim a reader would act on that is FALSE, or a load-bearing `file:line` anchor that does not resolve.
  Line drift / undercounts / omitted list members are **minors** (*"YELLOW with 0 blockers"* admits them).
- **Phase B — adjudicate.** Enumerate to `blocker-ledger.md`. **Verify each reported blocker before
  accepting it** (§5: verify a claim before escalating it, including a claim made by an audit — two of
  iter-22's handed corrections were themselves false).
- **Phase C — grade clause 5 against the READING.** A clause is met by a reading that returns zero, never by
  a repair that clears its own findings. iters 33, 34, 38 and 39 each refused to claim it on that ground.

## Escalation conditions — pre-committed, so the outcome cannot be re-litigated after the fact

- **Non-zero result → STOP. Do NOT open a seventh pass.** Close the iter honestly and exit
  `EXIT_REASON: user-blocker`, reporting (a) the count, (b) how much is repair-induced versus genuine, and
  (c) a recommendation on `CHECK-M257x-iter33-derived-fact-fence`. At that point the open question stops
  being *"fix the findings"* and becomes **whether a hand-maintained corpus can satisfy a zero-blocker
  clause at all** — which asks whether the gate clause needs re-cutting, and that is the user's call, not
  this iteration's.
- A platform commit landing mid-iter → **re-scope trigger occurrence 2** → STOP.
- A legal question (the AI-Act classification) → route to an owner outside this milestone; do not settle it.

## Acceptable close-no-lift outcomes

**The reading is the deliverable.** A pass returning non-zero is a *complete* iter — and under this pass's
design, a large non-zero is the more informative of the two outcomes, because it measures the instrument's
repeatability rather than the corpus's residual. Repairs are **out of scope for this iter**: fixing findings
here would destroy the byte-identical property that makes the number mean anything.

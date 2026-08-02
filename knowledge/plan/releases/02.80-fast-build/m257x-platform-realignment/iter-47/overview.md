---
iter: 47
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-02
---

# iter-47 — `READ-M257x-iter41-instrument`: the seventh pass, instrument frozen

**Active strategy reference:** [`TOK-02`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02) — **step 5**, the last, verbatim: *"Then ONE full 7-auditor read,
instrument held fixed at iter-41's — same seven auditors, same briefing, same partition method, all 40
files top-to-bottom. That reading is what meets or fails clause 5. Nothing else does."*

## Step 0 — re-survey before targeting

Platform origin HEAD re-fetched at open: `2adcf71`, **unchanged** (re-scope trigger stays at occurrence 1
of 2). Steps 1–4 are complete: four fences exist, run at commit time, and report **0 sites** on a tree
whose 18 known blockers were repaired at iter-46 by claim. TOK-02's named target for this slot is
unchanged and this is the only step left.

## What this iteration is, and what it is not

**It is a MEASUREMENT.** It repairs nothing. Clause 5 is met by *a reading that returns zero* — never by a
repair that clears its own findings, a rule iters 33, 34, 38, 39 and 41 each refused to break.

**Four fences GREEN is not this number.** Reading it as one is the mistake iter-38 measured (narrowing to
high-density files would have found 11 of 17) and iter-21 measured before it (a term-scoped `11→5→2`
preceded a full read finding 53). The fences check three mechanical classes and one claim ledger; a
7-auditor full read checks everything else.

## Instrument — HELD FIXED at iter-41's, deliberately

Every knob identical to iter-41 (which was itself identical to iter-39):

- **Seven auditors** — six full-read partitions (A–F) + one adversarial diff-reader (G).
- **Same briefing** — every §5 rule, each file's repair history, the same blocker/minor grading rule.
- **Same partition METHOD** — files sorted by line count descending, snake-dealt A→F then F→A.
- **All 40 files read in full, top-to-bottom.** No narrowing. Both narrowing strategies are disproven in
  this milestone (by search term, iter-21; by file set, iter-38).

**Rule 18(b) is satisfied by the method, not by ignoring it.** Holding the *method* fixed does not hold
the *partition* fixed: iter-46's repair moved the corpus from 9,163 to **9,243 lines**, so the same
size-sort deals a **different hand**. Checked against iter-41's own list, and the separations it relied on
survive or improve: `ai_architecture` (A→**B**) and `security_compliance` (E→**B**) now share a hand while
`external_services` (A) and `architecture_overview` (C) split from both; `backend` moved B→**A**, `cms`
B→**C**, `sentinel` D→**C**, `service_taxonomy` F→**E**, `platform-migration-status` stays **F**. **20 of
the 40 files changed hands.**

### Partition (computed, 40 files / 9,243 lines)

| auditor | lines | files |
|---|---|---|
| **A** | 1726 | external_services · backend · graphql-wundergraph · academy-backend · coursebuilder · skiller · TEMPLATE |
| **B** | 1559 | ai-readiness · ai_architecture · security_compliance · ai-labs · clerk-integration · customerio-sync · architecture/README |
| **C** | 1520 | alignment_testing · architecture_overview · cms · sentinel · messenger · services/README · db-backup |
| **D** | 1481 | studio-room · clerkenstein · chronos · roadrunner · next-web-app · gotenberg · intelligence |
| **E** | 1459 | service_taxonomy · hiring · shared_libraries · storage · askengine · dependency_map |
| **F** | 1498 | ant-academy · studio-desk · jobsimulation · platform-migration-status · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of `29eb414..301d61a -- corpus/ CLAUDE.md` — iter-46's own 17-file repair |

## PRE-REGISTERED PREDICTION — written before any auditor reports

**Count: fewer than 9 blockers.** This is TOK-02's own pre-registered prediction, carried here unmodified
so it can be refuted. Its derivation: iter-41 measured 18, of which **9 were manufactured by the repair
that preceded them, 8 of the 9 in the single self-contradiction class**. TOK-02 attacks that induced term
specifically — the fence now runs at the commit — so the projection is the ~9 non-induced residual, minus
the 5 that iter-45's mechanical fences now also cover at commit time, plus whatever iter-46's repair
induced that no fence sees.

**Named specific predictions, so the pass can miss them:**

1. **Fewer than 4 blockers of the self-contradiction class.** This is the direct test of TOK-02. If the
   class returns at iter-41's rate, the commit-time fence has not touched the mechanism and the strategy
   is refuted, not merely under-delivering.
2. **At least one blocker in text iter-46 wrote to *explain* a correction** — the retraction-as-new-claim
   shape that has appeared in every adversarial pass, and iter-46 wrote a great deal of such text.
3. **The residual is NOT concentrated in the 17 files iter-46 edited.** Every prior pass measured a 4–9×
   density ratio in repaired text. If that holds again, the induced term survives the fence.

**Consequent prediction: clause 5 does not close.** Registering that is what makes a zero result
meaningful if it happens, and what stops a non-zero result being re-framed as a disappointment. **Four
consecutive passes refuted their own predictions**, and iter-41's held only once the instrument was fixed.

## Phase plan

- **Phase A — read.** Seven auditors in parallel. Each verifies every falsifiable claim against the
  platform clones in `stack-demo/` at origin `2adcf71`, or `docker-compose.yml`/`repos.yml`. **Blocker** =
  a claim a reader would act on that is FALSE, or a load-bearing `file:line` anchor that does not resolve.
  Line drift / undercounts / omitted list members are **minors** (*"YELLOW with 0 blockers"* admits them).
- **Phase B — adjudicate.** Enumerate to `blocker-ledger.md`. **Verify each reported blocker before
  accepting it** — two of iter-22's handed corrections were themselves false, and iter-01 refuted five
  inherited claims by measurement.
- **Phase C — grade clause 5 against the READING**, and against nothing else.

## Escalation conditions

- **A blocker count at or above iter-41's 18** refutes TOK-02 outright and is a re-scope conversation, not
  a seventh repair.
- **A count in 9–17** means the strategy helped and did not close it; that is a user decision about
  whether clause 5 is reachable, not an eighth pass to be started unilaterally.

## Acceptable close-no-lift outcomes

Any count, honestly derived and adjudicated, is this iteration's deliverable. **A non-zero reading is a
complete iteration**, not a failure — the failure mode this milestone has guarded against six times is
re-cutting the instrument until the number is acceptable.

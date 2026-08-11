---
iter: 43
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: closed-fixed
opened: 2026-08-02
closed: 2026-08-02
handler: FENCE-M257x-iter42-claim-twin
---

# iter-43 — the claim-twin fence, watched going RED

**Active strategy:** [`TOK-02: fence the prose the way the anchors are fenced`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02)
(milestone-root `decisions.md`), approved by the user. **This iteration is step 1 of five, and step 1 only.**

## Step 0 — re-survey (mandatory, measured at open, not inherited)

| what | measured |
|---|---|
| rosetta tree | CLEAN at `48ca53c`, branch `m257x/platform-realignment`, 0 behind `main` |
| the 18-blocker fixture | **INTACT** — `git diff 103ad31..HEAD -- corpus/services/ corpus/architecture/` is **empty** (`103ad31` is iter-41's own closing commit), and four blockers re-verified live in the files: #18 `architecture_overview.md:243`, #9 `service_taxonomy.md:136`, #10/#11 `sentinel.md:12`/`:22`, #12 `graphql-wundergraph.md:79` |
| gate | **4 of 5.** Clauses 1–4 hold; clause 5 at **18** |
| platform origin | `2adcf71`, **unchanged** — re-scope trigger stays at occurrence 1 of 2 |
| rext | `main` @ `069c238`, CLEAN, pin `fast-build-m257x-iter-37`, both pins match |

TOK-02's `Next-tik direction` names this iteration's target explicitly and the re-survey confirms it is
still the right one: nothing has been repaired, so the fixture is still perishable and still available.

## Cluster / target identified

`FENCE-M257x-iter42-claim-twin` — the claim ledger + the fence that reads it, subsuming
`CHECK-M257x-iter33-derived-fact-fence` (which has sat unbuilt since iter-33).

**Why this and nothing else.** TOK-02 measured that **13 of the 18 blockers are the corpus contradicting
itself** and that **8 of the 9 repair-induced blockers are that same single class** — a claim repaired at
one site and left standing at another. That class needs no ground-truth read: the verdict is already
written down, five times over, in the blocker-ledgers of iters 33/34/38/39/41. Nothing in the tree checks
whether an adjudicated claim has come back.

## Hypothesis

A claim ledger **derived** (never hand-assembled — §5 rule 19's list-derivation clause) from the existing
blocker-ledgers, matched tree-wide against a markdown-normalized corpus, will fire at the anchors the
ledgers name — demonstrating detection on a fixture with a known answer key, before that fixture is spent.

## Expected lift

**None on the primary metric, by construction, and that is the point.** The gate metric is the blocker
count; this iteration repairs nothing, so it stays at 18. The iteration's deliverable is an *instrument*
and a *measurement of that instrument*.

**Pre-registered success criterion, so the result can be refuted:**

- The fence detects **≥ 12 of the 18** at the anchors iter-41 recorded. *(Rationale for the floor: 13 are
  self-contradiction and 3 are anchor-shaped; the 2 derived scalars have no quoted refuted form in the
  ledger's claim column and are expected to be OUT of a claim-twin fence's reach — they are step 3's value
  fence, not step 1's.)*
- A **GREEN control** exists and passes: claims from the earlier ledgers whose refuted form has genuinely
  been repaired must produce **no** hit, and the fence must say so as a positive statement of coverage
  rather than by silence (§5 rule 8 — a check that SKIPS reads exactly like a check that PASSES).
- A **mutation battery** with at least one **declared-GREEN no-op mutant that survives** and at least one
  **inverted mutant that goes RED** (§8 rule 5 — removal mutants do not catch inversion).

## Phase plan

- **Phase A** — re-survey + verify the fixture. *(done above)*
- **Phase B** — build the **derivation**: parse the blocker-ledgers into a machine-readable claim ledger.
  The source set is derived from table STRUCTURE, not from a hand-typed list of five filenames.
- **Phase C** — build the **fence**: normalize markdown, match tree-wide (`corpus/**`, `.claude/skills/**`,
  `CLAUDE.md`, `README.md`), report `file:line`, exit non-zero on any unacknowledged hit.
- **Phase D** — **watch it go RED** against the 18, with the GREEN control and the coverage statement.
- **Phase E** — tests + mutation battery; run the section suites; close; commit; tag rext; push the tag;
  verify it on origin.

## Escalation conditions

- A platform commit landing mid-iter → re-scope occurrence 2 → **STOP**.
- Any temptation to repair a corpus sentence → **refuse**. `D-M257x-42-3` is binding: fixing even one of
  the 18 destroys the only fixture that will ever carry a known, anchored, re-verified answer key.
- The fence needing a waiver that suppresses one of the 18 → that is Trap A's tune-until-it-catches-nothing.
  The answer-key test is deliberately the control that makes it impossible.

## Acceptable close-no-lift outcomes

If the derivation proves the ledgers' claim columns are too heterogeneous to yield matchable patterns for
a majority of the 18, **that is a real, publishable falsification of TOK-02's step 1** and the iteration
closes on it with the measurement — not on a fence tuned until it looked good.

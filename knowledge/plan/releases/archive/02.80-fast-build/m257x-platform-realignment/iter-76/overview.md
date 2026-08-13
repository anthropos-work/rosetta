---
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
active_strategy: TOK-05
---

# iter-76 — the graded READ (readings #13 / #14)

**Type:** tik, under [`TOK-05`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04),
whose `Next-tik direction` rung 5 is *"the next paired reading — the first ever taken against a
corpus whose largest predicate classes are fenced rather than prose. Only then is a zero reading
arithmetically reachable."*

## Why now, and not before

iter-73 declined this read and was right to: three known-bad classes were open, and **a reading
taken over a known-unrepaired class measures the instrument, not the corpus.** Both are now closed,
each by adjudication rather than by repair:

| class | closed at | verdict |
|---|---|---|
| **39 ambiguous-block citations** | iter-74 | growth was **100% reach** (`path` partition unchanged at 12); the fence defect underneath it fixed; residual **19 of 20 inert**, the 20th self-naming |
| **92 unresolvable bare citations** | iter-75 | **0 defects** — 77 unreachable (now reachable), 26 undecidable, **0 ABSENT** |

## Scope and instrument — refs stated (TOK-04 P1)

- **Corpus scope: 40 files / 9,544 lines** — `corpus/architecture/*.md` + `corpus/services/*.md`,
  which is exactly clause 5's stated scope. (Was 9,395 at #11/#12; the milestone's repairs since
  have added 149 net lines.)
- **Instrument HELD FIXED** at iter-41's on every knob — seven seats per reading, six full-read
  partitions A–F plus the adversarial diff seat G, the same briefing, the same partition METHOD
  (files sorted by line count descending, snake-dealt A→F then F→A), `wc -l` positive control per
  file, seats blind to each other and to `knowledge/plan/**`.
- **The briefing is STORED, not described** — `instrument/briefing-iter76-AS-RUN.md`. Its only
  deltas from the iter-53 AS-RUN copy are the ground-truth shas and one added paragraph on choosing
  the ref for an `app` claim; both are recorded there rather than applied silently (§5 rule 25).
- **Ground truth, re-derived at open:** `app b948604f` (v1.366.0) · `platform 0dab54df` (**level
  with origin/main**) · `next-web-app bb3313bc` · `sentinel 88bc5592` · `storage 4ce8ece5` ·
  `messenger fa47850d` · `cms ca50c817` · `graphql-wundergraph 60c229f3` · `roadrunner 87d8d443` ·
  `jobsimulation 462343b0` · `studio-desk 14a5442a` · `ant-academy 9c3843cd`.
- **`app`'s `origin/main` is four commits past `v1.367.0`** and the checkout is not. The briefing
  therefore carries the corpus's own rule (§5 rule 33) explicitly: **a claim is settled at the ref
  the claim itself names**, the checkout settles an unpinned claim, and a pin's scope is the claim's
  own block. This is a *stated* instrument change, not a silent one.
- **Seat G's diff scope:** `0e35b1a..HEAD -- corpus/ CLAUDE.md .claude/` — 32 files, 898
  insertions / 257 deletions, i.e. every corpus repair since the last graded read.

### Partition (computed by the fixed method — 40 files / 9,544 lines)

| seat | lines | files |
|---|---|---|
| **A** | 1791 | external_services · security_compliance · backend · messenger · academy-backend · skiller · TEMPLATE |
| **B** | 1621 | ai-readiness · ai_architecture · graphql-wundergraph · ai-labs · coursebuilder · customerio-sync · architecture/README |
| **C** | 1537 | alignment_testing · architecture_overview · platform-migration-status · sentinel · clerk-integration · services/README · db-backup |
| **D** | 1495 | studio-room · clerkenstein · cms · roadrunner · next-web-app · gotenberg · intelligence |
| **E** | 1542 | service_taxonomy · hiring · chronos · storage · askengine · dependency_map |
| **F** | 1558 | ant-academy · studio-desk · shared_libraries · jobsimulation · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of every corpus repair since `0e35b1a`, 32 files |

The method is fixed and therefore deals a **different hand** than #11/#12 (file sizes moved). That
is the honest consequence of a fixed method over a changed corpus, and it is recorded rather than
engineered away.

**Readings #13 and #14 get the IDENTICAL partition** — same hand, same diff. Any disagreement
between them is seat variance and nothing else (§5 rule 23's control).

## PRE-REGISTERED PREDICTIONS — written before any seat was launched, before any report was read

1. **Per-reading count.** Each of `N₁₃`, `N₁₄` lands in **[0, 10]**.
2. **Overlap.** Recall stays **below 60%** for both readings — the 43–48% single-pass figure has
   held across every paired measurement in this milestone and nothing in iters 54–75 was aimed at
   raising it.
3. **Union.** `|#13 ∪ #14| > max(N₁₃, N₁₄)` — each reading books at least one blocker the other
   misses. **This prediction is falsified if either reading returns 0**, which is the outcome the
   gate needs; it is pre-registered anyway, unsoftened, because a prediction written to be safe is
   not a prediction.
4. **Induced term** — findings attributable to the iters 54–75 repair diff (seat G's scope) — is
   **below 6**.
5. **The two closed classes contribute ZERO blockers.** iter-74 and iter-75 both closed by
   adjudication with 0 corpus repairs; if a seat books a blocker inside either class, one of those
   adjudications was wrong and the close is retracted rather than defended.

## Expected lift

Clause 5 is graded **only** by a reading that returns **zero**. Anything else is a repair list plus
a re-read, and the honest expectation at open — with `N̂ ≈ 14` at the floor from iter-52 and 22
iterations of predicate fencing since — is that **this reading is more likely to produce a small
repair list than a zero**.

## Escalation conditions

- **Any blocker inside the iter-74 or iter-75 classes** → the corresponding close is **retracted**,
  not defended, and the iter re-scopes to that.
- **More than ~15 union blockers** → measure and route; do not repair inside this iter.
- **A seat that cannot state a `wc -l` for an assigned file** → that seat's reading is void and is
  re-run, because a partial pass is not a reading.

## Acceptable close-no-lift outcomes

- **A non-zero union, fully adjudicated and repaired with pre-commit double-reads.** The gate stays
  at 4 of 5 and the next reading decides it. That is the protocol working, not failing.

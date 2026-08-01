---
milestone: M257x
iter: 15
---

# iter-15 — decisions

## D1 — Run the suite UNSCOPED, even though a scoped run would have been faster

`run-playthroughs.sh:300-307` makes the ptreport gate **binding on a full run and advisory on a scoped one**,
and `:292` records the residual honestly: a pattern broad enough to select everything (`--grep '@pt'`) is
*still* graded advisory, so a deleted spec would be swallowed. Clause 2 asks for a verdict, and only the
unscoped invocation produces one. Cost: ~7 min per run instead of seconds. Taken deliberately.

## D2 — Fix the plan hazard with a BARRIER, not by rewriting the predicate

The obvious repair for `column "sequence_catalog" of relation "sequences" does not exist` is to make
`pg_get_serial_sequence` tolerant — wrap it, or filter attnames first. Both are the same defect in costume:
they leave a query whose correctness depends on the planner's choice, and the planner's choice changes with
the fifth execution of a prepared statement.

What landed instead pins the relation behind two `AS MATERIALIZED` CTEs (a hard optimization fence in PG12+)
and hands the function `tgt.reloid::regclass::text` — **the OID the column list was built from**. There is no
plan in which the two arguments can name different relations, because they are the same object.
`platform-alignment.md` §8 rule 4: *prefer a construct that cannot express the drift over a fence that
catches it.*

The fence is kept as well: `TestSeqDiscoverySQL_PinsTheRelationBehindAMaterializedFence` asserts both CTEs are
MATERIALIZED, that the OID form is used, that the parameter-respelled form is **gone**, and that the candidate
set is scoped to `tgt.reloid`. 3 mutants RED, control GREEN, every mutant built before it was run (§8 rule 5).

## D3 — Restore the sequence ownership on the APPLY side, not in the capture

The missing statement is `ALTER SEQUENCE … OWNED BY …`, which the capture *could* emit. It should not, for a
reason that is about reach rather than correctness: a capture-side fix only reaches a cache that has been
**re-captured from prod**, which this host cannot do (`HOST-M257x-toolchain`), and every existing cache on
every box would stay broken until someone did. The apply side derives the edges from the target's own catalog
and therefore self-heals for caches old and new — §2's *resolve it from the environment at the point of use*.

**Corollary, deliberately not taken:** do NOT heal inside `AdvanceIdentitySequences`'s refusal. That refusal
is the M256 guard; making it repair what it detects turns it into a probe that satisfies itself (§5 rule 7).
It stays the fence — if the reconciliation ever stops running, the replay fails loud again, naming the column.

A reconciliation failure is **fatal** to the provision rather than a warning: continuing would load rows the
sequence advance is about to refuse anyway, producing a stack that looks provisioned and replays red — the
exact state the whole auto-provision path exists to avoid.

## D4 — Do NOT read run 2's 17/31 as a regression against run 1's 20/31

Run 2 was executed **without `--reset`** after a mutating run had already completed onboarding for the
pt-world heroes. Its three net-new failures are all `onboarding.*`, whose negative controls assert
*"onboarding is INCOMPLETE"* — the stale-world state `run-playthroughs.sh:9-12` names as FORBIDDEN precisely
because it produces exactly this. The two numbers describe two different worlds and are not comparable.

The measurement that IS comparable was taken at the layer that changed: replay **rc=1 → rc=0**, **0 → 11 986
rows**, `403 → 200` on two collections. This is the same discipline as §5 rule 12 (*say which invocation
produced the number*), one level up: **say which world.**

Consequence, stated rather than hidden: **the fix did not move the clause-2 metric.** It removed one of at
least four causes.

## D5 — Land the replay fix; ROUTE the readiness-probe defect, do not fix it here

`run-playthroughs.sh:161`'s readiness gate announces `✓ fake-FAPI ready (HTTP 000000)` on a connection that
never happened, because `curl -w '%{http_code}'` writes `000` on a failed transfer **and** exits non-zero, so
`|| echo 000` appends a second one and `!= "000"` is true. Reproduced against a dead port. It is a four-line
fix and it is correct and ready.

It is routed anyway. The iter had already opened its third line of investigation (suite → readiness probe →
replay), and the replay chain turned out to be the planned target's actual root cause. Landing an unrelated
correct fix at that point is what the scope-creep tripwire exists to stop. The sweep's finding is recorded in
full so the next iter starts from evidence: **one verdict-flipping site** (playthroughs) and **three
message-degrading siblings** (`stack-verify/lib/services.sh:132`, `readiness.sh:151`, `:185`) where the
concatenation only mangles the diagnostic and the `down` verdict still comes out right. That distinction is
the expensive part and it is already paid for.

**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*). The class iter-138
reached for and could not have: **a published receipt names its own command, pathspec and ref, so unlike a
bare `:NN` pin it is self-contained and genuinely re-runnable.**

# iter-140 — the receipts, and whether they reproduce

## Population, enumerated before any verdict

**22 published receipts carrying a claimed count, across 15 files** — a backticked
`git grep` / `grep` / `git log -S` / `git ls-tree` / `git show` followed within 160 characters by a result
verb *and* a number. Concentration: `sentinel.md` 4 · `external_services.md` 3 ·
`security_compliance.md` 2 · `studio-desk.md` 2.

**Check rule sealed in `overview.md` before any receipt was run:** run **verbatim as published**, in
listed order; **no re-wording to make it pass** — a receipt that needs re-wording is *not-reproduced*; and
*not-checkable* (repo or ref absent) is published as **its own count**, never folded into either other
bucket (iter-139's lesson).

## Result — 9 checkable, 7 reproduce, 2 do not

| # | receipt | claimed | measured | verdict |
|---|---|---|---|---|
| R1 | `security_compliance.md:95` — `grep -c 'OrganizationMixin{}' schema/*.go` | 30 | **30** | ✅ |
| R2 | `security_compliance.md:235` — `git grep -n ValidateToken ad9f3c498 -- '*.go'` | 8 hits, `token.go` + `token_test.go` | **8**, those two files | ✅ |
| R3 | `dependency_map.md:59` — `git grep -n SKILLER_STREAM ad9f3c49 -- '*.go'` | 6 lines, 3 named files | **6**, `main.go` · `subscriber_merge_test.go` · `subscriber_wiring.go` | ✅ |
| R4 | `studio-room.md:367` — `git grep -i mistral aeec036a` | 22 hits / 3 files | **22 / 3** | ✅ |
| R5 | `safety.md:1082` — `grep -rn "BUNNY_RECORDING" .agentspace/rosetta-extensions/` | 0 | **0** | ✅ |
| R7 | `build-budget.md:87` — `grep -rn "BRINGUP_ANCHORS" .` | exactly 2 | **2** | ✅ |
| R8 | `ai_architecture.md:54` — `grep -rin 'bedrock\|boto3' app/studio/` | 0 | **0** | ✅ |
| **R9** | **`sentinel.md:5`** — `git grep "authorization\|AUTHORIZATION_ADDRESS" fa47850d -- '*.go' go.mod` | *"one unrelated hit"* | **0 lines, exit 1** | ❌ |
| **R6** | **`latency-budget.md:365`** — `grep -n "redirect_url" clerkenstein/clerk-frontend/*.go` | *"one non-test occurrence"* | **22 lines / 4 files; 3 non-test; 1 CODE** | ❌ |

**Denominator stated: 9 of 22 were checkable on this box** at the ref each names. The other 13 are
**not-checkable here** and are published as that, not as passes.

## The two failures share one shape, and it is not carelessness about the conclusion

**Both conclusions survive.** R9's is *strengthened* — messenger's Go source imports no authorization
client, and the true count is **zero**, not one. R6's is intact — there is exactly **one code**
occurrence of `redirect_url` (`server.go:414`); the other two non-test lines are **comments** at `:150`
and `:155` describing the same handshake bounce.

> **`D-M257x-140-1`. The number in each receipt was written from the CONCLUSION, not from the command's
> output.** That is a specific and repeatable authoring failure: you know the answer, you write the
> command that demonstrates it, and you fill in the count from what you know rather than from what it
> printed. **The defect is not the number — it is that the receipt no longer demonstrates anything.** A
> reader who runs R6 sees **22** where the page says **one**, and the rational response is to distrust
> the paragraph, including the parts that are exactly right.

**R9 was found by an adjudicator; R6 was found by this census** — nobody had looked at it. That is the
argument for censusing a class rather than sampling it, made on a class where the census actually works.

## Why this class is censusable and iter-138's was not

iter-139 retracted the bare-`:NN` census because those pins have **no resolvable head** — the machine
cannot tell which file a continuation pin continues. **A receipt carries its own head**: the command names
the pathspec and the ref. The same strategy (`TOK-08`) applied to a class with a resolvable subject
returns a usable answer; applied to one without, it returned 0-for-12. **The strategy was never the
variable — the subject's decidability was.**

## The gate caught MY OWN edit, in the same iter — which is iter-138's lesson paying out

The first gate run went **RED**: `anchor_construct_guard` (2 findings) + `repair_postcondition`, plus 2
scoped-suite failures — all on **three bare cross-repo pins this iter had just written**
(`cmd/root.go:14`, `cmd/send.go:1`, `cmd/trigger.go:2`, published as the positive-control receipt). Bare
heads, so the resolver bound them to the **wrong clone**, where one lands on a blank line.

**That is `§5` rule 63(c) biting its own author inside the iter that cites it.** Fixed by **removing the
pins and naming the package** — not by re-pinning them at the right repo.

**The contrast with iter-137 is the point.** iter-137 shipped an equivalent defect and it survived a
whole iter, because that iter picked its suites **by topic**. iter-138's `D-M257x-138-5` said *choose the
suites by what you CHANGED*; iters 139 and 140 have run the anchor set every time since, and here it
**caught the defect before the commit**. One iter of latency → zero.

## Test gates

| gate | result |
|---|---|
| **Guard family** (`--repo-root` + `--platform stack-demo/platform @ 0c91421`) | **18 GREEN · 0 RED · 4 not-run** — *after* a genuine RED on this iter's own edit (above) |
| **Scoped fence suites** — chosen by what this iter CHANGED | **102 passed / 0 failed** — *after* 2 failures on the same self-inflicted defect |
| **Whole suite** | **NOT re-run — §5 rule 60 requires saying so.** Zero `rosetta-extensions` files changed; iter-132's clean run stands on the same rext tree (`223e4a6`) |
| **Suite wall-time** | not quoted — `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` stands |

## Close — 2026-08-08

**Outcome:** the published-receipt class enumerated (**22 across 15 files**) and the **9 checkable ones
run verbatim: 7 reproduce, 2 do not** — `sentinel.md:5` claims *"one unrelated hit"* and returns **zero**;
`latency-budget.md:365` claims *"one non-test occurrence"* and returns **22 lines / 3 non-test / 1 code**.
Both repaired, both conclusions intact, and **one of the two was found by this census and by nothing
else**. The class is censusable **because a receipt carries its own head** — which is exactly what the
retracted iter-138 subject lacked.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged; no reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 134–140 took no reading, so the metric is UNMEASURED not unmoved — §9's iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**4 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-140-1` (a receipt's number written from the conclusion instead of the output stops
the receipt demonstrating anything) · `D-M257x-140-2` (a class is censusable iff its subject carries its
own head — the variable was never the strategy).
**Side-deliverables:** none.
**Routes carried forward:**
- **`FIX-M257x-iter140-receipts-not-checkable-here` (NEW)** — **13 of the 22 receipts could not be run on
  this box** (repo or ref absent). They are neither passes nor failures. Re-run them where those clones
  exist, or record each as unverifiable at its site.
- **`FIX-M257x-iter140-receipt-fence` (NEW)** — this class is fenceable and its population is small (22).
  A fence that re-runs each published receipt and compares the printed count is buildable **once head
  resolution exists for the pathspec** — the same prerequisite iter-139 named, arriving from the other
  side.
- `FIX-M257x-iter138-anchor-rot-fence` (re-specified: head resolution first) ·
  `FIX-M257x-iter135-adjudicated-live-defects` (remainder: `clerk-integration.md:126` ·
  `backend.md:13`'s *UNEVEN* cross-ref · `ai-readiness.md:18-20` · `org-repos.md:227`,`:370`,`:43` ·
  `ai_architecture.md:111`,`:224` · `next-web-app.md:17`,`:186` · `external_services.md:368`) ·
  `FIX-M257x-iter134-fence-family-has-no-shared-predicate-layer` ·
  `FIX-M257x-iter133-two-fives-need-a-fence` · `FIX-M257x-iter132-suite-walltime-is-not-a-measurement` ·
  `FIX-M257x-iter131-predicate-sets-not-enumerated`.
- **CLOSED this iter:** `adj-B`'s **P-1** (`sentinel.md:5`).
**Lessons:**
1. **Write the number from the output, never from the conclusion.** Both failures came from an author who
   already knew the answer. The receipt then stops demonstrating anything, and a reader who runs it
   learns to distrust the surrounding prose — including the parts that are right.
2. **A class is censusable iff its subject carries its own head.** Receipts do; bare `:NN` pins do not.
   `TOK-08` was not wrong at iter-138 and right at iter-140 — the subjects differed.
3. **Publish the not-checkable count.** 9 of 22 were runnable here; saying "7 of 9 reproduce" without the
   22 would be the same over-claim iter-139 retracted.

---
iter: 124
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-07
---

# iter-124 — tier 2, triaged by consequence

**Run 80, tik 1.** The user's directed scope, priority 1: the census's tier-2 population — the factual
assertions that carry **no citation and no hedge** — has never been touched by anything in this milestone.
`D-M257x-122-6` built a **ratchet** over it (direction, per file) precisely because 1,151 could not be
repaired in one iter. This iter does the thing the ratchet cannot: it **decides what each one actually is**.

## Active strategy reference

**There is no `TOK-09` and there will not be one.** `TOK-08`'s sealed rule bars a successor strategy, it
was REFUTED at iter-119 by its own arithmetic, and the user has since supplied the scope directly. This tik
runs under **the user's directed scope of run 80**, recorded in the milestone-root `decisions.md` at
iter-122's framing: *the census is an INSTRUMENT, not a strategy.* `F4` (`D-M257x-122-2`) binds every
sentence below: **nothing here is `P`, is `N`, or is a clause-5 verdict.** The gate stays **4 of 5**.

## Cluster / target identified — and the re-survey that moved it

Phase 1 Step 0 re-survey, run before targeting:

```
cd .agentspace/rosetta-extensions/stack-core
/usr/bin/python3 claim_census_guard.py --repo-root <rosetta> --census \
    --tier2-export /tmp/m257x-iter124-tier2.json
```

| | at iter-122 (the directive's figures) | at iter-124 open, corpus `1f757cd` |
|---|---|---|
| assertion candidates | 3,292 | **3,479** |
| tier-2 unevidenced | 1,151 | **1,164** |
| hedged blocks | 24 | **116** |
| files in scope | 40 | **41** |

**The target grew, and the growth is iter-123's own.** `org-repos.md` is net-new (+27) and 7 existing files
went down by 13. The ratchet holds at 1,164. **`hedged blocks` 24 → 116 is NOT a 92-hedge improvement** —
it is a different field: the directive's "24" is *uncited-but-hedged assertion sentences*, this is *blocks
carrying a hedge token*. Two denominators, and they are not comparable. Stated here so the run's own report
cannot conflate them.

## Hypothesis

Of the 1,164, the great majority are **citable** — the evidence exists and was simply never given — and a
small minority are **wrong**. If that holds, the corpus is **under-cited, not unfounded**, and that is a
materially different disease with a materially different treatment. Nobody has measured which.

## The pre-registered triage protocol (sealed in this iter's FIRST commit, before any verdict)

### The four fates, and they are exhaustive

Every tier-2 sentence resolves to exactly one:

| fate | test |
|---|---|
| **cite** | the evidence exists and is reachable from here; the sentence simply never gave it |
| **hedge** | it cannot be checked from here, and the sentence must say so (iter-093's principle) |
| **fix** | it is wrong |
| **drop** | it asserts nothing checkable and earns nothing (placeholders, template scaffolding, pure restatement) |

### Ordering is by CONSEQUENCE, not by class

The user's rule, and this milestone has earned it: consequence-ordering surfaced three security-surface
understatements that class-ordering never would have. The consequence class is defined by a **token
predicate over the sentence**, not by filename, so it cuts across files:

```
C1 = auth* | token | secret | password | credential* | encrypt* | TLS | HTTPS | SSL | PII | GDPR
   | residen* | tenan* | isolat* | backup* | retention | RBAC | ABAC | Casbin | JWT | session
   | permission* | unauthenticated | public | firewall | VPN | CORS | Clerk | privacy | compliance
   | SOC 2 | ISO 27001 | DPA | sub-processor | delete* | erasure | audit
```

**Measured at open: `|C1| = 344` of 1,164 = 29.6 %**, spread over **34 of the 40 files** that carry tier-2
debt. That is this iter's denominator and it is named before the first verdict.

### What lands, and what does not — stated in advance so the close cannot be flattered

- **Every `fix` verdict in C1 lands in this iter.** A wrong sentence about a security or data-handling
  surface is the highest-consequence thing in the population.
- **`drop` and `hedge` land where they fall out.**
- **`cite` is REPORTED, not bulk-landed.** Citing ~n sentences is a volume of evidence-gathering, not a
  triage; bulk-landing it here would be the repair-volume mistake `TOK-07` was refuted for. The count is
  the deliverable; a **proof-of-citability sample** is taken so "citable" is a measurement and not a hope.
- **The residual (1,164 − 344) is named as untriaged**, never implied to be clean.

### Pre-registered falsification

> **If `fix` ≥ 15 % of C1**, the hypothesis is wrong: the corpus is not under-cited, it is **unfounded** in
> its highest-consequence class, and that is a finding that outranks everything else in this run — report
> it as the headline and stop adding citations until it is understood.
>
> **If `cite` < 50 % of C1**, "the evidence exists and was simply never given" is not what tier 2 is, and
> the triage's own framing needs replacing rather than extending.

Both branches are checked at close and reported with the number either way.

## Expected lift

**No `N` movement is claimed, and no reading is taken** (§9's guard-rail 1, in its own words). Clause 5 is
met only by a graded read that returns zero; a triage is not a read. The lift claimed is the **split
itself** — the first measurement of what the largest untouched class in the corpus actually consists of.

## Phase plan

1. Seal this pre-registration (commit 1).
2. Triage C1 exhaustively, in consequence order, file by file, reading each sentence in its block context.
3. Land every `fix`; land cheap `drop`/`hedge`.
4. Take the proof-of-citability sample.
5. Re-run the census + the affected guards; check both falsification branches; close.

## Escalation conditions

- A `fix` that inverts a shipped security property → land it, and **file it to the platform-defect
  register** if the defect is the platform's rather than the corpus's.
- No legal/compliance/policy escalation under any verdict (§5 rule 48). Filing is not escalation.

## Acceptable close-no-lift outcomes

If the triage shows the split but no `fix` verdicts land, the iter still closed its planned scope — the
split was the deliverable. A falsification that fires is a first-class outcome, not a failure.

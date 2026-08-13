# iter-124 — tier 2, triaged by consequence

**Type:** tik · **Run 80, tik 1.** Priority 1 of the user's directed scope: decide what the census's
1,164 unevidenced assertions actually ARE. The ratchet (`D-M257x-122-6`) could only hold their count.

## 0. The answer, first — because it is the thing nobody had measured

> ### The corpus is **UNDER-CITED**, not unfounded.

| fate | printed | corrected by the audit | what it means |
|---|---|---|---|
| **cite** | **331 = 96.2 %** | **≈ 298 = 86.6 %** | the evidence exists and was simply never given |
| **hedge** | 7 = 2.0 % | ≈ 42 = 12.2 % combined with `drop` | genuinely uncheckable from here |
| **drop** | 2 = 0.6 % | *(see above)* | asserts nothing checkable |
| **fix** | **4 = 1.2 %** | **a FLOOR; ≈ 11 ≈ 3.2 % recall-corrected** | measured false |

**Denominator: `|C1| = 344`**, the sealed consequence class — the 29.6 % of the tier-2 population whose
sentences touch security, data-handling, identity, tenancy, residency, backup or access. The predicate
is committed (`triage-predicate.py`) so the denominator is reproducible, not a filename list.

Over the **whole** tier-2 population (denominator **1,164**) the printed split is `cite` 1,146 = 98.5 %,
`drop` 7, `hedge` 7, `fix` 4. The C1 numbers are the ones to read: the residual 820 were **not
consequence-ordered and not hand-audited**, and are named untriaged below rather than implied clean.

**Both pre-registered branches were checked and neither fired** — `fix ≥ 15 %` (measured 1.2 %) and
`cite < 50 %` (measured 96.2 %). The pre-registration was sealed in this iter's first commit `8a7eb40`,
before a single verdict.

## 1. What the answer changes

*"1,151 uncited assertions"* has been quotable for two iters and has read, to anyone who met it, as
**1,151 things that might be wrong**. It is not that. It is **~298 sentences whose evidence is sitting in
a clone nobody linked**, ~42 that genuinely cannot be checked from here, and a **floor of 4 that are
false**. Those are three different diseases with three different treatments, and only the third is
urgent.

**The instrument that could not tell them apart is the same one that counted them.** The census measures
*unevidenced*, and says so in its own `KNOWN_WEAKNESS` — but a number with no fate attached defaults, in
a reader's head, to the worst fate available.

## 2. The `fix` class — and it is the same predicate iter-123 measured, unrepaired at 24 twins

Consequence-ordering did what the user said it would. The largest live falsehood in the clause-5 surface
was not a subtle one:

> **iter-123 measured that the Cosmo Router's production module is DESTROYED.** That correction reached
> **`graphql-wundergraph.md`'s fold cell and `org-repos.md` § 3, and nothing else.** **24 sites across 13
> files** still said *"prod only"*, *"still declared in production"*, *"survives in production only"*, or
> drew production traffic through it — **including the fenced `platform-migration-status.md` row whose
> three siblings (`cms`, `roadrunner`, `messenger`) iter-123 DID repair in the same pass.**

The evidence, restated where every repaired site now carries it: `module.wundergraph_euwest1` is deleted
from `infrastructure/terraform/production/services.tf` @ `13c248e6`; `:509-517` records that the apply
destroyed *"its ECS service, task definition, target group, ALB rule (priority 810), Cloud Map entry, log
group, ACM cert and the `wundergraph.anthropos.work` alias"*, leaving only a `removed{}` for the ECR
(`:521`), hand-deleted 2026-08-05 — *"so production-wundergraph is gone and this block is now inert."*

**And the half that is NOT measurable is now stated as such rather than guessed:** where production's
frontends send GraphQL today is Vercel runtime configuration, in no clone set. Three repaired sites say
so in those words. That is the `hedge` fate used correctly — not manufactured for a fact somebody can go
and measure, but supplied for one nobody here can.

**Second `fix` class, same shape:** the `db-backup` claim. iter-123 established **Bash not Go, two
targets not three, never Azure, no trigger since `7dd1b80` (2025-05-29)** — and that *"every 6 hours"*
**never had a source** (the disabled value was `rate(12 hours)`). **Three sites outside `db-backup.md`
still carried the refuted version, two of them in the security-and-compliance posture.** Durability is
now stated correctly there: RDS Multi-AZ plus an hourly AWS Backup plan with PITR carries it; what the
stalled job costs is the **offsite, non-AWS leg**, not recoverability.

**Why inspection could not have found either.** Every one of the 27 sites is *locally plausible*. Nothing
about *"the Cosmo Router — prod only"* looks wrong on the page; it looks like a careful hedge. The class
is visible only when the predicate is enumerated corpus-wide — `TOK-02`'s method and `TOK-07`'s unit of
repair, both vindicated on a class neither was aimed at. → `D-M257x-124-3`.

## 3. Three fences fired on this iter's own edits, and one refused the commit

The milestone's standing question is whether the fences catch **the author**:

| fence | caught | how it presented |
|---|---|---|
| `claim_census_guard` ratchet | `service_taxonomy.md` **68 → 69** — a repair split one uncited sentence in two | `RATCHET BROKEN`, named the file |
| `unreadable_repo_claim_guard` | a new `module.*_euwest1` mention with **no ref pin** | `RED`, named the line |
| **`repair_postcondition`** | **4 citations to `architecture_overview.md:335` became blank-line anchors** — this iter's own edits moved the target 8 lines down | **rejected the commit at the pre-commit hook** |

All three were repaired **by adding the evidence**. None was silenced, waived or baselined away. The
third is the one worth keeping: a repair that corrects 24 false sentences while silently breaking 4 true
citations is a net loss, and it was caught only because the fence runs at commit time.

## 4. Reach, with its denominator (§5 rule, iter-114)

| statement | number | denominator |
|---|---|---|
| C1 triaged | **344 / 344 = 100 %** | the sealed consequence class |
| C1 as a share of tier 2 | 344 / 1,164 = **29.6 %** | the tier-2 population at open |
| tier-2 population **left untriaged** | **820** | *named, not implied clean* |
| triage rules hand-audited | **30 / 344 = 8.7 %** | seeded, committed, published |
| rule accuracy | **27 / 30 = 90.0 %** — `R3` 21/21, `R4` 6/9 | the audit sample |
| `fix` sites repaired | **27** across 14 files | 4 were tier-2-flagged; **23 were not** |

**That last row is the honest limit of the census as a defect-finder.** Of the 27 false sentences
repaired, the tier-2 enumeration flagged **4**. The other 23 sat in blocks that carry *a* citation
somewhere, which exonerates every sentence in them — the census's own declared blind side (`KNOWN
WEAKNESS` clause 1), measured here rather than assumed. **The census pointed at the neighbourhood; the
consequence-ordered reading found the class.**

## 5. Guards

**22 members · 17 GREEN · 0 RED · 1 could-not-check · 4 not-run.** **Not a whole-family green, and the
runner's own summary line is what this sentence quotes** (`guard-family: 17 GREEN · 0 RED · 1
could-not-check · 4 not-run`; the family exits **2**). Invocation:

```
cd .agentspace/rosetta-extensions/stack-core
/usr/bin/python3 guard_family.py --repo-root <rosetta> --platform <rosetta>/stack-demo/platform
```

- `platform_alignment_guard` remains the **could-not-check** — the corpus cites repos no stack clones.
  **Unchanged by this iter and deliberately not touched**: resolving it is a named item of run 80 and is
  a separate line of investigation; opening it here would have fired the scope-creep tripwire. Routed.
- `claim_census_guard`: tier 2 **1,164 → 1,160**; ratchet **holds**, no file rose at close.
- `unreadable_repo_claim_guard`: **22 mentions — 9 by an unmeasurable marker, 12 by a ref-pinned
  reading.** Its own closing note flags the 9 as a reconciliation debt → `D-M257x-124-6`.

## 6. Routes carried forward

- `FIX-M257x-iter124-desired-count-predicate-reach` — enumerate the *"a service repo's own
  `service_desired_count` is evidence of production state"* predicate across `corpus/ops/**` too; this
  iter covered the clause-5 surface plus `CLAUDE.md`.
- `FIX-M257x-iter124-stale-hedges-on-infrastructure` — **9** sites hedge about a repo that iter-123 made
  readable. Denominator is the guard's own. Needs an `infrastructure` clone at a named sha.
- `FIX-M257x-iter124-tier2-residual` — **820** tier-2 members outside C1, untriaged.
- `corpus/ops/demo/coverage-protocol.md:349` instructs *"restart cms+router+directus"* — services that no
  longer exist. Stale ops instruction, outside the clause-5 surface.
- **The `platform_alignment_guard` resolution** — carried to the next iter of this run.

## Close — 2026-08-07

**Outcome:** the largest untouched class in the corpus is triaged over its consequence subset — **`cite`
96.2 % printed / ≈ 86.6 % audit-corrected, `hedge`+`drop` ≈ 12.2 %, `fix` 4 as a floor**, denominator
**344**; both pre-registered branches checked and neither fired; **27 false sentences repaired across 14
files**, all of them one predicate iter-123 had measured and whose correction had reached 2 sites.
**No `N` movement is claimed and no reading was taken.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — 4 of 5, unchanged; `P` unmeasured — (2) triggered-tok: n (**a
successor strategy remains FORBIDDEN by `TOK-08`'s sealed rule; this iter runs under the user's directed
scope, and `F4` books any sentence treating the census as the grader as a defect**) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-124-1` … `D-M257x-124-6` (see [`decisions.md`](decisions.md))
**Side-deliverables:** the `CLAUDE.md` M810/`cms` correction (iter-123 refuted *"has not moved for cms"*
and *"NOT MEASURABLE from our clone set"*; the correction had not reached the file every agent loads).
**Co-committed rather than committed separately — a process miss, disclosed in `D-M257x-124-5`**, and it
does not change the close status, which grades planned scope only.
**Lessons:** **A correction that reaches one cell is not a correction.** iter-123 repaired three of four
rows sharing one predicate; the fourth, and 24 downstream twins, survived a week of readings because each
is locally plausible. When a measurement retracts a claim, the unit of repair is the **predicate,
corpus-wide, in the same iter** — and the enumeration must be recorded so the reach can be checked. →
`platform-alignment.md` §5 **rule 53**.

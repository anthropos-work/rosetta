# iter-127 — the hedges that expired

**Type:** tik · **Run 80, tik 4.** `FIX-M257x-iter124-stale-hedges-on-infrastructure`.

## 1. The finding: the stale thing was not the hedge token, it was the VERDICT built on it

`unreadable_repo_claim_guard` counts *marker-carrying paragraphs*. Chasing that counter would have been
chasing the wrong noun. **What actually expired is the CONCLUSION those markers licensed** — the
*"report both, assert neither"* verdict on `cms`'s production state:

> `cms`'s prod ECS module is **not** a settled rollback path … `cms/terraform/main.tf:39` still declares
> it at `service_desired_count = 0` … the destruction happens in infrastructure's `services.tf`, **which
> we cannot read**.

**iter-123 read it.** `infrastructure` @ `13c248e6` declares **no `module "cms"` at all**, and
`infrastructure/terraform/production/services.tf:64-70` records the apply destroying the ECS service, task
definition, ECR repository, IAM roles, security group, Cloud Map entry, log group, alarms and the ten
`/production/cms/*` SSM parameters. **`cms/terraform/main.tf:39` is orphaned dead code.**

**iter-123 propagated that to `platform-migration-status.md`'s `cms` row, and iter-124 to `CLAUDE.md`.
Five more sites kept the retracted verdict**, two of them in fenced summary tables:

| site | what it still said |
|---|---|
| `cms.md` banner | *"scaled to zero, not deleted … the one M810 row whose terraform module block has not moved"* |
| `cms.md` Infrastructure bullet | *"this repo's own module block has not moved, and that is all that can be said from here"* |
| `backend.md` | *"The M810 prod teardown is UNEVEN — do not state the two together"* |
| `architecture_overview.md` | *"the prod ECS module is **not** a settled rollback path — report both, assert neither"* |
| `service_taxonomy.md` | same, in a second fenced table |

All five repaired, each stating what it used to say. **This is the third instance of one class in this
run** — after the router (24 sites) and `db-backup` (3) — and the class now has a rule (§5 **rule 54**).

## 2. What is NOT claimed

**The guard's reconciliation NOTE is unchanged: `9` sites hedge about `infrastructure` while `13` report
having read it.** This iter repaired the *verdicts*, not the *marker tokens*, and the counter did not move.
Saying so is the point — a repair that does not move the number it was routed against must **report the
number**, not imply it did.

Whether those 9 markers are individually stale needs a paragraph-by-paragraph read against the
`infrastructure` clone. Routed forward, unchanged, as
`FIX-M257x-iter124-stale-hedges-on-infrastructure` — **not closed by this iter.**

## 3. Reach

| statement | number | denominator |
|---|---|---|
| sites asserting `cms`'s prod state is unsettled | **5 → 0** | `git grep` for `cms/terraform/main.tf:39` + *"assert neither"* over `corpus/` + `CLAUDE.md` |
| guard reconciliation note | **9 hedged / 13 measured** | unchanged — stated, not implied closed |

## 4. Guards

`unreadable_repo_claim_guard` **OK** (23 mentions: 9 by marker, 13 ref-pinned) · `corpus_citation_guard`
**OK** · `anchor_construct_guard` **OK** · `claim_census_guard` **OK, ratchet holds (1,160 / baseline
1,164)** · `markdown_structure_guard` **OK**. Whole-family reading taken at run close.

## Close — 2026-08-07

**Outcome:** the stale thing was the **verdict**, not the hedge token — five sites (two of them fenced
summary tables) still called `cms`'s production state unsettled four days after it was measured;
all five repaired. **The guard's 9/13 reconciliation note did not move and is reported as unchanged
rather than implied closed.** **No `N` movement is claimed and no reading was taken.**
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — 4 of 5, unchanged — (2) triggered-tok: n (**successor strategy
FORBIDDEN by `TOK-08`'s sealed rule**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n
(4 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-127-1` (below)
**Side-deliverables:** none.
**Routes carried forward:** `FIX-M257x-iter124-stale-hedges-on-infrastructure` stays **open** — the 9
marker-carrying paragraphs are untouched and need a per-paragraph read against the `infrastructure` clone.
**Lessons:** **A hedge has an expiry, and nothing in this corpus watches it.** Three fences check that a
claim about an unreadable boundary *carries* a marker; none checks whether the boundary is **still**
unreadable. The moment a repo enters the clone set, every hedge about it becomes a candidate manufactured
hedge — and the guard that counts them cannot tell, because *"is this still true"* is not a property of
the sentence. **Status is closed-fixed-partial precisely because the routed item is not closed.**

---

## `D-M257x-127-1` — repair the VERDICT, and report the counter you did not move

The item was routed against a guard counter (9 hedged / 13 measured). The substantive defect was
somewhere else: five sites publishing a **conclusion** the corpus had already retracted. Fixing those was
worth more than moving the counter, **and it did not move the counter.**

**Decision: do both halves out loud.** Repair the verdicts; then state the counter's unchanged value and
leave the routed item open. The tempting close — *"stale hedges reconciled"* — would be true of the
sentences and false of the number, and this milestone has spent iterations on exactly that gap between a
conclusion and the quantifier under it (`D-M257x-121-2`).

**And the general form is worth keeping:** a fence can check that a claim **carries** its qualifier. No
fence in this family checks whether the qualifier is **still warranted** — that requires re-measuring the
boundary, which is the work the qualifier exists to say nobody did. **Qualifiers rot in the one direction
fences cannot see.**

# iter-92 — decisions

## `D-M257x-92-1` — the brief's premise did not hold at HEAD, and re-doing the sweep would have been the error

The run brief carries *"~14 passages across 11 files treat it as one future event. Sweep them."* Re-measured
before acting: **15 files, 40 occurrences**, and the great majority **already state the split correctly** —
`corpus/README.md:16`, `CLAUDE.md:189`, `services/backend.md:36`, `services/cms.md:9`,
`services/README.md:17`, `services/jobsimulation.md` (5 occurrences, all correct), and both fenced map rows.

So the sweep landed in an earlier iter and the brief's count is stale. **Re-running it would have re-landed
landed work** — precisely what Phase 1 Step 0 exists to prevent, and the failure mode `D-M257x-59-1`
described as repairing by claim when the claim is already repaired.

Recorded because the instruction came from the user and was followed by measuring it rather than by doing
it: **a task description is a claim too.** The right response to *"sweep them"* is to look first, and to
report back that they are swept — with the count — rather than to produce motion.

## `D-M257x-92-2` — the real defect: a FENCED claim restated UNFENCED, and stronger

What the re-survey did find is more interesting than the sweep would have been.

The fenced map's `cms` row (`platform-migration-status.md:88`) is careful and explicit:

> **Whether that rollback declaration still stands is not something this map can see** — it never could,
> since infrastructure has never been in the clone set.

But `services/backend.md:36-37` said, flatly:

> `cms` has not moved: `module.cms_euwest1` **is still declared** as the rollback path and takes no traffic

and `CLAUDE.md:241` said the same. **Both assert as fact exactly what the fenced document says it cannot
see.** The map is machine-fenced; its restatements are not, and the unfenced copies drifted *upward* in
confidence — the fence held the line in the one file a guard reads and nowhere else.

**This is the generalisable finding, and it is a limit on the whole TOK-02/TOK-05 fencing method:**
fencing a document does not fence its paraphrases. A claim that is hedged where it is fenced and flat where
it is quoted is *worse* than an unfenced claim, because the hedge exists and creates the impression the
system checked it.

Corrected in `backend.md`, `CLAUDE.md`, `cms.md` and the map itself; each now reports both measured facts
and asserts neither.

## `D-M257x-92-3` — cms HAS moved, in the direction opposite to what the corpus recorded

Measured at `origin/main` in the `cms` repo (fetched this session):

| commit | date | what |
|---|---|---|
| `8f4840b` | — | `service_desired_count` → 0 (already in the corpus) |
| `6efa1d5` → merged `f38c0c4` | 2026-08-04 | **deletes** `.github/workflows/build-production.yml` — *"the cms ECR repository is decommissioned (M810)"* |

`6efa1d5`'s body states that M810 *"deletes `module "cms_euwest1"` from the platform's `services.tf`, which
destroys the ECS service and the production-cms ECR repository"*, and that the workflow was dropped because
it *"would try to push an image into a registry that no longer exists."*

So the `cms` repo now contains **two measured facts pointing in opposite directions**: a terraform block
that still declares the module at `cms/terraform/main.tf:39`, and a CI commit asserting the registry is
already gone. The corpus was quoting only the first, and calling it *"the one M810 row that has not
moved"* — true of the module block, and misleading about the repo.

**The destruction itself lands in `infrastructure/services.tf`, and `infrastructure` has never been in any
clone set we have.** So this is UNMEASURABLE — and it is now unmeasurable *with contrary evidence on both
sides*, which is a different and more honest epistemic position than the corpus had. Both are reported;
neither is asserted. This is the deferred-by-rule boundary the brief names, respected rather than argued
past.

## `D-M257x-92-4` — iter-91's new fence caught its own author within the hour

The map edit above introduced a citation reading `` `main.tf:39` `` instead of `` `cms/terraform/main.tf:39` ``.
The family run immediately went:

```
platform_alignment_guard  CANNOT-CHECK  rc=2
  1 citation(s) could not be resolved at all (main.tf) — each is a claim this run did not check.
```

**Before iter-91 that citation would have been counted as `unresolvable` and printed, and the guard would
have exited 0 GREEN.** The map would have shipped with a dead citation and a green fence over it.

Two things worth recording. First, it is the **first live catch** of the third verdict, and it caught the
person who built it — the eighth-consecutive-iteration pattern §8 already notes, now with the fence
actually stopping it instead of the pattern merely being observed. Second, it is a **positive control that
arrived for free**: the fence was proven discriminating on real prose within minutes of landing, which is
worth more than the synthetic mutant, because nobody constructed it.

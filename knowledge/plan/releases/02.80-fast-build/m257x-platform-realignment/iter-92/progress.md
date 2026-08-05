**Type:** tik, under `TOK-05`.

# iter-92 — the M810 sweep: already done, and the one real defect was a fenced claim restated unfenced

## The re-survey changed the target, which is the protocol working

The brief carries *"~14 passages across 11 files treat M810 as one future event. Sweep them."* Measured at
HEAD first: **15 files, 40 occurrences**, and the great majority **already correct** — `corpus/README.md:16`
(*"is **uneven**: **landed for jobsimulation**"*), `CLAUDE.md:189`, `backend.md:36` (*"UNEVEN — do not state
the two together"*), `cms.md:9`, `services/README.md:17`, `jobsimulation.md` (5 of 5), both fenced map rows.

**The sweep had already landed.** Re-running it would have re-landed landed work. A task description is a
claim too, and the right answer to *"sweep them"* was to look first and report the count.

## Ground truth, measured from the platform (not from either document)

| service | measured state | evidence |
|---|---|---|
| `jobsimulation` | **module block DELETED** — ECS service, task definition, ECR repo, IAM roles, SG, Cloud Map, log group, alarms all destroyed. The module *file* survives to own the LiveKit/Chime buckets `backend` reads by literal name | `6092c6d2` (merged `caf36c96`); `jobsimulation/terraform/main.tf:15-22` |
| `cms` | **module block still declared** at `service_desired_count = 0` — **and** the build-production workflow **deleted** on 2026-08-04 saying *"the cms ECR repository is decommissioned (M810)"* | `cms/terraform/main.tf:39`; `6efa1d5` merged `f38c0c4` |

## The two things the original framing would have missed

### 1. A FENCED claim was being restated UNFENCED, and stronger

The fenced map says of `cms`: *"**Whether that rollback declaration still stands is not something this map
can see** — it never could, since infrastructure has never been in the clone set."*

`backend.md:36-37` and `CLAUDE.md:241` both asserted, flatly, that `module.cms_euwest1` **is still declared
as the rollback path**. That is precisely what the fenced document says it cannot see.

**Fencing a document does not fence its paraphrases.** The claim was hedged where a guard reads it and flat
everywhere else — which is *worse* than an unfenced claim, because the hedge exists and implies the system
checked it. This is a real limit on the TOK-02/TOK-05 method and it is recorded as one.

### 2. `cms` HAS moved, opposite to what the corpus recorded

`6efa1d5` (2026-08-04) deleted `.github/workflows/build-production.yml` because it *"would try to push an
image into a registry that no longer exists"*, its body stating that M810 *"deletes `module "cms_euwest1"`
… which destroys the ECS service and the production-cms ECR repository."*

So the repo holds **two measured facts pointing opposite ways**. The corpus was quoting only the older one
and calling cms *"the one M810 row that has not moved"* — true of the module block, misleading about the
repo. The destruction lands in `infrastructure/services.tf`, **which has never been in any clone set we
have**, so it stays UNMEASURABLE — now unmeasurable *with contrary evidence on both sides*, which is a
better epistemic position than the corpus had, and is reported as such. The deferred-by-rule boundary is
respected, not argued past.

## iter-91's fence caught its own author within the hour

The map edit above first went in with a citation reading `` `main.tf:39` `` instead of
`` `cms/terraform/main.tf:39` ``:

```
platform_alignment_guard  CANNOT-CHECK  rc=2
  1 citation(s) could not be resolved at all (main.tf) — each is a claim this run did not check.
```

**Before iter-91 that would have been printed as `unresolvable` and the guard would have exited 0 GREEN** —
the map would have shipped a dead citation under a green fence. First live catch of the third verdict, and
a positive control that arrived for free: proven discriminating on real prose, by nobody's construction.

## Close — 2026-08-05

**Outcome:** the M810 sweep was measured as already complete and NOT re-run; the two real residual defects
were found and fixed instead — a fenced claim restated more strongly in unfenced prose (`backend.md`,
`CLAUDE.md`), and a cms M810 step from 2026-08-04 the corpus had not seen, which points opposite to what it
recorded. The prod-side state is named UNMEASURABLE with evidence on both sides rather than asserted.
**Type:** tik
**Status:** closed-fixed — the declared scope landed; the substituted target was the right one
**Gate:** NOT MET — **4 of 5, unchanged.** No reading taken.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-92-1 … D-M257x-92-4 (iter-92/decisions.md)
**Side-deliverables:** none.
**Routes carried forward:**
- `CHECK-M257x-iter92-fenced-claim-restatements` → the class `D-M257x-92-2` names: find where else a fenced
  document's hedge is dropped by its restatements. `claim_twin_guard` already fences *adjudicated* claims
  across the tree; this is the adjacent case of a **fenced hedge** being restated flat, which nothing checks.
- `CHECK-M257x-iter91-claim-twin-answer-key-stale` → **still open**, and iter-92 sharpened why: the claim it
  fires on (C-2) asserts the `cms`/`jobsimulation` husks still run and import `colony/authn`, which the
  `838d907`/`0c91421` move ended. The answer key is stale for the same reason this iter existed.
- All iter-90/91 CHECKs remain open as recorded.
**Lessons:**
- **A task description is a claim too — measure it before executing it.** The brief's "~14 passages" was
  stale; acting on it would have produced motion and no repair.
- **Fencing a document does not fence its paraphrases.** A hedge that survives only where a guard reads it
  is a hedge that has already failed.

**Type:** tik

# iter-26 — clause 2 moves for the first time: `20 / 10 / 1` → `23 / 7 / 1`

## The measurement

Full suite, `--reset`, from the pinned consumption clone, on the stack carrying iter-24's Directus re-point.
ptreport's own verdict:

```
Playthroughs coverage: 23/31 passing (74.2%)
  passing=23  failing=7  unimplemented=1  unimplementable=0
ptreport: GATE no-regressions FAILED (a Playthrough is failing)
```

Checked in at `evidence/pt-summary-iter26.txt` (the full 31-row ptreport table). The raw runner log is at
`evidence/pt-run-iter26.log`, which `.gitignore:89` (`*.log`) excludes — so the **summary** is the durable
artefact and the log is local-only. Worth knowing before citing a `.log` path in this corpus as evidence.

## The diff, which is the actual result

`20/10/1` → `23/7/1` could be three fixed and three different ones broken. It is not — the sorted failing
ids `diff` cleanly against iter-19's ten, **three removals, zero additions**:

```
< hiring.recruiter-comparison.UC1
< skill-paths.legacy.UC1
< skill-paths.save-for-later.UC1
```

**Both skill-path Playthroughs flipped to passing.** They are the two whose failure text was the
`directus_versions` 403 — the class iter-24 traced to `backend` reading *prod* Directus anonymously and
closed by adding `backend` to `DIRECTUS_DATA_CONSUMERS`. The gate's own instrument now confirms, end to end,
what iter-24 proved at the HTTP layer.

`hiring.recruiter-comparison.UC1` flipping is **not claimed as understood.** It is consistent with the same
cause (the recruiter's scoreboard renders sim metadata that comes from the content layer), but nothing was
measured to establish that, and this milestone has already been burned once by reasoning from "the layer I
changed was broken" to "therefore this symptom was downstream." Recorded as an observation with an open
question, not as an attribution.

## Validating the number before quoting it

Two preconditions, both **checked rather than assumed** — iter-25 had flagged the second as a reason to
discard this run entirely:

1. **Full run?** The ptreport gate is binding only on a full run (`run-playthroughs.sh:300-307`). 31 rows, no
   scoping flag, gate verdict emitted → yes.
2. **Is the stale roster material?** The run's `stackseed --roster-export` failed (iter-25's second-pass
   defect), so the fake-FAPI served the *previous* seed's roster against a freshly reset-and-reseeded DB. If
   hero ids had moved, every *"the seeded hero is among the results"* assert would fail for a reason that is
   not the product — and **four of the seven survivors are exactly that shape**, so this was not a
   formality.

   Measured: the roster's `pt-employee` entry carries `eid=23f24e3f-38fb-5027-9e07-2ef49a644af5`; after the
   reset+reseed, `select id,email from public.users where id='23f24e3f…'` returns
   `pat.ellis1@pt-meridian-labs.com`. The seed is deterministic, so the stale roster is identical in the
   load-bearing field. **Confound defused by a query, not by an argument about UUID versions.**

## The seven that remain

| id | shape |
|---|---|
| `workforce-intelligence.organization-feedback.UC1` | *"the seeded hero (Pat Ellis) appears among the aggregated feedback"* — got nothing |
| `workforce-intelligence.skills-funnel.UC1` | *"her card carries her seeded role (DevOps Engineer)"* |
| `workforce-intelligence.talent-pool.UC1` | *"the succession/at-risk projection names the seeded hero"* |
| `assignment-monitoring.assign-and-track.UC2` | *"the org's seeded hero is among the per-member results"* — got 0 |
| `assignment-monitoring.assign-and-track.UC1` | the assignable-affordance count does not drop by one |
| `onboarding.enterprise-hiring.UC1` | her assigned position does not render as a startable org-scoped link |
| `org-admin.roles.UC1` | `page.waitForURL` timeout, 60 000 ms |

The first **four** share a signature — *a manager-vantage read returns the seeded hero as absent* — which is
`CHECK-M257x-iter15-manager-reads-empty` (iter-15 counted five; the fifth,
`workforce-funnel`, now passes). **Four of seven is one cluster**, and with the roster confound eliminated it
is a real product/seed-data question rather than an instrument artefact. That is the next tik's target and
the single highest-value item left on clause 2.

## Close — 2026-08-01

**Outcome:** clause 2 moves for the first time in the milestone — `20 live / 10 failing / 1` →
**`23 live / 7 failing / 1`**, three removals and **zero** additions on a sorted-id `diff`, with both
`directus_versions`-403 skill-path Playthroughs flipping green. Still **NOT MET** (the gate wants 30/0/0).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close) — (4) user-blocker: n — (5) cap-reached: **n on the numeric cap — 4 tiks
closed this run (23–26), not 5; iter-22 belongs to the PREVIOUS run and its commit `f0eb176` is this run's
starting point** — (6) protocol-stop: n — Outcome: **exit-5 by session budget, not by count**. Recorded
this way deliberately: `cap-reached` is the enum value the orchestrator parses, and there is none for
"session budget exhausted", but writing "5 tiks" to make the exit look mechanical would be the same
unverified-count error this milestone exists to eliminate. The cap was NOT hit; the session was.
**Decisions:** D-M257x-26-1 (this iter's `decisions.md`)
**Side-deliverables:** none.
**Routes carried forward:**
- **`CHECK-M257x-iter15-manager-reads-empty`** — now **4 of the 7** remaining failures and one coherent
  signature. Roster confound eliminated, so it is a genuine data/read question. Next tik.
- `onboarding.enterprise-hiring.UC1` · `assign-and-track.UC1` · `org-admin.roles.UC1` — three singletons,
  each its own cause; do not batch them on the strength of being "the rest."
- **Open question, not an attribution:** why `hiring.recruiter-comparison.UC1` flipped.
- A confirming re-run at `fast-build-m257x-iter-25b` — now **low value** (the confound it would rule out has
  been ruled out directly), but it would exercise the second-pass runner fix end to end.

**Lessons:**
1. **Re-survey before discarding, not just before targeting.** iter-25 routed this measurement forward as
   un-landed; the run had in fact completed minutes later. The Step-0 re-survey is normally used to check a
   *target* is still meaningful — it is equally good at finding that a *deliverable* already exists.
2. **Defuse a confound by querying it.** The stale-roster worry was a good one and would have justified an
   hour-long re-run; one `select` on the id the roster names settled it in seconds. Ask what observation
   would distinguish the two worlds before paying for the safer one.
3. **The diff is the result, the totals are not.** Three removals and zero additions is a different claim
   from "the number went up," and only the first supports attributing the lift.

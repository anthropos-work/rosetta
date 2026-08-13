**Type:** tik — under `TOK-08`, corpus half of the user redirect.

## Phase 1 — sealed

Predictions `P-227-1..4` sealed before any repo `CLAUDE.md` was read. Subject: the 6 archived repos, all
at origin tip since iter-224.

## Phase 2 — the census

Each repo's own `CLAUDE.md` read at `origin/main`:

| repo | what the repo says about ITSELF | verdict |
|---|---|---|
| `messenger` | **⚠️ FROZEN — this service no longer runs** | **agrees** |
| `storage` | **⚠️ FROZEN — the service no longer runs, but the terraform module is LIVE** | **agrees** |
| `cms` | *"Content layer for the platform: serves job simulations… **via GraphQL Federation**"* — **no status statement anywhere** | **corpus-ahead** |
| `jobsimulation` | a live-service doc; **no status statement**, though `6092c6d2` destroyed its ECS service and ECR repository | **corpus-ahead** |
| `graphql-wundergraph` | *"**it is currently the sole subgraph**"* — a live router, **no mention** that platform `2adcf71` deleted it; also carries the **`2 → 1`** figure this corpus corrected to **`3 → 1`** | **corpus-ahead ×2** |
| `roadrunner` | *"Sandboxed code execution service… **Deploy: Docker -> ECR -> ECS**… used exclusively by `jobsimulation`"* | **corpus-ahead** |

**`P-227-1` REFUTED — 2 of 6, not ≥ 4.** Only `messenger` and `storage` carry a status banner. The freeze
wave of 2026-08-05 reached those two and stopped.

**`P-227-2` REFUTED, and the direction is the finding.** Not one of the six states a status this corpus
does not already hold. The disagreements are real and **all four run the other way**: the repo describes a
running service and the corpus correctly contradicts it.

**`P-227-3` HOLDS.** `roadrunner` and `graphql-wundergraph` — the two that did not advance — carry no
freeze commit; `roadrunner` returned **zero** hits for any status vocabulary at all.

**`P-227-4` HOLDS.** Zero cases of the corpus being wrong on direction. Every gap is the repo being behind.

## Phase 3 — what this does to iter-224's lesson

iter-224 closed with *"read the repo's **retraction surface** — `CLAUDE.md`, `README.md` — not only the
anchor."* This census says **that method has a two-in-six hit rate here.** For four of the six there is no
retraction surface to read: the repo's own doc is a live-service doc.

The trap is aimed at exactly this corpus's reader — **an agent that clones `cms` or `roadrunner` to check
a claim will be told the service is live and deployed to ECS.** The authorities remain the platform's
`repos.yml` and `infrastructure`'s `services.tf`, which is what [§4 The fence](#4-the-fence) already
asserts mechanically.

## Phase 4 — repair, and the fence catching this iter twice

Landed as a boxed section above `## 1. How to read a row` in `platform-migration-status.md`, carrying the
6-row table and the two-in-six qualification.

**`guard_family` went RED on this iter's own edit, twice, and both were correct:**

1. **`derived_count_guard` arm D** — the heading claimed *"4 of 6 say they are live"*, an `N of M` count on
   the clause-5 surface with **no entry in `N_OF_M_DISPOSITIONS`**. The registry lives in rext, so adding a
   disposition would have meant an rext commit **plus a push** for a piece of prose. Rewritten to *"four of
   the six"* instead — the guard's pattern is `(\d[\d,]*)\s+of\s+(\d[\d,]*)`, digits only, and the spelled
   form is better prose anyway.
2. **`repair_postcondition`** — cleared with the same edit.

**And the insertion moved 27 lines out from under three citations**, which the fence family does *not*
check for an un-ranged run and which iter-224's own Lesson names. Found by grepping before committing,
each verified against the row it names, all three re-pointed: `dependency_map.md` `:88 → :115` (the `app`
row), `service_taxonomy.md` `:89 → :116` (`cms`), `services/README.md` `:101 → :128` (`intelligence`).

## Close — 2026-08-09

**Outcome:** the six archived repos were asked what they say about themselves. **Two carry a freeze
banner; four still describe live, ECS-deployed services** — none contradicts this map on direction, but
"check the repo's own docs" is a method with a two-in-six hit rate, and that is now written down where a
reader will hit it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-227-1` (a refuted prediction whose refutation IS the deliverable),
`D-M257x-227-2` (prose rewritten to clear a fence, rather than the fence's registry widened).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Predictions, graded:**

| id | prediction | result |
|---|---|---|
| `P-227-1` | ≥ 4 of 6 carry a status/freeze statement | **REFUTED — 2 of 6** |
| `P-227-2` | ≥ 1 repo states a status the corpus's row lacks | **REFUTED — 0; but 4 state a LIVE status the corpus correctly contradicts** |
| `P-227-3` | `roadrunner` + `graphql-wundergraph` carry no freeze commit | **HELD** |
| `P-227-4` | 0 cases of the corpus being wrong on direction | **HELD** |

**Suite state at close** — `guard_family` with `--platform stack-demo/platform`: **24 GREEN · 0 RED · 5
not-run** after repair (it was **22 GREEN · 2 RED** on the first run of this iter's edit — both REDs this
iter's own, both repaired). Not a whole-family green: the 5 not-run are commit/ledger-scoped members with
no input supplied. No pytest section run; this iter changed no rext code.

**Routes carried forward:**
- `ROUTE-M257x-227-archived-repo-selfdesc-is-stale` → **not ours to fix.** Four `anthropos-work` repos
  describe themselves as live. Repairing them is a platform edit, which this milestone does not do
  (escalation condition, stated in this iter's `overview.md`). Recorded in the corpus instead so a reader
  is warned; if it is ever raised upstream, this table is the evidence.
- All prior routes (`225-no-profile-for-sanctioned-host`, `225-profile-vs-host-identity-check`,
  `225-hostprofile-role-strings…`, `226-build-budget-argues-for-a-retired-host`,
  `224-drift-guard-blind-to-stale-clone`, `222-pin-advance-needs-a-reproof`,
  `223-classify-the-ten-drifted-baselines`) → open, unchanged.

**Lessons:**
1. **A refutation can be the deliverable.** `P-227-2` predicted the messenger defect would recur and it
   does not — nothing else in the corpus quotes a claim its repo has retracted. **The messenger defect is
   bounded to one site**, which is worth more than another repair would have been, and could only be known
   by asking all six.
2. **"Read the source's own docs" needs a measured hit rate before it is offered as a method.**
   iter-224 generalised from one repo; the census puts it at two in six. A method stated without its
   reach is the same defect this milestone books against numbers.
3. **When a fence REDs on a prose construct, changing the prose can beat widening the registry.** The
   disposition registry is rext code — widening it means a commit and a push to origin for a sentence.
   Spelling the number out cleared it, and reads better.
4. **The guard family does not check citation offsets on an un-ranged run.** `anchor_offset_guard` needs
   `--range`, so a 27-line insertion showed **24 GREEN** while three citations pointed at the wrong rows.
   Grep for citers before committing any insertion — the family being green is not the same as the tree
   being right.

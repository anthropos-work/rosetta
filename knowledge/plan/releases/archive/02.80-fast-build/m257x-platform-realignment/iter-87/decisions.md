# iter-87 — decisions

## `D-M257x-87-1` — the clone-advance rule: FETCH ALL; ADVANCE ONLY WHAT A DERIVED SET READS

The hand-off asked for a **stated rule** rather than a case-by-case call, and anticipated that advancing
`app` (93 commits behind, 65 corpus citations) would be "a large wave better run as its own iteration."

**Rule:**

> Every clone is **fetched** in any iter that takes a measurement. A clone's **checkout** is advanced only
> when the checked-out tree is an input to a derived legal set or to a build.

**It is derived, not preferred.** The measurement is in the guards' own source:

- `anchor_construct_guard.resolve()` and `platform_alignment_guard.cited_text()` both default to
  `CITE_REF=auto`, whose ladder is **`origin/main` first**, checkout second — landed at iter-68 as
  `FENCE-M257x-iter68-citation-resolution`, precisely because *"the demo stack pins its clones to a build
  tag while the exit gate names origin HEAD."* A **fetched** clone is therefore already graded at origin
  HEAD regardless of where its HEAD sits.
- Only two things read a checked-out tree: `platform_predicate_guard` (the dir passed to `--platform`) and
  `platform_alignment_guard` assertion B (the `repos.yml` path passed as argv). Both are **`platform`**.

**Consequence — the anticipated deferral dissolves.** Advancing `app` would move nothing, because no fence
reads `app`'s checkout. The 65 `app/*:N` citations were **already being graded at origin HEAD in this
iter's reading**, and the whole 93-commit exposure surfaced as **2** RED anchors (`app/main.go:471`,
`:637`), both repaired here. The wave was not deferred; it was already in the measurement.

**What IS deferred under the rule:** nothing citation-bearing. `stack-demo/rosetta-extensions` (+34) is a
**pin**, not a citation target — advancing it is a bring-up act governed by §7 rule 4 and belongs to a
clause-1/clause-2 iter. **Fate 2** (already covered by clause-1/2 work).

Generalised into the protocol doc as **§5 rule 41**.

## `D-M257x-87-2` — the re-scope trigger is graded NOT FIRED, and the count it is graded against was wrong

The hand-off carried *"occurrence 1 of 2"*, sourced from `state.md:6`. **Re-derived: that field was 73
iters stale and the trigger had already fired.**

| # | date | commit | outcome |
|---|---|---|---|
| 1 | 2026-07-31 | `2adcf71` (WunderGraph router deleted) | recorded, iter-12 |
| 2 | 2026-08-03 | `ef32d4c` (cms/jobsim/roadrunner pruned) | **FIRED**, iter-53 — `D-M257x-53-6`: *"Occurrence 2 of 2. `EXIT_REASON` corrected from `user-blocker` to `re-scope-trigger`."* Escalated; remedy built as **TOK-04** |
| 3 | 2026-08-05 | `0c91421` (support containers dropped) | **this iter** |

**Grading: NOT FIRED**, on two independent grounds.

1. **The predicate is false on its own words.** The condition is *"TWO **CONSECUTIVE** full-alignment
   attempts are invalidated."* Occurrences 2 and 3 are separated by **33 iters during which no platform
   commit landed at all** (the platform's own log is empty between 2026-08-03 and 2026-08-05). Two
   invalidations separated by 33 clean iters are not consecutive.
2. **The prescribed remedy already exists and performed on this exact event.** The trigger says the answer
   is *"a pinning-and-tracking POLICY (how we choose a platform ref, how we notice it moved, who
   re-points), not more alignment work."* That policy is TOK-04 P1/P2/P3. On occurrence 3 it delivered:
   the move was **noticed by a fence within hours**, the refs are **stated in the artifact**
   (`ground-truth.md`), and **the detecting iter re-pointed it** — this one. Firing a trigger whose remedy
   is in place and working would be ceremony, and would replace working machinery with a request to build
   it again.

`state.md` repaired: it now spells the history out rather than carrying a count, and says why.

## `D-M257x-87-3` — the M810 sweep lands NOW as a side-deliverable, rather than routing forward

Packet E surfaced (and I independently re-derived) that **platform M810 has already landed for
`jobsimulation`** — `6092c6d2` deleted the `module "jobsimulation"` block outright, destroying the ECS
service, task definition and ECR repository; `jobsimulation/terraform/main.tf` is 56 lines at `origin/main`
and contains no `service_desired_count` at all. **`cms` has not moved** (`cms/terraform/main.tf:39`, still
`= 0`). The corpus asserted M810 as future work in ~14 passages across 11 files.

This is **not** in this iter's planned scope, so the scope-creep tripwire applies. It is landed anyway, as
a **separate concern with its own commit**, because of §5 rule 19: the map — packet E's file — had already
been repaired to the correct verdict, and **half-repairing a uniformly-wrong corpus is worse than leaving
it alone.** Routing forward would have shipped exactly the state that rule forbids: *a claim corrected in
the file its owner held while the identical claim stood in a twin owned by somebody else.* Fate 1.

It does **not** upgrade the iter's close status, which grades planned scope only.

## `D-M257x-87-4` — the family runner's RED headline is derived from the producer's ordering

`guard_family.py` reported `lines[-1]` as each guard's headline. On the run that detected this platform
move, `platform_alignment_guard` went RED with 21 findings whose **first two** were the `[B departure]`
lines naming the two services that had just left `repos.yml` — and the family view showed a `gotenberg`
citation nit from the bottom of the list. **The instrument caught the event and its summary hid it.**

Repaired to *"N finding(s); first: …"*, with the finding set derived structurally (guards print summaries
flush-left and prefixed with their module name; findings are indented or bulleted), so no per-guard list
has to be kept in step — §2's rule against hand-maintained lists, applied to the runner built to enforce
§5 rule 8. Falls back to the last line when a RED itemises nothing, because swallowing that would be the
silent-skip this runner exists to refuse. +5 tests. Generalised as **§5 rule 42**.

## `D-M257x-87-5` — three hand-off figures re-derived; two confirmed, one refuted

§5 rule 32 (*re-derive the hand-off's numbers, including the orchestrator's*), applied and earning again:

| figure | source | verdict |
|---|---|---|
| `repos.yml` 6 → 4, compose 8 → 5, topology 10 → 7, 3 profiles gone, `STORAGE_S3_BUCKET` persists at `:82` | hand-off | **all confirmed** |
| `app` is ~93 behind (iter-86 said 60) | hand-off | **confirmed 93**; iter-86's 60 was a day and 33 commits stale |
| guard family at the old ref was *"13 GREEN · 0 RED · 3 not-run"* | hand-off | **refuted — 10 GREEN · 3 RED · 3 not-run** at the identical checkout, after a fetch. See `D-M257x-87-1` and rule 41: the reading was taken against an **unfetched** clone, and a citation fence resolving at `origin/main` reads GREEN until you fetch |

And one of my own inputs was refuted by a packet: my verdict sheet said `docker-compose.yml` went
**293 → 186**; measured, it is **271 → 186**. Recorded because a repairer catching its own briefing's
arithmetic is the property the packet design is for.

## `D-M257x-87-6` — the §2 time bomb is retired: it did not fire, and the derivation is why

§2 forecast that *"the day they leave the clone set, `[ -d ] || continue` silently skips them, both schemas
become empty shells, and **13 write targets fail with 42P01 at once**."* `838d907` is that day — `storage`
and `messenger` left `repos.yml`. Measured across the move:

| | `0dab54d` | `0c91421` |
|---|---|---|
| `repos_yml_migration_pairs` | `app:public` | `app:public` |
| `repos_yml_schemas_to_create` | `extensions sentinel public` | `extensions sentinel public` |

**Identical, and identical correctly, with zero human action.** The hand-maintained tuple deleted at
iter-02 would have skipped both repos on this commit. This is the **third** consecutive platform change the
derived layer has absorbed unaided, and it is the strongest evidence yet for P4's *derive → fence →
declare* ordering. Recorded in §2 and §9; the emptied debt list and its shrink-fence are **kept**.

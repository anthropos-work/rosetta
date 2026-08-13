# Where the milestone actually stands, clause by clause

**Taken 2026-08-03, against platform origin HEAD `ef32d4cd8e0ceecf528a74c37d5e2ae5804ce021`.**
Job 2 of the three the user asked for. Job 1 is [`platform-before-after.md`](platform-before-after.md);
job 3 is `TOK-04` in the milestone-root `decisions.md`.

The headline first, because it moves in the unwelcome direction and burying it would be the exact failure
this milestone exists to end:

> **The gate has been booked at 4 of 5 since iter-40. Verified against the ref the gate itself names —
> origin HEAD — it is 2 of 5 today.**
>
> Nothing regressed. Clauses 1 and 2 were met, honestly, against `2adcf71`; origin HEAD is now `ef32d4c`,
> and the gate's own wording is *"against platform @ **origin HEAD** (never a pinned pre-drift commit)."*
> Two clauses went stale by definition, not by failure. Both are cheap to restore. Clause 5 is the one
> that genuinely moved away.

| clause | booked | today | cost to restore |
|---|---|---|---|
| 1 — three green cold cycles | MET (iter-14, re-met iter-18) | **STALE** | ~35 min |
| 2 — full Playthrough run `30 / 0 / 0` | MET (iter-37) | **STALE**, and it recorded no ref | ~5 min |
| 3 — the fenced migration-status map | MET | **MET**, re-met today, one caveat | — |
| 4 — zero rext writes to a dead schema | MET | **MET — and now met *under test*** | — |
| 5 — KB-fidelity GREEN / YELLOW-0-blockers | NOT MET | **NOT MET, and further away** | see below |

---

## Clause 1 — three consecutive green cold cycles. **STALE.**

Met at **iter-14**, explicitly *"against platform origin HEAD `2adcf71`"* (`iter-14/progress.md:19`), three
`demo-down --purge` → `demo-up` cycles at `11:43:02Z` / `11:53:04Z` / `12:03:33Z`, `warnings:0 / green:true`,
~11 min each. Re-met at **iter-18** with an instrument that could see served content.

It is stale, and not merely on a technicality:

1. **The topology it measured no longer exists.** `d11a403` deletes the **cms**, **jobsimulation** and
   **roadrunner** compose services. The green was measured on a stack that started three more containers
   than the current one starts. The same commit had to add `--remove-orphans` to the `up` targets precisely
   because those containers otherwise linger — i.e. the platform itself treats the transition as one that
   changes what a bring-up produces.
2. **The bring-up timing contract changed.** `6060315` gives postgres `start_period: 120s` (`common.yml:22`)
   because permission re-application on a grown data dir outlasted the 25 s the retries allowed. Clause 1's
   ~11-minute cycle is a number about a different startup.
3. **Even a re-run here is not a cold-box proof.** `stack-demo/{cms,jobsimulation,roadrunner}` still exist
   on disk from before 2026-08-03, so no local run exercises a genuinely fresh `make init` against this
   HEAD. That residual is real and should be stated in whatever re-run happens, not quietly inherited.

**Expected outcome of a re-run: green** — because clause 4's derivation is ref-independent (below) and
`d11a403` *restored* env onto `backend` rather than removing it. Stating that in advance makes the re-run
refutable rather than confirmatory.

## Clause 2 — the full Playthrough suite, `30 live / 0 failing / 0 error`. **STALE — and its own record does not say against what.**

Met at **iter-37**: `29 / 1 / 1` → `30 / 0 / 1`. The trailing `1` is the declared in-manifest TODO, not an
error, so the clause's `0 error` is satisfied as written.

**But `iter-37/progress.md` contains no platform sha at all.** The only sha-shaped token in the entire file
is `ad524614`. Its platform ref is *inferred* from iter-36, which did re-fetch `2adcf71` at open and close
and said so. So the milestone's gate-meeting clause-2 measurement is one whose target ref exists only by
adjacency to the iter before it.

**That is the fourth instance of this tok's class, and it is inside the gate itself** — see TOK-04.

Substantively, the exposure is real but points toward green: Playthroughs drive the product through a
browser against `backend`, and `d11a403` moved `JUDGE0_BASE_URL`, `DIRECTUS_PUBLIC_BASE_ADDR`,
`REDIS_WORKER_INDEX`, the LiveKit and Chime blocks and the `~/.aws/credentials` mount **onto `backend`** —
code-exec, content and recording paths, a large share of what the suite exercises. Env restored, not removed.

**This is the cheapest clause in the gate — 4 min 50 s wall including the reset (iter-32's measurement,
against an inherited "~35–40 min" estimate carried unchecked across seven hand-offs).** It should be re-run
first, before anything else in the milestone.

## Clause 3 — the checked-in, machine-fenced migration-status map. **MET. Re-met today. One caveat that matters.**

Run against the new HEAD **before any corpus edit**, the fence returned:

```
platform_alignment_guard: 3 finding(s)
  [B departure] the map claims cms is in repos.yml, and it is not
  [B departure] the map claims jobsimulation is in repos.yml, and it is not
  [B departure] the map claims roadrunner is in repos.yml, and it is not
EXIT=1
```

After the map was updated: `EXIT=0`, re-verified again after this iter's further prose corrections.

**This is the first time a fence this milestone built caught a real departure it was not shown** — every
prior RED was a defect deliberately staged to watch the guard fire. Direction B is the direction that fired
in anger in all three historical occurrences (skiller, skillpath, jobsimulation) and it fired again,
correctly, unprompted, within hours, on a tree nobody had touched.

**The caveat, and it is not a small one.** This very iteration committed a **false claim** into that map:
*"the armed failure is now armed"*, citing `demo-stack/migrate-demo.sh:81-85` and `:106` — line anchors into
code that rext `54bccf7` (M257x **iter-02**) deleted. iter-01's finding was quoted forward without
re-measuring against this milestone's own repair. The guard cannot see it, and says so in its own header:

> *"This guard cannot tell you whether a cited sha says what the row claims it says. It checks the one thing
> that is mechanically checkable and that actually broke: who is in the clone set… Everything else in the
> map is prose under human review."*

A second false claim was found in the same pass and fixed: §5's "rows to watch" gave a **dead signal** —
*"when `repos.yml` flips `storage`/`messenger` to `migrations: false`, the fold has landed"* — when both have
read `migrations: false` since long before the fold was announced (`repos.yml:18-23` @ `ef32d4c`). The map
committed the very Trap A its own §1 exists to warn about.

**Honest reading of clause 3:** *membership is fenced and correct; the prose around it produced two false
claims inside one working day, one of them written by the fence's own milestone.* Met — but "the map is
fenced" and "the map is true" are different statements, and only the first is mechanically held.

## Clause 4 — zero rext writes to a schema the platform no longer creates. **MET — and this is the strongest result the milestone has produced, because nobody set it up.**

Verified **live**, not by inspection. Sourcing `stack-core/lib/repos_yml.sh` against `repos.yml` @ `ef32d4c`:

```
migration pairs:     app:public
schemas to create:   extensions  sentinel  public
transitional debt:   (empty)
```

Byte-identical to the reading at `2adcf71`, and identical **correctly**: the three departing repos declared
`migrations: false` and no `schema:` key, so a set derived from those two fields never named them. The
removal passed through the derivation as a no-op *because the derivation is a derivation.*

The counterfactual is concrete and dated. Before rext `54bccf7` (iter-02) the same file carried a
hand-maintained `app:public cms:cms jobsimulation:jobsimulation skillpath:skillpath` tuple behind a silent
`[ -d ] || continue`; iter-01 predicted the removal would make it *"silently skip … and 13 write targets
42P01 at once."* iter-02 replaced the skip with a loud `mig_fail=1` refusal, and iters 06/07 re-pointed the
last cms/jobsimulation writes so `REXT_TRANSITIONAL_SCHEMAS` could go empty. **The condition the time bomb
was waiting for arrived on 2026-08-03, six weeks ahead of the M810 everyone expected, and nothing happened.**

**Clause 4 has moved from met-by-construction to met-under-test, with zero human action in between.** That
is the milestone's founding thesis surviving an event it did not arrange.

Two limits, stated rather than glossed:
- This proves the **derivation**, by running it against the real new `repos.yml`. It does not prove a **cold
  bring-up** — that is clause 1's job.
- Our clones of the three departed repos still exist, so the fresh-`make init` path remains unexercised.

## Clause 5 — KB-fidelity GREEN, or YELLOW with 0 blockers. **NOT MET, and the distance grew today.**

**The series is non-comparable, and that is now settled rather than suspected.** `25 → 13 → 11 → 17 → 37 →
18 → 7 → 12 → 14 → 7` was produced by an instrument that was **never frozen**: the briefing that *is* the
instrument lived at a git-ignored scratch path and was re-authored from a summary on every pass, with six
knobs drifting — the load-bearing one an **inverted tie-break** (canonical: *"if you cannot cite the
refutation, it is not a blocker"*; iter-53 as-run: *"when in doubt, book it as a BLOCKER"*). A rule that
resolves doubt upward cannot produce a number comparable to one that resolves it downward. The canonical
briefing is now committed at [`instrument/`](../instrument/), with the as-run drift preserved beside it.

**What survives the instrument problem — the only instrument-independent fact — is recall ≈ 43–48% per
7-seat pass.** Two independent readings union to roughly two-thirds of what is there.

Latest paired reading: as-run **32 & 26** (union **46**, N̂≈68); canonical re-grade **23 & 23** (union **35**,
N̂≈47). *Which of 46 or 35 is the working residual is a user decision, `D-M257x-53-5`, still open.*
**9 of the 46 were induced** by iter-52's own 18-claim repair — a ~50% induction rate.

**And then the platform moved.** Sites in `corpus/**` + root `CLAUDE.md` asserting the now-false shape
(*"the container still starts"*, `running_but_unfederated`, *"`repos.yml` still lists"*, and the stale
`docker-compose.yml:{83,144,281}` anchors): **81, across 21 files** — 20 of the 21 inside the 40-file
partition the readings sweep. These are not the same claims re-mis-read; they are the same **files** newly
falsified. The vocabulary itself died: `running_but_unfederated` was coined at iter-20 for exactly these
three services and now describes none of them.

**Plus one more, today: the over-claim this iter committed and corrected** (clause 3 above). A defect induced
by the milestone's own repair activity, in the map that is clause 3's deliverable, caught by re-measurement
and by no reading.

### How they interact — the arithmetic that decides TOK-04

**Three platform commits, landed inside one working day, created a drift surface larger than the entire
union that ten readings have been trying to close.**

| term | rate |
|---|---|
| repair | ~18 claims per repair-iter |
| induction | ~9 per 18 repaired (~50%) |
| external drift | **81 in one working day** |

A pass that repairs 18 while inducing 9, against a platform that adds 81, is **net −72**.

**The clause-5 residual is not a stock being drained. It is a balance, and it is currently negative.** Ten
readings have been optimising the wrong term. That is not a criticism of the readings — the drift rate was
not visible until an event made it visible, and today was that event.

**Clause 5 is not re-cut.** The user has ruled three times. It is met only by a reading that returns zero.
TOK-04 changes the method and the accounting; it does not touch the bar.

---

## Summary for the orchestrator

- **2 of 5 verified at origin HEAD** (clauses 3 and 4), not 4 of 5.
- **Clauses 1 and 2 are stale by the gate's own wording**, not failed. ~40 minutes of machine time restores
  both, and the expected result is green — pre-registered here so it can be refuted.
- **Clause 4 is the milestone's best result**: it was tested by an event nobody arranged and passed with
  zero human action, because it is derived rather than maintained.
- **Clause 5 is further away than it was yesterday**, by 81 sites, and the reason is external to every
  method the last ten readings tried.

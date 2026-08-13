# M258 iter-08 — decisions

## D29 — the auto-arming waiter caught a window manual polling would have missed

iter-07 concluded *"the window never opened"* from ~30 minutes of **hand polling at ~2-minute
granularity**, minimum `load1` 11.93 against a limit of 10. iter-08 replaced the human trigger with a
continuous sampler (`autoarm-campaign.sh`: 15 s interval, fire on 3 consecutive samples ≤ 5.0).

**It armed at 08:17:41Z and fired at 08:19:12Z — 91 seconds later**, on a decay the sampler watched
happen: `7.88 → 6.98 → 6.34 → 5.37 → 4.56 → 4.11 → 3.87`. The dip from 7.88 to under 5.0 took **75
seconds**; at iter-07's 2-minute hand-poll granularity it is a coin flip whether that window is seen at
all, and iter-07's own series (`07:33 11.93 · 07:34 32.90`) shows the box moving by 21 points inside one
minute.

**The lesson is not "iter-07 lacked patience".** It polled for 30 minutes. The lesson is that **a
trigger sampled by hand cannot catch a window shorter than its own interval**, and on this host the
windows are shorter than the interval. Automating the trigger — not extending the wait — is what
converted "no window" into a launched campaign.

Firing is cheap by construction, which is what makes an aggressive threshold safe:
`launch-iter07-campaign.sh` asserts the user's stacks resident and asserts headroom **before** the
teardown, so a spike between arm and launch exits 8 with `demo-1` intact. The waiter re-arms on rc ∈
{8, 9} rather than giving up.

## D30 — the contention returned 60 s after launch, and it is named

The campaign passed `assert-headroom` at **load1 3.87** (08:19:12Z) and completed its teardown at
`load1 4.44`. By 08:21:02Z the top of `ps -r` was **three `a8-cart-runner` processes at 94–97 % CPU**
plus two Python workers, all from `hyperspace/anima8/` — the *same third-party project* iter-07 named,
resuming its next batch roughly one minute after this campaign started.

So the 08:11–08:19 lull was **a gap between anima8's iterations, not the end of its campaign.** This is
recorded because it changes what the run can be read as: rep-01's `samples.tsv` carries `load1` of
**31.39 / 35.46 / 40.09 / 39.82** during the bring-up, against the instrument's ceiling of 10. The reps
are **functionally valid and timing-invalid**, and those two halves must be reported separately.

## D31 — `D27`'s arithmetic prediction was CORRECT: `autoverify.green` flipped `false → true`

iter-07 refused to spend a 3-rep campaign testing iter-06's routed hypothesis (*"run from the
consumption clone"*) and reduced it instead to two path expressions plus one query, predicting that
`autoverify`'s `postgres-schemas` probe **would assert and would pass** from the consumption clone.

Measured on rep-01: `stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/autoverify.json` =
`{"project": "demo-1", "offset": 10000, "warnings": 0, "green": true, "ts": "2026-08-12T08:30:17Z"}`.

**`green: true`, `warnings: 0`** — up from iter-07's `{"green": false, "warnings": 1}`. The prediction
held exactly, and the field the campaign existed to move, moved. *Settle by arithmetic when you can.*

## D32 — the D17 attribution check PASSES, and the artifacts agree with each other

`D17` is the defect where a 166 s phase was reported inside a 2 s one **and the phase table still
summed** — *a table that adds up is not a table that attributes*. Checked on rep-01, from the anchors in
`cycle.log`:

| anchor | line | timestamp | derived duration |
|---|---|---|---|
| `▶ autoverify demo-1 (offset 10000)` | 506 | `08:30:12Z` (t=651.92) | **≈ 5.3 s** |
| `▶ batch gate: driving the full Playthrough suite` | 523 | `08:30:18Z` (t=657.24) | **≈ 153.8 s** |
| `P4_BRINGUP_END` | — | `08:32:52Z` (t=811.05) | — |

`autoverify` reads **≈ 5 s, not ≈ 154 s** — the two phases attribute **separately**, which is the
property `buildbench.py:148-151` was written to protect.

And the two artifacts **agree with each other independently** (the `D19`/`D20` relationship check, not
just the step check): `batch-gate.json` reports `batch_seconds 142 + restore_seconds 11 = 153`, and the
anchor-derived `batch_gate` span is **153.8 s**. Two instruments, one number.

Their timestamps also agree and are **fresh** — `autoverify` `08:30:17Z`, `batch-gate` `08:30:18Z`, one
second apart, both inside rep-01.

## D33 — the STALE-ARTIFACT trap this run walked past (the D19 hazard, one level up)

`demo-1`'s stack dir exists in **both** clones, and the campaign runs from the **consumption** one:

| clone | `demo-1/batch-gate.json` `ts` | what it is |
|---|---|---|
| `.agentspace/rosetta-extensions/` (authoring) | **`2026-08-12T06:54:15Z`** | **iter-06's**, 1.5 h stale, survived `--purge` |
| `stack-demo/rosetta-extensions/` (consumption) | `2026-08-12T08:30:18Z` | **this run's** |

Both say `verdict green · red_count 0`. **A reader who reached for the authoring path would have
reported a stale green as this campaign's result and been entirely consistent with themselves** — the
verdict matches, only the timestamp betrays it. This is the live instance of the carried route
`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`, and it is why every artifact read in this iter is
path-qualified and `ts`-checked rather than found by name.

## D34 — the red set stayed EMPTY under load1 31–40, which falsifies iter-07's stated fear

iter-07 declined to launch partly on the grounds that *"a batch run at `load1` 40–60 would have
manufactured FALSE REDS"* — the suite runs `retries: 0` by contract, so browser timeouts under load
become failures, and clause 4 escalates the red set to the user as a **product** verdict.

rep-01 ran its batch with `samples.tsv` recording `load1` **31.39 → 40.09** during the cycle. Result:

```
verdict green · red_count 0 · red_set [] · runner_exit 0
summary: passing 30 · failing 0 · unimplemented 1 · unimplementable-without-platform-edit 0 · total 31
```

**Identical to iter-07's clean B1 baseline** (`passing=30 unimplemented=1`). The suite produced no false
red at 3–4× its headroom ceiling.

**Stated with its limits:** this is **one** observation at 31–40, not a characterization, and it does not
license running the gate contended — the *timings* are still invalid, which is the whole reason the
ceiling exists. What it does retire is the specific claim that contention at this magnitude
**manufactures reds** in this suite. The concern was reasonable and is now measured; the honest record
is that the batch half is more robust to load than its timing half.

## D35 — the campaign RAN: 3/3 reps functionally green, 0/3 timing-usable

`buildbench run 1 --reps 3 --profile macmini --no-public-host` completed all three reps
(08:19:19Z → 09:02:12Z, `CAMP_RC=1`). **This is the first complete composed 3-rep cold campaign the
milestone has ever had.**

```
rep 1/3 total=811.06s up_rc=0 green=True headroom=FAIL isolation=OK phases=complete
rep 2/3 total=840.01s up_rc=0 green=True headroom=FAIL isolation=OK phases=complete
rep 3/3 total=859.06s up_rc=0 green=True headroom=FAIL isolation=OK phases=complete
  total cycle   p50 840.01s   min 811.06s   max 859.06s   (never a mean)
```

**Every functional column is green on every rep** — `up_rc=0`, `autoverify green`, `isolation OK`,
`phases complete`, `host identity MATCH {match: 3}`, and a `red_count 0 / passing 30 / failing 0` batch
each time. **Every timing column is refused** — `headroom=FAIL` 3/3 on the single clause `peak_load1`
(40.09 / 74.77 / 51.80, all over the limit of 10), so buildbench declares the campaign RED with
*"rep(s) [1, 2, 3] are not usable measurements."*

**840.01 s is therefore NOT the composed p50 the gate asks for.** It is the p50 of three reps the
instrument rejects, and quoting it as the milestone's headline number would be reporting a figure the
measuring device disowns.

**One number worth keeping:** the total-cycle spread is **811.06 → 859.06 = 1.06×**, across peak loads
spanning 40 → 75. `TOK-01` requires the spread be published beside the p50 because M256 escalated a
**2.04×** spread as evidence its suite was *"not decidable at n=3 on this host."* At 1.06× under
*adverse* conditions, the composed cycle does not show that pathology — the decidability caveat is not
retired (it was raised about the suite in isolation, and this is a different, larger quantity), but the
evidence points the other way.

## D36 — `set_dress` 81.23 → 283.53 s: the excess is real, its CAUSE is not established

The same command (`--reps 3 --profile macmini --no-public-host`) measured `set_dress` at **81.23 s
[78.92–117.51]** in iter-05 and **283.53 s [282.91–306.22]** here. Both are n=3, both stable. My first
reading — that `set_dress` is load-insensitive and therefore genuinely ~283 s — **was wrong, and is
retracted before it was written anywhere but this paragraph.** It rested on a within-campaign
observation (peak_load1 40/75/52 does not order 283.53/282.91/306.22, and rep-2 carried the *highest*
load with the *lowest* `set_dress`), which says only that the relationship **saturates above ~40** — not
that there is no relationship. The cross-campaign comparison is the one that matters:

| phase | iter-05 | iter-08 | ratio |
|---|---|---|---|
| `compose_up` | 43.87 | 59.60 | 1.36× |
| `ui_hiring` | 45.32 | 86.99 | 1.92× |
| `host_preflight` | 9.33 | 19.08 | 2.05× |
| `ui_studio_desk` | 7.12 | 17.16 | 2.41× |
| `ui_next_web` | 49.01 | 130.85 | 2.67× |
| **`set_dress`** | **81.23** | **283.53** | **3.49×** |

**Everything inflated** — cohort median **2.05×**. So the campaign as a whole paid roughly 2× for
running contended, exactly as the `headroom=FAIL` verdict says it should. `set_dress` paid **3.49×**,
which is **~117 s more than the cohort ratio predicts.**

Two hypotheses survive, and this campaign cannot separate them:

1. **`set_dress` is simply the most CPU-exposed phase.** Its dominant span is a `docker run … node
   cli.js bootstrap` plus a 330,261-row COPY that ends by **reindexing two pgvector indexes**
   (`public.skill_embeddings.small_embedding3`, `public.job_role_embeddings.small_embedding3`). Vector
   index construction is pure CPU; a super-cohort ratio is exactly what one would predict.
2. **Something changed between iter-05 and iter-08.** The campaign moved clones (authoring →
   consumption, per `D27`) and the consumption clone had no `stacks/demo-1/` at all, so rep-1 ran a
   colder start — including a Directus **structure auto-provision** (`schema b4cb55bcee08 →
   ea2e187a1605`, the M21 digest-miss path) that a warm stack dir may skip.

**It must not be resolved by preference, because the two answers point at different milestones' work.**
Under (1) the fix is a quiet host; under (2) `LEVER-M257-L5-setdress` has a much bigger target than the
~30–50 s the v2.8 plan priced, and the composition arithmetic changes. **Resolving it needs one clean
`set_dress` measurement, not a campaign** — the `D27`/`A3` move applied to a phase.

**The attribution gap that blocks aiming L5 either way.** Inside `set_dress`, the span
`370.98 → 623.71` — **252.73 s, 89 % of the phase** — passes with **no intermediate log line**: it runs
from `[directus] bootstrapping the directus_* system schema` straight to `replayed "taxonomy" … 330261
row(s) loaded`. Two distinct operations (a node CLI bootstrap; a bulk COPY + reindex) are inside it and
nothing distinguishes them. **This is `D17`'s shape one level down** — the phase table attributes
`set_dress` correctly against its siblings, and then hides a 4-minute span behind one anchor. L5 cannot
be aimed until that span is split. Routed as
`FIX-M258-iter08-set-dress-has-no-internal-attribution`.

## D37 — the re-scope trigger is graded and does NOT fire

The milestone's `re_scope_trigger` reads: *"If the composed p50 exceeds 600 s after 3 tiks, split the
suite into a fast smoke lane gating the bring-up + a full lane run after, and renegotiate the gate with
the user."* The measured composed p50 is **840.01 s** and this is the 4th tik, so the trigger's literal
text is satisfied. **It is still graded `n`, on two independent grounds — the second is the load-bearing
one.**

1. **The reps are instrument-rejected.** All three carry `headroom=FAIL`; buildbench's own words are
   *"not usable measurements."* Firing a re-scope on a number the measuring device disowns is the
   category error iter-04, iter-06 and iter-07 each refused in turn.
2. **The remedy does not fit the diagnosis.** The trigger's prescribed action is to *split the suite*.
   But the suite is **not** where the time is: `batch_gate` p50 is **179.37 s [153.81–203.95]** —
   inside the 200 s M256 was budgeted, **while contended**. The dominant phase is `set_dress` at
   283.53 s, which the smoke/full split does not touch at all. Splitting the suite would cost the
   milestone its batch-gate semantics (`D-v28-3`: *one* consolidated red set) and buy nothing.

**And the projection the trigger would override still holds.** Scaling this campaign's own best
observations back to the clean single-box figures already in the record:

```
247.79 s   iter-05/06 gateable single-box bring-up (n=2, agreeing within 1.34 s)
+ 153.81 s iter-08 best-rep batch_gate (batch 142 + restore 11), measured CONTENDED
= 401.60 s vs the 480 s ceiling
```

**401.60 s, and the batch term in it was measured under adverse load** — i.e. the estimate is
conservative on its one contended input. Nothing measured this iter refutes the 414.15 s projection;
the campaign is consistent with it.

**What is honestly still unknown** is whether `set_dress` returns to ~81 s on a quiet box (hypothesis 1
above) or stays near 283 s (hypothesis 2). If the latter, the clean composed total is ~600 s, the
ceiling is unreachable without L5, and **that** is the renegotiation to bring the user — with the
`set_dress` measurement attached. Raising it now, on contended reps and an unsplit 252.73 s span, would
be escalating a guess.

## D38 — D36 SETTLED by scoped diff: the `set_dress` excess is ENVIRONMENTAL, not a regression

`D36` left two live hypotheses for `set_dress` 81.23 → 283.53 s. Rather than spend a campaign, both were
attacked with cheap checks — the `D27`/`A3` move.

**Check 1 — is it a cold-rep-1 artifact?** No. All three reps run **identical** work:

| rep | bootstrap starts | taxonomy replayed | structure provisioned | span |
|---|---|---|---|---|
| 1 | t=370.98 | t=623.71 | `b4cb55bcee08 → ea2e187a1605` | **252.73 s** |
| 2 | t=374.94 | t=636.90 | `b4cb55bcee08 → ea2e187a1605` | **261.96 s** |
| 3 | t=350.53 | t=632.09 | `b4cb55bcee08 → ea2e187a1605` | **281.56 s** |

Same 330,261 rows, same two pgvector reindexes, **same schema digest transition every time** — `--purge`
wipes the directus schema each rep, so the structure auto-provision is a **constant of this
configuration**, not a first-rep cold cost. A warm-vs-cold split cannot explain a phase that costs the
same on all three.

**Check 2 — did the TOOLING change between the two campaigns?** **No — and this is decisive.** iter-05's
authoring HEAD sits between `fast-build-m258-iter-03` and `fast-build-m258-iter-06`; the consumption
clone is pinned at the latter. Across that entire span **11 files changed**, and they are:

```
demo-stack/up-injected.sh (+34)          playthroughs/e2e/batch-gate.sh (+279)
playthroughs/e2e/check-cockpit-roster.py playthroughs/e2e/restore-presenter-world.sh
playthroughs/e2e/stack-paths.sh          playthroughs/manifest/batch_gate_test.go
playthroughs/.gitignore                  stack-core/.gitignore
stack-core/buildbench.py (+36/-7)        stack-core/tests/test_buildbench.py
stack-core/tests/test_m255_mutation_battery.py
```

`git diff --name-only fast-build-m258-iter-03..fast-build-m258-iter-06` filtered for
`setdress|snapshot|stacksnap` returns **zero files**. The entire diff is the **batch-gate feature** plus
the **buildbench instrument**; the 34 lines in `up-injected.sh` are the tail hook at `:2844`, which runs
*after* `autoverify`, itself *after* `set_dress`. **The set-dress code path that measured 81.23 s and
the one that measured 283.53 s are byte-identical.**

**Conclusion: hypothesis 2 is refuted; hypothesis 1 stands.** `set_dress` is the most CPU-exposed phase
in the cycle — a `node cli.js bootstrap` plus a 330 K-row COPY ending in two pgvector index builds — and
it inflated **3.49×** under contention where the cohort inflated **2.05×**. That is a property of the
phase's workload, not a change in what it does.

**What this buys the milestone.** The composed clean projection is *not* threatened:

```
247.79 s  gateable single-box bring-up (iter-05/06, set_dress ≈ 81 s)
+153.81 s best-rep batch_gate measured CONTENDED this iter
= 401.60 s   vs the 480 s ceiling
```

`LEVER-M257-L5-setdress` therefore remains a **reserve for a ~81 s phase**, not a newly-urgent ~283 s
one, and the v2.8 plan's ~30–50 s pricing of it is not invalidated. **The gate needs a quiet host, not a
new lever** — which is a materially better position than `D36` left the milestone in an hour ago.

**Scope stated honestly:** this compares two *tags*. If iter-05's authoring clone carried uncommitted
set-dress edits at measurement time, the argument does not cover them — nothing in its `progress.md`
suggests it did, and `rext_dirty` on this campaign's own clone resolved to a single untracked
`.agentspace/rext.tag` marker, but it is a tag-to-tag comparison and is recorded as one.

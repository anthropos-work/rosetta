**Type:** tik · **Active strategy:** `TOK-02` step 3 — *take the baseline on the contended box, and label it*

## Line 1 — the first campaign refused at the pin guard, and it paid for itself anyway

The first `buildbench run 1 --reps 3 --profile macmini` launched at **20:44:26Z** and both reps it reached
died **37.79 s** and **25.52 s** in, at `up-injected.sh`'s M217 pin guard:

```
==> ✗ FATAL: rext pin mismatch.
    the consumption clone is at : main
    .agentspace/rext.tag pins   : fast-build-m257x-iter-288
    This stack would run TOOLING THAT IS NOT THE ONE YOU THINK IT IS.
    Or, to run un-pinned authoring work deliberately: DEMO_ALLOW_UNPINNED_REXT=1
```

**Correct refusal, and the hatch it names is the right answer here.** `buildbench`'s `rext_root()` is the
tree it lives in, so a campaign driven from the **authoring copy** necessarily runs an un-tagged `main` —
which is exactly the state the guard exists to catch, and exactly the state this measurement *wants*: the
tooling under test is HEAD (`64bbb81`, iter-06's units fix), deliberately. The alternative — cut a tag and
`git push --tags` — is a **publish** action taken purely to satisfy a guard, on a session whose brief says
nothing is pushed. So the run was relaunched at **20:46:30Z** with `DEMO_ALLOW_UNPINNED_REXT=1`, and the
deviation is **labelled** in the campaign label itself (`m257-iter07-contended-unpinned`) rather than left
for a reader to infer. Same discipline as the contention label: state the deviation with the number.

Evidence from the refused run is kept at `.agentspace/scratch/work-m257/baseline-n3-pinrefused/`.

### And it produced the fix's first live confirmation, for free

Both refused reps still ran the sampler and the post-rep headroom assert. Their verdicts:

| rep | wall | `rc_up` | **peak load1** | headroom under the **corrected** basis (limit 10) | under the **old** basis (limit 6) |
|---|---|---|---|---|---|
| 01 | 37.79 s | 1 | **6.51** | **OK** | **would have FAILED** |
| 02 | 25.52 s | 1 | **7.41** | **OK** | **would have FAILED** |

This is what a unit test cannot give: **two real campaign reps, on the real gate host, that the pre-iter-06
clause would have refused as "CPU-saturated" while the host demonstrably had headroom.** The defect was
never hypothetical — it was one campaign away from mislabelling every rep of this milestone's baseline as
contended, and a baseline whose every rep is stamped *"timings are contended and not comparable"* is not a
baseline.

*(Note what did **not** happen: neither rep was refused for load. The pin guard stopped them, and the
headroom clause — the thing this milestone spent iter-06 fixing — passed both. Two independent gates, two
different verdicts, and only one of them was the reason for the failure.)*

## Line 2 — interrupting a campaign costs two cleanup steps, and neither is obvious

Relaunching after the pin refusal hit a second wall immediately:

```
hostlock: (holder pid 63877 is not running on Marcos-Mac-mini.local; age=27s, ttl=7200s).
up-injected: REFUSING to bring up demo-1 — the demo host is already held by another cycle
```

Two mechanisms compounded, and the interesting part is that **each is individually correct**:

1. **`run_campaign` orphans its child.** It drives the cycle with
   `subprocess.Popen([...up-injected.sh...])` and no process group, so `pkill`-ing the harness leaves
   `up-injected.sh` running with `ppid=1`. Measured: after killing the python parent, pids 64500/64508 were
   still alive 17 s later, mid-bring-up.
2. **The hostlock refuses a lock whose holder it can already prove is dead.** `hostlock.sh:138` requires
   **both** `gone=1` **and** `age > HOSTLOCK_TTL`, and the TTL is **7200 s** — generous on purpose, because
   *"a full cycle is ~15–45 min."* So a killed campaign blocks its own slot for **up to two hours**, while
   printing that the holder is not running.

Cleared deliberately and in order, each step verified rather than assumed: killed the orphaned
`up-injected.sh` (then confirmed **none** remained), confirmed `demo-1` had **no containers** left behind,
confirmed the lock holder was dead (`kill -0 63877` → dead), and only then removed the lockfile — which is
one of the two remedies the guard's own message offers (`HOSTLOCK_TTL=0`, or remove the file).

**The finding is the *pairing*, not either half.** The TTL is right for a live cycle, and the guard is
right to be conservative about a lock it did not take. But an operator who interrupts a campaign — which
this milestone will do repeatedly while pricing levers — is left with a *silently orphaned bring-up* that
can still be mutating the stack, plus a lock that reads as held for two hours. Routed forward as
`FIX-M257-campaign-kill-orphans-bringup`.

## Line 3 — side finding: one of the gate's two falsifiable asserts has no code yet

While confirming what the campaign's post-rep gate actually asserts: **HEADROOM is implemented
(`headroom_assert`); ISOLATION is not implemented anywhere in `rosetta-extensions`.** Searched the whole
tree — the only `isolation`-named artefact is `demo-stack/tests/test_host_isolation.py`, which fences the
**hostlock** (concurrent cycles on one host, `GUARD-M221`) and is an unrelated meaning of the word.

**This is not a defect, and saying why matters.** `TOK-01` is explicit that each falsifiable assert lands
*"together with the lever that can trip it — never after"*, and the gate itself scopes ISOLATION to the
layers **L1/L3** rewrite. No lever has landed, so there is nothing yet for it to inspect.

**What is worth writing down is that nobody had.** A gate clause with no instrument is precisely the class
v2.8 exists to retract ("sampled, not asserted"), and this milestone's own opening lesson is that a gate
should be checked for *gradeability* before it is checked for *satisfaction*. Recorded here so that when
L1 lands, ISOLATION lands with it rather than being discovered missing at the gate.
Route: `ASSERT-M257-isolation-with-L1`.

## Line 4 — the instrument that would make a clause-1 refusal *attributable* does not run on this host

Reading the live sampler mid-rep, two of its nine columns are **empty on every row**:

```
elapsed_s  load1  mem_used_mb  swap_used_mb  disk_avail_gb  containers  disk_util_pct  disk_write_mbs  top_cpu_proc
      232  14.31       3849.3            ""             ""          ""             ""              ""  com.apple.Virtua:232.5
      282   8.53       4044.7            ""             ""          ""             ""              ""  a8-cart-runner:99.0
```

`disk_util_pct` / `disk_write_mbs` come from `disk_busy_snapshot()`, which reads `/proc/diskstats` and says
so in its own docstring: **"Linux only (macOS has no /proc); None elsewhere, and the summary simply omits
the columns."** Nothing is hidden and nothing is broken. What is *new* is the consequence of `D-v28-15`,
and it appears to be unrecorded anywhere:

> **The gate host is now a Mac, and spike (d)'s discriminator is Linux-only — so on the gate host, a
> clause-1 refusal cannot be attributed.**

That matters because of what clause 1 *says* when it fires: *"the run was CPU-bound, so its phase timings
are contended and not comparable."* The instrument built specifically to test that inference —
`disk_busy_snapshot`, whose docstring states the case in as many words (*"a run blocked on disk shows up in
load1 without being CPU-bound at all … without a %util series you cannot tell a CPU plateau from an I/O
ceiling"*) — **produces nothing here**. The clause still refuses; it simply cannot say why.

This is the live, reproducible form of `INVESTIGATE-M257-load1-48`, which has been carried unstarted since
iter-02 and whose original subject (odysseus) no longer exists. Re-aimed here as
`FIX-M257-io-sampler-macos`.

**One thing deliberately NOT claimed.** The Linux caveat is about Linux's load-average definition
(uninterruptible-sleep tasks count). Whether Darwin's `getloadavg()` shares that property is a *separate*
question this iter did not measure, and asserting it either way would be exactly the unmeasured
host-generalisation this milestone keeps retracting. What is measured is narrower and sufficient: **the
%util series is absent on this host**, so nothing here can distinguish the two cases.

## Line 5 — protocol compliance: the gated variant is *cold images, WARM layer cache*, and this host qualified

`build-budget.md` § The campaign protocol carries a **step zero**: on a host whose BuildKit cache is
**empty**, one full warm-up cycle must run and be **discarded**, because rep 1 there would measure the
*truly-cold* variant `D-v28-8` deliberately cut from the gate — *"a different measurement class, not an
unlucky rep"*. Checked before launching rather than after:

| check, 20:40Z (pre-campaign) | value |
|---|---|
| `docker system df` — Build Cache | **129 records, 31.17 GB** (23.28 GB reclaimable) |
| Images | 39 (14 active), 39.32 GB |
| pnpm store, observed live in rep-01's hiring build | `resolved 2411, reused 2390, **downloaded 0**` |

The cache is **populated**, and the live `reused 2390 / downloaded 0` line is the confirmation that matters:
it is the cache being *used*, not merely present. So no discard-warm-up was owed, and rep-01 counts.

*(Worth stating because the opposite mistake is cheap and invisible: `odysseus` was probed with 0 images,
0 containers and 0 build cache, and a campaign started there would have silently graded the excluded
variant. This host is the other case, and it was verified rather than assumed.)*

## Line 6 — THE FINDING: a campaign driven from the authoring copy is **not the gated configuration**

Mid-rep-01, `demopatch.log` reads **17 REFUSED, 0 applied** — every patch in the set:

```
next-web-studio-url:            REFUSED — the Studio nav link will EJECT TO PROD
next-web-ssr-graphql-origin:    REFUSED — authenticated renders will block ~37 s on an unreachable SSR GraphQL origin
next-web-members-pagination:    REFUSED — the members fetch stays limit:1000
next-hiring-members-pagination: REFUSED — the dashboard hangs and the scoreboards never mount
… 13 more, including both interview-flag gates, both back-to-cockpit items and studio-desk-shell-first-paint
```

**Root cause, read from `demo-stack/patches/demopatch` and confirmed by probe — both of G6's arms are False,
for two independent reasons:**

| G6 arm | mechanism | why it is False here |
|---|---|---|
| `_is_demo_workspace` | requires `workspace_root` to be **realpath-equal** to *this demopatch binary's own clone-set workspace* (`_HERE/../../..`) | buildbench's `rext_root()` is the tree it lives in, so the binary's own workspace is **`…/rosetta/.agentspace`**, while `up-injected.sh` passes **`…/rosetta/stack-demo`** (where the platform clones are). Not equal → False |
| `_registry_has_demo` | any registry row with `type == "demo"` | `stack_registry.list_stacks(reconcile=False)` returns **`[]`** on this box (probed directly) → False |

**Why this invalidates the campaign as a baseline, in three independent ways:**

1. **The gate requires *"all 7 demopatch guards (G1–G7) passing."*** Here G6 is refusing every patch. That
   clause cannot be satisfied by this configuration at all.
2. **It changes the timing, in the direction that matters.** `next-web-ssr-graphql-origin` REFUSED means
   *"authenticated renders will block ~37 s on an unreachable SSR GraphQL origin"* — the exact
   *blackholing* signature `latency-budget.md` teaches you to recognise arithmetically. A cycle whose verify
   phase blocks on that is not measuring the same work.
3. **It is a different artifact.** Ejects to prod, an un-paginated members grid, both interview gates
   closed — this is not the demo the gate describes.

**So the number this campaign produces is NOT `BASELINE-M257-macmini-n3`, and must not be written into
`macmini.json`'s `gated_baseline`.** Recording it as one would be precisely the number-shaped defect v2.8
exists to retract — a figure that looks like a baseline, carries a host name, and measures a configuration
nobody will ever ship.

**What a valid baseline requires** is the documented rung-zero workflow: `rosetta-extensions` **tagged and
`git push --tags`ed to origin**, `stack-demo/rosetta-extensions` re-pinned to that tag, and `buildbench`
driven **from that consumption clone** — where `own_ws` and `workspace_root` are the same directory and G6
arm 1 holds. That is a **publish** action, and this session's brief states plainly that nothing is pushed.
It is not mine to take unasked.

*(Note the shape of this: the corpus's rung-zero rule — **"tagging is not publishing"** — is written about a
**remote** stack being unable to obtain a tag. This is the same rule biting **locally**, on the box that
authors the tooling, and for a different reason: not that the tag is unreachable, but that the authoring
copy is structurally the wrong workspace to run a demo from. The guard is right; the harness's
`rext_root()` assumption is what does not hold.)*

## Line 7 — rep-01: a complete cycle, and what it does and does not tell us

The bring-up **succeeded** — `rc {teardown: 0, bringup: 0}`, all 11 `demo-1` containers up, migrations,
snapshot replay and seed all run. So this is a real, finished cold cycle, not a crash.

```
total_s            681.71
rc                 {'teardown': 0, 'bringup': 0}
verdict            {'green': False, 'warnings': 2, 'ts': '2026-08-11T20:59:21Z'}
headroom           ok=False  failures=['peak_load1']
observed           peak_load1 18.71 · load1_core_basis 12.0 · profile_cores 8.0
                   heap_commitment 5388 MiB · disk_avail 68.7 GiB · max_parallel_ui_lanes 2
samples (n=68)     peak_load1 18.71 · avg_load1 6.59 · peak_mem 4642.7 MB · peak_swap 1336.1 MB
```

### Three independent reasons this rep does not satisfy the gate

1. **`autoverify green: False`, 2 warnings** — the gate demands `green:true / 0 warnings`.
2. **HEADROOM clause 1 FAILED**: peak load1 **18.71** > **10** (`cores-2` on the 12-core basis).
3. **17/17 demopatches refused** (Line 6) — the gate demands all 7 guards passing.

**And note what the ledger now carries because of iter-06:** `load1_core_basis: 12.0` **beside**
`profile_cores: 8.0`. The refusal is *attributable to a machine* in the permanent record, rather than being
a bare number a later reader would have to re-derive. Under the old code this rep would have failed the
same clause against a limit of **6**, and nothing in the artefact would have said which machine that 6
described.

### The number that matters, and it is a **bound**, not a baseline

**681.71 s for a cycle with every demopatch refused** — i.e. a cycle doing *less* work than the gated one.
Set against the milestone's own arithmetic:

| | |
|---|---|
| `billion` n=3 p50 (the release's starting number) | **666.29 s** |
| iter-04's **estimate** for this host, scaled from billion's phase table | ~420–455 s |
| **this host, measured, degraded configuration, contended** | **681.71 s** |

**iter-04's estimate looks optimistic by roughly 50 %, and this is the first evidence either way.** It was
honestly labelled an estimate — a scaling of billion's phases by one measured image — and it is now doing
what estimates do. The consequence for the milestone is concrete: iter-05's re-cut rationale said *"≤ 360 s
looks reachable, plausibly on L1 alone."* Against a **681 s lower bound**, 360 s is a **~47 % cut** — the
same order as billion's 46 %, not the gentle one the estimate implied. **L1 alone will not do it here.**

**This does not reopen the gate re-cut.** The target was never conditional on the estimate; iter-05 kept
360 s *unchanged* precisely because it is the release thesis. What changes is the **expected lever
programme**, and `re_scope_trigger` is the mechanism that reads it — which is why it must be re-derived
against a *real* `gated_baseline`, exactly as its re-cut text now demands.

**One more headroom signal, recorded not acted on:** peak swap **1336 MB** against a 3072 MiB VM swap
budget. The VM swapped during the build with `MemAvailable` at 9.87 GiB pre-run. Worth watching when L2
(two concurrent lanes, which this host's profile derives as *possible*) is priced — a lane count that fits
on paper and swaps in practice is the M255 lesson in a different currency.

## Line 8 — the same units family again, one artefact apart: the sampler's disk column is the HOST's

rep-01's own ledger carries both numbers, 2.7× apart, on the axis the M239-F1 ENOSPC trap lives on:

| field | value | source |
|---|---|---|
| `headroom.observed.disk_avail_gib` | **68.7 GiB** | `effective_disk_avail_gib()` — VM first, host as fallback (`:606`) |
| `samples.min_disk_gib` | **186.9 GiB** | `host_disk_avail_gib()` — host `/` (`Sampler.run`, `:781`) |

**The assert is right; the human-readable series is not.** Clause 3 consumes the VM number, so the guard is
sound. What is wrong is the **time series an operator reads to watch disk trend toward the floor** — it
reports the *host's* free space, over-stating it by 2.7× on this host class, in exactly the direction that
produced M239-F1 (*"a mid-campaign ENOSPC presents as the cryptic `redis exited (1)`, not as a disk
error"*).

Same family as iter-06's clause-1 defect and the same shape: **one file, two probes, and the host/VM
distinction honoured in one place and dropped in another.**

**Deliberately NOT fixed here** — the scope tripwire, and for a substantive reason rather than tidiness:
`effective_disk_avail_gib()` shells out to `docker run --rm busybox df`, so calling it on a 10-second
sampler cadence would both cost real time and *perturb the measurement it is sampling*. The fix needs a
cached or cheaper VM probe, which is its own piece of work. Routed as `FIX-M257-sampler-disk-units-vm`.

## Line 9 — THE UNIFICATION: three symptoms, one root cause, and the stack itself was fine

rep-01's `autoverify` reports **2 failed checks**, and reading them settles what this campaign actually
measured. The healthy asserts first — this stack **worked**:

```
✓ backend /api/health 200 on :18082
✓ sentinel.casbin_rules = 1251 (authz policy loaded)
✓ directus.directus_collections = 21 (the content model is REGISTERED)
✓ directus DB is per-stack-local (not prod)
✓ frontend builds: ok (the running images are this run's)
```

And the two failures:

```
✗ postgres-schemas  fail: cannot derive the expected schema set
                    (missing …/.agentspace/rosetta-extensions/stack-verify/lib/../../../platform/repos.yml)
                    — refusing to assert a hand-maintained list
⚠ a demo-patch was NOT APPLIED — refused or skipped   [the 17 from Line 6]
```

**Both resolve to the same sentence.** `…/stack-verify/lib/../../../platform/repos.yml` is
**`.agentspace/platform/repos.yml`** — and the platform clone lives in **`stack-demo/platform`**. So:

| symptom | mechanism | one root cause |
|---|---|---|
| 17/17 demopatches REFUSED | G6 arm 1 compares the binary's own workspace (`.agentspace`) to the passed `workspace_root` (`stack-demo`) | ⎫ |
| `postgres-schemas` probe FAILED | resolves `repos.yml` relative to the rext tree → `.agentspace/platform/`, which does not exist | ⎬ **`buildbench.rext_root()` is the tree buildbench lives in, and every workspace-relative lookup inherits it. Run from the authoring copy, they all resolve one directory tree away from the stack.** |
| `autoverify green: False` | is *caused by* the two above | ⎭ |

**So the gate's own verdict clause was RED for a reason that has nothing to do with the demo.** Three
symptoms, one cause, and every one of them is an artefact of *where the harness was invoked from* — not a
defect in the stack, the platform, or the tooling under test. The `postgres-schemas` probe in particular
did exactly the right thing: it **refused to assert a hand-maintained list** rather than passing on a guess.

**This is the iter's real deliverable**, and it is worth more than the number it displaced:
`BASELINE-M257-macmini-n3` is not blocked by contention, by disk, by memory, or by the host — it is blocked
by **where `buildbench` can legitimately be run from**. That was not knowable from reading the code; it took
a real cold cycle to surface, and it would have silently produced a wrong "baseline" for anyone who trusted
the total_s and skipped the verdict.

## Line 10 — the fix for the baseline, VERIFIED rather than asserted

The obvious remedy — *"run `buildbench` from the pinned `stack-demo` consumption clone instead"* — was
checked against the actual predicates rather than assumed to work:

| requirement | check | result |
|---|---|---|
| G6 arm 1: `own_ws == workspace_root` | `realpath(stack-demo/rosetta-extensions/demo-stack/patches/../../..)` vs the `workspace_root` `up-injected.sh` passes | both **`…/rosetta/stack-demo`** ✅ |
| G6 arm 1: the workspace is demo-shaped | `stack-demo/rosetta-extensions/demo-stack/stacks/` | **exists** ✅ |
| `postgres-schemas` probe | `stack-verify/lib/../../../platform/repos.yml` → `stack-demo/platform/repos.yml` | **exists** ✅ |
| `buildbench --profile macmini` | `stack-demo/rosetta-extensions/stack-core/hostprofiles/` | **billion.json, laptop.json only — `macmini.json` ABSENT** ❌ |

**Three of four hold today. The fourth is the whole blocker, and it is the rung-zero one.** That clone is
pinned at `fast-build-m257x-iter-279`, which predates both `macmini.json` (iter-04) and the clause-1 units
fix (iter-06). `load_host_profile("macmini")` there is **exit 2** by design — *"a missing profile is exit-2
territory, never a pass."*

**So a valid `BASELINE-M257-macmini-n3` needs, in order:**

1. `rosetta-extensions` **tagged** at a commit containing `macmini.json` + `load1_core_basis`;
2. **`git push --tags` to origin** — the tag must be *fetchable*, per the M217 FATAL pin guard and the
   corpus's own rung-zero rule (*"tagging is not publishing"*);
3. `stack-demo/rosetta-extensions` **re-pinned** to that tag (`.agentspace/rext.tag` updated);
4. the campaign re-run **from that clone**, after tearing down the authoring-copy-driven `demo-1` (same
   compose project name, so the two cannot coexist).

**Step 2 is a publish action and this iter does not take it.** The session's brief states that nothing is
pushed, and publishing tooling to a shared origin — tooling whose *full* guard sweep was deliberately
killed mid-run (D2) to protect this very measurement — is a decision with blast radius past this milestone.
Surfaced for the user rather than taken quietly.

*(The irony is worth one line, because it is the milestone's own recurring lesson: the corpus's rung-zero
rule was written for a **remote** stack that cannot fetch an unpushed tag. It has now bitten on the
**authoring box itself** — and not because the tag is unreachable, but because the authoring copy is
structurally the wrong workspace to run a demo from. Same rule, a place nobody had looked.)*

## Line 11 — rep-02 CORRECTS Line 7's bound, and vindicates the `n ≥ 3` rule

| rep | total | `rc_up` | green | headroom | peak load1 |
|---|---|---|---|---|---|
| 01 | **681.71 s** | 0 | False | **FAIL** (18.71 > 10) | 18.71 |
| 02 | **472.82 s** | 0 | False | **OK** | 6.25 |

**A 209 s spread — 44 % — between two consecutive reps of the same command on the same box.**

**Line 7's reading was wrong and is corrected here.** I wrote that *"iter-04's estimate looks optimistic by
roughly 50 %"* on the strength of rep-01 alone. rep-02 lands at **472.82 s**, close to iter-04's
**~420–455 s** estimate — which is now the better-supported figure, not the refuted one. **The n=1 reading
inverted the conclusion.**

**This is exactly the failure `build-budget.md` legislates against**, and it caught me in the same session I
quoted the rule in: *"the baseline is a `p50` over `n ≥ 3` and never a mean"*, because a single rep can
carry a one-off cache eviction (M255 measured one at **+173 s**; this spread is 209 s). I verified the
build cache was *populated* before launching (Line 5) and concluded rep-01 would be a steady-state rep.
**Populated is not the same as populated *for this stack's layers*** — `demo-1`'s images bake
offset-specific build args (offset URLs + a minted publishable key), so its layers are not `demo-2`'s, and
rep-01 paid for cache the earlier stack's builds had never created. **rep-01 was a de-facto warm-up.**

So the protocol's step-zero rule needs a sharper form than *"is the cache empty?"*, and this is the
generalisation worth carrying:

> **Cache warmth is per-STACK, not per-HOST.** A box with 31 GB of BuildKit records can still be cold for
> the stack you are about to measure, because per-stack build args cut new layer chains. Grade the rule on
> *"has THIS demo-N been built here before"*, not on `docker system df`.

**And rep-02 is a second live confirmation of iter-06's fix**: peak load1 **6.25** passes against the
corrected limit of 10 and would have **FAILED** the old limit of 6. Two of the three reps so far sit in the
band between the wrong limit and the right one — this defect was not an edge case on this host, it was the
common case.

## Line 12 — the campaign, complete: n=3, and what it is allowed to be used for

```
rep-01 total=681.71  rc_up=0  green=False  warn=2  headroom=FAIL  peak_load1=18.71
rep-02 total=472.82  rc_up=0  green=False  warn=1  headroom=OK    peak_load1=6.25
rep-03 total=489.90  rc_up=0  green=False  warn=1  headroom=FAIL  peak_load1=16.88

n=3   min 472.82   p50 489.90   max 681.71
```

**All three bring-ups SUCCEEDED (`rc_up=0`).** Every failure above is a *verdict*, not a crash.

### What this is NOT

**Not `BASELINE-M257-macmini-n3`, and it must not be written into `macmini.json`'s `gated_baseline`.**
Three independent disqualifications, each already evidenced above:

1. **`green: False` on all three** — the gate demands `green:true / 0 warnings`.
2. **17/17 demopatches refused on all three** (Lines 6, 9) — the gate demands all 7 guards passing, and the
   cycle is doing *less work* than the gated one.
3. **HEADROOM refused 2 of 3** — peak load1 **18.71** and **16.88** against the corrected limit of 10.

### What it IS — a defensible order-of-magnitude anchor, and it changes the outlook

| | seconds | vs the 360 s gate |
|---|---|---|
| `billion` n=3 p50 (the release's starting number) | 666.29 | a **46 %** cut |
| iter-04's estimate for this host (scaled, one image) | ~420–455 | ~14–21 % |
| **this host, n=3 p50, contended + degraded + unpinned** | **489.90** | a **~27 %** cut |

**iter-04's estimate holds up far better than rep-01 alone suggested** — optimistic by roughly 8–16 %, not
the ~50 % Line 7 inferred before rep-02 landed. And the milestone's outlook is the one iter-05 argued for:
**this host needs a ~27 % cut where billion needed 46 %.** Note the direction of the residual bias — the
gated cycle applies 17 patches this one skipped, so the true figure sits *above* 489.90; how far above is
not knowable from here, which is precisely why this is an anchor and not a baseline.

### The contended-box result, reported as a result

**2 of 3 reps were refused by HEADROOM clause 1** — the honest outcome of measuring on a workstation that
cannot be freed. What the run *would* have measured is recorded above rather than discarded: 681.71 s at
peak load1 18.71, and 489.90 s at 16.88. This is the `laptop.json` situation handled the other way — that
profile records a refusal at load1 10.69 and **no cycle number at all**; this one records both, and says
which is which.

**And the units fix earned its place three times in one campaign.** Under the pre-iter-06 basis (limit 6
instead of 10), rep-02's 6.25 would ALSO have been refused — making it **3 of 3 refused**, with no rep
surviving to show that a clean cycle on this box is even possible. The defect would not have skewed this
campaign; it would have erased its only passing rep.

## Close — 2026-08-11

**Outcome:** The `n ≥ 3` campaign RAN — 3 complete cold cycles, all `rc_up=0`, **p50 489.90 s** (min
472.82 / max 681.71) — and **it is not a baseline, for a reason worth more than the number**: a campaign
driven from the **authoring copy** resolves every workspace-relative path one directory tree away from the
stack, so **17/17 demopatches are refused** and the `postgres-schemas` probe cannot find `repos.yml`. That
single root cause produces all three of the run's disqualifications (`green:False`, patches refused,
`autoverify` RED). The remedy is verified against the actual predicates, and its one missing precondition
is a **publish**: `rosetta-extensions` tagged and pushed to origin, `stack-demo` re-pinned, campaign re-run
from that clone. `gated_baseline` is deliberately left **empty**.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(tik; the streak reset at iter-05)* — (3) re-scope: n *(the trigger reads a p50 after L1+L2+L3; no lever has landed)* — (4) **user-blocker: y** *(a valid baseline requires `git push --tags` to a shared origin — a publish action this session's brief excludes, on tooling whose full guard sweep was deliberately killed to protect the measurement. Separately, `corpus/ops/demo/build-budget.md` is correct-but-uncommittable behind two PRE-EXISTING fence REDs in a file this iter never touched)* — (5) cap-reached: n *(tik 2 of 5)* — (6) protocol-stop: n — (7) budget-exhausted: n — **Outcome: exit-4**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D3)
**Side-deliverables:**
- `corpus/ops/demo/build-budget.md` — the protocol doc's SUPERSEDED banner carried the **retracted
  `overlay2` claim** and two items fixed since it was written; plus two line citations that **iter-06's own
  code change broke**. Corrected. **UNCOMMITTED** — see D3.
- `stack-core/hostprofiles/billion.json` — its `gated_baseline` note still named the retired `odysseus`
  (committed, `13e9638`).
**Routes carried forward** (Fate 3, named handlers):
- **`BASELINE-M257-macmini-n3`** → still owed. **Unblocked on evidence, blocked on a publish.** Steps 1–4
  in Line 10; three of four preconditions already hold.
- `FIX-M257-campaign-kill-orphans-bringup` → `run_campaign`'s `Popen` has no process group, so killing the
  harness orphans `up-injected.sh`; combined with a 7200 s hostlock TTL, an interrupted campaign blocks its
  own slot for up to two hours while printing that the holder is dead (Line 2).
- `FIX-M257-io-sampler-macos` → `disk_busy_snapshot()` is `/proc/diskstats`-only, so on the (Mac) gate host
  the %util series is absent and **a clause-1 refusal cannot be attributed** (Line 4). Re-aims
  `INVESTIGATE-M257-load1-48`, whose original host no longer exists.
- `FIX-M257-sampler-disk-units-vm` → the sampler's disk column is the **host's**, the assert's is the
  **VM's** — 186.9 vs 68.7 GiB in one ledger (Line 8). Needs a cached VM probe, not a one-liner.
- `FIX-M257-anchor-guard-resolution` → the anchor guard resolves two anchors to content **neither of its own
  declared clone roots contains** (D3).
- `ASSERT-M257-isolation-with-L1` → the gate's ISOLATION clause has **no implementation** yet; correctly
  deferred by `TOK-01` to land *with* L1, but nobody had written that down (Line 3).
- iter-05/06's remaining routes carry unchanged (`MEASURE-M257-macmini-true-idle`,
  `PROFILE-M257-provisional-fields`, and iter-03/04's tail).
**Lessons:**
- **An n=1 measurement can invert the conclusion an n=3 supports.** rep-01 alone said iter-04's estimate was
  ~50 % optimistic; reps 1–3 say ~8–16 %. I quoted the *"p50 over n ≥ 3, never a mean"* rule in this same
  iter (Line 5) and still drew a conclusion from one rep (Line 7) before the others landed. **The rule is
  not about arithmetic robustness — it is about not publishing the first number you see.**
- **Cache warmth is per-STACK, not per-HOST.** 31 GB of BuildKit records did not make `demo-1` warm, because
  per-stack build args (offset URLs, minted keys) cut new layer chains. rep-01 was a de-facto warm-up and
  the protocol's step-zero check — *"is the cache empty?"* — cannot see that.
- **A harness that assumes it lives in the workspace it measures cannot measure from anywhere else.**
  `rext_root()` is `Path(__file__).parents[1]`, and every workspace-relative lookup downstream inherits it.
  Run from the authoring copy, demopatch's G6, the `postgres-schemas` probe and the `autoverify` verdict all
  fail for that one reason — and the *timings still look plausible*, which is what makes it dangerous.
- **The corpus's rung-zero rule bit locally.** *"Tagging is not publishing"* was written about a remote stack
  that cannot fetch an unpushed tag. It has now blocked a measurement on the box that authors the tooling —
  same rule, a place nobody had looked.

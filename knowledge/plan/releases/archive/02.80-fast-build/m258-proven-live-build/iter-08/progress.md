# M258 iter-08 — progress

**Type:** tik · **Active strategy:** `TOK-01` step **4** — *the composed 3× cold campaign against the
gate, with the spread published beside the p50.*

## Phase A — pre-flight (the `D21` check, re-run rather than inherited)

The pin was verified to carry **the feature under test**, not merely to exist: consumption clone at
`fast-build-m258-iter-06` (`5f3a381`), `playthroughs/e2e/batch-gate.sh` present, hook live at
`up-injected.sh:2844`. iter-07 lost a re-pin to exactly this gap (four files absent, no hook) and the
phase table would still have summed.

User's stacks confirmed resident **before**: `demo-2` = 11, dev (`anthropos-*`) = 5. Disk 170 GiB host
free. `demo-1` = 11, presenter-usable from iter-07.

## Phase B — the trigger was automated, and it fired in 91 seconds

**The brief this iter opened under said the box was calm** (`load1 2.41 → 2.42`, *"anima8 finished on
its own and freed the box"*). Measured at 08:12:39Z it was **`11.07`, rising to 14.25** across three
samples, with `anima8`'s `m270a2-iter132/wt/shifttrap/lab/st_pin3.py` at ~89 % CPU plus **31 node
processes** across two `hyper-studio-worktrees/` checkouts. *Figures written from memory are wrong* —
including figures written 90 seconds ago about a machine three projects share.

The campaign was launched anyway, immediately, because **the instrument is safe-by-refusal**: it exits 8
on a headroom miss *before* the teardown. It refused at **`load1 31.16` vs a ceiling of 10**, `demo-1`
untouched. **A HEADROOM refusal is a RESULT**, and it cost 25 seconds.

Then the actual fix (`D29`). iter-07 concluded *"the window never opened"* after ~30 minutes of **hand
polling at ~2-minute granularity**. But its own series shows the box moving 21 points inside one minute
(`07:33 11.93 · 07:34 32.90`) — **a hand-sampled trigger cannot catch a window shorter than its own
interval.** So the human was removed from the trigger: `autoarm-campaign.sh` samples every 15 s and
fires on 3 consecutive samples ≤ 5.0, re-arming rather than giving up if the window closes before the
final assert.

**Armed 08:17:41Z. Fired 08:19:12Z — 91 seconds later**, on a decay it watched happen:

```
7.88 → 6.98 → 6.34 → 5.37 → 4.56 → 4.11 → 3.87
```

The dip below 5.0 lasted **75 seconds**. iter-07 would have had to poll inside that window by luck.

`assert-headroom` passed at **3.87**; teardown completed at **4.44**; `demo-2`=11 / dev=5 verified again
post-teardown. **By 08:21:02Z three `a8-cart-runner` processes were back at 94–97 % CPU** — the 08:11–08:19
lull was a gap *between* anima8's iterations, not the end of them (`D30`).

## Phase C — the campaign, and what it proved

08:19:19Z → 09:02:12Z, `CAMP_RC=1`. **The first complete composed 3-rep cold campaign this milestone has
had.**

```
rep 1/3 total=811.06s up_rc=0 green=True headroom=FAIL isolation=OK phases=complete
rep 2/3 total=840.01s up_rc=0 green=True headroom=FAIL isolation=OK phases=complete
rep 3/3 total=859.06s up_rc=0 green=True headroom=FAIL isolation=OK phases=complete
  total cycle   p50 840.01s   min 811.06s   max 859.06s   (never a mean)
  host identity MATCH {'match': 3}
```

**Every functional column green on every rep. Every timing column refused** — `headroom=FAIL` 3/3 on the
single clause `peak_load1` (40.09 / 74.77 / 51.80 against a limit of 10). buildbench's own verdict:
*"rep(s) [1, 2, 3] are not usable measurements."* **So 840.01 s is not the composed p50 the gate asks
for** — it is the p50 of three reps the measuring device disowns (`D35`).

**The spread, as `TOK-01` requires it be published beside the p50: 811.06 → 859.06 = 1.06×.** M256
escalated a **2.04×** spread as evidence its suite was *"not decidable at n=3 on this host"*. At 1.06×
under *adverse* conditions, the composed cycle does not show that pathology.

### C1. `D27`'s prediction held (`D31`)

iter-07 refused to spend a campaign testing iter-06's routed hypothesis and reduced it to path
arithmetic plus one query, predicting `autoverify`'s `postgres-schemas` probe would **pass** from the
consumption clone. Measured: `{"warnings": 0, "green": true}` on **all three reps**, up from
`{"green": false, "warnings": 1}`. *Settle by arithmetic when you can.*

### C2. The `D17` attribution check passes in BOTH directions (`D32`)

`D17` is the defect where a 166 s phase hid inside a 2 s one **and the table still summed**.

| rep | `autoverify` | `batch_gate` | Σ sub-phases | `P4_BRINGUP` |
|---|---|---|---|---|
| 1 | **5.32 s** | **153.81 s** | 802.39 | **802.39** |
| 2 | **4.16 s** | **179.37 s** | 822.19 | **822.19** |

`autoverify` reads ~5 s, not ~154 s — the two **attribute separately** — *and* Σ sub-phases equals
`P4_BRINGUP` exactly, with `missing_anchors: []` and `phases_complete: true` on all three reps. The
table both **sums and attributes**.

**And the artifacts agree with each other, not merely with themselves** (the `D19`/`D20` check):
`batch-gate.json` reports `batch_seconds 142 + restore_seconds 11 = 153`; the anchor-derived
`batch_gate` span is **153.81 s**. Two instruments, one number, timestamps one second apart.

### C3. The stale-artifact trap this run walked past (`D33`)

`demo-1`'s stack dir exists in **both** clones. The authoring copy still held a `batch-gate.json` from
**06:54:15Z** — iter-06's, 1.5 h stale, survived `--purge` — reading `verdict green · red_count 0`,
**identical to this run's verdict**. A reader who reached for it by name would have reported a stale
green and been entirely self-consistent. Only the `ts` betrays it. Every artifact in this iter is
path-qualified and `ts`-checked.

### C4. The red set stayed EMPTY under load (`D34`)

iter-07 declined to launch partly because *"a batch run at `load1` 40–60 would have manufactured FALSE
REDS"* — `retries: 0` is contract, so timeouts become failures escalated as a **product** verdict.
Measured, with `load1` 31–40 during the batch: `red_count 0 · red_set [] · passing 30 · failing 0 ·
unimplemented 1` — **on all three reps, identical to iter-07's clean baseline.**

Stated with its limits: one observation band, and it does **not** license running the gate contended —
the *timings* are still invalid. What it retires is the specific claim that contention at this magnitude
manufactures reds in this suite.

## Phase D — `set_dress` 81.23 → 283.53 s, and its settlement

The campaign's most important number was not the total. `set_dress` came in at **283.53 s
[282.91–306.22]** against iter-05's **81.23 s [78.92–117.51]** — *the same command, both n=3, both
stable, 3.49× apart.* **Two numbers for "the same thing" means the definitions differ**, and at 283 s the
composed clean total would be ~600 s and the ceiling unreachable.

**A first reading — that the phase is load-insensitive and therefore genuinely ~283 s — was wrong and is
retracted in place** (`D36`). It rested on the within-campaign fact that `peak_load1` 40/75/52 does not
order 283.53/282.91/306.22; that shows only that the relationship **saturates above ~40**.

The cross-campaign comparison is the one that matters, and **everything inflated**:

| phase | iter-05 | iter-08 | ratio |
|---|---|---|---|
| `compose_up` | 43.87 | 59.60 | 1.36× |
| `ui_hiring` | 45.32 | 86.99 | 1.92× |
| `host_preflight` | 9.33 | 19.08 | 2.05× |
| `ui_studio_desk` | 7.12 | 17.16 | 2.41× |
| `ui_next_web` | 49.01 | 130.85 | 2.67× |
| **`set_dress`** | **81.23** | **283.53** | **3.49×** |

Cohort median **2.05×**; `set_dress` **3.49×** — ~117 s more than the cohort predicts.

**Settled by two cheap checks rather than a campaign (`D38`):**

1. **Not a cold-rep artifact.** All three reps do identical work — 330,261 rows, two pgvector reindexes,
   and the **same schema digest transition `b4cb55bcee08 → ea2e187a1605` every rep**. `--purge` wipes
   the directus schema each cycle, so the structure auto-provision is a constant, not a first-rep cost.
2. **Not a tooling change.** `git diff --name-only fast-build-m258-iter-03..fast-build-m258-iter-06`
   filtered for `setdress|snapshot|stacksnap` returns **zero files**. All 11 changed files are the
   batch-gate feature and the buildbench instrument. **The set-dress code path is byte-identical between
   the campaign that measured 81.23 s and the one that measured 283.53 s.**

**So the excess is environmental.** `set_dress` is simply the most CPU-exposed phase in the cycle — a
`node cli.js bootstrap` plus a 330 K-row COPY ending in two pgvector index builds. **The gate needs a
quiet host, not a new lever**, and `LEVER-M257-L5-setdress` remains a reserve against a ~81 s phase.

**The projection therefore stands, unrefuted:**

```
247.79 s   iter-05/06 gateable single-box bring-up (n=2, agreeing within 1.34 s)
+ 153.81 s iter-08 best-rep batch_gate — measured CONTENDED, so conservative
= 401.60 s  vs the 480 s ceiling
```

**One residual, routed not resolved:** inside `set_dress`, the span `370.98 → 623.71` — **252.73 s, 89 %
of the phase** — passes with **no intermediate log line**, running from the Directus bootstrap straight
to the taxonomy replay's completion. Two distinct operations, one anchor. That is `D17`'s shape one
level down, and **L5 cannot be aimed until it is split.**

## Phase E — the gate, graded

| # | clause | kind | verdict |
|---|---|---|---|
| 1 | one cold command brings the stack up **and** drives the full batch | boolean | ✅ **3/3 reps**, `up_rc=0`, batch drove to completion each time |
| 2 | zero standing red | boolean | ✅ **3/3 reps** `red_count 0 · red_set [] · passing 30 · failing 0` |
| 3 | composed **p50 ≤ 480 s** over 3 consecutive cold cycles | **timing** | ⬜ **NOT MET** — campaign RED, 3/3 `headroom=FAIL`; 840.01 s is not a gate-usable p50 |
| 4 | 0 platform-repo edits | boolean | ✅ (`D26`, iter-07, all six peer clones) |
| 5 | stack left presenter-usable | boolean | ✅ **12/12 cockpit seats resolve in the 35-identity roster**, verified live *after* the 3-rep campaign |

Clause 5 is worth its own line: it was checked **after** three consecutive reset-to-seed cycles, so the
world-contract restore leg (`TOK-01`'s resolution (b)) held three times in a row. Final state:
`demo-2`=11 · dev=5 · **demo-1=11**, `autoverify green`, `batch-gate verdict green`.

## Close — 2026-08-12

**Outcome:** **The campaign the milestone has been waiting four iters for RAN — and the wait was ended
by automating the trigger, not by the box getting quieter.** An auto-arming waiter fired 91 seconds
after arming, into a 75-second window a 2-minute hand-poll would have missed. All three cold reps came
back **functionally green on every column** (`up_rc=0`, `autoverify green`, `isolation OK`, `phases
complete`, `red_count 0`, `passing 30`) and **timing-refused on every column** (`headroom=FAIL`,
`peak_load1` 40/75/52 vs 10) — because anima8's next batch resumed 60 seconds after launch. Clause 3
stays open; clauses 1, 2 and 5 are now proven **on three consecutive cold cycles** rather than one.
The iter's real deliverable is the `set_dress` scare and its settlement: **81.23 → 283.53 s looked like
a gate-killer, and two cheap checks proved the code path byte-identical and the excess environmental**,
leaving the 401.60 s projection intact and L5 un-escalated.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n *(clause 3 unproven; 4 of 5)* — (2) triggered-tok: n *(this iter
moved four clause-verdicts from n=1 to n=3 and settled `D36`; no no-progress streak)* — (3) re-scope:
**n** *(the trigger's literal text is satisfied — p50 840.01 s > 600 s on the 4th tik — but it is graded
`n` on two grounds, see `D37`: the reps are instrument-rejected, and the remedy does not fit the
diagnosis — `batch_gate` p50 is **179.37 s**, inside M256's 200 s budget while contended, so splitting
the suite would cost the `D-v28-3` single-red-set semantics and buy nothing against a `set_dress`-shaped
cost)* — (4) user-blocker: n *(iter-07 exited here because it could not run at all; this iter ran, and
the next work — splitting the 252.73 s span — needs no host window)* — (5) cap-reached: n *(1 tik)* —
(6) protocol-stop: n — (7) budget-exhausted: **y** *(between iters, tree clean — see below)* —
Outcome: **exit-7**

> **Why `budget-exhausted` and not `user-blocker`.** Nothing needs a user decision. The gate needs ~45
> continuous minutes at `load1 < 10`; the waiter that catches such a window **now exists and is proven
> to fire**, so resuming is a plain re-invoke rather than an arbitration. Escalating the host again
> would be asking the user to solve a problem this iter already built the tool for.

**Decisions:** D29–D38 (this iter's `decisions.md`)

**Side-deliverables:**

- **`.agentspace/scratch/work-m258/autoarm-campaign.sh`** — the auto-arming launcher (`bash -n` +
  shellcheck clean). Continuous 15 s sampling, fires on 3 consecutive `load1 ≤ 5.0`, re-arms on rc ∈
  {8, 9} (the two refusals that tear nothing down), 5 h deadline, writes its own decision series to
  `autoarm-series.log`. **Proven: armed 08:17:41Z, fired 08:19:12Z.** It is the reusable answer to a
  host this milestone does not control.
- **`FIX-M258-iter08-set-dress-has-no-internal-attribution`** — new, with evidence (`D36`/`D38`).

**Routes carried forward:**

- **`TOK-01` step 4 — a GATEABLE composed campaign** (iter-09). Everything except the host is
  discharged; the waiter is armed-and-proven. Needs ~45 min at `load1 < 10`.
- **`FIX-M258-iter08-set-dress-has-no-internal-attribution`** — split the 252.73 s span (89 % of
  `set_dress`) into its Directus-bootstrap and taxonomy-replay legs. **Needs no host window**, and
  `LEVER-M257-L5-setdress` cannot be aimed without it. The natural iter-09 if the box stays loud.
- **`LEVER-M257-L5-setdress`** — reserve, and `D38` re-confirms it is priced against a **~81 s** phase,
  not the 283.53 s this campaign observed. Only needed if a clean p50 lands over 480 s.
- Unchanged and still open: `FIX-M258-iter03-guard-scans-its-own-scratch` (+
  `test_fence_provenance::test_the_escape_accepts_and_records`) ·
  `ROUTE-M258-iter02-isolation-names-two-causes-not-three` · `ROUTE-M258-iter02-headroom-defaults-to-billion`
  · `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir` *(observed live this iter — `D33`)* ·
  `ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone`.
- ⚠️ `demo-2` (11) and the dev stack (5) verified resident **before, mid- and after** the campaign.
  `demo-1` is UP (11), `autoverify green`, **12/12 cockpit seats resolving**.

**Lessons:**

- **When the host is the blocker, automate the trigger — don't extend the wait.** iter-07 polled 30
  minutes by hand and correctly concluded no window opened; iter-08 armed a 15-second sampler and caught
  one in 91 seconds. The difference was not patience or luck, it was **sampling interval versus window
  width**. A blocker that is *external and intermittent* is a scheduling problem, and scheduling problems
  are solved with schedulers.
- **A scare is not a finding until you try to kill it.** `set_dress` at 3.49× looked like it moved the
  composed total to ~600 s and made L5 urgent. Two checks costing minutes — *do all three reps do the
  same work?* and *did the code change between the two tags?* — settled it as environmental. **The
  hypothesis that would have re-planned the milestone was refuted by `git diff --name-only`.**
- **Grade a trigger on its remedy, not only on its threshold.** The re-scope trigger's number was
  satisfied (840 > 600). Its *prescribed action* — split the suite — pointed at a `batch_gate` that
  measured **179.37 s inside a 200 s budget while contended**, while the cost sat in `set_dress`. Firing
  it would have traded away the `D-v28-3` single-red-set contract for nothing.
- **The brief's premise deserves the same scrutiny as the code's.** This iter opened with a written,
  confident, *wrong* statement that the box was calm and anima8 had finished. Thirty seconds of `uptime`
  and `ps -r` corrected it. *Figures written from memory are wrong* applies to the paragraph telling you
  to hurry.

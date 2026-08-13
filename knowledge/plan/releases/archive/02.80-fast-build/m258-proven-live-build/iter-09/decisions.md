# M258 iter-09 — decisions

## D39 — The set-dress split is NESTED. A phase's children are not its siblings.

The obvious implementation was four more rows in `BRINGUP_ANCHORS`. It is wrong, and the reason is
the invariant iter-08 verified in `D32`: **Σ `sub_phases` == `P4_BRINGUP`, exactly.** Anchors are a
flat tiling — each one's timestamp ends the previous phase — so adding `set_dress`'s children to
that list would make them its *peers*, `set_dress` would shrink to its own head segment, and every
second the children describe would be counted twice against the parent that still nominally contains
them.

Two things would have broken silently:

1. the flat table's sum invariant (the one thing iter-08 proved about it), and
2. **comparability** — `set_dress` is the series this milestone has been tracking across campaigns
   (iter-05 **81.23 s**, iter-08 **283.53 s**, and `D38`'s environmental settlement rests on
   comparing them). Redefining the label mid-milestone would have invalidated the comparison that
   settled the release's biggest scare, one iter after it was settled.

So the split lives in its own key (`set_dress_attribution`), aggregates into its own report block,
and prints under a header that says it is *"a SPLIT of `set_dress` above, not more sub-phases"*.
Verified in both directions: `set_dress` still reads 283.53 s in the flat table, and no `sd_*` name
appears in it.

## D40 — The target is named: the taxonomy replay, ~88 % of the phase.

| rep | `sd_replay_taxonomy` | % of `set_dress` |
|---|---|---|
| 1 | 249.69 s | 88.1 % |
| 2 | 258.63 s | 91.4 % |
| 3 | 266.66 s | 87.1 % |

p50 **258.63 s**, contended (`peak_load1` 40.09 / 74.77 / 51.80). Everything else in the phase is
noise by comparison: the Directus bootstrap is 3.11–14.96 s, the seed 12.25–15.66 s, the whole
Directus boot under 9 s.

**`LEVER-M257-L5-setdress` therefore has exactly one target**, and it is a single
`stacksnap replay --surface taxonomy` moving **~1.47 GB** of payload
(`public.skill_embeddings.copy` **825 MB** / `public.job_role_embeddings.copy` **364 MB**, measured
in the local snapshot store) and rebuilding **two** pgvector indexes over 42,790 + 22,470 vectors.

⚠️ **These are CONTENDED figures and must not be quoted as the lever's price.** `D38` established
the phase is environmentally inflated (cohort median 2.05×); against iter-05's quiet
`set_dress` of 81.23 s the same share prices the replay at roughly **70 s**. The share (~88 %) is
the durable finding; the seconds are a dated measurement.

## D41 — Settled retroactively, by re-reading logs we already had. No campaign was spent.

The boundary lines the split needs — `snapshot store root:`, `replayed <surface> into <stack>`,
the Directus bootstrap and health lines — **already existed** in every captured log. What was
missing was a deriver, not instrumentation. So the entire coarse split was obtained by re-parsing
iter-08's three cold reps, at zero host cost, on a box at `load1` 20–37 where no campaign could run.

This is the release's `D27` lesson applied again: *settle by arithmetic when you can.* It is also
why this iter could do useful gate work on a saturated host, which was the premise it opened under.

## D42 — The thesis, demonstrated mechanically: summing is not attributing.

`D17` is the defect where a 166 s phase hid inside a 2 s one **and the phase table still summed**.
This iter's own test suite now proves that failure mode is real rather than rhetorical.

A mutant was introduced into `replay.Run` billing the reindex to the copy (`t.CopyS += tt.ReindexS`):

- `TestTimings_PartsSumToTotal` → **ok**
- `TestTimings_EachPhaseAttributesSeparately` → **FAIL** (`CopyS = 21, want 10`)

The arithmetic check passes a breakdown that describes something that never happened. **Both checks
are required, and neither substitutes for the other.** Same result on the Python side: a mutant
letting the surface regex also match stacksnap's quoted line (the phantom-segment defect) was caught
by three tests.

## D43 — A mutant SURVIVED, and it was a test gap, not a code gap.

Deleting the deriver's lower span bound (`parent.start_s < elapsed`) left all 12 tests green. The
bound is load-bearing in production — a set-dress that ran earlier in the same log (a retry, or a
dev pass before the demo one) emits exactly the cut lines, and stealing them gives the first segment
a **negative** duration — but the fixture had no such lines, so nothing exercised it.

`test_cuts_before_the_parent_span_are_ignored` was added and kills the mutant. Recorded because the
instinct on a surviving mutant is to distrust the code; here the code was right and **the fence was
incomplete**, which is only discoverable by mutating.

## D44 — The literal ratchets are POLLUTED by the demo stack dir. Second member of a known family.

Measuring `RATCHET-M257-literal-ceilings-breached` against the working tree gave DOCSTRING **258**,
COMMENT **237**, TEST_MODULE **672**. Against a pristine `git archive HEAD` extract: **248 / 236 /
657**. The entire difference in two of the three is
`demo-stack/stacks/demo-1/clones/app/studio/**` — **the platform's own Python, inside a demo's
ephemeral clone**, being censused as if it were rext source:

| ratchet | mine | demo-clone pollution |
|---|---|---|
| DOCSTRING | +0 | **+10** |
| COMMENT | +0 | +0 |
| TEST_MODULE | +0 | **+9** |

This is verbatim `FIX-M258-iter03-guard-scans-its-own-scratch` — *"walks `demo-stack/stacks/**`,
which `demo-stack/.gitignore:8` ignores, and reports the platform's source inside a demo's ephemeral
clone"* — reaching a **third** consumer (`derivation_registry`'s three literal censuses), after
`test_decommissioned_instruction_guard` and `test_fence_provenance`. It fires only on a box that has
ever run a demo, which is why it keeps being mistaken for a regression.

**Consequence worth stating plainly:** any ratchet figure measured on this box without excluding
`stacks/` is not a measurement of this repo. M257's recorded 254 / 663 may itself be polluted.
**Routed, not fixed** — the fix is a root-selection change shared by ≥3 consumers, which is a
different iter's scope.

**My own contribution was paid down to +0 on all three**, rather than raised or waived: the
re-quoted row/table counts were dropped from two fixture line tails (marked with `…`), which no
assertion reads. *Never raise a ratchet ceiling.*

## D45 — The replay breakdown carries its residual instead of implying it.

`Timings.UnattributedS` is a named field, computed as `TotalS − Σ(parts)` and printed beside the
parts. The alternative — five phase numbers and a total — invites the reader to assume the parts are
exhaustive, which is precisely how a gap goes unnoticed. If the residual ever grows, that is itself
the finding. `TestTimings_PartsSumToTotal` asserts the identity holds to 0.02 s.

## D46 — The pin write went to the wrong file: a step that succeeded, a relationship that did not.

Re-pinning wrote `fast-build-m258-iter-09` into `stack-demo/rosetta-extensions/.agentspace/rext.tag`
— a stray untracked copy **inside** the consumption clone — while the single source-of-truth every
skill reads is `<rosetta>/.agentspace/rext.tag` (`demo-up/SKILL.md:39`, *"the single
source-of-truth, M49 #1"*). The write succeeded, the file existed, the content was right, and the
pin was still stale. Caught by asking *which file does the reader read*, then asserting the two
agree — not by checking that the write worked. `D19`/`D20`'s lesson, hit live again in the same
milestone.

## D47 — Clause 3 is armed, not awaited.

The host was at `load1` **20.31–37.27** throughout this iter. Rather than poll (iter-07's failure)
or wait (iter-08's opening premise), `autoarm-campaign.sh` was re-armed at **09:40:37Z** against a
**fresh** output dir. That last part is not cosmetic: `build_report` globs `rep-*/ledger.json`, so
pointing a new campaign at `campaign-iter07/` would have silently aggregated six reps — three of
them instrument-rejected — into one p50. The launcher and its log path are now parameters.

The campaign now runs at the `fast-build-m258-iter-09` pin, so its ledgers will carry the set-dress
attribution live, at whatever load the window offers.

# M258 iter-09 — progress

**Type:** tik · **Active strategy:** `TOK-01` (bootstrap) — *measure the composition before
engineering it*, applied one level down: the largest phase in the cycle was a single un-attributed
span, and it had to be measured before a lever was spent on it.

## Phase 0 — re-survey, and the substitution

`TOK-01`'s next-tik direction is step 4 (the composed 3× cold campaign). iter-08 **ran** it, so the
direction is not stale — but it returned `headroom=FAIL` 3/3 and its only remaining input is ~45
minutes at `load1 < 10`, which is not a work item. Measured at 09:22:06Z: **`load1 24.86`**
(1/5/15 = 24.86 / 20.88 / 22.57), rising to **37.27** by 09:33 and never below **19.72** all iter.

Target substituted **inside** the strategy, per iter-08's own routing — *"the natural iter-09 if the
box stays loud"* — to `FIX-M258-iter08-set-dress-has-no-internal-attribution`. Step 4 was **armed**
(Phase C), not abandoned.

## Phase 0d — pre-flight (the gate tools, before touching them)

139 buildbench tests green; `stack-snapshot` `go build ./...` + `go test ./replay/... ./cmd/...`
green. Interpreter noted: `python3` on this box is **3.14.6 and has no pytest**;
`/usr/bin/python3` (3.9.6) does. That is `FIX-M257-census-interpreter-namespace-import`'s territory
and it cost a round trip here too — recorded, not fixed.

## Phase A — the split, proven RETROACTIVELY on three real cold reps

The boundary lines were **already in the logs**; what was missing was a deriver (`D41`). So the
coarse split was obtained by re-parsing iter-08's three captured reps at zero host cost:

| segment | rep 1 | rep 2 | rep 3 |
|---|---|---|---|
| `sd_directus_preamble` | 2.59 | 1.37 | 5.54 |
| `sd_directus_bootstrap` | 3.11 | 3.39 | 14.96 |
| **`sd_replay_taxonomy`** | **249.69** | **258.63** | **266.66** |
| `sd_replay_directus` | 2.36 | 0.72 | 0.80 |
| `sd_directus_boot` | 8.90 | 4.62 | 4.67 |
| `sd_replay_sim_embeddings` | 0.29 | 0.16 | 0.23 |
| `sd_seed` | 15.66 | 13.07 | 12.25 |
| `sd_post_seed` | 0.93 | 0.95 | 1.11 |
| **Σ** | **283.53** | **282.91** | **306.22** |
| parent `set_dress` | 283.53 | 282.91 | 306.22 |
| **residual** | **0.0** | **0.0** | **0.0** |

**The 252.73 s span is one operation, not two** — the hypothesis held. The taxonomy replay is
**87–91 %** of the phase (`D40`); everything else is sub-16 s.

Both directions checked, because summing and attributing are different claims (`D42`):

- **sums** — residual 0.0 on all three, and the segments are contiguous (each starts exactly where
  the previous ended, asserted per-pair, so a sum cannot be right *through* an overlap-plus-gap);
- **attributes** — eight distinct values, the largest 88.1 %, the smallest 0.1 %.

And the flat table is **undisturbed**: Σ `sub_phases` == `P4_BRINGUP` still exact
(802.39 / 822.19 / 840.91 — reproducing iter-08's `D32` figures), `set_dress` still reads its full
283.53 s, no `sd_*` name leaks into it (`D39`).

## Phase B — the level below, which no captured log can supply

`replay.Run` moved ~1.47 GB and rebuilt two pgvector indexes while emitting **nothing** between
start and finish. It now attributes its five documented phases (verify / clear / copy / reindex /
advance-sequences) with per-table detail and an **explicit** unattributed residual (`D45`), printed
by `stacksnap` on every replay.

This is the number `LEVER-M257-L5-setdress` actually needs: **copy versus reindex**. The two have
entirely different remedies (a bulk-load path versus an index-build strategy), and the coarse split
above cannot distinguish them. Unmeasured as of this iter — it needs a run, and the armed campaign
is that run.

## Phase B′ — mutation testing, and what it found

| mutant | sum test | attribution test |
|---|---|---|
| reindex billed to copy (`t.CopyS += tt.ReindexS`) | **ok** | **FAIL** |
| surface regex also matches stacksnap's quoted line | — | **FAIL ×3** |
| deriver's lower span bound deleted | — | **SURVIVED → fixed** |

The first row is the iter's thesis made mechanical (`D42`). The third was a **test** gap, not a code
gap (`D43`) — fixed by `test_cuts_before_the_parent_span_are_ignored`, which now kills it.

## Phase B″ — the ratchets, and a pollution finding

Measuring `RATCHET-M257-literal-ceilings-breached` on the working tree vs a pristine `git archive
HEAD` extract showed that **two of the three ratchets are polluted by
`demo-stack/stacks/demo-1/clones/app/studio/**`** — the platform's own Python inside a demo's
ephemeral clone (`D44`). DOCSTRING +10, TEST_MODULE +9, none of it rext's. Third consumer of the
`FIX-M258-iter03-guard-scans-its-own-scratch` family. **Routed, not fixed.**

**My own contribution: +0 on all three**, paid down rather than waived — two fixture line tails
re-quoting row/table counts were truncated at a marked `…`, which no assertion reads.

## Phase C — armed, not awaited

Committed `fe50735`, tagged **`fast-build-m258-iter-09`**, **pushed and verified on origin**
(`git ls-remote --tags origin` → 1 hit; rung zero). Consumption clone re-pinned and the **feature
verified present at the pin** (the M236/`D21` check: both new test files exist, the deriver is
referenced 4×) — a tag that exists is not a tag that carries the feature.

⚠️ The pin write initially went to a stray copy **inside** the consumption clone rather than the
single source-of-truth the skills read; caught by asserting the two agree (`D46`).

The user's stacks were verified resident **before and after** every operation near them:
`demo-2`=11 · dev=5 · `demo-1`=11, throughout.

`autoarm-campaign.sh` re-armed **09:40:37Z** against a **fresh** `campaign-iter09/` output dir —
`build_report` globs `rep-*/ledger.json`, so re-using iter-08's dir would have silently aggregated
six reps, three of them instrument-rejected, into one p50 (`D47`).

## Phase E — the gate, graded

| # | clause | verdict this iter |
|---|---|---|
| 1 | one cold command brings the stack up **and** drives the full batch | ✅ unchanged — 3/3 reps (iter-08) |
| 2 | zero standing red | ✅ unchanged — 3/3 reps `red_count 0` |
| 3 | composed **p50 ≤ 480 s** over 3 cold cycles | ⬜ **NOT MET** — no window opened; `load1` 19.72–37.27 all iter. Armed. |
| 4 | 0 platform-repo edits | ✅ unchanged — 0 platform files touched this iter (rext + plan only) |
| 5 | stack left presenter-usable | ✅ `demo-1` UP (11), never torn down this iter |

**The projection is unchanged and unrefuted at 401.60 s vs 480**, and this iter did not move it —
by design. What it moved is the *aimability* of the reserve behind it.

## Close — 2026-08-12

**Outcome:** **`LEVER-M257-L5-setdress` has a target for the first time.** The 252.73 s span that
was 89 % of the largest phase in the cycle turned out to be **one operation, not two**: the taxonomy
replay, **87–91 %** of `set_dress` across three real cold reps, tiling the parent with **residual
0.0** on every one. And it was settled **retroactively, on logs already captured**, at zero host
cost, on a box that never dropped below `load1` 19.72 — the boundary lines existed; only the deriver
was missing. The level below shipped too (`replay.Run` now attributes copy-vs-reindex per table with
a carried residual), unmeasured until the next run. **The iter's most transferable result is a
mutation:** billing the reindex to the copy leaves the **sum test passing** and fails only the
attribution test — `D17`'s lesson stops being a story and becomes a fence.

**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n *(clause 3 unproven; 4 of 5 stand from iter-08)* — (2)
triggered-tok: n *(a tik, and the last three tiks all showed progress)* — (3) re-scope: n *(no new
p50 was produced at all this iter; iter-08's `D37` grading is unchanged and nothing new bears on
it)* — (4) user-blocker: n *(nothing needs a user decision — the host is armed for, not escalated;
and this iter's whole scope needed no host at all)* — (5) cap-reached: n *(1 tik)* — (6)
protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**

**Decisions:** D39–D47 (this iter's `decisions.md`)

**Side-deliverables:**

- **`.agentspace/scratch/work-m258/launch-iter09-campaign.sh`** — the campaign launcher against a
  **fresh** output dir, and `autoarm-campaign.sh` parameterised (`LAUNCH` / `LAUNCHLOG`) so a
  campaign can never again write into a previous campaign's ledger glob (`D47`).
- **`ROUTE-M258-iter09-literal-ratchets-scan-the-demo-clone`** — net-new finding with evidence
  (`D44`), routed.

**Routes carried forward:**

- **`TOK-01` step 4 — a GATEABLE composed campaign** (iter-10). Armed at the new pin; needs ~45 min
  at `load1 < 10`. ⚠️ Fresh output dir per campaign.
- **`SPLIT-M258-iter09-copy-vs-reindex`** — the level-two instrument is shipped but **unmeasured**.
  Do not assert COPY-bound vs REINDEX-bound until a run prints the line; the remedies differ.
- **`ROUTE-M258-iter09-literal-ratchets-scan-the-demo-clone`** — third consumer of
  `FIX-M258-iter03-guard-scans-its-own-scratch`'s root cause; the shared fix is a root-selection
  change, and it needs no host window.
- **`LEVER-M257-L5-setdress`** — still unspent, still not needed (401.60 s vs 480), but now priced
  against a named operation rather than an opaque span.
- Unchanged and still open: `FIX-M258-iter03-guard-scans-its-own-scratch` (+
  `test_fence_provenance::test_the_escape_accepts_and_records`) ·
  `ROUTE-M258-iter02-isolation-names-two-causes-not-three` ·
  `ROUTE-M258-iter02-headroom-defaults-to-billion` ·
  `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir` ·
  `ROUTE-M258-iter07-demopatch-G5-does-not-revert-the-native-clone`.
- ⚠️ `demo-2` (11) and the dev stack (5) verified resident before and after every operation.
  `demo-1` UP (11), never torn down this iter.

**Lessons:**

- **The instrument you need is often already in the log.** Four iters treated `set_dress` as opaque
  and priced a lever against a guess. The boundaries had been printed in every run since the feature
  shipped; what was missing was something that read them. **Before instrumenting, re-read.**
- **A surviving mutant is not automatically a code bug.** Deleting the deriver's lower span bound
  left every test green — and the code was *right*; the fence was incomplete. Mutation testing finds
  gaps in tests as often as in logic, and the instinct to "fix the code" would have broken a correct
  guard.
- **When adding a child to a measured thing, the risk is to the PARENT's series.** The tempting
  implementation would have silently redefined `set_dress` one iter after `D38` used that exact
  series to settle the release's biggest scare. *Adding detail must not re-cut the number people are
  already tracking.*
- **A write that succeeds is not a pin that moved.** The re-pin wrote the right content to a real
  file that no reader reads. The only check that catches it is *which file does the consumer read,
  and do the two agree* — steps versus relationships, again.

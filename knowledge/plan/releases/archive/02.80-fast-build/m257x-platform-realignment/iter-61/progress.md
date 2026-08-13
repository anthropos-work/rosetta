**Type:** tik — under [`TOK-05`](../decisions.md). Step 0 **substituted** the routed target; see
`D-M257x-61-5`.

# iter-61 — the fence could not see the class's larger half

## What happened

iter-60 fenced the profile predicate in its **command** form (`PROFILE=x`, `--profile x`) and its
**table-first-cell** form, repaired those to zero, and read **GREEN**. Re-surveying the same tree at
the same platform ref found the class's **larger half** still standing in a form the guard had no
construct for: the **noun phrase** — *"the default `graphql` profile"* — **34 raw sites across 17
files**.

**A fence whose reach is narrower than its class over-reports its own GREEN** — and the over-report
is invisible, because the fence is the thing you would use to check.

It also vindicated the briefing. `D-M257x-60-7` recorded the handed-down *"17 files / 30
occurrences"* as an **undercount**, on the strength of a broader grep. It was not: **17 files** is
exactly what the noun-phrase construct yields, and it is the construct that carries the class. That
row of `D-M257x-60-7` is **withdrawn** (`D-M257x-61-1`); its two `main.go` line-number corrections
stand.

## The two new constructs, and why they are still constructs

| form | shape | why iter-60 missed it |
|---|---|---|
| **noun phrase** | a backticked token adjacent to the literal word `profile`, or `profile (\`tok\`)` | no command verb anywhere near it |
| **table row** | `\| \`CMS_RPC_ADDR\` \| \`http://cms:8091\` \| … \|` | states the binding with **no `=`**; `messenger.md` held two stale values in this shape the entire time the fence read GREEN |

Neither is a substring widening. **"GraphQL" the API — named constantly in this corpus — still cannot
match**, because the token must sit inside a backtick span *and* be adjacent to the word `profile`.

## Prose can do two things a command line cannot

35 raw hits → **22 real; 13 were the guard's own**, all removed by rules derived from the corpus's
own writing rather than by exceptions (`D-M257x-61-3`):

* **Negation** — *"and no `cms` profile"* asserts the token's **absence**. Anchored to the end of the
  preceding text so only an *adjacent* negation counts. **The recursion is the interesting part:**
  this is the exact form iter-60's own corrections are written in (`D-M257x-60-5` made the corpus say
  *"there is no cms profile"*), so an undiscriminating widening reads **this milestone's repairs** as
  fresh defects — a fence that punishes the fix it asked for.
* **Ref-pin** — the exemption G2/G4/G5 already had and G1 did not, plus the **bare backticked sha**
  (`` `b56d731` only parked the block behind a `wundergraph-deprecated` profile ``), which is how the
  corpus opens a sentence about a commit and which carries most of the historical narrative.

## Where it stands

**RED at 2 findings / 22 sites / 12 files**, enumerated in
[`evidence/residual.md`](evidence/residual.md) with the command to regenerate it. **Not repaired
here, and deliberately not repaired in part** — §5 rule 19's scope-edge corollary says a claim leaks
to the edge of the previous repair's scope and pools there, so a subset repair leaves a
half-consistent corpus that costs the next auditor its budget in adjudication. Routed **whole**.

**RED and correct beats GREEN and narrow.** The instrument landing ahead of its repair is the honest
split, not a shortfall in it.

35 tests (was 28), including a regression that iter-60's GREEN fixture stays green under the widened
rules. rext tagged `fast-build-m257x-iter-61`, **verified on origin**; pin advanced.

## Close — 2026-08-04

**Outcome:** G1/G4 widened to the noun-phrase and table-row constructs, which exposed that iter-60's
GREEN covered only the smaller half of its own predicate class — **22 real sites / 12 files** remain,
now enumerated and routed whole. 13 of 35 raw hits were the guard's own and were removed by two
derived discriminators (adjacent negation, ref-pin). The briefed "17 files" figure was **vindicated**
and `D-M257x-60-7`'s contrary row withdrawn.
**Type:** tik
**Status:** closed-fixed-partial — the instrument (widened fence + 7 new tests) landed; the repair it
names did **not**, and is routed whole rather than in subsets.
**Gate:** NOT MET — 4 of 5, unchanged. Clause 5 remains the only open one.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-61-1` (a narrow fence over-reports its own GREEN; the briefed figure
vindicated) · `D-M257x-61-2` (two new constructs, still constructs) · `D-M257x-61-3` (negation +
ref-pin; 13 of 35 were the guard's own) · `D-M257x-61-4` (route the residual WHOLE) ·
`D-M257x-61-5` (Step-0 substitution, and the citation class re-measured before re-routing)
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter61-profile-prose-class` → **iter-62**, handler = the enumerated 22 sites in
  `evidence/residual.md`. Repair **whole**.
- `FIX-M257x-iter58-mainline-shift` → **iter-62**, with a **refreshed** measurement (5 of 16 distinct
  `app/main.go:N` citations still land on their claimed construct at app `v1.366.0`; the rest moved).
- `DOC-M257x-iter59-storage-mid-fold` → **iter-62**: the map's 8th state token. The *measurement*
  landed in iter-60 (`storage.md`, G6-fenced); the vocabulary + assertion-C change did not.
- `CHECK-M257x-iter60-stale-pin-exemption` → open, and **now load-bearing**: a ref-pin exempts a
  claim from G1 and G4 alike, so `messenger.md:108-109`'s two stale RPC values sit behind an
  `@ 2adcf71` pin and the fence cannot reach them. The pins are printed; turning "pinned to a
  **superseded** ref" into a finding is the fix and is not built.
- `CHECK-M257x-iter60-g6-citation-subject` → open.
- `FIX-M257x-iter53-union-set` → **PENDING USER DECISION**, untouched.
- `FIX-M257x-iter56-assignment-flake` → **NOT DECIDED**; needs a failure *rate*.
- `CHECK-M257x-iter38-ai-act-classification` → needs an owner outside this milestone.
- Unchanged: `-cold-daemon-registry` · `-grep-vs-failclosed` · `-empty-stdout-class` · `-baseline-refs` ·
  `CHECK-M257x-iter58-derive-preregistrations` · `FIX-M257x-iter57-within-block-drift` ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `CHECK-M257x-iter52-second-ai-manager` · RF-2/3/7–13.

**Lessons:**

1. **Enumerate a predicate's FORMS before believing a fence's GREEN.** The same predicate was written
   four ways here — command, table first cell, noun phrase, table row — and a fence that covers two
   of them reports the same GREEN as one that covers four. The reach line must name the *forms*, not
   just the site count.
2. **A widened fence will read the previous iteration's corrections as defects.** iter-60 repaired by
   writing *"there is no `cms` profile"*; iter-61's noun-phrase rule sees a documented `cms` profile.
   Any rule that widens across a repaired class needs a negation discriminator, or it punishes the
   fix it asked for.
3. **RED and correct beats GREEN and narrow** — and shipping the instrument ahead of the repair is a
   legitimate split, provided the residual is enumerated, regenerable, and routed whole.
4. **Re-measure a routed item before routing it on.** The citation class carried iter-58's "21 of 22";
   re-resolved at `v1.366.0` it is 5 of 16 holding. Passing a stale count forward is the same defect
   as inheriting one.

**Type:** tik — under `TOK-08`, on the user redirect (`D-M257x-222-1`), gate-clause-1 work.

## Probe — sealed before any fence

See [`probe-evidence.md`](probe-evidence.md) and the pre-registered P1–P6 table in
[`overview.md`](overview.md). Sealed as this iter's first commit.

## What the census found

| # | pre-registered | measured | verdict |
|---|---|---|---|
| P1 | **23** manifests | 23 | held |
| P2 | 4 repos — app 2 · next-web-app **11** · studio-desk 5 · ant-academy 5 | same | held **after in-flight correction** (first written as `9`; `2+9+5+5 = 21 ≠ 23`, and the arithmetic caught it before the seal — the two `next-hiring-*` manifests target `next-web-app`, their prefix naming the app inside the monorepo) |
| P3 | **0** paths missing at `origin/main` | 0 | held |
| P4 | anchor occurs exactly once: **23 of 23** | 23 / 23 | held |
| P5 | **10 of 23** `pre_sha256` no longer match | 10 | held |
| P6 | of those 10, **0** would be REFUSED | 0 | held — `assert_pre_patch` returns `pristine` on sha drift with the anchor intact 1× |

**The answer to the question that motivated the iter: advancing the clones would not break the patch
layer.** 23/23 anchors resolve at `origin/main` — against clones iter-222 measured at `app` **+28**,
`next-web-app` **+12**, `ant-academy` **+9** behind it.

**And the finding the probe did not pre-register, because it took a second reading to see:** the guard
run at `HEAD` returns **the same 10 drifts**, and the two sets are **identical member for member**. Those
baselines were **already** stale at the ref the demo builds today — the platform advancing did not make
them so. At least one is stale **by design**: `§6`'s documented chain case, where
`next-web-public-website-url`'s `pre_sha256` *is* studio's `post_sha256` and reads DRIFTED against a
pristine file on purpose. **Which of the remaining nine are chain cases and which are simply un-repinned
is NOT settled here**, and the guard is deliberately silent on the cause — it reports the count.

## What landed

1. **`stack-core/patch_anchor_guard.py`** (FENCE-M257x-iter223). Two arms, per manifest, at a **named
   ref**: the path exists; the anchor occurs **exactly once**. These are precisely the two conditions
   `demopatch`'s own G2 refuses on, hoisted from a per-manifest apply-time check to a whole-set
   pre-flight. Seconds, no build — against ~11 minutes and a warning line to notice.

2. **The design decision, and it is the one worth defending: sha drift is counted, printed *inside the
   verdict line*, and is NEVER a finding.** Since M217 the gate self-heals — drifted sha + intact anchor
   is `pristine`, WARNed, applied — so a fence reddening on drift would contradict the shipped mechanism
   and go RED on a set that demonstrably works. A fence that cries wolf gets suppressed, which is worse
   than no fence.

3. **`tests/test_patch_anchor_guard.py`** — 15 tests. A clean sweep is only readable next to the
   mutations that prove the instrument, so every verdict the guard can reach is fired against a
   synthetic git repo: anchor **GONE**, anchor **AMBIGUOUS**, path **GONE**, and a broken sibling
   alongside a green member. Four anti-vacuity controls, plus a test that pins the **non**-behaviour
   (drift stays green) — *"we chose not to check X" is invisible unless something says so*.

4. **Registered in `guard_family.INVOCATIONS`**, with the clones root **derived from `--platform`**
   rather than taking its own flag, so the five platform-facing guards cannot end up green about
   different trees. The registry proved itself first again: **exit 2**, *"patch_anchor_guard is on disk
   but this runner has no invocation for it."*

5. **Docs:** `demopatch-spec.md` §6-bis (the pre-flight + both readings + the reach statement);
   `stack-core/README.md` row.

## The three things the fences did to this iter's own work

Worth recording together, because all three were caught by machinery rather than by review:

- **`platform_predicate_guard` went RED on iter-222's §8 prose.** *"the pin named **11** repos"* — its G2
  arm reads a number beside the word `repos` as a claim about the clone set, and `repos.yml` has 4. The
  fence could not know the sentence was about the *pin's* population. Repaired by **rephrasing, not
  waiving**: *"the pin declared **11 keys**"*, with the reason recorded in place. That is `§8` iter-98
  — *write the claim in the vocabulary the fence enumerates* — applied to a fence's own documentation.
- **Three literal ratchets breached, and the arrows had to be SPLIT.** `DOCSTRING` 210→212→**218**,
  `TEST_MODULE` 567→569→**581**, `COMMENT` 178→**184**. The first arrow of each is this iter's own
  writing; the second is **reach**, because the iter also widened `_MEASURED_NOUNS` by four words
  (`manifests`/`paths`/`shas`/`drifts` — the sixth time the vocabulary has closed on the sentence that
  widened it). **Isolated the iter-217 way, holding the TREE fixed and varying the MATCHER:** the
  pre-widening registry run against *this* tree returns exactly **212 / 178 / 569**, the post-widening
  one **218 / 184 / 581**. So none of the second arrow is prose — and `COMMENT` is the clean control,
  where the old matcher returns the old ceiling *exactly*, meaning this iter's comments contributed
  **zero** to it.
- **The derivation registry rejected an entry in BOTH directions inside one iter.** `patch_anchor_guard
  ::manifests` *"became executable-here and was not graded"* → graded. Then `::check`, added out of
  symmetry with `clone_pin_guard::check`, was rejected by the *"a decision must not outlive its site"*
  arm — it is not in the executable-here population, because unlike its sibling its arguments are not
  all path-like. **Grading a site that does not exist is the same defect as failing to grade one that
  does**, and the registry caught both within the hour.

## A correction to iter-222, one iter old (and it is the milestone's own lesson landing again)

iter-222 closed on *"the floor caught 9; the cluster was **15**."* **It is 17.** The whole-section run
this iter took surfaced two more, in a document iter-222 had already edited:

| site | cited | quotes | actually at |
|---|---|---|---|
| `ai-readiness.md:26` | `readiness.go:770` | `maxLevel := int(m.workforce.LevelsCount(ctx, orgID))` | **776** |
| `ai-readiness.md:503` | `readiness.go:684` | `queryInCycleStep1Completers` | **690** (the `func`; the name also occurs at 687 and 721) |

Both re-derived and repaired here. Two things are worth separating:

1. **A different instrument found them.** `anchor_subject_census` checks that the cited line carries the
   **literal the prose quotes** — the *"resolves to the WRONG construct"* class `anchor_construct_guard`
   states it cannot reach. iter-222 routed that class forward as a gap
   (`ROUTE-M257x-222-anchor-guard-floor-leaves-siblings`) **while an instrument for part of it was
   already in the suite and already RED.** The route is not wrong, but its framing was: the gap is
   narrower than iter-222 described, and one more reading of the failing-module list would have shown it.
2. **`15` was stale before its own iter closed.** iter-222's whole-section run was captured through a
   truncated log — 14 `FAILED` lines survived of 19 — and these two were among the five lost. A number
   read off an incomplete capture is not a measurement, and *"a number can go stale inside the iter that
   writes it"* (`§5`, iter-206) has now happened to this milestone with the author present.

## Close — 2026-08-09

**Outcome:** the question *"would advancing the platform clones break the demo's patch layer?"* had no
cheap answer and now has a standing one — 23/23 anchors intact at both `HEAD` and `origin/main`, with the
10 sha drifts shown to predate the platform's advance entirely.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: y — Outcome: exit-7
**Decisions:** D-M257x-223-1 (drift is counted, never a finding), D-M257x-223-2 (the cause of the 10 drifts is not adjudicated here).
**No `N` movement is claimed** — this iter took no graded reading.

**Side-deliverables:** the `platform_predicate_guard` rephrasing of iter-222's §8 paragraph (a repair to
the *previous* iter's prose, surfaced by this iter's family run); `_MEASURED_NOUNS` widened by four; two
further rotted anchors in `ai-readiness.md` re-derived (the iter-222 correction above); and three stale
published figures re-derived where the tree moved under them — `claim_census_guard`'s basename share
**292 → 293 of 704**, and the fence-index reach triple **17/28 · 17/27 · 16/27 → 19/30 · 19/29 · 18/29**
(both new guards entered the family AND the prose index in their own commits, so numerator and
denominator moved in lockstep — the shape that does not widen the blind spot).

**Suite state at close** — `stack-core`, pytest 8.4.2 / `/usr/bin/python3`, Python. The whole-section run
came back **1,908 passed / 3 failed / 3 skipped**; all 3 were the figures and anchors named above, each
repaired and re-run green (`test_anchor_subject_census_m257x` + `test_frozen_expectation_census_m257x`
**126 passed**; `test_claim_census_substrate_m257x` **34 passed**; `test_fence_registry_population_m257x`
**16 passed**). **No post-repair whole-section re-read is claimed** — what is claimed is those four
modules and the family verdict. `guard_family` (with `--platform`): **24 GREEN · 0 RED · 5 not-run**, and
"not-run" is 5 commit/ledger-scoped members with no input supplied, which is *not* a whole-family green
and the runner says so itself.

**Routes carried forward:**
- `ROUTE-M257x-223-classify-the-ten-drifted-baselines` → a later tik. Of the 10, `§6`'s urls.ts chain
  explains at least one **by design**. The rest are candidates for `demopatch --repin`, which refuses
  unless the pre-image round-trips — so the classification is mechanical, and it is not this iter's.
- `ROUTE-M257x-222-pin-advance-needs-a-reproof` → **narrowed, not closed.** The patch layer is now
  proven safe for an advance. What is still unproven is everything else a bring-up touches; the advance
  still needs gate clause 1's three cold cycles.
- `ROUTE-M257x-222-other-clones-never-fetched` and
  `ROUTE-M257x-222-anchor-guard-floor-leaves-siblings` — both still open, unchanged.

**Lessons:**
1. **A per-item check is not a set check, and the difference is a whole bring-up.** `demopatch preflight`
   had answered this question for one manifest since M217. Nobody could ask it of the 23, so nobody did.
   When a guard exists at the item grain, ask what the *set*-grain question costs to answer.
2. **State the ref inside the verdict.** The same guard returns two different, both-true answers at
   `HEAD` and at `origin/main`; a verdict that omitted which one it took would be unreadable a day later.
3. **A ratchet arrow that mixes writing with reach hides both.** Split it, isolate the reach by holding
   the tree fixed and varying the matcher, and name the control case where the two can be told apart.

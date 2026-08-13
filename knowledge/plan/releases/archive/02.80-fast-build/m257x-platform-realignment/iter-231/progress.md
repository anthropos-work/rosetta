**Type:** tik — under `TOK-08`, closing the open half of `ROUTE-M257x-228`.

# iter-231 — the corpus's CURRENT-ref claims

## Two populations, and the difference is the whole iter

A **broad** pattern (any sha within 90 characters of `origin/main` or `HEAD`) returns **75 sites / 28
distinct shas**, and most of them look damning: `ad9f3c49` at 21 sites against an `app` `origin/main` of
`3eaadae6`, `2035f9a4` at 6, `8297c684` at 4 against `next-web-app` `19423a1f`.

**Nearly all of it is correct writing.** Those are **dated, past-tense** sentences — *"`ad9f3c49`, which
was `origin/main` on 2026-08-06"* — the exact form this milestone spent iters teaching the corpus to use.
Grading them stale would punish the fix. This iter's own `overview.md` sealed that as an escalation
condition before the first match was run, which is the only reason the broad number was not published.

The **strict** population — present-tense assertions that a sha *is* a repo's current ref — is **6 sites**:

| verdict | site | claim |
|---|---|---|
| **AGREES** | `platform-migration-status.md:118` | roadrunner HEAD `87d8d44` |
| **AGREES** | `platform-alignment.md:1344` | studio-room HEAD `aeec036` |
| **AGREES** | `CLAUDE.md:254` | app is `3eaadae6` |
| **AGREES** | `CLAUDE.md:332` | app is `3eaadae6` |
| **UNMEASURED** | `observability.md:9` | `ant-observability` HEAD `b49eb7af` — repo in no clone set (iter-230's partition) |
| **STALE** | `platform-alignment.md:1476` | *"Origin HEAD `2adcf71`"* — platform `origin/main` is `0c91421d`, **6 commits ahead** |

## The one defect, and why the repair is not a bump

`platform-alignment.md:1476` sits inside a past-tense narrative about iter-21/22, and the sha is **right**:
the paragraph's entire argument is that iter-21's corrections would have been false *at that ref*. What was
wrong is the **noun phrase** — *"Origin HEAD `2adcf71`"*, undated and present-tense, asserting a positional
fact that expired 6 commits ago.

**Repaired by dating the phrase, not by bumping the sha.** Bumping it would have destroyed the paragraph's
subject to fix its tense — the mirror image of the anchor-repair rule (*repair the citation against the
subject, never bump the offset*), and worth naming because the instinct runs the other way when the stale
thing is a ref rather than a line number.

`§5`'s *a clause that WAS a measurement, copied forward until it reads as a property* — in its smallest
possible form: two words of missing date.

## The extractor, proved both ways

`§9`: a census reporting one finding must show it can see the class **and** that it does not over-fire.

- **Fires** on all four present-tense shapes (`origin/main is X`, `HEAD X`, `X is now origin/main`,
  `it is X as of …`).
- **Correctly skips** the three dated past-tense forms it is required to skip — including the exact
  `CLAUDE.md` sentences that dominate the broad population. Without control B the strict count of 6 would be
  indistinguishable from a broken regex.

### Predictions, graded — 2 HELD, 2 REFUTED

| id | prediction | result |
|----|-----------|--------|
| `P-231-1` | ≥ 3 distinct repos carry a current-ref claim | **HELD — 5** (roadrunner, studio-room, app, platform, ant-observability) |
| `P-231-2` | ≥ 1 repo has two **different** shas claimed as current in different docs | **REFUTED — 0.** Both `app` sites name the same sha |
| `P-231-3` | ≥ 1 claimed-current sha disagrees with the clone's actual `origin/main` | **HELD — 1** (`2adcf71`, platform) |
| `P-231-4` | for ≥ 1 repo the corpus is **split**, one doc right where another is wrong | **REFUTED — 0**, and the reason is that **iter-228 already repaired the only known instance.** A prediction written from a three-iters-old finding, against a corpus that had since moved |

## Close — 2026-08-10

**Outcome:** the class iter-228 found by hand is now enumerated. **6 present-tense current-ref claims: 4
agree, 1 unmeasurable, 1 stale** — and the broad 75-site reading that looks alarming is almost entirely the
corpus writing its refs correctly, in the dated past tense this milestone taught it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-231-1` (the strict population is published and the broad one is disclosed, not buried),
`D-M257x-231-2` (a stale ref phrase is repaired by DATING it, never by bumping the sha).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Suite state at close** — no pytest section run; this iter changed no rext code.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-228-corpus-disagrees-with-itself-about-refs` → **CLOSED by measurement.** The corpus does
  not currently disagree with itself about any repo's current ref: 0 repos carry two different
  claimed-current shas. The one instance iter-228 found was real and is repaired; the class is empty today.
- `ROUTE-M257x-231-ref-claims-have-no-fence` → **new.** The 6-site population is small, the extractor is
  written and proved both ways, and a standing guard would catch the next `2adcf71` the day it is typed.
  Second deliverable → tripwire → routed.
- All prior routes (`230-82-sites-on-uncloneable-evidence`, `230-sha-census-has-no-fence`,
  `229-anchor-rot-is-19-of-22-invisible`, `225-no-profile-for-sanctioned-host`,
  `225-hostprofile-role-strings…`, `222-pin-advance-needs-a-reproof`,
  `223-classify-the-ten-drifted-baselines`, `224-drift-guard-blind-to-stale-clone`,
  `227-archived-repo-selfdesc-is-stale`) → open, unchanged.

**Lessons:**
1. **A stale REF is repaired by dating it; a stale ANCHOR is repaired by re-finding it.** Opposite moves,
   and the instinct generalises the wrong one. When the sha is the paragraph's *subject*, bumping it
   destroys the claim to fix the tense.
2. **Seal the exclusion, not just the prediction.** The escalation condition *"a dated past-tense claim is
   not a current-ref claim"* was written before the first match ran, which is the only reason a 75-site
   scare number was never published. An exclusion invented *after* seeing the data is indistinguishable
   from motivated reasoning.
3. **A prediction written from an old finding grades the old corpus.** `P-231-4` was refuted because
   iter-228 had already repaired the instance three iters earlier. Re-survey covers the target; it does not
   cover the *belief*.

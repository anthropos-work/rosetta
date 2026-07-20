# Release Retro — v2.5 "the playbill" (M229 → M236)

**Shipped 2026-07-20.** The **Content stories** release: *take REAL production job-simulation sessions,
scrub and re-tenant them into a demo org, project them into a second cockpit tab, and PROVE they land
non-empty for both the player and the manager — 29/29 live on `billion`, cold reset-to-seed.* 8 milestones
(**M229 → M230** [academy thread] **→ M231 → M232 → M233 → M234 → M235 → M236** [content-stories thread]),
tooling + docs only, **0 platform-repo edits across all 8**. Consolidates the eight milestone retros
([M229](m229-academy-content-model-re-ground/retro.md) · [M230](m230-academy-demo-fill/retro.md) ·
[M231](m231-content-stories-feasibility-spike/retro.md) · [M232](m232-session-clone-sourcing-seeder/retro.md) ·
[M233](m233-content-stories-manifest/retro.md) · [M234](m234-content-stories-cockpit-tab/retro.md) ·
[M235](m235-prove-it-lands/retro.md) · [M236](m236-prove-on-billion/retro.md)), the release-level
[`release-review.md`](release-review.md), and the fate ledger [`release-deferrals.md`](release-deferrals.md).

## What shipped

Two threads, one spine.

- **The academy re-ground (M229–M230).** M229 corrected a **materially-stale** `ant-academy.md` (+ 3 siblings)
  from a false "no backend / static JSON / only Clerk" model to the true DB-authoritative catalog — the doc
  whose wrongness had **mis-routed the F4 empty-grid fix into the platform repo for a full release**. M230 then
  built + runtime-proved the production-faithful fill (the `academy-fs-published-fallback` sha-pinned demopatch):
  the patched grid served 59 real skill-path cards, **0 Draft chips**, through the exact DB-authoritative path.
  The KB-fidelity prerequisite that unblocked everything after it.
- **Content stories (M231–M236).** M231 was the **hard go/no-go spike** — *GO* — resolving the central unknown
  by prove-by-render: the sim result page reads a **persisted row** (plain Ent SELECTs, no live recompute), so a
  cloned session renders. M232 built the `ContentStorySeeder` that **copies** real prod sessions (feedback,
  transcript, submission, interview report, real skill node-ids), best-effort PII-scrubbed and re-tenanted.
  M233 projected the honesty-gated `content-manifest.json`; M234 rendered the 2nd cockpit tab. M235 landed
  everything provable-without-a-browser (the 13-session matrix + 3 non-sim sections, 4 products / 18 sessions,
  both honesty gates green). M236 **re-proved it live on `billion`**: 29/29 landable (session × action) pairs,
  65 academy cards / 0 Draft chips, hero p95 click→ACCESS **1.22 s / 1.51 s** against a 5 s budget — all cold,
  no intervention.

## The defining thesis: *a check can report success while proving nothing*

Every release names one lesson. v2.5's is the most transferable this project has minted, because it is not a
bug — it is a **shape**, and it appeared at every altitude of the work.

It bit nine times across M235–M236 alone: a scrub that removed **zero** names with every test green; three
tests **asserting the very defect they should catch**; a `0/0` aggregator exiting green (0/0 is also
arithmetically 100%); an e2e suite passing by **collecting 0 tests** (silently taking 61 tests offline for 8
iters); a grader with **no negative tests at all**; and — the sharpest — a regression test for the manager-route
defect that was a **self-consistent tautology** (it asserted the projected path ended in `slot.MembershipID`,
produced by the same expression under test; *the fix for the defect contained a test that could not detect the
defect's return*).

**And then the release review found the class alive inside the close that named it.** The headline `29/29` had
three independent integrity problems: one of the 29 pairs was graded by the **wrong check** (`shapeFor` misgrades
a scored assessment whose slug contains "interview" down the no-floor `player-interview` branch — a live false
PASS in the shipped canonical manifest, §CQ-1); the four new `*.unit.spec.ts` files — **including the one that
carries the only literal `29`** — are executed by **no runner** (§CQ-2, *a pin nobody runs is not a pin*); and
the denominator is **self-derived** from the same served manifest the sweep reads (§CQ-6). Add the org-scrub arm
that has **never fired, fails open, and sits outside the leak gate** (§CQ-3), four more tautological tests
(§CQ-5), and two vacuous-pass tests (§CQ-7).

**Stated once, as the release's keeper:** ***ask of every layer that reports a number — what does it print when
nothing happened?*** A layer that cannot distinguish "everything passed" from "nothing ran" is not a gate.

The single most important consolidation insight is that this shape recurred at **three altitudes**, and only one
of them was folded into the thesis:

1. **The test/measurement layer** — the nine instances above. Named and now fenced (a collection-integrity
   guard with floors, a golden pin on the membership derivation, mutation verification on both).
2. **The build/decision layer** — **M232-P1** (below). The build-evaluation gate checked *severity* and found
   nothing wrong; it did not check whether the built thing was the thing the user asked for. A gate that grades
   the wrong property reports green.
3. **The review layer** — **DC-P5**: this very close template can issue a false green, demonstrated by v2.4's
   archived review asserting four false greens and pointing at a "Completeness Ledger below" **that the file does
   not contain**. A green review is the artifact downstream audits trust *most*, which is exactly why a false
   green there is the most expensive.

The thesis as filed scopes only altitude 1. It is the same class at all three.

## Incidents this cycle (P0 / P1 / notable P2)

| # | Sev | Milestone | What happened | Lesson |
|---|-----|-----------|---------------|--------|
| 1 | **P1** | M232 | **The build implemented the approach the user had explicitly REJECTED.** A safety-conscious sub-agent built "anonymize-by-construction" (synthesize free-text, never copy) — but the user had been explicit: *"literally copy the customer session, anonymize where possible."* Caught at the build-evaluation gate (not silently merged), surfaced as a fidelity choice, user chose copy-real; reworked the content layer, infrastructure survived. | **A safety-conscious agent will default to the safer design even against an explicit user decision — the orchestrator MUST diff the built approach against the stated instruction, not just check SEVERITY.** (This is altitude 2 of the release thesis.) |
| 2 | **P1** | M235 | **USER-BLOCKER-M235-01 — the scrub removed zero names.** 8 of 9 shipped M232 fixtures **leaked a real customer first name**, every test green. Root cause: the scrub map was built only from `jobsimulation.actors.username/.alias` (empty for these sessions), **never from the session owner's `public.users` identity** — where the candidate's first name actually lives, threaded through the LLM feedback. Fixed by owner-identity sourcing + token-split + word-boundary matching + a **capture-time fail-closed `SurvivingToken` post-condition** + a standing CI tripwire; re-captured 9 fixtures provably clean (0 leaked names, 545 placeholders). | **A status artifact shipped a claim the code did not enforce.** Bind the claim with an executable fail-closed gate, not prose. Surfaced in an iter-02 re-survey *before* extending the fixture with 4 more real sessions — surfacing-then-deciding, not building-then-discovering. |
| 3 | **P2** | M232 | **The weekly API limit terminated the rework mid-doc-edit.** Recovered cleanly: uncommitted work was intact in the working tree (verified before touching anything — no reset/checkout/clean/stash), the agent resumed from transcript once the limit unlocked, finished, committed, re-tagged. **No work lost.** | The forbidden-operations discipline (never `reset --hard`/`checkout --`/`clean`/`stash`) is what made a mid-flight kill a non-event. |
| 4 | **P2** | M236 | **The green gate failed OPEN west of UTC** (iter-09). The age check parsed a UTC `ts` as local time on BSD, so a 121-second-old verdict aged as 7321 s. West of UTC the same arithmetic reads a **stale** verdict as **fresh** — the exact hazard the check exists to prevent, inverted, for half the world. Fixed + regression-tested with a full timezone sweep and a shipped-line pin. | A safety check with a sign-flipped failure mode is worse than none — it manufactures confidence in the direction it was built to deny. |
| 5 | **P2** | M236 | **61 tests offline for 8 iters** — the whole e2e suite passed by collecting 0 tests after a module-scope throw. Found only in the final harden. | The collection layer can hide all the other false greens at once; it is now floor-guarded. |
| 6 | P3 | M236 | A **false PASS in the gate reading itself** (iter-07): a skill-path manager pair scored green off chrome served by a different query than the one that had failed. | A probe must not be constructed from the artifact under suspicion. |

Also worth recording as a not-an-incident: **M235-02** (a planning-assumption miss, not a defect) — TOK-01's
planned "coverage descriptor" mechanism **did not exist**; the exact-path crawl harness structurally cannot
reach the dynamic-URL, cockpit-seat-reached content-stories result pages. Resolved by building the non-sim
seeders and routing the new seat-login sweep plumbing to M236 — where it was **authored against a live render**,
and iter-04's calibration found **six shapes where the roadmap assumed one**, two of which a blind author would
have gotten wrong in opposite directions. M235-02 is *why* M236 authored the harness live rather than blind, and
that decision paid off directly.

## Cross-milestone patterns

**1. A check can report success while proving nothing.** (The thesis, above.) The release's single most reliable
predictor of a future defect.

**2. Prose does not propagate; only a shared definition or an executable fence does.** This recurred **inside the
close that named it**. The knob count did not propagate (`demo-up-defaults.md` says 27, `README.md` still says 25,
the skill still says 26 — DC-9). The correction sweep meant to enforce the rule (`302a32e`, "correct the claims
M236 refuted") touched **16 files, zero under `corpus/services/`** — leaving the whole service-doc tier asserting
refuted claims (D-1..D-6, one root cause). The non-integer-`N` guard was added to 2 of 4 runners; the membership
key was a bare literal at **9 sites** (one writes the row, eight hope to match it). Every instance is now a shared
definition or a fence. *The positive corollary M231 found:* a spike that **prove-by-renders naturally re-audits
the docs it reads** — but only the ones it reads; it does not sweep the siblings it doesn't touch.

**3. Offline-authored, never driven — *unit-proven ≠ route-proven* (DC-P3).** Four artifacts shipped wrong for
exactly this reason, each defended by a green test: the manager route built from a user id; `/library/<slug>`
which 404s; `managerKind: skill-paths` pointed at a "Coming soon" surface; the skill-path version `"2"` guess.
The class was never stated at the individual sites; it is now.

**4. The measurement-hygiene cluster (DC-P4).** Four instances of a measurement lying: the suppressed-stderr
`git fetch` measuring stale-vs-stale; the TZ age-check failing **open** west of UTC (incident #4); a denominator
counted from **survivors**; a probe hand-built from the artifact under suspicion. Individually landed,
collectively unrecorded until this close.

**5. The clone-staleness root cause — the generator of the entire "pin-drift" class.** `F-M236-CLOSE-1` /
`KB-A`: **`/demo-up` rebuilds images from clones it never updates.** Measured `app` **249** commits behind `main`
and `next-web-app` **202**, identically on both boxes; `clones.lock.json` records only local `{ref, sha}`, never
compared to the remote, so a 249-behind clone is **structurally indistinguishable from a fresh one** — and the
healthy-looking branch name is what misleads. **This was the user-reported stale left menu**, and the upstream
generator of everything that *looked* like pin drift this release. Anchored at close in a new `## Clone freshness`
section in `rosetta_demo.md`, cross-linked from `verification.md`'s pre-flight rung.

## The deferral story

**v2.4 shipped with no release-scope deferral audit at all** (`releases/archive/02.40-casting-call/` has four
milestone-scope audit dirs and none at release root). That is the structural cause of this close's headline: **8
items aged out unchecked, 7 of them newly detected here.** One (`DEF-M226-01`, the pre-bind serve reap) **aged
out twice** — M228 fired without it, then M236 fired without it — while `state.md` still named M228, stale for
two releases. The M204 **assign-WRITE** item has **~10 routings across 5 releases**, and *both* v2.5 milestone
audits recorded it as "correctly routed" — true of the *routing*, false of the *destination*.

All 22 items are now fated in [`release-deferrals.md`](release-deferrals.md), decided interactively by the
release owner 2026-07-20:

- **`M237 — re-prove-on-billion`** (a reservation, v2.6's declared first work) discharges **eight** items in one
  live bring-up: `CLOSE-D3` · the 39 live specs (`T-3`) · the anon-academy twin (`S-1`, = surviving `F4`) ·
  the `apps/web` `:5050` endpoint · `DEF-M226-01` · `BURNIN-M221-dev-public-host` · `F-M220-4` ·
  `PROBE-M218-c3-rerun`. That one-bring-up-discharges-eight is the whole argument for M237 being one milestone,
  not eight backlog rows.
- **`DEF-M226-01` was given teeth**: M237 must **TEST** the never-tested *"self-resolves in the default flow"*
  claim — the justification that let it pass three times untested — **or DROP it**. A justification that survives
  three passes untested has forfeited its claim on a fourth.
- **`M238 — playthrough-assign-write`** takes the assign-WRITE item with a **drop-expiry**: ten routings is the
  limit; if M238 isn't designed into v2.6/v2.7, it drops.

**Two failure classes were named, and they are the takeaway the next close must inherit:**

- ***"An arch-doc pass is not a milestone — a fate needs a milestone."*** (S-2 class.) A Fate-3 whose
  destination is a *kind of activity* ("a sweep", "the next close", "standing backlog") rather than a *unit of
  work* has no `In:` list to appear in and no close that can fire on it. It is indistinguishable from a drop
  except that nobody agreed to drop it. Three of the eight aged-out items failed exactly this way.
- ***"Confirmed-covered is the deferral form that leaves no ledger entry."*** (S-8 class.) A Fate-2 says
  "another milestone's gate covers this," so it is written as a **note**, never a carry, never a `DEF-` id. If the
  covering gate doesn't run — or runs and measures something weaker — **nothing anywhere notices**. A Fate-3 that
  fails leaves an aged-out trigger; a Fate-2 that fails leaves **silence**. Mitigation for the next close: a
  Fate-2 must name the covering gate's *assertion*, not just the milestone; and a Fate-2 whose covering milestone
  closes `closed-incomplete` **auto-reverts to an open Fate-3**.

And the meta-finding, **DC-P5**: v2.4's review is the proof that *the close template itself can issue a false
green*. It asserted "no undelivered Fate-3 items" while its own same-day retro listed two as inherited carries,
and pointed readers at a completeness ledger the file did not contain. **That dangling pointer is the concrete
mechanism** by which the standing test-debt carry and `DEF-M226-01` left v2.4 with no landing record. v2.5's
answer: a **real per-item ledger with a destination-still-valid check** (which is what `release-deferrals.md`
*is*), every fate naming a milestone, and the escape hatch written *where the number is*. Suggested fence for the
next close: refuse to emit a review containing a forward reference to a section heading absent from the file.

## What went well

- **Zero platform-repo edits held across all 8 milestones** — every platform-source wall routed to a sha-pinned
  demopatch or a config seam. The demopatch mechanism scaled to the interview flag-gates and the academy fill
  without a single canonical-repo touch.
- **Both M235 blockers surfaced BEFORE any wrong code landed.** The scrub leak was caught in a re-survey before
  the PII footprint was widened; the missing coverage mechanism was caught in planning before a mis-shaped
  descriptor was written. Surfacing-then-deciding kept the close clean.
- **The gate denominator was corrected rather than defended.** Driving the live surface showed 2 pairs were not
  provable; the target **shrank 31 → 29**, was argued inline with product-source evidence, and the 31 was
  **struck through**, not rewritten. It would have been easy — and wrong — to report 31/31.
- **The review listed what it got wrong.** `release-review.md`'s "Findings this review got WRONG" section (6
  refuted findings + 6 it missed, found during the fix pass) is the antidote to its own thesis: *a review that
  lists only its hits is the same false-green shape the release is named for.* KB-2 (ports) resolved harder than
  filed (roadrunner has **no caller at all**); D-5 (skill-path manager) was over-broad (the cohort scoreboard
  works; only the per-member drill-down is dead) — retracting the mirror guidance wholesale would have destroyed
  correct advice.
- **M229's cheap doc-fidelity milestone earned its place.** A false claim in one doc had mis-routed a real fix to
  the wrong repo for a release; the corpus-wide grep at close caught the *same* claim in three sibling docs.
  Correcting the map before building on it is what let M230–M236 land where they belonged.

## What hurt

- **The correction sweep violated the rule the release authored.** `302a32e` swept `corpus/ops/**` but never
  `corpus/services/**` — so the very pass meant to propagate M236's refuted-claim corrections left an entire doc
  tier asserting them. Prose does not propagate; a scoped-by-directory sweep is not a fence.
- **Two cross-repo doc-truth guards were RED and correct for three milestones, and were read as noise** — because
  each guard's *own test* hardcoded the superseded fact, so the suite was red for a reason that looked like the
  guard being broken. **A test asserting a stale fact is indistinguishable from a real regression**, and it trains
  readers to ignore a working alarm.
- **The standing carry was briefed as 14 and measured as 19** (asserted three ways: `state.md` 14, `roadmap.md`
  19, M236 14). Five failures had **no ledger entry anywhere**. A carry that is copied forward rather than
  re-measured stops being a fact. Re-baselined at close: **8 macOS / 7 Linux, 0 real defects, 0 pin drift** — the
  dirty-clone reading (716/14) was an artifact of the clone-staleness root cause, not real debt.
- **M236 could not cleanly self-close.** Its deferral audit returned RED and escalated the blocking 14-carry to
  the release (CLOSE-D2); the merge was HELD pending the user's release-close fates. The terminal milestone of a
  release is where the un-audited debt of the *prior* release comes due.

## Metrics delta

Sources: [`metrics.json`](metrics.json) vs [`../archive/02.40-casting-call/metrics.json`](../archive/02.40-casting-call/metrics.json).

| | v2.4 "casting call" | **v2.5 "the playbill"** | Δ / note |
|---|---|---|---|
| Go test funcs (`git grep '^func Test'`) | 1879 | **1976** | **+97** (like-for-like; 1879 is the reproducible v2.4 baseline, not the non-reproducible 1902 M233 quoted) |
| TS specs collected | 151 | **235** | **+84** — but **39 live-browser specs unexecuted** at close (24 stack-verify + 15 playthroughs; need a running stack) |
| Go suite | pass | **2461 / 0**, 6 modules | GREEN |
| Python suite | 644 / 14 (dirty-clone reading) | **1399 / 10** (8 standing debt + 2 flakes), re-baselined 8 macOS / 7 Linux, **0 real defects** | YELLOW; the "14" was a clone-staleness artifact |
| Seeders coverage | 96.8% | **96.1%** | −0.7pp, **within** the 2pp tolerance |
| p95 click→ACCESS | 1270 ms (recruiter) | **1220 / 1510 ms** (employee / manager) | flat, far inside the 5000 ms gate; **vantages differ — not strict like-for-like** |
| **Flakes** | 0 | **0 → 2 → 0** | 2 found at Phase 4b, **FIXED test-side at Phase 7, verified 3/3 serial-green — not waived, not `KEEP-REGRESSION`** |
| Supply chain | GREEN | **GREEN** | **0 net-new deps** (Go + both TS packages) |
| Platform-repo edits | 0 | **0** | held 8/8 |
| Alignment (Clerkenstein) | 100 / 100 | **100 / 100** | untouched |

## Stats delta (Phase 8c snapshot)

**Snapshot:** [`../../../journal/stats/2026-07-20.json`](../../../journal/stats/2026-07-20.json) · **prior:**
[`../../../journal/stats/2026-07-09.json`](../../../journal/stats/2026-07-09.json) (captured mid-v2.4, so this
window spans the v2.4 tail + all of v2.5 — not a clean per-release cut).

| | 2026-07-09 | **2026-07-20** | Δ |
|---|---|---|---|
| Total commits | 719 | **1017** | **+298** |
| Commits last 7d | 52 | **239** | +187 |
| Commits/day avg | 4.3 | **5.7** | +1.40 |
| Lines added | 105,938 | **150,382** | +44,444 |
| Lines removed | 22,084 | **27,773** | +5,689 |
| Churn % | 21 | **18** | −3 |
| Co-authored commits | 635 | **892** | +257 |
| Milestones done | 2 | **8** | +6 |
| Decisions recorded | 0 | **17** | +17 |

**Read the snapshot's zeros correctly:** the project-stats skill measures the **rosetta docs corpus**, whose
`code` and `doc-line` counters read **0** in every snapshot since 2026-07-01 — the release's actual code lives in
the git-ignored `.agentspace/rosetta-extensions`, and the doc-line counter has never resolved this repo's
`knowledge/` layout. The load-bearing signal from the snapshot is therefore **git velocity**, which is real:
**5.7 commits/day sustained, +298 commits, +44 K lines added.** The release's own code-size figure —
`rosetta-extensions 1d97861..main`, **34 commits / 94 files / +20,354 −108** — is recorded in
[`release-review.md`](release-review.md), not the stats snapshot, and is the number to quote for v2.5's build
size. *(Gap noted per the close protocol: the docs/code counters returning 0 is a pre-existing script
limitation, not a transient failure of this run; the snapshot wrote cleanly.)*

## The headline honesty note

**v2.5's primary metric — `29/29` landable (session × action) pairs — is UNIT-PROVEN, NOT LIVE-RE-PROVEN.** It
is tagged on that basis, as a **conscious user decision**, and this retro states it plainly rather than burying
it.

What happened, in order: the `29/29` reading was taken **live on `billion`**, cold reset-to-seed, at rext tag
`playbill-m236-hardened` — that measurement was real. The M236 close then fixed **~10 defects in that same
harness**; Phase 7 changed it **again**, and two of those changes bear directly on the number (CQ-1: one of the
29 was graded by the wrong check; CQ-2: the spec that pins the literal `29` is executed by no runner). **No live
re-run has happened since. The tag that actually shipped has never faced a browser.**

Fate 1 (land it now) **did not fail on capability** — `billion` is up and reachable, and v2.5 ran ten live iters
against it. It failed on an **explicit user choice**: tag v2.5 now, re-prove as v2.6's opening work. Recorded as
an **escape hatch, not an impossibility**, in [`release-deferrals.md`](release-deferrals.md) §A, with a concrete
acceptance condition for **M237**: a live `run-content-stories.sh` green **at the tag that actually shipped**,
with the CQ-1 grader fix and CQ-2 runner wiring in place, and `EXPECTED_PAIRS` sourced from something other than
the artifact under test (CQ-6). Anything less re-issues the same false green — which would be this release's own
thesis, one more time.

## What the next release must inherit

1. **The live re-prove is v2.6's first work, not a nice-to-have.** The headline number has never faced the
   shipped harness. M237's acceptance condition (§A) is the gate; treat it as a debt, not a formality.
2. **Every fate names a milestone; every "confirmed-covered" names the covering assertion.** The two named
   deferral classes (S-2, S-8) are what let 8 items age out. A per-item destination-still-valid ledger must be a
   standing close artifact — v2.4's absence of one is the whole reason this close's audit read RED.
3. **Prose does not propagate — fence it or it drifts.** The correction sweep that violated this rule while
   enforcing it is the proof. Doc-truth guards, knob-count fences, golden-value pins: only the executable ones
   held.
4. **Ask of every layer that reports a number: what does it print when nothing happened?** The collection-,
   aggregation-, grading-, coverage-, and regression-test layers all reported green on nothing this release. The
   floor guards landed at close are the down payment; extend them to any new grader.
5. **Fix the clone-staleness generator at its source.** `/demo-up` rebuilding images from un-updated clones
   (`app` 249 behind, `next-web-app` 202) is the upstream cause of the entire pin-drift-looking class, and of the
   user-reported stale menu. The anchor doc exists now; the mechanism fix does not.

## The one line worth keeping

**A check can report success while proving nothing — at the test, the build, and the review layer alike.** v2.5
named the class, fenced it at nine sites, and then found it alive in the close that named it, in v2.4's archived
review, and in its own `29/29`. What held was never the *name*; it was the *fence that executed the thing* — the
collection-integrity floor, the golden pin graded by mutation, and, above all, the cold reset-to-seed on the live
box that no unit spec can stand in for. The one the release has not yet run.

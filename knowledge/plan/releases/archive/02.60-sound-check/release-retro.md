# Release Retro — v2.6 "sound check" (M237 → M244)

**Shipped 2026-07-23.** The **reliability / field-hardening** release: *make everything that's built
actually get built + provisioned — and PROVE the whole thing live on `billion`, cold reset-to-seed.* 8
milestones (**M237** barrier **→ {M238 ∥ M239 ∥ M240→M241→M242 ∥ M243} → M244** the iterative
prove-on-billion closer), tooling + docs only, **0 platform-repo edits across all 8**. Consolidates the
eight milestone retros ([M237](m237-clean-stage/retro.md) · [M238](m238-ant-academy-reliability/retro.md) ·
[M239](m239-enterprise-surfaces/retro.md) · [M240](m240-content-stories-fidelity/retro.md) ·
[M241](m241-content-stories-language/retro.md) · [M242](m242-cockpit-ux/retro.md) ·
[M243](m243-assign-write-playthrough/retro.md) · [M244](m244-prove-on-billion/retro.md)), the release-level
[`release-review.md`](release-review.md), and the deferral audit
[`audit-deferrals/`](audit-deferrals/deferral-audit-2026-07-23-release-close.md).

## The headline

**v2.6 proved the whole release live on `billion`, cold, 0 platform edits.** M244's multi-part exit gate
(a–h) was discharged **8/8 GREEN** driving from a tailnet peer against a cold reset-to-seed demo on the
`billion` Tailscale VM: ORG-CLEAN 0 tokens · content-stories **47/47** of the 49-pair denominator (2 voice
player cells presence-only) · the **40 live-browser specs executed green** (24 stack-verify + 16/16
Playthroughs, 96 cases in one clean full run) · the anon academy twin · serve-reap 7→0 · the 3 v2.3
drift-carries · interview plan-section alignment · all 6 v2.6 fixes + **p95 click→ACCESS 1.46 s / 1.31 s**
against a 5 s budget. This discharged v2.5's headline debt — its `29/29` had shipped **unit-proven, not
live-re-proven** (the harness was fixed ~10× after the measurement); v2.6 re-proved it live at the grown
denominator.

## What shipped

The house shape was **barrier → parallel fixes → prove-on-billion**.

- **M237 clean stage (the HARD go/no-go barrier).** Fetch-verified clone-freshness in `ensure-clones.sh` (a
  real **7-state pin model** distinguishing deliberate-pin from stale-by-neglect) + the directory-driven R1
  pristine sweep (all 14 manifests, was a hard-coded 3). Its confirmed-defect ledger re-triaged the
  ambiguous v2.5 UI defects on a correctly-built demo — and **refuted the release's own premise**: billion's
  clones were 0–2 behind, not the "~202 behind" the design assumed; the "202" was itself the suppressed-fetch
  artifact §1 fixes.
- **The parallel fan-out (M238–M243).** M238 one FS-published chapter-body demopatch fixed **both** academy
  defects (Start→404 + the "Italian errors" symptom were the same dead-backend path). M239 **talk-to-data
  went FULL** — a real AWS Bedrock credential class for `app`, values-blind, bridged to the container,
  live-round-tripped ("Cervato Systems has 51 members"). M240 content-stories fidelity (selection/document/
  pass-rate) + **voice presence-only** as the faithful deliverable. M241 gave each session its **real played
  language** (11 of 13 were mislabeled `english`) + the cockpit **EN|IT** toggle (denominator 29→49). M242
  cockpit render/UX (tuple regroup + header tab + AA avatars). M243 landed the **FIRST MUTATING Playthrough**
  `pt-assignment-assign`.
- **M244 prove-on-billion (the iterative terminal closer).** 27 iters (24 tiks / 3 toks), gate MET 8/8.

## The defining thesis: *not-provisioned / wrong-version-built — caught only by proving it live*

Every release names one lesson. v2.6's is that **a build can be correct in source and still be wrong on the
box** — the thing that ships is not the thing that was built, because a dependency was never provisioned or
the wrong version got compiled — and **the only check that catches it is executing the real thing on a fresh
box.** The class bit three times, at three layers, and each time a live prove-on-billion was what surfaced
it:

1. **The clone layer (M237).** `/demo-up` rebuilt images from clones it never updated, and
   `clones.lock.json` recorded only local `{ref, sha}`, never compared to the remote — so a stale clone was
   **structurally indistinguishable from a fresh one**. The user-reported "stale menu" was this. Fixed with a
   fetch-verified freshness assertion + a 7-state pin model.
2. **The credential layer (M239).** Talk-to-data was *coded* but *inert* — the real AWS Bedrock creds were
   never provisioned to the container (the M217 override drops the `~/.aws` mount, so env vars are the only
   vehicle). "Does the feature work?" could only be answered by a live Bedrock round-trip, not a unit
   assertion.
3. **The build-pin layer (M244, iter-25).** `up-injected.sh` built the backend `:injected` image from the
   **highest fetched tag** (v1.351.0, which has `ai_readiness_cycles.launched_by`) instead of the source's
   **pinned checkout** (v1.341.0, the migrated schema) — so the binary SELECTed a column the schema never
   created, the cycles endpoint 500'd, and every ai-readiness surface rendered the zero-state, *looking* like
   a seed gap. Fixed durably (build-scratch resolves the pinned ref + an M217-style preflight), blended into
   `verification.md` rung zero.

**Stated once, as the release's keeper:** ***a green source tree is not a provisioned box. Only executing
the real artifact on a fresh box proves what actually got built and wired.*** This is precisely the
"reliability / field-hardening" charter working on itself — the release exists to catch this class, and it
kept catching it in its own tooling.

## Cross-milestone patterns

**1. The not-provisioned / wrong-version defect class** (above) — clone-staleness (M237), unprovisioned
Bedrock creds (M239), pinned-ref build-skew (M244). The single most transferable finding of the release:
each was invisible to source review and green tests, and visible only live.

**2. The anti-toothlessness thesis held — the hardens caught 3 toothless tests.** A release whose proof is a
browser harness lives or dies on whether an assertion can false-PASS. The two M244 harden sessions
mutation-verified every hardenable fix and found three checks that would pass on a broken page: a
year-less `dueDate` matcher that matched "24 hours"/"5 days" (would false-pass a page with no real
deadline), the iter-16 academy marker that false-passed an empty grid (moved to a card-count floor), and the
iter-08 route-shape scope test. M242 independently found + fixed 2 more toothless render tests (wrong-column
placement + unescaped title). Every one was mutation-verified RED-on-break before it was trusted — *a
regression test nobody has ever seen fail is a hypothesis, not a guard.*

**3. The coarse binary-per-gate metric is an artifact on an iterative proof-closer.** M244's primary metric
counts a gate part only when **fully** green (gate b at 47/47, gate c at 40/40), so it read flat over the
milestone's most productive tiks — the two triggered toks (**TOK-02, TOK-03**) both fired mechanically on
the metric floor, and the bootstrap TOK-01 opened the milestone. **All three HELD the strategy** (a
coarse-metric artifact, not a stall — no revision, no lost work). Recorded for future iterative-closer
design: a per-gate-part fractional metric would read progress more honestly.

**4. Honest triage shrank scope instead of manufacturing work.** M238's #2 "language bug" was investigated
and found to be the *same* dead-backend path as #3 plus two non-defects — the deliverable shrank to one
coherent patch. M239's #1 menu + #4 library were **no-defect verdicts recorded honestly** (with `file:line`
evidence + a live GREEN), not dressed up as fixes. The three-fate rule's discipline applied to triage.

## Incidents this cycle (notable P1 / P2)

| # | Sev | Milestone | What happened | Lesson |
|---|-----|-----------|---------------|--------|
| 1 | **P1** | M244 | **`launched_by` version-skew** — the `:injected` image compiled the highest fetched tag, not the pinned checkout → SELECT of a non-existent column → 500 → every ai-readiness surface zero-stated, *looking* like a seed gap. | The demo image must compile the **pinned** ref, not the highest fetched tag. Blended into `verification.md` rung zero. |
| 2 | **P1** | M244 | **Interview report rendered EMPTY on billion** (iter-04/05) — `directus.simulations_extraction` (the interview plan) was never captured → null plan → empty report. | A live gate caught a capture-surface gap no unit test could see; the plan-section-id alignment assertion then caught a second real drift. |
| 3 | **P2** | M239 | **A harden fix shipped a regression the harden didn't catch** — the F1 disk-probe fix omitted the `|| true` its own file documents (ISSUE-7), so a wedged Docker daemon aborted the whole bring-up before the fallback ran. Caught + fixed at close with a mutation-verified `set -e` regression test. | A new `docker run` in a `set -euo pipefail` script needs its **errexit path** tested, not just its happy path. |
| 4 | **P2** | M244 | **Seed-gap mischaracterization** (iter-26) — 3 ai-readiness "failures" were recorded as a plausible seed-gap story; iter-27 **falsified it** against the live DB: all three were harness locator mismatches vs a correctly-rendering v1.341.0 UI. | Check the locator against the actual rendered UI before hypothesizing a data gap. |
| 5 | P3 | M241 | **The build under-tested its own headline fix** — the whole milestone writes the *real* language, yet no build test asserted the seeded `sessions.language` column carried it (reverting to hard-coded `english` passed every Go suite). Closed at harden. | A fix's regression test belongs in the section that lands it, not the hardening pass. (Recurred M242/M243 — see below.) |

## What went well

- **Zero platform-repo edits held across all 8 milestones** — every platform-source wall routed to a
  sha-pinned demopatch (the M238 chapter-body, the interview flag-gates) or a config/secret seam (the M239
  Bedrock credential class was a *secret class + a bridge*, not a code change).
- **The gate bound the anti-toothlessness thesis.** Every M244 gate part was proven by *executing* the check
  live — ORG-CLEAN ran read-only first, the 40 specs actually executed, DEF-M226-01 was TESTED not assumed —
  and the two real defects (#1, #2 above) surfaced **only because the proof was live**.
- **Values-blind held end-to-end** on a real cloud credential (M239) — provisioned, bridged,
  live-round-tripped, even stripped-and-restored as a harden mutation — without any value reaching a log, a
  test, a commit, or reasoning. Customer media (M240) never entered agent context; the video exhibit is
  by-reference (no byte-port).
- **A 5-release carry discharged as Fate-1, not re-deferred.** `DEF-M235-03` / M204 assign-WRITE — the exact
  "a fate needs a MILESTONE" class the v2.5 close named — got a real milestone (M243) and landed in full: the
  FIRST mutating Playthrough, asserted by a **read-back FLIP** (assignable count drops by exactly one, or the
  poll times out RED). 16 live Playthroughs, 0 TODO.

## What hurt

- **The build repeatedly under-tested its own net-new logic** (M241 language write · M242 column split,
  twice · M243 `isOnAssignments` parity). A render/write change that introduces new *layout* or a new
  *predicate* needs tests that pin *which slot / which route*, not presence — the toothless default. Every
  instance was caught at harden or close, none shipped, but the pattern recurred across four milestones.
- **The standing demo-stack test debt rode the entire release** (8 fails, ridden ≥5 milestones), surfacing
  through `test_cockpit.py` on every milestone that touched it and costing each close a provenance
  re-confirmation. M244 is proof-only and structurally cannot fix them; the release close is the named
  expiry — resolved by user fate (below).
- **The coarse binary-per-gate metric under-reported** the closer's most productive tiks and fired 2
  mechanical toks (no lost work, but the metric reads a stall where there was progress).

## The deferral story

The terminal release-close deferral audit returned **RED** — not a "something broke" RED (0 real defects, 0
platform edits, 0 pin drift, flake 0), but the designed terminal gate firing on debt explicitly routed here:
1 CHRONIC repeat + 2 AGED_OUT. Resolved by the release owner's authoritative fate 2026-07-23, **"Tag now,
carry to v2.7"**:

- **DEF-M239-01 (ENOSPC loud-build-fail) → DROPPED.** The disk-full class is already caught by M239's
  pre-flight Docker-VM disk-measure; the loud-abort is redundant belt-and-braces, un-validatable without a
  real ENOSPC. Retired honestly rather than parked (M244 `decisions.md` D2).
- **The 8 standing demo-stack test failures + the stack-verify `run-unit.sh` roster nit → v2.7 "test-health"
  (KEEP-DEFERRED-WITH-SIGNOFF, a NAMED milestone).** 0 real product defects, host-dependent; the whole set →
  a dedicated test-health pass. Why Fate-1 declined: the user chose tag-now for a proof-complete release
  (gate 8/8 MET); these are non-defects best batched (M244 `decisions.md` D3; landed in `roadmap-vision.md`
  under the v2.6 → v2.7 carry).
- **Resolved in-release (left the ledger):** reap-17700 (Fate-1, M244 iter-10) · DEF-M240-01 real-video
  exhibit (dispositioned player-presence-only, iter-07) · DEF-M226-01 serve-reap (TESTED 7→0, iter-11) ·
  DEF-M235-03 assign-WRITE (Fate-1, M243).

## Metrics delta

Source: [`metrics.json`](metrics.json) vs [`../archive/02.50-the-playbill/metrics.json`](../archive/02.50-the-playbill/metrics.json).

| | v2.5 "the playbill" | **v2.6 "sound check"** | Δ / note |
|---|---|---|---|
| Go test funcs (`git grep '^func Test'`) | 1976 | **2010** | **+34** (like-for-like reproducible anchor) |
| TS unit specs (executed) | 196 | **257** | **+61** (playthroughs 85 + stack-verify 172) |
| Python demo-stack (passed) | 730 | **839** | **+109** (host-sensitive section; 8 standing fails, **0 real defects**) |
| Live-browser specs | 39 unexecuted | **40 EXECUTED GREEN live on `billion`** | the v2.5-deferred execution, discharged (24 stack-verify + 16/16 Playthroughs) |
| content-stories | 29/29 unit-proven | **47/47 landed live** of the 49-pair denom | 2 voice player cells presence-only |
| Touched coverage | — | scrub 100% · seeders 96.1% · contentsession 93.6% | flat / in-tolerance (largest drop −1.1pp < 2pp gate) |
| **Flake** | 0 | **0** | recorded 0 every milestone; final-harden 3/3 clean |
| p95 click→ACCESS (cold, billion) | 1.22 / 1.51 s | **1.46 / 1.31 s** (employee / manager) | far inside the 5 s gate |
| Supply chain | GREEN | **GREEN** | **0 net-new deps** (docs / render / tooling-test release) |
| Platform-repo edits | 0 | **0** | held 8/8 |
| Alignment (Clerkenstein) | 100 / 100 | **100 / 100** | untouched |

## What the next release must inherit

1. **v2.7 "test-health" is a named, owed milestone — not standing backlog.** The 8 standing demo-stack fails
   (a mechanical stale-assertion subset + a docker/live-gated residue) + the `run-unit.sh` roster nit
   (`content-denominator` + `run-discrete` unrostered) are batched there by explicit user fate. DEF-M239-01
   is DROPPED, not carried.
2. **A green source tree is not a provisioned box.** The release's keeper lesson: the only check that catches
   not-provisioned / wrong-version-built is executing the real artifact on a fresh box. Every new
   provisioning seam (a credential, a pinned build ref, a captured surface) needs a live prove-it rung, not a
   unit assertion.
3. **A net-new predicate / layout change is born with its full parity block.** The recurring "under-tested
   its own new logic" pattern (4 milestones) is cheap to prevent and expensive to catch retroactively.
4. **A per-gate-part fractional metric for iterative proof-closers.** M244's binary-per-gate metric fired 2
   mechanical toks on productive tiks; a fractional read would not.

## The one line worth keeping

**A green source tree is not a provisioned box — only executing the real artifact on a fresh box proves what
actually got built and wired.** v2.6 named the class, caught it three times in its own tooling
(clone-staleness, unprovisioned creds, a pinned-ref build skew), and then discharged v2.5's own headline
debt by doing exactly that: the whole feature, cold reset-to-seed, live on `billion`, 8/8, 0 platform edits.

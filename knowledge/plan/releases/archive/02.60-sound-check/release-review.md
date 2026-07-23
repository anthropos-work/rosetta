# Release Review: v2.6 "sound check"

**Date:** 2026-07-23
**Milestones:** M237, M238, M239, M240, M241, M242, M243, M244 (8, all closed — M244 the iterative
terminal closer, closed-on-gate 8/8)
**Release diff:** rosetta `main..release/02.60-sound-check` — ~102 commits / 165 files (all markdown +
plan `metrics.json` artifacts; **0 platform-repo files, 0 non-doc code in the corpus**). rext
code-of-record advanced to `498b1a5` (final consumption tag `sound-check-m244-content-sweep-robustness`,
on origin).

## Gate verdicts

| Phase | Gate | Verdict |
|---|---|---|
| 0 — supply chain | — | **GREEN** — 0 net-new deps (docs / render / tooling-test release); per-milestone audits stand; corpus has no lockfile to audit |
| 1 — release scope | — | **GREEN** — every in-release Fate-3 routing accounted for; `DEF-M235-03`/M204 assign-WRITE LANDED at M243 (Fate-1) |
| 1b — deferral re-audit | blocking | **RED → RESOLVED-BY-USER-DECISION** — 4 items, 2 resolved in-release, 2 open (CHRONIC + AGED_OUT) → **user fate 2026-07-23 "tag now, carry to v2.7"** (see Deferrals) |
| 2 — code quality | — | **GREEN** — 0 must-fix; toolchain clean (all rext code lives + audited in rosetta-extensions, unchanged at close) |
| 3 — documentation | — | 0 must-fix · **2 should-fix (FIXED, Phase 7)** — 1 `coverage-protocol` 29-vs-49 contradiction + 1 `CLAUDE.md` demo-index gap |
| 3b — KB consolidation | blocking | **GREEN** — no structural-debt blockers (the v2.5-close KB split + service index already landed) |
| 4 — tests | — | Go **GREEN** (2010 funcs / 0 failed) · TS unit **GREEN** (257 / 0) · Python demo-stack **GREEN** (839 pass / 8 standing, 0 real defects) · **40 live-browser specs executed GREEN live on `billion`** |
| 4b — metrics regression | blocking | **GREEN** — Go +34, TS +61, python passed-count +109, flake 0, touched coverage flat/in-tolerance; 2 YELLOW annotations (standing-8, roster-drift), neither trips a hard condition |
| 5 — decisions | — | **GREEN** — per-milestone decisions blended; 3 close-release decisions recorded (D1 re-audit at M244 close + D2/D3 the user fates) |

**No gate is RED at close.** The single blocking gate that flagged RED — the terminal deferral re-audit
(1b) — is **resolved by an explicit user fate-decision** ("tag now, carry to v2.7"), not left open. The
release is GREEN on every hard condition: **0 real defects, 0 platform-repo edits, 0 pin drift, flake 0.**

## The release's own thesis, recurring — and holding

v2.6 "sound check" is the **reliability / field-hardening** release: *make everything that's built actually
get built + provisioned, proven live on `billion`.* Its central adversarial question — for a release whose
proof is a live browser harness — is **could a check report success while proving nothing?** (the v2.5
thesis it inherited). This review confirms the thesis **held**: every M244 gate part was discharged by
*executing* the check live (ORG-CLEAN ran read-only first; the 40 specs actually executed; DEF-M226-01 was
TESTED not assumed), and the two live-only defects the proof surfaced — the interview-plan capture gap
(iter-04/05) and the `launched_by` pinned-ref build skew (iter-25) — **could only have been caught live**.
The two harden sessions mutation-verified every hardenable fix bites; three toothless tests were found and
given teeth. No un-mitigated false-pass scenario remains (M244 `decisions.md` § Adversarial review).

---

## Scope

- [x] **S-1** `DEF-M235-03` / M204 **assign-WRITE** — the sole in-manifest Playthrough TODO, ridden 5
      releases — **LANDED as `pt-assignment-assign` at M243 (Fate-1)**: a manager assigns a skill path, a
      real `organization_assignments` row is written + read back (assignable count drops by exactly one).
      Corpus stands at **16 live Playthroughs, 0 TODO**. A 5-release carry leaving the ledger by landing is
      the deferral audit's success story.
- [x] **S-2** The reserved milestone numbers resolved cleanly: reserved `M237` (re-prove) → realized as the
      iterative closer **M244**; reserved `M238` (assign-WRITE) → realized as **M243**. `roadmap-vision.md`'s
      v2.5→v2.6 carry section records the remap.

## Code Quality

- [x] **CQ-1** No must-fix findings. All tooling code (rext `stack-seeding` / `stack-verify` / `demo-stack`
      / `playthroughs`) was reviewed per-milestone and is **unchanged at release close** — v2.6 added no
      net-new third-party dependency. The one code-quality-adjacent finding is the stack-verify roster drift,
      dispositioned under Tests / Deferrals below.

## Documentation (Phase 7 fixes)

- [x] **D-1 [should-fix, FIXED]** `corpus/ops/demo/coverage-protocol.md` L987 narrated a present-tense
      "landable count is **still 29**" describing the M236 route-contract test, contradicting the reconciled
      **49** at L917–919 (and `content-stories-spec.md:157` "29 → 49"). **Historically fenced** ("29 at M236
      iter-08; reconciled to 49 since M241's EN/IT growth — see the denominator note above") so the M236
      narrative stays accurate without reading as a live contradiction.
- [x] **D-2 [should-fix, FIXED]** `CLAUDE.md` "Demo Environments" doc index: (a) the M240
      `media-substrate-spec.md` (raw-media / voice-video substrate) was **omitted** while every other demo
      doc is listed — **added** with a full bullet (the Bunny-CDN-reference recorded-call facet + the inline
      document-body facet + the PII discipline); (b) the four content-stories-family entries
      (`coverage-protocol` "29/29", `content-stories-routes`, `session-clone-spec`, `content-stories-spec`)
      **terminated at v2.5 M236** — each got a tight v2.6 clause (denominator 29→49, 47 landed / 2 voice
      player cells presence-only, the EN/IT toggle M241, media-fidelity M240).

## Tests & Benchmarks

**Counts are from parsed runner output, never guessed. All numbers are the release-close reading in
`metrics.json`.**

| Stack | Result |
|---|---|
| Go | **GREEN** — 2010 reproducible test funcs (v2.5 close 1976, **+34**); 2461 testcases / 0 failed, 6 modules |
| TypeScript (unit) | **GREEN** — 257 `*.unit.spec.ts` (v2.5 close 196, **+61**); playthroughs 85 + stack-verify 172 |
| Python (demo-stack) | **GREEN** — 839 pass (v2.5 close 730, **+109**) / **8 standing fails, 0 real defects** |
| Live-browser (Playwright) | **GREEN** — **40 specs EXECUTED live on `billion`** (24 stack-verify + 16/16 Playthroughs, 96 cases, cold reset-to-seed) — the v2.5-deferred execution, discharged |

- [~] **T-1 [should-fix → DEFERRED, v2.7]** **8 standing demo-stack Python failures** — pre-M244 inherited
      (rext `04babf8`, 2026-07-15), **0 real product defects**, host-dependent: 6 `test_cockpit`
      academy/overlay (stale academy-shape / removed-window assertions vs deliberately-changed
      M218/M238/M220 behaviour) + `test_public_host` (port-13001 hiring drift) + `test_purge`
      (docker-integration). Chronic (ridden ≥5 v2.6 milestones). **User fate 2026-07-23: KEEP-DEFERRED →
      v2.7 test-health** (see Deferrals).
- [~] **T-2 [should-fix → DEFERRED, v2.7]** **stack-verify `run-unit.sh` roster drift** (net-new at the
      final tag) — the canonical runner's `UNIT_SPECS` roster lists 7 of 9 `*.unit.spec.ts`; the 2 the M244
      final harden added (`content-denominator`, `run-discrete`) were never rostered, so the runner exits 2
      and its `test_e2e_collection_integrity::UnitSpecsAreExecuted` guard is RED. **All 9 specs pass when
      run directly** (172) — the recorded count is accurate; only the rostered runner + guard are out of
      sync. A 2-line rext edit, **deferred to v2.7 test-health** with the standing-8 (same test-health
      class; not touched now — the rext repo is not modified in this close).
- [x] **T-3** Flake gate **0** across every milestone + M244 final-harden (3/3 clean on all touched tests).
      Touched-package coverage flat / in-tolerance (scrub 100%, seeders 96.1%, contentsession 93.6% =
      −1.1pp, within the 2pp tolerance).

## Deferrals (Phase 1b re-audit + user fates)

The terminal release-close deferral audit
(`audit-deferrals/deferral-audit-2026-07-23-release-close.md`) returned **RED** — not a "something broke"
RED (0 real defects), but the designed terminal gate firing on debt explicitly routed here: 1 CHRONIC
repeat + 2 AGED_OUT. The user's authoritative fate 2026-07-23 was **"Tag now, carry to v2.7"**:

- [x] **DEF-M239-01 — DROPPED** (M244 `decisions.md` D2). "Make the demo build fail loudly on ENOSPC." The
      disk-full class is **already caught** by M239's pre-flight Docker-VM disk-measure; the loud-abort is
      redundant belt-and-braces, un-validatable without a real ENOSPC. Retired honestly rather than parked.
      **Not carried.**
- [~] **STANDING-8 + roster-nit — KEEP-DEFERRED-WITH-SIGNOFF → v2.7 "test-health"** (M244 `decisions.md`
      D3). Whole set → a **named** v2.7 milestone (the fallback path: no late-merge branch opened). Why
      Fate-1 declined: the user chose **tag-now for a proof-complete release** (gate 8/8 MET); these are
      **non-defects** best batched into a dedicated test-health pass. Landed in `roadmap-vision.md` under the
      v2.6 → v2.7 carry.
- [x] **Resolved in-release (leave the ledger):** reap-17700 / M239-D13 (Fate-1 LANDED, M244 iter-10);
      DEF-M240-01 real-video exhibit (dispositioned player-presence-only, M244 iter-07, denominator 49→47);
      DEF-M226-01 pre-bind serve reap (TESTED 7→0, iter-11); DEF-M235-03 assign-WRITE (Fate-1, M243).

## Decision Consolidation

- [x] **DC-1** All per-milestone decisions were blended into the corpus at each close (each milestone's
      `decisions.md` + retro). Three release-level decisions are recorded at this close: **D1** (M244 close
      re-audit re-fating the 2 inherited carries Fate-3 → close-release) and the two user fates **D2**
      (DROP DEF-M239-01) + **D3** (RELEASE-SCOPE-DEFER standing-8 + roster-nit → v2.7).
- [x] **DC-2** The coarse **binary-per-gate primary metric** fired the tok-trigger 3× on M244's most
      productive tiks (a gate part only counts fully-green). TOK-01/02/03 all HELD the strategy (a
      coarse-metric artifact, not a stall). Recorded as a nice-to-have process note (below), no action this
      release.

### Nice-to-have (recorded, not actioned)

- [x] **NTH-1** **A per-gate-part fractional metric would read progress more honestly on an iterative
      proof-closer.** M244's binary-per-gate metric under-reported over three productive tiks and fired 3
      mechanical toks with no lost work. Captured for future iterative-closer design (M244 retro "What
      Didn't"); not a defect, no change this release.

---

## Phase 7 disposition summary

**Legend:** `[x]` landed / resolved in this release · `[~]` routed forward with a recorded fate + named
destination (v2.7 "test-health").

**Landed this close:** the 2 should-fix documentation reconciliations (D-1 coverage-protocol fence + D-2
CLAUDE.md demo-index media-substrate bullet + v2.6 clauses), all in `rosetta` corpus; **0 code fixes, 0
rext edits, 0 platform edits.**

**Routed to v2.7 "test-health" (user fate 2026-07-23 "tag now, carry to v2.7"):** the 8 standing
demo-stack test failures + the stack-verify `run-unit.sh` roster nit (same test-health class). DEF-M239-01
DROPPED (not carried).

**All review items are resolved** — either landed (Phase 7), dispositioned by user fate (deferrals), or
recorded as a non-actioned process note. Nothing is left open blocking the tag; the orchestrator's Phase 11
(release→main merge + `v2.6` tag) can proceed once the two fates above are recorded (they are — D2/D3).

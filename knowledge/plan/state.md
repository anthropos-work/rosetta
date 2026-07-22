---
active_release: "v2.6 «sound check» — the reliability / field-hardening release (designed 2026-07-20): make everything that's built actually get built + provisioned. Barrier→parallel-fixes→prove-on-billion, 8 milestones M237→M244. Branch release/02.60-sound-check; tag will be v2.6."
active_branch: "release/02.60-sound-check"
active_milestone: "M242 ∥ M243 — the post-barrier fix fan-out (M237–M241 CLOSED); M242 cockpit-UX (serial after M241) ∥ M243 assign-WRITE"
last_closed: "M241 — 2026-07-22 (content-stories language)"
phase: "v2.6 in development — M237–M241 CLOSED. M241 content-stories language: the EN/IT axis — real per-session language (cs.Language, was hard-coded english — the core bug), the cockpit EN|IT toggle (11/12 tuples bilingual, INTERVIEW Italian-only per R2), a fail-closed ValidateLanguageConsistency gate, fixture 13→23, denominator 29→49; 3 harden passes closed the core-bug write-side gap. Next: M242 cockpit-UX (serial after M241) ∥ M243 assign-WRITE; M244 owns the standing-9 + DEF-M239-01/M240-01."
last_updated: "2026-07-22"
---

# State

**v2.6 "sound check" — IN DEVELOPMENT.** The **reliability / field-hardening release** (the v1.3b / v1.10b / v2.1 /
v2.3 lineage), designed 2026-07-20, triggered by **live demo defects** — *"still not all gets built and provisioned as
expected."* The job: make everything that's *built* actually *build + provision* on a fresh box. House shape **barrier →
parallel fixes → prove-on-billion**. **8 milestones M237 → M244**, branch `release/02.60-sound-check` (cut from **local**
`main`); tag will be **`v2.6`**. **Tooling + docs only — zero platform-repo edits.**

```
M237 clean stage (HARD go/no-go barrier)
  ├─▶ M238 academy reliability ─────┐
  ├─▶ M239 enterprise surfaces ─────┤
  ├─▶ M240 content-fidelity ─┐      │   (HARD media-safety gate)
  │      └─▶ M241 language ─┐ │      │
  │            └─▶ M242 cockpit-UX
  ├─▶ M243 assign-WRITE ────────────┤   [realizes reserved M238]
  └─────────────────────────────────▶ M244 prove-on-billion (iterative closer) [realizes reserved M237]
```

- **M237 clean stage** [`section`, HARD go/no-go] — ✅ **CLOSED 2026-07-21.** Fetch-verified clone-freshness +
  7-state pin model + R1-all-14-manifests, dogfooded on `billion`; "202-behind" REFUTED. (Detail: roadmap.md M237.)
- **M238 academy reliability** [`section`] — ✅ **CLOSED 2026-07-21.** ONE chapter-body FS-published demopatch fixed
  **both #3 (Start→404) and #2 (language)** live on `billion`; academy sweep extended. (Detail: roadmap.md M238.)
- **M239 enterprise surfaces** [`section`] — ✅ **CLOSED 2026-07-21.** talk-to-data **FULL** — a real AWS Bedrock cred
  class (values-blind, R3) bridged into the demo backend, **proven live** ("51 members"); the close fixed 2 own-code
  defects. DEF-M239-01 + a 9th reap 17700 → M244. (Detail: roadmap.md M239.)
- **M240 content-stories fidelity** [`section`, HARD media-safety gate] — ✅ **CLOSED 2026-07-22.** 5 fixes (selection
  type-match · document `input_data` · pass-rate 70–95 band · media-substrate + §3.8.1 VIDEO) + **voice presence-only**
  (the faithful `not_available` IS the deliverable); the real-video exhibit → M244 (`DEF-M240-01`). (Detail: roadmap.md M240.)
- **M241 content-stories language** [`section`, pool-count go/no-go] — ✅ **CLOSED 2026-07-22.** The EN/IT axis: real
  per-session language (`cs.Language`, was hard-coded `english` — the core bug the pool query exposed: 11 of 13 pins were
  Italian) + the cockpit **EN|IT toggle** (11/12 tuples bilingual; INTERVIEW Italian-only per R2) + a fail-closed
  `ValidateLanguageConsistency` gate. Fixture 13→23, denominator 29→49. **3 harden passes** closed the core-bug
  write-side test gap (`TestContentStorySeeder_WritesRealLanguage`). rext tag `sound-check-m241…` @ `17beede`.
- **M242 cockpit-UX ∥ M243 assign-WRITE Playthrough** — the remaining fan-out (M242 serial after M241 → wants the
  pass/fail + language variants; M243 ∥).
- **M244 prove-on-billion** [`iterative`, terminal] — re-prove v2.5's headline `29/29` AND every v2.6 fix live on
  `billion`, cold reset-to-seed. Multi-part exit gate (a–h), opens with the read-only `ORG-CLEAN` settling check.

**3 binding user decisions (2026-07-20):** **(1) talk-to-data → FULL** — real AWS Bedrock creds via `/stack-secrets` +
a secret-coverage DNA extension for `app` (not just a flag) [M239]; **(2) media → PORT IT** — capture + re-host the
Chime/S3 voice recording + document blobs behind a **HARD internal PII gate** (fresh sign-off + a `safety.md` raw-media
amendment + a voice/document anonymization decision — a voice cannot be token-scrubbed) [M240]; **(3) language → EN-only
fallback per tuple** — M241 opened with a read-only IT-session pool-count go/no-go query [M241].

**Next:** the **post-barrier fix fan-out** continues (M237–M241 done) — **M242 cockpit-UX** (serial after M241) **∥ M243
assign-WRITE Playthrough**. Each fix is scoped against the fresh, correctly-built demo M237 established. The v2.5 headline
live re-prove stays reserved for the terminal **M244** (which opens with the read-only `ORG-CLEAN` settling check,
standing backlog below) — M244 also owns **DEF-M240-01** (the content-stories real-video exhibit, Fate-3, user
pre-approved), the **standing 9 demo-stack test failures** (M238-D5 / M239-D13 reap-17700), and **DEF-M239-01**
("fail the BUILD loudly on ENOSPC", M239-D12).

## ⚠️ The v2.5 headline shipped UNIT-PROVEN, not LIVE-RE-PROVEN — v2.6/M244 re-proves it

**`29/29` is UNIT-PROVEN, NOT LIVE-RE-PROVEN.** It was measured live on `billion` at `playbill-m236-hardened`; the close
then fixed **~10 defects in that same harness** (incl. a grader that mis-shapes one of the 29 pairs + four unit-spec files
no runner executes — among them the only literal pin of `29`). The measuring instrument changed after the measurement.
Also unverified: **39 live-browser specs** (24 stack-verify + 15 playthroughs). **This is v2.6's `M244` job** (the realized
reserved `M237`): `billion` is up + reachable. Ledger: `CLOSE-D3` + `T-3`.

## Recently closed (milestones, newest first — max 5)

- **M241 content-stories language** — 2026-07-22 (section, pool-count go/no-go; fourth post-barrier fix). The EN/IT
  language axis: real per-session `language` (`cs.Language`, was hard-coded `english` — the core defect the pool query
  exposed, 11 of 13 pins actually Italian) + the cockpit **EN|IT toggle** (11/12 tuples bilingual; INTERVIEW Italian-only
  per R2) + a fail-closed `ValidateLanguageConsistency` gate (with teeth) + the TS mirror. Fixture 13→23, denominator
  29→49. **3 harden passes** mutation-verified every gate + closed the **core-bug write-side gap** (no test asserted the
  seeded `sessions.language` — reverting to `english` passed every Go suite; `TestContentStorySeeder_WritesRealLanguage`
  closes it). rext tag `sound-check-m241…` @ `17beede`. Go 2005 funcs (+6); Py 142 (136/6 pre-existing); TS 151; flake
  5/5; deferral audit YELLOW (0 new); 0 platform edits.
- **M240 content-stories fidelity** — 2026-07-22 (section, HARD media-safety gate; third post-barrier fix). 5 fixes
  (selection `d.type` type-match · document `input_data` at seed [inline text, no S3 blob] · pass-rate 70–95 band · the
  media-substrate + `safety.md` §3.8.1 VIDEO Delivers) + **voice presence-only** — the faithful
  `chime_status='not_available'` IS the deliverable (user decision 2026-07-22); the real-video exhibit is
  Bunny-recording-key-blocked → M244 (`DEF-M240-01`, Fate-3, user pre-approved). rext tag `sound-check-m240…` @
  `ae0e869`. Harden mutation-verified 3 fixes + 6 tests, flake 5/5; PII held (by-reference video); 0 platform edits.
- **M239 enterprise surfaces** — 2026-07-21 (section; second post-barrier fix). talk-to-data **FULL** — a real AWS
  Bedrock cred class (values-blind, R3 not-critical) provisioned + bridged into the demo backend, **proven live** (a
  manager asked member count → "Cervato Systems has 51 members" via a real ~7 s Bedrock round-trip). #4 library + #1 menu
  were **no-defect verdicts**; F1 disk pre-flight landed (Docker VM disk, not host `/`). The close fixed **2 own-code
  defects** (disk probe + bridge append), mutation-verified. Deferral audit YELLOW: DEF-M239-01 + a 9th reap 17700
  collision → Fate-3 → M244. Py 106, Go secretdna PASS, flake 5/5.
- **M238 ant-academy reliability** — 2026-07-21 (section; first post-barrier fix). ONE chapter-body FS-published
  demopatch fixed **both #3 (Start→404) and #2 (language)** (#2 not a distinct bug — `?lang=`) live on `billion`;
  academy sweep extended + a demopatch-inventory fence. Deferral audit YELLOW (standing-8 → M244, D5). 0 platform edits.
- **M237 clean stage** — 2026-07-21 (section, HARD go/no-go barrier; **opens v2.6**). Fetch-verified clone-freshness +
  7-state pin model + R1-all-14-manifests, dogfooded on `billion`; **"202-behind" REFUTED** (clones 0-2 behind). #1
  RESOLVED, #4→M239, #2→M238. 160 tests, 0 platform edits.

## Recently shipped (releases, newest first — max 3)

- **v2.5 "the playbill"** — 2026-07-20 (tag `v2.5`). The **content-vantage** release: **29/29** landable
  (session × action) pairs live on `billion` both vantages + the academy grid filled (65 cards), **0 platform edits**
  across 8 milestones. ⚠️ **The 29/29 is UNIT-PROVEN, NOT LIVE-RE-PROVEN** — the harness was fixed ~10× after the
  measurement; the live re-prove is v2.6's `M244` (realized reserved `M237`).
- **v2.4 "casting call"** — 2026-07-18 (tag `v2.4`). The **recruiter-vantage / hiring-org** release: a 4th hiring org
  on the cockpit (45 candidates on 5 shared positions), proven live on `billion` (M228 7/7, recruiter p95 1.27 s).
- **v2.3 "cue to cue"** — 2026-07-15 (tag `v2.3`). The **presenter-speed** release: click→ACCESS < 5 s proven live
  8/8 on `billion` — login p95 2.11 s / 1.31 s vs a ~39/38 s baseline. Remote default-on.

## Headline numbers (v2.5, as measured at close 2026-07-20 — the v2.6 baseline; re-measured at the M244/release close)
- **Go:** **1976** test funcs at the v2.5 baseline → **2005** now (release-cumulative M237–M241). **2461 testcases / 0
  failed** across 6 modules at baseline.
- **Python:** **1409 testcases** — 1399 pass / 8 deterministic / 2 flake / 8 skip (v2.5 baseline). demo-stack full suite
  at M241 close: **808 pass / 9 fail** (the standing set, 0 new from M241 → M244).
- **TypeScript/Playwright:** **196 unit specs / 0 failed** at baseline; **39 live-browser specs (24 + 15) NOT executed**
  (v2.6/M244 executes them).
- **p95 click→ACCESS:** employee **1.22 s** · manager **1.51 s** (M236, COLD on `billion`), vs the < 5000 ms gate.
- **Flake:** **0.** **Alignment (Clerkenstein):** **100% / 100% critical**. **Platform-repo edits:** **0**. **Supply
  chain:** GREEN — 0 net-new deps (v2.6 M239 adds a Bedrock **secret class**, not a dep).

## D17 — the carried-forward signature hazard
**D17: *a status artifact that outlives the thing it describes, and is then read as evidence.*** **The keeper:**
***a named hazard is not a fence; only an executable probe binds.*** **v2.5's sibling thesis:** ***a check can report
success while proving nothing.*** v2.6/M244's gate is built to bind it: `ORG-CLEAN` runs FIRST, the 39 specs must
actually execute, `DEF-M226-01` must be TESTED or DROPPED.

## Branch model / shipped tags
**v2.6** branch `release/02.60-sound-check` cut from **local** `main` 2026-07-20. ⚠️ **v2.5's `release → main` merge +
`v2.5` tag are LOCAL-ONLY** — `main` + tag not yet pushed to origin (R5; flag to user, do not auto-push). rext
code-of-record at v2.5 close **`playbill-m236-close-fixes`** (on origin). The `billion` demo LEFT UP. **Shipped tags:**
**v2.5** · **v2.4** · **v2.3** · **v2.2** · **v2.1** · **v2.0** · **v1.10b** `v1.10.1` · … · **v1.0**.
(Detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

## Standing backlog (unscheduled, cross-release)

> **Every item has a FATED destination.** Full per-item reasoning:
> [`releases/archive/02.50-the-playbill/release-deferrals.md`](releases/archive/02.50-the-playbill/release-deferrals.md).
> **Class named at v2.5 close: a fate needs a MILESTONE — not "a sweep", "the next close", or "standing backlog".**
> **v2.6 remap:** reserved `M237` re-prove → **`M244`**; reserved `M238` assign-WRITE → **`M243`**.

- **Standing demo-stack test debt — RE-BASELINED.** **9 fails at the M241 full-suite close** (6 academy+overlay
  `test_cockpit.py` + `test_host_prereqs_m215` + `test_purge` + `test_reap` reap-17700), **0 real defects**, **0 pin
  drift**, **0 new from M240/M241**. **Host-dependent — always state the host.** → `rebaseline-standing-failures.md`.
  Fate-2 → **M244** (the named expiry; discharge by editing the tests — 6 of 9 need no live stack). **Ridden ≥4
  v2.6 milestones — M244 is the expiry point.**
- **→ `M244` (v2.6's live closer).** Opens with the read-only **`ORG-CLEAN`** settling check (13 copied session exhibits
  — resolve source-org names via one prod query + an offline `scrub.OrgTokens`/`SurvivingToken` pass; 0 surviving tokens
  or each dispositioned; VPN-scoped + data-controller-accepted). M244 then discharges: `CLOSE-D3` (29/29 live re-prove) ·
  the **39 unexecuted live-browser specs** · the anon `/library`+`/free` twin · the `:5050` non-offset client endpoint ·
  **`DEF-M226-01`** (pre-bind serve reap — AGED OUT TWICE; **TEST or DROP**) · **`BURNIN-M221` / `F-M220-4` /
  `PROBE-M218-c3`** (the v2.3 `DRIFT_DEFER` tail) · the M232 interview plan-section alignment assertion ·
  **DEF-M239-01** (ENOSPC loud-build-fail) · **DEF-M240-01** (the real-video exhibit, user pre-approved).
- **→ `M243` (assign-WRITE):** **`DEF-M235-03`** / M204 **assign-WRITE** in-manifest TODO — ~10 routings across 5
  releases; **fresh dated KEEP-DEFERRED-WITH-SIGNOFF 2026-07-20**. **Expiry: if `M243` does not land it, DROP.**
- **→ next `stack-seeding` build-iter:** **`DEF-M215-03(a)` / `F11`** — seed hero identity-key vs generated profile
  display-name mismatch (cosmetic). Enumerated by id so it stays findable.
- **Older, still unscheduled (re-confirmed 2026-07-20):** **DEF-M10-01** (S3 blob bytes — **CONSUMED by M240** for the
  document facet [inline text, no blob]; the voice/video facet is DEF-M240-01 → M244), DEF-M21-01 (`replayCmd` hermetic
  test), CAVEAT-1 (clean-box literal full `/dev-up`), M314b (prod frozen-read hydration), **M205**-residual (tier gates +
  ATS). Playthroughs futures **M206–M207** stay in vision.

_Last updated: 2026-07-22 — v2.6 "sound check" IN DEVELOPMENT (branch `release/02.60-sound-check`); **M237–M241
CLOSED**, the M242 ∥ M243 fix fan-out continues (M242 serial after M241)._

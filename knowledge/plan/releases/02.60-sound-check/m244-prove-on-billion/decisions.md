# M244 — Decisions

_(implementation choices with rationale accumulate here during the iter loop)_

## TOK-01: staged cold billion bring-up → gate-parts a–h one-cluster-per-tik — 2026-07-22

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Stage a cold, green, reset-to-seed demo on `billion` at the m243 rext pin, then discharge the exit-gate parts (a–h) in **dependency order, one cluster per tik**, driving and asserting from a **tailnet PEER** (this workstation), recording live evidence. Tik order: (1) Foundation — pre-flight rung zero on billion + gate (a) ORG-CLEAN read-only FIRST + cold reset-to-seed bring-up → fresh green `autoverify.json`; (2) gate (h) latency p95 < 5 s + v2.6-fix smoke; (3) gate (b) content-stories sweep 49/49; (4) gate (c) the 40 live-browser specs; (later) gates (d) academy twin, (e) DEF-M226-01 serve-reap test-or-DROP, (f) 3 drift-carries, (g) interview alignment assertion, plus inherited DEF-M239-01 / reap-17700 / DEF-M240-01.
**Rationale:** the live-proof lineage (M215/M221/M226/M228/M236) makes the cold billion bring-up the critical path and highest-risk step (F1–F12 host findings, the login-shell PATH trap F2b, the M217 pin trap, the peer-vs-VM loopback-TLS trap). Every downstream proof (sweeps, playthroughs, latency) **gates on a fresh green `autoverify.json`**, so a green cold bring-up is the enabling precondition — it is tik-1. De-risk front-loaded: verify the pin lands on origin + on billion (rung zero), drive through a login shell (`ssh host 'bash -lc "…"'`), assert from a tailnet peer (never the VM).
**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).
**Distance-to-gate context:** metric = gate parts (a–h) discharged green live on billion, **0/8** at start (billion is bare — `docker ps` empty, no workspace/pin/cache). Gate met = all (a–h) green cold reset-to-seed + the inherited defers dispositioned + **0 platform edits**. NB (KB-1): gate (b) target is **49/49** pairs (`content-denominator.json` `expected_pairs=49` after M241 EN/IT growth), NOT the historical v2.5 "29/29".
**Next-tik direction:** iter-02 (tik) — the Foundation tik: set billion's `.agentspace/rext.tag` to `sound-check-m243-assign-write-playthrough`, lay out the `stack-demo/` workspace + PAT-over-HTTPS, scp `.agentspace/secrets` + `.agentspace/snapshots`, clone+checkout rext at the pin (`git describe --exact-match` MUST print the tag), run gate (a) ORG-CLEAN read-only, then launch the cold reset-to-seed `--public-host` bring-up through a login shell and drive to a fresh green `autoverify.json`.

## TOK-02: consolidate + sharpen TOK-01 after 5 tiks working the gate-b cluster — 2026-07-22

**Tok type:** triggered (5-no-metric-delta streak since iter-06 — the skill's literal primary-metric floor)
**Prior strategy:** TOK-01 (staged cold billion bring-up → discharge gate parts a–h one-cluster-per-tik).
**Why it stopped (moving the metric):** the primary metric is BINARY PER GATE. Iters 07–11 all worked the single hardest bucket (gate b) plus a gate-d measurement + an inherited-carry cleanup — real, load-bearing progress every iter (voice presence-only, interview player-ack demopatch scope-fix, gate-d falsification, reap-17700 discharge, and iter-11's the-cold-reset-to-seed-DONE-green-at-m244 + the LOAD-BEARING cross-check fix + gate-b fully root-caused) — but NONE flipped a gate bucket, because gate b only counts at 47/47 and sits at 45/47. So the metric reads a flat 3/8 over five productive tiks.
**New strategy:** TOK-01 HOLDS — it is SOUND, not stalled. Refinement: stop DIAGNOSING gate b (it is fully root-caused) and APPLY the fix, then sweep the cheap remaining gates on the now-green m244 seed. Precisely: (1) land gate (b) 47/47 by scoping the `next-web-interview-flag-container` FETCH demopatch to `isManagerScope` (mirror iter-08's result-gate scope) so the interview PLAYER skips the slow extraction fetch → ack renders promptly → the existing settle+retry lands both pairs (re-tag/push rext → re-pin billion → re-bake next-web/hiring → re-sweep); (2) drive gates (c) 40 live-browser specs, (f) 3 drift-carries, (h) v2.6 fixes + p95<5s, and (d) anon /library ant-academy demopatch — all live verifications on the green seed (only b's re-bake and d's academy demopatch need image work); (3) DEF-M239-01 ENOSPC loud-build-fail as budget allows.
**Strategy class:** more-granular — drill from "discharge parts" to the named gate-b fix + a live sweep of the remaining gates (NOT a new direction; TOK-01 is re-affirmed).
**Distance-to-gate context:** 3/8. Gate b is 45/47 (one fix away → 4/8). Gates c/f/h/d unblocked by the green m244 seed. Expect the next few tiks to flip buckets quickly now that the foundation + hardest-cluster diagnosis are done (diminishing risk, not diminishing returns).
**Cross-refs to prior TOKs:** TOK-01 (bootstrap) — this is the first triggered TOK; it does NOT pivot away from TOK-01, it consolidates it after the foundation + gate-b diagnosis and sharpens the next-tik targeting.
**Next-tik direction:** iter-13 (run 6, tik 1) — apply the container-demopatch gate-b fix + re-bake + re-sweep for 47/47; confirm the slow-fetch cause with a quick probe first.

## TOK-03: HOLD TOK-01/02 + sequence the now-narrow final push (gate-c stack-verify half DONE) — 2026-07-23

**Tok type:** triggered (3-no-prog streak: iter-16/17/18 all no-metric-delta on the binary gate-parts metric)
**Prior strategy:** TOK-02 (which itself re-affirmed TOK-01: staged cold billion bring-up → discharge gate parts a–h one-cluster-per-tik, sweeping the cheap remaining gates on the green m244 seed).
**Why it stopped (moving the metric):** the metric is BINARY PER GATE. iters 16/17/18 all did real, load-bearing gate work — iter-16 recalibrated the coverage academy-marker (both hero sweeps GATE MET), iter-17 proved 2 discrete specs + mapped the retrofit, and **iter-18 completed the ENTIRE gate-c stack-verify half (all 10 discrete gate specs green on billion) + fixed a real Bedrock provisioning defect (talk-to-data now answers live) + proved 4/6 gate-h fixes as byproducts** — but NONE ticked a full gate PART, because gate (c) only counts once the 16 Playthroughs land (seed-order-forced LAST) and gate (h) only counts at all-6-fixes+p95. So the metric reads a flat 5/8 over three of the milestone's most productive tiks — the EXACT coarse-binary-per-gate artifact TOK-02 already diagnosed and pre-registered.
**New strategy:** TOK-01/02 HOLD — SOUND, not stalled (this is a false-positive trigger on a coarse metric, fired mechanically per the skill's unconditional floor; no strategy revision is warranted). The refinement is pure SEQUENCING, now that the foundation + gate-b + gate-d + the gate-c stack-verify half are all done and the final push is fully de-risked and NARROW:
  (1) **gate (h) COMPLETION** — 4/6 fixes already proven (talk-to-data/library/cockpit/content-fidelity); remaining = academy course-start + language toggle (M241 EN/IT) + **p95 click→ACCESS < 5s** hero vantages (REFRESH the stale >4h autoverify FIRST; run-latency.sh both heroes, LATENCY_SCHEME=https, STACK_DIR set, mind the M236 UTC age-check) → ticks **6/8**;
  (2) **gate (f)** — 3 v2.3 drift-carries: BURNIN-M221 (needs a `/dev-up --public-host` remote **dev** burn-in — a separate heavy bring-up), F-M220-4 (ant-academy.sh re-runnable on a live public-host demo), PROBE-M218-c3 (Cosmo cms/Directus 403 re-check) → ticks **7/8**;
  (3) **gate (c) tick** — the 16 Playthroughs LAST on pt-world (reset destroys the demo seed) → ticks **8/8 = GATE MET**;
  (4) **DEF-M239-01** (ENOSPC loud-build-fail) as budget allows.
**Strategy class:** more-granular — drill from TOK-02's "sweep the cheap remaining gates" to the named, de-risked final-push sequence, now that the stack-verify half is done + gate-h is 4/6. NOT a new direction (TOK-01/02 re-affirmed).
**Distance-to-gate context:** 5/8. Three buckets remain (h → f → c-playthroughs), each now a clear single-focus tik. Diminishing RISK, not diminishing returns — the hard clusters (foundation, gate-b diagnosis, gate-d academy, gate-c retrofit + the Bedrock defect) are all behind us.
**Cross-refs to prior TOKs:** TOK-02 consolidated after the foundation + the gate-b diagnosis; TOK-03 consolidates after the gate-c stack-verify half + the Bedrock provisioning fix, with the final push fully sequenced. Both are the same coarse-metric-artifact re-affirmation — TOK-01 is the sound spine throughout.
**Next-tik direction:** run 8 iter-20 (fresh tik) = **gate (h) completion**: refresh billion's autoverify (STACK_DIR set), run-latency.sh both hero vantages for p95<5s (LATENCY_SCHEME=https, M236 UTC age-check), + live-check academy course-start + the EN/IT language toggle → tick 6/8.

## D1 (close): Deferral re-audit — 2 inherited carries re-fated Fate-3 → close-release — 2026-07-23

**Context:** Phase 1b deferral re-audit at milestone close (report:
`audit-deferrals/deferral-audit-2026-07-23-m244-close.md`, verdict **YELLOW**).

**Decision:** the two remaining open, pre-M244-inherited carries are re-fated **LAND-NEXT (Fate-3) →
`/developer-kit:close-release`**:
- **M238-D5 — standing demo-stack test debt (8 fails)** (6 `test_cockpit` academy/overlay + `test_public_host`
  port-13001 + `test_purge` docker-integration). Chronic (ridden ≥5 v2.6 milestones). Was 9; the M244 final
  harden fixed the ONE M244-introduced fence (M215 no-pipe, iter-25 describe-form reconcile, Pass 8, Fate-1) →
  8, all provably pre-M244 (SUT untouched by M244).
- **DEF-M239-01 — ENOSPC loud-build-fail** (a build-path change validatable only by inducing a real ENOSPC).

**Rationale:** M244 is a **proof milestone** (`delivers: none`, 0 platform edits) — its charter is driving +
asserting a cold billion bring-up, not authoring demo-stack test fixes or an un-validatable build-path change.
The named expiry across the M238–M243 closes was "M244"; now that M244 closes proof-only, the genuine final
expiry for release-wide test debt is the **release-level deferral audit** (close-release Phase 1b, release
scope, extra scrutiny), which is the very next step. Both resolved carries — **reap-17700 LANDED** (Fate-1,
iter-10) and **DEF-M240-01 dispositioned** (voice player-presence-only, the pre-approved else-branch, iter-07)
— leave the ledger. Not a silent auto-accept: an explicit Fate-3 re-fate with a filed report + a chronic-pattern
flag for close-release. No blocker raised.

## Adversarial review (close) — 2026-07-23

The release thesis is **anti-toothlessness**, so the central adversarial question for a proof milestone whose
"code" is a browser harness is: **could an assertion false-PASS — report a page as rendered/correct when it is
not?** A locator that matches chrome/boilerplate, or a settle that returns on the nav shell, would let a broken
surface sail through green and make the whole live proof toothless.

- **Scenario — a locator matches non-content and false-passes.** iter-27 is itself the proof this risk is real:
  it FALSIFIED iter-26's seed-gap story — the 3 ai-readiness "failures" were HARNESS locator mismatches against
  a **correctly-rendering** v1.341.0 UI (byTeam "Team" not "Tag" / interview panel scoped to a 24-char heading
  span not the findings card / dueDate short-date with no 4-digit year). The mirror risk is the opposite: a
  locator loose enough to pass on an EMPTY page. Pass 7's harden surfaced exactly that — the year-less
  `\w{3,9}` dueDate matcher matched "24 hours"/"5 days" (would false-pass a page with no real deadline) and was
  **month-constrained + mutation-verified RED-on-break**. The iter-16 academy-home marker was likewise moved
  from a `<title>`-only token (false-passed an empty grid) to a card-count floor, mutation-verified.
- **Scenario — settle returns before content paints.** gate (b)'s interview-player residual: `settle()`
  early-exited on the 126–128-char nav shell (cleared `MIN_SETTLED_CHARS=120`) before the ack painted → a false
  `mainLen 0`. Fixed with shape-aware `contentReady(shape)` gating player-interview on the ack in `<body>`, +4
  mutation-verified specs.
- **Response:** every M244 harness assertion with a pure-logic surface was mutation-verified to bite across the
  two harden sessions (Passes 1–9); the operative signal is mutation-verify (no line-cov tool wired). The one
  un-swept surface (`pickFirstSkillPath` live-tailnet retry timing) has no pure-logic seam and is gold-plating.
  No un-mitigated false-pass scenario remains. The gate's teeth stand.

## D2 (close-release): DROP DEF-M239-01 (ENOSPC loud-build-fail) — 2026-07-23

**Context:** the release-close deferral audit
(`../audit-deferrals/deferral-audit-2026-07-23-release-close.md`, verdict **RED** — AGED_OUT). User
fate-decision 2026-07-23: **"Tag now, carry to v2.7."**

**Decision:** **DROP** `DEF-M239-01` — "make the demo build FAIL LOUDLY on ENOSPC."

**Rationale:** the disk-full failure class is **ALREADY caught** by M239's pre-flight Docker-VM disk-measure
(the demo aborts before the build when the VM is low on space). The build-time loud-abort is redundant
belt-and-braces of marginal residual value, and it is **un-validatable** without inducing a real ENOSPC on a
build box — not reproducible in a clean docs-only close. Cutting it retires an aged item honestly rather than
parking un-validatable code. **NOT carried to v2.7** (a DROP, not a defer). Recorded as DROPPED in
`roadmap-vision.md`.

## D3 (close-release): RELEASE-SCOPE-DEFER standing-8 + roster-nit → v2.7 test-health — 2026-07-23

**Context:** the same release-close deferral audit (**RED** — CHRONIC + AGED_OUT). User fate-decision
2026-07-23: **"Tag now, carry to v2.7"** — the fallback path (no late-merge branch opened; the whole set → a
**named** v2.7 milestone, per the v2.5-close rule *a fate needs a MILESTONE, not a sweep/standing-backlog*).

**Decision:** **KEEP-DEFERRED-WITH-SIGNOFF → v2.7 "test-health" (a NAMED milestone):**
- **STANDING-DEMOSTACK-TESTDEBT (8 fails)** — pre-M244 inherited (rext `04babf8`, 2026-07-15), **0 real
  product defects**, host-dependent: 6 `test_cockpit` academy/overlay (stale academy-shape /
  removed-30 s-window assertions vs deliberately-changed M218/M238/M220 behaviour) + `test_public_host`
  (port-13001 hiring-app drift) + `test_purge` (docker-integration, needs a live docker box).
- **rext `stack-verify/e2e/run-unit.sh` roster nit** (same test-health class) — the canonical runner's
  `UNIT_SPECS` roster is missing the 2 specs the M244 final-harden added (`content-denominator`,
  `run-discrete`), so the runner exits 2 and its `test_e2e_collection_integrity::UnitSpecsAreExecuted` guard is
  RED. **All 9 specs pass when run directly** (172) — a roster-drift nit, not a product defect.

**Why Fate-1 (land-now) was DECLINED:** the user chose **tag-now for a proof-complete release** — v2.6's
charter was proving the whole feature live on `billion`, which is MET (exit gate 8/8, content-stories 47/47, 16/16
Playthroughs, p95 1.46/1.31 s, flake 0). These items are **non-defects** (0 real product impact): the standing-8
splits into a mechanical test-assertion subset + a docker/live-gated residue that can only be validated on a live
box, and the roster nit is a 2-line rext edit — all better batched into a dedicated v2.7 test-health pass than
bolted onto a proof close. Recorded under the v2.6 → v2.7 carry in `roadmap-vision.md`.

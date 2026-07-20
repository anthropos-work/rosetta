---
title: "v2.5 «the playbill» — Release-Scope Deferral Ledger + Escape-Hatch Record"
date: 2026-07-20
scope: release
release: v2.5
phase: "close-release Phase 7 / Phase 9"
audit_input: audit-deferrals/deferral-audit-2026-07-20-release-close.md
status: decided
---

# Release-Scope Deferral Ledger — v2.5 "the playbill"

This is the **per-item destination-still-valid ledger** that v2.4's close did not carry (see
[§ Process findings](#process-findings-what-this-close-owes-the-next-one) — `DC-P5`). Every item that
leaves v2.5 undone appears here **once**, with:

1. its **fate** (LAND-NOW / DROP / KEEP-DEFERRED-WITH-SIGNOFF),
2. the concrete **why Fate 1 / Fate 2 / Fate 3 each failed**, per item — not a blanket rationale,
3. a **named handler and a named destination milestone** — never a phase, a pass, or "the next X",
4. a **destination-still-valid check** — is the declared target real, dated, and reachable?

Input: `audit-deferrals/deferral-audit-2026-07-20-release-close.md` (verdict RED · 22 items · 8 aged-out ·
3 unhomed · 5 blocking). Decisions taken by the **user (release owner), 2026-07-20**, interactively at the
release close.

> **Destination reservations.** v2.6 is **not yet designed**, so the milestone ids below (`M237`, `M238`)
> are **reservations** under the same precedent as vision `M205`–`M207` and the retired `M216` — they may be
> renumbered at the v2.6 `/developer-kit:design-roadmap` run, but the **destination is a milestone with a
> declared scope**, not a vibe. This is deliberate: *"the next prove-on-VM"* is precisely the phrasing that
> let `DEF-M226-01` age out twice. Both reservations are recorded in
> [`roadmap-vision.md`](../../roadmap-vision.md) § v2.5 → v2.6 carry.

---

## A. The headline escape hatch — the release ships a number it did not re-prove live

### `CLOSE-D3` — the shipped harness has no live green

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → `M237` (v2.6's FIRST work). Signed off by the user, 2026-07-20.**

**State it plainly: v2.5's headline metric — `29/29` landable (session × action) pairs — is
UNIT-PROVEN, NOT LIVE-RE-PROVEN.** It is tagged on that basis, as a conscious decision.

What actually happened, in order:

1. The `29/29` reading was taken **live on `billion`**, cold reset-to-seed, at rext tag
   **`playbill-m236-hardened`**. That measurement was real.
2. The M236 close then fixed **~10 defects in that same harness** (`content-result-page.ts`,
   `content-stories.spec.ts`, `run-content-stories.sh`, `aggregate-content.py`, `content-pairs.ts`) and
   shipped as `playbill-m236-close-fixes`. **The measuring instrument changed after the measurement.**
3. **This Phase 7 is changing it again**, and two of the changes bear directly on the number:
   - **CQ-1** — `shapeFor` misgrades `asmt-voice-pass` (an ASSESSMENT whose sim slug contains the
     substring "interview") down the `player-interview` branch, which has **no length floor, no score
     check, no feedback check**. **One of the 29 pairs was graded by the wrong check** — a live false PASS
     in the shipped canonical manifest.
   - **CQ-2** — the four new `*.unit.spec.ts` files, **including M236's route-contract pin that carries
     the only literal `29`**, are **executed by no runner**. They are *collected* (they count toward
     `MIN_TESTS=120`) and never *run*. A pin nobody runs is not a pin.
   - Compounded by **CQ-6**: `EXPECTED_PAIRS` is computed from the same served manifest the sweep reads,
     so the denominator is **self-derived** — if the Go projection dropped 3 sessions the sweep would
     report a flawless `26/26`.
4. **No live re-run has happened since.** The current tag has never faced a browser.

**Why Fate 1 (land now) failed — concretely.** It did **not** fail on capability. `billion` is up and
reachable from this workstation; v2.5 ran ten live iters against it this release. A live re-prove is a
next-web rebuild + a cold reset-to-seed + `./run-content-stories.sh 1 --host billion.taildc510.ts.net`.
**It failed on an explicit user choice**: the user elected to **tag v2.5 now and re-prove as v2.6's opening
work**, accepting the residual. This is an escape hatch, not an impossibility, and it is recorded as one.

**Why Fate 2 (covered elsewhere in-release) failed — concretely.** The substitute for a live run is the
unit-spec layer, and **CQ-2 proved four of those specs never execute** — including the one that pins the
number. There is nothing else in the release that touches the live path.

**Why Fate 3 (route to a later in-release milestone) failed — concretely.** **M236 was the final v2.5
milestone.** There is no downstream destination inside the release. This is structural, not a scheduling
miss.

**Handler:** the **release owner (user, `kiralise`)**, executing reserved milestone **`M237 —
re-prove-on-billion`** as the declared first work of v2.6.

**Destination-still-valid check:** ✅ valid. `billion` is live and reachable (verified this release, 10
live iters + a cold bring-up at stable `main`). No infra prerequisite is outstanding.

**Acceptance condition for `M237` on this item:** a live `run-content-stories.sh` green **at the tag that
actually shipped** (post-Phase-7), with the CQ-1 grader fix and CQ-2 runner wiring in place, and with
`EXPECTED_PAIRS` sourced from something other than the artifact under test (CQ-6). Anything less
re-issues the same false green.

### `T-3` — the 39 unexecuted live-browser specs

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → `M237`. Signed off by the user, 2026-07-20.**

**24 `stack-verify` + 15 `playthroughs` live-browser specs were not executed at this close.** They require
a running demo stack. **The entire prove-by-render layer is unverified in this close** — which is to say
the release's *other* pillar (`coverage-protocol.md` proves presence; `playthroughs.md` proves function)
has no green either.

- **Why Fate 1 failed:** same as `CLOSE-D3` — possible, deliberately not taken. Same user decision, same
  batch.
- **Why Fate 2 failed:** the executed TS layer is **196 unit specs**, none of which drive a browser. Unit
  coverage is not a substitute for the specs that exist precisely because unit coverage is not enough —
  **the release's own `DC-P3` names this class: *unit-proven ≠ route-proven*.**
- **Why Fate 3 failed:** no in-release milestone remains.

**Handler:** release owner (user), via **`M237`** — the same single bring-up that discharges `CLOSE-D3`.
**Destination-still-valid check:** ✅ valid (same live host).

> **One live bring-up discharges eight items** (`CLOSE-D3` · `T-3` · `S-1` · the `apps/web` `:5050`
> endpoint · `DEF-M226-01` · `BURNIN-M221-dev-public-host` · `F-M220-4` · `PROBE-M218-c3-rerun`). That is
> the whole argument for `M237` being one milestone and not eight backlog rows.

---

## B. The aged-out set — per-item fates

### S-1 · `ACADEMY-M236-iter08-public-catalog-twin` — the release's only Fate-3-UNDELIVERED item

Anonymous academy routes (`/library`, `/free`) render **0 cards** — `getPublicCatalogView`'s `new Set()`
branch is uncovered by the M230 patch. Chain M230 → M235 → M236 → this close; **carried in the destination
`In:` list at every hop**, so this is *disclosed non-delivery*, not a silent drop. Aged out across 3
milestones.

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → `M237`, batched with the live re-prove. Signed off 2026-07-20.**

**The audit recommended LAND-NOW-batched; the batch is what moved, so the item moved with it.** The
argument for not splitting it out:

- **Why Fate 1 failed:** landing it needs a **2nd demopatch manifest + a next-web image rebuild + a live
  re-prove**. The first two without the third would ship an **unverified fix to a render defect** — i.e.
  a change whose only possible proof is the live run this close has consciously deferred. Shipping a
  render fix nobody rendered is exactly the defect class this release is named for
  (*"a check can report success while proving nothing"*). Landing it half-way is **worse** than routing it
  whole.
- **Why Fate 2 failed:** nothing else in v2.5 touches the anonymous catalog path. The signed-in path is a
  different code branch (M230's patch covers it; that half **did** land — see S-7).
- **Why Fate 3 failed:** M236 was the last milestone.

**Handler:** release owner (user), via **`M237`**. **Destination-still-valid check:** ✅ valid.
**Note:** this item **IS** the surviving half of the v2.3-era `F4` (see S-7) — they are one item, fated
once, here.

### S-2 · M231 `KB-2` / `KB-6` / `KB-8` — UNHOMED (Fate-3 with no destination milestone)

M231 classed these *"Tracked (Fate-3, arch-doc pass)"*. They appeared in **no** milestone `In:` list, **no**
`carry-forward.md`, and **not** in `state.md`'s standing backlog. They were not deferred — they were
**untracked**, and all three verified still stale in the corpus at this close.

**Fate: LAND-NOW (Fate 1) — IN FLIGHT THIS CLOSE.** Being landed now by the concurrent Phase 7 agent that
owns `corpus/services/**`: `jobsimulation.md:17,44` (ports 8400/8401 → 8080/8081) ·
`roadrunner.md:63` (the orphaned-repo / in-process-runner claim) · `backend.md:25,109`
(`internal/labsession` → `internal/labs/session`, which is also `D-11`). `KB-5` is **moot** — discharged
incidentally; the `ai_architecture.md` claim is no longer present. **`KB-6` is ~6 docs wider than M231
recorded.**

**Handler:** Phase 7 corpus-docs agent, this close. **Destination-still-valid check:** ✅ — landing now.

> **CLASS NAMED — the invalid destination.** ***"An arch-doc pass" is not a milestone. Neither is "a
> sweep", "the next close", "a future build-iter", or "standing backlog". A fate needs a MILESTONE — a
> named, dated, schedulable unit of work with an owner.*** A Fate-3 whose destination is a *kind of
> activity* rather than a *unit of work* has no `In:` list to appear in, no ledger to be checked against,
> and no close that can fire on it. It is indistinguishable from a drop, except that nobody agreed to
> drop it. Three of this release's eight aged-out items (S-2, S-10, and half of S-6) failed exactly this
> way.

### S-4 · `DEF-M235-03` / M204 **assign-WRITE** — ~10 routings across 5 releases

The declared in-manifest TODO half of the M204 assign flow. Declared out at M201, honest and tracked
in-manifest since v2.0. Its declared destination was **the v2.4 close, which fired 2026-07-18 with no fate
taken**; it then silently re-anchored on v2.5. **Both v2.5 milestone audits recorded it as "correctly
routed"** (`m236/decisions.md:210`, `m236/carry-forward.md:91`) — which was true of the *routing* and false
of the *destination*. Verbatim `AGED_OUT`.

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → `M238 — playthrough-assign-write`. FRESH DATED SIGN-OFF 2026-07-20.**
Not a rubber-stamp of the prior sign-off; the prior destination is void.

- **Why Fate 1 failed:** the assign-WRITE half needs a **live browser drive against a running demo stack**
  to stabilize (it is the write side of a flow whose read side is already green), and it is **not** in the
  `M237` batch — `M237` is a *re-prove of what shipped*, and assign-WRITE is *net-new coverage*. Folding
  net-new work into a re-prove milestone is how `M237` becomes a bucket and stops firing.
- **Why Fate 2 failed:** nothing in v2.5 touched the playthroughs assign path. v2.5's content-stories work
  is a different surface entirely.
- **Why Fate 3 failed:** no v2.5 milestone owned playthroughs coverage; M236 was the last.

**Handler:** release owner (user), via reserved **`M238`**, alongside the reserved Playthroughs futures
`M206`/`M207` in [`roadmap-vision.md`](../../roadmap-vision.md).
**Destination-still-valid check:** ✅ valid — `M238` is newly declared at this close and has not yet fired.
**Expiry condition:** if `M238` is not designed into v2.6 or v2.7, this item is **DROPPED**, not
re-anchored a sixth time. Ten routings is the limit.

### S-5 · `DEF-M226-01` — the pre-bind serve reap — **AGED OUT TWICE**

Clear stale `tailscale serve` fronts on a demo's offset ports before bind (M226 Finding-3). Deferred M226
(2026-07-17) with destination *"a follow-up build-iter / **M228** (the next prove-on-VM)"*.
**Trigger 1:** M228 closed 2026-07-18 without it. **Trigger 2:** M236 — the next prove-on-VM — closed
2026-07-20 without it either. `state.md` **still names M228**, stale for two releases.

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → `M237`, WITH TEETH. Signed off 2026-07-20.**

- **Why Fate 1 failed:** it is a **bring-up-path change on a live-only surface** — it cannot be proven
  without a live bring-up, which is the deferred `M237` batch. Landing the code without the bring-up
  changes the bring-up path on faith.
- **Why Fate 2 failed:** its standing justification — ***"self-resolves in the default flow"*** — is the
  reason it has passed **three** times, and **it has never been tested.** An untested self-resolution
  claim is not coverage; it is the `DRIFT_DEFER` signature.
- **Why Fate 3 failed:** the destination it was routed to (*"the next prove-on-VM"*) is a **category, not a
  milestone** — the exact anti-pattern named in S-2 — so it aged out silently, twice, with nothing to fire.

**Handler:** release owner (user), via **`M237`**, which **must TEST the self-resolution claim, not restate
it.** **Expiry condition — this is the teeth:** if `M237` completes **without** testing whether the reap
self-resolves in the default flow, `DEF-M226-01` is **DROPPED at the v2.6 close**, not carried a fourth
time. A justification that survives three passes untested has forfeited its claim on a fourth.

**Destination-still-valid check:** ✅ valid — and note that this is the **first time** this item has a
destination that is a milestone rather than a phrase.

### S-6 · The four v2.3 tail carries — reclassified `DRIFT_DEFER`

Signed off at the v2.3 close (2026-07-15) as KEEP-DEFERRED-WITH-SIGNOFF → **v2.4**. **v2.4 shipped
2026-07-18 without any of them and without a release-scope audit.** `roadmap-vision.md` still read
*"v2.4: investigate…"* and had **no v2.5 carry section at all** — so for two releases the four have been
described as v2.4 work while v2.4 was archived.

| item | fate | destination | note |
|---|---|---|---|
| **F4** — academy grid renders 0 cards | **SPLIT — partly DISCHARGED** (see S-7) | remainder = S-1 → `M237` | the signed-in half genuinely landed at M230/M236 |
| **BURNIN-M221-dev-public-host** — dev-path `--public-host` live burn-in | KEEP-DEFERRED-WITH-SIGNOFF | `M237` | `DRIFT_DEFER` |
| **F-M220-4** — `ant-academy.sh` re-runnable on a live public-host demo | KEEP-DEFERRED-WITH-SIGNOFF | `M237` | `DRIFT_DEFER` |
| **PROBE-M218-c3-rerun** — Cosmo cms/Directus 403 re-check | KEEP-DEFERRED-WITH-SIGNOFF | `M237` | `DRIFT_DEFER` |

**The stated rationale is now FALSE and is corrected here.** All three surviving items were deferred as
*"needs live infra."* That was true when written (2026-07-15). **It is not true now:** `billion` is up,
reachable from this workstation, and v2.5 ran **ten live iters** plus a full cold bring-up against it this
release. *"Needs live infra"* has silently become *"nobody batched it"* — the textbook `DRIFT_DEFER`
failure mode, where a once-true blocker is re-asserted after it has dissolved.

- **Why Fate 1 failed for all three:** each needs the live bring-up that is the deferred `M237` batch.
  Not an infra blocker — a **batching** decision, and it is recorded as one rather than dressed as an
  impossibility.
- **Why Fate 2 failed:** v2.5's live work (M236) drove the *demo* path on `billion` for the
  content-stories and academy surfaces. None of the three sits on that path (dev-path `--public-host`;
  academy re-run contention; Cosmo cms 403), so M236's green says nothing about them.
- **Why Fate 3 failed:** M236 was the last v2.5 milestone.

**Handler:** release owner (user), via **`M237`**.
**Destination-still-valid check:** ✅ valid — and, unlike the v2.3 sign-off, the destination is now a
milestone with a live host already proven reachable, recorded in a **v2.5 → v2.6 carry section that now
exists** in `roadmap-vision.md`.

### S-7 · `F4` is partly discharged — the debt was over-stated

**Fate: DISCHARGE RECORDED (Fate 1, retroactive). No further work owed on the discharged half.**

`state.md` listed **F4 whole** as parked, which over-states the standing debt. The truth:

- **Signed-in half — LANDED.** v2.5 M230 + M236: the academy grid renders **65 real cards, 0 Draft chips**,
  483 chapter links, proven **cold on `billion`** (iter-10). This was the substance of `F4`.
- **Anonymous half — SURVIVES**, and it is **literally S-1** (`ACADEMY-M236-iter08-public-catalog-twin`).

**One item, not two.** `F4` is retired as an id; its remainder is tracked under the S-1 handler. `state.md`
is corrected at this close to stop listing `F4` whole.

### S-8 / S-9 · M232's interview plan-section render fidelity — **ONE item, double-routed, never run**

These are **the same item**, which is part of why it vanished:

- `m232/spec-notes.md:40` — *"render fidelity (exact plan match) is **M234** coverage's concern."*
- `m232/progress.md:49` — *"**Confirmed-covered (Fate 2):** INTERVIEW report exact plan-section render
  fidelity → **M235**."*

**Neither ran it.** M235 closed `closed-incomplete` and **its coverage gate never ran** — the property was
deferred into a gate that did not fire. The item appears in **no** M235 carry-forward and in **none** of
`DEF-M235-01..04`. Where a gate *did* eventually run (M236, live, 29/29), it measured a **weaker property**:
that the interview result page renders **non-empty**, not that the seeded section-ids **match the replayed
CMS `ExtractionPlan`**. A believable interview report needs the plan match; `29/29` does not assert it.

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → `M237`, as an EXPLICIT ADDED ASSERTION. Signed off 2026-07-20.**

- **Why Fate 1 failed:** proving plan-section alignment requires driving the live interview result page
  against the replayed CMS plan — the deferred live batch again.
- **Why Fate 2 failed:** the Fate-2 was *claimed* (twice) and is **exactly what failed**. Re-asserting it
  would be the third claim of coverage by a gate that has never checked this property.
- **Why Fate 3 failed:** M236 was the last milestone; and the routing it would have used had already
  demonstrated it does not survive a hop.

**Handler:** release owner (user), via **`M237`**, which must add a **specific plan-section-id alignment
assertion** — not inherit the non-empty check and call it covered.
**Destination-still-valid check:** ✅ valid.

> **CLASS NAMED — the invisible deferral.** ***"Confirmed-covered" is the deferral form that leaves no
> ledger entry, and is therefore invisible to every downstream audit.*** A Fate-2 says *"another milestone's
> gate already covers this"* — so it is written as a **note**, not a **carry**. It never enters a
> `carry-forward.md`, never gets a `DEF-` id, never appears in an `In:` list. **If the covering gate does
> not run — or runs and measures something weaker — nothing anywhere notices.** A Fate-3 that fails leaves
> an aged-out trigger; a Fate-2 that fails leaves **silence**. Mitigation, for the next close to adopt:
> **a Fate-2 must name the covering gate's assertion, not just the covering milestone**, and the covering
> milestone's close must verify that assertion actually ran. A Fate-2 whose covering milestone closes
> `closed-incomplete` **automatically reverts to an open Fate-3** and must be re-fated.

### S-10 · `F11` / `DEF-M215-03(a)` — routed to "standing backlog", never arrived

Seed hero identity-key vs generated profile display-name mismatch (`maya-thriving` ≠ the generated display
name). Cosmetic; login and render both work. Routed at the v2.2 close (2026-07-11) to **"standing
backlog"** — and it is **absent from `state.md:120-147`**, i.e. absent from the standing backlog it was
routed to. Half (b) — the committed remote-origin Playwright gate — **was** genuinely discharged at
M218/M221.

**Fate: KEEP-DEFERRED-WITH-SIGNOFF → the standing backlog, NOW ACTUALLY WRITTEN DOWN. Signed off
2026-07-20.**

- **Why Fate 1 failed:** it is a `stack-seeding` identity/display-name polish item with **no live proof
  required** — it *is* cheaply landable, but it is genuinely cosmetic (zero functional impact, verified
  across three releases of live demos), and this close's Phase 7 scope is already the full finding list.
  Landing it here would be scope-creep on a close that is closing; it is honest to say so.
- **Why Fate 2 failed:** no seeding milestone in v2.5 touched hero identity keys.
- **Why Fate 3 failed:** no v2.5 milestone owned seed polish.

**Handler:** the **seed-polish owner at the next `stack-seeding` build-iter**; tracked in `state.md`'s
standing backlog **by id** (`DEF-M215-03(a)` / `F11`) so it is findable — which it was not for three
releases. **Destination-still-valid check:** ⚠️ **weak but now real** — "standing backlog" is not a
milestone (per the S-2 class), and this is a **deliberate exception for a verified-cosmetic item**: the
mitigation is that it is now **enumerated with an id in `state.md`**, so the next audit can see it. If it
survives another release un-landed, it should be **DROPPED** rather than carried — a cosmetic item nobody
will schedule is a drop wearing a carry's clothes.

---

## C. Re-confirmed carries (no change, fresh date)

Per the audit's item 12 — already in `roadmap-vision.md`, re-confirmed **2026-07-20**, no other action:

| item | fate |
|---|---|
| `DEF-M10-01` — cloud `SnapshotStore` / S3 blob bytes | KEEP-DEFERRED-WITH-SIGNOFF (gated on eu-west-1 S3 read access, still not wired) |
| `DEF-M21-01` — `replayCmd` conn-seam hermetic test | KEEP-DEFERRED-WITH-SIGNOFF (needs an injectable connector seam) |
| `CAVEAT-1` — clean-box literal full destructive `/dev-up` | KEEP-DEFERRED-WITH-SIGNOFF (this box is committed to native-app content-line dev) |
| `M314b` — prod frozen-read whole-org AI-readiness hydration | KEEP-DEFERRED (platform-team pointer; no rosetta work owed) |
| `M205`-residual — tier gates + ATS pipeline | reserved in vision (no Stripe mirror engine; the platform does not model ATS) |

## D. Landed this close (recorded for the destination-still-valid trail)

| item | fate | where |
|---|---|---|
| M231 `KB-2`/`KB-6`/`KB-8` | **LAND-NOW** | Phase 7 corpus-docs sweep (S-2) |
| `metrics-history.md` missing v2.0 / v2.2 / v2.4 rows | **LAND-NOW** | Phase 7 (`KB-D`) — `state.md` under-reported this as "v2.0 + v2.2" |
| `F4` signed-in half | **DISCHARGED** | v2.5 M230 + M236, live on `billion` (S-7) |
| M230 cluster 2 (next-web clone re-anchor) | **DISCHARGED by falsification** | M236 Phase-0b `F20` — the clone was never drifted |
| the standing demo-stack test-debt carry | **RE-BASELINED + fated at the M236 close** | 8 macOS / 7 Linux, 0 real defects, 0 pin drift |
| `F-M236-CLOSE-1` / `F-M236-CLOSE-2` | **LAND-NOW** (Phase 7, rext-side) | clone freshness + the R1 pristine sweep |

---

## Process findings (what this close owes the next one)

### `DC-P5` — the close template can issue a false green, and did

**v2.4's own `release-review.md` is the proof.** On 2026-07-18 it asserted:

- `:37` — *"✅ Per-milestone decisions blended into corpus at each close"* **and** *"✅ No cross-milestone
  decision conflicts"* — **while six items were aging out**, four of which no audit had recorded.
- `:16` — *"✅ no undelivered Fate-3 items"* — while `release-retro.md:41`, **written the same day**, listed
  two as inherited standing carries.
- `:34` — pointed the reader to a *"**Completeness Ledger below**"*. **The file ends at line 42. There is
  no such section.**

**That dangling pointer is the concrete mechanism** by which the standing test-debt carry and
`DEF-M226-01` left v2.4 with no landing record. The review told the reader where to check, the place did
not exist, and nobody followed the pointer to find out. A ✅ next to a claim whose evidence section was
never written is **worse than a ⛔** — it actively closes the question.

This is the same class as the release's own thesis (*"a check can report success while proving nothing"*),
one level up: **a review can report a green while proving nothing, and a green review is the thing
downstream audits trust most.** v2.4 also ran **no release-scope deferral audit at all** — the structural
cause of 7 of this close's 8 aged-out items.

**What v2.5 does differently, and what the next close must inherit:**

1. **This ledger exists** — a real per-item ledger, not a pointer to one, with a per-item
   destination-still-valid check.
2. **Every fate names a milestone**, never a phase, a pass, or "the next X" (the S-2 class).
3. **Fate-2 "confirmed-covered" is treated as a first-class deferral** with the mitigation in the S-8
   class note — because it is the form that leaves no trace.
4. **A ✅ in a release review must cite the artifact that proves it.** A checkmark whose evidence is
   "below" is invalid unless "below" exists. *Suggested fence: the close should refuse to emit a review
   containing a forward reference to a section heading that is not in the file.*
5. **The escape hatch is written where the number is** (§A), not in a footnote — because the number will
   be quoted and the footnote will not.

v2.4's archived `release-review.md` has been **annotated in place** with a dated correction block. It was
**not silently rewritten**: the false greens stay visible, because the record of *how a close issued them*
is the finding.

### Reservation trail

- **`M237 — re-prove-on-billion`** — v2.6's declared FIRST work. Discharges: `CLOSE-D3` · `T-3` (39 live
  specs) · `S-1` (anon academy twin, = surviving `F4`) · the `apps/web` non-offset `:5050` client GraphQL
  endpoint · `DEF-M226-01` (with a TEST of the self-resolution claim) · `BURNIN-M221-dev-public-host` ·
  `F-M220-4` · `PROBE-M218-c3-rerun` · the S-8 interview plan-section assertion.
- **`M238 — playthrough-assign-write`** — `DEF-M235-03` / M204 assign-WRITE, with a drop-expiry.

Both recorded in [`roadmap-vision.md`](../../roadmap-vision.md) § v2.5 → v2.6 carry.

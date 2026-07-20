---
title: "Deferral Audit — v2.5 «the playbill» release close"
date: 2026-07-20
scope: release
release: v2.5
invoked-by: close-release (Phase 1b)
---

## Verdict

**RED** — `SEVERITY=blocker`.

Not because the v2.5 milestone chain leaked (it did not — every in-release Fate-3 routing landed
in its declared target except one, which is disclosed and routed here). RED because **five aged-out
groups arrive at this close without a fate**, four of which **no prior audit had recorded as aged
out**, and one of which was **never homed in any milestone plan at all**.

The single structural cause: **v2.4 shipped with no release-scope deferral audit** (`releases/archive/
02.40-casting-call/` has four milestone-scope `audit-deferrals/` dirs and none at release root). Every
item whose declared destination was "the v2.4 release close" or "the v2.4 window" therefore fired
without anyone checking. This close is the first release-scope audit to run since v2.3.

## Summary

- Total deferrals in scope: **22** across 6 groups
- Landed in-release (Fate 1/2/3 verified): **11**
- Discharged / falsified: **2**
- Single deferrals awaiting fate: **4**
- **Repeat deferrals: 1** (the standing test-debt carry)
- **AGED_OUT: 8** (1 previously recorded, **7 newly detected by this audit**)
- **Unhomed (Fate-3 with no destination milestone): 3**
- Chronic patterns flagged: **2** (`CHRONIC_DEFER` on the test-debt carry; `DRIFT_DEFER` on the
  v2.3 tail-carry side-track)

---

## 1. In-release routings — VERIFIED LANDED

Every Fate-3 routed inside v2.5 was checked against the target milestone's `overview.md` `In:` list
**and** against the target's delivery ledger. These are clean:

| routing | target | `In:` carries it? | delivered? |
|---|---|---|---|
| M231 D3 — interview PostHog flag-enablement | M232 | yes (overview edited at M231 close) | yes — 2 sha-pinned demopatches |
| M231 D4 — AI-labs presence-only | M234 | yes | yes — presence row, no CTA |
| M231 D5 — academy real-progress form | M234 | yes | yes (render disposition); remainder → M235 |
| M232 D2 — per-session player seats | M234 | yes | yes — `content-player-<idx>` via `roster.go` |
| M232 D6 — interview report render fidelity | M235 | yes | yes — via M236 live (29/29) |
| M233 DEF-M233-01 — bring-up export wiring + tab read + seat registration | M234 | **partly** (see §5.1) | yes |
| M233 DEF-M233-02 — non-sim player-path builders | M234/M235 | yes | yes — `content_nonsim.go` |
| M234 DEF-M234-01 — non-sim fixtures + prove-CTAs-land | M235 | yes (`In:` + exit_gate) | yes → proven live M236 |
| M235 cluster 1 — LIVE proof + seat-login harness | M236 | yes (iter-08, `54eaefe`) | yes — authored + calibrated, 29/29 |
| M235 cluster 2 — per-section live-calibration | M236 | yes | yes — skill-path/ai-labs/academy all worked |
| M230 cluster 1 — formal cold card-count proof | M235→M236 | yes | yes — 65 cards, 0 Draft chips (iters 08/10) |

**M230 cluster 2** (next-web clone re-anchor) — **DISCHARGED by falsification**: M236's own Phase-0b
audit (F20) proved the clone was never drifted; the working tree was left-patched by an un-reverted
run. Correctly recorded.

---

## 2. THE ONE IN-RELEASE FATE-3 THAT DID NOT LAND

### DEF-M230-03 → `ACADEMY-M236-iter08-public-catalog-twin`

- **Item:** the `getPublicCatalogView` 2nd manifest — anonymous academy `/library` + `/free` render 0 cards.
- **Chain:** M230 iter-02 (2026-07-19, Fate-3 → "M235 next-iter queue") → M235 `overview.md` `In:`
  (inherited, verified present at lines 32–36) → M235 carry-forward cluster 3 (Fate-3 → M236) →
  M236 `overview.md` `In:` (verified present, lines 70–74) → **NOT delivered** → routed to this close.
- **Status:** the destination `In:` list carried it at **every** hop — so this is *not* the silent
  failure mode. It is a disclosed non-delivery, recorded in M236 `progress.md` §Carried forward and
  `carry-forward.md` Route 4.
- **Ageing:** deferred across **3 milestones** (M230→M235→M236) → **AGED_OUT**. Requires a fresh fate.
- **Note:** this item **is** the surviving half of the v2.3-era `F4` carry (see §4.3). The signed-in
  half of F4 was genuinely landed by M230/M236; only the anonymous half remains.

---

## 3. THE REPEAT — the standing demo-stack test failures

### REPEAT + AGED_OUT: "N pre-existing demo-stack test failures"

- **First deferred:** M224, 2026-07-17, as **8**.
- **Carried through:** 10 milestones, 2 releases (v2.4, v2.5). Re-confirmed at M232 (D8), M233
  (DEF-REL-testdebt), M234 (DEF-M234-02), M235, M236.
- **Declared destination:** *the v2.4 release close* — which **fired 2026-07-18 without landing it**
  (AGED_OUT trigger; first recorded at the M236 close, not before).
- **Label drift:** 8 → 14 under a fixed label, with the stated *class* changing from stale-tests to
  `pre_sha256` pin drift.
- **Pattern:** `CHRONIC_DEFER`.
- **User decision 2026-07-20 (M236 `decisions.md` CLOSE-D2 → RESOLUTION):** *re-baseline now, decide
  the fate at release close.* Executed — `m236-prove-on-billion/rebaseline-standing-failures.md`.

**Re-baseline result (verified present and internally consistent):**

| | |
|---|---|
| Reproduced before any change | **14** — the count was accurate |
| On a clean stable-`main` clone set | **8** (macOS; **7** expected on Linux) |
| Real product/tooling defects | **0** |
| `pre_sha256` pin drift | **0 — diagnosis REFUTED** |

Six of the fourteen were a **dirty clone** reporting itself as a test failure. The remedy the old
label implied — re-anchor the drifted pins — would have re-pinned five manifests to *patched*
content, permanently disarming the drift detector. **The deferral's own proposed fix was the
dangerous action.**

**This audit's recommendation: LAND-NOW (Fate 1).** Concurs with the re-baseline. All 8 are
test-side edits with no product risk, no live stack, no platform edit: re-point 4 academy assertions
at M234 semantics (or retire them with the M53 feature they describe), delete/invert 2 overlay
assertions, add `13001` to one expected port list, convert the purge test's precondition to
`skipUnless(Linux)`. **Sequencing: fix `F-M236-CLOSE-2` first** (§4.1) or the clone-dependent
failures return on the next interrupted build.

---

## 4. NEWLY DETECTED — items no prior audit recorded

### 4.1 F-M236-CLOSE-1 / F-M236-CLOSE-2 (new at M236 close, no prior deferral)

| id | item | fate recommendation |
|---|---|---|
| `F-M236-CLOSE-1` | `/demo-up` rebuilds images from clones it **never updates** — `make init` is skip-if-present, no fetch/pull/checkout anywhere in the bring-up. Measured identically on both boxes: `app` **249** commits behind `main`, `next-web-app` **202**. `clones.lock.json` records `ref: "HEAD"` for every detached clone, so it is structurally unable to distinguish *pinned* from *stale*. Surfaced as a **user-reported** stale left menu. | **LAND-NOW, highest priority.** A ref policy for the demo clone set + honest `ref` recording in the lock file. rext-side, zero platform edits. Treat **separately** from the 14 — it is their upstream generator and must not inherit their fate. |
| `F-M236-CLOSE-2` | `ensure-clones.sh` R1 "ensure pristine" sweep covers **3** manifests; `patches/` carries **~15**. Any patch outside those three, left applied by an interrupted build, survives every subsequent bring-up. Both boxes were carrying leftovers in disjoint sets. | **LAND-NOW, before the 14.** Extend R1 to enumerate `patches/` rather than a hard-coded triple. Cheap and mechanical. |

### 4.2 DEF-M226-01 — the pre-bind serve reap — **AGED_OUT TWICE, never recorded**

- **Item:** clear stale `tailscale serve` fronts on a demo's offset ports before bind (M226 Finding-3).
- **First deferred:** M226, 2026-07-17. Destination: *"a follow-up build-iter / **M228** (the next prove-on-VM)"*.
- **Trigger 1:** **M228 closed 2026-07-18 without landing it.** v2.4's `release-retro.md:41` lists it
  as an "inherited standing carry" — the destination fired, and because v2.4 ran **no release-scope
  audit**, nothing flagged it.
- **Trigger 2:** **M236 — the next prove-on-VM — closed 2026-07-20 without landing it either**, and
  `state.md:137-139` *still* names the destination as "M228". The destination string has been stale
  for two releases.
- **Standing justification:** *"self-resolves in the default flow."* That claim has been the reason
  for **three** passes and **has never been tested**. It is the signature of `DRIFT_DEFER`.
- **Recommendation:** **LAND-NOW, batched** into the single live re-prove §6 already requires
  (`billion` is up and reachable). If the user declines the live re-prove, then
  **KEEP-DEFERRED-WITH-SIGNOFF with a NAMED milestone** — never again "the next prove-on-VM", which
  is precisely the phrasing that let it age out twice.

### 4.3 The four v2.3 tail carries — **AGED_OUT ×4, never re-fated**

Signed off at the v2.3 close (2026-07-15) as KEEP-DEFERRED-WITH-SIGNOFF → **v2.4**, with
`roadmap-vision.md` §"v2.4 tail-carry SIDE-TRACK" entries reading *"v2.4: investigate…"*.
**v2.4 shipped 2026-07-18 without landing any of them and without a release-scope audit.**
`roadmap-vision.md` has **no v2.5 carry section at all**; the four still read as v2.4 work.

| item | status now | recommendation |
|---|---|---|
| **F4** — academy grid renders 0 cards | **SPLIT / partly discharged.** The signed-in half was **genuinely landed** by v2.5 M230 + M236 (65 cards, 0 Draft chips, live on billion). Only the **anonymous** half survives — and it is literally §2's `ACADEMY-M236-iter08-public-catalog-twin`. `state.md:140-142` still lists F4 whole as parked, which now over-states the debt. | **Record the discharge**; fold the remainder into §2 and fate them as one item. |
| **BURNIN-M221-dev-public-host** — dev-path `--public-host` live burn-in | needs live infra | **LAND-NOW batched** into the single re-prove, or fresh KEEP-DEFERRED-WITH-SIGNOFF → v2.6 |
| **F-M220-4** — `ant-academy.sh` re-runnable on a live public-host demo | needs live infra | same |
| **PROBE-M218-c3-rerun** — Cosmo cms/Directus 403 re-check | needs live infra | same |

All three "needs live infra" reasons were true when written. **They are no longer obviously true**:
`billion` is up, reachable from this workstation, and v2.5 just ran ten live iters against it.
"Needs live infra" has silently become "nobody batched it", which is the `DRIFT_DEFER` failure mode.

### 4.4 M231 KB-2 / KB-6 / KB-8 — **UNHOMED: Fate-3 with no destination milestone**

M231 classed four KB-fidelity findings as *"Tracked (Fate-3, arch-doc pass)"*. **"An architecture-doc
pass" is not a milestone.** These appear in **no** milestone `In:` list, **no** `carry-forward.md`,
and **not** in `state.md`'s standing backlog. They are not deferred — they are untracked.

Verified **still stale in the corpus today**:

| id | claim | verified |
|---|---|---|
| KB-2 | `corpus/services/jobsimulation.md:17` — *"Ports: 8400 (GraphQL/HTTP), 8401 (Connect-RPC)"*; repo CLAUDE.md says 8080/8081 | **still present** |
| KB-6 | `corpus/services/roadrunner.md:63` — *"jobsimulation (the only caller — `ROADRUNNER_RPC_ADDR=…`)"*; code execution is in-process in `jobsimulation/internal/runner/`, the repo is orphaned | **still present** |
| KB-8 | `corpus/services/backend.md:25,109` — labs path `internal/labsession`; actual is `internal/labs/session` | **still present** |
| KB-5 | `ai_architecture.md` *"Judge0 via the Roadrunner service"* | **no longer present** — discharged incidentally; mark moot |

**Recommendation: LAND-NOW (Fate 1).** Three one-line corpus corrections, zero risk, no code, no
stack. There is no defensible reason for these to leave the release untracked.

### 4.5 `metrics-history.md` missing v2.0 + v2.2 rows — AGED_OUT

Destination recorded as *"next close-release"*. The next close-release was **v2.4's**, which fired
without landing it. **Recommendation: LAND-NOW** — this close writes metrics at Phase 8/9 anyway.

---

## 5. Process findings (record integrity, not scope)

1. **v2.4 ran no release-scope deferral audit.** The single structural cause of §4.2 and §4.3. This
   close should record it, because the same omission would re-hide the same class next release.
2. **M233's `DEF-M233-01` routing claim was over-stated at source.** Its audit justified the Fate-2
   on the basis that *"M234's `In:` list carries it"*; the **bring-up export-wiring** half is not in
   that list. M234 delivered it anyway (ledger Fate-1, §3), so nothing stranded — but the provenance
   is unsound, and this is exactly how a Fate-2 becomes a silent drop when the target *doesn't*
   happen to deliver.
3. **M234 `In:` item 6 was narrowed without a partial label.** The bullet promised an academy CTA
   deep-linking to *the chapter with real progress*; what shipped was a generic academy-origin link,
   with the remainder moved to Fate-2 → M235. M236 iter-08 then found the route named in that CTA
   (`/library/[slug]`) **does not exist** and its unit test *required* the nonexistent prefix.
4. **M235's `progress.md` cites a "Gate Outcome Ledger below" that is not in the file.** The gate
   data lives only in `carry-forward.md` frontmatter. Minor record gap; the substance exists.
5. **M234's ledger claims `(#M234-DK)` back-ref tags.** The tags that shipped are `#M234-D1/-D2/-D4/-D5`;
   `M234-DK` appears nowhere. Cosmetic typo in the close record.

---

## 6. Recommendations — consolidated

**One live re-prove discharges six items.** `billion` is up and reachable. A single next-web rebuild
+ cold reset-to-seed + sweep would land: the anon academy catalog twin (§2), the `apps/web` non-offset
`:5050` endpoint, CLOSE-D3's live re-run of the content-stories sweep, DEF-M226-01, and plausibly
BURNIN-M221 / F-M220-4 / PROBE-M218-c3. Recommend batching rather than fating them one by one.

| # | item | recommended fate |
|---|---|---|
| 1 | `F-M236-CLOSE-2` — pristine sweep 3 of ~15 | **LAND-NOW** (first) |
| 2 | `F-M236-CLOSE-1` — demo rebuilds from never-updated clones | **LAND-NOW** (highest priority; separate from the 14) |
| 3 | standing demo-stack test failures (8 macOS / 7 Linux) | **LAND-NOW** |
| 4 | M231 KB-2 / KB-6 / KB-8 corpus staleness | **LAND-NOW** |
| 5 | `metrics-history.md` v2.0 + v2.2 rows | **LAND-NOW** |
| 6 | anon academy `/library` + `/free` (= surviving half of F4) | **LAND-NOW batched** w/ the single re-prove |
| 7 | `apps/web` client GraphQL on non-offset `:5050` | **LAND-NOW batched** |
| 8 | CLOSE-D3 — live re-run of the close-modified content-stories harness | **LAND-NOW batched** — the 29/29 headline is currently unit-proven, not live-re-proven |
| 9 | DEF-M226-01 — pre-bind serve reap | **LAND-NOW batched**, else KEEP-DEFERRED **with a named milestone** |
| 10 | BURNIN-M221-dev-public-host · F-M220-4 · PROBE-M218-c3-rerun | **LAND-NOW batched**, else fresh KEEP-DEFERRED-WITH-SIGNOFF → v2.6 + a **v2.5 carry section in `roadmap-vision.md`** (none exists) |
| 11 | `DEF-M235-03` — M204 assign-WRITE in-manifest TODO | **KEEP-DEFERRED-WITH-SIGNOFF** — honest since v2.0, declared out at M201, tracked in-manifest; but 3 releases old ⇒ needs a **fresh dated** sign-off + a `roadmap-vision.md` entry (M206/M207 futures) |
| 12 | DEF-M10-01 · DEF-M21-01 · CAVEAT-1 · M314b · M205 residual | **KEEP-DEFERRED-WITH-SIGNOFF** — already in `roadmap-vision.md`; re-confirm with today's date, no other action |

---

## 7. Applied changes

**None.** This audit is read-only by instruction. Phase 5 application is deliberately withheld:
every blocking item needs a user fate first, and the `RELEASE-SCOPE-DEFER:` decisions +
`roadmap-vision.md` entries that items 10–12 require must record the *user's* words, not this
audit's recommendation.

## 8. Blocking items (require an explicit user decision)

1. **The standing demo-stack test failures** (REPEAT + AGED_OUT) — recommendation LAND-NOW.
2. **DEF-M226-01** (AGED_OUT ×2, both triggers unrecorded until now).
3. **The three surviving v2.3 tail carries** — BURNIN-M221 · F-M220-4 · PROBE-M218-c3-rerun
   (AGED_OUT; destination release shipped without them; no v2.5 vision section exists).
4. **The anon academy catalog twin** (AGED_OUT across 3 milestones) + the `apps/web` `:5050`
   endpoint + CLOSE-D3, as one batch.
5. **`DEF-M235-03`** — needs a fresh dated escape-hatch sign-off, not a rubber-stamp.

---

`DEFERRALS: RED` · total 22 · repeat 1 · aged-out 8 (7 newly detected) · unhomed 3 · blocking 5 · `SEVERITY=blocker`

---
title: "Deferral Audit — M219 close (milestone scope)"
date: 2026-07-14
scope: milestone
invoked-by: close-milestone
---

## Verdict

**YELLOW**

**No repeat deferrals. No chronic patterns. No aged-out items in scope. Nothing blocks the merge.**

The headline: **all three items M218 routed INTO M219 (Fate 3) LANDED COMPLETELY** — none bounced. That is the
outcome the three-fate rule exists to produce, and it is worth stating plainly because the failure mode this audit
guards is exactly the opposite (an item that keeps being re-routed until it quietly becomes backlog).

M219's own four deferrals are all **first** deferrals, all **Fate 3**, all with the receiving milestone's
`overview.md` **actually edited** (not left as prose), and all with their **fences landed here reporting RED on
purpose** — so none of them can go green while still broken.

## Summary

| Category | Count |
|---|---|
| LAND-NOW (Fate 1) | **0** — nothing left to land; Phase 7 addressed every finding inline |
| LAND-NEXT (Fate 2 — already owned, confirm only) | **0** |
| LAND-NEXT (Fate 3 — annotate a sibling's `overview.md`) | **4** |
| DROP | 0 |
| KEEP-DEFERRED-WITH-SIGNOFF (escape hatch, cross-release) | **0** |
| Repeat deferrals | **0** |
| Aged-out | **0** (in milestone scope) |
| Chronic patterns | **0** |

## Inherited from M218 — the three routed INTO M219

| id | item | destination | outcome |
|----|------|-------------|---------|
| DEF-M218-08 | **F-11** — the BAPI fabricates `organization.public_metadata.eid` (`"org_eid_"+orgID`) instead of the roster's real org UUID | `FIX-M219-bapi-org-eid` | ✅ **LANDED** (S3). Gene `MembershipOrgIdentity/real-org-eid` **GREEN**; Go surface 97.2% → **100.0% overall / 100% critical, 27/27**. Gene **retained** as a permanent fence (D-M219-5). RED-proven pre-fix. |
| DEF-M218-09 | `expressrun` is **UNMEASURABLE** (no `@clerk/express` → rc=2, *no score*, read as a pass) | `TEST-M219-expressrun-dep-gate` | ✅ **LANDED** (S4). `alignctl` exit codes **split**: `3` = UNMEASURABLE vs `2` = REGRESSED, with a banner that cannot be mistaken for a pass (D-M219-7). RED-proven: collapsing the codes fails `TestExitCodes_UnmeasurableIsDistinctFromRegressed`. |
| DEF-M218-10 | the demo-patch live-clone **freshness gate SKIPS** when `stack-demo/next-web-app` is absent | `TEST-M219-freshness-gate-skips` | ✅ **LANDED** (S4). The silent skip now **speaks**; three unit tests that deferred to a *"live-verify gate that does not exist"* now report themselves as **coverage holes, not passes**. |

**Zero bounced.** No REPEAT group exists for this milestone.

## Still open, routed to M220/M221 (NOT M219's to land)

`DEF-M218-02` (Clerk telemetry) · `-03` (C-5 clerk-js `Timeout: 0`) · `-04` (ant-academy gets the REAL
`CLERK_SECRET_KEY`) · `-05` (F-7 `NEXT_PUBLIC_BACKEND_API_URL` blackhole) · `-06` (F-5 ad/analytics egress) ·
`-07` (C-3 cms/Directus 403s) · `-11` (`clerkenstein.md:3-4` header) · `DEF-M217-01` (pre-bind reap never
field-proven).

**Not aged out:** each was first deferred 2026-07-13/14, sits at **one** milestone of age, and its **destination
milestone (M220 / M221) has not closed**. No aging trigger fires. They are M220's and M221's business.

## M219's own deferrals — all Fate 3, all first-time

### DEF-M219-01 — studio-desk bounces the presenter OUT of the demo
- **Fate-1 feasible: NO.** `:19000` → 302 → `:13000/login`. Root cause is **demo-up wiring** (which session /
  cookie / origin studio-desk rejects), diagnosable only against a **running demo with a browser** — M220's
  domain, not a seeding/rendering question.
- **Fate 3 → M220** (D-M219-22). `m220-cue-sheet/overview.md` **item (j) edited**.
- **The honesty landed here:** the fence reports **RED** until M220 fixes it.

### DEF-M219-02 — the ant-academy keyless bounce, **escalated: it POISONS THE DEMO SESSION**
- **Fate-1 feasible: NO.** Same class — live browser iteration against the demo's Clerk wiring. M220 already
  carried item **(i)** for exactly this.
- **Fate 3 → M220** (D-M219-22). `m220-cue-sheet/overview.md` **item (i) edited — SEVERITY RAISED at the harden
  pass.** It does not merely render blank: its keyless Clerk returns `Set-Cookie __session=; Expires=1970` +
  `__client_uat=0; Domain=…` — **domain-wide**, and cookies are scoped by **HOST, not port**, so `:13077`
  clobbers next-web's session on `:13000`. **A presenter who clicks the academy link is LOGGED OUT of the demo**
  into `ERR_TOO_MANY_REDIRECTS`, and **every employee coverage sweep aborts** — so the employee vantage has **no
  runnable sweep** while this is live.
- **The honesty landed here:** the launcher now reads the **body** and fails loud (`SERVES BUT DOES NOT RENDER`);
  `ANT_ACADEMY_HOME_SECTION` requires an `AI Academy` marker + a 400-char floor. **Both RED until M220 — intended.**

### DEF-M219-03 — `GUARD-M221-host-isolation` (NEW, from the harden pass)
- **The battery-integrity finding.** Two agents can run cycles against **one** demo host and nothing stops them.
  M219's battery was corrupted by exactly this — two batteries run concurrently, one purged the stack
  mid-measurement, so **cycle 5's `no-junk-skills` gate went UNEXECUTED** (a FINDING, not a pass).
- **Fate-1 feasible: NO.** A host lock must be **proven on the host**, and M221 *is* the on-host milestone. Worse,
  it is a **prerequisite for M221's own gate**, which is itself a multi-cycle battery on that same single host —
  without it, M221 can corrupt its own evidence exactly as M219 did.
- **Fate 3 → M221.** `m221-prove-on-billion/overview.md` **edited** (§ *Inherited from M219*), with a DoD: a second
  concurrent cycle must **fail loud with the holder's identity** — never queue, never proceed.

### DEF-M219-04 — `FIX-M221-reap-native-academy` (NEW, from the harden pass)
- `down --purge` does not reap the **host-native** academy, so it holds `:13077` across cycles: the next bring-up's
  academy dies `EADDRINUSE` while **the OLD process keeps answering** — cycle N+1 silently measures cycle N.
- **Fate-1 feasible: NO.** Same surface as `DEF-M217-01` (the pre-bind reap, already M221's), and provable only on
  the host.
- **Fate 3 → M221.** `m221-prove-on-billion/overview.md` **edited**, noting the reap must cover the **native**
  processes (cockpit **and** academy), not only container ports.

## Out of milestone scope — flagged for `/developer-kit:close-release`

**The `assignment-monitoring.assign-and-track.UC1` in-manifest TODO** (the assign-WRITE half, a two-backend
org-admin write) — declared out of scope in **v2.0 M204** and still `unimplemented=1` in `ptreport`.

**Deliberately NOT fated here**, on the audit's own scoping rule (*"grep the source tree **restricted to files the
milestone touched**"*): `playthroughs/manifest/assignment-monitoring.yaml` is **not** in M219's diff, and the item
belongs to an unrelated product surface. Manufacturing a blocker out of it would be scope theatre.

It is also **not a hidden deferral**: the Playthroughs 4-state model treats `unimplemented` as a **first-class,
reported** state, and `ptreport` surfaces it on **every run**. That is the opposite of silent erosion.

**But it has now outlived the release it was declared in**, and M219 *did* touch the playthroughs suite
substantively (it added four). **`/developer-kit:close-release` (scope=release) should fate it explicitly** —
LAND-NOW in v2.3, or an explicit cross-release `KEEP-DEFERRED-WITH-SIGNOFF`. Recording it here so it cannot be
missed there.

## Applied changes

- `m220-cue-sheet/overview.md` — item (i) **severity raised** (session-poisoning + the two consequences); item (j)
  already present.
- `m221-prove-on-billion/overview.md` — new **§ Inherited from M219 (Fate-3)**: `GUARD-M221-host-isolation` +
  `FIX-M221-reap-native-academy`.
- `m219-readiness-renders/decisions.md` — D-M219-22 (the M220 routing) already recorded; D-M219-23/24 added at close
  for the M221 routings.

## Blocking items

**None.** Zero repeat deferrals, zero aged-out items in scope, zero escape-hatch entries.

**DEFERRALS: YELLOW** — 4 items · 0 LAND-NOW · 4 LAND-NEXT (all Fate-3, all target `overview.md`s edited) ·
0 DROP · 0 escape-hatch · 0 repeats. `SEVERITY=warning` — **non-blocking; the close proceeds.**

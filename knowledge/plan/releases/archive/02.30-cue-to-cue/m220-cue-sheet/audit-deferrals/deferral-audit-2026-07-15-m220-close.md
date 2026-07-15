---
title: "Deferral Audit — M220 close (milestone scope)"
date: 2026-07-15
scope: milestone
invoked-by: close-milestone
---

## Verdict

**YELLOW** — and it would have been **RED**. The one chronic repeat-deferral in the ledger was investigated
rather than re-deferred, and it **landed (Fate 1) in one line**. What remains are single deferrals with clear
Fate-3 destinations on M221, each with a stated DoD. **Zero blocking items. Zero escape-hatch deferrals.**

## Summary

| | |
|---|---|
| Total deferrals in scope | **17** |
| Single deferrals | 13 |
| **Repeat deferrals** | **1 group** (the `dev-stack` suite — 5 milestones) |
| Chronic patterns flagged | **1** → **RESOLVED by LAND-NOW this pass** |
| Blocking items remaining | **0** |
| Escape-hatch (cross-release) | **0** |

---

## 🔴 The one that mattered — a 5-milestone CHRONIC deferral, resolved in one line

### REPEAT / CHRONIC: *"the `dev-stack` suite cannot be run"*

- **First deferred:** v2.2, as *"the `dev-stack` suite fails on this box — **environmental** (needs a live
  Postgres on `:15432`; this box's `.agentspace/secrets` is also incomplete)"*.
- **Deferred again:** M217, M218, M219 — carried verbatim in `state.md` § Headline numbers as a **"Known issue
  (flagged, not a regression) … Identical at v2.2, M217, M218."**
- **Deferred again:** M220 S0–S2, re-characterised as **`FIX-M221-devstack-test-spin`** — *"BUSY-SPINS forever,
  8 min wall, 145 % CPU, `rc=124`"* — and written into **M221's `overview.md`**.
- **Deferred again:** M220 S7, re-characterised **a third time** as **`F-M220-1b`** — *"20 tests, 486 s, **19
  failures**, dialling a real Postgres on `:15432`; its stub seam doesn't hold."*
- **Time in limbo:** 5 milestones across 2 releases. Flag: **`DRIFT_DEFER`** — the reason was rewritten every
  single time, which is exactly why nobody ever looked *under* it. Each new characterisation felt like
  progress and licensed one more punt.

**Fate-1 investigation (this pass).** The sibling class `DevSetdress.run_sd` sets
`DEV_SETDRESS_USE_STUB_BINS=1`; **`DevSetdressLocalContent.run_sd` never did.** That flag is the only thing
`build_cli` (`dev-setdress.sh:114`) consults before deciding whether to honour the stubs in `$DEV_SETDRESS_BIN`
or do a **real `go build`**. Without it, the class's own docstring claim — *"so the bootstrap →
apply-structure(replay) → boot chain is exercised **offline**"* — was simply **false**: it compiled the real
`stackseed` and ran it against a **real Postgres that was never meant to be there**. The "145 % CPU" was the Go
compiler. The 19 failures were `connection refused`.

**It was one missing environment variable.** It is the **same defect class D29 already named one harness over**:
*a fix that landed in one of two sibling harnesses is a fix that did not sweep.* D29 swept the `HOME` isolation;
nobody swept the stub-bins flag.

**Verdict: LAND-NOW (Fate 1). Landed.**

| | before | after |
|---|---|---|
| `DevSetdressLocalContent` | 20 tests · **486 s** · **19 FAILURES** | **20 passed · 28 s** |
| the whole `dev-stack` suite | **`rc=124`, never reaches a summary** | **116 passed · 4 skipped · 127 s** |
| the whole **rext Python suite** | **`unittest discover` never completes** | **1208 tests · 1192 passed · 16 skipped · 0 FAILURES · 6 m 14 s** |

**This was going to block `/developer-kit:close-release`** (M221's own `overview.md` says so in as many words).
It no longer does. The v2.3 close can now run the suites whole — **for the first time in the release.**

> **The lesson, and it is the release's own:** *"environmental"* was a **status artifact that outlived the thing
> it described, and was then read as evidence** — **D17**, sitting in `state.md`'s headline numbers for four
> releases, re-authored by every milestone that touched it and investigated by none. An unrunnable suite must
> never be recorded as a known-and-accepted cost; **it is a FINDING.**

### Consequence: an M221 route is now STALE and must be corrected, not merely removed

`FIX-M221-devstack-test-spin` sits in **M221's `overview.md`** describing a busy-spin that **no longer exists**
and a diagnosis (**"it spins, it does not block on I/O"**) that was **already known to be wrong** at the M220 S7
close (F-M220-1: *"THE DIAGNOSIS WAS WRONG… it named one defect where there were two"*). Leaving it there would
hand M221 a **confident, precise, entirely false** work item — **D17 again, one milestone forward.** Corrected in
Phase 5.

---

## Deferral inventory + fate verdicts

### Resolved this pass (Fate 1 — LAND-NOW)

| id | item | verdict |
|---|---|---|
| **F-M220-1b** | `DevSetdressLocalContent` stub seam doesn't hold | ✅ **LANDED** (one env var; see above) |
| **FIX-M221-devstack-test-spin** | the `dev-stack` suite "busy-spins forever" | ✅ **DISCHARGED** — (a) fixed at M220 S7 (D29), (b) fixed here. **Route removed from M221.** |
| *(state.md known-issue)* | *"the dev-stack suite fails on this box — environmental"* | ✅ **RETRACTED** — it was never environmental |

### Fate 3 — annotate M221 (single deferrals, clear DoD, applied in Phase 5)

| id | item | why not Fate 1 |
|---|---|---|
| **F-M220-2** → `FIX-M221-academy-empty-catalog` | academy home renders **0 PATHS / 0 COURSES** (348 chars) while `[build-catalog]` finds **2,705** entries; the home reads the **local** catalog, which emits 0 | Needs an `ant-academy` content-pipeline investigation + probably a new rext-owned demo-patch. **M219's 400-char floor stays honestly RED and was NOT weakened.** |
| **F-M220-5** → `FIX-M221-academy-loopback-bind` | ant-academy binds **`*:13077`** and answers on the tailnet IP **even on a localhost demo** (`BIND_HOST=""` passes no `-H`; `next dev` defaults to `0.0.0.0`) | The **code** fix changes the very localhost path S3's HARD INVARIANT is fenced on. Bundling an unproven behaviour change into the milestone whose thesis is *"the docs must match the code"* is precisely wrong. **Docs corrected in-milestone (Fate 1).** |
| **F-M220-4** | `ant-academy.sh` is not re-runnable on a live public-host demo (serve holds the port; a standalone re-run bakes `localhost` URLs into a public demo) | Needs a live public-host demo to verify. Same surface as the two above. |
| **BURNIN-M221-dev-public-host** | the dev `--public-host` is **fenced, not live-proven** | The ladder + serve generator were proven live in S3/v2.2; the ~60 lines of net-new **dev** wiring are covered by the tripwire fence + 9 mutants only. Needs a box. |

### Fate 2 — already owned by M221 (confirm; no plan edit)

M217: pre-bind reap never run live · compose-range preflight · freshness preflight.
M218: `PROBE-M218-backend-api-url-twin` (F-7) · `PROBE-M218-c3-rerun` (C-3).
M219: `GUARD-M221-host-isolation` · `FIX-M221-reap-native-academy` · `REPROVE-M221-battery-at-final-code`.
v2.2 residual: **DEF-M215-02** (remote-VM snapshot cache).

All eight are already in M221's `overview.md` `In:` surface. **No edit required — confirmed covered.**

### Out of milestone scope (release-level; carried, not re-fated here)

`metrics-history.md` has no v2.2 row; no release-scope deferral audit exists for v2.1/v2.2 → **next
`/developer-kit:close-release`**. Older unscheduled items (DEF-M10-01, DEF-M21-01, CAVEAT-1, M314b) live in
`roadmap-vision.md` and are not in milestone scope.

---

## Applied changes (Phase 5)

1. **rext** — `dev-stack/tests/test_dev_stack.py`: `DevSetdressLocalContent.run_sd` sets
   `DEV_SETDRESS_USE_STUB_BINS=1` (the Fate-1 landing).
2. **`m221-prove-on-billion/overview.md`** — the stale `FIX-M221-devstack-test-spin` entry **replaced** with an
   honest discharge note; the **4 new Fate-3 routes** added under an *Inherited from M220* section.
3. **`m220-cue-sheet/decisions.md`** — **D31** records the chronic-deferral resolution + the retraction of the
   "environmental" claim.
4. **`state.md`** — the "Known issue: the dev-stack suite fails on this box — environmental" line is
   **retracted** (Phase 10).

## Blocking items (require user decision)

**None.** The single repeat/chronic group was resolved by a complete Fate-1 landing. No escape-hatch deferrals
exist in this milestone.

`DEFERRALS: YELLOW` · single=13 · repeat=1 (**resolved**) · chronic=1 (**resolved**) · blocking=0 ·
escape-hatch=0 · `SEVERITY=warning`

---
iter: 281
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-11
---

# iter-281 — make the control tree runnable, then attribute the RED

**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

## Active strategy reference

`TOK-08`. The residual RED is a **mechanical** question — a test either passes in company or it does
not — and the instrument that would enumerate it (a full run on a control tree) is the one that cannot
run. Fix the instrument, then census, exactly as TOK-08 prescribes.

## Step 0 — Re-survey (mandatory), and it changed the plan

Three targets were named for this iter. **Two are already closed**, found by re-verifying before working:

| named target | re-survey result |
|---|---|
| two map rows call a destroyed service `live-standalone` (`roadrunner`, `graphql-wundergraph`) | **CLOSED** by iter-280 `b5bd4e7e` — both prod cells now read `decommissioned`, each with a scope caveat |
| `services.sh` cites `arm E`, a fence that never existed | **CLOSED** by iter-280 — `stack-verify/lib/services.sh:164-168` now carries the retraction |
| `FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT` | **OPEN** — 4 sites in 2 files still resolve `/Users/marco/...` |

This is the third consecutive iter where a route list was stale. It is recorded, not just noted.

## Cluster / target identified

The residual **13 failures / 4 files** that left iter-280's section gate RED, and the open
`FIX-M257x-278` that **blocks the measurement needed to triage them**. iter-280 could not say whether
the 13 are pre-existing or its own, because the control-tree full run aborts at collection: the census
substrate resolves an absolute ROOT and does not survive cloning.

Order is forced, not chosen: **the instrument first**.

## Hypothesis

1. Deriving `ROOT` from `__file__` (the idiom already used by 14 sibling modules,
   `os.environ["ROSETTA_ROOT"] or Path(__file__).resolve().parents[4]`) makes the census substrate
   survive cloning, so a control-tree full run collects.
2. With the control tree runnable, the 13 become **attributable**: pre-existing, iter-280's, or
   full-run-interaction. The instrument-inside-its-subject reading (`test_suite_census` writing its
   probe into the directory it censuses) is the leading hypothesis for the largest group (7) and is
   **not assumed** — two cheap subset runs have already REFUTED two candidate causes (see progress.md).

## Expected lift

The section gate returns **GREEN** — or the residual is attributed with named evidence and the
un-attributable part is stated as such. **No ceiling is bumped and no growth is grandfathered**;
iter-280 set that standard by deleting growth rather than raising a ceiling.

## Phase plan

- **A.** Close `FIX-M257x-278` — derive every hardcoded absolute root in the census substrate; prove
  the derivation RED-first (a control clone that used to abort must now collect).
- **B.** Build the control PAIR at HEAD and run the whole stack-core section suite on it. This is the
  attribution instrument; it runs on a frozen copy so working-tree edits cannot contaminate it.
- **C.** Triage the residual against that run; fix the class, not the case.
- **D.** Re-run the whole section suite on the working tree; report the number that comes back.

## Escalation conditions

- If the residual proves to need a platform-repo edit → route forward, never edit.
- If a fix would require weakening a fence or bumping a ceiling → STOP and record; that is the defect,
  not the remedy.
- Clause 5 is **parked with the orchestrator by design** — no population reading is attempted here.

## Acceptable close-no-lift outcomes

A measured, named attribution of the 13 that shows they are **pre-existing and out of this milestone's
reach** would satisfy the iter even with the gate still RED — provided the evidence is a run, not an
argument.

---
iter: 16
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-01
---

# iter-16 — the two bring-up verdicts that grade themselves green

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`).
Step 1 of TOK-01 is *"unblock the gate's instrument"* and step 3 is *"land the fences, each watched going
RED, **before trusting any green**."* Both routed-forward findings below are instances of the second: a
bring-up step that reports a verdict it did not measure, in the instrument clause 1 is measured through.

## Step 0 — re-survey (mandatory)

Re-measured before targeting, and it **changed the target's meaning**:

- The hardening ledger routes RF-4 as a `dev-stack` finding. It is not scoped to dev.
  `demo-stack/up-injected.sh:212,2214,2263` **reuses `dev-stack/dev-setdress.sh` verbatim** via
  `--stack-type demo`. So the `*)` arm at `dev-setdress.sh:390-391` that turns any unclassified replay rc
  into `SNAP_SUMMARY=…skipped(error)` and returns 0 is **the exact code path that printed
  `directus=skipped(error)` on all three of clause 1's cold cycles** while the closing line said
  `set-dressed`. RF-4 is not "the clause-1 signature"; it is the clause-1 site.
- `up-injected.sh:2262` already wraps the call in `if ! …; then log "⚠ set-dressing did not fully
  complete…"`. That warning **has never been reachable for a replay error**, because the script exits 0.
  The caller-side handling the demo path needs already exists and is dead code.
- RF-1's twin asymmetry re-confirmed at source: `migrate-demo.sh:150-177` has `mig_fail` + output capture
  + `exit 1`; `migrate-dev.sh:95-106` classifies **every** non-zero atlas exit as
  `"had migration warnings (non-fatal — see atlas output)"` with the output discarded by
  `>/dev/null 2>&1`, and its absent-clone branch (`:100`) logs `✗` and `continue`s without recording a
  failure at all. The closing line at `:131` says `done — … the derived migration set applied.`

## Cluster / target identified

Two bring-up scripts whose closing verdict is independent of what they measured. Both are the milestone's
own dominant class (*a check that reports a state without measuring it*), both are in files this milestone
has already edited, and one of them is the mechanism behind clause 1's three compromised verdicts.

## Hypothesis

Making each script's **verdict and exit code a function of its own per-step outcomes** removes the false
green at its source, without making either pass fatal to a bring-up (the callers already have the
warn-and-continue handling; RF-4's is currently unreachable).

Explicitly NOT in scope: the cold ×3 re-prove of clause 1 (iter-17 — it needs these landed first, plus the
harden pass's `probe_directus_serves_content` in the consumed tag) and any clause-2 root cause.

## Expected lift

No movement on clause 2's metric — this iter is not aimed at it. The lift is on **clause 1's
re-provability**: after this iter a cold cycle that fails to replay content reports it in the bring-up
transcript's own verdict, so iter-17's three cycles are measured with an honest instrument rather than
re-running the original mistake with a better probe bolted on the side.

## Phase plan

Two planned lines (this is the iter's declared multi-step shape; the scope-creep tripwire counts a **third**
unplanned line):

1. **RF-4** — `dev-setdress.sh`: separate *documented degradation* (rc 4 unprovisioned / rc 5 cache-miss)
   from *unclassified error* (any other rc); the verdict word and the exit code follow. The seed floor must
   still run — that is the pass's stated contract and it is correct.
2. **RF-1** — `migrate-dev.sh`: port the demo twin's capture + classify + `mig_fail` + refuse-to-report-OK,
   including the absent-clone branch.

Then: fences watched going RED (mutation), full section suites vs the recorded baselines, tag, **push the
tag to origin**, re-pin `.agentspace/rext.tag` + the `stack-demo` consumption clone.

## Escalation conditions

- A second platform commit invalidating an alignment attempt → `re-scope-trigger` (occurrence 2 of 2).
- A section suite moving off its recorded baseline in a way this iter caused → user-blocker.

## Acceptable close-no-lift outcomes

If measurement shows either verdict is already honest through a path the ledger did not read (e.g. a caller
that inspects `SNAP_SUMMARY` textually), the finding is refuted and recorded as such — that is a complete
iter under this protocol, and the refutation is the deliverable.

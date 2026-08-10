---
iter: 257
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
controlling_strategy: TOK-08
---

# iter-257 — the workspace pin fights the canonical one

**Type:** tik · **Active strategy:** `TOK-08` (census the mechanical classes; stop sampling them).

## Cluster / target identified

`ROUTE-M257x-256-workspace-pin-is-not-the-canonical-pin`, opened by iter-256 and sharper than it
looked when it was routed.

`clone_pin_guard` (FENCE-M257x-iter222) fences **`rosetta-extensions/demo-stack/clones.pin.json`**, the
canonical pin, and iter-222 removed five phantom keys from it. But `ensure-clones.sh:203-210` seeds the
workspace copy **copy-if-absent**:

```sh
CANONICAL_PIN="$HERE/clones.pin.json"
if [ ! -f "$PIN_FILE" ] && [ -f "$CANONICAL_PIN" ]; then
```

so a workspace created before that fix **never** gets the correction. Measured on this box:
`stack-demo/clones.pin.json` still names **11** repos — the five phantoms `cms`, `jobsimulation`,
`storage`, `messenger`, `roadrunner` included — and all five directories are **present on disk** with
git checkouts, so the pin's entries are not inert.

**And iter-256 made it worse in a way worth naming:** the canonical pin now says `app` `3eaadae68`,
the workspace copy still says `ad9f3c498`. `DEMO_ADVANCE_CLONES=pinned` reads the **workspace** copy
(`ensure-clones.sh:181` `PIN_FILE="$DEMO/clones.pin.json"`), so on this box that mode would **check the
clones back out at the pre-advance shas** — undoing iter-256's advance while reporting `pinned`.
*A barrier that disagrees with the barrier it was copied from is worse than no barrier.*

## Hypothesis

The fence's **subject** is one file and the **mechanism's** subject is another, and nothing compares
them. The repair is a comparison arm plus a decision about what copy-if-absent should mean when the
canonical has moved — and the decision is not obviously "always overwrite", because a workspace pin an
operator edited deliberately is a legitimate declaration (`ensure-clones.sh:362` calls that state
`pinned` and its disagreement `pin-drift`).

## Pre-registered claims — sealed in this iter's FIRST commit, before any measurement

CLAIM and PREDICTION kept separate.

- **PR-1 — CLAIM: `clone_pin_guard` has an arm that reads the WORKSPACE copy.** PREDICTION:
  **REFUTED — zero arms**; it resolves the canonical path by default and its CLI positional is a
  caller-supplied path nothing routes to a workspace.
- **PR-2 — CLAIM: exactly one rext call site reads the workspace copy.** PREDICTION: **REFUTED — two
  or more** (`ensure-clones.sh` reads it, and at least one status/registry surface does too).
- **PR-3 — CLAIM: `DEMO_ADVANCE_CLONES=pinned` is the default, so the phantom checkout is ACTIVE on
  every bring-up.** PREDICTION: **REFUTED — the default is `0`** and the hazard is latent, which is
  why four releases have not tripped over it.
- **PR-4 — CLAIM: the workspace pin is the ONLY canonical rext artifact seeded into a stack
  copy-if-absent.** PREDICTION: **REFUTED — at least one more** artifact is seeded the same way.
- **PR-5 — CLAIM: `stack-demo/clones.lock.json`, the runtime record beside the pin, agrees with the
  clones as they are now.** PREDICTION: **REFUTED — it is stale** and still records pre-advance
  `behind` counts.

## Phase plan

- **Phase A** — census the readers: every rext site that reads a `clones.pin.json`, partitioned by
  which file it means. Grade PR-1…PR-4.
- **Phase B** — decide the semantics (copy-if-absent vs reconcile), and fence the comparison.
- **Phase C** — repair this box's workspace copy under the decided semantics.
- **Phase D** — grade the pre-registrations; re-run the guard family and the touched tests.

## Escalation conditions

If the right semantics turn out to need a `/demo-up` run to validate, route it — the box is not quiet
and iter-256 already booked the bring-up as unproven.

## Acceptable close-no-lift outcomes

If the workspace copy turns out to be read by nothing that can act on the phantoms, the finding is a
**bounded negative** and the deliverable is the fence that keeps it bounded — stated with the number,
not asserted.

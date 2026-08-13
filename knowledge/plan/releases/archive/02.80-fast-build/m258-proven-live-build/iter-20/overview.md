---
iteration_type: tik
status: in-flight
milestone: M258
iter: 20
opened: 2026-08-12
---

# iter-20 — the last RED fence, and the anchors no fence can see

**Type:** tik — under `TOK-01`, on the user's *"i want no debt"* ruling.

## Step 0 — re-survey before targeting

Re-measured at open:

- **`platform_predicate_guard` is the only RED fence on the box** (`rc=1`, one finding). Ten other
  corpus fences are `rc=0` after iter-18.
- The finding is `[G1 dead-token] docker-desktop-vm`, and iter-18 established it is a **false
  positive**: `docker-desktop-vm` is a **host** profile (`hostprofiles/*.json`, M255), not a compose
  profile. Pre-existing — both sites sit outside iter-18's diff.
- **`ROUTE-M258-iter18-app-row-anchors-are-at-2035f9a`** is open: the migration map's `app` row pins
  seven `app/main.go` wiring anchors at `2035f9a`, while `origin/main` is `c52dbc51e`. They pass
  **range-only** — no fence grades them — and iter-18 found a sibling anchor in that same row already
  landing on a closing brace.
- One stack up (`demo-4`), verified green; nothing here touches it.

## Active strategy reference

`TOK-01`. Both targets are corpus/tooling debt with no host-time cost, which is what makes them the
right pair for a session whose remaining measurements all need a quiet box.

## Cluster / target identified

The two items iter-18 routed rather than took, for the reason the tripwire exists. They are taken now
because the first is the **only RED fence left** and the second is the **class the fences cannot
see** — a citation that resolves, lands in range, and points at the wrong thing.

## Hypothesis

1. **G1** — adding a *domain* discriminator to `_PROSE_PROFILE`, symmetrical with the negation and
   ref-pin discriminators already there, removes the false positive **without** widening the detector:
   a mutation battery must still kill a detector that has been loosened.
2. **The `app` row** — re-resolving seven anchors at `c52dbc51e` makes the row's citations mean what
   they say. No fence will confirm it, so each one is checked by reading the line.

## Expected lift

- `platform_predicate_guard` → **`rc=0`**, with the G1 arm still firing on a genuinely dead compose
  token (proven, not assumed).
- The `app` row's seven anchors resolve to the constructs the prose names, at the ref the row states.
- rext tagged and **pushed**, tag verified on origin.

## Phase plan

- **A** — measure G1's false positive precisely; write the discriminator; RED-prove it with mutants
  (a real dead token must still fire, and a widened detector must die).
- **B** — re-resolve the `app` row's anchors at `c52dbc51e`, each verified by reading the target line.
- **C** — gates: the full fence set + the touched rext test modules, against iter-18's measured
  baseline of 46 pre-existing failures.
- **D** — tag, push, verify on origin, re-pin the declaration; close.

## Escalation conditions

- If the discriminator cannot be written without weakening G1's real arm, **do not write it** —
  report the fence as a known false positive with a waiver instead, and say why the loosening was
  refused. A detector that fails open is worse than a fence with a documented exception.
- Any new suite failure beyond iter-18's 46 must be attributed before close, not after.

## Acceptable close-no-lift outcomes

Refusing the G1 change with a written falsification is a complete iter: *"the repair would disarm the
arm that catches the real class"* is exactly the kind of result this protocol asks for.

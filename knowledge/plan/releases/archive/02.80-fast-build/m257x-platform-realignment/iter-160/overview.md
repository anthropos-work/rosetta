---
iter: 160
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-160 — the value side: a test literal the tree can already derive

**Active strategy reference:** [`TOK-08`](../decisions.md) — *census the mechanical classes.* TOK-08
directs working the classes **in descending measured size**. iter-159 split the spelling-pin class in two
and censused the larger half; this iter takes the half it exposed.

## Step 0 — re-survey before targeting

`FIX-M257x-iter159-value-side-subsignature-is-unfenced` was routed **one iter ago** and is current by
construction. It is the sharper of the two routes iter-159 produced: **3 of 7 confirmed instances** pin a
hand-written **value** rather than a haystack, and `§5` rule 71's own prescribed structural repair —
*derive the expectation from the same source the code derives from* — targets exactly this half, which
**nothing enumerates**.

## Phase 0d pre-flight — RUN, and it passed

The predicate depends on being able to **execute a derivation**, so the derivation was dry-run before the
plan settled rather than assumed:

```
platform_topology.default_profile(stack-demo/platform)  → 'core'
platform_topology.default_services(stack-demo/platform) → ['postgresql','redis','sentinel','backend','gotenberg']
```

The frozen literal in iter-155's confirmed instance is
`"postgresql redis sentinel backend gotenberg"` — **byte-for-byte what the tree derives.** The pre-flight
did not merely pass; it confirmed the hypothesis before a line was written.

## Cluster / target identified

Sub-signature **(b1) — the frozen literal**: a test supplies or expects a hand-written literal whose value
**the tree can already derive**. The repair is not a rewording; it is `import the derivation`.

## Hypothesis

**A test literal is a candidate iff its token set equals a value some non-test module derives.** Unlike
iter-159's haystack clause this one cannot be decided by reading — the derivation has to be **run** — and
that is the point: the instrument compares the frozen copy against the live value.

## Expected lift

The unfenced half of the class gets an instrument, proved against the labeled instance it was built from,
plus a stated population.

## Phase plan

- **A** — a registry of derivable values, each carrying its provenance.
- **B** — the census; **prove it fires on iter-155's confirmed instance** at its pre-repair commit.
- **C** — population + denominator.
- **D** — fence, incl. the anti-vacuity control this instrument specifically needs (below); gates; close.

## The trap this instrument must not fall into

**On a box with no platform clone the derivation yields nothing and the census reports a silent zero** —
`§9`'s exact failure mode, and this instrument is unusually exposed to it because its predicate depends on
an external checkout. It must exit `CANNOT RUN`, never `0`, when a derivation is unavailable. That is a
required Phase D control, not a nicety.

## Escalation conditions

- The predicate does not fire on iter-155's pre-repair instance → **REFUTED**; record and close on the
  falsification rather than weakening the predicate until it fires.
- The derivable registry cannot be built without a live stack → route forward; do not ship an instrument
  that only runs on one box.

## Acceptable close-no-lift outcomes

A measured refutation of the frozen-literal predicate, with the labeled-instance result.

## Explicitly OUT of scope (tripwire pre-declared)

Sub-signature **(b2)** — iter-157's `assertEqual(on_disk, registry)`, an over-strict *direction* rather
than a frozen literal — is **not** this iter's target and the instrument is expected to be blind to it.
Declared here in advance so the blind spot is a prediction, not an excuse. Sweeping either population is
also out of scope.

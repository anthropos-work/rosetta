---
iter: 64
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
---

# iter-64 — the map's eighth state

**Active strategy reference:** `TOK-05`, step 3 of its next-tik direction (*"the map's `mid-fold`
state + the storage row, with G6 fencing the split"*), routed forward from iters 59 → 62 → 63.

## Cluster / target identified

`DOC-M257x-iter59-storage-mid-fold`. The *measurement* landed at iter-59 and is G6-fenced in
`storage.md` as a two-sided record. What did **not** land is the map's own vocabulary: the protocol
states the gap in its own words —

> *"The seven-token vocabulary above has no token for **mid-fold**, so a split like storage's is
> currently recorded nowhere."*

So the map still calls `storage` `live-standalone` on both sides, which is the collapse §6 exists to
prevent: the config side reads *removed* (`STORAGE_RPC_ADDR` set by no compose file, absent from
`.env_example`, service moved to `profiles: [storage-legacy]`), the consumer side reads *live* (`app`
v1.366.0 reads it at `main.go:446`, `:524`, `:992` and two `cmd/` tools hard-require it).

## Hypothesis

Adding the eighth token and applying it to the one row that has it makes the half-landed fold
*visible in the map* rather than only in a service doc — and the map is the fenced artifact.

## Expected lift

- `mid-fold` in the map's §1 vocabulary, defined and cited.
- `storage`'s fresh-local-stack cell reading `mid-fold`, with both sides named.
- `platform_alignment_guard.py`'s assertion C vocabulary widened 7 → 8, its docstring corrected, and a
  test that the eighth token is accepted **and** that a made-up token is still refused.
- The protocol's *"has no token for mid-fold"* sentence retired — it describes a gap this iter closes.

## Phase plan

A. Re-derive the split from platform artifacts (not from `storage.md` — §5, adjudicate against
   artifacts, never against another document).
B. Land the token: map §1, the `storage` row, the guard, the protocol.
C. Test: the eighth token accepted, an invented token still refused, guard watched RED.
D. Re-measure: alignment guard + predicate guard + anchor + structure + index; the intra-corpus
   re-point (§5 rule 34) for any line-count change.

## Escalation conditions

- If the artifacts show the fold has *finished* since iter-59 (`app` no longer reads it), the token is
  still correct vocabulary but `storage` is not its instance — record that and say so.

## Acceptable close-no-lift outcomes

- The split has resolved in either direction since iter-59 → record the falsification, retire the
  routed item, no token needed.

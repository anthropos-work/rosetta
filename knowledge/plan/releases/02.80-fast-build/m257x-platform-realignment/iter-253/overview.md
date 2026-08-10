---
iter: 253
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: ROUTE-M257x-249-fresh-checkout-hostile-tests
---

# iter-253 — the fresh-checkout class, from a count to a named census

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* This route is
the one `TOK-08` itself named at iter-249's close: *"This should become a fence: the class is mechanically
decidable and is still being manufactured."*

## Step 0 — re-survey (mandatory)

`TOK-08`'s standing direction after iter-252 lists `ROUTE-M257x-249-fresh-checkout-hostile-tests` as
**open, and the largest thing that run found**. Re-surveyed before targeting:

- The route is **untouched** by iters 250–252 (they worked the *guard* layer of the same defect class;
  this is the *test* layer).
- Substrate re-counted on disk at rext `d739952`: **136 python test files** across 6 sections
  (`stack-core` 84 · `demo-stack` 35 · `stack-injection` 7 · `dev-stack` 5 · `stack-verify` 5).
- **51 of 136** name operator-local state by at least one of `stack-dev` / `stack-demo` / `.agentspace`.

Target is current and meaningful. No substitution.

## Cluster / target identified

iter-249 measured the class as **23 failures across 13 files** on a frozen clone and left it as a
**count with no names** — which its own `ROUTE-M257x-249-a-reading-must-name-its-failures` says is not a
reading. Meanwhile three of the fifteen holding files were authored *that week*, so the class is still
being manufactured with nothing watching.

## Hypothesis

The class is mechanically decidable, but **not statically at the file grain.** A file-level predicate
("names operator-local state AND carries no skip idiom") is already measured below and is the wrong
instrument — `test_toolchain_floor_guard` carried a `skipUnless` and failed anyway, because it declared
*half* a precondition. The decidable form is **dynamic and named**: run the suite against a frozen clone,
record the failing node-ids, and ratchet against a checked-in ledger so the class cannot grow silently.

## Expected lift

A census instrument + a named ledger for the class, derived from a real run at rosetta `e87daf3` /
rext `d739952`. Repair of the tranche is explicitly **out of this iter's planned scope** and routes forward.

## Phase plan

- **A** — seal pre-registrations (this commit).
- **B** — freeze a `git clone --local --shared` pair at both HEADs; run the `stack-core` section there;
  record the failing node-ids **by name**.
- **C** — control: run the same node-ids on the live tree; anything that fails on both is a real defect
  and leaves the class.
- **D** — ship the census + the ledger + tests; close.

## Pre-registrations (sealed in this iter's first commit, before any run)

Disclosed substrate for PR-1, measured before sealing: the **static** file-level predicate selects
**8 files total / 6 in `stack-core`** (`test_anchor_subject_census_m257x`, `test_buildbench`,
`test_gen_override_home_binds`, `test_guard_family`, `test_m257x_corpus_file_citations`,
`test_write_target_schema_fence`).

| # | claim | prediction |
|---|---|---|
| **PR-1** | the static file-level predicate reproduces the dynamic class (6 == 13 files) | **false** — it under-selects; the grain is the test function |
| **PR-2** | the frozen-pair run reproduces iter-249's **23 failures / 13 files** exactly | **false** — iters 250–252 both repaired and authored |
| **PR-3** | ≥ 1 fresh-checkout-hostile test was manufactured by iters 250–252 | **true** |
| **PR-4** | the 51-file grep candidate set is a strict superset of the failing files | **true** |
| **PR-5** | every node-id failing on the frozen pair **passes** on the live tree (0 real corpus defects) | **true** |

Three of five predict against the convenient answer; PR-2 and PR-1 predict against my own prior iter's
published numbers.

## Escalation conditions

- A failure that reproduces on **both** trees is a real corpus/tooling defect → route it out of this class
  immediately and name it; do not fold it into the ledger.
- If the frozen-pair run cannot complete, the iter closes `closed-no-lift` with the instrument's refusal
  recorded — a census that cannot run must say so, never estimate.

## Acceptable close-no-lift outcomes

- The class is measured to be **empty or near-empty** at current HEAD (iters 250–252 having repaired it
  incidentally) — that falsifies the route's premise and closes it, which is a complete iter.

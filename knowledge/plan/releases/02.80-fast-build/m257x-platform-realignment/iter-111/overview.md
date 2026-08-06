---
iter: 111
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: in-progress
opened: 2026-08-06
---

# iter-111 — TOK-07 step 0: the two hardening items with teeth

**Active strategy reference:** `TOK-07: enumerate the predicate, not the anchor` (milestone-root
`decisions.md`), **step 0**.

## Step 0 re-survey (mandatory, done before targeting)

`TOK-07` was authored **this session**, so its `Next-tik direction` cannot be stale — but the survey was
run anyway, on its own terms:

- **Is item 1 still open?** Yes. `--json` on 12 guards is still preceded by the provenance line;
  `TestKnownWeaknessJsonIsPolluted` still asserts the defect and the workaround; the four call sites
  still set `FENCE_PROVENANCE_STAMPED=1`.
- **Is item 2 still open?** Yes as *routed*, and the survey immediately widened it — see the escalation
  note below.

Target confirmed, unchanged.

## Cluster / target identified

The two items `TOK-07` step 0 names, and nothing else:

1. **`FIX-M257x-harden23-json-polluted-by-provenance-stamp`** — the routed **design decision**: where
   does provenance belong when the payload is machine-readable? The undocumented env var must stop
   being the mechanism that holds the suite green.
2. **`FIX-M257x-iter108-stackcore-suite-hangs`** — *either make the suite complete, or make every count
   state its invocation.*

Both are step 0 rather than later because **step 1 ships a tool whose entire output is machine-read**,
and a count it prints would inherit both defects: an unparseable document, and a total nobody can
reproduce.

## Hypothesis

- Item 1: the stdout-vs-stderr framing is a false dilemma. iter-105's rationale (*print FIRST so
  `run_one`'s `lines[-1]` is the guard's own summary; flush-left so `headline()` does not count it*) is
  about **order and shape**, not about the **stream** — so text mode can be preserved byte-for-byte
  while machine mode carries its provenance **inside the document**.
- Item 2: unknown at open. The routed claim is *"blocks indefinitely"*; the measurement behind it is
  *"12.6 s CPU over 3 m 43 s, frozen at 442 results."* Those are not the same statement, and this iter
  tests the second rather than assuming the first.

## Expected lift

**No `N` movement, and none is claimed** — this is a tooling step under a strategy that reads last.
The deliverable is: `--json` parseable on every guard that offers it, with no hidden env var; and a
statable answer about the suite, with its invocation named.

## Phase plan

Two planned lines (a tooling-iter's shape — the scope-creep tripwire counts unplanned lines):
1. the machine-mode fix + its controls;
2. the suite question, measured rather than inherited.

## Escalation conditions

- If the suite genuinely does not complete, that is a **route-forward with a named handler**, not a
  user-blocker — the alternative half of the routed instruction ("make every count state its
  invocation") is landable either way.
- If the design decision turns out to need a user ruling (e.g. it would weaken a gate clause), stop and
  surface. It does not: no gate clause reads a guard's `--json`.

## Acceptable close-no-lift outcomes

A measured refutation of either routed item's premise is a first-class outcome. In particular, if the
"hang" is not a hang, **say so and correct the record** — the milestone's own standard is that an
observation can be real while the inference from it is wrong.

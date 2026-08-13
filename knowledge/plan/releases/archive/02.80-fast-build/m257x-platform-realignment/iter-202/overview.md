---
iter: 202
milestone: M257x
iteration_type: tik
iter_shape: tooling
status: closed-fixed
opened: 2026-08-09
---

# iter-202 — materialize at a REF, then adjudicate the nineteen

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* This iter
works the class TOK-08 named first — **intra-corpus / corpus→source citation resolution** — at the one
place the census admits it cannot resolve honestly: every excerpt it hands an adjudicator is read from a
**stale working tree**.

**Step 0 — re-survey (mandatory).** Ran the census at corpus HEAD `b9f0833`, fence-tree
`e1b7345a1`: **3,089 tier-1 pairs · 949 materialize from the clone set · EXPOSURE 19** pairs whose cited
lines differ at `origin/main`, over **9 distinct sites**, all terraform in `jobsimulation` / `storage` /
`messenger`. Six clones behind their own fetched `origin/main` (`storage` by 20). The two iter-198 routes
are both still live and still unbuilt. Target confirmed — no substitution.

**Cluster / target identified.** iter-198 closed with two routes that are one problem seen from two
sides:

- `SURVEY-M257x-iter198-materialization-reads-the-working-tree-by-construction` — the *structural* half.
  `materialize()` resolves a candidate path and reads its bytes. There is no way to ask it for the bytes
  at a named ref, so **every** tier-1 excerpt is working-tree bytes, stale ones included.
- `SURVEY-M257x-iter198-the-nineteen-exposed-pairs-are-unadjudicated` — the *evidence* half. The 19 are
  identified, never resolved, and they sit in the exact M810 / `service_desired_count` region this
  milestone has already been wrong about twice (iter-122's two false DOES-NOT-SUPPORT verdicts; iter-123
  and iter-124's corrections).

The first blocks the second. iter-198 said so explicitly: resolving the 19 *"needs the clone set updated
or a `git show`-based materialization path that this iter did not build."* The clone set belongs to a
live demo stack this milestone may not touch — so the `git show` path is the only sanctioned one, and
**it needs no network**: each clone already carries a fetched `origin/main` ref that differs from its
checkout.

**Hypothesis.** Teaching `materialize()` to read at a ref turns the stale-substrate disclosure from a
caveat into a *choice*, and reading the 19 pairs at both refs side by side decides each one. The
dangerous failure shape is a **silent fallback** — a caller that asks for `origin/main` and is handed
working-tree bytes because the file was not under a clone. That must be impossible to receive without
being told.

**Expected lift.** Two planned lines (a tooling-iter's declared multi-step shape):

1. **Tooling** — `materialize(..., ref=…)` + a `--exposure-adjudicate` verb that prints, per exposed
   pair, the citing corpus unit and BOTH excerpts. Fenced with a mutation control and an anti-vacuity
   control that can actually fire.
2. **Use it** — read all 19 and classify each: does the corpus statement survive at `origin/main`?

**Phase plan.** A: build the ref path + CLI. B: tests (mutation + anti-vacuity + no-silent-fallback).
C: run it over the 19 and adjudicate. D: land corpus repairs for anything that does not survive, or
record the falsification. E: close.

**Escalation conditions.** If adjudication finds a corpus claim contradicted at `origin/main`, that is a
**Fate-1 repair in this iter** (it is exactly clause 5's subject). If it needs the platform's own history
read beyond these clones, route forward.

**Acceptable close-no-lift outcomes.** If all 19 turn out to be substrate-dependent in lines the corpus
statement does not rest on — i.e. the exposure is real but inert — that is a complete iter: the
falsification is *"the 19 are ref-dependent and claim-independent"*, and it is only sayable because the
instrument was built.

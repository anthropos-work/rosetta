---
milestone: M257x
iter: 06
iteration_type: tik
status: closed-fixed
---

# iter-06 — `REPOINT-M257x-jobsim-writes`

**Active strategy:** `TOK-01` ("instrument first, then follow"), step 2 → step 3. The mechanism half is
done (iter-02 derived the migration set; iter-05 derived the verifier's expectation). This iter pays the
**symptom** half: the writes themselves. TOK-01 explicitly forbade doing it in the other order.

## Step 0 — re-survey (done before targeting; 2026-07-31 17:31–17:40)

Re-measured against the **live `demo-1`** stack (16 containers up), not from the pre-computed note.

1. **The write surface is 7 files, and it is exactly the 7 failing seeders.** Measured by the schema
   *string literal* rather than by the SQL text, which is what a `CopyRows` call actually carries:
   `rg '"jobsimulation"' -g '*.go' -g '!*_test.go' stack-seeding/` → `activity.go` · `ai_readiness_funnel.go`
   · `content_stories.go` · `hiring_funnel.go` · `jobsim_sessions.go` · `persona_write.go` · `succession.go`.
2. **The precomputed mapping was short by one table.** It listed 11; the live write surface is **12** —
   `validation_check_results` (`content_stories.go:204`) was absent from it. Recorded because the
   pre-compute commit explicitly asked for verification rather than trust, and verification found something.
3. **All 12 targets exist in `public`; `jobsimulation` holds 0 tables.** Confirmed by
   `information_schema.tables` on demo-1. `public.sessions` does not exist (`D-M257x-1`), so `sessions` is
   the sole rename → `public.job_simulation_sessions`.
4. **The `cms` half of the same defect is live and observable, and it is NOT this iter's target.** The
   bring-up log carries `sim-embeddings replay skipped (rc=4) — the stack's "cms" schema is
   missing/empty`: `stack-snapshot/simembeddings` replays `cms.similarities` + 3 children, and those four
   tables have moved to `public` too. Routed forward as a **new** named handler (see below), because it is
   a different module with its own firewall/parent-scope semantics and needs its own live proof.

## Planned scope — a declared 3-step shape

The scope-creep tripwire counts against these three, not against a single-target tik.

1. **Re-point** all 12 write targets through **one declared mapping**, not 7 hand-edited files' worth of
   find-and-replace. A scattered replace would recreate the hand-maintained-list defect this milestone
   exists to end. Leave the history in a comment (§7 rule 3) and assert the positive replacement (§7 rule 2).
2. **Fence it**, mechanism-pinned and mutation-verified RED: the seeders must not be able to name a schema
   the platform no longer creates, and the fence must pin *where the target comes from*, not *what it
   happens to be* (§8 rule 3).
3. **Shrink the debt and re-measure live.** `REXT_TRANSITIONAL_SCHEMAS` goes `cms jobsimulation` → `cms`;
   its no-growth fence is designed to fail on a shrink with *"this failure is good news"*, so watch it go
   RED and update it deliberately. Then re-run the set-dress seed against demo-1.

## Hypothesis

The 12 targets are column-compatible with their `public` counterparts (the col sets look like subsets), so
re-pointing is sufficient and no row shape needs to change. **This is a hypothesis, not a finding** — Trap A
says a fidelity check against the wrong reference passes, and "same table name" does not entail "same
columns". It gets asserted mechanically against the live DB before it is believed.

## Expected lift

- `stackseed: 7 seeder(s) failed` → **0**, on the live stack.
- Gate **clause 4** moves from *measured-firing* to *paid-down-and-fenced* for the jobsimulation half.
- Downstream: the `hiring org UNDER-SET-DRESSED` autoverify ⚠ should clear, since iter-05 established it is
  a consequence of these 42P01s and not a separate defect. That would take autoverify **2 → 1**.

## Phase plan

Per `corpus/ops/platform-alignment.md`: §7 (re-point procedure) then §8 (fence, watched going RED), then a
live re-measure. Column fidelity is asserted against `information_schema.columns` on the migrated stack —
§8's "live schema assert" layer, which today asserts schema existence but not column existence.

## Escalation conditions

- A column set that is **not** a subset of its `public` target ⇒ the re-point is not 1:1; stop, record the
  drift table-by-table, and decide per-table rather than blanket-mapping.
- The session PAIR (`jobsimulation.sessions` + `public.job_simulation_sessions` written together by the
  hiring funnel) collapsing into a duplicate-key conflict once both halves point at the same table ⇒ a real
  design question, not a mechanical one.

## Acceptable close-no-lift outcomes

Finding that the mapping is **not** 1:1 — with the drift measured and written down — satisfies the protocol
even if the seed still fails, because the pre-computed mapping would then be refuted by measurement, which
is the more valuable output.

## Routes opened by Step 0 (Fate 3, named)

| item | why | target |
|---|---|---|
| `REPOINT-M257x-cms-similarity-writes` | `stack-snapshot/simembeddings` replays `cms.similarities` + `similarity_{categories,features,skills}`; all four now live in `public` and the `cms` schema is an empty shell. Already firing on demo-1 as `sim-embeddings replay skipped (rc=4)`. Blocks `REXT_TRANSITIONAL_SCHEMAS` reaching empty. | next tik |

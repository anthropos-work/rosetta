---
iter: 241
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-241 — the clone set is a MEASUREMENT INSTRUMENT, and nobody stated its reach

**Active strategy reference:** `TOK-08`.
**Route worked:** `ROUTE-M257x-236-host-is-the-unreliable-witness` — open since iter-236, never worked.

## Step 0 — re-survey

iter-240's close pointed at the next thing while grading its own exclusion: `app`/`sentinel` are readable
here **because this box has them**. Re-surveying the workspace makes the general form measurable.

`stack-demo/` holds **13** git clones. `repos.yml` names **4** (`app`, `sentinel`, `next-web-app`,
`studio-desk`), and `clone_pin_guard` accepts **2** sanctioned extras (`ant-academy`, `platform`), plus the
`rosetta-extensions` consumption clone — **7 accounted for.** The remaining **6** — `cms`,
`graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage` — are exactly the repos
`repos.yml` says in its own header comment are frozen legacy that *"`make init` therefore does not clone…
clone them by hand if you need to read the pre-merge source."*

**They are leftovers from before `838d907`, and a fresh box does not have them.**

That matters because `platform_alignment_guard.py:480` derives its unclonable set from
`not (clones_root / head).is_dir()` — **disk presence on the host that runs it.** Its shipped verdict reads
*"11 citation(s) into 2 repo(s) the map documents but no stack clones (db-backup, infrastructure)"*, and
that **2** is a property of this laptop, not of the corpus. The milestone's gate clause 3 asks for claims
*"machine-fenced against `repos.yml`"*; a fence whose reach silently shrinks on a clean machine is fenced
here and unfenced there.

## Hypothesis

The corpus's most load-bearing platform claims — the whole merge/teardown map — are cited into repos a
**fresh** stack does not clone, so on a fresh box those citations flip from **verified** to
**excused-as-unreadable** without any verdict changing colour. `§9` iter-178/iter-208: *a NOT-COVERED
clause is a MEASUREMENT or it is a mood*, and *it names a RUN and gets read as a PROPERTY.*

## Pre-registered claims — SEALED IN THIS COMMIT

- **`P-241-1`.** `stack-demo` holds exactly **6** clones unaccounted for by `repos.yml` + the sanctioned
  extras + `rosetta-extensions`: `cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`,
  `storage`. **Predict 6, those names.**
- **`P-241-2`.** Against a clone set restricted to what a fresh bring-up creates, the alignment guard's
  *"repo(s) the map documents but no stack clones"* count rises **2 → 8**. **Predict 8.**
- **`P-241-3`.** The corpus carries **≥ 20** citations into those 6 repos. **Predict ≥ 20.**
- **`P-241-4`.** The guard's GREEN sentence names the *count* of unclonable repos but **not** the clone set
  it read, so the verdict is not reproducible from its own words. **Predict: the clone set is absent from
  the verdict.**
- **`P-241-5`.** `ensure-clones.sh` does not clone any of the 6. **Predict: confirmed, 0 of 6.**

## Phase plan

1. Seal this pre-registration.
2. Measure `P-241-1`…`P-241-5`, each against the artifact, never against a doc quoting it.
3. Deliverable: make the reach **stated in the verdict** — the excused repos named, and the verified /
   excused citation split printed — so a green cannot be read as "everything was checked".
4. Regression-test the reach statement, including the fresh-box case.

## Escalation conditions

If the fresh-box reading shows a guard would go **green while checking materially less**, that is a
gate-clause-3 disclosure and it is written down in the corpus, not only in the guard.

## Acceptable close-no-lift outcomes

If the guard already states its clone set somewhere a reader of the verdict sees, `P-241-4` is refuted and
the iter closes on the disclosure measurement alone.

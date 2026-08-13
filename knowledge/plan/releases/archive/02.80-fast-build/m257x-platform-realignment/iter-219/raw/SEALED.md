# iter-219 — sealed BEFORE the by-effect run

## V1 — three static heuristics, three different answers, none of them a population

Measured at rosetta `d0979e7` / rext `5635f88`, over `stack-core/tests`:

| heuristic | answer |
|---|---|
| test modules containing any `.write_text(` | **54** of 77 |
| `write_text` calls whose receiver expression does not look temp-ish | **369** |
| `write_text` calls with a `.py` path in the surrounding 14 lines | **74** |

**Only 3 of the 54 handle mtime at all** (`os.utime`/`touch`/`st_mtime`):
`test_frozen_expectation_census_m257x`, `test_mutation_proof_cache_hazard_m257x`,
`test_suite_census_collection`.

The three answers are **not three readings of one population** — they are three different questions,
and none of them is *"which writes can the hazard actually reach."* Sealed as a refutation of the
static approach before the by-effect recorder is written, so the pivot cannot be narrated afterwards
as the plan.

## V2 — the by-effect method, fixed before it runs

Patch the write primitives (`Path.write_text`, `Path.write_bytes`, `builtins.open` in a write mode),
run the whole `stack-core` suite, and record per write: absolute path · inside-repo · existed-before ·
old and new size · mtime before and after. **Exposure** is then decidable rather than inferred:

> a write is EXPOSED iff its target is a `.py` **inside the repo** that the running interpreter may
> re-read, the write is **size-preserving**, and the mtime **is not forced** to a fresh value.

## V3 — pre-registered stop condition

The recorder must not change what the suite does. **The instrumented run's pass/fail/skip counts must
equal the uninstrumented run's** (`1,864 passed · 1 failed · 3 skipped` at iter-217, and this iter's own
uninstrumented control). A recorder that perturbs its subject is refused, not corrected.

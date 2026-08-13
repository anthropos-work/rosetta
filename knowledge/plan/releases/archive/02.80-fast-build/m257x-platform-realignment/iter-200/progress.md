**Type:** tik — under `TOK-08`.

# iter-200 — rule 77's hazard is not live here, and rule 77 is not the whole hazard

## The reading

Three censuses in `derivation_registry`, over the whole of `rosetta-extensions`:

| census | population | split |
|---|---|---|
| `mutation_rewrite_sites` — a repeated write to ONE target inside one function | **25** | **`data` 25 · `py-no-utime` 0 · `py-utime` 0** |
| `py_writing_tests` — every test function writing a `.py` path at all (the upper bound) | **35** | `py-no-utime` 35 |
| `memoised_disk_readers` — `path → content` memos in non-test modules | **14** | `module-dict` 10 · `lru` 4 |

## Half one — rule 77's hazard has no live instance, and the reason is not the one you would guess

`§5` rule 77's hazard needs **re-imported Python source**: CPython invalidates cached bytecode on
`(mtime-seconds, size)`, so a same-length edit inside the same second is served from cache. **A `.md`,
`.go`, `.yml` or `.txt` has no bytecode at all.**

Every one of the **25** repeated in-place rewrites in this repo targets a data file — `tgt` (markdown),
`main.go`, `docker-compose.yml`, `f.txt`, `doc`. **Zero rewrite Python source.** So `h42`'s question has
a structural answer rather than a probabilistic one: the mutation proofs it was worried about could not
have been vacuous *in the way rule 77 describes*, because none of them mutates a thing that gets
compiled.

**The repeat predicate's blind spot is stated, not left implied.** A helper that stages the file and a
test that mutates it are two different target expressions, so a stage-then-mutate split is invisible to
it — and its `py` zero would then read as *"no test touches Python source"*, which is **false**:
**35 test functions write a `.py` path**, and **none bumps mtime**. Four of them are the `_stage` helpers
of the mutation batteries. Those are safe by a different mechanism — each mutation gets a **fresh tmp
directory**, so no cache entry from a prior case can be hit — but **safe by isolation is not safe by
discipline, and nothing in those files says so.** An optimisation that reused one directory would arm
the hazard silently. That is the residual, and it is routed rather than described away.

## Half two — the hazard rule 77 does not name, and it is strictly broader

A mutation control is vacuous whenever the second read **does not reach the disk**. Bytecode is one
route. **In-process memoisation is another, and it ignores mtime *and* size *and* file type.**

**14 memo sites across 4 guards** — `_git_out` (twice, in `anchor_construct_guard` and
`platform_alignment_guard`), `service_repo_map`, `tracked_basenames`, `RESOLVE_ROUTES`, `NOT_CITATIONS`,
`_SUBSTRATE`, `_BASENAME_INDEX`, `_HISTORY_CACHE`, `_REF_RESOLVES_CACHE`. Every one is keyed on a path or
a root and returns disk-derived content.

They are **deliberate and load-bearing** — `anchor_construct_guard`'s own docstring records that removing
one took the guard from ~1 s to 10.9 s — so the fence **enumerates** them and does not assert them to
zero. What the enumeration buys is that a future mutation control written against one of those guards
meets a list instead of a surprise. Measured while surveying: **exactly one test module in the repo ever
calls `cache_clear()`**, and it clears two of the fourteen.

**The two halves point opposite ways, and that is the finding.** Bumping an mtime — rule 77's whole
prescription — does nothing whatsoever for the memo hazard, which has more surface here than the one the
rule names.

## Close — 2026-08-09

**Outcome:** `SURVEY-M257x-h42-…` asked whether earlier size-preserving mutation proofs were vacuous, and
was careful not to claim they were. Answered with a census, and the answer has two halves. **(1)** Rule
77's hazard needs re-imported Python source; all **25** repeated in-place rewrites in this repo target
data files (`.md` / `.go` / `.yml` / `.txt`), so the hazard is **structurally absent**, not merely
unobserved — with the predicate's blind spot bounded from above (**35** test functions write a `.py`
path, none bumps mtime, four are battery stagers that are safe **by fresh-directory isolation rather
than by any stated discipline**). **(2)** Rule 77 is **not the whole hazard**: **14 in-process
`path → content` memos across 4 guards** make a mutation control vacuous regardless of mtime, size or
file type, and **one test module in the repo has ever cleared one**. Bumping an mtime does nothing for
the larger surface.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirty-second consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted:** iters 197, 198, 199, 200 = **four** tiks this run; the fifth is
available — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-200-1` … `D-M257x-200-3` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **58 passed** in
`test_frozen_expectation_census_m257x.py` (**52 → 58**, +6 arms this iter), green under **both** runners
(unittest 3.9.6: `Ran 58 tests … OK`); **33 passed** in `test_fence_registry_population_m257x.py` +
`test_derived_count_guard.py`. `route_disposition_guard` **OK** after the pre-iter repair. *Scope:
`stack-core` only, Python only, changed-code reach (`§5` r60) — no Go, no TypeScript, whole-section
figure not re-taken.*

**Side-deliverables:**
- `fix(M257x/199)` — `route_disposition_guard` was **RED at HEAD** because iters 198 and 199 abbreviated
  a route id with an ellipsis. Repaired before iter-200 opened, both ids spelled in full, and the third
  site of iter-199's own unmeasured scope figure corrected with it.

**Routes carried forward:**
- `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — **CLOSED.** Answered by census in both
  directions, with the exposed shape at zero, the repeat predicate's blind spot bounded from above, and
  the broader hazard enumerated rather than asserted away.
- `SURVEY-M257x-iter200-battery-stagers-are-safe-by-isolation-not-by-discipline` — **NEW.** The four
  `_stage` helpers avoid every cache hazard because each mutation gets a fresh tmp directory. Nothing in
  those modules states that this is load-bearing, so a directory-reuse optimisation would arm the hazard
  with no fence in the way.
- `SURVEY-M257x-iter200-only-one-test-module-ever-clears-a-memo` — **NEW.** 14 memo sites; exactly one
  test module calls `cache_clear()`, covering two of them. `_SUBSTRATE` and the other module-dict memos
  have no clearing API at all — a test must reach into a private global.
- Unchanged and still open: `FIX-M257x-h44-claim-census-guard-is-single-runner` ·
  `SURVEY-M257x-iter199-the-literal-census-reads-PRINTS-only` ·
  `SURVEY-M257x-iter199-the-noun-list-is-a-declared-vocabulary` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter198-the-nineteen-exposed-pairs-are-unadjudicated` ·
  `SURVEY-M257x-iter198-materialization-reads-the-working-tree-by-construction` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **A rule names a mechanism, not a hazard class.** Rule 77 is correct and complete about bytecode, and
  bytecode turned out to be the *smaller* of the two ways a mutation control can fail to re-read here.
  Auditing a rule's compliance is not the same as auditing what the rule is for.
- **"Safe" and "safe for a stated reason" are different states.** The batteries pass; nothing records
  that fresh-directory isolation is why, so the property is one refactor from gone.
- **Zero is only readable next to its predicate's blind spot.** The `py` zero would have read as a much
  stronger claim than it is, until the 35-site upper bound was put beside it.

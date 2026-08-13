**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.*

# iter-221 — the corpus cites 2,117 files and nothing checked that any of them exist

## The finding

iter-220 landed direction B over `rosetta-extensions`' READMEs. The same construct **on the surface
clause 5 is actually about** was unfenced: `corpus_citation_guard` grades markdown **links** (C1
resolution, C2 anchors), and **a backticked filename is not a link.** Measured over
`fence_provenance.corpus_sources`' **114** documents: **2,117** file citations —
`.md` 876 · `.go` 428 · `.sh` 173 · `.py` 149 · `.yml` 133 · `.ts` 112 · `.json` 100 · `.yaml` 75 ·
`.tf` 39 · `.tsx` 32. **Nothing read any of them.**

The `.md` half is fenced here **and only that half, on purpose**: every pool a corpus `.md` citation can
resolve against is on this box, while `.go`/`.ts`/`.tf` citations point into service repos that may or
may not be cloned — and *"I could not check it"* and *"it is correct"* are different findings of which
only one is coverage. Widening needs an UNMEASURED bucket keyed on clone presence: **routed, not
assumed.**

## The residual is a TAXONOMY, not a defect list — 19 citations, 0 defects

| class | members |
|---|---|
| **negated** — the corpus cites the file to record its ABSENCE | `guidance.md` (`studio-room.md:467`: *"…nor `guidance.md` … anywhere on disk"*) |
| **explicitly future** | `deploy_guide.md`, `debug_guide.md` — both under `corpus/ops/README.md`'s *"## Future Operations — This directory may grow to include"* |
| **cross-repo, not cloned** | the six `07*-*.md` of `anthropos-knowledge-base`; `ant-singularity/…` ×3; `reference_devserver.md` (`kb-ant-business`, and the citing sentence calls it stale) |
| **git-ignored workspace artefact** | `.agentspace/profile_gaps.md`, `.agentspace/seeding_gaps.md`, `stack-dev/setup_progress.md` ×3, an ops report |

The **negated** class is why this could not be written by pattern alone: a census that flagged it would
be telling the corpus to stop saying true things. iter-214 refused an entire widening over the same
shape. Declarations are reconciled **both ways** — an undeclared non-resolver fails, and a declaration
whose citation has vanished fails too.

## ⚠️ The pool scope was wrong TWICE more, and it is now an argument the control varies

| pool | `.md` citations resolving nowhere |
|---|---|
| rosetta only | **45** (38 distinct) |
| rosetta + `rosetta-extensions` + the `stack-demo` clone set | **19** (17 distinct) |

The first reading was **2.4×** the second and every citation it lost was a real file in a pool the probe
had excluded — the **third and fourth** occurrences of that class inside two iters, the later ones by
the session that had just kept iter-220's narrow-scope RED as a control against precisely this.
*Remember the scope* has now failed four times, so `test_03` **varies the pool** and asserts the
sensitivity, rather than restating the lesson in a docstring.

## ⚠️ A shipped fence went RED on this session's own writing — and the root cause outlives the fix

`route_disposition_guard` failed the run: **`SURVEY-M257x-h42` was CLOSED at iter-200** — *"answered by
census in both directions, with the exposed shape at zero"* — and **iter-219 published it open nineteen
iters later.**

**The premise came from `hardening-ledger.md`, which re-listed h42 as *"routed forward"* in passes 48,
49, 50 and 53 — every one of them after the closure.** The ledger is a **second disposition surface that
the route registry does not read**, so a closed route went on being published as open until an iter
spent its opening claim on it.

Corrected in place at iter-219 in the registry's own grammar, appended and not substituted; the guard
now reports **57 closures · 0 contradictions**. What iter-219 actually produced is **not** withdrawn —
it is an **independent reproduction of iter-200's answer by a different instrument** (exposed shape
**0**) plus a new obstacle belonging to that instrument. *A retraction that reaches the prose and not
the code has not landed* — here the reverse: the closure reached the registry and not the ledger.

## Scope, stated rather than implied (`§5` r60)

`/usr/bin/python3 -m pytest` (**pytest 8.4.2 / CPython 3.9.6**), **Python**, `stack-core` only,
changed-code reach: **179 passed / 0 failed** across the new module, iter-220's module,
`test_fence_provenance` (RED before the correction, green after), `test_route_disposition_guard` and the
frozen-expectation census (2 m 34 s). `route_disposition_guard --repo-root` exits **0**.
`--ceilings` exits **0**, all three `exact +0` after re-pinning **563 → 566**. No whole-section run —
the tree was edited throughout. No Go, no TypeScript, no non-`stack-core` Python section.

## Close — 2026-08-09

**Outcome:** the corpus's own file citations — 2,117 of them, in no fence's population — are now
enumerated for the `.md` half, with all 19 non-resolvers adjudicated into four declared classes and
**zero defects**. A shipped fence caught this session's own false premise; the root cause is that the
hardening ledger is a second, ungraded disposition surface, and it is routed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
**the RED was this session's own prose, repaired in-iter in the registry's grammar; the gates are green
at close** — (5) cap-reached: **y** — **counted, not felt: iters 217, 218, 219, 220, 221 = five tiks
this run against a cap of five** — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **exit-5**
**Decisions:** `D-M257x-221-1` … `D-M257x-221-3` (see [`decisions.md`](decisions.md))

**Side-deliverables:** the marked correction to iter-219's routes block — recorded separately because it
repairs a *prior* iter's claim, not this iter's planned scope.

**Routes carried forward:**
- `SURVEY-M257x-iter221-the-hardening-ledger-is-a-second-ungraded-route-surface` — **NEW, and it cost an
  iter's premise this run.** `route_disposition_guard` reads iter routes-blocks; the ledger's *"routed
  forward"* lists are outside its population and contradicted it four times.
- `SURVEY-M257x-iter221-non-markdown-citations-need-an-UNMEASURED-bucket` — **NEW.** 1,241 of the 2,117
  citations name non-markdown files in repos that may not be cloned.
- `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` — unchanged.
- All routes from iters 207–220 unchanged, plus the standing queue.

**Lessons:**
- **A registry with two disposition surfaces has none.** The closure reached the route registry and not
  the ledger, and the ledger is what the next brief is built from.
- **When a lesson has failed four times, stop writing it down and make it an argument.** The pool is
  now varied by a control instead of remembered by a reader.
- **Fence the half you can actually resolve, and route the half you cannot** — a census that guessed at
  uncloned repos would report absence as correctness.

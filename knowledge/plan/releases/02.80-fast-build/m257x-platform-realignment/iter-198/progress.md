**Type:** tik — under `TOK-08`.

# iter-198 — the staleness table warned about a substrate; the reader needed a number about the evidence

## The reading

`claim_census_guard`, this tree, `stack-demo` clone set, **4.2 s** total (the exposure pass adds ~2 s):

| | value |
|---|---|
| clones behind their own fetched `origin/main` | **6** — `cms` 2, `jobsimulation` 4, `messenger` 7, `next-web-app` 4, `rosetta-extensions` 1, `storage` 20 |
| files differing HEAD..origin/main across those six | **30** |
| tier-1 pairs total | **3,089** |
| … that resolve to a file | **1,303** |
| … that materialize **from the clone set** | **949** |
| … sitting on a **drifted file** | **20** |
| … whose **cited lines actually differ** at origin | **19** |
| … cited lines unchanged despite the file drifting | **1** |
| … citing a file **absent** at origin | **0** |

**The warning implied 949 and the answer is 19.** Every one of those 19 is a **terraform** citation:
`jobsimulation/terraform/main.tf` (`:15-22`, `:15-40`, whole-file), `storage/terraform/main.tf` (`:9-11`,
`:13`, whole-file) and `storage/terraform/storage.tf` (`:22-25`, `:24-38`), plus
`messenger/terraform/main.tf:29`. Nine distinct sites.

**That is not a diffuse exposure — it lands exactly on the milestone's most contested claims.** Those are
the anchors behind `service_desired_count = 0`, the cms/jobsimulation/messenger M810 questions, and the
`module.messenger_euwest1` orphan — the same region where iter-122 booked two false verdicts, iter-123
cloned `infrastructure` to settle it, and iter-124 corrected the corpus. The pairs whose bytes depend on
the substrate are the pairs this milestone has already been wrong about twice.

## The direction — retracted, not extended

`KNOWN_WEAKNESS` clause (5) and `substrate_of`'s docstring both said a stale substrate *"produces evidence
**AGAINST** a claim that is TRUE"*. True, measured, and **one of two directions**.

The mirror is equally reachable and structurally *more likely*: where the corpus states something that
**was** true and no longer is, a stale clone **CONFIRMS** it — a false GREEN. Corpus and clone fell behind
**together**, so agreeing-and-both-wrong is the **correlated** case, not the exotic one. Nothing in this
family can say which case a given pair is in, because that requires adjudicating the claim, and `F4` is
the standing statement that it does not. So the honest form is: **substrate-dependent, direction
undetermined, and here are the nineteen.** Both publishing sites now say that, and the retraction is
named in place.

## The verdict consequence — derived, and the first draft of it was false

The batch's live rule is *an instrument that states its own invalidity must not exit 0*. Whether this
guard is in that class is a question, and it was answered the wrong way first.

**Draft:** *"the exit code grades the tier-2 ratchet, computed from the corpus text and not from any
clone's contents."* **False.** `census` filters every tier-2 candidate through
`has_subject_token(sent, _live_names(root, clones_root))`, and `_live_names` reads the clone **directory
names** plus `platform/repos.yml` and `platform/docker-compose.yml`. The exit code **is**
substrate-dependent.

Replaced with a derivation: `NAME_SOURCE_FILES` names those two files, and the report **checks** whether
either has drifted. Today neither has (`platform` is not in the stale set, and directory names do not move
with commits), so the report prints the not-tainted branch **with the check named**; if one ever does, it
prints the ⚠ branch and says the verdict is substrate-dependent too. The claim that would have been prose
is now a condition — *checked, not assumed* — which is the only form that can go wrong out loud.

The exit code is deliberately **not** changed. Recorded as `D-M257x-198-4`.

## Close — 2026-08-09

**Outcome:** `SURVEY-M257x-h46-…` offered *declare the direction, or grade it*; both were done, and both
found something. **Graded:** the six-clone staleness warning, which reads as a caveat over the **949**
pairs materialized from the clone set, is really about **19 of 3,089** — measured at cited-line grain,
with the `same`/`absent` buckets kept so the headline cannot be over-quoted — and all 19 are terraform
citations in `jobsimulation` / `storage` / `messenger`, the exact region this milestone has already been
wrong about twice. **Declared:** the module asserted the false-RED direction at two sites; the false-GREEN
mirror is equally reachable and *more* likely, since corpus and clone fell behind together — retracted in
place. The verdict-scope statement was **false in its first draft** (tier 2 does read the clone set, via
`_live_names`) and is now a derived check over `NAME_SOURCE_FILES` rather than a sentence.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (thirtieth consecutive `closed-fixed`; **no
`P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted:** iters 197, 198 = **two** tiks this run — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-198-1` … `D-M257x-198-5` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **43 passed** across the three
`claim_census*` modules (including the net-new `test_claim_census_substrate_m257x.py`, **9 arms, green
under BOTH runners**), and **45 passed** in `test_guard_family.py`, which drives the guard end-to-end.
*Scope: `stack-core` only, Python only, changed-code reach (`§5` r60). The whole-section figure is not
re-taken here; the other 10 sections, the 6 Go sections and the 424 TypeScript tests were not run.*

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-h46-stale-substrate-direction-undeclared` — **CLOSED.** Both branches of its either/or
  are landed: the direction is declared UNDETERMINED at both publishing sites with the single-direction
  wording retracted, and the exposure is graded at pair grain with an offline fence.
- `SURVEY-M257x-iter198-the-nineteen-exposed-pairs-are-unadjudicated` — **NEW.** The 19 are *identified*,
  not *resolved*. Each is substrate-dependent in an undetermined direction, and the corpus statements
  they support are exactly the M810 / `service_desired_count` region. Resolving them means reading those
  files at `origin/main` and re-adjudicating — judgement work, outside `F4`, and it needs the clone set
  updated or a `git show`-based materialization path that this iter did not build.
- `SURVEY-M257x-iter198-materialization-reads-the-working-tree-by-construction` — **NEW.** The exposure
  measurement proves the drift exists; it does not remove it. `materialize` still reads working-tree
  bytes, so every excerpt an adjudicator sees is still the stale one. A `--ref origin/main`
  materialization mode is the structural repair and was not attempted.
- Unchanged and still open: `SURVEY-M257x-h42-…` · `SURVEY-M257x-h45-printed-measurement-literals-
  uncensused` · `FIX-M257x-h44-claim-census-guard-is-single-runner` ·
  `SURVEY-M257x-iter197-the-derivation-registry-sees-only-NAME-shaped-derivations` ·
  `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` · and the standing queue.

**Lessons:**
- **A warning about a SUBSTRATE is not a measurement of the EVIDENCE.** "Six clones are behind" reads as
  a caveat over everything materialized from them; the affected set was 2 % of that, and naming it turned
  an unusable disclaimer into nine addressable sites.
- **A stated error direction is a claim, and one-sided is the easy mistake.** The direction that gets
  written down is the one that was *observed* — here, two false negatives in iter-122. The unobserved
  mirror was the more likely one, because the two substrates drift together.
- **When asking "does this instrument invalidate itself", derive the answer from what its verdict READS.**
  The prose answer was wrong on the first try, in the sentence written to be careful.

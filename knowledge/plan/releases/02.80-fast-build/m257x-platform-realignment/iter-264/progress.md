# iter-264 — progress

**Type:** tik
**Active strategy:** `TOK-08`.

## Phase B — the measurement, and PR-2 is REFUTED in the most useful way

**The corpus already stated this exact failure.** It is filed under the wrong service, in the wrong tense:

| site | what it says | why it did not help |
|---|---|---|
| `corpus/services/cms.md:271` | *"The Python studio submodule **had to be** cloned **before** any docker build, otherwise `make up` failed with `"/studio": not found`"* | **past tense**, under the **decommissioned** `cms` service |
| `corpus/ops/staging-bringup.md:428` | *"Quirk #3 — `cms/Dockerfile.dev` references removed `studio/` submodule … Comment them out"* | frames it as a **cms quirk to patch OUT**, the opposite of the current requirement |
| `corpus/ops/setup_guide.md:308-315` | *"there is no `make init-studio` step … To run the pipeline by hand, clone `anthropos-studio-room` yourself"* | frames the clone as **optional**, for a reader who wants the pipeline |
| `CLAUDE.md:574` | the tree is present *"on a box where a **build** or a hand-clone populated it"* | **inverted causality** — a local build cannot populate it, it FAILS on it |

So the defect is **not absence, it is misfiling plus staleness**: when the dependency moved from `cms` to
`app` at `fdb8034a`, the documentation of it did not move with it, and the surviving copy reads as dead
history about a service that no longer exists. That is `platform-alignment.md` §5's **"a named-consumer
list survives the merge that moved the consumer"** (M257x iter-23) — the milestone's own recurring class,
found again in its own corpus.

Per this iter's escalation clause, the refutation **changed the fix shape**: not "document a missing step"
but "**correct three sites that actively mislead**, and state the dependency where the reader meets it."

## Phase C — the edits

1. **`corpus/ops/setup_guide.md`** — section retitled *"Acquire the Studio runtime — REQUIRED before
   `make up`, or the `backend` build fails"*, carrying the verbatim failure, the two Dockerfile lines, why
   nothing in the documented flow supplies the tree (`.gitignore:79`, absent from `repos.yml`, no
   `.gitmodules` since `851cf3fb`, `init-studio` is a **cms** target), the exact `git clone`, and the
   **derive-it-from-the-Dockerfile** rule so the next service to grow the dependency is covered.
2. **`corpus/ops/setup_guide.md` § Starting the Services** — an inline warning at the `make up` step
   itself, for the reader who jumps straight there.
3. **`CLAUDE.md`** — the Studio-Room block's **inverted causality** corrected, plus the hard-build-dependency
   statement and why the demo path hides it.

## Phase D — grading

| | prediction | outcome |
|---|---|---|
| PR-1 | the guide is the only site naming the acquisition | **REFUTED** — 14 corpus files mention `anthropos-studio-room` |
| PR-2 | no corpus file states the image hard-fails without `studio/` | **REFUTED — and this was the valuable one.** `cms.md:271` states it precisely, in past tense, under a dead service |
| PR-3 | `studio-room.md` describes the pipeline, not the build dependency | **HELD** |
| PR-4 | the fences accept the edit | **HELD** — `repair-postcondition: OK`, 6 participating fences, 0 sites reported |
| PR-5 | `CLAUDE.md` omits the build dependency | **HELD, and worse than predicted** — it did not merely omit it, it asserted a **build** could populate the path |

## Close — 2026-08-10

**Outcome:** The corpus half of `D-M257x-262-2` is landed. The guide now states the `app/studio`
acquisition as the **required** pre-`make up` step it is, and **two actively-misleading statements are
corrected** — the guide's *"optional, to run the pipeline by hand"* framing and `CLAUDE.md`'s claim that a
local **build** populates the path (it cannot; it fails on it). **PR-2's refutation is the finding**: the
corpus knew this failure exactly and had filed it under the decommissioned service in past tense.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5

**Decisions:** `D-M257x-264-1` (the dependency moved from `cms` to `app`; its documentation did not).

**Side-deliverables:** none.

**Routes carried forward:**
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` → **corpus half CLOSED; tooling half OPEN.** Hoist
  `demo-stack/lib/studio.sh` so the dev bring-up shares it, and fence that the dev path uses it. A guide
  step is a mitigation; the fence is the fix.
- `FIX-M257x-264-cms-md-past-tense-dependency` → **new.** `cms.md:271` and `staging-bringup.md:428` still
  describe the live `app` dependency as dead `cms` history. Sweep for the general class: **requirements
  that migrated with a merged service but whose documentation stayed behind.**
- `FIX-M257x-263-dev-bringup-must-run-the-check`, `ROUTE-M257x-261-succession-projection-is-empty`,
  `FIX-M257x-262-demo-env-append-is-not-idempotent`, `ROUTE-M257x-258-the-pin-is-157-iters-stale` → open.

**Lessons:**
1. **When a service is merged away, its REQUIREMENTS migrate but its documentation does not.** The corpus
   had this failure written down, verbatim, and the merge left the sentence attached to a corpse. Grepping
   for the *symptom* (`"/studio": not found`) found it instantly; grepping for the *service* never would.
2. **"Undocumented" is a conclusion to earn, not assume.** iter-262 booked this as a documentation gap.
   It was a documentation **misfiling** — a different defect with a different fix, and the pre-registration
   that predicted absence is what forced the check.
3. **An inverted causal claim is worse than a missing one.** `CLAUDE.md` said a build could populate
   `app/studio`. A reader hitting the failure would have concluded their build was broken rather than that
   a step was missing.

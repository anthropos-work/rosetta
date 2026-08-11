**Type:** tik, under `TOK-08` (*census the mechanical classes; stop sampling them*).
Route: `FIX-M257x-264-cms-md-past-tense-dependency`, worked as an **enumeration**, per the milestone's own
rule that a class is closed by its population and never by its last member (§8, iter-169).

# iter-265 — the setup guide told you to delete the line the build requires

## What was measured

The class iter-264 named: **an operational requirement that migrated when a service folded into `app`,
whose documentation stayed attached to the decommissioned service.** Enumerated corpus-wide, at `cf9c469`.

| reading | number |
|---|---|
| decommissioned services, **derived** from the local-state column of `platform-migration-status.md` | **16** |
| corpus sites issuing an operational instruction naming one | **17**, across 10 files |
| of those, unmarked (no historical marker in scope) **before** repair | **1** |
| corpus instructions to **DELETE code that is LIVE in the clone set**, before repair | **2 reached, 3 known** |
| refused and bucketed (mention without an imperative; listing, not a command; the map itself) | 107 + 2 + 2 |

## The finding

`app/Dockerfile:45-46` hard-COPYs `/build/studio`, so `make up` cannot build `backend` without it —
iter-262 measured that, iter-264 documented it. This iter found **the corpus was simultaneously telling
operators to remove those exact lines**, in three places:

| site | what it said | why it is wrong now |
|---|---|---|
| `corpus/ops/setup_guide.md:771` | *"Edit `cms/Dockerfile.dev` and remove these two lines"* | in **the same file** as iter-264's *"Acquire the Studio runtime — REQUIRED"* |
| `corpus/ops/staging_from_dump.md:473` | *"Edit `cms/Dockerfile.dev` and remove:"* | same text, staging path |
| `corpus/ops/staging-bringup.md:428` | Quirk #3 — *"Comment them out … the Go binary runs fine without the Python studio runner"* | a long-lived skip-worktree patch, re-applied on every fresh clone |

**`RUN pip install --no-cache-dir -r studio/requirements.txt` is byte-identical between the obsolete
remedy and live `app/Dockerfile:46`.** The entries are titled by the **symptom** (`COPY … studio` fails),
and the symptom outlived the fold — so an operator greps their build error and lands on a page that
inverts the fix. Every clause was wrong for a current stack: the failing image is `app`, not `cms`;
deleting the lines is a **platform-repo edit** this release forbids; and the Go binary does *not* run
without the Python runtime, because `app/internal/cms/` now hosts the embedded studio-room pipeline.

`platform-alignment.md` §5 recorded this shape for **rext's** consumer lists (*"a named-consumer list
survives the merge that moved the consumer"*, iter-23). iter-264 found it inside this corpus. iter-265
found that inside this corpus it does not merely go stale — **it inverts**.

## The fence, and the reason it has four assertions instead of three

`rosetta-extensions/stack-core/decommissioned_instruction_guard.py` (`FENCE_KIND = "standalone"`),
committed at rext `25cdb84`.

- **A–C — the marker assertions.** Every enumerated instruction must carry a historical marker in scope.
  The decommissioned set is **derived** from the checked-in migration-status map — itself fenced against
  `repos.yml` by `platform_alignment_guard.py` — so a service leaving the clone set widens this
  population on the next run with no edit here (§2: derive it at the point of use).
- **D — no corpus instruction to DELETE code may name a line that is LIVE in the clone set.**

**D exists because A–C were measured to miss the defect, and that measurement is the iter's most useful
output.** Run read-only against the pre-repair tree (`git archive HEAD | tar -x`, no working-tree
mutation), the marker assertion fired on **1 site of 17** — `roadrunner.md:249` — and on **none of the
three that mattered.** The reason is exact: the obsolete remedy's own words, *"the `studio/` submodule
**has been removed from** `cms/main`"*, satisfy the marker vocabulary. **A fence whose green is produced
by the defect's own phrasing is worse than no fence**, and the only thing that surfaced it was
controlling the instrument against the tree it was built from (§9, iter-149).

D then fired on 2 of the 3 known instances, each cited to `stack-dev/app/Dockerfile:46`. **Its reach is
stated, not implied:** D reads only deletion instructions that *quote the code*, so
`staging-bringup.md:428` — which says "comment them out" in prose with no fence — is **out of reach and
was repaired by hand.** The first version of D's regex reached only 1 of 3 and would have published a
1-of-1 as a clean sweep; the colon form (*"and remove:"*) was added and is pinned by its own test.

Controls: `--selftest` A (17 sites resolve and still name their service) · B (a staged unmarked
instruction fires) · C (the same instruction with a marker is accepted — the **marker** is what is
graded) · D (deletion of a live line fires) · **D′ (the same instruction is SILENT when nothing is
live** — so D reads the clone set rather than the corpus alone). 12 tests, including
`test_marker_assertion_is_satisfied_by_the_defects_own_words`, which pins the A–C weakness as a
characterisation so a future reader knows why D was added rather than deleting it as redundant.

## Pre-registration grading

| PR | prediction | outcome |
|---|---|---|
| **PR-1** | ≥ 5 sites document the studio requirement as a `cms` matter | **HELD** — 6 (`cms.md` ×3 + the three guides) |
| **PR-2** | the mechanical slice is 15–70 sites | **HELD** — **17** |
| **PR-3** | ≥ 8 of the 10 `cd <decommissioned>` fences carry a marker | **REFUTED, at the margin — 7 of 10** (`messenger.md:166`, `roadrunner.md:249`, `storage.md:238`). One of the three, `messenger.md:166`, carries its marker in the paragraph *after* the fence, which the ±window read as absent: **the refutation is partly an instrument artifact and is reported as such** rather than as three findings |
| **PR-4** | no existing guard fires on any member of this class | **HELD** — `fence_command_guard.py` exits 1 on 10 findings, **all `ant-academy`/e2e path misses, none a class member.** `cd cms` is refused as an *unanchored single segment*: the guard cannot distinguish *a repo we did not provision* from *a repo that no longer exists*, which is precisely the seam this class lives in |
| **PR-5** | the class is not studio-only | **REFUTED** — every `migrated-contradicts-live` member is a studio member. The other 13 sites are *dead-but-marked* (`cd storage`, `cd messenger`, `cd roadrunner` in Testing sections of redirect docs), not migrated requirements. **The class is real and narrow: one requirement, six documentation sites** |

**PR-5's refutation is the honest correction to this iter's own framing.** iter-264 routed a class; the
enumeration says it is one requirement with six homes. What survives as a *class* is not "studio" but the
**shape** — and that is what the fence watches, so the population grows by itself when the next fold lands.

## Side note, recorded not chased

`stack-demo/` still carries clones of **`cms`, `graphql-wundergraph`, `jobsimulation`, `messenger`,
`roadrunner`, `storage`** — six repos that left `repos.yml`. Harmless (nothing builds from them) and
outside this iter's scope, but it is the literal form of the user's *"deprecated repos no longer treated
as part of the project"*, and it is why a `cd cms` fence can resolve on this box and not on a fresh one
(§8, iter-241: a fence's reach is a property of the clone set). Routed.

## Close — 2026-08-10

**Outcome:** The class is enumerated (**17 sites / 16 services / 10 files**, denominator stated), the
three inverted instructions and the past-tense filing are repaired, and the population is fenced by a new
standalone guard whose **fourth assertion exists because its first three were measured to miss the very
defect the iter repaired.** Pre-repair control: **RED** (1 unmarked + 2 live-deletion). Post-repair:
**green**, selftest A–D′ OK, 12 tests pass.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-265-1` (the requirement migrated; three copies of its remedy inverted),
`D-M257x-265-2` (a marker fence can be satisfied by the defect's own words — assertion D).

**Side-deliverables:** none. The README row + disclosed-triple bump (`23 of 35` → `24 of 36`, union) are
part of shipping the fence, not a separate fix.

**Routes carried forward:**
- `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach` — **new.** D reads only deletion
  instructions that quote the code. Prose-only ones (`staging-bringup.md`'s *"comment them out"`) were
  repaired by hand and nothing watches for the next.
- `ROUTE-M257x-265-stack-demo-carries-six-dead-clones` — **new.** Six repos that left `repos.yml` are
  still cloned in `stack-demo/`.
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` (tooling half), `FIX-M257x-263-dev-bringup-must-run-the-check`,
  `ROUTE-M257x-261-succession-projection-is-empty`, `FIX-M257x-262-demo-env-append-is-not-idempotent`,
  `ROUTE-M257x-258-the-pin-is-157-iters-stale` → open.

**Lessons:**
1. **A marker fence can be satisfied by the defect's own words.** *"has been removed from"* is both the
   vocabulary of a retraction and the vocabulary of the obsolete instruction it should have caught. Any
   guard that greys a *presence* rather than a *relation* has this failure mode; the only thing that
   surfaced it was running the new guard against the pre-repair tree instead of trusting its green.
2. **Grade the enumeration against the instances you already know.** D's first regex reached 1 of 3 known
   instances. Had the known set not been in hand, `1 live finding, 0 remaining` would have read as a
   complete sweep.
3. **A troubleshooting entry is indexed by its SYMPTOM, and symptoms outlive folds.** This is why the
   inverted remedy kept being reachable: the operator does not grep the repo name, they grep the error.

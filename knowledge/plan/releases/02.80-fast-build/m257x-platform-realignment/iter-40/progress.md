**Type:** tik (`iter_shape: cleanup` — a multi-claim sweep; the claim list in `overview.md` IS the planned
multi-step scope, per the skill's protocol-codified-shape carve-out).

# iter-40 — `FIX-M257x-iter39-claim-scoped-repair`

## Step 0 re-survey — the target's shape changed before a line was written

The routed framing was *"finish iter-39's cross-file drift inside the clause-5 file set."* A whole-tree
grep for every claim iters 38/39 adjudicated refuted that:

**Inside the 40 in-scope files the corpus is UNIFORM on all of them.** Taxonomy figures, `--template`, the
subgraph count, `organization_id on every table`, the "five internal Go libraries", `SkillPathSessionService`,
the `:5050` port — each is corrected or fenced everywhere in `corpus/services/**` + `corpus/architecture/**`.
iter-39's adversarial half **did** land its C1–C8 fixes corpus-wide *within scope*.

**Every surviving instance sat just outside** — `corpus/ops/**`, `.claude/skills/**`, `corpus/README.md`,
and `CLAUDE.md`. That is the finding, and it is not what the hand-off predicted:

> **A claim leaks to the EDGE of the previous repair's scope and stops there.** The repair partition *was*
> the 40 files. A claim does not respect it, so the drift ran outward until it hit a file nobody read, and
> pooled there. **The highest-propagation site in the tree — `CLAUDE.md`, read by every agent before any
> doc is — carried five of the eight claims.**

## What landed

**§5 rule 19 — *repair by CLAIM, not by FILE*** — authored into `corpus/ops/platform-alignment.md`
(`D-M257x-39-2` promoted; the protocol-evolution rule requires the generalising lesson to land in the same
commit). Carries the measurement (5 of iter-39's 8 self-inflicted defects were cross-file drift), the
*why-worse-than-nothing* argument, the grep-then-re-grep procedure, **the scope-edge corollary this iter
measured**, and the *must-not-adjudicate* clause.

**Eight claims swept whole-tree**, 20 files, +129/−54:

| claim | sites fixed | note |
|---|---|---|
| A taxonomy "60K / 18K" | 12 | **Both verdicts carried separately** — 18K **REFUTED** (below the 22,470 public floor), 60K **UNVERIFIED, not refuted**. Collapsing them was the named hazard |
| B `gen.py --template` | 3 | The flag does not exist and `parse_known_args` **silently absorbs** it, so every documented command succeeds and generates something unrelated. Each site now says so |
| C supergraph "2 → 1" | 2 | → **3 → 1**, with the reason the wrong figure spread (`915da06`'s own commit subject says 2→1; the tree it was committed against lists three) |
| D "the five internal Go libraries" | 2 | **Four** are imported; `authn` is a dependency of no service |
| E "`organization_id` on every table" | 2 | Retracted at both, incl. `CLAUDE.md` — the claim `security_compliance.md` withdrew at iter-33 and that was re-asserted three files away at iter-33's adversarial pass |
| F `:5050` as a live port | 8 | The Cosmo router was deleted from compose at platform `2adcf71`. Two sites were **executable**: a health-check `curl` and an `.env.local` assembly line writing a dead `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` |
| G "scoring is deterministic, NOT AI-scored" | 1 | `playthroughs.md` P2 — **the one that changes behaviour**: it granted Playthrough authors an exemption to assert simulation scores *exactly*. Withdrawn; the conjunction has one true conjunct |
| H `SkillPathSessionService` + singular `academy_*` | 8 | The RPC was **removed**, not re-hosted (six Connect handlers, not seven). Tables are plural; the Ent schema **file** is singular — the split is now stated so the next reader does not "fix" the file path |

## Two self-caught defects — the adversarial half over my own diff

Mandatory per §5 rule 18(a), and it was **not clean for a fifth consecutive pass**. Both defects were mine,
found by re-reading the diff rather than the anchors — the same asymmetry every prior pass has reported.

**(1) I violated rule 19 while writing rule 19.** I struck the `anthropos-graphql-1` row out of `dev-up`'s
service-set table and left its neighbours standing — while `anthropos-skillpath-1` is *equally* gone
(decommissioned at platform M507) and cms/jobsimulation/roadrunner are husks. A half-struck table teaches
the reader the unstruck rows are current: **exactly the failure the rule I had just authored describes.**
Repaired by re-deriving the whole table from the platform clone's own `docker-compose.yml` at origin
`2adcf71` — no `skillpath`, no `graphql`, no `skiller` service; cms/jobsimulation/roadrunner present as
husks — and the stale **"11 healthy containers"** count (which counted the two now-gone services) swept in
the same pass at all three of its sites.

**(2) I introduced a derived claim in repaired text.** My `playthroughs.md` retraction quoted a live figure
(*"1462 llm-backed checks vs 17 deterministic"*) that exists **only in iter-39's blocker ledger** — nowhere
in the corpus. My own overview had forbidden precisely this ("propagate the adjudicated verdict verbatim;
derive nothing new"), and the rule I authored an hour earlier ends with *"a repair pass must NOT
adjudicate."* Removed; the text now carries only the qualitative verdict — *most, not all* — which **is**
established in-scope at `ai_architecture.md:7,197`, verified.

> **Both defects are the same shape: the repair pass reached for authority it had not earned.** Once by
> half-applying a rule, once by importing a number from a plan document into the corpus. Neither would have
> been caught by checking anchors; both came from reading the diff as prose.

## Verification

- **Post-condition re-grep per claim** (rule 19's own procedure) — and it **earned its keep immediately**:
  it surfaced **three more `academy_*` sites** the first pass missed, including one in `CLAUDE.md`. Without
  the re-grep this iter would have shipped the very defect it exists to prevent.
- **All 12 links introduced resolve** (mechanical check, path-normalised against the filesystem).
- **All 5 corpus guards GREEN**; the orphaned-continuation grep clean (rc=1).
- Every figure written is traced to an in-scope established statement — checked one by one, which is how
  defect (2) was found.

## Close — 2026-08-02

**Outcome:** §5 rule 19 authored; 8 adjudicated claims swept whole-tree across 20 files (+129/−54), every
surviving instance found to be **outside** the previous repair's 40-file scope; two self-inflicted defects
caught and fixed by the mandatory adversarial pass over my own diff.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5. This iter cannot move clause 5 by construction; a repair does not meet a clause
that asks for a reading. It removes a confound from the reading.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
unchanged, re-fetched at open; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik) —
(6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-40-1`, `D-M257x-40-2`, `D-M257x-40-3`
**Side-deliverables:** the `dev-up` skill's service-set table re-derived from the platform clone (surfaced
by defect (1), landed because leaving it half-struck was worse than either extreme).
**Routes carried forward:**
- `MEASURE-M257x-iter41-clause5-sixth-pass` — next iter, with the instrument held **fixed**.
- `CHECK-M257x-iter40-migrate-tuple-still-lists-skillpath` — `dev-up/reference.md:39` and `SKILL.md`
  describe `migrate-dev.sh` atlas-migrating a 4-tuple that **includes `skillpath`**, a service with no
  compose entry at `2adcf71`. iter-01 documented this hand-maintained tuple as the milestone's founding
  time bomb. **Deliberately not touched here** — whether the tuple has since been derived from `repos.yml`
  is an rext-source question this repair pass must not adjudicate.
- `DOC-M257x-iter39-minors` / `DOC-M257x-iter38-minors` — unchanged.
**Lessons:**
1. **A claim-scoped repair's post-condition re-grep is not ceremony.** It caught 3 of 11 sites for one claim
   — a 27% miss rate on a sweep run by someone who had just written the rule.
2. **Writing a rule does not confer compliance with it.** The rule-19 violation was committed *in the same
   iteration that authored rule 19*, on the first table touched afterwards. Promoted into the rule text as
   the scope-edge corollary.
3. **The scope boundary of the previous repair is the highest-yield place to look next** — not the files it
   edited (rule 18's density result), and not the files it read. The ones it *couldn't* reach.

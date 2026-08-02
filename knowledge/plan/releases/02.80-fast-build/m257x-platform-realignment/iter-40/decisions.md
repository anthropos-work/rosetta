# iter-40 — decisions

## D-M257x-40-1: promote candidate rule 19, and extend it with what this iter measured — 2026-08-02

`D-M257x-39-2` proposed *repair by CLAIM, not by FILE* from iter-39's evidence (5 of 8 self-inflicted
defects were cross-file drift). Promoted verbatim into `corpus/ops/platform-alignment.md` §5 as **rule 19**,
plus two clauses iter-39 could not have known:

- **The scope-edge corollary.** iter-39's drift was assumed to be *inside* the 40 audited files. Measured,
  it is not: those 40 are uniform on every adjudicated claim, and 100% of the survivors sit immediately
  outside the audit's scope — `corpus/ops/**`, `.claude/skills/**`, `CLAUDE.md`. An audit's scope is a
  legitimate boundary for reading and never one for repair.
- **The must-not-adjudicate clause.** A repair pass propagates a verdict already settled inside the audited
  scope. If a site's correct form is not already established, route it. (Written before this iter violated
  it — see `D-M257x-40-3`.)

## D-M257x-40-2: re-derive the whole `dev-up` service table rather than strike one row — 2026-08-02

Striking only the `anthropos-graphql-1` row left `anthropos-skillpath-1` standing beside it, which is
*equally* gone (platform M507), and left cms/jobsimulation/roadrunner unmarked as husks. A half-struck table
asserts that the unstruck rows are current — rule 19's own failure mode, committed inside the iteration that
authored rule 19.

Two options were weighed. **Fence the table without fixing it** (cheap, honest, no derivation) versus
**re-derive it**. Re-derivation won because the ground truth is a file already on disk and already sanctioned
as authoritative for this milestone — the platform clone's `docker-compose.yml` at origin `2adcf71`, the same
source the migration-status map is machine-fenced against. That makes it *propagation from an established
source*, not fresh adjudication, so it stays inside `D-M257x-40-1`'s must-not-adjudicate clause.

The stale **"11 healthy containers"** figure was swept in the same pass at all three sites — it counted the
two now-absent services, so it is the same claim wearing a number.

## D-M257x-40-3: remove my own `1462 / 17` figure from the repaired text — 2026-08-02

The `playthroughs.md` retraction I wrote quoted *"live on `demo-1`: 1462 llm-backed checks vs 17
deterministic `EngineTextDiff` ones."* Those numbers appear **only in iter-39's blocker ledger** — a plan
document — and nowhere in the corpus.

**Removed rather than verified.** Verifying it live would have made it correct and *still* wrong for this
iter: it would be a measurement authored during a repair pass, in text pass six is about to read, which is
rule 18's highest-risk category and the exact thing `D-M257x-40-1`'s final clause forbids. The text now
carries only the qualitative verdict (*most, not all*), which is established in-scope at
`ai_architecture.md:7,197` and was checked.

The general point, recorded because it will recur: **a plan document is not a corpus source.** Milestone
ledgers accumulate precise measurements that never entered the corpus; a repairer with both open will reach
for them, and the resulting sentence looks better-evidenced than anything around it while being unsourceable
by a reader who has only the corpus.

## D-M257x-40-4: do not touch the `migrate-dev.sh` skillpath tuple — 2026-08-02

`dev-up/reference.md:39` and `SKILL.md` both describe the dev migrate step applying a 4-tuple that includes
**`skillpath`** — a service with no compose entry at origin `2adcf71`. iter-01 identified this
hand-maintained tuple as the milestone's founding time bomb.

**Not repaired here.** Whether rext's tuple has since been derived from `repos.yml` (TOK-01 step 2) is a
question about rext source, and answering it is adjudication, not propagation. Editing the docs to describe
a tuple I have not read would substitute one unverified claim for another — the shape iter-23 named as *"a
correction that is TRUE but INCOMPLETE."* Routed as `CHECK-M257x-iter40-migrate-tuple-still-lists-skillpath`.

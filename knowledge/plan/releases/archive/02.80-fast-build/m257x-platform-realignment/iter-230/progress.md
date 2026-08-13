**Type:** tik — under `TOK-08` (census the mechanical classes; stop sampling them), target chosen by the
user's redirect: **the corpus's claims about the platform**. A cited sha is the most mechanical claim the
corpus makes.

# iter-230 — do the corpus's commit shas exist?

## The census

Every backticked 7–40-hex token in `corpus/**` + `CLAUDE.md` (all-digit tokens excluded), resolved with
`git cat-file -e <sha>^{commit}` against **14 platform clones** (`stack-demo/`, `stack-dev/`) plus this repo
and the `rosetta-extensions` **authoring** copy.

| quantity | value |
|---|---:|
| documents scanned | 93 |
| distinct sha-shaped tokens | **138** |
| citation sites | **1,497** |
| resolve in exactly one repo | **132** |
| resolve in more than one repo | **0** |
| resolve nowhere | **6** |
| **demonstrably wrong** | **0** |

Per-repo, unique resolutions: `platform` 34 · `app` 21 · `rosetta-extensions` 10 · `jobsimulation` 8 ·
`messenger` 7 · `ant-academy` / `graphql-wundergraph` / `sentinel` / `next-web-app` / `storage` 6 each ·
`cms` / `roadrunner` 5 · `studio-desk` / `studio-room` 2 · `rosetta` (self) 8.

## The instrument, proved before the zero was published

`§9`: *a census that returns ZERO must prove its instrument.* Both directions, both run:

- **Fabricated shas resolve nowhere** — `deadbee`, `0123456789abcdef`, `aaaaaaa1`: 0 hits each.
- **Real shas resolve to their own repo** — `HEAD~3` read out of `app`, `platform` and `next-web-app`
  (`11664538`, `2adcf714`, `699fd846`) each resolved to exactly one repo, the right one.

## The residual is UNMEASURED, and that is the finding

The first pass reported **14** unresolvable. Eight of those resolved once the clone set was widened to
include this repo and the rext authoring copy — **an instrument defect, caught before it was published**,
and exactly `§9`'s *state the substrate before booking a failure*. Of the remaining **6**, every one belongs
to a repo **in no clone set on this box**:

| sha | sites | repo the prose names |
|---|---:|---|
| `13c248e6` | **61** | `infrastructure` |
| `7dd1b80` | 11 | `db-backup` |
| `6e1fb15b` | 4 | `db-backup` |
| `5d944e4a` | 3 | `ant-singularity` |
| `b810b28` | 2 | `colony` |
| `b49eb7af` | 1 | `ant-observability` |

**82 citation sites rest on evidence this box cannot re-read** — headed by `13c248e6`, the `infrastructure`
commit that settles the standing **cms M810** question and which `CLAUDE.md`, `corpus/README.md`,
`service_taxonomy.md`, `platform-migration-status.md` and `org-repos.md` all lean on. iter-123 cloned
`infrastructure` to establish it; the clone is gone.

### Predictions, graded — 3 of 4 REFUTED

| id | prediction | result |
|----|-----------|--------|
| `P-230-1` | ≥ 150 distinct shas cited | **REFUTED — 138** |
| `P-230-2` | ≥ 1 sha resolves in no clone | **HELD — 14, then 6 after the instrument was widened.** But the *reading* behind the prediction is refuted: none is a corpus error |
| `P-230-3` | `app` is the repo the most distinct shas resolve in | **REFUTED — `platform` 34, `app` 21.** The corpus cites the orchestrator's history more than the monolith's |
| `P-230-4` | ≥ 1 short sha resolves in ≥ 2 repos | **REFUTED — 0 of 138.** Short-sha ambiguity is not a hazard in this corpus; *absence* is |

`P-230-2` is the one to read carefully. It is graded **HELD** because the literal claim held, and the close
says so — but the belief that motivated it (*"some of these will be typos or inventions"*) is **refuted at
0 of 132 measurable**. Recording that split rather than a bare HELD is the point of sealing predictions:
a prediction can be right about the number and wrong about the world.

## Close — 2026-08-10

**Outcome:** the corpus's 138 cited commit shas were censused for the first time. **132 of 132 measurable
resolve; 0 are wrong.** The 6 that do not resolve are all in repos no clone set contains — **82 citation
sites, headed by the 61-site `infrastructure` sha that settles cms M810** — so the corpus's most-leaned-on
platform evidence is currently unverifiable from this box. Disclosed in the protocol doc.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-230-1` (an unresolvable sha is partitioned, never graded false),
`D-M257x-230-2` (no clone was fetched to close the residual).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Suite state at close** — no pytest section run; this iter changed no rext code. The corpus edit is prose
only, fenced by `guard_family` at the next run.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-230-82-sites-on-uncloneable-evidence` → **new.** The cheapest closure is to clone
  `infrastructure` (and `db-backup`) into a stack workspace so the 76 highest-value sites become
  re-readable. iter-123 did it once and the clone was not kept. **This is a decision about the clone set,
  which is a `repos.yml`-adjacent policy question, not a corpus edit** — routed, not taken.
- `ROUTE-M257x-230-sha-census-has-no-fence` → the census is a script this iter wrote and threw away. As a
  standing guard it would catch a fabricated sha the moment it is written. Second deliverable; tripwire.
- `ROUTE-M257x-229-anchor-rot-is-19-of-22-invisible`, `ROUTE-M257x-225-no-profile-for-sanctioned-host`,
  `ROUTE-M257x-225-hostprofile-role-strings-name-a-retired-gate-host`,
  `ROUTE-M257x-222-pin-advance-needs-a-reproof`, `ROUTE-M257x-223-classify-the-ten-drifted-baselines`,
  `ROUTE-M257x-224-drift-guard-blind-to-stale-clone`,
  `ROUTE-M257x-228-corpus-disagrees-with-itself-about-refs`,
  `ROUTE-M257x-227-archived-repo-selfdesc-is-stale` → all open, unchanged.

**Lessons:**
1. **An unresolvable reference is two findings wearing one symptom** — a wrong citation, or a repo you do
   not have. Partition on clone-set membership *before* reporting a rate.
2. **A census's own clone set is part of its instrument.** Eight of the first 14 "failures" were the
   instrument's, not the corpus's, and only widening the substrate told them apart.
3. **A prediction can be right about the number and wrong about the world.** `P-230-2` held literally and
   its motivating belief was refuted at 0/132; the close records both rather than the flattering one.
4. **The repos the corpus most needs are the ones `repos.yml` never lists**, so every guard built on
   "resolve it in a clone" is structurally blind to exactly the evidence the migration story rests on.

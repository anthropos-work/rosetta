**Type:** tik

# iter-23 — the batchA residual, and the consumer list that outlived its consumer

## What was planned

Two declared lines (so the scope-creep tripwire counts against a 2-step shape):

1. `DOC-M257x-iter22-batchA-residual` — the 18 enumerated clause-5 blockers, **with every CORRECTION
   re-derived against platform source** before application (`platform-alignment.md` §5, the rule iter-22 paid
   for).
2. `DOC-M257x-iter22-ops-guides-5050` — the dead `:5050` references in `corpus/ops/**`.

## Phase A — re-derivation, and what it caught

Sixteen of the eighteen handed corrections **confirmed** against `app` @ `5ba17044` / platform `2adcf71`
(the evidence table is in this iter's `overview.md`). The other two are the iter's first finding, and they are
two *different* failure shapes:

**(a) A correction can be INCOMPLETE rather than wrong.** The hand-off said colony is *"split: `app` +
`messenger` @ `v0.35.2`; `sentinel` + `storage` @ `v0.34.3`"*. Both halves are true. It names **four of six**
— the `cms` and `jobsimulation` containers the default `graphql` profile still starts are on a **third** pin,
`v0.35.1`. Applied verbatim it would have replaced one incomplete claim with another, and the row would have
read as freshly verified. Promoted to protocol §5: *when a correction enumerates, re-derive the ENUMERATION.*

**(b) Two blockers the handed list did not carry at all**, both in `hiring.md`:

- `:142-144` describes **our own tooling** as it was three releases ago — *"`persona_write.go:68-73` writes
  both `jobsimulation.sessions` and `public.local_jobsimulation_sessions` (col builders `sessionCols()` /
  `localSessionCols()`)"*. M257 re-pointed this: `persona_write.go:91` writes
  `{"public", "job_simulation_sessions", sessionCols(), …}` and **`localSessionCols` no longer exists**.
- the doc's **thesis**. Its own framing — *"the headline of this doc is the ONE table that actually feeds the
  score"* — names `public.local_jobsimulation_sessions`, **dropped** by `20260729133514.sql:58-62`. Correcting
  the table rows while leaving the thesis would have left the doc arguing for the wrong table. It now opens
  with a RE-GROUNDED banner stating the three facts that changed.

## Phase B — the finding the sweep was not looking for, measured on a live stack

Writing the correction for blockers #7–#11 required knowing which service actually reads Directus at HEAD.
It is **`backend`**, in-process: `app/cms_reader_switch.go` swaps the cms content reader to the in-process cms
server once Directus is configured — *"a DIRECT domain call — no proto round-trip … and no internal traffic to
a standalone cms"* — and `app/main.go:971-973` `log.Fatalf`s without `DIRECTUS_BASE_ADDR`.

rext's `--local-content` cutover re-points `DIRECTUS_BASE_ADDR` for every service in
`DIRECTUS_DATA_CONSUMERS` = `cms`, with `test_only_cms_is_repointed_not_other_services` **explicitly asserting
`backend` must not carry it**. That was correct when `cms` was the consumer.

Confirmed by measurement rather than reasoning — one `docker inspect` on the standing `demo-1`:

```
backend : DIRECTUS_BASE_ADDR=https://content.anthropos.work   DIRECTUS_TOKEN=(empty)
cms     : DIRECTUS_BASE_ADDR=http://directus:8055
```

**The per-stack Directus serves a consumer that no longer reads, and the reader that does is pointed at prod
anonymously.** This is the founding class wearing a different face: not a schema name, a **service name in a
consumer list**. Nothing errors — the list still names a real, running container that still starts and still
holds the var; the read simply happens elsewhere. And, for the third time in this milestone, **the suite was
not silent about the defect, it was arguing for it** (§8 rule 3).

It is a strong candidate cause for two of clause 2's open failure classes
(`FIX-M257x-iter15-directus-versions-403` — an anonymous read against *prod* Directus is a 403 machine — and
plausibly the `library_category` shape drift, since prod content and the replayed local catalog are different
content models). **Deliberately not concluded**: iter-19 proved those two independent of the *serving* defect,
which is a different question from this one. It is a third line of investigation, so per the tripwire it is
**routed, not landed** — it needs an rext change, a tag, and a cold cycle to prove.

## Phase C — the sweep

**52 edits across 13 files, 52/52 anchors matched exactly once**, in three passes
(`.agentspace/scratch/work-m257x/iter23/sweep23{,b,c}.py`; deliberately non-idempotent — a re-run SHOULD fail,
and that is the guard). One anchor in pass a reported `0x` because the hand-transcribed quote wrapped
differently from the file; re-anchored in pass b rather than loosened.

| file | edits | what was false at HEAD |
|---|---|---|
| `services/hiring.md` | 14 | the thesis + the score source + the two-row write-set + the cross-subgraph NULL-bubble + the `validation_*` schema and the middle table's real name (`validation_attempt_skill_results`) + our own seeder |
| `architecture/external_services.md` | 5 | the Directus env inverted onto `backend`; *"only Postgres + backend run"* (it is **nine** containers); the mermaid's `Frontend → CMS` edges; the smart-proxy prose; `docker compose logs cms` |
| `services/ai-readiness.md` | 4 | `jobsimulation.` qualifiers on five `public` tables — contradicting the doc's own table header at `:149` |
| `architecture/shared_libraries.md` | 2 | one colony pin where there are three; `ai v1.40.1` → `v1.40.2` |
| `services/clerkenstein.md` | 3 | the `clerk-deploy-1` artifact's `v0.34.3` asserted as the platform's current pin, in three places |
| `services/studio-room.md` | 1 | `cd studio/studio-room` — no such path; the root **is** `app/studio/` |
| `ops/` × 10 files | 22 | dead `:5050` — the retired Cosmo router, incl. `run_guide.md` / `setup_guide.md` / `staging-bringup.md` (the onboarding path) and the `:15050` tailscale front row iter-13 deleted |

Two judgement calls worth naming:

- **`staging-clerk.md`'s two `allowed_origins` lists were NOT edited.** They are a *record of what the Clerk
  instance holds*; changing the recorded values would make the record wrong. Annotated instead: `:5050` is
  dead weight rather than a needed origin, and `:8082` is what now needs allowing.
- **Explicitly-fenced history was kept.** `hiring.md` still names the dropped tables eight times — every one
  inside a "this is what it was before the drop" fence, which the grading rule treats as not-a-blocker. A
  seeder author needs to know the shape changed, not just what it changed to.

**Guards, all five green after:** `corpus_index_guard` (84 docs) · `platform_alignment_guard` (both
directions) · `demo_knob_guard` (30 env + 10 CLI, both directions) · `dev_flag_guard` (6 flags) ·
`story_org_count_guard`.

## Close — 2026-08-01

**Outcome:** the enumerated clause-5 residual is exhausted — 19 handed blockers + 2 found at re-derivation
closed via 52 exactly-once edits — and the re-derivation itself produced the iter's largest finding: **rext
re-points the per-stack Directus at `cms`, a container that has not been the reader since cms-in-app**,
measured live on `demo-1`.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close, unchanged; trigger stays at occurrence 1 of 2) — (4) user-blocker: n —
(5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-23-1 … D-M257x-23-4 (this iter's `decisions.md`)
**Side-deliverables:** none — both lines were planned scope. The protocol-doc update is a close-obligation,
not a side-fix.
**Routes carried forward:**
- **`FIX-M257x-iter23-backend-directus-not-repointed`** (rext, next tik) — add `backend` to
  `DIRECTUS_DATA_CONSUMERS`, invert `test_only_cms_is_repointed_not_other_services` (it pins the pre-merge
  shape), watch it RED first, prove on a cold cycle. Measure its effect on
  `FIX-M257x-iter15-directus-versions-403` **before** attributing.
- **`DOC-M257x-iter23-rext-stale-session-comment`** (rext, low) — `stack-seeding/cmd/stackseed/main.go:533`
  still describes the hiring funnel as writing *"`jobsimulation.sessions` + `public.job_simulation_sessions`"*.
  Comment-only; batch it with the next rext change rather than spending a tag on it.
- **The clause-5 full re-read** (next iter) — 40 files, 5 sub-agents, `wc -l` positive control per file. The
  list is exhausted; the corpus is not measured. iter-22 fixed a 29-item list and the tree still held 53.

**Lessons:**
1. **A correction can be incomplete rather than wrong**, and that is the harder half of the re-derive rule —
   an enumeration that is right about everything it names can still be wrong about the set. Promoted to §5.
2. **After a fold, grep the tooling for the folded service's NAME as a value**, not only for its schema.
   Consumer lists, `depends_on`, front-port tables, probe targets, env re-point maps. Nothing errors when
   one of these goes stale — the named container is still real and still running. Promoted to §5 with the
   one-command check (`docker inspect <container> | grep <VAR>` beats any amount of reasoning).
3. **Writing the correction is itself a measurement.** The Directus finding did not come from auditing; it
   came from needing to know what was true in order to write one paragraph. A sweep that only pattern-matches
   replacements would have missed it.

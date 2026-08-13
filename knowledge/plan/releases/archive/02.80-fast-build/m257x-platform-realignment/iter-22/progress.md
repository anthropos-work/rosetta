**Type:** tik — under `TOK-01: instrument first, then follow`.

# iter-22 — the enumerated residual, executed; and the residual was 2.5× larger than enumerated

## What was planned

Apply iter-21's 21 enumerated clause-5 blockers, re-audit by full read, close clause 5.

## What happened

**All 21 anchors verified.** Two of the *corrections* did not survive re-derivation. And the full-read
re-audit — the one the protocol now requires — returned **53 blockers**, not the ~0 a clean sweep would
predict, and not the 21 that were inherited.

### Phase 1–3 — pass 1, and the refutation

29 edits applied as `(file, old, new)` tuples, each asserting its anchor occurs **exactly once**. 29/29
matched; 28 consumed on re-check (the 29th is `old ⊂ new` by construction, applied once).

**But `D-M257x-22-1`:** hand-off items #8 and #10 said `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`
was stale and should read `backend:8083`. Measured at origin `2adcf71` **before** applying:

    docker-compose.yml:52   JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401   (backend)
    docker-compose.yml:258  JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401   (messenger)
    docker-compose.yml:256  CMS_RPC_ADDR=http://cms:8091                       (messenger)
    docker-compose.yml:265  SKILLER_RPC_ADDR=http://backend:8083               (messenger)

Only `SKILLER_RPC_ADDR` was re-pointed — and the husk addresses are **deliberate**, per
`app/main.go:1196-1202`: *"additive + DORMANT: external callers (messenger) keep hitting the standalone cms
via `CMS_RPC_ADDR` **until the M809 re-point**."* Applying the correction would have replaced two true
statements with false ones.

**The source of the error is the 22nd blocker.** The refuting citation iter-21 trusted was
`corpus/services/backend.md:175` — a corpus line asserting messenger points *all four* addresses at
`backend:8083`. It points two. One false corpus line, cited as authority, produced two false corrections in a
hand-off authored to be applied mechanically.

### Phase 5 — the full-read re-audit

40 files, fanned across 5 sub-agents, each reading whole files with a `wc -l` positive control (**7,700+
lines read**, every file 100%).

| batch | blockers | dominant cause |
|---|---|---|
| A | **19** | the `jobsimulation` schema is gone (tables → `public`, `sessions` → `job_simulation_sessions`, `local_*` mirrors DROPPED) · `cms` is a husk but docs route work through it |
| B | **10** | *"merged into `app`"* propagated as *"gone from local compose"* — in two files that contradict their own tables |
| C | **12** | the **post-M809 end state** asserted as current |
| D | **7** | next-web-app's **Next 15 → 16** upgrade, uncaught for four releases (`middleware.ts` → `proxy.ts`) |
| E | **5** | the frontend's real transport; the in-process (not RPC) cms read path |
| **total** | **53** | |

### Phases 6+ — passes 2, 3, 4

**+64 further edits**, same exactly-once discipline, 64/64 matched. Closed batches **C, E, B, D** (34 of the
53) plus the lower-grade findings in those batches.

Notable individual fixes: `make migrate S=jobsimulation` (re-materialises the dead schema — `repos.yml:17-19`
sets `migrations: false`, so the migrating set is `app` alone) · the skill-path CMS hop, in-process since
cms-in-app (`app/internal/skillpath/session.go:205-207`) · `platform-migration-status.md:60` citing
`app/internal/roadrunner/`, which does not exist · `backend.md`'s labs-api client "lands in PR 6" (it landed,
`main.go:735-738`) · `services/README.md`, the one file that never took M257x's roadrunner + chronos
corrections, so the index contradicted the docs it indexes.

**And a defect this iter introduced and the audit caught:** pass 1 replaced `docker compose up -d graphql`
with `-d cms` — the husk, which serves no subgraph. It is `backend`. Fixed before commit.

## Re-measurement

| metric | pre-iter | post-iter |
|---|---|---|
| clause-5 blockers (**full read**, 40 files) | **53** (measured this iter; the inherited figure was 21) | **≈19 residual** — batch A's 18 + 3 lower-grade cross-file citations |
| gate clauses met | 3 of 5 | **3 of 5** — clause 5 NOT closed |
| corpus files corrected | — | **24** |
| enumerated edits applied | — | **93**, 93/93 anchors matched exactly once, 0 misses |

**Clause 5 did not close, and claiming otherwise would repeat exactly the error this iter exists to correct.**

## Close — 2026-08-01

**Outcome:** 93 corrections across 24 corpus files closed 34 of 53 clause-5 blockers — but the iter's real
finding is that the inherited residual was **2.5× under-counted** (21 vs 53), and that **two of the inherited
corrections were false**, traceable to a false corpus line cited as authority.
**Type:** tik
**Status:** closed-fixed-partial (the sweep landed clean and the re-audit is complete and honest; the clause
did not close, and the ~19 residual is routed with a named handler)
**Gate:** NOT MET (3 of 5 — clauses 1, 3, 4)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik moved the metric: 53 → ~19) —
(3) re-scope: n (platform origin HEAD `2adcf71` re-checked at open AND close, unchanged — **occurrence stays
1 of 2**) — (4) user-blocker: n (no platform edit needed; every fix was corpus-side) — (5) cap-reached: n
(1 tik of 5) — (6) protocol-stop: n — Outcome: continue.
**Session note:** the session exits here on **budget**, not on an enum condition. The numeric 5-tik cap was
NOT reached (1 tik). Closing cleanly with the residual enumerated beats opening batch A's 18-item sweep and
dying mid-edit — runs 6 and 9 of this milestone did exactly that.
**Decisions:** D-M257x-22-1 … D-M257x-22-6.
**Side-deliverables:** `platform-alignment.md` §5 gains *"re-derive the CORRECTION, not just the anchor"* +
the corollary *"merged-in-production is not removed-from-compose"*, and a **struck correction to §5's own
example list**, which cited the RPC address as an audit miss when it is correct at origin HEAD.
**Routes carried forward:**
- `DOC-M257x-iter22-batchA-residual` — batch A's 18 blockers (hiring 6 · external_services 5 · ai-readiness 3
  · shared_libraries 2 · clerkenstein 1 · studio-room 1), enumerated in `HANDOFF-next.md`. **Next tik.**
- `DOC-M257x-iter22-ops-guides-5050` — 13 dead `:5050` refs in `corpus/ops/**` (out of clause-5 scope, but on
  the onboarding path). `D-M257x-22-5`.
- `CHECK-M257x-iter22-clerk-sdk-drift` — `app` is on `clerk-sdk-go/v2 v2.7.0`, not the `v2.6.0` the
  Clerkenstein Alignment DNA targets. `colony` is likewise split (`app`+`messenger` `v0.35.2`;
  `sentinel`+`storage` `v0.34.3`). An alignment score run against the stale pin is not measuring the platform.
- 3 lower-grade cross-file citation errors (`backend.md:71` sha, `messenger.md:110` line, taxonomy Directus flow).
- clause 2's three causes, unchanged.

**Lessons:**
- **Re-derive the CORRECTION, not just the anchor.** All 21 anchors were fine; two corrections were false.
  The failure mode a mechanical hand-off invites is not a moved anchor — it is an **inherited falsehood
  wearing a `file:line`**. And an audit that reads the corpus to correct the corpus is circular: the citation
  has to terminate in platform source.
- **A clean sweep is not a clean tree.** 29 correct edits landed, and the tree still held 53 blockers. Fixing
  what you were handed measures the hand-off, not the corpus. Only the full read measures the corpus — and
  every batch's dominant cause was something no prior term-scoped sweep had ever named.
- **The commonest false claim was a true one over-applied.** *"Merged into `app`"* is true; *"therefore gone
  from local compose"* is false, and both edits were made in the same pass. Three containers still start on
  every `make up`. The map already had the word — `running_but_unfederated` — and the service docs had not
  adopted it.
- **The audit caught a defect the sweep introduced** (`-d cms` for `-d graphql`). A sweep needs its own
  reader, and the reader must not be the author.

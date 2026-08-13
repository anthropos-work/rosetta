# iter-48 — clause-5 EIGHTH pass: the blocker ledger

Seven auditors, the instrument frozen at iter-41's on every knob — same seat count, same briefing, same
partition, all 40 files read top-to-bottom with a per-file `wc -l` positive control, plus the diff seat.

---

## THE HEADLINE: iter-47's "the pre-existing residual measured ZERO" is REFUTED

| pass | iter | auditors | corpus read | blockers | pre-existing | induced by the immediately preceding repair |
|---|---|---|---|---|---|---|
| 1 | 21/33 | grep-scoped → 5 | pre-repair | 25 | — | — |
| 2 | 34 | ~5 | post-33 repair | 13 | — | — |
| 3 | 38 | 6 | post-34 repair | 11 → 17 | — | — |
| 4 | 39 | **7** | post-38 repair | **37** | — | — |
| 5 | 41 | **7 — fixed instrument** | post-39 repair | **18** | **9** | **9** |
| 6 | 47 | **7 — SAME fixed instrument** | post-46 repair | **7** | **0** | **7** |
| **7** | **48** | **7 — SAME fixed instrument** | post-48 repair | **12** | **10** | **2** |

**Counting convention:** the same one every prior pass used — all seats plus the diff seat, and the diff
seat's findings count wherever they land. iter-47's 7 included `ops/demo/coverage-protocol.md:614`, so
excluding `corpus/ops/**` here to make the number smaller would be changing the instrument between passes
to flatter effect. Under clause 5's *literal* file scope (`corpus/services/**` + `corpus/architecture/**`,
which is exactly the 40-file partition) the count is **10**; both figures are reported below.

**The decisive observation is not the count. It is WHERE the findings are.**

iter-47 read these same 40 files with these same seven seats and reported **zero** blockers in text
iter-46 had not touched. iter-48 — same instrument, one repair later — books multiple blockers in
**files neither repair edited**, authored months before either:

- `external_services.md:662` — the LiveKit agent names. Authored **2026-03-02**. In seat A's set both
  passes.
- `hiring.md:189-196` — a NOT NULL, UNIQUE, undefaulted column missing from a "minimal write-set".
  Authored **2026-07-15**. iter-47 saw this exact passage and booked it a **MINOR**; iter-48's seat
  escalated it to a blocker after checking the DDL, the Ent schema and the shipped seeder.
- `dependency_map.md:19` — Storage's infrastructure row, contradicted by its own twin doc. Authored
  **2026-05-11**.
- `ops/demo/stories-spec.md:599` — the refuted live-recompute claim, still unfenced. Authored
  **2026-07-01**.

None of these is new. None was introduced by a repair. **They were present, unchanged, while a
seven-auditor full read one iteration earlier reported the pre-existing residual as zero.**

> **So the two-term model that iter-47 introduced — corpus term plus repair term — is right about the
> arithmetic and wrong about one of its measurements. The corpus term did not reach zero. It was
> measured as zero by one reading, and a second reading of the same corpus with the same instrument
> measured it as non-zero.**

That is a statement about the **instrument**, not the corpus, and it is the finding this pass exists to
report. Clause 5 asks for *a reading that returns zero blockers*. Eight passes have returned
`25 → 13 → 11 → 17 → 37 → 18 → 7 → …`, and the variance between two runs of the *same frozen instrument
on the same files* is now measurable and is comparable in size to the residual being chased.

---

## The findings

Class key: **PRE** = authored before this milestone · **MILE** = authored by an earlier iteration of this
milestone · **THIS** = authored by iter-48's own repair.

### Seat A — 2

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 1 | `graphql-wundergraph.md:171` | CI/prod pulls schema artifacts from GitHub Releases on `anthropos-work/{app,jobsimulation,cms}` | `ci/update-subgraph.sh` has exactly one `gh release download`, `-R anthropos-work/app`; the other two were deleted at `915da06`. The two bullets above it carry a historical fence; this one does not | PRE (2026-07-23) |
| 2 | `external_services.md:662` | the LiveKit agents are `anthropos-agent-eu` / `anthropos-agent-us` | `anthropos-agent-eu` exists nowhere. The EU agent is the bare `anthropos-agent` (`livekit.go:110,120`); only US is suffixed (`:126`). The eu/us split lives on the **endpoint** (`azure-eu`/`azure-us`) | PRE (2026-03-02) |

### Seat B — 2

Both re-derived against `app @ 5ba17044` by the adjudication, not accepted on the seat's report.

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 11 | `ai-readiness.md:410` | *"`keepStartedMembers` **excludes members with no step-1 signal** from the aggregate"* | It reads no step-1 signal at all. `queryReadinessStarters` (`steps.go:915-938`) is `SELECT DISTINCT user_id FROM public.ai_readiness_user_step_progresses … AND status <> 'not_started'` — a **progress row**, not a `user_skill_evidences` row. Wrong in **both** directions: a member with step-1 evidence but no progress row is dropped; one with an `in_progress` row and zero evidences is kept. The platform says so itself (`steps.go:907-914`: *"This DB signal is the only real 'has started' check"*). The sentence is the stated justification for the seeding contract three lines below it | **PRE** — `12b587b`, **2026-06-29**, five weeks before this milestone opened |
| 12 | `ai_architecture.md:202` | *"LiveKit + OpenAI Realtime is the engine for **new** sessions (gated by the `flag_use_realtime_openai` PostHog flag)"* | The flag does not select LiveKit over ElevenLabs and gates no "new sessions". Engine choice is **per sequence, from the CMS `voice_engine` field** (4-member enum, `jobsimulation.go:1079-1085`; nil → `gptrealtime`, `:1594-1600`). The flag is read **inside** `CreateAgentDispatch` — the LiveKit path already entered — and only swaps the dispatched **endpoint** to `openai-hosted` (`calls/livekit.go:131-135` read, `:140-144` effect). Residency-relevant: flipping it moves a live voice session off an Azure-EU endpoint, which the doc's routing section does not cover | **PRE** — `d791487`, **2026-06-01**, two months before this milestone opened |

Seat B also booked **14 minors** (anchor drift, an ambiguous bare filename resolving to two real files, a
compressed 4-entry role set) and one cross-doc finding it correctly routed OUT of its own file set —
`ops/demo/media-substrate-spec.md:33-34`, *"the bytes … never in prod S3"*, which is false about **capture**
(`recording/chime.go:38`, `:188-189`, `:262-265` all sink to S3; `:325` reads it back; `SaveBunnyVideoId`
`:361` runs *after*). Bunny is the **serving** layer. Out of the 40-file partition and out of clause 5's
scope; routed forward, not counted.

### Seat C — 1

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 3 | `architecture_overview.md:298` | *"**16** schemas carry an `organization_id` with no policy at all"* | **23.** 16 is the neither-mixin **subset**. Re-derived at `app @ 5ba17044`: 139 schema files, 30 `OrganizationMixin{}`, 7 `OrganizationIDMixin{}`, only **4** files declaring any `Policy()`. The doc's own link target says 23 in almost the same words (`security_compliance.md:76-77`). The error runs in the *"isolation is handled"* direction | MILE (iter-34) — **and a LEAK of iter-46's repair**: `301d61a` changed `**16 carry an` → `**23 carry an` in `security_compliance.md` and left this twin standing |

### Seat D — 0 · Seat F — 0

Both seats read their full sets and found no false claim. Seat D re-derived Studio-Room's provider set
independently and **confirmed the corpus is right**: `ai.py:704-724` maps exactly
`{openai, azure, anthropic}` and raises on anything else; `grep -rin 'bedrock|boto3' app/studio/` → 0,
with a passing positive control. Seat F executed the corpus's own `platform_alignment_guard.py` against
`repos.yml` (exit 0) and verified all four per-domain wiring call sites, the 23 jobsim tables, both mirror
`DROP TABLE`s, and all 17 cited platform shas with their dates.

### Seat E — 4

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 4 | `hiring.md:28` (repeated `:148`, `:276`) | *"`jobsimulation.sessions` was dropped (`20260722104506.sql:79`)"* | That line is `DROP TABLE "sessions"` under `search_path=public` — it dropped **`public.sessions`**. No `app` migration touches the `jobsimulation` schema at all, and `askengine/registry.go:192` says that schema survives frozen until M710. Contradicted by the corpus twins `service_taxonomy.md:52` and `dependency_map.md:78` | MILE (iter-23) |
| 5 | `hiring.md:189-196` | the "minimal write-set" for `public.job_simulation_sessions` | omits **`token`** — `NOT NULL`, UNIQUE, no default (`20260722104506.sql:13,29`). The only required-and-undefaulted column missing, so an INSERT built from this contract fails. The shipped seeder writes it (`persona_write.go:152-158`); the word "token" appears nowhere in the doc | PRE (2026-07-15) — **iter-47 booked this passage as a MINOR** |
| 6 | `hiring.md:20-22` (repeated `:147-148`) | *"`20260729133514.sql` … back-fills [the mirror] into the canonical entity"* | No back-fill exists. That migration re-points *referencing* assignment-session link ids and drops the mirrors. `SET "score"` → 0 hits across the migration set | MILE (iter-23) |
| 7 | `dependency_map.md:19` | Storage's Infrastructure is *"Postgres, Redis, S3"* | Storage has zero DB/Redis code, no redis in `go.mod`, and no `DB_CONNECTION`/`REDIS_ADDR` in compose. Contradicts its own twin `storage.md:14,21` (*"owns no database"*, *"**Database**: none"*) | PRE (2026-05-11) |

### Seat G (diff) — 3

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 8 | `ops/demo/coverage-protocol.md:616` | *"the `ai_readiness_live_snapshots` read was gated behind a *closed* `CycleID`"* | The closed-cycle-gated read is on **`ai_readiness_snapshots`** — `buildResponseFromSnapshots` calls `ListAIReadinessSnapshots` (`readiness.go:771-772`). `ai_readiness_live_snapshots` is the **askengine / Talk-to-Data write-side mirror**, and `live_snapshots.go:54-57` says so in its own docstring. The file's own twin 12 lines below names the right table | **THIS** — but see the note below |
| 9 | the *"**Four** ways a request leaves the EU"* enumeration — `external_services.md:569`, propagated by iter-48 to `architecture_overview.md:249` and `security_compliance.md:186` | that the list is exhaustive | It is not. Studio-Room's `TARGET SERVICE` — which **iter-48's own item 3 admits into the list** — has a third value `openai` that builds a bare `OpenAI(api_key=…)` against `https://api.openai.com` (`studio/services/ai.py:383`, `:706-708`; `config_template.ini:30-31`). **The same undercount defect as the blocker this repair was fixing, one scope level up** | root PRE; **the inconsistency is THIS** — naming Studio-Room in item 3 is what pulled it into a list that then undercounts it |
| 10 | `ops/demo/stories-spec.md:599` | *"the live-recompute never completes in the coverage harness's budget"*, unfenced | Refuted by M219 (2.09 s), and recorded as refuted in `seeding-spec.md:496-498`, `services/ai-readiness.md:371,449-450` and `CLAUDE.md:324`. **The paraphrase-sibling of the exact paragraph blocker #7 repaired** — and `repair_leak_guard` is verbatim-only, so it cannot see it (ran GREEN) | PRE (2026-07-01) |

**Note on #8, and it is the sixth consecutive occurrence of one shape.** iter-48 rewrote that sentence
(present tense → past tense, plus a retraction) and **preserved the false table name it did not notice**.
That is verbatim iter-46's blocker #2 — *"rewrote this sentence and preserved the false conjunct"* — and
the fifth-then-sixth instance of *the author of a correction violating it while writing it*.

---

## Adjudication — 12 raw, 12 unique, 12 held

iter-41 took 21 raw reports down to 18 unique, so adjudication is not a formality here. Every finding was
re-derived against platform source at `app @ 5ba17044` **before** acceptance rather than taken on the
seat's word. Spot-record of the re-derivations that could have collapsed the count:

| check | result |
|---|---|
| **Duplicates across seats** | none. A#2 (`external_services.md:662`) and B-minor-1 (`ai_architecture.md:180`) publish the **same** false agent name at **two distinct sites** in two files — two sites, not one finding double-counted. Graded consistently as one blocker (A#2) + one minor (B-1) only because B's site is a **diagram label** and A's is prose asserting the platform "runs" those agents |
| A#1 `gh release download` | `update-subgraph.sh:9` — **exactly one**, `-R anthropos-work/app`. Corpus claims `{app,jobsimulation,cms}`. **HELD** |
| A#2 agent names | `anthropos-agent-eu` → **0 hits**. EU is bare `anthropos-agent` (`livekit.go:110,115`); only US is suffixed (`:126`). eu/us lives on the **endpoint**. **HELD** |
| C#3 16-vs-23 | `security_compliance.md:76-77` says **23**, and names 16 as the neither-mixin subset, in almost the same words. **HELD**, and it is a leak of `301d61a` |
| E#5 `token` | `20260722104506.sql:13` `NOT NULL`, `:29` UNIQUE, no default. The word "token" appears **0 times** in `hiring.md`. An INSERT built from that write-set fails. **HELD** — and iter-47 read this passage and booked it a MINOR |
| E#7 Storage infra | `dependency_map.md:19` = *"Postgres, Redis, S3"*; `storage.md:14` = *"owns no database"*; `storage/go.mod` redis → **0**. **HELD** |
| B#11, B#12 | re-derived above, line by line. **HELD** |

## The induced / pre-existing split — and the instrument that produced it

**Class key:** **PRE** = authored before this milestone · **MILE** = authored by an earlier iteration of
this milestone · **THIS** = authored by iter-48's own repair.

| class | n | findings |
|---|---|---|
| **PRE** — predates M257x entirely | **7** | #1 (2026-07-23) · #2 (2026-03-02) · #5 (2026-07-15) · #7 (2026-05-11) · #10 (2026-07-01) · #11 (2026-06-29) · #12 (2026-06-01) |
| **MILE** — earlier iters of this milestone | **3** | #3 (iter-34) · #4 (iter-23) · #6 (iter-23) |
| **THIS** — induced by the repair this pass just made | **2** | #8 · the inconsistency half of #9 (whose root claim is PRE) |

> **Induced by the immediately preceding repair: 2. Not induced: 10.**
> iter-47 measured that same split as **7 induced / 0 not**. It is a complete inversion.

**The instrument that produced it — stated because the split is now known to be instrument-dependent,
and this is the pass that established that.** The reading was taken with **iter-41's instrument frozen on
every knob**: seven seats, the same briefing text, the same 6-way file partition (the 40 files of
`corpus/services/**` + `corpus/architecture/**`), every file read top-to-bottom in a single un-`offset`
`Read` with a per-file `wc -l` positive control, plus a seventh diff seat reading the repair commit. No
knob was touched between iter-47 and iter-48.

**That is the point.** The instrument did not change; the measurement did. Seven of these twelve sit in
text that neither iter-46's nor iter-48's repair touched, authored between **2026-03-02** and
**2026-07-23** — present, unchanged, and *in seats' assigned file sets* while iter-47's reading reported
the pre-existing residual as **zero**. So iter-47's zero was a property of that reading, not of the
corpus, and the two-term model it introduced is right about the arithmetic and wrong about one of its
inputs.

**Were the pre-existing findings reachable by the mechanical fences?** **No — none of the 7 was.** Traced
individually rather than assumed:

- `repair_leak_guard` ran **GREEN** and is verbatim-only by construction. `D-M257x-48-4` pins that a
  **paraphrase** (#10) is out of its reach; `D-M257x-48-9` pins that a **number-only** correction (#3) is
  too, and that lowering K buys two false positives without catching it.
- `claim_twin_guard` fires on claims **already in a ledger**. A claim never previously reported is
  invisible to it — which is every one of the 7.
- `anchor_construct_guard` resolves anchors. All 12 sit at **resolvable** anchors; a false claim behind a
  valid anchor is exactly what it does not look at. (It is what caught the 10 silently-broken
  cross-references in `D-M257x-48-8`, a different class.)
- `platform_alignment_guard` fences `repos.yml` states, not prose.

Two of the 7 (#3, #7) are **cross-doc numeric/factual contradictions** — a class a not-yet-built
consistency fence could reach, and `FENCE-M257x-iter49-numeric-leak` is already routed forward for
exactly it. The other five required someone to read a sentence and check it against platform source.
**There is no shipped fence whose blast radius includes them.**

## Pre-registered predictions — graded

Written into `overview.md` and committed **before any auditor ran** (`cabc3b1`).

| # | prediction | outcome |
|---|---|---|
| **headline** | **3 blockers**, fewer than iter-47's 7 and not zero | **REFUTED — and in the direction that matters.** Fewer than 7 was wrong: the count went UP |
| 1 | zero of the **verbatim** self-contradiction class | **HELD for the verbatim class** — `repair_leak_guard` ran GREEN on the repair commit and no finding is a verbatim survival. **But #3 and #10 are both self-contradictions the fence structurally cannot see** (`D-M257x-48-9`, `D-M257x-48-4`), so the prediction was true and the inference it invited was wrong |
| 2 | **at least one** blocker in text written to explain a correction | **CONFIRMED — #8.** Sixth consecutive iteration |
| 3 | if there is a residual, `ai_architecture.md:42-68` carries it | **REFUTED.** That block was read by two seats and verified clean line-by-line, including every mechanism independently re-derived rather than trusted. Seat C reported the newly-written AI-Providers paragraph *"100% clean"* |
| consequent | clause 5 does **not** close | **CONFIRMED** |

**Three of five refuted, and the two that held were the least informative.** The heavily-repaired block
predicted to fail was clean; the untouched files predicted to be fine were not.

---

## Clause 5

**NOT MET.** Gate stays **4 of 5**. The adjudicated reading is **12** (10 under clause 5's literal file
scope); clause 5 is met only by a reading returning **zero**, and the user has ruled it is not re-cut.

## Is zero reachable? — the honest read this pass owes the user

Eight passes, all measuring the same corpus modulo repairs: **25 → 13 → 11 → 17 → 37 → 18 → 7 → 12.**

The series does not converge, and the reason is now visible rather than inferred. **Every time the
instrument got better, it found more**, and the last three readings were taken with an instrument that
did **not** change at all — 18, then 7, then 12. The run-to-run variance of the *frozen* instrument is
therefore **±5 or more**, which is **larger than the entire residual being chased**.

That has a consequence clause 5 cannot absorb: **a reading that returns zero would not be evidence that
the corpus is clean.** It is evidence about the reading — which is precisely what iter-47's zero turned
out to be, one pass later, on text no repair had touched. Chasing zero by re-reading rewards the pass
that happens to read least well.

**My read: zero is not reachable by this method.** Not because the corpus is bad — the repair work is
real and the same seats verified large tracts clean line-by-line, including every block this pass
predicted would fail — but because the measurement's noise floor sits above zero. Continuing to iterate
would be grinding an instrument against its own variance.

**The two things that would change the answer**, offered as options rather than a recommendation, since
this is a user decision:

1. **Make the reading a fence rather than a reading.** Everything mechanized so far has *stayed* fixed;
   the residual lives entirely in the un-mechanized class (*"is this newly written sentence true?"*),
   which has produced findings in eight consecutive passes. `FENCE-M257x-iter49-numeric-leak` reaches 2
   of the 7; nothing reaches the other 5.
2. **Re-cut clause 5 as a bounded-variance criterion** (e.g. *no blocker of the classes the fences claim
   to cover*, or *two consecutive readings agreeing within N*) — **explicitly ruled out by the user**, and
   recorded here only so the option set is complete, not as a proposal.

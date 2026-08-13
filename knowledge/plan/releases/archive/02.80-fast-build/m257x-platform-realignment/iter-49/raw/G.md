# Seat G — the DIFF seat (M257x iter-49, ninth clause-5 reading)

## Coverage

**Subject:** the uncommitted working-tree diff `git -C /Users/marco/workspace/anthropos/rosetta diff -- corpus/`
on branch `m257x/platform-realignment` (11 files, +100 / −35).

**Reviewed: 19 of 19 hunks across 11 of 11 touched files.**

| File | Hunks | Reviewed |
|---|---|---|
| `corpus/architecture/ai_architecture.md` | 2 (`@@ -177`, `@@ -199`) | ✅ |
| `corpus/architecture/architecture_overview.md` | 2 (`@@ -246`, `@@ -295`) | ✅ |
| `corpus/architecture/dependency_map.md` | 1 (`@@ -16`) | ✅ |
| `corpus/architecture/external_services.md` | 3 (`@@ -566`, `@@ -585`, `@@ -659`) | ✅ |
| `corpus/architecture/security_compliance.md` | 1 (`@@ -183`) | ✅ |
| `corpus/ops/demo/coverage-protocol.md` | 1 (`@@ -613`) | ✅ |
| `corpus/ops/demo/stories-spec.md` | 1 (`@@ -596`) | ✅ |
| `corpus/ops/platform-alignment.md` | 1 (`@@ -487`) | ✅ |
| `corpus/services/ai-readiness.md` | 2 (`@@ -406`, `@@ -415`) | ✅ |
| `corpus/services/graphql-wundergraph.md` | 1 (`@@ -168`) | ✅ |
| `corpus/services/hiring.md` | 4 (`@@ -17`, `@@ -143`, `@@ -189`, `@@ -273`) | ✅ |

**Ground truth read directly:** `stack-demo/app` @ `5ba17044` (`terraform/migrations/*.sql` — note the
migrations live in `terraform/migrations/`, **not** `internal/data/migrate/migrations/` as the brief
stated; `internal/aireadiness/`, `internal/cms/`, `internal/jobsimulation/calls/`,
`internal/data/ent/schema/`, `internal/organization/`, `internal/askengine/`, `studio/` = Studio-Room),
`stack-demo/graphql-wundergraph` @ `60c229f` (incl. `git show 915da06`), `stack-demo/storage` @ `4ce8ece`,
`stack-demo/platform` @ `2adcf71`, and the `rosetta-extensions` seeder.

### Verdict on the 12 claimed repairs

| # | Claim | Verdict |
|---|---|---|
| 1 | GH Releases on `anthropos-work/app` only; one `gh release download` | ✅ VERIFIED (`ci/update-subgraph.sh:9` is the sole hit; `subgraphs.conf` = `BACKEND=v1.360.0` alone; `git show 915da06 -- ci/update-subgraph.sh` deletes both other lines) |
| 2 | EU agent is bare `anthropos-agent`; only US suffixed | ✅ VERIFIED (`calls/livekit.go:110,120,126,142`; 0 hits for `anthropos-agent-eu`) |
| 3 | 23 schemas with `organization_id` and no policy (16 = neither-mixin subset) | ✅ VERIFIED (independent count converges on 23 = 16 + the 7 `OrganizationIDMixin{}` users; matches the `security_compliance.md:77` twin) |
| 4 | `20260722104506.sql:79` dropped `public.sessions`, not `jobsimulation.sessions` | ✅ VERIFIED (`:79` = `DROP TABLE "sessions";`; `atlas.hcl:8` pins `search_path=public`; 0 `jobsimulation.` refs in the migration set; `askengine/registry.go:192` = M710) |
| 5 | `token` NOT NULL / UNIQUE / no default | ✅ column facts VERIFIED (`:13`, `:29`) — ❌ but the *characterisation* is wrong, see **B4** |
| 6 | `20260729133514.sql` re-points link ids, no back-fill | ✅ VERIFIED (`:15-23` UPDATE `session_id`/`js_session_id`; `SET "score"` = 0 hits set-wide) — twin left standing, see **B3** |
| 7 | Storage = S3 only | ✅ VERIFIED (`go.mod` has no pg/redis driver; 0 grep hits; `cmd/root.go:53,56` = two S3 buckets; `internal/storage/storage.go:162,269` = the two `file://` fallbacks) |
| 8 | Closed-cycle read is `ai_readiness_snapshots` | ✅ VERIFIED (`readiness.go:771-772` → `ListAIReadinessSnapshots`; `live_snapshots.go:54-57` is the askengine mirror) |
| 9 | FIVE ways out of the EU (5th = Studio-Room `openai`) | ⚠️ code path real, **operationally unreachable** — see **B5**; count-twin left standing, see **M1** |
| 10 | "never completes" refuted (2.09 s) | ✅ VERIFIED against the corpus twins; anchor defect **M3** |
| 11 | `keepStartedMembers` gates on a PROGRESS ROW, not step-1 evidence | ✅ VERIFIED (`steps.go:915-938` = `SELECT DISTINCT user_id FROM public.ai_readiness_user_step_progresses … status <> 'not_started'`; docstring at `:907-914`) |
| 12 | `flag_use_realtime_openai` does not select LiveKit; engine is per-sequence CMS `voice_engine` | ✅ VERIFIED (`livekit.go:131-135` read / `:140-144` effect, inside `CreateAgentDispatch`; `jobsimulation.go:1079-1085` enum, `:1594-1610` nil→`gptrealtime`, `:1350` `seq.VoiceEngine`) — **twins left standing in three places, see B1/B2** |

---

## BLOCKERS

| # | Location | The false claim a reader would act on | Refutation (platform / corpus `file:line`) | Class |
|---|---|---|---|---|
| **B1** | `corpus/architecture/external_services.md:672` | *"**Coexists with ElevenLabs**: LiveKit + OpenAI Realtime powers new sessions (gated by `flag_use_realtime_openai`)"* — the **exact claim this same diff retracts** four lines above it, in a file this diff edited. The repair rewrote `:668` (agent names) and left the refuted sentence in the same bullet list. | `app/internal/jobsimulation/calls/livekit.go:131-144` — the flag is read **inside** `CreateAgentDispatch`, i.e. after the LiveKit path is entered, and its only effect is `agentName="anthropos-agent"` + `agentEndpoint="openai-hosted"`. Engine selection is `app/internal/cms/directus/collections/jobsimulation.go:1350` (`seq.VoiceEngine`) → `:1594-1610`. The corpus now literally contradicts itself: `ai_architecture.md:207` says the flag *"gates no 'new sessions'"*. | twin left standing (dominant induced class) — **and in a touched file** |
| **B2** | `corpus/services/jobsimulation.md:123` **and** `:126` | `:123` *"new sessions increasingly use LiveKit + OpenAI Realtime (gated by the `flag_use_realtime_openai` PostHog flag)"*; `:126` *"the OpenAI Realtime voice path is gated by the `flag_use_realtime_openai` PostHog flag"*. Same refuted claim, two more places. Claim 12 was repaired in exactly one file. | same as B1 (`livekit.go:131-144`; `jobsimulation.go:1350,1594-1610`) | twin left standing |
| **B3** | `corpus/services/jobsimulation.md:102` | *"`20260729133514.sql:58-62` (*"5. Drop the mirrors."*) **back-fills** then `DROP TABLE`s both `local_jobsimulation_sessions` and `local_skill_path_sessions`"* — the **exact "back-fills" wording** this diff refutes in `hiring.md:20-23` and `:153-155`. | `app/terraform/migrations/20260729133514.sql` — the only `UPDATE`s are `:15-23` (`SET "session_id"` / `SET "js_session_id"`) and `:36-44` (orphan nulling). `grep -rn 'SET "score"' terraform/migrations/` = **0**. No `INSERT INTO` anywhere in the file. | twin left standing |
| **B4** | `corpus/services/hiring.md:210` (**new text**) | *"`token` … carries **no default** — the **only** required-and-undefaulted column in the table"*. A reader building the minimal INSERT from this contract concludes the other columns are defaulted and omittable; the INSERT then errors on three of them. | `app/terraform/migrations/20260722104506.sql` — **four** columns are `NOT NULL` with no `DEFAULT`: `owner_id` (`:6`), `sim_id` (`:7`), `sim_type` (`:10`), `token` (`:13`). Verified by `awk 'NR>=3&&NR<=27' … \| grep 'NOT NULL' \| grep -v DEFAULT`. The same defect makes the lead sentence *"`token` is the one column that makes the INSERT itself fail"* false. | **repair-manufactured overshoot** |
| **B5** | `corpus/architecture/external_services.md:588-593` (new item 5), replicated as the count "five" into `architecture_overview.md:246-249` and `security_compliance.md:186` | *"**Five** things **can** send a request outside the EU … 5. **Studio-Room's own `openai` `TARGET SERVICE`**"* — presented flatly, alongside four paths that need no code change, with no note that **no shipped configuration selects it**. A residency/DPIA reader acts on this list as an egress inventory. | Every `*_AI_*_MODEL` line in **all three** checked-in configs is `azure`: `app/studio/configs/production_config.ini:26-36`, `development_config.ini:26-36`, `config_template.ini:37-51`. `studio/gen.py:31-33` picks the ini by `ENVIRONMENT`; `:44-53` permits env override of **only** `{AZURE,OPENAI,ANTHROPIC}_{API_KEY,ENDPOINT}` — **never the service**. `services/ai.py:704-724` `get_client` reads `target_engine['service']` from that ini and has no `openai` default. Outside `tests/`, `service: openai` appears nowhere. Selecting it requires editing a checked-in file. | **repair-manufactured overshoot** (softest of the five; the code path is real, the reachability is not) |

---

## MINORS

1. **`external_services.md:577-578` — the count-twin the same hunk left standing.** Item 4 still reads *"the easiest of the **four** to miss"* inside a list the same hunk renumbered to **five** (`:569`). The repair changed the header count and the item count and missed the count *inside* an item.
2. **`external_services.md:595-600` — dangling demonstrative.** The *"**Why this one was missed**"* blockquote explains the **cms content-layer vendor default** — i.e. item **4**. After item 5 was inserted above it, *"this one"* reads as item 5, whose miss had an entirely different cause (already given in item 5's own last sentence). Referential drift induced by the insertion.
3. **`stories-spec.md:603` — anchor points at the wrong claim.** `[`services/ai-readiness.md:371,449-450`]` — `:371` is correct; `:449-450` is `effectiveCycleId = …` / `cyclesQ.isFetched`, i.e. the **cycle-param misattribution** correction. The 2.09 s measurement is at `ai-readiness.md:459-460` (*"LIVE `GET /api/workforce/ai-readiness` → HTTP 200 · 2.09 s · 304 KB"*).
4. **`hiring.md:23` — unsupported explanatory clause.** *"no score was copied from the mirror to the canonical row — **the canonical row already carried it**"*. Neither `20260722081626_jobsim_data_model.sql` nor `20260722104506.sql` contains a single `INSERT INTO`; the migration set moves **no session data at all** (the new table is created empty at `:2`, the old one dropped at `:79`). The operative half ("no back-fill") is right; the causal half has no basis in the source cited.
5. **`hiring.md:20-21` — anchor attributes the wrong statement.** *"`20260729133514.sql:58-62` — *"5. Drop the mirrors."* — **re-points the referencing rows** … and then `DROP TABLE`"*. `:58-62` is only the drop block; the re-point is at `:13-23`.
6. **`hiring.md:32-33` — "as the twins already said" overstates the twins.** `service_taxonomy.md:52` says only *"the `cms`, `jobsimulation` and `skillpath` schemas are legacy husks"*; `dependency_map.md:78` says only *"the legacy `jobsimulation` schema is non-authoritative"*. Neither mentions **M710** or that the schema *survives*. Only `app/internal/askengine/registry.go:192` supports the M710 half.
7. **`hiring.md:202` — the edited line preserves a DDL-false conjunct.** *"Non-null `status`, `started_at`, `ended_at`, …"* — `started_at`/`ended_at` are `timestamptz **NULL**` (`20260722104506.sql:14-15`) and `status` carries `DEFAULT 'pending'` (`:11`). Tolerable as "write a value" shorthand until iter-49 inserted a DDL-precise NOT-NULL/UNIQUE/no-default sentence two lines below, which now reads as making the same kind of claim about the same list.
8. **`ai-readiness.md:410` — ambiguous antecedent + inverted relative clause.** *"⚠️ **It reads no step-1 signal at all**, which this sentence claimed until M257x iter-49"*. "It" = `keepStartedMembers` (two clauses back), but the immediately preceding sentence correctly states the live path **does** read `user_skill_evidences` (`readiness.go:74-89`, `:330`) — so "It" scans as the dashboard, producing an apparent self-contradiction. And *"which this sentence claimed"* grammatically attaches to the **corrected** fact, not the retracted one.
9. **`ai_architecture.md:200-201` — sequence vs simulation.** *"Engine choice is per SEQUENCE … a 4-member enum on the authored **simulation**"*. The field is on the sequence struct (`jobsimulation.go:911` `VoiceEngine *SimulationVoiceEngine`, consumed at `:1350` as `seq.VoiceEngine`); the simulation carries no such field.
10. **`external_services.md:590-591` — citation does not support the mechanism.** `config_template.ini:30-31` is cited without its path (`app/studio/configs/`), and its `OPENAI_ENDPOINT` is **discarded**: `studio/services/ai.py:383` is `return OpenAI(api_key=self.api_key)` — the `endpoint` kwarg passed by `get_client` (`:719`) is never used. The `https://api.openai.com` conclusion holds only via the SDK's default `base_url`, not via the ini line cited.
11. **`ai_architecture.md:196-201` — unfenced three-way tension introduced.** The new sentence makes `gptrealtime` the nil-default (`:1595-1596`, `:1609`), two lines under a table row marking `gptrealtime` **Deprecated** and under a heading reading *"### Active Engine: LiveKit + GPT Realtime"*. All three are individually true; nothing reconciles them.
12. **`content-stories-routes.md:202` vs the new hiring.md claim.** `hiring.md:296` now asserts `jobsimulation.sessions` is *"frozen and unwritten"*, while `content-stories-routes.md:202` still routes the per-session result read to *"`jobsimulation.sessions` by `sessionId` (M248)"*. Pre-existing text, but the repair's new assertion is what makes it contradictory.
13. **`stories-spec.md:600-601` — blockquote abuts a paragraph with no blank line.** CommonMark permits a block quote to interrupt a paragraph, so it renders; it is the only such splice in the file and diverges from the corpus's blank-line convention.

### Explicitly checked and CLEAN (no twin found)

`anthropos-agent-eu` (0 corpus hits outside the corrected sentence) · `{app,jobsimulation,cms}` release enumeration (0 hits) · *"`jobsimulation.sessions` was dropped"* (0 unfenced hits) · Storage `Postgres, Redis` (0 hits) · *"16 schemas"* (0 hits) · `ai_readiness_live_snapshots` as the dashboard source (0 hits) · `dependency_map.md` table structure intact (12-row table, no broken pipes) · `platform-alignment.md:490` anchor `hiring.md:86 → :93` correct (old `:86` and new `:93` are byte-identical) · every other introduced anchor resolves: `livekit.go:110,120,126,131-135,140-144` · `jobsimulation.go:1079-1085,1594-1600` · `readiness.go:771-772,308-312,289` · `live_snapshots.go:54-57` · `steps.go:907-914,915-938` · `20260722104506.sql:2,13,29,79` · `20260729133514.sql:58-62` · `askengine/registry.go:192` · `intelligence.go:1700` · `atlas.hcl:8` · `persona_write.go:152-158` · `storage.md:14,21` · `seeding-spec.md:496-498` · `ai-readiness.md:371` · `stories-spec.md:599` · `ai.py:383,704-724,706-708` · `config_template.ini:30-31` · `ci/update-subgraph.sh:9`.

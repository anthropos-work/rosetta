# iter-49 — clause-5 NINTH pass: the blocker ledger

Seven auditors, the instrument frozen at iter-41's on every knob — same seat count, same briefing, same
6-way partition of the 40 files, all read top-to-bottom with a per-file `wc -l` positive control, plus the
diff seat. Ground-truth clones **unchanged from iter-48** (`app @ 5ba17044`, `platform @ 2adcf71`,
`graphql-wundergraph @ 60c229f`, `next-web-app @ bb3313bc`), so the instrument is frozen in its inputs as
well as its method.

---

## THE HEADLINE: the two new fences closed the gaps they were built for, and the induced term GREW anyway

| pass | iter | auditors | blockers | pre-existing | induced by the immediately preceding repair |
|---|---|---|---|---|---|
| 1 | 21/33 | grep-scoped → 5 | 25 | — | — |
| 2 | 34 | ~5 | 13 | — | — |
| 3 | 38 | 6 | 11 → 17 | — | — |
| 4 | 39 | **7** | **37** | — | — |
| 5 | 41 | **7 — fixed instrument** | **18** | 9 | 9 |
| 6 | 47 | **7 — SAME fixed instrument** | **7** | 0 | 7 |
| 7 | 48 | **7 — SAME fixed instrument** | **12** | 10 | 2 |
| **8** | **49** | **7 — SAME fixed instrument** | **14** | **7** | **7** |

**Pre-registered before any report was read** (`overview.md`): **6 blockers, 2 induced / 4 pre-existing.**
**Actual: 14 blockers, 7 induced / 7 pre-existing.** The prediction is **refuted in every term**, and it is
refuted hardest on the one the whole iteration was designed to move: the induced count was predicted to
hold at 2 and came back at **7**.

Per seat: **A 1 · B 1 · C 2 · D 2 · E 3 · F 0 · G (diff) 5**. 14 raw, 14 unique, 14 held.

---

## The findings

**Class key:** **PRE** = present before this iteration's repair · **THIS** = manufactured by it.

### Seat A — 1

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 1 | `external_services.md:668` | *"**only the US one is suffixed**, `anthropos-agent-us`"* | There is a **third** agent name: `calls/livekit.go:115` sets `anthropos-agent-chain` for the `livekitchain` engine, selected **before** the eu/us branch. A reader provisioning LiveKit dispatch would miss it | **THIS** — an **overshoot inside this pass's own correction**. The rest of the sentence verifies: `anthropos-agent-eu` really is 0 hits, and the eu/us split really is on the endpoint |

### Seat B — 1

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 2 | `ai-readiness.md:406-419` | the active-cycle population filter is `keepStartedMembers` / `queryReadinessStarters` (`status <> 'not_started'`, any step, unbounded) — *"one with an `in_progress` row and zero evidences is kept"* | On the active-cycle branch `buildLiveResponse` calls **`keepInCycleStep1`**, not `keepStartedMembers` (`readiness.go:387-392`; `keepStartedMembers` is the `cyc == nil` else-branch). The real predicate is `queryInCycleStep1Completers` (`:638-660`): `StepSkillMapping` **AND** `StepCompleted` **AND** `CompletedAtGTE(cycle.StartDate)`. So an `in_progress` row is **dropped**, a step-2/3 completion alone is dropped, and a pre-cycle `completed_at` is dropped | **THIS** — this pass correctly retired the old `user_skill_evidences` claim and **attached the right mechanism to the wrong branch**. A seeder built to it yields an empty-but-error-free dashboard |

### Seat C — 2

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 3 | `security_compliance.md:76`,`:120` + `architecture_overview.md:299-300` | *"**31** schemas auto-filter by ORGANIZATION (the 30 `OrganizationMixin{}` users + `Membership`)"* | **32.** `Organization` declares its own `Policy()` (`schema/organization.go:56`) whose Query rule is `rule.FilterSameOrganizations()` (`:96`), pinning the query to `organization.ID(org.ID())` (`rule/organization.go:41-49`). It uses neither mixin (`:17-23`). The fence's own re-measurement (`:92-97`) **names `organization.go` as a `Policy()` file and then never counts it** | **PRE** — and this is the passage whose own text records it has been wrong **four times**. This pass edited the adjacent 16→23 figure and left the 31 standing |
| 4 | `shared_libraries.md:145` | authn *"Imported by: via colony: app (the former cms / jobsimulation / skillpath usage is all folded in)"* | The still-running `cms` and `jobsimulation` **husks import `colony/authn` directly** in their own Go source (7 and 9 sites) and both containers start in the default `graphql` profile (`docker-compose.yml:144`, `:83`). The colony (`:42`), ai (`:105`) and taxonomy (`:181`) rows in this same file all name the husks; the authn row alone drops them | **PRE** |

### Seat D — 2 · both in `roadrunner.md`, both pre-existing

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 5 | `roadrunner.md:87` | *"`SubmissionPackage(...)` — Submit a **batch** of runs in one call"* | It calls the same `CreateSubmission` with runtime `"zip"` and returns **ONE** token (`rpcsrv/rpc.go:43-57`), hitting Judge0's **single**-submission endpoint (`runner/runner.go:62`); `languages.go:32-33` maps `"zip"` to language_id 89 — a multi-**file** program, not a batch | **PRE** |
| 6 | `roadrunner.md:49` | *"runs an Asynq worker pool for asynchronous **batch submissions**"* | One poll task per **single** submission (`runner.go:126-127`; one task type at `worker/worker.go:48`) — and it contradicts **this same file** at `:97` (*"Every submission enqueues exactly one poll task"*) | **PRE** |

### Seat E — 3 · all in `hiring.md`, all pre-existing

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 7 | `hiring.md:66-68` | `CreateOrganizationSimInvitationLink` is *"the very call the `HiringConfigSeeder` uses to write the 5 positions"* | The seeder **never calls it**. It writes raw rows via `CopyRowsIdempotent(ctx, "public", "organization_sim_invitation_links", …)` (`stack-seeding/seeders/hiring_config.go:99`) and never reads `organizations.is_hiring` | **PRE** |
| 8 | `hiring.md:113-115` | with Clerk-only wiring *"the `HiringConfigSeeder` cannot write the 5 positions in the first place"* | With `is_hiring=false` it **still writes all 5** — it bypasses the hard-erroring manager (`organization/siminvitationlink.go:63`) and gates only on the blueprint predicate (`hiring_config.go:65`). **This mis-routes any empty-positions debug** | **PRE** |
| 9 | `hiring.md:293-296` (recurs `:33-34`, `:157-158`) | *"`jobsimulation.sessions` still exists, frozen and unwritten, until M710"*, stated unqualified **in a demo/dev-stack section** | **No local stack ever creates that schema** — `platform/repos.yml:17-19` has no `schema:` key, the only `CREATE SCHEMA` in app's whole migration set is `auth`, and rext says so verbatim (`persona_write.go:58-61`). The M710 survival fact is **prod-only** | **PRE** — *and note this is a refinement of, not a contradiction to, this pass's claim-4 repair, which Seat G independently re-verified as correct* |

### Seat G (diff) — 5 · the highest-yield seat again

| # | site | the false claim | what is true | class |
|---|---|---|---|---|
| 10 | `external_services.md:672` | *"LiveKit + OpenAI Realtime powers new sessions (gated by `flag_use_realtime_openai`)"* | The exact claim this pass **retracts** at `ai_architecture.md:207` — refuted by `calls/livekit.go:131-144`. It stands **four lines below a line this same diff edited, in a file this same diff edited** | **THIS** (leak) |
| 11 | `jobsimulation.md:123`,`:126` | the same refuted `flag_use_realtime_openai` claim | Claim 12 was repaired in **exactly one of three places** | **THIS** (leak) |
| 12 | `jobsimulation.md:102` | the mirror migration *"**back-fills** then `DROP TABLE`s"* | The exact word this pass refutes in `hiring.md`: `SET "score"` = **0 hits**; the only UPDATEs (`20260729133514.sql:15-23`) touch link ids | **THIS** (leak) |
| 13 | `hiring.md:210` | `token` is *"the **only** required-and-undefaulted column in the table"* | **Four** are (`20260722104506.sql:6,7,10,13`: `owner_id`, `sim_id`, `sim_type`, `token`). So a minimal write-set a reader derives from this contract **still errors** — the defect this repair was fixing, one level down | **THIS** — manufactured, in text written to explain a correction |
| 14 | the new 5th EU-egress path — `external_services.md`, propagated to `architecture_overview.md` + `security_compliance.md` | that Studio-Room's `openai` `TARGET SERVICE` is a live way a request leaves the EU | It is selected by **no shipped config**: all three `app/studio/configs/*.ini` pin `azure`, and `gen.py:44-53` lets env override only keys and endpoints, **never the service**. The arm exists; nothing reaches it | **THIS** — an overshoot, and the count *"five"* was propagated into **two more files** before it was checked |

---

## The split, and the mechanism — this is the finding

| class | n | findings |
|---|---|---|
| **THIS** — manufactured by this pass's repair | **7** | 1, 2, 10, 11, 12, 13, 14 |
| **PRE** — present before it | **7** | 3, 4, 5, 6, 7, 8, 9 |

### The two new fences did their jobs. They are not the binding constraint.

Both closed the gap they were named for, and both were watched RED before they were trusted:

- `value_change_guard` detects the `16 → 23` class `repair_leak_guard` is structurally blind to — asserted
  against the verbatim fence directly, and it **found `seeding-spec.md:497` on this very repair**, a site
  the verbatim fence had already passed GREEN.
- `--audit-commit` admits an audit and refuses a repair wearing the same flag.

And the commit-time post-condition **caught this repair twice before the commit** (`D-M257x-49-5`): an
anchor drifted to a blank line, and a leak into `ai-readiness.md`. Both were repaired pre-commit.

**So why did the induced term go 2 → 7?** Because the seven induced findings partition into exactly three
classes, and **not one of them is mechanically reachable**:

| induced class | n | findings | why no fence reaches it |
|---|---|---|---|
| **paraphrase leak** — the twin states the claim in *different words* | 3 | 10, 11, 12 | `repair_leak_guard` is verbatim; the removed and surviving forms share no 8-token run (*"is the engine for"* vs *"powers"*; *"back-fills it into"* vs *"back-fills then"*). **This is the limit `D-M257x-48-4` pinned and `value_change_guard`'s own docstring re-pins.** Both fences behaved exactly as documented |
| **overshoot in NEW text** — a correction that over-corrects | 3 | 1, 13, 14 | There is **no old form to leak and no value to diff**. The defect is in prose that did not exist before the commit, so every diff-relative fence is silent by construction |
| **wrong mechanism, correctly cited** | 1 | 2 | Semantic. The anchor resolves, the construct is real, the citation is honest — and the sentence attaches it to the wrong branch |

> **The measured conclusion: the induced term is no longer dominated by the mechanical class TOK-02
> targeted. It is dominated by paraphrase and overshoot — and TOK-02 step 2's premise, *"the 8-of-9
> induced class cannot survive the commit"*, is now true of a class that has stopped being the
> majority.** Mechanising the mechanical half did not lower the total; it changed which class the
> remainder is made of.

**Three of the seven induced findings are the SEVENTH consecutive occurrence of one shape** — *the author
of a correction violating it while writing it* (1, 13, 14 are overshoots inside corrective text; 10 sits
four lines from a line the same diff edited). Six of nine passes have now produced a blocker in text
written to explain a correction. **That class still has nothing behind it but the author.**

## What this says about clause 5, honestly

Nine readings: `25 → 13 → 11 → 17 → 37 → 18 → 7 → 12 → 14`. §5 rule 22 (*a frozen instrument is not a
precise instrument*) holds and is reinforced: the run-to-run variance remains ~±5, comparable to the
residual being chased.

**The corpus term is real and is not shrinking by being repaired.** Seven pre-existing blockers this pass,
in files six of seven seats read top-to-bottom one iteration earlier — `roadrunner.md` (2), `hiring.md` (3),
`shared_libraries.md`, and a 31-vs-32 count in the passage whose own text says it has been wrong four times.
Seat F read 1,498 lines across six files and found **zero**, with its audited zeros named — so the finding
is not "auditors always find something."

This reading did not return zero. Clause 5 is **NOT MET**.

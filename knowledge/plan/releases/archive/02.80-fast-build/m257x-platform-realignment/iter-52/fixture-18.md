# Fixture — the UNION answer key (`#9 ∪ #10` = 18), captured BEFORE the repair

**Why this file exists.** §5 rule 21: *the fixture is perishable.* A corpus carrying a known, anchored,
independently re-verified answer key exists exactly once, and repairing it destroys the only thing that can
falsify an instrument built to protect it. iter-52 is the repair pass, so the key must be captured first or
it is gone.

**This key SUPERSEDES [`../iter-50/fixture-14.md`](../iter-50/fixture-14.md) as the repair target** — it does
not replace it as an artefact. `fixture-14.md` remains the answer key of reading **#9 alone** and stays
byte-untouched, as do `claim_twin/red/` (18), `claim_twin_iter48/` (4) and the iter-45 key. Those fixtures are
*supposed* to contain false claims; that is what makes them fixtures.

**Tree state this key is valid for:** rosetta `283ab1f` (the 40 audited files byte-identical to `47c9b7d`
and `57dfbfd` — every commit since is `knowledge/plan/**`).

**Ground truth, all thirteen clones re-read at iter-52 open and identical to the pair-experiment's:**
`app 5ba17044` · `app/studio aeec036a` · `platform 2adcf714` · `next-web-app bb3313bc` · `sentinel 88bc5592` ·
`storage 4ce8ece5` · `messenger fa47850d` · `cms ca50c817` · `graphql-wundergraph 60c229f3` ·
`roadrunner 87d8d443` · `jobsimulation 462343b0` · `studio-desk 14a5442a` · `ant-academy 9c3843cd`.

## How the union was formed

Not by hand (§5 rule 19's closing clause). The two readings' own blocker-ledgers are the source:

- reading **#9** — [`../iter-49/blocker-ledger.md`](../iter-49/blocker-ledger.md), indexed by
  [`../iter-50/fixture-14.md`](../iter-50/fixture-14.md) — **14** findings.
- reading **#10** — [`../iter-50/blocker-ledger.md`](../iter-50/blocker-ledger.md) — **7** findings.
- the matching is [`../iter-50/variance.md`](../iter-50/variance.md)'s published overlap table: **4 matched**
  (#10's row 3 covers #9's rows 10+11, which #9 had split by site; #10's 4 ≡ #9's 9; #10's 7 ≡ #9's 13).

`14 + 7 − 3 collapsed rows = 18`. Union **18**, coverage **78 %** of `N̂ ≈ 23`, against 61 % for #9 alone.

## The key

**Class:** **PRE** = present before iter-49's repair · **THIS** = manufactured by it.
**Seen by:** which reading(s) named it. Rows seen by only one reading are the ones a single-reading repair
would have left standing — **8 of the 18**, which is the whole reason this file exists.

| U# | anchor | the false claim, in one line | class | seen by |
|---|---|---|---|---|
| 1 | `external_services.md:668` | only the US LiveKit agent is suffixed — misses `anthropos-agent-chain` | THIS (overshoot) | #9 |
| 2 | `ai-readiness.md:406-419` | the active-cycle filter is named `keepStartedMembers`; it is `keepInCycleStep1` | THIS (wrong mechanism) | #9 |
| 3 | `security_compliance.md:76`,`:120` + `architecture_overview.md:299-300` | **31** org-filtered Ent schemas; it is **32** (`Organization`'s own `Policy()`) | PRE | #9 |
| 4 | `shared_libraries.md:145` | authn *"imported by … app"* only; the cms + jobsimulation husks import `colony/authn` directly | PRE | #9 |
| 5 | `roadrunner.md:87` | `SubmissionPackage` submits a *batch*; it is one multi-**file** submission, one token | PRE | #9 |
| 6 | `roadrunner.md:49` | Asynq pool for *batch* submissions; one poll task per single submission (contradicts `:97`) | PRE | #9 |
| 7 | `hiring.md:66-68` | `CreateOrganizationSimInvitationLink` is *the very call the seeder uses*; the seeder never calls it | PRE | #9 |
| 8 | `hiring.md:113-115` | with Clerk-only wiring the seeder *cannot* write the 5 positions; it still writes all 5 | PRE | #9 |
| 9 | `hiring.md:293-296` (recurs `:33-34`, `:157-158`) | *`jobsimulation.sessions` still exists*, unqualified in a local-dev section; no local stack creates that schema | PRE | **#9 + #10** |
| 10 | `external_services.md:672` | LiveKit + OpenAI Realtime *powers new sessions* — the claim retracted 4 lines from an edited line | THIS (paraphrase leak) | **#9 + #10** |
| 11 | `jobsimulation.md:123`,`:126` | the same refuted `flag_use_realtime_openai` claim — repaired in 1 of 3 places | THIS (paraphrase leak) | **#9 + #10** |
| 12 | `jobsimulation.md:102` | the mirror migration *back-fills then DROPs*; `SET "score"` = 0 hits, no `INSERT` in the file | THIS (paraphrase leak) | #9 |
| 13 | `hiring.md:209-210` | `token` is the *only* required-and-undefaulted column; **four** are | THIS (overshoot) | **#9 + #10** |
| 14 | `external_services.md` + `architecture_overview.md` + `security_compliance.md` | a **5th** EU-egress path via Studio-Room `openai`; selected by no shipped config | THIS (overshoot) | #9 |
| 15 | `external_services.md:604-618` | the `platform/.env` AI block names `OPENAI_API_KEY` / `OPENAI_ORG_ID` / `AZURE_OPENAI_ENDPOINT` / `AZURE_OPENAI_DEPLOYMENT`; the app reads **`OPENAI_KEY`** (33 hits) and `AZURE_OPENAI_ENDPOINT_URL` (24), and the one variable it reads is absent from the block | PRE | **#10 only** |
| 16 | `studio-room.md:388` | *"its **only** outbound API call is to the skills taxonomy service"*; the primary egress is the AI providers — contradicted by `:36` and `:261-266` of the same file | PRE | **#10 only** |
| 17 | `hiring.md:241-242` | *"**Neither table exists** … **both halves collapsed**"* — verbatim the claim retracted 83 lines above at `:157-159`; the two sites are a mutual contradiction | THIS (paraphrase leak) | **#10 only** |
| 18 | `dependency_map.md:19` | Storage *"No Postgres, no Redis"* with `-` in **Depends On**; `docker-compose.yml:213-217` declares `depends_on: redis`, `postgresql`. Contradicts the unedited twin `service_taxonomy.md:418` | THIS (overshoot) | **#10 only** |

**Split — 9 / 9.** THIS = **9** → #1, 2, 10, 11, 12, 13, 14, 17, 18. PRE = **9** → #3, 4, 5, 6, 7, 8, 9, 15, 16.

**Per-reading exclusivity:** #9-only = 10 (1, 2, 3, 4, 5, 6, 7, 8, 12, 14) · #10-only = 4 (15, 16, 17, 18) ·
both = 4 (9, 10, 11, 13).

## What this fixture is FOR, beyond being repaired

Three uses, and only the first is spent by iter-52's repair:

1. **The repair target.** Every row must be repaired by CLAIM, tree-wide (§5 rule 19), not by file.
2. **The post-condition.** `claim_twin_guard` derives its claim set from these same ledgers; at iter-52 open
   it reports **RED — 31 published sites / 19 unique**. The repair's post-condition is that number falling,
   with every survivor named and explained rather than silently absent.
3. **The pre-registration for reading #11/#12** (iter-53). TOK-03 pre-registered `N̂ < 12` and the induced
   term `< 4`. This key is what those readings will be scored against, and it is why iter-52 takes **no
   reading of its own** — a repairer who also reads is not blind to its own work.

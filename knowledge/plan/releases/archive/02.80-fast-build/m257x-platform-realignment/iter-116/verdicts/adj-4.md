# adj-4 — verdicts for seat E (readings r29-E + r30-E), M257x iter-116

**Trees read (stated per the brief's two-clone-set rule):** all platform clones under
`stack-demo/` at their own HEADs, re-verified at this reading's open with **no fetch** —
`app ad9f3c49` · `platform 0c91421d` · `messenger fa47850d` · `next-web-app 8297c684` ·
`sentinel f2c46190` · `storage 4ce8ece5` · `cms ca50c817` · `jobsimulation 462343b0` ·
`roadrunner 87d8d443` · `graphql-wundergraph 60c229f3` · `studio-desk 41ee3575` ·
`ant-academy 22df69dd` · nested `app/studio` + `cms/studio` both `aeec036a` ·
`stack-demo/rosetta-extensions 09d06070` (pinned) · `.agentspace/rosetta-extensions 1dc1eb82`
(authoring). **No booking in this seat-group is a fence-verdict or fence-config claim**, so
neither `rosetta-extensions` tree was load-bearing for any verdict here; no `wrong-tree`
rejection applies.

**Three-instrument discipline.** Every absence claim below (`posthog` in `internal/analytics/`,
`linkedin` in `internal/companysearch/`, Watermill consumers in `internal/worker/`, Liquid in
`messenger/internal/messenger/message/`, Redis key-writes in `messenger`) was checked with all
three mechanisms: `git grep` at the named ref, a raw filesystem `grep` over the same working
tree (which `.gitignore` cannot hide from and which sees untracked files), and a per-file NUL
byte census (`tr -dc '\000' < FILE | wc -c`) so no NUL-bearing file was silently skipped. Every
absence reported here returned the SAME answer on all three, with a working positive control in
the same command. No nested-repo blind spot applies — none of these paths is under
`app/studio` or `cms/studio`.

---

## Verdicts

r29-E B1 | `corpus/services/messenger.md:127` | UPHELD | IN-SCOPE | PREDICATE: Messenger uses Redis for scheduled-message storage.
   evidence: `messenger@fa47850d` constructs Redis in exactly two places, both the Watermill
   stream client — `cmd/root.go:108` (`redis.NewClient(os.Getenv("REDIS_ADDR"),
   redisStreamsIndex)`, feeding `pubsub.NewSubscriberServer` at `:113`) and `cmd/trigger.go:23`
   (the publisher CLI). I swept every non-test Go file for a key-write op
   (`.Set(|.HSet(|.SetEx(|.SetNX(|.ZAdd(|.LPush(|.RPush(|.Expire(|.SAdd(`) and got **one** hit,
   `internal/flow/organizations.go:225` `h.cache.Set(cacheKey, settingsMap, 5*time.Minute)` —
   and `cmd/root.go:27` imports `github.com/patrickmn/go-cache`, i.e. an in-process cache, not
   Redis. There is nothing to store either: `internal/rpcsrv/rpcsrv.go:25` and `:28` both return
   `connect.NewError(connect.CodeUnimplemented, …)` for `CancelScheduledMessage` and `Schedule`;
   the only scheduling anywhere is Brevo-side (`internal/messenger/brevo/brevo.go:289`
   `scheduledAt := time.Now().Add(30 * time.Second)`, handed to Brevo at `:305` as
   `SendSmtpEmail.ScheduledAt`). **Self-contradiction inside the same corpus file:**
   `messenger.md:17` says scheduling RPCs "are not yet implemented — they return Unimplemented"
   and `:112`/`:113` book both methods as `Unimplemented` stubs, while `:127` names a storage
   surface for that non-existent feature. Not a retraction (brief rule 5) — both halves are
   asserted live.

r29-E B2 | `corpus/services/messenger.md:90` | UPHELD | IN-SCOPE | PREDICATE: `messenger/internal/messenger/message/` performs Liquid rendering.
   evidence: `messenger@fa47850d` `git ls-tree -r internal/messenger/message/` = four files —
   `message.go`, `validator.go`, `errors.go`, `message_test.go`. I read `message.go` in full: it
   is `New` / `NewWithDefaultSender` / `DefaultSenderUser` / `ConvertProps` (a JSON round-trip
   into `map[string]any`). No template engine. Liquid at that ref lives in exactly three
   non-test files, none of them this package: `internal/flow/assignments.go:18,489`,
   `internal/flow/whitelabel.go:8,17`, `internal/messenger/console/console.go:16,26,33,71`. All
   three instruments agree the package has zero Liquid (git grep rc=1; filesystem `grep -ril
   liquid internal/messenger/message/` rc=1; NUL census 0 bytes on all four files). The same
   corpus file annotates the package that *does* render at `:98`
   (`whitelabel.go  Per-org whitelabel rendering (subject + body)`), so the gloss routes a
   reader away from the correct construct.

r29-E B3 | `corpus/services/messenger.md:115` | UPHELD | IN-SCOPE | PREDICATE: Messenger renders the email body with Liquid before the Brevo send.
   evidence: `messenger@fa47850d:internal/messenger/brevo/brevo.go:288-307` — `func
   (s *brevoSender) send(...)` builds `brevo.SendSmtpEmail{ Sender, To, ReplyTo, Bcc,
   ScheduledAt, TemplateId: templateId, Params: props }` and posts it at `:310` via
   `TransactionalEmailsApi.SendTransacEmail`. **No body is composed locally on that path** — a
   template id plus a parameter map go to Brevo and Brevo renders. The repo states the direction
   itself in the doc comment on the one function that renders locally,
   `internal/flow/whitelabel.go:13-15`: *"renderWhitelabelBody parses and renders a Liquid
   template using the provided data struct bound under a 'params' key, **mirroring how Brevo
   exposes template parameters as `{{ params.xxx }}` inside the email body**."* Local Liquid
   exists on exactly two paths — the whitelabel / CMS-custom composition, whose output is passed
   onward as a `custom_body` **param** (`flow/whitelabel.go:16-24`, `flow/assignments.go:486-489`;
   `brevo.go:36-77` shows templates 259 / 249 / 313 / 314 consuming it) — and the console sender
   used only when `BREVO_KEY` is empty (`console/console.go:71`). The corpus sentence states the
   special case as the general mechanism for messages that "carry user info, template ID, and
   template params", which is precisely the case where messenger renders nothing. Distinct from
   B2: B2 is a package-location proposition, this is a send-time-behaviour proposition, so they
   do not collapse.

r29-E B4 | `corpus/services/backend.md:377` | UPHELD | IN-SCOPE | PREDICATE: `app` has one Atlas migration pipeline; its versioned migrations live only in `terraform/migrations/`.
   evidence: The sentence names no ref of its own; the refs later in the same paragraph
   (`b948604f`, `2035f9a4`, **`ad9f3c49`** — the last explicitly called `origin/main` on
   2026-08-06) date the *"no top-level `migrations/` dir"* clause, and brief rule 1 keeps a pin's
   scope to its own claim. Graded at the checkout `ad9f3c49`, and cross-checked at the two other
   refs the paragraph names. `git show ad9f3c49:atlas.hcl` declares **two** envs: `env "local"`
   at `:6` (quoted correctly by the corpus) and, under the header *"sentinel-in-app v10.0 /
   M1001 — the SECOND Atlas pipeline"*, `env "sentinel"` at `:50` with
   `src = "file://terraform/sentinel/schema.sql"`, `dir = "file://terraform/migrations-sentinel"`,
   `revisions_schema = "sentinel"` and its own in-file comment: *"a SEPARATE env with its OWN
   source, migration directory, atlas.sum and revisions table."* The directory is real and
   populated — `terraform/migrations-sentinel/20260804151548_adopt_casbin_rules.sql` +
   `terraform/migrations-sentinel/atlas.sum` + `terraform/sentinel/schema.sql`. `app`'s CHANGELOG
   head records it: **v1.369.0 — 2026-08-05, "(atlas) second Atlas pipeline for the sentinel
   schema (M1001)" (`6827200`)**. I re-derived the ref window: `env "sentinel"` is present at
   **`2035f9a4` AND `ad9f3c49`** and absent only at the demo pin `b948604f` — so the claim is
   false at two of the three refs the paragraph itself names, including the one it calls
   origin/main. Operationally material and not merely terse: both documented commands apply the
   public pipeline only — `platform@0c91421d:Makefile:87` and `:95` run
   `atlas migrate apply --env local`, and `repos.yml` gives `app` `migrations: true`,
   `schema: public`. And the second pipeline is invisible corpus-wide: `git grep -n
   "migrations-sentinel\|sentinel-in-app\|M1001" -- corpus/ CLAUDE.md` returns **0**.

r29-E B5 | `corpus/services/backend.md:314` | REJECTED | — | PREDICATE: (would have been) `lab.v1.LabSessionService` is currently the third RPC handler registered in `app`'s `main.go`.
   evidence: The clause sits inside a *Recent Feature Additions (Q1-Q2 2026)* bullet whose head
   names its own ref — *"**AI Labs LabSession** (Phase B PR 2, #896)"* — and the verb is past
   ("Registered as a third RPC handler … after Users and Organizations"), i.e. a record of what
   that PR did. I resolved the ref: `app` `9ecade240` *"feat(labsession): add LabSession Ent
   entity + Connect-RPC service (Phase B PR 2) (#896)"*, and at that commit `main.go:239` is
   `mux := http.NewServeMux()` with exactly three registrations — `:240` users, `:243`
   organizations, `:249` LabSession. **The claim is true at the ref it names.** Separately
   re-derived the present shape at `ad9f3c49` (`main.go:1295` mux; registrations at `:1297`
   users, `:1298` organizations, `:1307` skiller, `:1315` jobsimulation, `:1322-1324` cms behind
   `if cmsRPCServer != nil`, `:1338` LabSession = **six**), which the same corpus file states
   correctly twice (`:30`, `:106`) — so the count a reader ends up with is right. The later
   sentence in the same bullet names its own ref (`b948604` v1.366.0) and is unaffected by this
   one, per brief rule 1's paragraph caveat. Note the sibling reading of this same seat (r30-E)
   independently **cleared** this exact clause as "historically correct"; my re-derivation agrees
   with that side.
   class: ref-discipline — a dated, past-tense change-log claim booked because later evidence
   contradicts it. That is the pin working (brief rule 2; also rule 7, historical anchor).

r30-E B1 | `corpus/services/backend.md:319` | UPHELD | IN-SCOPE | PREDICATE: `app/rpc.go` is the top-level Connect-RPC wire-up holding the implemented services.
   evidence: `git show ad9f3c49:rpc.go` is 383 lines and contains **only two handler
   implementations** — `type backendRPCServer` (`:22`, methods `GetUser` `:40`,
   `CanPerformFeatureAction` `:80`, `GetExperiencePoints` `:139`) and `type orgRPCServer`
   (`:183`, methods `GetOrganizationRoles` `:198`, `GetUserOrganizations` `:227`,
   `GetOrganizationDetails` `:255`, `GetOrganizationSetting` `:283`, `GetOrganizationSettings`
   `:347`). It contains **no** `http.NewServeMux`, **no** `mux.Handle`, **no** `rpc.NewServer` —
   there is no wire-up in the file at all. The wire-up is `main.go:1295`
   (`mux := http.NewServeMux()`), the six registrations at `:1297-1338`, served by
   `rpc.NewServer(mux, cfg.RPCPort, rpc.WithWriteTimeout(60*time.Second))` at `main.go:1353`.
   Four of the six services are implemented elsewhere — `internal/rpc/skillerrpc/skiller.go`,
   `internal/labs/session/labsession.go`, `internal/jobsimulation/rpcsrv/`,
   `internal/cms/rpcsrv/`. So a reader who follows "look there for the implemented services"
   finds 2 of 6 and none of the registration. **Self-contradiction inside the same corpus file:**
   `:30-31` anchors the mux at `main.go:1297-1338` @ `ad9f3c49` (and `:1185-1228` @ `b948604f`)
   and `:106` repeats the `main.go` enumeration — both of which I verified resolve exactly; this
   sentence points somewhere else. The bullet even names `internal/rpc/skillerrpc/` as
   `SkillerService`'s home later in its own text.

r30-E B2 | `corpus/services/backend.md:299` | UPHELD | IN-SCOPE | PREDICATE: `app/internal/worker/` holds the Redis Streams (Watermill) consumers.
   evidence: `app@ad9f3c49:internal/worker/worker.go:15` imports `github.com/hibiken/asynq`; the
   package is an Asynq server + cron registry (`const customerIOSyncCron = "*/10 * * * *"` at
   `:21`, `type asynqLoggerAdapter` at `:23`). Measured over the whole package: **20 of the 25
   files** match `asynq` (`git grep -cil asynq ad9f3c49 -- internal/worker/`; identical count
   from a raw filesystem `grep -ril`), while `watermill|pubsub` matches **one file, three lines**
   — `internal/worker/tasks/tasks.go:27` (import), `:39` (`publisher pubsub.Publisher` field),
   `:86` (constructor parameter) — all a *publisher* injected for emitting, never a subscriber.
   Absence re-checked with all three instruments: `grep -ril
   "watermill\|AddSubscriber\|NewSubscriberServer" internal/worker/` on the working tree rc=1,
   and a per-file NUL census over all 25 tracked files returned 0 bytes each, so nothing was
   skipped. The real Watermill/Redis-Streams consumers are where this same corpus file already
   locates them correctly — `buildStreamSubscribers` in `subscriber_wiring.go:203-248` applied by
   `main.go:1579-1581` (cited at `backend.md:68-70` and `:342`) — so the gloss both misnames the
   package and duplicates a role the document has already placed.

r30-E B3 | `corpus/services/backend.md:294` | UPHELD | IN-SCOPE | PREDICATE: `app/internal/templates/` holds email / message templates.
   evidence: `app@ad9f3c49:internal/templates/` is a **single file**, `templates.go`, 15 lines,
   exposing one function over `text/template`: `func Render(t *template.Template, name string,
   data any) (string, error)`. I enumerated its importers —
   `git grep -ln 'app/internal/templates' ad9f3c49 -- '*.go'` → **7 files, none of them email**:
   `internal/embeddings/embeddings.go`, `internal/jobrole/jobrole.go`,
   `internal/jobrole/jobrole_ant908_test.go`, `internal/rag/rag.go`, `internal/skill/skill.go`,
   `internal/skilltaxonomy/skill.go`, `internal/translation/translation.go`. I opened two to
   confirm the call shape: `rag.go:60` `prompt, err := templates.Render(tmpl, templateFileName,
   …)` and `skill.go:323` `prompt, err := templates.Render(tmpl, "skillClusters.tmpl", …)` —
   these are AI **prompt** templates. The transactional-email templates live under
   `internal/messenger/` (`message/`, `brevo/`, `console/`, `sender/`, `flow/`, `adapters/`,
   `aireadinessemail/` — enumerated by `git ls-tree -d ad9f3c49 internal/messenger/`), which this
   same document's merge banner names at `:15`.

r30-E B4 | `corpus/services/backend.md:257` | UPHELD | IN-SCOPE | PREDICATE: `app/internal/app/` is the component wire-up.
   evidence: `git ls-tree -r --name-only ad9f3c49 internal/app/` returns **32 files, all under
   `internal/app/users/`** — `importer/chain2d/{agent,extract,extractjson,patch,pipe,review,
   skills,skills_match,verify}.go`, `importer/model/*`, `importer/profile/profile2d.go`,
   `importer/steps/steps.go`, `models/`, `services/` — the résumé/profile 2-D import chain.
   `internal/app/users/services/service.go:13-27` is `type UserService` whose own constructor
   comment reads *"NewUserService takes its dependencies as explicit parameters (**no `*app.App`**,
   no grouping struct)"*. The construct the gloss describes is gone: `git grep -n "^type App
   struct" ad9f3c49 -- '*.go'` returns **rc=1, nothing**, and `1db398ea9` *"refactor: remove the
   \*app.App service-locator (explicit dependencies)"* is confirmed an ancestor of the graded ref
   by `git merge-base --is-ancestor`. The five surviving `app.App` mentions are all comments
   naming it as removed/not-used. Wire-up in this repo is `main.go` plus
   `internal/jobsimwiring/wiring.go` — which the same block already lists correctly at `:278`.

r30-E B5 | `corpus/services/backend.md:256` | UPHELD | IN-SCOPE | PREDICATE: `app/internal/analytics/` contains the PostHog integration.
   evidence: Absence re-derived with all three instruments at `ad9f3c49`. (1) `git grep -in
   posthog ad9f3c49 -- internal/analytics/` → **rc=1, no output**; positive control
   `git grep -cil organization` over the same path returns 11 of the 14 files, so the pipeline
   works. (2) Raw filesystem `grep -rin posthog internal/analytics/` → rc=1. (3) NUL census over
   all 14 tracked files → 0 bytes each, so no file was skipped by `grep -I`/`git grep`. The
   package's own header (`internal/analytics/analytics.go:15-33`) says what it is: *"This file
   owns the **Member activity analytics read model** … the summary KPIs, the stacked activity
   series, and the time split"*, sourced from `jobsimulation.sessions`,
   `public.skill_path_sessions`, `public.academy_chapter_progresses`, `public.lab_sessions`.
   PostHog is real in `app` and lives elsewhere — 20+ Go files including
   `internal/clerk/events/events.go`, `internal/invitations/feature_flag.go`,
   `internal/jobsimulation/ai/ai.go`, `internal/organization/manager.go`. The "internal
   analytics" half of the gloss is right; the PostHog half names a vendor the package has none
   of, and it is the searchable half.

r30-E B6 | `corpus/services/backend.md:269` | UPHELD | IN-SCOPE | PREDICATE: `app/internal/companysearch/` sources company data from LinkedIn.
   evidence: Absence re-derived with all three instruments at `ad9f3c49`. (1) `git grep -in
   linkedin ad9f3c49 -- internal/companysearch/` → **rc=1, no output**. (2) Raw filesystem
   `grep -rin linkedin internal/companysearch/` → rc=1. (3) NUL census over all three tracked
   files → 0 bytes each. The package is three files — `companysearch.go`,
   `companysearch_test.go`, `logodev.go` — and its one external source is **logo.dev**:
   `logodev.go:12-15` declares `type BrandSearcher interface { BrandSearch(...); LogoURL(domain
   string) string }`, `:17` `type LogoDevManager`, `:36-38` `LogoURL` returning
   `https://img.logo.dev/%s?token=%s`. LinkedIn import lives in `internal/linkedin/`, which the
   same code-map block lists separately and correctly at `:281`, so the gloss attributes a
   different package's subject to this one.

---

## PREDICATE ROLL-UP

P1 | Messenger uses Redis for scheduled-message storage. | anchors: r29-E B1 @ corpus/services/messenger.md:127
P2 | `messenger/internal/messenger/message/` performs Liquid rendering. | anchors: r29-E B2 @ corpus/services/messenger.md:90
P3 | Messenger renders the email body with Liquid before the Brevo send. | anchors: r29-E B3 @ corpus/services/messenger.md:115
P4 | `app` has one Atlas migration pipeline; versioned migrations live only in `terraform/migrations/`. | anchors: r29-E B4 @ corpus/services/backend.md:377
P5 | `app/rpc.go` is the top-level Connect-RPC wire-up holding the implemented services. | anchors: r30-E B1 @ corpus/services/backend.md:319
P6 | `app/internal/worker/` holds the Redis Streams (Watermill) consumers. | anchors: r30-E B2 @ corpus/services/backend.md:299
P7 | `app/internal/templates/` holds email / message templates. | anchors: r30-E B3 @ corpus/services/backend.md:294
P8 | `app/internal/app/` is the component wire-up. | anchors: r30-E B4 @ corpus/services/backend.md:257
P9 | `app/internal/analytics/` contains the PostHog integration. | anchors: r30-E B5 @ corpus/services/backend.md:256
P10 | `app/internal/companysearch/` sources company data from LinkedIn. | anchors: r30-E B6 @ corpus/services/backend.md:269

**Dedup notes.** No two bookings in this seat-group collapse. P2 and P3 are the nearest pair —
both concern Liquid in `messenger` — but they are different propositions (where the rendering
code lives vs. whether rendering happens before the Brevo send), and a seat booking one would
not write the other's sentence; brief's "do not collapse two anchors that merely look similar"
applies. P6–P10 all sit in one construct (`backend.md`'s *Key directories* map at `:245-301`)
and share a defect **class** — *a package gloss naming a technology or role the package does not
have* — but a class is not a predicate: each names a different package and a different
falsehood, and each has a separate repair.

---

BOOKED=11 UPHELD=10 REJECTED=1 IN-SCOPE-UPHELD-BLOCKERS=10 DISTINCT-IN-SCOPE-PREDICATES=10

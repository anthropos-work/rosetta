# Adjudicator 1 — verdicts for seats A · B · C · D (reading #21)

Re-derived from the platform clones at the ground-truth shas. Corpus graded at rosetta `964b7a3a`
(`git diff e858fd45..HEAD` touches `knowledge/plan/**` only — every corpus line number the seats
cite is unchanged). Every anchor below was opened, not grepped-to.

---

## Seat A

### A B1 | `corpus/services/chronos.md:27` | REJECTED | IN-SCOPE | Doc's 8080/8081 are the binary's defaults, not a compose-published pair

- evidence: `stack-demo/platform` @ `045857c^` `docker-compose.yml:339-368` — the chronos block does
  publish `"8500:8500"`/`"8501:8501"` with `PORT=8500`, `RPC_PORT=8501`, `REDIS_STREAMS_INDEX=4`, so the
  seat's compose reading is exact. But that is an *override*, not a refutation of a default. The
  platform-wide binary convention is `cmp.Or(os.Getenv("PORT"), "8080")` / `("RPC_PORT"), "8081"` —
  measured in **all four** cloned peers: `messenger/cmd/root.go:63-64`, `jobsimulation/cmd/root.go:77-78`,
  `storage/cmd/root.go:45-46`, `sentinel/cmd/root.go:47` (HTTP only). `chronos.md:27` sits under
  *Architecture & Code Map* and is the same fact class `storage.md:210-211` states explicitly
  (*"binary default 8080, overridden in compose"*); the doc's own `:209` CLI default
  (`--server` `http://localhost:8081`) is consistent with it.
- class: **mis-read** — the seat's own report concedes *"I cannot refute it directly"*; chronos is in no
  clone set, so the claim is unmeasurable-but-corroborated, not false. A compose override does not make
  a stated binary default wrong.

### A B2 | `corpus/architecture/external_services.md:208-211` | REJECTED | IN-SCOPE | rext anchors resolve exactly in the pinned per-stack clone

- evidence: both rext clones are on the ground-truth sheet. In `stack-demo/rosetta-extensions`
  @ `ab81527a` (tag `fast-build-m257x-iter-58` — the copy a stack actually consumes)
  `stack-injection/gen_injected_override.py:84` **is** `DIRECTUS_DATA_CONSUMERS = ("cms", "backend")`
  and `:669-670` **is** the `if with_directus and name in DIRECTUS_DATA_CONSUMERS:` /
  `env.append(f"DIRECTUS_BASE_ADDR=…")` pair — byte-for-byte the constructs the corpus names. In the
  authoring copy `5fb0915e` the same constructs sit at `:86` and `:698-699` (`:84` is a comment, `:669-670`
  the backend-CORS comment block) — the drift the seat measured, which I reproduced. The substantive
  predicate (*"re-points every service in `DIRECTUS_DATA_CONSUMERS`, which is `("cms","backend")`"*) is
  **TRUE in both copies**; only the offsets differ.
- class: **mis-read** — a citation exact at one of the two designated clones of one repo is anchor drift
  between witnesses, not a false or wrong-construct claim. (Checked the alternative: at the authoring
  copy's `f9ac72f`, cited two paragraphs below for the *fix*, neither `:84` nor `:669-670` names these
  constructs either — so the per-stack pin is the only reading under which the passage resolves, and it
  resolves perfectly.)

---

## Seat B

### B B1 | `corpus/services/README.md:21` | REJECTED | IN-SCOPE | Index echoes the fenced map's own `prod: live-standalone` roadrunner verdict

- evidence: `corpus/architecture/platform-migration-status.md:90` is the `roadrunner` row and its **prod
  column reads `live-standalone`**, defined at `:47` as *"its own process, still on the traffic path."*
  `README.md:21`'s *"so it **does** still deploy"* is that column restated, and the same sentence names
  the same contradiction the map and `roadrunner.md:19-20` name (*"the one row where prod and the
  platform's own `repos.yml` contradict each other"*). The *"recorded, not resolved"* language scopes the
  prod-vs-clone-set disagreement, not the deploy fact. The cms asymmetry the seat charges is
  **evidence-driven, not arbitrary**: cms carries a contrary artifact in its own repo — `6efa1d5` deleted
  `.github/workflows/build-production.yml` (verified: absent at `cms` origin/main `f38c0c4a`, present at
  the checkout `ca50c817`) — while `stack-demo/roadrunner` still ships `.github/workflows/build-production.yml`
  and `terraform/main.tf:19` `service_desired_count = 1` (verified at `87d8d443` = origin/main).
- class: **mis-read** — the index agrees with the authority it cites; there is no second measured fact on
  the roadrunner side to trigger the *report both, assert neither* posture.

### B B2 | `corpus/services/README.md:37` | UPHELD | IN-SCOPE | Lists a folded `roadrunner` domain in `app`; no such package, and `:11`/`:20` deny it

- evidence: `README.md:11-12` enumerates the folded set as **seven** (skiller, skillpath, jobsimulation,
  cms, storage, messenger, customerio-sync — no roadrunner); `README.md:20` states in bold
  *"`roadrunner` is the eighth, and it is different: orphaned, not merged-and-undeployed"*; `README.md:37`
  then writes *"the folded skiller …, skillpath, jobsimulation, cms and **roadrunner** domains"* — adding
  roadrunner and dropping storage/messenger/customerio-sync, a third mutually inconsistent enumeration in
  one 88-line file. Ground truth: `git -C stack-demo/app ls-tree --name-only HEAD internal/ | grep -i road`
  → **exit 1**, and identically at `origin/main` — `app/internal/roadrunner/` exists at neither ref. The
  fenced map says the same in the `app` row (`platform-migration-status.md:87`:
  *"**`app/internal/roadrunner/` does not exist**"*).

### B B3 | `corpus/services/coursebuilder.md:77-79` | UPHELD | IN-SCOPE | `cost` documented on the SSE wire; the code filters it off by design

- evidence: `stack-demo/app` @ `b948604f`,
  `internal/web/backend/coursebuilder/handler.go:2709-2717` — `case cb.EventCost:` returns `"", nil` under
  a comment reading *"the terminal cost readout is COGS and MUST NOT reach the customer SSE stream …
  filtered here at the wire boundary by returning an empty event name"*, and `:1474-1481` is the guard
  `if name != "" { writeSSE(w, flusher, name, payload) }`. So the doc lists on the wire contract the one
  event the platform fenced *off* it. Second half re-derived independently from `renderEvent`
  (`handler.go:2604-2775`): the live names are `text`, `score`, `patch_applied`, `patch_skipped`, `stage`,
  `outline`, `progress`, `preview_ready`, `draft_kept`, `error`, `translation_ready`, `rebuild_required`,
  `steering_received`, `steering_applied`, plus `session` (`:1458`) and `done` — **16**; the doc lists 13
  and omits four, two of which (`steering_received`/`steering_applied`) `handler.go:3182-3183` documents
  as the observable contract of `POST /sessions/:id/steer`.

### B B4 | `corpus/services/storage.md:215` | UPHELD | IN-SCOPE | Table records `ENVIRONMENT` as never set by compose; the block set it

- evidence: the column's contract is explicit — header `:208` *"Compose value"*, banner `:203-204`
  *"The middle column records what `docker-compose.yml` set on the `storage` service block."* The block set
  it at **both** refs the surrounding prose names: `platform@0dab54d docker-compose.yml:119`
  `- ENVIRONMENT=development`, and `platform@2adcf71 docker-compose.yml:206` `- ENVIRONMENT=development`.
  The `(empty)` convention is used correctly in the neighbouring rows — `STORAGE_S3_BUCKET` and
  `SENTRY_DSN` genuinely do not appear in that block — which is what makes this row a wrong value rather
  than a loose convention. The `HISTORICAL` banner scopes relevance, not accuracy.

### B B5 | `corpus/services/ai-readiness.md:305` | UPHELD | IN-SCOPE | `urls.ts:52` names `ORGANIZATION_FEEDBACK_URL`, not `AI_READINESS_URL`

- evidence: `stack-demo/next-web-app` @ `bb3313bc`,
  `packages/core-js/src/constants/urls.ts:50` `export const AI_READINESS_URL = '/ai-readiness';`;
  `:52` is `export const ORGANIZATION_FEEDBACK_URL = '/enterprise/organization-feedback';` — a different
  constant for a different route. Checked the other candidate ref so a pin cannot rescue it: at
  next-web-app `origin/main` `8297c684` the constant is at `:51`. No ref makes `:52` correct. Class-2
  wrong-construct anchor; the rest of the sentence (`useNavbarSections.tsx:4`, `:398-400`, `:547`)
  resolves, which is what lets a reader run past it.

### B B6 | `corpus/services/ai-readiness.md:46` | UPHELD | IN-SCOPE | Self-citation `:458` used as evidence is a blank line

- evidence: read `corpus/services/ai-readiness.md` with literal line numbers — `:456-457` close the
  *"Route 2 …"* paragraph, **`:458` is empty**, `:459` opens the `> **✅ CORRECTED M219 …**` blockquote.
  The sentence asserts a property of that line (*"contradicted `:458` of this same file, which is already
  in the past tense"*), and the line has no content to carry it. The citation is the stated *evidence* for
  a repair, so it is not decoration. Control: the sibling self-citations in the same seat's files
  (`cms.md:211` → `:44-47`, `cms.md:235` → `:70-71`) both resolve, so this is one drifted anchor and not a
  convention.

---

## Seat C

### C B1 | `corpus/services/backend.md:29` | UPHELD | IN-SCOPE | "All of their Connect-RPC surfaces are on `app`'s mux" is false for storage/messenger/skillpath/roadrunner

- evidence: enumerated the mux from `stack-demo/app` `main.go:1185-1228` @ `b948604f` — **six** handlers:
  `UsersService` `:1187`, `OrganizationsService` `:1188`, `SkillerService` `:1196`,
  `JobSimulationService` `:1204`, `CMSService` `:1212-1214` (conditional), `LabSessionService` `:1228`;
  identical set at `origin/main` `2035f9a4` (`:1297`, `:1298`, `:1306`, `:1314`, `:1323`, `:1338`). Of the
  **eight** services the banner's own table (`:9-16`) quantifies over, only three are represented. The
  falsity is **non-vacuous**, not merely a missing surface: `storage` declares a real
  `StorageService` Connect surface (`storage/internal/rpcsrv/rpcsrv.go`, registered
  `storage/cmd/root.go:63`) and `messenger` a `MessengerService` (`messenger/internal/rpcsrv/rpcsrv.go`,
  registered `messenger/cmd/root.go:84`) — neither appears on `app`'s mux at either ref.
  `SkillPathSessionService` → `git grep -c … -- '*.go'` exit 1 at **both** `b948604f` and `origin/main`
  (control: `UsersService` → 3 files). The same file contradicts itself at `:66` (LabSession fifth of five
  unconditional handlers) and `:68` (*"There is no `SkillPathSessionService`"*), and
  `shared_libraries.md:70` says *"skillpath and roadrunner were REMOVED, not re-hosted."*

### C B2 | `corpus/services/backend.md:33-34` | UPHELD | IN-SCOPE | Producer+consumer stream set omits `backend` while asserting exhaustiveness

- evidence: derived the set, not the sum. `git grep -n "NewPublisher(" b948604f -- '*.go'` → five
  application publishers — `main.go:287` (`serviceName`, defaulted `"backend"` at `main.go:211-215`),
  `main.go:637` (`SKILLPATH_STREAM`), `main.go:1039` (`CMS_STREAM`),
  `internal/jobsimwiring/wiring.go:127` (`AI_USAGE_STREAM`), `:180` (`JOBSIMULATION_STREAM`) — plus the
  DLQ (`internal/deadletterqueue/dead_letter_queue.go:38`, not an application stream).
  `AddSubscriber` → `main.go:1274, 1276, 1285, 1303, 1305, 1320`. Both-ways set = **five**:
  backend, skillpath, jobsimulation, cms, ai_usage; `SKILLER_STREAM` consumer-only (`:1276`). `:33-34`
  names four, drops `backend`, and then closes the set — *"`skiller` is NOT a fifth member"* — making the
  omission an explicit false exhaustiveness claim. Self-contradiction confirmed at the authority the
  banner itself cites: `backend.md:264` enumerates *"`backend`, `skillpath`, `jobsimulation`, `cms` —
  plus the `AI`/`ai_usage` usage stream"*, an incompatible set. `:264` is the correct one.

### C B3 | `corpus/services/backend.md:236` | REJECTED | IN-SCOPE | True at PR #896, inside a section headed "Q1-Q2 2026"

- evidence: the bullet names its own ref — *"AI Labs LabSession (Phase B PR 2, **#896**)"* — under
  `## Recent Feature Additions (Q1-Q2 2026)` (`backend.md:225`). Resolved the PR:
  `git log -S "labv1connect.NewLabSessionServiceHandler" --reverse -- main.go` → **`9ecade240`**,
  *"feat(labsession): … (Phase B PR 2) (#896)"*, **2026-05-29**. At that commit
  `git show 9ecade240:main.go | grep -n mux.Handle` returns exactly three lines — `:240` Users, `:243`
  Organizations, `:249` **LabSession**. The claim *"Registered as a third RPC handler in `main.go` after
  Users and Organizations"* is literally, exactly true at the ref its own bullet names. It is false today
  (`:1228` at `b948604f`, last of six) only because the skiller/jobsim/cms folds landed afterwards.
- class: **ref-discipline** — a dated claim booked because newer evidence contradicts it. The pin is
  working. (Noted: the seat itself rated this medium-low and named the reading it then declined to take.)

---

## Seat D

### D B1 | `corpus/services/jobsimulation.md:136` | UPHELD | OUT-OF-SCOPE | Corpus still mandates seeding two DROPPED mirror tables — but every offending site is outside scope

- evidence: the ground truth confirms the booked file is the **correct** side.
  `stack-demo/app/terraform/migrations/20260729133514.sql:58-61` is the *"5. Drop the mirrors."* comment,
  `:62` `DROP TABLE "local_jobsimulation_sessions";`, `:63` `DROP TABLE "local_skill_path_sessions";`;
  `grep -ciE 'INSERT[[:space:]]+INTO'` on that file → **0**; the only `CREATE TABLE "local_*"` statements in
  the whole migration tree are `20240312152917.sql:46` and `20240527131926.sql:2`, years earlier — nothing
  recreates them. The contradiction is live and verified at:
  `corpus/ops/demo/session-clone-spec.md:175` (*"**Co-write the manager MIRROR** (`public.local_jobsimulation_sessions`)
  or the manager scoreboard is blank"*, landmine #1), `corpus/ops/seeding-spec.md:388-391`,
  `corpus/ops/verification.md:248` and `:260` (a `≥ 40` cardinality gate on the dropped table), and
  `CLAUDE.md:387` (the *"generalized manager-view MIRROR trap"*).
- **scope**: I enumerated every corpus occurrence of either table name
  (`grep -rn 'local_jobsimulation_sessions\|local_skill_path_session' corpus/ CLAUDE.md`). Inside the
  audited partition the corpus is **uniformly on the correct side** — `jobsimulation.md:132-136`,
  `skillpath.md:81-86` (an explicit RETRACTION), `hiring.md:20-24, 170, 253, 264, 326-327, 402` all record
  the drop. Every mandating site is in `corpus/ops/**` or `CLAUDE.md`. Real finding, correctly booked,
  **OUT-OF-SCOPE**; it does not enter `N`.

### D B2 | `corpus/services/jobsimulation.md:203-204` | UPHELD | IN-SCOPE | Labels `7177374` as `origin/main`; at the real origin/main the anchor names another construct

- evidence: `stack-demo/app` — `git rev-parse origin/main` → **`2035f9a4`**;
  `git rev-list --count 7177374..origin/main` → **33** (the seat said six; the arithmetic is wrong, the
  predicate is not). `git show origin/main:main.go | sed -n 216p` → `if strings.Contains(dsn, "://") {`;
  `func main()` is at **`:229`** there. So a reader who resolves the label as written lands on an
  unrelated construct. The sha-named halves are all true and I confirmed each: `:216` **is** `func main()`
  at `7177374` and at `9d00a313`, and `:212` at `b948604f` — which is why this is a false *currency label*
  rather than a stale line number. Cross-file contradiction confirmed:
  `corpus/architecture/platform-migration-status.md:87` names *"`app` **origin/main `2035f9a`** (post-v1.369.0)"*,
  repeated at `:89`, `:90`, `:92`. Two corpus files assign two different shas to one moving ref; only
  `2035f9a4` is right.

### D B3 | `platform-migration-status.md:89` · `:270` · `corpus/services/jobsimulation.md:12` | UPHELD | IN-SCOPE | Three sites flatly assert "cms has not moved" / "sits untouched"; the cms row refuses exactly that

- evidence: one predicate, three anchors. Located the rows precisely: `:87` = `app`, **`:88` = `cms`**,
  `:89` = `jobsimulation`, `:90` = `roadrunner`. `:88` reads *"**M257x iter-92 — cms has since taken an
  M810 step, and it points the OTHER way** … this repo now holds **two measured facts pointing opposite
  ways** … **report both, assert neither**."* `:89` then says *"**Do not generalise M810 from this row:**
  `cms` **has not moved**"*; `:270` says *"while `cms` **sits untouched** at `service_desired_count = 0`"*
  inside the *"the teardown phase … is uneven"* passage — i.e. a claim about M810 progress, which is the
  exact axis on which cms **has** moved; `jobsimulation.md:12` repeats the flat clause in a second file
  where `:88`'s qualification is absent. Both halves re-derived in the `cms` clone: `terraform/main.tf:39`
  is `service_desired_count = 0`, **and** `6efa1d5` (2026-08-04, *"chore(ci): drop build-production — the
  cms ECR repository is decommissioned (M810)"*) deletes `.github/workflows/build-production.yml`
  (`--stat` = 18 deletions; the file is present at the checkout `ca50c817` and **gone** at origin/main
  `f38c0c4a`, of which `6efa1d5` is an ancestor). `cms.md:78-84` states the fence in words
  (*"**Do NOT extend 'the module block has not moved' to 'the rollback path is intact'**"*), which the
  three imperative-voice sentences cross.

---

## Deduplication

No two bookings share a predicate. The nearest pair — B B2 (a *folded roadrunner domain* in `app`) and
C B1 (the *universal RPC-mux* claim) — touch roadrunner but assert different things at different anchors;
they are kept separate. D B3 is **one** predicate carried at **three** anchors across two files and is
counted once.

---

BOOKED=14 UPHELD=10 REJECTED=4 IN-SCOPE-UPHELD-BLOCKERS=9

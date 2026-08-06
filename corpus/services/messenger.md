# Messenger Service

> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05) **deleted the compose service outright** and
> removed `messenger` from `repos.yml`. There is nothing to start locally and no `messenger` profile to
> opt into; the domain runs in-process inside `backend`, gated by **`MESSENGER_ENABLED`**, which is unset
> — and therefore **off** — on a developer machine. **The measured two-sided detail is in the fuller
> banner below**; this one exists so a reader who opens the file learns it before the first paragraph.

## Role & Responsibility

Messenger is the platform's **centralized notification subsystem** — a standalone service until the v9.0
fold, a domain inside `backend` since. It sends and schedules transactional emails, using **Brevo**
(formerly Sendinblue) as the delivery backend and **Liquid** templating for the bodies.

Callers don't talk to Brevo directly — they **publish Redis Stream events that the messenger flow consumes** (`messenger/internal/flow/flow.go:72-104` @ `fa47850` adds a subscriber on the `backend` stream with **21** handlers — 22 `pubsub.EventHandler(…)` lines of which one, `OrgJobSimulationAssignmentPastDueHandler`, is commented out *"not implemented"*; `app` runs that same subscriber now, on messenger's own consumer group). Messenger then decides whether to send immediately, apply org-level whitelabel branding, or skip the message entirely based on per-domain notification rules (e.g., it skips job-sim emails for stale/re-triggered sessions). (Scheduling RPCs exist in the proto but are not yet implemented — they return Unimplemented.)

> **⚠️ Nobody "fires a Messenger RPC", and nobody ever did.** The standalone *exposed* a `MessengerService`
> Connect-RPC surface, but **no service in the platform ever constructed a client for it**: `MESSENGER_RPC_ADDR`
> occurs in **no** repo — measured at each clone's own named ref, and including the two NESTED repos a
> host-ref grep cannot see (`app/studio` and `cms/studio`, both `anthropos-studio-room` @ `aeec036`).
> And `git -C stack-demo/platform log -S 'MESSENGER_RPC' --oneline 0c91421d | wc -l` returns **0**
> commits over the platform's whole 121-commit history (**positive control**, same repo and ref:
> `-S 'SKILLER_RPC'` returns **7**; add `--all` and both become 8-and-0, the 8th being `464dfe3` on a
> non-main branch — so state the scope). The earlier control figure here was **3**, which is reproducible
> at no repo, ref, spelling or scope; corrected M257x iter-96. The RPC traffic ran the other way — messenger
> called **out** to `backend` on four addresses, all `http://backend:8083` (`messenger/cmd/root.go:118-142`).
> **Compose set those four on the `messenger` service block and nowhere else**, so deleting the block
> deleted them: since `838d907` **no compose file sets any `*_RPC_ADDR` at all**, and there is no
> messenger process left to hold a client. Corrected M257x iter-85, re-derived at iter-87; the same
> sentence stood in [`README.md`](README.md) and was repaired in the same pass.

> **⚠️ MERGED INTO `app` — the v9.0 fold landed 2026-08-04**, in the same program that folded
> `storage`, and on the same morning; the container went the next day, with storage's and
> customerio-sync's, at platform `838d907` (merged `0c91421`, 2026-08-05, *"drop the storage,
> messenger and customerio-sync containers"*). Re-derived at `app` `9d00a313` v1.367.0 /
> `messenger` `e9421c6`.
>
> **Which tree settles which row, stated because getting it backwards inverts the answer.** `e9421c6` is
> `messenger`'s **`origin/main`**, and the **prod** row is a claim about *production infrastructure*, so
> origin/main is the tree that settles it. The demo's `messenger` clone is a frozen legacy checkout **7
> commits behind** at `fa47850d`, where `terraform/main.tf:19` reads `service_desired_count = 1` — **one,
> not zero** — and there is no `:29` declaration at all; grading the prod row against that clone returns
> the opposite verdict. The **local** row, by contrast, is a claim about a stack and is settled by platform `0c91421`
> + the clone set. (Noted M257x iter-102.)
>
> | side | measured |
> |---|---|
> | **prod** | `messenger/terraform/main.tf:29` `service_desired_count = 0` **@ `messenger` `e9421c6` (= `origin/main`)** — the compute is stopped, the cms precedent again. Image and task definition stay declared: this is the rollback path, a one-line revert plus an apply (`:27-28`, the in-file comment saying exactly that) |
> | **consumer** | `app` imports `internal/messenger/{flow,adapters,sender}` (`app/main.go:15`, `:62`, `:63` @ `app` **`ad9f3c49`**) and runs a **second subscriber server on messenger's OWN Redis consumer group** (`:1450`, wired at `:1471`, sender at `:1473` via `msgsender.NewFromEnv`). It does **not** merge messenger's handlers onto app's own subscribers — it **takes the group over**, so Redis keeps the cursor and there is no gap. The group name is messenger's, and it is a literal on purpose: the standalone read `cmp.Or(os.Getenv("SERVICE_NAME"), "messenger")` and nothing in terraform ever set `SERVICE_NAME` for it (`:1416-1421`). **Every anchor in this row moved between `9d00a313` and `2035f9a` without the code moving** — re-derived M257x iter-87, ref re-stated M257x iter-100, and re-stated **again** at iter-102. **This cell is the reference specimen of the stale-currency-pin class:** iter-100 correctly replaced the bare `9d00a313` — but replaced it with the *moving label* `origin/main` rather than with the sha that label then denoted, and on 2026-08-06 `origin/main` advanced `2035f9a4 → ad9f3c49` (5 commits) and the citation went stale a second time. A pin is a pin; a branch name refills on the platform's release cadence. The anchors themselves are unaffected — `git -C stack-demo/app diff --stat 2035f9a4 ad9f3c49 -- main.go` is **empty**, so every line number above resolves at both refs |
> | **local** | **nothing left.** No compose service, no `repos.yml` entry, so `make init` does not clone it. `838d907` deleted both; the rollback path is production-side only. (At `0dab54d` it was still `repos.yml:21-23` and a compose block behind an opt-in profile — that release also dropped it from `all`, because running it beside `backend` puts **two consumers on one Redis group**. That hazard is now unreachable rather than merely discouraged.) |
>
> **Everything below this banner describes the standalone service** — the frozen repo where the
> templates, the Brevo client and the notification rules still live, and the source `app` ported.
> Read it as the description of a rollback target, not of anything a local stack can start.

> **You cannot run it locally at all.** There is no messenger container and no `messenger` profile —
> both were deleted at `838d907`. `make up` never started it even before that (it was opt-in), and
> since the v9.0 fold `backend` does its work in-process, gated by `MESSENGER_ENABLED`
> (`app/env_guards.go:61` @ `app` **`ad9f3c49`** — `origin/main` and the demo's build pin on 2026-08-06;
> identical at `2035f9a4`), which defaults to **off** on a developer machine. **The ref is not optional
> on this anchor** (M257x iter-102): `env_guards.go` **did not exist** at the demo's former pin
> `b948604f` (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so this citation used to
> resolve at no ref the document named.

## Architecture & Code Map

* **Codebase**: `messenger` — repo `git@github.com:anthropos-work/messenger`. **Not cloned by `make init`**: `838d907` removed the `repos.yml` entry with the container. Clone it by hand to read the pre-merge source; the live code is `app/internal/messenger/`
* **Language**: Go 1.25
* **Framework**: Connect-RPC
* **Email backend**: Brevo via `getbrevo/brevo-go v1.1.3`
* **Templating**: `osteele/liquid v1.8.1`
* **Ports**: `8200` (HTTP) and `8201` (Connect-RPC) were published 1:1 by the compose block until `838d907` deleted it; **nothing publishes them on a stack now**. The binary's own defaults are 8080/8081 (`cmd/root.go:63`, `:64`)
* **Profile**: **none — there is no `messenger` compose service.** Platform `838d907` (merged `0c91421`, 2026-08-05) deleted the service block outright, and the `messenger` profile is gone with it. It was opt-in for its whole life; `0dab54d` had already dropped it from `all`, because running it beside `backend` puts two consumers on one Redis group.

### Key directories

```
cmd/                         Entrypoints
internal/
  rpcsrv/rpcsrv.go           Connect-RPC handler (Send, Schedule, Cancel)
  messenger/
    messenger.go             Top-level Messenger dispatcher
    brevo/                   Brevo client
    console/                 Console sender for local dev
    message/                 Message types + Liquid rendering
  flow/
    flow.go                  Notification-flow dispatcher
    assignments.go           Assignment notification rules
    cms.go                   CMS studio-task simulation completion rules (success/failure email)
    jobsimulations.go        Job-simulation completion / reminder rules
    organizations.go         Org invitation / membership rules
    organizations_db.go      Org DB lookups (read-only)
    whitelabel.go            Per-org whitelabel rendering (subject + body)
```

### Whitelabel rendering (2026-Q2)

Recent work in v0.34.0 added **whitelabel support**: when an org has custom branding (logo URL, custom invitation templates), Messenger renders subject and body separately so the Brevo send can include the org's logo and styling. The org lookup uses a **read-only Postgres connection** (`READONLY_DB_CONNECTION`, formerly `COPILOT_DB_CONNECTION` — see `cmd/root.go:147`) so the rendering path doesn't contend with the write-heavy backend load.

## Interface Discovery

### Connect-RPC (`MessengerService`)

| Method | Purpose | Status |
|--------|---------|--------|
| `Send(message)` | Send an email immediately | Implemented |
| `Schedule(message, schedule_for)` | Schedule a future email | Stub — returns `Unimplemented` (`internal/rpcsrv/rpcsrv.go:25-30`) |
| `CancelScheduledMessage(id)` | Cancel a previously scheduled message | Stub — returns `Unimplemented` (`internal/rpcsrv/rpcsrv.go:25-30`) |

Messages carry user info, template ID, and template params; the body is rendered through Liquid against those params before the Brevo send.

### What triggers Messenger?

Most messenger sends are reactive — driven by **Redis Streams** events on the `jobsimulation`, `cms` and `backend` streams. The stream *names* outlived the services: since the merges they are published from inside `app` (e.g. the `CMS_STREAM` publisher at `app/main.go:1095`, and the whole subscriber stream binding at `:1478-1484` @ `app` `9d00a313` v1.367.0), so there is no separate producer service in compose behind any of them. The corresponding flow handlers in `internal/flow/` decide whether a stream event should produce an email, what template to use, and whether to apply staleness guards (e.g., for job-sim completions it drops the email if the session ended >2h ago, or has no end time and started >12h ago — `internal/flow/jobsimulations.go:140-151`). See `internal/flow/jobsimulations.go` for examples.

## Dependencies

* **RPC clients**: the binary still constructs four Connect-RPC clients — CMS, backend users + organizations, skiller, and jobsimulation (`cmd/root.go:118-142`) — each reading its address from the environment. At `0dab54d` compose supplied all four, and all four resolved to the one `backend` mux (`http://backend:8083`); `838d907` deleted the service block, and with it **every `*_RPC_ADDR` compose ever set** — there are now zero across `docker-compose.yml`, `common.yml` and `.env_example`. So on a stack today those clients are neither constructed nor addressed: the process does not run. The `cms` and `jobsimulation` services they were named for went earlier, at `d11a403`; their surfaces are registered on `app`'s RPC server. Skill-path notifications arrive as Redis Streams events on the `backend` subscriber (`OrgSkillPath*` handlers in `internal/flow/flow.go:74-78`), not via a direct Skillpath RPC.
* **Downstream**:
  * **Brevo API** — outbound email delivery (`BREVO_KEY`)
  * **PostgreSQL** — read-only `public` schema access for org / whitelabel lookups
  * **Redis** — Watermill stream subscriber + scheduled-message storage

> **Staging safety**: if you ever restore a production DB dump into local staging, `BREVO_KEY` **must be blanked** in `platform/.env` before `make up` to prevent real customer emails from going out. See [staging_from_dump.md](../ops/staging_from_dump.md).

## Local Development

### Run in Docker — **not possible; there is no service to start**

`838d907` (merged `0c91421`, 2026-08-05) deleted the compose block, so no selection of profiles brings
a messenger container up. Asking for the retired token does not error either: compose **exits 0** and
starts only the always-on floor (`postgresql`, `redis`, `sentinel`), which looks like a live stack.

For the record, this is what it used to take, and why it was two flags rather than one: at `0dab54d`,
selecting the opt-in profile **alone** exited 1 — the block declared `depends_on: backend`, which that
selection did not include, so compose
rejected the project as invalid — you had to add the default profile alongside it. The `cms` and
`jobsimulation` `depends_on` entries had gone with those services at `d11a403`; `skillpath`'s went
when it merged into `app`.

**To exercise the mail path today you enable it inside `backend`**: set `MESSENGER_ENABLED=true` in
`platform/.env` (compose deliberately sets no value for it — pinning one there would override `.env`;
see the comment on the `backend` block, `docker-compose.yml:84-92`). Know what that does: `app`
attaches to messenger's **live** Redis consumer group, and a non-empty `BREVO_KEY` sends real mail.

### Run natively

```bash
cd messenger
go run main.go
```

The old first step — `cd platform && make dev S=messenger`, which stopped the container so the native
process could take its port — is moot: there is no container, and the repo is not a sibling clone any
more (`make init` does not fetch it), so `cd messenger` assumes you cloned it by hand.

Set `BREVO_KEY=""` to route through the **console sender** (`internal/messenger/console/`) instead of
hitting Brevo — emails print to stdout. (That fallback is standalone-only: `app` did **not** port it —
`app/main.go:295-300` @ `app` **`ad9f3c49`** (identical at `2035f9a4`): the condition is at `:295`
(`MESSENGER_ENABLED` **or** `CUSTOMERIO_SYNC_ENABLED` on with an empty `BREVO_KEY`) and the `log.Fatalf`
at `:296`. Ref re-stated M257x iter-102 — the citation was previously unpinned and present-tense, and
`:295` at the demo's former pin `b948604f` is a different construct.)

## Environment Variables

> **Nothing injects these any more.** The middle column is what the `messenger` compose block set,
> read at `0dab54d` — the last ref that had one. `838d907` deleted the block, so **every value below
> is now unset** unless you supply it yourself, and the four `*_RPC_ADDR` rows in particular are set
> by **no compose file at all**: they were the only four in the platform, and they went with the
> service.

| Variable | Value in the deleted compose block | Description |
|----------|---------|-------------|
| `PORT` | `8200` | HTTP port |
| `RPC_PORT` | `8201` | Connect-RPC port |
| `BREVO_KEY` | (empty) | Brevo API key. Empty → console sender. **MUST be empty for prod-dump staging.** |
| `REDIS_ADDR` | `redis:6379` | Redis address |
| `REDIS_STREAMS_INDEX` | `4` | Redis DB index for streams |
| `REDIS_WORKER_INDEX` | `0` | Was set in docker-compose (=0) but NOT read by the code — there is no worker pool / separate worker Redis index; only `REDIS_STREAMS_INDEX` is consumed (`cmd/root.go:107`). |
| `BACKEND_USERS_RPC_ADDR` | *(unset — was `http://backend:8083`)* | Backend RPC for user lookups |
| `CMS_RPC_ADDR` | *(unset — was `http://backend:8083`)* | CMS RPC. M809 re-pointed it off the standalone `cms` onto the `backend` mux at `d11a403`; `838d907` then removed the variable altogether. The earlier `http://cms:8091` was true at `2adcf71` only. `app`'s own comment at `app/main.go:1205-1211` (@ `b948604` v1.366.0) still says *"additive + DORMANT … until the M809 re-point"* and is **stale in `app`** |
| `JOBSIMULATION_RPC_ADDR` | *(unset — was `http://backend:8083`)* | Jobsimulation RPC, same history as the row above. The earlier `http://jobsimulation:8401` was true at `2adcf71` only; the husk container went at `d11a403` |
| `SKILLER_RPC_ADDR` | *(unset — was `http://backend:8083`)* | Skiller RPC surface — served by `backend` since the skiller→app merge, and reached in-process now that no consumer runs outside it |
| ~~`SKILLPATH_RPC_ADDR`~~ | *(removed earlier)* | **Gone from docker-compose** since skillpath was decommissioned into `app` ("skillpath-in-app", M502→M507) — only the residual `SKILLPATH_STREAM=skillpath` remains, on `backend`. Messenger never had a Skillpath RPC client anyway; skill-path data is read via the CMS client (`internal/flow/assignments.go:828`, in `getSkillPath`). |

> The binary's built-in fallbacks when the env var is unset are `PORT=8080` (`cmd/root.go:63`), `RPC_PORT=8081` (`cmd/root.go:64`), `REDIS_STREAMS_INDEX=2` (`cmd/root.go:107`).

## Testing

```bash
cd messenger
go test ./...
```

The flow handlers have unit tests (`assignments_test.go`, `jobsimulations_test.go`, `organizations_test.go`) covering the suppression / whitelabel branches.

## Related Documentation

* [Backend (app)](./backend.md) — main caller
* [staging_from_dump.md](../ops/staging_from_dump.md) — outbound-email kill switch
* [Dependency Map](../architecture/dependency_map.md)
* [Service Taxonomy](../architecture/service_taxonomy.md)

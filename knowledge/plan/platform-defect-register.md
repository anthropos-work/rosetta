# Platform-defect register

**Defects in the `anthropos-work` PLATFORM repositories that Rosetta's tooling found, and that Rosetta
cannot fix.** Zero platform-repo edits is a standing constraint, so every entry here is a **report**, not a
work item for this repo.

## Why this file exists

It was created at the **M256 close (v2.8 "fast build", 2026-07-30)** because the deferral audit found a
structural gap: **there was no platform-defect register anywhere in this repo.** M256 alone routed four
defects "to the platform", and all four lived only in that milestone's `decisions.md` — a file that flips to
`archived` at close and is never read again. A defect recorded inside a closed milestone has been *filed
where it cannot be found*.

The failure mode is the one the milestone spent 32 iterations on, one level up: a routing that looks
discharged because it was written down somewhere. M255's close had already been bitten by the sibling version
(four items routed to *"M255 harden resume"*, which was not a milestone and could not hold them).

## How to use it

- **Append, never rewrite.** An entry is evidence with a date.
- **Every entry carries `file:line`** for the deciding code. A defect report a platform engineer has to
  re-derive is half a report — and re-deriving it may mean re-doing the measurement that found it.
- **Distinguish MEASURED from INFERRED, per entry.** Several of these were found by driving a live demo; some
  siblings are reasoned from a shared code path and were never driven. The distinction is load-bearing and
  the entries say which is which.
- **Mark an entry `FIXED` with the platform commit** when it lands upstream. Do not delete it.

---

## Open

### `PLATFORM-2026-08-28-promote-path-authorizes-on-studio_task-org-never-on-the-simulation-tenant`
**Found:** 2026-08-28, adversarial review of the publish-endpoint design · **Filed here:** same day ·
**Repo:** `app` · **Status:** open · **Severity:** MEDIUM — a defence-in-depth gap that becomes a
cross-tenant write **in composition with** the sibling slug defect, not on its own ·
**Provenance: SOURCE READ at `app` `4bccda085`.** No request was issued. **An adversarial reviewer rated
this HIGH and claimed it lets any caller re-tenant-and-publish any simulation by naming its id; that claim
was checked and is WRONG — the calibration is recorded below, because a register entry that overstates its
own severity teaches a reader to discount the next one.**

**The gap.** The write primitive performs no read of the row it modifies:

```go
// internal/cms/directus/collections/jobsimulation.go:123-142
func (c *JobSimulationCollection) PublishJobSimulations(ctx, simulationId uuid.UUID, organizationId *string, superUser bool) (bool, error) {
	if !superUser && organizationId == nil { return false, fmt.Errorf("organization id is required") }
	var body = map[string]any{"private": false, "status": "published"}
	if organizationId != nil { body["private"] = true; body["tenant_id"] = organizationId }
	_, err := c.Query().PatchRaw(ctx, simulationId.String(), body)
```

It checks that an organization is **present**, never that it **owns** `simulationId`, and the patch sets
`tenant_id` and `status: "published"`. Its whole safety comes from its caller.

**Why it is NOT a standalone hole.** `StudioManager.PublishSimulation` (`studio/studioManager.go:440-470`)
gates it with `GetStudioTaskBySimulationID`, which for a non-super-user **does** scope on the organization
(`repository/studio.go:183-189`, `predicates = append(predicates, studiotask.OrganizationID(*organizationId))`),
followed by an explicit org-mismatch check (`studioManager.go:462-467`). Tenant A naming tenant B's
simulation gets no studio task back and stops at *"simulation not found"*. The denial-of-service variant
the reviewer proposed additionally requires creating a `studio_task` that names a victim's `simulation_id`;
no reachable path offers that (`SetSimulationID` appears only in generated ent code and in an unrelated
skillpath repository).

**Why it is still a defect.** The gate authorizes against **`studio_task.organization_id`** and never
against **`simulations.tenant_id`**. Those are two different records, and nothing keeps them in step. The
sibling defect
(`PLATFORM-2026-08-28-simulation-import-REPLACE-crosses-TENANTS-and-re-tenants-the-victim`) is a mechanism
that makes them diverge: it re-tenants a simulation row without touching the studio task that points at it.
After that, the gate consults a record whose organization no longer describes the row being written, passes,
and the promote re-tenants and publishes. **Two medium defects that compose into a cross-tenant write.**

**Fix.** Re-anchor promotion on the simulation's own stored `tenant_id` — read the row, compare, refuse —
rather than on the studio task's organization. Doing this also removes the composition risk independently of
whether the sibling defect is fixed first.

**Not measured:** whether any `studio_task.organization_id` and `simulations.tenant_id` pair currently
disagree in production. That is one join away and worth answering.


### `PLATFORM-2026-08-28-simulation-import-REPLACE-crosses-TENANTS-and-re-tenants-the-victim`
**Found:** 2026-08-28, designing the lodge publication endpoint · **Filed here:** same day · **Repo:** `app` ·
**Status:** open · **Severity:** HIGH — cross-tenant overwrite plus silent change of ownership ·
**Provenance: SOURCE READ, line by line, at `app` `4bccda085`.** No cross-tenant import was executed and this
entry does not claim one was. Every line below was read; the *consequence* is derived from those lines.

**The defect.** The simulation importer decides create-vs-replace with a lookup that has **no tenancy
filter**, then patches whatever row it finds.

```go
// internal/cms/directus/collections/jobsimulation.go:847
existingSim, err := c.GetJobSimulationBySlug(ctx, newSimInput.Slug, nil, true, nil)
//                                                          organizationIds ^^^^  ^^^^ superUser
```

`GetJobSimulationBySlug` (`:169-205`) applies `SetFilter("[slug][_eq]", slug)` and puts the
`tenant_id` / `private` filters inside `if !superUser` (`:195-205`). This call passes `superUser=true` and
`organizationIds=nil`, so **the only filter is the slug**, across every tenant, returning `Data[0]`.

What then happens with the foreign row (`:843-880`):

- `replace=false` → `return … fmt.Errorf("job simulation with slug %s already exists", …)` (`:851`).
  A **cross-tenant existence disclosure**: tenant A learns tenant B holds that slug.
- `replace=true` → `c.Query().PatchRaw(ctx, existingSim.ID.String(), newSimInput)` (`:874`).
  **Tenant B's simulation is overwritten by tenant A's content.**

And it is not only an overwrite. `JobSimulationInput` carries `TenantId *string \`json:"tenant_id,omitempty"\``
and `Private bool` (`jobsimulation_input.go:23`, `:36`), and the importer sets them from the CALLER's org
(`jobsimimport.go:403-410`). So the patch body re-tenants the victim's row: **B does not get a corrupted
simulation, B loses it to A.**

**Why this is live rather than theoretical.**

1. The legacy studio generator always passes `replace=true` — `internal/cms/studio/studioManager.go:409`.
   This is the path running in production today.
2. The slug is **caller-controlled**: it is read from the uploaded package, so a collision does not have to
   be waited for.
3. There is **no tenancy guard anywhere between the caller and the patch**. The importer's only tenancy code
   (`jobsimimport.go:403-410`) *sets* the input's tenant; it never checks the existing row's owner.

**The codebase already knows how to do this.** `applyPrivateCorpusFilters` (`jobsimulation.go:326`) exists to
scope a query to one tenant and is used by the LIST paths (`:442`, `:518`). The write path does not use it.

**Minimal safe fix (for the platform team, not this repo).** Scope the create-vs-replace lookup to the
caller's organization, and refuse rather than patch when the found row's `tenant_id` differs from the
caller's. **Honest caveat:** scoping the lookup changes behaviour for legitimate re-imports of *public*
simulations (`tenant_id` null), so the fix is not a one-line filter — the public case needs its own rule.

**Not measured here:** how many simulations currently share a slug across tenants, and whether any
production collision has already occurred. Both are answerable with a query against `directus.simulations`
and are worth answering before assuming the exposure is hypothetical.


### `PLATFORM-M257x-graphql-authz-middleware-FAILS-OPEN-and-REST-has-no-blanket-gate`
**Found:** M257x iter-120 (2026-08-07) · **Filed here:** iter-121 (2026-08-07) · **Repo:** `app` ·
**Status:** open · **Severity:** high (the platform's own source calls one half of it *"fail open"* and
attributes a shipped cross-tenant IDOR to it) · **Provenance: SOURCE READ, line by line, at `app`
`ad9f3c49` (`origin/main`).** No live request was issued and this entry does not claim one was.

**This is not the corpus defect.** iter-120 repaired four Rosetta claims that said *"Sentinel validates
**every** API request"* — that repair is done. What is filed here is the underlying platform property,
which the corpus repair does not address and Rosetta cannot fix.

**1 — The blanket GraphQL gate fails open, and the platform says so in a post-mortem comment.**
`internal/web/backend/graphql/graph/resolver_skiller_taxonomy_authz.go:53-66`:

> *"The skiller-in-app M207 port dropped every one of those guards and leaned on app's blanket
> `AuthorizationMiddleware` — but that gate is keyed on a `userId` operation variable and **FAILS OPEN**
> for taxonomy operations (which carry `{jobRoleId, organization}` and no `userId`): an authenticated
> caller with no org short-circuits to allow, and one with an org hits `errUnknownTarget` → allow. That
> left every taxonomy read/write reachable by any authenticated caller (**cross-tenant IDOR + privilege
> escalation**). … **Do NOT rely on the blanket gate for this surface — it fails open here.**"*

**2 — Six paths reach the resolver before the single Sentinel call.** The call is
`authorizationManager.OrgCheckUserPermission` at `internal/authorization/gqlauthz/gqlauthz.go:222`. Every
row below returns `next(ctx)` without it:

| `gqlauthz.go` | condition |
|---|---|
| `:160-161` | the operation failed to parse/validate |
| `:174-178` | no viewer **and** the op is `@public` / a federation query / dev introspection |
| `:190-191` | **the viewer has no active org** (`org == nil \|\| org.ID() == uuid.Nil`) |
| `:196-197` | `errUnknownTarget` — the operation carries **no `userId` variable** |
| `:202-203` | the target is nil, or the target **is** the viewer |
| `:209-219` | `@resolverAuthorized` — grants `authorization.Allow` **without calling Sentinel** |

The target is **one hardcoded variable name** — `grapqlTargetVar = "userId"`
(`internal/authorization/gqlauthz/target.go:11`), read as `ctx.Variables["userId"]` (`:20`). **An id
inlined as a document literal rather than passed as a variable does not reach Sentinel at all.** The
middleware's own doc comment (`gqlauthz.go:149`) says it *"gates every operation **on a viewer**"* —
authentication, which is what it does and not what a blanket authorization gate would do.

**3 — There is no blanket authz middleware on the REST surface.** `git grep -nE
'AuthzMiddleware|EchoAuthzMiddleware' -- '*.go'` at `ad9f3c49` returns **0**. Per-group, in
`internal/web/backend/backend.go`: `/api` (`:121-141`), `/ask` (`:171-173`), `/assignment-builder`
(`:194-196`) and `/admin/backfill` (`:210-212`) are `cors` + (`swagger`) + `authnEcho.EchoAuthnMiddleware`
only; `/coursebuilder` (`:229-232`) and `/credits` (`:273-276`) additionally carry `cbGate`
(`courseBuilderAccessGate`, `internal/web/backend/gate.go:27-49`), which **is** a Sentinel-backed group
middleware (`OrgCheckFeaturePermission(OrgFeatureMembersEdit, orgID)`). Everything else authorizes
per handler (`requireAdmin`, `checkTaxonomyWritePermission`, …) — the source states that intent inline at
`:183-185` and `:207-209`.

**What this does NOT claim.** Not a live vulnerability. The named taxonomy hole **was closed** by
restoring skiller's per-resolver checks (`checkTaxonomyWritePermission` et al, same file `:68-` onward).
The report is that the **general** property remains: a surface added to this codebase is authorized only
if its author writes the check, the blanket gate will not catch the omission, and **no enumeration exists
of which surfaces currently depend on a per-resolver or per-handler check**. That enumeration is the
thing a platform engineer can produce and we cannot.

**Also recorded, without editorialising** (it is a fact about the deciding line, not a recommendation):
the **admin impersonation mutation** — which mints a Clerk sign-in token for an arbitrary user by email —
gates on `permission.ActionObjectTaxonomy` / `permission.UserActionWrite`
(`internal/web/backend/graphql/graph/resolver_admin_audit.go:20-24`). That is a **taxonomy-write**
permission rather than a dedicated impersonation one. The sibling `adminAuditLogs` query at `:50-54` uses
the identical pair. Everything else on that path is present: a non-nil actor is required (`:25-28`), a nil
manager is refused (`:29-31`), and `internal/admin/impersonation/manager.go:1-4` writes an audit row on
**every** attempt, success or failure.

**Why it is filed rather than fixed:** every deciding line is platform source, and zero platform-repo
edits is a standing constraint of this milestone.

---

### `PLATFORM-M257x-dev-login-routes-mint-a-full-session-for-any-email-behind-one-NODE_ENV-boolean`
**Found:** M257x iter-119/120 (2026-08-07) · **Filed here:** iter-121 (2026-08-07) · **Repos:**
`next-web-app`, `studio-desk`, `ant-academy` · **Status:** open · **Severity:** medium-high (a
full-authentication bypass whose only control is one process-environment comparison) · **Provenance:
SOURCE READ across the whole clone set with three independent instruments** (`/usr/bin/grep -rn`,
per-clone `git grep` at each ref, and a Python `os.walk` — all three return the same five sites, per §5
rule 44). **Never driven.** No token was minted and no session was created.

> **⚠️ ONE PREMISE OF THIS FILING WAS RE-DERIVED AND IS FALSE — stated first, because the correction is
> the useful part.** The item was routed here as *"`ant-academy/code/app/api/dev/login-as/route.js:78` has
> **no `NODE_ENV` gate**."* **It does.** `route.js:34` refuses on `!DEV_LOGIN_ENABLED`, and
> `code/src/lib/devLogin.js:29` is `export const DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'`.
> The **ungated** minting site is a different one: `next-web-app/e2e/auth.setup.ts:72`, a Playwright
> runner file. The *"skips both factors"* comment is that file's too (`:57-62`), not ant-academy's. The
> two had been conflated. What follows is what is true.

**The five minting sites, enumerated** (a Clerk sign-in token is a one-shot ticket a browser exchanges via
`signIn.create({ strategy: 'ticket' })` for a **genuine session as the named user**):

| # | site | ref | gate |
|---|---|---|---|
| 1 | `app/internal/admin/impersonation/manager.go:101` | `ad9f3c49` | the product feature — Sentinel check + non-nil actor + an audit row on every attempt (see the entry above for the permission it uses) |
| 2 | `next-web-app/apps/web/src/app/api/dev/login-as/route.ts:79` | `8297c684` | `NODE_ENV !== 'production'` (`apps/web/src/lib/devLogin.ts:28`) |
| 3 | `next-web-app/e2e/auth.setup.ts:72` | `8297c684` | **none** — a test-runner file, never in an app build |
| 4 | `studio-desk/src/routes/dev.ts:83` | `41ee357` | `NODE_ENV !== 'production'` (`src/lib/devLogin.ts:33`) |
| 5 | `ant-academy/code/app/api/dev/login-as/route.js:78` | `22df69dd` | `NODE_ENV !== 'production'` (`code/src/lib/devLogin.js:29`) |

**What is worth a platform engineer's attention, stated as properties rather than as alarm.**

1. **Sites 2, 4 and 5 are UNAUTHENTICATED by design.** The same `DEV_LOGIN_ENABLED` boolean that mounts
   the route also adds it to the **public** route list in middleware — `ant-academy/code/proxy.js:178` and
   `next-web-app/apps/web/src/proxy.ts:56` both append `/api/dev/login-as` and `/dev/accept` to the
   unauthenticated set. So when the gate is open, `GET /api/dev/login-as?email=<anyone>` needs no
   credential of any kind.
2. **There is no allowlist and no shared secret.** The email is taken straight from the query string
   (`route.js:39`) and resolved with `users.getUserList({ emailAddress })` against whatever instance
   `CLERK_SECRET_KEY` points at. **Any** user of that instance can be assumed, not a fixed test account.
3. **The whole control is one comparison against a process-environment variable.** It is a *build-mode*
   gate, not a *deployment* gate: it holds because `next build`/`next start`, `npm start` and Vercel
   (Production **and** Preview) all set `NODE_ENV=production` — which the source comments state and which
   is true of those paths. It does not hold for anything that serves a dev-mode process, and **ant-academy
   is documented as running natively via `npm run dev`**. A second, independent control (an allowlist, a
   shared secret, or an explicit opt-in variable that is not `NODE_ENV`) would make the property depend on
   something other than how the process was started.
4. **Site 3 is a deliberate second-factor bypass against a real Clerk instance, and says so.**
   `e2e/auth.setup.ts:57-62`: the e2e account *"enforces 2FA (email_code as second factor); password
   signin returns `needs_second_factor` and never produces a session"*, and via a ticket *"Clerk treats it
   as fully authenticated and **skips both factors**."* That is a correct description of the primitive.
   It is recorded because it is the clearest statement in the codebase of what a leaked
   `CLERK_SECRET_KEY` buys, and because it means MFA enforcement is not a property of accounts that any
   holder of that key can reach.

**Not claimed:** that any deployed environment is currently exposed. Nothing here was driven, and whether a
given host serves a dev-mode process with a real secret key was not measured. **A sixth consumer is an
operator recipe, not code** — the `curl -X POST https://api.clerk.com/v1/sign_in_tokens` bypass Rosetta
documents in `corpus/ops/staging-clerk.md:58`, `staging_from_dump.md:384` and `staging-bringup.md:461`;
it mints against whichever instance the operator's key points at.

**Rosetta's side, already done:** `corpus/services/clerk-integration.md:40` claimed sign-in tokens are
minted *"only"* for admin impersonation. That word is retracted and the page now enumerates all five.

---

### `PLATFORM-M257x-compose-points-local-backend-at-the-PRODUCTION-S3-buckets`
**Found:** M257x iter-80 (2026-08-05) · **Filed here:** iter-102 (2026-08-06), after 21 iterations escalated
and undecided · **Repo:** `platform` · **Status:** open · **Severity:** high (a default `make up` writes into
production-named buckets)

`docker-compose.yml:82-83` @ `0c91421` — **re-derived at platform origin HEAD for this filing, still true** —
sets, **inside the `backend` service block**:

```
- STORAGE_S3_BUCKET=production-storage20240826131618541000000005
- STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
```

`backend` is the service the **default `core` profile starts**, so this is not an opt-in path: a plain
`make up` on a developer machine configures object storage against the **production** bucket names. The
compose file's own comment three lines above says these *"MUST be set"* because `app`'s boot guard *"only
fires outside a developer machine and `ENVIRONMENT=development` disarms it"* — i.e. the alternative the
platform authors were avoiding was silent writes to ephemeral container disk. **They chose the production
literal to avoid a silent local failure**, which is why this is a design disagreement to report rather than
an obvious slip to fix.

**Why it is filed rather than fixed:** the deciding line is **platform source**, and zero platform-repo edits
is a standing constraint of this milestone. Filing does not pre-empt a platform-side fix; it makes the
disagreement documented and permanent.

**Our side of it, stated so the two halves are not confused.**
`rosetta-extensions/stack-seeding/isolation/isolation.go:106` still registers `s3-private` as
`PerStackIsolated`. That registration is **ours** and is a separate, still-open question — re-classing it is
a rext change, not a platform one, and the two are not exclusive. **What was NOT done is assert our way out
of it:** iter-98 found `safety.md:207` claiming the `s3-private` registry entry had been *removed* when it
had not, and the assertion was **withdrawn rather than made true**.

**Provenance:** source read at three refs (iter-80 at the then-current platform ref; iter-98; re-derived
here at `0c91421` == `git ls-remote origin HEAD`). **Never driven live** — no upload was performed against
the configured bucket, and the entry does not claim one was. The claim is about **what the compose file
configures**, which is definitive by source; whether a given local run actually reaches those buckets
depends on credentials in `.env` and was not measured.

---

### `PLATFORM-M256-onboarding-step-not-resumed` — the org-prepared onboarding flow cannot resume
**Found:** M256 iter-31 (2026-07-30) · **Repo:** `next-web-app` · **Status:** open · **Severity:** high (the
first thing a real member does)

A member whose org pre-filled her profile confirms her role. The confirmation **persists** server-side
(`public.user_params.onboarding` gains a `role` step — verified in the DB on every attempt). She reloads
`/onboarding` and is back on **step one**, progress `0`, with no trace of what she confirmed; the screen is
**byte-identical to the pre-state**. Six fresh navigations across three browser sessions, hours apart, all
agree. She can never advance past the first step across a page load.

**Mechanism — one array, two consumers, opposite ends:**

```
packages/graphql/src/hooks/onboarding/useGetOnboardingStatus.tsx:25-27
    result.onboarding.steps?.sort((a, b) => sorterFn({ first: b.updatedAt, second: a.updatedAt }))
    → the array handed to the component is sorted NEWEST-FIRST
packages/ui/src/Onboarding/OnboardingUser.tsx:130-132
    const lastStep = reimport ? Import : steps?.[steps.length - 1]?.step;
    → takes the LAST element of a newest-first array, i.e. the OLDEST step ever taken
```

So `lastStep` is the *first* step the user ever completed, `managerImport` is true again, and the initial step
resolves to `Import` forever. The **host page reads the same array from index 0**
(`apps/web/.../onboarding/page.tsx:141-143`) and is therefore correct — which is why completion redirects
properly and nothing else ever looked wrong.

**Why nobody had seen it:** invisible for a NULL or single-element `steps` array (`length-1 == 0`, so both
readings coincide) — which is **every one of the 191 seeded users** and every hero any earlier iter could
reach. Only a multi-step array exposes it, and the only multi-step user in existence is a seat M256 iter-28
minted to reach the surface at all.

**Second defect on the same journey.** The prepared flow **cannot be completed on a demo**: one `Next` past
the skills screen reaches *"Add more skills"*, which renders *"We're having trouble loading your skills at the
moment"*, and its `Next` is **inert** (clicked five times, identical screen, progress stuck at 100).
`useClusterizeSkills` is the surface behind it.

**Provenance:** source read + six live observations. The mechanism above is a **source read**, stated as such.

---

### `DEFECT-M256-silent-forbidden-mutation` — a refused mutation renders nothing at all
**Found:** M256 iter-20, measured iter-23 (2026-07-30) · **Repo:** `next-web-app` · **Status:** open ·
**Severity:** high (it hid a real authorization gap for fifteen iterations)

A mutation the backend **refuses** is, from the user's side, indistinguishable from one that was never sent.
Reproduced deliberately on `demo-2` only (the `p3 admin → org:feature:taxonomy:write` grant revoked, the
journey driven, the grant restored byte-identically, `--policy-check` rc 0 afterwards).

**Measured across every channel a user or operator could learn from:** HTTP **200** with the error inside it ·
`[role=alert]` **present and EMPTY** · no `[role=status]` · antd `message`/`notification`/`form-item-explain`
all empty · the dialog **stays open with `Save` still ENABLED**, inviting an identical retry · URL unchanged ·
catalog total **49 → 49** · browser console says nothing about it · one **uncaught page error**.

**Two defects, one symptom:**

1. `packages/ui/src/JobRoles/Form/AddJobRole.tsx` `handleSubmit` handles exactly one error shape and
   `throw error`s the rest out of an async click handler — an unhandled rejection React renders nothing for.
   `onClose()` sits after the try/catch, which is why the dialog stays open. The empty `[role=alert]` is the
   **duplicate-warning slot**, never populated — *the app has one error surface here and it is reserved for a
   different error*. (`throw error;` from a catch appears **exactly once** in all of `packages/ui`.)
2. **The systemic half:** `apps/web/src/providers/Query.provider.tsx` sets
   `mutations: { onError: (e) => { captureException(e); PosthogClient.captureException(e) } }` — Sentry and
   PostHog, **no user surface**. Every mutation in the app is silent on failure unless it builds its own.

**And a dead contract that makes it look handled:** six mutations across four `hooks/organization/*` files
declare `meta: { error: '…' }` human-readable failure sentences. **No handler reads them** — there is **no
`MutationCache` anywhere** (0 occurrences); the only `meta.error` consumer is `QueryCache.onError`, which uses
it as a **Sentry tag**. So the strings are inert, on precisely the org-admin write set. *The authors wrote
failure messages and the framework never wired them up*, which is a more useful report than "the form is
silent" because it names a fix using a convention the codebase already believes it has.

**Suggested fix (not applied):** add a `MutationCache` whose `onError` reads `mutation.meta.error` and renders
it; replace `AddJobRole`'s `throw error` with the same path. Turns six dead strings live and gives every
future mutation a default surface.

**⚠️ Sweep residual — MEASURED vs INFERRED, stated because the claim is a negative.** Only `createJobRole` was
refused **LIVE**. The dead `meta.error` strings and the Sentry-only global handler are **definitive by
source**. That a refused tags-create / member-tag / settings-toggle would look **equally silent** is an
**inference** from the shared global handler — it was never driven, because each would have meant another
revoke/restore cycle on a stack later iters depended on. *Driving one sibling refusal closes that gap in one
revoke.*

---

### `PLATFORM-M256-keyrole-nondeterminism` — a succession key-role card appears nondeterministically
**Found:** M256 iter-26 (2026-07-30) · **Repo:** platform (succession ranking) · **Status:** open ·
**Severity:** low-medium (not a defect a presenter would see; it reddens automated batches)

A succession **key-role card**'s presence varies between page loads once its role has **2 occupants** —
measured **4 of 5 loads at occupancy 2** against **5 of 5 at occupancy 1**. Most plausibly a top-N ranking
with an unstable tiebreak.

**Cost, recorded because it is the reason this is filed rather than shrugged at:** two gate cycles. The
iter-14 cross-tenant negative control anchors its LIVENESS floor on that card, so it went RED reading
*"succession failed to compute for the contrast tenant"* — and a 45 s timeout did **not** fix it, because the
cause was the seed's role occupancy, not the clock.

**Mitigated our side, not fixed:** hero roles must be pairwise distinct within a story
(`playthroughs/e2e/tests/seed-facts-fence.unit.spec.ts`, mutant N1 RED). A batch can still redden on this.

---

### `PLATFORM-M256-cv-upload-never-parses` — a valid CV upload POSTs 200 and never advances
**Found:** M256 iter-18 (2026-07-29) · **Repo:** `next-web-app` / the import pipeline · **Status:** open ·
**Severity:** medium — **it is the reason a curated use case is `will-not-build`**

The profile-import CV route POSTs **200** for a valid PDF **and** for a docx alike, while the forward control
**never enables** (waited 100 s+). Measured with a purpose-built synthetic fixture
(`playthroughs/fixtures/synthetic-cv-sre.{pdf,docx}`, a wholly invented CV whose employers and school occur
nowhere in the seed, the taxonomy, or any real registry — so an assertion naming them can only be satisfied by
*that file having been imported*).

**Consequence for coverage, and it is the honest kind:** this is the deterministic alternative that would have
let `onboarding.enterprise-workforce-standard.UC1` be a Playthrough. With it blocked, the only advancing path
**scrapes a live public third-party profile** on a site that blocks automation — so the use case carries a
machine-checked `disposition: will-not-build` verdict instead (M256 `D104`/`D122`). **The two fixture files
ARE the evidence for that verdict**, which is why they ship despite having no consumer.

---

### `PLATFORM-M257x-directus-ext-logs-env` — a shipped Directus operation writes the whole environment to the log
**Found:** M257x iter-123 (2026-08-07) · **Repo:** `anthropos-work/directus` · **Status:** open ·
**Severity:** high — it is at the **deployed** pin

> **RE-DERIVED AT SOURCE, M257x iter-124/125 — the finding CONFIRMS and one anchor was WRONG.** Both repos
> were re-cloned at the exact refs this entry names (`directus` `d6325731`, `infrastructure` `13c248e6`).
> The line and the pin reproduce **verbatim**. **The environment inventory did not**: `services.tf:47-57`
> does *not* thread `SECRET` or `KEY` — those enter from the directus module's own terraform, and the real
> list is **longer** than the entry claimed. Corrected below, correction first, per `D-M257x-121-1`.

### The line — confirmed verbatim

`directus/extensions/directus-extension-youtube-meta/src/directus-extension-youtube-meta-operation/api.ts:9`
is a bare **`console.log(env);`** inside the operation handler:

```ts
export default defineOperationApi({                                    // :6
	id: 'ant-youtube-operation',                                         // :7
	handler: async ({ videoId }: Option, { env }) => {                   // :8
		console.log(env);                                                  // :9
		return await processUrl(videoId, true, env.GCLOUD_SERVICE_ACCOUNT); // :10
```

It is the **only** `console.log(env` in the entire repository (`git grep -n 'console\.log(env' -- '*.ts' '*.js'`
→ 1 hit). Present at `d6325731`, which `git tag --points-at HEAD` confirms **is** tag `v0.20.15`.

### The pin — confirmed, and it is named TWICE

`infrastructure/terraform/production/services.tf` @ `13c248e6` pins `v0.20.15` in **both** places that
select the deployed artifact — `:24` the module source (`?ref=v0.20.15`) and `:30` the ECR image tag
(`production-directus:v0.20.15`). So production runs exactly the tree that carries line 9.

### What the handler's `env` actually contains — CORRECTED, and it is worse than filed

A Directus operation handler's `env` is the instance's resolved configuration. The task definition
(`directus/terraform/main.tf` @ `d6325731`) injects **six secrets** via ECS `secrets` (`:224-249`):

| variable | source |
|---|---|
| `SECRET` | `aws_ssm_parameter.directus_secret` — Directus's own signing secret |
| `ADMIN_PASSWORD` | `aws_ssm_parameter.directus_admin_password` |
| **`DB_PASSWORD`** | `aws_ssm_parameter.db_password` — **the Postgres password** |
| `AUTH_GOOGLE_CLIENT_SECRET` | `aws_ssm_parameter.auth_google_client_secret` |
| `DB_SSL__CA` | `aws_ssm_parameter.database_ca` |
| `GCLOUD_SERVICE_ACCOUNT` | `aws_ssm_parameter.gcloud_service_account` |

**Plus `KEY` (`:111-114`), which is NOT in the `secrets` block** — it is a plain `environment` entry whose
value interpolates `aws_ssm_parameter.directus_key.value`, so it is materialised into the task-definition
JSON in clear as well as into the container env. That is a **second, independent** exposure of the same
value and it is recorded here because the re-derivation found it, not because it was looked for.

**What `services.tf:47-57` really threads** (the anchor this entry had): `database_host`, `database_name`,
`database_username`, `database_password`, `database_schema`, `database_ca`, `auth_google_client_id`,
`auth_google_client_secret`, `elasticache_cluster_primary_endpoint`, `gcloud_service_account`,
`admin_email` — the root module's *inputs*, not the container's environment. **The two are not the same
list, and conflating them is what put `SECRET`/`KEY` at the wrong anchor.**

**So: every invocation of the YouTube-meta operation writes the database password, the Directus signing
secret, the admin password and the Google client secret to the container log group.**

### Why this is reportable rather than merely noted

**Rosetta cannot fix it (zero platform edits binding), and it is FILED, not escalated** — the register is
where a platform defect this corpus measures goes (§5 rule 48). Three things make it reproducible by a
platform engineer without re-deriving anything: the value is not inferred (a literal `console.log(env)`),
the pin is not guessed (named twice in `services.tf`), and the environment inventory is now read from the
task definition itself rather than from the module's call site.

**Not measurable from here:** whether the operation is actually reachable in the production flow set, and
what the log group's retention and access policy are. Both change the exposure and neither is in a clone.
**Nothing in `corpus/` asserts the opposite** — checked at iter-125; the corpus's only statements about
this extension are `org-repos.md`'s inventory row, which now points here.

---

### `PLATFORM-M257x-akb-taxonomy-figures-contradict-measurement` — a customer-facing KB asserts a taxonomy figure this corpus refuted, in 14 unsourced places
**Found:** M257x iter-123, filed iter-125 (2026-08-07) · **Repo:** `anthropos-work/anthropos-knowledge-base` ·
**Status:** open · **Severity:** medium — no runtime impact; the exposure is **customer-facing collateral
and every engineer's editor**

`anthropos-knowledge-base` (AKB) carries a second, parallel platform-architecture corpus (six files under
`knowledge/`, ≈1,773 lines) covering this project's subject. It asserts **"60,000 skills … mapped to 18,000
roles"** in **14 places and cites no source in any of them**. The figure is **load-bearing in four
customer-facing competitor-comparison tables**, and AKB ships as a **Claude Code plugin** that injects
*"full Anthropos context (product details, architecture, …)"* into every engineer's editor on every
Anthropos repo.

**What was measured, and how.** A read-only production capture of the **public subset only**
(`organization_id IS NULL`), 2026-06-29, manifest `source: primary-read` / `public_only: true` /
`predicate: org-null`; both counts reproduce exactly against a live stack database with
`select count(*) … where organization_id is null`:

| | measured | AKB |
|---|---|---|
| public job roles | **22,470** | 18,000 |
| public skills | **42,790** | 60,000 |

**The two rows fail differently and must not be merged.** *"18K roles"* is **REFUTED** — public ⊆ total,
so production holds **≥ 22,470** and 18,000 is below the floor. *"60K skills"* is **UNVERIFIED, not
refuted** — a public-only capture cannot see org-private skills, so nothing measured supports 60K and
nothing rules it out. **A reconciliation that reports "AKB is wrong" over-claims on the skills row.**

**Candidate provenance for the 18K, offered as a lead and NOT as a measurement:** `public.job_role_embeddings`
holds **18,919** rows — a different table from the role count, and a plausible mis-transcription. Nothing
here can measure what AKB's author read, so this stays a hypothesis; it is recorded because it gives the
owner somewhere to start.

**This is filed, not escalated, and it is explicitly NOT one corpus correcting another.** On the
WunderGraph router's production residue **AKB was right and this corpus was wrong, in a fenced table** —
AKB reads the `infrastructure` repo this corpus had never cloned (M257x iter-123 cloned it and confirmed
AKB's reading). The two corpora have **different blind spots, not a ranking**: this one is authoritative
for measured local/runtime state and ops, AKB for `infrastructure`-derived production state and
product/GTM. **Neither cites the other.**

**Why it needs an owner outside M257x:** AKB is a different repository and outside this milestone's
two-repo scope (`rosetta` + `rosetta-extensions`), so no edit here can reach it. What M257x has done is
the part it can: the contradiction is stated **where a reader meets the figures**
(`corpus/architecture/shared_libraries.md#taxonomy-figures`) with both figures, both provenances and which
is measured; and `corpus/tools/toolchain_overview.md`'s **install recommendation now carries the warning
on the install line**, because that recommendation is what puts the refuted figure into an editor.

**Not measurable from here:** which of the four comparison tables have been published externally, and to
whom. That governs whether this is a documentation defect or a customer-communication one, and it is not
in any clone.

---

## `studio-desk` — the copilot's user avatar can never load (M257x iter-286, 2026-08-11)

**Reported by the user against a live demo:** *"in the right panel (studio copilot) we use avatar to
represent the user and the bot. The very first avatar shown there before the 'Assisting with: Scenario
design' is a not loaded image, tho the other one gets loaded properly."*

**Two independent defects in three lines of `app/sim-advanced-builder/builderAssistant.js`**, and either
alone is sufficient to produce a broken image:

```js
let userAvatarUrl = '/default_avatar.png';   // Default avatar
try {
  userAvatarUrl = userService.getUserPicture();
} catch (error) { console.error('Error getting user picture:', error); }
```

1. **The fallback is overwritten by a getter that can legitimately return nothing.**
   `userService.getUserPicture()` is `return this.clerk?.user?.imageUrl` — an optional chain, so a user
   with no image yields `undefined` **without throwing**. The assignment is unconditional and the `catch`
   never runs, so `avatar.src = undefined` and the browser requests a nonexistent path. This is the
   *"a check that skips reads exactly like a check that passes"* shape in assignment form: the error path
   is guarded, the **empty-success** path is not.
2. **The declared fallback asset does not exist.** `/default_avatar.png` appears **once** in the repo — at
   the line above — and there is no such file in `app/public/`, whose only avatar assets are
   `avatar_bot.png`, `avatar_bot_nobg.png` and the `avatar_traits/` directory. So even with defect 1
   fixed, the fallback lands on a 404.

**Why the bot avatar loads and the user's does not**, which is exactly what the user observed: the bot
image is a literal — `<img class="bot-avatar" src="/avatar_bot_nobg.png">` — and that file **is** in
`app/public/`.

**Not fixed here, and the reason is the constraint, not the difficulty.** This is `studio-desk` platform
source, and v2.8 holds 0 platform edits. It is **not demo-only** either: any user without a Clerk image
hits it in production, so a demo-patch would hide the symptom on the one surface where it does not
matter. **The fix belongs in the platform repo** — guard the assignment (`const url =
userService.getUserPicture() || FALLBACK`) *and* add the missing asset, or drop the fallback and render
initials.

**Measured at** `stack-demo/studio-desk`, `app/sim-advanced-builder/builderAssistant.js:741-754` and
`app/services/userService.ts:204-206`.

---

## PD-v28-A — ❌ RETRACTED 2026-08-14 — NOT a platform defect, a demo flag gate

**This entry was wrong and is withdrawn.** It called `EnterpriseWrapper.tsx:56` a *"stale redirect left
behind by the migration"*. It is nothing of the kind: it is the **gate itself**, and it is correct code.

    const v2Flag    = useFeatureFlagEnabled('flag_enable_assignments_v2');
    const v2Enabled = isDev || v2Flag === true;
    if (v2Resolved && !v2Enabled && V2_ONLY_PATHS.some(...)) router.replace('/enterprise/assignments');

`flag_enable_assignments_v2` is a **PostHog rollout flag**. A demo bakes no PostHog, so it resolves
`undefined` **forever**, `v2Enabled` is false on every demo, and the redirect correctly protects the
operator from the new (empty-until-backfilled) list. The five "already migrated" consumers I cited as
evidence use the v2 path because they live **inside** the gated surface — they were never the contrast
I read them as.

**The user identified this** (*"flag_enable_assignments_v2 should be on by default on all stacks type"*)
after I had filed it here. **Fixed where it belongs — in the demo tooling**, as two demo-patches shipped
in `v2.8.3-rext`, mirroring the M219 aireadiness and M232 interview flag patches that solve the identical
"no PostHog ⇒ no rollout gate" problem for their own flags.

**The lesson worth keeping:** a redirect to an older path is not evidence of staleness. I inferred a
migration defect from a *code shape* without reading what guarded it — and filed it against the platform,
where nobody could have acted on it, because the defect was ours.

**Superseded by** `next-web-assignments-v2-flag-{wrapper,home}` in `rosetta-extensions`.

<details><summary>Original (withdrawn) entry</summary>

## ~~PD-v28-A — `next-web-app`: the enterprise landing redirect still points at the RETIRED assignments page~~

**Reported** 2026-08-13 by the user against the rebuilt `demo1` (platform `766df6c` / `app` v2.0.1 /
`next-web-app` v2.141.0): *"the menu item on the left that says assign content brings me to the old page
`/enterprise/assignments` … it should bring me to `/enterprise/assignments-list`."*

**Measured**, `stack-demo/next-web-app`:

- `apps/web/src/app/(authenticated)/(verified)/enterprise/EnterpriseWrapper.tsx:56` —
  `router.replace('/enterprise/assignments')`. **This is the only occurrence of the old path in the app.**
- **Every other consumer already uses the new one:** `BuilderV2Container.tsx:47`
  (`LIST_PATH = '/enterprise/assignments-list'`), `AssignmentBuilderContainer.tsx:18`,
  `AssignmentDetailContainer.tsx:45` and `:77`, `AssignmentsListContainer.tsx:79`.
- Both pages exist — `enterprise/assignments-list/page.tsx` and `.../[planId]/page.tsx` are present.

So the migration to `assignments-list` landed everywhere **except the one redirect that decides where the
section opens**, which is why the user reaches the old interface without ever choosing it.

**Not a demo defect and NOT demo-patched.** Verified no demopatch manifest mentions `assignments-list`;
the stale redirect is platform source, and it hits production the same way. A demo-patch would hide the
symptom on the only surface where it does not matter. **The fix is one line in the platform repo.**

</details>

## PD-v28-B — `ant-academy`: clicking a module in the course-page list does not navigate

**Reported** 2026-08-13 by the user against the rebuilt `demo1`: on
`/chapters/coding-agents-landscape-intro/`, *"if i click a module from the list on the course page nothing
happens (it should move the ui to that subpage); if i click from the side menu the module it moves
properly."*

**Two list surfaces render the same modules and only one navigates** — the side menu works, the in-page
course list does not. That asymmetry is the diagnostic: the data is present and the routes resolve (the
side menu proves both), so the defect is in the in-page list's click handling, not in routing or content.

**Not investigated further here** and deliberately not demo-patched — `ant-academy` is platform source and
v2.8 holds 0 platform edits. Reproduce on the rebuilt demo, then compare the two list components'
handlers; the working side-menu path is the reference implementation.

**Measured environment:** `demo1` rebuilt 2026-08-13 on `billion`, ant-academy `7ae25e95b`, served at
`https://demo1.anthropos.work:13077`.

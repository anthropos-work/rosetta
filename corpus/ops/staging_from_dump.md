# Personal Staging from a Prod DB Dump

This guide takes a fresh machine that already has the basic Anthropos stack running (per [setup_guide.md](setup_guide.md)) and turns it into a **per-engineer staging environment** populated with real customer data so you can develop and test against the same shape of state as production — without touching any shared infrastructure or sending email to real users.

It is the bridge between "the stack starts" and "I can log in as my own admin account, see my org's members, simulations, and skill paths, and iterate with real data".

## When to use this

- You're a new engineer onboarding and want a faithful local environment.
- You're rebuilding a VM or replacing a dev box.
- You're spinning up a sales-grade demo or a customer-shaped sandbox.
- You want to test a feature against the same volume / shape of data as prod, not a hand-seeded toy DB.

## When NOT to use this

- You only need to develop a single isolated component (use the empty-DB flow from `setup_guide.md` and seed what you need).
- You don't have a prod dump or aren't authorized to handle one.

## Prerequisites

You should already have, per `setup_guide.md`:
- `platform/` plus the **three** `repos.yml` entries at `766df6c` — `app/`, `next-web-app/`, `studio-desk/` — cloned as siblings. ⚠️ **This read *"the four `repos.yml` entries — `app/`, `sentinel/`, …"* until M258 iter-18** — true at platform `0c91421`, and RETRACTED at `766df6c` (v11.0, 2026-08-11), which folded `sentinel` into `app` as `app/internal/sentinel/` and deleted its `repos.yml` entry with its compose service, so `make init` no longer clones it. (`skiller`, `skillpath`, `cms`, `jobsimulation`, `roadrunner`, `storage`, `messenger`, `customerio-sync` **and now `sentinel`** are all folded into `app` and no longer cloned or built; `graphql-wundergraph` went with the router at `2adcf71`.)
- `platform/.env` with `GH_PAT`, `CLERK_SECRET_KEY`, `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` filled in.
- `make up postgresql` succeeds and Postgres is healthy.

You also need:
- A recent prod dump, accessible at a known path (e.g. `~/prod_dump.dump`). ⚠️ **This read *"a `pg_dump` SQL file (plain SQL, not custom format)"* until 2026-09-03.** The scheduled backups are **`PGDMP` custom format**, and have been since at least the 2026-08 rotation — verify with `head -c5 <file>`, which prints `PGDMP`. See § 2 for where they live and how to restore one; the old `cat … | psql` path only works on a hand-made plain-SQL dump.
- A Postgres **17** client for the restore (`brew install postgresql@17`, or just `docker run --rm postgres:17-alpine`). The archive is written at **dump version 1.15** by a pg_dump 17 client, and `pg_restore` 15.x refuses it outright with `unsupported version (1.15) in file header`. The *server* can stay 15.x — new client against old server is the supported direction.
- Access to a **dev Clerk app** (NOT production Clerk). If you don't have one, create a fresh dev application in the Clerk dashboard and copy its publishable + secret keys.
- Your prod email address — the same email that exists in `public.users` of the dump.

> **DO NOT use prod Clerk keys on a dev/staging machine.** Any user-mutating action (create / ban / update metadata) you trigger from staging will execute against the live Clerk org. Always use a dev Clerk app.

---

## 1. Outbound-email kill switch (mandatory, do this FIRST)

A staging stack restored from a prod dump contains real customer email addresses in `public.users`. Many code paths trigger transactional notifications through the messenger subsystem → Brevo. If `BREVO_KEY` is set to a real value, those emails will go out to real people the moment you exercise the relevant flow.

> **⚠️ Corrected 2026-08-07 — three of the four examples this paragraph used to give do not exist.** It
> read *"(welcome emails, invitation flows, weekly recaps, password resets)"*. Measured in the messenger
> domain at `app` `ad9f3c498` (`git grep -in <term> -- internal/messenger`) and in the frozen `messenger`
> repo at `fa47850d9`:
>
> | Old example | Verdict | Measurement |
> |---|---|---|
> | invitation flows | **REAL** | `internal/messenger/flow/organizations.go` + `internal/messenger/flow/invitation_reminders.go`; goldens `organization_member_invited_whitelabel_{en,it}*`, `invitation_reminder_{workforce,hiring}_r1…r5_{en,it}` |
> | welcome emails | **NO SUCH FLOW** | `welcome` → **0** hits in `app/internal/messenger`. The only hit in the whole `messenger` repo is a test-fixture body string, `pkg/aireadinessemail/override_test.go:191` |
> | weekly recaps | **NO SUCH FLOW** | `recap` → **0** hits in `app/internal/messenger` and **0** in the `messenger` repo. The only weekly thing is the AI-readiness **manager** digest (`internal/messenger/brevo/brevo.go:265` *"Weekly manager digest"*, `internal/messenger/flow/organizations.go:152`) — a cycle report to managers, not a user recap |
> | password resets | **NOT OURS — Clerk's** | `password` → **0** hits in `app/internal/messenger` and **0** in the entire `messenger` repo. Widened to all of `app/internal/`: 16 hits, none a sender (a Clerk-webhook payload field at `internal/clerk/events/types.go:116`, LLM prompt templates, a meta-server config). The login surface is Clerk's own `<SignIn>` component — `next-web-app/apps/web/src/app/(unauthenticated)/login/[[...login]]/page.tsx:4` imports it from `@clerk/nextjs` and renders it at `:40` — so the reset mail is issued by **Clerk**, over Clerk's own transport. **Blanking `BREVO_KEY` does not suppress it.** Use a dev Clerk app (which this guide already mandates); that, not the Brevo key, is what keeps reset mail off real inboxes |
>
> The real Brevo-backed families, from the golden corpus in `internal/messenger/flow/testdata/emails/`:
> org member **invitations** + 5-slot invitation **reminders** (workforce + hiring), **assignment**
> mail (assigned / unassigned / completed / due-date-updated / past-due, for skill paths and job
> simulations), **job-simulation results** (passed / failed / dropped / interview-completed),
> **course-builder** mail (build completed / failed / course published), and the **AI-readiness**
> cycle set (invited / reminders r1–r5 / completed / cycle-launched / manager digest). Blanking
> `BREVO_KEY` is still mandatory — it is those flows it protects.

> **⚠️ There is no `messenger` container any more.** Platform `838d907` (2026-08-05) deleted it — messenger runs in-process inside `backend` (v9.0 "support-in-app"), gated by `MESSENGER_ENABLED`. Any `docker compose … messenger` command now fails with *no such service*; target **`backend`** instead. `MESSENGER_ENABLED` unset means off while `ENVIRONMENT=development`, but do not lean on that — a staging `.env` copied from elsewhere may well set it, which is exactly why blanking the key is the mandatory step.

**Blank `BREVO_KEY` in `platform/.env` and restart `backend` BEFORE running any flow that could enqueue a notification:**

```bash
sed -i.bak 's/^BREVO_KEY=.*/BREVO_KEY=/' platform/.env
docker compose -f platform/docker-compose.yml restart backend
```

Verify against **`backend`**:

```bash
docker compose -f platform/docker-compose.yml exec -T backend env | grep -E '^(BREVO_KEY|MESSENGER_ENABLED)='
# Expected: BREVO_KEY=  (empty); MESSENGER_ENABLED absent, or explicitly false
```

With the key blank, every API call to Brevo fails at the 401 layer and no email is delivered.

Apply the same caution to any other live-customer integration you don't intend to fire from staging:
- `HEYGEN_WEBHOOK_SECRET` (third-party webhooks — won't fire if not exposed publicly anyway, but blank it to be safe).
- `BUNNY_*`, `LIVEKIT_*`, `ELEVENLABS_*` (media / voice — these don't email but can incur cost or bandwidth charges; use sandbox keys if available).

The cheap heuristic: if disabling the integration would only break "email/notification went out", keep it disabled until you specifically need it.

### Also disable third-party analytics (page-load speed + don't pollute prod analytics)

The frontend's root layout eagerly loads Plausible, Google Tag Manager (GTM-PXRTBZK fans out to GA + LinkedIn pixel + Facebook pixel + Google Ads), BetterStack, and analytics.bellasio.com — that's ~10 third-party blocking requests on every page load. On staging this slows everything down over Tailscale and pollutes prod dashboards with staging traffic.

Set `NEXT_PUBLIC_DISABLE_ANALYTICS=true` in `platform/.env` and rebuild `next-web-app`. Also blank `POSTHOG_API_KEY` and `POSTHOG_SERVER_SIDE_KEY` for the same reasons:

```bash
cat >> platform/.env <<'EOF'
NEXT_PUBLIC_DISABLE_ANALYTICS=true
EOF
sed -i 's/^POSTHOG_API_KEY=.*/POSTHOG_API_KEY=/' platform/.env
sed -i 's/^POSTHOG_SERVER_SIDE_KEY=.*/POSTHOG_SERVER_SIDE_KEY=/' platform/.env

docker compose -f platform/docker-compose.yml build next-web-app
docker compose -f platform/docker-compose.yml up -d --no-deps next-web-app
```

Verify the analytics scripts are gone from the served HTML:

```bash
curl -s http://localhost:3000/login | grep -oE "plausible|googletagmanager|bellasio|betterstack|GTM-" | sort -u
# Expected: empty output
```

The flag is gated in `apps/web/src/app/layout.tsx` — production builds (which leave the flag unset) keep all analytics; only staging skips them.

---

## 2. Restore the prod DB dump

> ⚠️ **Rewritten 2026-09-03.** This section piped a plain-SQL file into `psql`
> (`cat ~/prod_dump.sql | … psql …`). The scheduled backups are **`PGDMP` custom format**, so
> that command fails on a real backup. Custom format is a net win — it takes `--schema`
> filters and `-j`, so you restore less and faster — but the ordering rule in 2b is new and
> **silently corrupts the restore if you skip it**.

### 2a. Fetch a backup

Scheduled RDS backups land in S3, roughly twice daily, ~2.5 GiB each:

```bash
export AWS_PAGER=""
BUCKET=production-db-backup20240909101503214700000001

# newest full dump
aws s3api list-objects-v2 --bucket "$BUCKET" --prefix rds/full/ \
  --query "reverse(sort_by(Contents,&LastModified))[:5].[LastModified,Size,Key]" --output text

aws s3 cp "s3://$BUCKET/rds/full/backup_<id>-<n>.dump" ~/prod_dump.dump
```

There is a `rds/per-schema/` tree alongside `rds/full/` if you only need one schema. If the
target machine has no AWS access, fetch on a machine that does and copy it across — verify with
`shasum -a 256` on both ends. (On macOS, `rsync` is **openrsync**: it rejects `--append-verify`
*and exits 0 while having transferred nothing*. Use `scp`, and always compare checksums.)

Read the archive before restoring — this also tells you which schemas are inside:

```bash
pg_restore --list ~/prod_dump.dump | head -20        # needs a pg17 client (see Prerequisites)
pg_restore --list ~/prod_dump.dump | awk '$4=="TABLE" {print $5}' | sort | uniq -c
```

### 2b. Create the schemas FIRST — including `extensions`

**`pg_restore --schema=X` does not create schema X**, and a missing `extensions` schema is the
single most expensive mistake in this guide. `extensions` holds the `vector` type. If `public`
is restored before it exists, **every table with a vector column fails to create** — and the
restore still reports success. You find out much later, when the app throws
`relation "similarities" does not exist` on the simulations library.

```bash
cd platform
docker compose up -d postgresql
until docker compose exec -T postgresql pg_isready -U postgres > /dev/null 2>&1; do sleep 1; done

docker compose exec -T postgresql psql -U postgres -d postgres < /dev/null -c "
  CREATE SCHEMA IF NOT EXISTS extensions;
  CREATE SCHEMA IF NOT EXISTS sentinel;
  CREATE SCHEMA IF NOT EXISTS directus;
  CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions;"
```

### 2c. Restore

`--schema` selects what you want and leaves the rest behind. The `zz_dropped_*` schemas are
unscrubbed shadow copies of decommissioned services — skip them unless you specifically need
them. Add `--schema=directus` only if you want CMS content locally (see § 5).

```bash
NET=$(docker inspect -f '{{range $k,$v := .NetworkSettings.Networks}}{{$k}}{{end}}' \
        $(docker compose ps -q postgresql))

docker run --rm --network "$NET" -v "$HOME:/dumps:ro" postgres:17-alpine \
  pg_restore --dbname="postgresql://postgres@postgresql:5432/postgres?sslmode=disable" \
    --no-owner --no-privileges \
    --schema=public --schema=sentinel --schema=extensions \
    --jobs=4 /dumps/prod_dump.dump 2> /tmp/restore.err
```

### 2d. Verify by counting tables, not by reading stderr

A partial restore is quiet. Compare what the archive holds against what landed:

```bash
pg_restore --list ~/prod_dump.dump | awk '$4=="TABLE" && $5=="public" {print $6}' | sort -u | wc -l
docker compose exec -T postgresql psql -U postgres -d postgres -At < /dev/null \
  -c "SELECT count(*) FROM pg_tables WHERE schemaname='public';"
```

**These two numbers must match.** If the second is lower, list the difference and restore just
the missing tables with repeated `--table=` flags — do **not** re-run a full `--schema=public`,
which re-inserts data into the tables that did land.

```bash
docker compose exec -T postgresql psql -U postgres -d postgres < /dev/null -c "
  SELECT 'users' tbl, COUNT(*) FROM public.users
  UNION ALL SELECT 'organizations', COUNT(*) FROM public.organizations
  UNION ALL SELECT 'memberships', COUNT(*) FROM public.memberships
  UNION ALL SELECT 'casbin_rules', COUNT(*) FROM sentinel.casbin_rules;"
```

**Expected warnings during restore** (all harmless):
- `ERROR: role "<name>" does not exist` for `backend`, `cms`, `skiller`, `chronos`, `customerio`, `simulator`, `sentinel`, `skillsgateway`, `skillpath` — these are GRANT/ALTER OWNER statements that no-op against a fresh box. Data tables load fine.
- `invalid command \unrestrict` at the very end — psql 15 doesn't recognize the `\restrict` / `\unrestrict` markers emitted by the pg_dump 16 client. Cosmetic.

If you see anything else (especially `relation already exists` collisions or `permission denied`), check `/tmp/restore.log` and clean the DB before retrying:

```bash
docker compose down
sudo rm -rf platform/data/postgresql
sudo mkdir -p platform/data/postgresql && sudo chown -R 1001:1001 platform/data/postgresql
docker compose up -d postgresql
```

(The `chown` is needed because the Bitnami Postgres image runs as uid 1001 and Docker creates bind-mount roots as root.)

---

## 3. Rebind your engineer account to the dev Clerk app

After restore, every `users.clerk_id` and `organizations.clerk_id` in the DB points at **prod** Clerk IDs that don't exist in your dev Clerk app. If you log in now, Clerk will authenticate you fine but the backend won't find your user record → blank/profile state, no admin context, all enterprise routes redirect to `/profile`.

The fix is a three-step rebind. Throughout, set:

```bash
export CLERK_SECRET=sk_test_…           # your dev app's secret key
export YOUR_EMAIL=stefano@anthropos.work  # the email you want to log in as
```

### 3a. Create your Clerk user, get its ID, set external_id to your DB UUID

```bash
# Create user in the dev Clerk app
curl -s -X POST https://api.clerk.com/v1/users \
  -H "Authorization: Bearer $CLERK_SECRET" \
  -H "Content-Type: application/json" \
  -d "{\"email_address\":[\"$YOUR_EMAIL\"],\"password\":\"<a-strong-password>\",\"first_name\":\"…\",\"last_name\":\"…\",\"skip_password_checks\":true}"
# Save the returned "id" — it looks like user_3DI…
export CLERK_USER_ID=user_3DI…

# Get your DB user UUID
export DB_USER_UUID=$(docker compose -f platform/docker-compose.yml exec -T postgresql \
  psql -U postgres -d postgres -At -c "SELECT id FROM public.users WHERE email='$YOUR_EMAIL';")
echo "DB UUID: $DB_USER_UUID"

# Set external_id on the Clerk user → DB UUID
curl -s -X PATCH "https://api.clerk.com/v1/users/$CLERK_USER_ID" \
  -H "Authorization: Bearer $CLERK_SECRET" \
  -H "Content-Type: application/json" \
  -d "{\"external_id\":\"$DB_USER_UUID\"}"

# Rewrite your DB row to point at the new Clerk user
docker compose -f platform/docker-compose.yml exec -T postgresql psql -U postgres -d postgres -c "
  UPDATE public.users SET clerk_id='$CLERK_USER_ID', updated_at=now() WHERE email='$YOUR_EMAIL';
"
```

### 3b. Enable Organizations on the dev Clerk app, create matching dev orgs

Most Anthropos pages use Clerk's Organizations feature. By default, dev apps ship with it disabled.

```bash
curl -s -X PATCH https://api.clerk.com/v1/instance/organization_settings \
  -H "Authorization: Bearer $CLERK_SECRET" \
  -H "Content-Type: application/json" \
  -d '{"enabled":true,"max_allowed_memberships":50,"creator_role":"org:admin","admin_delete_enabled":true}'
```

For each org you're an admin of in the prod dump, create a matching Clerk dev org and remap the DB:

```bash
# List orgs you're admin of (output: db_org_uuid|name lines)
docker compose -f platform/docker-compose.yml exec -T postgresql psql -U postgres -d postgres -At -c "
  SELECT o.id || '|' || o.name FROM public.organizations o
  JOIN public.memberships m ON m.organization=o.id
  JOIN public.users u ON u.id=m.\"user\"
  WHERE u.email='$YOUR_EMAIL' AND m.role='admin' ORDER BY o.name;
"
```

Then for each line, run (scripted in a `for` loop is fine):

```bash
DB_ORG_UUID=…
NAME='Acme Corp'

# Create the Clerk dev org (omit slug — dev apps default-disable slugs)
NEW_CLERK_ORG=$(curl -s -X POST https://api.clerk.com/v1/organizations \
  -H "Authorization: Bearer $CLERK_SECRET" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"$NAME\",\"created_by\":\"$CLERK_USER_ID\"}" \
  | python3 -c "import json,sys;print(json.load(sys.stdin)['id'])")

# Set public_metadata.eid → DB org UUID (so JWTs can carry it once you customize the session token)
curl -s -X PATCH "https://api.clerk.com/v1/organizations/$NEW_CLERK_ORG" \
  -H "Authorization: Bearer $CLERK_SECRET" -H "Content-Type: application/json" \
  -d "{\"public_metadata\":{\"eid\":\"$DB_ORG_UUID\"}}"

# Rewrite the DB org's clerk_id
docker compose -f platform/docker-compose.yml exec -T postgresql psql -U postgres -d postgres -c "
  UPDATE public.organizations SET clerk_id='$NEW_CLERK_ORG', updated_at=now() WHERE id='$DB_ORG_UUID';
"
```

### 3c. Sync sentinel casbin grants

The dump's `sentinel.casbin_rules.g2` is sometimes inconsistent with `public.memberships` — you'll be `admin` in N orgs in the DB but only have casbin grants for fewer. Without the casbin grant, the Casbin PDP rejects every `org:feature:*` check (members:list, workforce, etc.) with `forbidden`, which the UI surfaces as empty Members tables, empty Activity Dashboard, etc.

> **⚠️ Since `766df6c` the PDP is in-process and there is no `sentinel` container to restart.** v11.0
> folded it into `app` as `app/internal/sentinel/` (wired once at `app/main.go:305`, *"There is no switch
> and no RPC path: app IS the PDP"*) and deleted the compose service. The **policy tables did not move** —
> they stay in the `sentinel` schema (`SENTINEL_DB_CONNECTION`, `search_path=sentinel`), so the SQL below
> is unchanged and still correct. Only the reload step changed.

Sync them in one shot:

```bash
docker compose -f platform/docker-compose.yml exec -T postgresql psql -U postgres -d postgres -c "
  INSERT INTO sentinel.casbin_rules (p_type, v0, v1, v2)
  SELECT 'g2', m.organization::text, m.\"user\"::text, m.role::text
    FROM public.memberships m
    JOIN public.users u ON u.id=m.\"user\"
    WHERE u.email='$YOUR_EMAIL'
  ON CONFLICT DO NOTHING;
"

# Reload the policies. The enforcer calls LoadPolicy() at startup, so restarting the process
# that HOSTS it is the reload — and since 766df6c that process is `backend`, not `sentinel`.
# `restart sentinel` now ERRORS ("no such service"); it does not silently no-op, so you will see it.
docker compose -f platform/docker-compose.yml restart backend
```

> Non-restart alternative on a live stack: the in-process PDP subscribes to the Redis **Pub/Sub** channel
> `sentinel:policy:invalidate` for cross-replica invalidation (`app/internal/sentinel/watcher.go:55` —
> deliberately Pub/Sub fan-out, *not* the Watermill consumer-group plumbing, which would deliver to one
> consumer only). Publishing to that channel reloads without a bounce. `restart backend` is the blunter
> path and always works.

### 3d. Customize the dev Clerk session token — MANDATORY, not optional

> ⚠️ **Corrected 2026-09-03. This step was labelled "(Optional, recommended)" and its claim set
> was incomplete — following it verbatim leaves every org-scoped page empty.** The old JSON
> supplied only `org.eid`. `colony v0.35.2` needs **three** org claims and returns nil unless all
> three are non-empty, which surfaces as `forbidden: organization mismatch` on
> `organizationMembers` and `org-context is missing: ent/privacy: deny rule` on everything else.

From `colony/authn/provider/clerk/clerk_user.go:148-175`:

```go
clerkId,   _ = u.tokenClaims.Extra["org_id"].(string)
clerkRole, _ = u.tokenClaims.Extra["org_role"].(string)
if orgPM, ok := u.tokenClaims.Extra["org"].(map[string]any); ok {
    orgEid, _ = orgPM["eid"].(string)
}
if clerkId == "" || clerkRole == "" || orgEid == "" {
    return nil          // -> ErrOrgMismatch on every org-scoped query
}
```

In the Clerk dashboard → Sessions → "Customize session token", use **exactly** this:

```json
{
  "eid": "{{user.external_id}}",
  "email": "{{user.primary_email_address}}",
  "firstname": "{{user.first_name}}",
  "lastname": "{{user.last_name}}",
  "org_id": "{{org.id}}",
  "org_role": "{{org.role}}",
  "org": "{{org.public_metadata}}"
}
```

Note `"org"` takes the **whole `public_metadata` object**, not `.eid` — colony indexes into it
itself. This matches the `long-session` JWT template already configured on the shared dev
instances, which is why environments cloned from a working one behave and hand-built ones did not.

Two things that make this hard to diagnose:

- **The JWT template and the session token are different Clerk features** with different claim
  shapes. Copying a working instance's *template* does not fix its *session token*, and the
  dashboard shows them in separate places.
- **The shape is baked in at token issue time.** After changing it you must sign out and back in;
  an existing session keeps the old claims and keeps failing.

To confirm rather than guess, mint a token for a live session and decode it:

```bash
SID=$(curl -s "https://api.clerk.com/v1/sessions?user_id=$CLERK_USER_ID&status=active" \
       -H "Authorization: Bearer $CLERK_SECRET" | python3 -c "import json,sys;print(json.load(sys.stdin)[0]['id'])")
curl -s -X POST "https://api.clerk.com/v1/sessions/$SID/tokens" \
  -H "Authorization: Bearer $CLERK_SECRET" \
  | python3 -c "import base64,json,sys;p=json.load(sys.stdin)['jwt'].split('.')[1];print(json.dumps(json.loads(base64.urlsafe_b64decode(p+'='*(-len(p)%4))),indent=1))"
```

`org_id`, `org_role` and `org.eid` must all be present and non-empty.

This is dashboard-only as of 2026-05; there's no public REST endpoint to script it.

---

## 4. Apply the colony patches

> **⚠️ SUPERSEDED 2026-08-07 — do not vendor colony. Bump the pin.** Bug 1 below is **fixed upstream at
> colony `v0.34.4`** (*"fix(clerk): ensure client is not nil when fetching user"*, `b810b28`,
> 2026-06-15 — read from `.colony-fork/CHANGELOG.md` as checked into `app` at `fc5607a`), and `app`
> took the bump the same day at `b30f25e`. **`app` at `ad9f3c498` pins `colony v0.35.2` (`go.mod:15`)
> with no colony `replace` directive** (its only `replace`, `:295`, is for `sentry-go/echo`), **no
> `vendor-colony/` and no `.colony-fork/`.** The platform itself vendored colony exactly once, for one
> day — `fc5607a` (2026-06-16) added `.colony-fork/`, `984c50b6` (2026-06-17) deleted all 55 files of
> it, dropped the `replace` from `go.mod` and the `COPY .colony-fork/` from `Dockerfile.dev`, and
> removed app's own hand-rolled v2-claim reader from `internal/web/backend/api/api.go` in the same
> commit. Whether colony `v0.35.2` reads the v2 claim shape is **not measurable from this corpus** —
> `colony` is not in the clone set and only `v0.34.3` is in the local module cache — so do not assert
> it either way; what *is* certain is that `app` ships stock upstream colony with no patch of any kind.
> **Full evidence in [`staging-bringup.md` Quirk #11](./staging-bringup.md#bringup-quirks-consolidated-as-a-procedural-narrative).**
>
> The rest of § 4 is kept as the diagnosis (the symptoms below are exactly what an old build shows) and
> as an unwind guide for a host that still carries a `vendor-colony/`.

The two issues, as they present against dev Clerk apps:

1. **Nil-deref panic on every authenticated GraphQL query** — `colony@v0.33.2`/`v0.34.0` (and still `v0.34.3`, the newest we can read locally: `clerk.go:62-66`) `authn/provider/clerk.GetUser` returns a `*User` with the `client` field unset. Any later call to `User.Email()` or `User.fetchUser()` panics on a nil receiver. `Email()` also nil-derefs when `PrimaryEmailAddressID` is missing on the fetched Clerk user. **Fixed at `v0.34.4`.**
2. **Clerk JWT v2 format unsupported** — modern Clerk dev apps issue tokens with claims under `o.{id,rol,slg}` instead of `org_id`/`org_role`/`org`. Colony's `GetOrganization()` only reads v1 names (`colony@v0.34.3/authn/provider/clerk/clerk_user.go:148-167`), so it returns `nil`, and downstream resolvers see "no active organization" → `forbidden: organization mismatch` on every org-scoped query.

The historical fix was to vendor a patched colony into each Go service and add a `replace` directive in its `go.mod`. Patched source was on Ithaca at `/home/devops/colony/`; the diff against upstream `v0.34.0` was small enough to copy by hand. Key changes:

`authn/provider/clerk/clerk.go` `GetUser`:
```go
return &User{
    sessClaims:  sessClaims,
    tokenClaims: tokenClaims,
    provider:    c,
    client:      c.client,   // ← was missing upstream
}, nil
```

`authn/provider/clerk/clerk_user.go` `Email()`:
```go
user := u.fetchUser()
if user.PrimaryEmailAddressID == nil {
    if len(user.EmailAddresses) > 0 {
        return user.EmailAddresses[0].EmailAddress
    }
    return ""
}
for _, e := range user.EmailAddresses {
    if e.ID == *user.PrimaryEmailAddressID {
        return e.EmailAddress
    }
}
return ""
```

`authn/provider/clerk/clerk_user.go` `GetOrganization()` — fall back to v2 claim names + lazy-fetch `public_metadata.eid` via Clerk API with a process-wide cache (Clerk rate-limits otherwise).

Each consuming service needed the steps below. On a current stack the Go service is **`app`, and only `app`** (this said *"the Go services are **`app` and `sentinel`**"* until M258 iter-18 — `766df6c` folded `sentinel` in as the 8th merge) — when this recipe was written it also covered `cms`, `jobsimulation`, `messenger` and `storage`, all since folded into `app`. **Do not run these; see the box at the top of § 4.** Reproduced so you can recognise and unwind a `vendor-colony/` on an old host:

1. `cp -r <patched-colony> <service>/vendor-colony`
2. Append to `<service>/go.mod`:
   ```
   replace github.com/anthropos-work/colony => ./vendor-colony
   ```
3. In `<service>/Dockerfile.dev`, add `COPY vendor-colony ./vendor-colony` immediately after `COPY go.sum ./` (before `RUN go mod download`).
4. `cd <service> && go mod tidy && cd ..`
5. `docker compose -f platform/docker-compose.yml build <service>`

**To unwind one:** unmark the `skip-worktree` on `go.mod`, `go.sum` and `Dockerfile.dev`, `git checkout --` all three, `rm -rf vendor-colony/`, drop the `vendor-colony/` line from `.git/info/exclude`, and rebuild. Upstream is where the fix lives now — `app` pins a plain colony version and nothing else.

---

## 5. Bring up the rest of the stack

```bash
cd platform
docker compose --profile all up --build -d
```

Wait for all services to report healthy:

```bash
docker compose ps --format "table {{.Service}}\t{{.Status}}"
```

You should see **6** services running at platform `766df6c` — `--profile all` selects the whole
effective topology (`backend`, `gotenberg`, `next-web-app`, `studio-desk` + the always-on
`postgresql`/`redis` floor). ⚠️ **This read *"**7** services … at platform `0c91421` … the always-on
`postgresql`/`redis`/`sentinel` floor"* until M258 iter-18** — true at that ref, and RETRACTED at
`766df6c` (v11.0), which deleted the `sentinel` service, taking the floor from three to two and the
`--profile all` count from 7 to 6. The count was ~14 before the cms/jobsimulation/roadrunner
fold and the `graphql` deletion, 8 at `0dab54d`, and 7 since `838d907` dropped the `storage`,
`messenger` and `customerio-sync` containers (corrected M257x iter-78, re-measured iter-87). If any
service crashes on boot, check its logs (`docker compose logs <svc> --tail 30`) — most failures are missing env vars in `.env` or a Dockerfile gap; see Troubleshooting below.

---

## 6. Cross-device access via Tailscale (optional)

If you want to open the staging from another device on your Tailscale network (e.g., your laptop while the stack runs on a remote VM):

1. The host (where the stack runs) must be on Tailscale. Look up its Tailscale IP: `ip -4 addr show tailscale0 | grep inet`.
2. The frontend baked the wrong host at build time unless you set `PUBLIC_HOST` before `docker compose build`. Edit `platform/.env`:
   ```
   PUBLIC_HOST=100.x.y.z
   ```
   …then `docker compose build next-web-app && docker compose up -d next-web-app`.
3. Add the Tailscale origin to `CORS_EXTRA_ORIGINS` in `platform/.env` (comma-separated) and `docker compose restart backend` — **no rebuild, and do not edit `cors.go`.** The env var landed at `app` `f664473` (2026-05-14) and was hard-gated out of production at `13410de` (2026-05-19); both are ancestors of `app` `ad9f3c498`. It is read at `app/internal/cors/cors.go:24` and applied at `:78-82` under `if !environment.IsProduction()`.
4. Add the same origin in your dev Clerk app's "Allowed origins" list.
5. From the other device, open `http://100.x.y.z:3000/login`.

### Tailscale aliases (so URLs aren't IP-numbers)

Tailscale's MagicDNS gives each device one canonical short name (e.g., `calypso.taildc510.ts.net`). To have *additional* friendly aliases (e.g., let the same machine answer to both `calypso` and `calypsostaging`), add them to the **`hosts:`** section of the tailnet ACL:

```hcl
{
  "hosts": {
    "calypso":        "100.83.121.80",
    "calypsostaging": "100.83.121.80",
    "ithaca":         "100.120.254.65",
    "ithacastaging":  "100.120.254.65",
  },
  // ... rest of ACL ...
}
```

Save in https://login.tailscale.com/admin/acls/file (or `POST /api/v2/tailnet/<tailnet>/acl`) and resolution is instant tailnet-wide. After this, both `http://calypso:3000` and `http://calypsostaging:3000` work from any tailnet device. Remember to add each new alias to:

- **Clerk allowed origins** (https://api.clerk.com/v1/instance with `allowed_origins`).
- **Backend CORS** — append the alias to `CORS_EXTRA_ORIGINS` in `platform/.env`, then restart `backend`. (The old instruction here was to edit `app/internal/cors/cors.go`'s `colony.Development` block; that is obsolete — see step 3 above.)

Both lists must contain `http://<alias>:3000` for the browser to trust it.

---

## 7. Verify

After everything is up and you've completed the rebind, log in via:

```
http://localhost:3000/login   (or http://<tailscale-ip>:3000/login)
```

with your engineer email + the password you set in 3a. If your dev Clerk app still has the "new device" sign-in challenge enabled and you don't want to receive the email code, bypass it with a one-shot ticket:

```bash
TOKEN=$(curl -s -X POST https://api.clerk.com/v1/sign_in_tokens \
  -H "Authorization: Bearer $CLERK_SECRET" -H "Content-Type: application/json" \
  -d "{\"user_id\":\"$CLERK_USER_ID\",\"expires_in_seconds\":600}" \
  | python3 -c "import json,sys;print(json.load(sys.stdin)['token'])")
echo "http://localhost:3000/login?__clerk_ticket=$TOKEN"
```

Smoke test with Playwright (or by hand):
- `/home` shows your name, AI Sims count, XP — real data from the dump.
- `/library/ai-simulations` lists public simulations from Directus.
- `/library/skill-paths` lists public skill paths.
- `/enterprise/members` lists members of your active org.
- `/enterprise/activity-dashboard` renders without `forbidden`.

If `/enterprise/members` shows `0 / 50, No data`, re-check step 3c (casbin sync) — it almost always means a g2 row is missing for your user × your active org and sentinel is rejecting `org:feature:members:list`.

---

## Troubleshooting

### "forbidden: organization mismatch" on enterprise routes
Backend's `colony.GetOrganization()` returned `nil` because the JWT's `o.id` (or `org_id`) didn't resolve to a valid `eid`. Either the org's `public_metadata.eid` (step 3b) isn't set, or `app` is on a colony version that can't read your token's claim shape (check `grep colony ~/app/go.mod` against § 4 — the remedy is a version bump, **not** a vendor tree). **Note the blast radius:** this is not REST-only. A nil org context also denies GraphQL, because `OrganizationMixin`'s Ent privacy policy opens with `DenyIfNoOrganizationInContext()` on **30** schemas — see [`staging-bringup.md` Quirk #11](./staging-bringup.md#bringup-quirks-consolidated-as-a-procedural-narrative). Verify:
```bash
docker compose logs backend --since 1m | grep -i "organization mismatch\|colony: failed"
```

### "forbidden" on members/workforce queries
Casbin doesn't know you're an admin of the active org. Re-run step 3c and restart **`backend`** — since
`766df6c` the PDP is in-process, there is no `sentinel` container, and `restart sentinel` errors out.

### Members table shows skeleton rows forever
The query is panicking on the backend with `nil pointer dereference` — the colony nil-client bug, same root cause as "organization mismatch", different code path. **Fixed upstream at colony `v0.34.4`**; if you are seeing it, `app` is pinned below that. Bump `~/app/go.mod` and rebuild backend — do not vendor (§ 4).

### 422 on `publicJobSimulations`
Same colony root cause. **There is no `cms` service to patch** — the cms domain runs in-process inside `backend` (`app/internal/cms/`), and `cms` is neither a compose service nor a `repos.yml` entry at platform `0c91421d`. Nor is there a subgraph: platform `2adcf71` (2026-07-31) deleted the federation router, and `backend` serves one GraphQL schema at `:8082/graphql/query` (`app/internal/web/backend/backend.go:317`). Fix it on `backend` and only `backend`.

### "FATAL: role X does not exist" during DB restore
Harmless — these are GRANT statements on missing roles. Data tables load fine.

### Bitnami postgres won't start: "Permission denied"
The bind-mount root needs to be owned by uid 1001:
```bash
sudo chown -R 1001:1001 platform/data/postgresql
docker compose restart postgresql
```

### studio-desk fails to bind port 9100
> ⚠️ **NOTHING LISTENS ON 9100 SINCE THE NEXT MIGRATION** — it was the Vite dev-server port, present only
> under the old `npm run dev` and never inside the container. Compose still publishes it, so the bind
> conflict with `node_exporter` is still reachable, but the correct fix is to stop publishing a dead port
> rather than to remap it. The live port is `9000 + N*OFFSET`. Historical remap:

Conflicts with `node_exporter` (Prometheus monitoring) on the host. Edit `platform/docker-compose.yml`:
```yaml
studio-desk:
  ports:
    - "9101:9100"   # was 9100:9100
```

### Image build fails on `COPY … studio` — ACQUIRE the tree, do not delete the lines
On a current stack the failing image is **`backend` (`app`)**, not `cms`: `app/Dockerfile:45-46` hard-COPYs
`/build/studio`, nothing in the documented flow puts that tree on disk, and `make up` cannot complete
without it. Clone it — [`setup_guide.md` § Acquire the Studio
runtime](setup_guide.md#acquire-the-studio-runtime--required-before-make-up-or-the-backend-build-fails).

> **⚠️ Corrected M257x iter-265.** This entry read *"the `studio/` submodule was removed from `cms/main`.
> Edit `cms/Dockerfile.dev` and remove … The Go binary runs without the Python studio runner."* That was
> true of the frozen `cms` repo and is **false for every image a current staging builds** — `cms` is
> decommissioned, and `app` needs the Python runtime because it now hosts the embedded studio-room
> pipeline. **The requirement migrated with the fold; this troubleshooting entry did not**
> (`D-M257x-265-1`; two sibling copies in `setup_guide.md` and `staging-bringup.md` carried the same text).

### `backend` exits 1 immediately: "webhook secret may not be empty"

```
can't init clerk events manager: can't init svix webhook handler: webhook secret may not be empty
```

`CLERK_WEBHOOK_SECRET` is blank. The svix handler is constructed at boot and refuses an empty
secret, so the container dies before serving anything. Clerk cannot reach a laptop, so no real
webhook will ever arrive — any well-formed value works, and an obviously fake one is preferable
so nobody mistakes it for real:

```bash
CLERK_WEBHOOK_SECRET=whsec_RkFLRS1MT0NBTC1ERVYtTk9ULUEtUkVBTC1TRUNSRVQ=
CLERK_WEBHOOK_REDIRECT_URL=https://fake.invalid/clerk-webhook-not-configured-local-dev
```

### `directus` is `unhealthy`, 503 on `/server/health`, sims library empty

```
CredentialsProviderError: Could not load credentials from any providers
```

`DIRECTUS_STORAGE=s3` on a machine with no AWS credentials. `docker-compose.yml` reads it as
`STORAGE_LOCATIONS=${DIRECTUS_STORAGE:-local}` and defaults to `local` for exactly this reason —
someone copying a working `.env` from a machine that *does* have AWS credentials carries the `s3`
value across. Set `DIRECTUS_STORAGE=local` and recreate the container.

With `local`, the restored `directus_files` rows still point at S3 keys, so locally-served asset
bytes 404. The backend reads production `content.anthropos.work` for the public asset plane, so
library imagery still renders; only locally-authored assets are missing.

### 403 on `/_next/static/chunks/*` when reached by hostname

```
⚠ Blocked cross-origin request to Next.js dev resource /_next/… from "<host>".
```

Next.js 16 blocks cross-origin dev resources by default. Reaching the dev server as anything
other than `localhost` needs the host listed in `apps/web/next.config.mjs`:

```js
allowedDevOrigins: ['<short-host>', '<host>.<tailnet>.ts.net', '<tailscale-ip>'],
```

### Login loops forever on `?__clerk_handshake=…`

Clerk sets `__session` / `__client_uat` with `Secure; SameSite=None`. Browsers grant secure-context
status only to `localhost`/`127.0.0.1`, so over **plain HTTP to any other hostname the cookies are
silently dropped** and the handshake repeats indefinitely. This is browser policy — no app or
Clerk setting changes it. Pick one:

- browse `http://localhost:3000` on the box itself;
- `ssh -L 3000:127.0.0.1:3000 -L 8082:127.0.0.1:8082 <host>` and use `localhost` remotely (the
  ports must match, because the bundle hardcodes them);
- `tailscale serve --bg --https=443 http://127.0.0.1:3000` for a real certificate;
- or, for a throwaway profile, launch Chrome with
  `--unsafely-treat-insecure-origin-as-secure=http://<host>:3000 --user-data-dir=/tmp/<profile>`.

### Go build killed: "cannot allocate memory" / `signal: killed`

`docker compose --profile all up --build` builds every service concurrently. A **cold** build
compiles `internal/data/ent` (large, generated) while next-web-app runs its Turbo build, and the
pair exceeds a default Docker Desktop allocation. Machines that already have images built never
hit this, because their rebuilds are incremental — so "it works on mine" proves nothing here.
Build one service at a time:

```bash
for svc in backend next-web-app studio-desk directus-setup; do
  docker compose --profile all build "$svc" || break
done
```

### A remote setup script stops silently partway through

`docker compose exec -T` reads stdin. A script piped in as `ssh host 'zsh -s' < script.sh` has
**the rest of itself consumed** by the first such command, and the shell exits 0 as though it
finished. Either redirect every one (`docker compose exec -T … < /dev/null`) or — simpler — `scp`
the script and run it as a file.

### Clerk works from one machine but not another (`api.clerk.com` hijack)

A machine that has run the **rosetta demo stack** may still carry its Clerk interception:
`/etc/hosts` maps `127.0.0.1 api.clerk.com`, with a root `socat` on `:443` presenting an
mkcert-forged `api.clerk.com` certificate and forwarding to the demo's `fake-bapi`. Every Clerk
call from that machine's browser then hits the fake API. It is invisible from any other machine,
which makes it look like an app bug. Check and clear:

```bash
grep -c clerk /etc/hosts                       # must be 0
ping -c1 api.clerk.com                         # must NOT be 127.0.0.1
echo | openssl s_client -connect api.clerk.com:443 -servername api.clerk.com 2>/dev/null \
  | openssl x509 -noout -issuer                # must be a real CA, not "mkcert development CA"

sudo cp /etc/hosts /etc/hosts.bak && sudo sed -i '' '/api\.clerk\.com/d' /etc/hosts
sudo dscacheutil -flushcache; sudo killall -HUP mDNSResponder
```

The demo also parks user-owned `socat` forwarders on 3000/3001/9000 (loopback-only), which
collide with `pnpm dev`. Killing them does not touch its containers or volumes.

### `docker compose --profile all` fails: `path "…/directus" not found`

`--profile all` includes `directus-setup`, whose build context is `../directus`. `make init` does
not clone that repo — clone it as a sibling, or bring the stack up without it
(`--profile core --profile frontend`).

### Next.js build crashes with "STRIPE_SECRET_KEY is not configured"
Next.js statically evaluates server routes at build time and reads from `process.env`. Compose `env_file` is runtime-only. Drop a gitignored `next-web-app/apps/web/.env.production` containing the keys the routes need (Stripe, OpenAI, Azure OpenAI, Clerk publishable, Wundergraph endpoint, etc.) before `docker compose build`.

---

## Reset and start over

If the rebind goes wrong, you can wipe Postgres and re-run from step 2:

```bash
cd platform
docker compose down
sudo rm -rf data/postgresql
sudo mkdir -p data/postgresql && sudo chown -R 1001:1001 data/postgresql
docker compose up -d postgresql
# then re-run from step 2
```

Clerk-side cleanup (delete the dev orgs and your dev user) can be done by hand in the Clerk dashboard or via `DELETE /v1/organizations/<id>` and `DELETE /v1/users/<id>`.

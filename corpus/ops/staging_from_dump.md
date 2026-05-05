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
- `platform/`, `app/`, `cms/`, `skiller/`, `jobsimulation/`, `skillpath/`, `sentinel/`, `storage/`, `messenger/`, `roadrunner/`, `next-web-app/`, `studio-desk/`, `graphql-wundergraph/` cloned as siblings.
- `platform/.env` with `GH_PAT`, `CLERK_SECRET_KEY`, `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` filled in.
- `make up postgresql` succeeds and Postgres is healthy.

You also need:
- A recent prod `pg_dump` SQL file (plain SQL, not custom format), accessible at a known path (e.g. `~/prod_dump.sql`).
- Access to a **dev Clerk app** (NOT production Clerk). If you don't have one, create a fresh dev application in the Clerk dashboard and copy its publishable + secret keys.
- Your prod email address — the same email that exists in `public.users` of the dump.

> **DO NOT use prod Clerk keys on a dev/staging machine.** Any user-mutating action (create / ban / update metadata) you trigger from staging will execute against the live Clerk org. Always use a dev Clerk app.

---

## 1. Outbound-email kill switch (mandatory, do this FIRST)

A staging stack restored from a prod dump contains real customer email addresses in `public.users`. Many code paths trigger transactional notifications via `messenger` → Brevo (welcome emails, invitation flows, weekly recaps, password resets). If `BREVO_KEY` is set to a real value, those emails will go out to real people the moment you exercise the relevant flow.

**Blank `BREVO_KEY` in `platform/.env` and restart `messenger` BEFORE running any flow that could enqueue a notification:**

```bash
sed -i.bak 's/^BREVO_KEY=.*/BREVO_KEY=/' platform/.env
docker compose -f platform/docker-compose.yml restart messenger
```

Verify:

```bash
docker compose -f platform/docker-compose.yml exec -T messenger env | grep BREVO_KEY
# Expected: BREVO_KEY=  (empty)
```

The messenger boots with `INFO Brevo Messenger` either way; with the key blank, every API call to Brevo fails at the 401 layer and no email is delivered.

Apply the same caution to any other live-customer integration you don't intend to fire from staging:
- `CUSTOMERIO_*` (marketing/lifecycle email — disable unless you're testing customer.io integration explicitly).
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

Bring up Postgres only, then pipe the dump into `psql`:

```bash
cd platform
docker compose up -d postgresql
# wait for it to be healthy
until docker compose exec -T postgresql pg_isready -U postgres > /dev/null 2>&1; do sleep 1; done

# Restore (~5-10 min for a 500MB dump)
cat ~/prod_dump.sql | docker compose exec -T postgresql psql -U postgres -d postgres -v ON_ERROR_STOP=0 > /tmp/restore.log 2>&1

# Sanity check
docker compose exec -T postgresql psql -U postgres -d postgres -c "
  SELECT 'users' tbl, COUNT(*) FROM public.users
  UNION ALL SELECT 'organizations', COUNT(*) FROM public.organizations
  UNION ALL SELECT 'memberships', COUNT(*) FROM public.memberships
  UNION ALL SELECT 'casbin_rules', COUNT(*) FROM sentinel.casbin_rules;
"
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

The dump's `sentinel.casbin_rules.g2` is sometimes inconsistent with `public.memberships` — you'll be `admin` in N orgs in the DB but only have casbin grants for fewer. Without the casbin grant, sentinel rejects every `org:feature:*` check (members:list, workforce, etc.) with `forbidden`, which the UI surfaces as empty Members tables, empty Activity Dashboard, etc.

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

# Restart sentinel so it reloads policies
docker compose -f platform/docker-compose.yml restart sentinel
```

### 3d. (Optional, recommended) Customize the dev Clerk session token

Clerk's session token doesn't include your DB UUID by default, so the Anthropos backend has to fetch your user/org from the Clerk Backend API on every request. Clerk rate-limits aggressively (HTTP 429), and the first few page loads after a fresh login can fail with "organization mismatch" errors.

In the Clerk dashboard → Sessions → "Customize session token", add these custom claims:

```json
{
  "eid": "{{user.external_id}}",
  "email": "{{user.primary_email_address}}",
  "firstname": "{{user.first_name}}",
  "lastname": "{{user.last_name}}",
  "org": {
    "eid": "{{org.public_metadata.eid}}"
  }
}
```

This is dashboard-only as of 2026-05; there's no public REST endpoint to script it.

---

## 4. Apply the colony patches

The shared `colony` library used by every Go service has two unfixed-upstream issues that surface only against dev Clerk apps:

1. **Nil-deref panic on every authenticated GraphQL query** — `colony@v0.33.2`/`v0.34.0` `authn/provider/clerk.GetUser` returns a `*User` with the `client` field unset. Any later call to `User.Email()` or `User.fetchUser()` panics on a nil receiver. `Email()` also nil-derefs when `PrimaryEmailAddressID` is missing on the fetched Clerk user.
2. **Clerk JWT v2 format unsupported** — modern Clerk dev apps issue tokens with claims under `o.{id,rol,slg}` instead of `org_id`/`org_role`/`org`. Colony's `GetOrganization()` only reads v1 names, so it returns `nil`, and downstream resolvers see "no active organization" → `forbidden: organization mismatch` on every org-scoped query.

The fix is to vendor a patched colony into each Go service and add a `replace` directive in its `go.mod`. Patched source is on Ithaca at `/home/devops/colony/`; the diff against upstream `v0.34.0` is small enough to copy by hand. Key changes:

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

Each consuming service (`app`, `cms`, `skiller`, `jobsimulation`, `skillpath`, `messenger`, `storage`, `sentinel`) needs:

1. `cp -r <patched-colony> <service>/vendor-colony`
2. Append to `<service>/go.mod`:
   ```
   replace github.com/anthropos-work/colony => ./vendor-colony
   ```
3. In `<service>/Dockerfile.dev`, add `COPY vendor-colony ./vendor-colony` immediately after `COPY go.sum ./` (before `RUN go mod download`).
4. `cd <service> && go mod tidy && cd ..`
5. `docker compose -f platform/docker-compose.yml build <service>`

Once this lands upstream in `colony`, the vendor + replace can be removed and each service can pin a fixed colony version directly.

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

You should see ~15 services running. If any service crashes on boot, check its logs (`docker compose logs <svc> --tail 30`) — most failures are missing env vars in `.env` or a Dockerfile gap; see Troubleshooting below.

---

## 6. Cross-device access via Tailscale (optional)

If you want to open the staging from another device on your Tailscale network (e.g., your laptop while the stack runs on a remote VM):

1. The host (where the stack runs) must be on Tailscale. Look up its Tailscale IP: `ip -4 addr show tailscale0 | grep inet`.
2. The frontend baked the wrong host at build time unless you set `PUBLIC_HOST` before `docker compose build`. Edit `platform/.env`:
   ```
   PUBLIC_HOST=100.x.y.z
   ```
   …then `docker compose build next-web-app && docker compose up -d next-web-app`.
3. Add the Tailscale origin to the backend's CORS allowlist (`app/internal/cors/cors.go`, in the `colony.Development` block). Rebuild backend.
4. Add the same origin in your dev Clerk app's "Allowed origins" list.
5. From the other device, open `http://100.x.y.z:3000/login`.

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
Backend's `colony.GetOrganization()` returned `nil` because the JWT's `o.id` (or `org_id`) didn't resolve to a valid `eid`. Either the colony patch (step 4) isn't applied, or the org's `public_metadata.eid` (step 3b) isn't set. Verify:
```bash
docker compose logs backend --since 1m | grep -i "organization mismatch\|colony: failed"
```

### "forbidden" on members/workforce queries
Casbin doesn't know you're an admin of the active org. Re-run step 3c and restart sentinel.

### Members table shows skeleton rows forever
The query is panicking on the backend with `nil pointer dereference`. Apply the colony patch (step 4) and rebuild backend — this is the same root cause as "organization mismatch", just a different code path.

### CMS subgraph 422 on `publicJobSimulations`
Same colony root cause — the `cms` service needs the same vendor-colony patch as `app`.

### "FATAL: role X does not exist" during DB restore
Harmless — these are GRANT statements on missing roles. Data tables load fine.

### Bitnami postgres won't start: "Permission denied"
The bind-mount root needs to be owned by uid 1001:
```bash
sudo chown -R 1001:1001 platform/data/postgresql
docker compose restart postgresql
```

### studio-desk fails to bind port 9100
Conflicts with `node_exporter` (Prometheus monitoring) on the host. Edit `platform/docker-compose.yml`:
```yaml
studio-desk:
  ports:
    - "9101:9100"   # was 9100:9100
```

### CMS image build fails on `COPY studio/`
The `studio/` submodule was removed from `cms/main`. Edit `cms/Dockerfile.dev` and remove:
```dockerfile
COPY studio/ ./studio/
RUN pip install --no-cache-dir -r studio/requirements.txt
```
The Go binary runs without the Python studio runner.

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

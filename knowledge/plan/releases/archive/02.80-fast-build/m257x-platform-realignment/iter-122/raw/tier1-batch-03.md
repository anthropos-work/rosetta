# TIER-1 ADJUDICATION BATCH 03 — 44 pairs

Each item: the CLAIMING UNIT (corpus prose) and the CITED CONTENT it points at (source lines,
line-numbered, +-3 lines of context). Answer ONE question per item.

## 03-001
- **id**: `B03-001`
- **corpus site**: `corpus/services/clerk-integration.md:68-68` (table-row)
- **citation**: `src/lib/devLogin.ts:33`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/lib/devLogin.ts`  (39 lines)

**CLAIMING UNIT**

```md
| 4 | `studio-desk/src/routes/dev.ts:83` | `studio-desk` `41ee357` | Dev login harness → `dev-accept.html` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`src/lib/devLogin.ts:33`) |
```

**CITED CONTENT**

```
    30  // (`npm start` / any deploy) sets NODE_ENV=production, so the route is never
    31  // mounted and hard-404s. It also does nothing without CLERK_SECRET_KEY.
    32  
    33  export const DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production';
    34  
    35  // Optional convenience: when set, `GET /api/dev/login-as` with NO `email=` query
    36  // param signs you in as this address. Lets the agentic workflow use a single bare
```

## 03-002
- **id**: `B03-002`
- **corpus site**: `corpus/services/clerk-integration.md:69-69` (table-row)
- **citation**: `ant-academy/code/app/api/dev/login-as/route.js:78`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/app/api/dev/login-as/route.js`  (98 lines)

**CLAIMING UNIT**

```md
| 5 | `ant-academy/code/app/api/dev/login-as/route.js:78` | `ant-academy` `22df69dd` | Dev login harness → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`code/src/lib/devLogin.js:29`); hard-404 otherwise |
```

**CITED CONTENT**

```
    75      }
    76  
    77      // 5. Mint a one-time sign-in token (10-minute validity is plenty for a redirect).
    78      const { token } = await client.signInTokens.createSignInToken({
    79        userId: user.id,
    80        expiresInSeconds: 600,
    81      })
```

## 03-003
- **id**: `B03-003`
- **corpus site**: `corpus/services/clerk-integration.md:69-69` (table-row)
- **citation**: `code/src/lib/devLogin.js:29`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/ant-academy/code/src/lib/devLogin.js`  (35 lines)

**CLAIMING UNIT**

```md
| 5 | `ant-academy/code/app/api/dev/login-as/route.js:78` | `ant-academy` `22df69dd` | Dev login harness → `/dev/accept` | `DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'` (`code/src/lib/devLogin.js:29`); hard-404 otherwise |
```

**CITED CONTENT**

```
    26  // genuinely local-dev-only — the endpoint hard-404s anywhere it is deployed.
    27  // It also does nothing without CLERK_SECRET_KEY configured locally.
    28  
    29  export const DEV_LOGIN_ENABLED = process.env.NODE_ENV !== 'production'
    30  
    31  // Optional convenience: when set, `GET /api/dev/login-as` with NO `email=` query
    32  // param signs you in as this address. Lets the agentic workflow use a single
```

## 03-004
- **id**: `B03-004`
- **corpus site**: `corpus/services/clerk-integration.md:76-80` (paragraph)
- **citation**: `e2e/auth.setup.ts:55-62`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/next-web-app/e2e/auth.setup.ts`  (87 lines)

**CLAIMING UNIT**

```md
**Site 3 is the one worth reading twice.** `e2e/auth.setup.ts:55-62` states its own reason: the e2e
account *"enforces 2FA (email_code as second factor); password signin returns `needs_second_factor` and
never produces a session"*, and *"Clerk treats it as fully authenticated and **skips both factors**."* A
sign-in token is therefore a **documented, deliberate second-factor bypass** against a real Clerk instance
— which is precisely the kind of fact an emphatic **only** on one unrelated site hides.
```

**CITED CONTENT**

```
    52      );
    53    }
    54  
    55    // Mint a one-time sign-in ticket on the backend instead of going through the
    56    // password form. Two reasons:
    57    //   1. Stefano's account enforces 2FA (email_code as second factor); password
    58    //      signin returns `needs_second_factor` and never produces a session.
    59    //   2. Tickets don't burn dev-tier signin rate limits the way repeated
    60    //      password attempts do.
    61    // The ticket is consumed in-page via `clerk.signIn({ strategy: 'ticket' })`
    62    // — Clerk treats it as fully authenticated and skips both factors.
    63    const clerkClient = createClerkClient({ secretKey });
    64    const { data: users } = await clerkClient.users.getUserList({
    65      emailAddress: [email],
```

## 03-005
- **id**: `B03-005`
- **corpus site**: `corpus/services/clerk-integration.md:82-87` (paragraph)
- **citation**: `staging-clerk.md:58`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/ops/staging-clerk.md`  (466 lines)

**CLAIMING UNIT**

```md
**A sixth consumer is an operator recipe, not checked-in code**, and this corpus already documents it in
three other places: the staging/CI `curl -s -X POST https://api.clerk.com/v1/sign_in_tokens` bypass for
Clerk's "new device" challenge — [`staging-clerk.md:58`](../ops/staging-clerk.md),
[`staging_from_dump.md:384`](../ops/staging_from_dump.md), and
[`staging-bringup.md:461`](../ops/staging-bringup.md) (*"Quirk #13"*). It mints against whichever
instance the operator's key points at.
```

**CITED CONTENT**

```
    55  For programmatic CI flows where Clerk's "new device" challenge blocks you (Quirk #13), use a one-shot sign-in token instead:
    56  
    57  ```bash
    58  TOKEN=$(curl -s -X POST https://api.clerk.com/v1/sign_in_tokens \
    59    -H "Authorization: Bearer $CLERK_SECRET_KEY" -H "Content-Type: application/json" \
    60    -d "{\"user_id\":\"user_3DIYdXgwlr0Q0R12qDNbk4z95aZ\",\"expires_in_seconds\":600}" \
    61    | python3 -c "import json,sys;print(json.load(sys.stdin)['token'])")
```

## 03-006
- **id**: `B03-006`
- **corpus site**: `corpus/services/clerk-integration.md:82-87` (paragraph)
- **citation**: `staging_from_dump.md:384`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/ops/staging_from_dump.md`  (464 lines)

**CLAIMING UNIT**

```md
**A sixth consumer is an operator recipe, not checked-in code**, and this corpus already documents it in
three other places: the staging/CI `curl -s -X POST https://api.clerk.com/v1/sign_in_tokens` bypass for
Clerk's "new device" challenge — [`staging-clerk.md:58`](../ops/staging-clerk.md),
[`staging_from_dump.md:384`](../ops/staging_from_dump.md), and
[`staging-bringup.md:461`](../ops/staging-bringup.md) (*"Quirk #13"*). It mints against whichever
instance the operator's key points at.
```

**CITED CONTENT**

```
   381  with your engineer email + the password you set in 3a. If your dev Clerk app still has the "new device" sign-in challenge enabled and you don't want to receive the email code, bypass it with a one-shot ticket:
   382  
   383  ```bash
   384  TOKEN=$(curl -s -X POST https://api.clerk.com/v1/sign_in_tokens \
   385    -H "Authorization: Bearer $CLERK_SECRET" -H "Content-Type: application/json" \
   386    -d "{\"user_id\":\"$CLERK_USER_ID\",\"expires_in_seconds\":600}" \
   387    | python3 -c "import json,sys;print(json.load(sys.stdin)['token'])")
```

## 03-007
- **id**: `B03-007`
- **corpus site**: `corpus/services/clerk-integration.md:82-87` (paragraph)
- **citation**: `staging-bringup.md:461`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/ops/staging-bringup.md`  (638 lines)

**CLAIMING UNIT**

```md
**A sixth consumer is an operator recipe, not checked-in code**, and this corpus already documents it in
three other places: the staging/CI `curl -s -X POST https://api.clerk.com/v1/sign_in_tokens` bypass for
Clerk's "new device" challenge — [`staging-clerk.md:58`](../ops/staging-clerk.md),
[`staging_from_dump.md:384`](../ops/staging_from_dump.md), and
[`staging-bringup.md:461`](../ops/staging-bringup.md) (*"Quirk #13"*). It mints against whichever
instance the operator's key points at.
```

**CITED CONTENT**

```
   458  
   459  10. **Quirk #12 — Dev Clerk needs Organizations enabled + per-user/org `external_id` set.** Documented as the rebind procedure in §6 below.
   460  
   461  11. **Quirk #13 — Dev Clerk "new device" sign-in challenge** blocks programmatic login. Bypass with `POST /v1/sign_in_tokens` for Playwright / CI. Real-user login through the form is fine (Clerk emails the code on first sign-in, then trusts the device).
   462  
   463  12. **Quirk #15 — `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` must be in next-web-app's runtime `environment:` block, not just `env_file:`.** Clerk middleware reads it from `process.env` at runtime. If only `VITE_CLERK_PUBLISHABLE_KEY` is in the runtime env, Clerk's server-side init falls into the "infinite redirect loop" detector → blank pages. Fix: list `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` explicitly in compose's `next-web-app.environment:` array. Plus the four sign-in/up URL vars. Restart container — no rebuild needed (runtime-only). Sibling client-side symptom: stale `__clerk_db_jwt` cookies from a prior origin keep the loop alive after the env fix; clear cookies for the staging origin to recover.
   464  
```

## 03-008
- **id**: `B03-008`
- **corpus site**: `corpus/services/clerk-integration.md:92-97` (paragraph)
- **citation**: `ant-academy.md:324`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/ant-academy.md`  (504 lines)

**CLAIMING UNIT**

```md
**Corpus cross-references for the harness half** (this page previously mentioned it nowhere):
[`ant-academy.md:324`](./ant-academy.md) (the `DEV_LOGIN_ENABLED` public-route pair),
[`studio-desk.md:90`](./studio-desk.md) (*"Dev-only acceptance harness"*), and — added at **iter-121**, the
**pair-half iter-120 left open** — [`next-web-app.md` § the two minting sites in this repo](./next-web-app.md#the-two-clerk-sign-in-token-minting-sites-in-this-repo-added-m257x-iter-121). That page held **2 of the 5 sites, including the
only ungated one**, and said nothing about either; a repaired enumeration whose per-repo pages stay silent
is half a repair.
```

**CITED CONTENT**

```
   321  | Release provenance | **`/api/_meta(.*)`, `/api/meta(.*)`** | The academy's mirror of the Go services' `/_meta`, so uptime probes can read name/version/build-date without a session. Both spellings, because `next.config.js` rewrites `_meta` → `meta` and the middleware sees the pre-rewrite path |
   322  | Crawler / agent files | `/robots.txt`, `/sitemap.xml`, `/sitemap(.*)`, `/llms.txt`, `/llms-full.txt`, `/.well-known/(.*)` | Static, must be fetchable anonymously |
   323  | Assets & machine indexes | `/local-content/(.*)`, `/catalog.json`, `/academy-manifest.json` | Public-by-design: `/local-content/*` for `<audio>` Range requests + cover previews, `/catalog.json` for the external Anthropos backend Talk-to-Data indexer, `/academy-manifest.json` for the PWA manifest. Gating any of them 307s the fetch through sign-in and breaks it |
   324  | **Dev-only**, `DEV_LOGIN_ENABLED` | `/api/dev/login-as`, `/dev/accept` | The real-Clerk-user login shortcut; must be reachable before a session exists. Production drops both entries *and* the route handler hard-404s |
   325  | **Dev-only**, `BENCHMARK_VISUAL_BYPASS=1` ∧ `NODE_ENV==='development'` | `/my-certificates`, `/my-activity`, `/bookmarks` | Opens the remaining authed-only surfaces for the benchmark/e2e Playwright pass. `NODE_ENV` is whitelisted, not blacklisted, so unset/`test`/typos stay closed |
   326  
   327  Everything not matched: missing session → `/sign-in`; signed-in with zero org memberships → `/no-organization`.
```

## 03-009
- **id**: `B03-009`
- **corpus site**: `corpus/services/clerk-integration.md:92-97` (paragraph)
- **citation**: `studio-desk.md:90`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/corpus/services/studio-desk.md`  (450 lines)

**CLAIMING UNIT**

```md
**Corpus cross-references for the harness half** (this page previously mentioned it nowhere):
[`ant-academy.md:324`](./ant-academy.md) (the `DEV_LOGIN_ENABLED` public-route pair),
[`studio-desk.md:90`](./studio-desk.md) (*"Dev-only acceptance harness"*), and — added at **iter-121**, the
**pair-half iter-120 left open** — [`next-web-app.md` § the two minting sites in this repo](./next-web-app.md#the-two-clerk-sign-in-token-minting-sites-in-this-repo-added-m257x-iter-121). That page held **2 of the 5 sites, including the
only ungated one**, and said nothing about either; a repaired enumeration whose per-repo pages stay silent
is half a repair.
```

**CITED CONTENT**

```
    87  │   ├── academy/        # Academy UI
    88  │   ├── home/           # Home page
    89  │   ├── skills/         # Skills management UI
    90  │   ├── dev-accept/     # Dev-only acceptance harness (dev-accept.html; not in the prod build)
    91  │   ├── shared/         # Shared frontend utilities
    92  │   ├── services/       # Frontend services
    93  │   │   ├── graphql/    # GraphQL queries/mutations
```

## 03-010
- **id**: `B03-010`
- **corpus site**: `corpus/services/clerk-integration.md:159-159` (paragraph)
- **citation**: `app/go.mod:31`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/go.mod`  (296 lines)

**CLAIMING UNIT**

```md
> **SDK versions:** `@clerk/nextjs` is **NOT aligned** across the four surfaces — this line asserted *"all four on `^6.39.2`"* until M257x iter-108. Measured at each clone's own HEAD: `next-web-app` `apps/{web,hiring,integration}` are all on **`^6.39.6`** (`8297c684`), while `ant-academy` (`code/package.json:52`) is on **`^6.39.2`** (`22df69dd`) — **three on one version, one a patch-range behind**. (`next-web-app/apps/maintenance` declares no `@clerk/nextjs`, so it is not a fourth surface.) **`@clerk/clerk-expo` is NOT aligned:** `next-web-app/apps/mobile/package.json:6` pins **`~2.6.18`** while `ant-academy/mobile/package.json:18` pins **`~2.19.36`** — thirteen minor versions apart, on the two mobile surfaces. The Go side has **drifted again**: `app/go.mod:31` @ `5ba17044` reads **`clerk-sdk-go/v2 v2.7.0`**, not the `v2.6.0` this doc previously asserted for both. Since `v2.6.0` is *the version the Clerkenstein Alignment DNA targets*, re-verify `colony`'s pin and the DNA before trusting an alignment score (`CHECK-M257x-iter22-clerk-sdk-drift`).
```

**CITED CONTENT**

```
    28  	github.com/aws/smithy-go v1.27.4
    29  	github.com/clerk/clerk-sdk-go/v2 v2.7.0
    30  	github.com/dustin/go-humanize v1.0.1
    31  	github.com/gabriel-vasile/mimetype v1.4.15
    32  	github.com/getbrevo/brevo-go v1.1.3
    33  	github.com/getkin/kin-openapi v0.145.0
    34  	github.com/getsentry/sentry-go v0.48.0
```

## 03-011
- **id**: `B03-011`
- **corpus site**: `corpus/services/clerkenstein.md:37-69` (paragraph)
- **citation**: `alignment/cmd/alignctl/run.go:134-135`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/alignment/cmd/alignctl/run.go`  (155 lines)

**CLAIMING UNIT**

```md
> ### ⚠ The score, honestly (corrected in v2.3 M218 — it used to read "100% / 100% on all five surfaces")
>
> | surface | score | note |
> |---|---|---|
> | **Go SDK** (`clerk-2.6.0`, M1) | **100% overall · 100% critical** — **27/27 genes**, 14 capabilities | Gate is ≥95 / =100 ⇒ **MET**. (Was 97.2% / 26-of-27 until M219 landed the org-eid fix — see below.) |
> | **JS/FAPI** (`clerk-js-5`, M2) | 100% / 100% (9 genes) | |
> | **multi-identity seat-switch** (`clerk-multi-1`, M37) | 100% / 100% (9 genes) | |
> | **deployment/injection** (`clerk-deploy-1`, M3) | 100% / 100% (7 genes) | |
> | **`@clerk/express`** (`clerk-express-1`, M2c) | **UNMEASURABLE on a box without `@clerk/express` `node_modules`** — the runner cannot build, exits **rc=3 (`ExitUnmeasurable`), with NO score**. | **Not** a pass. ⚠️ **rc=2 is
> now `ExitRegressed` — a MEASURED regression** (`alignment/cmd/alignctl/run.go:134-135`); do not read a 2
> as a missing Node module. |
>
> **So "all five surfaces at 100%" is still false — but on ONE count now, not two.** It was two at M218
> (a RED gene *and* an unmeasurable surface); M219 closed the RED gene. What remains: four surfaces are
> measured and at 100%; the fifth is dependency-gated and frequently produces *no number at all* — which
> nothing treated as a failure.
>
> **The formerly-RED gene (M218 D16) — ✅ RESOLVED at M219.** `MembershipOrgIdentity/real-org-eid` shipped
> **failing on purpose** for one milestone. The fake BAPI *fabricated* `organization.public_metadata.eid` as
> `"org_eid_" + orgID` instead of the roster's real org UUID. It could not be fixed inside M218 (the milestone's exit gate was a p95 over 5
> cold reset-to-seed cycles graded on a specific binary; a runtime change restarts that count), so rather
> than **omit the field and keep a clean 100%**, the divergence was named in the report on **every single
> run** until it landed. **It has landed:** `clerk-backend/store.go:138` (`SeedOrgIdentity`) and `:151`
> (`LookupOrgEid`) ship the real roster org UUID, and the DNA records it —
> `alignment/dna/clerk-2.6.0.json:131`: *"M219 landed the fix … taking the Go surface 97.2% -> 100%."*
> The Go surface is **27/27**. `FIX-M219-bapi-org-eid` is CLOSED.
>
> **Why this matters more than the number.** Before M218, Clerkenstein scored **100% critical / 100%
> overall / 0 divergences while its fake BAPI returned the wrong human for ev
```

**CITED CONTENT**

```
   131  // THE RULE THIS ENCODES: **absence of a score is not a passing score.** An unmeasurable surface must be
   132  // impossible to mistake for a measured one — so it gets its own code, and a banner that says so.
   133  const (
   134  	ExitRegressed    = 2
   135  	ExitUnmeasurable = 3
   136  )
   137  
   138  // unmeasurable fails LOUD. It never returns a score, and it never returns ExitRegressed.
```

## 03-012
- **id**: `B03-012`
- **corpus site**: `corpus/services/clerkenstein.md:37-69` (paragraph)
- **citation**: `clerk-backend/store.go:138`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/clerkenstein/clerk-backend/store.go`  (294 lines)

**CLAIMING UNIT**

```md
> ### ⚠ The score, honestly (corrected in v2.3 M218 — it used to read "100% / 100% on all five surfaces")
>
> | surface | score | note |
> |---|---|---|
> | **Go SDK** (`clerk-2.6.0`, M1) | **100% overall · 100% critical** — **27/27 genes**, 14 capabilities | Gate is ≥95 / =100 ⇒ **MET**. (Was 97.2% / 26-of-27 until M219 landed the org-eid fix — see below.) |
> | **JS/FAPI** (`clerk-js-5`, M2) | 100% / 100% (9 genes) | |
> | **multi-identity seat-switch** (`clerk-multi-1`, M37) | 100% / 100% (9 genes) | |
> | **deployment/injection** (`clerk-deploy-1`, M3) | 100% / 100% (7 genes) | |
> | **`@clerk/express`** (`clerk-express-1`, M2c) | **UNMEASURABLE on a box without `@clerk/express` `node_modules`** — the runner cannot build, exits **rc=3 (`ExitUnmeasurable`), with NO score**. | **Not** a pass. ⚠️ **rc=2 is
> now `ExitRegressed` — a MEASURED regression** (`alignment/cmd/alignctl/run.go:134-135`); do not read a 2
> as a missing Node module. |
>
> **So "all five surfaces at 100%" is still false — but on ONE count now, not two.** It was two at M218
> (a RED gene *and* an unmeasurable surface); M219 closed the RED gene. What remains: four surfaces are
> measured and at 100%; the fifth is dependency-gated and frequently produces *no number at all* — which
> nothing treated as a failure.
>
> **The formerly-RED gene (M218 D16) — ✅ RESOLVED at M219.** `MembershipOrgIdentity/real-org-eid` shipped
> **failing on purpose** for one milestone. The fake BAPI *fabricated* `organization.public_metadata.eid` as
> `"org_eid_" + orgID` instead of the roster's real org UUID. It could not be fixed inside M218 (the milestone's exit gate was a p95 over 5
> cold reset-to-seed cycles graded on a specific binary; a runtime change restarts that count), so rather
> than **omit the field and keep a clean 100%**, the divergence was named in the report on **every single
> run** until it landed. **It has landed:** `clerk-backend/store.go:138` (`SeedOrgIdentity`) and `:151`
> (`LookupOrgEid`) ship the real roster org UUID, and the DNA records it —
> `alignment/dna/clerk-2.6.0.json:131`: *"M219 landed the fix … taking the Go surface 97.2% -> 100%."*
> The Go surface is **27/27**. `FIX-M219-bapi-org-eid` is CLOSED.
>
> **Why this matters more than the number.** Before M218, Clerkenstein scored **100% critical / 100%
> overall / 0 divergences while its fake BAPI returned the wrong human for ev
```

**CITED CONTENT**

```
   135  // The alignment seeds (NewSeeded / NewDemoSeeded) never call it, so with no roster the map is EMPTY and
   136  // organizationWithEid falls back to the demo-org eid / the historical stub. No DNA gene moves except the
   137  // one this fixes.
   138  func (s *Store) SeedOrgIdentity(org, eid string) {
   139  	if org == "" || eid == "" {
   140  		return // a half-formed roster entry must never shadow the working fallback
   141  	}
```

## 03-013
- **id**: `B03-013`
- **corpus site**: `corpus/services/clerkenstein.md:138-143` (paragraph)
- **citation**: `src/index.ts:96`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/src/index.ts`  (321 lines)

**CLAIMING UNIT**

```md
The data needed no new assembly: `/v1/me` already returns `userRes.OrganizationMemberships`. The role keeps
Clerk's **prefixed** form (`org:admin`) as a **fidelity** choice — it is what real Clerk emits on this route.
It is *not* a hard gate requirement: studio-desk's `STUDIO_ACCESS_ROLES` accepts **both** forms
(`['admin', 'org:admin', 'content_creator', 'org:content_creator']` — `src/index.ts:96` and
`app/services/userService.ts:16`, each carrying the comment "Both the prefixed (`org:*`) and bare role keys
are accepted"), so an unprefixed `admin` would pass too.
```

**CITED CONTENT**

```
    93  
    94  // Clerk org roles permitted to access Studio: admins and content creators.
    95  // Both the prefixed (org:*) and bare role keys are accepted.
    96  const STUDIO_ACCESS_ROLES = ['admin', 'org:admin', 'content_creator', 'org:content_creator'];
    97  
    98  // Middleware to check enterprise and Studio-eligible role status
    99  const checkEnterpriseAndAdmin = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
```

## 03-014
- **id**: `B03-014`
- **corpus site**: `corpus/services/clerkenstein.md:138-143` (paragraph)
- **citation**: `app/services/userService.ts:16`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/studio-desk/app/services/userService.ts`  (279 lines)

**CLAIMING UNIT**

```md
The data needed no new assembly: `/v1/me` already returns `userRes.OrganizationMemberships`. The role keeps
Clerk's **prefixed** form (`org:admin`) as a **fidelity** choice — it is what real Clerk emits on this route.
It is *not* a hard gate requirement: studio-desk's `STUDIO_ACCESS_ROLES` accepts **both** forms
(`['admin', 'org:admin', 'content_creator', 'org:content_creator']` — `src/index.ts:96` and
`app/services/userService.ts:16`, each carrying the comment "Both the prefixed (`org:*`) and bare role keys
are accepted"), so an unprefixed `admin` would pass too.
```

**CITED CONTENT**

```
    13  
    14  // Clerk org roles permitted to access Studio: admins and content creators.
    15  // Both the prefixed (org:*) and bare role keys are accepted.
    16  const STUDIO_ACCESS_ROLES = ['admin', 'org:admin', 'content_creator', 'org:content_creator'];
    17  
    18  class UserService {
    19    private clerk: Clerk | null = null;
```

## 03-015
- **id**: `B03-015`
- **corpus site**: `corpus/services/clerkenstein.md:274-288` (paragraph)
- **citation**: `sentinel/go.mod:8`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/sentinel/go.mod`  (54 lines)

**CLAIMING UNIT**

```md
> **⚠️ The `v0.34.3` pin is the ARTIFACT's, and it is BEHIND the platform** (v2.8 M257x iter-23). At platform
> `2adcf71`, `app/go.mod` reads `colony v0.35.2` — and `app` is the service an injected demo actually runs.
> **`sentinel` has since moved too** — it is on `colony v0.35.2` at `f2c46190` (`sentinel/go.mod:8`), taken
> there by `88036d7` *"chore(deps): update dependencies to latest versions"* — so the softening clause this
> paragraph used to carry (*"`sentinel` and `storage` are still on `v0.34.3`"*) is **false for `sentinel`**
> and now rests on the frozen `storage` repo alone, which nothing clones or builds. **Both live Go services
> are on `v0.35.2`; the artifact's pin is behind the whole live platform, not part of it.** A
> `clerk-deploy-1` score taken against `v0.34.3` **is not measuring the binary under test**. This is precisely
> the drift the deployment DNA exists to catch, so re-run `deployrun` against `v0.35.2` before quoting 7/7 as
> current. (`app` is likewise on `clerk-sdk-go/v2 v2.7.0`, not `v2.6.0` — `CHECK-M257x-iter22-clerk-sdk-drift`.)
That drop-in is **identity-agnostic** (straight-through claim mapping — it extracts whatever the token
carries, not a hard-coded user). Its contract is checked at *compile time* and scored by the
`alignment/cmd/deployrun` runner (the `clerk-deploy-1` DNA). `cmd/` ships the supporting standalone tools:
`mintpk` (the authoritative publishable-key minter) and `fake-fapi` / `fake-bapi` (standalone fake servers
for demos).
```

**CITED CONTENT**

```
     5  require (
     6  	connectrpc.com/connect v1.20.0
     7  	github.com/Blank-Xu/sql-adapter v1.2.1
     8  	github.com/anthropos-work/colony v0.35.2
     9  	github.com/anthropos-work/proto v1.210.0
    10  	github.com/casbin/casbin/v3 v3.10.0
    11  	github.com/google/uuid v1.6.0
```

## 03-016
- **id**: `B03-016`
- **corpus site**: `corpus/services/clerkenstein.md:309-325` (bullet)
- **citation**: `clerk-frontend/server.go:35-67`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/rosetta-extensions/clerkenstein/clerk-frontend/server.go`  (715 lines)

**CLAIMING UNIT**

```md
- **…and it WAS unbounded and uncached until M220 — ✅ FIXED, kept here because the failure mode is still
  worth recognising.** As documented at M218, `clerk-frontend/server.go` fetched the bundle with a bare
  **`http.Get`** (`http.DefaultClient`, **`Timeout: 0`**) and held no server-side cache, so every cold page
  load re-fetched from jsdelivr. **M220 closed it:** `clerk-frontend/server.go:35-67` now serves the
  clerk-js bundle **from disk** with the CDN as a *bounded* fallback — `clerkJSFetchTimeout = 15s` on an
  explicit `clerkJSClient` (commented *"Explicitly NOT http.DefaultClient"*), a disk cache at
  `FAKE_FAPI_CLERKJS_CACHE`, and a test asserting no `http.Get(` survives on that path. **A slow or
  blocked jsdelivr is therefore NO LONGER a plausible cause of a long demo login** — look elsewhere.
  The consequences below describe the pre-M220 behaviour, in order of severity:
  - next-web's **entire authenticated tree is client-gated on clerk-js**, so this sits squarely **on the
    login path**. Measured at **0.17–0.19 s healthy** — but **~127 s if egress blackholes**, with *no
    timeout to cut it short*. It is an **unbounded internet dependency in the login path of a demo the
    corpus describes as self-contained**.
  - The bring-up's egress pre-check curls from the **host, not from inside the container**
    (`up-injected.sh`), so it can pass green while the container cannot reach the CDN at all.
  - **No DNA gene covers `GET /npm/`** — the proxy is **alignment-invisible**, so vendoring the bundle and
    bounding the timeout is a **gate-free** change.
```

**CITED CONTENT**

```
    32  	return defaultClerkJSCDN
    33  }
    34  
    35  // ── M220 (S6/h): the clerk-js bundle is served FROM DISK; the CDN is a BOUNDED fallback. ────────────────
    36  //
    37  // WHAT WAS WRONG. handleClerkJSBundle proxied clerk.browser.js live from cdn.jsdelivr.net on EVERY full page
    38  // load, with `http.Get` — i.e. `http.DefaultClient`, whose `Timeout` is **0: unbounded**. There was no cache
    39  // of any kind. next-web's entire authenticated tree is client-gated on clerk-js, so this put an UNBOUNDED
    40  // INTERNET DEPENDENCY ON THE LOGIN PATH of a demo the corpus calls self-contained: 0.2 s when the CDN is
    41  // healthy, ~127 s when egress blackholes (a blackholed TCP connect, not a refused one — nothing ever fires a
    42  // timeout, so the browser just hangs on a white page and the presenter's demo is dead).
    43  //
    44  // THE FIX, in the order the item names it:
    45  //  1. SERVE FROM DISK. A cache dir (FAKE_FAPI_CLERKJS_CACHE) is consulted first. A hit never touches the
    46  //     network — the steady state of every demo after the first fetch of a given bundle version.
    47  //  2. CDN AS A BOUNDED FALLBACK. A miss fetches with a client that HAS a timeout, then writes the bytes to
    48  //     the cache (atomic tmp+rename) so the next load — and every later demo sharing the dir — is a disk hit.
    49  //  3. FAIL FAST AND LOUD. A dead CDN now costs clerkJSFetchTimeout, not forever.
    50  //
    51  // The cache dir is mounted from a BOX-LEVEL path (not per-demo), so demo-2 on a box that has already run
    52  // demo-1 never reaches the internet for this at all. Cache-key = the request path, flattened (see cachePath):
    53  // the path carries the exact package@version, so two clerk-js versions never collide and a stale entry is
    54  // impossible — an immutable-by-construction key, the same property the M45 prompt-hash cache relies on.
    55  const (
    56  	// clerkJSFetchTimeout bounds the CDN fallback. The whole point of the item: `Timeout: 0` on the login
    57  	// path is not a timeout, it is a hang. 15 s is generous for a ~600 KB script on a healthy link and still
    58  	// bounded enough that a presenter sees an error instead of a white page.
    59  	clerkJSFetchTimeout = 15 * time.Second
    60  	// clerkJSMaxBytes caps what we will read from the CDN and commit to the cache.
```

## 03-017
- **id**: `B03-017`
- **corpus site**: `corpus/services/cms.md:3-92` (paragraph)
- **citation**: `cms/terraform/main.tf:39`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/terraform/main.tf`  (192 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **cms-in-app v8.0** (`app` **v1.360.0**, July 2026), the standalone `cms` Go microservice has been
> **merged into the `app` monolith** (the service the platform calls "backend"). CMS no longer runs as a
> separate service **in production**. Its subgraph is gone from the supergraph, and its ECS service is
> **scaled to zero, not deleted** — `cms/terraform/main.tf:39` `service_desired_count = 0` — and this is
> the one M810 row whose **terraform module block** has not moved: do not read jobsimulation's teardown
> onto it (`6092c6d2` destroyed that module's service block outright).
> **⚠️ But cms HAS taken an M810 step since, and the corpus's "it has not moved" was becoming stale:**
> `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml` under the
> subject *"the cms ECR repository is decommissioned (M810)"*, its body stating that M810 *"deletes
> `module "cms_euwest1"` from the platform's `services.tf`, which destroys the ECS service and the
> production-cms ECR repository"* — the workflow went because it *"would try to push an image into a
> registry that no longer exists."* **So the two measured facts in this repo point opposite ways** (a
> module block that still declares the service; a CI commit asserting the registry is already gone), and
> the deletion itself lands in `infrastructure`, **which has never been in any clone set we have.**
> **Do not assert either way** — see the scope note below, and the fenced map, which states the same limit. It was the **fourth** engine consolidated into `app`, after
> [skiller](./skiller.md), [skillpath](./skillpath.md) and [jobsimulation](./jobsimulation.md) — **not the
> last.** The v9.0 program (2026-08-04) then folded [`storage`](./storage.md),
> [`messenger`](./messenger.md) and [`customerio-sync`](./customerio-sync.md), and platform `838d907`
> (merged `0c91421`, 2026-08-05) deleted all three containers the next day. See the fenced map,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> **✅ The husk is GONE locally, and M809 has landed (re-measured at platform `0c91421`).**
> There is no `cms` compose service, no `cms` entry in `repos.yml` (4 entries: app, sentinel,
> next-web-app, studio-desk) and no `cms` profile. Nor is there a `CMS_RPC_ADDR` any more: M809
> re-pointe
```

**CITED CONTENT**

```
    36    tags                           = var.tags
    37    aws_region                     = var.aws_region
    38    project                        = local.project
    39    service_desired_count          = 0
    40    service_cpu                    = local.service_cpu
    41    service_memory                 = local.service_memory
    42    health_check_path              = "/_meta"
```

## 03-018
- **id**: `B03-018`
- **corpus site**: `corpus/services/cms.md:3-92` (paragraph)
- **citation**: `app/knowledge/service-dependencies.md:52`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/service-dependencies.md`  (122 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **cms-in-app v8.0** (`app` **v1.360.0**, July 2026), the standalone `cms` Go microservice has been
> **merged into the `app` monolith** (the service the platform calls "backend"). CMS no longer runs as a
> separate service **in production**. Its subgraph is gone from the supergraph, and its ECS service is
> **scaled to zero, not deleted** — `cms/terraform/main.tf:39` `service_desired_count = 0` — and this is
> the one M810 row whose **terraform module block** has not moved: do not read jobsimulation's teardown
> onto it (`6092c6d2` destroyed that module's service block outright).
> **⚠️ But cms HAS taken an M810 step since, and the corpus's "it has not moved" was becoming stale:**
> `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml` under the
> subject *"the cms ECR repository is decommissioned (M810)"*, its body stating that M810 *"deletes
> `module "cms_euwest1"` from the platform's `services.tf`, which destroys the ECS service and the
> production-cms ECR repository"* — the workflow went because it *"would try to push an image into a
> registry that no longer exists."* **So the two measured facts in this repo point opposite ways** (a
> module block that still declares the service; a CI commit asserting the registry is already gone), and
> the deletion itself lands in `infrastructure`, **which has never been in any clone set we have.**
> **Do not assert either way** — see the scope note below, and the fenced map, which states the same limit. It was the **fourth** engine consolidated into `app`, after
> [skiller](./skiller.md), [skillpath](./skillpath.md) and [jobsimulation](./jobsimulation.md) — **not the
> last.** The v9.0 program (2026-08-04) then folded [`storage`](./storage.md),
> [`messenger`](./messenger.md) and [`customerio-sync`](./customerio-sync.md), and platform `838d907`
> (merged `0c91421`, 2026-08-05) deleted all three containers the next day. See the fenced map,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> **✅ The husk is GONE locally, and M809 has landed (re-measured at platform `0c91421`).**
> There is no `cms` compose service, no `cms` entry in `repos.yml` (4 entries: app, sentinel,
> next-web-app, studio-desk) and no `cms` profile. Nor is there a `CMS_RPC_ADDR` any more: M809
> re-pointe
```

**CITED CONTENT**

```
    49  >
    50  > **There are no external callers of app's RPC mux left.** `messenger` was the last one — it used to
    51  > reach the users, cms, jobsimulation and skiller surfaces at
    52  > `http://backend.internal.anthropos:8081`, and folding it in at v9.0 closed that edge. The mux is
    53  > kept because it is how the in-process domains are wired, not because something outside dials it.
    54  > App also keeps emitting `JOBSIMULATION_STREAM` as an in-process loopback — it feeds the
    55  > real consumers (XP/skills/quota/assignment link/AI Readiness); the `LocalJobsimulationSession`
```

## 03-019
- **id**: `B03-019`
- **corpus site**: `corpus/services/cms.md:3-92` (paragraph)
- **citation**: `app/main.go:1205-1211`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> As of **cms-in-app v8.0** (`app` **v1.360.0**, July 2026), the standalone `cms` Go microservice has been
> **merged into the `app` monolith** (the service the platform calls "backend"). CMS no longer runs as a
> separate service **in production**. Its subgraph is gone from the supergraph, and its ECS service is
> **scaled to zero, not deleted** — `cms/terraform/main.tf:39` `service_desired_count = 0` — and this is
> the one M810 row whose **terraform module block** has not moved: do not read jobsimulation's teardown
> onto it (`6092c6d2` destroyed that module's service block outright).
> **⚠️ But cms HAS taken an M810 step since, and the corpus's "it has not moved" was becoming stale:**
> `6efa1d5` (merged `f38c0c4`, 2026-08-04) **deleted** `.github/workflows/build-production.yml` under the
> subject *"the cms ECR repository is decommissioned (M810)"*, its body stating that M810 *"deletes
> `module "cms_euwest1"` from the platform's `services.tf`, which destroys the ECS service and the
> production-cms ECR repository"* — the workflow went because it *"would try to push an image into a
> registry that no longer exists."* **So the two measured facts in this repo point opposite ways** (a
> module block that still declares the service; a CI commit asserting the registry is already gone), and
> the deletion itself lands in `infrastructure`, **which has never been in any clone set we have.**
> **Do not assert either way** — see the scope note below, and the fenced map, which states the same limit. It was the **fourth** engine consolidated into `app`, after
> [skiller](./skiller.md), [skillpath](./skillpath.md) and [jobsimulation](./jobsimulation.md) — **not the
> last.** The v9.0 program (2026-08-04) then folded [`storage`](./storage.md),
> [`messenger`](./messenger.md) and [`customerio-sync`](./customerio-sync.md), and platform `838d907`
> (merged `0c91421`, 2026-08-05) deleted all three containers the next day. See the fenced map,
> [`platform-migration-status.md`](../architecture/platform-migration-status.md).
>
> **✅ The husk is GONE locally, and M809 has landed (re-measured at platform `0c91421`).**
> There is no `cms` compose service, no `cms` entry in `repos.yml` (4 entries: app, sentinel,
> next-web-app, studio-desk) and no `cms` profile. Nor is there a `CMS_RPC_ADDR` any more: M809
> re-pointe
```

**CITED CONTENT**

```
  1202  	// standalone cms. Active whenever the Directus edge is configured (the release sets it);
  1203  	// the external client the switch was seeded with is only the construction-time placeholder.
  1204  	cmsReaderSw.set(cmsRPCServer)
  1205  	// M805: consume the cms studio + ai_video Asynq queue in-process (the app is the sole
  1206  	// consumer post-release — the standalone cms takes no traffic). The consumer polls the SAME
  1207  	// DB index the enqueue client writes to (audit R2). The studio gen.py/postgen.py pipeline
  1208  	// is argv-safe (M809b H-1 fixed).
  1209  	cmsWorker := cmsworker.NewServer(redisAddr, cmsWorkerIndex, logger)
  1210  	wg.Go(func() {
  1211  		defer cancelServerContext()
  1212  		if err := cmsWorker.Start(serverContext, cmsManagers.Studio, cmsManagers.AiVideo); err != nil {
  1213  			logger.Info("shutting down the cms worker", "error", err)
  1214  		}
```

## 03-020
- **id**: `B03-020`
- **corpus site**: `corpus/services/cms.md:100-100` (bullet)
- **citation**: `app/internal/cms/studio/studioManager.go:1099-1101`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/studioManager.go`  (1224 lines)

**CLAIMING UNIT**

```md
3. **Runs the AI generation pipeline** in-process. The Python project `anthropos-studio-room` is pulled into the image and dispatched as a subprocess **in argv (exec) form — never through a shell** (`app/internal/cms/studio/studioManager.go:1099-1101` @ `app` `ad9f3c49`; `git grep -n '"bash"' ad9f3c49 -- '*.go'` over the whole tree returns **0**); the Go side dispatches generation work, the Python code executes it against **OpenAI, Azure OpenAI or Anthropic** — those three and no others. The provider registry is a three-entry dict: `{'openai': OpenAIProvider, 'azure': AzureProvider, 'anthropic': AnthropicProvider}` (`services/ai.py:705-708` @ `anthropos-studio-room` `aeec036` v0.51.1), and `services/ai.py:1-2` imports only `openai`/`anthropic`. **There is no Mistral path in the Python *generation* engine** — but `mistralai` is **not** unimported. `tools/pdf2md.py:24` does `from mistralai import Mistral` (client at `:96`, `model="mistral-ocr-latest"` at `:127`): a **standalone CLI OCR utility**, one leg of the `tools/r3.py` offline PDF→markdown chain, that nothing on the generation path calls — `gen.py` never imports `tools`, nothing outside `tools/` references it, and no Go caller exists (Go execs **two** studio scripts and neither is `pdf2md.py`: `studio/gen.py` at `studioManager.go:119` and `studio/postgen.py` at `:1045`, both @ `app b948604f`). `git -C app/studio grep -i mistral aeec036a` returns **22 hits in 3 files** (`requirements.txt:8`, `tools/pdf2md.py`, `tools/r3.py`), not one. So Mistral is **OCR-only on both sides** — Go-side for studio attachments, Python-side for that offline tool — and on the generation path on neither (see the Downstream-dependencies bullet below, and [`studio-room.md`](./studio-room.md) for the grep caveat that hid `tools/`).
```

**CITED CONTENT**

```
  1096  // runCommand executes name+args in argv (exec) form — NEVER through a shell. Callers pass a
  1097  // program and a discrete argument slice; nothing is string-interpolated into a command line,
  1098  // so shell metacharacters in any argument are inert (M809b H-1/M-1).
  1099  func (s StudioManager) runCommand(ctx context.Context, name string, args []string) error {
  1100  	s.logger.Info("Running command", "command", name, "args", args)
  1101  	pycmd := exec.CommandContext(ctx, name, args...)
  1102  	pycmd.Env = studioSubprocessEnv()
  1103  
  1104  	// Get stderr pipe
```

## 03-021
- **id**: `B03-021`
- **corpus site**: `corpus/services/cms.md:104-109` (paragraph)
- **citation**: `app/internal/skillpath/session.go:205-207`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/skillpath/session.go`  (1058 lines)

**CLAIMING UNIT**

```md
> [!IMPORTANT]
> **CMS owns content; the runtime engines own state.** Do not conflate the **skill-path engine** with skill-path content, or the **`jobsimulation`** service with simulation content. Those are **runtime/session engines** that hold *no* content and reference CMS artifacts **by ID**:
> - **The [skill-path engine](./skillpath.md)** (merged into `app` — "skillpath-in-app", M502→M507; formerly the standalone `skillpath` service) tracks per-user progression *state* (`SkillPathSession → ChapterSession → StepSession`, progress %); it reads the skill-path *structure* it tracks against from the **cms domain in-process** — `app/internal/skillpath/session.go:205-207` (`// cms-in-app deseam: cms is in-process`) calls `contentread.CmsContentReader.GetSkillPathDomain`. It was a `CMS_RPC_ADDR` Connect-RPC hop until both merged into `app`.
> - **[`jobsimulation`](./jobsimulation.md)** runs the interactive simulation *session*; it reads the simulation *definition* it runs from the cms domain **in-process** (it was a `cms.GetSimulation` Connect-RPC hop until both merged into `app`) — it has no `DIRECTUS_BASE_ADDR` of its own, so all its content reads go *through* CMS.
>
> So **content = CMS/Directus; the like-named service = the state machine over that content.** This split is the source of a recurring naming confusion — see the [Service Taxonomy](../architecture/service_taxonomy.md) and [Architecture Overview](../architecture/architecture_overview.md) content-vs-runtime callouts.
```

**CITED CONTENT**

```
   202  }
   203  
   204  func (u *SessionManager) getSkillPath(ctx context.Context, skillPathId uuid.UUID, version *string) (*skillpath.SkillPath, error) {
   205  	// cms-in-app deseam: cms is in-process — read the hydrated domain struct
   206  	// directly (no proto round-trip).
   207  	skillPathDomain, err := u.cms.GetSkillPathDomain(ctx, skillPathId, version)
   208  	if err != nil {
   209  		u.logger.Error("failed to fetch skill path", "error", err)
   210  		return nil, fmt.Errorf("failed to fetch skill path: %w", err)
```

## 03-022
- **id**: `B03-022`
- **corpus site**: `corpus/services/cms.md:116-116` (bullet)
- **citation**: `cms/go.mod:3`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/go.mod`  (127 lines)

**CLAIMING UNIT**

```md
* **Language**: Go 1.26 (primary — `cms/go.mod:3` `go 1.26.4`) + Python 3.11 (studio-room)
```

**CITED CONTENT**

```
     1  module github.com/anthropos-work/cms
     2  
     3  go 1.26.4
     4  
     5  require (
     6  	connectrpc.com/connect v1.20.0
```

## 03-023
- **id**: `B03-023`
- **corpus site**: `corpus/services/cms.md:118-118` (bullet)
- **citation**: `cms/cmd/root.go:77`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/cmd/root.go`  (287 lines)

**CLAIMING UNIT**

```md
* **Ports**: **8080 (GraphQL/HTTP), 8081 (Connect-RPC) — the binary's own defaults**, and now the only ones there are: `cms/cmd/root.go:77` `cmp.Or(os.Getenv("PORT"), "8080")` / `:78` `cmp.Or(os.Getenv("RPC_PORT"), "8081")`. The **8090 / 8091** pair quoted throughout this corpus was **compose-supplied by a service that no longer exists**: `docker-compose.yml` set `PORT=8090` (`:169`) / `RPC_PORT=8091` (`:173`) and published `8090:8090` / `8091:8091` (`:154-155`) — **at `2adcf71`**. At `0dab54d` there is no `cms` service, so nothing sets them and nothing is published; **8090/8091 are historical, not an address you can reach.** The domain's live surface is `backend`'s (`:8082/graphql/query`, RPC on `:8083`)
```

**CITED CONTENT**

```
    74  		)
    75  		defer colony.FlushLogEvents()
    76  
    77  		port := cmp.Or(os.Getenv("PORT"), "8080")
    78  		rpcPort := cmp.Or(os.Getenv("RPC_PORT"), "8081")
    79  
    80  		authnClerk, err := clerk.NewProvider(os.Getenv("CLERK_SECRET_KEY"))
```

## 03-024
- **id**: `B03-024`
- **corpus site**: `corpus/services/cms.md:201-201` (bullet)
- **citation**: `app/knowledge/service-dependencies.md:52`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/knowledge/service-dependencies.md`  (122 lines)

**CLAIMING UNIT**

```md
* **RPC**: `app/internal/cms/rpcsrv` — served on app's single RPC mux, and **every caller is in-process**. `messenger` was the last external one: M809 pointed its `CMS_RPC_ADDR` at `http://backend:8083` — **`d11a403` moved exactly two variables on that block, `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`; the other two already read `http://backend:8083` at `d11a403^` and were untouched (measured M257x iter-115)** — and `838d907` deleted the messenger service and that variable with it, so compose sets it nowhere. **No `.tf` file in any clone names `http://backend.internal.anthropos:8081`** — 0 hits measured 2026-08-06 over all 44 tracked `.tf` files in the 13 `stack-demo` repos at each clone's own HEAD, and 0 again over the 59 `.tf` files a raw filesystem sweep of that workspace finds. The literal does occur in the clone set — **6 times, none of them terraform**; the count and its per-repo derivation are stated once, in [`backend.md`](./backend.md)'s *RPC re-pointed, then un-set* bullet, and are not restated here. The one that matters is a **markdown KB page** — `app/knowledge/service-dependencies.md:52` @ `app` `ad9f3c49` — which is not terraform, and which puts it in the **past** tense: *"it used to reach the users, cms, jobsimulation and skiller surfaces at `http://backend.internal.anthropos:8081`, and folding it in at v9.0 closed that edge"*, under the heading *"**There are no external callers of app's RPC mux left.**"* **And the production declaration is not measurable from this repo at all:** it lives in `infrastructure`, which has never been in any clone set — as `:18` of this same file already says — so no *"still names"* claim can be made here in either direction. See [`platform-migration-status.md`](../architecture/platform-migration-status.md) for the fenced unmeasurable-claims convention.
```

**CITED CONTENT**

```
    49  >
    50  > **There are no external callers of app's RPC mux left.** `messenger` was the last one — it used to
    51  > reach the users, cms, jobsimulation and skiller surfaces at
    52  > `http://backend.internal.anthropos:8081`, and folding it in at v9.0 closed that edge. The mux is
    53  > kept because it is how the in-process domains are wired, not because something outside dials it.
    54  > App also keeps emitting `JOBSIMULATION_STREAM` as an in-process loopback — it feeds the
    55  > real consumers (XP/skills/quota/assignment link/AI Readiness); the `LocalJobsimulationSession`
```

## 03-025
- **id**: `B03-025`
- **corpus site**: `corpus/services/cms.md:220-234` (bullet)
- **citation**: `app/internal/cms/studio/markdownManager.go:10`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/markdownManager.go`  (144 lines)

**CLAIMING UNIT**

```md
* AI providers — **OpenAI / Azure OpenAI / Anthropic** for the `studio/` Python generation pipeline
  (`services/ai.py:705-708`). **Mistral is NOT one of them**: every use of it is **OCR**, never generation. The Go one —
  `app/internal/cms/studio/markdownManager.go:10` imports `internal/cms/studio/mistralocr` and `:30`
  builds the client (`mistralocr.New(aiKey)` inside `NewMarkdownManager`, re-derived at `app`
  `2035f9a` — a **pin**, not a moving label: that was origin/main on 2026-08-05, and both offsets still name the same two constructs at today's origin/main, `ad9f3c49` (re-checked 2026-08-06); it was `:11`/`:19` and a `mistral.NewMistral(nil, MISTRAL_API_KEY)` call before the key-plumbing
  fix that stopped it reading `os.Getenv` behind the caller's back).
  Its single use is `OCRProcess` (document → markdown) on the studio attachment path
  (`studioManager.go:531` *"supported ocr content types for mistral ocr"*, `:583`, and `xlsx.go:13` — xlsx is
  rendered locally precisely because Mistral OCR rejects it). Nothing generates through it. There is also a
  **Python-side** Mistral OCR user in the same image — `app/studio/tools/pdf2md.py:24`
  `from mistralai import Mistral` (`mistral-ocr-latest`), a standalone CLI **no Go caller and no `gen.py`
  path dispatches** — `tools/r3.py:139`/`:190`/`:199-206` DOES exec it as step 2 of the offline chain, so
  the flat *"nothing dispatches it"* this line carried is withdrawn
  (`git -C app/studio grep -i mistral aeec036a` → 22 hits / 3 files; `git -C app grep -- studio/`
  returns 0 because `studio/` is untracked in `app`, `app/.gitignore:79`)
```

**CITED CONTENT**

```
     7  	"strings"
     8  
     9  	"github.com/anthropos-work/app/internal/ai"
    10  	"github.com/anthropos-work/app/internal/cms/studio/mistralocr"
    11  )
    12  
    13  type MarkdownManager struct {
```

## 03-026
- **id**: `B03-026`
- **corpus site**: `corpus/services/cms.md:220-234` (bullet)
- **citation**: `app/studio/tools/pdf2md.py:24`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/studio/tools/pdf2md.py`  (233 lines)

**CLAIMING UNIT**

```md
* AI providers — **OpenAI / Azure OpenAI / Anthropic** for the `studio/` Python generation pipeline
  (`services/ai.py:705-708`). **Mistral is NOT one of them**: every use of it is **OCR**, never generation. The Go one —
  `app/internal/cms/studio/markdownManager.go:10` imports `internal/cms/studio/mistralocr` and `:30`
  builds the client (`mistralocr.New(aiKey)` inside `NewMarkdownManager`, re-derived at `app`
  `2035f9a` — a **pin**, not a moving label: that was origin/main on 2026-08-05, and both offsets still name the same two constructs at today's origin/main, `ad9f3c49` (re-checked 2026-08-06); it was `:11`/`:19` and a `mistral.NewMistral(nil, MISTRAL_API_KEY)` call before the key-plumbing
  fix that stopped it reading `os.Getenv` behind the caller's back).
  Its single use is `OCRProcess` (document → markdown) on the studio attachment path
  (`studioManager.go:531` *"supported ocr content types for mistral ocr"*, `:583`, and `xlsx.go:13` — xlsx is
  rendered locally precisely because Mistral OCR rejects it). Nothing generates through it. There is also a
  **Python-side** Mistral OCR user in the same image — `app/studio/tools/pdf2md.py:24`
  `from mistralai import Mistral` (`mistral-ocr-latest`), a standalone CLI **no Go caller and no `gen.py`
  path dispatches** — `tools/r3.py:139`/`:190`/`:199-206` DOES exec it as step 2 of the offline chain, so
  the flat *"nothing dispatches it"* this line carried is withdrawn
  (`git -C app/studio grep -i mistral aeec036a` → 22 hits / 3 files; `git -C app grep -- studio/`
  returns 0 because `studio/` is untracked in `app`, `app/.gitignore:79`)
```

**CITED CONTENT**

```
    21  from typing import Optional
    22  
    23  import tqdm
    24  from mistralai import Mistral
    25  from dotenv import load_dotenv
    26  
    27  MISTRAL_API_KEY = None
```

## 03-027
- **id**: `B03-027`
- **corpus site**: `corpus/services/cms.md:295-313` (paragraph)
- **citation**: `app/internal/cms/studio/studioManager.go:1096-1098`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/cms/studio/studioManager.go`  (1224 lines)

**CLAIMING UNIT**

```md
> ⚠️ **CORRECTED M257x iter-115 — this note asserted the exact inversion of a shipped security property.**
> It said the Go service invokes `python3 studio/gen.py ...` / `studio/postgen.py` **"via `bash -c`"**, in the
> present tense and with no HISTORICAL marker (unlike both of its neighbours in this section). The live code is
> the opposite, deliberately: at `app` `ad9f3c49`,
> `app/internal/cms/studio/studioManager.go:1096-1098` reads *"runCommand executes name+args in **argv (exec)
> form — NEVER through a shell**… nothing is string-interpolated into a command line (M809b H-1/M-1)"*, and
> `:1101` is `pycmd := exec.CommandContext(ctx, name, args...)`. `:100-103` says it in the caller's own words —
> *"It MUST NOT be interpolated into a shell … **No `bash -c`**"* — and `:119` is
> `s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))`. Measured: `git grep -n '"bash"'
> ad9f3c49 -- '*.go'` over the whole `app` tree returns **0**.
>
> **What survives.** Dev mode does auto-provision a venv at `studio/studio-venv` and run
> `pip3 install -r studio/requirements.txt` — as **fixed argv** (`:126`, `:129`), *"previously chained into the
> same `bash -c` string that carried the tainted args"* (`:122-124`). Paths are still `studio/...`, not from
> inside `studio/`. For standalone Python work, use a venv to match the service's behavior.
>
> **Why it read as true:** the claim is still correct about the **frozen** `cms` repo
> (`ca50c817:internal/studio/studioManager.go:967` = `exec.Command("bash", "-c", command)`) — right about the
> dead code, wrong about the shipped code, and wrong about the direction of a deliberate hardening.
```

**CITED CONTENT**

```
  1093  	return "...(truncated)..." + s[len(s)-max:]
  1094  }
  1095  
  1096  // runCommand executes name+args in argv (exec) form — NEVER through a shell. Callers pass a
  1097  // program and a discrete argument slice; nothing is string-interpolated into a command line,
  1098  // so shell metacharacters in any argument are inert (M809b H-1/M-1).
  1099  func (s StudioManager) runCommand(ctx context.Context, name string, args []string) error {
  1100  	s.logger.Info("Running command", "command", name, "args", args)
  1101  	pycmd := exec.CommandContext(ctx, name, args...)
```

## 03-028
- **id**: `B03-028`
- **corpus site**: `corpus/services/cms.md:295-313` (paragraph)
- **citation**: `internal/studio/studioManager.go:967`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/cms/internal/studio/studioManager.go`  (1093 lines)

**CLAIMING UNIT**

```md
> ⚠️ **CORRECTED M257x iter-115 — this note asserted the exact inversion of a shipped security property.**
> It said the Go service invokes `python3 studio/gen.py ...` / `studio/postgen.py` **"via `bash -c`"**, in the
> present tense and with no HISTORICAL marker (unlike both of its neighbours in this section). The live code is
> the opposite, deliberately: at `app` `ad9f3c49`,
> `app/internal/cms/studio/studioManager.go:1096-1098` reads *"runCommand executes name+args in **argv (exec)
> form — NEVER through a shell**… nothing is string-interpolated into a command line (M809b H-1/M-1)"*, and
> `:1101` is `pycmd := exec.CommandContext(ctx, name, args...)`. `:100-103` says it in the caller's own words —
> *"It MUST NOT be interpolated into a shell … **No `bash -c`**"* — and `:119` is
> `s.runCommand(ctx, pyBin, append([]string{"studio/gen.py"}, tokens...))`. Measured: `git grep -n '"bash"'
> ad9f3c49 -- '*.go'` over the whole `app` tree returns **0**.
>
> **What survives.** Dev mode does auto-provision a venv at `studio/studio-venv` and run
> `pip3 install -r studio/requirements.txt` — as **fixed argv** (`:126`, `:129`), *"previously chained into the
> same `bash -c` string that carried the tainted args"* (`:122-124`). Paths are still `studio/...`, not from
> inside `studio/`. For standalone Python work, use a venv to match the service's behavior.
>
> **Why it read as true:** the claim is still correct about the **frozen** `cms` repo
> (`ca50c817:internal/studio/studioManager.go:967` = `exec.Command("bash", "-c", command)`) — right about the
> dead code, wrong about the shipped code, and wrong about the direction of a deliberate hardening.
```

**CITED CONTENT**

```
   964  
   965  func (s StudioManager) runCommand(_ context.Context, command string) error {
   966  	s.logger.Info("Running command", "command", command)
   967  	pycmd := exec.Command("bash", "-c", command)
   968  
   969  	// Get stderr pipe
   970  	stderr, err := pycmd.StderrPipe()
```

## 03-029
- **id**: `B03-029`
- **corpus site**: `corpus/services/coursebuilder.md:48-55` (bullet)
- **citation**: `internal/coursebuilder/bedrock.go:105-114`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/coursebuilder/bedrock.go`  (276 lines)

**CLAIMING UNIT**

```md
*   **LLM usage — the backend is SELECTED AT START-UP, and production is the first-party Anthropic API, not Bedrock.** `internal/coursebuilder/bedrock.go:105-114` returns an `api.anthropic.com` client with bare model ids whenever `ANTHROPIC_API_KEY` is set, reporting `ModelBackendName() == "anthropic-api"` (`:98-104`, logged at `main.go:770` @ `app` `b948604` v1.366.0); the Bedrock `eu-west-1` path via `internal/askengine/bedrock.go` is the fallback when it is not. **In production the key is required** — at `app` `ad9f3c49`, `terraform/variables.tf:759-763` declares it `sensitive` with no default, `ssm.tf:328-333` creates the SecureString parameter and `main.tf:757-758` injects it from the SSM ARN — so the shipped path is the direct API. ⚠️ **These read `variables.tf:635-638` / `main.tf:555` until M257x iter-115, and both had drifted onto DIFFERENT SUBJECTS**, which is the maximally misleading failure: at `ad9f3c49`, `variables.tf:631-645` is a cms-in-app secrets comment block and `main.tf:555` is `"name": "DIRECTUS_BASE_ADDR"`, so a reader opening the cited line saw a Directus variable and read the whole production-key claim as wrong. **The substantive claim is true and was re-derived, not assumed** — only the citation pair was false. (`ssm.tf:328` verified at both `b948604f` and `ad9f3c49`; the sentence carries no ref of its own — the only pin in the bullet is the parenthetical `@ app b948604` attached to `main.go:770`, which is a different file — so it grades at the checkout, and now names it.) Models:
    *   **Author/patch model**: Opus 4.8 (`eu.anthropic.claude-opus-4-8`, env `CB_AUTHOR_MODEL`; streaming, no
        sampling params — Opus 4.8 rejects them — at 32 K max_tokens).
    *   **Grader model**: Sonnet 4.6 (`eu.anthropic.claude-sonnet-4-6`, env `CB_GRADER_MODEL`; deliberately a
        different model to avoid self-grading bias; `temperature=0`).
    *   **Cover images**: OpenAI `gpt-image-2` (`imagegen/openai.go`, separate key).
    *   **Prompt caching** (Wave 2b): static system prompts marked ephemeral cache-control (~85 % input-side saving
        on the static prefix), tracked via `CacheUsageTracker`.
```

**CITED CONTENT**

```
   102  	return "bedrock"
   103  }
   104  
   105  // newUnderlyingClient picks the backend for one model role:
   106  // ANTHROPIC_API_KEY present → the first-party Anthropic API (with the
   107  // model id normalized to its bare form); absent → AWS Bedrock, the
   108  // legacy path, byte-for-byte what shipped before the switch existed.
   109  func newUnderlyingClient(ctx context.Context, modelID string) (*askengine.BedrockClient, error) {
   110  	if key := strings.TrimSpace(os.Getenv(AnthropicAPIKeyEnv)); key != "" {
   111  		return askengine.NewAnthropicClientWithModel(key, directModelID(modelID))
   112  	}
   113  	return askengine.NewBedrockClientWithModel(ctx, modelID)
   114  }
   115  
   116  // authorNoSamplingUnderlying is the subset of *askengine.BedrockClient
   117  // the author path needs. Kept as an interface so tests can supply a
```

## 03-030
- **id**: `B03-030`
- **corpus site**: `corpus/services/coursebuilder.md:48-55` (bullet)
- **citation**: `main.go:770`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
*   **LLM usage — the backend is SELECTED AT START-UP, and production is the first-party Anthropic API, not Bedrock.** `internal/coursebuilder/bedrock.go:105-114` returns an `api.anthropic.com` client with bare model ids whenever `ANTHROPIC_API_KEY` is set, reporting `ModelBackendName() == "anthropic-api"` (`:98-104`, logged at `main.go:770` @ `app` `b948604` v1.366.0); the Bedrock `eu-west-1` path via `internal/askengine/bedrock.go` is the fallback when it is not. **In production the key is required** — at `app` `ad9f3c49`, `terraform/variables.tf:759-763` declares it `sensitive` with no default, `ssm.tf:328-333` creates the SecureString parameter and `main.tf:757-758` injects it from the SSM ARN — so the shipped path is the direct API. ⚠️ **These read `variables.tf:635-638` / `main.tf:555` until M257x iter-115, and both had drifted onto DIFFERENT SUBJECTS**, which is the maximally misleading failure: at `ad9f3c49`, `variables.tf:631-645` is a cms-in-app secrets comment block and `main.tf:555` is `"name": "DIRECTUS_BASE_ADDR"`, so a reader opening the cited line saw a Directus variable and read the whole production-key claim as wrong. **The substantive claim is true and was re-derived, not assumed** — only the citation pair was false. (`ssm.tf:328` verified at both `b948604f` and `ad9f3c49`; the sentence carries no ref of its own — the only pin in the bullet is the parenthetical `@ app b948604` attached to `main.go:770`, which is a different file — so it grades at the checkout, and now names it.) Models:
    *   **Author/patch model**: Opus 4.8 (`eu.anthropic.claude-opus-4-8`, env `CB_AUTHOR_MODEL`; streaming, no
        sampling params — Opus 4.8 rejects them — at 32 K max_tokens).
    *   **Grader model**: Sonnet 4.6 (`eu.anthropic.claude-sonnet-4-6`, env `CB_GRADER_MODEL`; deliberately a
        different model to avoid self-grading bias; `temperature=0`).
    *   **Cover images**: OpenAI `gpt-image-2` (`imagegen/openai.go`, separate key).
    *   **Prompt caching** (Wave 2b): static system prompts marked ephemeral cache-control (~85 % input-side saving
        on the static prefix), tracked via `CacheUsageTracker`.
```

**CITED CONTENT**

```
   767  	appWorker := worker.NewServer(redisAddr)
   768  	workerHandler := tasks.NewHandler(
   769  		ent,
   770  		repo,
   771  		aiClient,
   772  		ocrClient,
   773  		pub,
```

## 03-031
- **id**: `B03-031`
- **corpus site**: `corpus/services/coursebuilder.md:48-55` (bullet)
- **citation**: `terraform/variables.tf:759-763`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/terraform/variables.tf`  (764 lines)

**CLAIMING UNIT**

```md
*   **LLM usage — the backend is SELECTED AT START-UP, and production is the first-party Anthropic API, not Bedrock.** `internal/coursebuilder/bedrock.go:105-114` returns an `api.anthropic.com` client with bare model ids whenever `ANTHROPIC_API_KEY` is set, reporting `ModelBackendName() == "anthropic-api"` (`:98-104`, logged at `main.go:770` @ `app` `b948604` v1.366.0); the Bedrock `eu-west-1` path via `internal/askengine/bedrock.go` is the fallback when it is not. **In production the key is required** — at `app` `ad9f3c49`, `terraform/variables.tf:759-763` declares it `sensitive` with no default, `ssm.tf:328-333` creates the SecureString parameter and `main.tf:757-758` injects it from the SSM ARN — so the shipped path is the direct API. ⚠️ **These read `variables.tf:635-638` / `main.tf:555` until M257x iter-115, and both had drifted onto DIFFERENT SUBJECTS**, which is the maximally misleading failure: at `ad9f3c49`, `variables.tf:631-645` is a cms-in-app secrets comment block and `main.tf:555` is `"name": "DIRECTUS_BASE_ADDR"`, so a reader opening the cited line saw a Directus variable and read the whole production-key claim as wrong. **The substantive claim is true and was re-derived, not assumed** — only the citation pair was false. (`ssm.tf:328` verified at both `b948604f` and `ad9f3c49`; the sentence carries no ref of its own — the only pin in the bullet is the parenthetical `@ app b948604` attached to `main.go:770`, which is a different file — so it grades at the checkout, and now names it.) Models:
    *   **Author/patch model**: Opus 4.8 (`eu.anthropic.claude-opus-4-8`, env `CB_AUTHOR_MODEL`; streaming, no
        sampling params — Opus 4.8 rejects them — at 32 K max_tokens).
    *   **Grader model**: Sonnet 4.6 (`eu.anthropic.claude-sonnet-4-6`, env `CB_GRADER_MODEL`; deliberately a
        different model to avoid self-grading bias; `temperature=0`).
    *   **Cover images**: OpenAI `gpt-image-2` (`imagegen/openai.go`, separate key).
    *   **Prompt caching** (Wave 2b): static system prompts marked ephemeral cache-control (~85 % input-side saving
        on the static prefix), tracked via `CacheUsageTracker`.
```

**CITED CONTENT**

```
   756  # Distinct from studio_anthropic_api_key, which belongs to the cms Studio
   757  # subprocess — separate keys keep one domain's rate limit or key rotation
   758  # from affecting the other.
   759  variable "anthropic_api_key" {
   760    type        = string
   761    description = "Anthropic API key for the Course Builder model clients (ANTHROPIC_API_KEY)."
   762    sensitive   = true
   763  }
   764  
```

## 03-032
- **id**: `B03-032`
- **corpus site**: `corpus/services/coursebuilder.md:63-72` (bullet)
- **citation**: `main.go:766-779`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
*   **How to find the API**: **HTTP + SSE only — there is NO GraphQL subgraph and NO Connect-RPC for Course
    Builder.** The routes are an Echo group mounted in `internal/web/backend/backend.go` under **`/coursebuilder`**,
    behind `cors + authn (Clerk JWT via colony/authn) + courseBuilderAccessGate` (**org-admin-gated**). Route table:
    `internal/web/backend/coursebuilder/handler.go:Register`. The whole group is **unmounted** when NEITHER
    model backend can build a client — i.e. when the author or grader constructor fails, which on the
    Bedrock path is typically missing AWS creds, but **not** on missing AWS creds alone once
    `ANTHROPIC_API_KEY` is set (see the backend-selection note above). When that happens the routes stay
    unmounted deliberately, **so callers get a clean 404 instead of a half-wired endpoint** — there is no
    half-working surface (`main.go:766-779` @ `app` `b948604` v1.366.0 — the comment at `:766-769`, the
    two `logger.Warn(… routes disabled)` arms at `:774` / `:778`).
```

**CITED CONTENT**

```
   763  	var wg conc.WaitGroup
   764  
   765  	// Asynq worker. The task handler is built here in the root with explicit
   766  	// dependencies (no *app.App) and passed into the server's Start.
   767  	appWorker := worker.NewServer(redisAddr)
   768  	workerHandler := tasks.NewHandler(
   769  		ent,
   770  		repo,
   771  		aiClient,
   772  		ocrClient,
   773  		pub,
   774  		cfg,
   775  		jobsimDj.SimManager,
   776  		aiReadinessManager,
   777  		resourceManager,
   778  		linkedinService,
   779  		orgManager,
   780  		authnManager,
   781  		logoDevManager,
   782  		storagePublicClient,
```

## 03-033
- **id**: `B03-033`
- **corpus site**: `corpus/services/coursebuilder.md:113-118` (bullet)
- **citation**: `main.go:824-826`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
*   **Key env vars**: `CB_AUTHOR_MODEL` (default `eu.anthropic.claude-opus-4-8`), `CB_GRADER_MODEL` (default
    `eu.anthropic.claude-sonnet-4-6`), `CB_IMAGE_MODEL` (default `gpt-image-2`), **`OPENAI_KEY`** (the cover generator reads this — `main.go:824-826` @ `app` `b948604` v1.366.0; the `COURSEBUILDER_OPENAI_IMAGE_KEY` this doc used to name was deleted at app `68c24512` and survives only in stale in-repo markdown, so setting it fixes nothing),
    `COURSEBUILDER_PLANNER_ENABLED` (multi-chapter kill-switch), `CB_SOURCE_DISTILL`, `COURSEBUILDER_EMAILS_ENABLED`,
    `COURSEBUILDER_MAX_MONTHLY_COGS_USD` (default **500**, the primary per-org ceiling), `COURSEBUILDER_MAX_DAILY_COGS_USD`
    (default 0 = off), `AWS_REGION`, `CLERK_SECRET_KEY`. Cost/rate: session cap `DefaultSessionsPerOrgPerDay=50`
    (code constant); credits `course.build`=**5/chapter**, `course.refine`=**1**, `course.translate`=**1/locale**.
```

**CITED CONTENT**

```
   821  	swJobRoleManager := jobrole.NewJobRoleManager(logger, skillerWorkerEnt, orgManager, swSkillTaxonomyManager, skillerAIManager, swEmbeddingManager, pub, workerClient.Client, redisClientStream)
   822  	swLocalizationManager := localization.NewManager(repository.NewEntRepository(skillerWorkerEnt))
   823  	swSkillerManager := skiller.NewSkillerManager(logger, swJobRoleManager, swSkillTaxonomyManager, swLocalizationManager)
   824  	swSkillManager := skill.NewSkillManager(logger, skillerWorkerEnt, aiClient, swSkillerManager, jobsimDj.SimManager, pub)
   825  	swRoleManager := roles.NewRoleManager(logger, skillerWorkerEnt, authz, authnManager, swSkillerManager, swSkillManager)
   826  
   827  	skillerWorker := worker.NewSkillerServer(os.Getenv("REDIS_ADDR"), skillerWorkerConcurrency)
   828  	wg.Go(func() {
   829  		defer cancelServerContext()
```

## 03-034
- **id**: `B03-034`
- **corpus site**: `corpus/services/customerio-sync.md:3-28` (paragraph)
- **citation**: `app/internal/customeriosync/doc.go:4-5`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/customeriosync/doc.go`  (80 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4-5` — the sentence wraps the line break).
>
> **Every `app` anchor in this file is read at `app` `ad9f3c49`** — `origin/main` *and* the demo's build
> pin on 2026-08-06, and byte-identical at `2035f9a4`. Ref pinned M257x iter-102: these citations were
> unpinned and present-tense, and `env_guards.go` **did not exist** at the demo's former pin `b948604f`
> (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so several of them resolved at no
> ref this document named.
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`
> @ platform `0c91421`).
> Turning it on writes **real** contacts, which is the whole reason for the switch.
>
> **The name is a fossil.** The destination has been **Brevo**, not Customer.io, since long before
> the fold; the package doc says so outright, and the read model `public.customer_io_sync_table`
> carries the same fossil. The in-app manager is constructed with `os.Getenv("BREVO_KEY")`
> (`app/main.go:395`).
```

**CITED CONTENT**

```
     1  // Package customeriosync is the internalized customerio-sync service: the one-way
     2  // push of platform users into Brevo as marketing contacts.
     3  //
     4  // Provenance. Ported out of the customerio-sync repository (v0.19.3), the last of the
     5  // Go services to be folded into app. The name is historical — the service moved off
     6  // Customer.io to Brevo long ago and never got renamed; the DB view it reads
     7  // (public.customer_io_sync_table) carries the same fossil.
     8  //
```

## 03-035
- **id**: `B03-035`
- **corpus site**: `corpus/services/customerio-sync.md:3-28` (paragraph)
- **citation**: `app/env_guards.go:62`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4-5` — the sentence wraps the line break).
>
> **Every `app` anchor in this file is read at `app` `ad9f3c49`** — `origin/main` *and* the demo's build
> pin on 2026-08-06, and byte-identical at `2035f9a4`. Ref pinned M257x iter-102: these citations were
> unpinned and present-tense, and `env_guards.go` **did not exist** at the demo's former pin `b948604f`
> (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so several of them resolved at no
> ref this document named.
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`
> @ platform `0c91421`).
> Turning it on writes **real** contacts, which is the whole reason for the switch.
>
> **The name is a fossil.** The destination has been **Brevo**, not Customer.io, since long before
> the fold; the package doc says so outright, and the read model `public.customer_io_sync_table`
> carries the same fossil. The in-app manager is constructed with `os.Getenv("BREVO_KEY")`
> (`app/main.go:395`).
```

**CITED CONTENT**

```
    59  // predicate they never read.
    60  const (
    61  	envMessengerEnabled      = "MESSENGER_ENABLED"
    62  	envCustomerIOSyncEnabled = "CUSTOMERIO_SYNC_ENABLED"
    63  )
    64  
    65  // resolveSubsystemSwitch reads one switch. It is strict in BOTH directions, which is
```

## 03-036
- **id**: `B03-036`
- **corpus site**: `corpus/services/customerio-sync.md:3-28` (paragraph)
- **citation**: `app/main.go:286`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4-5` — the sentence wraps the line break).
>
> **Every `app` anchor in this file is read at `app` `ad9f3c49`** — `origin/main` *and* the demo's build
> pin on 2026-08-06, and byte-identical at `2035f9a4`. Ref pinned M257x iter-102: these citations were
> unpinned and present-tense, and `env_guards.go` **did not exist** at the demo's former pin `b948604f`
> (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so several of them resolved at no
> ref this document named.
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`
> @ platform `0c91421`).
> Turning it on writes **real** contacts, which is the whole reason for the switch.
>
> **The name is a fossil.** The destination has been **Brevo**, not Customer.io, since long before
> the fold; the package doc says so outright, and the read model `public.customer_io_sync_table`
> carries the same fossil. The in-app manager is constructed with `os.Getenv("BREVO_KEY")`
> (`app/main.go:395`).
```

**CITED CONTENT**

```
   283  	// marketing contacts — and neither is inferred from the environment any more: see
   284  	// resolveSubsystemSwitch for why unset is off on a laptop and fatal in production.
   285  	messengerEnabled := mustSubsystemSwitch(envMessengerEnabled)
   286  	customerIOSyncEnabled := mustSubsystemSwitch(envCustomerIOSyncEnabled)
   287  	logger.Info("subsystem switches",
   288  		envMessengerEnabled, messengerEnabled,
   289  		envCustomerIOSyncEnabled, customerIOSyncEnabled)
```

## 03-037
- **id**: `B03-037`
- **corpus site**: `corpus/services/customerio-sync.md:3-28` (paragraph)
- **citation**: `docker-compose.yml:84-92`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4-5` — the sentence wraps the line break).
>
> **Every `app` anchor in this file is read at `app` `ad9f3c49`** — `origin/main` *and* the demo's build
> pin on 2026-08-06, and byte-identical at `2035f9a4`. Ref pinned M257x iter-102: these citations were
> unpinned and present-tense, and `env_guards.go` **did not exist** at the demo's former pin `b948604f`
> (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so several of them resolved at no
> ref this document named.
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`
> @ platform `0c91421`).
> Turning it on writes **real** contacts, which is the whole reason for the switch.
>
> **The name is a fossil.** The destination has been **Brevo**, not Customer.io, since long before
> the fold; the package doc says so outright, and the read model `public.customer_io_sync_table`
> carries the same fossil. The in-app manager is constructed with `os.Getenv("BREVO_KEY")`
> (`app/main.go:395`).
```

**CITED CONTENT**

```
    81        - AWS_DEFAULT_REGION=eu-west-1
    82        - STORAGE_S3_BUCKET=production-storage20240826131618541000000005
    83        - STORAGE_S3_PUBLIC_BUCKET=production-storage-public20240919130721114900000001
    84        # messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container
    85        # too, but deliberately have NO variables here. Both reach outside the process on a
    86        # stream or a timer — they send mail and rewrite Brevo contacts — so app gates them
    87        # behind MESSENGER_ENABLED / CUSTOMERIO_SYNC_ENABLED, which default to OFF on a
    88        # developer machine (ENVIRONMENT=development is what makes unset mean off).
    89        # Pinning them to `false` here would override .env and make opting in impossible
    90        # without editing this file. To exercise either one locally, set it in .env — and
    91        # know that messenger then attaches to the LIVE Redis consumer group and
    92        # customerio-sync writes real Brevo contacts.
    93        - SUPABASE_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    94        - COPILOT_DB_CONN=postgresql://postgres@postgresql:5432/postgres?sslmode=disable
    95      networks:
```

## 03-038
- **id**: `B03-038`
- **corpus site**: `corpus/services/customerio-sync.md:3-28` (paragraph)
- **citation**: `app/main.go:395`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
> ## ⚠️ Merged into `app` — no longer a standalone service
>
> Platform **`838d907`** (merged **`0c91421`**, 2026-08-05, *"drop the storage, messenger and
> customerio-sync containers"*) deleted the compose service. The domain now runs **in-process inside
> `backend`** as `app/internal/customeriosync/` — a **relocation, not a rewrite**, ported out of
> `customerio-sync` v0.19.3 and, by its own package doc, *"the last of the Go services to be folded
> into app"* (`app/internal/customeriosync/doc.go:4-5` — the sentence wraps the line break).
>
> **Every `app` anchor in this file is read at `app` `ad9f3c49`** — `origin/main` *and* the demo's build
> pin on 2026-08-06, and byte-identical at `2035f9a4`. Ref pinned M257x iter-102: these citations were
> unpinned and present-tense, and `env_guards.go` **did not exist** at the demo's former pin `b948604f`
> (`git -C stack-demo/app ls-tree b948604f -- env_guards.go` → empty), so several of them resolved at no
> ref this document named.
>
> It does **not** run by default. `app` gates it behind `CUSTOMERIO_SYNC_ENABLED`
> (`app/env_guards.go:62`), resolved before anything connects to anything (`app/main.go:286`), and
> unset means **off** on a developer machine — `ENVIRONMENT=development` is what makes that so.
> Compose deliberately sets **no value** for it on the `backend` block: pinning one there would
> override `.env` and make opting in impossible without editing compose (`docker-compose.yml:84-92`
> @ platform `0c91421`).
> Turning it on writes **real** contacts, which is the whole reason for the switch.
>
> **The name is a fossil.** The destination has been **Brevo**, not Customer.io, since long before
> the fold; the package doc says so outright, and the read model `public.customer_io_sync_table`
> carries the same fossil. The in-app manager is constructed with `os.Getenv("BREVO_KEY")`
> (`app/main.go:395`).
```

**CITED CONTENT**

```
   392  	// count(*)) are recorded in customeriosync/store.go.
   393  	var customerIOSyncManager *customeriosync.Manager
   394  	if customerIOSyncEnabled {
   395  		customerIOSyncManager = customeriosync.New(logger, copilotDB, os.Getenv("BREVO_KEY"))
   396  	}
   397  	// ent here is the primary-DB ORM client (public schema). The AI Readiness
   398  	// cycles/snapshots/narratives tables it owns live there; the analytics reads
```

## 03-039
- **id**: `B03-039`
- **corpus site**: `corpus/services/customerio-sync.md:58-60` (paragraph)
- **citation**: `app/main.go:393-396`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
It runs on `app`'s **shared** analytics pool (`copilotDB`), not a pool of its own — one query every
ten minutes does not earn a standing allocation against the platform's connection budget
(`app/main.go:393-396`).
```

**CITED CONTENT**

```
   390  	// the sync window, because the read model's refresh_date is computed and the
   391  	// predicate can't be pushed down. Numbers and their trap (don't benchmark this with
   392  	// count(*)) are recorded in customeriosync/store.go.
   393  	var customerIOSyncManager *customeriosync.Manager
   394  	if customerIOSyncEnabled {
   395  		customerIOSyncManager = customeriosync.New(logger, copilotDB, os.Getenv("BREVO_KEY"))
   396  	}
   397  	// ent here is the primary-DB ORM client (public schema). The AI Readiness
   398  	// cycles/snapshots/narratives tables it owns live there; the analytics reads
   399  	// use copilotDB.
```

## 03-040
- **id**: `B03-040`
- **corpus site**: `corpus/services/customerio-sync.md:138-145` (paragraph)
- **citation**: `app/env_guards.go:98-104`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/env_guards.go`  (202 lines)

**CLAIMING UNIT**

```md
Runs inside the `backend` ECS task, gated by the same `CUSTOMERIO_SYNC_ENABLED` switch — where unset
is **fatal** rather than off, so a deployed environment must state its intent. The mechanism is
`app/env_guards.go:98-104` (`resolveSubsystemSwitch`'s `case "":` returns an error when `deployed`) via
`mustSubsystemSwitch`'s `log.Fatalf` at `:87`. (`app/main.go:284` is only the **comment** pointing at it,
not the mechanism — anchor corrected M257x iter-102.)
**Scope note:** whether its own ECS task / image / terraform module have been torn down was **not
measured** in this pass — the fold and the container deletion are local-compose and `app`-source
facts. Do not read them as the production teardown.
```

**CITED CONTENT**

```
    95  		return true, nil
    96  	case "false", "0", "no", "off":
    97  		return false, nil
    98  	case "":
    99  		if deployed {
   100  			return false, fmt.Errorf("%s is not set. Deployed environments must state this "+
   101  				"explicitly (\"true\" or \"false\") — an unset switch would silently disable the "+
   102  				"subsystem, and for messenger that means every email is dropped while the service "+
   103  				"reports healthy. Set it in app/terraform/main.tf's container environment", key)
   104  		}
   105  		return false, nil
   106  	default:
   107  		return false, fmt.Errorf("%s=%q is not a boolean. Use true/1/yes/on or false/0/no/off; "+
```

## 03-041
- **id**: `B03-041`
- **corpus site**: `corpus/services/customerio-sync.md:138-145` (paragraph)
- **citation**: `app/main.go:284`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/main.go`  (1640 lines)

**CLAIMING UNIT**

```md
Runs inside the `backend` ECS task, gated by the same `CUSTOMERIO_SYNC_ENABLED` switch — where unset
is **fatal** rather than off, so a deployed environment must state its intent. The mechanism is
`app/env_guards.go:98-104` (`resolveSubsystemSwitch`'s `case "":` returns an error when `deployed`) via
`mustSubsystemSwitch`'s `log.Fatalf` at `:87`. (`app/main.go:284` is only the **comment** pointing at it,
not the mechanism — anchor corrected M257x iter-102.)
**Scope note:** whether its own ECS task / image / terraform module have been torn down was **not
measured** in this pass — the fold and the container deletion are local-compose and `app`-source
facts. Do not read them as the production teardown.
```

**CITED CONTENT**

```
   281  	// Resolve the outbound-effect switches FIRST, before anything connects to
   282  	// anything. Both subsystems act on the world — one sends mail, one rewrites
   283  	// marketing contacts — and neither is inferred from the environment any more: see
   284  	// resolveSubsystemSwitch for why unset is off on a laptop and fatal in production.
   285  	messengerEnabled := mustSubsystemSwitch(envMessengerEnabled)
   286  	customerIOSyncEnabled := mustSubsystemSwitch(envCustomerIOSyncEnabled)
   287  	logger.Info("subsystem switches",
```

## 03-042
- **id**: `B03-042`
- **corpus site**: `corpus/services/gotenberg.md:7-7` (paragraph)
- **citation**: `app/internal/web/backend/coursebuilder/extract.go:77-81`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/web/backend/coursebuilder/extract.go`  (86 lines)

**CLAIMING UNIT**

```md
In the Anthropos platform it exists for one consumer — but **not to produce a PDF anybody ever sees.** `backend` uses it as a **text-extraction / OCR intermediate**: an uploaded document the text extractor can't read (or reads and finds no text in) is converted to PDF *in memory*, the text is pulled straight back out of those bytes, and the PDF is discarded. It is **never stored, never served, never displayed**; no PDF here is a platform artifact. Both call sites throw it away in the next statement — `app/internal/web/backend/coursebuilder/extract.go:77-81` converts, then immediately `converter.ConvertFromReader(bytes.NewReader(pdf), "application/pdf")` to get the course-builder source text; `app/internal/worker/tasks/user_import_resume_2d.go:68-74` converts a DOCX résumé **only** to feed the OCR client (`ocrInput = pdfBytes`) after the plain-text path found nothing readable. Measured at `app` `9d00a313` v1.367.0.
```

**CITED CONTENT**

```
    74  			if err != nil {
    75  				return "", err
    76  			}
    77  			pdf, err := converter.ConvertToPDF(context.Background(), gotenbergURL, data, name)
    78  			if err != nil {
    79  				return "", fmt.Errorf("convert %s via gotenberg: %w", mimeType, err)
    80  			}
    81  			return converter.ConvertFromReader(bytes.NewReader(pdf), "application/pdf")
    82  		}
    83  		return converter.ConvertFromReader(r, mimeType)
    84  	}
```

## 03-043
- **id**: `B03-043`
- **corpus site**: `corpus/services/gotenberg.md:7-7` (paragraph)
- **citation**: `app/internal/worker/tasks/user_import_resume_2d.go:68-74`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/app/internal/worker/tasks/user_import_resume_2d.go`  (155 lines)

**CLAIMING UNIT**

```md
In the Anthropos platform it exists for one consumer — but **not to produce a PDF anybody ever sees.** `backend` uses it as a **text-extraction / OCR intermediate**: an uploaded document the text extractor can't read (or reads and finds no text in) is converted to PDF *in memory*, the text is pulled straight back out of those bytes, and the PDF is discarded. It is **never stored, never served, never displayed**; no PDF here is a platform artifact. Both call sites throw it away in the next statement — `app/internal/web/backend/coursebuilder/extract.go:77-81` converts, then immediately `converter.ConvertFromReader(bytes.NewReader(pdf), "application/pdf")` to get the course-builder source text; `app/internal/worker/tasks/user_import_resume_2d.go:68-74` converts a DOCX résumé **only** to feed the OCR client (`ocrInput = pdfBytes`) after the plain-text path found nothing readable. Measured at `app` `9d00a313` v1.367.0.
```

**CITED CONTENT**

```
    65  			h.logger.Info("document has no readable text, attempting OCR fallback")
    66  			ocrInput := fileBytes
    67  			if mimetype == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && h.cfg.GotenbergURL != "" {
    68  				pdfBytes, convErr := converter.ConvertToPDF(ctx, h.cfg.GotenbergURL, fileBytes, "resume.docx")
    69  				if convErr != nil {
    70  					h.logger.Error("Gotenberg DOCX-to-PDF conversion failed", "error", convErr)
    71  				} else {
    72  					h.logger.Info("DOCX converted to PDF for OCR", "pdf_size", len(pdfBytes))
    73  					ocrInput = pdfBytes
    74  				}
    75  			}
    76  			ocrText, ocrErr := h.ocrClient.OCRProcess(ctx, ocrInput)
    77  			if ocrErr != nil {
```

## 03-044
- **id**: `B03-044`
- **corpus site**: `corpus/services/gotenberg.md:14-14` (bullet)
- **citation**: `docker-compose.yml:183`
- **resolved file**: `/Users/marco/workspace/anthropos/rosetta/stack-demo/platform/docker-compose.yml`  (187 lines)

**CLAIMING UNIT**

```md
* **Profile**: `core` (the default), `backend`, `all` — `profiles: [core, backend, all]` (`docker-compose.yml:183`, re-derived at platform `0c91421`). The default profile is `core`, not `graphql`: `0dab54d` renamed it. Corrected M257x iter-68, re-anchored iter-87 (the anchor was `:268` at `0dab54d`; `838d907` deleted three service blocks above it and the file is now 186 lines)
```

**CITED CONTENT**

```
   180        - "3200:3200"
   181      networks:
   182        - app-network
   183      profiles: [core, backend, all]
   184  
   185  networks:
   186    app-network:
```

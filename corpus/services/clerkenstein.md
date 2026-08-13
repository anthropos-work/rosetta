# Clerkenstein

**Status:** v0.3 (v1.0 "body double" · M1 + M2 + M2b + M2c `@clerk/express` + M3 deploy/injection; v1.9 "storytelling" M37 multi-identity seat-switch; v1.10 "method acting" M39 roster org-name threading; v2.2 "panorama" M213 MagicDNS/egress; v2.3 "cue to cue" M217 self-healing demopatch gate + **M218 the roster-aware fake BAPI**; v2.8 "fast build" M257x iter-108 pin re-derivation) · **Last updated:** 2026-08-06
**Repo:** the `clerkenstein/` **section** of the `rosetta-extensions` monorepo — authored at
`.agentspace/rosetta-extensions/clerkenstein`, consumed per-stack at a pinned tag as
`stack-demo/rosetta-extensions/clerkenstein`. (It is **not** "its own git" — see *One monorepo, two clone
roles* below, which this line used to contradict.) · **Measured by:** the
[alignment framework](../architecture/alignment_testing.md)

> **The demo-patch mechanism is specified in [`../ops/demo/demopatch-spec.md`](../ops/demo/demopatch-spec.md).** It is the sanctioned **zero-platform-edit escape hatch**: patch the demo's own ephemeral clone before the image build, revert after — the canonical repos are never touched. Read it before adding or re-pinning a patch. Since M217 the gate is **self-healing**: the *anchor* is the contract, the whole-file sha is only a baseline.

> **This is a pointer.** The full, self-contained documentation now lives **in the `clerkenstein/` section's
> own knowledge base** (added in M2b): start at
> `.agentspace/rosetta-extensions/clerkenstein/knowledge/kb-index.md` — the **authoring copy**, per *One
> monorepo, two clone roles* below; a stack's pinned-tag clone carries the same tree at
> `stack-demo/rosetta-extensions/clerkenstein/…`. This page
> keeps only the platform-side orientation + the cross-links a rosetta reader needs — it deliberately does
> **not** duplicate that KB.

> **One monorepo, two clone roles.** `rosetta-extensions` is ONE private monorepo with sections
> — **eleven** of them: `alignment`, `clerkenstein`, `demo-stack`, `dev-stack`, `playthroughs`,
> `stack-core`, `stack-injection`, `stack-secrets`, `stack-seeding`, `stack-snapshot`, `stack-verify`
> (⚠️ **this listed six until M257x iter-129**, and `CLAUDE.md` listed nine — the SAME under-count in two
> places, and the repair that fixed `CLAUDE.md` this iter reached only that one cell until the complement
> read found this twin. `§5` rule 54, committed by the run that was applying it). It is
> authored / built / tested / aligned in the **authoring copy** at `.agentspace/rosetta-extensions/`, then
> **tagged**, and consumed by each stack via that stack's own **pinned-tag** clone
> (`stack-demo/rosetta-extensions @ <tag>`). So the KB above is **read from the authoring copy**; a running
> stack uses its pinned consumption copy.

## Role (platform-side orientation)

Clerkenstein is a **drop-in mock of the Clerk library** the platform uses — the *same interface*, with all
security and sync **disarmed**. It lets **demo** environments create users / orgs / admins and log in/out
with no Clerk friction (one universal credential, no live API, no webhooks, no rate limits), while platform
repos keep "thinking" they use Clerk with **zero source changes**.

It is the **first mirror produced by the M0 alignment process** (not a hand-built mock): its fidelity is
*measured* as a 0–100% alignment score against a Clerk **Alignment DNA**.

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
> `clerkenstein/alignment/dna/clerk-2.6.0.json:131`: *"M219 landed the fix … taking the Go surface
> 97.2% -> 100%."*
> The Go surface is **27/27**. `FIX-M219-bapi-org-eid` is CLOSED.
>
> **Why this matters more than the number.** Before M218, Clerkenstein scored **100% critical / 100%
> overall / 0 divergences while its fake BAPI returned the wrong human for every hero** — `GET
> /v1/users/{id}` had no gene in any of the five DNAs, and the three genes that *did* name identity all
> asserted the stub itself. **The goldens ratified the defect.** That cost ~6 s on every authenticated
> render for four releases. **A 100% that hides a lie is worse than an honest 97.2%** — which is exactly why
> restoring a clean 100% by looking away from the next stub was rejected.

Driven to the gate across: the Go surface (`clerk-sdk-go/v2 @ v2.6.0`, M1), the JS/FAPI surface (9/9 genes, `@clerk/clerk-js` v5 / `@clerk/nextjs` v6, M2), the
**multi-identity seat-switch** surface (9/9 genes, `clerk-multi-1` — the v1.9 M37 registry + active-seat
selection, so a demo can present as any seeded hero; the multi-session FAPI semantics real clerk-js exhibits
with `single_session_mode=false`), the **`@clerk/express`** Node-backend surface (**13 genes across 5
capabilities** — dependency-gated, so frequently *unmeasured*; see the ⚠ box above — `@clerk/express`
^1.3.47, M2c — RS256/JWKS, the genuine SDK *satisfied*, not reimplemented), and the **deployment/injection**
surface (7/7 genes, `clerk-deploy-1` — the disarmed `colony/authn/provider/clerk` drop-in compiles against
a real `colony`, pinned at **`v0.34.3`**, and satisfies its contract; added after **M3** showed *behavioural*
alignment ≠ *deployability* — see [`alignment_testing.md`](../architecture/alignment_testing.md#what-alignment-proves--and-what-it-doesnt-the-m3-lesson)).
The DNAs + mirror + goldens + runners live in the `clerkenstein/` section; the `/align-dna` + `/align-run`
skills + the [`alignment_testing.md`](../architecture/alignment_testing.md) doc live in rosetta, while the
`alignctl` harness is the `rosetta-extensions/alignment/` section (a sibling of `clerkenstein/`).

> **What "100%" means (and doesn't).** The score measures the mirror as *indistinguishable from the source
> goldens*. Those goldens are **hand-authored / hybrid** (decision M1-D1) — the reference behavior derived
> from the real libraries' documented + observed semantics (and, for `@clerk/express`, confirmed by driving
> the *genuine* SDK), **not** captured from a live, network-connected real-Clerk tenant. So 100% means "the
> mirror reproduces the behavior we encoded as the reference," not "diffed byte-for-byte against a running
> Clerk instance." Re-capturing goldens against a live source on a Clerk version bump is the M1b drift loop's
> job. This is the right bar for a *demo* mock; it is not a conformance certificate against production Clerk.

## Repo structure (library-named, since M2b)

The `clerkenstein/` section is organised **one dir per mocked dependency** (M2b reorg, decision M2b-D2) —
every dir below is a subdir of `rosetta-extensions/clerkenstein/`, not a repo root:

| Dir | Mocks | What it is |
|---|---|---|
| `authn/` | `colony/authn` | the provider twin — **verifies** session JWTs (offline) |
| `clerk-backend/` | `clerk-sdk-go/v2` | fake Backend API + the in-memory org store, merged |
| `clerk-frontend/` | `@clerk/clerk-js` + `@clerk/nextjs` | fake Frontend API + publishable-key codec — **mints** JWTs |
| `clerk-webhook/` | `svix` | the signed-webhook injector |
| `shared/` | — | universal-key HS256 JWT (the mint side + verify side agree here) |
| `deploy/` | `colony/authn/provider/clerk` | the disarmed provider drop-in — **deployable** into a vendored colony fork (compiles against real `colony`, **pinned `v0.34.3` — behind `app`'s `v0.35.2`**, see the ⚠️ below) |
| `cmd/` | — | standalone binaries: `mintpk` (authoritative publishable-key minter) · `fake-fapi` / `fake-bapi` (standalone fake servers for demos; `fake-fapi` loads `FAKE_FAPI_ROSTER` for M37 multi-identity) |
| `alignment/` | — | the measurement harness: `cmd/{clerkrun,jsfapirun,multirun,expressrun,deployrun}` + `dna/` (five) + `golden{,-js,-multi,-express,-deploy}/` + `scripts/` |

The browser-login → backend-verify coherence chain runs through `shared`: `clerk-frontend` mints the
HS256 universal-key JWT, `authn` verifies that exact token — pinned by the JS DNA's
`SessionToken/decoded-identity` gene (operator `exact`).

### BAPI ≠ FAPI — both twins must exist (`fix/studio`, 2026-07-27)

**A route registered on the BAPI is NOT reachable by the browser, and vice versa.** The two dirs mock two
different callers: `clerk-backend/` answers **server-to-server** SDK calls (`@clerk/express`, next-web SSR);
`clerk-frontend/` answers **clerk-js in the browser**. A membership list exists on *both* sides of real Clerk
under different paths, and Clerkenstein shipped only one of them for four releases:

| caller | route | dir |
|---|---|---|
| `@clerk/express` · next-web SSR | `GET /v1/users/{userID}/organization_memberships` | `clerk-backend/` (registered since M2) |
| **clerk-js** (`clerk.user.getOrganizationMemberships()`) | **`GET /v1/me/organization_memberships`** | `clerk-frontend/` (**added `fix/studio`**) |

The missing FAPI twin 404'd, clerk-js's `const { data } = response` threw, and studio-desk's boot burned a
**~4.05 s** 3-attempt retry ladder on **every** load (`latency-budget.md` §"Time-to-usable"). Three properties
make the route correct, each pinned by a test in `clerk-frontend/meorgmemberships_test.go`:

- **The paginated envelope, not a bare array** — a `?paginated=true` request answers
  `{"response":{"data":[…],"total_count":N},…}`, because clerk-js destructures `{ data, total_count }`
  straight off it. A bare array **is** the *"Cannot destructure property 'data'"* throw.
- **`limit`/`offset` are honoured** — clerk-js sends `limit=10&offset=0` and pages when `total_count` exceeds
  what it received, so serving the full list against `limit=10` makes it chase a second page that disagrees
  with the first. `total_count` always reports the **true** total; an offset past the end is an empty **array**
  (never `null` — clerk-js maps over it).
- **Unauthenticated ⇒ 401, never 404** — a 404 is what clerk-js *retries*; a clean 401 it handles. A signed-in
  user with **no** memberships gets an empty list at 200 (mirroring the BAPI contract).

The data needed no new assembly: `/v1/me` already returns `userRes.OrganizationMemberships`. The role keeps
Clerk's **prefixed** form (`org:admin`) as a **fidelity** choice — it is what real Clerk emits on this route.
It is *not* a hard gate requirement: studio-desk's `STUDIO_ACCESS_ROLES` accepts **both** forms
(`['admin', 'org:admin', 'content_creator', 'org:content_creator']` — `src/index.ts:96` and
`app/services/userService.ts:16`, each carrying the comment "Both the prefixed (`org:*`) and bare role keys
are accepted"), so an unprefixed `admin` would pass too.

> **Not yet a measured gene.** `clerkenstein/alignment/dna/clerk-js-5.json` has a `Me` capability for
> `GET /v1/me` but
> **no** gene for this route, so alignment scoring does not cover it — the unit tests do. Adding one needs a
> real-Clerk golden capture (`/align-dna`), i.e. a milestone, not a patch. Tracked as a known DNA gap.

### Multi-identity

**(v1.9 M37)** — `clerk-frontend` now holds a **users/orgs registry** (replacing the single
`DefaultDemoUser`) + an **active-seat selection** so a demo can **switch the active browser identity** among
the seeded heroes/orgs (the M35 stories roster) — the seat-switch the presenter cockpit's "login as" needs.
Selection is **server-authoritative** (the FAPI holds the active key, so the client view, `/v1/me`, the
token mint, and the handshake cookies all resolve the same hero): `?__clerk_identity=<key>` on the handshake
(the cockpit's [Login as] deep-link) + the `/v1/demo/{identities,select}` control plane. The single-identity
path is byte-identical (a one-member registry). Measured by the `clerk-multi-1` DNA (`alignment/cmd/multirun`,
9 genes, 100%/100%) — a *new measured surface* that holds while the existing four stay green.

> **⚠️ "Server-authoritative" means SINGLE-TENANT: one active seat per stack, no client scoping** (documented
> v2.8 M256 pre-flight; the same limitation is disclosed from the presenter side in
> [`../ops/demo/cockpit-spec.md`](../ops/demo/cockpit-spec.md) § *Limitation — one seat per stack*). The
> coherence the paragraph above sells is bought by holding the seat **process-wide**, not per client:
> `clerk-frontend/registry.go` keeps a single `activeKey` (`Registry.active()` / `Registry.Select()`), and
> `clerk-frontend/server.go`'s `type Server` keeps **one** `signedIn`, **one** `clientID`
> (`"client_clerkenstein"`, a constant) and **one** `sessID` (`"sess_clerkenstein"`, also a constant, minted in
> `establishLocked`). Three consequences a consumer must design around:
> - **`POST /v1/demo/select` (`handleSelectIdentity`) is destructive to the current session.** It re-points the
>   seat **and** sets `signedIn = false; sessID = ""` — globally. A second seat-switch anywhere on the stack
>   signs the first browser out.
> - **The read path takes NO request input.** `handleMe`, `handleToken`, `handleClient` and
>   `handleMeOrganizationMemberships` all discard (or ignore) the `*http.Request` and answer from
>   `activeUserLocked()`. `r.Cookie(...)` is called **nowhere** in `clerkenstein/` — cookies are only ever
>   *emitted*. So **per-browser `storageState` cannot isolate two identities**, and a token refresh silently
>   re-mints whoever the *current* seat is. `handleSignOut` likewise ignores its `{id}` route param and logs the
>   whole stack out.
> - **Therefore: concurrency is one-identity-at-a-time per stack.** Two people on two deeplinks, or two
>   parallel Playwright workers, will swap identities mid-flight. The sanctioned workaround is **a stack each**
>   (a fake FAPI each). Making the seat per-client — keying the registry by `__client`/cookie and threading it
>   through the `/v1/me`, token-mint, client-view and handshake surfaces — is an **auth-model change with an
>   alignment-DNA consequence** (the `clerk-multi-1` DNA has no gene for concurrent-seat isolation), not a
>   config knob.

> **⚠️ An explicit sign-out is STICKY until an explicit login — and a SEAT SWITCH is not a sign-out** (D81,
> v2.8 M256 iter-16/iter-25). `Server` carries a **`signedOut`** flag alongside `signedIn`. Three rules:
>
> - **`POST /v1/client/sessions?_method=DELETE`** — what `@clerk/clerk-js`'s `signOut()` actually sends, a POST
>   with a `_method` override and **not** a `DELETE` — sets it. Before the fix no `DELETE` route was registered
>   and nothing read `_method`, so the request **404'd**, `handleSignOut` never ran, and the next handshake
>   silently re-established the same seat. The user-visible symptom was *"I have to click logout twice"*.
>   `_method` is a **dispatch whitelist**: an override the server does not understand is ignored, never obeyed.
> - **While `signedOut` is set, a BARE handshake DECLINES** — it will not re-establish a session on its own.
>   Every *explicit* establish path (a handshake carrying an identity, a sign-in form, `POST /v1/demo/select`)
>   clears the flag, so a demo can always get back in; a missing clear on any ONE of them **strands the stack
>   signed-out**, which is why there is a test per entry door rather than one per fix.
> - **`/v1/demo/select` drops the session but must NOT set the flag** — it is a seat switch, and setting it
>   would make the cockpit's own `[Log in as]` land on a signed-out browser. This was found by driving the
>   cockpit live after five green unit tests had passed over the same code.
>
> **The guard is a FRONT-DOOR guard, not a revocation — stated because it is a real limitation.** Every test
> observes it through `GET /v1/me`, which reads the server's in-memory flag; the browser's state comes from the
> handshake cookies. A *declined* handshake still 303s back having minted an RS256 `__session`, and the only
> differentiator is an **empty `sid`** claim. Nothing revokes an already-issued token either (no `jti`, no
> denylist, 1 h `exp`), so a token captured before the sign-out keeps working. Acceptable for a deliberately
> disarmed mock on a demo — see [`../ops/safety.md`](../ops/safety.md) §3 — but it means the flag governs
> *establishment*, not *access*. Pinned by `clerk-frontend/server_test.go`:
> `TestServer_signOutOnThePathClerkJSActuallySends`, `TestServer_signOutIgnoresAMethodOverrideItDoesNotUnderstand`,
> `TestServer_seatSwitchIsNotASignOut`, `TestServer_signedOutFlagIsClearedByEveryEstablishPath`,
> `TestServer_seatSelectAfterSignOutCanLogBackIn`.

**Roster org-name threading (v1.10 M39).** The roster now carries each hero's **story org name + slug**, so a
logged-in hero's **top bar reads her real company** (e.g. "Cervato Systems") instead of the hardcoded
"Clerkenstein Demo Org". The thread is a **paired change** kept in lockstep by the roster's
`DisallowUnknownFields` decoder — the producer (`stack-seeding/seeders/roster.go`) and the consumer
(`clerk-frontend`) add the same two `org_name`/`org_slug` snake_case fields in one change, and the rext repo is
re-tagged as a whole:

- **Producer** — `RosterIdentity` (roster.go) gains `org_name`/`org_slug`, filled in `BuildRoster` from
  `st.Org.Name` + the single-sourced `orgSlugFor` (the **same** slug rule `OrgSeeder` writes to
  `public.organizations.slug`, so the roster-carried org and the seeded org can never disagree — #M39-D2).
- **Consumer** — `RosterEntry` (`clerk-frontend/registry.go`) gains the matching `org_name`/`org_slug` and
  threads them through `toDemoUser` into `DemoUser` (`resources.go`); `DemoUser.orgMemberships()` renders them
  on the FAPI org resource (`/v1/me` → the SDK's active-org → the top bar).
- **No-roster default fallback** — an empty `OrgName`/`OrgSlug` (the `DefaultDemoUser`, or any roster that omits
  the fields) falls back to the `orgNameDefault`/`orgSlugDefault` constants (`"Clerkenstein Demo Org"` /
  `"clerkenstein-demo"`), so the single-identity path is **byte-identical** and a pre-M39 roster still loads
  (the decoder rejects *unknown* fields, not *missing* ones — forward-compatible — #M39-D3).

Alignment held: the **multi-identity** (`clerk-multi-1`) + **JS/FAPI** (`clerk-js-5`) surfaces stay **9/9,
100%/100%** (the `DefaultDemoUser` goldens are unchanged — they take the default-name fallback).

**Roster org `isHiring` threading (v2.4 "casting call" M224).** The same roster→FAPI thread extends to a hiring
org's `public_metadata.isHiring`. The fake **FAPI** emits org `public_metadata.isHiring = true` in
`clerk-frontend/resources.go::orgMemberships()` — the org resource `@clerk/clerk-js`'s
`useOrganization().publicMetadata` reads — fed by `RosterEntry.org_is_hiring` → `DemoUser.OrgIsHiring`, produced by
the seeder (`RosterIdentity.org_is_hiring` ← `ResolvedStory.IsHiringOrg()` in `BuildRoster`). Only a **hiring**
story's heroes carry `true`. It is the client-side half of the `is_hiring` dual-write: the DB column
(`public.organizations.is_hiring`, the seeder's write) drives the *server*; this FAPI field drives the *browser
re-skin* (`useGetClerkOrganization` derives `isHiringOrg` from it → the "Results" nav framing / hiring cohort
treatment). Without it a demo org whose DB row says `is_hiring=true` renders as a **normal Workforce org** in the
browser. See [`hiring.md`](hiring.md) § `isHiringOrg`.

- **The align-safety rule this pins — CONDITIONAL-EMIT (#M224-D-align).** A new FAPI field is emitted **only when
  its non-default value applies** (`if u.OrgIsHiring { pm["isHiring"] = true }`, else omit). The goldens are
  captured from the existing identities; adding a key to a `shape`-graded response (`Client/signed-in`,
  `Me/universal-user`) for the *default* case would flag the gate or force a golden re-capture. Conditional-emit
  keeps every non-hiring org's `public_metadata` **byte-identically `{eid}`** — generalizing the
  `Picture`/`OrgLogo` `omitempty` pattern to non-string additions.
- **`/align-run` record (BLOCKING for any `clerk-frontend/` change).** `clerk-js-5` **100.0%/100.0%** (9/9),
  `clerk-multi-1` **100.0%/100.0%** (9/9, incl. Roster 2/2) — GREEN, no identity gene perturbed (the named
  `SessionToken/decoded-identity` critical/exact gene unaffected).
- **BAPI intentionally NOT wired (#M224-D-bapi).** The server derives hiring from the
  `public.organizations.is_hiring` **DB column**, not Clerk BAPI metadata; a `clerk-backend` change would add
  Go-SDK align surface for **no render benefit**. Optional, only if a server-side consumer ever reads
  `organization.publicMetadata.isHiring`.

**`@clerk/express` (M2c) added no new dir** — it's a *consumer* (a Node backend verifier we satisfy), so
its support is **additive**: an RS256 path (RS256 minting in `shared/` + a real JWKS from `clerk-frontend/`
+ read endpoints in `clerk-backend/`), measured by the `alignment/cmd/expressrun` runner driving the
**genuine `@clerk/backend`** — the same "verify against the real library" discipline `clerk-webhook/` uses
with `svix`. `@clerk/express` verifies RS256-via-JWKS and rejects HS256, so the RS256 path is additive (the
HS256 seams + M1/M2 gates stay green). Its `clerk-express-1` DNA includes the **`dotless-pk-rejected`** gene:
`@clerk/backend`'s `assertValidPublishableKey` (run by `clerkMiddleware` on every request) rejects a pk whose
decoded FAPI host has **no dot** — which is why the demo pk host is a **dotted** `127.0.0.1` (not `localhost`)
and why a MagicDNS FQDN (`billion.taildc510.ts.net`, also dotted) validates natively (v2.2 M213).

**The deployment/injection surface (M3) *did* add `deploy/` + `cmd/`.** Unlike the `authn/` twin (which
mocks the standalone `colony/authn` interface), the platform actually consumes `colony/authn/provider/clerk`
*inside* the `colony` module. So the **deployable** drop-in lives in `deploy/colony-authn/`: the disarmed
provider — same package, same `Clerk` type, same `NewProvider(apiKey)` signature — compiled against the
**real** `colony`, pinned at `v0.34.3`, so an injected demo app accepts Clerkenstein-minted tokens with zero
source changes.

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

### Remote HTTPS over the tailnet (v2.2 "panorama" M213)

Making a demo reachable from another machine on a **Tailscale** tailnet touches three Clerkenstein-adjacent
seams. ⚠️ **This said "opt-in via `/demo-up --public-host <magicdns>`" until run 81 and that is FALSE for
the demo path**: since v2.3 M220 (`D-DESIGN-3`) remote reach is **DEFAULT-ON via auto-discovery** — a bare
`/demo-up N` probes tailscale, mints a trusted cert and publishes the stack on the tailnet. It is
**opt-OUT** (`--no-public-host` / `DEMO_NO_PUBLIC_HOST=1`, `demo-stack/up-injected.sh:27-40,108-116`);
`--public-host` now only *forces a host and skips discovery*. **Only `/dev-up` is still opt-in.** This is
the exposure axis — a reader believed a bare `/demo-up N` stayed local. Seams — all **gated** so an unset host is byte-identical:

- **FAPI cert → `tailscale cert`.** For a MagicDNS host the fake-FAPI cert is minted via `tailscale cert` (a real
  Let's Encrypt cert **trusted tailnet-wide, no per-machine CA install**) instead of mkcert/openssl — **same output
  paths** (`<stack>/certs/fapi.{crt,key}`), so the path-only mount + `cmd/fake-fapi` `ListenAndServeTLS` are
  untouched. Falls back to the local mkcert/openssl mint (non-fatal). 90-day LE cert → renew-then-reload (M215).
  **VM caveat (proven on billion, M215):** the bring-up calls `tailscale cert` **un-sudo'd**, so the deploy VM must
  have the Tailscale **operator** set once — `sudo tailscale set --operator=<user>` — or the un-sudo'd call fails
  and the cert silently falls back to mkcert (local-trust-only → a *remote* browser sees an untrusted cert). See
  [`../ops/setup_guide.md`](../ops/setup_guide.md) §"Linux host prerequisites (for a remote/VM demo over Tailscale)". (#M213-D-CERT-1)
- **pk host stays dotted.** The publishable key is minted host-parametrically (the `--fapi-host` is the MagicDNS
  FQDN); the demo wiring pre-checks the dotted-host rule (the `dotless-pk-rejected` gene) and fails loud on a
  dotless `--public-host`. The **codec** (`clerk-frontend/key.go` `MintPublishableKey`) stays permissive — the
  alignment gene deliberately mints a dotless pk to test the consumer's rejection. (#M213-D-PK-1)
- **clerk-js egress is overridable.** The FAPI proxies the clerk-js bundle from `cdn.jsdelivr.net` (its one outbound
  dependency); **`FAKE_FAPI_CLERKJS_CDN`** overrides that base so a locked-down network can point at a mirror. (#M213-D-EGRESS-1)
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

  M218 measured it and confirmed it was **not** the cause of the 38-second login (it was healthy on
  `billion`), so the fix was **not** taken there — a runtime change would have restarted the milestone's
  5-cycle cold battery. **Routed forward to M220** (vendor the bundle; serve from disk; keep the CDN proxy
  only as a *bounded* fallback). Until then, treat a slow/blocked jsdelivr as a **plausible cause of an
  arbitrarily long demo login**.

A fourth seam — **the origins & links emission (M214)** — admits the MagicDNS/HTTPS origin everywhere a
browser→backend or cross-surface call is gated, again all gated on the knob:

- **CORS + redirects → HTTPS MagicDNS.** The injected override appends `https://$HOST:{3000,3001,9000}+off` to
  the backend's `CORS_EXTRA_ORIGINS` (the `localhost` trio is kept), and emits studio-desk's
  `CLERK_SIGN_IN_URL`/`WEB_APP_URL` requireAuth fallback at `https://$HOST:3000+off`. **Per-port HTTPS**, because
  `tailscale serve` preserves the offset port (M213 D-PROXY-2) — the browser origin is `https://$HOST:<offsetport>`,
  not a port-less 443. One scheme predicate (`browser_scheme`) flips http→localhost / https→MagicDNS uniformly.
- **The bounded patch tail** rides the **existing** sha-pinned mechanism (never a canonical repo edit): ant-academy's
  `next dev` `allowedDevOrigins` admits the MagicDNS host (the `ant-academy-dev-origins` demo-patch, env-var
  indirection so the post-hash stays fixed), and studio-desk's SPA `VITE_CLERK_SIGN_IN_URL` bakes via a gitignored
  `.env.production.local` overlay (no Dockerfile ARG). (#M214-D-SCHEME-1 / D-VITE-SIGNIN-1)

The **live cross-machine acceptance** is **M215**. The full remote-access recipe + topology:
[`../ops/demo/tailscale-serve.md`](../ops/demo/tailscale-serve.md); bring-up mechanics:
[`recipe-browser-login.md §B`](../ops/demo/recipe-browser-login.md).

## Read next (in the `clerkenstein/` section)

- **`knowledge/kb-index.md`** — the KB entry point (scope, architecture, alignment, injection, coverage).
- **`knowledge/scope.md`** — what it is/isn't + the disarmed-by-design properties.
- **`knowledge/architecture.md`** — the dir layout, public API surface, and the universal-key JWT flow.
- **`knowledge/injection.md`** — the four per-library injection recipes (each labelled built+gated /
  spike-proven / recipe-only) for disarming the platform's Clerk with no platform-code change.
- **`knowledge/alignment.md`** — how fidelity is measured against a pinned Clerk version + the **drift
  runbook** (M1b: `gate.sh` / `drift-check.sh` exit-code contract; re-`/align-dna` + re-`/align-run` on a
  Clerk bump). The scripts are **mirror-side**, at `clerkenstein/alignment/scripts/` — the reusable
  `rosetta-extensions/alignment/` harness section has no `scripts/` dir; `ALIGN_DIR` (default
  `../../alignment`, relative to `clerkenstein/alignment/`) is how they find the sibling harness's
  `alignctl`. The weekly-cron CI workflow they reference (`clerkenstein/.github/workflows/alignment.yml`)
  is **git-tracked but inert** — GitHub Actions only reads `.github/workflows` at the *repository root*, so
  the gate is a manual `/align-run`; see
  [`alignment_testing.md`](../architecture/alignment_testing.md#how-m1-m1b-m2-and-m2c-consume-this).
- **`knowledge/coverage-index.md`** — per-package test coverage + known gaps.
- Per-library `README.md` in each dir for the code-level entry point.

## See also (rosetta)
- [Alignment Testing](../architecture/alignment_testing.md) — the framework that measures this mirror.
- [Clerk integration](clerk-integration.md) — the real Clerk surface Clerkenstein mirrors.
- [Frontend architecture](../architecture/frontend_architecture.md) · [next-web-app](next-web-app.md) — the
  `@clerk/nextjs` consumers the `clerk-frontend/` server stands in for.
- [Webhook setup](../ops/webhook_setup.md) — the real Clerk webhook path the `clerk-webhook/` injector replays into.

# iter-04 — Decisions

## D11 — The mock's two halves must agree about who the user is

**The defect.** Clerkenstein has **two** identity surfaces, and they disagreed:

| surface | what it told the platform | correct? |
|---|---|---|
| **FAPI** — mints the session JWT that `app`'s `authn` middleware authenticates | the hero's **real** eid (`40921b2e-…`), from the roster | ✅ |
| **BAPI** — answers next-web's **server-side** `currentUser()` | **`11111111-…`**, a hardcoded stub, for **every** id | ❌ |

`clerk-backend/resources.go` said so in its own comment — **`// Disarmed: any id → the demo user.`** That was
*correct* when a demo had exactly one user. The **Stories & Heroes** model (M35) gave every hero their own eid,
and this line was never revisited.

**Why it cost 6 seconds.** next-web's `(authenticated)/layout.tsx` passes `user.externalId` — **the BAPI's
value** — into GraphQL as the user id. `app`'s resolver is not an authorization check but an **identity equality
check**:

```go
if currentUser.ID() != userID {   // JWT's hero  vs  BAPI's stub
    return "", fmt.Errorf("user does not have permission to set studio preferences")
```

The fetchers rethrow under `retry: 2` / `retryDelay 2 s → 4 s`, so every authenticated render paid
`3 × ~33 ms + (2 s + 4 s) ≈ 6.0 s` of **pure backoff**. Both vantages, because both block on the same layout.

**Decision.** The BAPI becomes **roster-aware**: it already mounted the roster (`FAKE_FAPI_ROSTER`) but used it
only to seed **memberships** — now it also seeds each hero's **identity** (`Store.SeedUserIdentity`), and
`getUser` reports that hero's real eid/email/name.

**The general lesson (→ protocol doc).** *A mirror that is faithful on one surface and stubbed on another is not
"partially faithful" — it is **inconsistent**, and inconsistency is worse than an honest stub.* The platform
cross-checks the two surfaces against each other; the mock must therefore keep them **coherent**, not merely
individually plausible. Every identity the FAPI can mint, the BAPI must be able to resolve **to the same person**.

## D12 — The fallback is the alignment contract

The fix could have simply *replaced* the stub. It doesn't: with **no roster** the identity map is **empty** and
`userRes` falls back to the historical universal-demo-user payload, **byte-identically**.

That is deliberate and load-bearing: the **alignment runner never mounts a roster**, so its payloads must not
move. Proven, not assumed — `/align-run` across **all five surfaces** after the change:

| surface | genes | overall | critical | divergences |
|---|---|---|---|---|
| `clerkrun` | 22/22 | **100.0%** | **100.0%** | 0 |
| `jsfapirun` | 9/9 | **100.0%** | **100.0%** | 0 |
| `multirun` | 9/9 | **100.0%** | **100.0%** | 0 |
| `deployrun` | 7/7 | **100.0%** | **100.0%** | 0 |
| `expressrun` | 13/13 | **100.0%** | **100.0%** | 0 |

Gate = critical ≥ 100 / overall ≥ 95. **All five exit 0. Nothing moved.**

**Note on the guard's scope.** The milestone's BLOCKING clause names `clerkenstein/clerk-frontend/`; this change
is in `clerk-backend/`. The guard was run **anyway** — the `DistinctIdentity/*` genes are `critical` and the
critical gate has **no partial credit**, so "the clause does not technically apply" is not a reason to skip a
five-minute proof. (`expressrun` is dependency-gated — it needs a `node_modules` carrying `@clerk/backend`; it
was pointed at a real one rather than left unscored, because an unmeasured surface is not a passing surface.)

## D13 — The measurement chain now closes end-to-end

Three independent observations, taken at the same instants on the live green stack, name the same defect:

1. **the harness** — `[slow-body] … BODY took 6104 ms`, reproducible to **±3 ms** across both vantages;
2. **the Cosmo router** — SSR calls (`user_agent: "node"`) logged at **ERROR**, HTTP 200, latency 2.4 ms, at
   **T → T+2 s → T+4 s** (the retry ladder, visible in the timestamps);
3. **`app`** — at those same instants: `userPreferences user does not have permission to set studio preferences`,
   with the sentry scope naming the **real** hero (`dan.rossi3@cervato-systems.com`) while the BAPI was handing
   next-web the stub.

The arithmetic had already predicted the *shape* before any of it was read: 6.1 s **cannot** be a blackhole
(10.5 s/attempt), so it had to be a **fast-failing** fetch under the 2 s/4 s ladder. **The number told us what
kind of bug to look for.** That is the whole value of the iter-02 harness.

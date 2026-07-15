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

> ### ⚠ CORRECTED AT THE M218 CLOSE (2026-07-14) — this table is not reproducible, and the Go row has moved
>
> Two corrections, recorded rather than quietly overwritten:
>
> 1. **`expressrun` is DEPENDENCY-GATED and is now UNMEASURABLE.** Without `@clerk/express` `node_modules`
>    the runner cannot build: it exits **rc=2 with NO score** — and *nothing in the tooling treats that as a
>    failure*. The final harden pass reproduced this **identically at baseline `f296e5e`**, so it is
>    **pre-existing, not a regression** — but it means **only 4 of the 5 surfaces could be re-measured**, and
>    the "all five" claim above **cannot be reproduced on this box**. The `13/13 100%` row must therefore be
>    read as *"what was recorded on the box that ran iter-04"*, not as a standing, checkable fact.
>    **Absence of a score is not a passing score.** → `TEST-M219-expressrun-dep-gate`.
> 2. **The `clerkrun` (Go) row is now `26/27` → 97.2% overall / 100% critical, deliberately.** The final
>    harden pass added the `GetUser` per-hero identity genes that this milestone's own fix should always have
>    had — and, in doing so, surfaced **F-11**: the fake BAPI *also* fabricates the **org** eid. It ships as a
>    **deliberately RED** gene rather than being hidden by omitting the field (**D16**). The gate (≥95/=100)
>    is still **MET**.
>
> The `22/22` in the `clerkrun` row was, at the time, a **hollow 100%** — `GET /v1/users/{id}` had no gene at
> all, which is precisely the bug iter-04 was fixing. The score did not move across that fix **because
> nothing was measuring the thing that broke**. That is the lesson, and it is why this table is annotated
> instead of deleted.

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

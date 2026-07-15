---
iter: 4
milestone: M218
iteration_type: tik
status: closed-fixed
opened: 2026-07-13
---

# iter-04 (tik) — kill the ~6.10 s retry ladder

**Active strategy:** **TOK-01** — *"reachability-first"*, step 3: *"then re-measure and let the next bottleneck
name itself. **Expect a new bottleneck to appear** — the fix does not end the milestone, it unblocks the
measurement of the next layer."* It appeared. This iter kills it.

## Step 0 — re-survey

Re-measured at iter-03 close on a **cold, green** stack: employee p95 **7.90 s**, manager p95 **7.00 s** (from
39.45 / 38.30). The gate is **< 5 s**, so the remaining gap is **~2.9 s** — and it is **entirely** one term:

```
[slow-body] HTTP 200 headers in 62 ms but BODY took 6104 ms (streamed/blocked SSR)
```

**6,104 / 6,107 ms — reproducible to ±3 ms across both vantages.** Target confirmed and unchanged.

## Cluster / target identified

**The `~6.10 s` is a retry ladder, not a timeout — and it is Clerkenstein lying to next-web about who the hero is.**

The arithmetic named the shape first: a blackhole costs **10.5 s/attempt**, so 6.1 s cannot be one. But
`prefetchUserStatus` configures exactly:

```ts
retry: 2,
retryDelay: (attemptIndex) => Math.min(2000 * Math.pow(2, attemptIndex), 20000),   // 2 s → 4 s
```
⇒ **3 attempts × ~33 ms + (2 s + 4 s) ≈ 6.0 s.** A **fast-failing** fetch, retried twice.

Caught in the act — the Cosmo router logged the SSR calls (`user_agent: "node"`) at **ERROR**, HTTP 200,
latency 2.4 ms, at **20:12:16 → :18 → :22** (**+2 s, +4 s** — the ladder). And `app`, at those same instants:

```
ERROR graphql resolver error  user:{ID:40921b2e-…, Email:dan.rossi3@cervato-systems.com}
  error="input:1:21: userPreferences user does not have permission to set studio preferences"
```

The resolver is not an authorization check at all — it is an **identity equality check**:

```go
func (r *queryResolver) UserPreferences(ctx context.Context, userID uuid.UUID) (string, error) {
	currentUser := authn.UserFromContext(ctx)          // ← from the JWT (the FAPI minted it: the REAL hero)
	if currentUser.ID() != userID {                    // ← userID = next-web's user.externalId (the BAPI)
		return "", fmt.Errorf("user does not have permission to set studio preferences")
```

**The two halves of Clerkenstein's identity disagreed.** Proven against the live fake BAPI:

| source | `external_id` | email |
|---|---|---|
| **FAPI** — the client + the JWT `app` authenticates (roster-correct) | `40921b2e-4b27-524e-ace5-c4699156bad9` | `dan.rossi3@cervato-systems.com` |
| **BAPI** — next-web's **server-side** `currentUser()` | **`11111111-1111-1111-1111-111111111111`** | **`demo@anthropos.test`** |

`clerk-backend/resources.go` says so in its own comment: **`// Disarmed: any id → the demo user.`** That was
correct when a demo had **one** user. Since the **Stories & Heroes** model (M35) every hero has their **own**
eid — and the BAPI kept handing out the stub for **everyone**.

## Hypothesis

Make the fake BAPI's `getUser` **roster-aware** — report the requested hero's **real** eid/email/name — and the
`userPreferences` refusal disappears, taking the whole 6 s retry ladder with it.

## Expected lift

**p95 7.90 s → ~1.8 s** (employee) and **7.00 s → ~0.9 s** (manager) — i.e. **under the 5 s gate on both
vantages**, if the 6.1 s term is the only thing left. If a *new* residual appears, TOK-01's "assume at least one
more bottleneck" holds once more and iter-05 takes it.

## Phase plan

1. Add a per-hero **identity** seed to the BAPI store (it already mounts the roster — it just used it only for
   **memberships**, never the user resource).
2. Make `getUser` consult it; **fall back to the historical stub when no roster** — that fallback is what keeps
   the alignment gate unmoved.
3. **`/align-run` across all 5 surfaces** (the milestone's BLOCKING guard) — prove nothing moved.
4. Cold reset-to-seed on `billion`, re-measure both vantages.

## Escalation conditions

- **Any** alignment movement (critical < 100% or overall < 95%) → stop and re-shape the fix.
- If the fix needs a platform-repo edit → escalate (it does not: the defect is entirely in rext's own mock).

## Acceptable close-no-lift outcomes

If the refusal disappears but p95 does **not** fall, the attribution model is wrong — a first-class falsification.

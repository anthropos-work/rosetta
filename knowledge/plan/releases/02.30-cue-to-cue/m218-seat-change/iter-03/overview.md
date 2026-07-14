---
iter: 3
milestone: M218
iteration_type: tik
status: closed-fixed
opened: 2026-07-13
---

# iter-03 (tik) — land the real C-1 fix, and re-measure on a stack that is actually wired

**Active strategy:** **TOK-01** — *"reachability-first: fix the address, not the variable."* Step 2 (make the
baked public URL reachable from inside the container) + step 3 (re-measure; let the next bottleneck name itself).

> **Run history.** This iter was opened in **run 1**, which authored + landed the fix and then **stalled**
> (harness watchdog) at the moment of measuring. **Run 2** recovered it. The sections below are the merged plan;
> everything run 1 established is preserved, and the **re-survey** records what run 2 found when it went to take
> the number.

## Step 0 — re-survey (mandatory before targeting)

TOK-01's next-tik direction named `FIX-M218-iter03-ssr-origin`. Re-surveyed against the live stack:

**The target is unchanged, and the fix is already landed** — rext `9a8b6d5`: the sha-pinned demo-patch
`next-web-ssr-graphql-origin` (prepend a **server-only** `WUNDERGRAPH_SSR_ENDPOINT` to the `||` chain in
`packages/graphql/src/server/server.graphql.ts`) plus the matching runtime env in `gen_injected_override.py`.
Verified live on `billion`: the demo clone carries the patched chain, the compiled server chunk contains the new
var, and the container's runtime env supplies `http://graphql:8080/graphql`.

**What run 2 found when it went to measure — the recovered run says two things at once.**
Run 1's post-fix measurement *did* execute (recovered from the rext working tree, `capturedAt 19:44Z`):

| | iter-02 baseline | run-1 post-fix |
|---|---|---|
| SSR document **body** | **37,533 ms** | **6,107 ms** |
| reached ACCESS | **3 / 3** | **0 / 1** — bounced to `/login` |

So **the fix works on its own leg** (a 6.2× collapse of the dominant term, exactly as TOK-01 predicted) — but the
**stack it was measured on is no longer wired**, so the number is **not gradeable**.

## Cluster / target identified

**The stack's next-web image is Clerkenstein-DEWIRED, and a latent rext defect is what let it stay that way.**

Established, mechanism first:

- The minted publishable key reaches next-web **only** through the gitignored `apps/web/.env.local` overlay that
  `build_frontend_next_web()` writes (`up-injected.sh:357`) and trap-removes after the build.
- Run 1 rebuilt the image **out-of-band** — it passed the three `--build-arg` offset URLs but never wrote the
  overlay. (Proof: `build-next-web.log` is still dated **Jul 11**; the image is dated **Jul 13 15:12**.)
- `NEXT_PUBLIC_*` is **build-inlined** ⇒ next-web baked the **repo-resident real-Clerk pk** into the browser
  bundle: `pk_test_b3JpZW50ZWQtbGFiLTMzLmNsZXJrLmFjY291bnRzLmRldiQ` = `oriented-lab-33.clerk.accounts.dev$`.
- ⇒ the browser's clerk-js talks to the **real Clerk app**, not the fake FAPI ⇒ no session ⇒ the SSR GraphQL call
  errors ⇒ the `retry: 2` / 2 s + 4 s ladder fires (**the 6.1 s residual**) ⇒ redirect to `/login`.

**F-6 — the defect behind the defect.** The next-web image **cache-validator** (`up-injected.sh:344-353`, M211)
guards exactly one baked constant: the `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` **image ENV**. The publishable key is
**not an image ENV** — it is inlined into the bundle from the overlay — so `docker image inspect` **structurally
cannot see it**. An image with the **right offset** and the **wrong, real-Clerk key** therefore *passes* the
validator and is **silently reused** on the next bring-up. That is:

1. a **correctness** hazard — login is simply broken (this is what bit run 1), and
2. a **safety** hazard — a demo phoning **production auth**, which `safety.md` forbids outright,

and it would have **silently corrupted this milestone's own 5-cold-run gate battery**. It is the **same defect
class as C-1 itself**: a build-inlined `NEXT_PUBLIC_*` constant that a rebuild path can fail to supply.

**Note the compounding failure:** the stack's `autoverify.json` still said `{"green":true}` — written at 09:49, i.e.
**before** the 15:12 out-of-band image swap. A stale green verdict is exactly the M217 hazard, one level up.

## Hypothesis

1. Closing **F-6** makes the sanctioned bring-up **self-heal** the dewired image: the validator sees the wrong key
   → forces the rebuild through `build_frontend_next_web()` → the overlay lands → Clerkenstein is re-wired.
2. On that correctly-built stack, the SSR-origin fix removes the ~37.5 s blocking term and click→ACCESS collapses
   from **p95 39.45 s / 38.30 s** toward the low seconds.

## Expected lift

The 37.6 s term → ~0. Whether that **alone** lands p95 under **5 s** is explicitly unknown — TOK-01 says *"assume
at least one more bottleneck."* The 6.1 s retry-ladder residual observed in the broken run should **vanish** once
the session is valid (it fires only because the query errors); if it survives, it is the next target.

## Phase plan

1. **F-6** — harden the next-web cache-validator to also gate on the **minted pk**, read from the **bundle**
   (rext-only, fail-safe toward rebuild).
2. **Rebuild through the sanctioned path** on `billion` at the new rext tag; confirm the validator itself forces
   the rebuild and the minted pk lands.
3. **Re-establish green** — a **fresh** `autoverify.json` (the 09:49 one is stale w.r.t. the image swap).
4. **Re-measure** both vantages with the iter-02 harness; report per-leg attribution.

## Escalation conditions

- If the **employee** vantage is not explained by the SSR fix once the stack is correctly wired → the milestone's
  **RE-SCOPE TRIGGER** fires: stop and re-measure; do **not** push a manager-only fix set.
- If the rebuild cannot be driven through the sanctioned path without a platform-repo edit → escalate.

## Acceptable close-no-lift outcomes

A re-measurement that **refutes** the 6.2× collapse would itself be a first-class finding (and would mean the
attribution model is wrong) — that would close no-lift with documented falsification.

## Out (routed forward, not dropped)

- **The twin `NEXT_PUBLIC_BACKEND_API_URL`** (run-1 discovery). Run 2 checked its blast radius: every reader is
  **client-side** (hooks, UI components, form `action=` targets), so it is **not** on the SSR blocking path and
  **not** in this iter's rebuild. Stays routed → **iter-04** (`PROBE-M218-backend-api-url-twin`).
- **C-3's re-probe** (`PROBE-M218-c3-rerun`) — the federation only becomes exercisable once a login completes.

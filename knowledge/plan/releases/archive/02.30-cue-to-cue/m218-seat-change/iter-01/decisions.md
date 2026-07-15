# iter-01 — Decisions

## D1 — The root cause, and the no-op the plan would have shipped

**Evidence:** the compiled SSR bundle (`/app/apps/*/.next/server/`) contains the **build-inlined literal**
`https://billion.taildc510.ts.net:15050/graphql` (9 occurrences). It does **not** read `process.env` at runtime.
From inside the container that URL **connect-times-out after 10,564 ms** (`UND_ERR_CONNECT_TIMEOUT` — undici's 10 s
default); DNS resolves fine (`100.110.136.3`), the **TCP connect from the docker bridge to the tailscale IP
blackholes**. `http://graphql:8080/graphql` from the same container → **HTTP 200 in 94 ms**.

**Decision:** C-1 is the root cause, **with its mechanism corrected**. The milestone's stated C-1 fix — *"add
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT: http://graphql:8080/graphql` to the next-web **runtime** env in
`gen_injected_override.py`"* — is a **NO-OP** and must not be shipped. SSR never reads that variable.

**Why this matters:** it was the milestone's #1 ranked fix, described as "one line, rext-only." It would have
landed, measured **zero**, and burned an iter — and it would have looked like a refutation of C-1 when C-1 is in
fact correct. `frontend-tier.md:240` said so all along; the Phase-0b audit surfaced the contradiction; the
strengthened probe settled it.

## D2 — Fix shape: make the baked URL REACHABLE; do not try to re-point it

**The constraint.** `NEXT_PUBLIC_*` is a **single build-time constant** with **two consumers that need different
reachability**:
- the **browser** needs the **public origin** (`https://billion.taildc510.ts.net:15050/graphql`) — and it has it,
  correctly baked;
- **SSR** needs a **container-reachable** origin (`http://graphql:8080/graphql`).

One constant cannot be both. Re-baking it for SSR **breaks the browser**. Re-pointing it at runtime is D1's no-op.

**Decision:** do not fight the constant — **make the address it already holds resolve to something reachable from
inside the container.** The candidate (to be validated in the fix tik) is a compose-level `extra_hosts` mapping on
the next-web service so `billion.taildc510.ts.net` resolves to the docker host gateway, letting the container's
connect to `:15050` land on the host's published port instead of blackholing at the tailnet IP. This is **pure rext
injection** (`stack-injection`), **zero platform-repo edits**, and it leaves the browser bundle untouched.

Alternatives, ranked, if that fails validation:
1. a `hostAliases`/gateway route on the demo network (same class, same file);
2. a **sha-pinned demo-patch** to next-web that gives SSR its own origin (the sanctioned hatch —
   `demopatch-spec.md`) — heavier, and only if the reachability route is genuinely closed;
3. **escalate.** Never edit platform source.

**Not chosen:** changing the build arg (breaks the browser); `NEXT_PUBLIC_*` renaming (platform edit).

## D3 — `autoverify.json` path drift (doc vs reality)

`corpus/ops/verification.md:207` documents the file as **`<stack>/autoverify.json`**. It is actually written to
**`rosetta-extensions/demo-stack/stacks/<project>/autoverify.json`** (`stack-verify/live/autoverify.sh:250` writes
`"$STACK_DIR/autoverify.json"`, and `STACK_DIR` is the per-project stacks dir, not the stack root).

**Why this is not cosmetic:** M217 shipped this file *specifically* so **M218 could gate its measurements on it**
(`verification.md:213` says exactly that). A gate whose documented path does not exist is a gate that **silently
never fires** — the exact failure class M217 existed to kill. Fix the doc (and have the harness read the real path).

## D4 — The harness must be a real browser; curl is structurally blocked

Two independent walls, both measured:
- the fake-FAPI **validates `redirect_url`** against the public origin → **HTTP 400** for a loopback redirect;
- next-web's middleware **307s** any non-https origin (so a `Host:`-header rewrite onto `http://localhost:13000`
  cannot reach the authenticated render either).

⇒ the login flow **cannot be driven by curl**, at all. The harness must drive the **real https origin in a real
browser** — which is also the only way to see the client-gate (clerk-js) leg. This *independently* confirms the
milestone's Playwright mandate, and it is a constraint documented **nowhere**. It goes in `latency-budget.md`.

**Corollary for the harness build:** `stack-verify/e2e/lib/cockpit-login.ts::loginAs` (`:64`) navigates with
`waitUntil: 'networkidle'`, which the milestone **explicitly bans** (never-idle long-polls). The instruction to
"reuse it, do not fork it" and the `networkidle` ban **collide**. Resolution (for iter-02): reuse `selectSeat` /
`loginAs` for the **handshake**, but the latency harness must **not** inherit the `networkidle` wait — it polls for
**content presence** instead. Extend the shared lib with a wait-strategy parameter rather than forking it.

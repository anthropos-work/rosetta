# M218 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | The 4-leg experiment, zero code. **C-1 confirmed as root cause but its MECHANISM was wrong in the plan — SSR reads the BUILD-INLINED public URL, which connect-times-out (10.56 s) from inside the container; the planned runtime-env one-liner is a NO-OP.** Budget now sums to the symptom (~112 s). C-2 refuted (M217 fixed it — field proof). C-5 refuted as cause, retained as a free win. C-3 parked (unexercisable until a login completes). TOK-01 authored. | baseline **UNMEASURED** — no instrument exists (iter-02 builds it) | closed-fixed — see `iter-01/progress.md` |

## Next iter

**iter-02 (tik) — the latency harness.** Under **TOK-01** (`decisions.md`).
Build the Playwright per-leg latency harness as a **new `stack-verify` surface**; run it against the **current,
unfixed** `billion` demo; record the **honest pre-fix p95 baseline** for both vantages. **Write no fix.**
Binding constraints in `spec-notes.md` § Harness contract. Handler `HARNESS-M218-iter02-latency-probe`.

## Carry-forward queue

| handler | item |
|---|---|
| `FIX-M218-iter03-ssr-origin` | The **real** C-1 fix — make the baked public URL reachable from inside the container (iter-01 **D2**); **not** the planned runtime-env no-op (**D1**). rext `stack-injection` only. |
| `PROBE-M218-c3-rerun` | Re-run leg 2 (C-3, router retries) **after** the fix — the federation is only exercisable once a login completes. |
| `FIX-M218-c5-clerkjs` | Vendor clerk-js + **bound the unbounded timeout** (`server.go:187`, `Timeout: 0`). Alignment-invisible ⇒ gate-free. |
| `DOC-M218-audit-corrections` | M43-D5 correction (**`cockpit-spec.md:58,187`** — 155 was wrong) · CI-inert (**`alignment_testing.md:232,233,239`**) · `clerkenstein.md:3-4` header · **preserve** `frontend-tier.md:240` (it was right). |
| `DOC-M218-f2-autoverify-path` | **F-2:** `verification.md:207` documents the wrong path for `autoverify.json` — the file M218 gates on. |
| `DOC-M218-f3-harness-constraint` | **F-3:** curl structurally cannot drive the login flow (redirect_url validation + https 307). Undocumented. |
| `PROBE-M218-f4-hosting-url` | **F-4:** `NEXT_PUBLIC_HOSTING_URL=http://localhost:3000` — same un-offset-loopback class as C-1. On the login path? |
| _(also-in-scope, from overview)_ | Clerk telemetry off · ant-academy real-Clerk-secret leak · `x/crypto@v0.52.0` |

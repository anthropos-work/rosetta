---
milestone: M257x
iter: 10
iteration_type: tik
status: closed-fixed
opened: 2026-07-31
closed: 2026-07-31
---

# iter-10 — `FIX-M257x-academy-not-serving` (the fix, under iter-09's binding constraint)

**Active strategy reference:** `TOK-01` step 5 — *prove it cold*. The academy is the last genuine `✗` in
autoverify on a stack that is otherwise UP, and **gate clause 1 needs three consecutive green cold cycles**.

## Step 0 — re-survey (mandatory, and it changed the target's shape)

iter-09 handed off a **mechanism claim** and a **pre-computed fix menu**. Standing pattern says treat an
inherited hand-off as evidence to re-measure, not as fact. Re-measured on demo-1 (up 5 h, 16 containers),
academy relaunched by hand under controlled binds:

| bind | `GET /` | `x-middleware-rewrite` |
|---|---|---|
| `-H 127.0.0.1` | **500 in 30.077 s** | `http://localhost:13077/` (absolute) |
| `-H 0.0.0.0` | 200 in 0.686 s | (relativized — request served) |
| `-H ::1` | **500 in 30.1 s** | `http://localhost:13077/` (absolute) |
| `-H 127.0.0.1` + `NODE_OPTIONS=--dns-result-order=ipv4first` | **500 in 30.123 s** | `http://localhost:13077/` |

**iter-09's stated mechanism is REFUTED.** It attributed the failure to `localhost` resolving to `::1`
before `127.0.0.1`, so that an IPv4-only listener could not be self-dialled. Three measurements kill that:

1. `curl http://127.0.0.1:13077/` — a dial that never touches the name `localhost` — fails **identically**
   (500 / 30.06 s). A client-side resolution story cannot explain that.
2. `--dns-result-order=ipv4first`, the direct repair for a resolution-order bug, changes **nothing**.
3. `-H ::1` — which makes `localhost`'s first-resolved address the one being listened on — **also fails**.

## Cluster / target identified

The real mechanism is an **origin-STRING comparison**, and it is two lines of Next 16 that disagree:

- `next/dist/server/web/next-url.js:15-20` — `REGEX_LOCALHOST_HOSTNAME` normalizes **every** loopback
  hostname (`127.x.x.x`, `[::1]`, `localhost`) to the literal `"localhost"`. So the middleware's
  `request.nextUrl` origin is **always** `http://localhost:$PORT`, whatever the bind, whatever the `Host`
  header. Clerk's `dev-browser-missing` path rewrites to that URL.
- `next/dist/server/lib/router-utils/resolve-routes.js:117` — the router's `initUrl` is built from the
  **raw** `opts.hostname` (i.e. the literal string passed to `-H`), with **no** such normalization.
- `next/dist/shared/lib/router/utils/relativize-url.js` — `getRelativeURL` keeps the rewrite absolute
  unless `relative.origin === baseURL.origin`, a **plain string equality**.

So `-H 127.0.0.1` ⇒ `http://localhost:13077` ≠ `http://127.0.0.1:13077` ⇒ the in-app rewrite is
mis-classified as an **external** proxy target ⇒ Next proxies the server **to itself**
(`router-server.js:377` → `proxy-request.js`), which recurses until the **30 000 ms `proxyTimeout`
default** — that is the flat 30.0 s, and it is a constant in Next's source, not a coincidence.

Evidence, captured live: `x-middleware-rewrite: http://localhost:13077/` on the 500, and
`x-middleware-rewrite: /` on the 200 under the candidate fix.

## Hypothesis

Passing the **literal string `localhost`** to `-H` makes `initUrl`'s origin equal the string Next
normalizes every loopback hostname to, so the rewrite relativizes — while `localhost` still resolves only
to a **loopback** address on every host family, so **M221's de-exposure is preserved, not traded away**.

The elegance that matters: the comparison Next performs is on the string we supply, so the fix is
**resolver-independent by construction**. Only *which* loopback address gets bound is OS-dependent, and
both loopback addresses are equally un-exposed.

## Expected lift

`autoverify` FAILED 1 → 0 for the academy check; the academy answers 200 with real body content.

## Phase plan (2 planned lines — the second is the protocol's own paired-check rule)

1. **The bind fix** in `demo-stack/ant-academy.sh`, re-proved against `stack-injection/exposure_claim_guard.py`
   (not merely against the readiness probe), plus a regression fence over the bind literal.
2. **The paired "does it still work?" check** — iter-09's decision 3. The readiness probe polls `/` with
   `curl -fsS --max-time 3`; the failure it is watching for takes **30 s**, so every attempt times out and
   the probe reports *"never answered"* for a server that was answering (with a 500). A probe whose
   per-attempt timeout is shorter than the failure mode cannot describe what it saw. `platform-alignment.md`
   §5's family, on the launcher's own reporting path.

`FIX-M257x-autoverify-evidence-log-path` is a **third** line and is deliberately NOT in this iter's scope —
routed to iter-11 (it is on clause 1's critical path but is independent of this fix).

## Two-OS honesty constraint (binding, from the orchestrator + iter-09)

This box is **arm64 macOS**; the platform's real target is Linux. The origin-equality half is a pure string
operation in the same JS bundle on any OS and is argued from cited `file:line`. The **bind-address** half is
genuinely OS-dependent and will be measured on **linux/arm64** separately. Any claim is scoped to what was
actually measured on each family — no generalisation.

## Escalation conditions

If the only route to a green check weakens M221's de-exposure — i.e. if the exposure guard goes RED, or the
academy ends up on a routable interface — **STOP and surface it**; close honestly rather than land it.

## Acceptable close-no-lift outcomes

A falsification that the loopback bind can be kept at all (e.g. every loopback literal fails the origin
match) would be a complete iter: it would convert the fix into a user decision about the exposure/usability
trade, which is not mine to take.

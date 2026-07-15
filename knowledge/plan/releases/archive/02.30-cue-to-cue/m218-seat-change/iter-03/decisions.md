# iter-03 — Decisions

## D7 — The SSR origin gets a **server-only** env var, supplied at RUNTIME, via a sha-pinned demo-patch

**The wall.** next-web resolves its **server-side** GraphQL origin from `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`. That
is **one build-time constant serving two consumers with incompatible reachability needs**: the **browser** needs
the *public* origin; the **SSR pass** needs a *container* origin. `NEXT_PUBLIC_*` is **build-inlined**, so the one
constant cannot differ between them.

**Why every cheaper option is exhausted** (`demopatch-spec.md` § 1 mandates trying them in order):

| option | verdict |
|---|---|
| set `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` in the **runtime** env | **NO-OP** — the SSR bundle reads the inlined literal, never `process.env`. *This was the milestone's own planned "one-line fix"* (iter-01 **D1**): it would have shipped, measured **zero**, and looked like a refutation of a correct root cause. |
| re-bake the **build arg** to a container origin | **breaks the browser** — same single constant. |
| docker **network-alias** the public host | yields a *fast* `ECONNREFUSED`, not a success (~6.2 s — still over gate). |
| a TLS-terminating **reverse-proxy sidecar** aliased as the public host | works — but adds a **new container + a new third-party image to every `--public-host` demo**, purely to work around a missing platform env var. Strictly **more** invasive than one line, and it would **hide** the platform finding instead of disclosing it. |

**Decision.** Author the sha-pinned demo-patch `next-web-ssr-graphql-origin` — **prepend** a server-only
`WUNDERGRAPH_SSR_ENDPOINT` to the existing `||` chain, and supply its value from the container's **runtime** env
(`gen_injected_override.py` → `http://graphql:8080/graphql`). `WUNDERGRAPH_SSR_ENDPOINT` is **deliberately not a
`NEXT_PUBLIC_*` name** — that is the entire point: non-public env vars in server code **are** real runtime
`process.env` reads in Next.js, so the value actually lands. Hence **no `build_env`** in the manifest: baking it
would defeat the fix. The change is **behaviour-identical when the var is unset** (a strict prepend to a `||`
chain), so it is safe to upstream as-is.

**Disclosed (the real platform fix).** next-web should read its server-side GraphQL origin from a **server-only**
env var. Any containerized deployment behind a proxy/ingress whose public origin is not reachable from inside its
own container hits this. The demo only makes it **visible** — and expensive, because the unreachable address
**blackholes** rather than refusing.

## D8 — Validate baked constants by reading the **bundle**, not the image ENV (F-6)

**What happened.** An out-of-band next-web rebuild passed the three `--build-arg` offset URLs but never wrote the
`apps/web/.env.local` overlay — the **only** carrier of the minted publishable key. `NEXT_PUBLIC_*` being
build-inlined, the bundle silently baked the **repo-resident real-Clerk pk**
(`…b3JpZW50ZWQtbGFiLTMz…` = `oriented-lab-33.clerk.accounts.dev$`). The demo's browser clerk-js then talked to
the **real Clerk app**: no session → `/login` loop. Login was **broken**, and the stack was **phoning production
auth** — which `safety.md` forbids outright.

**Why the tooling did not catch it.** The M211 image cache-validator compares the baked
`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` **image ENV**. The pk is **not an image ENV** — it is *inlined into the
bundle* — so `docker image inspect` **structurally cannot see it**. An image with the **right offset** and the
**wrong key** therefore *passed* the validator and would have been **silently reused** on the next bring-up,
corrupting this milestone's own 5-cold-run gate battery.

**Decision.** The validator also asserts the stack's minted pk is **present in the built bundle**:

```sh
docker run --rm --entrypoint sh "$img" -c "grep -rqs -- '$PK_DEMO' /app/apps/web/.next/static"
```

All four overlay-borne constants (pk, `STUDIO_URL`, `ACADEMY_URL`, `PUBLIC_WEBSITE_URL`) come from that single
file, so **one probe detects the whole "overlay was not applied" class**; host/offset/scheme drift stays covered
by the existing `want_ep` check. **Fail-safe by design:** anything that cannot be positively verified (missing
path, no shell, unreadable bundle) **rebuilds** — a needless ~3 min build is strictly cheaper than serving a
real-Clerk-wired demo to a customer.

**Generalization.** This is the **same defect class as C-1**: a build-inlined `NEXT_PUBLIC_*` constant that a
rebuild path can silently fail to supply. *A cache-validator that reads image ENV can only see the constants that
happen to be ENV.* **Validate baked constants by reading the artifact they were baked into.**

## D9 — A green verdict is only as fresh as the image it graded

`autoverify.json` said `{"green":true}` — written **09:49**, i.e. **before** the **15:12** out-of-band image swap.
M217 shipped that file precisely so M218 would never measure a broken stack; a **stale** green is the same hazard
one level up. Every M218 measurement therefore re-establishes green on the **actual image under test** (a fresh
bring-up writes a fresh verdict), rather than trusting the file's mere existence. **The green gate grades a
stack-at-an-instant, not a stack-forever.**

## D10 — The `NEXT_PUBLIC_BACKEND_API_URL` "twin" is **not** a twin on the login path

Run 1 flagged it as a twin of C-1 and routed it to iter-04. Run 2 checked the blast radius before spending a
build cycle on it: **every reader is client-side** (`useManageSubscription.tsx`, `core-js/workforce/api.ts`,
`core-js/talkToData/api.ts`, UI components, form `action=` targets). The **browser** genuinely wants the *public*
origin there, and the baked value is *correct* for the browser — so there is **no SSR blocking read**, and it is
**not on the login path**. It stays routed to **iter-04** for confirmation, but it is **not** a gate-blocker and
was correctly **excluded** from this iter's rebuild.

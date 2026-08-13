# iter-16 — decisions

## D79 — the mock listened on a path the real client does not use, and its test drove the same unused path

`handleSignOut` is correct code. It was registered on `POST /v1/client/sessions/{id}/remove`. Measured from a
real browser, **clerk-js 5.127.1 signs out with `POST /v1/client/sessions?_method=DELETE`** — the collection,
with a method override. That pattern was not on the mux, there is no catch-all, and nothing in the server read
`_method` at all. So the request 404'd, the handler never ran, and `/v1/me` returned **200 through an entire
logout**.

The part that makes this a first-class finding rather than a typo: **`TestServer_signOutClearsSession` drove
`/v1/client/sessions/sess_clerkenstein/remove` and had passed since the mock was written.** A green test about
the mock talking to itself, on the single behaviour a user had complained about out loud. That is the 18th
could-not-fail check this milestone has found, and the first inside Clerkenstein — where the class is
*structurally* easiest to commit, because both ends of the conversation are ours.

**The generalisable remedy:** a mock's test must drive **the exact request the real client sends** — verb,
path, and query string, copied from a trace — not the shape the mock's own route table implies.

## D80 — `_method` is a dispatch key, not a hint

The naive fix is to register `POST /v1/client/sessions` and call it sign-out. clerk-js POSTs that collection
for other reasons, so that would turn one bug into its mirror image. Only `_method=DELETE` (case-insensitive)
signs out; anything else returns the client snapshot. Pinned by its own test over `""`, `?_method=PATCH` and
`?_method=GET`, because a rule with no test is a comment.

## D81 — an explicit sign-out should be sticky until an explicit login

Fixing the route is **necessary and not sufficient**, and the same trace proves it: clerk-js clears `__session`
but leaves `__client_uat`, so next-web's middleware bounces to the handshake
(`__clerk_hs_reason=client-uat-but-no-session-token`) and `handleHandshake` calls `establishLocked()`
**unconditionally**. The session comes straight back.

The rule, stated so the next iter implements it rather than re-derives it: **a handshake carrying
`__clerk_identity` is an intentional login and must always establish** (the cockpit's path, and every
Playthrough's); a **bare** handshake — the middleware's automatic re-entry — **must not resurrect a session the
user explicitly ended**.

Not shipped this iter, deliberately. That handler is on every login path in the suite; the change needs a live
re-proof; and the live re-proof needs the stack's clone re-pinned and `fake-fapi` rebuilt. iter-07 D31 already
recorded the governing rule — *an unverified lifecycle fix is the worst outcome available*. (a) and (b) get one
rebuild and one proof, together.

## D82 — say what the evidence cannot cover

`demo-2`'s `fake-fapi` container runs the previously-pinned binary, so the Playwright suite cannot see this
change at all. Running it would have produced `172 passed` and zero information about the fix. Reporting that
green as verification would be the favourable-sample dishonesty D-v28-13 was written about, in a different
currency. The available evidence is stated for what it is: module tests RED-before/GREEN-after on the real
request shape, plus a browser trace of the defect.

Also checked rather than assumed: **no gene in `alignment/` references sign-out or the session collection**, so
the mock's measured fidelity score is untouched and the DNA escalation condition in this iter's `overview.md`
does not fire.

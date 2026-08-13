**Type:** tik · shape: standard (single target: D-v28-5)

# iter-16 — D-v28-5 root-caused: the mock listened for sign-out on a path the client never uses

## Phase A — measure the flow (and read the surface first)

**Read before probing**, which immediately narrowed it: the cockpit has **no logout affordance at all** (no
logout / sign-out / session-remove path anywhere in `cockpit.py` — it is a launcher). So "logout to cockpit"
is the **app's** sign-out, after which the presenter returns to the cockpit for another hero. The app's
sign-out is `apps/web/src/app/(unauthenticated)/logout/[[...logout]]/page.tsx`: `useClerk().signOut()` →
clear storage → `router.replace('/login')`.

**Then measured, in a real browser on `demo-2`, step by step.**

| step | observed |
|---|---|
| 1. login as `pt-employee` via the cockpit's handshake | lands `/profile`; cookies `__session`(878) `__client_uat`(10) `__clerk_db_jwt`(503); `/v1/me` → **200** |
| 2. drive `/logout` | lands **`/home`**, not `/login`; cookies **all still present** (plus suffixed twins); `/v1/me` → **200** |
| 3. visit `/profile` again | **STILL IN THE APP**, `bodyLen 2147`, `"Pat Ellis"` = 1 — the logged-out hero's own profile |
| 4. ONE login as a different hero | `"Morgan Reyes"` = 0 **and** `"Pat Ellis"` = 0 — a mixed/broken identity on the first click |

**The defect is reproduced.** Step 3 is the user's report in one line: after logging out, the app still serves
the logged-out hero.

### The request trace named the cause

The FAPI calls during `/logout` include, in order:

```
POST /v1/client/sessions?__clerk_api_version=2025-11-10&_clerk_js_version=5.127.1&_method=DELETE
GET  /sign-in?redirect_url=…
GET  /v1/client/handshake?…&__clerk_hs_reason=client-uat-but-no-session-token
```

**clerk-js 5.127.1 signs out by POSTing the sessions COLLECTION with a `_method=DELETE` override** — the same
method-override convention it already uses for `/v1/environment?_method=PATCH`. The mock registered sign-out
on `POST /v1/client/sessions/{id}/remove`, **that collection pattern was not on the mux at all**, there is no
catch-all route, and **nothing in the server read `_method`** anywhere. So the request 404'd and
`handleSignOut` — which is correct code — **never ran**. Server state was never cleared: `/v1/me` stayed 200
through the entire logout.

**And the pre-existing unit test drove the same unused path.** `TestServer_signOutClearsSession` has POSTed
`/v1/client/sessions/sess_clerkenstein/remove` since the mock was written and has passed ever since — a green
test about the mock talking to itself, on the one behaviour a presenter complains about. That is this
milestone's signature defect, found for the 18th time, this time in the Clerk mock's own suite.

*(The `curl` route-probe route was closed off by the documented macOS LibreSSL-vs-mkcert issue —
`FIX-M256-autoverify-fapi-libressl`, still open in the routing table — so the browser trace and a Go test were
the available instruments. Both were sufficient.)*

## Phase B — the fix, failing-test-first

1. **A failing test written before any fix**, driving the exact request shape measured in the browser
   (query string and all) → **404, RED**, naming the defect.
2. `handleSessionsCollection` added: `POST /v1/client/sessions` dispatches on `_method`, and **only**
   `DELETE` (case-insensitive) signs out. Also `DELETE /v1/client/sessions` for a client that ever sends the
   honest verb. The legacy `/{id}/remove` route is **kept** — older clerk-js uses it, and Go 1.22's mux
   prefers the more specific `/{id}/…` patterns so nothing is shadowed.
3. **A second test pins the other direction:** `_method` is a *dispatch key*, so a bare `POST
   /v1/client/sessions` (or `?_method=PATCH`, or `?_method=GET`) must **not** sign the user out. Without it,
   "register the path" degrades into "any POST here logs you out" — the same bug facing the other way, and
   clerk-js does POST that collection for other reasons.
4. Both new tests **RED before, GREEN after**; the legacy test still passes, so the fix is additive.

## What this fix does NOT close — measured, not assumed

**Part (a) is necessary and not sufficient, and the same trace says so.** After `/logout`, clerk-js clears the
browser's `__session` but leaves `__client_uat`, so next-web's middleware bounces to the handshake with
`__clerk_hs_reason=client-uat-but-no-session-token` — and `handleHandshake` calls `establishLocked()`
**unconditionally**. With (a) in place the server *will* be signed out at that moment, and the handshake will
still sign it straight back in.

So the honest reading is: **D-v28-5 is root-caused and half-fixed.** The remaining half is a session-lifecycle
rule, and it is specified rather than guessed at:

> An explicit sign-out should be **sticky until an explicit login**. A handshake carrying `__clerk_identity`
> is an intentional login and must always establish (that is the cockpit's and every Playthrough's path). A
> **bare** handshake — no identity, i.e. the middleware's automatic re-entry — must not resurrect a session
> the user explicitly ended.

**Why it is not being shipped in this iter, deliberately.** That handler sits on *every* login path in the
suite, the change needs a live re-proof, and the live re-proof needs the stack's `rosetta-extensions` clone
re-pinned and the `fake-fapi` container rebuilt (the running container is the **old** binary, so a suite run
right now would be green about code that is not deployed — which is worth saying out loud rather than
banking). iter-07 D31 already recorded the rule this obeys: *an unverified lifecycle fix is the worst outcome
available*. (a) and (b) should be proven together, once, on one rebuild.

**The Alignment-DNA escalation did not fire:** no gene in `alignment/` references sign-out or the session
collection, so the mock's measured fidelity is untouched by this change (checked, not assumed).

## Phase D — verification available without a rebuild

- `clerkenstein`: **all packages `ok`, 0 FAIL** — including the two net-new tests and the legacy one.
- `alignment`, `playthroughs`: rc 0, 0 FAIL. `gofmt -l`: clean. `go build ./...`: clean.
- The Playwright suite is **deliberately not re-run as evidence for this change**: `demo-2`'s `fake-fapi`
  container is the previously-pinned binary, so the suite cannot see the fix either way. It was last green at
  `172 passed` ×3 in iter-15 and no harness file changed in this iter.

## Close — 2026-07-29

**Outcome:** **D-v28-5 root-caused and half-fixed.** The mock listened for sign-out on
`/v1/client/sessions/{id}/remove` while clerk-js 5.127.1 sends `POST /v1/client/sessions?_method=DELETE` — a
pattern that was not on the mux, with no catch-all and no `_method` handling anywhere — so `handleSignOut` was
**dead code** and `/v1/me` stayed 200 through an entire logout. Fixed additively, failing-test-first, with a
second test pinning that only `DELETE` dispatches. The remaining half (the handshake re-establishing
unconditionally) is measured, specified, and routed for a joint live proof.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET — clause 2 mutating **6/5 MET**, `blocked` **1/1 MET**, negative controls **21 of 24**;
clause 3 verdict half **COMPLETE**, landed half short (org-admin 2/4, onboarding 1/5); clause 1 leg half
**N/A**, flake half **MET**; **D-v28-5 root-caused + half-fixed** (was: unstarted).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this tik landed a measured, test-proven fix to the item it targeted) — (3) re-scope: n — (4) user-blocker: n (the remaining half is specified and routed, not a question needing an answer) — (5) cap-reached: n (3rd tik of this invocation) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D79 (**the mock was listening on a path the real client does not use, and its test drove the same unused path** — a green test about the mock talking to itself, on the one behaviour a user complained about; the 18th could-not-fail check of this milestone and the first inside Clerkenstein), D80 (`_method` is a DISPATCH KEY, not a hint — register the path *and* pin that only `DELETE` signs out, or "route the collection" becomes "any POST logs you out"), D81 (**an explicit sign-out should be sticky until an explicit login** — a handshake with `__clerk_identity` is an intentional login and must always establish; a bare handshake must not resurrect an ended session), D82 (a fix to a binary the running container does not carry must not be regression-"proven" by a suite run against the old build — say what the evidence cannot cover).
**Side-deliverables:** none.
**Routes carried forward:**
- `D-v28-5-cockpit-logout` → **HALF DONE. The remaining half is specified, not open-ended.** Implement D81's
  rule (`handleHandshake` must not auto-establish on a *bare* handshake after an explicit sign-out; a
  handshake carrying `__clerk_identity` always establishes), then prove (a) **and** (b) together on one
  rebuild: push the tag, re-pin `stack-demo/rosetta-extensions`, rebuild the `fake-fapi` container (the
  iter-11 precedent), then re-run this iter's four-step browser measurement and show step 3 bounce to
  `/login` and step 4 name the right hero on the FIRST click. Blast radius to check on that run: every
  Playthrough and the cockpit pass `__clerk_identity` so they should be unaffected, but `autoverify` may
  handshake bare — that is the thing to watch. Still **no Playthrough** (the user's explicit call).
- `NEGCTL-M256-cross-vantage` → unchanged at **21 of 24**; `pt-hiring-recruiter-compare` still needs a
  same-vantage control whose absence half is unmeasured (priced at iter-15).
- `FIX-M256-autoverify-fapi-libressl` → **second sighting.** It cost this iter a route probe: `curl` cannot
  handshake the mkcert leaf on macOS, so `POST /v1/client/sessions` could not be checked from the shell and
  had to be read from a browser trace and a Go test instead. Still routed, now with a second use case.
- Everything else from iter-15's list stands unchanged.
**Lessons:**
1. **Read the surface before probing it.** Ten minutes of grep established that the cockpit has no logout at
   all, which turned "the cockpit's logout is broken" into "the app's sign-out does not reach the mock" — a
   different investigation with a different instrument.
2. **A mock's test suite is the easiest place in a codebase to write a check that cannot fail**, because both
   sides of the conversation are yours. The fix for that class is to drive **the exact request the real client
   sends**, query string included, copied from a trace rather than from the mock's own route table.
3. **Say what the evidence cannot cover.** The suite would have gone green on this change and proven nothing,
   because the running container predates it. Reporting that green as verification would have been the
   favourable-sample dishonesty D-v28-13 was written about, in a different currency.

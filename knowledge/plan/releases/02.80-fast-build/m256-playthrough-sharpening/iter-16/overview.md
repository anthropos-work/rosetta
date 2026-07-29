---
iteration_type: tik
iter_shape: standard
status: in-progress
opened: 2026-07-29
---

# iter-16 — D-v28-5: measure the logout double-click, then fix it if the mechanism is ours

**Type:** tik · **Active strategy:** `TOK-01` move 4 ("close the honesty items … the D-v28-5 cockpit
Back-to-Cockpit / logout double-click fix")

## Step 0 — re-survey (mandatory)

- Both trees clean at iter open; `demo-2` up (16 containers, recovered at iter-15); drifted cockpit fixture
  intact at sha `99e2f315`.
- `@pt-negative-control` registry: **21 of 24**. The one sharpenable Playthrough left
  (`pt-hiring-recruiter-compare`) needs a same-vantage control whose *absence* half is unmeasured — priced
  and routed at iter-15, deliberately not opened here.
- **D-v28-5 is the last gate item that has never been attempted.** iter-10 tried and closed `closed-no-lift`
  because it was **not measurable**: the cockpit manifest had drifted from the roster, so no seat selection
  could be trusted. That blocker shipped in the harden pass (rext `0f36f71` made the roster authoritative),
  so the defect is measurable for the first time.

## Cluster / target identified

**D-v28-5**, the user's own report: *"logout to cockpit doesn't work: i have to click it twice+"*. A gate
clause in its own right, and by the user's explicit call it gets **no Playthrough** — so the deliverable is a
fix plus a recorded measurement, not a test.

What the surface actually is, established by reading before probing:

- The cockpit is a **launcher with no logout affordance at all** (grep: no logout / sign-out / session-remove
  path anywhere in `cockpit.py`). So "logout to cockpit" is the **app's** sign-out, after which the presenter
  returns to the cockpit to pick another hero.
- The app's sign-out is `apps/web/src/app/(unauthenticated)/logout/[[...logout]]/page.tsx`: `useClerk()`
  `signOut()` → clear `localStorage`/`sessionStorage` → `router.replace('/login')`.
- `signOut()` reaches Clerkenstein at `POST /v1/client/sessions/{id}/remove` → `handleSignOut`, which sets
  `s.signedIn = false; s.sessID = ""` **and nothing else** — notably it does **not** reset the registry's
  active seat, and it cannot clear the app-origin cookies (it is a different origin).
- The re-entry path is `handleHandshake`, which calls `establishLocked()` **unconditionally** and honours
  `__clerk_identity` best-effort.

## Hypothesis (explicitly NOT to be trusted without measurement)

The candidate mechanism is a **cookie-vs-server-state split**: the backend verifies `__session`
**networklessly** (`CLERK_JWT_KEY`) and never asks the FAPI, so clearing `signedIn` server-side does not make
the browser unauthenticated — the app-origin cookies outlive the sign-out. If so, the first click appears to
do nothing and a later one takes.

**This is a hypothesis and this milestone has refuted its own plan in 7 of 15 iters — including twice today,
where a stat accessor and a probe settle predicate were both wrong on first principles.** So Phase A measures
the real flow and the plan does not commit to a fix shape before it does.

## Expected lift

D-v28-5 **FIXED** — or, if the mechanism turns out to sit outside rext (a next-web behaviour, i.e. a
platform edit, which is forbidden), a **written diagnosis with the measurement attached** and a routed
verdict. iter-10 already set that precedent for this exact item and it was the right outcome then.

## Phase plan

- **Phase A — measure the real flow.** Drive a browser: log in as a hero via the cockpit's own handshake
  link, confirm the identity, then exercise the app's `/logout` and record, per step: the FAPI requests made,
  the app-origin cookies before/after, the landing URL, whether `/v1/me` 401s, and whether a *second*
  action is needed before a different hero's login takes. **Count the clicks** — the defect is a count.
- **Phase B — fix, only if the mechanism is rext-owned** (Clerkenstein / the cockpit / an injected artifact).
  A platform-source change is out of bounds unless it can be a sha-pinned `demopatch`, and a demopatch for a
  presenter-convenience defect would need its own justification.
- **Phase C — prove it**, by re-running the measured sequence and showing the click count drop. No
  Playthrough (the user's call), so the proof is the recorded before/after plus a unit test on whatever
  Clerkenstein behaviour changes.
- **Phase D — re-measure the suite** (the mock is on every login path, so a Clerkenstein change must be
  proven not to regress it): full suite ×3 cold reset-to-seed, plus the Clerkenstein module's own tests, plus
  its **Alignment DNA** score if the change touches a DNA'd capability.
- **Phase E — close**: commit both repos, tag + push + verify, corpus backfill.

## Escalation conditions

- **If the fix would require a platform-repo edit, STOP** and write the diagnosis. Zero platform edits is a
  hard release constraint; this defect is presenter convenience and does not earn a demopatch by itself.
- **If a Clerkenstein change touches a capability its Alignment DNA measures**, the DNA score must be
  re-taken — a mock that drifts from its source to fix a demo wart is a worse outcome than the wart.
- If the double-click proves **not reproducible** under measurement, that is a finding: record it with the
  click counts and route it back to the user, who reported it from real use.

## Acceptable close-no-lift outcomes

A measured diagnosis that locates the mechanism outside rext's reach, with the per-step evidence — the same
shape iter-05, iter-07, iter-10 and iter-12 produced, and better than a speculative fix to a mock that sits
on every login path in the suite.

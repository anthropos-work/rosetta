# M254 · iter-06 — decisions

## D1 — (f) root cause: wrong-default-hero measurement artifact (NOT a session-carry defect)
studio-desk's server (`src/index.ts`) gates in two stages: `requireAuth({signInUrl})` (networkless JWT verify —
PASSES on `--public-host`; the session carries), then `checkEnterpriseAndAdmin` which calls
`clerkClient.users.getOrganizationMembershipList({userId})` (a per-request round-trip to `CLERK_API_URL`, which
is the docker-network alias `api.clerk.com` → the fake-bapi, verified resolving from the studio container) and
admits only `STUDIO_ACCESS_ROLES = {admin, org:admin, content_creator, org:content_creator}`, 303-redirecting
everyone else to `WEB_APP_URL`. The FCP probe defaulted to `maya-thriving` (roster OrgRole=**member**) → bounced
to `:13000/home` by design → the probe measured the wrong app (iter-04's "302 → :13000/login" was the *unauth*
form of the same gate). The fake-bapi returns the CORRECT roles (dan-manager user_seed_demo-1_3 →
`"role":"admin"` total_count 1; maya → `"role":"member"`), seeded from `FAKE_FAPI_ROSTER` by
`cmd/fake-bapi/seedRosterMemberships`. Live: dan-manager / dana-manager / rae-recruiter reach the studio shell
(`:19000`, skeleton paints); maya is correctly bounced.

## D2 — fix: FCP default identity maya-thriving → dan-manager (rung-zero)
`run-studio-fcp.sh:37` + `tests/studio-fcp.spec.ts:15` default `maya-thriving` → `dan-manager` (org_clerkenstein
OrgRole=admin, present in the default Workforce demo), with corrected comments explaining the studio-eligibility
gate. rext commit `cbe9256`, tag `july-jitter-m254-studio-fcp-identity` pushed to origin (verified via
`git ls-remote`). Zero platform edits. No re-bring-up needed — dan-manager reaches the shell on the current demo.

## D3 — (f) FCP <1 s cold p95 on billion: p50 <1 s, p95 tailnet-jitter-bound → disposition-pending
With dan-manager, `reached_shell_all: true` across 15 cold samples (3 batches). Skeleton paint is **bimodal**:
fast runs 628-726 ms (well under the 1 s gate → the M253 shell-first-paint fix HOLDS on billion), but frequent
outliers (1443 / 1379 / 1631 / 2014 / 1566 / **4943** ms) push p95 to 1.4-5 s. reachedShell is always true, so
the outliers are latency, not failures. The studio server's per-request auth round-trip is IN the shell-paint
path, compounding tailnet HTTPS RTT variance. This is the roadmap's anticipated "M253's <1 s FCP is
environment-sensitive... re-confirmed cold on billion at M254 — a fix iteration may surface there." The <1 s
(p95 AND max ≤ gate) was set on localhost (M253). → surfaced for **coordinator disposition** (analogous to the
coordinator-approved (b) presence-only disposition): the shell-paint fix demonstrably works on billion (p50
< 1 s); the p95 gate as literally stated is tailnet-network-bound. NOT a studio-code regression.

## D4 — (c)-studio render LIVE-confirmed (unblocked by the (f) fix)
As dan-manager on billion: studio shell reached (`:19000`, 200), the SPA boots to `up-user-profile` (dan's
`up-admin-badge`), and the account dropdown carries a "Back to Cockpit" item (text present; it is a `<button>`
reading `import.meta.env.VITE_COCKPIT_URL`, not an `<a href>`, so it renders text not an anchor). **0 prod-eject
links** (`app.anthropos.work`) in the studio shell — the M249 studio-desk-logo-url / logout-url patches hold.
`:17700` is baked in `dist/public/assets/main-*.js` (resolve-to-stack). So (c) render side is now LIVE on 3/4
apps (next-web + hiring + studio); academy alone carries the iter-05 durability defect.

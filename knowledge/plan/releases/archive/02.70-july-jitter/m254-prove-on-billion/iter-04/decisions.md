# M254 · iter-04 — decisions

## D1 — gate (b) presence-only disposition (coordinator-approved)
The 2 voice MANAGER views (`hire-voice-fail`, `asmt-voice-pass-en`) that render ~230 chars are the symmetric
extension of the already-ACCEPTED `DEF-M240-01` (voice PLAYER cells presence-only): the customer interview
VIDEO is Bunny-CDN-hosted and the demo box is Bunny-keyless, so neither the player nor the manager voice result
can render real media. Presence-only is the faithful default. Gate (b) = 45/45 landable all landed + 4 voice
cells presence-only (2 player + 2 manager). Follow-up (non-blocking, tracked): `manager_presence_only` flag in
buildPairs + seeder + `content-denominator.json` 47→45 + a re-seed.

## D2 — (f) studio-desk session-carry is the roadmap-flagged env-sensitive FCP risk
studio-desk `:19000` → 302 → `:13000/login` (baked `VITE_CLERK_SIGN_IN_URL` = web app login). The cockpit
session isn't carried to studio-desk on the --public-host demo, so the FCP flow lands on the stack web app and
the studio shell never paints. NOT a prod-eject (:13000 is the stack's own web app). This is exactly the risk
the roadmap flagged for M254 ("M253's <1s FCP is environment-sensitive... re-confirmed cold on billion at M254
— a fix iteration may surface there"). Routed to a fix-iter.

## D3 — (g) host-sensitive test failures categorized
Isolated-clean run (stray :23077 listener killed) still yields 9 failures: 2 pure nvm env-artifacts
(`test_missing_node_documents` ×2) that should get an env-guard/skip-with-reason, + 7 real/host-sensitive
(launcher reap/stop intra-run listener leakage, 2 mutation meta-tests, `test_apply_revert_round_trip_on_the_real_next_config`).
Routed to a rext fix-iter (rung-zero).

## D4 — checkpoint / cap-reached
Dense measurement cluster (4 gate parts measured + 2 residuals characterized) consumed the invocation's
practical budget with a heavily-loaded context. Exiting cap-reached so the orchestrator re-invokes a FRESH
agent (the prove-on-billion "fresh agent per run" discipline) with a tight resume brief. billion demo left UP
+ fresh-green (autoverify 06:25:41Z) for the next agent's (c)/(e)/(h) + the (f)/(g) fix-iters.

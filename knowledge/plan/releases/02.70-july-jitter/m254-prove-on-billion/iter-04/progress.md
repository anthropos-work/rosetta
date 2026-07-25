**Type:** tik (measurement-cluster) — TOK-01 clusters 2+3. Read-only sweeps + studio-FCP leg.

# M254 · iter-04 — progress

## Close — 2026-07-25

**Outcome:** measured 4 gate parts against the live billion demo — (a) re-confirmed green, (b) MET-with-disposition
(45/45 + 4 voice presence-only, coordinator-approved), (d) MET both vantages (8/8 sections incl 3 drift-fixes),
(c) prod-eject side proven (0 escapes/133 pages). Surfaced 2 residuals: (f) studio-desk session-carry on
--public-host (FCP not gradeable — lands on web app) and (g) 9 host-sensitive test fails (2 nvm env-artifact +
7 to fix). Pending: (c) Back-to-Cockpit render, (e) builder Playthrough, (h) latency solo.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (overall) — parts a/b/d MET; c/e/h pending; f/g residual → fix-iters
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
(the f/g residuals are routed-forward Fate-3 fix-iters, not blockers) — (5) cap-reached: **y**
(dense measurement cluster consumed the invocation's practical budget; context heavily loaded — checkpoint for
a fresh agent per the prove-on-billion "fresh agent per run" discipline) — (6) protocol-stop: n —
Outcome: exit-5 (cap-reached / checkpoint)
**Decisions:** (b) presence-only disposition (coordinator-approved, DEF-M240-01 extension); (f)/(g) residuals
characterized + routed to fix-iters; verify.sh skillpath-default drift noted.
**Side-deliverables:** none (measurement + the iter-03 rext fix already shipped separately).
**Routes carried forward:** (c) Back-to-Cockpit render check; (e) builder Playthrough; (h) latency solo +
Playthroughs; **fix-iter (f)** studio-desk session-carry on --public-host; **fix-iter (g)** host-sensitive
tests (7 disentangle/fix + 2 nvm env-guard); (b) denominator/manager_presence_only follow-up (non-blocking);
verify.sh default-services skillpath drop (non-blocking).
**Lessons:** (1) The (f)/(h)/studio-FCP gates require a FRESH green autoverify (<4h); a ~9h-old verdict is
refused. Regenerate via `autoverify.sh --project demo-1 --offset 10000 --services "<explicit-no-skillpath>"` —
and ALWAYS pass `--services` (the default still carries stale skillpath). (2) The --public-host demo bakes
studio-desk's sign-in to the web app's `:13000/login`; the studio session-carry is the env-sensitive FCP risk
the roadmap called out. (3) The foreground-poll pattern (detached op + sentinel + foreground blocking poll)
replaces background+yield, which does not reliably re-invoke in this environment (cost ~7.5h across 2 stalls).

---
iteration_type: tik
milestone: M254
iter: 06
status: closed-fixed
---

# M254 · iter-06 — (f) studio session-carry investigation + fix (the linchpin)

**Type:** tik · **Active strategy:** TOK-01 (cluster-per-tik live re-prove) — cluster 4 fix-forward, (f).

## Cluster / target
(f) studio session-carry — iter-04 flagged `:19000 → 302 → :13000/login` on `--public-host` as a
session-carry defect blocking (c)-studio render + (e) the builder Playthrough. Investigate the true cause,
fix in tooling (rung-zero), OR escalate if it needs a platform edit.

## Hypothesis
The studio shell doesn't paint on `--public-host` because the cockpit session isn't carried to studio-desk.

## Outcome (falsified + re-rooted)
The hypothesis was WRONG. Studio auth works on `--public-host`. Root cause: the FCP probe's default identity
is `maya-thriving`, an EMPLOYEE (roster OrgRole=member). studio-desk's server-side `checkEnterpriseAndAdmin`
(`src/index.ts`) admits only `STUDIO_ACCESS_ROLES = {admin, org:admin, content_creator, org:content_creator}`
and 303-redirects everyone else to the web app. maya is bounced BY DESIGN → the probe measured the wrong app.
Admin heroes (dan-manager / dana-manager / rae-recruiter, roster OrgRole=admin) REACH the studio shell live.
Fix: default identity `maya-thriving` → `dan-manager` (studio-eligible) in `run-studio-fcp.sh` +
`studio-fcp.spec.ts` + corrected the misleading comment. rext `cbe9256`, tag
`july-jitter-m254-studio-fcp-identity` on origin. Zero platform edits, no re-bring-up needed.

## Expected lift → realized
(f) session-carry MET (fix shipped, dan reachedShell all true live) + unblocks (c)-studio (now LIVE 3/4) +
(e). FCP shell paint p50 637-726 ms < 1 s on billion (M253 fix holds); p95 tailnet-jitter-bound → disposition.

## Phase plan
verification.md investigate→fix-forward. rung-zero (tag on origin) before any consumer use.

## Escalation conditions
If the fix needed a platform-repo edit → ESCALATE. It did NOT (tooling-only default-identity change).

## Acceptable close-no-lift
A documented falsification of the session-carry hypothesis satisfies the protocol even without a metric move.
(Realized: falsified + a real tooling fix shipped → closed-fixed.)

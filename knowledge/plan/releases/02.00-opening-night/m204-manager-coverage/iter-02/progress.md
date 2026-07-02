**Type:** tik — first tik under TOK-01 (manager-surface-per-iter). Protocol:
`corpus/ops/demo/playthroughs.md` § iteration protocol.

# iter-02 — Workforce funnel + member roster

Probed the real manager UI as Morgan (pt-manager) → all four manager surfaces render real data at Org A
(Meridian Labs, size 40), no perf wall. Authored the manager product manifest (`workforce.yaml`), two page
objects (`WorkforcePage` funnel/gap anchors + `MembersPage` roster anchors), anchored route shapes
(`WORKFORCE_URL`/`MEMBERS_URL` in `url-shapes.ts` + pure-logic unit coverage), and two Playthroughs
(`pt-workforce-funnel`, `pt-workforce-roster`). Both green on demo-1; `ptreport` reconciles them `[PASS]`.

## Close — 2026-07-02

**Outcome:** +2 passing manager UCs (`workforce-intelligence.skills-funnel.UC1` +
`workforce-intelligence.roster.UC1`); metric 0 → 2. Both Playthroughs green + reconciled `[PASS]` by ptreport.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (2 of the 3 declared manager journeys' UCs passing; drill-down + succession remain; 5-run gate not yet run)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (iter-02/decisions.md — the runner reporter-override fix, in-scope Fate-1).
**Side-deliverables:** the `run-playthroughs.sh` reporter-override fix + inline `ptreport` reconciliation — a
latent M202/M203 wiring defect (a CLI `--reporter=list` REPLACED the config reporter list, suppressing the
JSON-file reporter ptreport reads → stale `last-run.json`, decoupled reconciliation). Fixed + documented in
`playthroughs.md`. Separate concern from the manager UCs; committed in the same iter commit, does not upgrade
the tik's status (it grades planned-scope: +2 manager UCs).
**Routes carried forward:**
- iter-03 (next tik) → **Member drill-down (activity-dashboard)** — `assignment-monitoring.assign-and-track.UC2`
  (`/enterprise/activity-dashboard` + per-member drill-down; the probe confirmed the surface renders a 20-row
  table + stat cards at Org A).
- then → **Succession / at-risk** — `workforce-intelligence.talent-pool.UC1` (`/enterprise/workforce/succession`;
  probe confirmed "Succession Planning" + 10 rows of role→candidate data + At-Risk labels render).
**Lessons:** (1) the base pt-world stories model DOES render the M36 org-dashboard aggregates at Org A size 40 —
no seed expansion needed for the manager READ surfaces (the key Phase 0b known-context, resolved positively).
(2) The WI SPA tabs are client-side state (the URL stays `/enterprise/workforce`), so anchor the funnel on the
default landing view, not a tab route. (3) The `(new)` in the succession route is a Next.js route group —
invisible in the URL (`/enterprise/workforce/succession` works). (4) Always let the config reporter set fire —
a CLI `--reporter` override silently decouples reconciliation (added to the protocol doc).

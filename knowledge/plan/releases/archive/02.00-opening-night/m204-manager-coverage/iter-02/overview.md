---
iter: 02
milestone: M204
iteration_type: tik
status: closed-fixed
created: 2026-07-02
---

# M204 iter-02 — tik: Workforce funnel + member roster

**Type:** tik

**Active strategy reference:** TOK-01 (manager-surface-per-iter) — the first tik under the bootstrap strategy.

**Step 0 re-survey:** metric baseline = 0 manager UCs declared / 0 passing (confirmed — the manifest held only
M203's 3 employee products / 6 UCs). TOK-01's named iter-02 target (Workforce funnel + roster) is untouched +
meaningful. No substitution.

**Cluster / target identified:** journey 1 of the 3 declared manager journeys — the manager's landing surface
(`jump_to /enterprise/workforce`), the lowest-risk drivable READ surface + the natural first manager-login proof.
Maps to two M201-corpus UCs: `workforce-intelligence.skills-funnel.UC1` (the mapped→verified funnel + org-scale
gap on `/enterprise/workforce`) + `workforce-intelligence.roster.UC1` (the member roster on `/enterprise/members`).

**Hypothesis:** the pt-world base stories model (Org A / Meridian Labs, size 40) already renders the M36
org-dashboard aggregates (funnel + roster) — so the two manager Playthroughs go green with page objects + specs,
no seed expansion. (The key known-context from the Phase 0b audit.)

**Expected lift:** +2 passing manager UCs (of the declared manager set).

**Phase plan:** measure (probe the real manager UI) → declare (`workforce.yaml`) → page-object
(`workforce-page.ts` + `members-page.ts`, route shapes in `url-shapes.ts`) → play (`--grep @pt:pt-workforce`) →
diagnose → re-measure (`ptreport`).

**Escalation conditions:** a manager surface un-drivable without a platform edit →
`unimplementable-without-platform-edit` (escalate, declare in `unimplementable.yaml`, NEVER edit the platform); a
perf-wall stall at Org A (size 40 — not expected per the audit) → escalate.

**Acceptable close-no-lift outcomes:** if the probe showed the funnel/roster rendered EMPTY at size 40 (a real
seed-scale gap), the falsification "the base model does not render M36 aggregates at size 40 → seed needs
expansion" would satisfy the protocol even without the +2 lift (routed to a seed-expansion iter).

## Outcome
Both hypothesis-confirming: the probe showed all four manager surfaces render REAL data at Org A size 40 (no
perf wall). Authored `workforce.yaml` + `WorkforcePage`/`MembersPage` + the two specs; both Playthroughs pass
green; `ptreport` reconciles them as `[PASS] passing`. **+2 passing manager UCs.**

**Side-discovery (foundation fix):** the runner passed `--reporter=list`, which REPLACES the config reporter
list and suppressed the JSON-file reporter `ptreport` reads → `last-run.json` was stale, decoupling the
four-state reconciliation from the run (a latent M202/M203 defect). Fixed the runner to not override the
reporter + to reconcile inline. Separate concern from the manager UCs; recorded as a side-deliverable.

## Close
See `progress.md`.

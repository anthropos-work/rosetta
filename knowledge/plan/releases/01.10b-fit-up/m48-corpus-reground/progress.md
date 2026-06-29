# M48 — progress

## Section checklist

### S1 — Drift survey (corpus/architecture + corpus/services vs current clones)
- [ ] survey `corpus/architecture/*` + `corpus/services/*` against the M47-current clones — produce a material-lag gap list (what changed since the docs were written: services, flows, signatures, new features)
- [ ] flag the **member-AI-readiness** feature as the headline undocumented addition (→ S2)
- [ ] record the gap list in `spec-notes.md`

### S2 — Member-AI-readiness flow doc (LOAD-BEARING — the M51 seeder contract)
- [ ] map the data model: tables, org-enablement flag, the **3-step onboarding/evaluation**, scoring, started-vs-completed states (across app / jobsimulation / skiller / skillpath / cms)
- [ ] map the surfaces: the manager dashboard (funnel / snapshot / how-we-measure tabs) + the member-facing onboarding element (next-web `ai-readiness/`)
- [ ] author the doc (a `corpus/services/` or `corpus/architecture/` home — decide during the read) — the contract M51 builds its seeder against (which tables to write, at what cardinality, for the 80%/3-step funnel)
- [ ] cross-link from the parent index + CLAUDE.md

### S3 — Reconcile the surveyed material drift
- [ ] update the architecture/service docs that materially lag current code (from the S1 gap list, material-first)
- [ ] update `last_updated` dates on touched docs

### S4 — Retire stale claims + cross-link
- [ ] correct the ant-academy "RESOLVED / cloned" claim (it is NOT in `repos.yml` → not cloned today; the FIX is M49 #5 — M48 makes the DOC tell the truth + notes M49 owns the fix)
- [ ] note (don't fix) the stale rext-tag references → M49 #1 owns the `.agentspace/rext.tag` source-of-truth fix
- [ ] verify all new/edited cross-references resolve

## Notes
- **Docs-vs-code only — never touches the live demo** → M48 ∥ M49 (disjoint file clusters: `architecture`+`services`
  here vs `ops`+rext there).
- **Material-lag-first** (per the overview): the AI-readiness contract (S2) is the non-negotiable deliverable
  (M51 is blocked without it); the broader drift sweep (S1/S3) is time-boxed to what materially lags.

# M224 — Decisions

_Implementation + strategy choices with rationale (incl. tok strategy revisions)._

| # | Decision | Rationale | Date |
|---|----------|-----------|------|
| — | _(intra-iter decisions live in `iter-NN/decisions.md`)_ | | |

---

## TOK-01: recruiter-render-first — 2026-07-16

**Tok type:** bootstrap (iter-01)

**Initial strategy:** Drive the **render loop** the protocol prescribes — seed → render → **attribute** → fix →
re-render — against a **LOCAL** demo, ordered **recruiter-scoreboard-FIRST**. The exit gate IS the recruiter's
comparison surface (`/enterprise/activity-dashboard → AI-Simulations → [simId]`) painting ≥40 comparable non-junk
candidate rows per **each** of the 5 seeded HIRING positions, non-degenerate distribution, closure green, 0
prod-eject, over ≥3 cold reset-to-seed runs. So every early tik targets **rows-on-the-recruiter-scoreboard**; the
2 candidate `/profile` heroes (In-scope but not the gate) are layered **after** the scoreboard is green.

Each tik's fix is chosen by **attribution against `hiring.md`'s traced read-path**, in the M219-trap order:
  1. **Clerkenstein identity gap** — the client re-skin (`useGetClerkOrganization`) reads org
     `publicMetadata.isHiring` from the **FAPI**; without it the org renders as a normal Workforce org and the
     "Results" framing / hiring cohort treatment never appears. **Prefer the Clerkenstein FAPI wiring** (roster+
     resource, M39 `org_name`/`org_slug` precedent) over any patch (D-DESIGN-2). **Touches `clerk-frontend/` →
     carries the BLOCKING `/align-run` on `clerk-js-5`** (keep `Client/signed-in` critical/shape +
     `Me/universal-user` standard/shape green; the named `SessionToken/decoded-identity` exact gene is unaffected).
  2. **Render-gate bypass** — an M219-class gate (a `CycleID==nil→buildLiveResponse` analog, an undefined
     PostHog flag, or the silent-403 on the `OrgFeatureInsights` Casbin permission) that ignores the seed → route
     to a **sha-pinned demo-patch** on the demo's ephemeral clone (read `demopatch-spec.md` first) only if there is
     no env/config/Clerkenstein seam.
  3. **Seed gap** — a row the scoreboard reads that the M223 funnel didn't write (e.g. a missing membership, a
     NULL-bubbling federated `Session!` twin, a drill-down `validation_*` row for the assessed candidate hero) →
     `stack-seeding`.

**Rationale:** The opening move is "reach + measure + attribute" NOT "fix", because M223 **already** writes the
score source (the `local_jobsimulation_sessions` mirror pair, 45 candidates × 5 sims, verified ALIGNED in Phase
0b). The render risk is therefore concentrated in the render-GATE and the Clerkenstein re-skin identity — exactly
where a fix-before-measure would be a guess (the M219 trap the milestone is named against). Front-loading a real
baseline render makes the first fix land against a measured, attributed gap.

**Strategy class:** new-direction (bootstrap — no prior strategy to compare against).

**Distance-to-gate context:** metric = `min(rows-painted across the 5 sims)` on a cold-seeded recruiter scoreboard
(+ the distribution / closure / eject sub-gates). Baseline = **UNMEASURED** (the surface has never been rendered on
a local stack — no recruiter seat, no isHiring wiring). Gate = **≥40** per each of 5 sims, non-degenerate, closure
green, 0 ejects, ≥3 cold runs.

**Next-tik direction:** iter-02 (first tik) = the **enabling-scaffold + baseline-render** tik. (a) Add the
recruiter (`vantage: manager`) hero seat to the Meridian Talent 4th story in `presets/stories.seed.yaml` with
`jump_to` at the comparison surface; (b) wire Clerkenstein FAPI `isHiring` (`clerk-frontend/registry.go` RosterEntry
`org_is_hiring` → `DemoUser.OrgIsHiring` → `resources.go::orgMemberships` `PublicMetadata["isHiring"]`) **+ the
`/align-run` gate**; (c) author a **recruiter render-probe** under `stack-verify/e2e/` (log in via the cockpit
handshake, reach each of the 5 `[simId]` scoreboards, count comparable rows + capture the score distribution + run
closure/eject checks); (d) tag rext; (e) bring up a **LOCAL** demo consuming the tag; (f) log in as the recruiter,
**measure the baseline**, and **attribute** the per-sim gap — no fix before the attribution. If reaching/measuring
is blocked on a missing probe capability, iter-03 becomes a **tooling-iter** (protocol refinement).


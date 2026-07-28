# M256 — decisions

Release-level binding decisions **D-v28-1 … D-v28-12** live in
[`../../../roadmap.md`](../../../roadmap.md) § Active — v2.8. This file carries the milestone's own
**strategy chain** (`TOK-NN`) plus milestone-level decisions. Intra-iter decisions live in each
`iter-NN/decisions.md`.

---

## TOK-01: cheap-lever speed, then the cluster that discharges two clauses at once — 2026-07-28

**Tok type:** bootstrap (iter-01)

**Initial strategy.** Four moves, strictly ordered, each landing before the next opens:

1. **Measure the denominator first, on this box, and change nothing until it exists.** A relative gate
   (D-v28-12: median per-Playthrough ≤ **0.79×** a same-stack pre-work baseline) is unfalsifiable without a
   measured starting point. n=3 on the local `demo-2`, environment stated with the number, recorded in
   `progress.md`. Report **median per non-LLM Playthrough** (the gated metric), the **suite wall-clock**
   (reported, not gated), and the **studio lane** separately (budgeted separately per D-v28-9).
2. **Take the per-test latency lever, not the parallelism lever.** Clause 1 is a **per-test median**;
   worker count cannot move a per-test median (iter-01 D1). The real lever is the residual
   **`networkidle`**: **12 of 18** login call sites omit `waitUntil` and inherit `cockpit-login.ts`'s
   `'networkidle'` default on an app whose own helper doc records that `networkidle` "resolves late and for
   the wrong reason"; plus **8 further unfenced violations** in the harness (2 page-object `goto` overrides
   + **6 unbounded `waitForLoadState`** sites the Phase-0b audit found). M254 iter-10 measured 13 min →
   3.8 min from exactly this class. Land the fix **with a widened fence** — the existing
   `home-login-networkidle.unit.spec.ts` guards only `/home`-landing specs, which is precisely why the
   other holes survived. Ship the machine-checked per-spec **`MUTATES` / `READ-ONLY` / `UNKNOWN`** tag in
   the same iter, following the `@pt:`-grammar **twin lockstep test** shape the audit identified as prior
   art (never a third unfenced regex).
3. **Land org-admin before onboarding, because org-admin discharges two clauses with one body of work.**
   All four curated org-admin UCs declare a persist-then-observe final → 4 mutating Playthroughs → **5**
   with `pt-assignment-assign`, which is exactly clause 2's `≥ 5 mutating` floor, while being half of
   clause 3's scope (D-v28-4). Onboarding is **seed-blocked** (audit F5 answers the overview's Open
   Question 1: **no** — `UsersSeeder` writes a membership for every user unconditionally; there is no
   pre-onboarding state, and none can be *declared* into existence), so it is ordered second and its cost
   is a **seeder + capability + roster seat**, not just specs. A seed wall must not be allowed to starve
   the clauses org-admin already discharges.
4. **Close the honesty items last, deliberately, not as leftovers.** Negative controls; the `blocked`
   outcome from an **RBAC/Sentinel deny** rather than an entitlement tier (iter-01 D4, refuted in-iter by
   the Phase-0b audit — `actor.entitlement` reaches no DB column, and `ptvalidate`'s precondition check
   **fail-opens** on it); the D-v28-5 cockpit Back-to-Cockpit / logout double-click fix; and a **written
   verdict for every remaining uncovered curated UC** including the 5-release-old M206/M207 reservations.

**Rationale — why this is the right opening move.** The milestone's own plan named parallelism as the
headline lever and the plan review already broke half of it. iter-01 broke the other half: **the D-v28-12
re-cut dissolved the requirement entirely**, because a median per-test metric is indifferent to worker
count. That removes the largest, riskiest item (a refactor of an Alignment-DNA-gated mirror engine) from
the critical path and frees the whole milestone budget for work that actually moves the three clauses. What
replaces it is cheaper *and* better-evidenced: a defect class with a measured precedent (M254 iter-10), a
known target list (12 + 8 sites), and an existing fence to widen rather than invent. The ordering rule
throughout is **discharge-per-unit-of-work**, which is why org-admin — the only cluster that serves two
clauses simultaneously — goes before the coverage cluster that is seed-blocked.

**Strategy class:** `new-direction` (bootstrap — no prior strategy to compare against).

**Distance-to-gate context.** Gate metric: **median per-Playthrough**, target **≤ 0.79× baseline**, on the
**post-coverage** suite (denominator 18 → ~27). Starting value: **not yet measured on this host** — that is
iter-02's entire job, and it is the first thing that happens. The only comparable prior number is billion's
228 s / 18 tests (~6.4 s per non-LLM test), which **must not be quoted as comparable** to any M256 number
(D-v28-12; the absolute billion re-measure is routed to M258). Clause 2 starts at **1** mutating
Playthrough and **0** `blocked` / **0** negative controls; clause 3 starts at **0** onboarding and **0**
org-admin.

**Known context carried from the Phase-0b audit (verdict YELLOW, `kb-fidelity-audit.md`).**
- **F4** `actor.entitlement` is declared-only → the `blocked` outcome needs a different refusal surface.
- **F5** no pre-onboarding state exists → onboarding needs a seeder, not just tests.
- **F6** `--reset` is **whole-stack** (`doReset` takes no org filter; it truncates
  `public.organizations`/`users`), and `pt-world.seed.yaml`'s header comment claims the opposite. Every
  `run-playthroughs.sh --reset` on `demo-2` therefore **destroys the showcase world** — acceptable, because
  `demo-2` is dedicated to this milestone, but it must be stated, not discovered.
- **Gap 4** 8 unfenced `networkidle` violations (2 `goto` + 6 unbounded `waitForLoadState`) — folded into
  move 2 above.
- **Gap 7** `run-content-stories.sh` recomputes a 47-vs-pinned-45 pair count and `sys.exit(2)`s, so the
  content-stories sweep **refuses to start**. Not M256's suite; **Fate 3 → M257/M258**, which compose it.

**Next-tik direction (iter-02).** Measure the baseline and nothing else. `run-playthroughs.sh 2 --reset`
from the authoring copy with `stack-demo/rosetta-extensions/demo-stack/stacks/demo-2/bin` prepended to
`PATH` (the M204 iter-05 gate-run prereq), n=3, `PT_HOST=localhost`, `PT_APP_SCHEME=http`. Record: per-test
durations from `report/last-run.json`, the median over the **non-LLM** subset, the studio lane separately,
the suite wall-clock, and the environment. Do **not** change harness code in iter-02 — a baseline measured
on already-modified code is not a baseline.

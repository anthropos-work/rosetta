---
iter: 1
milestone: M236
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-07-20
metric_pre: "0/31 landable (session × action) pairs proven live on billion"
metric_post: "0/31 (tok — no gate progress by design)"
---

# iter-01 — bootstrap tok: author the first strategy

## Inputs

- `../overview.md` — scope, exit gate, the expanded `In:` list (3 Fate-3 clusters inherited from M235)
- `../../m235-prove-it-lands/carry-forward.md` — the 3 root-cause clusters routed here
- `../../m230-academy-demo-fill/carry-forward.md` — Cluster 3's origin
- Protocol: `corpus/ops/verification.md` + `corpus/ops/demo/tailscale-serve.md`
- Live baseline measured this iter (B1–B5 in `progress.md`)

There are no prior iters and no prior TOK chain. This iter's job is to author the *first* strategy.

## Initial strategy — "publish, then prove, then triage by arm"

Four phases. The ordering is forced by the baseline, not chosen.

**Phase P — publish (tik-1).** `billion` cannot obtain the v2.5 tooling: `origin/main` is the exact M228
commit `billion` already runs, and 0 of the 13 `playbill-*` tags are on origin (B2). Push
`main` (fast-forward, 20 commits) + the `playbill-*` tags, re-pin `billion`'s
`.agentspace/rext.tag` to `playbill-m235-hardened`, and check the consumption clone out at it. Prune the
107.6 GB reclaimable build cache first (B5). **Nothing else in the milestone can start until this lands** —
this is the pre-flight rung the milestone plan never named.

**Phase C — cold bring-up (tik-2).** A cold reset-to-seed `/demo-up` on `billion` at the new tag, with the
default-on public-host path (D-DESIGN-3). This is where the M230 Cluster-3 prerequisite bites — the 2
drifted next-web demopatch manifests must re-anchor or the build refuses. Success criterion: stack UP,
`autoverify.json` fresh-green, cockpit serves a non-404 `content-manifest.json` with all 4 products.

**Phase L — land the arms (tik-3…N).** Drive the 31 (session × action) pairs, **one product arm per iter**,
in descending evidence-density order: `simulation` (26 actions, the M235 fixture matrix — the arm most
likely to work first try) → `skill-path-legacy` (4, carries Cluster-2's version-match / status-vocabulary /
mirror-uniqueness calibration) → `skill-path-new`/academy (1, depends on the M230 catalog fill + the
`app/cmd/academy-seed` wiring) → `ai-labs` (0 landable; prove the 2 presence rows + the `lab_sessions`
NOT-NULL DDL calibration). Each iter: measure the arm, triage every non-landing pair to a root cause, fix
in the seeder/manifest, re-measure.

**Phase H — harness + budget (interleaved).** The new content-stories seat-login coverage plumbing
(Cluster 1) is **authored against the first live seeded render**, not up front — M235 proved authoring it
blind ships an incorrect load-bearing harness. So it is authored in the same iter that first lands a real
simulation result page, then extended per arm. The p95 click→ACCESS measurement rides the existing
`rext stack-verify/e2e/run-latency.sh` harness once ≥1 arm lands.

## Rationale

- **Publish-first is not optional sequencing** — it is the only order in which any other phase is testable.
  The baseline (B2) turns what the plan treated as an implicit given into the milestone's first deliverable.
- **One arm per iter** keeps the scope-creep tripwire enforceable: 31 pairs across 4 arms with genuinely
  different failure modes (fixture replay vs. version-match vs. catalog-fill vs. DDL) would otherwise
  collapse into one unbounded iter with an unclear close status — exactly the anti-pattern Phase 2 names.
- **Descending evidence-density** front-loads the arm with the most offline proof behind it (M235
  unit-proved the 13-session matrix), so the harness (Phase H) gets calibrated against a *known-good*
  render rather than co-debugged with a broken one.
- **Harness-after-first-render** is a direct consequence of M235's USER-BLOCKER-M235-02 finding, carried
  forward as Cluster 1.

## Strategy class

`new-direction` — bootstrap tok; there is no prior strategy to compare against.

## Distance-to-gate context

Primary metric **0 / 31** landable (session × action) pairs proven live. Secondary components all at 0:
academy grid cards (Thread A), p95 click→ACCESS measurement, ai-labs presence rows. The distance is
dominated by a **structural** gap (the tooling is not on the host), not by defect count — so the first
tik should move the metric's *reachability*, and the first real numerator movement is expected in Phase L.

## Next-tik direction

**iter-02 (tik) executes Phase P**: prune `billion`'s build cache, push `rosetta-extensions` `main` +
the 13 `playbill-*` tags to origin, re-pin `.agentspace/rext.tag` to `playbill-m235-hardened`, and check
`billion`'s consumption clone out at that tag — verifying the M217 FATAL pin guard passes. Expected lift on
the primary metric: **0** (this is a reachability precondition, and iter-02 must close honestly as such —
`closed-fixed` on planned scope, `Gate: NOT MET`). Acceptable close-no-lift outcome: if push is refused for
an access/policy reason, that is a genuine user-blocker (the milestone cannot proceed) and exits the session.

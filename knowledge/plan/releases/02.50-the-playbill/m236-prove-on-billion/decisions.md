# M236 — Decisions

## TOK-01: publish-then-prove — 2026-07-20

**Tok type:** bootstrap (iter-01)

**Initial strategy:** four ordered phases, the ordering forced by the live baseline rather than chosen.

- **Phase P — publish (iter-02).** Push `rosetta-extensions` `main` (fast-forward, 20 commits) + the 13
  `playbill-*` tags to origin; re-pin `billion`'s `.agentspace/rext.tag` to `playbill-m235-hardened` and
  check its consumption clone out at that tag; prune the host's 107.6 GB reclaimable build cache first.
- **Phase C — cold bring-up.** A cold reset-to-seed `/demo-up` on `billion` at the new tag, public-host
  default-on (D-DESIGN-3). Success = stack UP, fresh-green `autoverify.json`, cockpit serving a non-404
  `content-manifest.json` with all 4 products.
- **Phase L — land the arms, one product arm per iter,** in descending evidence-density order:
  `simulation` (26 actions) → `skill-path-legacy` (4) → `skill-path-new`/academy (1) → `ai-labs`
  (0 landable; prove the 2 presence rows). Per iter: measure the arm, triage every non-landing pair to a
  root cause, fix in seeder/manifest, re-measure.
- **Phase H — harness + budget, interleaved.** The new content-stories seat-login coverage plumbing is
  authored **against the first live seeded render**, not up front; the p95 click→ACCESS measurement rides
  `rext stack-verify/e2e/run-latency.sh` once ≥1 arm lands.

**Rationale:** the baseline measurement (iter-01 B1–B5) turned up a structural blocker the milestone plan
never named — `billion` **cannot obtain the feature under test**. It consumes `rosetta-extensions` only at
a tag named in `.agentspace/rext.tag`, guarded FATALly since M217; it pins the M228 tag; `origin/main` **is**
that M228 commit; and **0 of 13** `playbill-*` tags are on origin. Every v2.5 milestone's tooling
(M230/M232/M233/M234/M235) lives only in the local authoring copy. So publish-first is not preferred
sequencing — it is the only order in which any other phase is testable.

The rest of the ordering follows from evidence density and from M235's Cluster-1 finding: one arm per iter
keeps the scope-creep tripwire enforceable across 4 arms with genuinely different failure modes (fixture
replay vs. version-match vs. catalog-fill vs. DDL), and harness-after-first-render is forced by
USER-BLOCKER-M235-02 (authoring the descriptor blind ships an *incorrect*, not merely uncalibrated,
load-bearing harness).

**Strategy class:** new-direction — bootstrap tok; no prior strategy exists to compare against.

**Distance-to-gate context:** primary metric **0 / 31** landable (session × action) pairs proven live;
all secondary components (academy real cards, ai-labs presence rows, p95 measurement) also at 0. The
distance is dominated by a **structural reachability gap**, not by defect count — so iter-02 is expected to
move *reachability* with a **0** metric delta, and the first numerator movement comes in Phase L.

**Next-tik direction:** iter-02 executes Phase P — prune `billion`'s build cache, push `main` + the 13
`playbill-*` tags, re-pin `.agentspace/rext.tag` to `playbill-m235-hardened`, check the consumption clone
out at it, and verify the M217 FATAL pin guard passes. Expected lift **0**; iter-02 must close honestly as
`closed-fixed` on planned scope with `Gate: NOT MET`. If push is refused for an access or policy reason,
that is a genuine user-blocker and exits the session.

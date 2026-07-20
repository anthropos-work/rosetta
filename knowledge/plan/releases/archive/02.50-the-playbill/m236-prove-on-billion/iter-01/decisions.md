# iter-01 — Decisions

## D1 — the primary metric is `landed-NON-EMPTY / 31`, computed from the canonical manifest

The milestone's exit gate says "every in-scope (session × action)" without naming a denominator. Computed
it from the M233 honesty-gated canonical projection
(`rosetta-extensions/stack-seeding/presets/content-manifest.json`) rather than from prose, so the metric
cannot drift from what the cockpit actually renders.

`has_manager_view` is a **per-session** field, not per-product — an early mis-read at product level
returned `None` for all 4 products and would have under-counted the denominator to 18. Per-session
iteration gives:

- `simulation` 13 sessions × (player + manager) = 26
- `skill-path-legacy` 2 × (player + manager) = 4
- `ai-labs` 2 sessions, **0** landable (presence-only — M231's verdict: nil client, `grade_result` not
  GraphQL-exposed, `/labs/[id]` reads live)
- `skill-path-new` (academy) 1 × player only = 1

**Denominator = 31.** The 2 ai-labs rows are a separate presence assertion, not part of the 31.

## D2 — publish-before-prove: the first tik is a reachability precondition, not gate progress

`billion` consumes `rosetta-extensions` at the tag named in `.agentspace/rext.tag`, and the M217 pin guard
is **FATAL** on mismatch. `billion` pins `casting-call-m228-hiring-scope-fix`; `origin/main` is that same
commit; all 13 `playbill-*` tags are local-only. So the milestone's entire subject-under-test is
unobtainable on the host.

Decision: make publication **iter-02's whole planned scope**, and grade it honestly — `closed-fixed` on
planned scope with `Gate: NOT MET` and a **0** metric delta. Folding it into a "cold bring-up" iter would
have produced exactly the mis-classified close status Phase 4 Step 0 warns about (a big iter that "moved
the metric" only because it also did unrelated setup).

The push itself is the documented workflow, not an escalation: `CLAUDE.md` specifies tools are
"built and tested in the `.agentspace/rosetta-extensions/` authoring copy and tagged, then consumed
per-stack via a pinned-tag clone" — consumption *requires* the tag be on origin. M230–M235 were offline
milestones, so the publish half was simply never exercised. `git push --dry-run` confirms a clean
fast-forward (`1d97861..60eff14`), so this is a mechanical step, not a merge negotiation.

## D3 — host-capacity prune folds into iter-02, it is not its own iter

`billion` has 40 G free with **107.6 GB reclaimable Docker build cache**. A cold UI-tier rebuild needs
headroom. `docker builder prune` is a one-command precondition of the same reachability work as the tag
re-pin, so it belongs in iter-02's planned scope rather than fragmenting into a "cleanup-iter" — it
carries no independent evidence and closing an iter on it alone would be throughput theatre.

## D4 — the coverage/Playthrough harness is authored AFTER the first live render, by design

Cluster 1 of M235's carry-forward is the biggest single work item routed here, and the instinct is to
front-load it. Deliberately not doing that: M235's USER-BLOCKER-M235-02 established that the content-stories
result pages are dynamic-URL + cockpit-seat-reached, so the descriptor's selectors, the mirror-table manager
scoreboard shape, and the per-session score/feedback fence are all **unobservable offline** — authoring blind
ships an *incorrect* (not merely uncalibrated) descriptor into a load-bearing harness.

So the harness is authored in the iter that first lands a real `simulation` result page, against that render.
This is Fate 2 (already planned — Phase L/H of TOK-01), not a deferral.

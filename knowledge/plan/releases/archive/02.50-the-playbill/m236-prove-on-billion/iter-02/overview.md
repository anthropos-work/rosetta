---
iter: 2
milestone: M236
iteration_type: tik
status: closed-fixed
date: 2026-07-20
metric_pre: "0/31"
metric_post: "0/31 (unchanged — reachability precondition, as designed)"
---

# iter-02 — tik: Phase P, publish the v2.5 tooling and re-pin billion

## Step 0 — re-survey (mandatory)

Re-ran the baseline check before committing to TOK-01's named target. Unchanged since iter-01:
`origin/main` is still `1d97861` (the M228 commit), 0 of 13 `playbill-*` tags are on origin, and
`billion`'s `.agentspace/rext.tag` still reads `casting-call-m228-hiring-scope-fix`. The TOK-directed
target is **still untouched and still meaningful** — no substitution needed.

## Active strategy reference

**TOK-01 "publish-then-prove", Phase P.** This tik is Phase P in its entirety.

## Cluster / target identified

The single structural blocker from iter-01's baseline (B2): `billion` cannot obtain the feature under
test. Everything else in the milestone is downstream of this.

## Hypothesis

Publishing `main` + the 13 `playbill-*` tags to origin, then re-pinning `billion`'s
`.agentspace/rext.tag` to `playbill-m235-hardened` and checking its consumption clone out at that tag,
makes the v2.5 content-stories tooling **obtainable** on the host — turning the primary metric from
*unreachable* to *measurable*.

## Expected lift

**0 on the primary metric, by design.** This iter changes reachability, not the numerator. iter-02 must
close honestly as `closed-fixed` on planned scope with `Gate: NOT MET` — folding any seeding or bring-up
work in here to manufacture a delta would be exactly the mis-classification Phase 4 Step 0 warns about.

The *verifiable* outcomes of this iter are binary and host-side:
1. All 13 `playbill-*` tags resolve on origin.
2. `billion`'s consumption clone is at `playbill-m235-hardened` and carries `contentsession/` +
   `presets/content-manifest.json`.
3. The M217 FATAL pin guard passes (pin SoT == clone checkout).
4. Host has headroom for a cold UI-tier rebuild (build cache pruned).

## Phase plan

1. Prune `billion`'s reclaimable Docker build cache (107.6 GB of 109 GB) — precondition for Phase C's
   cold rebuild against 40 G free.
2. `git push origin main` (fast-forward `1d97861..60eff14`, 20 commits) + `git push origin` the 13
   `playbill-*` tags. Purely additive: all tags are ancestors of `main`, 0 name collisions on origin.
3. Update `billion:/home/devops/panorama/.agentspace/rext.tag` → `playbill-m235-hardened`.
4. `git fetch --tags origin && git checkout -f playbill-m235-hardened` in `billion`'s consumption clone.
5. Verify: tags on origin, clone content present, pin guard agreement, host disk headroom.

## Escalation conditions

- **Push refused** (access, branch protection, policy) → genuine user-blocker; the milestone cannot
  proceed. Exit the session.
- **Pruning frees materially less than expected** and disk stays too tight for a UI-tier rebuild →
  route forward a host-capacity item; do not turn this iter into a disk-management project.
- Anything touching platform repos → hard stop (the milestone's 0-platform-edits constraint).

## Acceptable close-no-lift outcomes

If the publish lands but the consumption clone cannot be moved to the tag for a host-side reason, the
falsification ("tag published, host checkout blocked by X") is itself a complete iter outcome and closes
`closed-no-lift` with the blocker routed.

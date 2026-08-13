---
iter: 223
milestone: M257x
iteration_type: tik
status: open
created: 2026-08-09
---

# iter-223 — 23 sha-pinned patches stand between the platform's source and a working demo, and nothing checks they still apply

**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*
**User redirect:** `D-M257x-222-1` — target the platform and the claims made about it; gate clauses 1+2
are the working-stack clauses.

## Step 0 — re-survey (mandatory)

iter-222 left three routes. Re-surveyed at HEAD against the redirect, the one that is *gate-clause-1 work*
rather than corpus work wins:

`ROUTE-M257x-222-pin-advance-needs-a-reproof` asks whether the canonical pin can be advanced past
`app` **+28** / `next-web-app` **+12** / `ant-academy` **+9**. That question decomposes, and one part of
it is decidable **today, for free, without a bring-up**:

> **`demopatch` is the only way the demo changes platform source.** 23 manifests, each pinning a `path` +
> an `anchor` + a `pre_sha256`. `up-injected.sh` applies them to the demo's ephemeral clone just before
> the image build. If an anchor no longer resolves at the ref the demo builds, **the patch is refused and
> the bring-up carries on without it** — which is precisely how *"a silently-refused perf patch shipped a
> 76 s members grid for four releases"* (`build-budget.md`, M255).

**Nothing pre-flights the set.** `demopatch preflight` exists, but it takes **one** manifest and reads the
**checked-out** tree; there is no way to ask *"would advancing the clones break the patch layer?"* before
spending ~11 minutes finding out. That is the gap, and it sits directly on the working-stack clause.

## Cluster / target identified

Every manifest under `rosetta-extensions/demo-stack/patches/*/*.yaml`, resolved against each target
repo's clone at a **named ref** — the pinned `HEAD` the demo builds today, and `origin/main`, which is
what a pin advance would build.

## Hypothesis

The anchors are more durable than the shas, because that is what the mechanism was designed for
(M217: *"the anchor is the contract; the whole-file sha is only a baseline"*). So the census should return
**anchor-clean with substantial sha drift** — and if it does, the sha drift must **not** be reported as a
defect, or the fence contradicts the shipped self-healing gate.

## Pre-registered numeric claims — SEALED in this iter's first commit, before any repair or fence

Derived 2026-08-09 by resolving each manifest's `path` at each repo's `origin/main` and counting
`body.count(anchor)`, the same expression `demopatch` itself uses.

| # | claim | value |
|---|---|---|
| **P1** | demopatch manifests on disk | **23** |
| **P2** | target repos they touch | **4** — app (2) · next-web-app (**11**) · studio-desk (5) · ant-academy (5); **0** touch `platform`. *(Written as `next-web-app (9)` on the first pass and corrected against the projection before sealing — `2+9+5+5 = 21 ≠ 23`, and the arithmetic is what caught it. The two `next-hiring-*` manifests target `next-web-app`; their id prefix names the app inside the monorepo, not the repo.)* |
| **P3** | manifests whose `path` is missing at `origin/main` | **0** |
| **P4** | manifests whose anchor occurs **exactly once** at `origin/main` | **23 of 23** |
| **P5** | manifests whose `pre_sha256` no longer matches the file at `origin/main` | **10 of 23** |
| **P6** | of those 10, how many `demopatch` would REFUSE | **0** — `assert_pre_patch` returns `pristine` on sha drift when the anchor is intact 1×; it WARNs (`demopatch: WARN … but the anchor is intact (1x)`) and applies |

**A census returning zero must prove its instrument** (`§9` iter-149) — P4 is a clean sweep, so the fence
ships with a mutation control that makes each verdict (`anchor gone`, `anchor ambiguous`, `path gone`)
actually fire, and an anti-vacuity control on the manifest population and on an empty anchor.

## Expected lift

No repair is expected — the expected deliverable is a **standing pre-flight**: the question
*"does the patch layer still apply at ref X"* answerable in seconds instead of eleven minutes, and
answered on every suite run instead of never.

## Phase plan

1. Seal the probe (this file + the measured table) as commit 1.
2. Ship the census as a fence, with the sha-drift column advisory by construction.
3. Register it; watch the family go RED on the unregistered member first (the iter-222 control).
4. Run it at both refs; record both answers.
5. Document; close.

## Escalation conditions

- If any anchor is **gone or ambiguous** at `origin/main` → that is a real blocker for the pin advance and
  is reported as one, not repaired by re-authoring the anchor inside this iter.
- If the fence reddens on sha drift → it is wrong, not the manifests; the shipped gate self-heals.

## Acceptable close-no-lift outcomes

P4 measuring below 23 would refute the hypothesis and would be the more valuable finding. Either way the
census is the deliverable.

---
iteration_type: tik
status: closed-fixed
milestone: M257
iter: 09
opened: 2026-08-11
closed: 2026-08-12
gate: MET
---

# iter-09 — L1: multi-stage the two Next images, and the ISOLATION assert that ships with it

**Type:** tik · **Active strategy:** `TOK-02` step 4 — *then levers, unchanged: largest-measured-second
first, one per iter, re-measured at `n ≥ 3`, each landing together with the falsifiable assert that can
trip it — but re-priced for this host.*

## Step 0 — re-survey (mandatory, and it changed one number)

`TOK-02`'s next-tik direction named the baseline campaign; iter-08 took it. The routing then named
**`LEVER-M257-L1-multistage-next`** for iter-09. Re-surveyed at open against the **measured** artefacts
rather than the plan's prose:

| premise | re-checked | verdict |
|---|---|---|
| no app `next.config` sets `output` | `grep -n "output" apps/{web,hiring}/next.config.mjs` → rc=1; repo-wide `output:` in any `next.config.*` → none | **HOLDS** (now verified ×5) |
| `NEXT_PRIVATE_STANDALONE` flips Next 16's frozen `defaultConfig` | read out of the **real installed** Next 16.2.12 inside `demo-1-next-web`: `config-shared.js:129` → `output: !!process.env.NEXT_PRIVATE_STANDALONE ? 'standalone' : undefined` | **HOLDS — measured, not quoted** |
| L1 is worth ~136–152 s here | rep-03's own per-build step table (below) | **HOLDS at 136.4 s**, and it is now a measurement rather than a scaling |

**The estimate is superseded by an attribution.** The overview priced L1 at *"~136–152 s by iter-04's
arithmetic"* — a scaling of billion's phase table by one measured image. iter-08's campaign recorded the
real per-step breakdown, so L1's target is now readable directly off this host:

| build (rep-03, the headroom-clean rep) | total | **`exporting to image`** | of which export | of which unpack | `turbo build` | `pnpm install` |
|---|---|---|---|---|---|---|
| `next-web` | 114.7 s | **70.2 s** | 53.0 | 17.1 | 28.5 | 14.4 |
| `hiring` | 107.5 s | **66.2 s** | 50.1 | 16.0 | 25.0 | 14.9 |
| **both** | 222.2 s | **136.4 s** | 103.1 | 33.1 | 53.5 | 29.3 |

**136.4 s is the export/unpack block L1 attacks, against the 89.51 s the gate needs.** It is 55.4 % of the
two Next lanes and 30.3 % of the whole 449.51 s cycle. The compile legs (53.5 s) and the install legs
(29.3 s) are **not** L1's — they are L4's and L3's respectively, and must not be credited here.

**And the image anatomy says why.** Measured inside `demo-1-next-web` (4.04 GB):

| layer / path | bytes |
|---|---|
| `RUN pnpm install --frozen-lockfile` layer | **2.78 GB** |
| `RUN pnpm turbo build` layer | 300 MB |
| `COPY . .` layer | 46.6 MB |
| `/app/node_modules` in the shipped image | **2,630 MB** |
| `/app/apps/web/.next/server` | 219 MB |
| `/app/apps/web/.next/static` | 22 MB |

The image ships **2.6 GB of dev dependencies to serve a 241 MB build output**, and every byte of it is
exported and unpacked on this host — which is the leg iter-04 proved is real here and the retracted
*"a Mac pays no unpack leg"* claim said was free.

## Cluster / target identified

**L1** — the milestone's largest single lever, and the first one this milestone has been able to price
against a real baseline rather than a scaling. Two Next images, one shape:

- **`hiring`** — already builds from an **rext-owned** `demo-stack/frontend/hiring.Dockerfile`. Edit in place.
- **`next-web`** — builds from the platform clone's `Dockerfile.dev` (`up-injected.sh:879`). A multi-stage
  form cannot land there without a platform-repo edit, so it lands as a **new rext-owned
  `demo-stack/frontend/next-web.Dockerfile`** *in the shape `hiring.Dockerfile` already sanctions*
  (`overview.md` § Hard constraints names exactly this alternative to a demopatch). The clone stays a build
  **context** only.

## Hypothesis

`ENV NEXT_PRIVATE_STANDALONE=1` + `turbo … --env-mode=loose` makes `next build` emit `.next/standalone`;
a second stage copies **only** `standalone/` + `.next/static` + `public/` onto a bare `node:24-alpine` and
runs `node apps/<app>/server.js`. The dev-dependency tree never reaches the final image, so the exported
layer set collapses from ~4 GB to a few hundred MB and the export/unpack leg collapses with it.

**`--env-mode=loose` is not optional** and is the M255 finding this iter must carry: Turbo 2 defaults to
`strict` and filters `NEXT_PRIVATE_STANDALONE` out **before `next build` sees it**, so the flag silently
no-ops and the build stays green with the old full-size image. That failure is silent by construction, so
the builder stage asserts the standalone dir exists rather than trusting it.

## Expected lift

**~120 s** off the cycle (136.4 s of export/unpack → an estimated ~15 s at ~1/10th the bytes), i.e. a p50
of roughly **330 s** against the 449.51 s baseline. Success criterion: **≥ 80 s realized**, which is what
makes the gate's 89.51 s reachable at all. A realized lift under ~50 s would mean the export leg is not
size-proportional here after all — a falsification worth more than the lever, and one that would put the
`re_scope_trigger` (p50 > 400 s after L1+L2+L3) in view early.

## `ASSERT-M257-isolation-with-L1` — lands WITH the lever, never after

The exit gate's second falsifiable clause (**D-v28-11**) — *no built image contains another stack's baked
publishable key or offset origin, asserted by post-build image inspect* — **has no implementation**. `TOK-01`
correctly deferred it (*land each falsifiable assert together with the lever that can trip it*), and it went
unrecorded until iter-07 noticed. **L1 changes exactly the layers that carry those values**, so this is the
iter it belongs to: the multi-stage `COPY --from` is precisely the step that could carry a stale builder
stage's bundle across, and the `demo.patchset` / endpoint / pk reuse-validators all read paths that L1 moves.

Shipped as an executable assert driven from the same evidence the gate names (image inspect + in-bundle
grep), proven able to fail with a negative control.

## Phase plan (declared multi-step — the tripwire counts against THIS shape)

1. **Land L1** — the two rext-owned multi-stage Dockerfiles + the `up-injected.sh` wiring, preserving every
   contract the reuse-validators depend on (the baked `NEXT_PUBLIC_*` ENV, the `demo.patchset` label, and
   the `/app/apps/<app>/.next/static` path the minted-pk probe greps).
2. **Land `ASSERT-M257-isolation-with-L1`** with a negative control that proves it RED.
3. **Measure** — an A/B of the two UI lanes on this host (image size, export leg, unpack leg, total),
   every figure carrying its `load1`; then a cold `n ≥ 3` campaign if the budget allows.
4. **Close** — write the numbers, route the residue.

## Escalation conditions

- **Standalone builds but the app does not serve** (SSR 500s, missing static, dead `/login`) → this is a
  *lever that does not land*, not a user decision: revert the wiring, keep the falsification, close
  `closed-no-lift`.
- **`--env-mode=loose` still filters the var** → the fallback the lever names is a `next.config.mjs`
  demopatch per app; if that is needed, it is a Fate-3 route to iter-10, not a mid-iter third line.
- **A gate tool goes RED on something this iter did not touch** → user-blocker per Phase 5 § 4.

## Acceptable close-no-lift outcomes

A documented falsification that **the export leg on this host is not size-proportional**, or that the
standalone output cannot serve this monorepo's apps, would satisfy the protocol without the metric moving —
either one re-prices the milestone's largest lever and is worth an iter on its own.

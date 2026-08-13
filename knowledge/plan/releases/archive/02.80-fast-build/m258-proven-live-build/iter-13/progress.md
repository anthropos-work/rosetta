# M258 iter-13 — progress

**Type:** tik · **Active strategy:** `TOK-02`, `TIK-A` (Class B — image size, the coupling that is
*favourable*)

Measured 2026-08-12 10:15–10:50Z on `macmini` (Apple M4 Pro, arm64, Docker Desktop VM, **overlayfs**),
`load1` **20.04–30.86** throughout. Every figure below carries that environment.

## Phase A — the Dockerfile

`demo-stack/frontend/studio-desk.Dockerfile`, in the shape `next-web.Dockerfile` (M257 L1) and
`hiring.Dockerfile` (M224) already sanction: **rext-owned**, the clone consumed as a build **CONTEXT
only**, **zero platform-repo edits**, not even a demopatch. Builder installs and builds; runner gets a
fresh `node:24-alpine` and only `dist` + the pruned tree.

**Prune-and-copy, not re-install** (`D61`): `npm prune --omit=dev` in the builder, then
`COPY --from=builder /app/node_modules`. One resolution, one fetch. L1 could use a fresh-install-free
runner because `next build` *emits* `.next/standalone`; studio-desk has no standalone emitter, so the
prune is the equivalent — and this was chosen **because** of the time constraint, not despite it.

## Phase B — the assert fired, and it was the assert that was wrong

The first build **failed**, on this iter's own anti-silence assert. That is the assert working — but the
diagnosis matters (`D65`):

```
=== who requires typescript ===
studio-desk@1.0.0 /app
`-- @clerk/clerk-js@5.125.10
  +-- @base-org/account@2.0.1 -> ox@0.6.9 -> abitype@1.2.4 -> typescript@5.9.3
```

`typescript` is a genuine **production** transitive dependency and cannot leave this image. `vite` was
correctly pruned (`npm ls vite --omit=dev` → `(empty)`). **The prune had worked perfectly; the sentinel
was wrong.** Corrected to `vite`-only, with the reasoning written into the Dockerfile so nobody re-adds
it. *A sentinel must be provably reachable only from devDependencies.*

## Phase C — the measurement, and a prediction refuted

| | old (single-stage) | new (multi-stage) |
|---|---|---|
| **image** | **1.7 GB** | **1.35 GB** |
| dependency layer | `npm ci` **1.04 GB** | pruned `node_modules` **838 MB** |
| source layer | `COPY . .` 61.8 MB | — |
| build output | 63.2 MB | 60.5 MB |

**350 MB per stack (20.6 %)** — against `D62`'s ~1.25 GB and a mid-iter revision of ~670 MB. **Both
refuted by measurement** (`D64`), and the reason is the real finding:

> **studio-desk's image is dominated by PRODUCTION dependencies, not by the toolchain.** 838 MB of the
> 1.04 GB survives `--omit=dev`. `@clerk/clerk-js` carries a crypto-wallet tree — `viem` 68.2 MB,
> `@solana` 20.6 MB, `ox` 9.2 MB, `@base-org` 8.2 MB — plus React-Native/Hermes, all declared as
> *runtime* deps. The single-stage shape was only ~20 % of the problem.

## Phase D — it boots, and it serves identically

A smaller image that 404s is not a win, so this was verified rather than assumed. Container started from
the probe image with the running stack's env, against the **live `demo-1` studio-desk as the control**:

```
route      new (:19500)   control (:19000)
/          302            302
/home      302            302
/assets    302            302
```

302 is the correct healthy response (production `NODE_ENV` mounts `clerkMiddleware()` globally → the
Clerkenstein handshake). Server log: `Server is running on http://localhost:9000`. And the ISOLATION
contract holds — `grep pk_test_ /app/dist` returns the demo's minted
`pk_test_MTI3LjAuMC4xOjE1NDAwJA`, so `buildbench`'s probe still finds its build output at `/app/dist`.

## Phase E — the trap that would have made this ship nothing

`build_frontend_studio_desk`'s reuse check keys on the baked endpoint and the `demo.patchset`
fingerprint. **Neither moves when the Dockerfile changes** — so on every box already holding a
single-stage image, both would match, the image would be reused, and the lever would silently not ship
while the log said success (`D66`). The Dockerfile is now hashed into the fingerprint.
`next-web.Dockerfile` and `hiring.Dockerfile` have the same latent gap — **routed, not fixed here**.

## Phase F — the time axis, withdrawn rather than fudged

`D62` priced the time win at 7–10 s from `build-budget.md`'s **5.73–8.05 s/GB**. That constant was
measured on **`billion` (x86_64, containerd)**; this host is **arm64, overlayfs**. Applying it across
host classes is the precise error `build-budget.md`'s own first rule exists to prevent, and I made it
(`D67`). The available logs do not rescue it either: the new image's cold export+unpack is
**21.6 + 11.6 = 33.2 s**, but the only old-image log is a **warm** re-export (`exporting layers done`,
no duration, 9.6 s unpack). Quoting 33.2 against 9.6 would compare a cold export with a warm one.

**No time claim is made for `TIK-A`.** It is not deferred into the dark: `TIK-C` must run a cold bring-up
anyway, and its phase table gives `ui_studio_desk` cold against iter-05's **115.35 s** baseline for free.

## Phase G — the test gate, with the pre-existing failures PROVEN pre-existing

`test_frontend_build.py`: **105 passed** (from 97 + 8 mid-iter). Three fixture/fence updates were needed
and each grades the *new* invariant rather than weakening the old one — the harness now provisions
`studio-desk.Dockerfile` and a clone `package.json`, and
`test_platform_repo_is_a_build_context_only` asserts all three UI images build from `$HERE/frontend/`
**and `assertNotIn('-f "$ctx/Dockerfile.dev"')`** — no build reads a recipe out of the clone any more.
Both `Dockerfile.dev` stubs are deliberately kept so that assertion cannot pass vacuously.

Full section suite: **9 failed / 1080 passed**. Falsified rather than assumed (`D68`) — `git archive HEAD`
extracted to `.agentspace/rext-pristine-iter13/` at the **same directory depth** (so the live-clone paths
resolve identically) and the four failing files run there:

```
PRISTINE HEAD 5acedd2 (none of iter-13's changes):   9 failed, 131 passed
WITH iter-13's changes:                              9 failed, 1080 passed
```

**The same nine, by name.** All are live-clone / live-container tests; none touch the studio-desk build.

## Close — 2026-08-12

**Outcome:** studio-desk stops being the demo's largest UI image — **1.7 GB → 1.35 GB (350 MB/stack,
20.6 %)**, via an rext-owned multi-stage prune-and-copy Dockerfile with **zero platform-repo edits**,
**verified booting and serving identically to the live control**, and wired so a stale single-stage image
can no longer be silently reused. The headline is real but **one-third of what was predicted**, and the
refutation is the more useful result: **838 MB of the 1.04 GB dependency layer is PRODUCTION deps** —
`@clerk/clerk-js`'s wallet/React-Native tree — so the toolchain was only ~20 % of this image.
**The time half is withdrawn, not claimed**: its prediction came from a constant measured on a different
host class, and `TIK-C`'s cold bring-up settles it for free.
**Type:** tik
**Status:** closed-fixed
**Gate:** N/A — the milestone is achieved by user ruling (`D52`); clause 3 remains NOT MET, unmeasured
under load, and is never to be recorded as met. `TOK-02` (space) sets a goal, not a threshold.
**Phase 5 grading:** (1) gate-met: n *(never, by ruling)* — (2) triggered-tok: n — (3) re-scope: n —
(4) user-blocker: n *(the 9 suite failures are proven pre-existing, so not "a regression in an unrelated
suite")* — (5) cap-reached: n *(1 tik)* — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**

**Decisions:** D64–D68

**Side-deliverables:** none — the cache-key fix (`D66`) is not a side discovery, it is a precondition of
this iter's own deliverable actually shipping.

**Routes carried forward:**

- **`ROUTE-M258-iter13-dockerfile-not-in-cache-key`** (net-new, `D66`) — `next-web.Dockerfile` and
  `hiring.Dockerfile` are not in their own reuse fingerprints, so a tooling-side build change to either
  is silently ignored on any box holding an older image. Same class as the studio-desk gap fixed here.
  Cheapest moment to land it is a cold build, i.e. `TIK-C`.
- **`TARGET-M258-iter13-browser-only-deps-in-the-runtime-image`** (net-new, priced) — the server
  (`src/`) never imports `@clerk/clerk-js`; only the browser `app/` does, and vite inlines it into
  `dist/public`. So `@clerk/clerk-js` **22.8 MB** + `viem` 68.2 + `@solana` 20.6 + `ox` 9.2 +
  `@base-org` 8.2 + `abitype` 2.1 = **131 MB**, plus `react-native` 36.8 + `hermes-compiler` 46.3, are
  dead weight in the runtime image — **~200–260 MB, comparable to or larger than this tik's whole win**.
  Not taken here: deleting packages by hand from a hoisted flat tree is a different risk class from
  letting npm prune by its own graph, and it needs a live serve test per package.
- **`SETTLE-M258-iter13-studio-desk-cold-time`** — `TIK-C`'s cold bring-up yields `ui_studio_desk` cold
  vs iter-05's 115.35 s. Read it there; make no time claim before.

**Lessons:**

- **A prediction derived from a constant measured on another host is not a prediction.** `D62` was mine,
  it was wrong, and `build-budget.md`'s own opening rule says why. *State the environment with every
  number* applies to the numbers you borrow, not just the ones you take.
- **When an assert fires, suspect the assert.** `typescript` survived a prune that had worked; the fix
  was to pick a sentinel provably reachable only from devDependencies.
- **Ask what the cache key covers before shipping a build change.** The lever was one silent image-reuse
  away from landing as a no-op with a green log.
- **Prove "pre-existing", never assume it.** Two runs, same nine names, one from a pristine extract at
  the same directory depth — two minutes to convert an assumption into evidence.

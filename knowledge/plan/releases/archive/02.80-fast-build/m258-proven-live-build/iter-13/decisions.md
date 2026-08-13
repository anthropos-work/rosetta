# M258 iter-13 — decisions

## D64 — studio-desk: 1.7 GB → 1.35 GB, and the win is ~⅓ of what `D62` predicted.

Measured on `macmini`, `load1` 20–31, build args copied verbatim from the running `demo-1` image:

| | old (single-stage) | new (multi-stage) |
|---|---|---|
| image | **1.7 GB** | **1.35 GB** |
| dependency layer | `npm ci` **1.04 GB** | pruned `node_modules` **838 MB** |
| source layer | `COPY . .` **61.8 MB** | — (not in the runtime image) |
| build output | 63.2 MB | **60.5 MB** (`dist` only) |

**Saving: 350 MB per stack (20.6 %).** `D62` predicted ~1.25 GB and even the mid-iter revision said
~670 MB. Both were wrong, and the reason is the finding: **studio-desk's image is dominated by PRODUCTION
dependencies, not by the toolchain.** Of the 1.04 GB `npm ci` layer, **838 MB survives `--omit=dev`** —
`@clerk/clerk-js` carries a full crypto-wallet tree (`viem` 68.2 MB, `@solana` 20.6 MB, `ox` 9.2 MB,
`@base-org` 8.2 MB) plus React-Native/Hermes, all declared as *runtime* dependencies. The single-stage
shape was only ~20 % of the problem.

The lever is still worth shipping — 350 MB × every stack, permanent, zero platform edits — but it must be
quoted at 350 MB and never at `D62`'s number.

## D65 — The assert fired on a prune that had WORKED. `typescript` is not a valid sentinel.

The first anti-silence assert tested `typescript` and `vite` and failed the build. Diagnosis:
`npm ls typescript --omit=dev` resolves a real **production** path —
`studio-desk → @clerk/clerk-js → @base-org/account → ox → abitype → typescript@5.9.3` — so typescript
cannot leave the image without breaking the production graph. `vite` was correctly gone
(`npm ls vite --omit=dev` → `(empty)`).

**The assert was wrong, not the prune.** Corrected to a `vite`-only sentinel, with the reasoning recorded
in the Dockerfile so the next person does not re-add typescript. The general rule: *a sentinel must be
provably reachable only from devDependencies* — and this is the third time this release that a routed fix
was a hypothesis rather than an instruction.

## D66 — The Dockerfile had to go INTO the cache key, or the lever would silently not ship.

`build_frontend_studio_desk`'s reuse check keys on (a) the baked `VITE_GRAPHQL_ENDPOINT` and (b) the
`demo.patchset` fingerprint. **Neither moves when the Dockerfile changes.** So on any box holding an image
built by the old single-stage file — every box that has run a demo — both checks would match, the image
would be reused, and the multi-stage lever would not ship while the bring-up logged success. That is
verbatim the "applied is not shipped" class the fingerprint was created to kill, wearing a new costume.

Fixed by passing `"$dockerfile"` to `next_web_patchset_fp`, which already hashes arbitrary files. The
harness mirrors it by DRIVING the shipped function with the real file (never re-typing the algorithm).

⚠️ **`next-web.Dockerfile` and `hiring.Dockerfile` have the identical latent gap** and are not yet in
their fingerprints — routed as `ROUTE-M258-iter13-dockerfile-not-in-cache-key`. Not fixed here: it would
force a rebuild of both Next images on every box, which is `TIK-C`'s cold build to spend, not this tik's.

## D67 — `D62`'s TIME prediction was derived from a constant measured on a different host, and is withdrawn.

`D62` priced the time win at **7–10 s** from `build-budget.md`'s **5.73–8.05 s/GB** export/unpack figure.
That constant was measured on **`billion` — x86_64, containerd**. This host is **arm64, overlayfs**, a
different snapshotter with a different unpack cost, and `build-budget.md`'s own first rule is that the
same Dockerfile yields 4.84 GB on `billion` and 2.88 GB on an arm64 machine. **I applied a host-specific
constant across host classes — the exact error that rule exists to prevent.**

The one datum taken here does not settle it either: the new image's cold export+unpack was
**21.6 s export + 11.6 s unpack = 33.2 s** for 1.35 GB, but the only old-image log available
(`stacks/demo-1/build-studio-desk.log`) shows `exporting layers done` with **no duration** — a warm
re-export — beside a 9.6 s unpack. **Comparing 33.2 s against 9.6 s would be comparing a cold export with
a warm one**, which is `D53`'s trap in a stopwatch: two numbers for the same thing means the definitions
differ.

**So the time axis is recorded as UNMEASURED on this host**, and no time claim is made for `TIK-A`. It is
not deferred into the dark: `TIK-C` must do a cold bring-up anyway, and its phase table yields
`ui_studio_desk` cold against iter-05's **115.35 s** baseline **for free**. Settle it there — the iter-09
lesson (settle retroactively / from work already required) rather than spend a contended window now.

## D68 — The 9 suite failures are pre-existing and were PROVEN so, not assumed.

The full section suite is **9 failed / 1080 passed**. All nine are live-clone or live-container tests
(`test_ant_academy`, `test_demopatch::TestRealManifest`, `test_ssr_origin_chain::…AgainstLiveClone`,
`test_migrate_race_live`), none touch the studio-desk build, and iter-13's diff is three files
(`up-injected.sh`, `tests/test_frontend_build.py`, the new `frontend/studio-desk.Dockerfile`).

Assuming that would have been enough is how this milestone has mis-graded before, so it was falsified the
way iter-06 did it: `git archive HEAD` extracted to `.agentspace/rext-pristine-iter13/` — **the same
directory depth**, so the tests resolve the same live clones — and the same four files run there. See
`progress.md` for the verdict. `test_frontend_build.py` itself is **105 passed**, up from 97 + 8 failures
mid-iter.

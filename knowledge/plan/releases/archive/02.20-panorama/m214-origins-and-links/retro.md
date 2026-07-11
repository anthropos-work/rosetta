# M214 — retro (origins & links)

## Summary
The v2.2 "panorama" CORS + cross-surface-links + patch-tail milestone (∥ with the now-closed M213). It admits the
MagicDNS/HTTPS origin **everywhere a browser→backend or cross-surface call is gated**, closes the cross-surface link
ejects, and lands the **bounded platform-family patch tail** via the EXISTING rext sha-pinned mechanism — all flipped
by **one scheme predicate** (`browser_scheme` in `gen_injected_override.py`, mirrored by `$SCHEME` in
`up-injected.sh`/`ant-academy.sh`) that the M212 knob drives: `http` on the localhost default (byte-identical), `https`
for a dotted MagicDNS host. Delivered across 8 sections: the `CORS_EXTRA_ORIGINS` https-MagicDNS trio (kept alongside
the localhost trio), the studio-desk `CLERK_SIGN_IN_URL`/`WEB_APP_URL` host+scheme substitution, the
`VITE_CLERK_SIGN_IN_URL` gitignored-overlay bake (which also fixes the un-offset `:3000` for **every** demo), the NEW
`apply-ant-academy-dev-origins.sh` + `ant-academy-dev-origins` patch (env-var indirection admits the MagicDNS host into
`next dev`'s `allowedDevOrigins` at a fixed post_sha256), the scheme-flip confirming the two shipped demopatches + the
mixed-content check, and the NEW `corpus/ops/demo/tailscale-serve.md` remote-access recipe. All gated on
`--public-host`, byte-identical when unset. Config/generation + unit tests only — the live cross-machine run is M215.
Closed clean on the first close pass — merged `--no-ff` into `release/02.20-panorama`.

## Incidents This Cycle
None. No P1/P2, no flakes (gate 5/5 both suites: stack-injection 147p/8s ×5, demo-stack 383p ×5, 0 failures), no
regressions. Build + harden Pass 1 (+5 shell-behavioral tests, 0 bugs) + 3 close adversarial scenarios (ADV-1..ADV-3)
surfaced zero defects in the production code.

## What Went Well
- **One predicate, no drift.** Every host-bearing browser surface (CORS origins, studio-desk redirects, all the baked
  `NEXT_PUBLIC_*`/`VITE_*` endpoints, the cache-validators, the content-URL rewrite, the cockpit deep-links) derives
  its scheme from a single `browser_scheme`/`$SCHEME` predicate — so localhost stays byte-identical and a public host
  flips http→https uniformly. No site can drift out of step, and the byte-identity contract held under the flake gate.
- **The required patch rides the existing mechanism with a fixed post_sha256.** ant-academy's `allowedDevOrigins`
  varies per bring-up (the host is dynamic), but a demopatch pins an EXACT `post_sha256`. Env-var indirection
  (`allowedDevOrigins` reads `ANT_ACADEMY_ALLOWED_DEV_ORIGIN`, prepended to the kept originals) makes the SOURCE change
  static while the host arrives at `next dev` launch — so drift-refuse + idempotency stay exact, and the patch is
  behaviour-identical (upstream-safe) when the env var is unset. Zero platform-repo edits held.
- **The VITE bake gap closed without a Dockerfile ARG.** studio-desk declares no `ARG VITE_CLERK_SIGN_IN_URL` (adding
  one is a platform edit) and its `.dockerignore` excludes `.env*`. A gitignored `.env.production.local` overlay +
  a transient `!.env.production.local` re-include, both trap-reverted on the build's RETURN, baked the offset+scheme
  value — the same apply→build→revert contract as next-web's overlay. Bonus: it fixes the un-offset `:3000` for EVERY
  demo, not just public ones.
- **The conditional item was decided with evidence, not deferred.** next-web `urls.ts` `WEB_APP_URL`/`HIRING_APP_URL`
  would prod-eject IF traversed — but the M42e+M42m coverage sweeps gate at 0 prod-ejects and never surfaced them
  (only `STUDIO_URL` + `PUBLIC_WEBSITE_URL`, both fixed). Documented residual (D-URLS-1), not a speculative demopatch —
  and re-confirmed still-correct in the close deferral audit.
- **Harden targeted the right surface.** The Python emitter was already 99% (the flat, uncoverable `__main__`
  ceiling); the declared TOP RISKS are shell-behavioral (trap-revert, drift-refuse), so harden Pass 1 deepened those
  two shell surfaces directly — leaving the close adversarial review with the traps already pinned.

## What Didn't
- **A rext README index row can't be added at milestone close.** Neither `stack-injection/README.md` (its apply-script
  table) nor `demo-stack/patches/README.md` (its "Shipped manifests" list) indexes the new
  `apply-ant-academy-dev-origins.sh` helper + the `ant-academy-dev-origins` patch. Because the rext code-of-record is
  FROZEN at tag `panorama-m214` (a rosetta-only close must not re-point it), the fix routes to close-release rather
  than being papered over — the same class as M212's D-CLOSE-1 + M213's D-CLOSE-2. Low severity: both rext READMEs are
  already illustrative (one canonical example each) with no numeric count claim, so this is an index-completeness row,
  not a drift.
- **The python3.14 tooling friction recurred (non-blocking).** The box's default `python3` (3.14) has no usable
  pytest; the suites + JUnit run under `python3.12`, as at M212/M213. No milestone impact; noted again for the next
  close.

## Carried Forward
- **D-CLOSE-3** — the rext READMEs don't index the new `apply-ant-academy-dev-origins.sh` helper + the
  `ant-academy-dev-origins` patch → **v2.2 close-release** (bundled with M212's D-CLOSE-1 + M213's D-CLOSE-2 in the one
  rext commit when rext re-tags). Fate 2, in-release.
- **Live cross-machine acceptance + the loopback-vs-0.0.0.0/offset serve reconciliation + cert renewal (90-day LE) +
  RAM/swap burn-down + the cockpit-HTTPS-fronting (7700) polish** → **M215** (the exit-gate milestone). Fate 2,
  confirmed-owned.
- **The two items M212/M213 routed INTO M214** (DEF-M212-01 CORS/Clerk-URL emission; DEF-M213-01 the recipe + link
  admission) **landed complete as Fate-1 here** — the ledger is discharged, not carried.

## Metrics Delta
- **Tests:** stack-injection 152→**155** (+3: `browser_scheme` + the https CORS trio + scheme-flipped redirects);
  demo-stack 367→**383** (+16 net-new funcs: the ant-academy patch apply/revert/drift-refuse/error paths + the
  studio-desk VITE overlay/`.dockerignore` trap-revert + never-clobber). JUnit-authoritative, both exit 0. M214 touched
  zero Go (rext Go stays **1771**) and zero TS.
- **Coverage:** `gen_injected_override.py` 99% statements (flat; uncoverable `__main__`); the two shell trap/refuse
  surfaces deepened behaviorally at harden Pass 1.
- **Flake:** 0 (5/5 gate, both suites). **Supply-chain:** 0 net-new deps. **Platform-repo edits:** 0. **rext close
  edits:** 0 (frozen tag `panorama-m214` @ `99c86b7`).
- Full machine-readable record: `metrics.json`.

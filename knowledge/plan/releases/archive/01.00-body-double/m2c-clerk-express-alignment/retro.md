# M2c — Retro

**Summary:** Brought the last un-gated Clerk consumer — **`@clerk/express`** (studio-desk's Node backend) —
under the alignment framework, **100%/100%** on a 3rd DNA (`clerk-express-1.json`, 9 genes), **closed-on-gate**.
The defining risk (the RS256 wall) was solved the **additive** way: an RS256 path (RSA keypair + real JWKS +
RS256 minting) that the *genuine* `@clerk/express`/`@clerk/backend` accepts networkless via `jwtKey` — **no
HS256 migration**, so M1 (22/22) + M2 (9/9) stayed green throughout. `@clerk/express` is *verified, not
reimplemented* (no mock dir): we satisfy the real SDK like `clerk-webhook/` does with `svix`.

## Incidents This Cycle
- **None during the build** — the crux (real `@clerk/backend` accepts our RS256 token) verified first-try
  (no `iss`/`azp` tuning needed); 0 build-phase bugs, score arc 0 → gate in one runner tik after the crux.
- **Close-review (P2, both fixed):** (1) the `expressrun` runner binary wasn't gitignored (M2c added it);
  (2) an adversarial-review **latent flake** — the bad-signature tamper (`valid[:len-3]+"AAA"`) could be a
  no-op since the sig tail varies with `iat`/`nbf` → fixed to `tamperSig` (deterministic flip) + a regression test.

## What Went Well
- **Additive-first (TOK-01) paid off** — the feared M1/M2 re-gating never happened; the RS256 path slotted
  in beside the HS256 seams. The placement risk the user accepted (M2c-D3) didn't materialize.
- **The crux-first iteration** (prove the real SDK accepts our token before building the full runner) de-risked
  the whole milestone in iter-04; the runner was then mechanical.
- **The svix discipline generalized** — "verify against the genuine library" worked a 2nd time (svix → @clerk/express).

## What Didn't
- **The express gate has a Node dependency** (`@clerk/express` via `EXPRESSRUN_NODE_PATH`) — the only gate
  that isn't pure-Go, so it can't yet run in the (pure-Go) CI without a Node setup + the dependency present.
- The design's "new `clerk-express/` seam" became an additive distribution (no dir) once we realized
  `@clerk/express` is a *consumer* to satisfy, not a library to reimplement — a design assumption corrected at build time.

## Carried Forward
- **Wire the express gate into CI** (`alignment/.github/workflows/alignment.yml`) — needs Node + `@clerk/express`
  in the CI environment → **v1.1 demo-stack** (M3+), where studio-desk's `node_modules` is set up. Recorded in `metrics.json`.

## Metrics Delta (from metrics.json)
- Gate: **MET** (100%/100%, 9/9). Tests: 113 → **128** funcs (122 test + 6 fuzz) across **8** packages.
  Coverage: clerk-backend 97.2%, shared 95.9% (rsa.go added), expressrun 49.1% (mirror gate-covered).
  Flakes: 0. All four gates green (Go 22/22, JS 9/9, Express 9/9, drift ALL PASS).

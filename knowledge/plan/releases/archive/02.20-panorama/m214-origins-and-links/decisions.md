# M214 — decisions

_(implementation choices with rationale, recorded as the milestone is built)_

## D-PATCH-1 — the patch tail rides the existing sha-pinned mechanism
The one required platform-family change (ant-academy `allowedDevOrigins`) and the conditional one (next-web
`urls.ts`) go through the rext `apply-*.sh` / demopatch surface (drift-refuse, idempotent, non-fatal), never a raw
clone edit. **Why:** the platform stays read-only (CLAUDE.md hard rule); drift-refusal makes an upstream change fail
loudly by design.

## D-SCHEME-1 — one predicate flips the whole emission (http localhost / https MagicDNS), per-port
The emission M212 wired-but-deferred is flipped by a single predicate — `browser_scheme(host)` in
`gen_injected_override.py` + the mirrored `SCHEME` var in `up-injected.sh`/`ant-academy.sh`: `http` when the host
is `localhost`/unset (byte-identical to pre-M214), `https` for a dotted MagicDNS FQDN. **Per-port, NOT port-less
443:** M213's `tailscale serve` fronts each browser-facing offset port with HTTPS on the SAME offset port
(D-PROXY-2), so the browser origin is `https://$HOST:<offsetport>`. **Why one predicate:** every host-bearing
site (CORS origins, studio-desk redirects, all NEXT_PUBLIC_*/VITE_* bakes, the cockpit deep-links, the content-URL
rewrite) derives its scheme from it, so localhost stays byte-identical and a public host flips http→https
uniformly — no site can drift. The cache validators (`want_ep`) embed `$SCHEME` too (D-REBUILD-1 extended: an
http-baked image is correctly stale under an https host). The cockpit's OWN serving URL stays http — it is NOT in
M213's `tailscale serve` front list, so it serves plain-http on its port; an http launcher page → https demo
surfaces is fine (a nav/POST from http to https is not mixed content). Fronting the cockpit (7700) too is a
live-acceptance polish for **M215**, not M214's CORS/links scope (disjoint from M213's frozen proxy component).

## D-VITE-SIGNIN-1 — the VITE_CLERK_SIGN_IN_URL bake via a gitignored overlay, not a Dockerfile ARG
studio-desk's `Dockerfile.dev` declares no `ARG VITE_CLERK_SIGN_IN_URL`, so `config.ts`'s SPA sign-in redirect
falls back to the un-offset `http://localhost:3000/login`. **Declaring the ARG is a platform-repo edit
(forbidden); a naive build-context `.env` is dropped by studio-desk's `.dockerignore` (`.env*`).** Resolution: a
gitignored `.env.production.local` overlay at the context root (vite production mode loads it into
`import.meta.env`; `.env.*.local` matches studio-desk's gitignore so the clone never shows it untracked),
admitted past the `.env*` exclusion by a TRANSIENT `!.env.production.local` re-include appended to the
`.dockerignore` — both reverted on the RETURN trap (clone left git-clean), the same apply→build→revert contract
as next-web's transient `.dockerignore` + `.env.local`. Baked at `$SCHEME://$HOST:3000+offset/login` — **fixes
the un-offset `:3000` default for EVERY demo** (not just public) and is https for a public host. **Why not the
Dockerfile:** keeps the "unmodified platform Dockerfile" property + zero platform edit.

## D-URLS-1 — next-web `urls.ts` WEB_APP_URL/HIRING_APP_URL: DOCUMENTED RESIDUAL (evidence-decided, not deferred)
The conditional item. `WEB_APP_URL`/`HIRING_APP_URL` are `NEXT_PUBLIC_NODE_ENV` ternaries → prod
(`app./hiring.anthropos.work`) with no per-URL override, so IF a demo flow traversed them it would prod-eject.
**Investigated (Fate-1 discipline, not a default deferral) — decided with evidence they are NOT a new demopatch:**
1. **The coverage sweeps gate at 0 prod-ejects and did NOT surface them.** `coverage-protocol.md` (:53,66,94,108)
   gates M42e (employee) + M42m (manager) at `0 prod-eject escapes`, classifying every `<a href>` host. The two
   escapes it DID surface — `STUDIO_URL`, `PUBLIC_WEBSITE_URL` — are both fixed (demopatches). `WEB_APP_URL`/
   `HIRING_APP_URL` were never surfaced ⇒ the demo's target flows don't render them as off-demo links.
2. **The `apps/web` usages are non-demo surfaces:** PUBLIC marketing chrome (`PublicHeader`/`PublicFooter`/
   `BlackFridayBanner` — anonymous-only, the demo personas are authenticated enterprise users), PDF/SEO metadata
   (non-navigation), the Clerk.provider `HOSTING_URL` FALLBACK (dead — the demo BAKES `NEXT_PUBLIC_HOSTING_URL`),
   and HIRING-product features (share-sim/invite/start-sim — the demo is a Workforce demo, not recruiting).
3. **The remote dimension doesn't change reachability:** the same authenticated surfaces are browsed remotely; and
   under HTTPS-everywhere the prod hosts are https-prod (not mixed-content), only an eject on flows the demo never
   exercises.
`coverage-protocol.md` makes "add a demopatch mirroring `next-web-studio-url`" a **re-scope trigger only when a
0-eject sweep surfaces the escape** — it hasn't. Adding one speculatively for an unrendered link is gold-plating.
**Decision: documented residual** (recorded here + in `tailscale-serve.md`); if a future coverage sweep surfaces
one of these hosts, the fix is a demopatch mirroring `next-web-studio-url` (the mechanism is proven and ready).

## KB findings from Phase 0b (YELLOW — tracked for the Document phase)
Full audit: `kb-fidelity-audit.md` (2026-07-11). Verdict YELLOW — no blind area that isn't already a declared
`Delivers →` deliverable, no stale load-bearing claim. Every doc gap below is an M214 doc DELIVERABLE, authored
in the Document phase (not a pre-existing stale claim to correct up-front).
- **KB-1** (blind area, expected): `corpus/ops/demo/tailscale-serve.md` does not exist — it IS the milestone's
  declared `Delivers →` deliverable. Author it in the Document phase (the remote-access recipe).
- **KB-2** (completeness): `rosetta_demo.md` has no `--public-host` coverage; `clerkenstein.md` covers the M213
  cert/dotted-pk but not the CORS / `allowedDevOrigins` / per-port-serve origin shape; `frontend-tier.md` CORS
  example is localhost-only + no `allowedDevOrigins` patch note. All three are declared M214 doc updates.
- **KB-3** (carry-forward from M213 Phase 0b KB-4): `frontend-tier.md` "Browser-trusted FAPI cert (M31)" callout
  still describes the pre-M213 mkcert/`127.0.0.1` world — reconcile in the Document phase.

## Adversarial review (close, 2026-07-11)

Scenarios simulated against the M214 rext surface (frozen at `panorama-m214`). All handled by existing
code/tests — recorded here (the scenario, not the fix) so future reviewers see what was considered. No code fix
needed (nothing was broken; the rext tag stays put).

- **ADV-1 — `ACADEMY_ROOT` resolves wrong (patch target absent).** `ant-academy.sh` derives
  `ACADEMY_ROOT="$(cd "$ACADEMY/.." … || echo "$DEV/ant-academy")"`; if `$ACADEMY/..` can't `cd`, the fallback
  path may not hold `code/next.config.js`. **Handled:** `apply-ant-academy-dev-origins.sh` asserts
  `[ -f "$TARGET" ]` and exits non-zero, and `ant-academy.sh` invokes it as `if … apply … ; then … else log ⚠
  (non-fatal)` — a bad root just logs the documented residual and the academy still launches. Pinned by the
  new `test_apply_missing_target_is_refused` clone-independent test.
- **ADV-2 — concurrent studio-desk overlay in the SHARED clone.** Two demo-N builds writing
  `.env.production.local` + the transient `!.env.production.local` `.dockerignore` re-include into the same
  studio-desk context could race. **Handled / pre-existing:** this is the SAME shared-clone serial-build model
  as next-web's established `.env.local` overlay (not an M214-introduced risk); the RETURN trap reverts both on
  every build exit (proven by `test_studio_desk_failed_build_still_reverts_overlay_and_dockerignore`), and a
  pre-existing overlay is never clobbered (`test_studio_desk_does_not_clobber_a_preexisting_env_production_local`).
- **ADV-3 — `browser_scheme` on a dotted non-MagicDNS host (e.g. `127.0.0.1`).** `browser_scheme("127.0.0.1")`
  returns `https` (truthy + `!= "localhost"`), which would be wrong for a bare IP. **Handled upstream:** the
  bring-up contract requires a dotted MagicDNS FQDN and M213's `require_dotted_host` / `assertValidPublishableKey`
  gate rejects a misuse; `browser_scheme` is downstream of that guard. The documented, supported input is a
  MagicDNS name (see `tailscale-serve.md` TL;DR "must be a dotted MagicDNS FQDN"). Not an M214 code defect.

## D-CLOSE-3 — M214's new rext helper + patch are not indexed in the rext READMEs → close-release (Fate-2, rext-frozen)
The NEW `stack-injection/apply-ant-academy-dev-origins.sh` helper and the
`demo-stack/patches/ant-academy-dev-origins/` manifest are not listed in the rext index docs
(`stack-injection/README.md`'s apply-script table + `demo-stack/patches/README.md`'s "Shipped manifests"). The
rext code-of-record is FROZEN at annotated tag `panorama-m214` @ `99c86b7`; editing a rext README now re-points
that tag, which is **`/developer-kit:close-release`'s** job (it bumps the box-level `.agentspace/rext.tag` +
reconciles the rext READMEs). **Fate 2 → close-release**, bundled with M212's D-CLOSE-1 (demo-stack test-count
drift) + M213's D-CLOSE-2 (stack-injection index gap) in the single rext commit close-release makes at the rext
re-tag. NB: both rext READMEs are already ILLUSTRATIVE (each documents ONE canonical example — `apply-authn.sh` /
`next-web-studio-url` — and omits the other pre-existing helpers/patches), and carry **no numeric count claim**, so
this is an index-completeness row to add at the re-tag, not a NEW count-drift. Not a repeat deferral (a distinct
item from D-CLOSE-1/-2, same destination — the designed per-milestone rext-freeze pattern). Provenance:
`audit-deferrals/deferral-audit-2026-07-11.md`.

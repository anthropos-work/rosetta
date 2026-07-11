# M213 — decisions

_(implementation choices with rationale, recorded as the milestone is built)_

## D-SCHEME-1 — HTTPS everywhere, one MagicDNS origin
Front the whole browser surface with the `tailscale cert` via a reverse proxy; teammates hit one `https://<magicdns>`.
**Why:** user decision (2026-07-11) + Clerk clerk-js needs a secure context (Web Crypto), which a plain-`http://`
MagicDNS origin is not — so HTTPS on the app origin is effectively required, not cosmetic.

## D-CERT-1 — tailscale cert, not mkcert, for the remote case
`tailscale cert` (Let's Encrypt) is trusted tailnet-wide with no per-machine CA install — precisely what mkcert
cannot give a remote browser. Consumer mount is path-only, so it is a drop-in at the same `/certs/fapi.{crt,key}`.

## D-PROXY-1 — reverse proxy = `tailscale serve`, NOT Caddy (implementation choice, the overview's open question)
Front the browser-facing ports with **`tailscale serve`**, not a bundled Caddy. **Why:**
1. **Zero net-new dependency** — the decisive factor. `state.md` sets a "0 net-new deps" supply-chain target and
   flags "a reverse-proxy component, if not a stdlib/OS package, is the one supply-chain item to weigh in M213."
   `tailscale serve` ships inside the `tailscale` CLI that is ALREADY on every target VM (Tailscale is the whole
   release's premise; `tailscaled` runs there). Caddy would be a new binary/container to fetch, pin, and trust.
2. **Auto-uses the node's `tailscale cert`** — no cert file to hand `tailscale serve`; it terminates TLS with the
   same Let's Encrypt MagicDNS cert the FAPI mounts (D-CERT-1). One cert, one trust story.
3. **Proven in-corpus** — `corpus/ops/staging-bringup.md` §7 already documents the exact live pattern
   (`sudo tailscale serve --bg https://<host>.taildc510.ts.net http://localhost:3000`), and `staging-clerk.md`
   documents the Secure/SameSite cookie-drop-over-HTTP root cause this HTTPS layer fixes.

## D-PROXY-2 — the serve model built in M213 (per-port HTTPS on the MagicDNS host; live reconciliation → M215)
The generator emits per-port `tailscale serve --bg --https=<offsetport> http://127.0.0.1:<offsetport>` entries for
the browser-facing PLAINTEXT services {app 3000, cosmo 5050, backend 8082, studio-desk 9000, academy 3077}+offset
— the model that PRESERVES the offset-port URL scheme (M214's URL emission is then just http→https, same port).
**FAPI (5400+offset) is excluded from the proxy** — it serves its OWN TLS with the mounted tailscale cert
(D-CERT-1 / scope item a), reachable directly at `https://<host>:5400+offset` (the pk target). Built + unit-tested
here; **NOT stood up live** (no tailnet host in the build env — the wiring is gated + non-fatal, warns/skips when
`tailscale` is absent). The genuinely-live-only reconciliations — the docker/native listener loopback-vs-0.0.0.0
port-binding that makes the per-port serve conflict-free, and whether to collapse further to a literal single
port-less `https://<host>` via 443 path-routing — are **M215's declared scope** ("the live cross-machine
acceptance"): Fate-2, not a new deferral (see §"M215 live burn-down" below).

## D-PK-1 — dotted-host validation moves UP into the demo wiring (refines M212's D-IMPL-1)
M212's D-IMPL-1 held that `inject.py` "stays permissive" — it mints ANY non-`$` host, and the dotted-host
rule is @clerk/backend's downstream concern. M213 REFINES this (not contradicts): a dotless public host would
bake a pk that 500s **every** request (the `dotless-pk-rejected` gene), so failing at first-request runtime is
a bad operator experience. **The split:** the CODEC `mint_pk` STAYS permissive (unchanged — a faithful mirror
of `clerk-frontend/key.go::MintPublishableKey`, which the alignment runner drives with a dotless host on
purpose; the `test_mint_matches_clerkenstein_source` drift guard still passes). The new validation lives in
the demo-WIRING entry `inject.py::require_dotted_host` (called from `main()`), plus an early clear-message
guard in `up-injected.sh` (because inject.py's stderr is suppressed at the mint call). So the minter is still
permissive; the wiring now fails loud on a dotless `--public-host` before any work. KB-3 (a MagicDNS accept
gene) is covered by the codec round-trip + `require_dotted_host` tests + the existing dotted-`127.0.0.1`
express-gate accept gene (same "has-a-dot" predicate — a multi-label host is not a distinct validator path);
a new express gene is NOT added because the express alignment gate is not runnable in the build env (no
`@clerk/express` node_modules) — an unverifiable golden would be lower quality than the code-verified predicate.

## D-TOPO-1 — topology guard: FAPI same host as app; the PSL claim, VERIFIED + corrected
The fake FAPI stays on the SAME host as the app (different port). `up-injected.sh` asserts
`[ "$HOST" != "$FAPI_HOST" ]` ⇒ exit 1 on the public path — equal-by-construction today (both from the one
STACK_PUBLIC_HOST knob), a regression tripwire for a future split. **PSL VERIFIED 2026-07-11** (fetched
`publicsuffix.org`): the entry is **`ts.net`** (+ `*.c.ts.net`), NOT `*.ts.net` — so `ts.net` is the public
suffix and **`taildc510.ts.net` is the registrable domain (eTLD+1)**. This CORRECTS spec-notes' pre-verification
guess ("split needs SameSite=None; Secure"): two subdomains under `taildc510.ts.net` are SAME-SITE (SameSite=Lax
already satisfied). The real blocker to a split is the HOST-ONLY handshake cookie (no `Domain` ⇒ exact-host),
which a split would need `Domain=taildc510.ts.net` to cross. Same-host sidesteps it — keep it. spec-notes
Cookies/topology updated with the verified finding.

## D-REBUILD-1 — build-rebuild guard already trips on a HOST change (CONFIRMED, no new code)
Scope item (In) "confirm the M212 build-rebuild guard trips on a HOST change" is a CONFIRM, and it holds: the
M211 frontend reuse guard compares the cached image's baked `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` /
`VITE_GRAPHQL_ENDPOINT` against `want_ep="http://$HOST:$((5050+OFFSET))/graphql"` — which EMBEDS `$HOST`. So a
localhost→MagicDNS change flips `want_ep`, the cached localhost-baked image mismatches → `docker image rm -f`
+ rebuild (re-baking the new pk too, since pk derives from the same host). Already TEST-COVERED by M212's harden
pass: `test_next_web_tag_guard_rebuilds_when_host_knob_makes_localhost_endpoint_stale`,
`test_next_web_tag_guard_reuses_when_host_and_offset_both_match`, `test_studio_desk_tag_guard_rebuilds_on_host_change`
(all green; 80/80 frontend-build). No new code — confirmed + re-verified under M213's diff.

## D-EGRESS-1 — fake-fapi cdn.jsdelivr.net egress made explicit + testable + operator-overridable
The fake-fapi PROXIES the instance-agnostic clerk-js bundle from `cdn.jsdelivr.net` (`server.go`
handleClerkJSBundle) — its ONE outbound-egress dependency. "Confirm egress" done properly: turned the
hardcoded, untested `"https://cdn.jsdelivr.net"` into `clerkJSCDNBase()` + a `defaultClerkJSCDN` const +
a `FAKE_FAPI_CLERKJS_CDN` override — so (a) unit tests point it at an httptest server (no real network),
(b) an operator behind a jsdelivr block can point at a mirror. +4 clerk-frontend Go tests (proxies the
configured CDN with forwarded path/query + JS MIME; default is jsdelivr; override trims trailing slash; 502
clerkjs_proxy on egress failure). up-injected.sh adds a NON-FATAL host-side egress pre-check (gated on the
public host — a laptop localhost demo already has internet) that warns early with the mirror-override hint;
+1 static pin. Behavior-preserving: JS/FAPI alignment gate re-scored **100%/100% (9/9)**. The live
container→CDN reach is M215 (needs the running stack).

## KB findings from Phase 0b (YELLOW — tracked for the Document phase)
- **KB-1** (blind-area / missing xref, medium): the reverse-proxy / one-HTTPS-origin-over-MagicDNS topology is
  covered by EXISTING corpus docs not referenced by the milestone — `corpus/ops/staging-bringup.md` §7 (the
  `tailscale serve` pattern) + `corpus/ops/staging-clerk.md` (MagicDNS origins + the Secure/SameSite cookie-drop
  root cause). READ both before the proxy work (done). Document-phase: add a short proxy-topology note to
  `clerkenstein.md`; confirm M214 owns the `tailscale-serve.md` recipe (Fate-2, already in M214's `In:`).
- **KB-2** (thin anchor, low-med): `alignment_testing.md` cites but never NAMES the `dotless-pk-rejected` gene
  (real in `alignment/dna/clerk-express-1.json`). Document-phase: add a one-line gene reference there (or in
  `clerkenstein/knowledge/alignment.md`).
- **KB-3** (gene completeness, low): the `dotless-pk-rejected` gene's ACCEPT case uses dotted loopback
  `127.0.0.1`, not a multi-label MagicDNS host. Evaluate adding a MagicDNS accept-case to the DNA (inside scope
  item b "verify the dotted MagicDNS host passes").
- **KB-4** (docs going stale, medium): `recipe-browser-login.md` §B step 2 (remote = untrusted cert /
  proceed-anyway), `frontend-tier.md` "Browser-trusted FAPI cert (M31)" callout, and
  `clerkenstein/knowledge/architecture.md:150-152` all describe the pre-M213 mkcert/`127.0.0.1`-only world.
  Document-phase update list.
- **KB-5** (line-anchor drift, low): fixed in `spec-notes.md` (symbol anchors). See its "Line-anchor drift" block.

## M215 live burn-down (Fate-2 — M215 already owns "the live cross-machine acceptance")
Surfaced during M213 build, confirmed covered by M215 (no new deferral, no plan edit needed):
- Execute `tailscale cert` + `tailscale serve` LIVE on billion; browse the demo end-to-end from a remote tailnet
  machine (the cert-swap + proxy config M213 generates + unit-tests).
- Resolve the port-binding so the per-port serve is conflict-free live (docker/native loopback-bind vs a serve
  port scheme) + decide whether to collapse to a literal single `https://<host>` (443 path-routing).
- Cert renewal (90-day LE) + RAM/swap burn-down (per overview "Live foundation").

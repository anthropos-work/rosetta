---
active_release: "v2.2 panorama — IN DEVELOPMENT (branch release/02.20-panorama; tag v2.2 at close); designed 2026-07-11"
active_branch: "release/02.20-panorama"
active_milestone: "(between milestones) — M213 auth-over-tailnet CLOSED+merged 2026-07-11; next: M214 origins-and-links (not yet started) → M215 (final, iterative)"
last_closed: "v2.1 quick change — 2026-07-09 (tag v2.1, 4 milestones M208..M211) — the skiller-in-app re-ground | latest milestone: M213 — 2026-07-11"
phase: "in development — M212 + M213 closed+merged; next: M214 → M215"
last_updated: "2026-07-11"
---

# State

**Active release:** **v2.2 "panorama"** — the **external-shareability release**: make dev/demo stacks reachable from
other machines on a **Tailscale** tailnet (run a stack on a Tailscale VM, e.g. `billion.taildc510.ts.net` on the
odyssey Proxmox host; a teammate with Tailscale up browses the demo end-to-end). Designed 2026-07-11 via
`/developer-kit:design-roadmap`; branch `release/02.20-panorama` cut from `main`; tag **`v2.2`** at close. **4
milestones M212 → { M213 ∥ M214 } → M215** (+ optional/deferrable **M216**). **Decisions:** HTTPS-everywhere under
one MagicDNS origin (Clerk needs a secure context); external access **opt-in, default off** (an explicit
`/demo-up --public-host` flag); **demo-first** (dev-path parity = optional M216). **Tooling + docs + an opt-in flag
only — zero platform-repo edits** (two platform-family files ride the EXISTING rext sha-pinned patch mechanism). The
**sanctioned re-proposal** of the v1.4 seed "external stack shareability (Tailscale/ingress)" dropped 2026-06-11.
Records: [`releases/02.20-panorama/`](releases/02.20-panorama/).

**Last shipped:** **v2.1 "quick change" — 2026-07-09, tag `v2.1`** (rext code-of-record `v2.1` =
`quick-change-m211`; `.agentspace/rext.tag` = `v2.1`). The **skiller-in-app re-ground** — the merged platform stands
up cold on both stacks via the re-grounded tooling. 4 milestones M208..M211. Records:
[`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).

**Active milestone:** **(between milestones)** — **M213 — Auth over the tailnet** (`section`) **CLOSED 2026-07-11**,
merged `--no-ff` into `release/02.20-panorama`; the `m213/auth-over-tailnet` branch is deleted. It serves Clerkenstein
auth + the whole browser surface under **one trusted HTTPS MagicDNS origin**: the `tailscale cert` FAPI mint swap
(path-only drop-in + local mkcert/openssl fallback), dotted-pk validation in the demo wiring, the NEW
`gen_tailscale_serve.py` reverse-proxy generator (per-port HTTPS, 0 net-new deps), the FAPI-same-host topology guard,
the confirmed build-rebuild-on-HOST guard, and the overridable `cdn.jsdelivr.net` egress (`FAKE_FAPI_CLERKJS_CDN`) —
all gated on `--public-host`, byte-identical when unset. rext code-of-record FROZEN at tag `panorama-m213` @ `b9f41dd`
(rext re-tag deferred to close-release). Close: go clerk-frontend +7 / stack-injection **152** / demo-stack **367**,
flake 5/5, 7 findings (0 must-fix; 1 rext-README residual → close-release, D-CLOSE-2), deferral audit GREEN. **Next
active:** **M214 — Origins & links** via `/developer-kit:build-milestone`.

**Phase:** **in development — 2 of 4 milestones closed (M212, M213).** The release branch carries the M212 + M213
merges; the knob foundation + the HTTPS-over-tailnet auth surface are in. The remaining contracts (M214 origins/links,
M215 the live cross-machine acceptance) are scaffolded under `releases/02.20-panorama/`; the feasibility is established
(workflow `wf_bea3be47` — config-only core + a 2-item patch tail; the `tailscale cert` swap proven live on billion).

**Next up:** **M214 — Origins & links** (`section`, ∥-ready with the now-closed M213, different files) — admit the
MagicDNS origin everywhere gated + close cross-surface ejects + land the bounded patch tail via the EXISTING rext
patch mechanism: extend `CORS_EXTRA_ORIGINS` emission, studio-desk runtime redirects + the `VITE_CLERK_SIGN_IN_URL`
bake gap, a NEW `apply-*.sh` for ant-academy `allowedDevOrigins`, and the conditional next-web `urls.ts` demopatch.
**Delivers →** the NEW `corpus/ops/demo/tailscale-serve.md` remote-access recipe. Run `/developer-kit:build-milestone`.
Then **M215** (the FINAL, iterative cross-machine acceptance gate on odyssey — the live `tailscale cert`/`serve` run +
cert-renewal + RAM burn-down + the loopback-vs-0.0.0.0 serve reconciliation). Optional **M216** (dev-path parity)
stays roadmap-only until promoted. Two M213-surfaced Fate-2 items are owned here/M215; the rext-README residual
(D-CLOSE-2) + M212's D-CLOSE-1 land at close-release when rext re-tags.

**Origin sync (2026-07-11):** origin has `main` + all shipped tags (`v1.10.1`, `v2.0`, `v2.1`) for **both** rosetta
and rosetta-extensions, **and** the `release/02.20-panorama` branch (pushed 2026-07-11, `765528d`). **Local ahead of
origin (push at release close):** the rosetta `release/02.20-panorama` now carries the **M212 + M213 merges** (+ their
close commits); rext has **two new local annotated tags — `panorama-m212` @ `770f81b` and `panorama-m213` @ `b9f41dd`**
(the per-milestone code-of-record; the box-level `.agentspace/rext.tag` stays `v2.1` until `/developer-kit:close-release`
bumps it + reconciles the D-CLOSE-1/D-CLOSE-2 rext-README residuals). No push performed by the close.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), CAVEAT-1 (clean-box literal full `/dev-up` — belt-and-suspenders), M314b (prod
frozen-read hydration — a prod-team follow-up). Reserved **Playthroughs futures** M205–M207 stay in vision. All
tracked in [`roadmap-vision.md`](roadmap-vision.md).

## Recently shipped releases
- **v2.1 "quick change"** — **2026-07-09**, tag `v2.1`. The **skiller-in-app re-ground** (M208..M211). Records:
  [`releases/archive/02.10-quick-change/`](releases/archive/02.10-quick-change/).
- **v2.0 "opening night"** — **2026-07-02**, tag `v2.0`. The **Playthroughs** pillar (M201..M204); 10 live
  Playthroughs GREEN on cold reset-to-seed. Records:
  [`releases/archive/02.00-opening-night/`](releases/archive/02.00-opening-night/).
- **v1.10b "fit-up"** — **2026-07-01**, tag `v1.10.1`. Field-hardening backfill (M47..M53). Records:
  [`releases/archive/01.10b-fit-up/`](releases/archive/01.10b-fit-up/).

_(Earlier v1.x — v1.0 … v1.10 — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.)_

## Headline numbers (v2.2 in progress — baseline v2.1 final; M212 + M213 closed)
- **rext Go test funcs:** **1771** across 6 modules (v2.1 baseline 1764 + **7** M213 `clerkjs_proxy` funcs; M212
  touched zero Go). `go vet ./...` + `gofmt` clean; clerk-frontend `-race`+shuffle clean.
- **rext Python (M213-touched suites):** stack-injection **152** (144p/8s), demo-stack **367** (0f) — the M213 delta
  is +33 build + 11 harden test funcs. **TS unit specs:** **103**. **Live Playthroughs:** **10** + 1 in-manifest TODO.
- **Flake:** **0** (5/5 gate). **Supply-chain:** **0 net-new deps held** — M213 chose `tailscale serve` (already on
  every target VM) over Caddy for the reverse proxy (D-PROXY-1). **Alignment gates:** **100%/100%** on all 5
  Clerkenstein surfaces — M213 re-scored the JS/FAPI surface 100%/100% after the egress refactor (M215 re-scores live).

## Branch model / shipped tags
**v2.2 IN DEVELOPMENT:** `release/02.20-panorama` cut from `main` 2026-07-11. Milestone branches `m{N}/{slug}` cut
from it per `/developer-kit:build-milestone`; merge back at `/developer-kit:close-milestone`; the release merges →
`main` + tags `v2.2` at `/developer-kit:close-release`.
**Shipped tags:** **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` ·
**v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` ·
**v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-11 (M213 "auth over the tailnet" CLOSED + merged into `release/02.20-panorama` — Clerkenstein
auth + the whole browser surface under one trusted HTTPS MagicDNS origin: the `tailscale cert` FAPI swap + dotted-pk
validation + the `gen_tailscale_serve.py` reverse-proxy generator + the topology guard + the overridable jsdelivr
egress, all gated on `--public-host`, byte-identical when unset; rext code-of-record @ tag `panorama-m213` @ `b9f41dd`;
go +7 / stack-injection 152 / demo-stack 367 / flake 0; 7 close findings (0 must-fix, 1 rext-README residual →
close-release D-CLOSE-2); deferral audit GREEN. Next: `/developer-kit:build-milestone` → M214 → M215.)_

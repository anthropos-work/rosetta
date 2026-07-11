---
active_release: "v2.2 panorama — IN DEVELOPMENT (branch release/02.20-panorama; tag v2.2 at close); designed 2026-07-11"
active_branch: "release/02.20-panorama"
active_milestone: "M215 prove-on-odyssey — the FINAL, iterative cross-machine acceptance gate (planned; not yet started). M212+M213+M214 closed+merged"
last_closed: "v2.1 quick change — 2026-07-09 (tag v2.1, 4 milestones M208..M211) — the skiller-in-app re-ground | latest milestone: M214 — 2026-07-11"
phase: "in development — M212 + M213 + M214 closed+merged (3 of 4); next: M215 (final, iterative)"
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

**Active milestone:** **(between milestones)** — **M214 — Origins & links** (`section`) **CLOSED 2026-07-11**, merged
`--no-ff` into `release/02.20-panorama`; the `m214/origins-and-links` branch is deleted. It admits the MagicDNS/HTTPS
origin everywhere a browser→backend or cross-surface call is gated, closes the cross-surface link ejects, and lands
the bounded platform-family patch tail via the EXISTING rext sha-pinned mechanism: the `CORS_EXTRA_ORIGINS`
https-MagicDNS emission (`browser_scheme` + the offset https trio), the studio-desk `CLERK_SIGN_IN_URL`/`WEB_APP_URL`
host+scheme substitution, the `VITE_CLERK_SIGN_IN_URL` gitignored-overlay bake (trap-reverted; also fixes the
un-offset `:3000` for every demo), the NEW `apply-ant-academy-dev-origins.sh` + `ant-academy-dev-origins` patch
(env-var indirection admits the MagicDNS host into `next dev`'s `allowedDevOrigins` at a fixed post_sha256), the
`$SCHEME` flip confirming the two shipped demopatches + the mixed-content check, and the NEW
`corpus/ops/demo/tailscale-serve.md` recipe — all gated on `--public-host`, byte-identical when unset. The conditional
next-web `urls.ts` landed as the evidence-decided **documented residual** (D-URLS-1). rext code-of-record FROZEN at tag
`panorama-m214` @ `99c86b7` (rext re-tag deferred to close-release). Close: stack-injection **155** (147p/8s) /
demo-stack **383** (0f), flake 5/5, 4 findings (0 must-fix; 2 docs fixed; D-CLOSE-3 → close-release), deferral audit
GREEN. **Next active:** **M215 — Prove it on odyssey** (the FINAL, iterative acceptance gate).

**Phase:** **in development — 3 of 4 milestones closed (M212, M213, M214).** The release branch carries the M212 +
M213 + M214 merges; the knob foundation + the HTTPS-over-tailnet auth surface + the origins/links/patch tail are all
in. The one remaining contract (M215 — the live cross-machine acceptance) is scaffolded under
`releases/02.20-panorama/`; the feasibility is established (workflow `wf_bea3be47` — config-only core + a 2-item patch
tail; the `tailscale cert` swap proven live on billion).

**Next up:** **M215 — Prove it on odyssey** (`iterative`, complexity **large**, the FINAL v2.2 milestone) — the live
cross-machine acceptance gate on the odyssey `billion.taildc510.ts.net` VM: `/demo-up N --public-host …` (opt-in) →
a teammate on a DIFFERENT tailnet machine completes a full employee AND manager hero journey with 0 localhost/prod
ejects, 0 CORS blocks, 0 cert-untrusted, 0 mixed-content, assets rendering, on a cold reset-to-seed (unset knob =
byte-identical). Its iteration loop fixes any live breakage back into the M212/M213/M214 surfaces (never a platform
edit). Run `/developer-kit:work-mstone-iters` (or `build-mstone-iters`). Then `/developer-kit:close-release` (rolls
the rext `v2.2` tag, bumps `.agentspace/rext.tag`, reconciles the rext-README residuals D-CLOSE-1/-2/-3, merges →
`main` + tags `v2.2`). Optional **M216** (dev-path parity) stays roadmap-only until promoted. Four Fate-2 items are
owned by M215 (live acceptance + loopback-vs-0.0.0.0/offset reconciliation + cert renewal + RAM + the cockpit-fronting
polish); three rext-README index residuals (D-CLOSE-1/-2/-3) land at close-release when rext re-tags.

**Origin sync (2026-07-11):** origin has `main` + all shipped tags (`v1.10.1`, `v2.0`, `v2.1`) for **both** rosetta
and rosetta-extensions, **and** the `release/02.20-panorama` branch (pushed 2026-07-11, `765528d`). **Local ahead of
origin (push at release close):** the rosetta `release/02.20-panorama` now carries the **M212 + M213 + M214 merges**
(+ their close commits); rext has **three new local annotated tags — `panorama-m212` @ `770f81b`, `panorama-m213` @
`b9f41dd`, and `panorama-m214` @ `99c86b7`** (the per-milestone code-of-record; the box-level `.agentspace/rext.tag`
stays `v2.1` until `/developer-kit:close-release` bumps it + reconciles the D-CLOSE-1/-2/-3 rext-README residuals). No
push performed by the close.

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

## Headline numbers (v2.2 in progress — baseline v2.1 final; M212 + M213 + M214 closed)
- **rext Go test funcs:** **1771** across 6 modules (v2.1 baseline 1764 + **7** M213 `clerkjs_proxy` funcs; M212 +
  M214 touched zero Go). `go vet ./...` + `gofmt` clean.
- **rext Python (v2.2-touched suites):** stack-injection **155** (147p/8s), demo-stack **383** (0f) — the M214 delta
  is +3 stack-injection + 16 demo-stack net-new funcs (CORS emission tests + the ant-academy patch apply/revert/drift
  paths + the studio-desk VITE overlay trap). **TS unit specs:** **103**. **Live Playthroughs:** **10** + 1 TODO.
- **Flake:** **0** (5/5 gate, both v2.2 suites). **Supply-chain:** **0 net-new deps held**. **Alignment gates:**
  **100%/100%** on all 5 Clerkenstein surfaces (M215 re-scores live).

## Branch model / shipped tags
**v2.2 IN DEVELOPMENT:** `release/02.20-panorama` cut from `main` 2026-07-11. Milestone branches `m{N}/{slug}` cut
from it per `/developer-kit:build-milestone`; merge back at `/developer-kit:close-milestone`; the release merges →
`main` + tags `v2.2` at `/developer-kit:close-release`.
**Shipped tags:** **v2.1** `v2.1` · **v2.0** `v2.0` · **v1.10b** `v1.10.1` · **v1.10** `v1.10` · **v1.9** `v1.9` ·
**v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5` · **v1.3b** `v1.3.1` · **v1.3** `v1.3` ·
**v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`. (Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-11 (M214 "origins & links" CLOSED + merged into `release/02.20-panorama` — the MagicDNS/HTTPS
origin admitted everywhere a browser call is gated + the cross-surface ejects closed + the bounded patch tail via the
existing sha-pinned mechanism [CORS https-trio emission + studio-desk redirects + the VITE overlay bake + the NEW
ant-academy allowedDevOrigins patch] + the NEW corpus/ops/demo/tailscale-serve.md recipe, all gated on --public-host,
byte-identical when unset; rext code-of-record @ tag panorama-m214 @ 99c86b7; stack-injection 155 / demo-stack 383 /
flake 0; 4 close findings [0 must-fix, 2 docs fixed, D-CLOSE-3 → close-release]; deferral audit GREEN. Next: M215, the
final iterative live-acceptance gate.)_

---
title: "Deferral Audit — M214 origins-and-links (close)"
date: 2026-07-11
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (no single item deferred across ≥2 milestones of v2.2).
- No chronic / drift patterns.
- No aged-out items (every record dated 2026-07-11; the release is 2 days old; no destination milestone has
  closed; M214's own area — origins & links — was just built).
- The **two** items prior milestones routed **into M214** both **LANDED complete as Fate-1** in this milestone
  (CORS/Clerk-URL emission; the `tailscale-serve.md` recipe + link admission). Every remaining open item is a
  clean Fate-2 routing to a v2.2 milestone that already owns it (M215 live acceptance; close-release rext
  re-tag), each re-verified this pass against the target's `overview.md`/exit-gate.

## Summary
- Total deferrals in scope: 8 records (2 inherited→M214 RESOLVED · 3 inherited→M215 open · 2 inherited→close-release open · 1 new M214→close-release)
- Single deferrals: 6 open (all confirmed-owned, no user decision required)
- Resolved this milestone (Fate-1 landings of inherited routings): 2
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Evidence-decided non-deferrals reviewed: 1 (D-URLS-1 — confirmed still NOT a deferral)

## Deferral Inventory

### Inherited routings that LANDED in M214 (Fate-1 — removed from the ledger)

```yaml
- id: DEF-M212-01
  item: "gen_injected_override CORS (CORS_EXTRA_ORIGINS) + studio-desk CLERK_SIGN_IN_URL/WEB_APP_URL emission at the MagicDNS/HTTPS origin"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  resolved_in: M214 (gen_injected_override.py browser_scheme + fe_origins https trio; scheme-flipped studio-desk redirects — commit bf3edd1)
  destination: "RESOLVED — landed complete as Fate-1 in M214 (its declared In: emission)"
  partial_attempted: no

- id: DEF-M213-01
  item: "reverse-proxy topology recipe (tailscale-serve.md) + CORS/link emission for the MagicDNS origin"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  resolved_in: M214 (NEW corpus/ops/demo/tailscale-serve.md [198 lines] + the CORS/link/patch tail — commits bf3edd1/ca4cb0b/4599a2d + the docs commit 1d019e9)
  destination: "RESOLVED — landed complete as Fate-1 in M214"
  partial_attempted: no
```

### Inherited, still open — owned by M215 (Fate-2)

```yaml
- id: DEF-M213-02
  item: "execute tailscale cert + tailscale serve LIVE on billion; remote tailnet browser completes the full journey"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  last_seen_in: m213/audit-deferrals/deferral-audit-2026-07-11.md
  destination: "M215 prove-on-odyssey (the exit gate)"
  reason_recorded: "the live cross-machine acceptance is M215's exit gate; the build env has no tailnet host"
  partial_attempted: no

- id: DEF-M213-03
  item: "docker/native loopback-vs-0.0.0.0 per-port serve conflict + offset appears correctly in every baked MagicDNS URL"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  destination: "M215 prove-on-odyssey (overview risk 3 + the iteration protocol)"
  reason_recorded: "live-only reconciliation — surfaces only on a running cross-machine stack"
  partial_attempted: no

- id: DEF-M213-04
  item: "tailscale cert renewal (90-day LE) re-issue+reload + RAM/swap fit on the ballooned billion VM"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  destination: "M215 prove-on-odyssey (overview risks 1 + 4)"
  reason_recorded: "needs the long-lived running stack"
  partial_attempted: no
```

### Inherited + new — the rext-frozen README reconcile, owned by close-release (Fate-2)

```yaml
- id: D-CLOSE-1
  item: "demo-stack README test_tooling count-drift reconcile (M212)"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  destination: "v2.2 close-release (the rext re-tag + box-level rext.tag bump moment)"
  reason_recorded: "rext code-of-record frozen at tag panorama-m212; editing the README re-points the annotated tag"
  partial_attempted: no

- id: D-CLOSE-2
  item: "stack-injection README apply-script index gap (M213 authn/inject row set)"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  destination: "v2.2 close-release (bundled with D-CLOSE-1 in the single rext commit)"
  reason_recorded: "rext frozen at panorama-m213; the index reconcile lands at the rext re-tag"
  partial_attempted: no

- id: D-CLOSE-3   # NEW this close
  item: "rext READMEs don't index M214's new apply-ant-academy-dev-origins.sh helper + the ant-academy-dev-origins patch manifest"
  origin_milestone: M214
  first_deferred_on: 2026-07-11
  last_seen_in: m214/decisions.md D-CLOSE-3 (recorded at close, Phase 7)
  destination: "v2.2 close-release (bundled with D-CLOSE-1/-2 in the one rext README commit at the rext re-tag)"
  reason_recorded: "rext code-of-record FROZEN at annotated tag panorama-m214 @ 99c86b7; editing stack-injection/README.md or demo-stack/patches/README.md now re-points that tag, which is close-release's job. NB: consistent with pre-existing illustrative doc style — both READMEs already document one canonical example (apply-authn.sh / next-web-studio-url) and omit the other pre-existing helpers/patches, so this is no NEW numeric count-drift, only an index-completeness row the rext re-tag will add."
  partial_attempted: no
```

## Repeat-Deferral Patterns
None. No single item has been deferred across ≥2 distinct milestones.

- The two items routed INTO M214 (DEF-M212-01, DEF-M213-01) were deferred once, then **landed** here — the
  opposite of a repeat.
- **D-CLOSE-1 / -2 / -3 are three DISTINCT items** (M212's demo-stack test count · M213's stack-injection
  index · M214's ant-academy patch index) that happen to share a destination (close-release). They are **not the
  same item re-deferred** — so not a repeat by the definition. The shared destination is the **designed**
  rext-freeze pattern: each milestone freezes the rext code-of-record at its own annotated tag
  (`panorama-m212/-m213/-m214`), and the box-level `.agentspace/rext.tag` bump + the accompanying README
  reconcile is structurally close-release's job (it is the skill that re-tags rext + re-points the consumption
  pin). This is benign accumulation, NOT a CHRONIC_DEFER — close-release discharges all three in one rext commit.

## Aging
No aged-out records: all dated 2026-07-11, none ≥3 months, no destination milestone (M215, close-release) has
closed, and no item has been deferred across ≥2 milestones. M214's own feature-area was just built (this close).

## Fate-1 Investigation

### DEF-M212-01 / DEF-M213-01 — CORS/Clerk-URL emission + the recipe/link admission
- **Fate-1 (land now, complete) feasible:** YES — and both were **LANDED complete** in M214. The emission
  (browser_scheme + the https CORS trio + scheme-flipped studio-desk redirects), the two platform-family patches
  (ant-academy `allowedDevOrigins` + the studio-desk `VITE_CLERK_SIGN_IN_URL` overlay), the mixed-content scheme
  flip, and the NEW `tailscale-serve.md` recipe are all in this milestone. Not slices — full landings. Mark
  resolved.

### DEF-M213-02 / -03 / -04 — live acceptance, port-binding/offset, cert-renewal + RAM
- **Fate-1 feasible:** no — the build env has no tailnet host; these surface only on a live cross-machine run.
- **If no:** Fate 2 — **M215** (the FINAL iterative acceptance milestone) owns all three, re-verified against its
  `overview.md`: the exit gate is the live `--public-host billion.taildc510.ts.net` remote journey at
  `0 mixed-content / 0 CORS / 0 ejects / 0 cert-untrusted`; risk 3 names the offset-in-every-baked-URL /
  single-vs-multi-stack reconciliation; risks 1+4 name RAM fit + `tailscale cert` renewal. No plan edit (confirm).

### D-CLOSE-1 / -2 / -3 — rext README reconcile
- **Fate-1 feasible:** no — the rext code-of-record is FROZEN at the per-milestone annotated tags; touching a
  rext README now re-points that tag. **If no:** Fate 2 — close-release owns the rext re-tag + the README
  reconcile (it bumps the box-level `.agentspace/rext.tag`). Correctly routed; no plan edit.

## Evidence-decided non-deferral reviewed (NOT a deferral)

### D-URLS-1 — next-web `urls.ts` WEB_APP_URL / HIRING_APP_URL (documented residual)
Re-examined this pass per the audit's "ask fresh" discipline. **Confirmed: still NOT a deferral.** It is a
Fate-1-*investigated* decision whose complete, correct outcome is "no patch needed — documented residual," decided
with evidence (not for lack of time): the M42e+M42m coverage sweeps gate at 0 prod-ejects and never surfaced these
two hosts (only `STUDIO_URL` + `PUBLIC_WEBSITE_URL`, both fixed); the `apps/web` usages are non-demo surfaces
(public marketing chrome / PDF-SEO metadata / a dead Clerk.provider fallback / hiring-product features). Adding a
speculative demopatch for an unrendered link would be gold-plating; `coverage-protocol.md` makes "add a demopatch"
a re-scope trigger only when a 0-eject sweep surfaces the escape — it hasn't. The reasoning holds. Documented in
`decisions.md` D-URLS-1 + `corpus/ops/demo/tailscale-serve.md` §"Documented residual". If a future sweep surfaces
one of these hosts, the fix (a demopatch mirroring `next-web-studio-url`) is proven and ready. No re-fating needed.

### Cockpit-HTTPS-fronting → M215 (Fate-2)
The presenter cockpit's own page (port 7700+off) is deliberately NOT in M213's `tailscale serve` front list, so it
serves plain HTTP even under a public host (an http launcher → https demo surfaces is not mixed content — a nav/POST
from http to https is fine). Fronting 7700 too is a **live-acceptance polish**, confirmed owned by **M215**: its
iteration protocol is "fix in the M212 knob / M213 TLS / M214 CORS+patch surface" and its exit gate drives the full
journey (which uses the cockpit seat-switch for hero login) to 0 mixed-content. Within M215's iteration remit — no
plan edit.

## Recommendations
- DEF-M212-01 → **LAND-NOW (Fate 1) — already landed complete in M214.** Mark resolved.
- DEF-M213-01 → **LAND-NOW (Fate 1) — already landed complete in M214.** Mark resolved.
- DEF-M213-02 / -03 / -04 → **LAND-NEXT (Fate 2)** — confirmed owned by M215 (exit gate + risks); no plan edit.
- D-CLOSE-1 / -2 → **LAND-NEXT (Fate 2)** — confirmed owned by close-release; no plan edit.
- D-CLOSE-3 → **LAND-NEXT (Fate 2)** — NEW; owned by close-release (rext re-tag), bundled with D-CLOSE-1/-2.
  Record the decision in M214 `decisions.md` at close (Phase 7).

## Applied Changes
- No `overview.md` edits (every open item is a Fate-2 confirm — the destinations M215 / close-release already own
  their items; Fate 2 = confirm, don't edit).
- D-CLOSE-3 is recorded as a new decision in `m214/decisions.md` during close-milestone Phase 7 (the review-fixes
  commit), with this audit as its provenance.
- DEF-M212-01 + DEF-M213-01 marked resolved (landed Fate-1 in M214) — recorded here; no ledger edit beyond this.

## Blocking Items (require user decision)
None. Verdict GREEN — the calling close-milestone proceeds without a user prompt.

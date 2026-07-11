---
title: "Deferral Audit — M213 auth-over-tailnet (close)"
date: 2026-07-11
scope: milestone
invoked-by: close-milestone
---

## Verdict
GREEN

- No repeat deferrals (no item deferred across ≥2 milestones of v2.2).
- No aged-out items (every record dated 2026-07-11; no destination milestone has closed; the one
  feature-area touched by a later milestone — the clerkenstein.md dotted-host doc — was **landed**, not
  re-deferred).
- Every in-scope item is a clean Fate-2 routing to a future v2.2 milestone that **already owns it**
  (confirmed against the target `overview.md`/exit-gate), or was resolved early as Fate-1 in M213.

## Summary
- Total deferrals in scope: 5 open + 1 resolved-early = 6 records
- Single deferrals: 5 (all confirmed-owned, no user decision required)
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Landed early (Fate-1, inherited): 1 (DEF-M212-02)

## Deferral Inventory

### M213's own routings

```yaml
- id: DEF-M213-01
  item: "reverse-proxy topology recipe (tailscale-serve.md) + CORS/link emission for the MagicDNS origin"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  last_seen_in: m213/decisions.md:91 ("confirm M214 owns the tailscale-serve.md recipe (Fate-2, already in M214's In:)")
  destination: "M214 origins-and-links"
  reason_recorded: "different files (CORS/patches vs cert/proxy/pk); the recipe + origin admission are M214's declared In:"
  partial_attempted: no

- id: DEF-M213-02
  item: "execute tailscale cert + tailscale serve LIVE on billion; remote tailnet browser completes the full journey"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  last_seen_in: m213/decisions.md:104-107 (§"M215 live burn-down")
  destination: "M215 prove-on-odyssey (the exit gate)"
  reason_recorded: "no tailnet host in the build env; the config + unit tests are here, the live cross-machine acceptance is M215's exit gate"
  partial_attempted: no  # config generation + unit tests fully landed in M213; the LIVE run is a distinct, owned deliverable

- id: DEF-M213-03
  item: "resolve the docker/native loopback-vs-0.0.0.0 port-binding so the per-port serve is conflict-free live"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  last_seen_in: m213/decisions.md:108-109 (D-PROXY-2 + §"M215 live burn-down")
  destination: "M215 prove-on-odyssey"
  reason_recorded: "genuinely-live-only reconciliation — surfaces only on a running cross-machine stack"
  partial_attempted: no

- id: DEF-M213-04
  item: "cert renewal (90-day LE) re-issue+reload step + RAM/swap burn-down on the ballooned billion VM"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  last_seen_in: m213/decisions.md:110 + overview.md:43-48 ("Live foundation")
  destination: "M215 prove-on-odyssey"
  reason_recorded: "needs the long-lived running stack; M215 owns the live burn-down"
  partial_attempted: no
```

### Inherited from M212 (still open at M213 close)

```yaml
- id: DEF-M212-01
  item: "CORS + Clerk-URL (MagicDNS origin) emission via gen_injected_override.py"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  last_seen_in: m212/audit-deferrals/deferral-audit-2026-07-11.md (LAND-NEXT → M214)
  destination: "M214 origins-and-links"
  reason_recorded: "confirmed covered by M214 (its In: list); not load-bearing for M212 (127.0.0.1 default byte-identical)"
  partial_attempted: no

- id: D-CLOSE-1
  item: "rext-frozen README test-count drift reconcile"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  last_seen_in: state.md:32 + roadmap.md M212 block ("1 routed finding → close-release")
  destination: "close-release (rext re-tag moment)"
  reason_recorded: "rext code-of-record is frozen at panorama-m212; the README count reconcile lands when close-release bumps the box-level rext.tag"
  partial_attempted: no
```

### Resolved early in M213 (Fate-1 — removed from the ledger)

```yaml
- id: DEF-M212-02
  item: "clerkenstein.md dotted-host (assertValidPublishableKey) constraint doc"
  origin_milestone: M212
  first_deferred_on: 2026-07-11
  resolved_in: M213 (corpus/services/clerkenstein.md — the dotless-pk-rejected gene ref in the alignment
    section + the new "Remote HTTPS over the tailnet (v2.2 M213)" section; also rext knowledge
    architecture.md + alignment.md)
  destination: "RESOLVED — landed as Fate-1 in M213's Document phase"
  reason_recorded: "M213 owns the dotted-pk validation surface; documenting the constraint alongside the
    tailscale-cert path was the natural home — landed complete, not partial"
```

## Repeat-Deferral Patterns
None. No item has been deferred in ≥2 distinct milestones. DEF-M212-01 was deferred once (M212 close) to
M214; M214 has not yet run, so it is a **pending single deferral**, not a repeat. It was NOT re-deferred at
M213 (M213 does not own CORS emission).

## Fate-1 Investigation

### DEF-M213-01 — recipe + CORS/link emission
- **Fate-1 (land now, complete) feasible:** no
- **If no:** Fate 2 — M214 "origins & links" already owns it (`overview.md` In: — `CORS_EXTRA_ORIGINS`
  emission, the ant-academy `allowedDevOrigins` apply-*.sh, the studio-desk/next-web links, AND the NEW
  `tailscale-serve.md` recipe). Disjoint file set from M213 (CORS/patches vs cert/proxy/pk); an additive
  merge by design. No plan edit needed (confirm, don't edit).

### DEF-M213-02 / -03 / -04 — live acceptance, port-binding, cert-renewal + RAM
- **Fate-1 feasible:** no
- **If no:** Fate 2 — M215 "prove it on odyssey" (the FINAL, iterative acceptance milestone) owns all three.
  Its exit gate literally is the live cross-machine `--public-host billion.taildc510.ts.net` run on a cold
  reset-to-seed, and its overview names the loopback-bind reconciliation + the `tailscale cert` PEM shape +
  RAM fit as the live-only reconciliations. The build env has no tailnet host — landing these now is
  impossible, not merely out of scope. Config generation + unit tests (the landable half) ARE fully in M213.

### DEF-M212-01 — CORS/Clerk-URL emission (inherited)
- **Fate-1 feasible:** no — M213's file set (cert/proxy/pk) does not include the CORS emitter
  (`gen_injected_override.py:304-306`). Fate 2 — still correctly owned by M214, which has not yet run.

### D-CLOSE-1 — rext README count-drift (inherited)
- **Fate-1 feasible:** no — the rext code-of-record is FROZEN at tag `panorama-m212`/`panorama-m213`; touching
  the rext README now would re-point the annotated tag, which is close-release's job. Fate 2 — close-release
  owns the rext re-tag + the README reconcile. Correctly routed.

### DEF-M212-02 — clerkenstein.md dotted-host (inherited)
- **Fate-1 feasible:** YES — and it was LANDED in M213. M213 owns the dotted-pk validation surface, so the
  constraint doc found its natural home alongside the tailscale-cert path. Resolved complete (not a slice).

## Recommendations
- DEF-M213-01 → **LAND-NEXT (Fate 2)** — confirmed owned by M214; no plan edit.
- DEF-M213-02 → **LAND-NEXT (Fate 2)** — confirmed owned by M215 (exit gate); no plan edit.
- DEF-M213-03 → **LAND-NEXT (Fate 2)** — confirmed owned by M215; no plan edit.
- DEF-M213-04 → **LAND-NEXT (Fate 2)** — confirmed owned by M215; no plan edit.
- DEF-M212-01 → **LAND-NEXT (Fate 2)** — still owned by M214 (not yet run); no plan edit; not a repeat.
- D-CLOSE-1 → **LAND-NEXT (Fate 2)** — owned by close-release; no plan edit.
- DEF-M212-02 → **LAND-NOW (Fate 1) — already landed in M213.** Mark resolved.

## Applied Changes
None required. All open items are pre-existing, correctly-fated Fate-2 routings whose destinations were
re-verified this pass against the live `overview.md`/exit-gate text (M214 In: + M215 exit gate). Fate 2 =
confirm, don't edit. DEF-M212-02 resolved organically in M213's Document phase (no ledger edit needed beyond
this record). No `overview.md` edits, no new `decisions.md` entries.

## Blocking Items (require user decision)
None. Verdict GREEN — the calling close-milestone proceeds without a user prompt.

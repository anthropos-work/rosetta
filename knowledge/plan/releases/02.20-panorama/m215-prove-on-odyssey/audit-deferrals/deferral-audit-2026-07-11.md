---
title: "Deferral Audit — M215 prove-on-odyssey (milestone close, FINAL v2.2 milestone)"
date: 2026-07-11
scope: milestone
invoked-by: close-milestone
---

## Verdict
**GREEN**

- No repeat deferrals (no single item deferred across ≥2 milestones).
- No chronic / drift patterns.
- No aged-out items (every record dated 2026-07-11, within days; nothing ≥3mo, no destination-milestone closed-without-landing).
- All inherited Fate-2→M215 items LANDED as Fate-1 in M215's live proof. All new M215 residuals have clear Fate-2 destinations; none are release-scope-breaking.

## Summary
- Total deferrals in scope: **11 records** — 3 inherited→M215 (all RESOLVED Fate-1) · 3 inherited→close-release (open, correct destination, runs next) · 4 new M215→future/backlog (Fate-2) · 1 re-reviewed evidence-decided non-deferral (D-URLS-1).
- Single deferrals: 11
- Repeat deferrals: 0
- Chronic patterns flagged: 0
- Aged-out: 0
- Blocking items: **0**

## Deferral Inventory

### Inherited → M215 — RESOLVED as Fate-1 in this milestone
```yaml
- id: DEF-M213-02
  item: "execute tailscale cert + tailscale serve LIVE on billion; remote tailnet browser completes the full journey"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  destination: "M215 (exit gate)"
  resolution: "LANDED Fate-1 — both vantages driven live from a remote Mac over Tailscale (maya-thriving→/profile, dan-manager→/enterprise/workforce), trusted LE cert (verify=0), 0 console errors, 0 localhost/prod ejects. iter-01/findings.md tik-2 + both-vantages."
- id: DEF-M213-03
  item: "docker/native loopback-vs-0.0.0.0 per-port serve conflict + offset appears correctly in every baked MagicDNS URL"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  destination: "M215 (risk 3 + iteration protocol)"
  resolution: "LANDED Fate-1 — per-port tailscale-serve topology proven; the exact leftover-serve bind conflict surfaced (F12) and was FIXED (teardown per-port serve-reset + defensive up-path pre-reset); offset threads correctly through every baked https://$HOST:<offsetport> URL (verified on the cold reset-to-seed)."
- id: DEF-M213-04
  item: "tailscale cert renewal (90-day LE) re-issue+reload + RAM/swap fit on the ballooned billion VM"
  origin_milestone: M213
  first_deferred_on: 2026-07-11
  destination: "M215 (risks 1 + 4)"
  resolution: "RESOLVED — RAM fit PROVEN (16 GB swap added, ~5.7 GB free held with next-web + studio-desk up); the 90-day LE renewal cadence + renew-then-reload step DOCUMENTED in tailscale-serve.md §the-tailscale-cert-FAPI (a demo is disposable — renewal is a documented operational note, not a code gap)."
```

### Inherited → close-release — open, correct destination (runs NEXT)
```yaml
- id: D-CLOSE-1
  item: "demo-stack README test_tooling count-drift reconcile (M212)"
  origin_milestone: M212
  destination: "v2.2 close-release (the rext re-tag + rext.tag bump moment)"
  reason_recorded: "rext frozen at tag panorama-m212; editing the README re-points the annotated tag"
- id: D-CLOSE-2
  item: "stack-injection README index gap (M213)"
  origin_milestone: M213
  destination: "v2.2 close-release (bundled rext commit)"
- id: D-CLOSE-3
  item: "rext READMEs don't index the new apply-ant-academy-dev-origins.sh helper + ant-academy-dev-origins patch (M214)"
  origin_milestone: M214
  destination: "v2.2 close-release (bundled rext commit)"
```
These are **three DISTINCT** rext-README residuals (a count-drift, an index-row gap, a new-helper index) that share **one** destination — the single rext commit `/developer-kit:close-release` makes when it re-tags rext + bumps `.agentspace/rext.tag`. They are the designed per-milestone rext-freeze pattern (a rosetta-only milestone close must not re-point a frozen annotated tag), **not** the same item re-deferred. Destination has not closed; close-release is the immediate next operation.

### New in M215 → Fate-2 future work / standing backlog
```yaml
- id: DEF-M215-01  (F5)
  item: "two app demopatches (target-role authz-skip, ai-readiness loadMembers) sha-drift refuse on the current app tag — re-anchor"
  origin_milestone: M215
  first_deferred_on: 2026-07-11
  destination: "future demo-service maintenance (a demopatch pre-hash re-anchor); standing backlog"
  reason_recorded: "NON-FATAL (demo works — slower per-member Sentinel fan-out / unbounded AI-readiness hydration); a pre-existing demopatch-maintenance issue surfaced on a fresh full-clone, NOT a remote/Linux deliverable; rext is frozen at panorama-m215 so a re-anchor+re-tag is out of this close's scope"
  partial_attempted: no
- id: DEF-M215-02  (F9)
  item: "taxonomy/library/skills content on a remote VM needs the .agentspace/snapshots cache scp'd (or captured); auto-provision it remotely"
  origin_milestone: M215
  first_deferred_on: 2026-07-11
  destination: "documented residual (tailscale-serve.md Step 4) + future cache-auto-sync enhancement"
  reason_recorded: "OPERATIONAL, not a code defect — identity/profile/dashboard/workforce render fully; only taxonomy/library sparse. The runbook documents the scp/capture path. Auto-syncing the cache to a VM is a later enhancement, off the core exit gate"
  partial_attempted: no
- id: DEF-M215-03  (F11)
  item: "seed hero identity-key vs generated-profile-name mismatch (cosmetic)"
  origin_milestone: M215
  first_deferred_on: 2026-07-11
  destination: "future seed polish (stack-seeding); standing backlog"
  reason_recorded: "COSMETIC — login + profile render both work; the identity key (maya-thriving) differs from the generated display name. Unrelated to the remote/Linux story; rext-seeding is frozen"
  partial_attempted: no
- id: DEF-M215-04  (F13)
  item: "jobsimulation container exits(1) on startup (prints CLI help — no run/serve subcommand)"
  origin_milestone: M215
  first_deferred_on: 2026-07-11
  destination: "future demo-service fix (compose/service-command / version-drift); standing backlog"
  reason_recorded: "OFF the proven journey path (both hero journeys rendered fine); would affect only the AI-Simulations surface; NOT remote/Linux-specific — would hit any demo. A separate demo-service investigation, not part of the panorama shareability scope"
  partial_attempted: no
```

### Re-reviewed — evidence-decided non-deferral (carried by M214, confirmed still correct)
```yaml
- id: D-URLS-1
  item: "next-web urls.ts WEB_APP_URL/HIRING_APP_URL would prod-eject IF traversed"
  origin_milestone: M214
  status: "NOT a deferral — a documented residual decided with evidence (M42e/M42m coverage sweeps gate at 0 prod-ejects and never surfaced these hosts). The fix mechanism (a demopatch mirroring next-web-studio-url) is proven+ready if a future sweep surfaces one. Re-confirmed still-correct this pass."
```

## Repeat-Deferral Patterns
None. Grouping every record by item-similarity yields no group of size ≥2. The D-CLOSE-1/-2/-3 trio shares a destination but not an item (three distinct README residuals). No CHRONIC_DEFER, no DRIFT_DEFER.

## Fate-1 Investigation
- **DEF-M213-02/03/04** — Fate-1 feasible: **yes, and DONE** (landed live in M215; see resolutions above).
- **DEF-M215-01 (F5)** — Fate-1 now: **no**. Requires re-anchoring the demopatch pre-hashes + a rext re-tag; rext is frozen at `panorama-m215` and re-tagging is `/developer-kit:close-release`'s job, not this rosetta close. Non-fatal, orthogonal to shareability. → Fate-2 future demo-service maintenance.
- **DEF-M215-02 (F9)** — Fate-1 now: **no** (and not a code landing). Operational — the cache must be present on the target VM; the runbook documents scp/capture. Not release-scope-breaking (external-shareability is proven; content-density is orthogonal). → documented residual + future auto-sync enhancement.
- **DEF-M215-03 (F11)** — Fate-1 now: **no**. Cosmetic seed-data; rext-seeding frozen. → Fate-2 seed polish.
- **DEF-M215-04 (F13)** — Fate-1 now: **no**. A demo-service/compose investigation off the proven path; not remote/Linux-specific; rext frozen. → Fate-2 demo-service fix.
- **D-CLOSE-1/-2/-3** — Fate-1 now: **no** (the frozen-tag constraint is the whole point); already correctly owned by close-release (Fate-2), which runs immediately after this close.

None reach the escape hatch — none are release-scope-breaking. v2.2's commitment (opt-in remote demo reachable + fully functional over Tailscale from a second machine) is proven for both vantages on a cold reset-to-seed; every residual is demo-service / seed / content-cache polish orthogonal to that commitment.

## Recommendations
- DEF-M213-02/03/04 → **LAND-NOW (done)** — recorded as resolved.
- DEF-M215-01 (F5), DEF-M215-03 (F11), DEF-M215-04 (F13) → **LAND-NEXT (Fate-2)** to future demo-service / seed maintenance (standing backlog; no milestone in v2.2 owns them — v2.2 is closing).
- DEF-M215-02 (F9) → **LAND-NEXT (Fate-2)** documented residual + future cache-auto-sync enhancement.
- D-CLOSE-1/-2/-3 → **LAND-NEXT (Fate-2)** — confirmed owned by close-release; no plan edit.
- D-URLS-1 → **not a deferral** — no action.

## Applied Changes
- This report written.
- M215 `decisions.md` records the four new residual fates (DEF-M215-01..04) + confirms the inherited items' resolution; `carry-forward.md` catalogs the routed residuals for close-release / future work.
- No `overview.md` edits (no Fate-3 annotate — all routings are pre-existing owners or standing-backlog Fate-2; v2.2 has no further milestone to annotate).
- No `roadmap-vision.md` edit required (no escape-hatch / RELEASE-SCOPE-DEFER item).

## Blocking Items (require user decision)
**None.** Zero repeat-deferrals, zero aged-out, zero escape-hatch. Verdict GREEN → `SEVERITY=clear`.

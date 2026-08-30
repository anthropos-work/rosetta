# `/demo-update` — Out-of-scope / next-release parking lot

> Items parked out of the initial [`spec.md`](spec.md) `v0.1` build. Each carries a trigger for
> reconsideration; nothing here blocks M-A..M-E.

## Parked

- **D1 — Continuous / scheduled `demo-update`.** A cron loop that keeps a demo forever-at-prod-latest.
  *Trigger to reopen:* stable `/demo-update` in the field for ≥1 release cycle + operator ask for
  hands-off refresh (e.g. a permanently-warm demo-N on Ithaca for sales).
- **D2 — Multi-`N` batch update.** `demo-update --all` iterating every live `demo-N`. *Trigger:* ≥3 concurrent
  demos on one box.
- **D3 — Update from `demo-down` state.** Would require a "resurrect" verb (compose network+volume rehydration
  without full re-inject). *Trigger:* operator report of a stopped-but-must-preserve demo-N.
- **D4 — Full schema rollback (auto pg_restore).** Reverse-migrate isn't universally safe (destructive
  migrations exist). *Trigger:* operator demand + a documented reverse-migration convention across services.
- **D5 — `--with-playthroughs` as default.** Playthroughs are expensive (T4). Default remains T1+T2+T3.
  *Trigger:* a fast-lane playthroughs subset (~1 min) exists.
- **D6 — Cross-stack update fan-out (dev + demo together).** Out of scope; different contracts. *Trigger:*
  never (belongs to a coordination layer above the skills).
- **D7 — Update-time platform-repo edits.** Explicitly forbidden. Any surface that can't be driven without a
  platform edit **escalates**, does not edit. Mirrors coverage-protocol / playthroughs.

## Explicitly not deferred — decided out

- Anything modifying `/stack-update` (kept as-is; peer, not extension).
- Anything modifying `/demo-up` or `/demo-down` (this spec strictly adds a third verb).
- Values-reading of `.env` (values-blind is a load-bearing principle; not a deferral, a decision).

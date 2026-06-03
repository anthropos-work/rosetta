# M3 — Decisions

## M3-D1 — per-demo service-repo clones (user-chosen, 2026-06-03)
Each `demo-N` re-clones the platform service repos under its own `anthropos-demo/stacks/demo-N/`, rather than
all demos sharing the single `anthropos-dev/*` clones. **Chosen for full filesystem isolation** + future-proofing
per-demo config divergence, accepting ~N× disk + clone time. Note: since every demo uses the *same* Clerkenstein
injection, contention was not the deciding factor — isolation was. (Alternative: shared clones, disk-cheaper.)

## M3-D2 — manual teardown only (user-chosen, 2026-06-03)
`/demo-down [N]` is the only reclaim path; no nightly auto-reaper in M3. Keeps M3 tight + avoids a teardown-safety
concern. (Alternative: a cron/systemd reaper of `demo-*` older than X — deferred; add later only if forgotten
stacks become a real problem.)

## Open (resolve during build)
- Max-N + exact port-offset sizing (below the ephemeral range, no overlap with the 24 base mappings).
- The `clerk-backend` `api.clerk.com` → fake-BAPI redirect mechanism **inside Docker** (extra_hosts + trusted CA
  vs a base-URL env override). Load-bearing — spike in S3.
- Whether the v1.0 express-gate CI carry-forward lands here vs M5 (default: M5).

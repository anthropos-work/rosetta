# M215 — decisions

_(strategy + implementation choices with rationale; recorded at close for the direct-drive iterative milestone —
the canonical run record is [`iter-01/findings.md`](iter-01/findings.md).)_

## D-STRAT-1 — drive M215 directly (not via background sub-agents)
M215 is **live shared-infrastructure work**: a single real Linux VM (`billion`, odyssey Proxmox host) reached over
Tailscale, with one deploy user and shared serve/cert state. Parallel background actors would collide on the VM's
ports, `tailscale serve` config, and Docker state. So the milestone was driven **directly and serially**: a trimmed
tik-1 (backend + Clerkenstein only, `DEMO_NO_UI=1`) to de-risk the clone + amd64 Go builds + migrate + Clerkenstein
FAPI + `tailscale cert`/`serve` **before** adding the UI tier (tik-2) for the actual remote browser login. The
iter-01 findings.md is the running record; progress.md + this file are backfilled at close.

## Inherited Fate-2 → M215 — all resolved
- **DEF-M213-02** (execute cert+serve LIVE; remote browser completes the journey) → **LANDED Fate-1**: both vantages
  driven live from a remote Mac over Tailscale, trusted LE cert (`verify=0`), 0 ejects, 0 console errors.
- **DEF-M213-03** (loopback-vs-0.0.0.0 per-port serve conflict + offset in every baked URL) → **LANDED Fate-1**: the
  per-port `tailscale serve` topology proven; the exact leftover-serve bind conflict surfaced (F12) and was FIXED.
- **DEF-M213-04** (90-day LE cert renewal + RAM/swap fit) → **RESOLVED**: RAM fit proven (16 GB swap; ~5.7 GB free
  held with the UI tier); the renewal cadence documented in `tailscale-serve.md` (a demo is disposable).

## Residual fates (new in M215 — Fate-2, documented + routed; none release-scope-breaking)
The full audit is [`audit-deferrals/deferral-audit-2026-07-11.md`](audit-deferrals/deferral-audit-2026-07-11.md)
(GREEN); the durable routing record is [`carry-forward.md`](carry-forward.md).

## DEF-M215-01 (F5) — two `app` demopatches sha-drift refuse → demopatch re-anchor
`app: target-role authz-skip` + `app: ai-readiness loadMembers bound` demopatches refused on the current `app` tag
(pinned pre-hash drifted). **NON-FATAL** — the demo works (slower per-member Sentinel fan-out / unbounded
AI-readiness hydration). A pre-existing demopatch-maintenance issue surfaced on a fresh full-clone, **not** a
remote/Linux deliverable; rext is frozen at `panorama-m215`. **Fate-2** → future demo-service maintenance
(standing backlog). v2.2 has no remaining milestone to own it.

## DEF-M215-02 (F9) — remote VM taxonomy/library needs the snapshot cache → documented residual
After migrate, `public.skills=0` and the seed logs `taxonomy=skipped(cache-miss)` — the taxonomy (~42,790 public
skills) + Directus content are set-dressed from `.agentspace/snapshots`, which a fresh VM lacks. **Operational, not
a code defect**: identity/profile/dashboard/workforce render fully; only taxonomy/library are sparse. The runbook
documents the scp/capture path (`tailscale-serve.md` Step 4). **Fate-2** → documented residual + a future
cache-auto-sync enhancement. Off the core exit gate (both hero journeys rendered without it).

## DEF-M215-03 (F11) — seed hero identity-key vs profile-name mismatch (cosmetic)
Logged in as `maya-thriving`; nav shows "Maya" but the profile person renders as a different generated name. A
seed-data naming inconsistency (identity key ≠ generated display name); **login + render both work**. **Fate-2** →
future seed polish (stack-seeding, frozen here). Unrelated to the remote/Linux story.

## DEF-M215-04 (F13) — jobsimulation container exits(1) on startup
Its binary printed CLI help + exited (no run/serve subcommand). **OFF the proven journey path** (both hero journeys
rendered fine); would affect only the AI-Simulations surface. Likely a service-command/compose or version-drift
issue that would hit **any** demo — **not** remote/Linux-specific. **Fate-2** → separate demo-service fix (standing
backlog), not part of the panorama shareability scope.

## Adversarial review

### ADV-1 — F12 up-path defensive pre-reset ordering vs its stated failure mode
**Scenario:** the F12 fix adds two guards — (a) `tailscale serve` per-port reset on **teardown** (`/demo-down`),
and (b) a **defensive up-path pre-reset**. The up-path pre-reset sits **after** `docker compose up`. Its comment
says a stale serve listener means "the new backend container can't bind `0.0.0.0:<offsetport>`" — but that bind
happens at compose-up, *before* the pre-reset runs. So if the real conflict were the backend bind, the up-path
pre-reset could not prevent the first compose-up failure.
**Assessment:** **not a defect that blocks close.** The **teardown** reset (the primary mechanism) covers the normal
`/demo-down` → `/demo-up` cycle — after a clean teardown there is no leftover serve at the next up. The up-path
pre-reset is a belt-and-suspenders for a serve-*reconfiguration* conflict; the by-hand-teardown case (where a stale
listener survives) is exactly what the runbook's manual `tailscale serve reset` note (Step 7) covers. The block is
non-fatal and a localhost no-op. rext is **frozen** at `panorama-m215`, so no code change is made here.
**Fate:** a one-line comment/placement reconciliation is routed to the **rext close-release re-tag** (Fate-2,
bundled with the D-CLOSE-1/-2/-3 rext work when rext next advances) — never a re-point of the frozen tag.

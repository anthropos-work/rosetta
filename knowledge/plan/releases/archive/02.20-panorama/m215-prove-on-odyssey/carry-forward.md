---
title: "M215 Carry-Forward — Routes from prove-on-odyssey"
date: 2026-07-11
status: archived
close_status: closed-on-gate
gate_target: "remote demo reachable at https://<magicdns>; a teammate on a DIFFERENT tailnet machine completes a full employee AND manager hero journey with 0 localhost/prod ejects, 0 CORS blocks, 0 cert-untrusted, 0 mixed-content, assets rendering — reproducibly on a cold reset-to-seed; unset knob byte-identical."
gate_achieved: "MET (core) — both vantages (maya-thriving→/profile, dan-manager→/enterprise/workforce) driven live from a remote Mac over Tailscale, trusted LE cert (verify=0, no CA install), 0 console errors, 0 localhost/prod ejects, assets rendering; a clean cold reset-to-seed one-shot bring-up with the fixed tooling."
gate_distance: "gate met (core); 4 non-blocking documented residuals routed forward"
---

## TL;DR
M215's **core exit gate is MET** — the first remote Linux-VM demo deploy over Tailscale was driven end-to-end for
**both** vantages, on a **trusted** Let's Encrypt cert, with 0 ejects and assets rendering, reproducibly on a cold
reset-to-seed. This is a **clean close-on-gate**; the items below are **non-gate-blocking documented residuals**
surfaced during the live run, routed to future work per the three-fate rule (none release-scope-breaking). v2.2 is
closing (M215 is the FINAL milestone), so the three-in-release rext residuals go to `/developer-kit:close-release`
(the next operation) and the four demo-service/seed/content residuals go to the standing backlog.

## Root-cause clusters (documented residuals — gate MET, none blocking)

### Cluster 1: content set-dressing on a remote VM (F9)
- **Affected items:** taxonomy/library/skills surfaces (the manager mapped→verified funnel counts read 0); identity/profile/dashboard/workforce are **unaffected** (render fully).
- **Root cause:** the taxonomy (~42,790 public skills) + Directus content are set-dressed from the `.agentspace/snapshots` cache, not migrations. A fresh VM has no cache → `public.skills=0`. **Operational, not a code defect** — the runbook (`tailscale-serve.md` Step 4) documents scp/capture; the enhancement is to auto-sync/capture the cache to the VM.
- **Estimated scope:** small — a runbook-documented manual scp today; ~0.5 day for an auto-sync helper later.
- **Fate:** Fate 2 (documented residual + future enhancement).
- **Target milestone:** none in v2.2 (closing) → standing backlog / a future demo-tooling milestone.
- **Provenance:** iter-01 tik-1 + both-vantages (manager funnel 0s).

### Cluster 2: `app` demopatch sha-drift re-anchor (F5)
- **Affected items:** two `app` demopatches (`target-role authz-skip`, `ai-readiness loadMembers bound`) refuse on the current `app` tag — **NON-FATAL** (demo works; slower per-member Sentinel fan-out / unbounded AI-readiness hydration).
- **Root cause:** the demopatches' pinned pre-hash drifted vs the current `app` release. A demopatch-maintenance issue surfaced on a fresh full-clone; orthogonal to remote/Linux.
- **Estimated scope:** small — re-anchor the two pre-hashes against the current `app` tag + re-tag rext.
- **Fate:** Fate 2 (rext is frozen at `panorama-m215`; a re-anchor+re-tag is out of this rosetta close).
- **Target milestone:** none in v2.2 → standing backlog (demo-service maintenance).
- **Provenance:** iter-01 tik-1.

### Cluster 3: jobsimulation service-command crash (F13)
- **Affected items:** the AI-Simulations surface only (jobsimulation container exits(1) — printed CLI help, no run/serve subcommand). **OFF** the proven journey path (both hero journeys rendered fine).
- **Root cause:** likely a service-command/compose or version-drift issue that would hit **any** demo — not remote/Linux-specific.
- **Estimated scope:** unknown until investigated — a demo-service debug pass.
- **Fate:** Fate 2 (separate demo-service fix).
- **Target milestone:** none in v2.2 → standing backlog.
- **Provenance:** iter-01 cold reset-to-seed capstone.

### Cluster 4: seed hero-name cosmetic + repeatable remote-gate automation (F11 + automation)
- **Affected items:** (a) F11 — the seed identity key (`maya-thriving`) differs from the generated profile display name (login + render both work); (b) a committed, **repeatable** multi-journey Playwright suite pointed at the remote origin (the M215 proof used live one-off Chromium/Playwright runs; a checked-in remote-origin gate is a nice-to-have automation follow-up).
- **Root cause:** (a) seed-data naming inconsistency; (b) the M42/M202 e2e harness was pointed at the remote origin ad-hoc for the proof but not committed as a standing remote-origin gate.
- **Estimated scope:** small each — a seed polish; a harness-config follow-up.
- **Fate:** Fate 2.
- **Target milestone:** none in v2.2 → standing backlog (stack-seeding / stack-verify).
- **Provenance:** iter-01 tik-2 (F11) + "remaining toward the full exit gate" note.

### Cluster 5: rext-frozen README index reconcile (D-CLOSE-1/-2/-3) — owned by close-release
- **Affected items:** three DISTINCT rext-README residuals — M212's demo-stack `test_tooling` count-drift, M213's stack-injection index gap, M214's new `apply-ant-academy-dev-origins.sh` helper + patch not indexed.
- **Root cause:** each milestone's rext code-of-record is frozen at its annotated tag; a rosetta-only close must not re-point it. All reconcile in the ONE rext commit `/developer-kit:close-release` makes when it re-tags rext + bumps `.agentspace/rext.tag`.
- **Estimated scope:** small — one rext README-reconcile commit at release close.
- **Fate:** Fate 2 (already owned by close-release — the immediate next operation).
- **Target milestone:** v2.2 `/developer-kit:close-release`.
- **Provenance:** M212/M213/M214 closes (inherited); re-confirmed this audit.

## Projected post-resolution state
The **core gate is already met** — resolution of these residuals does not change the acceptance verdict. When they
land: the remote demo's taxonomy/library surfaces populate from the cache (Cluster 1), the two `app` demopatches
apply (Cluster 2), the AI-Simulations surface boots (Cluster 3), the hero name renders consistently + a standing
remote-origin gate exists (Cluster 4), and the rext READMEs index the panorama additions (Cluster 5 — at
close-release). No further work is needed for v2.2's external-shareability commitment.

## Cross-references
- Iter ledger: [`progress.md`](progress.md) + the canonical [`iter-01/findings.md`](iter-01/findings.md)
- Decisions: [`decisions.md`](decisions.md)
- Deferral audit (GREEN): [`audit-deferrals/deferral-audit-2026-07-11.md`](audit-deferrals/deferral-audit-2026-07-11.md)
- Propagation close-gate: [`propagation-checklist.md`](propagation-checklist.md)
- Iteration protocol used: `corpus/ops/verification.md` (+ `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/demo/playthroughs.md`) — the remote-origin variant of the cold-reset-to-seed gates
- Runbook delivered: [`corpus/ops/demo/tailscale-serve.md`](../../../../../corpus/ops/demo/tailscale-serve.md)

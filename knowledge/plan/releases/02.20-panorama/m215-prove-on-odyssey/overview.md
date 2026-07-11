---
milestone: M215
slug: prove-on-odyssey
version: v2.2 "panorama"
milestone_shape: iterative
status: archived
created: 2026-07-11
last_updated: 2026-07-11
complexity: large
depends_on: M212, M213, M214
exit_gate: "From a fresh box, /demo-up N --public-host billion.taildc510.ts.net (OPT-IN) brings up a demo reachable at https://billion.taildc510.ts.net; a teammate on a DIFFERENT tailnet machine completes a full employee AND manager hero journey (Clerkenstein login + a real journey) with 0 localhost/prod ejects, 0 CORS blocks, 0 cert-untrusted warnings, 0 mixed-content, and assets rendering — reproducibly on a cold reset-to-seed bring-up. The UNSET knob (default) remains byte-identical to today (regression-safe)."
iteration_protocol_ref: corpus/ops/verification.md (+ corpus/ops/demo/coverage-protocol.md + corpus/ops/demo/playthroughs.md — a remote-origin variant of the cold-reset-to-seed gates)
delivers: proof the demo is reachable + fully functional over Tailscale from a second machine, end-to-end
---

# M215 — Prove it on odyssey

## Goal
The acceptance gate. Bring a demo up on **`billion.taildc510.ts.net`** (the odyssey Proxmox host) with the
**opt-in** `--public-host` flag, and have a teammate on a **DIFFERENT** tailnet machine browse it **end-to-end** —
every in-app link resolves, Clerkenstein login works, content/images load — driving each real breakage to green on
a cold reset-to-seed.

## Exit gate (measurable)
From a fresh box:
- `/demo-up N --public-host billion.taildc510.ts.net` (OPT-IN) brings up a demo reachable at
  `https://billion.taildc510.ts.net` (one HTTPS origin, trusted `tailscale cert`);
- a **remote** tailnet browser (Playwright from a second tailnet host, reusing the M42 e2e harness + the M202
  Playthrough seat-switch for hero login) completes a **full employee AND manager hero journey** —
- with **0 localhost/prod ejects, 0 CORS blocks, 0 cert-untrusted warnings, 0 mixed-content, assets rendering** —
- **reproducibly on a cold reset-to-seed** bring-up;
- and the **UNSET knob (default) is byte-identical to today** (the regression contract).

## Why iterative (not section)
The reconfiguration is fully specified by M212–M214, but the **last breakages only surface on a live cross-machine
run**: secure-context (Web Crypto off-localhost), mixed-content under the proxy, cookie same-site over the real
host, the `tailscale cert` PEM shape vs Go `ListenAndServeTLS`, and RAM fit on the ballooned VM. A fixed `In:` list
would be speculative; the exit gate is the commitment. (Mirrors v1.10b M53 / v2.1 M211's acceptance-closer role,
but iterative because the remote path is unproven.)

## Iteration protocol
The fit-up/dress-rehearsal **bring-up → drive from a second machine → capture every eject/block/warning → fix in
the M212 knob / M213 TLS / M214 CORS+patch surface → re-run** loop, driven by `corpus/ops/verification.md` + the
coverage/playthroughs gates run from a **remote origin**. Tik/tok until the gate holds on a cold reset-to-seed.
**No new platform edits invented during iteration** — a surfaced platform-source hardcode routes to a NEW sha-pinned
rext patch (M214's mechanism), never a clone edit.

## Risks to burn down (first iters)
1. **RAM:** `billion` is ballooned to ~13 GB available; the demo UI tier wants ~12 GB. Iter-01 validates it fits,
   or exercises the `--no-ui` / `DEMO_NO_UI` escape, or bumps the VM RAM on odyssey (`qm set`).
2. **Secure-context / mixed-content:** confirm the HTTPS-everywhere proxy gives Clerk a secure context and no page
   makes a plain-`http://` browser call.
3. **Single-vs-multi stack per VM:** confirm N and that the offset appears correctly in every baked MagicDNS URL.
4. **`tailscale cert` renewal** for a long-lived stack (PEM shape confirmed vs Go TLS; renewal cadence).
5. **`@clerk/nextjs clerkMiddleware`** boot-time pk-host assumptions (expected host-agnostic — confirm live).

## Depends / parallel
- **Depends on:** M212 + M213 + M214 (all three surfaces).
- **Parallel with:** none — it is the joint re-measure of the whole release.

## Delivers → knowledge
The acceptance record + any live-surfaced fix routed back into the M212–M214 surfaces; finalizes
`corpus/ops/demo/tailscale-serve.md` with the proven walkthrough.

**Propagation close-gate (user directive 2026-07-11).** This is the FIRST Linux + remote-over-Tailscale
deployment, so **M215 does not close until every finding is propagated into all three surfaces — tools (rext),
corpus (KB), and skills** — so the next Linux/remote stack build is covered properly. Tracked in
[`propagation-checklist.md`](propagation-checklist.md); each `iter-NN/findings.md` item must land in ≥1 surface
(no orphan finding), and the corpus runbook must let a fresh reader stand up a remote demo on a new Linux VM
unaided. Close verifies this checklist is complete.

## KB dependencies
`corpus/ops/verification.md`, `corpus/ops/demo/coverage-protocol.md`, `corpus/ops/demo/playthroughs.md`,
`corpus/ops/demo/tailscale-serve.md` (M214), the odyssey server reference (kb-ant-business `odyssey` skill).

# M215 — Propagation checklist (HARD closure gate; user directive 2026-07-11) — ✅ COMPLETE

> **User directive:** since this is the FIRST time we deploy a demo on a Linux machine (and make it accessible
> from remote), by the END of M215 every finding + new fact MUST be propagated into the **knowledge base**, the
> **tools (rext)**, and the **skills** — so the next time a stack is built on Linux, or made remote-accessible,
> this is covered properly. This is a **close-gate**: M215 does not close until every row lands.

**STATUS: satisfied.** Every deployment-relevant finding (F1, F2, F4, F6, F8, F9, F12) is in tools + KB + skills; a
fresh reader can stand up a remote demo on a new Linux VM from `corpus/ops/demo/tailscale-serve.md` alone (verified).

## 1. Tools (rext — authoring copy, tag `panorama-m215`) — ✅
- [x] **F3** `git tag --list | head` SIGPIPE → `for-each-ref --count=1` + regression test (400-tag fixture).
- [x] **F4** `ensure_ssh_agent` — keyless agent before compose build (EXIT-trap reaped; no-op when SSH_AUTH_SOCK set).
- [x] **F6** `precreate_linux_data_dirs` — Linux-only, 0777 the Bitnami bind-mount data dirs before compose up.
- [x] **F1/F2/F8** `preflight_host_prereqs` — Go + atlas (always) + tailscale operator (public-host) → **fail loud**
      with exact fix lines; `DEMO_NO_HOST_PREFLIGHT=1` opt-out.
- [x] **F8** `migrate-demo.sh` fails loud on a genuine atlas migration failure (no longer a masked "warning").
- [x] **F12** teardown (`rosetta-demo` down) + up-path reset `tailscale serve` (offset-scoped per-port off; idempotent
      re-deploy); `gen_tailscale_serve.py --reset`.
- [x] Tests: demo-stack **424 passed** (383 baseline → +25 host-prereqs → +16 F12), stack-injection 147p/8s;
      shellcheck clean; macOS/dev path byte-identical (Linux-only / missing-prereq-only / public-host-only guards).
- [x] rext tag **`panorama-m215`** @ `00ba6b6` (annotated). `.agentspace/rext.tag` re-pin → `/developer-kit:close-release`.

## 2. Corpus / KB — ✅
- [x] **`corpus/ops/demo/tailscale-serve.md`** — full remote Linux-VM deploy runbook (Step 0 prereqs w/ install
      commands → GitHub-PAT clone → workspace → secrets → snapshot cache → `--public-host` bring-up → `tailscale serve`
      → verify [exact curls + cockpit login, both vantages] → teardown [F12 serve-reset]) + an F1–F12 host/deploy
      finding-set table (F13 noted out-of-scope — the jobsimulation crash, off the proven journey path).
- [x] **`corpus/ops/setup_guide.md`** — "Linux host prerequisites (for a remote/VM demo over Tailscale)" section.
- [x] **`corpus/ops/rosetta_demo.md`** — remote/Linux deploy proven cross-machine (both vantages) + prereq pointer.
- [x] **`corpus/services/clerkenstein.md`** — `tailscale cert` needs-operator caveat.
- [x] Stale **odyssey KB** (kb-ant-business) flagged as a cross-repo follow-up (13 VMs not 4; no Go 1.26) — noted,
      out-of-corpus (not fixed from here).

## 3. Skills — ✅
- [x] **`/demo-up`** (SKILL.md) — remote Linux-VM path + `--public-host` + host prereqs + runbook pointer.
- [x] **`/stack-secrets`** (SKILL.md) — provisioner needs Go on the host.
- [x] **`/dev-up`** (SKILL.md) — bare-Linux-VM note (shared F4/F6 fixes + PAT clone).

## Verification (close gate) — ✅
- [x] No orphan finding — every deployment finding lands in ≥1 surface (verify agent GREEN).
- [x] The runbook is followable end-to-end; cross-references + anchors resolve.
- [x] **Live proof:** both vantages (employee Maya `/profile` + manager Dan `/enterprise/workforce`) driven from a
      remote Mac over Tailscale (trusted LE cert, 0 console errors, 0 localhost ejects) + a clean cold reset-to-seed
      one-shot bring-up with the fixed tooling.

_Residuals (documented, not blocking — routed to future tiks / separate fixes): F9 snapshot cache for content
surfaces; F5 demopatch re-anchor; F11 seed hero-name cosmetic; F13 jobsimulation service-command (AI-sim surface)._

---

**Close re-verification (2026-07-11, `/developer-kit:close-milestone` — NOT rubber-stamped).** Independently
confirmed at close: all 9 `tailscale-serve.md` See-also cross-references resolve; the `setup_guide.md` §"Linux host
prerequisites (for a remote/VM demo over Tailscale)" + `clerkenstein.md` §"Remote HTTPS over the tailnet" anchors
exist; the `/demo-up`, `/stack-secrets`, `/dev-up` skill diffs carry the Linux/remote content; every deployment
finding (F1/F2/F4/F6/F8/F9/F12) maps to ≥1 surface (no orphan); the rext fix set (F3/F4/F6/F12 + host pre-flight) at
tag `panorama-m215` @ `00ba6b6` reviewed GREEN with demo-stack **424** + stack-injection **147p/8s** re-run passing +
shellcheck clean. Two minor `tailscale-serve.md` doc-accuracy nits fixed at close (the stale `F1–F11` finding range →
`F1–F12`/`F1–F13`; a truncated setup_guide §-anchor label). Gate: **CLOSED — propagation close-gate SATISFIED.**

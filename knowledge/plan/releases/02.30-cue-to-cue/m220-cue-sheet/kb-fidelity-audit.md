---
title: "KB Fidelity Audit — M220 (cue sheet)"
date: 2026-07-14
scope: milestone:M220
invoked-by: build-milestone
---

## Verdict

**YELLOW** — proceed with tracking.

No **unfilled** blind area: both blind areas (`safety.md` Part 3, the `/demo-up` defaults table) are **declared
`Delivers →` lines in the milestone's own `overview.md`** — which is exactly the sanctioned way to clear a
blind-area finding. No stale claim survives into implementation: the four stale claims found are either **the
milestone's own work product** (S0) or were **corrected inline** (the plan's stale anchors, below).

**The audit's real catch is KB-1/KB-3.** M220's plan itself carried the release's signature hazard (D17): its S0
site list was a **prose count that had never been checked against the tree**. Fixing "the 4 sites" would have left
**3 live sites** still lying, and the milestone would have reported S0 done. The plan was corrected before Phase 1.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Org count in the demo seed | `demo/README.md`, `rosetta_demo.md`, `demo-up/SKILL.md`, `stack-seed/SKILL.md` | rext `stack-seeding/presets/stories.seed.yaml`, `demo-stack/up-injected.sh` | **PAIRED — STALE** |
| Demo network exposure / port binding | `demo/demo/tailscale-serve.md` | rext `stack-injection/gen_injected_override.py`, `demo-stack/up-injected.sh` | **PAIRED — STALE (false safety claim)** |
| Safety contract — **exposure axis** | — (`safety.md` has Part 1 read-side + Part 2 write-side **only**) | rext `demo-stack/up-injected.sh`, `stack-injection/`, `clerkenstein/` | **BLIND-AREA** *(declared: Delivers 1+2)* |
| `/demo-up` defaults contract | — (no `knob \| default` table anywhere in `corpus/`) | rext `demo-stack/up-injected.sh`, `demo-stack/rosetta-demo` | **BLIND-AREA** *(declared: Scope.In (b))* |
| Tailscale capability ladder / host auto-discovery | `tailscale-serve.md` (v2.2 runbook) | rext — **no auto-discovery exists** (S3 builds it) | DOC-ONLY *(S3 — not audited here)* |

## Fidelity Findings

### KB-1 — the "2 orgs" claim is stale at **7** sites, not 4 · **STALE · load-bearing for S0**
- **Source:** the milestone plan's own S0 site list (`overview.md`, `progress.md`).
- **Expected (plan):** 4 sites.
- **Actual (tree):** **7**. The plan missed `corpus/ops/rosetta_demo.md:49` and
  `.claude/skills/stack-seed/SKILL.md:50`, and its `up-injected.sh` anchor was stale (`:1081` → **`:1317`**).
- **Ground truth:** `stories.seed.yaml` ships **three** `org:` entries — `:37` Cervato Systems
  (`ai-transformation`, 220) · `:75` Solvantis (`onboarding-ramp`, 120) · `:136` Northwind Aviation
  (`ai-readiness`, 200). `seed-generation-manifest.yaml:8` already says **"all 3 orgs"** — so the manifest and
  the prose disagree *inside the same repo*.
- **Verdict:** STALE. **Fix owner:** doc (7 sites). **Applied:** plan corrected; the doc fix is S0's work.

### KB-2 — the false exposure claim is at **`tailscale-serve.md:452-453`**, not `:405-407` · **STALE · load-bearing for S0**
- **Expected (doc):** *"no open 0.0.0.0-on-the-LAN surprise… Binding `0.0.0.0` is gated on the knob precisely so
  it is never ambient."*
- **Actual (code):** `gen_injected_override.py` emits published ports as **bare `"<hostport>:<target>"` pairs with
  no `127.0.0.1` prefix** at **three** sites — `:210` (directus), `:276-277` (frontends:
  `lines.append(f'      - "{host_port + offset}:{target}"')`), `:308` (backends). Docker's default for a bare
  `host:container` mapping is **`0.0.0.0`**. So **every demo container is published on all interfaces on every
  `demo-up`, today, with or without `--public-host`.** `BIND_HOST` (`up-injected.sh:76`) gates only the two
  **host-native** servers (cockpit, ant-academy) — not one container.
- **The doc already contradicts itself:** **`tailscale-serve.md:239`** states the truth —
  *"`docker-proxy` binds the demo's offset ports on **`0.0.0.0`**, which includes the VM's own `100.x`
  tailscale [address]"*. The correction is the doc **agreeing with itself**, not new knowledge.
- **Verdict:** STALE (a **false safety claim** — the worst class). **Fix owner:** doc. Ships regardless of the
  flip decision. Stale anchor corrected in `overview.md`, `progress.md`, `roadmap.md:422`.

### KB-3 — `--profile` / `--services` are doc-promised on a parser that **hard-errors** on them · **STALE · LIVE false promise**
- **Source:** `.claude/skills/demo-up/SKILL.md` `argument-hint`:
  `[N] [--public-host <magicdns>] [--profile P] [--services "a b"]`.
- **Actual:** there are **two** entry points, and the hint conflates them:
  - **`up-injected.sh`** — the one the skill actually invokes (`SKILL.md:52`) — accepts **only** `<N>` and
    `--public-host`, and on anything else prints `unknown argument` and **`exit 1`** (`:26-27`).
  - **`rosetta-demo`** (the wrapper) is where `--profile` / `--services` live (`:110-113`); `SKILL.md:62` uses
    them correctly against *that* binary.
- So **`up-injected.sh --profile X` exits 1 today** while the skill's flag list implies it works.
- **Verdict:** STALE. This is the **pre-existing RED** that pre-proves S2's both-directions fence — the fence is
  not theatre, it fails against the current tree. **Fix owner:** doc (S2's table must record *which entry point
  reads which knob*).

### KB-4 — `D-DESIGN-1` is an **overloaded ID across releases** · **UNVERIFIABLE-by-grep · load-bearing for S1**
- **v2.2's D-DESIGN-1** = *"public reach is never default-on"* (`demo-up/SKILL.md:79` — note: **`:79`**, the plan
  said `:78`). This is the one S1 must supersede.
- **v2.3 has its OWN D-DESIGN-1** = *"the <5 s gate is on ACCESS, not full first-page render"*
  (`roadmap.md:127`, `state.md:105`, `latency-budget.md:21`).
- A bare `D-DESIGN-1` in this release **resolves to the wrong decision**. `roadmap.md:129` and `state.md:107`
  already qualify it correctly as *"v2.2's D-DESIGN-1"* — S1's supersession prose **must do the same, every
  time, never bare**.
- **Verdict:** ambiguity, not divergence. **Fix owner:** S1's authoring discipline. Recorded in `progress.md`.

### KB-5 — `safety.md`'s exposure grep is **2 hits, not 0** (substance unchanged) · ALIGNED-in-substance
- **Expected (plan):** `tailscale|remote|expose|network|localhost` → **0 hits** in `safety.md`.
- **Actual:** **2** hits — `:146` and `:215` — but **both are incidental** (`in-network` Directus addressing,
  `http://directus:8055`). Neither is an exposure contract.
- **The section list is the real proof:** `safety.md` runs **Part 1 — the read side** → **Part 2 — the write
  side** → *How this relates to the platform's own isolation* → *Future* → *See also*. **There is no Part 3.**
- **Verdict:** the BLIND-AREA finding **stands**; only its evidence was mis-stated. Corrected in `progress.md`.
  *Lesson (D17): grep the **sections**, not the **words** — a keyword hit is not a contract, and a keyword miss
  is not proof of absence.*

## Completeness Gaps

1. **`DEMO_*` knob surface is undocumented as a contract (critical — S2 owns it).** Measured: **35** raw `DEMO_*`
   tokens across rext; **~25** are real user-facing knobs. The residue is internals (`DEMO_WS`, `DEMO_N`,
   `DEMO_STACK`, `DEMO_OFFSET`, `DEMO_PORT_OFFSET`, `DEMO_LOCAL_BASES`, `DEMO_ONLY`), a computed name
   (`DEMO_1_DIRECTUS_DSN`, `DEMO_1_STAGING_ZONE_DIRECTUS_DSN`), and a grep artifact (`DEMO_NO_`). The table must
   **classify**, not dump. No `knob | default` table exists anywhere under `corpus/` (verified).

2. **Fence homes exist — no new harness needed (incidental).** `stack-core/tests/test_corpus_index_guard.py` is
   the **existing doc-vs-code fence precedent** (stdlib `unittest`, reads the real corpus tree) → the org-count
   fence belongs beside it. `stack-core/tests/test_gen_override.py` + `stack-injection/tests/test_injection.py`
   are the port-emission fence homes. Both are Python/`unittest`; no new dep.

3. **`safety.md` Part 3's subject matter is real and enumerable (critical — S1 owns it).** The exposure surface a
   Part 3 must cover, all verified present in code: Clerkenstein disarms token verification; the authz-skip
   demo-patch is **default-on** (`DEMO_NO_AUTHZ_SKIP:-0`, `up-injected.sh:791,886`); the cockpit is a
   password-free hero launcher; and — per KB-2 — **container ports are already `0.0.0.0`-published today**.

## Applied Fixes

| File | Fix |
|---|---|
| `m220-cue-sheet/progress.md` | S0 site list `4 → 7` (+ the 2 missed sites, + `up-injected.sh:1081→1317`); `tailscale-serve.md:405-407→452-453` + the self-contradiction at `:239`; the 3 real port-emission sites; S1's "0 hits" → "2 incidental hits, no Part 3"; the **KB-4 `D-DESIGN-1` ID-collision warning**; S2's **two-parser** RED; fence homes named |
| `m220-cue-sheet/overview.md` | Same anchor + site-count corrections in the ask-vs-reality table, Scope.In (b), and Delivers 3 |
| `knowledge/plan/roadmap.md` (`:422`) | `tailscale-serve.md:405-407` → `:452-453`; `gen_injected_override.py:259-260,292-294` → `:210,276-277,308` |

## Open Items (require user decision)

**None.** No item blocks Phase 1.

## Gate Result

**YELLOW — proceed with tracking.** Both blind areas are declared milestone deliverables (the sanctioned
clearance). All stale claims are either S0/S2's work product or corrected inline. Findings KB-1…KB-5 are tracked
above and reflected in `progress.md`.

**Two findings are pre-proven REDs the milestone's fences must reproduce** — they fail against the current tree,
which is what makes them fences and not theatre:
- **KB-1/KB-2** → S0's org-count + published-port-shape fences.
- **KB-3** → S2's CLI-flag ↔ docs both-directions fence.

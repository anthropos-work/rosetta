# M220 — Spec notes

_Technical detail accumulated during the milestone: file:line maps, command transcripts, measured numbers._

## Pre-flight audits — S0 (the two lies the docs tell)

**Phase 0b — `/developer-kit:audit-kb-fidelity --milestone=M220` → YELLOW.** Report:
[`kb-fidelity-audit.md`](./kb-fidelity-audit.md). Commit `d395946`.

- 0 unfilled blind areas — `safety.md` Part 3 (S1) and the defaults table (S2) are both **declared `Delivers →`
  lines** in `overview.md`, which is the sanctioned clearance path.
- 5 fidelity findings (**KB-1…KB-5**), 3 completeness gaps. None blocks Phase 1.
- **The audit's own catch:** M220's plan carried the release's D17 hazard — its S0 site list was a prose count
  never checked against the tree (4 claimed, **7** actual). Anchors corrected in the plan before any code landed.

Reused for **S1** and **S2** per §"Audit reuse": same subsystem (the demo/doc surface), no load-bearing knowledge
doc changed between sections, and the audit already covered all three sections' topics in one pass.

## The verified file:line map (S0)

### The org-count sites — 7, all saying "2", preset ships 3

| # | Site | Repo |
|---|------|------|
| 1 | `.claude/skills/demo-up/SKILL.md:109` | rosetta |
| 2 | `.claude/skills/demo-up/SKILL.md:153` | rosetta |
| 3 | `.claude/skills/stack-seed/SKILL.md:50` | rosetta |
| 4 | `corpus/ops/demo/README.md:34` | rosetta |
| 5 | `corpus/ops/rosetta_demo.md:49` | rosetta |
| 6 | `demo-stack/up-injected.sh:1317` (`seed_label`) | rext |
| 7 | `stack-seeding/presets/stories.seed.yaml:1` (header) | rext |

**Ground truth** — `stories.seed.yaml` `org:` entries:

| Line | Org | Narrative | Size |
|------|-----|-----------|------|
| `:37` | Cervato Systems | `ai-transformation` | 220 |
| `:75` | Solvantis | `onboarding-ramp` | 120 |
| `:136` | Northwind Aviation | `ai-readiness` | 200 |

`seed-generation-manifest.yaml:8` already says **"all 3 orgs"** — the manifest and the prose disagree inside the
same repo. The manifest was right.

### The `/demo-up` knob surface (S2) — derived from the parsers

**25 env knobs + 9 CLI flags, across TWO entry points.**

| Entry point | Flags accepted | On an unknown arg |
|---|---|---|
| `demo-stack/up-injected.sh` (what `/demo-up` runs) | `<N>` + **`--public-host`** — that is all | **`exit 1`** (`:26-27`) |
| `demo-stack/rosetta-demo` (lifecycle wrapper) | `--profile` `--services` `--ref` `--only` `--resolve-only` `--fapi-host` `--bapi-ip` `--webhook-secret` | per-subcommand |

**The shape:** every feature knob is an **opt-OUT** (`DEMO_NO_*`, default `0`). The **only** default-off knob is
`STACK_PUBLIC_HOST` (`""`). Confirmed against the parser: `DEMO_STORIES=1`, and `DEMO_NO_UI` / `DEMO_NO_SETDRESS`
/ `DEMO_NO_STORIES` / `DEMO_NO_AUTHZ_SKIP` all `0`.

⇒ **"pull all the data + seed the 3 orgs" was ALWAYS the default.** The failure people attribute to a missing
default is a **cold snapshot cache** (replay is cache-first and *never* captures), not a knob.

`DEMO_DISK_AVAIL_KB` is carved out as internal (a test seam for the disk pre-flight). `DEMO_LOCAL_BASES` /
`DEMO_ONLY` appear in the repo-wide grep but exist **only** in docs/tests — not real knobs.

### The published-port emission — 3 sites, all bare

`stack-injection/gen_injected_override.py`:

| Line | Carrier | Emission |
|------|---------|----------|
| `:210` | directus | `'      - "%d:8055"' % (8055 + offset)` |
| `:276-277` | frontends | `lines.append(f'      - "{host_port + offset}:{target}"')` |
| `:308` | backends | `f'"{int(p["published"]) + offset}:{p["target"]}"'` |

No `127.0.0.1` prefix at any site ⇒ Docker's default for a bare `host:container` mapping is **`0.0.0.0`**.
`BIND_HOST` (`up-injected.sh:76`) is set to `0.0.0.0` only under a public host — but it is consumed **solely** by
the two **host-native** servers (cockpit, ant-academy), never by a container.

## Pre-flight audits — S5 + S6

**Phase 0b reused** from S0 per §"Audit reuse": same subsystem (the demo bring-up / injected-env / doc surface),
the milestone-scoped `/developer-kit:audit-kb-fidelity --milestone=M220` (YELLOW, commit `d395946`) already
covered S5/S6's topics in one pass, and the only knowledge docs changed since are M220's own S0–S2 deliverables.

## The live cycle (S5 + S6) — one demo, `billion`, demo-1, offset 10000

Bring-up: `up-injected.sh 1 --public-host billion.taildc510.ts.net` (foreground, attached). Asserted from a
**tailnet peer** (the Mac), never on-host. Cold: migrate ✓ (4 services + sentinel policy), replay ✓ (taxonomy
330 261 rows / directus 11 986 / sim-embeddings 1 490), stories seed ✓ (3 orgs × 3-hero trio, 9-identity roster),
autoverify ✓ (`public.skills = 42790`). Demo clone left **0-dirty**; `demopatch.log` **empty** (nothing refused).

### S5(i) — the academy session A/B (the DoD)

| arm | one variable | result |
|---|---|---|
| **A** (control) | login as `maya-thriving` → `/profile` | signed in — page renders **"Maya Chen · DevOps Engineer · Berlin"** |
| **B** (the fix) | login → `/profile` → **ACADEMY `:13077`** → `/profile` | **STILL signed in as "Maya Chen"** |

Cookies across arm B: `__session` **present throughout**; `__client_uat` a **live timestamp** (`1784052754` →
`…756` → `…759`), **never `0`**. Direct `curl -D -` at `:13077`: **zero `Set-Cookie` headers** — no
`__session=; Expires=1970`, no `__client_uat=0; Domain=…`, no keyless bounce.

**Values-blind secret check** (sha16 of the value, never the value):
`platform/.env` `CLERK_SECRET_KEY` = `b47f934a4c92f784` · academy `.env.local` = `7adefe7a43b3497a` ⇒ **different**,
and the academy's is `sk_test_…` (Clerkenstein). Publishable-key lines in `.env.local`: **1** (was 11 → last-wins).

### S6 — egress, measured in a real browser on an authenticated load

**0 hits** across all 11 denied hosts (GTM · GA · DoubleClick · Google Ads · LinkedIn · Plausible · Bellasio ·
BetterStack · clerk-telemetry · jsdelivr · real Clerk). The fence also asserts it captured traffic at all.

**(g) artifact RED → GREEN** (files in the built bundle containing each id):

| id | pre-fix image | post-fix image |
|---|---|---|
| `GTM-PXRTBZK` | 2 | **0 in `.next/static`** (1 `.js.map` only) |
| `plausible.io` | 6 | **0 in `.next/static`** (5 server chunks / maps) |
| `analytics.bellasio.com` | 2 | **0 in `.next/static`** (1 map) |
| `uptime.betterstack.com` | 2 | **0 in `.next/static`** (1 map) |

The client bundle — the only thing the browser receives — carries **none** of them; the survivors are dead server
chunks and source maps. The browser capture is the load-bearing proof.

**(h) clerk-js from disk.** Cache dir `demo-stack/stacks/.clerkjs-cache/` (box-level) holds 4 chunks after the
first load, incl. `npm_@clerk_clerk-js@5_dist_clerk.browser.js` (**321 927 B**) + `ui-common` (442 KB). The
browser fetches clerk-js **from the FAPI host**, never `cdn.jsdelivr.net`.

**Alignment after touching the FAPI** (the item claimed *alignment-invisible*; verified, not assumed):

| DNA | genes | overall | critical | gate |
|---|---|---|---|---|
| `clerk-2.6.0.json` (Go) | **27/27** | 100.0% | 100.0% | **MET** |
| `clerk-js-5.json` (JS/FAPI) | **9/9** | 100.0% | 100.0% | **MET** |

### The patch-set fingerprint fired on its first live run

```
next-web: cached image demo-1-next-web was built with a DIFFERENT demo-patch set
  (<none: predates the fingerprint> != cee1e4ff4cf9cd1e…) — removing + rebuilding
```

Without it the stale image (matching endpoint + pk) would have been **reused**, and `next-web-no-thirdparty`
would never have reached the bundle — a green bring-up over a demo still phoning home to seven third parties.

### rext rolls

`cue-to-cue-m220-r1` → `-r2` (fingerprint) → `-r3` (port reap) → **`-r4`** (the live proof spec). Host pinned at
`-r4`-equivalent code; `.agentspace/rext.tag` updated (the pin guard **correctly refused** the first bring-up
when the clone and the tag disagreed — it worked exactly as designed).

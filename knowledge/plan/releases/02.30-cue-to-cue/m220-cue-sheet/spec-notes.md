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

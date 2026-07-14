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

## Pre-flight audits — S3 + S4

**Phase 0b reused** from S0 per §"Audit reuse": same subsystem (the demo bring-up / exposure / doc surface); the
milestone-scoped `/developer-kit:audit-kb-fidelity --milestone=M220` (YELLOW, commit `d395946`) covered S3/S4's
topics in the same pass; and the only load-bearing knowledge doc changed since is **`safety.md` Part 3 — which is
M220's own S1 deliverable, and is exactly the doc S3 was blocked on**. Re-auditing it against itself would prove
nothing.

## The live cycle (S3 + S4) — `billion`, demo-1, offset 10000

**One demo host, one agent, driven synchronously.** Both bring-ups run in bounded foreground polls; asserted from
a **tailnet peer** (the Mac), never on-host (an on-host run fakes a RED via an SSL artifact).

### Run 1 — DEFAULT-ON (the flip). `up-injected.sh 1` — **no flag**

```
demo-up: public-host AUTO-DISCOVERED — billion.taildc510.ts.net (all 6 tailscale capability rungs passed:
  daemon running · dotted MagicDNS name · MagicDNS enabled · operator set · cert MINTED).
==> [demo-1] tailscale serve: the presenter cockpit (:17700) is now fronted with HTTPS too —
    the demo's entry point is no longer the one plain-HTTP surface on the tailnet
==> [demo-1] presenter cockpit serving on https://billion.taildc510.ts.net:17700
```

**From the peer (no `-k` anywhere — the cert is genuinely trusted):**

| surface | port | result |
|---|---|---|
| **cockpit (S4)** | 17700 | **HTTPS 200, `ssl_verify=0`** — was plain HTTP before |
| next-web | 13000 | 307 (login redirect), `ssl_verify=0` |
| ant-academy | 13077 | 200, `ssl_verify=0` |
| studio-desk | 19000 | 302, `ssl_verify=0` |
| cosmo graphql | 15050 | 200, `ssl_verify=0` |
| backend | 18082 | 404 (bare root), `ssl_verify=0` |

`tailscale serve status` listed **`:17700`** alongside the other five — the S4 listener, live.

**Hero login end-to-end over the cert (a controlled A/B, one variable):**

| arm | request | result |
|---|---|---|
| **A** (control) | `/profile`, **no session** | **307** → an endless `__clerk_handshake` loop (curl bails at 50 redirects) — the middleware **refuses** |
| **B** | cockpit `[Log in as] maya-thriving` → `/profile` | **200**, final URL `/profile`, `__session` + `__client_uat` = **live timestamp** |
| **B'** | cockpit `[Log in as] dan-manager` → `/enterprise/workforce` | **200** |

Cockpit renders **all 3 orgs** (Cervato · Solvantis · Northwind) + heroes. autoverify **OK** (`casbin_rules = 1150`,
`public.skills = 42790`). Demo clones left **0 tracked-modified** — every demo-patch reverted after the bake.

### Run 2 — THE FALLBACK. `tailscale` made **genuinely** unavailable

First attempt at the simulation was **wrong and would have passed**: `PATH=/usr/bin:/bin` does not hide
`/usr/bin/tailscale`. The ladder duly discovered the host — *an unexecuted check reported as a pass.* Redone with
`chmod -x /usr/bin/tailscale` (`shutil.which` → `None`; tailscaled untouched, so the tailnet/SSH stayed up).

```
demo-up: public-host auto-discovery STOPPED at rung 1/6 — 'tailscale' is not on PATH.
  Falling back to a LOCALHOST demo: byte-identical to a no-flag bring-up — the demo works fully, it just is
  not reachable from another machine. Fix: install Tailscale — https://tailscale.com/download
```

| assertion | measured |
|---|---|
| cockpit line | `presenter cockpit serving on **http://localhost:17700**` |
| `tailscale serve` applies | **0** |
| serve listeners on the node | **0** |
| baked browser URLs | `NEXT_PUBLIC_{HOSTING,BACKEND_API,WUNDERGRAPH}` = **`http://localhost:*`** |
| cockpit bind | **`127.0.0.1:17700`** — and **connection REFUSED** from the node's `100.x` tailnet IP |
| demo works | cockpit `http` **200**; autoverify **OK — verified-working** |

**Byte-identical. The invariant holds.**

> **…and the same run surfaced F-M220-5.** `ss -ltnp` on that localhost demo: cockpit `127.0.0.1:17700` ✅ but
> **ant-academy `*:13077`, answering HTTP 200 on the tailnet IP** ❌. `BIND_HOST=""` passes no `-H`, and
> **`next dev` defaults to `0.0.0.0`** — so the corpus's *"`BIND_HOST` gates the two host-native servers"* is
> **half false**, the same lie S0 retracted, one layer up. Docs corrected; code routed to M221.

### The `tailscale cert` open question — SETTLED (it gated the flip)

```
mint #1  wall=0.01s   serial=05777C48EFDF12EFED0512F46E8B53AC466C
mint #2  wall=0.01s   serial=05777C48EFDF12EFED0512F46E8B53AC466C   ← IDENTICAL
journalctl -u tailscaled | grep acme  →  0 real ACME orders
issuer=Let's Encrypt  notBefore=Jul 11  notAfter=Oct 9
```

**tailscaled CACHES.** The repo's *"`tailscale cert` RE-ISSUES on re-run"* (2 sites) is a **doc bug**, corrected.
This was load-bearing, not trivia: rung 6 **mints on every default-on bring-up**, and **`ts.net` is a PSL entry**,
so the LE duplicate-certificate bucket is **per-tailnet** — had the claim been true, default-on would have
rate-limited the whole tailnet.

### Fences — RED-proven by MUTATION (a fence that only tests the happy path is theatre)

| mutant (the naive thing a reasonable engineer writes) | verdict |
|---|---|
| presence-probe rung 6 (`the binary exists ⇒ good`) | **RED — 20 tests** |
| accept `.Self.DNSName` verbatim (no dotted check) | **RED — 2** |
| soft rung 4 (missing `CurrentTailnet` is fine) | **RED — 1** |
| trust `rc==0` as a mint (never open the cert file) | **RED — 2** |
| drop the `\|\| true` (a tailscale hiccup **aborts every demo**) | **RED — 2** |
| probe despite `--no-public-host` | **RED — 2** |
| second-guess an explicit `--public-host` | **RED — 2** |
| capture stderr into the host (`2>&1`) — the fix-line **becomes** the hostname | **RED — 3** |
| PRE-S4: cockpit absent from the front list (the shipped defect) | **RED — 3** |
| file the cockpit under the UI tier | **RED — 2** |
| front the cockpit in the FIRST apply (pre-bind ⇒ `EADDRINUSE`) | **RED — 1** |
| `--no-cockpit` on **both** applies (S4 becomes a silent no-op) | **RED — 1** |

> **My own ordering fence was THEATRE on the first cut** — it scanned the raw script text and matched
> `--no-cockpit` inside the **comment** above the command, so deleting the flag from the *command* still passed.
> Only the mutation run exposed it. It now strips comments before asserting. **D17, reproduced inside the fence
> built to catch it.**

**Suites:** demo-stack **569** · stack-injection **214** · stack-core **129** — all green, including the **11 that
were already RED at `r4`** (F-M220-6).

### rext rolls

`cue-to-cue-m220-r1` → `-r2` (fingerprint) → `-r3` (port reap) → `-r4` (the live proof spec) → **`-r5`**
(S3 ladder + S4 cockpit-fronting + the cert-claim correction + the 11 stale-test fixes). Host pinned at
`-r4`-equivalent code; `.agentspace/rext.tag` updated (the pin guard **correctly refused** the first bring-up
when the clone and the tag disagreed — it worked exactly as designed).

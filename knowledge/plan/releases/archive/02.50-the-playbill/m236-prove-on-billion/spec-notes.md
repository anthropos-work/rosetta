# M236 — Spec notes

## Pre-flight audits — iter-01

**`/developer-kit:audit-kb-fidelity --milestone=M236`: VERDICT = RED (blocker).**
Report: [`kb-fidelity-audit.md`](kb-fidelity-audit.md) — 224 lines; 24 fidelity findings (5 blocker-class),
4 completeness gaps (2 critical), 2 blind areas.

The audit was dispatched at iter-01 open and returned mid-iter-02. Its verdict **blocks** per Phase 0b, and
iter-02 was stopped **before** its publish step executed (see `iter-02/decisions.md` INCOMPLETE-EXIT).

### The systemic finding

All six "method" docs — `tailscale-serve.md`, `verification.md`, `coverage-protocol.md`, `playthroughs.md`,
`latency-budget.md`, `demo-up-defaults.md` — contain **zero** mentions of content-stories,
`content_products`, `content-manifest`, or `content-player`. Both docs M236 declares as its
`iteration_protocol_ref` were last edited 2026-07-16/17 (M225/M226). **M236's declared method describes a
v2.4 world**; the feature it must prove is v2.5.

### The five blocker classes (each changes M236's shape)

1. **`overview.md:37` cites a page-object that does not exist.** "Reuses the shared
   `AISimulationResultContainer`" — zero hits in `rosetta-extensions`; it is a **next-web `.tsx` component**
   (`content-stories-routes.md:52`), not a harness object. The nearest harness object
   (`simulation-page.ts:9-10`) stops at the *launch* boundary and has no `/result/` locator. It must be
   **authored**, not reused — so Cluster 1 is larger than M235's carry-forward described.
2. **The exit gate contains an unprovable clause.** "Demo reachable only over the tailnet" contradicts
   `safety.md:405` ("every demo container is published on `0.0.0.0` — all interfaces — on EVERY `demo-up`,
   flag or no flag"; `:413` measured 14 ports → `0.0.0.0`) and `tailscale-serve.md:626-627` (on Linux this
   bypasses the host firewall — `ufw deny` does **not** block it). Provable only as a property of
   `billion`'s **network placement**, and only via an explicit **off-tailnet probe** — never as an assertion.
3. **Half the gate is unmeasurable with today's harness.** The content-player CTA
   (`cockpit.py:421-425`) emits **no `data-login-as`** attribute — and that attribute *is* the ACCESS
   predicate: `latency.ts:123-127` throws without it, before t0. Additionally `run-latency.sh:42-47`
   hard-rejects non-hero vantages (`exit 2`). So "p95 click→ACCESS < 5 s" cannot currently be measured for
   any of the 31 content-story actions.
4. **Two inherited carry-forward tasks are misdescribed.** The "2 drifted demopatch manifests" are **not
   drifted** — clone HEAD matches the pin exactly; the working tree is merely left-patched from an
   un-reverted run, which classifies `ALREADY_PATCHED` → idempotent no-op. Fix is a one-line
   `git checkout -- packages/core-js/src/constants/urls.ts`. And the ANT_ACADEMY "rendered-card count"
   descriptor M236 is told to **run** does not exist — the shipped one
   (`coverage-manifest.ts:709-713`) is a text-marker + length floor that **would pass on a zero-card grid**,
   i.e. it cannot detect the exact Thread-A failure it is cited to catch. It must be authored.
5. **The declared `iteration_protocol_ref` is hollow.** `verification.md` (335 lines) contains no
   measure→triage→fix loop and no gate — it supplies a gate *input*, not a protocol.

### Additional finding that reverses a documented rule

`coverage-protocol.md:421-431` mandates **excluding** the exact pages M236 must prove
(`skipPaths` contains `/\/result\/[0-9a-f-]{8,}/`). M236 must consciously **reverse** a documented rule —
that is a spec decision, not an implementation detail.

### Notable stale claims (non-blocking but live-risk)

- **`demo-up-defaults.md:146`** claims the host pre-flight is "non-fatal, never blocks" — it **`exit 1`s**
  (`up-injected.sh:345-347`). Highest-risk row for a cold `billion` bring-up.
- **24 of 28 line anchors in `demo-up-defaults.md` are stale** (e.g. `DEMO_NO_VERIFY` documented at `:1594`,
  actually `:2110`). Its own guard exists and was run: it flags 1 real disagreement
  (**`DEMO_NO_ACADEMY_FILL` undocumented — and it gates Thread A**) but has **no line-anchor check**.
- **Seats are `content-player-23` … `content-player-35`**, not 0-indexed as `overview.md:36` implies;
  unknown seat keys return 400.
- `session-clone-spec.md:66` says 9 sessions / 2-2-1; actual is **13** and 3-2-2. The manifest's real shape
  (4 products / 18 sessions) is stated **nowhere** in the corpus.
- `safety.md:424` ("no `127.0.0.1` prefix anywhere") is **false** — `gen_injected_override.py:577`.

### Scope-back (the audit also found work is cheaper than planned)

`hero-login.ts` already accepts arbitrary seat keys (no allowlist); `player_result_path` /
`manager_result_path` are **pre-resolved strings** (no runtime URL derivation needed); no new tailnet port
is required; both new ant-academy demopatches verified sha-current. **Hidden cost:**
`VantageManifest.identityKey` is **singular** — 13 seats means **13 manifests**, not one sweep.

### Standing finding fed INTO the audit (established live at iter-01, B2)

The v2.5 tooling is unpublished — `origin/main` is the M228 commit, 0 of 13 `playbill-*` tags are on
origin — while `billion` consumes tooling only via a pinned tag from origin. Flagged to the auditor as a
candidate blind area: no corpus doc names **publication** as a step in the authoring-copy → tagged-clone
consumption path. `CLAUDE.md` and `rosetta_demo.md` both describe the path as "built and tested in the
authoring copy and tagged, then consumed per-stack via a pinned-tag clone" — which reads as though tagging
alone makes a tag consumable by a remote host. It does not.

## Thresholds and denominators

- **Primary metric denominator = 31** landable (session × action) pairs. Derived per-session from
  `stack-seeding/presets/content-manifest.json`, not from prose. See `iter-01/decisions.md` D1 — note the
  `has_manager_view` per-session vs per-product trap (a product-level read returns `None` for all 4
  products and silently under-counts to 18).
- **p95 click→ACCESS < 5 s**, ACCESS defined per `corpus/ops/demo/latency-budget.md` (authenticated shell
  rendered + interactive with the hero's identity present). **State the environment with every number** —
  the same defect measured ~6 s on a laptop and ~112 s on the tailnet VM.
- **ai-labs is presence-only** (M231 verdict) — 2 rows must render, 0 landable result pages. Not a gap.

## Environment facts (billion)

- Workspace `/home/devops/panorama/`; stack workspace `stack-demo/`; rext pin SoT at
  `/home/devops/panorama/.agentspace/rext.tag`.
- 7.3 GiB RAM + 15 GiB swap; 193 G disk, 40 G free at iter-01. `docker system df`: **109 GB build cache,
  107.6 GB reclaimable** — prune before any cold UI-tier rebuild.
- The M217 rext pin guard is **FATAL** on mismatch (`demo-stack/ensure-clones.sh`); escape hatch is
  `DEMO_ALLOW_UNPINNED_REXT=1` for deliberate un-tagged authoring work.

## Publishing milestone tags is the established workflow (not an escalation)

origin carries **160 tags**, including the complete `casting-call-m225…m228` milestone series of the
previous release — and `casting-call-m228-hiring-scope-fix` (the tag `billion` currently runs) is itself on
origin. So pushing v2.5's milestone tags is routine release practice, not a novel action; M230–M235 simply
ran offline and never exercised the publish half. All 13 `playbill-*` tags are ancestors of local `main`,
`main` is a clean fast-forward over `origin/main` (`1d97861..60eff14`, 20 commits), and there are **0 tag
name collisions** on origin — so the publish is purely additive.

## Local toolchain gap (non-blocking)

`pytest` is not installed in the workstation's python3.14 — the `stack-core` / `demo-stack` Python suites
cannot be run locally as a pre-publish check. Go side was verified instead: `stack-seeding` builds clean
and **16 packages pass, 0 fail**; `stack-snapshot` / `stack-secrets` / `alignment` all build. The Python
surfaces are exercised on `billion` at bring-up. Routed as an observation, not an iter.

---

## Pre-flight audits — iter-02 RE-VERDICT (2026-07-20, post-user-resolution)

**Standing verdict: RED → DISCHARGED (proceed as YELLOW).**

The iter-01 audit's RED rested on **five blocker classes, all of which were spec-scope** — they asserted
that M236's *declared gate and method* were wrong, not that its *code or knowledge* was defective. The
user resolved all five on 2026-07-20 (see `decisions.md` → USER-BLOCKER-M236-01 → **RESOLUTION**). Class
by class:

| # | Blocker class | Disposition |
|---|---|---|
| 1 | cited page-object doesn't exist | **ACCEPTED as scope** (B3) — harness authored from scratch, 13 seats → 13 manifests; `overview.md` In-list amended |
| 2 | "tailnet-only" unprovable | **CLAUSE DROPPED** (B1) — removed from the gate; `safety.md` §3 Part 3 disclosure stands as-is |
| 3 | p95 unmeasurable for content seats | **GATE SCOPED to HERO vantages** (B2) — content-seat latency explicitly out of scope for v2.5 |
| 4 | must reverse a documented `skipPaths` rule | **AMENDMENT IS NOW IN-SCOPE WORK** (B4) — `coverage-protocol.md` amended in the same change as the reversal |
| 5 | hollow `iteration_protocol_ref` | **REPOINTED** (B5) → `coverage-protocol.md` + `playthroughs.md`; their content-stories sections backfilled by this milestone |

The systemic finding (six method docs describe a v2.4 world) is **not discharged** — it is **converted into
milestone deliverables** under B4/B5. It remains true, and M236 now owns closing it. That is a YELLOW
posture, not a RED one: the gaps are known, named, assigned, and in-scope, which is exactly Phase 0b's
YELLOW definition ("proceed; the gaps become the iter's known-context").

**Nothing in the audit's remaining 19 fidelity findings blocks Phase P**, which touches no knowledge doc
and no platform repo. Phase P resumed and completed under this re-verdict.

### Independent denominator confirmation (unplanned, valuable)

Phase P's verification step read the **published** `content-manifest.json` on `billion` at
`playbill-m235-hardened` and recomputed the gate denominator from the artifact itself:

```
products: 4
  simulation          sessions=13  manager_views=13  -> 26
  skill-path-legacy   sessions= 2  manager_views= 2  ->  4
  ai-labs             sessions= 2  manager_views= 0  ->  0  (presence-only, not landable)
  skill-path-new      sessions= 1  manager_views= 0  ->  1
TOTAL sessions 18 | manager views 15 | raw pairs 33 | LANDABLE 31
```

**31 confirmed against the shipped artifact**, matching `metrics.json` exactly. Note the arithmetic that
makes the trap concrete: a naive read gives **33** (counting ai-labs' 2 presence-only player actions) and a
product-level `has_manager_view` read gives **18**. Only the per-SESSION read with ai-labs excluded yields
**31**.

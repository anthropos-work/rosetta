---
milestone: M217
slug: clean-stage
version: v2.3 "cue to cue"
milestone_shape: section
status: planned
created: 2026-07-13
last_updated: 2026-07-13
complexity: medium
depends_on: none
delivers: a /demo-up that comes up GREEN — so that every number measured afterwards is real. Plus corpus/ops/demo/demopatch-spec.md (the sanctioned escape hatch this release depends on has no corpus doc)
issues: "the last real run on billion: cockpit CRASHED on a leaked port (stale manifest → dead clerk-ids); all 3 snapshot replays SKIPPED (cold cache → structural-only catalog); autoverify ended FAILING; jobsimulation exits(1); both app perf demo-patches REFUSE on sha-drift with the reason piped to /dev/null"
---

# M217 — Clean stage

## Goal
A `/demo-up` that comes up **green**. This milestone builds no feature — it removes the confounds that make every
downstream measurement a lie.

## Why this is a HARD BARRIER (not just a nice-to-have first step)

The user's reported defect (1–2 min cockpit login) was measured on a stack where:

- **The cockpit had CRASHED.** `OSError: [Errno 98] Address already in use` at `cockpit.py:567` — port `7700+off`
  was still held by a **leaked cockpit from the prior run**. That predecessor kept serving a **stale
  `cockpit-manifest.json`** (dead `__clerk_identity` keys) against a **freshly re-seeded DB**. The run *before* that
  aborted entirely on a leaked `0.0.0.0:18082`. **Two of the last three runs on `billion` were broken by leaked
  ports.**
- **Both `app` perf demo-patches silently REFUSED** (sha-drift: pinned @ v1.295.0/v1.315.0, box runs **v1.337.0**),
  leaving the un-patched 76 s members grid and the 180 s AI-readiness read live.
- **All three snapshot replays SKIPPED** (cold cache) → a structural-only catalog, and cms read content **live from
  prod over the WAN**.
- **Autoverify ended FAILING** ("1 check(s) FAILED — the stack is UP but may be non-functional").

**No number taken before this milestone lands is trustworthy.** M218 must not measure anything until M217 is green.

## Why section
The deliverables are enumerable up front — every one is a known, file:line-mapped defect with a known fix surface.
No exploratory path.

## Scope

### In
1. **Reap the leaked cockpit port** (`7700+offset`) on bring-up, and make `demo-down` reap the **whole offset port
   range** — not just the containers. (Root cause of 2 of the last 3 broken runs.)
2. **Un-swallow the demo-patch REFUSE reason.** `up-injected.sh:701,717` pipes the applier's stderr to `/dev/null`
   while the applier prints the exact sha mismatch (`apply-app-authz-skip.sh:60-61`). A refusal must be **visible**.
3. **Re-pin the two `app` perf demo-patches** + add a **patch-freshness preflight that FAILS LOUD**. The anchors
   were mechanically verified to still occur **exactly once** in the current files — so this is a **re-pin, not a
   re-authoring**. (**DEF-M215-01 / F5**, Fate-1.)
   > **A perf patch that silently degrades a demo from 5 s to 120 s is worse than a patch that refuses to apply.**
4. **Fix `jobsimulation` exits(1)** — it prints CLI help and dies on startup, so **AI-Simulations is dead in every
   demo today**. Investigate a **compose-command fix (rext-side)** before escalating. (**DEF-M215-04 / F13**,
   Fate-1.)
5. **Prime the snapshot cache on `billion`** so replay stops skipping. (**DEF-M215-02 / F9**, Fate-1. Also feeds
   M218's **C-3** hypothesis — a cold federation tier is the leading both-hero latency suspect.)
6. **Re-pin the drifted rext consumption clones** — SoT `.agentspace/rext.tag` = `v2.2`; local `stack-demo` =
   `quick-change-m211` (5 tags stale); remote = `panorama-m214-3-g41a28aa` (not even on a tag; the box warns about
   itself every run).

### Out
- **Any latency fix.** That is M218 — and M218 must **measure before it guesses**. Do not scaffold a fix here.
- Any change to `/demo-up`'s defaults (M220) or the AI-readiness render path (M219).

## Delivers → knowledge/corpus

**`corpus/ops/demo/demopatch-spec.md`** — **BLIND AREA.** The demo-patch mechanism is the **sanctioned
zero-platform-edit escape hatch that this entire release depends on**, and it has **no corpus doc**. Its 6-guard
contract (G1 path-assert, G2 pre-sha drift-refuse + single-occurrence anchor, G3 never-commit, G4 idempotent,
G5 self-revert, G6 demo-only) exists **only in a Python module docstring**; the corpus mentions it in a single
routing-table *cell*. Must document: G1–G6, the sha-vs-anchor gate decision (BD-3), the freshness preflight, and the
**re-pin runbook**.

## Open questions

**BD-3 — the demo-patch gate.** Keep the **file-sha** gate (safe; rots on every `app` bump) or move to
**anchor-only single-occurrence** matching (survives bumps; weaker drift safety)?

- **Recommendation: keep the sha gate, add an auto-repin verb + the loud preflight.**
- **Evidence it matters:** two agents computed *different* `ai_readiness.go` shas days apart (`b3216968…` @ local
  v1.334.1 vs `dc9e167e…` @ remote v1.337.0) — the pin is **stale again by the time you commit it**. A one-shot
  re-pin is a band-aid; the freshness gate is the fix.
- **Note:** ref-pinning the source clones would **NOT** stop the rot — the demo builds from a **scratch clone
  force-checked-out at the newest `v*` tag on every bring-up** (`up-injected.sh:669,679`).

## KB dependencies
- `corpus/ops/rosetta_demo.md` (the demo lifecycle) · `corpus/ops/verification.md` (the autoverify contract)
- `corpus/ops/snapshot-spec.md` + `corpus/ops/snapshot-cold-start.md` (the cache-prime path)
- `corpus/ops/demo/coverage-protocol.md` (the fix-surface routing table)
- rext: `demo-stack/patches/demopatch` (the 6 guards, in the module docstring — the thing being documented)

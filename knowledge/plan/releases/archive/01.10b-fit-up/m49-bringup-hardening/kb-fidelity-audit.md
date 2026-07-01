---
title: "KB Fidelity Audit — M49 Bring-up hardening + truth-up"
date: 2026-06-30
scope: milestone:M49
invoked-by: build-milestone
---

## Verdict
YELLOW — proceed with tracking. No blind areas; every topic PAIRED. The stale claims found are
**precisely M49's enumerated In-scope deliverables** (the truth-up), so they do not mislead the
implementation — they are the work. Tracked as KB-1..KB-4 in `decisions.md`; resolved in each section's
Phase 5 doc truth-up.

## Topic Inventory

| Topic | Knowledge doc | Code paths (rext authoring copy) | Status |
|---|---|---|---|
| #1 rext-tag source-of-truth | `corpus/ops/rosetta_demo.md`, `.claude/skills/demo-up/SKILL.md`, `corpus/ops/demo/frontend-tier.md` | `demo-stack/ensure-clones.sh`, (new) `.agentspace/rext.tag` | PAIRED (stale pins) |
| #3 .env-guard ordering | `corpus/ops/rosetta_demo.md` | `demo-stack/up-injected.sh:258` (guard) vs `:283-326` (provision) | PAIRED |
| #4 INVITATION_HMAC_SECRET | `corpus/ops/secrets-spec.md` | `stack-secrets/secretdna/secret-dna.json` (absent), `demo-stack/up-injected.sh:283-326` (provision) | PAIRED (gene absent) |
| #5 ant-academy clone | `corpus/services/ant-academy.md` | `demo-stack/ensure-clones.sh:69-90`, `stack-demo/platform/repos.yml` | PAIRED (M48 corrected the false "in repos.yml" claim) |
| #6 disk pre-flight + down image cleanup | `corpus/ops/demo/frontend-tier.md` | `demo-stack/up-injected.sh:90-115` (RAM check), `demo-stack/rosetta-demo:139-162` (cmd_down) | PAIRED |
| #7 non-fatal frontend | `corpus/ops/demo/frontend-tier.md` | `demo-stack/up-injected.sh:222` (build_frontends), `:525` (call), `:528` (compose up) | PAIRED (claim is aspirational, not true) |
| #8 demopatch re-anchor | `corpus/ops/rosetta_demo.md`, `corpus/ops/demo/frontend-tier.md` | `demo-stack/patches/next-web-studio-url/next-web-studio-url.yaml` | PAIRED (hashes anchored to v1.10 ref) |

## Fidelity Findings

1. **#1 — conflicting tag-pins (STALE).** Three prose pins disagree: `SKILL.md:84` → `storytelling-postfix-2`,
   `frontend-tier.md:254` → `storytelling-postfix-2`, `rosetta_demo.md:15` → `storytelling-postfix-1`. All three
   are stale vs the current consumption tag (`fit-up-m47`). **Fix owner:** doc — reconcile to read from the new
   `.agentspace/rext.tag` SoT (M49 #1, In-scope).
2. **#5 — ant-academy "in repos.yml" (already corrected by M48).** `services/ant-academy.md` no longer claims a
   working repos.yml clone; ensure-clones.sh:69-90 confirms the stub-sweep only covers repos.yml entries, so
   ant-academy (absent from repos.yml) is never cloned. **Fix owner:** code (M49 #5, In-scope) — doc reconcile to
   describe the real clone path.
3. **#7 — "non-fatal frontend" (STALE/aspirational).** `frontend-tier.md` describes the frontend build as
   non-fatal, but under `set -euo pipefail` a missing/failed frontend image aborts compose up. **Fix owner:**
   code (M49 #7) makes it true; doc stays/clarifies.
4. **#8 — demopatch pre_sha256 anchored to v1.10 ref (STALE).** `pre_sha256: b3d62db…` pins the pristine file at
   the v1.10 release ref; the clone is now v2.89.0. **Fix owner:** code (M49 #8) re-anchors; doc note.

## Completeness Gaps
None critical. The disk-headroom pre-flight (#6) and INVITATION_HMAC_SECRET gene (#4) are new behaviors M49
introduces; their docs land in Phase 5 of their sections (no pre-existing claim to drift from).

## Applied Fixes
None applied inline — every finding is a tracked M49 In-scope deliverable (Fate-2). Applying them now would
duplicate the milestone's own section work. Recorded as KB-1..KB-4 in `decisions.md`; resolved per-section.

## Open Items (require user decision)
None at gate. The AI-provider-keys policy is M50's call (not M49) per the milestone overview — noted, not a gate item.

## Gate Result
YELLOW: proceed with tracking. No blind areas, no claim that would mislead the implementation (the stale claims
are the milestone's deliverables). Block lifts; build-milestone Phase 1 may proceed.

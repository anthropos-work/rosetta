# M49 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## KB-fidelity tracked items (Phase 0b, YELLOW verdict)
These are stale doc claims found at the pre-flight audit. Each is a **Fate-2 already-planned** item — it is one of
M49's own In-scope deliverables, resolved in the named section's Phase 5 doc truth-up. No new backlog entry.
- **KB-1** — 3 conflicting tag-pins (`SKILL.md:84` / `frontend-tier.md:254` `storytelling-postfix-2` vs
  `rosetta_demo.md:15` `storytelling-postfix-1`, both stale vs `fit-up-m47`). Resolved by §1 (rext.tag SoT).
- **KB-2** — ant-academy "in repos.yml" already corrected by M48; doc reconciled to the real clone path in §4.
- **KB-3** — "non-fatal frontend" claim is aspirational under `set -euo pipefail`; made true + doc clarified in §6.
- **KB-4** — demopatch `pre_sha256` anchored to the v1.10 ref; re-anchored to v2.89.0 + doc note in §7.

## #5 ant-academy clone — explicit step in ensure-clones.sh, NOT repos.yml
**Context:** the platform `repos.yml` (in the EPHEMERAL gitignored `stack-demo/platform` clone) does NOT list
ant-academy, so `make init` (the canonical clone loop) never clones it and the `repos.yml`-driven stub-sweep
never sees it — the real "down in the demo" root cause (NOT the FA-token theory the old comments claimed).
**Options:** (A) edit `stack-demo/platform/repos.yml` — rejected: non-durable (overwritten on re-clone) AND a
platform-repo edit (the canonical repos.yml has no ant-academy); (B) inject ant-academy into the cloned
repos.yml at bring-up — rejected: still mutates a platform-owned file; (C) an explicit rext-owned clone step
in `ensure-clones.sh`, mirroring the `make init-studio` submodule-pattern exception. **Choice: C.** It is
durable, touches no platform file, and matches the existing precedent for "a clone make init doesn't cover."
**NON-FATAL** (ant-academy is a Vercel-native peripheral the cockpit/next-web/studio-desk carry; a clone
failure must not abort a good demo). **Not** recorded in `clones.lock.json` provenance — consistent with
cms/studio (the other explicit-clone exception), which is also not recorded. M48 already corrected the docs'
false "in repos.yml" claim; this lands the real code fix + truth-ups ant-academy.sh's stale log.

## AI-provider-keys policy — DEFERRED to M50 (Fate-2/Fate-3)
The overview lists the AI-keys policy "jointly with M50". The bring-up itself needs **no** AI keys, so M49 does NOT
provision them. The policy (which of OPENAI/ANTHROPIC/MISTRAL/ELEVENLABS/LIVEKIT become throwaway/sandbox values vs
documented-as-absent) is **M50's** call. When §3 touches `secrets-spec.md`, it leaves a one-line note pointing the
AI-keys policy at M50. Confirmed coverage — no new backlog entry.

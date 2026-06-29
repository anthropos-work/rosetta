---
milestone: M49
slug: bringup-hardening
version: v1.10b "fit-up"
milestone_shape: section
status: planned
created: 2026-06-29
last_updated: 2026-06-29
complexity: medium
delivers: corpus/ops/{rosetta_demo,demo/frontend-tier,secrets-spec}.md + corpus/services/ant-academy.md truth-up; the .agentspace/rext.tag source-of-truth
issues: demo-up #1, #3, #4, #5, #6, #7, #8
---

# M49 — Bring-up hardening + truth-up

## Goal
A from-cold `/demo-up` on a **`stack-demo`-only box** reaches set-dress + seed + verify + cockpit cleanly, and the
bring-up docs tell the truth — closing the 7 remaining demo-up issues (#2 went to M47).

## Why section
Every deliverable is a known, enumerable edit — a guard re-order, a secret-DNA gene, a `repos.yml` entry, a disk
pre-flight, a compose scale-flag, a hash re-anchor, a tag-pin file. No emergent path. (The v1.3b "M16 land-fixes"
analog — a coherent grab-bag of small bring-up fixes + doc truth-up.)

## Repo split
- **`rosetta-extensions`** (authoring copy → tag `fit-up-m49` → consume per-stack): `demo-stack/up-injected.sh`,
  `demo-stack/rosetta-demo`, `demo-stack/ensure-clones.sh`/`ant-academy.sh`, `stack-secrets/secretdna/secret-dna.json`,
  `demo-stack/patches/next-web-studio-url/`, the new `.agentspace/rext.tag` reader.
- **`rosetta`** (this corpus): the `corpus/ops/` truth-up (disjoint from M48's `architecture`+`services`).

## Scope
- **In (the 7 issues):**
  - **#3 `.env`-guard ordering** — move the `platform/.env` presence check (`up-injected.sh:258`) to **after** the
    provision step (`:293-327`), so a `stack-demo`-only box provisions-then-checks instead of aborting first.
  - **#4 `INVITATION_HMAC_SECRET`** — add as a **critical** gene in `stack-secrets/secretdna/secret-dna.json`
    (`key-present`+`nonempty`); **auto-generate a throwaway non-prod value** at provision (so the pre-flight catches
    its absence instead of the backend dying silently at runtime — the silent `app Exited (0)` class).
  - **#5 ant-academy clone** — add ant-academy to `platform/repos.yml` (or an explicit clone in
    `ensure-clones.sh`/`ant-academy.sh`); reconcile the docs (retire the false "RESOLVED" — `make init` skips what
    `repos.yml` omits, so the postfix-2 stub-sweep never re-cloned it).
  - **#6 disk pre-flight + `demo-down` image cleanup** — add a **disk-headroom pre-flight** (warn + offer
    `docker system prune`, non-blocking — mirror the RAM check at `up-injected.sh:91-115`); have `cmd_down`
    (`rosetta-demo:139-162`) **remove the stack's images** (`demo-N-next-web`, `demo-N-studio-desk`) so dead stacks
    don't accumulate (the ENOSPC contributor).
  - **#7 non-fatal frontend** — when a frontend image is **absent**, drop that service from `compose up` / scale it
    to **0** (post-`build_frontends()` image-existence check before `up-injected.sh:532`), so a failed UI build no
    longer aborts backend + set-dress + verify + cockpit under `set -e`. Makes the documented "non-fatal" *true*.
  - **#8 demopatch re-anchor** — re-anchor `patches/next-web-studio-url/next-web-studio-url.yaml` `pre_sha256` /
    `post_sha256` to the **current (M47-synced) next-web source**, so the manager's "Studio" nav link rewrites to
    the demo's offset studio-desk instead of staying baked to prod.
  - **#1 single rext-tag source-of-truth** — introduce **`.agentspace/rext.tag`** that the `/demo-up` skill +
    `ensure-clones.sh` read to check out the pinned tag; reconcile the 3 conflicting prose pins
    (`SKILL.md:84`, `frontend-tier.md:254`, `rosetta_demo.md:15` — the last even says `storytelling-postfix-1`) to
    read from it. **Doubles as the note-2 reproducible-pin mechanism** (the pin is a file, not "whatever was last
    checked out").
  - **AI-provider-keys decision** (jointly with M50): which of OPENAI/ANTHROPIC/MISTRAL/ELEVENLABS/LIVEKIT become
    throwaway/sandbox values in the demo secret source vs **documented-as-absent** — record + apply the policy in
    `secrets-spec.md`.
- **Out:** the clone re-sync (M47); content seeding (M50/M51); corpus architecture/service re-ground (M48); the
  manifest (M52).

## Depends on
**M47** (#8 anchors to the synced next-web hash; #3/#4 verified on a current bring-up). **Parallel with:** **M48**
(disjoint clusters; M49 owns the live demo, M48 doesn't).

## Open questions (resolve during build)
- AI-provider-keys policy — which are demo-credible throwaway/sandbox values vs documented-as-absent (security).
- ant-academy via `repos.yml` vs an explicit clone step — *lean:* `repos.yml` (makes `make init` the single owner).
- `.agentspace/rext.tag` format — bare tag string vs a small keyed file. *Lean:* bare string, one line.

## KB dependencies (read as contract)
- `corpus/ops/rosetta_demo.md`, `corpus/ops/demo/frontend-tier.md`, `corpus/ops/secrets-spec.md`,
  `corpus/ops/snapshot-cold-start.md`, `corpus/ops/idempotency.md`, `corpus/services/ant-academy.md`.

## Delivers
- **→ rosetta-extensions:** the 7 fixes + `.agentspace/rext.tag`, tagged `fit-up-m49`.
- **→ rosetta:** `rosetta_demo.md` (the rext.tag source-of-truth + ordering), `demo/frontend-tier.md` (disk
  pre-flight + true non-fatal frontend + demo-down cleanup), `secrets-spec.md` (`INVITATION_HMAC_SECRET` critical +
  AI-keys policy), `services/ant-academy.md` (clone reconciliation).

## Risk
**(scope)** a 7-issue grab-bag is prone to growth. **Hard line:** only these enumerated issues; anything the M53
cold rebuild surfaces routes back to its owning milestone (or defers), it does not expand M49.

---
title: "KB Fidelity Audit — M42e employee 100% coverage"
date: 2026-06-25
scope: milestone:M42e
invoked-by: build-mstone-iters
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Demo UI tier (next-web + studio-desk + ant-academy, offset ports, baked pk/URLs) | `corpus/ops/demo/frontend-tier.md` | `demo-stack/up-injected.sh`, `stack-injection/gen_injected_override.py` | PAIRED |
| Bring-up auto-verify net (offset/project/scope-aware probes) | `corpus/ops/verification.md` | `stack-verify/live/autoverify.sh`, `stack-verify/live/verify.sh`, `stack-verify/lib/target.sh` | PAIRED |
| Demo lifecycle + Clerkenstein injection + offset ports | `corpus/ops/rosetta_demo.md` | `demo-stack/up-injected.sh`, `stack-injection/` | PAIRED |
| Cockpit-handshake browser login (fake FAPI, minted pk, mkcert TLS) | `corpus/ops/demo/recipe-browser-login.md` | `clerkenstein/clerk-frontend/`, `demo-stack/up-injected.sh` | PAIRED |
| Verified-skill chain + roster hero (Maya) | `corpus/ops/demo/stories-spec.md` | `stack-seeding/seeders/{roster,persona,profile,users}.go` | PAIRED |
| Fix surface: stack-seeding (empty section) | `corpus/ops/seeding-spec.md` | `stack-seeding/seeders/` (ProfileSeeder, roster org_name) | PAIRED |
| Fix surface: stack-snapshot serve-grants (federation/content) | `corpus/ops/snapshot-spec.md` | `stack-snapshot/directus/structure.go` (directus_versions + library-category synth grants) | PAIRED |
| Fix surface: demo injection + env link-rewriting (escapes) | overview.md (declared deliverable) + `frontend-tier.md` (baked-URL mechanism it extends) | `demo-stack/up-injected.sh` (NEXT_PUBLIC_* build-args), `stack-injection/gen_injected_override.py` | DOC-ONLY (planned deliverable — link-rewrite not yet built; sweep-discovered) |
| **Iteration protocol (the sweep + triage + fix loop)** | `corpus/ops/demo/coverage-protocol.md` (**NEW** — authored by iter-01) | the Playwright coverage harness (NEW rext section, this milestone) | DOC-ONLY (declared `Delivers`; authored in iter-01) |

## Fidelity Findings

1. **Demo UI tier baked-URL claim** — `frontend-tier.md` says next-web bakes `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` / `_BACKEND_API_URL` / `_HOSTING_URL` to offset ports. **Verified ALIGNED** against `up-injected.sh:153-155` (exact ARGs). No studio-host URL ARG is baked — consistent with the doc (it lists only the 3 URLs + pk). The escape risk (left-menu "Studio" → prod) is therefore real and un-rewritten — exactly the gap the sweep targets. ALIGNED.
2. **G6 serve-grants (directus_versions + library-category)** — `snapshot-spec.md` / the M40 record describe synthesized public-read `directus_permissions` rows. **Verified ALIGNED** against `stack-snapshot/directus/structure.go` + `serve_test.go` (the synthesized `directus_versions` grant + library-category closure are in code with dedicated tests). ALIGNED.
3. **G1 org_name threading** — `stories-spec.md` / M39 record describe threading `st.Org.Name` through the roster into the FAPI org resource. **Verified ALIGNED** against `roster.go:48-111` (`OrgName` field + `st.Org.Name`) and `clerk-frontend/resources.go:32,234` (the no-roster default fallback). ALIGNED.
4. **G3/G5 ProfileSeeder (M41)** — `stories-spec.md` describes the work-history/education timeline + claimed-but-unverified tail. **Verified ALIGNED** against `stack-seeding/seeders/profile.go` + `profile_write.go` (present, with tests). ALIGNED.
5. **Auto-verify offset/project/scope plumbing** — `verification.md` describes `target.sh` offset/project/scope resolution + the existing `stack-verify/e2e/` Playwright smoke harness. **Verified ALIGNED**: `stack-verify/e2e/{package.json,playwright.config.ts,tests/smoke.spec.ts}` exist with Playwright `^1.49.0` already pinned (the coverage harness extends this, not a from-scratch new dependency). ALIGNED.

## Completeness Gaps

1. **`corpus/ops/demo/coverage-protocol.md` does not exist yet** — this is NOT a blind-area blocker: the milestone's `overview.md` `Delivers →` list and `iteration_protocol_ref` both declare it a M42e deliverable, authored in iter-01 (the bootstrap tok). Per the gate rule, a topic the milestone explicitly delivers as knowledge is covered. (incidental/by-design)
2. **Link-rewriting fix surface (escapes) not yet built** — no studio-host URL override in the demo build/injection. By-design: it's a sweep-discovered fix surface, a declared milestone deliverable, built when/if the sweep finds an escape. The mechanism it extends (per-stack URL baking) is fully documented in `frontend-tier.md`. (planned deliverable, not a gap)

## Applied Fixes
None required — all PAIRED topics ALIGNED; the two DOC-ONLY items are declared milestone deliverables, not blind areas.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed. Every topic the milestone reads-as-truth is PAIRED + ALIGNED. The two DOC-ONLY items (`coverage-protocol.md`, the link-rewriting surface) are explicit M42e deliverables per the overview, not blind areas. The bootstrap tok may author its strategy against verified knowledge docs.

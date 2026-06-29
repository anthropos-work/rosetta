---
title: "KB Fidelity Audit — M26 self-contained-demo"
date: 2026-06-15
scope: milestone:M26
invoked-by: build-milestone
---

## Verdict
GREEN

M26 is a **doc-vs-code reconciliation** milestone: the rosetta corpus already documents `stack-demo` as a peer
workspace with "its cloned platform service repos" (CLAUDE.md stack-*/ table), but the code (`up-injected.sh`)
builds from `stack-dev`. M26 makes the code match the documented model AND re-authors the doc-half to describe
the now-real self-contained flow. The implementation's contract is the **verified orphan diff** (ext repo
@ `25ab855`, tag `prop-room-m26`) + the rosetta doc-half (§7), not a third-party spec doc — so there is no blind
area: every implemented topic is anchored either in the orphan (the spec) or in the doc-half this milestone
delivers. The rosetta docs that describe the OLD (stack-dev-sourced) flow are DOC-ONLY findings that §7 updates
by design — they are deliverables, not stale load-bearing claims the implementation reads as truth.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| peer-clone bootstrap (ensure-clones) | `GUIDE.md` (delivered §6) + orphan spec | `demo-stack/ensure-clones.sh` (NEW) | DOC-ONLY → delivered |
| demo build SOURCE = stack-demo | `corpus/ops/rosetta_demo.md`, `demo/README.md`, `frontend-tier.md` (§7) | `up-injected.sh`, `migrate-demo.sh`, `ant-academy.sh`, `rosetta-demo` | PAIRED (doc describes OLD flow → §7 updates) |
| `reuse_dev_images` gate | `frontend-tier.md` + orphan spec | `gen_injected_override.py` | PAIRED |
| stack-demo "own clone set" / true peer | `CLAUDE.md` stack-*/ table | (M26 makes true) | DOC-AHEAD (the gap M26 closes) |
| M30 secret provision (must-preserve) | `corpus/ops/secrets-spec.md`, `stack-secrets/SKILL.md` | `up-injected.sh` L197-239, `stack-secrets/` | PAIRED |
| M31 mkcert FAPI cert (must-preserve) | `frontend-tier.md`, `rosetta_demo.md` | `up-injected.sh` FapiCertStep region | PAIRED |
| M32 studio-desk single-port (must-preserve) | `frontend-tier.md` | `gen_injected_override.py` (disjoint) | PAIRED |

## Fidelity Findings

1. **`corpus/ops/demo/frontend-tier.md:162`** — `cd stack-dev/ant-academy/code` in the documented manual
   fallback. **Verdict: STALE-after-implementation.** M26 repoints `ant-academy.sh` to `stack-demo/ant-academy/code`.
   **Fix owner: doc** — §7 of this milestone (the doc-half re-author). NOT a blocker: the implementation reads the
   orphan spec, not this prose; the prose IS the deliverable being updated.
2. **`corpus/ops/demo/README.md:135`** — "The platform clones are consumed as-is" describes the build source
   generically without naming stack-demo. **Verdict: incomplete (not wrong).** §7 adds the ensure-clones step-0 +
   the stack-demo-own-clones source. **Fix owner: doc** — §7.
3. **`CLAUDE.md` stack-*/ table (L115-122)** — describes `stack-demo` as holding "its cloned platform service
   repos." **Verdict: DOC-AHEAD of code on `main`** (the exact gap M26 closes). After M26 the code delivers it.
   **Fix owner: code (the milestone) + a CLAUDE.md reconcile** — §2/§4 (code) + §7 (CLAUDE.md note).

None of findings 1–3 is a blind area or a stale-claim-the-implementation-reads-as-truth. All three are the
planned doc-half work (§7) of a reconciliation milestone.

## Completeness Gaps

None critical. The `reuse_dev_images` opt-in (D5) is a new flag — its doc surface lands in §6 (GUIDE) + §7
(frontend-tier). The `DEMO_REUSE_DEV_IMAGES=1` env knob joins the documented `DEMO_NO_*` family. Flagged for
the §2/§3/§6/§7 CLI-flag↔docs consistency check at commit time.

## Applied Fixes
None applied pre-build (the doc-half is the milestone's own §7 deliverable, authored during the build loop, not
a pre-flight patch). The topic→doc→code triples are recorded in `spec-notes.md`.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed. The milestone's implementation contract is the verified orphan diff + the GUIDE/doc-half it
delivers; no blind area, no load-bearing stale claim the implementation would read as truth. The DOC-ONLY/
DOC-AHEAD findings are §7 deliverables, tracked in progress.md.

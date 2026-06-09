---
title: "KB Fidelity Audit — M16 Land the field fixes + restore doc truth"
date: 2026-06-08
scope: milestone:M16
invoked-by: build-milestone
---

## Verdict
GREEN

M16 is a *doc-truth-restoration* milestone: most of its deliverables ARE the act of correcting stale docs.
The audit's job here is to confirm (a) the code reality M16 documents against is what M16 assumes, and (b) no
topic M16 touches is a blind area (no anchor at all). Both hold. The stale claims M16 will fix (78 tests,
"no remote", `/demo-status`, v1.1/M3, `anthropos-dev/platform` as default) are precisely the deliverables —
they are *known* stale, enumerated in the milestone overview + `.agentspace/demo-up-issue.md`, and fixing them
is the milestone, not a surprise. That makes them in-scope work, not a blocking pre-flight finding.

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| stack-dev workspace layout + back-compat fallback | `corpus/ops/rosetta_demo.md` (+ CLAUDE.md establishes `stack-dev/`) | `.agentspace/rosetta-extensions/{demo-stack/up-injected.sh,migrate-demo.sh,rosetta-demo; dev-stack/dev-stack}` (the `stack-dev → anthropos-dev` fallback lines) | PAIRED (note in corpus is net-new deliverable) |
| The two applied fixes (devpath ISSUE-1, migrate-race ISSUE-7) | `.agentspace/demo-up-issue.md` ISSUE-1/ISSUE-7 | extensions commits `547de17`, `ed72e94` | PAIRED |
| demo-stack GUIDE/README header facts | `demo-stack/GUIDE.md`, `demo-stack/README.md` | `demo-stack/tests/test_tooling.py` (13 tests); `origin` remote | PAIRED (facts STALE — M16 fixes) |
| pytest invocation doc | `demo-stack/GUIDE.md:166-167` | the `pytest` entrypoint | PAIRED (STALE — M16 fixes) |
| extensions consumption/version model | `.agentspace/rosetta-extensions/knowledge/README.md` | the tag set / `stack-demo/rosetta-extensions` pinned tag | PAIRED (version-jump note net-new) |
| residual anthropos-dev in corpus | `corpus/` | n/a | DOC-ONLY (already clean — 0 refs) |

## Fidelity Findings

1. **demo-stack/GUIDE.md:161 — "78 unit tests" (55 + 23)** — Expected: 78 across `test_tooling.py` (55) + `test_inject_scripts.py` (23). Actual: `test_inject_scripts.py` no longer exists; `pytest tests/` collects/passes **13** in `test_tooling.py` alone. Verdict: STALE. Fix owner: update doc (M16 deliverable ISSUE-3/4). Win: doc → match code.
2. **demo-stack/GUIDE.md:5 — "no remote"** — Expected per doc: no git remote. Actual: `origin → https://github.com/anthropos-work/rosetta-extensions.git` (ls-remote succeeds; `origin/main` at `a31d70b`). Verdict: STALE. Fix: doc (M16 ISSUE-3).
3. **demo-stack/GUIDE.md:5 — `/demo-status`** — renamed `/stack-list` in M14 (CLAUDE.md confirms). Verdict: STALE. Fix: doc (M16 ISSUE-3).
4. **demo-stack/GUIDE.md:5 — "v1.1 'show floor' / M3"** — repo tagged through `v1.3.0`/`stack-party-m15`; v1.3b in dev. Verdict: STALE. Fix: doc (M16 ISSUE-3).
5. **demo-stack/GUIDE.md:167 — `python3 -m pytest tests/ -v`** — fragile across environments (Homebrew `python3` may lack the pytest module; verified here `python3`=3.9.6 with no pytest module, while the `pytest` entrypoint=8.4.2 works). Verdict: STALE/fragile. Fix: doc → `pytest tests/ -v` + a 3.11/3.12 note (M16 ISSUE-4).
6. **Prose `anthropos-dev/platform` as the DEFAULT path** at `demo-stack/GUIDE.md:17`, `demo-stack/README.md:12`, `dev-stack/README.md:73`, `stack-core/gen_override.py:4`. Code default is `stack-dev` (fallback `anthropos-dev`). Verdict: STALE. Fix: doc → `stack-dev` (M16 ISSUE-2). The fallback lines in the scripts (the single intentional legacy-alias mention) stay.

All six STALE findings ARE M16's enumerated deliverables — not blockers, in-scope corrections.

## Completeness Gaps

1. **stack-dev layout note in `corpus/ops/`** — no consolidated note today on the `stack-*/` workspace layout + the back-compat fallback, cross-linked from `rosetta_demo.md`. NET-NEW M16 deliverable (not a blind area — the topic is well-anchored by CLAUDE.md + the existing ops docs that already use `stack-dev/`). Classify: deferred-to-this-milestone (M16 delivers it).
2. **Version-jump expectation (v1.2.1 → fix tag)** — extensions `knowledge/README.md` carries no consumption-version note today; M16 adds the version-jump note (ISSUE-5). NET-NEW; not a blind area.

## Applied Fixes
None applied in-audit — every STALE finding is an explicit M16 deliverable that lands in Phase 1 (correcting them here would erase the milestone's own work). The triples above are recorded for the build to start fast.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to build. Every topic is anchored (no blind areas); every stale claim is a known, enumerated
M16 deliverable; the code reality the milestone documents against was verified (13 tests, remote exists,
`stack-dev` default with `anthropos-dev` fallback, corpus already free of `anthropos-dev`).

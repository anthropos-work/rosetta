---
milestone: M0
slug: alignment-framework
version: v1.0 "body double"
milestone_shape: section
status: in-progress
started: 2026-06-02
---

# M0 — Alignment measurement framework

## Goal

A reusable, engine-agnostic process — two skills + an executable harness + a test class + a manifest
format — that measures how faithfully any **mirror** engine reproduces any **source** engine,
producing a **0–100% alignment score**. This is the foundation M1 (build Clerkenstein) builds on and
M1b (Clerk drift detection) reuses.

## Context

Clerkenstein is one instance of a general pattern: mirror an external engine so a consumer can't tell
the difference. The way you *prove* a mirror is faithful is **alignment testing** — a third class of
test beside unit and integration: a differential test that feeds identical input to two targets
(source + mirror) and asserts behavioral equivalence. M0 extracts this measurement machinery so M1
isn't a hand-built mock with ad-hoc parity tests, but the first mirror produced by a repeatable,
scored process. (Roadmap: `knowledge/plan/roadmap.md` → v1.0 "In Development".)

## Scope

### In
- **`/align-dna` skill** (build & update alignment targets): given a source framework + version, pull
  the pinned source into `.agentspace/`; enumerate the **consumed** capabilities (scoped to what the
  consumer calls, not the whole SDK); enumerate standard + corner-case **variants** per capability;
  emit/update the **Alignment DNA** manifest (each gene: input fixture, expected-shape descriptor,
  equivalence operator, criticality weight); **diff DNA across source versions**; **scaffold
  alignment-test stubs from the DNA**.
- **`/align-run` skill** (measure alignment of 2 targets): given a DNA + source version + mirror, run
  every gene against **both** targets, assert equivalence per the gene's operator, compose the
  **0–100% score** + a per-capability divergence report.
- The **alignment test-class convention** (tagging so tests are discoverable + countable, distinct
  from unit/integration), the **DNA file format** (JSON), the **equivalence operators**
  (`exact` / `shape` / `normalized` / `error_class`), and **record/replay (golden capture)** so a
  live-SaaS source is measured reproducibly offline.
- The executable **`alignctl`** reference harness (stdlib-only Go, under `test/alignment/`) that the
  skills drive.
- A **toy reference mirror** (≥2 capabilities, with one intentional divergence) that compiles, runs,
  and yields a real score — proving the framework end-to-end *including its ability to catch
  misalignment*.

### Out
- The Clerk DNA + the real Clerkenstein mirror → **M1**.
- Drift CI wiring → **M1b**.
- The JS/browser surface → **M2**.
- Mirroring engines other than the toy (the framework is generic, but M0 only exercises it on the toy).

## Sections
See `progress.md`. Order: contract → harness → toy (these three co-define and must run together) →
skills → doc. The doc is authored last so it describes the *actual* working contract, not an
aspirational one.

## Delivers → knowledge
- `corpus/architecture/alignment_testing.md` — the canonical reference (net-new; alignment is a
  documentation blind area today).

## Exit (section checklist complete) when
- `alignctl run` against the toy emits a correct 0–100% score and a divergence report that catches
  the toy's intentional divergence;
- both skills exist and their documented invocations match `alignctl`'s actual flags;
- `corpus/architecture/alignment_testing.md` describes the working contract and is discoverable from
  the corpus index.

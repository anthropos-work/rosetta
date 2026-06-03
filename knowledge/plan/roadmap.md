# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills).

> **Designed 2026-06-02** from the Demo Environment + Clerkenstein brief, **refined 2026-06-02** to
> promote alignment measurement into a first-class discipline (new **M0**). 3 research agents over the
> Clerk integration, the staging/dev-env tooling, and the data/seeding surface — all verified against
> the cloned platform in `anthropos-dev/`. Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-02.md`](../../.agentspace/scratch/roadmap-research-2026-06-02.md).
>
> **Active version: v1.0 "body double"** — *designed, not yet branched.* The `release/01.00-body-double`
> branch and milestone dirs are **not** created yet (Phase 8 deferred by user choice). Next action:
> scaffold v1.0 + cut the branch, or run the build skills (which cut it on demand) starting at **M0**.

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.0** | **body double** | A *measured* stand-in the platform can't tell from the real thing | M0 → M1 → { M1b ∥ M2 } | **active (designed)** |
| v1.1 | show floor | Disposable, richly-seeded demo stacks on demand | M3 → M4 → M5 | next ([roadmap-vision.md](roadmap-vision.md)) |

The whole initiative layers a **second corpus + skill set on top of** the existing dev-environment
tooling, to build disposable demo environments. Hard constraints: **no modification to any platform
repo** (current or future) and **no disruption to the dev environment** — demo clones live under the
gitignored `anthropos-demo/` (mirroring `anthropos-dev/`). Full brief:
[`.agentspace/demo-environment-draft.md`](../../.agentspace/demo-environment-draft.md).

## In Development — v1.0 "body double"

**Theme:** Clerk authentication is the friction that blocks fast, throwaway demos. v1.0 delivers
**Clerkenstein** — a drop-in mock that mirrors the exact Clerk interface the platform uses, with
security/sync disarmed, injected via build-time `go.mod replace` + skip-worktree so **every platform
repo keeps "thinking" it uses Clerk with zero source changes**. The novelty: Clerkenstein isn't a
hand-built mock — it's the **first mirror produced by a reusable, measurable alignment process**
(M0). We don't just claim the stand-in is faithful; we **score** it (0–100%) against the real Clerk
and CI-gate that score against drift. This also removes Clerk's API rate limit as the blocker for
scale data-seeding in v1.1.

**Decided at design (2026-06-02):** two-version split (Clerkenstein first); **alignment is a
first-class test class** with its own framework (M0); **M1 is iterative** (its exit gate is an
alignment score); M2 frontend = attempt the fake Clerk FAPI server, **fall back to the real dev Clerk
app for the browser session** (backend stays fully mocked) if base-URL override proves too fragile.

### Alignment vocabulary (the M0 model, referenced by M1/M1b)
- **Target** — an engine exposing a surface. **Source target** = the canonical engine, version-pinned (Clerk `clerk-sdk-go/v2 @ v2.6.0`). **Mirror target** = our reimplementation (Clerkenstein).
- **Capability** — one endpoint/function of the source surface *(axis 1)*. **Variant** — one input/scenario class for a capability (standard + corner + error) *(axis 2)*.
- **Alignment test** — one **(capability × variant)** pair; feeds identical input to both targets and asserts behavioral equivalence. A **third test class** alongside unit & integration; **tagged** so it's parseable/countable/runnable as its own suite.
- **Alignment DNA** — the officially-enumerated complete set of (capability × variant) **genes** for a source target at a version; the machine-readable manifest that *defines* faithfulness and is the score's denominator.
- **Alignment score** — `aligned genes ÷ total genes × 100`, with a per-capability rollup. 100% = behaviorally indistinguishable across the whole DNA.

### M0: Alignment measurement framework
**Status:** `done` (2026-06-02)
**Shape:** `section`
**Goal:** A reusable, engine-agnostic process — two skills + a test class + a manifest format — that measures how faithfully any mirror reproduces any source engine, producing a 0–100% alignment score. This is the foundation M1 builds on and M1b reuses.

**Closed 2026-06-02** (build S1–S5 → harden 2 passes → close review → merged to `release/01.00-body-double`). Delivered: `test/alignment/` — `alignctl` (stdlib-only Go, builds/runs offline) with `run`/`capture`/`dna list|diff|validate`, engine-agnostic via a pluggable `--runner`; the 4 equivalence operators + weighted score (overall + separate critical gate); record/replay goldens; `internal/canon` precision-safe canonicalization; the `//go:build alignment` test class; and a toy reference proving end-to-end detection (**86.7% / 100% critical**, catches `Greet/padded-name`). Plus `/align-dna` + `/align-run` skills and `corpus/architecture/alignment_testing.md`. Open questions resolved: DNA format = **JSON** (M0-D1); capabilities enumerated from the consumed surface; goldens live per-mirror-repo. Close-review adversarial pass found + fixed a path-traversal must-fix + score-overflow (M0-D7); 45 test funcs (3 fuzz), 5/5 flake gate, core coverage 83–98%. Decisions M0-D1…D7. Retro: [m0-alignment-framework/retro.md](releases/01.00-body-double/m0-alignment-framework/retro.md). Resolved repo split: framework in rosetta; the Clerk DNA/tests/mirror land in the `clerkenstein` repo (M1).
**Scope:**
  - In:
    - **`/align-dna` skill** (build & update alignment targets): given a source framework + version, pull the pinned source into `.agentspace/`; enumerate the **consumed** capabilities (scoped to what the platform calls, not the whole SDK); enumerate standard + corner-case **variants** per capability; emit/update the **Alignment DNA** manifest (each gene: input fixture, expected-shape descriptor, equivalence operator, criticality weight); **diff DNA across source versions** (added/removed/changed genes); **scaffold alignment-test stubs from the DNA** so tests never drift from the manifest.
    - **`/align-run` skill** (measure alignment of 2 targets): given a DNA + source version + mirror, pull the source, run every gene against **both** targets, assert equivalence per the gene's operator, compose the **0–100% score** + a per-capability divergence report.
    - the **alignment test-class convention** (tagging/marking so tests are discoverable + countable, distinct from unit/integration), the **DNA file format**, the **equivalence operators** (exact / same-shape / normalized / same-error-class), and **record/replay (golden capture)** support so a live-SaaS source can be measured reproducibly offline.
    - a **tiny toy reference mirror** (≈2 capabilities) proving the framework runs + scores end-to-end, independent of Clerk.
  - Out: the Clerk DNA + the real Clerkenstein mirror (M1); drift CI wiring (M1b); the JS surface (M2).
**Depends on:** none.
**Parallel with:** none (gates M1, M1b).
**Estimated complexity:** large
**Open questions:** DNA manifest format (YAML vs Go structs); how capabilities are enumerated (parse source surface vs curated list); where golden captures live + how they're refreshed.
**KB dependencies:** none new (greenfield — alignment is a documentation blind area).
**Delivers → `corpus/architecture/alignment_testing.md`:** the alignment test class, the DNA format, the two skills, equivalence + record/replay — the canonical reference (net-new doc).

### M1: Clerkenstein backend mirror (Go)
**Status:** `done` (2026-06-03, **closed-on-gate**)
**Shape:** `iterative`
**Goal:** The first real mirror — a drop-in Go stand-in for `colony/authn`'s provider + the Clerk `orgclient`, built *by* the M0 process and injected via `go.mod replace` (zero platform-repo edits), so backend services authenticate with one universal credential and locally-minted JWTs.

**Closed 2026-06-03** (5 iters: bootstrap tok TOK-01 → DNA → authn twin → critical orgclient → standard orgclient → **gate** → final harden → close). The **Clerkenstein backend mirror** (in the gitignored `anthropos-demo/clerkenstein` repo, its own git) scores **100% alignment / 100% critical** against the `clerk@2.6.0` DNA (22 genes), built offline. authn implements the real `colony/authn.Provider` (HS256, one universal key); orgclient is a disarmed in-memory twin. Score arc: 0 → 21.1 → 68.4 → **100%**. Final harden: authn + orgclient **0 → 100%** unit coverage (+1 fuzz, 0 bugs). Decisions: D1 hybrid goldens; iter-01-D1 authn injects via `go.mod replace` whole-colony; **M1-D2 orgclient injects via a fake-Clerk-API-server → routed to M2** (shared HTTP-interception with the JS side). Delivered `corpus/services/clerkenstein.md`. Retro: [m1-clerkenstein-backend/retro.md](releases/01.00-body-double/m1-clerkenstein-backend/retro.md). The gate (alignment fidelity) is met; live injection into a running platform is demo-stack work (v1.1) / M2 (orgclient).
**Exit gate:** `/align-run` reports **100% alignment on the platform-consumed Clerk Go surface (critical capabilities) and ≥95% overall**, with any waived genes documented + justified in the divergence report.
**Iteration protocol:** `corpus/architecture/alignment_testing.md` (the M0-delivered alignment-measurement process) — the measure → fix-diverging-genes → re-measure loop.
**Why iterative (not section):** the deliverables are writable, but *which genes diverge and how costly each is to close* only emerges from measuring against the real Clerk — a fixed up-front checklist would be speculative. The score is the commitment; the path to it is open.
**Depends on:** M0 (its skills + DNA format + test class).
**Parallel with:** none (gates M1b and M2).
**Estimated complexity:** large
**Re-scope trigger:** if consecutive strategy iters (toks) can't close a diverging gene (e.g. a capability that's fundamentally unmockable offline), waive it with justification or escalate to the user — don't chase an unreachable 100%.
**Open questions:** which capabilities need live-Clerk record/replay vs pure local mint; the precise critical-capability set; stub just `authn`+`orgclient` or `replace` all of `colony` (`authn` is a package inside `colony`) — fallback is vendoring whole `colony`, as staging already does.
**KB dependencies:** `corpus/architecture/alignment_testing.md` (the iteration protocol), `corpus/services/clerk-integration.md`, `corpus/architecture/shared_libraries.md` (§ authn/colony), `corpus/ops/staging-clerk.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`** (the mirror design + injection mechanism — net-new) **+ the Clerk Alignment DNA** (`clerk@2.6.0` genome, authored via `/align-dna`).

### M1b: Clerk drift detection
**Closes the gap after:** M1 (Clerkenstein is aligned at v2.6.0 — but must *stay* aligned as the platform bumps `clerk-sdk-go` / `@clerk/*`).
**Goal:** Reuse M0 wholesale to make Clerk drift a flagged, mechanical event: on a version bump, `/align-dna` diffs the DNA (what changed) and `/align-run` re-scores the existing mirror against the new source (score drop = broken genes), CI-gated on "alignment ≥ threshold."
**Scope:**
  - In: the "bump pinned Clerk version → DNA-diff → re-score → report" workflow; the CI gate on alignment score; golden-capture refresh on bump.
  - Out: building the framework (M0); authoring the original mirror/DNA (M1); the JS surface (M2 owns its own genes).
**Depends on:** M1 (needs a built, aligned mirror + the Clerk DNA). Reuses M0's skills — **now automation/config over M0, not new machinery** (the right size for a B-milestone).
**Parallel with:** M2 (CI/automation vs JS code — disjoint surfaces).
**Acceleration effect:** every future Clerk bump becomes a flagged, scored update instead of a silent break — the brief's "follow platform updates within minutes" requirement, mechanized.

### M2: Clerkenstein — browser session + webhook coherence (JS)
**Status:** `planned`
**Shape:** `section`
**Goal:** The frontend logs in with no real Clerk, and created/seeded users/orgs reach the DB without real Clerk webhooks.
**Scope:**
  - In: a fake Clerk FAPI path for `@clerk/nextjs ^6.39.2` (next-web-app, ant-academy) and `@clerk/clerk-js ^5.52.3` (studio-desk) via publishable-key + base-URL/DNS override — **with the decided fallback**: keep the real dev Clerk app for the browser session while the backend stays fully mocked; a **webhook injector** feeding the existing `app/internal/clerk/events/` sync pipeline directly; **the JS surface's fidelity expressed as alignment genes via M0** where applicable (same score treatment as the Go side).
  - In (**routed from M1 close — M1-D2, Fate 3**): the **fake-Clerk-API-server** (HTTP interception of `api.clerk.com`) ALSO serves M1's **orgclient** injection — the Go `app/internal/clerk/orgclient` is app-internal + networked, so it can't `go.mod replace` like authn; it disarms via the same fake-API-server this milestone builds for the JS side. The Clerkenstein orgclient mirror behavior already exists + scores 100% (M1); M2 wires the HTTP redirect that makes the platform's real orgclient hit it.
  - Out: multi-instance stacks (M3); data seeding (M4).
**Depends on:** M1 (consumes the mock contract + minted-token shape). **Parallel with:** M1b (yes).
**Estimated complexity:** large — **highest technical risk in v1.0** (SDKs hard-code Clerk FAPI; no documented base-URL override).
**Open questions:** can `@clerk/*` be pointed at a fake FAPI without a fork? (the fallback exists because this is uncertain) — spike the override early.
**KB dependencies:** `corpus/architecture/alignment_testing.md`, `corpus/services/clerk-integration.md`, `corpus/architecture/frontend_architecture.md`, `corpus/services/next-web-app.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`:** extends the M1 doc with the JS path + webhook injection + the fallback decision.

### Execution graph

```
v1.0 "body double"   — a stand-in the platform can't tell apart, and we can prove it

  M0 (alignment framework: /align-dna + /align-run, test class, DNA format, golden capture, toy ref)
    │
    ↓
  M1 (Clerkenstein backend mirror — ITERATIVE: author Clerk DNA → drive alignment score to gate)
    │
    ├──→ M1b (Clerk drift detection — DNA-diff + re-score, CI-gated across version bumps)   ∥ M2
    └──→ M2 (browser session + webhook; reuses the alignment class for the JS surface)
```

### Parallelism

- **M0 → M1 → {M1b, M2}** sequential at the core: M1 needs M0's framework; M1b + M2 need M1's mirror/contract.
- **M1b ∥ M2:** disjoint surfaces — M1b is CI/automation over M0; M2 is JS + the webhook injector. Merge risk **low**.
- **M3 ∥ M2 (cross-version, yes-with-caveats):** sequenced cleanly by the version boundary (M3 starts after v1.0 closes).

### Risks (v1.0)

| Risk / decision | Severity | Mitigation |
|---|---|---|
| **Source is a live SaaS** — Clerk's API capabilities can't be hit freely/offline/deterministically | blocks-release (reproducibility) | M0 **record/replay golden captures** is a core requirement, not an afterthought — capture once, replay forever |
| **DNA completeness gaming** — 100% on a thin DNA is hollow | degrades-quality | `/align-dna` capability-coverage check (every platform-consumed endpoint present) + M1b version-bump DNA-diff keeps it complete |
| **Defining "equivalent"** — timestamps, generated IDs, error formats differ even when behavior matches | degrades-quality | M0 ships **equivalence operators** (exact / same-shape / normalized / same-error-class) chosen per gene |
| **JS/FAPI fake server** — SDKs hard-code Clerk FAPI, no base-URL override | blocks-release (full no-Clerk browser) | **Decided fallback:** real dev Clerk app for the browser, backend fully mocked; spike override early in M2 |
| **`colony` replace granularity** — `authn` is a package inside `colony`, not its own module | degrades-quality (M1 effort) | M1 early iter resolves it; fallback = vendor whole `colony` (staging precedent) |
| **Repo layout** — where the framework vs the Clerk mirror live | nice-to-resolve | **Decided:** the M0 framework (skills + format + doc) lives in rosetta; the Clerk DNA + alignment tests + mirror live in the `clerkenstein` repo, cloned into gitignored `anthropos-demo/` |
| **"Zero platform-code changes" interpretation** — `replace` edits the *clone's* go.mod | nice-to-resolve | build-time injection in the gitignored clone + skip-worktree; upstream repo never modified (same as staging's `vendor-colony/`) |

### Branch model

`release/01.00-body-double` from `main` (to be created at scaffold time — Phase 8, deferred). Milestone
branches: `m0/alignment-framework`, `m1/clerkenstein-backend`, `m1b/clerk-drift-detection`,
`m2/clerkenstein-frontend`. **M1 is iterative** → built by `/developer-kit:build-mstone-iters` (closes
on a Gate Outcome Ledger). M0/M1b/M2 are section → `/developer-kit:build-milestone`. All →
`/developer-kit:close-milestone` → `/developer-kit:close-release`.

### Out of scope (v1.0 — recorded for v1.1+)
- Multi-instance disposable stacks, data seeding, use-case recipes → all v1.1 "show floor".
- Mirroring engines other than Clerk with M0 (the framework is generic, but v1.0 only exercises it on Clerk).
- AI-generated demo content (transcripts/embeddings) → v1.1 stretch or deferred.

## Done

_(none yet — no version has shipped under this lifecycle.)_

## Notes

- Milestone numbering is **flat sequential** (M0, M1, M2, …); B-milestones append `b` (M1b). See [`context.md`](context.md).
- v1.0 mixes shapes: M0/M1b/M2 are **section**; **M1 is iterative** (alignment-score gate).
- v1.1 "show floor" (M3, M4, M5) is detailed in [`roadmap-vision.md`](roadmap-vision.md); it promotes into this file when v1.0 closes.

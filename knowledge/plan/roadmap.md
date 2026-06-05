# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills).

> **Designed 2026-06-02** from the Demo Environment + Clerkenstein brief, **refined 2026-06-02** to
> promote alignment measurement into a first-class discipline (new **M0**). 3 research agents over the
> Clerk integration, the staging/dev-env tooling, and the data/seeding surface — all verified against
> the cloned platform in `anthropos-dev/`. Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-02.md`](../../.agentspace/scratch/roadmap-research-2026-06-02.md).
>
> **v1.0 "body double" — SHIPPED 2026-06-03** (merged to `main`, tagged `v1.0`; full detail in `## Done` below).
> **v1.1 "show floor" — SHIPPED 2026-06-05** (merged to `main`, tagged `v1.1`; full detail in `## Done — v1.1`
> below). 8 milestones M3→M8: the 2-repo consolidation + demo/dev stacks + the production-safe seeding stack
> (framework + data-DNA + fleet) + the corpus product layer. Next action: `/developer-kit:design-roadmap` for v1.2.

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.0** | **body double** | A *measured* stand-in the platform can't tell from the real thing | M0 → M1 → { M1b ∥ M2 } → M2b → M2c | ✅ **SHIPPED 2026-06-03** (tag `v1.0`) |
| **v1.1** | **show floor** | The platform-operations extension framework (demo + dev, in 2 repos) | M3 ✅ → M4 ✅ → M5 ✅ → M6 ✅ → M7a ✅ → M7b ✅ → M7c ✅ → M8 ✅ | 🏁 **all milestones done** — ready for `/developer-kit:close-release` (`release/01.10-show-floor`) |

The whole initiative layers a **second corpus + skill set on top of** the existing dev-environment
tooling, to build disposable demo environments. Hard constraints: **no modification to any platform
repo** (current or future) and **no disruption to the dev environment** — demo clones live under the
gitignored `anthropos-demo/` (mirroring `anthropos-dev/`). Full brief:
[`.agentspace/demo-environment-draft.md`](../../.agentspace/demo-environment-draft.md).

## Done — v1.1 "show floor" (SHIPPED 2026-06-05 · tag `v1.1`)

**Theme (broadened 2026-06-04):** v1.0 made the platform run *without* Clerk; v1.1 started as "disposable
demo stacks" (M3 ✅) and now becomes **the platform-operations extension framework** — consolidate the repo
constellation into **two repos** (`rosetta` = the platform corpus + dev-env skills; `rosetta-extensions` = a
monorepo of operations sections), then deliver the seeded-demo capability *and* generalize the pattern to dev.
Everything stays **additive — zero change to any read-only platform repo**.

**Refactored 2026-06-04** (after M3 shipped, to keep the constellation from exploding): the standalone
`clerkenstein` + `rosetta-demo` repos collapse into `rosetta-extensions/{clerkenstein,demo-stack,…}`; the
former M4 (seeding) → **M7**, former M5 (recipes) → **M8**; new structural milestones M4–M6 inserted. Decisions:
**git subtree, history-preserving** (M4-D1) · **delete the old repos, not archive** (M4-D2, user) · **the
alignment framework stays in rosetta** (M4-D3) · per-demo clones (M3-D1) · clone-at-release-tag (M3-D3).

**Seeding redesigned 2026-06-04** (M3–M6 all shipped): the user asked to make seeding robust/resilient/drift-proof/
fast/**production-safe**, so the single `section` M7 splits into **M7a → M7b → M7c** (a section + section +
iterative "mix"). 3 research agents over the platform grounded it: the prod-pollution boundary is *small + fixed*
(Directus, S3-public, live Clerk/external SaaS — everything in the per-stack Postgres is isolated); the M0 alignment
pattern *extends to data* (new structural operators + schema-as-source); the perf bottleneck is *DB-IO, not CPU*
(Go-link-ent + `COPY` + fan-out; Rust buys nothing). Decisions: **3-way split, all in v1.1** (M7a-D1, user chose
keep-in-v1.1 over a v1.2 spin-out) · **the isolation guard is the load-bearing deliverable** (M7a-D2) · **extend
M0 to a data dimension, don't fork it** (M7b-D1) · **the data-DNA is the catalog that drives the fleet** (M7b-D2)
· **the fleet is iterative, gated on data-DNA coverage** (M7c-D1).

### M3: Disposable multi-instance demo stacks ✅ DONE (2026-06-03; extended close 2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m3-demo-stacks/](releases/archive/01.10-show-floor/m3-demo-stacks/)
Spun up `demo-N` as isolated, Clerkenstein-wired full stacks; the full Clerk-free injected stack + migrate are
LIVE-PROVEN; the deployment/injection alignment surface (`clerk-deploy-1`, 7/7) landed. 78 demo-stack tests, 218
clerkenstein funcs. **Delivered** `corpus/ops/rosetta_demo.md` + `/demo-*` skills.

### M4: Consolidate into the `rosetta-extensions` monorepo ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m4-consolidate-extensions/](releases/archive/01.10-show-floor/m4-consolidate-extensions/)
Created the **`rosetta-extensions`** monorepo (private, 73 commits); `git subtree`-imported `clerkenstein` +
`rosetta-demo`(→`demo-stack`) **with full history preserved**; the `knowledge/` nav; thinned rosetta to pointers;
fixed a +1-depth path break the verify gate caught (M4-D4); verified under the new paths (78 demo-stack tests +
deploy gate 7/7); pushed; **removed the old `clerkenstein` + `rosetta-demo` repos** (local + org, 404). Decisions
M4-D1 (subtree) / D2 (delete-not-archive) / D3 (alignment framework stays in rosetta) / D4 (path-depth fix).

### M5: Extract the reusable `stack-injection` layer ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m5-stack-injection/](releases/archive/01.10-show-floor/m5-stack-injection/)
Extracted the generic injection (`inject.py`, `gen_injected_override.py`, `apply-authn.sh`) into
`rosetta-extensions/stack-injection/`, consumable by any stack with a **demo-ON / dev-OFF** toggle; the mock stayed
in clerkenstein (dependency runs stack-injection→clerkenstein, M5-D1); the port-offset engine stayed in demo-stack
(M5-D2, settles the M4 open question — moves to shared in M6). Split the tests, repointed the consumers; **78
preserved**, flake 3/3, deploy gate 100%/100%.

### M6: `dev-stack` — tooled local dev environment ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Dir:** [m6-dev-stack/](releases/archive/01.10-show-floor/m6-dev-stack/)
Extracted the shared port-offset engine into a new **`stack-core/`** section (settles the M5-routed question —
demo + dev share it, M6-D1) and added a focused **`dev-stack/`**: isolated dev stacks (`dev-N`, offset ports,
guarded `-p dev-N`), **real Clerk by default**, Clerkenstein injection **optional** (reuses stack-injection).
Scoped to the proven value (M6-D2 — not speculative multi-dev). **87 tests** (+9), flake 3/3, deploy gate 100%/100%.

### M7a: Seeding framework + production-isolation safety ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Complexity:** large · **Dir:** [m7a-seeding-framework/](releases/archive/01.10-show-floor/m7a-seeding-framework/)
Built `rosetta-extensions/stack-seeding/` — a host Go module that seeds a stack by talking **directly to its
Postgres** (offset port, `COPY`; *not* ent-linking — `app/internal/bootstrap` is internal, unimportable, M7a-D3)
behind a **3-layer production-isolation guard** (CheckWrite · PreflightEnv · AssertClean). **LIVE-PROVEN**: a
fresh injected `demo-1` → `migrate-demo.sh` (now bootstraps the global Sentinel policy) → `stackseed` (org + 1000
users + the real `user_clerkenstein` identity + the casbin `g2` grant, isolation audit clean) → authenticated
login returns **HTTP 200** (`membershipsCount: 1001`). The proof caught + fixed **2 real bugs** (the g2 arg-order;
the missing global-policy bootstrap — M7a-D4). **68 tests**, all gates green. Delivered `corpus/ops/seeding-spec.md`.

### M7b: The data-alignment dimension ("data DNA") ✅ DONE (2026-06-04)
**Status:** `done` · **Shape:** `section` · **Complexity:** medium · **Dir:** [m7b-data-dna/](releases/archive/01.10-show-floor/m7b-data-dna/)
Extended the **M0 alignment framework** to a **data** dimension — the `datadna` harness (`rosetta-extensions/
stack-seeding/dna/`) that (a) enumerates the seedable surfaces (**4 seeded + 6 planned** — the M7c checklist) and
(b) measures a seeder's output conforms to the platform's **current schema** via **structural operators**
(type-match / constraint-satisfied [NOT-NULL + UNIQUE] / fk-valid / row-count) with **schema-as-source via
introspection**. A separate harness, not an alignctl runner (M7b-D3). **PROVEN live** on the M7a-seeded `demo-1`:
`measure` **100% / Critical 100%** across the 4 seeded surfaces; `diff` flags an injected column (exit 1) and
reads clean on revert. Caught + fixed the planned-surface introspection bug; hardened the UNIQUE leg (M7b-D4).
**dna 49 + cmd/datadna 10 + pg 17 tests.** Delivered the data dimension into `corpus/architecture/alignment_testing.md`.

### M7c: The seeder fleet, to a coverage gate ✅ DONE (2026-06-05, gate-met-over-reachable + waiver)
**Status:** `done` · **Shape:** `iterative` · **Complexity:** large · **Dir:** [m7c-seeder-fleet/](releases/archive/01.10-show-floor/m7c-seeder-fleet/)
Built the fleet across 5 iters (TOK-01 strategy → jobsim-sessions → skillpath-sessions → assignments → activity),
each a deterministic **backdated-activity** seeder (time-distributed, pass/fail per `pass_rate`, content refs as
free values — the believability core is reachable **without** the shared Directus). Drove data-DNA coverage
**40%→80%**, promoting each surface planned→seeded + conformance-gated. **Gate: 3 of 4 met outright** — (a)
login→**200** · (c) full 8-seeder seed **0.69s** (~8500 rows, <2min) · (d) isolation **clean**; (b) coverage is
**100% over the 8 reachable surfaces / critical 100%**, with **taxonomy + content waived** (the hard line —
skiller snapshot + shared Directus; Re-scope trigger, user-confirmed → ~v1.2). Caught + fixed 2 live bugs (the
skillpath UNIQUE constraint; the introspect-load harness bug). **20 seeder / 145 module tests.** Delivered
`rosetta-extensions/stack-seeding/seeders/` + the `waived` data-DNA status.

### M8: Corpus + use-case recipes + polish ✅ DONE (2026-06-05) — LAST v1.1 milestone
**Status:** `done` · **Shape:** `section` · **Complexity:** medium · **Dir:** [m8-corpus-recipes/](releases/archive/01.10-show-floor/m8-corpus-recipes/)
The consolidation/discoverability layer: a **`corpus/ops/demo/` family** (index + 3 end-to-end recipes —
enterprise-onboarding, skill-progression, browser-login [which lands the 2 M3-deferred injection recipes: the
`api.clerk.com` cert-redirect + the browser-login walk-through]); **3 seed presets** (small/mid/large, mid-500 +
large-1k seed-proven end-to-end); the **`/demo-seed` skill** + the CLAUDE.md skill table; the v1.0
**express-gate CI carry-forward** wired into clerkenstein `alignment.yml` (**validated 9/9** locally); and
cross-linking from corpus/README + root README + CLAUDE.md (all doc links resolve). **Next:** `/developer-kit:close-release`.

### Execution graph (v1.1)
```
v1.1 "show floor" — the platform-operations extension framework (demo + dev, in 2 repos)
   M3 ✅ ─→ M4 ✅ (consolidate) ─→ M5 ✅ (stack-injection) ─→ M6 ✅ (dev-stack)
                                            └──→ M7a (framework+safety) ─→ M7b (data-DNA) ─→ M7c (seeder fleet) ─→ M8 (recipes)
```
**Sequential.** M4–M6 shipped (the extension framework + demo/dev stacks). M7a lands the framework + the
isolation guard (a usable, safe demo); M7b builds the data-DNA catalog that lists + gates the seeders; M7c drives
the fleet to the coverage gate; M8 curates the output.

### Risks (v1.1)
- **(M7a, blocks-prod-safety)** a single un-guarded **shared-write reaching prod** (Directus / S3-public bucket) —
  mitigate with the hard isolation guard + the clean-audit assertion as a tested acceptance gate, not a convention.
- **(M7a, scope)** linking the platform's `app/internal/bootstrap`/ent client into a `rosetta-extensions/` Go
  module without a platform edit — confirm the import path early (fallback: `go run` CLIs, slower).
- **(M7b)** trustworthy schema-as-source — get ent introspection / `atlas inspect` golden right or the drift diff lies.
- **(M7c, scope)** the heaviest build: ~8–10 seeders + 1k-scale `COPY` perf + backdating fidelity, each gated on
  conformance — the believable-demo *subset* of surfaces is the real target (waive unreachable genes, don't chase 100%).

### Open decisions (resolve during build)
Directus snapshot-replay vs hard-block-and-skip for the demo MVP (M7a); ent-introspection vs `atlas inspect`
golden for schema-as-source (M7b); whether seed presets ship in M7c or M8; external shareability (Tailscale vs
ingress); the AI-content STRETCH trigger (now firmly v1.2, not M8).

## Done — v1.0 "body double" (SHIPPED 2026-06-03 · tag `v1.0`)

> **Shipped 2026-06-03.** All six milestones closed-on-gate / completeness-complete and merged to `main`;
> `release/01.00-body-double` deleted. Release records archived under
> [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/) (review · retro · metrics ·
> lockfile · stats). Headline: a *measured* drop-in Clerk mock at **100%/100% on all three surfaces**
> (Go · JS/FAPI · `@clerk/express`), built by a first-class alignment framework, zero platform-code change.
> Close-release caught + fixed 1 blocker (an `@clerk/express` gate regression from the M2c close) — see the
> release retro.

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

**Closed 2026-06-02** (build S1–S5 → harden 2 passes → close review → merged to `release/01.00-body-double`). Delivered: `test/alignment/` — `alignctl` (stdlib-only Go, builds/runs offline) with `run`/`capture`/`dna list|diff|validate`, engine-agnostic via a pluggable `--runner`; the 4 equivalence operators + weighted score (overall + separate critical gate); record/replay goldens; `internal/canon` precision-safe canonicalization; the `//go:build alignment` test class; and a toy reference proving end-to-end detection (**86.7% / 100% critical**, catches `Greet/padded-name`). Plus `/align-dna` + `/align-run` skills and `corpus/architecture/alignment_testing.md`. Open questions resolved: DNA format = **JSON** (M0-D1); capabilities enumerated from the consumed surface; goldens live per-mirror-repo. Close-review adversarial pass found + fixed a path-traversal must-fix + score-overflow (M0-D7); 45 test funcs (3 fuzz), 5/5 flake gate, core coverage 83–98%. Decisions M0-D1…D7. Retro: [m0-alignment-framework/retro.md](releases/archive/01.00-body-double/m0-alignment-framework/retro.md). Resolved repo split: framework in rosetta; the Clerk DNA/tests/mirror land in the `clerkenstein` repo (M1).
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

**Closed 2026-06-03** (5 iters: bootstrap tok TOK-01 → DNA → authn twin → critical orgclient → standard orgclient → **gate** → final harden → close). The **Clerkenstein backend mirror** (in the gitignored `anthropos-demo/clerkenstein` repo, its own git) scores **100% alignment / 100% critical** against the `clerk@2.6.0` DNA (22 genes), built offline. authn implements the real `colony/authn.Provider` (HS256, one universal key); orgclient is a disarmed in-memory twin. Score arc: 0 → 21.1 → 68.4 → **100%**. Final harden: authn + orgclient **0 → 100%** unit coverage (+1 fuzz, 0 bugs). Decisions: D1 hybrid goldens; iter-01-D1 authn injects via `go.mod replace` whole-colony; **M1-D2 orgclient injects via a fake-Clerk-API-server → routed to M2** (shared HTTP-interception with the JS side). Delivered `corpus/services/clerkenstein.md`. Retro: [m1-clerkenstein-backend/retro.md](releases/archive/01.00-body-double/m1-clerkenstein-backend/retro.md). The gate (alignment fidelity) is met; live injection into a running platform is rosetta-demo work (v1.1) / M2 (orgclient).
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
**Status:** `done` (2026-06-03)
**Closes the gap after:** M1 (Clerkenstein is aligned at v2.6.0 — but must *stay* aligned as the platform bumps `clerk-sdk-go` / `@clerk/*`).

**Closed 2026-06-03** (2 sections + 1 harden pass). Automation/config over M0 — no new measurement machinery. In the clerkenstein repo: `scripts/gate.sh` (alignment gate, built-binary so exit 0 met / 2 regressed) + `scripts/drift-check.sh` (DNA-diff + gate, exit-code contract **0** none / **1** DNA moved / **2** gate regressed / **3** usage) + `.github/workflows/alignment.yml` (push + **weekly** CI) + `scripts/drift-test.sh` (9-assertion regression harness pinning the contract + the 2 build-phase fixes). Delivered the "Drift detection (M1b)" runbook in `corpus/services/clerkenstein.md`. Verified across all exit paths against a simulated `clerk@2.7.0` bump; shellcheck clean, flake 5/5. Close review: 0 findings.
**Goal:** Reuse M0 wholesale to make Clerk drift a flagged, mechanical event: on a version bump, `/align-dna` diffs the DNA (what changed) and `/align-run` re-scores the existing mirror against the new source (score drop = broken genes), CI-gated on "alignment ≥ threshold."
**Scope:**
  - In: the "bump pinned Clerk version → DNA-diff → re-score → report" workflow; the CI gate on alignment score; golden-capture refresh on bump.
  - Out: building the framework (M0); authoring the original mirror/DNA (M1); the JS surface (M2 owns its own genes).
**Depends on:** M1 (needs a built, aligned mirror + the Clerk DNA). Reuses M0's skills — **now automation/config over M0, not new machinery** (the right size for a B-milestone).
**Parallel with:** M2 (CI/automation vs JS code — disjoint surfaces).
**Acceleration effect:** every future Clerk bump becomes a flagged, scored update instead of a silent break — the brief's "follow platform updates within minutes" requirement, mechanized.

### M2: Clerkenstein — browser session + webhook coherence (JS)
**Status:** `done` (2026-06-03)
**Shape:** `section`
**Goal:** The frontend logs in with no real Clerk, and created/seeded users/orgs reach the DB without real Clerk webhooks.

**Closed 2026-06-03** (5 sections S1–S5 → 4 harden passes → close review → merged to `release/01.00-body-double`). Closes the last two Clerk seams so a demo stack is **Clerk-free end to end**. Delivered (in the gitignored `anthropos-demo/clerkenstein` repo): the **fake FAPI server** (`fapi/`) + the publishable-key codec — the browser logs in via a *minted publishable key* that encodes the fake FAPI host, **config-only, no SDK fork** (M2-D1 spike resolved the milestone's defining risk in the strong direction; the real-dev-Clerk fallback is documented but un-exercised); the **fake BAPI server** (`bapi/`) that disarms the platform's networked `orgclient` via an `api.clerk.com` DNS/base-URL redirect (the **M1-D2 Fate-3 pickup**), backed by the M1 orgclient twin made **concurrency-safe** (M2-D2); the **svix-signed webhook injector** (`webhook/`) for the 12 consumed event types → `POST /api/webhook/clerk`; and a **second Alignment DNA** (`clerk-js-5`, 9 genes, runner `cmd/jsfapirun`) scored at **100%/100%** like the Go side — proving the M0 framework is **surface-generic**. Both gates 100%/100% (Go 22/22 + JS 9/9); 112 Go test/fuzz funcs; flake 5/5; gofmt/vet/shellcheck clean. **Close review** found + fixed an `orgclient.ChangeRole` nil-map panic + phantom-membership divergence the alignment gate missed (reachable via the `bapi/` server) — M2-D4, with regression tests; plus a gofmt fix + the repo README refresh; 0 scope gaps, 0 deferrals (deferral audit GREEN). Decisions M2-D1…D4. Retro: [m2-browser-webhook-coherence/retro.md](releases/archive/01.00-body-double/m2-browser-webhook-coherence/retro.md). **This was the last *feature* milestone of v1.0**; a cleanup B-milestone **M2b (repo consolidation)** was inserted after it (2026-06-03) to tidy the `clerkenstein` repo before `/developer-kit:close-release`.
**Scope:**
  - In: a fake Clerk FAPI path for `@clerk/nextjs ^6.39.2` (next-web-app, ant-academy) and `@clerk/clerk-js ^5.52.3` (studio-desk) via publishable-key + base-URL/DNS override — **with the decided fallback**: keep the real dev Clerk app for the browser session while the backend stays fully mocked; a **webhook injector** feeding the existing `app/internal/clerk/events/` sync pipeline directly; **the JS surface's fidelity expressed as alignment genes via M0** where applicable (same score treatment as the Go side).
  - In (**routed from M1 close — M1-D2, Fate 3**): the **fake-Clerk-API-server** (HTTP interception of `api.clerk.com`) ALSO serves M1's **orgclient** injection — the Go `app/internal/clerk/orgclient` is app-internal + networked, so it can't `go.mod replace` like authn; it disarms via the same fake-API-server this milestone builds for the JS side. The Clerkenstein orgclient mirror behavior already exists + scores 100% (M1); M2 wires the HTTP redirect that makes the platform's real orgclient hit it.
  - Out: multi-instance stacks (M3); data seeding (M4).
**Depends on:** M1 (consumes the mock contract + minted-token shape). **Parallel with:** M1b (yes).
**Estimated complexity:** large — **highest technical risk in v1.0** (SDKs hard-code Clerk FAPI; no documented base-URL override).
**Open questions:** can `@clerk/*` be pointed at a fake FAPI without a fork? (the fallback exists because this is uncertain) — spike the override early.
**KB dependencies:** `corpus/architecture/alignment_testing.md`, `corpus/services/clerk-integration.md`, `corpus/architecture/frontend_architecture.md`, `corpus/services/next-web-app.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`:** extends the M1 doc with the JS path + webhook injection + the fallback decision.

### M2b: Clerkenstein repo consolidation + knowledge base
**Status:** `done` (completed 2026-06-03)
**Shape:** `section`
**Dir:** [m2b-clerkenstein-consolidation/](releases/archive/01.00-body-double/m2b-clerkenstein-consolidation/)

**Closed 2026-06-03** (5 sections S1–S5 → 1 harden pass → close review → merged to `release/01.00-body-double`). A pure-cleanup B-milestone that reorganized the `clerkenstein` repo (gitignored `anthropos-demo/clerkenstein`, its own git on `main`) into a clean, self-documented **library-named** structure — **no behavior change**, both alignment gates (Go 22/22, JS 9/9) + the drift harness (9/9) stayed green throughout. Delivered: the **library-named dirs** (`authn/` mocks colony/authn · `clerk-backend/` mocks clerk-sdk-go/v2 = the bapi server + orgclient store **merged** · `clerk-frontend/` mocks @clerk/clerk-js+nextjs · `clerk-webhook/` mocks svix) + `shared/` (the universal-key HS256 JWT, extracted because `clerk-frontend` **mints** and `authn` **verifies** the same token — `parse`→`shared.Parse` exported, M2b-D4) + `alignment/` (the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`) via **69 history-preserving `git-mv` renames**; a self-contained **`knowledge/` base** (kb-index + scope + architecture + alignment + injection + coverage-index) + 6 per-library READMEs + slim root README; an `.agentspace/` (gitignored contents, dir preserved) + `.gitignore` baseline + asset hygiene; and `CLAUDE.md` + `singularity-manifest.md` (authored TO the `/singularity-kit:repo-consolidate` standard — the formal `repo-consolidate code` run is a **USER finalize**, M2b-D3/D8, since the skill is `disable-model-invocation`). Rosetta-side: slimmed `corpus/services/clerkenstein.md` 197→62 lines to a pointer at the repo's KB + fixed 2 stale refs in `alignment_testing.md`. **Close review** found + fixed 1 should-fix code-quality (a fuzz-test comment naming pre-reorg packages) + 2 doc findings (coverage-index count drift 112→113 after the harden test, state.md Headline refresh) — clerkenstein fixes on its own `main` (`ad87545`); 0 scope gaps, 0 deferrals (deferral audit **GREEN** — 2 inherited singles owned by close-release/M3, 0 repeat). Decisions M2b-D1…D8; D1/D2/D4 blended into the repo's own KB. Retro: [m2b-clerkenstein-consolidation/retro.md](releases/archive/01.00-body-double/m2b-clerkenstein-consolidation/retro.md). **This was the LAST milestone of v1.0** → next is `/developer-kit:close-release`.

**Goal:** The `clerkenstein` repo grew organically across M1/M1b/M2 into flat package dirs (`authn bapi orgclient fapi webhook cmd dna golden golden-js scripts`) with a single README and no knowledge base. M2b reorganizes it into a clean, self-documented **library-named** structure — one dir per mocked dependency + a shared dir + an alignment harness dir + a `knowledge/` base — following `/singularity-kit:repo-consolidate`, so the repo is navigable + operable by agents *before* v1.0 ships.
**Context (B-milestone — cleanup after M2):** pure reorg / docs / hygiene over the M2-complete repo. **No behavior change** — both alignment gates (Go 22/22, JS 9/9) and the drift harness stay green throughout; the move repoints imports + DNAs/goldens/runners/scripts, it does not alter the mocks. Class of work like M1b (tooling/cleanup over a shipped surface).
**Scope:**
  - In (**1 — Restructure**): one dir per mocked library/framework + one shared dir, **library-named** (user-chosen scheme): `authn/` (mocks `colony/authn`), `clerk-backend/` (mocks `clerk-sdk-go/v2` — the `bapi` server + the `orgclient` store **merged into one dir**), `clerk-frontend/` (mocks `@clerk/clerk-js` + `@clerk/nextjs` — the FAPI), `clerk-webhook/` (mocks `svix`); `shared/` (universal-key HS256 JWT + claims + canonical helpers — extracted because `clerk-frontend` mints and `authn` verifies with the same key); `alignment/` (the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`). **Tests stay co-located within each library dir.** Go package identifiers can't contain hyphens → each hyphenated dir declares a clean package (e.g. `clerk-backend/` → `package clerkbackend`) — M2b-D1, confirmed at build.
  - In (**2 — Knowledge base**): a self-contained `knowledge/` dir documenting Clerkenstein — scope/goal; how it's built (the 4 mocks + shared); how fidelity is **validated with alignment tests against a pinned Clerk version** (the M0 framework + the two DNAs + the gate); **per-library injection recipes** (`go.mod replace` for `authn`; `api.clerk.com` HTTP/DNS redirect for `clerk-backend`; config-only publishable-key override for `clerk-frontend`; direct svix-signed POST for `clerk-webhook`); a coverage index. Per-library `README.md`s + a top-level index. Solid, well-written, well-distributed.
  - In (**3 — Hygiene**): an `.agentspace/` dir with contents **gitignored**; `.gitignore` cleanup (the current comment is mismatched); built-binary + transient hygiene per `repo-consolidate`'s asset-hygiene checks.
  - In (**4 — Consolidate**): run `/singularity-kit:repo-consolidate code` to standardize the repo (emit `CLAUDE.md` + `singularity-manifest.md`, audit against the code-repo + asset-hygiene standards, apply fixes), then re-verify both gates + the drift harness. **Note:** `repo-consolidate` is `disable-model-invocation` (user-invoked) — the build authors the structure TO its standard so the run is a clean finalize; the **user types the skill** (pointed at the `clerkenstein` repo).
  - Out: new library support / new alignment genes (the `@clerk/express` coverage gap — **now picked up by M2c**); any live injection wiring into a running platform (still v1.1/M3); any change to rosetta's M0 framework or to the platform repos.
**Depends on:** M2 (consolidates the M2-complete repo). **Parallel with:** none (touches the whole repo). **Precedes:** `/developer-kit:close-release`.
**Estimated complexity:** medium — mechanical but wide (touches every package + the gate/drift scripts); the only real risk is import/script repointing, fully caught by the **green-gate invariant** (gates + drift re-run after each section).
**KB dependencies:** `corpus/services/clerkenstein.md`, `corpus/architecture/alignment_testing.md`; the `/singularity-kit:repo-consolidate` standards (base + code-repo + asset-hygiene).
**Delivers → the `clerkenstein` repo's own `knowledge/` base** (net-new, self-contained) **+ slims `corpus/services/clerkenstein.md`** (rosetta) to a pointer at the repo's `knowledge/` + the new structure.

### M2c: Clerkenstein — `@clerk/express` backend session verification (RS256/JWKS)
**Status:** `done` (2026-06-03, **closed-on-gate**)
**Shape:** `iterative` (alignment-score gate, like M1) — a **feature** milestone; the letter suffix marks *insertion after M2b*, not a B/tooling milestone.

**Closed 2026-06-03** (5 iters: bootstrap TOK-01 → DNA → RS256 foundation → **crux proof** → full runner → gate; 1 final harden pass). Brought the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend) under the alignment framework at **100%/100%** (3rd DNA `clerk-express-1.json`, 9 genes). The **RS256 wall fell to an additive path** (M2c-D1/D2): an RSA keypair + a real JWKS + RS256 minting that the *genuine* `@clerk/backend` accepts networkless via `jwtKey` — **no HS256 migration**, so M1 (22/22) + M2 (9/9) stayed green. `@clerk/express` is **verified, not reimplemented** (no mock dir — the svix discipline; M2c-D5); the `expressrun` runner mints tokens (Go) + drives the real SDK (embedded `verify.js`, Node). The `clerkClient` BAPI reads were already covered by `clerk-backend` (M2c-D4). Close: folded the surface into the knowledge base + corpus, fixed a gitignore gap + 1 adversarial flake (`tamperSig`); deferral audit GREEN; the express-gate CI-wiring (needs Node) routed to v1.1. 128 test/fuzz funcs / 8 packages; all four gates green. Retro: [m2c-clerk-express-alignment/retro.md](releases/archive/01.00-body-double/m2c-clerk-express-alignment/retro.md).
**Dir:** [m2c-clerk-express-alignment/](releases/archive/01.00-body-double/m2c-clerk-express-alignment/)
**Goal:** Bring the **last un-gated Clerk consumer — `@clerk/express`** (studio-desk's Node backend auth) under the alignment framework: a new **`clerk-express/`** seam + a **3rd Alignment DNA**, driven to a gate, so studio-desk's backend genuinely verifies Clerkenstein tokens (not via its `MOCK_CLERK=true` bypass). Completes v1.0's thesis — *no* Clerk seam left un-faithful before shipping.
**Why iterative + the defining unknown (the RS256 wall):** `@clerk/express` (via `@clerk/backend`) verifies **RS256 via JWKS only** and **hard-rejects HS256** (`assertHeaderAlgorithm` → `TokenInvalidAlgorithm`). Clerkenstein mints HS256 universal-key tokens + serves an **empty JWKS**, so an HS256 shim is a dead end. The milestone must add an **RS256 path** (RSA keypair + a real JWKS from the fake FAPI + RS256 minting + the real-`@clerk/express` verifier). **The central iteration question:** can RS256 be **additive/parallel**, or must the existing HS256 seams (`authn`/`clerk-frontend`/`shared`) **migrate to RS256** — re-capturing the Go DNA goldens + re-gating M1/M2? The gate-driven iterations resolve it.
**Scope:**
  - In: a new **`clerk-express/`** seam (library-named); an **RSA keypair + a real (non-empty) JWKS** served by the fake FAPI (`clerk-frontend`'s `/.well-known/jwks.json`); **RS256 token minting**; the `@clerk/express` **DNA** (`clerk-express-1.json`, source `@clerk/express ^1.3.47`); a runner that drives **the real `@clerk/express` SDK** against the mock (the svix-pattern — verify against the genuine library); the **alignment gate** as the exit criterion.
  - In (confirm, don't rebuild): `@clerk/express` also calls `clerkClient.{getOrganizationMembershipList, getOrganization}` — those are **BAPI**, already 100%-mocked by `clerk-backend/`; M2c adds *integration* genes confirming that path, not a new BAPI mock.
  - Out: changing studio-desk or any platform repo (the `MOCK_CLERK` bypass is the platform's own); a webhook (svix) DNA (separate future gap); live injection into a running studio-desk (rosetta-demo work, v1.1).
**Candidate genes (~8, `clerk-express-1.json`):** `ExpressAuth/{valid, expired, malformed, bad-signature, no-token}` (error_class) · `ExtractIdentity/universal-user` (exact: verified claims → `req.auth`) · `JWKS/non-empty-rsa` (shape) · `ClerkClientBAPI/{org-membership-list, get-organization}` (integration vs `clerk-backend`).
**Exit gate:** alignment **≥ 95% overall / 100% critical** on `clerk-express-1.json`, AND the load-bearing test passes (a **real `@clerk/express` instance accepts a Clerkenstein-minted token + extracts the right identity**).
**Depends on:** M2 (the FAPI + token machinery it extends) + M2b (the consolidated repo it adds a seam to). **Precedes:** `/developer-kit:close-release`.
**Estimated complexity:** large — **highest fidelity-risk in v1.0**: the RS256 path may force a token-algorithm migration of the existing 100%-gated seams.
**KB dependencies:** `corpus/architecture/alignment_testing.md`; the clerkenstein repo's own `knowledge/` (alignment / architecture / injection / sources); the `@clerk/express` + `@clerk/backend` source under `anthropos-dev/studio-desk/node_modules`.
**Delivered → the clerkenstein repo's `knowledge/`** (alignment/architecture/sources updates) **+ a 3rd DNA + the `expressrun` runner;** updated `corpus/services/clerkenstein.md`'s scorecard to a **3rd *measured surface*** (`@clerk/express`, **verified-not-mocked** — no new mock dir, per M2c-D5; the genuine SDK is *satisfied* via an additive RS256/JWKS path).

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
              │  (both closed — repo feature-complete)
              ↓
    M2b (repo consolidation — library-named dirs + self-contained knowledge base; gates stay green)
              │
              ↓
    M2c (ITERATIVE: @clerk/express RS256/JWKS — new clerk-express/ seam + 3rd DNA → alignment gate)
              │
              ↓
    /developer-kit:close-release → v1.0 ships to main
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

`release/01.00-body-double` (cut from `feat/demo-environment` at M0). Milestone branches:
`m0/alignment-framework`, `m1/clerkenstein-backend`, `m1b/clerk-drift-detection`,
`m2/browser-webhook-coherence`, `m2b/clerkenstein-consolidation`, `m2c/clerk-express-alignment`.
**M1 + M2c are iterative** → built by `/developer-kit:build-mstone-iters` (close on a Gate Outcome Ledger).
M0/M1b/M2/M2b are section → `/developer-kit:build-milestone`. All → `/developer-kit:close-milestone` →
`/developer-kit:close-release`.
The `clerkenstein` repo's own code commits stack on its `main` (its own gitignored git, no branch model);
the rosetta-side milestone records + corpus pointer land on the `m{N}/…` branch.

### Out of scope (v1.0 — recorded for v1.1+)
- Multi-instance disposable stacks, data seeding, use-case recipes → all v1.1 "show floor".
- Mirroring engines other than Clerk with M0 (the framework is generic, but v1.0 only exercises it on Clerk).
- AI-generated demo content (transcripts/embeddings) → v1.1 stretch or deferred.

## Shipped releases

- **v1.0 "body double"** — shipped **2026-06-03**, tag `v1.0`. The alignment-testing framework + Clerkenstein
  (100%/100% on Go · JS/FAPI · `@clerk/express`). Detail in the `## Done` section above; records archived at
  [`releases/archive/01.00-body-double/`](releases/archive/01.00-body-double/).

## Notes

- Milestone numbering is **flat sequential** (M0, M1, M2, …); a letter suffix has two uses: (1) a milestone **inserted after** the fact — `b` for tooling/cleanup (M1b drift CI, M2b consolidation), and the letter-suffixed *feature* milestone M2c (iterative, "inserted after M2b"); and (2) a **split** of one planned milestone into a sequential mini-arc — **M7a → M7b → M7c** is the single former M7 "seeding" split into framework+safety / data-DNA / fleet (2026-06-04, M7a-D1). Both reuse the letter suffix; context disambiguates. See [`context.md`](context.md).
- v1.0 mixes shapes: M0/M1b/M2/M2b are **section**; **M1 + M2c are iterative** (alignment-score gates).
- v1.1 "show floor" mixes shapes too: M3–M6 + M7a/M7b + M8 are **section**; **M7c is iterative** (data-DNA coverage gate).

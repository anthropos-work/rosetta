# Roadmap

Active development plan for **Project Rosetta** (the Anthropos documentation corpus + environment-
builder skills).

> **Designed 2026-06-02** from the Demo Environment + Clerkenstein brief (3 research agents over the
> Clerk integration, the staging/dev-env tooling, and the data/seeding surface — all verified against
> the cloned platform in `anthropos-dev/`). Gap analysis:
> [`.agentspace/scratch/roadmap-research-2026-06-02.md`](../../.agentspace/scratch/roadmap-research-2026-06-02.md).
>
> **Active version: v1.0 "body double"** — *designed, not yet branched.* The `release/01.00-body-double`
> branch and milestone dirs are **not** created yet (Phase 8 deferred by user choice). Next action:
> scaffold v1.0 + cut the branch, or run `/developer-kit:build-milestone` (which can cut it on demand).

## Version plan

| Version | Codename | Theme | Milestones | Status |
|---------|----------|-------|------------|--------|
| **v1.0** | **body double** | A Clerk stand-in the platform can't tell from the real thing | M1 → { M1b ∥ M2 } | **active (designed)** |
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
repo keeps "thinking" it uses Clerk with zero source changes** (the same mechanism staging already
uses to vendor a patched `colony`). Ships as a standalone, parity-tested, drift-gated win before the
demo machinery is built — and removes Clerk's API rate limit as the blocker for scale data-seeding in
v1.1.

**Decided at design (2026-06-02):** two-version split (Clerkenstein first); M2 frontend = attempt the
fake Clerk FAPI server, **fall back to the real dev Clerk app for the browser session** (backend stays
fully mocked) if base-URL override proves too fragile.

### M1: Clerkenstein — backend auth bypass (Go)
**Status:** `planned`
**Shape:** `section`
**Goal:** A drop-in Go mock of `colony/authn`'s provider + the Clerk `orgclient`, injected into the demo build via `go.mod replace` (zero platform-repo edits), so backend services authenticate with one universal credential and locally-minted JWTs — no real Clerk.
**Scope:**
  - In: the dedicated `clerkenstein` repo with a gitignored `.agentspace/` that clones the **pinned** official Clerk SDK (`clerk-sdk-go/v2 v2.6.0`) for interface extraction; a Go authn-provider twin (local JWT mint/verify with the platform claim shape — `eid`, `email`, `firstname`, `lastname`, `org.eid`, `org_id`, `org_role`); an `orgclient` twin covering the methods `app/internal/clerk/orgclient/clerk.go` calls (Invite/Create/Delete membership, Create org, BulkInvite, RevokeInvitation, ChangeRole, Update{User,Membership}Metadata, UpdateClerkOrganizationWithExternalId); per-method **parity tests** asserting same-shape responses vs the real interface; the `replace` + skip-worktree injection recipe.
  - Out: JS/browser auth and webhook sync (M2); the parity/drift CI harness (M1b); multi-instance stacks (M3); data seeding (M4).
**Depends on:** none (first milestone).
**Parallel with:** none (gates M1b and M2).
**Estimated complexity:** large
**Open questions:** stub just `authn` + `orgclient`, or `replace` all of `colony`? (`authn` is a package inside the `colony` module, not its own go.mod). **Resolve in the milestone's first section** — fallback is vendoring the whole `colony`, which staging already does.
**KB dependencies:** `corpus/services/clerk-integration.md`, `corpus/architecture/shared_libraries.md` (§ authn/colony), `corpus/ops/staging-clerk.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`:** the Clerkenstein design + injection mechanism (this is currently a documentation blind area — net-new doc).

### M1b: Clerk parity & drift harness
**Closes the gap after:** M1 (Clerkenstein mirrors Clerk's interface — but must *stay* mirrored as the platform bumps `@clerk/*` / `clerk-sdk-go`).
**Goal:** Tooling that clones the pinned official Clerk into `.agentspace/`, extracts the surface the platform actually calls, and diffs Clerkenstein's responses shape-for-shape — CI-gating drift. This mechanizes the brief's "follow platform updates within minutes" requirement.
**Scope:**
  - In: an interface-extraction step (parse the official SDK surface), an automated parity-diff runner over both libraries, a CI gate that fails on drift, and a "bump the pinned Clerk version → see what broke" workflow.
  - Out: fixing drift (that's a follow-on edit to Clerkenstein); the JS surface (M2 owns its own parity).
**Depends on:** M1 (the milestone it tools).
**Parallel with:** M2 (harness/CI vs JS code — disjoint surfaces).
**Acceleration effect:** de-risks M2 and makes every future Clerk version bump a flagged, mechanical update instead of a silent break.

### M2: Clerkenstein — browser session + webhook coherence (JS)
**Status:** `planned`
**Shape:** `section`
**Goal:** The frontend logs in with no real Clerk, and created/seeded users/orgs reach the DB without real Clerk webhooks.
**Scope:**
  - In: a fake Clerk FAPI path for `@clerk/nextjs ^6.39.2` (next-web-app, ant-academy) and `@clerk/clerk-js ^5.52.3` (studio-desk) via publishable-key + base-URL/DNS override — **with the decided fallback**: keep the real dev Clerk app for the browser session while the backend stays fully mocked; a **webhook injector** that feeds the existing `app/internal/clerk/events/` sync pipeline directly (so seeded users/orgs land in the DB + Sentinel without real svix webhooks).
  - Out: multi-instance stacks (M3); data seeding (M4).
**Depends on:** M1 (consumes the mock contract + locally-minted token shape).
**Parallel with:** M1b (yes); M3 (yes-with-caveats — see parallelism).
**Estimated complexity:** large — **highest technical risk in v1.0** (the SDKs hard-code Clerk FAPI with no documented base-URL override).
**Open questions:** can `@clerk/*` be pointed at a fake FAPI at all without a fork? (the fallback exists precisely because this is uncertain). Spike the override **early** in the milestone.
**KB dependencies:** `corpus/services/clerk-integration.md`, `corpus/architecture/frontend_architecture.md`, `corpus/services/next-web-app.md`, `corpus/ops/webhook_setup.md`.
**Delivers → `corpus/services/clerkenstein.md`:** extends the M1 doc with the JS path + webhook injection + the fallback decision.

### Execution graph

```
v1.0 "body double"

  M1 (Clerkenstein backend Go: authn + orgclient twins, replace-injected, parity-tested)
    │
    ├──→ M1b (Clerk parity & drift harness)        ◀── tools M1; CI-gates drift
    │
    └──→ M2 (browser session + webhook coherence)  ◀── fake FAPI, real-Clerk fallback
              (M1b ∥ M2)
```

### Parallelism

- **M1 → {M1b, M2}** sequential: both consume M1's mock contract / interface.
- **M1b ∥ M2:** disjoint surfaces — M1b is harness + CI; M2 is JS + the webhook injector. Merge risk **low** (only the `clerkenstein` repo's test/CI dirs vs its JS adapter dir).
- **M3 ∥ M2 (yes-with-caveats, cross-version):** M3 (v1.1) can build + smoke-test a stack on backend-only bypass while M2 finishes, but the *browser login* story isn't complete until M2 ships. Sequenced cleanly here by the version boundary (M3 starts after v1.0 closes).

### Risks (v1.0)

| Risk / decision | Severity | Mitigation |
|---|---|---|
| **JS/FAPI fake server** — SDKs hard-code Clerk FAPI, no base-URL override | blocks-release (full no-Clerk browser) | **Decided fallback:** real dev Clerk app for the browser, backend fully mocked; spike override early in M2 |
| **`colony` replace granularity** — `authn` is a package inside `colony`, not its own module | degrades-quality (M1 effort) | M1 first-section spike; fallback = vendor whole `colony` (staging precedent) |
| **Clerk SDK version drift** | degrades-quality | M1b drift harness |
| **Repo layout** — where Clerkenstein lives | blocks-kickoff | **Decided:** `clerkenstein` = own repo, cloned into gitignored `anthropos-demo/`; rosetta holds docs + skills |
| **"Zero platform-code changes" interpretation** — `replace` edits the *clone's* go.mod | nice-to-resolve | build-time injection in the gitignored clone + skip-worktree; upstream repo never modified (same as staging's `vendor-colony/`) |

### Branch model

`release/01.00-body-double` from `main` (to be created at scaffold time — Phase 8, deferred). Milestone
branches: `m1/clerkenstein-backend`, `m1b/clerk-parity-harness`, `m2/clerkenstein-frontend`. Standard
`/developer-kit:build-milestone` → `/developer-kit:close-milestone` → `/developer-kit:close-release`.

### Out of scope (v1.0 — recorded for v1.1+)
- Multi-instance disposable stacks, data seeding, use-case recipes → all v1.1 "show floor".
- AI-generated demo content (transcripts/embeddings) → v1.1 stretch or deferred.

## Done

_(none yet — no version has shipped under this lifecycle.)_

## Notes

- Milestone numbering is **flat sequential** (M1, M2, …); B-milestones append `b` (M1b). See [`context.md`](context.md).
- v1.1 "show floor" (M3, M4, M5) is detailed in [`roadmap-vision.md`](roadmap-vision.md); it promotes into this file when v1.0 closes.

---
milestone: M20
slug: lifecycle-convergence
version: v1.3b "dress rehearsal"
milestone_shape: section
status: planned
created: 2026-06-08
last_updated: 2026-06-08
complexity: medium-large
delivers: corpus/ops/snapshot-cold-start.md (net-new) + updated demo recipes/skills for auto set-dress + set-dress chaining + cold-start capture in the tooling
issues: ISSUE-10, ISSUE-13
---

# M20 — Lifecycle convergence: demo-up auto set-dress + cold-start capture

## Goal
Close the dev↔demo asymmetry — `/demo-up` **auto set-dresses** (snapshot → seed, default-on + non-fatal) like
`dev-up` already does — and **unblock the real catalog on a fresh box** that has no safe `--dsn`.

## Why section
The asymmetry is confirmed (ISSUE-10: `dev-up` runs the M13 set-dress pass; `demo-up` doesn't chain) and the M13
pass is proven + reusable. The cold-start question (ISSUE-13) has a known sanctioned answer (DSN-export / restore a
`pg_dump`) plus a bounded spike (MCP-paging adapter). Concrete scope; the closing milestone.

## Repo split
- **`rosetta-extensions`** (authoring → tag `dress-rehearsal-m20` → consume): the set-dress chaining in
  `up-injected.sh` + the cold-start capture path.
- **`rosetta`**: the net-new `corpus/ops/snapshot-cold-start.md` + the `demo-up`/`demo-down` skill + recipe updates.

## Scope
- **In:**
  - **`rosetta-extensions`:**
    - **Chain** `stacksnap replay` → `stackseed` into `up-injected.sh` after migrate, **reusing M13's proven
      `dev-setdress` pass** — **default-on + non-fatal**, `--no-setdress` escape. (ISSUE-10.)
    - Enforce the **atomicity contract** — a partial snapshot with no seed = 403s, so it's both-or-neither; the M17
      re-run guards make a retry safe.
    - The **cold-start capture path** (ISSUE-13) — a documented, prod-safe **DSN-export / restore-a-`pg_dump`-then-
      `--dsn`** workflow (the sanctioned route, behind M9a's capture-source policy + `AssertPublicOnly`), **plus a
      spike** on whether a thin MCP-paging capture adapter (read via the wired `postgres` MCP) is cheap enough to
      build vs document. Note: directus replay needs a per-stack Directus on an offset port (bootstrap→replay→boot).
  - **`rosetta`:** update the `demo-up`/`demo-down` skills + the `corpus/ops/demo/` recipe family for auto set-dress.
- **Out:** **S3 media blob bytes + the cloud `SnapshotStore` backend** (DEF-M10-01 → **v1.4**, signed — untouched);
  AI-generated content (v1.4).

## Depends on
**M18** (the post-set-dress verify) **+ M19** (the full experience) **+ M17** (re-run-safe primitives make
auto-chaining safe to retry). **Parallel with:** none (the closing milestone).

## Open questions (resolve during build)
- ISSUE-13 — build the MCP-paging capture adapter vs document the DSN-export step only — lean: **document the
  sanctioned DSN-export/restore path now**; spike the MCP adapter only if cheap (the MCP is a query tool, not a
  `--dsn` `stacksnap` can `COPY` through).
- demo auto-set-dress preset — `dev-min`-style light seed vs a fuller demo preset — lean: a demo preset (e.g.
  `small-200`), since demos want a fuller world than dev; confirm at build.

## KB dependencies (read as contract)
- `corpus/ops/snapshot-spec.md`, `corpus/ops/seeding-spec.md`, `corpus/ops/db-access.md`
- the `dev-setdress` mechanism (M13), `corpus/ops/safety.md` (the capture-source policy + firewall it must preserve)
- `corpus/ops/demo/README.md` + the recipe family

## Delivers
- **→ rosetta:** `corpus/ops/snapshot-cold-start.md` — net-new: the fresh-box capture workflow + the MCP limitation +
  the sanctioned paths; + the demo recipes/skills updated for auto set-dress.
- **→ rosetta-extensions:** the set-dress chaining in `up-injected.sh` + the cold-start capture path.

## Risk
**(prod-safety, blocks-prod-safety)** auto set-dress + cold-start capture must preserve M13/M15's guarantees —
**read-only bounded capture, the tenant firewall, per-stack isolation, and a confirmation before any prod-touching
read**. Mitigate: reuse M13's set-dress pass verbatim (replay-only, never captures on the stack); capture stays
behind M9a's capture-source policy + `AssertPublicOnly`; the safety.md drift guards still hold.

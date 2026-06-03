---
milestone: M2b
slug: clerkenstein-consolidation
version: v1.0 "body double"
milestone_shape: section
status: planned
started:
last_updated: 2026-06-03
---

# M2b — Clerkenstein repo consolidation + knowledge base

## Goal
The `clerkenstein` repo (gitignored `anthropos-demo/clerkenstein`, its own git on `main`) grew
organically across M1/M1b/M2 into **flat package dirs** (`authn bapi orgclient fapi webhook cmd dna
golden golden-js scripts`) with a **single README** and **no knowledge base**. M2b reorganizes it into
a clean, self-documented **library-named** structure — one dir per mocked dependency + a `shared/` dir +
an `alignment/` harness dir + a `knowledge/` base — following `/singularity-kit:repo-consolidate`, so the
repo is navigable + operable by agents **before v1.0 ships**.

## Context (B-milestone — cleanup after M2)
Pure reorganization / documentation / hygiene over the **M2-complete** repo. **No behavior change** —
the two alignment gates (Go 22/22, JS 9/9) and the M1b drift harness (9/9) stay green throughout. The
move repoints imports + the DNAs/goldens/runners/scripts; it does **not** alter the mocks. Same class of
work as M1b (tooling/cleanup over a shipped surface), so it's a **B-milestone** appended after M2,
before `/developer-kit:close-release`. The user chose the **library-named** directory scheme (named after
the real dependency each dir mocks) over a surface-named or minimal-move scheme (decision M2b-D2).

## Scope
### In
1. **Restructure** to one dir per mocked library/framework + one shared dir, **library-named**:
   - `authn/` — mocks `colony/authn` (the `authn.Provider` twin).
   - `clerk-backend/` — mocks `clerk-sdk-go/v2` (the **`bapi` server + the `orgclient` store merged
     into one dir** — they are one library's mock split across two dirs today).
   - `clerk-frontend/` — mocks `@clerk/clerk-js` + `@clerk/nextjs` (the FAPI server + publishable-key codec).
   - `clerk-webhook/` — mocks `svix` (the webhook injector + event builders).
   - `shared/` — universal-key HS256 JWT + claims + canonical helpers (extracted because `clerk-frontend`
     **mints** and `authn` **verifies** with the same universal key — genuinely shared).
   - `alignment/` — the M0-consumption harness: `cmd/{clerkrun,jsfapirun}` + `dna/` + `golden{,-js}/` + `scripts/`.
   - **Tests stay co-located within each library dir.**
2. **A self-contained `knowledge/` base** documenting Clerkenstein — scope/goal; how it's built; how
   fidelity is **validated with alignment tests against a pinned Clerk version**; **per-library injection
   recipes**; a coverage index. Per-library `README.md`s + a top-level index. Solid, well-distributed.
3. **Hygiene** — an `.agentspace/` dir with contents **gitignored**; `.gitignore` cleanup; built-binary +
   transient hygiene per `repo-consolidate`'s asset-hygiene checks.
4. **Consolidate** — run `/singularity-kit:repo-consolidate code` to standardize the repo (emit
   `CLAUDE.md` + `singularity-manifest.md`, audit against the code-repo + asset-hygiene standards, apply
   fixes), then re-verify both gates + the drift harness.

### Out
- New library support / new alignment genes (the `@clerk/express` coverage gap stays a v1.1 item).
- Any **live injection** wiring into a running platform (still v1.1 / M3).
- Any change to rosetta's M0 framework (`test/alignment/`) or to the platform repos under `anthropos-dev/`.

## Sections
See `progress.md`. **S1** = restructure (move/merge into library-named dirs, repoint imports + scripts,
**gates stay green**). **S2** = the `knowledge/` base + per-library READMEs. **S3** = `.agentspace/` +
`.gitignore`/asset hygiene. **S4** = run `/singularity-kit:repo-consolidate code` + re-verify. **S5** =
slim rosetta's `corpus/services/clerkenstein.md` to a pointer + milestone records.

## Delivers → knowledge
- **Net-new:** the `clerkenstein` repo's own self-contained `knowledge/` base (+ `CLAUDE.md` +
  `singularity-manifest.md` from `repo-consolidate`).
- **Updated (rosetta):** `corpus/services/clerkenstein.md` slimmed to a pointer at the repo's `knowledge/`
  + the new library-named structure.

## Where it lives
The reorg + the new `knowledge/` base live in the `clerkenstein` repo (commits stack on its `main`, no
branch model). The rosetta-side milestone records + the corpus pointer land on the
`m2b/clerkenstein-consolidation` branch → merged to `release/01.00-body-double` at close.

## ⚠️ Process note — `repo-consolidate` is user-invoked
`/singularity-kit:repo-consolidate` is `disable-model-invocation` (the model can't auto-run it). S4's
consolidation run is therefore **user-driven**: the build authors the structure **to** repo-consolidate's
standard (so the run reports mostly-compliant), and the **user types the skill** pointed at the
`clerkenstein` repo. The build can otherwise complete S1–S3 + S5 autonomously.

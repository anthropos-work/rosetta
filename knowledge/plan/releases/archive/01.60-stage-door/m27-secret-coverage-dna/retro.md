# M27 — Retro

_Secret-coverage DNA + source ingestion. The first milestone of v1.6 "stage door". Closed 2026-06-14._

## Summary

M27 built the foundation of the secret-provisioning release: a new `stack-secrets` extension (in
`rosetta-extensions`, tagged `stage-door-m27` @ `195ef93`) holding a **values-blind secret-coverage DNA**
(55 genes across 6 repos) and a **DNA-driven dir/zip source reader** that structurally cannot ingest a
`zEnvs/` backup mirror or a stray `.env` — it opens exactly `<root>/<repo>/<target_file>` and never enumerates
the tree (the layout-contract defence). It also ships the hybrid `introspect` (the required set is the union of
`platform/.env_example` + sentinel + each frontend's `.env.example` + a curated compose/build-arg set — 8 of 9 Go
repos ship no example), the DNA-scoped **two-tier keep-listed `diff` gate**, and the `stacksecrets` CLI
(`list`/`check`(=`measure`)/`introspect`/`diff`). The `check`/`measure` scorer was folded in Fate-1 (the natural
pairing with the DNA); `provision` + the pre-flight wiring + demo-aware scoring stay M28. Verified live against
the real stack-dev: `diff` exits 0, `check` reports coverage with no secret value ever printed.

The close was clean: 2 findings, both Fate-1 ext doc-hygiene (the section was missing from the ext README index;
the README quoted a stale test count) — fixed in the ext commit `537aeff`. The deferral audit was GREEN.

## Incidents This Cycle

None. Zero P2 flakes, zero regressions. The build code held under both the harden pass (0 bugs over 2 passes,
+40 tests) and the close review (0 code-quality findings). The flake gate was 5/5 sequential `-shuffle` clean.

## What Went Well

- **The `datadna` mold paid off exactly as designed.** Reusing the proven `stack-seeding/dna` structure
  (criticality 3/2/1 → weight, the two-metric Overall/Critical score, the `ratio()` empty-denominator + the
  anti-vacuous-100 guard, the 0/1/3 exit-code contract) meant the DNA core was concrete construction, not
  exploratory — the milestone shipped as a section, as planned.
- **The values-blind invariant held end-to-end and is now machine-checkable.** Only `ClassifyShape` is permitted
  to read a value, as a discarded local; the committed `secret-dna.json` carries zero secret-shaped tokens. The
  harden pass pinned this with a 200 KB-value + `=`-in-value + quote-only adversarial extraction test and an
  end-to-end zero-leakage regression against the real 55-gene DNA.
- **The two-tier gate (M27-D2) is the headline design win.** A naive "every declared key must be a gene" gate
  produced 111 false unlisted-required findings (the `.env.example` files mix secrets with config noise). Scoping
  the gate-fatal class to the DNA's own tracked-secret universe made it honest AND usable — it caught the
  dangerous omission (a tracked secret missing a gene → vacuously-green coverage) without policing config noise.
  The build's diff-vs-stack-dev caught 10 real cross-repo omissions this way → all fixed Fate-1.
- **Stdlib-only was the right call.** No `go.sum`, no third-party code that could ever see a value — the
  values-blind audit surface is trivially small (M27-D3.1).

## What Didn't

- **The section README test count drifted** (quoted "94" from the build phase; ground truth was 113 after the
  harden pass added 40). Caught + fixed at close (TEST-1) — the handbook-count-reconciliation contract did its
  job, but the harden pass should have updated the count in-place when it added the tests.
- **The new section wasn't indexed in the ext top-level README** at build time (DOC-1) — caught + fixed at close.
  A new top-level unit should add its index row in the same change that creates the unit.

Both are minor doc-hygiene misses (the per-unit-handbook + index contracts are exactly the close-time net for
this class). No functional gap.

## Carried Forward

- **DEF-M27-02 — per-gene profile tag → Fate-2, M28-owned-if-needed.** Profile scoping settled to the `graphql`
  profile + the waived-class device for v1; the per-gene-tag variant is conditional residual M28 may revisit if
  it wires non-default-profile bring-ups. Already captured in M27-D3.4; no M28 `overview.md` edit required.
- **DEF-M27-01 — encrypted-zip (age/gpg) → DROPPED** as a documented v1 scope boundary (M27-D3.5), not a pending
  deferral (a crypto + key-management surface no consumer needs; plain dir/zip covers the M30 field-bake).
- **The `DIRECTUS_TOKEN` handoff to M28** (the highest-risk M28 interaction): `DIRECTUS_TOKEN` is modeled
  `key-present`-only (no `nonempty`) because the injection override strips it to `""` on non-prod / `--local-content`
  — M28's `provision` must defer to that strip and never re-arm the prod-write path. Recorded in spec-notes +
  decisions; the M28 risk-map already names it blocks-release with a regression-test requirement.
- Inherited release-level backlog (DEF-M10-01 / DEF-M21-01 / M25-D9) — KEEP, re-signed at the v1.5 close, all
  orthogonal to secret provisioning.

## Metrics Delta

(from `metrics.json`)

- **Go test funcs:** 867 → **980** (+113, entirely the new `stack-secrets` section: build + a 2-pass harden of +40).
  Per-module unchanged for the prior 4 sections; `stack-secrets` 113.
- **Python tests:** **459** unchanged (M27 is Go-only; the rosetta branch is markdown-only).
- **Coverage (harden):** statements 93.6% → 98.3% (+4.7); cmd/stacksecrets 87.4→97.9, secretdna 96.9→99.2,
  source 91.0→96.2.
- **Flake:** **0** (5/5 sequential `-shuffle` at close). `-race` + `gofmt` + `go vet` clean.
- **Review findings:** 2 (0 scope / 0 code / 1 docs / 1 tests / 0 blend), both Fate-1; 0 escape-hatch.
- **Field bugs:** 0 (the field-bake is M30); the build's diff-vs-stack-dev caught 10 DNA omissions → fixed Fate-1.

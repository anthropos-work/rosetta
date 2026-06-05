# M7b — progress (section checklist)

**Milestone:** M7b — The data-alignment dimension ("data DNA") · **Shape:** section · **Status:** done

## S1 — the data-DNA manifest format + the catalog ✅
- [x] gene format `{surface, schema, table, status, criticality, operators, expected-shape}` — `dna/dna.go` + `data-dna.json`
- [x] the seedable-surface catalog: 4 seeded (organizations/users/memberships/casbin-grant) + 6 planned — the M7c checklist (`catalog.go`, the `datadna catalog` CLI)

## S2 — the structural conformance operators ✅
- [x] `type-match` · `constraint-satisfied` (NOT-NULL + UNIQUE) · `fk-valid` · `row-count` (`operators.go`)
- [x] **separate `datadna` harness, NOT an alignctl runner** (M7b-D3) — one-sided + structural; reuses the M7a `pg` layer

## S3 — schema-as-source (the gene denominator) ✅
- [x] introspect the live schema per surface → ExpectedShape (`introspect.go` + the `pg` introspection builders)
- [x] `datadna introspect --stack demo-N` captures the contract (seeded-only); re-introspectable so `diff` is trustworthy

## S4 — drift detection (the M1b analogue) ✅
- [x] `datadna diff` exit-code contract (0 none / 1 schema moved / 3 usage) — `diff.go`
- [x] **PROVEN live:** added a column to `public.users` → `diff` flagged "added-column" exit 1 → revert → clean exit 0

## S5 — docs + the M7c gate ✅
- [x] the data dimension written into `corpus/architecture/alignment_testing.md` § "The data dimension (M7b)" + `dna/README.md`
- [x] the coverage metric defined as M7c's exit gate (seeded surfaces passing conformance ≥ threshold)
- [x] **acceptance:** catalog validates; operators run; `measure` 100%/Critical 100% on demo-1; `diff` flags a real schema change

**Build + proof status:** `dna/` + `cmd/datadna` + `pg` additions — **dna 49 + cmd/datadna 10 + pg 17 tests**
(build/vet/`-race`/gofmt clean). **LIVE on demo-1:** introspect (seeded-only) → measure **100% / Critical 100%**
(4 seeded surfaces) → diff drift flagged + cleared. Caught + fixed the planned-surface introspection bug;
hardened the UNIQUE leg of constraint-satisfied (M7b-D4). All sections done.

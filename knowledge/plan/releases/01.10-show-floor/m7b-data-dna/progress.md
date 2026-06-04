# M7b — progress (section checklist)

**Milestone:** M7b — The data-alignment dimension ("data DNA") · **Shape:** section · **Status:** planned

## S1 — the data-DNA manifest format + the catalog
- [ ] gene format `{surface, scenario, expected-shape, operator, criticality}` (the data analogue of `clerk-*.json`)
- [ ] the full seedable-surface catalog (~8–10 capabilities, isolation-class cross-checked vs M7a) — the M7c checklist

## S2 — the structural conformance operators
- [ ] `fk-valid` · `constraint-satisfied` · `type-match` · `row-count` (where M0's value-operators don't fit)
- [ ] reuse `alignctl` operator interface OR a sibling `datadnarun` runner (decide + record)

## S3 — schema-as-source (the gene denominator)
- [ ] capture the platform's expected shape per surface (ent introspection / `atlas inspect` golden)
- [ ] refreshable snapshot wired so the diff is trustworthy

## S4 — drift detection (the M1b analogue)
- [ ] `data-dna diff` exit-code contract (0 none / 1 schema moved / 2 gene regressed / 3 usage)
- [ ] proof: an injected schema change (added/renamed column on a seeded surface) is flagged

## S5 — docs + the M7c gate
- [ ] the data dimension written into `corpus/architecture/alignment_testing.md`
- [ ] the coverage metric defined as M7c's exit gate (threshold recorded)
- [ ] **acceptance:** catalog validates, operators run, drift diff flags a real schema change

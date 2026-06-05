# M7b — Retro

**Summary:** Extended the v1.0 alignment framework to a **third dimension — DATA**: a `datadna` harness
(`rosetta-extensions/stack-seeding/dna/`) that enumerates the seedable surfaces (the M7c catalog) and measures
whether a seeder's output **conforms to the platform's current schema** — with **drift detection** when the
schema moves. **Proven live** against the M7a-seeded `demo-1`: `measure` 100% / Critical 100%, `diff` flags an
injected column (exit 1) and reads clean on revert. The draft was authored by a sub-agent against a tight spec;
I verified the gates, ran the live proof (which caught a real bug), and hardened the UNIQUE leg.

## What went well
- **The alignment pattern genuinely transferred to data.** The manifest/score/diff *structure* carried over
  cleanly; the honest divergence (structural operators, one-sided output→schema) made it a separate harness, not
  a forced fit. The catalog *is* the M7c checklist, and the conformance + drift jobs fall out of the same DNA.
- **`introspect` as schema-as-source worked beautifully** — it captured the *full* live shapes (the seeders
  write a subset: `organizations` 6 of 10 cols, `memberships` 8 of 26 / 2 of 4 FKs), which is exactly the
  baseline a drift diff needs. The spec-authored shapes were a guess; the live capture is the truth.
- **The live proof earned its cost again.** It caught the planned-surface introspection bug (self-contradictory
  DNA) that every unit test passed over, because the unit tests don't introspect a real multi-schema DB.
- **The harden closed a genuine gap, not a coverage number** — the UNIQUE leg of `constraint-satisfied` was
  silently a no-op (the live impl didn't implement `DupLister`); wiring `CountDuplicates` made the operator
  actually complete, verified live.

## What didn't / constraints
- **The shipped `data-dna.json` was authored from the spec, not live introspection** (the sub-agent had no DB).
  The data-type spellings + column sets were incomplete until `introspect --stack demo-1` refreshed them — a
  reminder that the manifest's expected shapes must come from a real migrated stack, not a hand list.
- **The planned-surface table names are guesses** (`cms.content`, `skillpath.sessions`, `public.assignments`,
  `jobsimulation.activity` all introspected to 0 cols — wrong names; `skiller.skills` + `jobsimulation.sessions`
  exist). They're inert for M7b (planned surfaces aren't measured), but M7c should correct them when it builds
  those seeders (they're the catalog's targets).

## Carried forward → M7c
- M7c builds the planned seeders (taxonomy, content via Directus snapshot-replay, skillpath/jobsim sessions,
  assignments, activity); as each lands, **promote its surface planned→seeded and `introspect` its real shape**
  (fixing the guessed table names) so the data-DNA coverage gate rises. The coverage metric (seeded surfaces
  passing conformance) is M7c's exit gate.
- The feature-credit/tier grants beyond the minimum identity (so *every* feature-gated route authorizes, not
  just org-membership queries) — surfaced during M7a's proof, owned by M7c.

## Metrics
See [metrics.json](metrics.json). dna 49 + cmd/datadna 10 + pg 17 (M7a 62 + M7b harden) — all gates green
(build/vet/`-race`/gofmt). Live: `measure` 100%/Critical 100% on demo-1's 4 seeded surfaces; `diff` drift
flagged + cleared. 1 bug caught (planned-surface introspection), 1 harden gap closed (UNIQUE leg). Monorepo +1 commit.

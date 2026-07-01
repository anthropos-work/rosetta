# M52 — spec notes

_Technical notes accumulate here during build (file:line surfaces, rext tag, schema findings)._

## Topic → doc → code triples (Phase 0b anchor)

| Topic | Knowledge doc | Code path(s) |
|---|---|---|
| Mother prompt (generation template) | `ai-generation-spec.md` §1/§2b | `stack-seeding/blueprint/batch.go` — `DefaultBatchPromptTemplate` const, `Batch.expand`, `EffectiveBatches` |
| Prompt-hash cache key integrity | `cache-spec.md` §2 | `blueprint/batch.go` — `MotherPromptHash`; `stack-seeding/batchcache/cache.go` — `DefaultRootRel`, `Open` |
| Population blueprint (3 orgs incl. AI-readiness) | `seeding-spec.md`, `stories-spec.md` | `stack-seeding/presets/stories.seed.yaml`; `blueprint/stories.go`, `blueprint/blueprint.go` |
| Batch config (`--max-cost`/concurrency/re-roll) | `ai-generation-spec.md` §2c | `stack-seeding/cmd/gen-batch/` |
| Snapshot sources (taxonomy + Directus capture version) | `snapshot-spec.md`, `cache-spec.md` §2 | `stack-snapshot/manifest/manifest.go` — `SchemaVersion`, `CapturedAt`, `Source` |
| Cockpit download surface | `cockpit-spec.md` §Served endpoints | `demo-stack/cockpit.py` (`/manifest.json`), `seeders/cockpit.go` (`WriteCockpitManifest`), `demo-stack/up-injected.sh` (`--cockpit-export` wiring) |
| **Consolidated seed+gen manifest** | **`seed-manifest-spec.md` (NEW — M52 S4 deliverable)** | net-new: `seed-generation-manifest.yaml` + Go loader + `stackseed` export |

## Key architecture facts (verified against code, Phase 0b)

- **`DefaultBatchPromptTemplate` is a Go `const`** (`batch.go:111`) — NOT file-resident today; changing it needs
  a recompile. The **only** non-comment reference is `batch.go:254` (`Batch.expand`'s default when
  `b.PromptTemplate == ""`). S1 must move it to a checked-in file (`go:embed`) that renders the **byte-identical**
  effective mother prompt so the M45 cache key (`sha256(motherPrompt || capture-version)`) is preserved.
- **The cockpit `/manifest.json` today serves `cockpit-manifest.json`** — a stories→heroes projection
  (`CockpitManifest`: stack, stories[]→heroes[], deep_link_catalog). Built by `stackseed --cockpit-export`
  (`main.go:208`), written to `$STACK/cockpit-manifest.json` by `up-injected.sh`, served with
  `Content-Disposition: attachment; filename="cockpit-manifest.json"` (cockpit.py:365-375). S3 repoints the
  DOWNLOAD to the consolidated file; the MENU stays the projection.
- **The 3 orgs** live in `presets/stories.seed.yaml`: Cervato Systems (ai-transformation, 220), Solvantis
  (onboarding-ramp, 120), Northwind Aviation (ai-readiness, 200, the M51 showcase org).
- **Cache excluded from the manifest** per user note-3: `.agentspace/.batchcache` (git-ignored per-box) +
  generated member envelopes are NOT inlined — the manifest inlines only the DIRECTION (prompts + config +
  sources), not the derived data.

## Pre-flight audits — S1 (Extract the mother prompt to file-resident YAML)
- Report: `knowledge/plan/releases/01.10b-fit-up/m52-seed-manifest/kb-fidelity-audit.md`
- Verdict: **GREEN** (all PAIRED topics ALIGNED; the one BLIND-AREA is `seed-manifest-spec.md`, an explicit
  `Delivers →` milestone deliverable authored in S4 — satisfies its own gap per the audit's blind-area rule).

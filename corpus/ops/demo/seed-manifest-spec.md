# The Consolidated Seed+Generation Manifest (v1.10b "fit-up" M52)

> **One auditable file for the whole demo-data intent.** A presenter or auditor should be able to read
> the ENTIRE seed+generation direction — every org, every hero, the generation prompt, the batch config,
> and the snapshot sources — in **ONE checked-in file**, WITHOUT reading Go code. That file is
> `rosetta-extensions/stack-seeding/presets/seed-generation-manifest.yaml`, and the presenter cockpit's
> **[Download seed manifest]** serves it.
>
> This closes the "scattered intent" gap: before M52 the population lived in `stories.seed.yaml`, the
> generation batches in `gen-batch-*.seed.yaml`, and — the core gap — the **mother prompt** was embedded
> in a **Go const**, unreadable without opening the source. M52 makes the prompt file-resident and
> consolidates all of it into one manifest. **Entirely in `rosetta-extensions` — zero platform-repo edits.**

---

## For PMs — what it is

A single file that spells out, in plain YAML, everything that shapes a demo world's fake data:

- **who** — the orgs and their hero trios (thriving / struggling / manager);
- **how the supporting population is generated** — the exact instruction (the "mother prompt") the cheap
  LLM gets, plus the budget cap and the retry rules;
- **what real catalog it's set against** — the taxonomy + content snapshot versions.

It deliberately leaves out the *generated data itself* (the cache of AI-written profiles) — that's derived
output, not intent. The manifest is the **recipe**, not the meal. A presenter clicks **[Download seed
manifest]** in the cockpit and gets this recipe to read or hand to an auditor.

---

## For engineers

### 1. What it inlines (the whole DIRECTION)

The manifest (`SeedGenerationManifest`, `stack-seeding/manifest/manifest.go`) is a YAML document with four
substantive blocks:

| Block | What it carries | Source |
|---|---|---|
| `population.orgs[]` | every seeded org (name/slug/industry/narrative/size) + its `heroes[]` (key, name, role, vantage, trajectory, annotation, login, jump_to) | projected from the Stories & Heroes blueprint (`stories.seed.yaml`) — all 3 orgs incl. the M51 AI-readiness org (Northwind Aviation) |
| `generation.prompt_template` | the per-member **MOTHER PROMPT** verbatim — the exact instruction the cheap model gets | the file-resident `blueprint/prompts/default_batch_prompt.tmpl` (M52 S1) |
| `generation.config` | the batch RUN config: `model`, the **MANDATORY** `max_cost_usd` ceiling, `max_concurrent`, `max_rerolls`, `call_timeout` | the `gen-batch` CLI defaults, made file-resident |
| `generation.batches[]` | the batch descriptors (per `story_id`): `count`/`fill`, `roles`, `seniority`, `industry`, `narrative`, `bias_mix` | the generation preset (`gen-batch-org-fill.seed.yaml`) |
| `snapshot_sources` | the taxonomy + Directus **capture versions** the world is set-dressed from + the cache key is extended with | provenance (unpinned by default → "the capture the stack replays at bring-up") |

Plus a top `format_version` / `stack` / `description` header and a self-documenting `excludes:` block (§2).

`generation` is **omitted** entirely when the population declares no `batch[]` (a non-generated demo) — so a
population-only world shows no empty generation block.

### 2. What it deliberately EXCLUDES (the cache/generated-data boundary)

The user's note-3 boundary, made explicit in the file's own `excludes:` block:

- **`.agentspace/.batchcache`** — the per-box, git-ignored prompt-hash cache (the generated per-member
  envelopes + cost metadata). It is **derived data**, regenerated from the prompt; not intent.
- **generated per-member envelopes** (`member-*.json`) — the AI-written content itself.

The manifest inlines the **INTENT** (prompts + config + sources), never the **DERIVED** data. The cache is
covered by its own [`cache-spec.md`](cache-spec.md); this manifest is the direction that *produces* it.

### 3. How it's produced — a PROJECTION, honesty-gated (the D9 property extended)

The manifest is **not hand-authored** — it is **projected** from the canonical presets by
`manifest.Build()` and emitted by the `stackseed --manifest-export` verb:

```bash
stackseed --manifest-export \
  --seed presets/stories.seed.yaml \                    # population (all 3 orgs + heroes)
  --gen-seed presets/gen-batch-org-fill.seed.yaml \     # generation batches (merged by story id)
  --manifest-out presets/seed-generation-manifest.yaml
```

- `--seed` supplies the population (the Stories & Heroes blueprint).
- `--gen-seed` supplies the generation batches; they are merged onto the matching population story **by
  story id** (`mergeGenerationBatches`), so the ONE manifest inlines BOTH the seed AND the generation
  intent from the two existing single-sources — **no fabrication**.
- `--manifest-max-cost` (default `0.30`) sets the inlined `max_cost_usd` ceiling.

The checked-in `presets/seed-generation-manifest.yaml` is the canonical authored artifact (a prose header
+ the projected body). **The honesty gate:** `TestManifest_CanonicalFileMatchesProjection` (in
`cmd/stackseed`) re-derives the projection from the canonical presets and asserts the checked-in body still
equals it. So the "single auditable file" is also a **true** one — if a preset, the embedded prompt, or the
config changes without regenerating the manifest, CI fails loud. This extends the cockpit's D9 single-source
property (the menu can't drift from the seed) to the whole manifest.

### 4. The mandatory `--max-cost` ceiling is file-resident + validated

`--max-cost` is a **MANDATORY** `gen-batch` flag (no batch ever runs uncapped — the user's hard
requirement). M52 makes it **visible in the file** (`generation.config.max_cost_usd`) and
`SeedGenerationManifest.Validate()` **requires it `> 0`** whenever a generation block is present. An auditor
sees the budget cap in the manifest, and a zero-cap manifest fails validation — mirroring the CLI's own
guard.

### 5. The cache-key integrity rule (why the extraction is safe)

The M45 prompt-hash cache keys on the **rendered mother prompt** (`sha256(motherPrompt || capture-version)`,
[`cache-spec.md` §2](cache-spec.md)). M52 S1 moved the mother prompt from a Go `const` to the embedded file
`blueprint/prompts/default_batch_prompt.tmpl` — **byte-for-byte**. The embedded file's bytes EQUAL the
former const exactly, so the rendered prompt (and thus the cache key) is **UNCHANGED**: an existing cache
stays valid, the `$0` reseed holds. `TestDefaultBatchPromptTemplate_FileResident` pins that the embed
renders a well-formed prompt (a broken/empty `//go:embed` fails at test, never at a live generation run).

> **The invariant:** editing the `.tmpl` (or any manifest input) is exactly as cache-affecting as editing
> the old const would have been — a re-word re-keys the affected members (the documented, intentional
> invalidation); an identity-preserving move (M52's extraction) keeps every key. Never a *silent* cache
> bust.

### 6. How the cockpit serves it

The presenter cockpit ([`cockpit-spec.md`](cockpit-spec.md)) serves TWO manifests, for two different needs:

| File | Role | Cockpit surface |
|---|---|---|
| `cockpit-manifest.json` (the MENU) | drives the **[Log in as]** CTAs — the stories→heroes deep-link/seat-switch projection | read by the panel; served at `/manifest.json` (back-compat + fallback download) |
| `seed-generation-manifest.yaml` (the CONSOLIDATED manifest) | the **[Download seed manifest]** target — the auditable seed+gen intent | served verbatim at `/seed-generation-manifest.yaml` as a YAML attachment |

`up-injected.sh` exports both during bring-up (the consolidated one via `--manifest-export --gen-seed` the
org-fill preset when present) and passes `--seed-manifest` to `cockpit.py`. The consolidated-manifest serve
is **NON-FATAL**: a missing/broken file drops the richer download and the footer link falls back to the menu
manifest — a cockpit is never blocked (the M18/M19 non-fatal pattern). The MENU is unaffected either way.

### 7. Scope boundary (what M52 is NOT)

M52 only **EXPRESSES** the existing seed+generation behavior auditably — it adds **no new seeding
behavior** (M50/M51 own the seeders; M45/M46 own the generation engine). The manifest is a read-only
projection + a served download, not a new seed path.

---

## See also

- [`ai-generation-spec.md`](ai-generation-spec.md) — the generation engine + the CODE-owns-structure /
  AI-owns-content boundary + the mother prompt this manifest inlines.
- [`cache-spec.md`](cache-spec.md) — the prompt-hash cache the manifest deliberately EXCLUDES + the
  cache-key integrity rule (§5).
- [`cockpit-spec.md`](cockpit-spec.md) — the presenter cockpit whose [Download seed manifest] serves this
  file (§6).
- [`../seeding-spec.md`](../seeding-spec.md) — the seeding blueprint + the Stories & Heroes population the
  manifest projects.
- [`stories-spec.md`](stories-spec.md) — the 3-org Stories & Heroes world (incl. the M51 AI-readiness org).

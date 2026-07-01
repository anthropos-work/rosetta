# M52 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## D1 — Prompt extraction mechanism: `go:embed`, byte-identical
- **Context:** S1 must make the generation mother prompt file-resident (auditable, no-recompile-to-read)
  WITHOUT busting the M45 prompt-hash cache (keyed on the exact rendered mother prompt).
- **Options:** (a) `go:embed` the exact const bytes into a checked-in `.tmpl` file; (b) load the template
  from a YAML/JSON at runtime (os.ReadFile); (c) deliberately re-word + re-key the cache.
- **Choice:** (a) — `//go:embed prompts/default_batch_prompt.tmpl` into the same-named var, bytes extracted
  from the former const via a one-shot Go writer so they are byte-identical.
- **Why:** (a) keeps the cache key UNCHANGED (existing caches stay valid, $0 reseed holds) — the exact
  cache-integrity invariant the overview's Risk flags — while still making the prompt a plain file an
  auditor reads without Go. (b) adds a runtime read + a failure mode (missing file at gen time); embed
  fails at BUILD instead (safer). (c) needlessly invalidates a valid cache. Same-name var means zero
  consumer/test churn. `embed` is stdlib → no supply-chain change.

## D2 — Manifest is a PROJECTION of the presets, checked in + honesty-gated (not a hand-authored file)
- **Context:** S2 must produce ONE auditable `seed-generation-manifest.yaml`. Options for how it stays
  correct: (a) hand-author the YAML; (b) generate it from the existing canonical presets + a checked-in
  copy pinned by a test.
- **Choice:** (b) — a `manifest` Go package with `Build()` PROJECTS the manifest from the real blueprint
  (`stories.seed.yaml`) + the generation batches (`gen-batch-org-fill.seed.yaml`) + the embedded prompt +
  a config block. A `stackseed --manifest-export` verb emits it. The checked-in
  `presets/seed-generation-manifest.yaml` is the canonical artifact; `TestManifest_CanonicalFileMatches
  Projection` re-derives the projection and diffs the body — a drift fails CI.
- **Why:** extends the D9 single-source property (the cockpit menu already can't drift from the seed) to
  the whole manifest — a hand-authored file WOULD drift silently the moment a preset or the prompt
  changes. The projection guarantees the "single auditable file" is also a TRUE one.

## D3 — Population from stories.seed.yaml; generation batches merged from gen-batch-org-fill.seed.yaml
- **Context:** the population lives in `stories.seed.yaml` (all 3 orgs incl. the M51 AI-readiness org);
  the generation batch descriptors live in `gen-batch-org-fill.seed.yaml` (Cervato + Solvantis fill;
  Northwind is the curated showcase org, no fill batch). The overview requires ALL 3 orgs.
- **Choice:** `--manifest-export --seed stories.seed.yaml --gen-seed gen-batch-org-fill.seed.yaml` —
  `mergeGenerationBatches` attaches each gen story's batch[] to the matching population story by story id.
- **Why:** `stories.seed.yaml` is the canonical 3-org demo (the source of truth for the population +
  the AI-readiness org); `gen-batch-org-fill.seed.yaml` is the canonical generation intent. Merging them
  by story id inlines BOTH into one manifest from the two existing single-sources — no fabrication, no
  new "manifest-only" org/batch data to maintain.

## D4 — max-cost ceiling is inlined + validated (the user's hard requirement, made file-resident)
- **Context:** `--max-cost` is a MANDATORY gen-batch CLI flag (no batch runs uncapped). It was a
  runtime flag, not visible in any file.
- **Choice:** inline it into `generation.config.max_cost_usd` (default 0.30, `--manifest-max-cost`
  overridable) and `Validate()` REQUIRES it > 0 whenever a generation block is present.
- **Why:** the manifest is the auditable source of the whole generation intent; the budget ceiling is
  part of that intent. Making it file-resident + validation-enforced means an auditor sees the cap and a
  zero-cap manifest fails loud, mirroring the CLI's own guard.

## D5 — [Download] serves the consolidated manifest; the MENU stays the stories→heroes projection
- **Context:** the cockpit's `/manifest.json` currently drives BOTH the [Log in as] menu AND the
  [Download] link. M52 repoints the download to the consolidated file. Option: replace the menu manifest
  outright vs keep both.
- **Choice:** keep BOTH files. The MENU (cockpit-manifest.json) still drives the [Log in as] CTAs (it's
  the deep-link/seat-switch source — a different shape). The [Download] serves the NEW consolidated
  `seed-generation-manifest.yaml` (the auditable intent). The download endpoint is
  `/seed-generation-manifest.yaml`; `/manifest.json` stays served for back-compat + the fallback.
- **Why:** the two files serve different needs — the menu is a login launcher (needs jump_to/vantage
  labels), the download is an auditor's read (needs the prompt/config/sources). Conflating them would
  either bloat the menu or strip the download. A NON-FATAL fallback (has_seed_manifest / missing file →
  menu-manifest download) keeps an old bring-up's link alive. Live-verified byte-identical on a throwaway
  cockpit; demo-1 untouched.

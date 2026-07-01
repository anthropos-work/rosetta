# M52 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._

- [x] **S1 — Extract the mother prompt to a file-resident file** (`stack-seeding/blueprint/batch.go`
      `DefaultBatchPromptTemplate` const → the checked-in `prompts/default_batch_prompt.tmpl`, `go:embed`-loaded
      into the same-named var, editable without recompile). Cache-integrity: the embed is **byte-identical** to
      the former const, so the rendered mother prompt + the M45 prompt-hash cache key are **unchanged** — proven
      by `TestDefaultBatchPromptTemplate_FileResident` + every existing MotherPrompt/determinism/hash test
      passing untouched. No new dep (embed is stdlib); go.mod/go.sum byte-identical. rext @ `e57665f`.
- [x] **S2 — Author the consolidated `seed-generation-manifest.yaml`** — NEW `stack-seeding/manifest`
      package (`SeedGenerationManifest` schema + `Build` projection + `Validate` + `Write`/`Parse`) + the
      `stackseed --manifest-export` verb (`--gen-seed`/`--manifest-out`/`--manifest-max-cost`) + the
      checked-in `presets/seed-generation-manifest.yaml` (all 3 orgs incl. AI-readiness + the file-resident
      mother prompt + the batch config incl. the MANDATORY `max_cost_usd` ceiling + the snapshot sources +
      a self-documenting `excludes:` block for the cache/generated-data boundary). Honesty-gated by
      `TestManifest_CanonicalFileMatchesProjection` (checked-in body == live projection). No new dep;
      go.mod/go.sum unchanged. rext @ `0828f7f`.
- [x] **S3 — Repoint the cockpit [Download manifest]** — `cockpit.py` gets `--seed-manifest` + a new
      `/seed-generation-manifest.yaml` endpoint (verbatim YAML attachment) + the footer link repoint; the
      MENU stays the stories→heroes projection (drives [Log in as]). NON-FATAL fallback to the menu-manifest
      download. `up-injected.sh` exports the consolidated manifest (`--manifest-export`, `--gen-seed` the
      org-fill preset) + passes `--seed-manifest`. **Live-verified** byte-identical on a throwaway cockpit
      (:17799); demo-1 untouched. 6 new Python tests; demo-stack suite 311 green. rext @ `104896e`.
- [x] **S4 — NEW rosetta doc `corpus/ops/demo/seed-manifest-spec.md`** — the consolidated single-file
      seed+generation contract (what's inlined, what's excluded, how the cockpit serves it, the cache-key
      integrity rule §5, the honesty-gated projection). Cross-referenced from `demo/README.md` + `CLAUDE.md`
      + `cockpit-spec.md` + `cache-spec.md` + `ai-generation-spec.md` + `seeding-spec.md`; the stale
      "prompt-in-Go-const" (ai-generation §2b) + "download = menu JSON" (cockpit Served-endpoints) claims
      reconciled to the M52 reality. All cross-refs resolve.

## Status
All 4 sections landed. rext authoring copy @ `104896e` (S1 `e57665f` + S2 `0828f7f` + S3 `104896e`); the
`fit-up-m52` tag is cut at CLOSE, not here. Live-verify DONE (S3 cockpit download repoint, byte-identical,
demo-1 untouched).

## M52: Hardening

### Scope manifest (Phase 1)
The rosetta corpus M52 diff is **docs-only** (prose specs + plan files — no tests). The load-bearing M52
**code** footprint lives in the **rext authoring copy** (`.agentspace/rosetta-extensions`, branch `main`).
Touched source (grouped by stack), each with its co-located test:

- **Go — `stack-seeding`:**
  - `blueprint/batch.go` (`go:embed` extraction) — test `blueprint/batch_test.go` ✓ (baseline 97.2%)
  - `blueprint/prompts/default_batch_prompt.tmpl` — data (the extracted template)
  - `manifest/manifest.go` (NEW package: schema + `Build`/`Validate`/`Write`/`Parse`) — test
    `manifest/manifest_test.go` ✓ (baseline 90.2%)
  - `cmd/stackseed/main.go` (M52 funcs: `doManifestExport`, `mergeGenerationBatches`) — test
    `cmd/stackseed/main_test.go` ✓ (baseline 60.7% pkg, but M52 funcs 85.3% / 72.7%)
  - `presets/seed-generation-manifest.yaml` — data (the checked-in honesty-gated projection)
- **Python — `demo-stack`:**
  - `cockpit.py` (`--seed-manifest` + `/seed-generation-manifest.yaml` endpoint) — test
    `tests/test_cockpit.py` ✓ (baseline 97%)
  - `up-injected.sh` — shell (no unit-test harness)

Baseline coverage (M52-touched): manifest **90.2%**, blueprint **97.2%**, cmd/stackseed M52-funcs
**85.3%/72.7%**, cockpit.py **97%**. No new-unit-without-handbook (the NEW `manifest` package carries a
full doc-comment header; the S4 `seed-manifest-spec.md` is its corpus doc). No new dep in either stack.

### Pass 1 — 2026-07-01 (branch-coverage on the M52 footprint)
**Coverage delta (M52-touched files):**
- `manifest/`: statements 90.2% → **100.0%** (+9.8)
- `cmd/stackseed/` M52 funcs: `doManifestExport` 85.3% → **97.1%**; `mergeGenerationBatches` 72.7% → **100.0%**
- `cockpit.py`: 97% → **98%** (+1)

**Tests added:**
- `manifest/manifest_test.go`: 4 (seatless-hero skip, legacy single-org projection, storyAnnotation
  no-match fallthrough, Write encoder-error) + (Pass-3 note: batchesForStory out-of-range moved to P1) — all branch/error-path.
- `cmd/stackseed/main_test.go`: 5 (`--seed` load error, `--gen-seed` load error, post-merge
  validate-fail, `--stack` override, legacy single-org root-batch merge + empty-gen no-clobber guard).
- `demo-stack/tests/test_cockpit.py`: 1 (main() end-to-end WITH a real `--seed-manifest` — the
  success-read+serve branch; the prior test only drove the OSError fallback).

**Bugs fixed inline:** none — the M52 build was already correct; hardening deepened branch/error-path
coverage, no production-code change.
**Flakes stabilized:** none seen.

### Pass 2 — 2026-07-01 (the two load-bearing invariants, behavioral depth)
**Coverage delta:** ~0% lines (behavioral/regression depth on already-covered render + projection paths).

**Tests added:**
- `blueprint/batch_test.go`: `TestDefaultBatchPromptTemplate_CacheKeyGolden` — the **cache-key re-key
  tripwire**: pins the EXACT sha256 (`b4c09f94…`, len 1326) of a fully-rendered mother prompt from the
  file-resident default for a fixed context. The marker/determinism tests would pass on a `.tmpl`
  re-word that keeps markers + renders deterministically yet silently re-keys the M45 cache — this
  fails loud. **Mutation-verified:** a trailing space in the `.tmpl` trips it; restored byte-identical.
- `cmd/stackseed/main_test.go`: `TestManifest_HonestyGateHasTeeth` — the **honesty-gate sensitivity
  meta-test**: seeds a deliberate drift (mutate one projected hero name) + asserts the exact body
  comparison `TestManifest_CanonicalFileMatchesProjection` uses now DIVERGES from the canonical file. A
  toothless gate that always passes is worse than none.

**Bugs fixed inline:** none. **Flakes stabilized:** none.

### Pass 3 — 2026-07-01 (artifact consumability + stop)
**Coverage delta:** ~0% lines (artifact-consumability depth).

**Tests added:**
- `cmd/stackseed/main_test.go`: `TestManifest_CanonicalFileParsesAndValidates` — pins the SHIPPED
  `presets/seed-generation-manifest.yaml` is itself CONSUMABLE (re-parses through `manifest.Parse` with
  `KnownFields(true)` — no typo/unknown key survives — + passes `Validate`: 3 orgs + a generation block
  with the mandatory positive ceiling + a non-empty prompt). The honesty gate proves "matches the
  projection"; this proves "an auditor / the cockpit download can re-read it" — together they fence the
  file from both drift AND corruption.

**Out-of-scope note:** `demo-stack/up-injected.sh` (M52-touched) is bring-up orchestration glue
(`[ -f ]` gen-seed guard + non-fatal export fallback + `--seed-manifest` arg-wiring). No shell unit
harness exists; its behavior is already fenced by the Go CLI tests (`doManifestExport` ±`--gen-seed`,
unwritable-out) + the Python cockpit tests (present/absent non-fatal). End-to-end exercise belongs to
M53's cold-rebuild proof.

**Bugs fixed inline:** none. **Flakes stabilized:** none.

**Knowledge backfill:** no KB-worthy NEW findings — every invariant the harden pass pinned (the
byte-identical embed / cache-key preservation, the honesty-gated projection, the mandatory ceiling, the
cache/generated-data exclusion) is ALREADY documented in `corpus/ops/demo/seed-manifest-spec.md` (§5
cache-key integrity, the projection + excludes contract) + `cache-spec.md` §2. The tests reference those
docs in their names/comments. No doc edit needed; the S4 spec already carries the truths the tests pin.

### Stop condition
Loop stopped after **3 passes**: the full Step 2b six-dimension scan found nothing new worth adding
(qualitative), coverage deltas negligible (<2% after Pass 1; manifest at 100%, blueprint 97.2% with the
residual being pre-M52 out-of-scope error paths, cockpit.py 98% with the residual being uninstrumentable
interpreter idioms — `except KeyboardInterrupt` / `__main__`), and no flaky tests (3 consecutive clean
sequential runs of both the Go and Python suites). Full verification: 13 Go packages ok / 0 failures,
312 Python tests passed, `go vet` clean.

**Coverage delta (start → end, M52-touched):** manifest 90.2% → **100%**; cmd/stackseed M52-funcs
85.3%/72.7% → **97.1%/100%**; cockpit.py 97% → **98%**; blueprint stayed 97.2% (Pass-2/3 added
behavioral re-key + consumability depth, not lines).

## M52: Final Review

_Close review (2026-07-01). Phases 1–5 ran as parallel scans (deferral audit blocking-gate GREEN →
`audit-deferrals/deferral-audit-2026-07-01-m52-close.md`; code-quality + adversarial + test-coverage in the rext
authoring copy; docs review in the corpus). Section shape → Phase 9 Completeness Ledger. Default fix-everything;
no escape-hatch deferrals surfaced. Ground truth at review: Go 13 pkgs pass / 0 fail (2 runs, no flake), `go vet`
clean; Python 312 pass._

### Scope
- [x] All 4 sections (S1–S4) checked off; harden Pass 1–3 + stop-condition complete. No scope gap.
- [x] DEF-M52-01 (up-injected.sh end-to-end glue) → Fate-2 confirmed → M53 (cold `/demo-up` runs up-injected.sh;
      M53 `overview.md In:` asserts the cockpit [Download] returns the complete inlined manifest). No plan edit.
- [x] No TODO/FIXME/HACK in any M52-touched rext source (grep clean).

### Code Quality
- [x] [should-fix] F1 — 3-way duplicated projection helpers (`normalizeVantage`/`storyAnnotation`/
      `batchesForStory`) across `blueprint`/`seeders`/`manifest`: drift risk the honesty gate can't catch
      (it compares to the checked-in FILE, not the seeder's projection). → export canonical `blueprint`
      helpers; `manifest` + `seeders` consume the ONE implementation.
- [x] [should-fix] F2 — `TestManifest_HonestyGateHasTeeth` hand-built description says `"generation data
      excluded"` vs production `doManifestExport`'s `"generated data excluded"` — the teeth-test's non-mutated
      fields must mirror production exactly. → fix the test string.
- [x] [nice] F9 — `manifest.Build` doc-comment references a vestigial `reservedNames` param the signature
      doesn't take. → correct the comment.
- [x] [nice] F10 — `SnapshotSources.TaxonomyCaptureVersion`/`DirectusCaptureVersion` never wired non-empty
      (no CLI flag). CONFIRMED intentional provenance scaffolding (documented "unpinned by default" +
      tested) — no code change; noted here for the audit trail.

### Adversarial (Fate-1 — risk is in code that just shipped)
- [x] F3 — cache-key golden pins ONE render context (non-empty ReservedNames/Industry/Narrative), blind to
      the `.tmpl`'s inactive `{{else}}(none){{end}}` branch: a fill batch in a hero-less (empty ReservedNames)
      story could silently re-key the M45 cache — the exact risk the overview's Risk section flags. → add a
      second golden pinning the EMPTY-ReservedNames render (covers the `(none)` branch).
- [x] F4 — `mergeGenerationBatches` silently DROPS an orphan gen-story id (a gen-id typo → a generation-less
      manifest, no error) + last-wins on dup gen ids + never validates the gen file independently. → warn to
      stderr on an unmatched gen story id (fail-visible, still non-fatal) + add multi-story merge tests
      (correct-story routing, orphan-drop-is-warned, dup-id, population-batch-preservation).
- [x] F5 — cockpit "non-fatal fallback" guards only `OSError`; an empty/corrupt-but-readable `--seed-manifest`
      serves a hollow 0-byte attachment with no fallback. → treat an empty/blank read as absent (fall back to
      the menu manifest + warn) + test the empty-file case.
- [x] F6 — honesty-gate teeth-test proves sensitivity on the POPULATION axis (hero name) only. → add a
      generation-axis perturbation (mutate `max_cost_usd`) asserting the gate diverges — pins teeth on both halves.

### Documentation
- [x] F7 — `corpus/ops/demo/cockpit-spec.md:63-65` (the "For PMs" prose) still claims [Download] "saves the
      cockpit's JSON menu" — stale after the M52 repoint + inconsistent with the same file's own updated
      Served-endpoints table (:161) + prose bullet (:139-146). S4 reconciled the table but missed this
      earlier block. → reconcile the prose to the consolidated-manifest reality.

### Decision Triage
- [x] D1 (go:embed byte-identical prompt extraction) → already blended into `seed-manifest-spec.md §5`
      (cache-key integrity) — verify + add `(#M52-D1)` ref-tag.
- [x] D2 (manifest is a honesty-gated PROJECTION, not hand-authored) → already blended into
      `seed-manifest-spec.md §3` (the honesty gate) — verify + add `(#M52-D2)` ref-tag.
- [x] D3/D4/D5 → archive (maintainer-only: population/gen-batch merge source split, max-cost inlining
      mechanics, dual-manifest keep-both rationale — all captured in the spec's contract prose; no separate blend).

### Regression / test-depth
- [x] F8 — `stripLeadingComments` (the honesty-gate's load-bearing comparison helper) is itself untested. →
      add a direct unit test (strips the header comment block, preserves the YAML body).

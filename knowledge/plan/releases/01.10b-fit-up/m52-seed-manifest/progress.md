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

# M52 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._

- [x] **S1 — Extract the mother prompt to a file-resident file** (`stack-seeding/blueprint/batch.go`
      `DefaultBatchPromptTemplate` const → the checked-in `prompts/default_batch_prompt.tmpl`, `go:embed`-loaded
      into the same-named var, editable without recompile). Cache-integrity: the embed is **byte-identical** to
      the former const, so the rendered mother prompt + the M45 prompt-hash cache key are **unchanged** — proven
      by `TestDefaultBatchPromptTemplate_FileResident` + every existing MotherPrompt/determinism/hash test
      passing untouched. No new dep (embed is stdlib); go.mod/go.sum byte-identical. rext @ `e57665f`.
- [ ] **S2 — Author the consolidated `seed-generation-manifest.yaml`** — ONE checked-in file inlining the
      population blueprint (all 3 orgs incl. the M51 AI-readiness org), the generation prompt templates, the
      batch config (`--max-cost` ceiling, `--max-concurrent`, re-roll rules), and the snapshot sources
      (taxonomy + Directus capture versions). EXCLUDES `.agentspace/.batchcache` + generated member
      envelopes. A reader/loader (Go) that parses it + a `stackseed --manifest-export` (or equivalent) that
      emits/validates it for the cockpit.
- [ ] **S3 — Repoint the cockpit [Download manifest]** (`demo-stack/cockpit.py` + the `up-injected.sh`
      export wiring) to serve the consolidated `seed-generation-manifest.yaml` as the download (replacing the
      stories→heroes projection as the *download* target; keep the projection as the menu source).
- [ ] **S4 — NEW rosetta doc `corpus/ops/demo/seed-manifest-spec.md`** — the consolidated single-file
      seed+generation contract (what's inlined, what's excluded, how the cockpit serves it, the cache-key
      integrity rule), cross-referenced from `demo/README.md` + `seeding-spec.md` + `ai-generation-spec.md` +
      `cockpit-spec.md`.

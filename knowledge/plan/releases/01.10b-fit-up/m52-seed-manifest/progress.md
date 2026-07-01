# M52 — progress

## Section checklist
_Checked off as each In-scope deliverable lands. Close when all boxes are ticked._

- [ ] **S1 — Extract the mother prompt to a file-resident YAML** (`stack-seeding/blueprint/batch.go`
      `DefaultBatchPromptTemplate` const → a checked-in `go:embed`-loaded prompt file, editable without
      recompile). Cache-integrity: the embedded template must render the **byte-identical** effective
      mother prompt (so the M45 prompt-hash cache stays valid) — proven by a test that diffs the old const
      vs the new embedded template's expansion.
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

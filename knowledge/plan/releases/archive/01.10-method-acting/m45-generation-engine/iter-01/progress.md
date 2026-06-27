**Type:** tok (bootstrap)

# iter-01 — bootstrap tok

Authored the M45 iteration protocol + the engine's opening strategy. No production engine code lands this
iter (the bootstrap tok is setup work); iter-02 is the first tik.

## Work done
- Ran the mandatory Phase 0b **KB-fidelity gate** -> **GREEN** (report `kb-fidelity-audit.md`): all 5
  KB-dependency topics PAIRED + ALIGNED; the 2 NEW docs are intentional DOC-ONLY Delivers targets.
- Probed the SANCTIONED new dep: `github.com/anthropos-work/ai@v1.40.1` fetches via GH_PAT (RC=0). Read
  the real API: `ai.AI.ChatCompletion(ctx, []ChatCompletionMessage, ...Option) (ChatCompletionMessage,
  *MetaData, error)`; constructors `openai.New`/`NewOpenAI`/`NewAzure`; options `WithModel`,
  `WithResponseFormat(JSONObject)`, `WithTemperature`, `WithSeed`, `WithMaxCompletionTokens`;
  `MetaData.Usage` -> `*openai.CompletionUsage` (`PromptTokens`/`CompletionTokens`/`TotalTokens int64`).
- Authored **`corpus/ops/demo/ai-generation-spec.md`** (NEW) — the engine reference AND the iteration
  protocol (the measure->fix->accept loop; the 5-metric gate vector; fixtures-first discipline; the
  CODE-owns-structure / AI-owns-content boundary).
- Authored **`corpus/ops/demo/cache-spec.md`** (NEW) — the prompt-hash cache (mother-prompt + capture-
  version key, atomic writes, `.lock`, the $0 re-seed). Placement resolved to `corpus/ops/demo/` for
  demo-index consistency (TOK-01 decision; overview said `corpus/ops/cache-spec.md` — superseded).
- Recorded **TOK-01** (the initial strategy) in the milestone-root `decisions.md` + the audit verdict in
  `spec-notes.md`.

## Close — 2026-06-26

**Outcome:** Initial strategy (TOK-01) authored — build the engine inside-out, fixtures-first, in
dependency order (services/ai -> blueprint.Batch -> cache -> cmd/gen-batch -> GeneratedBatchSeeder ->
collision-avoidance -> gate-proving). Iteration protocol + cache spec landed. KB-fidelity GREEN, `ai` dep
fetchable.
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap, does NOT exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** TOK-01 (milestone-root), D1 (cache-doc placement), D2 (dep pin)
**Side-deliverables (if any):** none
**Routes carried forward:** iter-02 builds the `services/ai/` wrapper (component 1) under TOK-01.
**Lessons:** the `ai` lib's split (transport in the lib; routing + cost in the consumer) means the wrapper
is genuinely thin — do NOT reimplement provider internals; the model->price table is the only cost logic
the wrapper owns. `MetaData.Usage` is `any` (a `*openai.CompletionUsage`) — type-assert it for token counts.

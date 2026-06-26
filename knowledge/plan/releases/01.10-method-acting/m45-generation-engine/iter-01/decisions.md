# iter-01 decisions

| ID | Decision | Rationale |
|----|----------|-----------|
| D1 | Place the cache doc at `corpus/ops/demo/cache-spec.md` (NOT `corpus/ops/cache-spec.md` as the overview's Delivers line said). | The overview explicitly allowed "NEW, or a Caching section in ai-generation-spec.md"; the demo-family index convention (CLAUDE.md, README) keeps demo docs under `corpus/ops/demo/`, and `ai-generation-spec.md` (also under `demo/`) links it as a sibling. Index-consistency wins. |
| D2 | Pin the `ai` dep at `v1.40.1` (the platform consumers' pin per `shared_libraries.md`). | Reproducible; matches the version the corpus documents + the platform uses, so the wrapper compiles against the same contract. Full license-vet of the transitive tree (openai-go/v3, retry-go, tidwall json) happens at the iter-02 `go get` + at close. |

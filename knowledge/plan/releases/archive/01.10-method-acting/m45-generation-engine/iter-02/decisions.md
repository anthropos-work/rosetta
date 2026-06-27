# iter-02 decisions

| ID | Decision | Rationale |
|----|----------|-----------|
| D1 | Pin `ai` at `v1.40.1` + accept its transitive tree (openai-go/v3 Apache; retry-go/tidwall/samber/json-iterator/regexp2 MIT; Azure SDK/MSAL/jwt MIT). | All permissive (MIT/Apache/BSD); `anthropos-work/ai` is first-party. The footprint grew (Azure SDK + openai-go/v3) — flagged for the close-time supply-chain review, but no blocker. |
| D2 | The wrapper owns ONLY routing + cost (model->price table mirroring `app/internal/aiusage`); it does NOT reimplement provider internals. | The corpus contract (`shared_libraries.md` §ai): the lib owns transport, the consumer owns EU-first routing + dollar cost. Keeps the wrapper genuinely thin. |
| D3 | `--max-cost` ceiling fails CLOSED (ceiling<=0 -> no spend allowed) via `WouldExceed`. | The user's hard requirement: no batch ever runs uncapped. A missing/zero ceiling must abort, not run free. |

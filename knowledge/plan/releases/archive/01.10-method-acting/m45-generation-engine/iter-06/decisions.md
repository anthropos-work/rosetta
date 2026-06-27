# iter-06 decisions

| ID | Decision | Rationale |
|----|----------|-----------|
| D1 | The `GeneratedBatchSeeder` reuses `resolveJobRoleRefs`/`resolveNamedSkillRefs` UNCHANGED — no resolver edit. | The resolvers already return zero values (blank role, empty pool) on a no-match — that IS the drop-not-fabricate seam. The seeder just skips the zero value. Keeps M45 within its "reuse the seeder machinery unchanged" scope + keeps the closure gate green by construction. |
| D2 | Generated members occupy a HIGH user-index band (generatedBaseIndex=100000 + a per-batch-id band of 1000). | The curated population caps well under 100000; the offset guarantees a generated user index never collides with a curated one without coordinating the two index spaces. |
| D3 | M45's GeneratedBatchSeeder routes ALL generated members into the FIRST resolved story's org. | M45 proves the engine on a BOUNDED batch within the existing org; per-story batch routing (which story's batch[] seeds which org) is an M46 concern (org-scale fill). Keeps M45 focused on the engine + cache, not the routing. |
| D4 | The email local-part is sanitized ([a-z0-9._-]) from the envelope's email_local (or name-derived). | An LLM-suggested local part is AI content; sanitizing it keeps an invalid/injected address out of the row (CODE owns identity, the AI only suggests). |

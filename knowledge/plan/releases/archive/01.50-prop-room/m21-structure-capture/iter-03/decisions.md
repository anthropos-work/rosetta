# M21 iter-03 — decisions (iter-local)

_(Cross-iter decisions are M21-D5 (refined) + M21-D6 in the milestone-root `decisions.md`.)_

- **iter-03-L1 — Why close partial rather than push the apply now.** With the real DDL in hand, creating the 9 tables
  is trivial, but a faithful stage-3 needs the 217 `directus_fields` + 43 `directus_relations` registry rows loaded,
  AND the digest-keying fork (option A/B) decided before the apply is wired into stacksnap (else replay still exits 5).
  Rather than hand-load 217 registry rows through MCP-JSON and bake in a keying choice unilaterally, iter-03 closes on
  the source-resolution + the architectural finding, and surfaces the keying fork to the operator. iter-04 applies.
- **iter-03-L2 — The prod structural read is bounded + read-only.** Only information_schema/pg_catalog types + the
  registry inventory (counts) were read; no tenant rows, no bulk content pull. Within the M9a capture-source policy
  (operator-confirmed) + AssertPublicOnly extended to structural metadata.

# M21 iter-05 — decisions (iter-local)

_(Cross-iter: M21-D9 in the milestone-root `decisions.md`.)_

- **iter-05-L1 — the serve recipe (empirically pinned).** Beyond stage 4 (tables + rows + digest), serving a
  collection anonymously needs: (1) the table has a PRIMARY KEY (Directus ignores PK-less collections); (2) a
  `directus_collections` registration row; (3) a `directus_permissions` read row on Directus's hardcoded public policy
  `abf8a154-5b1c-4a46-ac9c-7300570f4f17` (bootstrap creates it + the `(null,null)` access link), `fields='*'`, with
  the `status=published` filter for simulations/skill_paths. `directus_fields` rows are NOT required (DB-column
  introspection covers them).
- **iter-05-L2 — harness retained for iter-06** (the code-ification will test against / rebuild from this gate-passing
  state). Torn down at session end.

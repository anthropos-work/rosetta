# M21 iter-06 — decisions (iter-local)

_(Cross-iter: M21-D10 in the milestone-root `decisions.md`.)_

- **iter-06-L1 — scoped the iter to the capture CORE.** The full code-ification (capture → apply → serve-rows →
  exit-code redefinition → firewall admissibility) is multi-iter. iter-06 lands the data model + the structure-SQL
  generator + tests (a clean, DB-free-testable unit, live-validated against prod). The apply + serve-row capture +
  wiring are iter-07/08 — keeping each iter completable.
- **iter-06-L2 — dynamic capture over hardcoded list.** `CaptureStructure` enumerates collections from the catalog
  (intersected with information_schema), so it tracks the source content model across versions rather than pinning a
  26-name list.

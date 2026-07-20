# M241 — Decisions

_(implementation choices with rationale accumulate here during build)_

- **User decision (2026-07-20): language → EN-only fallback per tuple.** M241 opens with a read-only prod pool-count query (IT sessions per requirement tuple); toggle where IT exists, EN-only where absent. No blocking. Recorded at design time; carried here for build traceability.

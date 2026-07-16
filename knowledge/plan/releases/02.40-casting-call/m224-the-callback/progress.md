# M224 — Progress

_Iterative milestone: a **running ledger**, not a section checklist. `/developer-kit:build-mstone-iters` appends one
entry per iter and creates `iter-NN/` dirs as it goes. iter-01 is the BOOTSTRAP tok (authors the first strategy)._

## Running ledger

| iter | kind | what changed | gate metric | outcome |
|------|------|--------------|-------------|---------|
| iter-01 | tok (bootstrap) | KB-fidelity GREEN (hiring.md FAPI-pointer fix inline); authored TOK-01 (recruiter-render-first) | baseline UNMEASURED (presumed 0 rows) | closed-fixed — see iter-01/progress.md |
| iter-02 | tik | Clerkenstein org `publicMetadata.isHiring` wired end-to-end (seeder roster → FAPI); `/align-run` GREEN 100/100 ×2; rext tag `casting-call-m224-iter02` | UNCHANGED (fix-half/scaffold — no render yet) | closed-fixed — see iter-02/progress.md |

## Next iter

**iter-03 (tik, under TOK-01):** the **hero seats** — add the recruiter (`vantage: manager` → `jump_to` the
comparison surface) + 2 candidate exemplars (one assigned-AND-assessed, one only-assigned, same position →
`/profile`) to the Meridian Talent 4th story in `presets/stories.seed.yaml`, and make the `HiringFunnelSeeder`
**hero-aware** so a candidate hero is pinned to its declared funnel stage (assessed vs assigned-only) and the
recruiter resolves to `role=candidate`-org admin. Handle `roleForHero` in a candidate-only org (no `member`).
Unit-test the funnel hero-stage + roster seats; re-tag rext. Then **iter-04** = LOCAL demo bring-up at the tag +
the render-probe + the **baseline measurement + attribution** (the first gate reading).

---
iter: 01
milestone: M2c
iteration_type: tok
tok_flavor: bootstrap
status: closed-no-lift
created: 2026-06-03
---

# M2c / iter-01 — bootstrap tok (author TOK-01)

**Type:** tok (bootstrap)

## Inputs
`overview.md` (scope, exit_gate, the RS256 wall), `spec-notes.md` (the gap + 2 options + gene proposal),
the alignment protocol (`corpus/architecture/alignment_testing.md`), and the 3-agent research workflow.

## Bootstrap investigation — the unknowns this tok resolves
- **@clerk/express offline?** **YES** — `v1.7.79` in `studio-desk/node_modules` (+ `@clerk/backend`),
  node v22 → a **Node runner driving the REAL SDK** is feasible (the svix-pattern, highest fidelity; no
  Go-verifier fallback needed). (Note: declared `^1.3.47`, installed `1.7.79` — same major.)
- **Shared token (additive vs migration)?** studio-desk runs its **own** Clerk instance (own
  `CLERK_SECRET_KEY` / `CLERK_SIGN_IN_URL`) → a **separate token domain** from the main app → **additive
  RS256 is viable** (studio-desk RS256; main app stays HS256). Confirm per-tik via re-survey.
- **Runner shape:** Node (real `@clerk/express`), not a Go RS256-verifier fallback.

## Initial strategy → TOK-01
**RS256-native, additive-first, real-SDK runner.** Full record in the milestone-root `decisions.md` § TOK-01.

## Next-tik direction (iter-02)
Author `clerk-express-1.json` (the ~8 genes from `spec-notes.md`) + validate with `alignctl dna validate`.

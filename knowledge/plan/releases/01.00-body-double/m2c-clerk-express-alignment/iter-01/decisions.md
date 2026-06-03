# M2c / iter-01 — decisions

## iter-01-D1 — bootstrap findings (the 3 unknowns, resolved)
1. **`@clerk/express` is offline-available** — `v1.7.79` + `@clerk/backend` in `studio-desk/node_modules`,
   node v22 → a **Node runner driving the real SDK** is feasible (high-fidelity; svix-pattern). Resolves
   M2c-D5 toward a Node runner (not the Go-verifier fallback).
2. **studio-desk runs its own Clerk instance** (own `CLERK_SECRET_KEY` / `CLERK_SIGN_IN_URL`) → a
   **separate token domain** from the main app → **additive RS256 is viable**. Resolves M2c-D2 toward
   **additive** (escalate to migration only if a tik proves the token is shared across apps).
3. **Runner shape = Node** (real `@clerk/express`).

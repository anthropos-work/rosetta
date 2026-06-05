# M8 — Decisions

## M8-D1 — the open questions, resolved at build (2026-06-05)
The kickoff open questions all resolved during the build:
- **Recipe count:** 3 end-to-end recipes (enterprise-onboarding · skill-progression · browser-login) — the
  browser-login one absorbs the 2 M3-deferred injection recipes (cert-redirect + walk-through), so they share a
  home rather than spawning 2 thin docs.
- **`corpus/ops/demo/` subdir vs flat:** **subdir** with a `README.md` family index — a third indexed family
  (parallel to the `staging-*` cluster) that keeps the recipes grouped; the M3 lifecycle guide
  (`rosetta_demo.md`) + the M7 seeding spec (`seeding-spec.md`) stay flat in `ops/` (M3/M7 deliverables) and are
  indexed from the family README.
- **`/demo-seed` skill:** a **thin wrapper skill** (over `stackseed`) for parity with `/demo-up|down|status`.
- **Preset count:** **3** (small-200 / mid-500 / large-1k); `mid-500` + `large-1k` validated to seed end-to-end.
- **AI-content STRETCH:** **not pulled forward** — the hard line holds (M7c waived taxonomy + content; AI content
  is a v1.2 theme). Out of M8.

## M8-D2 — the express-gate CI carry-forward lands here, validated (2026-06-05)
The v1.0 `@clerk/express` gate was un-wired to CI ("needs a Node step"). M8 wires it into clerkenstein
`alignment.yml` (setup-node + `npm install @clerk/express` + the gate) and **validates it locally 9/9** with a
freshly-installed SDK — not authored-and-untested. The 4th surface (deployment/injection) **stays a local gate**
(it needs the platform's `colony` via GH_PAT, which a public runner can't supply cleanly) — documented, not
forced into CI.

## M8-D3 — the cert-redirect is a documented recipe, not a live-wired step (2026-06-05)
The `api.clerk.com` → fake-BAPI cert-redirect (the backend orgclient seam) is documented in
`recipe-browser-login.md` (consistent with M3's "recipe-only" status): the fake BAPI is built + alignment-gated,
but the DNS/cert wiring into a demo stack is the recipe. The backend authn seam (the 403→200 path) works without
it; the cert-redirect is for the orgclient reads.

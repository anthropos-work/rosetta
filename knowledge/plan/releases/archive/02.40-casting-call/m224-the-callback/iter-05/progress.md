# iter-05 — progress (tik, under TOK-01) — THE DIAGNOSIS

**Reconciled at the TOK-02 boundary** (executed + committed during a driven diagnosis leg; record completed here).

## What happened
With the data side MET (iter-04: 5 sims × 43 candidates seeded), the recruiter STILL could not reach the scoreboard
— the render-probe timed out sitting on the **real Clerk sign-in page** (login never settled against Clerkenstein).

**Root cause (code-cited, read-only against the demo clone):**
- `apps/web/src/context/UserStatusContext.tsx:141-173` — a user whose memberships are ALL hiring-orgs is **ejected
  out of the workforce app** to `hiring.anthropos.work` (`buildSwitchHandoffUrl`, `targetProduct:'hiring'`), **by
  design** (the platform enforcing the web↔hiring product boundary).
- `useGetClerkOrganization` filters hiring orgs out of the workforce list.
- The Hiring sub-app is **not in the demo**, so the eject lands the recruiter on a dead login page.

**Isolation:** `dan-manager` (workforce) logs in + renders `/enterprise/activity-dashboard` fine on the SAME warm
stack → the demo login + apps/web work generally; the divergence is specific to the **genuine-hiring-org** recruiter.

## Attribution
**Product-boundary eject — NOT a render-gate, NOT a Clerkenstein bug, NOT a seed gap.** The iter-02 isHiring wiring
WORKED (the app reads `org_8ff…=Meridian, isHiring=true`); that same flag is what the platform ejects on. So
"genuinely reads as hiring" and "scoreboard reachable in apps/web" are **mutually exclusive on the unmodified
platform** → **falsifies M222's premise** (its apps/web conclusion held only because the spike org lacked client
`publicMetadata.isHiring`).

## Fix landed (cheap)
Render-probe timeout `150s → 300s` (kept the fullPage-hang `shoot()` fix from iter-04; tsc clean). rext tag
`casting-call-m224-iter05`; rosetta `ae4974e` (spec-notes attribution).

## Outcome
**closed-fixed (attribution).** The attribution resolving to a product boundary (not a render-gate) is what
**triggers TOK-02** (iter-06): pivot from "patch the eject + re-skin apps/web" → "run the real Hiring app in the
demo." See **TOK-02** in `decisions.md`.

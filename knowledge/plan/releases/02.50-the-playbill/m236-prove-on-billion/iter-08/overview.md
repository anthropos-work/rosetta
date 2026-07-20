---
milestone: M236
iter: 8
iteration_type: tik
status: closed-fixed
created: 2026-07-20
handler: ACADEMY-M236-iterTBD-catalog-fill
---

# iter-08 — the academy pair (the last one)

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove) — Phase L, final arm.

## Step 0 — re-survey

28/29 after iter-07; the single residual is `skill-path-new/academy-foundation-of-ai/player`,
`route rendered a not-found`. Target current and unambiguous.

## Cluster / target identified

Handler `ACADEMY-M236-iterTBD-catalog-fill`, plus the M230 carry-forward cluster-1 item the release routed
here ("the academy grid renders real cards" — the Thread A half of this milestone's exit gate).

**The survey already refuted the routed-forward plan**, which was: *wire `app/cmd/academy-seed` into the
cold bring-up for real `academy_chapter_progress`, then re-point the CTA to the authed progress-bearing
chapter route.* Three facts, all obtained live, say otherwise:

1. **There is no `/library/[slug]` route in ant-academy at all** — only `app/(public)/library/page.jsx`, the
   index. The M235 seeder comment asserting "`/library/<slug>` — a REAL public course (resolves,
   non-fabricated)" was authored offline and is false. The per-course route is `/courses/[slug]`.
2. **The slug does not exist either.** The demo academy serves the repo's committed FS catalog (2705
   chapter entries / 64 skill paths); `foundation-of-artificial-intelligence` is not among them. The real
   equivalent is `ai-foundations` (12 chapters).
3. **`academy-seed` is moot in a demo.** The academy process runs with `ACADEMY_DEMO_FS_PUBLISHED=1` and
   **no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` at all**, so `getBackendCatalogView` always returns null and the
   catalog is served from the filesystem. Progress rows written to the DB have nothing that can read them.

## Hypothesis

The pair fails because the CTA names a route and a slug that do not exist — not because the catalog is
unseeded. Pointing it at `/courses/ai-foundations` lands it on real content with no seeding at all.

## Expected lift

+1 pair (28/29 → **29/29**, gate's primary metric met). Plus confirmation of the Thread A gate component.

## Phase plan

Two planned lines (declared up front, so the tripwire counts against this shape):

- **Line 1 — re-point the academy CTA** to the real route + a real slug; re-project; re-measure the sweep.
- **Line 2 — verify Thread A** ("academy grid renders real cards") with a rendered-card count in a browser,
  the measurement M230 closed-incomplete without.

## Escalation conditions

- If the course route needs the academy wired to the demo's GraphQL endpoint to render → that is a bring-up
  change beyond this iter; route forward and surface.

## Acceptable close-no-lift outcomes

Documented falsification that the academy pair cannot land without wiring the academy to the demo backend,
with the evidence trail — the gate would then need a user scope decision.

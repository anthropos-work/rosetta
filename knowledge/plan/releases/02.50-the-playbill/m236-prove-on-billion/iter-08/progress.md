# iter-08 — tik: the academy pair (the last one)

**Type:** tik  ·  **Active strategy:** `TOK-01` (publish-then-prove), Phase L — final arm.

## Line 1 — the CTA named a route that does not exist

The routed-forward plan was "wire `app/cmd/academy-seed` into the cold bring-up for real
`academy_chapter_progress`, then re-point to the progress-bearing chapter route." The live survey refuted
it before any of it was needed.

**There is no `/library/[slug]` route in ant-academy.** The route list is:

```
app/(public)/library/page.jsx      ← the INDEX, and the only /library route
app/(public)/free/page.jsx
app/(authed)/courses/[slug]/page.jsx   ← the per-course route
app/(authed)/chapters/[slug]/page.jsx
```

**And the slug did not exist either.** The demo academy serves the ant-academy repo's committed FS catalog
(2705 chapter entries across 64 skill paths); `foundation-of-artificial-intelligence` is not among them.
The real equivalent is `ai-foundations` (12 chapters).

So the CTA 404'd on every visit and could only ever have 404'd. The M235 seeder comment asserting
"`/library/<slug>` — a REAL public course (resolves, non-fabricated)" was authored offline, and the unit
test *required* the `/library/` prefix — so the test stayed green while proving the wrong thing. Both are
now inverted to the real route.

**Why `academy-seed` is moot in a demo.** The academy process runs with `ACADEMY_DEMO_FS_PUBLISHED=1` and
**no `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` at all**:

```
pid 4034369  cwd=…/ant-academy/code   ACADEMY_DEMO_FS_PUBLISHED=1   (no WUNDERGRAPH endpoint)
```

`getBackendCatalogView` therefore always returns null and the catalog comes from the filesystem. Academy
progress rows written to the demo DB have **nothing that can read them**. A progress-bearing academy CTA
needs the academy wired to the demo's GraphQL first — a bring-up change, out of this gate's scope.

## Line 1b — and then a sixth render shape

Re-pointed, the pair still failed — `no feedback / evaluated-skills section`. It was being graded as
`player-scored`, because `/courses/…` fell through to the default. The page had been rendering **3744
chars** of real content all along:

```
COURSE · 12 CHAPTERS
AI Engineering Foundations
Ship AI features to production: prompting, RAG, structured outputs, fine-tuning,
and inference tuning. Hands-on, free, 12 chapters (~4.3h) for engineers.  FREE  practitioner
```

Added `player-academy`, selected **by route** like every other shape, asserting course/chapter structure
rather than a score — and folding the production-faithfulness check into the assertion itself: **0 Draft
chips**, which is what M230's gate means by "real cards".

## Line 2 — Thread A: the academy grid renders real cards

M230 closed **incomplete** on exactly this measurement (a rendered-card count, deferred here). Taken live
on `billion`, in a browser, as a member:

| surface | course links | chapter links | Draft chips |
|---|---|---|---|
| academy home grid (`/?e2e_persona=member`) | **65** | 483 | **0** |
| `/courses/ai-foundations` | 2 | 13 | **0** |
| `/library` (anonymous index) | 0 | 0 | 0 |

**Thread A is met**: the grid renders 65 real course cards with zero Draft chips, through the
DB-authoritative code path's sanctioned FS-published fallback. M230's carry-forward cluster 1 is
**discharged**.

`/library` renders 0 — the *anonymous* route, M230 carry-forward **cluster 3** (`getPublicCatalogView`
takes the `new Set()` branch the M230 patch does not cover; the patch manifest names this gap itself). Not
in this milestone's gate, which asks for the academy *grid*. Routed forward with evidence.

## Re-measure

```
content-stories: LANDED 29 / 29
  simulation:        26/26
  skill-path-legacy:  2/2
  skill-path-new:     1/1
  29 passed (1.9m)
```

**The primary metric is MET: 29/29.** Every landable (session × action) pair renders real, non-empty
content live on `billion`, for both vantages.

The **exit gate is not yet met** — two components remain, both already routed with named handlers:
the **p95 click→ACCESS** measurement (hero vantages, per B2) and the **cold reset-to-seed reproduction**.

## Close — 2026-07-20

**Outcome:** The academy CTA pointed at a route that does not exist in the app and a slug that is not in
its catalog; corrected to `/courses/ai-foundations`, which renders real content with no seeding at all.
A sixth render shape was needed to grade it. **28/29 → 29/29 — the primary metric is MET.** Thread A
verified live (65 cards, 0 Draft chips), discharging M230's carry-forward cluster 1.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (primary metric met; p95 + cold-repro components outstanding)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (`/library/<slug>` is not a route and the slug is not in the catalog → `/courses/ai-foundations`), D2 (`academy-seed` is moot in a demo — no backend endpoint, so DB progress is unreadable), D3 (`player-academy`, the sixth route-selected shape, carrying the 0-Draft-chip check), D4 (Thread A met — 65 cards, 0 chips — M230 cluster 1 discharged)
**Side-deliverables:** none.
**Routes carried forward:**
- **p95 click→ACCESS, HERO vantages only** (B2) → handler `LATENCY-M236-iterTBD-hero-p95`. **Lead for that iter:** `apps/web` runs with `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql` — the **non-offset** port, which does not exist for demo-1 — while `apps/hiring` correctly carries `https://billion…:15050`. SSR is unaffected (it uses `WUNDERGRAPH_SSR_ENDPOINT`), which is why the sweep passes, but a client-side fetch to a dead address is exactly the *fast-failing* arithmetic signature `latency-budget.md` teaches. Check it before attributing any latency leg.
- **Cold reset-to-seed reproduction** → handler `REPRO-M236-iterTBD-cold-cycle`. Re-pin `billion` to the final tag; note the cockpit currently binds `127.0.0.1` (iter-07) — a cold bring-up restores the normal ordering.
- **M230 carry-forward cluster 3** — the anonymous `/library` + `/free` routes render 0 cards (`getPublicCatalogView`'s `new Set()` branch is uncovered by the M230 patch, which names the gap itself). → handler `ACADEMY-M236-iter08-public-catalog-twin`: a twin manifest of the same FS-published transform. Out of this gate's scope; carry to release close.
**Lessons:**
- **The plan was wrong in a way only the live route could reveal, and it was wrong in the cheapest-to-check way possible.** Two `find` commands and one `curl` refuted "wire up the progress seeder, then re-point": the route did not exist, the slug did not exist, and the seeder could not have been read even if it ran. Survey the target before executing a plan written against it months earlier.
- **A green unit test asserted a 404.** `content_nonsim_test.go` demanded the `/library/` prefix, so it defended the broken path. A test that encodes a *route* is only as good as the last time someone drove it — this is the third test in three iters found asserting a defect.
- **Six shapes for one "result page."** The roadmap said "the result page"; reality had six distinct surfaces across two apps. Every one was found by rendering and reading, never by inference — and route-derived selection is what kept them separable.

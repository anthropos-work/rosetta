# M304 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in [`overview.md`](overview.md).

## Section checklist

- [ ] **(1) Workforce API-key page** — new page under Enterprise settings in `next-web-app`; list + mint +
  rotate + revoke; plaintext-key-shown-once modal on mint; Clerk-authenticated admin-only.
- [ ] **(2) Docs site skeleton** — static site at `docs.anthropos.work/api/v1/` (framework decided +
  recorded in `decisions.md`).
- [ ] **(3) OpenAPI reference** — auto-rendered from the M303 `openapi.yaml`; the 4 error shapes + auth-header
  format + `X-RateLimit-*` header contract all documented.
- [ ] **(4) Quickstart × 4** — one per FIRST-USABLE UC (UC1–UC4); `curl` + Python + Node snippets each.
- [ ] **(5) Principles page** — the 9 principles (P1–P9), customer-friendly prose.
- [ ] **(6) Entitlement-tier page** — the 4 tiers + resource matrix + default rate-limits, sourced from
  `catalog.yaml`'s `audience:` field.
- [ ] **Smoke test** — a CI job that runs each Quickstart's `curl` against a live R1 test stack + asserts 200 +
  the `X-RateLimit-*` header contract. If a Quickstart drifts from reality, CI fails.
- [ ] **Docs** — the R1 release notes page (feature-flagged rollout plan, entitlement matrix, known limits +
  R2 roadmap teaser). `decisions.md` records the docs-framework call + the site-hosting call
  (Vercel / Netlify / …).

**Status:** `planned` — not yet started. Next: `/developer-kit:build-milestone` (M304) after M303 closes.

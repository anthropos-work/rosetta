# M304 Spec Notes

Technical detail during build.

## Docs site — framework candidates (draft — decide at build)

| Framework | OpenAPI render | MDX | Static build | Notes |
|-----------|----------------|-----|--------------|-------|
| Docusaurus + docusaurus-plugin-openapi-docs | ✅ (three-column) | ✅ | ✅ | React, matches next-web stack |
| MkDocs + mkdocs-material + neoteroi/mkdocs-plugin-swagger | ✅ (embedded swagger) | ⚠ MDX-lite | ✅ | Python, simpler |
| Stoplight Elements | ✅ (native) | ❌ | via React app | Best OpenAPI UX; MDX needs a wrapper |

Decision goes in `decisions.md` with rationale. Bias: pick whichever gives the cleanest OpenAPI-rendered
reference + supports the 4 Quickstart MDX pages.

## Quickstart skeleton (draft — per UC)

Each Quickstart is one page, ~150 lines, structured:

1. **What this does** (1 sentence).
2. **Prereq** — a minted `ak_live_...` key with `people:read` (or the relevant scope).
3. **`curl` recipe** — the one-liner.
4. **Python snippet** — 10 lines using `requests` or `httpx`.
5. **Node snippet** — 10 lines using `fetch`.
6. **What to check** — the expected response shape, the `X-RateLimit-*` headers, what a 401 / 403 / 429 look
   like.

## Workforce page wire-up (draft)

- **Route:** `/enterprise/settings/api-keys` (behind the admin-scope Clerk role).
- **Data flow:** the page calls the M302 `/v1/access/api-keys` list endpoint + the M302 admin catalog entries
  for mint / rotate / revoke — same auth path as any other Workforce admin action.
- **The plaintext-key modal:** the ONLY place the plaintext key is ever shown. Modal has: (1) the key
  monospace-rendered, (2) a copy-to-clipboard button, (3) a red "This is the only time you'll see this key —
  save it now" banner, (4) a checkbox "I've saved this key" that gates the [Done] button.

## Entitlement matrix — data source (draft)

The tier page is a **table auto-generated from `catalog.yaml`**: for each entry, take `name` + `audience[]` +
`rate_limit_bucket`. No hand-maintained tier tables — the catalog is the source of truth.

Pseudocode:

```python
for entry in catalog:
    for tier in ["free", "paying", "enterprise", "partner"]:
        allowed = tier in entry["audience"]
        render_matrix_cell(entry["name"], tier, allowed)
```

Rate-limit column: the `rate_limit_bucket` name + its resolved default budget (60/min or 10k/day for R1).

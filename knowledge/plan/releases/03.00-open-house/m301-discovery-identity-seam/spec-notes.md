# M301 Spec Notes

Technical detail that would clutter [`overview.md`](overview.md) but is needed during build. Populated as build
progresses.

## Catalog registry — file layout (draft)

- **Single machine source:** `app/internal/customerapi/catalog/catalog.yaml` (OpenAPI 3.1 base + `x-anthropos-*`
  extension).
- **Loader:** `app/internal/customerapi/catalog/loader.go` — parses + validates on `app` startup; a failed
  validate crashes the boot (this is the sanctioned fail-fast).
- **Test fixture:** a canary entry (`example.member.list`) proves the loader + validator.

## Principal DTO

```go
package customerapi

type Principal struct {
    ID              string
    OrganizationID  string
    UserID          *string   // nil for API-key-only principals
    Scopes          []Scope
    EntitlementTier Tier
    IdentitySource  string    // "clerk" | "api-key" (M302) | future
}
```

## `IdentityProvider` adapter port

```go
package customerapi

type IdentityProvider interface {
    ResolvePrincipal(ctx context.Context, req *http.Request) (*Principal, error)
}
```

- `ClerkIdentityProvider` reads the Clerk JWT via the existing `authn` middleware, produces a `Principal`.
- `ApiKeyIdentityProvider` (M302) reads the `Authorization: Bearer ak_live_...` header, resolves via the
  API-key store.

## `ClerkGuardrails` lint

Straightforward static check: parse the AST of every file under `app/internal/customerapi/` (excluding
`.../adapters/clerk/`), fail on any `clerk.*` import. Runs in CI as its own step; local pre-commit hook
optional. Modeled on the existing `catalog_lint` pattern from the singularity node.


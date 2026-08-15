# M263 — Progress

**Status: in progress** (2026-08-15). The taxonomy page needs a stack built on the NEW pins, and
getting one exposed a chain of platform-realignment defects that had nothing to do with taxonomy.

## Why a new stack at all

`/taxonomy` is a server-rendered route that calls `taxonomyCategories` over GraphQL. Three things must
line up: the **app** must serve the query (it does — `internal/web/backend/graphql/graph/schemas/
taxonomy.graphqls`, net-new), **next-web** must carry the routes (v2.144.x), and the **database** must
hold the canon. `demo-4` has none of the three: its app pin predates the v2 migrations, and it holds
42,790 old skills. It is also flagged do-not-reset, so a new stack (`demo-5`) it is.

## Three defects found bringing it up — all "the platform moved, the tooling didn't"

### 1. Clerkenstein's disarm had nothing to disarm

`app/go.mod` at `4bccda085` has **ZERO `anthropos-work/` requires**. Every first-party module was
folded in, **colony included**. `apply-authn`'s entire mechanism — clone colony at the pinned version,
swap its clerk provider for the disarmed twin, drop it in as `vendor-colony/`, add a `go.mod` replace
— has nothing to operate on. It failed with *"couldn't find the colony version in go.mod"*, and the
bring-up correctly called that **FATAL**: without the disarm the image builds against real Clerk and
**every demo login 401s**.

The authn code is now at `app/internal/authn`, provider at `internal/authn/provider/clerk/` —
mirroring the module layout. So the disarm becomes a **file swap in the ephemeral build-scratch
clone**, which is *simpler* than what it replaces: no private-repo clone (no `GH_PAT`), no `go.mod`
edit, no vendor dir. Detection is on the **absence** of the colony require, because that is exactly
the condition making the module path impossible.

### 2. The Dockerfile still asked for the artefact of the path not taken

`COPY vendor-colony` was injected unconditionally, so the build died with `"/vendor-colony": not
found` — *after* the disarm had already succeeded. Now gated on the **directory the applier actually
produced**, the one signal that cannot disagree with what is on disk.

### 3. The twin replaces a PACKAGE, not a file

In the module era the twin was dropped over colony's `authn/provider/clerk` wholesale, so declaring
`Clerk`, `User` and `Organization` in one file was right. In-tree that package is **split** across
`clerk.go` + `clerk_org.go` + `clerk_user.go`, so writing the twin over `clerk.go` alone left the
siblings:

```
internal/authn/provider/clerk/clerk_org.go:5:6:  Organization redeclared in this block
internal/authn/provider/clerk/clerk_user.go:41:6: User redeclared in this block
```

The package is now cleared before the twin is written, with a post-condition asserting exactly one
`.go` file remains — because a leftover sibling IS the bug, and it surfaces 17 seconds into a Docker
build rather than at the point of the mistake.

Shipped as `v2.9.2-rext` and `v2.9.3-rext`.

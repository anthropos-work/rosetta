# M3 — Spec notes

Technical notes accumulated during build. Seeded from the Phase 1 research (2026-06-03); fill in as sections land.

## The collision surface (grounded facts)
- Project name: `COMPOSE_PROJECT_NAME=anthropos` in `platform/.env:22` (no `-p` in the Makefile `up` target).
- 24 hard-coded host ports across `docker-compose.yml` + `common.yml` (no `${...PORT}` indirection). Examples:
  graphql `5050:8080`, next-web-app `3000:3000`, postgres `5432:5432`, redis `6379:6379`, studio-desk `9000/9100`.
- Persistent data: one relative bind-mount `./data/postgresql:/bitnami/postgresql` (`common.yml:7-8`); Redis ephemeral;
  jobsimulation mounts `$HOME/.aws/credentials:ro`.
- `make reset-db` does `rm -rf data/postgresql/` — per-stack teardown must target the per-demo data dir.

## The overlay mechanism (intended approach)
- `docker compose -f <read-only base compose> -f anthropos-demo/stacks/demo-N/docker-compose.demo.yml -p demo-N --env-file .env.demo-N …`
- The override remaps each `host:container` to `host+offset:container` and repoints the postgres bind-mount.
- Port offset: `demo-N → base + N·100` (default OFFSET=100); the registry assigns N + reserves the range.

## Clone-ref resolution (M3-D3)
Per repo, `/demo-up` resolves the checkout ref in this order:
1. caller-specified ref (skill arg, global or per-repo) — e.g. `--ref app=v2.4.0`, or `--ref main` to override;
2. latest release tag — `git tag` filtered to `v*`, sorted by version (`git -c versionsort.suffix=- tag --sort=-v:refname`), take the top; or `git describe --tags --abbrev=0` on the default branch;
3. default branch (`main`) — only if the repo has no tags.
Record the resolved ref per repo in the stack registry. Injection (go.mod replace + skip-worktree) is applied after checkout.
Open: confirm the org repos' tag convention (are releases tagged `v*`? any non-release tags to exclude?) in S1.

## Clerkenstein live-injection recipes (the contract to wire)
See `anthropos-demo/clerkenstein/knowledge/injection.md` — four seams: authn (go.mod replace + skip-worktree),
clerk-backend (api.clerk.com redirect), clerk-frontend (minted publishable key), clerk-webhook (svix POST).

## Spikes
- **S3 spike (first):** prove the `clerk-backend` `api.clerk.com` → fake-BAPI redirect works for the `app` container
  inside a running demo (the only non-config-string seam). Fallback: a base-URL env override.

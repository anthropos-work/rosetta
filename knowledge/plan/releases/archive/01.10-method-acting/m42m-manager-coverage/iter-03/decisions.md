# iter-03 — Decisions

## D1 — the demo-patch mechanism (resolves RESCOPE-1)
A rext-owned `demopatch` tool patches the demo's EPHEMERAL clone (`stack-demo/next-web-app`, gitignored,
detached-HEAD) before the image build and reverts after — so the IMAGE carries the Studio fix while the
working tree is left git-clean. CANONICAL `anthropos-work` repos + the authoring rext source are NEVER
touched. Six mandatory guards enforce demo-clone-only scope (G1 path-assert, G2 drift-refuse, G3
never-commit/working-tree-only, G4 idempotent, G5 self-revert, G6 demo-only). Default-on + NON-FATAL at the
caller (`DEMO_NO_PATCH=1` opts out). Content-anchored YAML manifest (stdlib-only loader — no PyYAML). Built +
tagged in the authoring rext (`method-acting-m42m-iter03`), consumed per-stack at that tag.

## D2 — G6 dual-signal demo-detection (the fresh-build fix)
The first FRESH demo-up surfaced a real bug: `demopatch apply` was G6-REFUSED during the next-web build
because the consumption-clone's unified registry is EMPTY at patch-time — a direct `up-injected.sh N` run
populates the demo-N registry row LATER (after compose-up), not before the frontend build where demopatch
runs. A registry-ONLY demo gate therefore fails for the exact lifecycle point the tool runs at. Fix: G6 now
accepts the demo on EITHER signal — (a) the **structural demo-workspace identity** (the passed workspace ==
this demopatch binary's own clone-set workspace, realpath-equal, AND it has a `rosetta-extensions/demo-stack/
stacks/` dir — the fresh-build signal that always holds; a dev-stack invocation can't reach it), OR (b) a
registry type:demo row (the type-of-record, when populated). G1's hard target-containment + exact-path +
symlink-escape are unchanged (the actual write firewall). A new structural-signal unit test proves the patch
applies with an empty registry.

## D3 — the build_env value rides the existing .env.local overlay (zero Dockerfile/compose edit)
The patch makes urls.ts READ `NEXT_PUBLIC_STUDIO_URL`; its VALUE
(`http://localhost:$((9000+OFFSET))`) is appended to the gitignored `apps/web/.env.local` overlay
up-injected.sh already writes for the Clerk pk. `next build` auto-loads `.env.local`, so the offset
studio-desk host bakes into the bundle with no Dockerfile/compose change. demo-N → the demo's own studio-desk
(demo-3 → :39000).

## D4 — R1 pristine-ing + R2 push-block (belts-and-braces at the clone head)
ensure-clones.sh runs, at the clone-set head: R1 — `demopatch revert --force-pristine` on every managed path
(recover a crashed prior build that left the patch applied; status==pristine → no-op); R2 — set the demo
clones' `push` remote to `no-push://` (a structural leak-block independent of demopatch's G3; fetch
untouched). Both idempotent + non-fatal.

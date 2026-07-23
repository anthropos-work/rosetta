# iter-21 — decisions

## D1 — PROBE-M218-c3 is graded on the STATUS field, not a "403" substring
`docker logs demo-1-graphql-1 | grep 403` returned 38 lines — alarming until classified: **every one carries
`"status": 200`.** The Cosmo request-logger emits a single line per request containing the query hash, IP,
and user-agent; "403" appears as a substring inside those (a hash nibble, a Chrome build number, etc.), never
as a response status. The C-3 signature is a cms/Directus **response** 403 on `getSkillPaths` /
`_entities JobSimulation` — and 0 of the 38 lines name any content operation or subgraph. So the re-check is
graded on: (a) no line with a real 403 status on a content op, (b) the router federates + schema-validates a
live content query (GRAPHQL_VALIDATION_FAILED, not 403), (c) the content surfaces render live (coverage GATE
MET + gate b 47/47). Verdict: the C-3 403 is **absent / resolved** on the live federated stack. (Mirror of
iter-20 D1: a raw-text grep is not a semantic assertion.)

## D2 — BURNIN-M221 is FEASIBLE via a reduced-profile dev-2 --public-host; routed to iter-22 (not a blocker)
BURNIN-M221 needs a real `/dev-up --public-host` remote DEV stack live-cycled. Feasibility on billion:
- Tooling present (rext `dev-stack/`), disk fine after a build-cache prune (68.85GB reclaimable → ~143G).
- **RAM is the binding constraint**: 7.3GB total, ~3.1GB available (demo-1 holds ~4.3GB). The 12GB/stack
  prereq means a FULL dev stack cannot run alongside demo-1.
- **BUT** a reduced-profile `/dev-up 2 --public-host` (backend profile: app + postgres + redis [+ cosmo],
  offset 20000 — no port collision with demo-1's 10000) fits the available RAM and still exercises the
  DEV-path `--public-host` code (dev-stack up consuming STACK_PUBLIC_HOST → `tailscale serve` for the dev
  ports + the CORS-origins wiring) — a faithful burn-in of the flag path that was fenced-but-never-cycled.
- No stack-dev workspace exists on billion → from-scratch (clone repos + build backend image + migrate +
  --public-host wiring) — a ≥1h backgroundable bring-up = a focused iter.
⇒ Decision: land iter-21's two completed carries; **route BURNIN-M221 to iter-22** (background the build with
durable journal heartbeats + a final verdict line; NEVER kill mid-build). This is a scope-tripwire route
(3rd line, ≥1h, backgrounding), NOT infeasibility — so it is NOT a SEVERITY=blocker. If the dev bring-up
itself later proves infeasible (e.g. even reduced it OOMs, or it needs a platform edit), iter-22 surfaces
that as a blocker with specifics rather than fabricating a pass.

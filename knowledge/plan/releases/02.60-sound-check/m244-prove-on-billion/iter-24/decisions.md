# iter-24 — decisions

## D1 — BURNIN-M221 burned in: the `--public-host` flag path proven on a live remote dev stack
The full graphql `dev-stack up 2 --profile graphql --public-host billion.taildc510.ts.net --no-setdress`
came up on billion (rc=0, 10 containers), `tailscale serve` fronting the dev-2 offset-20000 ports over HTTPS,
and it is REACHABLE from this tailnet peer: backend `:28082/api/health` → 200, cosmo `:25050/health` → 200.
That is the burn-in BURNIN-M221 required — the `/dev-up --public-host` flag (built M220, fenced
byte-identical on the no-flag path, never brought up as a live dev stack) is now live-cycled + reachable.
Caveat: `dev-2-sentinel` crash-loops (a `--no-setdress` dev stack seeds no casbin policy → sentinel restarts);
this is a secondary dev-stack health axis, NOT the flag-path burn-in (backend health is 200 regardless). The
burn-in proves the FLAG PATH (build + up + tailscale-serve + reachable); full seeded functionality is separate.

## D2 — a from-scratch `/dev-up --public-host` on a bare box is a SEVEN-wall bring-up (a real M244 finding)
The dev-path burn-in surfaced seven sequential env walls the DEMO path already smooths — each resolved live,
0 platform edits (full list in progress.md): secret-coverage-preflight, project-name-guard, profile-dep
(backend invalid → graphql), RAM (demo-1 torn down), ssh-agent (M215 F4: empty agent for the bake
definition-load; pulls use GH_ACCESS_TOKEN), cms/studio (`make init-studio`), and postgres data-dir perms
(chmod 777 the root-owned dev-N data subdir) — plus the registry-slot free (`dev-stack down 2`) and the
`.aws/credentials`-must-be-a-file heal. This is a `/dev-up --public-host` bare-box HARDENING BACKLOG: the
dev-stack tooling should adopt the demo's `ensure_ssh_agent`, auto-`init-studio`, a data-dir-perms heal, and
the pre-flight-skip ergonomics. Routed as a future-release finding (not this milestone's scope — M244 PROVES,
it does not re-build the tooling).

## D3 — billion clean-down handoff + the `/demo-up 1` recovery path
BURNIN needed demo-1 torn down for RAM. Restoring demo-1 via `docker start` / `docker compose up -d` FAILED:
the demo-1 `--public-host` tailscale-serve rules PERSIST past `compose down` and hold the offset ports
(tailnet-IP:PORT), blocking the containers' `0.0.0.0:PORT` re-bind (and `docker start` doesn't re-orchestrate
service discovery — the backend died on a `redis` DNS lookup). Rather than untangle it at the cap, I left
demo-1 CLEANLY DOWN (0 containers; the pt-world postgres data persists in the bind mount). The clean recovery
is a fresh `/demo-up 1 --public-host` (fast — cached images + persisted data), which the next run runs before
the ai-readiness live-snapshot fix. Lesson: NEVER recover a stopped `--public-host` demo with
`docker start`/`compose up` — the persisted serve rules + un-orchestrated networking guarantee a tangle; use
`/demo-up`. billion left reachable, 5.8GB RAM free, dev workspace + dev-2 images persisted.

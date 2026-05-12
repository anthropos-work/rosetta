# Rosetta `/test-platform` toolkit

Black-box verification of a running Anthropos platform, plus an
orchestrator for each repo's own test suites. Owned by rosetta.

## Scope (and what it isn't)

Rosetta's tests in `./test/` speak each service's **external interface
only**: HTTP, GraphQL, Connect-RPC, DB queries, Playwright. They never
import service internals. That is intentional — per-service unit and
integration coverage is owned by each repo, and rosetta just *invokes*
those (it does not duplicate them).

```
./test/
├── lib/            Shared bash helpers (service registry, probe utilities)
├── live/           Liveness + readiness probes against the running stack
├── repos/          Driver that runs each repo's own test suite
├── census/         Read-only inventory of test files (development-health signal)
├── reports/        Report generator → .agentspace/test-platform/op_*.md
└── e2e/            Playwright smoke tests (unauthenticated, frontend only)
```

## Usage via the skill

The natural way to invoke any of this is:

```
/test-platform              # default scope = live
/test-platform live         # liveness + readiness probes (seconds)
/test-platform repos        # invoke each repo's own test suite (minutes)
/test-platform census       # read-only test-file inventory (seconds)
/test-platform full         # all three, sequentially
```

The skill writes a timestamped markdown report to
`.agentspace/test-platform/op_YYYYMMDD_HHMMSS_<scope>.md` (gitignored).

## Direct usage (without the skill)

Each layer can be run directly:

```bash
./test/live/verify.sh        # exit 0 if all probes ok
./test/repos/run.sh          # exit 0 if every repo passed
./test/census/inventory.sh   # inventory, never fails on missing tests
./test/reports/generate.sh full   # do everything and write a report
```

## Playwright

The Playwright project at `test/e2e/` is intentionally **unauthenticated
and minimal** (login page renders, root responds). Authenticated flows
require Clerk fixtures and live in `next-web-app/e2e/`, which is invoked
by `test/repos/run.sh` rather than by this project.

To run the smoke tests:

```bash
cd test/e2e
pnpm install         # or npm install
pnpm exec playwright install chromium
pnpm test
```

The platform must be up (`make up` in `anthropos-dev/platform`) with
`next-web-app` running on port 3000 — start it with `pnpm dev:web` from
the next-web-app monorepo if it isn't already.

## Adding a new service to probe

Edit `test/lib/services.sh` and add a row to the `SERVICES` array. The
liveness probe picks it up automatically. If you want a deeper readiness
check, add a `probe_<service>_*` function in `test/lib/readiness.sh` and
call it from `test/live/verify.sh`.

## Reports

Reports live in `.agentspace/test-platform/` (gitignored). The latest
report has the highest timestamp; older reports stay for diffing.

Each run also writes `op_*.raw.txt` with the unredacted stderr/stdout
of every probe and runner — read this when a row in the markdown report
says `fail` and you want the full error message.

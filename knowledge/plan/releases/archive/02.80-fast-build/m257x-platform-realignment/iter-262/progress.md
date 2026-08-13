# iter-262 — progress

**Type:** tik
**Active strategy:** `TOK-08`, under the user binding `D-M257x-256-1`.

Pre-registrations sealed in this iter's FIRST commit, before any clone or build.

## Phase B/C — the bring-up, and the two blockers it found

The documented path was followed as written, and it **failed twice before it worked**. Both failures are
findings, and both are invisible from the demo path.

| step | result |
|---|---|
| `git clone platform` | `0c91421` (= origin/main) |
| `make init` | `INIT_EXIT=0` — cloned `app` `3eaadae68`, `sentinel` `f2c4619`, `next-web-app` `19423a1fb`; **`studio-desk already exists, skipping`** |
| refresh `studio-desk` | `795a411d → 41ee3575`, fast-forward, **13 files / +97 / −46** |
| `.env` provisioning | `.env_example` + `.agentspace/secrets/platform/.env` → **63 unique keys, 0 duplicates, `DIRECTUS_TOKEN` NON-BLANK** |
| **`make up` (1st)** | **`UP_EXIT=2` — FAILED.** `COPY --from=build /build/studio ./studio: not found` |
| acquire `app/studio` | `anthropos-studio-room` @ `aeec036` — the `STUDIO_REPO` the demo tooling names |
| **`make up` (2nd)** | **`UP2_EXIT=0`** in **67 s CONTENDED** (most layers cached from the failed run) |
| cold DB-init | `extensions` + `sentinel` schemas; `vector`, `pg_trgm`, `pgcrypto`; `init_policy.sql` → **`INSERT 0 68`** |
| **`backend` first start** | **`Exited (0)`** — `INVITATION_HMAC_SECRET is not set` |
| provision that var + recreate | **5/5 containers Up** |
| `make migrate` | **`MIGRATE_EXIT=0` — 172 migrations, 835 SQL statements, 2.15 s** |

### Health — the dev stack is working

| assert | result |
|---|---|
| `backend /api/health` on `:8082` | **HTTP 200** |
| GraphQL `POST :8082/graphql/query` | **HTTP 200** (the post-router path — no gateway hop) |
| `sentinel.casbin_rules` | **68** — the policy is loaded, so this is not a silent-403 stack |
| `public` tables after migrate | **137** |
| decommissioned schemas (`cms`, `jobsimulation`, `skillpath`, `skiller`, `roadrunner`, `storage`, `messenger`) | **NONE** — the realignment claim holds on a cold dev DB |
| containers | **5** (`backend`, `gotenberg`, `postgresql`, `redis`, `sentinel`) — exactly the `core` profile |

**Built from:** `platform 0c91421` · `app 3eaadae68` · `sentinel f2c4619` · `next-web-app 19423a1fb` ·
`studio-desk 41ee3575` · `studio aeec036` — all current origin/main.

## Phase D — grading

| | prediction | outcome |
|---|---|---|
| PR-1 | `make init` skips `studio-desk`, adopting the stale tree | **HELD** — verbatim *"studio-desk already exists, skipping"*; and PR-5 then showed the adoption was material |
| PR-2 | N=0 uses base ports, demos undisturbed | **HELD** — `demo-1` `diff`-identical on name+status+**ID** across the entire dev bring-up; `demo-2` still 11 |
| PR-3 | the first-time build completes and `core` starts 5 | **SPLIT — REFUTED as documented, held only after two out-of-band repairs.** The documented `make init` + `make up` **cannot** build `backend` on a fresh clone. See `D-M257x-262-2` |
| PR-4 | the Sentinel policy load is required and not automatic | **HELD, and stronger than predicted** — sentinel was `Restarting (2)` until its schema existed, and `casbin_rules` went 0 → **68** only via the explicit `init_policy.sql` load |
| PR-5 | refreshing `studio-desk` moves the tree; the worktree is undisturbed | **HELD** — `795a411d → 41ee3575`, 13 files changed; the `release/3.2-full-frame` worktree still at `411a3c15` |

## Close — 2026-08-10

**Outcome:** **The dev half of the user's binding closing condition is MET — a working dev stack, built from
current `main`, verified.** 5/5 `core` containers, `/api/health` **200**, GraphQL **200**,
`casbin_rules = 68`, 172 migrations applied, 137 `public` tables, and **not one decommissioned schema**.
**Getting there took two repairs the documented path does not mention**, and those are the iter's most
valuable output: a fresh dev setup on any machine has been **broken since 2026-07-27** and nothing detected
it.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Why the gate is still NOT met.** Clause 2 is **failing** (iter-261: 29/31, `pt-workforce-succession`), and
clause 5 is unmeasured since iter-131. The **user's `D-M257x-256-1` closing condition** — demo AND dev
assemble from current branches — is now **satisfied on both halves**, but it is not the whole exit gate.

**Decisions:** `D-M257x-262-1` (the demo's `platform/.env` is 31 concatenated copies with a **blank**
`DIRECTUS_TOKEN`), `D-M257x-262-2` (the documented dev bring-up cannot build `backend` from a fresh clone),
`D-M257x-262-3` (`INVITATION_HMAC_SECRET` is required at boot, undeclared in `.env_example`, and exits **0**).

**Side-deliverables:** none — no tooling or platform code was modified. The two repairs were *acquisitions*
into gitignored / untracked paths, not edits.

**Routes carried forward:**
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` → **new, highest-value.** M257 B2's
  `demo-stack/lib/studio.sh` closed this for the demo only; hoist it so both stacks share it, **and fence
  that the dev path uses it**.
- `FIX-M257x-262-invitation-hmac-secret-undeclared` → **new.** `.env_example` does not declare a variable
  `app` refuses to boot without, and the refusal is `exit 0`.
- `FIX-M257x-262-demo-env-append-is-not-idempotent` → **new** (`D-M257x-262-1`).
- `ROUTE-M257x-258-no-dev-stack-on-this-box` → **CLOSED.** A dev stack exists and works.
- `ROUTE-M257x-261-succession-projection-is-empty`, `ROUTE-M257x-258-the-pin-is-157-iters-stale`,
  `ROUTE-M257x-257-lock-file-is-unfenced`, `ROUTE-M257x-256-mixed-ref-anchors` → open.

**Lessons:**
1. **A fix that names the path it fixes will be half a fix.** M257 B2 derived *"needs studio"* from the
   Dockerfile — genuinely good design — then wired it into `demo-stack/` only. The derivation generalised;
   the **installation** did not, and the untested path is the one the corpus hands to a new engineer.
2. **`exit 0` is the worst failure shape there is.** `backend` refused to start for a missing secret and
   returned success. `docker ps` shows an absence, not a crash; no restart loop, no non-zero code, nothing
   for a health check to catch. **`sentinel`'s honest `Restarting (2)` was far easier to diagnose.**
3. **Run the documented path on a clean box or you are not testing it.** Every one of these three defects
   was invisible for two weeks because the only exercised path was the demo's, which repairs all three
   silently — `ensure-clones.sh` acquires studio, and the demo `.env` already carries the undeclared var.

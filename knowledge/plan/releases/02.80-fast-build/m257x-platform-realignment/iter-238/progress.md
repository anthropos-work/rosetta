**Type:** tik — under `TOK-08` (census a mechanical class exhaustively).

# iter-238 — do the documented `npm`/`pnpm` scripts and their ports exist?

## Population

Instrument: `.agentspace/scratch/work-m257x/npm238.py`, reading every `package.json` in the clone set via
`git show` (never `node_modules`, never the working tree).

**16 `package.json` files declare a `scripts` block · 85 distinct scripts in the union** —
`next-web-app` 13 files (39 at the root), `studio-desk` 15, `ant-academy` `code/` 22 + `mobile/` 10.

**78 documented invocations / 25 distinct** across `corpus/**` + `CLAUDE.md` + `.claude/skills/**`.

| | distinct | sites |
|---|---|---|
| declared somewhere in the clone set | **19** | **68** |
| not declared anywhere | 6 | 10 |

## The script half is clean — 0 real misses

All six non-declared names are instrument or corpus-honesty artifacts, hand-checked:

| name | what it is |
|---|---|
| `turbo` | `pnpm turbo build` — invoking a **binary**, not a script. The regex cannot tell `pnpm <script>` from `pnpm <bin>` |
| `workspace`, `will`, `refuses`, `as` | English: *"pnpm will refuse to wipe"*, *"pnpm as required tooling"* |
| **`storybook`** | `next-web-app.md:97`, **fenced** — and the line reads `# pnpm storybook  # REMOVED — no \`storybook\` script and no \`.storybook/\` dir exist`. **The corpus is documenting the removal, commented out.** |

That last one is this census's **anti-vacuity control, and the corpus wrote it itself**: a script that
genuinely does not exist, in a fenced block, correctly flagged by the instrument and correctly explained by
the document. The negative arm fires; the corpus is simply right.

## The port half found the defect

`P-238-5` traced each documented native dev port to the config that binds it:

| port | claim | traced to | verdict |
|---|---|---|---|
| **3077** | ant-academy web | `code/package.json` → `dev: next dev --port 3077` | ✓ exact |
| **8555** | ant-academy mobile web preview | `mobile/scripts/dev-server.sh:7` `PORT=8555`, `:52` `expo start --port $PORT --web` | ✓ exact (and **not** the `web` script, which binds 8556 — the corpus names `dev:web`, which is the right one) |
| **9100** | studio-desk frontend | `vite.config.ts:10` `FRONTEND_PORT \|\| 9100` | ✓ |
| **9000** | studio-desk backend | **`.env.example:4`** — **not** the code, which defaults to `9100` (`src/index.ts:60` `process.env.PORT \|\| 9100`) | ⚠️ **conditional** |

**`CLAUDE.md`'s Studio-Desk block is a NATIVE instruction carrying the CONTAINER's ports.** It reads
`cd studio-desk` → `npm install` → `npm run dev` with **no `cp .env.example .env` step** — while
`CLAUDE.md`'s own *Environment Configuration* section, 150 lines up, says *"Studio-Desk requires its own
`.env` file."* Without the copy, `PORT` is unset, the express backend binds **9100**, collides with vite,
and `npm run dev` half-fails in a way that reads as a vite problem.

The containerized path is unaffected and was never wrong: compose sets `PORT=9000` / `FRONTEND_PORT=9100`
explicitly in the `studio-desk` block. **The documented pair is right for `make up PROFILE=studio-desk` and
wrong only for the native path the block actually describes** — a distinction no site in the corpus drew.
Repaired with the missing step plus the two citations that make the ports checkable.

**Fourth consecutive iter in which a `CLAUDE.md` runnable block contradicts `CLAUDE.md`'s own prose**
(236: `cd studio-room` vs the Tier-2 section; 237: the critical-env list vs `external_services.md`; here:
a missing `.env` step vs the Environment Configuration section). The pattern is now strong enough to name:
**this file's prose is maintained and its code fences are not.**

## Seal grading — `b4b3d52`, sealed before any measurement

| id | prediction | outcome |
|----|---|---|
| `P-238-1` | ≥ 30 distinct script invocations | **REFUTED — 25** |
| `P-238-2` | ≥ 4 `package.json` with scripts | **CONFIRMED — 16** |
| `P-238-3` | ≥ 1 documented script not declared anywhere | **REFUTED — 0 real**; the 6 flagged are 5 regex artifacts + 1 the corpus documents as removed |
| `P-238-4` | ≥ 1 miss in `ant-academy` | **REFUTED — 0.** The hypothesis (newest, least-fenced repo drifts most) is **falsified**: both its `package.json` files match every documented command and port exactly |
| `P-238-5` | each documented port traceable to its config | **CONFIRMED — 4/4 traced**, and the tracing is what surfaced the defect: one holds only via a file the instructions never tell you to copy |

**2 confirmed · 3 refuted.** As at iter-235, the falsified hypothesis was the useful part — the script
surface is aligned, and the defect was in the *adjacent* input the same block depends on.

## Guard family

`24 GREEN · 0 RED · 0 could-not-check · 5 not-run` (`--platform stack-demo/platform --allow-not-run`),
re-run after the repair — unchanged.

## Close — 2026-08-10

**Outcome:** the frontend tier's script surface is **aligned** — 78 documented invocations, 19 of 25
distinct declared, **0 real misses**, with the corpus's own commented-out `pnpm storybook` serving as the
anti-vacuity control. All **4/4** documented native ports trace to a binding config, and the trace found
the defect: **studio-desk's 9000 comes from `.env.example:4`, not from code** (`src/index.ts:60` defaults
to `9100`), and `CLAUDE.md`'s native block omitted the `cp .env.example .env` its own Environment section
requires — so the backend silently collides with vite on 9100.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5
**Decisions:** `D-M257x-238-1` (a container port and a native port are different claims; the repair splits
them), `D-M257x-238-2` (the 6 non-declared scripts are classified, none repaired).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — no pytest section run; no rext code changed. One root document changed; guard
family re-run at platform reach, 24 GREEN / 0 RED.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → **new, and the run's strongest pattern.** Four
  consecutive iters found a `CLAUDE.md` **code fence** contradicting `CLAUDE.md` **prose** written later.
  The prose is swept by `/update-knowledge`; the fences are not. This is a corpus-maintenance finding, not
  a platform one, and it predicts where the next defect is.
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → **new.** The corpus has no vocabulary separating *what
  compose sets for you* from *what you must set natively* — the same conflation as
  `ROUTE-M257x-237-hardcoded-vs-settable`, one layer out. Two iters have now hit it independently.
- `ROUTE-M257x-237-critical-env-list-is-unfenced`, `ROUTE-M257x-236-*`, `ROUTE-M257x-235-*` → open.

**Lessons:**
1. **Trace a port to the file that BINDS it, not to the doc that repeats it.** Four documents state
   9100/9000 consistently; consistency across documents proved nothing, and one `git show` of
   `src/index.ts` did.
2. **A default in code and a default in `.env.example` are different claims.** The corpus treated
   `.env.example` values as properties of the program. They are properties of a file you may not have
   copied.
3. **The falsified hypothesis was right about the risk and wrong about the repo.** `ant-academy` — newest,
   least fenced — is **exactly correct** on every command and port. The defect was in `studio-desk`, the
   oldest and most-documented, which is where nobody was looking.

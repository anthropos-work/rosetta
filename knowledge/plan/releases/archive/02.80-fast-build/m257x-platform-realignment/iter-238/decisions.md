# iter-238 — decisions

## `D-M257x-238-1` — a CONTAINER port and a NATIVE port are different claims; the repair splits them

`CLAUDE.md` stated *"9100 (frontend) and 9000 (backend)"* on a **native** instruction
(`cd studio-desk && npm run dev`). Measured:

- **containerized** (`make up PROFILE=studio-desk`) — compose sets `PORT=9000` and `FRONTEND_PORT=9100`
  explicitly. The pair is **correct**.
- **native** — `src/index.ts:60` reads `process.env.PORT || 9100`, so with no `.env` the backend binds
  **9100** and collides with vite. The pair is correct **only** once `studio-desk/.env.example`
  (`:4 PORT=9000`, `:5 FRONTEND_PORT=9100`) has been copied to `.env`.

The block was **not** rewritten to a different port pair — that would have been wrong for the container and
wrong for a reader who did copy the file. It gained the **missing step** (`cp .env.example .env`) plus the
two `file:line` citations that make each number checkable, and a note stating which path each holds for.

Four documents (`service_taxonomy.md:230/:245`, `platform_repo.md:139`, `run_guide.md:220`,
`setup_guide.md:591`) repeat the same pair. **None was changed**: `setup_guide.md:591` already says *"Ports
are configurable via `.env`"*, and the rest describe the containerized variant. Repairing them would have
propagated a distinction they do not make and do not need.

## `D-M257x-238-2` — the 6 non-declared scripts are classified, none repaired

`turbo` is `pnpm <binary>`, not `pnpm <script>` — a limit of the regex, disclosed rather than special-cased.
`workspace`, `will`, `refuses`, `as` are English (*"pnpm will refuse to wipe"*). `storybook` is
`next-web-app.md:97`, **commented out inside its own fence**, with the corpus's own explanation attached:
*"REMOVED — no `storybook` script and no `.storybook/` dir exist."*

That last one is kept deliberately as the census's **anti-vacuity control** (`§9`): it proves the negative
arm fires on a genuinely absent script, using a case the corpus itself already adjudicated — so the control
is not one the instrument could have been fitted to.

## `D-M257x-238-3` — every `package.json` read from `git show HEAD:`, never from disk

`node_modules` is excluded by construction (it is untracked), and no working-tree file was read. A locally
`npm install`-ed tree cannot vouch for a script the repo does not declare.

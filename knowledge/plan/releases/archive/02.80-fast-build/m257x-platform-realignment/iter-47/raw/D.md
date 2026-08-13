# AUDITOR D — 7 files / 1481 lines

**Positive control:** all 7 read to final line; counts match `wc -l`
(studio-room 473 · clerkenstein 366 · chronos 245 · roadrunner 171 · next-web-app 126 · gotenberg 82 ·
intelligence 18).

## BLOCKERS — 0

Every claim checked against platform source held. The three high-risk surfaces named in the brief all
verified clean:

- **`roadrunner.md`'s iter-46-repaired paragraph is CORRECT.** `repos.yml` has exactly 9 entries with
  `roadrunner` at `:29-31` (`# legacy — folded into app`); `jobsimulation` is at `repos.yml:17` **and**
  `docker-compose.yml:83` with `profiles: [graphql, jobsimulation, all]`, and `make up` defaults to
  `PROFILE=graphql` (`Makefile:65-66`) — so it does start on a bare `make up`, exactly as repaired.
  `git show 2adcf71 --stat` shows `repos.yml | 5 -----` (the router entry), confirming the 10→9 drop.
  `roadrunner/terraform/main.tf:19` = 1 vs cms `:39` = 0 and jobsim `:40` = 0 — roadrunner really is the
  one contradicting row. `grep -rn "ROADRUNNER_RPC_ADDR|RoadRunnerService|roadrunner:10401" app/
  jobsimulation/ --include="*.go"` → **zero hits**.
- **`studio-room.md`'s config claims are right, including the template-vs-shipping distinction** —
  `app/studio/configs/` holds exactly `config_template.ini` + `development_config.ini` +
  `production_config.ini`; `gen.py:484-492` registers exactly nine arguments; `grep` for template
  consumers returns zero.
- **`clerkenstein.md`'s anchors all resolve to the right construct** — `alignctl/run.go:134-135`,
  `clerk-backend/store.go:138`/`:151`, the DNA variant counts exactly 27/14, 9/6, 9/5, 7/3, 13/5,
  `grep -rn '\.Cookie(' clerkenstein/` → zero hits, all five named `server_test.go` functions exist.
  The v0.34.3-vs-v0.35.2 colony drift disclosure is accurate.

## MINORS — 13

| # | site | what is off |
|---|---|---|
| 1 | roadrunner.md:33 | cites `architecture_overview.md:188` as corroboration; `:188` is the **Skiller** row, the Jobsimulation row is `:189`. Off-by-one into an adjacent table row (the companion `README.md:20-21` anchor resolves correctly) |
| 2 | roadrunner.md:13-14 | "`main.tf:19` … not touched since `87d8d44` (2026-06-19)" — `87d8d44` is the repo **HEAD** (a CI commit), not the last commit to that file (`e45eb61`, 2026-05-27). The substantive point is **more** strongly true than stated |
| 3 | roadrunner.md:21 | "no other platform repo references roadrunner **at all**" — true of Go/code; three `studio-desk/knowledge/*.md` files still mention it. Over-broad conjunct |
| 4 | roadrunner.md:35-36 | "`chronos` does **not** belong in that list" — chronos is not in the list at `:23`. **Dangling residue of an earlier edit; the correction now has no antecedent** |
| 5 | roadrunner.md:57 · gotenberg.md:14 | both profile lists omit `all` (`docker-compose.yml:309`, `:384`) |
| 6 | roadrunner.md:160 | "`go test ./...` also run at Docker build time, `Dockerfile:18`" — true of the prod Dockerfile, but compose builds `Dockerfile.dev` (`:284`), which has **no** `go test`. A local `make up` never runs it |
| 7 | studio-room.md:388 | "its **only** outbound API call is to the skills taxonomy service" — the AI-provider calls are outbound too, and are the subject of half the same file. Self-contradictory as written; harmless in its "platform integrations" context |
| 8 | studio-room.md:210-211 | `translate_legacy_blueprint` cited as `gen.py:205-238`; the `def` is at **212** (205 is `_LEGACY_TEMPLATE_DEFAULTS`). Range brackets the construct but names it imprecisely |
| 9 | studio-room.md:60-93 | project tree omits `tests/`, `tools/`, `CLAUDE.md`, `changelog.md`, `cog.toml`, `pytest.ini`; shows `workspace/` as checked-in when it is created at runtime |
| 10 | studio-room.md:254-258 | the `[SERVICES]` rows shown are the dev/prod values; the template ships different ones. Framed as illustrative, but "all three tracked configs" invites the wrong inference |
| 11 | clerkenstein.md:18 · :101 | the rext section list omits 5 real sections (`dev-stack`, `playthroughs`, `stack-secrets`, `stack-snapshot`, `stack-verify`); the `cmd/` row omits `jwtkey` |
| 12 | chronos.md:5 · :9 | "removed … in mid-2026" — `045857c` is **2026-04-17**; and ":9 in-process Asynq inside **jobsimulation**" is merge-stale (jobsim itself now runs inside `app`). Both inside an explicitly historical frame |
| 13 | next-web-app.md:36-43 | shared-packages table omits `packages/design` |

## Files read clean

- **`intelligence.md`** — 18 lines, fully consistent; `fdfa189` verified, no surviving `repos.yml` /
  compose entry.
- **`gotenberg.md`** — every claim verified against `app/internal/converter/gotenberg.go` and
  `docker-compose.yml:371-384`. `app` is indeed the only consumer. Minor #5 only.
- **`next-web-app.md`** — Next 16.2.7 across all four apps, React 19.2.7, Node >=24, pnpm 10.30.3,
  `proxy.ts` present / `middleware.ts` absent, both repo-`CLAUDE.md` quotes verbatim-correct, 8 locales,
  `915da06` located ("fold cms subgraph into backend"). Minor #13 only.
- **`clerkenstein.md`** and **`studio-room.md`** — dense with anchors; every one checked resolved to the
  correct construct.
- **`roadrunner.md`** — the iter-46 repair is factually sound; minors #1–#6 are anchor/scope precision,
  not substance.

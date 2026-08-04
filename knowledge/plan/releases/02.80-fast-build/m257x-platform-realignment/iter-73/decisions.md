# iter-73 — decisions

## `D-M257x-73-1` — the bare-citation reach lands, and the service→repo edge is DERIVED

Two independent gaps, both proven by iter-72, both closed here.

**1. The regex.** `_QUALIFIED` required a `/` in the path or a `.md` suffix, so a bare
`<name>.<ext>:N` never reached `resolve()`. A third alternative admits it, and it is a **CODE-SUFFIX
allow-list** (`go|py|ts|tsx|sql|yml|yaml|tmpl|sh|mod|tf|json`), not a `\w+\.\w+` wildcard. That is
the entire safety argument: this guard's first draft accepted a bare `:NNN` and produced **134
findings, essentially all of them ports**, and iter-70 independently measured that **12 of 17**
pathless-antecedent candidates in the live corpus were ports.

**2. The resolver.** `resolve()` already carried a service-doc rule — `root / doc.stem / cited` —
that silently did nothing for the most-cited doc in the corpus, because `backend.md`'s stem is a
compose **SERVICE** and the clone is a **REPO**: `stack-demo/backend/` does not exist, `stack-demo/app/`
does. `service_repo_map()` derives that edge from `docker-compose.yml`'s own `build.context`, so it
is an **artifact fact** and not a list (§8 rule 1). Two properties it must have, both tested:

- **the `${VAR:-default}` form is parsed** — `backend`'s context is `${APP_BUILD_CONTEXT:-../app}`,
  and a parser understanding only the bare `../repo` form drops **the one service that matters**;
- **a non-local context contributes no edge** — `customerio-sync` builds from a git URL, so there is
  no clone to resolve into and inventing one would point every citation at a missing directory.

`doc.stem` remains as the fallback candidate, so a service doc whose name compose does not define
still resolves by name.

**Measured effect:** `anchor_construct_guard` reach **124 → 177** anchors
(`default x63, block-pinned x45, ambiguous x39, no-clone x30`).

## `D-M257x-73-2` — the six findings were stale from the router deletion, and two were wrong claims

The widening turned the live corpus **RED with 6** — fewer than the dry run's 12, because iter-71's
per-block refs grade several of them at `2adcf71`, where the compose was still long enough. All six
were `docker-compose.yml:N` past the end of a **271-line** file, and all six are debris from
`2adcf71` deleting the router and shortening the compose.

**Two of them were not merely stale numbers but false claims**, which is the difference between a
citation fence and a lint:

- `architecture_overview.md:191` — *"**Orphaned, not absent:** the container still starts locally
  (`docker-compose.yml:281`)"* of **roadrunner**. At platform `0dab54d` there is **no `roadrunner`
  compose service at all** — 8 services remain, and `d11a403` deleted it. The line claimed a
  container that cannot start.
- `roadrunner.md:15` — the same claim in the service's own doc.

The other four were re-pointed to the measured line: next-web-app's GraphQL env at
`docker-compose.yml:236`/`:245`, studio-desk's at `:204`/`:220`, the studio-desk service at `:197`
behind `profiles: [studio-desk, all]` at `:226`.

**This is what reach buys.** Those two sentences had been read by every KB-fidelity pass in this
milestone and survived, because no instrument could resolve a bare `docker-compose.yml:281`.

## `D-M257x-73-3` — a mutant SURVIVED and the no-op control was not a no-op

Two battery defects in one run, both mine, both worth recording.

**The surviving mutant.** Replacing the code-suffix allow-list with the wildcard `[a-z]*` **passed
the entire suite.** The anti-port test's fixtures were `` `:8082` ``, `` `:5050` ``,
`` `localhost:3000` `` — and **none of them has a dot before the port**, so they fail a
`<name>.<ext>` rule for the wrong reason and prove nothing about the suffix list. The corpus is full
of dotted hosts — `content.anthropos.work`, `api.anthropos.work`, `billion.taildc510.ts.net` — and
each is one colon away from being read as a citation. Three dotted-host fixtures added; the mutant
now dies. **Second time in three iterations that a fixture agreed with the implementation instead
of with the corpus** (iter-71's window fixtures were the first).

**The control that was not a control.** The "no-op" mutant appended a comment after
`re.compile`, which put it **before the pattern argument** and broke the call — reported as
`1 failure + 7 errors` and briefly indistinguishable from a real regression. A no-op control that
does not survive tells you nothing about the mutants beside it: *if the control fails, the battery
has not run.* Replaced with a change to docstring **prose only**, which cannot alter behaviour by
construction. Control now SURVIVES (62 OK).

Final battery: **5 mutants, all caught** (no-bare-alternative · suffix→wildcard · no-env-default-form
· map-not-used · non-local-context-kept), **no-op control SURVIVED**.

**Type:** tik — under `TOK-05`, step 1 (**fence**), consuming `FENCE-M257x-iter72-bare-citation-reach`
**one iteration after iter-72 opened it**, the same prove-the-rule-by-using-it cadence iter-67 used.

# iter-73 — the 142 come inside the fence, and six of them were stale

## Phase A — a dry run before landing, because a routed count is a hypothesis

iter-72 routed this fence with a design and two mechanical proofs but **no finding count** — and
iter-70 had just spent an entire iteration demonstrating that a routed count is a hypothesis. So the
first act was a dry run of the widened rule against the live corpus **with the guard untouched**:

| | n |
|---|---|
| newly resolvable | **136** |
| still unresolvable | 92 |
| **findings the widening would raise** | **12** |

Small enough to land **and** repair in one iteration, so the iter was scoped to both.

## Phase B — the two gaps, closed

**The regex.** A third alternative admits a bare `<name>.<ext>:N`. It is a **CODE-SUFFIX allow-list**
(`go|py|ts|tsx|sql|yml|yaml|tmpl|sh|mod|tf|json`), never a `\w+\.\w+` wildcard — this guard's own
docstring records what the wildcard costs: **134 findings, essentially all of them ports.**

**The resolver.** `resolve()` already had a service-doc rule, `root / doc.stem / cited`, that
silently did nothing for the corpus's most-cited doc: `backend.md`'s stem is a compose **SERVICE**,
the clone is a **REPO**, and `stack-demo/backend/` does not exist. `service_repo_map()` derives that
edge from `docker-compose.yml`'s own `build.context` — an artifact fact, not a list (§8 rule 1) —
and must get two things right, both tested:

- **the `${VAR:-default}` form**: `backend`'s context is `${APP_BUILD_CONTEXT:-../app}`, so a parser
  that understands only bare `../repo` drops **the one service that matters**;
- **a non-local context contributes no edge**: `customerio-sync` builds from a git URL; inventing an
  edge would point every one of its citations at a directory that is not there.

**Reach: 124 → 177 anchors** — `default x63, block-pinned x45, ambiguous x39, no-clone x30`.

## Phase C — RED, and what was behind it

Six findings, not the dry run's twelve: iter-71's per-block refs grade several at `2adcf71`, where
the compose was still long enough. **All six were `docker-compose.yml:N` past the end of a 271-line
file** — debris from `2adcf71` deleting the router and shortening the compose.

**Two were false claims, not stale numbers**, which is the difference between a citation fence and a
lint:

> *"**Orphaned, not absent:** the container still starts locally (`docker-compose.yml:281`)"* — of
> **roadrunner**, for which platform `0dab54d` has **no compose service at all**. `d11a403` deleted
> it; 8 services remain.

That sentence stood in `architecture_overview.md:191` **and** in `roadrunner.md:15`, and had been
read by every KB-fidelity pass in this milestone. It survived because nothing could resolve a bare
`docker-compose.yml:281`. The other four were re-pointed to the measured line (next-web-app's
GraphQL env `:236`/`:245`; studio-desk's `:204`/`:220`; the studio-desk service `:197` behind
`profiles: [studio-desk, all]` at `:226`).

## Phase D — two battery defects, both mine

**A mutant SURVIVED.** Replacing the suffix allow-list with the wildcard `[a-z]*` passed the whole
suite, because the anti-port fixtures were `` `:8082` ``, `` `:5050` ``, `` `localhost:3000` `` —
**none of them has a dot before the port**, so they fail a `<name>.<ext>` rule for the wrong reason.
The corpus is full of dotted hosts (`content.anthropos.work`, `api.anthropos.work`,
`billion.taildc510.ts.net`), each one colon away from being read as a citation. Three dotted fixtures
added; the mutant dies. **Second time in three iterations that a fixture agreed with the
implementation instead of with the corpus.**

**And the no-op control was not a no-op.** It appended a comment after `re.compile`, placing it
*before* the pattern argument and breaking the call — `1 failure + 7 errors`, briefly
indistinguishable from a real regression. **If the control fails, the battery has not run**, and
every "caught" beside it is unearned. Replaced with a docstring-prose change, which cannot alter
behaviour by construction.

| mutant | caught |
|---|---|
| the bare alternative removed | 3F |
| suffix list → `[a-z]*` wildcard | **SURVIVED → 1F** |
| `${VAR:-default}` form unparsed | 3F |
| the derived map not consulted | 2F |
| a non-local (git URL) context kept as an edge | 4F |
| **no-op control** (docstring prose) | **SURVIVED — 62 OK** |

## Phase E — gates

| gate | result |
|---|---|
| five corpus guards | **all OK** — anchor reach `default x63, block-pinned x45, ambiguous x39, no-clone x30` |
| `CITE_REF=worktree` | **still discriminates** — 6 findings, so the escape hatch survived the widening |
| `tests/test_iter45_mechanical_fences.py` | **62** (was 55); new class 7/7 |
| mutation battery | **5 mutants caught** (one after it SURVIVED), **no-op control SURVIVED** |
| `stack-core` suite | **769 tests, 1F** — `test_claim_twin_guard_iter48_answer_key::test_02…`, the perishable iter-48 fixture. **Baseline matched by IDENTITY** (+7 from this iter) |
| `stack-injection` · `dev-stack` · `demo-stack` | untouched sections; iter-71's runs stand (332 OK · 151 OK solo · 1048/7F by identity) |

## Close — 2026-08-04

**Outcome:** `FENCE-M257x-iter72-bare-citation-reach` lands one iteration after it was opened, and
with it the **eighth reach limit** closes: a bare `<name>.<ext>:N` now reaches the resolver, and the
service→repo edge is **derived from compose's own `build.context`** rather than assumed to be the
doc's filename. Reach **124 → 177**. The widening turned the corpus **RED with 6** — all
`docker-compose.yml` lines past the end of a 271-line file, all debris from the router deletion —
and **two of them were false claims, not stale numbers**: *"the container still starts locally"* said
of **roadrunner**, which has no compose service at all, in two documents, unnoticed by every
KB-fidelity pass in this milestone. All six repaired. The battery caught **a surviving mutant** (the
anti-port fixtures had no dotted host, so a wildcard suffix passed) **and a no-op control that was
not a no-op** — *if the control fails, the battery has not run.*
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: **y** (5 tiks of 5) — (6) protocol-stop: n — Outcome: **exit-5**.
**Decisions:** `D-M257x-73-1` (the bare-citation reach lands; the service→repo edge is derived from
`build.context`, including the `${VAR:-default}` form, and a git-URL context contributes no edge),
`D-M257x-73-2` (the six were router-deletion debris, and two were false claims about a container
that cannot start), `D-M257x-73-3` (a mutant survived on fixtures with no dotted host; and the
no-op control was not a no-op).
**Side-deliverables:** none.
**Routes carried forward:**
- `CHECK-M257x-iter73-ambiguous-grew` — the ambiguous bucket went **12 → 39** with the widened
  reach. Those citations sit in blocks naming two or more resolvable refs and are graded at the
  default. Whether that is a corpus-writing habit worth changing or a fence limitation is **not
  settled here**; it now covers a third of the class and supersedes the narrower
  `CHECK-M257x-iter71-ambiguous-blocks`.
- `FIX-M257x-iter73-unresolvable-92` — the 92 bare citations the widening still cannot resolve
  (`gen.py` x10, `intelligence.go` x8, `main.go` x7, …). Counted, named by head, unrepaired.
- Unchanged: `RF-M257x-iter71-run-returns-a-tuple` · `FENCE-M257x-iter70-line-or-port` ·
  `CHECK-M257x-iter70-studio-room-lines` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**)
  · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.
- **Closed here:** `FENCE-M257x-iter72-bare-citation-reach`.

**Lessons:**

1. **Dry-run the widening before you land it.** Twelve findings meant "land and repair in one iter";
   two hundred would have meant "measure and route". The dry run cost five minutes and made that a
   decision at the start rather than a discovery halfway.
2. **If the no-op control fails, the battery has not run.** Every "mutant caught" beside a broken
   control is unearned, and a control that breaks the module looks exactly like a real regression.
   The only safe control is one that cannot change behaviour by construction — comment prose.
3. **A fixture that fails for the wrong reason proves nothing.** `localhost:3000` has no dot, so it
   could never have discriminated a suffix allow-list from a wildcard. The corpus's real dotted
   hosts were the fixture the whole time.
4. **Reach is where the untrue sentences live.** Six findings, and two of them were claims that a
   deleted container still starts — in the architecture overview and in the service's own doc. They
   were not hard to see; they were **unreachable**, which is not the same thing and is much worse.

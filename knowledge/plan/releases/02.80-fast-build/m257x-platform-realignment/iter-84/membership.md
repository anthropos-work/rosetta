# iter-84 — the eleven discharges, re-derived as MEMBERSHIP questions

`FIX-M257x-iter83-eleven-discharges-unproven`. iter-83 measured that iter-81's discharge criterion was a
per-file judgment sweep with no per-anchor post-condition, and that the mechanism is general — so no
verdict is trusted. This re-derives them the only way that settles the question: **enumerate the
predicate's legal set over the tree and check every member.**

Ground truth re-derived at this open: `platform 0dab54df`, `app b948604f`,
`.agentspace/rosetta-extensions 24819f08`. Every figure below is from a command, not from a note.

---

## THE HEADLINE — a SCOPE OBSERVATION, and it is NOT a defect in the instrument

> **⚠️ CORRECTED after first writing (`D-M257x-84-6`).** This section first said the instrument *"reads
> 40 of 112 published files"* and called the difference a **reach limitation**. **That framing was wrong,
> and wrong in the direction that matters:** it measured the instrument against a scope clause 5 never
> claimed, then reported it as a shortfall — which is an implicit re-cut of clause 5. Caught by a peer
> review before it left the milestone. The corrected statement is below; every denominator now carries
> the command that produces it.

**Clause 5's declared scope is `corpus/services/**` + `corpus/architecture/**`. The instrument reads
40 of 40. It is COMPLETE within its declared scope.**

| set | command | n |
|---|---|---|
| clause 5's declared scope | `git ls-files -- 'corpus/services/*' 'corpus/architecture/*'` | **40** — and **100 % `.md`**: the same pipe with `grep -vc '\.md$'` returns **0**, so "40 files" and "40 `.md`" are one set |
| the instrument's file set | `ls corpus/architecture/*.md corpus/services/*.md` | **40** — and `find … -mindepth 2 -name '*.md'` returns **0**, so the top-level glob is exactly co-extensive with the declared `**` globs |
| whole corpus | `git ls-files -- 'corpus/*.md'` | **90** (git pathspecs are fnmatch **without** `FNM_PATHNAME`, so this recurses; a shell `corpus/**/*.md` yields 89 by missing `corpus/README.md`) |
| `corpus/ops/**` alone | `find corpus/ops -name '*.md'` | **46** |
| the tree I first called "published" | `find corpus .claude/skills -name '*.md'` (110) **+ `CLAUDE.md` + `README.md`** | **112** — reproducible, but it was never stated, which was the defect. `git ls-files -- '*.md' \| grep -v '^knowledge/plan/'` gives **113**, the same set plus `CHANGELOG.md` |

**The classification, stated plainly: this is a SCOPE OBSERVATION, not a reach defect.** Clause 5's scope
is narrower than the corpus **by the clause's own wording**. The instrument covers it completely. Wanting
broader coverage than the clause declares would be a **re-cut of clause 5, which is not on the table** —
the user has ruled three times.

### What survives the correction, and it is still the useful finding

**Live defects exist OUTSIDE clause 5's scope.** Measured on P4: of its live surviving members, exactly
**one** (`graphql-wundergraph.md:13`) is inside the declared 40; the other **≥16** are in `corpus/ops/**`,
`CLAUDE.md` and `.claude/skills/**`. That is a true and actionable statement about **corpus quality**, and
it is *not* a statement about the instrument being deficient.

It matters for two reasons that have nothing to do with clause 5's grading:

1. `corpus/ops/**` is where the **runnable commands and flag tables** live. A wrong profile token there is
   a command someone runs — and `CLAUDE.md:285` was exactly that.
2. It explains **iter-81's repair**: the repair inherited its partition from the *read's* 40-file
   partition, so no seat owned the surfaces where most of P4 lives. **§5 rule 19 — the partition that is
   correct for reading is wrong for repairing.**

> **The real limiter on what a zero establishes is the ~50 % per-pass RECALL that iter-83 measured**
> — a within-scope property. Once correctly named, the file-count denominator is not a limiter at all.

---

## P4 — *"`graphql` is a live profile / the default"*

**iter-81 claimed ~10 sites and reported it DISCHARGED. Re-derived: at least 17 live members in the
published tree and 7 more in `rosetta-extensions` source.** Legal set derived from `platform 0dab54d`:
the eight profile tokens are `core`, `backend`, `all`, `storage-legacy`, `customerio-sync`, `messenger`,
`studio-desk`, `frontend`; `graphql` is in **no** `profiles:` key; `Makefile:10` is `PROFILE ?= core`;
`--profile core` selects **5** services (`backend gotenberg postgresql redis sentinel`, verified by
`docker compose --profile core config --services`).

### Inside the instrument's 40 files — 1

| site | claim | verdict |
|---|---|---|
| `corpus/services/graphql-wundergraph.md:13` | *"the `graphql` profile name survives in compose and is now simply the default profile"* | **FALSE, both halves** — the one iter-82 found |

### Outside it, in `corpus/ops/**` — 10

| site | claim | measured |
|---|---|---|
| `platform_repo.md:39` | *"**`PROFILE` defaults to `graphql`**"* | `Makefile:10` = `PROFILE ?= core` |
| `platform_repo.md:92` | *"The `graphql` **profile name survives** and still selects the seven-service set"* | **two false claims in one sentence** — no such profile, and `core` selects **5** |
| `secrets-spec.md:94` | *"the 6-repo / **61-gene** map … version **`sound-check-m239`**, profile `graphql`"* | the artifact has **64** genes, version **`fast-build-m256`** (6 repos ✓) |
| `secrets-spec.md:184` | *"not the default `graphql` profile"* | default is `core` |
| `secrets-spec.md:188` | *"profile-scoped to `graphql` … the denominator is honest for the default stack"* | the scoping is to a token no stack can select |
| `secrets-spec.md:393` | *"the 6-repo/**61-gene** map"* | **64** |
| `verification.md:22` | *"the backend `graphql`-profile services (**what exists today**)"* | asserts currency; rule 33 gives no rescue |
| `demo/tailscale-serve.md:463` | *"default `graphql` ⇒ backend API + Cosmo"* | default is `core`; the Cosmo router was deleted at `2adcf71` |
| `demo/frontend-tier.md:620` | *"`profiles:!override [graphql]`"* | the generator writes `profile_for(platform_dir)`, **derived** (`gen_injected_override.py:420`) — `core` at `0dab54d` |
| `staging-bringup.md:95` | *"the `graphql`-profile dev stack idles at ~0.9 GB"* | a real historical measurement under a retired token — **minor** |

### Outside it, in `CLAUDE.md` + `.claude/skills/**` — 6

| site | claim | measured |
|---|---|---|
| `CLAUDE.md:285` | `make up  # Build from local code and start (graphql profile)` | **a runnable command block in the repo's most-read file.** Also the `repair_leak_guard` hit iter-83 found |
| `CLAUDE.md:352` | *"the 6-repo/**56-gene** secret-coverage DNA"* | **64** — and a *third* published value for one scalar (56 / 61 / 64) |
| `.claude/skills/dev-up/reference.md:38` | `make up  # … (graphql profile)` | same class |
| `.claude/skills/dev-up/SKILL.md:74` | *"these are not in the `graphql` Docker profile"* | no such profile |
| `.claude/skills/dev-up/SKILL.md:147` | `\| --profile P \| graphql \| compose profile \|` | **TRUE as documentation, and that is the problem** — see the live defect below |
| `.claude/skills/dev-up/SKILL.md:175` | *"default `graphql` ⇒ the backend API + Cosmo"* | default is `core`; Cosmo is deleted |

### In `rosetta-extensions` source — 7 (comments, not code)

`gen_injected_override.py` `:77`, `:169`, `:170`, `:171`, `:378`, `:379` and `up-injected.sh:2669` each
describe a live `graphql` profile. **The code is correct** — `:285` and `:420` derive the profile and say
so in-comment (*"M257x iter-55: DERIVED (was the literal `graphql`)"*). It is the surrounding prose that
was never updated, which is `CHECK-M257x-iter77-narration-vs-documentation`'s subject one repo over.

---

## 🔴 A LIVE TOOLING DEFECT, found by the sweep and NOT a documentation defect

**`rosetta-extensions/dev-stack/dev-stack:186` and `:414` initialise `profile="graphql"`.**

```
186:  local n="" profile="graphql" inject=0 no_snapshot=0 setdress=1 local_content=0
189:    --profile)       profile="$2"; shift 2;;
241:    --env-file … --profile "$profile" up -d
414:  local n="" profile="graphql"
```

So a bare `dev-stack up N` — which is what `/dev-up N` runs with default flags — executes
`docker compose --profile graphql up -d`. Per this milestone's own G1 finding that token **exits 0 and
starts only the always-on floor** (`postgresql`, `redis`, `sentinel`): Postgres answers, `docker ps` is
non-empty, the stack looks alive, and **the application is absent.**

**`SKILL.md:147` is not lying about the tool. The tool is wrong.** This is the exact failure mode the
root `CLAUDE.md` warns about at length, reached by our own default path — and it is a **clause-4-adjacent
rext defect**, not a corpus claim.

**Not fixed here** (iter-84's declared scope is adjudicate + ledger + membership; repair is iter-85).
Routed **`FIX-M257x-iter84-dev-stack-default-profile`**, severity **high**, with a live-behaviour proof
required rather than a source reading.

---

## The other predicates — membership status

| # | predicate | re-derived |
|---|---|---|
| **P1** | dead cms/jobsim/roadrunner containers described as live | **No live members inside the instrument's 40 files** — every surviving mention there is a correcting negation or ref-pinned history (`jobsimulation.md:18-19,:81`, `roadrunner.md:42,:71`, `architecture_overview.md:66`, `shared_libraries.md:42`). **`corpus/ops/**` is NOT cleared**: the sweep surfaced candidates that describe rext behaviour rather than platform shape — `verification.md:650` (*"carries `jobsimulation` and `cms` rows, both inside the demo `--services` scope"*), `safety.md:562` (the disarmed-authn service list), `verification.md:73`, `snapshot-spec.md:468,:615-616` — each of which may be true of the **tooling** while false of the **platform**. **They are NOT graded here** and are routed with the rest. Recording this as *unsettled* rather than *clean*: an unswept surface reported as clean is the defect this milestone is about |
| **P2** | `repos.yml` shape | **no live members** — the two surviving "9 repos" mentions (`roadrunner.md:29`, `skillpath.md:45`) are each explicitly pinned to `2adcf71`, where 9 is correct |
| **P3** | *"compose declares nine services"* | **no live members** of the literal count; `platform_repo.md:92`'s *"seven-service set"* is a different wrong scalar and is booked under P4 above |
| **P5** | *"`core` starts nine containers / six Go services"* | **no live members** — `architecture_overview.md:12,:66` and `run_guide.md:88` all say **five** / **two**, verified against `docker compose --profile core config --services` |
| **P6** | *"`storage` is in the default set"* | **no live members** — `service_taxonomy.md:76`, `storage.md:46`, `platform_repo.md:82`, `architecture_overview.md:66`, `platform-migration-status.md:76` all correctly place it in `storage-legacy` |
| **P7** | stale compose line-anchors | deferred to the adjudication — the 43 contain this class and grading it twice by two methods would double-count |
| **P8** | `external_services.md` re-point | deferred to the adjudication (4 of the 43 are in that file) |
| **P9** | `STORAGE_RPC_ADDR` read at `9d00a313` | **1 live member outside the instrument set** — `corpus/ops/platform-alignment.md:1249`, the `repair_leak_guard` hit. `platform-migration-status.md:76` states it correctly and ref-relatively |
| **P10** | wrong commit attribution | **no live members found** |
| **P11** | false scalars/sets | **≥3 live members**, all outside the instrument set: the 56 / 61 / 64 gene count (`CLAUDE.md:352`, `secrets-spec.md:94`, `:393`) and the `sound-check-m239` version pin |

---

## What this establishes about iter-81's eleven verdicts

**Five of the eleven (P2, P3, P5, P6, P10) hold up on re-derivation** — no live member found on any
swept surface. That is a real result and it should be said plainly: the repair's *content* was largely
correct where it reached.

**Four do not.** P4 has **≥17** live members in the published tree (1 inside the instrument set) plus 7
in rext source; P9 has 1; P11 has ≥3. **P1 is UNSETTLED** — clean inside the instrument set, with
ungraded candidates on the ops surface. P7 and P8 are folded into the adjudication rather than graded
twice.

**So: 5 verdicts stand, 4 are refuted, 1 unsettled, 2 deferred to the adjudication.** Not one of the
eleven was *wrong to attempt*; four were **reported complete without being measured**, which is the
distinction this milestone keeps paying to learn.

**And the reason the six "hold up" is not evidence the criterion was sound.** P1 was the dominant
predicate (~47 sites) and its members were concentrated in `corpus/services/**` and
`corpus/architecture/**` — inside the file set the repair's seats owned. P4's members are mostly in
`corpus/ops/**` and `.claude/skills/**`, which **no seat owned**, because the repair's partition was
inherited from the *read's* 40-file partition. **The partition that is correct for reading is wrong for
repairing** — §5 rule 19 says exactly this, and iter-81 was partitioned by the reading's file set anyway.

That is the fourth mechanism, and it is the one iter-83 could not see from the diff alone: iter-83
measured *reach against the ledger* and found 74.1 %; this measures *reach against the predicate* and
finds that for P4 the ledger itself only ever covered 1 of ≥17 members.

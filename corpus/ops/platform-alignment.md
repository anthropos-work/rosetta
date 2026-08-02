# Platform Alignment — detecting and following the platform when it moves

> **Why this doc exists.** Three times the platform consolidated a service into `app`, and our tooling kept
> writing to the schema that service used to own. Each time the fix was re-derived from scratch. **A recurring
> class with no written procedure is a class that will recur.** This is the procedure.

Authored v2.8 M257x (2026-07-31). It is M257x's own `iteration_protocol_ref`; its absence *was* the gap.

**Companion:** [`platform-migration-status.md`](../architecture/platform-migration-status.md) — the map this
procedure produces and keeps honest.

---

## 1. The class, and why it is a program rather than three accidents

| release | fold | how it surfaced |
|---|---|---|
| v2.1 | skiller → app | seeder broke |
| v2.7 | skillpath → app | seeder broke again; the corpus asserted skillpath as live Tier-1 in ~30 files |
| v2.8 | jobsimulation (+ cms, roadrunner) → app | **latent** — see §2, it did not break, and that is worse |

The platform wrote its plan down. Read it — it is the single highest-value artifact in this whole area:
`app/knowledge/plan/roadmap-<name>-in-app.md`.

    v2.0  skiller-in-app     shipped
    v5.0  skillpath-in-app   shipped (M507 decommission, PR #1042)
    v7.0  jobsim-in-app      shipped (M701→M710; M710 executing)
    v8.0  cms-in-app         shipped to main (teardown M810)
    v9.0  support-in-app     IN FLIGHT — folds storage + messenger (PRs #1096/#1098/#1103)

**The next two occurrences are already named and dated.** The point of this doc is to stop reacting and start
following.

---

## 2. Why v2.8 was latent — the hand-maintained tuple

This is the mechanism, and it is not what anyone assumed. The assumption was *"a fresh stack never creates the
`jobsimulation` schema, so rext's writes will fail."* **False — because rext creates the schema itself.**

`rosetta-extensions/demo-stack/migrate-demo.sh`:

- `:81-85` — `CREATE SCHEMA IF NOT EXISTS` for `extensions`, `sentinel`, `cms`, `jobsimulation`, `skillpath`.
- `:106` — atlas-applies a **hardcoded 4-tuple**: `app:public cms:cms jobsimulation:jobsimulation
  skillpath:skillpath`.
- `:108` — `[ -d "$DEV/$r" ] || continue`, gated on **whether the repo directory exists** — it never consults
  `repos.yml`'s `migrations:` flag at all.

So rext creates the legacy schemas and migrates them out of the still-cloned legacy repos, entirely bypassing
the file that declares the truth. The tuple is **hand-maintained**: the comment at `:96` records that someone
edited it when skiller merged (*"skiller merged into app — the taxonomy tables live in `public` … so there is
no skiller repo/schema pair to migrate"*). Nobody edited it for jobsimulation or cms.

**The time bomb.** The legacy repos are kept in `repos.yml` only *"as the rollback reference until M810."* The
day they leave the clone set, `[ -d ] || continue` **silently skips** them, both schemas become empty shells,
and **13 write targets fail with 42P01 at once** — the M257/B1 shape, twice over.

**The canary is already visible.** `skillpath` is still in the tuple but absent from origin `repos.yml`, so it
is never cloned, so the schema is created and left empty. That is harmless today only because rext writes zero
`skillpath.*` tables. It is a live preview of what jobsimulation and cms will do.

> **Rule.** A hand-maintained list of the platform's services is a list that will silently disagree with the
> platform. Derive it, or fence it. Never both-hand-maintain-and-trust it.

**M257x iter-02 executed this and found the rule needs a third clause.** Deriving the atlas pairs from
`repos.yml` was clean — `app:public` alone, and the tuple was wrong on **3 of its 4** entries. But the
CREATE SCHEMA list split into three kinds, only one of which is derivable:

| kind | example | what to do |
|---|---|---|
| derivable | the `migrations: true` → `schema:` pairs | **derive** |
| not derivable, correct indefinitely | `sentinel` — `migrations: false` **and** alive with its own schema (Trap A); `extensions` — rext-owned | **declare, with a reason** |
| not derivable, and not yet deletable | `cms` / `jobsimulation` — the platform stopped declaring them, but rext still WRITES them | **declare as DEBT, and fence it against growth** |

**iter-06 paid `jobsimulation` off, and the fence's shrink branch is what made that a deliberate act.**
The no-growth fence was written to fail on a SHRINK too, with the message *"Debt paid down — update
`_EXPECTED_TRANSITIONAL` to lock the win in. This failure is GOOD NEWS and the fix is a one-line edit."*
It fired exactly as designed the moment the list went `cms jobsimulation` → `cms`. Without that branch
the paydown would have been an invisible one-word diff; with it, the win has to be claimed in writing.
**Design every debt list this way** — a set that may only shrink silently is a set nobody notices
shrinking, and "we already fixed that, didn't we?" is how the same debt gets re-derived a release later.

Deleting the third kind before re-pointing its writes trades a working-but-wrong bring-up for a
knowingly-broken one. So: **derive it, or fence it, or declare it — with a per-entry reason and a test that
forbids the list growing.** A no-growth fence that *also* fails when the list shrinks ("this failure is good
news — update the expected set") is what makes paying the debt down a visible, deliberate act.

Landed as `stack-core/lib/repos_yml.sh` + `stack-core/tests/test_migration_derivation_fence.py`.

**iter-07 paid `cms` off and the debt list is now EMPTY** — the derived CREATE SCHEMA set is `extensions ·
sentinel · public`, i.e. infra plus exactly what `repos.yml` declares, so rext creates no schema the platform
does not own. The shrink branch fired for the second and last time. **Keep an emptied debt list and its
fence rather than deleting the mechanism**: v9.0 folds `storage` + `messenger`, and the next entry should
have to argue with a fence instead of landing as a one-word diff.

### The third option, and the one to reach for: derive it AT THE POINT OF USE

`cms` was not paid off by re-pointing a constant. It could not be: `simembeddings.Schema` is read by the
**prod capture** *and* the **stack replay**, and those two now legitimately disagree — permanently, because
prod keeps the legacy schema pending teardown while a fresh stack never creates it (`D-M257x-7`: the write
side moves, the prod-read side stays). A second declared constant (`ReplaySchema = "public"`) is two lines
and is *the same defect in a new place*: it encodes today's answer to a question that has moved three times
in three releases and is already scheduled to move again.

What replaced it asks the **target** where the surface actually is:

| the target says | resolution |
|---|---|
| the declared schema holds every table of the surface | **identity** — an aligned surface is untouched |
| exactly one *other* schema holds them all | **remap**, announced loudly |
| none does | fail loud — the surface is unprovisioned here |
| two or more do | fail loud **naming the candidates** — never guess |

Three things were deliberately NOT built, each being the same defect in costume: an **allow-list** of
application schemas (Trap A in miniature — you tune it until it stops catching what it exists to catch); a
**preference for `public`** (a constant with extra steps); a **fallback to the declared schema when the
lookup errors** (a probe that satisfies itself — §5 rule 7).

> **Rule.** When a value is a property of the environment rather than of your code, resolve it *from the
> environment at the point of use*, and make every ambiguous answer a loud failure. Then the next fold needs
> no edit at all — which is the difference between following the platform and chasing it.

**And check whether the re-point is even expensive before designing around it.** The cached snapshot was
assumed possibly-stale (captured before the fold). It was not: `pg.SchemaVersionSQL` digests
`table.column:type` and uses the schema only as a `WHERE` filter, so for a table-narrowed surface the cache
key is **schema-independent** — the digest computed over `public` on a fresh stack equalled the manifest's
`cms`-captured version *exactly*. A one-line `md5` comparison answered a question that had been written down
as unanswerable. Measure the artifact before assuming a migration invalidates it.

Landed as `stack-snapshot/replay/schema.go` (`TargetSchema`, `ResolveTargetSchema`) +
`pg.SchemasHoldingAllTablesSQL`.

---

## 3. Why nobody noticed — pinning disables drift detection

`demo-stack/ensure-clones.sh` phase (e) has a genuinely good freshness subsystem: fetch-verified, refuses to
fabricate a behind-count, LOUD on `stale-by-neglect`. It has one structural blind spot.

1. The behind-count is computed **only** when `ref != "HEAD"` (`:393`).
2. A checkout at a tag **or** at a bare sha is a **detached HEAD**, so `ref == "HEAD"`.
3. Every `pinned-tag` / `pinned-detached` clone therefore gets `behind = null` — never measured.
4. `pin_state` is assigned *before* any staleness test, and neither pinned state increments `_fresh_problems`.

Measured on `stack-demo/clones.lock.json`, 2026-07-31: **11/11 clones detached, 11/11 `behind: null`, 0
freshness problems** — while the bring-up logs *"all clones provably fresh-or-pinned."* Nothing was proven
fresh; they were merely all pinned. And `DEMO_FRESHNESS_STRICT=1`, documented as the go/no-go that "MUST build
current code", escalates only `stale-by-neglect | pin-drift | unknown` — **states a pinned clone cannot
enter.** A pinned clone hundreds of commits behind passes it cleanly.

> **Rule.** Pinning buys reproducibility and pays for it in blindness. A pinned clone still needs its own drift
> signal — a `pin-stale` state meaning *pinned deliberately, and the pin is N commits behind origin's default
> branch* — and `DEMO_FRESHNESS_STRICT=1` must escalate it.

**Also check the pin is self-consistent.** Three things must agree, and on 2026-07-31 they did not:
`.agentspace/rext.tag` (the SoT) named a tag **63 commits** behind `main`, while the consumption clone was
checked out at a *different, newer* tag. `ensure-clones.sh:94-101` treats that mismatch as **FATAL `exit 1`**
(M217 flipped it WARN→FAIL) — so a bring-up aborts. Note that both `ensure-clones.sh:68`'s own header comment
and [`rosetta_demo.md`](./rosetta_demo.md) still described the guard as *"non-fatal"*; the code has exited 1
since M217. When a guard's severity flips, grep for every place that documents it.

**And note where that SoT lives:** `.agentspace/` is **git-ignored** (`.gitignore:138`), so `rext.tag` — the one
file declaring which tooling every demo consumes — is **version-controlled nowhere**. It never appears in a
diff, no review ever sees it change, and nothing makes it stale-visible. That is sufficient explanation for how
it reached 63 commits behind unnoticed. A pin that cannot be reviewed is a pin that will rot; either track it,
or have the bring-up assert its freshness against origin.

**Corollary — a freshness check that compares to a PIN cannot detect a stale CLONE** (M257x iter-12). The
`stack-demo/platform` clone sat **3 commits behind origin/main** while the bring-up printed
`⚠ FRESHNESS: … but clones.pin.json pins '28c5f0d' — PIN-DRIFT`. Both facts are true and the message names
only one of them, so a **stale clone reads as a stale pin** — the reader's eye goes to the pin file, which is
the thing that looks wrong. The pin is a *second* hand-maintained tuple element (§2), and comparing one
hand-maintained value to another can never establish currency against the platform. **Measure distance to
`origin/<default-branch>` explicitly, and say which of the three refs — checkout, pin, origin — is the one
that moved.** M257x's gate says *"against platform @ origin HEAD"* precisely because a clause that does not
name its ref will be satisfied against whatever happens to be checked out.

---

## 4. Detection — six signals, cheapest first

Run these against **platform origin HEAD**, never against a pinned clone. Minutes, not hours.

| # | signal | how | a change means |
|---|---|---|---|
| 1 | `repos.yml` repo set | `gh api repos/anthropos-work/platform/contents/repos.yml?ref=main` | a repo entered or left the clone set |
| 2 | migrating repos | `awk '/^  - name:/{n=$3} /migrations: true/{print n}'` | **the schema-ownership signal** |
| 3 | declared schemas | `awk '/^  - name:/{n=$3} /schema:/{print n" -> "$2}'` | a `schema:` key removed ⇒ that schema is no longer created by the platform |
| 4 | subgraph set | `graphql-wundergraph/supergraph-config-*.yaml` + `ls schemas/` | federation topology changed |
| 5 | compose service set | platform `docker-compose.yml`, **parsed as YAML**, never grepped — plus the OPEN PR list | containers added or removed |
| 6 | org repo census **+ the `archived` flag** | `gh repo list anthropos-work --limit 300` — or, on a box with no `gh`, `curl -H "Authorization: Bearer $GH_PAT" 'https://api.github.com/orgs/anthropos-work/repos?per_page=100&page=N'` with `GH_PAT` from `platform/.env` | net-new repos; **and the cheapest confirmation a fold has completed** |

At 2026-07-31 these read: 10 repos · `migrations: true` → **`app` alone** · `schema:` → **`app -> public`
alone** · **one** subgraph (`backend.graphqls`) · 14 compose services with cms/jobsimulation/roadrunner still
in the default profile · **93** org repos.

**Signals 2 and 3 are the load-bearing pair** — they are what moved under us all three times.

> **Signal 6's `archived` flag is the cheapest fold-confirmation there is, and nothing used to say so
> (M257x iter-20).** Measured 2026-08-01: `jobsimulation` and `skillpath` were archived **2026-07-31**,
> `graphql-wundergraph` **2026-07-30**, `skiller` **2026-07-01** — each within days of its fold, and each a
> one-field answer to a question that otherwise costs a terraform read plus a compose read plus a code read.
> It is a *confirmation*, never a *precondition*: `cms` is folded and **not** archived, `roadrunner` is
> declared folded and **not** archived. And it cuts the other way too — `chronos` is **not** archived while
> the corpus called it archived, which is a corpus error the other five signals cannot see.

### Trap A — `migrations: false` entails nothing on its own

`sentinel` is `migrations: false` **and** alive **with its own `sentinel` schema**
(`docker-compose.yml:43`, `search_path=sentinel`). So a check shaped *"`migrations: false` ⇒ that schema must
not be written"* is **wrong**, and the natural fix — allow-listing sentinel — tunes it until it also stops
catching jobsimulation.

> **A fidelity check against the wrong reference passes.** Derive the schema set from what **actually creates
> schemas** — empirically, `information_schema.schemata` on a cold freshly-migrated stack — not from one
> declarative flag.

### Trap B — the declared topology and the actual topology can disagree by design

At 2026-07-31 `repos.yml` + platform `CLAUDE.md` + `README.md` said the cms/jobsimulation/roadrunner
containers were gone. `docker-compose.yml` still defined all three in the default profile, because the compose
change was a **separate open PR (#20)**. The repo was internally contradictory *on purpose*: docs merged,
compose deferred. Record both, and label which is which.

### Trap C — the platform's own planning docs lag its own code

Do not read their milestone frontmatter as status. At 2026-07-31 `app`'s M709 and M710 `overview.md` both said
`status: planned` and `knowledge/plan/state.md` said "next M702", while the code, terraform and PR stream were
~9 days ahead. **Current truth = `app/CLAUDE.md` + `platform/repos.yml` + terraform `service_desired_count` +
the PR list.** Their prose rots exactly like ours.

### Trap D — the platform ships coordinated multi-repo changes

`repos.yml` moved **39 minutes after** the migration that dropped the `local_*` mirrors. One repo's diff is
never the whole change. When signal 2 or 3 moves, immediately check `app`, `graphql-wundergraph`, and the open
PR list before concluding anything.

---

## 5. Searching without fooling yourself

Every false conclusion in this area has come from a search that returned nothing. **The NUL-byte trap has
become folklore** — it is invoked as an explanation far more often than it is measured, and invoking it *feels*
like rigour while skipping the measurement. In M257x there were three genuine false-absence incidents and
**none** involved NUL bytes:

| incident | actual cause |
|---|---|
| a Playthrough count returning 0 across all 10 manifests | **wrong field name** — matched `id: pt-`; the field is `playthrough:` |
| a corpus mention-index returning 0 for all 93 org repos | ripgrep's default engine **rejects look-around**, and `subprocess` **swallowed stderr**, so an engine error read as absence |
| "`services/next-web-app.md` is binary-skipped by grep" | **did not happen** — clean UTF-8, `rg -c` identical with and without `-a` |

Meanwhile the one tree where NUL contamination was asserted (169 migration files) had **zero**.

Rules, in order of how often they actually catch something:

1. **Never let a search's stderr go unread.** An engine rejection is indistinguishable from "no matches" once
   stderr is swallowed. This is the most common failure by far.
2. **Run a positive control in the same pass** — a pattern you know matches. If it returns 0, the pipeline is
   broken, not the corpus.
3. **Check the field name before concluding absence.** The cheapest false absence is a wrong regex.
4. **Enumerate the search set** (`rg --files | wc -l`) so you know what was actually searched.
5. **Measure NUL bytes before blaming them**: `LC_ALL=C tr -dc '\000' < f | wc -c`.
6. Use `-a` anyway. It costs nothing and removes one variable.
7. **A probe must not be able to satisfy itself.** M257x iter-02 tracked a background script with
   `pgrep -f "ensure-clones.sh"` — a pattern contained in the *watcher's own* command line, so it reported
   `RUNNING` for minutes after the script had exited, and the iter reported a still-running process that did
   not exist. Same family as the swallowed-stderr and wrong-field-name failures above: the probe answered
   without measuring. For process checks use `pgrep -fl` with an anchored interpreter prefix, a PID file, or
   the `$!` captured at launch — and confirm with a question that cannot self-match (`ps -Ao command | grep`,
   then read what actually matched).

8. **A check that SKIPS reads exactly like a check that PASSES.** M257x iter-04: the rext suite's
   `test_all_three_scripts_are_shellcheck_clean` skips when `shellcheck` is not installed. It was not
   installed on the dev host, so it skipped — and the summary line said `1 skipped` next to a wall of
   dots, which everyone read as green. Installing shellcheck immediately produced a real finding that had
   been sitting there since the previous iter. Same for a suite nobody re-runs: the same iter found a test
   that had been RED since iter-02 because only the newly-written tests were run. **Read the skip count,
   and name what each skip covers.** A skip is a hole in the evidence, not a pass.
9. **When you fix a swallowed-output site, sweep its siblings in the same file.** M217 fixed exactly this
   defect — an applier invoked with `>/dev/null 2>&1`, so its diagnosis vanished — for
   `apply-app-authz-skip`. The identical call to `apply-authn.sh` sat **one line above** it and was left
   alone, so M257x iter-04's first bring-up on a new host died in 25 s showing nothing but `EXIT=128`. The
   fix is cheap and the sweep is cheaper than the second incident: `grep -n '>/dev/null 2>&1'` the file you
   just fixed, and justify every remaining one.

10. **Read the lines AROUND the line you are quoting.** M257x iter-07 grepped to
    `SCORED_SECTIONS = ("stack-seeding",)`, read that line and the module docstring, and reported the
    scope limit undocumented — routing a fix forward for it. The justification was the **ten lines
    immediately above the line it quoted**, and it was a good one. iter-08 refuted the finding by opening
    the file.

    This one is not on the list above because it does not look like a false absence: the substring was
    real, the line number was right, and the quote was accurate. **The search succeeded and the
    conclusion was still false**, because a constant's meaning lives in its surroundings. Grepping to a
    line and reading only that line is the cheapest way to be confidently wrong about code you have
    "checked". Open the file, or at least `sed -n 'N-15,N+5p'`.

11. **A probe must exercise the surface whose health it claims — and a hardening change must be re-measured
    on that surface, not on a neighbour.** M257x iter-10. `autoverify`'s academy check read
    `:PORT/library/` and printed *"✓ AI Academy renders its catalog"*. Measured on a live demo:
    `/library/` answered **200 in 9 ms** while `/` — the page the demo's own "AI Academy" link opens —
    answered **500 after 30.0 s**. `/library` is public and short-circuits in Clerk's middleware *before*
    the code path that was broken; `/` does not. The one route the check read was the one route the defect
    spared, so a demo whose landing page was a 500 graded green for four releases.

    The cause was a **security tightening** (M221 tightened `next dev`'s bind `0.0.0.0` → `127.0.0.1`, a
    real and correct de-exposure) whose *literal* had a second, invisible effect: next@16 normalizes every
    loopback hostname to `localhost` when its middleware builds a rewrite URL, keeps the raw `-H` string in
    the router, and compares the two **origins by string equality** — so the tightened bind made the app's
    own rewrite look external and the dev server proxied to itself until a 30 s timeout. The repair was to
    keep the loopback bind and change the literal to `localhost`, the only loopback literal that is its own
    normalized form.

    Two rules fall out, and they are the ones to carry:

    - **Every security tightening ships with a paired "does it still work?" check that exercises the
      affected surface.** M221 shipped its exposure fence (which correctly stayed green — the bind *was*
      loopback) and nothing else. The exposure claim was true; the app was broken.
    - **A boolean probe hides the state it cannot express.** The launcher's own readiness probe used
      `curl -fsS --max-time 3`: `-f` makes a 5xx indistinguishable from silence, and a 3 s per-attempt
      window is *shorter than the 30 s failure it is watching for*. It printed "alive but NEVER ANSWERED"
      over a server that was answering. Capture the status and the exit code separately, and name the state
      you actually measured — *absent* / *hung* / *answering wrong* have different repairs.

    Related to rule 7 (a probe must not satisfy itself) but distinct: here the probe measured something
    real. It just measured the wrong thing, and reported a conclusion about the thing it had not touched.

12. **Say which INVOCATION produced the number, not just which tool.** M257x iter-10 handed forward
    *"autoverify measured 2 FAILED on demo-1, and both are the evidence-log path"* — a precise, sourced,
    confidently-wrong claim. The bring-up's own verdict, sitting in the stack dir the whole time, said
    `warnings:1`. The 2 came from a **standalone re-run of the same script pointed at a different
    directory**: same tool, same stack, different vantage, different answer. The two verdict files differed
    by five hours and nobody compared their timestamps.

    A verifier that takes *where to look* as a parameter can be aimed at a place the thing under test never
    wrote to, and it will then report **absence of evidence in the tool's own voice** — which is exactly
    what a real defect looks like. Two rules:

    - **Record the invocation with the measurement**: the command, the vantage (which clone, which stack
      dir), and the artifact's own timestamp. `warnings:2` is not a measurement; `warnings:2 from
      $(clone A)/autoverify.sh --project demo-1 with STACK_DIR=<workspace root>, ts 20:37Z` is.
    - **Then remove the parameter.** iter-11's fix was not a better message: it was deriving the receipts
      directory from `--project` so the wrong-vantage run is unwritable (§2, and §8 rule 4). The message had
      *already* named the alternative — `"(or STACK_DIR is not the bring-up's $STACK)"` — and iter-10 quoted
      it truncated at the em-dash. A correct diagnostic that has to be read carefully is a weaker control
      than a parameter that no longer exists.

13. **A catalog query that is correct in `psql` can be broken in the program — parameters change the plan.**
    M257x iter-15. `stacksnap`'s sequence-discovery query raised
    `column "sequence_catalog" of relation "sequences" does not exist` on every directus replay, and pasting
    the same SQL into `psql` with the parameters substituted returned cleanly. Both observations were true.

    The predicate was `pg_get_serial_sequence(quote_ident($1)||'.'||quote_ident($2), a.attname) IS NULL`,
    which references only `a.attname` and the parameters — so it is a **restriction clause on
    `pg_attribute`** and the planner is free to push it below the joins that select the relation, evaluating
    it once per row of the whole catalog. `pg_get_serial_sequence` does not return NULL for a column of some
    other relation; it **raises**. With literals Postgres picks a custom plan that resolves the relation
    first and the hazard never fires. pgx PREPAREs, Postgres switches to a **generic plan on the sixth
    execution**, and it fires.

    So the reproduction that "proves the SQL is fine" is a different plan from the one that runs. Two rules:

    - **Reproduce it the way the program sends it** — `PREPARE` + at least six `EXECUTE`s, not a literal
      paste. The sixth is where a generic plan takes over.
    - **A function that RAISES on unexpected input must not sit in a pushable qual.** Pin the relation
      behind an evaluation barrier (`WITH … AS MATERIALIZED`) and derive the function's arguments from the
      **same resolved object** — an OID, not a name re-spelled from the parameters. Then the two arguments
      cannot disagree about which relation they mean under any plan (§8 rule 4). The repair that merely
      makes the function tolerant leaves the correctness plan-dependent, which is how this survived months.

    Blast radius, for calibration: the failed replay also cancelled the post-replay Directus restart (which
    runs only on success), so the whole content layer 403'd — while autoverify reported `green:true /
    0 warnings` on three consecutive cold cycles, because its Directus probe counts registry rows in Postgres
    and never asks the running Directus for an item. **When a step's success gates a side effect, a failure
    costs both**, and the second symptom looks like an unrelated bug.

14. **REGISTERED is not SERVED — a check must be a CLIENT of the surface it grades.** M257x, harden pass 1,
    generalising the blast-radius note in rule 13 from an observation into a rule.

    The only Directus check anywhere in the verify path counted rows in `directus.directus_collections`.
    That table is a **registry**, populated by the *structure* replay; the content is loaded by a *later,
    separate* step, and the anon read grants by a third. So the count is satisfied by a Directus holding
    zero content, and by one that 403s every read — both of which is what the stack actually was. Three
    consecutive cold cycles were graded `green:true / 0 warnings`, and those verdicts were checked in.

    The shape generalises well past Directus, because most content-bearing systems have exactly this split:

    | what you can count cheaply | what the user actually needs |
    |---|---|
    | rows in a registry / catalog / metadata table | an item returned by the running service |
    | a migration recorded as applied | a query against the table it created |
    | a role or grant row existing | an unauthenticated request that gets a 200 |
    | a container reported `running` | the port answering the request the app makes |

    The left column is populated by a *different step* from the right. Counting the left and reporting the
    right is rule 7's self-satisfying probe wearing a convincing disguise — it measures something real, and
    something that is genuinely necessary. It is just not the thing being claimed.

    Three properties make the replacement a measurement rather than a second opinion:

    - **Be a client.** Go over the wire the way the consumer does — the running service, the stack's own
      offset port, an unauthenticated request if that is how the content is consumed. A DB count and an
      HTTP read are *independent* measurements; two DB counts are one measurement twice.
    - **Derive the target, never hardcode it.** Ask the environment what to read (here: the non-system
      collection holding the most rows), so a re-modelled surface cannot make the check stale, and so the
      check cannot pick a target chosen to guarantee its own success (§2, §5 rule 7).
    - **Fail closed on an empty derivation.** "Nothing to check" and "nothing is there" are the same
      observation from the check's side and opposite verdicts from the operator's. If the derivation finds
      no target, that IS the defect — say so; do not pass. Every silent-skip in this milestone read as green.

    And name the states distinctly: *403* (holds it, serves it to nobody), *200 with an empty payload*
    (serving, but the content is not there), *no response* (not serving) have three different repairs, and a
    boolean collapses them (rule 11).

15. **"It only reproduces in the full pipeline" is a claim about the failing step's INPUTS — check them
    before you pay for a pipeline run.** M257x iter-17 routed a Directus bootstrap failure forward with an
    explicit *"do not try to reproduce this by hand — it has healed; the diagnosis arrives through a cold
    cycle"*. That was true of **one artefact** (the `directus` schema on one stack, which the later replay
    had since filled in) and was silently generalised to *the failure is not reproducible*. The failing
    step's actual inputs were a DSN, a container image and an empty schema. Recreating them took **four
    minutes** and refuted the routed hypothesis before a line was written: the command exits **0** on a
    fresh schema, so the failure was context, not command.

    The context turned out to be a **second actor nobody had enumerated** — the `directus/directus` image's
    own `CMD` is `node cli.js bootstrap && pm2-runtime start`, so the compose service bootstraps the schema
    itself and races the pass's one-shot. Enumerate who else writes the thing you are writing; a container
    image's entrypoint is a participant, not scenery.

    **And the second half, which is what makes this rule expensive to ignore: a nondeterministic defect
    makes a green run weak evidence.** The race went our way on 2 of 3 cold cycles; those two cycles are
    green on the *unfixed* code too, and prove only that the fix did not regress the winning path. Only the
    third cycle exercised the branch under repair. It also dissolved a standing puzzle — the same
    nondeterminism is why one set of three cycles read green and a later single cycle read red on the same
    code. **Record which path each run took**, or a battery of greens will certify a fix that was never
    invoked, and a flake will be filed as a regression.

16. **An UNREAD metric is indistinguishable from an UNMOVED one — and the no-progress rule cannot tell them
    apart.** M257x iter-32 opened with what the ledger showed as three consecutive no-progress iters, which
    is the streak that fires a strategy revision. It was false. Iters 30 and 31 had each fixed a failing
    Playthrough and verified it live; what they had declined to do was *re-read the headline number*,
    because this protocol forbids quoting a scoped run as a binding one and the full read had been budgeted
    as an entire iteration. The measurement, when finally taken, moved `25 → 27` and attributed both points.

    A protocol that makes its primary measurement expensive **manufactures phantom stalls**, and the streak
    rule will then revise a strategy that was working. **Decide a no-progress trigger by measuring, never by
    counting ledger rows** — and where the two disagree, the ledger is the thing that was wrong.

    **The second half is cheaper and matters more: re-measure your own cost estimates.** The budget that
    caused the stall — *"~35–40 min serial, 209 specs"* — had been carried across seven hand-offs unchecked.
    Measured: **4 min 50 s**, reset included. It was never a bad estimate so much as a bad *proxy*: only ~31
    of those 209 specs are Playthroughs, the rest are unit specs at 0–1 ms, and the count had been standing
    in for the duration. An inherited cost estimate is a claim like any other, and this one had been
    silently scheduling the work.

17. **A count can be exactly right while the claim it supports is false — verify the PREDICATE, not just
    the denominator.** M257x iter-34, before any auditor reported, independently re-derived the corpus's
    multi-tenancy fence — the one that had *already been wrong twice*, both times failing toward
    *"isolation is handled."* Every number reproduced exactly: 30 schemas with the policy mixin, 7 with the
    id-only mixin, 18 with a bare `organization_id` and neither, plus all nine named example files. The
    re-derivation concluded the fence held.

    It did not. The sentence read *"…with no mixin **and no policy at all**"*, and the check had tested only
    the first conjunct. One of the 18 — the one listed **first** — declares its own fail-closed `Policy()`.
    The real split is 31 policed / 17 unpoliced, and the fence had failed reassuringly for a **third**
    consecutive generation.

    **When a claim is a conjunction, a measurement of one conjunct is not a verification.** Read the
    sentence as a predicate and test every clause of it. Where a list is load-bearing, ship the derivation
    command next to it so the next reader re-derives instead of trusting.

18. **Text written to repair fidelity debt is the highest-risk text in the corpus — measure the residual
    split by swept vs unswept.** M257x iter-33 measured a **24 %** self-inflicted rate on its own repair
    pass and recorded it as a caution. iter-34 re-measured across all 40 files and found the effect is the
    *dominant* term, not a caution: **9 of 11** remaining blockers sat in the **13** files the repair had
    touched, **2** in the **27** it had never opened — **0.69 vs 0.074 per file, a ~9× density difference.**
    Two auditors reported their never-edited files clean across ~40 exact citations, unprompted.

    Two consequences. **(a)** Never close a corrective sweep without an adversarial pass over the sweep's
    own diff; the errors will not be in the `file:line` anchors — those verify — but in the surrounding
    prose, the summary line above the section, and the bullet three paragraphs down that restates the
    retracted claim in different words. **(b)** When a confirming pass runs, **partition it differently
    from the pass before it.** Correlated blind spots are a property of how the corpus was divided, not
    only of who read it.

19. **Repair by CLAIM, not by FILE — and half-repairing a uniformly-wrong corpus is worse than leaving it
    alone.** The partition that is *correct for reading* is *wrong for repairing*. Six auditors each owning
    a disjoint file set is precisely what surfaces independent double-finds (rule 18(b), which has now paid
    three times — M257x iter-38 had two auditors refute the same false EU-AI-Act premise from two different
    files; iter-39 had two find `hiring.md:86` independently). But a **claim does not respect a file
    boundary**, and a repairer who owns `external_services.md` cannot fix the same sentence where it also
    lives in `graphql-wundergraph.md`.

    Measured: **5 of iter-39's 8 self-inflicted defects were cross-file drift** — a claim corrected in the
    file its owner held while the identical claim stood in a twin owned by somebody else, or in no
    partition at all. The EU-first fallback ladder, the subgraph count, the taxonomy figures, the
    `--template` flag and the jobsimulation subscriptions each survived in one to seven sibling sites the
    same sweep had edited for other reasons.

    **Why it is worse than doing nothing.** A uniformly-wrong corpus is at least *self-consistent*: a
    reader who catches the error once has caught it everywhere. A half-repaired one teaches the reader that
    the corpus contradicts itself — and the next auditor spends its budget **adjudicating rather than
    measuring**, which is the one thing a fidelity measurement cannot afford.

    **The procedure.** Before editing, grep the **whole tree** for the claim — not the file set the audit
    partitioned, and not only the audit's own scope. Fix every instance in one pass. Then re-run the same
    grep as a post-condition.

    **Corollary, and it is the half that gets missed: a claim leaks to the EDGE of the previous repair's
    scope and stops there.** M257x iter-40 swept every claim iters 38/39 adjudicated and found the 40
    in-scope files **uniform on all of them** — while every surviving instance sat just outside:
    `corpus/ops/**`, `.claude/skills/**`, and `CLAUDE.md`, the repo's top-level instruction file and the
    **highest-propagation site in the tree**. An audit's scope is a legitimate boundary for *reading*; it is
    never one for *repair*. Sweep outward from the scope edge, and check the instruction files first — they
    are read by every agent before any doc is.

    **What a claim-scoped repair must NOT do: adjudicate.** It propagates a verdict already established
    inside the audited scope, verbatim or by link to its canonical anchor. If a surviving site's correct
    form is not already settled, route it — deriving a fresh claim during a repair pass is how rule 18's
    highest-risk text gets written.

    **Derive the claim LIST from the prior pass's ledger — never assemble it by hand.** M257x iter-40 swept
    eight claims, reported the audited scope "uniform on all of them", and had in fact verified **five of
    the eight**. The next pass found two consequences: a `:5050` port claim still live *inside* the audited
    scope (fixed at eight sites outside it), and a retracted academy-fallback claim that was never on the
    list at all. A hand-assembled list silently redefines "every instance" as "every instance I thought of",
    and the post-condition re-grep then confirms exactly that smaller thing. **Enumerate from the ledger,
    grep per ledger entry, and state the coverage as a fraction.**

20. **Measure what the repair INDUCES, not only what it leaves.** Every corrective pass reports a residual;
    almost none reports its own contribution to the next one. M257x ran six KB-fidelity passes returning
    25 → 13 → 11 → 17 → 37 → 18. Only the last is comparable to its predecessor — the sixth held the
    instrument fixed on every knob *and* read a corpus the intervening repair had not touched, so
    **37 → 18 is a real measurement of the repair: it halved the residual.**

    Then the decision-relevant number, which no earlier pass had computed: **9 of those 18 were
    manufactured by the repair that preceded them** — over-corrections inside newly-written blockquotes, a
    retraction that over-reached from *"does not exist now"* to *"never existed"*, a blockquote spliced into
    a bullet list orphaning the member that stated a legal consequence, and half-applied edits that
    corrected one twin and left the other. A clean **50/50 induced/genuine split**.

    > **A repair loop whose induction rate approaches its removal rate has a fixed point, and it is not
    > zero.** Without the induced/genuine split, `37 → 18` reads as convergence and justifies another pass.
    > With it, the next pass is predictable: repair 18, induce ~9, measure ~9–15.

    **So: before booking another pass, classify the current findings by whether the last repair created
    them.** If the induced share is near half, the answer is not another pass — it is a fence, a narrower
    clause, or an explicit decision to accept a bounded residual. **Two further regularities point the same
    way:** across five consecutive passes **every `file:line` anchor a sweep introduced resolved
    correctly**, and every defect was in surrounding prose — so the failures live exactly where a machine
    check could reach and a hand sweep cannot. And in two consecutive iterations **the author of a
    newly-written rule violated it while writing it**, which is evidence that a hand-applied discipline does
    not survive a corpus of this size.

21. **Classify a residual by the CHEAPEST INSTRUMENT that would catch each finding — before concluding it is
    irreducible.** Rule 20 gets you the induced *rate* and offers three exits: a fence, a narrower clause, or
    an accepted bounded residual. **Which of the three is available is decided by the induced findings'
    KIND, and a rate cannot tell you that.** A count is an aggregate, and an undifferentiated count invites
    an undifferentiated remedy — repair harder, or give up on the target.

    M257x iter-42 re-read the same 18 blockers iter-41 had escalated on, taking each row's class from the
    ledger's own *"what is true"* column, and asked one question per row: *what is the least expensive thing
    that could have caught this?* The answer split them **13 / 3 / 2** —

    - **13: the corpus contradicting ITSELF** — a twin site inside the same repository already states the
      opposite, five of them within a few lines *in the same file*. Detectable with **no ground-truth read
      at all**.
    - **3:** an anchor that **resolves but names the wrong construct** — the right line, the wrong function
      or table row. The existing anchor check proves a line *exists*; it never asks whether the line *says
      the thing*.
    - **2:** a derived scalar contradicting source with no corpus twin (a version string, a memory limit).

    And the same cut applied to the **induced** half was decisive: **8 of the 9 repair-induced findings were
    the single self-contradiction class** — a claim repaired at one site and left standing at another.

    > **A ~50% induction rate is a property of the repair method, not a law about corpora — and homogeneity
    > is what makes mechanization available.** Had the induced half been nine different kinds of wrong, no
    > fence could reach it and rule 20's pessimism would stand. Because it was one mechanical class, the
    > fence exit opened. **A rate says *stop*; a class says *what to build*.**

    Two corollaries worth carrying:

    - **The self-contradiction class is the cheapest one in any corpus, and almost nobody checks it** — it
      needs no access to the system being documented, only the documentation. If a corpus has been repaired
      more than once, look here first.
    - **The fixture is perishable.** A corpus carrying a known, anchored, independently re-verified answer
      key exists exactly once; repairing it destroys the only thing that can falsify the fence you are
      building to protect it. **Build the fence and watch it go RED before the repair, not after** (§8, and
      rule 8 — a check that SKIPS reads exactly like a check that PASSES).

And: **verify a claim before escalating it, including a claim made by an audit.** In M257x two probes
contradicted each other on whether `public.sessions` exists; measuring settled it (it does not — created then
dropped as a rename completed) and *inverted* the risk assessment that had been built on it.

### Trap E — the tooling's own host preconditions are invisible until a clean host

Everything above is about the platform moving. This one is about *us*: a tooling path that quietly depends on
something the developer's machine happens to have will work on every machine that has it and fail on the first
that does not — and it will fail at whatever moment the new machine arrives, which is never a convenient one.

M257x iter-04, on the first bring-up ever attempted on a new Mac: `apply-authn.sh` cloned **colony — a private
repo — from an anonymous `https://github.com/...` URL**. That URL can only succeed if git finds an ambient
credential helper or an `url.insteadOf` rewrite. Every other rext acquisition already used
`git@github.com:$ORG/...` and even told the operator to check `ssh -T git@github.com`; this was the lone
outlier, and it had been fine for as long as nobody started from a clean box.

> **Rule.** Acquisition uses ONE convention, and the failure names the credential it wanted. When a bring-up
> fails on a new host, check what the tooling assumed about the host *before* concluding the platform moved —
> the mirror image of §9's "check signals 2 and 3 first" on a familiar one.

---

### An audit scoped by SEARCH TERMS measures the terms, not the corpus (M257x iter-21)

Four KB-fidelity runs over the same two trees, same day, same ground truth:

| run | method | blockers |
|---|---|---|
| 1–3 | sweep the **drift surface** — grep the router/subgraph/schema vocabulary, read around the hits | 11 → 5 → **2** |
| 4 | **read all 40 files in full** | **21** |

Runs 2 and 3 each recorded that their findings were *pre-existing, not regressions*. So `11 → 5 → 2` was
never convergence; it was a grep converging on its own vocabulary.

**Why the terms miss.** The dominant failure mode is *a corrected banner at the top of a file contradicted by
prose further down* — and that prose rarely uses the banner's words. The misses were `make init-studio`,
`docker compose up -d graphql`, a `depends_on` naming a deleted service, and an **arrow in a mermaid
diagram**. No grep for "router", "subgraph" or "merged" reaches any of them. The mermaid edge is the sharpest
case: `Web --> GraphQL` contains no drift vocabulary at all, and survived three audits for that reason.

> **Correction (M257x iter-22).** This list originally also named
> `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`. **It was not a miss — it is correct at origin HEAD**
> (`docker-compose.yml:52`, `:258` @ `2adcf71`). See the rule below; the example is kept struck rather than
> deleted because it is the one that taught the rule.

> **If the deliverable is *"this tree is true"*, the audit reads the tree.** Scope by *file set*, not by
> search term; fan the read out across sub-agents if it is large. The cost difference here was minutes, and
> it was the difference between a claimable clause and a 10x under-count.

**And check the branch against its base first.** Eight of run 1's eleven blockers were already fixed on
`main`; the milestone branch had been 3 commits behind for its whole life and nothing measured it. One
command — `git rev-list --count HEAD..main` — and the audit is not measuring a tree no reader will ever see.

### Re-derive the CORRECTION, not just the anchor (M257x iter-22)

iter-21 handed iter-22 an enumerated residual: 21 blockers, each with `file:line`, a verbatim quote, a
refuting citation and a one-line correction — authored to be executed without re-exploration. **All 21
anchors verified.** Two of the *corrections* were false.

Items #8/#10 said `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` was stale and should read
`http://backend:8083`. Origin HEAD `2adcf71` says otherwise: **only `SKILLER_RPC_ADDR` was re-pointed**;
`CMS_RPC_ADDR=http://cms:8091` and `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` still address the husk
containers — deliberately, per `app/main.go:1196-1202`, *"additive + DORMANT: external callers (messenger)
keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**."* Applying the correction
would have replaced two true statements with false ones.

**Where it came from is the whole lesson.** The refuting citation iter-21 trusted was
`corpus/services/backend.md:175` — a corpus line asserting messenger points *all four* addresses at
`backend:8083`. It points two. **One false corpus line, cited as authority, produced two false corrections in
a hand-off designed to be applied mechanically.** An audit that reads the corpus to correct the corpus is
circular; the citation must terminate in platform source.

> **The failure mode a mechanical hand-off invites is not a moved anchor — it is an inherited falsehood
> wearing a `file:line`.** Anchors are cheap to verify and they were all fine. Verify the *correction* against
> platform source (`docker-compose.yml` / `repos.yml` / the service's Go) before you apply it, every time. And
> when a correction turns out to be wrong, the line that misled you is itself a blocker: hunt it.

Corollary, worth stating because it reads as pedantry until it costs you: **merged-in-production is not
removed-from-compose.** `cms` and `jobsimulation` are `service_desired_count = 0` in prod, folded into `app`,
subgraphs gone — and still start containers on every local `make up`, still answer RPC. Two service docs said
*"not in the local compose"*; both were false. The map's word for this is `running_but_unfederated`. Use it.

**A correction can be INCOMPLETE rather than wrong — and that is harder to catch (M257x iter-23).** iter-22's
hand-off said colony was *"split: `app` + `messenger` @ `v0.35.2`; `sentinel` + `storage` @ `v0.34.3`"*, which
is true of the four services it names. It names four of **six**: the `cms` and `jobsimulation` containers the
default profile still starts are on a **third** pin, `v0.35.1`. Applying it verbatim would have replaced one
incomplete claim about that table row with another, and the row would have read as freshly verified. **When a
correction enumerates, re-derive the ENUMERATION, not just the values** — the question is not "are these two
pins right?" but "is this the whole set?"

### A named-consumer list survives the merge that moved the consumer (M257x iter-23)

The founding class has a second face, and it is not about schema names. When a service folds into `app`, any
**list of service names** that encodes "who talks to X" is silently wrong the moment the domain moves — and
unlike a dropped table, **nothing errors**: the list still names a real, running container, which still starts,
still holds the env var, and still answers. The read simply happens somewhere else.

Measured instance: rext's `--local-content` cutover re-points `DIRECTUS_BASE_ADDR` at the per-stack Directus
for every service in `DIRECTUS_DATA_CONSUMERS`, which is `cms` — correct when `cms` was the Directus consumer,
with a test (`test_only_cms_is_repointed_not_other_services`) explicitly asserting `backend` must **not** carry
it. Since cms-in-app, `app/cms_reader_switch.go` swaps `backend`'s content reader to the **in-process** cms
server ("*no internal traffic to a standalone cms*"), so `backend` reads Directus directly, over its own
`DIRECTUS_BASE_ADDR` — which comes from `env_file: .env`, which nothing re-points. Live on `demo-1`:
`cms` → `http://directus:8055`, `backend` → `https://content.anthropos.work`, `DIRECTUS_TOKEN` empty. The
per-stack Directus serves a consumer that no longer reads, and the reader is pointed at prod anonymously.

**The test made it worse, not better.** It pinned the pre-merge shape as a contract (§8 rule 3) and would fail
on the fix. Two of the three previous occurrences had the same signature — *the suite was not silent about the
defect, it was arguing for it.*

> **The check to run after any fold:** for each service the fold touched, grep the tooling for its **name** as
> a value (consumer lists, `depends_on`, front/proxy port tables, probe targets, env re-point maps) — not just
> for its schema or its tables. Then ask, per hit, *does the code that actually performs this read still live
> in the named service?* A `docker inspect <container> | grep <VAR>` on a live stack answers it in one command
> and beats any amount of reasoning.

---

## 6. Classification — the map

Produce [`platform-migration-status.md`](../architecture/platform-migration-status.md): every service the
platform has ever had, each row cited to a **sha or file:line**.

States: `live-standalone` · `merged-into-app` · `running_but_unfederated` · `decommissioned` · `net-new` ·
`external` · `library`.

**Every row carries TWO states — one for production, one for a fresh local stack.** They genuinely differ, and
collapsing them is why this class recurred three times: the corpus never distinguished the two, so it had to
pick one and be wrong about the other.

| jobsimulation, 2026-07-31 | production | fresh local stack @ origin HEAD |
|---|---|---|
| process | scaled to zero (`terraform/main.tf:40`); `app` owns it in-process | **container still starts** in the default profile |
| schema | legacy `jobsimulation` schema still present, pending M710 | not created by the platform — **but created by rext itself** (§2) |
| subgraph | gone (1 subgraph) | gone (1 subgraph) |

**`merged-into-app` is not `mid-merge`.** Ask three separate questions: does `app` own the code and call it
unconditionally; do the tables live in `public` under `app/terraform/migrations/`; is the standalone scaled to
zero? All three yes ⇒ merged, regardless of whether the repo or its container still exists.

---

## 7. Re-point — the procedure

1. **Measure the write surface first**, by scan, not memory — and split **live code** from **comments**. They
   are not the same finding. In M257x, 69 lines named a dropped mirror and **0** were live.
2. **Re-point to the canonical target, not to nothing.** Deleting a write is satisfiable and silently
   re-empties the surface (the M219/M222 render-gate trap). Every removed write needs its replacement
   asserted.

   **The exception that proves it — the co-written PAIR.** Sometimes the canonical target is *already
   being written on the next line*. In M257x iter-06 three seeders wrote the session twice — the same
   rows, the same ids, into `jobsimulation.sessions` **and** `public.job_simulation_sessions` — because
   M257 had added the app-side half beside the legacy one rather than replacing it. Re-pointing the
   legacy half would have written one table twice. So the fix there is **removal**, and rule 2 is
   satisfied *by inspection of the very next step*, not by adding anything. Before deleting either half
   of a pair, check which half the platform reads today; before re-pointing either half, check they are
   not about to become the same statement.

   **And re-check the ORDER after a re-point.** Two writes in different schemas have no FK between them,
   so their order was free; once both land in `public` the FK is real and the order is load-bearing.
   iter-06 had to move `public.job_simulation_sessions` to the FRONT of two flush slices (every other row
   in the fan-out FKs it) and move `actors`/`interactions` ABOVE it in `resetTables` (they became true
   children of a list entry). **A re-point can silently turn an arbitrary ordering into a required one.**
3. **Leave the history in a comment.** Name what the relation *was* and which migration removed it, so the
   next reader who greps the dead name lands on the explanation rather than on nothing.
4. **Advance the pins deliberately** and record what the advance contained.
5. **Prove it cold.** `demo-down --purge` + `demo-up`. A warm cycle hid a four-day-old total breakage.

---

## 8. Fence — so it cannot silently recur

Three layers. Each must be **watched going RED** before it is trusted — M256 found **43** checks that reported
success without checking.

| layer | asserts | lives in |
|---|---|---|
| map ↔ `repos.yml`, both ways | every `repos.yml` service has a map row; no map row invents a repo; every state is one of the seven; every row cites a sha or `file:line`; the net-new census does not overlap the clone set | `stack-core/platform_alignment_guard.py` (M257x iter-20) — landed, with `tests/test_platform_alignment_guard.py`. **What runs unconditionally is C+D+E** (vocabulary, citations, census overlap): those are properties of the map alone, so they are checked against the **real** map on every suite run, everywhere. **A and B need a real `repos.yml`, and a platform clone is git-ignored and ephemeral** — the test searches `PLATFORM_REPOS_YML` then every `stack-*/platform/repos.yml`, and *skips A/B/E-against-the-clone-set* if it finds none, naming what it looked for. Until M257x harden pass 5 it looked at one hardcoded path and this row claimed the whole thing ran "on every suite run" — on a box with no demo stack it silently skipped, i.e. the fence's own documentation was the claim-without-a-measurement shape this protocol exists to catch |
| static schema fence | every schema a seeder WRITES to **through a statically-visible construct** is one the migrate step CREATES | `stack-core/tests/test_write_target_schema_fence.py` (M257x iter-06) — reads the legal set from `repos_yml_schemas_to_create`, so it **names no dead schema at all** |
| live schema assert | every schema rext writes exists in `information_schema.schemata` on the migrated stack | bring-up / autoverify (precedent: `dev-stack/tests/test_migrate_dev_live.py:144`) |

### A fourth layer, on a different axis: fence the PROSE against verdicts already recorded (M257x iter-43)

The three layers above all fence **tooling against the platform**. None of them reads a sentence, and §5
rules 20–21 measured what that costs: six KB-fidelity passes returning `25 → 13 → 11 → 17 → 37 → 18`,
with **9 of the last 18 manufactured by the repair that preceded them** — and **8 of those 9 were one
mechanical class**, a claim repaired at one site and left standing at another. The corpus contradicting
*itself*, which needs no read of the platform to detect, because the verdict is already written down.

| layer | asserts | lives in |
|---|---|---|
| claim-twin fence | no claim an audit has already refuted is still published anywhere in `corpus/**`, `.claude/skills/**` or `CLAUDE.md` | `stack-core/claim_twin_guard.py` + `claim_ledger.py`, with `tests/test_claim_twin_guard.py` and an 8-mutant battery |

Four properties are what make it a fence rather than a linter, and each is a rule from this document
applied to prose instead of to code:

1. **The claim list is DERIVED, from the audits' own blocker-ledgers** — §5 rule 19's list-derivation
   clause. It is recognised by table *structure*, not by a list of filenames, so it is not the
   hand-maintained tuple §2 deleted wearing a new hat.
2. **`## Minors` sections are excluded, and that cut decides whether the fence is usable.** A blocker
   row records a claim that is *false*; a minor row overwhelmingly records an anchor that drifted while
   *"the claim itself is TRUE"*. Matching those tree-wide fires on correct prose, and §8 rule 6 already
   says where that ends: a fence that cries wolf gets disabled. Measured before the cut, **12 of 17**
   minor-sourced hits were true prose with a stale anchor.
3. **Scope is tree-wide from the first run** — §5 rule 19's corollary, that a claim leaks to the edge of
   the previous repair's scope and stops there. Its first run found exactly that: a claim iter-34
   refuted, repaired inside the audited scope, still standing in `corpus/ops/**`.
4. **It was watched going RED on a fixture with a known answer key, before that fixture was spent.**
   §5 rule 21's perishability clause, obeyed literally: the 18 sites were captured to
   `tests/fixtures/claim_twin/` at rosetta `48ca53c`, together with a GREEN twin of each — because a
   battery of REDs cannot tell a discriminating fence from a brittle one.

**And it does revise one design decision below** — *"keep `.md` prose out of scope; that is review, not
a fence."* That bullet is right about the fence it was written for, and it stays right: a write-target
fence must not read prose, because per Trap B the prose can be false at the same sha. The claim-twin
fence does not judge prose either. **It only ever asserts a verdict some auditor already recorded, with
an anchor** — it never adjudicates, which is the same line §5 rule 19 draws for a claim-scoped repair.
The distinction is worth stating rather than leaving for a reader to notice as a contradiction.

### Run the fence at the COMMIT, or it cannot touch the number that matters (M257x iter-44)

The fourth layer above is an *instrument*. Six passes measured what an instrument alone buys, and the
answer is the one that ends the loop: **9 of iter-41's 18 findings were manufactured by the repair that
preceded them.** An audit-time fence cannot reduce that term — by the time it runs, the induced defect
is committed, and it is one of the findings being counted. The induced term is only reachable **between
the repair and the commit**.

| layer | asserts | lives in |
|---|---|---|
| repair post-condition | the set of published sites restating an already-refuted claim may **shrink or stay — never grow** across a commit | `stack-core/repair_postcondition.py`, with `tests/test_repair_postcondition.py` and a 12-mutant battery |

The premise it rests on is measurable rather than rhetorical: an induced self-contradiction **is** a new
`(file, claim)` pair — a repair wrote the correct sentence at one site while an adjudicated form of the
same claim stayed published at another. That is the shape of all eight of the nine, so the class becomes
unrepresentable in a commit rather than merely visible in the next pass.

Three decisions decide whether that is a fence or a slogan, and each is a rule from this document:

1. **The fence registry is DERIVED from disk, not held as a list.** Every `*_guard.py` declares a
   `FENCE_KIND`, read *statically* (an import would let a guard that fails to import look like a guard
   that declares nothing), and an undeclared guard makes the run **exit 2 naming itself**. §2 deleted a
   hand-maintained tuple of services; a runner with a hardcoded list of fences is that tuple wearing a
   new hat, and it fails the same way — silently, by omission, exactly as iter-08 measured (*a fence only
   ever asserts about what it already scans*).
2. **A site's identity carries no line number** — it is `(fence, path, claim_id)`. Keyed on `file:line`,
   every edit above a known site would read as an induced defect, and §8 rule 6 already says where that
   ends: a fence that cries wolf gets disabled. Lines are still reported; they are not the identity.
3. **The baseline is a ratchet, not a diary.** `--accept` lowers, and **refuses to raise** for a fence
   already recorded, naming the sites that grew. A fence *absent* from the baseline is the other case —
   its first measurement is a registration, not a regression — so it is admitted and announced as a
   baseline, never as a pass.

**Two vehicles, because each covers the other's blind spot.** The suite runs it against the real tree in
every clone; a `--install-hook` pre-commit runs it at the moment of the repair, when the fix is cheap.
Git hooks are per-clone and unversioned — a real limitation, stated here rather than left for a reader to
discover, and the reason the suite is the load-bearing vehicle.

> **And every non-failure it reports has a mutant that silences it.** Harden passes 7–9 found that two of
> the claim-twin fence's own honesty mechanisms did not exist: one deleted clean with 15/15 still green,
> and one was promised by a docstring and never implemented. **A reporting path with no test is a
> docstring**, which is the same claim-without-a-measurement shape this whole document exists to catch —
> committed, twice, inside the tooling written to catch it.

### Say which layer covers which part — and derive that too (M257x iter-08)

Three layers only help if something records **which one covers what**. Until iter-08 nothing did, and the
cost was immediate: iter-07 read the static fence's scored-sections constant, concluded its scope limit was
undocumented, and routed a fix forward. The limit *was* documented — in ten lines directly above the
constant it quoted — and re-measurement refuted the finding. **The milestone's own dominant defect,
committed by the milestone: a claim reported without being measured.** §5's closing rule already said it —
*verify a claim before escalating it, including a claim made by an audit* — and it applies to your own audit
of your own tooling.

Two things came out of that, and both are worth copying:

1. **A fence's SCOPE is a hand-maintained list of the system's parts, exactly like the migrate tuple §2
   deleted** — and it is the worst place for one, because *a fence only ever asserts about what it already
   scans.* An unclassified section cannot go RED; it is invisible by construction. So the scope is now
   derived: every Go-bearing section on disk must carry a declared `(layer, reason)`, and a new section goes
   RED naming itself. v9.0 adds surface; this is what makes that arrive loudly.

2. **A section classified "static" that yields ZERO scoreable constructs is mis-classified, not covered** —
   and it reports GREEN, which is strictly worse than leaving it out, because it *looks* fenced. Measured:
   widening the scored set to the other five Go sections would have scored **0** constructs. `stack-snapshot`
   genuinely belongs to the **live** layer, because after `D-M257x-8` its write target is resolved at run
   time and there is nothing static left to see.

   > **Rule.** Assert that each declared scope actually *matches something*. "I scanned it" and "I found
   > nothing to check in it" are different findings, and only one of them is coverage.

**And re-read a fence's stated rationale whenever the thing it justifies changes.** The pre-iter-08 comment
argued `stack-snapshot` was safely out of scope because its one stale target *"already fails LOUD at replay
time (rc=4)"* — a signal iter-07 had **removed** by making the replay resolve and succeed. A justification
whose evidence has been deleted still reads as live. When you close a route, grep for the comments that
cited it.

The fence reads **only machine-readable fields** (`name` / `type` / `migrations` / `schema`) — never the prose
comments, because per Trap B the prose can be false at the same sha, and fencing on it would mechanically
encode a falsehood.

**Preserve the existing fence's design decisions** — they are correct and were paid for:

- score the **occurrence**, not the line;
- **allow comments unconditionally** — a fence that forbade naming a dead relation would delete its own
  rationale;
- assert the **positive** replacement, not only the negative absence;
- keep `.md` prose out of scope; that is review, not a fence.

### A fence over source must assert against a parsed construct, never a whole-file substring

The corollary of "allow comments unconditionally", and M257x iter-02 paid for it twice in one iter.

`test_migrates_the_four_merged_services_and_never_skiller` (v2.1 M209) asserted the four pairs were present
in `migrate-dev.sh`. When the loop was changed to derive its set, the test **still passed** — satisfied by
the tuple appearing in the new *comment* explaining why it had been removed. A test whose entire purpose was
to pin the migrate set was satisfied by its own refutation. And the replacement fence's own prose-comment
fixture initially **could not fail**, because it placed the lying values where the parser resets past them.

Both are the M256 reports-success-without-checking class. Rules:

1. Assert against the **construct** — the loop body, the derived value, the AST node — not `file.read()`.
2. **Mutation-verify the fixtures too**, not only the fence. A fixture that cannot fail proves nothing, and
   reads exactly like a passing test.
3. A fence that pins **the current shape of the drift** is worse than no fence: it converts the bug into a
   contract, and every future correct change has to argue with it. Pin the *mechanism* (where the list comes
   from), not the *contents* (what happened to be in it).
4. **Prefer a construct that cannot express the drift over a fence that catches it.** A fence can only
   catch a drift someone has already written; a shape that makes the drift unwritable is strictly better
   when it is cheap. M257x iter-07: the pre-replay digest probe and the replay itself both had to read a
   schema, and moving only one of them leaves the surface skipping at `rc=4` before a row is copied — with
   a diff that looks complete in review. Instead of moving both and trusting review, they were merged into
   one function that computes the probe's schema argument *inside itself*: **there is no parameter for a
   caller to supply, so there is no way to supply the wrong one.** Keep the fence too (the construct can
   be refactored apart later), but spend the design first.

   The same instinct applies to optional parameters: when a value must be decided at every call site, make
   it **positional and required** rather than an option with a default. iter-07's `replay.Run` gained a
   required `TargetSchema`, so the compiler forced all 18 existing call sites to state which behaviour they
   meant. A forgettable optional with a silent identity default is a fence you have to remember to use.

5. **A mutation that does not COMPILE is not a RED fence.** iter-07's mutation battery reported
   `RED (good)` for a mutation that had merely removed the last use of an import: the package stopped
   compiling, `go test` returned non-zero, and the harness read that as the fence firing. The tell was an
   **empty list of failing test names**. A compile break proves the mutation was applied and proves nothing
   about whether anything noticed the behaviour change. Re-run with a compiling mutant, it fired for real.

   > **Rule.** Gate every mutation on an explicit build BEFORE the test run, and **name the test that went
   > red**. A battery that reports only exit codes will sign off on a fence that does not fence — the same
   > family as §5 rule 8 (*a check that SKIPS reads exactly like a check that PASSES*).

   **And read the COUNT together with the exit code — "no tests collected" is a non-zero exit.** M257x
   iter-11's own mutation harness reported a clean `RED` for a mutant nothing had tested: the harness had
   invoked a `python3` without pytest, so every run returned non-zero and every mutant "fired". The same
   shape occurs without any tooling accident, because **pytest exits 5 when a `-k` filter matches nothing**
   — a renamed test makes its own mutation battery report RED forever. The tell is identical to iter-07's:
   an empty list of failing tests, here surfaced as `collected=0`. Gate on `collected == 1` *and* an exit
   code that means failure rather than emptiness; never on non-zero alone.

   **A mutant that changes nothing is not a surviving fence.** The same iter recorded a `GREEN (mutation
   SURVIVED)` that was neither: the mutant added an unreachable `case` arm below an early `return`, so the
   function's behaviour was unchanged. The honest reading is *the mutation was a no-op*, and the reason it
   was a no-op is worth keeping — the invariant was enforced **twice** (an early return AND a closed
   `case`), which is §8 rule 4 working. When a single-point mutant survives, first ask whether it changed
   behaviour at all; if the property is doubly enforced, write the two-point mutant a future editor would
   actually write (here: *"make the helper total"* — drop the guard, add a default arm), and confirm it
   parses, collects, and goes red.

   **So put a no-op mutant in the battery ON PURPOSE, and require it to survive.** iter-11 arrived at
   "a mutant that changes nothing is not a survivor" as a *disqualification* rule — a way to throw out a
   bad GREEN after the fact. It is more useful as a **positive control** run alongside the real mutants.
   A battery of ten REDs tells you the tests fail when the source changes; it does not tell you they fail
   for the right *reason*, and a fence that is merely brittle — keyed on a line number, a whitespace run,
   a whole-file digest — produces exactly the same ten REDs as a fence that is discriminating. The
   distinguishing observation is a mutation that alters no behaviour and must therefore stay GREEN. If it
   goes red, the battery has been grading edits rather than behaviour, and every other verdict in it is
   uninterpretable.

   > **Rule.** Every mutation battery carries at least one **no-op mutant with an expected verdict of
   > GREEN**, and the harness compares each mutant's verdict against its *declared expectation* rather
   > than counting REDs. M257x iter-16 ran 11 mutants over the two bring-up verdict fixes: 10 declared-RED
   > (all killed) and 1 declared-GREEN no-op (survived). Also run the unmutated control **after** the
   > battery as well as before — a restore that silently failed otherwise reads as a result.

   **Include the INVERTED mutant, because presence is not meaning.** M257x iter-27's new fence asserted
   that a seeder's share-draw guard was "guarded by a condition referencing `isHero`" — and mutant M4
   flipped that guard from `if !isHero && !memberInShare(…)` to `if isHero && !memberInShare(…)`: heroes
   gated, everyone else free, **the exact opposite semantics**. The fence reported GREEN. It was checking
   that a *token appeared in a construct*, which the inverse satisfies just as well as the original.
   Removal-mutants (M1, M2, M3) all died correctly and could not have found this — deleting the guard and
   reversing it are different edits, and only one of them is what a careless future editor actually writes.

   > **Rule.** For any fence asserting that a guard is *present*, add a mutant that **inverts** it rather
   > than removing it, and expect RED. A fence that a sign-flip satisfies is measuring the presence of a
   > token, not the meaning of a guard — §5 rule 7's family (*a probe must not be able to satisfy itself*)
   > seen from the mutation side. The repair is to assert against the **negation as a parsed node** (here
   > an `*ast.UnaryExpr` with `token.NOT`), not against the identifier's appearance anywhere in the
   > condition. Note which control found it: the battery's declared-GREEN no-op is what made the M4 result
   > interpretable at all, so this rule and the one above are one instrument, not two.

   **And when the fix has two clauses, mutate BOTH — a single-clause mutant proves nothing about a
   conjunction whose clauses are individually sufficient.** M257x iter-30 repaired a page accessor that
   was picking the wrong one of eight matching cards, by adding a structural discriminator **and**
   switching `.first()` → `.last()`. Its mutant battery then read:

   | mutant | expected | actual |
   |---|---|---|
   | M0 no-op (comment) | GREEN | GREEN ✓ |
   | M1 **inverted**: `.last()` → `.first()`, discriminator kept | RED | **GREEN** |
   | M2 removal: discriminator dropped, `.last()` kept | RED | **GREEN** |
   | M3 **full revert** to the pre-fix accessor | RED | RED ✓ |

   M1 and M2 surviving looked, for a minute, like the fix was not attributable — the exact shape of a
   result that should be reported rather than explained away. Re-derived: each clause **on its own**
   selects a node carrying the asserted text, so each single-clause mutant is not a broken fix but a
   *different working one*. The only mutant that reproduces the defect is the one that removes **both**.

   > **Rule.** Before reading a surviving mutant as "the fix does not matter", check whether the fix is a
   > conjunction of **individually sufficient** clauses. If it is, the discriminating control is the
   > **full revert to the pre-fix construct**, and the single-clause mutants are measuring redundancy, not
   > attribution. Run the full revert and require RED; keep the redundant clause anyway (it is what makes
   > the accessor *mean* the thing it names), but say in the code comment that it is redundant today, so a
   > later reader does not delete it believing it load-bearing — or keep it believing it is the fix.

6. **Scope the construct to its BLOCK, or the fence cries wolf.** iter-06's first cut of the write-target
   fence recognised `{"<schema>", "<table>",` and `"<schema>.<table>",` anywhere in a file. It promptly
   flagged 40-odd casbin grants (`{"default", "admin", "org:feature:insights"}`) and the string
   `"clerk.com"`. That is the *same* mistake as asserting on `file.read()`: the regex WAS the construct.
   Recognising them only inside `var resetTables = []string{` and inside a `[]struct{ schema, table string
   … }{` literal removed every false positive without weakening the true one — where the tempting
   alternative, an allow-list, is Trap A's tune-until-it-catches-nothing in miniature. **A fence that cries
   wolf gets disabled, and a disabled fence is indistinguishable from never having written one.**

Static and live are **both** required. Static is the only honest offline check, because every seeder test
asserts against a recording fake `Conn` that accepts any table name — *a fake cannot know a table was
dropped*, which is why 2,617 offline tests passed while the bring-up was broken for four days. Live is the
only check that knows what the migration path actually produced.

---

## 9. Cadence

Detection is cheap. Run it on a schedule, not on an incident.

- **At every release open, and before any prove-it-live milestone:** run §4's six signals.
- **Whenever a bring-up fails oddly:** check signals 2 and 3 *before* debugging the tooling. Three times the
  answer was "the platform moved."
- **Watch the named next fold.** v9.0 folds `storage` + `messenger`. Expect: their compose services removed,
  `app/internal/storage` and `app/internal/messenger` to appear, their ECS to go to zero — and, because their
  env flags are being *deleted* rather than defaulted, anything reading `STORAGE_IN_APP` /
  `MESSENGER_IN_APP_SUBSCRIBER` to break rather than degrade.
- **When M810 lands** and the legacy repos leave the clone set, §2's time bomb fires. Fix the tuple before
  then, not after.

---

## 10. Reading list

- [`platform-migration-status.md`](../architecture/platform-migration-status.md) — the map
- [`platform_repo.md`](./platform_repo.md) — the Makefile, profiles, and `repos.yml`
- [`verification.md`](./verification.md) — pre-flight rung zero (*tagging is not publishing*)
- [`safety.md`](./safety.md) — the safety contract a re-point must not weaken
- [`idempotency.md`](./idempotency.md) — what happens when a bring-up step runs twice

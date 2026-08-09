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
| v2.8 | **storage + messenger + customerio-sync → app** (`838d907`, merged `0c91421`, 2026-08-05) | **a fence caught it, unaided, hours after it landed** — `platform_alignment_guard` assertion B, 2 for 2, on a tree nobody had touched. The first occurrence in this table found by an instrument instead of by a breakage |

The platform wrote its plan down. Read it — it is the single highest-value artifact in this whole area:
`app/knowledge/plan/roadmap-<name>-in-app.md`.

    v2.0  skiller-in-app     shipped
    v5.0  skillpath-in-app   shipped (M507 decommission, PR #1042)
    v7.0  jobsim-in-app      shipped (M701→M710; M710 executing)
    v8.0  cms-in-app         shipped to main (teardown M810)
    v9.0  support-in-app     SHIPPED 2026-08-05 (838d907) — folded storage + messenger
                                                            + customerio-sync, which no plan doc named

**The next two occurrences are already named and dated.** The point of this doc is to stop reacting and start
following.

---

## 2. Why v2.8 was latent — the hand-maintained tuple

This is the mechanism, and it is not what anyone assumed. The assumption was *"a fresh stack never creates the
`jobsimulation` schema, so rext's writes will fail."* **False — because rext creates the schema itself.**

`rosetta-extensions/demo-stack/migrate-demo.sh` **@ `38a4214`** — the pre-fix state this section diagnoses.
**These are historical anchors, deliberately pinned**: iter-02 (`54bccf7`) rewrote all three sites, so at rext
`415240f` every line number below resolves to unrelated code (`:81-90` is now `wait_pg()`, `:106` a `log` line,
`:108` an `ON_ERROR_STOP` comment). Read them at `38a4214`; the "where it lives now" pointers follow each one.

- `:81-85` — `CREATE SCHEMA IF NOT EXISTS` for `extensions`, `sentinel`, `cms`, `jobsimulation`, `skillpath`.
  *(Now `:117-130` @ `415240f`, and the list is **built** by `repos_yml_schemas_to_create` rather than spelled
  out — see `stack-core/lib/repos_yml.sh`.)*
- `:106` — atlas-applies a **hardcoded 4-tuple**: `app:public cms:cms jobsimulation:jobsimulation
  skillpath:skillpath`. *(Now `MIG_PAIRS`, assigned from `repos_yml_migration_pairs "$PLAT_REPOS_YML"` at
  `:41` — rationale at `:20-25` — and consumed at `:141`/`:150` @ `415240f`.)*
- `:108` — `[ -d "$DEV/$r" ] || continue`, gated on **whether the repo directory exists** — it never consults
  `repos.yml`'s `migrations:` flag at all. *(Now `:156-160` @ `415240f`: the silent `continue` became a loud
  `✗ … its clone is ABSENT` plus `mig_fail=1`, so the skip cannot pass as a pass.)*

So rext created the legacy schemas and migrated them out of the still-cloned legacy repos, entirely bypassing
the file that declares the truth. The tuple was **hand-maintained**: the comment at `:95-96` (@ `38a4214`;
`:136-137` @ `415240f`) records that someone edited it when skiller merged (*"skiller merged into app — the
taxonomy tables live in `public` … so there is no skiller repo/schema pair to migrate"*). Nobody edited it for
jobsimulation or cms.

**The time bomb** *(as diagnosed at `38a4214`; defused — see below)*. The legacy repos were kept in `repos.yml`
only *"as the rollback reference until M810."* The day they left the clone set, `[ -d ] || continue` would
**silently skip** them, both schemas would become empty shells, and **13 write targets would fail with 42P01 at
once** — the M257/B1 shape, twice over.

**The canary was already visible** *(also at `38a4214`)*. `skillpath` was still in the tuple but absent from
origin `repos.yml`, so it was never cloned, so the schema was created and left empty. That was harmless only
because rext writes zero `skillpath.*` tables. It was a live preview of what jobsimulation and cms would do.
It is gone at `415240f` — `migrate-demo.sh:112-116` records the removal in as many words.

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
fence rather than deleting the mechanism**: the next entry should have to argue with a fence instead of
landing as a one-word diff.

**That argument was cashed on 2026-08-05, and the emptied list is why nothing happened.** `838d907` removed
`storage` and `messenger` from `repos.yml` — the very *"day they leave the clone set"* the time bomb above is
about. Re-derived across the move, the migration pairs (`app:public`) and the CREATE SCHEMA set
(`extensions · sentinel · public`) are **identical at both refs, and identical correctly**. A hand-maintained
tuple would have silently skipped both repos; the derivation had nothing to skip, because it never named them.
See §9.

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
(`docker-compose.yml:18`, `search_path=sentinel`). So a check shaped *"`migrations: false` ⇒ that schema must
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
    answered **500 after 30.0 s**. The one route the check read was the one route the defect spared, so a
    demo whose landing page was a 500 graded green for four releases.

    > **Correction — it is NOT a public-vs-gated split in the middleware.** This rule used to explain the
    > divergence as *"`/library` is public and short-circuits in Clerk's middleware before the code path
    > that was broken; `/` does not."* **Both routes are public**, in the same matcher: ant-academy
    > `code/proxy.js:139-140` (`"/library"`, `"/library/(.*)"`) and `:170` (`"/"` — commented *"M4 public
    > catalog — the front door"*), all @ `22df69dd8`, and `:282` short-circuits them identically with
    > `if (isPublic(req)) return embedResponse(req, embed);`. The middleware cannot be what separated a
    > 9 ms 200 from a 30 s 500.
    >
    > What actually separates them is the **route group, i.e. which layout renders**: `/` resolves to
    > `app/(authed)/page.jsx`, wrapped by `app/(authed)/layout.jsx:24` `<ClerkProvider …>` plus a
    > `QueryProvider` and six sync providers; `/library` resolves under `app/(public)/layout.jsx`, a
    > seven-line pass-through whose own header states it *"renders ZERO Clerk-aware components so anonymous
    > visitors never hit Clerk's dev-browser handshake"* (`:1-4`). The probe read the one route with no
    > Clerk-bearing layout at all. **The rule is unchanged and the measurements stand — only the mechanism
    > was wrong**, and it was wrong in the direction that makes the probe look better-targeted than it was:
    > "public" was never the discriminator, so "we checked the public one" was never a reason to expect the
    > gated one to behave the same.

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
    files; iter-39 had two find `hiring.md:93` **as it stood then** independently — a HISTORICAL anchor,
    deliberately not re-pointed, because re-pointing it would falsify the record of what iter-39 actually
    found). But a **claim does not respect a file
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

22. **A FROZEN instrument is not a PRECISE instrument — measure the reading's own variance before you
    believe any single reading, especially one that returns zero.** Rules 18/20/21 all decompose a residual
    into a corpus term and an induced term. That arithmetic is sound. What it silently assumes is that the
    reading *measures* the corpus term — and a reading performed by auditors, however carefully specified,
    is an instrument with a variance of its own.

    M257x believed it froze its instrument at iter-41 and never touched a knob again — **a belief rule 25
    below refutes: the briefing was git-ignored and was re-authored every pass.** As believed at the time:
    seven seats, one briefing, one
    size-sorted partition, every file read top-to-bottom under a per-file `wc -l` positive control, plus an
    adversarial diff seat. Three readings were then taken with it, on a corpus changing only by the repairs
    between them: **18 → 7 → 12.**

    **iter-47 read 40 files with seven seats and reported the pre-existing residual as ZERO** — all 7 of its
    findings induced by the repair before it — and concluded the corpus term had converged. **iter-48 read
    the same 40 files with the same instrument and booked 12, of which 10 were NOT induced and 7 predated
    the milestone entirely**, authored between four and five months earlier. Those seven sat, unchanged, in
    seats' own assigned file sets during the reading that reported zero. One was a NOT NULL + UNIQUE +
    undefaulted column missing from a documented "minimal write-set" — a passage iter-47 **read and booked
    as a MINOR**.

    > **The run-to-run variance of the frozen instrument (±5) was larger than the residual it was being
    > used to measure (7). At that point the reading is not measuring the corpus; the corpus is inside the
    > reading's error bar.**

    Three consequences, and the third is the one that changes what you do:

    - **A zero is the least trustworthy reading, not the most.** Every other value is corroborated by the
      findings it names; a zero names nothing and is equally consistent with a clean corpus and a poor
      pass. Treat it as a hypothesis requiring a second independent reading, never as a closed gate.
    - **"Every better instrument found more" is a warning, not progress.** A series that rises whenever the
      instrument sharpens (`25 → 13 → 11 → 17 → 37 → 18 → 7 → 12` here, with the jumps at each instrument
      change) has not been measuring the corpus — it has been measuring reach. Rule 21's classification
      tells you what to build; this rule tells you when to stop believing the count.
    - **Do not write an exit gate as "a reading returns zero" for a reading whose variance you have not
      measured.** It is unreachable-by-construction whenever variance exceeds the residual, and it rewards
      the pass that happens to read least well. Gate instead on something with a floor of zero *by
      construction* — no finding of the classes the fences claim to cover, or two consecutive readings
      agreeing within a declared bound. **Measure the variance FIRST, by reading the same tree twice with
      no repair between; that is a cheap experiment and this milestone paid eight passes to learn it.**

    The corollary to rule 21's *fixture is perishable*: **capture the answer key of a reading that
    contradicts an earlier reading, even when — especially when — nothing will be repaired from it.** It is
    the only artifact that can support the claim *a full multi-auditor pass missed these while they sat in
    its own file sets*, and that claim is about the method, which outlives the corpus.

23. **Rule 22's experiment, actually run: read the same tree TWICE, blind, and treat the two readings as a
    capture–recapture sample. The recall it measures is low enough to change what you do.** Rule 22 says
    measure the variance and calls the experiment cheap. M257x iter-50 ran it, and it was cheaper still —
    reading #9 was already the first half, so the whole design cost **one** reading, in a window that
    existed exactly once, between a reading and the repair of its findings.

    The control was total: the 40 audited files byte-identical (`git diff --stat` empty), all 13
    ground-truth clones at the same shas, the partition therefore dealing the **same hand**, the diff seat
    given the identical diff, and every seat **blind** — fresh, told nothing of a prior reading, barred
    from the plan directory that holds the answer key. **No partition, ground-truth or corpus confound.**
    Whatever the two readings disagree about is the instrument.

    | | reading #9 | reading #10 |
    |---|---|---|
    | blockers | 14 | 7 |
    | per seat | A 1 · B 1 · C 2 · D 2 · E 3 · F 0 · G 5 | A 1 · B 0 · C 0 · D 1 · E 0 · F 0 · G 5 |

    **Matched: 4. Union: 18. Recall: 4/14 = 29% and 4/7 = 57%.** Chapman's estimator puts the tree at
    **~23** blockers — and because heterogeneous detectability biases capture–recapture **downward**, that
    is a **floor**.

    Three consequences, and the third is the one that changes the work:

    - **A single reading is a SAMPLE, not a census.** Two full multi-auditor passes over an unchanged tree
      named 18 findings between them and neither named more than 14. Every count this protocol has ever
      published is a draw from a larger pool.
    - **A "zero" reading is not weak evidence in the way rule 22 implies — it is weak *relative to a small
      residual*.** With recall ≈ 0.43, missing all of a residual of 23 has probability ≈ 10⁻⁵. The problem
      is not that a zero could be luck; it is that **a zero is not drawable while the residual is 23**.
    - **A repair pass can only repair what a reading NAMES.** With recall < 1 and a non-zero induction
      rate, repair-then-read has a **fixed point**, and it sits where the series has been sitting. A series
      that stops falling is not necessarily near zero; it may be at the equilibrium of its own method. Do
      not read a flat series as convergence without measuring recall.

    What to do instead, and it is not more passes of the same shape:

    - **Repair the UNION of two blind readings, never one reading.** Two cost 2× and cover far more than
      2× — 78% of the estimated pool here, against 61% for the richer reading alone.
    - **`N̂` is a metric with a floor of zero BY CONSTRUCTION**, which is exactly the property rule 22 asks
      a gate to have: when there is nothing to find, `m`, `n₁` and `n₂` collapse together.
    - **Record what a reading CLEARS, not only what it finds.** The paired design's most valuable output
      was the list of things reading #10 positively certified that reading #9 had booked as blockers —
      which is how the next rule was found.

24. **An AUDITED ZERO can be wrong, and a wrong audited zero is worse than a silence.** M257x iter-50:
    three independent seats each re-derived the corpus's *"31 of 135 schemas auto-filter by organization"*
    and each recorded it as a positively audited zero. One ran the document's own `comm`/`xargs`
    derivation; one re-computed the split by hand; one *"re-derived independently — every figure matches."*
    **All three were wrong about the SET — and right about the TOTAL, which nobody noticed for two
    iterations.** `schema/organization.go:56` does declare its own `Policy()` with
    `rule.FilterSameOrganizations()` and uses neither mixin, so the audit's *addition* was real. But the
    base was also wrong: `grep -c 'OrganizationMixin{}'` returns 30 and **one of them,
    `user_resource.go:22`, is commented out**, so the live mixin set is **29**. The true count is
    **29 + `Membership` + `Organization` = 31** — the number the corpus already had, reached by two
    compensating errors.

    > **iter-52 repaired the corpus from 31 to 32 on this ledger's authority, and its two blind pre-commit
    > readers independently refuted it.** The audit that corrected an audited zero was itself wrong, in the
    > same direction and for the same reason: it re-derived a *sum* over a *set* it never enumerated. This
    > rule now has three generations of the identical failure, the third committed by the rule's own author.
    > **A `grep -c` over source counts commented-out code.** Enumerate the set, exclude what does not
    > compile, and state the cardinality before the arithmetic.

    Every seat verified **the arithmetic the document showed** instead of **the predicate the document
    claimed**. That is rule 17 — *a count can be exactly right while the claim it supports is false* —
    violated three times in one pass by auditors briefed on rule 17.

    > **A document that shows its own derivation is harder to audit than one that does not.** The visible
    > arithmetic is an attractor: it is checkable, it checks out, and checking it feels like auditing the
    > claim. The incompleteness is never in the arithmetic; it is in the SET the arithmetic ranges over.

    So: when a passage shows its derivation, **re-derive the set from the source, not the sum from the
    set** — enumerate the predicate independently and compare cardinalities. And treat *"I re-derived it
    and it matches"* as the weakest form of clearance a report can contain, because a silence is merely
    uninformative while a wrong audited zero is evidence pointing the wrong way — and a pass reporting
    zero blockers is made of nothing else.

25. **An instrument that is DESCRIBED rather than STORED is not frozen — and a rising series then measures
    the re-authoring, not the corpus.** Rule 22 above records that M257x *"froze its instrument at iter-41
    and never touched a knob again."* **That sentence was false, and no diff could show it.** The auditor
    briefing — *"the whole instrument"*, in its own words — lived in a **git-ignored scratch directory**. It
    appeared in no commit and no iter directory, so every pass "held it fixed" by **re-authoring it from a
    one-line summary in the previous iter's `overview.md`.**

    M257x iter-53 discovered this by doing it: it looked for the briefing inside the milestone directory,
    did not find it, re-authored it, and its two blind readings came in at **32 and 26** where the previous
    pair had come in at **14 and 7**. Re-grading iter-53's union against the recovered canonical rule
    *verbatim* brings it to **23 and 23** — so roughly half the jump was grading drift and half was not, and
    **neither half is a statement about the corpus.**

    The drift was concentrated in one clause. The canonical rule resolved doubt **downward** — *"if you
    cannot cite the refutation, it is not a blocker"* — and carved out **undercount**, **omitted list
    member** and **line drift** as MINOR explicitly. The re-authored rule resolved doubt **upward** (*"when
    in doubt, book it as a BLOCKER"*) and carved out none of them. A grading rule that resolves doubt upward
    cannot produce a number comparable to one that resolves it downward, and no amount of care in the
    reading recovers the comparability afterwards.

    > **Rule 22's warning arrives from the direction it did not anticipate.** It assumed instrument changes
    > would be *deliberate*, and told you to distrust the count when you sharpen the instrument. The failure
    > mode is worse than that: an instrument kept only as prose is re-sharpened, or re-blunted, **every time
    > it is used**, by whoever restates it.

    So: **store the instrument as a versioned file in the repository, and diff it.** Not in scratch, not in
    `.agentspace`, not paraphrased into the next iter's plan. If a measurement's procedure is a prose
    description that the next run re-instantiates, then *"held fixed"* is a claim about nothing, and a
    metric built on it — including one with a floor of zero by construction — inherits the drift. **A
    gate-bearing metric requires a stored instrument before it requires anything else.**

    The corollary for capture–recapture specifically: **recall survives this confound and the count does
    not.** Across two paired experiments at two different grading rules and on two different trees, the
    per-finding detection probability of one 7-seat pass measured 43%, 42% and 48%. Prefer the quantity that
    replicates.

26. **An input that can change without appearing in a diff is not a controlled input — and a measurement
    that does not name its refs is an anecdote.** Rule 25 caught one instance of this and read it as a
    fact about instruments. It is wider. M257x has now been bitten **four times, in its own apparatus**:
    the rext version pin (`.agentspace/rext.tag` — **git-ignored**, so 11/11 clones reported `behind: null`
    while the log claimed *"provably fresh"*); the auditor briefing (**git-ignored**, rule 25); the
    ground-truth platform clone (free to move, and it moved *during* a reading, invalidating a seat's
    clearance by name); and — the one that hurts — **the gate-meeting run of a gate clause, which recorded
    no platform ref at all.** Three were git-ignored files; the fourth was an external ref nobody wrote
    down. Same class.

    So, operationally:
    - **State the refs in the artifact, at the moment of measurement** — platform sha, tooling tag, corpus
      HEAD, instrument-file sha. Not in the surrounding narrative, not by adjacency to the previous
      measurement.
    - **Nothing an instrument depends on may live on a git-ignored path.** Store it; cite its sha.
    - **Re-check the moving ref at close as well as open.** A close-time re-fetch that finds it moved
      **invalidates the measurement by construction** — it does not "probably still hold." The two
      occurrences of M257x's re-scope trigger were both found this way, voluntarily.
    - **The iter that detects the move re-points it, in that iter.** M257x iter-54 absorbed a three-commit
      move in under an hour; the expense is in deferring it, not in doing it.

27. **Derive, else fence, else DECLARE it prose-under-review — and that order is now measured, not
    preferred.** On 2026-08-03 three platform commits removed cms, jobsimulation and roadrunner from
    `repos.yml` and from compose, in one working day. The same event hit all three approaches at once:

    | approach | outcome |
    |---|---|
    | **derived** — the migration/schema sets read from `repos.yml` at runtime | tracked the removal with **zero human action**; reading identical before and after, and identical *correctly* |
    | **fenced** — the map↔`repos.yml` membership guard | caught it **unaided, 3 for 3**, within hours, on a tree nobody had touched — its first non-staged catch |
    | **hand-maintained prose** — the map's narrative sections, and 81 sites across 21 corpus files | **falsified the same day**; one falsehood was written *by the milestone itself*, while updating the map, citing line anchors its own earlier iter had deleted |

    Three approaches, one event, one day, a control in each row. **Before writing a claim, ask whether it
    can be derived; if not, whether it can be fenced; if neither, mark it explicitly as prose-under-review
    with a re-check date.** A document that mixes the three without marking which is which invites the
    reader to extend a fenced row's authority to an unfenced paragraph — which is exactly how a false
    narrative survived beside a guard that was, at that moment, green.

    And the sharpest corollary: **quoting a prior finding forward is not evidence.** The false claim above
    was true when first written and had been fixed two iterations later by the same milestone. It survived
    because it was *cited* rather than *re-run*. Re-measure, or cite the re-measurement.

28. **Three true facts do not make a cause — join them with one experiment.** M257x iter-55 diagnosed a
    `backend` container that exited 0 in silence. It measured, correctly and separately, that (a) the
    container exits 0, (b) `STORAGE_RPC_ADDR` is absent from its environment, and (c) the pinned `app`
    source reads that variable three times. It then reported the conjunction as the cause and routed a
    **version-pin advance** for it — the single most dangerous move in that milestone's history, the one
    that had broken the seeders twice.

    iter-56 refuted it with two `docker run`s against the same image on the same network. The same
    environment with **no mounts** starts fine and serves on `:8082`; the same environment **plus the
    host's `$HOME/.aws/credentials` bind mount** reproduces the dead 2-line signature exactly. The env var
    was never the cause. And the destination was unreachable anyway: `app` at **origin/main IS the newest
    tag**, and it reads the variable at the same three sites — so no advance could have restored it.

    The diagnosis was not sloppy. Every input was measured; it explained the exit, the silence *and* the
    137 ms timing; and it named a real platform inconsistency. It was simply not tested. This is the cheap
    half of the `D-M257x-13` correction (*a mechanism that explains the observation is not the mechanism
    that produced it*): when a diagnosis has the shape **"X is missing AND the code reads X, therefore X"**,
    supply X, or remove the other suspect, and watch. It usually costs one command, and the alternative
    here was a cold cycle plus a pin advance aimed at a release that does not contain the fix.

    **Corollary — check that the proposed remedy contains the fix, before taking the remedy.**
    `git rev-list --count <newest-tag>..origin/main` and one `git grep` at the target ref would have shown,
    in ten seconds and before any decision, that the advance was a no-op against the stated cause.

29. **A reading names INSTANCES; only a derivation can name a PREDICATE.** This is why the M257x
    ten-reading series had a fixed point rather than a slope. A seven-seat blind reading of the same tree
    recalls **43–48%** of the sites present (iter-50's paired experiment); it can find *this* wrong
    sentence and *that* one, but it has no way to say *"every sentence in the tree that assumes a `graphql`
    profile exists."* The residual it leaves is therefore not a shrinking pile — the un-named sites are
    re-drawn from the same distribution on every pass, while the platform adds new ones faster (one working
    day of commits produced **81** fresh sites across 21 files). Net repair rate over ten readings: **−72**.

    The three largest residual classes turned out to share **one false predicate each** — *a `graphql`
    profile exists* (26 live docs), *messenger reaches cms/jobsim at husk containers*, *this line number
    names this construct*. **Repair by PREDICATE, not by claim** — this extends rule 19 (*by claim, not by
    file*) one level down. A predicate whose legal set is derivable from a platform artifact is a rule-27
    row-one candidate: derive the set, fence it both ways, and the class closes in one build instead of
    N edits — and stays closed, which no edit does.

    **A reading is still how you find out you were wrong.** Predicate-scoping changes the unit of
    *repair*; it does not replace the reading, and a residual is still graded by a reading that returns
    zero.

30. **Grade a documented command on "does it still SELECT something", not "does it still parse."** The
    dominant failure mode after a rename is not an error — it is a **successful command that does nothing
    visible**. And the measured version is worse than the reported one. `docker-compose.yml` at platform
    `0dab54d` gives `postgresql`, `redis` and `sentinel` **no `profiles:` key at all**, so they belong to
    *every* selection. A `--profile` naming the retired `graphql` token therefore exits **0** and starts
    **three containers**: Postgres answers, `docker ps` is non-empty, and the application is simply
    absent. The corpus had recorded this as *"starts zero containers"* — "zero"
    would at least be unambiguous; **three** presents as a partially-working stack and sends the reader
    debugging the application instead of the invocation.

    Two corollaries, both learned the same afternoon:

    * **Enumerate the always-on floor first.** It is the set that makes a dead token look alive, and it is
      derivable: *the services declaring no `profiles:` key*. Resolving compose's `include:` is
      load-bearing here rather than tidy — two thirds of that floor lives in `common.yml`, so a parser
      reading only `docker-compose.yml` computes `floor = {sentinel}` and concludes, wrongly and quietly,
      that a dead token starts nothing.
    * **Do not spell a dead command in runnable form, even to warn about it.** A warning that contains a
      copy-pasteable `PROFILE=<retired>` is indistinguishable, to a reader in a hurry and to a fence, from
      an instruction. Name the token; do not write the invocation. (iter-60 hit this on its own repair
      text: the fence went RED on the corrections, which was the fence being right.)

31. **A REFUTATION is a measurement, and it expires exactly like the claim it refuted.** M257x iter-22
    correctly refuted a proposed correction: at `2adcf71`, `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR`
    really did still address the husk containers, and applying the "fix" would have replaced two true
    statements with false ones. That refutation was written up as standing guidance — *"That address is
    **CURRENT, not stale text**"* — and by `0dab54d` the M809 re-point had landed and inverted it.

    A refuted-correction note is **more** dangerous than an ordinary stale claim once it expires, because
    its emphatic anti-repair language reads as *already adjudicated* and survives readings that would have
    caught plain prose. **Pin a refutation to the ref at which it was taken, in every sentence that depends
    on it**, or it will be read as standing. The mechanical test: if removing the ref would not change how
    the sentence reads, the ref is decoration and the sentence is a standing claim.

32. **Re-derive the hand-off's numbers — including the orchestrator's.** Rule 1's oldest lesson does not
    stop applying because the sender is trusted. Two consecutive M257x iterations corrected an
    orchestrator-supplied fact: one reported *"demo-1 GONE — 0 containers"* when the Docker daemon was
    merely down and all 11 were present; the next handed down *"17 files / 30 occurrences"* for a class
    that measured **26 live docs / 56 lines**, and `cmd/academyImport/main.go:235` for a `Getenv` that is
    at `:231` (@ `app` `b948604` v1.366.0 — `:235` is the `is required` return; both true, different
    lines, and only one of them is the read site. The storage fold has since deleted **both**: at `app`
    `9d00a313` v1.367.0 that file names `STORAGE_RPC_ADDR` nowhere, which is why this clause carries a
    ref and rule 33 is the reason it must). None of the three was careless. Re-derivation costs one
    command; inheriting costs a milestone's credibility.

33. **A ref-pin is a DATE, not an exemption — and it exempts a CLAIM, never a neighbourhood.** Every
    fence in this family skips a claim that names a ref, on the sound reading that it was true when
    measured (TOK-04 P1). M257x iter-62 promoted the hole that creates; iter-63 measured it and found
    it is **three mechanisms wearing one name**, of which only the third is the one people describe:

    (All three instances are quoted **as they stood before iter-63's repair**; the lines have since
    been rewritten, so they are named by their content rather than by an anchor that would now
    resolve to the fix — rule 34's own trap, avoided rather than described.)

    | mechanism | measured instance |
    |---|---|
    | the pin crosses a **row** boundary | `shared_libraries.md`'s colony **Version pin** row dated itself *"at platform `2adcf71`"*; the **Imported by** row directly beneath it — a different claim, about which profile starts which containers — went silent with it |
    | the pin crosses a **cell** boundary | `service_taxonomy.md`'s CMS row: one row, two clauses. Cell 2 cited `915da06` for the supergraph count; cell 3 claimed the `cms` container still starts in the retired router's profile |
    | the pin is cited **as evidence of currency** | `service_taxonomy.md` headed its Services table *"(**current** local docker-compose @ platform `2adcf71`)"*; `messenger.md`'s `CMS_RPC_ADDR` row wrote *"**current, not stale**"* beside its own pin |

    And the inverse nobody had checked: a pin naming **the guard's own ref** exempted too, so a doc
    could immunise a false present-tense claim by citing the very commit that refutes it.

    The rule, all four parts derived from structure rather than tuned: **a pin's scope is the claim's
    own block — a markdown CELL in a table, a wrapped sentence in prose; a pin naming the ref the
    checker derived from exempts nothing; and a block that asserts currency (*current*, *not stale*,
    *today*) cannot be pinned into silence, because the pin and the word are making opposite claims
    and only one of them is what the reader acts on.** Fenced in
    `stack-core/platform_predicate_guard.py` (`_pin_exempts` / `_pin_window`), each part watched RED
    under its own inversion mutant. **The blast radius is small and that is the point** — the rule
    removes the exemption exactly where the document itself insists the claim is current.

34. **Line numbers move when YOU edit too — §7 rule 4 applies to a corpus repair, not just a pin
    advance.** The corpus cites itself by `<doc>.md:N` in ~200 places. A repair that changes a file's
    line count silently invalidates every citation into that file below the edit — the same mechanism
    as an advancing pin, one level in, and with the same catch rate. M257x iter-63's own repairs moved
    **9 intra-corpus citations across 6 files** (two of them in root `CLAUDE.md`), and
    `anchor_construct_guard` caught exactly **one** of the nine: the one that happened to land on a
    blank line. The other eight landed on content and read as correct. **Re-derive the line map from
    `git diff -U0` and re-point in the same commit** — it is a five-line script and it is not
    optional. Two cautions worth the sentence: the re-point is **not idempotent** (the map is computed
    against `HEAD`, so a second run double-applies), and a citation into a line the edit *replaced*
    needs a human, because "where did that content go" is not a question a diff answers.

35. **A count that grows in the same pass that REACH grows is not a regression until you SPLIT it by
    the reach dimension — and rule 16 has a mirror image.** Rule 16 says an unread metric is
    indistinguishable from an unmoved one. The mirror: **an unread metric that becomes read looks
    exactly like a metric that got worse.**

    M257x iter-74. iter-73 widened a citation resolver's reach (124 → 177 anchors) and the
    `ambiguous` bucket went **12 → 39** in the same run. Routed as *"grew — is this a corpus-writing
    habit worth changing, or a fence limitation?"*, which is a real question and the wrong first one.
    The first one is decidable and costs one derivation: partition the class by **the dimension reach
    moved along** — here, which regex alternative matched the citation.

    | ref source | newly-reachable partition (`bare-code`) | pre-existing partition (`path`) |
    |---|---|---|
    | ambiguous | **27** | **12 — unchanged, to the citation** |

    All of the growth sat inside the partition that **could not be counted at all** the day before.
    The corpus had not moved by one site. A class that grows *entirely inside the newly-reached
    partition* is the instrument improving; a class that grows *inside the old partition* is the
    corpus moving — **and those need opposite responses**, which is why the split comes before any
    repair. Two of the last three routed backlogs in this milestone were routed by pattern-match on a
    count and collapsed on adjudication (64 → 5, 23 → 1).

    **The corollary that made the same iteration worth twice its size: when the split says "the
    instrument", read the class anyway.** The 39 were not all legitimate. **21 of them sat in one
    document**, which turned out to be a property of the *fence*, not of the document: the window it
    called "the block" walked to the nearest **blank line**, and a markdown table has no blank lines
    between its rows — so a citation inside a table took **the entire table** as its window, and every
    sha named in any row chose the ref for the citations in **every** row. Measured instance:
    `external_services.md`'s provider table pins a ref in the **Anthropic Direct** row, and the
    citations in the **AWS Bedrock** and **Mistral** rows — three different providers — were being read
    at it.

    That is rule 33's own *"the pin crosses a ROW boundary"* mechanism, and rule 33 had **already
    ruled on it** (*a markdown CELL in a table, a wrapped sentence in prose*) and was **already
    implemented** — in the sibling guard. So the last part of the rule, and the cheapest of the three:

    > **When a rule is already derived, check every implementation of it, not the one you are
    > editing.** Two guards holding two definitions of the same construct is not a disagreement
    > anyone will notice: each is internally consistent, each is green, and the one that is wrong is
    > wrong *silently*. `grep` for the predicate's name across the whole tooling section before
    > assuming the rule is in force.

    Two smaller cautions from the same repair, both cheap and both earned the hard way. **The first
    draft of the fix contained the defect it removes**: the row test was copied verbatim from the
    sibling and anchors at `|`, so a **blockquoted** table (`> | side | … |`, 75 rows across 14 files
    in this corpus) failed it and fell straight back into the prose branch. And **an ambiguity is
    only worth repairing if it can change an answer** — classifying each residual citation at *every*
    ref its own window names showed **19 of 20 agree** and the 20th is a cell that names the ref it
    asserts in words. Counting a fallback is coverage; assuming it is a defect is not.

36. **A universe built over CLONES must say which copy is the witness — and a derivation that
    returns the convenient answer earns a positive control before you believe it.** Two halves of
    one iteration (M257x iter-75), both cheap, both paid for the hard way.

    **The witness half.** Resolving a bare `<name>.<ext>:N` citation needs an index of *"which files
    exist"*, and the honest way to build one is `git ls-files` per clone — the repository's own
    answer, immune to build output and untracked scratch. But `rosetta-extensions` is cloned
    **twice** in this tree: the per-stack consumption copy (pinned at a tag) and the authoring copy.
    Both directories carry the same **name**. An index keyed by `<clone-name>/<relpath>` collapses
    them; one keyed by absolute path splits them — and iter-75 wrote one of each, so **the same
    misunderstanding produced two different wrong answers inside one iteration**: an adjudication
    reporting a file in *"2 places"* while printing the same path twice, and a dry run silently
    finding 31 of 77.

    > **Two clones of one repo are one witness.** Decide which copy is authoritative *before* asking
    > whether a name is unique — and derive that from something already in the code (here,
    > `resolve()`'s pre-existing preference for the authoring copy) rather than inventing a
    > preference to make the numbers work.

    The rule this protects is worth stating too, because it is what makes such an index safe at all:
    **resolve only on UNIQUENESS, and count what stays unresolved.** `main.go` is 57 tracked files
    in this platform, `main.tf` is 10; `studioManager.go` is 2 — the merged copy and the standalone
    husk, exactly the pair a directory guess gets wrong and exactly the fold the map exists to
    document. 26 sites stayed unresolvable and were named. **An unresolvable citation that cannot be
    resolved without guessing is coverage, not debt.**

    **The control half.** The dry run came back *77 newly resolvable, 0 findings* — and 0 was the
    **surprising** answer, because the comparable widening one iteration earlier had turned the
    corpus RED with 6. A 0 from a pipeline that *cannot* report a finding looks identical to a 0 from
    a clean corpus. So the same code path was fed three known-bad inputs against a real 2,693-line
    target (`:99999` → out-of-range, a blank line → on-blank-line, `:1` → clean) before the 0 was
    written down.

    > §5 rule 2 is usually read as a rule about *searches*. It applies to **derivations**: when a
    > measurement returns the answer you were hoping for, run the control that would have caught the
    > instrument being broken. It cost one command.

    And the adjudication itself, which is the reason the iteration was small: the class had been
    routed as *"92 unrepaired citations"* and adjudicated to **0 defects · 77 unreachable · 26
    undecidable** — the **fourth** consecutive routed count in this milestone to collapse when
    someone finally derived it (64 → 5, 23 → 1, 21 → 0, 92 → 0).

37. **A vocabulary derived from CURRENT state is blind to removals — and removals are the drift.**
    M257x iter-77. `platform_predicate_guard`'s repo vocabulary was `set(repos) | set(compose.services)`,
    derived entirely from what exists **now**. So a repo left the vocabulary at the same commit it
    left `repos.yml`, and every corpus claim naming it became unreadable **at exactly the moment it
    became false** — the worst possible instant for an instrument to stop looking.

    Live, not theoretical: `setup_guide.md:504` enumerated the `migrations: true` repos as
    *"(currently: app, cms, jobsimulation …)"*; the name-resolver dropped `cms` and `jobsimulation`
    as unknown tokens, compared `{'app'} == {'app'}` and **passed a false claim.** The one migration
    claim of 24 the fence could reach was the one it read wrong: *effective reach 0, reported as 1.*

    > **Where the artifact is under version control, derive the vocabulary from its HISTORY.**
    > Here: every `- name:` ever written in `repos.yml`, across all 9 commits that touched it — **14
    > ever, 6 now, 8 removed.** A clone too shallow to answer must report `UNMEASURED` in its reach
    > line and say so; it must never degrade silently to the current set while implying coverage.

    Generalises past repos: the identical shape hides a deleted compose service, a retired profile
    token, a dropped env var. **Any fence keyed on a name list is only as historical as that list.**

38. **A discriminator that NARROWS a fence is evaluated after the verdict, never before it.**
    M257x iter-77, measured on the guard's own corpus. G9's ambiguous-subject rule (the citing block
    names a repo other than the document's own subject) is correct and necessary. Placed **before**
    grading it cost `storage.md:25` — a true, gradeable citation whose sentence merely mentions two
    other repos in passing — and the graded count fell **4 → 2**. Placed **after**, it can only fire
    on a citation that was already about to be reported, so it converts would-be false positives
    into declared UNREACHED and **spends no recall at all.**

    > The loss is invisible where you would look for it: a claim removed from the denominator does
    > not appear as a missed finding, it appears as *"nothing to check."* Same rule, two positions,
    > opposite instruments — and only the reach line can tell them apart.

39. **Committing is not pushing** — the sibling of rung zero, and it fails one step earlier.
    Rung zero (`verification.md`) says *tagging is not publishing*: a stack clones
    `rosetta-extensions` from **origin** at a pinned tag, so a tag living only in the local
    authoring copy is unreachable to it, and M236 lost an entire iteration to exactly that.

    M257x iter-77 opened on the same failure one step up the chain: `main` held **13 commits on a
    single disk** — 8 hardening passes and 5 fences, ~1,400 lines of guard work. Nothing a stack
    consumed was at risk, and **that is why it survived thirteen commits**: the FATAL pin guard
    checks what a stack *pulls*, and nothing checks what the author has not *pushed*. A pin verified
    on origin is entirely silent about the branch it was cut from.

    > **Push before you build on it, and verify with `git ls-remote`** — the same one-command proof
    > rung zero already demands of a tag. `git log --oneline origin/main..main` is the check; a
    > non-empty answer at the end of a session is a single point of failure, not a to-do.

40. **A repair UNIT is not a repair POST-CONDITION — and "the predicate is discharged" is a
    measurement, not a report.** Rule 19 says *repair by claim, not by file*; `D-M257x-59-1` extended
    it to *repair by predicate, not by claim*. Both settle **what you work on**. Neither says **when
    you are done**, and M257x iter-81 paid for the gap: it repaired eleven predicates across 33 files
    and reported all eleven **discharged**, while a sentence asserting that the retired profile token
    had *survived the rename* and *become the default* — a member of P4, booked as **B1** by *both*
    blind readings and **upheld** by the adjudication — stood untouched in
    `corpus/services/graphql-wundergraph.md` (at `:13` as of rosetta `8d6bb6c`; the line moves when
    the claim is repaired, which is the point).

    > **The false sentence is described here, not quoted.** It was reproduced verbatim in this rule
    > until M257x iter-86, when the first full-family guard run measured what that cost: the
    > quotation is itself a published site of P4, and `platform_predicate_guard` booked **this
    > document** for it — the protocol doc tripping the fence it exists to teach, for the second
    > time (`:1305` was the first, iter-84). The rule this file already states about retired
    > compose tokens applies to the file stating it: **do not spell a dead claim in a form that
    > reads as a live one.** A worked example does not need a copy-pasteable copy of the defect.

    **Measured at iter-83, against the repair's own input ledger:** of 147 gradeable booked findings,
    **109 landed inside a repair hunk (74.1 %)** and **38 did not** — of which **35 were in files the
    repair opened and edited**. So it was not a partition gap (only 3 misses were in unopened files,
    all 3 outside the read's own file set) and not estimated membership (the misses fall on predicates
    whose site counts were *exact* as readily as on the ones marked `~`). The single case that names
    the mechanism: in **one file**, the repair rewrote the passage claiming that `make up` starts the
    retired token's profile — a finding the adjudication had **REJECTED** — and left the
    *"survives in compose"* passage 164 lines above it, which the adjudication had **UPHELD**.

    > **The discharge criterion was *"I have swept this file for this predicate."*** Not *"no member
    > survives"*, and not even *"every booked member is fixed."* Nothing anywhere could report the
    > difference between a predicate that was discharged and one that was believed to be.

    Three consequences, and the third is the one that generalizes furthest:

    **(a) A discharge verdict is only as good as its post-condition.** A repair states, per finding,
    that it was **written** or that it was **dispositioned in writing**. A finding may be declined —
    it may not be declined silently, because a silent skip and an omission are indistinguishable
    afterwards, and at 38 sites they were.

    **(b) Fence the direction nothing else fences.** Both existing repair fences are keyed on **what
    the diff contains** — `repair_postcondition.py` on the tree the commit produced,
    `repair_leak_guard.py` on the prose it removed — and are therefore blind *by construction* to a
    finding the repair was handed and never opened: nothing was removed there, nothing was added
    there, and the site reads as it did before. Only the **input ledger** can see it.
    `repair_reach_guard.py` (`FENCE-M257x-iter83-repair-reach`) is that third question, watched RED on
    a real answer key. **Reach is necessary, not sufficient** — a green reach report says the repair
    opened everything it was given, never that it was correct.

    **(c) The forgettable class is the one that gets forgotten.** `repair_leak_guard` goes RED on
    iter-81's commit and **was not run**, because it declares `FENCE_KIND = "standalone"` and the
    DERIVED registry selects only `postcondition`-kind fences — **10 of the 14 guards standing at that
    repair were `standalone`** (4 `postcondition`; 11 of 15 once iter-83 adds one, and the census is
    derivable in one `ast` pass rather than counted by hand — which is how this very figure was
    caught wrong, see below). A
    repair choosing its guard list by hand is §2's hand-maintained tuple wearing a new hat, inside
    the machinery built to end it. **If a guard only runs when someone remembers it, it is not a
    fence; it is a habit.**

    Finally, on the numbers a repair-then-re-read loop produces: **check the units before
    subtracting.** iter-82 reported readings of 29 and 30 against a union of 41 and an overlap of 15
    — `29 + 30 − 41 = 18`, not 15. It was not an arithmetic slip: 29/30 count blocker **blocks** and
    41 counted distinct **anchors**. Re-derived consistently the figures are **28 / 29 / union 43 /
    overlap 14**, and the correction matters because the per-pass **recall** estimate is computed from
    the overlap, and recall is what decides whether a future **zero** means anything.

41. **A check that resolves against a REMOTE-TRACKING ref is only as current as your last fetch — and
    it reads GREEN until you fetch.** §3's corollary said a freshness check comparing to a *pin*
    cannot detect a stale *clone*. This is the same hole one level in, and it is worse, because the
    check that goes blind is the one you would use to find the drift.

    M257x iter-87. `anchor_construct_guard.resolve()` and `platform_alignment_guard.cited_text()`
    both default to `CITE_REF=auto`, whose ladder is **`origin/main` first**, checkout second — a
    deliberate and correct fix (iter-68) for the demo stack pinning its clones to a build tag while
    the exit gate names *origin HEAD*. The consequence nobody had written down: those guards read
    the **remote-tracking ref in the local clone**, which only moves when someone fetches. The
    platform had pushed two commits taking `docker-compose.yml` from 271 lines to **186**; the
    hand-off measured the family at **13 GREEN · 0 RED**. Re-measured minutes later at the
    **identical checkout**, after a `git fetch`: **10 GREEN · 3 RED**. Nothing about the corpus or
    the checkout changed. The fetch did.

    So the two halves of "advance a clone" are not one act, and they have opposite cost profiles:

    | act | what it changes | who reads it |
    |---|---|---|
    | **fetch** | the remote-tracking ref | every citation assertion (`CITE_REF=auto`) |
    | **checkout** | the working tree | only what reads the tree — for this family, `platform` alone |

    > **Rule.** **Fetch every clone in any iter that takes a measurement; advance a clone's CHECKOUT
    > only when the checked-out tree is an input to a derived legal set or to a build.** A fetched
    > clone is already graded at origin HEAD no matter where its HEAD sits.

    Two consequences worth having in advance. **A large "behind" count is not a large repair.** At
    iter-87 `app` was **93** commits behind and carried **65** corpus citations — and because every
    one was already being resolved at `origin/main`, the whole exposure showed up as **2** RED
    anchors. The wave people budget an iteration for is usually already in the last reading.
    **And a guard run against an unfetched clone is not a weaker measurement, it is a different
    one** — it silently grades the corpus against whatever the last fetch happened to capture, and
    reports the result in the guard's own confident voice (§5 rule 12). **Fetch, then measure, and
    say when you fetched.**

    ### 41a. Corollary — a reading's GROUND TRUTH INCLUDES THE CLONE REFS, so **no lane may fetch while a reading is in flight** (M257x iter-101 / iter-102)

    Rule 41 says *fetch, then measure*. It does not say **when you may stop fetching**, and that
    omission cost M257x a provability it cannot get back.

    M257x runs three lanes concurrently against one checkout. Path ownership was assigned to keep
    them from colliding: one lane owned `corpus/**` and the milestone records, another owned
    `stack-demo/**`, a third owned the deferral ledger. **The reading's adjudicators were grading
    claims against the platform clones — which live inside `stack-demo/**`, the tree assigned
    exclusively to another lane.** That lane refreshed the clone set at **11:18:16 – 11:20:51**;
    the reading's adjudication commit landed at **11:21:55**.

    The exposure is bounded — most of adjudication ran pre-fetch, the orchestrator observed the
    pre-fetch ref, no corpus file moved, and the reading's `N` stands. **But it cannot be PROVEN that
    no adjudicator read post-fetch**, and that is the whole problem: a fetch moves the very refs the
    citation guards resolve against (`CITE_REF=auto` → `origin/main` first). Five clones advanced in
    that window — `app` by 98 commits over 634 files, `next-web-app` by 41 over 192 — so an
    adjudicator who happened to re-derive after 11:20:47 graded a **different subject** from one who
    re-derived before, with nothing in either report saying which.

    > **Rule.** **A reading's ground truth is not just the corpus — it is the corpus PLUS every clone
    > ref the reading resolves against.** Freeze both for the reading's duration. **No lane may fetch
    > any clone while a reading is in flight**; a lane that needs a clone advanced says so, and the
    > fetch happens **between** readings and is recorded in the next reading's ground-truth sheet.

    **Path ownership was necessary and not sufficient, and that is the generalisable half.** Two
    lanes can hold disjoint sets of *writable paths* and still collide, because **the reading's
    SUBJECT is wider than the paths anyone declared** — it reaches into a tree the reading never
    writes and does not own. When you partition work by ownership, partition by **what each task
    READS to settle its claims**, not only by what it writes. An instrument whose inputs are outside
    every declared boundary is an instrument nobody is protecting.

    Two practical corollaries:

    - **Timestamp both sides.** A reading's ground-truth sheet already records the clone shas; it must
      also record the **fetch time** of each clone, so a mid-reading move is *detectable* rather than
      merely *suspected*. `guard_family.py` already prints `fetched Nm ago` — the sheet should carry
      the same fact.
    - **"It has since been fetched" is not a repair, it is a NEW ground truth.** The next pass
      re-derives against the moved refs and measures what the move injected; it does not
      retro-fit the old reading. Retro-fitting a reading to refs it never saw is the same error as
      re-anchoring a deleted fact (§5 Trap A).

    #### What 41a CAN and CANNOT enforce — a rule believed enforced is worse than a disclosed gap (M257x iter-103)

    41a binds **lanes**. It does not, and cannot, bind the tooling: **`ensure-clones.sh` runs
    `git fetch` on every bring-up as a freshness assertion, and there is no flag that suppresses it.**

    What that fetch moves is bounded, and it is the bad half: **`refs/remotes/*` only, never a working
    tree** (`DEMO_ADVANCE_CLONES` defaults to `0`). But `refs/remotes/*` is exactly what a citation
    guard resolves against — `CITE_REF=auto`'s ladder is `origin/main` first — so the one thing the
    fetch *does* move is the one thing 41a exists to freeze. Measured: one such fetch caught
    `next-web-app`'s `origin/main` advancing **4 commits** past the frozen ref mid-run.

    > **So state 41a honestly.** A reading may forbid *lanes* from fetching. It may **not** assume no
    > fetch occurred. **If a reading overlaps a bring-up, the reading RECORDS the fetch and treats the
    > affected refs as MOVED** — it does not assert the rule held.

    The corollary above — timestamp both sides — is what makes that detectable rather than merely
    suspected, and it is therefore **not optional**. iter-103 is the worked example in the other
    direction: every clone's HEAD, `origin/main` **and** fetch timestamp were recorded at the reading's
    open and re-read at its close, and all three were identical for all thirteen clones. That reading's
    ground truth is not *believed* frozen, it is **measured** frozen at both ends — which is the only
    form of the claim worth making.

    **A rule that is believed to be enforced but is structurally unenforceable is worse than a disclosed
    gap: the belief suppresses the check that would have caught the violation.**

42. **A RED summary line must name the EVENT, and "the last line of the output" does not.** The
    sibling of rule 11, found in the runner built to enforce rule 8. `guard_family.py` reported each
    guard's final output line as its headline. Guards emit findings in assertion order, so the
    headline was whichever assertion sorted last.

    Measured at iter-87: `platform_alignment_guard` went RED with **21** findings whose **first two**
    were `[B departure] the map claims messenger is in repos.yml, and it is not` and the same for
    `storage` — **two services leaving the platform's clone set, the precise event the fence exists
    to catch, and its first unaided catch of that event.** The family view showed
    `[F out-of-range] gotenberg: cites docker-compose.yml:268`, a citation nit from the bottom of
    the list. Both lines are true. Only one is the news, and the summary view — the one that claims
    to speak for the whole family — showed the other.

    The repair is **state how many, then show the first**, both derived from the guard's own output
    (a curated pick would be the runner deciding which finding matters, which is the hand-maintained
    list of §2 in a new costume). **When a summariser must choose, make it choose by the producer's
    own ordering, and always print the cardinality** — a single finding shown without a count reads
    as the whole verdict.

43. **A mitigation keyed on a service NAME dies when the service does — and its tripwire dies with it,
    skipping rather than failing.** The pairing is what makes this expensive: the fix and the check that
    guards the fix share the same stale key, so they fail together, in the quietest available direction.

    M257x iter-88. `docker-compose.yml` binds `$HOME/.aws/credentials` into a container; on a fresh Linux
    box that path does not exist, **Docker auto-creates it as an empty DIRECTORY**, and the AWS SDK opens
    it (opening a directory succeeds) then fails `EISDIR`. The container prints its full cobra usage block
    and exits 1 — a symptom misread for an entire release cycle as a missing `serve` subcommand. M217
    fixed it correctly, in the generated override, with `if name == "jobsimulation":`.

    Then `d11a403` deleted the `jobsimulation` service and `838d907` moved the identical bind onto
    **`backend`**. The hazard migrated to the stack's most important container; the mitigation stayed
    pointed at a name that no longer resolved. And the tripwire written to catch exactly this looked the
    service up first, did not find it, and called `skipTest("jobsimulation not in the compose")` —
    a skip, which reads exactly like a pass (rule 8). Its sibling *passed* while asserting *"exactly 1
    `$HOME` bind (jobsimulation's AWS creds)"*: the count right, the claim inside it false (rule 17).

    Two more instances of the same shape surfaced in the same sweep, which is what makes it a rule rather
    than an anecdote: an anti-vacuity assert reading `len(repos) >= 5` (the platform went to **4**, so a
    guard against vacuity failed for the one reason it must not be sensitive to), and a demopatch anchor
    check that restated its target path beside the manifest that declares it — the manifest was
    re-pointed at M254, the restated copy was not, and the check **skipped for four releases**, which is
    the precise failure its own docstring said it existed to prevent.

    > **Rule.** Key a mitigation on the PROPERTY that made the service special — here *carries a `$HOME`
    > bind* — never on its name. The property outlives the fold; the name is exactly what the platform has
    > been deleting for three releases. And **a check that looks its subject up must FAIL when the subject
    > is absent but its world is present**: "no clone on this box" is a real skip, "the clone is here and
    > the thing is not in it" is the drift.

    **The suite is not fenced, and that is where these live.** Every fence in this family watches the
    corpus or the platform; nothing watches the *tests*, so a hand-maintained platform constant inside a
    test suite is the least-observed place in the system. Two of the three instances above sit in test
    files, and one of them — an assertion that the generator's source still contained the dead literal —
    would have **failed anyone who tried to remove the defect**. When you sweep for stale platform
    constants, sweep the checks too, and read the skip count before the pass count.

And: **verify a claim before escalating it, including a claim made by an audit.** In M257x two probes
contradicted each other on whether `public.sessions` exists; measuring settled it (it does not — created then
dropped as a rename completed) and *inverted* the risk assessment that had been built on it.

44. **"`git grep` at a named ref" is necessary and NOT sufficient — name the tree AND its ref, PER TREE.**
    A ref alone does not make an instrument honest, and here the obvious fix is the worse one. Three
    mechanisms hide a tracked file, and they defeat different tools:

    - **Gitignored-but-tracked.** The shell's `grep` is a function over `ugrep -G --ignore-files`, so it
      skips any tracked file an active `.gitignore` matches. `git grep` sees these; bare `grep` does not.
      Measured across the clone set: **12** such non-empty text files — three times the 4 first counted,
      because `git check-ignore` needs **`--no-index`** to report *tracked* paths at all, and because the
      first census did not descend into the nested repos below. The census missed them by the very
      mechanism it was measuring.
    - **NUL-bearing source.** `grep -I` and `git grep` **both** skip a file containing NUL bytes, and
      `file(1)` calls it "data". Two exist: `next-web-app/apps/web/src/hooks/useCoursebuilder.ts`
      (50,433 bytes, **1** NUL) and `ant-academy/code/src/time/store.js` (15,307 bytes, **1** NUL).
      **ONE byte is enough** — that is the point, and it is why this is a trap rather than a curiosity.
      ⚠ This bullet read *"1,178 NULs"* until M257x iter-98. **1,178 is the file's LINE count**: it was
      produced by `grep -c $'\x00' <file>`, where `grep -c` counts matching *lines*, not bytes, and the
      zsh `$'\x00'` pattern degenerates to the empty string and so matches every line. **The rule about
      lying instruments was itself written with a lying instrument**, and the adversarial seat caught it
      in two independent readings. Count NUL **bytes**, and only this way:
      ```bash
      tr -dc '\000' < FILE | wc -c      # bytes. `grep -c` here counts LINES and the pattern matches all of them
      ```
    - **Nested untracked repos — where the naive fix is worse than no fix.** `stack-demo/app/studio` and
      `stack-demo/cms/studio` are each the `anthropos-studio-room` repo (177 files, own HEAD `aeec036`),
      hidden from their hosts by `app/.gitignore:79` and `cms/.gitignore:129`. `git -C app grep <anything>
      HEAD -- studio/` returns **0 for every predicate** — a guaranteed zero that reads like evidence.
      Only the nested repo's own ref sees them. On the `mistralai` predicate the three instruments
      returned **1 / 0 / 22**; the ref-named `git grep` was the one that scored **0**, and that false
      clearance minted an "imported nowhere" claim that stood in **four** documents until M257x iter-96.
      Worked example: [`studio-room.md`](../services/studio-room.md), the requirements callout.

    So: **enumerate the nested repos before claiming any tree-wide zero**, grep each at its own ref, and
    say which trees the number covers. A tree-wide zero that does not name its sub-repos is an unproven zero.

    ```bash
    # 0. enumerate every git tree, nested ones included — this is the search set
    find stack-demo -name .git -maxdepth 4 | sed 's|/\.git$||' | sort
    # 1. grep each at ITS OWN ref (never the host's). `grep -ci` prints `file:count`, so SUM the counts
    #    for lines and pipe to `wc -l` only when you want FILES. Keep -i: casing is not the predicate.
    for d in $(find stack-demo -name .git -maxdepth 4 | sed 's|/\.git$||'); do
      printf '%-34s lines=%-4s files=%-3s ref=%s\n' "$d" \
        "$(git -C "$d" grep -ci "$TERM" HEAD 2>/dev/null | awk -F: '{s+=$NF} END{print s+0}')" \
        "$(git -C "$d" grep -cil "$TERM" HEAD 2>/dev/null | wc -l | tr -d ' ')" \
        "$(git -C "$d" rev-parse --short HEAD)"
      # 2. name the holes the tools cannot see, then read them by hand — INSIDE the loop, where $d is bound
      git -C "$d" ls-files | git -C "$d" check-ignore --stdin --no-index   # mechanism 1
    done
    ```

    ⚠ **This recipe was wrong in three ways until M257x iter-98, and each way flatters the result.** It
    labelled a **file** count as `hits=` (`grep -c … | wc -l` counts lines of `file:count` output, i.e.
    files); it **dropped `-i`**; and its last line sat **after `done`**, where `$d` is unbound or holds
    whatever the loop left. Run against this rule's own worked example (`TERM=mistral`,
    `stack-demo/app/studio` @ `aeec036`) the printed form returned **2** where the prose two paragraphs
    above publishes **22** — the corrected form returns **22**. A recipe that disagrees with its own
    worked example is the §5-rule-8 failure in executable form.

45. **Name the TREE that settles a claim, not just the trees that exist — an audit briefing that names the
    wrong one manufactures false bookings by construction.** (M257x iter-100.) `rosetta-extensions` has two
    clone roles: the **authoring copy** at `.agentspace/rosetta-extensions` (current, on `main`) and the
    **per-stack consumption copy** pinned at a tag. A corpus claim about what the tooling *does on a stack*
    is settled by the **pinned per-stack clone**, because that is the code the stack runs — the authoring
    copy is where the next tag is written, not what any stack executes.

    The measured cost of leaving this implicit: M257x reading #21 and #22 each booked
    `external_services.md:208-211` as a defect, from **two independent blind seats**, and both bookings were
    rejected — the anchors resolve byte-exact in the pinned clone `ab81527a`. The seats were not careless.
    Their briefing's clone table named `.agentspace/rosetta-extensions` as *"rosetta-extensions (the
    tooling)"* and they graded what they were pointed at. **Two seats making the same error is an instrument
    finding, not seat noise** — and here the instrument told them to make it.

    So a briefing's clone table must carry, per repo, **which tree adjudicates** — not merely which trees are
    checked out and at what sha. Listing both rext clones (as M257x iter-99's ground-truth sheet did) is
    necessary and not sufficient; the sheet said which was which and never said which one *counts*.

46. **A fence's REACH is part of its verdict — read a green as "green over its reach", and make the reach
    gradeable.** (M257x iter-100.) `anchor_construct_guard` reported *"every **resolvable** anchor names a
    construct"* while resolving **360 of 555** citations, and at least seven upheld audit findings were
    citations landing on the wrong construct. The sentence was true. The coverage was 65 %. Nothing in the
    output let a reader turn one into the other, and the load-bearing word was doing its work in silence.

    Three consequences, each of which cost this milestone something:

    - **Set a pre-registration band against the fence's CLAIMED subject, not its measured reach.** The band
      that caught this (*wrong-construct intra-corpus citations ≤ 1*) worked precisely because an upheld
      member would mean a blind spot. It failed by ~7×, and that failure was the finding.
    - **A finding must carry the ref it was graded at.** A run-level *"adjudicated at …"* line names every
      ref the pass touched and cannot attribute one to a finding. `app/main.go:1450` is a closing brace at
      `9d00a313` and a constructor call at `2035f9a`; without per-finding provenance, deciding which one the
      fence meant costs a full re-derivation.
    - **Widening a fence is the easy half — the NARROWING is the deliverable.** The widened resolver went
      360 → 511 anchors and 0 → 23 findings, of which more than half were the guard's own documented
      failure mode returning in a new costume (ports resolving as anchors). Narrow on a **construct the
      corpus demonstrably uses**, never on which findings the narrowing removes — and give every narrowing
      a mutant, because a mutant that is a named kill is what separates *narrowed for a reason* from
      [Trap A](#trap-a--migrations-false-entails-nothing-on-its-own).

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

> **Correction (M257x iter-22), itself superseded (iter-60).** This list originally also named
> `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` @ `2adcf71`. **It was not a miss — it was correct
> at that ref** (`docker-compose.yml:52`, `:258` @ `2adcf71`). **It is no longer:** at `0dab54d` compose
> sets that variable to `http://backend:8083` like the other three. See the rule below; the example is
> kept struck rather than deleted because it is the one that taught the rule — twice, in both
> directions.

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

Items #8/#10 said, at `2adcf71`, that `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401` was stale and
should read `http://backend:8083`. Origin HEAD `2adcf71` said otherwise: **only `SKILLER_RPC_ADDR` was
re-pointed**; at `2adcf71` both `CMS_RPC_ADDR` and `JOBSIMULATION_RPC_ADDR` still addressed the husk
containers — deliberately, per `app/main.go:1196-1202` **@ `app` `5ba17044`** (the same comment stands at
`:1205-1211` @ `b948604` v1.366.0, and is itself now stale — see below), *"additive + DORMANT: external callers (messenger)
keep hitting the standalone cms via `CMS_RPC_ADDR` **until the M809 re-point**."* Applying the correction
**at that ref** would have replaced two true statements with false ones.

> **And then it flipped — twice, which is the second half of the lesson (M257x iter-60, extended at
> iter-87).** At platform `0dab54d` the M809 re-point **had landed**: compose set **four** `*_RPC_ADDR`
> variables, all of them reading `http://backend:8083`, with no cms or jobsimulation container left to
> address. So a passage written to warn *"do not apply this correction"* became a passage that
> **forbade the correction then required** — and it was fortified, which is worse than merely wrong:
> emphatic anti-repair language ("**That address is CURRENT, not stale text**") survives readings
> precisely because it looks adjudicated.
> **Then the correction expired too.** All four variables lived on the `messenger` service block and
> nowhere else; `838d907` (merged `0c91421`, 2026-08-05) deleted that block, so at HEAD compose sets
> **zero** `*_RPC_ADDR` values and exactly one service address —
> `AUTHORIZATION_ADDRESS=http://sentinel:8087` (`docker-compose.yml:48`). Three states in one week, and
> the corrected form of the correction was wrong again within four days.
> **A refutation is a measurement, and it expires exactly like the claim it refuted.**
> Pin it (this passage names its ref in every sentence that depends on it) or it will be read as
> standing. Fenced by `platform_predicate_guard.py` G4.

**Where it came from is the whole lesson.** The refuting citation iter-21 trusted was a corpus line
asserting messenger points *all four* addresses at `backend:8083` — the **`**Connect-RPC**` bullet under
*Interface Discovery*** in [`corpus/services/backend.md`](../services/backend.md). **This passage no longer
carries a line number for it, because the number has now rotted FOUR times** (M257x iter-115). The generations,
each recorded without a live-looking anchor so the count itself cannot rot: the first drifted onto a directory
listing (`bootstrap/ First-run / new-org provisioning`); its replacement drifted onto `askengine/ "Talk to
Data"`; iter-98 and iter-102 re-derived it again; and iter-115's own repair of that very bullet moved it a
fourth time, onto a blank line. **The same anchor rotting four times across four readings** is §5 rule 22's
failure *inside the rule that teaches it*, now with a measured recurrence rate rather than an anecdote — and
the fourth generation was produced by a repair whose whole subject was this class.
Cite the *claim* here, not the line: the sentence to look for is the `**Connect-RPC**` bullet under
*Interface Discovery*, whichever line it currently occupies. At `2adcf71` it pointed **two** (at `0dab54d` it does point all four — the claim was
premature, not permanently wrong, which is its own lesson). **One false corpus line, cited as authority, produced two false corrections in
a hand-off designed to be applied mechanically.** An audit that reads the corpus to correct the corpus is
circular; the citation must terminate in platform source.

> **The failure mode a mechanical hand-off invites is not a moved anchor — it is an inherited falsehood
> wearing a `file:line`.** Anchors are cheap to verify and they were all fine. Verify the *correction* against
> platform source (`docker-compose.yml` / `repos.yml` / the service's Go) before you apply it, every time. And
> when a correction turns out to be wrong, the line that misled you is itself a blocker: hunt it.

Corollary, worth stating because it reads as pedantry until it costs you: **merged-in-production is not
removed-from-compose** — *and the two events are separated by weeks, so the corollary is a phase, not a
permanent state.* At `2adcf71`, `cms` and `jobsimulation` were `service_desired_count = 0` in prod, folded
into `app`, subgraphs gone — and still started containers on every local `make up`, still answered RPC. Two
service docs said *"not in the local compose"*; both were false **at that ref**. The map's word for that
phase is `running_but_unfederated`. **At `0dab54d` the phase is over**: `d11a403` deleted both compose
services and both `repos.yml` entries, so *"not in the local compose"* is now the true statement and the
corollary's own example has expired. Keep the corollary; re-derive which phase you are in before applying it.

**A correction can be INCOMPLETE rather than wrong — and that is harder to catch (M257x iter-23).** iter-22's
hand-off said colony was *"split: `app` + `messenger` @ `v0.35.2`; `sentinel` + `storage` @ `v0.34.3`"*, which
is true of the four services it names. It named four of **six**: at `2adcf71` the `cms` and `jobsimulation`
containers the default selection still started were on a **third** pin, `v0.35.1` (at `0dab54d` neither
container exists, so the set is back to four — the enumeration moved again, exactly as the rule predicts). Applying it verbatim would have replaced one
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

47. **Close a routed item when its DEFECT is repaired — not when its SUBJECT is fully understood.** The
    ledger records defects, not curiosity.

    M257x iter-38 found the corpus asserting an **EU AI Act Limited-Risk** classification resting on a
    technical premise. It measured the premise, found it **false at platform HEAD**, and **removed the
    claim** — correctly refusing to substitute *"therefore High Risk"*, on the ground that replacing one
    unmeasured legal conclusion with another repeats the defect with the sign flipped (it cites the tenancy
    fence, wrong twice in **opposite** directions, as precedent).

    **That repair was complete.** The corpus asserts nothing; the retracted bullets sit inside an explicit
    *"what is STATED, not what this corpus asserts"* fence in
    [`security_compliance.md`](../architecture/security_compliance.md) — **named, not pinned.** ⚠️ This
    carried a line pin into that file (line two-two-seven, under two-three-one to two-five-six —
    **spelled in words on purpose, so this retraction cannot itself be read as a live citation**).
    Re-derived at M257x iter-120 those lines are a `coursebuilder/bedrock.go` citation and the tail of the
    iter-46/48 AI-provider correction; the fence itself was **37 lines away**. A **ninth** member of iter-119's wrong-construct class, and the
    one worth reading the guard note for: `anchor_construct_guard` was **GREEN over it for the whole
    milestone** and only went RED when an unrelated edit above shifted `:227` onto a *blank* line. The
    guard's qualifier is therefore weaker than even its own caveat records — it detects *resolves to
    blank*, not *resolves to the right construct*, so every wrong pin that happens to land on prose is
    invisible to it. Booked as `FIX-M257x-iter120-anchor-guard-detects-blank-not-wrong`.
    **Silence was the correct end state.** What went onto the ledger was not a defect but an **aspiration** —
    *that someone re-derive the true classification* — routed as *"needs an owner outside this milestone."*

    It was then carried as an **open question for 36 iterations**, and re-escalated by three separate
    passes, none of which re-read the item at source. There was never anything to do.

    > **Rule.** An item whose only remaining content is *"someone should determine X"* is **not deferred
    > work** — it is a **finished repair with a wish attached**. Close it at the moment the repair lands.
    > **Never route an item as "needs an owner."** A routing with no owner and no defect is a permanent
    > resident of every future audit.

    The diagnostic is one question: *if nobody ever does this, is any statement in the corpus false?*
    If no, the item is done.

48. **No legal, regulatory, compliance or policy question is escalated during delivery — route it, do not
    ask.** *(Binding user decision, 2026-08-06: "don't bother me with legal stuff during this migration:
    our goal is to close this release, not waste resources on legal matter.")*

    This is rule 47 from the other direction, and the two share a cause. **The corpus documents what the
    code does; it does not issue legal conclusions in either direction.** A compliance question is therefore
    almost never a documentation defect — and when it is, the defect is the *unsupported assertion*, which
    is repaired by **deleting the assertion**, not by sourcing a better one.

    When one surfaces, pick one of three **without asking**:

    | | condition | action |
    |---|---|---|
    | 1 | the false assertion is **already removed** | **close it.** Silence is a valid end state |
    | 2 | a real defect **this repo cannot fix** | **file it** to `platform-defect-register.md` and move on |
    | 3 | it **genuinely blocks delivery** | surface it — and state **exactly what it blocks** |

    Apply it **retroactively** to an existing ledger: anything whose remaining content is a legal or policy
    *determination* rather than a *repair* gets closed or filed, never carried.

49. **A measurement of a CONCURRENTLY-MUTATED surface is timestamped, not standing — and you cannot refute
    another observer's report of one with your own later snapshot.** *(M257x iter-102 / iter-103.)*

    M257x ran three lanes against one checkout. A deferral audit reported **three `iter-101` tags on
    origin** and flagged it urgent. A second observer ran `git ls-remote --tags`, saw **one**, and ruled the
    report a **FALSE POSITIVE** — attributing it to miscounting the peeled `^{}` lines that `ls-remote`
    prints for every annotated tag. The ruling was recorded as a decision and shipped.

    **The report was correct.** Another lane had cut `-101b` and `-101c` and deleted both, from origin and
    locally, in the interval between the two observations. Three tags existed, two of them on one commit,
    exactly as reported.

    The peeled-`^{}` miscount is **a real mechanism and it is not what happened** — which is precisely why
    the ruling was persuasive. It supplied a complete-sounding cause for an observation the ruling had
    already decided was wrong. The refuting decision even recorded that the mechanism *"is not a complete
    explanation on its own"*, and was banked anyway.

    > **Rule.** To refute another observer's report of a surface a concurrent actor can write, you need
    > **their timestamp** or **the surface's history** — never your own later snapshot. Your snapshot answers
    > *"what is true now"*; their report answered *"what was true then"*. Those are different questions, and
    > only one of them was asked.

    **Which surfaces this covers**, and the list is longer than it first looks: git tags and remote refs, the
    clone set, `.agentspace/*` pins, a live stack's containers and ports, `docker system df`, any registry a
    lane writes, and **the reading's own evidence directory**. A surface qualifies whenever an actor other
    than you can change it between two observations — which, under any concurrent-lane protocol, is most of
    them.

    **The generalisable half.** Rule 41a froze the clone refs a *reading* resolves against, because a
    reading's ground truth includes them. This is the same idea one level out: **the evidentiary STATUS of
    any concurrently-writable surface is itself time-scoped.** 41a protects a measurement from a mid-flight
    move; this protects a *report* from being overturned by one.

    Two practical consequences:

    - **Report the instant, not just the value.** *"`ls-remote` at 11:42 returns one tag"* is a fact.
      *"There is one tag"* is a claim about a surface you do not control, and it decays.
    - **A disagreement between two observers of a mutable surface is evidence the surface MOVED**, and that
      is the first hypothesis — not observer error. The move is usually the more interesting finding, and
      here it was: a lane was cutting and deleting tags mid-milestone with nothing recording it.

50. **A guard VERDICT is settled by the tree its CONFIGURATION lives in — print that tree, or you do not
    have a verdict.** *(M257x iter-103 / iter-105.)*

    Rule 45 says *the settling tree follows the claim's SUBJECT*: a claim about what the tooling **does on a
    stack** is settled by that stack's pinned clone. **A guard verdict is not stack behaviour.** It is a
    measurement taken with a fence's *configuration* — its waiver files, its baselines, its assertion set —
    so it is settled by the tree that configuration lives in.

    iter-103 ran the guard family after its edits and got **2 RED**, against its own ground-truth sheet of
    `14 GREEN · 0 RED`. It reproduced at the same corpus commit via a read-only `git archive`, so it was not
    the edit. Two quotable conclusions were drafted — *"the sheet asserted a verdict it did not have"* and
    *"a fence names 8 in-scope sites the double reading missed, so `N ≥ 41`"* — **and both were false.**

    The fence had been run from the **pinned per-stack clone**, not the **authoring copy**. The entire
    difference was one file, `claim_twin_waivers.json` (+40 lines), and the 8 RED sites were **exactly** the
    8 waived sites. Run a fence from last release's pin and you measure **last release's fence**: every
    waiver, baseline row and assertion added since reads as a fresh RED at sites nobody touched.

    > **Rule.** Every verdict carries the fence tree's **path + sha + dirty state**, beside — not instead of
    > — the subject's refs. A verdict whose deciding input is unstated is not re-checkable, and a
    > measurement nobody can re-check is not a measurement.

    **Three details that are the difference between a stamp and a fence:**

    - **DIRTY is a disclosure, not a refusal.** An uncommitted fence tree means the sha names a tree that is
      not the tree that ran — iter-103's failure mode in a subtler form. It is stated and the run proceeds,
      because refusing it would make the family unrunnable during exactly the iters that ship fences.
      *Undeterminable* is different and IS a refusal (`EXIT 2 — UNMEASURED`), with an escape that **records**
      the gap the way `--allow-not-run` does.
    - **Print it FIRST.** A family runner that reports each member's last output line would have every
      guard's own summary silently replaced by a trailing stamp — the iter-87 `headline()` defect exactly.
    - **State it once per run, not once per member**, when the runner executes every member from its own
      directory: the tree is then the same *by construction*, and seventeen repetitions are noise rather
      than evidence.

    **The retroactive half, and it is the expensive one.** Measured at iter-105: **52 recorded family
    verdicts across 26 milestone artifacts, and 0 of them state the fence tree.** None is thereby *wrong* —
    but every one is **provenance-unstated**, which is a weaker thing than a green and should be quoted as
    such until re-run. Treat a guard verdict recorded before its runner printed its own tree the way you
    would treat a search whose stderr nobody read.

51. **A suite that emits nothing is not a suite that is stuck — and until it can tell you which it is, you
    do not have a whole-suite claim.** *(M257x iter-108 → iter-121, thirteen iters.)*

    `stack-core`'s full run was booked as *"blocks indefinitely"* and carried for eleven iters. It does not
    block. It completes. What it does is run a nested suite **8×** at ~16 s each inside one test that prints
    a single `-q` dot for the whole of it — output sat at **522 bytes for 2 m 15 s** and then advanced. That
    interval *is* the reported hang, and an operator watching it kills the run every time.

    Read carefully, the cost was never the duration. **It was the ambiguity**: a running suite and a wedged
    suite emitted the same thing, so no total could be quoted, so every count in eleven iters had to be
    scoped, and the milestone's harden ledger says *"no whole-suite total is quoted anywhere in this entry"*
    three passes in a row. **That is the defect a progress line fixes.**

    > **Rule, two halves and you need both.**
    > **(a) Make the silence readable.** Any harness that blocks on a child process announces the child
    > before it starts and again when it returns, with elapsed time. Write to `/dev/tty` — pytest's default
    > capture is **fd-level**, so `print(file=sys.stderr)` is captured and surfaces only on failure, which is
    > exactly backwards. `rosetta-extensions/stack-core/tests/progress_beacon.py` is the implementation;
    > the wiring is asserted by an **AST walk over the tree**, never a hand-kept list.
    > **(b) State the invocation and the expected wall time with every count**, and treat a run that
    > finishes far off that time as an unexplained measurement rather than a faster one. A count without its
    > invocation is not reproducible, and this suite got **2.5× slower by being FIXED** (431 s → 1090 s: a
    > battery that dies on its baseline never runs its mutants). *A fast suite is not evidence of a healthy
    > one.*

52. **A derivation scoped to ONE field of a multi-field state reports zero and is wrong.** *(M257x iter-121,
    and it landed in the milestone's own close gate.)*

    `deferrals-audit.md` §8 — the section `close-milestone` reads to learn what the user must decide —
    carried a banner reading **"ZERO open user questions remain"** while `state.md`'s own `phase:` field
    said **AWAITING USER SCOPE DECISION**. Nothing was stale and no field was wrong. The iter that opened
    the question graded it **`re-scope: y`** with **`user-blocker: n`**, and the sweep behind §8 read
    `user-blocker`.

    Mechanized (`stack-core/blocking_state_guard.py`) over the three fields that route *out of the iter loop
    to a decision the loop cannot take* — `re-scope`, `user-blocker`, `protocol-stop` — it found **8**
    blocking gradings across 109 graded iters, of which the audit named **3**. Seven of the eight were
    already closed; **that is not the finding.** The finding is that the one genuinely open is the one a
    `user-blocker`-keyed sweep is structurally least able to see, because a *scope* decision is graded
    `re-scope` — and scope decisions reach the user latest and matter most.

    > **Rule.** When a state is graded across N fields, derive over **all N and check the list in both
    > directions** — a field nothing grades must fail loudly, or the list can hold a name that never fires.
    > *"Which field did you read?"* is the same question as *"which file did you read?"* (rule 44), one
    > layer up.

53. **Three consecutive harden passes at the cap without stabilizing is a finding about the METHOD, not a
    request for a fourth pass.** *(M257x passes 22, 25, 26.)*

    Each of the three found real, live defects and closed them; the supply did not thin. All twelve
    green-over-nothing instruments this milestone found share one shape — **the fence's control was aimed
    at something adjacent to its claim**: unreachable controls, the wrong noun, the wrong clause, the wrong
    pin. **Running the suite never surfaces that.** Only mutating the named mechanism does, so a fourth
    pass in the same mode would find a fourth instance and still not measure a plateau.

    **The method finding is load-bearing and outlives the milestone:** three of pass 26's own mutations
    **silently failed to apply** — a `.append` regex that missed multi-line calls, a name set omitting one
    value, a single-quoted pattern against a double-quoted file — and **each read as *"the controls
    survive."***

    > **Rule.** **Every mutation asserts it APPLIED before its result is interpreted** (`assert count == 1,
    > "MUTATION DID NOT APPLY"`). A mutant that did not apply and a mutant that survived are
    > indistinguishable in the output, and only one of them is a finding.
    > **And a mutation is only a control for the clause it can ISOLATE** — a mutant that leaves a second,
    > broader clause satisfying every assertion proves nothing about the clause it targeted.

54. **A correction that reaches ONE cell is not a correction. When a measurement retracts a claim, the
    unit of repair is the PREDICATE, corpus-wide, in the SAME iter that measures it.** *(M257x iter-124,
    against iter-123.)*

    iter-123 cloned `infrastructure` and measured one rule that settled four standing questions at once —
    *a service repo's own `service_desired_count` is not evidence of production state.* It then repaired
    **three** of the four rows sharing that predicate. The fourth (`graphql-wundergraph`) kept the refuted
    reading **in a fenced table**, and **24 further sites across 13 files** kept asserting that the router
    *"survives in production only"*. A second predicate from the same iter — the `db-backup` schedule and
    targets — survived unrepaired at **3** sites, two of them in the security-and-compliance posture.

    **Why no amount of care at the repair site prevents this.** Every surviving instance is *locally
    plausible*: *"the Cosmo Router — prod only"* reads as a careful hedge, not as a falsehood, to anyone
    who does not already hold the measurement. Re-reading the repaired file cannot find them, because they
    are not in it. **Only enumerating the predicate can**, which is [`TOK-02`]'s method and [`TOK-07`]'s
    unit of repair — both vindicated here on a class neither was aimed at.

    > **Rule.** A retraction ships with (a) the **enumeration** of every site publishing the retracted
    > predicate, (b) the repair of all of them, and (c) the **enumeration recorded in the iter** so a later
    > reader can check the reach rather than trust it. **Grade a correction by its reach, never by its
    > correctness at the site that prompted it.**
    >
    > **Corollary — a repair moves line numbers, and line numbers are cited.** iter-124's own repair
    > silently turned four true citations into blank-line anchors; `repair_postcondition` rejected the
    > commit. **Re-run the anchor fences over the WHOLE tree after a multi-file repair**, not over the
    > files you edited.

55. **A finding is placed where the DECISION is made, not where the investigation happened. A census is
    not a disclosure.** *(M257x iter-125, against iter-123.)*

    iter-123 documented the contradicting second corpus in full and correctly — 14 unsourced occurrences
    of a refuted taxonomy figure, load-bearing in four customer-facing tables, in a repo whose plugin
    injects it into every engineer's editor. It landed in a **93-row repo census**. Nothing a reader
    would see changed, because **a reader who needs the taxonomy figures opens the taxonomy section, and
    an engineer about to install the plugin opens the toolchain page.** Neither opens the census.

    **The gap was placement, not writing.** The correct placements were: the section that already holds
    this corpus's own version of the disputed number; the **install line** of the recommendation that
    causes the harm; and the defect register, so the part this milestone cannot fix has an owner.

    > **Rule.** Ask *"where is somebody standing when this finding would change what they do?"* and put
    > it there. An investigation's own write-up is the **record**, not the disclosure — and a finding
    > that only exists in the record is indistinguishable, to everyone downstream, from a finding nobody
    > made.
    >
    > **Sibling of rule 54, on the other axis.** 54: a *correction* must reach every site publishing the
    > retracted predicate. 55: a *disclosure* must reach the site where somebody acts on it.

56. **A guard stuck at "cannot check" is not a careful guard — it is a guard that is not looking. When a
    fence goes could-not-check on a CHANGED SUBJECT, give the new condition its own NAME and its own
    DISCLOSURE.** *(M257x iter-126.)*

    `platform_alignment_guard` went to exit 2 because the corpus began citing `infrastructure` and
    `db-backup` — repos no stack clones. The verdict was **correct on a changed subject**, and the
    subject changed because iter-123 cloned the repo that settled four standing questions. But exit 2 is
    all-or-nothing: **while the guard reported "cannot check", assertion F never ran to completion, and
    it was masking a genuinely mis-bound citation in the same table** (a row citing `:622`/`:664` against
    a 121-line file — the lines belong to a different repo).

    Two conditions had been sharing one bucket:

    | | what it is | whose defect |
    |---|---|---|
    | **unresolvable** | a path the fence cannot find *and cannot account for* | **the citation's** |
    | **unclonable** | a repo the subject document itself names, simply not in the clone set | **the substrate's** |

    > **Rule.** Split them. The substrate-limited class gets its own name, prints on **every** run, rides
    > in `--json`, and **qualifies the verdict sentence itself** — `guard_family.run_one` reports
    > `lines[-1]`, so a qualifier printed anywhere else is invisible where it matters.
    >
    > **And the excuse is GATED, or it is a silencer wearing a disclosure's clothes.** A head enters the
    > excused class only if the subject document *documents that repo*; anything else stays a blind spot.
    > Ship it with a mutant that **withdraws the documentation** and a negative control proving an
    > unqualified green is still reachable — otherwise the qualifier carries no information.
    >
    > **Blindness is not the conservative option.** A fence that refuses to grade anything grades nothing,
    > and the defects inside its reach go unreported for as long as it stays refusing.

57. **A count is only as wide as the SEARCH that produced it — and a repair that inherits its
    predecessor's search inherits its blind spot.** *(M257x iter-128. The sibling of rule 54, on the
    quantifier axis: 54 is about a correction reaching every SITE; 57 is about a correction re-deriving
    from a wide enough SOURCE.)*

    `security_compliance.md` said the REST surface has **"six"** Echo groups, two of which opt into
    `cbGate`. `app` mounts **eleven** non-test groups on the one Echo instance
    (`internal/web/web.go:124-163` @ `ad9f3c498`). Three of the five missing never touch the Clerk
    `authn` middleware.

    > **The rule bit its own worked example, twice, and both corrections are the rule restated.**
    > (a) iter-128 wrote that `/api/invitations` *"has no authentication at all"*, citing the **mount
    > comment** (`web.go:145-146`, *"no auth required"*). Run 82 read the `RegisterRoutes` **call site
    > and the manager** instead and found a required credential: a 256-bit HMAC-derived token checked by
    > lookup before anything is returned (`invite.go:159`, `:194`), with the source's own model
    > *"token possession is the authorization"* (`:154-155`). **A comment is testimony; the call site is
    > evidence** — the same distinction rule 22 draws for commit messages, applied to code comments.
    > *No Clerk middleware* is not *no authentication*, and the corrected sentence is **token-authenticated,
    > deliberately pre-login**.
    > (b) The count itself was widened from one file to the whole service — but **not from *groups* to
    > *routes***. `app` mounts **seven** further routes directly on the root `e`, in no group; a
    > group-level statement reaches none of them. Both are booked in
    > [`security_compliance.md` § Layer 2](../architecture/security_compliance.md#layer-2-authorization).
    > **The widest instrument that can see the population is rarely the one you widened to last time.**

    **The paragraph had already been corrected TWICE.** iter-120 over-stated it (*"every Echo group …
    and nothing else"*); iter-121 corrected the quantifier and cited the exact lines. **Both re-derived
    from `backend.go`** — because that is where the sentence they were repairing pointed. The
    *denominator* was never re-measured, so two careful repairs left a five-group under-count of an
    authentication surface standing.

    > **Rule.** When you repair a count, **re-derive the population from the widest instrument that can
    > see it**, not from the file the wrong sentence names. State the search with the number
    > (`e.Group(` across the whole service, not the one file that happened to declare six of them).
    >
    > **The tell is a repair that only ever quotes the source its predecessor quoted.** If two
    > consecutive corrections cite the same file, the third defect is probably outside that file.
    >
    > **This is not hypothetical humility — it recurred inside the run that wrote it.** The same iter
    > swept the *"Ant Academy is an internal `@anthropos.work`-only portal"* predicate with a regex
    > tuned to the phrasing of the four sites it already knew about; it found **4**. Widening the
    > pattern found **14**, across nine files including `CLAUDE.md`. **The near-miss was the
    > enumeration, not the repair** — and it happened an hour after this rule's own headline was
    > written down.

58. **A conjunction of two whole-document predicates is not an association check.** *(M257x harden pass
    27–28. Found in TWO different fences, one pass apart, and then swept.)*

    `blocking_state_guard` asked whether a blocking grading was *represented* in a 60 KB audit with
    `it in body and field in body`. `unreadable_repo_claim_guard` asked whether a paragraph reported a
    ref-pinned reading with `"infrastructure" in block and _SHA.search(block)`. **Neither conjunct has
    to refer to the same thing as the other**, so both were satisfiable by coincidence:

    * an audit naming `iter-19` for an unrelated reason, with `re-scope` in a paragraph about
      `iter-07` — reported `represented: True`;
    * a paragraph naming `infrastructure` as an **English noun** and carrying the **platform's** sha
      about a different fact — reported as a measurement of a repo in no clone set.

    > **Rule.** When a check means *"X is documented as Y"*, the predicate must bind X and Y in one
    > scope — a block, a sentence, a citation shape. Two independent searches over the same document
    > measure **co-presence**, and co-presence is not association. The larger the document, the more
    > certainly it passes.
    >
    > **Do not reach for a character window.** It is the obvious fix and it is usually *fitted*: here
    > the false green sat ~70 characters apart and the widest TRUE site at 61, so any constant between
    > them is tuned to whichever fixture you happened to write. The durable discriminator was that **a
    > citation names its subject** — as a backticked artifact, or bound to its ref by `@` — and an
    > English noun does not. Both idioms were *measured off the live corpus*, not invented.
    >
    > **A second instance justifies a sweep, not a third fix.** Two occurrences were closed, then every
    > non-test `.py` in `rosetta-extensions` was grepped for the shape (`X in blob and Y in blob`,
    > `.search(…) and .search(…)`, `any(…) and any(…)`). There is no third — **and that is on the record
    > so the next pass does not re-derive it.**

59. **A fence that FALSE-REDs is worse than one that misses, when its only remedy disarms it.** *(M257x
    harden pass 27. The sibling of rule 8, facing the other way: 8 is a green that checked nothing;
    59 is a red that found nothing.)*

    `claim_census_guard`'s per-file ratchet carried an entry for **every** in-scope file, zero included,
    so adding a brand-new corpus file whose every assertion was cited produced
    `[new-file] x.md: 0 unevidenced assertions` and exit 1. The clause says a count **RISES**; 0 has not
    risen.

    > **Rule.** Grade a false positive by **what it teaches the operator to do**. Here the only remedy
    > was `--update-baseline`, which re-seals every *other* file's debt at its current level — so one
    > spurious red would have erased the ratchet's memory across 41 files. **A guard whose false alarm
    > is cured by disarming the guard trains the operator to disarm it**, and the second time it fires
    > nobody reads it.
    >
    > **Its siblings are the same shape.** `gen_override` cleared `postgresql`'s volumes because a
    > predicate meant to catch `$HOME/.aws/credentials` had been widened to *anything under the user's
    > home directory* — which on a normal dev box is the whole workspace — under a comment asserting
    > that `postgresql` "never reaches here". **A fence that harms is not a stricter fence; it is a
    > broken one.**

60. **Nothing runs the whole suite per-iter, so a fence can go RED and stay RED.** *(M257x harden pass
    29. Both of that pass's defects were invisible to every scoped invocation.)*

    A cross-section emitter fence had been failing since iter-129 — it asserted the DEV emitter still
    contained the literal iter-129 had just deleted — and `baseline_mirror_fence` had been failing since
    the same iter, because a provenance block moved away from the heading that named its host and the
    fence's naming lookback resets at a blank line. **Three iters, nobody told.**

    > **Rule.** A scoped green is evidence about its scope and nothing else. iter-121 built
    > `tests/progress_beacon.py` so a 20-minute whole-suite run is *watchable* — but **watchable is not
    > watched**. Either an iter's close runs the whole suite, or the milestone states out loud that it
    > does not and that standing REDs are therefore possible between harden passes.
    >
    > **The corollary for the harden pass itself:** the whole-suite run is not a formality at the end,
    > it is an *instrument*, and on this milestone it is the only one that finds this class.

61. **"Not in the clone set" is a fact about our habits. It never entailed "not measurable" — and the
    corpus inferred one from the other for four iterations, at 15 sites.** *(M257x iter-132.)*

    Six passages said cms's production disposition was UNMEASURABLE, *because* the deciding declaration
    lands in `infrastructure` and `infrastructure` **"has never been in any clone set."** Both conjuncts
    were true. The inference was not: iter-123 had already cloned the repo and settled the question, and
    the corpus cited that read 28 times **while eleven other sites kept publishing the limit it
    retired.** iter-132 spent one `--depth 1 git clone` — well under a minute — and settled a *second*
    standing hedge from the same tree (the production RPC address, hedged in seven more places).

    > **Rule.** When a claim turns on a repo that is merely **not habitually cloned**, the remedy is a
    > `git clone`, not a hedge. Write *"not in the standing clone set, therefore not re-derivable in
    > place"* — which is about **re-checkability** — and never *"not measurable"*, which is about the
    > world. A hedge that outlives its premise is worse than a wrong claim, because it **advertises that
    > somebody checked.**
    >
    > **Two riders, both earned by this iter's own prose.**
    >
    > *(a) The fence catches you at exactly this.* `unreadable_repo_claim_guard` flagged **three
    > paragraphs of the repair itself** — assertions about `module.messenger_euwest1` (deleted at
    > `infrastructure` `13c248e6`, `terraform/production/services.tf:618-621`) with no ref in the
    > paragraph. **The third was this very rule**, which fired on its own worked example. A sweep that
    > retires a hedge must carry the ref *into every paragraph it touches*, not only into the one where
    > the reading is narrated.
    >
    > *(b) A substring fence cannot tell a RETRACTION from a live hedge, and will report your repair as
    > the disease.* The NOTE read **"11 sites still hedge"** — of which **8 also carried a ref-pinned
    > reading**, i.e. they quote the retired wording precisely in order to retract it. The fix is **not**
    > to launder the prose until the substring disappears; that is tuning the corpus to the instrument.
    > The fix is a **third bucket, disclosed** (`hedged` / `mixed` / `measured`), so the NOTE counts live
    > hedges and the mixed count is published as a thing the guard **cannot** decide — `D-M257x-121-4`'s
    > precedent, applied a second time. **Do not read a marker count as a defect count.**

62. **The search that PLANS a repair and the search that VERIFIES it must not be the same search — and
    a subject wrong in two OPPOSITE directions has no owner.** *(M257x iter-137. The procedural corollary
    of rule 57, plus its two-directional case.)*

    Rule 57 says a count is only as wide as its search. iter-137 obeyed it — four independent searches,
    two per conjunct, before any edit — repaired **26** sites, then re-ran the same searches **with the
    newly-written wording `grep -v`'d out** and found **three live survivors** no planning search could
    reach: one inside a paragraph about gRPC hops, one in an index row, and one whose falsity is carried
    by a **section heading**.

    > **Rule (a) — verify with the corrected vocabulary.** After a sweep, re-run the search **excluding
    > your own new phrasing**. What survives is exactly what the planning search was blind to. A repair
    > that verifies itself with its planning query is measuring its own vocabulary.
    >
    > **Rule (a′) — then read the HEADINGS.** The third survivor was a bullet under
    > ***"Domains inside Backend/App, not services"*** — the false predicate asserted **one level up** and
    > inherited by every bullet beneath it. **Grep, the anchor fences and the claim census all read
    > LINES**, so this class is invisible to all of them. A heading is exactly where a corpus states its
    > general claims, which makes it exactly where a general claim goes stale unseen.

    The same iter found the subject wrong **in two directions at once**: `roadrunner` was asserted to be
    **running in production** (*"prod terraform still reads `service_desired_count = 1`"*) **and** to be
    **one of eight domains folded into `app`**. Both false; both live; mutually exclusive. A reader could
    open two files of the same corpus and come away with incompatible pictures, each stated confidently.

    > **Rule (b) — two-directional error means no owner.** When a subject carries contradictory false
    > claims, neither is a typo. Repairing one conjunct makes the corpus *look* more consistent and no
    > more true. **Sweep the subject, not the sentence** — and check whether a rule you already wrote was
    > applied to **every row of the table it was written for**. § 3's *"a service repo's own
    > `service_desired_count` is not evidence of production state"* reached `cms`, `messenger` and
    > `graphql-wundergraph` at iter-123 and skipped `roadrunner` — **one row away, in the same table.**

    > **Rule (c) — never quote a retracted line-pin; describe the artifact.** `roadrunner.md` carried a
    > bare `:NN` pin **as its own worked example of a bad pin**. iter-137's repair shifted the file, the
    > quoted pin landed on a blank line, and **two fences went RED on a citation whose purpose was warning
    > against citations like it.** A form-matching fence cannot distinguish *asserting* a pin from
    > *quoting a retracted* one — the anchor-axis sibling of (b) above and of iter-134's 1-of-4
    > measurement. **Delete the pin; name the construct.** Re-pinning restarts the clock.

63. **An exclusion is only as narrow as the predicate that justified it — and ANCHOR ROT is mechanically
    decidable even where the anchor's CLAIM is not.** *(M257x iter-138.)*

    `corpus_citation_guard.py` excludes **bare `:NN` pins** *"outright"* as *"not mechanically
    decidable."* That is **true of the claim** — *does this line say what the sentence says it says?*
    requires reading the sentence. **It is false of rot:**

    > If the citing line was authored at commit `C`, and the text that stood at the cited line at `C` now
    > stands at a **different** line, the pin **rotted** — and the line it now stands at is the repair.
    > No sentence is interpreted. Git answers it alone.

    ⚠️ **iter-138 measured this corpus-wide and reported 127 rotted of 222 decidable (57.2 %).
    iter-139 audited that number — a stratified 12-case sample, strata sealed in advance — and found
    12 FALSE POSITIVES out of 12 (precision 0.0 %, Wilson95 [0.0, 24.3]). The figure is RETRACTED.**
    In this corpus a bare `` `:NN` `` is **overwhelmingly a cross-file continuation pin**
    (`` `app/main.go:15`, `:62`, `:63` ``) or a quoted/historical/negated pin — not a same-file
    self-citation. **Resolving the HEAD is the hard part, and it blocks the rot predicate exactly as it
    blocks the content predicate**, so the fence's exclusion is better founded than iter-138 credited.
    `adj-E`'s five genuine same-file rotted anchors were found by a **human reading five sentences**;
    they are a rare form, and a machine must solve head resolution before it can enumerate them.

    > **Rule (a).** When you exclude a class from a fence, **name the predicate you actually tested.**
    > The next predicate over *may* be free — **but check whether both predicates are blocked by the same
    > thing.** Here they were: head resolution. iter-138 assumed rot was free of it and was 0-for-12.
    >
    > **Rule (b) — a DISCLOSED floor is quarantined only if you show the boundary holds.** iter-138
    > disclosed `out-of-range-then` (241) and named its cause precisely — *cross-file continuation pins* —
    > then reported a number over the **222 it had not excluded them from**. A continuation pin only lands
    > in that bucket when its line number exceeds the **citing** file's length; in a 3,100-line doc almost
    > none does, so the same failure mode dominated the "decidable" set undisclosed. **Naming a floor is
    > not bounding it, and the disclosure made the number more persuasive, not less.** Sample the clean
    > bucket for the disease you just disclosed.
    >
    > **Rule (c) — repair by NAMING the construct, never by re-pinning.** A repair that restores the
    > failing *form* fixes an instance and preserves the class. `graphql-wundergraph.md`'s `5050` pointer
    > rotted **twice** (`:174-176` → iter-98 → `:193` → iter-138) and one of its paragraphs held **three**
    > rotted pins. The durable citation is the construct name plus a substring `grep` returns uniquely.
    >
    > **Rule (c′) — the corpus's own RETRACTION IDIOM is a rot generator, measured.** The house style
    > *"it was `:274` at `<sha>`"* / *"this cited `:116-117` until iter-NN"* keeps the retracted number
    > **live in the text**, where the next insertion above it moves its target. **In one session (M257x
    > iters 137–141) this turned fences RED three times, in three different files, always on a pin whose
    > own sentence existed to retract it** — `roadrunner.md`'s async-tasks paragraph,
    > `graphql-wundergraph.md`'s `5050` pointer (rotted **twice** on its own), and
    > `ai-readiness.md`'s `:326`/`:274` note. **Retract by describing the artifact, not by reproducing
    > it**: *"this doc carried two different line numbers for it in successive iters"* says everything the
    > quoted number said and cannot rot. A fence matching on **form** cannot tell the quotation from the
    > assertion, and it is right not to.
    >
    > **Rule (c″) — a cross-reference that names its target by a RETRACTED TITLE is invisible to every
    > anchor fence**, because the pointer still resolves. `backend.md:13` sent readers to *"the **M810
    > prod teardown is UNEVEN** bullet below"* — a bullet retitled *"…has now LANDED for both"* at
    > iter-127, whose body retracts *"UNEVEN"* explicitly. The reader arrives at a paragraph that opens by
    > contradicting the sentence that sent them. **Name your target by what it says now.**
    >
    > **Rule (c‴) — the class was CENSUSED, and the census is now a fence.** *(M257x iter-142.)*
    > `TOK-08` says a reading SAMPLES and a fence CENSUSES; (c′) named the class from three incidents,
    > and iter-142 enumerated it. `retracted_pin_guard.py` scans the published corpus for a **backticked
    > line pin inside a retraction clause** — a clause carrying both a reporting verb (`was` / `were` /
    > `cited` / `said` / `read` / `pinned` / `carried` / `listed`) and a supersession marker. Over
    > **2,185 line pins in 94 documents** it found **50 live instances across 20 files** — 44 bare,
    > 6 path-qualified — every one hand-read before a line of prose was repaired. All 50 are now
    > descriptions.
    >
    > ⚠️ **This said *"the class stands at 0 and the fence holds it there"*, and that 0 was over a
    > population that excluded a large share of the class.** *(Corrected at M257x harden pass 30.)*
    > The census was **LINE-scoped** — every clause window lived inside one line — and this corpus
    > **hard-wraps at ~100 columns**, so a retraction routinely straddles a soft break: the reporting
    > verb ends one line and the pin opens the next. Joining each line to the one above surfaces **10
    > more**, hand-read **8 true / 2 false**, all of them still live *after* the repair that reported
    > zero. The gating arm genuinely is at 0; **the fence now says which population that 0 is over**,
    > and carries the wrapped arm as a disclosed non-gating SURVEY (80 % precision reports, it does
    > not gate). The eight are routed as `FIX-M257x-h30-crossline-repair`.
    >
    > **The sharp part is that the family already knew.** `platform_predicate_guard`'s `_pin_window`
    > has joined `line[i-1] + line[i]` since **iter-63**, `_NEGATED` needed the same widening at
    > **iter-68**, and a third site records *"line-scoped it reached only 2 of the 4 — two of the live
    > sites wrap."* Three prior encounters, and the newest fence in the family still shipped
    > line-scoped — because the remedy lived in one guard's source comments and nowhere a new fence's
    > author would read. That is what rule 64 below is for.
    >
    > **The path-qualified half was found by the GUARD FAMILY, not by this fence, and that is the most
    > useful thing the iter produced.** The first draft matched only the bare form, reported the class
    > clean, and `repair_leak_guard` went RED **on the commit that repaired it** —
    > `shared_libraries.md`'s **analytics-go** row still published a twin of a `CLAUDE.md` site the fence *had* flagged,
    > missed for one reason only: its pin carried a path. **A fence whose regex defines its own
    > denominator will report a clean census over the population it happens to match.** The
    > commit-scoped guards are the only thing standing between that and a published "0".
    >
    > Four things the census established that the rule as written did not:
    >
    > 1. **Precision lives in the SECOND half of the predicate, not the first.** `was` / `read` /
    >    `said` alone match ordinary description — *"the score column, read at `:1820`"* is a live pin,
    >    not a retraction. The first draft put every supersession marker in one anywhere-in-the-clause
    >    set and produced **2 false REDs of 46** on its first run, both from markers that attach to
    >    something OTHER than the pin: a sha anchoring a *deletion*, and a *"no longer"* about a
    >    deleted FILE. Demoting those two to a tier that must sit **within 30 characters after the
    >    pin** — exactly where the genuine idiom puts them — took it to 44 of 44. Where iter-138 was
    >    0-for-12, this is 44-for-44, and **the difference is that the audit happened before the
    >    repair, not after the publication.**
    > 2. **A ref-qualified historical pin is IN the class, and that is not a technicality.** *"it was
    >    `:100` at `0dab54d`"* is a TRUE statement about an immutable ref — and it still turns a fence
    >    RED, **because the fence does not read the qualification.** The hazard is the token, not the
    >    truth of the claim wrapped around it. That is (c′)'s *"a fence matching on form cannot tell
    >    the quotation from the assertion"* stated from the other side.
    > 3. **Fence the TOKEN, not the digit — so the repair never has to destroy evidence.** *"rotted
    >    +8"*, *"iter-102 added +23 and +16 instead of re-measuring"*, *"it stood ten lines earlier"*
    >    keep everything the pin said and are invisible to every resolver. All 20 repaired files came
    >    out **line-count FLAT** (rewritten in place, 0 net shift per file, `git diff --numstat`), so
    >    the sweep that removes anchor rot could not itself induce any — the failure mode iter-141
    >    caused with its own repair.
    > 4. **The two arms need DIFFERENT windows, and the asymmetry is measured, not designed.** Requiring
    >    the supersession marker anywhere in the clause scores **44 / 44** on bare pins and **1 of 5** on
    >    path-qualified ones, because a path pin lives in long evidentiary sentences where `until` is
    >    doing other work — *"declared in prod **until** it was destroyed, `services.tf:509-517`"* is a
    >    **live** citation, and *"said `= 1` **until** iter-137 — `main.tf:19` is an input to a module
    >    never instantiated"* retracts a CLAIM while its pin is the evidence refuting it. **Retracting a
    >    claim does not retract its evidence.** Requiring the marker *after* the pin separates that
    >    population — 6 of 6 true, 4 of 4 false dropped. ⚠️ **That window is a TUNED CONSTANT on a
    >    five-site denominator, and saying so is the only thing that makes it honest**;
    >    `FIX-M257x-iter142-path-arm-window` carries the derivation forward.
    >
    > **Rule (d) — choose the test suites by what you CHANGED, not by what you were writing ABOUT.**
    > iter-137 rewrote **29 anchors** and picked its scoped suites by topic (`platform_alignment`,
    > `claim_twin`, …). `anchor_offset_guard` was **NOT-RUN** in the family (commit-scoped, no `--range`)
    > **and** absent from the scoped set — two mechanisms blind in the same direction on the same commit,
    > and an ambiguous `README.md` head shipped (six files share that basename). Caught one iter later by
    > the anti-vacuity control built for it. **A disclosed not-run bucket is not coverage.**

64. **A fence over WRAPPED PROSE must state its line reach, because "one line" is not a unit of meaning
    here — and this corpus has now learned it FOUR times.** *(M257x harden pass 30.)* Every guard in
    this family scans `splitlines()` and matches a window around a hit. That window silently defines
    the denominator, and corpus prose hard-wraps at ~100–110 columns, so an association whose two
    halves land either side of a soft break is **not merely missed — it is never enumerated**, and the
    guard reports a confident zero over the population it happened to see. The prior three:
    `_pin_window` joined `line[i-1] + line[i]` at **iter-63**; `_NEGATED` needed the same widening at
    **iter-68**; a third `platform_predicate_guard` predicate records *"line-scoped it reached only 2
    of the 4 — two of the live sites wrap."* **And §7 rule 4 of this very document already says it in
    the general form — *"the paragraph is the unit of publication"*.** `retracted_pin_guard` was
    written after all four and shipped line-scoped anyway, hiding 10 live members of the class it was
    built to census.

    The lesson is **not** "always join lines" — joining has its own hazards, and the same iter-63 note
    is where the exclusions come from: a blank line, a heading, a fence delimiter and a **table row**
    each end a paragraph, and adjacent `|`-rows are separate records that merely sit next to each
    other. The lesson is that **the reach is a design parameter that must be chosen and STATED**,
    exactly like `AssertPublicOnly`'s scope or a green's reach under rule 46. A fence that does not
    say whether its unit is the line, the paragraph or the document has not said what its number
    counts.

    **Why it kept recurring is the transferable part: the remedy existed only as source comments
    inside one guard.** Nothing a new fence's author reads — not §5, not the fence template, not a
    test — carried it, so each encounter paid full price and left the fix where the next author would
    not look. **A defect class that has been solved in code but not in the rulebook is unsolved.**
    When a fix generalises past the file it lands in, promoting it here is part of shipping it.

    **The remedy now has a home:** `rosetta-extensions/stack-core/prose_reach.py` —
    `continues_paragraph()` / `join_prev()`, one definition with both consumers asserted to hold the
    same object, and every exclusion iter-63 paid for carried as a named test. Deliberately not
    called `*_guard.py`, so the family's derived registry excludes it by construction.

    **The sweep that followed is the rule's own denominator, and it found a second instance.** All 22
    guards were classified by scanning unit: `unreadable_repo_claim_guard` already works in
    blank-line-delimited paragraphs and is immune; `platform_predicate_guard` has joined since
    iter-63; the rest do not associate two things across prose. **`clone_drift_guard` did.** Its D2
    rule wanted the `go.mod` citation and the module pin on one line, and the `staging-sync.md`
    colony-requires sentence splits between them — so the site fell through a `continue`
    **silently**, three lines below a comment reading *"Named, not silently dropped."* Joining moves
    that guard from **3 graded sites to 4**, and changes nothing else.

    ⚠️ **A WIDENED REACH MUST NOT INVENT SUBJECTS, and this pass proved it the expensive way — on
    itself.** The first cut of that join reported four further sites and a relabelled fifth. They
    were **phantoms**: the join fired on lines carrying no module pin at all, attributing a citation
    to the *blank line beneath it*, and the harden pass wrote all five into a commit message and into
    this rule before noticing. What noticed was **`anchor_construct_guard`**, which resolved two of
    the published coordinates and found blank lines — a fence on a different axis catching the
    fence-widening, exactly as the guard family caught the path-arm gap at iter-142. Two rules fall
    out, and the second is the one worth carrying:

      * `continues_paragraph` answers whether the PREVIOUS line continues; **only the caller knows
        whether the current line has anything to continue INTO.** The precondition belongs at the
        call site and is now a named regression test.
      * **A measurement taken with a just-changed instrument is a claim about the instrument until
        something independent confirms it.** The five phantoms read exactly like findings — they had
        file names, line numbers and a plausible mechanism. Nothing inside the changed guard could
        have separated them from real ones. The corrected delta is **one** site.

    Where the widened evidence genuinely matters is a case the corpus does not currently contain but
    will: when a ref pin sits on the CITATION's line rather than the pin's, a join that imported only
    the citation would grade a ref-scoped claim against HEAD. That ordering ships, guarded
    **prospectively** by a synthetic test — and it is labelled as such, rather than as something that
    was measured.

65. **An AUDIT is a predicate too, and it needs a control that is not another reading.** *(M257x
    iter-143)* Rule 63 and `D-M257x-142-1` say *audit the predicate before the repair, not after the
    publication* — iter-138 published a mechanical number and was audited to 0-for-12; iter-142
    audited first and repaired 44 of 44. iter-143 followed that instruction exactly, and **the audit
    was still wrong.**

    It hand-read **all 92** anchors a candidate head-inference admits, graded each against its full
    line context, and published nothing. The reader's verdict: **62 true / 30 false**, and the best
    structural predicate scored **90.2 %**. Then the 62 "true" sites were pushed through the guard's
    own `classify()` — and **nine came back `anchor-out-of-range` against files whose line count made
    the citation impossible.** All nine were one mechanism the reader had no way to see: a bare file
    **mention** sitting nearer to the anchor than the **qualified citation** actually governing it
    (*"`studio/gen.py` at `studioManager.go:119` and `studio/postgen.py` at `:1045`"* — `:1045` is a
    `studioManager.go` line). Corrected, the population is **53 true / 39 false** and the same
    predicate scores **74.5 %**. The audit had inflated it by **15.7 points**.

    > **A reading is evidence about a corpus. It is not evidence about itself.** When the deliverable
    > is a predicate's precision, the audit that establishes it needs an independent check — and the
    > cheapest one is usually already in the tree: run the audit's own TRUE set through the machine
    > and see whether the machine agrees. It costs one script and it caught nine.

    Two corollaries, both measured here:

      * **State the reader's error rate alongside the predicate's.** *"Hand-read, 92 of 92"* sounds
        like ground truth and was 90 % accurate. A precision figure resting on an unaudited audit
        inherits its error silently, and in the same direction every time — **towards shipping**,
        because the sites a reader mis-grades are the ones that look right.
      * **The two failure modes are not equally visible, so count them separately.** Of the 39 false
        admits: **21 ports** (loud — they resolve out-of-range and show up as a RED) and **16
        WRONG-HEAD** (silent — a real line anchor booked against a file the sentence never named,
        which can land on a real construct and PASS). A guard's own source comment had named only the
        loud half for five iters. **When you decline a widening, decline it for the hazard that
        cannot be seen when it is wrong.**

    The iter shipped the census and the two no-inference reach gains, and **did not ship the
    inference** — 77.3 % precision at 32.1 % recall is not fence quality, and this milestone has
    already retracted one mechanical publication over this exact population.

66. **A change-derived test scope is only as good as its DERIVATION — and "what imports this" is not
    the derivation, "what call sites does this break" is.** *(M257x iter-143)* Rule 63(d) says pick
    the suites by what the iter CHANGED. iter-143 did, and named three modules on the reasoning
    *"these consume the changed return value."* They passed, **106 of 106** — green, real, and
    irrelevant. The whole suite then returned **31 failed**, thirty of them the same iter's own
    change: a census had been added as an **eighth member of a function's return tuple**, and one
    module unpacks that tuple **positionally at six call sites**.

    The scope was chosen from a recollection of imports. The breaking consumption pattern —
    *positional unpacking* — is not visible in an import list and would have taken one `grep` for the
    call sites to find.

    > **Arity is a published interface.** Adding a member to a returned tuple is a **breaking change**
    > wearing the costume of an addition. So is renaming a dict key, reordering positional args, or
    > widening a return type a caller pattern-matches on. None of them look like removals, and all of
    > them are.

    Two rules fall out, and both are cheap:

      * **Derive the scope by searching for CALL SITES of the thing you changed**, not by listing the
        modules that import it. A `grep` for the symbol is the derivation; memory is not.
      * **When an idiom for this already exists in the file you are editing, use it.** The census
        belonged on a module-level accumulator cleared at entry — and *two* such accumulators sat ten
        lines above the edit, doing exactly that job. The tuple was the lazier reach past a pattern
        the module had already established, and it is the only reason there was a breaking change to
        make at all.

    **And the third instance of one pattern, which is why this sits next to rule 65.** iter-142's real
    miss was caught by a guard on a **different axis** run over the commit; iter-143's audit error was
    caught by **`classify()`**, not by a second reading; iter-143's arity break was caught by the
    **whole suite**, not by the scoped set. In all three, everything *inside* the thing being built was
    green. **The check that catches you is never the one you designed while making the change** — which
    is the standing argument for rule 60's whole-suite debt being a debt, and not a formality.

67. **A retraction clause contains the CORRECTED pin as often as the retracted one — and a fence
    matching the clause cannot tell which it is holding.** *(M257x iter-144)* Rule 63(c′) says the
    corpus's correction idiom keeps a retracted number live in the text. True, and incomplete: a
    correction has **two** halves, and the second half is a **live** pin that the corpus is right to
    publish.

    Measured on `retracted_pin_guard`'s wrapped arm, over its whole population of 10: **7 true, 3
    false — 70 % precision.** All three falses are one shape, and it is the shape of a *correction*
    rather than a *retraction*:

      * `ai_architecture.md:303` — *"Both anchors named the file nowhere in this bullet until iter-115
        — they read as bare `:98-99` / `:110-111`"*. The clause retracts the **absence of a filename**,
        not the numbers. Resolved at `app` `ad9f3c49`, `:98-99` **is**
        `default: aiModel = anthropic.Anthropic37SonnetAWS20250219` — the pins are live and right.
      * `hiring.md:304` — *"It cited an earlier range — the `job_position` bullet, now `:176-187`"*.
        The retracted value is *"an earlier range"* and is **not named**; `:176-187` is the
        **correction**. Read at HEAD, that range **is** the `job_position` bullet.

    > **The token a retraction must not reproduce is the OLD value. The token a correction must
    > publish is the NEW one.** They are the same shape, in the same sentence, inside the same
    > markers — so a form-matching fence sees one class where there are two.

    Consequences, and the first is the one that generalises past this fence:

      * **Grade a survey arm's findings before treating its count as a backlog.** Harden pass 30
        routed *"the 8 true sites across 6 files"* forward without grading them; measured, the arm
        held **10 findings, 7 of them true**. A routed count is an *estimate of work*, and quoting it
        as a *measurement of defects* is how a survey number becomes a fact.
      * **A survey arm is the right home for a 70 %-precision predicate**, and pass 30 was right to
        keep it non-gating. The lesson is not to tighten it — a fence cannot read which half of a
        correction it is looking at — but to **grade before repairing**, every time.
      * The three survivors stay, and the guard reporting them is **correct behaviour, not debt**.

68. **A suite you never run is not GREEN, it is UNMEASURED — and "not ours" derived from a diff SCOPE
    is a statement about the window, not about the code.** *(M257x iter-145)* The same
    UNMEASURED-is-not-UNMOVED discipline `§9` guard-rail 1 enforces on the primary metric, applied to
    two places it had never been applied: the test suite's **denominator**, and a failure's
    **attribution**.

    **The measurement.** M257x ran a "whole suite" at every iter close and 32 harden passes. *"The
    whole suite"* meant one of five `rosetta-extensions` sections — **1,281 of 3,062 tests, 42 %**,
    where *tests* means **executed = passed + failed + skipped**.

    > ⚠️ **This figure read `1,280 of ~3,040` until M257x iter-173, and the way it was wrong is the
    > rule it is evidence for.** The denominator was assembled as `2,978 passed + 11 skipped = 2,989`
    > from the harden ledger's own five-section table — **dropping that table's 22 failures** — and the
    > next entry carried the hole forward as `2,989 + 51 = 3,040`. So the rule that *"the whole suite"
    > must name its denominator* was published with a denominator that had silently changed its unit
    > from *executed* to *passed-and-skipped*: iter-172's *name the unit* defect, one level up, inside
    > the rule about denominators. Re-derived: `2,978 + 22 + 11 = 3,011` executed at that table, of
    > which `stack-core` was `1,229 + 1 = 1,230`; after the `+51` the section is **1,281** and the
    > suite **3,062**. The four other sections come to **1,781** in the ledger's table *and*
    > independently in iter-145's own re-run of them — two runs, same executed total. **The 42 % holds
    > either way**, which is exactly why nobody looked: a percentage can survive an error its operands
    > do not. Fenced going forward by `stack-core/derived_count_guard.py`, which proves the table this
    > figure is derived from — it cannot reach the `N of M` prose shape itself, and says so.

    Run once in full, the four never-run sections held **21 failures**, and they had been routed
    forward three times with one characterisation: *"provably not ours … pre-existing,
    environment-coupled"*, evidenced by `git diff --name-only <5-iters-ago>..HEAD` returning only
    files in the section that was being run.

    Graded individually, the 21 is **three populations**:

    | cause | n | decided by |
    |---|---|---|
    | a **real defect, and the milestone's own** | **12** | one root cause, bisected to a named commit |
    | whole-file **sha baseline drift** vs an advanced but **clean** clone | 6 | `git status` empty + the sibling **anchor** assertions GREEN |
    | **host environment** (no live postgres socket) | 3 | `pg_isready … no response` |

    So the routed characterisation was **false for 57 % of the population it described**, and six of
    the twelve are pure table arithmetic touching no clone and no container.

    **The defect is this document's own thesis, reproduced by the iter written to end it.** Platform
    `2adcf71` deleted the GraphQL router; **iter-13** dropped its row from the tooling's service table
    and left the *test side's* copy — a registry map plus **six independent count literals**. Twelve
    tests went RED that day and stayed RED for **132 iters**. iter-13's own commit message reads
    *"six copies of a platform fact is the hand-maintained-tuple defect M257x exists to end"* — and it
    left a seventh copy, in the one place nothing watched.

    Three sub-rules, each of which had to be true for it to survive four months:

    * **(a) A window that opens after the breaking change can never see it.** `git diff A..HEAD` asks
      *"did these iters touch it?"*, not *"who broke it?"*. The answer `no` is correct and says
      nothing about authorship. **Bisect the failure, or write "not attributable within this window"**
      — which is a claim about the measurement, and honest.
    * **(b) Keep the MEMBERSHIP literal, derive the COUNT, and fence the two against each other.** A
      hand-written set is often the *anti-vacuity control* (`§8`) and must stay independent of the
      subject; a **count** is never a control, only a restatement, and restatements rot. Six literals
      restating one table's size are six things to forget. One literal set + derived counts + a
      membership fence is one thing to update and one assertion that says so.
    * **(c) A failure message that names nothing gets ignored, and being ignored is how it survives.**
      `12 != 13`, `9 != 10`, `'all 14 …' not found in '… all 13 …'` were on every run for four months
      and named neither the row, the file, nor the service. The fence that replaced them fails once,
      naming the drifted rows in both directions — **both detect the drift; only one says what
      drifted**, and that difference is the whole value.

    **The cost argument for the narrow scope does not survive measurement:** the four excluded
    sections cost ~11½ min together, **less than half** what the one included section costs. The
    excluded 58 % was the cheap half.

    * **(d) A repair's completeness is a function of what gets EXERCISED, not of the author's care —
      so census the un-exercised paths FIRST.** *(M257x iter-146)* iter-145 established that iter-13's
      re-point had a hole; iter-146 censused the rest of it. Over the tooling repo: **84 references to
      the deleted router's port across 31 files**, and after classification —

      | class | n |
      |---|---|
      | correct re-point · fence asserting its absence · guard prose · test fixture | **82** |
      | **latent** — a build-arg DEFAULT baking a dead endpoint, never reached because the caller always passes `--build-arg` | 1 |
      | **LIVE** — an operator-facing URL for a port with no listener | 1 |

      **97.6 % complete, and the misses were not random.** Both landed where nothing executes: a
      never-run test section (iter-145) and a `--public-host` branch that needs tailscale plus a
      public host to reach. Everywhere the code actually runs, iter-13's re-point held.

      The live one is the sharpest thing in the census. `gen_tailscale_serve.py` deleted the router's
      `tailscale serve` row and explains why in its own words — *"fronting a port with no listener
      produced a trusted-cert HTTPS endpoint that always refused, **which is worse than no entry at
      all (it looks configured)**"* — and one file over, `dev-stack` still **printed that URL to the
      operator** as the first line after a successful bring-up. **The repair removed the mechanism and
      left the announcement.** Grep for the *emitters* of a retired fact, not only its consumers.

      And the fence over this class has to carve out **comments**, for rule 67's reason: the four
      files that document the deletion best are the four that name the dead port most. A predicate
      matching the bare token would go RED on exactly the right code. *Executable content only — a
      comment may name a dead endpoint; a command may not.*

69. **A token census finds a value that is WRONG; it can never find one that is ABSENT — and when a
    repair's own write-up names a sibling that lagged, the repair is not finished until a fence spans
    BOTH siblings.** *(M257x iter-147)* Two halves of one iter, and they compound.

    **The reach limit.** Rule 68(d) censuses a retired fact by grepping its token. That is exhaustive
    over *wrong* values and structurally blind to *missing* ones. iter-147's defect was an **empty**
    compose profile — `docker compose up -d` with no `--profile` selects only the services declaring
    no `profiles:` key (here: postgresql, redis, sentinel), **exits 0**, and the stack looks alive
    with the application absent. There is no string to search for; the defect presents as
    `--profiles ` followed immediately by the next flag. **So invert the search: enumerate what the
    tooling ANNOUNCES or CHOOSES and grade each against the platform**, rather than hunting a
    known-dead token across all files. Denominator, stated: **7** profile-selecting compose sites,
    **5** already deriving, **2** defective — and the 2 were the two nothing exercises, because every
    documented invocation passes the flag explicitly (`D-M257x-146-2` again).

    **The twin-lag half, which is the expensive one.** This exact defect was repaired three times
    before iter-147: iter-55 (demo teardown + injected-gen), iter-85 (both `dev-stack` verbs). And
    **iter-85's own comment names the lag it was closing** — *"gen_injected_override.py derived this
    at M257x iter-55 for the demo path; the dev path kept the literal for four more releases"* —
    while leaving the demo path's other two verbs untouched for another **62 iters**. An observation
    about a twin is not a fence over it. When a repair's rationale says *"the sibling has had this
    since X"*, the sibling set is already enumerated: **write the fence over the set, in that
    commit.** iter-147's fence asserts BOTH entry points derive on BOTH their verbs, and names which
    side drifted — rule 68's lesson, applied at the moment the twin is first noticed rather than at
    the moment it next breaks.

    **Corollary for a defaults contract.** `demo-up-defaults.md` promises *every knob with its real
    default*. The `--profile` row had **no default column entry at all**, and the omission read as
    "there is nothing to say" rather than "the default is none" — which is what hid this for four
    releases inside the one document written to prevent exactly that. **An omitted default is a
    claim, not a gap.**

70. **Derive from the artifact that DECIDES the fact, not from one that merely constrains it — and
    remember that an over-broad scope is loud while an under-broad one is silent.** iter-148 fixed
    `/test-platform`'s probe scope by deriving it from `$STACK_ROOT/platform/docker-compose.yml`. That
    is the right *shape* (derive, never hand-write) against the wrong *artifact*: the platform's
    unmodified compose constrains what a stack **can** run; the stack's own **generated override** is
    what decides what it **does** run. Measured at platform `0c91421` — platform set **5**, the stack's
    own override **11**, so `/test-platform` probed five of eleven and printed `✓ pass`.

    **The asymmetry is the lesson.** iter-148's over-broad scope printed four false `down`s and was
    repaired inside one iter, because a false failure is *loud*. The under-broad half sat behind the
    same repair for five iters and would have read as a clean bill of health indefinitely, because an
    unprobed service produces **no line at all**. When grading a scope, ask both questions — *what does
    it probe that it should not*, and *what runs that it never looks at* — and expect only the first to
    announce itself.

    **Corollary — an intersection must NAME what it drops.** Unioning the override in is not enough: a
    service the stack runs for which the probe registry has no row (measured: `hiring-app`, and
    Clerkenstein's `fake-fapi` / `fake-bapi`, whose death makes every login on the stack fail) cannot be
    graded. Dropping it silently converts *"running and ungraded"* into *"absent"*. Print it, and
    **declare it on both sides** so an arrival or a stale declaration goes RED — the same
    declared-not-inferred move as `SERVICES_NOT_IN_PLATFORM_COMPOSE` (`§5` rule 69's fence discipline;
    `D-M257x-151-1`: a fence whose absent arm reads a comment cannot fire).

    **Corollary — retiring a gap-disclosure fence.** When a repair closes the gap, the fence that pinned
    the *disclosure* of that gap will fail, and deleting it is the wrong move: it retires a real
    property along with an obsolete spelling of it. Re-point it at the property one level up. Here the
    old test asserted `generate.sh`'s **source** contained three service-name literals; the repair
    deliberately removed them (a hand-written exclusion list is the defect class). The re-pointed test
    asserts the stronger pair — the disclosure block carries **no** service-name literal, **and** the
    real script, when RUN, still names what it left out.

71. **A fence that QUOTES the line it guards fails identically on improvement and on breakage — write it
    against behaviour.** Found twice in two consecutive iters, and neither time by review; both times by
    improving the code the fence guarded.

    - iter-153: harden pass 35's fence asserted `generate.sh`'s **source** contained the literals
      `next-web-app` / `studio-desk` / `directus`. The repair removed them **deliberately** — a
      hand-written list of what a mechanism excludes is the defect class (rule 70).
    - iter-154: `dev-stack`'s own contract test asserted the literal source line
      `[ "$local_content" = 1 ] && verify_svcs="$verify_svcs directus"`. The repair replaced that
      decision with a read of the artifact that already made it.

    Both fences protected a **real** property. Both encoded it as the **current spelling** of the code
    that happened to implement it, so a correct rewrite of the subject produced a RED that reads exactly
    like a regression. The remedy in both cases was the same and it is not deletion: **re-point to the
    property.** Where a body-text contract test genuinely cannot execute its subject, keep only the
    structural half it can honestly assert (*this script calls X*; *no line here names a service*) and
    move the behavioural half to a fence that **extracts and runs** the block — `§5` rule 19's technique,
    which iter-147 established and iters 153–154 both used.

    **The test to apply when writing one:** *would this assertion survive a CORRECT rewrite of its
    subject?* If not, it is pinned to a spelling. Note this is not the same question as *"is the marker
    anchored"* — anchoring is a mechanism, the property is what survives, and the two come apart
    (`SURVEY-M257x-iter152-other-guards-may-read-prose-as-data`).

    **The prescribed repair is STRUCTURAL, not vigilance — added at iter-155, because this rule caught
    its own author one iter after he wrote it.** iter-154 shipped a fence alongside rule 71 that asserted,
    as literals, the current state of the world; iter-155's *correct* change made those literals false and
    the fence went RED indistinguishably from a regression. **Writing the rule did not prevent the
    mistake.** So the instruction is not "be careful": **derive the expectation from the same source the
    code derives from**, at test time. A fence whose expected value comes out of the same generator, the
    same registry or the same compose file its subject reads is correct on both sides of any correct
    change to that subject — and it is still RED for a wrong one, which is the whole point. **Four
    confirmed instances in three iters** (harden pass 35's disclosure fence; `dev-stack`'s contract test;
    the iter-154 fence; `test_verify.py`'s frontend-scope fence), none found by review, all four found by
    improving the code they guarded.

    **Corollary — an omission from a "not re-run" list reads as coverage.** Rule 60 requires a scoped run
    to name what it did NOT cover. iter-154 named two sections as not-re-run and silently omitted a third,
    where a stale fence was already RED; it surfaced an iter later looking like a new regression. **Name
    every section you did not run, not only the ones you thought about.**

    **Corollary — a site that LOOKS derived can carry a hand-written half.** Both bring-up verify tails
    opened with a correct `platform_topology.py` derivation and then hand-appended a conditional tuple.
    That opening is why neither was censused: the *base* set was repaired twice (iter-55, harden pass 3)
    and the *conditional* set was never derived at all. **Grade the whole expression, not the first
    assignment** — and when a repair's own comment names its twin, the fence over the pair is due **in
    that commit**, not at the next breakage (rule 69).

72. **A reporting layer must tell its subject's own voice from everything else on the wire — and
    "everything else" includes the interpreter.** M257x iter-156. `guard_family` merged each member's
    stdout and stderr and reported the **last line of the merged stream** as a GREEN member's verdict.
    `claim_census_guard` emitted a `DeprecationWarning` (positional `maxsplit`, py3.13+), CPython echoed
    the offending **source line** to stderr after the guard had finished speaking, and the family printed
    that line of Python as the guard's verdict — while `claim-census: OK — ratchet holds over 41 files`
    was invisible in the one view that claims to summarise the family. Nothing the guard wrote was wrong.
    **The runner read text the guard did not author as the guard's datum** — iter-152's prose-as-data
    class, one layer up from the corpus, and the census that found it was pointed at the corpus first and
    came back clean. **Look at the layer that REPORTS, not only the layer that reads.**

    The repair is a *derived speaker test*, never a heuristic: guards already print their summary
    flush-left prefixed with their own name, a convention asserted in a docstring since iter-87 and
    checked by nothing. A "does it look like a summary" test would re-create the defect one remove up —
    a warning's echo looks like whatever the source line looks like. **And the rung must be returned**:
    without it, *"the guard summarised itself"* and *"no summary was found, so here is the last thing on
    the wire"* print identically (rule 70's line-3 lesson).

    **Corollary — the RED path is where this bites hardest.** A warning echo is *indented*, which is
    exactly the shape a findings-headline selects on. The same defect that mislabels a GREEN verdict will
    report **a line of Python as a guard's first finding** — re-entering through a door the headline's
    own author could not see.

    **Corollary — declare the noise; do not drop it, and do not grade it.** Discarding non-subject output
    is the same swallow in the other direction, and is how this warning survived four releases *while
    being printed as a verdict*. But a warning is not a finding: turning it RED would be the runner
    inventing a verdict. **The run discloses, the fence gates** — `D-M255-1`'s two-contract precedent.

    **Corollary — a fence whose verdict depends on the interpreter must say which one it measured.**
    This defect exists on py3.13+ and not on py3.9, and the only interpreter on the host carrying pytest
    was 3.9: **the fence's first full run was 18/18 green while the defect was live.** State the
    environment (rule 51's discipline), and prove the mechanism *interpreter-independently* — here by
    driving the runner against a synthetic member that writes to stderr on purpose — so the live census
    is the scoped evidence it is and not a claim about all interpreters.

    **Corollary — the instrument is a claim too, and a census's raw signal is not its finding.** This
    iter's own enumerator first reported **13 of 32 guard modules as `IMPORT-FAIL`**; the cause was the
    loader (a `@dataclass` resolves its module through `sys.modules`, which `module_from_spec` does not
    populate), not any guard. And the structural-position sweep flagged **105 of 171 patterns**, almost
    all of it helpers applied to already-segmented text. Neither number was published as a defect count.
    iter-138's withdrawn 127 and iter-150's 30-to-1 are the same lesson: **hand-grade before you publish,
    and report the raw signal as raw signal.**

73. **A GLOB IS NOT A DERIVATION — and a partition with no third bucket reports its gap as a pass.**
    M257x iter-157. `repair_postcondition` selected its fence registry with `glob("*_guard.py")` while
    **both** claims about that selection were written in terms of the DECLARATION: its own docstring
    (*"a fence added tomorrow enrols itself or makes this fail loudly naming its own filename"*) and
    `guard_family.py:67` (*"`FENCE_KIND` — read STATICALLY by `repair_postcondition.py`"*). Measured:
    **25 modules declared a `FENCE_KIND` and 23 were enumerated.** The two that were not
    (`guard_family`, `predicate_enumerator`) neither enrolled nor failed — their filenames did not end in
    `_guard`, so they fell through a partition with **no third bucket** and were, in the function's own
    output, **indistinguishable from a module that is not a fence at all**. The family's report carried
    the arithmetic in plain sight — `5 participating … 18 standalone` against 25 declarations — and
    nothing subtracted.

    **A glob is a hand-written predicate with a wildcard in it.** It fails the way a hand-written list
    fails, and unlike a list it *looks* derived — rule 71's looks-derived corollary at a different
    construct. When a mechanism's prose says *declaration*, the code must select on the **declaration**.

    **The repair is iter-150's split, and it generalises: keep the partition DECLARED, derive its
    COMPLETENESS.** Which bucket a thing belongs in is a judgement no parse can make; whether every
    member of the domain reached a bucket is arithmetic. Report the residue as a finding that names it —
    never as a silent skip, and never as a `could-not-run` that would suppress the verdicts the
    mechanism exists to produce.

    **Corollary — widen a registry only when the consumer's verdict is demonstrated unchanged.** Both
    newly-enrolled modules declare `standalone`, so neither is asked for sites and the ratchet reads
    identically. Had either declared `postcondition`, the widening would have moved a published baseline
    and belonged in its own iter.

    **Corollary — the naming REQUIREMENT must survive the widening, and the widening creates a new
    refusal.** Enrolling by declaration could silently replace *"a `*_guard.py` must declare a kind"*.
    It must not; and it adds an edge that did not exist before — a **non**-guard-named module declaring
    an *illegal* kind, which the new skip branch would otherwise swallow. Fence the branch that may pass
    in silence, so widening it later is a deliberate act rather than a drift.

    **Corollary — anchors are load-bearing in BOTH directions, and the API decides which.** This iter's
    own fence asserted `assertRegex(src, r"^FENCE_KIND\s*=")` — `assertRegex` applies `re.search`
    **without** `re.M`, so `^` meant *start of file* and the arm reported a false NEGATIVE: *"guard_family
    declares nothing"*. iter-152's unanchored-`search()` defect in mirror image, inside the fence written
    for this class. Compile the pattern once with the flags it needs and reuse it; care is not the
    defence, a shared compiled constant is.

74. **A FALSE RED WEARING AN OPEN ROUTE'S SIGNATURE IS A DECOY — grade the red before you act on the
    route.** M257x iter-169. `test_m255_mutation_battery` was found **RED at HEAD**, and its failure mode —
    a BASELINE RED with no attributable test — is the exact published signature of
    `FIX-M257x-iter111-staged-battery-dependency-is-underived`, the class whose **last open member was that
    very battery**. Every prior occurrence of that class had been closed by appending one filename to a
    stage list, so the available move was obvious, cheap, and **would have been wrong**: the red came from
    `test_buildbench`'s sampler regression test asserting `callable(Thread._stop)`, and **CPython 3.14
    removed `Thread._stop`**. The subject was innocent, the fence was rotted, and the open route was a
    coincidence of symptom. **A route predicts a cause; it does not certify one.** Reproduce the failure
    *outside* the staged tree first — one command here, and it separated the two immediately (the same
    test failed in place, so staging was never implicated).

    **Corollary — the repair pins the PROPERTY, and the anti-vacuity control is what proves it.** The
    replacement asserted "no name `threading.Thread` occupies", computed rather than listed (rule 70/71).
    Its first form compared `set(vars(sub)) - set(vars(bare))` and the control caught it **in the same
    minute**: `_target` is set by `Thread.__init__` itself, so the subtraction deleted exactly the
    collisions worth catching, and a deliberately-shadowing subclass passed. Ownership is about **who
    assigned**, not who ends up in `__dict__`. A control that only ever confirms is not a control.

75. **"THE SUITE" IS NOT ONE COMMAND — this repo's Python tests need TWO interpreters, and neither runs
    the whole population.** M257x iter-170, censusing all **110** `test_*.py` modules across the five
    Python sections. `/usr/bin/python3` is **3.9.6** and is the only interpreter on the box **with
    pytest** — the fleet runner. The working interpreter is **3.14.6** and has **no pytest at all**.
    Measured, the two **disagree about four modules**: two `import pytest` and cannot load under 3.14; one
    relied on pytest putting a test's own directory on `sys.path`; one binds a server and fails only under
    the other. **A green from one runner is not a green** — rule 60 with the scope being the *interpreter*,
    which nothing had named as a scope before. The population under the fleet runner is **3,332 tests, 0
    collection errors**; under 3.14/unittest, **101 GREEN · 8 RED · 1 TIMEOUT** over 3,279 executed — and
    **4 of those 8 REDs are runner artifacts, not defects.** State the runner with every suite number, the
    way `§8` already requires stating the host with every timing number.

    **Corollary — and the instrument committed the defect it was built to prevent.** The census's third
    bucket (*needs a live stack or clone — neither green nor a defect*, `§5` rule 73) was first sniffed
    from **error-message substrings**, and it returned **ZERO against nine genuinely environment-gated
    failures**. A third bucket that never fires is a two-bucket partition wearing three labels. The
    repair is the split rule 73 already prescribes: **keep the partition DECLARED, derive its
    COMPLETENESS** — the nine are named in `stack-core/suite_census.py`'s `ENV_GATED`, a declared entry
    whose test no longer exists is reported STALE, and an undeclared RED is reported ACTIONABLE. Six of
    the nine are `FIX-M257x-iter145-sha-baseline-drift`, which is a **freshness signal and must not be
    re-pinned away**.

76. **AN UNEXPLAINED RUNNER DISAGREEMENT IS A SHIPPED DEFECT UNTIL SOMEONE PROVES IT IS AN ARTIFACT —
    "harness assumption" is a hypothesis, not a disposition.** M257x iter-171, closing the ONE
    disagreement rule 75 could not attribute to imports. Two `test_cockpit` tests bound a server, waited
    2 s, and failed under 3.14 while passing under 3.9.6. The plausible reading was a tight harness
    window. **The actual cause was in shipped code, on the platform's own critical path:** CPython's
    `http.server.HTTPServer.server_bind` sets `server_name = socket.getfqdn(host)` — a synchronous
    **reverse-DNS query** that must answer before `serve_forever` is reached. Same Mac, same address,
    same instant:

    | interpreter | `socket.getfqdn("127.0.0.1")` | cold | warm |
    |---|---|---|---|
    | `/usr/bin/python3` 3.9.6 (Apple system) | `'1.0.0.127.in-addr.arpa'` | **0.005 s** | 0.001 s |
    | `python3` 3.14.6 (homebrew) | `'localhost'` | **35.005 s** | 0.000 s |

    **Four orders of magnitude, from a call nothing reads.** `server_name` exists to fill the CGI
    `SERVER_NAME`; the cockpit builds no CGI environment. And the runtime consequence is not slowness:
    `up-injected.sh` polls the cockpit's `/healthz` **25 × 0.2 s ≈ 5 s**, so a 35 s bind presents as
    *"presenter cockpit FAILED to come up … there is NO working cockpit"*, **skips fronting it with
    `tailscale serve`**, and the cockpit then serves fine — unfronted — half a minute later. **A false
    negative, not a delay**, on the demo's single entry point.

    **The general shape, and it is `§8`'s twin.** `§8` says *state the environment with every number*
    because the same code costs 4.84 GB on one host and 2.88 GB on another. Rule 75 extended that to
    pass/fail. This rule closes the loop: when two environments disagree, **the disagreement is
    evidence about the CODE**, and the cheap dismissal ("that runner is weird") is the one reading that
    guarantees nothing gets learned. Dump the blocked thread's stack — `sys._current_frames()` — before
    theorising; it took one call here and named the frame outright.

    **Warm caches hide it, which is why the fence poisons the resolver rather than timing the bind.**
    The second call is free on both interpreters, so any measurement taken after a warm-up shows
    nothing. `demo-stack/tests/test_bind_no_reverse_dns.py` makes `socket.getfqdn` **raise**, so the
    predicate is *"was the resolver consulted?"* — cost-free, cache-independent, and with CPython's own
    class as the mutation control that must still raise. Censused by **property**
    (`cls.server_bind is not HTTPServer.server_bind`), never by spelling: at iter-171 exactly **4** `.py`
    files in all of `rosetta-extensions` mention `HTTPServer`, carrying **3** reachable server classes
    and **13** construction sites — all routed through the one fixed class. `socketserver.TCPServer`
    sites are out of the population **by property** (its `server_bind` calls `getsockname()` and stops),
    not by exemption.

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

### The platform's CONFIG is its documentation of record; its NARRATIVE docs are not (M257x iter-60)

When a row needs adjudicating, go to the platform's **configuration files** — not to its prose, and never to
another document of ours. Measured, not asserted: config is edited **in the same commit as the change** and
carries the rationale inline, because the person making the change has to touch it either way.

| artifact | what it says, in its own words |
|---|---|
| `repos.yml` header | *"`app` is the ONLY repo with migrations to run… they own no local schema."* |
| `docker-compose.yml:84-92` | *"messenger-in-app (v9.0) and customerio-sync-in-app are folded into this container too, but deliberately have NO variables here… Pinning them to `false` here would override .env and make opting in impossible without editing this file."* |
| `docker-compose.yml:102-103` | *"storage, messenger and customerio-sync are not services any more — this one container serves all three in-process."* |
| the commit subject at `0dab54d` | *"run without the standalone storage; rename graphql -> core"* |

> **The row this table used to carry is itself the lesson (M257x iter-87).** Until `838d907` the middle
> row cited `docker-compose.yml:130-133` @ `0dab54d`, where the `storage` service block carried its own
> in-comment warning that the container had left the default profiles and that starting it beside
> `backend` would put two writers onto one bucket. That block is **deleted**, so the anchor now points
> past the end of the construct it named, and the hazard it described is unreachable rather than merely
> discouraged. **The example is stronger for having expired.** The platform's configuration told the
> truth at that ref *and* at this one — four days apart, in the same file, each time in the commit that
> made the change. Its prose did neither.

Its narrative docs lag and are partly unmeasured. App `v1.366.0`'s own `knowledge/*.md` asserts *"60K+ skills"*
with no measurement attached, and **the repo contains no job-role count anywhere** — so the long-quoted
"18K roles" has **zero upstream provenance in the very repo we would be deferring to**. Deferring to their prose
would have imported a figure this corpus had already refuted by measurement.

**And two documents that agree are not two witnesses.** A read-only reconciliation of an entire external
documentation PR against this corpus returned **92 claims absorbed · 30 superseded · 5 contradictions standing ·
0 refuted · ZERO new information** — while every live defect found that week sat exactly where the two documents
*agreed*. Diffing two documents is structurally incapable of finding what they share. **Adjudicate against
platform artifacts.**

**Corollary — a half-landed fold needs a state, and it is recorded on both sides or not at all.** Stated as a
gap at iter-59 and **closed at iter-64**: the map's vocabulary now has an **eighth** token, `mid-fold`, and
`storage` **was** its instance, at `app` `b948604` v1.366.0: config set `STORAGE_RPC_ADDR` in no compose file
and not in `.env_example` (0 occurrences), the service had moved to `profiles: [storage-legacy]`, `repos.yml`
still cloned it, and `app` still read it at `main.go:446`, `:524`, `:992` and in three `cmd/` tools — two of
which hard-required it (`academyImport/main.go:235`, `academy-asset-upload/main.go:133`) while
`cmd/import/main.go:50` built a client against the empty string. Neither `live-standalone` nor
`merged-into-app`, and one side alone is not a claim — the config side read as "removed", the consumer side as
"live". Record both, cited, or record neither.

**The state's life was four iterations, and that is the lesson, not a footnote.** At `app` `9d00a313` v1.367.0
— 56 commits and one working morning later — `STORAGE_RPC_ADDR` is read by **nothing**: a Go grep at that
ref returns **3 hits, every one a comment**, and `main.go:451` says so in words. (At the older `b948604` it
is read by `main.go` **and** all three CLIs — 7 env lookups. "Read by `main.go` but by none of the CLIs"
was a middle state that never existed; corrected M257x iter-85.) Both prod counts are `0`; the fold is
complete on both sides and the row is `merged-into-app`
(M257x iter-68). **No row carries `mid-fold` today.** The token stays, because the fold program is not
finished and a state you can only name after you need it is the state you will get wrong. (Worked example: [`storage.md`](../services/storage.md); the
state itself is fenced by `platform_alignment_guard.py` assertion C, and the variable by
`platform_predicate_guard.py` G6, which refuses a mid-fold variable that no document cites a read site for.)

> **A vocabulary gap is a claim the map cannot make.** Between iter-59 and iter-64 the map called `storage`
> `live-standalone` on both sides — not because anyone believed it, but because the alternative was to invent
> a token in a fenced field. **The fence had eight things to say and seven words.** When a measurement has no
> legal way to be recorded, widen the vocabulary in the same iter that measures it, or the measurement lands
> in a service doc and the fenced artifact keeps the old answer.

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

   **Schema-safety and CITATION-safety are unrelated properties, and this rule only ever measured the
   first (M257x iter-58 → `D-M257x-59-3`).** That advance was vetted exactly as written above — 0
   migrations, 0 destructive DDL, 0 newly hard-required config, all three correct — and it still moved
   **22 of 23** corpus `main.go:N` citations into the wrong construct. The fence that watches citations
   caught **1**: a **4.5% catch rate**, because a whole-file line shift moves nearly every citation by the
   same delta and only the one that crossed a construct boundary is detectable as *wrong* rather than
   merely *moved*.

   So: **before taking an advance, re-resolve every corpus citation whose path lands in the advancing
   repo, at the new ref, and record moved / dead / held.** It is bounded and cheap — the whole set was
   **23**. The repair belongs to the **advancing iter**, for the same reason the ref rule already gives:
   *the iter that detects the move re-points, in that iter.* Deferred, it becomes a second milestone's
   problem and the citations rot in the interval.

   **4c — you do not have to TAKE an advance to measure it, and you should measure it first
   (M257x iter-68 → `D-M257x-68-1`).** `git show <ref>:<path>` reads any ref without touching a
   checkout, so the whole delta is available before a single file moves — which matters when the clone
   is *pinned on purpose* (a demo stack pins its build refs; changing the checkout would fight the pin
   and the running stack). iter-68 measured an `app` advance of **56 commits landed the same morning**
   and found that **25 of the 42 corpus citations holding at the pinned ref break at origin HEAD — 60%
   in one working day.** Fifteen minutes of measurement redirected the entire iteration; repairing
   against the pinned ref first would have been careful, correct-looking work that was false before it
   was committed.

   **4d — a guard that resolves a citation against a CHECKOUT has no ref, and its verdict is not a
   measurement (M257x iter-68).** Three fences were found doing it in one iteration —
   `platform_predicate_guard` (G6's consumer side), `platform_alignment_guard` (assertion F),
   `anchor_construct_guard` (every anchor into a clone) — each calling `read_text()` on whatever the
   clone happened to have checked out, and **none of them reporting which file that was**. The same
   corpus read **GREEN at origin HEAD and 4-findings RED at the pinned build ref**; both answers look
   like answers. The rule:

   - resolve **and** read at a **named ref**, and **print the provenance on every run**;
   - **existence is decided at the same ref the content is read at** — resolving against the checkout
     while reading against a ref makes a file *born* in the advance read as an unresolvable path, an
     instrument gap wearing the costume of a corpus defect;
   - an **untracked** worktree file is in no ref and correctly leaves scope;
   - a ref the caller **named** that does not resolve is **UNMEASURED**, never silently substituted
     (§5 rule 7).

   P1 — *state your refs* — is not satisfied by a number whose ref is a checkout.
5. **Prove it cold.** `demo-down --purge` + `demo-up`. A warm cycle hid a four-day-old total breakage.

---

## 8. Fence — so it cannot silently recur

Three layers. Each must be **watched going RED** before it is trusted — M256 found **43** checks that reported
success without checking.

| layer | asserts | lives in |
|---|---|---|
| map ↔ `repos.yml`, both ways | every `repos.yml` service has a map row; no map row invents a repo; every state is one of the **eight** (`mid-fold` added M257x iter-64); every row cites a sha or `file:line`; the net-new census does not overlap the clone set | `stack-core/platform_alignment_guard.py` (M257x iter-20) — landed, with `tests/test_platform_alignment_guard.py`. **What runs unconditionally is C+D+E** (vocabulary, citations, census overlap): those are properties of the map alone, so they are checked against the **real** map on every suite run, everywhere. **A and B need a real `repos.yml`, and a platform clone is git-ignored and ephemeral** — the test searches `PLATFORM_REPOS_YML` then every `stack-*/platform/repos.yml`, and *skips A/B/E-against-the-clone-set* if it finds none, naming what it looked for. Until M257x harden pass 5 it looked at one hardcoded path and this row claimed the whole thing ran "on every suite run" — on a box with no demo stack it silently skipped, i.e. the fence's own documentation was the claim-without-a-measurement shape this protocol exists to catch |
| static schema fence | every schema a seeder WRITES to **through a statically-visible construct** is one the migrate step CREATES | `stack-core/tests/test_write_target_schema_fence.py` (M257x iter-06) — reads the legal set from `repos_yml_schemas_to_create`, so it **names no dead schema at all** |
| live schema assert | every schema rext writes exists in `information_schema.schemata` on the migrated stack | bring-up / autoverify (precedent: `dev-stack/tests/test_migrate_dev_live.py:144`) |

### A fifth layer, and it watches the INFLOW rather than the stock (M257x iter-106)

The four layers above all answer *"is what the corpus says true?"* None answers *"has the thing it describes
moved since anyone looked?"* — and M257x iter-103 measured what that costs: of the `N = 33` false anchors a
14-seat double reading found, **20 — 61 % — were clone-advance drift**: a version literal, a `go.mod` pin, a
symbol name, a line offset. `TOK-06` ranked this fence **above repairing those 20**, because repairing them
without a fence re-arms the class at the next clone advance, which is the mechanism that produced them.

| layer | asserts | lives in |
|---|---|---|
| cited-clone advance | every cloned repo the corpus cites **by sha** has at least one cited sha equal to its current HEAD — i.e. no repo has moved past everything the corpus knows about it. Plus a conservative pin check: a site naming `` `<repo>/go.mod` `` **and** a `<module> v<semver>` must agree with that `go.mod` | `stack-core/clone_drift_guard.py` (M257x iter-106), with `tests/test_clone_drift_guard.py` + a 7-mutant battery |

**Three design decisions worth carrying to the next fence of this kind:**

1. **It does not adjudicate truth-at-a-ref, and refusing to was the whole design.** The obvious rule — *a
   version the corpus states must equal the clone's* — cries wolf on the live tree, and a fence that cries
   wolf gets suppressed. The worked example, **as the corpus stood at `e6aed2e`**: `shared_libraries.md`'s
   proto row stated `sentinel v1.200.0` and cited `sentinel/go.mod:9 @ 88bc5592`; **at that ref it was
   true.** (**Repaired at iter-108** — that row now reads `v1.210.0` @ `f2c46190`, so the example is
   history; it is kept because the *design* argument does not depend on the claim still being live.)
   §5 rules 41/44 make it a ref-scoped claim,
   so the fence reports the ADVANCE — *this repo moved past everything you cite, by N commits, and here are
   the citing sites* — which is mechanically true and accuses no sentence. A site pinned to a ref the clone
   is not at is **UNMEASURED and named**, never graded false.
2. **The baseline is DERIVED from the corpus's own citations, so there is no baseline file.** A checked-in
   `repo → last-reconciled-sha` map is §2's hand-maintained tuple in a new costume: it drifts, it gets
   re-accepted absent-mindedly, and its first value has to be *asserted* rather than measured. Instead every
   backticked sha in `corpus/**` is resolved with `git cat-file -t` against every clone — **a sha resolves
   in exactly the repo that contains it**, so attribution is exact and needs no naming convention, no list,
   and nothing to keep in step.
3. **State the REACH in the OK line** (§5 rule 46). This fence catches *"a whole repo advanced unreviewed"*.
   It does **not** catch *"one site cites HEAD while five others are stale"* — one fresh citation reconciles
   the repo. The green says so in its own words, because otherwise "no drift" reads as "the corpus is
   current," which is a different and much larger claim.

**What it caught on its first committed run, which is also the evidence it works.** One RED, zero false
positives: `sentinel` at `f2c46190`, **2 commits past** the newest sha the corpus cites, 5 citing sites. Those
two commits are `chore(deps): update dependencies to latest versions` and a version bump — moving colony
`v0.34.3 → v0.35.2` and proto `v1.200.0 → v1.210.0`. **Both of iter-103's booked pin-drift predicates** —
clerkenstein's *"sentinel is still on `v0.34.3`"* and shared_libraries' *"the live skew is two …
`sentinel v1.200.0`"*, **both quoted as they stood at `e6aed2e`** — are downstream of that single advance,
and the fence found it **without parsing one sentence**. **Both were repaired at iter-108** (colony split
CLOSED, proto skew ZERO); the two claims are quoted here as the fence's evidence, not as live corpus text,
which is why they carry a ref and no line number.

**The limit it also measured, and it is a corpus-side finding.** The conservative pin rule graded exactly
**1** pin, because the corpus writes pins as `<repo> <version>` with the module implied by a table heading —
a form no line-scoped fence can read. **Write a checkable pin claim as `<module> <version>
(`<repo>/go.mod:N`)`**, all three on one line. That is the fence-facing half of §5 rule 44: naming the tree
and the ref is what makes a claim settleable, and putting them where the checker can see them is what makes
it *checked*.

### A sixth layer, on the OTHER inflow: the repair's own induction (M257x iter-107)

The fifth layer watches the platform moving under us. This one watches **us moving things under ourselves**.
`TOK-06` counted it at **21 % of the residual**, and its largest measured shape has now occurred twice by
the same mechanism, one cycle apart, with §5 rule 34 already naming it both times:

* **iter-100 → booked by iter-101.** A two-line parenthetical pushed a `service_taxonomy.md` table down two
  rows; the numbers around it did not move, so a note that had been exactly correct came to cite Chronos and
  Intelligence.
* **iter-102 → booked by iter-103.** `architecture_overview.md` **line 321 as it stood at `8f04d3a`** was
  the correct local-stack line; an inserted production-topology block moved that wording down, and **4 sites
  went on citing the old number**, by then naming the opposite topology. The 14-seat double reading found 2
  of the 3 in-scope ones and **missed `backend.md:54` in both passes**, inside a seat's own file set.
  **Repaired at iter-108**, which re-measured the wording to **`architecture_overview.md` line 335 as it
  stood at that repair** — named in full here because a bare `:NNN` written after a *different* file's
  citation binds to that other document, which is how this very line resolved against `backend.md` instead
  until M257x iter-115 (and false-RED'd the moment a repair shifted that file) — and re-pointed all four
  citers
  (`backend.md`, `sentinel.md`, `jobsimulation.md`, `CLAUDE.md`). **The line numbers in this bullet are
  HISTORICAL and are deliberately written with their ref** — an unpinned `file:line` in a post-mortem is
  indistinguishable from a live citation, and `repair_postcondition` correctly refused this paragraph when
  it carried bare ones.

| layer | asserts | lives in |
|---|---|---|
| repair-induced anchor rot | over a revision range: no commit moves a line out from under an intra-corpus `` `<doc>.md:<N>` `` citation it did not also update. A citation the commit *authored beside a shift it made* is reported **CANNOT-TELL** | `stack-core/anchor_offset_guard.py` (M257x iter-107), with `tests/test_anchor_offset_guard.py` — whose controls are **the two real commits**, replayed |

**Three things it taught, and two of them came from its own controls refuting its first design:**

1. **The answer key is the commit, not a fixture.** Replaying `cd16967` (iter-102) surfaces **all four**
   `:321` citers — including the one both reading passes missed — and replaying `a229f8d` (iter-100)
   surfaces its `service_taxonomy.md` induction. A fence for a defect that has actually happened should be
   graded against the commit that caused it; a fixture only proves the fence does what its author expected.
2. **A file-level carve-out returns GREEN over the defect.** The first cut waived any citation in a file the
   commit had touched — reasonable, and wrong: iter-102 was a **98-site repair** that modified all three
   citing service docs while editing *other* claims in them. It graded 2 citations and passed. The carve-out
   must be **line-level**: the commit must have written the line the citation is on.
3. **Some of it is genuinely undecidable, and the fence says so rather than guessing.** When a commit authors
   a citation *and* shifts that target line, `X is post-move and correct` and `X is pre-move and stale` are
   both consistent with everything the diff records — a synthetic control proved it by showing a **correct**
   re-point (`:7 → :9`) is indistinguishable from iter-102's stale one. So that class is **counted, named,
   and excluded from the exit code**, and the OK line states that the green does not cover it (§8's
   *grade the cannot-tell*, iter-91). **A fence must not assert what it cannot decide** — the alternative is
   a RED that correct repairs trigger, which is a fence that gets suppressed.

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

7. **A test that asserts a LIVE defect has an expiry date, and the thing that expires is the assertion —
   not the defect.** M257x iter-45 built three fences and asserted all five of their target blockers
   against the live corpus. Every assertion passed. Every one of them would have **failed on the next
   iteration**, because that iteration's entire job is repairing those five sites — and the obvious way to
   restore a green suite would have been to edit the fence's own test to match the repair. That is how a
   fence stops asserting anything: not by being deleted, but by being *maintained*.

   > **Rule.** What must go RED goes RED against a **captured answer key**. The live tree is asserted only
   > on properties that **survive the repair** — that the resolver still resolves, that unmeasured inputs
   > are still named, that the ratchet still grades the tree. If you cannot state why an assertion will
   > still be true after the defect is fixed, it belongs in the fixture.

   Two corollaries the same iteration paid for. **Capture the fixture in the shape the fence reads.**
   iter-43's claim-twin fixture could snapshot ±2-line neighbourhoods because its fence matches text;
   two of iter-45's five defects are *relationships between line numbers* — `:788` citing `:447` in the
   same file, `:110` citing line 815 of another — and an excerpt destroys precisely the property under
   test. The fixture there had to be a pair of **line-faithful trees**, with the one large vendored source
   copied under line-preserving elision rather than compacted. And **the green twin must still be reached**:
   assert not only that the repaired tree is silent but that it was *resolved* and *measured*, because a
   fence that stopped reaching a site prints exactly the same clean run as one that passed it (§5 rule 8,
   arriving through the fixture instead of through the code).

   A last, smaller one, from the same iteration: a fixture is a **captured copy of someone else's source**,
   not this repo's. Another fence walked it, reported `stack-core` as an unclassified Go-bearing rext
   section — a section that ships no Go, over a file belonging to the platform — and would next have scored
   the platform's source for rext's schema writes. **Prune `tests/fixtures/` from any walk that derives
   what this repository contains.**

   > **⚠️ This rule was violated by its own milestone four iterations after it was written — and by the
   > author of the fence it was protecting.** M257x iter-49 shipped `--audit-commit` with five end-to-end
   > tests driven against the LIVE corpus at the commit `D-M257x-48-12` records as blocked. All five
   > passed. The same iteration then repaired the 12 blockers, the ratchet fell to 0 new sites, and all
   > five failed **inside the same session** — with nothing RED there is no admission to make, so the
   > runner returns before the audit path is reached.
   >
   > The tell was available and was missed: *"if you cannot state why an assertion will still be true
   > after the defect is fixed, it belongs in the fixture."* The fix was a **hermetic temp-git-repo
   > fixture** — a corpus file, an audit ledger, and two differently-shaped commits — which never goes
   > stale and exercises the whole pipeline. The live-tree run kept its value as *evidence the mode works
   > on the real commit*; it simply cannot be the durable assertion.
   >
   > **So the rule generalizes beyond fences: it binds any test whose subject is a repairable state of
   > this repository.** And its recurrence is itself the datum — this is the eighth consecutive
   > iteration in which the author of a rule broke it while writing the thing the rule governs.

Static and live are **both** required. Static is the only honest offline check, because every seeder test
asserts against a recording fake `Conn` that accepts any table name — *a fake cannot know a table was
dropped*, which is why 2,617 offline tests passed while the bring-up was broken for four days. Live is the
only check that knows what the migration path actually produced.

### Write the anti-vacuity control against the guard's SUBJECT, not its inputs (M257x iter-94)

§5 rule 8 says *a check that SKIPS reads exactly like a check that PASSES*. Everyone in this milestone has
read it. It was still violated **three times in one session, by three different guards**:

| iter | guard | the control that failed |
|---|---|---|
| 91 | `platform_alignment_guard` | graded **total** resolver failure (`subject_checked == 0`) but not **partial** blindness |
| 93 | `unreadable_repo_claim_guard`'s own live-corpus test | **silently skipped** on a hardcoded `parents[3]` |
| 94 | `story_org_count_guard` | its emptiness control **could never fire** |

The last is the clearest. Its `scan_roots()` returns the corpus **plus rext's own two directories** — and
the guard lives in rext, so those always exist. `if not roots: return 2` was therefore dead code, and a run
whose rosetta half was missing scanned only rext, found nothing to contradict, and printed *"and every doc
agrees"* with `rc=0`. **"Every doc agrees" over zero docs is vacuously true.**

All three are the **same substitution**: the control asked whether the guard had *inputs* — roots existed, a
file was found, *some* citations resolved — when what mattered was whether it had reached its **subject**.

> **Rule.** An anti-vacuity control must assert against the thing the guard makes claims about, not against
> whatever it happened to load. Ask *"did I look at the thing I am about to make a claim about?"* — and
> **print the count**, because a vacuous scan and a real one otherwise produce the identical sentence
> (`OK — all 116 scanned doc(s) agree` vs `OK — … every doc agrees`).

Corollary, from the same iter: **a shared symptom is not a shared cause.** Both `story_org_count_guard` and
`union_apply_guard` returned GREEN against a corpus-less tree; only one was a defect. `union_apply_guard`'s
subject IS rext's manifest set, so it is correct by design, and "fixing" it to honour `--repo-root` would
have broken it in the rext-only checkouts the family is consumed from per-stack.

### Then audit the guards' TESTS the same way — that is where the next three were (M257x harden pass 20)

The rule above was written after three guards were caught passing over nothing. The obvious next question
was not asked for four iterations: **the guards were audited for vacuity; their tests were not.** Swept in
harden pass 20, and the very first one was the worst possible case.

`guard_family.py` exists because *"all six corpus guards exit 0"* was a statement about a list somebody
remembered. Its single load-bearing property is that **a guard which says it checked nothing must not read
as GREEN.** The test named for that property was:

```python
for out in ("repair-leak: CANNOT RUN — no candidate shingles", ...):
    self.assertTrue("CANNOT RUN" in out or "Nothing was checked" in out)
```

Two string literals declared three lines above, asserted to contain substrings of themselves. It never
imported the behaviour and never called `run_one`. **Measured:** blinding the branch to `if False:` left
all 22 tests in the file green. *The runner built to refuse an unearned green was, on its one property,
covered by nothing.*

> **Rule.** A test that never calls the code is a fixture asserting against itself. The cheapest detector
> is the one that found this: **blind the branch and re-run.** A control you have not watched fail is a
> control you have not got — and this applies to the CONTROLS with exactly the force it applies to the
> guards, because a vacuous test and a real one produce the identical green.

Five more followed from the same sweep, and they cluster into three shapes worth naming:

| shape | instances |
|---|---|
| **a flag or hatch that silently does nothing** | `--verify-remote` performed no network call at all without `--platform`; `ALIGNMENT_ALLOW_UNMEASURED=1` promised in its own message to *"RECORD the gap rather than hide it"* and recorded **nothing**, emitting an unqualified `OK` |
| **a verdict published over a graded set of zero** | `repair_reach_guard` printed *"every booked finding was reached"* when every anchor failed to parse; `unreadable_repo_claim_guard`'s "PREMISE LIFTED" was an exit 0 the family counted as green |
| **a summary sentence contradicting the line above it** | `guard-family: N guard(s) NOT RUN and accepted` followed immediately by `OK — every member of the census was run and returned green` |

**The generalisation.** Every one of these is *the same* substitution as the guard-level rule, moved one
layer out: **the thing that reports is not the thing that measured.** So the question to ask of any
reporting path — a flag, an escape hatch, a summary line, a test — is the one the rule above asks of a
guard: *did the thing making this claim actually reach the thing it is claiming about?*

Two operational corollaries, both paid for:

1. **An accepted gap is still a gap.** An escape hatch that suppresses a refusal must print what it
   suppressed, and the verdict sentence must be qualified (`OK WITH AN ACCEPTED GAP … this is NOT a
   whole-map green`). Otherwise the hatch converts an honest UNMEASURED into a quotable green, which is
   the failure the refusal existed to prevent, reachable by one environment variable.
2. **Grade on the exit code, not on a substring of the output, wherever the exit code is available.**
   `guard_family`'s CANNOT-CHECK detection sniffed merged stdout+stderr unconditionally, so a guard that
   exited 1 while *echoing a corpus line* containing the phrase was downgraded from RED to
   could-not-check — and its findings vanished from the one view that summarises the family. Two guards
   in the census echo corpus lines verbatim. The sniff belongs only on `rc == 0`, the case it was for.

**And the recursion is real.** In this same pass, the mutation battery caught the *new* provenance test
passing over nothing: two git repos built by one helper in the same second are byte-identical commits and
share a sha, so `assertIn(corpus_head, out)` was satisfied by the **platform** reference line. iter-94's own
fix — the anti-vacuity control the rule above is named for — had meanwhile been appended *after* its file's
`if __name__ == "__main__"` guard, so direct execution collected 23 tests instead of 25 and printed `OK`.
**A control that could never fire, whose fix was placed where it could never run.** Assume you have done it
too; the only reliable check is mechanical.

### Fencing a document does not fence its PARAPHRASES (M257x iter-92 / iter-93)

The fenced map's `cms` row said, in its own voice:

> Whether that rollback declaration still stands is **not something this map can see** — it never could,
> since infrastructure has never been in the clone set.

Six other documents — `services/backend.md`, `CLAUDE.md`, `architecture/dependency_map.md`,
`services/cms.md`, `architecture/external_services.md`, `services/storage.md` — stated flatly that
`module.cms_euwest1` **is still declared as the rollback path**, a thing **not visible to this corpus** at
all, since `infrastructure` has never been in the clone set. **Only the fenced map carried the hedge.**

This is a limit on the whole fencing method of TOK-02 and TOK-05, and it is worse than an unfenced claim:
**a hedge that survives only in the one file a guard reads implies the system checked, when it did not.**
`claim_twin_guard` fences *adjudicated claims* across the tree; nothing fenced a **hedge**.

Two hard-won riders, both from the iter that named the class:

1. **The restatement count is not guessable.** The estimate at the moment of naming was *"two files."* It
   was six, and it took `repair_leak_guard` **two rounds** to reach the floor: RED on the iter commit
   itself, RED again on the first repair, GREEN on the second. The iter that named the class committed the
   class while fixing it — twice.
2. **`repair_leak_guard`'s scope is the DIFF, so run it again after the repair.** A repair commit
   introduces a fresh candidate set and surfaces sites the previous run could not see. **One green run is
   not the fixpoint; two consecutive greens are.** §8's *run the fence at the COMMIT* said when, not how
   many times.

> **Rule.** When a fenced document hedges a claim *because the evidence is unreachable*, the hedge is part
> of the claim. Fence the **hedge**, tree-wide — not the sentence in the one document that carries it.

iter-93 does that for the case at hand (`unreadable_repo_claim_guard.py`: every `module.*_euwest1` mention
must carry an unmeasurable marker **in its own paragraph**, because a reader lands on a paragraph and not
on a file). The guard **re-measures its own premise on every run** and **retires itself** the moment an
`infrastructure` clone appears — a fence that kept demanding a hedge after the hedge became unnecessary
would be pinning the current shape of our ignorance, which is rule 3 above turned on the fence itself.

### Write the RETRACTION in the vocabulary the fence enumerates (M257x iter-98)

A corpus that documents its own corrections **quotes the refuted sentence verbatim** — that is good writing,
and `claim_twin_guard` accommodates it: a site may be acknowledged in `claim_twin_waivers.json`. But the
waiver is deliberately **half a key**. It is honoured only while `_looks_retracted` independently finds a
retraction marker within 320 characters of the match, so deleting the retraction and leaving the sentence
standing silently re-arms the fence. That is rule 3 applied to waivers.

The consequence is easy to miss until it bites. iter-98 repaired `roadrunner.md` and wrote:

> The old parenthetical *"zero hits outside CHANGELOG"* … **are both FALSE** and are withdrawn

The waiver **did not take**, and the site stayed RED. `RETRACTION_MARKERS` contains `"is false"` and
`"was false"`; it does not contain `"are both false"`. Two ways to close that gap, and **only one is safe**:

- **Widen `RETRACTION_MARKERS`** — in the guard's own recorded words, *"the direction that can hollow a
  fence out."* You would be teaching the fence to accept a phrasing you invented a minute ago.
- **Move the prose** — the sentence became *"is RETRACTED — each is false"*, using a marker the guard
  already knew. Cost: eight words.

**The rule: a retraction a fence cannot recognise is, to every automated reader, an unretracted falsehood.**
When the fence disagrees with your retraction, the default repair is the sentence, not the marker list —
and if you genuinely need a new marker, that is a change to the *specification*, argued and tested against
the answer keys, not a quiet edit to make today's commit green.

**The same asymmetry governs answer-key fixtures.** iter-98's ledger turned `test_02` (*"the green twin of
every site stays SILENT"*) RED, because the **green** fixture `claim_twin/green/17.md` still contained a
sentence refuted 57 iters after that fixture was captured. The fixture was green **only with respect to the
blocker it was built for**. Repairing it was correct — it was violating its own stated contract — and the
control that makes that safe is `test_01`: *all 18 known-bad sites must still fire*, untouched and still
passing. **Editing a GREEN fixture to match a newly-proven truth is maintenance; weakening a ledger quote so
a stale fixture stays quiet is tuning-to-green.** The test that separates them is the positive control, so
never repair a fixture without re-running it.

### A PARTIAL skip is worse than a total one — grade the cannot-tell (M257x iter-91)

§5 rule 8 says *a check that SKIPS reads exactly like a check that PASSES*. Its sharper form: **a check that
skips PART of its work reads like a check that passed, and carries a real verdict to prove it.**

`platform_alignment_guard` had exactly one positive control — `subject_checked == 0`, which trips when the
resolver fails *entirely*. Partial blindness had none. Two things were printed on every run and graded by
nothing:

- **`unresolvable > 0`** — citations the run could not resolve at all. Each is a claim it did not check.
- **the silent worktree fallback** — `cited_text` tries `origin/main`, then `HEAD`, then reads the
  **checkout**, returning provenance `worktree(no-ref)`. The string was always there; nothing read it.

That second one is the dangerous shape, because the two references **disagree**. Measured on the shipped map:

| reference | verdict |
|---|---|
| `auto`, refs present | GREEN — 90 resolved, 0 unresolvable |
| the worktree fallback | RED — 8 findings, **4 unresolvable, ungraded** |

So a clone that could not see the objects it needed was graded against whatever it happened to have, and
told nobody. A RED is not safety here: it *looks* like diligence while 4 citations went unchecked.

> **Rule.** A guard has **three** verdicts, not two: pass, fail, and **UNMEASURED**. Reserve a distinct exit
> code for the third (2 here, which the family runner already renders as CANNOT-RUN), and route every
> partial-blindness signal into it. An accept-the-gap escape hatch is fine — `ALIGNMENT_ALLOW_UNMEASURED=1`,
> `--allow-not-run` — because it **records** the gap; silence does not.

Two riders:

1. **Fix it at the point of use, not in the runner.** Only the guard knows which refs it needs. A
   family-level "is this clone stale" heuristic would rebuild §2's hand-maintained tuple in a new costume —
   and *stale* is not even locally decidable (a clone fetched a minute ago can be behind). Reserve the
   runner's refusal for what IS locally decidable and unambiguous: a platform-facing run against a clone
   with **no `origin/main` at all**.
2. **Print the REFERENCE with every verdict.** The runner named a *directory* and never a *commit*, so every
   `13 GREEN` transcript in this milestone is unre-checkable after the fact — which is exactly how a green
   reading gets quoted forward into a brief. A verdict without its refs is an anecdote.
3. **And "the reference" is THREE trees, not two — the third is the fence's own** *(M257x iter-105)*. Rider 2
   was implemented for the corpus and the platform; the runner still could not say which
   `rosetta-extensions` tree it had loaded, which is the tree that actually decides the verdict. §5 rule 50
   has the incident and the design; the mechanism is `stack-core/fence_provenance.py`, and what keeps it from
   rotting is that **the conformance check is derived from `guard_family.census()` rather than from a list**
   — a new `*_guard.py` that does not state its tree turns `tests/test_fence_provenance.py` RED without
   anyone remembering to add it. Asserted over the **parsed** `if __name__ == "__main__"` block, per the
   *parsed construct, never a whole-file substring* rule above: a `grep` would pass a guard that mentions
   the module in a comment or imports it and never calls it, and both are the shapes a stamp regresses into.

### A verdict with a GRAMMAR states its provenance INSIDE the payload (M257x iter-111)

Rider 3 above is right and was implemented in the wrong place for a third of the family. `stamp()` printed
the tree **first, on stdout** — correct for a text verdict, and **twelve** of these guards also offer
`--json`, where stdout is a document. There the same line does not weaken the verdict, it **destroys** it:
`guard.py --json | jq` dies at char 0 and the consumer gets nothing at all.

**The suite did not see it for six days, because the suite worked around it.** Every test that parsed a
guard's `--json` set `FENCE_PROVENANCE_STAMPED=1` first — undocumented, load-bearing, and nowhere stated.
**A green bought with a hidden condition is the §5 rule 50 class**, and it had arrived inside the fix *for*
rule 50.

> **Rule.** Provenance follows the payload's grammar. **Text** → the line first, on stdout. **Machine** →
> the tree goes **inside the document** (a `fence_tree` key), the human line to stderr. And the mode is
> **derived from the invocation** — `--json` is in `argv` or it is not — never from a flag and never from
> the environment.

Three things this gets right that a flat *"default the stamp to stderr"* would not:

1. **Text mode is untouched**, so iter-105's two real reasons — `run_one` reports `lines[-1]`; `headline()`
   counts finding-shaped lines — keep holding. Both were arguments about **order and shape**; neither was
   ever an argument about the **stream**, which is why the apparent dilemma dissolved on re-reading.
2. **An archived `verdict.json` states its own tree** without the terminal that produced it. A stderr line
   is lost to `> verdict.json`; a stdout preamble is lost to a parser. The document is the only place that
   survives both.
3. **The workaround dies with the defect.** Repairing the pollution and leaving the env var in the tests
   would keep the *mechanism* — an undocumented variable that makes tests pass — alive for the next one.
   Removed at all five call sites and fenced: a scan over every `test_*.py` fails on any **setter**, with
   an anti-vacuity control that writes a synthetic setter and confirms the matcher still sees one.

**And the inverse false promise, which nothing was looking for.** `demo_knob_guard`'s rule is that a
**doc**-promised flag with no parser entry is a false promise. iter-111 found the halves swapped: a guard
**declared** `--json` in its parser and read it **nowhere**, so `--json` parsed, exited 0, and printed
prose. Fenced by walking each `--json`-declaring guard's AST for a read of the flag. **Check both
directions: promised-and-absent, and present-and-unread.**

### A battery that stages a SUBSET carries an underived dependency contract (M257x iter-111)

Five mutation batteries copy a hand-listed subset of `stack-core` into a temp tree and run a fence's suite
there. iter-111 added one module-scope `import` to eleven guards — and **16 tests across all five
batteries went RED at once**, reporting *"the unmutated baseline went RED"*, i.e. **the fence looks
broken** when the true cause is **a file missing from a copy list**.

The batteries already assert the **presence** direction — *"the dependency list names a file that does not
exist"*. There is no assertion in the **absence** direction, and that is the one that fires when a guard
grows an import.

> **Rule.** A staged tree is a **closure**, and a hand-written closure rots the moment the code it stages
> grows a dependency. Either derive it (an import graph, not a grep), or make the staged run's failure
> **name its cause** — an import smoke-check before the battery grades anything, so *"you forgot a file"*
> can never be reported as *"the fence is broken"*.

The second-order lesson is about the write-up, not the code: the tempting reading of 16 simultaneous REDs
is *"a hidden failure, exposed at last"*. It was **caused by the iter that found it**, and it is recorded
that way (`D-M257x-111-5`). The flattering reading is the one this protocol refuses on principle.

### A SILENT test is not a BLOCKED test — and CPU-idle is not evidence when the work is in a child (M257x iter-111)

iter-108 recorded that `stack-core`'s full `pytest tests/` **"BLOCKS INDEFINITELY"** — *"blocked, not
slow"* — on the evidence *"12.6 s CPU over 3 m 43 s, frozen at 442 results, reproduced in two runs"*, and
drew the corollary that the standing suite total **could not be produced on that host**.

**Measured at iter-111: the suite completes.** Three full runs — **414.14 s**, **431.36 s**, and, once the
batteries were repaired so their mutants actually execute instead of aborting on a red baseline,
**1090.88 s (18 m 10 s) for `1 failed, 1011 passed`.** That last number is worth pausing on: the suite got
**2.5× slower by being fixed**, because a battery that dies on its baseline never runs its mutants. *A
fast suite is not evidence of a healthy one.*

The freeze is real and fully accounted for:

- the named test is the suite's slowest at **132.72 s** (pytest `--durations`), and its module takes
  **142.38 s** measured alone;
- it prints **nothing** while running — it is one `-q` dot, and inside it a harness runs a nested suite
  **8 times at ~16 s each** through `subprocess.run(..., timeout=900)`;
- so **low parent CPU is expected, not diagnostic**: the work is in child processes and in waiting;
- watched at 45-second intervals, output sat at 522 bytes for **2 m 15 s** and then advanced.

> **Rule.** Before calling a stall a block, measure **how long the thing is supposed to take**. A
> long-running test that captures its children's output emits nothing and burns no parent CPU — the exact
> signature of a hang, produced by working normally. The cheap discriminators, in order: run the module
> **alone** and time it; ask pytest for `--durations`; sample the **process tree**, not the parent.

**What survives from the retracted finding, and it was always the better half: *state the invocation with
the count*.** A scoped pass must never read as a whole-suite pass. That rule is correct whether or not the
suite completes, and it is unchanged.

### Guards must be tested in PAIRS — a green suite proves each guard, not the specification (M257x iter-90)

Every rule above hardens **one** fence. This one is about what a set of fences can hide from a suite that
tests each of them correctly.

`demopatch` ships **seven** guards, and its suite tested all seven. It was 52-for-52 green while the
mechanism was broken in the most basic way it could be: **the patch applied and would not come off.**

The two guards involved were each individually right.

- **G2** (drift-refuse) was rewritten at M217 so that *the anchor is the contract; the whole-file sha is only
  a baseline* — apply **self-heals** onto a drifted base. Deliberate, and load-bearing: the strict whole-file
  gate is what silently refused two `app` perf patches for four releases and shipped a 76-second members grid.
- **G5** (self-revert) still compared **whole-file shas** against the manifest's recorded baseline.

A self-healed apply produces a file whose sha is **necessarily** neither `pre_sha256` nor `post_sha256`. So
on any clone whose base had moved — the normal state of a persistent, `make pull`ed clone — revert refused
and the clone was left dirty, contradicting the spec's own headline promise (*"the clone is left git-clean"*).

**G2 and G5 cannot both hold once the base is allowed to move, and the base is always allowed to move.** No
test said so, because every test asserted them **separately**:

| test | what it covers | what it never does |
|---|---|---|
| `…g2_drifted_sha_with_an_INTACT_anchor_SELF_HEALS` | apply onto a drifted base | **revert** |
| `…g5_revert_on_drifted_refuses_without_force` | revert a drifted target | one that was ever **applied** |

Neither is wrong. Their **conjunction** is where the specification lived, and nothing evaluated it.

> **Rule.** A specification with *n* guards needs at least one test per **PAIR that can interact**, not one
> test per guard. Enumerate the pairs deliberately; a pair with no test is an unstated claim that the two
> guards are independent — and independence is the thing that fails.

Three riders, each paid for in the same iter:

1. **The interaction can be between INVOCATIONS, not just between guards.** The sharpest pair here is the
   *chain*: two patches on one `urls.ts`, applied studio→pubweb and reverted pubweb→studio. No single-call
   test can see it, whatever it asserts.
2. **A documented defect can be documented as benign and still be live.** This exact asymmetry was already
   written down in `demopatch-spec.md` — and closed with *"it is not currently harmful."* The reasoning was
   true of the `app` build-scratch clone (force-checked-out every bring-up) and false of the persistent
   `next-web` clone the paragraph was about. It then cost M257x iters 88 and 89 in full. **A
   known-and-dismissed finding deserves the same re-derivation as a new one: the dismissal is a claim too.**
3. **When a conjunction test fails, check whether it encodes a real requirement before changing the code to
   satisfy it.** iter-90's first double-revert test failed against the fix; the premise was checked against
   `up-injected.sh:741` (one revert per `RETURN` trap, then `trap - RETURN`) and found false. It was replaced
   by the chain pair rather than satisfied by bending the design. **A design bent to satisfy a wrong test is
   worse than either.**

### A TEARDOWN is a write path too — and a stale override poisons it before the bring-up (M257x iter-55)

Every rule above is about detecting drift on the way **in**. The way **out** was never examined, and it
turned out to be where a stale statement of topology does the most damage.

`rosetta-demo down N --purge` printed a clean teardown and exited **0** with **eleven containers still
running**. The only signal was one line inside compose's own output:

    service "roadrunner" has neither an image nor a build context specified: invalid compose project

A demo's `docker-compose.injected.yml` is generated at bring-up and persists. This one was 45 hours old, from
before `ef32d4c` deleted `roadrunner`/`cms`/`jobsimulation` and `0dab54d` defaulted `storage` out. Overlaid
on the *new* base compose, a stale override contributes only **overrides for services that no longer have a
base definition** — no image, no build — and compose then refuses to act on the **entire project**, including
the eleven services that were perfectly fine.

What followed was worse than the failure, and is the part to internalise:

- the call site was `compose down … || true`, so the refusal was discarded;
- `purge_data_dir` then deleted the data directory **out from under eleven running containers**;
- `data purged` and `removing this stack's images` both printed;
- the registry slot was released, and `cmd_down` exited 0;
- the next bring-up started on top of all of it.

**Three rules follow.**

1. **A teardown must have a post-condition, and it cannot be the teardown's own exit code.** §5 rule 7 says a
   probe must not be able to satisfy itself; `compose down` was treated as its own evidence of teardown. A
   teardown never asked *"did anything survive?"* reports success by not looking. Ask Docker —
   `docker ps -aq --filter label=com.docker.compose.project=<project>` — remove what is named, **re-read**,
   and fail if anything remains.

2. **There is no compose-file-shaped fix, because the compose file is the thing that went stale.** Any repair
   phrased as *"regenerate the override first"* or *"tear down with the old base compose"* is another
   statement about topology that can itself go stale. Prefer the source that cannot: the container labels
   Docker maintains.

3. **`--remove-orphans` does not cover this, and the reason is worth knowing.** It failed twice for two
   different reasons. Once because the project was invalid, so nothing ran at all. And once — on the *fixed*
   run, which caught a surviving `storage` container — because `storage` is still **declared** in the base
   compose under `profiles: [storage-legacy]`, so it is **not an orphan**, while not being in the default
   profile, so it is **not selected** either. A container can be simultaneously not-an-orphan and
   not-selected, and fall through both. Expect one of these per fold, at the moment a service moves to a
   rollback profile rather than being deleted outright.

> **Corollary for the §2 tuple family.** The same iteration found the profile NAME, the verify service set
> and the injected-build set all held as hand-written literals, and the profile one is the most dangerous
> thing in this document: a `--profile` naming a renamed token, run against the renamed platform, is a
> **successful command that starts the always-on floor and nothing else** — measured at `0dab54d`: **three** containers
> (`postgresql`, `redis`, `sentinel`), not zero, because those three declare no `profiles:` key. That is worse
> than zero: Postgres answers, `docker ps` is non-empty, and the stack presents as *partially working*.
> No error, no non-zero exit, nothing for a log-reader to catch. Prefer
> a derivation whose correctness you can test at *two* refs — that is the only evidence that distinguishes a
> derivation from a literal that happens to be right today.

### Fence a registry's COMPLETENESS, never its contents (M257x iter-162)

Several fences in this family are driven by a **registry** — a list of the things the fence looks at.
`frozen_expectation_census.py` is the clean case: its docstring says it matches *"a value **some
non-test module derives**"*, and its first implementation executed **three hand-written entries**
against a population of **125**. Reach was **3 of 53** executable-here derivations — **5.7 %** — and
nothing in the tree measured that, so the census printed *"0 unexempted candidates"* in exactly the
words a complete one would use.

**A hand-list and a census print the same sentence. Only one of them earns it.**

The repair is not a bigger hand-list — that is the same defect with a better number, rotting the same
way. Split the registry in two:

- **Contents stay DECLARED.** Which entries to execute is a judgement (some derivations shell out to
  docker, some return an audit *verdict* rather than a reference set, some would make the instrument
  match its own output). Declare each with a **class and a reason**, per site — that is what makes an
  individual call arguable later.
- **Completeness is DERIVED.** Enumerate the population mechanically, subtract the table, assert the
  remainder empty — **and assert the reverse too**, so a decision for a site that no longer exists is
  caught as fiction rather than carried as reassurance.

Then a member added to the tree tomorrow turns the fence RED **with its own id in the failure
message**, and the fix is one line. iter-162's fired **twice inside its own commit** — once on a
function that entered the population because that commit rewrote it, once on seven declines that
named a class and no reason. Neither would have survived, and neither is the kind of thing a reviewer
looks for.

**Read the reach as a number before believing a zero** (`§9`: *a census that returns ZERO must prove
its instrument* — this is the reach half of that obligation, and it is the half a mutation control
does not cover: a mutation control proves the instrument fires on what it looks at, never that it
looks at everything it claims to).

### When a class is "too semantic to fence", fence the slice that carries its own evidence (M257x iter-163)

`anchor_construct_guard` closed a question in its own docstring for eighteen iters: catching an
anchor that lands on the *wrong ordinary code* "requires deciding what a sentence claims, which is
the line this whole fence family does not cross." True — and it stopped the class dead, while
`corpus/services/backend.md` carried a citation **5 lines off its subject** and stayed **green**,
because the line it landed on happened to be ordinary code.

**The move is not to cross the line. It is to find the sub-population where the corpus already
supplied the evidence.** Documentation usually *quotes* the thing it is describing, and when a
backticked literal sits beside a citation the question stops being *what does this sentence claim*
and becomes *is this string at that line* — a lookup, with a proposed repair attached (the line the
literal is actually on).

Three rules travel with this shape:

1. **State the reach as a fraction of the resolved population, every run.** iter-163's slice is
   **137 adjudicable pairs of 442 resolved single-citation lines**; the other 249 carry no quoted
   literal and are out of reach *by construction*. A partial fence that does not publish its
   denominator will be read as a complete one — `§5`'s standing rule about scoped reds and scoped
   greens.
2. **The PAIRING is the instrument.** iter-163 went **346 → 24 → 16 → 0** without ever changing
   *what* is compared; every step changed *which two things get compared* (one citation per line ·
   the enclosing block, not a ±window · the corpus's own `` `:97` ``/`` `:78-82` `` attribution · no
   pairing across a full stop). A first draft that pairs every literal with every citation reports a
   cross-product and reads exactly like a finding count.
3. **Never tune a clause until a known instance fires** (Trap A). Where a helper genuinely
   under-reaches, take the **declared exemption naming the helper's defect** and route the fix. Two
   of iter-163's nine exemptions exist for exactly that reason and say so at the site.

**And grade every survivor at source.** iter-163's fourth repair was not an anchor defect at all: the
prose named *two* constructs and cited *three* lines, so the missing name — not the anchor — was
wrong. A mechanical "bump the offset" repair would have looked correct, destroyed a true anchor, and
left the real defect in place.

### A fence publishes its FIRE side and hides its ACCEPT side (M257x iter-166)

Every reach number this family prints answers *what did the guard catch*. **Nothing answered the
other half — what did it agree to IGNORE, and does that agreement still apply?** Four guards carried
waiver files; three never named a honoured waiver anywhere, and the fourth printed a **count**
(`22 acknowledged site(s) skipped`), which cannot distinguish a waiver that fires every run from one
dead for months. One waiver file's own README had promised *"every one is reported on each run"* for
as long as it existed, and nothing implemented it.

**The accept side is as mechanical as the fire side — but only the guard may measure it.** iter-165
proved the negative first: re-implementing each guard's matching from outside produced 11 confident
findings and all 11 were the auditor's own normalisation bug. So the rule is not "audit the waivers",
it is:

> **Report the accept side from the guard's OWN matcher.** Have the suppression predicate return the
> KEY it matched on, and feed the report from that return value. The suppression decision and the
> report are then the same decision read twice, and cannot drift.

**And dormancy has three preconditions, not one.** A dormant waiver is evidence it may be dead only
when (a) the guard ran, (b) it graded ≥ 1 candidate, and — the one that is easy to miss —
(c) **the run's subject is the population the waiver was written against.**

(b) is the familiar `§9` zero-census rule: a diff-scoped guard outside a repair grades nothing, so its
waivers are dormant for a reason that has nothing to do with the waivers. **(c) is the trap, because
it survives a large, healthy-looking denominator.** iter-166 measured the same six `repair_reach`
waivers at **0 of 6 honoured over 152 candidates graded** against one ledger and **6 of 6** against
another. Nothing about the waivers changed: their keys are `path:line` coordinates into ONE named
ledger's anchors, while the other three files key on `path` + a quoted form and are subject-
independent. A report that printed the first number bare would have bought six confident deletions.

So a coordinate-keyed waiver set must **name its subject in the report and may never print bare
dormancy**, and any published dormancy figure must state the subject it was taken against — the
`§9` denominator rule, applied to the accept side.

### A test DOUBLE and a test STAGER are both registries, and both rot the same way (M257x iter-166)

Widening one guard function by one optional parameter turned two harnesses red, in the two directions
that matter:

* **A pass-through mutant re-declares its subject's signature.** `test_value_change_guard`'s
  no-suppression double wraps `find_survivors`; the new parameter made it raise `TypeError`, and the
  battery then reported **ERROR instead of a mutant verdict** — it had stopped measuring the guard
  rather than measuring it wrongly, which is harder to notice.
* **A mutation battery that stages its dependencies from a HAND-LIST is iter-162's rule one layer
  down.** The list named four files; the guard grew a fifth; the staged suite died on ImportError and
  the battery reported its **baseline RED**, making all five mutant verdicts uninterpretable *while
  still looking like real kills*. Derive the stage set from the guard's and suite's own imports —
  a stager that can silently omit a dependency is a battery that has stopped measuring anything.

The corollary is the cheerful one: the **derivation-registry completeness fence caught the net-new
module before any new test ran**, which is what an enumerated registry is for. When a change turns
three fences red, grade each one — *working as designed* and *defect* look identical in a summary line.

### A frozen fixture and a live derivation are two clocks (M257x iter-167)

`claim_twin_guard` re-derives its claim ledger from the milestone's blocker-ledgers **on every run**.
The iter-48 answer-key fixture is **pinned** at rosetta `cabc3b1`. Its green-twin assertion read *"no
refuted claim fires here"* — but the capture could only ever support *"no refuted claim **known at
capture** fires here"*. iter-49 re-adjudicated the same corpus region on a **different** form, the two
propositions came apart, and the test was RED at HEAD while three iters shipped over it.

**When a fence and a fixture disagree there are THREE subjects, not two:** the fence, the corpus, and
**the assertion joining them**. The reflex is to suspect the first two. Here the fence fired on a
sentence that genuinely is a refuted claim, the live corpus was clean, and the *assertion* had changed
meaning because one of its inputs kept moving. Name all three before repairing — the one-minute repair
(edit the fixture until it is quiet) would have spent a perishable answer key that cannot be re-taken.

> **An assertion joining a frozen artifact to a live derivation must be scoped to the frozen
> artifact's own denominator, and must say so at the site.**

Two clauses make that scoping honest rather than a silencer:

1. **Assert the residual.** Every hit the scope EXCLUDES must be shown to come from outside the
   capture (here: from a later iter). Without it the scope silently absorbs a real in-capture miss
   that arrives by a different path.
2. **Prove the narrowed predicate still fires** — a permanent control applying the *same scoped
   predicate* to the RED fixture, plus a mutation run at authoring time. A narrowing that has not been
   shown to still fire is not a repair. This is iter-158's rule (a proposed narrowing would have graded
   14 of 14 broken checks green) promoted to a standing requirement on any scope change to a fence.

**And close such a class by RUNNING the family, never by counting repairs.** Two fixes plus an
argument is not the same object as `guard_family.py` at HEAD reporting `17 GREEN · 0 RED · 0
could-not-check · 7 not-run` with every not-run named.

### The stage list of a mutation battery is a REGISTRY (M257x iter-168)

A mutation battery copies a guard and its suite into a scratch tree, mutates one line, and grades the
verdict. **The copy list is a registry**, and iter-162's rule holds for it: *fence its completeness,
never its contents.* Six batteries in this suite hand-listed it, and their own comments record five
occurrences of one class:

| when | dependency added | what the battery reported |
|---|---|---|
| harden pass 1 | `platform_topology.py` | baseline RED, **no attributable test** |
| iter-111 | `fence_provenance.py` | **RED BASELINE** |
| iter-121 | `corpus_citation_guard.py` | **RED BASELINE**, unseen for four iters |
| iter-166 | `waiver_ledger.py` | **RED BASELINE**, 5 mutant verdicts uninterpretable |
| iter-168 | `tests/frozen_capture.py` | **RED BASELINE** |

**The failure mode is what makes it expensive.** A staged tree missing a dependency dies on
`ModuleNotFoundError`, so the battery reports its BASELINE RED and every mutant verdict beside it is
uninterpretable *while still reading as a set of real kills*. It does not look like *"you forgot a
file"*; it looks like the fence is broken — so it is triaged as a fence defect, and twice on this
record it sat unnoticed for four iters.

**iter-111 routed it and wrote the fix down verbatim** — *"a battery that stages a SUBSET carries a
dependency contract, and nothing derives it"* — and the route stayed open while the next three
occurrences were each closed by appending one filename. That is this document's founding sentence
turned on itself: *a recurring class with no written procedure is a class that will recur.* Deriving
the set from the seeds' own imports is the procedure. **Disclose the residual at the same time:** an
import-following derivation cannot see a DATA dependency (a waiver JSON, a checked-in baseline), so
those stay named explicitly and the limit is written where the next reader meets it.

### Measure a hazard's size, or "the same problem exists elsewhere" is only a mood (M257x iter-168)

iter-167 repaired one frozen-fixture/live-derivation collision and argued its two siblings had the
same coupling. The argument was right and it was not yet a finding. Measured against the live ledger's
264 derived claims, the share adjudicated AFTER each capture is **86.4 % / 81.8 % / 75.0 %** — and
**the capture that actually collided had the SMALLEST post-capture surface.** That inverts the
intuition a triage would run on ("the one that broke is the worst one") and it converts *green by
luck* from a worry into a number. `§9`'s denominator rule is usually applied to findings; it applies
to hazards the same way.

**And when a census member is structurally identical but semantically different, do not apply the
pattern to it.** `repair_postcondition`'s green-twin assertion has the same shape, but its site set
feeds a RATCHET — filtering post-capture sites out of a ratchet's input changes what the ratchet
counts and could mask a real induced regression, the one thing it exists to catch. Measure it,
disclose the exposure, route it. iter-158's rule turned on your own repair: **a pattern that fits
three members is a hypothesis about the fourth, not a plan for it.**

### Closing a class means fencing its POPULATION, not its last member (M257x iter-169)

`FIX-M257x-iter111-staged-battery-dependency-is-underived` was opened at iter-111 with its own fix stated
verbatim, then recurred at iter-121, iter-166 and iter-168 — **five occurrences, each closed by appending
one filename to the list the route was about.** iter-168 derived five of the six known stage sets and routed
the sixth. iter-169 closed the sixth and found that **that still would not have closed the class**: nothing
prevented member seven from arriving hand-listed, and the record says member seven is not hypothetical.

**The deliverable is a fence over the population, and the population is derived by property.** A mutation
battery is *a test module binding a module-level MUTANT registry* — not `*mutation_battery*.py`, which is
the glob rule 73 forbids. Under that predicate the population is **seven**, and the seventh is **exempt with
a proof rather than a sentence**: `test_m220_mutation_battery` mutates one subject into a gitignored sibling
beside the real tree and copies no file set, which the fence re-establishes on every run. An exemption that
is only asserted is a hand-list with better manners.

Three things fell out of doing it, and each is the general lesson:

* **The sixth occurrence was already live and symptomless.** Deriving m255's stage set returned one file the
  hand-list did not carry — `fence_provenance.py` — imported inside `main()`, which a suite run never
  reaches. **A latent registry defect is not an averted one**; it is the same defect waiting for a caller.
* **Over-approximate on purpose, and assert that you do.** The derivation follows function-scope imports
  exactly like module-scope ones. Under-staging reports a BASELINE RED with no attributable test;
  over-staging costs one `shutil.copy2`. The asymmetry is the design, so a permanent test states it —
  otherwise someone "fixes" the imprecision and re-opens the class.
* **The widening is a resolution change, not a bigger haystack.** Cross-section staging works by resolving
  each import against **the importing file's own directory** — what the interpreter does — then falling back
  to the root-relative conventions. Proved a no-op on the five already-migrated batteries by measuring their
  derived sets before and after, per rule 73's widen-only-when-the-verdict-is-unchanged corollary. And what
  a derivation must **refuse**: a repo file that would shadow a stdlib module inside the staged tree is an
  error, never a copy — that is a staged-only divergence, the exact class the helper exists to end.

### A DERIVED number is censusable; an OBSERVED one is not — split the class before you scope the fence (M257x iter-173)

**The distinction is the whole finding, and it is what made the census cost seconds rather than the
50 minutes iter-172's two-runner sweep cost.** A published count is one of two things, and only one of
them is reachable:

| | can it be re-checked? | how |
|---|---|---|
| **observed** — read off a runner (`1 failed · 1229 passed`) | **no**, not without re-running at a ref that may no longer exist | re-run: expensive, and `§5` rule 51's timing leg is unusable on this host |
| **derived** — a function of other published numbers on the same page (a table total, a delta, a percentage) | **yes**, with no runner, no host and no clone | arithmetic, on the page itself |

iter-172 routed *"every pytest count published before this iter is a `passed` count, therefore an
undercount"* — a class spanning both halves. Taken whole it is unaffordable. **Split it, and the derived
half is a census that runs in under a second**; the observed half stays routed, and the iter says so
rather than implying the green covered it (`§5` rule 60).

**What the census found, and it is the rule about denominators failing its own rule.** The harden ledger
summarised its own five-section table as *"one section of five, **1,230 of 2,989** tests"*. The table
says `2,978 passed · 22 failed · 11 skipped`. **`2,989 = 2,978 + 11`** — the **22 failures were dropped**,
so the denominator had silently changed unit from *executed* to *passed-and-skipped*, and the executed
population is **3,011**. The next entry carried the hole forward as `2,989 + 51 = 3,040`, and from there
it reached this document as the **evidence for `§5` rule 68 — the rule that "the whole suite" must name
its denominator.** Corrected to **1,281 of 3,062**.

Three sub-rules, each earned:

1. **A percentage can survive an error its operands do not — so never audit the ratio in place of the
   operands.** `1,280/3,040` and `1,281/3,062` are both **42 %**. The figure was published, quoted and
   re-quoted across 28 iters with its headline conclusion intact and its arithmetic false. A ratio that
   still looks right is the most durable place for a wrong count to live.
2. **A fence that cannot reach the claim can still be worth building, if it proves the claim's OPERANDS.**
   `N of M` prose is not machine-derivable — `M` names no source, and attributing it to a nearby table is
   the inference this milestone has nine iters of evidence against. So `stack-core/derived_count_guard.py`
   fences the *table totals, deltas and percent-triples* instead: the repair is still done by hand, but it
   now rests on a fenced ground truth rather than on a second reading. **State that reach with its
   denominator, and state what it does not reach, inside the tool's own report** — this one prints its
   NOT-REACHED line on every run, green included.
3. **Two independent runs agreeing on a total is stronger evidence than either run.** The four
   never-run sections come to **1,781 executed** in the ledger's table (`1,749 passed + 21 failed +
   11 skipped`) *and* in iter-145's separate re-run of them (`1,758 + 21 + 2`) — different pass/skip
   splits, identical executed total. That agreement is what let the corrected denominator be published as
   a number rather than an estimate.

**And the ownership note, because it shaped the repair:** `hardening-ledger.md` is owned by
`/developer-kit:harden-mstone-iters`, so a tik must **route** a correction into it rather than write it.
Two of the five sites were routed with the derivation pre-computed; three were repaired in place.
A correction that respects file ownership arrives later and intact; one that does not arrives as a
merge conflict.

### A capability probe that fails OPEN disarms the check it guards (M257x iter-174)

```python
stdlib = getattr(sys, "stdlib_module_names", frozenset())   # ← shipped for six iters
```

`sys.stdlib_module_names` landed in **Python 3.10**. The dev box has two interpreters, and
`/usr/bin/python3` **3.9.6** — *the only one with pytest, i.e. the one that runs everything* — is not one
of them. So that expression was an **empty set**, the membership test under it was never true, and the
refusal it gated **could not fire at any input**. The check was not weak; it was **off**.

**The shape, and it generalises past Python:** `getattr(x, "capability", <empty default>)` silently
converts *"this environment cannot tell me"* into *"the answer is nothing."* The two directions are not
symmetric:

| default | when the probe fails | consequence |
|---|---|---|
| empty / permissive | the check passes everything | **silent** — nothing to notice, forever |
| broad / restrictive | the check refuses something legitimate | **loud** — someone fixes it that afternoon |

**So: derive it, or refuse. Never default to empty.** Where the capability is absent, compute the answer
another way (here: `sys.builtin_module_names` plus the `sysconfig` stdlib directory — 232 names on 3.9.6,
297 on 3.14.6). Where it cannot be computed at all, **raise**. *A check that cannot check must not report
OK* — the same direction as `§9` iter-149 (a census returning zero must prove its instrument) and M236's
green-gate, which parsed a UTC timestamp as local time and therefore **aged a stale verdict as fresh
everywhere west of UTC**: failing open, silently, for half the world.

**Two corollaries, both earned here:**

1. **Assert the CAPABILITY, not only the behaviour.** The only witness was the guard's own behavioural
   test, and its message — `RuntimeError not raised` — named *the refusal*, when the defect was *the set
   the refusal consults*, one function away. A whole iter characterised it as "fails on the old runner"
   and stopped there, because nothing asserted the set. The repair's controls now assert the set directly
   (non-empty · contains what the check must catch · contains nothing it must not · equals the native
   value where one exists), and restoring the shipped form kills **eight** of them instead of one.
2. **Measure the hazard's population before calling it a class** (`§8`, iter-168). All 13
   `getattr(x, "attr", <empty>)` sites in the tooling were classified: **1** is a capability probe
   deciding a verdict; **12** are attribute lookups where the default means *"not set"*, which is the
   true answer. Not systemic — and saying so with the denominator is the difference between a measurement
   and a mood.

### A pre-registered escalation names a SYMPTOM; it cannot name a cause (M257x iter-174)

The iter pre-registered: *"if arming the refusal turns other batteries RED, that is a finding to route,
not to suppress by weakening the derivation."* The whole-suite run then returned **4 failed** in a
mutation battery — the exact shape the clause describes.

**None of the four had anything to do with the change.** All were one cause: a fence the *previous* iter
had registered in the postcondition baseline was missing from that battery's fence-seed list. Applied on
its face, the clause would have had the iter weaken a correct derivation to silence an unrelated
regression.

Pre-registration is worth keeping — it is what stops an inconvenient result being re-interpreted after the
fact. But it binds the **response to a diagnosis**, never the diagnosis itself. **Read the failure before
applying the rule you wrote for it.** iter-169 earned *a route predicts a cause; it does not certify one*;
this is the same rule applied to one's own pre-registration, which is the harder case because the author
trusts it more.

**And the registry count is now five.** iter-173 enumerated four registries a new fence must join and
found the fourth only by running the whole suite; this iter found the fifth the same way, one iter later,
on iter-173's own commit — while iter-173's post-fix **scoped** re-run was green and structurally could
not see it. `repair_postcondition --accept` writes one of the five. **Nothing enumerates them**
(`FIX-M257x-iter174-accept-registers-one-registry-of-two`).

### Two derivations of ONE population must be COMPARED, or the weaker one is a silent census (M257x iter-175)

iter-174 left a sentence that is a class, not an item: *"five registries are now known; nothing enumerates
them."* The five was a **remembered list** — four at iter-173, five at iter-174, each reached by grepping
for a sibling's name. That is §2's hand-maintained tuple one level up: the registry **of registries**.

**Instrument first, and grade it at the claim's grain (§9, iter-159).** A file naming ≥2 fence modules
anywhere returns **39** sites and measures *mentions*. A **collection literal** (py `List`/`Tuple`/`Set`/
`Dict`-keys/call-args, or a JSON array/object) holding ≥2 fence-module names returns **5** — and the claim
was always about *a set that must track the tree*. The sharper instrument landed on the five the route was
groping for, plus the one nobody had named.

**What it found, one row above the routed item:** `guard_family.census()` derived the fence family from a
FILENAME SPELLING (`glob("*_guard.py")` + a hand-maintained `EXTRA_CENSUS_MEMBERS` escape tuple), while
`repair_postcondition.discover_fences()` — repaired at iter-157 for exactly this — derived **the same
population** from the `FENCE_KIND` DECLARATION. **26 modules declare; the family censused 25**, and the
sets disagreed in *both* directions. `predicate_enumerator.py:142` declares itself a fence, and the runner
whose docstring promises to *"run the WHOLE guard family, and name every member"* had **never run it and
never named it** — not NOT-RUN in the verdict, *absent from it*. Its own founding sentence is the
indictment: a guard that was not run reads exactly like a guard that passed.

> **The rule:** when two derivations of one population exist, **the fence is the comparison between them**,
> not either one's own completeness check. iter-157 shipped a completeness fence — for one module's
> registry. iter-169's *closing a class means fencing its POPULATION, not its last member*, one turn on:
> **the population here is the set of DERIVATIONS, and two of them had never been put side by side.**

### Repair a derivation by UNION, never by substitution (M257x iter-175)

The apparent repair — swap the glob for the declaration, the symmetry with iter-157 — is a **weakening**,
and it looks like a tidy-up. A `*_guard.py` that declares nothing would then leave the family in silence,
and that file is precisely the one worth catching.

**A member needs only ONE of the two properties to be counted.** `spelled ∪ declared ∪ extra` closes the
spelling gap and the declaration gap at once and can lose neither; it is strictly stronger in both
directions, which is the only form of this repair that iter-158's rule permits (*a narrowing that grades a
broken check green is a defect, not a fix*). The escape tuple survives — demoted to **additive only**: it
may add a member, never substitute for a property.

**A member the census reaches and the runner does not run needs a DISPOSITION, printed.** An exclusion
table with a per-member reason, held to the same two directions as the invocation map — an exclusion that
subtracts nothing is stale, an exclusion that is also invoked is ambiguous — and the reasons are **printed
on every run**, because a member omitted from the census is that founding sentence with the evidence
removed.

**And the fixture was carrying the same shape.** Four synthetic guard dirs staged a family with no runner
in it. One of them asserted `exit 2` and, for the window between the exclusion table landing and the
fixture being made faithful, got its 2 from a stale-exclusion complaint while the orphan it was written for
was never reached. **An exit code is not a diagnosis** — assert the sentence.

### A class is not closed by a repair; it is closed by an ENUMERATION THAT KEEPS RUNNING (M257x iter-176)

iter-175 repaired the largest member of the registry class and **measured** its population with a scratch
script it deliberately did not check in. The measurement was correct and the population would have gone
straight back to being *remembered* the moment the script was deleted — which is what iter-173 and
iter-174 had each already done, one registry at a time.

> **A reading SAMPLES; a fence CENSUSES** — and an instrument that ran once and was deleted is a reading
> wearing a fence's clothes.

**The tell is in the discovery record, and it is mechanical.** Four consecutive registries were found by
the most expensive instrument the milestone owns:

| iter | registry found | found by |
|---|---|---|
| 173 | three siblings | grep for a sibling's name |
| 173 | the ratchet baseline (4th) | whole-suite run, after the fact |
| 174 | the battery's fence-seed list (5th) | whole-suite run, one iter later |
| 175 | the derivation registry again, and a 6th by hand | whole-suite run + a hand check |

**When the most expensive instrument you own is the one making the discoveries, the cheap instrument does
not exist yet.** Shipped as a checked-in fence over the population — every collection literal holding ≥2
fence-module names must be classified `REGISTRY:<what keeps it in sync>` or `DECLINE:<class>: <reason>`,
unclassified is RED — discovery of the seventh moves from a 36-minute suite run four iters later to a
sub-second static check.

Two corollaries the iter paid for directly:

1. **The predicate must include the SHAPE the motivating case actually has.** The seed list that started
   the thread is not a literal at all — it is `helper(ROOT, "markdown_structure_guard.py", …)`, a **call**.
   A literals-only predicate is tidier and one sentence shorter and **structurally blind to the registry
   it was commissioned for.** Elegance that cannot see its own commissioning case is not elegance.
2. **Measure an exclusion before shipping it.** The first draft skipped `fixtures/` on reflex; measured,
   the site count was **identical with and without it**. An exclusion that changes nothing buys nothing
   and silently forecloses the case it excludes — and an unmeasured narrowing inside the fence written
   against unmeasured narrowings is the joke telling itself.

---

## 9. Cadence

### Measurement preconditions — the host facts an iter must not rediscover (M257x iter-145)

Two of these have now each cost an iter real time, and both were already written down **in the harden
ledger**, which is not a document the iter loop reads. A precondition recorded only where the next reader
will not look is a precondition that gets rediscovered.

* **The suite's interpreter is `/usr/bin/python3` — 3.9.6, and it is the only one on this host with
  pytest.** `python3` on this shell is homebrew 3.14 and has none; `pip3.12`/`pip3.14` do not have it
  either, and there is no venv, no `uv`, no `pyenv`. Invoking the suite as `python3 -m pytest` fails with
  `No module named pytest`, which reads as a broken environment rather than a wrong interpreter. iter-145
  lost two full runs to it.
* **`timeout(1)` does not exist on macOS.** A run wrapped in it dies with `command not found` per section
  and produces an empty log that looks like a suite that collected nothing.
* **Suite WALL-TIME is not a usable measurement on this host** (rule 51's timing leg) — the box shows CPU
  contention from unrelated processes. **Counts are.** And do not edit the tree while a suite runs: three
  runs have been discarded as confounded for exactly that.
* **"The whole suite" must name its denominator** — see `§5` rule 68. There are five
  `rosetta-extensions` sections; running one of them is a measurement of one of them.
* **The working tree contains EXHAUST, and some of it is byte-identical to defects already repaired**
  (M257x iter-149). The M220 mutation battery stages mutated copies of its subjects **beside** the real
  ones (`.m220-mutant-*`, gitignored) because those scripts resolve their siblings from `$HERE` and cannot
  run from `/tmp`; `tearDown` removes them, an interrupted run's survive. iter-149 found **33 of them,
  oldest four days**, each a verbatim copy of the emitter line iter-146 had repaired — **66 % of that
  iter's raw census signal**. Two consequences, both general: **(a)** establish what in the tree is a
  measurement subject and what is exhaust *before* reading any repo-wide count — iter-148's *state the
  substrate before booking a failure*, in its stronger form; **(b)** a harness that stages copies in the
  source tree owes a **self-healing sweep**, age-gated so it cannot delete a live sibling run's staging.

### A census that returns ZERO must prove its instrument (M257x iter-149)

A fence built on a clean reading is unfalsifiable from the tree alone: an arm that matches nothing is
indistinguishable from an arm over content that contains nothing. iter-149 widened the retired-service
emitter fence from one hard-coded port to a 12-name × 3-arm grid and measured **0** defects, so neither
the container arm nor the address arm had a single real occurrence to demonstrate it worked.

**Two independent proofs, both required before a zero is published:**

1. **A real answer key.** Run the fence against content from *before* the repair it generalises — here the
   pre-iter-146 `dev-stack` at `1a44b97^`, which returns the original defect line, against a current file
   that returns nothing. This is `§5` rule 21's perishability clause: the fixture is spent once the
   repair is old enough that nobody remembers what it looked like.
2. **Per arm, on synthetic content written to trip it.** Every arm with no real occurrence gets one.
   An arm proven only by a green tree is decorative, and the fence reads wider than it is.

And the sibling rule for the subject set: **bind it to a set that is already fenced, never re-declare it.**
iter-149's fence imports `claim_census_guard.ARCHIVED_SERVICE_NAMES` — the row set of the migration map,
itself fenced against the platform's `repos.yml` in both directions — so a service entering or leaving the
map reaches the fence with nobody re-typing anything. A hand-maintained thirteenth list is `§2`'s tuple
defect wearing a different hat, and this milestone has now found it in a bring-up script (iter-01), a
probe registry (iter-148), a skills enumeration (iter-129) and a guard's own data (iter-149).

**And when a value genuinely cannot be derived, its COMPLETENESS usually still can** (M257x iter-150).
`blocking_state_guard` partitions the iteration protocol's exit enum into blocking and non-blocking
tuples, under a comment claiming the split was *"derived from the iteration protocol's own Phase-5
grading"*. Nothing derived it — and the obvious repair is wrong, because which side a condition falls on
is a **judgement** about what it means, which no parse can make. What is decidable is the field
*universe*: the gradings are the one place the enum is written out field by field, and the guard already
parsed them. Subtracting both tuples from the fields actually graded now reports anything neither
classifies. **Split a claim into its decidable and undecidable halves** — usually cheaper than either
deriving everything or fencing nothing. The stake is not theoretical: `budget-exhausted` entered that
enum on 2026-08-06 and was hand-added to the tuple; had it not been, a whole exit class would have read
as non-blocking **by omission**, which in the output is indistinguishable from safe.

**Corollary, from that last one: a list whose comment claims it is DERIVED, and which is not, will drift.**
`claim_census_guard.REXT_SECTION_NAMES` said *"derived from the monorepo's own layout"* and was missing
`stack-secrets` — 10 declared against 11 on disk, so every claim naming that section resolved to no known
artifact and left the census silently. When the module genuinely cannot derive (no repo-root notion, and
imported from copies whose layout differs), the repair is **not** to make it derive. It is to fix the
list, say in the comment that it is declared **and why**, and move the derivation into a test that
compares the two in both directions — **fence the property the comment asserted.**

### And grade the instrument at the GRAIN OF ITS CLAIM (M257x iter-159)

The companion failure to the one above, and it points the other way: iter-149's rule stops you
publishing a zero from an instrument that cannot fire. This one stops you **discarding an instrument
that works, because the test of it was too blunt to see it succeed.**

iter-159 built a predicate for the spelling-pin class (`§5` rules 70/71) and proved it against a
**labeled set** — the seven confirmed instances of the class, each with the commit that repaired it,
ground truth this milestone had produced as a by-product and never used. Graded at **file** level —
*did the file's candidate count fall at the repair commit?* — it scored **1 of 4** and printed
*"measures the file, not the pin."* That is a refutation, and it was **false**: the repairs each removed
**one** pinned assertion from files carrying 9, 51 and 56 legitimate assertions of the same shape, so
the count physically cannot move. Re-graded at the grain the claim is about — ***was the assertion the
repair DELETED among the lines the predicate flagged?*** — the same predicate on the same data scored
**4 of 4**.

**The rule:** a claim of the form *"this instrument finds X"* must be tested by asking whether it found
**X**, not whether some aggregate that contains X moved. Aggregates are where a true positive goes to
hide — one correct hit inside 56 is invisible to a count and unmistakable to an intersection.

**Two corollaries worth the same weight:**

- **A labeled set is worth building before the instrument.** Any milestone that has repaired instances
  of a class already owns one — the repairs are the labels. It converts *"is this predicate sharp?"*
  from an argument into two numbers (**recall**: did it fire on the pre-repair form; **discrimination**:
  did it name the repaired line), and the second is the one that earns the instrument its keep.
- **When the labeled set says the predicate MISSES, check whether the misses are the same class before
  improving the predicate.** Here they were not: two of the seven pinned a hand-written **value** rather
  than searching a **haystack**, so no haystack clause could ever reach them. "Improve recall" would have
  been effort against a target that does not exist — the real finding was that **the class is two
  classes**. Report the recall on the pre-registered denominator and let the taxonomy be the deliverable;
  re-labelling the misses "blind" after seeing them lifts the number to 4/4 and deletes the finding.

### Iter-type refinement — the 3-no-prog tok-trigger reads UNMEASURED as UNMEASURED (M257x iter-108)

**This protocol's primary metric (`N`, the graded read) is expensive and is deliberately not measured every
iter.** A strategy under this protocol may legitimately sequence several iters that build instrumentation
before the next reading — `TOK-06` does exactly that, putting the read **last** so it measures a pool whose
inflows are already fenced rather than re-measuring the inflow.

That collides with the generic 3-consecutive-no-progress tok-trigger, which fires when *"the metric did not
move in any of those 3 tiks (zero or net-negative delta)."*

**The refinement, and it narrows nothing:** a **delta requires two measurements.** An iter that took **no
reading** has an **UNMEASURED** metric, not an unmoved one, and does not count toward the streak. Grading
"not measured" as "did not move" asserts something nobody measured — §8's *grade the cannot-tell* applied to
the trigger itself.

**The floor is not suppressed.** Three consecutive tiks that DID measure and did not move still fire the
trigger. What this rule excludes is only the case where the metric was never read.

**Two guard-rails, both mandatory:**

1. **The iter must SAY it took no reading**, in its own close, in the words *"no `N` movement is claimed"*.
   An iter that quietly omits the measurement and an iter that measured zero must not look alike.
2. **The strategy must have declared the sequence in advance** (a `TOK-NN` entry naming the step order). A
   post-hoc claim that the metric "wasn't being measured anyway" is the flattering reading, and this
   milestone refuses those on principle.

Worked case: at iter-108, iters 105/106/107 were three consecutive tiks with no `N` movement — and the
trigger correctly did **not** fire, because `TOK-06` had declared them steps 0–2 of five and each said so at
its close. Firing would have revised a strategy **3 of 5 steps in, before either of its metric-moving steps
had run** — revising it from evidence that step 4 exists to produce. It would also mean **no declared
multi-step strategy longer than three non-metric steps could ever complete**, since the tok terminates the
call.

### Grade a refuted strategy LEG BY LEG, against what each leg measured (M257x iter-110)

A strategy that turns out to rest on a wrong premise is not thereby wrong in all its parts, and the reflex
to revert it wholesale destroys work that measurement says was earning.

**The worked case.** `TOK-06` was authored on iter-103's decomposition — *"inflow is comparable to
outflow"* — and put two fences ahead of the next repair: one on **clone-advance drift** (61 % of the
residual) and one on the **repair's own induction** (21 %). iter-109 froze all 14 platform clones at one
sha, verified identical at open and close, and read the residual again. **Drift was still ~33 %** over a
subject in which literally nothing moved. The premise is refuted: the 61 % had measured the *composition of
what a reading detects*, not the *rate at which defects arrive*. The pool is standing, not flowing.

**And the induction leg worked anyway.** Repair-induced anchors went **21 % → 5.6 %** of the residual — the
lowest in the series, against a rate that had held steady for six prior cycles at a far smaller repair
size. The fences that produced that are in production and earning.

**The rule:**

> When a strategy's premise is refuted, **do not grade the strategy. Grade each leg against the number that
> leg moved.** A leg with its own measurement survives on that measurement. A leg whose only justification
> was the refuted premise is **de-ranked, not cancelled** — the work may still be worth doing, just not
> first.

Applied at iter-110: the induction fences were **kept**; the drift fence (`FIX-M257x-iter107-...`) was
**de-ranked and left open**, because drift that is *standing* rather than *arriving* still wants a fence —
it is simply no longer the lever it was ranked as.

**Why this is not special pleading.** The test is whether the leg has a measurement of its **own**,
recorded **before** the premise fell. Induction did (band #10, pre-registered in iter-109's sealed rule).
A leg whose defence is assembled after the refutation is the flattering reading, and this protocol refuses
those on principle (§5).

### A reach metric is settled by its DENOMINATOR's provenance (M257x iter-110)

The same sentence as §5 rule 50 (*a guard verdict is settled by the tree its configuration lives in*), one
layer up, and it cost this milestone a full repair cycle to learn twice.

iter-108 was machine-graded **46/46 = 100 % of the upheld union** by `repair_reach_guard`, and **the grade
was correct**. Its anchor list was derived — never hand-assembled — from `iter-103/raw/`, i.e. from *what
the previous reading detected*. **A predicate's site list and a reading's detection list are not the same
set.** Per-pass detection recall on that instrument has run **33–83 %**, so the repair closed every site
the reading saw and left the same falsehood standing wherever the reading had not looked. Two survivors
were measured directly at the next reading, and **one had become a self-contradiction, because the repair
fixed one side of a pair.**

> **A reach percentage whose denominator came from a prior detection is a check reporting a state it did
> not measure.** State the denominator's provenance with the number, and make a run that cannot state it
> unable to print a percentage at all.

**The multiplier is the cheap tell, and it should be reported per predicate:**

| repair pass | site list derived from | booked | sites | multiplier |
|---|---|---|---|---|
| M257x iter-96 | the corpus, per predicate | 13 | **51** | **3.92×** (it also counted the **38** an anchor-wise repair would have left) |
| M257x iter-98 | the corpus, per predicate + paraphrase | 20 | 37 | 1.85× |
| M257x iter-102 | 76 anchor assignments | 52 | 98 | 1.88× |
| M257x iter-108 | **a prior reading's `raw/` ledger** | 46 | 46 reached | **no expansion figure reported at all** |

**A multiplier near 1.0× is evidence the ENUMERATION is not working — not that the predicate is rare.** And
note the last row precisely: iter-108 did not report a *low* multiplier, it reported **none**, because
there was no expansion step to produce one. An absent step is easier to miss in a close than a bad number,
which is why the multiplier is worth printing even when it is boring.

### But "rare" still has to be EARNED — against a ceiling (M257x iter-113)

The rule above is only half a procedure, and iter-112 walked straight into the missing half. It reported
**12 of 24 predicates at `NO-EXPANSION`**, obeyed the rule, refused to bank its 7.28× headline — and could
go no further, because **"the form is too narrow" and "the class really is that small" produce the
identical number.** A rule that tells you to distrust a number without telling you how to settle it just
relocates the guess.

> **Give every predicate a SECOND, BROADER form tier — the SUBJECT — and the ceiling falls out.**
> `ceiling` = every site that so much as mentions the topic. `headroom` = subject sites − predicate sites,
> **named by `file:line`, never merely counted**. Zero headroom settles `NO-EXPANSION` as a small class;
> non-zero headroom is **RED until every candidate is either folded in or excluded with a reason**.

Four things this bought, all measured on the real ledger:

1. **It fixes the opposite failure with the same lever.** iter-112's four *vocabulary* forms — `Cosmo
   Router` ×37 for a claim about VPC public subnets — were never bad forms; they were forms **on the wrong
   tier**. Moved down, those four predicates enumerate **5 sites between them instead of 162**.
2. **The multiplier went 7.28× → 2.45×, and DOWN was the honest direction.** Report the number that
   indicts the previous one; a figure that mostly counts how often the corpus writes a component's name is
   not a reach measurement.
3. **The real twins surfaced where the flat forms had hidden them** — `P21` 6 → 22 sites, and a
   same-fact-different-pin pair (`:1594-1597` vs `:1594-1600`) whose one-sided repair would have
   manufactured exactly the self-contradiction `TOK-07` rule 3 forbids.
4. **The paragraph is the unit of publication.** Corpus prose wraps at ~110 columns, so a subject token
   routinely sits a line above its own claim — one publication seen twice. Suppress *within-paragraph*
   subject hits only; the `ai`-fold twin at `external_services.md:554`/`:565` is eleven lines apart and
   must survive. Test both halves together, and keep the coverage invariant at **file** granularity — a
   line-level subset test goes RED on ordinary wrapped prose, the `anchor_offset_guard` false-RED again.

**And split the verdict in the OUTPUT, because only part of it is mechanical.** iter-113 settled all 16 of
its flat predicates, but **1** had zero headroom (`SMALL-CLASS-PROVEN` — nothing was judged) and **15**
rested on **254 candidate sites read and excluded by hand** (`SMALL-CLASS-ADJUDICATED`). The enumeration
guarantees the candidate set is complete and that nothing was left unexamined; it does **not** guarantee
the reasons are right. Two warranties, two tokens. Collapsing them lets a pile of judgement calls wear a
measurement's voice — which is `fence_provenance`'s defect (§8, iter-111) one layer up.

---

Detection is cheap. Run it on a schedule, not on an incident.

- **At every release open, and before any prove-it-live milestone:** run §4's six signals.
- **Whenever a bring-up fails oddly:** check signals 2 and 3 *before* debugging the tooling. Three times the
  answer was "the platform moved."
- **Watch the named next fold.** ~~v9.0 folds `storage` + `messenger`.~~ **It landed on 2026-08-05**
  (`838d907`, merged `0c91421`) and took `customerio-sync` with it, which no plan doc had named. All three
  compose services were **deleted outright** — not moved to a rollback profile — and `storage` + `messenger`
  left `repos.yml`. The two folded subsystems are gated inside `app` by `MESSENGER_ENABLED` /
  `CUSTOMERIO_SYNC_ENABLED`, unset meaning off on a developer machine (`docker-compose.yml:84-92`).
  **The next fold is not yet named by the platform** — watch signal 6's `archived` flag and the PR list
  rather than waiting for a plan doc, since 4 of the last month's 5 structural changes had no doc PR at all.
- ~~**When M810 lands** and the legacy repos leave the clone set, §2's time bomb fires.~~ **The repos left
  the clone set on 2026-08-05 and the bomb did not fire, because §2's derivation had already replaced the
  tuple.** Measured across the move, with `repos_yml_schemas_to_create` run against both refs:

  | | `0dab54d` | `0c91421` |
  |---|---|---|
  | migration pairs | `app:public` | `app:public` |
  | CREATE SCHEMA set | `extensions sentinel public` | `extensions sentinel public` |

  **Identical, and identical *correctly*** — zero human action, on the exact event that was forecast to
  break 13 write targets with 42P01 at once. This is the **third** consecutive platform change the derived
  layer has absorbed unaided (§5 rule 27 records the first two). The hand-maintained tuple would have
  silently skipped both repos on this commit. **Keep the emptied debt list and its shrink-fence** — the
  argument for that in §2 is now paid off twice over.
- **M810 is still open, and it is UNEVEN — do not state it as one milestone-wide event.** `jobsimulation`'s
  ECS service, task definition and ECR repository were **destroyed** by `6092c6d2`; `cms` is still at
  `service_desired_count = 0`. Dropping the legacy `jobsimulation` schema is a further, separate M810 step.
  The fenced map is authoritative for the per-service state.

---

## 10. Reading list

- [`platform-migration-status.md`](../architecture/platform-migration-status.md) — the map
- [`platform_repo.md`](./platform_repo.md) — the Makefile, profiles, and `repos.yml`
- [`verification.md`](./verification.md) — pre-flight rung zero (*tagging is not publishing*)
- [`safety.md`](./safety.md) — the safety contract a re-point must not weaken
- [`idempotency.md`](./idempotency.md) — what happens when a bring-up step runs twice

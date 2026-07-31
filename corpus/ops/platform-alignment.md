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
| 6 | org repo census | `gh repo list anthropos-work --limit 300` | net-new repos; newly archived repos |

At 2026-07-31 these read: 10 repos · `migrations: true` → **`app` alone** · `schema:` → **`app -> public`
alone** · **one** subgraph (`backend.graphqls`) · 14 compose services with cms/jobsimulation/roadrunner still
in the default profile · **93** org repos.

**Signals 2 and 3 are the load-bearing pair** — they are what moved under us all three times.

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
| map ↔ `repos.yml`, both ways | every `repos.yml` service has a map row; no map row invents a repo | `stack-core/platform_alignment_guard.py` (precedent: `corpus_index_guard.py`) |
| static schema fence | every schema a seeder WRITES to is one the migrate step CREATES | `stack-core/tests/test_write_target_schema_fence.py` (M257x iter-06) — reads the legal set from `repos_yml_schemas_to_create`, so it **names no dead schema at all** |
| live schema assert | every schema rext writes exists in `information_schema.schemata` on the migrated stack | bring-up / autoverify (precedent: `dev-stack/tests/test_migrate_dev_live.py:144`) |

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
4. **Scope the construct to its BLOCK, or the fence cries wolf.** iter-06's first cut of the write-target
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

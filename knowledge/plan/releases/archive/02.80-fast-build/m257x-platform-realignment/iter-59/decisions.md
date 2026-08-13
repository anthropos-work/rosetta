# iter-59 — decisions

The five decisions TOK-05 was directed to make explicitly. Each is recorded here in full; TOK-05 in the
milestone-root `decisions.md` summarises them and carries the strategy.

All denominators below were **re-derived at this iter's open** against platform `0dab54d` / app `v1.366.0`
— none are inherited. See `overview.md` § Step 0 for the re-survey table and the one correction it made.

---

## `D-M257x-59-1` — clause 5's residual is scoped by PREDICATE, not by count

**Question directed:** *does clause 5's residual get re-scoped by class rather than by count?*

**Decision: yes.** The repair unit changes from *the claim* to *the predicate*. A unit of work is now
"every site in the tree asserting predicate P", adjudicated against **the platform artifact that defines
P's legal set**, and closed by a fence that makes P underivable when false.

**Why.** Three residuals, three predicates:

| residual | count | the single false predicate underneath it | the artifact that defines its legal set |
|---|---|---|---|
| the drift sites from 2026-08-03 | 81 sites / 21 files | *three services are live-local husks* | `docker-compose.yml` + `repos.yml` |
| the profile class | 17 files / 30 occurrences | *a `graphql` profile exists* | compose `profiles:` declarations (8 legal) |
| the mainline shift | 21 of 22 | *this line number names this construct* | the cited file at the pinned ref |

**119 sites, 3 predicates.** A reading names *instances*; it cannot name the predicate they are instances
of, which is why ten readings at 43–48% recall never converged. A predicate has a legal set; a legal set can
be derived; a derivation does not need to be found by a reader at all.

**What this does NOT change.** Clause 5 is **met only by a reading that returns zero** — the user has ruled
three times. Predicate-scoping changes how the residual is *repaired*, never how it is *graded*. The audit
instrument is untouched: union-of-two blind readings, the blind second reading, and the pre-commit double
reads all stay exactly as they are.

**On `FIX-M257x-iter53-union-set` (46 vs 35) — NOT RESOLVED HERE.** It is a **pending user decision** and
this iter does not touch it. What can be said without deciding it: predicate-scoping **subsumes** the
question rather than answering it. A predicate sweep covers every site sharing the predicate *whether or not
a reading named it*, which is exactly the recall gap the union was invented to paper over. So the union
count moves from being the **scoping** input to being a **validation** input — after the sweep, every member
of the union must be covered, and any member that is not **names a predicate we have not yet found**, which
is a more useful output than either number. Whether the set is 46 or 35 changes the validation, not the
work; the user's decision is still owed and still unblocked.

---

## `D-M257x-59-2` — the fence widening is the next build, as a NEW sibling guard

**Question directed:** *is the fence-widening the next build, and what exactly does it fence?*

**Decision: yes, it is iter-60 — and it is a new `*_guard.py` sibling, NOT a widening of
`platform_alignment_guard.py`.**

**Why a sibling, departing from the reconciliation's recommendation.** Measured at this open, that guard
declares `FENCE_KIND = "standalone"` and its docstring scopes it to *"a map↔platform property, not a set of
claim sites"*. Assertion F derives its clone roots **from `repos_yml_path`** on purpose, so the citation
check and the membership check cannot be run against different references — the guard's own comment names
that as §4 Trap A wearing a second hat. Re-targeting it at the whole corpus would destroy the property that
makes it trustworthy. Its `compose_blocks()` parser (written and tested at iter-57, including the
`context: ${APP_BUILD_CONTEXT:-../app}` alias derivation that teaches it `backend` **is** `app`) is reused
as an **importable primitive**. That is an extension of the code, not of the guard's subject. §8 rule 1's
derived-from-disk registry means the new file self-registers.

**Inputs:** `repos.yml` · `docker-compose.yml` **with its `include:` resolved** · `Makefile`.
The include is load-bearing and was the correction this iter's re-survey made: reading
`docker-compose.yml` alone sees **8** services; resolving `include: [common.yml]` gives **10**, and the two
it adds (`postgresql`, `redis`) declare **no** `profiles:` key, which is the entire mechanism behind G1.

**Assertions — six, each derived, each run in BOTH directions.** A doc-promised value with no artifact
backing is a **false promise**; an artifact value with no doc row is **undiscoverable**. The
`demo_knob_guard.py` precedent already does this both-ways (31 env knobs + 10 CLI flags across 2 entry
points) and is green in the suite.

| # | assertion | derived denominator at `0dab54d` |
|---|---|---|
| **G1** | every profile token in a documented runnable command is in the derived legal set **and selects at least one service beyond the always-on floor** | **10** services · **8** legal profiles (`all backend core customerio-sync frontend messenger storage-legacy studio-desk`) · floor **3** (`postgresql redis sentinel`) · `core` selects **5** |
| **G2** | any corpus "N repos cloned" claim | **6** (`app sentinel storage messenger next-web-app studio-desk`) |
| **G3** | any corpus default-bring-up container claim | **5** = \|select(`core`)\| = `postgresql redis sentinel backend gotenberg` |
| **G4** | any corpus-cited `*_RPC_ADDR` value | **4**, all `http://backend:8083` |
| **G5** | any corpus-named migration target | **1** repo with `migrations: true` (`app`) |
| **G6** | half-fold split — compose-sets vs consumer-reads, per env var | see `D-M257x-59-4`; first instance `STORAGE_RPC_ADDR` |

Also measured and available to the build: **13** published port mappings, **7** `profiles:` lines.

**G1 is the one to build first, and the rule it encodes is: grade on "does it still SELECT something",
never on "does it still parse".** The silent no-op is the dominant new failure mode and nothing in the
current fence set can see it. **Measured at this open — and this corrects the briefing:**

```
PROFILE=core     -> 5 : postgresql redis sentinel backend gotenberg
PROFILE=graphql  -> 3 : postgresql redis sentinel      ('graphql' declared by NO service)
PROFILE=cms      -> 3 : postgresql redis sentinel      ('cms'     declared by NO service)
PROFILE=storage  -> 3 : postgresql redis sentinel      ('storage' declared by NO service)
```

The briefing said these *"exit 0 and start nothing."* They start **three**. Postgres answers, Redis answers,
sentinel is up, `docker ps` is non-empty — and the application is absent. **A stack that starts nothing is
an honest failure; a stack that starts its infrastructure and not its application is the silent no-op
wearing the costume of a partially-working stack.** `run_guide.md:88` and `setup_guide.md:441` promise 11
containers here, and `platform_repo.md:88` asserts the `graphql` profile **in bold**.

**Watch every assertion RED before trusting it** (§8 rule 5 — collect the mutant *before* running it; a
mutation that does not compile is not a RED fence) and **mutation-verify the fixtures too** (rule 2).

---

## `D-M257x-59-3` — §7 rule 4 gains a citation-safety half

**Question directed:** *what replaces §7 rule 4 for pin advances, now that schema-safety demonstrably does
not imply citation-safety?*

**Decision: rule 4 is not replaced — it is HALVED and completed.** Its schema half is correct and stays. A
second half is added: **a pin advance is not vetted until its CITATION delta is measured.**

**The case study is iter-58, and it is unusually clean.** The advance was vetted under rule 4 and passed on
every dimension it names — **0** migrations, **0** destructive DDL, **0** new hard-required config, **0** new
env reads, `STORAGE_RPC_ADDR` unchanged. All of that was true. It still moved **22 of 23** `main.go:N`
citations in the corpus, and the fence caught **1** — a **4.5%** catch rate.

> **Schema-safety and citation-safety are unrelated properties, and §7 rule 4 only ever measured the first.**

An advance can be perfectly additive at the schema and contract level and still relocate every line the
corpus points at, because **adding code moves the lines below it**. Rule 4's dimensions are all about
*removal*; the citation class is caused by *addition*.

**The new half, operationally:**

1. Before taking an advance, enumerate every corpus citation whose path resolves into the advancing repo.
   Bounded and cheap — the whole `main.go:N` set is **23**.
2. Re-resolve each at the new ref and classify **moved / dead / held**.
3. **The repair belongs to the advancing iter**, not to a routed-forward handler — the same rule P3 already
   states for refs (*the iter that detects the move re-points, in that iter, as its first act*). iter-58
   proved the deferral cost: 21 sites are still outstanding.
4. Record the three counts alongside rule 4's existing table, so an advance's record shows both halves.

`FIX-M257x-iter58-mainline-shift` (**21 of 22** outstanding) is the retrofit case and is iter-61's target,
so the rule is proven by use rather than by assertion.

---

## `D-M257x-59-4` — a half-landed fold gets a state of its own, recorded on both sides

**Question directed:** *how does the corpus record a HALF-LANDED fold?* (Currently: nowhere. And it will
happen again for messenger.)

**Decision: the map gains an 8th state token, `mid-fold`, which assertion C accepts ONLY when the row
carries a two-sided citation — the config side and the consumer side, each resolving.** G6 fences the pair.

**Why a new token rather than prose.** The map carries **two states per row** (prod / fresh local stack) over
a **7-token** vocabulary — `live-standalone · merged-into-app · running_but_unfederated · decommissioned ·
net-new · external · library`. **None of them is true of storage today**, so the honest options were a
wrong token or silence, and silence is what happened.

**The storage split, measured at this open:**

| side | state | evidence |
|---|---|---|
| compose service | moved to `profiles: [storage-legacy]` — **not** in `core`, so a default bring-up never starts it | `docker-compose.yml`, rationale in-comment at `:130-133` |
| compose env | **`STORAGE_RPC_ADDR` is set nowhere** — absent from `docker-compose.yml` *and* `.env_example` | grep, exit 1 both |
| `repos.yml` | `storage` **still present** — still cloned | 1 of the 6 entries |
| app `v1.366.0` | **still reads it** at `main.go:446`, `:524`, `:992` | grep at the pinned ref |
| app tools | **hard-require** it: `cmd/academyImport/main.go:235` and `cmd/academy-asset-upload/main.go:133` both `return … "STORAGE_RPC_ADDR is required"` | same |

So on every stack we currently run green, `os.Getenv("STORAGE_RPC_ADDR")` returns `""` and a storage client
is constructed against the empty string — **a failure deferred to call time, not boot time**, which is why
three green cold cycles did not surface it. Neither side of this split is recorded anywhere in the corpus.

**The row shape:** `mid-fold` + a config-side citation + a consumer-side citation, both resolving under
assertion F's existing resolver. A `mid-fold` row with only one side is a **finding**, not a row — that is
the whole point of the token, because a half-fold recorded from one side is exactly how this one went
missing.

**Messenger is next by the developer's own account**, so the state will be needed again before it is
finished being written. Building the token now costs one vocabulary entry and one assertion clause; prose-ing
the storage case costs the same work twice and leaves the second one undiscoverable.

---

## `D-M257x-59-5` — ordering

**Question directed:** *what is the ordering — remaining clause-5 work vs the harden pass vs anything else?*

**Decision, in dependency order:**

1. **iter-60 — build the sibling guard (`D-M257x-59-2`); G1 closes the `graphql`-profile class.**
   First because it is the cheapest win available, because its six denominators were measured **today** and
   will be stale within days at the observed commit rate, and because it converts three predicate classes
   from prose into derivation. Every subsequent clause-5 reading is otherwise measuring a corpus the
   platform is still moving under.
   **Pre-registered, therefore refutable:** G1 goes RED naming `graphql` across **17 files / 30
   occurrences**, plus `cms` and `storage`; after repair it is GREEN, and the reverse direction names any of
   the 8 legal profiles the corpus documents nowhere.
2. **iter-61 — land §7's citation-safety half (`D-M257x-59-3`) and spend it on the 21 outstanding sites.**
   Rule and first application in one iter.
3. **iter-62 — the map's `mid-fold` state + the storage row (`D-M257x-59-4`), G6 fencing the split.**
   Before messenger folds, not after.
4. **Then `/developer-kit:harden-mstone-iters`.** The counter **restarted at iter-58** after pass 15 closed
   `STABILIZED`, so it stands at **1 tik against a threshold of 10 — NOT due.** When due it owns the new
   guard's assertions, which are the AST/call-site shape the three standing `HARDEN-CAP-ACCEPTED` entries
   said the residue needs.
5. **Then the next paired reading** — the first ever taken against a corpus whose three largest predicate
   classes are fenced rather than prose. A zero reading is not arithmetically reachable before that.

**Why the fences come before the reading and not after.** With single-pass recall at 43–48%, a reading over
an unfenced corpus is sampling a pool that the platform refills faster than a repair pass drains it — the
**−72** net. Fencing a predicate removes its whole class from the pool *and* stops it refilling. Reading
first would spend the expensive instrument on findings the cheap instrument makes unrepresentable.

**Also open, and deliberately not settled here:** `FIX-M257x-iter56-assignment-flake` is **NOT DECIDED** —
passes at `v1.366.0` are compatible with both hypotheses and it needs a failure **rate**, not another pass.
`CHECK-M257x-iter38-ai-act-classification` **needs an owner outside this milestone** and is not settled
here.

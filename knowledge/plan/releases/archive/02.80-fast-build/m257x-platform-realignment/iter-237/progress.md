**Type:** tik — under `TOK-08` (census a mechanical class exhaustively).

# iter-237 — are the environment variables the corpus calls critical actually read by anything?

## Why this is the third half of a runnable instruction

iters 235–236 closed the `make` target (aligned) and the `cd` directory (four repairs, one path). The
third input to a working stack is the **environment**, and it is the one that fails **silently**: a wrong
`cd` errors instantly, a wrong variable name gets you a stack that boots and does the wrong thing.

`CLAUDE.md`'s **"Critical environment variables"** list is a five-name hand-maintained tuple — exactly the
latent-drift construct `platform-alignment.md` § 2 names — and **nothing derives or fences it**.
`platform_predicate_guard` grades compose *profiles*; `demo_knob_guard` grades `DEMO_*` against the demo
parsers. Neither reads platform env-var names.

Env-var names are **exact tokens**, so iter-234's paraphrase refusal does not reach here: containment is
the correct test when the subject really is a literal.

## The instrument was wrong twice, and the controls are what caught it

Recorded because both errors are the *same* error iters 234–236 kept finding, in a third place:

1. **Read-set included markdown.** Grepping all tracked files made every known-dead `*_RPC_ADDR` appear
   "read by app" — because `app/CLAUDE.md` names them. Restricted to source extensions.
2. **Denominator omitted the tooling repo.** The first orphan list was topped by `STACK_PUBLIC_HOST`,
   `MOCK_CLERK`, `FAKE_FAPI_ROSTER`, `DEMO_*` — all **`rosetta-extensions`** variables, absent from the
   platform *by design*. Widened to `app`, `sentinel`, `platform`, `next-web-app`, `studio-desk`,
   `ant-academy`, both rext clones. Orphans fell **221 → 28 distinct** (448 → 37 sites).

**A census whose denominator is one repo will always find the other repos' variables missing.** Third
instance in four iters (`D-M257x-235-1`, `D-M257x-236-1`, here).

## Instrument validation — both controls behave exactly as the corpus predicts

| control | expectation | measured |
|---|---|---|
| `SKILLER_RPC_ADDR`, `CMS_RPC_ADDR`, `STORAGE_RPC_ADDR`, `JOBSIMULATION_RPC_ADDR`, `BACKEND_USERS_RPC_ADDR` — dead at `838d907` | absent from compose and `.env_example` | **5/5 `env_example=n compose=n`** ✓ (the strings persist in `app` Go source as legacy constants — which is precisely what `CLAUDE.md` already says: *"set nowhere in compose and read by nothing"*) |
| `AUTHORIZATION_ADDRESS`, `GOTENBERG_URL`, `JUDGE0_BASE_URL` — live | present in compose | **3/3 `compose=Y`** ✓ |
| `MESSENGER_ENABLED` — in-app gate, deliberately not in compose | absent from compose, present in source | **`compose=n`, in `app`+`platform` source** ✓ |

The dead/live controls separate cleanly on the compose axis, which is the axis the claim is about.

## Population

**381 distinct backticked `UPPER_SNAKE` names / 1,287 sites** across `corpus/**` + `CLAUDE.md`.
`platform/.env_example` @ `0c91421` declares **61**; `docker-compose.yml` names **45**.
**Orphans — named in the corpus, absent from every repo's source: 28 distinct / 37 sites** (1.9 % of
distinct names). Inspected, they are archived-doc references (`intelligence.md`'s
`DB_CONNECTION_BACKEND`/`_SKILLER`, `messenger.md`'s `SKILLPATH_RPC_ADDR`), psql client vars
(`PGHOST`/`PGDATABASE`), staging-script locals (`SMOKE_EMAIL`), and two regex truncations of prose
(`BACKEND_USERS_`, `JOBSIMULATION_`). **None is presented as a currently-required platform variable.**

## The finding: the list names the one Directus variable you CANNOT set, and omits the one you MUST

Graded individually — the only way to grade a five-item list — `CLAUDE.md`'s critical set came back
**4 of 5 clean and 1 wrong**:

| name | `.env_example` | compose | verdict |
|---|---|---|---|
| `GH_PAT` | **Y** | n | ✓ |
| `CLERK_SECRET_KEY` | **Y** | n | ✓ |
| `OPENAI_KEY` | **Y** | n | ✓ — the name really is `OPENAI_KEY`, not `OPENAI_API_KEY` |
| `VITE_CLERK_PUBLISHABLE_KEY` | **Y** | Y | ✓ |
| **`DIRECTUS_PUBLIC_BASE_ADDR`** | **n** | Y | ✗ **not a variable you set** |

`docker-compose.yml:53` **hardcodes** `DIRECTUS_PUBLIC_BASE_ADDR=https://content.anthropos.work`, and
`.env_example` never declares it — so putting it in your `.env` does **nothing**. The variable a reader
actually has to fill is **`DIRECTUS_TOKEN`**, which `.env_example:92` ships **blank** and which is the
classic *stack boots, catalog empty* failure. It was **not in the list at all.**

**The corpus already knew.** `corpus/architecture/external_services.md:133` states the compose-literal fact
outright. Only the list in the file **every session loads** disagreed — the same shape as iter-236's
`CLAUDE.md` self-contradiction, one section apart. Repaired toward the site that was already right, with
`DIRECTUS_TOKEN` promoted in and a retraction banner left in place of the removed name.

## Seal grading — `5373765`, sealed before any measurement

| id | prediction | outcome |
|----|---|---|
| `P-237-1` | ≥ 150 distinct env-var names in the corpus | **CONFIRMED — 381** |
| `P-237-2` | `.env_example` declares ≥ 40 | **CONFIRMED — 61** |
| `P-237-3` | ≥ 1 of the five critical names not in `.env_example` and not in compose | **PARTLY REFUTED — and the refutation is the finding.** `DIRECTUS_PUBLIC_BASE_ADDR` is absent from `.env_example` but **is** in compose, so it fails the prediction as written while being exactly the defect: it is *set for you*, which is why listing it as yours to set is wrong |
| `P-237-4` | ≥ 5 orphans absent from all source | **CONFIRMED — 28** (only after the denominator was fixed; the first reading's 221 was denominator, not evidence) |
| `P-237-5` | ≥ 1 orphan presented as currently required | **REFUTED — 0.** All 28 are archived-doc, client-tool or script-local names |

**3 confirmed · 1 refuted · 1 partly refuted.** The seal earned its keep twice: `P-237-3`'s wording would
have let a looser reading claim a clean pass, and `P-237-5` predicted a defect class that does not exist.

## Guard family

`24 GREEN · 0 RED · 0 could-not-check · 5 not-run` (`--platform stack-demo/platform --allow-not-run`),
re-run after the repair — unchanged.

## Close — 2026-08-10

**Outcome:** 381 corpus-named variables censused against a seven-repo denominator with both a dead-name and
a live-name control firing correctly. **28 orphans, 0 of them presented as required.** The one real defect
is in `CLAUDE.md`'s five-name critical list: **`DIRECTUS_PUBLIC_BASE_ADDR` is hardcoded by compose
(`:53`) and undeclared in `.env_example`** — you cannot set it — while **`DIRECTUS_TOKEN`, which ships
blank and gates the whole content surface, was missing from the list.** Repaired toward
`external_services.md:133`, which had it right all along.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-237-1` (the denominator is every repo, and the first reading is recorded),
`D-M257x-237-2` (the 28 orphans are classified and none repaired).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — no pytest section run; no rext code changed. One root document changed; guard
family re-run at platform reach, 24 GREEN / 0 RED.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → **new, and the strongest fence candidate this run.**
  The five-name list is now correct *and dated*. Every name is decidable against `.env_example` +
  `docker-compose.yml` at platform HEAD, with no paraphrase and no polarity problem — the cleanest
  mechanical fence surface the last four iters have surfaced.
- `ROUTE-M257x-237-hardcoded-vs-settable` → **new.** The corpus has no vocabulary distinguishing *a
  variable you must set* from *a variable compose sets for you*. The defect here was entirely that
  conflation, and it will recur wherever a doc lists "environment variables."
- All prior routes → open, unchanged.

**Lessons:**
1. **Grade a list item by item, or you grade nothing.** Four of the five names were fine; a whole-list
   check that answered "mostly yes" would have missed the one that matters.
2. **A variable being *present in compose* can be the evidence it is WRONG in a setup list.** Presence and
   settability are opposite here, and no existing guard draws that line.
3. **The denominator error is now four-for-four.** Every census this run began with too small a
   denominator (one Makefile, one repo's dirs, one repo's source). Ask *"what is this a number OF"*
   **before** the first reading, not after the first surprising result.

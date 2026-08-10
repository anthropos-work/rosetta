**Type:** tik (under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07))

# iter-247 — the env-var surface, fenced by INVERTING it

## Phase A — the obvious fence is a trap, and the substrate said so before anything was built

`ROUTE-M257x-237` called this *"the strongest fence candidate"*, meaning: **every env name the corpus
mentions must exist in `.env_example`.** That fence would have been a false-RED generator.

| | |
|---|---|
| distinct `UPPER_SNAKE` tokens in the live corpus | **327** |
| occurrences | **1,213** |
| most-cited | `DIRECTUS_TOKEN` 46 · `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` 43 · **`CMS_RPC_ADDR` 38** · `GH_PAT` 31 · **`SKILLER_RPC_ADDR` 19** · **`STORAGE_RPC_ADDR` 18** |

Most are not environment variables at all, and **the most-cited ones are variables the platform
DELETED**. A must-exist rule would redden the corpus's most careful writing. **Absence is the point of
those sentences.**

**So invert the polarity: fence the ABSENCE claims.** They have exactly one reading, they are claims
about the platform, and the platform can falsify them at any time — which is precisely the class this
milestone exists for.

## Phase B — three measurements, each of which killed a draft

**1. Single-variable absence claims MIS-ATTACH.** 12 variables carry an absence phrase across 29 sites.
Graded against `stack-demo/platform`, **6 of the 12 are actually PRESENT** — and on inspection **all six
are co-location artifacts**:

| variable | where it really is | what the sentence's absence clause was about |
|---|---|---|
| `AUTHORIZATION_ADDRESS` | `docker-compose.yml:48` | the same sentence *says so*; the clause is about `*_RPC_ADDR` |
| `JUDGE0_BASE_URL` (5 sites) | `docker-compose.yml:59` | *"zero `*_RPC_ADDR` variables anywhere in compose"* |
| `MESSENGER_ENABLED` | `docker-compose.yml:87` | a long migration-map row carrying presence AND absence |
| `STORAGE_S3_BUCKET`, `_PUBLIC_BUCKET` | `:82`, `:83` | same row |
| `GRAPHQL_SCHEMA_FOR_GEN` | `.env_example:105` | *"read by nothing"* — a claim about **code**, and the sentence itself says it IS declared |

**0 of 12 absence claims are false.** Six were the instrument.

**2. Restricting to lines with exactly ONE env token does not fix it.** That leaves **2 sites**, and one
is *still* mis-attached — because the real subject of *"zero `*_RPC_ADDR` variables anywhere in compose"*
is the **glob**, which an `UPPER_SNAKE` pattern cannot see at all.

**3. The glob was the subject all along, and it is the stronger claim.** Family tokens in an absence
sentence: **13 distinct across 41 sites**, dominated by **`*_RPC_ADDR` — asserted absent at 23 sites
across 15 documents.** That is the corpus's most-repeated platform claim, made as service after service
folded into `app`. **One re-added RPC address falsifies 23 sites simultaneously**, and nothing was
watching.

## Phase D — the fence: `stack-core/env_absence_guard.py` (FENCE-M257x-iter247)

A backticked family glob, in an absence sentence, **that names the file it is about** → verified against
`docker-compose.yml` / `common.yml` / `.env_example` at the platform clone.

**Live verdict: 17 family-absence claims hold (`*_RPC_ADDR` ×17). Zero findings.**

Three scope rules, each measured rather than assumed:

* **The claim must name its own file.** Of 41 family-absence sites, **23 name a platform file and 18 do
  not** — and the 18 are overwhelmingly about the frontend or the rext tooling (`DEMO_NO_*`, `BUNNY_*`,
  `VITE_*`), where grading against platform compose is a false RED.
* **A `read by nothing` claim is about CODE, not configuration** — a category error to grade here.
  `GRAPHQL_SCHEMA_FOR_GEN` is declared in `.env_example` **and** read by nothing; both are true and only
  one is about this file.
* **A bare `no` is NOT an absence quantifier.** This was the guard's own only two findings, and both were
  false: *"no rebuild needed"* and *"with no `NEXT_PUBLIC_*` / env / compose override"* — the latter about
  the absence of an override **seam** for one value. Worse, one of the two sentences asserts the exact
  **opposite** (*"all five `NEXT_PUBLIC_CLERK_*` keys appear … before this edit too"*). Dropping bare `no`
  took the guard to zero without weakening a single real claim.

Plus two matching rules that keep it from crying wolf: a **commented-out** setting is not a setting
(`.env_example` documents retired variables in comments), and a bare **mention** is not a setting (the
name must be followed by `=` or `:`). **17 unit tests**, including both mutation directions, an
every-site test, and an anchored-not-substring test (`*_RPC_ADDR` must not match `..._RPC_ADDRESS`).

**And the guard's own test improved the guard:** the first draft reported *"claims `<family>` absent — but
it is SET at file:line"*, telling a reader where to look but not what they were looking for. It now names
the offending variable.

Wired into `guard_family` with the same **no-default-platform** discipline as its siblings — a fidelity
check against the wrong reference passes. **28 GREEN → 29 GREEN.**

## Phase E — pre-registrations, graded after the last edit

| id | prediction | outcome |
|---|---|---|
| **P-247-1** | 2–4 of 12 single-variable claims actually PRESENT | **REFUTED — 6.** And all six were mis-attachments, so the *underlying* claim rate is **0 of 12 false** |
| **P-247-2** | ≥ 1 co-location artifact | **HELD — 6 of 6 of them** |
| **P-247-3** | no guard grades an env-var absence claim | **CONFIRMED** — and now one does |
| **P-247-4** | all 6 `*_RPC_ADDR` names genuinely absent | **HELD, 6 of 6** — `BACKEND_USERS`, `CMS`, `JOBSIMULATION`, `ROADRUNNER`, `SKILLER`, `STORAGE`, absent from compose **and** `.env_example` |
| **P-247-5** | **control** — the 5 critical names are declared | **HELD, 5 of 5** (`GH_PAT:2`, `CLERK_SECRET_KEY:3`, `OPENAI_KEY:10`, `VITE_CLERK_PUBLISHABLE_KEY:106`, `DIRECTUS_TOKEN:92`) |

## Close — 2026-08-10

**Outcome:** the environment-variable surface fenced — by **inverting** it. The obvious must-exist rule
was refuted from the substrate (the corpus's most-cited env names are ones the platform **deleted**), so
the fence grades **absence** claims instead, and **family globs** rather than single names because 6 of 12
single-variable phrases were co-location artifacts. `env_absence_guard` (FENCE-M257x-iter247) holds **17
claims — all `*_RPC_ADDR`, the corpus's most-repeated platform claim, asserted at 23 sites across 15
documents and previously checked by nothing.** Family **28 → 29 GREEN**. **0 of 12 absence claims were
false**; the corpus was right and undefended.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: exit-5
**Decisions:** `D-M257x-247-1` (invert the polarity — fence absence, not presence) · `D-M257x-247-2`
(grade FAMILY globs, not single variables, on the 6-of-12 mis-attachment measurement) · `D-M257x-247-3`
(a claim must name its own file, and a read-axis claim is a category error) · `D-M257x-247-4` (a bare `no`
is not an absence quantifier).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
`test_env_absence_guard` **17 passed**, plus `test_fence_registry_population` + `test_guard_family`
**87 passed / 0 failed** together. Guard family (`--platform`, from repo root): **29 GREEN / 0 RED /
0 could-not-check / 5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → **CLOSED**, though not as written: the *presence*
  list is still hand-maintained and dated, and this iter's measurement says a presence fence over it is
  the wrong instrument. The claim that would actually rot — the absence family — is now fenced.
- `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` → **HALF-CLOSED.** Environment
  variables are fenced; **frontend scripts/ports** (iter-238) remain unfenced.
- `ROUTE-M257x-237-hardcoded-vs-settable` → open, and this iter is fresh evidence for it: the corpus has
  no vocabulary separating *a variable you must set* from *one compose sets for you* from *one that is
  declared and read by nothing* — three distinct states the same list conflates.
- `ROUTE-M257x-245-guard-family-green-is-not-suite-green` · `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` (which-line half) ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` · `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-236-disclosure-scope-is-document-level` ·
  `ROUTE-M257x-235-fence-scope-is-unread` · `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **When a fence would redden careful writing, the polarity is wrong, not the corpus.** The must-exist
   reading was the obvious one and the substrate refuted it in one measurement — the three most-cited
   names are deliberately-absent variables.
2. **A claim's SUBJECT is not always a name; sometimes it is a pattern.** Two drafts died attaching
   absence phrases to `UPPER_SNAKE` tokens before the glob turned out to be what the sentences were
   about all along — and the glob is the *stronger* claim, covering 23 sites at once.
3. **A quantifier a fence acts on must be unambiguous alone.** Bare `no` reads as *"no rebuild needed"*
   and *"no override seam"*; it even appeared in a sentence asserting the opposite of absence.
4. **The corpus was right and undefended.** 0 of 12 absence claims false, 17 family claims holding — the
   deliverable is not a repair, it is that these can no longer rot silently.

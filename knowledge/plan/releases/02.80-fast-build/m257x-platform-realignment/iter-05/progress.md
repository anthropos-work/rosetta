---
milestone: M257x
iter: 05
---

# iter-05 — progress

**Type:** tik, under `TOK-01` step 1. Declared **2-step planned shape** (see overview.md), so the
scope-creep tripwire counts against that, not a single-target tik.

## Target 1 — `FIX-M257x-autoverify-skillpath-schema`

`stack-verify/lib/readiness.sh:probe_postgres_schemas` carried a hand-written list:

    local expected=(public sentinel cms jobsimulation skillpath extensions)

It demanded **`skillpath`** — the schema iter-02 correctly stopped creating, because origin `repos.yml` no
longer lists the repo at all. Measured on the live stack, the DB holds
`auth · cms · directus · extensions · jobsimulation · public · sentinel` and **no `skillpath`**, exactly as
the platform dictates. So the probe failed on a *correctly migrated* stack.

**This is the same defect as the hand-maintained atlas tuple iter-02 removed — living in the verifier
instead of the migrator.** Third instance of `platform-alignment.md` §8 rule 3 inside this one milestone
(after the v2.1 migrate-set test and the iter-02 schema-create test).

Fixed by **deriving** the expected set from `repos_yml_schemas_to_create` — the same
`stack-core/lib/repos_yml.sh` the migrator uses — so the probe and the thing it verifies cannot disagree
about which schemas a stack should have. A missing source now **fails loud** instead of falling back to a
stale list; that fallback is precisely how this survived. `REXT_REPOS_YML` overrides the discovered path.

**Proven live, A/B against the running demo-1:**

    NEW: ok: all expected schemas present            (rc 0)
    OLD: fail: missing schemas: skillpath            (rc 1)

## Target 2 — `FIX-M257-directus-coldstart-order` (carried since M257 iter-02)

    docker logs demo-1-directus-1
    Error: connect ECONNREFUSED 172.18.0.5:5432

Directus was the **only service on the stack with neither a restart policy nor a readiness dependency**, so
it raced Postgres on the cold bring-up, died, and — `restarts=0` — stayed `Exited (1)` for the life of the
stack, making `directus HTTP 000000` a guaranteed ✗ on every cycle. Carried since M257 iter-02 as
"platform-shape-dependent" and **never reproduced until a genuinely cold host ran one**.

**Hypothesis proven before fixing:** started by hand once Postgres was up, it came straight up and served
`/server/health` **200**. Purely ordering — not config, not schema.

Fixed in `stack-core/gen_override.py` with the platform's **own** convention for every DB-backed service in
its `docker-compose.yml` — `restart: on-failure` + `depends_on: postgresql: condition: service_healthy` —
copied rather than invented, so there is one shape to reason about.

## Live result — autoverify FAILED count 3 → 2

    ✓ verify live: all liveness + readiness probes passed      <- both targets cleared
    ✓ container liveness: all 16 expected container(s) running <- was 15 + directus Exited(1)

The two remaining are both already-routed and neither is new:

| remaining ⚠ | disposition |
|---|---|
| hiring org under-set-dressed (5 positions / **0** candidate sessions) | **downstream of the `jobsimulation.*` 42P01s** — `REPOINT-M257x-jobsim-writes`. autoverify's own text names this cause first. Not a separate defect |
| AI Academy serves `:13077/library` but renders **no course cards** | `FIX-M257x-academy-not-serving`, **re-characterized**: at bring-up this check PASSED (*"real course cards"*) and the startup probe failed; ~90 min later it serves but renders nothing. The FS-published demo-patch appears to have reverted under a long-running stack. A drift-over-time symptom, not a bring-up one |

## Side discovery — my own iter-04 edit broke a corpus fence, and the fence caught it

Running the wider suites surfaced `stack-core/tests/test_demo_knob_guard.py` **5 FAILED**:

    STALE ANCHOR: DEMO_NO_VERIFY cites `up-injected.sh:2523`, the parser reads it at `up-injected.sh:2541`
    …7 anchors, every one shifted by exactly +18 lines

iter-04's fix added **18 lines** to `up-injected.sh`, so every `file:line` citation in
`corpus/ops/demo/demo-up-defaults.md` below the edit point silently stopped pointing at the real read.
That doc is fenced against the parsers **in both directions** precisely so this cannot rot — and it worked.
Repaired with the guard's own `--fix` (`rewrote 7 'Read at' citation(s) from the parsers`); guard now
`OK — the defaults table and the parsers agree, both directions`, 27/27 tests green.

**Worth stating plainly: this was my regression, and I found it only because I ran the wider suite.** The
iter-04 close reported green on the suites it ran. §5 rule 8's sibling: *a suite you did not run is not a
suite that passed.*

## A third self-inflicted patch error, same family as the others

The scripted edit that added `REXT_REPOS_YML` to the test harness anchored on
`env = {**os.environ, "PATH": self.bin ...}` and used `replace(..., 1)`. That string occurs **six times**
in the file; the edit landed in an unrelated autoverify harness at line 346 instead of `_run_probe` at 962.
The tests then failed with *"cannot derive the expected schema set"* — which read like a bug in the fix and
was actually a bug in the patch. Caught by reading the error's own path rather than re-editing blindly, then
corrected with `assert s.count(anchor) == 1` on **both** the revert and the re-apply.

`platform-alignment.md` §8 already requires an **exactly-once anchor** for demopatch manifests. It applies
to any scripted edit, mine included.

## Suite state

| suite | result |
|---|---|
| `stack-core/tests` | 341 passed / 5 failed → **all 5 were the knob-guard anchors, now fixed → 346 green** |
| `demo-stack/tests/test_tooling.py` | 175 passed |
| `stack-verify/tests` | 207 passed / 3 failed — the 3 are `test_e2e_collection_integrity` (needs the Playwright/npm toolchain), **reproduced on the pristine control clone**, so pre-existing |
| `stack-core/tests/test_gen_override.py` | 21 passed; directus pair mutation-verified RED |
| schema probe tests | 5 passed; mechanism-pinned + mutation-verified |

## Close — 2026-07-31

**Outcome:** both planned targets landed and were proven on the live stack — autoverify FAILED **3 → 2**,
`verify live: all … probes passed`, containers **15 → 16/16**. The long-carried
`FIX-M257-directus-coldstart-order` is root-caused (a pure startup race) and fixed with the platform's own
convention; the schema probe now DERIVES its expectation from `repos.yml` instead of a hand-written list.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n — (4)
user-blocker: n (no guard refused; the one wider-suite regression was my own and was fixed in-iter; both
trees clean) — (5) cap-reached: n (2 tiks of 5) — (6) protocol-stop: n — Outcome: continue
**Decisions:** none new at milestone level.
**Metric:** clauses met **0/5 → 0/5** (clause 1 needs **3 consecutive fully-green cold cycles**; 2 checks
still fail). Sub-progress: autoverify FAILED **3 → 2**; readiness probes **1 failing → all passing**;
containers **15 → 16/16**; a route carried since **M257 iter-02** closed.
**Side-deliverables:** the `demo-up-defaults.md` anchor repair (7 citations) — a regression this iter's
predecessor introduced, caught by the corpus fence and fixed with the guard's own `--fix`.
**Routes carried forward:**
- `REPOINT-M257x-jobsim-writes` → **the next tik**. It is now the single largest remaining item and owns
  BOTH gate clause 4 and the last substantive autoverify ⚠.
- `FIX-M257x-academy-not-serving` — **re-characterized** to "renders no cards after the stack has been up a
  while", which is a different (and more interesting) defect than the bring-up-time one iter-04 saw.
- `HOST-M257x-toolchain` — residual grows slightly: the Playwright/npm e2e toolchain is absent too, so
  `stack-verify/tests/test_e2e_collection_integrity.py` (3 tests) cannot pass on this box.
- `DOC-M257x-claude-md-knob-count` (NEW, small) — the guard reports **30** env knobs; root `CLAUDE.md` says
  *"all 27 `DEMO_*`/`STACK_PUBLIC_HOST` env knobs"*. The corpus doc agrees with the parsers; `CLAUDE.md`
  does not.
- All other routes unchanged.
**Lessons:**
- **A fence only protects what someone runs.** The knob guard caught an iter-04 regression perfectly — one
  iter late, because iter-04 ran the suites it thought were relevant. The wider run is the cheap part.
- **Reproduce before fixing, even when the fix is obvious.** Starting directus by hand cost 25 seconds and
  turned "directus is broken" into "directus raced Postgres", which is what made the one-line fix the
  *right* one rather than a plausible one.
- **`replace(x, 1)` on a non-unique anchor is the scripted-edit version of a probe that cannot fail.**
  `assert count == 1` before every scripted replace. §8's exactly-once anchor rule is not just for
  demopatch manifests.

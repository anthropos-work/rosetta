---
milestone: M257x
iter: 02
---

# iter-02 — progress

**Type:** tik, under `TOK-01` step 2. Shape per `corpus/ops/platform-alignment.md` §2 (the mechanism),
§7 (re-point procedure), §8 (fence).

## Step 0 re-survey — TOK-01's next-tik direction did not survive the machine move

`TOK-01` directed iter-02 to `FIX-M257x-rext-pin` on **odysseus**. Both halves were measured and both are
stale (full table in this iter's `overview.md`). The pin is **clean**: `.agentspace/rext.tag` reads
`fast-build-m257x-iter-01` → `31d2b5df`, which is **on origin** and **== `origin/main`** (0 behind, 0 ahead),
and there is no `stack-demo/`, so the clone-vs-SoT mismatch that made `ensure-clones.sh:94-101` FATAL cannot
exist. **Target absorbed; substituted to TOK-01 step 2 under the same strategy.** Recorded as `D-M257x-4`.

## What was measured (§4 signals 1–3, against origin HEAD)

Fetched `repos.yml` from **platform origin HEAD `1e8e7540`** by shallow blobless clone — there is no local
`platform` clone on this box at all, so no stale copy could contaminate the read.

    app            migrations: true   schema: public    <-- the ONLY one of each
    cms            migrations: false  (schema: key DELETED)
    jobsimulation  migrations: false  (schema: key DELETED)
    roadrunner · sentinel · storage · messenger · next-web-app · studio-desk · graphql-wundergraph
                   migrations: false
    skillpath      ABSENT ENTIRELY

**The hand-maintained tuple was wrong on 3 of its 4 entries.**

§7.1 write-surface scan (live code vs comments, positive control run in the same pass): `jobsimulation.*`
appears in **110** files; `stack-seeding/cmd/stackseed/main.go:45-105` holds **live** reset/seed entries for
~12 tables. So the anticipated escalation condition fired: **the legacy schema creation could not be removed
in this iter** without shipping a knowingly-broken bring-up. That is why the debt is *declared and fenced*
rather than deleted — the write re-point is a separate iter.

## Deliverables landed

| what | where |
|---|---|
| the derivation (machine-readable fields only) | `stack-core/lib/repos_yml.sh` (new) |
| demo migrate script derives its set | `demo-stack/migrate-demo.sh` — 3 hardcoded sites → 0 |
| dev migrate script derives its set | `dev-stack/migrate-dev.sh` — 2 hardcoded sites → 0 |
| the fence, mutation-verified | `stack-core/tests/test_migration_derivation_fence.py` (new, 14 tests) |
| fixtures reconciled to the new dependency | `demo-stack/tests/test_host_prereqs_m215.py` |
| the test that pinned the WRONG contract, replaced | `dev-stack/tests/test_dev_stack.py` |

Silent skip → loud: `[ -d "$DEV/$r" ] || continue` became a named failure. A repo in the **derived** set is
one `repos.yml` says owns migrations, so an absent clone is a real problem and now says so (`mig_fail=1` on
the demo side).

`skillpath` is **gone from executable code in both scripts** — the canary §2 predicted.

## The fence, watched going RED (§8 — "each must be watched going RED before it is trusted")

Run in an isolated copied tree; the real checkout was never mutated.

| mutation | result |
|---|---|
| restore the literal 4-tuple in `migrate-demo.sh` | **RED** (2 tests) |
| grow `REXT_TRANSITIONAL_SCHEMAS` by one entry | **RED** (4 tests) |
| replace the derivation with a literal in `migrate-dev.sh` | **RED** (1 test) |
| loosen the parser so it reads prose comments | **RED** (1 test) — *after* the fixture was fixed |

**The 4th mutation initially passed, and that is the finding.** The prose-comment fixture placed its lying
comment *above* the first `- name:` line, where the parser's per-entry reset makes it harmless — so a
derivation loosened to match `migrations:` anywhere still passed. A fixture that cannot fail is the M256
reports-success-without-checking class, occurring **inside a test written to prevent exactly that**. Fixed by
moving the lying values inline on a field and into a commented-out `schema:`, which is where the real
`repos.yml` actually puts them.

## A test that passed by reading its own refutation

`dev-stack/tests/test_dev_stack.py::test_migrates_the_four_merged_services_and_never_skiller` (v2.1 M209)
**required** all four pairs to be present. It encoded the drift as a contract. After the loop was derived it
**still passed** — satisfied by the tuple appearing in the new explanatory *comment*, because it grepped
whole-file source and cannot tell code from prose.

Replaced by `test_migrate_set_is_derived_from_repos_yml_never_hardcoded`, which asserts against the **loop
body** and the **derived set**, and is mutation-verified RED on both a restored tuple and a
literal-but-plausible substitution.

## Host survey (routed, not fixed here)

Measured at iter open: **no container runtime at all** (docker · podman · colima · nerdctl · lima · orbstack
all absent), no `gh`, no `psql`, no `tailscale`. **Docker Desktop was installed by the user mid-iter** and
re-verified independently: `server=29.6.2 · linux/arm64 · overlayfs · cpus=8 · mem=12528664576 B`. Clause 1
is now *unblocked-in-principle* but there is still **no `stack-demo/` workspace**, and no bring-up has been
attempted. `pytest` is still absent, which is why two mutation batteries report failures unrelated to this
iter (below).

## Test gates

| suite | result |
|---|---|
| `test_migration_derivation_fence.py` (new) | **14 / 14 OK** |
| `dev-stack/tests/test_dev_stack.py` | **90 / 90 OK** (3 skipped) — was 2 failures, both mine, both fixed |
| `demo-stack/tests/test_host_prereqs_m215.py` | **41 / 41 OK** — was 1 failure, mine, fixed |
| `test_migrate_dev_live.py` · `test_frontend_build.py` · `test_apply_authn.py` · `test_test_collection_fence.py` | OK |
| `stack-core` remainder | 10 / 14 modules OK |
| `test_m220_mutation_battery` · `test_m255_mutation_battery` | **FAIL — environmental, not this iter** |

Both failing batteries shell out to `python3 -m pytest` (`test_m255…:165`, `test_m220…:386`), and **pytest is
not installed on this box** — producing rc≠0 with *zero named failures*, which is exactly the observed
signature (`m255`'s own unmutated-baseline check fails first). Neither battery references any file this iter
touched (`grep -c` → **0** in both). Routed as `HOST-M257x-toolchain`.

## Close — 2026-07-31

**Outcome:** the hand-maintained 4-tuple is gone from both migrate scripts — the migration set is derived
from `repos.yml`'s machine-readable fields, the M810 silent-skip time bomb is disarmed, `skillpath` is
removed, and the residual non-derivable schemas are declared as fenced debt. Gate clauses unchanged at 0/5,
exactly as the plan predicted.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik; 0 prior tiks, so the 3-no-prog
window cannot be full) — (3) re-scope: n (zero alignment attempts invalidated by mid-milestone platform
commits; the trigger needs two) — (4) user-blocker: n (Docker absence did not change what code landed, and it
was resolved mid-iter regardless; test gates green) — (5) cap-reached: n (1 tik of 5) — (6) protocol-stop: n —
Outcome: continue
**Decisions:** `D-M257x-4` (TOK-01's pin premise refuted by the machine move), `D-M257x-5` (the transitional
schema set is declared debt, not design), `D-M257x-6` (a source-grep fence cannot tell code from prose)
**Metric:** clauses met **0/5 → 0/5** (delta 0, as predicted). Sub-progress: hardcoded platform repo:schema
tuple sites **5 → 0**; underived dead-schema creations **3 → 0**; declared transitional debt **3 → 2**
(`skillpath` eliminated); new mutation-verified fence **+14 tests**.
**Side-deliverables:** the rext pin advanced to `fast-build-m257x-iter-02` and **verified on origin**
(`git ls-remote`), keeping the D-v28-15 git-ignored-pin hazard closed.
**Routes carried forward:**
- `HOST-M257x-stack-demo` → iter-03: there is no `stack-demo/` workspace; clause 1 needs one. Docker is now
  present, so this is the next executable step.
- `FIX-M257x-vmram-gib-unit` → iter-03: `up-injected.sh:258-262` floors bytes to integer GiB, so a VM set to
  the documented "12 GB" (decimal) measures 11.67 GiB → floors to 11 → trips the non-fatal `< 12 GiB` warning.
  A doc/code **unit mismatch** (decimal GB vs binary GiB), never re-measured — this milestone's own subject
  matter. Non-fatal; deliberately **not** opened as a 3rd line of investigation here (scope-creep tripwire).
- `HOST-M257x-toolchain` → iter-03: no `pytest`, no `gh`, no `psql`, no `tailscale` on this box. Two mutation
  batteries cannot run at all; the release's own test story is unmeasurable until pytest exists.
- `REPOINT-M257x-jobsim-writes` → later tik: ~12 `jobsimulation.*` tables (9 written) in
  `stack-seeding/cmd/stackseed/main.go`. Until they are re-pointed the transitional debt cannot shrink, and
  gate clause 4 cannot be met.
- `FIX-M257x-migrate-dev-swallows-atlas` → later tik: `migrate-dev.sh`'s atlas loop still does
  `>/dev/null 2>&1` and logs every failure as "non-fatal migration warnings" — the exact M215-F8 masking class
  its demo twin already fixed. Observed while editing; not opened (tripwire).
- Unchanged from iter-01: `FIX-M257x-pin-stale`, `DOC-M257x-guard-severity`, `DOC-M257x-subgraph-count`,
  `DOC-M257x-ai-labs-repo`, `DOC-M257x-livekit-agents`, `DOC-M257x-repo-states`, `KB-1`.
**Lessons:**
- **A strategy step whose premise dissolves is a finding, not a no-op.** TOK-01 step 1 was correct when
  written and false three hours later; the re-survey is what caught it, and skipping it would have spent an
  iter fixing a pin that was already clean.
- **A source-grep fence cannot distinguish code from prose** — so the comment you add to explain a removal
  can satisfy the very test that was supposed to forbid it. Assert against the parsed construct (the loop
  body, the derived value), not the file.
- **Mutation-verify the fixtures, not just the fence.** The one mutation that did not go RED here failed
  because the *fixture* was unfalsifiable, not because the fence was wrong.
- **"Derive it, or fence it" needs a third clause: *and declare what you cannot do either to*.** `sentinel`
  proves derivation alone is unsafe; the transitional debt proves deletion alone is unsafe. Writing the
  residual down, with a per-entry reason and a test that forbids growth, is the honest middle.

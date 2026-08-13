---
milestone: M257x
iter: 04
---

# iter-04 — progress

**Type:** tik, under `TOK-01` step 1 (*"unblock the gate's instrument"*).

## Step 0 re-survey — TOK-01's next-tik direction was stale again, in a new way

`TOK-01` said iter-02 should fix a FATAL rext-pin mismatch on **odysseus**. Both halves are now dead
(machine move; pin clean since iter-02; guard confirmed MATCHING live in iter-03). Measured instead:

| probe | reading |
|---|---|
| `stack-demo/` | 13 repos + both lock files — `HOST-M257x-stack-demo` genuinely DONE, not re-run |
| `.agentspace/rext.tag` vs consumption clone | `fast-build-m257x-iter-02` == `54bccf7`, consistent |
| target `.env` files | **5 of 5 ABSENT** — never provisioned on this host |
| `docker ps` | up, **0 containers**; `demo-1` free |

## Phase 0d pre-flight — PASS, and it was the right 30 seconds

`stacksecrets check --demo` **before** the long operation: **Critical 100.0% (exit 0)**, Overall 65.5%.
The 65.5% is non-critical breadth (LiveKit, Stripe, ElevenLabs, Sentry, FontAwesome, several AI providers).
Recorded so a later failure in one of those reads as *a known-absent optional key*, not a new defect.

Provisioned values-blind: **30 written, 2 blanked, 0 skipped**; `DIRECTUS_TOKEN` written **blank** on both
repos that carry it (the strip-on-non-prod / fix16-17 non-rearm safety, observed working).

## The first bring-up ever attempted on this host — and what it found

    START 16:30:04 … EXIT=128 at 16:30:29

**25 seconds, exit 128, and no error message at all.** The log's last line before the abort was
`inject disarmed colony into app`. Two independent defects stacked, and the second is what made the first
unreadable.

### Defect 1 — a PRIVATE repo cloned anonymously (`stack-injection/apply-authn.sh:28`)

    git clone --quiet --depth 1 --branch "$VER" https://github.com/anthropos-work/colony.git …

colony is private. An anonymous HTTPS clone succeeds **only** on a host whose git carries a credential
helper or an `url.insteadOf` rewrite. Measured on this box: **no credential helper, no `insteadOf`, and
`ssh -T git@github.com` authenticates fine as `kiralise`.** So git exits **128** —
`could not read Username for 'https://github.com'`.

Every *other* rext acquisition already uses SSH — `ensure-clones.sh:54,164` clone `git@github.com:$ORG/…`
and the failure text even says *"Check SSH access (`ssh -T git@github.com`)"*. This one line was the lone
outlier, and it was invisible for as long as nobody started from a clean box. Generalized as **Trap E** in
the protocol doc.

### Defect 2 — the call site threw the diagnosis away (`demo-stack/up-injected.sh:1690`)

    "$HERE/../stack-injection/apply-authn.sh" "$dst" >/dev/null 2>&1

Under `set -euo pipefail` that aborts the whole bring-up with the applier's explanation sent to `/dev/null`.

**The part that matters:** this is the *same* masking class **M217 already fixed on the very next call in
the same function**. The comment M217 left at `:1703` reads *"CAPTURE the applier's stderr instead of
discarding it. This call used to end in `>/dev/null 2>&1` … the demo silently shipped a 76-second members
grid."* The sibling one line above was never swept. Generalized as **§5 rule 9**.

## What landed (rext — zero platform-repo edits)

`fix(M257x/04)` — acquisition + surfacing:
- SSH first (the rext convention), `GH_PAT`/`GH_TOKEN` HTTPS fallback for a token-but-no-key host.
- **git's own exit code propagated verbatim**, not flattened to 1 — the pre-existing test
  `test_clone_failure_aborts_before_the_swap` pins exit 7 from its stub, and preserving that kept a
  correct existing contract green instead of editing the test to match me.
- The token is **redacted** from the diagnosis, and `vendor-colony/.git` is removed so a credential in a
  remote URL can never reach the build context the app image `COPY . ./`s.
- The call site captures + surfaces + names the step; still fatal (an app built against the real colony
  rejects every Clerkenstein token → every demo login 401s) but never silent.

Fenced by **6 new acquisition tests + 3 call-site tests**, the latter asserting against the *parsed
construct* (the unique non-comment call line) rather than a whole-file substring, per §8. The harness now
scrubs `GH_PAT`/`GH_TOKEN` so it cannot inherit a developer's real token and silently exercise a different
branch. **4 mutations RED-proven**, clean restore verified by diff.

## Two of my own probe errors, caught by the discipline rather than by luck

- The redaction test's git stub echoed `$3` as "the URL" — `$3` is `--depth`. Caught **by its own positive
  control** (`assertIn("could not read from", …)` passed while the redaction assert failed), which is
  exactly what §5 rule 2 exists for.
- A `grep -c` I used to confirm the fix reached the consumption clone returned `0` on a file that plainly
  contained the string. Re-measured with `-F`: present at 5 sites. Had I stopped at the first reading I
  would have reported the re-point as failed. §5 rule 3, on myself.

## Side discovery — two offline guards were RED or skipping, and both read as green

Neither is caused by this iter; both were confirmed **on the pristine consumption clone** (no iter-04
edits) before being touched, so the attribution is measured, not assumed.

1. **`test_migrate_demo_schema_create_is_set_e_guarded` has been RED since iter-02.** It anchored on the
   literal `CREATE SCHEMA IF NOT EXISTS extensions`; iter-02 correctly *derived* the schema list from
   `repos.yml`, so the literal became `CREATE SCHEMA IF NOT EXISTS "$_s"` in a loop. The test pinned the
   **contents** of the list, not the **mechanism** it protects (the `|| log` guard) — a second instance of
   the rule iter-02 itself wrote into §8. Re-anchored on the `CREATE EXTENSION` line ending the same exec,
   plus an assert that the list is still derived at all. Mutation-verified still-RED without the guard.
2. **`test_all_three_scripts_are_shellcheck_clean` was skipping** — shellcheck absent on this host. A skip
   next to a wall of dots reads as a pass. Installed shellcheck 0.11.0; it immediately reported SC1091 on
   the `. stack-core/lib/repos_yml.sh` **iter-02 added**. Fixed at source with
   `# shellcheck source-path=SCRIPTDIR` (a bare `source=` resolves against the runner's cwd — clean for the
   test, dirty for a developer; verified clean from two different cwds) and the baseline now runs `-x`.

Generalized as **§5 rule 8: a check that SKIPS reads exactly like a check that PASSES.**

## Host toolchain — `HOST-M257x-toolchain` partially closed (Fate 1, not deferred)

The routed item said "no pytest/gh/psql/tailscale". The moment it became blocking — a test existed for the
exact file being changed and could not run — it was closed rather than re-deferred:

| tool | before | now |
|---|---|---|
| `pytest` | absent | **8.4.2** in `/tmp/rextvenv` |
| `shellcheck` | absent | **0.11.0** (brew) |
| `gh` · `psql` · `tailscale` | absent | still absent — not needed by this iter; stays routed |

`/tmp/rextvenv` is ephemeral; a durable location is routed forward rather than pretended-permanent.

## Suite state

| suite | result |
|---|---|
| `stack-injection/tests` | **267 passed, 9 skipped, 0 failed** |
| `stack-injection` + `demo-stack` (full) | 1270 passed / 8 failed / 18 skipped — **all 8 reproduced on the pristine control clone**, so none is iter-04's |
| of those 8 | 1 was iter-02's regression (**fixed here**); 2 are demopatch sha-drift vs the live clone (`CHECK-M257x-demopatch-pristine` territory, and the bring-up preflight independently reported both as `SELF-HEALABLE — stale baseline`); 5 need a live docker/postgres stack |
| `corpus_index_guard` | OK — 83 docs, 6 index-bearing dirs |
| shellcheck `-x` on all three scripts | clean from any cwd |

## The second bring-up — it COMPLETED, and the milestone's founding hypothesis fired live

    START 16:46:46 … EXIT=0 at 17:04:59      (18m 13s, cold, 15 containers)

The first bring-up to complete in this milestone. Past the injection, five service images built, compose up,
migrate, policy, UI tier, set-dress, cockpit, autoverify.

**`migrate-demo.sh` ran iter-02's derivation for real, in a real bring-up:**

    schemas: extensions sentinel public cms jobsimulation
    atlas migrate the derived migration set … : app:public
      app ok

And `cms`, `jobsimulation` and `roadrunner` containers all came up in the default profile while owning no
migrated schema — `D-M257x-3`'s `running_but_unfederated`, observed rather than inferred.

### The time bomb detonated on schedule — 7 seeders, 3 relations, 42P01

    stackseed: 7 seeder(s) failed
      content-stories · hiring-funnel · jobsim-sessions · personas · activity · succession · ai-readiness-funnel

    relation "jobsimulation.sessions"                     does not exist   (5 surfaces)
    relation "jobsimulation.activity_events"              does not exist
    relation "jobsimulation.interview_extraction_results" does not exist

**This is `REPOINT-M257x-jobsim-writes` / gate clause 4, no longer a prediction.** And note *why* it fired
now: `platform-alignment.md` §2 says rext used to create the legacy schemas **and migrate them out of the
still-cloned legacy repos**, which is what kept the writes working. iter-02 correctly derived the migration
set, so `jobsimulation` is now created as **declared transitional debt and left EMPTY** — and every write
into it fails loudly. §2 predicted *"13 write targets 42P01 at once"*; measured, it is 3 distinct relations
across 7 surfaces.

That is the intended behaviour of a correct fix: the debt moved from **latent** to **visible and loud**. It
is also why clause 4 cannot be met by inspection — this bring-up is the fence going RED.

### Autoverify — 3 checks FAILED, stack UP (non-fatal, correctly)

| check | reading |
|---|---|
| ✓ backend `/api/health` 200 · casbin_rules **1251** · `directus_collections` 21 · directus DB per-stack-local | pass |
| ✓ taxonomy replayed **`public.skills` = 42790** · demo-patches all applied (none refused/skipped) · frontends are this run's · cockpit answering · fake-FAPI answering | pass |
| ✗ `postgres-schemas` — **`missing schemas: skillpath`** | **a stale rext assertion**: it demands the schema iter-02 correctly REMOVED (absent from origin `repos.yml`). §8 rule 3 again — a check pinning the drift as a contract, so a correct change has to argue with it. NEW: `FIX-M257x-autoverify-skillpath-schema` |
| ✗ `directus` HTTP 000000 | `demo-1-directus-1` **exited(1)**. NEW: `FIX-M257x-directus-container-exit1` |
| ✗ ant-academy not serving on `:13077` after 120 s | its own log says `✓ Ready in 193ms`. NEW: `FIX-M257x-academy-not-serving` |
| ⚠ hiring org under-set-dressed (5 positions / **0** candidate sessions) | **downstream of the 42P01s** — and autoverify's own warning text names this exact cause first (*"M257/B1 was precisely that"*). Not a separate defect |

The `< 12 GiB` VM-RAM warning fired exactly as `FIX-M257x-vmram-gib-unit` predicted (11 GiB) — non-fatal,
did not block.

### Two routed CHECK items answered as a by-product

- **`CHECK-M257x-demopatch-pristine` — benign, as iter-03's correction suspected.** This run logged
  `demopatch R1: ensured pristine <manifest>` for all 23 with **no ⚠ at all**, and autoverify independently
  reported `demo-patches: all applied (none refused, none skipped)`. iter-03's 23/23 `skipped/failed`
  warnings were the first-ever-bootstrap "nothing to revert" case. **The observability defect stands** — the
  message still collapses `skipped` and `failed` into one string at one severity — so the item stays open
  narrowed to that, not to a suspected refusal.
- **`CHECK-M257x-pin-state-on-fresh-clone` — explained, not a false positive.** `clones.pin.json` pins
  `platform` at `28c5f0dd` while the fresh clone of `main` is at `1e8e7540`; `graphql-wundergraph` likewise.
  `pin-drift` is therefore **correct** — the *pin file* is stale, not the clone. The escalation risk under
  `DEMO_FRESHNESS_STRICT=1` is real but would be flagging a true condition.

## Close — 2026-07-31

**Outcome:** the first bring-up of this milestone to COMPLETE (18m 13s cold, 15 containers, autoverify 3
FAILED / non-fatal) — after root-causing and fixing the two stacked defects that made a clean host
unbuildable; and the milestone's founding hypothesis fired live as **7 seeders / 3 relations / 42P01**,
turning clause 4 from a prediction into a measurement.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (nothing was
invalidated by a new platform commit; the platform did not move under us) — (4) user-blocker: n (no guard
refused a decision, no unrelated suite regressed — all 8 wider-suite failures reproduce on the pristine
control; tracked tree clean) — (5) cap-reached: n (1 tik of 5 this session) — (6) protocol-stop: n —
Outcome: continue
**Decisions:** none new at milestone level; `D-M257x-3`'s per-environment axis gained a second live
corroboration (cms/jobsimulation/roadrunner containers running with no migrated schema).
**Metric:** clauses met **0/5 → 0/5** (delta 0, as the plan predicted — clause 1 needs **3 consecutive**
green cold cycles and this one is not green). Sub-progress, all first-evers for this milestone: bring-up
**abort@25s → complete@18m13s**; containers **0 → 15**; migrate **never run → derived set applied live**;
seed **never run → ran and failed loudly at the exact predicted surface**; autoverify **never run → 3
named failures**.
**Side-deliverables:**
- `test(M257x/04)` — two offline guards that were RED-since-iter-02 and silently-skipping restored
  (separate commit; does not upgrade this iter's status).
- `HOST-M257x-toolchain` partially closed: pytest 8.4.2 + shellcheck 0.11.0 now on the box.
- First cold **build+bring-up** cost on this host: **18m 13s** for a full UI-tier demo — the leg M257's
  paused speed gate has to budget for, alongside iter-03's 673 s clone bootstrap.
**Routes carried forward:**
- `REPOINT-M257x-jobsim-writes` → **the next tik**, and no longer speculative: 3 named relations, 7 named
  surfaces, reproducible on demand. This is gate clause 4.
- `FIX-M257x-autoverify-skillpath-schema` (NEW) → next tik — autoverify demands a schema iter-02 correctly
  removed; a one-line stale contract, and it is one of the 3 failing checks blocking clause 1.
- `FIX-M257x-directus-container-exit1` (NEW) → next tik — blocks clause 1.
- `FIX-M257x-academy-not-serving` (NEW) → later tik — "Ready in 193ms" but never answers; not on clause 1's
  critical path but it is a ✗.
- `CHECK-M257x-demopatch-pristine` — **narrowed to the observability defect only** (split `skipped` from
  `failed` in the log); the refusal hypothesis is refuted.
- `CHECK-M257x-pin-state-on-fresh-clone` — **explained**; residual is only whether to refresh
  `clones.pin.json` or accept the true-positive.
- `HOST-M257x-toolchain` — residual: `gh`/`psql`/`tailscale` still absent; `/tmp/rextvenv` is ephemeral and
  needs a durable home.
- `FIX-M257x-vmram-gib-unit`, `FIX-M257x-migrate-dev-swallows-atlas`, and the iter-01 doc set: unchanged.
**Lessons:**
- **A correct fix can be what makes the bomb go off, and that is success.** iter-02's derivation is exactly
  why 7 seeders failed here. The prior state — rext creating and migrating the legacy schemas itself — was
  a working-but-wrong bring-up. Do not read the new noise as a regression; it is the debt becoming visible.
  Folded into §2 already ("derive it, or fence it, or declare it").
- **The sibling sweep is the cheap half of a masking fix.** M217 paid for the diagnosis and fixed one call;
  the identical call one line above cost this milestone a 25-second unreadable abort four releases later.
  `platform-alignment.md` §5 rule 9.
- **A skip is a hole in the evidence.** Two guards over the exact files this iter changed were not running
  — one RED since iter-02, one skipping for want of a binary — and both read as green. §5 rule 8.
- **My own probes failed twice more, and the discipline caught both** (a stub echoing `$3` instead of the
  URL, caught by its positive control; a `grep -c` reading 0 on a file that contained the string, caught by
  re-measuring with `-F`). Three iters running, the milestone's own instrumentation is its most reliable
  source of the defect class it studies.

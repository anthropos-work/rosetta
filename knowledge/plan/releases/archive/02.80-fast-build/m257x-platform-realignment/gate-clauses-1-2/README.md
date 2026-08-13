# M257x exit-gate clauses 1 & 2 — re-proof against platform origin HEAD `0c91421`

**Lane B evidence.** Run on the **dev host** (this Mac — `D-v28-15`: dev/test is local to the new Mac;
`odysseus` retired, `billion` demo-deployment only).

The gate's own words (`overview.md` frontmatter, quoted verbatim):

> Against platform @ **origin HEAD** (never a pinned pre-drift commit): (1) a cold `demo-down --purge` +
> `demo-up` on the **dev host** (D-v28-15: local to the new Mac; odysseus retired, billion demo-only)
> reaches `autoverify green:true / 0 warnings` across **3 consecutive cycles**; (2) the **full Playthrough
> suite passes on that stack** (30 live / 0 failing / 0 error) — presence AND function, so a green bring-up
> cannot mean an empty world.

Both clauses were last proven at platform **`2adcf71`**. Re-derived here, the drift since is **6 commits**
and **281 changed lines of `docker-compose.yml`** (40 insertions, 241 deletions) plus 29 changed lines of
`repos.yml`:

```
$ git -C stack-demo/platform log --oneline 2adcf71..0c91421
0c91421 Merge pull request #26 from anthropos-work/chore/drop-support-service-containers
838d907 chore(compose): drop the storage, messenger and customerio-sync containers
0dab54d chore(compose): run without the standalone storage; rename graphql -> core
ef32d4c Merge pull request #24 from anthropos-work/chore/prune-merged-services
6060315 fix(compose): give postgres a 120s healthcheck start_period
d11a403 chore(compose): drop roadrunner, prune dead env, repoint messenger

$ git -C stack-demo/platform diff --stat 2adcf71 0c91421 -- docker-compose.yml
 docker-compose.yml | 281 ++++++++---------------------------------------------
```

By the gate's own words those proofs are **void**; this directory re-establishes them.

### the `$HOME/.aws/credentials` bind — and a citation correction

The hazard: compose binds `$HOME/.aws/credentials` into a container. On a host where that path does not
exist Docker auto-creates it as an empty **directory**, `aws-sdk-go-v2` opens it successfully (opening a
directory succeeds) and then fails `EISDIR`; the cobra root sets neither `SilenceUsage` nor
`SilenceErrors`, so the container prints its full usage block and exits 1. rext mitigates it in the
injection override.

Re-derived, service that OWNS the bind, per platform commit:

| platform commit | owning service |
|---|---|
| `2adcf71` (last proof point) | `jobsimulation` |
| **`d11a403`** | **`backend`** |
| `6060315`, `ef32d4c`, `0dab54d`, `838d907`, `0c91421` | `backend` |

So the move landed at **`d11a403`**, the *first* of the six drift commits — not at `838d907`, which is
what rext commit `7844e97`'s own message says ("`838d907` moved the identical bind onto **backend**").
The fix that commit ships is correct; **its citation is not**. Recorded here rather than corrected in
history.

The mitigation is present in the pinned tooling: `demo-1`'s generated override carries
`backend: volumes: !reset null` (measured in
`stack-demo/rosetta-extensions/demo-stack/stacks/demo-1/docker-compose.injected.yml`), which is
`7844e97`'s derived-property keying doing its job. On this particular host the hazard could not have
fired anyway — `$HOME/.aws/credentials` exists as an empty **file**, not a directory — so this box is not
a valid negative control for it.

---

## 1. The refs proven against

### platform — verified `== origin HEAD`

```
$ git -C stack-demo/platform ls-remote origin HEAD
0c91421dfdb08dc75f17f1aabfb61394070e770b	HEAD
$ git -C stack-demo/platform rev-parse HEAD
0c91421dfdb08dc75f17f1aabfb61394070e770b
```

`0c91421` = *"Merge pull request #26 from anthropos-work/chore/drop-support-service-containers"*, whose
payload commit is `838d907`.

`repos.yml` @ `0c91421` lists **four** repos — `app`, `sentinel`, `next-web-app`, `studio-desk`.
`storage` and `messenger` are gone from it.

### every clone's ref (measured 2026-08-06, all advanced to their own `origin/main`)

| repo | sha | tag at HEAD | note |
|---|---|---|---|
| `platform` | `0c91421dfdb08dc75f17f1aabfb61394070e770b` | — | **== origin HEAD** |
| `app` | `ad9f3c498e9c244187440562f83c11e5408d6554` | — (`v1.369.0-7-gad9f3c498`) | was pinned `v1.366.0` = `b948604f`, **98 commits behind** |
| `sentinel` | `f2c461903de022a6a506a3a10355dbf503515ce5` | `v0.24.2` | was `88bc559`, 2 behind |
| `next-web-app` | `8297c684caacefb84ae2bcdbf0135795268d6341` | — | was `bb3313bc` `v2.133.0`, 41 behind |
| `studio-desk` | `41ee3575ddd94930148706fff05e18aa805cc19a` | — | was `14a5442` `v0.152.4`, 2 behind |
| `ant-academy` | `22df69dd81f1d718ecc9c088bbf96b6ae681c3a2` | — | was `9c3843cd` `v2.34.2`, 5 behind |

**Why `app` had to move too, and why leaving it at the old pin would have been a false proof.**
The canonical clone pin named `app` at `v1.366.0` (`b948604f`). At that sha `app` still *reads*
`STORAGE_RPC_ADDR` in three places (`stack-core/platform_predicate_guard.py:704` records the measurement)
— and platform `0c91421`'s compose sets it **nowhere**, because the `storage` container it addressed no
longer exists. Proving a bring-up green with `platform@0c91421 + app@v1.366.0` would have proved a
combination that exists in no repository. The whole live clone set is therefore at its own `origin/main`.

The five frozen legacy repos (`cms`, `jobsimulation`, `storage`, `messenger`, `roadrunner`) keep their old
pins: `repos.yml` @ `0c91421` no longer lists them, `make init` does not clone them, and the pins remain
useful only for reading pre-merge source by hand.

### the refs were FROZEN for the whole run — proven from the reflog, not asserted

`§5 rule 41a`. Every clone was advanced **once**, before cycle 1, in a single pass at **11:18:35 CEST**.
After all five cycles and seven suite runs, every clone's most recent reflog entry is still that one:

```
$ for r in platform app sentinel next-web-app studio-desk ant-academy; do
    git -C stack-demo/$r reflog --date=iso -n1; done
0c91421  HEAD@{2026-08-06 11:18:35 +0200}: reset: moving to origin/main
ad9f3c4  HEAD@{2026-08-06 11:18:35 +0200}: reset: moving to origin/main
f2c4619  HEAD@{2026-08-06 11:18:35 +0200}: reset: moving to origin/main
8297c68  HEAD@{2026-08-06 11:18:35 +0200}: reset: moving to origin/main
41ee357  HEAD@{2026-08-06 11:18:35 +0200}: reset: moving to origin/main
22df69d  HEAD@{2026-08-06 11:18:35 +0200}: reset: moving to origin/main
```

**All five cycles are the same subject.** No working tree moved between them.

**Disclosure — `/demo-up` fetches, and it cannot be asked not to.** `ensure-clones.sh`'s freshness
assertion runs `git fetch` per clone on **every** bring-up (that is where `clones.lock.json`'s `fetch_ok` /
`behind` come from). Measured: `FETCH_HEAD` mtimes are `12:15:46`–`12:15:52`, i.e. cycle 5's bring-up.
**A fetch moves `refs/remotes/origin/*` and nothing else** — `DEMO_ADVANCE_CLONES` defaults to `0`
(no-advance), so no working tree is touched and no file on disk changes. The reflog above is the proof.

**And a fetch during this run caught the platform moving under us.** `next-web-app`'s `origin/main`
advanced **4 commits** past our frozen `8297c684` (to `f97ba659`) mid-run. Our clone stayed put, which is
what a same-refs claim requires — but it means **this dossier's proof is against `next-web-app@8297c684`,
which was `origin/main` when the run started and is no longer.** That is a live instance of the
milestone's own `re_scope_trigger` arithmetic, observed rather than predicted.

### what the 98-commit `app` advance actually contained

Decomposed from local objects (no network):

```
b948604f..2035f9a4 = 93 commits      # the old v1.366.0 pin → the sha CLAUDE.md cites as origin/main
2035f9a4..ad9f3c49 =  5 commits      # that sha → what we proved against
b948604f..ad9f3c49 = 98 commits
```

The 5-commit tail touches **no Go source at all** (`git diff --name-only 2035f9a4 ad9f3c49 | grep -c '\.go$'`
→ **0**) and `main.go` is **byte-identical at 1639 lines**.

**One correction to the figure circulating for that tail.** It is *not* "the entire residual was a label",
and `terraform/main.tf` is *not* byte-identical. `git diff --numstat 2035f9a4 ad9f3c49`:

```
411   0  .claude/skills/publish/SKILL.md
  1   0  CLAUDE.md
 17   1  knowledge/deployment.md
  1   1  terraform/main.tf
 37  12  terraform/variables.tf
```

`terraform/main.tf` is 786 lines at both ends, but one line **changed**: the `error_message` string of the
`atlas_migration.sentinel_migration` lifecycle precondition was rewritten (a prose warning about
`atlas_sentinel_dev_url` being a disposable scratch database). The *conclusion* — no cited terraform
**construct** moved — holds. The *evidence* "byte-identical" does not, and `terraform/variables.tf` really
did change by 49 lines. Stated here because this figure is being handed to a blind reading as ground truth.

---

## 2. The rosetta-extensions tag

**One tag, cut by the iteration lane, verified here — not a second one.** (`D-M257x-101-3`.)

```
$ git ls-remote --tags origin refs/tags/fast-build-m257x-iter-101
0011c10aba0ff0950341cb410265ee59d070afe3	refs/tags/fast-build-m257x-iter-101
09d06070fd99c742d7a671c468abf93074278575	refs/tags/fast-build-m257x-iter-101^{}
```

`0011c10a` is the **annotated tag object**; it dereferences to commit **`09d06070`**. Both forms are on
origin, so a stack cloning the tag from origin gets `09d06070`.

**Contains the fix the pin exists to deliver:**

```
$ git merge-base --is-ancestor 7844e97 'fast-build-m257x-iter-101^{commit}' && echo YES
YES
$ git log --oneline -1 7844e97
7844e97 fix(stack-injection): key the volumes reset on a DERIVED property, not on a deleted service name
```

**Pin + clone:**

| | value |
|---|---|
| `.agentspace/rext.tag` | `fast-build-m257x-iter-101` |
| `stack-demo/rosetta-extensions` | checked out at that tag |

**rext `main` has moved past the tag, by one lane-B commit.** `4cb920a`
(*"pin(M257x/lane-B): advance the canonical clone pin to platform origin HEAD 0c91421"*) changes exactly
one file, `demo-stack/clones.pin.json`, and is **not** in `fast-build-m257x-iter-101`.

**It is not needed for clauses 1 or 2, and that is measured, not assumed.** `ensure-clones.sh` seeds the
canonical pin into the ephemeral workspace **copy-if-ABSENT** and never clobbers an existing one; the
workspace copy `stack-demo/clones.pin.json` was written by hand with the same advanced values before the
first cycle, so the canonical file is never read on these runs. Its only two consumers are the
**non-fatal** freshness assertion (which reported `all clones provably fresh-or-pinned (0 stale-by-neglect
/ pin-drift / fetch-unknown)`) and the opt-in `DEMO_ADVANCE_CLONES=pinned` path (not used — the default is
no-advance). **Recommendation, not an action taken:** fold `4cb920a` into the next tag the iteration lane
cuts, so a fresh box reproduces this clone set without hand-authoring.

**Two tags were cut by this lane before the scope message arrived and have been DELETED from origin and
locally:** `fast-build-m257x-iter-101b` (byte-identical content to `-101`) and `fast-build-m257x-iter-101c`
(`-101` + `4cb920a`). `git ls-remote --tags origin 'refs/tags/fast-build-m257x-iter-101*'` now returns
`fast-build-m257x-iter-101` and its `^{}` peel and nothing else.

---

## 3. Cycles — clause 1

Each cycle is `rosetta-demo down 1 --purge` followed by `up-injected.sh 1` on the dev host (this Mac).
`--purge` removes the stack's data dir **and its images**, so every cycle re-builds all six images; the
BuildKit cache is deliberately left alone (`-af` would make the next cycle truly-cold and is never used).

| # | purge | bring-up | `up-injected.sh` exit | `autoverify.json` | warnings | ts |
|---|---|---|---|---|---|---|
| **1** | — (pre-existing stack) | **621 s** | 0 | `green:true` | **0** | `2026-08-06T09:31:05Z` |
| **2** | 8 s | **402 s** | 0 | `green:true` | **0** | `2026-08-06T09:51:21Z` |
| **3** | 12 s | **370 s** | 0 | `green:true` | **0** | `2026-08-06T09:57:55Z` |
| 4 | 9 s | 370 s | 0 | `green:true` | 0 | `2026-08-06T10:10:11Z` |
| 5 | ~9 s | 371 s | 0 | `green:true` | 0 | `2026-08-06T10:21:56Z` |

Cycles **1, 2, 3** are the three consecutive the clause asks for; 4 and 5 were run to chase the clause-2
defect below and are recorded because they are two more consecutive greens, not because the clause needs
them. **Five for five, no restart of the count, no defect in the bring-up path.**

Cycle 1 is longer because it is the only one whose images were not built from a warm BuildKit cache for
this clone set; cycles 2–5 land in a tight **370–402 s** band.

Raw evidence per cycle in `raw/`: `cycleN-up.log`, `cycleN.rc`, `cycleN-autoverify.json`,
`cycleN-autoverify.log`.

### the bring-up handled the platform drift correctly — measured, not assumed

From `raw/cycle1-up.log`:

- `note: injection candidate 'cms' is no longer built by the platform compose — skipping (folded into app)`
- `note: injection candidate 'jobsimulation' is no longer built by the platform compose — skipping (folded into app)`
- `injecting: app (derived from the platform compose's build set: sentinel app)`
- `freshness: all clones provably fresh-or-pinned (0 stale-by-neglect / pin-drift / fetch-unknown)`
- the generated override carries `backend: volumes: !reset null` — `7844e97`'s derived-property keying
- the stack starts **11** containers: `backend`, `gotenberg`, `postgresql`, `redis`, `sentinel`,
  `directus`, `fake-fapi`, `fake-bapi`, `next-web-app`, `hiring-app`, `studio-desk`. **No `storage`,
  `messenger` or `customerio-sync`** — matching compose @ `0c91421`.

Every demo-patch **applied**. Sixteen reported `whole-file sha DRIFTED … but the anchor is intact (1x)` and
self-healed — the expected consequence of `app` advancing 98 commits and `next-web-app` 41. That is
`demopatch-spec.md`'s design working (*"the anchor is the contract; the whole-file sha is only a
baseline"*), not a defect: **0 patches refused, 0 skipped**. Re-recording the 16 stale baselines is
optional hygiene and was deliberately **not** done — it would churn on every platform advance, and the
freshness gate is self-healing by design.

---

## 4. Playthrough suite — clause 2

**The gate number is reached: `passing=30, failing=0, unimplementable-without-platform-edit=0` over 31
manifest use cases**, from the runner's own report JSON (`report/last-binding-report.json`), never grepped
stdout, with a `binding:true / scoped:false` provenance sidecar and `playwright_exit:0 / ptreport_exit:0`.

```json
{"failing": 0, "passing": 30, "unimplementable-without-platform-edit": 0, "unimplemented": 1}
```

The 1 `unimplemented` is the declared `will-not-build` **verdict** (the onboarding self-import journey,
`onboarding.enterprise-workforce-standard.UC1`) — the milestone's own accounting, not a gap.

### but it is intermittent, and the intermittency is characterised

Seven suite runs. The unit of measurement is a **full, unscoped run** — `run-playthroughs.sh 1 --reset` —
because that is the only shape whose gate is binding.

| run | stack | shape | result |
|---|---|---|---|
| **A** | cycle-1 | full `--reset`, first-ever suite run on that stack | **29/1** — `pt-assignment-assign` FAILED |
| B | cycle-1 | full, **no** `--reset` | 26/4 — **invalid by design**, see below |
| **C** | cycle-1 | full `--reset`, third run on that stack | **30/0/0**, exit 0 |
| **D** | cycle-3 | full `--reset`, first-ever suite run on that stack | **29/1** — same spec, same signature |
| E | cycle-4 | full `--reset`, preceded by a standalone probe that loaded the assign surface | 30/0/0 |
| F | cycle-5 | full `--reset`, with an in-suite probe loading the assign surface 8.6 s earlier | 30/0/0 |
| **G** | cycle-5 | full `--reset`, clean (probe removed) | **30/0/0**, exit 0, 140 s |

**Run B is kept as a negative control, not as a failure.** A full run *without* `--reset` cannot pass: four
mutating Playthroughs (`onboarding.completion`, `onboarding.individual`,
`onboarding.enterprise-workforce-ai-readiness`, `skill-paths.legacy`) require the reset-to-seed pre-state
and their negative controls correctly refuse a world that has already been played. `playthroughs.md`'s
*"additive re-seed FORBIDDEN"* is exactly this. **`--reset` is not optional for a binding measurement.**

---

## 5. Defects found

### DEF-1 — `pt-assignment-assign` fails on the first load of the assign surface (2 of 2), and the harness's own diagnosis of it is refuted

**Signature, identical both times:**

```
Error: the assignable-affordance count drops by exactly one (the assignment landed + was read back)
  Expected: 15      ← before - 1, so the baseline `before` was read as 16
  Received: 14      ← the settled post-write count
  Timeout 20000ms exceeded while waiting on the predicate
```

**The write always landed.** After each failing run the DB holds exactly **one** new row — never two:

```
$ docker exec demo-1-postgresql-1 psql -U postgres -d postgres -tAc \
    "select count(*) from public.organization_assignments
     where resource_type='skill_path' and created_at > now() - interval '20 minutes';"
1
```

So this is **not** a double-write and **not** a product defect. The **baseline is over-read by exactly
one**; the delta assertion is then unreachable and a correct write reports RED.

**The harness's stated cause is measured to be wrong.** `baseline-settle-fence.unit.spec.ts` (M257x harden
pass 1) attributes this exact flake to *"the members table is still filling"* and mandates a
`waitForMembersTableSettled()` before any count baseline feeding a strict delta. That call **is** in place
in `assignment-assign.spec.ts`. A 60-sample instrumented probe of the same surface, through the same page
object, contradicts the premise — on **both** a warm stack and a genuinely cold one (fresh `demo-up` +
`--reset-only`, first-ever load):

```
t=0.0s rows=20 assignable=0
t=0.5s rows=20 assignable=15
t=1.0s rows=20 assignable=15
…  (unchanged through t=30.6s)
```

The count reaches its settled value in **under 0.5 s** and never shows 16 for 30 s. **The table is not
still filling.** `waitForMembersTableSettled`'s criterion (two equal reads 1 s apart) is not what is
failing, so strengthening it would not fix this.

**What the seven runs do isolate:** the spec fails **iff it is the first load of
`/enterprise/assignments/skill-paths` inside a full-suite run on a freshly built stack** (A, D — 2 of 2).
It passes whenever *any* earlier load of that surface has happened in the same session (C, E, F, G — 4 of
4), including when the earlier load is a throwaway probe **8.6 s** before it in the same run. A standalone
first load outside a suite run reads 15 correctly — so the trigger involves the ~17 specs that precede it,
which log in as several different heroes through the Clerkenstein seat-switch. **Cross-Playthrough session
or org-context bleed is the surviving hypothesis; it is not yet proven and is deliberately not patched.**

**Not fixed here, and that is a decision, not an omission.** The milestone's own doctrine is *repair by
CLAIM, not by FILE* — the previous repair of this same symptom was built on a premise the measurement above
refutes, and a second guess would repeat that. It also needs a rext tag, which this lane was directed not
to cut. Handed back with the mechanism narrowed from "flaky" to a reproducible precondition.

### DEF-2 — `stack-demo/autoverify.json` was a stale orphan that reads as a gate RED

`stack-demo/autoverify.json` held `{"warnings":2,"green":false,"ts":"2026-07-31T20:37:55Z"}` — six days
old. **Nothing writes it and nothing removes it.** The live verdict every grader is supposed to read is
`demo-stack/stacks/demo-N/autoverify.json`; `up-injected.sh:87` and `rosetta-demo:283` both `rm -f` only
`$STACK/autoverify.json`, and `$STACK` is the per-N stack dir.

A grader looking in the obvious place — the stack workspace root — would have read a month-old **RED** for
a stack that is green. Quarantined to `raw/ORPHAN-stack-demo-root-autoverify.{json,log}` so nobody reads
it. **Left as a finding for the tooling owner:** the two paths should not both be plausible.

### DEF-3 — `7844e97`'s commit message cites the wrong platform commit

Recorded in §1. The fix is correct; the citation names `838d907` where the measurement says `d11a403`.

---

## 6. Verdicts

### Clause 1 — **MET at platform `0c91421`**

Three consecutive cold `demo-down --purge` + `demo-up` cycles on the dev host reached
`autoverify green:true` with **0 warnings**: `09:31:05Z`, `09:51:21Z`, `09:57:55Z`. Two further cycles
(`10:10:11Z`, `10:21:56Z`) were also green — **five consecutive, zero restarts of the count**. Every cycle
ran against the same frozen clone set, proven from the reflog in §1.

### Clause 2 — **MET at platform `0c91421`, with a disclosed intermittency**

The suite's own binding report reads `passing=30, failing=0, unimplementable=0` over 31 manifest use cases,
with `playwright_exit:0 / ptreport_exit:0` and a `binding:true, scoped:false` provenance sidecar —
reproduced on two different stacks (runs C and G).

**The honest qualifier, which must travel with the number:** on a freshly built stack the first full run
failed 29/1 in **2 of 2** attempts (DEF-1). The gate's number is real and reproducible; it is **not yet
reproducible on the first attempt against a cold stack**. Anyone re-running this for the gate should expect
to need a second run, and should treat DEF-1 as open.

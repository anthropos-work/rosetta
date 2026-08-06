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

## 3. Cycles

<!-- filled in below -->

---

## 4. Playthrough suite

<!-- filled in below -->

---

## 5. Defects found

<!-- filled in below -->

---

## 6. Verdicts

<!-- filled in below -->

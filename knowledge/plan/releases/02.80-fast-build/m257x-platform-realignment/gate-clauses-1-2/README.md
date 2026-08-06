# M257x exit-gate clauses 1 & 2 — re-proof against platform origin HEAD `0c91421`

**Lane B evidence.** Run on the **dev host** (this Mac — `D-v28-15`: dev/test is local to the new Mac;
`odysseus` retired, `billion` demo-deployment only).

The gate's own words (`overview.md` frontmatter, quoted verbatim):

> Against platform @ **origin HEAD** (never a pinned pre-drift commit): (1) a cold `demo-down --purge` +
> `demo-up` on the **dev host** (D-v28-15: local to the new Mac; odysseus retired, billion demo-only)
> reaches `autoverify green:true / 0 warnings` across **3 consecutive cycles**; (2) the **full Playthrough
> suite passes on that stack** (30 live / 0 failing / 0 error) — presence AND function, so a green bring-up
> cannot mean an empty world.

Both clauses were last proven at platform `0dab54d`. `838d907` then rewrote `docker-compose.yml` by 107
lines, deleted the `storage`, `messenger` and `customerio-sync` compose services, and removed two
`repos.yml` entries. By the gate's own words those proofs are **void**; this directory re-establishes them.

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

<!-- filled in below -->

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

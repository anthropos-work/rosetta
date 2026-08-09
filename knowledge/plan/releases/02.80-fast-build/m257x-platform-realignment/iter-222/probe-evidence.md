# iter-222 probe evidence — sealed 2026-08-09T17:45:46Z

## repos.yml @ platform origin/main
0c91421dfdb08dc75f17f1aabfb61394070e770b
-  app
-  sentinel
-  next-web-app
-  studio-desk

## canonical pin keys (rosetta-extensions/demo-stack/clones.pin.json)
count: 11
 - ant-academy 22df69dd8
 - app ad9f3c498
 - cms 93e6aa354
 - jobsimulation 5d3003f9f
 - messenger d41029217
 - next-web-app 8297c684c
 - platform 0c91421df
 - roadrunner 87d8d4438
 - sentinel f2c461903
 - storage 769660542
 - studio-desk 41ee3575d

## freshness vs origin/main (fetched this session: platform app sentinel next-web-app studio-desk ant-academy)
- platform pin=0c91421df origin/main=0c91421 behind=0
- app pin=ad9f3c498 origin/main=3eaadae68 behind=28
- sentinel pin=f2c461903 origin/main=f2c4619 behind=0
- next-web-app pin=8297c684c origin/main=19423a1fb behind=12
- studio-desk pin=41ee3575d origin/main=41ee3575 behind=0
- ant-academy pin=22df69dd8 origin/main=c885dab2 behind=9

## POST-REPAIR — and the finding a plain `git fetch` disclosed

Nobody had fetched. `anchor_construct_guard` and `repair_postcondition` resolve corpus anchors at
the app clone's `origin/main` — which pointed at `ad9f3c498` (= the clone's own HEAD) until this
iter fetched it forward to `3eaadae68`. Both guards were green because the reference was stale.
Captured 2026-08-09T17:53:03Z, app origin/main = 3eaadae68:

```
anchor-construct-guard: RED — 9 anchor(s) resolve to a non-construct:

corpus/architecture/platform-migration-status.md:94  [anchor-on-blank-line]
    cites  : app/main.go:1471  (read at origin/main@3eaadae)
    which is: (blank)

corpus/architecture/platform-migration-status.md:94  [anchor-on-closing-delimiter]
    cites  : app/main.go:1552  (read at origin/main@3eaadae)
    which is: }

corpus/architecture/security_compliance.md:213  [anchor-on-blank-line]
    cites  : backend.go:289  (read at origin/main@3eaadae)
    which is: (blank)

corpus/architecture/security_compliance.md:214  [anchor-on-blank-line]
    cites  : backend.go:301  (read at origin/main@3eaadae)
    which is: (blank)

corpus/architecture/security_compliance.md:215  [anchor-on-blank-line]
    cites  : backend.go:295  (read at origin/main@3eaadae)
    which is: (blank)

corpus/architecture/security_compliance.md:262  [anchor-on-closing-delimiter]
    cites  : backend.go:315  (read at origin/main@3eaadae)
    which is: }

corpus/architecture/security_compliance.md:264  [anchor-on-blank-line]
    cites  : backend.go:301  (read at origin/main@3eaadae)
    which is: (blank)

corpus/ops/observability.md:28  [anchor-on-closing-delimiter]
    cites  : app/main.go:277  (read at origin/main@3eaadae)
    which is: },

corpus/services/ai-readiness.md:483  [anchor-on-closing-delimiter]
    cites  : readiness.go:710  (read at origin/main@3eaadae)
    which is: }

```

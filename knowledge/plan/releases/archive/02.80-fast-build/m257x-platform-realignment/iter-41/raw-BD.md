# Auditor B — 7 files / 1541 lines — 1 blocker, 15 minors
## B-B1 `ai-readiness.md:37-43` [EDITED] — states LANDED work as OUTSTANDING
"The patch **must re-anchor** — this is the M246 drift-ledger D-07 item, owned by **v2.7 M250**."
The re-anchor ALREADY LANDED at v2.7 **M254**. The manifest
`demo-stack/patches/app-aireadiness-snapshot-loadmembers/*.yaml:42` reads
`path: internal/aireadiness/readiness.go` and its header :33 says "v2.7 M254 RE-POINT".
The same doc contradicts itself at :458 (present tense, patch bounding the read).
NOTE scope caveat: ground truth is the rext authoring copy, not platform HEAD.

# Auditor D — 7 files / 1472 lines — 2 blockers, 16 minors
## B-D1 `sentinel.md:12` [EDITED] — "Language: Go 1.25"
Sentinel is Go **1.26**: `sentinel/go.mod:3` = `go 1.26.0`; `Dockerfile:2`/`Dockerfile.dev:2` =
`golang:1.26-bookworm`; `sentinel/CLAUDE.md:9` = "Go 1.26". ACTIONABLE: the same doc's :115 says
`go run main.go`, so a reader who provisions 1.25 gets a hard `go.mod requires go >= 1.26.0` failure.

## B-D2 `sentinel.md:22` [EDITED] — "256 CPU / 256 MB on ECS"
`sentinel/terraform/locals.tf:4-5` = `service_cpu = 256`, `service_memory = **128**`. CPU right,
memory wrong by 2x — a conjunction false while one conjunct measures correctly (§5 rule 17's shape).

## D's notable CLEAN results (high-signal, these were repaired sections):
- sentinel's "there is NO `manager` role" verified THREE ways incl. a live `p_type='g2'` query -> admin/member/candidate only. (iter-39 blocker #22's fix HOLDS.)
- shared_libraries: all six main.go Connect-handler anchors exact + no seventh; standalone `authn` imported by NOTHING across all go.mod/go.sum/*.go. (iter-39 #29's fix HOLDS.)
- studio-room: `gen.py:484-492` = exactly nine add_argument; zero consumers of a request-level `template`. (iter-39 #33's fix HOLDS.)
- taxonomy 22,470 / 42,790 reproduce exactly live.

# iter-195 — decisions

## `D-M257x-195-1` — they were never unreadable; nobody had typed `go test`

The exclusion reason on all six sections is *"Go module — N `*_test.go`; no Python runner collects
it."* Every word true. For **nine iters** it was read as *unreadable*, and the route stayed open while
every whole-population figure inherited the gap.

Feasibility probe on the smallest section (`stack-secrets`, 20 files) — **4 packages, all ok, 2.7 s,
offline, no credentials.** The gap was never capability.

**First reading of the Go population** (`go1.26.5 darwin/arm64`, `GOFLAGS=-mod=mod`, `-count=1`):

| section | pass | fail | skip |
|---|---:|---:|---:|
| alignment | 126 | 0 | 0 |
| clerkenstein | 416 | 0 | 0 |
| playthroughs | 183 | 0 | 0 |
| stack-secrets | 195 | 0 | 0 |
| stack-seeding | 1,276 | 0 | 0 |
| stack-snapshot | 518 | 0 | 0 |
| **TOTAL** | **2,714** | **0** | **0** |

**Name the unit** (`§5`, iter-177). iter-186's **264** counts `*_test.go` **FILES**; this **2,714**
counts test **FUNCTIONS**. Same population, different measurements, and the ratio between them is not a
constant anyone may lean on. The file count was independently re-derived here and came to **264
exactly**, which corroborates iter-186's figure rather than replacing it.

## `D-M257x-195-2` — the declaration is now a derivation

`LANGUAGE_EXCLUDED_SECTIONS` is hand-written and its six reasons all say *"Go module"*. `go_sections()`
reads `go.mod` off disk instead, and a fence asserts the two agree **in both directions** — so the
registry can no longer name a section that stopped being a Go module, nor miss one that started.

Measured: the derived set equals the declared set exactly. And with both languages read, the sections
now **partition**: `SECTIONS` (5, Python) ∪ `go_sections` (6, Go) = `all_sections` (11), no overlap, no
remainder — asserted, not assumed.

## `D-M257x-195-3` — a section that FAILS TO BUILD emitted a silent zero, and I wrote it

`go test -json` on a section that does not compile emits **no test events**. The first cut of
`go_census` therefore tallied `0 pass / 0 fail / 0 skip` and marked it runnable — **indistinguishable
from a section with no tests, and it would sum into a clean total.**

Found while writing the function's own tests, not by review. Repaired: a non-zero exit with zero
observed test events is `unrunnable: 1`. Proved in both directions — a synthetic non-compiling section
is flagged, a synthetic healthy one is not.

This is iter-191's *false CANNOT-RUN* in the mirror. There the guard refused when it could have
answered; here the census would have **answered when it could not run**, and answered *green*. Of the
two, this direction is worse: refusing is loud.

## `D-M257x-195-4` — reading Go closes one language and leaves one, so the remainder is NAMED

TypeScript is now the only unread population: **45 Playwright specs in `playthroughs`** plus the **30**
inside `stack-verify` that iter-187 found. The `--go` report prints that line every time, and a fence
asserts `playthroughs`'s reason still says *TypeScript* — because reading its Go half must not silently
promote a **mixed-toolchain** section to fully-read. That is iter-186's rule at the exact point it would
be easiest to forget: the moment a long-open gap closes is when a remaining one becomes invisible.

**`--go` is off by default** — it shells out to `go` — but the scope line now says the sections *are*
runnable and quotes the reading, so the default can never again be mistaken for a limit.

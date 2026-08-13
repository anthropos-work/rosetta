# iter-69 — decisions

## `D-M257x-69-1` — a ref-pinned citation is a MEASUREMENT; the defect class is the UNPINNED half

`FIX-M257x-iter63-app-citation-residual` scope B2 was routed as **"64 unrepaired non-mainline
citations."** Graded at the ref the exit gate names (`app` `origin/main` `9d00a313` v1.367.0), the
126 distinct (corpus-file × citation) pairs partition:

| class | n | defect? |
|---|---|---|
| identical at `b948604` **and** `origin/main` | 62 | no |
| drifted, but the corpus **block names its ref** | **59** | **no — a measurement** |
| drifted and **UNPINNED** | **2** | **yes** |
| file absent at the ref | 3 | yes, but mis-rooted rather than dead |

**The residual was 5, not 64.** This is not a re-cut of anything: it is TOK-05's own thesis arriving
from the other side. The predicate under a citation is not *"this line names this construct"* — it is
**"this line names this construct AT THE REF THIS BLOCK CLAIMS"**, and a corpus that writes its refs
has already discharged 59 instances of it. §5 rule 33 (*a ref-pin is a DATE, not an exemption*)
states the same rule for count/migration/RPC claims; the same `_REF_PINNED` regex decides both, so
the two instruments cannot disagree about what a pin is.

Verified rather than assumed: **every ref-pinned mainline citation in the B2 files resolves exactly
at its own pin.** `backend.md:39`'s seven — `main.go:1185-1228`, `:1187`, `:1188`, `:1196`, `:1204`,
`:1212-1214`, `:1228` — each land on the construct the sentence names at `b948604`, and each land on
an unrelated `web.NewServer` argument at `origin/main`. The corpus is internally honest; only a
reader who ignores the pin is misled.

**Corollary, and the reason this is a decision and not an observation:** a pass that "repairs" the
59 would be **inducing** drift, not removing it — it would move 59 correct claims onto a ref that
moves again next week, and grow the class by every line number it rewrites (§5 rule 34).

## `D-M257x-69-2` — the screen's REACH, stated, because it hid two of the three real defects

The mechanical screen (`iter69_screen.py`) grades one shape: *a citation whose line CONTENT changed
between the ref the corpus was written against and the adjudication ref.* Two inverted mutants moved
its verdict (pin-blind → 61 candidates; comparison-blind → 0) and a no-op control survived, so it
discriminates. It is still **blind in three named directions**, each of which hid a live defect
found by READING, not by the screen:

1. **Bare `:N` continuations after a PATHLESS antecedent.** iter-68's enumerator sets its
   `last_path` only when a full `path:line` matched. A line that says *"`app/main.go` registers six
   Connect handlers (Users `:1178`, …)"* names the antecedent **in backticks, one clause away, with
   no line number to latch onto** — and contributes **zero** citations. Derived: **23 citations
   across 14 lines** are invisible to every count this class has ever reported.
   `shared_libraries.md:70` carries **six of them, five of which were wrong at every ref.**
2. **A path written without its clone root.** `anchor_construct_guard` reports
   `unresolvable head 'internal' x6` — six citations spelled `internal/…` that belong to `app` and
   are silently dropped rather than resolved. That is why the Judge0 defect below survived.
3. **Platform-rooted citations.** The screen's universe is the `app` clone only, so
   `docker-compose.yml:118` was never looked at.

And the same enumerator rule is **too greedy in the other direction**: `ai_architecture.md:96`'s bare
`` `:15-17` `` refers to **that document's own lines 15-17** (the ⚠️ retraction block), and the
inherit-the-last-path rule attributed it to `app/internal/skillerai/ai.go`. One rule, both a false
negative and a false positive, in prose versus in a table cell.

## `D-M257x-69-3` — the three defects, adjudicated against platform artifacts

- **`shared_libraries.md:70`** — *"`app/main.go` registers six Connect handlers (Users `:1178`,
  Organizations `:1179`, Skiller `:1187`, JobSimulation `:1195`, CMS `:1204`, LabSession `:1218`)"*.
  Measured at `b948604`: the mux is **1187 / 1188 / 1196 / 1204 / 1213 / 1228**. **Five of six wrong,
  unpinned, and contradicting `backend.md:39`** — which is pinned and correct. Repaired to the
  measured set, pinned, with the CMS handler's conditionality (`if cmsRPCServer != nil`) restored.
- **`shared_libraries.md:79`** — two defects on one line. Judge0 is called at
  `internal/jobsimwiring/wiring.go:118` — true at `b948604`, **`:123` at `origin/main`**, and the
  citation carried no pin. And `ROADRUNNER_RPC_ADDR` *"(`docker-compose.yml:118`)"* asserts the
  variable is in the compose; at platform `0dab54d` it is **in the compose zero times**, and
  `docker-compose.yml:118` is an unrelated `AWS_REGION` line. Both repaired; the path re-rooted to
  `app/internal/…` so the anchor guard can reach it.
- **`platform-alignment.md:872`** — rule **32**'s own worked example
  (`cmd/academyImport/main.go:235` / `:231`) is unpinned, and the storage fold has deleted **both**:
  at `9d00a313` that file names `STORAGE_RPC_ADDR` nowhere. Pinned — **rule 33 applied to rule 32's
  neighbour**, which is where it was needed and not written.
- **`external_services.md` ×3** — `app/studio/services/ai.py`, `app/studio/gen.py`. Not dead: the
  path is the **in-image** path, from `anthropos-studio-room`, which CI pulls into the `app` image as
  an `additional_repo`. It is in **no `app` commit at any ref**. One scoping note added where the
  paths first appear; the claims themselves were correct.

**The repair widened the fence, measurably:** `anchor_construct_guard`'s citations adjudicated at
`origin/main@9d00a31` went **43 → 49** — re-rooting `internal/…` to `app/internal/…` made six
citations reachable that the guard had been dropping as unresolvable. That is the TOK-05 move
landing in the guard's own reach line, not in prose.

## `D-M257x-69-4` — G8, the per-service profile bullet (`FENCE-M257x-iter68-profile-bullet`)

Eight service docs open with a `* **Profile**: …` bullet. iter-68 read them by hand and found
**seven of eight wrong**, all seven naming `graphql` — while every fence in the family read GREEN,
because the bullet is none of G1's three constructs: not a command, not a table cell, and the token
sits *inside* `profiles: [...]` rather than adjacent to the word "profile". **Seventh reach limit of
this milestone.**

**G8 is G7 inverted.** G7 reads a PROFILE and checks the services beside it; G8 reads a SERVICE —
**from the doc's own file stem, derived, so a new service doc is in scope the day it is written** —
and checks the profiles beside it. Three shapes, each decidable against `docker-compose.yml`:
`list` (the quoted `profiles: [...]` must equal the declaration, both directions), `no-service` (the
stem must indeed be absent from compose), `always-on` (the service must be in the always-on floor).
A bullet matching none is **UNREACHED**, never an empty claim.

Live: **8/8 reached, 0 unreached — `{list: 5, no-service: 2, always-on: 1}` — GREEN.** Green is the
right outcome one iteration after a hand repair and is **not** the evidence; the fixtures and the
mutants are.

**Shape `no-service` is the one worth having.** `jobsimulation.md` and `roadrunner.md` named a
profile for a service the platform had already deleted, and until now no assertion in this file
could contradict a claim about a service that does not exist.

**Every fixture is copied verbatim from the live corpus and mutated by exactly one fact.** That is
the harden-pass-16 rule and it is load-bearing here: G3's reach was **zero for three iterations**
because its fixtures used the bare `(default)` spelling its own regex wanted while the corpus writes
`*(default — PROFILE ?= core)*` — **the fixture agreed with the pattern instead of with the corpus**,
so the class was invisible and looked healthy.

Watched RED before trusted: **5 source mutants, all caught** (one-direction 2F · shape-blind 2F ·
stem-blind 4F · always-on-unchecked 1F · prose-graded-as-empty-claim 1F), plus an **ARTIFACT**
inversion in `TestMutants` — swap `storage` into the floor and `gotenberg` into `storage-legacy`,
and the two correct bullets must both become findings **naming the mutated truth**, which a G8 that
memorised the real topology would survive. **No-op control SURVIVED** (119 tests OK, unchanged).

## `D-M257x-69-5` — the instrument behind every count in this class crashes, and my own repair found it

Re-measuring at close, `iter68_cites.py` raised
`TypeError: '<' not supported between instances of 'NoneType' and 'int'`. Its `sorted({(path, a, b)})`
compares `b=None` (from `X:N`) with `b=int` (from `X:N-M`) whenever the **same path carries both
forms at the same start line** — which no corpus contained until this iter's repair wrote
`app/main.go:1212-1214` beside `app/main.go:1187`.

**Every count this class has ever reported — 86, 96, 104, 105, 123, 135 — came off an enumerator
that was correct by luck.** Fixed (explicit sort key); recorded because the lesson is not the crash,
it is that a scratch instrument accumulating six iterations of load-bearing numbers had never been
tested on a shape the corpus was one edit away from containing.

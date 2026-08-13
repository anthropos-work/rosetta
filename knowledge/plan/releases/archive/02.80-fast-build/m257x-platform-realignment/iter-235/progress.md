**Type:** tik — under `TOK-08` (census a mechanical class exhaustively).

# iter-235 — does every `make` target the corpus tells you to run still exist?

## Why this class, right after iter-234's refusal

iter-234 refused the containment census because a corpus **paraphrases** source. That refusal kills
instruments that must match free text; it says nothing about claims whose subject is an **enumerable set**.
A Makefile's targets are enumerable, and `make up` is not a paraphrase of anything — it is a command that
runs or fails. Same lesson the corpus already wrote down for compose profiles (*grade a documented command
on "does it still select anything"*), applied one layer up, where every guide opens.

## Population

Instrument: `.agentspace/scratch/work-m257x/make235.py`, over `corpus/**` + `CLAUDE.md` +
`.claude/skills/**`, graded against **every repo in the `stack-demo` clone set that has a Makefile**, read
from git rather than the working tree.

| repo | declared targets |
|---|---|
| `platform` | 15 |
| `app` | 10 |
| `cms` | 4 · `graphql-wundergraph` 3 · `jobsimulation` 2 · `messenger` 2 · `roadrunner` 2 · `sentinel` 2 · `storage` 2 |
| `ant-academy`, `next-web-app`, `studio-desk`, `rosetta-extensions` | no Makefile at HEAD |

**394 documented `make <target>` sites · 36 distinct targets.**

## The denominator was wrong first, and that is recorded

The first reading graded everything against `platform/Makefile` alone and reported **21 distinct / 61
sites** not resolving. That is a category error of the same family iter-234 named: **the corpus's `make`
invocations are repo-relative by design** (*"`cd cms && make init-studio`"*, *"per service"*), so the
denominator is the clone set, not one Makefile. Widened, the 61 classify cleanly:

| class | distinct | sites | what it is |
|---|---|---|---|
| resolves in `platform/Makefile` | 15 | **333** | the documented entry point, intact |
| **English false positive** | 11 | 17 | *"make opting in impossible"*, *"to make those believable"* — `make` as a verb. Instrument noise, reported not hidden |
| **repo-relative, real elsewhere in the clone set** | 9 | 43 | `gen` `setup` `test` `initdb` `migrations` `init-studio` `update-studio` `run` `updatesubg` — correct documentation |
| **nowhere in the clone set** | **1** | **1** | `make force-gen` |

## Verdict

**Zero documented `make` targets are provably dead.** The single non-resolver, `make force-gen`
(`shared_libraries.md:152`), is a **`proto`-repo** target — and `proto` is a Go module pulled at Docker
build via `GH_PAT`/`GOPRIVATE`, **never cloned**. That is **UNMEASURABLE from the clone set, not wrong** —
the same distinction iter-123 drew for `infrastructure`, and the reason it is not repaired.

**Reverse direction is a clean sweep: 0 of `platform`'s 15 targets are named by no document.** Perfect
discoverability — a fact this corpus has never had a number for.

> **Stated limit on the classifier.** It asks *"does this target exist in ANY cloned repo"*, not *"in the
> repo this fenced block addresses."* So `make gen` inside a `proto` block resolves via `app`'s Makefile
> by coincidence. The limit is disclosed rather than tuned away; it does not affect the verdict, because
> the only target that failed the weaker test also fails the stronger one.

## The class delivered a real defect — one layer over

Phase 5 said *repair any miss that sits in a runnable block*, so the same question was asked of the other
half of a runnable instruction: **the directory you are told to `cd` into.** 22 sites name an
archived-service directory; **13 are fenced and copy-pasteable.** Two repairs landed:

**1. `corpus/ops/quick_ops.md` "Apply all migrations" — 3 of 3 runnable lines failed.** The cookbook whose
entire purpose is copy-pasteable one-off recipes said:

- `cd backend` → **there is no such directory.** `backend` is the *deployed service name*; the repo
  `make init` clones is **`app`**. This was the only `cd backend` in the corpus and it was in the one
  document you reach for when you do not want to read a guide.
- `cd cms`, `cd jobsimulation` → two directories `make init` no longer creates **and two schemas the
  platform no longer creates**. `repos.yml` says it in the platform's own words: *"`app` is the ONLY repo
  with migrations to run."* Running those legs is precisely the failure **gate clause 4** exists to
  prevent, here on the corpus side rather than the tooling side.
- `cd stack-dev/backend` in "Check migration status" — same non-existent directory.

Repaired to a single `app` leg with the rule and its citation stated inline.

**2. `corpus/services/graphql-wundergraph.md:254` — the one archived-service doc missing its caveat.**
Measured across the six: `cms` `jobsimulation` `roadrunner` `messenger` `storage` all disclose *"`make init`
no longer clones it — clone it by hand"* beside their fenced `cd`. `graphql-wundergraph` did not — and it
is the strongest case, since the router was **deleted from the platform outright** at `2adcf71`. Caveat
added, matched to its five siblings.

## Instrument non-vacuity

The census does not return zero on either arm — 333 resolving sites and 61 non-resolving, with the
non-resolvers partitioned into three classes each independently checkable. The English-false-positive
class is itself the anti-vacuity control: an instrument that could not over-fire would not have produced
`make opting`.

## Seal grading — `46738fa`, sealed before any measurement

| id | prediction | outcome |
|----|---|---|
| `P-235-1` | ≥ 40 distinct `make <target>` invocations | **CONFIRMED — 394 sites / 36 distinct** |
| `P-235-2` | `platform/Makefile` declares ≥ 20 targets | **REFUTED — 15** |
| `P-235-3` | ≥ 1 documented target does not exist | **REFUTED as stated.** 1 non-resolver, and it is **unmeasurable** (`proto` is never cloned), not missing |
| `P-235-4` | ≥ 1 non-existent target in a fenced block | **REFUTED for `make` — but CONFIRMED for the sibling half**: 13 fenced `cd <dead-repo>` sites, 2 repaired |
| `P-235-5` | misses concentrate in ops guides | **REFUTED — they concentrate in `corpus/services/*` archived-service docs**, where they are correct-and-disclosed. The one real defect *was* in an ops guide (`quick_ops.md`), so the intuition was right for the wrong population |

**1 confirmed · 4 refuted.** The hypothesis — *"guides written before the merge program should carry a
dead target"* — is **falsified**: the `make` surface is aligned. The defect was one layer over, in the
directory half of the same instruction, and only asking the second question found it.

## Guard family

`24 GREEN · 0 RED · 0 could-not-check · 5 not-run` (`--platform stack-demo/platform --allow-not-run`),
re-run after both repairs — unchanged from the run-27 baseline.

## Close — 2026-08-10

**Outcome:** the `make` surface is **aligned** — 394 documented invocations, 333 resolving in
`platform/Makefile`, 43 correctly repo-relative, 17 English false positives, and **1 unmeasurable, 0
dead**; plus **0 of 15 platform targets undocumented**, a reverse-coverage number the corpus never had.
The hypothesis was falsified and the real defect was found by asking the *other* half of a runnable
instruction: **`quick_ops.md`'s migration recipe had 3 of 3 lines fail**, including the corpus's only
`cd backend` — a directory that does not exist, in the one document you open to avoid reading a guide.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-235-1` (a repo-relative command is graded against the clone set, never one
Makefile), `D-M257x-235-2` (`force-gen` is unmeasurable, not wrong — not repaired).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — no pytest section run; no rext code changed. Two corpus documents changed;
guard family re-run at platform reach, 24 GREEN / 0 RED.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-235-runnable-block-has-two-halves` → **new.** A fenced command is a **target** *and* a
  **directory**, and this corpus is aligned on the first while carrying 13 copy-pasteable instances of the
  second. Both halves are enumerable; only one has ever been asked.
- `ROUTE-M257x-235-fence-scope-is-unread` → **new.** The classifier cannot tell which repo a fenced block
  addresses. A block-scope reader (nearest preceding `cd`) would make this class exactly gradable and is
  the cheapest remaining mechanical win.
- All prior routes → open, unchanged.

**Lessons:**
1. **Ask what the declared number is a number OF — for commands too.** Grading repo-relative `make`
   invocations against one Makefile produced a 61-site "failure" that was entirely denominator.
   `§8`'s iter-229 rule, in a new place.
2. **A falsified hypothesis is where to look next, not where to stop.** The `make` half came back clean;
   the *directory* half of the identical instruction held the real defect. The iter's value came from the
   question asked after the seal was refuted.
3. **A cookbook is the highest-severity document in the corpus.** `quick_ops.md` exists to be pasted
   without reading, so a dead line there costs more than the same line in a guide — and it is the one
   place with no surrounding prose to disclose the caveat.

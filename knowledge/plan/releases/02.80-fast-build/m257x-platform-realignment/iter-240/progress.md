**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), continuing
`ROUTE-M257x-238-claude-md-fences-are-unmaintained`.

# iter-240 — the toolchain floor: the sixth runnable input, and the first that came back mostly right

## What was censused

The **toolchain version** — the sixth input of a runnable instruction, after `make` targets (235), `cd`
targets (236), environment variables (237), `npm`/`pnpm` scripts (238) and slash-commands (239). It is the
input that decides whether a person gets a working stack **before the first command runs**, and it fails in
the least legible way: a wrong `cd` gives a clear error; a wrong required-Go floor gives either a silent
background toolchain download or a `go.mod requires go >= …` that reads like a broken checkout.

⚠️ **Fifth consecutive iter whose naive population would have been mostly its own instrument.** A
tool-name-adjacent-to-a-number regex over the live corpus returns **471 distinct (tool, version) pairs
across 1,108 sites**, and its top hits are `go 14`, `docker 48`, `go 123` — **line numbers inside
`file:line` citations**. The graded population is therefore the **enumerated prerequisite claims of the
bring-up path**, read against each source of truth directly, never a doc quoting another doc.

## The reading — 6 axes, graded item by item

| axis | corpus says | source of truth | verdict |
|---|---|---|---|
| host **Go**, remote-VM block | `setup_guide.md:134` **Go 1.25.x**, quoting `toolchain go1.25.12` | 6 rext `go.mod`: **all `go 1.25.0` + `toolchain go1.25.12`** | ✅ **exact** |
| host **Go**, macOS + Linux blocks | `:36` / `:94` **v1.23+** | same | ❌ **LOW BY TWO MINOR VERSIONS** |
| **Node** | `:46` / `:99` *"v24+ required"*, quoting `engines.node ">=24.0.0"` | `next-web-app` `>=24.0.0` · `studio-desk` `>=24` · `ant-academy/code` `>=22` | ✅ covers all three |
| **pnpm** | `corepack enable` | `next-web-app` `packageManager: pnpm@10.30.3` | ✅ corepack honours it exactly |
| **Go base image** | `golang:1.26-bookworm`, 3 sites | `app/Dockerfile:2`, `sentinel/Dockerfile:2` | ✅ both readable clones confirm |
| **atlas** | installs `-latest`, no version claimed | — | not gradeable |

> ### 5 of the 6 axes are correct. The toolchain is the FIRST of the six runnable inputs to come back
> ### mostly clean — and the single defect is the same self-contradiction shape as 236 / 237 / 238 / 239.

**`setup_guide.md` stated the Go floor three times, in its three OS blocks, and disagreed with itself.**
`v1.23+` for macOS, `v1.23+` for Linux, **`Go 1.25.x` for the remote VM** — the last one derived correctly
and even quoting `toolchain go1.25.12`.

**The consequence is real and it is not a warning.** All six `rosetta-extensions` sections that own a
`go.mod` declare `go 1.25.0` + `toolchain go1.25.12`, and those tools (`stacksecrets` / `stacksnap` /
`stackseed`) build **on the host**. On Go 1.23 the build does not simply work: under the default
`GOTOOLCHAIN=auto` it silently downloads a 1.25 toolchain (so it needs network, on a box the guide has
just finished configuring offline-ish), and under `GOTOOLCHAIN=local` it hard-fails with
`go.mod requires go >= 1.25.0`. The guide's own remote-VM block already knew this — **and it is the block
fewest developers read.** (That block sat at `:124` when this iter's probe commit sealed the reading and at
`:134` after the repair added 10 lines above it; every line number in this document is the **post-repair**
tree, and the answer-key test pins the pre-repair one by sha rather than by offset.)

`app` and `sentinel` declare `go 1.26.x` and are **deliberately not** in the derivation: they build inside
Docker, so the host never compiles them, and a floor raised on their account would be a false requirement
(`D-M257x-240-2`, pinned by a regression test).

## The repair — 2 sites, with the derivation stated in place

`corpus/ops/setup_guide.md:36` (macOS) and `:94` (Linux) → **v1.25+**, each carrying *why* the number is
1.25 and not 1.23, and why `app`/`sentinel`'s 1.26 does not raise it. The remote-VM block was already
right and was not touched.

## The fence — `stack-core/toolchain_floor_guard.py` (rext `676c459`, pushed to origin)

Two arms: the stated Go floor ≥ the highest rext `go` directive; the stated Node floor ≥ the highest
frontend `engines.node` minimum.

* **Grades `<`, never `≠`** — a prerequisite is a floor, so a conservative round-up stays green
  (`D-M257x-240-1`). The guard says *"floors, not pins"* in its own output so a green cannot later be
  read as a claim of exactness.
* **Fails CLOSED** — it reads two trees that need not exist on every host (the rext checkout, a clone
  set); either absent is **exit 2**, never 0. `§9` iter-174: a capability probe that fails OPEN disarms
  the check it guards.
* **14 tests** — a mutation control per arm *including the realistic direction* (nobody edits the guide;
  rext bumps to 1.26 and all three stated floors go RED), **five** anti-vacuity paths, and a **real answer
  key** rebuilt from `git show 2f89b3c:corpus/ops/setup_guide.md` — this iter's own sealed
  pre-registration — which must report exactly the 2 `v1.23+` sites **and still see all 3**, proving the
  guard distinguishes the correct block rather than flagging the document wholesale.

**The reach fix is the part worth reading.** The first selector matched `**Go**` and reached **2 of 3**
sites — and the site it could not see was the one that had the number **right**. That green would have
rested entirely on the two sites just repaired, and would have persisted if they rotted back to a spelling
it also could not parse. Widened to `\*\*Go\b` and pinned by a reach test **before** the fence shipped
(`D-M257x-240-4`) — `§5`, *a verdict without its reach is not a verdict*.

## Pre-registration — scored 3 confirmed / 2 refuted

| claim | prediction | result |
|---|---|---|
| `P-240-1` the "Go 1.25.x" claim is stale; ≥1 rext section past 1.25 | REFUTED claim | **REFUTED — my prediction, not the corpus.** All 6 sections are exactly `go 1.25.0`/`toolchain go1.25.12`; `:134` is precisely right |
| `P-240-2` `next-web-app engines.node >=24` quoted correctly | confirmed | **CONFIRMED** — actual `>=24.0.0`; the guide quotes it verbatim, `dev-for-dummies` paraphrases to `>=24` (same claim) |
| `P-240-3` ant-academy `>=22` quoted correctly | confirmed | **CONFIRMED** — exact |
| `P-240-4` ≥ 1 stale host toolchain version | ≥ 1 | **CONFIRMED** — 2 sites, and it is the load-bearing one |
| `P-240-5` `golang:1.26-bookworm` stated nowhere | 0 sites | **REFUTED** — stated at **3** sites, all correct |

**Both refutations are the corpus being better than predicted, and that is worth stating plainly.** This
iter opened expecting the version axis to be rotten because the five before it were, and it is not: two of
the five pre-registrations bet against documents that turned out to be exactly right. A protocol that only
ever confirms its suspicions is measuring the suspicion. **The one real defect was found by the axis the
iter did *not* single out** — the same shape as iter-238, where `ant-academy` was predicted to drift and
`studio-desk` was the one that had.

## Close — 2026-08-10

**Outcome:** the toolchain-floor axis is censused and fenced; the one defect — `setup_guide.md` stating a
Go floor **two minor versions below what the rext host tools require**, at 2 of its 3 OS blocks, while its
third block had it right — is repaired with its derivation, and the guard family goes **25 → 26 GREEN /
0 RED**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-240-1` (floors, not pins — grade `<` not `≠`) · `D-M257x-240-2` (`app`/`sentinel`
are out of the host derivation, pinned by a test) · `D-M257x-240-3` (3 under-declared prerequisites
classified, none repaired) · `D-M257x-240-4` (reach widened before shipping, not after).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — `stack-core` (pytest 8.4.2 / CPython 3.9.6, Python):
`tests/test_toolchain_floor_guard.py` **14 passed / 0 failed**; iter-239's
`tests/test_skill_invocation_guard.py` still **17 passed**. No other section run. Guard family at platform
reach: **26 GREEN / 0 RED / 0 could-not-check / 5 not-run** (the five needing `--range`/`--ledger`).

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` → **new.** `setup_guide.md` states its
  prerequisites **once per OS block**, so every prerequisite is a claim made 2–3 times that nothing keeps
  in step; this iter's defect was two of three copies rotting while the third stayed right.
  `toolchain_floor_guard` now holds the Go and Node floors across all copies, but **Docker, atlas, git and
  VS Code are stated the same way and are not derived from anything** — the fence covers the two axes that
  have a machine-readable source, and says so.
- `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only` → open.
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → **open, six-for-six.**
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → open, and iter-240 is a third independent hit: the
  host-Go floor is a **native** requirement that `app`/`sentinel`'s **container** Go directives look like
  they should raise, and separating the two is exactly what `D-M257x-240-2` had to write down.
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → open.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → open.
- `ROUTE-M257x-236-host-is-the-unreliable-witness` → open.
- `ROUTE-M257x-235-fence-scope-is-unread` → open.
- `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **A version that is right in one OS block and wrong in the other two is invisible to every reader,
   because no reader reads more than one block.** The document was internally inconsistent for as long as
   it has had three blocks, and the block that was correct is the one for a remote VM.
2. **Derive the floor from what runs on the HOST, and prove the exclusion with a test.** `app`'s
   `go 1.26.4` is the biggest number in the tree and the wrong one to use; a comment saying so would have
   been reverted by the next contributor's instinct to "be thorough."
3. **Widen a fence's reach before shipping it, not in the next harden pass** — especially when the sites it
   cannot see are the sites that are *correct*, because then the green it prints is load-bearing on exactly
   the content that just changed.
4. **Bet against your own suspicion sometimes.** Two of five pre-registrations predicted rot and found
   none; the real defect was on the axis the iter did not single out. Five consecutive dirty inputs make
   "it will be dirty" feel like knowledge, and it is not.

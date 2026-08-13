**Type:** tik — under `TOK-08` (census a mechanical class exhaustively).

# iter-236 — can you `cd` where the corpus tells you to?

## Denominator, derived not remembered

`repos.yml` @ `platform` HEAD declares **4** clones — `app`, `sentinel`, `next-web-app`, `studio-desk` —
plus `platform` itself and the documented manual clone `ant-academy`. Six roots. Everything else a fenced
`cd` names is a subdirectory, a stack workspace, a `.agentspace` path, or a defect.

Instrument: `.agentspace/scratch/work-m257x/cd236.py` over `corpus/**` + `CLAUDE.md` +
`.claude/skills/**`, fenced blocks only, subdirectories resolved against the actual clone trees with
`git ls-tree` rather than the working tree.

## Population and classification

**158 fenced `cd <path>` sites · 58 distinct targets.**

| class | sites |
|---|---|
| clone-root | 75 |
| variable-or-relative (`$VAR`, `~`, `..`) | 23 |
| archived-repo | 15 |
| UNKNOWN-ROOT (incl. placeholders like `[service]`, `<stack>`) | 13 |
| agentspace-or-dotdir | 13 |
| subdir-of-clone | 10 |
| workspace-root | 7 |
| **SUBDIR-MISSING** | **2** |

## The finding: one directory, documented three ways, two of them silent

Every real defect collapsed onto the **same path** — `app/studio/`, the Studio-Room pipeline root — and the
corpus already knew the right answer in one place while contradicting it in two others:

| site | said | state |
|---|---|---|
| `corpus/services/cms.md:298-308` | `cd app/studio  # was: cms/studio`, with a full paragraph on why `cms` is not cloned | **correct and disclosed** — the model |
| `corpus/services/studio-room.md:347` | `cd app/studio` under a 5-line comment block citing `studioManager.go` | **silent**: never says `app/studio` is absent from the clone |
| **`CLAUDE.md:533`** | `cd studio-room` | **contradicts its own Tier-2 section 300 lines up**, which says Studio-Room is *"Embedded inside the `app` container … pulled into the image by CI; never a standalone deployment"* |
| `corpus/ops/run_guide.md:341` | `cd stack-dev/studio-room` | **silent**, and this is the guide `/dev-up` executes |
| `corpus/ops/update_guide.md:165` | `cd stack-dev/studio-room` | **silent**, and `make pull` does not touch this repo either |

**Measured ground truth:** `app/.gitignore:78-79` reads *"Python studio runtime
(anthropos-studio-room) — pulled at build via additional_repo, like cms"* → `studio/*`. `git ls-tree -d
HEAD studio` in the `app` clone returns **nothing**. So the directory is **absent from a fresh `make init`
clone** and present only where an image build or a hand-clone put it.

> **And it is present on this box** — `stack-demo/app/studio` and `stack-dev/studio-room` both exist on
> disk, which is exactly why five documents could say five things for months without anyone tripping. *A
> path that works on the author's machine and fails on a fresh clone is the hardest kind of doc defect,
> because the corpus's own host is the least reliable witness.* iter-233 had already flagged the
> git-ignored `studio/` embed as a clone-set anomaly; this iter is what it was an anomaly *of*.

**Four repairs landed** — `CLAUDE.md`, `studio-room.md`, `run_guide.md`, `update_guide.md` — each pointed
at `app/studio` with the `.gitignore` citation, modelled on the `cms.md` block that was already right.

## Prefix drift — real, and mostly benign; the honest reading

**6 of 16 repos are reached under more than one prefix**: `platform` 29 bare / 17 `stack-dev/`;
`ant-academy` 5 bare / 4 `stack-dev/` / 1 `stack-demo/`; also `app`, `next-web-app`, `studio-desk`,
`studio-room`. `P-236-4` is confirmed on the number.

**It is not scored as a defect.** A bare `cd platform` is correct after a preceding `cd stack-dev`, and the
instrument reads one line at a time, so it cannot see the preceding line. Calling 46 sites wrong on that
basis would be the iter-235 denominator error again. Reported as a measured property of the corpus, not a
finding — and it is genuine information: the workspace convention moved twice and the prose still shows both.

## Stated instrument limits (both disclosed, neither tuned away)

1. **The disclosure window is ±12/+6 lines**, so a **document-level** banner is invisible to it.
   `corpus/services/chronos.md:177/:196` were flagged UNDISCLOSED and are **not** defects — the doc opens
   with *"⚠️ Decommissioned … no longer cloned by `make init` … preserved for historical context."*
   Hand-checked, dismissed, and the limit recorded rather than the window widened to fit.
2. **13 UNKNOWN-ROOT sites are mostly placeholders** (`[service-name]`, `[repo-with-conflict]`,
   `<stack>/…`, `<the`) — a template, not a path. Reported in their own class.

## Instrument non-vacuity

Both arms fire: 75 clone-roots resolve positively and 2 SUBDIR-MISSING negatively, with the negative arm
independently confirmed against `.gitignore` and `git ls-tree`. The placeholder and workspace classes are
the over-fire control — an instrument that could not over-fire would not have produced `cd <the`.

## Seal grading — `b17e08c`, sealed before any measurement

| id | prediction | outcome |
|----|---|---|
| `P-236-1` | ≥ 100 fenced `cd` sites | **CONFIRMED — 158** |
| `P-236-2` | ≥ 20 distinct targets | **CONFIRMED — 58** |
| `P-236-3` | ≥ 1 undisclosed missing directory | **CONFIRMED — 4 repaired**, all one path (`app/studio`), incl. `CLAUDE.md` contradicting itself |
| `P-236-4` | same repo under ≥ 2 prefixes | **CONFIRMED — 6 repos**; scored as a property, not a defect |
| `P-236-5` | ≥ 1 path existing in no form | **REFUTED** — every non-placeholder target exists somewhere; the defects are *absent-from-a-fresh-clone*, which is a different and subtler thing |

**4 confirmed · 1 refuted**, and the refutation sharpened the finding: nothing the corpus names is
fictional; what it gets wrong is **when** a path is there.

## Guard family

`24 GREEN · 0 RED · 0 could-not-check · 5 not-run` (`--platform stack-demo/platform --allow-not-run`),
re-run after all four repairs — unchanged.

## Close — 2026-08-10

**Outcome:** 158 fenced `cd` sites censused against a denominator derived from `repos.yml`. Every real
defect collapsed onto **one path**, `app/studio` — absent from a fresh `make init` clone
(`app/.gitignore:78-79`), present on this box, and documented **five different ways**, of which one was
right and four were silent, including **`CLAUDE.md` contradicting its own architecture section**. All four
repaired against the one that was already correct. Prefix drift confirmed on 6 repos and deliberately
**not** scored as a defect.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-236-1` (prefix drift measured, not repaired), `D-M257x-236-2` (chronos.md is
disclosed at document level — the window is the limit, and it is recorded, not widened).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — no pytest section run; no rext code changed. Four corpus/root documents
changed; guard family re-run at platform reach, 24 GREEN / 0 RED.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-236-host-is-the-unreliable-witness` → **new, and general.** Three of the four defects are
  invisible on this box because a build populated the path. Any *"does the documented path exist"* check
  must read the **clone's git tree**, never the filesystem — the same shape as `§8`'s *a remote-tracking
  ref is a cache, not a remote*, one layer down.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → **new.** Disclosure is a property of a
  **document**, not a 12-line window. A checker for this class needs a doc-level banner reader.
- `ROUTE-M257x-235-fence-scope-is-unread` → **still open**, and now paired with the above.
- All prior routes → open, unchanged.

**Lessons:**
1. **A path that works on the author's machine and fails on a fresh clone is the hardest doc defect.** Five
   documents disagreed about `app/studio` for months because on every box that mattered, it was there.
2. **When a corpus contradicts itself, one side is usually already right — repair toward it.** `cms.md`
   had the correct, fully-reasoned block. The fix was to make four sites match it, not to invent wording.
3. **Confirm a prediction and still decline to act on it.** Prefix drift was predicted and measured, and
   scoring it as a defect would have re-run iter-235's denominator error at 46× the volume.

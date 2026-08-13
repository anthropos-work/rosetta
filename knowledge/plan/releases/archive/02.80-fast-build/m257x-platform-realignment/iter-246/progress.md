**Type:** tik (under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07))

# iter-246 — the six-for-six route, censused: the fences are CLEAN, and nothing was watching them

## Phase B — the `CLAUDE.md` census, complete

All **11 fenced blocks / 55 runnable command lines / 12 comment lines**, graded against the clone set:

| block | what was checked | verdict |
|---|---|---|
| ×4 slash-invocation blocks | `/dev-up`, `/dev-up 2`, `/setup-github`, `/stack-update`, `/stack-update dev-2`, `/update-knowledge` | ✔ (already fenced by `skill_invocation_guard`, iter-239) |
| the Makefile workflow (15 lines) | all **10** platform targets — `init up migrate pull status down ps logs dev reset-db` | ✔ all declared |
| | `make init # Clone the 4 repos in repos.yml: app, sentinel, next-web-app, studio-desk` | ✔ **the positive control**: `repos.yml` declares exactly those four |
| app / sentinel | `make setup`, `make gen`, `atlas migrate apply --env local`, `make proto`, *"NOT `make gen`"* | ✔ — `app/atlas.hcl:6` declares `env "local"`; sentinel declares `initdb proto` and **no `gen`**, so the warning is right |
| next-web monorepo | `pnpm dev`, `pnpm build`, `pnpm test` | ✔ all three declared |
| studio-desk | `cp .env.example .env`, `npm run dev`; *"frontend 9100 (`vite.config.ts:10`), backend 9000 (`.env.example:4`)"* | ✔ — **both cited lines are exactly those values** |
| studio-room | `cd app/studio`, `requirements.txt`, `gen.py` | ✔ all present |
| ant-academy | `cp .env.example .env.local`, `npm run dev # port 3077`; mobile `pnpm run dev:web # port 8555` | ✔ — `dev` is `next dev --port 3077`; `dev:web` → `scripts/dev-server.sh` → **`PORT=8555`**, a three-hop chain that checks out |

**`CLAUDE.md`'s fences are CLEAN. Zero defects in 67 lines.**

## The reading that matters — the route was RIGHT, and its diagnosis was wrong

`P-246-1`'s falsification said a result of ≤ 1 means *"six-for-six was coincidence, not structure."*
**That branch fires on the arithmetic and the conclusion it names is WRONG**, and saying so is the honest
close rather than following the pre-registered wording off a cliff.

**The zero is the route's own repairs holding.** Six iters found six defects and each fixed **the one it
found**. Nothing was left behind that would notice the seventh. And the mechanism generalises past this
one document: **iters 235–238 censused four runnable-input surfaces — `make` targets, `cd` directories,
environment variables, frontend scripts — repaired every defect, and fenced NONE of them.** Two of those
four still have no guard at all. So the class was never closed; only its instances were, one per iter, in
whichever document happened to be read.

*A class is not closed by a repair; it is closed by an enumeration that keeps running* — `§8` iter-176,
which this milestone wrote and then did not apply to its own four censuses.

## Phase D — the fence: `stack-core/fence_command_guard.py` (FENCE-M257x-iter246)

Corpus-wide, not `CLAUDE.md`-only, because the route's cause is not specific to one document. It tracks
the working directory each fenced block's own `cd` lines establish, then asserts three arms: **`cd`
targets exist · `make` targets are declared by that directory's Makefile · `npm run`/`pnpm` scripts by its
`package.json`.**

**Live verdict: 188 commands resolve (103 cd / 68 make / 17 npm) across 621 blocks. Zero findings.**

### Its first run reported 39 findings and every one was the instrument

Recorded in full, because this is now the fourth consecutive iter whose first pass was dominated by its
own tooling, and each reduction is a checked-in regression test:

| # | 39 → | the defect |
|---|---|---|
| 1 | 4 | **`stack-dev/` is provisioned with 2 repos on this host; `stack-demo/` has 13.** A literal reading turned **39 correct `cd stack-dev/platform` lines in 8 documents** RED. `cd stack-dev/platform && make up` is a claim about **`platform`'s Makefile**, not about which workspace you are in — so a `stack-<X>/` path is retried under any workspace that carries the repo, and **the substitution is printed** (`stack-dev x32`) |
| 2 | 2 | **cwd composition.** `cd a/code` then later `cd a/mobile` became `a/code/a/mobile` — a fence is not a shell transcript; stanzas assume a fresh start. Resolution is now **fresh-first** |
| 3 | 1 | **`cd stack-dev/app` then `cd stack-dev/sentinel`** still composed, because a workspace-anchored path is repo-root-absolute in this corpus's vocabulary and must never be appended to the current one |
| 4 | 0 | **the rext trees were not roots.** `cd playthroughs/e2e` is a real directory in the tooling monorepo — the last finding, and a reach gap |

And two holes its **own unit tests** caught before it shipped: a fabricated workspace (`stack-nowhere/…`)
was silently substituted to a real one — *a pardon that travelled with the repo name* — and arm A could
never fire for any `stack-*` path.

### Resolve or refuse, and the line is principled

**A repo-level miss is REFUSED; a sub-path miss is a FINDING.** `cd stack-dev/chronos` (decommissioned)
and `cd stack-dev/experiments` (a separate org repo you clone by hand) are indistinguishable from a typo
with the evidence available — that is iter-244's *wrong-vs-uncheckable* rule cutting the other way, and
the bucket is honest precisely because it does not claim to know. Once the repo IS located, its interior
is fully knowable, so a missing subdirectory fires.

Every refusal is bucketed by reason and printed: `make: no establishing cd ×40 · npm: no establishing cd
×28 · cd: placeholder ×24 · cd: unanchorable ×12 · cd: unanchored single segment ×7 · cd: no provisioned
workspace carries this repo ×3`.

## Phase E — pre-registrations, graded after the last edit

| id | prediction | outcome |
|---|---|---|
| **P-246-1** | ≤ 6 of 55 runnable lines fail | **HELD — 0.** The falsification branch fires arithmetically; its stated conclusion is rejected with reasons (above) |
| **P-246-2** | ≥ 1 comment-borne claim is false | **REFUTED** — every one checked out, including a three-hop port chain |
| **P-246-3** | no guard takes a fence body as subject, bar `skill_invocation_guard` for `/`-lines | **CONFIRMED** — and now one does |
| **P-246-4** | ≥ 1 `cd` target does not exist | **REFUTED** — all 7 exist |
| **P-246-5** | **control** — the 4-repo `make init` comment is correct | **HELD** — the instrument reports a correct line as correct |

**3 of 5 refuted or rejected.** A census whose predictions all held would have told us nothing we did not
already believe.

## Close — 2026-08-10

**Outcome:** the milestone's longest-standing structural lead, censused to completion. `CLAUDE.md`'s
fences are **clean — 0 defects in 67 lines** — and that zero is the route's own repairs holding, not a
refutation: **iters 235–238 censused four runnable-input surfaces, repaired every defect, and fenced
none.** Shipped the missing half — `fence_command_guard` (FENCE-M257x-iter246), corpus-wide, **188
commands resolving across 621 blocks, zero findings**, family **27 → 28 GREEN**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-246-1` (the falsification's arithmetic fires and its conclusion is rejected, with
the evidence) · `D-M257x-246-2` (workspace substitution, printed, and gated on the named workspace being
real) · `D-M257x-246-3` (repo-level miss refused, sub-path miss flagged) · `D-M257x-246-4` (corpus-wide
scope, not `CLAUDE.md`-only, because the cause is not document-specific).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
`test_fence_command_guard` **20 passed**, plus `test_fence_registry_population` + `test_guard_family`
**90 passed / 0 failed** together. Guard family (`--platform`, from repo root): **28 GREEN / 0 RED /
0 could-not-check / 5 not-run**.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → **CLOSED.** Censused to zero and fenced; the
  diagnosis is corrected in the closing note (the fences were not unmaintained — they were unfenced).
- `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` → **new.** `make` targets and `cd`
  directories are now covered by `fence_command_guard`; the **environment-variable** and **frontend
  script/port** censuses (iters 237, 238) still have none, and `ROUTE-M257x-237-critical-env-list-is-unfenced`
  is the same gap seen from the other side.
- `ROUTE-M257x-245-guard-family-green-is-not-suite-green` · `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` (repo-relative + clone-absent halves) ·
  `ROUTE-M257x-h59-range-anchors-are-ungraded` (which-line half) ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` · `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-critical-env-list-is-unfenced` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **A pre-registered falsification binds its ARITHMETIC, not its interpretation.** `P-246-1` fired and
   its stated conclusion — *"coincidence, not structure"* — is false. The seal exists to stop the number
   being argued after the fact; it does not license following a wrong inference off a cliff. Say which
   half fired and why the other is rejected.
2. **A repair rate of one-per-iter is the signature of an unfenced class**, not of a rare defect. Six
   iters found six defects in one document; the seventh was never going to be found by a seventh reading.
3. **A guard's roots ARE its reach, and a missing root reads as a corpus defect.** Three of this
   instrument's four reductions were roots it did not have — a workspace with a different name, a
   workspace-anchored path, and the tooling tree the corpus routinely `cd`s into.
4. **Test the pardon, not just the rule.** Both holes this guard's tests caught were in its *excusing*
   logic, not its *asserting* logic — a substitution that travelled with the repo name, and a refusal
   that swallowed an entire arm.

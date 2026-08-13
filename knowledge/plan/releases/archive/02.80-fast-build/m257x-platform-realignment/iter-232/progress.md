**Type:** tik — under `TOK-08`.

# iter-232 — does the corpus cite files the platform has DELETED?

## The census

Every backticked `<seg>/<path>:<line>` in `corpus/**` + `CLAUDE.md` whose first segment names a cloned
repo, resolved with `git cat-file -e origin/main:<path>`.

| quantity | value |
|---|---:|
| distinct (repo, path) pairs cited | **111** |
| citation sites | **345** |
| repos covered | 13 |
| reported missing at `origin/main` (first pass) | **11** |
| **genuinely gone** | **0** |

Per-repo: `app` 73 · `jobsimulation` 6 · `messenger` 5 · `cms` / `rosetta-extensions` / `next-web-app` 4 ·
`roadrunner` / `storage` 3 · `graphql-wundergraph` / `studio-desk` / `sentinel` / `ant-academy` 2 ·
`platform` 1.

## All 11 findings were the instrument's, and they failed in four distinct ways

| # | reported missing | what it actually is |
|---|---|---|
| 1 | `app/.../resolver_queries.go` | the corpus wrote a literal **ellipsis** (`app/.../…`); not a citation shape at all |
| 3 | `app/{core/main.ts, services/config.ts, services/userService.ts}` | **`app/` is a DIRECTORY inside `studio-desk`**, not the repo `app` — all three exist |
| 4 | `app/studio/{gen.py, services/ai.py, tools/pdf2md.py, .gitignore}` | **studio-room** paths — the repo is embedded in the `app` **image** by CI, so the corpus is citing an image path. All four exist in `studio-room` |
| 2 | `studio-desk/{.env.example, src/routes/dev.ts}` | real, and ref-pinned to `41ee357x` — the lookup hit a **broken clone** (below) |
| 1 | `jobsimulation/ai/ai.go` | **shorthand**: the same paragraph states the full path `internal/jobsimulation/ai/ai.go` one line above and then abbreviates it |

### The instrument defect worth keeping

`stack-dev/studio-desk` has **no `origin/main` at all** — an empty ref list — and the census built its
clone map as `{name: dir}` over two roots, so `stack-dev` silently overwrote `stack-demo`. Five real files
read as deleted because a *second clone of the same repo* won a dictionary assignment.

**This is iter-230's defect recurring one iter later** (there it was eight missing clones; here it is one
extra, broken clone) and it is the same shape both times: **the census's own substrate was never
enumerated before the census was believed.** The fix is the same both times — hold *all* clones per name
and try each, never one.

## The instrument, proved both ways

`§9`: a census returning ZERO must prove its instrument.

- **Fabricated paths report missing** — `internal/does_not_exist.go`, `main_nope.go`: both missing. ✓
- **Genuinely deleted paths report missing, and present at the parent of their deleting commit** —
  `app/internal/ai/pipe/pipe.go` (`caa12f400`), `app/internal/ai/mistral/completion.go` (`2b3a65cf0`),
  `next-web-app/packages/ui/src/Coursebuilder/markdownBody.helpers.ts` (`72d4a7405`): all three
  `before=y / origin-main=n`. ✓

The second control is the load-bearing one: control 1 only proves the tool can say "no". Control 2 proves it
says "no" **about the exact class the census exists to find** — a file the platform really deleted.

### Predictions, graded — 1 HELD (weakly), 3 REFUTED

| id | prediction | result |
|----|-----------|--------|
| `P-232-1` | ≥ 200 distinct (repo, path) pairs | **REFUTED — 111** |
| `P-232-2` | ≥ 1 cited path missing at `origin/main` | **HELD literally — 11 reported.** The belief behind it is **refuted: 0 genuine** |
| `P-232-3` | missing paths concentrate in `app` | **REFUTED.** 7 of 11 *looked* like `app` and none of them was — the concentration was an artifact of `app` being the most-cited repo *name*, not of `app` losing files |
| `P-232-4` | ≥ 1 missing path in a live present-tense claim | **REFUTED — 0**, there being no genuine misses to be present-tense about |

## Close — 2026-08-10

**Outcome:** **0 of 111 cited platform paths are gone.** On the file axis the corpus is aligned with the
platform at `origin/main`. Every one of the 11 apparent misses was this census's own instrument, failing in
four distinct ways — including a second clone of one repo, with an empty `origin/main`, silently winning a
dictionary assignment.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-232-1` (the 11 are published as instrument defects, not quietly dropped),
`D-M257x-232-2` (no prose was "repaired" — there was nothing wrong with it).
**No `N`/`P` movement is claimed** — this iter took no graded reading.

**Suite state at close** — no pytest section run; this iter changed no rext code and no corpus prose.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-232-stack-dev-studio-desk-is-a-broken-clone` → **new.** `stack-dev/studio-desk` has an empty
  `origin/main`. Harmless to this census once found, but it is a clone every `stack-dev` guard also reads.
  Not repaired here (touching the clone set is frozen behind gate clause 1 — `D-M257x-230-2`).
- `ROUTE-M257x-232-first-path-segment-is-not-a-repo` → **new.** Three of the four extractor failure modes
  are one bug: the corpus writes repo-relative, image-relative, directory-relative and shorthand paths in
  the same backticked shape, and nothing disambiguates them. Any future path fence must resolve
  **candidate-repo-set**, not first-segment.
- All prior routes → open, unchanged.

**Lessons:**
1. **Enumerate the census's substrate before believing the census.** Twice in three iters, the first-pass
   finding list was majority-instrument: 8 of 14 (iter-230), **11 of 11** (here).
2. **Control 1 proves a tool can say no; control 2 proves it says no about the right thing.** A
   fabricated-input control would have passed here even if the census had been pointed at the wrong ref
   entirely. Only a *really deleted* file exercises the class.
3. **A prediction can be HELD by its number and refuted by its meaning** — for the second iter running
   (`P-231-2`, now `P-232-2`). Grade both, and say which one the close is about.
4. **A backticked `a/b/c.ext:NN` is four different citation kinds in this corpus** — repo-relative,
   image-relative, directory-relative, and shorthand-after-a-full-path. Any tool that assumes one will
   report the other three as defects.

# M14 — Progress

**Shape:** section · **Status:** build complete (all 5 sections landed)

## Section checklist (from overview Scope.In)
- [x] `dev-up` (consolidate setup-platform + start-platform → dev bring-up, drives M13 flow) + `dev-down` — §1, commit `e6c2da6`
- [x] Hard-rename ops skills → `stack-list`/`stack-seed`/`stack-snapshot`/`stack-update` (dev-N|demo-N target) — §2, commit `d6afb70`
- [x] Remove old skill dirs (setup-platform, start-platform, update-platform, demo-status/seed/snapshot) — §3, commit `d6afb70`
- [x] Update every reference: CLAUDE.md skill table, READMEs, corpus/ops/ guides, demo/ recipes — §4, commits `3b8e80e` (rosetta) + `b37e831` (extensions)
- [x] `demo-up`/`demo-down` retained + aligned with `dev-up`/`dev-down` — §5, commit `23b2398`

## Build notes
- KB-fidelity Phase 0b: **GREEN** (`kb-fidelity-audit.md`) — docs ARE the deliverable; every driven CLI exists.
- PR review (Phase 3): **1 finding** (D1 — `stack-seed --preset` is a skill-level shorthand, not a CLI flag),
  fixed inline (commit `28cb5c4`, → M14-D5). A/B/C clean; all links resolve; renamed skills' argument-hints
  verified against the `dev-stack`/`stackseed`/`stacksnap` CLIs.
- Decisions: M14-D1…D6 (D2/D3/D4 resolve Q1/Q2/Q3; D5 = the review fix; D6 = CHANGELOG convention).
- Extensions: 1 doc-only commit on `main` (`b37e831`) for the 4 extension-clone references. No tag needed
  (no CLI/section change); the renamed skills drive the existing M12/M13 binaries at their pinned tags.
- Three-fate guard: **0 deferrals** — every item (incl. the `--preset` finding) landed Fate-1 in M14.

## M14: Hardening

### Pass 1 — 2026-06-07
**Shape:** M14 is a rename/consolidation + docs milestone. The milestone-touched surface
(`git diff release/01.30-stack-party...HEAD`) is **entirely markdown** — the renamed/consolidated
SKILL.md files, the CLAUDE.md skill table, the corpus/ops guides + demo recipes, and the plan records.
**No new Go/Python/TS code** was authored in the rosetta repo, so there is **no code-coverage surface**
(Step 2a coverage measurement is N/A). The high-value hardening here is **contract / reference-integrity
guards** — the M11 shape (docs↔CLI flag-drift). Scope manifest = the diff name-list above; no source
file lacks a "test" in the conventional sense because the deliverable IS the docs.

**Hardening dimensions exercised (reference-integrity, not code-coverage):**
1. **Hard-rename invariant** — swept all 6 retired skill names (`setup-platform`, `start-platform`,
   `update-platform`, `demo-status`, `demo-seed`, `demo-snapshot`) across the **entire** tracked tree
   (excl. `knowledge/plan/` dated history + `CHANGELOG.md` immutable entries). **Result: zero
   non-provenance live references** — every remaining mention is an intentional "formerly /X → /Y"
   migration marker that maps the old name forward. No live surface instructs a reader to run a retired
   skill. CLAUDE.md table, READMEs, corpus/ops guides, demo recipes all clean.
2. **Skill→CLI contract resolves** — every binary/subcommand/flag the renamed skills invoke verified real
   in the `.agentspace/rosetta-extensions/` clone: `dev-stack` (`up`/`status`/`down` + `--no-snapshot`/
   `--no-setdress`/`--inject`/`--profile`), `rosetta-demo`, `stack_registry.py`, `stackseed`
   (`--stack`/`--seed`/`--reset`/`--dry-run`), `stacksnap` (`replay`/`capture`/`status` +
   `--surface`/`--stack`/`--dsn`/`--dry-run`/`--source`/`--store`), `provision-plan`. The `--preset`
   shorthand (M14-D5) correctly documented as a skill-level → `--seed presets/NAME.seed.yaml` mapping;
   all 4 promised presets (`dev-min`/`small-200`/`mid-500`/`large-1k`) exist.
3. **CLAUDE.md table ⇄ skill dirs** — perfect bijection: the table lists exactly the 14 skills on disk,
   **zero ghosts, zero missing rows**; every SKILL.md `name:` frontmatter matches its dir name.
4. **dev-up consolidation completeness** — element-by-element diff of the consolidated `dev-up`
   SKILL.md + reference.md against the former `setup-platform` + `start-platform` bodies (pulled from
   `release/01.30-stack-party`): **no dropped step** — STEP RUN discipline, confirmation policy, ops
   reports, TodoWrite progress (12-container set, 5 migration services, studio submodule, Ant Academy,
   Node v24+), expected-service table, error recovery (`make logs S=` via run_guide.md), critical rules
   — all preserved or expanded; the M13 dev set-dress flow (per-stack Directus + auto-snapshot +
   `dev-min` seed, default-on + non-fatal) correctly referenced. All 6 cross-reference links resolve.

**Stale reference fixed inline + guarded (Fate-1):**
- The M11-established stacksnap **docs↔parser flag-drift guard**
  (`rosetta-extensions/stack-snapshot/cmd/stacksnap/main_drift_test.go`) named the **retired**
  `/demo-snapshot` skill (+ its deleted `.claude/skills/demo-snapshot/SKILL.md` path) as the doc-side
  source of truth in 6 comment lines. M14 hard-renamed that skill to `/stack-snapshot` and deleted the
  old dir, so those comments pointed a future maintainer at a path that no longer exists. **Fixed** the
  comments to the renamed skill (flag contract unchanged) and **added `TestDocSourceSkillRename_M14`** —
  it asserts `.claude/skills/stack-snapshot/SKILL.md` exists and the retired `demo-snapshot` dir stays
  gone, turning the silent-comment-rot failure mode into a test failure. Skips gracefully when the
  sibling rosetta repo isn't reachable (per-stack clone / CI) so it only guards in the authoring copy.
  Negative-tested (reappearing the retired dir fails the guard). Committed in the extensions clone on
  `main`: `33fc525`. gofmt + go vet clean; full stacksnap package green.

**Bugs fixed inline:** none in rosetta production surface (docs/rename milestone). The one stale-reference
fix above was in the extensions-clone test guard.

**Flakes stabilized:** none (no flaky surface — the extension Go suite is deterministic; `-count=1` green).

**Knowledge backfill:** no KB-worthy findings beyond the guard itself — the rename invariant is encoded
in the test, and the existing corpus docs already document the rename forward (snapshot-spec.md §, the
"formerly /X" markers). The question was asked; nothing else warranted a corpus edit.

### Stop condition
Single pass. The reference-integrity surface is finite and fully verified (4 dimensions, all clean except
the one stale guard-comment, now fixed + actively guarded). No code-coverage surface to iterate on
(docs/rename milestone). Step 2b scan found nothing else worth adding → loop terminates after pass 1.

## Final review
_(filled at close)_

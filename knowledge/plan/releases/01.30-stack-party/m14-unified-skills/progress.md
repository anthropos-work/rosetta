# M14 — Progress

**Shape:** section · **Status:** `archived` (completed 2026-06-07)

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

## M14: Final Review

**Closed 2026-06-07** — clean straight-through close, **0 findings**, 0 fixes needed.

🔍 **M14 review found 0 findings:** 0 scope · 0 code-quality · 0 docs · 0 tests · 0 decision-triage.
A rename/consolidation + reference-sweep milestone — the close re-verified the reference-integrity surface
independently (not trusting the harden record) and found everything already clean.

### Scope (Phase 1) — clean
- All 5 sections checked off in the section checklist; every `overview.md` `In:` item delivered Fate 1.
- The single `Out:` item — "the safety doc (M15)" — is **Fate 2 (confirmed-covered)**: M15's `In:` list owns
  `corpus/ops/safety.md` (verified). Not a deferral.
- 0 unrecorded implementation choices (M14-D1…D6 cover every decision; Q1/Q2/Q3 → D2/D3/D4).
- No orphan TODO/FIXME/HACK tied to a v1.3 milestone. (The lone pre-existing `setup_guide.md:439` keys-mgmt
  doc-TODO predates the release, isn't tied to any deferral ledger, and is out of audit scope — see the M14
  deferral audit.)

### Phase 1b deferral re-audit — **GREEN**
`audit-deferrals/deferral-audit-2026-06-07-m14-close.md`. 1 inherited (DEF-M10-01 S3 blob bytes + cloud store →
v1.4, signed escape-hatch, not aged out); 0 new, 0 repeat, 0 chronic, 0 aged-out. M14 added zero deferrals.

### Code-quality / reference-integrity (Phase 2/2c) — clean
- **Hard-rename invariant:** all 6 retired skill names swept project-wide (excl. plan history + immutable
  CHANGELOG) → **0 live stragglers**; every remaining mention is a "formerly /X → /Y" provenance marker.
- **CLAUDE.md table ⇄ skill dirs:** perfect bijection — exactly 14 skills both in the table and on disk; zero
  ghosts, zero missing rows.
- **Frontmatter ⇄ dir:** every SKILL.md `name:` matches its dir name (0 mismatches).
- **Skill→CLI contract:** every driven binary exists in the extensions clone (`dev-stack`, `stackseed`,
  `stacksnap`, `provision-plan`, `stack_registry.py`, `rosetta-demo`).

### Documentation (Phase 3) — clean
- 0 broken links to deleted skill dirs; all `../`-relative links in the new/renamed skill SKILL.md files resolve
  to real targets; `dev-up`/`stack-update` `reference.md` present.

### Tests & benchmarks (Phase 4/8) — green
- Go: all 4 modules `-race -count=1` green — alignment 46 · clerkenstein 218 · stack-seeding 233 ·
  stack-snapshot **224** (+1, `TestDocSourceSkillRename_M14` the rename guard) = **721** (720→721).
- Python: 174 passed (dev-stack 38 + stack-core/demo-stack/injection/verify).
- gofmt + go vet clean (all modules); all 3 CLIs shellcheck-clean; `stack_registry.py` py_compile clean.
- **Flake gate: 5/5** (sequential shuffled on the M14-touched `stacksnap` package + full stack-snapshot module).

### Decision triage (Phase 5)
M14-D1…D6 are codebase-maintainer rationale (rename policy, consolidation boundary, target-detection UX,
`--preset` shorthand mapping, CHANGELOG convention). All already flowed forward where user-facing: the rename is
documented in the corpus (CLAUDE.md table, snapshot-spec §, the "formerly /X" markers) + encoded as a test
(`TestDocSourceSkillRename_M14`). No further knowledge-doc blend warranted → all **archive** in `decisions.md`.

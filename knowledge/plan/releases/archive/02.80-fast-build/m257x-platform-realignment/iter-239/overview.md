---
iter: 239
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-239 — the SKILL surface: the fifth input of a runnable instruction

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*

## Step 0 — re-survey

`TOK-08`'s next-target field is exhausted in the sense that names a class; the run's controlling lead is
the route iter-238 opened and did not work:

> `ROUTE-M257x-238-claude-md-fences-are-unmaintained` — four consecutive iters found a `CLAUDE.md` **code
> fence** contradicting `CLAUDE.md` **prose**. The prose is swept by `/update-knowledge`; the fences are
> not.

Iters 235–238 censused **four** of the inputs a runnable instruction needs: `make` targets (235), `cd`
targets (236), environment variables (237), `npm`/`pnpm` scripts (238). **A fifth input is untouched and
it is the one an agent reaches for first: the SKILL invocation.** `CLAUDE.md` opens with a 16-row
*Available Skills* table and four `bash` fences whose entire content is slash-commands. Nothing in the
guard family reads `.claude/skills/**` as a target — `service_registry_guard` grades compose services,
`platform_predicate_guard` grades compose profiles, `demo_knob_guard` grades `DEMO_*`.

This class has a **proven prior instance in the corpus's own words**: `corpus/ops/demo/demo-up-defaults.md`
records that `/demo-up`'s `argument-hint` *"conflated"* the `rosetta-demo` wrapper's `--profile`/`--services`
with `up-injected.sh`'s two-argument contract — **for releases**. So the surface is known to rot and has
been repaired exactly once, reactively.

**Substrate, measured before any grading** (`§5` — enumerate a census's substrate before believing it):

| substrate | measure |
|---|---|
| skill dirs with a `SKILL.md` on disk | **16** |
| `CLAUDE.md` *Available Skills* table rows | **16** |
| `CLAUDE.md` code fences (total) | **11** — 10 `bash`, 1 unlabelled |
| tracked `*.md` files scanned (excl. `.agentspace/`, `stack-*/`) | the repo's `git ls-files '*.md'` set |
| backticked `/token` sites in that set | **3,156 across 168 distinct tokens** |
| …of which the token is a skill on disk | **1,528 across 16** |
| …of which the token is a **retired** skill name | **96 across 6** |
| skills declaring an `argument-hint` | **16 of 16** |

⚠️ **The 168-token figure is why this needs a stated denominator.** The naive selector `` `/word` ``
matches URL paths (`/home`, `/profile`, `/sim`, `/courses`) far more often than slash-commands — 152 of
the 168 distinct tokens are **not** skills. A census run on the naive population would have been
dominated by its own instrument, which is the failure the last three consecutive first-passes hit.

## Hypothesis

The skill surface is the **least-fenced** runnable input in the corpus and the **first** one an agent
executes. Its mechanically-decidable half — *does the name resolve, does the guide path resolve, does the
invocation's argument appear in the skill's own `argument-hint`, is a retired name presented as runnable*
— has never been enumerated.

## Pre-registered claims — SEALED IN THIS COMMIT, before any grading

Numeric and falsifiable. Each is scored **confirmed / refuted / partly refuted** at close.

- **`P-239-1` (control arm — expected CLEAN).** All **16** skill names in `CLAUDE.md`'s table resolve to
  `.claude/skills/<name>/SKILL.md`. **Predict 16/16.** If this arm fires, the finding is larger than the
  iter and the iter re-scopes to it.

- **`P-239-2` (control arm — expected CLEAN).** Every `.md` path in the table's **Guide** column resolves
  on disk. **Predict 0 missing.** (One row carries `N/A (meta-skill)` and declares no path; it is counted
  as a non-path row, not as a pass.)

- **`P-239-3` (the hypothesis — expected DIRTY).** Among **runnable** slash-invocations — a line inside a
  fenced code block whose first token is `/<skill>` — corpus-wide, **at least one** carries an argument
  that its skill's own `argument-hint` does not accept. **Predict ≥ 1.**

- **`P-239-4` (the zero, and it must prove its instrument).** In the **live** corpus — `CLAUDE.md`,
  `README.md`, `corpus/**`, `.claude/skills/**` — **zero** of the six retired skill names
  (`setup-platform`, `start-platform`, `update-platform`, `demo-status`, `demo-seed`, `demo-snapshot`)
  appear as a **runnable** invocation. **Predict 0.** Per `§9` iter-149, a zero is not publishable until
  the instrument is proven against a **real answer key** — a historical tree in which it fires.

- **`P-239-5` (the risk call — pre-registered so it can be wrong).** The `P-239-3` defect will sit on a
  **`dev-*`** skill rather than a `demo-*` one, because v2.5–v2.8 have repeatedly proven the demo path
  live on `billion` while the dev path has not been exercised once in this milestone. **Predict: the
  defect's skill name starts with `dev-`.**

## Phase plan

1. Seal this pre-registration (this commit).
2. Build the census instrument with a **stated denominator** and both controls (mutation + anti-vacuity).
3. Grade `P-239-1`…`P-239-4` item by item — never as a rate (`§5`: a five-item list is graded item by
   item).
4. Repair toward the site that is wrong, not the site that is loudest; split claims that are two claims.
5. Close: score every pre-registration, route what does not land.

## Escalation conditions

- If `P-239-1` fires (a table row naming a skill that does not exist), that is a **larger** finding than
  the argument class and the iter re-scopes onto it in place, recording the substitution.
- If the argument grading cannot be made mechanical without paraphrase judgement (iter-234's refusal
  class), the arm is **declared unfenceable and said so** rather than graded softly.

## Acceptable close-no-lift outcomes

`P-239-3` returning **0** — i.e. the skill surface is already aligned — is a complete iter, provided the
instrument is proven with an answer key. That result would falsify this iter's central hypothesis and is
worth the same as a repair: it would make the skill surface the **first** of the five runnable-instruction
inputs to come back clean, and that is itself a measurement about where the corpus rots.

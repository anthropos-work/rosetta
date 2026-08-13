**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*), working
`ROUTE-M257x-238-claude-md-fences-are-unmaintained`.

# iter-239 — the SKILL surface: the fifth runnable input, and the first one that was wrong everywhere

## What was censused

The slash-command — the **first** runnable thing an agent executes in this repo, and the **last** of the
five inputs of a runnable instruction that nobody had fenced. Iters 235–238 took the other four (`make`
targets · `cd` targets · environment variables · `npm`/`pnpm` scripts) and every one came back with a
defect. Nothing in the guard family reads `.claude/skills/**` at all.

### The denominator, stated before the grading (`§5`)

| population | measure |
|---|---|
| naive selector — a backticked `` `/word` `` anywhere in tracked markdown | **3,156 sites / 168 distinct tokens** |
| …of those 168 tokens that are **not** skills (URL paths: `/home`, `/profile`, `/sim`, `/courses`) | **152** |
| live scope — `CLAUDE.md` + `README.md` + `corpus/**` + `.claude/skills/**` | **114 documents** |
| fenced lines whose first token starts `/` | **60** |
| …that are **filesystem paths**, not invocations (30 × `/tmp/<binary>` build outputs, 1 × `/home/<you>/` tree) | **31** |
| **runnable skill invocations — the graded population** | **29**, across 9 distinct skills |

**A census run on the naive population would have been 90 % its own instrument by distinct token**
(152 of 168) **and 52 % by site** (1,628 of 3,156). Even restricted to fenced lines it is **52 % artifact**
(31 of 60). Both splits are printed by the guard on every run, so the number can be audited rather than
trusted — and note the two 52 %s are a coincidence of different denominators, not one figure quoted twice.

## The reading — 4 arms, graded item by item

| arm | subject | population | findings |
|---|---|---|---|
| **A** name resolution | every fenced `/name` resolves to a `SKILL.md` | 29 | **0** |
| **B** the *Available Skills* table, **both directions** | 16 rows ↔ 16 skills on disk | 16 / 16 | **0** |
| **C** Guide paths | every `.md` in the Guide column exists | 15 paths (1 row declares `N/A (meta-skill)`) | **0** |
| **D** the argument contract | invocations naming a target for a skill whose hint declares `dev-N\|demo-N` | **8** | **8** |

> ### Arm D read **8 of 8**. Every runnable invocation in the live corpus that named a target for one of
> those skills disagreed with that skill's own declared contract. Not a rate — a total.

Two shapes, five and three:

* **bare `N`** (5) — `/stack-seed 1 --preset mid-500` against a hint that declares `dev-N|demo-N`, and a
  skill body that writes `--stack demo-N` / `--stack dev-N` at **7 of 7** of its own examples.
  `recipe-skill-progression.md` is the tell: it writes `stack: demo-1` **in the YAML two lines below**
  `/stack-seed 1` — qualified where a machine reads it, bare where a human does.
* **verb-first** (3) — `/stack-snapshot replay 1` against a hint **and** a flow line inside the skill's own
  body (`SKILL.md:32`, `/stack-snapshot N replay`) that both read target-then-verb.

**The declared contract was the MINORITY spelling in the corpus it governs** — 3 of the 4 runnable
`/stack-snapshot` examples inverted it. Volume is not a contract (`D-M257x-239-1`).

### Graded by consequence, not by class

`ParseStackN("1")` **succeeds** — it takes everything after the first `-`, and a name with no `-` parses
whole — so the bare form is accepted silently rather than failing loudly. And `isolation.TargetFor` keys
`IsProd` on the literal `"production"`/`"prod"` only, so **the never-write-prod firewall is untouched**.
This is a correctness finding, and the guard's docstring says so explicitly rather than borrowing safety's
weight (`D-M257x-239-5`).

## The repair — 8 sites, toward the declaration

`README.md` ×2 · `recipe-enterprise-onboarding.md` ×2 · `recipe-skill-progression.md` ×2 ·
`recipe-snapshot-world.md` ×2. Nothing in the skills changed: they were right at every one of their own
sites.

## The fence — `stack-core/skill_invocation_guard.py` (rext `ad63d1d`, pushed to origin)

Enumerates the population corpus-wide, holds it at zero, and prints its own substrate split. **Arm D's
subject set is DERIVED from the argument-hints, never named** — and that caught this iter's own instrument:
the hand-written first pass hard-coded **five** "generic stack-ops skills", where the derived set is
**three**. `stack-list` declares `(no args)`; `stack-update` declares `[dev-N]` — a *dev-only* slot, so
grading `demo-N` against it would have been wrong in both directions (`D-M257x-239-2`).

**17 tests**, all three control classes:
* a **mutation control per arm** (A ×2 including the path-artifact negative · B ×2, both directions · C · D
  ×3 including a skill invented at test time that the guard must grade without being edited);
* **all three anti-vacuity paths exit 2, never 0** — no skills, no table rows, no runnable invocations;
* a **real answer key** (`§9` iter-149): the whole live tree rebuilt from `git show 2a0a939:<file>` — this
  iter's own pre-registration commit — which must report exactly **5 bare-`N` + 3 verb-first**. It does.

## Pre-registration — scored 4 confirmed / 1 refuted

| claim | prediction | result |
|---|---|---|
| `P-239-1` name resolution | 16/16 | **CONFIRMED** — control clean |
| `P-239-2` guide paths | 0 missing | **CONFIRMED** — control clean |
| `P-239-3` ≥ 1 argument-contract defect | ≥ 1 | **CONFIRMED**, and understated: 8, and 8 of 8 |
| `P-239-4` retired names, 0 runnable | 0 | **CONFIRMED** — 96 sites, all historical; instrument proven by arm A's mutation control |
| `P-239-5` the defect sits on a `dev-*` skill | `dev-*` | **REFUTED** — 0 of 8. All 8 sit on generic `stack-*` skills |

**`P-239-5`'s refutation is the iter's most useful output, because the reason is structural rather than
luck.** `/dev-up`, `/demo-up`, `/dev-down`, `/demo-down` declare `[N]` — a **bare** `N` is their correct
and only form. `/stack-seed`, `/stack-snapshot`, `/stack-secrets` declare `dev-N|demo-N`, because they are
the *generic* half of the v1.3 "stack party" split and a bare index cannot say which family it means. **The
corpus carried one habit across a boundary the skill set had deliberately drawn** — every defect is a
`stack-*` invocation written in the `demo-*` dialect, and each sits in a document whose own bring-up line
is a genuinely-bare-`N` `/demo-up`.

`recipe-snapshot-world.md` is the sharpest instance, because it states the rule and then breaks it **16
lines later, in prose it wrote itself**: `:28` reads *"`/demo-up N` **or** `/dev-up N` — dev is a peer;
replay works on `dev-N|demo-N`"*, and `:44` then writes the unqualified form. `CLAUDE.md`'s convergence
note does the same thing eleven lines above the fences that ignore it. **`ROUTE-M257x-238` firing a fifth
time, in a fifth input — and this time the contradicted prose is in the same document as the fence.**

## Close — 2026-08-10

**Outcome:** the skill-invocation surface is now censused and fenced; **8 of 8** target-bearing invocations
in the live corpus disagreed with their own skill's declared contract and are repaired, the fence holds
them at zero, and the guard family goes **24 → 25 GREEN / 0 RED**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-239-1` (the declaration is the contract, not the majority spelling) ·
`D-M257x-239-2` (arm D's subject set is derived, not named) · `D-M257x-239-3` (retired names get no arm) ·
`D-M257x-239-4` (2 undeclared presets classified, not repaired) · `D-M257x-239-5` (bare-`N` is correctness,
not safety).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — `stack-core` (pytest 8.4.2 / CPython 3.9.6, Python): the new
`tests/test_skill_invocation_guard.py` runs **17 passed / 0 failed**. No other section run; the whole-section
baseline is unchanged from harden-11's **1,985 passed / 0 failed / 3 skipped** plus these 17. Guard family
re-run at platform reach: **25 GREEN / 0 RED / 0 could-not-check / 5 not-run** (the five that need
`--range`/`--ledger`, unchanged).

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-239-stackseed-sentinel-reload-is-demo-only` → **new.**
  `stack-seeding/cmd/stackseed/main.go:677` reconstructs the stack name as `fmt.Sprintf("demo-%d", n)`
  inside `reloadStackSentinel`, which `shouldReloadSentinel` fires on **any** non-prod stack with `n > 0` —
  so on a `dev-N` the RPC leg is correct (it is offset-keyed) but the `docker restart` **fallback** names
  `demo-N-sentinel-1`. Non-fatal by construction, and found while grading this iter's consequence, not
  while looking for it — the scope-creep tripwire's 3rd line, routed rather than worked.
- `ROUTE-M257x-238-claude-md-fences-are-unmaintained` → **open, and now five-for-five.** Every one of the
  five runnable inputs censused (235–239) found a `CLAUDE.md`-family fence contradicting prose written
  later. The class is confirmed; what is still unbuilt is the sweep that would keep the *fences* current
  the way `/update-knowledge` keeps the prose current.
- `ROUTE-M257x-238-container-vs-native-is-undrawn` → open.
- `ROUTE-M257x-237-critical-env-list-is-unfenced` → open.
- `ROUTE-M257x-236-disclosure-scope-is-document-level` → open.
- `ROUTE-M257x-236-host-is-the-unreliable-witness` → open.
- `ROUTE-M257x-235-fence-scope-is-unread` → open.
- `ROUTE-M257x-235-runnable-block-has-two-halves` → open.

**Lessons:**
1. **A convention that is correct for one family and wrong for its sibling is the hardest kind to see.**
   The bare `N` is not a typo — it is right at 4 skills and wrong at 3, and the corpus applied one habit
   across a boundary the skill set had deliberately drawn. Nothing that greps for a *wrong string* can find
   this; only a check that reads each skill's own declaration can.
2. **Grade the consequence before booking the severity.** "Ambiguous stack target" invites a safety
   framing. Two reads — `ParseStackN` and `isolation.TargetFor` — showed the parse succeeds and the
   prod firewall is untouched, which moved the finding from safety to correctness *and* made the fence's
   docstring honest about what it does not measure.
3. **The instrument was wrong twice and both controls caught it**, for the fifth consecutive iter: the
   comment strip dropped only `#`-leading *tokens*, so `# the main dev stack` contributed `the` as a
   positional argument (2 of the first 10 findings were that artifact); and the hand-listed subject set was
   5 where the derived set is 3. **Derive the subject set, then re-derive after the last edit.**

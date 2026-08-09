---
iter: 238
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-238 — do the documented `npm`/`pnpm` scripts and their ports exist?

## Step 0 — Re-survey

Three iters have censused the three inputs of a runnable instruction on the **Go/compose** side: the `make`
target (235, aligned), the `cd` directory (236, four repairs), the environment variable (237, one repair in
`CLAUDE.md`'s critical list). All three barely touched the **frontend tier**, and that tier is half of what
*"build a working stack"* means — `next-web-app`, `studio-desk` and `ant-academy` are how you see anything.

Their entry points are **`npm run <script>` / `pnpm <script>`**, and a `package.json` `scripts` block is
**enumerable** exactly like a Makefile's targets. Same mechanical shape as iter-235, different toolchain,
and never asked. The ports are enumerable too — a documented `localhost:9100` either matches what the
package/config binds or it does not.

`service_registry_guard` grades **compose** services and their published host ports (12 registry rows vs 7
compose services, 10 host ports). It says nothing about a **natively-run** frontend — and `ant-academy`
runs natively by design (not in docker-compose at all), so its entire surface is outside every existing
fence.

**Active strategy reference:** `TOK-08` — census a mechanical class exhaustively.

## Hypothesis

`ant-academy` is the newest, least-fenced repo (Next.js 16 + Expo, added v1.10b M49, not in `repos.yml`),
and `CLAUDE.md` documents four separate commands for it across two directories with two package managers
(`npm` in `code/`, `pnpm` in `mobile/`). That is the most drift-prone surface in the corpus, and the
denominator lesson from iters 235–237 says the first reading will be wrong about *which* `package.json`
each command addresses.

## Predictions — SEALED BEFORE MEASUREMENT

| id | prediction |
|----|-----------|
| `P-238-1` | ≥ 30 distinct `npm run <x>` / `pnpm <x>` script invocations across `corpus/**` + `CLAUDE.md` + `.claude/skills/**` |
| `P-238-2` | ≥ 4 `package.json` files in the clone set declare a `scripts` block |
| `P-238-3` | ≥ 1 documented script is **not** declared in any clone-set `package.json` |
| `P-238-4` | ≥ 1 miss is in `ant-academy` |
| `P-238-5` | each documented native dev port (3077, 8555, 9100, 9000) is traceable to the repo config that binds it |

## Expected lift

No `N`/`P` reading. Deliverable: the documented-script population with its denominator, the declared-script
sets per `package.json`, the difference, each miss classified (renamed / wrong-package / never-existed /
tooling-only), the port trace, and repair of any miss in a runnable block.

## Phase plan

1. Enumerate every `package.json` with a `scripts` block from the clone set, read via `git show`.
2. Enumerate documented `npm run` / `pnpm` invocations, fenced-vs-prose.
3. Diff, **against the union first** (the iters 235–237 denominator lesson), then per-repo where the
   context names one.
4. Trace each documented port to a config that binds it.
5. Prove the instrument; repair only real misses.

## Escalation conditions

- 0 documented invocations → the regex is the finding (`§9`).
- A miss needs a platform edit → route, never edit.

## Acceptable close-no-lift outcomes

0 misses with a proven-non-vacuous instrument closes the class and refutes `P-238-3`/`P-238-4` on the
seal — as at iter-235, where the refutation was the useful half.

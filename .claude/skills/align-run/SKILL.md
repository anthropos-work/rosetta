---
name: align-run
description: Measure how faithfully a mirror engine reproduces a source engine — run every Alignment DNA gene against both source and mirror, compose a 0–100% alignment score (overall + critical), and surface a per-capability divergence report. Use to check a mirror's fidelity, to drive a mirror toward an alignment gate (the M1 build loop), or to re-score after a source version bump. Pairs with /align-dna (which authors the DNA). Full reference: corpus/architecture/alignment_testing.md.
argument-hint: [dna path + mirror runner, e.g. "dna.json + 'go run ./cmd/clerkrun'"]
---

# Measure alignment of two targets

Run a **mirror** against a **source** across an Alignment DNA and report how faithfully the mirror
reproduces the source as a **0–100% score**. Concepts:
[`corpus/architecture/alignment_testing.md`](../../../corpus/architecture/alignment_testing.md).
Harness: the `rosetta-extensions/alignment/` section (`alignctl`).

> **Where this runs.** This skill + the doc live in **rosetta**; `alignctl` is the
> `rosetta-extensions/alignment/` section. The DNA, goldens, mirror, and
> its runner live in the **mirror's own repo** (e.g. `clerkenstein`) under `stack-demo/`. Run
> this from the mirror repo so the `--runner` relative path and goldens resolve.
>
> Stack-dir rename: `anthropos-demo/` → `stack-demo/` (and `anthropos-dev/` → `stack-dev/` where
> present). Each gitignored `stack-*/` dir spans one full local stack — its platform service repos
> plus its own clone of rosetta-extensions consumed at a pinned tag (`stack-<role>/rosetta-extensions @ <tag>`).
> Same `.agentspace/` disambiguation as align-dna: the only `.agentspace/` path this rename
> introduces is the authoring copy `.agentspace/rosetta-extensions/` (spawn on demand to read/build/
> test tooling, then commit + tag); the mirror-repo-local `.agentspace/` that pins the source-SDK
> clone for goldens is a distinct, pre-existing meaning and stays as-is.

## Prerequisites
- A validated DNA (authored by `/align-dna`) and captured source **goldens** (or use `--source live`).
- The mirror's **runner**: an executable invoked as `RUNNER --target {source|mirror} --dna PATH` that
  prints the outcomes JSON (gene id → `{value, error_class}`).

## Your mission

1. **Run the measurement:**
   ```
   alignctl run --dna PATH --runner CMD --golden-dir DIR \
       [--source golden|live] [--report out.json] \
       [--gate-overall F] [--gate-critical F]
   ```
   - Default `--source golden` replays committed goldens (offline, reproducible). Use
     `--source live` only to compare against a freshly-run source (e.g. right after a goldens refresh).
   - Pass `--gate-overall` / `--gate-critical` to make `alignctl` exit non-zero when the mirror is
     below target (CI / the build loop). With no gate it always exits 0.
2. **Read the report.** It gives the overall score, the separate **critical %**, a per-capability
   rollup (which capabilities diverged), and a **divergence list** — for each diverged gene: the
   operator, the weight, whether it's critical, and the source-vs-mirror diff.
3. **Triage each divergence** — exactly one of:
   - **Fix the mirror** (default): the mirror is genuinely wrong; correct it and re-run.
   - **Fix the DNA gene** (rare, legitimate): the divergence is a field that *should* be ignored
     (a generated id/timestamp) → switch the gene to `normalized` with the right `normalize` path,
     or to `shape`, via `/align-dna`. Do **not** loosen an operator just to make a real bug pass.
   - **Waive** (last resort): the gene can't be matched (e.g. a capability that's genuinely
     un-mockable offline) → document the waiver in the report/decisions; it stays counted against
     the score honestly.
4. **Iterate to the gate (the M1 loop).** Re-run after each mirror fix; watch overall + critical
   climb. The milestone's exit gate (M1: **100% critical / ≥95% overall**) is met when
   `alignctl run --gate-overall 95 --gate-critical 100` exits 0.

## The native test surface (same score, different runner)
The mirror repo also ships a build-tagged alignment suite. `go test -tags alignment ./...` runs the
same comparison core over the same DNA + goldens, reporting one subtest per gene and asserting the
mirror's configured gate — use it in CI. `alignctl run` and the tagged suite agree by construction.

## Surface kinds + dependency-gated runners
`alignctl run` is identical for **behavioural** and **deployment/injection** DNAs — same scoring, same
gate. Two practical notes:
- A mirror's full fidelity is **both kinds**; if it has a deployment DNA (e.g. `clerk-deploy-1`), run that
  gate too, not just the behavioural ones. A green behavioural gate does **not** imply the injection
  artifact deploys (the M3 lesson — see `/align-dna` and `alignment_testing.md`).
- Some runners are **dependency-gated**: they need the consumer's real toolchain to even build/run — e.g.
  the deployment runner needs the consumer's private module (`colony` via `GH_PAT`); the `@clerk/express`
  runner needs `node` + a `node_modules`. Run those where the dependency exists (locally / a configured
  CI), and let the others stay pure. A runner that *won't build* against the consumer's pinned version is
  itself a RED result — that's the deployment contract failing, not a setup nuisance.

## After a source version bump (with /align-dna + M1b)
`/align-dna` diffs + updates the DNA and refreshes goldens for the new source version; then run this
skill to re-score the existing mirror against it. A score drop names exactly which genes the bump
broke — turning a silent break into a flagged, mechanical fix.

## `alignctl` commands this skill uses

| Command | Purpose |
|---|---|
| `alignctl run --dna P --runner CMD --golden-dir D [--source golden\|live] [--report out.json] [--gate-overall F] [--gate-critical F]` | run the mirror, score it, report divergences; exit non-zero if a gate is unmet. |
| `alignctl capture --dna P --runner CMD --golden-dir D` | (via /align-dna) refresh source goldens before a live re-score. |

## Done when
The score is read and every divergence is triaged (fixed / DNA-adjusted / waived-with-note). For a
build loop, done when `alignctl run --gate-overall <X> --gate-critical <Y>` exits 0.

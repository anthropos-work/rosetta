---
iter: 240
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
---

# iter-240 — the TOOLCHAIN VERSION: the sixth input of a runnable instruction

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.*
**Route:** `ROUTE-M257x-238-claude-md-fences-are-unmaintained`, now five-for-five (iters 235–239).

## Step 0 — re-survey

Iters 235–239 censused five inputs a runnable instruction needs — `make` targets · `cd` targets ·
environment variables · `npm`/`pnpm` scripts · slash-commands — and **every one came back with a defect.**
A sixth input has never been read, and it is the one that decides whether a person gets a working stack
**before the first command runs**: the **toolchain version**.

A wrong `cd` produces a clear error. A wrong required-Go-version produces a build failure inside a
container, or — worse — a host tool that silently will not compile. `corpus/ops/setup_guide.md:124` states
**"Go 1.25.x"** for the rext host tools; `.claude/skills/dev-for-dummies/reference.md:99` repeats it;
`:89` states **"Node ≥ 24"** and quotes `next-web-app`'s `engines.node >=24`;
`corpus/ops/demo/frontend-tier.md:155` quotes ant-academy's `engines.node >=22`. Every one of those is a
claim about a file, and no guard reads any of them.

**Measured substrate, before grading:** `stack-demo/app/go.mod` declares **`go 1.26.4`** and
`stack-demo/sentinel/go.mod` **`go 1.26.0`** — but those build inside Docker, so they are NOT the host
claim's source of truth. The host claim is about **`rosetta-extensions`' own sections**, which is what
must be read.

⚠️ **The naive selector is again mostly instrument.** A tool-name-adjacent-to-a-number regex over the live
corpus returns **471 distinct (tool, version) pairs across 1,108 sites**, and its top hits are
`go 14` / `docker 48` / `go 123` — **line numbers in `file:line` citations**, not versions. Fifth
consecutive iter where the naive population would have been dominated by its own instrument. The graded
population is therefore the **explicit prerequisite claims of the bring-up path**, enumerated by hand and
listed in this file, not a regex sweep.

## Pre-registered claims — SEALED IN THIS COMMIT, before any grading

- **`P-240-1`.** The corpus's host-Go claim (**"Go 1.25.x"**, two sites) is **stale** — at least one
  `rosetta-extensions` section's `go.mod` requires a Go newer than 1.25. **Predict: REFUTED claim / ≥ 1
  section over 1.25.**
- **`P-240-2`.** The quoted `next-web-app` `engines.node >=24` is verbatim-correct. **Predict: confirmed.**
- **`P-240-3`.** The quoted ant-academy `engines.node >=22` is verbatim-correct. **Predict: confirmed.**
- **`P-240-4`.** At least one bring-up document states a host toolchain version that no longer matches its
  source. **Predict: ≥ 1.**
- **`P-240-5`.** The Go version the platform's **own Docker build** pins (`app`'s Dockerfile `FROM golang:…`)
  is stated **nowhere** in the live corpus. **Predict: 0 sites** — an undocumented prerequisite is invisible
  until it fails.

## Phase plan

1. Seal this pre-registration (this commit).
2. Read each source of truth directly — `go.mod` per rext section, `package.json` `engines`, Dockerfile
   `FROM` — never a doc quoting another doc.
3. Grade item by item; repair toward the source; classify what is under-declared rather than wrong.
4. Extend the fence if and only if the class is mechanically decidable at a stated denominator.

## Escalation conditions

If the host-Go claim is stale in the direction that **blocks a build** (the corpus asks for a Go older than
a section requires), that is a working-stack blocker and the repair takes priority over any fence work.

## Acceptable close-no-lift outcomes

Every version claim resolving is a complete iter: it would make the toolchain the **first** of the six
runnable-instruction inputs to come back clean, and that is a measurement about where the corpus rots.

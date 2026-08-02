---
iter: 39
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-02
---

# iter-39 — `MEASURE-M257x-iter39-clause5-fifth-pass`

**Active strategy reference:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`).

## Cluster / target identified

Clause 5 is the **only** open gate clause (1, 2, 3, 4 all hold; clause 2 closed at iter-37 at `30/0/0`).
Four passes have run — iters 21/33, 34, and 38 — returning **25, 13, 11 and 17** blockers. Every one was
fixed. **A clause is met by a READING that returns zero, not by a repair that clears its own findings**;
iters 33, 34 and 38 each refused to claim it on that ground and this iter holds that line.

Routed forward by iter-38 as `MEASURE-M257x-iter39-clause5-fifth-pass`.

## Hypothesis

None about the corpus's content — this is a **measurement**, and its value is destroyed by predicting what
it should find. The hypothesis under test is procedural: **that a full read of all 40 in-scope files under a
partition that shares no boundary with iters 33/34/38 returns a countable blocker set**, and that the
count/location distribution can be compared against a prediction registered before any auditor reports.

## Scope — all 40 files, no narrowing

`corpus/services/**` (29 docs) + `corpus/architecture/**` (11 docs) = **40 files / 8 674 lines**.

**Read in full, top-to-bottom.** Two narrowing strategies are already disproven in this milestone:

- **By search term** — iter-21. A term-scoped audit's 11 → 5 → 2 convergence was the audit exhausting its own
  grep vocabulary; a full read then found 53. §5 *"An audit scoped by SEARCH TERMS measures the terms, not
  the corpus."*
- **By file set** — iter-38. Run 19 routed *"scope to the 9 changed files"* on iter-34's real 9× density
  measurement. iter-38 ran wide anyway: **11 blockers in the 8 repaired files (1.375/file) vs 6 in the 32
  untouched (0.19/file)**. The density ratio reproduced (~7.3×); the counts did not. Scoping **would have
  found 11 of 17 and declared the rest clean** — including the two most consequential claims in the corpus,
  both in `ai_architecture.md`, a file two prior full-read passes had already cleared.

> **§5 rule 18 licenses WEIGHTING, not NARROWING.** This iter weights by giving the 13 files iter-38 repaired
> **double coverage** (once in full by their partition owner, once as a diff by a dedicated adversarial
> auditor) — while still reading all 40 in full.

## Partition — snake draft by size, sharing no boundary with iters 33/34/38

Files sorted by line count descending, dealt A→F then F→A repeatedly. This deliberately breaks **topical**
adjacency, which is the property that produced iter-38's strongest signal (two auditors refuting the same
false claim from two different files — impossible when one auditor owns both).

| auditor | lines | files |
|---|---|---|
| **A** | 1566 | external_services · chronos · cms · academy-backend · coursebuilder · skiller · TEMPLATE |
| **B** | 1499 | ai-readiness · backend · security_compliance · ai-labs · messenger · customerio-sync · architecture/README |
| **C** | 1473 | alignment_testing · hiring · jobsimulation · sentinel · clerk-integration · services/README · db-backup |
| **D** | 1364 | studio-desk · architecture_overview · shared_libraries · roadrunner · next-web-app · gotenberg · intelligence |
| **E** | 1367 | studio-room · clerkenstein · ai_architecture · storage · askengine · dependency_map |
| **F** | 1405 | service_taxonomy · ant-academy · graphql-wundergraph · platform-migration-status · skillpath · frontend_architecture |
| **G** | (diff) | adversarial diff-read of `git diff a98692b..643ed2b -- corpus/` — iter-38's 13 repaired files |

Splits that matter, verified against the draft: the merged/archived family (cms · jobsimulation · roadrunner ·
skiller · skillpath · intelligence · chronos) lands in **four** different hands; `ai_architecture` (E) and
`security_compliance` (B) stay split, preserving the cross-check that caught the EU-AI-Act claim;
`backend` (B) / `architecture_overview` (D) / `service_taxonomy` (F) — the three big "what is `app`" docs —
split three ways; `studio-desk` (D) / `studio-room` (E) split; `academy-backend` (A) / `ant-academy` (F) split.

## PRE-REGISTERED PREDICTION — written before any auditor report is read

Registered so the measurement can refute it. iter-38's prediction was refuted on **both** count and location
and that refutation was the iter's most valuable output.

**Count: 10–16 blockers.**

**Location: 7–11 in the 13 files iter-38 repaired; 3–6 in the 27 it never opened.**

Derivation:
- The "blockers found in the *immediately prior* pass's repaired text" figure is the stablest number in this
  series: **9** (iter-34, in iter-33's repairs) and **11** (iter-38, in iter-34's repairs) — and it held at 11
  *despite* iter-34 having run the mandatory adversarial half. So an adversarial half does not drive it to 0.
  13 repaired files × the 1.375/file rate → ~18; the observed absolute has been 9–11, so 7–11.
- The untouched-set rate measured 0.19/file at iter-38. 27 untouched files × 0.19 ≈ **5**.

**Two named specific predictions:**
1. At least one blocker in `security_compliance.md` **or** `ai_architecture.md` arising from the AI-Act
   **retraction text itself** — a retraction is a new derived claim, and iter-38's own adversarial half
   already caught one dead-field citation inside it.
2. At least one blocker in `hiring.md`, repaired twice and defective after both.

**Consequent prediction: clause 5 does NOT close this iter.** Stating that in advance is what makes a
zero-blocker result meaningful if it happens — and what makes a non-zero result a confirmation rather than a
disappointment.

## Phase plan

- **Phase A — read.** Seven auditors in parallel (A–F full-read partitions + G diff-read). Each verifies every
  falsifiable claim against the platform clones at origin `2adcf71`, the live `demo-1` Postgres, or
  `docker-compose.yml`/`repos.yml`. Blocker = a claim a reader would act on that is FALSE, or a load-bearing
  `file:line` anchor that does not resolve. Line drift / undercounts / omitted list members are **minors**
  (*"YELLOW with 0 blockers"* admits them).
- **Phase B — reconcile + fix.** Enumerate to `evidence/blocker-ledger.md` with anchors; fix each with an
  exactly-once anchored edit.
- **Phase C — adversarial pass over iter-39's OWN corrections** (§5 rule 18(a); mandatory, never once clean in
  three attempts: 24% → 2 → 6). **Diff-read the sweep, not just its anchors** — all six of iter-38's
  self-inflicted defects were in surrounding prose, over-correction, or mechanical damage, while every
  `file:line` anchor it introduced resolved correctly.
- **Phase D — re-measure.** Corpus guards ×5. Grade clause 5 against the reading, not the repair.

## Escalation conditions

- A platform commit landing mid-iter that invalidates an alignment attempt → **re-scope trigger occurrence 2**
  → STOP and escalate (occurrence 1 already stands).
- A finding requiring a **platform-repo edit** → route forward; the v2.8 zero-platform-edit constraint is binding.
- A finding that is a **legal** question (the AI-Act classification) → route to an owner outside this
  milestone; do NOT let this iter settle it by asserting a classification.

## Acceptable close-no-lift outcomes

A pass that returns a non-zero blocker count is a **complete** iter: the reading is the deliverable. The iter
closes `closed-fixed` when the enumerated blockers are fixed and the guards are green, whether or not clause 5
flips. Clause 5 flipping is not this iter's success criterion — an honest count is.

**Type:** tik — under `TOK-05`, step 2 (citations), clearing the last **named** unrepaired class
before step 4. *Do not take the graded read while a known class is unrepaired.*

# iter-72 — the mainline class dissolves, and the guard cannot see 142 citations

## Phase A — `FIX-M257x-iter58-mainline-shift`, re-derived

The route has read **"21 of 22 outstanding"** since iter-59, and the rule underneath it changed
three times in iters 69–71. Re-derived at `app` `origin/main` `9d00a313`:

| | n |
|---|---|
| distinct `main.go` / `app/main.go` citations | **66** |
| graded at a ref **their own block names** | **37** |
| block naming **two or more** resolvable refs → graded at the default | **29** |
| **out of range or absent at the ref they are graded at** | **0** |

**Zero structurally broken.** The class dissolves exactly as B2 did at iter-69 and for the same
reason. **The route closes with a derived verdict**, and its "21 outstanding" joins iter-69's "64"
and iter-70's "23" as a carried number that did not survive re-derivation.

**A self-inflicted measurement bug, recorded rather than quietly fixed.** The first run printed
*"graded at a ref their own block names: **0**"*, flatly contradicting the guard's own
`block-pinned x31`. The script called `live.pop()` and then tested `len(live) == 1` on the mutated
set. **Two instruments disagreeing is a finding, and the one that agrees with nothing is usually
the new one** — the contradiction was visible in one glance precisely because the other instrument
had already been mutant-verified.

## Phase B — probing the guard's reach, and the eighth limit

| probe | result |
|---|---|
| `` `main.go:1187` `` matches `_QUALIFIED`? | **False** |
| `` `app/main.go:1187` `` | True |
| `` `internal/coursebuilder/bedrock.go:98` `` | True |
| `resolve('main.go')` from `backend.md` | **None** |
| `resolve('app/main.go')` from `backend.md` | the file |
| `stack-demo/backend/` exists? | **No** — the compose SERVICE is `backend`, the REPO is `app` |

`_QUALIFIED` requires a `/` in the path or a `.md` suffix, so a bare `<name>.<ext>:N` **never
reaches `resolve()` at all**. Derived over the live corpus: **142 distinct citations are outside the
guard's reach entirely**, led by **41 `docker-compose.yml:N`** and **32 `up-injected.sh:N`** — the
two most-cited artifacts in the ops corpus.

And the second half is independent: `resolve()` already carries a repo-relative rule for service
docs (`root / doc.stem / cited`), but for `backend.md` that is `stack-demo/backend/`, which does not
exist. **Two gaps, either of which alone would hide the class.**

**Eighth reach limit of this milestone, and by far the largest.** The pattern is now worth stating
plainly: **every one of the eight was found by reading or probing, never by a GREEN verdict.** A
fence reports on the class it can see; its silence about everything else is indistinguishable from
health.

## Phase C — the fix, designed and routed rather than half-landed

`FENCE-M257x-iter72-bare-citation-reach`:

- widen `_QUALIFIED` with a third alternative for a bare `<name>.<codeext>:N`;
- **derive the doc-stem → clone mapping from compose** rather than listing it — `docker-compose.yml`
  gives `backend` a `build.context: ../app`, so service→repo is an artifact fact, and
  `platform_predicate_guard.parse_compose` is already an importable tested primitive
  (`D-M257x-59-2` sanctions exactly this reuse);
- keep `AMBIGUOUS_BASENAMES`' discipline — `main.go` exists in **seven** clones, so a bare basename
  must resolve **through the doc's own service**, never by a tree-wide basename search. That is the
  over-match the guard's own docstring records as *"134 findings, essentially all of them ports."*

**Not landed here**: it is a third line in a two-line iter and it will turn the guard RED on real
sites, which is a repair pass of its own. Same disposition iter-68 used for two boundary defects it
measured unreachable — **recorded, with the proof, rather than patched on speculation.**

## Phase D — gates

| gate | result |
|---|---|
| five corpus guards | **all OK** — alignment · anchor (`ref chosen by default x57, block-pinned x31, no-clone x30, ambiguous x12`) · predicate (8 assertions) · markdown-structure · corpus-index |
| suites | **not re-run — zero code touched**, in either repo. This iter changed plan documents only; `git status` is the evidence. iter-71's runs stand: `stack-core` 762/1F by identity · `stack-injection` 332 OK · `dev-stack` 151 OK solo · `demo-stack` 1048/7F by identity |

## Close — 2026-08-04

**Outcome:** `FIX-M257x-iter58-mainline-shift` closes with a **derived** verdict — **66 distinct
mainline citations, 37 graded at a ref their own block names, 29 in ambiguous blocks, and ZERO out
of range or absent.** Its "21 of 22 outstanding" joins iter-69's "64" and iter-70's "23" as a
carried number that did not survive re-derivation. Probing the guard's reach in that class then
found the **eighth and largest reach limit of this milestone**: `_QUALIFIED` requires a `/` or a
`.md`, so **142 distinct bare `<name>.<ext>:N` citations never reach the resolver at all** — 41 of
them `docker-compose.yml:N`, 32 `up-injected.sh:N` — and the resolver's service-doc rule would miss
them anyway, because `backend.md`'s doc-stem maps to `stack-demo/backend/` while the clone is `app`.
The fix is **designed and routed with both proofs**, not half-landed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (4 tiks of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-72-1` (the mainline route closes: 66 distinct, 0 structurally broken; plus
the self-inflicted `live.pop()` measurement bug), `D-M257x-72-2` (the eighth reach limit — 142
citations outside the regex, and a doc-stem→clone gap behind it; fix designed and routed).
**Side-deliverables:** none.
**Routes carried forward:**
- **`FENCE-M257x-iter72-bare-citation-reach`** — the 142. Design, both mechanical proofs and the
  anti-over-match constraint are in `decisions.md` `D-M257x-72-2`. **A prerequisite for the graded
  read** — and unlike iter-69's routed prerequisite, which iter-70 falsified, this one is proven by
  two probes rather than by a count.
- `CHECK-M257x-iter71-ambiguous-blocks` — now with a number for the mainline slice: **29 of 66**.
- Unchanged: `RF-M257x-iter71-run-returns-a-tuple` · `FENCE-M257x-iter70-line-or-port` ·
  `CHECK-M257x-iter70-studio-room-lines` · `FIX-M257x-iter53-union-set` (**PENDING USER DECISION**)
  · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED**) ·
  `CHECK-M257x-iter38-ai-act-classification` (owner outside this milestone) ·
  `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13.
- **Closed here:** `FIX-M257x-iter58-mainline-shift`.

**Lessons:**

1. **Four carried numbers, four re-derivations, four collapses** — 64→5, 23→4, 21→0, and the class
   size 86→96→105→109 growing under repair. In this milestone a routed count has never once
   survived contact with a re-derivation. §5 rule 32 should be read as: *a number you did not
   measure this iteration is a hypothesis.*
2. **Every reach limit here was found by reading or probing; none by a GREEN verdict.** Eight for
   eight. The corollary is uncomfortable and worth keeping: the guards' GREEN is evidence about the
   class they reach and evidence about nothing else, so **the reach line is the load-bearing
   output**, not the verdict.
3. **Two instruments disagreeing is a finding.** The `live.pop()` bug announced itself as a flat
   contradiction with a mutant-verified guard, which is the cheapest possible signal — and only
   available because the other instrument existed and printed its reach.
4. **Design the fix in the iter that finds the gap, even when you do not land it.** The route
   carries both proofs and the anti-over-match constraint, so the next iter starts from evidence
   rather than from a sentence — which is what iter-70 wished iter-69 had left it.

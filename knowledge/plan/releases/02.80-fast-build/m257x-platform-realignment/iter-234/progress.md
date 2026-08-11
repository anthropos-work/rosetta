**Type:** tik — under `TOK-08` (census a mechanical class exhaustively).

# iter-234 — do the corpus's `file:NN` anchors LAND on the text they quote?

## What was measured

Population derived from `corpus/**` + `CLAUDE.md`, restricted to the **six live platform repos** a stack
actually builds, read at the **`stack-demo` clone set** (harden-54: the tree a demo builds, not
`origin/main`). Instrument: `.agentspace/scratch/work-m257x/cite234.py`.

| repo | anchor sites |
|---|---|
| `app` | 230 |
| `platform` | 165 |
| `sentinel` | 7 |
| `next-web-app` | 4 |
| `ant-academy` | 2 |
| `studio-desk` | 1 |
| **total** | **409** |

Gradable by co-quotation (a backticked content literal beside the anchor): **278**. Ungradable: **131** —
stated, not hidden.

## The instrument needed three corrections, and each moved the headline

Recorded because the sequence *is* the finding — the first two readings would have been published as
corpus defect rates:

| correction | why | exact-LANDS after |
|---|---|---|
| (none — naive) | every backticked token on the line counts | 63.6 % |
| exclude **metadata** literals (shas, bare repo/service names, anchors) | `` `app` `` appears in every Go file (false LANDS); a sha appears in none (false MISMATCH) | 55.7 % |
| exclude **pointer** literals (bare paths/filenames) | grading *"does the string `repos.yml` appear inside `sentinel/go.mod`"* is a category error | 56.8 % |
| associate each literal with its **nearest anchor**, not its line | a wide table row carries several anchors; line-pairing manufactures mismatches from a neighbour's literal | 56.8 % |

**Valid-instrument subset** — single-anchor line **and** no negation marker: **93 of 278 (33.5 %)**.
Exact-LANDS on it: **53.8 %**; LANDS-or-NEAR **55.9 %**; DRIFT 16.1 %; MISMATCH 25.8 %.

## The verdict, and it inverts the number

**Five of five unambiguous MISMATCH candidates, hand-verified against the clone, are the corpus being
exactly right.** Each failed for a different, nameable reason — the taxonomy is now in
`platform-alignment.md` § 5:

| # | site | verdict on hand-verification |
|---|---|---|
| 1 | `ai-readiness.md:98` → `enum/organization_settings.go:47` | **LANDS on the exact line.** Corpus drops the Go **type name** from the declaration |
| 2 | `backend.md:220` → `terraform/locals.tf:6` | **LANDS.** Corpus quotes the **reference** form `local.project`; the file **declares** `project` |
| 3 | `update_guide.md:96` → `Makefile:31` | **LANDS.** Corpus quotes the **invocation** `make pull`; the Makefile declares the **target** `pull:` |
| 4 | `backend.md:222` → `terraform/main.tf:58` | **Correct NEGATIVE claim** — the sentence continues *"appears in **no** `.tf` file in any clone"* **on the next markdown line**, invisible to a line-scoped polarity detector |
| 5 | `coverage-protocol.md:637` → `readiness.go:308` | Literal is a **fix-id** — corpus vocabulary that cannot appear in source by construction |

So the **53.8 % is a floor on demonstrable agreement, never a ceiling on correctness**, and the complement
is **not** a defect rate. The co-quotation containment census is **REFUSED as an instrument for this
class**, in writing, so a future iter does not spend a run "repairing" correct claims.

## What the census DID establish and keeps

Grading every site at **two clocks** — the clone HEAD a demo builds, and the sha the site itself names —
found **15 sites that miss at clone HEAD and LAND at their own stated sha**. Those are **correct and
dated**, not wrong (`external_services.md:203/:567`, `shared_libraries.md:130/:143`,
`service_taxonomy.md:108`, `content-stories-routes.md:351`, `backend.md:357`, …). Harden-54 in the
opposite direction: when a site names its ref, the ref it names is the tree that grades it.

## Instrument non-vacuity

The census did **not** return zero and therefore does not owe the `§9` zero-proof: it returns **158 LANDS**
at clone HEAD over 278, so the positive side demonstrably fires. The negative side is proven
**over-firing** by the 5/5 hand-verification above — both sides characterised in the same run.

## Seal grading — `3fa583c`, sealed before any measurement

| id | prediction | outcome |
|----|---|---|
| `P-234-1` | ≥ 150 anchor sites into live repos | **CONFIRMED — 409** |
| `P-234-2` | ≥ 60 gradable by co-quotation | **CONFIRMED — 278 (93 on the valid subset)** |
| `P-234-3` | exact-LANDS ≥ 70 % | **REFUTED — 56.8 % / 53.8 %** |
| `P-234-4` | ≥ 10 anchors MISS at the cited line | **CONFIRMED as a count (120 non-LANDS) — but the LABEL "miss" is WITHDRAWN**; the class is not misses |
| `P-234-5` | ≥ 1 real content mismatch | **REFUTED on the verified sample — 0 of 5.** Not proven zero corpus-wide; proven the instrument cannot find one |
| `P-234-6` | ≥ 1 site misses at clone HEAD but lands at its own sha | **CONFIRMED — 15** |

**4 confirmed · 2 refuted**, and both refutations are about the *instrument*, not the corpus — which is
the seal doing its job.

## Guard family

`24 GREEN · 0 RED · 0 could-not-check · 5 not-run` (`--platform stack-demo/platform --allow-not-run`) —
identical to the run-27 baseline. Scoped run without `--platform` is 19 GREEN / 10 not-run; the reach is
stated because *a verdict without its reach is not a verdict*.

## Close — 2026-08-10

**Outcome:** the citation ladder's top rung is measured and the instrument for it is **refused in
writing**. 409 anchor sites; 53.8 % exact-LANDS on the cleanest subset; **5 of 5 hand-verified
"mismatches" were the corpus being right**, in five distinct shapes now named in the protocol doc. The one
durable positive: **15 sites land at the sha they name and miss at clone HEAD** — correct-and-dated, two
clocks, not interchangeable.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-234-1` (the containment census is refused, not tuned), `D-M257x-234-2` (no corpus
prose was repaired — 0 defects found, so 0 edits made).
**No `N`/`P` movement is claimed** — this iter took no graded seat.

**Suite state at close** — no pytest section run; this iter changed no rext code. Corpus prose changed in
one file (`platform-alignment.md` § 5) and the full guard family was re-run at platform reach: 24 GREEN /
0 RED.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-234-polarity-is-a-sentence-property` → **new.** Every claim instrument in the family reads
  **lines**. The corpus's absence-claims routinely wrap, putting the negation on the next line. Any future
  polarity-aware guard must parse sentences, not lines.
- `ROUTE-M257x-234-two-clocks-is-fenceable` → **new.** The 15 correct-and-dated sites are a real, cheap,
  *non*-paraphrase-dependent signal (line-existence + stated-sha resolution only). Unlike containment, this
  one **is** mechanical and could become a guard arm.
- All prior routes → open, unchanged.

**Lessons:**
1. **A citation binds a CLAIM to a LOCATION; it does not promise a TRANSCRIPT.** Paraphrase is the correct
   scholarly form. An instrument that treats citation as quotation reports the best-written sentences as
   the worst.
2. **Hand-verify before publishing a rate.** Three successive instrument corrections each looked like the
   last one; only reading five actual files showed the residual was *entirely* artifact. A rate published
   after correction #3 would have been a 46 %-broken-corpus headline.
3. **Refusing an instrument is a deliverable.** The negative result is written into the protocol doc
   precisely so the next agent does not rediscover it at the cost of a run — and does not "repair" 24
   correct claims into wrong ones.

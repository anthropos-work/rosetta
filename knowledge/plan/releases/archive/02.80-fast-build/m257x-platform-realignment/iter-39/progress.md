**Type:** tik

# iter-39 — the fifth clause-5 pass: 37 blockers, and a prediction refuted by 2.3×

## The measurement

Seven auditors, **40 files / 8 674 lines**, every in-scope file read top-to-bottom under a partition
(size-sorted snake draft) sharing no boundary with iters 33/34/38, plus a dedicated adversarial diff-reader
over iter-38's 13 repaired files. Combined verification volume, self-reported: **~805 exact citations**
checked against the platform clones at origin `2adcf71`, the live `demo-1` Postgres, the read-only prod
taxonomy capture, or `docker-compose.yml`/`repos.yml`.

**Result: 37 unique blockers** (38 raw; `hiring.md:86` was found independently by two auditors). Enumerated
with anchors in `blocker-ledger.md`. All 37 fixed.

## The pre-registered prediction was refuted on count AND on location — in both directions

Written into `overview.md` before any auditor reported: *"10–16 blockers; 7–11 in the 13 files iter-38
repaired; 3–6 in the 27 it never opened."*

| | files | blockers | per file | predicted |
|---|---|---|---|---|
| repaired by iter-38 | 13 | **~25** | 1.9 | 7–11 |
| never opened by iter-38 | 27 | **~12** | 0.44 | 3–6 |
| **total** | 40 | **37** | 0.93 | **10–16** |

Both strata came in at roughly double their predicted ceiling. The repaired-vs-untouched density ratio is
**~4.4×** — down from iter-38's 7.3× and iter-34's 9× — while the absolute count rose **11 → 17 → 37**.

**The honest reading, and it is a caution against this milestone's own headline series.** These numbers came
from five different instruments: iter-33 ran 5 auditors, iter-38 ran 6, iter-39 ran 7 and briefed them with
the accumulated §5 rules plus each file's own repair history. **25 → 13 → 11 → 17 → 37 is not a corpus
trend; it is four instrument changes.** Nothing here licenses a claim that the residual is converging, and
nothing licenses a claim that it is growing. **The series is not a measurement of the corpus. It is a
measurement of successive audits, and it has been quoted as the former.**

## Ten blockers were in seven files no prior pass ever flagged

`studio-desk.md` (2), `sentinel.md` (2), `graphql-wundergraph.md` (2), `studio-room.md` (1),
`clerkenstein.md` (1), `alignment_testing.md` (1), `cms.md` (1). Every one had been read in full at least
twice and passed twice. This is iter-38's `ai_architecture.md` finding reproduced at scale: **what changes
between passes is the partition, not the diligence.**

Three of them are the kind a reader acts on and loses time to:

- **`studio-room.md` documented a `gen.py --template` flag that does not exist** — and because the parser
  uses `parse_known_args`, a stray `--template` is **silently absorbed**, so all four documented commands
  *succeed* and generate something unrelated to what was asked. Worse than a hard failure.
- **`sentinel.md` listed a `manager` Casbin role that does not exist.** The real fourth role is
  `content_creator`. Granting `manager` yields a membership with **no policy rows at all** — the silent-403
  mode this corpus warns about elsewhere.
- **`graphql-wundergraph.md`'s smoke test promised `{ __typename }` → `{"data":{"__typename":"Query"}}`.**
  It returns HTTP 200 with `unknown viewer: Forbidden`, and the platform pins **that exact query** in its own
  regression test. The documented "healthy" output is the one output the platform guarantees you cannot get.

## Four repairers corrected the hand-off rather than applying it — the iter-22 rule earning its keep

Each repairer was required to re-derive its correction against platform source before editing. Four did
**not** apply what they were given:

- The claim *"every `default:` arm resolves to `gpt-4.1`"* is false for the two Anthropic arms.
- *"The out-of-range enum row fails at the GraphQL marshal"* is **also** false — both gqlgen bindings are
  passthroughs. **Nothing rejects the value anywhere on the path**; it renders wrong rather than vanishing.
- *"A stray `--template` is always silently swallowed"* — not in `--blueprint` mode, which hard-errors.
- The `31 of 135` multi-tenancy figure was **REFUSED and routed** rather than changed, on the ground that its
  home derivation lives in a file the repairer did not own and flipping one site would manufacture a
  contradiction. That is the correct call and it is the fence that has been wrong four times.

## The adversarial pass: 8 self-inflicted, and the defect class has SHIFTED

Fourth consecutive non-clean adversarial pass (24 % → 2 → 6 → **8**). 122 hunks read in full with ±20 lines
of context; ~95 introduced anchors resolved.

**Three prior passes were dominated by mechanical damage. This one found ONE mechanical defect and FIVE
cross-file DRIFT defects** — a claim corrected in the file its owner held while the identical claim survived
in a twin file owned by somebody else. The corpus still said "2 → 1" in seven places after `backend.md` was
corrected to "3 → 1"; still published the EU-first ladder in `ai_architecture.md` after `external_services.md`
refuted it; still told readers to run `--template` in three places after `studio-room.md` warned it does not
exist — one of them **inside a mermaid edge label**, the blind spot this milestone has now been bitten by
three times.

Two of the eight are worth naming for their shape:

- **An over-correction with the same signature as the one it was fixing.** A repairer wrote *"Three things —
  and **only three** — can send a request outside the EU," derived by reading `jobsimulation/ai/ai.go`. There
  is a fourth: a sequence with no `ai_vendor` defaults to `openai` in the **cms content layer**, which the
  callee's dispatch file structurally cannot see. A universal quantifier is only as wide as the file it was
  derived from.
- **A citation to a commit that never merged.** `bba862f` was cited as switching `sse_post` → `ws`;
  `merge-base --is-ancestor` returns rc=1 and it lives only on `origin/feat/use-web-socket`. Mainline never
  carried `ws`. `git log` finding a commit is not evidence the commit is in your history.

I also caught one over-correction myself before the adversarial pass ran, and it is instructive: a repairer
concluded *"the compose service **always** built from `Dockerfile.dev`"* from `git log -S
"graphql-wundergraph/Dockerfile"` returning exactly two commits. That **prefixed** path only existed after
`67ba772` — so the search was structurally blind to the earlier era, when the block carried **no
`dockerfile:` key at all** and Docker silently defaulted to the production file. **An absent key is invisible
to every search for its value.** The corrected text now carries the three-era table and the caution.

## Two pre-existing mechanical defects, found by a regression grep rather than by reading

After the sweep I grepped the whole in-scope tree for orphaned continuation lines (`^ [A-Za-z]`) and found
**two broken blockquotes that predate this iteration** and survived four full-read audits —
`external_services.md` (a sentence split across a non-quoted line) and `clerkenstein.md` (a continuation
orphaned from its subject when a blockquote was inserted mid-sentence, leaving *"It is identity-agnostic"*
with no antecedent). Both fixed; the tree is now clean under that grep.

**They survived four passes because a broken sentence asserts nothing false.** Every auditor in this
milestone has been pointed at *claims*, and mechanical damage is invisible to a claim-checker. A two-second
grep found what ~800 citation checks did not.

## Clause 5 — NOT MET, and not close

**A clause is met by a READING that returns zero, not by a repair that clears its own findings.** This pass
returned 37, then 8 more from its own adversarial half, then 2 more from a regression grep — **47 found and
fixed**. iters 33, 34 and 38 each refused to claim the clause on exactly this ground; so does this one, and
with far less ambiguity than any of them.

## Close — 2026-08-02

**Outcome:** clause-5 fifth pass measured at **37 blockers** (+8 self-inflicted +2 pre-existing mechanical;
**47 closed** across 20 files, +738/−249 lines). Pre-registered prediction refuted on count by **2.3×** and on
location in both directions. Established that the 25→13→11→17→37 series measures **instruments, not the
corpus**. Clause 5 stays open.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (4 of 5 — clauses 1, 2, 3, 4 hold; clause 5 outstanding)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: y (session budget, not the 5-tik cap) — (6) protocol-stop: n — Outcome: exit-5
**Decisions:** D-M257x-39-1 (read all 40 again; weight by repair history, never narrow) · D-M257x-39-2 (repair by CLAIM, not by FILE — the new §5 rule 19 candidate) · D-M257x-39-3 (refuse the `31 of 135` change and route it, rather than manufacture a cross-file contradiction) — see `decisions.md`
**Side-deliverables:** two pre-existing broken blockquotes repaired corpus-wide, found by a net-new regression grep that is now worth keeping as a cheap standing check.
**Routes carried forward:**
- `MEASURE-M257x-iter40-clause5-sixth-pass` — the confirming reading. **Re-partition again.** Do NOT
  inherit iter-39's snake draft.
- `FIX-M257x-iter39-claim-scoped-repair` — **repair by claim, not by file.** 5 of 8 self-inflicted defects
  were the direct consequence of file-scoped ownership over a claim-scoped problem. Before editing, grep the
  whole corpus for the claim and fix every instance in one pass.
- `DOC-M257x-iter39-ops-collateral` — the refuted **"60K skills / 18K roles"** figures survive in ~12
  `corpus/ops/**` files, `.claude/skills/stack-snapshot/SKILL.md`, and **`CLAUDE.md:189,219`** — the repo's
  top-level instruction file, the highest-propagation site in the tree. Out of clause-5 scope; in scope for
  honesty. `--template` also survives in `ops/run_guide.md:347` and `ops/setup_guide.md:595`.
- `CHECK-M257x-iter39-tenancy-fence-off-by-one` — `architecture_overview.md:288`'s "31 of 135" is arguably an
  undercount by one (`organization.go` declares its own org-filtering `Policy()`), but the home derivation is
  in `security_compliance.md` and the **denominator** is itself ambiguous (135 by `grep -l ent.Schema`, 112 by
  counting embedded declarations). **This fence has been wrong four times in both directions — do not change a
  number here without settling both conjuncts AND the denominator.**
- `CHECK-M257x-iter39-archived-repo-contradicts-corpus` — `graphql-wundergraph/CLAUDE.md:39` asserts the
  opposite of the now-verified compose build path, in a commit titled *"correct the compose build path"*. A
  fencing sentence is in place; the general risk (an archived repo's own docs re-corrupting the corpus) is not.
- `DOC-M257x-iter39-minors` — ~60 minors with exact anchors across nine reports.
**Lessons:**
- **A series of measurements taken with changing instruments is not a trend.** This milestone has quoted
  25→13→11→17 as though it described the corpus. It describes the audits. Say which instrument produced each
  number, or the series will keep being read as convergence.
- **Repair by claim, not by file.** A partition that is right for *reading* (disjoint ownership, which is
  what produces independent double-finds) is wrong for *repairing*, because a claim does not respect a file
  boundary. Half-repairing a uniformly-wrong corpus is worse than leaving it: it teaches the reader that the
  corpus contradicts itself, and the next auditor spends its budget adjudicating instead of measuring.
- **A universal quantifier is only as wide as the file it was derived from.** Two of this pass's defects —
  "only three EU exits" and "always Dockerfile.dev" — were true of the file that was read and false of the
  system. When writing "always"/"never"/"only", name the search that would have found a counterexample and
  confirm it could have.
- **Mechanical damage is invisible to claim-checking.** Four full-read audits missed two broken sentences
  because a broken sentence asserts nothing false. Cheap structural greps catch a class that careful reading
  systematically cannot.

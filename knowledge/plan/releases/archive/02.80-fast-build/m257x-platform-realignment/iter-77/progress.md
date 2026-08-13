# iter-77 — progress

**Type:** tik, under `TOK-05` (*stop repairing claims; fence the predicates under them*).
Planned deliverable: **adjudicate `FIX-M257x-iter76-read-union`, then close the G5/G2 reach hole at
its cause** — under three binding conditions taken from iter-76's own evidence: *adjudicate before
repairing* · *repair by predicate, not by claim* · *not closed until the reach hole is closed*.

> **Resumed, not restarted.** A prior session was killed mid-Phase-D by a spend limit. It left
> `overview.md` written, `progress.md` and `decisions.md` **empty**, and the guard carrying
> +129/−3 uncommitted lines. This iter began by **reading that diff and re-deriving every number in
> it** rather than assuming it correct — which is how two of its four findings turned out to be
> false. See `D-M257x-77-2`.

---

## Phase A — adjudicate the union, mechanically

iter-76 routed **77 blockers in reading #13 and 75 in #14** and warned in its own close that
*"~150 is an upper bound, not a work item."* This phase did not re-read the union by hand; it
partitioned the classes the guard can decide and measured each.

The partition is by **ref-pin**, and it is mechanical. Every `repos.yml:A[-B]` citation in the
corpus (**20**, in 8 files) was resolved through the guard's **own** `_pin_window` / `_REF_PINNED` /
`_pin_exempts` machinery — not by eye, and not by a fresh regex written for the occasion.

| finding | adjudication |
|---|---|
| `jobsimulation.md:12` cites `repos.yml:17-19` | **FALSE POSITIVE.** Its block pins `2adcf71`, and at `2adcf71` lines 17-19 **are** the jobsimulation block, `migrations: false # legacy — folded into app (jobsim-in-app)` |
| `roadrunner.md:14` cites `repos.yml:29-31` | **TRUE.** Its block's only pin, `87d8d44`, **does not resolve in the platform clone at all** — it is a *roadrunner-repo* commit dating a terraform file |
| `roadrunner.md:29` cites `repos.yml:17` | **TRUE.** Unpinned; line 17 is `sentinel` |
| `jobsimulation.md:224` cites `repos.yml:17-19` | **TRUE.** Unpinned |

**4 → 3.** The fifth routed count in this milestone to shrink on adjudication (64→5 · 23→1 · 21→0 ·
92→0 · **4→3**), and it shrank by *exactly the mechanism iter-76 had already named one class
earlier* — grading a dated claim against today's checkout.

**Derived, not inherited:** platform `d11a403` (2026-08-03) is the commit that removed **both**
`roadrunner` and `jobsimulation` from `repos.yml`. Enumerated across all 9 commits that ever touched
the file. The entries did not move — they were deleted — which is why all three real defects read as
citations into some other repo's lines.

## Phase B — the vocabulary hypothesis, tested live and CONFIRMED

`overview.md` predicted a third answer to the briefing's design question: that *"the one claim G5
does reach is read WRONG"*, because the guard's repo vocabulary was `set(repos) | set(compose.services)`
— derived from what **currently exists**, so a name left the vocabulary at the same commit it left
`repos.yml`, going invisible exactly when it went false.

**Confirmed by counterfactual**, run rather than argued. The guard at `HEAD` and the guard with the
historical vocabulary were run over the identical corpus and identical platform ref:

```
committed guard  → 0 findings
patched  guard   → [G5 wrong-target-set] setup_guide.md:486 enumerates
                   ['app','cms','jobsimulation','skillpath']; repos.yml says ['app']
```

`_names_in` had resolved that enumeration to `{'app'}` — because `cms`, `jobsimulation` and
`skillpath` were no longer in the vocabulary — compared `{'app'} == {'app'}` and **passed a false
claim.** *G5's effective reach was 0 of 24, not 1.*

The fix is derived, not listed: the vocabulary is every `- name:` ever written in `repos.yml`, taken
from the platform artifact's own history. **Re-derived independently of the guard**, by iterating
the 9 commits: **14 names ever · 6 now · 8 removed** (chronos, cms, graphql-wundergraph,
intelligence, jobsimulation, roadrunner, skiller, skillpath). Exact agreement.

> **Environment trap, recorded because it silently returned an empty set.** In zsh,
> `git show "$SHA:repos.yml"` applies `:r` as a **parameter modifier**; the first derivation
> returned *0 names ever* and looked like a finding. `"${SHA}:repos.yml"` is required. An empty
> result from a mangled command is not evidence of absence (§8 rule 4).

## Phase C — can free prose be fenced? measured, and answered

The briefing asked it directly: *can free prose be fenced at all, or should the corpus restate those
21 claims in a form that CAN be checked?* Three candidate prose fences were built and measured
against the live corpus before any was shipped. All three are recorded **in the guard's own
docstring** and none shipped: membership-by-line **71% precision** (7 hits, 5 true, 2 false);
membership-by-clause **100% precision but 60% recall** (the corpus writes *"so both dropped to
`migrations: false`"* and the subject is a pronoun); membership-by-citation-cell rejected on
`storage.md:25`.

**The answer this iter implements is neither of the two the question offered** — see
`D-M257x-77-1`. Fence the predicates that are decidable **without reading a sentence** (the
vocabulary is a fact of git history; the citation is a fact of the file's line numbers), and leave
attribution honestly UNREACHED with its count printed on every run.

## Phase D — build, with the controls §8 rule 5 requires

**Shipped:**

1. **The historical vocabulary** (`repos_yml_history`) — memoised on the clone's own HEAD, so the
   key changes exactly when the history could have. A clone that cannot answer reports
   `UNMEASURED` in the reach line and falls back to the current set; it never pretends.
2. **G9 — `repos.yml` line citations.** Subject = the document's **filename** (a fact of the tree,
   exactly as G8 takes it). Each citation read **at the ref its own block names, and only if that
   ref resolves in the platform clone** — the rule this milestone already owned at iter-71 and had
   not applied here.

**The two halves are coupled, and the coupling is the finding.** `jobsimulation` and `roadrunner`
are in G9's subject vocabulary *only* because the vocabulary is historical. Without Phase B's fix
their docs have no derivable subject, their citations are counted-but-never-graded, and **G9 finds
zero defects.**

**Performance:** the history derivation costs ~3 s/call on a real clone. Memoised + an exact
substring pre-test, the guard now runs a live check in **9.3 s against a 10.7 s baseline** — faster
than before, with an assertion added.

### The widening that was built, measured, and NOT shipped

`_pin_exempts` applies `D-M257x-63-1`'s own-ref rule to the **leftmost** pin only, so a block
opening with a historical sha and closing with the guard's own is exempted whole. Widening it to
read *every* sha in the block was implemented and measured: **26** exempted blocks on the live
corpus name the guard's own ref other than first, and reading all of them changes exactly **one**
verdict.

**I graded that verdict TRUE by class-matching it to one of iter-76's dominant classes, and then
read the sites.** It is a **false positive**: `jobsimulation.md:72` and `roadrunner.md:57` read
*"**none — there is no `jobsimulation` compose service.** … the line that stood here named the
`graphql` profile, which `0dab54d` renamed `core` … Historical only."* Every fact is true and the
sentence says it is historical. The later `0dab54d` **narrates what a commit did**; it does not
assert currency at it, and no mechanical discriminator separates the two without reading English
(§4 Trap A). **Reverted, reported as a measured negative, and pinned by four tests** so it cannot be
silently re-tried. G9 never depended on it.

### Positive controls (§8 rule 5)

- **A no-op control that SURVIVES** — a citation narrowed from `9-11` to `10` is the same fact and
  stays GREEN; a fence matching only the exact `first-last` pair would fail here and would be
  matching text, not deciding a fact.
- **An INVERTED mutant** — the same citation moved one block over goes RED and names `sentinel`.
- **The inverted mutant on the derivation itself** (§5 rule 36) — the *identical* false enumeration
  passes GREEN when the history is removed, and the reach line says `UNMEASURED` rather than
  claiming historical coverage. The hypothesis is a test, not a paragraph.

### The guard went RED on this iteration's own repair — and that was the useful part

After the three repairs landed, G9 flagged `roadrunner.md:32`: a paragraph *inside `roadrunner.md`*
correcting a claim about **jobsimulation**, line-cited. The filename says one subject and the block
names another. That is an **ambiguous subject**, not a defect, so it is counted and declared
UNREACHED — the same honest degradation G5 makes for free prose.

**The order of that check is load-bearing, and the guard proved it on itself.** Tested *before*
grading, the rule took `storage.md:25` — a correct, gradeable citation whose sentence merely says
*"exactly as `cms` and `jobsimulation` are kept"* — from GRADED to UNREACHED, and G9's graded count
fell **4 → 2**. That is the guard going blind on a true claim because of a passing mention, which is
precisely the failure this iteration exists to end. Tested *after*, it can only convert a would-be
false positive into a declared unreached: **no recall is spent buying the precision.**

## Phase E — repair, by predicate

Four sites, three predicates, every replacement value derived from a platform artifact:

| site | predicate repaired |
|---|---|
| `setup_guide.md:486` | *the `migrations: true` set* → **`app` alone**; the others were not set false, `d11a403` **deleted their entries** |
| `jobsimulation.md:224` | same predicate, in a runnable comment that told the reader a dead line range |
| `roadrunner.md:14` | *repo X's entry is at lines A-B* → true **at `2adcf71`**, stated with the ref; gone at `0dab54d` |
| `roadrunner.md:29-32` | *the husk container still starts on a bare `make up`* — iter-76's single most-asserted false predicate — → true at `2adcf71`, and **neither entry nor service exists** at `0dab54d` |

## Phase F — re-measure, honestly

```
G9  4/19 repos.yml citation(s) graded (subject = the document's own filename;
    2 read at a historical ref, 1 AMBIGUOUS-SUBJECT and therefore UNREACHED)
repo vocabulary 19 (8 removed: chronos, cms, graphql-wundergraph, intelligence,
    jobsimulation, roadrunner, skiller, skillpath) from 9 commit(s) touching repos.yml
G5  23 migration claim(s) = 1 enumerated + 20 free prose UNREACHED + 2 ref-pinned
G2  3 repo-count claim(s)
platform_predicate_guard: OK
```

**The reach hole is closed at its cause, and the cause was not the twenty-one.** G5 still reaches
one enumerated claim of 23 — that number barely moved and **saying so is the point**. What changed
is that the one claim it reaches is now read **correctly**, where before it was read wrong and
passed. A denominator of 24 with an effective reach of 0 and a denominator of 23 with an effective
reach of 1 look almost identical in the reach line and are opposites in the field.

**G9 is net-new reach** into a construct no assertion in this family could see: **19 citations
enumerated, 4 graded, 1 declared ambiguous, all counted on every run.**

## Close — 2026-08-05

**Outcome:** the **G5/G2 reach hole is closed at its cause**, and the cause was not the twenty-one
unreached prose claims — it was that **the one claim G5 could reach was being read wrong**. The
guard's repo vocabulary was derived from what currently exists, so a repo left the vocabulary at the
same commit it left `repos.yml`; `setup_guide.md:486` enumerated *"app, cms, jobsimulation"*, the
resolver silently dropped the removed names, and the fence compared `{'app'} == {'app'}` and passed
a false claim. **Effective reach was 0 of 24, reported as 1.** Confirmed by counterfactual, not
argued. The vocabulary is now HISTORICAL — every `- name:` ever written in `repos.yml`, from the
artifact's own 9 commits (**14 ever · 6 now · 8 removed**, re-derived independently of the guard and
in exact agreement) — and **G9** adds net-new reach into a construct no assertion in this family
could see: **19 `repos.yml:N` citations enumerated, 4 graded, 1 declared AMBIGUOUS**, each read at
the ref its own block names **and only if that ref resolves in the platform clone**.
`FIX-M257x-iter76-read-union` was **adjudicated before repair** and shrank **4 → 3** — the fifth
routed count in this milestone to collapse — because `jobsimulation.md:12` is a *measurement* at
`2adcf71`, where lines 17-19 **are** its block. All three real defects repaired **by predicate**,
along with `setup_guide.md:486`.

**Type:** tik
**Status:** closed-fixed — every planned phase (A adjudicate · B hypothesis · C measure the prose
fences · D build + controls · E repair · F re-measure) ran and landed, including the pre-registered
*"report the negative"* outcome for a fence that could not clear precision.
**Gate:** NOT MET — 4 of 5. Clause 5 is not re-cut and is met only by a reading returning zero; this
iter repaired 4 sites and closed the instrument hole beneath them, it did not take a new reading.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n
— (5) cap-reached: n (1 tik of 5) — (6) protocol-stop: n — Outcome: **continue**.
**Decisions:** `D-M257x-77-1` (free prose is not fenceable AND the corpus should not be restated —
fence what is decidable without reading a sentence), `D-M257x-77-2` (inherited work is evidence, not
fact — two of a prior session's four findings were false), `D-M257x-77-3` (a derived vocabulary must
be HISTORICAL or the instrument goes blind exactly where the drift is), `D-M257x-77-4` (a narrowing
discriminator is checked AFTER the verdict, never before), `D-M257x-77-5` (*committing is not
pushing* — the sibling of rung zero).
**Side-deliverables:**
- `rosetta-extensions` `main` **pushed to origin** — 13 commits (8 hardening passes, 5 fences,
  ~1,400 lines) had existed on a single disk. Verified with `git ls-remote`.
- `corpus/ops/platform-alignment.md` §5 gains rules **37**, **38** and **39** — the three lessons
  that generalise past this guard, landed in the same commit as the work per the protocol-evolution
  rule.
- A guard-performance improvement that came free with the fix: memoisation on the clone's HEAD plus
  an exact substring pre-test take a live check from **10.7 s to 9.3 s** — faster than baseline
  **with** an assertion added.
**Routes carried forward:** `CHECK-M257x-iter77-narration-vs-documentation` (G1 reads historical
narration as documentation; latent, 2 sites) · `CHECK-M257x-iter77-cross-repo-pin` (145 pin-exempted
blocks name a sha that does not resolve in the platform clone; unmeasured whether any dates a
*platform* claim) · `CHECK-M257x-iter77-zsh-modifier` (zsh's `:r` mangles `git show "$SHA:file"` and
returns an empty set that reads exactly like a finding) · all iter-76 routes unchanged, including
`FIX-M257x-iter53-union-set` (**PENDING USER DECISION**) and
`CHECK-M257x-iter76-compose-service-count` (**explicitly unsettled**).

**Lessons:**

1. **Fix the instrument before reshaping the artifact it reads.** The briefing's restate-the-corpus
   option was the more attractive of the two on offer and would have been wasted work: 21 claims
   rewritten into enumerable form, read by a resolver that was mis-reading the enumerated form it
   already had.
2. **A denominator is not a reach.** *"1 enumerated of 24"* and *"0 effective of 24"* print almost
   identically and are opposites in the field. A reach line must state what the instrument
   **found**, not only that it looked — the same lesson G6 learned at harden pass 16, relearned one
   assertion over.
3. **Class-matching is not adjudication, and I did it twice in one iteration.** A new finding was
   graded TRUE because it matched a known dominant class; reading the two sites showed correct,
   self-labelled historical prose. iter-76's own warning applied to the iteration that inherited it.
4. **Order a narrowing rule against the verdict, not against the input.** The same correct rule cost
   half of G9's reach in one position and nothing at all in the other, and only the reach line could
   have told the difference.
5. **What a stack pulls is fenced; what the author has not pushed is not.** Thirteen commits
   survived on one disk precisely *because* the pin guard was green.

---
iter: 42
milestone: M257x
iteration_type: tok
tok_flavor: triggered
tok_trigger: user-directed (NOT the 3-no-prog streak — see "Why this tok fired" below)
status: closed-fixed
opened: 2026-08-02
---

# iter-42 — TOK-02, a strategy revision

**Prior strategy:** `TOK-01: instrument first, then follow` (milestone-root `decisions.md`), authored at
iter-01 and **never revised across 40 consecutive tiks.**

## Why this tok fired, and what it is not

**User-directed.** The 3-consecutive-no-progress trigger did **not** fire — the primary metric moved at
iter-41 (`37 → 18`). This tok fires because the user directed a strategy review after 40 tiks on one
strategy, in substance: *review the current strategy to address the slowdowns encountered and see how to
soften them, or accelerate progress in any case.*

It is recorded with `tok_flavor: triggered` because it has the **triggered tok's semantics** — it revises a
running strategy and it terminates the session so the revision is visible before the next tik commits to
it. The `tok_trigger` field states the honest provenance.

**Two user rulings bound this iteration and are not re-opened here:**

1. **Exit-gate clause 5 stays exactly as written.** The escalation iter-41 raised — *can a hand-maintained
   corpus satisfy a zero-blocker clause at all?* — was put to the user and answered: **the milestone
   continues against the clause as specified.** This tok therefore proposes no re-cut, no narrowing, and no
   alternative reading of "met". A pass returning zero is the only thing that meets it.
2. **The revision must not weaken the audit instrument.** iter-39 established that the headline series
   `25 → 13 → 11 → 17 → 37` measured the **instruments**, not the corpus; iter-38 measured that scoping to
   the highest-density files would have found **11 of 17** and declared the rest clean. A cheaper instrument
   produces an uncomparable number. The 7-auditor full read survives this revision intact.

## Step 0 — re-survey (mandatory before authoring)

Confirmed by measurement at open, not inherited:

- **The 18 blockers stand.** In-scope corpus (`corpus/services/` + `corpus/architecture/`) is **byte-identical**
  to what iter-41 measured — `git diff 103ad31..HEAD -- corpus/services/ corpus/architecture/` is empty, and
  103ad31 *is* iter-41's own closing commit. Four blockers spot-verified live in the files
  (#18 `architecture_overview.md:243` vs `external_services.md:537`; #15 `roadrunner.md:23-25`;
  #9 `service_taxonomy.md:136` vs `:75`; #8 `:145` vs `studio-desk.md:20`) — all four present, all four
  still contradicted by the twin the ledger names.
- **Gate 4 of 5.** Clauses 1–4 hold. Clause 5 at 18.
- **Platform origin `2adcf71`, unchanged** — re-scope trigger stays at occurrence 1 of 2.
- **rext `main` @ `069c238`**, clean, pin `fast-build-m257x-iter-37` on origin, both pins match.

## THE MEASUREMENT THIS REVISION RESTS ON

Every prior strategy discussion treated the 18 as *"18 wrong sentences"*. They are not homogeneous. I
classified all 18 by **the cheapest instrument that could have caught them**, reading iter-41's
`blocker-ledger.md` row by row and taking the class from the ledger's own *"what is true"* column:

| class | n | blockers | instrument that reaches it |
|---|---|---|---|
| **The corpus contradicts ITSELF** — a twin site inside the corpus already states the opposite | **13** | 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 15, 18 | a claim-twin fence. **Needs no platform read at all** |
| **The anchor resolves but names the wrong construct** — right line, wrong function/row | **3** | 13, 16, 17 | a symbol-aware anchor check |
| **A derived scalar vs platform source, with no corpus twin** | **2** | 10 (Go 1.25 vs `go.mod` 1.26), 11 (256 MB vs `locals.tf` 128) | a value fence |

**13 of 18 are the corpus disagreeing with itself.** Not with the platform — with another sentence in the
same repository, and in five of those cases (#1 `:102-103`, #5 `:69`, #6 `:202`, #12 `:174-176`, #14 `:458`)
**with another sentence in the same file, within a few lines.**

**And the split explains the dynamic that ends the loop.** Of the **9 repair-induced** blockers
(1, 2, 3, 4, 5, 6, 7, 13, 18), **8 are corpus-self-contradiction** — every one of them the signature of a
repair editing one site of a claim and leaving its twin: *"added the twin row, left this one"* (#4),
*"fixed one twin row, not the other"* (#18), *"the contradiction is inside its own text"* (#5), *"a
blockquote spliced into a bullet list"* (#6). The only induced blocker that is **not** self-contradiction is
#13, a cross-ref anchor.

> **This reframes the dominant cost term.** iter-41 concluded *"each repair injects defects at a rate
> comparable to what it removes, so the fixed point of this process is not zero."* That is right about the
> arithmetic and incomplete about the mechanism: **the injected defects are overwhelmingly ONE mechanical
> class — a claim repaired at one site and left standing at another — and that class is exactly the class a
> machine can check without knowing anything about the platform.**

**Stated weakness of this classification, up front.** The ledger names each twin because a *human auditor
found it*. A fence must find the twin **without being told**. That is not a fatal objection, because the
finding-by-grep-from-a-known-verdict direction is already **proven manually**: iter-40 swept 8 adjudicated
claims across the whole tree by grep, found 20 files, and its mandatory post-condition re-grep surfaced
**3 more sites the first pass missed — a 27% miss rate, by the author of the rule, applying it by hand.**
The mechanism works; the hand is what fails at it.

### The second measurement: some of these are RETURNING claims

At least four of the 18 restate a claim a **prior pass already adjudicated and recorded a verdict for**:
#12 (`:5050` — iter-40 swept it at 8 sites and missed this one, *inside* clause-5 scope), #18 (the EU-first
ladder, retracted at iter-39/40 and still published verbatim), #5 (the multi-tenancy fence, wrong for a
**fifth** consecutive time), #7 (a retraction contradicting `platform-migration-status.md:86` — the corpus's
own **machine-fenced** source of truth, linked by the very section that contradicts it).

Nothing in the corpus checks whether an already-refuted claim has come back. The verdicts exist, enumerated
with anchors, in five blocker-ledgers (iters 33, 34, 38, 39, 41).

### What is already fenced, and what that proves

For **five consecutive passes every `file:line` anchor a sweep introduced resolved correctly** — iter-41's
adversarial auditor verified ~110 across 91 hunks with **zero** failures, and ancestry-checked all 13 cited
shas. The anchors are not clean because anchors are easy. They are clean because **a machine checks them and
it runs on every pass.** The prose has never had an equivalent instrument. `stack-core` carries five corpus
guards — `corpus_index_guard`, `platform_alignment_guard`, `demo_knob_guard`, `story_org_count_guard`,
`dev_flag_guard`, plus `test_service_doc_status_fence` — and **every one is a single-subject membership or
table fence.** None reads a claim.

Blocker #15 makes the gap concrete: `roadrunner.md:23-25` asserts jobsimulation was *removed from
`repos.yml` + `docker-compose.yml`*. Verified at open: jobsimulation is at **`repos.yml:17`** and
**`docker-compose.yml:83`**. `platform_alignment_guard.py` **already parses both files** — it fences the
migration *map*'s membership rows and has no idea a service doc's prose asserts the opposite.

## Phase plan (tok)

- **Phase A** — re-survey + verify the gate state. *(done above)*
- **Phase B** — classify the 18 by cheapest reaching instrument; derive the induced/self-contradiction
  overlap. *(done above)*
- **Phase C** — author TOK-02 into the milestone-root `decisions.md`.
- **Phase D** — record the next-tik direction and the harden recommendation; close; commit; exit.

**No corpus text is repaired in this iteration** and no fence is built here — a tok authors strategy. The
18 remain enumerated and ready as `FIX-M257x-iter41-blocker-set`.

## Escalation conditions

- A platform commit landing mid-iter → re-scope occurrence 2 → STOP. *(checked: `2adcf71`, unchanged.)*
- Any finding that would require re-cutting clause 5 → **do not act**; the user has ruled. Record and route.

## Acceptable close-no-lift outcomes

A tok produces no metric delta by construction. The deliverable is the revised strategy plus the
measurement that justifies it. Gate stays 4 of 5.

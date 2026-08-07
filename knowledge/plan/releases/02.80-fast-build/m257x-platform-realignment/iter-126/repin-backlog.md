# iter-126 — the re-pin backlog, ENUMERATED

iter-123 measured the backlog and reported a count. **A count with no list cannot be closed**, which is
why priority 4 was still open a run later. This is the list.

## Denominator, and the substrate that decides it

**89 rows** — the non-`SUPPORTS` tier-1 verdicts: 68 `UNRESOLVABLE` + 17 `PARTIAL` + 4 `DOES-NOT-SUPPORT`,
read from `iter-122/verdicts/tier1-batch-*.tsv`. Each row's corpus site is a **line pin taken at the
census's ref**.

> ### The re-derivation was wrong the first time, and the reason is the milestone's own rule
>
> | corpus read at | NO-SHA-IN-BLOCK |
> |---|---|
> | **HEAD** (iter-126, ~30 commits later) | **22** |
> | **`afe58ac`** — the ref iter-123 named | **7** |
>
> **A 3.1× inflation, manufactured entirely by reading pinned sites at the wrong ref.** `D-M257x-122-4`
> established this for *platform* clones (*"a stale substrate FABRICATES defects"*); it applies
> identically when the substrate is **this corpus**. The blank-line-delimited block at
> `security_compliance.md:266` is simply a different paragraph today than it was at `afe58ac`.
>
> **Standing rule, one level up: a corpus-side derivation over line-pinned sites reads the corpus AT A
> REF, never at HEAD.** Reproducing iter-123's 7 exactly is the control that says the method is right.

## The 7, each with its disposition

| id | corpus site @ `afe58ac` | citation | disposition |
|---|---|---|---|
| `B03-039` | `customerio-sync.md:58-60` | `app/main.go:393-396` | **RE-PINNED** `@ 2035f9a` — verified: `:395` is `customeriosync.New(logger, copilotDB, …)`, which is the claim |
| `B07-030` | `ai_architecture.md:264` | `app/internal/cms/directus/collections/jobsimulation.go:983-990` | **RE-PINNED** `@ 2035f9a` — verified: the `AIModel` enum block, `anthropic-45-sonnet-aws` … `gpt-5.1` |
| `B11-022` | `security_compliance.md:293-302` | `app/internal/coursebuilder/bedrock.go:109-112` | **RE-PINNED** `@ 2035f9a` — verified: `newUnderlyingClient` → `AnthropicAPIKeyEnv` → `NewAnthropicClientWithModel` |
| `B04-018` | `hiring.md:241-252` | `persona_write.go:69-71` | **PATH-QUALIFIED + RE-PINNED** → `rosetta-extensions/stack-seeding/seeders/persona_write.go` @ rext `63ce41a`. A bare basename is refused by `D-M257x-122-5`, never resolved by proximity |
| `B04-020` | `hiring.md:258-294` | `persona_write.go:152-158` | **PATH-QUALIFIED + RE-PINNED**, same file/ref — verified: `sessionCols()` includes `token`, which is the claim |
| `B11-020` | `security_compliance.md:266` | `README.md:21` | **PIN REMOVED, NOT RE-PINNED** — see below |
| `B01-021` | `ai-readiness.md:499-507` | `ops/demo/stories-spec.md:599` | **STAYS, WITH THE REASON** — an **intra-corpus** citation. A sha does not apply: it would be a ref of *this* repo, and both `corpus_citation_guard` and `anchor_construct_guard` already resolve and grade every intra-corpus anchor on every run. That is a **stronger** control than a pin, not a weaker one. This member is a false positive of the no-sha class, and the class should exclude intra-corpus anchors |

## `B11-020` — the one that got worse before it got better, and the fence caught it

The citation was a bare `README.md:21` inside a sentence about the shared **`ai`** Go library.
**Qualifying it made it worse.** With the repo prefix added, the resolver bound it to **`studio-desk` @
`41ee357`** — a repo with nothing to do with the shared library — and landed on a **blank line**;
`anchor_construct_guard` went RED on the qualification, in the same session that wrote it.

**`ai` is a private Go module that no stack clones** — it is not in `repos.yml`, it is pulled at Docker
build via `GOPRIVATE`, and `ls stack-demo/` has no `ai` directory. **So no `file:line` into it is
verifiable from here at all.** The pin was therefore **removed**, and the sentence now names the document
in prose with the reason stated.

Two things this cost, both recorded:

1. **A manufactured hedge was written and then withdrawn.** The first correction claimed the anchor was
   *"not verifiable from here"* — true — **while leaving the unverifiable pin in place**, which is the
   worst of both. The directive's rule is *do not manufacture hedges for facts somebody can measure*; its
   mirror is **do not keep a pin you have just said nobody can check.**
2. **The retraction itself had to be rewritten in the fence's vocabulary** (§8, iter-98). The first draft
   explained the removal *by quoting the removed pin in backticks*, and the guard parsed the quotation as
   a live citation and stayed RED. **A retraction written in the vocabulary the fence enumerates is
   indistinguishable from the claim it retracts.**

## What is NOT closed here, and why

The other **6** of the ≤13 are iter-123's `PIN-DOES-NOT-RESOLVE` class — *"candidate rot; several are
blocks citing a second repo, not decay."* Deciding each needs a **full-history clone read at the block's
own named sha**, per row. Not attempted here: it is a third line of investigation and the scope-creep
tripwire governs. Routed as `FIX-M257x-iter126-pin-does-not-resolve-6`, with this file's method — read
the corpus at the census ref, resolve at the block's sha — as the procedure.

**Retired: 0.** No anchor among the 7 has a third-generation history, so iter-115's
retire-rather-than-re-derive precedent still does not fire.

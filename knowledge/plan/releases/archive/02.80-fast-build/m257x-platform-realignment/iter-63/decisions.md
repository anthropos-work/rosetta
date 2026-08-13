# iter-63 — decisions

## `D-M257x-63-1` — a ref-pin is a DATE, not an exemption

**Question the briefing asked:** should a pinned claim carry an expiry, a two-sided citation, or
something else?

**Answer: none of those three.** The premise that a pin is a *legitimate* exemption which merely needs
a shelf-life is wrong, and measuring it is what showed why. `CHECK-M257x-iter60-stale-pin-exemption`
is **three mechanisms wearing one name**, and only the third is the one the briefing describes:

| # | mechanism | measured instance | fix |
|---|---|---|---|
| 1 | the pin crosses a **row** boundary | `shared_libraries.md:41-42` — row 41 dates its colony pins *"at platform `2adcf71`"*; row 42's profile claim went silent with it | pin scope = the claim's block |
| 2 | the pin crosses a **cell** boundary | `service_taxonomy.md:98` — one row, two clauses; cell 2's `915da06` laundered cell 3's *"still starts in the default `graphql` profile"* | in a table the block is the **cell** |
| 3 | the pin is cited **as evidence of currency** | `service_taxonomy.md:55` — *"(**current** local docker-compose @ platform `2adcf71`)"*; `messenger.md:108` — *"**current, not stale**"* | a block asserting currency cannot be pinned into silence |

Mechanisms 1 and 2 are not policy questions at all — they are a **window bug**. The exemption was
correct in principle and simply reached further than the claim it was written for. An expiry would
not have touched either.

Mechanism 3 is the policy question, and the answer is narrower than an expiry: **the pin and the word
"current" are making opposite claims, and only one of them is what the reader acts on.** So the pin
loses, exactly there and nowhere else. That is derived from the corpus's own vocabulary — the same
class of discriminator as iter-61's adjacent-negation rule — rather than tuned to a threshold.

**And the inverse, which nothing had ever checked:** a pin naming the **guard's own ref** exempted
too. A document could immunise a false present-tense claim by citing the very commit that refutes it.
A pin earns silence only by naming a ref *other* than the one the reader is standing on.

**Why not an expiry.** An expiry is a threshold, and §4 Trap A is explicit that a threshold fitted to
the known-bad set is not a fence. It also gets the semantics backwards: *"`fdfa189` removed
intelligence"* is past-tense, pinned, and true forever; *"at `2adcf71` `CMS_RPC_ADDR` reads
`http://cms:8091`"* is present-tense, pinned, and was false within days. **Age is not the variable —
tense is**, and the three rules above catch tense structurally without parsing it.

**Why not a mandatory two-sided citation.** The two-sided form (`storage.md`'s G6-fenced mid-fold
record) is the right *repair* and several of this iter's repairs use it. As a *requirement* it would
force every historical sentence in the corpus to restate the present, which is how a fence starts
crying wolf. The pin stays legal; it just stops covering claims that were never its own.

**Blast radius, measured before adopting:** 19 blocks in the live corpus are both pinned and assert
currency; of those, exactly **1** also contained a construct any assertion checks. The rule removes
the exemption where the document itself insists the claim is current, and essentially nowhere else.

## `D-M257x-63-2` — the profile column is identified by its HEADER, not by its position

Fixing the pin was not enough, and the reason is iter-61's lesson recurring one construct along.
`service_taxonomy.md`'s Services table — the corpus's most-read infrastructure table — names its
profile column **fourth** (`| Service | Port(s) | Purpose | Profile | Source |`). G1 required the
profile column to be the table's **first** cell, which is the shape of a profile *reference* table.
So six rows naming a retired token for three deleted containers were **not exempt — they were
unreachable**, and the fence had never once looked at them.

Column identity now comes from the table's own header cell, in any position; cell contents are read
by shape (a bare comma list `graphql, backend` as readily as a backticked token), so `(always on)`
and `—` fall out without a stop-list.

**This is the fifth time in this milestone that a GREEN reading turned out to be a reach limit.**
Before trusting a GREEN, measure the fence's reach against the class it claims to cover.

## `D-M257x-63-3` — both prior sizings of the citation class were readings of a subset

iter-58 said *"21 of 22 moved"* (the protocol recorded *"22 of 23"*); iter-61 said *"5 of 16
distinct"*. Neither is the class.

- iter-58 counted **raw sites** of the string `main.go:N`, pooling `app/`-qualified, bare, and
  `cmd/*/main.go` forms — three different files under one name.
- iter-61 counted **distinct** citations of the two app-mainline forms, which is the right unit, but
  could not see the bare `:N` **continuation** construct the corpus uses for a run of citations into
  one file (`` `app/main.go:446`, `:524`, `:992` ``).

Derived enumeration over both constructs, resolved against the app clone on disk: **104 citation
sites / 86 distinct citations across 22 corpus files land in `app`**, of which **18 are the app
mainline** — 5 held, 13 moved. 18 is the routed class's denominator; **86** is the denominator of the
class §7 rule 4 actually names.

## `D-M257x-63-4` — a corpus repair moves the corpus's own line numbers

§7 rule 4's citation-safety half was written for a *pin advance*. It applies identically to a *corpus
edit*: this iter's repairs moved **9 intra-corpus citations across 6 files**, two of them in root
`CLAUDE.md`. `anchor_construct_guard` caught **one** — the one that landed on a blank line. Re-derive
the map from `git diff -U0` and re-point in the same commit. Recorded as §5 rule 34, with its two
sharp edges: the re-point is **not idempotent** (the map is computed against `HEAD`), and a citation
into a line the edit *replaced* needs a human.

## `D-M257x-63-5` — compose refutes four service docs at once; the map had it right

Adjudicating the citation repairs against platform artifacts (never against another document) surfaced
a live falsehood the citation work was not looking for. `docker-compose.yml:171-183` @ platform
`0dab54d` sets **all four** `*_RPC_ADDR` to `http://backend:8083`, under its own comment *"cms +
jobsimulation are folded into app: all four RPC edges are the one backend mux"*. **M809 has landed.**
`messenger.md`, `cms.md`, `jobsimulation.md`, `dependency_map.md` and `backend.md:195` all asserted
the two-of-four split as current — three of them emphatically. `platform-migration-status.md:76` was
correct the whole time. **The fenced map was right and the prose was wrong, which is what the map is
for**; and it is the second time this milestone that emphatic language marked the *falsest* claim.

## Routed forward

- **`FIX-M257x-iter63-app-citation-residual`** — the 86-citation `app`-resolving class beyond the
  mainline (68 non-mainline distinct citations). Graded mechanically this iter into
  HELD/MOVED?/GONE?/DEAD/UNNAMED buckets; the MOVED?/GONE? buckets carry known instrument artifacts
  (whole-line token pooling, and a continuation attaching to the nearest preceding path rather than
  the intended one). **Route WHOLE** — §5 rule 19's scope-edge corollary.
- **`CHECK-M257x-iter63-quoting-a-retired-token`** — the fence flags a *quotation* of the retired
  construct in the repair's own historical note (`service_taxonomy.md:58` on first draft). Rephrasing
  is the right answer for one site; if the class grows, a quotation discriminator is the build.
- **`CHECK-M257x-iter60-g6-citation-subject`** — untouched, cheap-if-reached.

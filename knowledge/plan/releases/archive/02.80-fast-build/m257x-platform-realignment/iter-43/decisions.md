# iter-43 — decisions

## D-M257x-43-1 — `## Minors` sections are excluded from the claim ledger

**Measured, not assumed.** The first run derived from every ledger-shaped table and returned **33 hits**.
17 came from minor sections, and **12 of those 17 were prose that is TRUE** with an anchor off by a few
lines — eight in `ai-readiness.md` alone, plus `alignment_testing.md`, `service_taxonomy.md`,
`next-web-app.md`, `ai_architecture.md`.

A blocker row records a claim that is **false**. A minor row overwhelmingly records an **anchor that
drifted** while *"the claim itself is TRUE"* — that phrase, or a paraphrase, is in the verdict column of
most of them. They are different KINDS of finding, not two severities of one kind, and matching a minor
row's quoted text tree-wide fires on correct prose. §8 rule 6 names where that ends: *"a fence that cries
wolf gets disabled, and a disabled fence is indistinguishable from never having written one."*

The exclusion is also aligned with iter-42's classification, which routed the wrong-construct class to a
**symbol-aware anchor check** — a different instrument. Those findings are not lost; they are somebody
else's job.

**The cost is real and is recorded rather than discovered later.** iter-41's blocker **#16**
(`messenger.md:110`) was first written down as minor `m-E3` at iter-34 and survived seven iterations to
become a blocker. A fence over minor rows would have caught it. Taking the trade anyway, because a fence
that fires on true sentences will not survive contact with a repair pass — and because #16 is squarely
inside the anchor-check's charter.

## D-M257x-43-2 — a waiver is TWO keys, and the second one is the machine's

A page may legitimately quote a refuted claim in order to retract it — three do. So there is a waiver
file, and Trap A ("tune it until it catches nothing") is the standing hazard for any such file.

A waiver is honoured only when **both** hold: the site is acknowledged in `claim_twin_waivers.json` with
a written reason, **and** `_looks_retracted` still finds a retraction marker within one block of the
quote (§8 rule 4 — scope the assertion to its enclosing block). Delete the retraction and leave the
sentence standing, and the waiver silently stops applying. That is §8 rule 3's demand: a fence must not
pin the current shape of the drift and convert the bug into a contract.

The detector and the human agreed on all four candidates independently — the three retraction sites read
`retracted_context=True`; `coverage-protocol.md:629` reads **False** and stays RED.

**And the load-bearing control is elsewhere**: `test_01` asserts all 18 answer-key sites still fire, so a
waiver that suppressed one turns the suite RED. The other two properties are only trustworthy because
that one exists.

## D-M257x-43-3 — the answer key is CAPTURED as a fixture, because the live corpus cannot be the test

§5 rule 21: *"The fixture is perishable. A corpus carrying a known, anchored, independently re-verified
answer key exists exactly once; repairing it destroys the only thing that can falsify the fence."*

TOK-02 step 4 repairs the 18. After that a test that reads the live tree can only ever assert GREEN — and
this milestone has caught a check reporting success without checking five times, plus M256's 43. So the
18 sites were snapshotted to `stack-core/tests/fixtures/claim_twin/red/` at rosetta `48ca53c`, **with a
GREEN twin of each** (`green/`, the same neighbourhood minus the offending line).

The GREEN twin is not decoration. A battery of REDs cannot distinguish a discriminating fence from a
brittle one, and exactly one mutant proves it: `fragment-floor-collapsed` (30 → 3 characters) leaves
**every** answer-key site firing and is caught **only** by the GREEN twin.

Consequence to carry: `test_01` is now the milestone's memory of what the corpus looked like on
2026-08-02. It must not be relaxed when the corpus is repaired — a repair should leave it untouched,
because it reads the fixture and not the tree.

## D-M257x-43-4 — the ledger is derived by table STRUCTURE, and the sources are not named

§5 rule 19's closing clause forbids a hand-assembled claim list, and TOK-02 restates it. Naming the five
ledger files would have re-created the hand-maintained tuple §2 deleted, one level up.

So a file is a ledger iff it contains a markdown table whose header carries both a claim-shaped and an
anchor-shaped column. Measured: **4 files, 85 blocker rows, 36 claims, 39 refuted forms** — and the two
ways a row falls out of reach are **counted and named**, never silent: 17 rows quoted nothing longer than
the fragment floor, 32 quoted no refuted form at all. A future audit that changes its ledger shape shows
up as a shrinking claim count against a steady row count, which is a finding a reader can act on (§5
rule 8 — a check that SKIPS reads exactly like a check that PASSES).

## D-M257x-43-5 — NOTHING was repaired, including the one thing it was tempting to repair

`D-M257x-42-3` is binding and was obeyed literally. The fence found a claim **no pass has ever caught** —
`corpus/ops/demo/coverage-protocol.md:629`, a claim iter-34 refuted that survived outside the audited
scope — and it was **routed, not fixed**, as `FIX-M257x-iter43-coverage-protocol-livepath`.

Two reasons, and the second is the one that matters. It is out of clause-5 scope, so fixing it moves no
metric. And verifying it needs the `app` repo, which is not cloned: §5 rule 19's closing rule is that a
claim-scoped repair must **propagate a verdict, never adjudicate one**, and deriving a fresh claim during
a repair pass is how rule 18's highest-risk text gets written.

The only corpus edit this iteration made is **new** text: the §8 fourth-layer section in
`platform-alignment.md`, documenting the fence. The fence was re-run as a **post-condition** over that
edit — 18 hits, byte-identical set. That post-condition is TOK-02 step 2 in embryo.

## D-M257x-43-6 — the protocol doc's *"keep `.md` prose out of scope"* bullet is reconciled, not contradicted

§8 carries the design decision *"keep `.md` prose out of scope; that is review, not a fence."* This
iteration ships a fence that reads `.md` prose, so leaving the two side by side would have manufactured
exactly the self-contradiction the fence exists to catch — in the document that governs the fence.

Reconciled explicitly in the new §8 section: that bullet is right about the **write-target** fence, and
stays right, because per Trap B prose can be false at the same sha and fencing on it would mechanically
encode a falsehood. The claim-twin fence does not judge prose either — **it only ever asserts a verdict
some auditor already recorded, with an anchor.** It never adjudicates.

## Routed forward

| handler | what | target |
|---|---|---|
| `FIX-M257x-iter43-coverage-protocol-livepath` | `coverage-protocol.md:629` restates a claim iter-34 refuted; out of clause-5 scope, needs the `app` repo to verify | TOK-02 step 4, or the ops-scope sweep |
| `CHECK-M257x-iter43-value-fence` | blocker #10 (`go.mod` 1.26 vs *"Go 1.25"*) and #11's memory half — derived scalars with no corpus twin | TOK-02 step 3 |
| `CHECK-M257x-iter43-symbol-anchor` | blockers #13, #16, #17 — an anchor that resolves but names the wrong construct. **#16 is the one this fence deliberately gave up** (D-M257x-43-1) | TOK-02 step 3 |
| `CHECK-M257x-iter43-markdown-lint` | blocker #6's spliced blockquote / orphaned bullet, and iter-38's *"The The"* — pure mechanical damage | TOK-02 step 3 |

# iter-86 decisions

## `D-M257x-86-1` — "all six corpus guards exit 0" was never a statement about the family, and the fix is a runner, not a habit

iter-83 diagnosed why `repair_leak_guard` was not run: it declares `FENCE_KIND = "standalone"` while
`repair_postcondition.py`'s **derived** registry selects only `postcondition` kinds, and 10 of the 14
guards then standing were standalone. That is true. **It is not what went wrong**, and iter-86 measured
the difference.

**The census, over iters 77–85, of every claim that guards were green:**

| iter | claimed | enumerated? | evidence |
|---|---|---|---|
| 77 · 78 · 79 | "5 corpus guards GREEN" + a captured `platform_predicate_guard: OK` | **no** (members never listed) | captured transcript of **one** guard |
| 80 | "All 5 corpus guards GREEN" — used to declare **gate clause 3 MET** | partial: 2 of 5 named | none |
| 81 | subject says **5**, body lists **6** | the only full enumeration in the window, and it is second-hand (reconstructed at iter-83) | none |
| 82 | — | — | no guard claim at all |
| 83 · 84 · 85 | "6 corpus guards exit 0 at open and at close" | **no** | **none captured** |

Two things follow, and the second is the one that matters.

**(1) The count moved 5 → 6 with no record of what joined.** No iteration records `claim_twin_guard`
entering the set. The only artifact that reconciles the two numbers contradicts itself inside one
message.

**(2) The assertion is FALSE for one of its own members.** Re-measured at iter-86 with **the guard
version each iter declared** (`rext 24819f08`) against **the tree each iter closed at**, with the
platform ref each iter declared (`0dab54df`, which is `origin/main`, 0 behind):

| tree | `platform_predicate_guard` |
|---|---|
| iter-83 `b4a4af3` / close `36cca77` | **rc=1** — and iter-83's own commit introduced the sites |
| iter-84 `c1d4a27` (iter-85's declared open) | **rc=1**, 2 sites |
| iter-85 `b4b6db8` (close) | **rc=1**, 3 sites — **iter-85's own repair added the third** |

> **The earlier `OK`s were real.** iters 77–79 captured genuine transcripts; today's guard finds more on
> those trees only because it has grown assertions since (G10 arrived at iter-78). **The claim did not
> become false when the guard changed. It became false when the evidence stopped being captured** — at
> exactly the iteration where "guards green" stopped being a pasted transcript and became a sentence.

**And 9 of the 15 guards are covered by no green claim anywhere in the window** — `value_change_guard`
(which appears **nowhere** in the milestone outside iter-49), `repair_leak_guard`, `repair_reach_guard`,
`demo_knob_guard`, `dev_flag_guard`, `evidence_visibility_guard`, `story_org_count_guard`,
`union_apply_guard`, and `derived_value_guard` — **which is a `postcondition` guard**. The registry-filter
finding cannot explain that last one. The list in use was not the derived registry and not its
complement; it was **a list somebody remembered**.

**So §2's deleted 4-tuple did return — not as a runner's hardcoded list, which is what
`repair_postcondition.py`'s docstring guards against, but as a human's memory, which is worse: a tuple
in source can at least be diffed.**

**Decision: `guard_family.py`** (`rext stack-core/`) — one command that runs the family and names every
member. Its census is **derived from `*_guard.py` on disk**; its invocation map is **declared**, because
guards genuinely answer different questions (tree-state · commit-scoped · needs-a-ledger) and no honest
invocation can be derived from a filename. What keeps the map from becoming the tuple is that it is
**reconciled against the census in BOTH directions** — a guard on disk with no entry exits 2 naming
itself, and an entry naming a guard not on disk exits 2 too. It also refuses to read a guard's own
*"CANNOT RUN … Nothing was checked; this is not GREEN"* as a pass.

**First full-family run: 16 members, 14 GREEN, 2 RED** (`platform_predicate_guard`, `value_change_guard`),
0 could-not-check, 0 not-run. Both REDs were invisible to every green claim this milestone has made.

## `D-M257x-86-2` — the seat-ref escalation: the sheet is NOT the instrument, and the measurement says so

`CHECK-M257x-iter76-seat-ref-discipline` fired at its 5th occurrence. The proposed fix — make the
re-derived ground-truth sheet authoritative for refs and require seats to cite from it, carrying each
clone's `origin/main` sha beside its checkout sha — carries a stated risk: **a sheet handed to seats is
arguably part of the instrument, and if it is, adopting it quietly breaks the comparability that makes
140 → 43 meaningful.**

**Two measurements settle it, and neither is a preference.**

**Measurement 1 — the sheet already varies between the readings being compared.** iter-84 re-derived 13
clone shas at its open; iter-83's sheet named different ones; iter-80's differed again. All three
readings are counted into the same 140 → 43 series. **A per-iter input that already changes between two
compared readings cannot be the invariant that makes them comparable** — if it were, the series was
never comparable and the number was never meaningful. What holds the series together is the frozen
briefing (§5), the 40-file set, the partition and the model. The sheet is none of those. **Adding a
column to it is not a re-cut.**

**Measurement 2 — the class has contributed ZERO to the graded number, 5 times out of 5.**

| # | iter | disposition |
|---|---|---|
| 1 | 76 | settled **correct** at adjudication (§5 rule 33); recorded as an upper-bound inflation, not a defect |
| 2 | — | routed, same class |
| 3 | 82 | `D-M257x-82-5` — adjudicated in-run as a **false positive** |
| 4–5 | 84 | both **REJECTED** by the adjudication |

**Every occurrence was filtered before it reached the graded count.** Clause 5 grades a
*post-adjudication* reading, so the class cannot block a zero and never could.

**Decision: ADOPT the sheet, and declare the RAW series discontinuous at iter-86 — out loud.**

The user's read is upheld, with one correction to its reasoning. The fix does belong outside the frozen
briefing, but **not** because it is costless: it changes what a seat books, and the seat-ref rejection
rate is measurable at **3 of 43 = 7.0 %** (iter-84). So the **raw** booked series takes a ~7 % step down
for a reason that has nothing to do with corpus quality, and anyone reading 152 → 43 → *n* across
iter-86 must know that. The **adjudicated** series — the one clause 5 actually grades — is untouched,
because adjudication was already removing this class.

> **Stated plainly so it cannot be smuggled: the raw-count series is re-baselined at iter-86. The
> adjudicated series is not.** This is the third of the three honest options (leave it to adjudication
> and pay the overhead · re-baseline deliberately and say so · change it silently). The third was never
> available; the first was affordable but pays 7 % overhead forever to preserve a series that is already
> discontinuous for other reasons.

The sheet ships as [`ground-truth.md`](ground-truth.md). **6 of 14 clones are behind `origin/main` right
now** — `app` by 60, `next-web-app` by 26 — which is exactly the invisible condition that produced
occurrences 3, 4 and 5.

**The §5 rule-33 amendment stays ROUTED, not written.** *"Grade at the ref the claim names — UNLESS the
sentence asserts currency"* is the correct statement of the rule, and amending the frozen briefing's
subject between readings is the one thing `TOK-04` protects. It goes in at the next deliberate
re-baseline, alongside this one, not mid-run.

## `D-M257x-86-3` — G1 could not reach zero on a correct corpus, and that is why it stopped being run

`platform_predicate_guard` had been RED since it was authored. Adjudicating its 3 sites:

| site | verdict |
|---|---|
| `graphql-wundergraph.md:13` | **guard defect.** *"The `graphql` profile **is gone** too"* is as flat a denial as English has. `_NEGATED` reads only the text **before** the noun phrase, so a POSTFIX denial read as a fresh claim |
| `platform-alignment.md:1060` · `:1071` | **corpus defect.** The protocol doc **quoted the false sentence verbatim** as a worked example — publishing a live-reading copy of the very claim it exists to kill. Second time (`:1305` was iter-84's) |

**Both repaired, and each on the correct side of the line.**

The doc sites were repaired in the corpus, not waived: `CLAUDE.md` already states the governing rule —
*"no retired token is spelled here in runnable form — a copy-pasteable command for a silent no-op is the
defect"* — and rule 40 is the file that teaches it. A worked example does not need a working copy of the
defect. The sentences are now **described**, not quoted.

The guard site was repaired in the guard, because contorting a correct English sentence to dodge a regex
would tax every correct sentence written after it. `_NEGATED_AFTER` completes a discriminator that was
half-implemented; the module's own comment already says *"parsing English negation is not fitting a rule
to an answer key"*, and prefix-only was that rule, unfinished.

**The predicate list denies EXISTENCE only** — *gone / removed / renamed / retired / dropped /
decommissioned / no longer / does not exist*. Not a bare *"is not"*, because *"the `storage-legacy`
profile is **not** started by default"* denies a profile's DEFAULTNESS, not the profile. **The two false
claims this guard was holding open both survive the change and stay RED** (*"…name **survives** in
compose"*, *"…**is now simply the default**"*) — that separation is the evidence the rule was fitted to
English rather than to the answer key, and it is a test, not an assurance.

**3 mutants, 3 kills, 3 distinct signatures:** restore prefix-only → 6 failures naming the defect ·
widen the predicate to any word after a copula → **the two CONTROL tests fail**, which is the whole
point of writing them · drop the next-line window → the wrap test alone fails. **+7 tests** (160 → 167).

`platform_predicate_guard` is **GREEN for the first time since iter-60.**

## `CHECK-M257x-iter86-value-change-weak-form` — a 3-token form of ubiquitous words is not evidence

Surfaced by running the family, and worth routing rather than tuning at the end of a long run.

`value_change_guard` builds a form from *the changed token run + the surviving context on each side*,
capped at `MAX_CHANGED_TOKENS = 3` with a floor of `MIN_CONTEXT = 2`, then matches it **in order within
a 90-character window** — not contiguously. On this iter's diff, one correction (`the graphql backend
stack` → `the core backend stack`) produced the form **`('graphql', 'backend', 'stack')`**: one changed
token plus the two most ubiquitous nouns in this corpus. It matched **two sentences that state the
CORRECTED fact**, both correct as written:

- `architecture_overview.md:256` — *"GraphQL — `backend` directly on a local stack (`:8082/graphql/query`)"*
- `backend.md:204` — *"GraphQL via `backend`'s own endpoint … on a local stack since `2adcf71` deleted the router"*

**0 of 2 precision on that form.** `MIN_CONTEXT = 2` is derived and defended in the module (the
motivating fixture's right context is exactly two tokens), so the floor is not the defect — **the defect
is that distinctiveness is measured in token COUNT and not in token RARITY.** Two occurrences of
`backend` and `stack` carry almost no evidence in a corpus about a backend stack.

**Not fixed here, and not waived — the waiver mechanism correctly refuses it.** `is_waived` requires
`retracted_context`, so a waiver only applies to a site that is *quoting a value in order to retract
it*. These sites are not retracting anything; they are simply correct. **A mechanism that cannot
express "this is a false positive" is the right mechanism** — it is what stops a waiver file from
pinning drift (§8 rule 3) — so the answer is a precision fix in the guard, not an entry in a JSON file.

**What cleared it was a better sentence, not a tuning.** The repaired cell was rewritten from a
one-token value swap into a real rewrite (*"together with whatever the default profile selects —
`Makefile:120` passes `--profile core --profile frontend`, i.e. …"*), which pushes the changed run past
`MAX_CHANGED_TOKENS` and hands the question to `repair_leak_guard`, where it belongs and where it is
GREEN. That is a legitimate outcome — **the sentence really was under-specified** — but it is a repair
of one instance, not of the class. **Routed** with the measurement attached.

## `D-M257x-86-4` — the EXIT_REASON contract, stated so a clean budget stop has exactly one name

iter-85 reported `EXIT_REASON: user-blocker` while its own Phase 5 grading read six `n` and `continue`.
It named the contradiction against itself rather than smoothing it, which is right, and the contract is
the thing to fix — **an honest agent should not be able to repeat it.**

The gap: the six values are `gate-met · tok-fired · re-scope-trigger · user-blocker · cap-reached ·
protocol-stop`, and `cap-reached` is described as *"hit the 5-tik cap"*. **A run that ends cleanly
because it ran out of ROOM — not tiks — matches no description**, so it gets reported as whatever is
nearest, and `user-blocker` is nearest in feel and furthest in meaning: it wakes a human for nothing.

**Binding for this milestone, until the skill's own contract says otherwise:**

> **A clean stop with no partial work, no gate, no tok, no re-scope and nothing needing a human is
> `cap-reached`** — the value that means *"this session's budget is spent; re-invoke to continue."* The
> cap is the session boundary whether it is measured in tiks or in room. **`user-blocker` requires a
> named question whose answer would change what code lands.** No question, no blocker.

Corollary, from iter-85's own honesty: **if the Phase 5 grading and the reported EXIT_REASON disagree,
the disagreement IS the finding.** Report both and let it escalate; do not reconcile them in prose.

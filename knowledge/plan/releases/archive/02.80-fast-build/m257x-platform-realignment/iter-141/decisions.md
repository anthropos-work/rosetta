# iter-141 — decisions

## `D-M257x-141-1` — the corpus's own RETRACTION IDIOM is a rot generator, and this session measured it

House style for a corrected citation is *"it was `:274` at `<sha>`"* / *"this cited `:116-117` until
iter-NN"*. **That keeps the retracted number live in the text**, where the next insertion above its target
moves it.

**In one session — iters 137 → 141 — this turned fences RED three times, in three different files, always
on a pin whose own sentence existed to retract it:**

| file | the retracting sentence | outcome |
|---|---|---|
| `roadrunner.md` (iter-137) | *"this said `:124` below, and at iter-120 `:124` was above this very line"* | landed on a blank line; 2 fences RED |
| `graphql-wundergraph.md` (iter-138) | the `5050` pointer | **rotted twice on its own** (`:174-176` → iter-98 → `:193` → iter-138) |
| `ai-readiness.md` (iter-141) | *"Read that `:326` as a pin … it was `:274` at `bb3313bc`"* | landed on a blank line; 2 fences RED |

> **Rule.** **Retract by describing the artifact, not by reproducing it.** *"This doc carried two
> different line numbers for it in successive iters"* says everything the quoted number said and **cannot
> rot.** A fence matching on *form* cannot distinguish the quotation from the assertion — and it is right
> not to; the reader cannot either.

Three occurrences in five iters is not anecdote. The idiom is load-bearing prose in a milestone whose
whole subject is citation integrity, and it is **generating** the defect class it exists to document.

## `D-M257x-141-2` — a cross-reference that names its target by a RETRACTED TITLE is invisible to every anchor fence

`backend.md:13` sent readers to *"the **M810 prod teardown is UNEVEN** bullet below."* That bullet was
**retitled at iter-127** to *"The M810 prod teardown has now LANDED for both"*, and its body **retracts
"UNEVEN" in the first sentence.**

**No anchor fence can see this**, and that is the point: the pointer is a *name*, not a line, so it still
resolves. The reader arrives at a paragraph that opens by contradicting the sentence that sent them.

> **Rule.** **Name your target by what it says now.** A title is a citation. When a section is retitled
> because its claim was retracted, every pointer that used the old title is now asserting the retracted
> claim — in the pointer, where no fence is looking.

This is the sibling of `D-M257x-137-3` one level up: 137 covered quoting a retracted *pin*; this covers
quoting a retracted *name*, and it is harder to catch precisely because nothing breaks.

## `D-M257x-141-3` — `adj-B`'s P-3 upheld and widened at source

The seam claim was *"the **only** remaining dependency on `workforce` is the member directory (the
`WorkforceDirectory` interface — `LoadMembers`/`LoadMembersByUserIDs`, whose implementations **stayed** in
`members.go`)."* Re-derived independently at `app` `ad9f3c498`:

- the interface declares **four** methods, not two — `LoadMembers` (`:43`), `LoadMembersByUserIDs`
  (`:45`), **`BaseMembers`** (`:48`), **`LevelsCount`** (`:50`), between `:40` and `:51`;
- **the source's own doc comment already said so** (`manager.go:36-39`): *"the active-member directory …
  **and the org's skill-scale setting**"* — so *"only the member directory"* is contradicted by the
  construct the sentence cites;
- **`LevelsCount` is an org setting, not a member call** (`readiness.go:770`);
- and *"stayed in `members.go`"* is **wrong for the fourth**: it lives at
  `internal/workforce/manager.go:90` (`git grep "func .*LevelsCount"` → three sites: the unexported
  `getLevelsCount` at `:61`, the exported one at `:90`, a test fake).

**An absolute quantifier over a coupling seam, refuted by the doc comment on the interface it names** —
the same shape as this milestone's four security-surface understatements, on a different axis.

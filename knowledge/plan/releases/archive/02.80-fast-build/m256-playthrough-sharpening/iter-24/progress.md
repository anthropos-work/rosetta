# iter-24 — progress

**Type:** tik (standard shape, `TOK-01` — the onboarding cluster's seeder half, priced before any spec)

## Phase A — measure hypothesis 2

**CONFIRMED, and it re-prices the next use case.** `onboarding.enterprise-workforce-ai-readiness.UC1`'s
routed blocker reads *"needs an Org C stage-0 seat"* — which implies the seat is declarable and merely
absent. Measured, by declaring an Org-C-shaped story, appending a 4th end-user hero and reading back the
stage the seeder assigns her:

| declared | stage assigned |
|---|---:|
| manager | **0** |
| end-user, `trajectory: struggling` | **1** |
| end-user, anything else (incl. `thriving`) | **3** |

**There is no stage-0 outcome for an end-user.** An appended Org C hero arrives **COMPLETED**, and a
Playthrough built on that seat would drive a hero who has already done the thing it means to watch her do
— and could well *pass*, by asserting the completed surface. "Declare her struggling" is not a workaround
either: that is stage 1, not 0. The only stage-0 seat in the system belongs to the **manager**, who is not
an onboarding member actor.

So the blocker is not a missing row in a YAML file; it is **a capability the seeder cannot express**. That
is a materially different piece of work, and iter-25 now plans against it instead of discovering it after
editing the seed.

## Phase B — pin the append-only rule, and hold the gap

`stack-seeding/seeders/seat_append_test.go`, three tests:

1. **`AppendingDoesNotRenumberExistingPersonas`** — the fence for iter-18 D89. It asserts the *baseline*
   first (1-based declaration order) rather than assuming it, because a comparison over a wrong baseline
   proves nothing; then appends and requires that no pre-existing index moved; then requires the appended
   seat takes the next free slot and that `personaIndexMapForStory` stays collision-free (two personas on
   one index = two heroes sharing a user row).
2. **`InsertingMidListRenumbers`** — the **self-test**, so the fence is discriminating rather than
   trivially true. It performs the wrong edit (insert at position 1 — *"next to the hero it relates to"*)
   and requires that exactly the 2 personas below the insertion point move.
3. **`NoStage0EndUserSeatCanBeDeclared`** — **holds Phase A's gap as a test.** Its failure message is the
   deliverable: when stage-0 support lands, this test SHOULD fail, and it tells whoever landed it to
   discharge the UC's verdict and delete the test.

**3 mutants, each restore byte-identical** (`cp` backups, never `git checkout`):

| # | mutation | result |
|---|---|---|
| N1 | the fence's own append becomes an insertion | RED — names `[pt-ai-started pt-ai-manager]` as renumbered |
| N2 | `personaUserIndexFor` stops indexing by declaration order | RED on the baseline assert (`indexed 3, want 1`) **and** on the self-test |
| N3 | `aiReadinessStageFor`'s hero default becomes 0 | RED, printing the discharge instruction |

**Deliberately not done:** changing the seeder to support a stage-0 end-user, and editing
`pt-world.seed.yaml`. Both are real work, and doing either now would start it on a premise measured five
minutes earlier — which is what Phase A exists to prevent.

## Phase D — verify

- `stack-seeding`: `go test ./...` **rc 0**, 0 FAIL, 16 packages.
- `gofmt -l` clean across all ten rext sections (after the side-fix below).
- **No live-stack change**, so no gate run is owed — nothing shipped touches the harness, the seed or the
  demo. `demo-2` was left exactly as iter-23 left it.

## Close — 2026-07-30

**Outcome:** the append-only seat rule is a **fence** instead of a comment, mutation-verified in both
directions — and the next onboarding UC's blocker turned out to be **mis-stated**: a stage-0 end-user seat
cannot be *declared* at all, so what was priced as "append a seat" is a seeder capability. The gap is held
by a test that will fail when it is closed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** recorded in this file (Phase A is the decision; no separate `decisions.md` entry was needed
— the measurement table IS the finding, and it is cross-referenced from the routing table).
**Side-deliverables:**
- **`gofmt -l` was not clean in the committed tree** — three files, all landed by iter-21
  (`stack-seeding/cmd/stackseed/main.go` alignment; `policy_grants.go` and
  `stack-snapshot/cmd/stacksnap/adapters.go` doc-comment reformatting, where Go 1.19+ rewrites a `''` pair
  inside a **doc** comment to a Unicode `”`). Fixed in its own commit (`b4c802d`), the two doc-comment
  cases **reworded to plain ASCII** rather than accepting the typographic quote. Does not change this
  iter's close status.

**Routes carried forward:**
- **`ONBOARD-M256-stage0-capability`** (net-new, replaces the "append a stage-0 seat" framing) → a
  stage-0 end-user vantage the AI-readiness funnel seeder can express, so a day-0 member exists in an
  AI-readiness org. **iter-25.**
- `ONBOARD-M256-seat-append` → **the rule is fenced**; the *use* of it still belongs to whichever UC lands
  first.
- The other 3 onboarding UCs, `D-v28-5` part (b), `PT-M256-readiness-step-asserts`,
  `NEGCTL-M256-studio-pair` → unchanged.

## Lessons

1. **A blocker recorded as a paraphrase gets re-planned as a paraphrase.** *"Needs an Org C stage-0 seat"*
   is a true sentence that hides the whole cost: it names the artifact and not the capability, and it reads
   like a YAML edit. Two commands measured that the capability does not exist. **A routed blocker should
   name the thing that is missing, not the thing you would add if it weren't.**
2. **The dangerous version of this gap is not a failure — it is a pass.** An appended Org C hero arrives
   *completed*, so an onboarding Playthrough written against her would drive a hero past the flow and could
   satisfy itself on the completed surface. That is the milestone's signature defect class arriving through
   the **seed** rather than through a locator, and it is the fourth distinct door it has come through
   (assertion, instrument, URL shape, seed).
3. **A gap held by a test is remembered; a gap in a routing table is re-derived.** Test 3 exists to fail
   later, and its failure message carries the instruction. That costs one test and removes the chance that
   the next author measures `aiReadinessStageFor` a second time.
4. **A verification claim inherits the moment it was made.** iter-21 reported `gofmt -l` clean and shipped
   three dirty files; the check was real and it ran before those files existed in their final form.

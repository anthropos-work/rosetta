# Adjudication adj-F — seats `iter-131/raw/r33-F.md`, `iter-131/raw/r34-F.md`

**Scope line.** Independent re-adjudication of the two Seat-F readings (#33 and #34) of M257x iter-131.
9 claimed BLOCKERs (r33: B1–B4; r34: B1–B5). I read only the adjudicator brief and these two seat
reports from `knowledge/plan/**`; no other iter dir, no `progress.md`/`decisions.md`/`adjudication.md`,
no other adjudicator's output. Every corpus claim was graded against the corpus text **as it stood when
the seats were dealt** — `a532493` (*"the read is SEALED — pre-registration committed before any seat is
dealt"*), the last corpus-affecting commit before `probe(M257x/131)` — and then re-checked against the
live tree so that since-repaired defects are labelled as such rather than rejected. Platform claims were
settled at the ground-truth shas in the brief, or at the ref the corpus passage itself names.

**Instrument note.** This shell's `git grep -E` silently returns nothing for `\b`; every route/group
enumeration below was run with `-P` **and** with a `-F` positive control on the same ref, and I verified
that every `*echo.Echo` parameter in `app/internal/**` is named `e`, so the `e.<METHOD>(` sweep is
complete rather than merely non-empty. `infrastructure` is in **no clone** in this tree (I enumerated all
19 `.git` dirs) — where that matters I say so and say what I settled instead.

## Counts

| | n |
|---|---|
| claimed BLOCKERs | 9 |
| **UPHELD** | **9** |
| REJECTED — `wrong-tree` | 0 |
| REJECTED — `misread` | 0 |
| REJECTED — `true-at-its-ref` | 0 |
| REJECTED — `retraction-not-contradiction` | 0 |
| REJECTED — `minor-not-blocker` | 0 |
| REJECTED — `not-in-scope` | 0 |
| CANNOT-SETTLE | 0 |
| **upheld rate** | **9/9 = 100 %** |
| **DISTINCT PREDICATES** | **4** |

One upheld blocker (r34-B5) is upheld **on a corrected predicate** — the seat's number is right and its
diagnosis is wrong. See § *Where I disagree with the seat's framing*.

## Verdict table

| seat | B# | anchor | verdict | rejection class | predicate (if upheld) | class | multi-pin | repair-induced (sha) |
|---|---|---|---|---|---|---|---|---|
| r33-F | B1 | `corpus/services/jobsimulation.md:12` | UPHELD (since-repaired) | — | P1 | self-contradiction + intra-corpus-citation | yes | no (`cd16967`, iter-102) |
| r33-F | B2 | `corpus/services/jobsimulation.md:54` | UPHELD (since-repaired) | — | P1 | self-contradiction | yes | no (`f8be5a1`, fix/108) |
| r33-F | B3 | `corpus/services/roadrunner.md:13` (+`:30-32`, `:53`, `:74`) | UPHELD | — | P2 | intra-corpus-citation + self-contradiction | yes | no (`21e7d4a`, fix/22) |
| r33-F | B4 | `corpus/architecture/architecture_overview.md:228` | UPHELD | — | P2 | self-contradiction (half-applied repair in one table) | yes | no (`904502c`, iter-87) |
| r34-F | B1 | `corpus/architecture/architecture_overview.md:83` (+`:233`, `:239`) | UPHELD (since-repaired) | — | P3 | platform-drift + self-contradiction | yes | no (`ae5c1db`, iter-86) |
| r34-F | B2 | `corpus/services/jobsimulation.md:12` | UPHELD (since-repaired) | — | P1 | self-contradiction | yes | no (`cd16967`, iter-102) |
| r34-F | B3 | `corpus/services/jobsimulation.md:54` | UPHELD (since-repaired) | — | P1 | self-contradiction | yes | no (`f8be5a1`, fix/108) |
| r34-F | B4 | `corpus/services/skiller.md:26` | UPHELD (since-repaired) | — | P1 | self-contradiction (cross-file drift) | yes | no (`f8be5a1`, fix/108) |
| r34-F | B5 | `corpus/architecture/architecture_overview.md:406` | UPHELD **(corrected predicate)** | — | P4 | arithmetic/count | yes | **yes — `3785b47` (iter-129)** |

*Repair-induced was evaluated with `git log -L<n>,<n>:<file> --oneline a532493` (the state the seats read),
not on the live file, so that iter-132/133's own repairs do not contaminate the answer.*

## Upheld predicates, deduplicated within my assignment

```
P1 | "`infrastructure` has never been in any clone set, so the production module declarations it holds
     are unmeasurable from this corpus — cms's ECS state and the production RPC address must be reported
     both-ways-and-asserted-neither"
     anchors: corpus/services/jobsimulation.md:12, corpus/services/jobsimulation.md:54,
              corpus/services/skiller.md:26
     class: self-contradiction (with an intra-corpus-citation failure at :12)

P2 | "roadrunner's retirement from production is unresolved — `roadrunner/terraform/main.tf:19`
     `service_desired_count = 1` is prod terraform and is evidence of a live production ECS service"
     anchors: corpus/services/roadrunner.md:13, :30-32, :53, :74;
              corpus/architecture/architecture_overview.md:228
     class: intra-corpus-citation + self-contradiction

P3 | "the `ai` library is one of the private Go modules the platform's services import — one of the four
     `colony, proto, ai, taxonomy`"
     anchors: corpus/architecture/architecture_overview.md:83, :233, :239
     class: platform-drift + self-contradiction

P4 | "`app` mounts SEVEN routes on the root Echo instance outside any group, and that seven-row table is
     the complete un-grouped surface"
     anchors: corpus/architecture/architecture_overview.md:406
              (same figure at corpus/architecture/security_compliance.md:250, :253-262 table, :265, :293)
     class: arithmetic/count
```

**Dedup notes.** P1 folds five booked blockers (r33-B1, r33-B2, r34-B2, r34-B3, r34-B4). The brief's own
worked example — *"cms's production ECS state is unmeasurable / infrastructure was never in a clone
set"* — bundles the premise and its cms conclusion into one predicate, so I bundled the RPC-address
conclusion with it too: the falsified proposition at all three anchors is the identical flat sentence
*"which has never been in any clone set"*, and both downstream hedges are wrong **because** it is. P2
folds r33-B3 and r33-B4: one proposition, two anchors, two files. That leaves **4 distinct predicates
from 9 booked blockers** — a 2.25× anchor-to-predicate multiplier inside this one seat's two readings.

## The evidence I opened, per upheld blocker

### P1 — the `infrastructure` clone-set premise (r33-B1, r33-B2, r34-B2, r34-B3, r34-B4)

Text verified verbatim at `a532493`:

- `jobsimulation.md:12` — *"The deciding declaration is in `infrastructure`, which has never been in any
  clone set: **report both, assert neither** — see `cms.md` and the fenced map."*
- `jobsimulation.md:54` — *"**And the production declaration is not measurable from this repo at all:** it
  lives in `infrastructure`, which has never been in any clone set — so no *"still names"* claim can be
  made here in either direction."*
- `skiller.md:26` — *"…the deciding declaration lives in the `infrastructure` repo, **which has never
  been in any clone set**…"*

Against, in the same file and in the two authorities `:12` points at:

- **Same file, 24 and 74 lines away.** `jobsimulation.md:78` opens *"**MEASURED at `infrastructure`
  `13c248e6` (iter-123), and it confirms this bullet exactly:** `module "jobsimulation_euwest1"` **IS
  still declared**, at `terraform/production/services.tf:475`…"*, and `:86` concludes *"…the opposite of
  `cms`, **whose block is gone**."* A document cannot hold both that a repo has never been in any clone
  set and quote it by `file:line` at a named sha — and `:86` asserts precisely the cms verdict `:12`
  forbids.
- **The fenced map `:12` directs the reader to.** `platform-migration-status.md:88` (cms row) reads
  *"**RESOLVED at M257x iter-123 — the ECS service is DESTROYED.** `infrastructure` was cloned and read
  (`13c248e6`, 2026-08-07)… **The blocker this cell named for four iterations — "the destruction happens
  in infrastructure's `services.tf`, which we cannot read" — was a clone-set limit, not a measurement
  limit, and the fix was to clone the repo.**"*
- **The other four sites.** `cms.md:8` (*"DESTROYED — corrected M257x iter-127"*),
  `service_taxonomy.md:169`, `architecture_overview.md:227` (which explicitly retracts the exact sentence
  `:12` still carries), `org-repos.md:134-145` (the ten `module "…"` declarations and the four orphaned
  repo-side counts).

I also re-measured the two cms facts `:12` calls *"opposite ways"* at the ref it names: `cms` `f38c0c4a`
`terraform/main.tf` is **191** lines with `:39` = `service_desired_count = 0` (true), and `6efa1d5`
(2026-08-04) deletes `.github/workflows/build-production.yml`, subject *"the cms ECR repository is
decommissioned (M810)"* (true). Both true; the corpus's own measurement says the first describes nothing,
so the pair was never a contradiction and *"assert neither"* is the wrong instruction.

**On the one thing I could not open.** There is no `infrastructure` clone anywhere in this tree (19 `.git`
trees enumerated; none is it), so I cannot independently verify `13c248e6`. **That does not block these
verdicts**, because the falsified proposition is not *"the measurement is right"* — it is the flat
sentence *"has never been in any clone set"*, which the corpus refutes eleven times in its own text and
which is a claim about **this corpus's own history**, not about AWS. The live tree confirms it: iter-132
rewrote all three anchors to *"not in the standing clone set and **HAS been read**"* — the corpus itself
now concedes the premise, in the corrected form the seats' finding implies.

**Status:** all three anchors are **since-repaired** (`ceba938`, iter-132). Upheld as booked.

### P2 — roadrunner's production state (r33-B3, r33-B4)

Verified at `a532493`: `roadrunner.md:12-14` (*"'There is no roadrunner service in production' overstates
it"*), `:30-32` (*"the **one row where prod and the platform's own declaration contradict each other** —
recorded, not resolved"*), `:53` (*"**prod's** `service_desired_count = 1`"*), `:74` (*"Treat retirement
as pending, not done"*) — and `architecture_overview.md:228` (*"**Gone locally, orphaned in prod** … while
**prod terraform still reads `= 1`**"*).

The fenced authority this banner names **twice** (`:22` and `:32`), `platform-migration-status.md:90`,
reads: *"**MEASURED AT LAST (M257x iter-123, `infrastructure` `13c248e6`):** a service repo's own
`service_desired_count` is **not evidence of production state** … **`roadrunner` appears in NO terraform
in `infrastructure` at all** … **There is no roadrunner ECS service, and `roadrunner/terraform/main.tf:19`
describes nothing.** The whole line of enquiry below is therefore SETTLED."* `org-repos.md:134-145`
carries the derivation and names `roadrunner/terraform/main.tf:19` `= 1` **orphaned** in its table.

**One half of this I settled without `infrastructure` at all, and it is the load-bearing half.** I read
`roadrunner/terraform/main.tf` at `87d8d443`: 95 lines, and `:10-11` is `module "roadrunner" { source =
"github.com/anthropos-work/infrastructure.git//modules/services/base_internal_service?ref=main"` with
`:19` `service_desired_count = 1` among inputs fed from `var.environment`, `var.platform_cluster_id`,
`var.platform_vpc_id`, `var.platform_private_subnets_ids`. **That file is a module, not a root module** —
its variables are unbound. Calling it *"prod terraform"* (`architecture_overview.md:228`) or *"prod's
`service_desired_count = 1`"* (`roadrunner.md:53`) misattributes a module input to production state,
which is exactly the error class `org-repos.md` § 3 was written to close, and it is demonstrable from the
roadrunner clone alone. `:30-32`'s *"recorded, not resolved"* directly contradicts the *"SETTLED"* of the
file it sends the reader to; `:74`'s *"retirement as pending"* is the conclusion drawn from the
misattribution.

`architecture_overview.md:228` is a **half-applied repair inside one table**: the CMS row one line above
(`:227`) carries the full iter-123/127 correction — *"the prod ECS service is **DESTROYED** …
`cms/terraform/main.tf:39` is **orphaned dead code**"* — while the Roadrunner row keeps the pre-iter-123
framing. Both anchors are still **live in the working tree** (unrepaired as of this adjudication).

I note the weakest point honestly: `:228` does say *"orphaned in prod"*, the corrected word, before the
clause that re-asserts the refuted evidence in the present tense. I still uphold — a top-level summary
table that tells a reader an ECS service is declared at desired-count 1 in production is a working
hazard, and the seat's own hesitation named the same tension.

### P3 — `ai` as an imported private module (r34-B1)

`architecture_overview.md:83` at `a532493`: *"**four** imported private modules — colony, proto, ai,
taxonomy"*, with `:233` (*"Imported as private Go modules — **not** cloned by `make init`"*) heading a
table whose `**ai**` row is `:239`.

I enumerated `github.com/anthropos-work/` in each of the seven Go clones' `go.mod` at its own ground-truth
ref:

| lib | requirers (of 7) |
|---|---|
| colony | 7 |
| proto | 7 |
| taxonomy | 6 (not `roadrunner`) |
| **ai** | **2 — `cms` `v1.40.2`, `jobsimulation` `v1.40.2`** |
| authn | 0 |

`app` `ad9f3c49` `go.mod:14-18` is `analytics-go v0.3.1`, `colony v0.35.2`, `proto v1.210.0`, `storage
v0.15.2`, `taxonomy v1.2.0` — **no `ai` line**, and **`storage` and `analytics-go` are absent from the
corpus's four**, so the list is wrong in both directions. The only two requirers of `ai` are frozen legacy
repos with no compose service at platform `0c91421` and no `repos.yml` entry — nothing a stack builds.
The sentence applies the exact test that excludes `authn` (*"no service's `go.mod` requires the standalone
module"*) and then keeps `ai` in the four. Corroborated by `shared_libraries.md:73` (*"Folded into `app` at
`1e457fa70`; **in no live `go.mod`**"*), `:170` (**Imported by** — *"**No repo a stack builds**"*) and
`jobsimulation.md:176`. **Since-repaired** at `9c86e0f` (iter-133): the live `:83` now reads *"**five**
imported private modules — analytics-go, colony, proto, storage, taxonomy … **`ai` is NOT among them**"*.

### P4 — the root-mounted route count (r34-B5)

`architecture_overview.md:406` at `a532493` and **still live today**: *"plus **seven routes mounted on the
root outside any group**, so no group-level statement reaches them at all"*, in a sentence that states its
own scope as *"re-measured **repo-wide** at `app` `ad9f3c498`"*. Same figure at
`security_compliance.md:250`, its 7-row table `:253-262`, `:265` (*"11 groups + 7 ungrouped root mounts"*)
and `:293`.

Measured at `app` `ad9f3c49`, non-test, `internal/**` (excluding the separate `cmd/labsdemo` binary), with
`-P` plus an `-F` positive control and after confirming every `*echo.Echo` parameter in `internal/**` is
named `e`:

**Echo groups = 11 ✓** (`invitations/handlers.go:31`, `academy_embeddings_admin.go:41`, `backend.go:121`,
`:171`, `:194`, `:210`, `:229`, `:273`, `content_admin.go:35`, `emailpreview/handler.go:66`,
`labs_admin.go:31`; the `graph.go:15466` complexity hit and the `emailpreview/handler.go:6` comment
excluded). Exactly **2** carry `cbGate` (`backend.go:232`, `:276`) ✓.

**Root-mounted routes = 8, not 7:** `aireadiness/notifications/handlers.go:41`, `backend.go:117`, `:309`,
`:315`, `:317`, `:324`, `content.go:23`, **`labs_admin.go:40`**.

The error is in the flattering direction — it under-counts, by one, the routes that escape group-level
authorization, in the passage whose whole purpose is to enumerate that escape surface. Upheld as a
blocker rather than a minor for that reason, and because the corpus states the number four times and
tabulates it once.

## Where I disagree with the seat's framing

**r34-B5 is right by number and wrong by diagnosis, and the diagnosis is what a repair would act on.**

The seat named `/ai-readiness/unsubscribe/:token`
(`internal/aireadiness/notifications/handlers.go:41`) as the eighth route, and built a theory on it:
*"7 is reachable only by scoping to `internal/web/backend/` — a scope that would also drop
`internal/invitations/handlers.go:31` and take the group count from 11 to 10. One sentence, two
incompatible scopes."*

I opened the corpus's own enumeration, which the seat did not: `security_compliance.md:253-262` is a
**seven-row table**, and its third row **is** `/ai-readiness/unsubscribe/:token`, cited to
`internal/aireadiness/notifications/handlers.go:41` and *"mounted `web.go:153`"*. The route the corpus
actually omits is **`/v1/labs/:slug/workspace.tar.gz`**, `internal/web/backend/labs_admin.go:40` — a route
**inside** `internal/web/backend/`.

So the scope theory is refuted by the corpus's own table, and it inverts the direction: the corpus
**includes** a route outside `internal/web/backend/` and **excludes** one inside it. A repair driven by
the seat's report — *widen the scope* — would change nothing and would leave the defect in place. The
corrected predicate is a plain enumeration miss, not a scope mismatch, and the missed route matters more
than the one the seat named: `labs_admin.go:36-40`'s own comment says *"Serve is **OUTSIDE the write
group** — it has **OPTIONAL auth** (a public Lab's workspace is served to anyone; a tenant-private Lab
requires a key with access)"*, it is wired unconditionally into the app server at `backend.go:301`, and it
serves a downloadable tarball. The route the seat named is HMAC-verified in-handler and already
documented as such. **P4 is stated above as the corpus's claim, not as the seat's reasoning.**

Two smaller framing notes, neither changing a verdict:

- **r33-B1 and r34-B2 lean on `platform-migration-status.md`'s cms row as a clean opposite.** It is not
  clean: that same map still carries the stale premise in its **jobsimulation** row (`:89`, *"the
  destruction itself lands in **infrastructure**, which is in no clone set"*) and at the tail of its
  **roadrunner** row (`:90`, *"the authoritative rollback declaration lives in **infrastructure's
  `services.tf`** — a repo this map has never read"*), inside the very row whose head says *"MEASURED AT
  LAST"*. Neither seat booked those, and neither is in my assignment; I record them because the P1 sweep
  is wider than either seat's anchor list and a repair scoped to the seats' anchors would miss them.
- **r33-B4's citation `org-repos.md:140` is one row off** — `:140` is the table header; the roadrunner row
  is `:143`. Immaterial to the verdict; the quoted content is at the cited section.

## Cannot-settle

None among the nine. The one substrate I could not open — `infrastructure` @ `13c248e6` — is not load-bearing
for any of these verdicts, for the reasons given under P1 and P2: P1 turns on a claim about this corpus's
own clone history (settled from corpus text, eleven ways), and P2's decisive half turns on
`roadrunner/terraform/main.tf` being a module rather than a root module, which I read directly at
`87d8d443`. **What would settle the residual** (whether the `13c248e6` measurements the corpus reports are
themselves accurate): the `infrastructure` clone at `13c248e6` in the ground-truth set, with its sha and
fetch time recorded like every other clone. Both seats asked for exactly that, and both were right to.

## Counts

```
UPHELD=9 REJECTED=0 (of which wrong-tree=0) CANNOT-SETTLE=0
DISTINCT-PREDICATES-IN-MY-SET=4
```

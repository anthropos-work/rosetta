# iter-137 — decisions

## `D-M257x-137-1` — *"merged into `app`"* and *"deleted"* are different claims, and only one of them predicts a package

The corpus carried **eight** services as folded into `app`. Seven have a package to show for it —
`app/internal/{cms,customeriosync,jobsimulation,messenger,skiller,skillpath,storage}/`. **`roadrunner` has
none, at any ref, ever**: `git log --all --diff-filter=A -- internal/roadrunner` returns **0 commits** in a
full 6,728-ref clone at `app` `ad9f3c498` (positive control `jobsimwiring` → 3 paths). Its job was
**replaced**, in-process, **inside the jobsimulation domain** — `app/internal/jobsimwiring/wiring.go:123`,
whose own comment says the runner *"replaces the removed roadrunner RPC edge"*.

> **Rule.** *Merged* is a claim about **where the code went**; *deleted* is a claim about **what stopped
> existing**. A doc that writes the first when the second is true sends every reader looking for a package
> that was never created — and it hides the fact that a **capability was re-implemented**, which is a
> different risk profile from a capability that was **moved**.

The corpus's own repaired side has said **seven** since iter-102 (`services/README.md`,
`platform-migration-status.md:88`). **This was a half-repair with a two-iter head start and 17 unswept
twins** — the class iter-133 and iter-135 both named.

## `D-M257x-137-2` — the two false predicates contradicted each other, and neither reader could see the other

Two adjudicators, blind to each other in iter-135, reported roadrunner defects pointing in **opposite
directions**:

| | predicate | what it tells a reader |
|---|---|---|
| `adj-F` P2 | *retirement is unresolved — prod terraform still reads `service_desired_count = 1`* | **still running in production** |
| `adj-B` P-2 | *`roadrunner` is one of eight domains folded into `app`* | **absorbed into the monolith** |

Both were live at HEAD, in the same corpus, and **both are false.** A reader could open
`service_taxonomy.md` and `dependency_map.md` in either order and come away with mutually exclusive
pictures, each stated with confidence.

> **Rule.** When a subject is wrong in **two directions at once**, neither error is a typo — the subject
> has no owner. Repairing one conjunct alone would have made the corpus *more* consistent-looking and no
> more true. **Sweep the subject, not the sentence.**

Recorded because iter-136 fixed the adjudicator brief for conflating conjuncts *within* one predicate;
this is the sibling case — **two predicates about one subject, each individually coherent, jointly
impossible.**

## `D-M257x-137-3` — a retraction that quotes the retracted line-pin re-publishes it

`roadrunner.md`'s § Async-tasks paragraph carried a bare pin **as its own worked example of a bad pin** —
*"this said `:124` below, and at iter-120 `:124` was above this very line."* This iter's repair shifted the
file, `:124` landed on a blank line, and **`anchor_construct_guard` + `repair_postcondition` both went
RED** on a citation whose entire purpose was to warn against citations like it.

Fixed by **deleting the pin, not re-pinning it**: the construct name (*the "Upstream consumers" bullet
under § Dependencies*) is the citation. Re-pinning would have restarted the same clock.

> **Rule.** A fence that matches on **form** cannot distinguish *asserting* a pin from *quoting a
> retracted* one. This is the same blindness iter-132 found for hedge markers and iter-134 measured at
> **1 of 4** fences — here it appears on the **anchor** axis rather than the prose axis, and it is the
> instrument's declared floor working correctly, not a bug. **Write retractions so the retracted artifact
> is described, never reproduced.**

## `D-M257x-137-4` — four planning searches still missed three sites; and one of them no search could reach

Four independent searches at iter open (§5 rule 57) enumerated the predicate. After repairing, a
**re-run of the same searches with the corrected vocabulary excluded** surfaced **three live survivors**
no original search reached — `architecture_overview.md:357` (*"it was folded in with jobsim-in-app"*,
inside a paragraph about gRPC hops) and `services/README.md:57` (*"but prod terraform still reads `= 1`"*,
in an index row). All three were repaired — the third (`architecture_overview.md:22-23`) carries its falsity in a **section heading** (*"Domains inside Backend/App, not services"*) inherited by its bullets, and was found by **reading the file**, not by any search.

> **Rule.** Rule 57 says a count is only as wide as its search. **The corollary is procedural: the search
> that plans a repair and the search that verifies it must not be the same search.** Verify with the
> *corrected* vocabulary and a `grep -v` of your own new wording — what survives is what your planning
> search could not see.

    **And the third survivor bounds even that.** `architecture_overview.md:22-23` is a Roadrunner bullet
    sitting under a heading that reads ***"Domains inside Backend/App, not services"*** — the false
    predicate is asserted by the **heading** and inherited by every bullet beneath it.

    > **Rule (second half).** A predicate carried by a **section heading** is invisible to every
    > line-oriented instrument this milestone owns — grep, the anchor fences, the claim census, all of
    > them read lines. **After a sweep, open the files and read the headings.** There is no cheaper way to
    > see this class, and it is not rare: a heading is exactly where a corpus states its general claims.

## Side-deliverable (does NOT grade this iter)

> **⚠️ Disclosure — it could NOT be given its own commit, and that is stated rather than papered over.**
> The protocol's default is a separate commit for an unrelated side-fix. This one lives in
> `corpus/README.md`, **the same file as planned-scope Q2 work**, and the only ways to split it are
> forbidden here (`git stash`, `git checkout --`, a revert-and-reapply dance on the working tree). It
> therefore rides in the `iter(M257x/137):` commit and is booked separately **here** instead. It does
> **not** upgrade or otherwise touch the close status, which grades planned scope only.


**`corpus/README.md:18` — the 16th escape of the cms-M810 predicate**, and it sat on the **front door of the
corpus**: *"M810 … is **uneven**: landed for jobsimulation, **not moved for cms**."* That claim was
retracted at iter-124, and swept corpus-wide **twice** (iter-127 at 5 sites, iter-132 at 15). Width
re-measured before repairing: **1 live site** — every other match is a retraction quoting the old wording.
Rule 55 exactly: the reader who wants the migration state stands at `corpus/README.md`, and neither sweep's
search reached it.

## Upheld claims counted as results (this milestone's practice)

- **`org-repos.md:143`** — *"`roadrunner` appears in NO terraform in `infrastructure`; 7 org-wide hits, all
  `judge0_*` secret names in two CI workflows plus one KB line."* **Re-derived independently at
  `13c248e6` and upheld byte-for-byte** (7 hits: 6 across `wf-terraform-deploy.yml` +
  `wf-terraform-plan-preview.yml`, 1 at `knowledge/service-dependencies.md:119`).
- **`platform-migration-status.md:91`** — the fenced map's roadrunner row was **right all along** and is the
  authority the service doc contradicted for four readings. Only its *"a repo this map has never read"*
  clause was stale (repaired).
- **The 31 `roadrunner`-as-domain hits inside `rosetta-extensions` are ALL frozen test fixtures**
  (`stack-core/tests/fixtures/repair_leak/{pre,post}/`) — deliberately-frozen fence inputs. **Not
  repaired, and repairing them would have corrupted what the fence measures.** The predicate has no live
  home in the tooling repo.

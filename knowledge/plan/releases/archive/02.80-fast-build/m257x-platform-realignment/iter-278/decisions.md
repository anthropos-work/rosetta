# iter-278 — decisions

## `D-M257x-278-1` — a drift fence's RED names a REVIEW that is owed, never a renumbering that is due

`clone_drift_guard` went RED naming **44 citing sites**, and the shape of that sentence invites exactly one
response: repair 44 sites. **That response would have been wrong at every one of them.** All 44 are
ref-scoped claims, and two instruments in this repo already say so in their own words — D1's docstring
(*"this fence does NOT adjudicate truth-at-a-ref … a fence that cries wolf gets suppressed, which is worse
than no fence"*) and `advance_impact_census`, which files a block-pinned citation under `pinned` and
**refuses to merge it into the subject** because *"counting it as `held` would inflate the safe bucket with
citations the advance was never able to reach."*

**Decision: read the fence's assertion, not its noun phrase.** D1 asserts *the clone advanced past every
commit the corpus knows about* — i.e. **a repo's worth of change has never been looked at**. The remedy is
a **review of the delta**, and the sha lands wherever that review changed a claim. **0 of 44 renumbered.**

**Why this needed deciding rather than assuming:** iter-256 took the other branch. It applied 27
renumbering repairs, found its premise refuted mid-iter, and reverted all 27 byte-for-byte — **seven of
which had moved correct citations onto comments.** A fence whose finding-count is a count of *mentions* will
keep proposing that work; the count is not the workload.

## `D-M257x-278-2` — the drift is paid down with a NEW-STATE claim, never with prose about the drift

`clone_drift_guard`'s own docstring books
`FIX-M257x-iter107-drift-fence-satisfiable-by-prose`: its first RED was on `sentinel`, the section
documenting that RED contained the sha, **and the next run was GREEN with nothing repaired.** *"Writing
about the drift satisfies the drift fence."*

That escape was available here — one sentence in this iter's own progress note would have done it — and
**it was not taken**. Every sha written this iter sits on a claim about what is true at HEAD: the studio
hardcode *is gone*, `pinned` *reads the workspace copy*, the test file *has 34 functions*. **Graded by
reading and recorded as PR-4**, because a fence that can be satisfied cheaply is only worth its green if
someone states, each time, which way it was earned.

## `D-M257x-278-3` — 0 of 13 anchors held, and 0 of 13 were repaired; the direction is not measurable

Comparing each unpinned anchor's cited line at `d739952` against HEAD: **13 of 13 differ.** The tempting
reading — *"the advance broke 13 anchors"* — is unsupported, and one site refutes it outright:
`safety.md:712` and `media-substrate-spec.md:122` cite the **same function** (`bridge_bedrock_creds`) at
**`:1358` and `:1364`** — the old ref and the new — with nothing in either sentence marking which clock it
is on. A third (`org-repos.md:195`) matches **neither** ref and was stale before this advance began.

**Decision: measure it, publish it, route it, repair none of it.** `D-M257x-122-5`'s rule transfers
unchanged from bare basenames to refs — *guessing between two candidates that equally satisfy a citation is
the wrong-construct error the instrument exists to find*. Routed as
`ROUTE-M257x-278-thirteen-unpinned-rext-anchors-are-on-undecidable-clocks`.

**The general form, and it is the milestone's own subject turned inward:** a corpus that cites a moving repo
without a ref is not *wrong*, it is **ungradeable** — and an ungradeable claim cannot be fixed by a fence,
only by an author writing the ref down. That is the argument for §5 rules 41/44 stated as a cost.

## `D-M257x-278-4` — a fence's registry entry is a CLAIM, and it decays exactly like any other

The two repairs this iter landed are both **documentation of our own instruments**: a fence's arm list
(three, shipping four) and its test count (16, shipping 34), plus the file its subject mechanism actually
reads. Nothing fences the fence registry.

**Decision: repair in place, do NOT build a registry-fencing fence this iter.** The instinct is to
mechanize — this milestone has done it 30-odd times — but the honest sizing is one row, and a fence over a
one-row population is the vacuity class this milestone has caught eight times. **Recorded so the next
occurrence has a second data point**, which is what would justify it.

**What made it findable at all** is worth stating: §7 of the *same document* has carried the correct
mechanism claim since iter-257. The document disagreed with itself for 21 iters, and the half that was
right was the half written by the iter that measured it. **A doc's newest paragraph is its most reliable
one; the registry table is where claims go to stop being re-derived.**

## `D-M257x-278-5` — the runbook half of iter-277's route is closed; the fence half is left open ON PURPOSE

iter-277 lost 53 minutes to a venv it created **inside the tree its census scans**, and offered two
resolutions: *"either the counters exclude virtualenvs or the runbook states the venv must live outside the
tree."* This iter needed a pytest and so met the same fork immediately.

**Decision: take the runbook half now (§5 rule 80, written from the invocation this iter actually ran), and
say in the rule itself that it is the weaker half.** A rule can be forgotten by the next operator; an
exclusion cannot. Closing both would have meant an rext edit — and an rext edit **advances the clone past
the sha this very iter has just reconciled the corpus to**, which is the coupling iter-277 named and this
iter is demonstrating. So the fence half is routed, not skipped.

## `D-M257x-278-6` — a route that must not be spent is named in full, in the corpus, at its point of use

`FIX-M257x-278-clone-pin-guard-docstring-says-three-arms`: `clone_pin_guard.py`'s header still enumerates
three arms and still carries the superseded *"checks each clone out at the ref it names"* framing — the very
sentence the corpus copied and this iter retracted. **The code is right; only the docstring is behind.**

Written into `platform-alignment.md` beside the repaired row rather than only into this ledger, because the
next reader of that row is the person who will otherwise copy the docstring forward a second time. Spelled
**in full** — iter-277's `D-M257x-277-2` cost two repairs to an **elided** id, the second of which re-tripped
the guard by quoting the elided form inside the footnote explaining the rule.

## `D-M257x-278-7` — the iter whose finding was "a closed fix published as open" published a closed route as open

The headline repair here is `services/cms.md` asserting `FIX-M257x-268` **open** after iter-270 closed it.
While writing that repair, this iter's own close block listed
`ROUTE-M257x-265-stack-demo-carries-six-dead-clones` under *"unchanged and not absorbed"* — **in the
corpus and in the milestone ledger** — and it was **closed at iter-268**. `route_disposition_guard` went
RED on exactly that, in the run before the commit.

**Decision: record it as the iter's own instance, not as a near-miss.** The route's deliverable was the
**census**, and *"nothing is deleted in this iter"* was one of iter-268's sealed pre-registrations — so
*"the clones are still on disk"* is a **measured, accepted state**, and reading a surviving symptom as a
surviving route is precisely the inference that produced the headline defect. Both are now phrased as
closed-with-residual.

**This is the second consecutive iter caught by this guard on its own prose** — iter-277's
`D-M257x-277-2` was an elided route id whose *first repair re-created the defect inside the footnote
explaining the rule.* The pattern is not carelessness about ids; it is that **a route's SYMPTOM outlives
its closure**, and prose reaches for the symptom. The fence is the only thing that distinguishes them,
which is the argument for running it before the commit rather than after — `guard_family` was run three
times in this iter's close and only the third was clean.

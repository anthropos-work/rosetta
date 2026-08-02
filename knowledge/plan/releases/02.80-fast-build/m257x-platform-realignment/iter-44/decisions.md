# iter-44 — decisions

## `D-M257x-44-1` — the ratchet is keyed `(fence, path, claim_id)`, deliberately WITHOUT the line number

The obvious key is `file:line` — it is what every RED prints, and it is what the audits' anchors use.
It is also wrong here, and the reason is §8 rule 6 rather than taste: under a line-keyed ratchet, adding
a paragraph *above* a known-and-accepted site changes its key, so the site reads as **new** and the
commit goes RED. The corpus is edited constantly for reasons unrelated to any claim; a fence that
reddens on those is a fence somebody adds to their `--no-verify` reflex within a week, and then it
protects nothing.

Line numbers are still reported on every RED — they are how a reader finds the site. They are simply not
part of its **identity**. `test_08_a_LINE_SHIFT_is_not_a_new_site` pins it and the
`key-carries-the-line-number` mutant proves the pin bites.

## `D-M257x-44-2` — the fence registry is DERIVED from disk, not held as a list in the runner

A post-condition runner naming its fences in a constant would be **exactly** the hand-maintained 4-tuple
§2 of the protocol deleted, one abstraction level up — and it would fail in exactly the same way: a
fence added later is simply absent, the runner reports OK, and nobody learns that a whole layer never
ran. iter-08 already measured the generalisation (*a fence only ever asserts about what it already
scans; an unclassified part is invisible by construction*).

So every `*_guard.py` on disk must declare a module-level `FENCE_KIND`, and one that declares none makes
the run **exit 2 naming its own filename**. Two sub-decisions:

- **Read statically, via `ast`, never by importing.** An import executes module-level code, so a guard
  that crashes on import would be indistinguishable from a guard that declares nothing — the
  swallowed-stderr shape (§5 rule 1) transplanted into a registry.
- **Two legal kinds, `postcondition` and `standalone`, and a typo is refused like an absence.** A third
  informal value would silently mean "not scanned"; `test_03` pins that `"postcondtion"` is a refusal.

## `D-M257x-44-3` — the baseline is a ratchet: lowerable, registrable, never raisable

`--accept` refuses to record growth **for a fence already in the baseline**, naming the sites that grew.
Without that refusal the file is a diary of whatever the tree happens to contain, which is worth nothing
against an induced-defect class.

A fence **absent** from the baseline is a genuinely different case: its first measurement is a
*registration*, not a regression, and grading 18 pre-existing sites as "induced" on a fence's first run
would cry wolf 18 times at exactly the moment the fence is least trusted. So it is admitted — and
**announced on stdout as a baseline rather than a pass**, with a `registered_at` sha in the file. This
matters directly for iter-45, which registers two more fences.

## `D-M257x-44-4` — every non-failure this module claims to report has a test that reads stdout, and a mutant that silences it

Straight from harden passes 7–9: two of the claim-twin fence's own honesty mechanisms **did not exist**
— the `UNMATCHABLE` reporting loop deleted clean with 15/15 green, and a docstring promised behaviour
(*"an audit in a new shape is reported as contributing zero rows rather than in silence"*) that was
never implemented. Both were in code written specifically to catch claims-without-measurements.

So the two non-failure halves here — *repaired since the baseline* and *newly registered fence* — are
each asserted through captured stdout (`test_16`, `test_17`), each duplicated in the JSON report
(`test_19`), and each has a dedicated silencing mutant (`repaired-report-goes-silent`,
`registration-report-goes-silent`, `json-drops-the-non-failure-halves`). **A reporting path with no
mutant is a docstring.**

## `D-M257x-44-5` — two vehicles, and the hook is the WEAKER one, said out loud

`--install-hook` writes `.git/hooks/pre-commit`. Git hooks are **per-clone and unversioned**: a fresh
clone has none, and nothing in a diff would ever reveal that. That is a real limitation, so it is stated
in the module docstring, in the protocol section, and in the installer's own output rather than left for
a reader to discover the hard way — the same failure mode as iter-01's git-ignored `rext.tag`, which
*"never appears in a diff and drifts unseen."*

The **suite** is therefore the load-bearing vehicle: `test_15` grades the live tree against the shipped
baseline on every `unittest discover`, in every clone, hook or no hook. The hook is a latency
optimisation on top — it moves the finding from "next test run" to "the moment of the repair" — and it
is scoped to commits that stage a published path so it cannot become the thing people bypass.

## `D-M257x-44-6` — the provider returns the fence's OWN live set, not a second opinion

`claim_twin_guard.postcondition_sites()` returns exactly the `live` list `main()` prints — same call,
same waiver handling. A post-condition that re-derived its own view could disagree with the fence a
reader had just run by hand, and then neither number would mean anything. `test_24` asserts the two sets
are equal, and `provider-reports-waived-sites-too` is the inversion mutant that proves the assertion
bites (it also reddens the live-tree test, because the three acknowledged retraction sites would
otherwise become permanent baseline entries).

## `D-M257x-44-7` — nothing was repaired, again

`D-M257x-42-3` still holds and for the same reason: the 18-blocker corpus is the only fixture with a
known answer key, and TOK-02 step 4 owns the repair. The one corpus edit in this iteration is **new
text** (the §8 subsection), and the post-condition was run over it as a post-condition — 18 sites,
byte-identical set, exit 0. That is the mechanism this iteration ships, used on the iteration that
shipped it.

## `D-M257x-44-8` — the induced-defect fixture is ASSEMBLED from the captured answer key, never written

`test_06` builds the anti-fixture by taking the 18 GREEN twins (a repaired tree — the fence is silent)
and dropping **one** captured RED file back in. A hand-written "induced defect" would be a guess about
what the class looks like; this one *is* the class, byte for byte at rosetta `48ca53c`. The test asserts
the repaired tree is silent **first** — without that control, a fence that reddens on everything would
pass it.

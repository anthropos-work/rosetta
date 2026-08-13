# iter-70 — decisions

## `D-M257x-70-1` — the routed class was 23; it is 4. Rule 32, applied to my own hand-off

iter-69 derived *"23 citations across 14 lines"* whose antecedent is a backticked path carrying no
line number, routed them as `FIX-M257x-iter69-pathless-antecedent`, and called the class **"a
prerequisite for the graded read."** Re-derived at this open: 6 were `shared_libraries.md:70`'s and
were repaired in iter-69 itself; of the remaining 17, **12 are PORT NUMBERS.**

The discriminator is mechanical and costs nothing: **does the antecedent file even have that many
lines?** `repos.yml` has 31, and the citation is `:5050`. `docker-compose.yml` has 271, and the
citation is `:8082`. `cors.go` has 110, and the citations are `:8000` and `:9000`.

| | n |
|---|---|
| impossible as a line number | **12** |
| line-plausible | 5 |
| …of which a port my heuristic mis-sized (`labsapi/client.go`, *"default `:7070`"*) | 1 |
| **real** | **4, across 3 lines** |

**And the heuristic's own defect is part of the record.** It sized the antecedent by `max()` over
every file on disk sharing that basename, so `labsapi/client.go` resolved to a 23,437-line
namesake and `:7070` came back "line-plausible." A sound version resolves the antecedent to **one**
file — which `platform_alignment_guard.resolve_citation` already does. Reported rather than
quietly corrected, because a discriminator that flatters its own recall is the failure mode this
milestone keeps finding.

§5 rule **32** says *re-derive the hand-off's numbers — including the orchestrator's*, on the
evidence of two M257x iterations that corrected an orchestrator-supplied fact. **This is the same
rule against my own hand-off, one iteration later**, and the number was off by a factor of six.

## `D-M257x-70-2` — `FENCE-M257x-iter69-citation-antecedent` must NOT be built as designed

iter-69 routed a fence to *"teach the shared citation parser that a backticked path with no line
number IS an antecedent for a following bare `:N`."* `D-M257x-70-1` is the evidence that decides it,
and the answer is **no**:

- the construct is **ambiguous by nature** — `` `docker-compose.yml` … `:8082` `` is a URL, and the
  corpus uses the port sense **12 times to the line sense 4**;
- so the proposed rule is a **3-to-1 false-positive generator**. §4 Trap A: a rule fitted to make
  the known-bad set visible, at the cost of drowning the true set, is not a fence;
- and the rule would be wrong in a third way — in `ai-readiness.md:45` the bare `` `:33` `` binds to
  the path **already named on the line** (`app-aireadiness-snapshot-loadmembers.yaml:42`), and only
  *looked* pathless because a different, line-less path follows it. The existing
  inherit-the-last-path rule was right there and my detector was wrong.

**The sound narrower version, named so it is not lost:** *a bare `:N` is a line citation only if the
antecedent resolves to exactly one file and `N ≤` that file's length at the adjudication ref;
otherwise it is not a citation and must not be counted as one.* That is decidable, has zero false
positives on all 12 port instances, and reuses `resolve_citation`. Routed as
`FENCE-M257x-iter70-line-or-port`, replacing the iter-69 design.

**iter-69's claim that this class blocked the graded read is retracted.** It is 4 citations, 3 of
which hold. The read is not held behind it.

## `D-M257x-70-3` — the four, adjudicated

| site | verdict |
|---|---|
| `security_compliance.md:71` — `org_membership.go:172-188` | **HELD.** At `app` `origin/main` `9d00a313` those lines are exactly `func (Membership) Policy() ent.Policy`, ending in `privacy.AlwaysDenyRule()`, as the sentence says |
| `ai-readiness.md:45` — `` `:33` `` | **HELD, and my detector was wrong.** It binds to `app-aireadiness-snapshot-loadmembers.yaml`, whose line 33 is `# v2.7 M254 RE-POINT: …` — the header the sentence quotes |
| `ai-labs.md:63` — `` `:7070` `` | **not a citation.** *"HTTP client to labs-api (default `:7070`, bearer `LABS_API_PLATFORM_TOKEN`)"* is a port |
| `studio-room.md:388` — `` `:36` `` / `` `:261-266` `` *"above"* | **one half repaired, one half UNMEASURED** |

**The `studio-room.md:388` split matters more than the edit.** The sentence read *"consistent with
`:36` and `:261-266` **above**"* — and `services/ai.py` is cited **nowhere earlier in that
document**, which is a corpus-internal fact needing no ref and no clone. The dangling cross-reference
is deleted (deletion > scoping edit > rewrite).

Whether `:36` and `:261-266` are the right *lines* is **UNMEASURED and stays so.** studio-room lives
in `anthropos-work/anthropos-studio-room`, which is **not cloned on this box**; the only copy here is
the untracked, CI-pulled `stack-demo/app/studio/` tree that iter-69 established is **in no `app`
commit at any ref**. A verdict read off it would have no ref, and P1 says a number whose ref is a
checkout is not a measurement. On that copy `:36` is `FAST = "fast"` and `:261-266` is an
error-handling block — suggestive, not decisive. Routed as `CHECK-M257x-iter70-studio-room-lines`
with the clone named as its precondition. **Recorded, not patched on speculation** — iter-68's
precedent, where two boundary defects were measured unreachable and left alone.

# iter-118 — decisions

## `D-M257x-118-1` — a URL is not a citation, and counting it as an unresolvable one corrupts the reach

`anchor_construct_guard` reported **675 resolved / 514 unresolvable** and named its unresolvable heads —
which is why nobody had looked *inside* them. **119 of the 514 are network addresses, not citations.**
`http://sentinel:8087` matches the qualified-anchor regex because `sentinel` is a path-ish head and
`8087` is a line number: `http:` ×87, `AUTHORIZATION_ADDRESS=http:` ×11, `GOTENBERG_URL=http:` ×11, plus
a tail of `VAR=http:` forms.

The fence was therefore reporting its coverage over a denominator **inflated by things that could never
resolve because they were never citations**. That is **iter-114's rule — a reach metric is settled by its
DENOMINATOR's provenance — arriving one layer over**, inside the guard that reports reach.

| | before | after |
|---|---|---|
| resolved | 675 | 675 |
| unresolvable | 514 | **395** |
| ratio printed | **none** | **675/1,070 = 63.1 %** |
| denominator provenance | — | **`citation-candidates-minus-non-citations`** |

**Counted and NAMED, never silently dropped.** A denominator that shrinks without saying so is the same
defect facing the other way, and this milestone has now been bitten from both directions.

## `D-M257x-118-2` — the anti-vacuity half is the load-bearing one when the fix REMOVES from a denominator

*"Exclude the non-citations"* is trivially satisfiable by a predicate that excludes everything — reach
goes to 100 % by deleting its own subject. So the controls are asymmetric on purpose: the mutation half
asserts each scheme form IS excluded, and the anti-vacuity half asserts **six real citation shapes stay
in** — including **`app/internal/httpclient/do.go`**, a genuine path containing the letters *http*. Plus
a subject-level control on the real corpus: exclusions must be **> 0** (the classifier still runs) and
**< 400** (it did not swallow the class).

## `D-M257x-118-3` — the sweep is COMPLETE as pre-registered, and the residual is named rather than closed

Against iter-117's sealed definition — *every class has a fence that enumerates its population
corpus-wide, stands at zero findings, and ships with controls that can fire*:

| class | population | findings | reach |
|---|---|---|---|
| 1 — intra-corpus citation | 1,520 | **0** | **1,520 / 1,520 = 100 %** |
| 2 — platform-source citation | 1,070 | **0** | **675 / 1,070 = 63.1 %** |

**The 36.9 % class-2 residual is not a gap that was hidden; it is one that is now measured.** 276 of the
395 are bare `` `:NNN` `` pins — the same shape class 1 measured as undecidable (a port, or a
continuation of a platform file named earlier). The rest are ambiguous bare basenames (`main.go` ×27 —
every Go repo has one, so the basename is not an identifier; `studioManager.go` ×8; a long single-file
tail). Closing them needs a repo-disambiguation rule, which is **routed, not attempted here** — the
scope-creep tripwire, with the iter's planned scope landed.

**So `TOK-08`'s trigger is now armed:** the next iter is the grading reading, and it grades
`P >= 19` (refuted) against `P <= 18` (working) on the baseline `P = 37` at `f581de09`.

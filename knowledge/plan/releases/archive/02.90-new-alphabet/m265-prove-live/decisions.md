# M265 — Decisions

## D-M265-1 — The content realignment lives in the REPLAY, not in the capture

**Decision.** Retired node-ids embedded in Directus content are repaired at **replay** time, per
stack, by `stack-snapshot/realign` — not fixed once in the captured snapshot.

**Why.** A snapshot is a faithful copy of a source at a moment; rewriting it during capture would
make the cache a *derived* artifact whose correctness depends on which taxonomy happened to be live
when it was taken. The same cached content is then wrong for any stack on a different canon. Doing
it at replay means the repair is always against **the taxonomy that stack actually has**, and the
cache stays a plain copy. It also costs nothing on a clean stack: a few counting queries.

**Cost accepted.** Every replay pays the scan. Measured at seconds on a content schema of this size.

## D-M265-2 — It runs on EVERY replay, not only the content surface

**Decision.** `realign` runs after every `stacksnap replay`, whatever the surface.

**Why.** Taxonomy and content are replayed by **separate invocations**, and both orders occur: a
cold bring-up replays taxonomy first, a re-replay of content onto a live stack is the other way
round. Gating the repair on "the content surface" would leave the first order broken, and gating it
on "the taxonomy surface" would leave the second. Running it unconditionally is what makes the
outcome independent of an ordering the caller chooses for unrelated reasons.

## D-M265-3 — Discovery, not a list

**Decision.** The set of columns to repair is read from the catalog every run, not maintained in
source.

**Why.** The list version shipped and was wrong within the hour: it named four columns, repaired all
four, verified clean, exited 0, and the next page load still failed on ids nested inside
`sequences.validation_evaluation_criteria` and `skill_paths.chapter_list`. **A hand-maintained list
of the places a value can hide is wrong the moment the content changes shape** — the same class as a
hand-maintained cache-key input list, which this project has been bitten by before.

**Consequence.** The rewrite must be shape-agnostic, so it is an exact-token substitution over the
document text rather than a path expression. That is what makes depth irrelevant.

## D-M265-4 — An unmappable id FAILS the bring-up; it is not dropped silently

**Decision.** If any retired id has no successor, `realign` returns an error naming the columns, and
the replay exits non-zero.

**Why.** The failure being guarded is *invisible*: a surface that renders empty while every liveness
probe passes. A half-applied realignment has exactly that signature, so "repaired what I could" must
not be reportable as success. Measured at M265 there were none — 187 of 187 were redirectable — so
the loud path costs nothing today and is the whole safety story tomorrow.

**Explicitly rejected:** dropping the unresolvable element. It is only well-defined for
array-of-objects shapes, and after D-M265-3 the package no longer knows a document's shape.

## D-M265-5 — "Unprovisioned" and "broken probe" are different verdicts

**Decision.** A content schema with **no tables** is a stated SKIP. A schema **with** tables and no
json column remains a failure.

**Why.** The first version collapsed them, and a cold bring-up — where taxonomy replays before
Directus is bootstrapped — recorded a perfectly good replay of 55,116 rows as a FAILED surface. **A
false alarm on the most common path is worse than no alarm**, because it teaches the reader to
discount the real one. The anti-vacuity rule was right; its trigger was too broad.

## D-M265-6 — The claim-census ratchet is NOT re-baselined here

**Decision.** Six corpus files exceed their unevidenced-assertion baseline; M265 records this and
routes it to `/developer-kit:close-release` rather than running `--update-baseline`.

**Why.** Five of the six grew during earlier v2.9 milestones, so this is a release-level question,
not a milestone one. More importantly, moving a ratchet baseline is a **target change**, which is the
user's call — repairing a stale reference is not. Doing it quietly at the end of the last milestone
is exactly how a ratchet stops meaning anything.

## No TOK chain — Before-You-Start case (b)

This milestone has **no `TOK-*` entries** and never had one. `iter-01/` existed as an empty scaffold
before the first iter ran, so `/developer-kit:build-mstone-iters` Phase 0 rule 1 (iter-01 = bootstrap
tok) could not fire — the bootstrap-less shape the skill calls case (b).

iter-01 therefore ran as a **tik planning from `overview.md` + the protocol doc directly**. Any future
iter of this milestone should name "no TOK chain (case (b))" as its active-strategy reference rather
than hunting for a TOK entry that does not exist.

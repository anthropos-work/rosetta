# iter-190 — decisions

## `D-M257x-190-1` — the routed selector was PROSE, and it was replaced rather than run

iter-189 routed its residual with the selector *a function whose docstring or name claims to be the same
derivation as another*. Re-surveyed across `stack-core`, that phrasing returns **one** hit: iter-189's own
founding case. **A selector that finds only the case it was written from is not an enumeration** — it is
the case, restated.

Replaced with a mechanical one that needs no prose: *a module-level constant used by BOTH a
filesystem-reading function and a git-reading function*, derived by AST over all 42 modules. **6 pairs**,
all in `platform_predicate_guard`. This is the substitution `§8`'s *a class is closed by an enumeration
that keeps running* asks for, and it is recorded because the prose selector was mine, one iter old.

## `D-M257x-190-2` — the census returned ZERO divergences, and proving the instrument is what found the defect

All 6 pairs agree today, measured: `parse_compose` → 7 services vs `compose_counts_at(None)` → `(5, 7)`;
`repos_yml_history` ⊇ `parse_repos_yml`, 0 current entries missing. `§9` then applies — *a census
returning ZERO must prove its instrument* — and the proof is what surfaced the finding.

**The pair the census structurally cannot see is the one that is broken.** `_parse_one_compose` and
`compose_counts_at` read the **same construct out of the same file** with **different regexes**, so they
share no constant and a shared-constant census misses them by construction. That is iter-184's rule (*a
fence's POPULATION is a registry too*) committed by the instrument written this iter — which is exactly
where iter-184 said to look.

## `D-M257x-190-3` — the narrow recogniser is the wrong one, and the direction of the error decides it

`_SVC_KEY` was `^  ([A-Za-z][A-Za-z0-9_.-]*):` — **letter-initial**. `_COMPOSE_SERVICE_KEY` was
`^  (?P<name>[A-Za-z0-9_.-]+):`. Over a 9-name candidate table they disagree on **5** (`3d-render`,
`_internal`, `-legacy`, `.hidden`, `9front`).

Compose's own service-name charset admits a leading digit, so the narrow one is wrong on the merits. The
direction settles it independently: `parse_compose` builds the topology **G1/G7/G8 grade profile
membership against**, so a service it cannot see is *absent* — claims about it read UNREACHED rather than
graded — while G10 counts it. **Under-count on the side that grades.**

Unified on one `_SVC_NAME` charset both patterns are built from, rather than making the two literals
match: agreement today is not the property; **sharing the source is** (iter-177).

## `D-M257x-190-4` — both escalation conditions were checked, with the measurements

`overview.md` pre-registered two. **(a)** *If widening makes `parse_compose` see a NEW service in a real
clone, this is a live under-count.* Measured before and after against `stack-demo/platform`: **7 services,
identical set**, and `compose_counts_at` unchanged at `(5, 7)`. **(b)** *If the guard's own suite goes RED,
the narrow charset was load-bearing.* **184 passed · 0 failed.** Neither fired; the defect is latent.

## `D-M257x-190-5` — the fixtures are synthetic, deliberately

The live comparisons could have run against `stack-demo/platform` and `stack-demo/app`. They do not: **a
comparison that only runs where a clone happens to exist is a comparison that stops being run** (`§8`),
and the arms would then pass by absence on any bare checkout — the fail-open shape this milestone has
paid for twice (iter-179's unrun battery, iter-182's `1 passed, 4 skipped`).

The compose fixture carries `9front` **on purpose**: without a non-letter-initial name every arm here is
green under either recogniser, which is the `§9` point made concrete — and mutation M6 confirms it, the
narrow recogniser turning the fixture RED where a realistic fixture would have stayed green.

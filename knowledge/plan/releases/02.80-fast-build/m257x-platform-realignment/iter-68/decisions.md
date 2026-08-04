# iter-68 — decisions

## `D-M257x-68-1` — a guard that resolves a citation against a CHECKOUT has no ref, and its verdict is not a measurement

Three guards — `platform_predicate_guard` (G6's consumer side), `platform_alignment_guard`
(assertion F), `anchor_construct_guard` (every anchor into a clone) — all read the cited file with
`read_text()` on whatever the clone happened to have checked out. None of them reported which file
that was.

That is not a latent tidiness problem. It is measured, and it is decisive:

| | at the demo's pinned build ref `b948604` v1.366.0 | at `origin/main` `9d00a313` v1.367.0 |
|---|---|---|
| `STORAGE_RPC_ADDR` (G6) | **mid-fold**, 6 app read sites | **unconfigured** — no reader at all |
| `app/main.go` length (F) | **1361** lines | **1569** |
| `app/internal/storage/service.go` | does not exist | exists |
| corpus verdict (anchor guard) | **RED, 4 findings** | **GREEN** |

**Same guard, same corpus, opposite verdicts, 56 commits and one working morning apart.** P1 —
*state your refs* — is not satisfied by a number whose ref is a checkout.

**Decision.** Every guard that resolves or reads a cited file does so **at a named ref**, and
**reports it**. The default is `auto`: prefer `origin/main` — the ref a cold `make init` clones and
the ref the exit gate names — falling back to `HEAD`. `worktree` restores the old behaviour, by
name. **A ref the caller named that does not resolve is UNMEASURED, never silently substituted**
(§5 rule 7).

Two corollaries that were both live defects, not hypotheticals:

- **Existence is decided at the same ref the content is read at.** Resolving against the checkout
  and reading against a ref is the worst of both: a file *born* in the advance
  (`app/internal/storage/service.go`) read as `unresolvable head 'app'` — an instrument gap wearing
  the costume of a corpus defect.
- **An untracked worktree file is in no ref, and correctly leaves scope.** The demo box carries an
  untracked `studio/` copy of `anthropos-studio-room`; three corpus citations (`app/studio/…`)
  resolve against it and against no commit of `app`. They are studio-room's, not app's.

## `D-M257x-68-2` — release 09.00 "support-in-app" landed; `mid-fold`'s only instance lived four iterations

Measured at platform `0dab54d` / `app` `9d00a313` v1.367.0 / `storage` `63bffc8` / `messenger`
`a0ec933`:

- **storage** — prod `service_desired_count = 0` (`storage/terraform/main.tf:38`); `app` serves
  object storage in-process (`main.go:471`, `:472`, consumed `:494`, `:1048`); `STORAGE_RPC_ADDR`
  read by `main.go` and by **none** of the three `cmd/` tools.
- **messenger** — prod `service_desired_count = 0` (`messenger/terraform/main.tf:29`); `app` imports
  `internal/messenger/{flow,adapters,sender}` (`main.go:15`, `:61`, `:62`) and **takes over
  messenger's own Redis consumer group** (`:1387`, `:1423`) rather than merging its handlers.

Both repos stay in `repos.yml` and both stay startable in compose — the rollback path, exactly as
`cms` and `jobsimulation` are kept.

**The eighth vocabulary token, `mid-fold`, was built at iter-64 for `storage` and its only instance
was gone by iter-68.** The token stays: the fold program is not finished, and a state you can only
name *after* you need it is the state you will get wrong. The map now records that **no row carries
it**, which is a different statement from the token not existing.

## `D-M257x-68-3` — the routed figure was stale when it was routed, and a corpus repair is what staled it

iter-63 routed *"the 68 non-mainline citations, WHOLE"*. Re-running **its own enumerator** against
today's corpus, same app ref:

| reading | sites | distinct | files | mainline | non-mainline |
|---|---|---|---|---|---|
| iter-63, as recorded | 104 | 86 | 22 | 18 | **68** |
| iter-68, same instrument | 123 | 96 | 22 | 32 | **64** |

§5 rule 34 says a corpus repair moves the corpus's own line numbers. **This is its sibling: a corpus
repair also enlarges the corpus's own citation class.** iters 63–67 wrote 10 net-new citations of
their own — the two-sided `mid-fold` row alone added six, and every one of those six died with the
row. A routed count is a snapshot of a population that the routing iter is itself growing.

## `D-M257x-68-4` — `CHECK-M257x-iter63-quoting-a-retired-token` is a WINDOW bug, not a policy question

Third instance in three iterations, and it fired on the very edit that was recording the rename.
G1 already had a negation discriminator (`_NEGATED`, adjacent particle). Its **window was one line**,
and prose wraps: *"…and there is no\n> `graphql` profile any more"* puts the particle on one line and
the noun phrase on the next, behind a blockquote marker.

**The second window bug of this milestone wearing a policy's name** — `D-M257x-63-1` was the first,
and its lesson was written down as *"a 'policy hole' is often a window bug wearing a policy's name."*
Written down, and then not applied to the next one.

- The negation prefix is now the preceding prose line + this line up to the claim, on `_pin_window`'s
  block boundaries. **Adjacency is unchanged** — two words — so a negation two lines up, or about
  something else, launders nothing.
- A leading blockquote/list marker is stripped: layout, not a word standing between them.
- The one remaining live site was a **postfix** denial (*"a `storage` profile that never did"*).
  **Rephrased to the prefix form rather than teaching the fence English.** A retired token spelled in
  profile position is exactly what G1 exists to catch; the corpus can always say it the other way
  round. Fitting the rule to that sentence would have been §4 Trap A.

## `D-M257x-68-5` — a fence's reach hole surfaced a defect class the fence could not see: the `* **Profile**:` bullet

Repairing messenger surfaced that **eight service docs carry a `* **Profile**: …` bullet** and G1
reaches **none** of them — its three constructs are a command, a table cell and a noun phrase, and a
list bullet is none of those. Seven of the eight were wrong, and **all seven named
`graphql`** — re-derived per bullet from `git show HEAD`, because the first draft of this entry said
*five* and was itself an unchecked count in the iteration that exists to stop those. `0dab54d`
renamed that profile `core`. Two of the seven (`roadrunner`, `jobsimulation`) additionally named a
profile for a service with **no compose entry at all**, both deleted by `d11a403`; `storage`
additionally named a `storage` profile that never existed. The eighth, `sentinel.md`
(*"always on — no `profiles:` declared"*), is right, and is also the only one naming no token.

All seven repaired from `docker-compose.yml` @ `0dab54d`. **Widening G1 to the bullet construct is
routed** (`FENCE-M257x-iter68-profile-bullet`) rather than done here — iter-67's lesson is to build
the fence in the iteration *after* the one that names the class, with a design rather than a title.

**Seventh time in this milestone that a GREEN reading turned out to be a reach limit.**

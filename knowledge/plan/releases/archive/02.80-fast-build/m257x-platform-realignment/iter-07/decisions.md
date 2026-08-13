---
milestone: M257x
iter: 07
---

# iter-07 — decisions

## D-M257x-8 — the replay schema is DERIVED from the target, never declared

The pre-compute offered two shapes for `REPOINT-M257x-cms-similarity-writes` and deliberately refused to
pick. **Adopted: the derived replay-time resolver. Rejected: a second declared constant.**

**Why the constant loses even though it is two lines.** `simembeddings.Schema = "cms"` is read by both the
prod CAPTURE and the stack REPLAY, and those two now legitimately disagree — permanently, not transiently
(`D-M257x-7`: the write side moves, the prod-read side stays). Adding `ReplaySchema = "public"` would encode
*today's* answer to a question the platform has changed three times in three releases (skiller v2.1,
skillpath v2.7, jobsim + cms v2.8) and has already scheduled changing again (v9.0 folds `storage` +
`messenger`, PRs open). It is the same hand-maintained-list defect this milestone exists to end, and its
failure mode is the bad one: the bring-up reports success.

**What was built instead.** `replay.ResolveTargetSchema` — a pure decision over the candidate schemas the
TARGET reports holding *all* of a surface's tables (`pg.SchemasHoldingAllTablesSQL`, one catalog query):

- declared schema holds them → **identity** (taxonomy and directus are untouched, by construction);
- exactly one other does → **remap**, announced loudly on stdout;
- none → `ErrSurfaceNotOnTarget` (exit 4, a provisioning problem);
- two or more → `ErrAmbiguousTargetSchema` **naming them** (exit 1 — *not* a provisioning problem; nothing
  the operator provisions will fix two schemas each holding a full copy).

**The three things that were deliberately NOT built**, each because it is the same defect in a different
costume: an allow-list of application schemas (Trap A in miniature — tune it until it stops catching what it
exists to catch); a preference for `public` (a constant with extra steps); a fallback to the declared schema
when the lookup errors (a probe that satisfies itself, §5 rule 7).

**Cost, measured before committing:** three types implement `Replayer`, and `replay.Run` had 18 call sites.
The parameter was made **positional and required** rather than an option, so the compiler forced all 18 to
state which behaviour they meant. Identity is spelled `Identity()` at every site that wants it — never a
silent zero-value default.

**What this buys at the next fold:** nothing to edit. `storage`/`messenger` moving into `public` is followed
automatically, and if they land somewhere ambiguous the tooling stops and names the candidates instead of
guessing.

## D-M257x-9 — the digest probe and the replay resolve as ONE construct

The pre-compute's "thing not to miss (a)": `stacksnap`'s pre-replay probe reads the manifest schema too, so
moving only the replay leaves the surface skipping at `rc=4` before a single row is copied — with a diff that
looks complete in review.

Rather than move both and rely on review to keep them together, they are **one function**:
`resolveThenProbe(ctx, prober, surface)` resolves, then probes the digest on `ts.For(surface.Schema)`
computed *inside itself*. **There is no parameter for a caller to supply**, therefore no way to supply the
wrong one.

This is `§8 rule 1` — assert against the construct, not the text — applied to the *implementation* rather
than to a test. A fence can only catch a drift that has already been written; a construct that cannot express
the drift is strictly better, and here it was cheap. `TestResolveThenProbe_ProbesTheRESOLVEDSchemaNotTheDeclaredOne`
still exists as the fence, mutation-verified, because the construct could later be refactored apart.

## D-M257x-10 — the resolution is LAZY, and the reason is an exit-code contract

Recorded because the eager version was written first, shipped RED, and the tests explained why better than
the code did.

`--schema-version` means *"do not ask the stack for the digest"*. The store resolve that follows it is a pure
local-filesystem question, and three pre-existing tests document — in their own comments — that a
**cache-miss verdict must be reachable without a live database**. Resolving the schema eagerly broke that,
and did so in the specific way `fix16` exists to prevent: it turned exit 5 (*an empty/stale cache; capture
fixes it*) into exit 4 (*the stack is unprovisioned; a capture cannot help*). An operator reading exit 4
would go and re-provision a stack that was fine.

So the resolution is memoized and runs at the first point that genuinely needs it — before the probe when
the probe runs, otherwise immediately before `replay.Run`. Both paths share one memo, so the two consumers
can never see two different answers.

**The transferable part:** an exit code is a contract with a human about *what to go and do next*. Adding a
new precondition check ahead of an existing decision tree can silently re-answer that question, and the
symptom is not a crash — it is an operator doing the wrong repair. The existing tests caught it only because
somebody had written down *why* the fixture was shaped the way it was.

## D-M257x-11 — a mutation that does not compile is not a RED fence

Method note, promoted to the protocol doc (`platform-alignment.md` §8 rule 5).

M3 of this iter's mutation battery (ambiguity resolved by picking `candidates[0]`) removed the last use of
the `strings` import. The package stopped **compiling**, `go test ./...` returned non-zero, and the harness
printed `RED (good)` — with an **empty list of failing test names**, which is the tell. A compile break
proves the mutation was applied; it proves nothing whatsoever about whether any fence noticed the behaviour
change.

Re-run with a compiling mutation (`_ = strings.Join(candidates, ", ")`), gated on an explicit `go build`
BEFORE the test run, the real fence fired by name.

Same family as `§5 rule 8` (*a check that SKIPS reads exactly like a check that PASSES*) and `§5 rule 1`
(*never let a search's stderr go unread*): the run failed for the wrong reason, and a failure for the wrong
reason is indistinguishable from a failure for the right one unless you read what actually failed.

**Rule: build the mutant before you trust its RED, and name the test that went red.** A mutation battery that
reports only exit codes can sign off on a fence that does not fence.

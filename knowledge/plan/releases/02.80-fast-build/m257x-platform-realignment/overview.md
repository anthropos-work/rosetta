---
milestone: M257x
title: "platform re-alignment"
release: v2.8 "fast build"
milestone_shape: iterative
status: archived
last_updated: 2026-08-11
closed: 2026-08-11
close_status: closed-incomplete
close_authority: "USER RULING — TOK-09 (2026-08-11). NOT a gate-met close. Clauses 1-4 met; clause 5 OUT OF SCOPE by that ruling and never met."
created: 2026-07-31
iteration_protocol_ref: corpus/ops/platform-alignment.md
exit_gate: "Against platform @ **origin HEAD** (never a pinned pre-drift commit): (1) a cold `demo-down --purge` + `demo-up` on the **dev host** (D-v28-15: local to the new Mac; odysseus retired, billion demo-only) reaches `autoverify green:true / 0 warnings` across **3 consecutive cycles**; (2) the **full Playthrough suite passes on that stack** (30 live / 0 failing / 0 error) — presence AND function, so a green bring-up cannot mean an empty world; (3) a **checked-in migration-status map** covering every service the platform has ever had — state ∈ {live-standalone, merged-into-app, decommissioned, net-new} — **each claim cited to platform source** (commit sha or file:line) and **machine-fenced against `repos.yml`** so it cannot silently drift; (4) **zero rext writes to a schema the platform no longer creates**, asserted by a FENCE that is watched going RED, not by inspection; (5) KB-fidelity audit **GREEN, or YELLOW with 0 blockers**, over `corpus/services/**` + `corpus/architecture/**`. Clauses 1/2/4 are rosetta-extensions; 3/5 are the rosetta corpus — both repos are in the gate by construction."
re_scope_trigger: "If TWO consecutive full-alignment attempts are invalidated by new platform commits landing mid-milestone — i.e. the target moves faster than we can track it — STOP and escalate. The answer then is a pinning-and-tracking POLICY (how we choose a platform ref, how we notice it moved, who re-points), not more alignment work. Grinding against a moving target is the failure mode this trigger exists to catch. NOTE the specific arithmetic that would signal it: origin `repos.yml` moved at 2026-07-29T14:14Z, **39 minutes after** the mirror-drop commit — the platform team ships coordinated multi-repo changes, so 'we re-pointed everything' has a short shelf life unless the policy exists."
---

# M257x — platform re-alignment

**The team is migrating the microservices back into `app`, toward a more monolithic design. Nobody on our
side knows how far that has got.** This milestone finds out, writes it down where it cannot rot, and makes
**both** repos work against the platform as it actually is.

## Why this exists, and why now

**It is not speculative — M257 hit the wall and stopped.** Its iters 02–03 found that **the demo could not
have been built cold on any machine for four days**, and that the gate's own health check was querying a
table the platform had **deleted**, with a `|| echo 0` turning the error into `0` — a number that reads
exactly like *"the demo seeded, just thinly."* Two blockers were fixed; the third is not a fix, it is this
milestone:

> Platform `repos.yml` @ origin `236771f10` (2026-07-29T14:14Z) sets `migrations: false` for **cms**,
> **jobsimulation** and **roadrunner**, stating *"`app` is the ONLY repo with migrations to run… they own no
> local schema."* A fresh stack therefore **never creates the `jobsimulation` schema** — while rext still
> writes **~15 `jobsimulation.*` tables**. It is latent only because our clones are days stale.

**And our local ground truth is already provably stale:** `stack-dev/platform/repos.yml` still lists
**`skillpath`** — decommissioned back in v2.7 — and still marks cms/jobsimulation `migrations: true`.

**This is the THIRD occurrence of one class.** v2.1 (skiller → app) broke the seeder. v2.7 (skillpath → app)
broke it again and the corpus asserted skillpath as live Tier-1 in ~30 files. Now jobsimulation. Each time we
re-derived the fix from scratch. **A recurring class with no written procedure is a class that will recur** —
so this milestone's protocol doc is a deliverable, not a formality.

## Scope

**Investigate → consolidate → re-point → prove.** Both repos, in one milestone, because splitting them would
let the corpus describe a platform the tooling cannot build.

### rosetta-extensions (the tooling — "probably most of the stuff doesn't run anymore")
- Every write path re-pointed off schemas the platform no longer owns (the `jobsimulation.*` set is known;
  **assume it is not the only one until measured**).
- A **fence** that fails when rext writes to a schema `repos.yml` says the platform does not create. Watched
  going RED. This is the guard that would have caught all three occurrences.
- Seeding, snapshot, verification, playthroughs, cockpit: each proven against origin HEAD, not assumed.

### rosetta (the corpus)
- The **migration-status map** — every service, its state, cited to platform source, fenced against `repos.yml`.
- `corpus/services/**` (29 docs) and `corpus/architecture/**` reconciled to what the map establishes.
- **Net-new repos** the org has grown that appear in neither `repos.yml` nor the corpus — the user asked for
  this explicitly, and nothing currently looks for them.
- **`Delivers → corpus/ops/platform-alignment.md`** — the re-ground procedure itself: how to detect
  platform/tooling drift, how to re-point, how to fence it. Authored by iter-01; it is this milestone's
  `iteration_protocol_ref` and does not exist yet, which is precisely the gap.

### Out
- Making the bring-up **fast** — that is M257, paused behind this.
- Baking Playthroughs into the bring-up — M258.
- **Any platform-repo edit.** The v2.8 constraint holds: 0 platform edits, sha-pinned `demopatch` or an
  rext-owned file.

## What is already known (inherited evidence — do not re-derive)

| fact | source |
|---|---|
| `repos.yml` @ origin `236771f10`: `migrations: false` for cms/jobsimulation/roadrunner; *"app is the ONLY repo with migrations"* | M257 iter-03 |
| Local `stack-dev/platform/repos.yml` still lists **skillpath**, still `migrations: true` for cms/jobsimulation | measured 2026-07-31 |
| rext writes ~15 `jobsimulation.*` tables | M257 iter-03 |
| Dropped `local_*` mirrors broke **6 seeders** — the re-point was **34 sites / 20 files** (fixed in M257) | M257 iter-03 |
| `app`/studio had **no rext acquisition path**, broken since 2026-07-27, invisible because nobody ran a cold cycle (fixed in M257) | M257 iter-03 |
| `app` was **~386 commits** ahead of the pin; skillpath fully decommissioned (M501–M507 → 3 subgraphs); **jobsimulation mid-merge — "the next shoe"** | v2.7 |
| `app` grew undocumented domains: coursebuilder · AI Labs + credits · askengine · a server-owned academy | v2.7 |
| Corpus says studio-room is CMS-only in **5 places**; nothing records that `app` embeds it | M257 `DOC-M257-studio-in-app` |

## Open questions iter-01 must answer

1. **Where is the migration actually?** Which services are still standalone, which are merged, which are husks?
2. **Is `jobsimulation` merged, mid-merge, or unchanged?** v2.7 called it "the next shoe"; `repos.yml` already
   says it owns no schema. Those are not the same claim.
3. **Are there NET-NEW repos** in the org that appear in neither `repos.yml` nor the corpus?
4. **Does the 3-subgraph count still hold?**
5. **How much of rext actually still runs?** The honest prior is "less than we think" — B1 and B2 were found by
   the first cold cycle anyone ran in four days.

## Shape

`iterative`, and not by preference. **A fixed `In:` list would be speculative** — the deliverable set depends
entirely on what the investigation finds, and the one thing we know for certain is that our picture of the
platform is out of date. Committing to a checklist now would be committing to the stale picture.


## ⏸️ MACHINE MOVE — 2026-07-31, mid-run

Work **stopped on odysseus and on the old laptop** (`D-v28-15`). Both repos move to a new Mac with a **local**
dev stack. **iter-01 is KEPT** — it is committed and pushed (`99f0aca`, rext tag `fast-build-m257x-iter-01` on
origin) and its output is *platform knowledge*, not machine state. Re-deriving it would be the very
"re-derived from scratch each time" waste this milestone exists to end.

**What iter-01 established (carry it forward, do not re-measure):**
- The consolidation is a **5-service PROGRAM**, with **the next two folds already in open PRs**.
- **Root cause of the recurring class: pinning silently disables rext's own drift detection.** 11/11 clones
  report `behind: null` while the log claims *"provably fresh"*, and the pin's source of truth
  (`.agentspace/rext.tag`) is **git-ignored** — so it never appears in a diff and drifts unseen.
- **Local mechanism:** `demo-stack/migrate-demo.sh:81-85` **creates the legacy schemas itself** and `:106`
  atlas-applies a **hand-maintained 4-tuple**, never consulting `repos.yml`'s `migrations:` flag. Someone
  edited it for skiller; nobody did for jobsim/cms. **Time bomb:** when the legacy repos leave the clone set,
  `[ -d ] || continue` silently skips them and **13 write targets 42P01 at once**. The canary is already
  visible — skillpath sits in the tuple but is absent from origin `repos.yml`.
- The real jobsimulation surface is **12 tables (9 write / 3 read-only)**, not ~15 — two inherited names were
  comments, one explicitly labelled a red herring.
- **5 inherited/audited claims refuted by measurement**, including one from the KB audit that had **inverted**
  the guard it described.
- `corpus/ops/platform-alignment.md` **authored** (16 KB) — this milestone's `iteration_protocol_ref`, and it
  proved the corpus index guard went RED on it before the index row was added, then GREEN after.

**Gate read at the stop: 0 of 5 clauses met.** Clause 1 was **BLOCKED** — `/demo-up` aborts on a FATAL rext-pin
mismatch, SoT 63 commits behind `main`. **That blocker follows us:** `rext.tag` is git-ignored, so the new Mac
starts with no pin at all and must create one deliberately.

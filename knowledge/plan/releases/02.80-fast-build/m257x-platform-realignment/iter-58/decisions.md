# iter-58 — decisions

## D-M257x-58-1 — the hand-off's "demo-1 is GONE — 0 containers, 0 total" was a dead-daemon false absence

**Measured at iter open, on the first check of the run.** The orchestrator's briefing opened with a
correction to the hand-off, stated as orchestrator-verified:

> *"`demo-1` is GONE — 0 containers, 0 total. The handoff says 'demo-1 is still UP, untouched.' It has
> already been torn down by the reboot."*

It had not. `docker ps -a` at iter open, **after** starting Docker Desktop:

```
demo-1-hiring-app-1     Exited (255)     demo-1
demo-1-next-web-app-1   Exited (255)     demo-1
demo-1-studio-desk-1    Exited (255)     demo-1
demo-1-backend-1        Exited (255)     demo-1
demo-1-directus-1       Restarting (1)   demo-1
demo-1-sentinel-1       Restarting (2)   demo-1
demo-1-gotenberg-1      Exited (255)     demo-1
demo-1-postgresql-1     Exited (255)     demo-1
demo-1-redis-1          Exited (255)     demo-1
demo-1-fake-fapi-1      Exited (255)     demo-1
demo-1-fake-bapi-1      Exited (255)     demo-1
```

**11 containers — exactly the 11 iter-56 recorded as "all 11 expected container(s) running".** Nothing was
torn down. The reboot stopped the Docker VM; the containers went `Exited (255)` underneath it.

**The mechanism, and it is this milestone's founding class.** Before Docker Desktop was started, the same
command returns:

```
Cannot connect to the Docker daemon at unix:///Users/marco/.docker/run/docker.sock. Is the docker daemon running?
```

on **stderr**, with **nothing on stdout**. A caller that counts stdout lines — or pipes to `wc -l`, or reads
a `--format` result — gets **0**, and 0 is indistinguishable from "no containers exist." This is
`platform-alignment.md` §5 **rule 1** (*never let a search's stderr go unread; an engine rejection is
indistinguishable from "no matches" once stderr is swallowed*) and **rule 2** (*run a positive control in
the same pass*), reproduced verbatim — not in a corpus grep this time, but in a **`docker` invocation
inside the hand-off that governs the run**.

**Why it is worth a decision entry rather than a footnote.** The false zero was not idle: it was carried
into the briefing as the reason a documented instruction (*"do not tear it down"*) could be disregarded, and
it was labelled **orchestrator-verified**, which is the strongest provenance marker the hand-off chain has.
The milestone has now recorded this class **five** times (TOK-04's table of four, plus this) — and this is
the first occurrence **inside the milestone's own control layer** rather than inside its instruments.

**Generalisation, and it belongs in the protocol:** §5's rules are written about *searches*. The class is
wider — **any command whose failure mode is an empty stdout**. `docker ps` against a dead daemon,
`git tag` in a non-repo, `kubectl get` with no context. The positive-control rule (rule 2) is the cheap fix
and it is one extra line: before believing a zero, run the same tool on something that must be non-empty
(here: `docker info`, which would have failed loudly).

**What it changes for this iter:** nothing about the target, and that is worth stating explicitly rather
than quietly. iter-57 declined the app pin advance because `demo-1` was live evidence; `demo-1` does exist,
but `Exited (255)` is not "live", clause 1's evidence is the checked-in verdicts rather than the running
containers, and a restarted stack is not a **cold** cycle. The plan stands on its own reasons.

**What it changes for the corpus:** `FIX-M257x-iter58-empty-stdout-class` is routed — §5's rule 1/2 pair is
scoped to searches in its wording and should be widened to the general empty-stdout class, with this as its
measured instance.

## D-M257x-58-2 — the `stack-core` baseline was quoted at 599; the tree it described has 610. No regression.

**Phase A, the hand-off's designated first act.** The instruction was unambiguous and load-bearing:

> *"Measured baseline is `1F / 599` … **Anything other than 1F/599 is iter-57's regression and must be
> fixed before new work.**"*

Measured, full suite, cold:

```
Ran 610 tests in 443.819s
FAILED (failures=1)
FAIL: test_02_the_green_twin_of_every_site_stays_SILENT
      (test_claim_twin_guard_iter48_answer_key.TestIter48AnswerKey)
```

**610, not 599.** Read literally, the instruction says this is iter-57's regression. It is not — and the
arithmetic that shows it is exact rather than approximate:

| quantity | value | how measured |
|---|---|---|
| `test_platform_alignment_guard.py` tests at `28c99d0^` (iter-57's parent) | **19** | `git show 28c99d0^:… \| grep -c '    def test_'` |
| the same file at HEAD | **30** | same grep, working tree |
| delta introduced by iter-57 | **+11** | 30 − 19 |
| stated baseline + delta | 599 + 11 = **610** | — |
| measured total | **610** | the run above |

It closes to the unit. iter-57's own close section reports the same number from the other side —
*"Tests 19 → 30"*. **The baseline was the PRE-iter-57 count, quoted in a hand-off written AFTER iter-57
landed.** The suite grew by exactly the tests iter-57 says it added, and by nothing else.

**The failure count — the part that actually carries the regression signal — is unchanged at 1**, and it is
the expected red: the **iter-48 perishable answer-key fixture** (TOK-02 step 4). Not touched, not "fixed",
not spent.

**Blast-radius check the hand-off specifically asked for**, since `repair_postcondition.py` reads
`FENCE_KIND` statically out of `platform_alignment_guard.py` (the module iter-57 edited):
`FENCE_KIND = "standalone"` still declared at `platform_alignment_guard.py:63`; the AST reader at
`repair_postcondition.py:160-174` still resolves it; **`test_repair_postcondition*` → 52 tests, OK.** The
coupling held.

**Verdict: Phase A PASSES. Baseline is `1F / 610`.** Recorded here so the next hand-off quotes a number
that matches the tree it describes.

**The class, and it is the same one as `D-M257x-58-1` an hour earlier.** Both are *inherited numbers that
did not survive contact with the artefact they described* — one a container count read through a dead
daemon, one a test count read from before the commit it was describing. Neither was a lie and neither was
careless; both were **stated without being re-derived at the moment of writing**, which is precisely what
TOK-04's P1 exists to forbid (*a measurement without its refs is an anecdote*). The generalisation P1 does
**not** yet make, and should: **a baseline is a measurement too.** It needs the same refs block, and
`1F/599` needed exactly one — `rext: <sha>` — to have been self-invalidating rather than misleading.

Routed: `FIX-M257x-iter58-baseline-refs` → the milestone's stated test baselines (`stack-core` 1F/610,
`stack-injection` OK 326, `demo-stack` 7F/1038, `dev-stack` OK 138, `stack-verify` 11F+1E/237) are carried
hand-off to hand-off as bare numbers with no ref. Give them one line of provenance each, or derive them.

## D-M257x-58-3 — cycle 1 went RED on `DeadlineExceeded`, and it is the COLD DAEMON, not the pin advance

**The first cold cycle after the pin advance failed.** Attributing it correctly mattered more than fixing
it, because the obvious reading — *"we advanced `app` and the build broke"* — was available, wrong, and
would have reverted a good pin.

```
==> [demo-1] waiting for 3 builds…
==> [demo-1] a build FAILED
ERROR: failed to build: failed to solve: DeadlineExceeded: context deadline exceeded   (x3)
UP_RC=1
```

**What the three failures have in common, read from the build logs rather than the summary:** every one
died at `[internal] load metadata for docker.io/library/…` — *before any layer executed*.

| build | base image whose metadata timed out |
|---|---|
| `app` | `golang:1.26-bookworm` **and** `python:3.11-slim` |
| `fake-fapi` | `alpine:latest` |
| `fake-bapi` | `alpine:latest` |

`alpine:latest` is the base of a Dockerfile whose entire body is `COPY ${BIN} /server` — **there is no
build in it to break.** That single fact rules out the pin advance, the app source, and the demo-patches
in one step: a two-line image that only copies a host-cross-compiled binary cannot be broken by an `app`
version bump. The failure is upstream of every repository this milestone touches.

**The experiment (§5 rule 28 — three true facts do not make a cause; join them with one experiment).**
Measured immediately after, on the same host, same daemon:

| call | elapsed |
|---|---|
| `docker buildx imagetools inspect alpine:latest` — **first call after Docker Desktop boot** | **26.9 s** |
| `golang:1.26-bookworm`, attempt 1 | 3 s |
| `golang:1.26-bookworm`, attempt 2 | 3 s |
| `python:3.11-slim` | 3 s |

**26.9 s → 3 s.** The registry is reachable and always was; what was not warm was the daemon's path to it
(DNS + TLS + registry auth token, all established lazily on first use). The bring-up was launched **~90
seconds** after Docker Desktop finished booting — the machine had been rebooted this morning and the daemon
was started at iter open — and it put **three concurrent metadata resolutions** onto a cold connection at
once. BuildKit's metadata deadline is per-resolution and does not care that the other two are competing
with it.

**Why the earlier builds in the same run survived:** `cms`, `hiring`, `next-web`, `studio-desk` and
`jobsimulation` ran in earlier waves and completed normally (`demo-1-cms:injected … DONE 2.2s`). By the time
the last wave started, the connection *should* have been warm — so the honest version is that the burst of
three simultaneous cold resolutions is the trigger, not simply "the daemon was cold." Both facts are needed.

**Verdict: environmental and transient. Not a clause-1 red, and not evidence about the advance.** The cycle
is re-run cold from a fresh `down --purge`, and the 3-consecutive-green count restarts at zero — a failed
cycle cannot be counted, skipped, or "resumed."

**But it IS a real finding, and it is routed rather than absorbed.**
`FIX-M257x-iter58-cold-daemon-registry` — a bring-up launched shortly after the Docker daemon starts dies
**~4 minutes in**, at the build phase, with `DeadlineExceeded: context deadline exceeded` and no mention of
a registry anywhere in the message. Nothing in the pre-flight resolves a single base image, so nothing
catches it. **The cheap fix is one warm-up call** (`docker buildx imagetools inspect alpine:latest`) in the
pre-flight, which would both warm the path and fail loudly-and-early with a message that names the cause.

This is the **sibling of `FIX-M257x-iter56-preflight-fails-late`** — there, a correct pre-flight ran *after*
~8 minutes of image builds; here, the check does not exist at all. Same class: **the bring-up's
preconditions are validated late or not at all, so an environmental fault presents as a build failure.**
Worth pairing them under one handler.

## D-M257x-58-4 — my own pre-registration carried a stale container count. Fourth inherited number this iter, and the first one I wrote myself.

**Pre-registration 2 in this iter's `overview.md` reads:** *"Container count is **15**, not 16 — the deleted
WunderGraph router stays deleted."*

**Measured, on all three green cycles: 11.**

```
demo-1-backend-1     demo-1-directus-1   demo-1-fake-bapi-1   demo-1-fake-fapi-1
demo-1-gotenberg-1   demo-1-hiring-app-1 demo-1-next-web-app-1
demo-1-postgresql-1  demo-1-redis-1      demo-1-sentinel-1    demo-1-studio-desk-1
```

`docker ps | grep -ci 'graphql\|router\|wunder\|cosmo'` → **0**. The half of the prediction that mattered —
*the router stays deleted* — is confirmed. The **number attached to it was four containers stale.**

**Where 15 came from, and why it was wrong.** It is iter-14's figure, and iter-14 measured it correctly:
*"15 containers not 16 — exactly the deleted router."* I lifted it from the milestone's `state.md` summary
instead of measuring, and four services have left the default profile since:

| service | why it is no longer a default-profile container | evidence |
|---|---|---|
| `cms` | folded into `app` (cms-in-app v8.0) | no compose service |
| `jobsimulation` | folded into `app` (jobsim-in-app) | no compose service |
| `storage` | **v9.0 fold landed** — `profiles: [storage-legacy]` | `docker-compose.yml`, and iter-57 recorded the same |
| `messenger` | `profiles: [messenger]` only — dropped from `all` | same |

15 − 4 = **11.** It closes exactly, and `backend: profiles: [core, backend, all]` confirms the
`graphql`-profile→`core` rename iter-57 recorded. The measured count is not a surprise; it is the
consolidation this milestone exists to track, arriving on schedule.

**Why this is a decision entry and not an erratum.** This iteration opened by recording two inherited
numbers that did not survive contact with the artefact they described — a container count read through a
dead daemon (`D-M257x-58-1`) and a test count read from before the commit it described (`D-M257x-58-2`).
It then wrote a third one **into its own pre-registration, in the same sitting**, from the same cause:
a number copied from a summary rather than derived at the moment of writing.

That is §5 **rule 20**'s closing observation, reproduced live: *"in two consecutive iterations the author of
a newly-written rule violated it while writing it, which is evidence that a hand-applied discipline does not
survive a corpus of this size."* It now has a third instance, and the strongest possible provenance — the
author had just finished writing the rule down twice.

**The conclusion is not "be more careful."** It is TOK-04 **P4**: this quantity is **derivable** — the
expected container set is a function of the platform's own compose profiles, which `platform_topology.py`
(built in iter-55, extended in iter-56) already parses. autoverify's assert
`container liveness: all 11 expected container(s) running` **already derives it**; my prose did not, and
prose is the only place the error could live.

Routed: `CHECK-M257x-iter58-derive-preregistrations` — pre-registered quantities in an iter's `overview.md`
are prose in the one document that is *supposed* to be refutable. Where the quantity is derivable (container
counts, test counts, subgraph counts), state the **derivation** and let the number be its output. Where it
is not, mark it prose-under-review like any other claim. A pre-registration whose number is stale is worse
than none: it invites a *false* refutation, and a false refutation of a good advance is exactly the outcome
`D-M257x-58-3` had to be careful to avoid an hour earlier.

## D-M257x-58-5 — a "purely additive" app advance moved 22 of 23 `main.go` citations, and the fence caught 1

**This is the iteration's most consequential finding, and it exists only because the advance was taken.**

Phase D's guard re-run went **RED**, unaided, on a corpus file nobody had touched:

```
corpus/services/storage.md:9  [anchor-on-closing-delimiter]
    cites  : app/main.go:983
    which is: }
```

Traced exactly: at `v1.365.0`, `main.go:983` was the *"Storage namespace is the S3 key PREFIX…"* comment
and the cited call `storage.NewClient(…, storagens.CMS)` sat at `:988`. At `v1.366.0` the call is at
**`:992`** and `:983` is a bare `}`. The cause is the one behavioural change in the whole 5-commit range —
the `clerkEventsManager` block moved from ~443 to ~452 so the `user.created` webhook can force-join —
which nets **+4 lines for everything below ~452**.

**So the advance the commit body called "the safest shape an advance can have" — 0 migrations, 0
destructive DDL, 0 removed contract — shifted essentially the entire `app/main.go` citation surface.**
Schema-safety and citation-safety are unrelated properties, and this milestone had only ever measured the
first.

### The measurement (`instrument/measure_mainline_shift.py`, committed per P2)

| | |
|---|---|
| `main.go:N` citations in `corpus/**`, `.claude/**`, `CLAUDE.md` | **23** |
| landing on **different content** at `v1.366.0` than at `v1.365.0` | **22** |
| caught by `anchor_construct_guard` | **1** |

**A catch rate of 1 in 22 — 4.5%.** The fence is not broken: it fires when a cited line is *not a construct
at all*. `storage.md:9` was caught only because its new line happens to be a bare `}`. The other 21 land on
comments and statements that still look like perfectly good anchors — `app/main.go:1196-1202` now opens on
`mux.Handle(skillerv1connect.NewSkillerServiceHandler…` instead of the messenger/cms comment it was chosen
for, in **six** separate files.

This is **`FIX-M257x-iter57-within-block-drift` and `CHECK-M257x-iter57-anchor-guard-bare-class`, both
confirmed and quantified on a real event** rather than reasoned about. iter-57 predicted the blind spot
from structure and could not size it; it is 21 of 22.

### What the number does and does not say — stated because it would otherwise be over-read

The instrument compares line N at two refs. **MOVED is a lower bound on instability, not a count of
newly-false claims.** Corpus citations were written against assorted app refs — `platform-migration-status.md:70`
names `5ba17044` v1.363.2 **in its own prose** — so an unknown share of the 22 was *already* stale before
this advance, falsified by iter-56's 37-commit `v1.363.2 → v1.365.0` step that nobody re-checked the corpus
against. The decisive question — *does each citation resolve to what its prose says it contains, today* —
needs a per-citation expected construct, which only some rows carry. That is the routed work.

### The judgement: the advance stays, and refusing it would have been the error

The tempting inference is *"the advance broke 22 citations, so revert it."* That is backwards. The gate
says **against origin HEAD**, and `v1.366.0` **is** app origin/main — confirmed unchanged at close. Holding
a stale pin to keep the corpus looking true is pinning the platform to the documentation, which is the
inversion this whole milestone exists to correct. The 22 sites were going to be falsified by that advance
whenever anyone took it; **taking it and measuring it is strictly better than not taking it**, and the
alternative — advancing later without this measurement — is how all four prior occurrences of the class
happened.

**Honest clause-5 accounting for this iteration:** 1 site repaired (`storage.md:9`, watched RED → GREEN),
**up to 21 newly falsified**, 0 induced by editing. Under TOK-04 change 1 the net is **negative** — and
that is the metric working exactly as designed: the old metric could not go negative and therefore could
never report that an iteration had cost the corpus more than it gained.

Routed: **`FIX-M257x-iter58-mainline-shift`** — repair the 21 as **ONE derived class** (TOK-04 change 3),
not 21 claims. They share a single mechanical predicate (*a line citation into `app/main.go` below ~452*)
and a single derivable remedy (*re-resolve the cited construct by content, not by offset*). The strategic
form is stronger still and belongs with it: **a line-number citation into a fast-moving file is a claim
with a half-life measured in days.** P4's ordering says derive-else-fence-else-declare; for this class the
derivation exists — cite the **construct** (`func`/symbol/quoted call) and resolve the line at check time —
and it would make the entire class self-maintaining. `measure_mainline_shift.py` is the seed.


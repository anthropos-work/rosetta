---
milestone: M257x
iter: 56
---

# iter-56 — decisions

## D-M257x-56-1 — iter-55's root cause is REFUTED, and the pin advance could never have fixed it

iter-55 closed `user-blocker` on a single claim: *"compose at `0dab54d` deleted `STORAGE_RPC_ADDR`; pinned
`app v1.363.2` still reads it at `main.go:446/516/983`, so `backend` exits 0 in silence. The compose half
of the storage fold landed; the app half is not in the pinned release."* The orchestrator resolved the
blocker by directing a pin advance to `v1.365.0` on exactly that reasoning.

**Both halves of the claim are false, and the first one is checkable in ten seconds:**

| measurement | result |
|---|---|
| `STORAGE_RPC_ADDR` at `v1.365.0` | still read at `main.go:450`, `:520`, `:988`, + `internal/jobsimwiring/wiring.go:115` — the same sites, shifted 4–5 lines |
| `git rev-list --count v1.365.0..origin/main` (app) | **0** — `v1.365.0` **is** app origin/main |
| the v9.0 storage-in-app app half | not released at any ref; the only v9.0 commits in the range are `docs(plan):` design commits |

So there is no app ref, anywhere, in which that env read is gone. *"The app half is not in the pinned
release"* is true but useless — it is not in **any** release, and advancing to the newest one changes
nothing about it.

**How the error was made, and it is a named class.** iter-55 observed (a) the container exits 0, (b) the
env var is absent from the container, (c) the app source reads that env var. Each is individually true and
was individually measured. The *conjunction* was then asserted as causation with no probe joining them —
no error naming storage, no run with the variable restored, no run without it. This is `D-M257x-13`'s
correction verbatim (*"a mechanism that explains the observation is not the same as the mechanism that
produced it"*), and it explained the observation well: it accounted for the exit, the silence, and the
timing, while being false.

**What actually produced it**, measured by running the same image on the same network three ways:

    env as-is, no mounts                     -> starts; 93 log lines; "Web server started at :8082"; >2 min
    env as-is + this host's ~/.aws mount     -> exit 0 in 0 s; 2 log lines; the exact stack signature
    env as-is + a regular EMPTY file         -> starts; still up at 25 s

`~/.aws/credentials` does not exist on this Mac. **Docker does not fail on a missing bind-mount source —
it creates the source as an empty DIRECTORY and mounts that.** (`~/.aws` and `~/.aws/credentials` were both
created at 16:12 today, by the compose run itself.) The AWS config load then dies with
`read /root/.aws/credentials: is a directory`, and the process exits **0**.

**The rule this adds** — `platform-alignment.md` §5, and it is the cheap half of `D-M257x-13`:

> **Three true facts do not make a cause. Join them with one experiment.** When a diagnosis is
> *"X is absent AND the code reads X, therefore X"*, the experiment is to supply X, or to remove the
> other suspect, and it usually costs one command. iter-55 routed a version-pin advance — the single most
> dangerous move in this milestone's history, the one that broke the seeders twice — on a conjunction that
> one `docker run` refutes.

## D-M257x-56-2 — the pin is advanced anyway, for a different and better reason

The advance is kept, because the reason it was ordered was wrong but the destination is right. The gate
says *"against platform @ **origin HEAD**, never a pinned pre-drift commit"*, and `v1.363.2` was three
releases behind `app`'s own origin HEAD. Measuring clause 1 against a stale app would have reproduced
exactly the defect TOK-04 exists to name.

**Recorded per §7 rule 4 — what the advance contains** (`v1.363.2 → v1.365.0`, 37 commits):

| dimension | finding |
|---|---|
| migrations | **2**, both `ALTER TABLE … ADD COLUMN` with a default or NULL — `course_builder_sessions.brief`/`credits_spent`, `academy_chapter_progresses.completed_at` (+ an idempotent backfill) |
| destructive DDL | **0** `DROP TABLE` / `DROP COLUMN` / `RENAME` across the whole range |
| new hard-required config | **0** — not one `log.Fatalf` added in any non-test Go file, so the advance cannot introduce a new missing-env boot failure |
| new env reads | `STRIPE_SECRET_KEY`, `BREVO_KEY` (constructor arguments at `main.go:382`/`:437`, not gated), `WORKFORCE_TEST_DB` (test-only) |
| RPC addresses | unchanged; `STORAGE_RPC_ADDR` still read, still the same three sites |
| feature surface | member-analytics (new `internal/analytics` package + 5 REST endpoints), course-builder credits/rename, academy `completed_at`, wundergraph drop |

That shape is the safest an advance can have: **purely additive schema, no removed contract.** It is the
reason PR-2 predicts the seeders survive rather than hoping they do — the class that broke at v2.1 and v2.7
was a *removed* table or schema, and there is no removal here.

The canonical `demo-stack/clones.pin.json` is advanced with it (`app` → `v1.365.0`, `platform` → `0dab54d`)
so the pin is a **committed file** naming the combination actually proven, per TOK-04 P2 — rather than the
de-facto pin that a checked-out clone had been silently supplying.

## D-M257x-56-3 — the host precondition is derived and fenced, and the fence's SHAPE is forced by the measurement

P4 says derive, else fence, else declare. This is derivable, so it is derived: host-absolute bind-mount
sources belonging to the **default profile's** services, read from the platform's own compose by
`platform_topology.py` — the module iter-55 built for the profile/service/build tuples. Scope comes from a
real property (host-absolute vs workspace-relative), so `./data/postgresql` is excluded because it is
stack-owned and legitimately auto-created, and `storage`'s mount is excluded because `storage` is not in
the default profile at `0dab54d`. Both exclusions follow the platform; neither is a list.

**The non-obvious part, and the reason this is a decision rather than a chore.** The obvious check is
*"does the bind source exist?"* — and it would have reported **GREEN on the exact host state that produced
the defect**, because Docker had already created the path. The residue of an auto-creation is a path that
**exists as an empty directory**. So the check tests for that state specifically, and
`test_this_fence_has_TEETH` runs the existence-only mutant and asserts it *misses* — the fence's own
weaker predecessor is pinned as a negative control, so nobody can quietly simplify it back.

Watched RED on the real host and real clone before it was trusted, then GREEN after the repair. It reports
the path, the reason and the exact remedy, and **never repairs `$HOME` itself.**

## D-M257x-56-4 — a fence that fires late is still a fence, and moving it is not free

The pre-flight sits immediately before `compose up`, which on a cold cycle is **after ~10 minutes of image
builds**. Failing fast would be better and the fix is small. It is deliberately **not** done in this
iteration: the cycles under way are the measurement, and changing runtime source mid-measurement is the
thing TOK-04 P2 exists to forbid. Routed as `FIX-M257x-iter56-preflight-fails-late`.

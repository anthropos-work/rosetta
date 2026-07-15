# iter-05 — Decisions

## D12 — A "cold reset-to-seed run" must be *proven* cold, not *assumed* cold

**The problem.** iter-04 graded the gate on a stack it called "cold reset-to-seed". It was not one. `demo-1`'s
postgres is a **bind mount** (`stacks/demo-1/data/postgresql`), and its `PG_VERSION` still carried the mtime
`initdb` wrote on **2026-07-11** — through every subsequent "cold" bring-up, including iter-04's. Fresh
*containers*, two-day-old *database*.

The cause is **F-9** (below): `--purge` could not delete the DB, so nothing ever reset it.

**The decision.** A cycle counts as a cold reset-to-seed run **only if the tooling can prove `initdb` re-ran** —
i.e. `PG_VERSION`'s mtime is **newer than the start of the cycle**. The battery asserts this as **GATE A**, before
it is allowed to measure anything. A cycle that cannot prove coldness **stops the battery**; it is not retried.

**Why it matters beyond bookkeeping.** The whole point of the 5-cycle gate (per the milestone) is to catch
bring-up **non-determinism** — the F-6 class, where a rebuild silently baked the real Clerk publishable key and
the stack phoned production auth while `autoverify.json` still said `green`. A cycle that reuses the previous
database **and** the previous images cannot exercise that variance at all. It is the one thing the gate was
designed to test, and it was the one thing not being tested.

---

## D13 — Fix the instrument before grading with it (F-9), and accept the count restart

**The defect (F-9).** `rosetta-demo down N --purge` was a bare `rm -rf "$stack/data"`. On a **Linux** host that
cannot work. `up-injected.sh` pre-creates the bind-mount dirs `chmod 777` (M215 F6) so the non-root Bitnami
containers (UID 1001) can write them — but postgres then creates its **own** cluster dir *inside*, as UID 1001,
mode **0700**:

```
drwx------ 19 1001 root  stacks/demo-1/data/postgresql/data     <-- the host user cannot unlink this
```

`rm -rf` exits `Permission denied`. And `rosetta-demo` runs under `set -euo pipefail`, so that non-zero `rm`
**aborted `cmd_down` on the spot**. Every consequence follows from that one line:

| what should have happened | what actually happened |
|---|---|
| the database is wiped | **it was not** — the cluster survived from 2026-07-11 |
| this demo's images are removed (M49 #6, the ENOSPC fix) | **never ran** — 13 images accumulated |
| `reg_del` / `ureg_release` free the registry slot | **never ran** — the slot leaked |
| a failure is reported | **a bare `rc=1`, no diagnosis** |

So `/demo-down --purge` + `/demo-up` silently reused the **same database and the same images**, indefinitely. It
never bit a Mac: Docker Desktop's userspace file sharing lets the host user unlink container-owned files. A
**Linux-host-only** defect — the M215 F1–F12 host-deploy family.

**The decision.** Fix it in rext (**no platform edit**): delete container-owned files **with a container**
(docker runs as root — no `sudo`, nothing host-specific), **verify**, and treat a purge that did not purge as a
**hard, loud failure** — never silently continue, because the next bring-up would *look* cold and *be* warm. A
**G1 path-assert** scopes the root container to this demo's own data dir and nothing else. Fenced by
`demo-stack/tests/test_purge.py` (5 tests; the regression case stages the exact UID-1001/0700 layout and asserts
the host user genuinely cannot remove it *before* purging).

**The cost, accepted.** This is a code change, so **the 5-cycle count restarts at 0** — iter-04's warm-DB number
is *not* counted toward the gate. That is the correct trade: a battery of five runs on an instrument that cannot
produce a cold run would measure nothing five times.

---

## D14 — F-10: a torn-down stack must not be able to hand you a `green`

**Caught in this iter's own instrument.** `autoverify.json` is **not removed on teardown**. When the first
battery cycle's bring-up *failed* (rext pin mismatch — 0 containers running), the file on disk still read:

```json
{"project":"demo-1","offset":10000,"warnings":0,"green":true}
```

— a `green` verdict for a stack that **did not exist**. Both the battery's green gate *and* `run-latency.sh`'s
own green gate read that file. This is the **F-6 class** exactly (the stale `autoverify.json` that let a
Clerkenstein-dewired stack be graded `green`), and it survived M217's hardening.

**The decision.** The battery **deletes `autoverify.json` before every bring-up**, so a green verdict can only
come from the run under test. GATE A (coldness) is what caught this; GATE B alone would have passed it.

**Routed:** the durable fix belongs in the tooling, not in my scratch harness — see **Fate-1 item 2** below.

---

## D15 — The alignment blind spot: Clerkenstein scored 100% while lying about every hero

**The observation that started it.** iter-04 fixed a defect where the fake **BAPI returned the wrong human for
every single request** — a hardcoded stub (`11111111-…`, `demo@anthropos.test`) for *every* hero.
`/align-run` scored Clerkenstein **100% critical / 100% overall, 0 divergences — both before and after the fix.**
Nothing moved. A mirror that serves a stub identity to every hero is **not** a 100% faithful mirror. So the score
was not measuring the thing that was broken.

**Verified — the goldens never assert per-hero identity on the BAPI.** Read across all five DNAs
(`clerkenstein/alignment/dna/`):

| DNA (surface) | identity-related capability | what it actually asserts |
|---|---|---|
| `clerk-2.6.0` (**the Go BAPI** — the surface the fake BAPI implements) | **none** | `VerifyToken` + org/membership/invite/metadata **writes**. **There is no `GetUser` capability at all.** |
| `clerk-express-1` | `ClerkClientBAPI` *(standard)* | BAPI **reads** — but only `get-organization`, `org-membership-list`. **Never the user.** |
| `clerk-express-1` | `ExtractIdentity` | variant **`universal-user`** — asserts *the stub* |
| `clerk-js-5` | `Me` *(standard)* | variant **`universal-user`** — asserts *the stub* |
| `clerk-deploy-1` | `DeployIdentity` | variant **`universal-user`** — asserts *the stub* |
| `clerk-multi-1` | `DistinctIdentity` *(critical)* | per-hero identity — but **FAPI/JWT only** (see below) |

Two independent confirmations:

1. **No runner anywhere calls the BAPI's user endpoint.** `grep` for `/v1/users/` across every runner
   (`clerkrun`, `expressrun`, `jsfapirun`, `multirun`, `deployrun`) → **zero** non-test hits. `getUser`
   (`clerk-backend/server.go:235`) — the exact function iter-04 fixed — is reached by **no gene**.
2. **`DistinctIdentity` — the only per-hero identity gene — exercises the half that was already correct.** Its
   `decodeActive()` drives FAPI `/v1/demo/select` → `/v1/client/sign_ins` → `…/tokens`, then decodes the
   **JWT**. It proves the **FAPI mints distinct per-hero tokens**. It never touches the BAPI.

And the iter-04 fix commit (`8ebc89e`) touched **5 files, zero DNAs, zero goldens**. "Nothing moved" is exactly
right — and exactly meaningless.

**The three genes that *do* mention identity encode the bug as correct.** `ExtractIdentity/universal-user`,
`Me/universal-user`, `DeployIdentity/universal-user` — all three assert the *single universal stub user*. They
stayed green **because** the mirror served the stub. The goldens ratified the defect.

**The corpus already named this hazard — and the named mitigation is what failed.**
`corpus/architecture/alignment_testing.md:169–172`, verbatim:

> **Honesty caveat:** the score is only as complete as the DNA. 100% on a thin DNA is hollow — it
> just means "matches across the genes we bothered to enumerate." Two things keep the DNA honest:
> `/align-dna`'s capability-coverage check (every consumed endpoint is present) and the
> version-bump DNA diff (M1b) that surfaces newly-added source behavior.

`GET /v1/users/{id}` **is** a consumed endpoint — next-web's server-side `currentUser()` consumes it on every
authenticated render. It has no capability. So the safeguard that exists precisely to prevent this **did not
catch it**. That is the finding: not merely "a gene is missing", but "the coverage check that guarantees genes
aren't missing is not actually binding."

### Routing (three-fate rule)

**Fate 1** — land it completely, in this milestone, at the **final harden pass**. It is small, it is a
*regression* golden for a defect we just shipped a fix for, and it is exactly what a harden pass is for. Named
handler: **`/developer-kit:harden-mstone-iters --final`**, before `/developer-kit:close-milestone`.

**Explicitly NOT done inline in iter-05.** Adding a DNA capability is a code change, and a code change
**restarts the 5-cycle count** (D13). Doing it mid-battery would invalidate the battery. It is sequenced *after*
the gate fires, deliberately.

| # | Fate-1 item for the final harden pass | Definition of done |
|---|---|---|
| **1** | **`GetUser` gene on the BAPI surface (`clerk-2.6.0`), with per-hero variants.** At minimum a `two-heroes-differ`-shaped variant: `getUser(hero_A)` and `getUser(hero_B)` must return **A** and **B** — not one stub twice. This gene must **FAIL against the pre-iter-04 mirror** (`8ebc89e^`) and **PASS after** — if it passes on both, it is not fencing the bug. | the gene exists, is `critical`, and is red-before/green-after |
| **2** | **Teardown removes `autoverify.json`** (F-10 / D14) — a torn-down stack cannot hand a stale `green` to the next grader. Fold the battery's workaround into `rosetta-demo`'s `cmd_down`. | `cmd_down` unlinks it; a test fences it |
| **3** | **Make the capability-coverage check binding for the BAPI's consumed endpoints** — or, if that is larger than a harden pass, state honestly in `alignment_testing.md` that the check does not currently cover the fake BAPI's read surface, so the "honesty caveat" mitigation is not over-claimed. | the doc matches reality |

**Note for the handler:** existing goldens for this surface are **source-captured** (`alignctl capture`). If a
real-Clerk capture for `users.Get` is not obtainable cheaply, the *invariant* still holds without one — Clerk's
`users.Get(id)` returning user `id` is the API's defining contract, and `two-heroes-differ` is a **two-sided**
assertion that needs no golden value, only that A ≠ B and each matches the requested id. Do not let a capture
problem become a reason to skip the gene.

**Also add (carry-forward, doc):** `DOC-M218-audit-corrections` already owns `alignment_testing.md` corrections;
item 3 above joins it.

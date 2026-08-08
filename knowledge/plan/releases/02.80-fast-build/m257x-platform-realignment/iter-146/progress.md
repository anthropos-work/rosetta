**Type:** tik — under [`TOK-08`](../decisions.md) (*census the mechanical classes; stop sampling
them*). The class: **tooling that still names the deleted GraphQL router as a reachable endpoint.** A
port either has a listener or it does not, so every instance is decidable without interpretation.

# iter-146 — iter-13's re-point was 97.6 % complete, and the misses were not random

## Step 0 — the re-survey, and the question iter-145 routed rather than answered

iter-145 proved M257x **iter-13**'s re-point off the deleted Cosmo/WunderGraph router was incomplete —
it left the test-side copy of `stack-verify`'s service table, RED for 132 iters. The obvious next
question is not *"what else is broken"* but the narrower, censusable one:

> **iter-13 missed one place. Did it miss others?**

## Phase A — the census, with its denominator

`\b5050\b` over `rosetta-extensions` (excluding `.git`, `node_modules`): **84 references across 31
files.** Classified by what the reference *does*:

| class | n | examples |
|---|---|---|
| guard prose / test fixtures — about the **corpus's** ports, not the platform's | 62 | `anchor_construct_guard.py`, `tests/fixtures/**` |
| a **fence asserting the port's absence** — the repair, working | 8 | `test_frontend_build.py:915`, `test_injection.py:162` |
| a **comment explaining the deletion** | 12 | `readiness.sh:27`, `gen_tailscale_serve.py:39`, `gen_injected_override.py:159`, `up-injected.sh:162` |
| **LATENT** — a build-arg DEFAULT baking a dead endpoint | **1** | `demo-stack/frontend/hiring.Dockerfile:36` |
| **LIVE** — an operator-facing URL for a port with no listener | **1** | `dev-stack/dev-stack:285` |

**82 of 84 correct — 97.6 %.** iter-13's work held everywhere the code runs.

## Phase B — verify each candidate against the MECHANISM, not its wording

**The live one.** `dev-stack:285` printed, on every successful `--public-host` dev bring-up:

```
==> dev-N: reachable from the tailnet — https://<host>:$((5050+n*OFFSET)) (graphql) · …
```

Three facts, each checked rather than assumed: platform `2adcf71` deleted the `graphql` compose
service, so **there is no container**; `stack-verify/lib/services.sh` has no `graphql` row, so
**nothing probes it**; and `stack-injection/gen_tailscale_serve.py:38-41` **deleted its
`tailscale serve` row**, so **nothing fronts it**. The URL resolves to a port with no listener behind
a proxy with no entry.

**And that file says exactly why that is the bad outcome:**

> *"The router's own (`"graphql"`, 5050) row **is deleted**: fronting a port with no listener produced
> a trusted-cert HTTPS endpoint that always refused, **which is worse than no entry at all (it looks
> configured)**."* — `gen_tailscale_serve.py:38-41`

**iter-13 removed the mechanism and left the announcement** (`D-M257x-146-1`), and the announcement is
the line the human acts on. Only the dev path carries it; `up-injected.sh` has no equivalent, so it is
one site.

**The latent one.** `hiring.Dockerfile:36` defaulted `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` to
`http://localhost:5050/graphql`. Checked, not assumed: `up-injected.sh:874` and `:1308` **always** pass
`--build-arg …=$(browser_graphql_endpoint)`, so the default is unreachable on the sanctioned path. It
is what a hand-run `docker build` gets, and Next.js bakes `NEXT_PUBLIC_*` into the **client bundle** —
so it would ship a dead endpoint with no failing signal. Graded **latent, repaired anyway**: the cost
is one line and the failure mode is silent.

## Phase C — repair + fence

**Repairs** (2 files, `rosetta-extensions`):
1. `dev-stack/dev-stack:285` — the `graphql` half dropped; the line now names `backend`'s own port
   **and the path**: *"backend api, and GraphQL at `/graphql/query`"*. **Both moved** (`:5050/graphql`
   → `:8082/graphql/query`), and a host-only correction would 404 rather than refuse — the quiet half.
2. `demo-stack/frontend/hiring.Dockerfile:36` — default re-pointed to `http://localhost:8082/graphql/query`.

**The fence** — `stack-core/tests/test_deleted_router_endpoints.py`, 3 tests. Seven **emitter** files
(anything that bakes a URL into an image, fronts a port, or prints one to an operator) must not name
`:5050` **in executable content**.

**The carve-out is the design decision** (`D-M257x-146-3`): comments are exempt. The four files that
document the deletion best are the four that name the dead port most, so a bare-token predicate would
go RED on exactly the right code. **This is `§5` rule 67 in a second domain** — the same token
carrying opposite obligations depending on whether the line asserts or retracts it. In prose that was
undecidable and the arm stayed SURVEY at 70 %; in **code** the language marks it, so here it is
decidable. Whole-line comments only: a trailing-comment parser would need three languages' quoting
rules, and being wrong permissively is how a fence stops fencing.

**Scope is an allowlist on purpose.** A repo-wide sweep returns 84 hits to grade 2 — iter-143's
measured lesson about widened reach over this material. `test_every_emitter_exists` fails if a listed
path stops existing, so the list cannot pass vacuously.

### The anti-vacuity control — run against the real pre-fix content, not a reconstruction

`§8`: *write the control against the guard's SUBJECT.* Both files were recovered from `HEAD` (their
pre-fix text, not a paraphrase) and run through the predicate:

```
PRE-FIX HITS (the control): 2
   dev-stack/dev-stack:285  echo "==> dev-$n: reachable from the tailnet — https://$pub:$((5050+n*OFFSET)) (graphql) …
   demo-stack/frontend/hiring.Dockerfile:36  ARG NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:5050/graphql
```

Post-fix: **0**. And a synthetic pair proves the carve-out in **both** directions — a documenting
comment is exempt, an emitted URL is not.

## Phase D — the sections the repair touches, re-run

Per rule 68's own instruction, and naming the denominator: this iter touched `dev-stack`,
`demo-stack` and `stack-core`.

## Close — 2026-08-08

**Outcome:** iter-13's re-point censused end-to-end — **84 references, 82 correct (97.6 %)**, and the
two misses are **both emitters on un-exercised paths**: an operator-facing tailnet URL for a port with
no listener, no probe and no front (**LIVE**), and a build-arg default baking a dead endpoint
(**latent**). Both repaired; a comment-carving **emitter fence** added, with its control run against
the real pre-fix text. The finding that generalises: **the repair removed the mechanism and left the
announcement**, one file away from a comment stating that exact hazard. `§5` rule 68 gains sub-rule
**(d)**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–146 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**2 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** `D-M257x-146-1` (grep the EMITTERS of a retired fact, not only its consumers — a
consumer fails loudly, an emitter is just a string and nothing can notice it) · `D-M257x-146-2` (a
repair's completeness tracks EXERCISE, not care — audit the un-exercised paths first, because that is
where the residual is by construction) · `D-M257x-146-3` (the fence carves out comments; rule 67's
shape in a domain where the language makes it decidable).
**Side-deliverables:** none.
**Routes carried forward:**
- **`SURVEY-M257x-iter146-other-retired-services-unaudited` (NEW)** — this iter censused **one**
  retired platform fact end-to-end. The same emitter question is unasked for every other one:
  `skiller`, `skillpath`, `chronos`, `intelligence`, the `storage`/`messenger`/`customerio-sync`
  containers deleted at `838d907`. `D-M257x-146-2` predicts the residual sits on un-exercised paths
  for each. The emitter fence is one allowlist and one regex away from covering them; **do not widen
  it un-audited** — grade a population first, per iter-144.
- `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` (⚠️ **iter-146 did NOT close this** — those ~9 sites are
  *inert mentions* in stub scopes, not emitters, so the new fence deliberately does not reach them).
- `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` (⚠️ **widened again by this iter's sub-rule
  (d) — fourth consecutive iter to grow it and say so**) · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **A consumer of a deleted service FAILS; an emitter of one is just a string.** The re-point held on
   every path that executes and broke on the two that don't. Nothing in a suite, a health check or a
   bring-up can notice a stale URL in an `echo` — which is exactly why that class needs its own fence
   rather than more care.
1. **Read the code that documents the repair — it often states the hazard the repair then commits.**
   `gen_tailscale_serve.py` wrote *"worse than no entry at all — it looks configured"* about fronting a
   dead port, and `dev-stack` advertised that dead port to the operator. The two files were edited in
   the same commit.
2. **When a fence's token appears in both the defect and its documentation, find the axis the language
   already marks.** In prose (rule 67) there is none and the arm stays a survey. In code there is one,
   and it is free.

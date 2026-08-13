**Type:** tik

# iter-37 — the role was never created, and the backend said so at the second it happened

## GATE CLAUSE 2 IS MET

    Playthroughs coverage: 30/31 passing (96.8%)
    passing=30  failing=0  unimplemented=1  unimplementable=0
    209 specs, 209 passed (1.8m)
    provenance: {"binding": true, "scoped": false, "grep_pattern": "", "playwright_exit": 0, "ptreport_exit": 0}

Exactly the figure pre-registered in this iter's `overview.md` before any confirming run existed. Sorted-id
diff against iter-36's own artifact: **one removal (`pt-orgadmin-role-create`), ZERO additions**, and the
failing set is now **empty**. The one `unimplemented` row is the declared in-manifest `will-not-build`
(`onboarding.enterprise-workforce-standard.UC1`); the gate's third figure is ERRORS, and there are none.

**Gate: 3 of 5 → 4 of 5.** Only clause 5 (KB-fidelity) remains.

## What it actually was

Not a front-end hang, and not iter-36's mechanism a second time. The backend logged the cause **at the
second of the click**, and it had been logging it for four releases:

    23:44:42 ERROR graphql resolver error  gql-operation:createJobRole
      user=morgan.reyes2@pt-meridian-labs.com  organization=ad524614-…
      error="createJobRole can't generate skill embedding: can't create embeddings:
             can't get client: azure client EU is not set"

`createJobRole` computes a job-role embedding, and `app/internal/embeddings/embeddings.go:226` asks for
vendor **`skillerai.Azure` hardcoded** — there is no OpenAI fallback in that path. The Azure client is
built only when **both** `SKILLER_AZURE_OPENAI_KEY` and `SKILLER_AZURE_OPENAI_ENDPOINT_URL` are set
(`app/internal/skillerai/ai.go:93`); neither was. So `getClient` returns the error, the mutation is
refused, the dialog closes, the list re-renders, and **nothing is created**.

**The write side settled it before any of that reading happened**, which is what kept the iter off
iter-36's path: `public.job_roles` held no `PT Role%` row and zero rows created in the preceding two
hours. A missing navigation can mean "the app stopped navigating"; a missing ROW cannot.

## The finding that makes this more than a fix: it had already been predicted, by inspection

`stack-secrets/secretdna/secret-dna.json`'s gene for `SKILLER_AZURE_OPENAI_KEY` says, in its own note,
written at **v2.8 M256 iter-21**:

> Gates EVERY taxonomy write: without it `app/internal/skillerai` builds no `azureClientEu` and
> `ComputeJobRoleEmbeddings` (hardcoded `skillerai.Azure`) fails *"azure client EU is not set"*, so
> **createJobRole / custom-skill creation is refused inside an HTTP 200 and renders as NOTHING**. …
> was set on NO demo or dev stack.

Every word of that is correct, and it was arrived at by reading source. **This iter is its live proof**,
and the loop between an inspected prediction and a measured confirmation closes. Also confirmed from the
same note and re-verified here: `SKILLER_OPENAI_KEY` does **not** substitute — embeddings hardcode Azure.

## The fix, and the line it deliberately does not cross

The injected override now gives `backend` a **demo-scoped fallback expression**:

    SKILLER_AZURE_OPENAI_KEY=${SKILLER_AZURE_OPENAI_KEY:-${AZURE_OPENAI_KEY:-}}
    SKILLER_AZURE_OPENAI_ENDPOINT_URL=${SKILLER_AZURE_OPENAI_ENDPOINT_URL:-${AZURE_OPENAI_ENDPOINT_URL:-}}

Three properties, each load-bearing:

1. **The operator's own value ALWAYS wins.** A stack with a real dedicated skiller key is untouched. The
   secret DNA calls these keys **DISTINCT-SIMILAR** from `AZURE_OPENAI_*` and says *"do NOT auto-alias"* —
   and that rule is right, because in **production** they address a different Azure resource and
   conflating two live credentials is how a secret reaches the wrong tenant. **That rule is about the
   SOURCE.** A demo has no second Azure resource and never had one; pointing the skills cluster at the
   stack's own pair is a wiring decision, not a claim that the two production secrets are one. The
   ordering is what encodes that distinction, so it is pinned by a test.
2. **No secret value passes through the tooling.** Compose resolves the expression at parse time: the
   generator never reads the value and the override file on disk never carries it. `/stack-secrets`'s
   values-blind contract survives **by construction**, not by care — mutant N3 makes the generator resolve
   the value itself and the suite kills it.
3. **Empty when neither is set**, so an operator with no Azure at all sees the platform's own unchanged
   error rather than a rext-invented one.

`D-v28-3` forbids a standing red, and this respects it: the DNA's `standard`/`required` classification is
untouched and no key nobody has is now demanded.

## Verification

**Mutants — 4, 4/4 matching declared expectation** (3 RED killed + 1 declared-GREEN no-op control that
survived), each `ast.parse`-gated, control green before and after:

| mutant | declared | actual |
|---|---|---|
| N1 the fallback is never emitted | RED | RED |
| N2 order inverted — `AZURE_*` wins over the operator's own key | RED | RED |
| N3 the generator RESOLVES the value itself (a credential in the override file) | RED | RED |
| N4 **no-op control** — the two emitted lines swap order | GREEN | GREEN |

`stack-injection` **297 → 299, OK** (the two new tests: the expression shape + operator-wins ordering, and
that the pair is **backend-only** — emitting it on cms/jobsimulation would be cargo-cult wiring of exactly
the kind that survives for releases because nothing errors).

**Live, in three steps, each one a control for the next:**

1. **Experiment first, tooling second.** The two keys were copied into `platform/.env` values-blind, the
   backend recreated, and the Playthrough run scoped: **PASS in 7.4 s**, against a 60 s timeout. So the
   stack's existing Azure pair genuinely serves the skills cluster — measured before a line of tooling was
   written, because the whole fix rests on that being true.
2. **Then the durable form.** The generator was re-run with the stack's real arguments and its output
   `diff`s against the RUNNING override as **exactly two added lines and nothing else** — which also
   proves the argument reconstruction was exact.
3. **Then the negative control that matters.** The hand-added `.env` keys were **removed**, and the
   backend **force-recreated** (a plain `up -d` reported `Running` — same resolved config, so no
   recreate; taking that as proof would have measured the *previous* container). With `platform/.env`
   carrying **no** `SKILLER_*` key at all, the container's pair is set. The override is doing the work.

Write side, after: `public.job_roles` holds **2** `PT Role%` rows — one per run — where it held 0.

## Suites

`stack-injection` **OK 299** (baseline 297 + 2). `stack-core` **14F of 396 — baseline** (measured this
run at iter-36 close, unchanged by an injection-only edit). All five corpus guards green.

## Close — 2026-08-02

**Outcome:** **GATE CLAUSE 2 MET** — `29 / 1 / 1` → `30 / 0 / 1` on a binding cold-reset run, zero
additions, failing set empty. Root cause was a missing env pair the secret DNA had predicted by
inspection at M256 iter-21; this is its live proof.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (4 of 5 — clause 5 outstanding)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-37-1 (a demo-scoped fallback is not an alias) · D-M257x-37-2 (resolve in Compose, never in the generator) — see `decisions.md`
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter37-dev-twin-has-no-fallback` — **`dev-stack` does not run `gen_injected_override.py`**
  (verified: `dev-setdress.sh` only *mentions* it, in a comment; `grep -rln` with a positive control). So
  a dev-N stack still has no `SKILLER_AZURE_OPENAI_*` and its taxonomy writes still fail silently. Named
  rather than swept, because the dev seam is a different one and this iter measured only the demo.
- `CHECK-M257x-iter37-other-hardcoded-azure-paths` — `createJobRole` is one caller;
  `GenerateSkillsEmbeddings` / `ComputeSkillEmbeddings` use the same hardcoded `skillerai.Azure`. Which
  other user-visible flows were silently refused is unmeasured.
- `DOC-M257x-iter37-secret-dna-live-proof` — the DNA gene's note should record that its prediction was
  confirmed live (M257x iter-37) and that a demo now falls back. Clause-5 adjacent.
**Lessons:**
- **When a UI action times out, ask the SERVER what it did.** The failure was a 60 s front-end timeout;
  the answer was one ERROR line in `docker logs`, timestamped to the second, naming the operation, the
  user and the org. Four releases of "a flaky role-create" were four releases of nobody reading it.
- **A missing ROW discriminates where a missing navigation does not.** Checking the write side first is
  what kept this off iter-36's (correct, but wrong-here) platform-moved-the-surface path.
- **A rule about a SOURCE is not a rule about a TARGET.** *"DISTINCT-SIMILAR — do NOT auto-alias"* is
  right about production credentials and says nothing about how a demo wires a resource it does not have.
  Reading it as a blanket prohibition would have left the last clause-2 failure standing.

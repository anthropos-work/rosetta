**Type:** tik, under `TOK-08`. Route: `FIX-M257x-263-dev-bringup-must-run-the-check`.

# iter-266 — the gate exists, is fatal, is tested, and the guide never mentions it

## The refutation that relocated the defect

**PR-1 predicted the dev path carries no secret handling at all. It is REFUTED, and that is the iter.**

`dev-stack/dev-stack:241-247` runs the secret-coverage pre-flight **before it builds anything**:

```
"$HERE/../stack-secrets/preflight.sh" --stack "dev-${n:-auto}" || pf_rc=$?
if [ "$pf_rc" = 1 ]; then
  die "secret-coverage pre-flight FAILED (a critical key is missing) — fix the secret source then re-run…"
fi
```

It is the **same wrapper** the demo path calls (`demo-stack/up-injected.sh:1511`), it is **fatal** on a
critical miss, it has an opt-out (`DEV_NO_SECRET_PREFLIGHT=1`), and its wiring is **asserted by the
tooling's own tests** (`dev-stack/tests/test_dev_stack.py:279-284`), which were green throughout.

The survey grep missed it because it searched for the **binary's** name (`stacksecrets`) and the dev path
invokes the **wrapper** (`preflight.sh`) — 0 hits across the whole of `dev-stack/`. That is §10 iter-190
exactly: *a selector that finds only the case it was written from is not an enumeration.* Recorded because
the pre-registration is the only reason it was caught before the repair was written.

## So where is the defect?

**In this corpus, not in the tooling.** `corpus/ops/setup_guide.md` is the by-hand path, it mentions
`/stack-secrets` **8 times — every one of them `--provision`, never the coverage `check`** — and at
`:181-196` it presents `/dev-up` as *"automate this entire setup process"*, i.e. as an **equivalent
convenience**. It is not equivalent: the automated path carries a fatal gate the manual path does not.

**That gate is the one that would have caught iter-262.** Re-measured on today's tree (PR-4, held):

```
▶ secret-coverage pre-flight (stack=dev-0, source=…/.agentspace/secrets)
  ⚠ platform  13/29 short: … INVITATION_HMAC_SECRET …
  ✗ secret-coverage: a CRITICAL secret key is missing — the stack would be broken without it.
rc=1
```

`/dev-up` would have refused to build and named the key. iter-262 followed this guide by hand, got no
check, and lost the bring-up to `backend` **exiting with code 0** — a refusal that returns success, so
`docker ps` shows an absence rather than a crash and no log line says *secret*.

**This is iter-265's class one layer up, and the layer matters.** iter-265: a requirement migrated and its
documentation stayed behind. Here: a *capability* exists on the path that gets exercised, and the document
that hands a new engineer the *other* path never learned it needs to say so. Neither is a tooling defect;
both are limb 3.

## Repair

`corpus/ops/setup_guide.md`, two sites, both on the by-hand path:

1. **§ Automated Setup with Claude Code** — a warning that the paths are **not equivalent**, naming the
   `dev-stack:244`/`:246` gate, its opt-out, and the exit-0 failure it prevents.
2. **New § "Verify secret coverage before you build — the gate `/dev-up` enforces and this guide does
   not"** — the runnable command (the same wrapper both stacks call), the **rc 0/1/2** contract, how to
   read the verdict **by class not by count** (a `13/29 short` warn line is not the failure; the `✗` line
   is), the iter-262 story as the motive, and the fix discipline: **add the key to the secret source and
   re-provision, never hand-edit `.env`**, which is how the gap survives the next re-provision.

**No skill file changed** — `stack-secrets/SKILL.md:22` and `dev-up/SKILL.md:137-142` both describe the
pre-flight accurately (PR-2 refuted: the `/dev-up` skill *does* claim it, and truthfully).

**Side-repair, forced by the edit:** the +44 lines shifted `setup_guide.md` under
`platform-alignment.md:1054`'s pin `setup_guide.md:514`, and `anchor_construct_guard` went RED. Re-derived
rather than re-pointed (§8, iter-22): the subject — the `migrations: true` enumeration — now lives at
**`:598`**, and the pin was stale in *substance* before it was stale in *position*, since iter-77 had
already rewritten the sentence it quoted.

## Pre-registration grading

| PR | prediction | outcome |
|---|---|---|
| **PR-1** | `dev-stack/` carries no secret handling at all (0 matches) | **REFUTED — and it is the iter's finding.** 0 matches for `stacksecrets`; the wiring is `preflight.sh`, fatal, at `dev-stack:244` |
| **PR-2** | `/dev-up`'s SKILL never claims a secrets pre-flight (one-sided disagreement) | **REFUTED** — `dev-up/SKILL.md:137-142` documents it fully and correctly. There was no disagreement to repair |
| **PR-3** | ≤ 3 sites assert the check rides inside `/dev-up` | **HELD** — 3 (`stack-secrets/SKILL.md:22`, `:159`, `dev-up/SKILL.md:137`), **all true** |
| **PR-4** | the dev-side critical miss still reproduces | **HELD** — `preflight.sh --stack dev-0` → **rc 1**, `INVITATION_HMAC_SECRET` among the 13/29 short |
| **PR-5** | doc half lands, tooling half routes behind a pin bump | **REFUTED, favourably** — there is **no tooling half**. The whole repair is documentation, and `D-M257x-258-1`'s frozen-pin control is **not** spent |

**Three of five refuted, and the iter is better for it.** PR-1/PR-2/PR-5 all encoded the same wrong prior —
*the tooling is behind the docs* — inherited from iter-262/263, where it was true. Here it is inverted:
the tooling is ahead, tested, and correct, and the corpus is what lags.

## Close — 2026-08-10

**Outcome:** The route is closed against a **refutation**: the dev bring-up already runs the check and
`die`s on a critical miss. The real defect is that `setup_guide.md` — the by-hand path, and the one
iter-262 followed — never mentions the gate while presenting `/dev-up` as an equivalent convenience. Both
sites repaired; the shifted anchor re-derived. All six corpus fences green.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-266-1` (the gate was missing from the document, not the tooling).

**Side-deliverables:** the `platform-alignment.md:1054` anchor re-derivation (`:514` → `:598`) — forced by
this iter's insert, committed with it, and re-derived rather than re-pointed.

**Routes carried forward:**
- `FIX-M257x-263-dev-bringup-must-run-the-check` → **CLOSED** by refutation + the guide repair.
- `FIX-M257x-266-manual-path-drops-gates-the-automated-path-enforces` — **new.** This iter found **one**
  instance. The general assertion — *the by-hand guide warns about every fatal gate `dev-stack up`
  enforces* — needs its population derived before it can be fenced (§8, iter-173), and iter-265's own
  lesson forbids scoping a fence from a single known instance.
- `FIX-M257x-262-dev-path-needs-the-studio-acquisition` (tooling half),
  `ROUTE-M257x-261-succession-projection-is-empty`, `FIX-M257x-262-demo-env-append-is-not-idempotent`,
  `FIX-M257x-265-prose-deletion-instructions-are-out-of-D-reach`,
  `ROUTE-M257x-265-stack-demo-carries-six-dead-clones`, `ROUTE-M257x-258-the-pin-is-157-iters-stale` → open.

**Lessons:**
1. **Grep the WRAPPER, not the binary.** `dev-stack/` has 0 occurrences of `stacksecrets` and calls it on
   every `up`. An absence measured through one vocabulary is a fact about the vocabulary.
2. **"Automates this entire process" is a claim about EQUIVALENCE, and it is usually false.** Wherever a
   corpus offers an automated path beside a manual one, the automated path has accumulated gates. The
   manual reader inherits the risk and none of the checks.
3. **A prior inherited from two consecutive iters is still a prior.** *The tooling lags the docs* was true
   at iter-262 and iter-263 and shaped three of this iter's five pre-registrations. It was false here.

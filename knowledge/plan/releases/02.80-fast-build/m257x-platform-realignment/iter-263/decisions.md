# iter-263 — decisions

## `D-M257x-263-1` — CORRECTS `D-M257x-262-3`: the tooling already declared it; the gap was the SOURCE and an unrun check

**`D-M257x-262-3` (iter-262, `56bcebc`) is corrected. Two of its three claims were wrong when written.**

| iter-262 claimed | measured at iter-263 |
|---|---|
| an **undeclared** boot requirement | **declared** — `platform/INVITATION_HMAC_SECRET` is a gene, **critical + required**, target `platform/.env`, scope shared, operators key-present + non-empty, with `secret_dna_json_test.go:141-161` pinning every field |
| a newly-named `exit 0` class | **already named verbatim** — `secretdna/demo.go:45-47`: *"`invitations.NewTokenManager` ERRORS when it's empty and main returns (**the silent `app Exited (0)` class**)"* |
| add it to the DNA + the docs | **already in both** — `corpus/ops/secrets-spec.md:105`, `:118` (*"critical/required — the `app` exits early when it is unset"*), `:273` (the demo-auto-generated family) |
| a dev `.env` from `.env_example` + this source cannot start `backend` | **SURVIVES** — this half was right |

**The real cause, measured.** `stacksecrets check --from .agentspace/secrets`, same DNA and same source,
one flag apart:

| target | `platform` | `INVITATION_HMAC_SECRET` | **Critical** |
|---|---|---|---|
| **dev** | 13/29 | **SHORT** | **92.3 %** → `check` exits **1** |
| **demo** (`--demo`) | 16/29 | satisfied without the source (`IsDemoGenerated`) | **100.0 %** → exits **0** |

So: **(a)** the secret source is genuinely missing a gene the DNA calls critical for dev, and **(b)**
iter-262 never ran the verb that says so — it hand-provisioned `.env` from `.env_example` plus an overlay
instead of driving `/stack-secrets`. **An operator error plus a source gap, not a documentation or tooling
gap.** The distinction decides what gets fixed: nothing needs declaring, the source needs a value, and the
**documented dev bring-up needs to actually run the check**.

**This also explains the demo asymmetry without appeal to luck.** Critical coverage is **100 % for a demo
and 92.3 % for a dev stack on the identical source**, because the demo auto-generates this key at provision.
Three green demo cycles could not have surfaced it.

**Unchanged by this correction:** `D-M257x-262-2` (the `app/studio` acquisition gap) stands in full, and it
is the more serious of the two — **no instrument would have caught it**, whereas this one had a check
sitting unrun.

**Routed:** `FIX-M257x-263-dev-bringup-must-run-the-check` — (a) add the missing critical gene to the secret
source; (b) wire `stacksecrets check` into the documented dev bring-up, because a check nobody runs is a
check that does not exist.

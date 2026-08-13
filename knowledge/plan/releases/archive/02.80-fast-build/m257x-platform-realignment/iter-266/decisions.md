# iter-266 — decisions

## Pre-registrations — SEALED BEFORE THE CONFIRMING MEASUREMENTS

Sealed in this iter's first commit, corpus at `79d63d1`. Known at seal time and NOT predicted: a grep of
`dev-stack/*.sh` + `dev-stack/dev-stack` for `stacksecrets` returned nothing. Everything below extends past
that.

**PR-1 — the dev path carries no secret handling AT ALL.**
`grep -ri 'stacksecrets\|secretdna' rosetta-extensions/dev-stack/` over **every** file (not just the two
greppped at survey) returns **0** matches. *Risk:* real — the check could ride via a shared helper in
`stack-core/`, via `dev-setdress.sh`, or be invoked by the skill's prose rather than the script.

**PR-2 — `/dev-up`'s own SKILL.md never claims a secrets pre-flight.**
It names `/stack-secrets` only as a manual provisioning step, so the disagreement is **one-sided**: the
`stack-secrets` skill asserts a rider that the `dev-up` skill does not claim to carry. *Risk:* if `/dev-up`
claims it too, this is a two-sided false claim and a wider repair.

**PR-3 — the assertion has few homes.**
A sweep of `.claude/skills/**` + `corpus/**` for the claim that a secrets check runs inside `/dev-up`
finds **≤ 3** sites. *Risk:* falsifiable both ways; a large population would make this iter-265's class
again rather than a single false claim.

**PR-4 — iter-263's asymmetry still reproduces on today's tree.**
`stacksecrets check` against the same source, one flag apart, still yields **dev exit 1 / demo exit 0**
(iter-263 measured `platform` 13/29 critical 92.3 % vs 16/29 critical 100 %). *Risk:* iter-262
hand-provisioned `dev`'s `.env` afterwards, so the dev side may now pass — which would change the repair
from *"the check would have caught it"* to *"the check no longer demonstrates it"*.

**PR-5 — the tooling half needs a pin bump and the doc half does not.**
Making `dev-stack` actually run the check requires a rext tag + the stack's pinned clone to move, which
would spend `D-M257x-258-1`'s frozen-pin control; correcting the claim + documenting the step does not.
Prediction: the iter lands the doc/skill half and routes the tooling half, mirroring iter-264's split of
`FIX-M257x-262`. *Risk:* the dev path may consume rext from a path that needs no tag at all, in which case
the whole fix lands here and the prediction is refuted.

## D-M257x-266-1 — the gate was missing from the DOCUMENT, not from the tooling

`dev-stack/dev-stack:241-247` runs `stack-secrets/preflight.sh` before any build and **`die`s** when a
critical key is short. Same wrapper as the demo path (`up-injected.sh:1511`), opt-out
`DEV_NO_SECRET_PREFLIGHT=1`, and the wiring is asserted by `dev-stack/tests/test_dev_stack.py:279-284`.
Re-measured today: `preflight.sh --stack dev-0` → **rc 1**, with `INVITATION_HMAC_SECRET` among
`platform 13/29 short`. So `/dev-up` would have refused iter-262's bring-up and named the key.

`corpus/ops/setup_guide.md` — the by-hand path iter-262 actually followed — mentions `/stack-secrets`
**eight times, all `--provision`, never the coverage `check`**, and at `:181-196` frames `/dev-up` as
*"automate this entire setup process"*. The framing is the defect: it asserts **equivalence** between a
path that carries a fatal gate and one that does not.

**Why this is a sharper defect than a missing feature.** Nothing is broken, nothing is untested, and no
instrument is wrong. The tooling's tests pin the gate; the skills document it accurately (all three
assertions checked, all true). The single artifact that omits it is the one handed to a new engineer, and
its omission is invisible from every other vantage — which is why three consecutive iters (262 hand-built,
263 relocated the cause, 266 opened predicting a tooling gap) each carried the prior forward.

Repaired at both sites: the equivalence claim now carries the disclaimer, and a new section gives the
runnable check, its rc contract, how to read the verdict by class rather than by count, and the discipline
that a critical miss is fixed in the **secret source** and re-provisioned — never by hand-editing `.env`,
which is precisely how the gap survives the next `--provision`.

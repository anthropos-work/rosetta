# iter-284 — decisions

## D-M257x-284-1 — count the POINTERS, not the guards

`STORAGE_S3_BUCKET` reaches a demo through **two** independent channels: the seeder's environment
(`PreflightEnv`) and the running container's compose environment. The existing guard covered the one the
tooling owns; the write the user hit came through the one the product reads. **A fence over a proper
subset of a variable's carriers is indistinguishable, from the outside, from a fence over all of them.**
Both are now covered, and the doc says which mechanism owns which.

## D-M257x-284-2 — empty is the platform's own fallback, and the cost is accepted and STATED

`app/internal/storage/storage.go` builds a tmp root when the bucket is `""`, and `main.go`'s fatal on an
empty bucket is gated on `deployedEnvironment()`. So `""` is not a sentinel this tooling invented — any
other value (a placeholder, a `demo-N`-derived name) would be a bucket the app genuinely tries to reach.
**The cost:** uploads land on ephemeral container disk and `GetPresignedUrl` returns `("", nil)`, so a
freshly uploaded asset renders as a broken image. Accepted: **a broken thumbnail beats a write to the
production bucket**, and it is written down rather than left to be discovered by the next demo-giver.

## D-M257x-284-3 — re-class `s3-private`; iter-98's refusal is spent

iter-98 declined to re-class the store because *"fixing the code is only available when the change is
yours to make"*, and recorded the doc-vs-registry disagreement openly instead. That was right at the time.
The user has now ruled this the top-priority item, so the change is ours to make. The registry's own note
— *"falls back to local /tmp on demo"* — described behaviour that is real **only when the bucket is
empty**, which compose guaranteed it was not. **A store whose default pointer is production was never
per-stack-isolated by any reading.**

## D-M257x-284-4 — fence the retraction's ACCEPT side, not only its FIRE side

Forcing the bucket without deleting the sentences that say it is unforced leaves the document asserting
both. The doc-drift test therefore pins the three stale phrasings as **FORBIDDEN** alongside the positive
assertion that the doc names the key and the code forces it. Without the accept side, the old wording
returns on the next paste and the fence stays green while the document is wrong again.

## D-M257x-284-5 — `demo-2` is NOT restarted, and that is a disclosure

The override is read at bring-up, so the live `demo-2` still carries the production bucket. The user is
validating on it right now; pulling its `backend` container out from under them to apply a fix whose
absence is currently masked by an IAM 403 is the wrong trade. **Recorded as an open route, not as
"fixed".** The next `/demo-up` carries the containment.

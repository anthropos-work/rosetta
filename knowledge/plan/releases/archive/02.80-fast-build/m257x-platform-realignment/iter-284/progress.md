# iter-284 — a demo pointed at the production S3 bucket, and only an IAM policy stopped it

**Type:** tik — under `TOK-09`.

## What was already true, and what the corpus already said

`corpus/ops/safety.md` **did not invent a guarantee it lacked** — its S3-private row named this exposure
precisely: *"unclassified — the guard does not cover it… a stack with working AWS credentials writes
private uploads into the production private bucket,"* filed as `DEF-M257x-iter80-storage-prod-bucket`,
severity high, **open for ten iters**. Every restatement of it was accurate.

**And it happened anyway.** That is the finding worth more than the fix: a hazard can be correctly
recorded, correctly re-derived at a newer ref, correctly flagged CHRONIC_DEFER — and none of that makes it
cheaper to hit or easier to see. It took a user clicking a button.

## The measurement: TWO pointers, and the guard covered the wrong one

| pointer | what read it | covered before |
|---|---|---|
| the **seeder's** env (`PreflightEnv`) | `stackseed` | **public bucket only** — `STORAGE_S3_BUCKET` was not in the forced set |
| the **running container's** env (compose) | `backend`, which makes the actual `PutObject` | **nothing** |

The write the user hit came through the second one. `backend` reads its own compose env and never sees
the seeder's, so a fence over `PreflightEnv` alone reads exactly like a fence over both — **and did**.
Platform compose **hardcodes** both buckets (`docker-compose.yml:81-82`), so there is no `.env` seam and
no platform-side fix available under v2.8's zero-platform-edit constraint.

## The fix, on both pointers, plus the classification that was false

1. **The injected compose override** (`stack-injection/gen_injected_override.py`) now strips
   `STORAGE_S3_BUCKET` and `STORAGE_S3_PUBLIC_BUCKET` to empty on **every emitted service**, beside the
   `DIRECTUS_TOKEN` strip that has been there since fix16/17 — same mechanism, same rationale, one class
   further. Scoped to the per-service emit path on purpose: the UI tier has its own env block and, unlike
   a credential, **a bucket NAME is inert in a container that never reads it.**
2. **`PreflightEnv`** forces both buckets on every target.
3. **The store registry was re-classed** — `s3-private: PerStackIsolated → SharedPollutionRisk`. Its old
   note said it *"falls back to local /tmp on demo"*, and that is only true when the bucket is **empty**;
   compose hardcoded a production bucket, so the fallback never happened. **A store whose default pointer
   is production was never per-stack-isolated by any reading.** iter-98 declined to re-class it — *"fixing
   the code is only available when the change is yours to make"* — which was right then and is now spent.

**EMPTY is `app`'s own local-fallback value, not a sentinel we invented:** `internal/storage/storage.go`'s
manager constructors build a tmp root when the bucket is `""` and `getKeyPath` returns a local path rather
than an `s3://` URL; `main.go`'s fatal on an empty bucket is gated on `deployedEnvironment()`, which a demo
is not. **The cost is stated rather than discovered later:** an upload then lands on the container's
ephemeral disk and its presigned URL comes back empty, so a freshly uploaded asset renders as a broken
image. That is the platform's own documented W2 behaviour, and for a demo it is the right trade.

## The document, corrected in both directions

`safety.md` asserted *"a demo cannot write prod"* **unqualified at four sites** while its own S3-private
row documented the hole — a document that contradicts itself, where a reader finds whichever comes first.
All four now carry *"over the covered pointers"*, and the PM-facing paragraph carries the caveat in full:

> **every claim of the form "a demo cannot write prod" is a claim about the set of pointers this tooling
> knows to override.** A pointer the platform adds tomorrow is outside it until someone notices, and
> *"nothing was written"* is the outcome, never the guarantee.

**Fenced in both directions**, because a correction that is not fenced is a correction that lasts until
the next paste. `stack-seeding/isolation/safety_doc_drift_test.go` gains the FIRE side (the doc must name
`STORAGE_S3_BUCKET` and `PreflightEnv` must actually force it, on prod and non-prod targets) **and the
ACCEPT side** — the three stale phrasings are pinned as FORBIDDEN, so the retraction cannot be silently
un-done. **Proven RED** on the pre-repair document before the repair landed.

## Verification

| scope | result |
|---|---|
| `stack-seeding/isolation` (Go) | **ok**, `-count=1` |
| `stack-injection` whole section | **337 passed** |
| the new bucket-strip arms, with the strip removed | **2 failed** — RED-proven, then restored |
| the doc-drift ACCEPT arm, against the pre-repair document | **RED**, naming all three stale phrasings |
| `prose_twin_guard` (this iter's own new prose) | **OK — 0 RED** |

**NOT COVERED, stated:** no stack was brought up, re-seeded, restarted or torn down. **`demo-2` is live
and the user is validating on it — its `backend` container still carries the production bucket**, because
the override is read at bring-up. The containment lands on the **next** `/demo-up`; until then the
observed behaviour on `demo-2` is unchanged (the IAM 403), and that is a disclosure, not a mitigation.

## Close — 2026-08-11

**Outcome:** the demo's production-S3 write path is closed at **both** pointers, the store is re-classed
to the class it always belonged to, and `safety.md`'s four unqualified *"cannot write prod"* claims are
qualified and fenced in both directions. `DEF-M257x-iter80-storage-prod-bucket` — **open for ten iters** —
is discharged.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**

**Decisions:** `D-M257x-284-1` (two pointers, and the guard covered the wrong one) · `D-M257x-284-2`
(empty is the platform's own fallback, and the broken-thumbnail cost is accepted and stated) ·
`D-M257x-284-3` (re-class `s3-private`; iter-98's refusal is spent) · `D-M257x-284-4` (fence the
retraction's ACCEPT side, not only its FIRE side).

**Routes carried forward:**
- **`ROUTE-M257x-284-demo-2-is-live-and-uncontained`** — the running stack predates the fix. A bring-up
  is the user's call, not this iter's.

**Lessons:**
1. **A correctly-recorded open hazard is still an open hazard.** Ten accurate restatements did not make
   it cheaper to hit. A ledger that keeps its entries honest is not the same as one that closes them.
2. **Count the POINTERS, not the guards.** One env variable lived in two places; the guard covered the
   one the tooling owned rather than the one the product reads.
3. **A doc that states a hole AND denies it elsewhere is a doc that says whatever the reader greps
   first** — so the retraction needs an ACCEPT-side fence, or the old wording comes back by paste.

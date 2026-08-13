---
iter: 284
milestone: M257x
iteration_type: tik
status: closed
opened: 2026-08-11
---

# iter-284 — a demo pointed at the production S3 bucket, and only an IAM policy stopped it

**Active strategy:** `TOK-09` — the closed defect list, safety item first.

**Cluster / target:** the user's defect 1. Studio-desk's *"fine-tune in Advanced mode"* on `demo-2`
attempted `s3:PutObject` against `s3://production-storage…/cms/<uuid>` and was refused **403** by IAM.
Nothing was written. **The refusal came from an account policy we do not control.**

**Hypothesis:** the demo's storage pointer is the production bucket, and `corpus/ops/safety.md`'s
*"a demo cannot write prod"* is stronger than the design supports.

**Expected lift:** the pointer overridden at every place a demo carries one; the document corrected
wherever it overstates; both fenced.

**Escalation:** if the only real fix were platform-side, say so and land demo-side containment plus the
honest correction rather than editing the platform (v2.8 holds: 0 platform edits).

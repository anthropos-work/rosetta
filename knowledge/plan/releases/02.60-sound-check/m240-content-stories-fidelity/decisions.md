# M240 — Decisions

_(implementation choices with rationale accumulate here during build)_

- **User decision (2026-07-20): media → PORT IT.** Capture + re-host the Chime/S3 voice recording + document blobs. Gated by a HARD internal PII gate (fresh data-controller sign-off + a `safety.md` §3.8 raw-media amendment + a voice/document anonymization decision) that MUST clear before any customer audio lands in a demo — a voice cannot be token-scrubbed. Likely consumes DEF-M10-01. Recorded at design time; carried here for build traceability.

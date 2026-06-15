# M31 — Decisions

_Implementation decisions with rationale, numbered `M31-D1`, `M31-D2`, … . Empty at scaffold; filled during build._

_Pre-decided at design (2026-06-15, see `.agentspace/scratch/roadmap-research-2026-06-15.md`):_
- _Fallback is openssl self-signed (NOT fail-loud) — the never-abort-a-good-bring-up contract._
- _BAPI is out of scope (plain HTTP, no browser TLS handshake)._
- _SANs = `127.0.0.1 localhost ::1`; cert CN is a non-issue (the pk validator checks SANs, not CN)._
- _`DEMO_NO_MKCERT=1` opt-out exists (dev-CA-in-trust-store is a real trust expansion)._

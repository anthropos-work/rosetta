# M15 — Decisions

_Implementation decisions with rationale. ID scheme: M15-D1, M15-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M15-D1 | Doc home = `corpus/ops/safety.md` (resolves Q1). | Ops-adjacent to seeding-spec / snapshot-spec / db-access (the three docs it cross-links and consolidates from); a reader operating a stack looks under `corpus/ops/`, not `corpus/architecture/`. | 2026-06-07 |
| M15-D2 | safety.md names the REAL Go functions (`AssertPlan` + `AssertCaptured`) alongside the conceptual umbrella name `AssertPublicOnly`. | KB-fidelity finding 1: `AssertPublicOnly` is a conceptual name (the firewall "invoked twice, defense in depth"), not a Go symbol. The existing docs use the umbrella name; safety.md keeps that convention BUT names the two grep-able functions so a reader can find them. Accuracy over prose. | 2026-06-07 |
| M15-D3 | safety.md must NOT claim an offline pg_dump-FILE reader. | KB-fidelity finding 3: the direct file reader was DROPPED in M9b-D9. A dump is ingested by RESTORING into Postgres and read over a DSN. The capture is always a DSN read (dump-ingest default → throttled primary-read fallback, both `SET TRANSACTION READ ONLY`). | 2026-06-07 |
| M15-D4 | safety.md describes the n=0-dev guard PRECISELY: it gates (a) unsolicited auto-set-dressing (`dev-setdress.sh`) and (b) destructive `--reset` (`stackseed`) — NOT all tools. `stacksnap` replay has no N=0 guard (correctly — replay writes only public data into the stack's own isolated stores). | KB-fidelity finding 5: the `dev-setdress.sh` source comment over-stated "stacksnap/stackseed independently refuse N=0 too" — stacksnap does not. The "doubled in M13" framing = two independent enforcement points (auto-set-dress + reset), which is accurate. **Resolution (Fate-1, landed in harden — ext `c8a0589`):** the over-claim was corrected in BOTH the `dev-setdress.sh:19-22` source comment AND the sibling `provision-plan/main_test.go:179-181` comment so both match the shipped safety.md §2.5 — a stale comment contradicting the safety doc is exactly what harden catches. (Originally noted as "flag, don't fix" during build to avoid scope creep into M13 code; harden re-evaluated and landed it completely rather than leave a doc↔code contradiction.) | 2026-06-07 |

## Open at design (RESOLVED)
- M15-Q1 → **resolved (M15-D1):** doc home is `corpus/ops/safety.md`.

## Release-scope deferral — formalized at v1.3 close-release (2026-06-07)

`RELEASE-SCOPE-DEFER: DEF-M10-01 → v1.4` — **Directus S3 media blob *bytes* + the cloud `SnapshotStore` backend.** Recorded here at v1.3 close-release because the prior records carried this item as Fate-2 (originally → v1.3) without ever formalizing the escape-hatch when its destination moved forward to v1.4. Sign-off: **user, 2026-06-07** (at v1.3 close-release; the move itself was the user's 2026-06-07 v1.3-scoping decision).

- **Why Fate-1 (land in v1.3) fails — concrete:** capturing the actual media blob bytes + a cloud (S3) snapshot store both require **eu-west-1 S3-read credentials the project does not have** — verified during the M9a capture-source investigation (`~/.aws/credentials` is 0 bytes; no read replica; no local AWS creds). The code (`media.go BlobBytesAvailable`, `provision.go` per-stack Directus) is built + unit-tested but unreachable from any entrypoint, gated on that access. No amount of v1.3 work unlocks the credential.
- **Why Fate-2 / Fate-3 (route to a v1.3 sibling) fail:** v1.3 "stack party" was deliberately scoped to the dev/demo *convergence* (registry, dev-peers, unified skills, safety doc) — the cloud-store / S3-blobs / AI-content / shareability seeds were **moved out to v1.4** by the user on 2026-06-07. No v1.3 milestone (M12–M15) owns media bytes; the only valid home is v1.4. v1.3 replays media **refs-only** (placeholder bytes) by design, which `safety.md` documents with a labelled `Future (v1.4)` pointer.
- **Repeat / `DRIFT_DEFER` acknowledgment (the Phase 1b finding):** this item is a **`DRIFT_DEFER`** — its destination moved forward (v1.2-close: Fate-2 → v1.3; v1.3-design 2026-06-07: → v1.4), making this its **2nd release carried**. It is **not** `CHRONIC_DEFER`: the reason is consistent and real (a hard credential gap), not "no time", and the forward move was a single, **dated, user-authored** re-scoping decision (2026-06-07), not a silent or repeated work-dodge. The four per-milestone audits' "no repeat" wording was mechanically imprecise (the skill counts destination-updated-forward as a repeat); this record corrects that to the accurate `DRIFT_DEFER` classification.
- **Forward tracking verified:** `roadmap-vision.md § v1.4 seeds` carries the entry (cloud snapshot store + S3 media blob bytes), backref'd to DEF-M10-01.
- **First deferred:** **2026-06-06** (M10 close, v1.2) — not 2026-06-07. (Corrects the `first_deferred_on` in `audit-deferrals/deferral-audit-2026-06-07-m15-close.md`.)

## KB-fidelity (Phase 0b) verdict
- **GREEN** (2026-06-07). Report: `kb-fidelity-audit.md`. Every read-side/write-side safety claim ALIGNED with the actual extensions code. Two accuracy guardrails carried into the build → M15-D3 (no offline file reader) + M15-D4 (precise n=0 scope).

## Adversarial review (close Phase 2c)
The milestone's value is a *claim surface* (`safety.md`) plus drift guards that pin it to code. The adversarial
question for a docs+guard milestone is the **false-negative guard**: could a drift guard pass while the doc is
actually wrong? If so, the guard is theatre and the safety contract can silently rot.

- **Scenario — guard false-negative (read-side).** Mutated `safety.md`'s taxonomy public predicate
  `organization_id IS NULL` → `... IS NOT NULL` (the single most load-bearing read-side literal). `TestSafetyDocPredicatesMatchCode`
  **FAILED** as designed (`doc<->code drift`). Restored byte-identical; guard green again.
- **Scenario — guard false-negative (write-side).** Renamed every `STORAGE_S3_PUBLIC_BUCKET` occurrence in
  `safety.md` (the forced-empty bucket override — the highest-risk prod-S3 vector). `TestSafetyDocForcedBucketOverride`
  **FAILED** as designed. Restored byte-identical; guard green again.
- **Scenario — consumption-clone breakage.** Both guard files resolve `safety.md` via a relative path up to the
  rosetta repo root and **`t.Skip` gracefully** when unreachable (per-stack pinned-tag clone / CI). So a pinned-tag
  consumption clone that lacks the sibling corpus does not red — the guard fires only in the authoring copy where
  `safety.md` is actually edited. Correct: it guards drift at the one place drift can be introduced, without coupling
  every stack's clone to the doc corpus.

Conclusion: the guards are genuinely fail-closed, not theatre. The risk that just shipped (doc rotting away from
code) is mechanically caught. No fix needed; scenarios recorded per Phase 2c.

# M254 · iter-03 — decisions

## D1 — demopatch re-point authored against app@3df8536 (manifest-only fix, self-healing gate)
The apply helper (`stack-injection/apply-app-aireadiness-loadmembers.sh`) reads `path` from the manifest and
delegates to `apply_patch.py` (self-healing: anchor 1× is the contract, whole-file sha is a baseline). So the
fix is manifest-only. But because the consolidation MOVED the file (whole `internal/workforce/ai_readiness.go`
→ `internal/aireadiness/readiness.go`), the self-heal alone couldn't help — it needs the anchor to exist in
the target, and the old target path didn't exist at all. Re-pointed `path` + re-authored the anchor/replacement
against the current `buildResponseFromSnapshots` (now `func (m *aireadiness.Manager)`, member load via the
exported `m.workforce.LoadMembers` → bounded via `m.workforce.LoadMembersByUserIDs`, type `workforce.Member`).
Recomputed `pre_sha256 130e58f7` / `post_sha256 ab611ed5`. The patch stays a PURE perf-only, data-identical
read-path relaxation (same rationale as M51/M219). rext `997272b`, tag `july-jitter-m254-aireadiness-repoint`
on origin. 0 platform-repo edits.

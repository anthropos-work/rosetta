**Type:** tik

# iter-04 — the prompt-hash cache (batchcache/)

Built component (3) of TOK-01's chain: the reproducibility + cost-control cache (corpus/ops/demo/
cache-spec.md).

## Work done
- **NEW `batchcache/cache.go`** — `Open(root, memberPrompts, captureVersion)` computes the batch hash +
  prepares `batch-${hash}/`; `Has`/`Get`/`Put` (atomic `.tmp`→`os.Rename`; `Put` REJECTS non-JSON so a
  malformed envelope never caches; a leftover `.tmp` is NOT a hit); `WriteManifest`; `BatchHash` (sha256
  of SORTED member prompts + NUL + capture version → order-independent batch dir, prompt/version change
  invalidates); `memberKey`/`MemberKey` (the per-member invalidation unit).
- **NEW `batchcache/lock.go`** — the `.lock` fence: `Lock` (OS-atomic `O_CREATE|O_EXCL`, records holder
  pid, returns `ErrLocked` if held), `Unlock` (nil-safe, idempotent), `BreakLock` (recover a stale lock),
  `Locked` (status).
- **NEW `batchcache/cache_test.go`** — 14 tests: hash determinism/order-independence/prompt+version
  invalidation, put/get/has, invalid-JSON reject, leftover-.tmp-not-a-hit, the REPRODUCIBLE RESEED
  scenario (run-2 hits both members, capture-version bump → different dir/miss), manifest, the lock fence
  (second lock fails, release re-locks), break-lock, nil-unlock, empty-root, member-key invalidation unit.

## Measurements
- `go build ./...` + `go vet ./...` clean; `gofmt -l batchcache/` clean.
- `go test ./batchcache/...` GREEN; full `go test ./...` GREEN. Test funcs 599 -> **613** (+14).
- The per-box cache root (`.agentspace/.batchcache`) is covered by the rosetta repo `.gitignore`
  (`.agentspace/`) — never committed.

## Close — 2026-06-26

**Outcome:** Component (3) landed — the prompt-hash cache (atomic writes, lock fence, capture-version
invalidation) with the $0 byte-identical reseed PROVEN in a unit test. 14 tests; full suite green. Gate
still 0/5 (the reproducibility MACHINERY exists + is unit-proven, but the gate's $0-reseed clause is
asserted on a REAL batch — that's the gate-proving run after cmd/gen-batch + the seeder land).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (0/5 — cache machinery proven in unit; the LLM-firing CLI + seeder + real run remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (sorted-prompt batch hash = order-independent), D2 (Put rejects non-JSON)
**Side-deliverables (if any):** none
**Routes carried forward:** iter-05 builds component (4): `cmd/gen-batch` — wire EffectiveBatches →
cache(hit?$0 : generate via services/ai) → the mandatory --max-cost guard + --max-concurrent + re-roll +
the cost report, under TOK-01.
**Lessons:** sorting the member prompts before hashing makes the batch dir independent of member order
(two equivalent descriptors share a dir) while per-member files still key by index — the right split.
Rejecting non-JSON at `Put` keeps the cache's $0-reseed invariant honest (a malformed response can't
poison the cache).

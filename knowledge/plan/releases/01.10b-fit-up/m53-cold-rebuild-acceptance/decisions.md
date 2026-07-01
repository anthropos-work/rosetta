# M53 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## KB-1 — AB2 is "prompt-free replay from a filled cache", not "auto-capture during bring-up"
**Context.** The Phase 0b audit (all 3 KB contract docs ALIGNED) clarified that `/demo-up` on a cold box
**replays only, never captures** (`up-injected.sh:665`; `snapshot-cold-start.md:110,198-211`). M47's
contribution is making the *operator's* one-time `stacksnap capture` turnkey via the MCP-configured DSN (no
`~/.pgpass`) — it does not auto-capture inside the bring-up.
**Choice.** M53 asserts AB2 as: the `.agentspace/snapshots/` cache is present + populated (verified: 1.4 GB,
taxonomy + directus + sim-embeddings each with COPY files + manifest.json), so the cold `/demo-up` set-dresses
by **replaying that cache with no prompt**. This is the accurate reading of "the cold-start MCP-DSN auto-capture
filled the snapshot with NO prompt" — the snapshot was filled (by M47's turnkey capture, already done), and the
cold rebuild consumes it prompt-free. **We do NOT wipe the cache** — that would require a live prod capture,
which is out of M53's scope and not what AB2 means.
**Why.** Matches doc+code; a truly-empty-cache assertion would contradict the (correct) replay-only bring-up
contract. Not a regression — the cache-fill is a completed M47 deliverable.

## KB-2 — AB5 asserts the shipped 78.4%/199, and validates render on the cockpit's actual link
**Context.** `ai-readiness.md:106` carries a round "~80% / ≈160 of 200" from the contract-writing phase; the
shipped funnel + `seeding-spec.md:369-375` are **78.4% / 199 frozen snapshots**. Also: the fast frozen read
path fires on a `?cycle=<closed>` deep-link, but the cockpit AI-readiness link (`cockpit.go:74`) is the bare
`/enterprise/workforce/ai-readiness` (no `?cycle=`); the M51 `app-aireadiness-snapshot-loadmembers` patch
bounds member hydration so the dashboard renders acceptably regardless.
**Choice.** M53 asserts AB5 against the shipped 78.4%/199 (1 completed + 1 started + manager), enabled/3-step,
and validates the dashboard **renders** (not a 180s timeout) on whatever link the cockpit + manager coverage
harness actually navigate to. A stale round-number in the doc prose is a doc-hygiene note (flag in §5), not an
acceptance failure.
**Why.** Assert against ground truth (code + seeding-spec), not the round contract number.

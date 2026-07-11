---
name: stack-list
description: List the live stacks — every dev-N and demo-N that's up, their type, N, status, and offset ports — from the one unified registry. Use when asked which stacks / environments are running.
argument-hint: (no args)
---

# Stack List — list the live dev + demo stacks

Reports the **unified stack registry** (M12) — which `dev-N` **and** `demo-N` exist, with the fields the
registry actually carries per stack: `type`, `n`, `status`, resolved host `ports` (the offset is `base + N·10000`),
`created`, and — for a stack brought up with `/demo-up --public-host <magicdns>` (v2.2 "panorama" M212) — the
additive `external_host` field (the reachable MagicDNS URL). One N-pool spans both kinds, so this is the single
source of truth for "what's running and which N is free". (Formerly `/demo-status`, now generalized to both stack
types.)

## Mission

Read the unified registry via either stack tooling — they share the same `stack-core/stack_registry.py`
N-pool, so either lists **all** stacks:
```bash
# via the demo tooling:
stack-demo/rosetta-extensions/demo-stack/rosetta-demo status
# or via the dev tooling (same unified registry):
stack-dev/rosetta-extensions/dev-stack/dev-stack status
```
The `stack-<role>/` prefix is the stack's per-stack consumption clone of rosetta-extensions (pinned at a
tag), distinct from the authoring copy at `.agentspace/rosetta-extensions/` where this tooling is built,
tested, and tagged. The `demo-stack` / `dev-stack` section names inside the repo stay as-is.

Then summarize for the user: which stacks are live (dev + demo), on which offset ports, and which N values are
**free** (re-allocatable). For a `demo-N`, `rosetta-demo status` *additionally* surfaces per-demo provenance
(the profile + the release tags each repo was cloned at) from demo-stack's **own** separate provenance registry
(`demo-stack/stacks/registry.json`) — that detail is demo-only; the unified registry itself does not store it, and
`dev-N` stacks have no equivalent provenance record.

To inspect one stack's containers directly: `docker compose -p dev-N ps` or `docker compose -p demo-N ps`.
The main `anthropos` dev stack (N=0) is a separate project; `docker ps | grep '^anthropos-'` shows it
independently.

## Related skills

| Skill | Use when |
|-------|----------|
| `/dev-up` · `/dev-down` | Bring up / tear down a dev stack |
| `/demo-up` · `/demo-down` | Bring up / tear down a demo stack |
| `/stack-seed` · `/stack-snapshot` | Seed / set-dress a listed stack |

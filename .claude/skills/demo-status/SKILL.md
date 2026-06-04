---
name: demo-status
description: Show the live demo stacks — which demo-N are up, their offset ports, profile, health, and the resolved per-repo refs. Use when asked about running demo environments.
argument-hint: (no args)
---

# Demo Status — list running demo stacks

Reports the rosetta-demo registry (which `demo-N` exist, their offset, profile, services, and resolved
per-repo clone refs) plus a live `ps` per demo project.

## Mission
```bash
anthropos-demo/rosetta-demo/rosetta-demo status
```
Then summarize for the user: which demos are live, on which offset ports, their health, and — useful for
reproduction — the release tag each repo was cloned at (from the registry's `clones` field).

To inspect one demo's containers directly: `docker compose -p demo-N ps`. The dev `anthropos` stack is a
separate project; `docker ps | grep '^anthropos-'` shows it independently.

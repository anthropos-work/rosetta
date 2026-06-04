# M3 — Decisions

## M3-D1 — per-demo service-repo clones (user-chosen, 2026-06-03)
Each `demo-N` re-clones the platform service repos under its own `anthropos-demo/stacks/demo-N/`, rather than
all demos sharing the single `anthropos-dev/*` clones. **Chosen for full filesystem isolation** + future-proofing
per-demo config divergence, accepting ~N× disk + clone time. Note: since every demo uses the *same* Clerkenstein
injection, contention was not the deciding factor — isolation was. (Alternative: shared clones, disk-cheaper.)

## M3-D2 — manual teardown only (user-chosen, 2026-06-03)
`/demo-down [N]` is the only reclaim path; no nightly auto-reaper in M3. Keeps M3 tight + avoids a teardown-safety
concern. (Alternative: a cron/systemd reaper of `demo-*` older than X — deferred; add later only if forgotten
stacks become a real problem.)

## M3-D3 — clone each repo at its latest release tag, not `main` (user-directed, 2026-06-03)
The `/demo-up` clone step (M3-D1) checks out **each platform service repo at its most recent release tag**, not
its default branch — so a demo runs a *released* version, reproducibly, rather than bleeding-edge `main`.
**Resolution order, per repo:**
1. If the caller passes a ref at skill-call time (global or per-repo, e.g. `/demo-up 1 --ref app=v2.4.0` or a
   `--ref main` override) → use that.
2. Else → the **latest release tag** (prefer semver `v*` tags by version order; fall back to the most recent tag).
3. Else (repo has **no tags**) → fall back to the default branch (`main`).
The resolved ref per repo is recorded in the stack registry (so `/demo-status` and reproduction show exactly what
each demo is running). Clerkenstein injection (go.mod replace + skip-worktree) is applied **on top of** the
checked-out tag. (Open: exact "release tag" pattern — `v*` only vs any tag — settle in S1 against the actual org
repos' tagging conventions.)

## M3-D4 — tooling home: a new gitignored `anthropos-demo/rosetta-demo/` repo (build, 2026-06-03)
The rosetta-demo tooling (override generator, lifecycle CLI, registry, clone/injection logic) lives in a
**new gitignored repo at `anthropos-demo/rosetta-demo/`** with its own git — the same pattern as
`clerkenstein/`. The user-facing **`/demo-*` slash-command skills** and the **`corpus/ops/rosetta_demo.md`**
guide live in the **rosetta** repo (tracked, on the `m3/rosetta-demo` branch). Rationale: `anthropos-demo/`
is gitignored from rosetta, so runtime tooling can't be rosetta-tracked; skills + docs are rosetta's domain.

## M3-D5 — acceptance bar adjusted to "1 demo stack alongside the dev stack" (user-directed, 2026-06-03)
The design acceptance was *two* concurrent full stacks. The dev box is **16 GB with Docker's VM capped at
~8 GB and a full `anthropos` dev stack already running** — two full stacks (~20-24 GB) is physically
impossible here. User direction: "do the best you can; 1 is ok; if you can make a demo live alongside the
dev stack without messing with it, that's already a great test." So the **acceptance is one isolated demo
stack co-resident with the dev stack**, proving the isolation mechanism. The full 12-service single-stack
run + end-to-end Clerkenstein browser login are **resource-gated** → verified on a bigger Docker VM / box
(documented in the ops guide), not on this hardware.

## Open (resolve during build)
- Max-N + exact port-offset sizing (below the ephemeral range, no overlap with the 24 base mappings).
- The `clerk-backend` `api.clerk.com` → fake-BAPI redirect mechanism **inside Docker** (extra_hosts + trusted CA
  vs a base-URL env override). Load-bearing — spike in S3.
- Whether the v1.0 express-gate CI carry-forward lands here vs M5 (default: M5).

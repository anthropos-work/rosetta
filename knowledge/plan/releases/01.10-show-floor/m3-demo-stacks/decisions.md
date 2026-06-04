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

## M3-D6 — tooling renamed `demo-stack` → `rosetta-demo` (user-directed, 2026-06-04)
The demo-environment tooling (repo dir, CLI, ops guide) was renamed from `demo-stack(s)` to **`rosetta-demo`**
and pushed as a **private org repo** `anthropos-work/rosetta-demo`. The CLI `demo-stack` → `rosetta-demo`; the
repo dir `anthropos-demo/demo-stacks` → `anthropos-demo/rosetta-demo`; the guide `corpus/ops/demo_stacks.md` →
`corpus/ops/rosetta_demo.md`. **Preserved** (milestone identity, not the tool): the `m3-demo-stacks` milestone
dir, `slug: demo-stacks`, the title "Disposable multi-instance demo stacks", the `/demo-up|down|status` skill
names, and `demo-N`/`-p demo-N` instance naming. Archived v1.0 records left frozen. Verified: 78 tests green +
deploy gate 100%/100% under the new name. (So the milestone keeps its name; its *deliverable* is now rosetta-demo.)

## M3-D7 — extended-work close: deferral routings (2026-06-04)
The post-close extended work (full injected stack + deployment/injection alignment surface + harden + corpus +
rename) was closed via `/developer-kit:close-milestone` against the already-`done` M3 (no branch merge — the work
was committed directly on `release/01.10-show-floor` + the two scratchpad mains, all pushed). The deferral audit
([audit-deferrals/deferral-audit-2026-06-04-m3-extended-close.md](audit-deferrals/deferral-audit-2026-06-04-m3-extended-close.md),
**YELLOW**) routed the open items: **clerk-backend cert/redirect → M5** (Fate 3, annotated), **clerk-webhook live
POST → M4** (Fate 3), **the demo login identity `user_clerkenstein` + the casbin plural/singular gotcha → M4**
(Fate 3, annotated — prevents the seed-the-wrong-user 403 trap), **express-gate CI → M5** (Fate 2, already owned).
M3-CF1 RESOLVED. The nightly auto-reaper stays a deliberate non-goal (M3-D2).

## Resolved (were "Open during build")
- **Port-offset sizing** — `demo-N → base + N·OFFSET`; **OFFSET defaults to 10000** (raised from 100 at the
  extended-work close, 2026-06-04, after the close review caught that 100 maps demo-1's storage `8300→8400`
  straight onto the dev stack's jobsimulation `8400` — a real collision the base `up` path had shipped with).
  Both the base `rosetta-demo up` path and `up-injected.sh` now use `N·10000`; trade-off is **max-N ≈ 5**
  (`base + N·10000` must stay < 65535). Fixed in `rosetta-demo` line 16 + the ops guide.
- **`clerk-backend` redirect mechanism** — resolved to `extra_hosts: api.clerk.com:<ip>` (base path) + a
  `network alias` on the fake-BAPI service (injected path); the TLS-cert trust is the M5 Fate-3 item above.
- **Express-gate CI location** — M5 (confirmed, M3-D7).

## Adversarial review (extended-work close, 2026-06-04)
Scenarios the close review considered against the extended orchestration code:
- **Port-offset collision (CONFIRMED bug, fixed).** With `OFFSET=100`, `demo-N`'s storage port `8300 + N·100`
  lands on a dev base port for small N (demo-1 → `8400` == dev jobsimulation). The base `rosetta-demo up` path
  had this; `up-injected.sh` didn't (`N·10000`). Fixed the default to `10000` (both paths now identical), with
  `max-N ≈ 5` documented. A misconfigured `DEMO_OFFSET` smaller than the ~3000 base-port spread re-introduces it
  — documented as an operator caveat in the ops guide.
- **registry.json lost-write under concurrency (fixed).** Two concurrent `up`/`clone` could read-modify-write
  the shared registry and clobber each other (multi-instance is the tool's purpose). Fixed with an `fcntl`
  exclusive lock around every registry RMW — portable across macOS + Linux (`flock(1)` is Linux-only). A
  deterministic regression test for a bash-internal RMW would be timing-sensitive (would risk the flake gate);
  the lock is correct-by-construction (proven inline against 20 concurrent writers) and the scenario is recorded
  here per the adversarial-review protocol.
- **Partial bring-up (accepted + documented).** If `up-injected.sh` fails mid-way under `set -e`, clones/containers
  linger; `rosetta-demo down N --purge` reclaims cleanly (`-p demo-N`-scoped). Documented as the recovery step in
  the ops guide rather than auto-rollback (a trap-based rollback is M5 polish, not load-bearing).

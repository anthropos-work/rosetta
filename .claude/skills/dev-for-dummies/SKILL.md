---
name: dev-for-dummies
description: One-command LIVE-development setup (or resume) for working on one or more TARGET platform repos through Rosetta. Pins rosetta + rosetta-extensions to clean updated main, brings up a remote-accessible demo stack (demo-N over Tailscale --public-host), and runs each TARGET natively with hot-reload from an isolated feat/<name> or fix/<scope> worktree (also remote-served), so live code edits show up instantly. Detects and offers to RESUME an earlier setup. Use to START or RESUME feature/fix work on the platform.
argument-hint: [target-repo ...] [--feat <function-name> | --fix <scope>] [--public-host <magicdns-fqdn>] [N] [--resume | --fresh]
---

# Dev-for-Dummies — live development setup (or resume) via Rosetta

One skill that puts you in front of a **running, remote-shareable copy of the platform** with the one or
more repos you actually want to change **running live** — so every edit you make is reflected immediately,
on an isolated branch, without ever touching anything you shouldn't.

It does the boring, error-prone setup for you and it is **safe to re-run**: if a setup already exists (or an
earthquake wiped your box and you're starting over), it finds it and offers to **resume** instead of
clobbering it.

> **Two words up front so nothing surprises you:**
> - **TARGET** = a platform repo you want to edit live this session (e.g. `next-web-app`). Everything you
>   change lives on a dedicated branch/worktree of a TARGET. **Nothing else** is ever edited.
> - **demo stack** = a throwaway, seeded, isolated copy of the whole platform (`demo-N`) that your TARGETs
>   run against. Brought up with `/demo-up` and reachable remotely over Tailscale.

Sources of truth this skill drives: [`/demo-up`](../demo-up/SKILL.md) ·
[`corpus/ops/rosetta_demo.md`](../../../corpus/ops/rosetta_demo.md) ·
[`corpus/ops/demo/tailscale-serve.md`](../../../corpus/ops/demo/tailscale-serve.md) ·
[`corpus/ops/demo/frontend-tier.md`](../../../corpus/ops/demo/frontend-tier.md).
Detailed commands (port map, native-run recipes, tailscale wiring, the resume manifest format) live in
[`reference.md`](reference.md) — read it before running the technical phases.

---

## Do these phases IN ORDER. Each has a hard gate.

Each step is **FATAL** (stop with one plain-English line + what to do) or **NON-FATAL** (warn plainly, keep
going). Never say "all set" if a FATAL step failed; never abort a good stack over a NON-FATAL one. Per-step tags
are in [`reference.md`](reference.md).

### Phase 0 — Session gate: model + effort (**do this first; block on failure**)

This work is only worth doing on a strong, high-effort session. Before anything else:

1. **Model** — confirm you are running on **Opus (4.8 or better)** or **Fable**. You know your own model from
   your system prompt. If you are on Sonnet/Haiku or anything weaker: **STOP** and tell the user, in one plain
   line, to run `/model` and pick **Opus 4.8** (or **Fable**), then re-invoke the skill. Do not continue.
2. **Effort** — the requirement is **ultrahigh** effort (`/effort max`). You **cannot reliably read your own
   effort setting**, so do not guess: tell the user "this skill needs ultrahigh effort — please run
   `/effort max` now (or confirm you already have)" and **wait for their confirmation**. If they can't/won't,
   stop and let them decide. Only proceed once model ✓ and effort ✓.

> Why block: the later phases make live, cross-repo, remotely-served changes. A weak or low-effort session
> here causes exactly the "it hacked the wrong repo" failures this skill exists to prevent.

### Phase 1 — Pin **rosetta** and **rosetta-extensions** to clean, updated `main`

The environment baseline must be current and unpolluted before you build on it. This covers **two** repos:

| Repo | Path | What "clean" means |
|------|------|--------------------|
| **rosetta** (this corpus) | `/home/devops/rosetta` | on `main`, no pending changes, == `origin/main` |
| **rosetta-extensions** (authoring copy) | `/home/devops/rosetta/.agentspace/rosetta-extensions` | on `main`, no pending changes, == `origin/main` (clone it if absent) |

For **each** repo: `git fetch origin`, then check the branch and working tree.

- **Already clean on `main` and current** → fast-forward (`git pull --ff-only`). Done.
- **On another branch, OR has any pending change** (staged / unstaged / untracked that isn't ignored) →
  1. **ALERT the user in plain English** — one or two lines: which branch it was on and a **named list** of the
     pending files (e.g. "5 screenshots + your uncommitted skill edit"), so someone watching files disappear
     knows exactly what moved. No jargon dumps.
  2. **Make it recoverable, never destroy** — stash everything including untracked:
     `git stash push -u -m "dev-for-dummies safety backup <UTC timestamp>"`. Give the user the exact way back:
     `git stash list` then `git stash pop` (or `git stash apply stash@{0}`). (`git stash -u` correctly leaves
     gitignored paths — `.agentspace/`, `stack-*/` — untouched, so the snapshot cache, secrets, and stack clones
     are safe.)
  3. **Force to clean updated main**: `git checkout main && git reset --hard origin/main`.
- **rosetta-extensions absent** → clone it on main (`git clone git@github.com:anthropos-work/rosetta-extensions
  .agentspace/rosetta-extensions`). Non-fatal if the clone fails — note it and continue; the demo doesn't need
  the authoring copy to come up.

> **Self-preservation gotcha (important).** This skill's own files live under
> `rosetta/.claude/skills/dev-for-dummies/`. If the skill has **not yet been committed to `origin/main`**, the
> `reset --hard origin/main` above would remove it from disk. The `git stash push -u` in step 2 keeps it
> **recoverable**, and you already hold the instructions in context, so execution continues — but say so
> plainly: *"your `dev-for-dummies` skill isn't committed to main yet; I've stashed it (recoverable) — commit
> it to main so it stops disappearing on clean-up."* Prefer committing the skill to `main` once, so this stops
> mattering.

> **Not a contradiction with the "don't touch rext" rule (Phase 6).** Phase 1 *pins* rosetta + the rext
> authoring copy to clean main as the **setup baseline**. The **demo** consumes rext from its **own** pinned
> clone at the tag in `.agentspace/rext.tag` (currently `v2.2`) — `/demo-up` manages that; **do not** repoint
> it to main. After setup, development never touches rosetta or rosetta-extensions again.

### Phase 2 — Decide the TARGET repo(s) and branch names

1. **TARGETs** — if repos were passed as arguments, use them. **If not, ASK the user** which repo(s) they want
   to work on live. Offer the common ones and let them pick one or more:
   `next-web-app` (the main frontend — the usual answer), `app` (Go backend/API + AI-readiness + skills),
   `cms`, `jobsimulation`, `skillpath`, `studio-desk`, `ant-academy`. (Frontend/UI targets are the smooth path;
   a backend Go target has extra wiring caveats — see [`reference.md`](reference.md) § *Backend targets*.)
   **Validate each pick** with `test -d stack-demo/<repo>`; if it doesn't exist, reject it and show the
   valid-repo list (reference.md § *Port map*) — don't let an opaque `git worktree` error surface later.
   **`hiring` is not its own repo** — it lives inside `next-web-app` (target `next-web-app`, run `pnpm
   dev:hiring` on `3001+OFF`).
2. **Branch name per target** — you work on a dedicated branch, never on the repo directly. Names are
   **human-readable and conventional, no junk**:
   - a new feature → **`feat/<function-name>`** (e.g. `feat/ai-readiness-export`)
   - a fix → **`fix/<scope>`** (e.g. `fix/profile-avatar-fallback`)

   If the user gave `--feat`/`--fix`, use it. Otherwise ask "feature or fix?" and get a short kebab-case name.
   Reject vague names (`feat/stuff`, `fix/bug`) — ask for something that names the actual thing.

### Phase 3 — Resume check (**before creating anything**)

An earlier setup may already exist (you started this before; or the box was wiped and you're rebuilding). **Do
not create blindly.** Look for prior state and offer to resume it:

1. Read any manifests in `.agentspace/dev-for-dummies/session-*.yaml` (format in [`reference.md`](reference.md)
   § *Session manifest*).
2. In each candidate TARGET's clone, run `git worktree list` and look for existing `feat/*` / `fix/*` worktrees.
3. Check `/stack-list` for a live `demo-N` recorded by a manifest.

If you find prior state, present it to the user in **plain English** and let them choose — use a real question,
not a guess. Give them what a human needs to recognise it:

> *"I found a setup from **Mon 14 Jul 2026, 15:32** — feature **`ai-readiness-export`** on **next-web-app**
> (branch `feat/ai-readiness-export`), demo-3 at `https://calypsostaging.taildc510.ts.net:13000`. Resume this,
> or start fresh?"*

- **Resume** → skip creation. Verify the worktree + branch still exist (if the branch exists but the worktree
  dir was wiped, re-attach **without** `-b` — see reference.md § *Resume checks*). **Check `/stack-list` first:**
  if demo-N is **up**, reuse it as-is — do **not** re-run `/demo-up` (a bare re-run re-does the slow
  set-dress/seed and can bounce the peers your native TARGET depends on); only run `/demo-up N --public-host …`
  when it's **down**. Relaunch the native TARGET process(es) + their `tailscale serve`; refresh the manifest
  timestamp. Don't recreate anything that's already there.
- **Fresh** → continue to Phase 4. (Leave the old worktree/branch/manifest alone unless the user asks to remove
  them.)

### Phase 4 — Bring up the remote-accessible demo stack

Bring up (or reuse) a `demo-N` **built for remote access** — the default assumption is that the demo, and the
live TARGETs, are reachable from another machine on the tailnet.

1. **Public host** — this box's own MagicDNS FQDN. Discover it (`tailscale status --json | jq -r '.Self.DNSName'`,
   strip the trailing dot) — on this box it is `calypsostaging.taildc510.ts.net`. If tailscale isn't up or has
   no FQDN, tell the user and either get one from them or fall back to a `localhost` demo (warn that remote
   access won't work).
2. **One-time host readiness** — ensure `sudo tailscale set --operator=$USER` has been run (so `tailscale
   cert`/`serve` work un-sudo'd and the remote browser sees a *trusted* cert). See [`reference.md`](reference.md)
   § *Host prereqs*.
3. **Bring it up** via `/demo-up` (do **not** re-implement it): `/demo-up N --public-host <fqdn>`. This gives you
   the full seeded Stories-&-Heroes world on offset ports, Clerkenstein-wired, HTTPS-served over Tailscale.
   Note the allocated **N** and the port map (`reference.md` § *Port map*).
4. **Known non-self-healing gotchas** on this box (from the memory + `tailscale-serve.md`): the 0.0.0.0
   container-bind vs `tailscale serve` shadowing fix, the snapshot-cache digest re-capture, the backend
   network-detach recreate. Apply them **only if** the demo verify surfaces them — recipe in
   [`reference.md`](reference.md) § *Known gotchas*.

### Phase 5 — Put each TARGET **live** (native + hot-reload + remote-served) from its worktree

This is the point of the skill: your edits reflect **immediately**, and the running copy is the demo everyone
can see. Full, **verified** commands are in [`reference.md`](reference.md) § *Run a target live* — follow them;
the summary below is the shape, not a substitute.

0. **Prereq (FATAL for a frontend target):** node **≥ 24** on PATH — **do not assume `nvm`** (it may be absent
   and the system node may be older). Check in a login shell; if `< 24`, STOP and ask the user to put node 24 on
   PATH, then resume.
1. **Isolated worktree + branch** off `stack-demo/<repo>` into `stack-demo/.worktrees/<repo>-<slug>`
   (`git -C stack-demo/<repo> worktree add -b feat/<name> ../.worktrees/<repo>-feat-<name>`).
2. **Stop that one container** so the native process owns the offset slot (`$DC stop <svc>` — leaves every other
   service running; also clears the F12 serve-shadow for that port).
3. **Assemble the native env + run it**, bound to `127.0.0.1` on the **offset** port (e.g. 13000, so it inherits
   the demo's baked CORS origins + minted pk), in a **detached login-shell tmux** session:
   - **Frontend:** author a fresh `apps/web/.env.local` — **do NOT copy the demo's `.env.local`; `/demo-up`
     trap-deletes it after every build.** Take the minted pk from `$STACK/.env.demo-$N` (FATAL if absent) + the
     offset `NEXT_PUBLIC_*` URLs, then `pnpm exec next dev -H 127.0.0.1 -p <offset>`.
   - **Backend (more caveated):** `go run .` reaches **nothing** until you **rewrite**
     `DB_CONNECTION`/`REDIS_ADDR`/`*_RPC_ADDR` to `localhost:<base+off>` (they live in `docker-compose.yml`, not
     `.env`). And router→native-subgraph federation needs host-gateway wiring on the router — **flag that as a
     tooling gap to raise, don't hack the stack.**
4. **Serve it remotely** (idempotent): `tailscale serve --bg --https=<offset> http://127.0.0.1:<offset>`.
5. **Verify (NON-FATAL):** open the live URL, log in, edit a string in the worktree, confirm it hot-reloads.

### Phase 6 — Record the session manifest

Write `.agentspace/dev-for-dummies/session-<feature-slug>.yaml` with everything a human (or a rebuild) needs:
created timestamp (human-readable + UTC), model + effort, demo N + public host + live URLs, per target (repo,
branch, worktree path, native port, tmux session), and **`allowed_edit_roots`** (the worktree paths Phase 7's
guard checks every edit against). Format in [`reference.md`](reference.md) § *Session manifest*. This is what
Phase 3 reads on the next run.

---

## After setup — how this session must run

### Phase 7 — Standing rules for the rest of THIS session (emit these to the agent + user)

Once setup is done, the session that launched this skill must operate under these rules until it ends. State
them back plainly so both you and the user are on the same page:

- **🔒 Edit ONLY the TARGET worktree(s).** For no reason change, hack, patch, or "quick-fix" **any other repo,
  any stack, the demo tooling, or rosetta / rosetta-extensions itself.** If something outside a TARGET looks
  broken or in the way, **stop and report it — do not touch it.** This is non-negotiable: touching anything
  outside the TARGETs is the failure mode this whole skill exists to prevent.
  - **Make it mechanical, not just a promise:** before **every** Edit/Write, assert the target file's absolute
    path is under one of the `allowed_edit_roots` recorded in the session manifest (Phase 6). If it is not,
    **refuse the edit and report it** — no exceptions, even if the fix looks trivial or "obviously needed."
- **🗣 Explain in plain English.** End every message — and the session, where it makes sense — with a short,
  jargon-free account of what you did and why. Introduce a term briefly if you must use one. Be **brief**: the
  user has a life and doesn't want to read walls of text.

### Phase 8 — Hand off to the human (brief)

Give the user a short intro (a few lines, plain English):
- **What's set up** — which TARGET(s) are running live, on which branch, against demo-N.
- **Where to look** — the live URL(s) to open in a browser to watch changes land as work is done (the
  tailscale-served TARGET URL + the demo app). The presenter cockpit is at `:7700+`, but it's **plain HTTP** and
  **not** fronted by `tailscale serve`, so treat it as reachable **on this box**, not a guaranteed remote link.
  Explain the loop plainly: "as I edit the code in `<worktree>`, the page hot-reloads — just refresh to see it."
- **One next step** — "tell me what to build/fix and I'll work only inside the TARGET."

### Phase 9 — When the user asks "what next / how do I move forward"

Tell them every development session must **end** with these three steps (in order):

1. **Hardening** — write tests for everything you wrote, **proportional to the change**: unit + integration +
   e2e as appropriate, **plus knowledge/documentation coverage** (docs updated for what changed).
2. **Formal close prep** — open a PR with a description a human can read in **under 5 minutes**: a short intro
   first (what & why), then the specifics.
3. **Formal auto PR review** — hunt for bugs, missing pieces, and misalignments across **all three**: the
   **code**, the **knowledge/docs**, and the **tests**.

### Phase 10 — Wrapping up (when the feature is truly done — user-initiated, AFTER Phase 9)

Don't leave the box in a half-state. Once the PR is opened + reviewed, **offer** the user a clean teardown
(confirm before each; full commands in [`reference.md`](reference.md) § *Wrap up / cleanup*):

- Stop the native process (`tmux kill-session`) and drop its remote proxy (`tailscale serve --https=<port> off`).
- **Put the stack back** — *either* restart the container you stopped (`$DC up -d --no-deps <svc>`, if you might
  return to this feature) *or* tear the whole demo down (`/demo-down N`, if you're finished — frees ~10–12 GB).
  Make that choice explicit; don't guess.
- Once merged, remove the worktree (`git worktree remove …`) + delete the branch, and remove the session
  manifest so a later run doesn't offer to resume finished work.

---

## Safety recap (the load-bearing part)

- All changes land in a `feat/<name>` / `fix/<scope>` **worktree** of a TARGET — isolated by construction.
- **Nothing outside the TARGETs is ever edited** — enforced mechanically (Phase 7): every edit's path is
  asserted against the manifest's `allowed_edit_roots` first, or it's refused. Rosetta + rext are pinned clean
  in Phase 1 and then frozen.
- The demo is `-p demo-N`-scoped and can never touch the dev stack (see [`/demo-up`](../demo-up/SKILL.md) safety).
- Phase 1's "drop pending changes" is **recoverable** — it stashes (incl. untracked) before any reset, never
  hard-deletes user work.

## Related skills

| Skill | Use when |
|-------|----------|
| `/demo-up` · `/demo-down` | The demo stack this runs on (bring up / tear down) |
| `/stack-list` | See the live demo-N + its ports |
| `/stack-seed` · `/stack-snapshot` | Re-seed / re-set-dress the demo if its data looks thin |
| `/test-platform` | Verify the running stack when something looks off |

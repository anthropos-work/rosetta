# iter-262 — decisions

## `D-M257x-262-1` — the demo's `platform/.env` is 31 concatenated copies with a BLANK `DIRECTUS_TOKEN`

Found while provisioning the dev stack's `.env`, and it is a defect in a file three green demo
bring-ups depend on — so it is recorded whatever this iter's outcome.

**Measured (values-blind — no secret value was read, echoed or logged; only key names and a
blank/non-blank predicate):**

| file | lines | keyed lines | **unique keys** | `DIRECTUS_TOKEN` |
|---|---|---|---|---|
| `stack-demo/platform/.env` | 470 | 408 | **18** | present **31×, BLANK every time** |
| `.agentspace/secrets/platform/.env` (the real source) | 23 | 15 | **15** | present **1×, NON-BLANK** |

`stack-demo/platform/.env` is **the same 18-key block appended 31 times**. Something in the demo path
**appends rather than replaces**, and it has run 31 times. Compose takes last-wins so the stack still
works, which is exactly why nobody noticed.

**The blank `DIRECTUS_TOKEN` is the interesting half.** `CLAUDE.md` names it as the one variable you must
fill and calls a missing value *"the classic stack-boots-catalog-empty failure."* The demo does not feel
it because a demo runs a **per-stack Directus** (`--local-content` default-on) and replays content into
it — the prod token is never needed. **A dev stack without `--local-content` reads content live from
prod**, so the same file would have produced an empty catalog on the dev side and it would have looked
like a platform fault.

**Decision:** do **not** provision the dev `.env` by copying the demo's. Provision it the documented way —
`platform/.env_example` overlaid with `.agentspace/secrets/platform/.env` — which yields **63 unique keys,
0 duplicates, `DIRECTUS_TOKEN` NON-BLANK**. The first attempt in this iter *did* copy the demo file and was
replaced before `make up` ran; both states are recorded rather than the second one alone.

`COMPOSE_PROJECT_NAME=anthropos` was verified explicitly (it is configuration, not a credential) because it
determines container naming and therefore collision risk with `demo-1` / `demo-2`. It is the dev project
name; no collision.

**Routed:** `FIX-M257x-262-demo-env-append-is-not-idempotent` — find the writer and make it replace-or-skip,
and decide whether a blank `DIRECTUS_TOKEN` should fail the demo's own secret-coverage gate rather than pass
it. `corpus/ops/secrets-spec.md` is the home; the `/stack-secrets` DNA is what should have caught it.

## `D-M257x-262-2` — the DOCUMENTED dev bring-up cannot build `backend` from a fresh clone

**This is the finding the user's "demo AND dev" closing condition was written to catch, and it could not
have been found on the demo path.**

`make init` + `make up`, exactly as `CLAUDE.md` and `corpus/ops/setup_guide.md` document them, **fail** on a
fresh `stack-dev/`:

```
#50 [backend stage-1 5/6] COPY --from=build /build/studio ./studio
#50 ERROR: failed to calculate checksum of ref …: "/build/studio": not found
target backend: failed to solve …
make: *** [up] Error 1     (UP_EXIT=2)
```

**Why it is structural, not local:**

| fact | evidence |
|---|---|
| `app`'s image hard-COPYs the studio tree | `stack-dev/app/Dockerfile:45` `COPY --from=build /build/studio ./studio`, then `:46` `pip install -r studio/requirements.txt` |
| the tree is **gitignored**, so no clone ever carries it | `stack-dev/app/.gitignore:79` `studio/*`, commented *"pulled at build via additional_repo, like cms"* |
| `repos.yml` does not list it and `make init` does not fetch it | `make init` clones exactly `app`, `sentinel`, `next-web-app`, `studio-desk` |
| on current main there is **no submodule and no Make target** either | `demo-stack/lib/studio.sh` header: `851cf3fb` deleted the `.gitmodules` + gitlink while **keeping** the Dockerfile COPY, so *"the only sanctioned acquisition is an out-of-band clone into the gitignored path"*; `app` has no `init-studio` target (that is a **cms** target) |

**And the fix for this exact bug already exists — for the demo only.** M257 B2 shipped
`rosetta-extensions/demo-stack/lib/studio.sh` (`studio_required` / `studio_populated` / `STUDIO_REPO`),
sourced by `ensure-clones.sh` **and** `up-injected.sh`, deriving the need from each clone's own Dockerfile.
It is why iter-258's demo log prints *"app: studio/ already populated — reusing (idempotent)"* and why three
green `--purge` cycles never saw this. **`rosetta-extensions/dev-stack/` contains no studio handling at
all** — grepped: its only `studio` hits are unrelated `studio-desk` Directus-token lines in
`dev-setdress.sh`.

So the bug M257 B2 believed it had closed is **half-closed**: closed on the path that gets exercised,
open on the path the corpus tells a new engineer to follow. A first-time dev setup on any box has been
broken since `fdb8034a` (2026-07-27) and nothing detected it, because nobody ran the documented dev path
cold.

**Unblocked in-iter the sanctioned way** — `git clone git@github.com:anthropos-work/anthropos-studio-room.git
stack-dev/app/studio` (`STUDIO_REPO` verbatim; landed at `aeec036`), which is exactly what `app`'s own CI
does via `additional_repo: "anthropos-studio-room:studio"`. **Zero platform-repo edits**: the path is
gitignored, so `app`'s tracked state is unchanged.

**Routed — and this one is a tooling deliverable, not a doc note:**
`FIX-M257x-262-dev-path-needs-the-studio-acquisition` — hoist `lib/studio.sh` to a section both stacks
share and call it from the dev bring-up, so the derivation stays Dockerfile-driven rather than
service-named. **A fence is the real ask**: the demo proves the predicate works, and nothing asserts the
dev path *uses* it.

## `D-M257x-262-3` — `INVITATION_HMAC_SECRET` is a boot requirement `.env_example` does not declare, and its absence exits **0**

`backend` started, logged, and **stopped**:

```
ERROR INVITATION_HMAC_SECRET is not set; invitation tokens will be insecure
ERROR can't create invitation token manager error="INVITATION_HMAC_SECRET is not set; …"
```

`docker inspect` → **ExitCode 0**, no error string.

| file | declares it? |
|---|---|
| `stack-dev/platform/.env_example` @ `0c91421` | **NO** |
| `.agentspace/secrets/platform/.env` (the secret source) | **NO** |
| `stack-demo/platform/.env` | **yes** — which is why the demo never saw this |

So a dev `.env` provisioned **exactly as documented** — `.env_example` overlaid with the sanctioned secret
source — **cannot start `backend`**, and the failure presents as a **successful exit**: no crash loop, no
non-zero code, no restart, nothing a health check or `docker ps` reads as broken. The container is simply
absent. Compare `sentinel`, which failed honestly in the same bring-up with `Restarting (2)` and was
diagnosable in one `docker logs`.

**Routed:** `FIX-M257x-262-invitation-hmac-secret-undeclared` — add it to `.env_example` (upstream, so a
platform PR, out of scope here) **and** to the `/stack-secrets` DNA so coverage-check catches it. The
`corpus/ops/setup_guide.md` critical-variable list and `CLAUDE.md`'s own hand-maintained tuple (which
iter-237 verified against `.env_example`, and which therefore could not have caught this — the variable is
in neither) both need the entry.

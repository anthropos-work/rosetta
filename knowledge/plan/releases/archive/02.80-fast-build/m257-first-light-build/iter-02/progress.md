**Type:** tik · **Active strategy:** `TOK-01` (steps 1–2)

## Work

### (a) odysseus provisioned — and it brings a stack up

`up-injected.sh 1 --public-host odysseus.taildc510.ts.net` → **rc=0**, 16/16 containers Up, remote
HTTPS reach working, 42,790 taxonomy skills replayed, 12 of 13 autoverify checks green.

| prereq | outcome |
|---|---|
| **Go** | `go1.26.5`, `GOROOT=/usr/local/go`. **Not installed over** — PATH fixed two ways (`/etc/profile.d/golang.sh` for login shells **and** `/usr/local/bin` symlinks, because the tooling does use non-login `ssh host 'cmd'`). Verified in both modes. |
| **atlas** | `v1.2.4-e282f76-canary`, the doc's exact command. |
| **Node** | **A 7th prereq the docs do not list** — `ant-academy.sh` aborts without Node ≥ 22, and autoverify grades the academy. nvm had v22.23.2 but only via interactive-only `.bashrc` — **the same login-shell hole as Go.** |
| **git/GitHub** | No org SSH key; used the documented PAT-over-HTTPS path, primed **values-blind** (stdin only). `make init` cloned 12 repos; private `colony v0.34.3` fetched. |
| **tailscale** | operator set; `tailscale cert` mints a **trusted LE** leaf without sudo (not mkcert). |
| **rext @ pinned tag** | `fast-build-m257-iter-01`, **rung zero verified on the host** via `git ls-remote` before cloning; clone left git-clean at the tag. |

**Go 1.26.5 vs the declared 1.25.12 — resolved, not assumed** (`TOK-01` known-context #6): **6/6 rext
modules build**, `go build ./...` rc=0 each. Root build fails by design.

### (b) both inherited `autoverify` fixes landed — rext `06761b5`

New check **(h) CONTAINER LIVENESS** + check **(d)** rebuilt as a 5-rung TLS-stack-independent ladder.
207 `stack-verify` tests pass (from 168 + **3 pre-existing failures, now fixed** — the M256 close's
Python roster omitted `stack-verify/tests`, which is how they survived). Negative controls proven by
**mutating production code back to the defect**: 10 / 13 / 8 / 7 tests go red per mutation.

**Two corrections to the audit's framing**, both verified first-hand:

1. **The row-less gap is 3 containers, not 2** — `hiring-app` was missed alongside `fake-fapi`/`fake-bapi`.
   13 checked vs 16 compose services, which is **exactly** the "14 of 16 Up" M256 measured.
2. **Check (d) does NOT fire on Linux/OpenSSL.** The failing client is macOS system `curl` on LibreSSL,
   while `openssl s_client` on the **same** LibreSSL completes a full handshake on the same leaf — so it
   is a **`curl`** defect, not a LibreSSL one, and not a cert problem. **It was never a gate risk on the
   gate host** (`TOK-01` known-context #2 answered — and the answer was "no", which is why the doc
   forbade assuming in either direction). Also established: the `http://` fallback **could never have
   succeeded on any demo** — the fake FAPI serves *either* TLS *or* plain HTTP, never both, and the
   override always sets the TLS cert. Both legs were always one boolean.

### (c) autoverify proven able to go RED — twice, and one was unplanned

- **By construction:** the mutation battery above.
- **In the wild, immediately:** the first odysseus bring-up produced `warnings:1, green:false`, and
  separately left `demo-1-directus-1` at **`Exited (1)`** on a cold-data-dir schema race —
  **precisely the class check (h) was built for.** The instrument earned its keep inside the iter that
  shipped it.

## 🔴 Two blockers surfaced that make the GATE CURRENTLY UNREACHABLE

Neither is a host problem and neither needs user authorisation — both are rext-side, in-remit, and
routed to iter-03 as its mandatory precondition. But **the v2.8 READY definition is unsatisfiable until
they land**, so no number measured before then is gateable.

### B1 — the platform dropped the `local_*` session mirrors; six seeders still write to them

`app/terraform/migrations/20260729133514.sql` (**2026-07-29**, *"Collapse the local_\* session
mirrors"*) does `DROP TABLE local_jobsimulation_sessions` / `local_skill_path_sessions`, re-pointing at
canonical `public.job_simulation_sessions` / `public.skill_path_sessions` (both confirmed present).
**Six `stackseed` seeders fail identically** — `hiring-funnel`, `personas`, `feedback`, `assignments`,
`content-stories`, `content-stories-nonsim`:

```
copy local_jobsimulation_sessions: ERROR: relation "public.local_jobsimulation_sessions"
does not exist (SQLSTATE 42P01)
```

→ the hiring org is under-set-dressed (5 positions, **0** candidate sessions against a want of ≥40)
→ **the one autoverify warning.** Host-independent; **no odysseus fix exists.** Fix = re-point the six
seeders in rext `stack-seeding`. It also **invalidates the "manager-view MIRROR trap" guidance** in
`corpus/ops/demo/content-stories-routes.md`.

**And the warning misattributes its own cause** — it blames a cold snapshot cache. The cache was
**warm** (2.98 GB staged, 330,261 taxonomy rows replayed), and its suggested remedy
(`stacksnap replay --surface directus`) cannot fix a seeder writing to a dropped table. Same
defect class as `verification.md:624`, in a different file.

### B2 — `app/studio` has no acquisition path in rext, so a cold `app` build cannot succeed

`app` gained the studio-room Python runtime in `fdb8034a` (*"feat(cms-in-app): M804 L3"*,
**2026-07-27 13:25:30Z**, first in tag `v1.360.0`), and its Dockerfile now hard-`COPY`s
`/build/studio`. rext handles `studio/` for **cms only** (`ensure-clones.sh:146` `make init-studio`,
`up-injected.sh:1654` `if [ "$svc" = cms ]`). `app` has **no `init-studio` target, no `.gitmodules` at
any ref**, and `.gitignore:78-79` says *"pulled at build via additional_repo"* — so **on any fresh box
`app/studio` is unobtainable** and the build dies:

```
#19 [stage-1 5/6] COPY --from=build /build/studio ./studio → "/build/studio": not found
```

**The timing is the finding.** billion's M255 baseline campaign ran **09:59–11:37Z on 2026-07-27** —
**~1 h 48 m before that commit landed.** So the cold bring-up path has been broken on **every** host
since, and it went unnoticed because **nobody has run a cold `--purge` + `demo-up` cycle since** (M256's
demo-2 was long-lived and never rebuilt). **M257 is the first milestone to actually exercise the path
its own gate is defined on.** Also: `ensure-clones.sh:144-145`'s premise *"studio is an embedded
pipeline, not a hard build dep"* is now **false for `app`** — it is a hard `COPY` and the build aborts.

The host was unblocked by hand (copying `cms/studio` → `app/studio`, a gitignored path — no tracked
file touched, byte-identical in shape to what rext's own sanctioned `make init-studio` produces). **That
hack is not reproducible by the tooling** and must become `ensure-clones.sh` + `up-injected.sh` coverage.

## ⚠️ And a measured gate risk: F3 is not theoretical — it is already biting

`TOK-01` reasoned that odysseus's **zero swap** turns a headroom overshoot into an OOM kill. Measured on
the first bring-up:

- free RAM hit **165 MB** during the 5 parallel Go backend builds; **217–235 MB** through the three
  Next.js lanes;
- **peak `load1` 48.7** — against HEADROOM **clause 1's limit of `cores − 2 = 6`**, and against
  billion's measured **4.06 / 4.56 / 4.22** for the same phase.

It survived (no OOM in `dmesg`, no ENOSPC — `df` and `docker system df` checked first per the M239-F1
rule; **142 G free**, build cache **27.78 GB** and now warm). But a `load1` of 48.7 is **8× clause 1's
limit**, and if it reproduces under `buildbench`'s sampler then **the gate's own HEADROOM clause cannot
pass on this host** — which is a re-scope conversation, not a lever problem. **iter-03's first
investigation**, before any campaign: is 48.7 real, reproducible, and attributable (memory-pressure
blocking with no swap to absorb it), or an artefact of how it was sampled?

## Close — 2026-07-31

**Outcome:** odysseus is a provisioned bench that reaches `rc=0`, and the gate's instrument is fixed in
both directions and **proven able to fail**. Metric delta: **none — zero by design** (no lever touched,
no gated number produced). Two blockers surfaced that make the gate currently unreachable, plus a
measured `load1` reading that may put its HEADROOM clause out of reach on this host.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(this was a tik; no 3-no-prog streak exists — one tik has run)* — (3) re-scope: n *(the trigger is "p50 > 420 s after L1+L2+L3"; no p50 exists yet. The `load1` finding is a **candidate** re-scope signal but is unverified — grading it as fired on one un-probed sample would be exactly the un-probed-lift dishonesty Phase 3 forbids, inverted)* — (4) user-blocker: n *(B1/B2 are rext-side, in-remit, need no authorisation → Fate-3 routed forward, which Phase 5 §4 lists explicitly as NOT a user-blocker)* — (5) cap-reached: n *(tik 1 of 5)* — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** see [`decisions.md`](decisions.md) (D1–D4)
**Side-deliverables:**
- **3 pre-existing `stack-verify` test failures fixed** — inherited, not planned. The M256 close's
  Python-suite roster omitted `stack-verify/tests` entirely, which is how they survived a close.
- The **swap figure corrected in 5 places**: billion has **16 GiB** (`billion.json:21` `swap_mib:
  16384`), not 15. `build-budget.md:122` had said 15 GiB **since M255**, disagreeing with its own
  machine-readable profile. Conclusion unchanged; the rule that *state the environment with every
  number* has no rounding exemption is the point.
- The recon table's `Go | absent` row corrected to agree with the F4 retraction below it.
**Routes carried forward** (all Fate 3, named handlers, all to **iter-03** as its precondition):
- `FIX-M257-seeders-local-mirror-drop` → re-point 6 `stack-seeding` seeders at
  `public.job_simulation_sessions` / `public.skill_path_sessions`; correct
  `content-stories-routes.md`'s MIRROR guidance; fix the under-set-dress warning's misattributed cause.
- `FIX-M257-app-studio-acquisition` → extend `ensure-clones.sh:146` + `up-injected.sh:1654` to cover
  `app`; correct `ensure-clones.sh:144-145`'s now-false premise.
- `INVESTIGATE-M257-load1-48` → is peak `load1` 48.7 real/reproducible/attributable, and does HEADROOM
  clause 1 survive on this host? **Blocks the campaign.**
- `FIX-M257-stacksnap-directus-sequences` → `stacksnap replay --surface directus` fails
  **deterministically**: `column "sequence_catalog" of relation "sequences" does not exist`. The string
  appears **nowhere** in rext source, so the server resolves it to `information_schema.sequences`, which
  the Directus content collection `directus.sequences` (relkind `r`) shadows by name. The identical SQL
  hand-run via psql **succeeds** — so the collision is in how the driver sends it, not the text.
  Non-blocking (content landed: `directus.simulations` = 307 rows).
- `DOC-M257-prereq-gaps` → **4 doc corrections** the provisioning found: `tailscale-serve.md:124`'s
  *"`/usr/local/go/bin` is added to `PATH` by the login profile"* is **false** (the Go tarball adds
  nothing) and **contradicts** `setup_guide.md:129`; **F2b needs a third branch** — *"installed, on no
  shell's PATH"* — because its stated logic reads that state as "Go genuinely missing" and its remedy
  would lay a **second** toolchain over a newer one (the reliable disproof is `ls /usr/local/go/bin/go`
  **before** installing); the prereq table needs a **Node ≥ 22** row; and the Go row should read
  **"≥ 1.25.12"** rather than "1.25.x".
- `FIX-M257-directus-coldstart-order` → compose starts directus **before** the per-stack provision
  creates its schema, so on a cold data dir it dies `Exited (1)` and nothing restarts it. Recovery is a
  non-destructive `docker start`. Now *detected* by check (h) — but detection is not a fix.
- `DOC-M257-autoverify-project-arg` → `autoverify.sh 1` prints *"ignoring unknown arg '1'"* and
  **silently skips**; it requires `--project demo-1`. Same shape as the `STACK_DIR` note already in
  `latency-budget.md`.
**Lessons:**
- **The gate was defined on a path nobody had walked.** B2 has been broken since 2026-07-27 13:25Z —
  1 h 48 m *after* the baseline campaign that priced every lever in this release, and 4 days before this
  iter. It survived because a long-lived demo-2 was never rebuilt. **A "cold bring-up" gate is worth
  exactly as much as the last time someone actually ran one cold**, and this milestone's first act
  should have been — and now was — to run it.
- **The instrument caught a real defect within hours of shipping.** Check (h) was justified on a
  *hypothetical* (an OOM kill masquerading as exit-0); what it actually caught first was an unrelated
  cold-start schema race. Instruments pay for themselves on the bugs you did not predict.
- **Three consecutive audit/brief claims were corrected by first-hand reading this iter** — the row-less
  gap (3 not 2), check (d)'s Linux behaviour (does not fire), and the swap figure (16 not 15 GiB).
  Combined with iter-01's four stale §8.5 anchors, the pattern is now unambiguous: **in this codebase, a
  claim's provenance matters as much as its content.** Verify before acting, every time.

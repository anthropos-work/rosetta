# iter-03 — tik: Phase C (cold bring-up on billion at the new tag)

**Type:** tik
**Shape:** standard tik under TOK-01, Phase C.

## Phase plan status

| # | Step | Status |
|---|---|---|
| 1 | Tear down stale `demo-1` (`--purge`) | ✅ DONE |
| 2 | Cold `demo-up 1`, public-host default-on | ✅ DONE |
| 3 | Verify stack health + fresh-green `autoverify.json` | ✅ DONE |
| 4 | Verify cockpit `/content-manifest.json` + tab presence | ✅ DONE |
| 5 | Triage a failed export (contingency) | n/a — export succeeded |
| 6 | Opportunistic first metric reading | ✅ DONE (substrate-level) |

## Step 1 — teardown

The stale stack was 17 containers built from the **M228-era tag** — it predates the feature under test.
`rosetta-demo down 1 --purge` removed all containers + the network, purged the container-owned data
(UID 1001 / 0700), and removed the stack's images. Note the correct entry point is
`demo-stack/rosetta-demo down <N> --purge`; `reap.sh` is a **sourced helper** (port-reaper), not a teardown
CLI — invoking it directly is a silent no-op (it exited 0 having done nothing).

## Step 2 — the cold bring-up, and the one real obstacle

**First launch FAILED the host pre-flight:**

```
up-injected: host pre-flight ✗ Go NOT on PATH — required: stacksecrets/stacksnap/stackseed are Go tools
                               that run ON the host …
up-injected: host pre-flight FAILED
```

**This was NOT a host defect and NOT a tooling defect.** Go **is** correctly installed on `billion` at
`/usr/local/go/bin/go`, version **go1.25.12** — exactly the `toolchain go1.25.12` the rext pins. The
failure is a **shell-invocation artifact**: `/usr/local/go/bin` is added to PATH by the login profile, and
a non-interactive `ssh host 'cmd'` does **not** source it. `atlas` happens to live in `/usr/local/bin`
(already on the default PATH) so it masked the pattern by passing.

**Fix: drive the bring-up through a login shell** — `ssh host 'bash -lc "…"'`. Relaunched that way and the
pre-flight passed. Recorded as D1 (with a corpus-doc route) because any agent driving `/demo-up` on a
remote host over non-interactive SSH will hit this and it looks exactly like a missing prereq.

The cold build then ran end-to-end: 6 injected Go images + Clerkenstein FAPI/BAPI + next-web + studio-desk
+ hiring, all from scratch (the iter-02 prune emptied the cache). The **M217 self-healing demopatch gate
worked as designed** — `app-aireadiness-snapshot-loadmembers` reported `whole-file sha DRIFTED … but the
anchor is intact (1x) → SELF-HEALING`, i.e. the anchor is the contract and the sha is only a baseline.
Both **M232 interview-flag demopatches** (`next-web-interview-flag-{container,result}`) applied to *both*
the web and hiring clones.

## Step 3 — health

```
✓ verify live: all liveness + readiness probes passed
✓ demo-patches: all applied (none refused, none skipped)
✓ frontend builds: ok (the running images are this run's)
✓ taxonomy replayed: public.skills = 42790
✓ directus.directus_collections = 21   ✓ directus DB is per-stack-local (not prod)
✓ presenter cockpit answering on :17700
✓ clerkenstein fake-FAPI answering on :15400 (hero login is possible)
▶ autoverify demo-1: OK — verified-working.
```

**Fresh-green `autoverify`. 17 containers. 0 platform-repo edits.**

## Step 4 — the content pipeline (the fail-closed trap, cleared)

```
content-export.log:  stackseed: exported content-stories manifest (18 sessions) → …/content-manifest.json
cockpit GET /content-manifest.json → HTTP 200
products: 4 | sessions: 18 | manager views: 15
```

The export exited 0, so `up-injected.sh` appended `--content-manifest` and the **2nd "Content stories" tab
is live**. The pre-flight's identified trap (a failed export silently yields a tab-less cockpit) did **not**
fire.

## Step 6 — opportunistic reading: the substrate is fully in place

Resolved all 13 simulation session ids from the **live** manifest's `player_result_path` and checked them
against the demo DB:

| Check | Result |
|---|---|
| present in `jobsimulation.sessions` | **13 / 13** |
| with `validation_attempt_results` rows (the M231 persisted-read contract) | **13 / 13** |
| present in `local_jobsimulation_sessions` (the M231 manager MIRROR trap) | **13 / 13** |
| `public.lab_sessions` (ai-labs presence target = 2) | **2** ✅ target met |

**Thread A gap confirmed, as scoped:** `academy_chapter_progresses` = **0**, `academy_skill_paths` = **0**,
`academy_chapters` = **0**. The academy catalog + progress fill is **not wired into the cold bring-up** —
which is precisely the M236 in-scope item ("wire `app/cmd/academy-seed …` into the cold bring-up"). Routed
to the academy arm, not a surprise.

## Metric

**0 / 31 — unchanged, and deliberately not claimed higher.**

The substrate reading above is strong evidence that up to 26 simulation pairs *can* land, but the primary
metric is **"renders real, non-empty content, live"** — and no render has been proven, because the
seat-login harness does not exist yet (Phase H). Per the protocol's honesty self-check, an un-probed
substrate count is **not** a claimable lift. Recording 26 here would be exactly the mis-attribution this
milestone was re-scoped to avoid.

## Close — 2026-07-20

**Outcome:** A cold, fully-rebuilt demo stack is live on `billion` at `playbill-m235-hardened` with a
fresh-green autoverify, the Content-stories tab serving all 4 products / 18 sessions, and the seeded
result fan-out verified present for 13/13 simulation sessions incl. the manager mirror. The v2.5 feature
is, for the first time, **running and measurable** on the host. Metric held honestly at 0/31.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (login-shell requirement for remote bring-up — corpus-routed), D2 (substrate present but not claimed as metric lift), D3 (academy fill absent — routed to the academy arm)
**Side-deliverables:** none (the teardown + build are planned scope).
**Routes carried forward:**
- **Phase H harness** — author the content-stories seat-login coverage plumbing against the now-LIVE seeded render (the calibration target USER-BLOCKER-M235-02 required). Now unblocked for the first time. → **iter-04**, handler `HARNESS-M236-iter04-seat-login`.
- **Academy fill** — wire `app/cmd/academy-seed` into the cold bring-up; academy tables are empty. → the academy arm, handler `ACADEMY-M236-iterTBD-catalog-fill`.
- **D1 corpus note** — the login-shell/PATH finding belongs in `tailscale-serve.md`'s F-series host-prereq set. → the B4/B5 doc pass, handler `DOC-M236-iterTBD-protocol-backfill`.
**Lessons:**
- **A host pre-flight that reads PATH is really testing the *shell invocation*, not the host.** `atlas` (in `/usr/local/bin`) passed while `go` (in `/usr/local/go/bin`) failed — the same box, the same provisioning, different PATH default. Any agent-driven remote bring-up should use a **login shell** by default; the failure otherwise mimics a missing prereq perfectly and invites a pointless reinstall.
- **Verify the entry point before trusting a silent success.** `reap.sh 1 --purge` exited 0 and printed nothing while tearing down nothing — a sourced helper invoked as a CLI. "Exit 0 + no output" on a destructive operation deserves an independent post-condition check (container count), which is what caught it.

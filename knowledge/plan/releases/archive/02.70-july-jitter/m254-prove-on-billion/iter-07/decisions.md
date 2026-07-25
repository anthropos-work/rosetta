# M254 · iter-07 — decisions

## D1 — captured the exact live failures on billion (10 fail + 1 error / 159)
Ran the host-sensitive files (`test_ant_academy`, `test_reap`, `test_host_isolation`,
`test_ant_academy_clerk_wiring`, `test_back_to_cockpit_m249`) via `python3 -m unittest` on billion (no pytest on
the box; tests are stdlib unittest). Unique failing tests (`TestAntAcademyPreBindReap` inherits
`TestAntAcademyLauncher`, so several appear twice): test_missing_node_documents (×2), test_second_launch_is_
idempotent_noop_while_running, test_stop_kills_the_recorded_pid, test_a_stale_academy_on_our_port_is_reaped,
test_MUTATION_without_the_reap_block_the_stale_academy_wins, test_apply_revert_round_trip_on_the_real_next_config,
test_mutant_no_term_trap_is_caught, + ERROR test_overlay_has_minted_pk_and_no_real_secret.

## D2 — FIXED: nvm/node host-robustness (test_missing_node_documents ×2) — rung-zero
`_run(no_node=True)` built its "node absent" PATH as `npm_only + /usr/bin + /bin`, ASSUMING node lives outside
those dirs. On billion `/usr/bin/node` is v18 AND `~/.nvm` holds v22, so `command -v node` resolved a node and
the M219 nvm rescue found v22 — the "node absent" path never ran. Fix: build a truly node-free bindir (symlink
every `/usr/bin`+`/bin` entry EXCEPT node/nodejs) + point HOME at a clean no-`.nvm` dir. Verified live on billion
(2 tests green). rext `dfdd9bc`, tag `july-jitter-m254-academy-nonode-hostrobust` on origin.

## D3 — ROUTED (a dedicated test-health batch, `FIX-M254-g-testhealth`): the remaining 6, root-caused
- **test_second_launch / test_stop_kills / test_a_stale_academy_reaped** — **intra-run listener leakage +
  M245 reconcile drift.** Two entangled causes: (a) tests spawn detached stub `python -m http.server` on the
  shared demo-2 port `:23077` (via `launch_detached`/setsid), and tearDown only kills the recorded pidfile pid
  — orphaned servers survive, so a later test finds `:23077` "held by a process we do NOT own" (reap refuses to
  kill it) or hits `OSError: Address already in use`. (b) The M245 academy-durable-fix added the reconcile+
  `render_ok` branch: a 2nd launch of a NON-rendering academy now logs "reconciled … relaunching to recover"
  (not the old "already running") — the stub `SimpleHTTPRequestHandler` doesn't serve a `/library/` with
  `href="/courses/…"`, so `render_ok` fails. Fix surface: (i) tearDown reaps ALL listeners on the test's port
  (or use a unique free port per test); (ii) the stub npm `run dev` serves a minimal rendering `/library/`;
  (iii) update the assertions to the M245 reconcile messages. Handler: `FIX-M254-g-academy-launch-isolation`.
- **test_apply_revert_round_trip_on_the_real_next_config** — **next.config.js sha drift.** The demo clone's
  `next.config.js` hashes to `0d58ea60…` but the dev-origins patch manifest's `pre_sha256` is `6837cab9…`
  (the upstream file moved, likely post-M246 consolidation). Fix surface: re-pin the manifest pre/post shas to
  the current pristine `next.config.js` (the aireadiness-repoint pattern, iter-03). Handler:
  `FIX-M254-g-devorigins-sha-repin`.
- **test_mutant_no_term_trap_is_caught** (host_isolation) + **test_MUTATION_without_the_reap_block** —
  **mutation meta-tests** that deliberately break a guard and assert the break is caught. On billion the mutated
  behaviour didn't manifest (`AssertionError not raised` / the stale academy didn't survive), i.e. the meta-test's
  environmental precondition differs on the VM. Fix surface: make the mutation's precondition host-deterministic
  (pin the port/proc it mutates). Handler: `FIX-M254-g-mutation-meta`.
- **ERROR test_overlay_has_minted_pk_and_no_real_secret** — the inlined `write_env_local` bash `-c` script
  exited **127** (command not found) on billion. A command the extracted snippet calls isn't on the test's PATH.
  Fix surface: ensure the extracted `write_env_local` snippet's PATH carries its deps (or run it under the real
  ant-academy.sh env). Handler: `FIX-M254-g-overlay-127`.

## D4 — (g) disposition: 1 fixed / 6 routed → coordinator fate
These are the chronic "host-sensitive" demo-stack tests (the gate's "8", host-sensitive membership) that have
historically slid their expiry to release close and been fated there by the user. iter-07 lands the one clear,
correct, verified host-robustness fix and precisely root-causes + routes the rest as a dedicated test-health
batch (`FIX-M254-g-testhealth`). They are test-harness / fixture / manifest issues (0 platform edits, 0 demo
runtime impact — the real academy on `:13077` serves 200). Surfaced for coordinator disposition; not
demo-proof-blocking. The core "prove on billion" demo-proof gates (e / h / c-academy) take budget priority.

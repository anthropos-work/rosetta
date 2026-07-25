# M254 — carry-forward

Items routed forward from M254's iters, for the close-milestone deferral audit to formally fate (three-fate
rule). None block the M254 exit gate — all gate parts (a–h) are MET (6 hard-MET + e/h via the live
Playthrough suite + the 3 coordinator-approved dispositions below).

## Coordinator-approved dispositions (recorded iter-10, fated at close)

- **`(f)-FCP-p95` — ACCEPTED environmental** (D-iter10-2). Studio first-paint shell holds on billion
  (p50 637–726 ms < 1 s, M253 fix holds); p95 outliers (1443/2014/4943 ms, reachedShell always true) are
  tailnet network-RTT jitter, per `latency-budget.md`'s "state the environment" + the (b) precedent. Gate (f)
  MET on app-side paint. **Fate 1** (disposition recorded).

- **`FIX-M254-c-academy-durable` — (c)-academy-durability follow-up** (D-iter10-3). A FRESH demo renders
  Back-to-Cockpit on all 4 apps (gate (c) MET, presenter path). The native academy dev-server reverts the
  patch only on a long-running demo. Route: make the academy back-to-cockpit patch durable across the native
  dev-server's lifecycle (rext `stack-injection`/`ant-academy.sh` reapply lifecycle, 0 platform edits).
  **Fate 3** (academy-durable follow-up).

- **`FIX-M254-g-testhealth` — 6 host-sensitive demo-stack test-health tests** (D-iter10-4). 2/8-class fixed +
  verified live; the 6 remaining are chronic host-sensitive test-HARNESS issues (intra-run `:23077` port-leak
  + M245 reconcile-message drift [`test_second_launch`/`stop_kills`/`a_stale`]; next.config sha re-pin
  [`test_apply_revert`]; 2 mutation-meta; overlay exit-127) with **0 demo-runtime impact** (real academy
  serves 200). Route: per-test disentangle + fix in rext `demo-stack` (a fix-iter's work), 0 platform edits.
  **Fate 3** (carry-forward).

## Harden-surfaced (final harden pass) — LANDED at close

- **`FIX-M254-h-patch-inventory-drift` — the demopatch inventory fence RED at HEAD — ✅ LANDED (Fate 1) at
  the M254 close (2026-07-25).** Surfaced by the final harden pass: `demo-stack/tests/test_patch_inventory.py`
  asserted `EXPECTED_TOTAL=21` / `EXPECTED_BY_REPO["studio-desk"]=3`, but `patches/` holds **23** manifests
  (studio-desk **5**), so 2 tests failed on any host. **Root cause = M253** (`b8969c0`) added
  `studio-desk-shell-first-paint` + `studio-desk-no-thirdparty` without bumping the fence — RED since the M253
  tip; **NOT an M254 change** (M254 only changed the aireadiness patch's content, not the patch count). A
  RED-at-HEAD must not ship, so the close re-fated this from the harden pass's Fate-3 to **Fate-1 LAND-NOW**
  and landed it in full:
  - **rext:** `EXPECTED_TOTAL 21→23`, `EXPECTED_BY_REPO["studio-desk"] 3→5`, history comment updated →
    `test_patch_inventory` 5/5 green; commit `02ac973`, tag **`july-jitter-m254-patch-inventory-fence`** ON
    ORIGIN (rung-zero verified).
  - **corpus:** `corpus/ops/demo/demopatch-spec.md §5` reconciled — the §5 header was already 23 (M253), but
    three secondary fence-description references still read 21/3/16: line 144 "three → **five** studio-desk",
    the fence-contract line "total 21→**23** / studio-desk 3→**5**", and the self-contradicting "current
    directory-fenced total is 16 → **23**".
  - **verification:** full demo-stack suite 909 tests, 900 pass; the 8+1 remaining are all the
    `FIX-M254-g-testhealth` host-sensitive carry (0 milestone regressions). Test-only; no demo re-pin; 0
    platform-repo edits.

## Prior-iter follow-ups (non-blocking)

- **`(b)-voice manager_presence_only`** — the content-stories denominator: gate (b) = 45/45 landable ALL
  LANDED + 4 voice presence-only (2 player + 2 manager, Bunny-keyless demo box, symmetric extension of
  `DEF-M240-01`). Follow-up: a rext `manager_presence_only` flag + content-denominator 47→45 + re-seed
  (tracked, non-blocking; the surface renders, no fabricated CTA). **Fate 3.**

- **verify.sh stale `skillpath` default** (iter-04 note) — verify.sh's DEFAULT service list (when `--services`
  omitted) still lists the decommissioned `skillpath`, so a standalone autoverify without `--services` probes
  a non-existent service (HTTP 000000). Not gate-blocking (bring-up passes explicit `--services`). Route: drop
  `skillpath` from verify.sh default. **Fate 3.**

- **studio-desk billion re-pin → `july-jitter-m254-studio-pt-retune` (4f1409e)** on the NEXT full cold
  re-prove. The iter-10 studio-builder + networkidle fix is **test-only** (runs from the local rext clone;
  billion serves the unchanged app), so it needs **no** demo re-pin for the current proof — but a future full
  cold reset-to-seed should pin billion's rext to 4f1409e so the whole toolchain is current. **Fate 3** (note).

## Observation (not a defect)

- **studio membership-check toast** — the studio-desk client-side org-membership poll to the fake-BAPI shows
  "Checking user organization memberships failed after 3 attempts. Please try again." over the tailnet, but it
  is **non-blocking** (both builders render fully behind it; auth works; generation works). Likely a
  tailnet-latency retry timeout. Noted for a future demo-wiring look; not gate-blocking.

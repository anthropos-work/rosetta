---
milestone: M18
slug: verify-net
version: v1.3b "dress rehearsal"
milestone_shape: section
status: planned
created: 2026-06-08
last_updated: 2026-06-08
complexity: medium-large
delivers: corpus/ops/verification.md (net-new) + offset-/scope-aware stack-verify + bring-up wiring
issues: ISSUE-12, ISSUE-14
---

# M18 — The verification safety net

## Goal
Teach `stack-verify` to target an **offset** stack and scope to the services actually brought up, then **auto-run it
(non-fatal) at the tail of every bring-up** — so "UP" means *verified-working*, not just *containers-started*.

## Why section
The gap is fully diagnosed (ISSUE-12/14, verified against code): the hardcoded service table, the missing offset
awareness, the `$DEVDIR` bug, and the absent auto-run are all known. Concrete edits + a wiring hook.

## Repo split
- **`rosetta-extensions`** (authoring → tag `dress-rehearsal-m18` → consume): the `stack-verify` changes + the
  bring-up-tail wiring.
- **`rosetta`**: the net-new `corpus/ops/verification.md`.

## Scope
- **In (`rosetta-extensions` code; one `rosetta` doc):**
  - **Project/offset awareness** — derive the `demo-N`/`dev-N` prefix + the N×10000 port offset from the unified
    registry's recorded host ports (today `lib/services.sh:25-39` hardcodes 12 `anthropos-*-1` names at **base**
    ports → reports an offset stack entirely `down`). (ISSUE-12b.)
  - **Service/profile filter** — intersect the checked set with what was requested (a reduced bring-up isn't a wall
    of false `down`s). The bring-up CLIs already parse `--services/--profile`; pass it through. (ISSUE-12b.)
  - **Fix the undefined `$DEVDIR` → `$STACK_ROOT` bug** (`repos/run.sh:108`, `census/inventory.sh:75`). (ISSUE-12.)
  - **Cheap-win asserts available today** — `curl -fsS .../api/health` + `SELECT count(*) FROM
    sentinel.casbin_rules > 0` on the stack's offset ports at the bring-up tail (the exact ISSUE-7 silent-failure
    catcher). (ISSUE-14.)
  - **Auto-wired scoped `verify live`** at the bring-up tail — **default-on + non-fatal** (mirror `dev-setdress`'s
    proven default-on + non-fatal pattern). (ISSUE-12c/ISSUE-14.)
- **Out:** verifying the *frontend* ports (added in M19, where the frontends first exist); deep behavioural/e2e
  probes (that's `/test-platform` — M18 is the always-on smoke net).

## Depends on
M16. **Parallel with:** M17 (yes-with-caveats — different surfaces; lean sequential).

## Open questions (resolve during build)
- Offset from `STACK_ROOT`-parse vs reading the registry's recorded host ports — lean: read the registry (M12
  already records resolved ports per stack).
- How loud a "non-fatal but failed" should be — lean: a clear ⚠️ block + a one-line "run `/test-platform N` to dig in".

## KB dependencies (read as contract)
- `corpus/ops/run_guide.md` (the "verify after start" intent), `corpus/ops/rosetta_demo.md` (the registry + ports)
- the `stack-verify` extension section, the `/test-platform` skill

## Delivers
- **→ rosetta:** `corpus/ops/verification.md` — net-new: the auto-verify contract (default-on + non-fatal) + the
  offset/scope model + the cheap-win asserts.
- **→ rosetta-extensions:** the offset-/scope-aware `stack-verify` (services.sh + verify.sh + readiness.sh) + the
  `$DEVDIR` fix + the bring-up-tail wiring.

## Risk
**(correctness, degrades-quality)** a mis-derived offset false-positives "down" — the very bug it fixes. Mitigate:
derive from the registry's recorded ports (not re-computed); **non-fatal** so a verify bug never blocks a genuinely
good stack; a regression test on an offset `demo-N` fixture.

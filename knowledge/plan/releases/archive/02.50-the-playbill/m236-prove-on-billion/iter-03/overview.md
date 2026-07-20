---
iter: 3
milestone: M236
iteration_type: tik
status: closed-fixed
date: 2026-07-20
metric_pre: "0/31"
metric_post: "0/31 (unchanged — substrate verified, render not yet proven)"
---

# iter-03 — tik: Phase C, cold bring-up on billion at the new tag

## Step 0 — re-survey (mandatory)

Re-ran the primary measurement's precondition before committing to TOK-01's named target. State **has**
changed since TOK-01 was authored — iter-02 landed Phase P — so the re-survey confirms Phase C is now the
correct next target rather than a stale one:

- `billion` pin SoT + consumption clone both at `playbill-m235-hardened`; M217 guard PASSES.
- Host has 139 G free (post-prune) — enough for a cold UI-tier rebuild.
- A **stale `demo-1` stack is still running** (17 containers, built from the M228-era tag). It predates the
  feature under test and **must be torn down**, not reused — reusing it would prove nothing about the v2.5
  tooling and would silently measure M228 behaviour.
- Metric unchanged at **0/31** — nothing is measurable until a stack is up at the new tag.

TOK-01's Phase C is **untouched and still meaningful**. No substitution.

## Active strategy reference

**TOK-01 "publish-then-prove", Phase C.** This tik is Phase C in its entirety.

## Phase 0d — pre-flight tooling check: PASS

Phase C wires artifacts (the seed → `--content-export` projection → cockpit) through a multi-stage
pipeline, so the pre-flight applies. Verified on the host at the pinned tag, before the build:

| Check | Result |
|---|---|
| `stack-seeding/contentsession/` + `presets/content-manifest.json` present | ✅ |
| `ContentStorySeeder` registered (`cmd/stackseed/main.go:513`) | ✅ |
| `--content-export` flag exists (`main.go:158`) | ✅ |
| `cockpit.py` serves `/content-manifest.json` + renders the 2nd tab (`:729`) | ✅ |
| `up-injected.sh` wires `--content-manifest` into the cockpit launch (`:1958-1984`, `:2024`) | ✅ |

**Pre-flight finding worth carrying:** the wiring is **fail-closed by design** — `up-injected.sh:1982`
only appends `--content-manifest` *if* `stackseed --content-export` exits 0; otherwise
`content_manifest_arg` stays empty and the cockpit comes up **with no Content-stories tab at all** and no
loud error at the cockpit layer. The export's stderr goes to `$STACK/content-export.log`. So *"the tab is
missing"* and *"the export failed"* are the same observable — **read `content-export.log` first** on any
missing-tab symptom, before suspecting the cockpit.

## Cluster / target identified

The reachability gap is closed but **unexercised**. Nothing in the v2.5 chain — seeder → export → manifest
→ cockpit tab → result routes — has ever run on this host. Phase C is the single step that converts
"published" into "running", and every Phase-L arm is downstream of it.

## Hypothesis

A cold `demo-down 1 --purge` + `demo-up 1` at `playbill-m235-hardened`, with public-host default-on, brings
up a stack whose cockpit serves a non-404 `/content-manifest.json` carrying all 4 products — making the 31
landable pairs **measurable for the first time**.

## Expected lift

**Primary metric: 0 on the numerator is an acceptable and expected outcome.** Phase C's deliverable is a
running, verified stack — not landed pairs. Any numerator movement observed here is a bonus reading, not
the iter's grade.

The *verifiable* outcomes are binary:
1. Stale `demo-1` fully torn down (no M228-era containers survive).
2. Stack UP at the new tag; `autoverify.json` **fresh green**.
3. Cockpit serves a **non-404** `/content-manifest.json` with **4 products / 18 sessions**.
4. The Content-stories tab renders (i.e. `--content-manifest` was actually wired, per the fail-closed trap).
5. 0 platform-repo edits.

## Phase plan

1. Tear down the stale `demo-1` (`--purge`).
2. Cold `demo-up 1`, public-host default-on. **Expect 30–45 min** — the build cache was pruned in iter-02,
   so every image rebuilds from scratch. Heartbeat throughout.
3. Verify stack health + a fresh-green `autoverify.json`.
4. Verify the cockpit's `/content-manifest.json` (4 products, 18 sessions) and the tab's presence.
5. If the export failed, read `content-export.log` FIRST (the fail-closed trap above) and triage from there.
6. Take an opportunistic first reading of the primary metric if the stack is healthy.

## Escalation conditions

- **Bring-up fails on a host-capacity or infrastructure cause** → triage once; if it is not a v2.5-tooling
  defect, route the host item forward rather than turning this iter into a host-ops project.
- **`--content-export` fails** → this IS a v2.5-tooling defect and is in-scope: triage it here.
- Anything requiring a **platform-repo edit** → hard stop (the 0-platform-edits constraint).
- Scope-creep tripwire: this iter is bring-up + verification. Landing *pairs* is Phase L (iter-04+) — a
  green stack with 0/31 landed still closes `closed-fixed`.

## Acceptable close-no-lift outcomes

If the stack comes up but the content pipeline is broken for a characterized reason, the falsification
("stack green at the new tag; export fails because X") is a complete iter outcome and closes
`closed-no-lift` with the defect routed to iter-04.

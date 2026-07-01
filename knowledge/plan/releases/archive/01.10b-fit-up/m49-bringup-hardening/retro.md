# M49 — Bring-up hardening + truth-up — retro

## Summary
Closed the 7 remaining demo-up issues so a from-cold `/demo-up` on a `stack-demo`-only box reaches set-dress + seed +
verify + cockpit cleanly — and made the bring-up docs tell the truth. The headline mechanics: #1 a single
`.agentspace/rext.tag` consumption-tag source-of-truth (+ a CRLF-tolerant `lib/rext_tag.sh` reader) retiring 4
conflicting prose pins; #3 the `.env`-guard reordered to provision-then-check (the `stack-demo`-only abort); #4
`INVITATION_HMAC_SECRET` as a critical secret-DNA gene + values-blind auto-gen + a `DemoGeneratedKeys` overlay
(fixing the silent `app Exited (0)` class); #5 ant-academy cloned **explicitly** in `ensure-clones.sh` (NOT
`repos.yml`); #6 a disk-headroom pre-flight + `demo-down --purge` per-demo image cleanup; #7 a *true* non-fatal
frontend (absent image → `--scale svc=0`); #8 the demopatch re-anchored to next-web v2.89.0. Two-repo split: rext
code of record @ tag `fit-up-m49` (`ba586d6`), consumed per-stack; rosetta carried the `corpus/ops` truth-up. The
**live-verify gate PASSED** (a from-cold `/demo-up` on a re-pinned `fit-up-m49` clone proved all 7 fixes end-to-end:
demo-1 UP, autoverify "verified-working") — that is the M49 acceptance.

## Incidents this cycle
- **P2 (caught + fixed in harden, not shipped to the live gate) — `rext_tag.sh` CRLF carriage-return leak.** The
  reader's `awk` captured a trailing `\r` into the picked token, so a CRLF-edited `.agentspace/rext.tag` (a Windows
  editor / any `\r\n`-writing tool) yielded `fit-up-m49\r`. The leaked CR would fail `git checkout <ref>` in
  `ensure-clones.sh` with a baffling "pathspec did not match" — an invisible CR in the ref. Fixed with
  `gsub(/\r/,"",$1)` (portable across GNU/BSD awk, unlike a BSD-skipped `sed s/\r//`); regression test
  negative-control-verified (fails on the pre-fix reader). Knowledge-backfilled into `rosetta_demo.md` (the pin
  reader is CRLF-tolerant). No flakes.

## What went well
- **The two-repo split held cleanly.** rext code froze at `fit-up-m49` and was consumed per-stack; rosetta carried
  only docs + records. close-milestone reviewed BOTH repos (Phase 2/4 in the rext authoring copy) and merged only the
  rosetta doc half — exactly the tooling-policy shape.
- **Live-verify-as-acceptance was the right call for a bring-up grab-bag.** A from-cold run proved all 7 fixes
  compose at runtime (which static suites can't), so close ran the STATIC suites as defense-in-depth without burning
  the single demo machine on a re-run.
- **The #5 root-cause correction stuck.** The real cause was "`make init` skips what `repos.yml` omits", not the old
  FA-token theory — and the fix (an explicit `ensure-clones.sh` clone, the cms/studio submodule precedent) is durable
  + touches no platform file.

## What didn't
- **A forward-reference went stale across the M48→M49 handoff.** M48's docs (and its retro carry-forward) predicted
  "M49 adds the `repos.yml` entry". M49 correctly pivoted to an explicit clone (repos.yml lives in the ephemeral
  gitignored platform clone — editing it is non-durable + a platform edit), which left the M48 prediction inaccurate.
  M49 §4 + the close caught most of it, but **two residual "cloned via `make init`" claims survived in the un-touched
  `corpus/architecture/service_taxonomy.md`** (lines 191, 383) and were only caught at close-Phase-3. Lesson: when a
  milestone's actual fix diverges from a prior milestone's forward-ref, sweep ALL docs that quote the predicted fact,
  not just the milestone's own delivers-list.
- **Test-count literals drifted in the records.** `progress.md` pinned "163 demo-stack tests" (a non-monotonic
  build-phase cumulative tally) where the runner actually collects 111; demopatch said 46 vs the actual 47. Caught +
  reconciled at close-Phase-4 per the handbook-count contract. Lesson: pin counts from the runner, not from prose
  running-totals.

## Carried forward
- **AI-provider-keys policy → M50** (Fate-2, audit-confirmed). Which of OPENAI/ANTHROPIC/MISTRAL/ELEVENLABS/LIVEKIT
  become throwaway/sandbox demo values vs documented-as-absent; gates the academy AI chat. M50 owns it (overview).
- **Consumption-clone re-pin → push-gated KEEP** (release-level). The `.agentspace/rext.tag` bump to a pushed
  `fit-up-m49` is tracked with the release's other pending origin pushes; authoritative bump is M53's cold-rebuild
  acceptance step.

## Metrics delta
- Tests: rext Go 1552 → **1555** (stack-secrets 160→163: the #4 gene asserts + the `TestMintedShape_DemoGeneratedKeyIsOpaque`
  harden probe); demo-stack Python **299** (demopatch 46→47). Coverage: `demo.go` `Shape` 88.9%→100%. Flake **0**
  (5/5). 0 new deps (rext unchanged supply-chain; rosetta docs-only). Close review: 6 findings, all Fate-1; deferral
  audit GREEN. (Full: `metrics.json`.)

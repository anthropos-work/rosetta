---
active_release: "v1.10b fit-up — COMPLETE (6/6+F6 GREEN from cold); ready for /developer-kit:close-release. v2.0 opening night PAUSED"
active_branch: "release/01.10b-fit-up"
active_milestone: "(between milestones — v1.10b's final milestone M53 closed; release ready to close)"
last_closed: "M53 — 2026-07-01 (Cold-rebuild acceptance, section)"
phase: "v1.10b COMPLETE — all 7 milestones (M47..M53) closed to release/01.10b-fit-up; next: /developer-kit:close-release"
last_updated: "2026-07-01"
---

# State

**Active release:** **v1.10b "fit-up" — COMPLETE, ready to close.** An **interposed field-hardening backfill** (the
v1.3b "dress rehearsal" lineage): a from-scratch `/demo-up` surfaced 8 bring-up issues + a tail of v1.10 content gaps.
7 milestones **M47 → { M48 ∥ M49 } → M50 → M51 → M52 → M53**, all closed to `release/01.10b-fit-up`. The clones were
found **current** (M47 — a trivial `make pull`, not the reported 5-week lag); the genuinely-stale surface was the
**rosetta corpus** (re-grounded in M48). The release: snapshot recaptured from current prod (M47), corpus re-grounded
(M48), the 8 bring-up issues + v1.10 content gaps fixed (M49/M50), a curated **AI-readiness showcase org** added (M51),
**one auditable seed+gen manifest** consolidated (M52), and **cold-rebuild acceptance** proven (M53). **Tooling + docs
only — zero platform-repo edits.** rext code of record @ tag **`v1.10.1`** (`576dbcb`).

```
M47 ──→ ┌ M48 corpus re-ground ───────────┐                (M48 ∥ M49 — disjoint clusters; M48 no-demo)
        └ M49 bring-up hardening ──────────┘ ──→ M50 ──→ M51 ──→ M52 ──→ M53 ✅
```

**Active milestone:** **(between milestones).** v1.10b's final milestone **M53 — Cold-rebuild acceptance** closed
2026-07-01 (see Recently closed). The release is complete; there is no next milestone to build.

**Phase:** **v1.10b building COMPLETE — M47..M53 all CLOSED.** The **single-demo serialization** (fix-on-live across
M47..M52, then M53 destroyed + cold-rebuilt as the single acceptance proof) ran to completion: demo-1 was purged +
cold-rebuilt from the `v1.10.1` tag by a single `/demo-up`; the acceptance bar asserted **6/6 + academy F6 GREEN from
cold**. (AB4 surfaced an M51-owned gate regression — an unconditional ai-readiness manager seedPath — fixed at the gate
[user-approved; org-conditional manager manifest, rext `117fe41`], and `v1.10.1` re-rolled to include it + the close
harden tests → `576dbcb`.)

**Next up:** **`/developer-kit:close-release`** — the release-level review + merge of `release/01.10b-fit-up` → `main`
+ tag `v1.10.1`. _(The orchestrator still owes origin the pushes: `main` + the `v1.10` tag + the v1.10 ext tags + the
`fit-up-m47..m52` rext tags + `v1.10.1` — the v1.10 LOCAL close did not push; the M201 close merged to `main` LOCALLY;
this v1.10b branch is cut from that local `main`. The consumption-clone re-pin to `v1.10.1` is DONE [box-level,
authoritative at M53]; the origin push is the sole push-gated KEEP.)_

**Last shipped:** **v1.10 — 2026-06-27** (`method acting`, 9 milestones M39→M46, tag `v1.10`,
`release/01.10-method-acting` merged `--no-ff` → `main`). The **last release of the v1.x major**; its history + the
full shipped log now live in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
[`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).

**Paused:** **v2.0 "opening night" (Playthroughs)** — paused 2026-06-29 after M201 closed, to interpose this **v1.10b
"fit-up"** backfill. M201 corpus preserved as the v2.0 spec; M202 ∥ M203 ∥ M204 not started. Resume after v1.10b ships.

**Standing backlog (unscheduled, cross-release):** DEF-M10-01 (cloud SnapshotStore / S3 blob bytes), DEF-M21-01
(`replayCmd` hermetic test), M25-D9 (dev taxonomy `rc=4`); **future v2 milestones** M205 Hiring + tier gates · M206
AI-sim mirror tier · M207 Academy coverage. All tracked in [`roadmap-vision.md`](roadmap-vision.md); none scheduled.

## Recently closed (v1.10b milestones)
- **M53 — Cold-rebuild acceptance** — **2026-07-01** (`section`; the FINAL v1.10b milestone; merged →
  `release/01.10b-fit-up`; rext release tag `v1.10.1` @ `576dbcb`). demo-1 purged + cold-rebuilt from `v1.10.1` by a
  single `/demo-up`; **6/6 acceptance criteria + academy F6 GREEN from cold.** AB4 surfaced + fixed an M51-owned gate
  regression at the gate (org-conditional manager manifest, `117fe41`). Close: 2 Fate-1 doc fixes; deferral audit
  **GREEN** (every carry landed here; academy-F6 REPEAT resolved by execution). Tests rext stack-seeding 786→791 ·
  Python 313→326 · TS 29; flake 0. Full narrative in `roadmap.md` § M53.
- **M52 — Single auditable seed+gen manifest** — **2026-07-01** (`section`; rext `fit-up-m52` @ `36d7430`). ONE
  checked-in `seed-generation-manifest.yaml` drives all seed+gen intent, projected + honesty-gated from the presets,
  served by the cockpit [Download]. 12 close findings all Fate-1 (F1 dedup projection; F4 warn on orphan gen-id). Tests
  rext stack-seeding 749→786 · Python 313; flake 0. Full narrative in `roadmap.md` § M52.
- **M51 — AI-readiness showcase org** — **2026-07-01** (`iterative`, closed-on-gate; rext `fit-up-m51` @ `a23f38d`).
  Manager coverage gate MET at iter-09 (Northwind 200, 78.4% all-3-complete, closed cycle + 199 frozen snapshots). 3
  net-new seeders + the `app-aireadiness-snapshot-loadmembers` read-path demo-patch. Deferral audit RED→CLEARED (academy
  F6 LAND-NEXT → M53). Full narrative in `roadmap.md` § M51.
- **M50 — Content & seeding fill** — **2026-06-30** (`iterative`, closed-on-gate; rext `fit-up-m50` @ `f0d984c`). M42
  coverage gate MET both vantages; new member-language/certificate/user-field seeders + Directus content-URL rewrite.
  Full narrative in `roadmap.md` § M50.
_(M49/M48/M47 closed 2026-06-30/29 — full narratives in `roadmap.md` §§ M47–M49.)_

## Recently shipped releases
- **v1.10 "method acting"** — **2026-06-27**, tag `v1.10`. The believable-profile release + presenter-grade /
  scalable-generation extension; Playwright SEMANTIC coverage gate at both vantages cold; 9 milestones. The **last v1.x
  release** — detail in [`roadmap-legacy.md`](roadmap-legacy.md). Records:
  [`releases/archive/01.10-method-acting/`](releases/archive/01.10-method-acting/).
- **v1.9 "storytelling"** — **2026-06-23**, tag `v1.9`. Declarative Stories & Heroes seeding + presenter cockpit.
  Records: [`releases/archive/01.90-storytelling/`](releases/archive/01.90-storytelling/).
- **v1.8 "understudy"** — **2026-06-15**, tag `v1.8`. Self-contained-demo release. Single milestone M26.
- **Earlier v1.x** (v1.0 … v1.7) — full shipped table in [`roadmap-legacy.md`](roadmap-legacy.md) § Shipped releases.

## Headline numbers (v1.10b — M53 close)
- **Go test funcs (rext):** M53-touched module **stack-seeding = 791** (`Test`+`Fuzz`; +5 vs M52's 786 — F6 academy
  DeepLink/AcademyDeepLink build + harden single-source tests). Other modules (unchanged this milestone): `alignment`
  52 · clerkenstein 270 · stack-snapshot 364 · stack-secrets 163. All Go modules green.
- **Python / TS:** `demo-stack` Python **326** (+13 vs M52's 313 — F6 authenticated-session + [Academy] deep-link +
  harden `_academy_catalog_entry` edge/escape tests). rext **e2e TS unit** **29** (AB4 org-gating + referential-stability
  edges, +2 vs the pre-AB4 27). `stack-injection` 117.
- **Flake:** **0** (M53 close flake gate 5/5 Go seeders `-shuffle=on` + 5/5 Python cockpit+academy [101 each] + 5/5 TS
  coverage-manifest [29 each]).
- **Supply-chain:** the v1.10 close carried 1 new dep (`github.com/anthropos-work/ai@v1.40.1`, M45). v1.10b added none.
  The rosetta corpus is docs-only. Lockfile inherited from
  [`releases/archive/01.10-method-acting/dependencies.lock`](releases/archive/01.10-method-acting/dependencies.lock).
- **Alignment gates:** **100%/100%** on all 5 Clerkenstein surfaces (v1.10 close; v1.10b touched no contract surface).

## Branch model
**v1.10b COMPLETE (active, ready to close):** `release/01.10b-fit-up` cut from `main` 2026-06-29 (LOCAL — origin push
is the orchestrator's step). All 7 milestone branches `m{47..53}/{slug}` merged `--no-ff` + deleted. rext code of record
(a SEPARATE repo) authored in `.agentspace/rosetta-extensions/`, tagged `fit-up-m47..m52` per milestone + rolled to the
**`v1.10.1`** release tag at M53 (`576dbcb`, re-rolled at close to include the harden tests); consumed via the
`.agentspace/rext.tag` source-of-truth (= `v1.10.1`) + the `stack-demo/rosetta-extensions` consumption clone (pinned to
`v1.10.1`).
**v2.0 PAUSED:** `release/02.00-opening-night` cut from `main` 2026-06-28 (LOCAL). M201 merged → `main` (LOCAL, no `v2.0`
tag); M202→M204 not started — resumes after v1.10b. A `playthroughs` rext section arrives at M202 build.
**v1.10 SHIPPED:** `release/01.10-method-acting` merged `--no-ff` → `main` + tagged `v1.10` at close (LOCAL). rext code
@ tags `method-acting-m39..m46-servegrant-closure`.
**Shipped:** **v1.10** `v1.10` · **v1.9** `v1.9` · **v1.8** `v1.8` · **v1.7** `v1.7` · **v1.6** `v1.6` · **v1.5** `v1.5`
· **v1.3b** `v1.3.1` · **v1.3** `v1.3` · **v1.2** `v1.2` · **v1.1** `v1.1` · **v1.0** `v1.0`.
(Full shipped detail: [`roadmap-legacy.md`](roadmap-legacy.md).)

_Last updated: 2026-07-01 (M53 "Cold-rebuild acceptance" CLOSED — the FINAL v1.10b milestone; 6/6+F6 GREEN from cold;
v1.10.1 re-rolled to `576dbcb`. Release complete → next: `/developer-kit:close-release`.)_

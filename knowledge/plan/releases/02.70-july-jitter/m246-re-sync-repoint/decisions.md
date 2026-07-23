# M246 — Decisions

_(Implementation decisions with rationale, D-numbered, recorded during build.)_

## KB-fidelity findings (Phase 0b, 2026-07-23) — verdict YELLOW

- **KB-1** — Seeder write-site anchors ALIGNED. `main.go:97` + `hero_activity.go:180` verified exact; full write-site inventory in spec-notes. In-scope build work.
- **KB-2** — `DEMO_ADVANCE_CLONES=pinned` path ALREADY WIRED (M237, `ensure-clones.sh:206`). M246's delta = author `stack-demo/clones.pin.json`, not wire the path. Scope clarification.
- **KB-3** — Injection comment `gen_injected_override.py:16` STALE (4→3 subgraphs). M246 in-scope fix (comment only).
- **KB-4** — Injection CODE (`INJECTED` dict `:17`, enum `:458`, `exposure_claim_guard.py:124`) still lists skillpath as a live service. Out of the comment-only declared scope → **go/no-go scope-watch** for the bring-up; observed behavior → M247 drift ledger.
- **KB-5** — Corpus asserts skillpath live Tier-1 (~30 files). **Fate-2 confirmed covered by M247** (corpus re-ground) — no M246 deferral, no M246 edit. M246 does not read these as truth.

## D-1 — Seeder re-point: hard-cut to `public`, no dual-schema tolerance (2026-07-23)
**Decision.** Re-point every `skill_path_sessions` write from schema `skillpath` → `public` in one
shot (live seeder code + tests + DNA + comments); do NOT add a runtime dual-schema fallback.
**Rationale.** The open question in overview.md asked dual-schema-tolerant vs hard-cut. Hard-cut wins:
(a) the demo builds from a SINGLE pinned clone set (this milestone bumps it to current `origin/main`,
where skillpath is decommissioned) — there is no mixed-topology demo to tolerate; (b) a dual-schema
seeder would need a live schema probe per write (complexity + a new failure mode) to serve a
transition window the demo never occupies; (c) `rext` is consumed per-stack at a pinned tag, so a
stack still on the old platform pin simply consumes the OLD rext tag — the version skew is handled by
the tag pin, not by seeder branching. Commit `97585f5` (rext).
**Scope kept tight.** Surface NAME `skillpath-sessions`, Go symbols, and the mirror
`public.local_skill_path_sessions` untouched (the mirror was already `public`).
**Empirical dependency.** Correctness of the target (`public.skill_path_sessions` existing) is proven
by the cold `/demo-up` — the seeder COPY to public succeeds only if the consolidated migrations create
it there. That is the go/no-go, not an assumption baked in here.

## D-2 — Clone pins: durable canonical pin in rext + copy-if-absent seam (2026-07-23)
**Decision.** The `DEMO_ADVANCE_CLONES=pinned` consume path was already wired (M237). Rather than
author an ephemeral workspace-only pin, ship a DURABLE canonical `demo-stack/clones.pin.json` in rext
(12 repos @ current origin/main HEAD shas; skillpath EXCLUDED — absent from current repos.yml) + a
copy-if-absent seam in `ensure-clones.sh` that seeds it into the git-ignored `stack-demo/` workspace,
never clobbering an operator's own pin. Commit `ee44b9a` (rext).
**Rationale.** The barrier's payoff is a REPRODUCIBLE consolidated topology that M247-M254 all build
against. An ephemeral pin evaporates on `/demo-down`+`/demo-up` (demos are disposable). A canonical
pin consumed at the pinned rext tag makes the topology reproducible on a fresh box. A SHA pin lands
`pinned-detached` → the freshness gate treats it as fresh → `DEMO_FRESHNESS_STRICT=1` passes for the
HARD go/no-go.

## D-3 — Section 3 scope expansion: de-skillpath the LIVE bring-up path (2026-07-23)
**Decision.** The declared Lane-D scope was "fix the gen_injected_override.py:16 comment." The compose
check (current origin/main `docker-compose.yml`) proved skillpath has **no service** — so
`up-injected.sh` would (a) BUILD `demo-N-skillpath:injected` from the stale skillpath clone and (b)
VERIFY a skillpath container that can't exist. Both BLOCK a green bring-up. So I expanded section 3 to
drop skillpath from `up-injected.sh` `INJECT_SVCS` + `verify_svcs` (rext tooling — **zero
platform-repo edits**). Commit `88bcdb8` (rext).
**Why Fate-1, not scope-creep.** These are REQUIRED for the milestone's core deliverable (green
bring-up on the consolidated platform); the barrier exists precisely to surface + fix such drift. The
`INJECTED`-dict/`test_injection.py`/`exposure_claim_guard._cfg` skillpath residue is NOT required for
green (inert without a skillpath compose service) → ledgered for M247 (the designed handoff), NOT a
disguised deferral (it has a documented home).

## D-4 — Go/no-go PASS: the consolidated platform builds + comes up green (2026-07-23)
**Decision.** The HARD go/no-go barrier **PASSES**. One cold `/demo-up` (demo-2, `DEMO_ADVANCE_CLONES=pinned
DEMO_FRESHNESS_STRICT=1 --no-public-host`) on the consolidated platform came up GREEN on every load-bearing
axis: build exit 0 (16 services), the re-pointed seeder wrote **561 rows to `public.skill_path_sessions`**
(legacy `skillpath.skill_path_sessions` = 0-row husk), **3 subgraph images + 0 skillpath**, health 200 +
casbin 1250 + all liveness/readiness probes pass. The ~386-commit `app` bump did **not** surface a
migration/schema/subgraph break — so downstream milestones (M247–M254) are safe to scope against this
proven topology; no re-scope needed (overview Open-question #2 → answered NO).
**The 3 autoverify warnings are non-firing** (drift/probe-artifact, not platform breaks): D-07 AI-readiness
`loadMembers` demopatch anchor went stale (file moved in the bump; non-fatal perf; → M250); D-08 fake-FAPI
"not answering" is an http-vs-TLS cheap-win probe artifact (container up + roster-loaded; → M251/M254);
D-09 academy peripheral (missing keys; non-fatal by design). All ledgered (`drift-ledger.md` §4).
**Environment (billion-last honored).** LOCAL demo-2 on this box; the stale demo-1 was torn down first to
free resources; the dev native worktrees + the `billion` VM were untouched. Zero platform-repo edits.
**Rung-zero.** The demo consumed rext at tag `july-jitter-m246-re-sync-repoint` fetched from origin (on
origin @ `44f1049`) — one tag covering the three rext commits 97585f5 / ee44b9a / 88bcdb8.

## D-5 — Harden: right-sized to ONE pass; deepen the seam, not the mechanical re-point (2026-07-23)
**Decision.** One honest hardening pass. Add NOTHING to the Go schema re-point (it is mechanical +
already asserted per-site + live-proven at 561 rows); deepen ONLY the 2 genuinely-untested branches of
the copy-if-absent clone-pin seam. Rext commit `c9fbf6b`; new consumption tag `july-jitter-m246-harden`
@ `9b29f3a` **on origin**.
**Rationale.** The re-point flips one schema qualifier (`skillpath`→`public`) for `skill_path_sessions`.
Every write-site's test already asserts the `public` target (reset list + FK-order in `main_test.go`;
`findCopy(t,"public",…)` in the seeder tests; `failTable:"public.skill_path_sessions"` in the copy-error
paths), and the barrier's cold demo-2 wrote 561 rows to `public.skill_path_sessions` — the strongest
possible proof. Forcing more passes there would be shallow-test throughput, not robustness. A
"no-`skillpath`-schema-anywhere" grep-guard was REJECTED (fragile: the surface NAME `skillpath-sessions`,
Go symbols, and `skillpath_sessions.go` legitimately keep the word). The 2 new seam tests close real
behavioral contracts the shipped land-tests miss: (a) pinned mode with NO pin anywhere → clean no-op
fallback; (b) the seam runs before the advance gate, so shipping the canonical pin must not silently turn
a default/main `/demo-up` into a pinned checkout (the deliberate-staleness contract). Flake gate 3/3;
full `test_tooling.py` 168/168; Go package green.
**Fate.** All Fate-1 (landed now). No deferrals; the M247 drift ledger is unchanged by this pass.

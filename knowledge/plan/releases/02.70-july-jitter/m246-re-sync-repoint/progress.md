# M246 — Progress

Section milestone. Checklist from the roadmap In-list. **HARD go/no-go barrier — gates M247–M254.**

## Sections

- [x] **Seeder re-point** — rext `stack-seeding` writes `skillpath.skill_path_sessions → public.skill_path_sessions` in the **live** seeder code + tests (`cmd/stackseed/main.go:97`, `seeders/hero_activity.go:180`, `skillpath_sessions.go`, `content_nonsim.go`, `dna/data-dna.json` + the in-package test assertions). Surface **names** (`skillpath-sessions`) + the mirror `public.local_skill_path_sessions` left untouched. **DONE** — rext `97585f5`; 8 live sites + DNA + comments + ~16 test assertions; build/vet/gofmt clean, package tests green, zero missed sites (all-file-types sweep).
- [x] **Demo clone pins** — author `stack-demo/clones.pin.json` + wire the `DEMO_ADVANCE_CLONES=pinned` advance path + bump the **demo** clone pins to current `origin/main` (jobsimulation stays standalone — still live). **DONE** — rext `ee44b9a`. The advance path was ALREADY wired (M237); delivered the durable canonical pin (`demo-stack/clones.pin.json`, 12 repos @ current origin/main HEAD shas, **skillpath excluded** — not in current repos.yml) + a copy-if-absent seam that seeds it into the ephemeral workspace (never clobbers an operator pin). 2 new tests; module green.
- [x] **Injection-comment fix** — `stack-injection/gen_injected_override.py:16` skillpath comment → 3 subgraphs. **DONE + EXPANDED** — rext `88bcdb8`. Scope grew to "de-skillpath the LIVE bring-up path" once the compose check proved skillpath has no service: also dropped skillpath from `up-injected.sh` `INJECT_SVCS` (was building `demo-N-skillpath:injected`) + `verify_svcs` (was verifying a skillpath container) — both required for a green bring-up. The `INJECTED`-dict/`test_injection.py`/`exposure_claim_guard` skillpath hygiene is inert (no compose service) → M247 drift ledger.
- [x] **Cold `/demo-up` GREEN + drift ledger** — prove ONE cold `/demo-up` GREEN on the consolidated platform; transcribe observed drift into the M247 ledger. **DONE — GO/NO-GO PASS.** Cold demo-2 (pinned advance + strict freshness, local, billion untouched) came up green: build exit 0, 16 services, **561 rows in `public.skill_path_sessions`** (re-point proven), 3 subgraphs + 0 skillpath, health 200 + casbin 1250, all probes pass. 3 autoverify warnings all non-firing (D-07 AI-readiness demopatch→M250, D-08 fake-FAPI probe artifact, D-09 academy peripheral). Drift ledger `drift-ledger.md` D-01..D-09 emitted for M247. rext consumed at tag `july-jitter-m246-re-sync-repoint` (on origin).

## Completeness Ledger

### Deferred

### Dropped

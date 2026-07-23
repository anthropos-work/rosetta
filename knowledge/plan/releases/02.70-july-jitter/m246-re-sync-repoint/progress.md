# M246 — Progress

Section milestone. Checklist from the roadmap In-list. **HARD go/no-go barrier — gates M247–M254.**

## Sections

- [x] **Seeder re-point** — rext `stack-seeding` writes `skillpath.skill_path_sessions → public.skill_path_sessions` in the **live** seeder code + tests (`cmd/stackseed/main.go:97`, `seeders/hero_activity.go:180`, `skillpath_sessions.go`, `content_nonsim.go`, `dna/data-dna.json` + the in-package test assertions). Surface **names** (`skillpath-sessions`) + the mirror `public.local_skill_path_sessions` left untouched. **DONE** — rext `97585f5`; 8 live sites + DNA + comments + ~16 test assertions; build/vet/gofmt clean, package tests green, zero missed sites (all-file-types sweep).
- [ ] **Demo clone pins** — author `stack-demo/clones.pin.json` + wire the `DEMO_ADVANCE_CLONES=pinned` advance path + bump the **demo** clone pins to current `origin/main` (jobsimulation stays standalone — still live).
- [ ] **Injection-comment fix** — `stack-injection/gen_injected_override.py:16` skillpath comment → 3 subgraphs.
- [ ] **Cold `/demo-up` GREEN + drift ledger** — prove ONE cold `/demo-up` GREEN on the consolidated platform; transcribe observed drift into the M247 ledger.

## Completeness Ledger

### Deferred

### Dropped

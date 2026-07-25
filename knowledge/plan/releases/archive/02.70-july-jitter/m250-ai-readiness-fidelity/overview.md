---
milestone_shape: iterative
milestone: M250
title: "AI-readiness fidelity"
status: archived
last_updated: 2026-07-24
release: v2.7 "july jitter"
exit_gate: "On a cold reset-to-seed, for a completed Northwind AI-readiness member: (1) step-1 AI Skill Mapping renders the platform 31 default readiness skills (19 core + 12 enabling), not invented; (2) step-2 AI Simulation shows the correct track-keyed named sim (tech=who-can-see-this-document-fc0 / business=use-ai-to-turn-survey-data-into-a-leadership-email) + interview, with a non-empty evaluated-skills list of that sim real evaluated node-ids; (3) the member profile carries the completed sim distributed verified skills (validation fan-out + user_skill_evidences); (4) the manager AI-readiness view shows the same faithfully; (5) 0 invented values, 0 prod-ejects, closure green, frozen-vs-live arithmetic agrees at the 31-skill repertoire."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
re_scope_trigger: "5 consecutive toks without a viable strategy -> user-strategic-replan"
depends_on: [M246]
complexity: large
created: 2026-07-23
---

# M250 — AI-readiness fidelity  (`iterative`, marquee)

**Status:** `archived` (completed 2026-07-24)  ·  **Shape:** `iterative`  ·  **Complexity:** large  ·  **Release:** v2.7 "july jitter"  ·  **Depends on:** M246

## Goal
The `/ai-readiness` page renders the platform's real fidelity: the **31 canonical mapping skills**, the **2
track-keyed named sims + interview**, a **non-empty evaluated-skills list**, and the completed sim's **verified
skills distributed to the employee** — faithful for **player AND manager**.

## Shape (why this shape)
`iterative`. The path is exploratory because it combines three unknowns that resist a single-shot section build:
the **8→31 arithmetic re-derivation** across ~200 members (the funnel + the M219 arithmetic fences + the
"Champion 30/30" beat must all be re-derived at the new 31-skill repertoire), a **net-new directus-write
set-dress** (snapshot replay is replay-only, so writing `directus.simulations.skills` has no existing seam), and
**live-render believability** (0 invented values, 0 prod-ejects proven only by rendering the page). The
measure→triage→fix→re-render loop of `coverage-protocol.md` + `verification.md` is the right instrument.
Risk-map R2 (blocks-scope) is the reason: reuse `content_stories_write.go`'s verified-skill fan-out, keep the
`iterative` shape, and re-derive the M219 fences + the "Champion 30/30" beat rather than assume the old numbers.

## Scope
### In
- **Arithmetic spine** — re-derive the **8→31** readiness-skill repertoire (**19 core + 12 enabling**) across
  ~200 members: the AI-readiness **config + funnel + the M219 arithmetic fences**, as **one atomic edit**;
  re-derive the **"Champion 30/30" beat** at 31 skills so the frozen-vs-live arithmetic agrees (gate parts 1 + 5).
- **Directus set-dress (net-new)** — a directus-write set-dress that populates step-2 **"AI Simulation"** with the
  correct **track-keyed named sim** (tech=`who-can-see-this-document-fc0` / business=`use-ai-to-turn-survey-data-into-a-leadership-email`)
  **+ interview**, each carrying a **non-empty evaluated-skills list** of that sim's **real evaluated node-ids**
  (write `directus.simulations.skills`). A net-new file — snapshot replay is replay-only (gate part 2).
- **Evidence distribution** — distribute the **completed sim's verified skills** to the member's profile: the
  **validation fan-out + `user_skill_evidences`** (reuse `content_stories_write.go`'s verified-skill fan-out).
  Behind both lanes above (gate part 3).
- **Manager-vantage fidelity** — the **manager** AI-readiness view shows the same faithfully (gate part 4).
- **Believability render loop** — measure→triage→fix→re-render on a cold reset-to-seed until **0 invented values,
  0 prod-ejects, closure green**, and the frozen-vs-live arithmetic agrees at the 31-skill repertoire (gate part 5;
  the serial iterative loop).

### Out
- **Any platform edit** — the fill routes through the **existing resolvers** / a **directus set-dress**, never a
  platform-repo change.

## Dependencies & parallelism
- **Depends on:** M246 (the HARD go/no-go re-sync barrier — fidelity work on stale pins is untrustworthy). Branches
  from **post-M246 HEAD**.
- **Parallel with:** M247 / M248 / M249 / M251 / M252 (fan-out worktrees). **Live-iteration contends with M253 for
  the box:** M250 + M253 are both live-measured iteratives and **serialize on one billion demo** (RAM won't hold
  two); M253 can bootstrap its FCP loop on a **local** demo, cold-p95 confirmed in M254 (coordination rule 9).
- **Intra-milestone LANE decomposition:** **~1.6× on iter-01 only.** Two concurrent lanes then a serial join:
  - **Lane A — arithmetic-spine** (config + funnel + M219 fences, **one atomic edit**)  ∥
  - **Lane B — directus-set-dress** (net-new file)
  - **→ evidence-distribution** (serial, behind **both** A and B — it distributes verified skills the set-dress
    named and the arithmetic sized).
  - **The iterative loop after iter-01 is serial** (measure→triage→fix→re-render is single-threaded).
  - **Recommended subagents:** ~**2** on iter-01 (Lane A ∥ Lane B), then **1 serial driver** for the join and the
    render loop.
- **Shared-file coordination:** `cmd/stackseed/main.go` — the single seeder registry, touched by **both M248 +
  M250**; each edits only its own `MustRegister`/truncate hunk → clean hand-merge (rule 2). Rung-zero every push —
  rext tags on **origin** before billion re-pins (rule 7). The one-line `CLAUDE.md` bullet defers to **M247** (sole
  owner, rule 5).
- **Merge/close order (release-level):** M251 → { M248, **M250** } → M249 → M253 → M252 → M247-reconcile → M254.

## KB dependencies
- `corpus/services/ai-readiness.md`
- `corpus/ops/seeding-spec.md`
- `corpus/ops/demo/stories-spec.md`

## Delivers
- `corpus/services/ai-readiness.md` + `corpus/ops/seeding-spec.md` — the **31-default + 2-named-sim + track +
  evaluated-skills set-dress + skill-distribution** seeding contract.

## Exit gate
On a **cold reset-to-seed**, for a **completed Northwind AI-readiness member**:
1. step-1 **"AI Skill Mapping"** renders the platform's **31 default readiness skills** (**19 core + 12 enabling**),
   not invented ones;
2. step-2 **"AI Simulation"** shows the correct **track-keyed named sim**
   (tech=`who-can-see-this-document-fc0` / business=`use-ai-to-turn-survey-data-into-a-leadership-email`) **+
   interview**, with a **non-empty evaluated-skills list** of that sim's **real evaluated node-ids**;
3. the member's profile carries the completed sim's **distributed verified skills** (validation fan-out +
   `user_skill_evidences`);
4. the **manager** AI-readiness view shows the same faithfully;
5. **0 invented values, 0 prod-ejects, closure green**, and the **frozen-vs-live arithmetic agrees** at the
   31-skill repertoire.

## Iteration protocol
- **Protocol:** `corpus/ops/demo/coverage-protocol.md` + `corpus/ops/verification.md` (measure→triage→fix→re-render);
  contract `corpus/services/ai-readiness.md`.
- **Re-scope trigger:** **5 consecutive toks without a viable strategy → user-strategic-replan.**

## Open questions
- How to **write `directus.simulations.skills`** in the per-stack Directus? (net-new set-dress — snapshot replay is
  replay-only, so there is no existing write seam.)
- The **tech/business track ↔ audience label mapping** — the platform pins the **opposite** of the annotation's
  framing; **confirm at live render**.
- **Re-derive** the M219 arithmetic fences + the **"Champion 30/30" beat** at 31 skills.

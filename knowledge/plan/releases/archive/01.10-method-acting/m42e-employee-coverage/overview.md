---
milestone: M42e
slug: employee-coverage
version: v1.10 "method acting"
milestone_shape: iterative
exit_gate: "Logged in as an EMPLOYEE hero (Maya) on a stack brought up by a FRESH demo-up, the coverage sweep asserts — per page AND per section/element of every employee-reachable page — (a) REAL semantic content: actual seeded user/catalog content, NOT placeholder/empty-state copy ('add something here') and NOT bare chrome; (b) SUBSTANTIAL cardinality: each content section shows a meaningful count of items (not just 1), except documented exceptions where 0/1 is genuinely correct; (c) PERSONA self-consistency: the hero's role, skills, bio, real-person avatar (consistent across menu + profile), work-history and activity cohere as one believable person, and the org has a name + logo; (d) NO prod-eject escape: no in-app nav/menu/button ejects the presenter to a prod anthropos surface (legitimate external editorial/reference links in content are allowed but disclosed as presenter-notes). Gate = 0 sections below the content/cardinality/consistency bar + 0 prod-eject escapes, AND the same result reproduces on a FRESH demo-up (not just the live-patched stack)."
iteration_protocol_ref: corpus/ops/demo/coverage-protocol.md
status: archived
created: 2026-06-24
last_updated: 2026-06-25
complexity: large
delivers: corpus/ops/demo/coverage-protocol.md (NEW — the Playwright demo-coverage sweep + triage + fix iteration protocol) + the Playwright coverage harness (a new rext section, demo-only — the first non-Go rext dev/test dependency) + updates across the touched specs (frontend-tier.md / verification.md / seeding-spec.md / snapshot-spec.md) as gaps close
depends_on: M39+M40+M41
spec_ref: .agentspace/profile_gaps.md (live-demo review, 2026-06-24; root-cause workflow w7t4wq2z4)
---

# M42e — Employee 100% demo coverage

## ⚠ Re-scope (2026-06-25 — user live-review feedback)
The original gate measured **DOM text-density** (`textLen > 40` in `<main>`) — far too weak. It passed pages
that render placeholder/empty-state cards ("add something here") + nav chrome, so the harness reported a green
`(0,0)` while a logged-in presenter saw `/profile` placeholder-only, `/library/ai-simulations` empty,
incoherent 3D skills for a backend dev, a silhouette avatar, and no org logo. The gate below is **re-scoped to
the believability bar**: real semantic content + substantial per-section cardinality + persona self-consistency,
reproducible on a **fresh demo-up**. (Detail: [[demo-coverage-semantic-content-gate]] in agent memory.)

## Goal
A hero of the **EMPLOYEE/member vantage** (Maya), logged into the demo platform, sees a **believable, fully-
populated person and catalog** — every reachable page and **every section/element** shows real seeded content
(no placeholders, no empty-state copy, no single-item "filler"), the persona is **internally coherent**
(role ↔ skills ↔ bio ↔ a real face ↔ work history), and nothing ejects the presenter to production. And it must
look this way on a **fresh `demo-up`**, not just a hand-patched stack.

## Exit gate (objective, machine-verifiable)
Logged in as Maya on a stack from a **FRESH `demo-up`**, the coverage sweep asserts — **per page AND per
section/element** of every employee-reachable page:
- **(a) Real semantic content** — actual seeded user/catalog content; **placeholder / empty-state copy and bare
  chrome do NOT count** as content.
- **(b) Substantial cardinality** — each content section shows a **meaningful count** of items (not just 1),
  **except** documented exceptions where 0/1 is the genuinely-correct state.
- **(c) Persona self-consistency** — the hero's **role, skills, bio, avatar (a real-person photo, consistent
  across the menu AND the profile), work history, and recent activity cohere as one believable person**; the
  **org has a name + logo**.
- **(d) No prod-eject escape** — no in-app nav / menu / feature-button takes the presenter to a **prod anthropos
  surface** (e.g. left-menu "Studio" → `studio.anthropos.work`). Legitimate **external editorial/reference links
  inside content** (citations; a LinkedIn-import help link) are **allowed but disclosed** in a presenter-notes
  list — they are not prod-ejects.

**Gate = 0 sections below the content/cardinality/consistency bar + 0 prod-eject escapes, reproduced on a FRESH
`demo-up`.** A new **semantic** coverage harness (not the text-density one) measures (a)–(d) per section, with a
documented exception list, and screenshot review is part of acceptance.

## Why iterative (not section)
The page set and the failure modes — empty section / federation error / out-of-demo link / missing seed — are
**discovered BY the sweep**, not enumerable up front. Each iter measures (run the sweep) → triages the failures
→ fixes in `rosetta-extensions` (stack-seeding seeds / stack-snapshot serve-grants / the demo injection+env
link-rewriting / a corpus doc) → re-sweeps, until the gate is GREEN for an employee hero. A fixed `In:`
checklist would be speculative — the gate is the commitment, the path emerges from each tik's evidence. Build
with `/developer-kit:build-mstone-iters`.

## Iteration protocol
The loop is authored **by this milestone** as [`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md)
— the Playwright demo-coverage sweep + triage + fix loop. M42e **also delivers** the Playwright coverage
harness itself (a new rext section — `stack-coverage`, or under `stack-verify` — demo-only; runs against the
live demo on offset ports). Note: **Playwright is the first non-Go rext dev/test dependency** — sanctioned by
the coverage requirement.

## Approach sketch
Build the harness (a rext section) that:
1. **logs in** as a roster hero via the cockpit handshake,
2. **crawls** the in-app nav as that vantage,
3. **per page** asserts DOM non-emptiness + captures a screenshot + asserts every link host is demo-local,
4. **emits** a coverage report (pages, pass/fail, escapes).

Then iterate the fixes into `rosetta-extensions` until the gate is GREEN for an employee hero. The fix surfaces
per failure mode:
- **empty section / missing seed** → `stack-seeding` (seed the data the page reads),
- **federation / content error** → `stack-snapshot` serve-grants (the page reads replayed content),
- **out-of-demo link** → the demo injection + env **link-rewriting** (rewrite the host to the offset port),
- **a documented gap** → a corpus doc update.

**Re-scope trigger:** if a 100%-blocking gap can **ONLY** be closed by a platform change (the platform repos
are read-only this release), **escalate** (the zero-edit line) and **record it** — do not edit
next-web/app/cms/jobsimulation.

## Depends on / Parallel with
**Depends on:** M39 + M40 + M41 — the targeted fills (profile identity, Directus serve-grant, profile depth)
land first so the coverage loop only chases the **tail** of remaining gaps, not the known-and-already-scoped
ones. **Parallel with:** none.

## Open questions
- Harness home: a **new `stack-coverage`** rext section vs **under `stack-verify`** — decide in the bootstrap
  iter against how much it reuses verify's offset/port plumbing.
- Crawl strategy: pure in-app nav-link discovery vs a seeded route manifest — decide against escape-detection
  fidelity (a manifest can't catch a nav that escapes the demo).
- "Non-empty semantic content" assertion shape: per-section DOM selectors vs a generic text-density floor —
  decide per-iter against the false-pass/false-fail rate the sweep surfaces.
- Playwright wiring: how the first non-Go dev/test dependency is pinned + invoked alongside the Go rext tooling
  (its own dir + lockfile under the rext section).

## KB dependencies
Read as contract:
- [`corpus/ops/demo/frontend-tier.md`](../../../../corpus/ops/demo/frontend-tier.md) — the demo UI tier
  (next-web + studio-desk + ant-academy, offset ports, minted-pk + offset-URL baked) the sweep crawls.
- [`corpus/ops/verification.md`](../../../../corpus/ops/verification.md) — the bring-up auto-verify net (the
  scoped, non-fatal `verify live` on the stack's own offset ports) the coverage harness extends/sits beside.
- [`corpus/ops/rosetta_demo.md`](../../../../corpus/ops/rosetta_demo.md) — the demo-stack lifecycle + offset
  ports + Clerkenstein injection the cockpit login rides.
- [`corpus/ops/demo/stories-spec.md`](../../../../corpus/ops/demo/stories-spec.md) — the verified-skill chain +
  the roster hero (Maya) the sweep logs in as.

## Delivers →
- **NEW** [`corpus/ops/demo/coverage-protocol.md`](../../../../corpus/ops/demo/coverage-protocol.md) — the
  iteration protocol (the Playwright demo-coverage sweep + triage + fix loop).
- The **Playwright coverage harness** — a new rext section (demo-only; the first non-Go rext dev/test
  dependency).
- **Updates** across the touched specs — `corpus/ops/demo/frontend-tier.md`, `corpus/ops/verification.md`,
  `corpus/ops/seeding-spec.md`, `corpus/ops/snapshot-spec.md` — as gaps close.

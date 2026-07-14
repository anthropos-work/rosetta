# M43 Decisions

Implementation decisions with rationale (recorded during build). Design-time decisions live in
[`overview.md`](overview.md) + the research note
[`.agentspace/scratch/roadmap-research-2026-06-26.md`](../../../../.agentspace/scratch/roadmap-research-2026-06-26.md).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| D1 | **Aesthetic = clean modern professional LIGHT** (card hierarchy, generous spacing, one indigo accent, high-contrast typography). | The user's ask: "too dark, ugly, unusable → slick, professional, nice". Light flat high-contrast is the decided default (overview Open-question 1). Not a ground-up design-system pass (out of scope). | 2026-06-26 |
| D2 | **FontAwesome via the FREE CDN** (cdnjs FA6-free `<link>` in `<head>`), NOT vendored. | The user explicitly asked for the free CDN link. A CDN `<link>` is a runtime asset, not a build dep — supply-chain stays GREEN, cockpit.py stays stdlib-only. ant-academy vendors FA Pro locally (`code/public/assets/fontawesome/`) as the offline-safe precedent **if** a future offline-demo need surfaces — recorded, not adopted now (overview Open-question 2). | 2026-06-26 |
| D3 | **One CTA per hero: `[Log in as]` → the hero's `jump_to`** (per-role landing), NOT the app root. Drop `[Jump to section]` entirely. | The user: login is the ONLY CTA. The manifest already carries a resolved `jump_to` per hero (declared or `defaultJumpForVantage`), so the unified CTA reuses existing data — no `cockpit.go` change required. The bare `[Log in as]` shows the name only (no jump-label), overview Open-question 3's "stays bare" default. | 2026-06-26 |
| D4 | **Manifest download = the JSON** (`/manifest.json`, `Content-Disposition: attachment`), not the source `stack.stories.yaml`. | The `/manifest.json` endpoint already exists + is the single-sourced projection the panel reads; readily available, no new endpoint (overview Open-question 4). | 2026-06-26 |
| D5 | ⚠️ **SUPERSEDED by v2.3 M218 — see the box below.** **Login-progress overlay = staged deterministic** (`Signing you in…` → `Loading your workspace…` → `Almost there…`) with `localStorage` state carried across the FAPI-handshake → next-web redirect; a generous final stage. | No real cross-origin progress signal exists from the handshake; the overlay is FEEDBACK, not telemetry. ~~The real ~2–5s latency is unavoidable — only the feedback improves~~ **(FALSE — see below)** (overview scope "Out"). localStorage persists the overlay across the cross-origin redirect so it survives the blank-load. | 2026-06-26 |

| KB-1 | **Reconcile `stories-spec.md` § "The presenter cockpit (M38)" after the CTA unification** — its two-CTA description (`[Login as]`→root + `[Jump to section]`→jump_to) is superseded by M43's single unified `[Log in as]`→jump_to. Add a supersession pointer to the new `cockpit-spec.md`; do NOT delete the M38 historical narrative. | Phase 0b (KB-fidelity, YELLOW) tracking item: today's docs accurately describe pre-M43 code; M43 deliberately supersedes the two-CTA model. Addressed in Phase 5 docs work alongside authoring the `Delivers → cockpit-spec.md`. | 2026-06-26 |

> ### ⚠️ D5 IS SUPERSEDED — formally re-opened by **v2.3 "cue to cue" M218** (2026-07-14)
>
> **D5's load-bearing claim was false, and it was never measured.** It asserted the cockpit-login latency was
> *"~2–5 s"* and *"unavoidable — only the feedback improves"*, and on that basis put the latency itself
> **out of scope**. The real number, when someone finally built an instrument and measured it, was
> **39.45 s (employee) / 38.30 s (manager)** — and up to **60–120 s** on the tailnet VM the demo actually
> ships on.
>
> It was **two tooling bugs**, both fixable, neither touching platform source:
> 1. next-web's **SSR** GraphQL origin was a build-inlined *public* URL that **blackholes from inside the
>    container** (`3 × 10.5 s` connect-timeout + a 2 s/4 s retry ladder ≈ **37.5 s**).
> 2. Clerkenstein's fake BAPI served **a hardcoded stub user to every hero**, so the JWT identity and the
>    BAPI identity disagreed and `app` refused `userPreferences` → another **~6 s** of retry ladder.
>
> **Shipped: cold p95 2413 ms / 1767 ms** over 5 consecutive cold reset-to-seed cycles — a **~16×**
> improvement on the honest cold number.
>
> **Why this decision is the cautionary tale of the release.** The "~2–5 s, unavoidable" claim propagated
> into **four** corpus sites, was booked as an M43 scope-`Out:` with **zero deferrals recorded** — so it
> never entered any ledger — and was therefore never revisited across **four releases** (v1.10b, v2.0, v2.1,
> v2.2). **Nobody investigated because the docs said there was nothing to investigate.**
> **Do not write "we can't fix this" about something you have never measured.**
>
> The overlay itself (the actual D5 deliverable) **stands** and is still shipped — it is now a *safety net*
> for a slow box rather than the normal path. See [`corpus/ops/demo/latency-budget.md`](../../../../../../corpus/ops/demo/latency-budget.md).


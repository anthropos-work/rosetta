**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-34 — the clause-5 confirming pass, and the pass that audits the repair

## What ran

The confirming full-read pass iter-33 routed forward. **40 files, 8 530 lines, five read-only auditors,
every file read top-to-bottom with a `wc -l` positive control. 40/40 read to their last line; 0 UNREAD.**

The partition was **deliberately re-cut** (`D-M257x-34-2`): pass 1 grouped by subject, this pass grouped by
line balance with swept/unswept files deliberately mixed, so no auditor inherited pass 1's group
boundaries. Correlated blind spots are a property of how a corpus is divided, not only of who reads it.

Ground truth was the iter-33 brief (derived against the same platform sha, **`2adcf71`, re-fetched
unchanged at this iter's open**) plus an addendum — `iter34-confirm-brief.md` — that gave every auditor
both halves of the prior: *do not trust the swept text*, and *do not assume the unswept text is clean.*

## Both pre-registered predictions were wrong. That is the result.

Recorded in `overview.md` at 00:36, before any report existed:

| prediction | outcome |
|---|---|
| **P1** — the pass returns **1–5** blockers, most likely 2–3 | ❌ **REFUTED — 11**, more than double the top of the range |
| **P2** — the largest cluster lands in the **27 unswept** files | ❌ **REFUTED, decisively and in the opposite direction** |

**Where the 11 actually landed:**

| | files | blockers | per file |
|---|---|---|---|
| **swept by iter-33's repair** | 13 | **9** | **0.69** |
| never edited | 27 | **2** | **0.074** |

**A file the repair pass touched was ~9× more likely to contain a blocker than a file it never opened.**
Two auditors volunteered this unprompted: group D's four never-edited files "produced **zero** blockers and
verified exact on every pin I checked, including the newest drift"; group A's two mechanically-verifiable
untouched files "produced **zero** findings across ~40 exact citations and three exact counts."

iter-33 measured a 24 % self-inflicted rate and treated it as a caution. **iter-34 measures the same
effect at 82 % of the residual and it is no longer a caution — it is the dominant term.** The corpus's
remaining fidelity debt is concentrated almost entirely in text written to repair fidelity debt.

## Grading: 11 reported, 11 verified, **0 downgraded**

`D-M257x-34-1` fixed the adjudication rule *before* any report was read, precisely so the count could not
drift toward whatever answer was convenient. Every blocker was re-derived against platform source by the
parent before being acted on:

| # | file | the false claim | verified against |
|---|---|---|---|
| 1 | `external_services.md` | the `--local-content` re-point targets `cms` only, `backend` must **not** carry it | `gen_injected_override.py:53` = `("cms","backend")`; the cited test was replaced by `test_injection.py:1005` |
| 2 | `ai-readiness.md` | default (`CycleID == nil`) GET is "hardcoded to `buildLiveResponse`" | `readiness.go:307-312` takes the frozen path on the no-active/has-closed shape — **exactly what M51 seeds** |
| 3 | `ai-readiness.md` | both gates must hold "for the UI to render"; `AIReadinessClient.tsx` checks the flag | that file has **0** posthog refs; gates on `orgEnabled` alone (`:133-134`) |
| 4 | `architecture_overview.md` | routing is Azure EU → Bedrock EU → Mistral EU → OpenAI US | `ai.go:262-276` — Azure EU → **Azure US** by flag; Bedrock is a per-call vendor, not a tier; Mistral is not in the cascade (it *is* in `app`, for Studio OCR — see S1 below, where my first correction overshot) |
| 5 | `hiring.md` | `anticheat_summary` is a column of `job_simulation_sessions` | 0 occurrences in `job_simulation_session.go`; it was on the **dropped** mirror (`20250416091037.sql:5`) |
| 6 | `hiring.md` | the hiring container is "wired to … Cosmo" | contradicted **4 lines earlier in its own paragraph** |
| 7 | `backend.md` | the RPC mux carries `SkillPathSessionService` | **0** occurrences in Go source; five unconditional handlers + a conditional `CMSService` at `main.go:1178-1218`, none skillpath. (Its 2 non-Go hits are in **`app`'s own docs** — Trap C: the platform's docs lag its code, and that is where the corpus got the claim) |
| 8 | `backend.md` | `aiacademy/` populates `aiacademy_courses` | package absent; `academy/academy.go:6-9` says both were removed |
| 9 | `skillpath.md` | the manager scoreboard reads mirror `local_skill_path_session`; "the mirror row must be co-written" | `intelligence.go:1144` queries `ent.SkillPathSession` directly; the mirror was `DROP TABLE`d at `20260729133514.sql:63` |
| 10 | `security_compliance.md` | `org_membership.go` heads the list of "no policy at all" | it declares its own fail-closed `Policy()` at `:172-188` — **the only one of the 18 that does** |
| 11 | `clerk-integration.md` | JS SDKs "aligned … both on `@clerk/clerk-expo ~2.6.18`" | `ant-academy/mobile:18` pins `~2.19.36` — thirteen minor versions apart |

**All 11 fixed.** Five corpus guards green before and after.

## The finding that matters most: my own re-derivation was right about the number and wrong about the claim

Before any report arrived, the parent independently re-derived the tenancy fence — the fence that had
already been wrong **twice**, both times failing toward *"isolation is handled"* (`evidence/parent-rederivation.md`).
It confirmed all three numbers exactly: **30 / 7 / 18**, plus all nine named example files. It concluded
the fence held.

**It did not hold.** Auditor E found that `org_membership.go` — listed *first* among the schemas with
"no mixin and no policy at all" — **declares its own `Policy()` ending in `AlwaysDenyRule`**.

The parent verified the **denominator** and never checked the **predicate**. The sentence says "no mixin
**and no policy at all**"; the check tested only the first conjunct. **A count can be exactly right while
the claim it supports is false.**

And the correction *that* produced was wrong too: the self-audit found `academy_feedback.go` also policed
(via `UserMixin{}`'s owner filter), taking the unpoliced count **17 → 16**, with **31** org-filtered. So
the fence has now been wrong **four** times — three toward *"isolation is handled"* and mine toward alarm.
**The error direction is not stable**, which is a stronger argument for re-deriving than any single
direction would be. The rewritten fence says so in its own text and ships a derivation command — one that,
unlike my first attempt, was **verified to reproduce its own numbers (18, then 16) before commit**.

## Clause 5 is NOT MET

The gate wants **GREEN, or YELLOW with 0 blockers**. The reading taken was **11 blockers**. They are
fixed, and the post-fix state is — by construction — unmeasured. iter-33 refused to grade a clause on an
absent measurement; that refusal binds symmetrically and pre-registered rule `D-M257x-34-1(4)` committed
to it in advance. **Gate stays 3 of 5.**

**Is this convergence?** 19 → 6 → 11 is not a curve, and it should not be read as one: the three passes
had different scopes (40 files unswept / 13 files swept / 40 files post-repair). The comparable pair is
**pass 1's 19 over 40 files vs pass 3's 11 over the same 40** — a real reduction, with the residual
now concentrated 9× in repaired text rather than spread across the corpus. The honest projection is that
a fourth pass over freshly-repaired text finds a **smaller but non-zero** number, and the way to end the
regress is not more passes but the structural answer already routed as
`CHECK-M257x-iter33-derived-fact-fence`.

An adversarial self-audit of this iteration's own 8 files was launched before close, on the same
reasoning — see `evidence/audit-selfcheck.md`.

## Evidence

`iter34-confirm-brief.md` (the addendum) · `evidence/audit-{a..e}.md` (five pass-3 reports with their
positive-control tables, 1 328 lines) · `evidence/parent-rederivation.md` (the pre-report re-derivation
that was right about the number and wrong about the claim) · `evidence/audit-selfcheck.md`.

## The self-audit — and the one finding that indicts this iteration's own method

The adversarial pass over this iteration's 8 changed files returned **2 blockers + 8 minors**, both
blockers **self-inflicted**, both in **surrounding prose**. Every `file:line` anchor written this iter
verified exact — for the second iteration running. Both verified and fixed:

**S1 — `architecture_overview.md`: I asserted an absence from a search that had FAILED.**
The correction said *"Mistral is not in `app`'s routing at all — its only platform reference is
`terraform/ssm.tf:291`."* Mistral is **live Go code in `app`**:
`internal/cms/studio/markdownManager.go:11,19` builds a Mistral client from `MISTRAL_API_KEY` for Studio
document OCR, called at `studioManager.go:583`.

The mechanism is the damning part. The check run was `grep -ril mistral stack-demo/app --include=*.go`,
which **zsh rejected** (`no matches found: --include=*.go`) — and the empty output was read as *absence*.
That is **§5 rule 1 of this milestone's own protocol** — *never let a search's stderr go unread; an engine
rejection is indistinguishable from "no matches"* — committed **in the same iteration that added rules 17
and 18 to that section.** Knowing a rule and executing it are different things, and the rule fired against
its own author within the hour.

**S2 — `security_compliance.md:154`: a GDPR residency claim left failing toward reassurance.**
*"US providers (OpenAI Direct, Anthropic Direct) used only as fallback / No customer data stored in US by
default"* — while `ai.go:263-277` routes to **Azure OpenAI US on a PostHog flag**, which is a switch, not
a fallback; Azure US was absent from the residency list entirely; and "Anthropic Direct" is never used at
all (Bedrock `eu-west-1`, `:85-95`). **I added the correct warning to `architecture_overview.md` and walked
past the same claim in the file I was editing** — structurally identical to iter-33's
`organization_id`-on-every-table miss, one iteration later.

**And the fence was wrong a fourth time — in the opposite direction.** My rewrite listed
`academy_feedback.go` among the unpoliced; it carries `UserMixin{}`, whose `Policy()` applies a row-level
owner filter. 17 → **16**. The first three generations of this fence failed toward *"isolation is
handled"*; mine failed toward alarm. **The error direction is not stable**, which is a stronger reason to
re-derive than any one direction would be — the doc now says exactly that.

**The derivation command I shipped did not reproduce the number I shipped beside it.** `grep -L` over a
bare `*.go` glob pulls in `skiller_mixins.go` and returns **19**, not 18. A recipe added so the next reader
would not have to trust the prose was itself untrustworthy. Now schema-restricted, and **verified to
return 18 then 16** before commit.

## Final tally

| | blockers |
|---|---|
| confirming pass over 40 files | **11** (9 in the 13 repaired files, 2 in the 27 untouched) |
| adversarial pass over this iter's own 8 files | **2** (both self-inflicted, both in prose) |
| **closed this iteration** | **13** |

Five corpus guards green. Protocol gained §5 **rule 17** (verify the predicate, not the denominator) and
**rule 18** (repaired text is the highest-risk text; re-partition the confirming pass).

## Close — 2026-08-02

**Outcome:** the clause-5 confirming pass ran — 40/40 files read in full — and returned **11 blockers**,
not the 0 the clause needs. All 11 verified against platform source and fixed; an adversarial pass over
those fixes found **2 more, both self-inflicted**, also fixed. **13 blockers closed.** The measurement's
real product is the **distribution**: 9 of 11 sat in the 13 files the previous repair had touched vs 2 in
the 27 it never opened — **~9× density** — so *repaired text*, not unswept text, is where this corpus's
fidelity debt now lives.
**Type:** tik
**Status:** closed-fixed (planned scope was *run the confirming pass, then fix by evidence rank*; both
landed, plus the self-audit the milestone's own rules require)
**Gate:** NOT MET (**3 of 5**. Clause 5 needs GREEN-or-YELLOW-with-0-blockers; the reading taken was 11)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik; the prior three tiks all
moved a metric) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open AND close, unchanged;
trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) —
(6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-34-1` … `D-M257x-34-5` (`iter-34/decisions.md`).
**Side-deliverables:** `platform-alignment.md` §5 **rule 17** + **rule 18** (protocol evolution, same
commit). **No rext change and no re-pin** — clause 5 is a rosetta-corpus clause; pin stays
`fast-build-m257x-iter-31b`.
**Routes carried forward:**
- `MEASURE-M257x-iter35-clause5-fourth-pass` — but **scope it to the 9 files this iter changed**, not all
  40. The 27 untouched files have now been read in full twice with 2 blockers between them; re-reading
  them is the low-yield half. Rule 18 says where to look.
- `CHECK-M257x-iter33-derived-fact-fence` — **promoted in priority.** Three passes have now each found
  blockers, and the residual is being *manufactured by the repair*. Passes do not converge on their own;
  a fence over derived facts is the only thing that ends the regress.
- `DOC-M257x-iter33-corpus-minors` — now ~66 more from this pass (the auditors' minor lists are in
  `evidence/audit-{a..e}.md`, each with exact anchors).
- `CHECK-M257x-iter22-clerk-sdk-drift` — **materially advanced**: the JS half is now measured, and
  `@clerk/clerk-expo` is **13 minor versions apart** across the two mobile surfaces.
- Clause 2's three survivors, unchanged. `pt-activity-drilldown`'s coupling was scoped read-only this
  iter (see `decisions.md` prep note); the fix shape is *select the drill target by hero participation
  rather than by grid position*.

**Lessons:**
- **Repaired text is the most dangerous text in a corpus.** 82 % of this pass's blockers were in 32 % of
  the files — the ones a repair had touched. Never close a corrective sweep without an adversarial pass
  over its own diff, and re-partition the confirming pass so it cannot inherit the last one's blind spots.
- **A count can be exactly right while the claim it supports is false.** Verify the predicate — every
  conjunct of it — not the denominator. The parent's pre-report re-derivation reproduced 30/7/18 exactly
  and still missed that the first-listed file polices itself.
- **Knowing a rule does not execute it.** This iteration added §5 rules 17 and 18 and, within the hour,
  violated §5 **rule 1** — reading an empty result from a shell-rejected `grep` as proof of absence, then
  writing that absence into the corpus. The protocol's oldest rule caught its newest author.
- **Check the direction of your own correction, not just its truth.** Three generations of the tenancy
  fence failed toward reassurance; mine failed toward alarm. Unstable error direction is itself the finding.
- **A derivation shipped to prevent trusting must be run.** The recipe added so readers need not trust the
  prose returned 19 where the prose said 18.

**Type:** tik — TOK-02 **step 4 of 5** (*"repair the 18 once, fence-assisted — by CLAIM not by FILE,
tree-wide, with the fence as the commit post-condition"*).

# iter-46 — the 18 repaired, by claim

## Phase A — re-derive, then repair

Every claim's "what is true" was re-measured against `stack-demo/` at `app` @ `5ba17044` **before a word
was written** (`D-M257x-46-1`). Two came back stronger than the ledger recorded: `#2`'s uncounted exit is
the `default:` arm reached by a **nullable-and-unset** `AIVendor` (the ordinary path, not a
misconfiguration) plus a **fifth** route nobody counted — a caller explicitly selecting `Openai`; and
`#4`'s Anthropic-direct path is selected by an **env var, not a flag**, so the `flag_use_azure_us` caveat
four lines below it does not cover it.

`#5`'s base count, which iter-41 **deliberately left unsettled** between two auditors, was settled by
counting (`D-M257x-46-2`): **E's 16 is right**, the total is **23**, and the defect was never the 16 — it
was a closing sentence that excluded the 7 `OrganizationIDMixin` schemas **three lines after naming them
as unpoliced**.

`#17` was repaired **by hand**, exactly as `D-M257x-45-3` routed it: its anchor resolves and names a
construct, just the wrong one for 3 of the 4 domains, and no fence in this family can see that.

## Phase B — the fence found what the repair missed

After the first pass all 18 anchored sites were fixed. `claim_twin_guard` still reported **three**
(`D-M257x-46-3`) — a correction appended while the original sentence stayed, a second EU-residency site
untouched because only the PM summary had been fixed, and a phrasing that survived a rewrite.

**All three are the class iter-41 measured as 8 of the 9 repair-induced blockers.** Under the previous
method they would have been committed and counted by the next full read. **This is the first pass in the
series where that class was caught by a machine instead of by the next audit.**

Two further live findings outside the 18 — a stray unterminated code fence and an anchor onto a blank
line — were **repaired rather than baselined** (`D-M257x-46-4`).

## Phase C — measurement

| fence | before | after |
|---|---|---|
| `claim_twin_guard` | 18 sites | **0** |
| `markdown_structure_guard` | 2 | **0** |
| `anchor_construct_guard` | 2 | **0** |
| `derived_value_guard` | 2 | **0** |
| ratchet baseline | 25 sites / 4 fences | **0 sites / 4 fences** |

**GREEN was verified to mean the corpus is clean, not that the fences broke** (`D-M257x-46-5`, §5 rule 8):
`anchor_construct_guard` resolves **101** anchors across 112 files (up — repaired anchors now resolve),
`derived_value_guard` measures **5** service docs, and **both perishable fixtures still go RED** with
their green twins still silent (53 tests). `stack-core` **491 tests / 14 failures**, exactly the
pre-existing baseline.

Files touched, by claim rather than by file: `ai_architecture.md` · `architecture_overview.md` ·
`architecture/README.md` · `external_services.md` · `platform-migration-status.md` ·
`security_compliance.md` · `service_taxonomy.md` · `shared_libraries.md` · `services/ai-readiness.md` ·
`cms.md` · `graphql-wundergraph.md` · `messenger.md` · `roadrunner.md` · `sentinel.md` ·
`ops/demo/coverage-protocol.md` · `.claude/skills/stack-update/reference.md` · `CLAUDE.md`.

## Close — 2026-08-02

**Outcome:** all 18 of iter-41's blockers repaired in one pass, by CLAIM and tree-wide, every "what is
true" re-derived from platform source rather than from the ledger — and **the fence caught three sites
the repair itself left standing**, the exact class that accounted for 8 of the 9 repair-induced defects
in every prior pass. All four fences GREEN, the ratchet baseline lowered **25 → 0**, and GREEN
independently verified to mean a clean corpus rather than a broken fence.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5. **Clause 5 is not graded here** (`D-M257x-46-7`): only TOK-02 step 5's full
7-auditor read at iter-41's frozen instrument grades it, and reading a fence's GREEN as that number is
the mistake iter-38 and iter-21 both paid for.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this was a tik) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open and close, unchanged; trigger stays at occurrence 1 of 2) — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — **Outcome: continue**
**Decisions:** `D-M257x-46-1` … `D-M257x-46-8`
**Side-deliverables:** the two live fence findings outside the 18 (`D-M257x-46-4`) — repaired in the same
pass because a baseline that accumulates what nobody wants to fix stops being a ratchet.
**Routes carried forward:**

- **`READ-M257x-iter41-instrument` — TOK-02 step 5, and the only thing that grades clause 5.** ONE full
  7-auditor read, instrument held fixed at iter-41's: same seven auditors, same briefing, same partition
  method, all 40 files top-to-bottom. **iter-47.**
- `CHECK-M257x-iter35-seeder-writes-one-instant` — still the highest-value open non-gate item.
- `CHECK-M257x-iter38-ai-act-classification` — needs an owner **outside** this milestone.

**Lessons:**

- **The fence's value showed up in the gap between "I fixed all 18" and "the fence agrees".** Three sites
  survived a careful, claim-scoped, tree-wide repair by its own author. Not a different defect class from
  the eighteen — *the same one*, arriving on schedule, and stopped at the commit for the first time.
- **Re-derive from source even when the ledger looks conclusive.** Two of the eighteen were understated
  by their own adjudication, and the corrected wording is materially different: an unset nullable is not
  a mistyped string, and an env var is not a feature flag.
- **Settle a deliberately-unsettled number by counting, not by choosing.** iter-41 recorded two auditors
  disagreeing and refused to pick. Counting took minutes and showed the disagreement was never the
  defect.
- **Four fences going 25 → 0 in one pass is exactly the shape §5 rule 8 warns about**, so it has to be
  falsified rather than celebrated. The two perishable fixtures are what made that a two-minute check
  instead of an argument.

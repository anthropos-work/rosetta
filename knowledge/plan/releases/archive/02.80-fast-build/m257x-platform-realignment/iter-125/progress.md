# iter-125 — the second corpus, and the register entry re-derived at source

**Type:** tik · **Run 80, tik 2.** Priorities 2 and 3 of the user's directed scope.

## 1. P3 — the directus finding: the conclusion strengthens, one anchor was wrong

`D-M257x-121-1` made re-derivation **mandatory at filing time**, including when the routing is our own.
This is the first entry filed under that rule to be re-derived after filing, and the rule earned its
place immediately.

**Both repos were re-cloned at the exact refs the entry names** — `directus` `d6325731`, `infrastructure`
`13c248e6` — and `git rev-parse HEAD` returned those refs, so the substrate is the one the entry claims
(the standing rule: *before believing a defect, read the substrate line*).

| what the entry claimed | re-derived |
|---|---|
| `api.ts:9` is a bare `console.log(env);` | **CONFIRMS VERBATIM.** And it is the **only** `console.log(env` in the entire repository — 1 hit |
| `d6325731` = tag `v0.20.15` | **CONFIRMS.** `git tag --points-at HEAD` → `v0.20.15` |
| production pins it (`services.tf:24`) | **CONFIRMS, and the pin is named TWICE** — `:24` the module `?ref=v0.20.15`, `:30` the ECR image tag `production-directus:v0.20.15` |
| `services.tf:47-57` threads `SECRET`, `KEY`, the Postgres connection and password, the admin credentials, `GCLOUD_SERVICE_ACCOUNT` | **WRONG ANCHOR.** `:47-57` are the **root module's INPUTS**, and `SECRET`/`KEY` are not among them |

**Where the container's environment actually comes from**, read from the task definition rather than from
the module's call site — `directus/terraform/main.tf` @ `d6325731`:

- **six ECS-injected secrets** (`:224-249`): `SECRET`, `ADMIN_PASSWORD`, **`DB_PASSWORD`**,
  `AUTH_GOOGLE_CLIENT_SECRET`, `DB_SSL__CA`, `GCLOUD_SERVICE_ACCOUNT`;
- **plus `KEY` (`:111-114`) — which is NOT in the `secrets` block at all.** It is a plain `environment`
  entry whose value interpolates `aws_ssm_parameter.directus_key.value`, so the secret is materialised
  into the **task-definition JSON in clear** as well as into the container env. **A second, independent
  exposure of the same value, which the original filing did not have and which nobody was looking for.**

**So the finding is confirmed and strengthened**: an invocation writes the database password, the Directus
signing secret, the admin password and the Google client secret to the container log group. The
correction is stated **first** in the entry, not folded in.

**Nothing in `corpus/` asserts the opposite** — checked. The corpus's only statement about this extension
was `org-repos.md`'s inventory row, which said what the operation does and not that it logs `env`; it now
points at the register entry, so a reader meeting the extension is told.

## 2. P2 — the second corpus, stated where a reader meets it

`org-repos.md` § 11 already carried the full comparison (iter-123). **What it could not do is reach the
reader**: a taxonomy reader opens `shared_libraries.md`, and an engineer installing the plugin opens
`toolchain_overview.md`. Neither opens a repo census. So the directive's instruction — *state it where a
reader would meet it* — is a **placement** requirement, and placement was the whole gap.

**Three placements landed:**

1. **`shared_libraries.md#taxonomy-figures`** — the section that already holds this corpus's own figures
   now carries the contradiction: both figures, both provenances, and which is measured.
2. **`toolchain_overview.md`, on the install line** — because that recommendation is what puts the
   refuted figure into an engineer's editor on every Anthropos repo.
3. **`platform-defect-register.md`** — `PLATFORM-M257x-akb-taxonomy-figures-contradict-measurement`,
   filed so it has an **owner outside this milestone**. AKB is a different repo; no edit here can reach it.

**Two verdicts, kept apart, because merging them over-claims:**

| | measured here | AKB | verdict |
|---|---|---|---|
| job roles | **≥ 22,470** (public subset, `organization_id IS NULL`, 2026-06-29, reproducible) | 18,000, unsourced ×14 | **REFUTED** — public ⊆ total, so 18K is below the floor |
| skills | **≥ 42,790** (same capture) | 60,000, unsourced ×14 | **UNVERIFIED, not refuted** — a public-only capture cannot see org-private skills |

*"AKB is wrong"* would be false on the second row. The register says so.

**A lead was recorded rather than a conclusion:** `public.job_role_embeddings` holds **18,919** rows — a
different table, a plausible mis-transcription onto the role count. Nothing here can measure what AKB's
author read, so it stays a hypothesis and is labelled one; it is filed because it gives the owner a place
to start.

### And the direction of the errors is recorded, because the directive is right about it

**On the WunderGraph router's production residue, AKB was RIGHT and this corpus was WRONG — in a fenced
table.** AKB reads the `infrastructure` repo this corpus had never cloned. The two have **different blind
spots, not a ranking**: this corpus is authoritative for measured local/runtime state and ops, AKB for
`infrastructure`-derived production state and product/GTM, and **neither cites the other**.

**Found while doing it, and not in the directive's framing:** `org-repos.md` § 11 item 1 still read
*"Unresolved"* and *"this corpus says the module is still declared"* — **one screen below § 3 of the same
file, which had measured it destroyed.** A document contradicting itself across two sections is the same
one-cell-reach failure iter-124 found 24 more of, arriving inside the very section that documents the
disagreement. Repaired, with AKB's correctness recorded.

## 3. Reach, with denominators

| statement | number | denominator |
|---|---|---|
| register anchors re-derived | **4 of 4** | the claims the entry makes |
| of those, confirmed verbatim | **3** | same |
| of those, corrected | **1** (the environment inventory) | same |
| net-new exposure found by re-derivation | **1** (`KEY` in the plain `environment` block) | — |
| AKB placements landed | **3** | the three the directive named |
| corpus sites asserting the opposite of the directus finding | **0** | `git grep` over `corpus/` |

## 4. Guards

Tree guards re-run after the edits: `claim_census_guard` **OK — ratchet holds (1,160, baseline 1,164)**;
`corpus_citation_guard` **OK**; `anchor_construct_guard` **OK**; `markdown_structure_guard` **OK**.
`repair_postcondition` **OK** at the pre-commit hook. The full family was last read at iter-124's close —
**17 GREEN · 0 RED · 1 could-not-check · 4 not-run** — and is re-read at run close rather than quoted
from arithmetic.

## 5. Routes carried forward

- **`platform_alignment_guard`'s could-not-check** — carried to iter-126, the next tik of this run.
- `FIX-M257x-iter125-akb-comparison-tables-exposure` — which of AKB's four customer-facing comparison
  tables have been published externally is **not measurable from here**, and it decides whether this is a
  documentation defect or a customer-communication one. Named in the register entry.

## Close — 2026-08-07

**Outcome:** the directus register entry **re-derives verbatim on 3 of its 4 claims and its environment
inventory was wrong** — corrected, correction first, and the re-derivation found a **net-new second
exposure** (`KEY` materialised in clear into the task definition). The AKB contradiction is stated in the
**three places a reader actually meets it** and filed with an owner outside this milestone; **AKB's
correctness on the router residue is recorded**, and `org-repos.md` § 11's self-contradiction repaired.
**No `N` movement is claimed and no reading was taken.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — 4 of 5, unchanged — (2) triggered-tok: n (**a successor strategy
remains FORBIDDEN by `TOK-08`'s sealed rule; this tik runs under the user's directed scope**) —
(3) re-scope: n — (4) user-blocker: n (**the register is FILING, not escalation — §5 rule 48, and the
directive says so explicitly**) — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n —
(7) budget-exhausted: n — Outcome: **continue**
**Decisions:** see [`decisions.md`](decisions.md)
**Side-deliverables:** none — both lines were planned scope.
**Lessons:** **A contradiction documented in a census is not a contradiction disclosed.** iter-123 wrote
the AKB comparison in full and it changed nothing a reader would see, because the reader meets the
figures in `shared_libraries.md` and the plugin in `toolchain_overview.md`. **Disclosure is a placement
problem, not a writing problem** — file the finding where the decision is made, not where the
investigation happened. → `platform-alignment.md` §5 **rule 55**.

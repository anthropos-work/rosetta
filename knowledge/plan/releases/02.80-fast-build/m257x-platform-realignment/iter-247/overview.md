---
iter: 247
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-247 — fence the corpus's ABSENCE claims about environment variables

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey

iter-246 closed `ROUTE-M257x-238` and opened
`ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence`: of the four runnable-input surfaces
iters 235–238 censused and repaired, `make` targets and `cd` directories are now fenced; **environment
variables and frontend scripts/ports are not.** `ROUTE-M257x-237-critical-env-list-is-unfenced` is the
same gap from the other side, and iter-237 called it *"the strongest fence candidate this run."*

**But the obvious fence is a trap, and the substrate says so before any of it is built.** The live corpus
carries **327 distinct `UPPER_SNAKE` tokens across 1,213 occurrences** — most are not environment
variables at all (enum members, rext knobs, Go constants), and, worse, the most-cited *are* env vars the
platform **deleted**: `CMS_RPC_ADDR` ×38, `SKILLER_RPC_ADDR` ×19, `STORAGE_RPC_ADDR` ×18. A
"name must exist in `.env_example`" fence would RED the corpus's most careful writing. **Absence is the
point of those sentences.**

## Hypothesis

**Invert it. Fence the ABSENCE claims, which is where the polarity is unambiguous and where rot actually
lands.** A sentence that says a variable is *"set nowhere"* is a claim about the platform that was true
when written and that the platform can falsify at any time — and unlike a presence claim, it has exactly
one reading.

Substrate, enumerated before sealing: **12 distinct variables carry an explicit absence claim, across 29
sites** (`set nowhere` / `is not set` / `read by nothing` / `occurs in zero` / `does not set` …), headed by
`JUDGE0_BASE_URL` ×6 and `STORAGE_RPC_ADDR` ×5.

## Pre-registered numeric claims — SEALED IN THIS COMMIT

Graded against `stack-demo/platform` (`docker-compose.yml` + `.env_example`), named.

| id | claim | prediction |
|---|---|---|
| **P-247-1** | of the 12, those actually PRESENT in compose or `.env_example` — i.e. the absence claim is false *or* the instrument mis-attached it | **2–4** |
| **P-247-2** | ≥ 1 is a **co-location artifact** — the absence phrase belongs to a different variable on the same line | **YES** |
| **P-247-3** | no existing guard grades an env-var absence claim | **CONFIRMED** |
| **P-247-4** | all **6** `*_RPC_ADDR` names (BACKEND_USERS, CMS, JOBSIMULATION, SKILLER, STORAGE, ROADRUNNER) are genuinely absent from compose | **YES, 6 of 6** |
| **P-247-5** | **control** — the 5 names CLAUDE.md lists as critical (`GH_PAT`, `CLERK_SECRET_KEY`, `OPENAI_KEY`, `VITE_CLERK_PUBLISHABLE_KEY`, `DIRECTUS_TOKEN`) are all declared in `.env_example` | **HELD** (iter-237 verified them; the control proves the instrument can report a present variable as present) |

**Falsification:** if **0 of 12** absence claims are false AND the co-location artifact does not appear,
the class is clean and trivially so; the deliverable is then the fence alone, and the iter says so rather
than dressing a zero as a finding.

## Phase plan

A — grade the 12 against the platform clone. B — separate false claims from mis-attachments.
C — repair. D — fence, with the attachment rule the census establishes. E — re-derive last.

## Escalation

A variable absent from a **demo** stack's compose but present under another profile is a **scope**
question, not a false claim — grade against the file, and say which file.

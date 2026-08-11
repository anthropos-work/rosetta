# AUDITOR C — 7 files / 1520 lines

**Positive control:** all 7 read to final line; counts match `wc -l`
(alignment_testing 521 · architecture_overview 341 · cms 254 · sentinel 166 · messenger 128 ·
services/README 79 · db-backup 31).

## BLOCKERS — 1

| # | site | the false claim | what is true |
|---|---|---|---|
| **C1** | `architecture_overview.md:246-247`, compounded at `:251-252` | *"the **two US paths** are a **feature flag** and a **429 retry target**, not fallback rungs"* … *"one feature flag routes traffic to the US"* | **There is a THIRD US path and it is unconditional.** `app/internal/cms/directus/collections/jobsimulation.go:1302` — `aiVendor := simulation.Openai` is the default when `seq.AIVendor == nil`. That vendor reaches `getClient`'s `case Openai: client = a.openaiClient`, built at `internal/jobsimulation/ai/ai.go:80` as `openai.NewOpenAI(openaiKey)` — direct `api.openai.com`, **no region override**. A sequence authored with no `ai_vendor` reaches **direct US OpenAI on the first attempt: no PostHog flag, no HTTP 429.** |

**Why this is a blocker and not a minor.** It is residency-relevant — a reader would act on it — and **the
corpus's own cited source says the opposite and warns against exactly this conclusion**:
`external_services.md:532` reads *"**two ways in**: (a) `vendor = Openai` from the caller — including the
case where the caller never chose, since a simulation sequence with `ai_vendor` unset defaults to `openai`
(`internal/cms/directus/collections/jobsimulation.go:1302`) … The 429 retry is the only automatic fallback
— but it is **not** the only route to US OpenAI. **Path (a) gets there on the first attempt.**"*

**Provenance: this text was written by M257x iter-46**, to correct an over-claim. It is the
**retraction-as-new-claim** class — a retraction that substitutes a new false enumeration and drops
precisely the path its own source flags as most consequential. The lead-in *"the default clients are
EU-resident"* inherits the same defect.

## MINORS — 9

| # | site | what is off |
|---|---|---|
| 1 | alignment_testing.md:193 | `gate.sh:61` is a **comment** line; the actual invocation is `:69`. Substance + `--if-declared` semantics verified exactly (bare exit `2` vs `--if-declared` exit `0`, measured on all five DNAs with a built binary) |
| 2 | alignment_testing.md:316 | *"live in the **clerkenstein repo**, not here"* — clerkenstein is a **section** of the rext monorepo, which this same file states correctly at `:490-497`/`:500-507`. Same-file inconsistency |
| 3 | alignment_testing.md:513-518 | Layout omits `internal/canon/`; the `cmd/alignctl` verb list omits `coverage`, documented in the same file at `:245` |
| 4 | alignment_testing.md:320 | the scripts dir also holds `drift-test.sh` |
| 5 | cms.md:110 | `requirements.txt` listed as including **`python-docx`**; the actual file has no such entry |
| 6 | cms.md:8-9 | *"the **fourth and last** engine consolidated into `app`"* vs `architecture_overview.md:62-64` *"**Five** former microservices…"*. Cross-doc count framing (roadrunner is elsewhere classed "orphaned, not merged") |
| 7 | services/README.md:20 | *"**three of the four**"* — roadrunner is explicitly *"the **fifth**"* nine lines above at `:15`. Should be "three of the five". The enumerated three and their compose lines are correct |
| 8 | messenger.md:110 | *"only the residual `SKILLPATH_STREAM=skillpath` remains"* sits in **messenger's** env table, but that var lives in the **backend** block (`docker-compose.yml:64`). True at compose scope, misleading at row scope. Also `:107` anchors `app/main.go:1199` for *"additive + DORMANT"*, which is on `:1198` |
| 9 | messenger.md:23-40 | `internal/flow/` omits `ai_readiness.go`, `content_assigned.go`, `content_completed.go`, `coursebuilder.go`, `invitation_reminders.go` |

## Files read clean

- **`sentinel.md` — 0 blockers, 0 minors.** Every checkable anchor resolves **and** names the right
  construct, including **both iter-46 repairs**: `go.mod:3` `go 1.26.0` ✓, `Dockerfile:2`/`Dockerfile.dev:2`
  `golang:1.26-bookworm` ✓, `terraform/locals.tf:4-5` `service_cpu = 256` / `service_memory = 128` ✓.
  Also: Casbin **v3** ✓; **6 request / 6 policy / 3 role-grouping / 6 matchers** exact against
  `casbin.go:14-44`; `AUTHORIZATION_ADDRESS` in exactly three compose blocks ✓; **no `manager` role** ✓;
  **`messenger` is not a caller** — `grep -rn authorization --include=*.go messenger/` returns **zero** ✓.
- **`db-backup.md`** — clean; no falsifiable claim against a cloned repo.
- **`alignment_testing.md`** — no blockers. Independently **executed**: the toy transcript at `:293-306`
  reproduces **byte-for-byte** (86.7% / 100.0% / 5-of-6, `FAIL Greet/padded-name`); all five DNA
  gene/capability counts exact; `alignment/` has **no** `scripts/` dir ✓; stdlib-only ✓.
- **`cms.md`** — no blockers beyond minors 5-6. The **3→1 subgraph correction is right**:
  `graphql-wundergraph@915da06` deletes **both** `cms.graphqls` and `jobsimulation.graphqls`, and the
  5→4→3 ladder checks out; `app/main.go:1196-1202` "additive + DORMANT … until the **M809** re-point" is an
  **exact range** ✓; `skillpath/session.go:205-207` exact ✓.
- **`architecture_overview.md`** — apart from **C1**, the merge/router rewrite is solid: `2adcf71` verified,
  **no `:5050` anywhere** ✓, six Go services on a bare `make up` ✓, and **the multi-tenancy numbers are
  exactly right** — `grep -l ent.Schema` = **135** schemas, `OrganizationMixin{}` = **30** ✓.
- **`services/README.md`** — apart from minor 7, clean.

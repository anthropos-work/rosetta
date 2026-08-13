#!/usr/bin/env python3
"""sweep33b.py — M257x iter-33 clause-5 corpus sweep, groups 2-5 BLOCKERS.

Same contract as sweep33.py: enumerated (file, old, new); `old` MUST occur EXACTLY ONCE;
two-phase (validate every anchor in every file, THEN write). Not idempotent by design.
"""
import sys
from pathlib import Path

ROOT = Path("/Users/marco/workspace/anthropos/rosetta/corpus")

EDITS = [
    # ================= g2-B1 — proto skew did NOT disappear =================
    ("architecture/shared_libraries.md",
     "| **Version pin** | one pin per live repo — the cms/jobsimulation version skew disappeared with the merge (they no longer have their own `go.mod`) |",
     "| **Version pin** | one pin per repo, and the skew is **live and wider than the merge suggests**: app/messenger `v1.210.0`, cms `v1.207.0`, jobsimulation `v1.205.0`, sentinel `v1.200.0`, storage/roadrunner `v1.196.0`. The husk repos **still carry their own `go.mod`** — `repos.yml:14-19` still clones them and `docker-compose.yml:83,144` still builds them |"),

    # ================= g2-B2 — skillpath/roadrunner RPC removed, not re-hosted =============
    ("architecture/shared_libraries.md",
     "| **Imported by** | every live Go service that does RPC (app, sentinel, storage, messenger; the cms / jobsimulation / skiller / skillpath / roadrunner RPC surfaces are all served in-process by app) |",
     "| **Imported by** | every live Go service that does RPC (app, sentinel, storage, messenger). The cms / jobsimulation / **skiller** RPC surfaces are served in-process by `app`; **skillpath and roadrunner were REMOVED, not re-hosted** — `app/main.go` registers six Connect handlers (Users `:1178`, Organizations `:1179`, Skiller `:1187`, JobSimulation `:1195`, CMS `:1204`, LabSession `:1218`) and neither `SkillPathSessionService` nor a RoadRunner service is among them |"),

    ("architecture/shared_libraries.md",
     "`SkillPathSessionService` (served by app since the skillpath merge, M502→M507),",
     "`SkillPathSessionService` (**contract still in `proto`, but NO LONGER SERVED** — like `ChronosService`. "
     "skillpath-in-app M506 *removed* the RPC rather than re-hosting it; "
     "`app/internal/skillpaths/skillpaths.go:27-31` calls its replacement \"the drop-in for the **removed** "
     "skillpath RPC client\". Likewise roadrunner: `backend` calls Judge0 over plain HTTP at "
     "`internal/jobsimwiring/wiring.go:118`, and `ROADRUNNER_RPC_ADDR` (`docker-compose.yml:118`) is read by "
     "no Go code in `app`),"),

    # ================= g3-B4 — the tenancy absolutes are false =================
    ("architecture/security_compliance.md",
     "### Layer 1: Database\n"
     "- Every table has an `organization_id` foreign key\n"
     "- Ent ORM policies auto-filter all queries by organization\n"
     "- No cross-tenant data access is possible at the query level",
     "### Layer 1: Database\n\n"
     "> **⚠️ Isolation is NOT automatic across the whole schema — do not rely on it as a blanket guarantee.**\n"
     "> Measured at `app` HEAD: of **139** Ent schemas, only **30** use `OrganizationMixin{}` — the one that\n"
     "> carries the privacy `Policy()` (`mixin.go:126`). Seven use `OrganizationIDMixin{}`, explicitly *\"a plain\n"
     "> nullable organization_id column\"* with **no policy**, and the rest never mention organization at all.\n"
     "> The platform states this itself: `job_simulation_session.go:5` — *\"L2: NO Ent privacy Policy;\n"
     "> owner/org/tenant are plain fields\"* — and `jobrole.go:18` / `category.go:15` note the taxonomy is\n"
     "> deliberately globally readable. **Scoping on the jobsim fan-out and the taxonomy is the caller's job.**\n\n"
     "- Org-scoped tables carry an `organization_id` column\n"
     "- Ent privacy policies auto-filter by organization **only on the 30 schemas using `OrganizationMixin{}`**\n"
     "- Cross-tenant reads are prevented at the query level **on those tables**; elsewhere isolation is\n"
     "  enforced by Layer 2 (Sentinel) and by explicit query scoping, not by the ORM"),

    # ================= g4-B1 — completion_status is spelled correctly =================
    ("services/hiring.md",
     "`completition_status` (**note the misspelled column**; values `passed`/`failed`/`pending`/`SIMULATION…`),",
     "`completion_status` (values `passed`/`failed`/`pending`/`SIMULATION…`) — **spelled correctly in the DB**\n"
     "   (`app/terraform/migrations/20260722104506.sql:12`, `ent/schema/job_simulation_session.go:39`); the\n"
     "   `completition` misspelling survives only in the GraphQL sort-field enum\n"
     "   (`enum.InsightsSortFieldCompletitionStatus`) and a JSON tag, **never as a column name**,"),

    # ================= g4-B2 — ai-readiness is its own package =================
    ("services/backend.md",
     "Engine: `internal/workforce/ai_readiness.go` + `readiness_steps.go` + `readiness_narrative.go`;",
     "Engine: its own top-level package **`app/internal/aireadiness/`** (`manager.go`, `cycles.go`,\n"
     "  `diagnosis.go`, `compare.go`, `csv.go`, …) — **not** `internal/workforce/`, which contains no `readi*`\n"
     "  file at HEAD;"),

    ("services/backend.md",
     "* **AI Readiness** (v1.266+, the `internal/workforce` subsystem):",
     "* **AI Readiness** (v1.266+, the `internal/aireadiness` package):"),

    # ================= g4-B3 — the mirrors were DROPPED =================
    ("services/jobsimulation.md",
     "> **not** read this service's `sessions` table at all — it reads an `app`-side MIRROR, `public.local_jobsimulation_sessions`\n"
     "> (the analog of skill-path's `local_skill_path_session`). Seed the runtime rows only and the manager scoreboard\n"
     "> is blank.",
     "> reads the **same** table — **the mirrors are GONE.** `app/terraform/migrations/20260729133514.sql:58-62`\n"
     "> (*\"5. Drop the mirrors.\"*) back-fills then `DROP TABLE`s both `local_jobsimulation_sessions` and\n"
     "> `local_skill_path_sessions`, and `intelligence.go:1700` now reads `m.ent.JobSimulationSession.Query()`.\n"
     "> **There is one row to seed, not a pair** — the older \"seed the mirror or the scoreboard is blank\"\n"
     "> guidance is superseded.",),

    # ================= g4-B4 — `sessions` was renamed =================
    ("services/jobsimulation.md",
     "`app/terraform/migrations/20260722081626_jobsim_data_model.sql`, with the **same table names**. The old",
     "`app/terraform/migrations/20260722081626_jobsim_data_model.sql`. **Most kept their names — but the\n"
     ">   headline one did NOT:** the very next migration, `20260722104506.sql`, creates\n"
     ">   `job_simulation_sessions` (`:2`) and `DROP TABLE \"sessions\"` (`:79`). **`public.sessions` does not\n"
     ">   exist**; the session table is `public.job_simulation_sessions`. The old"),

    # ================= g4-B5 — the cutover rides on `backend`, not `cms` =================
    ("services/jobsimulation.md",
     "So the M23 content cutover (re-pointing CMS's `DIRECTUS_BASE_ADDR` at the per-stack Directus) carries jobsimulation's content reads to local automatically; no jobsimulation env change is needed.",
     "**The M23 content cutover does NOT ride on the `cms` husk.** `backend` is the in-process Directus reader "
     "(`app/cms_reader_switch.go`; `app/main.go:971-973` `log.Fatalf`s without `DIRECTUS_BASE_ADDR`), so "
     "re-pointing `cms` alone leaves `backend` reading prod — measured live on `demo-1` at M257x iter-24 as 96 "
     "Directus log lines, all 403. rext therefore sets `DIRECTUS_DATA_CONSUMERS = (\"cms\", \"backend\")` in both "
     "twins. No jobsimulation env change is needed, but the cutover must include `backend`."),

    # ================= g4-B6 — same defect at its source =================
    ("services/cms.md",
     "and **M23 re-points `cms`'s `DIRECTUS_BASE_ADDR` at that local instance**",
     "and **M23 re-points `DIRECTUS_BASE_ADDR` at that local instance — for `cms` AND for `backend`** "
     "(⚠️ the `cms` re-point alone is **not sufficient**: since cms-in-app, `backend` is the in-process "
     "Directus reader via `app/cms_reader_switch.go`, so a stack that re-points only `cms` still reads prod. "
     "M257x iter-24 measured that as 96 all-403 Directus lines in `backend`'s log; `DIRECTUS_DATA_CONSUMERS` "
     "now names both)"),

    # ================= g5-B1 — storage's only live caller is `app` =================
    ("services/storage.md",
     "Storage is the **centralized file/blob service** for the platform. Other services (`jobsimulation`, `cms`, `app`) push and pull binary objects through it instead of dealing with S3 themselves.",
     "Storage is the **centralized file/blob service** for the platform.\n\n"
     "> **⚠️ Since the merges, the sole live caller is `app`.** The jobsimulation and cms domains run\n"
     "> **in-process inside `backend`** (`app/internal/jobsimulation/recording/recording.go:12`,\n"
     "> `anticheat.go:34`, `app/main.go:983` `storage.NewClient(…, storagens.CMS)`); their compose containers\n"
     "> are unfederated husks sitting off every storage path, and stay up only until platform **M810**.\n\n"
     "Callers push and pull binary objects through it instead of dealing with S3 themselves."),

    ("services/storage.md",
     "* **Upstream consumers**: jobsimulation (recordings, simulation documents), cms (content assets, media), app (user files, profile images)",
     "* **Upstream consumers**: **`app` only** — the jobsimulation domain (recordings, simulation documents),\n"
     "  the cms domain (content assets, media) and app itself (user files, profile images) all call from\n"
     "  inside the `backend` binary. The `jobsimulation`/`cms` husk containers call nothing (teardown M810)."),

    # ================= g3-B1 — the RED gene was FIXED at M219 =================
    ("services/clerkenstein.md",
     "> | **Go SDK** (`clerk-2.6.0`, M1) | **97.2% overall · 100% critical** — **26/27 genes**, 14 capabilities | Gate is ≥95 / =100 ⇒ **MET**. The 2.8% is **one deliberately RED gene** (see below). |",
     "> | **Go SDK** (`clerk-2.6.0`, M1) | **100% overall · 100% critical** — **27/27 genes**, 14 capabilities | Gate is ≥95 / =100 ⇒ **MET**. (Was 97.2% / 26-of-27 until M219 landed the org-eid fix — see below.) |"),

    ("services/clerkenstein.md",
     "> **The deliberately RED gene (M218 D16).** `MembershipOrgIdentity/real-org-eid` ships **failing, on\n"
     "> purpose**.",
     "> **The formerly-RED gene (M218 D16) — ✅ RESOLVED at M219.** `MembershipOrgIdentity/real-org-eid` shipped\n"
     "> **failing on purpose** for one milestone.",),

    ("services/clerkenstein.md",
     "> than **omit the field and keep a clean 100%**, the divergence is named in the report on **every single\n"
     "> run** until it lands. Routed forward as `FIX-M219-bapi-org-eid`.",
     "> than **omit the field and keep a clean 100%**, the divergence was named in the report on **every single\n"
     "> run** until it landed. **It has landed:** `clerk-backend/store.go:138` (`SeedOrgIdentity`) and `:151`\n"
     "> (`LookupOrgEid`) ship the real roster org UUID, and the DNA records it —\n"
     "> `alignment/dna/clerk-2.6.0.json:131`: *\"M219 landed the fix … taking the Go surface 97.2% -> 100%.\"*\n"
     "> The Go surface is **27/27**. `FIX-M219-bapi-org-eid` is CLOSED.",),

    # ================= g3-B2 — rc=2 now means REGRESSED =================
    ("services/clerkenstein.md",
     "the runner cannot build, exits **rc=2, with NO score**. | **Not** a pass. Routed forward as `TEST-M219-expressrun-dep-gate`. |",
     "the runner cannot build, exits **rc=3 (`ExitUnmeasurable`), with NO score**. | **Not** a pass. ⚠️ **rc=2 is\n"
     "> now `ExitRegressed` — a MEASURED regression** (`alignment/cmd/alignctl/run.go:134-135`); do not read a 2\n"
     "> as a missing Node module. |"),

    # ================= g3-B3 — the unbounded clerk-js fetch was fixed at M220 =================
    ("services/clerkenstein.md",
     "- **…and it is UNBOUNDED and UNCACHED — the proxy's real contract (documented in M218; it had never been\n"
     "  written down).** `clerk-frontend/server.go:187` fetches the bundle with a bare **`http.Get`**, which is\n"
     "  `http.DefaultClient` — i.e. **`Timeout: 0`, no timeout at all**. There is **no server-side cache**: the\n"
     "  only caching is a *response-side* `Cache-Control: public, max-age=3600` header (`:194`), so **every full\n"
     "  page load in a cold browser re-fetches from the CDN**, and the fake FAPI re-fetches from jsdelivr each\n"
     "  time. Consequences, in order of severity:",
     "- **…and it WAS unbounded and uncached until M220 — ✅ FIXED, kept here because the failure mode is still\n"
     "  worth recognising.** As documented at M218, `clerk-frontend/server.go` fetched the bundle with a bare\n"
     "  **`http.Get`** (`http.DefaultClient`, **`Timeout: 0`**) and held no server-side cache, so every cold page\n"
     "  load re-fetched from jsdelivr. **M220 closed it:** `clerk-frontend/server.go:35-67` now serves the\n"
     "  clerk-js bundle **from disk** with the CDN as a *bounded* fallback — `clerkJSFetchTimeout = 15s` on an\n"
     "  explicit `clerkJSClient` (commented *\"Explicitly NOT http.DefaultClient\"*), a disk cache at\n"
     "  `FAKE_FAPI_CLERKJS_CACHE`, and a test asserting no `http.Get(` survives on that path. **A slow or\n"
     "  blocked jsdelivr is therefore NO LONGER a plausible cause of a long demo login** — look elsewhere.\n"
     "  The consequences below describe the pre-M220 behaviour, in order of severity:"),

    # ================= g2-B3 — the legacy AI-readiness surface was DELETED =================
    ("services/ai-readiness.md",
     "> ⚠️ **There are TWO manager dashboards. Only one of them is the product.** Every AI-readiness demo pointer —\n"
     "> the cockpit deep-link catalog, the manager hero's `jump_to`, and the coverage sweep's page descriptor —\n"
     "> targeted the **legacy** one for four releases. Nothing ever failed, because the legacy page *does* render.\n"
     "> It just isn't the dashboard the product ships. **Establish which surface you are on before you conclude\n"
     "> anything about AI readiness.**",
     "> ⚠️ **HISTORICAL (M219) — there is now only ONE manager dashboard.** For four releases every AI-readiness\n"
     "> demo pointer (the cockpit deep-link catalog, the manager hero's `jump_to`, the coverage sweep's page\n"
     "> descriptor) targeted a **legacy** dashboard that still rendered but was not the shipped product. **That\n"
     "> surface no longer exists:** next-web-app `dae0fb2f7` (*\"drop orphaned container\"*, 2026-07-13) deleted\n"
     "> `AIReadinessContainer.tsx`, `AIReadinessIntro.tsx` and `AIReadinessView.tsx` (−653 lines), and\n"
     "> `/enterprise/workforce/ai-readiness` now **404s**. The section is kept because the *class* of defect —\n"
     "> an unlinked orphan surface that renders — is the one worth recognising, not because there is still a\n"
     "> choice to make."),

    ("services/ai-readiness.md",
     "| **Manager** | `AIReadinessContainer` → `AIReadinessView` — pre-v3.0 org-summary card + team table. **No cycle picker, no archetype matrix, no people, no How-we-measure, no What-to-do-next.** | `/enterprise/workforce/ai-readiness` | ❌ **LEGACY** |",
     "| **Manager** | `AIReadinessContainer` → `AIReadinessView` — pre-v3.0 org-summary card + team table. | `/enterprise/workforce/ai-readiness` | 🗑️ **DELETED** at next-web `dae0fb2f7` (2026-07-13); the route 404s |"),

    ("services/ai-readiness.md",
     "**How to tell them apart in code** (there is no `@deprecated` marker, no `-v2` naming, and no feature flag\n"
     "switching between them — the legacy one is simply *unlinked*):",
     "**How the orphan was identified, while it still existed** (there was no `@deprecated` marker, no `-v2`\n"
     "naming and no feature flag switching between them — the legacy one was simply *unlinked*). Retained as the\n"
     "recognition pattern; the anchors below are pre-`dae0fb2f7` and no longer resolve:"),

    ("services/ai-readiness.md",
     "- **The legacy route is an orphan**: no nav entry, no workforce tab (`WorkforceNewClient.tsx:125-151` omits it),",
     "- **The legacy route was an orphan**: no nav entry, no workforce tab (`WorkforceNewClient.tsx:125-151` omitted it),"),
]


def main() -> int:
    findings, applied = [], 0
    by_file: dict[str, list[tuple[str, str]]] = {}
    for fname, old, new in EDITS:
        by_file.setdefault(fname, []).append((old, new))

    staged: dict[Path, str] = {}
    for fname, pairs in by_file.items():
        path = ROOT / fname
        if not path.exists():
            findings.append(f"[missing] {fname}")
            continue
        text = path.read_text()
        for old, new in pairs:
            n = text.count(old)
            if n != 1:
                findings.append(
                    f"[anchor] {fname}: expected EXACTLY 1 occurrence, found {n} -> {old[:70]!r}")
                continue
            text = text.replace(old, new, 1)
            applied += 1
        staged[path] = text

    if not findings:
        for path, text in staged.items():
            path.write_text(text)

    if findings:
        print("SWEEP FAILED — no file written:", file=sys.stderr)
        for f in findings:
            print("  " + f, file=sys.stderr)
        return 1
    print(f"sweep33b: OK — {applied}/{len(EDITS)} edits applied across {len(by_file)} file(s)")
    return 0


if __name__ == "__main__":
    sys.exit(main())

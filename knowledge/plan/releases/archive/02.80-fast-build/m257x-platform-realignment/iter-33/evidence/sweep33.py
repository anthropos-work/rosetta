#!/usr/bin/env python3
"""sweep33.py — M257x iter-33 clause-5 corpus sweep, group 1.

Enumerated (file, old, new) tuples. `old` MUST occur EXACTLY ONCE in its file:
  0 occurrences -> the anchor moved (fail loudly)
  2+ occurrences -> the anchor is not unique (fail loudly)

Deliberately NOT idempotent: a re-run SHOULD fail with 0-occurrence errors.
That is the guard, not a defect (iter-22's harness shape).
"""
import sys
from pathlib import Path

ROOT = Path("/Users/marco/workspace/anthropos/rosetta/corpus/architecture")

EDITS = [
    # ---- external_services.md : BLOCKERS ----
    ("external_services.md",
     "NEXT_PUBLIC_GRAPHQL_ENDPOINT=http://localhost:8082/graphql/query   # was :5050/graphql",
     "NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT=http://localhost:8082/graphql/query   # was :5050/graphql\n"
     "# NB the var is WUNDERGRAPH, not GRAPHQL — `NEXT_PUBLIC_GRAPHQL_ENDPOINT` does not exist in\n"
     "# next-web-app. Set on the image at docker-compose.yml:352 (build arg) and :361 (runtime env)."),

    ("external_services.md",
     "  endpoint: process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT",
     "  endpoint: process.env.NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT"),

    ("external_services.md",
     "- Set up webhooks to production Sentinel endpoint",
     "- Set up Clerk webhooks to the production **backend** endpoint `/api/webhook/clerk` — **not**\n"
     "  Sentinel, which is authorization-only and exposes no webhook route"),

    ("external_services.md",
     "- Inspect Sentinel logs for sync errors",
     "- Inspect **backend** logs (`docker compose logs backend`) for `/api/webhook/clerk` errors — Clerk\n"
     "  user/org sync is app/backend's job (`app/internal/web/backend/backend.go:130`), not Sentinel's"),

    ("external_services.md",
     "### CMS Service Integration\n\nThe CMS service connects to Directus via:",
     "### cms-domain Directus integration\n\n"
     "> **⚠️ This is the cms DOMAIN inside `backend`, not the `cms` container.** Since cms-in-app the\n"
     "> Directus client lives at `app/internal/cms/directus/` and runs in-process in `backend`;\n"
     "> `app/cms_reader_switch.go` swaps the content reader to the in-process cms server, and\n"
     "> `app/main.go:971-973` makes `DIRECTUS_BASE_ADDR` a hard boot requirement **of `backend`**. The\n"
     "> `cms` container still starts until platform M810 but serves none of `backend`'s content reads.\n\n"
     "The cms domain connects to Directus via:"),

    ("external_services.md",
     "**Code Integration** (from CMS service):\n```go\n// internal/directus/",
     "**Code Integration** (`app/internal/cms/directus/`, compiled into `backend`):\n```go\n"
     "// app/internal/cms/directus/   (NOT the frozen cms repo's internal/directus/)"),

    # ---- external_services.md : minors ----
    ("external_services.md",
     "The Anthropos platform integrates with **three key external services**:",
     "The Anthropos platform integrates with **four key external services**:"),

    ("external_services.md",
     "that block (`:43-77` @",
     "that block (`:43-67` @"),

    ("external_services.md",
     "   - Configured in `studio-room/configs/*.ini`",
     "   - Configured in `anthropos-studio-room/configs/*.ini` (the repo is `anthropos-studio-room`;\n"
     "     it is baked into the `app` image and orchestrated from `app/internal/cms/studio/`)"),

    # ---- service_taxonomy.md ----
    ("service_taxonomy.md",
     "Frontend → CMS Service → Directus API (content.anthropos.work) → PostgreSQL",
     "Frontend/Studio-Desk → `backend` :8082/graphql/query (cms **domain**,\n"
     "`app/internal/cms/directus/`) → Directus API (content.anthropos.work) → PostgreSQL"),

    ("service_taxonomy.md",
     "The **CMS Service** acts as a smart proxy/adapter, adding business logic on top of Directus.",
     "The **cms domain inside `backend`** acts as a smart proxy/adapter, adding business logic on top of\n"
     "Directus. (Before cms-in-app this was a standalone `cms` service; that container still starts until\n"
     "platform M810 but no frontend reaches it — both are baked against `backend` at\n"
     "`docker-compose.yml:352`/`:361` and `:318`/`:334`.)"),

    ("service_taxonomy.md",
     "- **Studio-Room**: Direct integration with CMS service for blueprint retrieval",
     "- **Studio-Room**: runs inside the `app` image, orchestrated from `app/internal/cms/studio/` —\n"
     "  blueprint retrieval is in-process against the cms domain, not a call to a CMS service"),

    ("service_taxonomy.md",
     "- **Content Storage**: Directus API (via CMS proxy for core services)",
     "- **Content Storage**: Directus API (via the cms **domain** in `app`, for core services)"),

    ("service_taxonomy.md",
     "cp .env.example .env   # fill Clerk + AI keys",
     "cp .env.example .env.local   # fill Clerk + AI keys (the app reads code/.env.local)"),

    # ---- architecture_overview.md ----
    ("architecture_overview.md",
     "| **Next Web App** | Next.js 15 |",
     "| **Next Web App** | Next.js 16 |"),

    ("architecture_overview.md",
     "*   **Directus**: Proxied via CMS service (business logic layer)",
     "*   **Directus**: Proxied via the cms **domain** inside `backend` (business logic layer)"),

    ("architecture_overview.md",
     "    *   **Roadrunner**: Code execution proxy (via Judge0 sandbox)",
     "    *   **Roadrunner**: **orphaned husk** — the container still starts, but nothing calls it;\n"
     "        `backend` reaches Judge0 directly (`app/internal/jobsimwiring/wiring.go:118`)"),

    ("architecture_overview.md",
     "    *   **GraphQL/Cosmo Router**: API federation gateway",
     "    *   **GraphQL/Cosmo Router**: API federation gateway **(prod only — deleted from local dev at\n"
     "        platform `2adcf71`)**"),

    ("architecture_overview.md",
     "- **APIs**: GraphQL Federation v2 (WunderGraph Cosmo Router), gRPC/Connect-RPC (internal), Protocol Buffers",
     "- **APIs**: GraphQL Federation v2 (WunderGraph Cosmo Router — **prod only**; local dev talks to\n"
     "  `backend` directly), gRPC/Connect-RPC (internal), Protocol Buffers"),

    ("architecture_overview.md",
     "3. **External Services**: Clerk, Directus, GraphQL, AI providers, LiveKit, AWS Chime",
     "3. **External Services**: Clerk, Directus, GraphQL (**prod only**), AI providers, LiveKit, AWS Chime"),

    # ---- dependency_map.md ----
    ("dependency_map.md",
     "`docker-compose.yml:66-80`",
     "`docker-compose.yml:70-80`"),

    ("dependency_map.md",
     "inferred from configuration files (`docker-compose.yaml`)",
     "inferred from configuration files (`docker-compose.yml`)"),

    ("dependency_map.md",
     "| `backend` | App | CMS | User/org updates |",
     "| `backend` | App | App (cms **domain** in `app`; the `cms` husk also still subscribes until "
     "platform M810) | User/org updates |"),
]


def main() -> int:
    findings, applied = [], 0
    # group by file so each file is read once and written once
    by_file: dict[str, list[tuple[str, str]]] = {}
    for fname, old, new in EDITS:
        by_file.setdefault(fname, []).append((old, new))

    # PASS 1 — validate every anchor in every file. Nothing is written in this pass, so a
    # broken anchor in the LAST file cannot leave the FIRST file half-swept.
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

    # PASS 2 — write only if EVERY anchor in EVERY file validated.
    if not findings:
        for path, text in staged.items():
            path.write_text(text)

    if findings:
        print("SWEEP FAILED — no file written:", file=sys.stderr)
        for f in findings:
            print("  " + f, file=sys.stderr)
        return 1
    print(f"sweep33: OK — {applied}/{len(EDITS)} edits applied across {len(by_file)} file(s)")
    return 0


if __name__ == "__main__":
    sys.exit(main())

# Toolchain Overview

This document maps the essential tools required to interact with the Anthropos platform. It serves as a registry for what is used, why, and by whom, ensuring that tooling remains curated alongside the codebase.

## 1. Platform Setup & Infrastructure
Tools required to provision the environment and run the core infrastructure.

*   **Git**
    *   *Function*: Version Control & Source Code Management.
    *   *User*: All Engineers.
    *   *Context*: Cloning repositories (`rosetta`, `platform`, `app`, etc.) and managing code changes.

*   **Docker Desktop** (or Engine)
    *   *Function*: Container Runtime & Orchestration.
    *   *User*: All Engineers.
    *   *Context*: Hosting the local version of the platform (`platform/docker-compose.yml`), running DBs (Postgres, Redis) and services.

*   **XCode CLI Tools** (macOS) / **Build Essential** (Linux)
    *   *Function*: Native Compiler Toolchain (C/C++, Make).
    *   *User*: System / Backend Engineers.
    *   *Context*: Required by **Go** (CGO bindings) and **Makefiles**. Essential for `make setup` and `make gen`.

*   **Homebrew** (macOS) / **apt** (Linux)
    *   *Function*: System Package Manager.
    *   *User*: System.
    *   *Context*: Bootstrapping the initial environment (installing Git, Go, Node).

## 2. Service Development (Backend)
Tools specific to developing, building, and running the Go-based microservices.

*   **Go** (**v1.25+**)
    *   *Function*: Programming Language & Runtime.
    *   *User*: Backend Engineers.
    *   *Context*: Compiling and running service code locally, and — the binding constraint — the
        `rosetta-extensions` host tooling (`stacksecrets` / `stacksnap` / `stackseed`), whose six
        sections all declare `go 1.25.0` + `toolchain go1.25.12` and which build **on the host**.
        ⚠️ **This said `v1.23+` until M257x harden pass 60**, two minor versions low, and it named
        `cms` and `jobsimulation` as service code you compile — both are merged into `app` and neither
        repo is in the clone set. The service repos themselves declare `go 1.26.x` but build **inside
        Docker** on `golang:1.26-bookworm`, so they do not raise the host floor. Same derivation as
        [`corpus/ops/setup_guide.md`](../ops/setup_guide.md), which M257x iter-240 repaired; this
        fourth statement of the floor sat outside that iter's subject and stayed wrong.

*   **Make**
    *   *Function*: Task Runner & Build Automation.
    *   *User*: Backend Engineers.
    *   *Context*: Standardized interface for dev tasks. Used in service directories: `make setup` (deps), `make gen` (codegen), `make test`.

*   **Protoc / Gen Tools**
    *   *Function*: Code Generation.
    *   *User*: Backend Engineers.
    *   *Context*: *Managed via Make*. Generates gRPC/Protobuf definitions and boilerplate.

*   **PostgreSQL Client** (psql/GUI)
    *   *Function*: Database Interface.
    *   *User*: Backend Engineers.
    *   *Context*: Inspecting local database state exposed by Docker on forwarded ports.

*   **Atlas**
    *   *Function*: Database Schema Management.
    *   *User*: Backend Engineers.
    *   *Context*: **Required for Setup**. Manages PostgreSQL schema migrations — **two pipelines, both declared in `app`**: `env "local"` → the **`public`** schema (`app/atlas.hcl:6-19`, dir `terraform/migrations`) and `env "sentinel"` → the **`sentinel`** schema (`:50-64`, dir `terraform/migrations-sentinel`, `revisions_schema = "sentinel"`, added 2026-08-04). Used via `atlas migrate apply --env local` / `--env sentinel`. ⚠️ **This said `public, cms, jobsimulation, skillpath` until M257x iter-129** — `cms`/`jobsimulation`/`skillpath` are legacy husks whose repos are out of the clone set, and `repos.yml` @ platform `0c91421df` lists exactly one migrating repo (`app`, `schema: public`). A fresh stack never creates those three schemas.

## 3. Web Development (Frontend)
Tools specific to the Next.js monorepo and web applications.

*   **Node.js** (**v24+**)
    *   *Function*: JavaScript Runtime.
    *   *User*: Frontend Engineers.
    *   *Context*: Executing the dev server and build scripts. Derived from the highest
        `engines.node` in the clone set — `next-web-app` declares `">=24.0.0"`; `ant-academy` declares
        `">=22"`, which the higher floor already covers. ⚠️ **This said `v20+` until M257x harden pass
        60** — four major versions low, and low enough that `next-web-app` refuses to install.

*   **pnpm**
    *   *Function*: Package Manager.
    *   *User*: Frontend Engineers.
    *   *Context*: **Strictly required** (Corepack). Dependency management for the monorepo. Replacing `npm`/`yarn`.

*   **TurboRepo**
    *   *Function*: Monorepo Build System.
    *   *User*: Frontend Engineers.
    *   *Context*: *Internal dependency*. Orchestrates builds and caching within `next-web-app`.

## 4. Editor & Productivity
Recommended environment for efficiency.

*   **Visual Studio Code**
    *   *Function*: Integrated Development Environment (IDE).
    *   *User*: All Engineers.
    *   *Context*: Recommended editor. Configured with workspace settings for Go and ESLint/Prettier.

*   **Shell** (zsh/bash)
    *   *Function*: Command Line Interface.
    *   *User*: All Engineers.
    *   *Context*: Primary interface for all `git`, `make`, and `docker` commands.

## 5. Internal Applications & Experiments

Internal tools and sandboxes that support team workflows but are not part of the core platform.

*   **Anthropos Knowledge Base** (`anthropos-knowledge-base`)
    *   *Function*: Claude Code plugin providing product, technical, and design context.
    *   *User*: All Engineers.
    *   *Context*: Installed as a Claude Code plugin, gives Claude full Anthropos context (product details, architecture, design system, competitor analysis) when working in any Anthropos codebase. Includes skills like `/build-feature` and auto-triggered design system enforcement.
    *   *Setup*: Clone repo, then use `/plugin marketplace add` and `/plugin install` in Claude Code.
    *   ⚠️ **KNOWN CONTRADICTION — read before you trust its architecture context (M257x iter-125).**
        AKB carries a **second, parallel platform-architecture corpus** (six files under `knowledge/`,
        ≈1,773 lines) that **contradicts this one on the taxonomy figures**. It asserts *"60,000 skills
        … mapped to 18,000 roles"* in **14 places, citing no source in any of them**, and the figure is
        **load-bearing in four customer-facing competitor-comparison tables**. This corpus measured
        **42,790 public skills / 22,470 public job roles** (read-only production capture,
        `organization_id IS NULL`, 2026-06-29, reproducible against a live stack) — so **"18K roles" is
        REFUTED** (public ⊆ total, so 18K is below the floor) and **"60K skills" is UNVERIFIED**, not
        refuted. Derivation: [`shared_libraries.md § taxonomy figures`](../architecture/shared_libraries.md#taxonomy-figures).
        **Installing this plugin injects the refuted figure into your editor on every Anthropos repo** —
        which is why the warning lives on the install line rather than in a repo census nobody opens.
        **AKB is not simply less accurate: it was RIGHT and this corpus WRONG about the WunderGraph
        router's production residue**, because it reads the `infrastructure` repo this corpus had never
        cloned. Reconciliation is filed as `PLATFORM-M257x-akb-taxonomy-figures-contradict-measurement`
        in [`platform-defect-register.md`](../../knowledge/plan/platform-defect-register.md); the full
        comparison is [`org-repos.md` § 11](../architecture/org-repos.md).

*   **Anthropos Labs** (`experiments`)
    *   *Function*: Internal experiments hub for PoCs and prototypes.
    *   *User*: All Engineers.
    *   *Context*: Vanilla JS/HTML experiments with Clerk auth, hosted on Vercel. Used for UI prototyping, internal tools, and demos before platform integration.
    *   *Details*: [→ Anthropos Labs](./anthropos-labs.md)

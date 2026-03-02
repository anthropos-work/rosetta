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

*   **Go** (v1.23+)
    *   *Function*: Programming Language & Runtime.
    *   *User*: Backend Engineers.
    *   *Context*: Compiling and running service code (`app`, `cms`, `jobsimulation`) locally.

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
    *   *Context*: **Required for Setup**. Manages PostgreSQL schema migrations (`public`, `cms`, `jobsimulation`, `skiller`) and versioning. Used via `atlas migrate apply`.

## 3. Web Development (Frontend)
Tools specific to the Next.js monorepo and web applications.

*   **Node.js** (v20+)
    *   *Function*: JavaScript Runtime.
    *   *User*: Frontend Engineers.
    *   *Context*: Executing the dev server and build scripts.

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

*   **Anthropos Labs** (`experiments`)
    *   *Function*: Internal experiments hub for PoCs and prototypes.
    *   *User*: All Engineers.
    *   *Context*: Vanilla JS/HTML experiments with Clerk auth, hosted on Vercel. Used for UI prototyping, internal tools, and demos before platform integration.
    *   *Details*: [→ Anthropos Labs](./anthropos-labs.md)

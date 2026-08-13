# Anthropos Document Skill - Technical Reference

## Corpus Structure

**90 markdown files.** This tree is a shape, **not an inventory** — every section has a maintained
`README.md` index, and `ls` is the source of truth. (An earlier revision listed 13 files and omitted the
entire `ops/demo/` family plus the fenced migration map; sweeps that trusted it missed both.)

```
corpus/                              # 90 .md total — list, don't recall
├── README.md                        # corpus entry point
├── architecture/  (11)              # README.md is the index
│   ├── platform-migration-status.md # ⭐ FENCED one-row-per-service map — read FIRST for
│   │                                #    "does this service still exist"; outranks prose
│   │                                #    anywhere, incl. CLAUDE.md. Guarded vs repos.yml.
│   ├── architecture_overview.md · service_taxonomy.md · dependency_map.md
│   ├── frontend_architecture.md · external_services.md · shared_libraries.md
│   └── security_compliance.md · ai_architecture.md · alignment_testing.md
├── services/      (29)              # README.md = enumerated index of every service doc
│   ├── TEMPLATE.md                  # follow this for a new service
│   ├── backend.md · sentinel.md · next-web-app.md · studio-desk.md · studio-room.md · …
│   └── archived/merged redirects: skiller.md · skillpath.md · chronos.md · intelligence.md
├── ops/           (23)              # README.md is the index
│   ├── setup_guide.md · run_guide.md · update_guide.md · quick_ops.md
│   ├── platform-alignment.md · safety.md · verification.md · idempotency.md
│   ├── secrets-spec.md · seeding-spec.md · snapshot-spec.md · snapshot-cold-start.md
│   ├── db-access.md · directus-local.md · rosetta_demo.md · webhook_setup.md
│   ├── staging-bringup.md · staging-clerk.md · staging-sync.md · staging_from_dump.md
│   └── demo/  (23)                  # the demo family — README.md is the index
│       ├── demo-up-defaults.md      # fenced against the parsers in both directions
│       ├── build-budget.md · latency-budget.md · coverage-protocol.md · playthroughs.md
│       └── recipe-*.md · *-spec.md · content-stories-routes.md · frontend-tier.md · …
└── tools/         (3)
    └── README.md · toolchain_overview.md · anthropos-labs.md
```

## Evidence Analysis Patterns

### New Go Service

**Indicators**: `go.mod`, `main.go` or `cmd/`, Dockerfile with golang base

**Key files to read**:
- `go.mod` - Module name and dependencies
- `main.go` or `cmd/main.go` - Entry point
- `rpc.go` - API definitions
- `internal/data/ent/` - Database schema

**Corpus impact**:
- Create: `corpus/services/{name}.md`
- Update: `architecture_overview.md`, `service_taxonomy.md`
- Update: `setup_guide.md` if new clone/migration steps

### New TypeScript Service

**Indicators**: `package.json`, `tsconfig.json`, Vite/Next.js config

**Key files to read**:
- `package.json` - Name, scripts, dependencies
- Entry point (pages/, src/index.ts)
- API routes if present

**Corpus impact**:
- Create: `corpus/services/{name}.md`
- Update: `frontend_architecture.md` if frontend
- Update: `toolchain_overview.md` if new build tool

### New Python Service

**Indicators**: `requirements.txt` or `pyproject.toml`, Python files

**Key files to read**:
- `requirements.txt` - Dependencies
- Entry point (main.py, app.py, gen.py)
- Check for FastAPI/Flask or AI libraries

**Corpus impact**:
- Create: `corpus/services/studio-{name}.md` (usually Studio tier)
- Update: `service_taxonomy.md`
- Update: `setup_guide.md` for pip install steps

### Setup Feedback

**Indicators**: `setup_progress.md` with checkboxes and notes

**Analysis**:
- Parse `[ ]` (incomplete) items
- Parse error notes
- Categorize: missing steps, incorrect commands, OS-specific

**Corpus impact**:
- Update: `setup_guide.md` with fixes
- Add troubleshooting entries

## Tool Usage

### Discovery

```bash
# Find files by pattern
Glob: **/*.go           # Go files
Glob: **/docker-compose*.yml

# Search content
Grep: "func main", type: go    # Find entrypoints
Grep: "FROM.*golang", glob: Dockerfile*

# Read key files
Read: go.mod, README.md, package.json
```

### Determine Service Tier

| Indicator | Tier |
|-----------|------|
| Go service in docker-compose | Core |
| TypeScript with Next.js | Frontend |
| Python with AI libraries | Studio |
| Third-party SaaS | External |

### Writing Documentation

```bash
# Create new service doc
Write: corpus/services/new-service.md
# Follow corpus/services/TEMPLATE.md structure

# Update existing doc
Edit: corpus/architecture/architecture_overview.md
# Add entry to appropriate section
```

## TodoWrite Checklist Pattern

When documenting new evidence, create a checklist like:

```typescript
TodoWrite({
  todos: [
    { content: "Inspect evidence source", status: "in_progress", activeForm: "Inspecting evidence" },
    { content: "Check architecture_overview.md", status: "pending", activeForm: "Checking architecture_overview" },
    { content: "Check service_taxonomy.md", status: "pending", activeForm: "Checking service_taxonomy" },
    { content: "Check dependency_map.md", status: "pending", activeForm: "Checking dependency_map" },
    { content: "Create/update service doc", status: "pending", activeForm: "Updating service doc" },
    { content: "Check setup_guide.md", status: "pending", activeForm: "Checking setup_guide" },
    { content: "Check Claude skills", status: "pending", activeForm: "Checking Claude skills" },
    { content: "Verify discoverability", status: "pending", activeForm: "Verifying discoverability" }
  ]
})
```

Mark each as completed after reviewing (even if no changes needed).

## Documentation Quality Patterns

### Dual-Level Structure

```markdown
# Service Name

## Role & Responsibility
[High-level: 1-2 sentences for PMs]
- Primary Goal: What problem does it solve?
- Key Functions: Bullet list of capabilities

## Architecture & Code Map
[Deep dive: For engineers]
- Codebase location
- Language/framework
- Key directories explained
```

### Command Blocks

Always include verification:

```markdown
### Step: Install Tool

```bash
brew install tool-name
```

**Verify**:
```bash
tool-name --version
# Expected: v1.2.3 or higher
```
```

### Link to Parent Docs

New content must be linked from at least one parent:
- Service docs → linked from `architecture_overview.md` or `service_taxonomy.md`
- Setup steps → linked from `setup_guide.md` table of contents
- Tools → linked from `toolchain_overview.md`

## Error Recovery

| Situation | Action |
|-----------|--------|
| Evidence unclear | Ask user for clarification |
| Conflicts with existing docs | Flag for user decision |
| Scope expansion (found more undocumented things) | Document original evidence, note discoveries for follow-up |
| Can't access evidence | Ask user for access or path |

## Integration with Other Skills

| Scenario | Action |
|----------|--------|
| New setup requirements | Update `setup_guide.md` → suggest `/dev-up` to test |
| New run requirements | Update `run_guide.md` → update the `dev-up` skill if needed |
| Found issues during documentation | Create ops report for `/update-knowledge` follow-up |

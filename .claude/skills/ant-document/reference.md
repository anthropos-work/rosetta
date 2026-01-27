# Anthropos Document Skill - Technical Reference

## Corpus Structure

```
corpus/
├── README.md                    # Corpus entry point
├── architecture/
│   ├── architecture_overview.md # System design (update for new services)
│   ├── service_taxonomy.md      # Three-tier categorization
│   ├── frontend_architecture.md # Next.js monorepo details
│   ├── external_services.md     # Clerk, Directus, GraphQL
│   └── dependency_map.md        # Service interconnections
├── services/
│   ├── TEMPLATE.md              # Follow this for new services
│   ├── backend.md, sentinel.md, etc.
│   └── studio-desk.md, studio-room.md
├── ops/
│   ├── setup_guide.md           # Environment setup
│   ├── run_guide.md             # Platform startup
│   └── update_guide.md          # Sync and update
└── tools/
    └── toolchain_overview.md    # Development tools
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
| New setup requirements | Update `setup_guide.md` → suggest `/ant-setup` to test |
| New run requirements | Update `run_guide.md` → update `ant-run` skill if needed |
| Found issues during documentation | Create ops report for `/ant-integrate` follow-up |

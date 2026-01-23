# Anthropos Setup Skill - Technical Reference

## Architecture

This skill implements a **setup-and-improve** loop:
1. Execute setup steps from corpus documentation
2. Observe real-world execution results
3. Update documentation based on observations
4. Continue execution with improved docs

## File References

### Documentation Sources
- `corpus/setup/setup_guide.md` - Master setup guide
- `corpus/setup/setup_checklist.md` - Progress tracking checklist
- `corpus/architecture/` - Platform architecture context

### Working Directory
- `anthropos-dev/` - Git-ignored scratchpad for all setup activities

## Tool Usage Strategy

### Read-Only Tools (Verification)
- `Bash` for verification commands (`--version`, `docker ps`, etc.)
- `Read` for reading documentation and .env examples
- `Glob` for finding configuration files

### Write Tools (Execution & Documentation)
- `Bash` for installation and cloning (after user confirmation)
- `Edit` for updating documentation (auto-improvement)
- `Write` for creating .env files

### Interactive Tools
- `AskUserQuestion` for all confirmations before:
  - Installing system tools
  - Cloning repositories
  - Starting services
  - Populating secrets

## STEP RUN Guidelines Implementation

### Guideline 1: Verify Before Install
```bash
# Pattern: Check existence first
command -v tool_name || echo "Not installed"
tool_name --version 2>/dev/null || echo "Not installed"
```

### Guideline 2: Request Confirmation
```typescript
// Pattern: Always ask before destructive operations
AskUserQuestion({
  question: "Install Docker Desktop? This will download ~500MB.",
  header: "Install Tool",
  options: [
    { label: "Yes, install", description: "Proceed with installation" },
    { label: "Skip for now", description: "I'll install manually" }
  ]
})
```

### Guideline 3: Verify After Install
```bash
# Pattern: Same command as "verify before"
tool_name --version
# Exit code 0 = success
```

### Guideline 4: Document Improvements
```markdown
# Pattern: Update guide immediately
Edit({
  file_path: "corpus/setup/setup_guide.md",
  old_string: "1. Install Docker Desktop",
  new_string: "1. Install Docker Desktop\n   * Verification: `docker --version && docker compose version`"
})
```

### Guideline 5: Track Progress in Local Checklist
```markdown
# Pattern: Update LOCAL copy as you progress
Edit({
  file_path: "anthropos-dev/setup_progress.md",  # LOCAL working copy
  old_string: "- [ ] **Docker Desktop / Engine** installed & running",
  new_string: "- [x] **Docker Desktop / Engine** installed & running"
})

# Only update ORIGINAL when setup guide structure changes
# (new step, removed step, reordered steps)
Edit({
  file_path: "corpus/setup/setup_checklist.md",  # ORIGINAL - only for structure changes
  old_string: "## 3. Cloning Repositories",
  new_string: "## 3. GitHub SSH Access\n\n- [ ] SSH key generated\n\n## 4. Cloning Repositories"
})
```

## Docker Isolation Strategy

The skill uses `-p anthropos-rosetta` project name to create isolated Docker resources:

**Containers**: `anthropos-rosetta-{service}-1`
**Networks**: `anthropos-rosetta_app-network`
**Volumes**: `anthropos-rosetta_postgres_data`

This prevents conflicts with other Anthropos environments but may cause **port conflicts**. The skill should detect these and guide resolution.

## Error Recovery Patterns

### Pattern 1: Tool Not Found
```yaml
Error: command not found: pnpm
Recovery:
  1. Check if npm is installed
  2. Run: corepack enable
  3. Verify: pnpm --version
  4. Document in troubleshooting
```

### Pattern 2: Permission Denied (Docker on Linux)
```yaml
Error: permission denied while connecting to Docker daemon
Recovery:
  1. Check: groups | grep docker
  2. Fix: sudo usermod -aG docker $USER
  3. Reload: newgrp docker
  4. Verify: docker ps
  5. Update setup_guide.md with emphasis
```

### Pattern 3: Port Already in Use
```yaml
Error: bind: address already in use
Recovery:
  1. Identify: lsof -i :3000
  2. Options:
     a. Stop conflicting service
     b. Change port in config
  3. Document resolution
```

### Pattern 4: Missing Environment Variable
```yaml
Error: CLERK_SECRET_KEY is not set
Recovery:
  1. Check: .env file exists
  2. Verify: grep CLERK_SECRET_KEY .env
  3. Guide: Populate from 1Password
  4. Emphasize in setup_guide.md
```

## Auto-Improvement Examples

### Example 1: Add Missing Verification
```markdown
Before:
> 4. **Go** (v1.23+): `brew install go`

After:
> 4. **Go** (v1.23+): `brew install go`
>    * *Verification*: `go version`
```

### Example 2: Add Troubleshooting Entry
```markdown
Before:
> ## 9. Troubleshooting
> (existing entries)

After:
> ## 9. Troubleshooting
> (existing entries)
>
> ### "corepack: command not found"
> If pnpm installation fails, ensure Node.js v16+ is installed.
> * **Fix**: `npm install -g corepack && corepack enable`
```

### Example 3: Clarify Ambiguous Step
```markdown
Before:
> 2. **Populate secrets**: Edit `platform/.env`

After:
> 2. **Populate secrets**: Edit `platform/.env` and fill in all required secret values from 1Password vault "Engineering/Env".
>
>    **Critical Keys Required**:
>    * `CLERK_SECRET_KEY` & `CLERK_PUBLISHABLE_KEY` (Auth)
>    * `OPENAI_API_KEY` (AI services)
```

## Progress Tracking

### Two-Level Tracking System

**Level 1: Local Checklist** (`anthropos-dev/setup_progress.md`)
- Copied from `corpus/setup/setup_checklist.md` at the start
- Updated as each step completes
- Used for resuming interrupted setup
- Contains "Notes / Errors" table for issue reporting
- This is the USER'S working copy

**Level 2: TodoWrite Tool**
- High-level phase tracking
- Immediate next steps
- Blockers requiring user input

Example TodoWrite tracking:
```markdown
- [x] Phase 0: Copied checklist to anthropos-dev/
- [x] Phase 1: Prerequisites - Git verified
- [x] Phase 1: Prerequisites - Docker verified
- [ ] Phase 1: Prerequisites - Go verification
- [ ] Phase 2: GitHub SSH access
- [ ] Documentation: Add verification for Go
```

### When to Update Original Checklist

Only update `corpus/setup/setup_checklist.md` when:
- Adding a NEW step to setup_guide.md (structure change)
- Removing a step from setup_guide.md (structure change)
- Reordering steps in setup_guide.md (structure change)

Do NOT update original checklist when:
- Clarifying existing step details
- Adding verification commands
- Fixing typos or improving descriptions
- Adding troubleshooting information

## Success Validation

At the end of setup, the skill should verify:

1. **Services Running**: `docker ps` shows healthy containers
2. **Frontend Accessible**: `curl -s http://localhost:3000` returns HTML
3. **Studio-Desk Accessible**: `curl -s http://localhost:3100` returns HTML
4. **Environment Valid**: All required keys present in .env files
5. **Documentation Updated**: Git shows modifications to setup docs

## Repository Structure Post-Setup

```
anthropos-dev/
├── platform/
│   ├── docker-compose.yml
│   └── .env (secrets)
├── backend/ (cloned from app)
├── cms/
├── jobsimulation/
├── next-web-app/
└── studio/
    ├── studio-desk/
    └── studio-room/
```

## Skill Invocation

```bash
# Full setup from scratch
/anthropos-setup

# Run complete setup
/anthropos-setup full

# Jump to specific phase
/anthropos-setup repos
/anthropos-setup docker
/anthropos-setup frontend
```

## Integration with Project Rosetta

This skill embodies the **Recursive Inspection** objective:
- It uses the corpus to execute setup
- It observes real execution to improve the corpus
- It updates the corpus for future agents/engineers
- It validates the "Recreation Standard" criterion

The skill is both a **consumer** and **improver** of the documentation corpus.

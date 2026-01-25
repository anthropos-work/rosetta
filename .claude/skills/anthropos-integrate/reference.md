# Anthropos Integrate Skill - Technical Reference

## Architecture

This skill implements an **inspect-plan-implement** loop:
1. Identify new evidence about the Anthropos platform
2. Inspect and analyze the evidence
3. Plan documentation updates
4. Implement changes to the corpus
5. Verify documentation quality

## Evidence Type Decision Tree

```
User invokes /anthropos-integrate
    │
    ├─ Argument provided?
    │   ├─ Yes → Parse evidence type (A/B/C/D/E/F)
    │   └─ No → AskUserQuestion for evidence type
    │
    ├─ Type A: New Repo
    │   ├─ Clone to anthropos-dev/
    │   ├─ Read README, package.json/go.mod, entrypoint
    │   ├─ Determine tier (Core/Studio/External)
    │   └─ Plan: service doc + architecture updates
    │
    ├─ Type B: New Feature
    │   ├─ Locate code with Glob/Grep
    │   ├─ Read implementation and tests
    │   ├─ Map affected services
    │   └─ Plan: service doc updates + possible architecture
    │
    ├─ Type C: New Directory
    │   ├─ List contents
    │   ├─ Identify what was added
    │   └─ Route to Type A or B handling
    │
    ├─ Type D: Setup Feedback
    │   ├─ Read setup_progress.md
    │   ├─ Parse checkboxes and notes
    │   ├─ Categorize: missing/incorrect/env-specific
    │   └─ Plan: setup guide + checklists + troubleshooting
    │
    ├─ Type E: Missing Aspect
    │   ├─ AskUserQuestion for specifics
    │   ├─ Locate in platform code
    │   └─ Plan based on what's found
    │
    └─ Type F: Other
        ├─ AskUserQuestion for details
        └─ Custom handling
```

## File References

### Documentation Targets

| Corpus Location | Purpose | Update Trigger |
|-----------------|---------|----------------|
| `corpus/architecture/architecture_overview.md` | System design | New service, major architectural change |
| `corpus/architecture/service_taxonomy.md` | Tier classification | New service added |
| `corpus/architecture/frontend_architecture.md` | Frontend details | Frontend changes |
| `corpus/architecture/external_services.md` | Third-party integrations | New external service |
| `corpus/architecture/dependency_map.md` | Service dependencies | Dependency changes |
| `corpus/services/{service}.md` | Per-service docs | New or updated service |
| `corpus/tools/toolchain_overview.md` | Development tools | New tool required |
| `corpus/ops/platform-setup/setup_guide.md` | Setup instructions | New setup step |
| `corpus/ops/platform-setup/setup_checklist_macos.md` | macOS progress | Setup structure change |
| `corpus/ops/platform-setup/setup_checklist_linux.md` | Linux progress | Setup structure change |

### Claude Assets

| Asset | Purpose | Update Trigger |
|-------|---------|----------------|
| `.claude/skills/anthropos-setup/SKILL.md` | Setup automation | New setup phase |
| `.claude/skills/anthropos-integrate/SKILL.md` | This skill | Improve integration process |
| `.claude/settings.local.json` | Permissions | New command patterns needed |
| `CLAUDE.md` | Agent context | Significant platform changes |

### Working Directory

- `anthropos-dev/` - Git-ignored scratchpad for all evidence analysis

## Tool Usage Strategy

### Discovery Tools

```typescript
// Find files by pattern
Glob({ pattern: "**/*.go" })  // Find Go files
Glob({ pattern: "**/docker-compose*.yml" })  // Find Docker configs

// Search content
Grep({ pattern: "func main", type: "go" })  // Find entrypoints
Grep({ pattern: "FROM.*golang", glob: "Dockerfile*" })  // Find Go Dockerfiles

// Read files
Read({ file_path: "anthropos-dev/new-repo/README.md" })
```

### Analysis Patterns

#### Determine Service Tier

```typescript
// Check for indicators
Read({ file_path: "go.mod" })  // Go service → likely Core
Read({ file_path: "package.json" })  // Check for "next" → Frontend
Read({ file_path: "requirements.txt" })  // Python → likely Studio

// Check Docker integration
Grep({ pattern: "service-name", path: "anthropos-dev/platform/docker-compose.yml" })
```

#### Map Dependencies

```typescript
// Go imports
Grep({ pattern: "anthropos-work/", path: "go.mod" })

// NPM dependencies
Read({ file_path: "package.json" })  // Look at dependencies object

// Docker dependencies
Grep({ pattern: "depends_on:", path: "docker-compose.yml", "-A": 5 })
```

### Writing Tools

```typescript
// Create new service doc
Write({
  file_path: "corpus/services/new-service.md",
  content: "# New Service\n\n## Role & Responsibility\n..."
})

// Update existing doc
Edit({
  file_path: "corpus/architecture/architecture_overview.md",
  old_string: "## Core Services\n\n-",
  new_string: "## Core Services\n\n- **New Service**: Description\n-"
})
```

### Interactive Tools

```typescript
// Ask for evidence type
AskUserQuestion({
  questions: [{
    question: "What type of evidence would you like to integrate?",
    header: "Evidence",
    options: [
      { label: "New Repo/Tool", description: "A new repository or development tool" },
      { label: "New Feature", description: "New code added to existing service" },
      { label: "Setup Feedback", description: "Issues found during setup" },
      { label: "Missing Docs", description: "Platform aspect not in corpus" }
    ],
    multiSelect: false
  }]
})

// Confirm update plan
AskUserQuestion({
  questions: [{
    question: "Proceed with this update plan?",
    header: "Approve",
    options: [
      { label: "Yes, proceed", description: "Implement all planned changes" },
      { label: "Modify plan", description: "I want to adjust something" },
      { label: "Cancel", description: "Don't make changes" }
    ],
    multiSelect: false
  }]
})
```

## Evidence Analysis Patterns

### Pattern 1: New Go Service

```yaml
Indicators:
  - go.mod present
  - main.go or cmd/ directory
  - Dockerfile with golang base
  - Present in docker-compose.yml

Analysis Steps:
  1. Read go.mod for module name and dependencies
  2. Find main entrypoint
  3. Look for rpc.go or api definitions
  4. Check for internal/ structure
  5. Find migrations (atlas.hcl, migrations/)

Documentation Impact:
  - Create: corpus/services/{name}.md
  - Update: corpus/architecture/architecture_overview.md
  - Update: corpus/architecture/service_taxonomy.md
  - Update: corpus/ops/platform-setup/setup_guide.md (clone + migrations)
  - Update: Both setup checklists
```

### Pattern 2: New TypeScript Service

```yaml
Indicators:
  - package.json present
  - tsconfig.json present
  - Vite or Next.js config

Analysis Steps:
  1. Read package.json for name, scripts, dependencies
  2. Determine framework (Next.js, Vite, Express)
  3. Find entry point (pages/, src/index.ts)
  4. Check for API routes

Documentation Impact:
  - Create: corpus/services/{name}.md
  - Update: corpus/architecture/frontend_architecture.md (if frontend)
  - Update: corpus/tools/toolchain_overview.md (if new build tool)
```

### Pattern 3: New Python Service

```yaml
Indicators:
  - requirements.txt or pyproject.toml
  - *.py files
  - Dockerfile with python base

Analysis Steps:
  1. Read requirements.txt for dependencies
  2. Find entrypoint (main.py, app.py, cli.py)
  3. Check for FastAPI/Flask indicators
  4. Look for AI/ML libraries (openai, anthropic, torch)

Documentation Impact:
  - Create: corpus/services/studio-{name}.md (usually Studio tier)
  - Update: corpus/architecture/service_taxonomy.md
  - Update: corpus/ops/platform-setup/setup_guide.md (pip install steps)
```

### Pattern 4: Setup Feedback

```yaml
Indicators:
  - setup_progress.md exists
  - Contains checkboxes [ ] and [x]
  - Has "Notes / Errors" section

Analysis Steps:
  1. Parse all [ ] (incomplete) items
  2. Parse all error notes
  3. Cross-reference with setup_guide.md
  4. Categorize issues:
     - Missing verification commands
     - Incorrect commands
     - Missing steps entirely
     - OS-specific issues

Documentation Impact:
  - Update: corpus/ops/platform-setup/setup_guide.md
  - Update: Troubleshooting section
  - Update: Checklists if structure changed
```

## Documentation Quality Enforcement

### Pre-Write Checklist

Before creating/editing documentation:

```typescript
// Verify purpose is clear
const hasRoleSection = content.includes("## Role") || content.includes("## Purpose")

// Verify commands are complete
const hasPlaceholders = content.match(/\[YOUR_.*?\]/)  // Should be false
const hasVerification = content.includes("*Verification*") || content.includes("verify")

// Verify accessibility
const averageSentenceLength = sentences.map(s => s.split(' ').length).avg()
// Target: < 20 words average
```

### Post-Write Validation

After creating documentation:

```typescript
// Check all links resolve
const links = content.match(/\[.*?\]\((.*?)\)/g)
for (const link of links) {
  // Verify file exists
  Glob({ pattern: extractPath(link) })
}

// Check commands are runnable (if in anthropos-dev/)
const codeBlocks = content.match(/```bash\n([\s\S]*?)```/g)
// Flag any that look incomplete
```

## Update Plan Generation

### Template Variables

```typescript
const planTemplate = `
## Update Plan for: {{EVIDENCE_NAME}}

### Summary
{{ONE_SENTENCE_DESCRIPTION}}

### Evidence Type: {{TYPE_LETTER}}

### Evidence Location
- Path: {{EVIDENCE_PATH}}
- Repo: {{REPO_URL}} (if applicable)

### Analysis Summary
{{KEY_FINDINGS}}

### Files to Create
{{#EACH NEW_FILES}}
- [ ] \`{{PATH}}\` - {{PURPOSE}}
{{/EACH}}

### Files to Update
{{#EACH UPDATE_FILES}}
- [ ] \`{{PATH}}\` - {{CHANGE_DESCRIPTION}}
{{/EACH}}

### Claude Assets
{{#EACH CLAUDE_ASSETS}}
- [ ] \`{{PATH}}\` - {{CHANGE_DESCRIPTION}}
{{/EACH}}

### Verification Steps
{{VERIFICATION_STEPS}}

### Estimated Changes
- New files: {{NEW_COUNT}}
- Updated files: {{UPDATE_COUNT}}
- Lines of documentation: ~{{ESTIMATED_LINES}}
`
```

## Error Recovery Patterns

### Pattern 1: Can't Clone Repository

```yaml
Error: Permission denied (publickey)
Recovery:
  1. Check SSH key: ssh -T git@github.com
  2. If not authenticated, prompt user
  3. Option: Ask user to clone manually
  4. Continue with inspection once cloned
```

### Pattern 2: Evidence Not Found

```yaml
Error: File/directory doesn't exist
Recovery:
  1. Verify path with user
  2. Search for similar names: Glob({ pattern: "**/partial-name*" })
  3. Ask user for correct location
```

### Pattern 3: Conflicting Documentation

```yaml
Scenario: New evidence conflicts with existing docs
Recovery:
  1. Flag the conflict to user
  2. Present both versions
  3. Ask user which is correct
  4. Document resolution in corpus
```

### Pattern 4: Scope Expansion

```yaml
Scenario: Evidence reveals more undocumented aspects
Recovery:
  1. Document only the original evidence
  2. Note discovered gaps
  3. Create follow-up tasks (TodoWrite)
  4. Ask user if they want to continue with discoveries
```

## Progress Tracking

### TodoWrite Integration

```typescript
// Initial todos
TodoWrite({
  todos: [
    { content: "Identify evidence target", status: "in_progress", activeForm: "Identifying evidence target" },
    { content: "Inspect evidence", status: "pending", activeForm: "Inspecting evidence" },
    { content: "Create update plan", status: "pending", activeForm: "Creating update plan" },
    { content: "Get plan approval", status: "pending", activeForm: "Getting plan approval" },
    { content: "Implement updates", status: "pending", activeForm: "Implementing updates" },
    { content: "Verify documentation", status: "pending", activeForm: "Verifying documentation" }
  ]
})

// Update as phases complete
TodoWrite({
  todos: [
    { content: "Identify evidence target", status: "completed", activeForm: "Identifying evidence target" },
    { content: "Inspect evidence", status: "in_progress", activeForm: "Inspecting evidence" },
    // ...expand with specific implementation tasks
  ]
})
```

## Success Validation

At completion, verify:

1. **Plan Executed**: All planned files created/updated
2. **Links Valid**: All markdown links resolve
3. **Commands Work**: Bash commands are syntactically correct
4. **Discoverable**: New content linked from corpus/README.md or parent docs
5. **Consistent Style**: Follows dual-level documentation pattern
6. **Verified**: User confirms accuracy

## Skill Invocation

```bash
# Interactive - asks for evidence type
/anthropos-integrate

# Direct evidence type
/anthropos-integrate A              # New repo
/anthropos-integrate B              # New feature
/anthropos-integrate C              # New directory
/anthropos-integrate D              # Setup feedback
/anthropos-integrate E              # Missing aspect

# With path argument
/anthropos-integrate A studio-analytics
/anthropos-integrate D anthropos-dev/setup_progress.md
```

## Integration with Other Skills

### Handoff to anthropos-setup

If integration reveals new setup requirements:
1. Document in setup_guide.md
2. Update checklists
3. Suggest user run `/anthropos-setup` to test changes

### Recursive Improvement

This skill can improve itself:
1. After integration, review what worked/didn't
2. Update SKILL.md with better patterns
3. Update reference.md with new scenarios

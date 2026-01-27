---
name: ant-integrate
description: Integrate new evidence into the Rosetta corpus through reflection and analysis
argument-hint: [evidence-type] or interactive
---

# Anthropos Corpus Integration

Analyze new evidence about the Anthropos platform and update the Rosetta documentation corpus through structured reflection and implementation.

## Your Mission

You are a **documentation integrator** for Project Rosetta. When new elements of the Anthropos platform emerge (repos, features, tools, setup feedback), your job is to:

1. **Identify the evidence** - What new information exists?
2. **Inspect and understand** - What does it do? How does it fit?
3. **Plan corpus updates** - What docs need changing?
4. **Implement updates** - Write clear, accessible documentation

## Modus Operandi (From Project Rosetta)

Follow these core principles:

1. **Iterative & Goal-Oriented**: One clear goal at a time. Don't boil the ocean.
2. **Autoconsistent & Discoverable**: A new agent or engineer should land here and know exactly what to do.
3. **The "Recreation Standard"**: Can someone recreate the dev environment from this documentation?
4. **Dual-Level Documentation**: High-level for PMs + Deep dive for engineers.

## Phase 1: Identify Evidence Target

If no target is provided, **ask the user** which type of evidence to integrate:

### Evidence Types

| Type | Description | Example |
|------|-------------|---------|
| **A - New Repo/Project/Tool** | A new repository in anthropos-work org or a new tool | `studio-analytics`, new CLI tool |
| **B - New Feature/Code** | Significant addition to existing platform | New API endpoint, new service method |
| **C - New Directory** | New folder appeared in anthropos-dev | Cloned repo, generated output |
| **D - Setup Feedback** | A `setup_progress.md` with issues/gaps | Missing steps, failed commands |
| **E - Missing Aspect** | Known platform element not in corpus | Undocumented service, missing architecture |
| **F - Other** | User-specified evidence | External doc, Slack thread, etc. |

Use AskUserQuestion to clarify:
```
What evidence would you like to integrate into the corpus?
- A: New repository, project, or tool
- B: New feature or code in existing service
- C: New directory in anthropos-dev/
- D: Setup feedback (setup_progress.md with issues)
- E: Missing platform aspect
- Other: (describe)
```

## Phase 2: Inspect Evidence

### For Repos/Projects (Type A, C)

1. **Clone if needed**: Clone to `anthropos-dev/` workspace
2. **Read key files**:
   - `README.md` - Purpose and usage
   - `package.json` / `go.mod` / `requirements.txt` - Tech stack
   - `docker-compose.yml` or `Dockerfile` - Deployment pattern
   - Entry points (`main.go`, `index.ts`, etc.)
3. **Identify role**: What tier does this belong to? (Core, Studio, External)
4. **Map dependencies**: What does it depend on? What depends on it?

### For Features/Code (Type B)

1. **Locate the code**: Use Glob/Grep to find the new functionality
2. **Read implementation**: Understand what it does
3. **Check tests**: Are there test files showing expected behavior?
4. **Map impact**: Which services/docs are affected?

### For Setup Feedback (Type D)

1. **Read setup_progress.md**: Look for `[x]` completed, `[ ]` failed, and notes
2. **Identify gaps**: What steps failed? What was missing?
3. **Categorize issues**:
   - Missing verification commands
   - Incorrect instructions
   - Missing steps entirely
   - Environment-specific issues (macOS vs Linux)

### For Missing Aspects (Type E)

1. **User interview**: Ask what aspect is missing
2. **Locate in platform**: Find the relevant code/config
3. **Understand context**: Why wasn't it documented? What's its role?

## Phase 3: Plan Corpus Updates

Create an **Update Plan** before making changes. Consider:

### Documentation Locations

| Element | Where to Document |
|---------|-------------------|
| New service (Core) | `corpus/services/{service-name}.md` + update `architecture_overview.md` |
| New service (Studio) | `corpus/services/studio-{name}.md` + update `service_taxonomy.md` |
| New external integration | `corpus/architecture/external_services.md` |
| New setup step | `corpus/ops/setup_guide.md` |
| New tool | `corpus/tools/toolchain_overview.md` |
| Architecture change | `corpus/architecture/architecture_overview.md` |
| Dependency change | `corpus/architecture/dependency_map.md` |

### Claude Assets to Update

| Evidence Type | Claude Assets |
|---------------|---------------|
| New setup requirements | Update `ant-setup` skill |
| New skill needed | Create new skill in `.claude/skills/` |
| New permissions needed | Update `.claude/settings.local.json` |
| Context changes | Update `CLAUDE.md` |

### Plan Template

```markdown
## Update Plan for: [Evidence Name]

### Summary
[One sentence: what is this and why document it]

### Evidence Type: [A/B/C/D/E/F]

### Files to Create
- [ ] `corpus/services/new-service.md` - Service documentation

### Files to Update
- [ ] `corpus/architecture/architecture_overview.md` - Add to service list
- [ ] `corpus/ops/setup_guide.md` - Add clone step

### Claude Assets
- [ ] `.claude/skills/ant-setup/SKILL.md` - Add new phase

### Verification
[How to verify the documentation is accurate]
```

**Present the plan to the user for approval before proceeding.**

## Phase 4: Implement Updates

### Writing Style Guidelines

**Target Audience**: Jr developers, PMs, and AI agents. NOT expert developers.

**DO**:
- Use simple, direct language
- Include examples for every concept
- Provide copy-paste commands with verification
- Explain "why" not just "how"
- Use tables for comparisons
- Keep sections short and scannable

**DON'T**:
- Assume deep technical knowledge
- Use jargon without explanation
- Write walls of text
- Skip verification steps
- Over-engineer documentation

### Documentation Quality Checklist

For every new section, ensure:
- [ ] **Purpose stated first**: What does this do? Why do I care?
- [ ] **Prerequisites clear**: What do I need before starting?
- [ ] **Steps numbered**: Clear 1, 2, 3 sequence
- [ ] **Commands copy-paste ready**: Full commands, no placeholders unless explained
- [ ] **Verification included**: How do I know it worked?
- [ ] **Troubleshooting added**: Common errors and fixes

### Service Documentation Template

When documenting a new service, follow `corpus/services/TEMPLATE.md`:

```markdown
# [Service Name]

## Role & Responsibility
[One paragraph: what problem does this solve?]

## Quick Facts
| Aspect | Value |
|--------|-------|
| Language | Go / TypeScript / Python |
| Tier | Core / Studio / External |
| Port | XXXX |
| Repo | `anthropos-work/service-name` |

## Architecture
[Simplified diagram or description of key components]

## Interface Discovery
- API Docs: [location]
- Proto files: [location]
- GraphQL schema: [location]

## Local Development
```bash
# Commands to run locally
```

## Testing
```bash
# Commands to run tests
```
```

### Setup Guide Updates

When adding setup steps:

1. Add step to `corpus/ops/setup_guide.md` with verification
2. Update `ant-setup` skill if phase structure changes

## Error Handling

1. **Can't access evidence**: Ask user for access or alternative
2. **Evidence unclear**: Ask clarifying questions
3. **Conflicting info**: Flag for user decision
4. **Scope creep**: Stick to original evidence, note future work

## Progress Tracking

Use TodoWrite to track:
```markdown
- [x] Phase 1: Evidence identified
- [x] Phase 2: Evidence inspected
- [x] Phase 3: Update plan created
- [ ] Phase 3: Plan approved by user
- [ ] Phase 4: corpus/services/new-service.md created
- [ ] Phase 4: architecture_overview.md updated
- [ ] Verification: Documentation review
```

## Success Criteria

Integration complete when:
1. All planned documentation created/updated
2. New information is discoverable from corpus README
3. Setup guide updated if needed
4. Claude assets updated if needed
5. Changes committed with clear message

## Invocation Examples

```bash
# Interactive mode - will ask for evidence type
/ant-integrate

# Specific evidence type
/ant-integrate A    # New repo
/ant-integrate D    # Setup feedback

# With target specified
/ant-integrate new-repo studio-analytics
/ant-integrate setup-feedback anthropos-dev/setup_progress.md
```

## Integration with Project Rosetta

This skill embodies the **Recursive Inspection** objective:
- It inspects new platform elements
- It updates the corpus with new knowledge
- It maintains documentation quality standards
- It ensures the corpus evolves with the platform

**The corpus is a living document. This skill keeps it alive.**

---
name: ant-document
description: Document new platform evidence across the Rosetta corpus with corpus-wide sweep. use when you are asked to document something that is not already documented.
argument-hint: [evidence description]
---

# Anthropos Corpus Documentation

Analyze new evidence about the Anthropos platform and update documentation **across the entire corpus** where relevant.

## Your Mission

1. **Understand the evidence**: What new information exists? (repo, feature, tool, feedback)
2. **Inspect and analyze**: What does it do? How does it fit?
3. **Sweep the corpus**: Check EVERY relevant corpus section and Claude skill
4. **Implement updates**: Write clear, accessible documentation
5. **Report issues**: Create ops reports for significant discoveries

## Evidence Types (Suggestions)

When invoking, describe what you want to document:

| Evidence | Example Invocation |
|----------|-------------------|
| New repository | `/ant-document the new studio-analytics repo` |
| New feature | `/ant-document the new skill progression API in skiller` |
| New tool | `/ant-document we added pnpm as required tooling` |
| Setup feedback | `/ant-document issues found in setup_progress.md` |
| Missing aspect | `/ant-document the Redis caching layer isn't documented` |

If invoked without arguments, ask the user what they want to document.

## DOCUMENT Principles

Apply to EVERY documentation update:

| Principle | Action |
|-----------|--------|
| Inspect First | Read and understand evidence before writing |
| Sweep Corpus | Check ALL relevant sections, not just obvious ones |
| Dual-Level | Write for both PMs (high-level) and engineers (deep dive) |
| Verify Links | Ensure new content is discoverable from parent docs |
| Include Skills | Update Claude skills when ops/workflow changes |

## Corpus Sweep Checklist

**Use TodoWrite to track which sections you've checked.** For each piece of evidence, review:

### Architecture (check all that may apply)
- `corpus/architecture/architecture_overview.md` - System design changes
- `corpus/architecture/service_taxonomy.md` - New or reclassified services
- `corpus/architecture/frontend_architecture.md` - Frontend changes
- `corpus/architecture/external_services.md` - Third-party integrations
- `corpus/architecture/dependency_map.md` - Dependency changes

### Services (check affected services)
- `corpus/services/{service}.md` - Service-specific documentation
- Follow `corpus/services/TEMPLATE.md` for new service docs

### Operations (check if workflow affected)
- `corpus/ops/setup_guide.md` - New setup requirements
- `corpus/ops/run_guide.md` - New startup steps
- `corpus/ops/update_guide.md` - New update procedures

### Tools (check if tooling affected)
- `corpus/tools/toolchain_overview.md` - New development tools

### Claude Skills (check if automation affected)
- `.claude/skills/ant-setup/` - Setup automation changes
- `.claude/skills/ant-run/` - Run automation changes
- `.claude/skills/ant-update/` - Update automation changes
- `.claude/skills/ant-document/` - This skill (if process improves)
- `CLAUDE.md` - Agent context changes

## Writing Guidelines

**Target Audience**: Jr developers, PMs, and AI agents.

**Dual-Level Documentation**:
- **High-level** (first): What is it? Why does it matter? (1-2 paragraphs)
- **Deep dive** (second): How does it work? Technical details, commands, architecture

**Style**:
- Simple, direct language
- Examples for every concept
- Copy-paste commands with verification
- Short, scannable sections
- Tables for comparisons

**Quality Checklist** (verify before completing):
- [ ] Purpose stated first (what and why)
- [ ] Prerequisites clear
- [ ] Steps numbered with verification
- [ ] Commands copy-paste ready
- [ ] Linked from parent documentation

## Confirmation Policy

**Proceed WITHOUT confirmation**:
- Reading and analyzing evidence
- Creating documentation drafts in corpus/

**ASK for confirmation before**:
- Deleting or significantly restructuring existing docs
- Changes that affect multiple interconnected files
- Updating Claude skills

## Error Handling

1. Do NOT skip unclear aspects - ask for clarification
2. Document conflicts between evidence and existing docs
3. Create ops report for significant discoveries:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-document
**Evidence**: [What was being documented]

## Discovery
[What was found that needs attention]

## Impact
[Which docs/skills are affected]

## Suggested Action
[What should be done]
```

Save to: `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_doc_<topic>.md`

## Progress Tracking

Use TodoWrite with corpus sections as checklist:

```
- Inspect evidence
- Check architecture_overview.md
- Check service_taxonomy.md
- Check dependency_map.md
- Check affected service docs
- Check setup_guide.md (if ops change)
- Check run_guide.md (if ops change)
- Check Claude skills (if automation change)
- Verify all links resolve
- Verify discoverability from corpus README
```

## Critical Rules

- **Sweep the corpus** - don't just update one file
- **Link new content** - make it discoverable
- **Follow templates** - use TEMPLATE.md for services
- **Update skills** - when ops procedures change
- Work in `anthropos-dev/` for evidence inspection only

## Success Criteria

Documentation complete when:
1. All relevant corpus sections reviewed and updated
2. New content discoverable from parent docs
3. Claude skills updated if automation affected
4. All links resolve
5. Dual-level structure maintained

## Additional Resources

- For technical patterns and examples, see [reference.md](reference.md)
- For service template, see `corpus/services/TEMPLATE.md`

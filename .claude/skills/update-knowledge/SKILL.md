---
name: update-knowledge
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
| New repository | `/update-knowledge the new studio-analytics repo` |
| New feature | `/update-knowledge the new skill progression API in app` |
| New tool | `/update-knowledge we added pnpm as required tooling` |
| Setup feedback | `/update-knowledge issues found in setup_progress.md` |
| Missing aspect | `/update-knowledge the Redis caching layer isn't documented` |

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
- `.claude/skills/dev-up/` - Dev build + start + set-dress automation changes (← the former setup-platform + start-platform)
- `.claude/skills/dev-down/` - Dev teardown automation changes
- `.claude/skills/stack-update/` - Stack code/deps/schema sync automation changes (← update-platform)
- `.claude/skills/stack-list/`, `stack-seed/`, `stack-snapshot/` - Generic stack-ops changes (← demo-status/seed/snapshot)
- `.claude/skills/update-knowledge/` - This skill (if process improves)
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
**Skill**: /update-knowledge
**Evidence**: [What was being documented]

## Discovery
[What was found that needs attention]

## Impact
[Which docs/skills are affected]

## Suggested Action
[What should be done]
```

Save to: `stack-dev/ops-reports/op_YYYYMMDD_HHMMSS_doc_<topic>.md`

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
- Work in `stack-dev/` for evidence inspection only
- **Know the corpus/tooling boundary** - rosetta is a read-only doc corpus + dev-env skills; all code/scripts that operate the platform on a spawned stack live in **rosetta-extensions**, the executable stack tooling. When new evidence is EXECUTABLE stack tooling, it does NOT get documented as scripts into the rosetta corpus — it belongs in rosetta-extensions, authored and tested in the AUTHORING copy at `.agentspace/rosetta-extensions/`, committed, then **tagged**. Each stack consumes it as a pinned per-stack copy (`stack-<role>/rosetta-extensions @ <tag>`). Make the sweep aware of both: the `.agentspace/rosetta-extensions/` authoring copy and the per-stack `stack-*/rosetta-extensions @ <tag>` consumption copies.

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

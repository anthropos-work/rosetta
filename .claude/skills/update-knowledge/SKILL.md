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

**Use TodoWrite to track which sections you've checked.**

> **⚠️ Enumerate from the indexes, never from this page.** The corpus is **90 markdown files** and the skill
> set is **16**. This checklist used to hardcode **10** corpus docs and **8** skills — so the demo family
> (23 files), the fenced migration map, and half the skills were invisible to every sweep that trusted it.
> That is a live root cause, not a hypothetical: 12 of 20 skill files went unswept across three platform
> moves. **Start each sweep by listing the directory or its README**, then narrow.

### Step 0 — enumerate, then narrow

```bash
ls corpus/architecture/*.md corpus/ops/*.md corpus/ops/demo/*.md corpus/services/*.md corpus/tools/*.md
ls -d .claude/skills/*/          # every skill, not a remembered subset
```

Each section has a maintained index — read it before guessing a filename:
`corpus/README.md` · `corpus/architecture/README.md` · `corpus/services/README.md` (the enumerated index of
all service docs) · `corpus/ops/README.md` · `corpus/ops/demo/README.md` · `corpus/tools/README.md`.

### Architecture (`corpus/architecture/`, 11 files)
Start with **`platform-migration-status.md`** for anything touching which services exist — it is the
**fenced** one-row-per-service map (machine-checked against the platform's own `repos.yml` in both
directions by `stack-core/platform_alignment_guard.py`), and it outranks prose anywhere else, including
`CLAUDE.md`. Then as applicable: `architecture_overview.md` · `service_taxonomy.md` · `dependency_map.md` ·
`frontend_architecture.md` · `external_services.md` · `shared_libraries.md` · `security_compliance.md` ·
`ai_architecture.md` · `alignment_testing.md`.

### Services (`corpus/services/`, 29 files)
`corpus/services/README.md` is the enumerated index — start there rather than guessing. Follow
`corpus/services/TEMPLATE.md` for a new service doc. Remember the archived/merged redirects
(`skiller`, `skillpath`, `chronos`, `intelligence`) are docs too.

### Operations (`corpus/ops/`, 23 files + `corpus/ops/demo/`, 23 more)
Core: `setup_guide.md` · `run_guide.md` · `update_guide.md` · `platform-alignment.md` · `safety.md` ·
`verification.md` · `secrets-spec.md` · `seeding-spec.md` · `snapshot-spec.md` · `idempotency.md` ·
`db-access.md` · `directus-local.md` · `rosetta_demo.md` · `quick_ops.md` · the `staging-*` family.
**The whole `corpus/ops/demo/` family counts** — `demo-up-defaults.md` (fenced against the parsers),
`build-budget.md`, `latency-budget.md`, `coverage-protocol.md`, `playthroughs.md`, the `recipe-*` and
`*-spec.md` pages. List the dir; do not work from memory.

### Tools (`corpus/tools/`, 3 files)
`README.md` (the index) · `toolchain_overview.md` · `anthropos-labs.md`.

### Claude Skills — **all 16, check the list, don't recall it**
`ls -d .claude/skills/*/`. At time of writing: `align-dna` · `align-run` · `db-query` · `demo-down` ·
`demo-up` · `dev-down` · `dev-for-dummies` · `dev-up` · `setup-github` · `stack-list` · `stack-secrets` ·
`stack-seed` · `stack-snapshot` · `stack-update` · `test-platform` · `update-knowledge`. Several carry a
`reference.md` beside `SKILL.md` — **sweep both**; a SKILL.md and its own reference.md have contradicted
each other before, and the SKILL.md is what an agent actually follows.

Plus `CLAUDE.md` — agent context, and its skill table (every row's guide-doc pointer must resolve).

### A skill is executable — grade it harder than a doc
A stale doc misleads; a stale skill **runs**. When platform evidence changes, check each skill for:
service names · profile tokens · compose commands · `repos.yml` expectations · schema names · ports ·
env vars · paths into a platform clone. **The most dangerous shape is a command that still exits 0 while
selecting nothing** — `postgresql`, `redis` and `sentinel` declare no `profiles:` key, so a retired
profile token starts the floor and the stack looks alive with the application absent. Grade a documented
command on *"does it still select anything"*, never on *"does it still parse."*

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

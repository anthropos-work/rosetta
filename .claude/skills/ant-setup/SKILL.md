---
name: ant-setup
description: Build Anthropos dev environment following setup guide with auto-improvement
argument-hint: [step-name or 'full']
---

# Anthropos Development Environment Setup

Execute the Anthropos platform setup by following `corpus/ops/platform-setup/setup_guide.md` while applying STEP RUN guidelines and auto-improving documentation.

## Your Mission

1. **Follow the guide**: Use `corpus/ops/platform-setup/setup_guide.md` as your source of truth
2. **Apply STEP RUN guidelines**: Verify before/after, request confirmation, document improvements
3. **Track progress locally**: Copy checklist to `anthropos-dev/setup_progress.md` and update as you go
4. **Auto-improve docs**: Update setup_guide.md when you discover issues or better approaches
5. **Update structure only when needed**: Only modify OS-specific checklists (`corpus/ops/platform-setup/setup_checklist_macos.md`, `setup_checklist_linux.md`) and this skill when setup structure changes

## STEP RUN Guidelines (Apply to Every Step)

### 1. Verify Before Install
Check if tool exists before attempting installation:
```bash
tool --version 2>/dev/null || echo "Not installed"
```

### 2. Request Confirmation
**ALWAYS ask user before**:
- Installing system tools
- Cloning repositories
- Starting services
- Creating/modifying .env files

Use AskUserQuestion tool.

### 3. Verify After Install
Run same verification command to confirm success.

### 4. Document Improvements
**Immediately update setup_guide.md** when you discover:
- Missing verification commands
- Better installation approaches
- New steps required
- Errors needing troubleshooting entries

### 5. Track in Local Checklist
Update `anthropos-dev/setup_progress.md` as you complete steps (NOT the original in corpus/).

## Initial Setup

1. Copy checklist: `cp corpus/ops/platform-setup/setup_checklist_macos.md anthropos-dev/setup_progress.md` (or `setup_checklist_linux.md` for Linux)
2. Read `corpus/ops/platform-setup/setup_guide.md` to understand the process
3. Navigate to `anthropos-dev/` workspace
4. Follow guide section by section

## When to Update Original Checklist

**Only update OS-specific checklists when setup structure changes:**
- NEW step added to guide
- Step removed from guide
- Steps reordered

**Do NOT update original checklist for:**
- Clarifications to existing steps
- Added verification commands
- Typo fixes
- Troubleshooting additions

**When structure changes, also update this skill** (`.claude/skills/ant-setup/SKILL.md`) if needed.

## Error Handling

1. **Do NOT skip errors** - resolve them first
2. Document error message verbatim
3. Research solution
4. Test fix
5. Add to setup_guide.md troubleshooting section
6. Continue

## Progress Tracking

Use TodoWrite for high-level tracking:
```markdown
- [x] Copied checklist to anthropos-dev/
- [x] Prerequisites: Git, Docker verified
- [ ] Prerequisites: Go, Node, pnpm
- [ ] GitHub SSH access configured
- [ ] Repositories cloned
```

Use local checklist (`anthropos-dev/setup_progress.md`) for detailed step tracking and issue logging.

## Critical Rules

- Work in `anthropos-dev/` scratchpad only
- Never commit .env files
- Use `-p ant-rosetta` for Docker isolation
- Ask before every destructive operation
- Follow the guide - don't improvise unless needed
- Update docs immediately when improvements found

## Success Criteria

Setup complete when:
1. All tools installed and verified
2. All repositories cloned
3. Environment files configured
4. Docker services running and healthy
5. Frontend accessible at localhost:3000
6. Studio services running
7. Documentation improvements committed (if any)
8. Local checklist fully complete

**Follow `corpus/ops/platform-setup/setup_guide.md` as your primary reference. Apply these guidelines to execute it reliably and improve it recursively.**

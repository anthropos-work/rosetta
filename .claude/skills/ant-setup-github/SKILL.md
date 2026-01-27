---
name: ant-setup-github
description: Configure GitHub SSH access for Anthropos organization with single or dual account support
argument-hint: [single|dual]
---

# GitHub SSH Setup for Anthropos

Configure GitHub SSH access to contribute to `https://github.com/anthropos-work/` repositories.

## Your Mission

1. **Read the guide**: `corpus/ops/setup_github_guide.md` is your source of truth
2. **Apply STEP RUN principles**: Verify before/after, request confirmation
3. **Track progress**: Use TodoWrite for each phase (no separate checklist file)
4. **Report issues**: Create ops reports for errors and fixes discovered

## STEP RUN Principles

Apply to EVERY step in the guide:

| Principle | Action |
|-----------|--------|
| Verify Before | Check if SSH keys/config exist before generating |
| Request Confirmation | Ask user before generating keys or modifying SSH config |
| Verify After | Test SSH connection after each configuration change |
| Report Issues | Create ops report when errors or improvements found |

## First: Determine the Scenario

Ask the user which scenario applies:

| Scenario | Description | Guide Section |
|----------|-------------|---------------|
| **Single Account** | One GitHub account for all work (existing or new) | [Single Account Setup] |
| **Dual Accounts** | Personal account + separate work account | [Two Accounts Setup] |

## Confirmation Policy

**ALWAYS ask for confirmation before**:
- Generating SSH keys
- Modifying `~/.ssh/config`
- Running `ssh-add -D` (clears all keys)
- Adding keys to SSH agent

## Verification Commands

Use these to verify state before/after each step:

```bash
# Check existing SSH keys
ls -la ~/.ssh/

# Check SSH config exists
cat ~/.ssh/config 2>/dev/null || echo "No SSH config found"

# List keys in agent
ssh-add -l

# Test GitHub connection
ssh -T git@github.com

# Test personal alias (dual setup only)
ssh -T git@github.com-personal
```

## Error Handling

1. Do NOT skip errors - resolve first
2. Document error message verbatim
3. Research and test solution
4. Create ops report (see below)
5. Continue

## Ops Reports

When you discover errors, missing steps, or better approaches:

1. Create a report file: `anthropos-dev/ops-reports/op_YYYYMMDD_HHMMSS_github_<topic>.md`
2. Use this template:

```markdown
# Ops Report: [Brief Title]

**Date**: YYYY-MM-DD HH:MM
**Skill**: /ant-setup-github
**OS**: [macOS/Linux/version]
**Phase**: [Which setup phase]
**Scenario**: [Single/Dual account]

## Issue Encountered
[Exact error message or problem description]

## Context
[What step was being executed, what commands were run]

## Resolution
[How it was fixed, or "Unresolved" if still broken]

## Suggested Documentation Update
[What should be added/changed in setup_github_guide.md]
```

## Progress Tracking

Use TodoWrite with phases appropriate to the scenario:

### Single Account Phases
- Determine scenario (single vs dual)
- Check/Create GitHub account
- Generate SSH key (~/.ssh/id_github_work)
- Configure SSH (~/.ssh/config)
- Add key to SSH agent
- Add public key to GitHub
- Test connection
- Request organization access
- Clone test: rosetta repo to ./anthropos-dev/

### Dual Account Phases
- Determine scenario (single vs dual)
- Verify work GitHub account
- Generate work SSH key (~/.ssh/id_github_work)
- Generate personal SSH key (~/.ssh/id_github_personal)
- Configure SSH (~/.ssh/config) - work as default
- Clear and re-add keys (work first)
- Add keys to both GitHub accounts
- Test both connections
- Request organization access
- Clone test: rosetta repo to ./anthropos-dev/

## Critical Rules

- **Work account MUST be default**: The `github.com` host must use the work key for Docker compatibility
- **Order matters**: When adding keys to agent, add work key FIRST (`ssh-add -D` then `ssh-add ~/.ssh/id_github_work`)
- **Verify persistence**: Ensure keys survive terminal/computer restart
- **Never expose private keys**: Only share `.pub` files
- **Ask before modifying**: Always confirm before changing SSH config or clearing keys

## Docker Compatibility Note

Docker and many build tools cannot use SSH host aliases. They always connect to `github.com` directly. This is why:
1. The work key MUST be configured as default for `Host github.com`
2. Personal repos use the alias `github.com-personal`
3. When troubleshooting Docker git issues, verify with:
   ```bash
   ssh-add -D
   ssh-add ~/.ssh/id_github_work
   ssh -T git@github.com  # Must show work account
   ```

## Success Criteria

Setup complete when:

### Single Account
1. SSH key generated and secured
2. SSH config properly configured
3. Key added to SSH agent (persistent)
4. Key added to GitHub account
5. `ssh -T git@github.com` shows correct username
6. Organization access requested
7. Successfully cloned `rosetta` to `./anthropos-dev/rosetta`

### Dual Account
1. Both SSH keys generated and secured
2. SSH config has work as default, personal as alias
3. Both keys in agent (work added first)
4. Keys added to respective GitHub accounts
5. `ssh -T git@github.com` shows WORK username
6. `ssh -T git@github.com-personal` shows PERSONAL username
7. Organization access requested
8. Keys persist after restart
9. Successfully cloned `rosetta` to `./anthropos-dev/rosetta`

## Additional Resources

- For complete setup instructions, read `corpus/ops/setup_github_guide.md`

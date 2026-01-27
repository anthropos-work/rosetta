# GitHub Account Setup Guide for Anthropos

This guide helps Anthropos employees configure GitHub access to contribute to the organization repositories at `https://github.com/anthropos-work/`.

## Prerequisites

- macOS or Linux operating system
- Terminal access
- GitHub account (existing or will create one)

## Choose Your Scenario

Before starting, determine which scenario applies to you:

| Scenario | Description | Go to |
|----------|-------------|-------|
| **Single Account** | You have one GitHub account (or will create one) that you'll use for all work | [Single Account Setup](#single-account-setup) |
| **Two Accounts** | You have a personal GitHub account AND need a separate work account | [Two Accounts Setup](#two-accounts-setup) |

---

## Single Account Setup

Use this if you have only one GitHub account or are creating your first one.

### Phase 1: GitHub Account

#### 1.1 Check for Existing Account

If you already have a GitHub account you want to use for work:
- Verify you can log in at https://github.com
- Skip to [Phase 2: SSH Key Generation](#phase-2-ssh-key-generation-single)

#### 1.2 Create New Account (if needed)

1. Go to https://github.com/signup
2. Use your work email address
3. Choose a professional username
4. Complete email verification

### Phase 2: SSH Key Generation (Single) {#phase-2-ssh-key-generation-single}

#### 2.1 Check for Existing SSH Keys

```bash
ls -la ~/.ssh/
```

Look for files like `id_ed25519`, `id_rsa`, or `id_github_work`. If you have existing keys you want to use, skip to [Phase 3](#phase-3-ssh-config-single).

#### 2.2 Generate New SSH Key

```bash
ssh-keygen -t ed25519 -C "your-work-email@company.com" -f ~/.ssh/id_github_work
```

When prompted:
- **Passphrase**: Enter a secure passphrase (recommended) or leave empty

#### 2.3 Verify Key Creation

```bash
ls -la ~/.ssh/id_github_work*
```

Expected output:
```
-rw-------  1 user  staff  XXX  date time  /Users/user/.ssh/id_github_work
-rw-r--r--  1 user  staff  XXX  date time  /Users/user/.ssh/id_github_work.pub
```

### Phase 3: SSH Config (Single) {#phase-3-ssh-config-single}

#### 3.1 Create/Update SSH Config

Create or edit `~/.ssh/config`:

```bash
touch ~/.ssh/config
chmod 600 ~/.ssh/config
```

Add the following configuration:

```
# GitHub - Work Account (Default)
Host github.com
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_github_work
    AddKeysToAgent yes
    UseKeychain yes
    IdentitiesOnly yes
```

**Note for Linux users**: Remove the `UseKeychain yes` line (macOS-specific).

### Phase 4: Add Key to SSH Agent (Single) {#phase-4-ssh-agent-single}

#### 4.1 Start SSH Agent

```bash
eval "$(ssh-agent -s)"
```

#### 4.2 Add Key to Agent

**macOS:**
```bash
ssh-add --apple-use-keychain ~/.ssh/id_github_work
```

**Linux:**
```bash
ssh-add ~/.ssh/id_github_work
```

#### 4.3 Verify Key is Loaded

```bash
ssh-add -l
```

Expected output should show your key fingerprint.

### Phase 5: Add Key to GitHub (Single) {#phase-5-github-single}

#### 5.1 Copy Public Key

```bash
cat ~/.ssh/id_github_work.pub | pbcopy  # macOS
# OR
cat ~/.ssh/id_github_work.pub | xclip -selection clipboard  # Linux
```

#### 5.2 Add to GitHub

1. Go to https://github.com/settings/keys
2. Click "New SSH key"
3. Title: "Work Machine - [Your Device Name]"
4. Key type: "Authentication Key"
5. Paste the public key
6. Click "Add SSH key"

### Phase 6: Test Connection (Single) {#phase-6-test-single}

```bash
ssh -T git@github.com
```

Expected output:
```
Hi username! You've successfully authenticated, but GitHub does not provide shell access.
```

### Phase 7: Request Organization Access

1. Contact your team lead or organization admin
2. Provide your GitHub username
3. Wait for invitation to `anthropos-work` organization
4. Accept the invitation at https://github.com/orgs/anthropos-work/invitation

---

## Two Accounts Setup

Use this if you have a personal GitHub account and need a separate work account, keeping them isolated.

### Phase 1: GitHub Accounts

#### 1.1 Verify Personal Account

- Log in to your personal GitHub account
- Note your personal email address used

#### 1.2 Create Work Account (if needed)

1. Log out of personal account (or use incognito mode)
2. Go to https://github.com/signup
3. Use your **work email address**
4. Choose a professional username (can include company name)
5. Complete email verification

### Phase 2: SSH Key Generation (Dual)

#### 2.1 Check Existing Keys

```bash
ls -la ~/.ssh/
```

#### 2.2 Generate Work SSH Key

```bash
ssh-keygen -t ed25519 -C "your-work-email@company.com" -f ~/.ssh/id_github_work
```

#### 2.3 Generate Personal SSH Key (if not exists)

```bash
ssh-keygen -t ed25519 -C "your-personal-email@example.com" -f ~/.ssh/id_github_personal
```

#### 2.4 Verify Both Keys Exist

```bash
ls -la ~/.ssh/id_github_*
```

Expected output:
```
-rw-------  1 user  staff  XXX  date time  /Users/user/.ssh/id_github_personal
-rw-r--r--  1 user  staff  XXX  date time  /Users/user/.ssh/id_github_personal.pub
-rw-------  1 user  staff  XXX  date time  /Users/user/.ssh/id_github_work
-rw-r--r--  1 user  staff  XXX  date time  /Users/user/.ssh/id_github_work.pub
```

### Phase 3: SSH Config (Dual)

#### 3.1 Create/Update SSH Config

```bash
touch ~/.ssh/config
chmod 600 ~/.ssh/config
```

**Important**: The order matters! The **work** configuration must come FIRST for `github.com` to make it the default.

Edit `~/.ssh/config` with the following:

```
# ============================================
# DEFAULT: github.com uses WORK key
# This is critical for Docker and tools that
# don't support host aliases
# ============================================
Host github.com
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_github_work
    AddKeysToAgent yes
    UseKeychain yes
    IdentitiesOnly yes

# ============================================
# PERSONAL: Use github.com-personal in git URLs
# Example: git clone git@github.com-personal:user/repo.git
# ============================================
Host github.com-personal
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_github_personal
    AddKeysToAgent yes
    UseKeychain yes
    IdentitiesOnly yes

# ============================================
# EXPLICIT WORK ALIAS (optional)
# Use github.com-work when you want to be explicit
# Example: git clone git@github.com-work:anthropos-work/repo.git
# ============================================
Host github.com-work
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_github_work
    AddKeysToAgent yes
    UseKeychain yes
    IdentitiesOnly yes
```

**Note for Linux users**: Remove the `UseKeychain yes` lines (macOS-specific).

### Phase 4: Add Keys to SSH Agent (Dual)

#### 4.1 Clear Existing Keys (Clean Slate)

```bash
ssh-add -D
```

This ensures no conflicting keys are loaded.

#### 4.2 Start SSH Agent

```bash
eval "$(ssh-agent -s)"
```

#### 4.3 Add Work Key FIRST (Critical for Default)

**macOS:**
```bash
ssh-add --apple-use-keychain ~/.ssh/id_github_work
```

**Linux:**
```bash
ssh-add ~/.ssh/id_github_work
```

#### 4.4 Add Personal Key

**macOS:**
```bash
ssh-add --apple-use-keychain ~/.ssh/id_github_personal
```

**Linux:**
```bash
ssh-add ~/.ssh/id_github_personal
```

#### 4.5 Verify Keys are Loaded

```bash
ssh-add -l
```

Expected output should show both key fingerprints. The work key should be listed first.

### Phase 5: Add Keys to GitHub (Dual)

#### 5.1 Add Work Key to Work Account

1. Copy the work public key:
   ```bash
   cat ~/.ssh/id_github_work.pub | pbcopy  # macOS
   ```
2. Log in to your **work** GitHub account
3. Go to https://github.com/settings/keys
4. Click "New SSH key"
5. Title: "Work Machine - [Your Device Name]"
6. Paste the public key
7. Click "Add SSH key"

#### 5.2 Add Personal Key to Personal Account

1. Copy the personal public key:
   ```bash
   cat ~/.ssh/id_github_personal.pub | pbcopy  # macOS
   ```
2. Log in to your **personal** GitHub account
3. Go to https://github.com/settings/keys
4. Click "New SSH key"
5. Title: "Work Machine - Personal"
6. Paste the public key
7. Click "Add SSH key"

### Phase 6: Test Both Connections (Dual)

#### 6.1 Test Work Account (Default)

```bash
ssh -T git@github.com
```

Expected output:
```
Hi work-username! You've successfully authenticated, but GitHub does not provide shell access.
```

#### 6.2 Test Personal Account (Alias)

```bash
ssh -T git@github.com-personal
```

Expected output:
```
Hi personal-username! You've successfully authenticated, but GitHub does not provide shell access.
```

#### 6.3 Test Explicit Work Alias

```bash
ssh -T git@github.com-work
```

Expected output:
```
Hi work-username! You've successfully authenticated, but GitHub does not provide shell access.
```

### Phase 7: Request Organization Access

1. Contact your team lead or organization admin
2. Provide your **work** GitHub username
3. Wait for invitation to `anthropos-work` organization
4. Accept the invitation at https://github.com/orgs/anthropos-work/invitation

---

## Ensuring Persistence Across Restarts

### macOS: Keychain Integration

The `AddKeysToAgent yes` and `UseKeychain yes` options in your SSH config ensure keys are stored in the macOS Keychain and automatically loaded.

To verify persistence after a restart:
```bash
ssh-add -l
```

If keys are not loaded, the SSH agent will automatically use the Keychain when you attempt a connection.

### Linux: SSH Agent Startup Script

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
# Start SSH agent if not running
if [ -z "$SSH_AUTH_SOCK" ]; then
    eval "$(ssh-agent -s)"
    ssh-add ~/.ssh/id_github_work
    # For dual setup, also add:
    # ssh-add ~/.ssh/id_github_personal
fi
```

### Docker Compatibility

Docker and other tools that use git internally often cannot handle SSH host aliases. This is why **the work account must be the default** for `github.com`.

When Docker builds images that clone from `github.com`, it will use the work key automatically.

To verify the correct key is used:
```bash
# Clear and re-add work key first
ssh-add -D
ssh-add ~/.ssh/id_github_work

# Verify
ssh-add -l
ssh -T git@github.com  # Should show work account
```

---

## Git Configuration

### Per-Repository Config (Recommended for Dual Setup)

For work repositories:
```bash
cd ~/path/to/work-repo
git config user.name "Your Name"
git config user.email "your-work-email@company.com"
```

For personal repositories (using alias):
```bash
cd ~/path/to/personal-repo
git remote set-url origin git@github.com-personal:username/repo.git
git config user.name "Your Name"
git config user.email "your-personal-email@example.com"
```

### Global Config (Single Account Setup)

```bash
git config --global user.name "Your Name"
git config --global user.email "your-work-email@company.com"
```

---

## Troubleshooting

### Permission Denied (publickey)

1. Verify key is added to agent:
   ```bash
   ssh-add -l
   ```

2. Verify key is added to GitHub:
   - Go to https://github.com/settings/keys
   - Confirm your public key is listed

3. Test with verbose output:
   ```bash
   ssh -vT git@github.com
   ```

### Wrong Account Being Used

1. Check which key is being used:
   ```bash
   ssh -vT git@github.com 2>&1 | grep "Offering public key"
   ```

2. Reset and re-add keys in correct order:
   ```bash
   ssh-add -D
   ssh-add ~/.ssh/id_github_work
   # Then add personal if needed
   ```

3. Verify SSH config has work account first for `github.com`

### Keys Not Persisting After Restart

**macOS:**
- Ensure `UseKeychain yes` is in your SSH config
- Run: `ssh-add --apple-use-keychain ~/.ssh/id_github_work`

**Linux:**
- Add the SSH agent startup script to your shell profile (see [Linux: SSH Agent Startup Script](#linux-ssh-agent-startup-script))

### Docker Can't Clone Private Repos

1. Ensure work key is the default:
   ```bash
   ssh-add -D
   ssh-add ~/.ssh/id_github_work
   ssh -T git@github.com  # Should show work account
   ```

2. For Docker BuildKit, ensure SSH agent forwarding:
   ```bash
   DOCKER_BUILDKIT=1 docker build --ssh default .
   ```

---

## Quick Reference

### Commands Cheat Sheet

| Action | Command |
|--------|---------|
| List loaded keys | `ssh-add -l` |
| Clear all keys | `ssh-add -D` |
| Add work key (macOS) | `ssh-add --apple-use-keychain ~/.ssh/id_github_work` |
| Add work key (Linux) | `ssh-add ~/.ssh/id_github_work` |
| Test work connection | `ssh -T git@github.com` |
| Test personal connection | `ssh -T git@github.com-personal` |
| Verbose test | `ssh -vT git@github.com` |

### URL Formats for Dual Setup

| Account | Clone URL Format |
|---------|------------------|
| Work (default) | `git@github.com:anthropos-work/repo.git` |
| Personal | `git@github.com-personal:username/repo.git` |
| Work (explicit) | `git@github.com-work:anthropos-work/repo.git` |

---

## Final Verification: Clone Test

The ultimate test of your setup is successfully cloning an organization repository.

### Create Workspace Directory

```bash
mkdir -p ./anthropos-dev
cd ./anthropos-dev
```

### Clone Rosetta Repository

```bash
git clone git@github.com:anthropos-work/rosetta.git
```

Expected output:
```
Cloning into 'rosetta'...
remote: Enumerating objects: ...
remote: Counting objects: ...
remote: Compressing objects: ...
Receiving objects: 100% ...
Resolving deltas: 100% ...
```

### Verify Clone

```bash
cd rosetta
git remote -v
```

Expected output:
```
origin  git@github.com:anthropos-work/rosetta.git (fetch)
origin  git@github.com:anthropos-work/rosetta.git (push)
```

### Configure Git Identity for This Repo

```bash
git config user.name "Your Name"
git config user.email "your-work-email@company.com"
```

Verify:
```bash
git config user.name
git config user.email
```

**Note**: If the clone fails with "Permission denied (publickey)", ensure:
1. You have been added to the `anthropos-work` organization
2. Your work SSH key is loaded: `ssh-add -l`
3. The correct key is being used: `ssh -T git@github.com`

---

## Success Criteria

Setup is complete when:

1. SSH key(s) generated and secured
2. SSH config properly configured
3. Key(s) added to SSH agent
4. Key(s) added to GitHub account(s)
5. Connection test successful: `ssh -T git@github.com` shows work account
6. (Dual setup) Personal alias works: `ssh -T git@github.com-personal`
7. Organization access requested/granted
8. Keys persist after terminal/computer restart
9. Successfully cloned `rosetta` repo to `./anthropos-dev/rosetta`

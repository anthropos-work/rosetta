# Zero to Hero: Anthropos Development Setup Guide

This guide is designed to take you from a **fresh computer** (or a clean folder) to a fully running Anthropos development environment.

> **Companion Checklist**: This guide is paired with OS-specific setup checklists ([macOS](./setup_checklist_macos.md) | [Linux](./setup_checklist_linux.md)). We recommend AI Agents and Engineers copy the appropriate checklist to their workspace to track progress, pause/resume setup, and log any errors encountered.

## 1. Prerequisites

Before we write any code, ensure you have the following tools installed.

### Setup Execution Guidelines (STEP RUN)

When executing each step of this setup guide, follow these guidelines to ensure reliability and documentation:

1. **Verify Before Install**: Run verification commands to check if a tool is already installed before attempting installation.
2. **Verify After Install**: After installation, run verification commands to confirm successful installation.
3. **Request Confirmation**: Always ask for user confirmation before installing or modifying system tools.
4. **Document Verification Commands**: Update this guide with verification commands for each tool if not already present.
5. **Continuous Improvement**: If you discover new steps or issues during setup, document them in this guide for future users.

### Essential Tools

### Essential Tools

#### 1. OS-Specific Setup

<details open>
<summary><strong>MacOS</strong></summary>

We recommend using [Homebrew](https://brew.sh/) for package management.

1.  **Git**: `brew install git`
    *   *Verification*: `git --version`
2.  **Docker Desktop**: [Install Docker Desktop for Mac](https://www.docker.com/products/docker-desktop/).
    *   *Verification*: `docker --version && docker compose version`
3.  **Visual Studio Code**: [Install VS Code](https://code.visualstudio.com/).
    *   *Verification*: `code --version`
4.  **Go** (v1.23+): `brew install go`
    *   *Verification*: `go version`.
5.  **Node.js** (v20+) & **pnpm**:
    *   `brew install node`
    *   `corepack enable` or `npm install -g pnpm`.
    *   *Verification*: `node --version && pnpm --version`
6.  **Build Tools**:
    *   Ensure XCode CLI tools are installed: `xcode-select --install`
    *   Ensure XCode CLI tools are installed: `xcode-select --install`
    *   *Verification*: `xcode-select -p`
7.  **Python** (v3.8+ for Studio-Room):
    *   `brew install python`
    *   *Verification*: `python3 --version`

</details>

<details>
<summary><strong>Linux (Ubuntu/Debian)</strong></summary>

1.  **Git & Build Tools**:
    ```bash
    sudo apt-get update
    sudo apt-get install -y git build-essential
    ```
2.  **Docker**:
    *   Remove conflicting packages and install Docker Engine: [Official Guides](https://docs.docker.com/engine/install/ubuntu/).
    *   **Crucial Step**: Manage Docker as a non-root user (so you don't need `sudo`):
        ```bash
        sudo usermod -aG docker $USER
        newgrp docker
        ```
    *   *Verification*: `docker ps`.
3.  **Visual Studio Code**: [Install VS Code](https://code.visualstudio.com/docs/setup/linux).
4.  **Go** (v1.23+):
    *   [Official Install](https://go.dev/doc/install) is recommended to get the latest version, as apt repos are often outdated.
    *   *Verification*: `go version`.
5.  **Node.js** (v20+) & **pnpm**:
    *   Use [nodesource](https://github.com/nodesource/distributions) or `nvm` to get v20+.
    *   `corepack enable` or `npm install -g pnpm`.
6.  **Python** (v3.8+ for Studio-Room):
    *   `sudo apt-get install python3 python3-pip python3-venv`
    *   *Verification*: `python3 --version`

</details>

---

## Automated Setup with Claude Code

If you're using **Claude Code**, you can automate this entire setup process using the `/anthropos-setup` skill:

```bash
/anthropos-setup
```

The skill will:
*   Execute each step with verification before and after
*   Request your confirmation before installing tools or making changes
*   Copy and track progress in a local checklist
*   Auto-improve this documentation when it discovers issues

See [`.claude/skills/anthropos-setup/`](../../.claude/skills/anthropos-setup/) for details.

---

## 2. GitHub SSH Access Setup

Before cloning repositories, you need SSH access to the `anthropos-work` GitHub organization.

### Check Existing SSH Keys

First, verify if you already have SSH keys configured:

```bash
ls -al ~/.ssh
```

Look for files named `id_rsa.pub`, `id_ed25519.pub`, or similar `.pub` files.

### Test GitHub SSH Connection

```bash
ssh -T git@github.com
```

*Expected output*: `Hi [username]! You've successfully authenticated, but GitHub does not provide shell access.`
*Note*: Exit code 1 is normal for this command - it indicates successful authentication.

### If SSH Keys Don't Exist

1. **Generate a new SSH key**:
   ```bash
   ssh-keygen -t ed25519 -C "your_email@example.com"
   ```
   Press Enter to accept the default file location. Optionally set a passphrase.

2. **Start the ssh-agent**:
   ```bash
   eval "$(ssh-agent -s)"
   ```

3. **Add your SSH key to the ssh-agent** (macOS):
   ```bash
   ssh-add --apple-use-keychain ~/.ssh/id_ed25519
   ```

4. **Copy your public key**:
   ```bash
   pbcopy < ~/.ssh/id_ed25519.pub
   ```

5. **Add the SSH key to your GitHub account**:
   - Go to GitHub Settings > SSH and GPG keys
   - Click "New SSH key"
   - Paste your key and save

6. **Request access to anthropos-work organization**:
   - Contact the Engineering Manager to be added to the `anthropos-work` GitHub organization.

*Verification*: `ssh -T git@github.com` should show successful authentication.

---

## 3. Workspace Setup

We will create a dedicated workspace to house all the microservices and the frontend.

1.  Open your terminal.
2.  Navigate to the `rosetta` directory in this repository.
3.  Enter the scratchpad directory:
    ```bash
    cd anthropos-dev
    ```
    *(Note: This directory is git-ignored, so you can clone anything here without messing up the main repo).*

---

## 4. Cloning Repositories

We need to fetch the code for the platform. Run the following commands inside `rosetta/anthropos-dev`:

### Platform Configuration
This repository contains the `docker-compose.yml` to orchestrate everything.
```bash
git clone git@github.com:anthropos-work/platform.git
```
*Verification*: `ls -la platform` should show `docker-compose.yml` and other config files.

### Backend Services
Clone the core Go services.
```bash
git clone git@github.com:anthropos-work/app.git backend
git clone git@github.com:anthropos-work/cms.git
git clone git@github.com:anthropos-work/jobsimulation.git
# Remote-only services (optional source, strictly needed only if editing):
# git clone git@github.com:anthropos-work/sentinel.git
# git clone git@github.com:anthropos-work/skiller.git
```
*Verification*: `ls -la backend cms jobsimulation` should show all three directories with Go files.

### Frontend
Clone the Next.js monorepo.
```bash
git clone git@github.com:anthropos-work/next-web-app.git
```
*Verification*: `ls -la next-web-app` should show Next.js project files including `package.json`.

### Studio Services
Clone the studio monorepo containing Studio-Desk and Studio-Room.
```bash
git clone git@github.com:anthropos-work/studio.git
```
*Verification*: `ls -la studio` should show `studio-desk` and `studio-room` directories.

---

## 5. Environment Configuration

### The `.env` File
All services share a **single centralized `.env` file** located in the `platform` directory.

> **IMPORTANT**: You must obtain the master `.env` values from the Engineering Manager or the 1Password Vault "Engineering/Env".

1.  **Create the environment file**:
    ```bash
    cd platform
    cp .env_example .env
    ```
2.  **Populate secrets**: Edit `platform/.env` and fill in all required secret values from 1Password or the Engineering Manager.
    
    **Critical Keys Required**:
    *   `CLERK_SECRET_KEY` & `CLERK_PUBLISHABLE_KEY` (Auth)
    *   `OPENAI_API_KEY` (AI services)
    *   `ANTHROPIC_API_KEY` (AI services)
    *   `AZURE_API_KEY` (Optional, if using Azure OpenAI)
    *   `DIRECTUS_PUBLIC_BASE_ADDR` (Content)

3.  **Studio Environment**:
    *   Copy `studio/studio-desk/.env.example` to `studio/studio-desk/.env`
    *   Populate matching keys from `platform/.env` (Clerk, OpenAI)
3.  **Verification**: `ls -la platform/.env` should show the file exists.

**Note**: The docker-compose configuration uses this single `.env` file for all services (backend, cms, jobsimulation, etc.). Individual service repositories do not need their own `.env` files when running via Docker.

---

## 6. Running the Platform (Docker)

The easiest way to start is using Docker Compose.

### Docker Compose Project Name

We use the `-p anthropos-rosetta` flag to set a custom project name. This creates an isolated Docker stack that won't conflict with other Anthropos environments you may have running.

**What this does:**
- Creates containers named `anthropos-rosetta-postgresql-1`, `anthropos-rosetta-backend-1`, etc.
- Creates isolated networks: `anthropos-rosetta_app-network`
- Creates isolated volumes: `anthropos-rosetta_postgres_data`

**Note**: If you have another Anthropos stack running (e.g., "platform"), they will be completely separate. However, you may encounter **port conflicts** if both stacks try to use the same ports. Stop the other stack first or modify port mappings in docker-compose.yml.

### Starting the Services

1.  Navigate to the platform directory:
    ```bash
    cd platform
    ```
2.  Start the core infrastructure (Postgres, Redis):
    ```bash
    docker compose -p anthropos-rosetta up -d postgresql redis
    ```
    *Verification*: `docker ps` should show `anthropos-rosetta-postgresql-1` and `anthropos-rosetta-redis-1` containers running.

3.  Start the services:
    ```bash
    docker compose -p anthropos-rosetta up -d backend cms jobsimulation
    ```
    *(Note: Sentinel and others will be pulled as images if you didn't build them).*

4.  **Verification**:
    Run `docker ps`. You should see healthy containers: `anthropos-rosetta-backend-1`, `anthropos-rosetta-cms-1`, `anthropos-rosetta-postgresql-1`, etc.

---

## 7. Running Frontend

1.  Navigate to the frontend repo:
    ```bash
    cd ../next-web-app
    ```
2.  Install dependencies:
    ```bash
    pnpm install
    ```
3.  Run the development server:
    ```bash
    pnpm dev
    ```
4.  Open `http://localhost:3000` (or the port shown in terminal).

---

## 8. Running Studio Services

### Studio-Desk (Design Tool)
1.  Navigate to studio-desk:
    ```bash
    cd ../studio/studio-desk
    ```
2.  Install dependencies & start:
    ```bash
    npm install
    npm run dev
    ```
3.  Access at `http://localhost:3100`

### Studio-Room (AI Pipeline)
1.  Navigate to studio-room:
    ```bash
    cd ../studio-room
    ```
2.  Install requirements:
    ```bash
    pip3 install -r requirements.txt
    ```
3.  Run a test generation:
    ```bash
    python3 gen.py --media simulation --template default
    ```

---

## 9. Troubleshooting

### "Generated code missing" / "command not found: make"
If running Go services locally (outside Docker), you may hit errors about missing files.
*   **MacOS**: Ensure `xcode-select --install` is run.
*   **Linux**: Ensure `sudo apt-get install build-essential`.

Resolution:
```bash
cd [service-name]
make setup
make gen
```

### "Connection Refused" / Docker Issues
*   **General**: Ensure Docker containers are running (`docker ps`). If a service in Docker is failing, check logs: `docker compose logs -f [service_name]`.
*   **Linux Permission Denied**: If you see "permission denied while trying to connect to the Docker daemon", you likely skipped the `usermod` step. Run `sudo usermod -aG docker $USER`, then log out and back in (or `newgrp docker`).

---

## 10. Maintenance Guidelines

This `setup_guide.md`, the OS-specific checklists (`setup_checklist_macos.md`, `setup_checklist_linux.md`), and the `/anthropos-setup` Claude skill are interconnected documents that must be maintained together.

### When You Update This Setup Guide

If you modify the setup process (add/remove/reorder steps), you must update:

1.  **Setup Checklists** (`setup_checklist.md`): Add, remove, or reorder checkboxes to match the guide structure
2.  **Anthropos Setup Skill** (`.claude/skills/anthropos-setup/SKILL.md`): Update phase definitions, step sequences, and verification commands
3.  **This Guide**: Ensure all steps have verification commands documented

### Checklist Usage Pattern

The checklist is for **progress tracking**, not detailed instruction:

*   **User Workflow**: Copy the checklist to your `anthropos-dev/` workspace (e.g., `anthropos-dev/my_setup_progress.md`)
*   **Track Progress**: Check off items `[x]` as you complete them in YOUR local copy
*   **Resume Setup**: Use your local checklist to resume where you left off
*   **Report Issues**: Use the "Notes / Errors" table in your local copy to document problems for other developers
*   **Keep Checklists Lean**: The original checklist in `corpus/setup/` is only updated when the setup guide structure changes

### General Guidelines

*   **OS-Specific Differences**: When a step differs between macOS and Linux, ensure each checklist reflects the appropriate commands/tools for that OS
*   **Agent-Friendly**: Ensure all documents remain parseable and clear for autonomous agents
*   **Verification Commands**: Every installation step should have a documented verification command

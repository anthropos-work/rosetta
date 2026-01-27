# Zero to Hero: Anthropos Development Setup Guide

This guide is designed to take you from a **fresh computer** (or a clean folder) to a fully running Anthropos development environment.

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
5.  **Node.js** (v20 LTS - **not v22+**) & **pnpm**:
    *   **Recommended**: Use [nvm](https://github.com/nvm-sh/nvm) to manage Node versions:
        ```bash
        curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
        source ~/.nvm/nvm.sh
        nvm install 20
        nvm use 20
        ```
    *   Alternative: `brew install node@20` (not `brew install node` which installs latest)
    *   `corepack enable` or `npm install -g pnpm`
    *   *Verification*: `node --version && pnpm --version`
    *   **Warning**: Node.js v22+ removed `import ... assert` syntax used by the frontend. Use v20 LTS.
6.  **Build Tools**:
    *   Ensure XCode CLI tools are installed: `xcode-select --install`
    *   Ensure XCode CLI tools are installed: `xcode-select --install`
    *   *Verification*: `xcode-select -p`
7.  **Python** (v3.8+ for Studio-Room):
    *   `brew install python`
    *   *Verification*: `python3 --version`
8.  **Atlas** (Database Schema Manager):
    *   `curl -sSf https://atlasgo.sh | sh`
    *   *Verification*: `atlas version`

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
5.  **Node.js** (v20 LTS - **not v22+**) & **pnpm**:
    *   **Recommended**: Use [nvm](https://github.com/nvm-sh/nvm) to manage Node versions:
        ```bash
        curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
        source ~/.nvm/nvm.sh
        nvm install 20
        nvm use 20
        ```
    *   `corepack enable` or `npm install -g pnpm`
    *   **Warning**: Node.js v22+ removed `import ... assert` syntax used by the frontend. Use v20 LTS.
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
*   Track progress using TodoWrite
*   Create ops reports when issues are discovered

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
# Remote-only services (optional, clone only if editing source or applying migrations):
# git clone git@github.com:anthropos-work/sentinel.git
# git clone git@github.com:anthropos-work/skiller.git
# git clone git@github.com:anthropos-work/skillpath.git
# git clone git@github.com:anthropos-work/chronos.git
```
*Verification*: `ls -la backend cms jobsimulation` should show all three directories with Go files.

### Frontend
Clone the Next.js monorepo.
```bash
git clone git@github.com:anthropos-work/next-web-app.git
```
*Verification*: `ls -la next-web-app` should show Next.js project files including `package.json`.

### Studio Services
Clone the Studio-Desk (design tool) and Studio-Room (AI pipeline) repositories.
```bash
git clone git@github.com:anthropos-work/studio-desk.git
git clone git@github.com:anthropos-work/anthropos-studio-room.git studio-room
```
*Verification*: `ls -la studio-desk studio-room` should show both directories.

**CMS Dependency**: The CMS service requires studio-room to be present within its build context. Since Docker does not follow symlinks outside the build context, we must clone the repository directly:
```bash
git clone git@github.com:anthropos-work/anthropos-studio-room.git cms/studio
```
*Verification*: `ls -la cms/studio` should show the studio-room files (e.g., `gen.py`, `requirements.txt`).

### Internal Tools (Optional)
Clone the internal experiments hub for access to PoCs, prototypes, and internal tools.
```bash
git clone git@github.com:anthropos-work/experiments.git
```
*Verification*: `ls -la experiments` should show the experiments hub files (e.g., `package.json`, `vite.config.js`).

See [Anthropos Labs documentation](../tools/anthropos-labs.md) for usage details.

### Understanding Docker's Build-From-Git Architecture

**Most backend services do NOT need local clones to run.** The `docker-compose.yml` builds them directly from GitHub using SSH authentication:

| Service | Builds From | Local Clone Needed? |
|---------|-------------|---------------------|
| `graphql` | `git@github.com:anthropos-work/graphql-wundergraph.git` | No |
| `sentinel` | `git@github.com:anthropos-work/sentinel.git` | No |
| `backend` | `git@github.com:anthropos-work/app.git` | No |
| `skiller` | `git@github.com:anthropos-work/skiller.git` | No |
| `skillpath` | `git@github.com:anthropos-work/skillpath.git` | No |
| `storage` | `git@github.com:anthropos-work/storage.git` | No |
| `chronos` | `git@github.com:anthropos-work/chronos.git` | No |
| `intelligence` | `git@github.com:anthropos-work/intelligence.git` | No |
| `cms` | Local (`../cms`) | **Yes** |

**When to clone a repo:**
- You need to apply database migrations (Atlas requires local migration files)
- You're developing on that service (editing source code)
- You need to run the service outside Docker

**Why this matters:** When you run `docker compose up`, Docker will:
1. Clone the repo from GitHub (if not cached)
2. Build the image using the Dockerfile in that repo
3. Start the container

This requires your SSH agent to be running with GitHub access: `ssh-add -l`

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

3.  **Verification**: `ls -la platform/.env` should show the file exists.

### Studio-Desk Environment

Studio-Desk requires its own `.env` file with Clerk and OpenAI credentials.

1.  **Create the environment file**:
    ```bash
    cd ../studio-desk
    cp .env.example .env
    ```

2.  **Populate required keys** from `platform/.env`:

    Open `studio-desk/.env` and set:
    ```
    # Server
    CLERK_SECRET_KEY=<from platform/.env>

    # Frontend
    VITE_CLERK_PUBLISHABLE_KEY=<from platform/.env CLERK_PUBLISHABLE_KEY>
    VITE_GRAPHQL_ENDPOINT=http://localhost:5050/graphql

    # OpenAI (for Copilot)
    OPENAI_API_KEY=<from platform/.env>
    ```

3.  **Verification**: `ls -la studio-desk/.env` should show the file exists.

**Note**: The docker-compose configuration uses this single `.env` file for all services (backend, cms, jobsimulation, etc.). Individual service repositories do not need their own `.env` files when running via Docker.

---

## 6. Running the Platform (Docker)

The easiest way to start is using Docker Compose.

### Docker Compose Project Name

We use the `-p ant-rosetta` flag to set a custom project name. This creates an isolated Docker stack that won't conflict with other Anthropos environments you may have running.

**What this does:**
- Creates containers named `ant-rosetta-postgresql-1`, `ant-rosetta-backend-1`, etc.
- Creates isolated networks: `ant-rosetta_app-network`
- Creates isolated volumes: `ant-rosetta_postgres_data`

**Note**: If you have another Anthropos stack running (e.g., "platform"), they will be completely separate. However, you may encounter **port conflicts** if both stacks try to use the same ports. Stop the other stack first or modify port mappings in docker-compose.yml.

### Starting the Services

1.  Navigate to the platform directory:
    ```bash
    cd platform
    ```
2.  Start the core infrastructure (Postgres, Redis):
    ```bash
    docker compose -p ant-rosetta up -d postgresql redis
    ```
    *Verification*: `docker ps` should show `ant-rosetta-postgresql-1` and `ant-rosetta-redis-1` containers running.

3.  **Prepare PostgreSQL Extensions** (required before migrations):

    Wait for PostgreSQL to be healthy, then create the required schemas and extensions:
    ```bash
    # Wait for PostgreSQL to be ready
    sleep 5

    # Create extensions schema with pgvector (required by CMS and Skiller)
    docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions; CREATE EXTENSION IF NOT EXISTS pg_trgm SCHEMA extensions;"

    # Create sentinel schema (required by Sentinel for Casbin authorization)
    docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS sentinel;"
    ```
    *Verification*: Commands should complete without errors.

4.  **Initialize Database Schemas**:
    The Postgres database starts empty. You must create the schemas for the core services (`backend`, `cms`, `jobsimulation`) using Atlas.
    
    *   **Install Atlas** (if you skipped it in Prerequisites):
        ```bash
        curl -sSf https://atlasgo.sh | sh
        ```
    *   **Apply Migrations**:
        Run the following commands from the `anthropos-dev` directory (where you cloned the repos):
        ```bash
        # Backend Schema (public)
        (cd backend && atlas migrate apply --env local)
        
        # CMS Schema (cms)
        (cd cms && atlas migrate apply --env local)
        
        # JobSimulation Schema (jobsimulation)
        (cd jobsimulation && atlas migrate apply --env local)

        # Skiller Schema (skiller)
        (cd skiller && atlas migrate apply --env local)

        # Skillpath Schema (skillpath)
        (cd skillpath && atlas migrate apply --env local)

        # Chronos Schema (chronos)
        (cd chronos && atlas migrate apply --env local)
        ```
        *Note: The parenthesis `(...)` ensure you return to the current directory after the command.*
        *Note: Skillpath and Chronos services are run as Docker images. Only clone these repos if you need to run migrations or develop on them.*
    *   **Verification**: The commands should complete successfully without error, outputting the migration steps applied.

5.  Start all backend services using the `graphql` profile:
    ```bash
    docker compose -p ant-rosetta --profile graphql up -d
    ```
    This starts: `sentinel`, `backend`, `cms`, `skiller`, `skillpath`, `storage`, `chronos`, `jobsimulation`, `intelligence`, and `graphql`.

    *Note*: First run may take several minutes as Docker builds images from GitHub. Ensure your SSH agent is running (`ssh-add -l`).

6.  **Verification**:
    Run `docker ps --filter "name=ant-rosetta" --format "table {{.Names}}\t{{.Status}}"`. You should see all services running:
    - `ant-rosetta-postgresql-1` (healthy)
    - `ant-rosetta-redis-1` (healthy)
    - `ant-rosetta-sentinel-1`
    - `ant-rosetta-backend-1`
    - `ant-rosetta-cms-1`
    - `ant-rosetta-graphql-1`
    - And others...

### Ongoing Operations

For daily platform operations (starting, stopping, rebuilding services), see the [Run Guide](run_guide.md).

---

## 7. Running Frontend

The Next.js frontend is a monorepo with multiple apps. Each app needs its own `.env` file.

### Configure Environment Files

1.  Navigate to the frontend repo:
    ```bash
    cd ../next-web-app
    ```

2.  **Create `.env` files** for each app you want to run:
    ```bash
    # Main web app (required)
    cp apps/web/.env.example apps/web/.env

    # Hiring app (optional)
    cp apps/hiring/.env.example apps/hiring/.env

    # Integration app (optional)
    cp apps/integration/.env.example apps/integration/.env
    ```

3.  **Populate Clerk keys** from `platform/.env`:

    Open `apps/web/.env` and set:
    ```
    NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=<from platform/.env CLERK_PUBLISHABLE_KEY>
    CLERK_SECRET_KEY=<from platform/.env CLERK_SECRET_KEY>
    ```

    *Note*: The GraphQL and Backend URLs already default to `localhost:5050` and `localhost:8082` which are correct for local development.

    *Verification*: `ls apps/web/.env` should show the file exists.

### Install and Run

4.  Install dependencies:
    ```bash
    pnpm install
    ```

5.  Run the development server:
    ```bash
    pnpm dev
    ```
    This starts all apps in the monorepo (web, hiring, integration).

6.  Open `http://localhost:3000` (or the port shown in terminal).
    - Web app: http://localhost:3000
    - Hiring app: http://localhost:3001
    - Integration app: http://localhost:3002

---

## 8. Running Studio-Desk (Design Tool)

Studio-Desk is the simulation design tool - a required part of the full platform for content creation workflows.

1.  Navigate to studio-desk:
    ```bash
    cd ../studio-desk
    ```

2.  Install dependencies:
    ```bash
    npm install
    ```

3.  Start the development server:
    ```bash
    npm run dev
    ```
    This starts both the frontend (port 9100) and backend (port 9000). Ports are configurable via `.env`.

4.  Access at `http://localhost:9100`

    *Verification*: You should see the Studio-Desk login page (uses Clerk authentication).

---

## 9. Running Studio-Room (Optional - On-Demand)

Studio-Room is the AI generation pipeline. It runs **on-demand** for specific generation tasks, not as a persistent service.

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

## 10. Troubleshooting

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

### "SyntaxError: Unexpected identifier 'assert'" (Frontend)
The frontend uses `import ... assert { type: 'json' }` syntax which was removed in Node.js v22+.
*   **Solution**: Use Node.js v20 LTS. Install via nvm:
    ```bash
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
    source ~/.nvm/nvm.sh
    nvm install 20
    nvm use 20
    ```
*   Then run the frontend: `cd next-web-app && pnpm dev`

### "schema 'extensions' does not exist" (Atlas migrations)
CMS and Skiller services require the pgvector extension for vector embeddings.
*   **Solution**: Create the extensions schema and install pgvector:
    ```bash
    docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS extensions; CREATE EXTENSION IF NOT EXISTS vector SCHEMA extensions;"
    ```
*   Then retry the migrations: `cd [service] && atlas migrate apply --env local`

### Sentinel Crashing / Restarting
Sentinel requires its own database schema for Casbin authorization.
*   **Solution**: Create the sentinel schema:
    ```bash
    docker exec ant-rosetta-postgresql-1 psql -U postgres -c "CREATE SCHEMA IF NOT EXISTS sentinel;"
    ```
*   Then restart sentinel: `docker restart ant-rosetta-sentinel-1`

### Port Already In Use
If you have another Docker stack running (e.g., "platform"), ports may conflict.
*   **Solution**: Stop the other stack first:
    ```bash
    docker compose -p platform stop
    ```

### Docker Build Fails with "Permission denied (publickey)"
Docker builds services from GitHub and needs SSH access.
*   **Solution**: Ensure your SSH agent is running with keys loaded:
    ```bash
    # Check if SSH agent has keys
    ssh-add -l

    # If "no identities" or "agent not running", start it and add your key:
    eval "$(ssh-agent -s)"
    ssh-add ~/.ssh/id_ed25519

    # Verify GitHub access
    ssh -T git@github.com
    ```
*   Then retry: `docker compose -p ant-rosetta --profile graphql up -d --build`

---

## 11. Maintenance Guidelines

This `setup_guide.md` and the `/ant-setup` Claude skill are interconnected documents that must be maintained together.

### When You Update This Setup Guide

If you modify the setup process (add/remove/reorder steps), you must update:

1.  **Anthropos Setup Skill** (`.claude/skills/ant-setup/SKILL.md`): Update phase definitions, step sequences, and verification commands
2.  **This Guide**: Ensure all steps have verification commands documented

### General Guidelines

*   **OS-Specific Differences**: When a step differs between macOS and Linux, ensure this guide reflects the appropriate commands/tools for each OS
*   **Agent-Friendly**: Ensure all documents remain parseable and clear for autonomous agents
*   **Verification Commands**: Every installation step should have a documented verification command

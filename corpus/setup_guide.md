# Zero to Hero: Anthropos Development Setup Guide

This guide is designed to take you from a **fresh computer** (or a clean folder) to a fully running Anthropos development environment.

## 1. Prerequisites

Before we write any code, ensure you have the following tools installed.

### Essential Tools

### Essential Tools

#### 1. OS-Specific Setup

<details open>
<summary><strong>MacOS</strong></summary>

We recommend using [Homebrew](https://brew.sh/) for package management.

1.  **Git**: `brew install git`
2.  **Docker Desktop**: [Install Docker Desktop for Mac](https://www.docker.com/products/docker-desktop/).
3.  **Visual Studio Code**: [Install VS Code](https://code.visualstudio.com/).
4.  **Go** (v1.23+): `brew install go`
    *   *Verification*: `go version`.
5.  **Node.js** (v20+) & **pnpm**:
    *   `brew install node`
    *   `corepack enable` or `npm install -g pnpm`.
6.  **Build Tools**:
    *   Ensure XCode CLI tools are installed: `xcode-select --install`

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

</details>

---

## 2. Workspace Setup

We will create a dedicated workspace to house all the microservices and the frontend.

1.  Open your terminal.
2.  Navigate to the `rosetta` directory in this repository.
3.  Enter the scratchpad directory:
    ```bash
    cd anthropos-dev
    ```
    *(Note: This directory is git-ignored, so you can clone anything here without messing up the main repo).*

---

## 3. Cloning Repositories

We need to fetch the code for the platform. Run the following commands inside `rosetta/anthropos-dev`:

### Platform Configuration
This repository contains the `docker-compose.yml` to orchestrate everything.
```bash
git clone git@github.com:anthropos-work/platform.git
```

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

### Frontend
Clone the Next.js monorepo.
```bash
git clone git@github.com:anthropos-work/next-web-app.git
```

---

## 4. Environment Configuration

### The `.env` Files
Most services require an `.env` file to function.
> **IMPORTANT**: You must obtain the master `.env` values from the Engineering Manager or the 1Password Vault "Engineering/Env".

1.  **Backend**: Copy `.env.example` to `backend/.env` and populate secrets.
2.  **CMS**: Copy `.env.example` to `cms/.env`.
3.  **Platform**: Check `platform/.env`.

---

## 5. Running the Platform (Docker)

The easiest way to start is using Docker Compose.

1.  Navigate to the platform directory:
    ```bash
    cd platform
    ```
2.  Start the core infrastructure (Postgres, Redis):
    ```bash
    docker compose up -d postgresql redis
    ```
3.  Start the services:
    ```bash
    docker compose up -d backend cms jobsimulation
    ```
    *(Note: Sentinel and others will be pulled as images if you didn't build them).*

4.  **Verification**:
    Run `docker ps`. You should see healthy containers for `app` (backend), `cms`, and `postgres`.

---

## 6. Running Frontend

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

## Troubleshooting

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

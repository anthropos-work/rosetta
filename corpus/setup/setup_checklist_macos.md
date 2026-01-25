# Anthropos Development Setup Checklist (macOS)

This file is a companion to the `setup_guide.md`. Use it to track your progress when setting up a new macOS machine or environment.

**Instructions:**
1.  Copy this file to your working directory (e.g., `anthropos-dev/setup_progress.md`).
2.  Check off items `[x]` as you complete them.
3.  Use the "Notes / Errors" section to log any issues for future reference or debugging.

---

## 1. Prerequisites (Tools)

- [ ] **Homebrew** installed
- [ ] **Git** installed (`brew install git`)
    - [ ] Verified with `git --version`
- [ ] **Docker Desktop** installed & running
    - [ ] Verified with `docker --version && docker compose version`
- [ ] **Visual Studio Code** installed
    - [ ] Verified with `code --version`
- [ ] **Go** (v1.23+) installed (`brew install go`)
    - [ ] Verified with `go version`
- [ ] **Node.js** (v20+) installed (`brew install node`)
    - [ ] **pnpm** installed (`corepack enable` or `npm i -g pnpm`)
    - [ ] Verified with `node --version && pnpm --version`
- [ ] **XCode CLI Tools** installed (`xcode-select --install`)
    - [ ] Verified with `xcode-select -p`
- [ ] **Atlas** (Database Schema Manager) installed (`curl -sSf https://atlasgo.sh | sh`)
    - [ ] Verified with `atlas version`

## 2. GitHub SSH Access

- [ ] SSH keys configured
    - [ ] Check with `ls -al ~/.ssh` for `id_rsa.pub` or `id_ed25519.pub`
    - [ ] Generate if needed: `ssh-keygen -t ed25519 -C "your_email@example.com"`
    - [ ] Add to ssh-agent: `ssh-add --apple-use-keychain ~/.ssh/id_ed25519`
- [ ] SSH key added to GitHub account (Settings > SSH and GPG keys)
- [ ] SSH connection verified: `ssh -T git@github.com`
- [ ] Access to `anthropos-work` organization confirmed

## 3. Workspace Setup

- [ ] Created/Navigated to workspace directory (`anthropos-dev`)
    - *Note: Ensure this is the scratchpad dir ignored by git.*

## 4. Cloning Repositories

- [ ] **Platform** (`anthropos-work/platform`) cloned
    - [ ] Verified: `ls -la platform` shows `docker-compose.yml`
- [ ] **Backend Services**
    - [ ] `app` repo cloned as `backend` directory
    - [ ] `cms` repo cloned
    - [ ] `jobsimulation` repo cloned
    - [ ] `skiller` repo cloned
    - [ ] Verified: `ls -la backend cms jobsimulation skiller` shows Go files
- [ ] **Frontend**
    - [ ] `next-web-app` repo cloned
    - [ ] Verified: `ls -la next-web-app` shows `package.json`
- [ ] **Internal Tools (Optional)**
    - [ ] `experiments` repo cloned (Anthropos Labs)
    - [ ] Verified: `ls -la experiments` shows `package.json`, `vite.config.js`

## 5. Environment Configuration

- [ ] **Platform `.env`** created
    - [ ] Copied from `platform/.env_example` to `platform/.env`
    - [ ] Secrets populated from 1Password Vault "Engineering/Env" or Engineering Manager
    - [ ] Verified: `ls -la platform/.env` shows the file exists
    - *Note: All services share this SINGLE centralized .env file*
- [ ] **CMS `studio/` folder** obtained and placed
    - [ ] Contact Engineering Manager or check 1Password for studio folder
    - [ ] Place folder at `cms/studio/`
    - [ ] Verified: `ls -la cms/studio/requirements.txt` exists
    - *Note: This folder is gitignored and required for CMS Python AI components*

## 6. Running the Platform (Docker)

- [ ] Docker Compose: Core Infra started
    - [ ] Command: `docker compose -p anthropos-rosetta up -d postgresql redis`
    - [ ] Verified: `docker ps` shows `anthropos-rosetta-postgresql-1` and `anthropos-rosetta-redis-1` running
- [ ] **Database Schemas Initialized**
    - [ ] Backend schema applied: `(cd backend && atlas migrate apply --env local)`
    - [ ] CMS schema applied: `(cd cms && atlas migrate apply --env local)`
    - [ ] JobSimulation schema applied: `(cd jobsimulation && atlas migrate apply --env local)`
    - [ ] Skiller schema applied: `(cd skiller && atlas migrate apply --env local)`
- [ ] Docker Compose: Services started
    - [ ] Command: `docker compose -p anthropos-rosetta up -d backend cms jobsimulation`
    - [ ] Verified: `docker ps` shows healthy containers for backend, cms, jobsimulation
    - *Note: Using `-p anthropos-rosetta` creates an isolated stack that won't conflict with other environments*

## 7. Running Frontend

- [ ] Dependencies installed
    - [ ] Command: `pnpm install` in `next-web-app` directory
- [ ] Dev server running
    - [ ] Command: `pnpm dev`
    - [ ] Accessed `http://localhost:3000` successfully

---

## Notes / Errors

| Date | Step | Issue / Error Message | Resolution |
| :--- | :--- | :--- | :--- |
|      |      |                       |            |

---

## Status Summary

**Completed:**
-

**Blocked/Pending:**
-

**Next Steps:**
1.

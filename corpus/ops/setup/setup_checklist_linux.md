# Anthropos Development Setup Checklist (Linux)

This file is a companion to the `setup_guide.md`. Use it to track your progress when setting up a new Linux machine or environment.

**Instructions:**
1.  Copy this file to your working directory (e.g., `anthropos-dev/my_setup_checklist.md`).
2.  Check off items `[x]` as you complete them.
3.  Use the "Notes / Errors" section to log any issues for future reference or debugging.

---

## 1. Prerequisites (Tools)

- [ ] **System Update** (`sudo apt-get update`)
- [ ] **Git** installed (`sudo apt-get install -y git`)
- [ ] **Build Essential** installed (`sudo apt-get install -y build-essential`)
- [ ] **Docker Engine** installed
    - [ ] Added user to `docker` group (`sudo usermod -aG docker $USER`)
    - [ ] Verified with `docker ps` (no sudo required)
- [ ] **Visual Studio Code** installed
- [ ] **Go** (v1.23+) installed (from official source)
    - [ ] Verified with `go version`
- [ ] **Node.js** (v20+) installed (via nodesource or nvm)
- [ ] **pnpm** installed (`corepack enable` or `npm i -g pnpm`)
- [ ] **Atlas** (Database Schema Manager) installed (`curl -sSf https://atlasgo.sh | sh`)
    - [ ] Verified with `atlas version`

## 2. Workspace Setup

- [ ] Created/Navigated to workspace directory (`anthropos-dev`)
    - *Note: Ensure this is the scratchpad dir ignored by git.*

## 3. Cloning Repositories

- [ ] **Platform** (`anthropos-work/platform`) cloned
- [ ] **Backend Services**
    - [ ] `app` repo cloned as `backend` directory
    - [ ] `cms` repo cloned
    - [ ] `jobsimulation` repo cloned
    - [ ] `skiller` repo cloned
- [ ] **Frontend**
    - [ ] `next-web-app` repo cloned
- [ ] **Internal Tools (Optional)**
    - [ ] `experiments` repo cloned (Anthropos Labs)

## 4. Environment Configuration

- [ ] **Backend** `.env` created (from `.env.example` + secrets)
- [ ] **CMS** `.env` created (from `.env.example`)
- [ ] **Platform** `.env` verified

## 5. Running the Platform

- [ ] Docker Compose: Core Infra (`postgresql`, `redis`) up
- [ ] **Database Schemas Initialized** (after Core Infra is up)
    - [ ] Backend schema applied: `(cd backend && atlas migrate apply --env local)`
    - [ ] CMS schema applied: `(cd cms && atlas migrate apply --env local)`
    - [ ] JobSimulation schema applied: `(cd jobsimulation && atlas migrate apply --env local)`
    - [ ] Skiller schema applied: `(cd skiller && atlas migrate apply --env local)`
- [ ] Docker Compose: Services (`backend`, `cms`, `jobsimulation`) up
- [ ] Verification: `docker ps` shows healthy containers

## 6. Running Frontend

- [ ] Dependencies installed (`pnpm install` in `next-web-app`)
- [ ] Dev server running (`pnpm dev`)
- [ ] Accessed `http://localhost:3000` successfully

---

## Notes / Errors

| Date | Step | Issue / Error Message | Resolution |
| :--- | :--- | :--- | :--- |
|      |      |                       |            |

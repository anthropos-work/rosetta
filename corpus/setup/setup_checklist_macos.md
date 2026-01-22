# Anthropos Development Setup Checklist (macOS)

This file is a companion to the `setup_guide.md`. Use it to track your progress when setting up a new macOS machine or environment.

**Instructions:**
1.  Copy this file to your working directory (e.g., `anthropos-dev/my_setup_checklist.md`).
2.  Check off items `[x]` as you complete them.
3.  Use the "Notes / Errors" section to log any issues for future reference or debugging.

---

## 1. Prerequisites (Tools)

- [ ] **Homebrew** installed
- [ ] **Git** installed (`brew install git`)
- [ ] **Docker Desktop** installed & running
- [ ] **Visual Studio Code** installed
- [ ] **Go** (v1.23+) installed (`brew install go`)
    - [ ] Verified with `go version`
- [ ] **Node.js** (v20+) installed (`brew install node`)
- [ ] **pnpm** installed (`corepack enable` or `npm i -g pnpm`)
- [ ] **XCode CLI Tools** installed (`xcode-select --install`)

## 2. Workspace Setup

- [ ] Created/Navigated to workspace directory (`anthropos-dev`)
    - *Note: Ensure this is the scratchpad dir ignored by git.*

## 3. Cloning Repositories

- [ ] **Platform** (`anthropos-work/platform`) cloned
- [ ] **Backend Services**
    - [ ] `app` repo cloned as `backend` directory
    - [ ] `cms` repo cloned
    - [ ] `jobsimulation` repo cloned
- [ ] **Frontend**
    - [ ] `next-web-app` repo cloned

## 4. Environment Configuration

- [ ] **Backend** `.env` created (from `.env.example` + secrets)
- [ ] **CMS** `.env` created (from `.env.example`)
- [ ] **Platform** `.env` verified

## 5. Running the Platform

- [ ] Docker Compose: Core Infra (`postgresql`, `redis`) up
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

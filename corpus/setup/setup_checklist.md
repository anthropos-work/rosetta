# Anthropos Development Setup Checklist

This file is a companion to the `setup_guide.md`. Use it to track your progress when setting up a new machine or environment.

**Instructions:**
1.  **Copy this file** to your working directory: `cp corpus/setup/setup_checklist.md anthropos-dev/setup_progress.md`
2.  **Work from your copy**: All progress tracking happens in `anthropos-dev/setup_progress.md` (not this original file)
3.  **Check off items** `[x]` as you complete them in YOUR copy
4.  **Log issues**: Use the "Notes / Errors" table in your copy to document problems
5.  **Resume setup**: Use your copy to pick up where you left off
6.  **Report back**: Share your copy with other developers to report issues

**Note**: This original file in `corpus/setup/` is only modified when the setup guide structure changes (new/removed/reordered steps). Your personal tracking happens in your copy.

---

## 1. Prerequisites (Tools)

- [ ] **OS Setup**
    - [ ] MacOS: Homebrew & XCode CLI installed
    - [ ] Linux: `build-essential` & `git` installed
    - [ ] Python v3.8+ installed (for Studio-Room)
- [ ] **Git** installed
- [ ] **Docker Desktop / Engine** installed & running
    - [ ] Linux users: Added user to `docker` group (`docker ps` works without sudo)
- [ ] **Visual Studio Code** installed
- [ ] **Go** (v1.23+) installed
    - [ ] Verified with `go version`
- [ ] **Node.js** (v20+) installed
- [ ] **pnpm** installed (`corepack enable` or `npm i -g pnpm`)

## 2. Workspace Setup

- [ ] Created/Navigated to workspace directory (`anthropos-dev`)
    - *Note: Ensure this is the scratchpad dir ignored by git.*

## 3. Cloning Repositories

- [ ] **Platform** (`anthropos-work/platform`) cloned
- [ ] **Backend Services**
    - [ ] `app` repo cloned as `backend` directory
    - [ ] `cms` repo cloned
    - [ ] `jobsimulation` repo cloned
- [ ] **Studio Services**
    - [ ] `studio` repo cloned (includes `studio-desk` & `studio-room`)
- [ ] **Frontend**
    - [ ] `next-web-app` repo cloned

## 4. Environment Configuration

- [ ] **Backend** `.env` created (from `.env.example` + secrets)
    - [ ] **Secrets Populated**: Clerk keys, OpenAI/Anthropic keys, Directus URL
- [ ] **Studio-Desk** `.env` created (from `.env.example`)
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

## 7. Running Studio Services

- [ ] **Studio-Desk**: Dependencies installed (`npm install`)
- [ ] **Studio-Desk**: Dev server running (`npm run dev`)
- [ ] **Studio-Room**: Python requirements installed (`pip install -r requirements.txt`)
- [ ] **Studio-Room**: Test generation successful (`python gen.py ...`)

---

## Notes / Errors

| Date | Step | Issue / Error Message | Resolution |
| :--- | :--- | :--- | :--- |
|      |      |                       |            |

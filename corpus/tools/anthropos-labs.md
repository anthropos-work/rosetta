# Anthropos Labs (Experiments Hub)

Internal experiments hub for the Anthropos team. Hosts PoCs, UI prototypes, and internal tools. **Not part of the main platform** - this is a sandbox for team experimentation.

## Quick Facts

| Aspect | Value |
|--------|-------|
| **Repo** | `anthropos-work/experiments` |
| **Tech Stack** | Vite, Vanilla JS/HTML, Clerk Auth |
| **Hosting** | Vercel (automated CI/CD) |
| **Access** | `@anthropos.work` emails only |
| **Local Port** | 3002 |

## What's Inside

The experiments hub contains various PoCs and internal tools:

| Experiment | Purpose |
|------------|---------|
| `simulator-01/` | Job simulator prototype v1 |
| `simulator-02/` | Job simulator prototype v2 |
| `ant-library-02/` | UI component library experiments |
| `ant-library-cards/` | Card-based UI experiments |
| `tool-icon-picker/` | Icon selection tool |
| `labs/` | Miscellaneous experiments |

## Local Development

### Prerequisites

- Node.js v14+
- Clerk publishable key (for authentication)

### Setup

```bash
# Navigate to experiments directory
cd anthropos-dev/experiments

# Install dependencies
npm install

# Create environment file
cp .env.example .env
# Edit .env and add: VITE_CLERK_PUBLISHABLE_KEY=pk_test_your_key_here

# Start dev server
npm run dev
```

The hub runs at `http://localhost:3002`.

### Verification

```bash
# Check server is running
curl -s http://localhost:3002 | head -5
```

You should see the HTML for the experiments hub (authentication will be required to access experiments).

## Adding New Experiments

1. **Create directory**: `mkdir experiments/your-experiment-name`
2. **Create HTML file** with Clerk auth integration (see `simulator-01/` for template)
3. **Register in hub**: Add entry to `experiments` array in `index.html`
4. **Deploy**: Push to main branch (auto-deploys to Vercel)

### Experiment Template

Each experiment should:
- Import auth from `../src/auth.js`
- Call `authManager.requireAuth()` before initializing
- Include "Back to Hub" navigation
- Mount Clerk user button for sign-out

## Architecture Notes

### Authentication Flow

```
User visits experiment
    ↓
Clerk checks authentication
    ↓
Domain check (@anthropos.work?)
    ↓
[Yes] → Show experiment
[No]  → Show 403 error
```

### Dependencies

- `@clerk/clerk-js` - Authentication
- `monaco-editor` - Code editor (for code-heavy experiments)
- `marked` - Markdown rendering

## When to Use This

**Use Anthropos Labs for:**
- UI/UX prototypes before platform integration
- Testing new libraries or approaches
- Internal tools that don't fit the platform
- Demos and PoCs for stakeholders

**Don't use for:**
- Production features (use platform services)
- Customer-facing functionality
- Anything requiring platform data access

## Related

- [Toolchain Overview](./toolchain_overview.md) - Development tools registry
- [Studio-Desk](../services/studio-desk.md) - Production content design tool

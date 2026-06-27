# iter-02 — decisions (tik, TOK-01 line 1)

**D1 — The next-web Studio left-nav link is NOT env-rewritable to the demo-local studio-desk without a
platform-repo edit → RE-SCOPE TRIGGER.** Full evidence chain (all read-only; zero platform edit):

- **Source:** `STUDIO_URL` = `packages/core-js/src/constants/urls.ts:12-15`:
  ```ts
  export const STUDIO_URL =
    process.env.NEXT_PUBLIC_NODE_ENV === 'development'
      ? 'http://localhost:9000'
      : 'https://studio.anthropos.work';
  ```
  A `NEXT_PUBLIC_NODE_ENV`-gated ternary. **No `NEXT_PUBLIC_STUDIO_URL` override** exists — contrast
  `ACADEMY_URL` (line 16-17) which reads `process.env.NEXT_PUBLIC_ACADEMY_URL || '…'` (the override pattern
  ant-academy's demo uses to rewrite ITS Studio link — `ant-academy.sh:108` sets `NEXT_PUBLIC_STUDIO_URL`).
  next-web simply lacks the equivalent read.
- **Rendered link:** `useNavbarSections.tsx:283` — `studioMenuItem` `{ key: STUDIO_URL, type:
  MenuType.outbound }` — the left-nav "Studio" outbound link the manager (enterprise) nav shows on every page.
- **Live (demo-3):** `docker exec demo-3-next-web-app-1` → `NEXT_PUBLIC_NODE_ENV=[]` (unset),
  `NODE_ENV=production`; the built bundle (`/app/apps/web/.next`) carries `studio.anthropos.work` (grep-confirmed).
- **Demo build:** `up-injected.sh::build_frontend_next_web` bakes only `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` /
  `_BACKEND_API_URL` / `_HOSTING_URL` (the only ARGs `Dockerfile.dev` declares) + a gitignored
  `apps/web/.env.local` (pk). `NEXT_PUBLIC_NODE_ENV` is never set → the `else` (prod) branch.
- **Why the only knob fails:** setting `NEXT_PUBLIC_NODE_ENV=development` in the `.env.local` overlay would
  (a) point Studio at `:9000`, but the demo studio-desk is on the OFFSET port `:39000`
  (`0.0.0.0:39000->9000/tcp`) → still a broken/wrong link, not demo-local; and (b) ALSO flip `WEB_APP_URL`
  (→`:3000`) + `HIRING_APP_URL` (→`:3001`) + `TablePagination`, producing NEW wrong-port links across
  manager-reachable surfaces (`PublicFooter`, `skillPathRoute` public URLs, `ShareAISim`, PDF download). A
  global dev-flip ADDS escapes, it doesn't remove one.

**Conclusion:** the sole clean fix is a 1-line platform-source edit to `urls.ts` (add
`process.env.NEXT_PUBLIC_STUDIO_URL ||` before the ternary, mirroring `ACADEMY_URL`), which the v1.10
zero-platform-edit line forbids. → RE-SCOPE TRIGGER. The user decides: (a) carve the Studio link out of the
manager gate with a documented rationale + a presenter-note disclosure (the protocol already has the
presenter-notes mechanism for "don't click this live"); (b) own a tiny upstream platform PR adding the
`NEXT_PUBLIC_STUDIO_URL` read (then the demo injection sets it to `http://localhost:39000`); or (c) pivot.

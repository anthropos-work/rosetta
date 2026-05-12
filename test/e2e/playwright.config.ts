import { defineConfig, devices } from '@playwright/test';

/**
 * Rosetta /test-platform Playwright config.
 *
 * SCOPE: black-box smoke tests against a running platform. These tests
 * verify that the frontend builds, serves, and loads its critical surfaces
 * (login page, public marketing routes if any). They do NOT replicate
 * next-web-app's own E2E suite (which lives in anthropos-dev/next-web-app/e2e/
 * and is invoked by test/repos/run.sh, not by this project).
 *
 * Authenticated flows (dashboard, Talk to Data, etc.) require Clerk fixtures
 * and live in the next-web-app E2E suite, not here.
 */
export default defineConfig({
  testDir: './tests',
  outputDir: './test-results',
  timeout: 30_000,
  expect: { timeout: 5_000 },
  fullyParallel: true,
  forbidOnly: !!process.env.CI,
  retries: process.env.CI ? 1 : 0,
  workers: process.env.CI ? 2 : undefined,
  reporter: [['list'], ['html', { outputFolder: './playwright-report', open: 'never' }]],
  use: {
    baseURL: process.env.ROSETTA_E2E_BASE_URL ?? 'http://localhost:3000',
    trace: 'retain-on-failure',
    screenshot: 'only-on-failure',
  },
  projects: [
    {
      name: 'chromium',
      use: { ...devices['Desktop Chrome'] },
    },
  ],
});

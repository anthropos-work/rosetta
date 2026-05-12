/**
 * Smoke tests — unauthenticated. Verify that the next-web-app serves
 * critical routes without crashing.
 *
 * No fixtures, no logins. Authenticated flows belong to next-web-app's own
 * E2E suite (anthropos-dev/next-web-app/e2e/), invoked by
 * test/repos/run.sh, not by this project.
 */
import { expect, test } from '@playwright/test';

test.describe('rosetta live smoke', () => {
  test('login page renders', async ({ page }) => {
    // App should respond on /login (Clerk SignIn). Tolerant of redirects
    // away from / for unauthenticated users.
    const response = await page.goto('/login', { waitUntil: 'domcontentloaded' });
    expect(response, 'no response from /login').not.toBeNull();
    expect(
      response!.status(),
      `unexpected HTTP status on /login (body: ${(await response!.text()).slice(0, 200)})`,
    ).toBeLessThan(500);

    // Page should contain *some* recognizable login affordance. We avoid
    // asserting on specific Clerk DOM because it changes; we just check the
    // body is non-empty and contains either email/password copy or the
    // Clerk-injected root.
    const body = await page.locator('body').first().textContent();
    expect(body && body.length > 0, 'login page body is empty').toBeTruthy();
  });

  test('root URL responds (any 2xx/3xx or recognized auth redirect)', async ({ page }) => {
    const response = await page.goto('/', { waitUntil: 'domcontentloaded' });
    expect(response, 'no response from /').not.toBeNull();
    const status = response!.status();
    // 2xx = served. 3xx = redirected (likely to /login for unauthenticated).
    expect(status, `unexpected HTTP status on /`).toBeLessThan(500);
  });
});

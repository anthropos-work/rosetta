# Clerk Webhook Local Development Setup

This guide explains how to receive Clerk webhooks during local development so user and organization data syncs to your local database.

## Why Webhooks Matter

Clerk is the source of truth for authentication. When events happen in Clerk (user signup, org creation, membership changes), webhooks notify the backend to sync data locally.

**Without webhooks working**:
- Users created in Clerk won't appear in your local database
- Organization changes won't sync
- Membership and role changes won't be reflected

**Events the backend handles**:
- `user.created`, `user.updated`, `user.deleted`
- `organization.created`, `organization.updated`, `organization.deleted`
- `organizationMembership.created`, `organizationMembership.updated`, `organizationMembership.deleted`
- `organizationInvitation.created`, `organizationInvitation.accepted`, `organizationInvitation.revoked`

---

## The Problem

Clerk runs as a SaaS at `clerk.com`. Your backend runs locally at `localhost:8082`. Clerk cannot reach `localhost` - it needs a public URL.

**Solution**: Use a tunnel service to expose your local backend to the internet temporarily.

---

## Quick Setup: localtunnel (Zero Friction)

localtunnel requires no account signup - just run it via npx.

### Step 1: Start the Tunnel

```bash
# Start tunnel to backend port (8082)
npx localtunnel --port 8082
```

You'll see output like:
```
your url is: https://gentle-flies-think.loca.lt
```

**Copy this URL** - you'll need it for the Clerk dashboard.

### Step 2: Configure Clerk Webhook

1. Go to [Clerk Dashboard](https://dashboard.clerk.com)
2. Select your **development** application (the one with `pk_test_` keys)
3. Navigate to **Webhooks** in the sidebar
4. Click **Add Endpoint**
5. Configure the endpoint:
   - **Endpoint URL**: `https://<your-localtunnel-url>/api/webhook/clerk`
     - Example: `https://gentle-flies-think.loca.lt/api/webhook/clerk`
   - **Subscribe to events** (select all of these):
     - `user.created`
     - `user.updated`
     - `user.deleted`
     - `organization.created`
     - `organization.updated`
     - `organization.deleted`
     - `organizationInvitation.created`
     - `organizationInvitation.accepted`
     - `organizationInvitation.revoked`
     - `organizationMembership.created`
     - `organizationMembership.updated`
     - `organizationMembership.deleted`
6. Click **Create**
7. **Copy the Signing Secret** (starts with `whsec_`)

### Step 3: Update Environment

Add/update the signing secret in `platform/.env`:

```bash
CLERK_WEBHOOK_SECRET=whsec_your_signing_secret_here
```

### Step 4: Restart Backend

```bash
docker compose -p ant-rosetta restart backend
```

### Step 5: Verify

1. Keep the tunnel running in a terminal
2. Open http://localhost:3000 and log in (or create a new account)
3. Check backend logs:
   ```bash
   docker compose -p ant-rosetta logs -f backend 2>&1 | grep -i clerk
   ```
4. Look for: `received event from clerk` with event type `user.created`

---

## localtunnel Drawbacks

- **URL changes on every restart** - You must update the webhook URL in Clerk each time you restart the tunnel
- **Less reliable** - Occasional connection drops, especially on slower networks
- **First request may timeout** - localtunnel shows a warning page on first access; Clerk retries handle this

**For more stable development**, consider the alternatives below.

---

## Alternative: ngrok (More Reliable)

[ngrok](https://ngrok.com) is Clerk's official tunnel partner. More reliable but requires a free account.

### Setup

1. **Create account** at https://ngrok.com (free tier is sufficient)

2. **Install ngrok**:
   ```bash
   # macOS
   brew install ngrok

   # Linux
   curl -s https://ngrok-agent.s3.amazonaws.com/ngrok.asc | sudo tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null
   echo "deb https://ngrok-agent.s3.amazonaws.com buster main" | sudo tee /etc/apt/sources.list.d/ngrok.list
   sudo apt update && sudo apt install ngrok
   ```

3. **Configure auth token** (one-time):
   ```bash
   ngrok config add-authtoken <YOUR_AUTH_TOKEN>
   ```

4. **Start tunnel**:
   ```bash
   ngrok http 8082
   ```

5. Copy the HTTPS URL and configure in Clerk (same as localtunnel steps above)

**ngrok advantages**:
- More stable connections
- Built-in request inspector at http://127.0.0.1:4040
- Replay failed requests for debugging
- Paid tier offers stable URLs

---

## Alternative: Tailscale Funnel (Stable URLs)

If your organization uses Tailscale, Funnel provides stable URLs without changing them on restart.

### Setup

1. **Enable Funnel** for your Tailscale account (requires admin approval)

2. **Start Funnel**:
   ```bash
   # macOS (App Store install)
   /Applications/Tailscale.app/Contents/MacOS/Tailscale funnel http://localhost:8082

   # macOS/Linux (direct install)
   tailscale funnel http://localhost:8082
   ```

3. Use the provided Tailscale URL in Clerk webhook configuration

**Tailscale advantages**:
- Stable URLs that don't change
- No per-session URL updates needed
- Enterprise-grade security

---

## Development Without Webhooks

If you don't need real-time user sync (e.g., just testing frontend UI), you can skip webhook setup:

1. **Use existing test accounts** from the shared development Clerk instance
2. **Seed the database** with test users directly (if seed scripts are available)

This approach is faster but means your local database won't reflect Clerk changes until you set up webhooks.

---

## Troubleshooting

### "Webhook verification failed"

**Cause**: `CLERK_WEBHOOK_SECRET` doesn't match the signing secret in Clerk.

**Fix**:
1. Go to Clerk Dashboard > Webhooks > your endpoint
2. Copy the **Signing Secret**
3. Update `platform/.env`:
   ```bash
   CLERK_WEBHOOK_SECRET=whsec_correct_secret_here
   ```
4. Restart backend: `docker compose -p ant-rosetta restart backend`

### "Connection refused" or "No response"

**Cause**: Tunnel not running or backend not listening.

**Fix**:
1. Verify tunnel is running (check terminal output for URL)
2. Verify backend is running: `docker ps --filter "name=ant-rosetta-backend"`
3. Test backend directly: `curl http://localhost:8082/health`

### localtunnel shows warning page

**Cause**: localtunnel displays a warning page on first access asking to click through.

**Fix**: This is normal. Clerk webhook delivery will retry and succeed. You can also:
1. Visit the localtunnel URL in a browser first
2. Click "Click to Continue"
3. Future requests will work directly

### Webhook URL changed (localtunnel/ngrok free tier)

**Cause**: Tunnel restarted and got a new URL.

**Fix**:
1. Go to Clerk Dashboard > Webhooks
2. Edit your endpoint
3. Update the URL to the new tunnel URL
4. Save

### Events not appearing in logs

**Diagnosis**:
1. Check Clerk Dashboard > Webhooks > your endpoint > "Logs"
2. Look for delivery attempts and their status
3. If showing "Failed", check the error message

---

## Security Notes

- **Never commit** tunnel auth tokens to version control
- **Each developer** needs their own webhook endpoint in Clerk (or update the shared one when you start working)
- **Development only** - production uses a stable domain, not tunnels
- **Webhook secrets** are unique per endpoint - if you create a new endpoint, update your `.env`

---

## Related Documentation

- [External Services](../architecture/external_services.md) - Clerk integration overview
- [Setup Guide](setup_guide.md) - Initial environment setup
- [Run Guide](run_guide.md) - Running the platform locally

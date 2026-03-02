# Messenger Service

## Role & Responsibility

Messenger handles **email notifications** for the Anthropos platform. It integrates with **Brevo** (formerly Sendinblue) as the email delivery provider.

Typical notifications include:
- Simulation assignment notifications
- Completion alerts
- Organization invitations
- System notifications

## Architecture & Code Map

| Property | Value |
|:---------|:------|
| **Technology** | Go |
| **Deployment** | Docker (internal service, no exposed port) |
| **Email Provider** | Brevo (Sendinblue) |
| **Communication** | Called via RPC from other services |

Messenger is an **internal service** — it is not a GraphQL subgraph and has no external-facing API. Other services (primarily Backend/App) call it via Connect-RPC when they need to send emails.

## Local Development

Messenger runs as part of the Docker Compose stack. It does not need to be cloned locally unless you're modifying email templates or notification logic.

```bash
# Runs automatically with docker compose
cd platform
docker compose up -d messenger
```

## Related Documentation
- [Architecture Overview](../architecture/architecture_overview.md)
- [Service Taxonomy](../architecture/service_taxonomy.md)

# Roadrunner Service

## Role & Responsibility

Roadrunner is a **code execution proxy** that enables running user-submitted code in AI Simulations. It forwards code to a **Judge0** sandbox for safe, isolated execution and returns the results.

This powers the **code task type** in simulations, where players write and run code as part of their assessment.

## Architecture & Code Map

| Property | Value |
|:---------|:------|
| **Technology** | Go |
| **Deployment** | Docker (internal service) |
| **Execution Backend** | Judge0 (sandboxed code execution) |
| **Communication** | Called via RPC from Jobsimulation |

Roadrunner acts as a proxy layer between the Jobsimulation service and Judge0, handling:
- Code submission and language selection
- Execution timeout management
- Result formatting and error handling
- Security boundary enforcement

## Local Development

Roadrunner runs as part of the Docker Compose stack. It does not need to be cloned locally unless you're modifying code execution behavior.

```bash
# Runs automatically with docker compose
cd platform
docker compose up -d roadrunner
```

## Related Documentation
- [Architecture Overview](../architecture/architecture_overview.md)
- [Jobsimulation Service](./jobsimulation.md)

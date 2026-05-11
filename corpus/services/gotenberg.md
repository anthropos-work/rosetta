# Gotenberg Service

## Role & Responsibility

Gotenberg is a **third-party stateless conversion service**. It runs LibreOffice headless behind an HTTP API and converts Office documents (DOCX, XLSX, PPTX, etc.) and HTML into PDF.

In the Anthropos platform it exists for one consumer: the backend service uses it to render user-uploaded documents into PDFs for downstream display and storage.

## Architecture

* **Image**: `gotenberg/gotenberg:8` (pinned major version 8)
* **Source**: Upstream project — [gotenberg/gotenberg](https://github.com/gotenberg/gotenberg)
* **Local port**: `3200`
* **Profile**: `graphql` (default) and `backend`
* **Statelessness**: No database, no Redis, no persistence. Spin up / tear down freely.

### Compose command

```yaml
gotenberg:
  image: gotenberg/gotenberg:8
  command:
    - "gotenberg"
    - "--api-port=3200"
    - "--api-timeout=60s"
    - "--libreoffice-restart-after=50"
  ports: ["3200:3200"]
```

The `--libreoffice-restart-after=50` flag restarts the LibreOffice subprocess every 50 conversions to bound memory growth.

## Interface

Gotenberg exposes a multi-route HTTP API. Anthropos uses:

* `POST /forms/libreoffice/convert` — accepts a multipart form upload, returns the PDF bytes

The full API is documented at [gotenberg.dev](https://gotenberg.dev/docs/getting-started/introduction).

## Usage in the Platform

The backend service (`app`) is the only consumer.

* **Code**: [`app/internal/converter/gotenberg.go`](https://github.com/anthropos-work/app/blob/main/internal/converter/gotenberg.go)
* **Endpoint**: `POST {GOTENBERG_URL}/forms/libreoffice/convert`
* **Function**: `ConvertToPDF(ctx, gotenbergURL, document, filename)` returns `[]byte` of the rendered PDF
* **Timeout**: 90 seconds (client-side)
* **Env var**: `GOTENBERG_URL=http://gotenberg:3200` (injected via the backend's compose `environment:`)

## Local Development

Gotenberg starts automatically with the default profile:

```bash
cd platform
make up
```

Verify it's reachable:

```bash
curl -s http://localhost:3200/health
# {"status":"up"...}
```

To exercise the conversion path manually:

```bash
curl --request POST \
  --url http://localhost:3200/forms/libreoffice/convert \
  --form 'files=@example.docx' \
  -o example.pdf
```

## Why a third-party service?

Rendering Office documents requires LibreOffice (a large native binary with its own subprocesses and locale dependencies). Embedding it in the Go backend would bloat the image and complicate the build. Gotenberg packages that complexity behind a clean HTTP boundary and is well-maintained by the open-source community.

## Related Documentation

* [Backend (app)](./backend.md) — the consumer
* [Service Taxonomy](../architecture/service_taxonomy.md)

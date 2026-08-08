# Gotenberg Service

## Role & Responsibility

Gotenberg is a **third-party stateless conversion service**. It runs LibreOffice headless behind an HTTP API and converts Office documents (DOCX, XLSX, PPTX, etc.) and HTML into PDF.

In the Anthropos platform it exists for one consumer — but **not to produce a PDF anybody ever sees.** `backend` uses it as a **text-extraction / OCR intermediate**: an uploaded document the text extractor can't read (or reads and finds no text in) is converted to PDF *in memory*, the text is pulled straight back out of those bytes, and the PDF is discarded. It is **never stored, never served, never displayed**; no PDF here is a platform artifact. Both call sites throw it away in the next statement — `app/internal/web/backend/coursebuilder/extract.go:77-81` converts, then immediately `converter.ConvertFromReader(bytes.NewReader(pdf), "application/pdf")` to get the course-builder source text; `app/internal/worker/tasks/user_import_resume_2d.go:68-74` converts a DOCX résumé **only** to feed the OCR client (`ocrInput = pdfBytes`) after the plain-text path found nothing readable. Measured at `app` `9d00a313` v1.367.0.

## Architecture

* **Image**: `gotenberg/gotenberg:8` (pinned major version 8)
* **Source**: Upstream project — [gotenberg/gotenberg](https://github.com/gotenberg/gotenberg)
* **Local port**: `3200`
* **Profile**: `core` (the default), `backend`, `all` — `profiles: [core, backend, all]` (`docker-compose.yml:183`, re-derived at platform `0c91421`). The default profile is `core`, not `graphql`: `0dab54d` renamed it. Corrected M257x iter-68, re-anchored iter-87 (the anchor stood some eighty lines further down at `0dab54d`; `838d907` deleted three service blocks above it and the file is now 186 lines)
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
* **Function**: `ConvertToPDF(ctx, gotenbergURL, document, filename)` returns `[]byte` of the rendered PDF — **an in-memory intermediate, discarded by both callers** (see Role above)
* **Call sites** (two, `app` @ `9d00a313`): `internal/web/backend/coursebuilder/extract.go:77` — course-builder upload text extraction, for the nine MIME types `docconv` can't read directly (`extract.go:17-27` — `.xls`, `.xlsx`, `.ppt`, `.doc`, the three OpenDocument types, and RTF under both its MIME spellings; **DOCX is not among them**, it goes straight to `docconv`); and `internal/worker/tasks/user_import_resume_2d.go:68` — the résumé-import **OCR fallback**, DOCX only, reached only when the document has no readable text
* **Degrades gracefully**: an empty `GOTENBERG_URL` falls the gotenberg-only formats back to `docconv` (`extract.go:59`, `:43-45`), and a failed conversion on the résumé path only logs and OCRs the original bytes (`user_import_resume_2d.go:69-74`) — neither caller treats a missing PDF as fatal
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

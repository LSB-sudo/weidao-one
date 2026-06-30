# PRD: WebRTC Static Viewer and Niagara Embedding

## Goal

Make the existing RDK WebRTC camera preview easier to embed outside the board-hosted `/viewer` page.

## Requirements

- Keep the existing Go-based WebRTC service flow intact.
- Add a separate static viewer page served by the board.
- Add a standalone external HTML page that can be embedded in Niagara and connect to the board by absolute URL.
- Add CORS support on the board so the external page can call `/offer`, `/health`, and `/devices` from a different origin.
- Keep remote operations standardized on `ssh-skill` for this workspace.

## Deliverables

- Remote static page route: `/viewer-static`
- Remote static file: `web/viewer-static.html`
- Remote CORS support for external HTML embedding
- Local standalone page: `viewer_external.html`
- Updated remote engineering record in `AGENTS.md`

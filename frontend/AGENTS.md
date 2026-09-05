# AGENTS.md

Vue 3 frontend for the "abstracted_self" portfolio site. This file covers only `frontend/`; the repo root `AGENTS.md` has the full-stack picture (backend, Postgres, docker-compose).

## Stack

- Vue 3 + Vite + vue-router, **plain JavaScript** `.vue` SFCs (`<script setup>`) — no TypeScript.
- Tailwind CSS v4 via the `@tailwindcss/vite` plugin (declared in `vite.config.js`, not a PostCSS/tailwind.config). `src/style.css` does `@import "tailwindcss"` + `@plugin "@tailwindcss/aspect-ratio"`.
- Icons: `@heroicons/vue` (24/outline). `SidebarLink` takes a string `icon` prop mapped to a component in a local `icons` object, not a direct component import.
- `typed.js` is loaded from a CDN `<script>` in `index.html`; `src/pages/partials/Hero.vue` uses the global `Typed`.

## Commands

- Dev server: `npm install && npm run dev`
- Build: `npm run build` (default) or `npm run build -- --mode docker` (loads `.env.docker`)
- No tests, no linter, no CI.

## Gotchas

### Local dev can't reach the backend
`vite.config.js` has no proxy and there is no `.env` / `.env.development` — only `.env.docker`. So `import.meta.env.VITE_API_URL` and every `VITE_*_ENDPOINT` are `undefined` under `npm run dev`, and API calls fail. Do full-stack work via `docker compose`, where nginx (`nginx.conf`) serves the SPA and proxies `/api/*` to the backend.

### Env-driven API paths, no shared API client
There is no axios/fetch wrapper. Components call `fetch` directly, building URLs as `import.meta.env.VITE_API_URL + import.meta.env.VITE_<X>_ENDPOINT` (all endpoint vars are defined in `.env.docker`). Authenticated requests read the token from `localStorage.getItem("token")` and send `Authorization: Bearer <token>`.

### Ignore the stale `.opencode/` docs
`.opencode/agent/*` and `.opencode/context/*` describe an **unrelated project** ("Vector Vault", a RAG chat app) and prescribe plain CSS / no-Tailwind / TypeScript / composables-only. The real code is Tailwind v4 + plain JS + vue-router. Follow `src/` and `package.json`, not those docs.

### Auth & routing
- Token is stored in `localStorage.token`; login page is `/unlock` (`Login.vue`).
- Protected routes live under `/admin/*` and set `meta.requiresAuth`. The guard in `src/router/index.js` renews via `VITE_RENEW_TOKEN_ENDPOINT` when <2 min from expiry, then validates with `VITE_CHECK_TOKEN_ENDPOINT` before allowing navigation.
- Public routes (`/`) also fire a visitor-tracking POST to `VITE_PROFILE_VISIT_ENDPOINT` using a `visitor_uuid` js-cookie.

## Layout

- `src/main.js` — entry; mounts `App.vue` + router.
- `src/router/index.js` — all routes + auth guard.
- `src/pages/` — route pages; `src/pages/partials/` — sections of the public profile page (`Profile.vue`).
- `src/layouts/` — `ProfileLayout.vue` (public) and `AdminLayout.vue` (admin shell).
- `src/components/` — shared/reusable components (forms, cards, uploads).
- `src/style.css` — Tailwind import + global scrollbar styling only.

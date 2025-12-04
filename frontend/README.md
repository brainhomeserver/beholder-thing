# Beholder Frontend

This folder houses the React single-page application and Worker backend that power the Beholder UI. It follows [Cloudflare's React + Vite guide](https://developers.cloudflare.com/workers/framework-guides/web-apps/react/) and uses TypeScript with the SWC-powered React plugin for faster HMR and builds.

## Prerequisites

- Node.js 20+ and npm 10+
- A Cloudflare account configured for `wrangler` if you plan to deploy

## Available scripts

Run the following commands from this directory:

- `npm run dev` – start the Cloudflare-aware Vite dev server at <http://localhost:5173>
- `npm run lint` – run ESLint across the worker and React source
- `npm run build` – type-check and build both the Worker and the client bundles
- `npm run check` – full CI-style check (`tsc`, `vite build`, and a dry-run `wrangler deploy`)
- `npm run deploy` – deploy the Worker + assets to Cloudflare
- `npm run cf-typegen` – regenerate Worker binding types after updating `wrangler.json`

## Project layout

- `src/react-app` – the React SPA, built with Vite + SWC
- `src/worker` – the Hono Worker that exposes `/api/`
- `public/` – static assets copied verbatim into the final build
- `wrangler.json` – Worker configuration (assets routing, compatibility, etc.)

Both the Worker and the client are written in TypeScript with strict settings enabled to keep the surface area well-typed.

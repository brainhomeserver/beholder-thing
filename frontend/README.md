# React + Cloudflare Workers

This frontend follows Cloudflare's [React + Workers guide](https://developers.cloudflare.com/workers/framework-guides/web-apps/react/).  
It is a Vite + React app that uses TypeScript with SWC plus a Workers backend powered by the Cloudflare Vite plugin.

## Useful scripts

| Command | Description |
| --- | --- |
| `npm run dev` | Start the Vite dev server with the Workers runtime in watch mode. |
| `npm run build` | Type-check and produce a production build. |
| `npm run preview` | Build then run `vite preview` to serve the static assets locally. |
| `npm run deploy` | Build and deploy the Worker via `wrangler deploy`. |
| `npm run cf-typegen` | Regenerate `worker-configuration.d.ts` after adding bindings. |

## Cloudflare Worker

The Worker entry point lives in `worker/index.ts`. Requests to `/api/*` are handled by the Worker so you can
call Cloudflare bindings or other APIs without exposing credentials to the browser. Static assets are served from
Vite's build output and fall back to the Worker as a single-page application.

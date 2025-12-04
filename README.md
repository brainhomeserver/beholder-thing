# beholder-thing

This should be all coded by AI.

## Frontend

The `frontend` directory contains a Cloudflare Workers application that follows the [React + Vite + Workers](https://developers.cloudflare.com/workers/framework-guides/web-apps/react/) tutorial. It is written in TypeScript and uses the SWC-powered React plugin for Vite.

### Local development

```bash
cd frontend
npm install
npm run dev
```

### Linting and production build

```bash
cd frontend
npm run lint
npm run build
```

The build output lives in `frontend/dist/client` for the SPA assets and `frontend/dist/frontend` for the Worker bundle. Deploy with `npm run deploy` once you have authenticated Wrangler with your Cloudflare account.

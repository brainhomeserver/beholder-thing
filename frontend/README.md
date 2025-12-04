# Frontend

React + Vite (SWC) + Hono + Cloudflare Workers boilerplate following the Cloudflare React guide. The app ships with a simple `/api/` endpoint in a Worker (via Hono) and a React SPA that fetches from it. Tooling includes TypeScript, SWC-based React transforms, ESLint, and Wrangler for local dev and deploys.

## Development

Install dependencies:

```bash
npm install
```

Start the development server with:

```bash
npm run dev
```

Your application will be available at [http://localhost:5173](http://localhost:5173).

## Production

Build your project for production:

```bash
npm run build
```

Preview your build locally:

```bash
npm run preview
```

Deploy your project to Cloudflare Workers:

```bash
npm run build && npm run deploy
```

Monitor your workers:

```bash
npx wrangler tail
```

## Additional Resources

- [Cloudflare Workers Documentation](https://developers.cloudflare.com/workers/)
- [Vite Documentation](https://vitejs.dev/guide/)
- [React Documentation](https://reactjs.org/)
- [Hono Documentation](https://hono.dev/)

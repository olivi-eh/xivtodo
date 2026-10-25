import { readFile, writeFile } from "node:fs/promises";
import { resolve } from "node:path";
import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

const siteURL = "https://xivtodo.com";

// Routes served as copies of index.html, so loading them directly returns a 200 instead of going through 404.html.
// Keep in sync with src/router/index.js; dynamic routes (e.g. /character/:id) still fall back to 404.html.
const staticRoutes = [
  { path: "/encounters", sitemap: true },
  { path: "/questlines", sitemap: true },
  { path: "/challenges", sitemap: true },
  { path: "/checklist", sitemap: true },
  { path: "/help", sitemap: true },
  { path: "/privacy", sitemap: true },
  { path: "/auth" },
  { path: "/profile" },
  { path: "/settings" },
  { path: "/signout" },
  { path: "/sandbox" },
];

const replaceOrFail = (html, search, replacement) => {
  if (!html.includes(search)) {
    throw new Error(`static-pages: ${search} not found in index.html`);
  }
  return html.replaceAll(search, replacement);
};

// Emits <route>.html pages, which GitHub Pages serves at /<route> without a trailing slash redirect, and sitemap.xml
const staticPages = () => {
  let outDir;
  return {
    name: "static-pages",
    apply: "build",
    configResolved(config) {
      outDir = resolve(config.root, config.build.outDir);
    },
    async closeBundle() {
      const indexHTML = await readFile(resolve(outDir, "index.html"), "utf-8");

      for (const route of staticRoutes) {
        let html = replaceOrFail(indexHTML, `href="${siteURL}/"`, `href="${siteURL}${route.path}"`);
        html = replaceOrFail(html, `content="${siteURL}"`, `content="${siteURL}${route.path}"`);
        await writeFile(resolve(outDir, `${route.path.slice(1)}.html`), html);
      }

      const lastmod = new Date().toISOString().slice(0, 10);
      const urls = ["/", ...staticRoutes.filter((route) => route.sitemap).map((route) => route.path)]
        .map((path) => `  <url>\n    <loc>${siteURL}${path}</loc>\n    <lastmod>${lastmod}</lastmod>\n  </url>`)
        .join("\n");
      const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${urls}
</urlset>
`;
      await writeFile(resolve(outDir, "sitemap.xml"), sitemap);
    },
  };
};

export default defineConfig({
  plugins: [vue(), staticPages()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    // Matches the local Discord OAuth redirect URI
    port: 8080,
    strictPort: true,
  },
});

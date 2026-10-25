import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  plugins: [vue()],
  // Keep the VUE_APP_ prefix so existing .env files and the deploy workflow keep working
  envPrefix: "VUE_APP_",
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

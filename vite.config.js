import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  build: {
    lib: {
      entry: "client/application.ts",
      name: "application",
    },
    manifest: true,
    rollupOptions: {
      output: {
        dir: "dist",
        entryFileNames: "application.js",
        assetFileNames: "application.css",
        chunkFileNames: "chunk.js",
        manualChunks: undefined,
      },
    },
  },
  feature: {},
  plugins: [],
  css: {
    postcss: "./postcss.config.js",
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "client"),
      "@Components": path.resolve(__dirname, "client/components"),
      "@Helpers": path.resolve(__dirname, "client/helpers"),
    },
    extensions: [".js", ".jsx", ".ts", ".tsx", ".css"],
  },
  optimizeDeps: {
    include: [],
    exclude: [],
  },
});

import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  build: {
    outDir: "dist",
    manifest: true,
  },
  feature: {},
  plugins: [],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "client"),
      "@Components": path.resolve(__dirname, "client/components"),
    },
    extensions: [".js", ".jsx", ".ts", ".tsx", ".css"],
  },
  optimizeDeps: {
    include: [],
    exclude: [],
  },
});

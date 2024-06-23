import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  ssr: {
    optimizeDeps: {
      include: ["lodash"],
    },
    noExternal: ["lodash"],
  },
});

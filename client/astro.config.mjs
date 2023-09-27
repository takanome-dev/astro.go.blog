import { defineConfig } from "astro/config";
import nodejs from "@astrojs/node";
import tailwind from "@astrojs/tailwind";
import react from "@astrojs/react";

// https://astro.build/config
export default defineConfig({
  integrations: [tailwind(), react()],
  output: "server",
  adapter: nodejs({
    mode: "standalone",
  }),
  vite: {
    ssr: {
      noExternal: ["class-variance-authority", "tailwind-merge"],
    },
  },
});

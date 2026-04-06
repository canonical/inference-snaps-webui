#!/usr/bin/env node
// Vendors frontend assets into static/vendor/ for offline use.
// Run with: npm run vendor

import { copyFileSync, mkdirSync, writeFileSync } from "fs";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import * as sass from "sass";

const __dirname = dirname(fileURLToPath(import.meta.url));
const root = resolve(__dirname, "..");
const dest = resolve(root, "static/vendor");

mkdirSync(dest, { recursive: true });

// Vue: copy pre-built global prod bundle from node_modules
copyFileSync(
  resolve(root, "node_modules/vue/dist/vue.global.prod.js"),
  resolve(dest, "vue.global.prod.js"),
);
console.log("✓ vue.global.prod.js");

// Vanilla Framework: compile from SCSS
// silenceDeprecations suppresses known warnings from vanilla-framework's SCSS
// (it still uses legacy @import, global built-ins, and old color functions)
const result = sass.compile(
  resolve(root, "node_modules/vanilla-framework/_index.scss"),
  {
    style: "compressed",
    quietDeps: true,
    silenceDeprecations: ["import", "global-builtin", "color-functions"],
  },
);
writeFileSync(resolve(dest, "vanilla-framework.min.css"), result.css);
console.log("✓ vanilla-framework.min.css");

// Ubuntu CoF logo: download from Canonical CDN (not on npm)
const res = await fetch("https://assets.ubuntu.com/v1/82818827-CoF_white.svg");
if (!res.ok) throw new Error(`Failed to fetch logo: ${res.status}`);
writeFileSync(resolve(dest, "CoF_white.svg"), Buffer.from(await res.arrayBuffer()));
console.log("✓ CoF_white.svg");

console.log(`\nVendored assets written to static/vendor/`);

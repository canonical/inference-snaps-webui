# Inference Snaps UI

This project is based on a template of Vue 3 in Vite.

## Project specific configuration

The web UI loads configuration from the `/config` endpoint.
The production web server should host the configuration at this path.
For development a static file is included in `mock/config.json`.
In the development server this file will be served at `/config`, and can therefore be used to test different configurations.

## Recommended IDE Setup

[VS Code](https://code.visualstudio.com/) + [Vue (Official)](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Recommended Browser Setup

- Chromium-based browsers (Chrome, Edge, Brave, etc.):
  - [Vue.js devtools](https://chromewebstore.google.com/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)
  - [Turn on Custom Object Formatter in Chrome DevTools](http://bit.ly/object-formatters)
- Firefox:
  - [Vue.js devtools](https://addons.mozilla.org/en-US/firefox/addon/vue-js-devtools/)
  - [Turn on Custom Object Formatter in Firefox DevTools](https://fxdx.dev/firefox-devtools-custom-object-formatters/)

## Type Support for `.vue` Imports in TS

This project type-checks with the native TypeScript 7 compiler (`typescript@7.1.0-beta`).
TypeScript itself cannot read `.vue` Single-File Components, so the project registers
[`@vue/content-mapper`](https://github.com/vuejs/language-tools) (a Vue Language Tools v4
alpha) as a *content mapper* in `tsconfig.app.json`. TypeScript 7.1 runs the mapper as
external code (`tsc -b --runExternalCode`, see the `type-check` script), which turns each
`.vue` file into type-checkable TypeScript and maps diagnostics back to the original file.

This replaces the previous `vue-tsc` setup, which relied on the TypeScript 6 JavaScript
API that TypeScript 7 no longer ships.

### TypeScript 7 migration status

The build, type-check, unit tests and `npm ci` all work on TypeScript 7, but the
migration is mid-flight across the ecosystem. Until it settles:

- `typescript` is `7.1.0-beta` and `@vue/content-mapper` is a `4.0.0-alpha` release;
  pin exact versions (as done in `package.json`) to avoid accidental upgrades.
- ESLint's TypeScript integration (`typescript-eslint`) cannot run on TypeScript 7 yet.
  `npm run lint` therefore only runs `oxlint`; the ESLint config is kept intact and can
  be re-wired once [typescript-eslint supports TypeScript 7](https://typescript-eslint.io/users/dependency-versions/).
  For the same reason, `.npmrc` sets `legacy-peer-deps=true` (typescript-eslint's
  `typescript` peer range still excludes 7.x).
- Editor support is in transition: the Vue language server still expects the classic
  TypeScript API. See [vuejs/language-tools#6170](https://github.com/vuejs/language-tools/pull/6170)
  for the status of Vue tooling on TypeScript 7.

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Lint with [OxLint](https://oxc.rs/docs/guide/usage/linter) (ESLint temporarily disabled)

```sh
npm run lint
```

See "TypeScript 7 migration status" above for why the ESLint pass is currently
unwired from this command.

// Vue's built-in JSX attribute types do not include the standard `data-*`
// custom element attributes. Vue Language Tools v4 checks native element
// attributes strictly, so declare them to keep templates type-safe.
//
// This file must be a module (hence the `export {}`) for the `declare module`
// below to act as a module augmentation rather than an ambient module
// declaration that would shadow the real `@vue/runtime-dom` package.
export {}

declare module '@vue/runtime-dom' {
  interface HTMLAttributes {
    [key: `data-${string}`]: unknown
  }
}

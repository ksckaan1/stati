

export const index = 1;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/error.svelte.js')).default;
export const imports = ["_app/immutable/nodes/1.fda56a35.js","_app/immutable/chunks/scheduler.643286d3.js","_app/immutable/chunks/index.c2323441.js","_app/immutable/chunks/singletons.f28f9ba0.js"];
export const stylesheets = [];
export const fonts = [];

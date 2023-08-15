

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.b7cab716.js","_app/immutable/chunks/scheduler.643286d3.js","_app/immutable/chunks/index.c2323441.js"];
export const stylesheets = ["_app/immutable/assets/2.49338e76.css"];
export const fonts = [];

export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {"start":"_app/immutable/entry/start.c20d5c95.js","app":"_app/immutable/entry/app.64899922.js","imports":["_app/immutable/entry/start.c20d5c95.js","_app/immutable/chunks/scheduler.643286d3.js","_app/immutable/chunks/singletons.f28f9ba0.js","_app/immutable/entry/app.64899922.js","_app/immutable/chunks/scheduler.643286d3.js","_app/immutable/chunks/index.c2323441.js"],"stylesheets":[],"fonts":[]},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		}
	}
}
})();

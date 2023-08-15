import adapter from '@sveltejs/adapter-static';
import {vitePreprocess} from "@sveltejs/kit/vite";

const config =  {
	kit: {
		adapter: adapter({
			fallback: 'index.html' // may differ from host to host
		})
	},
	preprocess: vitePreprocess()
};

export default config;

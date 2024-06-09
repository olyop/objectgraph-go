/* eslint-disable unicorn/better-regex */
import { readFile } from "node:fs/promises";

import reactSwc from "@vitejs/plugin-react-swc";
import { Plugin, UserConfig, defineConfig, loadEnv } from "vite";
import checker from "vite-plugin-checker";
import tsconfigPaths from "vite-tsconfig-paths";

import { graphQL } from "./vite/vite-plugin-graphql";

type Mode = "development" | "production";

const checkerOptions: Parameters<typeof checker>[0] = {
	typescript: true,
	eslint: {
		lintCommand: "eslint",
		dev: {
			overrideConfig: {
				overrideConfig: {
					parserOptions: {
						project: "./tsconfig.eslint.json",
					},
				},
			},
		},
	},
};

export default defineConfig(async options => {
	const mode = options.mode as Mode;

	const environmentVariables = loadEnv(mode, process.cwd(), "");

	process.env = { ...process.env, ...environmentVariables };

	const config: UserConfig = {
		plugins: [tsconfigPaths(), reactSwc(), checker(checkerOptions), graphQL()],
		define: {
			__DEV__: JSON.stringify(mode === "development"),
			"globalThis.__DEV__": JSON.stringify(mode === "development"),
		},
		server: {
			https:
				mode === "development"
					? {
							cert: await readFile(process.env.TLS_CERT_PATH),
							key: await readFile(process.env.TLS_KEY_PATH),
						}
					: undefined,
		},
	};

	return config;
});

declare global {
	// eslint-disable-next-line @typescript-eslint/no-namespace
	namespace NodeJS {
		// eslint-disable-next-line unicorn/prevent-abbreviations
		interface ProcessEnv {
			TLS_CERT_PATH: string;
			TLS_KEY_PATH: string;
		}
	}
}

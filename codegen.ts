import { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
	schema: "server/graphql/schema/*.graphqls",
	documents: "client/src/**/*.graphql",
	ignoreNoDocuments: true,
	hooks: {
		afterAllFileWrite: "npx prettier --write client/src/types/graphql.ts",
	},
	config: {
		avoidOptionals: true,
		immutableTypes: true,
		useTypeImports: true,
		mergeFragmentTypes: true,
		nonOptionalTypename: true,
		useImplementingTypes: true,
		flattenGeneratedTypes: true,
		defaultScalarType: "unknown",
		inlineFragmentTypes: "combine",
		namingConvention: {
			enumValues: "change-case#upperCase",
		},
		scalars: {
			UUID: "string",
			Price: "number",
			Timestamp: "string",
		},
	},
	generates: {
		"client/src/types/graphql.ts": {
			plugins: ["typescript", "typescript-operations"],
		},
	},
};

export default config;

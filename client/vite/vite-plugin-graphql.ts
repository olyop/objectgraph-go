import graphqlTag from "graphql-tag";
import { Plugin } from "vite";

const transform: Plugin["transform"] = (source: string, id: string) => {
	if (id.endsWith(".graphql")) {
		return {
			map: null,
			code: `export default JSON.parse('${JSON.stringify(graphqlTag(source))}');`,
		};
	} else {
		return null;
	}
};

export const graphQL = (): Plugin => ({
	name: "vite-plugin-gql",
	transform,
});

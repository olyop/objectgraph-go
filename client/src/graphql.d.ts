declare module "*.graphql" {
	import type { DocumentNode } from "@apollo/client";

	const value: DocumentNode;

	export default value;
}

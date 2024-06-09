import { ApolloClient, ApolloProvider as ApolloProviderInternal, HttpLink, InMemoryCache } from "@apollo/client";
import { createElement } from "react";

const apolloCache = new InMemoryCache();

const apolloLink = new HttpLink({
	uri: "https://localhost:8080/graphql",
});

const apolloClient = new ApolloClient({
	cache: apolloCache,
	link: apolloLink,
});

export function ApolloProvider({ children }: { children: React.ReactNode }) {
	return <ApolloProviderInternal client={apolloClient}>{children}</ApolloProviderInternal>;
}

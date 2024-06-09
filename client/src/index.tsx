import { createElement } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import "tailwindcss/tailwind.css";

import { ApolloProvider } from "@/clients/apollo";
import { Routes } from "@/routes";

import "./index.css";

const container = document.getElementById("root");

if (!container) {
	throw new Error("No root element found");
}

const root = createRoot(container);

root.render(
	<BrowserRouter>
		<ApolloProvider>
			<Routes />
		</ApolloProvider>
	</BrowserRouter>,
);

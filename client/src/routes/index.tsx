import { createElement } from "react";
import { Route, Routes as RoutesInternal } from "react-router-dom";

import { Home } from "./home";

export function Routes() {
	return (
		<RoutesInternal>
			<Route path="" element={<Home />} />
		</RoutesInternal>
	);
}

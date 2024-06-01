import { FC, createElement } from "react";
import { BrowserRouter as Router } from "react-router-dom";

import { Content } from "./content";

export const Main: FC = () => (
	<Router>
		<Content />
	</Router>
);

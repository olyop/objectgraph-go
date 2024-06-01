import { FC, createElement } from "react";
import { Route, Routes } from "react-router-dom";

export const Content: FC = () => (
	<div className="h-screen w-screen bg-black font-sans text-white">
		<Routes>
			<Route path="" element={<p>Hello World</p>} />
		</Routes>
	</div>
);

/** @type {import('tailwindcss').Config} */

module.exports = {
	content: ["./src/**/*.{js,ts,jsx,tsx}"],
	plugins: [],
	theme: {
		fontFamily: {
			mono: ["monospace"],
			sans: ["CircularSP-Book", "sans-serif"],
		},
		extend: {
			colors: {
				primary: "#1db954",
				"primary-light": "#60ce87",
				"primary-dark": "#14813a",
				"spotify-base": "#121212",
				"spotify-hover": "#242424",
			},
			height: {
				"content-height": "calc(100vh - 7.5rem)",
				"scroll-height": "calc(100vh - 10rem)",
				"bar-height": "calc(8rem)",
			},
			width: {
				"sidebar-open": "24rem",
				"sidebar-closed": "11rem",
				"content-width": "calc(100vw - 24rem - 24rem)",
			},
			padding: {
				"sidebar-open": "24rem",
				"sidebar-closed": "11rem",
			},
			top: {
				"header-bar-height": "6rem",
			},
			borderColor: {
				"main": "#a7a7a7",
			},
		},
	},
};

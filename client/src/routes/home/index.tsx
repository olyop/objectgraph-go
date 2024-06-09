import { useQuery } from "@apollo/client";
import { createElement } from "react";

import { Product } from "@/components/product";
import { GetProductsQuery } from "@/types/graphql";

import GET_PRODUCTS from "./get-products.graphql";

export function Home() {
	const { data, loading } = useQuery<GetProductsQuery>(GET_PRODUCTS);

	if (loading) {
		return <p>Loading...</p>;
	}

	return (
		<div className="grid grid-cols-3 gap-2 p-2">
			{data?.getProducts.map(product => <Product key={product.productID} product={product} />)}
		</div>
	);
}

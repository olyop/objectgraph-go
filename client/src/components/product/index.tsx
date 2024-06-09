import { createElement } from "react";

import { Product as ProductType } from "@/types/graphql";

export function Product({ product }: ProductProps) {
	return (
		<div className="flex flex-col items-center gap-2 rounded-2xl border p-2">
			<p className="text-center text-xl">{product.name}</p>
			<p className="text-center text-xl">{product.brand.name}</p>
			<p className="text-center text-xl">{product.categories[0]?.classification.name}</p>
			<p className="text-center text-xl">{product.price}</p>
		</div>
	);
}

export interface ProductProps {
	product: ProductType;
}

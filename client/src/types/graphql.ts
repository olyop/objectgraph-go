export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
	ID: { input: string; output: string };
	String: { input: string; output: string };
	Boolean: { input: boolean; output: boolean };
	Int: { input: number; output: number };
	Float: { input: number; output: number };
	Price: { input: number; output: number };
	Timestamp: { input: string; output: string };
	UUID: { input: string; output: string };
};

export type Brand = {
	readonly __typename: "Brand";
	readonly brandID: Scalars["UUID"]["output"];
	readonly createdAt: Scalars["Timestamp"]["output"];
	readonly name: Scalars["String"]["output"];
	readonly updatedAt: Maybe<Scalars["Timestamp"]["output"]>;
};

export type Category = {
	readonly __typename: "Category";
	readonly categoryID: Scalars["UUID"]["output"];
	readonly classification: Classification;
	readonly createdAt: Scalars["Timestamp"]["output"];
	readonly name: Scalars["String"]["output"];
	readonly updatedAt: Maybe<Scalars["Timestamp"]["output"]>;
};

export type Classification = {
	readonly __typename: "Classification";
	readonly classificationID: Scalars["UUID"]["output"];
	readonly createdAt: Scalars["Timestamp"]["output"];
	readonly name: Scalars["String"]["output"];
	readonly updatedAt: Maybe<Scalars["Timestamp"]["output"]>;
};

export type Mutation = {
	readonly __typename: "Mutation";
	readonly updateProductByID: Maybe<Product>;
};

export type MutationUpdateProductByIdArgs = {
	input: UpdateProductInput;
};

export type Product = {
	readonly __typename: "Product";
	readonly abv: Maybe<Scalars["Float"]["output"]>;
	readonly brand: Brand;
	readonly categories: ReadonlyArray<Category>;
	readonly createdAt: Scalars["Timestamp"]["output"];
	readonly name: Scalars["String"]["output"];
	readonly price: Scalars["Price"]["output"];
	readonly productID: Scalars["UUID"]["output"];
	readonly promotionDiscount: Maybe<Scalars["Price"]["output"]>;
	readonly promotionDiscountMultiple: Maybe<Scalars["Int"]["output"]>;
	readonly updatedAt: Maybe<Scalars["Timestamp"]["output"]>;
	readonly volume: Maybe<Scalars["Int"]["output"]>;
};

export type Query = {
	readonly __typename: "Query";
	readonly getProducts: ReadonlyArray<Product>;
};

export type UpdateProductInput = {
	readonly name: Scalars["String"]["input"];
	readonly productID: Scalars["UUID"]["input"];
};

export type GetProductsQueryVariables = Exact<{ [key: string]: never }>;

export type GetProductsQuery = {
	readonly getProducts: ReadonlyArray<
		{
			readonly productID: string;
			readonly name: string;
			readonly price: number;
			readonly promotionDiscount: number | null;
			readonly promotionDiscountMultiple: number | null;
			readonly volume: number | null;
			readonly abv: number | null;
			readonly updatedAt: string | null;
			readonly createdAt: string;
			readonly brand: {
				readonly brandID: string;
				readonly name: string;
				readonly updatedAt: string | null;
				readonly createdAt: string;
			} & { readonly __typename: "Brand" };
			readonly categories: ReadonlyArray<
				{
					readonly categoryID: string;
					readonly name: string;
					readonly updatedAt: string | null;
					readonly createdAt: string;
					readonly classification: {
						readonly classificationID: string;
						readonly name: string;
						readonly updatedAt: string | null;
						readonly createdAt: string;
					} & { readonly __typename: "Classification" };
				} & { readonly __typename: "Category" }
			>;
		} & { readonly __typename: "Product" }
	>;
} & { readonly __typename: "Query" };

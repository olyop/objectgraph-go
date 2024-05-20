CREATE
OR REPLACE function get_now () returns bigint language sql stable AS $$
		SELECT cast(extract(epoch FROM now()) AS BIGINT) * 1000;
	$$;

CREATE TABLE brands (
	brand_id UUID DEFAULT gen_random_uuid (),
	name text NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT brands_pk PRIMARY KEY (brand_id),
	CONSTRAINT brands_name_unique UNIQUE (name),
	CONSTRAINT brands_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE classifications (
	classification_id UUID DEFAULT gen_random_uuid (),
	name text NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT classifications_pk PRIMARY KEY (classification_id),
	CONSTRAINT classifications_name_unique UNIQUE (name),
	CONSTRAINT classifications_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE categories (
	category_id UUID DEFAULT gen_random_uuid (),
	name text NOT NULL,
	classification_id UUID NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT categories_pk PRIMARY KEY (category_id),
	CONSTRAINT categories_name_unique UNIQUE (name),
	CONSTRAINT categories_classification_id_fk FOREIGN key (classification_id) REFERENCES classifications (classification_id),
	CONSTRAINT categories_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE products (
	product_id UUID NOT NULL DEFAULT gen_random_uuid (),
	name varchar(255) NOT NULL,
	brand_id UUID NOT NULL,
	price_id UUID NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT products_pk PRIMARY KEY (product_id),
	CONSTRAINT products_name_unique UNIQUE (name),
	CONSTRAINT products_brand_id_fk FOREIGN key (brand_id) REFERENCES brands (brand_id),
	CONSTRAINT products_price_fk FOREIGN key (price_id) REFERENCES products_prices (price_id),
	CONSTRAINT products_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE products_prices (
	price_id UUID NOT NULL DEFAULT gen_random_uuid (),
	product_id UUID NOT NULL,
	price numeric NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT products_prices_pk PRIMARY KEY (price_id),
	CONSTRAINT products_prices_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_prices_price_ck CHECK (price >= 0),
	CONSTRAINT products_prices_product_id_price_unique UNIQUE (product_id, price),
	CONSTRAINT products_prices_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE products_categories (
	product_id UUID NOT NULL,
	category_id UUID NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT products_categories_pk PRIMARY KEY (product_id, category_id),
	CONSTRAINT products_categories_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_categories_category_id_fk FOREIGN key (category_id) REFERENCES categories (category_id),
	CONSTRAINT products_categories_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE products_abv (
	product_id UUID NOT NULL,
	abv numeric NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT products_abv_pk PRIMARY KEY (product_id, abv),
	CONSTRAINT products_abv_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_abv_abv_ck CHECK (
		abv >= 0
		AND abv <= 100
	),
	CONSTRAINT products_abv_created_at_ck CHECK (created_at >= 0)
);

CREATE TABLE products_volumes (
	product_id UUID NOT NULL,
	volume numeric NOT NULL,
	created_at bigint NOT NULL DEFAULT get_now (),
	CONSTRAINT products_volumes_pk PRIMARY KEY (product_id),
	CONSTRAINT products_volumes_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_volumes_volume_ck CHECK (volume >= 0),
	CONSTRAINT products_volumes_created_at_ck CHECK (created_at >= 0)
);

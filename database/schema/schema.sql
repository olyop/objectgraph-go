CREATE OR
REPLACE function get_now () returns BIGINT language sql stable AS $$
		SELECT cast(extract(epoch FROM now()) AS BIGINT) * 1000;
	$$;

CREATE TABLE IF NOT EXISTS classifications (
	classification_id UUID DEFAULT gen_random_uuid (),
	classification_name VARCHAR(255) NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT classifications_pk PRIMARY KEY (classification_id),
	CONSTRAINT classifications_classification_name_uq UNIQUE (classification_name),
	CONSTRAINT classifications_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT classifications_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS categories (
	category_id UUID DEFAULT gen_random_uuid (),
	category_name VARCHAR(255) NOT NULL,
	classification_id UUID NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT categories_pk PRIMARY KEY (category_id),
	CONSTRAINT categories_classification_id_fk FOREIGN key (classification_id) REFERENCES classifications (classification_id),
	CONSTRAINT categories_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT categories_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS brands (
	brand_id UUID DEFAULT gen_random_uuid (),
	brand_name VARCHAR(255) NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT brands_pk PRIMARY KEY (brand_id),
	CONSTRAINT brands_brand_name_uq UNIQUE (brand_name),
	CONSTRAINT brands_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT brands_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS prices (
	price_id UUID NOT NULL DEFAULT gen_random_uuid (),
	price_value BIGINT NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT prices_pk PRIMARY KEY (price_id),
	CONSTRAINT prices_price_value_ck CHECK (price_value > 0),
	CONSTRAINT prices_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT prices_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS promotions (
	promotion_id UUID NOT NULL DEFAULT gen_random_uuid (),
	promotion_discount BIGINT,
	promotion_discount_multiple SMALLINT,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT promotions_pk PRIMARY KEY (promotion_id),
	CONSTRAINT promotions_promotion_discount_ck CHECK (promotion_discount > 0),
	CONSTRAINT promotions_promotion_discount_multiple_ck CHECK (promotion_discount_multiple > 0),
	CONSTRAINT promotions_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT promotions_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS products (
	product_id UUID NOT NULL DEFAULT gen_random_uuid (),
	product_name VARCHAR(255) NOT NULL,
	brand_id UUID NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT products_pk PRIMARY KEY (product_id),
	CONSTRAINT products_product_name_uq UNIQUE (product_name),
	CONSTRAINT products_brand_id_fk FOREIGN key (brand_id) REFERENCES brands (brand_id),
	CONSTRAINT products_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT products_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS products_abvs (
	product_id UUID NOT NULL,
	abv SMALLINT NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT products_abvs_pk PRIMARY KEY (product_id, abv),
	CONSTRAINT products_abvs_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_abvs_abv_ck CHECK (
		abv > 0 AND
		abv <= 100
	),
	CONSTRAINT products_abvs_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT products_abvs_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS products_volumes (
	product_id UUID NOT NULL,
	volume SMALLINT NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT products_volumes_pk PRIMARY KEY (product_id, volume),
	CONSTRAINT products_volumes_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_volumes_volume_ck CHECK (volume > 0),
	CONSTRAINT products_volumes_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT products_volumes_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS products_categories (
	product_id UUID NOT NULL,
	category_id UUID NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT products_categories_pk PRIMARY KEY (product_id, category_id),
	CONSTRAINT products_categories_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_categories_category_id_fk FOREIGN key (category_id) REFERENCES categories (category_id),
	CONSTRAINT products_categories_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT products_categories_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS products_prices (
	product_id UUID NOT NULL,
	price_id UUID NOT NULL,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT products_prices_pk PRIMARY KEY (product_id),
	CONSTRAINT products_prices_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_prices_price_id_fk FOREIGN key (price_id) REFERENCES prices (price_id),
	CONSTRAINT products_prices_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT products_prices_created_at_ck CHECK (created_at > 0)
);

CREATE TABLE IF NOT EXISTS products_promotions (
	product_id UUID NOT NULL,
	promotion_id UUID NOT NULL,
	active_from BIGINT NOT NULL DEFAULT get_now (),
	active_to BIGINT,
	updated_at BIGINT,
	created_at BIGINT NOT NULL DEFAULT get_now (),
	CONSTRAINT products_promotions_pk PRIMARY KEY (product_id),
	CONSTRAINT products_promotions_product_id_fk FOREIGN key (product_id) REFERENCES products (product_id),
	CONSTRAINT products_promotions_promotion_id_fk FOREIGN key (promotion_id) REFERENCES promotions (promotion_id),
	CONSTRAINT products_promotions_active_from_ck CHECK (active_from > 0),
	CONSTRAINT products_promotions_active_to_ck CHECK (active_to > 0),
	CONSTRAINT products_promotions_updated_at_ck CHECK (updated_at > 0),
	CONSTRAINT products_promotions_created_at_ck CHECK (created_at > 0)
);

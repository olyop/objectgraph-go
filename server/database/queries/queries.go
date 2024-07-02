package queries

import _ "embed"

//go:embed select-brand-by-id.sql
var SelectBrandByID string

//go:embed select-brands-by-ids.sql
var SelectBrandsByIDs string

//go:embed select-categories-by-product-id.sql
var SelectCategoriesByProductID string

//go:embed select-category-by-id.sql
var SelectCategoryByID string

//go:embed select-categories-by-ids.sql
var SelectCategoriesByIDs string

//go:embed select-classification-by-id.sql
var SelectClassificationByID string

//go:embed select-classifications-by-ids.sql
var SelectClassificationsByIDs string

//go:embed select-contacts-by-user-id.sql
var SelectContactsByUserID string

//go:embed select-product-by-id.sql
var SelectProductByID string

//go:embed select-products-by-ids.sql
var SelectProductsByIDs string

//go:embed select-top-1000-products.sql
var SelectTop1000Products string

//go:embed select-user-by-id.sql
var SelectUserByID string

//go:embed select-users-by-ids.sql
var SelectUsersByIDs string

//go:embed select-top-1000-users.sql
var SelectTop1000Users string

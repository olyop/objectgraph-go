package queries

import _ "embed"

//go:embed select-brand-by-id.sql
var SelectBrandByID string

//go:embed select-categories-by-product-id.sql
var SelectCategoriesByProductID string

//go:embed select-classification-by-id.sql
var SelectClassificationByID string

//go:embed select-contacts-by-user-id.sql
var SelectContactsByUserID string

//go:embed select-product-by-id.sql
var SelectProductByID string

//go:embed select-top-1000-products.sql
var SelectTop1000Products string

//go:embed select-top-1000-users.sql
var SelectTop1000Users string

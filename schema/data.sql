INSERT INTO brands (name) VALUES ('Budweiser');
INSERT INTO brands (name) VALUES ('Heineken');
INSERT INTO brands (name) VALUES ('Calton Dry');
INSERT INTO brands (name) VALUES ('Oyster Bay');
INSERT INTO brands (name) VALUES ('Verve Clicquot');
INSERT INTO brands (name) VALUES ('Smirnoff');
INSERT INTO brands (name) VALUES ('Bundaberg');
INSERT INTO brands (name) VALUES ('Jack Daniels');
INSERT INTO brands (name) VALUES ('Coca Cola');
INSERT INTO brands (name) VALUES ('Red Bull');
INSERT INTO brands (name) VALUES ('Gift Box Co');

-- "a21d943c-562f-4eed-a5f4-38d191bd8f21"	"Budweiser"
-- "4818462e-ff64-4cfe-ba07-d179fdc08f88"	"Heineken"
-- "a5f07a00-9045-404c-bca6-ef8b73abb63d"	"Calton Dry"
-- "2ef80489-06db-4572-9230-1dd3beb05e04"	"Oyster Bay"
-- "6fbac30f-18f3-49d7-a4f3-51a253b78180"	"Verve Clicquot"
-- "3b88df6b-df65-426d-902c-fb17b539fa28"	"Smirnoff"
-- "f17fd4de-f90f-4b20-932a-1edc507f31ea"	"Bundaberg"
-- "a51c92b3-dcdf-4b28-972f-9987ac86c33e"	"Jack Daniels"
-- "a059ab47-a826-4933-9072-6c495a988bf8"	"Coca Cola"
-- "a481be10-438a-443d-9627-795db727fb80"	"Red Bull"
-- "1e500cd0-629a-4325-91cc-5afc63152eae"	"Gift Box Co"

INSERT INTO classifications (name) VALUES ('Beer');
INSERT INTO classifications (name) VALUES ('Cider');
INSERT INTO classifications (name) VALUES ('Wine');
INSERT INTO classifications (name) VALUES ('Spirits');
INSERT INTO classifications (name) VALUES ('Mixers');
INSERT INTO classifications (name) VALUES ('Pre-Mix');
INSERT INTO classifications (name) VALUES ('Gifts');

-- "03b14647-65ed-4f13-8a3f-b43f28d16a71"	"Beer"
-- "c3408ba5-b6af-4130-8552-818234a47b35"	"Cider"
-- "10f4f982-9711-4ace-be75-f73a086cae22"	"Wine"
-- "b714fda4-3ab3-4783-bd83-a833ab06c10a"	"Spirits"
-- "298db8d1-36cf-483c-a0d7-8efebd34f89e"	"Mixers"
-- "a7d18925-b9ae-4393-ad34-3286b990a69b"	"Pre-Mix"
-- "c2b15e1b-d7a0-4446-acd0-c2923d0c4e94"	"Gifts"

INSERT INTO categories (name, classification_id) VALUES ('Lager', '03b14647-65ed-4f13-8a3f-b43f28d16a71');
INSERT INTO categories (name, classification_id) VALUES ('Ale', '03b14647-65ed-4f13-8a3f-b43f28d16a71');
INSERT INTO categories (name, classification_id) VALUES ('Stout', '03b14647-65ed-4f13-8a3f-b43f28d16a71');
INSERT INTO categories (name, classification_id) VALUES ('Pear', 'c3408ba5-b6af-4130-8552-818234a47b35');
INSERT INTO categories (name, classification_id) VALUES ('Apple', 'c3408ba5-b6af-4130-8552-818234a47b35');
INSERT INTO categories (name, classification_id) VALUES ('Savingnon Blanc', '10f4f982-9711-4ace-be75-f73a086cae22');
INSERT INTO categories (name, classification_id) VALUES ('Chardonnay', '10f4f982-9711-4ace-be75-f73a086cae22');
INSERT INTO categories (name, classification_id) VALUES ('Merlot', '10f4f982-9711-4ace-be75-f73a086cae22');
INSERT INTO categories (name, classification_id) VALUES ('Pinot Noir', '10f4f982-9711-4ace-be75-f73a086cae22');
INSERT INTO categories (name, classification_id) VALUES ('Champagne', '10f4f982-9711-4ace-be75-f73a086cae22');
INSERT INTO categories (name, classification_id) VALUES ('Sparkling', '10f4f982-9711-4ace-be75-f73a086cae22');
INSERT INTO categories (name, classification_id) VALUES ('Whiskey', 'b714fda4-3ab3-4783-bd83-a833ab06c10a');
INSERT INTO categories (name, classification_id) VALUES ('Vodka', 'b714fda4-3ab3-4783-bd83-a833ab06c10a');
INSERT INTO categories (name, classification_id) VALUES ('Rum', 'b714fda4-3ab3-4783-bd83-a833ab06c10a');
INSERT INTO categories (name, classification_id) VALUES ('Gin', 'b714fda4-3ab3-4783-bd83-a833ab06c10a');
INSERT INTO categories (name, classification_id) VALUES ('Tonic', '298db8d1-36cf-483c-a0d7-8efebd34f89e');
INSERT INTO categories (name, classification_id) VALUES ('Soda', '298db8d1-36cf-483c-a0d7-8efebd34f89e');
INSERT INTO categories (name, classification_id) VALUES ('Bags', 'c2b15e1b-d7a0-4446-acd0-c2923d0c4e94');

-- "414b54d8-946e-4be4-b452-bcd50555246d"	"Lager"	"03b14647-65ed-4f13-8a3f-b43f28d16a71"
-- "58d4bf34-a3ca-4011-9707-9dcf0487f82a"	"Ale"	"03b14647-65ed-4f13-8a3f-b43f28d16a71"
-- "3d1504a2-9128-44ff-8896-e7ea9684e131"	"Stout"	"03b14647-65ed-4f13-8a3f-b43f28d16a71"
-- "557fbce8-73d2-491c-8db7-d16be16a4f92"	"Pear"	"c3408ba5-b6af-4130-8552-818234a47b35"
-- "c7cccf4d-4670-4f73-970f-c33d23a0364b"	"Apple"	"c3408ba5-b6af-4130-8552-818234a47b35"
-- "11153020-048d-46bd-8c3e-593a4322db73"	"Savingnon Blanc"	"10f4f982-9711-4ace-be75-f73a086cae22"
-- "61d5afa9-2c4f-4fe3-ab17-d4b7049eade0"	"Chardonnay"	"10f4f982-9711-4ace-be75-f73a086cae22"
-- "a7c6f13d-60a9-48bd-93e4-230231712f1b"	"Merlot"	"10f4f982-9711-4ace-be75-f73a086cae22"
-- "f38641fe-6421-4015-a0e0-7c9679d0e98b"	"Pinot Noir"	"10f4f982-9711-4ace-be75-f73a086cae22"
-- "751f9e46-76ba-468a-bfb0-4a23df2a66b3"	"Champagne"	"10f4f982-9711-4ace-be75-f73a086cae22"
-- "f86e1ad4-afeb-41ea-bb32-8f6cbe3f2efa"	"Sparkling"	"10f4f982-9711-4ace-be75-f73a086cae22"
-- "8601ac70-a10d-4e55-a9f6-1e0e9b43c4e9"	"Whiskey"	"b714fda4-3ab3-4783-bd83-a833ab06c10a"
-- "f5e8a36a-d3ed-4ac6-ae4c-fc4bffb2e3a5"	"Vodka"	"b714fda4-3ab3-4783-bd83-a833ab06c10a"
-- "d5d16e60-4a76-43b3-8d86-664a76264a79"	"Rum"	"b714fda4-3ab3-4783-bd83-a833ab06c10a"
-- "16b2c77f-dcd2-44ae-a9de-0de81a9e7828"	"Gin"	"b714fda4-3ab3-4783-bd83-a833ab06c10a"
-- "a689c6b8-f185-43b5-93b7-5e291258e62a"	"Tonic"	"298db8d1-36cf-483c-a0d7-8efebd34f89e"
-- "9b9e3428-8fa0-4cae-be8f-7beede7c3b7e"	"Soda"	"298db8d1-36cf-483c-a0d7-8efebd34f89e"
-- "64c9db9f-f70b-4390-9095-52f87f421063"	"Bags"	"c2b15e1b-d7a0-4446-acd0-c2923d0c4e94"

INSERT INTO products (name, brand_id) VALUES ('Budweiser Lager Beer Bottle', 'a21d943c-562f-4eed-a5f4-38d191bd8f21');
INSERT INTO products (name, brand_id) VALUES ('Heineken Lager Beer Bottle', '4818462e-ff64-4cfe-ba07-d179fdc08f88');
INSERT INTO products (name, brand_id) VALUES ('Calton Dry Lager Beer Bottle', 'a5f07a00-9045-404c-bca6-ef8b73abb63d');
INSERT INTO products (name, brand_id) VALUES ('Oyster Bay Sauvignon Blanc', '2ef80489-06db-4572-9230-1dd3beb05e04');
INSERT INTO products (name, brand_id) VALUES ('Verve Clicquot Brut Champagne', '6fbac30f-18f3-49d7-a4f3-51a253b78180');
INSERT INTO products (name, brand_id) VALUES ('Smirnoff Red Vodka', '3b88df6b-df65-426d-902c-fb17b539fa28');
INSERT INTO products (name, brand_id) VALUES ('Bundaberg Rum', 'f17fd4de-f90f-4b20-932a-1edc507f31ea');
INSERT INTO products (name, brand_id) VALUES ('Jack Daniels Whiskey', 'a51c92b3-dcdf-4b28-972f-9987ac86c33e');
INSERT INTO products (name, brand_id) VALUES ('Coca Cola', 'a059ab47-a826-4933-9072-6c495a988bf8');
INSERT INTO products (name, brand_id) VALUES ('Red Bull', 'a481be10-438a-443d-9627-795db727fb80');
INSERT INTO products (name, brand_id) VALUES ('Gift Box Co Gift Box', '1e500cd0-629a-4325-91cc-5afc63152eae');

-- "54545658-d5e7-47c1-be4b-0e4718a93d02"	"Budweiser Lager Beer Bottle"	"a21d943c-562f-4eed-a5f4-38d191bd8f21"
-- "9e173255-21f3-46f8-a712-908f6381d169"	"Heineken Lager Beer Bottle"	"4818462e-ff64-4cfe-ba07-d179fdc08f88"
-- "335a9306-7ad0-46ff-81fc-ab383c9065cb"	"Calton Dry Lager Beer Bottle"	"a5f07a00-9045-404c-bca6-ef8b73abb63d"
-- "6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	"Oyster Bay Sauvignon Blanc"	"2ef80489-06db-4572-9230-1dd3beb05e04"
-- "f735067d-938c-443a-9271-8c29dbc2868c"	"Verve Clicquot Brut Champagne"	"6fbac30f-18f3-49d7-a4f3-51a253b78180"
-- "85231feb-1ceb-4e99-b54a-7233f327c915"	"Smirnoff Red Vodka"	"3b88df6b-df65-426d-902c-fb17b539fa28"
-- "1bb45678-87e6-464d-b8ac-fc627bf54e1e"	"Bundaberg Rum"	"f17fd4de-f90f-4b20-932a-1edc507f31ea"
-- "06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	"Jack Daniels Whiskey"	"a51c92b3-dcdf-4b28-972f-9987ac86c33e"
-- "1d021865-dbe4-49e9-b7cd-e117d079ac98"	"Coca Cola"	"a059ab47-a826-4933-9072-6c495a988bf8"
-- "abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5"	"Red Bull"	"a481be10-438a-443d-9627-795db727fb80"
-- "53c24139-1497-43c0-861d-847789f4adb7"	"Gift Box Co Gift Box"	"1e500cd0-629a-4325-91cc-5afc63152eae"

INSERT INTO products_categories (product_id, category_id) VALUES ('54545658-d5e7-47c1-be4b-0e4718a93d02', '414b54d8-946e-4be4-b452-bcd50555246d');
INSERT INTO products_categories (product_id, category_id) VALUES ('9e173255-21f3-46f8-a712-908f6381d169', '414b54d8-946e-4be4-b452-bcd50555246d');
INSERT INTO products_categories (product_id, category_id) VALUES ('335a9306-7ad0-46ff-81fc-ab383c9065cb', '414b54d8-946e-4be4-b452-bcd50555246d');
INSERT INTO products_categories (product_id, category_id) VALUES ('6fb7b4e7-b771-4246-8f4d-a56b3d16316f', '11153020-048d-46bd-8c3e-593a4322db73');
INSERT INTO products_categories (product_id, category_id) VALUES ('f735067d-938c-443a-9271-8c29dbc2868c', '751f9e46-76ba-468a-bfb0-4a23df2a66b3');
INSERT INTO products_categories (product_id, category_id) VALUES ('85231feb-1ceb-4e99-b54a-7233f327c915', 'f5e8a36a-d3ed-4ac6-ae4c-fc4bffb2e3a5');
INSERT INTO products_categories (product_id, category_id) VALUES ('1bb45678-87e6-464d-b8ac-fc627bf54e1e', 'd5d16e60-4a76-43b3-8d86-664a76264a79');
INSERT INTO products_categories (product_id, category_id) VALUES ('06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d', '8601ac70-a10d-4e55-a9f6-1e0e9b43c4e9');
INSERT INTO products_categories (product_id, category_id) VALUES ('1d021865-dbe4-49e9-b7cd-e117d079ac98', '9b9e3428-8fa0-4cae-be8f-7beede7c3b7e');
INSERT INTO products_categories (product_id, category_id) VALUES ('abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5', '9b9e3428-8fa0-4cae-be8f-7beede7c3b7e');

-- "54545658-d5e7-47c1-be4b-0e4718a93d02"	"414b54d8-946e-4be4-b452-bcd50555246d"
-- "9e173255-21f3-46f8-a712-908f6381d169"	"414b54d8-946e-4be4-b452-bcd50555246d"
-- "335a9306-7ad0-46ff-81fc-ab383c9065cb"	"414b54d8-946e-4be4-b452-bcd50555246d"
-- "6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	"11153020-048d-46bd-8c3e-593a4322db73"
-- "f735067d-938c-443a-9271-8c29dbc2868c"	"751f9e46-76ba-468a-bfb0-4a23df2a66b3"
-- "85231feb-1ceb-4e99-b54a-7233f327c915"	"f5e8a36a-d3ed-4ac6-ae4c-fc4bffb2e3a5"
-- "1bb45678-87e6-464d-b8ac-fc627bf54e1e"	"d5d16e60-4a76-43b3-8d86-664a76264a79"
-- "06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	"8601ac70-a10d-4e55-a9f6-1e0e9b43c4e9"
-- "1d021865-dbe4-49e9-b7cd-e117d079ac98"	"9b9e3428-8fa0-4cae-be8f-7beede7c3b7e"
-- "abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5"	"9b9e3428-8fa0-4cae-be8f-7beede7c3b7e"

INSERT INTO products_abv (product_id, abv) VALUES ('54545658-d5e7-47c1-be4b-0e4718a93d02', 4.5);
INSERT INTO products_abv (product_id, abv) VALUES ('9e173255-21f3-46f8-a712-908f6381d169', 4.5);
INSERT INTO products_abv (product_id, abv) VALUES ('335a9306-7ad0-46ff-81fc-ab383c9065cb', 4.5);
INSERT INTO products_abv (product_id, abv) VALUES ('6fb7b4e7-b771-4246-8f4d-a56b3d16316f', 13);
INSERT INTO products_abv (product_id, abv) VALUES ('f735067d-938c-443a-9271-8c29dbc2868c', 20);
INSERT INTO products_abv (product_id, abv) VALUES ('85231feb-1ceb-4e99-b54a-7233f327c915', 37.5);
INSERT INTO products_abv (product_id, abv) VALUES ('1bb45678-87e6-464d-b8ac-fc627bf54e1e', 37.5);
INSERT INTO products_abv (product_id, abv) VALUES ('06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d', 40);

-- "54545658-d5e7-47c1-be4b-0e4718a93d02"	4.5
-- "9e173255-21f3-46f8-a712-908f6381d169"	4.5
-- "335a9306-7ad0-46ff-81fc-ab383c9065cb"	4.5
-- "6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	13
-- "f735067d-938c-443a-9271-8c29dbc2868c"	20
-- "85231feb-1ceb-4e99-b54a-7233f327c915"	37.5
-- "1bb45678-87e6-464d-b8ac-fc627bf54e1e"	37.5
-- "06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	40

INSERT INTO products_prices (product_id, price, is_current) VALUES ('54545658-d5e7-47c1-be4b-0e4718a93d02', 2500, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('9e173255-21f3-46f8-a712-908f6381d169', 2500, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('335a9306-7ad0-46ff-81fc-ab383c9065cb', 2500, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('6fb7b4e7-b771-4246-8f4d-a56b3d16316f', 5000, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('f735067d-938c-443a-9271-8c29dbc2868c', 20000, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('85231feb-1ceb-4e99-b54a-7233f327c915', 20000, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('1bb45678-87e6-464d-b8ac-fc627bf54e1e', 20000, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d', 20000, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('1d021865-dbe4-49e9-b7cd-e117d079ac98', 500, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5', 500, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('53c24139-1497-43c0-861d-847789f4adb7', 1000, TRUE);
INSERT INTO products_prices (product_id, price, is_current) VALUES ('53c24139-1497-43c0-861d-847789f4adb7', 1200, FALSE);

-- "db910220-d497-4921-a454-e54fa9007a5c"	"54545658-d5e7-47c1-be4b-0e4718a93d02"	2500	true
-- "d56348b7-80f4-449d-bdae-c11564ec1b75"	"9e173255-21f3-46f8-a712-908f6381d169"	2500	true
-- "d37dec82-4522-429d-94ef-dfc2ba677183"	"335a9306-7ad0-46ff-81fc-ab383c9065cb"	2500	true
-- "9fb1f687-0de7-483d-880a-4733c8a6c2d1"	"6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	5000	true
-- "90d6d8d0-21b6-400e-b6a3-973f9a0bdbff"	"f735067d-938c-443a-9271-8c29dbc2868c"	20000	true
-- "8dd27a1e-277e-45d8-b56a-7175dba0d12b"	"85231feb-1ceb-4e99-b54a-7233f327c915"	20000	true
-- "30fd6f1c-9254-4f2f-aa1f-f5eeb26c2e4d"	"1bb45678-87e6-464d-b8ac-fc627bf54e1e"	20000	true
-- "148a968b-3196-43a5-979b-01c937efcd67"	"06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	20000	true
-- "67423091-b805-4f72-87b2-90a9af44185c"	"1d021865-dbe4-49e9-b7cd-e117d079ac98"	500	true
-- "45128933-9eed-4c8e-8bc6-57750f0096fb"	"abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5"	500	true
-- "1df5a677-60f6-4887-a91b-9f23a63f61fe"	"53c24139-1497-43c0-861d-847789f4adb7"	1000	true

INSERT INTO products_volumes (product_id, volume) VALUES ('54545658-d5e7-47c1-be4b-0e4718a93d02', 330);
INSERT INTO products_volumes (product_id, volume) VALUES ('9e173255-21f3-46f8-a712-908f6381d169', 330);
INSERT INTO products_volumes (product_id, volume) VALUES ('335a9306-7ad0-46ff-81fc-ab383c9065cb', 330);
INSERT INTO products_volumes (product_id, volume) VALUES ('6fb7b4e7-b771-4246-8f4d-a56b3d16316f', 750);
INSERT INTO products_volumes (product_id, volume) VALUES ('f735067d-938c-443a-9271-8c29dbc2868c', 750);
INSERT INTO products_volumes (product_id, volume) VALUES ('85231feb-1ceb-4e99-b54a-7233f327c915', 750);
INSERT INTO products_volumes (product_id, volume) VALUES ('1bb45678-87e6-464d-b8ac-fc627bf54e1e', 750);
INSERT INTO products_volumes (product_id, volume) VALUES ('06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d', 750);
INSERT INTO products_volumes (product_id, volume) VALUES ('1d021865-dbe4-49e9-b7cd-e117d079ac98', 330);
INSERT INTO products_volumes (product_id, volume) VALUES ('abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5', 330);

-- "54545658-d5e7-47c1-be4b-0e4718a93d02"	330
-- "9e173255-21f3-46f8-a712-908f6381d169"	330
-- "335a9306-7ad0-46ff-81fc-ab383c9065cb"	330
-- "6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	750
-- "f735067d-938c-443a-9271-8c29dbc2868c"	750
-- "85231feb-1ceb-4e99-b54a-7233f327c915"	750
-- "1bb45678-87e6-464d-b8ac-fc627bf54e1e"	750
-- "06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	750
-- "1d021865-dbe4-49e9-b7cd-e117d079ac98"	330
-- "abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5"	330

-- Fix Prices
-- Add price column to products table that references the price_id in products_prices

-- "54545658-d5e7-47c1-be4b-0e4718a93d02"	"Budweiser Lager Beer Bottle"	"a21d943c-562f-4eed-a5f4-38d191bd8f21"	
-- "9e173255-21f3-46f8-a712-908f6381d169"	"Heineken Lager Beer Bottle"	"4818462e-ff64-4cfe-ba07-d179fdc08f88"	
-- "335a9306-7ad0-46ff-81fc-ab383c9065cb"	"Calton Dry Lager Beer Bottle"	"a5f07a00-9045-404c-bca6-ef8b73abb63d"	
-- "6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	"Oyster Bay Sauvignon Blanc"	"2ef80489-06db-4572-9230-1dd3beb05e04"	
-- "f735067d-938c-443a-9271-8c29dbc2868c"	"Verve Clicquot Brut Champagne"	"6fbac30f-18f3-49d7-a4f3-51a253b78180"	
-- "85231feb-1ceb-4e99-b54a-7233f327c915"	"Smirnoff Red Vodka"	"3b88df6b-df65-426d-902c-fb17b539fa28"	
-- "1bb45678-87e6-464d-b8ac-fc627bf54e1e"	"Bundaberg Rum"	"f17fd4de-f90f-4b20-932a-1edc507f31ea"	
-- "06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	"Jack Daniels Whiskey"	"a51c92b3-dcdf-4b28-972f-9987ac86c33e"	
-- "1d021865-dbe4-49e9-b7cd-e117d079ac98"	"Coca Cola"	"a059ab47-a826-4933-9072-6c495a988bf8"	
-- "abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5"	"Red Bull"	"a481be10-438a-443d-9627-795db727fb80"	
-- "53c24139-1497-43c0-861d-847789f4adb7"	"Gift Box Co Gift Box"	"1e500cd0-629a-4325-91cc-5afc63152eae"

-- "db910220-d497-4921-a454-e54fa9007a5c"	"54545658-d5e7-47c1-be4b-0e4718a93d02"	2500
-- "d56348b7-80f4-449d-bdae-c11564ec1b75"	"9e173255-21f3-46f8-a712-908f6381d169"	2500
-- "d37dec82-4522-429d-94ef-dfc2ba677183"	"335a9306-7ad0-46ff-81fc-ab383c9065cb"	2500
-- "9fb1f687-0de7-483d-880a-4733c8a6c2d1"	"6fb7b4e7-b771-4246-8f4d-a56b3d16316f"	5000
-- "90d6d8d0-21b6-400e-b6a3-973f9a0bdbff"	"f735067d-938c-443a-9271-8c29dbc2868c"	20000
-- "8dd27a1e-277e-45d8-b56a-7175dba0d12b"	"85231feb-1ceb-4e99-b54a-7233f327c915"	20000
-- "30fd6f1c-9254-4f2f-aa1f-f5eeb26c2e4d"	"1bb45678-87e6-464d-b8ac-fc627bf54e1e"	20000
-- "148a968b-3196-43a5-979b-01c937efcd67"	"06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d"	20000
-- "67423091-b805-4f72-87b2-90a9af44185c"	"1d021865-dbe4-49e9-b7cd-e117d079ac98"	500
-- "45128933-9eed-4c8e-8bc6-57750f0096fb"	"abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5"	500
-- "1df5a677-60f6-4887-a91b-9f23a63f61fe"	"53c24139-1497-43c0-861d-847789f4adb7"	1000
-- "3d0fbecb-d7ed-4ddd-a489-b453e9f799e8"	"53c24139-1497-43c0-861d-847789f4adb7"	1200

UPDATE products SET price_id = 'db910220-d497-4921-a454-e54fa9007a5c' WHERE product_id = '54545658-d5e7-47c1-be4b-0e4718a93d02';
UPDATE products SET price_id = 'd56348b7-80f4-449d-bdae-c11564ec1b75' WHERE product_id = '9e173255-21f3-46f8-a712-908f6381d169';
UPDATE products SET price_id = 'd37dec82-4522-429d-94ef-dfc2ba677183' WHERE product_id = '335a9306-7ad0-46ff-81fc-ab383c9065cb';
UPDATE products SET price_id = '9fb1f687-0de7-483d-880a-4733c8a6c2d1' WHERE product_id = '6fb7b4e7-b771-4246-8f4d-a56b3d16316f';
UPDATE products SET price_id = '90d6d8d0-21b6-400e-b6a3-973f9a0bdbff' WHERE product_id = 'f735067d-938c-443a-9271-8c29dbc2868c';
UPDATE products SET price_id = '8dd27a1e-277e-45d8-b56a-7175dba0d12b' WHERE product_id = '85231feb-1ceb-4e99-b54a-7233f327c915';
UPDATE products SET price_id = '30fd6f1c-9254-4f2f-aa1f-f5eeb26c2e4d' WHERE product_id = '1bb45678-87e6-464d-b8ac-fc627bf54e1e';
UPDATE products SET price_id = '148a968b-3196-43a5-979b-01c937efcd67' WHERE product_id = '06861ced-2b4c-4d77-af5d-b4a0e7eb2a4d';
UPDATE products SET price_id = '67423091-b805-4f72-87b2-90a9af44185c' WHERE product_id = '1d021865-dbe4-49e9-b7cd-e117d079ac98';
UPDATE products SET price_id = '45128933-9eed-4c8e-8bc6-57750f0096fb' WHERE product_id = 'abc7c3d5-ece5-45f2-b0c1-b31fcd79a7d5';
UPDATE products SET price_id = '1df5a677-60f6-4887-a91b-9f23a63f61fe' WHERE product_id = '53c24139-1497-43c0-861d-847789f4adb7';

package tables

const ProductsTable = `create table if not exists products (
	id BIGSERIAL PRIMARY KEY,
	type product_type not null,
	title text not null,
	description text not null,
	price int not null,
	publish bool default true,
	created_at timestamp default now()
)`

// transfer category_id into a separated table?
// i don't want a category_id to be essential while creating a product,
// since when there is a need to create a constuctor type product you have to
// assign a product to a specific category, which can be undefined.
// This is why you want to create a table where category_id have to be referenced
// to a real one, but is not essential while creating a product.
// By the way, it makes managing more difficult, especially in terms of modyfing
// and extracting data.
//
// - Done

package tables

const ProductsCategoriesTable = `create table if not exists products_categories (
	id bigserial primary key,
	product_id bigint references products(id) not null,
	category_id bigint references categories(id) not null
)`

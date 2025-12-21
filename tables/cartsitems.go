package tables

const CartsItemsTable = `create table if not exists carts_items (
	id bigserial primary key,
	cart_id bigint references carts(id) on delete cascade not null,
	product_id bigint references products(id) on delete cascade not null,
	count int default 1
)`

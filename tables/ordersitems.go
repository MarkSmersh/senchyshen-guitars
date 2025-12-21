package tables

const OrdersItemsTable = `create table if not exists orders_items (
	id bigserial primary key,
	order_id bigint references orders(id) on delete cascade not null,
	product_id bigint references products(id) on delete set null,
	title text not null,
	description text not null,
	image_id bigint references images(id) on delete set null,
	price int NOT NULL,
	count int not null
)`

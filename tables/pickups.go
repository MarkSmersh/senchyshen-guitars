package tables

const PickupsTable = `create table if not exists pickups (
	id bigserial primary key,
	product_id bigint unique references products(id) on delete cascade not null,
	type pickup_type not null
)`

// why does a reference of a product_id accept null?..

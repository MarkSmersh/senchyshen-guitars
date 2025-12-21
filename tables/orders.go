package tables

const OrdersTable = `create table if not exists orders (
	id bigserial primary key,
	uuid uuid unique default gen_random_uuid(),
	tel text NOT NULL,
	comment VARCHAR(256) DEFAULT ''
)`

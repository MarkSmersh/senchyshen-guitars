package tables

const CartsTable = `create table if not exists carts (
	id bigserial primary key,
	uuid uuid unique default gen_random_uuid()
)`

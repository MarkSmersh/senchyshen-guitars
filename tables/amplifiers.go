package tables

const AmplifiersTable = `create table if not exists amplifiers (
	id bigserial primary key,
	product_id bigint unique references products(id) on delete cascade not null ,
	power bigint
)`

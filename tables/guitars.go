package tables

const GuitarsTable = `create table if not exists guitars (
	id bigserial primary key,
	product_id bigint unique references products(id) on delete cascade not null ,
	strings_count int NOT NULL,
	bodyshape text NOT NULL,
	color color NOT NULL,
	pickups_count int NOT NULL
)`

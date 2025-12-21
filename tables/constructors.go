package tables

const ConstructorsTable = `create table if not exists constructors (
	id bigserial primary key,
	product_id bigint unique references products(id) on delete cascade not null,
	bodyshape_id bigint references bodyshapes(id) not null,
	color color,
	created_at timestamp default now()
)`

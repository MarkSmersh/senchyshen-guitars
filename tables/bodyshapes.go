package tables

const BodyshapesTable = `create table if not exists bodyshapes (
	id bigserial primary key,
	product_id bigint unique references products(id) on delete cascade not null,
	color color NOT NULL,
	material text NOT NULL,
	height int NOT NULL,
	width int NOT NULL,
	thickness int NOT NULL,
	pickups_count int NOT NULL
)`

package tables

const ConstructorsPickupsTable = `create table if not exists constructors_pickups (
	id bigserial primary key,
	constructor_id bigint references constructors(id) on delete cascade,
	pickup_id bigint references pickups(id) on delete cascade not null,
	position pickup_pos NOT NULL
)`

package tables

const ImagesTable = `create table if not exists images (
	id bigserial primary key,
	uuid uuid unique default gen_random_uuid(),
	ext varchar(4) not null
)`

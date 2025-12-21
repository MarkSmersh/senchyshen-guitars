package tables

// Kategoria w danym wypadku nie jest zbiorem jednego product_type,
// ale jest po prostu zbiorem produktów, niezależnie od ich typu.
// Kategorii są stworzone nie dla tego, aby sortować według nich produkty -
// to można zrobić za pomocą SELECT DISTINCT, ale żeby ułatwić użytkowniku
// wyszukiwania potzebnego mu produktu
//
// Cochiaż, ze względu na optymizację, to nie jest najszybszy sposób,
// ale jest lepszy ze względu na go architekturę.

const CategoriesTable = `create table if not exists categories (
	id bigserial primary key,
	image_id bigint references images(id) on delete set null,
	title text not null,
	description text not null
)`

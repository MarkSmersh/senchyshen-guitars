package tables

const Enums = `
	create type product_type as enum
		('guitar', 'pickup', 'bodyshape', 'amplifier', 'crafted');
	create type pickup_pos as enum
		('top', 'bottom', 'middle');
	create type pickup_type as enum
		('single', 'humbucker', 'p90');
`

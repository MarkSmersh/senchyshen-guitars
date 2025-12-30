package tables

const Domains = `
	create domain color as varchar(6) check (value ~ '^[0-9A-Fa-f]{6}$');
`

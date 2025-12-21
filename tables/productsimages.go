package tables

const ProductsImagesTable = `create table if not exists products_images (
	id BIGSERIAL PRIMARY KEY,
	product_id bigint references products(id),
	image_id bigint references images(id) on delete cascade
)`

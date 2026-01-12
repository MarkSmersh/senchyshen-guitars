package models

import (
	"context"
	"log/slog"
)

func (p *Product) GetCategories() ([]Category, error) {
	row := p.conn.QueryRow(
		context.Background(),
		"select json_array(select (json_build_object('id', c.id, 'title', c.title, 'description', c.description, 'image', coalesce(i.uuid||'.'||i.ext::text, '.'))) from categories c left join images i on i.id = c.image_id)",
	)

	categories := []Category{}

	if err := row.Scan(&categories); err != nil {
		slog.Error(err.Error())
		return categories, InternalServerError()
	}

	return categories, nil
}

type CreateCategoryReq struct {
	Category
	Products []int `json:"products"`
}

func (p *Product) CreateCategory(params CreateCategoryReq) (int, error) {
	row := p.conn.QueryRow(
		context.Background(),
		"with cat as (insert into categories (title, description) values ($1, $2) returning id), pcs as (insert into products_categories (category_id, product_id) select id, unnest($3::bigint[]) from cat) select id from cat",
		params.Title,
		params.Description,
		params.Products,
	)

	id := 0

	if err := row.Scan(&id); err != nil {
		slog.Error(err.Error())
		return id, InternalServerError()
	}

	return id, nil
}

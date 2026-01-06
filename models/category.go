package models

import (
	"context"
	"log/slog"
)

func (p *Product) GetCategories() ([]Category, error) {
	row := p.conn.QueryRow(
		context.Background(),
		"select json_array(select (json_build_object('id', id,'title', title, 'description', description, 'image', coalesce(image_id::text, '.'))) from categories c)",
	)

	categories := []Category{}

	if err := row.Scan(&categories); err != nil {
		slog.Error(err.Error())
		return categories, InternalServerError()
	}

	return categories, nil
}

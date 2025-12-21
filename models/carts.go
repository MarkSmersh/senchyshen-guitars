package models

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Cart struct {
	conn *pgx.Conn
}

func NewCart(conn *pgx.Conn) Cart {
	return Cart{conn: conn}
}

func (c Cart) Create() (string, error) {
	row := c.conn.QueryRow(
		context.Background(),
		"insert into carts default values returning uuid",
	)

	var uuid string

	if err := row.Scan(&uuid); err != nil {
		slog.Error(err.Error())
		return "", InternalServerError()
	}

	return uuid, nil
}

func (c Cart) Clear(uuid string) error {
	_, err := c.conn.Exec(
		context.Background(),
		"delete from carts_items where cart_id = $1",
		uuid,
	)

	if err != nil {
		slog.Error(err.Error())
		return InternalServerError()
	}

	return nil
}

type CartAddProduct struct {
	UUID      string `json:"uuid"`
	ProductID int    `json:"productId"`
	Count     int    `json:"count"`
}

func (c Cart) AddProduct(params CartAddProduct) error {
	_, err := c.conn.Exec(
		context.Background(),
		"insert into carts_items (cart_id, product_id, count) values ((select id from carts where uuid = $1), $2, $3)",
		params.UUID,
		params.ProductID,
		params.Count,
	)

	if err != nil {
		slog.Error(err.Error())
		return InternalServerError()
	}

	return nil
}

func (c Cart) RemoveProduct(uuid string, productID string) error {
	tag, err := c.conn.Exec(
		context.Background(),
		"delete from carts_items where cart_id = (select id from carts where uuid = $1) and product_id = $2",
		uuid,
		productID,
	)

	if err != nil {
		slog.Error(err.Error())
		return InternalServerError()
	}

	if tag.RowsAffected() <= 0 {
		return NewApiError(404, "Koszyk nie zawiera danego produktu")
	}

	return nil
}

type GetProductsModel struct {
	ID          int    `json:"id"`
	Count       int    `json:"count"`
	Total       int    `json:"total"`
	Price       int    `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (c Cart) GetProducts(uuid string) ([]GetProductsModel, error) {
	rows, err := c.conn.Query(
		context.Background(),
		`
		select ci.product_id, SUM(ci.count*p.price), SUM(ci.count), p.price, p.title, p.description, i.uuid, i.ext from carts c join carts_items ci on ci.cart_id = c.id join products p on p.id = ci.product_id left join products_images pi on pi.product_id = p.id left join lateral (select uuid, ext from images i where i.id = pi.image_id limit 1) i on true where c.uuid = $1 group by ci.product_id, p.title, p.description, i.uuid, i.ext, p.price
		`,
		uuid,
	)

	products := []GetProductsModel{}

	if err != nil {
		slog.Error(err.Error())
		return products, InternalServerError()
	}

	for {
		if !rows.Next() {
			break
		}

		p := GetProductsModel{}

		var uuid, ext sql.NullString

		if err := rows.Scan(
			&p.ID, &p.Total, &p.Count, &p.Price, &p.Title, &p.Description, &uuid, &ext,
		); err != nil {
			slog.Error(err.Error())
			return products, InternalServerError()
		}

		if uuid.Valid && ext.Valid {
			p.Image = uuid.String + "." + ext.String
		}

		products = append(products, p)
	}

	return products, nil
}

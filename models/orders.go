package models

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	conn *pgxpool.Pool
}

func NewOrder(conn *pgxpool.Pool) Order {
	return Order{conn: conn}
}

type OrderCreate struct {
	CartUUID string `json:"cartUUID"`
	Tel      string `json:"tel"`
	Comment  string `json:"comment"`
}

// Tworzy zamówienie na podstawie koszyka, gdzie każdy element
// zawiera nazwę, opis, cenę i pierwsze najstarsze zdjęcię produktu;
// usuwa wszystkie cart_items stosujących podanego koszyka oraz
// zwraca uuid stworzonego zamówenia
//
// FIXME: Jeżeli wstawić niestniejący UUID koszyka, to zamównie będzie
// stworzone, cochiaż nie musze :(
// Dodatkowo, jeżeli sprobować znajdź przez /api/orders/:uuid takie zamówenie,
// to nie będzie znajdzione. Ale to raczej z powodu tego, że w nim jest brak
// produktow. Najprowdopobniej problem jest w JOIN
func (o Order) Create(params OrderCreate) (string, error) {
	//
	// Here’s an honest, expert review of your SQL.
	// Short version: It works, but it can be cleaner, safer, and faster.
	//
	// ⭐ Rating: 6.5 / 10
	//
	// - ChatGPT 12.10.2025

	row := o.conn.QueryRow(
		context.Background(),
		`
		with crt as (select id from carts where uuid = $1), rows as (select ci.product_id, p.price, p.title, p.description, pi.image_id, SUM(ci.count) as count from carts_items ci join products p on p.id = ci.product_id left join lateral (select * from products_images pi where pi.product_id = p.id limit 1) pi on true where ci.cart_id = (select id from crt) group by ci.product_id, p.price, p.title, p.description, pi.image_id, ci.count), ord as (insert into orders (tel, comment) values ($2, $3) returning id, uuid), ins as (insert into orders_items (order_id, product_id, title, description, image_id, price, count) select (select id from ord), rows.product_id, rows.title, rows.description, rows.image_id, rows.price, rows.count from rows), del as (delete from carts_items where cart_id = (select id from crt)) select uuid from ord;
		`,
		params.CartUUID,
		params.Tel,
		params.Comment,
	)

	var uuid string

	if err := row.Scan(&uuid); err != nil {
		slog.Error(err.Error())
		return "", InternalServerError()
	}

	return uuid, nil
}

type OrderItemModel struct {
	ProductID   int    `json:"productId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
}

type OrderModel struct {
	UUID    string           `json:"uuid"`
	Tel     string           `json:"tel"`
	Comment string           `json:"comment"`
	Items   []OrderItemModel `json:"items"`
}

func (o Order) Find(uuid string) (OrderModel, error) {
	rows, err := o.conn.Query(
		context.Background(),
		`
		select o.uuid, o.tel, o.comment, oi.product_id, oi.title, oi.description, oi.price, oi.count,
			i.uuid, i.ext
		from orders o
		join orders_items oi on oi.order_id = o.id
		left join images i on i.id = oi.image_id
		where o.uuid = $1;
		`,
		uuid,
	)

	order := OrderModel{}

	if err != nil {
		slog.Error(err.Error())
		return order, InternalServerError()
	}

	order = OrderModel{}

	for {
		if !rows.Next() {
			break
		}

		// var uuid, tel, comment, title, description string
		// var pid, price, count int
		var imgUuid, ext sql.NullString

		item := OrderItemModel{}

		if err := rows.Scan(
			&order.UUID, &order.Tel, &order.Comment, &item.ProductID,
			&item.Title, &item.Description, &item.Price, &item.Count,
			&imgUuid, &ext,
		); err != nil {
			slog.Error(err.Error())
			return order, InternalServerError()
		}

		if len(order.Items) <= 0 {
			order.Items = append(order.Items, item)
		}

		lastProduct := &order.Items[len(order.Items)-1]

		if lastProduct.ProductID != item.ProductID {
			order.Items = append(order.Items, item)
		}

		if imgUuid.Valid && ext.Valid {
			lastProduct.Image = imgUuid.String + "." + ext.String
		}
	}

	rows.Close()

	if rows.CommandTag().RowsAffected() <= 0 {
		return order, NewApiError(404, "Nie znalieoziono zamówenia")
	}

	return order, nil
}

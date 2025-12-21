package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Product struct {
	conn *pgx.Conn
}

func NewProduct(conn *pgx.Conn) Product {
	return Product{conn: conn}
}

type ProductData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p Product) Create(title string, description string, price float64) error {
	tag, err := p.conn.Exec(
		context.Background(),
		"INSERT INTO products (title, description, price)",
		title,
		description,
		price,
	)

	if err != nil {
		return InternalServerError()
	}

	if tag.RowsAffected() <= 0 {
		return NewApiError(500, "Produkt nie jest stworzony.")
	}

	return nil
}

type ImageModel struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

type FindManyParams struct {
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	Types      *[]string `json:"types"`
	PriceMin   float64   `json:"priceMin"`
	PriceMax   float64   `json:"priceMax"`
	CategoryID int       `json:"category"`
	ProductID  int       `json:"id"`
	// title, price, createdAt
	OrderBy string `json:"orderBy"`
	// desc | asc
	Order string `json:"order"`
	Query string `json:"query"`
}

var orderBys = []string{"title", "price", "createdAt"}
var order = []string{"desc", "asc"}

// wyciąga z bazy danych produkty oraz ich zdjęcia,
// przy wprowadzeniu parametru pid > 0, zwraca tylko produkt z podanym id;
// dla takich sytuacji jest wykorzystana metoda Find

func (p Product) FindMany(params FindManyParams) ([]ProductModel, error) {
	switch params.OrderBy {
	case "title":
		break
	case "price":
		break
	case "createdAt":
		params.OrderBy = "created_at"
		break
	default:
		params.OrderBy = "created_at"
		break
	}

	switch params.Order {
	case "desc":
		break
	case "asc":
		break
	default:
		params.Order = "desc"
		break
	}

	// - no you cannot pull variables via sprintf bfacause
	//   it will lead to an sql injection 😭😭😭
	// - idc, it just works

	q := fmt.Sprintf(`
WITH ROWS AS (
    SELECT
        id
    FROM
        products
    WHERE
        CASE
            WHEN $1 > 0 THEN id = $1
            ELSE price >= $3
            AND CASE
                WHEN $4 > 0 THEN price <= $4
                ELSE TRUE
            END
			AND ($2::text[] IS NULL OR type::text = ANY($2::text[]))
            AND CASE
                WHEN $8 != '' THEN
                to_tsvector(title || ' ' || description) @@ to_tsquery($8)
                ELSE TRUE
            END
        END
        AND publish = TRUE
    ORDER BY
        %[1]s %[2]s
		LIMIT $6 OFFSET $6::int * $7::int
)
SELECT
    p.id,
    p.title,
    p.description,
    p.price,
    p.type :: text,
    extract(
        epoch
        FROM
            p.created_at
    ) :: bigint,
    i.id,
    i.uuid,
    i.ext,
    c.id,
    c.title
FROM
    products p
    LEFT JOIN products_images pi ON pi.product_id = p.id
    LEFT JOIN images i ON i.id = pi.image_id
    LEFT JOIN products_categories pc ON pc.product_id = p.id
    LEFT JOIN categories c ON c.id = pc.category_id
WHERE
    p.id IN (
        SELECT
            id
        FROM
            ROWS
    )
    AND CASE
        WHEN $5 > 0 THEN pc.category_id = $5
        ELSE TRUE
    END
ORDER BY
    p.%[1]s %[2]s

		`,
		params.OrderBy, params.Order,
	)

	rows, err := p.conn.Query(
		context.Background(),
		q,
		params.ProductID,
		params.Types,
		params.PriceMin,
		params.PriceMax,
		params.CategoryID,
		params.Limit,
		params.Page,
		params.Query,
	)

	products := []ProductModel{}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return products, NewApiError(404, "Dalej już produktów nie ma")
		}

		slog.Error(err.Error())
		return products, InternalServerError()
	}

	for {
		if !rows.Next() {
			break
		}

		// wtf is that...
		// p.s. I didn't know that you can pass fields straight
		// to the Scan() function

		var title, description, ptype string
		var pid, createdAt int
		var price float64
		var uuid, ext, categoryTitle sql.NullString
		var iid, categoryId sql.NullInt32

		if err := rows.Scan(
			&pid, &title, &description, &price, &ptype, &createdAt, &iid, &uuid, &ext, &categoryId, &categoryTitle,
		); err != nil {
			slog.Error(err.Error())
			return products, InternalServerError()
		}

		if len(products) <= 0 || products[(len(products)-1)].ID != pid {
			p := ProductModel{
				ID:            pid,
				Title:         title,
				Description:   description,
				Price:         price,
				Type:          ptype,
				CreatedAt:     createdAt,
				CategoryID:    int(categoryId.Int32),
				CategoryTitle: categoryTitle.String,
				Images:        []ImageModel{},
			}

			products = append(products, p)
		}

		var p ProductModel = products[len(products)-1]

		if uuid.Valid {
			p.Images = append(p.Images, ImageModel{
				ID:   int(iid.Int32),
				Path: uuid.String + "." + ext.String,
			})
		}

		products[len(products)-1] = p
	}

	return products, nil
}

type GuitarModel struct {
	StringCount  int    `json:"stringsCount"`
	Bodyshape    string `json:"bodyshape"`
	Color        string `json:"color"`
	PickupsCount int    `json:"pickupsCount"`
}

type PickupModel struct {
	Type string `json:"type"`
}

type AmplifierModel struct {
	Power int `json:"power"`
}

type CraftedModel struct {
	BodyshapeID int      `json:"bodyshapeId"`
	Color       string   `json:"color"`
	Pickups     []Pickup `json:"pickups"`
}

type BodyshapeModel struct {
	Color        string `json:"color"`
	Material     string `json:"material"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Thickness    int    `json:"thickness"`
	PickupsCount int    `json:"pickupsCount"`
}

type ProductModel struct {
	ID            int             `json:"id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	Type          string          `json:"type"`
	Price         float64         `json:"price"`
	CategoryID    int             `json:"category,omitempty"`
	CategoryTitle string          `json:"categoryTitle,omitempty"`
	CreatedAt     int             `json:"createdAt"`
	Images        []ImageModel    `json:"images"`
	Guitar        *GuitarModel    `json:"guitar,omitempty"`
	Pickup        *PickupModel    `json:"pickup,omitempty"`
	Bodyshape     *BodyshapeModel `json:"bodyshape,omitempty"`
	Amplifier     *AmplifierModel `json:"amplifier,omitempty"`
	Crafted       *CraftedModel   `json:"crafted,omitempty"`
}

func (p Product) Find(id int) (ProductModel, error) {
	products, err := p.FindMany(FindManyParams{
		Limit:     1,
		ProductID: id,
	})

	if err != nil {
		return ProductModel{}, err
	}

	if len(products) <= 0 {
		return ProductModel{}, NewApiError(404, "Nie znajzdiono produktu")
	}

	product := &products[0]

	switch product.Type {
	case "guitar":
		row := p.conn.QueryRow(
			context.Background(),
			"select g.strings_count, g.bodyshape, g.color, g.pickups_count from products p join guitars g on g.product_id = p.id where p.id = $1",
			product.ID,
		)
		product.Guitar = &GuitarModel{}
		g := product.Guitar
		err = row.Scan(&g.StringCount, &g.Bodyshape, &g.Color, &g.PickupsCount)
		break
	case "pickup":
		row := p.conn.QueryRow(
			context.Background(),
			"select pi.type from products p join pickups pi on pi.product_id = p.id where p.id = $1",
			product.ID,
		)
		product.Pickup = &PickupModel{}
		err = row.Scan(&product.Pickup.Type)
		break
	case "bodyshape":
		row := p.conn.QueryRow(
			context.Background(),
			"select b.color::text, b.material, b.height, b.width, b.thickness, b.pickups_count from products p join bodyshapes b on b.product_id = p.id where p.id = $1",
			product.ID,
		)
		product.Bodyshape = &BodyshapeModel{}
		b := product.Bodyshape
		err = row.Scan(&b.Color, &b.Material, &b.Height, &b.Width, &b.Thickness, &b.PickupsCount)
		break
	case "amplifier":
		row := p.conn.QueryRow(
			context.Background(),
			"select a.power from products p join amplifiers a on a.product_id = p.id where p.id = $1",
			product.ID,
		)
		product.Amplifier = &AmplifierModel{}
		err = row.Scan(&product.Amplifier.Power)
		break
	case "crafted":
		rows, err := p.conn.Query(
			context.Background(),
			"select c.bodyshape_id, c.color, cp.pickup_id, cp.position from products p join constructors c on c.product_id = p.id join constructors_pickups cp on cp.constructor_id = c.id where p.id = $1",
			product.ID,
		)

		if err != nil {
			slog.Error(err.Error())
			return *product, InternalServerError()
		}

		product.Crafted = &CraftedModel{}
		c := product.Crafted

		for {
			if !rows.Next() {
				break
			}

			p := Pickup{}

			if err := rows.Scan(&c.BodyshapeID, &c.Color, &p.PickupID, &p.Position); err != nil {
				slog.Error(err.Error())
				return *product, InternalServerError()
			}

			c.Pickups = append(c.Pickups, p)
		}

		break
	}

	if err != nil {
		slog.Error(err.Error())
		return ProductModel{}, err
	}

	return *product, nil
}

package models

import (
	"context"
	"log/slog"
	"slices"
)

type Pickup struct {
	PickupID int    `json:"pickupId"`
	Position string `json:"position"`
}

type ConstructorCreate struct {
	Title       string   `json:"title"`
	BodyshapeID int      `json:"bodyshapeId"`
	Pickups     []Pickup `json:"pickups"`
	Color       string   `json:"color"`
}

var Positions = []string{"top", "middle", "bottom"}

// Tworzy produkt o typu "crafted", podlicza automatycznie cenę na podstawie wprowadzonych
// danych o przetwornika itd. Zwraza identyfikator swieżostworzonego produktu.
func (c Product) CreateConstructor(params ConstructorCreate) (int, error) {
	deposit := 1000
	description := "Guitara stworzona przez constuctor"

	pickupIDs := []int{}
	positions := []string{}

	for _, p := range params.Pickups {
		if !slices.Contains(Positions, p.Position) {
			return 0, NewApiError(400, "Nieprawidlowa pozycja dla przetwornika (pickup). Prawidlowymi są: top, middle, bottom.")
		}

		if slices.Contains(positions, p.Position) {
			return 0, NewApiError(400, "Pozycji nie mogą się powtarzacz w dwóch przetwornikach")
		}

		positions = append(positions, p.Position)
		pickupIDs = append(pickupIDs, p.PickupID)
	}

	// 	😄 loud and clear.
	// Score now: 8.5 / 10
	//
	// - ChatGPT (12.19.2025)

	row := c.conn.QueryRow(
		context.Background(),
		`
WITH pkps AS (
	SELECT
        pi.id,
        p.price,
        row_number() over () AS rn
    FROM
		unnest ($1::bigint[]) AS ids(id)
        JOIN pickups pi ON pi.id = ids.id
        JOIN products p ON p.id = pi.product_id

),
prds AS (
    SELECT
        gen_random_uuid(),
        price
    FROM
        pkps
    UNION
    (
        SELECT
            gen_random_uuid(),
            p.price
        FROM
            bodyshapes b
            JOIN products p ON p.id = b.product_id
        WHERE
            b.id = $2
    )
),
prd AS (
    INSERT INTO
        products (TYPE, title, description, price, publish)
    VALUES
        (
            'crafted',
            $3,
            $4,
            (
                SELECT
                    SUM(price) + $5
                FROM
                    prds
            ),
            false
        ) returning id
),
con AS (
    INSERT INTO
        constructors (product_id, bodyshape_id, color)
    VALUES
        (
            (
                SELECT
                    id
                FROM
                    prd
            ),
            $2,
            $6
        ) returning id
),
cp AS (
    INSERT INTO
        constructors_pickups (constructor_id, pickup_id, position)
    SELECT
        c.id,
        pkps.id,
        pos.position
    FROM
        con c
        JOIN pkps ON TRUE
        JOIN (
            SELECT
                position,
                row_number() over () AS rn
            FROM
                unnest (
                    $7 :: pickup_pos []
                ) AS position
        ) pos ON pos.rn = pkps.rn
)
SELECT
    id
FROM
    prd;

		`,
		pickupIDs,
		params.BodyshapeID,
		params.Title,
		description,
		deposit,
		params.Color,
		positions,
	)

	var pid int

	if err := row.Scan(&pid); err != nil {
		slog.Error(err.Error())
		return 0, InternalServerError()
	}

	return pid, nil
}

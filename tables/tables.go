package tables

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func Init(conn *pgx.Conn) {
	createTable(conn, Enums)
	createTable(conn, Domains)

	createTable(conn, ImagesTable)

	createTable(conn, CategoriesTable)

	createTable(conn, ProductsTable)
	createTable(conn, ProductsImagesTable)
	createTable(conn, ProductsCategoriesTable)

	createTable(conn, GuitarsTable)
	createTable(conn, AmplifiersTable)
	createTable(conn, BodyshapesTable)
	createTable(conn, PickupsTable)

	createTable(conn, ConstructorsTable)
	createTable(conn, ConstructorsPickupsTable)

	createTable(conn, CartsTable)
	createTable(conn, CartsItemsTable)

	createTable(conn, OrdersTable)
	createTable(conn, OrdersItemsTable)

}

func createTable(conn *pgx.Conn, query string) {
	_, err := conn.Exec(context.Background(), query)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code != "42710" {
				slog.Error(
					fmt.Sprintf(
						"While trying to create a scheme error occured. Query: %s",
						query,
					),
				)
				slog.Error(err.Error())
				os.Exit(1)
			}

		} else {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}
}

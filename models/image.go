package models

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Image struct {
	conn *pgx.Conn
}

func NewImage(conn *pgx.Conn) Image {
	return Image{conn: conn}
}

// tworzy zdjęcie i zwraca go identyfikator
func (i Image) Create(uuid string, ext string) (int, error) {
	row := i.conn.QueryRow(
		context.Background(),
		`
		INSERT INTO images (uuid, ext) VALUES ($1, $2) RETURNING id
		`,
		uuid,
		ext,
	)

	var id int
	err := row.Scan(&id)

	if err != nil {
		slog.Error(err.Error())
		return 0, InternalServerError()
	}

	return id, nil
}

// tworzy zdjęcie i dodaje go do produktu z podanym identyfikatorem
func (i Image) ProductCreate(uuid string, ext string, productId string) error {
	imageId, err := i.Create(uuid, ext)

	if err != nil {
		return err
	}

	_, err = i.conn.Exec(
		context.Background(),
		`INSERT INTO products_images (product_id, image_id) VALUES ($1, $2)`,
		productId,
		imageId,
	)

	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.SQLState() == "23503" {
				return NewApiError(400, "Produkt z podanym identyfikatorem nie istnieje.")
			}
		}

		slog.Error(err.Error())
		return InternalServerError()
	}

	return nil
}

// tworzy zdjęcie i ustanawia go dla kategorii z podanym identyfikatorem
func (i Image) CategoryCreate(uuid string, ext string, categoryId string) error {
	imageId, err := i.Create(uuid, ext)

	if err != nil {
		return err
	}

	_, err = i.conn.Exec(
		context.Background(),
		`UPDATE products_images SET image_id = $1 WHERE id = $2`,
		imageId,
		categoryId,
	)

	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.SQLState() == "23503" {
				return NewApiError(400, "Kategoria z podanym identyfikatorem nie istnieje.")
			}
		}

		slog.Error(err.Error())
		return InternalServerError()
	}

	return nil
}

// usuwa zdjęcie z podanym identyfikatorem i zwraca nazwę pliku
// np. "ce900cc6-64f0-473b-b7e9-34c3cdf478b5.png";
// dodatkowo usuwa relacji dla produktów i katetegorii
func (i Image) Delete(id string) (string, error) {
	rows, err := i.conn.Query(
		context.Background(),
		"DELETE from images WHERE id = $1 RETURNING uuid, ext",
		id,
	)

	defer rows.Close()

	if err != nil {
		slog.Error(err.Error())
		return "", InternalServerError()
	}

	if !rows.Next() {
		return "", NewApiError(400, "Nie ma zdjęcia z podanym identyfikatorem.")
	}

	var uuid, ext string

	err = rows.Scan(&uuid, &ext)

	if err != nil {
		slog.Error(err.Error())
		return "", InternalServerError()
	}

	return uuid + "." + ext, nil
}

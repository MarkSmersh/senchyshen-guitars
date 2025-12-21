package api

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"mime/multipart"
	"os"
	"slices"
	"strings"

	"github.com/MarkSmersh/senchyshen-guitars/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) ImagesProductsPost(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		slog.Error(err.Error())
		err := models.InternalServerError()
		c.String(err.Code(), err.Error())
		return
	}

	pid := c.Param("pid")

	images := form.File["images"]
	model := models.NewImage(s.Conn)

	for i, img := range images {
		uuid, ext, err := saveImage(img, c.SaveUploadedFile)

		if err != nil {
			slog.Error(err.Error())
			c.String(400, "Zdjęcie #%d. %s", i, err.Error())
			return
		}

		err = model.ProductCreate(uuid, ext, pid)

		var apierr models.ApiError

		if err != nil {
			if errors.As(err, &apierr) {
				c.String(apierr.Code(), "Zdjęcie #%d. %s", i, apierr.Error())
				return
			}

			apierr = models.InternalServerError()
			c.String(apierr.Code(), apierr.Error())
			return
		}
	}

	c.String(201, "Do wybranego produkto było dodano %d zdjęć.", len(images))
}

func (s *Server) ImagesCategoriesPost(c *gin.Context) {
	img, err := c.FormFile("image")

	if err != nil {
		slog.Error(err.Error())
		err := models.InternalServerError()
		c.String(err.Code(), err.Error())
		return
	}

	cid := c.Param("cid")
	model := models.NewImage(s.Conn)

	uuid, ext, err := saveImage(img, c.SaveUploadedFile)

	if err != nil {
		slog.Error(err.Error())
		c.String(400, err.Error())
		return
	}

	err = model.CategoryCreate(uuid, ext, cid)

	var apierr models.ApiError

	if err != nil {
		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		apierr = models.InternalServerError()
		c.String(apierr.Code(), apierr.Error())
		return
	}

	c.String(201, "Zdjęcie było zmienione dla danej kategorii.")
}

func (s *Server) ImagesDelete(c *gin.Context) {
	model := models.NewImage(s.Conn)

	id := c.Param("id")

	slog.Info(id)

	filename, err := model.Delete(id)

	if err != nil {
		var apierr models.ApiError

		if errors.As(err, &apierr) {
			c.String(apierr.Code(), apierr.Error())
			return
		}

		slog.Error(err.Error())
	}

	err = os.Remove("assets/" + filename)

	if err != nil {
		c.String(500, "Niemożliwe usunąć zdjęcie.")
	}

	c.String(200, "Zdjęcie jest usunięte.")
}

func saveImage(
	img *multipart.FileHeader,
	saveFunc func(*multipart.FileHeader, string, ...fs.FileMode) error,
) (string, string, error) {
	filename := strings.Split(img.Filename, ".")

	if len(filename) <= 1 {
		return "", "", errors.New("W nazwie pliku jest brak roszerzenia.")
	}

	ext := filename[1]

	acceptedExtension := []string{"webp", "jpg", "png", "jpeg"}

	if !slices.Contains(acceptedExtension, ext) {
		return "", "", errors.New(
			fmt.Sprintf(
				"Rozszerzenie nie jest przyjmowane. Dozwolone roszerzenia to: %s",
				strings.Join(acceptedExtension, ", "),
			),
		)
	}

	uuid := uuid.New()

	err := saveFunc(img, "assets/"+uuid.String()+"."+ext)

	if err != nil {
		return "", "", errors.New("Niemożliwe stworyć plik.")
	}

	return uuid.String(), ext, nil
}

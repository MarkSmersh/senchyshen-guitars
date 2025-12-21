package utils

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"os"
	"slices"
	"strings"

	"github.com/google/uuid"
)

// writes a file and returns it's generated uuid and extension
func SaveFormImage(header multipart.FileHeader, dir string) (string, string, error) {
	f, _ := header.Open()

	data, _ := io.ReadAll(f)

	filename := strings.Split(header.Filename, ".")

	if len(filename) <= 1 {
		return "", "", errors.New("W nazwie pliku jest brak roszerzenia.")
	}

	ext := filename[1]

	acceptedExtension := []string{"webp", "jpg", "png", "jpeg"}

	if !slices.Contains(acceptedExtension, ext) {
		return "", "", errors.New(
			fmt.Sprintf(
				"Rozszerzenie nie jest przyjmowany. Dozwolone roszenia to: %s",
				strings.Join(acceptedExtension, ", "),
			),
		)
	}

	uuid := uuid.New()

	err := os.WriteFile(
		fmt.Sprintf("%s/%s.%s", dir, uuid.String(), ext),
		data,
		os.FileMode(0666),
	)

	if err != nil {
		slog.Error(err.Error())
		return "", "", errors.New("Niemożliwe zapisać plik.")
	}

	return uuid.String(), ext, nil
}

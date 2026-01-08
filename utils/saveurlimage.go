package utils

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/google/uuid"
)

// zapisuje zdjęcie do podanego dyrektorium i zwraca nazwę pliku oraz go roszerzenie
func SaveUrlImage(url string, dir string) (string, string, error) {
	res, err := http.Get(url)

	splittedUrl := strings.Split(url, ".")
	splittedEnd := strings.Split(splittedUrl[len(splittedUrl)-1], "?")
	ext := splittedEnd[0]

	if !slices.Contains(AcceptedExtension, ext) {
		return "", "", errors.New(
			fmt.Sprintf(
				"Nie można pobrać zdjęcia. Dozwolone rozszerzenia to: %s",
				strings.Join(AcceptedExtension, ", "),
			),
		)
	}

	if err != nil {
		return "", "", errors.New("Niemożliwie dostać zdjęcie z podanego URL")
	}

	uuid := uuid.New()

	out, _ := os.Create(fmt.Sprintf("%s/%s.%s", dir, uuid.String(), ext))

	_, err = io.Copy(out, res.Body)

	if err != nil {
		slog.Error(err.Error())
		return "", "", errors.New("Niemożliwe zapisać plik")
	}

	return uuid.String(), ext, nil
}

package api

import (
	"net/http"
	"os"
)

type AssetsFS struct {
	fs http.FileSystem
}

// inspired by yt's cdn
func (fs AssetsFS) Open(name string) (http.File, error) {
	file, err := fs.fs.Open(name)

	if err != nil {
		file, err = os.Open("assets/cc.jpg")

		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func (s *Server) AssetsRouter() {
	s.Engine.StaticFS("/assets", AssetsFS{fs: http.Dir("assets")})
}

package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeStaticFile(w http.ResponseWriter, r *http.Request, baseDir string, fileServer http.Handler) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	fullpath := filepath.Join(baseDir, path)

	if strings.HasPrefix(r.URL.Path, "/api/") {
		http.NotFound(w, r)
		return
	}

	_, err := os.Stat(fullpath)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fileServer.ServeHTTP(w, r)
}

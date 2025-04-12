package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

const SERVE_MSG = `%sServing slide show from %s%s%s on%s:%s

    %sLocal%s: %shttp://localhost:%d%s
    %sNetwork%s: %shttp://10.0.0.16:%d%s

%sPress %sCtrl+C%s to stop server%s
`

func CreateServer(filePath string, port int, customThemes CustomThemesFlag) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	execDir := path.Dir(execPath)
	filePath = strings.Replace(filePath, "~", homeDir, 1)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("notebook at '%s' does not exist", filePath)
	}

	// Frontend
	frontendDir := path.Join(execDir, "../frontend")
	frontendFs := http.FileServer(http.Dir(frontendDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fePath := path.Join(frontendDir, r.URL.Path)

		// Notebook assets
		if _, err := os.Stat(fePath); os.IsNotExist(err) {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
			http.ServeFile(w, r, path.Join(path.Dir(filePath), r.URL.Path))
			return
		}
		// Frontend assets
		frontendFs.ServeHTTP(w, r)
	})

	// Themes
	themes, err := customThemes.ToArray()
	if err != nil {
		return fmt.Errorf("failed to load themes: %v", err)
	}
	http.HandleFunc("/assets/custom-themes.json", func(w http.ResponseWriter, r *http.Request) {
		if themes == nil {
			themes = []map[string]any{}
		}
		json.NewEncoder(w).Encode(themes)
	})

	// Notebook
	http.HandleFunc("/notebook.ipynb", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})

	fmt.Printf(SERVE_MSG,
		ansiGreen+ansiBold, ansiYellow, path.Base(filePath),
		ansiGreen, ansiReset+ansiBold+ansiDim, ansiReset,

		ansiBold, ansiDim, ansiReset+ansiCyan, port, ansiReset,
		ansiBold, ansiDim, ansiReset+ansiCyan, port, ansiReset,

		ansiYellow, ansiCyan, ansiYellow, ansiReset,
	)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

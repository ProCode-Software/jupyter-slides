package server

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	ansiBold   = "\033[1m"
	ansiReset  = "\033[0m"
	ansiYellow = "\033[33m"
	ansiGreen  = "\033[32m"
	ansiCyan   = "\033[36m"
	ansiDim    = "\033[2m"
)

func CreateServer(filePath string, port int) error {
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
			http.ServeFile(w, r, path.Join(path.Dir(filePath), r.URL.Path))
            return
		}
		// Frontend assets
		frontendFs.ServeHTTP(w, r)
	})

	// Notebook
	http.HandleFunc("/notebook.ipynb", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})

	fmt.Printf(`%sServing slide show from %s%s%s on%s:%s

    %sLocal%s: %shttp://localhost:%d%s
    %sNetwork%s: %shttp://10.0.0.16:%d%s

%sPress %sCtrl+C%s to stop server%s
`,
		ansiGreen+ansiBold, ansiYellow, path.Base(filePath),
		ansiGreen, ansiReset+ansiBold+ansiDim, ansiReset,
		ansiBold, ansiDim, ansiReset+ansiCyan, port, ansiReset,
		ansiBold, ansiDim, ansiReset+ansiCyan, port, ansiReset,
		ansiYellow, ansiCyan, ansiYellow, ansiReset,
	)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}

	return nil
}

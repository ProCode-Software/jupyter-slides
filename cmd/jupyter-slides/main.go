package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "0.0.1"
const DEFAULT_PORT = 8080

var (
	ansiBold   = "\033[1m"
	ansiReset  = "\033[0m"
	ansiRed    = "\033[31m"
	ansiYellow = "\033[33m"
	ansiGreen  = "\033[32m"
	ansiCyan   = "\033[36m"
	ansiBlue   = "\033[34m"
	ansiDim    = "\033[2m"
)

var helpString = fmt.Sprintf(
	`%sjupyter-slides %s(version %s)%s
%sA Jupyter Notebook slide show converter%s

%sUsage%s: %sjupyter-slides [options] <file>%s

%sOptions%s:%s
%s
%s

%shttps://github.com/ProCode-Software/jupyter-slides%s`,
	ansiBold+ansiYellow, ansiReset+ansiDim, VERSION, ansiReset,
	ansiGreen, ansiReset,
	ansiBold, ansiDim, ansiReset+ansiCyan, ansiReset,
	ansiBold, ansiDim, ansiReset,
	option(
		"--port"+param("port")+"       ",
		fmt.Sprintf(
			"The localhost port to serve the slideshow on %s(default: %d)%s",
			ansiDim, DEFAULT_PORT, ansiReset,
		),
	),
	option("--help", "Show this message"),
	ansiBlue, ansiReset,
)

func option(label string, description string) string {
	return fmt.Sprintf(ansiYellow+"  %-20s"+ansiReset+"%s", label, description)
}

func param(text string) string {
	return fmt.Sprintf("%s %s(%s)%s", ansiReset, ansiDim, text, ansiReset)
}

func main() {
	portFlag := flag.Int("port", 8000, "The localhost port to serve the slideshow on")
	helpFlag := flag.Bool("help", false, "Show this help message")
	flag.Parse()

	port := *portFlag
	filePath := flag.Arg(0)

	if *helpFlag || filePath == "" || filePath == "help" {
		fmt.Println(helpString)
		os.Exit(1)
	}

	if error := CreateServer(filePath, port); error != nil {
		fmt.Fprintf(
			os.Stderr,
			"%sError%s:%s Failed to start server: %s%s\n",
			ansiBold+ansiRed, ansiReset+ansiBold+ansiDim, ansiReset+ansiBold, ansiReset, error,
		)
		os.Exit(1)
	}
}

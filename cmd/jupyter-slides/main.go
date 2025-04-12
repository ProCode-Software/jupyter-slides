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
    ansiMagenta = "\033[35m"
	ansiBlue   = "\033[34m"
	ansiDim    = "\033[2m"
)

const HELP_FORMAT = `%sjupyter-slides %s(version %s)%s
%sA Jupyter Notebook slide show converter%s

%sUsage%s: %sjupyter-slides [options] <file>%s

%sOptions%s:%s
%s
%s
%s

%shttps://github.com/ProCode-Software/jupyter-slides%s`

var helpString = fmt.Sprintf(
	HELP_FORMAT,
	ansiBold+ansiYellow, ansiReset+ansiDim, VERSION, ansiReset,
	ansiGreen, ansiReset,
	ansiBold, ansiDim, ansiReset+ansiCyan, ansiReset,
	ansiBold, ansiDim, ansiReset,
	option("--port", "port", "            The localhost port to serve the slideshow on"+
		par(fmt.Sprintf("default: %d", DEFAULT_PORT))),
	option("--custom-theme", "path", "    Load a custom VSCode color theme"),
	option("--help", "", "Show this message"),
	ansiMagenta, ansiReset,
)

func option(label, parameter, description string) string {
    if parameter != "" {
        parameter = param(parameter)
    }
	return fmt.Sprintf(
		ansiYellow+"  %-25s"+ansiReset+"%s",
		fmt.Sprintf("%s%s", label, parameter),
		description,
	)
}

func par(text string) string {
	return fmt.Sprintf("%s %s(%s)%s", ansiReset, ansiDim, text, ansiReset)
}
func param(text string) string {
	return fmt.Sprintf("%s %s<%s>%s", ansiReset, ansiDim, text, ansiReset)
}

func main() {
	var customThemesFlag CustomThemesFlag
	portFlag := flag.Int("port", 8000, "The localhost port to serve the slideshow on")
	helpFlag := flag.Bool("help", false, "Show this help message")
	flag.Var(&customThemesFlag, "custom-theme", "Load a custom VSCode color theme")
	flag.Parse()

	port := *portFlag
	filePath := flag.Arg(0)

	if *helpFlag || filePath == "" || filePath == "help" {
		fmt.Println(helpString)
		os.Exit(2)
	}

	if err := CreateServer(filePath, port, customThemesFlag); err != nil {
		Error("Failed to start server", err)
	}
}

func Error(msg string, err error) {
	fmt.Fprintf(
		os.Stderr,
		"%sError%s:%s %s: %s%s\n",
		ansiBold+ansiRed, ansiReset+ansiBold+ansiDim, ansiReset+ansiBold, msg, ansiReset, err,
	)
	os.Exit(1)
}

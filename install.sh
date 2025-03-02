#!/usr/bin/env bash

JS_DIR=${JS_DIR:-"$HOME/.jupyter-slides"}
SCRIPT_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

function red() {
    echo -e "$(tput setaf 1)$1$(tput sgr0)"
}
function green() {
    echo -e "$(tput setaf 2)$1$(tput sgr0)"
}
function yellow() {
    echo -e "$(tput setaf 3)$1$(tput sgr0)"
}
function cyan() {
    echo -e "$(tput setaf 6)$1$(tput sgr0)"
}
function dim() {
    echo -e "$(tput dim)$1$(tput sgr0)"
}
function bold() {
    echo -e "$(tput bold)$1$(tput sgr0)"
}
function error() {
    text=$1
    echo -e "\033[1m$(red Error)\033[1m$(dim :)\033[0m $text"
    exit 1
}

function check-if-command-exists() {
    command=$1
    error=$2
    if ! which "$command" > /dev/null; then
        error "$error"
    fi
}

bold "$(yellow "Installing jupyter-slides to $(cyan "$JS_DIR")")"

mkdir -p "$JS_DIR"
rm -rf "${JS_DIR:?}/"* || exit
cd "$JS_DIR" || exit

check-if-command-exists go "Go is not installed"

echo -n "$(bold "💻 Building command line$(dim ...)")"
error=$(go build -C "$SCRIPT_DIR/cmd" -o "$JS_DIR/bin/jupyter-slides" main.go 2>&1)
if [ $? -ne 0 ]; then
    red "$(bold Failed)"
    error "Failed to build command line: $error"
fi
green "$(bold Done!)"

check-if-command-exists pnpm "pnpm is not installed"

echo -n "$(bold "🌐 Building frontend$(dim ...)")"
error=$(
    cd "$SCRIPT_DIR" || exit
    pnpm install 2>&1
    pnpm vite build --outDir "$JS_DIR/frontend" 2>&1
)
if [ $? -ne 0 ]; then
    red "$(bold Failed)"
    error "Failed to build frontend: $error"
fi
green "$(bold Done!)"
green "$(bold "\n✅ jupyter-slides has been installed successfully!")"
yellow "Next, add $(cyan "$JS_DIR/bin")""$(yellow " to PATH for easy access to jupyter-slides")"

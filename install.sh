#!/usr/bin/env bash
# shellcheck disable=SC2069,SC2181

JS_DIR=${JS_DIR:-"$HOME/.jupyter-slides"}
SCRIPT_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

function red() {
    echo -e "\033[31m$1\033[0m"
}
function green() {
    echo -e "\033[32m$1\033[0m"
}
function yellow() {
    echo -e "\033[33m$1\033[0m"
}
function cyan() {
    echo -e "\033[36m$1\033[0m"
}
function dim() {
    echo -e "\033[2m$1\033[0m"
}
function bold() {
    echo -e "\033[1m$1\033[0m"
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
error=$(go build -C "$SCRIPT_DIR" -o "$JS_DIR/bin/jupyter-slides" ./cmd/jupyter-slides 2>&1 >/dev/null)
if [ $? -ne 0 ]; then
    red "Failed"
    error "Failed to build command line: $error"
fi
green "Done!"

check-if-command-exists pnpm "pnpm is not installed"

echo -n "$(bold "🌐 Building frontend$(dim ...)")"
error=$(
    cd "$SCRIPT_DIR" || exit
    pnpm install
    pnpm vite build --outDir "$JS_DIR/frontend" --emptyOutDir
) 2>&1 >/dev/null
if [ $? -ne 0 ]; then
    red "Failed"
    error "Failed to build frontend: $error"
fi
green "Done!"
green "$(bold "\n✅ jupyter-slides has been installed successfully!")"
yellow "Next, add $(cyan "$JS_DIR/bin")""$(yellow " to PATH for easy access to jupyter-slides.")"

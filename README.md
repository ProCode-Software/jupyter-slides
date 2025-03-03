# jupyter-slides
[![Jupyter](https://img.shields.io/badge/Jupyter-%23F37626?logo=jupyter&logoColor=fff&style=for-the-badge)](#)
[![Vue.js](https://img.shields.io/badge/Vue.js-4FC08D?logo=vuedotjs&logoColor=fff&style=for-the-badge)](#)
[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white&style=for-the-badge)](#)
[![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=fff&style=for-the-badge)](#)
[![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?logo=typescript&logoColor=fff&style=for-the-badge)](#)

🪐 A tool to convert Jupyter Notebooks to slide shows.

## Features
- 📖 Frontend written in [Vue.js](https://vuejs.org) and TypeScript
- 💻 Backend written in Go
- 🖼️ Supports loading images
- 📄 Fast Markdown parsing with [markdown-it](https://github.com/markdown-it/markdown-it)
- 🎨 Code highlighting with [Shiki](https://shiki.style) with customizable theme

## Installation
[Go](https://go.dev) and [pnpm](https://pnpm.io) are required
1. Clone the repository
```shell
git clone https://github.com/ProCode-Software/jupyter-slides.git
cd jupyter-slides
```
2. Run the `./install.sh` script to build from source and install jupyter-slides into `~/.jupyter-slides`. To install into a different location, run `JS_DIR=folder ./install.sh`

## Usage
```
jupyter-slides [--port <port>] <notebook>
```
- **port:** Localhost port number to host the server on
- **notebook:** (Required) Path to Jupyter Notebook

## Uninstallation
To uninstall jupyter-slides, just delete the `~/.jupyter-slides` folder.

## Development
Clone the repo:
```shell
git clone https://github.com/ProCode-Software/jupyter-slides.git
cd jupyter-slides
pnpm install
pnpm run dev
```
Create a `notebook.ipynb` file in the folder for testing. Then open `http://localhost:5173` to view (this uses Vite's dev server).

## License
MIT License
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
- 📄 Markdown parsing with [markdown-it](https://github.com/markdown-it/markdown-it)
- 🎨 Code highlighting with [Shiki](https://shiki.style) with customizable theme

## Installation
[Go](https://go.dev) and [pnpm](https://pnpm.io) are required
1. Clone the repository
```shell
git clone https://github.com/ProCode-Software/jupyter-slides.git
cd jupyter-slides
```
2. Run the `./install.sh` script to build from source and install jupyter-slides into `~/.jupyter-slides`

## Usage
```
jupyter-slides [--port <port>] <notebook>
```

## License
MIT License
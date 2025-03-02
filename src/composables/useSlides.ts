import type { Notebook, Slide } from '@/types'

export function useSlides(notebook: Notebook): Slide[] {
    const slides: Slide[] = []

    // Title slide
    if (notebook.metadata?.name) {
        const titleSlide: Slide = {
            title: notebook.metadata.name,
            isTitleSlide: true,
            cells: [{ type: 'markdown', content: notebook.metadata.author }],
        }
        slides.push(titleSlide)
    }

    for (const cell of notebook.cells) {
        if (cell.cell_type == 'code') {
            slides.at(-1).cells.push({
                type: 'code',
                language: cell.metadata.vscode.languageId,
                content: cell.source.join(''),
            })
            continue
        }
        let isCodeBlock = false
        for (const line of cell.source.join('').split('\n')) {
            // Markdown code blocks
            if (isCodeBlock) {
                if (line.startsWith('```')) {
                    isCodeBlock = false
                    continue
                }
                slides.at(-1).cells.at(-1).content += line + '\n'
                continue
            }
            const codeBlockMatch = line.match(/^```(.*)$/)
            if (codeBlockMatch) {
                isCodeBlock = true
                slides.at(-1).cells.push({
                    type: 'code',
                    language: codeBlockMatch[1],
                    content: '',
                })
                continue
            }
            // Headings
            const headingMatch = line.match(/^#+ (.*)$/)
            if (headingMatch) {
                slides.push({
                    title: headingMatch[1],
                    cells: [],
                })
                continue
            }
            // Others
            const lastCell = slides.at(-1).cells.at(-1)
            if (lastCell?.type == 'markdown') {
                lastCell.content += line + '\n'
            } else {
                slides.at(-1).cells.push({
                    type: 'markdown',
                    content: line + '\n',
                })
            }
        }
    }
    return slides
}

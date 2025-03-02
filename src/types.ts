export interface Notebook {
    cells: NotebookCell[]
    metadata?: {
        name?: string
        author?: string
        language_info?: {
            name: string
        }
    }
    nbformat?: number
    nbformat_minor?: number
}

export type NotebookCell = {
    source?: string[]
} & (NotebookCodeCell | NotebookMarkdownCell)

export interface NotebookCodeCell {
    cell_type: 'code'
    execution_count?: number
    metadata?: {
        vscode?: {
            languageId: string
        }
    }
    outputs?: []
}
export interface NotebookMarkdownCell {
    cell_type: 'markdown'
    metadata?: {}
}

// Slides
export interface Slide {
    title: string
    cells: SlideCell[]
    isTitleSlide?: boolean
}

type SlideCell = SlideCodeCell | SlideMarkdownCell

export interface SlideCodeCell {
    type: 'code'
    language: string
    content: string
}
export interface SlideMarkdownCell {
    type: 'markdown'
    content: string
}
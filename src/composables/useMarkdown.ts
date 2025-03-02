import { useSettingsStore } from '@/stores'
import MarkdownIt from 'markdown-it'
import { bundledThemes, createHighlighter } from 'shiki/bundle/web'

const md = MarkdownIt({
    html: true,
    linkify: true,
    typographer: true,
    breaks: true,
})

export function useMarkdown(text: string): string {
    const result = md.render(text)
    return result
}

const lightTheme = 'github-light'
const darkTheme = 'github-dark'

export async function useHighlighter() {
    return await createHighlighter({
        themes: Object.keys(bundledThemes),
        langs: ['javascript', 'html', 'css'],
    })
}
export function highlight(highlighter, code: string, language: string) {
    const { theme } = useSettingsStore()
    const result = highlighter.codeToHtml(code, {
        lang: language,
        themes: {
            light: theme || lightTheme,
            dark: theme || darkTheme,
        },
    })
    return result
}

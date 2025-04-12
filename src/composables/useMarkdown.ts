import { useSettingsStore } from '@/stores'
import MarkdownIt from 'markdown-it'
import { createHighlighter } from 'shiki/bundle/web'
import { getAllThemes } from './useThemes'

const md = MarkdownIt({
    html: true,
    linkify: true,
    typographer: true,
    breaks: true,
})

export function useMarkdown(text: string): string {
    return md.render(text)
}

export const lightTheme = 'github-light'
export const darkTheme = 'github-dark'

export type Highlighter = Awaited<ReturnType<typeof useHighlighter>>

export async function useHighlighter() {
    return await createHighlighter({
        themes: await getAllThemes(),
        langs: ['javascript', 'html', 'css'],
    })
}

export function highlight(highlighter: Highlighter, code: string, language: string) {
    const { theme } = useSettingsStore()
    return highlighter.codeToHtml(code, {
        lang: language,
        themes: {
            light: theme || lightTheme,
            dark: theme || darkTheme,
        },
    })
}

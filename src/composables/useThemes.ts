import { bundledThemes } from 'shiki/bundle/web'
import { darkTheme, lightTheme, type Highlighter } from './useMarkdown'

export async function getCustomThemes(): Promise<Record<string, any>[]> {
    try {
        const res = await fetch('/assets/custom-themes.json')
        const themes = await res.json()
        return themes ?? []
    } catch {
        return []
    }
}

export async function getAllThemes() {
    const customThemes = await getCustomThemes()
    return [...Object.keys(bundledThemes), ...customThemes]
}
export async function getAllThemeNames() {
    const customThemes = await getCustomThemes()
    return [
        ...Object.keys(bundledThemes),
        ...customThemes.map(({ name }) => name),
    ].sort((a: string, b: string) => a.toLowerCase().localeCompare(b.toLowerCase()))
}

interface ThemeColors {
    type: 'light' | 'dark'
    foreground: string
    background: string
    codeForeground?: string
    linkForeground?: string
}

export function getThemeColors(
    highlighter: Highlighter,
    themeName: string
): ThemeColors {
    if (!themeName)
        themeName = document.body.classList.contains('dark') ? darkTheme : lightTheme

    const { bg, settings, colors, type, fg } = highlighter.getTheme(themeName)

    const findScope = (...n: string[]) =>
        settings.find(({ scope }) => {
            for (const s of n) {
                if (scope?.includes(s)) return true
            }
        })?.settings.foreground

    return {
        type,
        foreground: fg,
        background: bg,
        codeForeground: findScope('string'),
        linkForeground:
            colors['textLink.foreground'] ??
            findScope('string.other.link', 'meta.link'),
    }
}

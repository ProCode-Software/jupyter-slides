import { useSettingsStore } from '@/stores'

export function setFont() {
    const { font, setFont } = useSettingsStore()
    const selection = prompt('Font:', font ?? '')
    selection !== null && setFont(selection)
}
export function setMonoFont() {
    const { monoFont, setMonoFont } = useSettingsStore()
    const selection = prompt('Monospace Font:', monoFont ?? '')
    selection !== null && setMonoFont(selection)
}
export function setStoredSettings() {
    const { setFont, setMonoFont, setTheme } = useSettingsStore()
    setFont(localStorage.getItem('font'))
    setMonoFont(localStorage.getItem('monoFont'))
    setTheme(localStorage.getItem('theme'))
}

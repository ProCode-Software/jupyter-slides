import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import type { Notebook } from './types'

export const useNotebookStore = defineStore('notebook', () => {
    const rawNotebook = ref('{}')
    const notebookName = ref('')
    const notebookExists = ref(false)
    const notebookData = computed<Notebook>(
        () => JSON.parse(rawNotebook.value) as Notebook
    )
    function loadNotebook() {
        fetch('/notebook.ipynb').then(async (res: Response) => {
            let text: string, json: Notebook
            try {
                console.log(res.status);
                
                text = await res.text()
                json = JSON.parse(text)
            } catch (e) {
                alert('Invalid Jupyter notebook')
                throw new Error('Invalid Jupyter notebook')
            }
            rawNotebook.value = text
            notebookName.value = json.metadata?.name ?? 'Untitled Notebook'
            notebookExists.value = true
        })
    }
    return {
        rawNotebook,
        loadNotebook,
        notebookData,
        notebookName,
        notebookExists,
    }
})

export const useSettingsStore = defineStore('font', () => {
    const font = ref(''),
        monoFont = ref(''),
        theme = ref(''),
        isSelectingTheme = ref(false)

    return {
        font,
        monoFont,
        theme,
        isSelectingTheme,
        showThemeSelection() {
            isSelectingTheme.value = true
        },
        hideThemeSelection() {
            isSelectingTheme.value = false
        },
        setTheme(value: string) {
            value === 'null' && (value = null)
            isSelectingTheme.value = false
            theme.value = value
            localStorage.setItem('theme', value)
        },
        setFont(value: string = '') {
            value === 'null' && (value = null)
            font.value = value
            document.body.style.setProperty(`--setting-font`, value)
            localStorage.setItem('font', value)
        },
        setMonoFont(value: string = '') {
            value === 'null' && (value = null)
            monoFont.value = value
            document.body.style.setProperty(`--setting-monoFont`, value)
            localStorage.setItem('monoFont', value)
        },
    }
})

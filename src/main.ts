import '@/scss/style.scss'
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import App from './App.vue'
import { setStoredSettings } from './composables/useSettings'

const app = createApp(App)
app.use(createPinia())
app.mount('#app')

const isDark = window.matchMedia('(prefers-color-scheme: dark)')
if (isDark.matches) document.body.classList.add('dark')
isDark.addEventListener(
    'change',
    e => e.matches && document.body.classList.add('dark')
)
setStoredSettings()

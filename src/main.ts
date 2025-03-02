import '@/scss/style.scss'
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import App from './App.vue'
import { setStoredSettings } from './composables/useSettings'

const app = createApp(App)
app.use(createPinia())
app.mount('#app')

setStoredSettings()
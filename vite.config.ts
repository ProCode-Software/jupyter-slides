import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

const resolved = (p: string) => resolve(import.meta.dirname, p)

// https://vite.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: { open: true },
    resolve: {
        alias: {
            '@': resolved('./src'),
            '@composables': resolved('./src/composables'),
            '@components': resolved('./src/components')
        }
    },
    css: {
        preprocessorOptions: {
            scss: {
                additionalData: `@use '@/scss/colors' as *;`
            }
        }
    }
})

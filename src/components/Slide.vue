<script setup lang="ts">
import { computed } from 'vue'
import { highlight, useHighlighter } from '@composables/useMarkdown'
import { useSettingsStore } from '@/stores'
import { getThemeColors } from '@composables/useThemes'
import type { Slide } from '@/types'
import Markdown from './Markdown.vue'

const highlighter = await useHighlighter()
const themeData = computed(() =>
    getThemeColors(highlighter, useSettingsStore().theme)
)
const style = computed(() => {
    const { background, linkForeground, codeForeground, type, foreground } =
        themeData.value
    document.body.classList.remove('light')
    document.body.classList.remove('dark')
    document.body.classList.add(type)
    return {
        '--bg': background,
        '--code': codeForeground,
        '--link': linkForeground,
        '--mix': type == 'dark' ? '#000' : '#fff',
        color: type == 'dark' ? 'rgb(var(--dark-fg))' : foreground,
    }
})

const { slide } = defineProps<{ slide: Slide }>()
</script>
<template>
    <div class="Slide TitleSlide" v-if="slide.isTitleSlide" :style>
        <Markdown :content="`# ${slide.title}`" class="SlideTitle" />
        <Markdown
            v-if="slide.cells[0].content"
            :content="`${slide.cells[0].content}`"
            class="SlideAuthor" />
    </div>
    <div
        class="Slide CenterSlide"
        v-else-if="
            slide.cells.length < 1 ||
            (slide.cells.length === 1 && slide.cells[0].content.trim() == '')
        "
        :style>
        <Markdown :content="`# ${slide.title}`" class="SlideTitle" />
    </div>
    <div class="Slide" v-else :style>
        <Markdown :content="`# ${slide.title}`" class="SlideTitle" />
        <div class="SlideContent">
            <template v-for="cell in slide.cells">
                <Markdown v-if="cell.type == 'markdown'" :content="cell.content" />
                <div
                    class="CodeBlock"
                    v-else-if="cell.type == 'code'"
                    v-html="
                        highlight(highlighter, cell.content, cell.language)
                            .replace(
                                /background-color:\s*#f{3,6}/gi,
                                'background-color: #FAFAFA'
                            )
                            .replace(
                                /(background-color|--shiki-dark-bg):\s*#0{3,6}/gi,
                                '$1: #151515'
                            )
                    "></div>
            </template>
        </div>
    </div>
</template>
<style scoped src="@/scss/markdown.scss" />

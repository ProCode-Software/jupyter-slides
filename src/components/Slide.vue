<script setup lang="ts">
import { highlight, useHighlighter } from '@/composables/useMarkdown';
import type { Slide } from '@/types';
import Markdown from './Markdown.vue';

const highlighter = await useHighlighter()

const { slide } = defineProps<{ slide: Slide }>()
</script>
<template>
    <div class="Slide TitleSlide"
        v-if="slide.isTitleSlide">
        <Markdown :content="`# ${slide.title}`" class="SlideTitle" />
        <Markdown v-if="slide.cells[0].content"
            :content="`${slide.cells[0].content}`" class="SlideAuthor" />
    </div>
    <div class="Slide CenterSlide" v-else-if="slide.cells.length < 1">
        <Markdown :content="`# ${slide.title}`" class="SlideTitle" />
    </div>
    <div class="Slide" v-else>
        <Markdown :content="`# ${slide.title}`" class="SlideTitle" />
        <div class="SlideContent">
            <template v-for="cell in slide.cells">
                <Markdown v-if="cell.type == 'markdown'"
                    :content="cell.content" />
                <div class="CodeBlock" v-else-if="cell.type == 'code'"
                    v-html="highlight(highlighter, cell.content, cell.language).replace(/background-color:\s?#f{3,6}/, 'background-color:#fafafa')" />
            </template>
        </div>
    </div>
</template>
<style scoped src="@/scss/markdown.scss" />
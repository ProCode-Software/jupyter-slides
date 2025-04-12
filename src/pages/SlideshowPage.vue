<script setup lang="ts">
import SlideshowControls from '@components/SlideshowControls.vue'
import { useSlides } from '@composables/useSlides'
import { useNotebookStore } from '@/stores'
import type { Slide as SlideType } from '@/types'
import Slide from '@components/Slide.vue'
import { onKeyStroke } from '@vueuse/core'
import { ref, watch } from 'vue'

const { notebookData } = useNotebookStore()

const currentSlide = ref(+(localStorage.getItem('lastSlide') ?? 0))

const slides: SlideType[] = useSlides(notebookData)

onKeyStroke(['ArrowLeft', 'ArrowRight', ' '], ({ key }) => {
    if (key == 'ArrowLeft' && currentSlide.value !== 0) {
        currentSlide.value--
    } else if (
        (key == 'ArrowRight' || key == ' ') &&
        currentSlide.value < slides.length - 1
    ) {
        currentSlide.value++
    }
})

watch(currentSlide, n => localStorage.setItem('lastSlide', n.toString()))
</script>
<template>
    <div class="SlideShowView" @keydown.arrow-down="() => console.log('yes')">
        <Suspense>
            <Slide :slide="slides[currentSlide]" />
        </Suspense>
    </div>
    <Suspense>
        <SlideshowControls :slides="slides" v-model="currentSlide" />
    </Suspense>
</template>
<style scoped lang="scss"></style>

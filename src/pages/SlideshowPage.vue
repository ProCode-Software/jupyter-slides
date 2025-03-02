<script setup lang="ts">
import { useSlides } from '@/composables/useSlides';
import Slide from '@components/Slide.vue'
import { useNotebookStore } from '@/stores';
import { onKeyStroke } from '@vueuse/core';
import { ref, watch } from 'vue';
import SlideshowControls from '@/components/SlideshowControls.vue';
import type { Slide as SlideType } from '@/types';

const { notebookName, notebookData } = useNotebookStore()

document.title = notebookName

const currentSlide = ref(+(localStorage.getItem('lastSlide') ?? 0))

const slides: SlideType[] = useSlides(notebookData)

onKeyStroke(['ArrowLeft', 'ArrowRight'], ({ key }) => {
    if (key == 'ArrowLeft' && currentSlide.value !== 0) {
        currentSlide.value--
    } else if (key == 'ArrowRight' && currentSlide.value < slides.length - 1) {
        currentSlide.value++
    }
})

watch(currentSlide, (n) => localStorage.setItem('lastSlide', n.toString()))

</script>
<template>
    <div class="SlideShowView"
         @keypress.arrow-down="() => console.log('yes')">
        <Suspense>
            <Slide :slide="slides[currentSlide]" />
        </Suspense>
    </div>
    <SlideshowControls :slides="slides"
                       v-model="currentSlide" />
</template>
<style scoped lang="scss"></style>
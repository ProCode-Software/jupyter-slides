<script setup lang="ts">
import SlideshowPage from './pages/SlideshowPage.vue';
import { useNotebookStore } from './stores';

const data = useNotebookStore()
data.loadNotebook()
</script>

<template>
    <SlideshowPage v-if="data.notebookExists" />
</template>

<style scoped></style>
